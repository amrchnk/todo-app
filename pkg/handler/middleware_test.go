package handler

import (
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/amrchnk/todo-app/pkg/service"
    mock_service "github.com/amrchnk/todo-app/pkg/service/mocks"
    "net/http/httptest"
    "github.com/golang/mock/gomock"
    "fmt"
    "errors"
)

func TestHandler_userIdentity(t *testing.T){
    type mockBehavior func(s *mock_service.MockAuthorization,token string)

    testTable:=[]struct{
        name string
        headerName string
        headerValue string
        token string
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name: "OK",
            headerName: "Authorization",
            headerValue: "Bearer token",
            token:"token",
            mockBehavior: func(s *mock_service.MockAuthorization, token string){
                s.EXPECT().ParseToken(token).Return(1,nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:"1",
        },
        {
            name: "No header",
            headerName: "",
            mockBehavior: func(s *mock_service.MockAuthorization, token string){},
            expectedStatusCode:401,
            expectedRequestBody:`{"message":"empty auth header"}`,
        },
        {
            name: "Invalid Bearer",
            headerName: "Authorization",
            headerValue: "Bearr token",
            mockBehavior: func(s *mock_service.MockAuthorization, token string){},
            expectedStatusCode:401,
            expectedRequestBody:`{"message":"invalid auth header"}`,
        },
        {
            name: "Invalid token",
            headerName: "Authorization",
            headerValue: "Bearer ",
            mockBehavior: func(s *mock_service.MockAuthorization, token string){},
            expectedStatusCode:401,
            expectedRequestBody:`{"message":"token is empty"}`,
        },
        {
            name: "Service Failure",
            headerName: "Authorization",
            headerValue: "Bearer token",
            token:"token",
            mockBehavior: func(s *mock_service.MockAuthorization, token string){
                s.EXPECT().ParseToken(token).Return(1,errors.New("failed to parse token"))
            },
            expectedStatusCode:401,
            expectedRequestBody:`{"message":"failed to parse token"}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            auth:=mock_service.NewMockAuthorization(c)
            testCase.mockBehavior(auth,testCase.token)

            services:=&service.Service{Authorization:auth}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/protected",handler.userIdentity,func(c *gin.Context){
                id,_:=c.Get(userCtx)
                c.String(200,fmt.Sprintf("%d",id.(int)))
            })

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET","/protected",nil)
            req.Header.Set(testCase.headerName,testCase.headerValue)

            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}