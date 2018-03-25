# go-config-template
A config library you copy into your own project.

## Features

* simple
* easily hackable
* default values
* dev, test, and prod ( or whatever, easily changed if you hack the code ) set with GO_ENV=dev
* ENV variables will overwrite values

## Precedence pattern

#### Higher overwrites lower

1. ENV values 
2. mode values: GO_ENV=mode
3. default values

## Installation

copy config.go and data.go ( and optionally config_test.go ) into your config
directory

```
├── config
│   ├── config.go
│   ├── config_test.go
│   └── data.go
└── main.go
```

### Setup all data in data.go

Edit the main struct for all your data

```
type config struct {
  Host   string
  Port   string
  DbHost string
  DbUser string
  DbPass string
  DbName string
}
```

#### Edit default data
NOTE: keep stuff outta here you don't want accidentally in production, it could happen!
```
var defaultData = config{
  Host: "localhost",
  Port: "300",
}

```

#### Create structs for the modes
```
var devData = config{
  DbName: "site_dev",
}

var testData = config{
  DbName: "site_test",
}

var prodData = config{
  DbName: "site_prod",
}
```

#### Load and use
```
func main() {

  config.Setup()
  fmt.Println(config.Data.DbName)

}


```

#### Put this amazing package to use
```
GO_ENV=prod go run main.go
```

Without GO_ENV set it defaults to "dev"

