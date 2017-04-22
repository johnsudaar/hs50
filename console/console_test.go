package console

import (
	"testing"

	"github.com/johnsudaar/hs50/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPackDommand(t *testing.T) {
	Convey("Given a simple command", t, func() {
		command := config.SET_BUS
		params := []string{config.BUS_A, config.INPUT_1}

		query := PackCommand(command, params...)

		So(query[0], ShouldEqual, config.START_TRANSMISSION)
		So(query[len(query)-1], ShouldEqual, config.END_TRANSMISSION)

		c := string(query[1 : len(query)-1])
		So(c, ShouldEqual, config.SET_BUS+config.BUS_A+config.INPUT_1)
	})
}

func TestUnpackCommand(t *testing.T) {
	Convey("Given a wrong start packet", t, func() {
		query := []byte{0xff, 0xde, 0xad, 0xbe, 0xef, config.END_TRANSMISSION}
		_, _, err := UnPackCommand(query)

		So(err, ShouldNotBeNil)
	})

	Convey("Given a wrong end packet", t, func() {
		query := []byte{config.START_TRANSMISSION, 0xde, 0xad, 0xbe, 0xef, 0xff}
		_, _, err := UnPackCommand(query)

		So(err, ShouldNotBeNil)
	})

	Convey("Given an invalid command", t, func() {
		query := PackCommand("WRONG", ":00")

		_, _, err := UnPackCommand(query)
		So(err, ShouldNotBeNil)
	})

	Convey("Given an invalid parameter", t, func() {
		query := PackCommand(config.SET_BUS, ":000")

		_, _, err := UnPackCommand(query)
		So(err, ShouldNotBeNil)
	})

	Convey("Given a command without parameters", t, func() {
		query := PackCommand(config.SET_BUS)
		cmd, p, err := UnPackCommand(query)

		So(err, ShouldBeNil)
		So(cmd, ShouldEqual, config.SET_BUS)
		So(len(p), ShouldEqual, 0)
	})

	Convey("Given a command with many parameters", t, func() {
		query := PackCommand(config.SET_BUS, config.INPUT_1, config.INPUT_2, config.INPUT_3)
		cmd, p, err := UnPackCommand(query)

		So(err, ShouldBeNil)
		So(cmd, ShouldEqual, config.SET_BUS)
		So(len(p), ShouldEqual, 3)
		So(p[0], ShouldEqual, config.INPUT_1)
		So(p[1], ShouldEqual, config.INPUT_2)
		So(p[2], ShouldEqual, config.INPUT_3)
	})
}
