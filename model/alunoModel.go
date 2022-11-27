package model

type Aluno struct {
	ID        int     `json: "id"`
	NomeAluno string  `json:"nome"`
	Atividade string  `json: "atividade"`
	NotaAluno float32 `json: "notas"`
}

type ListaDeAlunos struct {
	ListaDeAlunos []Aluno `json:"aluno"`
}
