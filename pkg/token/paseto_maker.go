package token

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (p PasetoMaker) CreateToken(id uuid.UUID, username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(id, username, duration)
	if err != nil {
		return "", err
	}
	encrypt, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}

func (p PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	// 此版本要求密钥长度必须是32
	if len(symmetricKey) != 32 {
		return nil, errors.New("system key must be 32 bytes")
	}

	return &PasetoMaker{
		symmetricKey: []byte(symmetricKey),
		paseto:       paseto.NewV2(),
	}, nil
}
