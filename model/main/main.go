package main

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

//	func genMd5(s string) string {
//		h := md5.New()
//		h.Write([]byte(s))
//		return hex.EncodeToString(h.Sum(nil))
//	}
func main() {
	// Using custom options
	options := &password.Options{16, 100, 50, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	NewPasswoed := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	// fmt.Println(NewPasswoed)
	//fmt.Println(len(NewPasswoed))
	passwordInfo := strings.Split(NewPasswoed, "$")
	//fmt.Println(passwordInfo)
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
	// fmt.Println(salt, encodedPwd)
}
