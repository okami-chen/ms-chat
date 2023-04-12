package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/okamin-chen/chat/pkg/config"
	"github.com/okamin-chen/chat/pkg/global"
	"net/http"
	"reflect"
)

type AuthController struct {
}

type AuthWechatRequest struct {
	Code string `json:"code" form:"code" binding:"required"`
}

func StructToMap(obj any) (data map[string]interface{}, err error) {
	// 通过反射将结构体转换成map
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			data[fileName] = objV.Field(i).Interface()
		} else {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data, nil
}

func (auth AuthController) LoginByWechat(c *gin.Context) {

	var request AuthWechatRequest
	if c.ShouldBind(&request) != nil {
		c.JSON(401, gin.H{"status": "missing params"})
	}
	ctx := context.TODO()
	response, e := global.MiniProgram.Auth.Session(ctx, request.Code)

	if e != nil {

		c.JSON(http.StatusOK, gin.H{
			"message": "授权失败",
		})
		return
	}
	if response.ErrCode > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    4023,
			"message": response,
		})
		return
	}

	claims := config.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID: response.OpenID,
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"token": global.Conf.Jwt.Generate(claims),
	})
}
