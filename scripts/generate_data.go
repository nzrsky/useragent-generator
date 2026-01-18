//go:build ignore

// This script downloads real user-agent data from Intoli and extracts
// version components to update pkg/useragent/data.go
//
// Run with: go run scripts/generate_data.go
package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
)

const (
	dataURL    = "https://raw.githubusercontent.com/intoli/user-agents/main/src/user-agents.json.gz"
	outputFile = "pkg/useragent/data.go"
	topN       = 20 // top N versions to keep per category
)

// UserAgent represents a single user-agent entry from Intoli dataset
type UserAgent struct {
	UserAgent string  `json:"userAgent"`
	Weight    float64 `json:"weight"`
}

// VersionWeight holds version string and its cumulative weight
type VersionWeight struct {
	Version string
	Weight  float64
}

// ExtractedData holds all extracted version data
type ExtractedData struct {
	ChromeVersions   []string
	FirefoxVersions  []string
	SafariVersions   []string
	EdgeVersions     []string
	IOSVersions      []string
	MacVersions      []string
	AndroidVersions  []string
	AndroidDevices   []AndroidDevice
	WindowsVersions  []string
	LinuxDesktops    []string
}

type AndroidDevice struct {
	Model string
	Build string
}

func main() {
	fmt.Println("Downloading user-agents data...")
	agents, err := downloadAndParse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading data: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Downloaded %d user-agents\n", len(agents))

	fmt.Println("Extracting version components...")
	data := extractVersions(agents)

	fmt.Println("Generating Go code...")
	if err := generateCode(data); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating code: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outputFile)
	printStats(data)
}

func downloadAndParse() ([]UserAgent, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(dataURL)
	if err != nil {
		return nil, fmt.Errorf("HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status %d", resp.StatusCode)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gzip reader failed: %w", err)
	}
	defer gz.Close()

	data, err := io.ReadAll(gz)
	if err != nil {
		return nil, fmt.Errorf("read failed: %w", err)
	}

	var agents []UserAgent
	if err := json.Unmarshal(data, &agents); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return agents, nil
}

var (
	// Chrome/131.0.6778.85
	chromeVersionRe = regexp.MustCompile(`Chrome/(\d+\.\d+\.\d+\.\d+)`)
	// Firefox/133.0
	firefoxVersionRe = regexp.MustCompile(`Firefox/(\d+\.\d+)`)
	// Version/17.4 Safari
	safariVersionRe = regexp.MustCompile(`Version/(\d+\.\d+(?:\.\d+)?) Safari`)
	// Edg/131.0.2903.51
	edgeVersionRe = regexp.MustCompile(`Edg/(\d+\.\d+\.\d+\.\d+)`)
	// iPhone OS 17_4_1 or CPU OS 17_4
	iosVersionRe = regexp.MustCompile(`(?:iPhone|CPU) OS (\d+_\d+(?:_\d+)?)`)
	// Mac OS X 14_2_1 or Mac OS X 10_15_7
	macVersionRe = regexp.MustCompile(`Mac OS X (\d+_\d+(?:_\d+)?)`)
	// Android 14
	androidVersionRe = regexp.MustCompile(`Android (\d+)`)
	// Windows NT 10.0
	windowsVersionRe = regexp.MustCompile(`Windows NT (\d+\.\d+)`)
	// Linux; Android 14; SM-S911B Build/UP1A.231005.007
	androidDeviceRe = regexp.MustCompile(`Android \d+; ([^)]+) Build/([A-Z0-9.]+)`)
	// X11; Linux x86_64 or X11; Ubuntu; Linux x86_64
	linuxDesktopRe = regexp.MustCompile(`\((X11; [^)]+)\)`)
)

func extractVersions(agents []UserAgent) ExtractedData {
	// Maps to accumulate weights for each version
	chromeWeights := make(map[string]float64)
	firefoxWeights := make(map[string]float64)
	safariWeights := make(map[string]float64)
	edgeWeights := make(map[string]float64)
	iosWeights := make(map[string]float64)
	macWeights := make(map[string]float64)
	androidWeights := make(map[string]float64)
	windowsWeights := make(map[string]float64)
	linuxWeights := make(map[string]float64)
	deviceWeights := make(map[string]float64) // "model|build" -> weight

	for _, ua := range agents {
		s := ua.UserAgent
		w := ua.Weight

		// Chrome (but not Edge, Opera, Samsung)
		if strings.Contains(s, "Chrome/") && !strings.Contains(s, "Edg/") &&
			!strings.Contains(s, "OPR/") && !strings.Contains(s, "SamsungBrowser") {
			if m := chromeVersionRe.FindStringSubmatch(s); m != nil {
				chromeWeights[m[1]] += w
			}
		}

		// Firefox
		if m := firefoxVersionRe.FindStringSubmatch(s); m != nil {
			firefoxWeights[m[1]] += w
		}

		// Safari (not Chrome-based)
		if strings.Contains(s, "Safari/") && !strings.Contains(s, "Chrome/") {
			if m := safariVersionRe.FindStringSubmatch(s); m != nil {
				safariWeights[m[1]] += w
			}
		}

		// Edge
		if m := edgeVersionRe.FindStringSubmatch(s); m != nil {
			edgeWeights[m[1]] += w
		}

		// iOS
		if m := iosVersionRe.FindStringSubmatch(s); m != nil {
			iosWeights[m[1]] += w
		}

		// macOS
		if m := macVersionRe.FindStringSubmatch(s); m != nil {
			macWeights[m[1]] += w
		}

		// Android version
		if m := androidVersionRe.FindStringSubmatch(s); m != nil {
			androidWeights[m[1]] += w
		}

		// Windows
		if m := windowsVersionRe.FindStringSubmatch(s); m != nil {
			windowsWeights[m[1]] += w
		}

		// Android device
		if m := androidDeviceRe.FindStringSubmatch(s); m != nil {
			model := strings.TrimSpace(m[1])
			build := m[2]
			// Skip generic "K" model
			if model != "K" && len(model) > 2 {
				key := model + "|" + build
				deviceWeights[key] += w
			}
		}

		// Linux desktop (X11; Linux x86_64, etc.)
		if strings.Contains(s, "X11;") && strings.Contains(s, "Linux") && !strings.Contains(s, "Android") {
			if m := linuxDesktopRe.FindStringSubmatch(s); m != nil {
				platform := m[1]
				// Remove Firefox rv: suffix
				if idx := strings.Index(platform, "; rv:"); idx != -1 {
					platform = platform[:idx]
				}
				linuxWeights[platform] += w
			}
		}
	}

	data := ExtractedData{
		ChromeVersions:  topVersions(chromeWeights, topN),
		FirefoxVersions: topVersions(firefoxWeights, topN),
		SafariVersions:  topVersions(safariWeights, topN),
		EdgeVersions:    topVersions(edgeWeights, topN),
		IOSVersions:     topVersions(iosWeights, topN),
		MacVersions:     topVersions(macWeights, 30),
		AndroidVersions: topVersions(androidWeights, 10),
		WindowsVersions: topVersions(windowsWeights, 10),
		LinuxDesktops:   topVersions(linuxWeights, 20),
		AndroidDevices:  topDevices(deviceWeights, 50),
	}

	// Add fallback data if extracted data is sparse
	data.LinuxDesktops = mergeUnique(data.LinuxDesktops, fallbackLinuxDesktops)
	data.MacVersions = mergeUnique(data.MacVersions, fallbackMacVersions)
	data.AndroidDevices = mergeDevices(data.AndroidDevices, fallbackAndroidDevices)

	return data
}

// Fallback data for sparse categories
var fallbackLinuxDesktops = []string{
	"X11; Linux x86_64",
	"X11; Linux i686",
	"X11; Linux aarch64",
	"X11; Ubuntu; Linux x86_64",
	"X11; Fedora; Linux x86_64",
	"X11; Debian; Linux x86_64",
	"X11; Arch Linux; Linux x86_64",
	"X11; CentOS; Linux x86_64",
}

var fallbackMacVersions = []string{
	"10_15_7", "11_0", "11_6", "12_0", "12_6", "13_0", "13_6", "14_0", "14_4", "15_0",
}

var fallbackAndroidDevices = []AndroidDevice{
	{"Pixel 6", "TQ3A.230901.001"},
	{"Pixel 7", "TQ3A.230901.001"},
	{"Pixel 7 Pro", "TQ3A.230901.001"},
	{"Pixel 8", "UQ1A.231205.015"},
	{"Pixel 8 Pro", "UQ1A.231205.015"},
	{"SM-S901B", "TP1A.220624.014"},  // Samsung S22
	{"SM-S908B", "TP1A.220624.014"},  // Samsung S22 Ultra
	{"SM-S911B", "TP1A.220624.014"},  // Samsung S23
	{"SM-S918B", "TP1A.220624.014"},  // Samsung S23 Ultra
	{"SM-S921B", "UP1A.231005.007"},  // Samsung S24
	{"SM-S928B", "UP1A.231005.007"},  // Samsung S24 Ultra
	{"SM-A536B", "TP1A.220624.014"},  // Samsung A53
	{"SM-A546B", "UP1A.231005.007"},  // Samsung A54
	{"SM-G998B", "TP1A.220624.014"},  // Samsung S21 Ultra
	{"ONEPLUS A6013", "QKQ1.190716.003"},
	{"IN2025", "RKQ1.211119.001"},         // OnePlus Nord
	{"CPH2451", "TP1A.220905.001"},        // OPPO Find X5
	{"M2101K6G", "TKQ1.221114.001"},       // Xiaomi 11T Pro
	{"2201116SG", "TKQ1.221114.001"},      // Xiaomi 12
	{"23049PCD8G", "UKQ1.231003.002"},     // Xiaomi 14
	{"RMX3363", "TP1A.220905.001"},        // Realme GT 2 Pro
	{"V2111", "TP1A.220624.014"},          // Vivo X70 Pro
	{"LE2125", "RKQ1.211119.001"},         // OnePlus 9 Pro
}

func mergeUnique(a, b []string) []string {
	seen := make(map[string]bool)
	for _, v := range a {
		seen[v] = true
	}
	result := append([]string{}, a...)
	for _, v := range b {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

func mergeDevices(a, b []AndroidDevice) []AndroidDevice {
	seen := make(map[string]bool)
	for _, d := range a {
		seen[d.Model] = true
	}
	result := append([]AndroidDevice{}, a...)
	for _, d := range b {
		if !seen[d.Model] {
			result = append(result, d)
			seen[d.Model] = true
		}
	}
	return result
}

func topVersions(weights map[string]float64, n int) []string {
	var vw []VersionWeight
	for v, w := range weights {
		vw = append(vw, VersionWeight{v, w})
	}

	// Sort by weight descending
	sort.Slice(vw, func(i, j int) bool {
		return vw[i].Weight > vw[j].Weight
	})

	// Take top N
	if len(vw) > n {
		vw = vw[:n]
	}

	// Sort by version string for stable output
	sort.Slice(vw, func(i, j int) bool {
		return vw[i].Version < vw[j].Version
	})

	result := make([]string, len(vw))
	for i, v := range vw {
		result[i] = v.Version
	}
	return result
}

func topDevices(weights map[string]float64, n int) []AndroidDevice {
	var vw []VersionWeight
	for v, w := range weights {
		vw = append(vw, VersionWeight{v, w})
	}

	sort.Slice(vw, func(i, j int) bool {
		return vw[i].Weight > vw[j].Weight
	})

	if len(vw) > n {
		vw = vw[:n]
	}

	// Sort by model name
	sort.Slice(vw, func(i, j int) bool {
		return vw[i].Version < vw[j].Version
	})

	result := make([]AndroidDevice, len(vw))
	for i, v := range vw {
		parts := strings.SplitN(v.Version, "|", 2)
		if len(parts) == 2 {
			result[i] = AndroidDevice{Model: parts[0], Build: parts[1]}
		}
	}
	return result
}

func printStats(data ExtractedData) {
	fmt.Printf("\nExtracted:\n")
	fmt.Printf("  Chrome versions:  %d\n", len(data.ChromeVersions))
	fmt.Printf("  Firefox versions: %d\n", len(data.FirefoxVersions))
	fmt.Printf("  Safari versions:  %d\n", len(data.SafariVersions))
	fmt.Printf("  Edge versions:    %d\n", len(data.EdgeVersions))
	fmt.Printf("  iOS versions:     %d\n", len(data.IOSVersions))
	fmt.Printf("  macOS versions:   %d\n", len(data.MacVersions))
	fmt.Printf("  Android versions: %d\n", len(data.AndroidVersions))
	fmt.Printf("  Windows versions: %d\n", len(data.WindowsVersions))
	fmt.Printf("  Linux desktops:   %d\n", len(data.LinuxDesktops))
	fmt.Printf("  Android devices:  %d\n", len(data.AndroidDevices))
}

const codeTemplate = `// Code generated by scripts/generate_data.go. DO NOT EDIT.
// Source: https://github.com/intoli/user-agents
// Generated: {{.Timestamp}}

package ua

// Chrome versions extracted from real usage data (sorted by popularity)
var chromeVersions = []string{
{{- range .Data.ChromeVersions}}
	"{{.}}",
{{- end}}
}

// Firefox versions extracted from real usage data
var firefoxVersions = []string{
{{- range .Data.FirefoxVersions}}
	"{{.}}",
{{- end}}
}

// Safari versions extracted from real usage data
var safariVersions = []string{
{{- range .Data.SafariVersions}}
	"{{.}}",
{{- end}}
}

// Edge versions extracted from real usage data
var edgeVersions = []string{
{{- range .Data.EdgeVersions}}
	"{{.}}",
{{- end}}
}

// Windows NT versions (sorted by popularity)
var windowsVersions = []string{
{{- range .Data.WindowsVersions}}
	"{{.}}",
{{- end}}
}

// macOS versions (underscore format: 14_2_1)
var macVersions = []string{
{{- range .Data.MacVersions}}
	"{{.}}",
{{- end}}
}

// Linux desktop platforms extracted from real usage data
var linuxDesktops = []string{
{{- range .Data.LinuxDesktops}}
	"{{.}}",
{{- end}}
}

// iOS versions (underscore format: 17_4_1)
var iosVersions = []string{
{{- range .Data.IOSVersions}}
	"{{.}}",
{{- end}}
}

// Android versions
var androidVersions = []string{
{{- range .Data.AndroidVersions}}
	"{{.}}",
{{- end}}
}

// Android devices extracted from real usage data (model, build ID)
var androidDevices = []struct {
	model string
	build string
}{
{{- range .Data.AndroidDevices}}
	{"{{.Model}}", "{{.Build}}"},
{{- end}}
}

// WebKit version (used in Safari)
const webkitVersion = "605.1.15"

// AppleWebKit version for Chrome/Edge
const appleWebKitChrome = "537.36"
`

func generateCode(data ExtractedData) error {
	tmpl, err := template.New("code").Parse(codeTemplate)
	if err != nil {
		return fmt.Errorf("template parse failed: %w", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("create file failed: %w", err)
	}
	defer f.Close()

	templateData := struct {
		Timestamp string
		Data      ExtractedData
	}{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	if err := tmpl.Execute(f, templateData); err != nil {
		return fmt.Errorf("template execute failed: %w", err)
	}

	return nil
}
