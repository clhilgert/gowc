package main

import "testing"

func TestTestMain(t *testing.T) {
	got := testMain()
	want := "test"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
