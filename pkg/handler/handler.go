package handler

import (
    "github.com/gin-gonic/gin"
    "github.com/amrchnk/todo-app/pkg/service"
    cors "github.com/rs/cors/wrapper/gin"
)

type Handler struct{
    services *service.Service
}

func NewHandler(services *service.Service) *Handler{
    return &Handler{services:services}
}

func (h *Handler) InitRoutes() *gin.Engine{
    router:=gin.New()
    router.Use(cors.Default())
    //endpoints for authorization and registration pages
    auth:=router.Group("/auth")
    {
        auth.POST("/sign-up",h.signUp)
        auth.POST("/sign-in",h.signIn)
    }

    //endpoints for working with lists and todos
    api:=router.Group("/api",h.userIdentity)
    {
        //lists routes
        lists:=api.Group("/lists")
        {
            lists.POST("/",h.createList)
            lists.GET("/",h.getAllLists)
            lists.GET("/:id",h.getListById)
            lists.PUT("/:id",h.updateList)
            lists.DELETE("/:id",h.deleteList)

            //items routes
            items:=lists.Group(":id/items")
            {
                items.POST("/",h.createItem)
                items.GET("/",h.getAllItems)
            }
        }

        items:=api.Group("items")
        {
            items.GET("/:id",h.getItemById)
            items.PUT("/:id",h.updateItem)
            items.DELETE("/:id",h.deleteItem)
        }
    }

    return router
}