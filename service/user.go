package service

import "github.com/google/uuid"

type Id uuid.UUID

type PostBody struct {
	FirstName string
	LastName  string
	Biography string
}

type User struct {
	Id        Id
	FirstName string
	LastName  string
	Biography string
}

type Application struct {
	Data map[Id]User
}

func (app *Application) FindAll() []User {
	users := make([]User, 0, len(app.Data))

	for _, u := range app.Data {
		users = append(users, u)
	}

	return users
}

func (app *Application) FindById(id Id) (User, bool) {
	u, ok := app.Data[id]
	return u, ok
}

func (app *Application) Insert(body PostBody) User {
	newId := Id(uuid.New())
	newUser := User{
		Id:        newId,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Biography: body.Biography,
	}
	app.Data[newId] = newUser

	return newUser
}

func (app *Application) Update(id Id, body PostBody) (User, bool) {
	u, ok := app.Data[id]
	if !ok {
		return User{}, false
	}

	u.FirstName = body.FirstName
	u.LastName = body.LastName
	u.Biography = body.Biography
	app.Data[id] = u
	return u, true

}

func (app *Application) Delete(id Id) bool {
	_, ok := app.Data[id]
	if ok {
		delete(app.Data, id)
		return true
	}
	return false
}
