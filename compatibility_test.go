package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"testing"
)

// Shared fixtures were produced by the original Go encoder.
func TestCompatibility(t *testing.T) {
	data, err := os.ReadFile("testdata/compatibility.json")
	if err != nil {
		t.Fatal(err)
	}
	var vectors []struct {
		Hex     string
		Encoded string
	}
	if err := json.Unmarshal(data, &vectors); err != nil {
		t.Fatal(err)
	}
	for _, v := range vectors {
		input, err := hex.DecodeString(v.Hex)
		if err != nil {
			t.Fatal(err)
		}
		if got := encode(input); got != v.Encoded {
			t.Fatalf("encode(%x) = %q, want %q", input, got, v.Encoded)
		}
		got, err := decode(v.Encoded)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, input) {
			t.Fatalf("decode(%q) = %x, want %x", v.Encoded, got, input)
		}
	}
}
