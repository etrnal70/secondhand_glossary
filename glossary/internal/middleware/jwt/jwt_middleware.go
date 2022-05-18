package jwt

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"secondhand_glossary/internal/config"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CustomClaims struct {
	UserID uint
	Admin  bool
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userId uint, isAdmin bool, conf config.Config) (*TokenDetails, error) {
	token := &TokenDetails{}

	token.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	atUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	token.AccessUuid = atUuid.String()

	token.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	rtUuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	token.RefreshUuid = rtUuid.String()

	atClaims := CustomClaims{
		UserID: userId,
		Admin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			Id:        token.AccessUuid,
			ExpiresAt: token.AtExpires,
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		return nil, err
	}

	rtClaims := CustomClaims{
		UserID: userId,
		Admin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			Id:        token.RefreshUuid,
			ExpiresAt: token.RtExpires,
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(conf.REFRESH_SECRET))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractTokenFromHeader(e echo.Context) string {
	header := e.Request().Header
	authString := header["Authorization"]
	var tokenString string

	// TODO Needs tests
	if len(authString) == 2 {
		tokenString = authString[1]
	} else {
		tokenString = authString[0]
	}

	return tokenString
}

func VerifyUserToken(auth string, _ echo.Context) (interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected JWT signing method=%v", t.Header["alg"])
		}
		// FIXME Dirty hacks
		return os.Getenv("JWT_SECRET"), nil
	}

	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func VerifyAdminToken(auth string, _ echo.Context) (interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected JWT signing method=%v", t.Header["alg"])
		}
		// FIXME Dirty hacks
		return os.Getenv("JWT_SECRET"), nil
	}

	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	tokenClaims := token.Claims.(jwt.MapClaims)
	if !tokenClaims["isAdmin"].(bool) {
		return nil, errors.New("unauthorized: not an admin")
	}
	return token, nil
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}

	return 0
}

func JWTError(err error, c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"message": "Unauthorized",
	})
}
