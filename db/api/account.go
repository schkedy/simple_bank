package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

type getAccountRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

type getListAccountsRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"`
}

type deleteAccountRequest struct {
	Id int64 `uri:"id" binidng:"required,min=1"`
}

// type UpdateAccountRequest struct {
// 	Id int64 ``
// }

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id := req.Id

	account, err := server.store.GetAccount(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// account := db.Account{}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		fmt.Println(err)
		return
	}

	fmt.Println("ff")
	args := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)

}

func (server *Server) getListAccounts(ctx *gin.Context) {
	var req getListAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.GetListAccountsParams{
		Limit:  int64(req.PageSize),
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.GetListAccounts(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	id := req.Id
	result, err := server.store.DeleteAccount(ctx, id)
	fmt.Print(err)
	if err != nil {
		fmt.Print(err)
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
