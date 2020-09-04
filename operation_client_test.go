package toolkit_test

import (
	"fmt"
	"log"

	"github.com/opencars/toolkit"
)

func ExampleOperationClient_FindByNumber() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	operations, err := client.Operation().FindByNumber("AA9359PC")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(operations[0].Number)
	// Output: АА9359РС
}
