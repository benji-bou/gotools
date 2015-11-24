package tools

import (
	"errors"
	"runtime"
	// "log"
	"reflect"
	"strings"
)

func GetInnerTypeName(elem interface{}) string {
	elemValue := reflect.ValueOf(elem)
	if elemValue.Kind() == reflect.Ptr {
		elemValue = reflect.ValueOf(elemValue.Elem().Interface())
	}
	if elemValue.Kind() == reflect.Slice {
		return elemValue.Type().Elem().Name()
	}
	return elemValue.Type().Name()
}

// Dereference control if the interface is a pointer, if yes dereference the pointer and return the interface of the value
func Dereference(ptr interface{}) interface{} {
	if t := reflect.TypeOf(ptr); t.Kind() == reflect.Ptr {
		ptr = reflect.ValueOf(ptr).Elem().Interface()
	}
	return ptr
}

// CreatePtrToSliceFromInterface  create a ptr to slice from the input interface Use for MGO library All() function
// if the input is a Slice create ptr to slice and copy the input to the output
// else create a ptr to slice of type reflect.TypeOf(Interface)
//
func CreatePtrToSliceFromInterface(from interface{}) interface{} {
	var v reflect.Value
	from = Dereference(from)
	if t := reflect.TypeOf(from); t.Kind() == reflect.Slice {
		v = reflect.New(reflect.TypeOf(from))
		v.Elem().Set(reflect.ValueOf(from))
	} else {
		//Create an empty slice of type of resultInterface  by reflection
		vTmp := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(from)), 0, 20)
		//Create a ptr to Slice of type of resultInterface
		v = reflect.New(vTmp.Type())
		//Copy empty slice values to ptr slice -- This is done to produce an empty slice over a null value when marshalling the slice
		v.Elem().Set(vTmp)
	}
	//
	return v.Interface()
}

func getNameFromJSONTag(structType reflect.StructField) string {
	jsonTag := strings.Split(structType.Tag.Get("json"), ",")
	key := structType.Name
	if len(jsonTag) > 0 {
		key = jsonTag[0]
	}
	return key
}

func Map(elem interface{}) (map[string]interface{}, error) {
	elemValue := reflect.Indirect(reflect.ValueOf(elem))
	if elemValue.Kind() != reflect.Struct {
		return nil, errors.New("Unable top retrieve a correct structure type from argument")
	}
	result := make(map[string]interface{})
	elemValueType := elemValue.Type()
	for i := 0; i < elemValue.NumField(); i++ {
		fieldValue := elemValue.Field(i)
		structType := elemValueType.Field(i)
		if fieldValue.Kind() == reflect.Struct && NotEmpty(fieldValue.Interface()) {
			recursiveResp, _ := Map(fieldValue.Interface())
			result[getNameFromJSONTag(structType)] = recursiveResp
		} else if fieldValue.CanInterface() && fieldValue.IsValid() && NotEmpty(fieldValue.Interface()) {
			result[getNameFromJSONTag(structType)] = fieldValue.Interface()
		}
	}
	return result, nil
}

func NotEmpty(x interface{}) bool {
	return x != nil && !reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func Zero(x interface{}) interface{} {
	resultType := reflect.TypeOf(x)
	return reflect.Zero(resultType).Interface()
}

func SortArrayByType(input []interface{}) map[string][]interface{} {
	res := make(map[string][]interface{})
	for i, interface_in := range input {
		// if val := reflect.ValueOf(i); val.Kind() == reflect.Ptr {
		// 	interface_in = val.Elem().Interface()
		// }
		name := GetInnerTypeName(interface_in)
		value, ok := res[name]
		if ok == false {
			value = make([]interface{}, 0)
			res[name] = value
		}
		value = append(value, input[i])
		// Range copy the value it need to be ra assigned
		res[name] = value
	}
	return res
}

func ImplementsMethod(input interface{}, methodName string) bool {
	return NotEmpty(reflect.ValueOf(input).MethodByName(methodName))
}

func GetMethod(input interface{}, methodName string) (bool, reflect.Value) {
	v := reflect.ValueOf(input).MethodByName(methodName)
	if !NotEmpty(v) {
		return false, reflect.Value{}
	}
	return true, v
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
