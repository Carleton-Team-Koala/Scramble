// Special thanks to this package for some inspiration: https://pkg.go.dev/gitlab.com/ffe4/exercism-go/scrabble-score#section-readme

package main

import (
    "encoding/json"
    "fmt"
    "os"
    "net/http"
)

type alphabet struct {
    ListOfWords map[string]int `json:"listOfWords"`
}

var activeAlphabet = new(alphabet)

func letterReturner(w http.ResponseWriter, r *http.Request) {
  // recreate json dataset of active alphabet
  jsonData, err := json.Marshal(activeAlphabet)
    if err != nil {
        // Handle the error if marshaling fails
        http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
        return
    }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonData)
}

func importJSONdata(){
   // Open the JSON file
   fmt.Println("Opening file...")
   file, err := os.Open("englishAlphabet.json")
   if err != nil {
       fmt.Println("Error opening file:", err)
       return
   }
   defer file.Close()
 
   // Create a decoder to read the JSON data
   decoder := json.NewDecoder(file)
 
   // Decode the JSON data into the struct
   err = decoder.Decode(&activeAlphabet)
   if err != nil {
       fmt.Println("Error decoding JSON:", err)
       return
   }
}

func main() {
  importJSONdata()
  http.HandleFunc("/letterScores", letterReturner)
  http.ListenAndServe(":8080", nil)
}