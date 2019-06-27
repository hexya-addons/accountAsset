package account_asset

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AssetDepreciationConfirmationWizard().DeclareModel()

	h.AssetDepreciationConfirmationWizard().AddFields(map[string]models.FieldDefinition{
		"Date": models.DateField{
			String:   "Account Date",
			Required: true,
			Help: "Choose the period for which you want to automatically post" +
				"the depreciation lines of running assets",
			Default: func(env models.Environment) interface{} { return odoo.fields.Date.context_today },
		},
	})
	h.AssetDepreciationConfirmationWizard().Methods().AssetCompute().DeclareMethod(
		`AssetCompute`,
		func(rs m.AssetDepreciationConfirmationWizardSet) {
			//        self.ensure_one()
			//        context = self._context
			//        created_move_ids = self.env['account.asset.asset'].compute_generated_entries(
			//            self.date, asset_type=context.get('asset_type'))
			//        return {
			//            'name': _('Created Asset Moves') if context.get('asset_type') == 'purchase' else _('Created Revenue Moves'),
			//            'view_type': 'form',
			//            'view_mode': 'tree,form',
			//            'res_model': 'account.move',
			//            'view_id': False,
			//            'domain': "[('id','in',[" + ','.join(map(str, created_move_ids)) + "])]",
			//            'type': 'ir.actions.act_window',
			//        }
		})
}
