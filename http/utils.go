package http

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strconv"
)

const tagName = "param"

//SwitchBody switches response body data
func SwitchBody(resp Response, data []byte) Response {
	resp.ClearInternalBuffer()
	buf := ioutil.NopCloser(bytes.NewBuffer(data))
	resp.RawResponse().Body = buf

	return resp
}

//ToMap turn objects in maps
func ToMap(obj interface{}) map[string]string {

	mapString := make(map[string]string)
	t := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)

	if t != nil {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := values.Field(i)

			tag := field.Tag.Get(tagName)
			if tag != "" {
				switch v := value.Interface().(type) {
				case string:
					mapString[tag] = v
				case *string:
					if v != nil {
						mapString[tag] = *v
					}
				case int:
					mapString[tag] = strconv.Itoa(v)
				case *int:
					if v != nil {
						mapString[tag] = strconv.Itoa(*v)
					}
				}
			}
		}
	}

	return mapString
}
