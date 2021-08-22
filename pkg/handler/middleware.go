//interlayer for parsing tokens from requests and providing necessary endpoints
package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
    "errors"
)

const (
    authorizationHeader="Authorization"
    userCtx="userId"
)

//get and validate user token
func (h *Handler) userIdentity(c *gin.Context){
    header:=c.GetHeader(authorizationHeader)
    if header==""{
        newErrorResponse(c,http.StatusUnauthorized,"empty auth header")
        return
    }

    headerParts:=strings.Split(header," ")
    if len(headerParts)!=2{
        newErrorResponse(c,http.StatusUnauthorized,"invalid auth header")
        return
    }

    if headerParts[0]!="Bearer"{
        newErrorResponse(c,http.StatusUnauthorized,"invalid auth header")
        return
    }

    if headerParts[1]==""{
        newErrorResponse(c,http.StatusUnauthorized,"token is empty")
        return
    }

    userId,err:=h.services.Authorization.ParseToken(headerParts[1])
    if err!=nil{
        newErrorResponse(c,http.StatusUnauthorized,err.Error())
        return
    }

    c.Set(userCtx,userId)
}

func getUserId(c *gin.Context)(int,error){
    id,ok:=c.Get(userCtx)
    if !ok{
        newErrorResponse(c,http.StatusInternalServerError,"user id not found")
        return 0,errors.New("user isn't found")
    }
    idInt,ok:=id.(int)
    if !ok{
        newErrorResponse(c,http.StatusInternalServerError,"user id isn't valid")
        return 0,errors.New("user isn't found")
    }

    return idInt,nil
}
