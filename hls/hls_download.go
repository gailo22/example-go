package main

import (
	"fmt"

	"github.com/canhlinh/hlsdl"
	"github.com/tuya/pulsar-client-go/pkg/log"
)

func main() {

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	// url := ""https://bitdash-a.akamaihd.net/content/sintel/hls/video/1500kbit.m3u8""
	url := "https://wework-1-us.stream.iot-11.com:8080/cloudrecord/eb675542a3f236ea384fak/csb7jsa59bakkorc1t4g76AJ7i7i9Oyf.m3u8"
	hlsDL := hlsdl.New(url, nil, "download", "", 64, true)

	filepath, err := hlsDL.Download()
	if err != nil {
		log.Error("error:", err)
		panic(err)
	}

	fmt.Println(filepath)
}
