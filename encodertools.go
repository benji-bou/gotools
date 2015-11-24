package tools

//JSONResp write the interface into a Json Format  it uses json.Marshal if interface get a type name, the json as a parent object with this name
import "encoding/json"

func JSONResp(rep interface{}) (string, error) {

	name := GetInnerTypeName(rep)
	jsons, err := json.Marshal(rep)
	jsonName, _ := json.Marshal(name)
	if len(jsonName) > 2 {
		return "{ " + string(jsonName) + " : " + string(jsons) + "}", err
	}
	return string(jsons), err
}

//XMLResp write the interface into a XML Format the XML as a parent object with this name
func XMLResp(rep interface{}) (string, error) {
	return "", nil
}
