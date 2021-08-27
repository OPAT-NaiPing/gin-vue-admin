package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify"
	//email "github.com/flipped-aurora/gva-plug-email"   // 在线仓库模式
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email" // 本地插件仓库地址模式
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/example_plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
	//  添加开放权限的插件 示例
	PluginInit(PublicGroup, example_plugin.ExamplePlugin)

	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

	//  钉钉通知，暂时开放权限
	PluginInit(PublicGroup, notify.CreateDDPlug(
		"https://oapi.dingtalk.com/robot/send",
		"8ded23f91917dc4f6275f44ba5ef243e6ed1d2cc74de83f01a6f5f5f39905671",
		"SECaecf452bd6e671ab0d47469c3ad933e32fcc47b335333049a1b8961606192f38"))
}
