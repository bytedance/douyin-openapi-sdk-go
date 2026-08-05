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

type ItemReplyCommentResponse struct {
	Data   *ItemReplyCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ItemReplyCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemReplyCommentResponse) GoString() string {
	return s.String()
}

func (s *ItemReplyCommentResponse) SetData(v *ItemReplyCommentResponseData) *ItemReplyCommentResponse {
	s.Data = v
	return s
}

func (s *ItemReplyCommentResponse) SetErrNo(v int32) *ItemReplyCommentResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemReplyCommentResponse) SetErrMsg(v string) *ItemReplyCommentResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemReplyCommentResponse) SetLogId(v string) *ItemReplyCommentResponse {
	s.LogId = &v
	return s
}

type ItemReplyCommentResponseData struct {
	Extra *ItemReplyCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemReplyCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemReplyCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemReplyCommentResponseData) GoString() string {
	return s.String()
}

func (s *ItemReplyCommentResponseData) SetExtra(v *ItemReplyCommentResponseDataExtra) *ItemReplyCommentResponseData {
	s.Extra = v
	return s
}

func (s *ItemReplyCommentResponseData) SetData(v *ItemReplyCommentResponseDataData) *ItemReplyCommentResponseData {
	s.Data = v
	return s
}

type ItemReplyCommentResponseDataData struct {
	CommentId *string `json:"comment_id,omitempty" xml:"comment_id,omitempty"`
	NickName  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	Avatar    *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
}

func (s ItemReplyCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemReplyCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemReplyCommentResponseDataData) SetCommentId(v string) *ItemReplyCommentResponseDataData {
	s.CommentId = &v
	return s
}

func (s *ItemReplyCommentResponseDataData) SetNickName(v string) *ItemReplyCommentResponseDataData {
	s.NickName = &v
	return s
}

func (s *ItemReplyCommentResponseDataData) SetAvatar(v string) *ItemReplyCommentResponseDataData {
	s.Avatar = &v
	return s
}

type ItemReplyCommentResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ItemReplyCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemReplyCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemReplyCommentResponseDataExtra) SetSubDescription(v string) *ItemReplyCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemReplyCommentResponseDataExtra) SetLogid(v string) *ItemReplyCommentResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemReplyCommentResponseDataExtra) SetNow(v int64) *ItemReplyCommentResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemReplyCommentResponseDataExtra) SetErrorCode(v int32) *ItemReplyCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemReplyCommentResponseDataExtra) SetDescription(v string) *ItemReplyCommentResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemReplyCommentResponseDataExtra) SetSubErrorCode(v int32) *ItemReplyCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type JsGetticketRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s JsGetticketRequest) String() string {
	return tea.Prettify(s)
}

func (s JsGetticketRequest) GoString() string {
	return s.String()
}

func (s *JsGetticketRequest) SetHeader(v map[string]*string) *JsGetticketRequest {
	s.Header = v
	return s
}

func (s *JsGetticketRequest) SetAccessToken(v string) *JsGetticketRequest {
	s.AccessToken = &v
	return s
}

type JsGetticketResponse struct {
	Data  *JsGetticketResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *JsGetticketResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s JsGetticketResponse) String() string {
	return tea.Prettify(s)
}

func (s JsGetticketResponse) GoString() string {
	return s.String()
}

func (s *JsGetticketResponse) SetData(v *JsGetticketResponseData) *JsGetticketResponse {
	s.Data = v
	return s
}

func (s *JsGetticketResponse) SetExtra(v *JsGetticketResponseExtra) *JsGetticketResponse {
	s.Extra = v
	return s
}

type JsGetticketResponseData struct {
	Ticket        *string `json:"ticket,omitempty" xml:"ticket,omitempty" require:"true"`
	ExpiresIn     *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s JsGetticketResponseData) String() string {
	return tea.Prettify(s)
}

func (s JsGetticketResponseData) GoString() string {
	return s.String()
}

func (s *JsGetticketResponseData) SetTicket(v string) *JsGetticketResponseData {
	s.Ticket = &v
	return s
}

func (s *JsGetticketResponseData) SetExpiresIn(v int64) *JsGetticketResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *JsGetticketResponseData) SetGwErrorCode(v int32) *JsGetticketResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *JsGetticketResponseData) SetGwDescription(v string) *JsGetticketResponseData {
	s.GwDescription = &v
	return s
}

type JsGetticketResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s JsGetticketResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s JsGetticketResponseExtra) GoString() string {
	return s.String()
}

func (s *JsGetticketResponseExtra) SetDescription(v string) *JsGetticketResponseExtra {
	s.Description = &v
	return s
}

func (s *JsGetticketResponseExtra) SetSubErrorCode(v int32) *JsGetticketResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *JsGetticketResponseExtra) SetSubDescription(v string) *JsGetticketResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *JsGetticketResponseExtra) SetLogid(v string) *JsGetticketResponseExtra {
	s.Logid = &v
	return s
}

func (s *JsGetticketResponseExtra) SetNow(v int64) *JsGetticketResponseExtra {
	s.Now = &v
	return s
}

func (s *JsGetticketResponseExtra) SetErrorCode(v int32) *JsGetticketResponseExtra {
	s.ErrorCode = &v
	return s
}

type LedgerDetailedQueryRequest struct {
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Size        *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s LedgerDetailedQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryRequest) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryRequest) SetCursor(v string) *LedgerDetailedQueryRequest {
	s.Cursor = &v
	return s
}

func (s *LedgerDetailedQueryRequest) SetSize(v int64) *LedgerDetailedQueryRequest {
	s.Size = &v
	return s
}

func (s *LedgerDetailedQueryRequest) SetAccountId(v string) *LedgerDetailedQueryRequest {
	s.AccountId = &v
	return s
}

func (s *LedgerDetailedQueryRequest) SetBillDate(v string) *LedgerDetailedQueryRequest {
	s.BillDate = &v
	return s
}

func (s *LedgerDetailedQueryRequest) SetHeader(v map[string]*string) *LedgerDetailedQueryRequest {
	s.Header = v
	return s
}

func (s *LedgerDetailedQueryRequest) SetAccessToken(v string) *LedgerDetailedQueryRequest {
	s.AccessToken = &v
	return s
}

type LedgerDetailedQueryResponse struct {
	Extra *LedgerDetailedQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *LedgerDetailedQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s LedgerDetailedQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponse) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponse) SetExtra(v *LedgerDetailedQueryResponseExtra) *LedgerDetailedQueryResponse {
	s.Extra = v
	return s
}

func (s *LedgerDetailedQueryResponse) SetData(v *LedgerDetailedQueryResponseData) *LedgerDetailedQueryResponse {
	s.Data = v
	return s
}

type LedgerDetailedQueryResponseData struct {
	GwDescription *string                                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor        *string                                             `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                               `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	LedgerRecords []*LedgerDetailedQueryResponseDataLedgerRecordsItem `json:"ledger_records,omitempty" xml:"ledger_records,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s LedgerDetailedQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseData) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseData) SetGwDescription(v string) *LedgerDetailedQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *LedgerDetailedQueryResponseData) SetCursor(v string) *LedgerDetailedQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *LedgerDetailedQueryResponseData) SetHasMore(v bool) *LedgerDetailedQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *LedgerDetailedQueryResponseData) SetLedgerRecords(v []*LedgerDetailedQueryResponseDataLedgerRecordsItem) *LedgerDetailedQueryResponseData {
	s.LedgerRecords = v
	return s
}

func (s *LedgerDetailedQueryResponseData) SetGwErrorCode(v int32) *LedgerDetailedQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type LedgerDetailedQueryResponseDataLedgerRecordsItem struct {
	LedgerTime      *int64                                                           `json:"ledger_time,omitempty" xml:"ledger_time,omitempty"`
	OrderAttrribute *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute `json:"order_attrribute,omitempty" xml:"order_attrribute,omitempty"`
	PoiId           *string                                                          `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Id              *string                                                          `json:"id,omitempty" xml:"id,omitempty"`
	SkuOrderId      *string                                                          `json:"sku_order_id,omitempty" xml:"sku_order_id,omitempty"`
	IsMallStore     *bool                                                            `json:"is_mall_store,omitempty" xml:"is_mall_store,omitempty"`
	ItemOrderId     *string                                                          `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	ReceiptBillKey  *string                                                          `json:"receipt_bill_key,omitempty" xml:"receipt_bill_key,omitempty"`
	AccountId       *string                                                          `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Certificate     *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate     `json:"certificate,omitempty" xml:"certificate,omitempty"`
	Goods           *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods           `json:"goods,omitempty" xml:"goods,omitempty"`
	OrderId         *string                                                          `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Amount          *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount          `json:"amount,omitempty" xml:"amount,omitempty"`
	Status          *int64                                                           `json:"status,omitempty" xml:"status,omitempty"`
	LedgerId        *string                                                          `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItem) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetLedgerTime(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.LedgerTime = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetOrderAttrribute(v *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.OrderAttrribute = v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetPoiId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.PoiId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.Id = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetSkuOrderId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.SkuOrderId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetIsMallStore(v bool) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.IsMallStore = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetItemOrderId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.ItemOrderId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetReceiptBillKey(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.ReceiptBillKey = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetAccountId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.AccountId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetCertificate(v *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.Certificate = v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetGoods(v *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.Goods = v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetOrderId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.OrderId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetAmount(v *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.Amount = v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetStatus(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.Status = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItem) SetLedgerId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItem {
	s.LedgerId = &v
	return s
}

type LedgerDetailedQueryResponseDataLedgerRecordsItemAmount struct {
	UserDelivMerchantSubsidy     *int64 `json:"user_deliv_merchant_subsidy,omitempty" xml:"user_deliv_merchant_subsidy,omitempty"`
	MallStoreGoods               *int64 `json:"mall_store_goods,omitempty" xml:"mall_store_goods,omitempty"`
	LedgerTotal                  *int64 `json:"ledger_total,omitempty" xml:"ledger_total,omitempty"`
	ActualInsured                *int64 `json:"actual_insured,omitempty" xml:"actual_insured,omitempty"`
	PlatformDutyTip              *int64 `json:"platform_duty_tip,omitempty" xml:"platform_duty_tip,omitempty"`
	TotalMerchantPlatformService *int64 `json:"total_merchant_platform_service,omitempty" xml:"total_merchant_platform_service,omitempty"`
	MerchantTicket               *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	BrokerCommission             *int64 `json:"broker_commission,omitempty" xml:"broker_commission,omitempty"`
	TalentCommission             *int64 `json:"talent_commission,omitempty" xml:"talent_commission,omitempty"`
	Goods                        *int64 `json:"goods,omitempty" xml:"goods,omitempty"`
	MerchantCancelFreight        *int64 `json:"merchant_cancel_freight,omitempty" xml:"merchant_cancel_freight,omitempty"`
	LedgerPlatformTicket         *int64 `json:"ledger_platform_ticket,omitempty" xml:"ledger_platform_ticket,omitempty"`
	TotalOperationAgent          *int64 `json:"total_operation_agent,omitempty" xml:"total_operation_agent,omitempty"`
	PlatformDutyFreight          *int64 `json:"platform_duty_freight,omitempty" xml:"platform_duty_freight,omitempty"`
	ClerkCommission              *int64 `json:"clerk_commission,omitempty" xml:"clerk_commission,omitempty"`
	UserDelivPlatformSubsidy     *int64 `json:"user_deliv_platform_subsidy,omitempty" xml:"user_deliv_platform_subsidy,omitempty"`
	MerchantTip                  *int64 `json:"merchant_tip,omitempty" xml:"merchant_tip,omitempty"`
	ThridPartyCommission         *int64 `json:"thrid_party_commission,omitempty" xml:"thrid_party_commission,omitempty"`
	CouponPay                    *int64 `json:"coupon_pay,omitempty" xml:"coupon_pay,omitempty"`
	MerchantFreight              *int64 `json:"merchant_freight,omitempty" xml:"merchant_freight,omitempty"`
	PlatformTicket               *int64 `json:"platform_ticket,omitempty" xml:"platform_ticket,omitempty"`
	TotalAgentMerchant           *int64 `json:"total_agent_merchant,omitempty" xml:"total_agent_merchant,omitempty"`
	UserDelivCouponPay           *int64 `json:"user_deliv_coupon_pay,omitempty" xml:"user_deliv_coupon_pay,omitempty"`
	InstitutionCommission        *int64 `json:"institution_commission,omitempty" xml:"institution_commission,omitempty"`
	Original                     *int64 `json:"original,omitempty" xml:"original,omitempty"`
	PayDiscount                  *int64 `json:"pay_discount,omitempty" xml:"pay_discount,omitempty"`
	CraftsMan                    *int64 `json:"crafts_man,omitempty" xml:"crafts_man,omitempty"`
	Pay                          *int64 `json:"pay,omitempty" xml:"pay,omitempty"`
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetUserDelivMerchantSubsidy(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.UserDelivMerchantSubsidy = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetMallStoreGoods(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.MallStoreGoods = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetLedgerTotal(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.LedgerTotal = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetActualInsured(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.ActualInsured = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetPlatformDutyTip(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.PlatformDutyTip = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetTotalMerchantPlatformService(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.TotalMerchantPlatformService = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetMerchantTicket(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.MerchantTicket = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetBrokerCommission(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.BrokerCommission = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetTalentCommission(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.TalentCommission = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetGoods(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.Goods = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetMerchantCancelFreight(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.MerchantCancelFreight = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetLedgerPlatformTicket(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.LedgerPlatformTicket = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetTotalOperationAgent(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.TotalOperationAgent = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetPlatformDutyFreight(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.PlatformDutyFreight = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetClerkCommission(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.ClerkCommission = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetUserDelivPlatformSubsidy(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.UserDelivPlatformSubsidy = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetMerchantTip(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.MerchantTip = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetThridPartyCommission(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.ThridPartyCommission = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetCouponPay(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.CouponPay = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetMerchantFreight(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.MerchantFreight = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetPlatformTicket(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.PlatformTicket = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetTotalAgentMerchant(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.TotalAgentMerchant = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetUserDelivCouponPay(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.UserDelivCouponPay = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetInstitutionCommission(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.InstitutionCommission = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetOriginal(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.Original = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetPayDiscount(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.PayDiscount = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetCraftsMan(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.CraftsMan = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount) SetPay(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemAmount {
	s.Pay = &v
	return s
}

type LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate) SetCertificateId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate {
	s.CertificateId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate) SetCode(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemCertificate {
	s.Code = &v
	return s
}

type LedgerDetailedQueryResponseDataLedgerRecordsItemGoods struct {
	GroupName     *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) SetGroupName(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods {
	s.GroupName = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) SetMarketPrice(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods {
	s.MarketPrice = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) SetSkuId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods {
	s.SkuId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods) SetSoldStartTime(v int64) *LedgerDetailedQueryResponseDataLedgerRecordsItemGoods {
	s.SoldStartTime = &v
	return s
}

type LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute struct {
	RoomId  *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Source  *string `json:"source,omitempty" xml:"source,omitempty"`
	VideoId *string `json:"video_id,omitempty" xml:"video_id,omitempty"`
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) SetRoomId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.RoomId = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) SetSource(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.Source = &v
	return s
}

func (s *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute) SetVideoId(v string) *LedgerDetailedQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.VideoId = &v
	return s
}

type LedgerDetailedQueryResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s LedgerDetailedQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s LedgerDetailedQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *LedgerDetailedQueryResponseExtra) SetErrorCode(v int32) *LedgerDetailedQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *LedgerDetailedQueryResponseExtra) SetLogid(v string) *LedgerDetailedQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *LedgerDetailedQueryResponseExtra) SetNow(v int64) *LedgerDetailedQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *LedgerDetailedQueryResponseExtra) SetSubDescription(v string) *LedgerDetailedQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *LedgerDetailedQueryResponseExtra) SetSubErrorCode(v int32) *LedgerDetailedQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *LedgerDetailedQueryResponseExtra) SetDescription(v string) *LedgerDetailedQueryResponseExtra {
	s.Description = &v
	return s
}

type LedgerQueryByOrderRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderIds    []*string          `json:"order_ids,omitempty" xml:"order_ids,omitempty" require:"true" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s LedgerQueryByOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderRequest) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderRequest) SetHeader(v map[string]*string) *LedgerQueryByOrderRequest {
	s.Header = v
	return s
}

func (s *LedgerQueryByOrderRequest) SetAccessToken(v string) *LedgerQueryByOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *LedgerQueryByOrderRequest) SetOrderIds(v []*string) *LedgerQueryByOrderRequest {
	s.OrderIds = v
	return s
}

func (s *LedgerQueryByOrderRequest) SetAccountId(v string) *LedgerQueryByOrderRequest {
	s.AccountId = &v
	return s
}

type LedgerQueryByOrderResponse struct {
	Extra *LedgerQueryByOrderResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *LedgerQueryByOrderResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s LedgerQueryByOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponse) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponse) SetExtra(v *LedgerQueryByOrderResponseExtra) *LedgerQueryByOrderResponse {
	s.Extra = v
	return s
}

func (s *LedgerQueryByOrderResponse) SetData(v *LedgerQueryByOrderResponseData) *LedgerQueryByOrderResponse {
	s.Data = v
	return s
}

type LedgerQueryByOrderResponseData struct {
	LedgerSecondRecords map[string][]*LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem `json:"ledger_second_records,omitempty" xml:"ledger_second_records,omitempty"`
	OrderAmounts        map[string]*LedgerQueryByOrderResponseDataOrderAmountsValue              `json:"order_amounts,omitempty" xml:"order_amounts,omitempty"`
	GwErrorCode         *int32                                                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription       *string                                                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	LedgerRecords       map[string][]*LedgerQueryByOrderResponseDataLedgerRecordsValueItem       `json:"ledger_records,omitempty" xml:"ledger_records,omitempty"`
}

func (s LedgerQueryByOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseData) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseData) SetLedgerSecondRecords(v map[string][]*LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) *LedgerQueryByOrderResponseData {
	s.LedgerSecondRecords = v
	return s
}

func (s *LedgerQueryByOrderResponseData) SetOrderAmounts(v map[string]*LedgerQueryByOrderResponseDataOrderAmountsValue) *LedgerQueryByOrderResponseData {
	s.OrderAmounts = v
	return s
}

func (s *LedgerQueryByOrderResponseData) SetGwErrorCode(v int32) *LedgerQueryByOrderResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *LedgerQueryByOrderResponseData) SetGwDescription(v string) *LedgerQueryByOrderResponseData {
	s.GwDescription = &v
	return s
}

func (s *LedgerQueryByOrderResponseData) SetLedgerRecords(v map[string][]*LedgerQueryByOrderResponseDataLedgerRecordsValueItem) *LedgerQueryByOrderResponseData {
	s.LedgerRecords = v
	return s
}

type LedgerQueryByOrderResponseDataLedgerRecordsValueItem struct {
	Amount    *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	OrderId   *string                                                     `json:"order_id,omitempty" xml:"order_id,omitempty"`
	LedgerId  *string                                                     `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	Goods     *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods  `json:"goods,omitempty" xml:"goods,omitempty"`
	AccountId *string                                                     `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Status    *int64                                                      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItem) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItem) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetAmount(v *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Amount = v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetOrderId(v string) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.OrderId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetLedgerId(v string) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.LedgerId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetGoods(v *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Goods = v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetAccountId(v string) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.AccountId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItem) SetStatus(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItem {
	s.Status = &v
	return s
}

type LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount struct {
	ActualInsured  *int64 `json:"actual_insured,omitempty" xml:"actual_insured,omitempty"`
	CouponPay      *int64 `json:"coupon_pay,omitempty" xml:"coupon_pay,omitempty"`
	Goods          *int64 `json:"goods,omitempty" xml:"goods,omitempty"`
	MerchantTicket *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	Pay            *int64 `json:"pay,omitempty" xml:"pay,omitempty"`
	Original       *int64 `json:"original,omitempty" xml:"original,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetActualInsured(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.ActualInsured = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetCouponPay(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.CouponPay = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetGoods(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Goods = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetMerchantTicket(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.MerchantTicket = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetPay(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Pay = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount) SetOriginal(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemAmount {
	s.Original = &v
	return s
}

type LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods struct {
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetSkuId(v string) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.SkuId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetSoldStartTime(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.SoldStartTime = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods) SetMarketPrice(v int64) *LedgerQueryByOrderResponseDataLedgerRecordsValueItemGoods {
	s.MarketPrice = &v
	return s
}

type LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem struct {
	Status    *int64                                                            `json:"status,omitempty" xml:"status,omitempty"`
	Amount    *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	AccountId *string                                                           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Goods     *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods  `json:"goods,omitempty" xml:"goods,omitempty"`
	LedgerId  *string                                                           `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	OrderId   *string                                                           `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetStatus(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.Status = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetAmount(v *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.Amount = v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetAccountId(v string) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.AccountId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetGoods(v *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.Goods = v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetLedgerId(v string) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.LedgerId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem) SetOrderId(v string) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItem {
	s.OrderId = &v
	return s
}

type LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount struct {
	MerchantTip              *int64 `json:"merchant_tip,omitempty" xml:"merchant_tip,omitempty"`
	PlatformDutyFreight      *int64 `json:"platform_duty_freight,omitempty" xml:"platform_duty_freight,omitempty"`
	UserDelivCouponPay       *int64 `json:"user_deliv_coupon_pay,omitempty" xml:"user_deliv_coupon_pay,omitempty"`
	UserDelivMerchantSubsidy *int64 `json:"user_deliv_merchant_subsidy,omitempty" xml:"user_deliv_merchant_subsidy,omitempty"`
	UserDelivPlatformSubsidy *int64 `json:"user_deliv_platform_subsidy,omitempty" xml:"user_deliv_platform_subsidy,omitempty"`
	MerchantCancelFreight    *int64 `json:"merchant_cancel_freight,omitempty" xml:"merchant_cancel_freight,omitempty"`
	MerchantFreight          *int64 `json:"merchant_freight,omitempty" xml:"merchant_freight,omitempty"`
	PlatformDutyTip          *int64 `json:"platform_duty_tip,omitempty" xml:"platform_duty_tip,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantTip(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantTip = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetPlatformDutyFreight(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.PlatformDutyFreight = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivCouponPay(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivCouponPay = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivMerchantSubsidy(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivMerchantSubsidy = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetUserDelivPlatformSubsidy(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.UserDelivPlatformSubsidy = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantCancelFreight(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantCancelFreight = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetMerchantFreight(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.MerchantFreight = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount) SetPlatformDutyTip(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemAmount {
	s.PlatformDutyTip = &v
	return s
}

type LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods struct {
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) SetMarketPrice(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods {
	s.MarketPrice = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) SetSkuId(v string) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods {
	s.SkuId = &v
	return s
}

func (s *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods) SetSoldStartTime(v int64) *LedgerQueryByOrderResponseDataLedgerSecondRecordsValueItemGoods {
	s.SoldStartTime = &v
	return s
}

type LedgerQueryByOrderResponseDataOrderAmountsValue struct {
	MerchantIncome *int64 `json:"merchant_income,omitempty" xml:"merchant_income,omitempty"`
}

func (s LedgerQueryByOrderResponseDataOrderAmountsValue) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseDataOrderAmountsValue) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseDataOrderAmountsValue) SetMerchantIncome(v int64) *LedgerQueryByOrderResponseDataOrderAmountsValue {
	s.MerchantIncome = &v
	return s
}

type LedgerQueryByOrderResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s LedgerQueryByOrderResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryByOrderResponseExtra) GoString() string {
	return s.String()
}

func (s *LedgerQueryByOrderResponseExtra) SetDescription(v string) *LedgerQueryByOrderResponseExtra {
	s.Description = &v
	return s
}

func (s *LedgerQueryByOrderResponseExtra) SetErrorCode(v int32) *LedgerQueryByOrderResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *LedgerQueryByOrderResponseExtra) SetLogid(v string) *LedgerQueryByOrderResponseExtra {
	s.Logid = &v
	return s
}

func (s *LedgerQueryByOrderResponseExtra) SetNow(v int64) *LedgerQueryByOrderResponseExtra {
	s.Now = &v
	return s
}

func (s *LedgerQueryByOrderResponseExtra) SetSubDescription(v string) *LedgerQueryByOrderResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *LedgerQueryByOrderResponseExtra) SetSubErrorCode(v int32) *LedgerQueryByOrderResponseExtra {
	s.SubErrorCode = &v
	return s
}

type LedgerQueryRequest struct {
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Size        *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
}

func (s LedgerQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryRequest) GoString() string {
	return s.String()
}

func (s *LedgerQueryRequest) SetCursor(v string) *LedgerQueryRequest {
	s.Cursor = &v
	return s
}

func (s *LedgerQueryRequest) SetSize(v int64) *LedgerQueryRequest {
	s.Size = &v
	return s
}

func (s *LedgerQueryRequest) SetHeader(v map[string]*string) *LedgerQueryRequest {
	s.Header = v
	return s
}

func (s *LedgerQueryRequest) SetAccessToken(v string) *LedgerQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *LedgerQueryRequest) SetAccountId(v string) *LedgerQueryRequest {
	s.AccountId = &v
	return s
}

func (s *LedgerQueryRequest) SetBillDate(v string) *LedgerQueryRequest {
	s.BillDate = &v
	return s
}

type LedgerQueryResponse struct {
	Data  *LedgerQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *LedgerQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s LedgerQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponse) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponse) SetData(v *LedgerQueryResponseData) *LedgerQueryResponse {
	s.Data = v
	return s
}

func (s *LedgerQueryResponse) SetExtra(v *LedgerQueryResponseExtra) *LedgerQueryResponse {
	s.Extra = v
	return s
}

type LedgerQueryResponseData struct {
	Cursor        *string                                     `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                       `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	LedgerRecords []*LedgerQueryResponseDataLedgerRecordsItem `json:"ledger_records,omitempty" xml:"ledger_records,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s LedgerQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseData) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseData) SetCursor(v string) *LedgerQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *LedgerQueryResponseData) SetHasMore(v bool) *LedgerQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *LedgerQueryResponseData) SetLedgerRecords(v []*LedgerQueryResponseDataLedgerRecordsItem) *LedgerQueryResponseData {
	s.LedgerRecords = v
	return s
}

func (s *LedgerQueryResponseData) SetGwErrorCode(v int32) *LedgerQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *LedgerQueryResponseData) SetGwDescription(v string) *LedgerQueryResponseData {
	s.GwDescription = &v
	return s
}

type LedgerQueryResponseDataLedgerRecordsItem struct {
	OrderId         *string                                                  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Certificate     *LedgerQueryResponseDataLedgerRecordsItemCertificate     `json:"certificate,omitempty" xml:"certificate,omitempty"`
	SkuOrderId      *string                                                  `json:"sku_order_id,omitempty" xml:"sku_order_id,omitempty"`
	Amount          *LedgerQueryResponseDataLedgerRecordsItemAmount          `json:"amount,omitempty" xml:"amount,omitempty"`
	OrderAttrribute *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute `json:"order_attrribute,omitempty" xml:"order_attrribute,omitempty"`
	PoiId           *string                                                  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ReceiptBillKey  *string                                                  `json:"receipt_bill_key,omitempty" xml:"receipt_bill_key,omitempty"`
	ItemOrderId     *string                                                  `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	Id              *string                                                  `json:"id,omitempty" xml:"id,omitempty"`
	AccountId       *string                                                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	LedgerTime      *int64                                                   `json:"ledger_time,omitempty" xml:"ledger_time,omitempty"`
	Status          *int64                                                   `json:"status,omitempty" xml:"status,omitempty"`
	Goods           *LedgerQueryResponseDataLedgerRecordsItemGoods           `json:"goods,omitempty" xml:"goods,omitempty"`
	LedgerId        *string                                                  `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
}

func (s LedgerQueryResponseDataLedgerRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseDataLedgerRecordsItem) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetOrderId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.OrderId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetCertificate(v *LedgerQueryResponseDataLedgerRecordsItemCertificate) *LedgerQueryResponseDataLedgerRecordsItem {
	s.Certificate = v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetSkuOrderId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.SkuOrderId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetAmount(v *LedgerQueryResponseDataLedgerRecordsItemAmount) *LedgerQueryResponseDataLedgerRecordsItem {
	s.Amount = v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetOrderAttrribute(v *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) *LedgerQueryResponseDataLedgerRecordsItem {
	s.OrderAttrribute = v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetPoiId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.PoiId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetReceiptBillKey(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.ReceiptBillKey = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetItemOrderId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.ItemOrderId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.Id = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetAccountId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.AccountId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetLedgerTime(v int64) *LedgerQueryResponseDataLedgerRecordsItem {
	s.LedgerTime = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetStatus(v int64) *LedgerQueryResponseDataLedgerRecordsItem {
	s.Status = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetGoods(v *LedgerQueryResponseDataLedgerRecordsItemGoods) *LedgerQueryResponseDataLedgerRecordsItem {
	s.Goods = v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItem) SetLedgerId(v string) *LedgerQueryResponseDataLedgerRecordsItem {
	s.LedgerId = &v
	return s
}

type LedgerQueryResponseDataLedgerRecordsItemAmount struct {
	CouponPay      *int64 `json:"coupon_pay,omitempty" xml:"coupon_pay,omitempty"`
	Goods          *int64 `json:"goods,omitempty" xml:"goods,omitempty"`
	MerchantTicket *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	ActualInsured  *int64 `json:"actual_insured,omitempty" xml:"actual_insured,omitempty"`
	Original       *int64 `json:"original,omitempty" xml:"original,omitempty"`
	Pay            *int64 `json:"pay,omitempty" xml:"pay,omitempty"`
}

func (s LedgerQueryResponseDataLedgerRecordsItemAmount) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseDataLedgerRecordsItemAmount) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetCouponPay(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.CouponPay = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetGoods(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.Goods = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetMerchantTicket(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.MerchantTicket = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetActualInsured(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.ActualInsured = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetOriginal(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.Original = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemAmount) SetPay(v int64) *LedgerQueryResponseDataLedgerRecordsItemAmount {
	s.Pay = &v
	return s
}

type LedgerQueryResponseDataLedgerRecordsItemCertificate struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s LedgerQueryResponseDataLedgerRecordsItemCertificate) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseDataLedgerRecordsItemCertificate) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseDataLedgerRecordsItemCertificate) SetCertificateId(v string) *LedgerQueryResponseDataLedgerRecordsItemCertificate {
	s.CertificateId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemCertificate) SetCode(v string) *LedgerQueryResponseDataLedgerRecordsItemCertificate {
	s.Code = &v
	return s
}

type LedgerQueryResponseDataLedgerRecordsItemGoods struct {
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	GroupName     *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s LedgerQueryResponseDataLedgerRecordsItemGoods) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseDataLedgerRecordsItemGoods) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseDataLedgerRecordsItemGoods) SetSkuId(v string) *LedgerQueryResponseDataLedgerRecordsItemGoods {
	s.SkuId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemGoods) SetSoldStartTime(v int64) *LedgerQueryResponseDataLedgerRecordsItemGoods {
	s.SoldStartTime = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemGoods) SetGroupName(v string) *LedgerQueryResponseDataLedgerRecordsItemGoods {
	s.GroupName = &v
	return s
}

type LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute struct {
	Source  *string `json:"source,omitempty" xml:"source,omitempty"`
	VideoId *string `json:"video_id,omitempty" xml:"video_id,omitempty"`
	RoomId  *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

func (s LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) SetSource(v string) *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.Source = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) SetVideoId(v string) *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.VideoId = &v
	return s
}

func (s *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute) SetRoomId(v string) *LedgerQueryResponseDataLedgerRecordsItemOrderAttrribute {
	s.RoomId = &v
	return s
}

type LedgerQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s LedgerQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s LedgerQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *LedgerQueryResponseExtra) SetSubErrorCode(v int32) *LedgerQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *LedgerQueryResponseExtra) SetDescription(v string) *LedgerQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *LedgerQueryResponseExtra) SetErrorCode(v int32) *LedgerQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *LedgerQueryResponseExtra) SetLogid(v string) *LedgerQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *LedgerQueryResponseExtra) SetNow(v int64) *LedgerQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *LedgerQueryResponseExtra) SetSubDescription(v string) *LedgerQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type LinkGetRequest struct {
	AccessToken        *string                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExternalMerchantId *string                     `json:"external_merchant_id,omitempty" xml:"external_merchant_id,omitempty"`
	PhoneNo            *string                     `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
	EntranceCert       *LinkGetRequestEntranceCert `json:"entrance_cert,omitempty" xml:"entrance_cert,omitempty"`
	EntrancePoi        *LinkGetRequestEntrancePoi  `json:"entrance_poi,omitempty" xml:"entrance_poi,omitempty"`
	Header             map[string]*string          `json:"header,omitempty" xml:"header,omitempty"`
}

func (s LinkGetRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkGetRequest) GoString() string {
	return s.String()
}

func (s *LinkGetRequest) SetAccessToken(v string) *LinkGetRequest {
	s.AccessToken = &v
	return s
}

func (s *LinkGetRequest) SetExternalMerchantId(v string) *LinkGetRequest {
	s.ExternalMerchantId = &v
	return s
}

func (s *LinkGetRequest) SetPhoneNo(v string) *LinkGetRequest {
	s.PhoneNo = &v
	return s
}

func (s *LinkGetRequest) SetEntranceCert(v *LinkGetRequestEntranceCert) *LinkGetRequest {
	s.EntranceCert = v
	return s
}

func (s *LinkGetRequest) SetEntrancePoi(v *LinkGetRequestEntrancePoi) *LinkGetRequest {
	s.EntrancePoi = v
	return s
}

func (s *LinkGetRequest) SetHeader(v map[string]*string) *LinkGetRequest {
	s.Header = v
	return s
}

type LinkGetRequestEntranceCert struct {
	EntranceSubjectCert      *LinkGetRequestEntranceCertEntranceSubjectCert            `json:"entrance_subject_cert,omitempty" xml:"entrance_subject_cert,omitempty"`
	EntranceIndustryCertList []*LinkGetRequestEntranceCertEntranceIndustryCertListItem `json:"entrance_industry_cert_list,omitempty" xml:"entrance_industry_cert_list,omitempty" type:"Repeated"`
}

func (s LinkGetRequestEntranceCert) String() string {
	return tea.Prettify(s)
}

func (s LinkGetRequestEntranceCert) GoString() string {
	return s.String()
}

func (s *LinkGetRequestEntranceCert) SetEntranceSubjectCert(v *LinkGetRequestEntranceCertEntranceSubjectCert) *LinkGetRequestEntranceCert {
	s.EntranceSubjectCert = v
	return s
}

func (s *LinkGetRequestEntranceCert) SetEntranceIndustryCertList(v []*LinkGetRequestEntranceCertEntranceIndustryCertListItem) *LinkGetRequestEntranceCert {
	s.EntranceIndustryCertList = v
	return s
}

type LinkGetRequestEntranceCertEntranceIndustryCertListItem struct {
	EffectiveDate *string `json:"effective_date,omitempty" xml:"effective_date,omitempty"`
	QualType      *string `json:"qual_type,omitempty" xml:"qual_type,omitempty"`
	Uri           *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s LinkGetRequestEntranceCertEntranceIndustryCertListItem) String() string {
	return tea.Prettify(s)
}

func (s LinkGetRequestEntranceCertEntranceIndustryCertListItem) GoString() string {
	return s.String()
}

func (s *LinkGetRequestEntranceCertEntranceIndustryCertListItem) SetEffectiveDate(v string) *LinkGetRequestEntranceCertEntranceIndustryCertListItem {
	s.EffectiveDate = &v
	return s
}

func (s *LinkGetRequestEntranceCertEntranceIndustryCertListItem) SetQualType(v string) *LinkGetRequestEntranceCertEntranceIndustryCertListItem {
	s.QualType = &v
	return s
}

func (s *LinkGetRequestEntranceCertEntranceIndustryCertListItem) SetUri(v string) *LinkGetRequestEntranceCertEntranceIndustryCertListItem {
	s.Uri = &v
	return s
}

type LinkGetRequestEntranceCertEntranceSubjectCert struct {
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s LinkGetRequestEntranceCertEntranceSubjectCert) String() string {
	return tea.Prettify(s)
}

func (s LinkGetRequestEntranceCertEntranceSubjectCert) GoString() string {
	return s.String()
}

func (s *LinkGetRequestEntranceCertEntranceSubjectCert) SetUri(v string) *LinkGetRequestEntranceCertEntranceSubjectCert {
	s.Uri = &v
	return s
}

type LinkGetRequestEntrancePoi struct {
	PoiId *int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s LinkGetRequestEntrancePoi) String() string {
	return tea.Prettify(s)
}

func (s LinkGetRequestEntrancePoi) GoString() string {
	return s.String()
}

func (s *LinkGetRequestEntrancePoi) SetPoiId(v int64) *LinkGetRequestEntrancePoi {
	s.PoiId = &v
	return s
}

type LinkGetResponse struct {
	Data  *LinkGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *LinkGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s LinkGetResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkGetResponse) GoString() string {
	return s.String()
}

func (s *LinkGetResponse) SetData(v *LinkGetResponseData) *LinkGetResponse {
	s.Data = v
	return s
}

func (s *LinkGetResponse) SetExtra(v *LinkGetResponseExtra) *LinkGetResponse {
	s.Extra = v
	return s
}

type LinkGetResponseData struct {
	GwErrorCode    *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription  *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	EntranceH5Link *string `json:"entrance_h5_link,omitempty" xml:"entrance_h5_link,omitempty"`
}

func (s LinkGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s LinkGetResponseData) GoString() string {
	return s.String()
}

func (s *LinkGetResponseData) SetGwErrorCode(v int32) *LinkGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *LinkGetResponseData) SetGwDescription(v string) *LinkGetResponseData {
	s.GwDescription = &v
	return s
}

func (s *LinkGetResponseData) SetEntranceH5Link(v string) *LinkGetResponseData {
	s.EntranceH5Link = &v
	return s
}

type LinkGetResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s LinkGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s LinkGetResponseExtra) GoString() string {
	return s.String()
}

func (s *LinkGetResponseExtra) SetSubDescription(v string) *LinkGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *LinkGetResponseExtra) SetSubErrorCode(v int32) *LinkGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *LinkGetResponseExtra) SetDescription(v string) *LinkGetResponseExtra {
	s.Description = &v
	return s
}

func (s *LinkGetResponseExtra) SetErrorCode(v int32) *LinkGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *LinkGetResponseExtra) SetLogid(v string) *LinkGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *LinkGetResponseExtra) SetNow(v int64) *LinkGetResponseExtra {
	s.Now = &v
	return s
}

type LinkmicQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	RoomId      *string            `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s LinkmicQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkmicQueryRequest) GoString() string {
	return s.String()
}

func (s *LinkmicQueryRequest) SetAccessToken(v string) *LinkmicQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *LinkmicQueryRequest) SetAppId(v string) *LinkmicQueryRequest {
	s.AppId = &v
	return s
}

func (s *LinkmicQueryRequest) SetRoomId(v string) *LinkmicQueryRequest {
	s.RoomId = &v
	return s
}

func (s *LinkmicQueryRequest) SetHeader(v map[string]*string) *LinkmicQueryRequest {
	s.Header = v
	return s
}

type LinkmicQueryResponse struct {
	UserList []*LinkmicQueryResponseUserListItem `json:"user_list,omitempty" xml:"user_list,omitempty" type:"Repeated"`
	BaseInfo *LinkmicQueryResponseBaseInfo       `json:"base_info,omitempty" xml:"base_info,omitempty"`
}

func (s LinkmicQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkmicQueryResponse) GoString() string {
	return s.String()
}

func (s *LinkmicQueryResponse) SetUserList(v []*LinkmicQueryResponseUserListItem) *LinkmicQueryResponse {
	s.UserList = v
	return s
}

func (s *LinkmicQueryResponse) SetBaseInfo(v *LinkmicQueryResponseBaseInfo) *LinkmicQueryResponse {
	s.BaseInfo = v
	return s
}

type LinkmicQueryResponseBaseInfo struct {
	TotalCount *int32  `json:"total_count,omitempty" xml:"total_count,omitempty"`
	FreeCount  *int32  `json:"free_count,omitempty" xml:"free_count,omitempty"`
	LinkId     *string `json:"link_id,omitempty" xml:"link_id,omitempty"`
}

func (s LinkmicQueryResponseBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s LinkmicQueryResponseBaseInfo) GoString() string {
	return s.String()
}

func (s *LinkmicQueryResponseBaseInfo) SetTotalCount(v int32) *LinkmicQueryResponseBaseInfo {
	s.TotalCount = &v
	return s
}

func (s *LinkmicQueryResponseBaseInfo) SetFreeCount(v int32) *LinkmicQueryResponseBaseInfo {
	s.FreeCount = &v
	return s
}

func (s *LinkmicQueryResponseBaseInfo) SetLinkId(v string) *LinkmicQueryResponseBaseInfo {
	s.LinkId = &v
	return s
}

type LinkmicQueryResponseUserListItem struct {
	OpenId            *string                                  `json:"open_id,omitempty" xml:"open_id,omitempty"`
	CameraState       *int32                                   `json:"camera_state,omitempty" xml:"camera_state,omitempty"`
	AppInfo           *LinkmicQueryResponseUserListItemAppInfo `json:"app_info,omitempty" xml:"app_info,omitempty"`
	DisableMicrophone *int32                                   `json:"disable_microphone,omitempty" xml:"disable_microphone,omitempty"`
	NickName          *string                                  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	SecAvatarUrl      *string                                  `json:"sec_avatar_url,omitempty" xml:"sec_avatar_url,omitempty"`
	LinkState         *int32                                   `json:"link_state,omitempty" xml:"link_state,omitempty"`
	LinkPosition      *int32                                   `json:"link_position,omitempty" xml:"link_position,omitempty"`
	DisableCamera     *int32                                   `json:"disable_camera,omitempty" xml:"disable_camera,omitempty"`
	MicrophoneState   *int32                                   `json:"microphone_state,omitempty" xml:"microphone_state,omitempty"`
	AvatarUrl         *string                                  `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	SecNickName       *string                                  `json:"sec_nick_name,omitempty" xml:"sec_nick_name,omitempty"`
}

func (s LinkmicQueryResponseUserListItem) String() string {
	return tea.Prettify(s)
}

func (s LinkmicQueryResponseUserListItem) GoString() string {
	return s.String()
}

func (s *LinkmicQueryResponseUserListItem) SetOpenId(v string) *LinkmicQueryResponseUserListItem {
	s.OpenId = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetCameraState(v int32) *LinkmicQueryResponseUserListItem {
	s.CameraState = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetAppInfo(v *LinkmicQueryResponseUserListItemAppInfo) *LinkmicQueryResponseUserListItem {
	s.AppInfo = v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetDisableMicrophone(v int32) *LinkmicQueryResponseUserListItem {
	s.DisableMicrophone = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetNickName(v string) *LinkmicQueryResponseUserListItem {
	s.NickName = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetSecAvatarUrl(v string) *LinkmicQueryResponseUserListItem {
	s.SecAvatarUrl = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetLinkState(v int32) *LinkmicQueryResponseUserListItem {
	s.LinkState = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetLinkPosition(v int32) *LinkmicQueryResponseUserListItem {
	s.LinkPosition = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetDisableCamera(v int32) *LinkmicQueryResponseUserListItem {
	s.DisableCamera = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetMicrophoneState(v int32) *LinkmicQueryResponseUserListItem {
	s.MicrophoneState = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetAvatarUrl(v string) *LinkmicQueryResponseUserListItem {
	s.AvatarUrl = &v
	return s
}

func (s *LinkmicQueryResponseUserListItem) SetSecNickName(v string) *LinkmicQueryResponseUserListItem {
	s.SecNickName = &v
	return s
}

type LinkmicQueryResponseUserListItemAppInfo struct {
	HostAppStartAppAvailable *bool `json:"host_app_start_app_available,omitempty" xml:"host_app_start_app_available,omitempty"`
}

func (s LinkmicQueryResponseUserListItemAppInfo) String() string {
	return tea.Prettify(s)
}

func (s LinkmicQueryResponseUserListItemAppInfo) GoString() string {
	return s.String()
}

func (s *LinkmicQueryResponseUserListItemAppInfo) SetHostAppStartAppAvailable(v bool) *LinkmicQueryResponseUserListItemAppInfo {
	s.HostAppStartAppAvailable = &v
	return s
}

type ListPlanBySpuidRequest struct {
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	SpuId       *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
	SpuIdType   *int               `json:"spu_id_type,omitempty" xml:"spu_id_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageNo      *int32             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
}

func (s ListPlanBySpuidRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlanBySpuidRequest) GoString() string {
	return s.String()
}

func (s *ListPlanBySpuidRequest) SetPageSize(v int32) *ListPlanBySpuidRequest {
	s.PageSize = &v
	return s
}

func (s *ListPlanBySpuidRequest) SetSpuId(v int64) *ListPlanBySpuidRequest {
	s.SpuId = &v
	return s
}

func (s *ListPlanBySpuidRequest) SetSpuIdType(v int) *ListPlanBySpuidRequest {
	s.SpuIdType = &v
	return s
}

func (s *ListPlanBySpuidRequest) SetHeader(v map[string]*string) *ListPlanBySpuidRequest {
	s.Header = v
	return s
}

func (s *ListPlanBySpuidRequest) SetAccessToken(v string) *ListPlanBySpuidRequest {
	s.AccessToken = &v
	return s
}

func (s *ListPlanBySpuidRequest) SetPageNo(v int32) *ListPlanBySpuidRequest {
	s.PageNo = &v
	return s
}

type ListPlanBySpuidResponse struct {
	Data   *ListPlanBySpuidResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ListPlanBySpuidResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPlanBySpuidResponse) GoString() string {
	return s.String()
}

func (s *ListPlanBySpuidResponse) SetData(v *ListPlanBySpuidResponseData) *ListPlanBySpuidResponse {
	s.Data = v
	return s
}

func (s *ListPlanBySpuidResponse) SetErrMsg(v string) *ListPlanBySpuidResponse {
	s.ErrMsg = &v
	return s
}

func (s *ListPlanBySpuidResponse) SetErrNo(v int32) *ListPlanBySpuidResponse {
	s.ErrNo = &v
	return s
}

func (s *ListPlanBySpuidResponse) SetLogId(v string) *ListPlanBySpuidResponse {
	s.LogId = &v
	return s
}

type ListPlanBySpuidResponseData struct {
	Data      []*ListPlanBySpuidResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	PageCount *int64                                 `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
	Total     *int64                                 `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s ListPlanBySpuidResponseData) String() string {
	return tea.Prettify(s)
}

func (s ListPlanBySpuidResponseData) GoString() string {
	return s.String()
}

func (s *ListPlanBySpuidResponseData) SetData(v []*ListPlanBySpuidResponseDataDataItem) *ListPlanBySpuidResponseData {
	s.Data = v
	return s
}

func (s *ListPlanBySpuidResponseData) SetPageCount(v int64) *ListPlanBySpuidResponseData {
	s.PageCount = &v
	return s
}

func (s *ListPlanBySpuidResponseData) SetTotal(v int64) *ListPlanBySpuidResponseData {
	s.Total = &v
	return s
}

type ListPlanBySpuidResponseDataDataItem struct {
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	UpdateTime     *string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	UpdateTimeUnix *int64  `json:"update_time_unix,omitempty" xml:"update_time_unix,omitempty"`
	CommissionRate *int64  `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	ContentType    *int32  `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	CreateTime     *string `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	CreateTimeUnix *int64  `json:"create_time_unix,omitempty" xml:"create_time_unix,omitempty"`
	PlanId         *int64  `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s ListPlanBySpuidResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s ListPlanBySpuidResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *ListPlanBySpuidResponseDataDataItem) SetStatus(v int32) *ListPlanBySpuidResponseDataDataItem {
	s.Status = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetUpdateTime(v string) *ListPlanBySpuidResponseDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetUpdateTimeUnix(v int64) *ListPlanBySpuidResponseDataDataItem {
	s.UpdateTimeUnix = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetCommissionRate(v int64) *ListPlanBySpuidResponseDataDataItem {
	s.CommissionRate = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetContentType(v int32) *ListPlanBySpuidResponseDataDataItem {
	s.ContentType = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetCreateTime(v string) *ListPlanBySpuidResponseDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetCreateTimeUnix(v int64) *ListPlanBySpuidResponseDataDataItem {
	s.CreateTimeUnix = &v
	return s
}

func (s *ListPlanBySpuidResponseDataDataItem) SetPlanId(v int64) *ListPlanBySpuidResponseDataDataItem {
	s.PlanId = &v
	return s
}

type LiveDataAckRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AckType     *int64             `json:"ack_type,omitempty" xml:"ack_type,omitempty" require:"true"`
	Data        *string            `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	RoomId      *string            `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s LiveDataAckRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveDataAckRequest) GoString() string {
	return s.String()
}

func (s *LiveDataAckRequest) SetAccessToken(v string) *LiveDataAckRequest {
	s.AccessToken = &v
	return s
}

func (s *LiveDataAckRequest) SetAckType(v int64) *LiveDataAckRequest {
	s.AckType = &v
	return s
}

func (s *LiveDataAckRequest) SetData(v string) *LiveDataAckRequest {
	s.Data = &v
	return s
}

func (s *LiveDataAckRequest) SetRoomId(v string) *LiveDataAckRequest {
	s.RoomId = &v
	return s
}

func (s *LiveDataAckRequest) SetAppId(v string) *LiveDataAckRequest {
	s.AppId = &v
	return s
}

func (s *LiveDataAckRequest) SetHeader(v map[string]*string) *LiveDataAckRequest {
	s.Header = v
	return s
}

type LiveDataAckResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	Logid  *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s LiveDataAckResponse) String() string {
	return tea.Prettify(s)
}

func (s LiveDataAckResponse) GoString() string {
	return s.String()
}

func (s *LiveDataAckResponse) SetErrNo(v int32) *LiveDataAckResponse {
	s.ErrNo = &v
	return s
}

func (s *LiveDataAckResponse) SetErrMsg(v string) *LiveDataAckResponse {
	s.ErrMsg = &v
	return s
}

func (s *LiveDataAckResponse) SetLogid(v string) *LiveDataAckResponse {
	s.Logid = &v
	return s
}

type MatchTaskQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskId      *int64             `json:"task_id,omitempty" xml:"task_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s MatchTaskQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MatchTaskQueryRequest) GoString() string {
	return s.String()
}

func (s *MatchTaskQueryRequest) SetAccessToken(v string) *MatchTaskQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *MatchTaskQueryRequest) SetTaskId(v int64) *MatchTaskQueryRequest {
	s.TaskId = &v
	return s
}

func (s *MatchTaskQueryRequest) SetHeader(v map[string]*string) *MatchTaskQueryRequest {
	s.Header = v
	return s
}

type MatchTaskQueryResponse struct {
	Data  *MatchTaskQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *MatchTaskQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s MatchTaskQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MatchTaskQueryResponse) GoString() string {
	return s.String()
}

func (s *MatchTaskQueryResponse) SetData(v *MatchTaskQueryResponseData) *MatchTaskQueryResponse {
	s.Data = v
	return s
}

func (s *MatchTaskQueryResponse) SetExtra(v *MatchTaskQueryResponseExtra) *MatchTaskQueryResponse {
	s.Extra = v
	return s
}

type MatchTaskQueryResponseData struct {
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Results       []*MatchTaskQueryResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s MatchTaskQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s MatchTaskQueryResponseData) GoString() string {
	return s.String()
}

func (s *MatchTaskQueryResponseData) SetGwErrorCode(v int32) *MatchTaskQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *MatchTaskQueryResponseData) SetGwDescription(v string) *MatchTaskQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *MatchTaskQueryResponseData) SetResults(v []*MatchTaskQueryResponseDataResultsItem) *MatchTaskQueryResponseData {
	s.Results = v
	return s
}

type MatchTaskQueryResponseDataResultsItem struct {
	ExtId        *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	MatchMessage *string `json:"match_message,omitempty" xml:"match_message,omitempty"`
	MatchPoiId   *string `json:"match_poi_id,omitempty" xml:"match_poi_id,omitempty"`
	MatchResult  *int64  `json:"match_result,omitempty" xml:"match_result,omitempty"`
	MatchStatus  *int64  `json:"match_status,omitempty" xml:"match_status,omitempty"`
}

func (s MatchTaskQueryResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s MatchTaskQueryResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *MatchTaskQueryResponseDataResultsItem) SetExtId(v string) *MatchTaskQueryResponseDataResultsItem {
	s.ExtId = &v
	return s
}

func (s *MatchTaskQueryResponseDataResultsItem) SetMatchMessage(v string) *MatchTaskQueryResponseDataResultsItem {
	s.MatchMessage = &v
	return s
}

func (s *MatchTaskQueryResponseDataResultsItem) SetMatchPoiId(v string) *MatchTaskQueryResponseDataResultsItem {
	s.MatchPoiId = &v
	return s
}

func (s *MatchTaskQueryResponseDataResultsItem) SetMatchResult(v int64) *MatchTaskQueryResponseDataResultsItem {
	s.MatchResult = &v
	return s
}

func (s *MatchTaskQueryResponseDataResultsItem) SetMatchStatus(v int64) *MatchTaskQueryResponseDataResultsItem {
	s.MatchStatus = &v
	return s
}

type MatchTaskQueryResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s MatchTaskQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MatchTaskQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *MatchTaskQueryResponseExtra) SetDescription(v string) *MatchTaskQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *MatchTaskQueryResponseExtra) SetErrorCode(v int32) *MatchTaskQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MatchTaskQueryResponseExtra) SetLogid(v string) *MatchTaskQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *MatchTaskQueryResponseExtra) SetNow(v int64) *MatchTaskQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *MatchTaskQueryResponseExtra) SetSubDescription(v string) *MatchTaskQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *MatchTaskQueryResponseExtra) SetSubErrorCode(v int32) *MatchTaskQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

type MerchantClaimSubmitRequest struct {
	Datas       []*MerchantClaimSubmitRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" require:"true" type:"Repeated"`
	AccountId   *string                                `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MerchantClaimSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequest) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequest) SetDatas(v []*MerchantClaimSubmitRequestDatasItem) *MerchantClaimSubmitRequest {
	s.Datas = v
	return s
}

func (s *MerchantClaimSubmitRequest) SetAccountId(v string) *MerchantClaimSubmitRequest {
	s.AccountId = &v
	return s
}

func (s *MerchantClaimSubmitRequest) SetHeader(v map[string]*string) *MerchantClaimSubmitRequest {
	s.Header = v
	return s
}

func (s *MerchantClaimSubmitRequest) SetAccessToken(v string) *MerchantClaimSubmitRequest {
	s.AccessToken = &v
	return s
}

type MerchantClaimSubmitRequestDatasItem struct {
	AccountId       *string                                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	IndustryInfo    *MerchantClaimSubmitRequestDatasItemIndustryInfo    `json:"industry_info,omitempty" xml:"industry_info,omitempty" require:"true"`
	LegalPersonInfo *MerchantClaimSubmitRequestDatasItemLegalPersonInfo `json:"legal_person_info,omitempty" xml:"legal_person_info,omitempty"`
	LicenseInfo     *MerchantClaimSubmitRequestDatasItemLicenseInfo     `json:"license_info,omitempty" xml:"license_info,omitempty" require:"true"`
	PoiId           *string                                             `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
}

func (s MerchantClaimSubmitRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequestDatasItem) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequestDatasItem) SetAccountId(v string) *MerchantClaimSubmitRequestDatasItem {
	s.AccountId = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItem) SetIndustryInfo(v *MerchantClaimSubmitRequestDatasItemIndustryInfo) *MerchantClaimSubmitRequestDatasItem {
	s.IndustryInfo = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItem) SetLegalPersonInfo(v *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) *MerchantClaimSubmitRequestDatasItem {
	s.LegalPersonInfo = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItem) SetLicenseInfo(v *MerchantClaimSubmitRequestDatasItemLicenseInfo) *MerchantClaimSubmitRequestDatasItem {
	s.LicenseInfo = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItem) SetPoiId(v string) *MerchantClaimSubmitRequestDatasItem {
	s.PoiId = &v
	return s
}

type MerchantClaimSubmitRequestDatasItemIndustryInfo struct {
	Qualifications     []*MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem `json:"qualifications,omitempty" xml:"qualifications,omitempty" type:"Repeated"`
	MajorIndustryCode  *string                                                              `json:"major_industry_code,omitempty" xml:"major_industry_code,omitempty" require:"true"`
	MinorIndustryCodes []*string                                                            `json:"minor_industry_codes,omitempty" xml:"minor_industry_codes,omitempty" type:"Repeated"`
}

func (s MerchantClaimSubmitRequestDatasItemIndustryInfo) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequestDatasItemIndustryInfo) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfo) SetQualifications(v []*MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) *MerchantClaimSubmitRequestDatasItemIndustryInfo {
	s.Qualifications = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfo) SetMajorIndustryCode(v string) *MerchantClaimSubmitRequestDatasItemIndustryInfo {
	s.MajorIndustryCode = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfo) SetMinorIndustryCodes(v []*string) *MerchantClaimSubmitRequestDatasItemIndustryInfo {
	s.MinorIndustryCodes = v
	return s
}

type MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem struct {
	QualificationType       *int64    `json:"qualification_type,omitempty" xml:"qualification_type,omitempty" require:"true"`
	QualificationUrls       []*string `json:"qualification_urls,omitempty" xml:"qualification_urls,omitempty" require:"true" type:"Repeated"`
	QualificationExpiration *string   `json:"qualification_expiration,omitempty" xml:"qualification_expiration,omitempty" require:"true"`
}

func (s MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) SetQualificationType(v int64) *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem {
	s.QualificationType = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) SetQualificationUrls(v []*string) *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem {
	s.QualificationUrls = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem) SetQualificationExpiration(v string) *MerchantClaimSubmitRequestDatasItemIndustryInfoQualificationsItem {
	s.QualificationExpiration = &v
	return s
}

type MerchantClaimSubmitRequestDatasItemLegalPersonInfo struct {
	IdCardExpiration  *string `json:"id_card_expiration,omitempty" xml:"id_card_expiration,omitempty" require:"true"`
	IdCardFrontUrl    *string `json:"id_card_front_url,omitempty" xml:"id_card_front_url,omitempty" require:"true"`
	IdCardNo          *string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty" require:"true"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	QualificationType *int64  `json:"qualification_type,omitempty" xml:"qualification_type,omitempty" require:"true"`
	IdCardBackUrl     *string `json:"id_card_back_url,omitempty" xml:"id_card_back_url,omitempty" require:"true"`
}

func (s MerchantClaimSubmitRequestDatasItemLegalPersonInfo) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequestDatasItemLegalPersonInfo) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetIdCardExpiration(v string) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.IdCardExpiration = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetIdCardFrontUrl(v string) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.IdCardFrontUrl = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetIdCardNo(v string) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.IdCardNo = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetName(v string) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.Name = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetQualificationType(v int64) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.QualificationType = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLegalPersonInfo) SetIdCardBackUrl(v string) *MerchantClaimSubmitRequestDatasItemLegalPersonInfo {
	s.IdCardBackUrl = &v
	return s
}

type MerchantClaimSubmitRequestDatasItemLicenseInfo struct {
	Address         *string   `json:"address,omitempty" xml:"address,omitempty"`
	LegalPersonName *string   `json:"legal_person_name,omitempty" xml:"legal_person_name,omitempty" require:"true"`
	LicenseType     *int64    `json:"license_type,omitempty" xml:"license_type,omitempty" require:"true"`
	City            *string   `json:"city,omitempty" xml:"city,omitempty"`
	LicenseUrls     []*string `json:"license_urls,omitempty" xml:"license_urls,omitempty" require:"true" type:"Repeated"`
	LicenseId       *string   `json:"license_id,omitempty" xml:"license_id,omitempty" require:"true"`
	SalesRange      *string   `json:"sales_range,omitempty" xml:"sales_range,omitempty"`
	CompanyName     *string   `json:"company_name,omitempty" xml:"company_name,omitempty" require:"true"`
	Province        *string   `json:"province,omitempty" xml:"province,omitempty"`
	Expiration      *string   `json:"expiration,omitempty" xml:"expiration,omitempty" require:"true"`
}

func (s MerchantClaimSubmitRequestDatasItemLicenseInfo) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitRequestDatasItemLicenseInfo) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetAddress(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.Address = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetLegalPersonName(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.LegalPersonName = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetLicenseType(v int64) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.LicenseType = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetCity(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.City = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetLicenseUrls(v []*string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.LicenseUrls = v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetLicenseId(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.LicenseId = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetSalesRange(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.SalesRange = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetCompanyName(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.CompanyName = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetProvince(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.Province = &v
	return s
}

func (s *MerchantClaimSubmitRequestDatasItemLicenseInfo) SetExpiration(v string) *MerchantClaimSubmitRequestDatasItemLicenseInfo {
	s.Expiration = &v
	return s
}

type MerchantClaimSubmitResponse struct {
	Extra *MerchantClaimSubmitResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *MerchantClaimSubmitResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s MerchantClaimSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitResponse) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitResponse) SetExtra(v *MerchantClaimSubmitResponseExtra) *MerchantClaimSubmitResponse {
	s.Extra = v
	return s
}

func (s *MerchantClaimSubmitResponse) SetData(v *MerchantClaimSubmitResponseData) *MerchantClaimSubmitResponse {
	s.Data = v
	return s
}

type MerchantClaimSubmitResponseData struct {
	StoreClaimTaskInfos []*MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem `json:"store_claim_task_infos,omitempty" xml:"store_claim_task_infos,omitempty" type:"Repeated"`
	GwErrorCode         *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription       *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MerchantClaimSubmitResponseData) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitResponseData) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitResponseData) SetStoreClaimTaskInfos(v []*MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) *MerchantClaimSubmitResponseData {
	s.StoreClaimTaskInfos = v
	return s
}

func (s *MerchantClaimSubmitResponseData) SetGwErrorCode(v int32) *MerchantClaimSubmitResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *MerchantClaimSubmitResponseData) SetGwDescription(v string) *MerchantClaimSubmitResponseData {
	s.GwDescription = &v
	return s
}

type MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem struct {
	FailReason      *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	PoiId           *string `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	AccountId       *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ClaimResultType *int    `json:"claim_result_type,omitempty" xml:"claim_result_type,omitempty" require:"true"`
}

func (s MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) SetFailReason(v string) *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem {
	s.FailReason = &v
	return s
}

func (s *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) SetPoiId(v string) *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem {
	s.PoiId = &v
	return s
}

func (s *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) SetAccountId(v string) *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem {
	s.AccountId = &v
	return s
}

func (s *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem) SetClaimResultType(v int) *MerchantClaimSubmitResponseDataStoreClaimTaskInfosItem {
	s.ClaimResultType = &v
	return s
}

type MerchantClaimSubmitResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s MerchantClaimSubmitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MerchantClaimSubmitResponseExtra) GoString() string {
	return s.String()
}

func (s *MerchantClaimSubmitResponseExtra) SetSubErrorCode(v int32) *MerchantClaimSubmitResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *MerchantClaimSubmitResponseExtra) SetDescription(v string) *MerchantClaimSubmitResponseExtra {
	s.Description = &v
	return s
}

func (s *MerchantClaimSubmitResponseExtra) SetErrorCode(v int32) *MerchantClaimSubmitResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MerchantClaimSubmitResponseExtra) SetLogid(v string) *MerchantClaimSubmitResponseExtra {
	s.Logid = &v
	return s
}

func (s *MerchantClaimSubmitResponseExtra) SetNow(v int64) *MerchantClaimSubmitResponseExtra {
	s.Now = &v
	return s
}

func (s *MerchantClaimSubmitResponseExtra) SetSubDescription(v string) *MerchantClaimSubmitResponseExtra {
	s.SubDescription = &v
	return s
}

type MessageGetUserMessageRequest struct {
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Type        *int               `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Username    *string            `json:"username,omitempty" xml:"username,omitempty" require:"true"`
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MessageGetUserMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s MessageGetUserMessageRequest) GoString() string {
	return s.String()
}

func (s *MessageGetUserMessageRequest) SetEndTime(v int64) *MessageGetUserMessageRequest {
	s.EndTime = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetType(v int) *MessageGetUserMessageRequest {
	s.Type = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetUsername(v string) *MessageGetUserMessageRequest {
	s.Username = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetPageSize(v int64) *MessageGetUserMessageRequest {
	s.PageSize = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetPageNum(v int64) *MessageGetUserMessageRequest {
	s.PageNum = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetStartTime(v int64) *MessageGetUserMessageRequest {
	s.StartTime = &v
	return s
}

func (s *MessageGetUserMessageRequest) SetHeader(v map[string]*string) *MessageGetUserMessageRequest {
	s.Header = v
	return s
}

func (s *MessageGetUserMessageRequest) SetAccessToken(v string) *MessageGetUserMessageRequest {
	s.AccessToken = &v
	return s
}

type MessageGetUserMessageResponse struct {
	Data   *MessageGetUserMessageResponseData `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s MessageGetUserMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s MessageGetUserMessageResponse) GoString() string {
	return s.String()
}

func (s *MessageGetUserMessageResponse) SetData(v *MessageGetUserMessageResponseData) *MessageGetUserMessageResponse {
	s.Data = v
	return s
}

func (s *MessageGetUserMessageResponse) SetErrNo(v int32) *MessageGetUserMessageResponse {
	s.ErrNo = &v
	return s
}

func (s *MessageGetUserMessageResponse) SetErrMsg(v string) *MessageGetUserMessageResponse {
	s.ErrMsg = &v
	return s
}

func (s *MessageGetUserMessageResponse) SetLogId(v string) *MessageGetUserMessageResponse {
	s.LogId = &v
	return s
}

type MessageGetUserMessageResponseData struct {
	Messages []*MessageGetUserMessageResponseDataMessagesItem `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	Total    *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s MessageGetUserMessageResponseData) String() string {
	return tea.Prettify(s)
}

func (s MessageGetUserMessageResponseData) GoString() string {
	return s.String()
}

func (s *MessageGetUserMessageResponseData) SetMessages(v []*MessageGetUserMessageResponseDataMessagesItem) *MessageGetUserMessageResponseData {
	s.Messages = v
	return s
}

func (s *MessageGetUserMessageResponseData) SetTotal(v int64) *MessageGetUserMessageResponseData {
	s.Total = &v
	return s
}

type MessageGetUserMessageResponseDataMessagesItem struct {
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
}

func (s MessageGetUserMessageResponseDataMessagesItem) String() string {
	return tea.Prettify(s)
}

func (s MessageGetUserMessageResponseDataMessagesItem) GoString() string {
	return s.String()
}

func (s *MessageGetUserMessageResponseDataMessagesItem) SetTitle(v string) *MessageGetUserMessageResponseDataMessagesItem {
	s.Title = &v
	return s
}

func (s *MessageGetUserMessageResponseDataMessagesItem) SetContent(v string) *MessageGetUserMessageResponseDataMessagesItem {
	s.Content = &v
	return s
}

func (s *MessageGetUserMessageResponseDataMessagesItem) SetUserName(v string) *MessageGetUserMessageResponseDataMessagesItem {
	s.UserName = &v
	return s
}

func (s *MessageGetUserMessageResponseDataMessagesItem) SetType(v string) *MessageGetUserMessageResponseDataMessagesItem {
	s.Type = &v
	return s
}

func (s *MessageGetUserMessageResponseDataMessagesItem) SetSendTime(v string) *MessageGetUserMessageResponseDataMessagesItem {
	s.SendTime = &v
	return s
}

type MessageResourcesRequest struct {
	MessageId      *string            `json:"message_id,omitempty" xml:"message_id,omitempty" require:"true"`
	OpenId         *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ConversationId *string            `json:"conversation_id,omitempty" xml:"conversation_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MessageResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s MessageResourcesRequest) GoString() string {
	return s.String()
}

func (s *MessageResourcesRequest) SetMessageId(v string) *MessageResourcesRequest {
	s.MessageId = &v
	return s
}

func (s *MessageResourcesRequest) SetOpenId(v string) *MessageResourcesRequest {
	s.OpenId = &v
	return s
}

func (s *MessageResourcesRequest) SetConversationId(v string) *MessageResourcesRequest {
	s.ConversationId = &v
	return s
}

func (s *MessageResourcesRequest) SetHeader(v map[string]*string) *MessageResourcesRequest {
	s.Header = v
	return s
}

func (s *MessageResourcesRequest) SetAccessToken(v string) *MessageResourcesRequest {
	s.AccessToken = &v
	return s
}

type MessageResourcesResponse struct {
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *MessageResourcesResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s MessageResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s MessageResourcesResponse) GoString() string {
	return s.String()
}

func (s *MessageResourcesResponse) SetErrNo(v int32) *MessageResourcesResponse {
	s.ErrNo = &v
	return s
}

func (s *MessageResourcesResponse) SetErrMsg(v string) *MessageResourcesResponse {
	s.ErrMsg = &v
	return s
}

func (s *MessageResourcesResponse) SetLogId(v string) *MessageResourcesResponse {
	s.LogId = &v
	return s
}

func (s *MessageResourcesResponse) SetData(v *MessageResourcesResponseData) *MessageResourcesResponse {
	s.Data = v
	return s
}

type MessageResourcesResponseData struct {
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
	MediaType *string `json:"media_type,omitempty" xml:"media_type,omitempty"`
}

func (s MessageResourcesResponseData) String() string {
	return tea.Prettify(s)
}

func (s MessageResourcesResponseData) GoString() string {
	return s.String()
}

func (s *MessageResourcesResponseData) SetUrl(v string) *MessageResourcesResponseData {
	s.Url = &v
	return s
}

func (s *MessageResourcesResponseData) SetMediaType(v string) *MessageResourcesResponseData {
	s.MediaType = &v
	return s
}

type MsgGroupRequest struct {
	OpenId      *string                 `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Content     *MsgGroupRequestContent `json:"content,omitempty" xml:"content,omitempty"`
	Header      map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MsgGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequest) GoString() string {
	return s.String()
}

func (s *MsgGroupRequest) SetOpenId(v string) *MsgGroupRequest {
	s.OpenId = &v
	return s
}

func (s *MsgGroupRequest) SetContent(v *MsgGroupRequestContent) *MsgGroupRequest {
	s.Content = v
	return s
}

func (s *MsgGroupRequest) SetHeader(v map[string]*string) *MsgGroupRequest {
	s.Header = v
	return s
}

func (s *MsgGroupRequest) SetAccessToken(v string) *MsgGroupRequest {
	s.AccessToken = &v
	return s
}

type MsgGroupRequestContent struct {
	Image             *MsgGroupRequestContentImage             `json:"image,omitempty" xml:"image,omitempty"`
	TemplateCard      *MsgGroupRequestContentTemplateCard      `json:"template_card,omitempty" xml:"template_card,omitempty"`
	Text              *MsgGroupRequestContentText              `json:"text,omitempty" xml:"text,omitempty"`
	Video             *MsgGroupRequestContentVideo             `json:"video,omitempty" xml:"video,omitempty"`
	MsgType           *int                                     `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	RetainConsultCard *MsgGroupRequestContentRetainConsultCard `json:"retain_consult_card,omitempty" xml:"retain_consult_card,omitempty"`
	AppletCard        *MsgGroupRequestContentAppletCard        `json:"applet_card,omitempty" xml:"applet_card,omitempty"`
}

func (s MsgGroupRequestContent) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContent) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContent) SetImage(v *MsgGroupRequestContentImage) *MsgGroupRequestContent {
	s.Image = v
	return s
}

func (s *MsgGroupRequestContent) SetTemplateCard(v *MsgGroupRequestContentTemplateCard) *MsgGroupRequestContent {
	s.TemplateCard = v
	return s
}

func (s *MsgGroupRequestContent) SetText(v *MsgGroupRequestContentText) *MsgGroupRequestContent {
	s.Text = v
	return s
}

func (s *MsgGroupRequestContent) SetVideo(v *MsgGroupRequestContentVideo) *MsgGroupRequestContent {
	s.Video = v
	return s
}

func (s *MsgGroupRequestContent) SetMsgType(v int) *MsgGroupRequestContent {
	s.MsgType = &v
	return s
}

func (s *MsgGroupRequestContent) SetRetainConsultCard(v *MsgGroupRequestContentRetainConsultCard) *MsgGroupRequestContent {
	s.RetainConsultCard = v
	return s
}

func (s *MsgGroupRequestContent) SetAppletCard(v *MsgGroupRequestContentAppletCard) *MsgGroupRequestContent {
	s.AppletCard = v
	return s
}

type MsgGroupRequestContentAppletCard struct {
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	Schema         *string `json:"schema,omitempty" xml:"schema,omitempty"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	CardId         *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
}

func (s MsgGroupRequestContentAppletCard) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentAppletCard) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentAppletCard) SetCardTemplateId(v string) *MsgGroupRequestContentAppletCard {
	s.CardTemplateId = &v
	return s
}

func (s *MsgGroupRequestContentAppletCard) SetPath(v string) *MsgGroupRequestContentAppletCard {
	s.Path = &v
	return s
}

func (s *MsgGroupRequestContentAppletCard) SetQuery(v string) *MsgGroupRequestContentAppletCard {
	s.Query = &v
	return s
}

func (s *MsgGroupRequestContentAppletCard) SetSchema(v string) *MsgGroupRequestContentAppletCard {
	s.Schema = &v
	return s
}

func (s *MsgGroupRequestContentAppletCard) SetAppId(v string) *MsgGroupRequestContentAppletCard {
	s.AppId = &v
	return s
}

func (s *MsgGroupRequestContentAppletCard) SetCardId(v string) *MsgGroupRequestContentAppletCard {
	s.CardId = &v
	return s
}

type MsgGroupRequestContentImage struct {
	MediaId *string `json:"media_id,omitempty" xml:"media_id,omitempty"`
}

func (s MsgGroupRequestContentImage) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentImage) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentImage) SetMediaId(v string) *MsgGroupRequestContentImage {
	s.MediaId = &v
	return s
}

type MsgGroupRequestContentRetainConsultCard struct {
	CardId *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
}

func (s MsgGroupRequestContentRetainConsultCard) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentRetainConsultCard) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentRetainConsultCard) SetCardId(v string) *MsgGroupRequestContentRetainConsultCard {
	s.CardId = &v
	return s
}

type MsgGroupRequestContentTemplateCard struct {
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
}

func (s MsgGroupRequestContentTemplateCard) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentTemplateCard) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentTemplateCard) SetCardTemplateId(v string) *MsgGroupRequestContentTemplateCard {
	s.CardTemplateId = &v
	return s
}

type MsgGroupRequestContentText struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s MsgGroupRequestContentText) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentText) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentText) SetText(v string) *MsgGroupRequestContentText {
	s.Text = &v
	return s
}

type MsgGroupRequestContentVideo struct {
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s MsgGroupRequestContentVideo) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupRequestContentVideo) GoString() string {
	return s.String()
}

func (s *MsgGroupRequestContentVideo) SetItemId(v string) *MsgGroupRequestContentVideo {
	s.ItemId = &v
	return s
}

type MsgGroupResponse struct {
	Extra *MsgGroupResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	MsgId *string                `json:"msg_id,omitempty" xml:"msg_id,omitempty"`
	Data  *MsgGroupResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s MsgGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupResponse) GoString() string {
	return s.String()
}

func (s *MsgGroupResponse) SetExtra(v *MsgGroupResponseExtra) *MsgGroupResponse {
	s.Extra = v
	return s
}

func (s *MsgGroupResponse) SetMsgId(v string) *MsgGroupResponse {
	s.MsgId = &v
	return s
}

func (s *MsgGroupResponse) SetData(v *MsgGroupResponseData) *MsgGroupResponse {
	s.Data = v
	return s
}

type MsgGroupResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s MsgGroupResponseData) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupResponseData) GoString() string {
	return s.String()
}

func (s *MsgGroupResponseData) SetGwDescription(v string) *MsgGroupResponseData {
	s.GwDescription = &v
	return s
}

func (s *MsgGroupResponseData) SetGwErrorCode(v int32) *MsgGroupResponseData {
	s.GwErrorCode = &v
	return s
}

type MsgGroupResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MsgGroupResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MsgGroupResponseExtra) GoString() string {
	return s.String()
}

func (s *MsgGroupResponseExtra) SetErrorCode(v int32) *MsgGroupResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MsgGroupResponseExtra) SetLogid(v string) *MsgGroupResponseExtra {
	s.Logid = &v
	return s
}

func (s *MsgGroupResponseExtra) SetNow(v int64) *MsgGroupResponseExtra {
	s.Now = &v
	return s
}

func (s *MsgGroupResponseExtra) SetSubDescription(v string) *MsgGroupResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *MsgGroupResponseExtra) SetSubErrorCode(v int32) *MsgGroupResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *MsgGroupResponseExtra) SetDescription(v string) *MsgGroupResponseExtra {
	s.Description = &v
	return s
}

type MusicHotRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MusicHotRequest) String() string {
	return tea.Prettify(s)
}

func (s MusicHotRequest) GoString() string {
	return s.String()
}

func (s *MusicHotRequest) SetHeader(v map[string]*string) *MusicHotRequest {
	s.Header = v
	return s
}

func (s *MusicHotRequest) SetAccessToken(v string) *MusicHotRequest {
	s.AccessToken = &v
	return s
}

type MusicHotResponse struct {
	Data  *MusicHotResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *MusicHotResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s MusicHotResponse) String() string {
	return tea.Prettify(s)
}

func (s MusicHotResponse) GoString() string {
	return s.String()
}

func (s *MusicHotResponse) SetData(v *MusicHotResponseData) *MusicHotResponse {
	s.Data = v
	return s
}

func (s *MusicHotResponse) SetExtra(v *MusicHotResponseExtra) *MusicHotResponse {
	s.Extra = v
	return s
}

type MusicHotResponseData struct {
	List          []*MusicHotResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MusicHotResponseData) String() string {
	return tea.Prettify(s)
}

func (s MusicHotResponseData) GoString() string {
	return s.String()
}

func (s *MusicHotResponseData) SetList(v []*MusicHotResponseDataListItem) *MusicHotResponseData {
	s.List = v
	return s
}

func (s *MusicHotResponseData) SetGwErrorCode(v int32) *MusicHotResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *MusicHotResponseData) SetGwDescription(v string) *MusicHotResponseData {
	s.GwDescription = &v
	return s
}

type MusicHotResponseDataListItem struct {
	Cover    *string `json:"cover,omitempty" xml:"cover,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	Duration *int32  `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
	Author   *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	UseCount *int64  `json:"use_count,omitempty" xml:"use_count,omitempty" require:"true"`
	ShareUrl *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Rank     *int32  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s MusicHotResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s MusicHotResponseDataListItem) GoString() string {
	return s.String()
}

func (s *MusicHotResponseDataListItem) SetCover(v string) *MusicHotResponseDataListItem {
	s.Cover = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetTitle(v string) *MusicHotResponseDataListItem {
	s.Title = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetDuration(v int32) *MusicHotResponseDataListItem {
	s.Duration = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetAuthor(v string) *MusicHotResponseDataListItem {
	s.Author = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetUseCount(v int64) *MusicHotResponseDataListItem {
	s.UseCount = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetShareUrl(v string) *MusicHotResponseDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *MusicHotResponseDataListItem) SetRank(v int32) *MusicHotResponseDataListItem {
	s.Rank = &v
	return s
}

type MusicHotResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MusicHotResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MusicHotResponseExtra) GoString() string {
	return s.String()
}

func (s *MusicHotResponseExtra) SetSubErrorCode(v int32) *MusicHotResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *MusicHotResponseExtra) SetSubDescription(v string) *MusicHotResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *MusicHotResponseExtra) SetLogid(v string) *MusicHotResponseExtra {
	s.Logid = &v
	return s
}

func (s *MusicHotResponseExtra) SetNow(v int64) *MusicHotResponseExtra {
	s.Now = &v
	return s
}

func (s *MusicHotResponseExtra) SetErrorCode(v int32) *MusicHotResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MusicHotResponseExtra) SetDescription(v string) *MusicHotResponseExtra {
	s.Description = &v
	return s
}

type MusicOriginalRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MusicOriginalRequest) String() string {
	return tea.Prettify(s)
}

func (s MusicOriginalRequest) GoString() string {
	return s.String()
}

func (s *MusicOriginalRequest) SetHeader(v map[string]*string) *MusicOriginalRequest {
	s.Header = v
	return s
}

func (s *MusicOriginalRequest) SetAccessToken(v string) *MusicOriginalRequest {
	s.AccessToken = &v
	return s
}

type MusicOriginalResponse struct {
	Data  *MusicOriginalResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *MusicOriginalResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s MusicOriginalResponse) String() string {
	return tea.Prettify(s)
}

func (s MusicOriginalResponse) GoString() string {
	return s.String()
}

func (s *MusicOriginalResponse) SetData(v *MusicOriginalResponseData) *MusicOriginalResponse {
	s.Data = v
	return s
}

func (s *MusicOriginalResponse) SetExtra(v *MusicOriginalResponseExtra) *MusicOriginalResponse {
	s.Extra = v
	return s
}

type MusicOriginalResponseData struct {
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*MusicOriginalResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s MusicOriginalResponseData) String() string {
	return tea.Prettify(s)
}

func (s MusicOriginalResponseData) GoString() string {
	return s.String()
}

func (s *MusicOriginalResponseData) SetGwDescription(v string) *MusicOriginalResponseData {
	s.GwDescription = &v
	return s
}

func (s *MusicOriginalResponseData) SetList(v []*MusicOriginalResponseDataListItem) *MusicOriginalResponseData {
	s.List = v
	return s
}

func (s *MusicOriginalResponseData) SetGwErrorCode(v int32) *MusicOriginalResponseData {
	s.GwErrorCode = &v
	return s
}

type MusicOriginalResponseDataListItem struct {
	UseCount *int64  `json:"use_count,omitempty" xml:"use_count,omitempty" require:"true"`
	ShareUrl *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Rank     *int32  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	Cover    *string `json:"cover,omitempty" xml:"cover,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	Duration *int32  `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
	Author   *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
}

func (s MusicOriginalResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s MusicOriginalResponseDataListItem) GoString() string {
	return s.String()
}

func (s *MusicOriginalResponseDataListItem) SetUseCount(v int64) *MusicOriginalResponseDataListItem {
	s.UseCount = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetShareUrl(v string) *MusicOriginalResponseDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetRank(v int32) *MusicOriginalResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetCover(v string) *MusicOriginalResponseDataListItem {
	s.Cover = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetTitle(v string) *MusicOriginalResponseDataListItem {
	s.Title = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetDuration(v int32) *MusicOriginalResponseDataListItem {
	s.Duration = &v
	return s
}

func (s *MusicOriginalResponseDataListItem) SetAuthor(v string) *MusicOriginalResponseDataListItem {
	s.Author = &v
	return s
}

type MusicOriginalResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s MusicOriginalResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MusicOriginalResponseExtra) GoString() string {
	return s.String()
}

func (s *MusicOriginalResponseExtra) SetLogid(v string) *MusicOriginalResponseExtra {
	s.Logid = &v
	return s
}

func (s *MusicOriginalResponseExtra) SetNow(v int64) *MusicOriginalResponseExtra {
	s.Now = &v
	return s
}

func (s *MusicOriginalResponseExtra) SetErrorCode(v int32) *MusicOriginalResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MusicOriginalResponseExtra) SetDescription(v string) *MusicOriginalResponseExtra {
	s.Description = &v
	return s
}

func (s *MusicOriginalResponseExtra) SetSubErrorCode(v int32) *MusicOriginalResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *MusicOriginalResponseExtra) SetSubDescription(v string) *MusicOriginalResponseExtra {
	s.SubDescription = &v
	return s
}

type MusicSoarRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s MusicSoarRequest) String() string {
	return tea.Prettify(s)
}

func (s MusicSoarRequest) GoString() string {
	return s.String()
}

func (s *MusicSoarRequest) SetHeader(v map[string]*string) *MusicSoarRequest {
	s.Header = v
	return s
}

func (s *MusicSoarRequest) SetAccessToken(v string) *MusicSoarRequest {
	s.AccessToken = &v
	return s
}

type MusicSoarResponse struct {
	Extra *MusicSoarResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *MusicSoarResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s MusicSoarResponse) String() string {
	return tea.Prettify(s)
}

func (s MusicSoarResponse) GoString() string {
	return s.String()
}

func (s *MusicSoarResponse) SetExtra(v *MusicSoarResponseExtra) *MusicSoarResponse {
	s.Extra = v
	return s
}

func (s *MusicSoarResponse) SetData(v *MusicSoarResponseData) *MusicSoarResponse {
	s.Data = v
	return s
}

type MusicSoarResponseData struct {
	List          []*MusicSoarResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MusicSoarResponseData) String() string {
	return tea.Prettify(s)
}

func (s MusicSoarResponseData) GoString() string {
	return s.String()
}

func (s *MusicSoarResponseData) SetList(v []*MusicSoarResponseDataListItem) *MusicSoarResponseData {
	s.List = v
	return s
}

func (s *MusicSoarResponseData) SetGwErrorCode(v int32) *MusicSoarResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *MusicSoarResponseData) SetGwDescription(v string) *MusicSoarResponseData {
	s.GwDescription = &v
	return s
}

type MusicSoarResponseDataListItem struct {
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
	Duration *int32  `json:"duration,omitempty" xml:"duration,omitempty" require:"true"`
	Author   *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	UseCount *int64  `json:"use_count,omitempty" xml:"use_count,omitempty" require:"true"`
	ShareUrl *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Rank     *int32  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	Cover    *string `json:"cover,omitempty" xml:"cover,omitempty"`
}

func (s MusicSoarResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s MusicSoarResponseDataListItem) GoString() string {
	return s.String()
}

func (s *MusicSoarResponseDataListItem) SetTitle(v string) *MusicSoarResponseDataListItem {
	s.Title = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetDuration(v int32) *MusicSoarResponseDataListItem {
	s.Duration = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetAuthor(v string) *MusicSoarResponseDataListItem {
	s.Author = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetUseCount(v int64) *MusicSoarResponseDataListItem {
	s.UseCount = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetShareUrl(v string) *MusicSoarResponseDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetRank(v int32) *MusicSoarResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *MusicSoarResponseDataListItem) SetCover(v string) *MusicSoarResponseDataListItem {
	s.Cover = &v
	return s
}

type MusicSoarResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MusicSoarResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s MusicSoarResponseExtra) GoString() string {
	return s.String()
}

func (s *MusicSoarResponseExtra) SetSubErrorCode(v int32) *MusicSoarResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *MusicSoarResponseExtra) SetSubDescription(v string) *MusicSoarResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *MusicSoarResponseExtra) SetLogid(v string) *MusicSoarResponseExtra {
	s.Logid = &v
	return s
}

func (s *MusicSoarResponseExtra) SetNow(v int64) *MusicSoarResponseExtra {
	s.Now = &v
	return s
}

func (s *MusicSoarResponseExtra) SetErrorCode(v int32) *MusicSoarResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *MusicSoarResponseExtra) SetDescription(v string) *MusicSoarResponseExtra {
	s.Description = &v
	return s
}

type OauthAccessTokenRequest struct {
	Code         *string            `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	GrantType    *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty" require:"true"`
}

func (s OauthAccessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthAccessTokenRequest) SetCode(v string) *OauthAccessTokenRequest {
	s.Code = &v
	return s
}

func (s *OauthAccessTokenRequest) SetGrantType(v string) *OauthAccessTokenRequest {
	s.GrantType = &v
	return s
}

func (s *OauthAccessTokenRequest) SetClientKey(v string) *OauthAccessTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthAccessTokenRequest) SetHeader(v map[string]*string) *OauthAccessTokenRequest {
	s.Header = v
	return s
}

func (s *OauthAccessTokenRequest) SetClientSecret(v string) *OauthAccessTokenRequest {
	s.ClientSecret = &v
	return s
}

type OauthAccessTokenResponse struct {
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
	Data    *OauthAccessTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OauthAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthAccessTokenResponse) SetMessage(v string) *OauthAccessTokenResponse {
	s.Message = &v
	return s
}

func (s *OauthAccessTokenResponse) SetData(v *OauthAccessTokenResponseData) *OauthAccessTokenResponse {
	s.Data = v
	return s
}

type OauthAccessTokenResponseData struct {
	ErrorCode        *int64  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ExpiresIn        *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	OpenId           *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	RefreshToken     *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	LogId            *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Scope            *string `json:"scope,omitempty" xml:"scope,omitempty"`
	RefreshExpiresIn *int64  `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
	AccessToken      *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OauthAccessTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthAccessTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthAccessTokenResponseData) SetErrorCode(v int64) *OauthAccessTokenResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetDescription(v string) *OauthAccessTokenResponseData {
	s.Description = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetExpiresIn(v int64) *OauthAccessTokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetOpenId(v string) *OauthAccessTokenResponseData {
	s.OpenId = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetRefreshToken(v string) *OauthAccessTokenResponseData {
	s.RefreshToken = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetLogId(v string) *OauthAccessTokenResponseData {
	s.LogId = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetScope(v string) *OauthAccessTokenResponseData {
	s.Scope = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetRefreshExpiresIn(v int64) *OauthAccessTokenResponseData {
	s.RefreshExpiresIn = &v
	return s
}

func (s *OauthAccessTokenResponseData) SetAccessToken(v string) *OauthAccessTokenResponseData {
	s.AccessToken = &v
	return s
}

type OauthBusinessScopesRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OauthBusinessScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthBusinessScopesRequest) GoString() string {
	return s.String()
}

func (s *OauthBusinessScopesRequest) SetHeader(v map[string]*string) *OauthBusinessScopesRequest {
	s.Header = v
	return s
}

func (s *OauthBusinessScopesRequest) SetAccessToken(v string) *OauthBusinessScopesRequest {
	s.AccessToken = &v
	return s
}

type OauthBusinessScopesResponse struct {
	Data      []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrorCode *int64    `json:"error_code,omitempty" xml:"error_code,omitempty"`
	Message   *string   `json:"message,omitempty" xml:"message,omitempty"`
}

func (s OauthBusinessScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthBusinessScopesResponse) GoString() string {
	return s.String()
}

func (s *OauthBusinessScopesResponse) SetData(v []*string) *OauthBusinessScopesResponse {
	s.Data = v
	return s
}

func (s *OauthBusinessScopesResponse) SetErrorCode(v int64) *OauthBusinessScopesResponse {
	s.ErrorCode = &v
	return s
}

func (s *OauthBusinessScopesResponse) SetMessage(v string) *OauthBusinessScopesResponse {
	s.Message = &v
	return s
}

type OauthBusinessTokenRequest struct {
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Scope        *string            `json:"scope,omitempty" xml:"scope,omitempty" require:"true"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s OauthBusinessTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthBusinessTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthBusinessTokenRequest) SetOpenId(v string) *OauthBusinessTokenRequest {
	s.OpenId = &v
	return s
}

func (s *OauthBusinessTokenRequest) SetScope(v string) *OauthBusinessTokenRequest {
	s.Scope = &v
	return s
}

func (s *OauthBusinessTokenRequest) SetClientKey(v string) *OauthBusinessTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthBusinessTokenRequest) SetClientSecret(v string) *OauthBusinessTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *OauthBusinessTokenRequest) SetHeader(v map[string]*string) *OauthBusinessTokenRequest {
	s.Header = v
	return s
}

type OauthBusinessTokenResponse struct {
	Message   *string                         `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Data      *OauthBusinessTokenResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrorCode *int64                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s OauthBusinessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthBusinessTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthBusinessTokenResponse) SetMessage(v string) *OauthBusinessTokenResponse {
	s.Message = &v
	return s
}

func (s *OauthBusinessTokenResponse) SetData(v *OauthBusinessTokenResponseData) *OauthBusinessTokenResponse {
	s.Data = v
	return s
}

func (s *OauthBusinessTokenResponse) SetErrorCode(v int64) *OauthBusinessTokenResponse {
	s.ErrorCode = &v
	return s
}

type OauthBusinessTokenResponseData struct {
	BizRefreshExpiresIn *int64  `json:"biz_refresh_expires_in,omitempty" xml:"biz_refresh_expires_in,omitempty"`
	BizToken            *string `json:"biz_token,omitempty" xml:"biz_token,omitempty"`
	BizExpiresIn        *int64  `json:"biz_expires_in,omitempty" xml:"biz_expires_in,omitempty"`
	BizRefreshToken     *string `json:"biz_refresh_token,omitempty" xml:"biz_refresh_token,omitempty"`
}

func (s OauthBusinessTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthBusinessTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthBusinessTokenResponseData) SetBizRefreshExpiresIn(v int64) *OauthBusinessTokenResponseData {
	s.BizRefreshExpiresIn = &v
	return s
}

func (s *OauthBusinessTokenResponseData) SetBizToken(v string) *OauthBusinessTokenResponseData {
	s.BizToken = &v
	return s
}

func (s *OauthBusinessTokenResponseData) SetBizExpiresIn(v int64) *OauthBusinessTokenResponseData {
	s.BizExpiresIn = &v
	return s
}

func (s *OauthBusinessTokenResponseData) SetBizRefreshToken(v string) *OauthBusinessTokenResponseData {
	s.BizRefreshToken = &v
	return s
}

type OauthClientTokenRequest struct {
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	GrantType    *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
}

func (s OauthClientTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthClientTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthClientTokenRequest) SetClientSecret(v string) *OauthClientTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *OauthClientTokenRequest) SetHeader(v map[string]*string) *OauthClientTokenRequest {
	s.Header = v
	return s
}

func (s *OauthClientTokenRequest) SetGrantType(v string) *OauthClientTokenRequest {
	s.GrantType = &v
	return s
}

func (s *OauthClientTokenRequest) SetClientKey(v string) *OauthClientTokenRequest {
	s.ClientKey = &v
	return s
}

type OauthClientTokenResponse struct {
	Data    *OauthClientTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s OauthClientTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthClientTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthClientTokenResponse) SetData(v *OauthClientTokenResponseData) *OauthClientTokenResponse {
	s.Data = v
	return s
}

func (s *OauthClientTokenResponse) SetMessage(v string) *OauthClientTokenResponse {
	s.Message = &v
	return s
}

type OauthClientTokenResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

func (s OauthClientTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthClientTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthClientTokenResponseData) SetDescription(v string) *OauthClientTokenResponseData {
	s.Description = &v
	return s
}

func (s *OauthClientTokenResponseData) SetAccessToken(v string) *OauthClientTokenResponseData {
	s.AccessToken = &v
	return s
}

func (s *OauthClientTokenResponseData) SetExpiresIn(v int64) *OauthClientTokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *OauthClientTokenResponseData) SetErrorCode(v int64) *OauthClientTokenResponseData {
	s.ErrorCode = &v
	return s
}

type OauthRefreshBizTokenRequest struct {
	RefreshToken *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty"`
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s OauthRefreshBizTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshBizTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthRefreshBizTokenRequest) SetRefreshToken(v string) *OauthRefreshBizTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *OauthRefreshBizTokenRequest) SetClientKey(v string) *OauthRefreshBizTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthRefreshBizTokenRequest) SetClientSecret(v string) *OauthRefreshBizTokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *OauthRefreshBizTokenRequest) SetHeader(v map[string]*string) *OauthRefreshBizTokenRequest {
	s.Header = v
	return s
}

type OauthRefreshBizTokenResponse struct {
	Message   *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Data      *OauthRefreshBizTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrorCode *int64                            `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

func (s OauthRefreshBizTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshBizTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthRefreshBizTokenResponse) SetMessage(v string) *OauthRefreshBizTokenResponse {
	s.Message = &v
	return s
}

func (s *OauthRefreshBizTokenResponse) SetData(v *OauthRefreshBizTokenResponseData) *OauthRefreshBizTokenResponse {
	s.Data = v
	return s
}

func (s *OauthRefreshBizTokenResponse) SetErrorCode(v int64) *OauthRefreshBizTokenResponse {
	s.ErrorCode = &v
	return s
}

type OauthRefreshBizTokenResponseData struct {
	BizRefreshToken     *string `json:"biz_refresh_token,omitempty" xml:"biz_refresh_token,omitempty"`
	BizRefreshExpiresIn *int64  `json:"biz_refresh_expires_in,omitempty" xml:"biz_refresh_expires_in,omitempty"`
	BizToken            *string `json:"biz_token,omitempty" xml:"biz_token,omitempty"`
	BizExpiresIn        *int64  `json:"biz_expires_in,omitempty" xml:"biz_expires_in,omitempty"`
}

func (s OauthRefreshBizTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshBizTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthRefreshBizTokenResponseData) SetBizRefreshToken(v string) *OauthRefreshBizTokenResponseData {
	s.BizRefreshToken = &v
	return s
}

func (s *OauthRefreshBizTokenResponseData) SetBizRefreshExpiresIn(v int64) *OauthRefreshBizTokenResponseData {
	s.BizRefreshExpiresIn = &v
	return s
}

func (s *OauthRefreshBizTokenResponseData) SetBizToken(v string) *OauthRefreshBizTokenResponseData {
	s.BizToken = &v
	return s
}

func (s *OauthRefreshBizTokenResponseData) SetBizExpiresIn(v int64) *OauthRefreshBizTokenResponseData {
	s.BizExpiresIn = &v
	return s
}

type OauthRefreshTokenRequest struct {
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
	RefreshToken *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty" require:"true"`
	GrantType    *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s OauthRefreshTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthRefreshTokenRequest) SetClientKey(v string) *OauthRefreshTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthRefreshTokenRequest) SetRefreshToken(v string) *OauthRefreshTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *OauthRefreshTokenRequest) SetGrantType(v string) *OauthRefreshTokenRequest {
	s.GrantType = &v
	return s
}

func (s *OauthRefreshTokenRequest) SetHeader(v map[string]*string) *OauthRefreshTokenRequest {
	s.Header = v
	return s
}

type OauthRefreshTokenResponse struct {
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
	Data    *OauthRefreshTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OauthRefreshTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthRefreshTokenResponse) SetMessage(v string) *OauthRefreshTokenResponse {
	s.Message = &v
	return s
}

func (s *OauthRefreshTokenResponse) SetData(v *OauthRefreshTokenResponseData) *OauthRefreshTokenResponse {
	s.Data = v
	return s
}

type OauthRefreshTokenResponseData struct {
	Scope            *string `json:"scope,omitempty" xml:"scope,omitempty"`
	RefreshExpiresIn *int64  `json:"refresh_expires_in,omitempty" xml:"refresh_expires_in,omitempty"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	OpenId           *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	LogId            *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ExpiresIn        *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	RefreshToken     *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	AccessToken      *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ErrorCode        *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

func (s OauthRefreshTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthRefreshTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthRefreshTokenResponseData) SetScope(v string) *OauthRefreshTokenResponseData {
	s.Scope = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetRefreshExpiresIn(v int64) *OauthRefreshTokenResponseData {
	s.RefreshExpiresIn = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetDescription(v string) *OauthRefreshTokenResponseData {
	s.Description = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetOpenId(v string) *OauthRefreshTokenResponseData {
	s.OpenId = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetLogId(v string) *OauthRefreshTokenResponseData {
	s.LogId = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetExpiresIn(v int64) *OauthRefreshTokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetRefreshToken(v string) *OauthRefreshTokenResponseData {
	s.RefreshToken = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetAccessToken(v string) *OauthRefreshTokenResponseData {
	s.AccessToken = &v
	return s
}

func (s *OauthRefreshTokenResponseData) SetErrorCode(v int64) *OauthRefreshTokenResponseData {
	s.ErrorCode = &v
	return s
}

type OauthRenewRefreshTokenRequest struct {
	RefreshToken *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty" require:"true"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s OauthRenewRefreshTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthRenewRefreshTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthRenewRefreshTokenRequest) SetRefreshToken(v string) *OauthRenewRefreshTokenRequest {
	s.RefreshToken = &v
	return s
}

func (s *OauthRenewRefreshTokenRequest) SetClientKey(v string) *OauthRenewRefreshTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthRenewRefreshTokenRequest) SetHeader(v map[string]*string) *OauthRenewRefreshTokenRequest {
	s.Header = v
	return s
}

type OauthRenewRefreshTokenResponse struct {
	Data    *OauthRenewRefreshTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s OauthRenewRefreshTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthRenewRefreshTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthRenewRefreshTokenResponse) SetData(v *OauthRenewRefreshTokenResponseData) *OauthRenewRefreshTokenResponse {
	s.Data = v
	return s
}

func (s *OauthRenewRefreshTokenResponse) SetMessage(v string) *OauthRenewRefreshTokenResponse {
	s.Message = &v
	return s
}

type OauthRenewRefreshTokenResponseData struct {
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorCode    *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ExpiresIn    *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

func (s OauthRenewRefreshTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthRenewRefreshTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthRenewRefreshTokenResponseData) SetDescription(v string) *OauthRenewRefreshTokenResponseData {
	s.Description = &v
	return s
}

func (s *OauthRenewRefreshTokenResponseData) SetErrorCode(v int64) *OauthRenewRefreshTokenResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OauthRenewRefreshTokenResponseData) SetExpiresIn(v int64) *OauthRenewRefreshTokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *OauthRenewRefreshTokenResponseData) SetRefreshToken(v string) *OauthRenewRefreshTokenResponseData {
	s.RefreshToken = &v
	return s
}

type OauthStableClientTokenRequest struct {
	GrantType    *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty" require:"true"`
}

func (s OauthStableClientTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthStableClientTokenRequest) GoString() string {
	return s.String()
}

func (s *OauthStableClientTokenRequest) SetGrantType(v string) *OauthStableClientTokenRequest {
	s.GrantType = &v
	return s
}

func (s *OauthStableClientTokenRequest) SetClientKey(v string) *OauthStableClientTokenRequest {
	s.ClientKey = &v
	return s
}

func (s *OauthStableClientTokenRequest) SetHeader(v map[string]*string) *OauthStableClientTokenRequest {
	s.Header = v
	return s
}

func (s *OauthStableClientTokenRequest) SetClientSecret(v string) *OauthStableClientTokenRequest {
	s.ClientSecret = &v
	return s
}

type OauthStableClientTokenResponse struct {
	Data    *OauthStableClientTokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s OauthStableClientTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthStableClientTokenResponse) GoString() string {
	return s.String()
}

func (s *OauthStableClientTokenResponse) SetData(v *OauthStableClientTokenResponseData) *OauthStableClientTokenResponse {
	s.Data = v
	return s
}

func (s *OauthStableClientTokenResponse) SetMessage(v string) *OauthStableClientTokenResponse {
	s.Message = &v
	return s
}

type OauthStableClientTokenResponseData struct {
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

func (s OauthStableClientTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthStableClientTokenResponseData) GoString() string {
	return s.String()
}

func (s *OauthStableClientTokenResponseData) SetErrorCode(v int64) *OauthStableClientTokenResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OauthStableClientTokenResponseData) SetDescription(v string) *OauthStableClientTokenResponseData {
	s.Description = &v
	return s
}

func (s *OauthStableClientTokenResponseData) SetAccessToken(v string) *OauthStableClientTokenResponseData {
	s.AccessToken = &v
	return s
}

func (s *OauthStableClientTokenResponseData) SetExpiresIn(v int64) *OauthStableClientTokenResponseData {
	s.ExpiresIn = &v
	return s
}

type OauthUserinfoRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OauthUserinfoRequest) String() string {
	return tea.Prettify(s)
}

func (s OauthUserinfoRequest) GoString() string {
	return s.String()
}

func (s *OauthUserinfoRequest) SetOpenId(v string) *OauthUserinfoRequest {
	s.OpenId = &v
	return s
}

func (s *OauthUserinfoRequest) SetHeader(v map[string]*string) *OauthUserinfoRequest {
	s.Header = v
	return s
}

func (s *OauthUserinfoRequest) SetAccessToken(v string) *OauthUserinfoRequest {
	s.AccessToken = &v
	return s
}

type OauthUserinfoResponse struct {
	Data   *OauthUserinfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s OauthUserinfoResponse) String() string {
	return tea.Prettify(s)
}

func (s OauthUserinfoResponse) GoString() string {
	return s.String()
}

func (s *OauthUserinfoResponse) SetData(v *OauthUserinfoResponseData) *OauthUserinfoResponse {
	s.Data = v
	return s
}

func (s *OauthUserinfoResponse) SetErrMsg(v string) *OauthUserinfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *OauthUserinfoResponse) SetErrNo(v int32) *OauthUserinfoResponse {
	s.ErrNo = &v
	return s
}

func (s *OauthUserinfoResponse) SetLogId(v string) *OauthUserinfoResponse {
	s.LogId = &v
	return s
}

type OauthUserinfoResponseData struct {
	Avatar       *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	OpenId       *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	Nickname     *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	ClientKey    *string `json:"client_key,omitempty" xml:"client_key,omitempty"`
	EAccountRole *string `json:"e_account_role,omitempty" xml:"e_account_role,omitempty"`
	UnionId      *string `json:"union_id,omitempty" xml:"union_id,omitempty"`
}

func (s OauthUserinfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s OauthUserinfoResponseData) GoString() string {
	return s.String()
}

func (s *OauthUserinfoResponseData) SetAvatar(v string) *OauthUserinfoResponseData {
	s.Avatar = &v
	return s
}

func (s *OauthUserinfoResponseData) SetOpenId(v string) *OauthUserinfoResponseData {
	s.OpenId = &v
	return s
}

func (s *OauthUserinfoResponseData) SetNickname(v string) *OauthUserinfoResponseData {
	s.Nickname = &v
	return s
}

func (s *OauthUserinfoResponseData) SetClientKey(v string) *OauthUserinfoResponseData {
	s.ClientKey = &v
	return s
}

func (s *OauthUserinfoResponseData) SetEAccountRole(v string) *OauthUserinfoResponseData {
	s.EAccountRole = &v
	return s
}

func (s *OauthUserinfoResponseData) SetUnionId(v string) *OauthUserinfoResponseData {
	s.UnionId = &v
	return s
}

type OnlineGetRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductIds  []*string          `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutIds      []*string          `json:"out_ids,omitempty" xml:"out_ids,omitempty" type:"Repeated"`
}

func (s OnlineGetRequest) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetRequest) GoString() string {
	return s.String()
}

func (s *OnlineGetRequest) SetHeader(v map[string]*string) *OnlineGetRequest {
	s.Header = v
	return s
}

func (s *OnlineGetRequest) SetAccessToken(v string) *OnlineGetRequest {
	s.AccessToken = &v
	return s
}

func (s *OnlineGetRequest) SetProductIds(v []*string) *OnlineGetRequest {
	s.ProductIds = v
	return s
}

func (s *OnlineGetRequest) SetAccountId(v string) *OnlineGetRequest {
	s.AccountId = &v
	return s
}

func (s *OnlineGetRequest) SetOutIds(v []*string) *OnlineGetRequest {
	s.OutIds = v
	return s
}

type OnlineGetResponse struct {
	Data     *OnlineGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *OnlineGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *OnlineGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s OnlineGetResponse) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponse) GoString() string {
	return s.String()
}

func (s *OnlineGetResponse) SetData(v *OnlineGetResponseData) *OnlineGetResponse {
	s.Data = v
	return s
}

func (s *OnlineGetResponse) SetExtra(v *OnlineGetResponseExtra) *OnlineGetResponse {
	s.Extra = v
	return s
}

func (s *OnlineGetResponse) SetBaseResp(v *OnlineGetResponseBaseResp) *OnlineGetResponse {
	s.BaseResp = v
	return s
}

type OnlineGetResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s OnlineGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseBaseResp) SetExtra(v map[string]*string) *OnlineGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *OnlineGetResponseBaseResp) SetStatusCode(v int32) *OnlineGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *OnlineGetResponseBaseResp) SetStatusMessage(v string) *OnlineGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type OnlineGetResponseData struct {
	Description    *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductOnlines []*OnlineGetResponseDataProductOnlinesItem `json:"product_onlines,omitempty" xml:"product_onlines,omitempty" type:"Repeated"`
}

func (s OnlineGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseData) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseData) SetDescription(v string) *OnlineGetResponseData {
	s.Description = &v
	return s
}

func (s *OnlineGetResponseData) SetErrorCode(v int32) *OnlineGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OnlineGetResponseData) SetProductOnlines(v []*OnlineGetResponseDataProductOnlinesItem) *OnlineGetResponseData {
	s.ProductOnlines = v
	return s
}

type OnlineGetResponseDataProductOnlinesItem struct {
	Sku            *OnlineGetResponseDataProductOnlinesItemSku            `json:"sku,omitempty" xml:"sku,omitempty"`
	CommissionInfo *OnlineGetResponseDataProductOnlinesItemCommissionInfo `json:"commission_info,omitempty" xml:"commission_info,omitempty"`
	OnlineStatus   *int                                                   `json:"online_status,omitempty" xml:"online_status,omitempty"`
	Product        *OnlineGetResponseDataProductOnlinesItemProduct        `json:"product,omitempty" xml:"product,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItem) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItem) SetSku(v *OnlineGetResponseDataProductOnlinesItemSku) *OnlineGetResponseDataProductOnlinesItem {
	s.Sku = v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItem) SetCommissionInfo(v *OnlineGetResponseDataProductOnlinesItemCommissionInfo) *OnlineGetResponseDataProductOnlinesItem {
	s.CommissionInfo = v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItem) SetOnlineStatus(v int) *OnlineGetResponseDataProductOnlinesItem {
	s.OnlineStatus = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItem) SetProduct(v *OnlineGetResponseDataProductOnlinesItemProduct) *OnlineGetResponseDataProductOnlinesItem {
	s.Product = v
	return s
}

type OnlineGetResponseDataProductOnlinesItemCommissionInfo struct {
	PlatformTakeRate *int64 `json:"platform_take_rate,omitempty" xml:"platform_take_rate,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItemCommissionInfo) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItemCommissionInfo) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItemCommissionInfo) SetPlatformTakeRate(v int64) *OnlineGetResponseDataProductOnlinesItemCommissionInfo {
	s.PlatformTakeRate = &v
	return s
}

type OnlineGetResponseDataProductOnlinesItemProduct struct {
	Pois             []*OnlineGetResponseDataProductOnlinesItemProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	SoldEndTime      *int64                                                    `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	OutUrl           *string                                                   `json:"out_url,omitempty" xml:"out_url,omitempty"`
	OpenBizType      *int                                                      `json:"open_biz_type,omitempty" xml:"open_biz_type,omitempty"`
	SoldStartTime    *int64                                                    `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	OutId            *string                                                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	AttrKeyValueMap  map[string]*string                                        `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	BizLine          *int                                                      `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	UpdateTime       *int64                                                    `json:"update_time,omitempty" xml:"update_time,omitempty"`
	CategoryId       *int64                                                    `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	CreateTime       *int64                                                    `json:"create_time,omitempty" xml:"create_time,omitempty"`
	ProductName      *string                                                   `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	CategoryFullName *string                                                   `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	CreatorAccountId *int64                                                    `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	ProductId        *string                                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OwnerAccountId   *int64                                                    `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	ProductType      *int                                                      `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	AccountName      *string                                                   `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItemProduct) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItemProduct) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetPois(v []*OnlineGetResponseDataProductOnlinesItemProductPoisItem) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.Pois = v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetSoldEndTime(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.SoldEndTime = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetOutUrl(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.OutUrl = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetOpenBizType(v int) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.OpenBizType = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetSoldStartTime(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.SoldStartTime = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetOutId(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.OutId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetAttrKeyValueMap(v map[string]*string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetBizLine(v int) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.BizLine = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetUpdateTime(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.UpdateTime = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetCategoryId(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.CategoryId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetCreateTime(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.CreateTime = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetProductName(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.ProductName = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetCategoryFullName(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.CategoryFullName = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetCreatorAccountId(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetProductId(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.ProductId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetOwnerAccountId(v int64) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetProductType(v int) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.ProductType = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProduct) SetAccountName(v string) *OnlineGetResponseDataProductOnlinesItemProduct {
	s.AccountName = &v
	return s
}

type OnlineGetResponseDataProductOnlinesItemProductPoisItem struct {
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItemProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItemProductPoisItem) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItemProductPoisItem) SetSupplierId(v int64) *OnlineGetResponseDataProductOnlinesItemProductPoisItem {
	s.SupplierId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProductPoisItem) SetPoiId(v string) *OnlineGetResponseDataProductOnlinesItemProductPoisItem {
	s.PoiId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemProductPoisItem) SetSupplierExtId(v string) *OnlineGetResponseDataProductOnlinesItemProductPoisItem {
	s.SupplierExtId = &v
	return s
}

type OnlineGetResponseDataProductOnlinesItemSku struct {
	Status          *int                                             `json:"status,omitempty" xml:"status,omitempty"`
	OriginAmount    *int64                                           `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	OutSkuId        *string                                          `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId           *string                                          `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName         *string                                          `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	ActualAmount    *int64                                           `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	Stock           *OnlineGetResponseDataProductOnlinesItemSkuStock `json:"stock,omitempty" xml:"stock,omitempty"`
	AttrKeyValueMap map[string]*string                               `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItemSku) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItemSku) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetStatus(v int) *OnlineGetResponseDataProductOnlinesItemSku {
	s.Status = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetOriginAmount(v int64) *OnlineGetResponseDataProductOnlinesItemSku {
	s.OriginAmount = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetOutSkuId(v string) *OnlineGetResponseDataProductOnlinesItemSku {
	s.OutSkuId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetSkuId(v string) *OnlineGetResponseDataProductOnlinesItemSku {
	s.SkuId = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetSkuName(v string) *OnlineGetResponseDataProductOnlinesItemSku {
	s.SkuName = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetActualAmount(v int64) *OnlineGetResponseDataProductOnlinesItemSku {
	s.ActualAmount = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetStock(v *OnlineGetResponseDataProductOnlinesItemSkuStock) *OnlineGetResponseDataProductOnlinesItemSku {
	s.Stock = v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSku) SetAttrKeyValueMap(v map[string]*string) *OnlineGetResponseDataProductOnlinesItemSku {
	s.AttrKeyValueMap = v
	return s
}

type OnlineGetResponseDataProductOnlinesItemSkuStock struct {
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
}

func (s OnlineGetResponseDataProductOnlinesItemSkuStock) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseDataProductOnlinesItemSkuStock) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseDataProductOnlinesItemSkuStock) SetLimitType(v int) *OnlineGetResponseDataProductOnlinesItemSkuStock {
	s.LimitType = &v
	return s
}

func (s *OnlineGetResponseDataProductOnlinesItemSkuStock) SetStockQty(v int64) *OnlineGetResponseDataProductOnlinesItemSkuStock {
	s.StockQty = &v
	return s
}

type OnlineGetResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s OnlineGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OnlineGetResponseExtra) GoString() string {
	return s.String()
}

func (s *OnlineGetResponseExtra) SetLogid(v string) *OnlineGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *OnlineGetResponseExtra) SetNow(v int64) *OnlineGetResponseExtra {
	s.Now = &v
	return s
}

func (s *OnlineGetResponseExtra) SetSubDescription(v string) *OnlineGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OnlineGetResponseExtra) SetSubErrorCode(v int32) *OnlineGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OnlineGetResponseExtra) SetDescription(v string) *OnlineGetResponseExtra {
	s.Description = &v
	return s
}

func (s *OnlineGetResponseExtra) SetErrorCode(v int32) *OnlineGetResponseExtra {
	s.ErrorCode = &v
	return s
}

type OnlineListRequest struct {
	DyPoiId     *int64                 `json:"dy_poi_id,omitempty" xml:"dy_poi_id,omitempty"`
	PageNo      *int32                 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	Status      []*int                 `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	PageSize    *int32                 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Base        *OnlineListRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                `json:"account_id,omitempty" xml:"account_id,omitempty"`
	DishType    *int                   `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	Header      map[string]*string     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OnlineListRequest) String() string {
	return tea.Prettify(s)
}

func (s OnlineListRequest) GoString() string {
	return s.String()
}

func (s *OnlineListRequest) SetDyPoiId(v int64) *OnlineListRequest {
	s.DyPoiId = &v
	return s
}

func (s *OnlineListRequest) SetPageNo(v int32) *OnlineListRequest {
	s.PageNo = &v
	return s
}

func (s *OnlineListRequest) SetStatus(v []*int) *OnlineListRequest {
	s.Status = v
	return s
}

func (s *OnlineListRequest) SetPageSize(v int32) *OnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *OnlineListRequest) SetBase(v *OnlineListRequestBase) *OnlineListRequest {
	s.Base = v
	return s
}

func (s *OnlineListRequest) SetAccountId(v string) *OnlineListRequest {
	s.AccountId = &v
	return s
}

func (s *OnlineListRequest) SetDishType(v int) *OnlineListRequest {
	s.DishType = &v
	return s
}

func (s *OnlineListRequest) SetHeader(v map[string]*string) *OnlineListRequest {
	s.Header = v
	return s
}

func (s *OnlineListRequest) SetAccessToken(v string) *OnlineListRequest {
	s.AccessToken = &v
	return s
}

type OnlineListRequestBase struct {
	LogID      *string                          `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *OnlineListRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                          `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                          `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                          `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string               `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s OnlineListRequestBase) String() string {
	return tea.Prettify(s)
}

func (s OnlineListRequestBase) GoString() string {
	return s.String()
}

func (s *OnlineListRequestBase) SetLogID(v string) *OnlineListRequestBase {
	s.LogID = &v
	return s
}

func (s *OnlineListRequestBase) SetTrafficEnv(v *OnlineListRequestBaseTrafficEnv) *OnlineListRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *OnlineListRequestBase) SetAddr(v string) *OnlineListRequestBase {
	s.Addr = &v
	return s
}

func (s *OnlineListRequestBase) SetCaller(v string) *OnlineListRequestBase {
	s.Caller = &v
	return s
}

func (s *OnlineListRequestBase) SetClient(v string) *OnlineListRequestBase {
	s.Client = &v
	return s
}

func (s *OnlineListRequestBase) SetExtra(v map[string]*string) *OnlineListRequestBase {
	s.Extra = v
	return s
}

type OnlineListRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s OnlineListRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s OnlineListRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *OnlineListRequestBaseTrafficEnv) SetOpen(v bool) *OnlineListRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *OnlineListRequestBaseTrafficEnv) SetEnv(v string) *OnlineListRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type OnlineListResponse struct {
	Data     *OnlineListResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *OnlineListResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *OnlineListResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s OnlineListResponse) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponse) GoString() string {
	return s.String()
}

func (s *OnlineListResponse) SetData(v *OnlineListResponseData) *OnlineListResponse {
	s.Data = v
	return s
}

func (s *OnlineListResponse) SetExtra(v *OnlineListResponseExtra) *OnlineListResponse {
	s.Extra = v
	return s
}

func (s *OnlineListResponse) SetBaseResp(v *OnlineListResponseBaseResp) *OnlineListResponse {
	s.BaseResp = v
	return s
}

type OnlineListResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s OnlineListResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseBaseResp) GoString() string {
	return s.String()
}

func (s *OnlineListResponseBaseResp) SetStatusMessage(v string) *OnlineListResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *OnlineListResponseBaseResp) SetExtra(v map[string]*string) *OnlineListResponseBaseResp {
	s.Extra = v
	return s
}

func (s *OnlineListResponseBaseResp) SetStatusCode(v int32) *OnlineListResponseBaseResp {
	s.StatusCode = &v
	return s
}

type OnlineListResponseData struct {
	Description *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Dishs       []*OnlineListResponseDataDishsItem `json:"dishs,omitempty" xml:"dishs,omitempty" type:"Repeated"`
	ErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore     *bool                              `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Total       *int64                             `json:"total,omitempty" xml:"total,omitempty"`
}

func (s OnlineListResponseData) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseData) GoString() string {
	return s.String()
}

func (s *OnlineListResponseData) SetDescription(v string) *OnlineListResponseData {
	s.Description = &v
	return s
}

func (s *OnlineListResponseData) SetDishs(v []*OnlineListResponseDataDishsItem) *OnlineListResponseData {
	s.Dishs = v
	return s
}

func (s *OnlineListResponseData) SetErrorCode(v int32) *OnlineListResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OnlineListResponseData) SetHasMore(v bool) *OnlineListResponseData {
	s.HasMore = &v
	return s
}

func (s *OnlineListResponseData) SetTotal(v int64) *OnlineListResponseData {
	s.Total = &v
	return s
}

type OnlineListResponseDataDishsItem struct {
	Attributes        []*OnlineListResponseDataDishsItemAttributesItem       `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	DishGroups        []*OnlineListResponseDataDishsItemDishGroupsItem       `json:"dish_groups,omitempty" xml:"dish_groups,omitempty" type:"Repeated"`
	ImageList         []*OnlineListResponseDataDishsItemImageListItem        `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	Skus              []*OnlineListResponseDataDishsItemSkusItem             `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	DraftStatus       *int                                                   `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	DeliveryMethod    []*int                                                 `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	ApplyDate         *OnlineListResponseDataDishsItemApplyDate              `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	OutId             *string                                                `json:"out_id,omitempty" xml:"out_id,omitempty"`
	MerchantProductId *string                                                `json:"merchant_product_id,omitempty" xml:"merchant_product_id,omitempty"`
	CreateTime        *int64                                                 `json:"create_time,omitempty" xml:"create_time,omitempty"`
	AccountName       *string                                                `json:"account_name,omitempty" xml:"account_name,omitempty"`
	PoiCount          *int64                                                 `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	ProductId         *int64                                                 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	DishType          *int                                                   `json:"dish_type,omitempty" xml:"dish_type,omitempty"`
	DishDescription   *string                                                `json:"dish_description,omitempty" xml:"dish_description,omitempty"`
	CategoryFullName  *string                                                `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	ProductName       *string                                                `json:"product_name,omitempty" xml:"product_name,omitempty"`
	CategoryId        *int64                                                 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	AddDishGroups     []*OnlineListResponseDataDishsItemAddDishGroupsItem    `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	AccountId         *string                                                `json:"account_id,omitempty" xml:"account_id,omitempty"`
	IsBindMerchant    *bool                                                  `json:"is_bind_merchant,omitempty" xml:"is_bind_merchant,omitempty"`
	UpdateTime        *int64                                                 `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductSpecAttrs  []*OnlineListResponseDataDishsItemProductSpecAttrsItem `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	DishDetailInfo    *OnlineListResponseDataDishsItemDishDetailInfo         `json:"dish_detail_info,omitempty" xml:"dish_detail_info,omitempty"`
	Pois              []*OnlineListResponseDataDishsItemPoisItem             `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	SettleType        *int64                                                 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	OnlineStatus      *int                                                   `json:"online_status,omitempty" xml:"online_status,omitempty"`
}

func (s OnlineListResponseDataDishsItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItem) SetAttributes(v []*OnlineListResponseDataDishsItemAttributesItem) *OnlineListResponseDataDishsItem {
	s.Attributes = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDishGroups(v []*OnlineListResponseDataDishsItemDishGroupsItem) *OnlineListResponseDataDishsItem {
	s.DishGroups = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetImageList(v []*OnlineListResponseDataDishsItemImageListItem) *OnlineListResponseDataDishsItem {
	s.ImageList = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetSkus(v []*OnlineListResponseDataDishsItemSkusItem) *OnlineListResponseDataDishsItem {
	s.Skus = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDraftStatus(v int) *OnlineListResponseDataDishsItem {
	s.DraftStatus = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDeliveryMethod(v []*int) *OnlineListResponseDataDishsItem {
	s.DeliveryMethod = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetApplyDate(v *OnlineListResponseDataDishsItemApplyDate) *OnlineListResponseDataDishsItem {
	s.ApplyDate = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetOutId(v string) *OnlineListResponseDataDishsItem {
	s.OutId = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetMerchantProductId(v string) *OnlineListResponseDataDishsItem {
	s.MerchantProductId = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetCreateTime(v int64) *OnlineListResponseDataDishsItem {
	s.CreateTime = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetAccountName(v string) *OnlineListResponseDataDishsItem {
	s.AccountName = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetPoiCount(v int64) *OnlineListResponseDataDishsItem {
	s.PoiCount = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetProductId(v int64) *OnlineListResponseDataDishsItem {
	s.ProductId = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDishType(v int) *OnlineListResponseDataDishsItem {
	s.DishType = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDishDescription(v string) *OnlineListResponseDataDishsItem {
	s.DishDescription = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetCategoryFullName(v string) *OnlineListResponseDataDishsItem {
	s.CategoryFullName = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetProductName(v string) *OnlineListResponseDataDishsItem {
	s.ProductName = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetCategoryId(v int64) *OnlineListResponseDataDishsItem {
	s.CategoryId = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetAddDishGroups(v []*OnlineListResponseDataDishsItemAddDishGroupsItem) *OnlineListResponseDataDishsItem {
	s.AddDishGroups = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetAccountId(v string) *OnlineListResponseDataDishsItem {
	s.AccountId = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetIsBindMerchant(v bool) *OnlineListResponseDataDishsItem {
	s.IsBindMerchant = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetUpdateTime(v int64) *OnlineListResponseDataDishsItem {
	s.UpdateTime = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetProductSpecAttrs(v []*OnlineListResponseDataDishsItemProductSpecAttrsItem) *OnlineListResponseDataDishsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetDishDetailInfo(v *OnlineListResponseDataDishsItemDishDetailInfo) *OnlineListResponseDataDishsItem {
	s.DishDetailInfo = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetPois(v []*OnlineListResponseDataDishsItemPoisItem) *OnlineListResponseDataDishsItem {
	s.Pois = v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetSettleType(v int64) *OnlineListResponseDataDishsItem {
	s.SettleType = &v
	return s
}

func (s *OnlineListResponseDataDishsItem) SetOnlineStatus(v int) *OnlineListResponseDataDishsItem {
	s.OnlineStatus = &v
	return s
}

type OnlineListResponseDataDishsItemAddDishGroupsItem struct {
	GroupId   *int64                                                          `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                         `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItem) SetGroupId(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItem) SetGroupName(v string) *OnlineListResponseDataDishsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItem) SetItemList(v []*OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) *OnlineListResponseDataDishsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

type OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem struct {
	ProductId           *int64                                                               `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OutId               *string                                                              `json:"out_id,omitempty" xml:"out_id,omitempty"`
	OutSkuId            *string                                                              `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	ResidueStock        *int64                                                               `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	IsSetAutoComplement *bool                                                                `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	PackFee             *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	Price               *int64                                                               `json:"price,omitempty" xml:"price,omitempty"`
	SkuId               *int64                                                               `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	IsSetSellOut        *bool                                                                `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	Stock               *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock   `json:"stock,omitempty" xml:"stock,omitempty"`
	MaxStock            *int64                                                               `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	ProductName         *string                                                              `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutId(v string) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetOutSkuId(v string) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetResidueStock(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ResidueStock = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetAutoComplement(v bool) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetPackFee(v *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.PackFee = v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetSkuId(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetIsSetSellOut(v bool) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.IsSetSellOut = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetStock(v *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.Stock = v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetMaxStock(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.MaxStock = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem) SetProductName(v string) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

type OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee struct {
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFeeUnit(v int) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFeeUnit = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetStep(v int32) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.Step = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee) SetPackFee(v int32) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemPackFee {
	s.PackFee = &v
	return s
}

type OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetSoldCount(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.SoldCount = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetSoldQty(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.SoldQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetStockQty(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.StockQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetAvailQty(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.AvailQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetFrozenQty(v int64) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.FrozenQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock) SetLimitType(v int) *OnlineListResponseDataDishsItemAddDishGroupsItemItemListItemStock {
	s.LimitType = &v
	return s
}

type OnlineListResponseDataDishsItemApplyDate struct {
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
}

func (s OnlineListResponseDataDishsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemApplyDate) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemApplyDate) SetUseDateType(v int) *OnlineListResponseDataDishsItemApplyDate {
	s.UseDateType = &v
	return s
}

func (s *OnlineListResponseDataDishsItemApplyDate) SetUseEndDate(v string) *OnlineListResponseDataDishsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *OnlineListResponseDataDishsItemApplyDate) SetUseStartDate(v string) *OnlineListResponseDataDishsItemApplyDate {
	s.UseStartDate = &v
	return s
}

func (s *OnlineListResponseDataDishsItemApplyDate) SetDayDuration(v int32) *OnlineListResponseDataDishsItemApplyDate {
	s.DayDuration = &v
	return s
}

type OnlineListResponseDataDishsItemAttributesItem struct {
	ItemList  []*OnlineListResponseDataDishsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                      `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s OnlineListResponseDataDishsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAttributesItem) SetItemList(v []*OnlineListResponseDataDishsItemAttributesItemItemListItem) *OnlineListResponseDataDishsItemAttributesItem {
	s.ItemList = v
	return s
}

func (s *OnlineListResponseDataDishsItemAttributesItem) SetGroupName(v string) *OnlineListResponseDataDishsItemAttributesItem {
	s.GroupName = &v
	return s
}

type OnlineListResponseDataDishsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OnlineListResponseDataDishsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemAttributesItemItemListItem) SetName(v string) *OnlineListResponseDataDishsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type OnlineListResponseDataDishsItemDishDetailInfo struct {
	MaterialInfo []*OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem `json:"material_info,omitempty" xml:"material_info,omitempty" type:"Repeated"`
}

func (s OnlineListResponseDataDishsItemDishDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemDishDetailInfo) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemDishDetailInfo) SetMaterialInfo(v []*OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) *OnlineListResponseDataDishsItemDishDetailInfo {
	s.MaterialInfo = v
	return s
}

type OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem struct {
	Value []*string `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
	Key   *string   `json:"key,omitempty" xml:"key,omitempty"`
	Name  *string   `json:"name,omitempty" xml:"name,omitempty"`
}

func (s OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetValue(v []*string) *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Value = v
	return s
}

func (s *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetKey(v string) *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Key = &v
	return s
}

func (s *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem) SetName(v string) *OnlineListResponseDataDishsItemDishDetailInfoMaterialInfoItem {
	s.Name = &v
	return s
}

type OnlineListResponseDataDishsItemDishGroupsItem struct {
	GroupRankingScore *string `json:"group_ranking_score,omitempty" xml:"group_ranking_score,omitempty"`
	DishGroupId       *int64  `json:"dish_group_id,omitempty" xml:"dish_group_id,omitempty"`
	DishGroupName     *string `json:"dish_group_name,omitempty" xml:"dish_group_name,omitempty"`
}

func (s OnlineListResponseDataDishsItemDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemDishGroupsItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemDishGroupsItem) SetGroupRankingScore(v string) *OnlineListResponseDataDishsItemDishGroupsItem {
	s.GroupRankingScore = &v
	return s
}

func (s *OnlineListResponseDataDishsItemDishGroupsItem) SetDishGroupId(v int64) *OnlineListResponseDataDishsItemDishGroupsItem {
	s.DishGroupId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemDishGroupsItem) SetDishGroupName(v string) *OnlineListResponseDataDishsItemDishGroupsItem {
	s.DishGroupName = &v
	return s
}

type OnlineListResponseDataDishsItemImageListItem struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s OnlineListResponseDataDishsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemImageListItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemImageListItem) SetUrl(v string) *OnlineListResponseDataDishsItemImageListItem {
	s.Url = &v
	return s
}

type OnlineListResponseDataDishsItemPoisItem struct {
	PoiId *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s OnlineListResponseDataDishsItemPoisItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemPoisItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemPoisItem) SetPoiId(v string) *OnlineListResponseDataDishsItemPoisItem {
	s.PoiId = &v
	return s
}

type OnlineListResponseDataDishsItemProductSpecAttrsItem struct {
	ItemList  []*OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s OnlineListResponseDataDishsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItem) SetItemList(v []*OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) *OnlineListResponseDataDishsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItem) SetGroupName(v string) *OnlineListResponseDataDishsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

type OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem struct {
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *OnlineListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

type OnlineListResponseDataDishsItemSkusItem struct {
	Stock               *OnlineListResponseDataDishsItemSkusItemStock              `json:"stock,omitempty" xml:"stock,omitempty"`
	OutSkuId            *string                                                    `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuName             *string                                                    `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	MaxStock            *int64                                                     `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	SkuSpecAttrs        []*OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	Status              *int                                                       `json:"status,omitempty" xml:"status,omitempty"`
	PackFee             *OnlineListResponseDataDishsItemSkusItemPackFee            `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	ActualAmount        *int64                                                     `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	IsSetAutoComplement *bool                                                      `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	ResidueStock        *int64                                                     `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	SkuId               *int64                                                     `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	IsSetSellOut        *bool                                                      `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
}

func (s OnlineListResponseDataDishsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemSkusItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetStock(v *OnlineListResponseDataDishsItemSkusItemStock) *OnlineListResponseDataDishsItemSkusItem {
	s.Stock = v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetOutSkuId(v string) *OnlineListResponseDataDishsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetSkuName(v string) *OnlineListResponseDataDishsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetMaxStock(v int64) *OnlineListResponseDataDishsItemSkusItem {
	s.MaxStock = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetSkuSpecAttrs(v []*OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem) *OnlineListResponseDataDishsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetStatus(v int) *OnlineListResponseDataDishsItemSkusItem {
	s.Status = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetPackFee(v *OnlineListResponseDataDishsItemSkusItemPackFee) *OnlineListResponseDataDishsItemSkusItem {
	s.PackFee = v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetActualAmount(v int64) *OnlineListResponseDataDishsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetIsSetAutoComplement(v bool) *OnlineListResponseDataDishsItemSkusItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetResidueStock(v int64) *OnlineListResponseDataDishsItemSkusItem {
	s.ResidueStock = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetSkuId(v int64) *OnlineListResponseDataDishsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItem) SetIsSetSellOut(v bool) *OnlineListResponseDataDishsItemSkusItem {
	s.IsSetSellOut = &v
	return s
}

type OnlineListResponseDataDishsItemSkusItemPackFee struct {
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s OnlineListResponseDataDishsItemSkusItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemSkusItemPackFee) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemSkusItemPackFee) SetStep(v int32) *OnlineListResponseDataDishsItemSkusItemPackFee {
	s.Step = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemPackFee) SetPackFee(v int32) *OnlineListResponseDataDishsItemSkusItemPackFee {
	s.PackFee = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemPackFee) SetPackFeeUnit(v int) *OnlineListResponseDataDishsItemSkusItemPackFee {
	s.PackFeeUnit = &v
	return s
}

type OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem struct {
	ItemList  []*OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

type OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *OnlineListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

type OnlineListResponseDataDishsItemSkusItemStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s OnlineListResponseDataDishsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseDataDishsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetSoldCount(v int64) *OnlineListResponseDataDishsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetSoldQty(v int64) *OnlineListResponseDataDishsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetStockQty(v int64) *OnlineListResponseDataDishsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetAvailQty(v int64) *OnlineListResponseDataDishsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetFrozenQty(v int64) *OnlineListResponseDataDishsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *OnlineListResponseDataDishsItemSkusItemStock) SetLimitType(v int) *OnlineListResponseDataDishsItemSkusItemStock {
	s.LimitType = &v
	return s
}

type OnlineListResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s OnlineListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OnlineListResponseExtra) GoString() string {
	return s.String()
}

func (s *OnlineListResponseExtra) SetLogid(v string) *OnlineListResponseExtra {
	s.Logid = &v
	return s
}

func (s *OnlineListResponseExtra) SetNow(v int64) *OnlineListResponseExtra {
	s.Now = &v
	return s
}

func (s *OnlineListResponseExtra) SetSubDescription(v string) *OnlineListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OnlineListResponseExtra) SetSubErrorCode(v int32) *OnlineListResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OnlineListResponseExtra) SetDescription(v string) *OnlineListResponseExtra {
	s.Description = &v
	return s
}

func (s *OnlineListResponseExtra) SetErrorCode(v int32) *OnlineListResponseExtra {
	s.ErrorCode = &v
	return s
}

type OpenGetticketRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OpenGetticketRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGetticketRequest) GoString() string {
	return s.String()
}

func (s *OpenGetticketRequest) SetHeader(v map[string]*string) *OpenGetticketRequest {
	s.Header = v
	return s
}

func (s *OpenGetticketRequest) SetAccessToken(v string) *OpenGetticketRequest {
	s.AccessToken = &v
	return s
}

type OpenGetticketResponse struct {
	Extra *OpenGetticketResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *OpenGetticketResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OpenGetticketResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGetticketResponse) GoString() string {
	return s.String()
}

func (s *OpenGetticketResponse) SetExtra(v *OpenGetticketResponseExtra) *OpenGetticketResponse {
	s.Extra = v
	return s
}

func (s *OpenGetticketResponse) SetData(v *OpenGetticketResponseData) *OpenGetticketResponse {
	s.Data = v
	return s
}

type OpenGetticketResponseData struct {
	ExpiresIn     *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty" require:"true"`
	Ticket        *string `json:"ticket,omitempty" xml:"ticket,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OpenGetticketResponseData) String() string {
	return tea.Prettify(s)
}

func (s OpenGetticketResponseData) GoString() string {
	return s.String()
}

func (s *OpenGetticketResponseData) SetExpiresIn(v int64) *OpenGetticketResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *OpenGetticketResponseData) SetTicket(v string) *OpenGetticketResponseData {
	s.Ticket = &v
	return s
}

func (s *OpenGetticketResponseData) SetGwErrorCode(v int32) *OpenGetticketResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OpenGetticketResponseData) SetGwDescription(v string) *OpenGetticketResponseData {
	s.GwDescription = &v
	return s
}

type OpenGetticketResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s OpenGetticketResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OpenGetticketResponseExtra) GoString() string {
	return s.String()
}

func (s *OpenGetticketResponseExtra) SetErrorCode(v int32) *OpenGetticketResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OpenGetticketResponseExtra) SetDescription(v string) *OpenGetticketResponseExtra {
	s.Description = &v
	return s
}

func (s *OpenGetticketResponseExtra) SetSubErrorCode(v int32) *OpenGetticketResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OpenGetticketResponseExtra) SetSubDescription(v string) *OpenGetticketResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OpenGetticketResponseExtra) SetLogid(v string) *OpenGetticketResponseExtra {
	s.Logid = &v
	return s
}

func (s *OpenGetticketResponseExtra) SetNow(v int64) *OpenGetticketResponseExtra {
	s.Now = &v
	return s
}

type OpenItemIdToEncryptIdRequest struct {
	VideoIds    []*string          `json:"video_ids,omitempty" xml:"video_ids,omitempty" type:"Repeated"`
	AccessKey   *string            `json:"access_key,omitempty" xml:"access_key,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OpenItemIdToEncryptIdRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenItemIdToEncryptIdRequest) GoString() string {
	return s.String()
}

func (s *OpenItemIdToEncryptIdRequest) SetVideoIds(v []*string) *OpenItemIdToEncryptIdRequest {
	s.VideoIds = v
	return s
}

func (s *OpenItemIdToEncryptIdRequest) SetAccessKey(v string) *OpenItemIdToEncryptIdRequest {
	s.AccessKey = &v
	return s
}

func (s *OpenItemIdToEncryptIdRequest) SetHeader(v map[string]*string) *OpenItemIdToEncryptIdRequest {
	s.Header = v
	return s
}

func (s *OpenItemIdToEncryptIdRequest) SetAccessToken(v string) *OpenItemIdToEncryptIdRequest {
	s.AccessToken = &v
	return s
}

type OpenItemIdToEncryptIdResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *OpenItemIdToEncryptIdResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s OpenItemIdToEncryptIdResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenItemIdToEncryptIdResponse) GoString() string {
	return s.String()
}

func (s *OpenItemIdToEncryptIdResponse) SetErrMsg(v string) *OpenItemIdToEncryptIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *OpenItemIdToEncryptIdResponse) SetLogId(v string) *OpenItemIdToEncryptIdResponse {
	s.LogId = &v
	return s
}

func (s *OpenItemIdToEncryptIdResponse) SetData(v *OpenItemIdToEncryptIdResponseData) *OpenItemIdToEncryptIdResponse {
	s.Data = v
	return s
}

func (s *OpenItemIdToEncryptIdResponse) SetErrNo(v int32) *OpenItemIdToEncryptIdResponse {
	s.ErrNo = &v
	return s
}

type OpenItemIdToEncryptIdResponseData struct {
	ConvertResult map[string]*string `json:"convert_result,omitempty" xml:"convert_result,omitempty" require:"true"`
}

func (s OpenItemIdToEncryptIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s OpenItemIdToEncryptIdResponseData) GoString() string {
	return s.String()
}

func (s *OpenItemIdToEncryptIdResponseData) SetConvertResult(v map[string]*string) *OpenItemIdToEncryptIdResponseData {
	s.ConvertResult = v
	return s
}

type OrderApplyRefundRequest struct {
	RefundReason   *OrderApplyRefundRequestRefundReason `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	Header         map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId        *string                              `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	OutAfterSaleId *string                              `json:"out_after_sale_id,omitempty" xml:"out_after_sale_id,omitempty"`
}

func (s OrderApplyRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *OrderApplyRefundRequest) SetRefundReason(v *OrderApplyRefundRequestRefundReason) *OrderApplyRefundRequest {
	s.RefundReason = v
	return s
}

func (s *OrderApplyRefundRequest) SetHeader(v map[string]*string) *OrderApplyRefundRequest {
	s.Header = v
	return s
}

func (s *OrderApplyRefundRequest) SetAccessToken(v string) *OrderApplyRefundRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderApplyRefundRequest) SetOrderId(v string) *OrderApplyRefundRequest {
	s.OrderId = &v
	return s
}

func (s *OrderApplyRefundRequest) SetOutAfterSaleId(v string) *OrderApplyRefundRequest {
	s.OutAfterSaleId = &v
	return s
}

type OrderApplyRefundRequestRefundReason struct {
	Desc          *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	ReasonCode    []*int64  `json:"reason_code,omitempty" xml:"reason_code,omitempty" require:"true" type:"Repeated"`
	UserMediaList []*string `json:"user_media_list,omitempty" xml:"user_media_list,omitempty" type:"Repeated"`
}

func (s OrderApplyRefundRequestRefundReason) String() string {
	return tea.Prettify(s)
}

func (s OrderApplyRefundRequestRefundReason) GoString() string {
	return s.String()
}

func (s *OrderApplyRefundRequestRefundReason) SetDesc(v string) *OrderApplyRefundRequestRefundReason {
	s.Desc = &v
	return s
}

func (s *OrderApplyRefundRequestRefundReason) SetReasonCode(v []*int64) *OrderApplyRefundRequestRefundReason {
	s.ReasonCode = v
	return s
}

func (s *OrderApplyRefundRequestRefundReason) SetUserMediaList(v []*string) *OrderApplyRefundRequestRefundReason {
	s.UserMediaList = v
	return s
}

type OrderApplyRefundResponse struct {
	Extra *OrderApplyRefundResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *OrderApplyRefundResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrderApplyRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *OrderApplyRefundResponse) SetExtra(v *OrderApplyRefundResponseExtra) *OrderApplyRefundResponse {
	s.Extra = v
	return s
}

func (s *OrderApplyRefundResponse) SetData(v *OrderApplyRefundResponseData) *OrderApplyRefundResponse {
	s.Data = v
	return s
}

type OrderApplyRefundResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderApplyRefundResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderApplyRefundResponseData) GoString() string {
	return s.String()
}

func (s *OrderApplyRefundResponseData) SetGwErrorCode(v int32) *OrderApplyRefundResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderApplyRefundResponseData) SetGwDescription(v string) *OrderApplyRefundResponseData {
	s.GwDescription = &v
	return s
}

type OrderApplyRefundResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderApplyRefundResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderApplyRefundResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderApplyRefundResponseExtra) SetErrorCode(v int32) *OrderApplyRefundResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderApplyRefundResponseExtra) SetLogid(v string) *OrderApplyRefundResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderApplyRefundResponseExtra) SetNow(v int64) *OrderApplyRefundResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderApplyRefundResponseExtra) SetSubDescription(v string) *OrderApplyRefundResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderApplyRefundResponseExtra) SetSubErrorCode(v int32) *OrderApplyRefundResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderApplyRefundResponseExtra) SetDescription(v string) *OrderApplyRefundResponseExtra {
	s.Description = &v
	return s
}

type OrderAuditRequest struct {
	ExtOrderId       *string            `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty" require:"true"`
	OrderId          *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ReceiptStatus    *int               `json:"receipt_status,omitempty" xml:"receipt_status,omitempty" require:"true"`
	RefuseReason     *string            `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefuseReasonCode *int               `json:"refuse_reason_code,omitempty" xml:"refuse_reason_code,omitempty"`
	AccountId        *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s OrderAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderAuditRequest) GoString() string {
	return s.String()
}

func (s *OrderAuditRequest) SetExtOrderId(v string) *OrderAuditRequest {
	s.ExtOrderId = &v
	return s
}

func (s *OrderAuditRequest) SetOrderId(v string) *OrderAuditRequest {
	s.OrderId = &v
	return s
}

func (s *OrderAuditRequest) SetReceiptStatus(v int) *OrderAuditRequest {
	s.ReceiptStatus = &v
	return s
}

func (s *OrderAuditRequest) SetRefuseReason(v string) *OrderAuditRequest {
	s.RefuseReason = &v
	return s
}

func (s *OrderAuditRequest) SetHeader(v map[string]*string) *OrderAuditRequest {
	s.Header = v
	return s
}

func (s *OrderAuditRequest) SetAccessToken(v string) *OrderAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderAuditRequest) SetRefuseReasonCode(v int) *OrderAuditRequest {
	s.RefuseReasonCode = &v
	return s
}

func (s *OrderAuditRequest) SetAccountId(v string) *OrderAuditRequest {
	s.AccountId = &v
	return s
}

type OrderAuditResponse struct {
	Data  *OrderAuditResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *OrderAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s OrderAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderAuditResponse) GoString() string {
	return s.String()
}

func (s *OrderAuditResponse) SetData(v *OrderAuditResponseData) *OrderAuditResponse {
	s.Data = v
	return s
}

func (s *OrderAuditResponse) SetExtra(v *OrderAuditResponseExtra) *OrderAuditResponse {
	s.Extra = v
	return s
}

type OrderAuditResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s OrderAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderAuditResponseData) GoString() string {
	return s.String()
}

func (s *OrderAuditResponseData) SetDescription(v string) *OrderAuditResponseData {
	s.Description = &v
	return s
}

func (s *OrderAuditResponseData) SetErrorCode(v int32) *OrderAuditResponseData {
	s.ErrorCode = &v
	return s
}

type OrderAuditResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s OrderAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderAuditResponseExtra) SetDescription(v string) *OrderAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderAuditResponseExtra) SetErrorCode(v int32) *OrderAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderAuditResponseExtra) SetLogid(v string) *OrderAuditResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderAuditResponseExtra) SetNow(v int64) *OrderAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderAuditResponseExtra) SetSubDescription(v string) *OrderAuditResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderAuditResponseExtra) SetSubErrorCode(v int32) *OrderAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

type OrderConfirmRequest struct {
	OrderId       *string                         `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	SourceOrderId *string                         `json:"source_order_id,omitempty" xml:"source_order_id,omitempty" require:"true"`
	Header        map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ConfirmInfo   *OrderConfirmRequestConfirmInfo `json:"confirm_info,omitempty" xml:"confirm_info,omitempty" require:"true"`
}

func (s OrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequest) SetOrderId(v string) *OrderConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *OrderConfirmRequest) SetSourceOrderId(v string) *OrderConfirmRequest {
	s.SourceOrderId = &v
	return s
}

func (s *OrderConfirmRequest) SetHeader(v map[string]*string) *OrderConfirmRequest {
	s.Header = v
	return s
}

func (s *OrderConfirmRequest) SetAccessToken(v string) *OrderConfirmRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderConfirmRequest) SetConfirmInfo(v *OrderConfirmRequestConfirmInfo) *OrderConfirmRequest {
	s.ConfirmInfo = v
	return s
}

type OrderConfirmRequestConfirmInfo struct {
	FreeTravelInfo *OrderConfirmRequestConfirmInfoFreeTravelInfo `json:"free_travel_info,omitempty" xml:"free_travel_info,omitempty"`
	HotelInfo      *OrderConfirmRequestConfirmInfoHotelInfo      `json:"hotel_info,omitempty" xml:"hotel_info,omitempty"`
	PlayInfo       *OrderConfirmRequestConfirmInfoPlayInfo       `json:"play_info,omitempty" xml:"play_info,omitempty"`
	RejectCode     *int32                                        `json:"reject_code,omitempty" xml:"reject_code,omitempty"`
	ConfirmResult  *int32                                        `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
	ExtraMsg       *string                                       `json:"extra_msg,omitempty" xml:"extra_msg,omitempty"`
}

func (s OrderConfirmRequestConfirmInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfo) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfo) SetFreeTravelInfo(v *OrderConfirmRequestConfirmInfoFreeTravelInfo) *OrderConfirmRequestConfirmInfo {
	s.FreeTravelInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfo) SetHotelInfo(v *OrderConfirmRequestConfirmInfoHotelInfo) *OrderConfirmRequestConfirmInfo {
	s.HotelInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfo) SetPlayInfo(v *OrderConfirmRequestConfirmInfoPlayInfo) *OrderConfirmRequestConfirmInfo {
	s.PlayInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfo) SetRejectCode(v int32) *OrderConfirmRequestConfirmInfo {
	s.RejectCode = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfo) SetConfirmResult(v int32) *OrderConfirmRequestConfirmInfo {
	s.ConfirmResult = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfo) SetExtraMsg(v string) *OrderConfirmRequestConfirmInfo {
	s.ExtraMsg = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfo struct {
	OnedayTourList []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem `json:"oneday_tour_list,omitempty" xml:"oneday_tour_list,omitempty" type:"Repeated"`
	TravelNum      *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum            `json:"travel_num,omitempty" xml:"travel_num,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfo) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfo) SetOnedayTourList(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) *OrderConfirmRequestConfirmInfoFreeTravelInfo {
	s.OnedayTourList = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfo) SetTravelNum(v *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum) *OrderConfirmRequestConfirmInfoFreeTravelInfo {
	s.TravelNum = v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem struct {
	PlayInfoList  []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem  `json:"play_info_list,omitempty" xml:"play_info_list,omitempty" type:"Repeated"`
	Sequence      *int32                                                                             `json:"sequence,omitempty" xml:"sequence,omitempty"`
	HotelInfoList []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem `json:"hotel_info_list,omitempty" xml:"hotel_info_list,omitempty" type:"Repeated"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) SetPlayInfoList(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem {
	s.PlayInfoList = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) SetSequence(v int32) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem {
	s.Sequence = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem) SetHotelInfoList(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItem {
	s.HotelInfoList = v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem struct {
	PoiInfo        []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem   `json:"poi_info,omitempty" xml:"poi_info,omitempty" type:"Repeated"`
	RoomItems      []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem `json:"room_items,omitempty" xml:"room_items,omitempty" type:"Repeated"`
	HotelConfirmNo *string                                                                                         `json:"hotel_confirm_no,omitempty" xml:"hotel_confirm_no,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) SetPoiInfo(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem {
	s.PoiInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) SetRoomItems(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem {
	s.RoomItems = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem) SetHotelConfirmNo(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItem {
	s.HotelConfirmNo = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem struct {
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem) SetPoiName(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem {
	s.PoiName = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem) SetPoiId(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemPoiInfoItem {
	s.PoiId = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem struct {
	Meals    *int32  `json:"meals,omitempty" xml:"meals,omitempty"`
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem) SetMeals(v int32) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem {
	s.Meals = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem) SetRoomType(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemHotelInfoListItemRoomItemsItem {
	s.RoomType = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem struct {
	ShowCerts     []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem `json:"show_certs,omitempty" xml:"show_certs,omitempty" type:"Repeated"`
	BookEndTime   *string                                                                                        `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
	BookStartTime *string                                                                                        `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
	EntranceTypes []*int32                                                                                       `json:"entrance_types,omitempty" xml:"entrance_types,omitempty" type:"Repeated"`
	PoiInfo       []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem   `json:"poi_info,omitempty" xml:"poi_info,omitempty" type:"Repeated"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) SetShowCerts(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem {
	s.ShowCerts = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) SetBookEndTime(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem {
	s.BookEndTime = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) SetBookStartTime(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem {
	s.BookStartTime = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) SetEntranceTypes(v []*int32) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem {
	s.EntranceTypes = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem) SetPoiInfo(v []*OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItem {
	s.PoiInfo = v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem) SetPoiId(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem {
	s.PoiId = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem) SetPoiName(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemPoiInfoItem {
	s.PoiName = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem struct {
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	QrCodeImage  *string `json:"qr_code_image,omitempty" xml:"qr_code_image,omitempty"`
	CardNo       *string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	CertNo       *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertNoQrCode *string `json:"cert_no_qr_code,omitempty" xml:"cert_no_qr_code,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) SetName(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem {
	s.Name = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) SetQrCodeImage(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem {
	s.QrCodeImage = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) SetCardNo(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem {
	s.CardNo = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) SetCertNo(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem {
	s.CertNo = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem) SetCertNoQrCode(v string) *OrderConfirmRequestConfirmInfoFreeTravelInfoOnedayTourListItemPlayInfoListItemShowCertsItem {
	s.CertNoQrCode = &v
	return s
}

type OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum struct {
	DayNum   *int32 `json:"day_num,omitempty" xml:"day_num,omitempty"`
	NightNum *int32 `json:"night_num,omitempty" xml:"night_num,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum) SetDayNum(v int32) *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum {
	s.DayNum = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum) SetNightNum(v int32) *OrderConfirmRequestConfirmInfoFreeTravelInfoTravelNum {
	s.NightNum = &v
	return s
}

type OrderConfirmRequestConfirmInfoHotelInfo struct {
	PoiInfo        []*OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem   `json:"poi_info,omitempty" xml:"poi_info,omitempty" type:"Repeated"`
	RoomItems      []*OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem `json:"room_items,omitempty" xml:"room_items,omitempty" type:"Repeated"`
	HotelConfirmNo *string                                                 `json:"hotel_confirm_no,omitempty" xml:"hotel_confirm_no,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoHotelInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoHotelInfo) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoHotelInfo) SetPoiInfo(v []*OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem) *OrderConfirmRequestConfirmInfoHotelInfo {
	s.PoiInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoHotelInfo) SetRoomItems(v []*OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem) *OrderConfirmRequestConfirmInfoHotelInfo {
	s.RoomItems = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoHotelInfo) SetHotelConfirmNo(v string) *OrderConfirmRequestConfirmInfoHotelInfo {
	s.HotelConfirmNo = &v
	return s
}

type OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem) SetPoiId(v string) *OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem {
	s.PoiId = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem) SetPoiName(v string) *OrderConfirmRequestConfirmInfoHotelInfoPoiInfoItem {
	s.PoiName = &v
	return s
}

type OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem struct {
	Meals    *int32  `json:"meals,omitempty" xml:"meals,omitempty"`
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem) SetMeals(v int32) *OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem {
	s.Meals = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem) SetRoomType(v string) *OrderConfirmRequestConfirmInfoHotelInfoRoomItemsItem {
	s.RoomType = &v
	return s
}

type OrderConfirmRequestConfirmInfoPlayInfo struct {
	EntranceTypes []*int32                                               `json:"entrance_types,omitempty" xml:"entrance_types,omitempty" type:"Repeated"`
	PoiInfo       []*OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem   `json:"poi_info,omitempty" xml:"poi_info,omitempty" type:"Repeated"`
	ShowCerts     []*OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem `json:"show_certs,omitempty" xml:"show_certs,omitempty" type:"Repeated"`
	BookEndTime   *string                                                `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
	BookStartTime *string                                                `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoPlayInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoPlayInfo) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoPlayInfo) SetEntranceTypes(v []*int32) *OrderConfirmRequestConfirmInfoPlayInfo {
	s.EntranceTypes = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfo) SetPoiInfo(v []*OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem) *OrderConfirmRequestConfirmInfoPlayInfo {
	s.PoiInfo = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfo) SetShowCerts(v []*OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) *OrderConfirmRequestConfirmInfoPlayInfo {
	s.ShowCerts = v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfo) SetBookEndTime(v string) *OrderConfirmRequestConfirmInfoPlayInfo {
	s.BookEndTime = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfo) SetBookStartTime(v string) *OrderConfirmRequestConfirmInfoPlayInfo {
	s.BookStartTime = &v
	return s
}

type OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem) SetPoiId(v string) *OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem {
	s.PoiId = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem) SetPoiName(v string) *OrderConfirmRequestConfirmInfoPlayInfoPoiInfoItem {
	s.PoiName = &v
	return s
}

type OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem struct {
	CertNo       *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertNoQrCode *string `json:"cert_no_qr_code,omitempty" xml:"cert_no_qr_code,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	QrCodeImage  *string `json:"qr_code_image,omitempty" xml:"qr_code_image,omitempty"`
	CardNo       *string `json:"card_no,omitempty" xml:"card_no,omitempty"`
}

func (s OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) GoString() string {
	return s.String()
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) SetCertNo(v string) *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem {
	s.CertNo = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) SetCertNoQrCode(v string) *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem {
	s.CertNoQrCode = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) SetName(v string) *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem {
	s.Name = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) SetQrCodeImage(v string) *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem {
	s.QrCodeImage = &v
	return s
}

func (s *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem) SetCardNo(v string) *OrderConfirmRequestConfirmInfoPlayInfoShowCertsItem {
	s.CardNo = &v
	return s
}

type OrderConfirmResponse struct {
	Extra *OrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *OrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *OrderConfirmResponse) SetExtra(v *OrderConfirmResponseExtra) *OrderConfirmResponse {
	s.Extra = v
	return s
}

func (s *OrderConfirmResponse) SetData(v *OrderConfirmResponseData) *OrderConfirmResponse {
	s.Data = v
	return s
}

type OrderConfirmResponseData struct {
	ErrorCode     *string `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderOutId    *string `json:"order_out_id,omitempty" xml:"order_out_id,omitempty"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Descripiton   *string `json:"descripiton,omitempty" xml:"descripiton,omitempty"`
}

func (s OrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *OrderConfirmResponseData) SetErrorCode(v string) *OrderConfirmResponseData {
	s.ErrorCode = &v
	return s
}

func (s *OrderConfirmResponseData) SetOrderId(v string) *OrderConfirmResponseData {
	s.OrderId = &v
	return s
}

func (s *OrderConfirmResponseData) SetOrderOutId(v string) *OrderConfirmResponseData {
	s.OrderOutId = &v
	return s
}

func (s *OrderConfirmResponseData) SetGwDescription(v string) *OrderConfirmResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderConfirmResponseData) SetDescripiton(v string) *OrderConfirmResponseData {
	s.Descripiton = &v
	return s
}

type OrderConfirmResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderConfirmResponseExtra) SetErrorCode(v int32) *OrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderConfirmResponseExtra) SetLogid(v string) *OrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderConfirmResponseExtra) SetNow(v int64) *OrderConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderConfirmResponseExtra) SetSubDescription(v string) *OrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderConfirmResponseExtra) SetSubErrorCode(v int32) *OrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderConfirmResponseExtra) SetDescription(v string) *OrderConfirmResponseExtra {
	s.Description = &v
	return s
}

type OrderCreateOrderRequest struct {
	DiscountAmount         *int64                                         `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	OrderEntrySchema       *OrderCreateOrderRequestOrderEntrySchema       `json:"order_entry_schema,omitempty" xml:"order_entry_schema,omitempty"`
	TradeOption            *string                                        `json:"trade_option,omitempty" xml:"trade_option,omitempty"`
	PayExpireSeconds       *int64                                         `json:"pay_expire_seconds,omitempty" xml:"pay_expire_seconds,omitempty"`
	FeeList                []*OrderCreateOrderRequestFeeListItem          `json:"fee_list,omitempty" xml:"fee_list,omitempty" type:"Repeated"`
	Extra                  *string                                        `json:"extra,omitempty" xml:"extra,omitempty"`
	Header                 map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	OutOrderNo             *string                                        `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	ConsumeShopInfo        *OrderCreateOrderRequestConsumeShopInfo        `json:"consume_shop_info,omitempty" xml:"consume_shop_info,omitempty"`
	GoodsList              []*OrderCreateOrderRequestGoodsListItem        `json:"goods_list,omitempty" xml:"goods_list,omitempty" type:"Repeated"`
	BizLine                *int32                                         `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	CpExtra                *string                                        `json:"cp_extra,omitempty" xml:"cp_extra,omitempty"`
	PayNotifyUrl           *string                                        `json:"pay_notify_url,omitempty" xml:"pay_notify_url,omitempty"`
	CpBookInfo             *OrderCreateOrderRequestCpBookInfo             `json:"cp_book_info,omitempty" xml:"cp_book_info,omitempty"`
	LimitPayWayList        []*int64                                       `json:"limit_pay_way_list,omitempty" xml:"limit_pay_way_list,omitempty" type:"Repeated"`
	ContactName            *string                                        `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	TotalAmount            *int64                                         `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	PhoneNum               *string                                        `json:"phone_num,omitempty" xml:"phone_num,omitempty"`
	DeliveryInfo           *OrderCreateOrderRequestDeliveryInfo           `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
	OpenId                 *string                                        `json:"open_id,omitempty" xml:"open_id,omitempty"`
	UserAddress            *OrderCreateOrderRequestUserAddress            `json:"user_address,omitempty" xml:"user_address,omitempty"`
	SkuList                []*OrderCreateOrderRequestSkuListItem          `json:"sku_list,omitempty" xml:"sku_list,omitempty" type:"Repeated"`
	ReserveType            *int32                                         `json:"reserve_type,omitempty" xml:"reserve_type,omitempty"`
	PriceCalculationDetail *OrderCreateOrderRequestPriceCalculationDetail `json:"price_calculation_detail,omitempty" xml:"price_calculation_detail,omitempty"`
	AccessToken            *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OrderCreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequest) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequest) SetDiscountAmount(v int64) *OrderCreateOrderRequest {
	s.DiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequest) SetOrderEntrySchema(v *OrderCreateOrderRequestOrderEntrySchema) *OrderCreateOrderRequest {
	s.OrderEntrySchema = v
	return s
}

func (s *OrderCreateOrderRequest) SetTradeOption(v string) *OrderCreateOrderRequest {
	s.TradeOption = &v
	return s
}

func (s *OrderCreateOrderRequest) SetPayExpireSeconds(v int64) *OrderCreateOrderRequest {
	s.PayExpireSeconds = &v
	return s
}

func (s *OrderCreateOrderRequest) SetFeeList(v []*OrderCreateOrderRequestFeeListItem) *OrderCreateOrderRequest {
	s.FeeList = v
	return s
}

func (s *OrderCreateOrderRequest) SetExtra(v string) *OrderCreateOrderRequest {
	s.Extra = &v
	return s
}

func (s *OrderCreateOrderRequest) SetHeader(v map[string]*string) *OrderCreateOrderRequest {
	s.Header = v
	return s
}

func (s *OrderCreateOrderRequest) SetOutOrderNo(v string) *OrderCreateOrderRequest {
	s.OutOrderNo = &v
	return s
}

func (s *OrderCreateOrderRequest) SetConsumeShopInfo(v *OrderCreateOrderRequestConsumeShopInfo) *OrderCreateOrderRequest {
	s.ConsumeShopInfo = v
	return s
}

func (s *OrderCreateOrderRequest) SetGoodsList(v []*OrderCreateOrderRequestGoodsListItem) *OrderCreateOrderRequest {
	s.GoodsList = v
	return s
}

func (s *OrderCreateOrderRequest) SetBizLine(v int32) *OrderCreateOrderRequest {
	s.BizLine = &v
	return s
}

func (s *OrderCreateOrderRequest) SetCpExtra(v string) *OrderCreateOrderRequest {
	s.CpExtra = &v
	return s
}

func (s *OrderCreateOrderRequest) SetPayNotifyUrl(v string) *OrderCreateOrderRequest {
	s.PayNotifyUrl = &v
	return s
}

func (s *OrderCreateOrderRequest) SetCpBookInfo(v *OrderCreateOrderRequestCpBookInfo) *OrderCreateOrderRequest {
	s.CpBookInfo = v
	return s
}

func (s *OrderCreateOrderRequest) SetLimitPayWayList(v []*int64) *OrderCreateOrderRequest {
	s.LimitPayWayList = v
	return s
}

func (s *OrderCreateOrderRequest) SetContactName(v string) *OrderCreateOrderRequest {
	s.ContactName = &v
	return s
}

func (s *OrderCreateOrderRequest) SetTotalAmount(v int64) *OrderCreateOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *OrderCreateOrderRequest) SetPhoneNum(v string) *OrderCreateOrderRequest {
	s.PhoneNum = &v
	return s
}

func (s *OrderCreateOrderRequest) SetDeliveryInfo(v *OrderCreateOrderRequestDeliveryInfo) *OrderCreateOrderRequest {
	s.DeliveryInfo = v
	return s
}

func (s *OrderCreateOrderRequest) SetOpenId(v string) *OrderCreateOrderRequest {
	s.OpenId = &v
	return s
}

func (s *OrderCreateOrderRequest) SetUserAddress(v *OrderCreateOrderRequestUserAddress) *OrderCreateOrderRequest {
	s.UserAddress = v
	return s
}

func (s *OrderCreateOrderRequest) SetSkuList(v []*OrderCreateOrderRequestSkuListItem) *OrderCreateOrderRequest {
	s.SkuList = v
	return s
}

func (s *OrderCreateOrderRequest) SetReserveType(v int32) *OrderCreateOrderRequest {
	s.ReserveType = &v
	return s
}

func (s *OrderCreateOrderRequest) SetPriceCalculationDetail(v *OrderCreateOrderRequestPriceCalculationDetail) *OrderCreateOrderRequest {
	s.PriceCalculationDetail = v
	return s
}

func (s *OrderCreateOrderRequest) SetAccessToken(v string) *OrderCreateOrderRequest {
	s.AccessToken = &v
	return s
}

type OrderCreateOrderRequestConsumeShopInfo struct {
	ShopId      *string                                            `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	ShopName    *string                                            `json:"shop_name,omitempty" xml:"shop_name,omitempty" require:"true"`
	ShopPhone   *string                                            `json:"shop_phone,omitempty" xml:"shop_phone,omitempty" require:"true"`
	ShopAddress *OrderCreateOrderRequestConsumeShopInfoShopAddress `json:"shop_address,omitempty" xml:"shop_address,omitempty" require:"true"`
}

func (s OrderCreateOrderRequestConsumeShopInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestConsumeShopInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestConsumeShopInfo) SetShopId(v string) *OrderCreateOrderRequestConsumeShopInfo {
	s.ShopId = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfo) SetShopName(v string) *OrderCreateOrderRequestConsumeShopInfo {
	s.ShopName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfo) SetShopPhone(v string) *OrderCreateOrderRequestConsumeShopInfo {
	s.ShopPhone = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfo) SetShopAddress(v *OrderCreateOrderRequestConsumeShopInfoShopAddress) *OrderCreateOrderRequestConsumeShopInfo {
	s.ShopAddress = v
	return s
}

type OrderCreateOrderRequestConsumeShopInfoShopAddress struct {
	DoorPlateNum    *string                                                           `json:"door_plate_num,omitempty" xml:"door_plate_num,omitempty"`
	Gender          *int32                                                            `json:"gender,omitempty" xml:"gender,omitempty"`
	LocationAddress *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress `json:"location_address,omitempty" xml:"location_address,omitempty"`
	Phone           *string                                                           `json:"phone,omitempty" xml:"phone,omitempty"`
	UserAddressId   *string                                                           `json:"user_address_id,omitempty" xml:"user_address_id,omitempty"`
	ConnectName     *string                                                           `json:"connect_name,omitempty" xml:"connect_name,omitempty"`
	DetailAddress   *string                                                           `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

func (s OrderCreateOrderRequestConsumeShopInfoShopAddress) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestConsumeShopInfoShopAddress) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetDoorPlateNum(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.DoorPlateNum = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetGender(v int32) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.Gender = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetLocationAddress(v *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.LocationAddress = v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetPhone(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.Phone = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetUserAddressId(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.UserAddressId = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetConnectName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.ConnectName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddress) SetDetailAddress(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddress {
	s.DetailAddress = &v
	return s
}

type OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress struct {
	TownName        *string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	CityName        *string `json:"city_name,omitempty" xml:"city_name,omitempty" require:"true"`
	ProvinceCode    *string `json:"province_code,omitempty" xml:"province_code,omitempty" require:"true"`
	LocationId      *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	DistrictName    *string `json:"district_name,omitempty" xml:"district_name,omitempty" require:"true"`
	ProvinceName    *string `json:"province_name,omitempty" xml:"province_name,omitempty" require:"true"`
	TownCode        *string `json:"town_code,omitempty" xml:"town_code,omitempty"`
	DistrictCode    *string `json:"district_code,omitempty" xml:"district_code,omitempty" require:"true"`
	Lode            *string `json:"lode,omitempty" xml:"lode,omitempty"`
	LocationAddress *string `json:"location_address,omitempty" xml:"location_address,omitempty"`
	CityCode        *string `json:"city_code,omitempty" xml:"city_code,omitempty" require:"true"`
	LocationName    *string `json:"location_name,omitempty" xml:"location_name,omitempty"`
	Lade            *string `json:"lade,omitempty" xml:"lade,omitempty"`
}

func (s OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetTownName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.TownName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetCityName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.CityName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetProvinceCode(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.ProvinceCode = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetLocationId(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.LocationId = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetDistrictName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.DistrictName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetProvinceName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.ProvinceName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetTownCode(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.TownCode = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetDistrictCode(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.DistrictCode = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetLode(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.Lode = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetLocationAddress(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.LocationAddress = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetCityCode(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.CityCode = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetLocationName(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.LocationName = &v
	return s
}

func (s *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress) SetLade(v string) *OrderCreateOrderRequestConsumeShopInfoShopAddressLocationAddress {
	s.Lade = &v
	return s
}

type OrderCreateOrderRequestCpBookInfo struct {
	ItemBookInfoList []*OrderCreateOrderRequestCpBookInfoItemBookInfoListItem `json:"item_book_info_list,omitempty" xml:"item_book_info_list,omitempty" type:"Repeated"`
	OutBookNo        *string                                                  `json:"out_book_no,omitempty" xml:"out_book_no,omitempty"`
}

func (s OrderCreateOrderRequestCpBookInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfo) SetItemBookInfoList(v []*OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) *OrderCreateOrderRequestCpBookInfo {
	s.ItemBookInfoList = v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfo) SetOutBookNo(v string) *OrderCreateOrderRequestCpBookInfo {
	s.OutBookNo = &v
	return s
}

type OrderCreateOrderRequestCpBookInfoItemBookInfoListItem struct {
	BookExtra     *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra          `json:"book_extra,omitempty" xml:"book_extra,omitempty"`
	SkuId         *string                                                                  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	UserInfoList  []*OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem `json:"user_info_list,omitempty" xml:"user_info_list,omitempty" type:"Repeated"`
	ShopName      *string                                                                  `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	PoiId         *string                                                                  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtShopId     *string                                                                  `json:"ext_shop_id,omitempty" xml:"ext_shop_id,omitempty"`
	BookSkuInfo   *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo        `json:"book_sku_info,omitempty" xml:"book_sku_info,omitempty"`
	BookStartTime *int64                                                                   `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
	ItemOrderId   *string                                                                  `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	GoodsId       *string                                                                  `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	BookRangeType *int64                                                                   `json:"book_range_type,omitempty" xml:"book_range_type,omitempty"`
	BookEndTime   *int64                                                                   `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetBookExtra(v *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.BookExtra = v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetSkuId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.SkuId = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetUserInfoList(v []*OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.UserInfoList = v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetShopName(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.ShopName = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetPoiId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.PoiId = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetExtShopId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.ExtShopId = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetBookSkuInfo(v *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.BookSkuInfo = v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetBookStartTime(v int64) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.BookStartTime = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetItemOrderId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.ItemOrderId = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetGoodsId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetBookRangeType(v int64) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.BookRangeType = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem) SetBookEndTime(v int64) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItem {
	s.BookEndTime = &v
	return s
}

type OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra struct {
	HotelTravel *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel `json:"hotel_travel,omitempty" xml:"hotel_travel,omitempty"`
	BizType     *int32                                                                     `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra) SetHotelTravel(v *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra {
	s.HotelTravel = v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra) SetBizType(v int32) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtra {
	s.BizType = &v
	return s
}

type OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel struct {
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel) SetRoomType(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookExtraHotelTravel {
	s.RoomType = &v
	return s
}

type OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo struct {
	SkuIdType *int32  `json:"sku_id_type,omitempty" xml:"sku_id_type,omitempty" require:"true"`
	Price     *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SkuId     *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) SetSkuIdType(v int32) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo {
	s.SkuIdType = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) SetPrice(v int64) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo {
	s.Price = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo) SetSkuId(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemBookSkuInfo {
	s.SkuId = &v
	return s
}

type OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem struct {
	IdCardNo *string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) SetIdCardNo(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem {
	s.IdCardNo = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) SetName(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem {
	s.Name = &v
	return s
}

func (s *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem) SetPhone(v string) *OrderCreateOrderRequestCpBookInfoItemBookInfoListItemUserInfoListItem {
	s.Phone = &v
	return s
}

type OrderCreateOrderRequestDeliveryInfo struct {
	PickupEndTime    *int64  `json:"pickup_end_time,omitempty" xml:"pickup_end_time,omitempty"`
	PickupStartTime  *int64  `json:"pickup_start_time,omitempty" xml:"pickup_start_time,omitempty"`
	PoiId            *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ReserveEndTime   *int64  `json:"reserve_end_time,omitempty" xml:"reserve_end_time,omitempty"`
	ReserveStartTime *int64  `json:"reserve_start_time,omitempty" xml:"reserve_start_time,omitempty"`
	DeliveryMode     *int64  `json:"delivery_mode,omitempty" xml:"delivery_mode,omitempty"`
	ExpectEndTime    *int64  `json:"expect_end_time,omitempty" xml:"expect_end_time,omitempty"`
	ExpectStartTime  *int64  `json:"expect_start_time,omitempty" xml:"expect_start_time,omitempty"`
}

func (s OrderCreateOrderRequestDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestDeliveryInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetPickupEndTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.PickupEndTime = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetPickupStartTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.PickupStartTime = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetPoiId(v string) *OrderCreateOrderRequestDeliveryInfo {
	s.PoiId = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetReserveEndTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.ReserveEndTime = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetReserveStartTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.ReserveStartTime = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetDeliveryMode(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.DeliveryMode = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetExpectEndTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.ExpectEndTime = &v
	return s
}

func (s *OrderCreateOrderRequestDeliveryInfo) SetExpectStartTime(v int64) *OrderCreateOrderRequestDeliveryInfo {
	s.ExpectStartTime = &v
	return s
}

type OrderCreateOrderRequestFeeListItem struct {
	FeeType           *int32  `json:"fee_type,omitempty" xml:"fee_type,omitempty" require:"true"`
	OrderId           *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderIdType       *int32  `json:"order_id_type,omitempty" xml:"order_id_type,omitempty" require:"true"`
	FeeAmount         *int64  `json:"fee_amount,omitempty" xml:"fee_amount,omitempty" require:"true"`
	FeeDesc           *string `json:"fee_desc,omitempty" xml:"fee_desc,omitempty"`
	FeeDiscountAmount *int64  `json:"fee_discount_amount,omitempty" xml:"fee_discount_amount,omitempty"`
}

func (s OrderCreateOrderRequestFeeListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestFeeListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestFeeListItem) SetFeeType(v int32) *OrderCreateOrderRequestFeeListItem {
	s.FeeType = &v
	return s
}

func (s *OrderCreateOrderRequestFeeListItem) SetOrderId(v string) *OrderCreateOrderRequestFeeListItem {
	s.OrderId = &v
	return s
}

func (s *OrderCreateOrderRequestFeeListItem) SetOrderIdType(v int32) *OrderCreateOrderRequestFeeListItem {
	s.OrderIdType = &v
	return s
}

func (s *OrderCreateOrderRequestFeeListItem) SetFeeAmount(v int64) *OrderCreateOrderRequestFeeListItem {
	s.FeeAmount = &v
	return s
}

func (s *OrderCreateOrderRequestFeeListItem) SetFeeDesc(v string) *OrderCreateOrderRequestFeeListItem {
	s.FeeDesc = &v
	return s
}

func (s *OrderCreateOrderRequestFeeListItem) SetFeeDiscountAmount(v int64) *OrderCreateOrderRequestFeeListItem {
	s.FeeDiscountAmount = &v
	return s
}

type OrderCreateOrderRequestGoodsListItem struct {
	Quantity       *int32                                              `json:"quantity,omitempty" xml:"quantity,omitempty"`
	DateRule       *string                                             `json:"date_rule,omitempty" xml:"date_rule,omitempty"`
	MerchantUid    *string                                             `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	OrderValidTime *OrderCreateOrderRequestGoodsListItemOrderValidTime `json:"order_valid_time,omitempty" xml:"order_valid_time,omitempty"`
	GoodsBookInfo  *OrderCreateOrderRequestGoodsListItemGoodsBookInfo  `json:"goods_book_info,omitempty" xml:"goods_book_info,omitempty"`
	GoodsIdType    *int32                                              `json:"goods_id_type,omitempty" xml:"goods_id_type,omitempty"`
	GoodsId        *string                                             `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	GoodsPage      *OrderCreateOrderRequestGoodsListItemGoodsPage      `json:"goods_page,omitempty" xml:"goods_page,omitempty"`
	Price          *int64                                              `json:"price,omitempty" xml:"price,omitempty"`
	GoodsTitle     *string                                             `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
	Labels         *string                                             `json:"labels,omitempty" xml:"labels,omitempty"`
	DiscountAmount *int64                                              `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	GoodsImage     *string                                             `json:"goods_image,omitempty" xml:"goods_image,omitempty"`
}

func (s OrderCreateOrderRequestGoodsListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestGoodsListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestGoodsListItem) SetQuantity(v int32) *OrderCreateOrderRequestGoodsListItem {
	s.Quantity = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetDateRule(v string) *OrderCreateOrderRequestGoodsListItem {
	s.DateRule = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetMerchantUid(v string) *OrderCreateOrderRequestGoodsListItem {
	s.MerchantUid = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetOrderValidTime(v *OrderCreateOrderRequestGoodsListItemOrderValidTime) *OrderCreateOrderRequestGoodsListItem {
	s.OrderValidTime = v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsBookInfo(v *OrderCreateOrderRequestGoodsListItemGoodsBookInfo) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsBookInfo = v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsIdType(v int32) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsIdType = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsId(v string) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsPage(v *OrderCreateOrderRequestGoodsListItemGoodsPage) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsPage = v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetPrice(v int64) *OrderCreateOrderRequestGoodsListItem {
	s.Price = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsTitle(v string) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsTitle = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetLabels(v string) *OrderCreateOrderRequestGoodsListItem {
	s.Labels = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetDiscountAmount(v int64) *OrderCreateOrderRequestGoodsListItem {
	s.DiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItem) SetGoodsImage(v string) *OrderCreateOrderRequestGoodsListItem {
	s.GoodsImage = &v
	return s
}

type OrderCreateOrderRequestGoodsListItemGoodsBookInfo struct {
	CancelPolicy      *int32 `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	BookType          *int32 `json:"book_type,omitempty" xml:"book_type,omitempty"`
	CancelAdvanceHour *int64 `json:"cancel_advance_hour,omitempty" xml:"cancel_advance_hour,omitempty"`
}

func (s OrderCreateOrderRequestGoodsListItemGoodsBookInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestGoodsListItemGoodsBookInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestGoodsListItemGoodsBookInfo) SetCancelPolicy(v int32) *OrderCreateOrderRequestGoodsListItemGoodsBookInfo {
	s.CancelPolicy = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItemGoodsBookInfo) SetBookType(v int32) *OrderCreateOrderRequestGoodsListItemGoodsBookInfo {
	s.BookType = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItemGoodsBookInfo) SetCancelAdvanceHour(v int64) *OrderCreateOrderRequestGoodsListItemGoodsBookInfo {
	s.CancelAdvanceHour = &v
	return s
}

type OrderCreateOrderRequestGoodsListItemGoodsPage struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s OrderCreateOrderRequestGoodsListItemGoodsPage) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestGoodsListItemGoodsPage) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestGoodsListItemGoodsPage) SetParams(v string) *OrderCreateOrderRequestGoodsListItemGoodsPage {
	s.Params = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItemGoodsPage) SetPath(v string) *OrderCreateOrderRequestGoodsListItemGoodsPage {
	s.Path = &v
	return s
}

type OrderCreateOrderRequestGoodsListItemOrderValidTime struct {
	ValidStartTime *int64 `json:"valid_start_time,omitempty" xml:"valid_start_time,omitempty"`
	ValidDuration  *int64 `json:"valid_duration,omitempty" xml:"valid_duration,omitempty"`
	ValidEndTime   *int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
}

func (s OrderCreateOrderRequestGoodsListItemOrderValidTime) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestGoodsListItemOrderValidTime) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestGoodsListItemOrderValidTime) SetValidStartTime(v int64) *OrderCreateOrderRequestGoodsListItemOrderValidTime {
	s.ValidStartTime = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItemOrderValidTime) SetValidDuration(v int64) *OrderCreateOrderRequestGoodsListItemOrderValidTime {
	s.ValidDuration = &v
	return s
}

func (s *OrderCreateOrderRequestGoodsListItemOrderValidTime) SetValidEndTime(v int64) *OrderCreateOrderRequestGoodsListItemOrderValidTime {
	s.ValidEndTime = &v
	return s
}

type OrderCreateOrderRequestOrderEntrySchema struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s OrderCreateOrderRequestOrderEntrySchema) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestOrderEntrySchema) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestOrderEntrySchema) SetParams(v string) *OrderCreateOrderRequestOrderEntrySchema {
	s.Params = &v
	return s
}

func (s *OrderCreateOrderRequestOrderEntrySchema) SetPath(v string) *OrderCreateOrderRequestOrderEntrySchema {
	s.Path = &v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetail struct {
	OrderDiscountDetail   *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail       `json:"order_discount_detail,omitempty" xml:"order_discount_detail,omitempty"`
	UseMergedMarketingApi *bool                                                                   `json:"use_merged_marketing_api,omitempty" xml:"use_merged_marketing_api,omitempty"`
	CalculationType       *int32                                                                  `json:"calculation_type,omitempty" xml:"calculation_type,omitempty"`
	ExtCalculationNo      *string                                                                 `json:"ext_calculation_no,omitempty" xml:"ext_calculation_no,omitempty"`
	GoodsDiscountDetail   []*OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem `json:"goods_discount_detail,omitempty" xml:"goods_discount_detail,omitempty" type:"Repeated"`
	ItemDiscountDetail    []*OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem  `json:"item_discount_detail,omitempty" xml:"item_discount_detail,omitempty" type:"Repeated"`
}

func (s OrderCreateOrderRequestPriceCalculationDetail) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetail) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetOrderDiscountDetail(v *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) *OrderCreateOrderRequestPriceCalculationDetail {
	s.OrderDiscountDetail = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetUseMergedMarketingApi(v bool) *OrderCreateOrderRequestPriceCalculationDetail {
	s.UseMergedMarketingApi = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetCalculationType(v int32) *OrderCreateOrderRequestPriceCalculationDetail {
	s.CalculationType = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetExtCalculationNo(v string) *OrderCreateOrderRequestPriceCalculationDetail {
	s.ExtCalculationNo = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetGoodsDiscountDetail(v []*OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) *OrderCreateOrderRequestPriceCalculationDetail {
	s.GoodsDiscountDetail = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetail) SetItemDiscountDetail(v []*OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) *OrderCreateOrderRequestPriceCalculationDetail {
	s.ItemDiscountDetail = v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem struct {
	MerchantTotalDiscount *int64                                                                                         `json:"merchant_total_discount,omitempty" xml:"merchant_total_discount,omitempty"`
	PlatformTotalDiscount *int64                                                                                         `json:"platform_total_discount,omitempty" xml:"platform_total_discount,omitempty"`
	Quantity              *int32                                                                                         `json:"quantity,omitempty" xml:"quantity,omitempty"`
	TotalAmount           *int64                                                                                         `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	TotalDiscountAmount   *int64                                                                                         `json:"total_discount_amount,omitempty" xml:"total_discount_amount,omitempty"`
	GoodsId               *string                                                                                        `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	MarketingDetailInfo   []*OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem `json:"marketing_detail_info,omitempty" xml:"marketing_detail_info,omitempty" type:"Repeated"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetMerchantTotalDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.MerchantTotalDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetPlatformTotalDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.PlatformTotalDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetQuantity(v int32) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.Quantity = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetTotalAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.TotalAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetTotalDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.TotalDiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetGoodsId(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem) SetMarketingDetailInfo(v []*OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItem {
	s.MarketingDetailInfo = v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem struct {
	DiscountRange   *int32             `json:"discount_range,omitempty" xml:"discount_range,omitempty"`
	Note            *string            `json:"note,omitempty" xml:"note,omitempty"`
	Title           *string            `json:"title,omitempty" xml:"title,omitempty"`
	CreatorType     *int32             `json:"creator_type,omitempty" xml:"creator_type,omitempty"`
	Type            *int32             `json:"type,omitempty" xml:"type,omitempty"`
	Value           *int64             `json:"value,omitempty" xml:"value,omitempty"`
	Subtype         *string            `json:"subtype,omitempty" xml:"subtype,omitempty"`
	Id              *string            `json:"id,omitempty" xml:"id,omitempty"`
	Code            *string            `json:"code,omitempty" xml:"code,omitempty"`
	Kind            *int32             `json:"kind,omitempty" xml:"kind,omitempty"`
	MarketingExtend map[string]*string `json:"marketing_extend,omitempty" xml:"marketing_extend,omitempty"`
	DiscountAmount  *int64             `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetDiscountRange(v int32) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.DiscountRange = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetNote(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Note = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetTitle(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Title = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetCreatorType(v int32) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.CreatorType = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetType(v int32) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Type = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetValue(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Value = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetSubtype(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Subtype = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetId(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Id = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetCode(v string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Code = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetKind(v int32) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.Kind = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetMarketingExtend(v map[string]*string) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.MarketingExtend = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem) SetDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailGoodsDiscountDetailItemMarketingDetailInfoItem {
	s.DiscountAmount = &v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem struct {
	MerchantTotalDiscount *int64                                                                                        `json:"merchant_total_discount,omitempty" xml:"merchant_total_discount,omitempty"`
	PlatformTotalDiscount *int64                                                                                        `json:"platform_total_discount,omitempty" xml:"platform_total_discount,omitempty"`
	TotalAmount           *int64                                                                                        `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	TotalDiscountAmount   *int64                                                                                        `json:"total_discount_amount,omitempty" xml:"total_discount_amount,omitempty"`
	GoodsId               *string                                                                                       `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	MarketingDetailInfo   []*OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem `json:"marketing_detail_info,omitempty" xml:"marketing_detail_info,omitempty" type:"Repeated"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetMerchantTotalDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.MerchantTotalDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetPlatformTotalDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.PlatformTotalDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetTotalAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.TotalAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetTotalDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.TotalDiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetGoodsId(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem) SetMarketingDetailInfo(v []*OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItem {
	s.MarketingDetailInfo = v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem struct {
	MarketingExtend map[string]*string `json:"marketing_extend,omitempty" xml:"marketing_extend,omitempty"`
	Subtype         *string            `json:"subtype,omitempty" xml:"subtype,omitempty"`
	Type            *int32             `json:"type,omitempty" xml:"type,omitempty"`
	DiscountRange   *int32             `json:"discount_range,omitempty" xml:"discount_range,omitempty"`
	Code            *string            `json:"code,omitempty" xml:"code,omitempty"`
	Title           *string            `json:"title,omitempty" xml:"title,omitempty"`
	Kind            *int32             `json:"kind,omitempty" xml:"kind,omitempty"`
	Id              *string            `json:"id,omitempty" xml:"id,omitempty"`
	Value           *int64             `json:"value,omitempty" xml:"value,omitempty"`
	DiscountAmount  *int64             `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	Note            *string            `json:"note,omitempty" xml:"note,omitempty"`
	CreatorType     *int32             `json:"creator_type,omitempty" xml:"creator_type,omitempty"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetMarketingExtend(v map[string]*string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.MarketingExtend = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetSubtype(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Subtype = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetType(v int32) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Type = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetDiscountRange(v int32) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.DiscountRange = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetCode(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Code = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetTitle(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Title = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetKind(v int32) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Kind = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetId(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Id = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetValue(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Value = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.DiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetNote(v string) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.Note = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem) SetCreatorType(v int32) *OrderCreateOrderRequestPriceCalculationDetailItemDiscountDetailItemMarketingDetailInfoItem {
	s.CreatorType = &v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail struct {
	OrderPlatformDiscount    *int64                                                                                     `json:"order_platform_discount,omitempty" xml:"order_platform_discount,omitempty"`
	OrderTotalDiscountAmount *int64                                                                                     `json:"order_total_discount_amount,omitempty" xml:"order_total_discount_amount,omitempty"`
	SkuMerchantDiscount      *int64                                                                                     `json:"sku_merchant_discount,omitempty" xml:"sku_merchant_discount,omitempty"`
	SkuPlatformDiscount      *int64                                                                                     `json:"sku_platform_discount,omitempty" xml:"sku_platform_discount,omitempty"`
	GoodsTotalDiscountAmount *int64                                                                                     `json:"goods_total_discount_amount,omitempty" xml:"goods_total_discount_amount,omitempty"`
	MarketingDetailInfo      []*OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem `json:"marketing_detail_info,omitempty" xml:"marketing_detail_info,omitempty" type:"Repeated"`
	OrderMerchantDiscount    *int64                                                                                     `json:"order_merchant_discount,omitempty" xml:"order_merchant_discount,omitempty"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetOrderPlatformDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.OrderPlatformDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetOrderTotalDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.OrderTotalDiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetSkuMerchantDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.SkuMerchantDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetSkuPlatformDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.SkuPlatformDiscount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetGoodsTotalDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.GoodsTotalDiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetMarketingDetailInfo(v []*OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.MarketingDetailInfo = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail) SetOrderMerchantDiscount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetail {
	s.OrderMerchantDiscount = &v
	return s
}

type OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem struct {
	Kind            *int32             `json:"kind,omitempty" xml:"kind,omitempty"`
	Note            *string            `json:"note,omitempty" xml:"note,omitempty"`
	Subtype         *string            `json:"subtype,omitempty" xml:"subtype,omitempty"`
	Id              *string            `json:"id,omitempty" xml:"id,omitempty"`
	Type            *int32             `json:"type,omitempty" xml:"type,omitempty"`
	DiscountRange   *int32             `json:"discount_range,omitempty" xml:"discount_range,omitempty"`
	Code            *string            `json:"code,omitempty" xml:"code,omitempty"`
	DiscountAmount  *int64             `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	MarketingExtend map[string]*string `json:"marketing_extend,omitempty" xml:"marketing_extend,omitempty"`
	CreatorType     *int32             `json:"creator_type,omitempty" xml:"creator_type,omitempty"`
	Value           *int64             `json:"value,omitempty" xml:"value,omitempty"`
	Title           *string            `json:"title,omitempty" xml:"title,omitempty"`
}

func (s OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetKind(v int32) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Kind = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetNote(v string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Note = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetSubtype(v string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Subtype = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetId(v string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Id = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetType(v int32) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Type = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetDiscountRange(v int32) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.DiscountRange = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetCode(v string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Code = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetDiscountAmount(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.DiscountAmount = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetMarketingExtend(v map[string]*string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.MarketingExtend = v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetCreatorType(v int32) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.CreatorType = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetValue(v int64) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Value = &v
	return s
}

func (s *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem) SetTitle(v string) *OrderCreateOrderRequestPriceCalculationDetailOrderDiscountDetailMarketingDetailInfoItem {
	s.Title = &v
	return s
}

type OrderCreateOrderRequestSkuListItem struct {
	GoodsInfo      *OrderCreateOrderRequestSkuListItemGoodsInfo `json:"goods_info,omitempty" xml:"goods_info,omitempty"`
	Price          *int64                                       `json:"price,omitempty" xml:"price,omitempty"`
	Quantity       *int32                                       `json:"quantity,omitempty" xml:"quantity,omitempty"`
	SkuId          *string                                      `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuIdType      *int32                                       `json:"sku_id_type,omitempty" xml:"sku_id_type,omitempty"`
	Atts           map[string]*string                           `json:"atts,omitempty" xml:"atts,omitempty"`
	DiscountAmount *int64                                       `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
}

func (s OrderCreateOrderRequestSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestSkuListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestSkuListItem) SetGoodsInfo(v *OrderCreateOrderRequestSkuListItemGoodsInfo) *OrderCreateOrderRequestSkuListItem {
	s.GoodsInfo = v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetPrice(v int64) *OrderCreateOrderRequestSkuListItem {
	s.Price = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetQuantity(v int32) *OrderCreateOrderRequestSkuListItem {
	s.Quantity = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetSkuId(v string) *OrderCreateOrderRequestSkuListItem {
	s.SkuId = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetSkuIdType(v int32) *OrderCreateOrderRequestSkuListItem {
	s.SkuIdType = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetAtts(v map[string]*string) *OrderCreateOrderRequestSkuListItem {
	s.Atts = v
	return s
}

func (s *OrderCreateOrderRequestSkuListItem) SetDiscountAmount(v int64) *OrderCreateOrderRequestSkuListItem {
	s.DiscountAmount = &v
	return s
}

type OrderCreateOrderRequestSkuListItemGoodsInfo struct {
	Labels         *string                                                    `json:"labels,omitempty" xml:"labels,omitempty"`
	GoodsId        *string                                                    `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	GoodsBookInfo  *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo  `json:"goods_book_info,omitempty" xml:"goods_book_info,omitempty"`
	MerchantUid    *string                                                    `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty"`
	GoodsIdType    *int32                                                     `json:"goods_id_type,omitempty" xml:"goods_id_type,omitempty"`
	GoodsPage      *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage      `json:"goods_page,omitempty" xml:"goods_page,omitempty"`
	DateRule       *string                                                    `json:"date_rule,omitempty" xml:"date_rule,omitempty"`
	GoodsImage     *string                                                    `json:"goods_image,omitempty" xml:"goods_image,omitempty"`
	GoodsTitle     *string                                                    `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
	OrderValidTime *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime `json:"order_valid_time,omitempty" xml:"order_valid_time,omitempty"`
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetLabels(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.Labels = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsId(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsBookInfo(v *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsBookInfo = v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetMerchantUid(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.MerchantUid = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsIdType(v int32) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsIdType = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsPage(v *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsPage = v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetDateRule(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.DateRule = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsImage(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsImage = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetGoodsTitle(v string) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.GoodsTitle = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfo) SetOrderValidTime(v *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) *OrderCreateOrderRequestSkuListItemGoodsInfo {
	s.OrderValidTime = v
	return s
}

type OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo struct {
	CancelAdvanceHour *int64 `json:"cancel_advance_hour,omitempty" xml:"cancel_advance_hour,omitempty"`
	CancelPolicy      *int32 `json:"cancel_policy,omitempty" xml:"cancel_policy,omitempty"`
	BookType          *int32 `json:"book_type,omitempty" xml:"book_type,omitempty"`
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) SetCancelAdvanceHour(v int64) *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo {
	s.CancelAdvanceHour = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) SetCancelPolicy(v int32) *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo {
	s.CancelPolicy = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo) SetBookType(v int32) *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsBookInfo {
	s.BookType = &v
	return s
}

type OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage) SetParams(v string) *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage {
	s.Params = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage) SetPath(v string) *OrderCreateOrderRequestSkuListItemGoodsInfoGoodsPage {
	s.Path = &v
	return s
}

type OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime struct {
	ValidEndTime   *int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	ValidStartTime *int64 `json:"valid_start_time,omitempty" xml:"valid_start_time,omitempty"`
	ValidDuration  *int64 `json:"valid_duration,omitempty" xml:"valid_duration,omitempty"`
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) SetValidEndTime(v int64) *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime {
	s.ValidEndTime = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) SetValidStartTime(v int64) *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime {
	s.ValidStartTime = &v
	return s
}

func (s *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime) SetValidDuration(v int64) *OrderCreateOrderRequestSkuListItemGoodsInfoOrderValidTime {
	s.ValidDuration = &v
	return s
}

type OrderCreateOrderRequestUserAddress struct {
	UserAddressId   *string                                            `json:"user_address_id,omitempty" xml:"user_address_id,omitempty"`
	ConnectName     *string                                            `json:"connect_name,omitempty" xml:"connect_name,omitempty"`
	DetailAddress   *string                                            `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	DoorPlateNum    *string                                            `json:"door_plate_num,omitempty" xml:"door_plate_num,omitempty"`
	Gender          *int32                                             `json:"gender,omitempty" xml:"gender,omitempty"`
	LocationAddress *OrderCreateOrderRequestUserAddressLocationAddress `json:"location_address,omitempty" xml:"location_address,omitempty"`
	Phone           *string                                            `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s OrderCreateOrderRequestUserAddress) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestUserAddress) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestUserAddress) SetUserAddressId(v string) *OrderCreateOrderRequestUserAddress {
	s.UserAddressId = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetConnectName(v string) *OrderCreateOrderRequestUserAddress {
	s.ConnectName = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetDetailAddress(v string) *OrderCreateOrderRequestUserAddress {
	s.DetailAddress = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetDoorPlateNum(v string) *OrderCreateOrderRequestUserAddress {
	s.DoorPlateNum = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetGender(v int32) *OrderCreateOrderRequestUserAddress {
	s.Gender = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetLocationAddress(v *OrderCreateOrderRequestUserAddressLocationAddress) *OrderCreateOrderRequestUserAddress {
	s.LocationAddress = v
	return s
}

func (s *OrderCreateOrderRequestUserAddress) SetPhone(v string) *OrderCreateOrderRequestUserAddress {
	s.Phone = &v
	return s
}

type OrderCreateOrderRequestUserAddressLocationAddress struct {
	LocationId      *string `json:"location_id,omitempty" xml:"location_id,omitempty"`
	LocationName    *string `json:"location_name,omitempty" xml:"location_name,omitempty"`
	CityCode        *string `json:"city_code,omitempty" xml:"city_code,omitempty" require:"true"`
	LocationAddress *string `json:"location_address,omitempty" xml:"location_address,omitempty"`
	CityName        *string `json:"city_name,omitempty" xml:"city_name,omitempty" require:"true"`
	ProvinceCode    *string `json:"province_code,omitempty" xml:"province_code,omitempty" require:"true"`
	TownName        *string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	Lode            *string `json:"lode,omitempty" xml:"lode,omitempty"`
	DistrictCode    *string `json:"district_code,omitempty" xml:"district_code,omitempty" require:"true"`
	TownCode        *string `json:"town_code,omitempty" xml:"town_code,omitempty"`
	ProvinceName    *string `json:"province_name,omitempty" xml:"province_name,omitempty" require:"true"`
	Lade            *string `json:"lade,omitempty" xml:"lade,omitempty"`
	DistrictName    *string `json:"district_name,omitempty" xml:"district_name,omitempty" require:"true"`
}

func (s OrderCreateOrderRequestUserAddressLocationAddress) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderRequestUserAddressLocationAddress) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetLocationId(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.LocationId = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetLocationName(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.LocationName = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetCityCode(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.CityCode = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetLocationAddress(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.LocationAddress = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetCityName(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.CityName = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetProvinceCode(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.ProvinceCode = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetTownName(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.TownName = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetLode(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.Lode = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetDistrictCode(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.DistrictCode = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetTownCode(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.TownCode = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetProvinceName(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.ProvinceName = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetLade(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.Lade = &v
	return s
}

func (s *OrderCreateOrderRequestUserAddressLocationAddress) SetDistrictName(v string) *OrderCreateOrderRequestUserAddressLocationAddress {
	s.DistrictName = &v
	return s
}

type OrderCreateOrderResponse struct {
	Extra *OrderCreateOrderResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *OrderCreateOrderResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrderCreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponse) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponse) SetExtra(v *OrderCreateOrderResponseExtra) *OrderCreateOrderResponse {
	s.Extra = v
	return s
}

func (s *OrderCreateOrderResponse) SetData(v *OrderCreateOrderResponseData) *OrderCreateOrderResponse {
	s.Data = v
	return s
}

type OrderCreateOrderResponseData struct {
	GwDescription       *string                                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OpenBookInfo        *OrderCreateOrderResponseDataOpenBookInfo              `json:"open_book_info,omitempty" xml:"open_book_info,omitempty"`
	ItemOrderDetailList []*OrderCreateOrderResponseDataItemOrderDetailListItem `json:"item_order_detail_list,omitempty" xml:"item_order_detail_list,omitempty" type:"Repeated"`
	OrderId             *string                                                `json:"order_id,omitempty" xml:"order_id,omitempty"`
	GwErrorCode         *int32                                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	PayOrderToken       *string                                                `json:"pay_order_token,omitempty" xml:"pay_order_token,omitempty"`
	OutOrderNo          *string                                                `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	PayOrderId          *string                                                `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	ItemOrderInfoList   []*OrderCreateOrderResponseDataItemOrderInfoListItem   `json:"item_order_info_list,omitempty" xml:"item_order_info_list,omitempty" type:"Repeated"`
}

func (s OrderCreateOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseData) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseData) SetGwDescription(v string) *OrderCreateOrderResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetOpenBookInfo(v *OrderCreateOrderResponseDataOpenBookInfo) *OrderCreateOrderResponseData {
	s.OpenBookInfo = v
	return s
}

func (s *OrderCreateOrderResponseData) SetItemOrderDetailList(v []*OrderCreateOrderResponseDataItemOrderDetailListItem) *OrderCreateOrderResponseData {
	s.ItemOrderDetailList = v
	return s
}

func (s *OrderCreateOrderResponseData) SetOrderId(v string) *OrderCreateOrderResponseData {
	s.OrderId = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetGwErrorCode(v int32) *OrderCreateOrderResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetPayOrderToken(v string) *OrderCreateOrderResponseData {
	s.PayOrderToken = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetOutOrderNo(v string) *OrderCreateOrderResponseData {
	s.OutOrderNo = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetPayOrderId(v string) *OrderCreateOrderResponseData {
	s.PayOrderId = &v
	return s
}

func (s *OrderCreateOrderResponseData) SetItemOrderInfoList(v []*OrderCreateOrderResponseDataItemOrderInfoListItem) *OrderCreateOrderResponseData {
	s.ItemOrderInfoList = v
	return s
}

type OrderCreateOrderResponseDataItemOrderDetailListItem struct {
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuIdType   *int32  `json:"sku_id_type,omitempty" xml:"sku_id_type,omitempty"`
	GoodsId     *string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	GoodsIdType *int32  `json:"goods_id_type,omitempty" xml:"goods_id_type,omitempty"`
	ItemOrderId *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s OrderCreateOrderResponseDataItemOrderDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseDataItemOrderDetailListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetSkuId(v string) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.SkuId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetSkuIdType(v int32) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.SkuIdType = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetGoodsId(v string) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetGoodsIdType(v int32) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.GoodsIdType = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetItemOrderId(v string) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.ItemOrderId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderDetailListItem) SetPrice(v int64) *OrderCreateOrderResponseDataItemOrderDetailListItem {
	s.Price = &v
	return s
}

type OrderCreateOrderResponseDataItemOrderInfoListItem struct {
	GoodsId         *string                                                                 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	ItemOrderDetail []*OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem `json:"item_order_detail,omitempty" xml:"item_order_detail,omitempty" type:"Repeated"`
	ItemOrderIdList []*string                                                               `json:"item_order_id_list,omitempty" xml:"item_order_id_list,omitempty" type:"Repeated"`
	SkuId           *string                                                                 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s OrderCreateOrderResponseDataItemOrderInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseDataItemOrderInfoListItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItem) SetGoodsId(v string) *OrderCreateOrderResponseDataItemOrderInfoListItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItem) SetItemOrderDetail(v []*OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) *OrderCreateOrderResponseDataItemOrderInfoListItem {
	s.ItemOrderDetail = v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItem) SetItemOrderIdList(v []*string) *OrderCreateOrderResponseDataItemOrderInfoListItem {
	s.ItemOrderIdList = v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItem) SetSkuId(v string) *OrderCreateOrderResponseDataItemOrderInfoListItem {
	s.SkuId = &v
	return s
}

type OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem struct {
	ItemOrderId *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuIdType   *int32  `json:"sku_id_type,omitempty" xml:"sku_id_type,omitempty"`
	GoodsId     *string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	GoodsIdType *int32  `json:"goods_id_type,omitempty" xml:"goods_id_type,omitempty"`
}

func (s OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetItemOrderId(v string) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.ItemOrderId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetPrice(v int64) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.Price = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetSkuId(v string) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.SkuId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetSkuIdType(v int32) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.SkuIdType = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetGoodsId(v string) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.GoodsId = &v
	return s
}

func (s *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem) SetGoodsIdType(v int32) *OrderCreateOrderResponseDataItemOrderInfoListItemItemOrderDetailItem {
	s.GoodsIdType = &v
	return s
}

type OrderCreateOrderResponseDataOpenBookInfo struct {
	BookId *string `json:"book_id,omitempty" xml:"book_id,omitempty"`
}

func (s OrderCreateOrderResponseDataOpenBookInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseDataOpenBookInfo) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseDataOpenBookInfo) SetBookId(v string) *OrderCreateOrderResponseDataOpenBookInfo {
	s.BookId = &v
	return s
}

type OrderCreateOrderResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s OrderCreateOrderResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateOrderResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderCreateOrderResponseExtra) SetSubErrorCode(v int32) *OrderCreateOrderResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderCreateOrderResponseExtra) SetDescription(v string) *OrderCreateOrderResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderCreateOrderResponseExtra) SetErrorCode(v int32) *OrderCreateOrderResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderCreateOrderResponseExtra) SetLogid(v string) *OrderCreateOrderResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderCreateOrderResponseExtra) SetNow(v int64) *OrderCreateOrderResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderCreateOrderResponseExtra) SetSubDescription(v string) *OrderCreateOrderResponseExtra {
	s.SubDescription = &v
	return s
}

type OrderCreateRequest struct {
	CooperationContent *int               `json:"cooperation_content,omitempty" xml:"cooperation_content,omitempty" require:"true"`
	EndTime            *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	StartTime          *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId          *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ChargeType         *int               `json:"charge_type,omitempty" xml:"charge_type,omitempty" require:"true"`
	CommissionRatio    *string            `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
	GoodsChannel       *int               `json:"goods_channel,omitempty" xml:"goods_channel,omitempty"`
	Header             map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s OrderCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateRequest) GoString() string {
	return s.String()
}

func (s *OrderCreateRequest) SetCooperationContent(v int) *OrderCreateRequest {
	s.CooperationContent = &v
	return s
}

func (s *OrderCreateRequest) SetEndTime(v int64) *OrderCreateRequest {
	s.EndTime = &v
	return s
}

func (s *OrderCreateRequest) SetStartTime(v int64) *OrderCreateRequest {
	s.StartTime = &v
	return s
}

func (s *OrderCreateRequest) SetAccessToken(v string) *OrderCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderCreateRequest) SetAccountId(v string) *OrderCreateRequest {
	s.AccountId = &v
	return s
}

func (s *OrderCreateRequest) SetChargeType(v int) *OrderCreateRequest {
	s.ChargeType = &v
	return s
}

func (s *OrderCreateRequest) SetCommissionRatio(v string) *OrderCreateRequest {
	s.CommissionRatio = &v
	return s
}

func (s *OrderCreateRequest) SetGoodsChannel(v int) *OrderCreateRequest {
	s.GoodsChannel = &v
	return s
}

func (s *OrderCreateRequest) SetHeader(v map[string]*string) *OrderCreateRequest {
	s.Header = v
	return s
}

type OrderCreateResponse struct {
	Data  *OrderCreateResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *OrderCreateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s OrderCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateResponse) GoString() string {
	return s.String()
}

func (s *OrderCreateResponse) SetData(v *OrderCreateResponseData) *OrderCreateResponse {
	s.Data = v
	return s
}

func (s *OrderCreateResponse) SetExtra(v *OrderCreateResponseExtra) *OrderCreateResponse {
	s.Extra = v
	return s
}

type OrderCreateResponseData struct {
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateResponseData) GoString() string {
	return s.String()
}

func (s *OrderCreateResponseData) SetOrderId(v string) *OrderCreateResponseData {
	s.OrderId = &v
	return s
}

func (s *OrderCreateResponseData) SetGwErrorCode(v int32) *OrderCreateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderCreateResponseData) SetGwDescription(v string) *OrderCreateResponseData {
	s.GwDescription = &v
	return s
}

type OrderCreateResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s OrderCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderCreateResponseExtra) SetSubDescription(v string) *OrderCreateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderCreateResponseExtra) SetSubErrorCode(v int32) *OrderCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderCreateResponseExtra) SetDescription(v string) *OrderCreateResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderCreateResponseExtra) SetErrorCode(v int32) *OrderCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderCreateResponseExtra) SetLogid(v string) *OrderCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderCreateResponseExtra) SetNow(v int64) *OrderCreateResponseExtra {
	s.Now = &v
	return s
}

type OrderGetRequest struct {
	Header              map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken         *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId             *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	WithoutProductItems *string            `json:"without_product_items,omitempty" xml:"without_product_items,omitempty"`
}

func (s OrderGetRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderGetRequest) GoString() string {
	return s.String()
}

func (s *OrderGetRequest) SetHeader(v map[string]*string) *OrderGetRequest {
	s.Header = v
	return s
}

func (s *OrderGetRequest) SetAccessToken(v string) *OrderGetRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderGetRequest) SetOrderId(v string) *OrderGetRequest {
	s.OrderId = &v
	return s
}

func (s *OrderGetRequest) SetWithoutProductItems(v string) *OrderGetRequest {
	s.WithoutProductItems = &v
	return s
}

type OrderGetResponse struct {
	Data  *OrderGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *OrderGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s OrderGetResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderGetResponse) GoString() string {
	return s.String()
}

func (s *OrderGetResponse) SetData(v *OrderGetResponseData) *OrderGetResponse {
	s.Data = v
	return s
}

func (s *OrderGetResponse) SetExtra(v *OrderGetResponseExtra) *OrderGetResponse {
	s.Extra = v
	return s
}

type OrderGetResponseData struct {
	GoodsChannel       *int                                    `json:"goods_channel,omitempty" xml:"goods_channel,omitempty"`
	Status             *int                                    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	ProductItems       []*OrderGetResponseDataProductItemsItem `json:"product_items,omitempty" xml:"product_items,omitempty" type:"Repeated"`
	ProductName        *string                                 `json:"product_name,omitempty" xml:"product_name,omitempty"`
	StartTime          *int64                                  `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	GwErrorCode        *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	EndTime            *int64                                  `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	CommissionRatio    *string                                 `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
	CooperationContent *int                                    `json:"cooperation_content,omitempty" xml:"cooperation_content,omitempty" require:"true"`
	MerchantName       *string                                 `json:"merchant_name,omitempty" xml:"merchant_name,omitempty" require:"true"`
	CreateTime         *int64                                  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	AccountId          *string                                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Id                 *string                                 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	GwDescription      *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ChargeType         *int                                    `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
}

func (s OrderGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderGetResponseData) GoString() string {
	return s.String()
}

func (s *OrderGetResponseData) SetGoodsChannel(v int) *OrderGetResponseData {
	s.GoodsChannel = &v
	return s
}

func (s *OrderGetResponseData) SetStatus(v int) *OrderGetResponseData {
	s.Status = &v
	return s
}

func (s *OrderGetResponseData) SetProductItems(v []*OrderGetResponseDataProductItemsItem) *OrderGetResponseData {
	s.ProductItems = v
	return s
}

func (s *OrderGetResponseData) SetProductName(v string) *OrderGetResponseData {
	s.ProductName = &v
	return s
}

func (s *OrderGetResponseData) SetStartTime(v int64) *OrderGetResponseData {
	s.StartTime = &v
	return s
}

func (s *OrderGetResponseData) SetGwErrorCode(v int32) *OrderGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderGetResponseData) SetEndTime(v int64) *OrderGetResponseData {
	s.EndTime = &v
	return s
}

func (s *OrderGetResponseData) SetCommissionRatio(v string) *OrderGetResponseData {
	s.CommissionRatio = &v
	return s
}

func (s *OrderGetResponseData) SetCooperationContent(v int) *OrderGetResponseData {
	s.CooperationContent = &v
	return s
}

func (s *OrderGetResponseData) SetMerchantName(v string) *OrderGetResponseData {
	s.MerchantName = &v
	return s
}

func (s *OrderGetResponseData) SetCreateTime(v int64) *OrderGetResponseData {
	s.CreateTime = &v
	return s
}

func (s *OrderGetResponseData) SetAccountId(v string) *OrderGetResponseData {
	s.AccountId = &v
	return s
}

func (s *OrderGetResponseData) SetId(v string) *OrderGetResponseData {
	s.Id = &v
	return s
}

func (s *OrderGetResponseData) SetGwDescription(v string) *OrderGetResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderGetResponseData) SetChargeType(v int) *OrderGetResponseData {
	s.ChargeType = &v
	return s
}

type OrderGetResponseDataProductItemsItem struct {
	ActualPrice     *string `json:"actual_price,omitempty" xml:"actual_price,omitempty" require:"true"`
	CommissionRatio *string `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty" require:"true"`
	ProductId       *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	ProductStatus   *int    `json:"product_status,omitempty" xml:"product_status,omitempty" require:"true"`
}

func (s OrderGetResponseDataProductItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderGetResponseDataProductItemsItem) GoString() string {
	return s.String()
}

func (s *OrderGetResponseDataProductItemsItem) SetActualPrice(v string) *OrderGetResponseDataProductItemsItem {
	s.ActualPrice = &v
	return s
}

func (s *OrderGetResponseDataProductItemsItem) SetCommissionRatio(v string) *OrderGetResponseDataProductItemsItem {
	s.CommissionRatio = &v
	return s
}

func (s *OrderGetResponseDataProductItemsItem) SetProductId(v string) *OrderGetResponseDataProductItemsItem {
	s.ProductId = &v
	return s
}

func (s *OrderGetResponseDataProductItemsItem) SetProductStatus(v int) *OrderGetResponseDataProductItemsItem {
	s.ProductStatus = &v
	return s
}

type OrderGetResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderGetResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderGetResponseExtra) SetErrorCode(v int32) *OrderGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderGetResponseExtra) SetLogid(v string) *OrderGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderGetResponseExtra) SetNow(v int64) *OrderGetResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderGetResponseExtra) SetSubDescription(v string) *OrderGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderGetResponseExtra) SetSubErrorCode(v int32) *OrderGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderGetResponseExtra) SetDescription(v string) *OrderGetResponseExtra {
	s.Description = &v
	return s
}

type OrderMerchantRejectRequest struct {
	OrderId      *string                                 `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header       map[string]*string                      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RejectReason *OrderMerchantRejectRequestRejectReason `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s OrderMerchantRejectRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderMerchantRejectRequest) GoString() string {
	return s.String()
}

func (s *OrderMerchantRejectRequest) SetOrderId(v string) *OrderMerchantRejectRequest {
	s.OrderId = &v
	return s
}

func (s *OrderMerchantRejectRequest) SetHeader(v map[string]*string) *OrderMerchantRejectRequest {
	s.Header = v
	return s
}

func (s *OrderMerchantRejectRequest) SetAccessToken(v string) *OrderMerchantRejectRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderMerchantRejectRequest) SetRejectReason(v *OrderMerchantRejectRequestRejectReason) *OrderMerchantRejectRequest {
	s.RejectReason = v
	return s
}

type OrderMerchantRejectRequestRejectReason struct {
	ReasonCode []*int64 `json:"reason_code,omitempty" xml:"reason_code,omitempty" require:"true" type:"Repeated"`
	Desc       *string  `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s OrderMerchantRejectRequestRejectReason) String() string {
	return tea.Prettify(s)
}

func (s OrderMerchantRejectRequestRejectReason) GoString() string {
	return s.String()
}

func (s *OrderMerchantRejectRequestRejectReason) SetReasonCode(v []*int64) *OrderMerchantRejectRequestRejectReason {
	s.ReasonCode = v
	return s
}

func (s *OrderMerchantRejectRequestRejectReason) SetDesc(v string) *OrderMerchantRejectRequestRejectReason {
	s.Desc = &v
	return s
}

type OrderMerchantRejectResponse struct {
	Extra *OrderMerchantRejectResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *OrderMerchantRejectResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrderMerchantRejectResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderMerchantRejectResponse) GoString() string {
	return s.String()
}

func (s *OrderMerchantRejectResponse) SetExtra(v *OrderMerchantRejectResponseExtra) *OrderMerchantRejectResponse {
	s.Extra = v
	return s
}

func (s *OrderMerchantRejectResponse) SetData(v *OrderMerchantRejectResponseData) *OrderMerchantRejectResponse {
	s.Data = v
	return s
}

type OrderMerchantRejectResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s OrderMerchantRejectResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderMerchantRejectResponseData) GoString() string {
	return s.String()
}

func (s *OrderMerchantRejectResponseData) SetGwDescription(v string) *OrderMerchantRejectResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderMerchantRejectResponseData) SetGwErrorCode(v int32) *OrderMerchantRejectResponseData {
	s.GwErrorCode = &v
	return s
}

type OrderMerchantRejectResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s OrderMerchantRejectResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderMerchantRejectResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderMerchantRejectResponseExtra) SetSubErrorCode(v int32) *OrderMerchantRejectResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderMerchantRejectResponseExtra) SetDescription(v string) *OrderMerchantRejectResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderMerchantRejectResponseExtra) SetErrorCode(v int32) *OrderMerchantRejectResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderMerchantRejectResponseExtra) SetLogid(v string) *OrderMerchantRejectResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderMerchantRejectResponseExtra) SetNow(v int64) *OrderMerchantRejectResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderMerchantRejectResponseExtra) SetSubDescription(v string) *OrderMerchantRejectResponseExtra {
	s.SubDescription = &v
	return s
}

type OrderQueryCpsRequest struct {
	OutOrderNo  *string            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s OrderQueryCpsRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsRequest) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsRequest) SetOutOrderNo(v string) *OrderQueryCpsRequest {
	s.OutOrderNo = &v
	return s
}

func (s *OrderQueryCpsRequest) SetHeader(v map[string]*string) *OrderQueryCpsRequest {
	s.Header = v
	return s
}

func (s *OrderQueryCpsRequest) SetAccessToken(v string) *OrderQueryCpsRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderQueryCpsRequest) SetOrderId(v string) *OrderQueryCpsRequest {
	s.OrderId = &v
	return s
}

type OrderQueryCpsResponse struct {
	Extra *OrderQueryCpsResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *OrderQueryCpsResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrderQueryCpsResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponse) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponse) SetExtra(v *OrderQueryCpsResponseExtra) *OrderQueryCpsResponse {
	s.Extra = v
	return s
}

func (s *OrderQueryCpsResponse) SetData(v *OrderQueryCpsResponseData) *OrderQueryCpsResponse {
	s.Data = v
	return s
}

type OrderQueryCpsResponseData struct {
	OutOrderNo    *string                                `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	PaymentInfo   *OrderQueryCpsResponseDataPaymentInfo  `json:"payment_info,omitempty" xml:"payment_info,omitempty"`
	RefundInfo    *OrderQueryCpsResponseDataRefundInfo   `json:"refund_info,omitempty" xml:"refund_info,omitempty"`
	CpsInfo       *OrderQueryCpsResponseDataCpsInfo      `json:"cps_info,omitempty" xml:"cps_info,omitempty"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	DeliveryInfo  *OrderQueryCpsResponseDataDeliveryInfo `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
	OrderId       *string                                `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s OrderQueryCpsResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseData) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseData) SetOutOrderNo(v string) *OrderQueryCpsResponseData {
	s.OutOrderNo = &v
	return s
}

func (s *OrderQueryCpsResponseData) SetPaymentInfo(v *OrderQueryCpsResponseDataPaymentInfo) *OrderQueryCpsResponseData {
	s.PaymentInfo = v
	return s
}

func (s *OrderQueryCpsResponseData) SetRefundInfo(v *OrderQueryCpsResponseDataRefundInfo) *OrderQueryCpsResponseData {
	s.RefundInfo = v
	return s
}

func (s *OrderQueryCpsResponseData) SetCpsInfo(v *OrderQueryCpsResponseDataCpsInfo) *OrderQueryCpsResponseData {
	s.CpsInfo = v
	return s
}

func (s *OrderQueryCpsResponseData) SetGwErrorCode(v int32) *OrderQueryCpsResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderQueryCpsResponseData) SetGwDescription(v string) *OrderQueryCpsResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderQueryCpsResponseData) SetDeliveryInfo(v *OrderQueryCpsResponseDataDeliveryInfo) *OrderQueryCpsResponseData {
	s.DeliveryInfo = v
	return s
}

func (s *OrderQueryCpsResponseData) SetOrderId(v string) *OrderQueryCpsResponseData {
	s.OrderId = &v
	return s
}

type OrderQueryCpsResponseDataCpsInfo struct {
	CpsItemList           []*OrderQueryCpsResponseDataCpsInfoCpsItemListItem `json:"cps_item_list,omitempty" xml:"cps_item_list,omitempty" type:"Repeated"`
	TotalCommissionAmount *int64                                             `json:"total_commission_amount,omitempty" xml:"total_commission_amount,omitempty"`
}

func (s OrderQueryCpsResponseDataCpsInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataCpsInfo) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataCpsInfo) SetCpsItemList(v []*OrderQueryCpsResponseDataCpsInfoCpsItemListItem) *OrderQueryCpsResponseDataCpsInfo {
	s.CpsItemList = v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfo) SetTotalCommissionAmount(v int64) *OrderQueryCpsResponseDataCpsInfo {
	s.TotalCommissionAmount = &v
	return s
}

type OrderQueryCpsResponseDataCpsInfoCpsItemListItem struct {
	TaskId                 *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	ItemId                 *int64  `json:"item_id,omitempty" xml:"item_id,omitempty"`
	CommissionUserDouyinid *string `json:"commission_user_douyinid,omitempty" xml:"commission_user_douyinid,omitempty"`
	CommissionUserNickname *string `json:"commission_user_nickname,omitempty" xml:"commission_user_nickname,omitempty"`
	SellAmount             *int64  `json:"sell_amount,omitempty" xml:"sell_amount,omitempty"`
	CommissionAmount       *int64  `json:"commission_amount,omitempty" xml:"commission_amount,omitempty"`
	Status                 *int64  `json:"status,omitempty" xml:"status,omitempty"`
	ItemOrderId            *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	SourceType             *int32  `json:"source_type,omitempty" xml:"source_type,omitempty"`
	CommissionRate         *int32  `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
}

func (s OrderQueryCpsResponseDataCpsInfoCpsItemListItem) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataCpsInfoCpsItemListItem) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetTaskId(v string) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.TaskId = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetItemId(v int64) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.ItemId = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetCommissionUserDouyinid(v string) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.CommissionUserDouyinid = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetCommissionUserNickname(v string) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.CommissionUserNickname = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetSellAmount(v int64) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.SellAmount = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetCommissionAmount(v int64) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.CommissionAmount = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetStatus(v int64) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.Status = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetItemOrderId(v string) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.ItemOrderId = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetSourceType(v int32) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.SourceType = &v
	return s
}

func (s *OrderQueryCpsResponseDataCpsInfoCpsItemListItem) SetCommissionRate(v int32) *OrderQueryCpsResponseDataCpsInfoCpsItemListItem {
	s.CommissionRate = &v
	return s
}

type OrderQueryCpsResponseDataDeliveryInfo struct {
	DeliveryItems       []*OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem `json:"delivery_items,omitempty" xml:"delivery_items,omitempty" type:"Repeated"`
	TotalDeliveryAmount *int64                                                    `json:"total_delivery_amount,omitempty" xml:"total_delivery_amount,omitempty"`
}

func (s OrderQueryCpsResponseDataDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataDeliveryInfo) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataDeliveryInfo) SetDeliveryItems(v []*OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) *OrderQueryCpsResponseDataDeliveryInfo {
	s.DeliveryItems = v
	return s
}

func (s *OrderQueryCpsResponseDataDeliveryInfo) SetTotalDeliveryAmount(v int64) *OrderQueryCpsResponseDataDeliveryInfo {
	s.TotalDeliveryAmount = &v
	return s
}

type OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem struct {
	DeliveryAmount *int64  `json:"delivery_amount,omitempty" xml:"delivery_amount,omitempty"`
	DeliveryAt     *int64  `json:"delivery_at,omitempty" xml:"delivery_at,omitempty"`
	DeliveryStatus *string `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	ItemOrderId    *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
}

func (s OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) SetDeliveryAmount(v int64) *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem {
	s.DeliveryAmount = &v
	return s
}

func (s *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) SetDeliveryAt(v int64) *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem {
	s.DeliveryAt = &v
	return s
}

func (s *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) SetDeliveryStatus(v string) *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem {
	s.DeliveryStatus = &v
	return s
}

func (s *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem) SetItemOrderId(v string) *OrderQueryCpsResponseDataDeliveryInfoDeliveryItemsItem {
	s.ItemOrderId = &v
	return s
}

type OrderQueryCpsResponseDataPaymentInfo struct {
	TotalFee     *int64  `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	ChannelPayId *string `json:"channel_pay_id,omitempty" xml:"channel_pay_id,omitempty"`
	PayTime      *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	SellerUid    *string `json:"seller_uid,omitempty" xml:"seller_uid,omitempty"`
	ItemId       *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	Message      *string `json:"message,omitempty" xml:"message,omitempty"`
	CpExtra      *string `json:"cp_extra,omitempty" xml:"cp_extra,omitempty"`
	OrderStatus  *string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	PayChannel   *int32  `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
}

func (s OrderQueryCpsResponseDataPaymentInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataPaymentInfo) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetTotalFee(v int64) *OrderQueryCpsResponseDataPaymentInfo {
	s.TotalFee = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetChannelPayId(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.ChannelPayId = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetPayTime(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.PayTime = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetSellerUid(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.SellerUid = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetItemId(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.ItemId = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetMessage(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.Message = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetCpExtra(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.CpExtra = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetOrderStatus(v string) *OrderQueryCpsResponseDataPaymentInfo {
	s.OrderStatus = &v
	return s
}

func (s *OrderQueryCpsResponseDataPaymentInfo) SetPayChannel(v int32) *OrderQueryCpsResponseDataPaymentInfo {
	s.PayChannel = &v
	return s
}

type OrderQueryCpsResponseDataRefundInfo struct {
	RefundItems       []*OrderQueryCpsResponseDataRefundInfoRefundItemsItem `json:"refund_items,omitempty" xml:"refund_items,omitempty" type:"Repeated"`
	TotalRefundAmount *int64                                                `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
}

func (s OrderQueryCpsResponseDataRefundInfo) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataRefundInfo) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataRefundInfo) SetRefundItems(v []*OrderQueryCpsResponseDataRefundInfoRefundItemsItem) *OrderQueryCpsResponseDataRefundInfo {
	s.RefundItems = v
	return s
}

func (s *OrderQueryCpsResponseDataRefundInfo) SetTotalRefundAmount(v int64) *OrderQueryCpsResponseDataRefundInfo {
	s.TotalRefundAmount = &v
	return s
}

type OrderQueryCpsResponseDataRefundInfoRefundItemsItem struct {
	RefundAt     *int64  `json:"refund_at,omitempty" xml:"refund_at,omitempty"`
	RefundStatus *string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	ItemOrderId  *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	OutRefundNo  *string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	RefundAmount *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

func (s OrderQueryCpsResponseDataRefundInfoRefundItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseDataRefundInfoRefundItemsItem) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseDataRefundInfoRefundItemsItem) SetRefundAt(v int64) *OrderQueryCpsResponseDataRefundInfoRefundItemsItem {
	s.RefundAt = &v
	return s
}

func (s *OrderQueryCpsResponseDataRefundInfoRefundItemsItem) SetRefundStatus(v string) *OrderQueryCpsResponseDataRefundInfoRefundItemsItem {
	s.RefundStatus = &v
	return s
}

func (s *OrderQueryCpsResponseDataRefundInfoRefundItemsItem) SetItemOrderId(v string) *OrderQueryCpsResponseDataRefundInfoRefundItemsItem {
	s.ItemOrderId = &v
	return s
}

func (s *OrderQueryCpsResponseDataRefundInfoRefundItemsItem) SetOutRefundNo(v string) *OrderQueryCpsResponseDataRefundInfoRefundItemsItem {
	s.OutRefundNo = &v
	return s
}

func (s *OrderQueryCpsResponseDataRefundInfoRefundItemsItem) SetRefundAmount(v int64) *OrderQueryCpsResponseDataRefundInfoRefundItemsItem {
	s.RefundAmount = &v
	return s
}

type OrderQueryCpsResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderQueryCpsResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryCpsResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderQueryCpsResponseExtra) SetErrorCode(v int32) *OrderQueryCpsResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderQueryCpsResponseExtra) SetLogid(v string) *OrderQueryCpsResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderQueryCpsResponseExtra) SetNow(v int64) *OrderQueryCpsResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderQueryCpsResponseExtra) SetSubDescription(v string) *OrderQueryCpsResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderQueryCpsResponseExtra) SetSubErrorCode(v int32) *OrderQueryCpsResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderQueryCpsResponseExtra) SetDescription(v string) *OrderQueryCpsResponseExtra {
	s.Description = &v
	return s
}

type OrderQueryRequest struct {
	AccountId           *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	StartTime           *int64             `json:"start_time,omitempty" xml:"start_time,omitempty"`
	AccessToken         *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	GoodsChannel        *int               `json:"goods_channel,omitempty" xml:"goods_channel,omitempty"`
	Status              []*int             `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	IsAsc               *bool              `json:"is_asc,omitempty" xml:"is_asc,omitempty"`
	CooperationContents []*int             `json:"cooperation_contents,omitempty" xml:"cooperation_contents,omitempty" type:"Repeated"`
	Page                *int32             `json:"page,omitempty" xml:"page,omitempty"`
	Header              map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	Size                *int32             `json:"size,omitempty" xml:"size,omitempty"`
	EndTime             *int64             `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s OrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryRequest) GoString() string {
	return s.String()
}

func (s *OrderQueryRequest) SetAccountId(v string) *OrderQueryRequest {
	s.AccountId = &v
	return s
}

func (s *OrderQueryRequest) SetStartTime(v int64) *OrderQueryRequest {
	s.StartTime = &v
	return s
}

func (s *OrderQueryRequest) SetAccessToken(v string) *OrderQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderQueryRequest) SetGoodsChannel(v int) *OrderQueryRequest {
	s.GoodsChannel = &v
	return s
}

func (s *OrderQueryRequest) SetStatus(v []*int) *OrderQueryRequest {
	s.Status = v
	return s
}

func (s *OrderQueryRequest) SetIsAsc(v bool) *OrderQueryRequest {
	s.IsAsc = &v
	return s
}

func (s *OrderQueryRequest) SetCooperationContents(v []*int) *OrderQueryRequest {
	s.CooperationContents = v
	return s
}

func (s *OrderQueryRequest) SetPage(v int32) *OrderQueryRequest {
	s.Page = &v
	return s
}

func (s *OrderQueryRequest) SetHeader(v map[string]*string) *OrderQueryRequest {
	s.Header = v
	return s
}

func (s *OrderQueryRequest) SetSize(v int32) *OrderQueryRequest {
	s.Size = &v
	return s
}

func (s *OrderQueryRequest) SetEndTime(v int64) *OrderQueryRequest {
	s.EndTime = &v
	return s
}

type OrderQueryResponse struct {
	Extra *OrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *OrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s OrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryResponse) GoString() string {
	return s.String()
}

func (s *OrderQueryResponse) SetExtra(v *OrderQueryResponseExtra) *OrderQueryResponse {
	s.Extra = v
	return s
}

func (s *OrderQueryResponse) SetData(v *OrderQueryResponseData) *OrderQueryResponse {
	s.Data = v
	return s
}

type OrderQueryResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Orders        []*OrderQueryResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
	Total         *int32                              `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s OrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *OrderQueryResponseData) SetGwErrorCode(v int32) *OrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderQueryResponseData) SetGwDescription(v string) *OrderQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *OrderQueryResponseData) SetOrders(v []*OrderQueryResponseDataOrdersItem) *OrderQueryResponseData {
	s.Orders = v
	return s
}

func (s *OrderQueryResponseData) SetTotal(v int32) *OrderQueryResponseData {
	s.Total = &v
	return s
}

type OrderQueryResponseDataOrdersItem struct {
	EndTime            *int64                                              `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	ProductName        *string                                             `json:"product_name,omitempty" xml:"product_name,omitempty"`
	AccountId          *string                                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	MerchantName       *string                                             `json:"merchant_name,omitempty" xml:"merchant_name,omitempty" require:"true"`
	Status             *int                                                `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CooperationContent *int                                                `json:"cooperation_content,omitempty" xml:"cooperation_content,omitempty" require:"true"`
	ProductItems       []*OrderQueryResponseDataOrdersItemProductItemsItem `json:"product_items,omitempty" xml:"product_items,omitempty" type:"Repeated"`
	GoodsChannel       *int                                                `json:"goods_channel,omitempty" xml:"goods_channel,omitempty"`
	ChargeType         *int                                                `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	Id                 *string                                             `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	CreateTime         *int64                                              `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	StartTime          *int64                                              `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	CommissionRatio    *string                                             `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
}

func (s OrderQueryResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *OrderQueryResponseDataOrdersItem) SetEndTime(v int64) *OrderQueryResponseDataOrdersItem {
	s.EndTime = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetProductName(v string) *OrderQueryResponseDataOrdersItem {
	s.ProductName = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetAccountId(v string) *OrderQueryResponseDataOrdersItem {
	s.AccountId = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetMerchantName(v string) *OrderQueryResponseDataOrdersItem {
	s.MerchantName = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetStatus(v int) *OrderQueryResponseDataOrdersItem {
	s.Status = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetCooperationContent(v int) *OrderQueryResponseDataOrdersItem {
	s.CooperationContent = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetProductItems(v []*OrderQueryResponseDataOrdersItemProductItemsItem) *OrderQueryResponseDataOrdersItem {
	s.ProductItems = v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetGoodsChannel(v int) *OrderQueryResponseDataOrdersItem {
	s.GoodsChannel = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetChargeType(v int) *OrderQueryResponseDataOrdersItem {
	s.ChargeType = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetId(v string) *OrderQueryResponseDataOrdersItem {
	s.Id = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetCreateTime(v int64) *OrderQueryResponseDataOrdersItem {
	s.CreateTime = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetStartTime(v int64) *OrderQueryResponseDataOrdersItem {
	s.StartTime = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItem) SetCommissionRatio(v string) *OrderQueryResponseDataOrdersItem {
	s.CommissionRatio = &v
	return s
}

type OrderQueryResponseDataOrdersItemProductItemsItem struct {
	ActualPrice     *string `json:"actual_price,omitempty" xml:"actual_price,omitempty" require:"true"`
	CommissionRatio *string `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty" require:"true"`
	ProductId       *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	ProductStatus   *int    `json:"product_status,omitempty" xml:"product_status,omitempty" require:"true"`
}

func (s OrderQueryResponseDataOrdersItemProductItemsItem) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryResponseDataOrdersItemProductItemsItem) GoString() string {
	return s.String()
}

func (s *OrderQueryResponseDataOrdersItemProductItemsItem) SetActualPrice(v string) *OrderQueryResponseDataOrdersItemProductItemsItem {
	s.ActualPrice = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItemProductItemsItem) SetCommissionRatio(v string) *OrderQueryResponseDataOrdersItemProductItemsItem {
	s.CommissionRatio = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItemProductItemsItem) SetProductId(v string) *OrderQueryResponseDataOrdersItemProductItemsItem {
	s.ProductId = &v
	return s
}

func (s *OrderQueryResponseDataOrdersItemProductItemsItem) SetProductStatus(v int) *OrderQueryResponseDataOrdersItemProductItemsItem {
	s.ProductStatus = &v
	return s
}

type OrderQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s OrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderQueryResponseExtra) SetSubDescription(v string) *OrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *OrderQueryResponseExtra) SetSubErrorCode(v int32) *OrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderQueryResponseExtra) SetDescription(v string) *OrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderQueryResponseExtra) SetErrorCode(v int32) *OrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderQueryResponseExtra) SetLogid(v string) *OrderQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderQueryResponseExtra) SetNow(v int64) *OrderQueryResponseExtra {
	s.Now = &v
	return s
}

type OrderSyncStatusRequest struct {
	RiderPhone      *string            `json:"rider_phone,omitempty" xml:"rider_phone,omitempty"`
	AcceptTime      *int64             `json:"accept_time,omitempty" xml:"accept_time,omitempty"`
	RiderPhoneType  *int64             `json:"rider_phone_type,omitempty" xml:"rider_phone_type,omitempty"`
	RiderFlowTime   *int64             `json:"rider_flow_time,omitempty" xml:"rider_flow_time,omitempty"`
	OrderId         *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RiderLng        *float64           `json:"rider_lng,omitempty" xml:"rider_lng,omitempty"`
	RiderName       *string            `json:"rider_name,omitempty" xml:"rider_name,omitempty"`
	ThreeSource     *int               `json:"three_source,omitempty" xml:"three_source,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ThreeDeliveryId *string            `json:"three_delivery_id,omitempty" xml:"three_delivery_id,omitempty"`
	Behavior        *int               `json:"behavior,omitempty" xml:"behavior,omitempty" require:"true"`
	RiderLat        *float64           `json:"rider_lat,omitempty" xml:"rider_lat,omitempty"`
}

func (s OrderSyncStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderSyncStatusRequest) GoString() string {
	return s.String()
}

func (s *OrderSyncStatusRequest) SetRiderPhone(v string) *OrderSyncStatusRequest {
	s.RiderPhone = &v
	return s
}

func (s *OrderSyncStatusRequest) SetAcceptTime(v int64) *OrderSyncStatusRequest {
	s.AcceptTime = &v
	return s
}

func (s *OrderSyncStatusRequest) SetRiderPhoneType(v int64) *OrderSyncStatusRequest {
	s.RiderPhoneType = &v
	return s
}

func (s *OrderSyncStatusRequest) SetRiderFlowTime(v int64) *OrderSyncStatusRequest {
	s.RiderFlowTime = &v
	return s
}

func (s *OrderSyncStatusRequest) SetOrderId(v string) *OrderSyncStatusRequest {
	s.OrderId = &v
	return s
}

func (s *OrderSyncStatusRequest) SetRiderLng(v float64) *OrderSyncStatusRequest {
	s.RiderLng = &v
	return s
}

func (s *OrderSyncStatusRequest) SetRiderName(v string) *OrderSyncStatusRequest {
	s.RiderName = &v
	return s
}

func (s *OrderSyncStatusRequest) SetThreeSource(v int) *OrderSyncStatusRequest {
	s.ThreeSource = &v
	return s
}

func (s *OrderSyncStatusRequest) SetHeader(v map[string]*string) *OrderSyncStatusRequest {
	s.Header = v
	return s
}

func (s *OrderSyncStatusRequest) SetAccessToken(v string) *OrderSyncStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *OrderSyncStatusRequest) SetThreeDeliveryId(v string) *OrderSyncStatusRequest {
	s.ThreeDeliveryId = &v
	return s
}

func (s *OrderSyncStatusRequest) SetBehavior(v int) *OrderSyncStatusRequest {
	s.Behavior = &v
	return s
}

func (s *OrderSyncStatusRequest) SetRiderLat(v float64) *OrderSyncStatusRequest {
	s.RiderLat = &v
	return s
}

type OrderSyncStatusResponse struct {
	Extra *OrderSyncStatusResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *OrderSyncStatusResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrderSyncStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderSyncStatusResponse) GoString() string {
	return s.String()
}

func (s *OrderSyncStatusResponse) SetExtra(v *OrderSyncStatusResponseExtra) *OrderSyncStatusResponse {
	s.Extra = v
	return s
}

func (s *OrderSyncStatusResponse) SetData(v *OrderSyncStatusResponseData) *OrderSyncStatusResponse {
	s.Data = v
	return s
}

type OrderSyncStatusResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s OrderSyncStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrderSyncStatusResponseData) GoString() string {
	return s.String()
}

func (s *OrderSyncStatusResponseData) SetGwErrorCode(v int32) *OrderSyncStatusResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *OrderSyncStatusResponseData) SetGwDescription(v string) *OrderSyncStatusResponseData {
	s.GwDescription = &v
	return s
}

type OrderSyncStatusResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s OrderSyncStatusResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s OrderSyncStatusResponseExtra) GoString() string {
	return s.String()
}

func (s *OrderSyncStatusResponseExtra) SetSubErrorCode(v int32) *OrderSyncStatusResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *OrderSyncStatusResponseExtra) SetDescription(v string) *OrderSyncStatusResponseExtra {
	s.Description = &v
	return s
}

func (s *OrderSyncStatusResponseExtra) SetErrorCode(v int32) *OrderSyncStatusResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *OrderSyncStatusResponseExtra) SetLogid(v string) *OrderSyncStatusResponseExtra {
	s.Logid = &v
	return s
}

func (s *OrderSyncStatusResponseExtra) SetNow(v int64) *OrderSyncStatusResponseExtra {
	s.Now = &v
	return s
}

func (s *OrderSyncStatusResponseExtra) SetSubDescription(v string) *OrderSyncStatusResponseExtra {
	s.SubDescription = &v
	return s
}

type OrientedPlanTalentDetailRequest struct {
	PlanId       *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	DouyinIdList []*string          `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" require:"true" type:"Repeated"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s OrientedPlanTalentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OrientedPlanTalentDetailRequest) GoString() string {
	return s.String()
}

func (s *OrientedPlanTalentDetailRequest) SetPlanId(v int64) *OrientedPlanTalentDetailRequest {
	s.PlanId = &v
	return s
}

func (s *OrientedPlanTalentDetailRequest) SetDouyinIdList(v []*string) *OrientedPlanTalentDetailRequest {
	s.DouyinIdList = v
	return s
}

func (s *OrientedPlanTalentDetailRequest) SetHeader(v map[string]*string) *OrientedPlanTalentDetailRequest {
	s.Header = v
	return s
}

func (s *OrientedPlanTalentDetailRequest) SetAccessToken(v string) *OrientedPlanTalentDetailRequest {
	s.AccessToken = &v
	return s
}

type OrientedPlanTalentDetailResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *OrientedPlanTalentDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s OrientedPlanTalentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s OrientedPlanTalentDetailResponse) GoString() string {
	return s.String()
}

func (s *OrientedPlanTalentDetailResponse) SetErrMsg(v string) *OrientedPlanTalentDetailResponse {
	s.ErrMsg = &v
	return s
}

func (s *OrientedPlanTalentDetailResponse) SetErrNo(v int32) *OrientedPlanTalentDetailResponse {
	s.ErrNo = &v
	return s
}

func (s *OrientedPlanTalentDetailResponse) SetLogId(v string) *OrientedPlanTalentDetailResponse {
	s.LogId = &v
	return s
}

func (s *OrientedPlanTalentDetailResponse) SetData(v *OrientedPlanTalentDetailResponseData) *OrientedPlanTalentDetailResponse {
	s.Data = v
	return s
}

type OrientedPlanTalentDetailResponseData struct {
	Data map[string]*OrientedPlanTalentDetailResponseDataDataValue `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Date *string                                                   `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s OrientedPlanTalentDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s OrientedPlanTalentDetailResponseData) GoString() string {
	return s.String()
}

func (s *OrientedPlanTalentDetailResponseData) SetData(v map[string]*OrientedPlanTalentDetailResponseDataDataValue) *OrientedPlanTalentDetailResponseData {
	s.Data = v
	return s
}

func (s *OrientedPlanTalentDetailResponseData) SetDate(v string) *OrientedPlanTalentDetailResponseData {
	s.Date = &v
	return s
}

type OrientedPlanTalentDetailResponseDataDataValue struct {
	MediaSellInfo    []*OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem `json:"media_sell_info,omitempty" xml:"media_sell_info,omitempty" require:"true" type:"Repeated"`
	PlayCnt          *int64                                                            `json:"play_cnt,omitempty" xml:"play_cnt,omitempty" require:"true"`
	TalentCommission *int64                                                            `json:"talent_commission,omitempty" xml:"talent_commission,omitempty" require:"true"`
	UsedGmv          *int64                                                            `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	Gmv              *int64                                                            `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
}

func (s OrientedPlanTalentDetailResponseDataDataValue) String() string {
	return tea.Prettify(s)
}

func (s OrientedPlanTalentDetailResponseDataDataValue) GoString() string {
	return s.String()
}

func (s *OrientedPlanTalentDetailResponseDataDataValue) SetMediaSellInfo(v []*OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) *OrientedPlanTalentDetailResponseDataDataValue {
	s.MediaSellInfo = v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValue) SetPlayCnt(v int64) *OrientedPlanTalentDetailResponseDataDataValue {
	s.PlayCnt = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValue) SetTalentCommission(v int64) *OrientedPlanTalentDetailResponseDataDataValue {
	s.TalentCommission = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValue) SetUsedGmv(v int64) *OrientedPlanTalentDetailResponseDataDataValue {
	s.UsedGmv = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValue) SetGmv(v int64) *OrientedPlanTalentDetailResponseDataDataValue {
	s.Gmv = &v
	return s
}

type OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem struct {
	PlayCnt          *int64  `json:"play_cnt,omitempty" xml:"play_cnt,omitempty" require:"true"`
	TalentCommission *int64  `json:"talent_commission,omitempty" xml:"talent_commission,omitempty" require:"true"`
	UsedGmv          *int64  `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	ContentId        *string `json:"content_id,omitempty" xml:"content_id,omitempty" require:"true"`
	ContentOpenId    *string `json:"content_open_id,omitempty" xml:"content_open_id,omitempty" require:"true"`
	ContentType      *int32  `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	Gmv              *int64  `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
}

func (s OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) String() string {
	return tea.Prettify(s)
}

func (s OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) GoString() string {
	return s.String()
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetPlayCnt(v int64) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.PlayCnt = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetTalentCommission(v int64) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.TalentCommission = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetUsedGmv(v int64) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.UsedGmv = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetContentId(v string) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.ContentId = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetContentOpenId(v string) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.ContentOpenId = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetContentType(v int32) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.ContentType = &v
	return s
}

func (s *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem) SetGmv(v int64) *OrientedPlanTalentDetailResponseDataDataValueMediaSellInfoItem {
	s.Gmv = &v
	return s
}

type PassportOpenGetAccessCodeRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PassportOpenGetAccessCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetAccessCodeRequest) GoString() string {
	return s.String()
}

func (s *PassportOpenGetAccessCodeRequest) SetHeader(v map[string]*string) *PassportOpenGetAccessCodeRequest {
	s.Header = v
	return s
}

func (s *PassportOpenGetAccessCodeRequest) SetAccessToken(v string) *PassportOpenGetAccessCodeRequest {
	s.AccessToken = &v
	return s
}

type PassportOpenGetAccessCodeResponse struct {
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Data    *PassportOpenGetAccessCodeResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PassportOpenGetAccessCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetAccessCodeResponse) GoString() string {
	return s.String()
}

func (s *PassportOpenGetAccessCodeResponse) SetMessage(v string) *PassportOpenGetAccessCodeResponse {
	s.Message = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponse) SetData(v *PassportOpenGetAccessCodeResponseData) *PassportOpenGetAccessCodeResponse {
	s.Data = v
	return s
}

type PassportOpenGetAccessCodeResponseData struct {
	AccessCode  *string `json:"access_code,omitempty" xml:"access_code,omitempty"`
	Captcha     *string `json:"captcha,omitempty" xml:"captcha,omitempty"`
	DescUrl     *string `json:"desc_url,omitempty" xml:"desc_url,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
	LogId       *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
}

func (s PassportOpenGetAccessCodeResponseData) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetAccessCodeResponseData) GoString() string {
	return s.String()
}

func (s *PassportOpenGetAccessCodeResponseData) SetAccessCode(v string) *PassportOpenGetAccessCodeResponseData {
	s.AccessCode = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponseData) SetCaptcha(v string) *PassportOpenGetAccessCodeResponseData {
	s.Captcha = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponseData) SetDescUrl(v string) *PassportOpenGetAccessCodeResponseData {
	s.DescUrl = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponseData) SetDescription(v string) *PassportOpenGetAccessCodeResponseData {
	s.Description = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponseData) SetErrorCode(v int64) *PassportOpenGetAccessCodeResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PassportOpenGetAccessCodeResponseData) SetLogId(v string) *PassportOpenGetAccessCodeResponseData {
	s.LogId = &v
	return s
}

type PassportOpenGetClientCodeRequest struct {
	ClientSecret *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	ClientKey    *string            `json:"client_key,omitempty" xml:"client_key,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PassportOpenGetClientCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetClientCodeRequest) GoString() string {
	return s.String()
}

func (s *PassportOpenGetClientCodeRequest) SetClientSecret(v string) *PassportOpenGetClientCodeRequest {
	s.ClientSecret = &v
	return s
}

func (s *PassportOpenGetClientCodeRequest) SetClientKey(v string) *PassportOpenGetClientCodeRequest {
	s.ClientKey = &v
	return s
}

func (s *PassportOpenGetClientCodeRequest) SetHeader(v map[string]*string) *PassportOpenGetClientCodeRequest {
	s.Header = v
	return s
}

type PassportOpenGetClientCodeResponse struct {
	Data    *PassportOpenGetClientCodeResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
}

func (s PassportOpenGetClientCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetClientCodeResponse) GoString() string {
	return s.String()
}

func (s *PassportOpenGetClientCodeResponse) SetData(v *PassportOpenGetClientCodeResponseData) *PassportOpenGetClientCodeResponse {
	s.Data = v
	return s
}

func (s *PassportOpenGetClientCodeResponse) SetMessage(v string) *PassportOpenGetClientCodeResponse {
	s.Message = &v
	return s
}

type PassportOpenGetClientCodeResponseData struct {
	DescUrl     *string `json:"desc_url,omitempty" xml:"desc_url,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
	LogId       *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ClientCode  *string `json:"client_code,omitempty" xml:"client_code,omitempty"`
	Captcha     *string `json:"captcha,omitempty" xml:"captcha,omitempty"`
}

func (s PassportOpenGetClientCodeResponseData) String() string {
	return tea.Prettify(s)
}

func (s PassportOpenGetClientCodeResponseData) GoString() string {
	return s.String()
}

func (s *PassportOpenGetClientCodeResponseData) SetDescUrl(v string) *PassportOpenGetClientCodeResponseData {
	s.DescUrl = &v
	return s
}

func (s *PassportOpenGetClientCodeResponseData) SetDescription(v string) *PassportOpenGetClientCodeResponseData {
	s.Description = &v
	return s
}

func (s *PassportOpenGetClientCodeResponseData) SetErrorCode(v int64) *PassportOpenGetClientCodeResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PassportOpenGetClientCodeResponseData) SetLogId(v string) *PassportOpenGetClientCodeResponseData {
	s.LogId = &v
	return s
}

func (s *PassportOpenGetClientCodeResponseData) SetClientCode(v string) *PassportOpenGetClientCodeResponseData {
	s.ClientCode = &v
	return s
}

func (s *PassportOpenGetClientCodeResponseData) SetCaptcha(v string) *PassportOpenGetClientCodeResponseData {
	s.Captcha = &v
	return s
}

type PhysicalRoomOperateRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OpType      *int               `json:"op_type,omitempty" xml:"op_type,omitempty" require:"true"`
	RoomIds     []*string          `json:"room_ids,omitempty" xml:"room_ids,omitempty" require:"true" type:"Repeated"`
}

func (s PhysicalRoomOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomOperateRequest) GoString() string {
	return s.String()
}

func (s *PhysicalRoomOperateRequest) SetHeader(v map[string]*string) *PhysicalRoomOperateRequest {
	s.Header = v
	return s
}

func (s *PhysicalRoomOperateRequest) SetAccessToken(v string) *PhysicalRoomOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *PhysicalRoomOperateRequest) SetAccountId(v string) *PhysicalRoomOperateRequest {
	s.AccountId = &v
	return s
}

func (s *PhysicalRoomOperateRequest) SetOpType(v int) *PhysicalRoomOperateRequest {
	s.OpType = &v
	return s
}

func (s *PhysicalRoomOperateRequest) SetRoomIds(v []*string) *PhysicalRoomOperateRequest {
	s.RoomIds = v
	return s
}

type PhysicalRoomOperateResponse struct {
	Extra *PhysicalRoomOperateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PhysicalRoomOperateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PhysicalRoomOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomOperateResponse) GoString() string {
	return s.String()
}

func (s *PhysicalRoomOperateResponse) SetExtra(v *PhysicalRoomOperateResponseExtra) *PhysicalRoomOperateResponse {
	s.Extra = v
	return s
}

func (s *PhysicalRoomOperateResponse) SetData(v *PhysicalRoomOperateResponseData) *PhysicalRoomOperateResponse {
	s.Data = v
	return s
}

type PhysicalRoomOperateResponseData struct {
	GwDescription *string                                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	RoomList      []*PhysicalRoomOperateResponseDataRoomListItem `json:"room_list,omitempty" xml:"room_list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PhysicalRoomOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomOperateResponseData) GoString() string {
	return s.String()
}

func (s *PhysicalRoomOperateResponseData) SetGwDescription(v string) *PhysicalRoomOperateResponseData {
	s.GwDescription = &v
	return s
}

func (s *PhysicalRoomOperateResponseData) SetRoomList(v []*PhysicalRoomOperateResponseDataRoomListItem) *PhysicalRoomOperateResponseData {
	s.RoomList = v
	return s
}

func (s *PhysicalRoomOperateResponseData) SetGwErrorCode(v int32) *PhysicalRoomOperateResponseData {
	s.GwErrorCode = &v
	return s
}

type PhysicalRoomOperateResponseDataRoomListItem struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	RoomId  *string `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	Code    *int32  `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s PhysicalRoomOperateResponseDataRoomListItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomOperateResponseDataRoomListItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomOperateResponseDataRoomListItem) SetMessage(v string) *PhysicalRoomOperateResponseDataRoomListItem {
	s.Message = &v
	return s
}

func (s *PhysicalRoomOperateResponseDataRoomListItem) SetRoomId(v string) *PhysicalRoomOperateResponseDataRoomListItem {
	s.RoomId = &v
	return s
}

func (s *PhysicalRoomOperateResponseDataRoomListItem) SetCode(v int32) *PhysicalRoomOperateResponseDataRoomListItem {
	s.Code = &v
	return s
}

type PhysicalRoomOperateResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s PhysicalRoomOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *PhysicalRoomOperateResponseExtra) SetSubDescription(v string) *PhysicalRoomOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PhysicalRoomOperateResponseExtra) SetSubErrorCode(v int32) *PhysicalRoomOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PhysicalRoomOperateResponseExtra) SetDescription(v string) *PhysicalRoomOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *PhysicalRoomOperateResponseExtra) SetErrorCode(v int32) *PhysicalRoomOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PhysicalRoomOperateResponseExtra) SetLogid(v string) *PhysicalRoomOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *PhysicalRoomOperateResponseExtra) SetNow(v int64) *PhysicalRoomOperateResponseExtra {
	s.Now = &v
	return s
}

type PhysicalRoomQueryRequest struct {
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId    *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	NeedRatePlan *bool              `json:"need_rate_plan,omitempty" xml:"need_rate_plan,omitempty"`
	RoomIds      []*string          `json:"room_ids,omitempty" xml:"room_ids,omitempty" type:"Repeated"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PhysicalRoomQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryRequest) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryRequest) SetAccessToken(v string) *PhysicalRoomQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *PhysicalRoomQueryRequest) SetAccountId(v string) *PhysicalRoomQueryRequest {
	s.AccountId = &v
	return s
}

func (s *PhysicalRoomQueryRequest) SetNeedRatePlan(v bool) *PhysicalRoomQueryRequest {
	s.NeedRatePlan = &v
	return s
}

func (s *PhysicalRoomQueryRequest) SetRoomIds(v []*string) *PhysicalRoomQueryRequest {
	s.RoomIds = v
	return s
}

func (s *PhysicalRoomQueryRequest) SetHeader(v map[string]*string) *PhysicalRoomQueryRequest {
	s.Header = v
	return s
}

type PhysicalRoomQueryResponse struct {
	Extra *PhysicalRoomQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PhysicalRoomQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PhysicalRoomQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryResponse) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryResponse) SetExtra(v *PhysicalRoomQueryResponseExtra) *PhysicalRoomQueryResponse {
	s.Extra = v
	return s
}

func (s *PhysicalRoomQueryResponse) SetData(v *PhysicalRoomQueryResponseData) *PhysicalRoomQueryResponse {
	s.Data = v
	return s
}

type PhysicalRoomQueryResponseData struct {
	GwDescription *string                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	RoomList      []*PhysicalRoomQueryResponseDataRoomListItem `json:"room_list,omitempty" xml:"room_list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PhysicalRoomQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryResponseData) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryResponseData) SetGwDescription(v string) *PhysicalRoomQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *PhysicalRoomQueryResponseData) SetRoomList(v []*PhysicalRoomQueryResponseDataRoomListItem) *PhysicalRoomQueryResponseData {
	s.RoomList = v
	return s
}

func (s *PhysicalRoomQueryResponseData) SetGwErrorCode(v int32) *PhysicalRoomQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type PhysicalRoomQueryResponseDataRoomListItem struct {
	Status       *int                                                         `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AuditMessage *string                                                      `json:"audit_message,omitempty" xml:"audit_message,omitempty"`
	CnName       *string                                                      `json:"cn_name,omitempty" xml:"cn_name,omitempty" require:"true"`
	RatePlanList []*PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem `json:"rate_plan_list,omitempty" xml:"rate_plan_list,omitempty" type:"Repeated"`
	RoomId       *string                                                      `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
}

func (s PhysicalRoomQueryResponseDataRoomListItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryResponseDataRoomListItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryResponseDataRoomListItem) SetStatus(v int) *PhysicalRoomQueryResponseDataRoomListItem {
	s.Status = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItem) SetAuditMessage(v string) *PhysicalRoomQueryResponseDataRoomListItem {
	s.AuditMessage = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItem) SetCnName(v string) *PhysicalRoomQueryResponseDataRoomListItem {
	s.CnName = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItem) SetRatePlanList(v []*PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) *PhysicalRoomQueryResponseDataRoomListItem {
	s.RatePlanList = v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItem) SetRoomId(v string) *PhysicalRoomQueryResponseDataRoomListItem {
	s.RoomId = &v
	return s
}

type PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem struct {
	Status       *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanName *string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty" require:"true"`
	RatePlanType *int    `json:"rate_plan_type,omitempty" xml:"rate_plan_type,omitempty" require:"true"`
}

func (s PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) SetStatus(v int) *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem {
	s.Status = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) SetRatePlanId(v string) *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem {
	s.RatePlanId = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) SetRatePlanName(v string) *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem {
	s.RatePlanName = &v
	return s
}

func (s *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem) SetRatePlanType(v int) *PhysicalRoomQueryResponseDataRoomListItemRatePlanListItem {
	s.RatePlanType = &v
	return s
}

type PhysicalRoomQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PhysicalRoomQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *PhysicalRoomQueryResponseExtra) SetLogid(v string) *PhysicalRoomQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *PhysicalRoomQueryResponseExtra) SetNow(v int64) *PhysicalRoomQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *PhysicalRoomQueryResponseExtra) SetSubDescription(v string) *PhysicalRoomQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PhysicalRoomQueryResponseExtra) SetSubErrorCode(v int32) *PhysicalRoomQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PhysicalRoomQueryResponseExtra) SetDescription(v string) *PhysicalRoomQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *PhysicalRoomQueryResponseExtra) SetErrorCode(v int32) *PhysicalRoomQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type PhysicalRoomSaveRequest struct {
	RoomInfo    *PhysicalRoomSaveRequestRoomInfo `json:"room_info,omitempty" xml:"room_info,omitempty" require:"true"`
	AccountId   *string                          `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ByFileds    []*string                        `json:"by_fileds,omitempty" xml:"by_fileds,omitempty" type:"Repeated"`
	PoiId       *string                          `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	Header      map[string]*string               `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                          `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PhysicalRoomSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequest) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequest) SetRoomInfo(v *PhysicalRoomSaveRequestRoomInfo) *PhysicalRoomSaveRequest {
	s.RoomInfo = v
	return s
}

func (s *PhysicalRoomSaveRequest) SetAccountId(v string) *PhysicalRoomSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PhysicalRoomSaveRequest) SetByFileds(v []*string) *PhysicalRoomSaveRequest {
	s.ByFileds = v
	return s
}

func (s *PhysicalRoomSaveRequest) SetPoiId(v string) *PhysicalRoomSaveRequest {
	s.PoiId = &v
	return s
}

func (s *PhysicalRoomSaveRequest) SetHeader(v map[string]*string) *PhysicalRoomSaveRequest {
	s.Header = v
	return s
}

func (s *PhysicalRoomSaveRequest) SetAccessToken(v string) *PhysicalRoomSaveRequest {
	s.AccessToken = &v
	return s
}

type PhysicalRoomSaveRequestRoomInfo struct {
	FacilityMap  map[string]*PhysicalRoomSaveRequestRoomInfoFacilityMapValue `json:"facility_map,omitempty" xml:"facility_map,omitempty"`
	EnName       *string                                                     `json:"en_name,omitempty" xml:"en_name,omitempty"`
	RoomNum      *int32                                                      `json:"room_num,omitempty" xml:"room_num,omitempty"`
	CnName       *string                                                     `json:"cn_name,omitempty" xml:"cn_name,omitempty" require:"true"`
	OutRoomId    *string                                                     `json:"out_room_id,omitempty" xml:"out_room_id,omitempty" require:"true"`
	Window       *int                                                        `json:"window,omitempty" xml:"window,omitempty"`
	BedGroups    []*PhysicalRoomSaveRequestRoomInfoBedGroupsItem             `json:"bed_groups,omitempty" xml:"bed_groups,omitempty" type:"Repeated"`
	Active       *bool                                                       `json:"active,omitempty" xml:"active,omitempty"`
	Images       []*PhysicalRoomSaveRequestRoomInfoImagesItem                `json:"images,omitempty" xml:"images,omitempty" require:"true" type:"Repeated"`
	RoomId       *string                                                     `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Descriptions []*string                                                   `json:"descriptions,omitempty" xml:"descriptions,omitempty" type:"Repeated"`
	Floor        []*PhysicalRoomSaveRequestRoomInfoFloorItem                 `json:"floor,omitempty" xml:"floor,omitempty" type:"Repeated"`
	Smoking      *int                                                        `json:"smoking,omitempty" xml:"smoking,omitempty"`
	Internets    []*int                                                      `json:"internets,omitempty" xml:"internets,omitempty" type:"Repeated"`
	MaxOccupancy *int32                                                      `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty" require:"true"`
	CategoryId   *int64                                                      `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	IsTestData   *bool                                                       `json:"is_test_data,omitempty" xml:"is_test_data,omitempty"`
	Area         *PhysicalRoomSaveRequestRoomInfoArea                        `json:"area,omitempty" xml:"area,omitempty"`
}

func (s PhysicalRoomSaveRequestRoomInfo) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfo) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetFacilityMap(v map[string]*PhysicalRoomSaveRequestRoomInfoFacilityMapValue) *PhysicalRoomSaveRequestRoomInfo {
	s.FacilityMap = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetEnName(v string) *PhysicalRoomSaveRequestRoomInfo {
	s.EnName = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetRoomNum(v int32) *PhysicalRoomSaveRequestRoomInfo {
	s.RoomNum = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetCnName(v string) *PhysicalRoomSaveRequestRoomInfo {
	s.CnName = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetOutRoomId(v string) *PhysicalRoomSaveRequestRoomInfo {
	s.OutRoomId = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetWindow(v int) *PhysicalRoomSaveRequestRoomInfo {
	s.Window = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetBedGroups(v []*PhysicalRoomSaveRequestRoomInfoBedGroupsItem) *PhysicalRoomSaveRequestRoomInfo {
	s.BedGroups = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetActive(v bool) *PhysicalRoomSaveRequestRoomInfo {
	s.Active = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetImages(v []*PhysicalRoomSaveRequestRoomInfoImagesItem) *PhysicalRoomSaveRequestRoomInfo {
	s.Images = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetRoomId(v string) *PhysicalRoomSaveRequestRoomInfo {
	s.RoomId = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetDescriptions(v []*string) *PhysicalRoomSaveRequestRoomInfo {
	s.Descriptions = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetFloor(v []*PhysicalRoomSaveRequestRoomInfoFloorItem) *PhysicalRoomSaveRequestRoomInfo {
	s.Floor = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetSmoking(v int) *PhysicalRoomSaveRequestRoomInfo {
	s.Smoking = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetInternets(v []*int) *PhysicalRoomSaveRequestRoomInfo {
	s.Internets = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetMaxOccupancy(v int32) *PhysicalRoomSaveRequestRoomInfo {
	s.MaxOccupancy = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetCategoryId(v int64) *PhysicalRoomSaveRequestRoomInfo {
	s.CategoryId = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetIsTestData(v bool) *PhysicalRoomSaveRequestRoomInfo {
	s.IsTestData = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfo) SetArea(v *PhysicalRoomSaveRequestRoomInfoArea) *PhysicalRoomSaveRequestRoomInfo {
	s.Area = v
	return s
}

type PhysicalRoomSaveRequestRoomInfoArea struct {
	ValueList []*int32 `json:"value_list,omitempty" xml:"value_list,omitempty" type:"Repeated"`
	ValueType *int     `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s PhysicalRoomSaveRequestRoomInfoArea) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoArea) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoArea) SetValueList(v []*int32) *PhysicalRoomSaveRequestRoomInfoArea {
	s.ValueList = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfoArea) SetValueType(v int) *PhysicalRoomSaveRequestRoomInfoArea {
	s.ValueType = &v
	return s
}

type PhysicalRoomSaveRequestRoomInfoBedGroupsItem struct {
	BedItems []*PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem `json:"bed_items,omitempty" xml:"bed_items,omitempty" require:"true" type:"Repeated"`
}

func (s PhysicalRoomSaveRequestRoomInfoBedGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoBedGroupsItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoBedGroupsItem) SetBedItems(v []*PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) *PhysicalRoomSaveRequestRoomInfoBedGroupsItem {
	s.BedItems = v
	return s
}

type PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem struct {
	BedSize *int32 `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
	BedType *int   `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	BedNum  *int32 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
}

func (s PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) SetBedSize(v int32) *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem {
	s.BedSize = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) SetBedType(v int) *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem {
	s.BedType = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem) SetBedNum(v int32) *PhysicalRoomSaveRequestRoomInfoBedGroupsItemBedItemsItem {
	s.BedNum = &v
	return s
}

type PhysicalRoomSaveRequestRoomInfoFacilityMapValue struct {
	FacilityIds []*int64 `json:"facility_ids,omitempty" xml:"facility_ids,omitempty" require:"true" type:"Repeated"`
}

func (s PhysicalRoomSaveRequestRoomInfoFacilityMapValue) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoFacilityMapValue) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoFacilityMapValue) SetFacilityIds(v []*int64) *PhysicalRoomSaveRequestRoomInfoFacilityMapValue {
	s.FacilityIds = v
	return s
}

type PhysicalRoomSaveRequestRoomInfoFloorItem struct {
	ValueList []*int32 `json:"value_list,omitempty" xml:"value_list,omitempty" type:"Repeated"`
	ValueType *int     `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s PhysicalRoomSaveRequestRoomInfoFloorItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoFloorItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoFloorItem) SetValueList(v []*int32) *PhysicalRoomSaveRequestRoomInfoFloorItem {
	s.ValueList = v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfoFloorItem) SetValueType(v int) *PhysicalRoomSaveRequestRoomInfoFloorItem {
	s.ValueType = &v
	return s
}

type PhysicalRoomSaveRequestRoomInfoImagesItem struct {
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty" require:"true"`
	ImageUri *string `json:"image_uri,omitempty" xml:"image_uri,omitempty"`
}

func (s PhysicalRoomSaveRequestRoomInfoImagesItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveRequestRoomInfoImagesItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveRequestRoomInfoImagesItem) SetImageUrl(v string) *PhysicalRoomSaveRequestRoomInfoImagesItem {
	s.ImageUrl = &v
	return s
}

func (s *PhysicalRoomSaveRequestRoomInfoImagesItem) SetImageUri(v string) *PhysicalRoomSaveRequestRoomInfoImagesItem {
	s.ImageUri = &v
	return s
}

type PhysicalRoomSaveResponse struct {
	Data  *PhysicalRoomSaveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PhysicalRoomSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PhysicalRoomSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveResponse) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveResponse) SetData(v *PhysicalRoomSaveResponseData) *PhysicalRoomSaveResponse {
	s.Data = v
	return s
}

func (s *PhysicalRoomSaveResponse) SetExtra(v *PhysicalRoomSaveResponseExtra) *PhysicalRoomSaveResponse {
	s.Extra = v
	return s
}

type PhysicalRoomSaveResponseData struct {
	Ok              *bool   `json:"ok,omitempty" xml:"ok,omitempty" require:"true"`
	RoomId          *string `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	DuplicateRoomId *string `json:"duplicate_room_id,omitempty" xml:"duplicate_room_id,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	GwErrorCode     *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription   *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PhysicalRoomSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveResponseData) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveResponseData) SetOk(v bool) *PhysicalRoomSaveResponseData {
	s.Ok = &v
	return s
}

func (s *PhysicalRoomSaveResponseData) SetRoomId(v string) *PhysicalRoomSaveResponseData {
	s.RoomId = &v
	return s
}

func (s *PhysicalRoomSaveResponseData) SetDuplicateRoomId(v string) *PhysicalRoomSaveResponseData {
	s.DuplicateRoomId = &v
	return s
}

func (s *PhysicalRoomSaveResponseData) SetMessage(v string) *PhysicalRoomSaveResponseData {
	s.Message = &v
	return s
}

func (s *PhysicalRoomSaveResponseData) SetGwErrorCode(v int32) *PhysicalRoomSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PhysicalRoomSaveResponseData) SetGwDescription(v string) *PhysicalRoomSaveResponseData {
	s.GwDescription = &v
	return s
}
