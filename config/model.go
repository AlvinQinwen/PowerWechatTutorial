package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type Configuration struct {
	Payment      Payment
	MiniProgram  MiniProgram
	WeCom        WeCom
	OffiAccount  OffiAccount
	OpenPlatform OpenPlatform
	OpenAI       OpenAI
	Discord      Discord
}

type Payment struct {
	AppID              string `required:"true" env:"app_id"`
	MchID              string `required:"true" env:"mch_id"`
	CertPath           string `required:"true" env:"wx_cert_path"`
	KeyPath            string `required:"true" env:"wx_key_path"`
	SerialNo           string `required:"true" env:"serial_no"`
	CertificateKeyPath string `required:"false" env:"certificate_key_path"`
	WechatPaySerial    string `required:"false" env:"wechat_pay_serial"`
	RSAPublicKeyPath   string `required:"false" env:"wx_rsa_public_key_path"`
	MchApiV3Key        string `env:"mch_api_v3_key"`
	Key                string `env:"key"`
	ResponseType       string `default:"map"`
	NotifyURL          string `env:"notify_url"`
	SubMchID           string `env:"sub_mch_id"`
	SubAppID           string `env:"sub_app_id"`
	HttpDebug          bool   `default:"false"`
	RedisAddr          string `env:"redis_addr"`
}

type MiniProgram struct {
	AppID         string `required:"true" env:"miniprogram_app_id"`
	Secret        string `required:"true" env:"miniprogram_secret"`
	RedisAddr     string `env:"redis_addr"`
	MessageToken  string `env:"message_token"`
	MessageAesKey string `env:"message_aes_key"`
}

type WeCom struct {
	CorpID          string `env:"corp_id"`
	AgentID         int    `env:"wecom_agent_id"`
	Secret          string `env:"wecom_secret"`
	MessageToken    string `env:"app_message_token"`
	MessageAesKey   string `env:"app_message_aes_key"`
	MessageCallback string `env:"app_message_callback_url"`
	OAuthCallback   string `env:"app_oauth_callback_url"`
	ContactSecret   string `env:"contact_secret"`
	ContactToken    string `env:"contact_token"`
	ContactAESKey   string `env:"contact_aes_key"`
	ContactCallback string `env:"contact_callback_url"`
	RedisAddr       string `env:"redis_addr"`
}

type OffiAccount struct {
	AppID         string `required:"true" env:"appid"`
	AppSecret     string `required:"true" env:"appsecret"`
	RedisAddr     string `env:"redis_addr"`
	MessageToken  string `env:"message_token"`
	MessageAesKey string `env:"message_aes_key"`
}

type OpenPlatform struct {
	AppID         string
	AppSecret     string
	MessageToken  string
	MessageAesKey string
	RedisAddr     string
}

type OpenAI struct {
	OpenAPIKey   string
	Organization string
	HttpDebug    bool
	ProxyURL     string
}

type Discord struct {
	Token string
}

func configFiles() []string {
	return []string{"config.yml", "/etc/gotify/config.yml"}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{}).Load(conf, configFiles()...)
	if err != nil {
		log.Printf("%#v", conf)
		panic(err)
	}
	return conf
}
