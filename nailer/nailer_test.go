package nailer

import (
	"testing"
)

func TestProcessSingleUrl(t *testing.T) {
	resized := Thumbnail("http://global3.memecdn.com/duh-im-winning-sort-of_o_873601.jpg", "20")
	if resized == "" {
    t.Error("expected an http output to be returned")
  }
}
func TestProcessEmptyUrl(t *testing.T) {
	resized := Thumbnail("", "20")
	if resized == "" {
    t.Error("expected an http output to be returned")
  }
}
// func TestProcessEmptyWidth(t *testing.T) {
// 	resized := Thumbnail("http://global3.memecdn.com/duh-im-winning-sort-of_o_873601.jpg", "")
// 	if resized == "" {
//     t.Error("expected an http output to be returned")
//   }
// }
// func TestProcessEmptyUrlAndEmptyWidth(t *testing.T) {
// 	resized := Thumbnail("", "")
// 	if resized == "" {
//     t.Error("expected an http output to be returned")
//   }
// }
