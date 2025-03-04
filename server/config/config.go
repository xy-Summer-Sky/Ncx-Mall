package config

// WXPayConfig 微信支付配置
type WXPayConfig struct {
    AppID     string `mapstructure:"appid" json:"appid" yaml:"appid"`
    MchID     string `mapstructure:"mch_id" json:"mch_id" yaml:"mch_id"`
    NotifyURL string `mapstructure:"notify_url" json:"notify_url" yaml:"notify_url"`
    ApiV3Key    string `mapstructure:"api_key" json:"api_key" yaml:"api_key"`
}
type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	
	// oss
	Local        Local        `mapstructure:"local" json:"local" yaml:"local"`
	Minio        Minio        `mapstructure:"minio" json:"minio" yaml:"minio"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
	// auto 
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// 微信配置
	WXPay WXPayConfig `json:"wxpay" yaml:"wxpay"`
	//grpc
	GrpcServer GrpcServer `mapstructure:"grpc-server" json:"grpc-server" yaml:"grpc-server"`

}
