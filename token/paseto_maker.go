package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	key, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	maker := &PasetoMaker{
		symmetricKey: key,
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	token, err := paseto.MakeToken(payload.GetMap(), nil)
	if err != nil {
		return "", err
	}
	tkn := token.V4Encrypt(maker.symmetricKey, nil)
	return tkn, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.MakeParser(nil)
	tkn, err := parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	m := tkn.Claims()

	payload, err := PayloadFromMap(m)
	if err != nil {
		fmt.Println("ОШИБКААААААААААААААААААААААААА ТУТТТТТТТТТТТ", err)
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		fmt.Println("ОШИБКААААААААААААААААААААААААА ТУТТТТТТТТТТТ", err)
		return nil, err
	}
	fmt.Println("ОШИБКИ ТУТТТТТТТТТТТ НЕТТТ", err)

	return payload, nil

}
