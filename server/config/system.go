package config

type System struct {
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	ServerPort    int    `mapstructure:"server-port" json:"server-port" json:"server-port" yaml:"server-port"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`
}
