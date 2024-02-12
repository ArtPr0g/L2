package pattern

/*
	Патерн использоваться для управления состоянием пользовательского интерфейса или поведения веб-приложения в зависимости от различных условий или событий.
*/


import (
    "fmt"
)

// Интерфейс Состояния
type OrderState interface {
    ConfirmOrder()
    CancelOrder()
}

type ConfirmedState struct{}

func (cs *ConfirmedState) ConfirmOrder() {
    fmt.Println("Заказ уже подтвержден.")
}

func (cs *ConfirmedState) CancelOrder() {
    fmt.Println("Отмена заказа.")
}

type CanceledState struct{}

func (cs *CanceledState) ConfirmOrder() {
    fmt.Println("Подтверждение заказа.")
}

func (cs *CanceledState) CancelOrder() {
    fmt.Println("Заказ уже отменен.")
}

// Структура заказа
type Order struct {
    state OrderState
}

func (o *Order) SetState(state OrderState) {
    o.state = state
}

func (o *Order) Confirm() {
    o.state.ConfirmOrder()
}

func (o *Order) Cancel() {
    o.state.CancelOrder()
}

func main() {
    order := &Order{state: &ConfirmedState{}}

    order.Confirm()
    order.Cancel()

    order.SetState(&CanceledState{})
    order.Confirm()
    order.Cancel()
}
