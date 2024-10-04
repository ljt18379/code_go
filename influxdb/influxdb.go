package main
import(
	"influxDemo/test"
)
func main(){
	userSetup:=test.UserSetup{
		Url:    "http://192.168.230.130:8086",
		Org:    "org1",
		Bucket: "bucket1",
		Token:  "IqpL4hCulj0cWhkDf5mZDGScnv3u1taSU4UX4b_UQtKoWX2hqgoqM4YtwJsP3hmbiyHQX_Ki5Pbt6H2GpqMQ7Q==",
	}
	fmt.printf("")
	test.Write(userSetup)
	//test.Query(userSetup)

}
