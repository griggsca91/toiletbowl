package rename

import "time"

type User struct {
	Email          string
	HashedPassword string
	Salt           string
	RegisteredOn   time.Time
}

func (u *User) Save() {

}
