package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtKey = []byte("asdasdasdasdf")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	// 设置 token 过期时间为 10 天
	expirationTime := time.Now().Add(240 * time.Hour)
	// 创建 token 对象，并指定 token 中包含的信息
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// 使用 JWT 签名方法 HS256 签名，并使用指定的 secret key 生成 token 字符串
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		// 解析 token 字符串
		if ok, claims := CheckToken(tokenString); ok {
			c.Set("username", claims.Username)
			c.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "token无效"})
		}
	}
}

func CheckToken(tokenstring string) (bool, *Claims) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, nil
	}

	// 检查 token 是否有效，并获取其中的信息
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return false, nil
	}
	return true, claims
}
