package core

import "errors"

type CloudProviders int

const (
	GCP CloudProviders = 1
)

func (provider CloudProviders) Name() (string, error) {
	if provider == GCP {
		return "GCP", nil
	}

	return "", errors.New("unknown provider")
}
