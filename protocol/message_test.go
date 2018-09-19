package protocol

import (
	"testing"
	"github.com/hedianbin/godcoin/util/hashx"
	"encoding/gob"
	"bytes"
)



func Test_Gob_MsgGetBlocks(t *testing.T) {
	msg := NewMsgGetBlocks(*hashx.ZeroHash())
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(msg)
	if err != nil{
		t.Error("gob.encode error", err)
	}

	bufDecode := bytes.NewBuffer(buf.Bytes())
	dec := gob.NewDecoder(bufDecode)
	var out MsgGetBlocks
	err = dec.Decode(&out)
	if err != nil{
		t.Error("gob.Decode error", err)
	}

	t.Log(out)
}