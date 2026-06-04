package edit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (e *EditorArgs) JwtDecode(input string) (string, error) {
	key := e.Key
	if len(key) == 0 {
		key = "x"
	}

	if len(e.Key) > 0 {
		token, err := jwt.Parse(input, func(token *jwt.Token) (any, error) {
			return []byte(key), nil
		})

		if errors.Is(err, jwt.ErrTokenMalformed) {
			return input, fmt.Errorf("Token malformed")
		}

		claimsJson, jsonerr := getTokenJson(token)
		if jsonerr != nil {
			return input, fmt.Errorf("Error serializing claims: %v", jsonerr)
		}

		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return claimsJson, fmt.Errorf("Invalid signature")
		}

		if errors.Is(err, jwt.ErrTokenExpired) {
			return claimsJson, fmt.Errorf("Token expired")
		}

		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return claimsJson, fmt.Errorf("Token not valid yet")
		}

		if !token.Valid {
			return input, fmt.Errorf("Invalid token %v", err)
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

func (e *EditorArgs) JwtEncode(input string) (string, error) {
	claims := &jwt.MapClaims{}
	err := json.Unmarshal([]byte(input), claims)
	if err != nil {
		return input, fmt.Errorf("Error converting JSON input to claims: %v", err)
	}
	alg, err := getAlgorithm(e.Option)
	if err != nil {
		return input, err
	}

	token := jwt.NewWithClaims(alg, claims)

	tokenStr, err := token.SignedString([]byte(e.Key))
	if err != nil {
		return input, fmt.Errorf("Error signing token: %v", err)
	}

	return tokenStr, nil
}

func allowedAlgorithms() (map[string]jwt.SigningMethod) {
	return map[string]jwt.SigningMethod {
		"HMAC": jwt.SigningMethodHS256,
		"HS256": jwt.SigningMethodHS256,
		"HS384": jwt.SigningMethodHS384,
		"HS512": jwt.SigningMethodHS512,
		"ECDSA": jwt.SigningMethodES256,
		"ES256": jwt.SigningMethodES256,
		"ES384": jwt.SigningMethodES384,
		"ES512": jwt.SigningMethodES512,
		"RSA": jwt.SigningMethodRS256,
		"RS256": jwt.SigningMethodRS256,
		"RS384": jwt.SigningMethodRS384,
		"RS512": jwt.SigningMethodRS512,
	}
}

func JwtSigningMethods() []string {
	i := 0;
	sm := allowedAlgorithms()
	output := make([]string, len(sm))
	for k := range sm {
		output[i] = k
		i++
	}

	return output
}

func getAlgorithm(agName string) (jwt.SigningMethod, error) {
	algs := allowedAlgorithms()
	if strings.Trim(agName, " ") == "" {
		agName = "HMAC"
	}

	alg, ok := algs[strings.ToUpper(agName)]
	if !ok {
		algKeys := JwtSigningMethods()
		return jwt.SigningMethodHS256, fmt.Errorf("No signing algorithm %s, use one of:\n%v", agName, algKeys)
	}

	return alg, nil
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
