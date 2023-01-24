package config

type RssHubConfig struct {
	UrlPrefixs []string `mapstructure:"url_prefixes"`
	ServiceUrl string   `mapstructure:"serviceUrl"`
}

type RssConfig struct {
	Hub *RssHubConfig `mapstructure:"hub"`
}

var _conf *RssConfig

var DefaultConfig *RssConfig = &RssConfig{
	Hub: &RssHubConfig{UrlPrefixs: []string{"rsshub:", "https://rsshub.app"}},
}

func Config() *RssConfig {
	if _conf == nil {
		return SetConfig(DefaultConfig)
	}
	return _conf
}
func SetConfig(conf *RssConfig) *RssConfig {
	_conf = conf

	if conf.Hub == nil {
		conf.Hub = DefaultConfig.Hub
	} else {
		if len(conf.Hub.UrlPrefixs) == 0 {
			conf.Hub.UrlPrefixs = DefaultConfig.Hub.UrlPrefixs
		}
	}

	return conf
}
