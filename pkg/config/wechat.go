package config

type MiniProgram struct {
	AppId     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	Debug     bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	HttpDebug bool   `mapstructure:"http_debug" json:"http_debug" yaml:"http_debug"`
	File      string `mapstructure:"file" json:"file" yaml:"file"`
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	AesKey    string `mapstructure:"aes_key" json:"aes_key" yaml:"aes_key"`
}

type Official struct {
	AppId     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	Debug     bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
}

type Work struct {
	AppId     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
	AgentId   string `mapstructure:"agent_id" json:"agent_id" yaml:"agent_id"`
	CallBack  string `mapstructure:"callback" json:"callback" yaml:"callback"`
	Debug     bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
}

type Wechat struct {
	MiniProgram *MiniProgram
	Official    *Official
	Work        *Work
}
