package fpclient

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"

	"github.com/playwright-community/playwright-go"
)

// Load tls fingerprint and return *TlsFingerprint.
func LoadTlsFingerprint(config *LoadingConfig) (*TlsFingerprint, error) {
	var fpStr string
	
	if config.FilePath != "" {
		data, err := ioutil.ReadFile(config.FilePath)
		if err != nil {
			return nil, err
		}

		fpStr = string(data)
	}
	
	if config.String != "" {
		fpStr = config.String

		if config.Isbase64String {
			str, err := base64.StdEncoding.DecodeString(fpStr)
			if err != nil {
				return nil, err
			}

			fpStr = string(str)
		}
	}

	var fp TlsFingerprint
	if err := json.Unmarshal([]byte(fpStr), &fp); err != nil {
		return nil, err
	}

	return &fp, nil
}

// Dump TLS fingerprint from chromedriver
func DumpTlsFingerprint(config *DumpTlsConfig) (*TlsFingerprint, error) {
	err := playwright.Install()
	if err != nil {
		return nil, err
	}

	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(config.Headless),
	})
	if err != nil {
		return nil, err
	}

	page, err := browser.NewPage(playwright.BrowserNewContextOptions{
		IgnoreHttpsErrors: playwright.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	resp, err := page.Goto("https://tls.peet.ws/api/all")
	if err != nil {
		return nil, err
	}

	response, err := resp.Text()
	if err != nil {
		return nil, err
	}

	var fp TlsFingerprint
	if err := json.Unmarshal([]byte(response), &fp); err != nil {
		return nil, err
	}
	
	err = ioutil.WriteFile(config.OutputPath, []byte(response), 0644)
	if err != nil {
		return nil, err
	}

	return &fp, nil
}
