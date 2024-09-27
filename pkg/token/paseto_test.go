package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"simple_bank/pkg"
)

var symmetricKey = pkg.RandomString(32)

func TestCreateToken(t *testing.T) {
	token := createToken(t)
	require.NotEmpty(t, token)
	require.NotNil(t, token)
}

func TestVerifyToken(t *testing.T) {
	token := createToken(t)
	require.NotEmpty(t, token)
	require.NotNil(t, token)

	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotNil(t, payload)
	t.Log(payload)
}

func createToken(t *testing.T) string {
	maker, err := NewPasetoMaker(symmetricKey)
	require.NoError(t, err)
	require.NotNil(t, maker)
	require.NotEmpty(t, maker)

	tokenID := uuid.New()
	username := pkg.RandomString(5)
	token, err := maker.CreateToken(tokenID, username, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	t.Log("tokenID", tokenID)
	t.Log("username", username)
	t.Log(token)

	return token
}
