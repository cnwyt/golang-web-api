package package1

import "testing"

func TestMax(t *testing.T) {
    input := "WORLD"
    got := SayHello(input)
    expected := "Hello,WORLD"
    if got != expected {
        t.Error("--> SayHello 出错了！")
        t.Errorf("--> 出错了: SayHello(): input: %q, expected: %q, got: %q", input, got, expected)
    }
}