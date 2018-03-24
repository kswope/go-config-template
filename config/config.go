package config

import (
	"os"
	"regexp"
	"strings"

	"github.com/fatih/structs"
	"github.com/imdario/mergo"
)

// Data this is where you get your data
var Data config

// easier to call mergo.Merge with WithOverride option
func merge(dest *config, src config) {
	mergo.Merge(dest, src, mergo.WithOverride)
}

// https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Setup loads the data which is accessible from config.Data
func Setup() {

	// first merge in defaults
	merge(&Data, defaultData)

	mode, ok := os.LookupEnv("GO_ENV")

	// merge in env specific configs
	if ok {
		switch mode {
		case "dev", "": // "" means default is dev
			merge(&Data, devData)
		case "test":
			merge(&Data, testData)
		case "prod":
			merge(&Data, prodData)
		}
	}

	// try to overwrite from env variables
	for _, f := range structs.New(&Data).Fields() {
		nameCamel := toSnakeCase(f.Name())
		nameUpper := strings.ToUpper(nameCamel)
		envVal, ok := os.LookupEnv(nameUpper)
		if ok { // ok means env variable exists, even if its value is blank
			f.Set(envVal)
		}
	}

}
