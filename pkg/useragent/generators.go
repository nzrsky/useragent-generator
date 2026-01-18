package ua

import "strings"

func (g *Generator) SafariIOS() string {
	iosVer := pick(g.rng, iosVersions)
	safariVer := pick(g.rng, safariVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (iPhone; CPU iPhone OS ")
	b.WriteString(iosVer)
	b.WriteString(" like Mac OS X) AppleWebKit/")
	b.WriteString(webkitVersion)
	b.WriteString(" (KHTML, like Gecko) Version/")
	b.WriteString(safariVer)
	b.WriteString(" Mobile/15E148 Safari/")
	b.WriteString(webkitVersion)

	return b.String()
}

func (g *Generator) SafariIPad() string {
	iosVer := pick(g.rng, iosVersions)
	safariVer := pick(g.rng, safariVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (iPad; CPU OS ")
	b.WriteString(iosVer)
	b.WriteString(" like Mac OS X) AppleWebKit/")
	b.WriteString(webkitVersion)
	b.WriteString(" (KHTML, like Gecko) Version/")
	b.WriteString(safariVer)
	b.WriteString(" Mobile/15E148 Safari/")
	b.WriteString(webkitVersion)

	return b.String()
}

func (g *Generator) ChromeIOS() string {
	iosVer := pick(g.rng, iosVersions)
	chromeVer := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (iPhone; CPU iPhone OS ")
	b.WriteString(iosVer)
	b.WriteString(" like Mac OS X) AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) CriOS/")
	b.WriteString(chromeVer)
	b.WriteString(" Mobile/15E148 Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

func (g *Generator) ChromeAndroid() string {
	androidVer := pick(g.rng, androidVersions)
	chromeVer := pick(g.rng, chromeVersions)
	device := pick(g.rng, androidDevices)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Linux; Android ")
	b.WriteString(androidVer)
	b.WriteString("; ")
	b.WriteString(device.model)
	b.WriteString(" Build/")
	b.WriteString(device.build)
	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Mobile Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

func (g *Generator) AndroidWebView() string {
	androidVer := pick(g.rng, androidVersions)
	chromeVer := pick(g.rng, chromeVersions)
	device := pick(g.rng, androidDevices)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Linux; Android ")
	b.WriteString(androidVer)
	b.WriteString("; ")
	b.WriteString(device.model)
	b.WriteString(" Build/")
	b.WriteString(device.build)
	b.WriteString("; wv) AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Version/4.0 Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Mobile Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

func (g *Generator) FirefoxAndroid() string {
	androidVer := pick(g.rng, androidVersions)
	ffVer := pick(g.rng, firefoxVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Android ")
	b.WriteString(androidVer)
	b.WriteString("; Mobile; rv:")
	b.WriteString(ffVer)
	b.WriteString(") Gecko/")
	b.WriteString(ffVer)
	b.WriteString(" Firefox/")
	b.WriteString(ffVer)

	return b.String()
}

func (g *Generator) SamsungBrowser() string {
	androidVer := pick(g.rng, androidVersions)
	chromeVer := pick(g.rng, chromeVersions)
	device := pick(g.rng, androidDevices)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Linux; Android ")
	b.WriteString(androidVer)
	b.WriteString("; ")
	b.WriteString(device.model)
	b.WriteString(" Build/")
	b.WriteString(device.build)
	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) SamsungBrowser/25.0 Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Mobile Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

func (g *Generator) EdgeAndroid() string {
	androidVer := pick(g.rng, androidVersions)
	chromeVer := pick(g.rng, chromeVersions)
	edgeVer := pick(g.rng, edgeVersions)
	device := pick(g.rng, androidDevices)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Linux; Android ")
	b.WriteString(androidVer)
	b.WriteString("; ")
	b.WriteString(device.model)
	b.WriteString(" Build/")
	b.WriteString(device.build)
	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Mobile Safari/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" EdgA/")
	b.WriteString(edgeVer)

	return b.String()
}

// Browsers

// Chrome generates a Chrome User-Agent for random desktop OS
func (g *Generator) Chrome() string {
	version := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (")

	switch g.rng.intn(3) {
	case 0: // Windows
		b.WriteString("Windows NT ")
		b.WriteString(pick(g.rng, windowsVersions))
		b.WriteString("; Win64; x64")
	case 1: // macOS
		b.WriteString("Macintosh; Intel Mac OS X ")
		b.WriteString(pick(g.rng, macVersions))
	case 2: // Linux
		b.WriteString(pick(g.rng, linuxDesktops))
	}

	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(version)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

// ChromeWindows generates a Chrome User-Agent for Windows
func (g *Generator) ChromeWindows() string {
	version := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Windows NT ")
	b.WriteString(pick(g.rng, windowsVersions))
	b.WriteString("; Win64; x64) AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(version)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

// ChromeMac generates a Chrome User-Agent for macOS
func (g *Generator) ChromeMac() string {
	version := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Macintosh; Intel Mac OS X ")
	b.WriteString(pick(g.rng, macVersions))
	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(version)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

// ChromeLinux generates a Chrome User-Agent for Linux
func (g *Generator) ChromeLinux() string {
	version := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (")
	b.WriteString(pick(g.rng, linuxDesktops))
	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(version)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)

	return b.String()
}

// Firefox generates a Firefox desktop User-Agent
func (g *Generator) Firefox() string {
	version := pick(g.rng, firefoxVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (")

	switch g.rng.intn(3) {
	case 0: // Windows
		b.WriteString("Windows NT ")
		b.WriteString(pick(g.rng, windowsVersions))
		b.WriteString("; Win64; x64; rv:")
		b.WriteString(version)
		b.WriteString(") Gecko/20100101 Firefox/")
	case 1: // macOS
		b.WriteString("Macintosh; Intel Mac OS X ")
		b.WriteString(pick(g.rng, macVersions))
		b.WriteString("; rv:")
		b.WriteString(version)
		b.WriteString(") Gecko/20100101 Firefox/")
	case 2: // Linux
		b.WriteString(pick(g.rng, linuxDesktops))
		b.WriteString("; rv:")
		b.WriteString(version)
		b.WriteString(") Gecko/20100101 Firefox/")
	}

	b.WriteString(version)

	return b.String()
}

// FirefoxWindows generates a Firefox User-Agent for Windows
func (g *Generator) FirefoxWindows() string {
	version := pick(g.rng, firefoxVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Windows NT ")
	b.WriteString(pick(g.rng, windowsVersions))
	b.WriteString("; Win64; x64; rv:")
	b.WriteString(version)
	b.WriteString(") Gecko/20100101 Firefox/")
	b.WriteString(version)

	return b.String()
}

// FirefoxMac generates a Firefox User-Agent for macOS
func (g *Generator) FirefoxMac() string {
	version := pick(g.rng, firefoxVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Macintosh; Intel Mac OS X ")
	b.WriteString(pick(g.rng, macVersions))
	b.WriteString("; rv:")
	b.WriteString(version)
	b.WriteString(") Gecko/20100101 Firefox/")
	b.WriteString(version)

	return b.String()
}

// Safari generates a Safari desktop User-Agent (macOS only)
func (g *Generator) Safari() string {
	version := pick(g.rng, safariVersions)
	macVer := pick(g.rng, macVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Macintosh; Intel Mac OS X ")
	b.WriteString(macVer)
	b.WriteString(") AppleWebKit/")
	b.WriteString(webkitVersion)
	b.WriteString(" (KHTML, like Gecko) Version/")
	b.WriteString(version)
	b.WriteString(" Safari/")
	b.WriteString(webkitVersion)

	return b.String()
}

// Edge generates an Edge desktop User-Agent
func (g *Generator) Edge() string {
	edgeVer := pick(g.rng, edgeVersions)
	chromeVer := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (")

	switch g.rng.intn(2) {
	case 0: // Windows
		b.WriteString("Windows NT ")
		b.WriteString(pick(g.rng, windowsVersions))
		b.WriteString("; Win64; x64")
	case 1: // macOS
		b.WriteString("Macintosh; Intel Mac OS X ")
		b.WriteString(pick(g.rng, macVersions))
	}

	b.WriteString(") AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" Edg/")
	b.WriteString(edgeVer)

	return b.String()
}

// EdgeWindows generates an Edge User-Agent for Windows
func (g *Generator) EdgeWindows() string {
	edgeVer := pick(g.rng, edgeVersions)
	chromeVer := pick(g.rng, chromeVersions)

	var b strings.Builder
	b.Grow(useragentBufSize)

	b.WriteString("Mozilla/5.0 (Windows NT ")
	b.WriteString(pick(g.rng, windowsVersions))
	b.WriteString("; Win64; x64) AppleWebKit/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" (KHTML, like Gecko) Chrome/")
	b.WriteString(chromeVer)
	b.WriteString(" Safari/")
	b.WriteString(appleWebKitChrome)
	b.WriteString(" Edg/")
	b.WriteString(edgeVer)

	return b.String()
}

// Search engine bots
const (
	googlebotUA = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

	googlebotMobileUA = "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.6778.85 Mobile Safari/537.36 " +
		"(compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

	bingbotUA = "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"

	bingbotMobileUA = "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.6778.85 Mobile Safari/537.36 " +
		"(compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"

	yandexBotUA = "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)"

	yandexBotMobileUA = "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.6778.85 Mobile Safari/537.36 " +
		"(compatible; YandexBot/3.0; +http://yandex.com/bots)"

	baiduspiderUA = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"

	duckduckbotUA = "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)"
)

// Social media bots
const (
	facebookBotUA = "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)"
	twitterBotUA = "Twitterbot/1.0"
	linkedinBotUA = "LinkedInBot/1.0 (compatible; Mozilla/5.0; Apache-HttpClient +http://www.linkedin.com)"
	slackBotUA = "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)"
	telegramBotUA = "TelegramBot (like TwitterBot)"
	discordBotUA = "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)"
	whatsappBotUA = "WhatsApp/2.23.20.0"
	pinterestBotUA = "Pinterest/0.2 (+http://www.pinterest.com/bot.html)"
)

// SEO tool bots
const (
	ahrefsBotUA = "Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)"
	semrushBotUA = "Mozilla/5.0 (compatible; SemrushBot/7~bl; +http://www.semrush.com/bot.html)"
	mozBotUA = "Mozilla/5.0 (compatible; DotBot/1.2; +https://opensiteexplorer.org/dotbot; help@moz.com)"
	majesticBotUA = "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)"
	screaminFrogBotUA = "Screaming Frog SEO Spider/19.0"
	sitebulbBotUA = "Mozilla/5.0 (compatible; SitebulbBot/0.11.10; +https://sitebulb.com/crawler/)"
)

// Googlebot returns Google's web crawler User-Agent
func Googlebot() string { return googlebotUA }

// GooglebotMobile returns Google's mobile web crawler User-Agent
func GooglebotMobile() string { return googlebotMobileUA }

// Bingbot returns Bing's web crawler User-Agent
func Bingbot() string { return bingbotUA }

// BingbotMobile returns Bing's mobile web crawler User-Agent
func BingbotMobile() string { return bingbotMobileUA }

// YandexBot returns Yandex's web crawler User-Agent
func YandexBot() string { return yandexBotUA }

// YandexBotMobile returns Yandex's mobile web crawler User-Agent
func YandexBotMobile() string { return yandexBotMobileUA }

// Baiduspider returns Baidu's web crawler User-Agent
func Baiduspider() string { return baiduspiderUA }

// DuckDuckBot returns DuckDuckGo's web crawler User-Agent
func DuckDuckBot() string { return duckduckbotUA }

// FacebookBot returns Facebook's web crawler User-Agent
func FacebookBot() string { return facebookBotUA }

// TwitterBot returns Twitter's web crawler User-Agent
func TwitterBot() string { return twitterBotUA }

// LinkedInBot returns LinkedIn's web crawler User-Agent
func LinkedInBot() string { return linkedinBotUA }

// SlackBot returns Slack's link preview bot User-Agent
func SlackBot() string { return slackBotUA }

// TelegramBot returns Telegram's link preview bot User-Agent
func TelegramBot() string { return telegramBotUA }

// DiscordBot returns Discord's link preview bot User-Agent
func DiscordBot() string { return discordBotUA }

// WhatsAppBot returns WhatsApp's link preview bot User-Agent
func WhatsAppBot() string { return whatsappBotUA }

// PinterestBot returns Pinterest's web crawler User-Agent
func PinterestBot() string { return pinterestBotUA }

// AhrefsBot returns Ahrefs SEO crawler User-Agent
func AhrefsBot() string { return ahrefsBotUA }

// SemrushBot returns Semrush SEO crawler User-Agent
func SemrushBot() string { return semrushBotUA }

// MozBot returns Moz SEO crawler User-Agent
func MozBot() string { return mozBotUA }

// MajesticBot returns Majestic SEO crawler User-Agent
func MajesticBot() string { return majesticBotUA }

// ScreamingFrogBot returns Screaming Frog SEO Spider User-Agent
func ScreamingFrogBot() string { return screaminFrogBotUA }

// SitebulbBot returns Sitebulb crawler User-Agent
func SitebulbBot() string { return sitebulbBotUA }