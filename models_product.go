package account_asset

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProductTemplate().DeclareModel()

	h.ProductTemplate().AddFields(map[string]models.FieldDefinition{
		"AssetCategoryId": models.Many2OneField{
			RelationModel: h.AccountAssetCategory(),
			String:        "Asset Type",
			//company_dependent=True
			OnDelete: `restrict`,
		},
		"DeferredRevenueCategoryId": models.Many2OneField{
			RelationModel: h.AccountAssetCategory(),
			String:        "Deferred Revenue Type",
			//company_dependent=True
			OnDelete: `restrict`,
		},
	})
	h.ProductTemplate().Methods().GetAssetAccounts().DeclareMethod(
		`GetAssetAccounts`,
		func(rs m.ProductTemplateSet) {
			//        res = super(ProductTemplate, self)._get_asset_accounts()
			//        if self.asset_category_id:
			//            res['stock_input'] = self.property_account_expense_id
			//        if self.deferred_revenue_category_id:
			//            res['stock_output'] = self.property_account_income_id
			//        return res
		})
}
