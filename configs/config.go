package configs

import "github.com/spf13/viper"

type conf struct {
	MaxRequestsIP    int    `mapstructure:"MAX_REQUESTS_IP"`
	MaxRequestsToken int    `mapstructure:"MAX_REQUESTS_TOKEN"`
	BlockTime        int    `mapstructure:"BLOCK_TIME"`
	RedisADDR        string `mapstructure:"REDIS_ADDR"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
