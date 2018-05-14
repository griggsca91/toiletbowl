package toiletbowl

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Poo struct {
	gorm.Model
	Content string
	UserId  uint
}

func (u *Poo) Save() {

}

func (u *Poo) AfterSave(scope *gorm.Scope) error {
	log.Println("AfterSave")
	return nil
}
