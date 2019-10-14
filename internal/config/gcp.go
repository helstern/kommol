package config

type GcpConfig interface {
	GetCredentialFilePath() string
}

type protoGcpConfig struct {
	credentialsFile string
}

func (c protoGcpConfig) GetCredentialFilePath() string {
	return c.credentialsFile
}

func NewGcpCliConfig(credentialsFile string) GcpConfig {
	t := struct{ protoGcpConfig }{}
	t.credentialsFile = credentialsFile
	return t
}
