package config_test

import (
	"os"
	"testing"

	"github.com/kswope/viper-experiment/config"
	. "github.com/smartystreets/goconvey/convey"
)

//// NOTE: This is wacky to test because we're testing hardcoded values.

func Test(t *testing.T) {

	Convey("dev", t, func() {

		os.Setenv("GO_ENV", "dev")
		config.Setup()

		Convey("some defaults", func() {
			So(config.Data.Port, ShouldEqual, "300")
			So(config.Data.Host, ShouldEqual, "localhost")
		})

		Convey("data", func() {
			So(config.Data.DbName, ShouldEqual, "wheatt_dev")
		})

	})

	Convey("test", t, func() {

		os.Setenv("GO_ENV", "test")
		config.Setup()

		Convey("some defaults", func() {
			So(config.Data.Port, ShouldEqual, "300")
			So(config.Data.Host, ShouldEqual, "localhost")
		})

		Convey("data", func() {
			Print(config.Data)
			So(config.Data.DbName, ShouldEqual, "wheatt_test")
		})

	})

	Convey("dealing with env variable", t, func() {

		os.Setenv("GO_ENV", "test")
		config.Setup()

		originalDbName := config.Data.DbName
		originalPort := config.Data.Port

		// restore so we don't screw up other tests
		defer func() {
			config.Data.DbName = originalDbName
			config.Data.Port = originalPort
		}()

		os.Setenv("DB_NAME", "bogus_database_name")
		os.Setenv("PORT", "5555")
		config.Setup()

		So(config.Data.DbName, ShouldEqual, "bogus_database_name")
		So(config.Data.Port, ShouldEqual, "5555")

		//// make sure blank env variables count (blank passwords, etc)

		os.Setenv("DB_NAME", "")
		os.Setenv("PORT", "")
		config.Setup()

		So(config.Data.DbName, ShouldEqual, "")
		So(config.Data.Port, ShouldEqual, "")

	})

}
