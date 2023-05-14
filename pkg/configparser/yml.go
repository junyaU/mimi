package configparser

import (
	"fmt"
	"github.com/spf13/viper"
)

type YmlConfig struct {
	Version  string          `yaml:"version"`
	Commands []ConfigCommand `yaml:"commands"`
}

type ConfigCommand struct {
	Name       string        `yaml:"name"`
	Parameters CommandParams `yaml:"parameters"`
}

type CommandParams struct {
	Path              string `yaml:"path"`
	DirectThreshold   int    `yaml:"directThreshold"`
	InDirectThreshold int    `yaml:"inDirectThreshold"`
}

type Command struct {
	Name              string
	Path              string
	DirectThreshold   int
	IndirectThreshold int
}

func (c *ConfigCommand) IsVaild() bool {
	switch c.Name {
	case "check", "list", "table":
	default:
		return false
	}

	if c.Parameters.Path == "" {
		return false
	}

	return true
}

func NewYmlConfig(path string) (*YmlConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("mimi")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config *YmlConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %v", err)
	}

	return config, nil
}

func (c *YmlConfig) GetCommands() ([]Command, error) {
	var commands []Command
	for _, command := range c.Commands {
		if !command.IsVaild() {
			return nil, fmt.Errorf("invalid command: %v", command.Name)
		}

		commands = append(commands, Command{
			Name:              command.Name,
			Path:              command.Parameters.Path,
			DirectThreshold:   command.Parameters.DirectThreshold,
			IndirectThreshold: command.Parameters.InDirectThreshold,
		})
	}

	return commands, nil
}
