package encrypt

import "golang.org/x/crypto/bcrypt"

//对明文密码进行加密
func GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//传入哈希和明文密码，进行比较
func ComparePW(hashedPW string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(password))
	return err == nil
}
