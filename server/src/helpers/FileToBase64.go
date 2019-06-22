package helpers

import (
	"bytes"
	"encoding/base64"
	"io"
	multipart "mime/multipart"
)

func FileToBase64(f multipart.File) string {
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, f); err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buff.Bytes())

	return base64Str
}
