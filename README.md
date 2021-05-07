# genshin-api

[![Go Reference](https://pkg.go.dev/badge/github.com/darylhjd/genshin-api.svg)](https://pkg.go.dev/github.com/darylhjd/genshin-api)

Golang wrapper for a fan-made Genshin Impact API.

To install, do `go get -u github.com/darylhjd/genshin-api`.

API can be found [here](https://api.genshin.dev/). GitHub repo for API is [here](https://github.com/genshindev/api).

## Usage

```golang
package main

import (
	"fmt"
	
	genshinapi "github.com/darylhjd/genshin-api"
)

func main() {
	// Get list of data types provided by API: artifacts, characters, etc...
	dataTypes, _ := genshinapi.GetDataTypes()
	fmt.Println(dataTypes)

	// Get list of character names
	characters, _ := genshinapi.GetCharacters()
	fmt.Println(characters)
}
```

Refer to each DataType's corresponding file to see what helper functions are available. As a general rule, most
DataTypes support getting the full list of item names as well as a particular DataEntry of that type.

This library provides a wildcard function, `GetCustomBody`, for your own requests. You can find this function
in `api.go`.

```golang
package main

import (
	"encoding/json"
	"fmt"
	
	genshinapi "github.com/darylhjd/genshin-api"
)

func main() {
	// Get data on characters/amber
	urlArgs := []string{
		// Package provides a few helper strings
		genshinapi.CharactersDType, 
		"amber",
	}
	// Function will parse URL for you
	result, _ := genshinapi.GetCustomBody(urlArgs...)

	var character genshinapi.Character
	_ = json.Unmarshal(result, &character)
	fmt.Println(character)
	
	// Get portrait image of Amber
	// The function GetImage in types.go also provides this functionality
	urlArgs = append(urlArgs, genshinapi.CharacterPortrait)
	imagedata, _ := genshinapi.GetCustomBody(urlArgs...)
	_ = imagedata
}
```

## Contributing

Feel free to contribute to this project!