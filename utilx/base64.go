package utilx

import (
	"encoding/base64"
	"strings"
)


func EnBase64(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

func DeBase64(v string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


func EnUrlBase64(v string) string {
	return base64.URLEncoding.EncodeToString([]byte(v))
}

func DeUrlBase64(v string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EnParam(v string) string {
	data := EnBase64(v)
	data = strings.Replace(data, "/", "_x", -1)
	data = strings.Replace(data, "+", "_y", -1)
	data = strings.Replace(data, "=", "_z", -1)
	return data
}

func DeParam(v string) (string, error) {
	data := v
	data = strings.Replace(data, "_x", "/", -1)
	data = strings.Replace(data, "_y", "+", -1)
	data = strings.Replace(data, "_z", "=", -1)
	data, err := DeBase64(data)
	if err != nil {
		return "", err
	}
	return data, nil
}
