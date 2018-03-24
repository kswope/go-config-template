package config

// All strings. Reasoning: all these are overwritable by env variables, and env
// variables are always strings.  Also, the nil value of number types are a
// surprise for configs.
type config struct {
	Host   string
	Port   string
	DbHost string
	DbUser string
	DbPass string
	DbName string
}

// NOTE: DB_HOST in ENV will overwrite DbHost
// NOTE: PORT in ENV will overwrite Port

// defaults inherited by all other modes below
var defaultData = config{
	Host: "localhost",
	Port: "300",
}

// dev ( default mode, or activate using GO_ENV=dev )
var devData = config{
	DbName: "wheatt_dev",
}

// test ( set using GO_ENV=test )
var testData = config{
	DbName: "wheatt_test",
}

// test ( set using GO_ENV=prod )
var prodData = config{
	DbName: "wheatt_prod",
}
