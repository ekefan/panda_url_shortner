package util

import (
	"fmt"
	"math/rand"
	"strings"
)
func RandomShortCode(l int) (string, error){
	letters := "abcdefghijklmnopqrstuvwxyz"
	var shortCode strings.Builder
	for i := 0; i < l; i++{
		idx := rand.Intn(len(letters))
		err := shortCode.WriteByte(letters[idx])
		if err != nil {
			return "", fmt.Errorf("couldn't build shortcode: %v", err)
		}
	}
	return shortCode.String(), nil
}