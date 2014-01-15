// slope_one_test.go
package slopeone

import (
	"testing"
)

const (
	ItemSalt  = "salt"
	ItemCandy = "candy"
	ItemDog   = "dog"
	ItemCat   = "cat"
	ItemWar   = "war"
	ItemFood  = "strange food"
)

func TestSlopeOne(t *testing.T) {
	users := make([]map[string]float32, 4)

	users[0] = map[string]float32{ItemCandy: 1.0, ItemDog: 0.5, ItemWar: 0.1}
	users[1] = map[string]float32{ItemCandy: 1.0, ItemCat: 0.5, ItemWar: 0.2}
	users[2] = map[string]float32{ItemCandy: 0.9, ItemDog: 0.4, ItemCat: 0.5, ItemWar: 0.1}
	users[3] = map[string]float32{ItemCandy: 0.1, ItemWar: 1.0, ItemFood: 0.4}

	so := NewSlopeOne(users)

	user := make(map[string]float32)
	user[ItemFood] = 0.4
	t.Log(so.Predict(user))
	user[ItemWar] = 0.2
	t.Log(so.Predict(user))
}
