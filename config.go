package nataelb

type Config struct {
	Port                int       `yaml:"Port"`
	URLInfos            []URLInfo `yaml:"URLInfos"`
	HealthcheckInterval int       `yaml:"HealthcheckInterval"`
}

type URLInfo struct {
	URL             string `yaml:"URL"`
	HealthcheckPath string `yaml:"HealthcheckPath"`
}
