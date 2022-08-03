package day00

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func part1(input string) int {
	target := "00000"
	i := 1
	for {
		in := input + strconv.Itoa(i)
		hash := getMD5String(in)
		if string(hash[:5]) == target {
			return i
		}
		i++
	}
	return -1
}

func part2(input string) int {
	target := "000000"
	i := 1
	for {
		in := input + strconv.Itoa(i)
		hash := getMD5String(in)
		if string(hash[:6]) == target {
			return i
		}
		i++
	}
	return -1
}

func getMD5String(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
