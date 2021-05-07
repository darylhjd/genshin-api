package genshinapi

import "encoding/json"

/*
NOTE: For specific data types and their helper functions, please refer to their corresponding file names.
*/

// DataEntry : This interface is implemented by each particular DataType provided by the API.
type DataEntry interface {
	EntryName() string
}

// getDataAndUnmarshal : Get required data from API and unmarshal to required struct.
// The function accepts a pointer to a required data type.
func getDataAndUnmarshal(req []string, datptr interface{}) error {
	bytes, err := GetCustomBody(req...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, datptr)
	return err
}

// DataTypes : Struct to hold list of data types provided by API.
// This is a struct because the API returns a JSON containing the list, rather than just a list.
type DataTypes struct {
	Types []string `json:"types"`
}

// GetDataTypeItemsList : Get a list of item names for a particular data type.
func GetDataTypeItemsList(t string) ([]string, error) {
	resp, err := GetCustomBody(t)
	if err != nil {
		return nil, err
	}

	var list []string
	err = json.Unmarshal(resp, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GetDataTypes : Get data types provided by the API
func GetDataTypes() ([]string, error) {
	bytes, err := GetCustomBody()
	if err != nil {
		return nil, err
	}

	var dts DataTypes
	err = json.Unmarshal(bytes, &dts)
	if err != nil {
		return nil, err
	}
	return dts.Types, nil
}

// GetImage : Return a byte array of image data for a data type entry
// Returns error if request is not valid (validity of request is guaranteed).
func GetImage(dtype, name, imagetype string) ([]byte, error) {
	return GetCustomBody(dtype, name, imagetype)
}
