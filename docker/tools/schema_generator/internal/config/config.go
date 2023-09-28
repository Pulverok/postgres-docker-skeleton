package config

import (
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"
)

type AppConfig struct {
	DataPath string `long:"data_path" env:"DATA_PATH" description:"Path to the data folder" required:"true"`
}

// ErrHelp is returned when --help flag is
// used and application should not launch.
var ErrHelp = errors.New("help")

// New that corresponds to the values read.
func New() (*AppConfig, error) {
	var config AppConfig

	if _, err := flags.Parse(&config); err != nil {
		var flagsErr *flags.Error
		if errors.As(err, &flagsErr) && flagsErr.Type == flags.ErrHelp {
			return nil, ErrHelp
		}
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &config, nil
}
