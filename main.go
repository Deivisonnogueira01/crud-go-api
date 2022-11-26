package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/alunos", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {

			var aluno model.Aluno
			err := json.NewDecoder(req.Body).Decode(&aluno)
			if err != nil {
				fmt.Printf("Desculpe, Não consegui Ler :( : %s", err.Error())
				http.Error(resp, "Erro  ao Criar Dado, la ele", http.StatusBadRequest)
				return
			}

			if aluno.ID <= 0 {
				http.Error(resp, "Error", http.StatusBadRequest)
				return
			}

			resp.WriteHeader(http.StatusCreated)
			return
		}

		http.Error(resp, "Error", http.StatusBadRequest)

	})
}
