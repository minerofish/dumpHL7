package dumpHL7

import (
	"fmt"
	"reflect"
	"time"
)

const (
	Value string = "value"
	Key   string = "key"
	Slice string = "slice"
)

func dump(i interface{}, name string) {
	level := 0
	skipLine := make(map[int]bool)
	skipLine[0] = true
	fmt.Print(name)
	recursiveDump(i, "", &level, false, &skipLine)
}

func recursiveDump(i interface{}, name string, level *int, lastElement bool, skipLine *(map[int]bool)) {
	v := reflect.ValueOf(i)
	skip := *skipLine
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if name != Slice {
		printPaths(Key, name, v, level, lastElement, &skip)
	}
	*level++

	switch v.Kind() {
	case reflect.Interface:
		fmt.Printf("Type of interface: %s\n", v.Kind().String())
	case reflect.Slice, reflect.Array:
		*level--
		for j := 0; j < v.Len(); j++ {
			recursiveDump(v.Index(j).Interface(), Slice, level, lastElement, &skip)
		}
	case reflect.Struct:
		if !lastElement && *level > 1 {
			skip[*level-1] = false
		}
		for j := 0; j < v.NumField(); j++ {
			if v.Type() == reflect.TypeOf(time.Time{}) {
				// time is a 3x struct (wall, ext, *loc), so have to bodge :/
				j += 2
				*level--
				printPaths(Value, "", v, level, lastElement, &skip)
				*level++
				continue
			}
			recursiveDump(v.Field(j).Interface(), v.Type().Field(j).Name, level, j == v.NumField()-1, &skip)
		}
		*level--
	default:
		*level--
		printPaths(Value, "", v, level, lastElement, &skip)
	}
}

func printPaths(toPrint string, name string, v reflect.Value, level *int, lastElement bool, skipLine *map[int]bool) {
	road := "│"
	junction := ""
	skip := *skipLine

	for j := 0; j < *level; j++ {
		if skip[j] {
			road = " "
		}
		fmt.Print(road, "  ")
		road = "│"
	}

	if toPrint == Key {
		if *level > 0 {
			junction = "┝"
		}
		if lastElement {
			junction = "└"
			skip[*level] = true
		}
		fmt.Print(junction, " ")
		fmt.Printf("%s (Type: %s, Kind: %s) \n", name, v.Type().String(), v.Kind().String())
		return
	}

	if toPrint == Value {
		if lastElement {
			road = " "
		}
		fmt.Printf(road+"   Value: %v\n", v.Interface())
	}
}
