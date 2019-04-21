# Toolkit

## Overview

Download opencars toolkit with following command

```
$ go get github.com/opencars/toolkit
```

## Development

Simple example of toolkit usage

```go
package main

import (
	"fmt"

	"github.com/opencars/toolkit"
)

func main() {
  client := toolkit.New("https://api.opencars.pp.ua")
  transport, err := client.Search("AA1111BX")

  if err != nil {
  	panic(err)
  }

  fmt.Println(transport)
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).
