package config

type OfficialAccount struct {
	AppId         string `mapstructure:"appId" json:"appId" yaml:"appId"`
	Token         string `mapstructure:"token" json:"token" yaml:"token"`
	AppSecret     string `mapstructure:"appSecret" json:"appSecret" yaml:"appSecret"`
	MessageToken  string `mapstructure:"messageToken" json:"messageToken" yaml:"messageToken"`
	MessageAesKey string `mapstructure:"messageAesKey" json:"messageAesKey" yaml:"messageAesKey"`
}
