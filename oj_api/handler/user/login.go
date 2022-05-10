package user

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	md "github.com/wannabea/OnlineJudge/oj_api/middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func Login(c *gin.Context) {
	var loginReq LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := LoginCheck(c, loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败" + err.Error(),
				"data":   nil,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户数据解析失败",
			"data":   nil,
		})
	}
}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user UserIdt) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := md.NewJWT()

	// 构造用户claims信息(负荷)
	claims := md.CustomClaims{
		user.UserId,
		user.UserName,
		user.Email,
		user.IsAdmin,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "bgbiao.top",                    // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 封装一个响应数据,返回用户名和token
	data := LoginResult{
		Name:  user.UserName,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	return

}
