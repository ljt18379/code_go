package main

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

func main() {

}
func InfluxdbWrite() {
	url := "http://192.168.230.130:8086"
	token := "UdabIU-3xzrWaQ6OUHyz0K6IbLIcM_QXrmWbyOVkkUBT2Qc0QlPVLchzZZ5D-IpN1RRXGQui5tkmXcM7CKmefg=="
	org := "org1"
	bucket := "bucket1"
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	index := make(map[string]string)
	value := make(map[string]interface{})
	index["module"] = "module-tebt2"
	index["user"] = "user-test2"
	index["pid"] = "12"
	index["unum"] = "0f2"
	value["cost"] = 12
	value["uid"] = "1f0f4ceb0002642f7b89"
	p := influxdb2.NewPoint("poedb", index, value, time.Now())
	writeAPI.WritePoint(context.Background(), p)
}
func InfluxdbQuery() [][]interface{} {
	startTime := time.Now().AddDate(-1, 0, -2).Format("2006-01-02T00:00:00Z")
	stopTime := time.Now().AddDate(0, 0, -2).Format("2006-01-02T00:00:00Z")
	var queryString = `from(bucket:"bucket1")
	          |> range(start: ` + startTime + `,stop: ` + stopTime + `)
	          |> filter(fn: (r) => r._measurement == "poedb")`
	fmt.Printf("result:%v\n", queryString)
	url := "http://192.168.230.130:8086"
	token := "UdabIU-3xzrWaQ6OUHyz0K6IbLIcM_QXrmWbyOVkkUBT2Qc0QlPVLchzZZ5D-IpN1RRXGQui5tkmXcM7CKmefg=="
	org := "org1"
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)
	data, err := queryAPI.Query(context.Background(), queryString)
	var result [][]interface{}
	if err == nil {
		for data.Next() {
			var rs []interface{}
			//rs = append(rs, data.Record().Measurement())
			rs = append(rs, data.Record().Time())
			rs = append(rs, data.Record().ValueByKey("module"))
			rs = append(rs, data.Record().ValueByKey("pid"))
			rs = append(rs, data.Record().ValueByKey("unum"))
			rs = append(rs, data.Record().ValueByKey("user"))
			rs = append(rs, data.Record().Value())
			_, b := utils.SearchMatrix(result, rs)
			if b {
				for k, v := range result {
					if v[0] == rs[0] {
						result[k] = append(result[k], rs[len(rs)-1])
					}
				}
			} else {
				result = append(result, rs)
			}
		}

		if data.Err() != nil {
			fmt.Printf("query parsing error: %s\n", data.Err().Error())
		}
	} else {
		panic(err)
	}
	client.Close()
	fmt.Println(result)
	return result

}

func InfluxdbQuery2() [][]interface{} {
	result := InfluxdbQuery()
	var res [][]interface{}
	var f = []string{"module", "pid", "user", "sum"}
	var field []interface{}
	for _, v := range f {
		field = append(field, v)
	}
	res = append(res, field)
	for _, v := range result {
		var vv []interface{}
		vv = append(vv, v[1], v[2], v[4])
		row, b := utils.SearchMatrix(res, vv)
		if !b {
			vv = append(vv, 1)
			res = append(res, vv)
		} else {
			if res[row][3] != "sum" {
				res[row][3] = res[row][3].(int) + 1
			}
		}
	}
	fmt.Println(res)
	return res
}

func Query() {

	var queryString = `from(bucket:"bucket1")
	          |> range(start: ` + "2021-01-01T00:00:00Z" + `)
	          |> filter(fn: (r) => r._measurement == "poedb")`
	url := "http://192.168.230.130:8086"
	token := "UdabIU-3xzrWaQ6OUHyz0K6IbLIcM_QXrmWbyOVkkUBT2Qc0QlPVLchzZZ5D-IpN1RRXGQui5tkmXcM7CKmefg=="
	org := "org1"
	client := influxdb2.NewClient(url, token)
	queryAPI := client.QueryAPI(org)
	data, err := queryAPI.Query(context.Background(), queryString)
	var result []interface{}
	count := 0
	if err == nil {

		for data.Next() {
			count = count + 1
			var rs []interface{}
			rs = append(rs, data.Record().Time())
			rs = append(rs, data.Record().ValueByKey("module"))
			rs = append(rs, data.Record().ValueByKey("pid"))
			rs = append(rs, data.Record().ValueByKey("unum"))
			rs = append(rs, data.Record().ValueByKey("user"))
			rs = append(rs, data.Record().Field())
			rs = append(rs, data.Record().Value())
			result = append(result, rs)
		}

		if data.Err() != nil {
			fmt.Printf("query parsing error: %s\n", data.Err().Error())
		}
	} else {
		panic(err)
	}
	client.Close()
	fmt.Println(result)
}

// 并发
func concurrent() {

}
