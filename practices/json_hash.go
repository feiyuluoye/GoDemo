package practices

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func JsonHashDemo() {
	type Customer struct {
		Name    string    `json:"name"`
		Created time.Time `json:"created"`
	}

	data := []byte(`
    {
        "name": "J.B",
        "created": "2018-04-09T23:00:00Z"
    }`)

	var c Customer
	if err := json.Unmarshal(data, &c); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(c.Name)
}
