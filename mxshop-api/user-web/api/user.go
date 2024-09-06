package api

import (
	"context"
	"fmt"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

var trans ut.Translator

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorErrors(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)), //翻译错误
	})
	return

}

func HandleGrpcError2Http(err error, c *gin.Context) {
	//将grpc code转化为http状态吗
	if err != nil {
		if e, ok := status.FromError(err); ok {
			//具体分析grpccode
			switch e.Code() {
			case codes.NotFound:
				//404
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				//500
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				//400
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误" + e.Message(),
				})
			}
			return
		}
	}
}

func GetUserList(ctx *gin.Context) {
	//从注册中心拉取到用户服务的信息

	claims, _ := ctx.Get("claims")
	zap.S().Infof("访问用户：%d", claims.(*models.CustomClaims).ID) //类型转换
	//生成grpc的client并调用接口

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	psize := ctx.DefaultQuery("psize", "5")
	pSizeInt, _ := strconv.Atoi(psize)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})

	if err != nil {
		zap.S().Errorw("GetUserlist 查询用户列表失败")
		HandleGrpcError2Http(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//data := make(map[string]interface{})

		user := response.UserResponse{
			Id:       value.Id,
			Nickname: value.Nickname,
			Birthday: time.Time(time.Unix(int64(value.Birthday), 0)).Format("2006-01-02"),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		//data["id"] = value.Id //符合go的编码规范，首字母转换成大写
		//data["name"] = value.Nickname
		//data["birthday"] = value.Birthday
		//data["gender"] = value.Gender
		//data["mobile"] = value.Mobile
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)

}

func PasswordLogin(c *gin.Context) {
	//表单验证
	passwordloginform := forms.PasswordLoginForm{}
	//var loginForm LoginForm
	if err := c.ShouldBind(&passwordloginform); err != nil { //参数绑定
		HandleValidatorErrors(c, err)
		return
	}

	//验证码校验,提交时，使用id做验证,每一次验证完关闭
	if !store.Verify(passwordloginform.CaptchaId, passwordloginform.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captchaId": "验证码错误",
		})
		return
	}

	//登陆的逻辑,手机号登陆
	if rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordloginform.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登陆失败",
				})
			}
			return

		}
	} else {
		//查询到用户，还需检查密码
		if PwdRsp, Pwderr := global.UserSrvClient.CheckUserPassword(context.Background(), &proto.PasswordCheck{
			Password:          passwordloginform.Password,
			Encryptedpassword: rsp.Password,
		}); Pwderr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登陆失败",
			})
		} else {
			if PwdRsp.Success {
				//生成token
				j := middlewares.NewJWT()
				//对指定的model进行签名
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.Nickname,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),              //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*7, //七天过期
						Issuer:    "xyyzmcsw",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nickname":   rsp.Nickname,
					"token":      token,
					"expired_at": time.Now().Unix() + 60*60*24*7,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登陆失败",
				})
			}
		}
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登录成功",
	//})
}

func RegisterUser(c *gin.Context) {
	//用户注册
	//验证码校验
	registerform := forms.RegisterForm{}
	if err := c.ShouldBind(&registerform); err != nil {
		HandleValidatorErrors(c, err)
		return
	}
	//校验验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	fmt.Printf("Redis Host: %s, Port: %d\n", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port)
	value, err := rdb.Get(context.Background(), registerform.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "key不存在",
		})
		return
	} else {
		if value != registerform.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":              "验证码错误",
				"value":             value,
				"registerform.Code": registerform.Code,
			})
			return
		}
	}
	//if err := c.ShouldBind(&registerform); err != nil {
	//	HandleValidatorErrors(c, err)
	//	return
	//}
	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Nickname: registerform.Mobile,
		Password: registerform.Password,
		Mobile:   registerform.Mobile,
	})

	if err != nil {
		zap.S().Errorw("Register 新建用户失败:%s", err.Error())
		HandleGrpcError2Http(err, c)
		return
	}

	j := middlewares.NewJWT()
	//对指定的model进行签名
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.Nickname,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),              //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //七天过期
			Issuer:    "xyyzmcsw",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nickname":   user.Nickname,
		"token":      token,
		"expired_at": time.Now().Unix() + 60*60*24*7,
	})
}
