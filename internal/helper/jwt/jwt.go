package jwt

import (
	"errors"
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
	"github.com/malikfajr/halo-suster/config"
)

type JWTClaim struct {
	ID   string            `json:"id"`
	Nip  int               `json:"nip"`
	Role string            `json:"role"`
	Exp  *gjwt.NumericDate `json:"exp"`
	gjwt.RegisteredClaims
}

func CreateToken(ID string, nip int, role string) string {
	claim := &JWTClaim{
		ID:   ID,
		Nip: nip,
		Role: role,
		Exp:  gjwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
	}

	token := gjwt.NewWithClaims(gjwt.SigningMethodHS512, claim)

	signedToken, err := token.SignedString([]byte(config.JWT.SECRET))
	if err != nil {
		panic(err)
	}

	return string(signedToken)
}

func ClaimToken(token string) (*JWTClaim, error) {
	parsed, err := gjwt.ParseWithClaims(token, &JWTClaim{}, func(token *gjwt.Token) (interface{}, error) {
		return []byte(config.JWT.SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if parsed.Method != gjwt.SigningMethodHS512 {
		return nil, errors.New("invalid token")
	}

	if claims, ok := parsed.Claims.(*JWTClaim); ok {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}

}
