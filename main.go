package main

import (
	"Patterns-IngSoft-II/builder"
	"Patterns-IngSoft-II/factory"
	"Patterns-IngSoft-II/prototype"

	"Patterns-IngSoft-II/adapter"
	"Patterns-IngSoft-II/bridge"
	"Patterns-IngSoft-II/facade"

	"Patterns-IngSoft-II/observer"
	s "Patterns-IngSoft-II/strategy"
	"fmt"
)

func main() {

	fmt.Printf("---------------------------\n")

	fmt.Printf("Prototype Pattern:\n")
	customer := prototype.NewCustomer("Cristtian", "ian@gmail.com")
	fmt.Printf("Original Object, Name: %v, Email: %v\n", customer.Name, customer.Email)

	customer2 := customer.WithName("Jared")
	fmt.Printf("Cloned Object, New Customer name: %v, Email: %v\n", customer2.Name, customer2.Email)

	fmt.Printf("Original Object, Customer name: %v, Email: %v\n", customer.Name, customer.Email)

	fmt.Printf("---------------------------\n")
	fmt.Printf("Builder Pattern:\n")

	chef := builder.ChefDirector{}

	_steakBuilder := builder.SteakBuilder{}
	food := chef.MixMaterials(_steakBuilder)
	food.Cook()

	_pastaBuilder := builder.PastaBuilder{}
	food = chef.MixMaterials(_pastaBuilder)
	food.Cook()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Factory Pattern:\n")
	factory.Factory(factory.StringType1, "Capital_Campaña").CreateCampaign()
	factory.Factory(factory.StringType2, "Campaña Ejemplo").CreateCampaign()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Adapter Pattern:\n")
	a := adapter.NewAdapter()
	a.AdaptFoodsAndDrinks()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Bridge Pattern:\n")

	app := bridge.App{
		CurrentUser: 20,
	}
	saleApp := bridge.SaleApp{}
	app.AppContext = &saleApp

	app.GetTasks()

	adminApp := bridge.AdminApp{}

	app.CurrentUser = 30
	app.AppContext = &adminApp

	app.GetAccount()

	fmt.Printf("---------------------------\n")

	fmt.Printf("Facade Pattern:\n")
	f := facade.NewFacade()
	f.CompleteProcess("Microsoft", "Microsoft Azul", "10 tips abouts Microsoft Azul")

	fmt.Printf("---------------------------\n")

	fmt.Printf("Iterator Pattern:\n")
	Iterate(func(n int) { fmt.Println(n) })

	fmt.Printf("---------------------------\n")

	fmt.Printf("Strategy Pattern:\n")
	ctx := s.NewCampaignContext("Karl Marx", &s.CampaignType1{})
	ctx.CreateCampaign()

	ctx = s.NewCampaignContext("Hannah Arendt", &s.CampaignType2{})
	ctx.CreateCampaign()

	fmt.Printf("Observer Pattern:\n")

	subject := observer.NewSubject()
	concreteObserverA := observer.NewConcreteNotifierA()
	concreteObserverA1 := observer.NewConcreteNotifierA()
	concreteObserverB := observer.NewConcreteNotifierB()
	concreteObserverB1 := observer.NewConcreteNotifierB()

	subject.Attach(concreteObserverA)
	subject.Attach(concreteObserverA1)
	subject.Attach(concreteObserverB)
	subject.Attach(concreteObserverB1)

	fmt.Printf("First Scan\n")
	subject.ScanData()

	subject.DeAttach(concreteObserverB)
	subject.DeAttach(concreteObserverA)

	fmt.Printf("Second Scan\n")
	subject.ScanData()

	fmt.Printf("---------------------------\n")
}

func Iterate(f func(n int)) {
	for i := 1; i <= 3; i++ {
		f(i)
	}
}
