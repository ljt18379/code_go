package test

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)
func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.230.130:8086",
		Username: "lijiangtao",
		Password: "12345678",
	})
	if err != nil {
		log.Fatal(err)
	}
	return cli
}

// query
func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "bucket1",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// insert
func writesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "bucket1",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	tags := map[string]string{
		"province": "黑龙江",
		"station": "大庆",
	}
	fields := map[string]interface{}{
		"temperature": -9,
		"wind": 10,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}

//func main(){
//	writesPoints2()
//}

//func main() {
//  Go语言操作influxDB
//  https://www.fdevops.com/docs/golang-tutorial/influxdb-operating
//	conn := connInflux()
//	fmt.Println(conn)
//
//	// insert
//	//writesPoints(conn)
//
//	// 获取10条数据并展示
//	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "weather", 3)
//	res, err := queryDB(conn, qs)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, row := range res[0].Series[0].Values {
//		for j, value := range row {
//			log.Printf("j:%d value:%v\n", j, value)
//		}
//	}
//}





