package account_asset

import (
	"github.com/hexya-erp/pool/h"
)

//vars

var ()

//rights
func init() {
	h.AccountAssetCategory().Methods().Load().AllowGroup(GroupAccountUser)
	h.AccountAssetAsset().Methods().Load().AllowGroup(GroupAccountUser)
	h.AccountAssetCategory().Methods().AllowAllToGroup(GroupAccountManager)
	h.AccountAssetAsset().Methods().AllowAllToGroup(GroupAccountManager)
	h.AccountAssetDepreciationLine().Methods().Load().AllowGroup(GroupAccountUser)
	h.AccountAssetDepreciationLine().Methods().AllowAllToGroup(GroupAccountManager)
	h.AssetAssetReport().Methods().Load().AllowGroup(GroupAccountUser)
	h.AssetAssetReport().Methods().AllowAllToGroup(GroupAccountManager)
	h.AccountAssetCategory().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.AccountAssetAsset().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.AccountAssetAsset().Methods().Create().AllowGroup(GroupAccountInvoice)
	h.AccountAssetDepreciationLine().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.AccountAssetDepreciationLine().Methods().Create().AllowGroup(GroupAccountInvoice)
}
