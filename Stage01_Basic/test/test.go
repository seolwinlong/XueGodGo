package test

var Str1 = "我可以被外部调用"
var str2 = "我不可以被外部调用"

func Test() string {
	return str2
}
