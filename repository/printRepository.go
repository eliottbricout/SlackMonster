package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PrintRepository struct {
	url string
}

func CreatePrintRepository(url string) PrintRepository {
	return PrintRepository{url}
}

func (repo *PrintRepository) Print(text string) {
	requestBody, _ := json.Marshal(map[string]string{ "text": text})
	http.Post(repo.url, "application/json", bytes.NewBuffer(requestBody))
}