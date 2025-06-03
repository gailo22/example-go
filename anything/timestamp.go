package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	ts := 1726642038960

	// loc, _ := time.LoadLocation("UTC")

	// // setup a start and end time
	// createdAt := time.Now().In(loc).Add(1 * time.Hour)
	// expiresAt := time.Now().In(loc).Add(4 * time.Hour)

	// // get the diff
	// diff := expiresAt.Sub(createdAt)
	// fmt.Printf("Lifespan is %+v", diff)

	localTime := time.Now().Local().Unix()
	utcTime := time.Now().Unix()

	fmt.Println("ts:", ts)
	fmt.Println("localTime:", localTime)
	fmt.Println("utcTime:", utcTime)

	// layout := "2006-01-02 15:04:05 +"
	// value := "2024-09-20 00:00:00"

	// date, _ := time.Parse(layout, value)
	// fmt.Printf("date: %v\n", date)
	// date.UTC()

	// then := time.Date(2024, 9, 18, 20, 34, 58, 651387237, time.UTC)

	timeString := "2024-09-19T00:00:00+07:00"

	tm, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tm) // Output: 2024-09-20 00:00:00 +0700 +07 ICT

	tmTs := tm.UTC().UnixMilli()
	nowTs := time.Now().UTC().UnixMilli()
	fmt.Println("tmTs:", tmTs, " nowTs:", nowTs)
	if tmTs < nowTs {
		fmt.Println("xxx")
	}

	currentTS := time.Now().UTC().Unix()
	fmt.Println(currentTS)

	current := time.Now()
	fmt.Println("origin : ", current.String())
	fmt.Println("yyyy-mm-dd : ", current.Format("2006-01-02"))

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println(err)
	}
	unixTimeUTC := time.Unix(1730934652, 0).In(loc)
	unitTimeInDateOnly := unixTimeUTC.Format(time.DateOnly)
	fmt.Println(unitTimeInDateOnly)

}
