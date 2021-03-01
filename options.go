package generator

// Options for Document
type Options struct {
	AutoPrint bool `json:"auto_print,omitempty"`

	CurrencySymbol    string `default:"€ " json:"currency_symbol,omitempty"`
	CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
	CurrencyDecimal   string `default:"" json:"currency_decimal,omitempty"`
	CurrencyThousand  string `default:"." json:"currency_thousand,omitempty"`

	TextTypeInvoice      string `default:"INVOICE" json:"text_type_invoice,omitempty"`
	TextTypeQuotation    string `default:"QUOTATION" json:"text_type_quotation,omitempty"`
	TextTypeDeliveryNote string `default:"DELIVERY NOTE" json:"text_type_delivery_note,omitempty"`

	TextRefTitle         string `default:"Ref." json:"text_ref_title,omitempty"`
	TextVersionTitle     string `default:"Version" json:"text_version_title,omitempty"`
	TextDateTitle        string `default:"Date" json:"text_date_title,omitempty"`
	TextPaymentTermTitle string `default:"Payment term" json:"text_payment_term_title,omitempty"`

	TextItemsDescriptionTitle string `default:"Description" json:"text_items_description_title,omitempty"`
	TextItemsUnitCostTitle    string `default:"Unit price" json:"text_items_unit_cost_title,omitempty"`
	TextItemsQuantityTitle    string `default:"Quantity" json:"text_items_quantity_title,omitempty"`
	TextItemsTotalHTTitle     string `default:"Total" json:"text_items_total_ht_title,omitempty"`
	TextItemsTaxTitle         string `default:"Tax" json:"text_items_tax_title,omitempty"`
	TextItemsTotalTTCTitle    string `default:"Total with tax" json:"text_items_total_ttc_title,omitempty"`

	TextTotalTotal    string `default:"TOTAL" json:"text_total_total,omitempty"`
	TextTotalDiscount string `default:"DISCOUNT" json:"text_total_discount,omitempty"`
	TextTotalTax      string `default:"TAX" json:"text_total_tax,omitempty"`
	TextTotalWithTax  string `default:"TOTAL WITH TAX" json:"text_total_with_tax,omitempty"`
}
