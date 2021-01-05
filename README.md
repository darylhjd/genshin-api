# genshin-api

Go wrapper around a Genshin Impact API.

API can be found [here](https://api.genshin.dev/). GitHub repo for API is [here](https://github.com/genshindev/api).

## Example
```golang
package main

import (
	"fmt"
	genshinapi "github.com/darylhjd/genshin-api"
)

func main() {
   // Get data types provided by API: artifacts, characters, etc...
   dataTypes, _ := genshinapi.GetDataTypes()
   fmt.Println(dataTypes)
   
   // Get list of characters
   characters, _ := genshinapi.GetCharacters()
   fmt.Println(characters)
}
```
