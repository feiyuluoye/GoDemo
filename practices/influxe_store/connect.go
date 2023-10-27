package influxe_store

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	config "godemo.com/config"

	"os"
)

func ConnectMain() {
	token := os.Getenv(config.INFLUXDB_TOKEN)
	url := "http://172.16.30.139:8086"
	client := influxdb2.NewClient(url, token)
	fmt.Println(client)
	org := "wudun"
	bucket := "testing"
	create_bucket(org, bucket, client, write)

}
