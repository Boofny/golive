package golive

import (
	"testing"
)

type User struct{
	ID 		int 		`json:"id"`
	Name 	string 	`json:"name"`
}

const (
	userJSON  = `{"id": 1, "name": "David"}`
	usersJSON = `[{"id": 1, "name": "David"}]`
)

const (
	PretyJSON = `
	{
		"id": 1,
		"name": "David"
	}
	`
)

func TestGolive(t *testing.T){
}

