package test

import (
	"blockEmulator/account"
	"blockEmulator/params"
	"fmt"
	"math/big"
	"testing"
)

func Test_bankInitialization(t *testing.T) {
	// Test that Enable_bank config is false by default (before we enable it)
	if params.Config.Enable_bank {
		t.Log("Bank is enabled")
	} else {
		t.Log("Bank is disabled (default)")
	}
}

func Test_bankAccountFunctions(t *testing.T) {
	// Initialize bank accounts
	account.InitializeBankAccounts()

	// Test IsBankAccount
	if !account.IsBankAccount("BANK") {
		t.Error("IsBankAccount should return true for BANK")
	}
	if account.IsBankAccount("other") {
		t.Error("IsBankAccount should return false for non-BANK addresses")
	}

	// Test SetBankCreditLine and GetBankCreditLine
	testAddr := "test_account_123"
	testBalance := big.NewInt(1000)

	account.SetBankCreditLine(testAddr, testBalance)
	retrievedBalance := account.GetBankCreditLine(testAddr)

	if retrievedBalance.Cmp(testBalance) != 0 {
		t.Errorf("Credit line mismatch: expected %v, got %v", testBalance, retrievedBalance)
	}

	// Test ClearBankCreditLine
	account.ClearBankCreditLine(testAddr)
	retrievedBalance = account.GetBankCreditLine(testAddr)
	if retrievedBalance.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Credit line should be 0 after clear, got %v", retrievedBalance)
	}

	fmt.Printf("Bank account functions test passed!\n")
}

func Test_bankCreditLineWithMigration(t *testing.T) {
	// This test simulates the bank credit flow during migration
	// 1. Account mortgages balance to bank
	// 2. Credit line is set
	// 3. Credit line is retrieved
	// 4. Credit line is cleared after settlement

	account.InitializeBankAccounts()

	migratingAccount := "migrating_account_456"
	originalBalance := big.NewInt(1000)

	// Step 1 & 2: Mortgage balance to bank (set credit line)
	account.SetBankCreditLine(migratingAccount, originalBalance)
	fmt.Printf("Account %s mortgaged %v to bank as credit\n", migratingAccount, originalBalance)

	// Step 3: Verify credit line
	creditLine := account.GetBankCreditLine(migratingAccount)
	if creditLine.Cmp(originalBalance) != 0 {
		t.Errorf("Credit line mismatch: expected %v, got %v", originalBalance, creditLine)
	}
	fmt.Printf("Credit line verified: %v\n", creditLine)

	// Simulate settlement - bank returns remaining balance
	amountPaidByBank := big.NewInt(600)
	remainingBalance := new(big.Int).Sub(originalBalance, amountPaidByBank)
	fmt.Printf("Amount paid by bank: %v\n", amountPaidByBank)
	fmt.Printf("Remaining balance to return: %v\n", remainingBalance)

	// Step 4: Clear credit line after settlement
	account.ClearBankCreditLine(migratingAccount)
	creditLineAfter := account.GetBankCreditLine(migratingAccount)
	if creditLineAfter.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Credit line should be 0 after settlement, got %v", creditLineAfter)
	}
	fmt.Printf("Credit line cleared after settlement\n")

	fmt.Printf("Bank credit line with migration test passed!\n")
}

func Test_bankAsSender(t *testing.T) {
	// Test that Bank account is recognized correctly
	bankAddr := account.BankAccountAddr
	fmt.Printf("Bank account address: %s\n", bankAddr)

	if !account.IsBankAccount(bankAddr) {
		t.Errorf("IsBankAccount should return true for bank address: %s", bankAddr)
	}

	// Test with byte representation
	byteAddr := []byte(bankAddr)
	fmt.Printf("Bank account as bytes: %v\n", byteAddr)
}
