package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Settings - Actual service settings
var Settings Config

type Config struct {
	AppName        string
	DefaultAppPort string
	DB             DataBase
	CurrencyPair   CurrencyConversion
}
type DataBase struct {
	Host     string
	Username string
	Password string
	Database string
	Debug    bool
}
type CurrencyConversion struct {
	Fsyms string
	Tsyms string
}

func Init() (Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	logrus.Println("DB NAME", viper.GetString("app.name"))
	viper.SetEnvPrefix("APP")
	viper.SetDefault("NAME", viper.GetString("app.name"))
	viper.SetDefault("PORT", viper.GetString("app.port"))

	viper.SetDefault("POSTGRES_HOSTNAME", viper.GetString("postgres.db_host"))
	viper.SetDefault("POSTGRES_USERNAME", viper.GetString("postgres.db_user"))
	viper.SetDefault("POSTGRES_PASSWORD", viper.GetString("postgres.db_password"))
	viper.SetDefault("POSTGRES_DATABASE", viper.GetString("postgres.db_name"))

	viper.SetDefault("POSTGRES_DEBUG", viper.GetString("postgres.db_debug"))

	viper.SetDefault("CURRENCY_FSYMS", viper.GetString("currency_pairs.fsyms"))
	viper.SetDefault("CURRENCY_TSYMS", viper.GetString("currency_pairs.tsyms"))

	viper.AutomaticEnv()
	//viper.Debug()
	Settings = Config{
		AppName:        viper.GetString("NAME"),
		DefaultAppPort: viper.GetString("PORT"),
		DB: DataBase{
			Host:     viper.GetString("POSTGRES_HOSTNAME"),
			Username: viper.GetString("POSTGRES_USERNAME"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Database: viper.GetString("POSTGRES_DATABASE"),
			Debug:    viper.GetBool("POSTGRES_DEBUG"),
		},
		CurrencyPair: CurrencyConversion{
			Fsyms: viper.GetString("CURRENCY_FSYMS"),
			Tsyms: viper.GetString("CURRENCY_TSYMS"),
		},
	}
	return Settings, nil
}
