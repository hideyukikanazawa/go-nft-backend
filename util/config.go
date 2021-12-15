package util

// Config stores all configurations of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	ALCHEMY_API_KEY string `mapstructure:"ALCHEMY_API_KEY"`
}

type LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")    // json, xml, yaml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return config
}
