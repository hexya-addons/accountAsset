package account_asset

import (
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "account_asset"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/account_asset/src/js/account_asset.js",
	)
	controllers.Backend = append(controllers.Backend,
		"/static/account_asset/src/less/account_asset.less",
	)

}
