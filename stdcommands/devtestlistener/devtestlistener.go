package devtestlistener

import (
//	"net/http"
//	"io"
//	"fmt"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
)
// This is a commenter
var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "devtestlistener",
	Description: "Queries a test listener",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	Arguments: []*dcmd.ArgDef{
		{Name: "UserID", Type: dcmd.BigInt},
		{Name: "Key", Type: dcmd.String},
	},
	RunFunc: func(data *dcmd.Data) (interface{}, error) {

//	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:9000/%d/%s", data.Args[0].Int64(), data.Args[1].Str())
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	
//	if string(body) = "none" {
//		return "No entry found", err
//	}

//	return string(body), err
return "this command is on hold for a sec, i'll work on it closer to testing environment", nil
	}}