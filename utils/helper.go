package utils

import (
	"regexp"
	"strconv"
)

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
