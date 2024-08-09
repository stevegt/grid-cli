package grid_cli

import (
	"testing"

	. "github.com/stevegt/goadapt"
)

// TestMarshalMessage tests the message.Marshal function
func TestMarshalMessage(t *testing.T) {
	promise, err := NewPromise("I will say hello", "sha256")
	msg := Message{
		Promise: promise,
		Parms:   []string{"hello"},
		Payload: "world",
	}

	// Marshal the message
	data, err := Marshal(&msg)
	Tassert(t, err == nil, "Failed to marshal message: %v", err)

	Pl(string(data))

	// Check the marshalled data
	// XXX
}

/*
// TestUnMarshalMessage tests the message.UnMarshal function
func TestUnMarshalMessage(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/hello.msg")
	Tassert(t, err == nil, "Failed to read test message file: %v", err)

	// Unmarshal the message
	var msg Message
	err = Unmarshal(data, &msg)
	Tassert(t, err == nil, "Failed to unmarshal message: %v", err)

	// Check the message fields
	spew.Dump(msg)

	Tassert(t, len(msg.Parms) == 1, "Expected 1 argument but got %d", len(msg.Parms))
	Tassert(t, msg.Parms[0] == "hello", "Expected string %q but got %q", "hello", msg.Parms[0])

	Tassert(t, msg.Payload == "world", "Expected string %q but got %q", "world", msg.Payload)

}
*/
