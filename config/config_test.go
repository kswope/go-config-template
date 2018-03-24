package config_test

import (
	"os"
	"testing"

	"github.com/kswope/go-config-template/config"
	. "github.com/smartystreets/goconvey/convey"
)

//// NOTE: This is wacky to test because we're testing hardcoded values.

func Test(t *testing.T) {

	Convey("GO_ENV doesn't exist", t, func() {

		os.Unsetenv("GO_ENV")
		config.Setup()

		Convey("defaults are working", func() {
			So(config.Data.Port, ShouldEqual, "default-port")
			So(config.Data.Host, ShouldEqual, "default-host")
		})

	})

	Convey("GO_ENV=dev", t, func() {

		os.Setenv("GO_ENV", "dev")
		config.Setup()

		Convey("defaults are working", func() {
			So(config.Data.Port, ShouldEqual, "default-port")
			So(config.Data.Host, ShouldEqual, "default-host")
		})

		Convey("mode data is populating", func() {
			So(config.Data.DbName, ShouldEqual, "site_dev")
		})

	})

	Convey("GO_ENV=test", t, func() {

		os.Setenv("GO_ENV", "test")
		config.Setup()

		Convey("defaults are working", func() {
			So(config.Data.Port, ShouldEqual, "default-port")
			So(config.Data.Host, ShouldEqual, "default-host")
		})

		Convey("mode data is populating", func() {
			So(config.Data.DbName, ShouldEqual, "site_test")
		})

	})

	Convey("ENV variables work", t, func() {

		os.Setenv("GO_ENV", "prod")
		config.Setup()

		originalDbName := config.Data.DbName
		originalPort := config.Data.Port

		// restore so we don't screw up other tests
		defer func() {
			config.Data.DbName = originalDbName
			config.Data.Port = originalPort
		}()

		// override with env variables
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
