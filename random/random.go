package random

import (
	"strings"
	"time"
	"math/rand"
	//"log"
	"github.com/foolin/goutils/uuid"
)

var ran = rand.New(rand.NewSource(time.Now().UnixNano()))
var chars = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var alphabets = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}


//NewUid Unique Identifier
func Uid() string {
	objUuid := uuid.NewV4()
	return strings.Replace(objUuid.String(), "-", "", -1)
}


func TimeUid(isAfter ...bool) string {
	subTime := time.Now().Format("20060102150405")
	subUid := Uid()[:18]
	if len(isAfter) > 0  && isAfter[0]{
		return subUid + subTime
	}
	return subTime + subUid
}


//RandomNumber random number between min and max ([min, max]), include max
func Number(min, max int) int {
	if min > max{
		return -1
	}
	return ran.Intn(max - min + 1) + min
}

//RandomString random string
func String(length int) string {
	if length <= 0 {
		length = 32
	}
	result := make([]string, length)
	charsLen := len(chars)
	for i := 0; i < length; i++ {
		idx := ran.Intn(charsLen)
		result[i] = chars[idx]
	}
	return strings.Join(result, "")
}

func Alphabet(length int) string {
	if length <= 0 {
		length = 32
	}
	result := make([]string, length)
	charsLen := len(alphabets)
	for i := 0; i < length; i++ {
		idx := ran.Intn(charsLen)
		result[i] = alphabets[idx]
	}
	return strings.Join(result, "")
}

//Weight
func WeightIndex(weightArray []int) int {
	type Scope struct{
		Idx int
		Min int
		Max int
	}
	sum := 0	//总数
	min := 0	//上标
	idx := -1	//索引
	scopeList := make([]Scope, 0)
	for key, val := range weightArray{
		sum += val //下标
		item := Scope{key, min, sum}
		scopeList = append(scopeList, item)
		min = sum + 1
	}
	rdValue := Number(1, sum);
	for _, v := range scopeList{
		if rdValue >= v.Min && rdValue <= v.Max{
			idx = v.Idx
			break
		}
	}
	//log.Printf("list: %v, randomVal: %v, idx: %v", scopeList, rdValue, idx)
	return idx
}

//Weight
func WeightKey(weightMap map[interface{}]int) interface{} {
	keyList := make([]interface{}, 0)
	weightList := make([]int, 0)
	for key, val := range weightMap{
		keyList = append(keyList, key)
		weightList = append(weightList, val)
	}
	idx := WeightIndex(weightList)
	if idx < 0{
		return nil
	}
	return keyList[idx]
}

//Weight
func WeightKeys(weightMap map[interface{}]int, num int) []interface{} {
	retList := make([]interface{}, 0)
	if len(weightMap) <= num{
		for key, _ := range retList{
			retList = append(retList, key)
		}
		return retList
	}
	for i := 0; i < num; i++{
		retKey := WeightKey(weightMap)
		if retKey != nil{
			retList = append(retList, retKey)
			delete(weightMap, retKey)
		}
	}
	return retList
}

func UniqueArray(num, min, max int) []int {
	if num <= 0{
		return []int{}
	}
	sub := max - min + 1
	ret := make([]int, 0)
	m := make(map[int]bool, num)
	if num >= sub{
		num = sub
		for i := min; i <= max; i++{
			ret = append(ret, i)
		}
		//随机打乱顺序
		for i := 0; i < num * 2; i++{
			idx1 := Number(1, num) - 1
			idx2 := Number(1, num) - 1
			ret[idx1], ret[idx2] = ret[idx2], ret[idx1]
		}
		return ret
	}
	for {
		v := Number(min, max)
		if _, ok := m[v]; ok{
			continue
		}
		m[v] = true
		ret = append(ret, v)
		if len(ret) >= num{
			break
		}
	}
	return ret
}

func Array(num, min, max int) []int {
	ret := make([]int, 0)
	for {
		v := Number(min, max)
		ret = append(ret, v)
		if len(ret) >= num{
			break
		}
	}
	return ret
}

//是否命中概率，0-1之间的小数点
func Hit(rate float32) bool  {
	rateInt := int(rate * 100)
	return HitInt(rateInt)
}

// HitInt 是否命中概率，1-100之间的小数点
func HitInt(rateInt int) bool  {
	if rateInt < 1{
		return false
	}
	if rateInt >= 100{
		return true
	}
	isHit := false
	rdVal := Number(1, 100)
	if rdVal <= rateInt{
		isHit = true
	}
	return isHit
}