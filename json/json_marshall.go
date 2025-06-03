package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	str := `{
				"online": true,
						"owner_id": "202231538",
						"product_id": "tda8yu7ankwszilh",
						"product_name": "Security Camera",
						"status": [
							{
								"code": "basic_flip",
								"value": false
							},
							{
								"code": "basic_osd",
								"value": true
							},
							{
								"code": "basic_private",
								"value": false
							},
							{
								"code": "motion_sensitivity",
								"value": "1"
							},
							{
								"code": "sd_storge",
								"value": "896|896|0"
							},
							{
								"code": "sd_status",
								"value": 5
							},
							{
								"code": "sd_format",
								"value": false
							},
							{
								"code": "motion_record",
								"value": false
							},
							{
								"code": "movement_detect_pic",
								"value": "$"
							},
							{
								"code": "ptz_stop",
								"value": true
							},
							{
								"code": "sd_format_state",
								"value": 0
							},
							{
								"code": "ptz_control",
								"value": "7"
							},
							{
								"code": "record_loop",
								"value": true
							},
							{
								"code": "motion_switch",
								"value": false
							},
							{
								"code": "record_switch",
								"value": true
							},
							{
								"code": "record_mode",
								"value": "1"
							},
							{
								"code": "motion_tracking",
								"value": false
							},
							{
								"code": "device_restart",
								"value": false
							},
							{
								"code": "motion_area_switch",
								"value": false
							},
							{
								"code": "motion_area",
								"value": "{\"num\":1,\"region0\":{\"x\":0,\"y\":0,\"xlen\":100,\"ylen\":100}}"
							},
							{
								"code": "alarm_message",
								"value": ""
							}
						],
						"sub": false
				}`

	var result Result
	err := json.Unmarshal([]byte(str), &result)

	if err != nil {
		fmt.Println(err)
	}

	if isMotionDetectionEnable(result) {
		fmt.Println("motion detection enable")
	}

}

func isMotionDetectionEnable(result Result) bool {
	for _, item := range result.Status {

		if item.Code == "motion_switch" && item.Value == true {
			return true
		}
	}

	return false
}

type Result struct {
	Online      bool   `json:"online"`
	OwnerID     string `json:"owner_id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Status      []struct {
		Code  string      `json:"code"`
		Value interface{} `json:"value"`
	} `json:"status"`
	Sub bool `json:"sub"`
}
