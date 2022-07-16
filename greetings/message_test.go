package greetings_test

import (
	"testing"

	"github.com/myugen/tdd-go-template/greetings"
	"github.com/stretchr/testify/assert"
)

func TestMessage_Say(t *testing.T) {
	type fields struct {
		to string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should show greeting message to someone",
			fields: fields{to: "John Doe"},
			want:   "Hello John Doe!",
		},
		{
			name:   "should show greeting message to everyone",
			fields: fields{},
			want:   "Hello everyone!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := greetings.NewMessage(tt.fields.to)
			assert.Equal(t, tt.want, message.Say())
		})
	}
}
