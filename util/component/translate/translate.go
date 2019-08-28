package translate

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

func Utf8ToGbk(r io.Reader) io.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewEncoder())
}

func GbkToUtf8(r io.Reader) io.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewDecoder())
}
