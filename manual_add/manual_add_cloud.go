package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type IPCDeviceAdded struct {
	DeviceID       string `json:"deviceId"`
	UID            string `json:"uid"`
	SpaceID        string `json:"spaceId"`
	ProductID      string `json:"productId"`
	UUID           string `json:"uuid"`
	Timestamp      int64  `json:"timestamp"`
	XCorrelationID string `json:"x-correlation-id"`
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file: %v", err)
	}

	return records, nil
}

func filterFields(headers []string, record []string, fieldsToInclude []string) ([]string, error) {
	var filteredRecord []string

	// Create a map of headers to their index positions
	headerIndexMap := make(map[string]int)
	for i, header := range headers {
		headerIndexMap[header] = i
	}

	// Filter only the fields that we want to include
	for _, field := range fieldsToInclude {
		if index, exists := headerIndexMap[field]; exists {
			filteredRecord = append(filteredRecord, record[index])
		} else {
			return nil, fmt.Errorf("field %s not found in CSV headers", field)
		}
	}

	return filteredRecord, nil
}

func publishToNATS(nc *nats.Conn, subject string, message []byte) error {
	err := nc.Publish(subject, message)
	if err != nil {
		return fmt.Errorf("error publishing message to NATS: %v", err)
	}
	return nil
}

func main() {
	natsURL := "nats://localhost:5222"
	subject := "IPC_ADDED_DEVICE.created"
	// filePath := "data.csv"
	filePath := "Export A2 list - Import.csv"

	records, err := readCSV(filePath)
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	// fieldsToInclude := []string{"devId", "uid", "product_id", "spaceId", "uuid"}
	fieldsToInclude := []string{"device_id", "product_id", "uid", "uuid"}

	// The first row of the CSV file is expected to be the header
	headers := records[0]

	// Connect to NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()

	count := 0
	for i, record := range records[1:] {
		filteredRecord, err := filterFields(headers, record, fieldsToInclude)
		if err != nil {
			log.Printf("Failed to filter row %d: %v", i+2, err)
			continue
		}

		deviceAdded := IPCDeviceAdded{
			DeviceID:       filteredRecord[0],
			ProductID:      filteredRecord[1],
			UID:            filteredRecord[2],
			UUID:           filteredRecord[3],
			SpaceID:        "",
			Timestamp:      time.Now().UnixMilli(),
			XCorrelationID: uuid.New().String(),
		}
		message, _ := json.Marshal(deviceAdded)

		err = publishToNATS(nc, subject, message)
		count++
		if err != nil {
			log.Printf("Failed to publish row %d: %v", i+2, err)
		} else {
			fmt.Printf("Published row %d: %s\n", i+2, message)
		}

		fmt.Println(string(message))
	}

	fmt.Println("total: ", count)
}
