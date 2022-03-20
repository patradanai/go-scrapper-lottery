package utils

import (
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func GetDate(date string) map[string]string {
	r := regexp.MustCompile("(?P<Day>\\d{1,2}) (?P<Monthly>.*) (?P<Year>\\d{4})")
	matchs := r.FindStringSubmatch(date)
	params := make(map[string]string)

	for i, name := range r.SubexpNames() {
		params[name] = matchs[i]
	}

	return params
}

func ConvMonthlyToNum(monthly string) int {
	textMonthly := map[string]int{
		"มกราคม":     1,
		"กุมภาพันธ์": 2,
		"มีนาคม":     3,
		"เมษายน":     4,
		"พฤษภาคม":    5,
		"มิถุนายน":   6,
		"กรกฎาคม":    7,
		"สิงหาคม":    8,
		"กันยายน":    9,
		"ตุลาคม":     10,
		"พฤศจิกายน":  11,
		"ธันวาคม":    12,
	}

	return textMonthly[monthly]
}

func Filter(text string) string {
	r := regexp.MustCompile("รางวัลละ (?P<Prize>.*?) บาท")
	matchs := r.FindStringSubmatch(text)
	params := make(map[string]string)

	for i, name := range r.SubexpNames() {
		params[name] = matchs[i]
	}

	return params["Prize"]
}

func ConvToFloat32(myString string) float32 {
	value, _ := strconv.ParseFloat(myString, 32)

	return float32(value)
}

func ConvToInteger(myString string) int {
	value, _ := strconv.Atoi(myString)

	return value
}

func Contains(s []string, t string) bool {
	for _, val := range s {

		if val == t {
			return true
		}
	}

	return false
}

func GenUUID() string {
	uuidNew := uuid.New()
	reg := regexp.MustCompile("[-]")

	return reg.ReplaceAllString(uuidNew.String(), "")
}
