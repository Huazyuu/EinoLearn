package stage05tools

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type Game struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type InputParams struct {
	Name string `json:"name" jsonschema:"description=the name of game"`
}

func GetGame(_ context.Context, params *InputParams) (string, error) {
	GameSet := []Game{
		{Name: "原神", Url: "https://ys.mihoyo.com/tool"},
		{Name: "鸣潮", Url: "https://mc.kurogames.com/tool"},
		{Name: "明日方舟", Url: "https://ak.hypergryph.com/tool"},
	}
	for _, g := range GameSet {
		if g.Name == params.Name {
			return g.Url, nil
		}
	}
	return "", fmt.Errorf("game not found")
}

func CreateTool() tool.InvokableTool {
	getGameTool := utils.NewTool(&schema.ToolInfo{
		Name: "getGame",
		Desc: "get the url of game",
		ParamsOneOf: schema.NewParamsOneOfByParams(
			map[string]*schema.ParameterInfo{
				"name": &schema.ParameterInfo{
					Type:     schema.String,
					Desc:     "game's name",
					Required: true,
				},
			},
		),
	}, GetGame)
	return getGameTool
}
