package model

//domain

//PERSON
type Aluno struct {
	ID        int     `json: "id"`
	NomeAluno string  `json:"nome"`
	Atividade string  `json: "atividade"`
	NotaAluno float32 `json: "notas"`
}

//PEOPLE
type ListaDeAlunos struct {
	ListaDeAlunos []Aluno `json:"aluno"`
	//people
}

/*
 pasta person = pasta regras
   Service
*/

// PERSON É A ALUNO MODEL

// PEOPLE É A LISTADEALUNOS
