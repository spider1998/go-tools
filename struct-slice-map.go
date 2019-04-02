package routing

import (
	"reflect"
	"strconv"
)

type sliceType int8

const (
	sliceInt sliceType = iota
	sliceUInt
	sliceString
)

// Map 可以提取一个结构体切片某个 int 或 string 字段。
// a := []struct{Foo int}{{1}, {2}}
// Map(a, "Foo") -> []int{1, 2}
func MapByField(a interface{}, field string) interface{} {
	T := reflect.TypeOf(a)
	if T.Kind() != reflect.Slice {
		panic("only slice allowed in map")
	}
	V := reflect.ValueOf(a)
	elemT := T.Elem()
	if elemT.Kind() == reflect.Ptr {
		elemT = elemT.Elem()
	}
	elem := reflect.Zero(elemT)
	if elem.Kind() != reflect.Struct {
		panic("only slice of struct type allowed")
	}
	kind := elem.FieldByName(field).Kind()

	var (
		ret interface{}
		t   sliceType
	)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = make([]int, V.Len())
		t = sliceInt
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret = make([]int, V.Len())
		t = sliceUInt
	case reflect.String:
		ret = make([]string, V.Len())
		t = sliceString
	default:
		panic("invalid type " + kind.String())
	}

	for i := 0; i < V.Len(); i++ {
		elem := reflect.Indirect(V.Index(i)).FieldByName(field)
		if elem.Kind() != kind {
			panic(kind.String() + " and " + elem.Kind().String() + " is not same")
		}
		switch t {
		case sliceInt:
			ret.([]int)[i] = int(elem.Int())
		case sliceUInt:
			ret.([]int)[i] = int(elem.Uint())
		case sliceString:
			ret.([]string)[i] = elem.String()
		}
	}
	return ret
}

// SliceIntToSliceString convert slice []int to []string
func SliceIntToSliceString(from []int) (to []string) {
	for _, i := range from {
		to = append(to, strconv.Itoa(i))
	}
	return
}

// SliceIntToSliceString convert slice []int to []string
func SliceStringToSliceInt(from []string) (to []int) {
	for _, i := range from {
		str, err := strconv.Atoi(i)
		if err == nil {
			to = append(to, str)
		}
	}
	return
}

// SliceAnyToSliceInterface convert slice []T to []interface{}
func SliceAnyToSliceInterface(from interface{}) (to []interface{}) {
	switch reflect.TypeOf(from).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(from)
		to = make([]interface{}, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			to = append(to, s.Index(i).Interface())
		}
	default:
		panic("only slice allowed to be converted")
	}
	return
}