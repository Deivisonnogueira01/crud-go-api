package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Deivisonnogueira01/crud-go-api/model"
)

type Service struct {
	dbFilePath string
	alunos     model.ListaDeAlunos
}

func NewService(dbFilePath string) (Service, error) {
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			err = createEmptyFile(dbFilePath)
			if err != nil {
				return Service{}, err
			}
			return Service{
				dbFilePath: dbFilePath,
				alunos:     model.ListaDeAlunos{},
			}, nil
		} else {
			return Service{}, err
		}
	}

	jsonFile, err := os.Open(dbFilePath)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to open file that contains all aluno: %s", err.Error())
	}

	jsonFileContentByte, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return Service{}, fmt.Errorf("Error trying to read aluno file: %s", err.Error())
	}

	var totalAlunos model.ListaDeAlunos
	json.Unmarshal(jsonFileContentByte, &totalAlunos)

	return Service{
		dbFilePath: dbFilePath,
		alunos:     totalAlunos,
	}, nil
}

func (s *Service) AddAluno(alunoModel model.Aluno) error {
	s.alunos.ListaDeAlunos = append(s.alunos.ListaDeAlunos, alunoModel)
	return s.saveFile()
}

func (s Service) saveFile() error {
	allPeopleJSON, err := json.Marshal(s.alunos)
	if err != nil {
		return fmt.Errorf("Error trying to encode aluno as JSON: %s", err.Error())
	}
	return ioutil.WriteFile(s.dbFilePath, allPeopleJSON, 0755)
}

func (s *Service) Create(alunoModel model.Aluno) error {
	if s.exists(alunoModel) {
		return fmt.Errorf("There is already a person with this ID registered")
	}

	err := s.AddAluno(alunoModel)
	if err != nil {
		return fmt.Errorf("Error trying to add Aluno to file: %s", err.Error())
	}

	return nil
}

func (s Service) exists(person model.Aluno) bool {
	for _, currentPerson := range s.alunos.ListaDeAlunos {
		if currentPerson.ID == person.ID {
			return true
		}
	}
	return false
}

func (s Service) List() model.ListaDeAlunos {
	return s.alunos
}

func (s Service) GetByID(personID int) (model.Aluno, error) {
	for _, currentPerson := range s.alunos.ListaDeAlunos {
		if currentPerson.ID == personID {
			return currentPerson, nil
		}
	}
	return model.Aluno{}, fmt.Errorf("Aluno not found")
}

func (s *Service) DeleteByID(personID int) error {
	var indexToRemove int = -1
	for index, currentPerson := range s.alunos.ListaDeAlunos {
		if currentPerson.ID == personID {
			indexToRemove = index
			break
		}
	}
	if indexToRemove < 0 {
		return fmt.Errorf("There is no person with the provided ID")
	}

	s.alunos.ListaDeAlunos = append(
		s.alunos.ListaDeAlunos[:indexToRemove],
		s.alunos.ListaDeAlunos[indexToRemove+1:]...,
	)

	return s.saveFile()
}

func (s *Service) Update(person model.Aluno) error {
	var indexToUpdate int = -1
	for index, currentPerson := range s.alunos.ListaDeAlunos {
		if currentPerson.ID == person.ID {
			indexToUpdate = index
			break
		}
	}
	if indexToUpdate < 0 {
		return fmt.Errorf("There is no person with the given ID to be updated")
	}

	s.alunos.ListaDeAlunos[indexToUpdate] = person
	return s.saveFile()
}

func createEmptyFile(dbFilePath string) error {
	var aluno model.ListaDeAlunos = model.ListaDeAlunos{
		ListaDeAlunos: []model.Aluno{},
	}
	peopleJSON, err := json.Marshal(aluno)
	if err != nil {
		return fmt.Errorf("Error trying to encode aluno as JSON: %s", err.Error())
	}

	err = ioutil.WriteFile(dbFilePath, peopleJSON, 0755)
	if err != nil {
		return fmt.Errorf("Error trying to writing aluno file: %s", err.Error())
	}

	return nil
}
