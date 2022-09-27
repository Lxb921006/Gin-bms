package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type GoogleAuthenticator struct {
	Secret string //The base32NoPaddingEncodedSecret parameter is an arbitrary key value encoded in Base32 according to RFC 3548. The padding specified in RFC 3548 section 2.2 is not required and should be omitted.
	Expire uint64 //更新周期单位秒
	Digits int    //数字数量
}

func (m *GoogleAuthenticator) GaCode() (code string, err error) {
	count := uint64(time.Now().Unix()) / m.Expire
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(m.Secret)
	if err != nil {

		return
	}
	codeInt := hotp(key, count, m.Digits)
	intFormat := fmt.Sprintf("%%0%dd", m.Digits) //数字长度补零
	return fmt.Sprintf(intFormat, codeInt), nil
}

//QrString google authenticator 扫描二维码的二维码字符串
func (m *GoogleAuthenticator) QrUrl(label, user string) (qr string) {
	m.CreateSecret(user)
	flabel := url.QueryEscape(label) //有一些小程序MFA不支持
	//规范文档 https://github.com/google/google-authenticator/wiki/Key-Uri-Format
	//otpauth://totp/ACME%20Co:john.doe@email.com?secret=HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ&issuer=ACME%20Co&algorithm=SHA1&digits=6&period=30
	qr = fmt.Sprintf(`otpauth://totp/%s?secret=%s&issuer=%s&algorithm=SHA1&digits=%d&period=%d`, user, m.Secret, flabel, m.Digits, m.Expire)
	return
}

func (m *GoogleAuthenticator) CreateSecret(u string) {
	date := strconv.Itoa(int(time.Now().Nanosecond()))
	data := "ajksduk912J3KDAKJKASD" + u + date
	hash := sha1.New()
	hash.Write([]byte(data))
	nd := hash.Sum(nil)
	nnd := hex.EncodeToString(nd)
	key := base32.StdEncoding.EncodeToString([]byte(nnd))
	sha1String := strings.Split(key, "=")
	m.Secret = sha1String[0]
}

func NewGoogleAuthenticator(key string) *GoogleAuthenticator {
	return &GoogleAuthenticator{
		Secret: key,
		Expire: 30,
		Digits: 6,
	}
}

func hotp(key []byte, counter uint64, digits int) int {
	//RFC 6238
	//只支持sha1
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	//取sha1的最后4byte
	//0x7FFFFFFF 是long int的最大值
	//math.MaxUint32 == 2^32-1
	//& 0x7FFFFFFF == 2^31  Set the first bit of truncatedHash to zero  //remove the most significant bit
	// len(sum)-1]&0x0F 最后 像登陆 (bytes.len-4)
	//取sha1 bytes的最后4byte 转换成 uint32
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)
	//取十进制的余数
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}
