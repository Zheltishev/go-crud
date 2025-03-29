package database

import (
	"crud/model"

	"github.com/gocraft/dbr"
)

type PersonRepository struct {
    Session *dbr.Session
}

func NewPersonRepository(session *dbr.Session) *PersonRepository {
    return &PersonRepository{Session: session}
}

func (r *PersonRepository) GetAll() ([]model.Person, error) {
    var persons []model.Person
    _, err := r.Session.Select("*").From("persons").Load(&persons)
    return persons, err
}

func (r *PersonRepository) GetByID(id int) (*model.Person, error) {
    var person model.Person
    err := r.Session.Select("*").From("persons").Where("id = ?", id).LoadOne(&person)
    return &person, err
}

func (r *PersonRepository) Create(person *model.Person) error {
    _, err := r.Session.InsertInto("persons").
        Columns("email", "phone", "firstname", "lastname").
        Record(person).
        Exec()
    return err
}

func (r *PersonRepository) Update(id int, person *model.Person) error {
    _, err := r.Session.Update("persons").
        Set("email", person.Email).
        Set("phone", person.Phone).
        Set("firstname", person.FirstName).
        Set("lastname", person.LastName).
        Where("id = ?", id).
        Exec()
    return err
}

func (r *PersonRepository) Delete(id int) error {
    _, err := r.Session.DeleteFrom("persons").
        Where("id = ?", id).
        Exec()
    return err
}
