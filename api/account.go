package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rishitashaw/bank-api/database/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}


func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err:=ctx.ShouldBindJSON(&req); err!=nil{
		ctx.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}

	arg:=db.CreateAccountParams{
		Owner: req.Owner,
		Balance: 0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}
	ctx.JSON(http.StatusOK,account)
}

type getAccountRequest struct{
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){
	var req getAccountRequest

	if err:=ctx.ShouldBindUri(&req); err!=nil{
		ctx.JSON(http.StatusBadRequest, errorRequest(err))
		return
	}
	account, err:=server.store.GetAccount(ctx,req.ID)
	if err != nil{
		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorRequest(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorRequest(err))
		return
	}
	ctx.JSON(http.StatusOK,account)
}