package grid_cli

import (
	"bytes"
	"context"
	"encoding/binary"
	"io/ioutil"
	"testing"
)

// TestParseMessage tests the ParseMessage function, ensuring it can parse a multihash and a string from an io.Reader.
func TestParseMessage(t *testing.T) {
	ctx := context.Background()
	data, err := ioutil.ReadFile("testdata/hello.msg")
	if err != nil {
		t.Fatalf("Failed to read test message file: %v", err)
	}

	promiseHash, args, err := deserializeMessage(data)
	if err != nil {
		t.Fatalf("Failed to deserialize message: %v", err)
	}

	expectedHash := [...]byte{0x12, 0x20, 0x50, 0x1d, 0x4b, 0x4d, 0x2e, 0x95, 0x19, 0xd5, 0x15, 0x59, 0x4e, 0x2b, 0xde, 0x57, 0xb2, 0x8f, 0x39, 0xc0, 0x4e, 0xb8, 0x9e, 0xb4, 0x3d, 0x76, 0x2a, 0x20, 0xb9, 0x1d, 0x06, 0x2b, 0x7e, 0xd2}
	expectedString := "hello world"

	if !bytes.Equal(promiseHash, expectedHash[:]) {
		t.Errorf("Expected hash %x but got %x", expectedHash[:], promiseHash)
	}

	if len(args) != 1 {
		t.Fatalf("Expected 1 argument but got %d", len(args))
	}

	arg, ok := args[0].([]byte)
	if !ok {
		t.Fatalf("Expected argument type []byte but got %T", args[0])
	}

	if string(arg) != expectedString {
		t.Errorf("Expected string %q but got %q", expectedString, string(arg))
	}
}
