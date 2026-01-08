package main

import (
	"flag"
	"fmt"
	"lazyparam/pkg/config"
	"lazyparam/pkg/logger"
	"lazyparam/pkg/runner"
	"os"
)

func main() {
	// تعریف فلگ‌ها به صورت کوتاه و بلند
	var url string
	flag.StringVar(&url, "u", "", "Target URL")
	flag.StringVar(&url, "url", "", "Target URL")

	var silent bool
	flag.BoolVar(&silent, "s", false, "Silent mode")
	flag.BoolVar(&silent, "silent", false, "Silent mode")

	var noColor bool
	flag.BoolVar(&noColor, "nc", false, "Disable colors")
	flag.BoolVar(&noColor, "no-color", false, "Disable colors")

	// شخصی‌سازی راهنما (Help)
	flag.Usage = func() {
		logger.Banner()
		fmt.Printf("Usage: ./lazyparam [options]\n\n")
		fmt.Printf("Options:\n")
		fmt.Printf("  -u,  --url        Target URL (e.g., https://example.com)\n")
		fmt.Printf("  -s,  --silent     Silent mode (no banner/logs)\n")
		fmt.Printf("  -nc, --no-color   Disable colored output\n")
		fmt.Printf("  -h,  --help       Show this help menu\n")
	}

	flag.Parse()

	// اعمال تنظیمات لاگر
	logger.Silent = silent
	logger.NoColor = noColor

	// چاپ بنر فقط اگر حالت سایلنت غیرفعال باشد
	if !logger.Silent {
		logger.Banner()
	}

	// چک کردن اجباری بودن URL
	if url == "" {
		logger.Error("URL is required! Use -u or --url")
		fmt.Printf("\nExample: ./lazyparam -u https://github.com\n")
		os.Exit(1)
	}

	// لود کردن کانفیگ (که حالا شامل User-Agent هم هست)
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		logger.Error("Config Error: %v", err)
		os.Exit(1)
	}

	// شروع عملیات اصلی
	runner.Run(url, cfg)
}
