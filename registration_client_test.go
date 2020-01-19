package toolkit_test

import (
	"fmt"
	"log"

	"github.com/opencars/toolkit"
)

func ExampleRegistrationClient_FindByVIN() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	registrations, err := client.Registration().FindByVIN("5YJXCCE40GF010543")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*registrations[0].VIN)
	// Output: 5YJXCCE40GF010543
}

func ExampleRegistrationClient_FindByNumber() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	registrations, err := client.Registration().FindByNumber("AA9359PC")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(registrations[0].NRegNew)
	// Output: AA9359PC
}

func ExampleRegistrationClient_FindByCode() {
	client := toolkit.New("https://api.opencars.app", "jIidbA8wivROjFNv1H8SiEoFQFHZ0VzL")
	registration, err := client.Registration().FindByCode("CXH484154")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(registration.SDoc + registration.NDoc)
	// Output: CXH484154
}
