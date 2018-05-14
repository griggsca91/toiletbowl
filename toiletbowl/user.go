package toiletbowl

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Poos     []Poo
}

func (u *User) Save() {

}
