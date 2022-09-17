package roll

import (
	"github.com/botlabs-gg/yagpdb/v2/common"
	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/common/templates"
	"github.com/botlabs-gg/yagpdb/v2/customcommands"
)

var Command = &commands.YAGCommand{
	CmdCategory:     commands.CategoryFun,
	Name:            "devtestdatabase",
	Description:     "Testing database functions in commands.",
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {

	func(userID int64, key interface{}, incrBy interface{}) (interface{}, error) {
		vNum := templates.ToFloat64(1)
		var b bytes.Buffer
		enc := msgpack.NewEncoder(templates.LimitWriter(&b, 100000))
		err := enc.Encode(vNum)
		valueSerialized, err := b.Bytes()
		if err != nil {
			return "", err
		}

		keyStr := "devtest"

		const q = `INSERT INTO templates_user_database (created_at, updated_at, guild_id, user_id, key, value_raw, value_num)
VALUES (now(), now(), $1, $2, $3, $4, $5)
ON CONFLICT (guild_id, user_id, key)
DO UPDATE SET
	value_num =
		-- Don't increment expired entry
		CASE WHEN (templates_user_database.expires_at IS NULL OR templates_user_database.expires_at > now()) THEN templates_user_database.value_num + $5
		ELSE $5
		END,
	updated_at = now(),
	created_at =
		-- Reset created_at if the entry expired
		CASE WHEN (templates_user_database.expires_at IS NULL OR templates_user_database.expires_at > now()) THEN templates_user_database.created_at
		ELSE now()
		END,
	expires_at =
		-- Same for expires_at
		CASE WHEN (templates_user_database.expires_at IS NULL OR templates_user_database.expires_at > now()) THEN templates_user_database.expires_at
		ELSE NULL
		END

RETURNING value_num`

		result := common.PQ.QueryRow(q, 1019386105063284828, 0, keyStr, valueSerialized, vNum)

		var newVal float64
		err = result.Scan(&newVal)
		return newVal, err
	},
}
}