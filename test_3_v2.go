package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type BeefCountResponse struct {
	BeefCounts map[string]int `json:"beef"`
}

type ErrorMessageResponse struct {
	ErrorMessage string `json:"error_message"`
}

func main() {
	http.HandleFunc("/api/beef/summary", getBeefSummary)

	port := ":5000"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func getBeefSummary(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		log.Printf("Error fetching data from Bacon Ipsum API: %v", err)
		response := ErrorMessageResponse {
			ErrorMessage: "can't get data.",
		}

		jsonData, _ := json.Marshal(response)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}
	defer resp.Body.Close()

	textData := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			break
		}
		textData = append(textData, buf[:n]...)
	}

	beefWords := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
		"chuck",
		"filet mignon",
		"rib",
		"loin",
		"shoulder",
		"brisket",
		"tri-tip",
	}

	beefCounts := make(map[string]int)

	words := strings.Fields(strings.ToLower(string(textData)))
	for _, word := range words {
		for _, beefWord := range beefWords {
			if strings.Contains(word, beefWord) {
				beefCounts[beefWord]++
			}
		}
	}

	response := BeefCountResponse{
		BeefCounts: beefCounts,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling JSON response: %v", err)
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
