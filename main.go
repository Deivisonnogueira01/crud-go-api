package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Deivisonnogueira01/crud-go-api/model"
	"github.com/Deivisonnogueira01/crud-go-api/model/regras"
)

func main() {
	service, err := regras.NewService("regras.json")
	if err != nil {
		fmt.Printf("Error trying to creating personService: %s\n", err.Error())
	}

	http.HandleFunc("/aluno/", func(resposta http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			path := strings.TrimPrefix(req.URL.Path, "/aluno/")
			if path == "" {
				// list all people
				resposta.WriteHeader(http.StatusOK)
				resposta.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(resposta).Encode(service.List())
				if err != nil {
					http.Error(resposta, "Error trying to list people", http.StatusInternalServerError)
					return
				}
			} else {
				idAluno, err := strconv.Atoi(path)
				if err != nil {
					http.Error(resposta, "Invalid id given. aluno ID must be an integer", http.StatusBadRequest)
					return
				}
				aluno, err := service.GetByID(idAluno)
				if err != nil {
					http.Error(resposta, err.Error(), http.StatusNotFound)
					return
				}
				resposta.WriteHeader(http.StatusOK)
				resposta.Header().Set("Content-Type", "application/json")
				err = json.NewEncoder(resposta).Encode(aluno)
				if err != nil {
					http.Error(resposta, "Error trying to get aluno", http.StatusInternalServerError)
					return
				}
			}
			return
		}
		if req.Method == "POST" {
			var aluno model.Aluno
			err := json.NewDecoder(req.Body).Decode(&aluno)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s\n", err.Error())
				http.Error(resposta, "Error trying to create aluno", http.StatusBadRequest)
				return
			}
			if aluno.ID <= 0 {
				http.Error(resposta, "aluno ID should be a positive integer", http.StatusBadRequest)
				return
			}

			err = service.Create(aluno)
			if err != nil {
				fmt.Printf("Error trying to create aluno: %s\n", err.Error())
				http.Error(resposta, "Error trying to create aluno", http.StatusInternalServerError)
				return
			}
			resposta.WriteHeader(http.StatusCreated)
			return
		}
		if req.Method == "DELETE" {
			path := strings.TrimPrefix(req.URL.Path, "/aluno/")
			if path == "" {
				http.Error(resposta, "ID is required to delete a aluno", http.StatusBadRequest)
				return
			} else {
				idAluno, err := strconv.Atoi(path)
				if err != nil {
					http.Error(resposta, "Invalid id given. aluno ID must be an integer", http.StatusBadRequest)
					return
				}
				err = service.DeleteByID(idAluno)
				if err != nil {
					fmt.Printf("Error trying to delete aluno: %s\n", err.Error())
					http.Error(resposta, "Error trying to delete aluno", http.StatusInternalServerError)
					return
				}
				resposta.WriteHeader(http.StatusOK)
			}
			return
		}
		if req.Method == "PUT" {
			var aluno model.Aluno
			err := json.NewDecoder(req.Body).Decode(&aluno)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body should be a json. Error: %s\n", err.Error())
				http.Error(resposta, "Error trying to update aluno", http.StatusBadRequest)
				return
			}
			if aluno.ID <= 0 {
				http.Error(resposta, "aluno ID should be a positive integer", http.StatusBadRequest)
				return
			}

			err = service.Update(aluno)
			if err != nil {
				fmt.Printf("Error trying to update aluno: %s\n", err.Error())
				http.Error(resposta, "Error trying to update aluno", http.StatusInternalServerError)
				return
			}
			resposta.WriteHeader(http.StatusOK)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
