package devtestexec

import (
	"context"

	"github.com/botlabs-gg/yagpdb/v2/bot"
	"github.com/botlabs-gg/yagpdb/v2/common/templates"
	"github.com/botlabs-gg/yagpdb/v2/customcommands/models"
	"github.com/botlabs-gg/yagpdb/v2/customcommands"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)
// This is a commenter
var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "devtestexec",
	Description: "Executes a specified command",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RequiredArgs:  2,
	Arguments: []*dcmd.ArgDef{
		{Name: "CCID", Type: dcmd.Int},
		{Name: "Args", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		cmd, err := models.CustomCommands(qm.Where("guild_id = ? AND local_id = ?", data.GuildData.GS.ID, data.Args[0].Int())).OneG(context.Background())
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		gs := bot.State.GetGuild(data.GuildData.GS.ID)
		cs := gs.GetChannel(data.ChannelID)
		ms := data.GuildData.MS
		tmplCtx := templates.NewContext(gs, cs, ms, fmt.Sprintf("%d;;%s;;%s;;%d;;%d;;%s", 0, data.Args[1].Str(), data.SlashCommandTriggerData.Interaction.Token, data.SlashCommandTriggerData.Interaction.ID, 0, ""))
		body, _ := customcommands.ExecuteCustomCommand(cmd, tmplCtx, true)
		return body, nil
	}}//