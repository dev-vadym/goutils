package random

import "testing"


func TestTimeUid(t *testing.T)  {
	t.Logf("-----------------\n")
	for i := 0; i < 100; i++{
		t.Log(TimeUid())
	}
	t.Logf("-----------------\n")
	for i := 0; i < 100; i++{
		t.Log(TimeUid(true))
	}
	t.Logf("-----------------\n")

	threads := 1000
	times := 10000
	c := make(chan []string)
	for n := 0; n < threads; n++{
		go func() {
			list := []string{}
			for i := 0; i < times; i++{
				uid := TimeUid()
				list = append(list, uid)
			}
			c <- list
		}()
	}
	uniMap := map[string]int{}
	for n := 0; n < threads; n++{
		list := <-c
		t.Logf("recieve %v count uids, first:%v", len(list), list[0])
		for _, v := range list{
			count := uniMap[v] + 1
			uniMap[v] = count
			if count > 1{
				t.Errorf("uid %v has % repeat %v times", v, count)
			}
		}
	}
	t.Logf("New %v threads, %v times done!!!", threads, times)
	t.Logf("-----------------\n")
}

//func TestRandom(t *testing.T)  {
//	t.Logf("-----------------\n")
//	min := 0
//	max := 0
//	num := Number(min, max)
//	t.Logf("min: %v, max: %v, num: %v", min, max, num)
//	t.Logf("-----------------\n")
//	var weightList = []struct{
//		x int
//		y int
//	}{
//		{1, 4},
//	}
//	weightList = append(weightList, struct{
//		x int
//		y int
//	}{1, 2})
//	t.Logf("list is %v", weightList)
//}

//func TestRandom(t *testing.T)  {
//	t.Logf("-----------------\n")
//	min,max := 10, 99
//	for i:=0; i < 1000; i++{
//		num := Number(min, max)
//		if num <= min || num >= max{
//			t.Logf("%v %v", num, String(30))
//		}
//		if num < min || num > max{
//			t.Fatalf("%v is not between %v and %v", num, min, max)
//		}
//	}
//	t.Logf("-----------------\n")
//}

//func TestWeightIndex(t *testing.T)  {
//	t.Logf("-----------------\n")
//	//arr := make([]int, 0, 4)
//	//arr = append(arr, 123)
//	arr := []int{1,2, 3}
//	mCount := make(map[int]int, 0)
//	n := 1000000
//	for i:=0; i < n; i++{
//		idx := WeightIndex(arr)
//		if idx < 0{
//			t.Errorf("idx is -1, arr is %v", arr)
//		}
//		val := arr[idx]
//		if num, ok := mCount[val]; ok{
//			mCount[val] = num + 1
//
//		}else{
//			mCount[val] = 1
//		}
//		//t.Logf("array %v random value is %v, index is :%v",  arr, arr[idx], idx)
//	}
//	t.Logf("%v times result is: %#v", n, mCount)
//	t.Logf("-----------------\n")
//}

//func TestWeightKey(t *testing.T)  {
//	t.Logf("-----------------\n")
//	//arr := make([]int, 0, 4)
//	//arr = append(arr, 123)
//	var x interface{}
//	x = nil
//	ix, _ := x.(int)
//	t.Logf("parser x to int is %v", ix)
//	mFoo := make(map[interface{}]int, 0)
//	mFoo["a"] = 5
//	mFoo["b"] = 15
//	mFoo["c"] = 30
//	mFoo["d"] = 50
//	mCount := make(map[interface{}]int, 0)
//	n := 1000000
//	for i:=0; i < n; i++{
//		key := WeightKey(mFoo)
//		if key == nil{
//			t.Fatalf("WeightKey() return nil")
//		}
//		if num, ok := mCount[key]; ok{
//			mCount[key] = num + 1
//
//		}else{
//			mCount[key] = 1
//		}
//		//t.Logf("array %v random value is %v, index is :%v",  arr, arr[idx], idx)
//	}
//	//t.Logf("%v times result is: %#v", n, mCount)
//	for k, v := range mCount{
//		t.Logf("key: %v, weight: %v, count: %v, percent: %.2f%%", k, mFoo[k], v, (float64(v) * 100 / float64(n)) )
//	}
//	t.Logf("-----------------\n")
//}

//
//func TestUniqueArray(t *testing.T)  {
//	t.Logf("-----------------\n")
//	min,max := 10, 16
//	for i:=0; i < 100; i++{
//		arr := UniqueArray(8, min, max)
//		t.Logf("%v is between %v and %v", arr, min, max)
//	}
//	t.Logf("-----------------\n")
//}
//
//
//func TestArray(t *testing.T)  {
//	t.Logf("-----------------\n")
//	min,max := 10, 16
//	for i:=0; i < 100; i++{
//		arr := Array(8, min, max)
//		t.Logf("%v is between %v and %v", arr, min, max)
//	}
//	t.Logf("-----------------\n")
//}