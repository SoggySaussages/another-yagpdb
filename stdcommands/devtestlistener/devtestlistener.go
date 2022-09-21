package devtestlistener

import (
	"net/http"
)
// This is a commenter
var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "devtestlistener",
	Description: "Queries a test listener",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {

	resp, err := http.Get("http://127.0.0.1:9000/")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err

	}}