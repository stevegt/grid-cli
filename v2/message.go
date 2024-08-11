package grid_cli

import (
	"bytes"
	"crypto/sha256"
	"strings"

	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multihash"
	. "github.com/stevegt/goadapt"
)

type Message struct {
	Promise *multihash.DecodedMultihash
	Parms   []string // Promise Parameters
	Payload string
}

// Message marshalling and unmarshalling follows the Go
// marshal/unmarshal pattern.  A message is a promise hash followed by
// zero or more parameters, space separated.  The optional payload is
// separated by a double newline.

// NewPromise creates a new promise multihash from the given text and algorithm.
func NewPromise(promiseTxt string, algoName string) (mHash *multihash.DecodedMultihash, err error) {
	defer Return(&err)
	// convert promiseTxt to a multihash
	switch algoName {
	case "sha256":
		// get the sha256 hash of the promiseTxt
		digest := sha256.Sum256([]byte(promiseTxt))
		// create a multihash from the digest
		m, err := multihash.Encode(digest[:], multihash.SHA2_256)
		Ck(err)
		// convert the multihash to a DecodedMultihash
		mHash, err = multihash.Decode(m)
		Ck(err)
	default:
		Assert(false, "unsupported algorithm")
	}
	return mHash, nil
}

// NewMessage creates a new message with the given promise,
// parameters, and payload.
func NewMessage(promiseStr, algoName string, parms []string, payload string) (msg *Message, err error) {
	defer Return(&err)

	mHash, err := NewPromise(promiseStr, algoName)
	Ck(err)

	msg = &Message{
		Promise: mHash,
		Parms:   parms,
		Payload: payload,
	}

	return msg, nil
}

// Marshal a message to a byte slice
func Marshal(msg *Message) (buf []byte, err error) {
	defer Return(&err)

	// convert the decoded multihash to a multihash
	m, err := multihash.Encode(msg.Promise.Digest, msg.Promise.Code)
	Ck(err)

	// convert the multihash to a multibase string
	promiseStr, err := multibase.Encode(multibase.Base58BTC, m)
	Ck(err)

	parms := strings.Join(msg.Parms, " ")
	header := Spf("%s %s", promiseStr, parms)
	txt := header
	if len(msg.Payload) > 0 {
		txt = Spf("%s\n\n%s", txt, msg.Payload)
	}

	buf = []byte(txt)
	return buf, nil
}

func Unmarshal(data []byte, m *Message) (err error) {
	defer Return(&err)
	// Split the data into the header and payload
	parts := bytes.SplitN(data, []byte("\n\n"), 2)
	// payload is optional
	if len(parts) == 2 {
		m.Payload = string(parts[1])
	}
	// Split the header into the promise and parameters
	header := strings.Fields(string(parts[0]))
	// promise is required
	Assert(len(header) > 0, "invalid message header; missing promise hash")
	promiseBuf := header[0]
	// parameters are optional
	if len(header) > 1 {
		m.Parms = strings.Fields(string(header[1]))
	}
	// decode the promise hash using multibase and multihash
	_, buf, err := multibase.Decode(promiseBuf)
	Ck(err)
	m.Promise, err = multihash.Decode(buf)
	Ck(err)
	return nil
}
