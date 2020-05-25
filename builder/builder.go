package builder

import "fmt"

type foodBuilder interface {
	GetFood() *Food
}

type ChefDirector struct {
}

type Food struct {
	Name       string
	TimeToCook int
	Materials  string
}

func (f *Food) Cook() {
	fmt.Printf("%s en %v minutos esta listo el/la %s \n", f.Materials, f.TimeToCook, f.Name)
}

//MixMaterials
func (d *ChefDirector) MixMaterials(builder foodBuilder) *Food {
	return builder.GetFood()
}

type SteakBuilder struct {
}

func (b SteakBuilder) GetFood() *Food {
	return &Food{Name: "Filete", TimeToCook: 20, Materials: "Carne, Cebolla, Arroz"}
}

type PastaBuilder struct {
}

func (b PastaBuilder) GetFood() *Food {
	return &Food{Name: "Pasta", TimeToCook: 30, Materials: "Tomate, Pasta, Carne, Cebolla"}
}
