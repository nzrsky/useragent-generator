# useragent-generator

[![CI](https://github.com/nzrsky/useragent-generator/actions/workflows/ci.yml/badge.svg)](https://github.com/nzrsky/useragent-generator/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/nzrsky/useragent-generator/graph/badge.svg)](https://codecov.io/gh/nzrsky/useragent-generator)
[![Go Reference](https://pkg.go.dev/badge/github.com/nzrsky/useragent-generator/pkg/useragent.svg)](https://pkg.go.dev/github.com/nzrsky/useragent-generator/pkg/useragent)
[![Go Report Card](https://goreportcard.com/badge/github.com/nzrsky/useragent-generator)](https://goreportcard.com/report/github.com/nzrsky/useragent-generator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

High-performance User-Agent string generator for Go.

## Features

- **Auto-updated browser versions** from [Intoli](https://github.com/intoli/user-agents) real usage data
- **Zero-alloc bot User-Agents** (~2ns per call)
- **Fast browser UA generation** (~40ns, 1 alloc)
- **Seed-based reproducibility** for testing
- **xorshift64 PRNG** - faster than math/rand
- Desktop: Chrome, Firefox, Safari, Edge
- Mobile: iOS Safari, Android Chrome, WebView
- Bots: Google, Bing, Yandex, Baidu, Facebook, Twitter, LinkedIn, SEO tools

## Installation

```bash
go get github.com/nzrsky/useragent-generator/pkg/useragent
```

## Quick Start

```go
package main

import (
    "fmt"
    ua "github.com/nzrsky/useragent-generator/pkg/useragent"
)

func main() {
    fmt.Println(ua.Chrome())
    fmt.Println(ua.SafariIOS())
    fmt.Println(ua.Googlebot())

    fmt.Println(ua.Random())        // any type
    fmt.Println(ua.RandomDesktop()) // desktop only
    fmt.Println(ua.RandomMobile())  // mobile only
}
```

## Reproducible Results

```go
// Set global seed for reproducibility
ua.Seed(12345)
fmt.Println(ua.Chrome()) // always same result with same seed

// Or create your own generator
g := ua.WithSeed(12345)
fmt.Println(g.Chrome())
fmt.Println(g.Firefox())
```

## Available Functions

### Desktop Browsers

| Function | Description |
|----------|-------------|
| `Chrome()` | Chrome (random OS) |
| `ChromeWindows()` | Chrome on Windows |
| `ChromeMac()` | Chrome on macOS |
| `ChromeLinux()` | Chrome on Linux |
| `Firefox()` | Firefox (random OS) |
| `FirefoxWindows()` | Firefox on Windows |
| `FirefoxMac()` | Firefox on macOS |
| `Safari()` | Safari on macOS |
| `Edge()` | Edge (random OS) |
| `EdgeWindows()` | Edge on Windows |

### Mobile Browsers

| Function | Description |
|----------|-------------|
| `SafariIOS()` | Safari on iPhone |
| `SafariIPad()` | Safari on iPad |
| `ChromeIOS()` | Chrome on iOS |
| `ChromeAndroid()` | Chrome on Android |
| `AndroidWebView()` | Android WebView |
| `FirefoxAndroid()` | Firefox on Android |
| `SamsungBrowser()` | Samsung Internet |
| `EdgeAndroid()` | Edge on Android |

### Bots (Zero-Allocation)

| Function | Description |
|----------|-------------|
| `Googlebot()` | Google web crawler |
| `GooglebotMobile()` | Google mobile crawler |
| `Bingbot()` | Bing web crawler |
| `BingbotMobile()` | Bing mobile crawler |
| `YandexBot()` | Yandex crawler |
| `YandexBotMobile()` | Yandex mobile crawler |
| `Baiduspider()` | Baidu crawler |
| `DuckDuckBot()` | DuckDuckGo crawler |
| `FacebookBot()` | Facebook crawler |
| `TwitterBot()` | Twitter crawler |
| `LinkedInBot()` | LinkedIn crawler |
| `SlackBot()` | Slack link preview |
| `TelegramBot()` | Telegram link preview |
| `DiscordBot()` | Discord link preview |
| `WhatsAppBot()` | WhatsApp link preview |
| `PinterestBot()` | Pinterest crawler |
| `AhrefsBot()` | Ahrefs SEO crawler |
| `SemrushBot()` | Semrush SEO crawler |
| `MozBot()` | Moz SEO crawler |
| `MajesticBot()` | Majestic SEO crawler |
| `ScreamingFrogBot()` | Screaming Frog |
| `SitebulbBot()` | Sitebulb crawler |

### Utilities

| Function | Description |
|----------|-------------|
| `Random()` | Random UA (any type) |
| `RandomDesktop()` | Random desktop browser |
| `RandomMobile()` | Random mobile browser |
| `RandomBot()` | Random bot |
| `Seed(uint64)` | Set global generator seed

## Performance

Benchmarks on Apple M3 Max:

```
BenchmarkGooglebot          644734208    1.780 ns/op    0 B/op    0 allocs/op
BenchmarkChrome              30398343   38.90 ns/op  192 B/op    1 allocs/op
BenchmarkFirefox             36187087   32.63 ns/op  128 B/op    1 allocs/op
BenchmarkSafariIOS           32038802   34.46 ns/op  192 B/op    1 allocs/op
BenchmarkChromeAndroid       25472996   46.88 ns/op  224 B/op    1 allocs/op
BenchmarkRandom              29923258   39.43 ns/op  129 B/op    0 allocs/op
```

## Example Output

```
Chrome (Windows):
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.6778.85 Safari/537.36

Safari iOS:
Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4 Mobile/15E148 Safari/605.1.15

Googlebot:
Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)
```

## License

MIT License - see [LICENSE](LICENSE) file.
