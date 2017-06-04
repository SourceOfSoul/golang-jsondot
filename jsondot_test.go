package jsondot

import (
	"testing"
)

const JSON_TEXT = `{
	"name": "Bill",
	"age": 18,
	"hobby": ["sport", "music", "game"],
	"lucknum": [1,3,7,12],
	"card": {
		"tel": 13311223344,
		"bank": [
			{
				"bankname": "ICBC",
				"cardnum": 3568111100002231987,
				"credit": 100000
			},
			{
				"bankname": "CCB",
				"cardnum": 8782266352638784672,
				"credit": 50000
			}
		]
	}
}`

func initDot(t *testing.T) *JsonDot {
	byteData := []byte(JSON_TEXT)
	if byteData == nil {
		t.Error("byteData is NIL")
		panic("convert string to []byte error!")
	}
	dot, err := NewDot(byteData)
	if err != nil {
		t.Error(err)
		panic("NewDot error !")
	}
	return dot
}

func Test_NewDot(t *testing.T) {
	dot := initDot(t)
	if dot != nil {
		t.Log("Test_NewDot success!")
	} else {
		t.Log("Test_NewDot error!")
	}
}

func Test_GetValue(t *testing.T) {
	dot := initDot(t)

	t.Log(dot.GetString("name"))
	t.Log(dot.GetNumber("age"))

	t.Log("Test_GetValue success!")
}

func Test_SetValue(t *testing.T) {
	dot := initDot(t)

	t.Log("name=>", dot.GetString("name"))
	dot.SetString("name", "Frank")
	t.Log("name=>", dot.GetString("name"))

	t.Log("age=>", dot.GetNumber("age"))
	dot.SetNumber("age", 99)
	t.Log("age=>", dot.GetNumber("age"))

	t.Log("Test_SetValue success!")
}

func Test_GetStringArray(t *testing.T) {
	dot := initDot(t)

	t.Log("hobby=>", dot.GetStringArray("hobby"))

	slice := dot.GetStringArray("hobby")

	t.Log("[0]=>", slice[0])

	t.Log("[2]=>", slice[2])

	t.Log("Test_GetStringArray success!")
}

func Test_GetIntArray(t *testing.T) {
	dot := initDot(t)

	t.Log("lucknum=>", dot.GetNumberArray("lucknum"))

	slice := dot.GetNumberArray("lucknum")

	t.Log("[0]=>", slice[0])

	t.Log("[2]=>", slice[2])

	t.Log("Test_GetIntArray success!")
}

func Test_SetStringArray(t *testing.T) {
	dot := initDot(t)

	t.Log("before Change: hobby=>", dot.GetStringArray("hobby"))

	slice := dot.GetStringArray("hobby")

	t.Log("[0]=>", slice[0])

	slice[0] = "swimming"

	t.Log("[0]=>", slice[0])

	t.Log("after reassign: hobby=>", dot.GetStringArray("hobby"))

	dot.SetStringArray("hobby", slice)

	t.Log("after Set: hobby=>", dot.GetStringArray("hobby"))

	t.Log("Test_SetStringArray success!")
}

func Test_SetNumberArray(t *testing.T) {
	dot := initDot(t)

	t.Log("lucknum=>", dot.GetNumberArray("lucknum"))

	slice := dot.GetNumberArray("lucknum")

	t.Log("[0]=>", slice[0])

	slice[0] = 123

	t.Log("[0]=>", slice[0])

	t.Log("after reassign: lucknum=>", dot.GetNumberArray("lucknum"))

	dot.SetNumberArray("lucknum", slice)

	t.Log("after Set: lucknum=>", dot.GetNumberArray("lucknum"))

	t.Log("Test_SetStringArray success!")
}
