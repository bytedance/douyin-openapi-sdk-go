/*
Copyright 2024 ByteDance Ltd. and/or its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client

import (
	"github.com/alibabacloud-go/tea/tea"
)

type DetailTripResponseDataLedgerRecordsItem struct {
	FundAmountType          *int64            `json:"fund_amount_type,omitempty" xml:"fund_amount_type,omitempty"`
	SettleSubType           *int64            `json:"settle_sub_type,omitempty" xml:"settle_sub_type,omitempty"`
	FundAmount              *int64            `json:"fund_amount,omitempty" xml:"fund_amount,omitempty"`
	BookNumber              *string           `json:"book_number,omitempty" xml:"book_number,omitempty"`
	SettleTime              *int64            `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	GroupName               *string           `json:"group_name,omitempty" xml:"group_name,omitempty"`
	BizType                 *int              `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	Amount                  map[string]*int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	OuterVerifyCouponNumber *string           `json:"outer_verify_coupon_number,omitempty" xml:"outer_verify_coupon_number,omitempty"`
	ItemOrderId             *string           `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	Code                    *string           `json:"code,omitempty" xml:"code,omitempty"`
	BillTradeType           *int              `json:"bill_trade_type,omitempty" xml:"bill_trade_type,omitempty"`
	AccountId               *string           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	VerifyPoiAccountId      *string           `json:"verify_poi_account_id,omitempty" xml:"verify_poi_account_id,omitempty"`
	LedgerId                *string           `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	PoiId                   *string           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	OuterVerifyOrderNumber  *string           `json:"outer_verify_order_number,omitempty" xml:"outer_verify_order_number,omitempty"`
	VerifyId                *string           `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	SettleType              *int64            `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	VerifySerialNum         *int32            `json:"verify_serial_num,omitempty" xml:"verify_serial_num,omitempty"`
	SkuId                   *string           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SettleTypeDesc          *string           `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	FlowId                  *string           `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	CertificateId           *string           `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	OrderExtId              *string           `json:"order_ext_id,omitempty" xml:"order_ext_id,omitempty"`
	ShopOrderId             *string           `json:"shop_order_id,omitempty" xml:"shop_order_id,omitempty"`
}

func (s DetailTripResponseDataLedgerRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s DetailTripResponseDataLedgerRecordsItem) GoString() string {
	return s.String()
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetFundAmountType(v int64) *DetailTripResponseDataLedgerRecordsItem {
	s.FundAmountType = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetSettleSubType(v int64) *DetailTripResponseDataLedgerRecordsItem {
	s.SettleSubType = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetFundAmount(v int64) *DetailTripResponseDataLedgerRecordsItem {
	s.FundAmount = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetBookNumber(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.BookNumber = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetSettleTime(v int64) *DetailTripResponseDataLedgerRecordsItem {
	s.SettleTime = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetGroupName(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.GroupName = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetBizType(v int) *DetailTripResponseDataLedgerRecordsItem {
	s.BizType = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetAmount(v map[string]*int64) *DetailTripResponseDataLedgerRecordsItem {
	s.Amount = v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetOuterVerifyCouponNumber(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.OuterVerifyCouponNumber = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetItemOrderId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.ItemOrderId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetCode(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.Code = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetBillTradeType(v int) *DetailTripResponseDataLedgerRecordsItem {
	s.BillTradeType = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetAccountId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.AccountId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetVerifyPoiAccountId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.VerifyPoiAccountId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetLedgerId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.LedgerId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetPoiId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.PoiId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetOuterVerifyOrderNumber(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.OuterVerifyOrderNumber = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetVerifyId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetSettleType(v int64) *DetailTripResponseDataLedgerRecordsItem {
	s.SettleType = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetVerifySerialNum(v int32) *DetailTripResponseDataLedgerRecordsItem {
	s.VerifySerialNum = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetSkuId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.SkuId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetSettleTypeDesc(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.SettleTypeDesc = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetFlowId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.FlowId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetCertificateId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetOrderExtId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.OrderExtId = &v
	return s
}

func (s *DetailTripResponseDataLedgerRecordsItem) SetShopOrderId(v string) *DetailTripResponseDataLedgerRecordsItem {
	s.ShopOrderId = &v
	return s
}

type DetailTripResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s DetailTripResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DetailTripResponseExtra) GoString() string {
	return s.String()
}

func (s *DetailTripResponseExtra) SetNow(v int64) *DetailTripResponseExtra {
	s.Now = &v
	return s
}

func (s *DetailTripResponseExtra) SetSubDescription(v string) *DetailTripResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DetailTripResponseExtra) SetSubErrorCode(v int32) *DetailTripResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DetailTripResponseExtra) SetDescription(v string) *DetailTripResponseExtra {
	s.Description = &v
	return s
}

func (s *DetailTripResponseExtra) SetErrorCode(v int32) *DetailTripResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DetailTripResponseExtra) SetLogid(v string) *DetailTripResponseExtra {
	s.Logid = &v
	return s
}

type DetailedQueryByOrderRequest struct {
	OrderIds    []*string          `json:"order_ids,omitempty" xml:"order_ids,omitempty" require:"true" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DetailedQueryByOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderRequest) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderRequest) SetOrderIds(v []*string) *DetailedQueryByOrderRequest {
	s.OrderIds = v
	return s
}

func (s *DetailedQueryByOrderRequest) SetAccountId(v string) *DetailedQueryByOrderRequest {
	s.AccountId = &v
	return s
}

func (s *DetailedQueryByOrderRequest) SetHeader(v map[string]*string) *DetailedQueryByOrderRequest {
	s.Header = v
	return s
}

func (s *DetailedQueryByOrderRequest) SetAccessToken(v string) *DetailedQueryByOrderRequest {
	s.AccessToken = &v
	return s
}

type DetailedQueryByOrderResponse struct {
	Extra *DetailedQueryByOrderResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *DetailedQueryByOrderResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DetailedQueryByOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponse) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponse) SetExtra(v *DetailedQueryByOrderResponseExtra) *DetailedQueryByOrderResponse {
	s.Extra = v
	return s
}

func (s *DetailedQueryByOrderResponse) SetData(v *DetailedQueryByOrderResponseData) *DetailedQueryByOrderResponse {
	s.Data = v
	return s
}

type DetailedQueryByOrderResponseData struct {
	GwDescription       *string                                                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	LedgerSecondRecords map[string][]*DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem `json:"ledger_second_records,omitempty" xml:"ledger_second_records,omitempty"`
	OrderAmounts        map[string]*DetailedQueryByOrderResponseDataOrderAmountsValue              `json:"order_amounts,omitempty" xml:"order_amounts,omitempty"`
	LedgerRecords       map[string][]*DetailedQueryByOrderResponseDataLedgerRecordsValueItem       `json:"ledger_records,omitempty" xml:"ledger_records,omitempty"`
	GwErrorCode         *int32                                                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DetailedQueryByOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseData) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseData) SetGwDescription(v string) *DetailedQueryByOrderResponseData {
	s.GwDescription = &v
	return s
}

func (s *DetailedQueryByOrderResponseData) SetLedgerSecondRecords(v map[string][]*DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) *DetailedQueryByOrderResponseData {
	s.LedgerSecondRecords = v
	return s
}

func (s *DetailedQueryByOrderResponseData) SetOrderAmounts(v map[string]*DetailedQueryByOrderResponseDataOrderAmountsValue) *DetailedQueryByOrderResponseData {
	s.OrderAmounts = v
	return s
}

func (s *DetailedQueryByOrderResponseData) SetLedgerRecords(v map[string][]*DetailedQueryByOrderResponseDataLedgerRecordsValueItem) *DetailedQueryByOrderResponseData {
	s.LedgerRecords = v
	return s
}

func (s *DetailedQueryByOrderResponseData) SetGwErrorCode(v int32) *DetailedQueryByOrderResponseData {
	s.GwErrorCode = &v
	return s
}

type DetailedQueryByOrderResponseDataLedgerRecordsValueItem struct {
	OrderId   *string                                                       `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Status    *int64                                                        `json:"status,omitempty" xml:"status,omitempty"`
	Goods     *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods  `json:"goods,omitempty" xml:"goods,omitempty"`
	LedgerId  *string                                                       `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	AccountId *string                                                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Amount    *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItem) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItem) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetOrderId(v string) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.OrderId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetStatus(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Status = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetGoods(v *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Goods = v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetLedgerId(v string) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.LedgerId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetAccountId(v string) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.AccountId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItem) SetAmount(v *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) *DetailedQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Amount = v
	return s
}

type DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount struct {
	CouponPay                    *int64 `json:"coupon_pay,omitempty" xml:"coupon_pay,omitempty"`
	PayDiscount                  *int64 `json:"pay_discount,omitempty" xml:"pay_discount,omitempty"`
	TotalAgentMerchant           *int64 `json:"total_agent_merchant,omitempty" xml:"total_agent_merchant,omitempty"`
	LedgerTotal                  *int64 `json:"ledger_total,omitempty" xml:"ledger_total,omitempty"`
	ActualInsured                *int64 `json:"actual_insured,omitempty" xml:"actual_insured,omitempty"`
	MerchantTicket               *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	TalentCommission             *int64 `json:"talent_commission,omitempty" xml:"talent_commission,omitempty"`
	Pay                          *int64 `json:"pay,omitempty" xml:"pay,omitempty"`
	Goods                        *int64 `json:"goods,omitempty" xml:"goods,omitempty"`
	InstitutionCommission        *int64 `json:"institution_commission,omitempty" xml:"institution_commission,omitempty"`
	PlatformTicket               *int64 `json:"platform_ticket,omitempty" xml:"platform_ticket,omitempty"`
	Original                     *int64 `json:"original,omitempty" xml:"original,omitempty"`
	LedgerPlatformTicket         *int64 `json:"ledger_platform_ticket,omitempty" xml:"ledger_platform_ticket,omitempty"`
	TotalMerchantPlatformService *int64 `json:"total_merchant_platform_service,omitempty" xml:"total_merchant_platform_service,omitempty"`
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetCouponPay(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.CouponPay = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetPayDiscount(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.PayDiscount = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetTotalAgentMerchant(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.TotalAgentMerchant = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetLedgerTotal(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.LedgerTotal = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetActualInsured(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.ActualInsured = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetMerchantTicket(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.MerchantTicket = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetTalentCommission(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.TalentCommission = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetPay(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Pay = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetGoods(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Goods = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetInstitutionCommission(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.InstitutionCommission = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetPlatformTicket(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.PlatformTicket = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetOriginal(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Original = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetLedgerPlatformTicket(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.LedgerPlatformTicket = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetTotalMerchantPlatformService(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.TotalMerchantPlatformService = &v
	return s
}

type DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods struct {
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetMarketPrice(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.MarketPrice = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetSkuId(v string) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.SkuId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetSoldStartTime(v int64) *DetailedQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.SoldStartTime = &v
	return s
}

type DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem struct {
	Status    *int64                                                              `json:"status,omitempty" xml:"status,omitempty"`
	LedgerId  *string                                                             `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	OrderId   *string                                                             `json:"order_id,omitempty" xml:"order_id,omitempty"`
	AccountId *string                                                             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Amount    *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetStatus(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.Status = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetLedgerId(v string) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.LedgerId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetOrderId(v string) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.OrderId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetAccountId(v string) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.AccountId = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetAmount(v *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.Amount = v
	return s
}

type DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount struct {
	UserDelivCouponPay       *int64 `json:"user_deliv_coupon_pay,omitempty" xml:"user_deliv_coupon_pay,omitempty"`
	MerchantCancelFreight    *int64 `json:"merchant_cancel_freight,omitempty" xml:"merchant_cancel_freight,omitempty"`
	UserDelivMerchantSubsidy *int64 `json:"user_deliv_merchant_subsidy,omitempty" xml:"user_deliv_merchant_subsidy,omitempty"`
	MerchantFreight          *int64 `json:"merchant_freight,omitempty" xml:"merchant_freight,omitempty"`
	PlatformDutyFreight      *int64 `json:"platform_duty_freight,omitempty" xml:"platform_duty_freight,omitempty"`
	PlatformDutyTip          *int64 `json:"platform_duty_tip,omitempty" xml:"platform_duty_tip,omitempty"`
	UserDelivPlatformSubsidy *int64 `json:"user_deliv_platform_subsidy,omitempty" xml:"user_deliv_platform_subsidy,omitempty"`
	MerchantTip              *int64 `json:"merchant_tip,omitempty" xml:"merchant_tip,omitempty"`
}

func (s DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivCouponPay(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivCouponPay = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantCancelFreight(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantCancelFreight = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivMerchantSubsidy(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivMerchantSubsidy = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantFreight(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantFreight = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetPlatformDutyFreight(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.PlatformDutyFreight = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetPlatformDutyTip(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.PlatformDutyTip = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivPlatformSubsidy(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivPlatformSubsidy = &v
	return s
}

func (s *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantTip(v int64) *DetailedQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantTip = &v
	return s
}

type DetailedQueryByOrderResponseDataOrderAmountsValue struct {
	MerchantIncome *int64 `json:"merchant_income,omitempty" xml:"merchant_income,omitempty"`
}

func (s DetailedQueryByOrderResponseDataOrderAmountsValue) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseDataOrderAmountsValue) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseDataOrderAmountsValue) SetMerchantIncome(v int64) *DetailedQueryByOrderResponseDataOrderAmountsValue {
	s.MerchantIncome = &v
	return s
}

type DetailedQueryByOrderResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s DetailedQueryByOrderResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DetailedQueryByOrderResponseExtra) GoString() string {
	return s.String()
}

func (s *DetailedQueryByOrderResponseExtra) SetDescription(v string) *DetailedQueryByOrderResponseExtra {
	s.Description = &v
	return s
}

func (s *DetailedQueryByOrderResponseExtra) SetErrorCode(v int32) *DetailedQueryByOrderResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DetailedQueryByOrderResponseExtra) SetLogid(v string) *DetailedQueryByOrderResponseExtra {
	s.Logid = &v
	return s
}

func (s *DetailedQueryByOrderResponseExtra) SetNow(v int64) *DetailedQueryByOrderResponseExtra {
	s.Now = &v
	return s
}

func (s *DetailedQueryByOrderResponseExtra) SetSubDescription(v string) *DetailedQueryByOrderResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DetailedQueryByOrderResponseExtra) SetSubErrorCode(v int32) *DetailedQueryByOrderResponseExtra {
	s.SubErrorCode = &v
	return s
}

type DeveloperCancelAuthOrderRequest struct {
	AuthOrderId *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperCancelAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCancelAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCancelAuthOrderRequest) SetAuthOrderId(v string) *DeveloperCancelAuthOrderRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperCancelAuthOrderRequest) SetHeader(v map[string]*string) *DeveloperCancelAuthOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperCancelAuthOrderRequest) SetAccessToken(v string) *DeveloperCancelAuthOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperCancelAuthOrderResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperCancelAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCancelAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCancelAuthOrderResponse) SetErrMsg(v string) *DeveloperCancelAuthOrderResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperCancelAuthOrderResponse) SetLogId(v string) *DeveloperCancelAuthOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperCancelAuthOrderResponse) SetErrNo(v int32) *DeveloperCancelAuthOrderResponse {
	s.ErrNo = &v
	return s
}

type DeveloperCheckDeductPeriodRequest struct {
	MerchantUid    *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	ExtAuthOrderId *string            `json:"ext_auth_order_id,omitempty" xml:"ext_auth_order_id,omitempty"`
	AuthOrderId    *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperCheckDeductPeriodRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCheckDeductPeriodRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCheckDeductPeriodRequest) SetMerchantUid(v string) *DeveloperCheckDeductPeriodRequest {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperCheckDeductPeriodRequest) SetExtAuthOrderId(v string) *DeveloperCheckDeductPeriodRequest {
	s.ExtAuthOrderId = &v
	return s
}

func (s *DeveloperCheckDeductPeriodRequest) SetAuthOrderId(v string) *DeveloperCheckDeductPeriodRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperCheckDeductPeriodRequest) SetHeader(v map[string]*string) *DeveloperCheckDeductPeriodRequest {
	s.Header = v
	return s
}

func (s *DeveloperCheckDeductPeriodRequest) SetAccessToken(v string) *DeveloperCheckDeductPeriodRequest {
	s.AccessToken = &v
	return s
}

type DeveloperCheckDeductPeriodResponse struct {
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperCheckDeductPeriodResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperCheckDeductPeriodResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCheckDeductPeriodResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCheckDeductPeriodResponse) SetLogId(v string) *DeveloperCheckDeductPeriodResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperCheckDeductPeriodResponse) SetData(v *DeveloperCheckDeductPeriodResponseData) *DeveloperCheckDeductPeriodResponse {
	s.Data = v
	return s
}

func (s *DeveloperCheckDeductPeriodResponse) SetErrNo(v int32) *DeveloperCheckDeductPeriodResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperCheckDeductPeriodResponse) SetErrMsg(v string) *DeveloperCheckDeductPeriodResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperCheckDeductPeriodResponseData struct {
	InDeductPeriod        *bool   `json:"in_deduct_period,omitempty" xml:"in_deduct_period,omitempty" require:"true"`
	DeductPeriodStartTime *string `json:"deduct_period_start_time,omitempty" xml:"deduct_period_start_time,omitempty"`
	DeductPeriodEndTime   *string `json:"deduct_period_end_time,omitempty" xml:"deduct_period_end_time,omitempty"`
	PreNotifyAmount       *int64  `json:"pre_notify_amount,omitempty" xml:"pre_notify_amount,omitempty"`
}

func (s DeveloperCheckDeductPeriodResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCheckDeductPeriodResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCheckDeductPeriodResponseData) SetInDeductPeriod(v bool) *DeveloperCheckDeductPeriodResponseData {
	s.InDeductPeriod = &v
	return s
}

func (s *DeveloperCheckDeductPeriodResponseData) SetDeductPeriodStartTime(v string) *DeveloperCheckDeductPeriodResponseData {
	s.DeductPeriodStartTime = &v
	return s
}

func (s *DeveloperCheckDeductPeriodResponseData) SetDeductPeriodEndTime(v string) *DeveloperCheckDeductPeriodResponseData {
	s.DeductPeriodEndTime = &v
	return s
}

func (s *DeveloperCheckDeductPeriodResponseData) SetPreNotifyAmount(v int64) *DeveloperCheckDeductPeriodResponseData {
	s.PreNotifyAmount = &v
	return s
}

type DeveloperClosePayOrderRequest struct {
	PayOrderId  *string            `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperClosePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperClosePayOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperClosePayOrderRequest) SetPayOrderId(v string) *DeveloperClosePayOrderRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperClosePayOrderRequest) SetHeader(v map[string]*string) *DeveloperClosePayOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperClosePayOrderRequest) SetAccessToken(v string) *DeveloperClosePayOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperClosePayOrderResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperClosePayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperClosePayOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperClosePayOrderResponse) SetLogId(v string) *DeveloperClosePayOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperClosePayOrderResponse) SetErrNo(v int32) *DeveloperClosePayOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperClosePayOrderResponse) SetErrMsg(v string) *DeveloperClosePayOrderResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperCreateAuthOrderRequest struct {
	NotifyUrl       *string            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	AdmissibleToken *string            `json:"admissible_token,omitempty" xml:"admissible_token,omitempty" require:"true"`
	OutAuthOrderNo  *string            `json:"out_auth_order_no,omitempty" xml:"out_auth_order_no,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperCreateAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCreateAuthOrderRequest) SetNotifyUrl(v string) *DeveloperCreateAuthOrderRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperCreateAuthOrderRequest) SetAdmissibleToken(v string) *DeveloperCreateAuthOrderRequest {
	s.AdmissibleToken = &v
	return s
}

func (s *DeveloperCreateAuthOrderRequest) SetOutAuthOrderNo(v string) *DeveloperCreateAuthOrderRequest {
	s.OutAuthOrderNo = &v
	return s
}

func (s *DeveloperCreateAuthOrderRequest) SetHeader(v map[string]*string) *DeveloperCreateAuthOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperCreateAuthOrderRequest) SetAccessToken(v string) *DeveloperCreateAuthOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperCreateAuthOrderResponse struct {
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperCreateAuthOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperCreateAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCreateAuthOrderResponse) SetLogId(v string) *DeveloperCreateAuthOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperCreateAuthOrderResponse) SetData(v *DeveloperCreateAuthOrderResponseData) *DeveloperCreateAuthOrderResponse {
	s.Data = v
	return s
}

func (s *DeveloperCreateAuthOrderResponse) SetErrNo(v int32) *DeveloperCreateAuthOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperCreateAuthOrderResponse) SetErrMsg(v string) *DeveloperCreateAuthOrderResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperCreateAuthOrderResponseData struct {
	AuthOrderId *string `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
}

func (s DeveloperCreateAuthOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateAuthOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCreateAuthOrderResponseData) SetAuthOrderId(v string) *DeveloperCreateAuthOrderResponseData {
	s.AuthOrderId = &v
	return s
}

type DeveloperCreatePayOrderRequest struct {
	AuthOrderId   *string                                            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	OutPayOrderNo *string                                            `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty" require:"true"`
	TotalAmount   *int64                                             `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
	FeeDetailList []*DeveloperCreatePayOrderRequestFeeDetailListItem `json:"fee_detail_list,omitempty" xml:"fee_detail_list,omitempty" require:"true" type:"Repeated"`
	NotifyUrl     *string                                            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	Header        map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperCreatePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreatePayOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCreatePayOrderRequest) SetAuthOrderId(v string) *DeveloperCreatePayOrderRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetOutPayOrderNo(v string) *DeveloperCreatePayOrderRequest {
	s.OutPayOrderNo = &v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetTotalAmount(v int64) *DeveloperCreatePayOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetFeeDetailList(v []*DeveloperCreatePayOrderRequestFeeDetailListItem) *DeveloperCreatePayOrderRequest {
	s.FeeDetailList = v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetNotifyUrl(v string) *DeveloperCreatePayOrderRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetHeader(v map[string]*string) *DeveloperCreatePayOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperCreatePayOrderRequest) SetAccessToken(v string) *DeveloperCreatePayOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperCreatePayOrderRequestFeeDetailListItem struct {
	Title       *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Quantity    *int64  `json:"quantity,omitempty" xml:"quantity,omitempty" require:"true"`
}

func (s DeveloperCreatePayOrderRequestFeeDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreatePayOrderRequestFeeDetailListItem) GoString() string {
	return s.String()
}

func (s *DeveloperCreatePayOrderRequestFeeDetailListItem) SetTitle(v string) *DeveloperCreatePayOrderRequestFeeDetailListItem {
	s.Title = &v
	return s
}

func (s *DeveloperCreatePayOrderRequestFeeDetailListItem) SetAmount(v int64) *DeveloperCreatePayOrderRequestFeeDetailListItem {
	s.Amount = &v
	return s
}

func (s *DeveloperCreatePayOrderRequestFeeDetailListItem) SetDescription(v string) *DeveloperCreatePayOrderRequestFeeDetailListItem {
	s.Description = &v
	return s
}

func (s *DeveloperCreatePayOrderRequestFeeDetailListItem) SetQuantity(v int64) *DeveloperCreatePayOrderRequestFeeDetailListItem {
	s.Quantity = &v
	return s
}

type DeveloperCreatePayOrderResponse struct {
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperCreatePayOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperCreatePayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreatePayOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCreatePayOrderResponse) SetErrMsg(v string) *DeveloperCreatePayOrderResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperCreatePayOrderResponse) SetLogId(v string) *DeveloperCreatePayOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperCreatePayOrderResponse) SetData(v *DeveloperCreatePayOrderResponseData) *DeveloperCreatePayOrderResponse {
	s.Data = v
	return s
}

func (s *DeveloperCreatePayOrderResponse) SetErrNo(v int32) *DeveloperCreatePayOrderResponse {
	s.ErrNo = &v
	return s
}

type DeveloperCreatePayOrderResponseData struct {
	PayOrderId *string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
}

func (s DeveloperCreatePayOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreatePayOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCreatePayOrderResponseData) SetPayOrderId(v string) *DeveloperCreatePayOrderResponseData {
	s.PayOrderId = &v
	return s
}

type DeveloperCreateRefundRequest struct {
	NotifyUrl         *string                                          `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	PayOrderId        *string                                          `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	OutPayRefundNo    *string                                          `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty" require:"true"`
	RefundTotalAmount *int64                                           `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	FeeDetailList     []*DeveloperCreateRefundRequestFeeDetailListItem `json:"fee_detail_list,omitempty" xml:"fee_detail_list,omitempty" require:"true" type:"Repeated"`
	Header            map[string]*string                               `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                          `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefundReason      *string                                          `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
}

func (s DeveloperCreateRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateRefundRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCreateRefundRequest) SetNotifyUrl(v string) *DeveloperCreateRefundRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperCreateRefundRequest) SetPayOrderId(v string) *DeveloperCreateRefundRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperCreateRefundRequest) SetOutPayRefundNo(v string) *DeveloperCreateRefundRequest {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperCreateRefundRequest) SetRefundTotalAmount(v int64) *DeveloperCreateRefundRequest {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperCreateRefundRequest) SetFeeDetailList(v []*DeveloperCreateRefundRequestFeeDetailListItem) *DeveloperCreateRefundRequest {
	s.FeeDetailList = v
	return s
}

func (s *DeveloperCreateRefundRequest) SetHeader(v map[string]*string) *DeveloperCreateRefundRequest {
	s.Header = v
	return s
}

func (s *DeveloperCreateRefundRequest) SetAccessToken(v string) *DeveloperCreateRefundRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperCreateRefundRequest) SetRefundReason(v string) *DeveloperCreateRefundRequest {
	s.RefundReason = &v
	return s
}

type DeveloperCreateRefundRequestFeeDetailListItem struct {
	Quantity    *int64  `json:"quantity,omitempty" xml:"quantity,omitempty" require:"true"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s DeveloperCreateRefundRequestFeeDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateRefundRequestFeeDetailListItem) GoString() string {
	return s.String()
}

func (s *DeveloperCreateRefundRequestFeeDetailListItem) SetQuantity(v int64) *DeveloperCreateRefundRequestFeeDetailListItem {
	s.Quantity = &v
	return s
}

func (s *DeveloperCreateRefundRequestFeeDetailListItem) SetTitle(v string) *DeveloperCreateRefundRequestFeeDetailListItem {
	s.Title = &v
	return s
}

func (s *DeveloperCreateRefundRequestFeeDetailListItem) SetAmount(v int64) *DeveloperCreateRefundRequestFeeDetailListItem {
	s.Amount = &v
	return s
}

func (s *DeveloperCreateRefundRequestFeeDetailListItem) SetDescription(v string) *DeveloperCreateRefundRequestFeeDetailListItem {
	s.Description = &v
	return s
}

type DeveloperCreateRefundResponse struct {
	Data   *DeveloperCreateRefundResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperCreateRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateRefundResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCreateRefundResponse) SetData(v *DeveloperCreateRefundResponseData) *DeveloperCreateRefundResponse {
	s.Data = v
	return s
}

func (s *DeveloperCreateRefundResponse) SetErrNo(v int32) *DeveloperCreateRefundResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperCreateRefundResponse) SetErrMsg(v string) *DeveloperCreateRefundResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperCreateRefundResponse) SetLogId(v string) *DeveloperCreateRefundResponse {
	s.LogId = &v
	return s
}

type DeveloperCreateRefundResponseData struct {
	PayRefundId *string `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty" require:"true"`
}

func (s DeveloperCreateRefundResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateRefundResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCreateRefundResponseData) SetPayRefundId(v string) *DeveloperCreateRefundResponseData {
	s.PayRefundId = &v
	return s
}

type DeveloperCreateSignPayRequest struct {
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpireSeconds     *int64             `json:"expire_seconds,omitempty" xml:"expire_seconds,omitempty"`
	TotalAmount       *int64             `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	AuthOrderId       *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	NotifyUrl         *string            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	PlatformPreNotify *int32             `json:"platform_pre_notify,omitempty" xml:"platform_pre_notify,omitempty"`
	OutPayOrderNo     *string            `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty" require:"true"`
	MerchantUid       *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty" require:"true"`
}

func (s DeveloperCreateSignPayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignPayRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignPayRequest) SetAccessToken(v string) *DeveloperCreateSignPayRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetExpireSeconds(v int64) *DeveloperCreateSignPayRequest {
	s.ExpireSeconds = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetTotalAmount(v int64) *DeveloperCreateSignPayRequest {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetAuthOrderId(v string) *DeveloperCreateSignPayRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetNotifyUrl(v string) *DeveloperCreateSignPayRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetHeader(v map[string]*string) *DeveloperCreateSignPayRequest {
	s.Header = v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetPlatformPreNotify(v int32) *DeveloperCreateSignPayRequest {
	s.PlatformPreNotify = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetOutPayOrderNo(v string) *DeveloperCreateSignPayRequest {
	s.OutPayOrderNo = &v
	return s
}

func (s *DeveloperCreateSignPayRequest) SetMerchantUid(v string) *DeveloperCreateSignPayRequest {
	s.MerchantUid = &v
	return s
}

type DeveloperCreateSignPayResponse struct {
	Data   *DeveloperCreateSignPayResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperCreateSignPayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignPayResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignPayResponse) SetData(v *DeveloperCreateSignPayResponseData) *DeveloperCreateSignPayResponse {
	s.Data = v
	return s
}

func (s *DeveloperCreateSignPayResponse) SetErrNo(v int32) *DeveloperCreateSignPayResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperCreateSignPayResponse) SetErrMsg(v string) *DeveloperCreateSignPayResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperCreateSignPayResponse) SetLogId(v string) *DeveloperCreateSignPayResponse {
	s.LogId = &v
	return s
}

type DeveloperCreateSignPayResponseData struct {
	PayOrderId *string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
}

func (s DeveloperCreateSignPayResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignPayResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignPayResponseData) SetPayOrderId(v string) *DeveloperCreateSignPayResponseData {
	s.PayOrderId = &v
	return s
}

type DeveloperCreateSignRefundRequest struct {
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PayOrderId        *string            `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	OutPayRefundNo    *string            `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty" require:"true"`
	RefundTotalAmount *int64             `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	RefundReason      *string            `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	NotifyUrl         *string            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
}

func (s DeveloperCreateSignRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignRefundRequest) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignRefundRequest) SetHeader(v map[string]*string) *DeveloperCreateSignRefundRequest {
	s.Header = v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetAccessToken(v string) *DeveloperCreateSignRefundRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetPayOrderId(v string) *DeveloperCreateSignRefundRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetOutPayRefundNo(v string) *DeveloperCreateSignRefundRequest {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetRefundTotalAmount(v int64) *DeveloperCreateSignRefundRequest {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetRefundReason(v string) *DeveloperCreateSignRefundRequest {
	s.RefundReason = &v
	return s
}

func (s *DeveloperCreateSignRefundRequest) SetNotifyUrl(v string) *DeveloperCreateSignRefundRequest {
	s.NotifyUrl = &v
	return s
}

type DeveloperCreateSignRefundResponse struct {
	Data   *DeveloperCreateSignRefundResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperCreateSignRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignRefundResponse) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignRefundResponse) SetData(v *DeveloperCreateSignRefundResponseData) *DeveloperCreateSignRefundResponse {
	s.Data = v
	return s
}

func (s *DeveloperCreateSignRefundResponse) SetErrNo(v int32) *DeveloperCreateSignRefundResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperCreateSignRefundResponse) SetErrMsg(v string) *DeveloperCreateSignRefundResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperCreateSignRefundResponse) SetLogId(v string) *DeveloperCreateSignRefundResponse {
	s.LogId = &v
	return s
}

type DeveloperCreateSignRefundResponseData struct {
	PayRefundId *string `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty" require:"true"`
}

func (s DeveloperCreateSignRefundResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperCreateSignRefundResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperCreateSignRefundResponseData) SetPayRefundId(v string) *DeveloperCreateSignRefundResponseData {
	s.PayRefundId = &v
	return s
}

type DeveloperFinishAuthOrderRequest struct {
	AuthOrderId *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperFinishAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperFinishAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperFinishAuthOrderRequest) SetAuthOrderId(v string) *DeveloperFinishAuthOrderRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperFinishAuthOrderRequest) SetHeader(v map[string]*string) *DeveloperFinishAuthOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperFinishAuthOrderRequest) SetAccessToken(v string) *DeveloperFinishAuthOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperFinishAuthOrderResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperFinishAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperFinishAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperFinishAuthOrderResponse) SetErrNo(v int32) *DeveloperFinishAuthOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperFinishAuthOrderResponse) SetErrMsg(v string) *DeveloperFinishAuthOrderResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperFinishAuthOrderResponse) SetLogId(v string) *DeveloperFinishAuthOrderResponse {
	s.LogId = &v
	return s
}

type DeveloperFulfillPushStatusRequest struct {
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ToStatus        *string            `json:"to_status,omitempty" xml:"to_status,omitempty" require:"true"`
	OrderId         *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ItemOrderIdList []*string          `json:"item_order_id_list,omitempty" xml:"item_order_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperFulfillPushStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperFulfillPushStatusRequest) GoString() string {
	return s.String()
}

func (s *DeveloperFulfillPushStatusRequest) SetHeader(v map[string]*string) *DeveloperFulfillPushStatusRequest {
	s.Header = v
	return s
}

func (s *DeveloperFulfillPushStatusRequest) SetAccessToken(v string) *DeveloperFulfillPushStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperFulfillPushStatusRequest) SetToStatus(v string) *DeveloperFulfillPushStatusRequest {
	s.ToStatus = &v
	return s
}

func (s *DeveloperFulfillPushStatusRequest) SetOrderId(v string) *DeveloperFulfillPushStatusRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperFulfillPushStatusRequest) SetItemOrderIdList(v []*string) *DeveloperFulfillPushStatusRequest {
	s.ItemOrderIdList = v
	return s
}

type DeveloperFulfillPushStatusResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperFulfillPushStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperFulfillPushStatusResponse) GoString() string {
	return s.String()
}

func (s *DeveloperFulfillPushStatusResponse) SetErrNo(v int64) *DeveloperFulfillPushStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperFulfillPushStatusResponse) SetErrMsg(v string) *DeveloperFulfillPushStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperFulfillPushStatusResponse) SetLogId(v string) *DeveloperFulfillPushStatusResponse {
	s.LogId = &v
	return s
}

type DeveloperOrderQueryRequest struct {
	OutOrderNo  *string            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *DeveloperOrderQueryRequest) SetOutOrderNo(v string) *DeveloperOrderQueryRequest {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperOrderQueryRequest) SetOrderId(v string) *DeveloperOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperOrderQueryRequest) SetHeader(v map[string]*string) *DeveloperOrderQueryRequest {
	s.Header = v
	return s
}

func (s *DeveloperOrderQueryRequest) SetAccessToken(v string) *DeveloperOrderQueryRequest {
	s.AccessToken = &v
	return s
}

type DeveloperOrderQueryResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperOrderQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *DeveloperOrderQueryResponse) SetLogId(v string) *DeveloperOrderQueryResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperOrderQueryResponse) SetData(v *DeveloperOrderQueryResponseData) *DeveloperOrderQueryResponse {
	s.Data = v
	return s
}

func (s *DeveloperOrderQueryResponse) SetErrNo(v int32) *DeveloperOrderQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperOrderQueryResponse) SetErrMsg(v string) *DeveloperOrderQueryResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperOrderQueryResponseData struct {
	AppId               *string                                             `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	PayTime             *int64                                              `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	PayStatus           *string                                             `json:"pay_status,omitempty" xml:"pay_status,omitempty" require:"true"`
	SubmisssionTime     *int64                                              `json:"submisssion_time,omitempty" xml:"submisssion_time,omitempty"`
	Currency            *string                                             `json:"currency,omitempty" xml:"currency,omitempty"`
	ItemOrderList       []*DeveloperOrderQueryResponseDataItemOrderListItem `json:"item_order_list,omitempty" xml:"item_order_list,omitempty" type:"Repeated"`
	ServiceCancelReason *string                                             `json:"service_cancel_reason,omitempty" xml:"service_cancel_reason,omitempty"`
	TotalAmount         *int64                                              `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
	OrderServiceStatus  *string                                             `json:"order_service_status,omitempty" xml:"order_service_status,omitempty"`
	TotalCurrencyAmount *int64                                              `json:"total_currency_amount,omitempty" xml:"total_currency_amount,omitempty"`
	OrderId             *string                                             `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	TradeTime           *int64                                              `json:"trade_time,omitempty" xml:"trade_time,omitempty" require:"true"`
	DiscountAmount      *int64                                              `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	OutOrderNo          *string                                             `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	MerchantUid         *string                                             `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	PayChannel          *int                                                `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	ChannelPayId        *string                                             `json:"channel_pay_id,omitempty" xml:"channel_pay_id,omitempty"`
}

func (s DeveloperOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperOrderQueryResponseData) SetAppId(v string) *DeveloperOrderQueryResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetPayTime(v int64) *DeveloperOrderQueryResponseData {
	s.PayTime = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetPayStatus(v string) *DeveloperOrderQueryResponseData {
	s.PayStatus = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetSubmisssionTime(v int64) *DeveloperOrderQueryResponseData {
	s.SubmisssionTime = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetCurrency(v string) *DeveloperOrderQueryResponseData {
	s.Currency = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetItemOrderList(v []*DeveloperOrderQueryResponseDataItemOrderListItem) *DeveloperOrderQueryResponseData {
	s.ItemOrderList = v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetServiceCancelReason(v string) *DeveloperOrderQueryResponseData {
	s.ServiceCancelReason = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetTotalAmount(v int64) *DeveloperOrderQueryResponseData {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetOrderServiceStatus(v string) *DeveloperOrderQueryResponseData {
	s.OrderServiceStatus = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetTotalCurrencyAmount(v int64) *DeveloperOrderQueryResponseData {
	s.TotalCurrencyAmount = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetOrderId(v string) *DeveloperOrderQueryResponseData {
	s.OrderId = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetTradeTime(v int64) *DeveloperOrderQueryResponseData {
	s.TradeTime = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetDiscountAmount(v int64) *DeveloperOrderQueryResponseData {
	s.DiscountAmount = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetOutOrderNo(v string) *DeveloperOrderQueryResponseData {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetMerchantUid(v string) *DeveloperOrderQueryResponseData {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetPayChannel(v int) *DeveloperOrderQueryResponseData {
	s.PayChannel = &v
	return s
}

func (s *DeveloperOrderQueryResponseData) SetChannelPayId(v string) *DeveloperOrderQueryResponseData {
	s.ChannelPayId = &v
	return s
}

type DeveloperOrderQueryResponseDataItemOrderListItem struct {
	ItemOrderId             *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
	SkuId                   *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	ItemOrderAmount         *int64  `json:"item_order_amount,omitempty" xml:"item_order_amount,omitempty" require:"true"`
	ItemOrderCurrencyAmount *int64  `json:"item_order_currency_amount,omitempty" xml:"item_order_currency_amount,omitempty"`
}

func (s DeveloperOrderQueryResponseDataItemOrderListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperOrderQueryResponseDataItemOrderListItem) GoString() string {
	return s.String()
}

func (s *DeveloperOrderQueryResponseDataItemOrderListItem) SetItemOrderId(v string) *DeveloperOrderQueryResponseDataItemOrderListItem {
	s.ItemOrderId = &v
	return s
}

func (s *DeveloperOrderQueryResponseDataItemOrderListItem) SetSkuId(v string) *DeveloperOrderQueryResponseDataItemOrderListItem {
	s.SkuId = &v
	return s
}

func (s *DeveloperOrderQueryResponseDataItemOrderListItem) SetItemOrderAmount(v int64) *DeveloperOrderQueryResponseDataItemOrderListItem {
	s.ItemOrderAmount = &v
	return s
}

func (s *DeveloperOrderQueryResponseDataItemOrderListItem) SetItemOrderCurrencyAmount(v int64) *DeveloperOrderQueryResponseDataItemOrderListItem {
	s.ItemOrderCurrencyAmount = &v
	return s
}

type DeveloperPreNotifyRequest struct {
	AuthOrderId    *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	MerchantUid    *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	Amount         *int64             `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	ExtAuthOrderId *string            `json:"ext_auth_order_id,omitempty" xml:"ext_auth_order_id,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperPreNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperPreNotifyRequest) GoString() string {
	return s.String()
}

func (s *DeveloperPreNotifyRequest) SetAuthOrderId(v string) *DeveloperPreNotifyRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperPreNotifyRequest) SetMerchantUid(v string) *DeveloperPreNotifyRequest {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperPreNotifyRequest) SetAmount(v int64) *DeveloperPreNotifyRequest {
	s.Amount = &v
	return s
}

func (s *DeveloperPreNotifyRequest) SetExtAuthOrderId(v string) *DeveloperPreNotifyRequest {
	s.ExtAuthOrderId = &v
	return s
}

func (s *DeveloperPreNotifyRequest) SetHeader(v map[string]*string) *DeveloperPreNotifyRequest {
	s.Header = v
	return s
}

func (s *DeveloperPreNotifyRequest) SetAccessToken(v string) *DeveloperPreNotifyRequest {
	s.AccessToken = &v
	return s
}

type DeveloperPreNotifyResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperPreNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperPreNotifyResponse) GoString() string {
	return s.String()
}

func (s *DeveloperPreNotifyResponse) SetErrMsg(v string) *DeveloperPreNotifyResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperPreNotifyResponse) SetLogId(v string) *DeveloperPreNotifyResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperPreNotifyResponse) SetErrNo(v int32) *DeveloperPreNotifyResponse {
	s.ErrNo = &v
	return s
}

type DeveloperQueryAdmissibleAuthRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TotalAmount *int64             `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
	MerchantUid *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	ServiceId   *string            `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Scene       *int64             `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
}

func (s DeveloperQueryAdmissibleAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAdmissibleAuthRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetOpenId(v string) *DeveloperQueryAdmissibleAuthRequest {
	s.OpenId = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetTotalAmount(v int64) *DeveloperQueryAdmissibleAuthRequest {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetMerchantUid(v string) *DeveloperQueryAdmissibleAuthRequest {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetServiceId(v string) *DeveloperQueryAdmissibleAuthRequest {
	s.ServiceId = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetHeader(v map[string]*string) *DeveloperQueryAdmissibleAuthRequest {
	s.Header = v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetAccessToken(v string) *DeveloperQueryAdmissibleAuthRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthRequest) SetScene(v int64) *DeveloperQueryAdmissibleAuthRequest {
	s.Scene = &v
	return s
}

type DeveloperQueryAdmissibleAuthResponse struct {
	ErrMsg *string                                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQueryAdmissibleAuthResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperQueryAdmissibleAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAdmissibleAuthResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAdmissibleAuthResponse) SetErrMsg(v string) *DeveloperQueryAdmissibleAuthResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthResponse) SetLogId(v string) *DeveloperQueryAdmissibleAuthResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthResponse) SetData(v *DeveloperQueryAdmissibleAuthResponseData) *DeveloperQueryAdmissibleAuthResponse {
	s.Data = v
	return s
}

func (s *DeveloperQueryAdmissibleAuthResponse) SetErrNo(v int32) *DeveloperQueryAdmissibleAuthResponse {
	s.ErrNo = &v
	return s
}

type DeveloperQueryAdmissibleAuthResponseData struct {
	Admissibility   *int    `json:"admissibility,omitempty" xml:"admissibility,omitempty" require:"true"`
	AdmissibleToken *string `json:"admissible_token,omitempty" xml:"admissible_token,omitempty"`
	ExpireSeconds   *int64  `json:"expire_seconds,omitempty" xml:"expire_seconds,omitempty"`
}

func (s DeveloperQueryAdmissibleAuthResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAdmissibleAuthResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAdmissibleAuthResponseData) SetAdmissibility(v int) *DeveloperQueryAdmissibleAuthResponseData {
	s.Admissibility = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthResponseData) SetAdmissibleToken(v string) *DeveloperQueryAdmissibleAuthResponseData {
	s.AdmissibleToken = &v
	return s
}

func (s *DeveloperQueryAdmissibleAuthResponseData) SetExpireSeconds(v int64) *DeveloperQueryAdmissibleAuthResponseData {
	s.ExpireSeconds = &v
	return s
}

type DeveloperQueryAuthOrderRequest struct {
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AuthOrderId    *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty"`
	OutAuthOrderNo *string            `json:"out_auth_order_no,omitempty" xml:"out_auth_order_no,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DeveloperQueryAuthOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAuthOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAuthOrderRequest) SetAccessToken(v string) *DeveloperQueryAuthOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperQueryAuthOrderRequest) SetAuthOrderId(v string) *DeveloperQueryAuthOrderRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperQueryAuthOrderRequest) SetOutAuthOrderNo(v string) *DeveloperQueryAuthOrderRequest {
	s.OutAuthOrderNo = &v
	return s
}

func (s *DeveloperQueryAuthOrderRequest) SetHeader(v map[string]*string) *DeveloperQueryAuthOrderRequest {
	s.Header = v
	return s
}

type DeveloperQueryAuthOrderResponse struct {
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQueryAuthOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperQueryAuthOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAuthOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAuthOrderResponse) SetErrMsg(v string) *DeveloperQueryAuthOrderResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponse) SetLogId(v string) *DeveloperQueryAuthOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponse) SetData(v *DeveloperQueryAuthOrderResponseData) *DeveloperQueryAuthOrderResponse {
	s.Data = v
	return s
}

func (s *DeveloperQueryAuthOrderResponse) SetErrNo(v int32) *DeveloperQueryAuthOrderResponse {
	s.ErrNo = &v
	return s
}

type DeveloperQueryAuthOrderResponseData struct {
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	ServiceId      *string `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	OutAuthOrderId *string `json:"out_auth_order_id,omitempty" xml:"out_auth_order_id,omitempty" require:"true"`
	OpenId         *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	SignTime       *int64  `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	NotifyUrl      *string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Scene          *int64  `json:"scene,omitempty" xml:"scene,omitempty" require:"true"`
	AuthOrderId    *string `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
}

func (s DeveloperQueryAuthOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryAuthOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQueryAuthOrderResponseData) SetAppId(v string) *DeveloperQueryAuthOrderResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetServiceId(v string) *DeveloperQueryAuthOrderResponseData {
	s.ServiceId = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetOutAuthOrderId(v string) *DeveloperQueryAuthOrderResponseData {
	s.OutAuthOrderId = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetOpenId(v string) *DeveloperQueryAuthOrderResponseData {
	s.OpenId = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetSignTime(v int64) *DeveloperQueryAuthOrderResponseData {
	s.SignTime = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetNotifyUrl(v string) *DeveloperQueryAuthOrderResponseData {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetStatus(v string) *DeveloperQueryAuthOrderResponseData {
	s.Status = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetScene(v int64) *DeveloperQueryAuthOrderResponseData {
	s.Scene = &v
	return s
}

func (s *DeveloperQueryAuthOrderResponseData) SetAuthOrderId(v string) *DeveloperQueryAuthOrderResponseData {
	s.AuthOrderId = &v
	return s
}

type DeveloperQueryCpsRequest struct {
	OutOrderNo  *string            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperQueryCpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryCpsRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQueryCpsRequest) SetOutOrderNo(v string) *DeveloperQueryCpsRequest {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperQueryCpsRequest) SetOrderId(v string) *DeveloperQueryCpsRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperQueryCpsRequest) SetHeader(v map[string]*string) *DeveloperQueryCpsRequest {
	s.Header = v
	return s
}

func (s *DeveloperQueryCpsRequest) SetAccessToken(v string) *DeveloperQueryCpsRequest {
	s.AccessToken = &v
	return s
}

type DeveloperQueryCpsResponse struct {
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQueryCpsResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperQueryCpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryCpsResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQueryCpsResponse) SetErrMsg(v string) *DeveloperQueryCpsResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperQueryCpsResponse) SetLogId(v string) *DeveloperQueryCpsResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQueryCpsResponse) SetData(v *DeveloperQueryCpsResponseData) *DeveloperQueryCpsResponse {
	s.Data = v
	return s
}

func (s *DeveloperQueryCpsResponse) SetErrNo(v int32) *DeveloperQueryCpsResponse {
	s.ErrNo = &v
	return s
}

type DeveloperQueryCpsResponseData struct {
	OrderId               *string                                         `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	OutOrderNo            *string                                         `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	TotalCommissionAmount *int64                                          `json:"total_commission_amount,omitempty" xml:"total_commission_amount,omitempty" require:"true"`
	CpsItemList           []*DeveloperQueryCpsResponseDataCpsItemListItem `json:"cps_item_list,omitempty" xml:"cps_item_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperQueryCpsResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryCpsResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQueryCpsResponseData) SetOrderId(v string) *DeveloperQueryCpsResponseData {
	s.OrderId = &v
	return s
}

func (s *DeveloperQueryCpsResponseData) SetOutOrderNo(v string) *DeveloperQueryCpsResponseData {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperQueryCpsResponseData) SetTotalCommissionAmount(v int64) *DeveloperQueryCpsResponseData {
	s.TotalCommissionAmount = &v
	return s
}

func (s *DeveloperQueryCpsResponseData) SetCpsItemList(v []*DeveloperQueryCpsResponseDataCpsItemListItem) *DeveloperQueryCpsResponseData {
	s.CpsItemList = v
	return s
}

type DeveloperQueryCpsResponseDataCpsItemListItem struct {
	CommissionAmount       *int64  `json:"commission_amount,omitempty" xml:"commission_amount,omitempty" require:"true"`
	SourceType             *int32  `json:"source_type,omitempty" xml:"source_type,omitempty" require:"true"`
	ItemOrderId            *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
	CommissionUserNickname *string `json:"commission_user_nickname,omitempty" xml:"commission_user_nickname,omitempty" require:"true"`
	CommissionUserDouyinid *string `json:"commission_user_douyinid,omitempty" xml:"commission_user_douyinid,omitempty" require:"true"`
	ItemId                 *int64  `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Status                 *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CommissionRate         *int32  `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	SellAmount             *int64  `json:"sell_amount,omitempty" xml:"sell_amount,omitempty" require:"true"`
	TaskId                 *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s DeveloperQueryCpsResponseDataCpsItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryCpsResponseDataCpsItemListItem) GoString() string {
	return s.String()
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetCommissionAmount(v int64) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.CommissionAmount = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetSourceType(v int32) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.SourceType = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetItemOrderId(v string) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.ItemOrderId = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetCommissionUserNickname(v string) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.CommissionUserNickname = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetCommissionUserDouyinid(v string) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.CommissionUserDouyinid = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetItemId(v int64) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.ItemId = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetStatus(v int32) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.Status = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetCommissionRate(v int32) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.CommissionRate = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetSellAmount(v int64) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.SellAmount = &v
	return s
}

func (s *DeveloperQueryCpsResponseDataCpsItemListItem) SetTaskId(v string) *DeveloperQueryCpsResponseDataCpsItemListItem {
	s.TaskId = &v
	return s
}

type DeveloperQueryPayOrderRequest struct {
	OutPayOrderNo *string            `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty"`
	PayOrderId    *string            `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperQueryPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryPayOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQueryPayOrderRequest) SetOutPayOrderNo(v string) *DeveloperQueryPayOrderRequest {
	s.OutPayOrderNo = &v
	return s
}

func (s *DeveloperQueryPayOrderRequest) SetPayOrderId(v string) *DeveloperQueryPayOrderRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQueryPayOrderRequest) SetHeader(v map[string]*string) *DeveloperQueryPayOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperQueryPayOrderRequest) SetAccessToken(v string) *DeveloperQueryPayOrderRequest {
	s.AccessToken = &v
	return s
}

type DeveloperQueryPayOrderResponse struct {
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQueryPayOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperQueryPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryPayOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQueryPayOrderResponse) SetLogId(v string) *DeveloperQueryPayOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQueryPayOrderResponse) SetData(v *DeveloperQueryPayOrderResponseData) *DeveloperQueryPayOrderResponse {
	s.Data = v
	return s
}

func (s *DeveloperQueryPayOrderResponse) SetErrNo(v int32) *DeveloperQueryPayOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperQueryPayOrderResponse) SetErrMsg(v string) *DeveloperQueryPayOrderResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperQueryPayOrderResponseData struct {
	ChannelPayId  *string                                                `json:"channel_pay_id,omitempty" xml:"channel_pay_id,omitempty"`
	AuthOrderId   *string                                                `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	TotalAmount   *int64                                                 `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
	MerchantUid   *string                                                `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	PayChannel    *int32                                                 `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	OutPayOrderNo *string                                                `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty" require:"true"`
	Status        *string                                                `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AppId         *string                                                `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	PayTime       *int64                                                 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	FeeDetailList []*DeveloperQueryPayOrderResponseDataFeeDetailListItem `json:"fee_detail_list,omitempty" xml:"fee_detail_list,omitempty" require:"true" type:"Repeated"`
	PayOrderId    *string                                                `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	NotifyUrl     *string                                                `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
}

func (s DeveloperQueryPayOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryPayOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQueryPayOrderResponseData) SetChannelPayId(v string) *DeveloperQueryPayOrderResponseData {
	s.ChannelPayId = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetAuthOrderId(v string) *DeveloperQueryPayOrderResponseData {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetTotalAmount(v int64) *DeveloperQueryPayOrderResponseData {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetMerchantUid(v string) *DeveloperQueryPayOrderResponseData {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetPayChannel(v int32) *DeveloperQueryPayOrderResponseData {
	s.PayChannel = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetOutPayOrderNo(v string) *DeveloperQueryPayOrderResponseData {
	s.OutPayOrderNo = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetStatus(v string) *DeveloperQueryPayOrderResponseData {
	s.Status = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetAppId(v string) *DeveloperQueryPayOrderResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetPayTime(v int64) *DeveloperQueryPayOrderResponseData {
	s.PayTime = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetFeeDetailList(v []*DeveloperQueryPayOrderResponseDataFeeDetailListItem) *DeveloperQueryPayOrderResponseData {
	s.FeeDetailList = v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetPayOrderId(v string) *DeveloperQueryPayOrderResponseData {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseData) SetNotifyUrl(v string) *DeveloperQueryPayOrderResponseData {
	s.NotifyUrl = &v
	return s
}

type DeveloperQueryPayOrderResponseDataFeeDetailListItem struct {
	Quantity    *int64  `json:"quantity,omitempty" xml:"quantity,omitempty" require:"true"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DeveloperQueryPayOrderResponseDataFeeDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryPayOrderResponseDataFeeDetailListItem) GoString() string {
	return s.String()
}

func (s *DeveloperQueryPayOrderResponseDataFeeDetailListItem) SetQuantity(v int64) *DeveloperQueryPayOrderResponseDataFeeDetailListItem {
	s.Quantity = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseDataFeeDetailListItem) SetTitle(v string) *DeveloperQueryPayOrderResponseDataFeeDetailListItem {
	s.Title = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseDataFeeDetailListItem) SetAmount(v int64) *DeveloperQueryPayOrderResponseDataFeeDetailListItem {
	s.Amount = &v
	return s
}

func (s *DeveloperQueryPayOrderResponseDataFeeDetailListItem) SetDescription(v string) *DeveloperQueryPayOrderResponseDataFeeDetailListItem {
	s.Description = &v
	return s
}

type DeveloperQueryRefundRequest struct {
	OutPayRefundNo *string            `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty"`
	PayRefundId    *string            `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperQueryRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryRefundRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQueryRefundRequest) SetOutPayRefundNo(v string) *DeveloperQueryRefundRequest {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperQueryRefundRequest) SetPayRefundId(v string) *DeveloperQueryRefundRequest {
	s.PayRefundId = &v
	return s
}

func (s *DeveloperQueryRefundRequest) SetHeader(v map[string]*string) *DeveloperQueryRefundRequest {
	s.Header = v
	return s
}

func (s *DeveloperQueryRefundRequest) SetAccessToken(v string) *DeveloperQueryRefundRequest {
	s.AccessToken = &v
	return s
}

type DeveloperQueryRefundResponse struct {
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQueryRefundResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperQueryRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryRefundResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQueryRefundResponse) SetErrMsg(v string) *DeveloperQueryRefundResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperQueryRefundResponse) SetLogId(v string) *DeveloperQueryRefundResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQueryRefundResponse) SetData(v *DeveloperQueryRefundResponseData) *DeveloperQueryRefundResponse {
	s.Data = v
	return s
}

func (s *DeveloperQueryRefundResponse) SetErrNo(v int32) *DeveloperQueryRefundResponse {
	s.ErrNo = &v
	return s
}

type DeveloperQueryRefundResponseData struct {
	Message           *string                                              `json:"message,omitempty" xml:"message,omitempty"`
	RefundTotalAmount *int64                                               `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	PayOrderId        *string                                              `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	PayRefundId       *string                                              `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty" require:"true"`
	RefundReason      *string                                              `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	OutPayRefundNo    *string                                              `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty" require:"true"`
	AppId             *string                                              `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	RefundAt          *int64                                               `json:"refund_at,omitempty" xml:"refund_at,omitempty"`
	FeeDetailList     []*DeveloperQueryRefundResponseDataFeeDetailListItem `json:"fee_detail_list,omitempty" xml:"fee_detail_list,omitempty" require:"true" type:"Repeated"`
	Status            *string                                              `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	NotifyUrl         *string                                              `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
}

func (s DeveloperQueryRefundResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryRefundResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQueryRefundResponseData) SetMessage(v string) *DeveloperQueryRefundResponseData {
	s.Message = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetRefundTotalAmount(v int64) *DeveloperQueryRefundResponseData {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetPayOrderId(v string) *DeveloperQueryRefundResponseData {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetPayRefundId(v string) *DeveloperQueryRefundResponseData {
	s.PayRefundId = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetRefundReason(v string) *DeveloperQueryRefundResponseData {
	s.RefundReason = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetOutPayRefundNo(v string) *DeveloperQueryRefundResponseData {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetAppId(v string) *DeveloperQueryRefundResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetRefundAt(v int64) *DeveloperQueryRefundResponseData {
	s.RefundAt = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetFeeDetailList(v []*DeveloperQueryRefundResponseDataFeeDetailListItem) *DeveloperQueryRefundResponseData {
	s.FeeDetailList = v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetStatus(v string) *DeveloperQueryRefundResponseData {
	s.Status = &v
	return s
}

func (s *DeveloperQueryRefundResponseData) SetNotifyUrl(v string) *DeveloperQueryRefundResponseData {
	s.NotifyUrl = &v
	return s
}

type DeveloperQueryRefundResponseDataFeeDetailListItem struct {
	Quantity    *int64  `json:"quantity,omitempty" xml:"quantity,omitempty" require:"true"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s DeveloperQueryRefundResponseDataFeeDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQueryRefundResponseDataFeeDetailListItem) GoString() string {
	return s.String()
}

func (s *DeveloperQueryRefundResponseDataFeeDetailListItem) SetQuantity(v int64) *DeveloperQueryRefundResponseDataFeeDetailListItem {
	s.Quantity = &v
	return s
}

func (s *DeveloperQueryRefundResponseDataFeeDetailListItem) SetTitle(v string) *DeveloperQueryRefundResponseDataFeeDetailListItem {
	s.Title = &v
	return s
}

func (s *DeveloperQueryRefundResponseDataFeeDetailListItem) SetAmount(v int64) *DeveloperQueryRefundResponseDataFeeDetailListItem {
	s.Amount = &v
	return s
}

func (s *DeveloperQueryRefundResponseDataFeeDetailListItem) SetDescription(v string) *DeveloperQueryRefundResponseDataFeeDetailListItem {
	s.Description = &v
	return s
}

type DeveloperQuerySignOrderRequest struct {
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AuthOrderId    *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty"`
	OutAuthOrderNo *string            `json:"out_auth_order_no,omitempty" xml:"out_auth_order_no,omitempty"`
}

func (s DeveloperQuerySignOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignOrderRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignOrderRequest) SetHeader(v map[string]*string) *DeveloperQuerySignOrderRequest {
	s.Header = v
	return s
}

func (s *DeveloperQuerySignOrderRequest) SetAccessToken(v string) *DeveloperQuerySignOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperQuerySignOrderRequest) SetAuthOrderId(v string) *DeveloperQuerySignOrderRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperQuerySignOrderRequest) SetOutAuthOrderNo(v string) *DeveloperQuerySignOrderRequest {
	s.OutAuthOrderNo = &v
	return s
}

type DeveloperQuerySignOrderResponse struct {
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQuerySignOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperQuerySignOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignOrderResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignOrderResponse) SetLogId(v string) *DeveloperQuerySignOrderResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponse) SetData(v *DeveloperQuerySignOrderResponseData) *DeveloperQuerySignOrderResponse {
	s.Data = v
	return s
}

func (s *DeveloperQuerySignOrderResponse) SetErrNo(v int32) *DeveloperQuerySignOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperQuerySignOrderResponse) SetErrMsg(v string) *DeveloperQuerySignOrderResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperQuerySignOrderResponseData struct {
	AuthOrderId    *string `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CancelSource   *int64  `json:"cancel_source,omitempty" xml:"cancel_source,omitempty"`
	OpenId         *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ServiceId      *string `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	SignTime       *int64  `json:"sign_time,omitempty" xml:"sign_time,omitempty"`
	OutAuthOrderId *string `json:"out_auth_order_id,omitempty" xml:"out_auth_order_id,omitempty" require:"true"`
	NotifyUrl      *string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
}

func (s DeveloperQuerySignOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignOrderResponseData) SetAuthOrderId(v string) *DeveloperQuerySignOrderResponseData {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetStatus(v string) *DeveloperQuerySignOrderResponseData {
	s.Status = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetCancelSource(v int64) *DeveloperQuerySignOrderResponseData {
	s.CancelSource = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetOpenId(v string) *DeveloperQuerySignOrderResponseData {
	s.OpenId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetServiceId(v string) *DeveloperQuerySignOrderResponseData {
	s.ServiceId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetAppId(v string) *DeveloperQuerySignOrderResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetSignTime(v int64) *DeveloperQuerySignOrderResponseData {
	s.SignTime = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetOutAuthOrderId(v string) *DeveloperQuerySignOrderResponseData {
	s.OutAuthOrderId = &v
	return s
}

func (s *DeveloperQuerySignOrderResponseData) SetNotifyUrl(v string) *DeveloperQuerySignOrderResponseData {
	s.NotifyUrl = &v
	return s
}

type DeveloperQuerySignPayRequest struct {
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PayOrderId    *string            `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	OutPayOrderNo *string            `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty"`
}

func (s DeveloperQuerySignPayRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignPayRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignPayRequest) SetHeader(v map[string]*string) *DeveloperQuerySignPayRequest {
	s.Header = v
	return s
}

func (s *DeveloperQuerySignPayRequest) SetAccessToken(v string) *DeveloperQuerySignPayRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperQuerySignPayRequest) SetPayOrderId(v string) *DeveloperQuerySignPayRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQuerySignPayRequest) SetOutPayOrderNo(v string) *DeveloperQuerySignPayRequest {
	s.OutPayOrderNo = &v
	return s
}

type DeveloperQuerySignPayResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQuerySignPayResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperQuerySignPayResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignPayResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignPayResponse) SetLogId(v string) *DeveloperQuerySignPayResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQuerySignPayResponse) SetData(v *DeveloperQuerySignPayResponseData) *DeveloperQuerySignPayResponse {
	s.Data = v
	return s
}

func (s *DeveloperQuerySignPayResponse) SetErrNo(v int32) *DeveloperQuerySignPayResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperQuerySignPayResponse) SetErrMsg(v string) *DeveloperQuerySignPayResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperQuerySignPayResponseData struct {
	TotalAmount   *int64  `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
	UserBillPayId *string `json:"user_bill_pay_id,omitempty" xml:"user_bill_pay_id,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	ChannelPayId  *string `json:"channel_pay_id,omitempty" xml:"channel_pay_id,omitempty"`
	NotifyUrl     *string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	PayOrderId    *string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	PayTime       *int64  `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	AppId         *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	PayChannel    *int32  `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	OutPayOrderNo *string `json:"out_pay_order_no,omitempty" xml:"out_pay_order_no,omitempty" require:"true"`
	MerchantUid   *string `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	AuthOrderId   *string `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
}

func (s DeveloperQuerySignPayResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignPayResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignPayResponseData) SetTotalAmount(v int64) *DeveloperQuerySignPayResponseData {
	s.TotalAmount = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetUserBillPayId(v string) *DeveloperQuerySignPayResponseData {
	s.UserBillPayId = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetStatus(v string) *DeveloperQuerySignPayResponseData {
	s.Status = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetChannelPayId(v string) *DeveloperQuerySignPayResponseData {
	s.ChannelPayId = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetNotifyUrl(v string) *DeveloperQuerySignPayResponseData {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetPayOrderId(v string) *DeveloperQuerySignPayResponseData {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetPayTime(v int64) *DeveloperQuerySignPayResponseData {
	s.PayTime = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetAppId(v string) *DeveloperQuerySignPayResponseData {
	s.AppId = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetPayChannel(v int32) *DeveloperQuerySignPayResponseData {
	s.PayChannel = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetOutPayOrderNo(v string) *DeveloperQuerySignPayResponseData {
	s.OutPayOrderNo = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetMerchantUid(v string) *DeveloperQuerySignPayResponseData {
	s.MerchantUid = &v
	return s
}

func (s *DeveloperQuerySignPayResponseData) SetAuthOrderId(v string) *DeveloperQuerySignPayResponseData {
	s.AuthOrderId = &v
	return s
}

type DeveloperQuerySignRefundRequest struct {
	OutPayRefundNo *string            `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PayOrderId     *string            `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	PayRefundId    *string            `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty"`
}

func (s DeveloperQuerySignRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignRefundRequest) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignRefundRequest) SetOutPayRefundNo(v string) *DeveloperQuerySignRefundRequest {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperQuerySignRefundRequest) SetHeader(v map[string]*string) *DeveloperQuerySignRefundRequest {
	s.Header = v
	return s
}

func (s *DeveloperQuerySignRefundRequest) SetAccessToken(v string) *DeveloperQuerySignRefundRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperQuerySignRefundRequest) SetPayOrderId(v string) *DeveloperQuerySignRefundRequest {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQuerySignRefundRequest) SetPayRefundId(v string) *DeveloperQuerySignRefundRequest {
	s.PayRefundId = &v
	return s
}

type DeveloperQuerySignRefundResponse struct {
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperQuerySignRefundResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeveloperQuerySignRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignRefundResponse) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignRefundResponse) SetErrNo(v int32) *DeveloperQuerySignRefundResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperQuerySignRefundResponse) SetErrMsg(v string) *DeveloperQuerySignRefundResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperQuerySignRefundResponse) SetLogId(v string) *DeveloperQuerySignRefundResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperQuerySignRefundResponse) SetData(v *DeveloperQuerySignRefundResponseData) *DeveloperQuerySignRefundResponse {
	s.Data = v
	return s
}

type DeveloperQuerySignRefundResponseData struct {
	RefundList []*DeveloperQuerySignRefundResponseDataRefundListItem `json:"refund_list,omitempty" xml:"refund_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperQuerySignRefundResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignRefundResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignRefundResponseData) SetRefundList(v []*DeveloperQuerySignRefundResponseDataRefundListItem) *DeveloperQuerySignRefundResponseData {
	s.RefundList = v
	return s
}

type DeveloperQuerySignRefundResponseDataRefundListItem struct {
	RefundAt          *int64  `json:"refund_at,omitempty" xml:"refund_at,omitempty"`
	RefundReason      *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	Message           *string `json:"message,omitempty" xml:"message,omitempty"`
	PayOrderId        *string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty" require:"true"`
	RefundTotalAmount *int64  `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	OutPayRefundNo    *string `json:"out_pay_refund_no,omitempty" xml:"out_pay_refund_no,omitempty" require:"true"`
	PayRefundId       *string `json:"pay_refund_id,omitempty" xml:"pay_refund_id,omitempty" require:"true"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	NotifyUrl         *string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	AppId             *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CreateAt          *int64  `json:"create_at,omitempty" xml:"create_at,omitempty" require:"true"`
}

func (s DeveloperQuerySignRefundResponseDataRefundListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperQuerySignRefundResponseDataRefundListItem) GoString() string {
	return s.String()
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetRefundAt(v int64) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.RefundAt = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetRefundReason(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.RefundReason = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetMessage(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.Message = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetPayOrderId(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.PayOrderId = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetRefundTotalAmount(v int64) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetOutPayRefundNo(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.OutPayRefundNo = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetPayRefundId(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.PayRefundId = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetStatus(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.Status = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetNotifyUrl(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetAppId(v string) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.AppId = &v
	return s
}

func (s *DeveloperQuerySignRefundResponseDataRefundListItem) SetCreateAt(v int64) *DeveloperQuerySignRefundResponseDataRefundListItem {
	s.CreateAt = &v
	return s
}

type DeveloperRefundAuditCallbackRequest struct {
	RefundId          *string            `json:"refund_id,omitempty" xml:"refund_id,omitempty" require:"true"`
	RefundAuditStatus *int32             `json:"refund_audit_status,omitempty" xml:"refund_audit_status,omitempty" require:"true"`
	DenyMessage       *string            `json:"deny_message,omitempty" xml:"deny_message,omitempty"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperRefundAuditCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundAuditCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeveloperRefundAuditCallbackRequest) SetRefundId(v string) *DeveloperRefundAuditCallbackRequest {
	s.RefundId = &v
	return s
}

func (s *DeveloperRefundAuditCallbackRequest) SetRefundAuditStatus(v int32) *DeveloperRefundAuditCallbackRequest {
	s.RefundAuditStatus = &v
	return s
}

func (s *DeveloperRefundAuditCallbackRequest) SetDenyMessage(v string) *DeveloperRefundAuditCallbackRequest {
	s.DenyMessage = &v
	return s
}

func (s *DeveloperRefundAuditCallbackRequest) SetHeader(v map[string]*string) *DeveloperRefundAuditCallbackRequest {
	s.Header = v
	return s
}

func (s *DeveloperRefundAuditCallbackRequest) SetAccessToken(v string) *DeveloperRefundAuditCallbackRequest {
	s.AccessToken = &v
	return s
}

type DeveloperRefundAuditCallbackResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperRefundAuditCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundAuditCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeveloperRefundAuditCallbackResponse) SetErrMsg(v string) *DeveloperRefundAuditCallbackResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperRefundAuditCallbackResponse) SetLogId(v string) *DeveloperRefundAuditCallbackResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperRefundAuditCallbackResponse) SetErrNo(v int32) *DeveloperRefundAuditCallbackResponse {
	s.ErrNo = &v
	return s
}

type DeveloperRefundCreateRequest struct {
	ItemOrderDetail   []*DeveloperRefundCreateRequestItemOrderDetailItem `json:"item_order_detail,omitempty" xml:"item_order_detail,omitempty" type:"Repeated"`
	OrderId           *string                                            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RefundTotalAmount *int64                                             `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	AccessToken       *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefundReason      []*DeveloperRefundCreateRequestRefundReasonItem    `json:"refund_reason,omitempty" xml:"refund_reason,omitempty" require:"true" type:"Repeated"`
	RefundAll         *bool                                              `json:"refund_all,omitempty" xml:"refund_all,omitempty"`
	OutRefundNo       *string                                            `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty" require:"true"`
	CpExtra           *string                                            `json:"cp_extra,omitempty" xml:"cp_extra,omitempty"`
	NotifyUrl         *string                                            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	OrderEntrySchema  *DeveloperRefundCreateRequestOrderEntrySchema      `json:"order_entry_schema,omitempty" xml:"order_entry_schema,omitempty" require:"true"`
	Header            map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DeveloperRefundCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateRequest) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateRequest) SetItemOrderDetail(v []*DeveloperRefundCreateRequestItemOrderDetailItem) *DeveloperRefundCreateRequest {
	s.ItemOrderDetail = v
	return s
}

func (s *DeveloperRefundCreateRequest) SetOrderId(v string) *DeveloperRefundCreateRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetRefundTotalAmount(v int64) *DeveloperRefundCreateRequest {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetAccessToken(v string) *DeveloperRefundCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetRefundReason(v []*DeveloperRefundCreateRequestRefundReasonItem) *DeveloperRefundCreateRequest {
	s.RefundReason = v
	return s
}

func (s *DeveloperRefundCreateRequest) SetRefundAll(v bool) *DeveloperRefundCreateRequest {
	s.RefundAll = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetOutRefundNo(v string) *DeveloperRefundCreateRequest {
	s.OutRefundNo = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetCpExtra(v string) *DeveloperRefundCreateRequest {
	s.CpExtra = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetNotifyUrl(v string) *DeveloperRefundCreateRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperRefundCreateRequest) SetOrderEntrySchema(v *DeveloperRefundCreateRequestOrderEntrySchema) *DeveloperRefundCreateRequest {
	s.OrderEntrySchema = v
	return s
}

func (s *DeveloperRefundCreateRequest) SetHeader(v map[string]*string) *DeveloperRefundCreateRequest {
	s.Header = v
	return s
}

type DeveloperRefundCreateRequestItemOrderDetailItem struct {
	ItemOrderId  *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
	RefundAmount *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
}

func (s DeveloperRefundCreateRequestItemOrderDetailItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateRequestItemOrderDetailItem) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateRequestItemOrderDetailItem) SetItemOrderId(v string) *DeveloperRefundCreateRequestItemOrderDetailItem {
	s.ItemOrderId = &v
	return s
}

func (s *DeveloperRefundCreateRequestItemOrderDetailItem) SetRefundAmount(v int64) *DeveloperRefundCreateRequestItemOrderDetailItem {
	s.RefundAmount = &v
	return s
}

type DeveloperRefundCreateRequestOrderEntrySchema struct {
	Path   *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s DeveloperRefundCreateRequestOrderEntrySchema) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateRequestOrderEntrySchema) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateRequestOrderEntrySchema) SetPath(v string) *DeveloperRefundCreateRequestOrderEntrySchema {
	s.Path = &v
	return s
}

func (s *DeveloperRefundCreateRequestOrderEntrySchema) SetParams(v string) *DeveloperRefundCreateRequestOrderEntrySchema {
	s.Params = &v
	return s
}

type DeveloperRefundCreateRequestRefundReasonItem struct {
	Code *int64  `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Text *string `json:"text,omitempty" xml:"text,omitempty" require:"true"`
}

func (s DeveloperRefundCreateRequestRefundReasonItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateRequestRefundReasonItem) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateRequestRefundReasonItem) SetCode(v int64) *DeveloperRefundCreateRequestRefundReasonItem {
	s.Code = &v
	return s
}

func (s *DeveloperRefundCreateRequestRefundReasonItem) SetText(v string) *DeveloperRefundCreateRequestRefundReasonItem {
	s.Text = &v
	return s
}

type DeveloperRefundCreateResponse struct {
	Data   *DeveloperRefundCreateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeveloperRefundCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateResponse) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateResponse) SetData(v *DeveloperRefundCreateResponseData) *DeveloperRefundCreateResponse {
	s.Data = v
	return s
}

func (s *DeveloperRefundCreateResponse) SetErrNo(v int32) *DeveloperRefundCreateResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperRefundCreateResponse) SetErrMsg(v string) *DeveloperRefundCreateResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperRefundCreateResponse) SetLogId(v string) *DeveloperRefundCreateResponse {
	s.LogId = &v
	return s
}

type DeveloperRefundCreateResponseData struct {
	RefundId            *string `json:"refund_id,omitempty" xml:"refund_id,omitempty" require:"true"`
	RefundAuditDeadline *int64  `json:"refund_audit_deadline,omitempty" xml:"refund_audit_deadline,omitempty" require:"true"`
}

func (s DeveloperRefundCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundCreateResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperRefundCreateResponseData) SetRefundId(v string) *DeveloperRefundCreateResponseData {
	s.RefundId = &v
	return s
}

func (s *DeveloperRefundCreateResponseData) SetRefundAuditDeadline(v int64) *DeveloperRefundCreateResponseData {
	s.RefundAuditDeadline = &v
	return s
}

type DeveloperRefundQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefundId    *string            `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutRefundNo *string            `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DeveloperRefundQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryRequest) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryRequest) SetAccessToken(v string) *DeveloperRefundQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperRefundQueryRequest) SetRefundId(v string) *DeveloperRefundQueryRequest {
	s.RefundId = &v
	return s
}

func (s *DeveloperRefundQueryRequest) SetOrderId(v string) *DeveloperRefundQueryRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperRefundQueryRequest) SetOutRefundNo(v string) *DeveloperRefundQueryRequest {
	s.OutRefundNo = &v
	return s
}

func (s *DeveloperRefundQueryRequest) SetHeader(v map[string]*string) *DeveloperRefundQueryRequest {
	s.Header = v
	return s
}

type DeveloperRefundQueryResponse struct {
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperRefundQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeveloperRefundQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryResponse) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryResponse) SetErrNo(v int32) *DeveloperRefundQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperRefundQueryResponse) SetErrMsg(v string) *DeveloperRefundQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperRefundQueryResponse) SetLogId(v string) *DeveloperRefundQueryResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperRefundQueryResponse) SetData(v *DeveloperRefundQueryResponseData) *DeveloperRefundQueryResponse {
	s.Data = v
	return s
}

type DeveloperRefundQueryResponseData struct {
	RefundList []*DeveloperRefundQueryResponseDataRefundListItem `json:"refund_list,omitempty" xml:"refund_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperRefundQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryResponseData) SetRefundList(v []*DeveloperRefundQueryResponseDataRefundListItem) *DeveloperRefundQueryResponseData {
	s.RefundList = v
	return s
}

type DeveloperRefundQueryResponseDataRefundListItem struct {
	RefundId            *string                                                              `json:"refund_id,omitempty" xml:"refund_id,omitempty" require:"true"`
	RefundStatus        *string                                                              `json:"refund_status,omitempty" xml:"refund_status,omitempty" require:"true"`
	CreateAt            *int64                                                               `json:"create_at,omitempty" xml:"create_at,omitempty" require:"true"`
	RefundSource        *int32                                                               `json:"refund_source,omitempty" xml:"refund_source,omitempty"`
	MerchantAuditDetail *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail   `json:"merchant_audit_detail,omitempty" xml:"merchant_audit_detail,omitempty"`
	ItemOrderDetail     []*DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem `json:"item_order_detail,omitempty" xml:"item_order_detail,omitempty" type:"Repeated"`
	OrderId             *string                                                              `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RefundAt            *int64                                                               `json:"refund_at,omitempty" xml:"refund_at,omitempty"`
	Message             *string                                                              `json:"message,omitempty" xml:"message,omitempty"`
	RefundTotalAmount   *int64                                                               `json:"refund_total_amount,omitempty" xml:"refund_total_amount,omitempty" require:"true"`
	OutRefundNo         *string                                                              `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
}

func (s DeveloperRefundQueryResponseDataRefundListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryResponseDataRefundListItem) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetRefundId(v string) *DeveloperRefundQueryResponseDataRefundListItem {
	s.RefundId = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetRefundStatus(v string) *DeveloperRefundQueryResponseDataRefundListItem {
	s.RefundStatus = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetCreateAt(v int64) *DeveloperRefundQueryResponseDataRefundListItem {
	s.CreateAt = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetRefundSource(v int32) *DeveloperRefundQueryResponseDataRefundListItem {
	s.RefundSource = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetMerchantAuditDetail(v *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) *DeveloperRefundQueryResponseDataRefundListItem {
	s.MerchantAuditDetail = v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetItemOrderDetail(v []*DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem) *DeveloperRefundQueryResponseDataRefundListItem {
	s.ItemOrderDetail = v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetOrderId(v string) *DeveloperRefundQueryResponseDataRefundListItem {
	s.OrderId = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetRefundAt(v int64) *DeveloperRefundQueryResponseDataRefundListItem {
	s.RefundAt = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetMessage(v string) *DeveloperRefundQueryResponseDataRefundListItem {
	s.Message = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetRefundTotalAmount(v int64) *DeveloperRefundQueryResponseDataRefundListItem {
	s.RefundTotalAmount = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItem) SetOutRefundNo(v string) *DeveloperRefundQueryResponseDataRefundListItem {
	s.OutRefundNo = &v
	return s
}

type DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem struct {
	RefundAmount *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
	ItemOrderId  *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
}

func (s DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem) SetRefundAmount(v int64) *DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem {
	s.RefundAmount = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem) SetItemOrderId(v string) *DeveloperRefundQueryResponseDataRefundListItemItemOrderDetailItem {
	s.ItemOrderId = &v
	return s
}

type DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail struct {
	RefundAuditDeadline *int64  `json:"refund_audit_deadline,omitempty" xml:"refund_audit_deadline,omitempty" require:"true"`
	AuditStatus         *string `json:"audit_status,omitempty" xml:"audit_status,omitempty" require:"true"`
	DenyMessage         *string `json:"deny_message,omitempty" xml:"deny_message,omitempty"`
	NeedRefundAudit     *int64  `json:"need_refund_audit,omitempty" xml:"need_refund_audit,omitempty" require:"true"`
}

func (s DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) String() string {
	return tea.Prettify(s)
}

func (s DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) GoString() string {
	return s.String()
}

func (s *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) SetRefundAuditDeadline(v int64) *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail {
	s.RefundAuditDeadline = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) SetAuditStatus(v string) *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail {
	s.AuditStatus = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) SetDenyMessage(v string) *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail {
	s.DenyMessage = &v
	return s
}

func (s *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail) SetNeedRefundAudit(v int64) *DeveloperRefundQueryResponseDataRefundListItemMerchantAuditDetail {
	s.NeedRefundAudit = &v
	return s
}

type DeveloperSettleCreateRequest struct {
	Ext          *string            `json:"ext,omitempty" xml:"ext,omitempty"`
	OutOrderNo   *string            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	SettleParams *string            `json:"settle_params,omitempty" xml:"settle_params,omitempty"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	OutSettleNo  *string            `json:"out_settle_no,omitempty" xml:"out_settle_no,omitempty" require:"true"`
	NotifyUrl    *string            `json:"notify_url,omitempty" xml:"notify_url,omitempty"`
	SettleDesc   *string            `json:"settle_desc,omitempty" xml:"settle_desc,omitempty" require:"true"`
	ItemOrderId  *string            `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperSettleCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleCreateRequest) GoString() string {
	return s.String()
}

func (s *DeveloperSettleCreateRequest) SetExt(v string) *DeveloperSettleCreateRequest {
	s.Ext = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetOutOrderNo(v string) *DeveloperSettleCreateRequest {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetSettleParams(v string) *DeveloperSettleCreateRequest {
	s.SettleParams = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetAppId(v string) *DeveloperSettleCreateRequest {
	s.AppId = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetOutSettleNo(v string) *DeveloperSettleCreateRequest {
	s.OutSettleNo = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetNotifyUrl(v string) *DeveloperSettleCreateRequest {
	s.NotifyUrl = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetSettleDesc(v string) *DeveloperSettleCreateRequest {
	s.SettleDesc = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetItemOrderId(v string) *DeveloperSettleCreateRequest {
	s.ItemOrderId = &v
	return s
}

func (s *DeveloperSettleCreateRequest) SetHeader(v map[string]*string) *DeveloperSettleCreateRequest {
	s.Header = v
	return s
}

func (s *DeveloperSettleCreateRequest) SetAccessToken(v string) *DeveloperSettleCreateRequest {
	s.AccessToken = &v
	return s
}

type DeveloperSettleCreateResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperSettleCreateResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeveloperSettleCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleCreateResponse) GoString() string {
	return s.String()
}

func (s *DeveloperSettleCreateResponse) SetLogId(v string) *DeveloperSettleCreateResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperSettleCreateResponse) SetData(v *DeveloperSettleCreateResponseData) *DeveloperSettleCreateResponse {
	s.Data = v
	return s
}

func (s *DeveloperSettleCreateResponse) SetErrNo(v int32) *DeveloperSettleCreateResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperSettleCreateResponse) SetErrMsg(v string) *DeveloperSettleCreateResponse {
	s.ErrMsg = &v
	return s
}

type DeveloperSettleCreateResponseData struct {
	SettleId       *string `json:"settle_id,omitempty" xml:"settle_id,omitempty" require:"true"`
	WalletSettleId *string `json:"wallet_settle_id,omitempty" xml:"wallet_settle_id,omitempty" require:"true"`
}

func (s DeveloperSettleCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleCreateResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperSettleCreateResponseData) SetSettleId(v string) *DeveloperSettleCreateResponseData {
	s.SettleId = &v
	return s
}

func (s *DeveloperSettleCreateResponseData) SetWalletSettleId(v string) *DeveloperSettleCreateResponseData {
	s.WalletSettleId = &v
	return s
}

type DeveloperSettleQueryRequest struct {
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	SettleId    *string            `json:"settle_id,omitempty" xml:"settle_id,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	OutOrderNo  *string            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	OutSettleNo *string            `json:"out_settle_no,omitempty" xml:"out_settle_no,omitempty"`
}

func (s DeveloperSettleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleQueryRequest) GoString() string {
	return s.String()
}

func (s *DeveloperSettleQueryRequest) SetOrderId(v string) *DeveloperSettleQueryRequest {
	s.OrderId = &v
	return s
}

func (s *DeveloperSettleQueryRequest) SetHeader(v map[string]*string) *DeveloperSettleQueryRequest {
	s.Header = v
	return s
}

func (s *DeveloperSettleQueryRequest) SetAccessToken(v string) *DeveloperSettleQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *DeveloperSettleQueryRequest) SetSettleId(v string) *DeveloperSettleQueryRequest {
	s.SettleId = &v
	return s
}

func (s *DeveloperSettleQueryRequest) SetAppId(v string) *DeveloperSettleQueryRequest {
	s.AppId = &v
	return s
}

func (s *DeveloperSettleQueryRequest) SetOutOrderNo(v string) *DeveloperSettleQueryRequest {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperSettleQueryRequest) SetOutSettleNo(v string) *DeveloperSettleQueryRequest {
	s.OutSettleNo = &v
	return s
}

type DeveloperSettleQueryResponse struct {
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   []*DeveloperSettleQueryResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DeveloperSettleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleQueryResponse) GoString() string {
	return s.String()
}

func (s *DeveloperSettleQueryResponse) SetErrNo(v int32) *DeveloperSettleQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperSettleQueryResponse) SetErrMsg(v string) *DeveloperSettleQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperSettleQueryResponse) SetLogId(v string) *DeveloperSettleQueryResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperSettleQueryResponse) SetData(v []*DeveloperSettleQueryResponseDataItem) *DeveloperSettleQueryResponse {
	s.Data = v
	return s
}

type DeveloperSettleQueryResponseDataItem struct {
	OutOrderNo     *string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	OutSettleNo    *string `json:"out_settle_no,omitempty" xml:"out_settle_no,omitempty" require:"true"`
	CpExtra        *string `json:"cp_extra,omitempty" xml:"cp_extra,omitempty" require:"true"`
	Commission     *int64  `json:"commission,omitempty" xml:"commission,omitempty" require:"true"`
	SettleId       *string `json:"settle_id,omitempty" xml:"settle_id,omitempty" require:"true"`
	SettleDetail   *string `json:"settle_detail,omitempty" xml:"settle_detail,omitempty" require:"true"`
	SettleAt       *int64  `json:"settle_at,omitempty" xml:"settle_at,omitempty"`
	Rake           *int64  `json:"rake,omitempty" xml:"rake,omitempty" require:"true"`
	SettleAmount   *int64  `json:"settle_amount,omitempty" xml:"settle_amount,omitempty" require:"true"`
	SettleStatus   *string `json:"settle_status,omitempty" xml:"settle_status,omitempty" require:"true"`
	OrderId        *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ItemOrderId    *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
	PlatformTicket *int64  `json:"platform_ticket,omitempty" xml:"platform_ticket,omitempty"`
}

func (s DeveloperSettleQueryResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperSettleQueryResponseDataItem) GoString() string {
	return s.String()
}

func (s *DeveloperSettleQueryResponseDataItem) SetOutOrderNo(v string) *DeveloperSettleQueryResponseDataItem {
	s.OutOrderNo = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetOutSettleNo(v string) *DeveloperSettleQueryResponseDataItem {
	s.OutSettleNo = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetCpExtra(v string) *DeveloperSettleQueryResponseDataItem {
	s.CpExtra = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetCommission(v int64) *DeveloperSettleQueryResponseDataItem {
	s.Commission = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetSettleId(v string) *DeveloperSettleQueryResponseDataItem {
	s.SettleId = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetSettleDetail(v string) *DeveloperSettleQueryResponseDataItem {
	s.SettleDetail = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetSettleAt(v int64) *DeveloperSettleQueryResponseDataItem {
	s.SettleAt = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetRake(v int64) *DeveloperSettleQueryResponseDataItem {
	s.Rake = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetSettleAmount(v int64) *DeveloperSettleQueryResponseDataItem {
	s.SettleAmount = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetSettleStatus(v string) *DeveloperSettleQueryResponseDataItem {
	s.SettleStatus = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetOrderId(v string) *DeveloperSettleQueryResponseDataItem {
	s.OrderId = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetItemOrderId(v string) *DeveloperSettleQueryResponseDataItem {
	s.ItemOrderId = &v
	return s
}

func (s *DeveloperSettleQueryResponseDataItem) SetPlatformTicket(v int64) *DeveloperSettleQueryResponseDataItem {
	s.PlatformTicket = &v
	return s
}

type DeveloperTagQueryRequest struct {
	GoodsType   *int               `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperTagQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryRequest) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryRequest) SetGoodsType(v int) *DeveloperTagQueryRequest {
	s.GoodsType = &v
	return s
}

func (s *DeveloperTagQueryRequest) SetHeader(v map[string]*string) *DeveloperTagQueryRequest {
	s.Header = v
	return s
}

func (s *DeveloperTagQueryRequest) SetAccessToken(v string) *DeveloperTagQueryRequest {
	s.AccessToken = &v
	return s
}

type DeveloperTagQueryResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DeveloperTagQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DeveloperTagQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponse) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponse) SetErrNo(v int32) *DeveloperTagQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *DeveloperTagQueryResponse) SetErrMsg(v string) *DeveloperTagQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperTagQueryResponse) SetLogId(v string) *DeveloperTagQueryResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperTagQueryResponse) SetData(v *DeveloperTagQueryResponseData) *DeveloperTagQueryResponse {
	s.Data = v
	return s
}

type DeveloperTagQueryResponseData struct {
	TagDetailList []*DeveloperTagQueryResponseDataTagDetailListItem `json:"tag_detail_list,omitempty" xml:"tag_detail_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperTagQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponseData) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponseData) SetTagDetailList(v []*DeveloperTagQueryResponseDataTagDetailListItem) *DeveloperTagQueryResponseData {
	s.TagDetailList = v
	return s
}

type DeveloperTagQueryResponseDataTagDetailListItem struct {
	RuleList   []*DeveloperTagQueryResponseDataTagDetailListItemRuleListItem `json:"rule_list,omitempty" xml:"rule_list,omitempty" require:"true" type:"Repeated"`
	TagGroupId *string                                                       `json:"tag_group_id,omitempty" xml:"tag_group_id,omitempty" require:"true"`
	TagList    []*DeveloperTagQueryResponseDataTagDetailListItemTagListItem  `json:"tag_list,omitempty" xml:"tag_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperTagQueryResponseDataTagDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponseDataTagDetailListItem) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponseDataTagDetailListItem) SetRuleList(v []*DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) *DeveloperTagQueryResponseDataTagDetailListItem {
	s.RuleList = v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItem) SetTagGroupId(v string) *DeveloperTagQueryResponseDataTagDetailListItem {
	s.TagGroupId = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItem) SetTagList(v []*DeveloperTagQueryResponseDataTagDetailListItemTagListItem) *DeveloperTagQueryResponseDataTagDetailListItem {
	s.TagList = v
	return s
}

type DeveloperTagQueryResponseDataTagDetailListItemRuleListItem struct {
	RuleType   *int                                                                  `json:"rule_type,omitempty" xml:"rule_type,omitempty" require:"true"`
	RefundRule *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule `json:"refund_rule,omitempty" xml:"refund_rule,omitempty"`
	RuleId     *string                                                               `json:"rule_id,omitempty" xml:"rule_id,omitempty" require:"true"`
}

func (s DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) SetRuleType(v int) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem {
	s.RuleType = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) SetRefundRule(v *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem {
	s.RefundRule = v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem) SetRuleId(v string) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItem {
	s.RuleId = &v
	return s
}

type DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule struct {
	RefundType        *int    `json:"refund_type,omitempty" xml:"refund_type,omitempty" require:"true"`
	MerchantAuditType *int    `json:"merchant_audit_type,omitempty" xml:"merchant_audit_type,omitempty" require:"true"`
	BizStatus         *string `json:"biz_status,omitempty" xml:"biz_status,omitempty" require:"true"`
}

func (s DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) SetRefundType(v int) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule {
	s.RefundType = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) SetMerchantAuditType(v int) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule {
	s.MerchantAuditType = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule) SetBizStatus(v string) *DeveloperTagQueryResponseDataTagDetailListItemRuleListItemRefundRule {
	s.BizStatus = &v
	return s
}

type DeveloperTagQueryResponseDataTagDetailListItemTagListItem struct {
	TagType   *int      `json:"tag_type,omitempty" xml:"tag_type,omitempty" require:"true"`
	BizLine   *int      `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	GoodsType *int      `json:"goods_type,omitempty" xml:"goods_type,omitempty" require:"true"`
	TagId     *string   `json:"tag_id,omitempty" xml:"tag_id,omitempty" require:"true"`
	TagName   *string   `json:"tag_name,omitempty" xml:"tag_name,omitempty" require:"true"`
	TagDescs  []*string `json:"tag_descs,omitempty" xml:"tag_descs,omitempty" require:"true" type:"Repeated"`
}

func (s DeveloperTagQueryResponseDataTagDetailListItemTagListItem) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTagQueryResponseDataTagDetailListItemTagListItem) GoString() string {
	return s.String()
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetTagType(v int) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.TagType = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetBizLine(v int) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.BizLine = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetGoodsType(v int) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.GoodsType = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetTagId(v string) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.TagId = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetTagName(v string) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.TagName = &v
	return s
}

func (s *DeveloperTagQueryResponseDataTagDetailListItemTagListItem) SetTagDescs(v []*string) *DeveloperTagQueryResponseDataTagDetailListItemTagListItem {
	s.TagDescs = v
	return s
}

type DeveloperTerminateSignRequest struct {
	AuthOrderId *string            `json:"auth_order_id,omitempty" xml:"auth_order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeveloperTerminateSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTerminateSignRequest) GoString() string {
	return s.String()
}

func (s *DeveloperTerminateSignRequest) SetAuthOrderId(v string) *DeveloperTerminateSignRequest {
	s.AuthOrderId = &v
	return s
}

func (s *DeveloperTerminateSignRequest) SetHeader(v map[string]*string) *DeveloperTerminateSignRequest {
	s.Header = v
	return s
}

func (s *DeveloperTerminateSignRequest) SetAccessToken(v string) *DeveloperTerminateSignRequest {
	s.AccessToken = &v
	return s
}

type DeveloperTerminateSignResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeveloperTerminateSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DeveloperTerminateSignResponse) GoString() string {
	return s.String()
}

func (s *DeveloperTerminateSignResponse) SetErrMsg(v string) *DeveloperTerminateSignResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeveloperTerminateSignResponse) SetLogId(v string) *DeveloperTerminateSignResponse {
	s.LogId = &v
	return s
}

func (s *DeveloperTerminateSignResponse) SetErrNo(v int32) *DeveloperTerminateSignResponse {
	s.ErrNo = &v
	return s
}

type DevtoolGetMountLegalRequest struct {
	MicappId    *string            `json:"micapp_id,omitempty" xml:"micapp_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DevtoolGetMountLegalRequest) String() string {
	return tea.Prettify(s)
}

func (s DevtoolGetMountLegalRequest) GoString() string {
	return s.String()
}

func (s *DevtoolGetMountLegalRequest) SetMicappId(v string) *DevtoolGetMountLegalRequest {
	s.MicappId = &v
	return s
}

func (s *DevtoolGetMountLegalRequest) SetHeader(v map[string]*string) *DevtoolGetMountLegalRequest {
	s.Header = v
	return s
}

func (s *DevtoolGetMountLegalRequest) SetAccessToken(v string) *DevtoolGetMountLegalRequest {
	s.AccessToken = &v
	return s
}

type DevtoolGetMountLegalResponse struct {
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DevtoolGetMountLegalResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DevtoolGetMountLegalResponse) String() string {
	return tea.Prettify(s)
}

func (s DevtoolGetMountLegalResponse) GoString() string {
	return s.String()
}

func (s *DevtoolGetMountLegalResponse) SetErrMsg(v string) *DevtoolGetMountLegalResponse {
	s.ErrMsg = &v
	return s
}

func (s *DevtoolGetMountLegalResponse) SetLogId(v string) *DevtoolGetMountLegalResponse {
	s.LogId = &v
	return s
}

func (s *DevtoolGetMountLegalResponse) SetData(v *DevtoolGetMountLegalResponseData) *DevtoolGetMountLegalResponse {
	s.Data = v
	return s
}

func (s *DevtoolGetMountLegalResponse) SetErrNo(v int32) *DevtoolGetMountLegalResponse {
	s.ErrNo = &v
	return s
}

type DevtoolGetMountLegalResponseData struct {
	Extra *DevtoolGetMountLegalResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *DevtoolGetMountLegalResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DevtoolGetMountLegalResponseData) String() string {
	return tea.Prettify(s)
}

func (s DevtoolGetMountLegalResponseData) GoString() string {
	return s.String()
}

func (s *DevtoolGetMountLegalResponseData) SetExtra(v *DevtoolGetMountLegalResponseDataExtra) *DevtoolGetMountLegalResponseData {
	s.Extra = v
	return s
}

func (s *DevtoolGetMountLegalResponseData) SetData(v *DevtoolGetMountLegalResponseDataData) *DevtoolGetMountLegalResponseData {
	s.Data = v
	return s
}

type DevtoolGetMountLegalResponseDataData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	IsLegal     *bool   `json:"is_legal,omitempty" xml:"is_legal,omitempty" require:"true"`
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DevtoolGetMountLegalResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s DevtoolGetMountLegalResponseDataData) GoString() string {
	return s.String()
}

func (s *DevtoolGetMountLegalResponseDataData) SetDescription(v string) *DevtoolGetMountLegalResponseDataData {
	s.Description = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataData) SetIsLegal(v bool) *DevtoolGetMountLegalResponseDataData {
	s.IsLegal = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataData) SetErrorCode(v int64) *DevtoolGetMountLegalResponseDataData {
	s.ErrorCode = &v
	return s
}

type DevtoolGetMountLegalResponseDataExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DevtoolGetMountLegalResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s DevtoolGetMountLegalResponseDataExtra) GoString() string {
	return s.String()
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetDescription(v string) *DevtoolGetMountLegalResponseDataExtra {
	s.Description = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetSubErrorCode(v int32) *DevtoolGetMountLegalResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetSubDescription(v string) *DevtoolGetMountLegalResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetLogid(v string) *DevtoolGetMountLegalResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetNow(v int64) *DevtoolGetMountLegalResponseDataExtra {
	s.Now = &v
	return s
}

func (s *DevtoolGetMountLegalResponseDataExtra) SetErrorCode(v int32) *DevtoolGetMountLegalResponseDataExtra {
	s.ErrorCode = &v
	return s
}

type DishBindGetRequest struct {
	ProductId   *int64                  `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	Base        *DishBindGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PageNo      *int32                  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32                  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Header      map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiId       *int64                  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DishBindGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetRequest) GoString() string {
	return s.String()
}

func (s *DishBindGetRequest) SetProductId(v int64) *DishBindGetRequest {
	s.ProductId = &v
	return s
}

func (s *DishBindGetRequest) SetBase(v *DishBindGetRequestBase) *DishBindGetRequest {
	s.Base = v
	return s
}

func (s *DishBindGetRequest) SetAccountId(v string) *DishBindGetRequest {
	s.AccountId = &v
	return s
}

func (s *DishBindGetRequest) SetPageNo(v int32) *DishBindGetRequest {
	s.PageNo = &v
	return s
}

func (s *DishBindGetRequest) SetPageSize(v int32) *DishBindGetRequest {
	s.PageSize = &v
	return s
}

func (s *DishBindGetRequest) SetHeader(v map[string]*string) *DishBindGetRequest {
	s.Header = v
	return s
}

func (s *DishBindGetRequest) SetAccessToken(v string) *DishBindGetRequest {
	s.AccessToken = &v
	return s
}

func (s *DishBindGetRequest) SetPoiId(v int64) *DishBindGetRequest {
	s.PoiId = &v
	return s
}

type DishBindGetRequestBase struct {
	LogID      *string                           `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishBindGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                           `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                           `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                           `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishBindGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetRequestBase) GoString() string {
	return s.String()
}

func (s *DishBindGetRequestBase) SetLogID(v string) *DishBindGetRequestBase {
	s.LogID = &v
	return s
}

func (s *DishBindGetRequestBase) SetTrafficEnv(v *DishBindGetRequestBaseTrafficEnv) *DishBindGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishBindGetRequestBase) SetAddr(v string) *DishBindGetRequestBase {
	s.Addr = &v
	return s
}

func (s *DishBindGetRequestBase) SetCaller(v string) *DishBindGetRequestBase {
	s.Caller = &v
	return s
}

func (s *DishBindGetRequestBase) SetClient(v string) *DishBindGetRequestBase {
	s.Client = &v
	return s
}

func (s *DishBindGetRequestBase) SetExtra(v map[string]*string) *DishBindGetRequestBase {
	s.Extra = v
	return s
}

type DishBindGetRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishBindGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishBindGetRequestBaseTrafficEnv) SetEnv(v string) *DishBindGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishBindGetRequestBaseTrafficEnv) SetOpen(v bool) *DishBindGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishBindGetResponse struct {
	Data     *DishBindGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishBindGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishBindGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s DishBindGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetResponse) GoString() string {
	return s.String()
}

func (s *DishBindGetResponse) SetData(v *DishBindGetResponseData) *DishBindGetResponse {
	s.Data = v
	return s
}

func (s *DishBindGetResponse) SetExtra(v *DishBindGetResponseExtra) *DishBindGetResponse {
	s.Extra = v
	return s
}

func (s *DishBindGetResponse) SetBaseResp(v *DishBindGetResponseBaseResp) *DishBindGetResponse {
	s.BaseResp = v
	return s
}

type DishBindGetResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishBindGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishBindGetResponseBaseResp) SetStatusCode(v int32) *DishBindGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishBindGetResponseBaseResp) SetStatusMessage(v string) *DishBindGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishBindGetResponseBaseResp) SetExtra(v map[string]*string) *DishBindGetResponseBaseResp {
	s.Extra = v
	return s
}

type DishBindGetResponseData struct {
	HasMore     *bool                                     `json:"has_more,omitempty" xml:"has_more,omitempty"`
	ProductList []*DishBindGetResponseDataProductListItem `json:"product_list,omitempty" xml:"product_list,omitempty" type:"Repeated"`
	Total       *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
	Description *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishBindGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetResponseData) GoString() string {
	return s.String()
}

func (s *DishBindGetResponseData) SetHasMore(v bool) *DishBindGetResponseData {
	s.HasMore = &v
	return s
}

func (s *DishBindGetResponseData) SetProductList(v []*DishBindGetResponseDataProductListItem) *DishBindGetResponseData {
	s.ProductList = v
	return s
}

func (s *DishBindGetResponseData) SetTotal(v int32) *DishBindGetResponseData {
	s.Total = &v
	return s
}

func (s *DishBindGetResponseData) SetDescription(v string) *DishBindGetResponseData {
	s.Description = &v
	return s
}

func (s *DishBindGetResponseData) SetErrorCode(v int32) *DishBindGetResponseData {
	s.ErrorCode = &v
	return s
}

type DishBindGetResponseDataProductListItem struct {
	StoreProductId *int64           `json:"store_product_id,omitempty" xml:"store_product_id,omitempty"`
	PoiId          *int64           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SkuMapping     map[int64]*int64 `json:"sku_mapping,omitempty" xml:"sku_mapping,omitempty"`
}

func (s DishBindGetResponseDataProductListItem) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetResponseDataProductListItem) GoString() string {
	return s.String()
}

func (s *DishBindGetResponseDataProductListItem) SetStoreProductId(v int64) *DishBindGetResponseDataProductListItem {
	s.StoreProductId = &v
	return s
}

func (s *DishBindGetResponseDataProductListItem) SetPoiId(v int64) *DishBindGetResponseDataProductListItem {
	s.PoiId = &v
	return s
}

func (s *DishBindGetResponseDataProductListItem) SetSkuMapping(v map[int64]*int64) *DishBindGetResponseDataProductListItem {
	s.SkuMapping = v
	return s
}

type DishBindGetResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishBindGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishBindGetResponseExtra) GoString() string {
	return s.String()
}

func (s *DishBindGetResponseExtra) SetLogid(v string) *DishBindGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishBindGetResponseExtra) SetNow(v int64) *DishBindGetResponseExtra {
	s.Now = &v
	return s
}

func (s *DishBindGetResponseExtra) SetSubDescription(v string) *DishBindGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishBindGetResponseExtra) SetSubErrorCode(v int32) *DishBindGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishBindGetResponseExtra) SetDescription(v string) *DishBindGetResponseExtra {
	s.Description = &v
	return s
}

func (s *DishBindGetResponseExtra) SetErrorCode(v int32) *DishBindGetResponseExtra {
	s.ErrorCode = &v
	return s
}

type DishDraftGetRequest struct {
	DishType    *int                     `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	PoiId       *string                  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ProductIds  []*int64                 `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base        *DishDraftGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DishDraftGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetRequest) GoString() string {
	return s.String()
}

func (s *DishDraftGetRequest) SetDishType(v int) *DishDraftGetRequest {
	s.DishType = &v
	return s
}

func (s *DishDraftGetRequest) SetPoiId(v string) *DishDraftGetRequest {
	s.PoiId = &v
	return s
}

func (s *DishDraftGetRequest) SetProductIds(v []*int64) *DishDraftGetRequest {
	s.ProductIds = v
	return s
}

func (s *DishDraftGetRequest) SetBase(v *DishDraftGetRequestBase) *DishDraftGetRequest {
	s.Base = v
	return s
}

func (s *DishDraftGetRequest) SetAccountId(v string) *DishDraftGetRequest {
	s.AccountId = &v
	return s
}

func (s *DishDraftGetRequest) SetHeader(v map[string]*string) *DishDraftGetRequest {
	s.Header = v
	return s
}

func (s *DishDraftGetRequest) SetAccessToken(v string) *DishDraftGetRequest {
	s.AccessToken = &v
	return s
}

type DishDraftGetRequestBase struct {
	Extra      map[string]*string                 `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                            `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishDraftGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                            `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                            `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                            `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s DishDraftGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetRequestBase) GoString() string {
	return s.String()
}

func (s *DishDraftGetRequestBase) SetExtra(v map[string]*string) *DishDraftGetRequestBase {
	s.Extra = v
	return s
}

func (s *DishDraftGetRequestBase) SetLogID(v string) *DishDraftGetRequestBase {
	s.LogID = &v
	return s
}

func (s *DishDraftGetRequestBase) SetTrafficEnv(v *DishDraftGetRequestBaseTrafficEnv) *DishDraftGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishDraftGetRequestBase) SetAddr(v string) *DishDraftGetRequestBase {
	s.Addr = &v
	return s
}

func (s *DishDraftGetRequestBase) SetCaller(v string) *DishDraftGetRequestBase {
	s.Caller = &v
	return s
}

func (s *DishDraftGetRequestBase) SetClient(v string) *DishDraftGetRequestBase {
	s.Client = &v
	return s
}

type DishDraftGetRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s DishDraftGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishDraftGetRequestBaseTrafficEnv) SetOpen(v bool) *DishDraftGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *DishDraftGetRequestBaseTrafficEnv) SetEnv(v string) *DishDraftGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type DishDraftGetResponse struct {
	Data     *DishDraftGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishDraftGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishDraftGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s DishDraftGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponse) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponse) SetData(v *DishDraftGetResponseData) *DishDraftGetResponse {
	s.Data = v
	return s
}

func (s *DishDraftGetResponse) SetExtra(v *DishDraftGetResponseExtra) *DishDraftGetResponse {
	s.Extra = v
	return s
}

func (s *DishDraftGetResponse) SetBaseResp(v *DishDraftGetResponseBaseResp) *DishDraftGetResponse {
	s.BaseResp = v
	return s
}

type DishDraftGetResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishDraftGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseBaseResp) SetStatusCode(v int32) *DishDraftGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishDraftGetResponseBaseResp) SetStatusMessage(v string) *DishDraftGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishDraftGetResponseBaseResp) SetExtra(v map[string]*string) *DishDraftGetResponseBaseResp {
	s.Extra = v
	return s
}

type DishDraftGetResponseData struct {
	Description *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Dishs       []*DishDraftGetResponseDataDishsItem `json:"dishs,omitempty" xml:"dishs,omitempty" type:"Repeated"`
	ErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishDraftGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseData) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseData) SetDescription(v string) *DishDraftGetResponseData {
	s.Description = &v
	return s
}

func (s *DishDraftGetResponseData) SetDishs(v []*DishDraftGetResponseDataDishsItem) *DishDraftGetResponseData {
	s.Dishs = v
	return s
}

func (s *DishDraftGetResponseData) SetErrorCode(v int32) *DishDraftGetResponseData {
	s.ErrorCode = &v
	return s
}

type DishDraftGetResponseDataDishsItem struct {
	Attributes        []*DishDraftGetResponseDataDishsItemAttributesItem       `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	DishDescription   *string                                                  `json:"dish_description,omitempty" xml:"dish_description,omitempty"`
	DishDetailInfo    *DishDraftGetResponseDataDishsItemDishDetailInfo         `json:"dish_detail_info,omitempty" xml:"dish_detail_info,omitempty"`
	OutId             *string                                                  `json:"out_id,omitempty" xml:"out_id,omitempty"`
	UpdateTime        *int64                                                   `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductSpecAttrs  []*DishDraftGetResponseDataDishsItemProductSpecAttrsItem `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	DishType          *int                                                     `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	ImageList         []*DishDraftGetResponseDataDishsItemImageListItem        `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	IsBindMerchant    *bool                                                    `json:"is_bind_merchant,omitempty" xml:"is_bind_merchant,omitempty"`
	CreateTime        *int64                                                   `json:"create_time,omitempty" xml:"create_time,omitempty"`
	PoiCount          *int64                                                   `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	MerchantProductId *string                                                  `json:"merchant_product_id,omitempty" xml:"merchant_product_id,omitempty"`
	AddDishGroups     []*DishDraftGetResponseDataDishsItemAddDishGroupsItem    `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	Skus              []*DishDraftGetResponseDataDishsItemSkusItem             `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	CategoryId        *int64                                                   `json:"category_id,omitempty" xml:"category_id,omitempty"`
	ApplyDate         *DishDraftGetResponseDataDishsItemApplyDate              `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	ProductName       *string                                                  `json:"product_name,omitempty" xml:"product_name,omitempty"`
	DishGroups        []*DishDraftGetResponseDataDishsItemDishGroupsItem       `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
	Pois              []*DishDraftGetResponseDataDishsItemPoisItem             `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	SettleType        *int64                                                   `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	AccountId         *string                                                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DraftStatus       *int                                                     `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	ProductId         *int64                                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	DeliveryMethod    []*int                                                   `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
}

func (s DishDraftGetResponseDataDishsItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItem) SetAttributes(v []*DishDraftGetResponseDataDishsItemAttributesItem) *DishDraftGetResponseDataDishsItem {
	s.Attributes = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDishDescription(v string) *DishDraftGetResponseDataDishsItem {
	s.DishDescription = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDishDetailInfo(v *DishDraftGetResponseDataDishsItemDishDetailInfo) *DishDraftGetResponseDataDishsItem {
	s.DishDetailInfo = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetOutId(v string) *DishDraftGetResponseDataDishsItem {
	s.OutId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetUpdateTime(v int64) *DishDraftGetResponseDataDishsItem {
	s.UpdateTime = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetProductSpecAttrs(v []*DishDraftGetResponseDataDishsItemProductSpecAttrsItem) *DishDraftGetResponseDataDishsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDishType(v int) *DishDraftGetResponseDataDishsItem {
	s.DishType = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetImageList(v []*DishDraftGetResponseDataDishsItemImageListItem) *DishDraftGetResponseDataDishsItem {
	s.ImageList = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetIsBindMerchant(v bool) *DishDraftGetResponseDataDishsItem {
	s.IsBindMerchant = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetCreateTime(v int64) *DishDraftGetResponseDataDishsItem {
	s.CreateTime = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetPoiCount(v int64) *DishDraftGetResponseDataDishsItem {
	s.PoiCount = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetMerchantProductId(v string) *DishDraftGetResponseDataDishsItem {
	s.MerchantProductId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetAddDishGroups(v []*DishDraftGetResponseDataDishsItemAddDishGroupsItem) *DishDraftGetResponseDataDishsItem {
	s.AddDishGroups = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetSkus(v []*DishDraftGetResponseDataDishsItemSkusItem) *DishDraftGetResponseDataDishsItem {
	s.Skus = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetCategoryId(v int64) *DishDraftGetResponseDataDishsItem {
	s.CategoryId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetApplyDate(v *DishDraftGetResponseDataDishsItemApplyDate) *DishDraftGetResponseDataDishsItem {
	s.ApplyDate = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetProductName(v string) *DishDraftGetResponseDataDishsItem {
	s.ProductName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDishGroups(v []*DishDraftGetResponseDataDishsItemDishGroupsItem) *DishDraftGetResponseDataDishsItem {
	s.DishGroups = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetPois(v []*DishDraftGetResponseDataDishsItemPoisItem) *DishDraftGetResponseDataDishsItem {
	s.Pois = v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetSettleType(v int64) *DishDraftGetResponseDataDishsItem {
	s.SettleType = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetAccountId(v string) *DishDraftGetResponseDataDishsItem {
	s.AccountId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDraftStatus(v int) *DishDraftGetResponseDataDishsItem {
	s.DraftStatus = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetProductId(v int64) *DishDraftGetResponseDataDishsItem {
	s.ProductId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItem) SetDeliveryMethod(v []*int) *DishDraftGetResponseDataDishsItem {
	s.DeliveryMethod = v
	return s
}

type DishDraftGetResponseDataDishsItemAddDishGroupsItem struct {
	GroupId   *int64                                                            `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                           `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItem) SetGroupId(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItem) SetGroupName(v string) *DishDraftGetResponseDataDishsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItem) SetItemList(v []*DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) *DishDraftGetResponseDataDishsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

type DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem struct {
	Price               *int64                                                                 `json:"price,omitempty" xml:"price,omitempty"`
	IsSetAutoComplement *bool                                                                  `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	ProductId           *int64                                                                 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ResidueStock        *int64                                                                 `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	OutId               *string                                                                `json:"out_id,omitempty" xml:"out_id,omitempty"`
	PackFee             *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	SkuId               *int64                                                                 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	IsSetSellOut        *bool                                                                  `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	OutSkuId            *string                                                                `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	MaxStock            *int64                                                                 `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	ProductName         *string                                                                `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetAutoComplement(v bool) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetResidueStock(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ResidueStock = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutId(v string) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetPackFee(v *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.PackFee = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetSkuId(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetSellOut(v bool) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutSkuId(v string) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetMaxStock(v int64) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.MaxStock = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductName(v string) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

type DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee struct {
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetStep(v int32) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.Step = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFee(v int32) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFeeUnit(v int) *DishDraftGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFeeUnit = &v
	return s
}

type DishDraftGetResponseDataDishsItemApplyDate struct {
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemApplyDate) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemApplyDate) SetUseDateType(v int) *DishDraftGetResponseDataDishsItemApplyDate {
	s.UseDateType = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemApplyDate) SetUseEndDate(v string) *DishDraftGetResponseDataDishsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemApplyDate) SetUseStartDate(v string) *DishDraftGetResponseDataDishsItemApplyDate {
	s.UseStartDate = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemApplyDate) SetDayDuration(v int32) *DishDraftGetResponseDataDishsItemApplyDate {
	s.DayDuration = &v
	return s
}

type DishDraftGetResponseDataDishsItemAttributesItem struct {
	ItemList  []*DishDraftGetResponseDataDishsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                        `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemAttributesItem) SetItemList(v []*DishDraftGetResponseDataDishsItemAttributesItemItemListItem) *DishDraftGetResponseDataDishsItemAttributesItem {
	s.ItemList = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemAttributesItem) SetGroupName(v string) *DishDraftGetResponseDataDishsItemAttributesItem {
	s.GroupName = &v
	return s
}

type DishDraftGetResponseDataDishsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemAttributesItemItemListItem) SetName(v string) *DishDraftGetResponseDataDishsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type DishDraftGetResponseDataDishsItemDishDetailInfo struct {
	MaterialInfo []*DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem `json:"material_info,omitempty" xml:"material_info,omitempty" type:"Repeated"`
}

func (s DishDraftGetResponseDataDishsItemDishDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemDishDetailInfo) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemDishDetailInfo) SetMaterialInfo(v []*DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) *DishDraftGetResponseDataDishsItemDishDetailInfo {
	s.MaterialInfo = v
	return s
}

type DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem struct {
	Name  *string   `json:"name,omitempty" xml:"name,omitempty"`
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Key   *string   `json:"key,omitempty" xml:"key,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetName(v string) *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Name = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetValue(v []*string) *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Value = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetKey(v string) *DishDraftGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Key = &v
	return s
}

type DishDraftGetResponseDataDishsItemDishGroupsItem struct {
	DishGroupId       *int64  `json:"dish_group_id,omitempty" xml:"dish_group_id,omitempty"`
	DishGroupName     *string `json:"dish_group_name,omitempty" xml:"dish_group_name,omitempty"`
	GroupRankingScore *string `json:"group_ranking_score,omitempty" xml:"group_ranking_score,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemDishGroupsItem) SetDishGroupId(v int64) *DishDraftGetResponseDataDishsItemDishGroupsItem {
	s.DishGroupId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemDishGroupsItem) SetDishGroupName(v string) *DishDraftGetResponseDataDishsItemDishGroupsItem {
	s.DishGroupName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemDishGroupsItem) SetGroupRankingScore(v string) *DishDraftGetResponseDataDishsItemDishGroupsItem {
	s.GroupRankingScore = &v
	return s
}

type DishDraftGetResponseDataDishsItemImageListItem struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s DishDraftGetResponseDataDishsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemImageListItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemImageListItem) SetUrl(v string) *DishDraftGetResponseDataDishsItemImageListItem {
	s.Url = &v
	return s
}

type DishDraftGetResponseDataDishsItemPoisItem struct {
	PoiId *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemPoisItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemPoisItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemPoisItem) SetPoiId(v string) *DishDraftGetResponseDataDishsItemPoisItem {
	s.PoiId = &v
	return s
}

type DishDraftGetResponseDataDishsItemProductSpecAttrsItem struct {
	ItemList  []*DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                              `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItem) SetItemList(v []*DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) *DishDraftGetResponseDataDishsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItem) SetGroupName(v string) *DishDraftGetResponseDataDishsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

type DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem struct {
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *DishDraftGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

type DishDraftGetResponseDataDishsItemSkusItem struct {
	IsSetAutoComplement *bool                                                        `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	OutSkuId            *string                                                      `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	MaxStock            *int64                                                       `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	ResidueStock        *int64                                                       `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	IsSetSellOut        *bool                                                        `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	SkuSpecAttrs        []*DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	SkuId               *int64                                                       `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName             *string                                                      `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	PackFee             *DishDraftGetResponseDataDishsItemSkusItemPackFee            `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	ActualAmount        *int64                                                       `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemSkusItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetIsSetAutoComplement(v bool) *DishDraftGetResponseDataDishsItemSkusItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetOutSkuId(v string) *DishDraftGetResponseDataDishsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetMaxStock(v int64) *DishDraftGetResponseDataDishsItemSkusItem {
	s.MaxStock = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetResidueStock(v int64) *DishDraftGetResponseDataDishsItemSkusItem {
	s.ResidueStock = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetIsSetSellOut(v bool) *DishDraftGetResponseDataDishsItemSkusItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetSkuSpecAttrs(v []*DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) *DishDraftGetResponseDataDishsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetSkuId(v int64) *DishDraftGetResponseDataDishsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetSkuName(v string) *DishDraftGetResponseDataDishsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetPackFee(v *DishDraftGetResponseDataDishsItemSkusItemPackFee) *DishDraftGetResponseDataDishsItemSkusItem {
	s.PackFee = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItem) SetActualAmount(v int64) *DishDraftGetResponseDataDishsItemSkusItem {
	s.ActualAmount = &v
	return s
}

type DishDraftGetResponseDataDishsItemSkusItemPackFee struct {
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemSkusItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemSkusItemPackFee) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemSkusItemPackFee) SetPackFeeUnit(v int) *DishDraftGetResponseDataDishsItemSkusItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemPackFee) SetStep(v int32) *DishDraftGetResponseDataDishsItemSkusItemPackFee {
	s.Step = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemPackFee) SetPackFee(v int32) *DishDraftGetResponseDataDishsItemSkusItemPackFee {
	s.PackFee = &v
	return s
}

type DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem struct {
	ItemList  []*DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                  `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

type DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *DishDraftGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

type DishDraftGetResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishDraftGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishDraftGetResponseExtra) GoString() string {
	return s.String()
}

func (s *DishDraftGetResponseExtra) SetLogid(v string) *DishDraftGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishDraftGetResponseExtra) SetNow(v int64) *DishDraftGetResponseExtra {
	s.Now = &v
	return s
}

func (s *DishDraftGetResponseExtra) SetSubDescription(v string) *DishDraftGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishDraftGetResponseExtra) SetSubErrorCode(v int32) *DishDraftGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishDraftGetResponseExtra) SetDescription(v string) *DishDraftGetResponseExtra {
	s.Description = &v
	return s
}

func (s *DishDraftGetResponseExtra) SetErrorCode(v int32) *DishDraftGetResponseExtra {
	s.ErrorCode = &v
	return s
}

type DishGroupQueryRequest struct {
	PageNo      *int32                     `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32                     `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PoiId       *int64                     `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Base        *DishGroupQueryRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	GroupType   *int                       `json:"group_type,omitempty" xml:"group_type,omitempty" require:"true"`
}

func (s DishGroupQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryRequest) GoString() string {
	return s.String()
}

func (s *DishGroupQueryRequest) SetPageNo(v int32) *DishGroupQueryRequest {
	s.PageNo = &v
	return s
}

func (s *DishGroupQueryRequest) SetPageSize(v int32) *DishGroupQueryRequest {
	s.PageSize = &v
	return s
}

func (s *DishGroupQueryRequest) SetPoiId(v int64) *DishGroupQueryRequest {
	s.PoiId = &v
	return s
}

func (s *DishGroupQueryRequest) SetBase(v *DishGroupQueryRequestBase) *DishGroupQueryRequest {
	s.Base = v
	return s
}

func (s *DishGroupQueryRequest) SetAccountId(v string) *DishGroupQueryRequest {
	s.AccountId = &v
	return s
}

func (s *DishGroupQueryRequest) SetHeader(v map[string]*string) *DishGroupQueryRequest {
	s.Header = v
	return s
}

func (s *DishGroupQueryRequest) SetAccessToken(v string) *DishGroupQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *DishGroupQueryRequest) SetGroupType(v int) *DishGroupQueryRequest {
	s.GroupType = &v
	return s
}

type DishGroupQueryRequestBase struct {
	TrafficEnv *DishGroupQueryRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                              `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                              `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                              `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                   `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                              `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s DishGroupQueryRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryRequestBase) GoString() string {
	return s.String()
}

func (s *DishGroupQueryRequestBase) SetTrafficEnv(v *DishGroupQueryRequestBaseTrafficEnv) *DishGroupQueryRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishGroupQueryRequestBase) SetAddr(v string) *DishGroupQueryRequestBase {
	s.Addr = &v
	return s
}

func (s *DishGroupQueryRequestBase) SetCaller(v string) *DishGroupQueryRequestBase {
	s.Caller = &v
	return s
}

func (s *DishGroupQueryRequestBase) SetClient(v string) *DishGroupQueryRequestBase {
	s.Client = &v
	return s
}

func (s *DishGroupQueryRequestBase) SetExtra(v map[string]*string) *DishGroupQueryRequestBase {
	s.Extra = v
	return s
}

func (s *DishGroupQueryRequestBase) SetLogID(v string) *DishGroupQueryRequestBase {
	s.LogID = &v
	return s
}

type DishGroupQueryRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishGroupQueryRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishGroupQueryRequestBaseTrafficEnv) SetEnv(v string) *DishGroupQueryRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishGroupQueryRequestBaseTrafficEnv) SetOpen(v bool) *DishGroupQueryRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishGroupQueryResponse struct {
	Data     *DishGroupQueryResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishGroupQueryResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishGroupQueryResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s DishGroupQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryResponse) GoString() string {
	return s.String()
}

func (s *DishGroupQueryResponse) SetData(v *DishGroupQueryResponseData) *DishGroupQueryResponse {
	s.Data = v
	return s
}

func (s *DishGroupQueryResponse) SetExtra(v *DishGroupQueryResponseExtra) *DishGroupQueryResponse {
	s.Extra = v
	return s
}

func (s *DishGroupQueryResponse) SetBaseResp(v *DishGroupQueryResponseBaseResp) *DishGroupQueryResponse {
	s.BaseResp = v
	return s
}

type DishGroupQueryResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishGroupQueryResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishGroupQueryResponseBaseResp) SetExtra(v map[string]*string) *DishGroupQueryResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishGroupQueryResponseBaseResp) SetStatusCode(v int32) *DishGroupQueryResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishGroupQueryResponseBaseResp) SetStatusMessage(v string) *DishGroupQueryResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishGroupQueryResponseData struct {
	ErrorCode   *int32                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore     *bool                                       `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Total       *int32                                      `json:"total,omitempty" xml:"total,omitempty"`
	Description *string                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	DishGroups  []*DishGroupQueryResponseDataDishGroupsItem `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
}

func (s DishGroupQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryResponseData) GoString() string {
	return s.String()
}

func (s *DishGroupQueryResponseData) SetErrorCode(v int32) *DishGroupQueryResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishGroupQueryResponseData) SetHasMore(v bool) *DishGroupQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *DishGroupQueryResponseData) SetTotal(v int32) *DishGroupQueryResponseData {
	s.Total = &v
	return s
}

func (s *DishGroupQueryResponseData) SetDescription(v string) *DishGroupQueryResponseData {
	s.Description = &v
	return s
}

func (s *DishGroupQueryResponseData) SetDishGroups(v []*DishGroupQueryResponseDataDishGroupsItem) *DishGroupQueryResponseData {
	s.DishGroups = v
	return s
}

type DishGroupQueryResponseDataDishGroupsItem struct {
	GroupName    *string  `json:"group_name,omitempty" xml:"group_name,omitempty"`
	IsMust       *bool    `json:"is_must,omitempty" xml:"is_must,omitempty"`
	RankingScore *float64 `json:"ranking_score,omitempty" xml:"ranking_score,omitempty"`
	GroupDesc    *string  `json:"group_desc,omitempty" xml:"group_desc,omitempty"`
	GroupId      *int64   `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DishGroupQueryResponseDataDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryResponseDataDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishGroupQueryResponseDataDishGroupsItem) SetGroupName(v string) *DishGroupQueryResponseDataDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *DishGroupQueryResponseDataDishGroupsItem) SetIsMust(v bool) *DishGroupQueryResponseDataDishGroupsItem {
	s.IsMust = &v
	return s
}

func (s *DishGroupQueryResponseDataDishGroupsItem) SetRankingScore(v float64) *DishGroupQueryResponseDataDishGroupsItem {
	s.RankingScore = &v
	return s
}

func (s *DishGroupQueryResponseDataDishGroupsItem) SetGroupDesc(v string) *DishGroupQueryResponseDataDishGroupsItem {
	s.GroupDesc = &v
	return s
}

func (s *DishGroupQueryResponseDataDishGroupsItem) SetGroupId(v int64) *DishGroupQueryResponseDataDishGroupsItem {
	s.GroupId = &v
	return s
}

type DishGroupQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s DishGroupQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishGroupQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *DishGroupQueryResponseExtra) SetSubErrorCode(v int32) *DishGroupQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishGroupQueryResponseExtra) SetDescription(v string) *DishGroupQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *DishGroupQueryResponseExtra) SetErrorCode(v int32) *DishGroupQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishGroupQueryResponseExtra) SetLogid(v string) *DishGroupQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishGroupQueryResponseExtra) SetNow(v int64) *DishGroupQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *DishGroupQueryResponseExtra) SetSubDescription(v string) *DishGroupQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type DishGroupSaveRequest struct {
	GroupInfo   *DishGroupSaveRequestGroupInfo `json:"group_info,omitempty" xml:"group_info,omitempty"`
	GroupType   *int                           `json:"group_type,omitempty" xml:"group_type,omitempty" require:"true"`
	Header      map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiId       *int64                         `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Base        *DishGroupSaveRequestBase      `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s DishGroupSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveRequest) GoString() string {
	return s.String()
}

func (s *DishGroupSaveRequest) SetGroupInfo(v *DishGroupSaveRequestGroupInfo) *DishGroupSaveRequest {
	s.GroupInfo = v
	return s
}

func (s *DishGroupSaveRequest) SetGroupType(v int) *DishGroupSaveRequest {
	s.GroupType = &v
	return s
}

func (s *DishGroupSaveRequest) SetHeader(v map[string]*string) *DishGroupSaveRequest {
	s.Header = v
	return s
}

func (s *DishGroupSaveRequest) SetAccessToken(v string) *DishGroupSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *DishGroupSaveRequest) SetPoiId(v int64) *DishGroupSaveRequest {
	s.PoiId = &v
	return s
}

func (s *DishGroupSaveRequest) SetBase(v *DishGroupSaveRequestBase) *DishGroupSaveRequest {
	s.Base = v
	return s
}

func (s *DishGroupSaveRequest) SetAccountId(v string) *DishGroupSaveRequest {
	s.AccountId = &v
	return s
}

type DishGroupSaveRequestBase struct {
	Addr       *string                             `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                             `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                             `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                  `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                             `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishGroupSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s DishGroupSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveRequestBase) GoString() string {
	return s.String()
}

func (s *DishGroupSaveRequestBase) SetAddr(v string) *DishGroupSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *DishGroupSaveRequestBase) SetCaller(v string) *DishGroupSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *DishGroupSaveRequestBase) SetClient(v string) *DishGroupSaveRequestBase {
	s.Client = &v
	return s
}

func (s *DishGroupSaveRequestBase) SetExtra(v map[string]*string) *DishGroupSaveRequestBase {
	s.Extra = v
	return s
}

func (s *DishGroupSaveRequestBase) SetLogID(v string) *DishGroupSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *DishGroupSaveRequestBase) SetTrafficEnv(v *DishGroupSaveRequestBaseTrafficEnv) *DishGroupSaveRequestBase {
	s.TrafficEnv = v
	return s
}

type DishGroupSaveRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s DishGroupSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishGroupSaveRequestBaseTrafficEnv) SetOpen(v bool) *DishGroupSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *DishGroupSaveRequestBaseTrafficEnv) SetEnv(v string) *DishGroupSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type DishGroupSaveRequestGroupInfo struct {
	IsMust    *bool   `json:"is_must,omitempty" xml:"is_must,omitempty"`
	GroupDesc *string `json:"group_desc,omitempty" xml:"group_desc,omitempty"`
	GroupId   *int64  `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishGroupSaveRequestGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *DishGroupSaveRequestGroupInfo) SetIsMust(v bool) *DishGroupSaveRequestGroupInfo {
	s.IsMust = &v
	return s
}

func (s *DishGroupSaveRequestGroupInfo) SetGroupDesc(v string) *DishGroupSaveRequestGroupInfo {
	s.GroupDesc = &v
	return s
}

func (s *DishGroupSaveRequestGroupInfo) SetGroupId(v int64) *DishGroupSaveRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *DishGroupSaveRequestGroupInfo) SetGroupName(v string) *DishGroupSaveRequestGroupInfo {
	s.GroupName = &v
	return s
}

type DishGroupSaveResponse struct {
	BaseResp *DishGroupSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishGroupSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishGroupSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishGroupSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveResponse) GoString() string {
	return s.String()
}

func (s *DishGroupSaveResponse) SetBaseResp(v *DishGroupSaveResponseBaseResp) *DishGroupSaveResponse {
	s.BaseResp = v
	return s
}

func (s *DishGroupSaveResponse) SetData(v *DishGroupSaveResponseData) *DishGroupSaveResponse {
	s.Data = v
	return s
}

func (s *DishGroupSaveResponse) SetExtra(v *DishGroupSaveResponseExtra) *DishGroupSaveResponse {
	s.Extra = v
	return s
}

type DishGroupSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishGroupSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishGroupSaveResponseBaseResp) SetStatusCode(v int32) *DishGroupSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishGroupSaveResponseBaseResp) SetStatusMessage(v string) *DishGroupSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishGroupSaveResponseBaseResp) SetExtra(v map[string]*string) *DishGroupSaveResponseBaseResp {
	s.Extra = v
	return s
}

type DishGroupSaveResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishGroupSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveResponseData) GoString() string {
	return s.String()
}

func (s *DishGroupSaveResponseData) SetDescription(v string) *DishGroupSaveResponseData {
	s.Description = &v
	return s
}

func (s *DishGroupSaveResponseData) SetErrorCode(v int32) *DishGroupSaveResponseData {
	s.ErrorCode = &v
	return s
}

type DishGroupSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s DishGroupSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *DishGroupSaveResponseExtra) SetDescription(v string) *DishGroupSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *DishGroupSaveResponseExtra) SetErrorCode(v int32) *DishGroupSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishGroupSaveResponseExtra) SetLogid(v string) *DishGroupSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishGroupSaveResponseExtra) SetNow(v int64) *DishGroupSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *DishGroupSaveResponseExtra) SetSubDescription(v string) *DishGroupSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishGroupSaveResponseExtra) SetSubErrorCode(v int32) *DishGroupSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type DishGroupSortRequest struct {
	AccountId   *string                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GroupType   *int                      `json:"group_type,omitempty" xml:"group_type,omitempty" require:"true"`
	PoiId       *int64                    `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SortList    []*int64                  `json:"sort_list,omitempty" xml:"sort_list,omitempty" require:"true" type:"Repeated"`
	Base        *DishGroupSortRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DishGroupSortRequest) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortRequest) GoString() string {
	return s.String()
}

func (s *DishGroupSortRequest) SetAccountId(v string) *DishGroupSortRequest {
	s.AccountId = &v
	return s
}

func (s *DishGroupSortRequest) SetGroupType(v int) *DishGroupSortRequest {
	s.GroupType = &v
	return s
}

func (s *DishGroupSortRequest) SetPoiId(v int64) *DishGroupSortRequest {
	s.PoiId = &v
	return s
}

func (s *DishGroupSortRequest) SetSortList(v []*int64) *DishGroupSortRequest {
	s.SortList = v
	return s
}

func (s *DishGroupSortRequest) SetBase(v *DishGroupSortRequestBase) *DishGroupSortRequest {
	s.Base = v
	return s
}

func (s *DishGroupSortRequest) SetHeader(v map[string]*string) *DishGroupSortRequest {
	s.Header = v
	return s
}

func (s *DishGroupSortRequest) SetAccessToken(v string) *DishGroupSortRequest {
	s.AccessToken = &v
	return s
}

type DishGroupSortRequestBase struct {
	Client     *string                             `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                  `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                             `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishGroupSortRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                             `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                             `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s DishGroupSortRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortRequestBase) GoString() string {
	return s.String()
}

func (s *DishGroupSortRequestBase) SetClient(v string) *DishGroupSortRequestBase {
	s.Client = &v
	return s
}

func (s *DishGroupSortRequestBase) SetExtra(v map[string]*string) *DishGroupSortRequestBase {
	s.Extra = v
	return s
}

func (s *DishGroupSortRequestBase) SetLogID(v string) *DishGroupSortRequestBase {
	s.LogID = &v
	return s
}

func (s *DishGroupSortRequestBase) SetTrafficEnv(v *DishGroupSortRequestBaseTrafficEnv) *DishGroupSortRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishGroupSortRequestBase) SetAddr(v string) *DishGroupSortRequestBase {
	s.Addr = &v
	return s
}

func (s *DishGroupSortRequestBase) SetCaller(v string) *DishGroupSortRequestBase {
	s.Caller = &v
	return s
}

type DishGroupSortRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishGroupSortRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishGroupSortRequestBaseTrafficEnv) SetEnv(v string) *DishGroupSortRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishGroupSortRequestBaseTrafficEnv) SetOpen(v bool) *DishGroupSortRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishGroupSortResponse struct {
	Extra    *DishGroupSortResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishGroupSortResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishGroupSortResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DishGroupSortResponse) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortResponse) GoString() string {
	return s.String()
}

func (s *DishGroupSortResponse) SetExtra(v *DishGroupSortResponseExtra) *DishGroupSortResponse {
	s.Extra = v
	return s
}

func (s *DishGroupSortResponse) SetBaseResp(v *DishGroupSortResponseBaseResp) *DishGroupSortResponse {
	s.BaseResp = v
	return s
}

func (s *DishGroupSortResponse) SetData(v *DishGroupSortResponseData) *DishGroupSortResponse {
	s.Data = v
	return s
}

type DishGroupSortResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishGroupSortResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishGroupSortResponseBaseResp) SetExtra(v map[string]*string) *DishGroupSortResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishGroupSortResponseBaseResp) SetStatusCode(v int32) *DishGroupSortResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishGroupSortResponseBaseResp) SetStatusMessage(v string) *DishGroupSortResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishGroupSortResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DishGroupSortResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortResponseData) GoString() string {
	return s.String()
}

func (s *DishGroupSortResponseData) SetErrorCode(v int32) *DishGroupSortResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishGroupSortResponseData) SetDescription(v string) *DishGroupSortResponseData {
	s.Description = &v
	return s
}

type DishGroupSortResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s DishGroupSortResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishGroupSortResponseExtra) GoString() string {
	return s.String()
}

func (s *DishGroupSortResponseExtra) SetNow(v int64) *DishGroupSortResponseExtra {
	s.Now = &v
	return s
}

func (s *DishGroupSortResponseExtra) SetSubDescription(v string) *DishGroupSortResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishGroupSortResponseExtra) SetSubErrorCode(v int32) *DishGroupSortResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishGroupSortResponseExtra) SetDescription(v string) *DishGroupSortResponseExtra {
	s.Description = &v
	return s
}

func (s *DishGroupSortResponseExtra) SetErrorCode(v int32) *DishGroupSortResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishGroupSortResponseExtra) SetLogid(v string) *DishGroupSortResponseExtra {
	s.Logid = &v
	return s
}

type DishOnlineGetRequest struct {
	PoiId       *int64                    `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ProductIds  []*int64                  `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base        *DishOnlineGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DishType    *int                      `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
}

func (s DishOnlineGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetRequest) GoString() string {
	return s.String()
}

func (s *DishOnlineGetRequest) SetPoiId(v int64) *DishOnlineGetRequest {
	s.PoiId = &v
	return s
}

func (s *DishOnlineGetRequest) SetProductIds(v []*int64) *DishOnlineGetRequest {
	s.ProductIds = v
	return s
}

func (s *DishOnlineGetRequest) SetBase(v *DishOnlineGetRequestBase) *DishOnlineGetRequest {
	s.Base = v
	return s
}

func (s *DishOnlineGetRequest) SetHeader(v map[string]*string) *DishOnlineGetRequest {
	s.Header = v
	return s
}

func (s *DishOnlineGetRequest) SetAccessToken(v string) *DishOnlineGetRequest {
	s.AccessToken = &v
	return s
}

func (s *DishOnlineGetRequest) SetAccountId(v string) *DishOnlineGetRequest {
	s.AccountId = &v
	return s
}

func (s *DishOnlineGetRequest) SetDishType(v int) *DishOnlineGetRequest {
	s.DishType = &v
	return s
}

type DishOnlineGetRequestBase struct {
	Caller     *string                             `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                             `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                  `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                             `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishOnlineGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                             `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s DishOnlineGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetRequestBase) GoString() string {
	return s.String()
}

func (s *DishOnlineGetRequestBase) SetCaller(v string) *DishOnlineGetRequestBase {
	s.Caller = &v
	return s
}

func (s *DishOnlineGetRequestBase) SetClient(v string) *DishOnlineGetRequestBase {
	s.Client = &v
	return s
}

func (s *DishOnlineGetRequestBase) SetExtra(v map[string]*string) *DishOnlineGetRequestBase {
	s.Extra = v
	return s
}

func (s *DishOnlineGetRequestBase) SetLogID(v string) *DishOnlineGetRequestBase {
	s.LogID = &v
	return s
}

func (s *DishOnlineGetRequestBase) SetTrafficEnv(v *DishOnlineGetRequestBaseTrafficEnv) *DishOnlineGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishOnlineGetRequestBase) SetAddr(v string) *DishOnlineGetRequestBase {
	s.Addr = &v
	return s
}

type DishOnlineGetRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s DishOnlineGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishOnlineGetRequestBaseTrafficEnv) SetOpen(v bool) *DishOnlineGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *DishOnlineGetRequestBaseTrafficEnv) SetEnv(v string) *DishOnlineGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type DishOnlineGetResponse struct {
	BaseResp *DishOnlineGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishOnlineGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishOnlineGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishOnlineGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponse) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponse) SetBaseResp(v *DishOnlineGetResponseBaseResp) *DishOnlineGetResponse {
	s.BaseResp = v
	return s
}

func (s *DishOnlineGetResponse) SetData(v *DishOnlineGetResponseData) *DishOnlineGetResponse {
	s.Data = v
	return s
}

func (s *DishOnlineGetResponse) SetExtra(v *DishOnlineGetResponseExtra) *DishOnlineGetResponse {
	s.Extra = v
	return s
}

type DishOnlineGetResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DishOnlineGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseBaseResp) SetStatusMessage(v string) *DishOnlineGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishOnlineGetResponseBaseResp) SetExtra(v map[string]*string) *DishOnlineGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishOnlineGetResponseBaseResp) SetStatusCode(v int32) *DishOnlineGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

type DishOnlineGetResponseData struct {
	Dishs         []*DishOnlineGetResponseDataDishsItem `json:"dishs,omitempty" xml:"dishs,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DishOnlineGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseData) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseData) SetDishs(v []*DishOnlineGetResponseDataDishsItem) *DishOnlineGetResponseData {
	s.Dishs = v
	return s
}

func (s *DishOnlineGetResponseData) SetGwErrorCode(v int32) *DishOnlineGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DishOnlineGetResponseData) SetGwDescription(v string) *DishOnlineGetResponseData {
	s.GwDescription = &v
	return s
}

type DishOnlineGetResponseDataDishsItem struct {
	AccountName       *string                                                   `json:"account_name,omitempty" xml:"account_name,omitempty"`
	ProductName       *string                                                   `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductSpecAttrs  []*DishOnlineGetResponseDataDishsItemProductSpecAttrsItem `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	MerchantProductId *string                                                   `json:"merchant_product_id,omitempty" xml:"merchant_product_id,omitempty"`
	AccountId         *string                                                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutId             *string                                                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	IsBindMerchant    *bool                                                     `json:"is_bind_merchant,omitempty" xml:"is_bind_merchant,omitempty"`
	OnlineStatus      *int                                                      `json:"online_status,omitempty" xml:"online_status,omitempty"`
	DraftStatus       *int                                                      `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	CategoryId        *int64                                                    `json:"category_id,omitempty" xml:"category_id,omitempty"`
	Skus              []*DishOnlineGetResponseDataDishsItemSkusItem             `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	DeliveryMethod    []*int                                                    `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	DishDetailInfo    *DishOnlineGetResponseDataDishsItemDishDetailInfo         `json:"dish_detail_info,omitempty" xml:"dish_detail_info,omitempty"`
	ImageList         []*DishOnlineGetResponseDataDishsItemImageListItem        `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	SettleType        *int64                                                    `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	Attributes        []*DishOnlineGetResponseDataDishsItemAttributesItem       `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	AddDishGroups     []*DishOnlineGetResponseDataDishsItemAddDishGroupsItem    `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	DishDescription   *string                                                   `json:"dish_description,omitempty" xml:"dish_description,omitempty"`
	CreateTime        *int64                                                    `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Pois              []*DishOnlineGetResponseDataDishsItemPoisItem             `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	DishType          *int                                                      `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	ProductId         *int64                                                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ApplyDate         *DishOnlineGetResponseDataDishsItemApplyDate              `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	UpdateTime        *int64                                                    `json:"update_time,omitempty" xml:"update_time,omitempty"`
	DishGroups        []*DishOnlineGetResponseDataDishsItemDishGroupsItem       `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
	PoiCount          *int64                                                    `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItem) SetAccountName(v string) *DishOnlineGetResponseDataDishsItem {
	s.AccountName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetProductName(v string) *DishOnlineGetResponseDataDishsItem {
	s.ProductName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetProductSpecAttrs(v []*DishOnlineGetResponseDataDishsItemProductSpecAttrsItem) *DishOnlineGetResponseDataDishsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetMerchantProductId(v string) *DishOnlineGetResponseDataDishsItem {
	s.MerchantProductId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetAccountId(v string) *DishOnlineGetResponseDataDishsItem {
	s.AccountId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetOutId(v string) *DishOnlineGetResponseDataDishsItem {
	s.OutId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetIsBindMerchant(v bool) *DishOnlineGetResponseDataDishsItem {
	s.IsBindMerchant = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetOnlineStatus(v int) *DishOnlineGetResponseDataDishsItem {
	s.OnlineStatus = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDraftStatus(v int) *DishOnlineGetResponseDataDishsItem {
	s.DraftStatus = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetCategoryId(v int64) *DishOnlineGetResponseDataDishsItem {
	s.CategoryId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetSkus(v []*DishOnlineGetResponseDataDishsItemSkusItem) *DishOnlineGetResponseDataDishsItem {
	s.Skus = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDeliveryMethod(v []*int) *DishOnlineGetResponseDataDishsItem {
	s.DeliveryMethod = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDishDetailInfo(v *DishOnlineGetResponseDataDishsItemDishDetailInfo) *DishOnlineGetResponseDataDishsItem {
	s.DishDetailInfo = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetImageList(v []*DishOnlineGetResponseDataDishsItemImageListItem) *DishOnlineGetResponseDataDishsItem {
	s.ImageList = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetSettleType(v int64) *DishOnlineGetResponseDataDishsItem {
	s.SettleType = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetAttributes(v []*DishOnlineGetResponseDataDishsItemAttributesItem) *DishOnlineGetResponseDataDishsItem {
	s.Attributes = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetAddDishGroups(v []*DishOnlineGetResponseDataDishsItemAddDishGroupsItem) *DishOnlineGetResponseDataDishsItem {
	s.AddDishGroups = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDishDescription(v string) *DishOnlineGetResponseDataDishsItem {
	s.DishDescription = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetCreateTime(v int64) *DishOnlineGetResponseDataDishsItem {
	s.CreateTime = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetPois(v []*DishOnlineGetResponseDataDishsItemPoisItem) *DishOnlineGetResponseDataDishsItem {
	s.Pois = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDishType(v int) *DishOnlineGetResponseDataDishsItem {
	s.DishType = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetProductId(v int64) *DishOnlineGetResponseDataDishsItem {
	s.ProductId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetApplyDate(v *DishOnlineGetResponseDataDishsItemApplyDate) *DishOnlineGetResponseDataDishsItem {
	s.ApplyDate = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetUpdateTime(v int64) *DishOnlineGetResponseDataDishsItem {
	s.UpdateTime = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetDishGroups(v []*DishOnlineGetResponseDataDishsItemDishGroupsItem) *DishOnlineGetResponseDataDishsItem {
	s.DishGroups = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItem) SetPoiCount(v int64) *DishOnlineGetResponseDataDishsItem {
	s.PoiCount = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAddDishGroupsItem struct {
	GroupName *string                                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupId   *int64                                                             `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItem) SetGroupName(v string) *DishOnlineGetResponseDataDishsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItem) SetItemList(v []*DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) *DishOnlineGetResponseDataDishsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItem) SetGroupId(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem struct {
	Stock               *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock   `json:"stock,omitempty" xml:"stock,omitempty"`
	ResidueStock        *int64                                                                  `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	IsSetSellOut        *bool                                                                   `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	OutSkuId            *string                                                                 `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	IsSetAutoComplement *bool                                                                   `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	PackFee             *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	OutId               *string                                                                 `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductName         *string                                                                 `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductId           *int64                                                                  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	MaxStock            *int64                                                                  `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	SkuId               *int64                                                                  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Price               *int64                                                                  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetStock(v *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Stock = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetResidueStock(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ResidueStock = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetSellOut(v bool) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutSkuId(v string) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetAutoComplement(v bool) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetPackFee(v *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.PackFee = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutId(v string) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductName(v string) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetMaxStock(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.MaxStock = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetSkuId(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee struct {
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFee(v int32) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFeeUnit(v int) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetStep(v int32) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.Step = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock struct {
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetFrozenQty(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.FrozenQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetLimitType(v int) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.LimitType = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetSoldCount(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.SoldCount = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetSoldQty(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.SoldQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetStockQty(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.StockQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetAvailQty(v int64) *DishOnlineGetResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.AvailQty = &v
	return s
}

type DishOnlineGetResponseDataDishsItemApplyDate struct {
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
}

func (s DishOnlineGetResponseDataDishsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemApplyDate) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemApplyDate) SetUseEndDate(v string) *DishOnlineGetResponseDataDishsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemApplyDate) SetUseStartDate(v string) *DishOnlineGetResponseDataDishsItemApplyDate {
	s.UseStartDate = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemApplyDate) SetDayDuration(v int32) *DishOnlineGetResponseDataDishsItemApplyDate {
	s.DayDuration = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemApplyDate) SetUseDateType(v int) *DishOnlineGetResponseDataDishsItemApplyDate {
	s.UseDateType = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAttributesItem struct {
	ItemList  []*DishOnlineGetResponseDataDishsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                         `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAttributesItem) SetItemList(v []*DishOnlineGetResponseDataDishsItemAttributesItemItemListItem) *DishOnlineGetResponseDataDishsItemAttributesItem {
	s.ItemList = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemAttributesItem) SetGroupName(v string) *DishOnlineGetResponseDataDishsItemAttributesItem {
	s.GroupName = &v
	return s
}

type DishOnlineGetResponseDataDishsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemAttributesItemItemListItem) SetName(v string) *DishOnlineGetResponseDataDishsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type DishOnlineGetResponseDataDishsItemDishDetailInfo struct {
	MaterialInfo []*DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem `json:"material_info,omitempty" xml:"material_info,omitempty" type:"Repeated"`
}

func (s DishOnlineGetResponseDataDishsItemDishDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemDishDetailInfo) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemDishDetailInfo) SetMaterialInfo(v []*DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) *DishOnlineGetResponseDataDishsItemDishDetailInfo {
	s.MaterialInfo = v
	return s
}

type DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem struct {
	Name  *string   `json:"name,omitempty" xml:"name,omitempty"`
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Key   *string   `json:"key,omitempty" xml:"key,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetName(v string) *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Name = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetValue(v []*string) *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Value = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetKey(v string) *DishOnlineGetResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Key = &v
	return s
}

type DishOnlineGetResponseDataDishsItemDishGroupsItem struct {
	GroupRankingScore *string `json:"group_ranking_score,omitempty" xml:"group_ranking_score,omitempty"`
	DishGroupId       *int64  `json:"dish_group_id,omitempty" xml:"dish_group_id,omitempty"`
	DishGroupName     *string `json:"dish_group_name,omitempty" xml:"dish_group_name,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemDishGroupsItem) SetGroupRankingScore(v string) *DishOnlineGetResponseDataDishsItemDishGroupsItem {
	s.GroupRankingScore = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemDishGroupsItem) SetDishGroupId(v int64) *DishOnlineGetResponseDataDishsItemDishGroupsItem {
	s.DishGroupId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemDishGroupsItem) SetDishGroupName(v string) *DishOnlineGetResponseDataDishsItemDishGroupsItem {
	s.DishGroupName = &v
	return s
}

type DishOnlineGetResponseDataDishsItemImageListItem struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s DishOnlineGetResponseDataDishsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemImageListItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemImageListItem) SetUrl(v string) *DishOnlineGetResponseDataDishsItemImageListItem {
	s.Url = &v
	return s
}

type DishOnlineGetResponseDataDishsItemPoisItem struct {
	PoiId *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemPoisItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemPoisItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemPoisItem) SetPoiId(v string) *DishOnlineGetResponseDataDishsItemPoisItem {
	s.PoiId = &v
	return s
}

type DishOnlineGetResponseDataDishsItemProductSpecAttrsItem struct {
	GroupName *string                                                               `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DishOnlineGetResponseDataDishsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItem) SetGroupName(v string) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItem) SetItemList(v []*DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

type DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem struct {
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *DishOnlineGetResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

type DishOnlineGetResponseDataDishsItemSkusItem struct {
	Stock               *DishOnlineGetResponseDataDishsItemSkusItemStock              `json:"stock,omitempty" xml:"stock,omitempty"`
	SkuId               *int64                                                        `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Status              *int                                                          `json:"status,omitempty" xml:"status,omitempty"`
	PackFee             *DishOnlineGetResponseDataDishsItemSkusItemPackFee            `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	SkuSpecAttrs        []*DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	ActualAmount        *int64                                                        `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	MaxStock            *int64                                                        `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	ResidueStock        *int64                                                        `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	IsSetAutoComplement *bool                                                         `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	SkuName             *string                                                       `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	IsSetSellOut        *bool                                                         `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	OutSkuId            *string                                                       `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemSkusItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetStock(v *DishOnlineGetResponseDataDishsItemSkusItemStock) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.Stock = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetSkuId(v int64) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetStatus(v int) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.Status = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetPackFee(v *DishOnlineGetResponseDataDishsItemSkusItemPackFee) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.PackFee = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetSkuSpecAttrs(v []*DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetActualAmount(v int64) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetMaxStock(v int64) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.MaxStock = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetResidueStock(v int64) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.ResidueStock = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetIsSetAutoComplement(v bool) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetSkuName(v string) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetIsSetSellOut(v bool) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItem) SetOutSkuId(v string) *DishOnlineGetResponseDataDishsItemSkusItem {
	s.OutSkuId = &v
	return s
}

type DishOnlineGetResponseDataDishsItemSkusItemPackFee struct {
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemSkusItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemSkusItemPackFee) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemPackFee) SetPackFee(v int32) *DishOnlineGetResponseDataDishsItemSkusItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemPackFee) SetPackFeeUnit(v int) *DishOnlineGetResponseDataDishsItemSkusItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemPackFee) SetStep(v int32) *DishOnlineGetResponseDataDishsItemSkusItemPackFee {
	s.Step = &v
	return s
}

type DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem struct {
	ItemList  []*DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                   `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

type DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *DishOnlineGetResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

type DishOnlineGetResponseDataDishsItemSkusItemStock struct {
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
}

func (s DishOnlineGetResponseDataDishsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseDataDishsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetStockQty(v int64) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetAvailQty(v int64) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetFrozenQty(v int64) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetLimitType(v int) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetSoldCount(v int64) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *DishOnlineGetResponseDataDishsItemSkusItemStock) SetSoldQty(v int64) *DishOnlineGetResponseDataDishsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

type DishOnlineGetResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s DishOnlineGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishOnlineGetResponseExtra) GoString() string {
	return s.String()
}

func (s *DishOnlineGetResponseExtra) SetSubErrorCode(v int32) *DishOnlineGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishOnlineGetResponseExtra) SetDescription(v string) *DishOnlineGetResponseExtra {
	s.Description = &v
	return s
}

func (s *DishOnlineGetResponseExtra) SetErrorCode(v int32) *DishOnlineGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishOnlineGetResponseExtra) SetLogid(v string) *DishOnlineGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishOnlineGetResponseExtra) SetNow(v int64) *DishOnlineGetResponseExtra {
	s.Now = &v
	return s
}

func (s *DishOnlineGetResponseExtra) SetSubDescription(v string) *DishOnlineGetResponseExtra {
	s.SubDescription = &v
	return s
}

type DishOperateStatusRequest struct {
	ProductIds      []*int64                      `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base            *DishOperateStatusRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header          map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId       *string                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DyPoiId         *int64                        `json:"dy_poi_id,omitempty" xml:"dy_poi_id,omitempty"`
	MerchantOperate *bool                         `json:"merchant_operate,omitempty" xml:"merchant_operate,omitempty"`
	OperateType     *int                          `json:"operate_type,omitempty" xml:"operate_type,omitempty" require:"true"`
}

func (s DishOperateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusRequest) GoString() string {
	return s.String()
}

func (s *DishOperateStatusRequest) SetProductIds(v []*int64) *DishOperateStatusRequest {
	s.ProductIds = v
	return s
}

func (s *DishOperateStatusRequest) SetBase(v *DishOperateStatusRequestBase) *DishOperateStatusRequest {
	s.Base = v
	return s
}

func (s *DishOperateStatusRequest) SetHeader(v map[string]*string) *DishOperateStatusRequest {
	s.Header = v
	return s
}

func (s *DishOperateStatusRequest) SetAccessToken(v string) *DishOperateStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *DishOperateStatusRequest) SetAccountId(v string) *DishOperateStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DishOperateStatusRequest) SetDyPoiId(v int64) *DishOperateStatusRequest {
	s.DyPoiId = &v
	return s
}

func (s *DishOperateStatusRequest) SetMerchantOperate(v bool) *DishOperateStatusRequest {
	s.MerchantOperate = &v
	return s
}

func (s *DishOperateStatusRequest) SetOperateType(v int) *DishOperateStatusRequest {
	s.OperateType = &v
	return s
}

type DishOperateStatusRequestBase struct {
	Client     *string                                 `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                 `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishOperateStatusRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                 `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                 `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s DishOperateStatusRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusRequestBase) GoString() string {
	return s.String()
}

func (s *DishOperateStatusRequestBase) SetClient(v string) *DishOperateStatusRequestBase {
	s.Client = &v
	return s
}

func (s *DishOperateStatusRequestBase) SetExtra(v map[string]*string) *DishOperateStatusRequestBase {
	s.Extra = v
	return s
}

func (s *DishOperateStatusRequestBase) SetLogID(v string) *DishOperateStatusRequestBase {
	s.LogID = &v
	return s
}

func (s *DishOperateStatusRequestBase) SetTrafficEnv(v *DishOperateStatusRequestBaseTrafficEnv) *DishOperateStatusRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishOperateStatusRequestBase) SetAddr(v string) *DishOperateStatusRequestBase {
	s.Addr = &v
	return s
}

func (s *DishOperateStatusRequestBase) SetCaller(v string) *DishOperateStatusRequestBase {
	s.Caller = &v
	return s
}

type DishOperateStatusRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishOperateStatusRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishOperateStatusRequestBaseTrafficEnv) SetEnv(v string) *DishOperateStatusRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishOperateStatusRequestBaseTrafficEnv) SetOpen(v bool) *DishOperateStatusRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishOperateStatusResponse struct {
	Data     *DishOperateStatusResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishOperateStatusResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishOperateStatusResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s DishOperateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusResponse) GoString() string {
	return s.String()
}

func (s *DishOperateStatusResponse) SetData(v *DishOperateStatusResponseData) *DishOperateStatusResponse {
	s.Data = v
	return s
}

func (s *DishOperateStatusResponse) SetExtra(v *DishOperateStatusResponseExtra) *DishOperateStatusResponse {
	s.Extra = v
	return s
}

func (s *DishOperateStatusResponse) SetBaseResp(v *DishOperateStatusResponseBaseResp) *DishOperateStatusResponse {
	s.BaseResp = v
	return s
}

type DishOperateStatusResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishOperateStatusResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishOperateStatusResponseBaseResp) SetStatusCode(v int32) *DishOperateStatusResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishOperateStatusResponseBaseResp) SetStatusMessage(v string) *DishOperateStatusResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishOperateStatusResponseBaseResp) SetExtra(v map[string]*string) *DishOperateStatusResponseBaseResp {
	s.Extra = v
	return s
}

type DishOperateStatusResponseData struct {
	Description *string           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrMap      map[int64]*string `json:"err_map,omitempty" xml:"err_map,omitempty"`
	ErrorCode   *int32            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishOperateStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusResponseData) GoString() string {
	return s.String()
}

func (s *DishOperateStatusResponseData) SetDescription(v string) *DishOperateStatusResponseData {
	s.Description = &v
	return s
}

func (s *DishOperateStatusResponseData) SetErrMap(v map[int64]*string) *DishOperateStatusResponseData {
	s.ErrMap = v
	return s
}

func (s *DishOperateStatusResponseData) SetErrorCode(v int32) *DishOperateStatusResponseData {
	s.ErrorCode = &v
	return s
}

type DishOperateStatusResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishOperateStatusResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishOperateStatusResponseExtra) GoString() string {
	return s.String()
}

func (s *DishOperateStatusResponseExtra) SetLogid(v string) *DishOperateStatusResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishOperateStatusResponseExtra) SetNow(v int64) *DishOperateStatusResponseExtra {
	s.Now = &v
	return s
}

func (s *DishOperateStatusResponseExtra) SetSubDescription(v string) *DishOperateStatusResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishOperateStatusResponseExtra) SetSubErrorCode(v int32) *DishOperateStatusResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishOperateStatusResponseExtra) SetDescription(v string) *DishOperateStatusResponseExtra {
	s.Description = &v
	return s
}

func (s *DishOperateStatusResponseExtra) SetErrorCode(v int32) *DishOperateStatusResponseExtra {
	s.ErrorCode = &v
	return s
}

type DishSaveOutIdRequest struct {
	FeedOutIdMap map[int64]*string         `json:"feed_out_id_map,omitempty" xml:"feed_out_id_map,omitempty"`
	PoiId        *int64                    `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	AccessToken  *string                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base         *DishSaveOutIdRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId    *string                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutId        *string                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId    *int64                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuOutIdMap  map[int64]*string         `json:"sku_out_id_map,omitempty" xml:"sku_out_id_map,omitempty"`
	Header       map[string]*string        `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DishSaveOutIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdRequest) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdRequest) SetFeedOutIdMap(v map[int64]*string) *DishSaveOutIdRequest {
	s.FeedOutIdMap = v
	return s
}

func (s *DishSaveOutIdRequest) SetPoiId(v int64) *DishSaveOutIdRequest {
	s.PoiId = &v
	return s
}

func (s *DishSaveOutIdRequest) SetAccessToken(v string) *DishSaveOutIdRequest {
	s.AccessToken = &v
	return s
}

func (s *DishSaveOutIdRequest) SetBase(v *DishSaveOutIdRequestBase) *DishSaveOutIdRequest {
	s.Base = v
	return s
}

func (s *DishSaveOutIdRequest) SetAccountId(v string) *DishSaveOutIdRequest {
	s.AccountId = &v
	return s
}

func (s *DishSaveOutIdRequest) SetOutId(v string) *DishSaveOutIdRequest {
	s.OutId = &v
	return s
}

func (s *DishSaveOutIdRequest) SetProductId(v int64) *DishSaveOutIdRequest {
	s.ProductId = &v
	return s
}

func (s *DishSaveOutIdRequest) SetSkuOutIdMap(v map[int64]*string) *DishSaveOutIdRequest {
	s.SkuOutIdMap = v
	return s
}

func (s *DishSaveOutIdRequest) SetHeader(v map[string]*string) *DishSaveOutIdRequest {
	s.Header = v
	return s
}

type DishSaveOutIdRequestBase struct {
	Addr       *string                             `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                             `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                             `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                  `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                             `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishSaveOutIdRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s DishSaveOutIdRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdRequestBase) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdRequestBase) SetAddr(v string) *DishSaveOutIdRequestBase {
	s.Addr = &v
	return s
}

func (s *DishSaveOutIdRequestBase) SetCaller(v string) *DishSaveOutIdRequestBase {
	s.Caller = &v
	return s
}

func (s *DishSaveOutIdRequestBase) SetClient(v string) *DishSaveOutIdRequestBase {
	s.Client = &v
	return s
}

func (s *DishSaveOutIdRequestBase) SetExtra(v map[string]*string) *DishSaveOutIdRequestBase {
	s.Extra = v
	return s
}

func (s *DishSaveOutIdRequestBase) SetLogID(v string) *DishSaveOutIdRequestBase {
	s.LogID = &v
	return s
}

func (s *DishSaveOutIdRequestBase) SetTrafficEnv(v *DishSaveOutIdRequestBaseTrafficEnv) *DishSaveOutIdRequestBase {
	s.TrafficEnv = v
	return s
}

type DishSaveOutIdRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishSaveOutIdRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdRequestBaseTrafficEnv) SetEnv(v string) *DishSaveOutIdRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishSaveOutIdRequestBaseTrafficEnv) SetOpen(v bool) *DishSaveOutIdRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishSaveOutIdResponse struct {
	Extra    *DishSaveOutIdResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishSaveOutIdResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishSaveOutIdResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DishSaveOutIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdResponse) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdResponse) SetExtra(v *DishSaveOutIdResponseExtra) *DishSaveOutIdResponse {
	s.Extra = v
	return s
}

func (s *DishSaveOutIdResponse) SetBaseResp(v *DishSaveOutIdResponseBaseResp) *DishSaveOutIdResponse {
	s.BaseResp = v
	return s
}

func (s *DishSaveOutIdResponse) SetData(v *DishSaveOutIdResponseData) *DishSaveOutIdResponse {
	s.Data = v
	return s
}

type DishSaveOutIdResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishSaveOutIdResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdResponseBaseResp) SetExtra(v map[string]*string) *DishSaveOutIdResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishSaveOutIdResponseBaseResp) SetStatusCode(v int32) *DishSaveOutIdResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishSaveOutIdResponseBaseResp) SetStatusMessage(v string) *DishSaveOutIdResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishSaveOutIdResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DishSaveOutIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdResponseData) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdResponseData) SetErrorCode(v int32) *DishSaveOutIdResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishSaveOutIdResponseData) SetDescription(v string) *DishSaveOutIdResponseData {
	s.Description = &v
	return s
}

type DishSaveOutIdResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s DishSaveOutIdResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishSaveOutIdResponseExtra) GoString() string {
	return s.String()
}

func (s *DishSaveOutIdResponseExtra) SetDescription(v string) *DishSaveOutIdResponseExtra {
	s.Description = &v
	return s
}

func (s *DishSaveOutIdResponseExtra) SetErrorCode(v int32) *DishSaveOutIdResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishSaveOutIdResponseExtra) SetLogid(v string) *DishSaveOutIdResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishSaveOutIdResponseExtra) SetNow(v int64) *DishSaveOutIdResponseExtra {
	s.Now = &v
	return s
}

func (s *DishSaveOutIdResponseExtra) SetSubDescription(v string) *DishSaveOutIdResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishSaveOutIdResponseExtra) SetSubErrorCode(v int32) *DishSaveOutIdResponseExtra {
	s.SubErrorCode = &v
	return s
}

type DishSaveSubmitRequest struct {
	SettleType      *int                                      `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	Skus            []*DishSaveSubmitRequestSkusItem          `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	Header          map[string]*string                        `json:"header,omitempty" xml:"header,omitempty"`
	DishDetailInfo  *DishSaveSubmitRequestDishDetailInfo      `json:"dish_detail_info,omitempty" xml:"dish_detail_info,omitempty"`
	AddDishGroups   []*DishSaveSubmitRequestAddDishGroupsItem `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	Base            *DishSaveSubmitRequestBase                `json:"Base,omitempty" xml:"Base,omitempty"`
	DishDescription *string                                   `json:"dish_description,omitempty" xml:"dish_description,omitempty"`
	Pois            []*DishSaveSubmitRequestPoisItem          `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ProductType     *int                                      `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ApplyDate       *DishSaveSubmitRequestApplyDate           `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	DishGroups      []*DishSaveSubmitRequestDishGroupsItem    `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
	ProductName     *string                                   `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductId       *int64                                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Attributes      []*DishSaveSubmitRequestAttributesItem    `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	DeliveryMethod  []*int                                    `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	CategoryId      *int64                                    `json:"category_id,omitempty" xml:"category_id,omitempty"`
	AccountId       *string                                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccessToken     *string                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ImageList       []*DishSaveSubmitRequestImageListItem     `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
}

func (s DishSaveSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequest) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequest) SetSettleType(v int) *DishSaveSubmitRequest {
	s.SettleType = &v
	return s
}

func (s *DishSaveSubmitRequest) SetSkus(v []*DishSaveSubmitRequestSkusItem) *DishSaveSubmitRequest {
	s.Skus = v
	return s
}

func (s *DishSaveSubmitRequest) SetHeader(v map[string]*string) *DishSaveSubmitRequest {
	s.Header = v
	return s
}

func (s *DishSaveSubmitRequest) SetDishDetailInfo(v *DishSaveSubmitRequestDishDetailInfo) *DishSaveSubmitRequest {
	s.DishDetailInfo = v
	return s
}

func (s *DishSaveSubmitRequest) SetAddDishGroups(v []*DishSaveSubmitRequestAddDishGroupsItem) *DishSaveSubmitRequest {
	s.AddDishGroups = v
	return s
}

func (s *DishSaveSubmitRequest) SetBase(v *DishSaveSubmitRequestBase) *DishSaveSubmitRequest {
	s.Base = v
	return s
}

func (s *DishSaveSubmitRequest) SetDishDescription(v string) *DishSaveSubmitRequest {
	s.DishDescription = &v
	return s
}

func (s *DishSaveSubmitRequest) SetPois(v []*DishSaveSubmitRequestPoisItem) *DishSaveSubmitRequest {
	s.Pois = v
	return s
}

func (s *DishSaveSubmitRequest) SetProductType(v int) *DishSaveSubmitRequest {
	s.ProductType = &v
	return s
}

func (s *DishSaveSubmitRequest) SetApplyDate(v *DishSaveSubmitRequestApplyDate) *DishSaveSubmitRequest {
	s.ApplyDate = v
	return s
}

func (s *DishSaveSubmitRequest) SetDishGroups(v []*DishSaveSubmitRequestDishGroupsItem) *DishSaveSubmitRequest {
	s.DishGroups = v
	return s
}

func (s *DishSaveSubmitRequest) SetProductName(v string) *DishSaveSubmitRequest {
	s.ProductName = &v
	return s
}

func (s *DishSaveSubmitRequest) SetProductId(v int64) *DishSaveSubmitRequest {
	s.ProductId = &v
	return s
}

func (s *DishSaveSubmitRequest) SetAttributes(v []*DishSaveSubmitRequestAttributesItem) *DishSaveSubmitRequest {
	s.Attributes = v
	return s
}

func (s *DishSaveSubmitRequest) SetDeliveryMethod(v []*int) *DishSaveSubmitRequest {
	s.DeliveryMethod = v
	return s
}

func (s *DishSaveSubmitRequest) SetCategoryId(v int64) *DishSaveSubmitRequest {
	s.CategoryId = &v
	return s
}

func (s *DishSaveSubmitRequest) SetAccountId(v string) *DishSaveSubmitRequest {
	s.AccountId = &v
	return s
}

func (s *DishSaveSubmitRequest) SetAccessToken(v string) *DishSaveSubmitRequest {
	s.AccessToken = &v
	return s
}

func (s *DishSaveSubmitRequest) SetImageList(v []*DishSaveSubmitRequestImageListItem) *DishSaveSubmitRequest {
	s.ImageList = v
	return s
}

type DishSaveSubmitRequestAddDishGroupsItem struct {
	GroupId   *int64                                                `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                               `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DishSaveSubmitRequestAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DishSaveSubmitRequestAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestAddDishGroupsItem) SetGroupId(v int64) *DishSaveSubmitRequestAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItem) SetGroupName(v string) *DishSaveSubmitRequestAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItem) SetItemList(v []*DishSaveSubmitRequestAddDishGroupsItemItemListItem) *DishSaveSubmitRequestAddDishGroupsItem {
	s.ItemList = v
	return s
}

type DishSaveSubmitRequestAddDishGroupsItemItemListItem struct {
	Price               *int64                                                     `json:"price,omitempty" xml:"price,omitempty"`
	ProductName         *string                                                    `json:"product_name,omitempty" xml:"product_name,omitempty"`
	IsSetAutoComplement *bool                                                      `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	IsSetSellOut        *bool                                                      `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	PackFee             *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	ProductId           *int64                                                     `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ResidueStock        *int64                                                     `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	MaxStock            *int64                                                     `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	SkuId               *int64                                                     `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s DishSaveSubmitRequestAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetPrice(v int64) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetProductName(v string) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetIsSetAutoComplement(v bool) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetIsSetSellOut(v bool) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetPackFee(v *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.PackFee = v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetProductId(v int64) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetResidueStock(v int64) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.ResidueStock = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetMaxStock(v int64) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.MaxStock = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItem) SetSkuId(v int64) *DishSaveSubmitRequestAddDishGroupsItemItemListItem {
	s.SkuId = &v
	return s
}

type DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee struct {
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) SetPackFee(v int32) *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) SetPackFeeUnit(v int) *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee) SetStep(v int32) *DishSaveSubmitRequestAddDishGroupsItemItemListItemPackFee {
	s.Step = &v
	return s
}

type DishSaveSubmitRequestApplyDate struct {
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
}

func (s DishSaveSubmitRequestApplyDate) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestApplyDate) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestApplyDate) SetDayDuration(v int32) *DishSaveSubmitRequestApplyDate {
	s.DayDuration = &v
	return s
}

func (s *DishSaveSubmitRequestApplyDate) SetUseDateType(v int) *DishSaveSubmitRequestApplyDate {
	s.UseDateType = &v
	return s
}

func (s *DishSaveSubmitRequestApplyDate) SetUseEndDate(v string) *DishSaveSubmitRequestApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *DishSaveSubmitRequestApplyDate) SetUseStartDate(v string) *DishSaveSubmitRequestApplyDate {
	s.UseStartDate = &v
	return s
}

type DishSaveSubmitRequestAttributesItem struct {
	ItemList  []*DishSaveSubmitRequestAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DishSaveSubmitRequestAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestAttributesItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestAttributesItem) SetItemList(v []*DishSaveSubmitRequestAttributesItemItemListItem) *DishSaveSubmitRequestAttributesItem {
	s.ItemList = v
	return s
}

func (s *DishSaveSubmitRequestAttributesItem) SetGroupName(v string) *DishSaveSubmitRequestAttributesItem {
	s.GroupName = &v
	return s
}

type DishSaveSubmitRequestAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DishSaveSubmitRequestAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestAttributesItemItemListItem) SetName(v string) *DishSaveSubmitRequestAttributesItemItemListItem {
	s.Name = &v
	return s
}

type DishSaveSubmitRequestBase struct {
	Extra      map[string]*string                   `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                              `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishSaveSubmitRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                              `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                              `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                              `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s DishSaveSubmitRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestBase) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestBase) SetExtra(v map[string]*string) *DishSaveSubmitRequestBase {
	s.Extra = v
	return s
}

func (s *DishSaveSubmitRequestBase) SetLogID(v string) *DishSaveSubmitRequestBase {
	s.LogID = &v
	return s
}

func (s *DishSaveSubmitRequestBase) SetTrafficEnv(v *DishSaveSubmitRequestBaseTrafficEnv) *DishSaveSubmitRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishSaveSubmitRequestBase) SetAddr(v string) *DishSaveSubmitRequestBase {
	s.Addr = &v
	return s
}

func (s *DishSaveSubmitRequestBase) SetCaller(v string) *DishSaveSubmitRequestBase {
	s.Caller = &v
	return s
}

func (s *DishSaveSubmitRequestBase) SetClient(v string) *DishSaveSubmitRequestBase {
	s.Client = &v
	return s
}

type DishSaveSubmitRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishSaveSubmitRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestBaseTrafficEnv) SetEnv(v string) *DishSaveSubmitRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishSaveSubmitRequestBaseTrafficEnv) SetOpen(v bool) *DishSaveSubmitRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishSaveSubmitRequestDishDetailInfo struct {
	MaterialInfo []*DishSaveSubmitRequestDishDetailInfoMaterialInfoItem `json:"material_info,omitempty" xml:"material_info,omitempty" type:"Repeated"`
}

func (s DishSaveSubmitRequestDishDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestDishDetailInfo) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestDishDetailInfo) SetMaterialInfo(v []*DishSaveSubmitRequestDishDetailInfoMaterialInfoItem) *DishSaveSubmitRequestDishDetailInfo {
	s.MaterialInfo = v
	return s
}

type DishSaveSubmitRequestDishDetailInfoMaterialInfoItem struct {
	Value []*DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Key   *string                                                         `json:"key,omitempty" xml:"key,omitempty"`
}

func (s DishSaveSubmitRequestDishDetailInfoMaterialInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestDishDetailInfoMaterialInfoItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestDishDetailInfoMaterialInfoItem) SetValue(v []*DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem) *DishSaveSubmitRequestDishDetailInfoMaterialInfoItem {
	s.Value = v
	return s
}

func (s *DishSaveSubmitRequestDishDetailInfoMaterialInfoItem) SetKey(v string) *DishSaveSubmitRequestDishDetailInfoMaterialInfoItem {
	s.Key = &v
	return s
}

type DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem) SetValue(v string) *DishSaveSubmitRequestDishDetailInfoMaterialInfoItemValueItem {
	s.Value = &v
	return s
}

type DishSaveSubmitRequestDishGroupsItem struct {
	DishGroupId *int64 `json:"dish_group_id,omitempty" xml:"dish_group_id,omitempty"`
}

func (s DishSaveSubmitRequestDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestDishGroupsItem) SetDishGroupId(v int64) *DishSaveSubmitRequestDishGroupsItem {
	s.DishGroupId = &v
	return s
}

type DishSaveSubmitRequestImageListItem struct {
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s DishSaveSubmitRequestImageListItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestImageListItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestImageListItem) SetUri(v string) *DishSaveSubmitRequestImageListItem {
	s.Uri = &v
	return s
}

type DishSaveSubmitRequestPoisItem struct {
	PoiId *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DishSaveSubmitRequestPoisItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestPoisItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestPoisItem) SetPoiId(v string) *DishSaveSubmitRequestPoisItem {
	s.PoiId = &v
	return s
}

type DishSaveSubmitRequestSkusItem struct {
	SkuSpecAttrs        []*DishSaveSubmitRequestSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	PackFee             *DishSaveSubmitRequestSkusItemPackFee            `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	MaxStock            *int64                                           `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	ActualAmount        *int64                                           `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	IsSetSellOut        *bool                                            `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	IsSetAutoComplement *bool                                            `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	SkuName             *string                                          `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	ResidueStock        *int64                                           `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	SkuId               *int64                                           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s DishSaveSubmitRequestSkusItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestSkusItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestSkusItem) SetSkuSpecAttrs(v []*DishSaveSubmitRequestSkusItemSkuSpecAttrsItem) *DishSaveSubmitRequestSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetPackFee(v *DishSaveSubmitRequestSkusItemPackFee) *DishSaveSubmitRequestSkusItem {
	s.PackFee = v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetMaxStock(v int64) *DishSaveSubmitRequestSkusItem {
	s.MaxStock = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetActualAmount(v int64) *DishSaveSubmitRequestSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetIsSetSellOut(v bool) *DishSaveSubmitRequestSkusItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetIsSetAutoComplement(v bool) *DishSaveSubmitRequestSkusItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetSkuName(v string) *DishSaveSubmitRequestSkusItem {
	s.SkuName = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetResidueStock(v int64) *DishSaveSubmitRequestSkusItem {
	s.ResidueStock = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItem) SetSkuId(v int64) *DishSaveSubmitRequestSkusItem {
	s.SkuId = &v
	return s
}

type DishSaveSubmitRequestSkusItemPackFee struct {
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DishSaveSubmitRequestSkusItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestSkusItemPackFee) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestSkusItemPackFee) SetPackFee(v int32) *DishSaveSubmitRequestSkusItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItemPackFee) SetPackFeeUnit(v int) *DishSaveSubmitRequestSkusItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItemPackFee) SetStep(v int32) *DishSaveSubmitRequestSkusItemPackFee {
	s.Step = &v
	return s
}

type DishSaveSubmitRequestSkusItemSkuSpecAttrsItem struct {
	GroupName *string                                                      `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DishSaveSubmitRequestSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestSkusItemSkuSpecAttrsItem) SetGroupName(v string) *DishSaveSubmitRequestSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItemSkuSpecAttrsItem) SetItemList(v []*DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) *DishSaveSubmitRequestSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

type DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem struct {
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *DishSaveSubmitRequestSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

type DishSaveSubmitResponse struct {
	Extra    *DishSaveSubmitResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DishSaveSubmitResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishSaveSubmitResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DishSaveSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitResponse) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitResponse) SetExtra(v *DishSaveSubmitResponseExtra) *DishSaveSubmitResponse {
	s.Extra = v
	return s
}

func (s *DishSaveSubmitResponse) SetBaseResp(v *DishSaveSubmitResponseBaseResp) *DishSaveSubmitResponse {
	s.BaseResp = v
	return s
}

func (s *DishSaveSubmitResponse) SetData(v *DishSaveSubmitResponseData) *DishSaveSubmitResponse {
	s.Data = v
	return s
}

type DishSaveSubmitResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishSaveSubmitResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitResponseBaseResp) SetExtra(v map[string]*string) *DishSaveSubmitResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishSaveSubmitResponseBaseResp) SetStatusCode(v int32) *DishSaveSubmitResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishSaveSubmitResponseBaseResp) SetStatusMessage(v string) *DishSaveSubmitResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishSaveSubmitResponseData struct {
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishSaveSubmitResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitResponseData) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitResponseData) SetProductId(v int64) *DishSaveSubmitResponseData {
	s.ProductId = &v
	return s
}

func (s *DishSaveSubmitResponseData) SetDescription(v string) *DishSaveSubmitResponseData {
	s.Description = &v
	return s
}

func (s *DishSaveSubmitResponseData) SetErrorCode(v int32) *DishSaveSubmitResponseData {
	s.ErrorCode = &v
	return s
}

type DishSaveSubmitResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s DishSaveSubmitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishSaveSubmitResponseExtra) GoString() string {
	return s.String()
}

func (s *DishSaveSubmitResponseExtra) SetNow(v int64) *DishSaveSubmitResponseExtra {
	s.Now = &v
	return s
}

func (s *DishSaveSubmitResponseExtra) SetSubDescription(v string) *DishSaveSubmitResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishSaveSubmitResponseExtra) SetSubErrorCode(v int32) *DishSaveSubmitResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishSaveSubmitResponseExtra) SetDescription(v string) *DishSaveSubmitResponseExtra {
	s.Description = &v
	return s
}

func (s *DishSaveSubmitResponseExtra) SetErrorCode(v int32) *DishSaveSubmitResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishSaveSubmitResponseExtra) SetLogid(v string) *DishSaveSubmitResponseExtra {
	s.Logid = &v
	return s
}

type DishSortRequest struct {
	AccountId   *string              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GroupId     *int64               `json:"group_id,omitempty" xml:"group_id,omitempty" require:"true"`
	GroupType   *int                 `json:"group_type,omitempty" xml:"group_type,omitempty" require:"true"`
	PoiId       *int64               `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SortList    []*int64             `json:"sort_list,omitempty" xml:"sort_list,omitempty" type:"Repeated"`
	Base        *DishSortRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DishSortRequest) String() string {
	return tea.Prettify(s)
}

func (s DishSortRequest) GoString() string {
	return s.String()
}

func (s *DishSortRequest) SetAccountId(v string) *DishSortRequest {
	s.AccountId = &v
	return s
}

func (s *DishSortRequest) SetGroupId(v int64) *DishSortRequest {
	s.GroupId = &v
	return s
}

func (s *DishSortRequest) SetGroupType(v int) *DishSortRequest {
	s.GroupType = &v
	return s
}

func (s *DishSortRequest) SetPoiId(v int64) *DishSortRequest {
	s.PoiId = &v
	return s
}

func (s *DishSortRequest) SetSortList(v []*int64) *DishSortRequest {
	s.SortList = v
	return s
}

func (s *DishSortRequest) SetBase(v *DishSortRequestBase) *DishSortRequest {
	s.Base = v
	return s
}

func (s *DishSortRequest) SetHeader(v map[string]*string) *DishSortRequest {
	s.Header = v
	return s
}

func (s *DishSortRequest) SetAccessToken(v string) *DishSortRequest {
	s.AccessToken = &v
	return s
}

type DishSortRequestBase struct {
	LogID      *string                        `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishSortRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                        `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                        `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                        `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string             `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishSortRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishSortRequestBase) GoString() string {
	return s.String()
}

func (s *DishSortRequestBase) SetLogID(v string) *DishSortRequestBase {
	s.LogID = &v
	return s
}

func (s *DishSortRequestBase) SetTrafficEnv(v *DishSortRequestBaseTrafficEnv) *DishSortRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishSortRequestBase) SetAddr(v string) *DishSortRequestBase {
	s.Addr = &v
	return s
}

func (s *DishSortRequestBase) SetCaller(v string) *DishSortRequestBase {
	s.Caller = &v
	return s
}

func (s *DishSortRequestBase) SetClient(v string) *DishSortRequestBase {
	s.Client = &v
	return s
}

func (s *DishSortRequestBase) SetExtra(v map[string]*string) *DishSortRequestBase {
	s.Extra = v
	return s
}

type DishSortRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishSortRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishSortRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishSortRequestBaseTrafficEnv) SetEnv(v string) *DishSortRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishSortRequestBaseTrafficEnv) SetOpen(v bool) *DishSortRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishSortResponse struct {
	BaseResp *DishSortResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishSortResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishSortResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishSortResponse) String() string {
	return tea.Prettify(s)
}

func (s DishSortResponse) GoString() string {
	return s.String()
}

func (s *DishSortResponse) SetBaseResp(v *DishSortResponseBaseResp) *DishSortResponse {
	s.BaseResp = v
	return s
}

func (s *DishSortResponse) SetData(v *DishSortResponseData) *DishSortResponse {
	s.Data = v
	return s
}

func (s *DishSortResponse) SetExtra(v *DishSortResponseExtra) *DishSortResponse {
	s.Extra = v
	return s
}

type DishSortResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishSortResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishSortResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishSortResponseBaseResp) SetExtra(v map[string]*string) *DishSortResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishSortResponseBaseResp) SetStatusCode(v int32) *DishSortResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishSortResponseBaseResp) SetStatusMessage(v string) *DishSortResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishSortResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DishSortResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishSortResponseData) GoString() string {
	return s.String()
}

func (s *DishSortResponseData) SetDescription(v string) *DishSortResponseData {
	s.Description = &v
	return s
}

func (s *DishSortResponseData) SetErrorCode(v int32) *DishSortResponseData {
	s.ErrorCode = &v
	return s
}

type DishSortResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s DishSortResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishSortResponseExtra) GoString() string {
	return s.String()
}

func (s *DishSortResponseExtra) SetSubDescription(v string) *DishSortResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishSortResponseExtra) SetSubErrorCode(v int32) *DishSortResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishSortResponseExtra) SetDescription(v string) *DishSortResponseExtra {
	s.Description = &v
	return s
}

func (s *DishSortResponseExtra) SetErrorCode(v int32) *DishSortResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishSortResponseExtra) SetLogid(v string) *DishSortResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishSortResponseExtra) SetNow(v int64) *DishSortResponseExtra {
	s.Now = &v
	return s
}

type DishSyncTaskCreateRequest struct {
	ProductIds  []*int64                       `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base        *DishSyncTaskCreateRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId       []*int64                       `json:"poi_id,omitempty" xml:"poi_id,omitempty" type:"Repeated"`
	Header      map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DishSyncTaskCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateRequest) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateRequest) SetProductIds(v []*int64) *DishSyncTaskCreateRequest {
	s.ProductIds = v
	return s
}

func (s *DishSyncTaskCreateRequest) SetBase(v *DishSyncTaskCreateRequestBase) *DishSyncTaskCreateRequest {
	s.Base = v
	return s
}

func (s *DishSyncTaskCreateRequest) SetAccountId(v string) *DishSyncTaskCreateRequest {
	s.AccountId = &v
	return s
}

func (s *DishSyncTaskCreateRequest) SetPoiId(v []*int64) *DishSyncTaskCreateRequest {
	s.PoiId = v
	return s
}

func (s *DishSyncTaskCreateRequest) SetHeader(v map[string]*string) *DishSyncTaskCreateRequest {
	s.Header = v
	return s
}

func (s *DishSyncTaskCreateRequest) SetAccessToken(v string) *DishSyncTaskCreateRequest {
	s.AccessToken = &v
	return s
}

type DishSyncTaskCreateRequestBase struct {
	Caller     *string                                  `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                  `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                       `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                  `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DishSyncTaskCreateRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                  `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s DishSyncTaskCreateRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateRequestBase) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateRequestBase) SetCaller(v string) *DishSyncTaskCreateRequestBase {
	s.Caller = &v
	return s
}

func (s *DishSyncTaskCreateRequestBase) SetClient(v string) *DishSyncTaskCreateRequestBase {
	s.Client = &v
	return s
}

func (s *DishSyncTaskCreateRequestBase) SetExtra(v map[string]*string) *DishSyncTaskCreateRequestBase {
	s.Extra = v
	return s
}

func (s *DishSyncTaskCreateRequestBase) SetLogID(v string) *DishSyncTaskCreateRequestBase {
	s.LogID = &v
	return s
}

func (s *DishSyncTaskCreateRequestBase) SetTrafficEnv(v *DishSyncTaskCreateRequestBaseTrafficEnv) *DishSyncTaskCreateRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishSyncTaskCreateRequestBase) SetAddr(v string) *DishSyncTaskCreateRequestBase {
	s.Addr = &v
	return s
}

type DishSyncTaskCreateRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s DishSyncTaskCreateRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateRequestBaseTrafficEnv) SetOpen(v bool) *DishSyncTaskCreateRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *DishSyncTaskCreateRequestBaseTrafficEnv) SetEnv(v string) *DishSyncTaskCreateRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type DishSyncTaskCreateResponse struct {
	BaseResp *DishSyncTaskCreateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishSyncTaskCreateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishSyncTaskCreateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishSyncTaskCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateResponse) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateResponse) SetBaseResp(v *DishSyncTaskCreateResponseBaseResp) *DishSyncTaskCreateResponse {
	s.BaseResp = v
	return s
}

func (s *DishSyncTaskCreateResponse) SetData(v *DishSyncTaskCreateResponseData) *DishSyncTaskCreateResponse {
	s.Data = v
	return s
}

func (s *DishSyncTaskCreateResponse) SetExtra(v *DishSyncTaskCreateResponseExtra) *DishSyncTaskCreateResponse {
	s.Extra = v
	return s
}

type DishSyncTaskCreateResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DishSyncTaskCreateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateResponseBaseResp) SetStatusCode(v int32) *DishSyncTaskCreateResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishSyncTaskCreateResponseBaseResp) SetStatusMessage(v string) *DishSyncTaskCreateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DishSyncTaskCreateResponseBaseResp) SetExtra(v map[string]*string) *DishSyncTaskCreateResponseBaseResp {
	s.Extra = v
	return s
}

type DishSyncTaskCreateResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	TaskId      *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DishSyncTaskCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateResponseData) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateResponseData) SetDescription(v string) *DishSyncTaskCreateResponseData {
	s.Description = &v
	return s
}

func (s *DishSyncTaskCreateResponseData) SetErrorCode(v int32) *DishSyncTaskCreateResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishSyncTaskCreateResponseData) SetTaskId(v int64) *DishSyncTaskCreateResponseData {
	s.TaskId = &v
	return s
}

type DishSyncTaskCreateResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DishSyncTaskCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *DishSyncTaskCreateResponseExtra) SetErrorCode(v int32) *DishSyncTaskCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishSyncTaskCreateResponseExtra) SetLogid(v string) *DishSyncTaskCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishSyncTaskCreateResponseExtra) SetNow(v int64) *DishSyncTaskCreateResponseExtra {
	s.Now = &v
	return s
}

func (s *DishSyncTaskCreateResponseExtra) SetSubDescription(v string) *DishSyncTaskCreateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DishSyncTaskCreateResponseExtra) SetSubErrorCode(v int32) *DishSyncTaskCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishSyncTaskCreateResponseExtra) SetDescription(v string) *DishSyncTaskCreateResponseExtra {
	s.Description = &v
	return s
}

type DishSyncTaskRecordRequest struct {
	Base        *DishSyncTaskRecordRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageNo      *int32                         `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize    *int32                         `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TaskId      *int64                         `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s DishSyncTaskRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordRequest) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordRequest) SetBase(v *DishSyncTaskRecordRequestBase) *DishSyncTaskRecordRequest {
	s.Base = v
	return s
}

func (s *DishSyncTaskRecordRequest) SetAccountId(v string) *DishSyncTaskRecordRequest {
	s.AccountId = &v
	return s
}

func (s *DishSyncTaskRecordRequest) SetHeader(v map[string]*string) *DishSyncTaskRecordRequest {
	s.Header = v
	return s
}

func (s *DishSyncTaskRecordRequest) SetAccessToken(v string) *DishSyncTaskRecordRequest {
	s.AccessToken = &v
	return s
}

func (s *DishSyncTaskRecordRequest) SetPageNo(v int32) *DishSyncTaskRecordRequest {
	s.PageNo = &v
	return s
}

func (s *DishSyncTaskRecordRequest) SetPageSize(v int32) *DishSyncTaskRecordRequest {
	s.PageSize = &v
	return s
}

func (s *DishSyncTaskRecordRequest) SetTaskId(v int64) *DishSyncTaskRecordRequest {
	s.TaskId = &v
	return s
}

type DishSyncTaskRecordRequestBase struct {
	TrafficEnv *DishSyncTaskRecordRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                  `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                  `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                  `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                       `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                  `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s DishSyncTaskRecordRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordRequestBase) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordRequestBase) SetTrafficEnv(v *DishSyncTaskRecordRequestBaseTrafficEnv) *DishSyncTaskRecordRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishSyncTaskRecordRequestBase) SetAddr(v string) *DishSyncTaskRecordRequestBase {
	s.Addr = &v
	return s
}

func (s *DishSyncTaskRecordRequestBase) SetCaller(v string) *DishSyncTaskRecordRequestBase {
	s.Caller = &v
	return s
}

func (s *DishSyncTaskRecordRequestBase) SetClient(v string) *DishSyncTaskRecordRequestBase {
	s.Client = &v
	return s
}

func (s *DishSyncTaskRecordRequestBase) SetExtra(v map[string]*string) *DishSyncTaskRecordRequestBase {
	s.Extra = v
	return s
}

func (s *DishSyncTaskRecordRequestBase) SetLogID(v string) *DishSyncTaskRecordRequestBase {
	s.LogID = &v
	return s
}

type DishSyncTaskRecordRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishSyncTaskRecordRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordRequestBaseTrafficEnv) SetEnv(v string) *DishSyncTaskRecordRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishSyncTaskRecordRequestBaseTrafficEnv) SetOpen(v bool) *DishSyncTaskRecordRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishSyncTaskRecordResponse struct {
	BaseResp *DishSyncTaskRecordResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishSyncTaskRecordResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishSyncTaskRecordResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishSyncTaskRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordResponse) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordResponse) SetBaseResp(v *DishSyncTaskRecordResponseBaseResp) *DishSyncTaskRecordResponse {
	s.BaseResp = v
	return s
}

func (s *DishSyncTaskRecordResponse) SetData(v *DishSyncTaskRecordResponseData) *DishSyncTaskRecordResponse {
	s.Data = v
	return s
}

func (s *DishSyncTaskRecordResponse) SetExtra(v *DishSyncTaskRecordResponseExtra) *DishSyncTaskRecordResponse {
	s.Extra = v
	return s
}

type DishSyncTaskRecordResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishSyncTaskRecordResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordResponseBaseResp) SetExtra(v map[string]*string) *DishSyncTaskRecordResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishSyncTaskRecordResponseBaseResp) SetStatusCode(v int32) *DishSyncTaskRecordResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishSyncTaskRecordResponseBaseResp) SetStatusMessage(v string) *DishSyncTaskRecordResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishSyncTaskRecordResponseData struct {
	Total          *int32                                              `json:"total,omitempty" xml:"total,omitempty"`
	Description    *string                                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32                                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore        *bool                                               `json:"has_more,omitempty" xml:"has_more,omitempty"`
	SubTaskRecords []*DishSyncTaskRecordResponseDataSubTaskRecordsItem `json:"sub_task_records,omitempty" xml:"sub_task_records,omitempty" type:"Repeated"`
}

func (s DishSyncTaskRecordResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordResponseData) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordResponseData) SetTotal(v int32) *DishSyncTaskRecordResponseData {
	s.Total = &v
	return s
}

func (s *DishSyncTaskRecordResponseData) SetDescription(v string) *DishSyncTaskRecordResponseData {
	s.Description = &v
	return s
}

func (s *DishSyncTaskRecordResponseData) SetErrorCode(v int32) *DishSyncTaskRecordResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishSyncTaskRecordResponseData) SetHasMore(v bool) *DishSyncTaskRecordResponseData {
	s.HasMore = &v
	return s
}

func (s *DishSyncTaskRecordResponseData) SetSubTaskRecords(v []*DishSyncTaskRecordResponseDataSubTaskRecordsItem) *DishSyncTaskRecordResponseData {
	s.SubTaskRecords = v
	return s
}

type DishSyncTaskRecordResponseDataSubTaskRecordsItem struct {
	SrcProductId *int64  `json:"src_product_id,omitempty" xml:"src_product_id,omitempty"`
	Status       *int32  `json:"status,omitempty" xml:"status,omitempty"`
	DstProductId *int64  `json:"dst_product_id,omitempty" xml:"dst_product_id,omitempty"`
	FailReason   *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	PoiId        *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DishSyncTaskRecordResponseDataSubTaskRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordResponseDataSubTaskRecordsItem) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordResponseDataSubTaskRecordsItem) SetSrcProductId(v int64) *DishSyncTaskRecordResponseDataSubTaskRecordsItem {
	s.SrcProductId = &v
	return s
}

func (s *DishSyncTaskRecordResponseDataSubTaskRecordsItem) SetStatus(v int32) *DishSyncTaskRecordResponseDataSubTaskRecordsItem {
	s.Status = &v
	return s
}

func (s *DishSyncTaskRecordResponseDataSubTaskRecordsItem) SetDstProductId(v int64) *DishSyncTaskRecordResponseDataSubTaskRecordsItem {
	s.DstProductId = &v
	return s
}

func (s *DishSyncTaskRecordResponseDataSubTaskRecordsItem) SetFailReason(v string) *DishSyncTaskRecordResponseDataSubTaskRecordsItem {
	s.FailReason = &v
	return s
}

func (s *DishSyncTaskRecordResponseDataSubTaskRecordsItem) SetPoiId(v int64) *DishSyncTaskRecordResponseDataSubTaskRecordsItem {
	s.PoiId = &v
	return s
}

type DishSyncTaskRecordResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s DishSyncTaskRecordResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishSyncTaskRecordResponseExtra) GoString() string {
	return s.String()
}

func (s *DishSyncTaskRecordResponseExtra) SetSubErrorCode(v int32) *DishSyncTaskRecordResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishSyncTaskRecordResponseExtra) SetDescription(v string) *DishSyncTaskRecordResponseExtra {
	s.Description = &v
	return s
}

func (s *DishSyncTaskRecordResponseExtra) SetErrorCode(v int32) *DishSyncTaskRecordResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishSyncTaskRecordResponseExtra) SetLogid(v string) *DishSyncTaskRecordResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishSyncTaskRecordResponseExtra) SetNow(v int64) *DishSyncTaskRecordResponseExtra {
	s.Now = &v
	return s
}

func (s *DishSyncTaskRecordResponseExtra) SetSubDescription(v string) *DishSyncTaskRecordResponseExtra {
	s.SubDescription = &v
	return s
}

type DishUpdateStockRequest struct {
	AccountId       *string                                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DishSkuStocks   []*DishUpdateStockRequestDishSkuStocksItem `json:"dish_sku_stocks,omitempty" xml:"dish_sku_stocks,omitempty" type:"Repeated"`
	Header          map[string]*string                         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DyPoiId         *int64                                     `json:"dy_poi_id,omitempty" xml:"dy_poi_id,omitempty"`
	MerchantOperate *bool                                      `json:"merchant_operate,omitempty" xml:"merchant_operate,omitempty"`
	Base            *DishUpdateStockRequestBase                `json:"Base,omitempty" xml:"Base,omitempty"`
}

func (s DishUpdateStockRequest) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockRequest) GoString() string {
	return s.String()
}

func (s *DishUpdateStockRequest) SetAccountId(v string) *DishUpdateStockRequest {
	s.AccountId = &v
	return s
}

func (s *DishUpdateStockRequest) SetDishSkuStocks(v []*DishUpdateStockRequestDishSkuStocksItem) *DishUpdateStockRequest {
	s.DishSkuStocks = v
	return s
}

func (s *DishUpdateStockRequest) SetHeader(v map[string]*string) *DishUpdateStockRequest {
	s.Header = v
	return s
}

func (s *DishUpdateStockRequest) SetAccessToken(v string) *DishUpdateStockRequest {
	s.AccessToken = &v
	return s
}

func (s *DishUpdateStockRequest) SetDyPoiId(v int64) *DishUpdateStockRequest {
	s.DyPoiId = &v
	return s
}

func (s *DishUpdateStockRequest) SetMerchantOperate(v bool) *DishUpdateStockRequest {
	s.MerchantOperate = &v
	return s
}

func (s *DishUpdateStockRequest) SetBase(v *DishUpdateStockRequestBase) *DishUpdateStockRequest {
	s.Base = v
	return s
}

type DishUpdateStockRequestBase struct {
	TrafficEnv *DishUpdateStockRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                               `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                               `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                               `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                    `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                               `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s DishUpdateStockRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockRequestBase) GoString() string {
	return s.String()
}

func (s *DishUpdateStockRequestBase) SetTrafficEnv(v *DishUpdateStockRequestBaseTrafficEnv) *DishUpdateStockRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DishUpdateStockRequestBase) SetAddr(v string) *DishUpdateStockRequestBase {
	s.Addr = &v
	return s
}

func (s *DishUpdateStockRequestBase) SetCaller(v string) *DishUpdateStockRequestBase {
	s.Caller = &v
	return s
}

func (s *DishUpdateStockRequestBase) SetClient(v string) *DishUpdateStockRequestBase {
	s.Client = &v
	return s
}

func (s *DishUpdateStockRequestBase) SetExtra(v map[string]*string) *DishUpdateStockRequestBase {
	s.Extra = v
	return s
}

func (s *DishUpdateStockRequestBase) SetLogID(v string) *DishUpdateStockRequestBase {
	s.LogID = &v
	return s
}

type DishUpdateStockRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DishUpdateStockRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DishUpdateStockRequestBaseTrafficEnv) SetEnv(v string) *DishUpdateStockRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DishUpdateStockRequestBaseTrafficEnv) SetOpen(v bool) *DishUpdateStockRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DishUpdateStockRequestDishSkuStocksItem struct {
	ProductId *int64                                                  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuStocks []*DishUpdateStockRequestDishSkuStocksItemSkuStocksItem `json:"sku_stocks,omitempty" xml:"sku_stocks,omitempty" type:"Repeated"`
}

func (s DishUpdateStockRequestDishSkuStocksItem) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockRequestDishSkuStocksItem) GoString() string {
	return s.String()
}

func (s *DishUpdateStockRequestDishSkuStocksItem) SetProductId(v int64) *DishUpdateStockRequestDishSkuStocksItem {
	s.ProductId = &v
	return s
}

func (s *DishUpdateStockRequestDishSkuStocksItem) SetSkuStocks(v []*DishUpdateStockRequestDishSkuStocksItemSkuStocksItem) *DishUpdateStockRequestDishSkuStocksItem {
	s.SkuStocks = v
	return s
}

type DishUpdateStockRequestDishSkuStocksItemSkuStocksItem struct {
	AvailStock *int64 `json:"avail_stock,omitempty" xml:"avail_stock,omitempty"`
	SkuId      *int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s DishUpdateStockRequestDishSkuStocksItemSkuStocksItem) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockRequestDishSkuStocksItemSkuStocksItem) GoString() string {
	return s.String()
}

func (s *DishUpdateStockRequestDishSkuStocksItemSkuStocksItem) SetAvailStock(v int64) *DishUpdateStockRequestDishSkuStocksItemSkuStocksItem {
	s.AvailStock = &v
	return s
}

func (s *DishUpdateStockRequestDishSkuStocksItemSkuStocksItem) SetSkuId(v int64) *DishUpdateStockRequestDishSkuStocksItemSkuStocksItem {
	s.SkuId = &v
	return s
}

type DishUpdateStockResponse struct {
	BaseResp *DishUpdateStockResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DishUpdateStockResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DishUpdateStockResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DishUpdateStockResponse) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockResponse) GoString() string {
	return s.String()
}

func (s *DishUpdateStockResponse) SetBaseResp(v *DishUpdateStockResponseBaseResp) *DishUpdateStockResponse {
	s.BaseResp = v
	return s
}

func (s *DishUpdateStockResponse) SetData(v *DishUpdateStockResponseData) *DishUpdateStockResponse {
	s.Data = v
	return s
}

func (s *DishUpdateStockResponse) SetExtra(v *DishUpdateStockResponseExtra) *DishUpdateStockResponse {
	s.Extra = v
	return s
}

type DishUpdateStockResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DishUpdateStockResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DishUpdateStockResponseBaseResp) SetExtra(v map[string]*string) *DishUpdateStockResponseBaseResp {
	s.Extra = v
	return s
}

func (s *DishUpdateStockResponseBaseResp) SetStatusCode(v int32) *DishUpdateStockResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DishUpdateStockResponseBaseResp) SetStatusMessage(v string) *DishUpdateStockResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type DishUpdateStockResponseData struct {
	ErrorCode   *int32             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrMap      map[string]*string `json:"err_map,omitempty" xml:"err_map,omitempty"`
}

func (s DishUpdateStockResponseData) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockResponseData) GoString() string {
	return s.String()
}

func (s *DishUpdateStockResponseData) SetErrorCode(v int32) *DishUpdateStockResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DishUpdateStockResponseData) SetDescription(v string) *DishUpdateStockResponseData {
	s.Description = &v
	return s
}

func (s *DishUpdateStockResponseData) SetErrMap(v map[string]*string) *DishUpdateStockResponseData {
	s.ErrMap = v
	return s
}

type DishUpdateStockResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s DishUpdateStockResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DishUpdateStockResponseExtra) GoString() string {
	return s.String()
}

func (s *DishUpdateStockResponseExtra) SetSubErrorCode(v int32) *DishUpdateStockResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DishUpdateStockResponseExtra) SetDescription(v string) *DishUpdateStockResponseExtra {
	s.Description = &v
	return s
}

func (s *DishUpdateStockResponseExtra) SetErrorCode(v int32) *DishUpdateStockResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DishUpdateStockResponseExtra) SetLogid(v string) *DishUpdateStockResponseExtra {
	s.Logid = &v
	return s
}

func (s *DishUpdateStockResponseExtra) SetNow(v int64) *DishUpdateStockResponseExtra {
	s.Now = &v
	return s
}

func (s *DishUpdateStockResponseExtra) SetSubDescription(v string) *DishUpdateStockResponseExtra {
	s.SubDescription = &v
	return s
}

type DouyinCreateInteractTaskRequest struct {
	InteractRules map[string]*DouyinCreateInteractTaskRequestInteractRulesValue `json:"interact_rules,omitempty" xml:"interact_rules,omitempty" require:"true"`
	Tags          []*string                                                     `json:"tags,omitempty" xml:"tags,omitempty" require:"true" type:"Repeated"`
	AppId         *string                                                       `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header        map[string]*string                                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	EndTime       *int64                                                        `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	MountLink     *string                                                       `json:"mount_link,omitempty" xml:"mount_link,omitempty" require:"true"`
	MaxCount      *int64                                                        `json:"max_count,omitempty" xml:"max_count,omitempty" require:"true"`
	TaskType      *int32                                                        `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	StartTime     *int64                                                        `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	PublishType   []*int32                                                      `json:"publish_type,omitempty" xml:"publish_type,omitempty" require:"true" type:"Repeated"`
}

func (s DouyinCreateInteractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateInteractTaskRequest) GoString() string {
	return s.String()
}

func (s *DouyinCreateInteractTaskRequest) SetInteractRules(v map[string]*DouyinCreateInteractTaskRequestInteractRulesValue) *DouyinCreateInteractTaskRequest {
	s.InteractRules = v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetTags(v []*string) *DouyinCreateInteractTaskRequest {
	s.Tags = v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetAppId(v string) *DouyinCreateInteractTaskRequest {
	s.AppId = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetHeader(v map[string]*string) *DouyinCreateInteractTaskRequest {
	s.Header = v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetAccessToken(v string) *DouyinCreateInteractTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetEndTime(v int64) *DouyinCreateInteractTaskRequest {
	s.EndTime = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetMountLink(v string) *DouyinCreateInteractTaskRequest {
	s.MountLink = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetMaxCount(v int64) *DouyinCreateInteractTaskRequest {
	s.MaxCount = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetTaskType(v int32) *DouyinCreateInteractTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetStartTime(v int64) *DouyinCreateInteractTaskRequest {
	s.StartTime = &v
	return s
}

func (s *DouyinCreateInteractTaskRequest) SetPublishType(v []*int32) *DouyinCreateInteractTaskRequest {
	s.PublishType = v
	return s
}

type DouyinCreateInteractTaskRequestInteractRulesValue struct {
	StageCount  *int64 `json:"stage_count,omitempty" xml:"stage_count,omitempty"`
	TargetCount *int64 `json:"target_count,omitempty" xml:"target_count,omitempty" require:"true"`
}

func (s DouyinCreateInteractTaskRequestInteractRulesValue) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateInteractTaskRequestInteractRulesValue) GoString() string {
	return s.String()
}

func (s *DouyinCreateInteractTaskRequestInteractRulesValue) SetStageCount(v int64) *DouyinCreateInteractTaskRequestInteractRulesValue {
	s.StageCount = &v
	return s
}

func (s *DouyinCreateInteractTaskRequestInteractRulesValue) SetTargetCount(v int64) *DouyinCreateInteractTaskRequestInteractRulesValue {
	s.TargetCount = &v
	return s
}

type DouyinCreateInteractTaskResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DouyinCreateInteractTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DouyinCreateInteractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateInteractTaskResponse) GoString() string {
	return s.String()
}

func (s *DouyinCreateInteractTaskResponse) SetErrMsg(v string) *DouyinCreateInteractTaskResponse {
	s.ErrMsg = &v
	return s
}

func (s *DouyinCreateInteractTaskResponse) SetLogId(v string) *DouyinCreateInteractTaskResponse {
	s.LogId = &v
	return s
}

func (s *DouyinCreateInteractTaskResponse) SetData(v *DouyinCreateInteractTaskResponseData) *DouyinCreateInteractTaskResponse {
	s.Data = v
	return s
}

func (s *DouyinCreateInteractTaskResponse) SetErrNo(v int32) *DouyinCreateInteractTaskResponse {
	s.ErrNo = &v
	return s
}

type DouyinCreateInteractTaskResponseData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s DouyinCreateInteractTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateInteractTaskResponseData) GoString() string {
	return s.String()
}

func (s *DouyinCreateInteractTaskResponseData) SetTaskId(v string) *DouyinCreateInteractTaskResponseData {
	s.TaskId = &v
	return s
}

type DouyinCreateTaskRequest struct {
	MountLink   *string            `json:"mount_link,omitempty" xml:"mount_link,omitempty" require:"true"`
	TaskType    *int32             `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	PublishType []*int32           `json:"publish_type,omitempty" xml:"publish_type,omitempty" require:"true" type:"Repeated"`
	RuleType    *int32             `json:"rule_type,omitempty" xml:"rule_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	Tags        []*string          `json:"tags,omitempty" xml:"tags,omitempty" require:"true" type:"Repeated"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TargetCount *int64             `json:"target_count,omitempty" xml:"target_count,omitempty" require:"true"`
}

func (s DouyinCreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateTaskRequest) GoString() string {
	return s.String()
}

func (s *DouyinCreateTaskRequest) SetMountLink(v string) *DouyinCreateTaskRequest {
	s.MountLink = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetTaskType(v int32) *DouyinCreateTaskRequest {
	s.TaskType = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetPublishType(v []*int32) *DouyinCreateTaskRequest {
	s.PublishType = v
	return s
}

func (s *DouyinCreateTaskRequest) SetRuleType(v int32) *DouyinCreateTaskRequest {
	s.RuleType = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetHeader(v map[string]*string) *DouyinCreateTaskRequest {
	s.Header = v
	return s
}

func (s *DouyinCreateTaskRequest) SetTags(v []*string) *DouyinCreateTaskRequest {
	s.Tags = v
	return s
}

func (s *DouyinCreateTaskRequest) SetStartTime(v int64) *DouyinCreateTaskRequest {
	s.StartTime = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetEndTime(v int64) *DouyinCreateTaskRequest {
	s.EndTime = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetAppId(v string) *DouyinCreateTaskRequest {
	s.AppId = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetAccessToken(v string) *DouyinCreateTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *DouyinCreateTaskRequest) SetTargetCount(v int64) *DouyinCreateTaskRequest {
	s.TargetCount = &v
	return s
}

type DouyinCreateTaskResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DouyinCreateTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DouyinCreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateTaskResponse) GoString() string {
	return s.String()
}

func (s *DouyinCreateTaskResponse) SetLogId(v string) *DouyinCreateTaskResponse {
	s.LogId = &v
	return s
}

func (s *DouyinCreateTaskResponse) SetData(v *DouyinCreateTaskResponseData) *DouyinCreateTaskResponse {
	s.Data = v
	return s
}

func (s *DouyinCreateTaskResponse) SetErrNo(v int32) *DouyinCreateTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *DouyinCreateTaskResponse) SetErrMsg(v string) *DouyinCreateTaskResponse {
	s.ErrMsg = &v
	return s
}

type DouyinCreateTaskResponseData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s DouyinCreateTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s DouyinCreateTaskResponseData) GoString() string {
	return s.String()
}

func (s *DouyinCreateTaskResponseData) SetTaskId(v string) *DouyinCreateTaskResponseData {
	s.TaskId = &v
	return s
}

type DouyinQueryUserTaskRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TaskId      []*string          `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true" type:"Repeated"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DouyinQueryUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DouyinQueryUserTaskRequest) GoString() string {
	return s.String()
}

func (s *DouyinQueryUserTaskRequest) SetOpenId(v string) *DouyinQueryUserTaskRequest {
	s.OpenId = &v
	return s
}

func (s *DouyinQueryUserTaskRequest) SetTaskId(v []*string) *DouyinQueryUserTaskRequest {
	s.TaskId = v
	return s
}

func (s *DouyinQueryUserTaskRequest) SetAppId(v string) *DouyinQueryUserTaskRequest {
	s.AppId = &v
	return s
}

func (s *DouyinQueryUserTaskRequest) SetHeader(v map[string]*string) *DouyinQueryUserTaskRequest {
	s.Header = v
	return s
}

func (s *DouyinQueryUserTaskRequest) SetAccessToken(v string) *DouyinQueryUserTaskRequest {
	s.AccessToken = &v
	return s
}

type DouyinQueryUserTaskResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DouyinQueryUserTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DouyinQueryUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DouyinQueryUserTaskResponse) GoString() string {
	return s.String()
}

func (s *DouyinQueryUserTaskResponse) SetLogId(v string) *DouyinQueryUserTaskResponse {
	s.LogId = &v
	return s
}

func (s *DouyinQueryUserTaskResponse) SetData(v *DouyinQueryUserTaskResponseData) *DouyinQueryUserTaskResponse {
	s.Data = v
	return s
}

func (s *DouyinQueryUserTaskResponse) SetErrNo(v int32) *DouyinQueryUserTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *DouyinQueryUserTaskResponse) SetErrMsg(v string) *DouyinQueryUserTaskResponse {
	s.ErrMsg = &v
	return s
}

type DouyinQueryUserTaskResponseData struct {
	TaskInfoList map[string]*DouyinQueryUserTaskResponseDataTaskInfoListValue `json:"task_info_list,omitempty" xml:"task_info_list,omitempty" require:"true"`
}

func (s DouyinQueryUserTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s DouyinQueryUserTaskResponseData) GoString() string {
	return s.String()
}

func (s *DouyinQueryUserTaskResponseData) SetTaskInfoList(v map[string]*DouyinQueryUserTaskResponseDataTaskInfoListValue) *DouyinQueryUserTaskResponseData {
	s.TaskInfoList = v
	return s
}

type DouyinQueryUserTaskResponseDataTaskInfoListValue struct {
	ErrNo        *int32                                                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg       *string                                                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	TaskId       *string                                                          `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	VideoInfo    []*DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem `json:"video_info,omitempty" xml:"video_info,omitempty" require:"true" type:"Repeated"`
	IsValid      *bool                                                            `json:"is_valid,omitempty" xml:"is_valid,omitempty" require:"true"`
	SuccessCount *int64                                                           `json:"success_count,omitempty" xml:"success_count,omitempty" require:"true"`
	TargetCount  *int64                                                           `json:"target_count,omitempty" xml:"target_count,omitempty" require:"true"`
	Completed    *bool                                                            `json:"completed,omitempty" xml:"completed,omitempty" require:"true"`
}

func (s DouyinQueryUserTaskResponseDataTaskInfoListValue) String() string {
	return tea.Prettify(s)
}

func (s DouyinQueryUserTaskResponseDataTaskInfoListValue) GoString() string {
	return s.String()
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetErrNo(v int32) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.ErrNo = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetErrMsg(v string) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.ErrMsg = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetTaskId(v string) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.TaskId = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetVideoInfo(v []*DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.VideoInfo = v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetIsValid(v bool) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.IsValid = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetSuccessCount(v int64) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.SuccessCount = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetTargetCount(v int64) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.TargetCount = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValue) SetCompleted(v bool) *DouyinQueryUserTaskResponseDataTaskInfoListValue {
	s.Completed = &v
	return s
}

type DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem struct {
	VideoId     *string `json:"video_id,omitempty" xml:"video_id,omitempty"`
	VideoStatus *int32  `json:"video_status,omitempty" xml:"video_status,omitempty"`
}

func (s DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem) GoString() string {
	return s.String()
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem) SetVideoId(v string) *DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.VideoId = &v
	return s
}

func (s *DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem) SetVideoStatus(v int32) *DouyinQueryUserTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.VideoStatus = &v
	return s
}

type DraftGetRequest struct {
	AccountId    *string              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GetDraftType *int                 `json:"get_draft_type,omitempty" xml:"get_draft_type,omitempty"`
	Header       map[string]*string   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductIds   []*string            `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base         *DraftGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
}

func (s DraftGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DraftGetRequest) GoString() string {
	return s.String()
}

func (s *DraftGetRequest) SetAccountId(v string) *DraftGetRequest {
	s.AccountId = &v
	return s
}

func (s *DraftGetRequest) SetGetDraftType(v int) *DraftGetRequest {
	s.GetDraftType = &v
	return s
}

func (s *DraftGetRequest) SetHeader(v map[string]*string) *DraftGetRequest {
	s.Header = v
	return s
}

func (s *DraftGetRequest) SetAccessToken(v string) *DraftGetRequest {
	s.AccessToken = &v
	return s
}

func (s *DraftGetRequest) SetProductIds(v []*string) *DraftGetRequest {
	s.ProductIds = v
	return s
}

func (s *DraftGetRequest) SetBase(v *DraftGetRequestBase) *DraftGetRequest {
	s.Base = v
	return s
}

type DraftGetRequestBase struct {
	Addr       *string                        `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                        `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                        `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string             `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                        `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DraftGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s DraftGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DraftGetRequestBase) GoString() string {
	return s.String()
}

func (s *DraftGetRequestBase) SetAddr(v string) *DraftGetRequestBase {
	s.Addr = &v
	return s
}

func (s *DraftGetRequestBase) SetCaller(v string) *DraftGetRequestBase {
	s.Caller = &v
	return s
}

func (s *DraftGetRequestBase) SetClient(v string) *DraftGetRequestBase {
	s.Client = &v
	return s
}

func (s *DraftGetRequestBase) SetExtra(v map[string]*string) *DraftGetRequestBase {
	s.Extra = v
	return s
}

func (s *DraftGetRequestBase) SetLogID(v string) *DraftGetRequestBase {
	s.LogID = &v
	return s
}

func (s *DraftGetRequestBase) SetTrafficEnv(v *DraftGetRequestBaseTrafficEnv) *DraftGetRequestBase {
	s.TrafficEnv = v
	return s
}

type DraftGetRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DraftGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DraftGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DraftGetRequestBaseTrafficEnv) SetEnv(v string) *DraftGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DraftGetRequestBaseTrafficEnv) SetOpen(v bool) *DraftGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DraftGetResponse struct {
	BaseResp *DraftGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *DraftGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DraftGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DraftGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponse) GoString() string {
	return s.String()
}

func (s *DraftGetResponse) SetBaseResp(v *DraftGetResponseBaseResp) *DraftGetResponse {
	s.BaseResp = v
	return s
}

func (s *DraftGetResponse) SetData(v *DraftGetResponseData) *DraftGetResponse {
	s.Data = v
	return s
}

func (s *DraftGetResponse) SetExtra(v *DraftGetResponseExtra) *DraftGetResponse {
	s.Extra = v
	return s
}

type DraftGetResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DraftGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DraftGetResponseBaseResp) SetStatusCode(v int32) *DraftGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DraftGetResponseBaseResp) SetStatusMessage(v string) *DraftGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DraftGetResponseBaseResp) SetExtra(v map[string]*string) *DraftGetResponseBaseResp {
	s.Extra = v
	return s
}

type DraftGetResponseData struct {
	Description *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Products    []*DraftGetResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
}

func (s DraftGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseData) GoString() string {
	return s.String()
}

func (s *DraftGetResponseData) SetDescription(v string) *DraftGetResponseData {
	s.Description = &v
	return s
}

func (s *DraftGetResponseData) SetErrorCode(v int32) *DraftGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *DraftGetResponseData) SetProducts(v []*DraftGetResponseDataProductsItem) *DraftGetResponseData {
	s.Products = v
	return s
}

type DraftGetResponseDataProductsItem struct {
	ProductSpec   *int                                               `json:"product_spec,omitempty" xml:"product_spec,omitempty"`
	Sku           *DraftGetResponseDataProductsItemSku               `json:"sku,omitempty" xml:"sku,omitempty"`
	Skus          []*DraftGetResponseDataProductsItemSkusItem        `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	AddDishList   []*DraftGetResponseDataProductsItemAddDishListItem `json:"add_dish_list,omitempty" xml:"add_dish_list,omitempty" type:"Repeated"`
	AuditMsg      *string                                            `json:"audit_msg,omitempty" xml:"audit_msg,omitempty"`
	DraftStatus   *int                                               `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	IsSuperSkuSub *bool                                              `json:"is_super_sku_sub,omitempty" xml:"is_super_sku_sub,omitempty"`
	Product       *DraftGetResponseDataProductsItemProduct           `json:"product,omitempty" xml:"product,omitempty"`
}

func (s DraftGetResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItem) SetProductSpec(v int) *DraftGetResponseDataProductsItem {
	s.ProductSpec = &v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetSku(v *DraftGetResponseDataProductsItemSku) *DraftGetResponseDataProductsItem {
	s.Sku = v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetSkus(v []*DraftGetResponseDataProductsItemSkusItem) *DraftGetResponseDataProductsItem {
	s.Skus = v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetAddDishList(v []*DraftGetResponseDataProductsItemAddDishListItem) *DraftGetResponseDataProductsItem {
	s.AddDishList = v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetAuditMsg(v string) *DraftGetResponseDataProductsItem {
	s.AuditMsg = &v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetDraftStatus(v int) *DraftGetResponseDataProductsItem {
	s.DraftStatus = &v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetIsSuperSkuSub(v bool) *DraftGetResponseDataProductsItem {
	s.IsSuperSkuSub = &v
	return s
}

func (s *DraftGetResponseDataProductsItem) SetProduct(v *DraftGetResponseDataProductsItemProduct) *DraftGetResponseDataProductsItem {
	s.Product = v
	return s
}

type DraftGetResponseDataProductsItemAddDishListItem struct {
	ItemList  []*DraftGetResponseDataProductsItemAddDishListItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                        `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DraftGetResponseDataProductsItemAddDishListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemAddDishListItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemAddDishListItem) SetItemList(v []*DraftGetResponseDataProductsItemAddDishListItemItemListItem) *DraftGetResponseDataProductsItemAddDishListItem {
	s.ItemList = v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItem) SetGroupName(v string) *DraftGetResponseDataProductsItemAddDishListItem {
	s.GroupName = &v
	return s
}

type DraftGetResponseDataProductsItemAddDishListItemItemListItem struct {
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	OutSkuId  *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuId     *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Stock     *int32  `json:"stock,omitempty" xml:"stock,omitempty"`
	AddPrice  *int32  `json:"add_price,omitempty" xml:"add_price,omitempty"`
	ItemName  *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

func (s DraftGetResponseDataProductsItemAddDishListItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemAddDishListItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetOutId(v string) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.OutId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetOutSkuId(v string) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetProductId(v string) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetSkuId(v string) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetStock(v int32) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.Stock = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetAddPrice(v int32) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.AddPrice = &v
	return s
}

func (s *DraftGetResponseDataProductsItemAddDishListItemItemListItem) SetItemName(v string) *DraftGetResponseDataProductsItemAddDishListItemItemListItem {
	s.ItemName = &v
	return s
}

type DraftGetResponseDataProductsItemProduct struct {
	Desc             *string                                            `json:"desc,omitempty" xml:"desc,omitempty"`
	OutId            *string                                            `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductExt       *DraftGetResponseDataProductsItemProductProductExt `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	Extra            *string                                            `json:"extra,omitempty" xml:"extra,omitempty"`
	AccountId        *string                                            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	CreatorAccountId *int64                                             `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	CategoryId       *int64                                             `json:"category_id,omitempty" xml:"category_id,omitempty"`
	BizLine          *int                                               `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	Version          *int64                                             `json:"version,omitempty" xml:"version,omitempty"`
	CreateTime       *int64                                             `json:"create_time,omitempty" xml:"create_time,omitempty"`
	UpdateTime       *int64                                             `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductName      *string                                            `json:"product_name,omitempty" xml:"product_name,omitempty"`
	AccountName      *string                                            `json:"account_name,omitempty" xml:"account_name,omitempty"`
	CategoryFullName *string                                            `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	OwnerAccountId   *int64                                             `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	SoldStartTime    *int64                                             `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	ProductSpec      *int                                               `json:"product_spec,omitempty" xml:"product_spec,omitempty"`
	OutUrl           *string                                            `json:"out_url,omitempty" xml:"out_url,omitempty"`
	ProductType      *int                                               `json:"product_type,omitempty" xml:"product_type,omitempty"`
	AttrKeyValueMap  map[string]*string                                 `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	SoldEndTime      *int64                                             `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	ProductId        *string                                            `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s DraftGetResponseDataProductsItemProduct) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemProduct) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemProduct) SetDesc(v string) *DraftGetResponseDataProductsItemProduct {
	s.Desc = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetOutId(v string) *DraftGetResponseDataProductsItemProduct {
	s.OutId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetProductExt(v *DraftGetResponseDataProductsItemProductProductExt) *DraftGetResponseDataProductsItemProduct {
	s.ProductExt = v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetExtra(v string) *DraftGetResponseDataProductsItemProduct {
	s.Extra = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetAccountId(v string) *DraftGetResponseDataProductsItemProduct {
	s.AccountId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetCreatorAccountId(v int64) *DraftGetResponseDataProductsItemProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetCategoryId(v int64) *DraftGetResponseDataProductsItemProduct {
	s.CategoryId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetBizLine(v int) *DraftGetResponseDataProductsItemProduct {
	s.BizLine = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetVersion(v int64) *DraftGetResponseDataProductsItemProduct {
	s.Version = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetCreateTime(v int64) *DraftGetResponseDataProductsItemProduct {
	s.CreateTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetUpdateTime(v int64) *DraftGetResponseDataProductsItemProduct {
	s.UpdateTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetProductName(v string) *DraftGetResponseDataProductsItemProduct {
	s.ProductName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetAccountName(v string) *DraftGetResponseDataProductsItemProduct {
	s.AccountName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetCategoryFullName(v string) *DraftGetResponseDataProductsItemProduct {
	s.CategoryFullName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetOwnerAccountId(v int64) *DraftGetResponseDataProductsItemProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetSoldStartTime(v int64) *DraftGetResponseDataProductsItemProduct {
	s.SoldStartTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetProductSpec(v int) *DraftGetResponseDataProductsItemProduct {
	s.ProductSpec = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetOutUrl(v string) *DraftGetResponseDataProductsItemProduct {
	s.OutUrl = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetProductType(v int) *DraftGetResponseDataProductsItemProduct {
	s.ProductType = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetAttrKeyValueMap(v map[string]*string) *DraftGetResponseDataProductsItemProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetSoldEndTime(v int64) *DraftGetResponseDataProductsItemProduct {
	s.SoldEndTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProduct) SetProductId(v string) *DraftGetResponseDataProductsItemProduct {
	s.ProductId = &v
	return s
}

type DraftGetResponseDataProductsItemProductProductExt struct {
	TestExtra *DraftGetResponseDataProductsItemProductProductExtTestExtra `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
}

func (s DraftGetResponseDataProductsItemProductProductExt) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemProductProductExt) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemProductProductExt) SetTestExtra(v *DraftGetResponseDataProductsItemProductProductExtTestExtra) *DraftGetResponseDataProductsItemProductProductExt {
	s.TestExtra = v
	return s
}

type DraftGetResponseDataProductsItemProductProductExtTestExtra struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
}

func (s DraftGetResponseDataProductsItemProductProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemProductProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemProductProductExtTestExtra) SetTestFlag(v bool) *DraftGetResponseDataProductsItemProductProductExtTestExtra {
	s.TestFlag = &v
	return s
}

func (s *DraftGetResponseDataProductsItemProductProductExtTestExtra) SetUids(v []*string) *DraftGetResponseDataProductsItemProductProductExtTestExtra {
	s.Uids = v
	return s
}

type DraftGetResponseDataProductsItemSku struct {
	BindSkus        []*string                                       `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuId           *string                                         `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	AttrKeyValueMap map[string]*string                              `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	OutSkuId        *string                                         `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Stock           *DraftGetResponseDataProductsItemSkuStock       `json:"stock,omitempty" xml:"stock,omitempty"`
	OriginAmount    *int64                                          `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ActualAmount    *int64                                          `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	Specs           []*DraftGetResponseDataProductsItemSkuSpecsItem `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	Status          *int                                            `json:"status,omitempty" xml:"status,omitempty"`
	Extra           *string                                         `json:"extra,omitempty" xml:"extra,omitempty"`
	SkuName         *string                                         `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	UpdateTime      *int64                                          `json:"update_time,omitempty" xml:"update_time,omitempty"`
	CreateTime      *int64                                          `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s DraftGetResponseDataProductsItemSku) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSku) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSku) SetBindSkus(v []*string) *DraftGetResponseDataProductsItemSku {
	s.BindSkus = v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetSkuId(v string) *DraftGetResponseDataProductsItemSku {
	s.SkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetAttrKeyValueMap(v map[string]*string) *DraftGetResponseDataProductsItemSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetOutSkuId(v string) *DraftGetResponseDataProductsItemSku {
	s.OutSkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetStock(v *DraftGetResponseDataProductsItemSkuStock) *DraftGetResponseDataProductsItemSku {
	s.Stock = v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetOriginAmount(v int64) *DraftGetResponseDataProductsItemSku {
	s.OriginAmount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetActualAmount(v int64) *DraftGetResponseDataProductsItemSku {
	s.ActualAmount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetSpecs(v []*DraftGetResponseDataProductsItemSkuSpecsItem) *DraftGetResponseDataProductsItemSku {
	s.Specs = v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetStatus(v int) *DraftGetResponseDataProductsItemSku {
	s.Status = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetExtra(v string) *DraftGetResponseDataProductsItemSku {
	s.Extra = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetSkuName(v string) *DraftGetResponseDataProductsItemSku {
	s.SkuName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetUpdateTime(v int64) *DraftGetResponseDataProductsItemSku {
	s.UpdateTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSku) SetCreateTime(v int64) *DraftGetResponseDataProductsItemSku {
	s.CreateTime = &v
	return s
}

type DraftGetResponseDataProductsItemSkuSpecsItem struct {
	GroupName *string                                                     `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DraftGetResponseDataProductsItemSkuSpecsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DraftGetResponseDataProductsItemSkuSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkuSpecsItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItem) SetGroupName(v string) *DraftGetResponseDataProductsItemSkuSpecsItem {
	s.GroupName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItem) SetItemList(v []*DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) *DraftGetResponseDataProductsItemSkuSpecsItem {
	s.ItemList = v
	return s
}

type DraftGetResponseDataProductsItemSkuSpecsItemItemListItem struct {
	OutSkuId  *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId     *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Unit      *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight    *string `json:"weight,omitempty" xml:"weight,omitempty"`
	AddPrice  *int32  `json:"add_price,omitempty" xml:"add_price,omitempty"`
	IsDefault *bool   `json:"is_default,omitempty" xml:"is_default,omitempty"`
	ItemName  *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

func (s DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetOutSkuId(v string) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetSkuId(v string) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetUnit(v string) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetWeight(v string) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetAddPrice(v int32) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.AddPrice = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetIsDefault(v bool) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.IsDefault = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem) SetItemName(v string) *DraftGetResponseDataProductsItemSkuSpecsItemItemListItem {
	s.ItemName = &v
	return s
}

type DraftGetResponseDataProductsItemSkuStock struct {
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
}

func (s DraftGetResponseDataProductsItemSkuStock) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkuStock) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetStockQty(v int64) *DraftGetResponseDataProductsItemSkuStock {
	s.StockQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetAvailQty(v int64) *DraftGetResponseDataProductsItemSkuStock {
	s.AvailQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetFrozenQty(v int64) *DraftGetResponseDataProductsItemSkuStock {
	s.FrozenQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetLimitType(v int) *DraftGetResponseDataProductsItemSkuStock {
	s.LimitType = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetSoldCount(v int64) *DraftGetResponseDataProductsItemSkuStock {
	s.SoldCount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkuStock) SetSoldQty(v int64) *DraftGetResponseDataProductsItemSkuStock {
	s.SoldQty = &v
	return s
}

type DraftGetResponseDataProductsItemSkusItem struct {
	SkuName         *string                                              `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	OutSkuId        *string                                              `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	AttrKeyValueMap map[string]*string                                   `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	OriginAmount    *int64                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	UpdateTime      *int64                                               `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ActualAmount    *int64                                               `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	Specs           []*DraftGetResponseDataProductsItemSkusItemSpecsItem `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	Stock           *DraftGetResponseDataProductsItemSkusItemStock       `json:"stock,omitempty" xml:"stock,omitempty"`
	BindSkus        []*string                                            `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuId           *string                                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	CreateTime      *int64                                               `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Extra           *string                                              `json:"extra,omitempty" xml:"extra,omitempty"`
	Status          *int                                                 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DraftGetResponseDataProductsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkusItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetSkuName(v string) *DraftGetResponseDataProductsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetOutSkuId(v string) *DraftGetResponseDataProductsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetAttrKeyValueMap(v map[string]*string) *DraftGetResponseDataProductsItemSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetOriginAmount(v int64) *DraftGetResponseDataProductsItemSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetUpdateTime(v int64) *DraftGetResponseDataProductsItemSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetActualAmount(v int64) *DraftGetResponseDataProductsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetSpecs(v []*DraftGetResponseDataProductsItemSkusItemSpecsItem) *DraftGetResponseDataProductsItemSkusItem {
	s.Specs = v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetStock(v *DraftGetResponseDataProductsItemSkusItemStock) *DraftGetResponseDataProductsItemSkusItem {
	s.Stock = v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetBindSkus(v []*string) *DraftGetResponseDataProductsItemSkusItem {
	s.BindSkus = v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetSkuId(v string) *DraftGetResponseDataProductsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetCreateTime(v int64) *DraftGetResponseDataProductsItemSkusItem {
	s.CreateTime = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetExtra(v string) *DraftGetResponseDataProductsItemSkusItem {
	s.Extra = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItem) SetStatus(v int) *DraftGetResponseDataProductsItemSkusItem {
	s.Status = &v
	return s
}

type DraftGetResponseDataProductsItemSkusItemSpecsItem struct {
	ItemList  []*DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                          `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DraftGetResponseDataProductsItemSkusItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkusItemSpecsItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItem) SetItemList(v []*DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) *DraftGetResponseDataProductsItemSkusItemSpecsItem {
	s.ItemList = v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItem) SetGroupName(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItem {
	s.GroupName = &v
	return s
}

type DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem struct {
	ItemName  *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	OutSkuId  *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId     *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Unit      *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight    *string `json:"weight,omitempty" xml:"weight,omitempty"`
	AddPrice  *int32  `json:"add_price,omitempty" xml:"add_price,omitempty"`
	IsDefault *bool   `json:"is_default,omitempty" xml:"is_default,omitempty"`
}

func (s DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetItemName(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetOutSkuId(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetSkuId(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetUnit(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetWeight(v string) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetAddPrice(v int32) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.AddPrice = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem) SetIsDefault(v bool) *DraftGetResponseDataProductsItemSkusItemSpecsItemItemListItem {
	s.IsDefault = &v
	return s
}

type DraftGetResponseDataProductsItemSkusItemStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s DraftGetResponseDataProductsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseDataProductsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetSoldCount(v int64) *DraftGetResponseDataProductsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetSoldQty(v int64) *DraftGetResponseDataProductsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetStockQty(v int64) *DraftGetResponseDataProductsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetAvailQty(v int64) *DraftGetResponseDataProductsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetFrozenQty(v int64) *DraftGetResponseDataProductsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *DraftGetResponseDataProductsItemSkusItemStock) SetLimitType(v int) *DraftGetResponseDataProductsItemSkusItemStock {
	s.LimitType = &v
	return s
}

type DraftGetResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s DraftGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DraftGetResponseExtra) GoString() string {
	return s.String()
}

func (s *DraftGetResponseExtra) SetSubDescription(v string) *DraftGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DraftGetResponseExtra) SetSubErrorCode(v int32) *DraftGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DraftGetResponseExtra) SetDescription(v string) *DraftGetResponseExtra {
	s.Description = &v
	return s
}

func (s *DraftGetResponseExtra) SetErrorCode(v int32) *DraftGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DraftGetResponseExtra) SetLogid(v string) *DraftGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *DraftGetResponseExtra) SetNow(v int64) *DraftGetResponseExtra {
	s.Now = &v
	return s
}

type DraftListRequest struct {
	AccessToken *string               `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DishType    *int                  `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	PageNo      *int32                `json:"page_no,omitempty" xml:"page_no,omitempty"`
	Base        *DraftListRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string    `json:"header,omitempty" xml:"header,omitempty"`
	AccountId   *string               `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PageSize    *int32                `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PoiId       *int64                `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Status      []*int                `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s DraftListRequest) String() string {
	return tea.Prettify(s)
}

func (s DraftListRequest) GoString() string {
	return s.String()
}

func (s *DraftListRequest) SetAccessToken(v string) *DraftListRequest {
	s.AccessToken = &v
	return s
}

func (s *DraftListRequest) SetDishType(v int) *DraftListRequest {
	s.DishType = &v
	return s
}

func (s *DraftListRequest) SetPageNo(v int32) *DraftListRequest {
	s.PageNo = &v
	return s
}

func (s *DraftListRequest) SetBase(v *DraftListRequestBase) *DraftListRequest {
	s.Base = v
	return s
}

func (s *DraftListRequest) SetHeader(v map[string]*string) *DraftListRequest {
	s.Header = v
	return s
}

func (s *DraftListRequest) SetAccountId(v string) *DraftListRequest {
	s.AccountId = &v
	return s
}

func (s *DraftListRequest) SetPageSize(v int32) *DraftListRequest {
	s.PageSize = &v
	return s
}

func (s *DraftListRequest) SetPoiId(v int64) *DraftListRequest {
	s.PoiId = &v
	return s
}

func (s *DraftListRequest) SetStatus(v []*int) *DraftListRequest {
	s.Status = v
	return s
}

type DraftListRequestBase struct {
	Caller     *string                         `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                         `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string              `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                         `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *DraftListRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                         `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s DraftListRequestBase) String() string {
	return tea.Prettify(s)
}

func (s DraftListRequestBase) GoString() string {
	return s.String()
}

func (s *DraftListRequestBase) SetCaller(v string) *DraftListRequestBase {
	s.Caller = &v
	return s
}

func (s *DraftListRequestBase) SetClient(v string) *DraftListRequestBase {
	s.Client = &v
	return s
}

func (s *DraftListRequestBase) SetExtra(v map[string]*string) *DraftListRequestBase {
	s.Extra = v
	return s
}

func (s *DraftListRequestBase) SetLogID(v string) *DraftListRequestBase {
	s.LogID = &v
	return s
}

func (s *DraftListRequestBase) SetTrafficEnv(v *DraftListRequestBaseTrafficEnv) *DraftListRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *DraftListRequestBase) SetAddr(v string) *DraftListRequestBase {
	s.Addr = &v
	return s
}

type DraftListRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s DraftListRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s DraftListRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *DraftListRequestBaseTrafficEnv) SetEnv(v string) *DraftListRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *DraftListRequestBaseTrafficEnv) SetOpen(v bool) *DraftListRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type DraftListResponse struct {
	Data     *DraftListResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *DraftListResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *DraftListResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s DraftListResponse) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponse) GoString() string {
	return s.String()
}

func (s *DraftListResponse) SetData(v *DraftListResponseData) *DraftListResponse {
	s.Data = v
	return s
}

func (s *DraftListResponse) SetExtra(v *DraftListResponseExtra) *DraftListResponse {
	s.Extra = v
	return s
}

func (s *DraftListResponse) SetBaseResp(v *DraftListResponseBaseResp) *DraftListResponse {
	s.BaseResp = v
	return s
}

type DraftListResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s DraftListResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseBaseResp) GoString() string {
	return s.String()
}

func (s *DraftListResponseBaseResp) SetStatusCode(v int32) *DraftListResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *DraftListResponseBaseResp) SetStatusMessage(v string) *DraftListResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *DraftListResponseBaseResp) SetExtra(v map[string]*string) *DraftListResponseBaseResp {
	s.Extra = v
	return s
}

type DraftListResponseData struct {
	HasMore     *bool                             `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Total       *int32                            `json:"total,omitempty" xml:"total,omitempty"`
	Description *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Dishs       []*DraftListResponseDataDishsItem `json:"dishs,omitempty" xml:"dishs,omitempty" type:"Repeated"`
	ErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DraftListResponseData) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseData) GoString() string {
	return s.String()
}

func (s *DraftListResponseData) SetHasMore(v bool) *DraftListResponseData {
	s.HasMore = &v
	return s
}

func (s *DraftListResponseData) SetTotal(v int32) *DraftListResponseData {
	s.Total = &v
	return s
}

func (s *DraftListResponseData) SetDescription(v string) *DraftListResponseData {
	s.Description = &v
	return s
}

func (s *DraftListResponseData) SetDishs(v []*DraftListResponseDataDishsItem) *DraftListResponseData {
	s.Dishs = v
	return s
}

func (s *DraftListResponseData) SetErrorCode(v int32) *DraftListResponseData {
	s.ErrorCode = &v
	return s
}

type DraftListResponseDataDishsItem struct {
	DraftStatus       *int                                                  `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	ApplyDate         *DraftListResponseDataDishsItemApplyDate              `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	AccountId         *string                                               `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DishGroups        []*DraftListResponseDataDishsItemDishGroupsItem       `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
	DeliveryMethod    []*int                                                `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	AccountName       *string                                               `json:"account_name,omitempty" xml:"account_name,omitempty"`
	ProductId         *int64                                                `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SettleType        *int64                                                `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	DishDetailInfo    *DraftListResponseDataDishsItemDishDetailInfo         `json:"dish_detail_info,omitempty" xml:"dish_detail_info,omitempty"`
	UpdateTime        *int64                                                `json:"update_time,omitempty" xml:"update_time,omitempty"`
	AddDishGroups     []*DraftListResponseDataDishsItemAddDishGroupsItem    `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	Attributes        []*DraftListResponseDataDishsItemAttributesItem       `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	Skus              []*DraftListResponseDataDishsItemSkusItem             `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	IsBindMerchant    *bool                                                 `json:"is_bind_merchant,omitempty" xml:"is_bind_merchant,omitempty"`
	CreateTime        *int64                                                `json:"create_time,omitempty" xml:"create_time,omitempty"`
	DishDescription   *string                                               `json:"dish_description,omitempty" xml:"dish_description,omitempty"`
	ImageList         []*DraftListResponseDataDishsItemImageListItem        `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	CategoryId        *int64                                                `json:"category_id,omitempty" xml:"category_id,omitempty"`
	ProductSpecAttrs  []*DraftListResponseDataDishsItemProductSpecAttrsItem `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	OutId             *string                                               `json:"out_id,omitempty" xml:"out_id,omitempty"`
	DishType          *int                                                  `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	PoiCount          *int64                                                `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	MerchantProductId *string                                               `json:"merchant_product_id,omitempty" xml:"merchant_product_id,omitempty"`
	Pois              []*DraftListResponseDataDishsItemPoisItem             `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ProductName       *string                                               `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s DraftListResponseDataDishsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItem) SetDraftStatus(v int) *DraftListResponseDataDishsItem {
	s.DraftStatus = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetApplyDate(v *DraftListResponseDataDishsItemApplyDate) *DraftListResponseDataDishsItem {
	s.ApplyDate = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetAccountId(v string) *DraftListResponseDataDishsItem {
	s.AccountId = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetDishGroups(v []*DraftListResponseDataDishsItemDishGroupsItem) *DraftListResponseDataDishsItem {
	s.DishGroups = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetDeliveryMethod(v []*int) *DraftListResponseDataDishsItem {
	s.DeliveryMethod = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetAccountName(v string) *DraftListResponseDataDishsItem {
	s.AccountName = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetProductId(v int64) *DraftListResponseDataDishsItem {
	s.ProductId = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetSettleType(v int64) *DraftListResponseDataDishsItem {
	s.SettleType = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetDishDetailInfo(v *DraftListResponseDataDishsItemDishDetailInfo) *DraftListResponseDataDishsItem {
	s.DishDetailInfo = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetUpdateTime(v int64) *DraftListResponseDataDishsItem {
	s.UpdateTime = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetAddDishGroups(v []*DraftListResponseDataDishsItemAddDishGroupsItem) *DraftListResponseDataDishsItem {
	s.AddDishGroups = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetAttributes(v []*DraftListResponseDataDishsItemAttributesItem) *DraftListResponseDataDishsItem {
	s.Attributes = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetSkus(v []*DraftListResponseDataDishsItemSkusItem) *DraftListResponseDataDishsItem {
	s.Skus = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetIsBindMerchant(v bool) *DraftListResponseDataDishsItem {
	s.IsBindMerchant = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetCreateTime(v int64) *DraftListResponseDataDishsItem {
	s.CreateTime = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetDishDescription(v string) *DraftListResponseDataDishsItem {
	s.DishDescription = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetImageList(v []*DraftListResponseDataDishsItemImageListItem) *DraftListResponseDataDishsItem {
	s.ImageList = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetCategoryId(v int64) *DraftListResponseDataDishsItem {
	s.CategoryId = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetProductSpecAttrs(v []*DraftListResponseDataDishsItemProductSpecAttrsItem) *DraftListResponseDataDishsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetOutId(v string) *DraftListResponseDataDishsItem {
	s.OutId = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetDishType(v int) *DraftListResponseDataDishsItem {
	s.DishType = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetPoiCount(v int64) *DraftListResponseDataDishsItem {
	s.PoiCount = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetMerchantProductId(v string) *DraftListResponseDataDishsItem {
	s.MerchantProductId = &v
	return s
}

func (s *DraftListResponseDataDishsItem) SetPois(v []*DraftListResponseDataDishsItemPoisItem) *DraftListResponseDataDishsItem {
	s.Pois = v
	return s
}

func (s *DraftListResponseDataDishsItem) SetProductName(v string) *DraftListResponseDataDishsItem {
	s.ProductName = &v
	return s
}

type DraftListResponseDataDishsItemAddDishGroupsItem struct {
	GroupId   *int64                                                         `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                        `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DraftListResponseDataDishsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DraftListResponseDataDishsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItem) SetGroupId(v int64) *DraftListResponseDataDishsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItem) SetGroupName(v string) *DraftListResponseDataDishsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItem) SetItemList(v []*DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) *DraftListResponseDataDishsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

type DraftListResponseDataDishsItemAddDishGroupsItemItemListItem struct {
	Price               *int64                                                              `json:"price,omitempty" xml:"price,omitempty"`
	ProductId           *int64                                                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ResidueStock        *int64                                                              `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	SkuId               *int64                                                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OutSkuId            *string                                                             `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	ProductName         *string                                                             `json:"product_name,omitempty" xml:"product_name,omitempty"`
	OutId               *string                                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	IsSetAutoComplement *bool                                                               `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	MaxStock            *int64                                                              `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	PackFee             *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	IsSetSellOut        *bool                                                               `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
}

func (s DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetResidueStock(v int64) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ResidueStock = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetSkuId(v int64) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutSkuId(v string) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductName(v string) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutId(v string) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetAutoComplement(v bool) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetMaxStock(v int64) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.MaxStock = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetPackFee(v *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.PackFee = v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetSellOut(v bool) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetSellOut = &v
	return s
}

type DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee struct {
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFee(v int32) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFeeUnit(v int) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetStep(v int32) *DraftListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.Step = &v
	return s
}

type DraftListResponseDataDishsItemApplyDate struct {
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
}

func (s DraftListResponseDataDishsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemApplyDate) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemApplyDate) SetUseDateType(v int) *DraftListResponseDataDishsItemApplyDate {
	s.UseDateType = &v
	return s
}

func (s *DraftListResponseDataDishsItemApplyDate) SetUseEndDate(v string) *DraftListResponseDataDishsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *DraftListResponseDataDishsItemApplyDate) SetUseStartDate(v string) *DraftListResponseDataDishsItemApplyDate {
	s.UseStartDate = &v
	return s
}

func (s *DraftListResponseDataDishsItemApplyDate) SetDayDuration(v int32) *DraftListResponseDataDishsItemApplyDate {
	s.DayDuration = &v
	return s
}

type DraftListResponseDataDishsItemAttributesItem struct {
	GroupName *string                                                     `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DraftListResponseDataDishsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DraftListResponseDataDishsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemAttributesItem) SetGroupName(v string) *DraftListResponseDataDishsItemAttributesItem {
	s.GroupName = &v
	return s
}

func (s *DraftListResponseDataDishsItemAttributesItem) SetItemList(v []*DraftListResponseDataDishsItemAttributesItemItemListItem) *DraftListResponseDataDishsItemAttributesItem {
	s.ItemList = v
	return s
}

type DraftListResponseDataDishsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DraftListResponseDataDishsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemAttributesItemItemListItem) SetName(v string) *DraftListResponseDataDishsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type DraftListResponseDataDishsItemDishDetailInfo struct {
	MaterialInfo []*DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem `json:"material_info,omitempty" xml:"material_info,omitempty" type:"Repeated"`
}

func (s DraftListResponseDataDishsItemDishDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemDishDetailInfo) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemDishDetailInfo) SetMaterialInfo(v []*DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) *DraftListResponseDataDishsItemDishDetailInfo {
	s.MaterialInfo = v
	return s
}

type DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem struct {
	Name  *string   `json:"name,omitempty" xml:"name,omitempty"`
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Key   *string   `json:"key,omitempty" xml:"key,omitempty"`
}

func (s DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetName(v string) *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Name = &v
	return s
}

func (s *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetValue(v []*string) *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Value = v
	return s
}

func (s *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetKey(v string) *DraftListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Key = &v
	return s
}

type DraftListResponseDataDishsItemDishGroupsItem struct {
	DishGroupId       *int64  `json:"dish_group_id,omitempty" xml:"dish_group_id,omitempty"`
	DishGroupName     *string `json:"dish_group_name,omitempty" xml:"dish_group_name,omitempty"`
	GroupRankingScore *string `json:"group_ranking_score,omitempty" xml:"group_ranking_score,omitempty"`
}

func (s DraftListResponseDataDishsItemDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemDishGroupsItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemDishGroupsItem) SetDishGroupId(v int64) *DraftListResponseDataDishsItemDishGroupsItem {
	s.DishGroupId = &v
	return s
}

func (s *DraftListResponseDataDishsItemDishGroupsItem) SetDishGroupName(v string) *DraftListResponseDataDishsItemDishGroupsItem {
	s.DishGroupName = &v
	return s
}

func (s *DraftListResponseDataDishsItemDishGroupsItem) SetGroupRankingScore(v string) *DraftListResponseDataDishsItemDishGroupsItem {
	s.GroupRankingScore = &v
	return s
}
