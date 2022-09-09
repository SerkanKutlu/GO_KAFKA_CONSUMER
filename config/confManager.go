package config

import (
	"github.com/spf13/viper"
)

type configurationManager struct {
	applicationConfig *AppConfig
}

func NewConfigurationManager(path string, file string, env string) *configurationManager {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	//viper.KeyDelimiter(",")
	appConfig := readApplicationConfigFile(env, file)
	return &configurationManager{
		applicationConfig: appConfig,
	}
}

func (cm *configurationManager) GetKafkaConfiguration() *KafkaConfig {
	return &cm.applicationConfig.KafkaConfig
}
func (cm *configurationManager) GetMongoConfiguration() *MongoConfig {
	return &cm.applicationConfig.MongoConfig
}

func readApplicationConfigFile(env string, file string) *AppConfig {

	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	var appConfig AppConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	return &appConfig
}
