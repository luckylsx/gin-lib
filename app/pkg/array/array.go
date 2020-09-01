package array

import (
	"reflect"
)

// IsArray return true of the value if it is array or slice or not
func IsArray(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

// InArray return true or false if needle in interface list or not
func InArray(needle, v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice:
		lists := reflect.ValueOf(v)
		for i := 0; i < lists.Len(); i++ {
			if reflect.DeepEqual(needle, lists.Index(i).Interface()) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// StrInArray return string if string needle in array slice lists or not
func StrInArray(needle string, lists []string) bool {
	for _, v := range lists {
		if needle == v {
			return true
		}
	}
	return false
}

// IntInArray return a int if needle in array slice lists or not
func IntInArray(needle int, lists []int) bool {
	for _, v := range lists {
		if needle == v {
			return true
		}
	}
	return false
}

// Search return the index if the needle in haystack or return -1 if it not in haystack
func Search(needle, haystack interface{}) int {
	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Array, reflect.Slice:
		lists := reflect.ValueOf(haystack)
		for i := 0; i < lists.Len(); i++ {
			if reflect.DeepEqual(needle, lists.Index(i).Interface()) {
				return i
			}
		}
		return -1
	default:
		return -1
	}
}

// StrSearch return the index if string is in slice
func StrSearch(needle string, lists []string) int {
	for index, v := range lists {
		if needle == v {
			return index
		}
	}
	return -1
}

// IntSearch return the index if int is in slice
func IntSearch(needle int, lists []int) int {
	for index, v := range lists {
		if needle == v {
			return index
		}
	}
	return -1
}

// Sum return the value of sum with lists slice
func Sum(lists []int) int {
	var sum int
	for _, v := range lists {
		sum += v
	}
	return sum
}

// RemoveRepeatedInt delete the repeat int of arr slice
func RemoveRepeatedInt(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// RemoveRepeatedString delete the repeat string of arr slice
func RemoveRepeatedString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// IndexRemove return new slice after remove element by index
func IndexRemove(s []interface{}, i int) []interface{} {
	return append(s[:i], s[i+1:]...)
}

// ElemRemove return new slice after remove element by index
func ElemRemove(slice []interface{}, elem interface{}) []interface{} {
	if len(slice) == 0 {
		return slice
	}
	for i, v := range slice {
		if v == elem {
			slice = append(slice[:i], slice[i+1:]...)
			return ElemRemove(slice, elem)
		}
	}
	return slice
}

// Keys return keys of maps
func Keys(maps map[interface{}]interface{}) []interface{} {
	var keys []interface{}
	for key := range maps {
		keys = append(keys, key)
	}
	return keys
}

// Shift return the new slice after delete the element of header
func Shift(lists []interface{}) []interface{} {
	return lists[1:]
}

// Unshift return the new slice after the new element append header
func Unshift(lists []interface{}, elem interface{}) []interface{} {
	newList := make([]interface{}, len(lists)+1)
	newList[0] = elem
	return append(newList, lists...)
}

// Push return the new slice after the new element append header
func Push(lists []interface{}, elem interface{}) []interface{} {
	return append(lists, elem)
}

// Pop return the new slice after delete the element of tail
func Pop(lists []interface{}) []interface{} {
	return lists[:len(lists)-1]
}
