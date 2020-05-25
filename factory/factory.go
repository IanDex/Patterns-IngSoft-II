package factory

import "Patterns-IngSoft-II/strategy"

const (
	StringType1 string = "Ejemplo Tipo 1"
	StringType2 string = "Ejemplo Tipo 2"
)

func Factory(typeCampaign string, name string) (ctx *strategy.CampaignContext) {
	switch typeCampaign {
	case StringType1:
		return strategy.NewCampaignContext(name, &strategy.CampaignType1{})
	case StringType2:
		return strategy.NewCampaignContext(name, &strategy.CampaignType2{})
	}

	return nil
}
