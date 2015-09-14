package com

import (
	"bytes"
	"encoding/gob"
)

func EncodeGob(src interface{}) (string, error) {

	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(src)
	if err != nil {
		return "", err
	}

	return string(b.Bytes()), nil
}

func DecodeGob(data string, des interface{}) error {
	b := bytes.NewReader([]byte(data))

	dec := gob.NewDecoder(b)

	err := dec.Decode(des)
	if err != nil {
		return err
	}

	return nil
}
