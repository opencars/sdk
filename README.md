# Toolkit

## Overview

Download opencars toolkit with following command

```sh
go get -u github.com/opencars/toolkit
```

## Example

Simple example of toolkit usage

```go
package main

import (
  "fmt"

  "github.com/opencars/toolkit"
)

func main() {
  client := toolkit.NewSDK("https://api.opencars.pp.ua")
  operation, err := client.Operation("AA1111BX")

  if err != nil {
    panic(err)
  }

  fmt.Println(operation)
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).
