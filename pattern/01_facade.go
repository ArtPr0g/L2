package pattern

import "fmt"
/*
Допустим, у нас есть следующие компоненты банковской системы:

Авторизация пользователя
Проверка баланса
Перевод денег
Создадим интерфейсы для каждого из этих компонентов:
*/

type Authorization interface {
    Authorize(username string, password string) bool
}

type AccountBalanceChecker interface {
    GetAccountBalance(accountNumber string) float64
}

type MoneyTransferer interface {
    TransferMoney(senderAccountNumber string, receiverAccountNumber string, amount float64) bool
}
// Затем реализуем конкретные классы для каждого из интерфейсов:

go
Copy code
type AuthorizationService struct {
    // реализация интерфейса Authorization
}

func (a *AuthorizationService) Authorize(username string, password string) bool {
    // реализация авторизации пользователя
}

type AccountBalanceCheckerService struct {
    // реализация интерфейса AccountBalanceChecker
}

func (a *AccountBalanceCheckerService) GetAccountBalance(accountNumber string) float64 {
    // реализация получения баланса
}

type MoneyTransferService struct {
    // реализация интерфейса MoneyTransferer
}

func (m *MoneyTransferService) TransferMoney(senderAccountNumber string, receiverAccountNumber string, amount float64) bool {
    // реализация перевода денег
}
// Теперь создадим структуру Фасада, которая будет предоставлять простой интерфейс для работы с банковской системой:

go
Copy code
type BankSystemFacade struct {
    authorization     Authorization
    balanceChecker    AccountBalanceChecker
    moneyTransferer   MoneyTransferer
}

func NewBankSystemFacade() *BankSystemFacade {
    return &BankSystemFacade{
        authorization:     &AuthorizationService{},
        balanceChecker:    &AccountBalanceCheckerService{},
        moneyTransferer:   &MoneyTransferService{},
    }
}

func (f *BankSystemFacade) Login(username string, password string) bool {
    return f.authorization.Authorize(username, password)
}

func (f *BankSystemFacade) CheckBalance(accountNumber string) float64 {
    return f.balanceChecker.GetAccountBalance(accountNumber)
}

func (f *BankSystemFacade) TransferMoney(senderAccountNumber string, receiverAccountNumber string, amount float64) bool {
    return f.moneyTransferer.TransferMoney(senderAccountNumber, receiverAccountNumber, amount)
}
//Теперь можем использовать Фасад для упрощенного взаимодействия с банковской системой:

go
Copy code
func main() {
    facade := NewBankSystemFacade()

    authorized := facade.Login("username", "password")
    if authorized {
        balance := facade.CheckBalance("accountNumber")
        fmt.Println("Balance:", balance)

        transferred := facade.TransferMoney("senderAccountNumber", "receiverAccountNumber", 100.0)
        if transferred {
            fmt.Println("Перевод денег выполнен")
        } else {
            fmt.Println("Ошибка перевода денег")
        }
    } else {
        fmt.Println("Ошибка авторизации")
    }
}
//Таким образом, мы использовали паттерн Фасад для создания упрощенного интерфейса для работы с банковской системой, скрывая сложность ее компонентов от клиента.
