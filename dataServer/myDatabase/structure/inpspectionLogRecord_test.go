package structure

import (
	"fmt"
	"testing"
)

func TestMyMakeJSON(t *testing.T) {
	s, err := MyMakeJSON()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}
