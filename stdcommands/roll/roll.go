package roll

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/dice"
)

var Command = &commands.YAGCommand{
	CmdCategory:     commands.CategoryFun,
	Name:            "devtestdatabase",
	Description:     "Testing database functions in commands.",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: tmplDBSet(ctx *templates.Context) interface{} {
			return func(userID int64, key interface{}, value interface{}) (string, error) {
				return (tmplDBSetExpire(ctx))(0, "devtest", 1, -1)
			}
		}
}
