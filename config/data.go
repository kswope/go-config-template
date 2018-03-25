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
// NOTE: keep stuff outta here you don't want accidentally in production, it could happen!
var defaultData = config{
	Host: "default-host",
	Port: "default-port",
}

// dev ( default mode, or activate using GO_ENV=dev )
var devData = config{
	DbName: "site_dev",
}

// test ( set using GO_ENV=test )
var testData = config{
	DbName: "site_test",
}

// test ( set using GO_ENV=prod )
var prodData = config{
	DbName: "site_prod",
}
