package influxe_store

import (
	"context"
	"fmt"
	"log"
	"time"
)

func create_bucket(org, bucket string, client any, write any) {
	fmt.Println("创建buckets")
	fmt.Println(bucket)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"fields1": value,
		}
		point := write.NewPoint("measeurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println(point)
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}

	}

}
