package utilx

import (
	"fmt"
	"strings"
)

func MbStrLen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}


func SplitInt(val string, sep string) []int {
	val = strings.TrimSpace(val)
	list := make([]int, 0)
	if val != ""{
		arr := strings.Split(val, sep)
		for _, v := range arr{
			iv := Int(v, 0)
			list = append(list, iv)
		}
	}
	return list
}

func JoinInt(sep string, list ...int) string {
	strs := make([]string, 0)
	for _, v := range list{
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	return strings.Join(strs, sep)
}

func JoinInt64(sep string, list ...int64) string {
	strs := make([]string, 0)
	for _, v := range list{
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	return strings.Join(strs, sep)
}

func SplitInt64(val string, sep string) []int64 {
	val = strings.TrimSpace(val)
	list := make([]int64, 0)
	if val != ""{
		arr := strings.Split(val, sep)
		for _, v := range arr{
			iv := Int64(v, 0)
			list = append(list, iv)
		}
	}
	return list
}

func Join(sep string, list ...interface{}) string {
	strs := make([]string, 0)
	for _, v := range list{
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	return strings.Join(strs, sep)
}



//
//func MbStrSub(s string, l int) string {
//	if len(s) <= l {
//		return s
//	}
//	ss, sl, rl, rs := "", 0, 0, []rune(s)
//	for _, r := range rs {
//		rint := int(r)
//		if rint < 128 {
//			rl = 1
//		} else {
//			rl = 2
//		}
//
//		if sl + rl > l {
//			break
//		}
//		sl += rl
//		ss += string(r)
//	}
//	return ss
//}