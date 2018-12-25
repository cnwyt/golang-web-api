package tools

import "fmt"
import "testing"

func TestDatetime(t *testing.T) {
	input := "1545742069"
    got := GetDatetime(input)
    expected := "2018-12-25 20:47:49"
    if got != expected {
        t.Errorf("--> 出错了: GetDatetime(): input: %q, expected: %q, got: %q", input, got, expected)
    }
}