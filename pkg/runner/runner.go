package runner

import (
	"fmt"
	"lazyparam/pkg/config"
	"lazyparam/pkg/installer"
	"lazyparam/pkg/logger"
	"lazyparam/pkg/utils"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Run(targetURL string, cfg *config.Config) {
	installer.EnsureTools()


	wordlistsDir := "wordlists"
	if _, err := os.Stat(wordlistsDir); os.IsNotExist(err) {
		os.Mkdir(wordlistsDir, 0755)
	}

	cwd, _ := os.Getwd()
	fallPath := filepath.Join(cwd, "tools", "fallparams")
	x8Path := filepath.Join(cwd, "tools", "x8")


	if _, err := os.Stat(cfg.Wordlists.CustomWordlistPath); os.IsNotExist(err) {
		os.WriteFile(cfg.Wordlists.CustomWordlistPath, []byte(""), 0644)
	}

	logger.Info("Checking target: %s", targetURL)
	if !utils.CheckURL(targetURL, cfg.HTTPSettings.UserAgent) {
		logger.Error("Target is down or blocked!")
		os.Exit(1)
	}


	parsedURL, _ := url.Parse(targetURL)
	domainName := strings.ReplaceAll(parsedURL.Host, ".", "_")
	extractedFile := filepath.Join(wordlistsDir, fmt.Sprintf("%s_extracted.txt", domainName))


	logger.Info("Running FallParams...")
	_ = exec.Command(fallPath, "-u", targetURL, "-o", extractedFile).Run()


	foundParams, _ := utils.ReadLines(extractedFile)
	customParams, _ := utils.ReadLines(cfg.Wordlists.CustomWordlistPath)
	
	countFound := len(foundParams)
	countCustom := len(customParams)
	
	finalParams := utils.Deduplicate(append(foundParams, customParams...))
	countFinal := len(finalParams)


	logger.Success("Stats: Extracted: %d | Custom: %d | Combined (Unique): %d", 
		countFound, countCustom, countFinal)

	if countFinal == 0 {
		logger.Error("No parameters to fuzz!")
		os.Exit(1)
	}

	finalWordlistPath := "final_fuzz_temp.txt"
	utils.WriteLines(finalWordlistPath, finalParams)


	for _, method := range cfg.X8Settings.Methods {
		logger.Info("Starting x8 | Method: %s", method)

		args := []string{
			"-u", targetURL,
			"-w", finalWordlistPath,
			"-X", method,
			"-m", strconv.Itoa(cfg.X8Settings.ChainCount),
			"--disable-progress-bar",
			"-H", "User-Agent: " + cfg.HTTPSettings.UserAgent,
		}

		if cfg.X8Settings.Delay > 0 {
			args = append(args, "--delay", strconv.Itoa(cfg.X8Settings.Delay))
		}

		if cfg.X8Settings.Headers != "" {
			args = append(args, "-H", cfg.X8Settings.Headers)
		}

		c := exec.Command(x8Path, args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		_ = c.Run()
	}


	os.Remove(finalWordlistPath)
	logger.Success("Done. Extracted params saved in: %s", extractedFile)
}
