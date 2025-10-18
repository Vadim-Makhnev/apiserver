package store_test

import (
	"testing"

	"github.com/Vadim-Makhnev/apiserver/internal/app/model"
	"github.com/Vadim-Makhnev/apiserver/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)

	u, err = s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
