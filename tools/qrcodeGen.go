package tools

import (
	qrcode "github.com/skip2/go-qrcode"
)

func QrcodeGen(urlstr string) []byte {
	qrpng, err := qrcode.Encode(urlstr, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}
	return qrpng
}
