package config

type Config interface {
	GetCredentialFilePath() string
}

type protoConfig struct {
	credentialsFile string
}

func (c protoConfig) GetCredentialFilePath() string {
	return c.credentialsFile
}

func NewCliConfig(credentialsFile string) Config {
	t := struct{ protoConfig }{}
	t.credentialsFile = credentialsFile
	return t
}
