package profile

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Address   string
}

func (u *User) GetProfile() {
	// u.Firstname = "Ambar"
	// u.Lastname = "Shafaa"
	// u.Address = "Jogja"

	fmt.Println(*u)
}

func (u *User) SetProfile(fname string, jname string, address string) {
	u.Firstname = fname
	u.Lastname = jname
	u.Address = address

	fmt.Println(*u)
}
