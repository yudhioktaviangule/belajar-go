package login

type UserLogin struct {
	ID       int
	Email    string
	Password string
	Token    string
	Name     string
}
type User struct {
	ID    int
	Email string
	Name  string
	Token string
}
type UserLoginPost struct {
	Email    string
	Password string
}
