package main

import (
	"flag"

	"github.com/jeevangb/project-portal-management/internal/configs"
	"github.com/jeevangb/project-portal-management/internal/server"
)

func main() {
	env := flag.String("env", "", "")
	flag.Parse()
	// If 'env' is not provided, set it to the default value
	if *env == "" {
		*env = "dev" // Set your default environment value here
	}
	cfg, err := configs.LoadConfig(env)
	if err != nil {
		return
	}
	err = server.RegisterGrpcServer(cfg.Port)
	if err != nil {
		return
	}
}
