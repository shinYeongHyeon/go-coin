package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ToBytes parameter Anything
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleError(encoder.Encode(i))

	return aBuffer.Bytes()
}
