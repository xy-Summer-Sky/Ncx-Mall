package config

type GrpcServer struct {
	Port 		int			`mapstructure:"port" json:"port" yaml:"port"`
	Type 		string 		`mapstructure:"type" json:"type" yaml:"type"`
}