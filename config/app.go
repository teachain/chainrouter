package config

type AppConfig struct {
	Common *CommonConfig `yaml:"common"`
	Chains *ChainsConfig `yaml:"chains"`
}
type StubConfig struct {
	Common *CommonStubConfig `yaml:"common"`
}

type CommonStubConfig struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}
