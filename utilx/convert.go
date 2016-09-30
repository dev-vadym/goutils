package utilx

import (
	"reflect"
	"strconv"
	"math"
	"time"
	"fmt"
	"strings"
)


func Int(value string, defval int) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		num = defval
	}
	return num
}

func IntArray(defval int, values ...string) []int {
	arr := make([]int, 0)
	for _, v := range values{
		arr = append(arr, Int(v, defval))
	}
	return arr
}

func IntVal(value interface{}, defval int) int {
	iv, ok := value.(int)
	if !ok{
		iv = defval
	}
	return iv
}

func Int64(value string, defval int64) int64 {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defval
	}
	return num
}

func Int32(value string, defval int32) int32 {
	num, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defval
	}
	return int32(num)
}

func Float32(value string, defval float32) float32 {
	num, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defval
	}
	return float32(num)
}

func Float64(value string, defval float64) float64 {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defval
	}
	return num
}

func HumanNumber(value int, decimal int) string {
	return HumanFloat(float64(value), decimal)
}


func HumanFloat(value float64, decimal int) string {
	unit := ""
	if value >= 10000{
		value = float64(value) / 10000
		unit = "w"
	}else if (value >= 1000){
		value = float64(value) / 1000
		unit = "k"
	}
	strval := fmt.Sprintf("%.*f",  decimal, value)
	//去掉后面的0,例如：5.00变成5
	strval = strings.TrimRight(strings.TrimRight(strval, "0"), ".")
	return fmt.Sprintf("%v%v", strval, unit)
}

func Round(val float64) int {
	return int(Floor(val + 0.5))
}

func Floor(val float64) int {
	return int(math.Floor(val))
}

func Ceil(val float64) int {
	return int(math.Ceil(val + 0.5))
}


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0 // return correctly abs(-0)
	}
	return x
}

//uinx time
func Unix(t time.Time) int64 {
	unix := t.Unix()
	if unix < 0{
		unix = 0
	}
	return unix
}

//BindValues 绑定值
func BindValues(obj *interface{}, data map[string]interface{})  {
	t := reflect.TypeOf(obj)
    for i := 0; i < t.NumField(); i++ {
		objKey := t.Field(i).Name
		objVal := reflect.ValueOf(obj)
        if val, ok := data[objKey]; ok && val != nil{
			objVal.Set(reflect.ValueOf(val))
		}
    }
}

func IsNumber(value string) bool {
	if _, err := strconv.Atoi(value); err == nil {
		return true
	}else{
		return false
	}
}