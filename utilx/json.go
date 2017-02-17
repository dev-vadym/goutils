package utilx

import "encoding/json"

func EnJson(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func EnJsonStr(v interface{}) string {
	jsonstr, _ := EnJson(v)
	return jsonstr
}

func DeJson(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}