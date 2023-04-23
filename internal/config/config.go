package config

import "github.com/spf13/viper"

var (
	Port     string
	DbName   string
	Host     string
	User     string
	Password string

	Domain  string
	AppPort string

	Origins string

	AuthUrl string
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	//Db conf
	Port = viper.GetString("DBserver.port")
	DbName = viper.GetString("DBserver.name")
	Host = viper.GetString("DBserver.host")
	User = viper.GetString("database.user")
	Password = viper.GetString("database.password")

	//App conf
	Domain = viper.GetString("AppServer.domain")
	AppPort = viper.GetString("AppServer.port")

	//CORS conf
	Origins = viper.GetString("CORS.origins")

	//Auth url
	AuthUrl = viper.GetString("Auth.authUrl")

	return nil
}
