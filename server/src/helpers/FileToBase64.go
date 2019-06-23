package helpers

import (
	"bytes"
	"encoding/base64"
	"io"
	multipart "mime/multipart"
)

// FileToBase64 receives a binary representation of file from
// a request with Content-Type: "multipart/form-data" and
// returns the file encoded in base64
func FileToBase64(f multipart.File) string {
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, f); err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buff.Bytes())

	return base64Str
}
