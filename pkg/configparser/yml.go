package configparser

import (
	"fmt"
	"github.com/spf13/viper"
)

// YmlConfig represents the entire structure of a YAML configuration file.
type YmlConfig struct {
	Version  string          `yaml:"version"`
	Commands []ConfigCommand `yaml:"commands"`
}

// ConfigCommand represents a single command from the configuration.
type ConfigCommand struct {
	Name       string        `yaml:"name"`
	Parameters CommandParams `yaml:"parameters"`
}

// CommandParams holds the parameters of a command from the configuration.
type CommandParams struct {
	Path               string  `yaml:"path"`
	DirectThreshold    int     `yaml:"directThreshold"`
	IndirectThreshold  int     `yaml:"indirectThreshold"`
	DepthThreshold     int     `yaml:"depthThreshold"`
	LinesThreshold     int     `yaml:"linesThreshold"`
	DependentThreshold int     `yaml:"dependentThreshold"`
	WeightThreshold    float32 `yaml:"weightThreshold"`
	EnableWeight       bool    `yaml:"enableWeight"`
}

// Command represents a validated command ready for execution.
type Command struct {
	Name               string
	Path               string
	DirectThreshold    int
	IndirectThreshold  int
	DepthThreshold     int
	LinesThreshold     int
	DependentThreshold int
	WeightThreshold    float32
	EnableWeight       bool
}

func (c *ConfigCommand) IsValid() bool {
	switch c.Name {
	case "check", "list", "table", "deps":
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
		if !command.IsValid() {
			return nil, fmt.Errorf("invalid command: %v", command.Name)
		}

		commands = append(commands, Command{
			Name:               command.Name,
			Path:               command.Parameters.Path,
			DirectThreshold:    command.Parameters.DirectThreshold,
			IndirectThreshold:  command.Parameters.IndirectThreshold,
			DepthThreshold:     command.Parameters.DepthThreshold,
			LinesThreshold:     command.Parameters.LinesThreshold,
			DependentThreshold: command.Parameters.DependentThreshold,
			WeightThreshold:    command.Parameters.WeightThreshold,
			EnableWeight:       command.Parameters.EnableWeight,
		})
	}

	return commands, nil
}
