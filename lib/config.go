package lib

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration is the base yaml object.
type Configuration struct {
	Token        string `yaml:"token"`
	BotOwner     string `yaml:"bot_owner"`
	BotID        string `yaml:"bot_id"`
	ElevatedRole string `yaml:"elevated_role"`
	Prefix       string `yaml:"prefix"`
}

// Config retrieves the app's configuration form config.json.
func Config() Configuration {
	file, err := os.Open(LocGet("config.yml"))
	Check(err)
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	Check(err)

	var config Configuration
	err = yaml.Unmarshal(contents, &config)
	Check(err)

	return config
}
