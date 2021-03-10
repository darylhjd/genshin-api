/*
API at : https://api.genshin.dev/
GitHub Repo : https://github.com/genshindev/api
*/
package genshinapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	BaseAPI = "https://api.genshin.dev/"
)

// APIError : Struct to check for API errors.
// This struct is used in the GetCustomBody API function.
type APIError struct {
	Error           *string  `json:"error"`                     // If API encounters an error, Error will not be nil.
	AvailableTypes  []string `json:"availableTypes,omitempty"`  // For wrong data type
	AvailableImages []string `json:"availableImages,omitempty"` // For wrong image type
}

// GetCustomBody : Function to get data from API. Provide the function with a slice of URL-part strings.
// The function will concatenate the strings with backslashes so you do not need to do so.
func GetCustomBody(ext ...string) ([]byte, error) {
	reqUrl := fmt.Sprintf("%s%s", BaseAPI, strings.Join(ext, "/"))
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("request returned non-200 status code")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Read body into byte array
	byteArr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check error. If there is an error, then the Error field will not be empty.
	var apierror APIError
	err = json.Unmarshal(byteArr, &apierror)
	if apierror.Error != nil {
		return nil, errors.New(*apierror.Error)
	}

	return byteArr, nil
}
