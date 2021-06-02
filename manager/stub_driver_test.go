package manager

import (
	"fmt"
	"testing"
)

func TestNewStubDriver(t *testing.T) {
	app := "/Users/daminyang/Downloads/hello/check"
	driver, err := NewStubDriver(app)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer driver.Exited()
	fmt.Println(driver.Greet())
}
