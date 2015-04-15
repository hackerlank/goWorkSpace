package config

import (
	"io/ioutil"
	"path/filepath"
	"github.com/bitly/simplejson"
	"reflect"
	"fmt"
)

func GetSpecConfig(filename string, keys ...string) (string, error) {
	jsonText, err := getJsonFile(filename)
	if err != nil {
		//panic(err)
		return "", err
	}
	jsonBody, err := getJsonBody(jsonText)
	if err != nil {
		//panic(err)
		return "", err
	}

	return getByKeys(jsonBody, keys)
}

func GetSpecJsonValue(jsonStr []byte, keys ...string) (string, error) {
	jsonBody, err := getJsonBody(jsonStr)
	if err != nil {
		return "", err
	}
	return getByKeys(jsonBody, keys)
}

func GetJsonValueInt(jsonStr []byte, keys ...string) (int, error) {
	jsonBody, err := getJsonBody(jsonStr)
	if err != nil {
		return 0, err
	}
	return getIntByKeys(jsonBody, keys)
}

func ShowParseJsonMap(jsonBodyMap map[string]interface {}) {
	for _, v := range jsonBodyMap {
		switch reflect.ValueOf(v).Kind() {
		case reflect.Map:
			ShowParseJsonMap(v.(map[string]interface {}))
		case reflect.String:
			fmt.Printf("%v\n", v)
		}
	}
}




// ###
func getJsonFile(filename string) ([]byte, error) {
	absFilePath, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(absFilePath)
}
// ###
func getJsonBody(jsonText []byte) (*simplejson.Json, error) {
	return simplejson.NewJson(jsonText)
}
// ###
func getByKeys(jsonBody *simplejson.Json, keys []string) (string, error) {
	for _, v := range keys {
		jsonBody = jsonBody.Get(string(v))
	}

	return jsonBody.String()
}
// ###
func getIntByKeys(jsonBody *simplejson.Json, keys []string) (int, error) {
	for _, v := range keys {
		jsonBody = jsonBody.Get(string(v))
	}

	return jsonBody.Int()
}
