package pages

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
)

func GetTools(ctx *context.Context) (types.Panel, error) {

	allBtns := make(types.Buttons, 0)

	act := action.Ajax("calcAll", func(ctx *context.Context) (success bool, msg string, data interface{}) {

		return true, "ok", nil
	})

	data := map[string]interface{}{
		"calc": "insterestAll",
	}

	// Add a ajax button action
	allBtns = append(allBtns, types.GetDefaultButton("生成利息还款计划", icon.Bank, act.SetUrl("/admin/tools/").SetData(data)))
	btns, btnsJs := allBtns.Content()
	cbs := make(types.Callbacks, 0)
	for _, btn := range allBtns {
		cbs = append(cbs, btn.GetAction().GetCallbacks())
	}
	btnJsStr := `<script>` + string(btnsJs) + `</script>`

	// 设置页面布局
	components := template.Default()
	btnContent := btns + template.HTML(btnJsStr)
	col1 := components.Col().SetContent(btnContent).SetSize(types.SizeXS(6).SM(3)).GetContent()

	row1 := components.Row().SetContent(col1).GetContent()

	return types.Panel{
		Content:     row1,
		Title:       "Tools",
		Description: "数据处理",
		Callbacks:   cbs,
	}, nil
}
