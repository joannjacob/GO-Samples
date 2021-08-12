package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main_project/models"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Property struct {
	csvPos          float64
	rdsColumnName   string
	datatype        string
	hasDefaultValue bool
	// ifdefault       nil `json:"default,omitempty"`
}
type Fields struct {
	SKUCode     Property `json:SKUCode`
	Name        Property
	Description Property
	Color       Property
	Size        Property
}

var wg sync.WaitGroup

func main() {
	ImportData()
}

func LoadConfiguration() map[string]int {
	// var config Fields
	var result map[string]map[string]interface{}

	configFile, err := os.Open("./config/csv_config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.NewDecoder(configFile).Decode(&result)
	// json.NewDecoder(configFile).Decode(&config)

	configMap := make(map[string]map[string]string)
	for k, v := range result["Fields"] {
		innerMap := make(map[string]string)
		for i, val := range v.(map[string]interface{}) {
			switch val.(type) {
			case int:
				innerMap[i] = strconv.Itoa(val.(int))
			case string:
				innerMap[i] = val.(string)
			case bool:
				innerMap[i] = strconv.FormatBool(val.(bool))
			case float64:
				innerMap[i] = fmt.Sprintf("%f", val.(float64)-1)
			}
		}
		configMap[k] = innerMap
	}
	modelMap := make(map[string]int)
	for _, value := range configMap {
		position, _ := strconv.ParseFloat(value["csvPos"], 64)
		modelMap[value["rdsColumnName"]] = int(position)

	}
	return modelMap
}

func saveProduct(line []string, modelMap map[string]int) {
	product := &models.Product{}
	productMap := make(map[string]interface{})
	for key, value := range modelMap {
		productMap[key] = line[value]

	}
	mapstructure.Decode(productMap, &product)
	product.Create()

	wg.Done()
}

func processRecords(records chan []string) {
	modelMap := LoadConfiguration()
	for record := range records {
		wg.Add(1)
		go saveProduct(record, modelMap)
	}
}

var ImportData = func() {
	startTime := time.Now()
	// modelMap := LoadConfiguration()

	csvFile, err := os.Open("./config/input1.csv")
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvFile.Close()
	row1, err := bufio.NewReader(csvFile).ReadSlice('\n')
	if err != nil {
		log.Fatal(err)
	}
	_, err = csvFile.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	// reader := csv.NewReader(bufio.NewReader(csvFile))

	// for {
	// 	wg.Add(1)
	// 	line, error := reader.Read()
	// 	if error == io.EOF {
	// 		break
	// 	} else if error != nil {
	// 		log.Fatal(error)
	// 	}
	// 	go saveProduct(line, modelMap)
	// }

	records := make(chan []string)
	go func() {
		reader := csv.NewReader(bufio.NewReader(csvFile))

		defer close(records)
		for {
			line, error := reader.Read()
			if error == io.EOF {
				break
			} else if error != nil {
				log.Fatal(error)
			}
			// wg.Add(1)
			// saveProduct(line, modelMap)
			records <- line
		}
	}()
	processRecords(records)

	fmt.Println("Successfully imported data to database in ", time.Now().Sub(startTime))
	wg.Wait()
}
