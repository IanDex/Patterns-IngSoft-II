package strategy

import "fmt"

type CampaignContext struct {
	Name     string
	campaign CampaignStrategy
}

type CampaignStrategy interface {
	CreateCampaign(*CampaignContext)
}

func (c *CampaignContext) CreateCampaign() {
	c.campaign.CreateCampaign(c)
}

func NewCampaignContext(name string, campaign CampaignStrategy) *CampaignContext {
	return &CampaignContext{
		Name:     name,
		campaign: campaign,
	}
}

type CampaignType1 struct{}

func (*CampaignType1) CreateCampaign(ctx *CampaignContext) {
	fmt.Printf("Campaña Tipo 1: %s creada.\n", ctx.Name)
}

type CampaignType2 struct{}

func (*CampaignType2) CreateCampaign(ctx *CampaignContext) {
	fmt.Printf("Campaña Tipo 2: %s creada.\n", ctx.Name)
}
