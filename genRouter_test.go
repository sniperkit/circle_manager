package circle_manager

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}
func TestRun(t *testing.T) {
	// rawRouterFile, err := getRouter()
	// assert.Nil(t, err)

	// spew.Dump(rawRouterFile)
	// //spew.Printf("%+v", rawRouterFile)
	// assert.NotNil(t, rawRouterFile)
	// assert.NotEqual(t, 0, len(rawRouterFile.RouterItems))
}
