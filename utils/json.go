package utils

import (
	"regexp"
	"strings"
)

var newKeys, oldKeys []string

func updateOldsKeys(key string) string {
	oldKeys = append(oldKeys, key)
	return key
}

// ProcessJSON function process the JSON returned by server to a golang readable format
func ProcessJSON(data []byte) []byte {
	// Compile regular expression
	reg := regexp.MustCompile("\"(.+?)\"")
	keys := reg.FindAllString(string(data), -1)

	// Alocate memory to store key name equivalents
	newKeys = make([]string, len(keys))
	oldKeys = make([]string, len(keys))

	// Append the new key names to array
	for _, key := range keys {
		// Split keyname from ','
		keyName := strings.ReplaceAll(strings.Split(key, ",")[0], " ", "")
		if strings.LastIndex(keyName, "\"") <= 0 {
			keyName += "\""
		}
		newKeys = append(newKeys, keyName)
	}

	// Convert input data to string
	sdata := string(data)

	// Use regex to do the conversion
	reg.ReplaceAllStringFunc(string(sdata), updateOldsKeys)

	// Replace the old keys with the new valid key names
	for j, old := range oldKeys {
		sdata = strings.ReplaceAll(sdata, old, newKeys[j])
	}

	return []byte(sdata)
}
