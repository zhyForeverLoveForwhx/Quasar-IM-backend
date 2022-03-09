package token

import (
	"demo/util"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32)) //生成密匙
	require.NoError(t, err)


	username := util.RandomUsername()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.NotZero(t, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T)  {
	maker, err := NewJWTMaker(util.RandomString(32)) //生成密匙
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomUsername(), -time.Minute)//生成负的时间
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t,err,ErrExpiredToken.Error())
	require.Nil(t,payload)
}

func TestInvaildJWTTokenAlgNone(t *testing.T)  {
	payload,err := NewPayload(util.RandomUsername(),time.Minute)
	require.NoError(t,err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)//需要应用TestInvaildJWTTokenAlgNone才能进行测试
	require.NoError(t,err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t,err)

	payload,err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}