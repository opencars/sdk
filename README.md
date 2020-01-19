# OpenCars Toolkit

[godoc]: https://godoc.org/github.com/opencars/toolkit
[godoc-img]: https://godoc.org/github.com/opencars/toolkit?status.svg
[goreport]: https://goreportcard.com/report/github.com/opencars/toolkit
[goreport-img]: https://goreportcard.com/badge/github.com/opencars/toolkit
[version]: https://img.shields.io/github/v/tag/opencars/toolkit?sort=semver

[![Docs][godoc-img]][godoc]
[![Go Report][goreport-img]][goreport]
[![Version][version]][version]

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
    "log"
    "os"

    "github.com/opencars/toolkit"
)

func main() {
    client := toolkit.New("https://api.opencars.app", os.Getenv("OPENCARS_API_KEY"))

    operations, err := client.Operation().FindByNumber("АА9359РС")
    if err != nil {
        log.Fatal(err)
    }

    for _, op := range operations {
        fmt.Println(op)
    }
```

## License

Project released under the terms of the MIT [license](./LICENSE).
