package ua

import "testing"

func BenchmarkGooglebot(b *testing.B) {
	for b.Loop() {
		_ = Googlebot()
	}
}

func BenchmarkBingbot(b *testing.B) {
	for b.Loop() {
		_ = Bingbot()
	}
}

func BenchmarkChrome(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Chrome()
	}
}

func BenchmarkChromeWindows(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.ChromeWindows()
	}
}

func BenchmarkFirefox(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Firefox()
	}
}

func BenchmarkSafari(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Safari()
	}
}

func BenchmarkEdge(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Edge()
	}
}

func BenchmarkSafariIOS(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.SafariIOS()
	}
}

func BenchmarkChromeAndroid(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.ChromeAndroid()
	}
}

func BenchmarkAndroidWebView(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.AndroidWebView()
	}
}

func BenchmarkRandom(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Random()
	}
}

func BenchmarkRandomDesktop(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.RandomDesktop()
	}
}

func BenchmarkRandomMobile(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.RandomMobile()
	}
}

func BenchmarkXorshift64(b *testing.B) {
	rng := newXorshift64(42)
	b.ResetTimer()
	for b.Loop() {
		_ = rng.next()
	}
}

func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		_ = New()
	}
}

func BenchmarkWithSeed(b *testing.B) {
	for b.Loop() {
		_ = WithSeed(42)
	}
}

func BenchmarkClone(b *testing.B) {
	g := WithSeed(42)
	b.ResetTimer()
	for b.Loop() {
		_ = g.Clone()
	}
}

func BenchmarkChromeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		g := WithSeed(42)
		for pb.Next() {
			_ = g.Chrome()
		}
	})
}

func BenchmarkRandomParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		g := WithSeed(42)
		for pb.Next() {
			_ = g.Random()
		}
	})
}
