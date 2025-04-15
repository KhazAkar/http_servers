package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		// these tags indicate how the keys in the JSON should be mapped to the struct fields
		// the struct fields must be exported (start with a capital letter) if you want them parsed
		Body string `json:"body"`
	}
	type returnVals struct {
		// the key will be the name of struct field unless you give it an explicit JSON tag
		Valid bool `json:"valid"`
	}
	type errorJson struct {
		Error string `json:"error"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		errorResp := errorJson{
			Error: "Something went wrong",
		}
		dat, err := json.Marshal(errorResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			w.Write(dat)
			return
		}
	}
	respBody := returnVals{
		Valid: true,
	}
	if len(params.Body) > 41 {
		respBody.Valid = false
		w.WriteHeader(400)
		errorJson := errorJson{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(errorJson)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(dat)
		return
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
