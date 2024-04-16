package logged

type Logged_in struct {
	ID   uint
	Role string
}

// our logged account that is a global variable
var loggedAcc Logged_in

// func that set the account
func SetLoggedAcc(id uint, role string) {
	loggedAcc.ID = id
	loggedAcc.Role = role
}
