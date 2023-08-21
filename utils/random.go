package utils

import (
	"math/rand"
	"time"
)

var (
	num       = "0123456789"
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r         *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(n int64) int64 {
	return r.Int63n(n)
}

func RandomString(min, max int, ch ...string) string {
	desiredLength := r.Intn(max-min-len(ch)+1) + min + len(ch)

	randomString := make([]byte, desiredLength)
	remainingChars := ""
	for i, c := range ch {
		randomIndex := r.Intn(len(c))
		randomString[i] = c[randomIndex]
		remainingChars += c
	}

	for i := len(ch); i < desiredLength; i++ {
		randomString[i] = remainingChars[r.Intn(len(remainingChars))]
	}

	r.Shuffle(len(randomString), func(i, j int) {
		randomString[i], randomString[j] = randomString[j], randomString[i]
	})

	return string(randomString)
}

func RandomStringWithAlp(min, max int) string {
	return RandomString(min, max, lowerCase, upperCase)
}

func RandomStringWithAllChar(min, max int) string {
	return RandomString(min, max, lowerCase, upperCase, num)
}
