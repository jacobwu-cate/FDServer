package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	// "strconv"
)

type person struct {
	name string
	id int
  timesServed int
	haveMet []int
	previousAssignments []string
  currentAssignment string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n\nIncoming Request: ") // Print info
	fmt.Println("Method: ", r.Method, " ", r.URL)

	headerKeys := make([]string, len(r.Header)) // Placeholder
	i := 0
	for k := range r.Header {
		headerKeys[i] = k
		i++
	} // Get specific information
	
	for _, line := range headerKeys {
		fmt.Println("  > ", line, ":", r.Header.Get(line))
	} // Show Client Headers
	
	jsonString := string(encodeJSON())
	fmt.Fprintf(w, jsonString, r.URL.Path[1:]) // Answer the Client request
} // From PowerSchool Site, McF.
 
func encodeJSON() []byte {
  csvFile, err := os.Open("myFile.csv")
  if err != nil {
    fmt.Println(err)
  }
  defer csvFile.Close()
	
	reader := csv.NewReader(csvFile)
  reader.FieldsPerRecord = -1
 
  csvData, err := reader.ReadAll()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

	var allPersons [][]string
  var persons []string

	for _, table := range csvData {
		persons = make([]string, 0)
		for i := 1; i < len(table); i ++ {
			// fmt.Print(table[i])
			persons = append(persons, table[i])
		}
		// fmt.Print(persons,"\n\n")
		allPersons = append(allPersons, persons)
	}
 
	// Convert to JSON
	jsonData, err := json.Marshal(allPersons)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return jsonData

	/* fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./tableData.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close() */
} // http://www.cihanozhan.com/converting-csv-data-to-json-with-golang/

func main() {
	fmt.Print(string(encodeJSON()))
  // setupRoutes()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}