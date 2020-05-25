package adapter

import "fmt"

type ITargetFood interface {
	GetFoods() []string
}

type TargetFood struct {
}

func (t *TargetFood) GetFoods() []string {
	return []string{"Carne", "Pollo", "Hamburguesa"}
}

type Adaptee struct{}

func (a *Adaptee) GetDrinks() []string {
	return []string{"Vino", "Limonada", "Coca Cola"}
}

type ChefAdapter struct {
	adaptee *Adaptee
	foods   *TargetFood
}

func NewAdapter() *ChefAdapter {
	return &ChefAdapter{new(Adaptee), new(TargetFood)}
}

func (c *ChefAdapter) AdaptFoodsAndDrinks() {
	drinks := c.adaptee.GetDrinks()
	foods := c.foods.GetFoods()

	for i := 0; i < len(foods); i++ {
		fmt.Printf("Servir %s con %s\n", foods[i], drinks[i])
	}
}
