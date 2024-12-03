package model_test

import (
	"github.com/bedrock-gophers/cooldown/cooldown"
	"github.com/restartfu/solar/internal/core/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser(t *testing.T) {
	mockUser := model.User{
		Name:        "test",
		DisplayName: "TEST",
		Invitations: make(cooldown.MappedCoolDown[string]),
	}

	u := model.NewUser("TEST")
	require.Equal(t, mockUser, u)
}
