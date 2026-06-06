package golive

import (
	"testing"
)

type User struct{
	ID 		int 		`json:"id"`
	Name 	string 	`json:"name"`
}

const (
	userJson = `{"id": 1, "name": "David"}`
	usersJson = `[{"id": 1, "name": "David"}]`
)

const (
	PretyJson = `
	{
		"id": 1,
		"name": "David"
	}
	`
)

func TestGolive(t *testing.T){
}

