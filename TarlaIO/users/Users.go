package users

//User's Json data to Go Struct to marshall

//We are creating struct in a hierarchical method
type Geo struct{
	Lat string
	Lng string
}

type Address struct{
	Street string
	Suite string
	City string
	ZipCode string
	Geo Geo
}

type Company struct {
	CompanyName string
	CatchPhrase string
	Bs string
}

//To collect every structure init, we should defined before declaration of struct
type User struct {
	Id int
	Name string
	Username string
	Email string
	Address Address
	Phone string
	Website string
	Company Company
	PostNumber int
}
