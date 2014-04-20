package nailer

import (
	"testing"
)

func TestProcessSingleExample(t *testing.T) {
	resized := Process("https://pbs.twimg.com/profile_images/378800000667589936/19ade83790a92ef4cb4d5d1fb3a6e5dd_bigger.jpeg", "20")
	if resized == "" {
    t.Error("expected an http output to be returned")
  }
}
