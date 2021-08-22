package handler

import (
    "bytes"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/amrchnk/todo-app"
    "github.com/amrchnk/todo-app/pkg/service"
    mock_service "github.com/amrchnk/todo-app/pkg/service/mocks"
    "net/http/httptest"
    "github.com/golang/mock/gomock"
    "errors"
)

func TestHandler_signUp(t *testing.T){
    type mockBehavior func(s *mock_service.MockAuthorization,user todo.User)

    testTable:=[]struct{
        name string
        inputBody string
        inputUser todo.User
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            inputBody:`{"name":"Test","username":"test","password":"qwerty"}`,
            inputUser: todo.User{
                Name:"Test",
                Username:"test",
                Password:"qwerty",
            },
            mockBehavior: func(s *mock_service.MockAuthorization, user todo.User){
                s.EXPECT().CreateUser(user).Return(1,nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{"id":1}`,
        },
        {
            name:"Empty fields",
            inputBody:`{"username":"test","password":"qwerty"}`,
            mockBehavior: func(s *mock_service.MockAuthorization, user todo.User){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"invalid input body"}`,
        },
        {
            name:"Service Failure",
            inputBody:`{"name":"Test","username":"test","password":"qwerty"}`,
            inputUser: todo.User{
                Name:"Test",
                Username:"test",
                Password:"qwerty",
            },
            mockBehavior: func(s *mock_service.MockAuthorization, user todo.User){
                s.EXPECT().CreateUser(user).Return(1,errors.New("service failure"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"service failure"}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            auth:=mock_service.NewMockAuthorization(c)
            testCase.mockBehavior(auth,testCase.inputUser)

            services:=&service.Service{Authorization:auth}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.POST("/sign-up",handler.signUp)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("POST","/sign-up",
                bytes.NewBufferString(testCase.inputBody))

            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}