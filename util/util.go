package util

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

// RequireJSON ...
func RequireJSON(filepath string, refer interface{}) error {
	raw, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}
	re := regexp.MustCompile("(?s)[^https?:]//.*?\n|/\\*.*?\\*/")
	valid := re.ReplaceAll(raw, nil)
	if err := json.Unmarshal(valid, &refer); err != nil {
		return err
	}
	return nil
}
