package test

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"strconv"
	"time"
)
type UserSetup struct{
	Url    string
	Org    string
	Bucket string
	Token    string
}
func Write(userSetup UserSetup) {
	//GO操作influxdb_go influxdb
	//https://blog.csdn.net/weixin_48101150/article/details/115083236
	client := influxdb2.NewClient(userSetup.Url, userSetup.Token)
	writeAPI := client.WriteAPIBlocking(userSetup.Org, userSetup.Bucket)
	for i:=0;i<10000;i++{
		p := influxdb2.NewPoint("weather",
			map[string]string{"province": "河北"+strconv.Itoa(i),"station": "保定"+strconv.Itoa(i)},
			map[string]interface{}{"temperature": -1, "wind": 7},
			time.Now())
		writeAPI.WritePoint(context.Background(), p)
	}

	client.Close()
}

func Query(userSetup UserSetup){
	client := influxdb2.NewClient(userSetup.Url, userSetup.Token)
	queryAPI := client.QueryAPI(userSetup.Org)
	result, err := queryAPI.Query(context.Background(),
		`from(bucket:"bucket2")
	           |> range(start: -1d)
	           |> filter(fn: (r) => r._measurement == "weather")`)
	if err == nil {
		for result.Next() {
			//if result.TableChanged() {
			//	fmt.Printf("table: %s\n", result.TableMetadata().String())
			//}
			fmt.Printf("measurement: %v, timestamp: %v, tagvalue1: %v,tagvalue2: %v,fieldkey: %v, fieldvalue: %v\n",
				result.Record().Measurement(),
				result.Record().Time(),
				result.Record().ValueByKey("province"),
				result.Record().ValueByKey("station"),
				result.Record().Field(),
				result.Record().Value())
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} /*else {
		panic(err)
	}*/
	client.Close()
}