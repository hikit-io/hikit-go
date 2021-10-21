// @title Atlas API
// @version 1.0

// @contact.name Nekilc

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hfunc/hfunc-go/hfctx"
	"github.com/hfunc/hfunc-go/hflog"
	"github.com/hfunc/hfunc-go/hfutils/hfgin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type _Controller struct{}

func (_ _Controller) GroupName() string {
	return "hfunc"
}

type logger struct{}

func (l logger) Info(ctx hfctx.Ctx, msg string, keyAndValues ...interface{}) {
	hflog.Info(ctx, msg, keyAndValues...)
}

func (_ _Controller) Middlewares() (prefix, suffix []gin.HandlerFunc) {
	return []gin.HandlerFunc{
			hfgin.Tracer("trace_id"),
			hfgin.Logger(logger{}),
			hfgin.Paramer(logger{}),
			func(c *gin.Context) {
				hflog.Info(c, "global prefix")
				c.Next()
			},
		}, []gin.HandlerFunc{
			func(c *gin.Context) {
				hflog.Info(c, "global subfix")
				c.Next()
			},
		}
}

type GetUserByIdParams struct {
}

type GetUserByIdResp struct {
}

// GetUserById godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param query query GetUserByIdParams true " "
// @Success 200 {object} GetUserByIdResp
// @Router /v1/hfunc/user/:id [Get]
// @Security ApiKeyAuth
func (a _Controller) GetUserById() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &GetUserByIdParams{}
				resp   = &GetUserByIdResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			id := c.Param("id")
			hflog.Info(c, "GetUserById", zap.String("id", id))
			hfgin.Ok(c, resp)
		},
	}
}

type PostUserByIdParams struct {
}

type PostUserByIdResp struct {
}

// PostUserById godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param query query PostUserByIdParams true " "
// @Success 200 {object} PostUserByIdResp
// @Router /v1/hfunc/user/:id [Post]
// @Security ApiKeyAuth
func (a _Controller) PostUserById() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &PostUserByIdParams{}
				resp   = &PostUserByIdResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
		func(c *gin.Context) {
			hflog.Info(c, "subfix")
			c.Next()
		},
	}
}

type PutUserByIdParams struct {
}

type PutUserByIdResp struct {
}

// PutUserById godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param query query PutUserByIdParams true " "
// @Success 200 {object} PutUserByIdResp
// @Router /v1/hfunc/user/:id [Put]
// @Security ApiKeyAuth
func (c _Controller) PutUserById() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &PutUserByIdParams{}
				resp   = &PutUserByIdResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
		func(c *gin.Context) {
			hflog.Info(c, "subfix")
			c.Next()
		},
	}
}

type DeleteUserByIdParams struct {
}

type DeleteUserByIdResp struct {
}

// DeleteUserById godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param body body DeleteUserByIdParams true " "
// @Success 200 {object} DeleteUserByIdResp
// @Router /v1/hfunc/user/:id [Delete]
// @Security ApiKeyAuth
func (c _Controller) DeleteUserById() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &DeleteUserByIdParams{}
				resp   = &DeleteUserByIdResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
	}
}

type PatchUserByIdParams struct {
}

type PatchUserByIdResp struct {
}

// PatchUserById godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param body body PatchUserByIdParams true " "
// @Success 200 {object} PatchUserByIdResp
// @Router /v1/hfunc/user/:id [Patch]
// @Security ApiKeyAuth
func (c _Controller) PatchUserById() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &PatchUserByIdParams{}
				resp   = &PatchUserByIdResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
	}
}

type PostUserByIdNameParams struct {
}

type PostUserByIdNameResp struct {
}

// PostUserByIdName godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param body body PostUserByIdNameParams true " "
// @Success 200 {object} PostUserByIdNameResp
// @Router /v1/hfunc/user/:id/name [Post]
// @Security ApiKeyAuth
func (c _Controller) PostUserByIdName() (version string, handlers []gin.HandlerFunc) {
	return "v`", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &PostUserByIdNameParams{}
				resp   = &PostUserByIdNameResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
	}
}

type GetUserListParams struct {
}

type GetUserListResp struct {
}

// GetUserList godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param query query GetUserListParams true " "
// @Success 200 {object} GetUserListResp
// @Router /v1/hfunc/user/list [Get]
// @Security ApiKeyAuth
func (c _Controller) GetUserList() (version string, handlers []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			var (
				err error

				params = &GetUserListParams{}
				resp   = &GetUserListResp{}
			)
			if err = c.ShouldBind(params); err != nil {

				return
			}
			hfgin.Ok(c, resp)
		},
	}
}

func main() {
	r := gin.New()
	i := r.Group("/", gin.Recovery())
	i.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	hfgin.RegisterController(i, &_Controller{})
	r.Run(":8081")
}
