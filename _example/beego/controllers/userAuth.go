package controllers

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jungju/circle_manager/_example/beegoapp/envs"
	"github.com/jungju/circle_manager/_example/beegoapp/errors"
	"github.com/jungju/circle_manager/_example/beegoapp/models"
	"github.com/jungju/circle_manager/_example/beegoapp/requests"
	"github.com/jungju/circle_manager/_example/beegoapp/responses"
	"github.com/jungju/circle_manager/modules"
)

const (
	userMapClaimsName  = "UserID"
	groupMapClaimsName = "GroupID"
)

//  UserAuthController operations for auth
type UserAuthController struct {
	modules.BaseController
}

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

func (c *UserAuthController) URLMapping() {
	//c.Mapping("Regist", c.Regist)
	c.Mapping("Login", c.Login)
}

// Login ...
// @Title Login
// @Description login User
// @Param	body		body 	requests.LoginUser	true		"body for token content"
// @Success 201 {int} responses.UserToken
// @Failure 403 body is empty
// @router /login [post]
func (c *UserAuthController) Login() {
	//Request Body 얻기 및 기본적인 유효성 체크
	reqBody := &requests.LoginUser{}
	c.SetRequestDataAndValid(reqBody)

	var userToken *responses.UserToken
	user, err := models.GetUserByUsername(reqBody.Username)
	if err == nil {
		userToken, err = login(user.ID, user.EncryptedPassword, reqBody.Password)
		if err != nil {
			c.ErrorAbort(401, nil)
		}
	} else {
		c.ErrorAbort(401, nil)
	}

	if userToken == nil {
		c.ErrorAbort(401, nil)
	}

	c.Success(http.StatusOK, userToken)
}

func login(userID uint, encryptedPassword string, reqBodyPassword string) (*responses.UserToken, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(reqBodyPassword)); err != nil {
		return nil, err
	}

	now := time.Now()

	claims := CircleCustomClaims{
		UserID:  userID,
		GroupID: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(5 * time.Hour).Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(envs.SecretKey))
	if err != nil {
		return nil, err
	}

	userToken := &responses.UserToken{
		Token:     tokenString,
		ExpiresAt: claims.ExpiresAt,
	}
	return userToken, nil
}

// Logout ...
// @Title Logout
// @Description logout User
// @Success 204 {int}
// @Failure 403 body is empty
// @router /logout [post]
func (c *UserAuthController) Logout() {
	tokenString := c.Ctx.Request.Header.Get(envs.UserTokenHeaderName)

	if userMeta, err := modules.GetCurrentUserMeta(tokenString); err == nil {
		_, err := models.GetUserByID(userMeta.UserID)
		if err != nil {
			c.ErrorAbort(401, nil)
		}
	} else {
		c.ErrorAbort(401, err)
	}

	c.ErrorAbort(204, nil)
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
