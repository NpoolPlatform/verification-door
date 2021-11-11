package google

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func toBytes(value int64) []byte {
	result := []byte{}
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func oneTimePassword(key, data []byte) uint32 {
	hash := hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] &= 0x7F
	number := toUint32(hashParts)
	return number % 1000000
}

// get secret
func GetSecret() (string, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, time.Now().UnixNano()/1000/30)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(base32encode(hmacSha1(buf.Bytes(), nil))), nil
}

// get qrcode url
func GetQrcodeURL(user, secret string) string {
	qrcode := fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
	width := "200"
	height := "200"
	data := url.Values{}
	data.Set("data", qrcode)
	return "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
}

// get code
func getTrueCode(secret string) ([]string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := base32decode(secretUpper)
	if err != nil {
		return nil, err
	}
	codeNow := oneTimePassword(secretKey, toBytes(time.Now().UTC().Unix()/30))
	codePast30 := oneTimePassword(secretKey, toBytes((time.Now().UTC().Unix()-30)/30))
	codeAfter30 := oneTimePassword(secretKey, toBytes((time.Now().UTC().Unix()+30)/30))

	codeGroup := []string{
		fmt.Sprintf("%06d", codeNow),
		fmt.Sprintf("%06d", codePast30),
		fmt.Sprintf("%06d", codeAfter30),
	}
	return codeGroup, nil
}

// verify code
func VerifyCode(secret, code string) (bool, error) {
	trueCodes, err := getTrueCode(secret)
	if err != nil {
		return false, err
	}

	for _, trueCode := range trueCodes {
		if code != trueCode {
			continue
		}
		return true, nil
	}
	return false, nil
}
