package account_asset

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

//import odoo.addons.decimal_precision as dp
func init() {
	h.AccountInvoice().DeclareModel()

	h.AccountInvoice().Methods().RefundCleanupLines().DeclareMethod(
		`RefundCleanupLines`,
		func(rs m.AccountInvoiceSet, lines interface{}) {
			//        result = super(AccountInvoice, self)._refund_cleanup_lines(lines)
			//        for i, line in enumerate(lines):
			//            for name, field in line._fields.items():
			//                if name == 'asset_category_id':
			//                    result[i][2][name] = False
			//                    break
			//        return result
		})
	h.AccountInvoice().Methods().ActionCancel().DeclareMethod(
		`ActionCancel`,
		func(rs m.AccountInvoiceSet) {
			//        res = super(AccountInvoice, self).action_cancel()
			//        self.env['account.asset.asset'].sudo().search(
			//            [('invoice_id', 'in', self.ids)]).write({'active': False})
			//        return res
		})
	h.AccountInvoice().Methods().ActionMoveCreate().DeclareMethod(
		`ActionMoveCreate`,
		func(rs m.AccountInvoiceSet) {
			//        result = super(AccountInvoice, self).action_move_create()
			//        for inv in self:
			//            context = dict(self.env.context)
			//            # Within the context of an invoice,
			//            # this default value is for the type of the invoice, not the type of the asset.
			//            # This has to be cleaned from the context before creating the asset,
			//            # otherwise it tries to create the asset with the type of the invoice.
			//            context.pop('default_type', None)
			//            inv.invoice_line_ids.with_context(context).asset_create()
			//        return result
		})
	h.AccountInvoiceLine().DeclareModel()

	h.AccountInvoiceLine().AddFields(map[string]models.FieldDefinition{
		"AssetCategoryId": models.Many2OneField{
			RelationModel: h.AccountAssetCategory(),
			String:        "Asset Category",
		},
		"AssetStartDate": models.DateField{
			String:   "Asset Start Date",
			Compute:  h.AccountInvoiceLine().Methods().GetAssetDate(),
			ReadOnly: true,
			Stored:   true,
		},
		"AssetEndDate": models.DateField{
			String:   "Asset End Date",
			Compute:  h.AccountInvoiceLine().Methods().GetAssetDate(),
			ReadOnly: true,
			Stored:   true,
		},
		"AssetMrr": models.FloatField{
			String:   "Monthly Recurring Revenue",
			Compute:  h.AccountInvoiceLine().Methods().GetAssetDate(),
			ReadOnly: true,
			//digits=dp.get_precision('Account')
			Stored: true,
		},
	})
	h.AccountInvoiceLine().Methods().GetAssetDate().DeclareMethod(
		`GetAssetDate`,
		func(rs h.AccountInvoiceLineSet) h.AccountInvoiceLineData {
			//        self.asset_mrr = 0
			//        self.asset_start_date = False
			//        self.asset_end_date = False
			//        cat = self.asset_category_id
			//        if cat:
			//            months = cat.method_number * cat.method_period
			//            if self.invoice_id.type in ['out_invoice', 'out_refund']:
			//                self.asset_mrr = self.price_subtotal_signed / months
			//            if self.invoice_id.date_invoice:
			//                start_date = datetime.strptime(
			//                    self.invoice_id.date_invoice, DF).replace(day=1)
			//                end_date = (start_date + relativedelta(months=months, days=-1))
			//                self.asset_start_date = start_date.strftime(DF)
			//                self.asset_end_date = end_date.strftime(DF)
		})
	h.AccountInvoiceLine().Methods().AssetCreate().DeclareMethod(
		`AssetCreate`,
		func(rs m.AccountInvoiceLineSet) {
			//        if self.asset_category_id:
			//            vals = {
			//                'name': self.name,
			//                'code': self.invoice_id.number or False,
			//                'category_id': self.asset_category_id.id,
			//                'value': self.price_subtotal_signed,
			//                'partner_id': self.invoice_id.partner_id.id,
			//                'company_id': self.invoice_id.company_id.id,
			//                'currency_id': self.invoice_id.company_currency_id.id,
			//                'date': self.invoice_id.date_invoice,
			//                'invoice_id': self.invoice_id.id,
			//            }
			//            changed_vals = self.env['account.asset.asset'].onchange_category_id_values(
			//                vals['category_id'])
			//            vals.update(changed_vals['value'])
			//            asset = self.env['account.asset.asset'].create(vals)
			//            if self.asset_category_id.open_asset:
			//                asset.validate()
			//        return True
		})
	h.AccountInvoiceLine().Methods().OnchangeAssetCategoryId().DeclareMethod(
		`OnchangeAssetCategoryId`,
		func(rs m.AccountInvoiceLineSet) {
			//        if self.invoice_id.type == 'out_invoice' and self.asset_category_id:
			//            self.account_id = self.asset_category_id.account_asset_id.id
			//        elif self.invoice_id.type == 'in_invoice' and self.asset_category_id:
			//            self.account_id = self.asset_category_id.account_asset_id.id
		})
	h.AccountInvoiceLine().Methods().OnchangeUomId().DeclareMethod(
		`OnchangeUomId`,
		func(rs m.AccountInvoiceLineSet) {
			//        result = super(AccountInvoiceLine, self)._onchange_uom_id()
			//        self.onchange_asset_category_id()
			//        return result
		})
	h.AccountInvoiceLine().Methods().OnchangeProductId().DeclareMethod(
		`OnchangeProductId`,
		func(rs m.AccountInvoiceLineSet) {
			//        vals = super(AccountInvoiceLine, self)._onchange_product_id()
			//        if self.product_id:
			//            if self.invoice_id.type == 'out_invoice':
			//                self.asset_category_id = self.product_id.product_tmpl_id.deferred_revenue_category_id
			//            elif self.invoice_id.type == 'in_invoice':
			//                self.asset_category_id = self.product_id.product_tmpl_id.asset_category_id
			//        return vals
		})
	h.AccountInvoiceLine().Methods().SetAdditionalFields().DeclareMethod(
		`SetAdditionalFields`,
		func(rs m.AccountInvoiceLineSet, invoice interface{}) {
			//        if not self.asset_category_id:
			//            if invoice.type == 'out_invoice':
			//                self.asset_category_id = self.product_id.product_tmpl_id.deferred_revenue_category_id.id
			//            elif invoice.type == 'in_invoice':
			//                self.asset_category_id = self.product_id.product_tmpl_id.asset_category_id.id
			//            self.onchange_asset_category_id()
			//        super(AccountInvoiceLine, self)._set_additional_fields(invoice)
		})
}
