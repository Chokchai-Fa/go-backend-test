package api

import (
	db "backend_test/db"
	"database/sql"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Tel   string `json:"tel" binding:"required"`
}

func createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.InsertNewUserParams{
		Name:  req.Name,
		Email: req.Email,
		Tel:   req.Tel,
	}

	res, err := db.CreateNewUser(arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type getAccountByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func getUserById(ctx *gin.Context) {
	var req getAccountByIdRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.GetUserByIdParams{
		ID: req.ID,
	}

	res, err := db.GetUserById(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func getAllUser(ctx *gin.Context) {
	res, err := db.GetAllUser(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type updateAllDetailRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Tel   string `json:"tel" binding:"required"`
}

func updateAllDetail(ctx *gin.Context) {
	var req updateAllDetailRequest

	paramID := ctx.Param("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAllDetailParams{
		ID:    int64(ID),
		Name:  req.Name,
		Email: req.Email,
		Tel:   req.Tel,
	}

	res, err := db.UpdateAllDetail(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)

}

type updateSomeDetailRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Tel   string `json:"tel"`
}

func updateSomeDetail(ctx *gin.Context) {
	var req updateSomeDetailRequest

	paramID := ctx.Param("id")

	ID, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Email != "" {
		_, err = mail.ParseAddress(req.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

	}

	arg := db.UpdateSomeDetailParams{
		ID:    int64(ID),
		Name:  req.Name,
		Email: req.Email,
		Tel:   req.Tel,
	}

	res, err := db.UpdateSomeDetail(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, res)

}

type deleteByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func deleteUserbyId(ctx *gin.Context) {
	var req deleteByIdRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteUserbyIdParams{
		ID: req.ID,
	}

	err := db.DeleteUserbyId(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
