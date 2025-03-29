package logic

import (
	"crud/model"
	"crud/database"
)

type PersonLogic struct {
    Repo *database.PersonRepository
}

func NewPersonLogic(repo *database.PersonRepository) *PersonLogic {
    return &PersonLogic{Repo: repo}
}

func (l *PersonLogic) GetAllPersons() ([]model.Person, error) {
    return l.Repo.GetAll()
}

func (l *PersonLogic) GetPersonByID(id int) (*model.Person, error) {
    return l.Repo.GetByID(id)
}

func (l *PersonLogic) CreatePerson(person *model.Person) error {
    return l.Repo.Create(person)
}

func (l *PersonLogic) UpdatePerson(id int, person *model.Person) error {
    return l.Repo.Update(id, person)
}

func (l *PersonLogic) DeletePerson(id int) error {
    return l.Repo.Delete(id)
}
