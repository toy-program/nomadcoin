package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/toy-program/nomadcoin/blockchain"
	"github.com/toy-program/nomadcoin/utils"
)

const port string = ":3000"

type URL string

func (u *URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, *u)
	return []byte(url), nil
}

type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Data string `json:"data"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{URL: URL("/"), Method: "GET", Description: "See Documentations"},
		{URL: URL("/blocks"), Method: "POST", Description: "Add a Block", Payload: "data: string"},
	}

	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		err := json.NewDecoder(r.Body).Decode(&addBlockBody)
		utils.HandleErr(err)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Data)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
