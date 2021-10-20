package main

import "testing"

func TestHelloYou(t *testing.T) {
    got := HelloYou("Chris")
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}