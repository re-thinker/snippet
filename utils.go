package snippet

import "strings"

type MapStringInterface map[string]interface{}

// StringToMap 转换字符串到map结构
// 把XXX1=YYY;XX2=ZZZ 转换成map
func StringToMap(params string) MapStringInterface {
	m := MapStringInterface{}
	mapStrings := strings.Split(params, ";")
	for _, mapWord := range mapStrings {
		words := strings.Split(mapWord, "=")
		m[words[0]] = words[1]
	}
	return m
}
