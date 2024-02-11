package pattern

/*
	Этот патрн нужен для того, чтобы управлять выполнением различных действий, таких как обработка запросов пользователя, выполнение операций на сервере или управление динамическими действиями на стороне клиента.
*/
package main

import "fmt"

// Интерфейс команды
type Command interface {
    Execute()
}

// Конкретная команда
type ConcreteCommand struct {
    Receiver Receiver
}

func (cc *ConcreteCommand) Execute() {
    cc.Receiver.Action()
}

// Получатель команды
type Receiver struct{}

func (r *Receiver) Action() {
    fmt.Println("Выполняется действие получателя")
}

// Инициатор команды
type Invoker struct {
    commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
    i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
    for _, command := range i.commands {
        command.Execute()
    }
}

func main() {
    receiver := Receiver{}

    concreteCommand := ConcreteCommand{Receiver: receiver}

    invoker := Invoker{}
    invoker.StoreCommand(&concreteCommand)

    invoker.ExecuteCommands()
}
