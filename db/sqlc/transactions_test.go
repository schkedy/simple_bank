package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 3
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)
	// pomp := make(chan interface{}, 1)
	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			// pomp <- true
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// transfer := result.Transfer
		// require.NotEmpty(t, transfer)
		// require.Equal(t, account1.ID, transfer.FromAccountID)
		// require.Equal(t, account2.ID, transfer.ToAccountID)
		// require.Equal(t, amount, transfer.Amount)
		// require.NotZero(t, transfer.ID)
		// require.NotZero(t, transfer.CreatedAt)

		// _, err = store.GetTransfer(context.Background(), transfer.ID)
		// require.NoError(t, err)

		// // check entries
		// fromEntry := result.FromEntry
		// require.NotEmpty(t, fromEntry)
		// require.Equal(t, account1.ID, fromEntry.AccountID)
		// require.Equal(t, -amount, fromEntry.Amount)
		// require.NotZero(t, fromEntry.ID)
		// require.NotZero(t, fromEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), fromEntry.ID)
		// require.NoError(t, err)

		// // check to entry
		// toEntry := result.ToEntry
		// require.NotEmpty(t, toEntry)
		// require.Equal(t, account2.ID, toEntry.AccountID)
		// require.Equal(t, amount, toEntry.Amount)
		// require.NotZero(t, toEntry.ID)
		// require.NotZero(t, toEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), toEntry.ID)
		// require.NoError(t, err)

		// // check accounts
		fromAccount := result.FromAccount
		// require.NotEmpty(t, fromAccount)
		// require.Equal(t, account1.ID, fromAccount.ID)
		// // require.Equal(t, account1.Balance-int64(i+1)*amount, fromAccount.Balance)

		toAccount := result.ToAccount
		// require.NotEmpty(t, toAccount)
		// require.Equal(t, account2.ID, toAccount.ID)
		// // require.Equal(t, account2.Balance+int64(i+1)*amount, toAccount.Balance)

		fmt.Println(">> after:", fromAccount.Balance, toAccount.Balance)

		// <-pomp
	}
}
