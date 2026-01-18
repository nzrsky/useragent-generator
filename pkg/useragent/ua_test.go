package ua

import (
	"regexp"
	"strings"
	"testing"
)

func TestSeedReproducibility(t *testing.T) {
	seed := uint64(12345)

	g1 := WithSeed(seed)
	g2 := WithSeed(seed)

	for i := 0; i < 100; i++ {
		ua1 := g1.Chrome()
		ua2 := g2.Chrome()

		if ua1 != ua2 {
			t.Errorf("iteration %d: expected same UA, got:\n%s\n%s", i, ua1, ua2)
		}
	}
}

func TestChromeFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \([^)]+\) AppleWebKit/537\.36 \(KHTML, like Gecko\) Chrome/\d+\.\d+\.\d+\.\d+ Safari/537\.36$`)

	for i := 0; i < 50; i++ {
		ua := g.Chrome()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Chrome UA format: %s", ua)
		}
	}
}

func TestFirefoxFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \([^)]+; rv:\d+\.\d+\) Gecko/20100101 Firefox/\d+\.\d+$`)

	for i := 0; i < 50; i++ {
		ua := g.Firefox()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Firefox UA format: %s", ua)
		}
	}
}

func TestSafariFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \(Macintosh; Intel Mac OS X [^)]+\) AppleWebKit/\d+\.\d+\.\d+ \(KHTML, like Gecko\) Version/\d+\.\d+(?:\.\d+)? Safari/\d+\.\d+\.\d+$`)

	for i := 0; i < 50; i++ {
		ua := g.Safari()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Safari UA format: %s", ua)
		}
	}
}

func TestEdgeFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \([^)]+\) AppleWebKit/537\.36 \(KHTML, like Gecko\) Chrome/\d+\.\d+\.\d+\.\d+ Safari/537\.36 Edg/\d+\.\d+\.\d+\.\d+$`)

	for i := 0; i < 50; i++ {
		ua := g.Edge()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Edge UA format: %s", ua)
		}
	}
}

func TestSafariIOSFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \(iPhone; CPU iPhone OS [^)]+\) AppleWebKit/\d+\.\d+\.\d+ \(KHTML, like Gecko\) Version/\d+\.\d+(?:\.\d+)? Mobile/\w+ Safari/\d+\.\d+\.\d+$`)

	for i := 0; i < 50; i++ {
		ua := g.SafariIOS()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Safari iOS UA format: %s", ua)
		}
	}
}

func TestChromeAndroidFormat(t *testing.T) {
	g := WithSeed(42)

	pattern := regexp.MustCompile(`^Mozilla/5\.0 \(Linux; Android \d+; [^)]+\) AppleWebKit/537\.36 \(KHTML, like Gecko\) Chrome/\d+\.\d+\.\d+\.\d+ Mobile Safari/537\.36$`)

	for i := 0; i < 50; i++ {
		ua := g.ChromeAndroid()
		if !pattern.MatchString(ua) {
			t.Errorf("invalid Chrome Android UA format: %s", ua)
		}
	}
}

func TestBotConstants(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
		want string
	}{
		{"Googlebot", Googlebot, "Googlebot"},
		{"GooglebotMobile", GooglebotMobile, "Googlebot"},
		{"Bingbot", Bingbot, "bingbot"},
		{"YandexBot", YandexBot, "YandexBot"},
		{"Baiduspider", Baiduspider, "Baiduspider"},
		{"FacebookBot", FacebookBot, "facebookexternalhit"},
		{"TwitterBot", TwitterBot, "Twitterbot"},
		{"LinkedInBot", LinkedInBot, "LinkedInBot"},
		{"AhrefsBot", AhrefsBot, "AhrefsBot"},
		{"SemrushBot", SemrushBot, "SemrushBot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := tt.fn()
			if !strings.Contains(ua, tt.want) {
				t.Errorf("%s() = %q, want contains %q", tt.name, ua, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	g1 := WithSeed(12345)

	for i := 0; i < 10; i++ {
		g1.Chrome()
	}

	g2 := g1.Clone()

	for i := 0; i < 50; i++ {
		ua1 := g1.Random()
		ua2 := g2.Random()

		if ua1 != ua2 {
			t.Errorf("cloned generator produced different UA: %s vs %s", ua1, ua2)
		}
	}
}

func TestRandomDistribution(t *testing.T) {
	g := WithSeed(42)

	counts := make(map[string]int)
	n := 1000

	for i := 0; i < n; i++ {
		ua := g.RandomDesktop()
		switch {
		case strings.Contains(ua, "Chrome/") && !strings.Contains(ua, "Edg/"):
			counts["Chrome"]++
		case strings.Contains(ua, "Firefox/"):
			counts["Firefox"]++
		case strings.Contains(ua, "Safari/") && !strings.Contains(ua, "Chrome/"):
			counts["Safari"]++
		case strings.Contains(ua, "Edg/"):
			counts["Edge"]++
		}
	}

	// Each browser should get roughly 25% (with some variance)
	for browser, count := range counts {
		ratio := float64(count) / float64(n)
		if ratio < 0.15 || ratio > 0.35 {
			t.Errorf("uneven distribution for %s: %d/%d = %.2f", browser, count, n, ratio)
		}
	}
}

func TestMobileGenerators(t *testing.T) {
	g := WithSeed(42)

	generators := []struct {
		name string
		fn   func() string
	}{
		{"SafariIOS", g.SafariIOS},
		{"SafariIPad", g.SafariIPad},
		{"ChromeIOS", g.ChromeIOS},
		{"ChromeAndroid", g.ChromeAndroid},
		{"AndroidWebView", g.AndroidWebView},
		{"FirefoxAndroid", g.FirefoxAndroid},
		{"SamsungBrowser", g.SamsungBrowser},
		{"EdgeAndroid", g.EdgeAndroid},
	}

	for _, gen := range generators {
		t.Run(gen.name, func(t *testing.T) {
			ua := gen.fn()
			if ua == "" {
				t.Errorf("%s() returned empty string", gen.name)
			}
			if len(ua) < 50 {
				t.Errorf("%s() returned suspiciously short UA: %s", gen.name, ua)
			}
		})
	}
}

func TestXorshift64(t *testing.T) {
	rng := newXorshift64(12345)

	seen := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		val := rng.next()
		if seen[val] {
			t.Errorf("duplicate value after %d iterations: %d", i, val)
		}
		seen[val] = true
	}
}

func TestPackageLevelFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
	}{
		{"Chrome", Chrome},
		{"ChromeWindows", ChromeWindows},
		{"Firefox", Firefox},
		{"Safari", Safari},
		{"Edge", Edge},
		{"SafariIOS", SafariIOS},
		{"ChromeAndroid", ChromeAndroid},
		{"Random", Random},
		{"RandomDesktop", RandomDesktop},
		{"RandomMobile", RandomMobile},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := tt.fn()
			if ua == "" {
				t.Errorf("%s() returned empty string", tt.name)
			}
			if len(ua) < 50 {
				t.Errorf("%s() returned suspiciously short UA: %s", tt.name, ua)
			}
		})
	}
}

func TestRandomBot(t *testing.T) {
	ua := RandomBot()
	if ua == "" {
		t.Error("RandomBot() returned empty string")
	}
}

func TestGlobalSeed(t *testing.T) {
	Seed(99999)
	ua1 := Chrome()
	ua2 := Chrome()

	Seed(99999)
	ua3 := Chrome()
	ua4 := Chrome()

	if ua1 != ua3 {
		t.Errorf("Seed() didn't reset: expected %q, got %q", ua1, ua3)
	}
	if ua2 != ua4 {
		t.Errorf("Seed() didn't reset sequence: expected %q, got %q", ua2, ua4)
	}
}

func TestAllDesktopBrowsers(t *testing.T) {
	g := WithSeed(42)
	browsers := []struct {
		name string
		fn   func() string
	}{
		{"ChromeMac", g.ChromeMac},
		{"ChromeLinux", g.ChromeLinux},
		{"FirefoxWindows", g.FirefoxWindows},
		{"FirefoxMac", g.FirefoxMac},
		{"EdgeWindows", g.EdgeWindows},
	}
	for _, b := range browsers {
		t.Run(b.name, func(t *testing.T) {
			ua := b.fn()
			if len(ua) < 50 {
				t.Errorf("%s() too short: %s", b.name, ua)
			}
		})
	}
}

func TestAllMobileBrowsersPackage(t *testing.T) {
	fns := []struct {
		name string
		fn   func() string
	}{
		{"SafariIPad", SafariIPad},
		{"ChromeIOS", ChromeIOS},
		{"AndroidWebView", AndroidWebView},
		{"FirefoxAndroid", FirefoxAndroid},
		{"SamsungBrowser", SamsungBrowser},
		{"EdgeAndroid", EdgeAndroid},
	}
	for _, f := range fns {
		t.Run(f.name, func(t *testing.T) {
			ua := f.fn()
			if len(ua) < 50 {
				t.Errorf("%s() too short: %s", f.name, ua)
			}
		})
	}
}

func TestAllBots(t *testing.T) {
	bots := []struct {
		name     string
		fn       func() string
		contains string
	}{
		{"BingbotMobile", BingbotMobile, "bingbot"},
		{"YandexBotMobile", YandexBotMobile, "YandexBot"},
		{"DuckDuckBot", DuckDuckBot, "DuckDuckBot"},
		{"SlackBot", SlackBot, "Slackbot"},
		{"TelegramBot", TelegramBot, "TelegramBot"},
		{"DiscordBot", DiscordBot, "Discordbot"},
		{"WhatsAppBot", WhatsAppBot, "WhatsApp"},
		{"PinterestBot", PinterestBot, "Pinterest"},
		{"MozBot", MozBot, "DotBot"},
		{"MajesticBot", MajesticBot, "MJ12bot"},
		{"ScreamingFrogBot", ScreamingFrogBot, "Screaming Frog"},
		{"SitebulbBot", SitebulbBot, "Sitebulb"},
	}
	for _, b := range bots {
		t.Run(b.name, func(t *testing.T) {
			ua := b.fn()
			if !strings.Contains(ua, b.contains) {
				t.Errorf("%s() = %q, want contains %q", b.name, ua, b.contains)
			}
		})
	}
}

func TestNewGenerator(t *testing.T) {
	g := New()
	ua := g.Chrome()
	if len(ua) < 50 {
		t.Errorf("New().Chrome() too short: %s", ua)
	}
}

func TestGeneratorState(t *testing.T) {
	g := WithSeed(12345)
	g.Chrome()
	state := g.State()
	if state == 12345 {
		t.Error("State should have changed after generating UA")
	}
}
