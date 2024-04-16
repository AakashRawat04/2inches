package lib

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateShortCode() (string, error) {
	return gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 6)
}
