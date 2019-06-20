package helpers

import (
	"io"
	"fmt"
	"bytes"
	"encoding/base64"
	multipart "mime/multipart"
)

func FileToBase64(f multipart.File) {
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, f); err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buff.Bytes())
	fmt.Println(base64Str)
}
