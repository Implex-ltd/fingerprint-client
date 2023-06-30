package fpclient

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

// Load fingerprint and return *Fingerprint.
func LoadFingerprint(config *LoadingConfig) (*Fingerprint, error) {
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

	var fp Fingerprint
	if err := json.Unmarshal([]byte(fpStr), &fp); err != nil {
		return nil, err
	}

	return &fp, nil
}
