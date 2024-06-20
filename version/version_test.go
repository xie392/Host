package version_test

import (
	"fmt"
	"github.com/xie392/restful-api/version"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println("version_test.TestVersion", version.FullVersion())
}
