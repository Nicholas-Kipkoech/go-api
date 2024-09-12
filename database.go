package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "nickey968@gmail.com",
		Id:    "User1",
		Name:  "Nicholas Kipkoech",
		Token: "123",
	},
	"user2": {
		Email: "reen@123.com",
		Id:    "User2",
		Name:  "Maureen Chebet",
		Token: "124",
	},
}
