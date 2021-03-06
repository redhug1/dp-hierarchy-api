package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetReturnsDefaultValues(t *testing.T) {
	t.Parallel()
	Convey("When loading a configuration, default values are returned", t, func() {
		config, err := Get()
		So(err, ShouldBeNil)
		So(config.BindAddr, ShouldEqual, ":22600")
		So(config.HierarchyAPIURL, ShouldEqual, "http://localhost:22600")
		So(config.CodelistAPIURL, ShouldEqual, "http://localhost:22400")
	})
}
