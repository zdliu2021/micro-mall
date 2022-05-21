package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// 结构体赋值:v1=v2

func Assignment(v1 interface{}, v2 interface{}) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ok = false
		}
	}()
	val1 := reflect.Indirect(reflect.ValueOf(v1))
	val2 := reflect.Indirect(reflect.ValueOf(v2))
	typ1 := val1.Type()
	typ2 := val2.Type()

	allType := []reflect.Kind{reflect.Int64, reflect.Int, reflect.Int32, reflect.Int16, reflect.String, reflect.Float32, reflect.Float64, reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint16, reflect.Uint8}

	elementInSlice := func(e reflect.Kind) bool {
		for i := 0; i < len(allType); i++ {
			if e.String() == allType[i].String() {
				return true
			}
		}
		return false
	}

	for i := 0; i < val2.NumField(); i++ {
		if !elementInSlice(typ2.Field(i).Type.Kind()) {
			continue
		}
		if _, ok := typ1.FieldByName(typ2.Field(i).Name); ok {
			if val1.FieldByName(typ2.Field(i).Name).Kind() != val2.Field(i).Kind() {
				k1 := val1.FieldByName(typ2.Field(i).Name).Kind().String()
				k2 := val2.Field(i).Kind().String()

				if strings.Contains(k1, "int") && strings.Contains(k2, "int") {
					var tmp int64
					switch k2 {
					case "int32":
						tmp = int64(val2.Field(i).Interface().(int32))
					case "int64":
						tmp = int64(val2.Field(i).Interface().(int64))
					case "int":
						tmp = int64(val2.Field(i).Interface().(int))
					default:
						tmp = 0
					}
					switch k1 {
					case "int32":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(int32(tmp)))
					case "int":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(int(tmp)))
					case "int64":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(tmp))
					}
				}
			} else {
				val1.FieldByName(typ2.Field(i).Name).Set(val2.Field(i))
			}
		}
	}

	return true
}

// 结构体赋值:v1=v2
// 只有v2 的非零值字段才会赋值给v1

func NotZeroAssignment(v1 interface{}, v2 interface{}) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ok = false
		}
	}()
	val1 := reflect.Indirect(reflect.ValueOf(v1))
	val2 := reflect.Indirect(reflect.ValueOf(v2))
	typ1 := val1.Type()
	typ2 := val2.Type()

	for i := 0; i < val2.Type().NumField(); i++ {
		if _, ok := typ1.FieldByName(typ2.Field(i).Name); ok && !val2.Field(i).IsZero() {
			if val1.FieldByName(typ2.Field(i).Name).Kind() != val2.Field(i).Kind() {
				k1 := val1.FieldByName(typ2.Field(i).Name).Kind().String()
				k2 := val2.Field(i).Kind().String()

				if strings.Contains(k1, "int") && strings.Contains(k2, "int") {
					var tmp int64
					switch k2 {
					case "int32":
						tmp = int64(val2.Field(i).Interface().(int32))
					case "int64":
						tmp = int64(val2.Field(i).Interface().(int64))
					case "int":
						tmp = int64(val2.Field(i).Interface().(int))
					default:
						tmp = 0
					}
					switch k1 {
					case "int32":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(int32(tmp)))
					case "int":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(int(tmp)))
					case "int64":
						val1.FieldByName(typ2.Field(i).Name).Set(reflect.ValueOf(tmp))
					}
				}
			} else {
				val1.FieldByName(typ2.Field(i).Name).Set(val2.Field(i))
			}
		}
	}

	return true
}
