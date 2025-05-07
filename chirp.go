package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

type parameters struct {
	Body string `json:"body"`
}

type returnVals struct {
	CleanedBody string `json:"cleaned_body"`
}

type errorJson struct {
	Error string `json:"error"`
}

func handlerChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := replaceProfaneWords(params.Body)
	respBody := returnVals{
		CleanedBody: cleanedBody,
	}

	respondWithJSON(w, http.StatusOK, respBody)
}

func replaceProfaneWords(text string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(text, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if slices.Contains(profaneWords, lowerWord) {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, errorJson{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
