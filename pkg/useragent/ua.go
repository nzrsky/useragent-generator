// Package ua provides high-performance User-Agent string generation
// for web crawlers and scrapers.
//
// Features:
//   - Zero-allocation bot User-Agents (constants)
//   - Seed-based reproducible randomization
//   - Fast xorshift64 PRNG
//   - Desktop browsers: Chrome, Firefox, Safari, Edge
//   - Mobile: iOS Safari, Android Chrome, WebView
//   - Bots: Google, Bing, Yandex, Baidu, social, SEO
//
// Quick usage with global generator (auto-seeded from time):
//
//	ua.Chrome()        // random Chrome UA
//	ua.SafariIOS()     // random iOS Safari UA
//	ua.Random()        // random UA from any category
//
// For reproducible results, create your own generator:
//
//	g := ua.WithSeed(12345)
//	g.Chrome()  // same sequence every time
//
// Thread safety: Package-level functions are goroutine-safe.
// Individual Generator instances are NOT goroutine-safe.
package ua

import (
	"sync"
	"time"
)

var (
	globalGen  = newTimeSeeded()
	globalLock sync.Mutex
)

func newTimeSeeded() *Generator {
	return &Generator{
		rng: newXorshift64(uint64(time.Now().UnixNano())),
	}
}

// Seed sets a new seed for the global generator
// Useful for reproducible results in tests
func Seed(seed uint64) {
	globalLock.Lock()
	globalGen = WithSeed(seed)
	globalLock.Unlock()
}

// Generator generates random but reproducible User-Agent strings.
// Individual Generator instances are NOT goroutine-safe.
// For concurrent use, create separate generators per goroutine.
type Generator struct {
	rng *xorshift64
}

// New creates a new Generator with a time-based seed
// For reproducible results, use WithSeed instead
func New() *Generator {
	return newTimeSeeded()
}

// WithSeed creates a new Generator with a specific seed
// Same seed produces same sequence of User-Agents
func WithSeed(seed uint64) *Generator {
	return &Generator{
		rng: newXorshift64(seed),
	}
}

// UAType represents a category of User-Agent
type UAType int

const (
	TypeDesktop UAType = iota
	TypeMobile
	TypeBot
)

// Random returns a random User-Agent from any category
func (g *Generator) Random() string {
	switch g.rng.intn(3) {
	case 0:
		return g.RandomDesktop()
	case 1:
		return g.RandomMobile()
	default:
		return g.RandomBot()
	}
}

// RandomDesktop returns a random desktop browser User-Agent
func (g *Generator) RandomDesktop() string {
	switch g.rng.intn(4) {
	case 0:
		return g.Chrome()
	case 1:
		return g.Firefox()
	case 2:
		return g.Safari()
	default:
		return g.Edge()
	}
}

// RandomMobile returns a random mobile browser User-Agent
func (g *Generator) RandomMobile() string {
	switch g.rng.intn(4) {
	case 0:
		return g.SafariIOS()
	case 1:
		return g.ChromeAndroid()
	case 2:
		return g.ChromeIOS()
	default:
		return g.AndroidWebView()
	}
}

// Package-level bot UA slice (zero allocation on access)
var botUAs = []string{
	googlebotUA,
	googlebotMobileUA,
	bingbotUA,
	bingbotMobileUA,
	yandexBotUA,
	baiduspiderUA,
	facebookBotUA,
	twitterBotUA,
	linkedinBotUA,
	ahrefsBotUA,
	semrushBotUA,
}

// RandomBot returns a random bot User-Agent
func (g *Generator) RandomBot() string {
	return botUAs[g.rng.intn(len(botUAs))]
}

// State returns the internal PRNG state for debugging/serialization
func (g *Generator) State() uint64 {
	return g.rng.state
}

// Clone creates a copy of the generator with the same state
// Useful for creating checkpoints
func (g *Generator) Clone() *Generator {
	return &Generator{
		rng: &xorshift64{state: g.rng.state},
	}
}

// --- Package-level functions using global generator (thread-safe) ---

// Chrome returns a random Chrome desktop User-Agent
func Chrome() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.Chrome()
}

// ChromeWindows returns a Chrome User-Agent for Windows
func ChromeWindows() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.ChromeWindows()
}

// ChromeMac returns a Chrome User-Agent for macOS
func ChromeMac() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.ChromeMac()
}

// ChromeLinux returns a Chrome User-Agent for Linux
func ChromeLinux() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.ChromeLinux()
}

// Firefox returns a random Firefox desktop User-Agent
func Firefox() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.Firefox()
}

// FirefoxWindows returns a Firefox User-Agent for Windows
func FirefoxWindows() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.FirefoxWindows()
}

// FirefoxMac returns a Firefox User-Agent for macOS
func FirefoxMac() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.FirefoxMac()
}

// Safari returns a Safari desktop User-Agent (macOS)
func Safari() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.Safari()
}

// Edge returns a random Edge desktop User-Agent
func Edge() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.Edge()
}

// EdgeWindows returns an Edge User-Agent for Windows
func EdgeWindows() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.EdgeWindows()
}

// SafariIOS returns a Safari User-Agent for iPhone
func SafariIOS() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.SafariIOS()
}

// SafariIPad returns a Safari User-Agent for iPad
func SafariIPad() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.SafariIPad()
}

// ChromeIOS returns a Chrome User-Agent for iOS
func ChromeIOS() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.ChromeIOS()
}

// ChromeAndroid returns a Chrome User-Agent for Android
func ChromeAndroid() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.ChromeAndroid()
}

// AndroidWebView returns an Android WebView User-Agent
func AndroidWebView() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.AndroidWebView()
}

// FirefoxAndroid returns a Firefox User-Agent for Android
func FirefoxAndroid() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.FirefoxAndroid()
}

// SamsungBrowser returns a Samsung Internet Browser User-Agent
func SamsungBrowser() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.SamsungBrowser()
}

// EdgeAndroid returns an Edge User-Agent for Android
func EdgeAndroid() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.EdgeAndroid()
}

// Random returns a random User-Agent from any category
func Random() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.Random()
}

// RandomDesktop returns a random desktop browser User-Agent
func RandomDesktop() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.RandomDesktop()
}

// RandomMobile returns a random mobile browser User-Agent
func RandomMobile() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.RandomMobile()
}

// RandomBot returns a random bot User-Agent
func RandomBot() string {
	globalLock.Lock()
	defer globalLock.Unlock()
	return globalGen.RandomBot()
}

// useragentBufSize is the pre-allocated buffer size for UA string building.
// Sized for the longest possible UA (Edge Android ~270 bytes).
// Using one size avoids branching and simplifies code.
const useragentBufSize = 280
