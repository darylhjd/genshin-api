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

// GetCustomBody : Function to get data from API
// ext: Slice of URL arguments for API.
// Function will concatenate with backslashes.
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

	// Check error
	var apierror APIError
	err = json.Unmarshal(byteArr, &apierror)
	if apierror.Error != nil {
		return nil, errors.New(*apierror.Error)
	}

	return byteArr, nil
}
