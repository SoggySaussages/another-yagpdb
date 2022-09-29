package devtestexec

import (
	"context"

	"github.com/botlabs-gg/yagpdb/v2/bot"
	"github.com/botlabs-gg/yagpdb/v2/common/templates"
	"github.com/botlabs-gg/yagpdb/v2/customcommands/models"
	"github.com/botlabs-gg/yagpdb/v2/customcommands"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/volatiletech/sqlboiler/queries/qm"
)
// This is a commenter
var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "devtestexec",
	Description: "Executes a specified command",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	Arguments: []*dcmd.ArgDef{
		{Name: "CCID", Type: dcmd.Int},
		{Name: "Args", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		cmd, err := models.CustomCommands(qm.Where("guild_id = ? AND local_id = ?", data.GuildData.GS.ID, data.Args[0].Int())).OneG(context.Background())
		gs := bot.State.GetGuild(data.GuildData.GS.ID)
		cs := gs.GetChannel(data.ChannelID)
		ms := data.GuildData.MS
		tmplCtx := templates.NewContext(gs, cs, ms)
		tmplCtx.Data["SlashArgs"] = data.Args[1].Str
		return customcommands.ExecuteCustomCommand(cmd, tmplCtx, true)
	}}