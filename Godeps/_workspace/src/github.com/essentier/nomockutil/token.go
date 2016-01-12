package nomockutil

import (
	"encoding/json"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-errors/errors"
)

func GetSubjectInToken(tokenString string) (subject string, err error) {
	tokenParts := strings.Split(tokenString, ".")
	if len(tokenParts) != 3 {
		return "", errors.Errorf("The token %v is not a valid JWT token.", tokenString)
	}

	jsonData, err := jwt.DecodeSegment(tokenParts[1])
	if err != nil {
		return "", errors.Wrap(err, 1)
	}

	var jwtClaims map[string]interface{}
	err = json.Unmarshal(jsonData, &jwtClaims)
	if err != nil {
		return "", errors.Wrap(err, 1)
	}

	sub, exists := jwtClaims["sub"]
	if !exists {
		return "", errors.New("The token string does not contain the subject.")
	}

	return sub.(string), nil
}
