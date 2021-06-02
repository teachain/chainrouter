package config

//基础配置
type CommonConfig struct {
	Zone    string `yaml:"manager"`
	Visible bool `yaml:"visible"`
}

//链插件所在目录
type ChainsConfig struct {
	Path string `yaml:"path"`//到这里目录下扫描有哪些插件 一个目录表示一个插件（也表示一条链）
}
