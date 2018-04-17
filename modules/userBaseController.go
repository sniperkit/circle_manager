package modules

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jungju/aha/errors"
	"github.com/sirupsen/logrus"
)

var (
	userTokenHeaderName = ""
	secretKeys          = ""
)

type UserMeta struct {
	UserID  uint
	GroupID uint
	Token   string
}

type CircleCustomClaims struct {
	UserID  uint
	GroupID uint
	jwt.StandardClaims
}

type getExistsObjectInUserFunc func(uint) (bool, error)

//BaseController ...
type UserBaseController struct {
	BaseCircleController
	CurrentUserMeta *UserMeta
}

// Prepare ...
func (c *UserBaseController) Prepare() {
	logrus.Debug("UserBaseController", "Prepare")
	var err error
	if c.CurrentUserMeta, err = c.GetCurrentUserMeta(); err != nil {
		c.ErrorAbort(401, nil)
	}

	//TODO: User 권한 맵. 403.

	//TODO: User 승인 여부 맵. 403.
}

func (c *UserBaseController) GetCurrentUserMeta() (*UserMeta, error) {
	tokenString := c.Ctx.Request.Header.Get(userTokenHeaderName)
	return GetCurrentUserMeta(tokenString)
}

func GetCurrentUserMeta(tokenString string) (*UserMeta, error) {
	if tokenString == "" {
		return nil, ErrUnauthorized
	}
	logrus.Debug("전달 받은 토큰", tokenString)

	token, err := parseToken(tokenString)
	if err != nil {
		logrus.Debug("전달 받은 토큰 에러", err)
		return nil, err
	}

	return getCurrentUserMetaByToken(token)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CircleCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKeys), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (c *UserBaseController) checkExistsObject(getExistsObjectFunc getExistsObjectInUserFunc, id uint, ErrRecordNotFoundMsg string) {
	exists, err := getExistsObjectFunc(id)
	if err != nil {
		c.ErrorAbort(500, err)
	}
	if !exists {
		c.ErrorAbort(400, errors.New(ErrRecordNotFoundMsg))
	}
}

func getCurrentUserMetaByToken(token *jwt.Token) (*UserMeta, error) {
	userMeta := &UserMeta{
		Token: token.Raw,
	}

	if claims, ok := token.Claims.(*CircleCustomClaims); ok && token.Valid {
		userMeta.UserID = claims.UserID
		userMeta.GroupID = claims.GroupID
	} else {
		return nil, errors.ErrInvalidToken
	}
	return userMeta, nil
}
