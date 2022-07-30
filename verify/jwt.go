package verify

import (
	"time"
	"wechatdemo/model"

	"github.com/golang-jwt/jwt/v4"
)

//定义一个加密的密钥
var jwtKey = []byte("secretKey")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) { //发放Token
	expireTime := time.Now().Add(1048576 * time.Hour) //几乎永久过期时间
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //发放时间
			Issuer:    "web",
			Subject:   "user token",
		},
	}

	//获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //利用密钥生成token字符串

	if err != nil {
		return "获取失败", err
	}

	//返回字符串
	return tokenString, nil
}

//解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
