package config

import (
	"os"

	"github.com/juju/errors"
	"gopkg.in/yaml.v3"
)

var Settings SettingsRoot

type SettingsRoot struct {
	Notifier Notifier `yaml:"notifier"`
	NATS     NATS     `yaml:"nats"`
}
type Notifier struct {
	Port string `yaml:"port"`
}

type NATS struct {
	URL   string `yaml:"url"`
	Topic string `yaml:"topic"`
}

func ParseSettings() error {
	f, err := os.Open(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return errors.Trace(err)
	}

	return errors.Trace(yaml.NewDecoder(f).Decode(&Settings))
}
