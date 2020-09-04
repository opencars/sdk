package toolkit_test

import (
	"fmt"
	"log"

	"github.com/opencars/toolkit"
)

func ExampleWantedClient_FindByNumber() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	wanted, err := client.Wanted().FindByNumber("AT4750CP")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*wanted[0].Number)
	// Output: АТ4750СР
}

func ExampleWantedClient_FindByVIN() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	wanted, err := client.Wanted().FindByVIN("5YJSA1E28HF176944")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*wanted[0].BodyNumber)
	// Output: 5YJSA1E28HF176944
}
