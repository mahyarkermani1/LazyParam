package installer

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"io"
	"lazyparam/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
)

const (
	FallParamsURL = "https://github.com/ImAyrix/fallparams/releases/download/v1.0.12/fallparams_1.0.12_linux_amd64.zip"
	X8URL         = "https://github.com/Sh1Yo/x8/releases/download/v4.3.0/x86_64-linux-x8.gz"
	ToolsDir      = "tools"
)

func EnsureTools() {
	if _, err := os.Stat(ToolsDir); os.IsNotExist(err) {
		os.Mkdir(ToolsDir, 0755)
	}
	checkAndInstall("fallparams", FallParamsURL)
	checkAndInstall("x8", X8URL)
}

func checkAndInstall(toolName, url string) {
	toolPath := filepath.Join(ToolsDir, toolName)
	if _, err := os.Stat(toolPath); os.IsNotExist(err) {
		logger.Info("Tool '%s' not found. Downloading...", toolName)
		
		resp, err := http.Get(url)
		if err != nil {
			logger.Error("Download failed: %v", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)

		if filepath.Ext(url) == ".zip" {
			extractZip(data, toolName)
		} else if filepath.Ext(url) == ".gz" {
			extractGzip(data, toolName)
		}
		
		os.Chmod(toolPath, 0755)
		logger.Success("'%s' installed!", toolName)
	}
}

func extractZip(data []byte, toolName string) {
	reader, _ := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	for _, f := range reader.File {
		if f.Name == toolName {
			rc, _ := f.Open()
			defer rc.Close()
			outFile, _ := os.Create(filepath.Join(ToolsDir, toolName))
			defer outFile.Close()
			io.Copy(outFile, rc)
			return
		}
	}
}

func extractGzip(data []byte, toolName string) {
	reader, _ := gzip.NewReader(bytes.NewReader(data))
	defer reader.Close()
	outFile, _ := os.Create(filepath.Join(ToolsDir, toolName))
	defer outFile.Close()
	io.Copy(outFile, reader)
}
