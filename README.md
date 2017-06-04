jsondot is utility to easy use json.

we assume have a json data like below:

`{
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


you can test jsondot like:

$go test -v -test.run Test_GetValue
=== RUN   Test_GetValue
--- PASS: Test_GetValue (0.00s)
        jsondot_test.go:55: Bill
        jsondot_test.go:56: 18
        jsondot_test.go:58: Test_GetValue success!
PASS
ok      jsondot 0.324s

$go test -v -test.run Test_SetValue
=== RUN   Test_SetValue
--- PASS: Test_SetValue (0.00s)
        jsondot_test.go:64: name=> Bill
        jsondot_test.go:66: name=> Frank
        jsondot_test.go:68: age=> 18
        jsondot_test.go:70: age=> 99
        jsondot_test.go:72: Test_SetValue success!
PASS
ok      jsondot 0.344s


$go test -v -test.run Test_GetIntArray
=== RUN   Test_GetIntArray
--- PASS: Test_GetIntArray (0.00s)
        jsondot_test.go:78: name=> [1 3 7 12]
        jsondot_test.go:82: [0]=> 1
        jsondot_test.go:84: [0]=> 7
        jsondot_test.go:86: Test_GetIntArray success!
PASS
ok      jsondot 0.297s


$go test -v -test.run Test_GetStringArray
=== RUN   Test_GetStringArray
--- PASS: Test_GetStringArray (0.00s)
        jsondot_test.go:78: hobby=> [sport music game]
        jsondot_test.go:82: [0]=> sport
        jsondot_test.go:84: [0]=> game
        jsondot_test.go:86: Test_GetStringArray success!
PASS
ok      jsondot 0.305s


$go test -v -test.run Test_SetStringArray
=== RUN   Test_SetStringArray
--- PASS: Test_SetStringArray (0.00s)
        jsondot_test.go:106: before Change: hobby=> [sport music game]
        jsondot_test.go:110: [0]=> sport
        jsondot_test.go:114: [0]=> swimming
        jsondot_test.go:116: after reassign: hobby=> [sport music game]
        jsondot_test.go:120: after Set: hobby=> [swimming music game]
        jsondot_test.go:122: Test_SetStringArray success!
PASS
ok      jsondot 0.285s


$go test -v -test.run Test_SetNumberArray
=== RUN   Test_SetNumberArray
--- PASS: Test_SetNumberArray (0.00s)
        jsondot_test.go:128: lucknum=> [1 3 7 12]
        jsondot_test.go:132: [0]=> 1
        jsondot_test.go:136: [0]=> 123
        jsondot_test.go:138: after reassign: lucknum=> [1 3 7 12]
        jsondot_test.go:142: after Set: lucknum=> [123 3 7 12]
        jsondot_test.go:144: Test_SetStringArray success!
PASS
ok      jsondot 0.375s


