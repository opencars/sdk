package toolkit_test

import (
	"fmt"
	"log"

	"github.com/opencars/toolkit"
)

func ExampleOperationClient_FindByNumber() {
	client := toolkit.New("https://api.opencars.app", "RMk86CHp8ulX3EnmVTJwffIYu6uAg3Wj")
	operations, err := client.Operation().FindByNumber("АА9359РС")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(operations[0].Number)
	// Output: АА9359РС
}
