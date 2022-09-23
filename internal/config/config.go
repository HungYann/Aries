package config

import "server/internal/tool"

type Application struct {
	Address string `mapstructure:"address"`
}

type Log struct {
	OutputPath string `mapstructure:"output-path"`
	MaxTime    int    `mapstructure:"max-time"`
	RotateTime int    `mapstructure:"rotate-time"`
}

type Target struct {
	EndPoint string `mapstructure:"endpoint"`
}

type DataSource struct {
	MySQL string `mapstructure:"mysql"`
}

type global struct {
	Application `mapstructure:",squash"`
	Log         Log        `mapstructure:"log"`
	Target      Target     `mapstructure:"target"`
	DataSource  DataSource `mapstructure:"datasource"`
}

var (
	g = global{}
)

func GetApplication() Application {
	return g.Application
}

func GetLog() Log {
	return g.Log
}

func GetTarget() Target {
	return g.Target
}

func GetDataSource() DataSource {
	return g.DataSource
}

func init() {
	if err := tool.UnmarshalConfig(&g, "achilles"); err != nil {
		panic(err)
	}
}
