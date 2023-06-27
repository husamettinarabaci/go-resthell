package domain

import (
	"testing"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
)

func Test_IsCommandRequestEntityValid(t *testing.T) {

	var tests = []struct {
		name           string
		commandRequest me.CommandRequest
		want           error
	}{
		{name: "empty", commandRequest: me.CommandRequest{}, want: mo.ErrInvalidInput},
		{name: "non_values", commandRequest: me.CommandRequest{Id: uuid.New(), Command: mo.Command{}}, want: mo.ErrInvalidInput},
		{name: "valid", commandRequest: me.CommandRequest{Id: uuid.New(), Command: mo.Command{Values: []string{"ls", "-al"}}}, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NewService().IsCommandRequestEntityValid(tt.commandRequest)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
