package test

import "testing"

func TestTest2(t *testing.T) {
	Test2()
	t.Errorf("something %q", "abc")
}
