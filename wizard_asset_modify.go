package account_asset

import (
	"github.com/hexya-erp/hexya-base/web/webdata"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AssetModify().DeclareModel()

	h.AssetModify().AddFields(map[string]models.FieldDefinition{
		"Name": models.TextField{
			String:   "Reason",
			Required: true,
		},
		"MethodNumber": models.IntegerField{
			String:   "Number of Depreciations",
			Required: true,
		},
		"MethodPeriod": models.IntegerField{
			String: "Period Length",
		},
		"MethodEnd": models.DateField{
			String: "Ending date",
		},
		"AssetMethodTime": models.CharField{
			Compute:  h.AssetModify().Methods().GetAssetMethodTime(),
			String:   "Asset Method Time",
			ReadOnly: true,
		},
	})
	h.AssetModify().Methods().GetAssetMethodTime().DeclareMethod(
		`GetAssetMethodTime`,
		func(rs h.AssetModifySet) h.AssetModifyData {
			//        if self.env.context.get('active_id'):
			//            asset = self.env['account.asset.asset'].browse(
			//                self.env.context.get('active_id'))
			//            self.asset_method_time = asset.method_time
		})
	h.AssetModify().Methods().FieldsViewGet().Extend(
		`FieldsViewGet`,
		func(rs m.AssetModifySet, view_id webdata.FieldsViewGetParams, view_type interface{}, toolbar interface{}, submenu interface{}) {
			//        result = super(AssetModify, self).fields_view_get(
			//            view_id, view_type, toolbar=toolbar, submenu=submenu)
			//        asset_id = self.env.context.get('active_id')
			//        active_model = self.env.context.get('active_model')
			//        if active_model == 'account.asset.asset' and asset_id:
			//            asset = self.env['account.asset.asset'].browse(asset_id)
			//            doc = etree.XML(result['arch'])
			//            if asset.method_time == 'number' and doc.xpath("//field[@name='method_end']"):
			//                node = doc.xpath("//field[@name='method_end']")[0]
			//                node.set('invisible', '1')
			//                setup_modifiers(node, result['fields']['method_end'])
			//            elif asset.method_time == 'end' and doc.xpath("//field[@name='method_number']"):
			//                node = doc.xpath("//field[@name='method_number']")[0]
			//                node.set('invisible', '1')
			//                setup_modifiers(node, result['fields']['method_number'])
			//            result['arch'] = etree.tostring(doc)
			//        return result
		})
	h.AssetModify().Methods().DefaultGet().Extend(
		`DefaultGet`,
		func(rs m.AssetModifySet, fields interface{}) {
			//        res = super(AssetModify, self).default_get(fields)
			//        asset_id = self.env.context.get('active_id')
			//        asset = self.env['account.asset.asset'].browse(asset_id)
			//        if 'name' in fields:
			//            res.update({'name': asset.name})
			//        if 'method_number' in fields and asset.method_time == 'number':
			//            res.update({'method_number': asset.method_number})
			//        if 'method_period' in fields:
			//            res.update({'method_period': asset.method_period})
			//        if 'method_end' in fields and asset.method_time == 'end':
			//            res.update({'method_end': asset.method_end})
			//        if self.env.context.get('active_id'):
			//            active_asset = self.env['account.asset.asset'].browse(
			//                self.env.context.get('active_id'))
			//            res['asset_method_time'] = active_asset.method_time
			//        return res
		})
	h.AssetModify().Methods().Modify().DeclareMethod(
		` Modifies the duration of asset for calculating depreciation
        and maintains the history of old values, in the chatter.
        `,
		func(rs m.AssetModifySet) {
			//        asset_id = self.env.context.get('active_id', False)
			//        asset = self.env['account.asset.asset'].browse(asset_id)
			//        old_values = {
			//            'method_number': asset.method_number,
			//            'method_period': asset.method_period,
			//            'method_end': asset.method_end,
			//        }
			//        asset_vals = {
			//            'method_number': self.method_number,
			//            'method_period': self.method_period,
			//            'method_end': self.method_end,
			//        }
			//        asset.write(asset_vals)
			//        asset.compute_depreciation_board()
			//        tracked_fields = self.env['account.asset.asset'].fields_get(
			//            ['method_number', 'method_period', 'method_end'])
			//        changes, tracking_value_ids = asset._message_track(
			//            tracked_fields, old_values)
			//        if changes:
			//            asset.message_post(subject=_('Depreciation board modified'),
			//                               body=self.name, tracking_value_ids=tracking_value_ids)
			//        return {'type': 'ir.actions.act_window_close'}
		})
}
