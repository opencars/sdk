package toolkit_test

import (
	"fmt"
	"log"

	"github.com/opencars/toolkit"
)

func ExampleALPRClient_Recognize() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	results, err := client.ALPR().Recognize("https://raw.githubusercontent.com/opencars/alpr/master/pkg/api/http/test/example.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(results[0].Plate)
	// Output: AA9008MT
}
