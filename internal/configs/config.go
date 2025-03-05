package configs

import (
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string `mapstructure:"GRPC_SERVER_PORT"`
}

var envs = []string{
	"GRPC_SERVER_PORT",
}

func LoadConfig(env *string) (Config, error) {
	var cfg Config
	viper.AddConfigPath("./")
	viper.SetConfigFile("internal/configs/" + *env + "/application.env")
	viper.ReadInConfig()
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			log.Error().Err(err).Msg("could not bind envs")
			return cfg, err
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error().Err(err).Msg("could not unmarshal the configs")
		return cfg, err
	}
	if err := validator.New().Struct(&cfg); err != nil {
		log.Error().Err(err).Msg("could not validate the config struct")
	}
	return cfg, nil
}
