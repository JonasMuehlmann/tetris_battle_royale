package userService

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/argon2"
)

const saltLength = 20
const AASCIAlnumRange = int('~' - ' ')
const ASCIIAlnumOffsetFromStart = int(' ')

func generateSalt(length int) []byte {
	salt := make([]byte, 0, saltLength)

	for i := 0; i < length; i++ {
		saltByte, _ := rand.Int(rand.Reader, big.NewInt(int64(AASCIAlnumRange)))
		salt = append(salt, byte(int(saltByte.Int64())+ASCIIAlnumOffsetFromStart))
	}

	return salt
}

func hashPw(pw, salt []byte) []byte {
	return argon2.IDKey(pw, salt, 1, 64*1024, 4, 32)
}
