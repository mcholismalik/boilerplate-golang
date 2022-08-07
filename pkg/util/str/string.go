package str

import (
	"math/rand"
	"strings"

	"github.com/dustin/go-humanize"
)

func GenerateRandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ConvertDayFromEnToIdn(day string) string {
	switch day {
	case "monday":
		return "senin"
	case "tuesday":
		return "selasa"
	case "wednesday":
		return "rabu"
	case "thursday":
		return "kamis"
	case "friday":
		return "jum'at"
	case "saturday":
		return "sabtu"
	case "sunday":
		return "minggu"
	}
	return ""
}

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}
