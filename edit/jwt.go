package edit

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (e *EditorArgs) JwtDecode(input string) (string, error) {
	key := e.Option
	if len(key) == 0 {
		key = "x"
	}

	if len(e.Option) > 0 {
		token, err := jwt.Parse(input, func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		})

		if errors.Is(err, jwt.ErrTokenMalformed) {
			return input, fmt.Errorf("Token malformed")
		}

		if !token.Valid {
			return input, fmt.Errorf("Invalid token")
		}

		claimsJson, err := getTokenJson(token)

		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return claimsJson, fmt.Errorf("Invalid signature")
		}

		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return claimsJson, fmt.Errorf("Token expired")
		}

		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return claimsJson, fmt.Errorf("Token not valid yet")
		}

		return claimsJson, nil
	} else {
		token, _, err := new(jwt.Parser).ParseUnverified(input, jwt.MapClaims{})
		if err != nil {
			return input, err
		}

		claimsJson, err := getTokenJson(token)
		if err != nil {
			return input, err
		}

		return claimsJson, nil
	}
}

func getTokenJson(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid claims type")
	}

	claimsJson, err := json.MarshalIndent(claims, "", settings().TabString)
	if err != nil {
		return "", err
	}

	return string(claimsJson), nil
}
