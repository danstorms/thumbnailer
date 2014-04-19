package nailer

import (
	"testing"
	"os"
)

func TestProcessSingleExample(t *testing.T) {
	resized := Process("../public/storms.jpg")
	if _, err := os.Stat(resized); err != nil {
    if os.IsNotExist(err) {
        t.Error("expected the output file to be existing")
    } else {
        t.Error("expected to not get another error")
    }
  }
}
