package pattern

package main

import "fmt"

// Шаги построения дома
type HouseBuilder interface {
	SetWalls() HouseBuilder
	SetRoof() HouseBuilder
	SetWindows() HouseBuilder
	Build() House
}

// Дом
type House struct {
	walls   string
	roof    string
	windows string
}

// Рабочий для строительства дома
type ConstructionWorker struct {
	houseBuilder HouseBuilder
}

// Установка стен дома
func (c *ConstructionWorker) SetWalls() HouseBuilder {
	c.houseBuilder.walls = "Стены установлены"
	return c.houseBuilder
}

// Установка крыши дома
func (c *ConstructionWorker) SetRoof() HouseBuilder {
	c.houseBuilder.roof = "Крыша установлена"
	return c.houseBuilder
}

// Установка окон дома
func (c *ConstructionWorker) SetWindows() HouseBuilder {
	c.houseBuilder.windows = "Окна установлены"
	return c.houseBuilder
}

// Построить дом
func (c *ConstructionWorker) Build() House {
	return House{
		walls:   c.houseBuilder.walls,
		roof:    c.houseBuilder.roof,
		windows: c.houseBuilder.windows,
	}
}

func main() {
	// Создаем строителя
	houseBuilder := &HouseBuilderImpl{}

	// Создаем рабочего для строительства дома
	constructionWorker := &ConstructionWorker{
		houseBuilder: houseBuilder,
	}

	// Строим дом
	house := constructionWorker.
		SetWalls().
		SetRoof().
		SetWindows().
		Build()

	// Выводим информацию о построенном доме
	fmt.Printf("Дом построен:\nСтены: %s\nКрыша: %s\nОкна: %s\n",
		house.walls, house.roof, house.windows)
}
