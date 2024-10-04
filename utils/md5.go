/*************************************************
Copyright: cmri
Author: sunao@chinamobile.com
Date: 2021-03-26
Description: generate request id
**************************************************/
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)

func GetStringMD5(s string) string {
    md5 := md5.New()
    md5.Write([]byte(s))
    md5Str := hex.EncodeToString(md5.Sum(nil))
    return md5Str
}

func main(){
    s :=GetStringMD5("ljt")
    fmt.Println(s)
}
