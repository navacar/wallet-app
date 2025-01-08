package endpoint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type DB interface {
	Balance(id int) (float64, error)
	DepOrWithdraw(walletId int, operationType string, amount float64) (float64, error)
}

type Endpoint struct {
	db DB
}

type postRequestHandler struct {
	WalletId      int     `json:"walletId"`
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
}

func New(db DB) *Endpoint {
	return &Endpoint{
		db: db,
	}
}

func (e *Endpoint) Balance(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "No such waaletId")
		return err
	}

	balance, err := e.db.Balance(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error getting balance")
		return err
	}

	ctx.String(http.StatusOK, fmt.Sprintf("Balance is: %f", balance))
	return nil
}

func (e *Endpoint) DepOrWithdraw(ctx echo.Context) error {
	// walletId, err := strconv.Atoi(ctx.FormValue("walletId"))
	// if err != nil {
	// 	fmt.Println(ctx.FormValue("walletId"))
	// 	ctx.String(http.StatusBadRequest, "No such walletId")
	// 	return err
	// }

	// operType := ctx.FormValue("operationType")
	// amount, err := strconv.ParseFloat(ctx.FormValue("amount"), 64)
	// if err != nil || amount <= 0 {
	// 	ctx.String(http.StatusBadRequest, "Incorrect amount")
	// 	return err
	// }

	pR := &postRequestHandler{}
	if err := ctx.Bind(pR); err != nil {
		ctx.String(http.StatusBadRequest, "Invalid request body")
		return err
	}

	balance, err := e.db.DepOrWithdraw(pR.WalletId, pR.OperationType, pR.Amount)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error updating balance")
		return err
	}

	ctx.String(http.StatusOK, fmt.Sprintf("New balance is: %f", balance))
	return nil
}
