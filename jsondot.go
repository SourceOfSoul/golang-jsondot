Copyright (c) 2017 The JsonDot Authors. All Rights Reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

package jsondot

import "encoding/json"

type JsonDot struct {
	JsonObject interface{}
}

func NewDot(byteData []byte) (*JsonDot, error) {
	dataMap := make(map[string]interface{}, 0)
	err := json.Unmarshal(byteData, &dataMap)
	if err != nil {
		return nil, err
	}
	dot := new(JsonDot)
	dot.JsonObject = dataMap
	return dot, nil
}

func (q *JsonDot) GetBool(key string) bool {
	return q.JsonObject.(map[string]interface{})[key].(bool)
}
func (q *JsonDot) GetNumber(key string) float64 {
	return q.JsonObject.(map[string]interface{})[key].(float64)
}
func (q *JsonDot) GetString(key string) string {
	return q.JsonObject.(map[string]interface{})[key].(string)
}
func (q *JsonDot) GetJson(key string) JsonDot {
	return JsonDot{q.JsonObject.(map[string]interface{})[key]}
}

func (q *JsonDot) GetBoolArray(key string) []bool {
	val := q.JsonObject.(map[string]interface{})[key]
	val_array := val.([]interface{})
	array := make([]bool, len(val_array))
	for i, j := range val_array {
		array[i] = j.(bool)
	}
	return array
}
func (q *JsonDot) GetNumberArray(key string) []float64 {
	val := q.JsonObject.(map[string]interface{})[key]
	val_array := val.([]interface{})
	array := make([]float64, len(val_array))
	for i, j := range val_array {
		array[i] = j.(float64)
	}
	return array
}
func (q *JsonDot) GetStringArray(key string) []string {
	val := q.JsonObject.(map[string]interface{})[key]
	val_array := val.([]interface{})
	array := make([]string, len(val_array))
	for i, j := range val_array {
		array[i] = j.(string)
	}
	return array
}
func (q *JsonDot) GetJsonArray(key string) []JsonDot {
	obj_array := q.JsonObject.(map[string]interface{})[key]
	array := obj_array.([]interface{})
	qArray := make([]JsonDot, len(array))
	for i, j := range array {
		qArray[i] = JsonDot{j}
	}
	return qArray
}

func (q *JsonDot) SetBool(key string, value bool) {
	q.JsonObject.(map[string]interface{})[key] = value
}
func (q *JsonDot) SetNumber(key string, value float64) {
	q.JsonObject.(map[string]interface{})[key] = value
}
func (q *JsonDot) SetString(key, value string) {
	q.JsonObject.(map[string]interface{})[key] = value
}
func (q *JsonDot) SetJson(key string, json *JsonDot) {
	q.JsonObject.(map[string]interface{})[key] = json
}

func (q *JsonDot) SetBoolArray(key string, values []bool) {
	obj_array := q.JsonObject.(map[string]interface{})[key]
	array := obj_array.([]interface{})
	for i, v := range values {
		array[i] = interface{}(v)
	}
}

func (q *JsonDot) SetStringArray(key string, values []string) {
	obj_array := q.JsonObject.(map[string]interface{})[key]
	array := obj_array.([]interface{})
	for i, v := range values {
		array[i] = interface{}(v)
	}
}

func (q *JsonDot) SetNumberArray(key string, values []float64) {
	obj_array := q.JsonObject.(map[string]interface{})[key]
	array := obj_array.([]interface{})
	for i, v := range values {
		array[i] = interface{}(v)
	}
}
