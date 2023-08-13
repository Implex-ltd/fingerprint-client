package main

import (
	"fmt"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
)

func main() {
	// load fp from json file
	browser_fp, err := fpclient.LoadFingerprint(&fpclient.LoadingConfig{
		FilePath: "./fp.json",
	})

	if err != nil {
		panic(err)
	}

	// load fp from string
	_, _ = fpclient.LoadFingerprint(&fpclient.LoadingConfig{
		Isbase64String: true, // you can ignore if the fingerprint is not in base64 format.
		String:         "ZXhlbXBsZSBiYXNlNjQ=",
	})

	tls_fp, err := fpclient.DumpTlsFingerprint(&fpclient.DumpTlsConfig{
		Headless:   false,
		OutputPath: "./tls_fp.json",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(tls_fp.TLS.Ja3)
	fmt.Println(browser_fp.GPU.Renderer)
	fmt.Println(browser_fp.Navigator.UserAgent)
	fmt.Println(browser_fp.Navigator.DeviceMemory)
}
