package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/goellavish10/skillbee-assignment/interfaces"
	"github.com/goellavish10/skillbee-assignment/utils"
)

func GenerateStaticPages() {
	pages := os.Getenv("PAGES")

	if pages == "" {
		pages = "10"
	}

	pageCount, err := strconv.Atoi(pages)

	if err != nil {
		panic("Invalid number of pages")
	}
	utils.CreateDir("dist")
	var apiResponses []interfaces.ApiResponse
	_, err = os.Stat("dist/data.json")
	isJsonExists := false

	if err == nil {
		fmt.Println("JSON Exists! Skipping fetching data from API")
		isJsonExists = true
	}

	for i := 0; i < pageCount; i++ {
		if !isJsonExists {
			if i == 0 {
				fmt.Println("Initiating data fetch from API...")
			}
			url := "https://www.boredapi.com/api/activity"

			// Making a GET Request only when JSON already doesn't exist
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				log.Fatal("Error in fetching API Response")
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("Status Code: ", resp.StatusCode)
				log.Fatal("Error in fetching API Response")
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				log.Fatal("Error in reading API Response")
			}

			var responseData interfaces.ApiResponse
			err = json.Unmarshal(body, &responseData)

			if err != nil {
				fmt.Println("Error parsing JSON: ", err)
				log.Fatal("Error parsing JSON")
			}

			// Check if the 'Key' already exists in the local array of objects
			keyExists := false
			for _, item := range apiResponses {
				if item.Key == responseData.Key {
					keyExists = true
					break
				}
			}

			if !keyExists {
				apiResponses = append(apiResponses, responseData)
			}

			if i == pageCount-1 {
				fmt.Println("Data fetch from API completed!")
			}
		}
	}

	if !isJsonExists {
		fmt.Println("Creating a local copy of API data...")
		file, err := os.Create("dist/data.json")
		if err != nil {
			fmt.Println("Error creating JSON file: ", err)
			log.Fatal("Error creating JSON file")
		}

		defer file.Close()
		jsonData, err := json.Marshal(apiResponses)
		if err != nil {
			log.Fatal("\nError converting JSON to string: ", err)
		}
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatal("\nError writing to JSON file: ", err)
		}
		fmt.Println("Local copy of API data created!")
	}

	if len(apiResponses) == 0 {
		fileData, err := os.ReadFile("dist/data.json")

		if err != nil {
			log.Fatal("Error reading JSON file: ", err)
		}

		// Unmarshal the JSON data into the apiResponses slice
		err = json.Unmarshal(fileData, &apiResponses)
		if err != nil {
			log.Fatal("Error unmarshaling JSON data:", err)
		}
	}

	for i := 0; i < len(apiResponses); i++ {
		_, err := os.Stat(fmt.Sprintf("dist/page-%d.html", i+1))
		if err == nil {
			fmt.Println("Page Exists! Skipping generating page")
			continue
		}

		fmt.Printf("Generating static html page %d", i+1)
		var htmlString string

		templateFile, err := os.ReadFile("views/template.html")
		if err != nil {
			log.Fatal("Error reading template file: ", err)
		}
		htmlString = string(templateFile)

		htmlString = strings.Replace(htmlString, "{{TITLE}}", apiResponses[i].Activity, -1)
		htmlString = strings.Replace(htmlString, "{{KEY}}", apiResponses[i].Key, -1)
		htmlString = strings.Replace(htmlString, "{{PRICE}}", strconv.FormatFloat(apiResponses[i].Price, 'f', 2, 64), -1)
		htmlString = strings.Replace(htmlString, "{{TYPE}}", apiResponses[i].Type, -1)

		outputFile, err := os.Create(fmt.Sprintf("dist/page-%d.html", i+1))

		if err != nil {
			log.Fatal("Error creating output file: ", err)
		}
		defer outputFile.Close()

		err = os.WriteFile(fmt.Sprintf("dist/page-%d.html", i+1), []byte(htmlString), 0644)

		if err != nil {
			log.Fatal("Error writing to output file: ", err)
		}
	}

	fmt.Println("🎊 Static site generated!")

}
