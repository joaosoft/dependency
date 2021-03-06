package main

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Dependency *DependencyConfig `json:"dependency"`
}

// DependencyConfig ...
type DependencyConfig struct {
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
	Log      struct {
		Level string `json:"level"`
	} `json:"log"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
