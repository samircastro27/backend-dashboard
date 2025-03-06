package httprouter

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Request interface {
	GetData() []byte
	GetQueryString() string
	GetApiKey() (string, error)
	Parse(target interface{}) error
	QueryString(target interface{}) error
}

type request struct {
	data        []byte
	queryString string
	method      string
}

func (req *request) GetApiKey() (string, error) {
	values, err := url.ParseQuery(req.queryString)
	if err != nil {
		return "", err
	}

	queryMap := make(map[string]string)
	for key, value := range values {
		if len(value) > 0 {
			queryMap[key] = value[0]
		}
	}

	return queryMap["apiKey"], nil
}

func (req *request) GetQueryString() string {
	return req.queryString
}

func (req *request) GetData() []byte {
	return req.data
}

func (req *request) Parse(target interface{}) error {

	if target != nil {
		return json.Unmarshal(req.data, target)
	}

	return errors.New("target is nil")
}

func (req *request) QueryString(target interface{}) error {
	values, err := url.ParseQuery(req.queryString)
	if err != nil {
		return err
	}

	queryMap := make(map[string]string)
	for key, value := range values {
		if len(value) > 0 {
			queryMap[key] = value[0]
		}
	}

	targetValue := reflect.ValueOf(target).Elem()

	for key, value := range queryMap {
		fieldName := strings.Title(key)

		field := targetValue.FieldByName(fieldName)

		if field.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Int, reflect.Int64:
				if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					field.SetInt(intValue)
				}
			case reflect.Float64:
				if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
					field.SetFloat(floatValue)
				}
			case reflect.Bool:
				if boolValue, err := strconv.ParseBool(value); err == nil {
					field.SetBool(boolValue)
				}
			}
		}
	}

	return nil
}
