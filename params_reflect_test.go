package snippet

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestM_AssignTo(t *testing.T) {
	type Room struct {
		ID   int    `param:"-" json:"id"`
		Name string `param:"name" json:"name"`
	}

	type School struct {
		ID     int    `param:"-" json:"id"`
		Name   string `param:"name" json:"name"`
		RoomID int    `param:"room_id" json:"-"`
		Room   *Room  `param:"-" json:"room"`
	}
	params := MapStringInterface {"id": "123", "name":"primary school", "room_id":"1"}
	var s *School
	params.AssignTo(&s, "param")
	fmt.Println(s)
	assert := assert.New(t)
	assert.Equal(s.Name, "primary school")
	assert.Equal(s.RoomID, 1)
}
