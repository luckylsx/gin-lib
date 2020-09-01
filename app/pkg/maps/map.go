// As a quick start:
// 		function Values():
// 		maps := map[interface{}]interface{}{"name":"go", "age": 14}
// 		list := Values(maps)
// 		// output:
// 		// [go 12]
//
// 		function Column():
// 		map2 := []map[interface{}]interface{}{
//			{"name": "go", "age": "18"},
//			{"name": "java", "age": "24"},
// 		}
// 		list := Column(maps, "name")
// 		// output:
// 		// [go java]
//
// 		function ColumnKey():
// 		map2 := []map[interface{}]interface{}{
//			{"name": "go", "age": "18"},
//			{"name": "java", "age": "24"},
// 		}
// 		list := ColumnKey(maps, "age", "name)
// 		// output:
// 		// map[go:18 java:24]
package maps

// Values return values of maps
func Values(maps map[interface{}]interface{}) []interface{} {
	var values []interface{}
	for _, value := range maps {
		values = append(values, value)
	}
	return values
}

// Column get the slice of specify key
func Column(maps []map[interface{}]interface{}, key interface{}) []interface{} {
	var lists []interface{}
	for _, value := range maps {
		if v, ok := value[key]; ok {
			lists = append(lists, v)
		}
	}
	return lists
}

// ColumnKey get the maps of specify key
func ColumnKey(maps []map[interface{}]interface{}, value interface{}, key interface{}) map[interface{}]interface{} {
	lists := make(map[interface{}]interface{})
	for index, mapValue := range maps {
		v, VOK := mapValue[value]
		k, KeyOk := mapValue[key]
		if VOK && KeyOk {
			lists[k] = v
		} else if VOK && (!KeyOk) {
			lists[index] = v
		}
	}
	return lists
}
