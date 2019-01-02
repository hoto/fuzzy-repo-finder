package config

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// all fields must be public in order for marshalling to work
type ymlConfig struct {
	ProjectRoots []string `yaml:"project_roots,flow"`
}

func newYmlConfig() ymlConfig {
	return readYmlConfig()
}

func readYmlConfig() ymlConfig {
	ymlConfig := ymlConfig{}
	configYmlBytes := readConfigFromDisk()
	err := yaml.Unmarshal([]byte(configYmlBytes), &ymlConfig)
	check(err)
	ymlConfigBytes, err := yaml.Marshal(&ymlConfig)
	check(err)

	return expandEnvs(ymlConfigBytes)
}

func expandEnvs(ymlConfigBytes []byte) ymlConfig {
	expandedYmlConfig := ymlConfig{}
	expandedStringYmlConfig := os.ExpandEnv(string(ymlConfigBytes))
	err := yaml.Unmarshal([]byte(expandedStringYmlConfig), &expandedYmlConfig)
	check(err)
	//debugLog(ymlConfigBytes, expandedStringYmlConfig, expandedYmlConfig)
	return expandedYmlConfig
}

func debugLog(rawConfig []byte, expandedConfig string, structConfig ymlConfig) {
	fmt.Printf("Config path=%s\n\n", Cyan(configFile))
	fmt.Printf("Raw yml config file:\n%s\n", Cyan(string(rawConfig)))
	fmt.Printf("Expanded yml config file:\n%s\n", Cyan(expandedConfig))
	fmt.Printf("Expanded struct config file:\n%s\n", Cyan(structConfig))
}

func readConfigFromDisk() []byte {
	bytes, err := ioutil.ReadFile(configFile)
	check(err)
	return bytes
}

func (c ymlConfig) String() string {
	return fmt.Sprintf("ProjectRoots=[%s]", c.ProjectRoots)
}
