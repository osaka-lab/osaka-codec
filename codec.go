package main

import (
	"fmt"
	"strings"
)

func getMap() (map[int]string, map[string]int) {
	reverseMapping := make(map[string]int)
	mapping := map[int]string{
		0: "Osaka Shoe Kick",
		1: "Hello Everynyan",
		2: "America ya",
		3: "Oh my gah",
		4: "Why did you become a teacher, Mr. Kimura?",
		5: "Sata Andagi!!!",
		6: "Mai Waifu",
		7: "You know those times when you can see the dust in your eyes? I'm chasing it right now.",
		8: "Elevator, escalator, escalator, elevator...",
		9: "Kamineko",
	}

	for i := 10; i < 256; i++ {
		intVal := fmt.Sprintf("%d", i)
		var sb strings.Builder

		for _, v := range intVal {
			singleByte := int(v - '0')
			p := mapping[singleByte]
			sb.WriteString(p)
		}

		mapping[i] = sb.String()
	}

	for k, v := range mapping {
		reverseMapping[v] = k
	}

	return mapping, reverseMapping
}

func Encode(bytes []byte) (string, error) {
	var encoded strings.Builder
	mapping, _ := getMap()

	for _, v := range bytes {
		intByte := int(v)

		str, exists := mapping[intByte]
		if exists {
			encoded.WriteString(str)
		} else {
			return "", fmt.Errorf("no mapping for byte '%d'", intByte)
		}
		encoded.WriteString(";")
	}

	return encoded.String(), nil
}

func Decode(encoded string) ([]byte, error) {
	var decoded []byte
	divided := strings.Split(encoded, ";")

	_, reverseMapping := getMap()

	for _, str := range divided {
		if str == "" || str == "\n" {
			continue
		}

		if key, found := reverseMapping[str]; found {
			decoded = append(decoded, byte(key))
		} else {
			return nil, fmt.Errorf("no byte found for encoded string '%s'", str)
		}
	}

	return decoded, nil
}
