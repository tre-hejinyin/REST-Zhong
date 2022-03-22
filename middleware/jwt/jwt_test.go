package jwt

import (
	"encoding/base64"
	"testing"
)

func TestBase64Decoding(t *testing.T) {
	str := "VFJFLWNoaW5hLWFwcGFyZWw="
	bs, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bs))
}
