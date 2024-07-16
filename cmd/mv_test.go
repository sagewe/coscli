package cmd

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMvCmd(t *testing.T) {
	fmt.Println("TestMvCmd")
	setUp(testBucket, testAlias, testEndpoint)
	defer tearDown(testBucket, testAlias, testEndpoint)
	Convey("Test coscli mv", t, func() {
		Convey("success", func() {
			cmd := rootCmd
			args := []string{"lsparts", fmt.Sprintf("cos://%s", testAlias)}
			cmd.SetArgs(args)
			e := cmd.Execute()
			So(e, ShouldBeNil)
		})
	})
}
