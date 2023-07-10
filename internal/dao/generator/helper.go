package generator

import (
	"math/rand"
	"time"
)

func getRandomUint8(maxExcluding int) uint8 {
	return uint8(rand.Intn(maxExcluding))
}

func getRandomStringByLength(length uint) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const num = len(letters)
	randomBytes := make([]byte, length)
	for i := range randomBytes {
		randomBytes[i] = letters[rand.Intn(num)]
	}
	return string(randomBytes)
}

func getRandomHourDuration(maxHours int) time.Duration {
	return time.Hour * time.Duration(rand.Intn(maxHours))
}
