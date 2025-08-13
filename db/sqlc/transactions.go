package db

import (
	"context"
	"fmt"
)

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// Make transfer, make two entries, update accounts balance
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry from")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry to")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			fmt.Println(txName, "get account from")
			result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
			if err != nil {
				return err
			}
			fmt.Println(txName, "get account to")
			result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}

			fmt.Println(txName, "update account from")
			result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.FromAccountID,
				Balance: result.FromAccount.Balance - arg.Amount,
			})
			if err != nil {
				return err
			}

			fmt.Println(txName, "update account to")
			result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.ToAccountID,
				Balance: result.ToAccount.Balance + arg.Amount,
			})
			if err != nil {
				return err
			}

		} else {
			fmt.Println(txName, "get account to")
			result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
			if err != nil {
				return err
			}
			fmt.Println(txName, "get account from")
			result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
			if err != nil {
				return err
			}

			fmt.Println(txName, "update account to")
			result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.ToAccountID,
				Balance: result.ToAccount.Balance + arg.Amount,
			})
			if err != nil {
				return err
			}

			fmt.Println(txName, "update account from")
			result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
				ID:      arg.FromAccountID,
				Balance: result.FromAccount.Balance - arg.Amount,
			})
			if err != nil {
				return err
			}

		}

		return nil
	})

	return result, err
}
