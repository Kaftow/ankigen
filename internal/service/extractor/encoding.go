package extractor

import (
	"bytes"
	"errors"
	"io"
	"log"
	"unicode/utf8"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// decoderMap contains a wide range of common encodings
var decoderMap = map[string]*encoding.Decoder{
	"gbk":          simplifiedchinese.GBK.NewDecoder(),
	"gb2312":       simplifiedchinese.HZGB2312.NewDecoder(),
	"gb18030":      simplifiedchinese.GB18030.NewDecoder(),
	"big5":         traditionalchinese.Big5.NewDecoder(),
	"big5hkscs":    traditionalchinese.Big5.NewDecoder(),
	"shift_jis":    japanese.ShiftJIS.NewDecoder(),
	"euc-jp":       japanese.EUCJP.NewDecoder(),
	"iso-2022-jp":  japanese.ISO2022JP.NewDecoder(),
	"euc-kr":       korean.EUCKR.NewDecoder(),
	"latin1":       charmap.ISO8859_1.NewDecoder(),
	"iso-8859-1":   charmap.ISO8859_1.NewDecoder(),
	"iso-8859-2":   charmap.ISO8859_2.NewDecoder(),
	"iso-8859-3":   charmap.ISO8859_3.NewDecoder(),
	"iso-8859-4":   charmap.ISO8859_4.NewDecoder(),
	"iso-8859-5":   charmap.ISO8859_5.NewDecoder(),
	"iso-8859-6":   charmap.ISO8859_6.NewDecoder(),
	"iso-8859-7":   charmap.ISO8859_7.NewDecoder(),
	"iso-8859-8":   charmap.ISO8859_8.NewDecoder(),
	"iso-8859-9":   charmap.ISO8859_9.NewDecoder(),
	"windows-1250": charmap.Windows1250.NewDecoder(),
	"windows-1251": charmap.Windows1251.NewDecoder(),
	"windows-1252": charmap.Windows1252.NewDecoder(),
	"windows-1253": charmap.Windows1253.NewDecoder(),
	"windows-1254": charmap.Windows1254.NewDecoder(),
	"windows-1255": charmap.Windows1255.NewDecoder(),
	"windows-1256": charmap.Windows1256.NewDecoder(),
	"windows-1257": charmap.Windows1257.NewDecoder(),
	"windows-1258": charmap.Windows1258.NewDecoder(),
	"koi8-r":       charmap.KOI8R.NewDecoder(),
	"koi8-u":       charmap.KOI8U.NewDecoder(),
}

// fallbackDecoders: ordered list to try if detection fails or unsupported
// fallbackDecoders: ordered list to try if detection fails or unsupported
var fallbackDecoders = []*encoding.Decoder{
	simplifiedchinese.GBK.NewDecoder(),
	simplifiedchinese.GB18030.NewDecoder(),
	traditionalchinese.Big5.NewDecoder(),
	japanese.ShiftJIS.NewDecoder(),
	charmap.ISO8859_1.NewDecoder(),
	charmap.Windows1252.NewDecoder(),
}

// getDecoder returns decoder from map
func getDecoder(encName string) *encoding.Decoder {
	return decoderMap[encName]
}

// tryDecode attempts decoding with a decoder and validates UTF-8
func tryDecode(data []byte, dec *encoding.Decoder) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(data), dec)
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, reader)
	if err != nil {
		log.Printf("error decoding bytes: %v", err)
		return nil, err
	}
	if !utf8.Valid(buf.Bytes()) {
		log.Printf("invalid UTF-8 detected")
		return nil, errors.New("resulting bytes are not valid UTF-8")
	}
	return buf.Bytes(), nil
}

func ToUTF8(data []byte) (string, error) {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(data)

	if utf8.Valid(data) {
		return string(data), nil
	}

	if err != nil {
		result = &chardet.Result{Charset: ""}
	}

	if dec := getDecoder(result.Charset); dec != nil {
		if out, err := tryDecode(data, dec); err == nil {
			return string(out), nil
		}
	}

	for _, dec := range fallbackDecoders {
		if out, err := tryDecode(data, dec); err == nil {
			return string(out), nil
		}
	}

	log.Printf("unable to decode bytes to valid UTF-8")
	return "", errors.New("unable to decode bytes to valid UTF-8")
}
