package reflects

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// GetStructFieldTags get struct field tags
//
// # Examples
// ```
// type User struct {
//     ID   string `db:"id"`
//     Name string `db:"name"`
// }
//
// GetStructFieldTags(User{ID: "1", Name: "jason"}, "db")
// GetStructFieldTags(&User{ID: "1", Name: "jason"}, "db")
// ```
func GetStructFieldTags(s interface{}, key string) (tags []string, err error) {
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		return GetVStructFieldTags(s, key)
	}
	if reflect.TypeOf(s).Kind() == reflect.Ptr && reflect.TypeOf(s).Elem().Kind() == reflect.Struct {
		return GetRStructFieldTags(s, key)
	}
	err = fmt.Errorf("%#v is not a struct type", s)
	return
}

// GetVStructFieldTags get value struct filed tags
//
// # Examples
// ```
// type User struct {
//     ID   string `db:"id"`
//     Name string `db:"name"`
// }
//
// GetVStructFieldTags(User{ID: "1", Name: "jason"}, "db")
// ```
func GetVStructFieldTags(vs interface{}, key string) (tags []string, err error) {
	var t = reflect.TypeOf(vs)
	for i := 0; i < t.NumField(); i++ {
		if tag, ok := t.Field(i).Tag.Lookup(key); !ok {
			err = fmt.Errorf("tag name is not %s", key)
			return
		} else {
			tags = append(tags, tag)
		}
	}
	return
}

// GetRStructFieldTags get reference struct filed tags
//
// # Examples
// ```
// type User struct {
//     ID   string `db:"id"`
//     Name string `db:"name"`
// }
//
// GetRStructFieldTags(&User{ID: "1", Name: "jason"}, "db")
// ```
func GetRStructFieldTags(rs interface{}, key string) (tags []string, err error) {
	var v = reflect.ValueOf(rs).Elem()
	for i := 0; i < v.NumField(); i++ {
		if tag, ok := v.Type().Field(i).Tag.Lookup(key); !ok {
			err = fmt.Errorf("tag name is not %s", key)
			return
		} else {
			tags = append(tags, tag)
		}
	}
	return
}

// GetStructFieldValues 获取结构体字段的值
//
// # Examples
// ```
// type User struct {
//     ID   string
//     Name string
// }
//
// GetStructFieldValues(User{ID: "1", Name: "jason"})
// GetStructFieldValues(&User{ID: "1", Name: "jason"})
func GetStructFieldValues(s interface{}, key string) (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	if reflect.TypeOf(s).Kind() == reflect.Struct {
		return GetVStructFieldValues(s, key)
	}
	if reflect.TypeOf(s).Kind() == reflect.Ptr && reflect.TypeOf(s).Elem().Kind() == reflect.Struct {
		return GetRStructFieldValues(s, key)
	}
	err = fmt.Errorf("%#v is not a struct type", s)
	return
}

func GetVStructFieldValues(vs interface{}, key string) (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	var v = reflect.ValueOf(vs)
	for i := 0; i < v.NumField(); i++ {
		if tag, ok := v.Type().Field(i).Tag.Lookup(key); !ok {
			err = fmt.Errorf("tag name is not %s", key)
			return
		} else {
			values[tag] = v.Field(i).Interface()
		}
	}
	return
}

func GetRStructFieldValues(rs interface{}, key string) (values map[string]interface{}, err error) {
	values = make(map[string]interface{})
	var v = reflect.ValueOf(rs).Elem()
	for i := 0; i < v.NumField(); i++ {
		if tag, ok := v.Type().Field(i).Tag.Lookup(key); !ok {
			err = fmt.Errorf("tag name is not %s", key)
			return
		} else {
			values[tag] = v.Field(i).Interface()
		}
	}
	return
}

// TagDeepEqual compare whether the tags of two struct are equal
//
// # Examples
// ```
// type User struct {
//     UID   int    `db:"id"`
//     UName string `db:"name"`
// }
//
// type Person struct {
//     PID   int    `db:"ID"`
//     PName string `db:"Name"`
// }
//
// type Human struct {
//     HID   int    `db:"hid"`
//     HName string `db:"name"`
// }
//
// true, nil := TagDeepEqual(User{}, Person{}, "db")
// false, nil := TagDeepEqual(User{}, Human{}, "db")
// ```
func TagDeepEqual(s1, s2 interface{}, key string) (err error) {
	s1t, err := GetStructFieldTags(s1, key)
	if err != nil {
		return
	}
	s2t, err := GetStructFieldTags(s2, key)
	if err != nil {
		return
	}
	if !StringSliceEqualFold(s1t, s2t) {
		err = fmt.Errorf("tag for %#v and tag for %#v is different", s1, s2)
		return
	}
	return
}

// StringSliceEqualFold compare whether two slice elements of same length are consistent
func StringSliceEqualFold(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	sort.Strings(s1)
	sort.Strings(s2)
	for index, tag := range s1 {
		if !strings.EqualFold(tag, s2[index]) {
			return false
		}
	}
	return true
}
