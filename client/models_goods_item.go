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
	util "github.com/bytedance/douyin-openapi-util-go/client"
)

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo struct {
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) SetBrandActivityId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) SetPlanId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) SetPromoId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem struct {
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetConstantVal(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetPriceRel(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetSharedQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetStockRel(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetBizId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) SetCoefficient(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetSoldCount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetSoldQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetStockQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetAvailQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetFrozenQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) SetLimitType(v int) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock {
	s.LimitType = &v
	return s
}

type GoodsProductDraftGetResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseExtra) SetLogid(v string) *GoodsProductDraftGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *GoodsProductDraftGetResponseExtra) SetNow(v int64) *GoodsProductDraftGetResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsProductDraftGetResponseExtra) SetSubDescription(v string) *GoodsProductDraftGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsProductDraftGetResponseExtra) SetSubErrorCode(v int32) *GoodsProductDraftGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseExtra) SetDescription(v string) *GoodsProductDraftGetResponseExtra {
	s.Description = &v
	return s
}

func (s *GoodsProductDraftGetResponseExtra) SetErrorCode(v int32) *GoodsProductDraftGetResponseExtra {
	s.ErrorCode = &v
	return s
}

type GoodsProductFreeAuditRequest struct {
	ProductId    *string                           `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Base         *GoodsProductFreeAuditRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId    *string                           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutId        *string                           `json:"out_id,omitempty" xml:"out_id,omitempty"`
	SoldEndTime  *int64                            `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	StockQty     *int64                            `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	ActualAmount *int64                            `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	Header       map[string]*string                `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                           `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GoodsProductFreeAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditRequest) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditRequest) SetProductId(v string) *GoodsProductFreeAuditRequest {
	s.ProductId = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetBase(v *GoodsProductFreeAuditRequestBase) *GoodsProductFreeAuditRequest {
	s.Base = v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetAccountId(v string) *GoodsProductFreeAuditRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetOutId(v string) *GoodsProductFreeAuditRequest {
	s.OutId = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetSoldEndTime(v int64) *GoodsProductFreeAuditRequest {
	s.SoldEndTime = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetStockQty(v int64) *GoodsProductFreeAuditRequest {
	s.StockQty = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetActualAmount(v int64) *GoodsProductFreeAuditRequest {
	s.ActualAmount = &v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetHeader(v map[string]*string) *GoodsProductFreeAuditRequest {
	s.Header = v
	return s
}

func (s *GoodsProductFreeAuditRequest) SetAccessToken(v string) *GoodsProductFreeAuditRequest {
	s.AccessToken = &v
	return s
}

type GoodsProductFreeAuditRequestBase struct {
	TrafficEnv *GoodsProductFreeAuditRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                     `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                     `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                     `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                     `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s GoodsProductFreeAuditRequestBase) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditRequestBase) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditRequestBase) SetTrafficEnv(v *GoodsProductFreeAuditRequestBaseTrafficEnv) *GoodsProductFreeAuditRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *GoodsProductFreeAuditRequestBase) SetAddr(v string) *GoodsProductFreeAuditRequestBase {
	s.Addr = &v
	return s
}

func (s *GoodsProductFreeAuditRequestBase) SetCaller(v string) *GoodsProductFreeAuditRequestBase {
	s.Caller = &v
	return s
}

func (s *GoodsProductFreeAuditRequestBase) SetClient(v string) *GoodsProductFreeAuditRequestBase {
	s.Client = &v
	return s
}

func (s *GoodsProductFreeAuditRequestBase) SetExtra(v map[string]*string) *GoodsProductFreeAuditRequestBase {
	s.Extra = v
	return s
}

func (s *GoodsProductFreeAuditRequestBase) SetLogID(v string) *GoodsProductFreeAuditRequestBase {
	s.LogID = &v
	return s
}

type GoodsProductFreeAuditRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s GoodsProductFreeAuditRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditRequestBaseTrafficEnv) SetEnv(v string) *GoodsProductFreeAuditRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *GoodsProductFreeAuditRequestBaseTrafficEnv) SetOpen(v bool) *GoodsProductFreeAuditRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type GoodsProductFreeAuditResponse struct {
	Extra    *GoodsProductFreeAuditResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *GoodsProductFreeAuditResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *GoodsProductFreeAuditResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GoodsProductFreeAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditResponse) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditResponse) SetExtra(v *GoodsProductFreeAuditResponseExtra) *GoodsProductFreeAuditResponse {
	s.Extra = v
	return s
}

func (s *GoodsProductFreeAuditResponse) SetBaseResp(v *GoodsProductFreeAuditResponseBaseResp) *GoodsProductFreeAuditResponse {
	s.BaseResp = v
	return s
}

func (s *GoodsProductFreeAuditResponse) SetData(v *GoodsProductFreeAuditResponseData) *GoodsProductFreeAuditResponse {
	s.Data = v
	return s
}

type GoodsProductFreeAuditResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GoodsProductFreeAuditResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditResponseBaseResp) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditResponseBaseResp) SetStatusCode(v int32) *GoodsProductFreeAuditResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *GoodsProductFreeAuditResponseBaseResp) SetStatusMessage(v string) *GoodsProductFreeAuditResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *GoodsProductFreeAuditResponseBaseResp) SetExtra(v map[string]*string) *GoodsProductFreeAuditResponseBaseResp {
	s.Extra = v
	return s
}

type GoodsProductFreeAuditResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsProductFreeAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditResponseData) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditResponseData) SetErrorCode(v int32) *GoodsProductFreeAuditResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsProductFreeAuditResponseData) SetDescription(v string) *GoodsProductFreeAuditResponseData {
	s.Description = &v
	return s
}

type GoodsProductFreeAuditResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s GoodsProductFreeAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductFreeAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsProductFreeAuditResponseExtra) SetNow(v int64) *GoodsProductFreeAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsProductFreeAuditResponseExtra) SetSubDescription(v string) *GoodsProductFreeAuditResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsProductFreeAuditResponseExtra) SetSubErrorCode(v int32) *GoodsProductFreeAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsProductFreeAuditResponseExtra) SetDescription(v string) *GoodsProductFreeAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *GoodsProductFreeAuditResponseExtra) SetErrorCode(v int32) *GoodsProductFreeAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsProductFreeAuditResponseExtra) SetLogid(v string) *GoodsProductFreeAuditResponseExtra {
	s.Logid = &v
	return s
}

type GoodsProductOperateRequest struct {
	OutId       *string                         `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId   *string                         `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Header      map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base        *GoodsProductOperateRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                         `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OpType      *int                            `json:"op_type,omitempty" xml:"op_type,omitempty" require:"true"`
}

func (s GoodsProductOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateRequest) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateRequest) SetOutId(v string) *GoodsProductOperateRequest {
	s.OutId = &v
	return s
}

func (s *GoodsProductOperateRequest) SetProductId(v string) *GoodsProductOperateRequest {
	s.ProductId = &v
	return s
}

func (s *GoodsProductOperateRequest) SetHeader(v map[string]*string) *GoodsProductOperateRequest {
	s.Header = v
	return s
}

func (s *GoodsProductOperateRequest) SetAccessToken(v string) *GoodsProductOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *GoodsProductOperateRequest) SetBase(v *GoodsProductOperateRequestBase) *GoodsProductOperateRequest {
	s.Base = v
	return s
}

func (s *GoodsProductOperateRequest) SetAccountId(v string) *GoodsProductOperateRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsProductOperateRequest) SetOpType(v int) *GoodsProductOperateRequest {
	s.OpType = &v
	return s
}

type GoodsProductOperateRequestBase struct {
	Client     *string                                   `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                        `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                   `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *GoodsProductOperateRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                   `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                   `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s GoodsProductOperateRequestBase) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateRequestBase) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateRequestBase) SetClient(v string) *GoodsProductOperateRequestBase {
	s.Client = &v
	return s
}

func (s *GoodsProductOperateRequestBase) SetExtra(v map[string]*string) *GoodsProductOperateRequestBase {
	s.Extra = v
	return s
}

func (s *GoodsProductOperateRequestBase) SetLogID(v string) *GoodsProductOperateRequestBase {
	s.LogID = &v
	return s
}

func (s *GoodsProductOperateRequestBase) SetTrafficEnv(v *GoodsProductOperateRequestBaseTrafficEnv) *GoodsProductOperateRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *GoodsProductOperateRequestBase) SetAddr(v string) *GoodsProductOperateRequestBase {
	s.Addr = &v
	return s
}

func (s *GoodsProductOperateRequestBase) SetCaller(v string) *GoodsProductOperateRequestBase {
	s.Caller = &v
	return s
}

type GoodsProductOperateRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s GoodsProductOperateRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateRequestBaseTrafficEnv) SetOpen(v bool) *GoodsProductOperateRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *GoodsProductOperateRequestBaseTrafficEnv) SetEnv(v string) *GoodsProductOperateRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type GoodsProductOperateResponse struct {
	BaseResp *GoodsProductOperateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *GoodsProductOperateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *GoodsProductOperateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s GoodsProductOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateResponse) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateResponse) SetBaseResp(v *GoodsProductOperateResponseBaseResp) *GoodsProductOperateResponse {
	s.BaseResp = v
	return s
}

func (s *GoodsProductOperateResponse) SetData(v *GoodsProductOperateResponseData) *GoodsProductOperateResponse {
	s.Data = v
	return s
}

func (s *GoodsProductOperateResponse) SetExtra(v *GoodsProductOperateResponseExtra) *GoodsProductOperateResponse {
	s.Extra = v
	return s
}

type GoodsProductOperateResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GoodsProductOperateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateResponseBaseResp) SetStatusMessage(v string) *GoodsProductOperateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *GoodsProductOperateResponseBaseResp) SetExtra(v map[string]*string) *GoodsProductOperateResponseBaseResp {
	s.Extra = v
	return s
}

func (s *GoodsProductOperateResponseBaseResp) SetStatusCode(v int32) *GoodsProductOperateResponseBaseResp {
	s.StatusCode = &v
	return s
}

type GoodsProductOperateResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s GoodsProductOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateResponseData) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateResponseData) SetDescription(v string) *GoodsProductOperateResponseData {
	s.Description = &v
	return s
}

func (s *GoodsProductOperateResponseData) SetErrorCode(v int32) *GoodsProductOperateResponseData {
	s.ErrorCode = &v
	return s
}

type GoodsProductOperateResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsProductOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsProductOperateResponseExtra) SetErrorCode(v int32) *GoodsProductOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsProductOperateResponseExtra) SetLogid(v string) *GoodsProductOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *GoodsProductOperateResponseExtra) SetNow(v int64) *GoodsProductOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsProductOperateResponseExtra) SetSubDescription(v string) *GoodsProductOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsProductOperateResponseExtra) SetSubErrorCode(v int32) *GoodsProductOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsProductOperateResponseExtra) SetDescription(v string) *GoodsProductOperateResponseExtra {
	s.Description = &v
	return s
}

type GoodsSkuBatchSaveRequest struct {
	AccessToken  *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Skus         []*GoodsSkuBatchSaveRequestSkusItem `json:"skus,omitempty" xml:"skus,omitempty" require:"true" type:"Repeated"`
	Base         *GoodsSkuBatchSaveRequestBase       `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId    *string                             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ProductId    *string                             `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductOutId *string                             `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	Header       map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
}

func (s GoodsSkuBatchSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequest) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequest) SetAccessToken(v string) *GoodsSkuBatchSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetSkus(v []*GoodsSkuBatchSaveRequestSkusItem) *GoodsSkuBatchSaveRequest {
	s.Skus = v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetBase(v *GoodsSkuBatchSaveRequestBase) *GoodsSkuBatchSaveRequest {
	s.Base = v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetAccountId(v string) *GoodsSkuBatchSaveRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetProductId(v string) *GoodsSkuBatchSaveRequest {
	s.ProductId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetProductOutId(v string) *GoodsSkuBatchSaveRequest {
	s.ProductOutId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequest) SetHeader(v map[string]*string) *GoodsSkuBatchSaveRequest {
	s.Header = v
	return s
}

type GoodsSkuBatchSaveRequestBase struct {
	Addr       *string                                 `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                 `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                 `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                 `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *GoodsSkuBatchSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s GoodsSkuBatchSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestBase) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestBase) SetAddr(v string) *GoodsSkuBatchSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestBase) SetCaller(v string) *GoodsSkuBatchSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestBase) SetClient(v string) *GoodsSkuBatchSaveRequestBase {
	s.Client = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestBase) SetExtra(v map[string]*string) *GoodsSkuBatchSaveRequestBase {
	s.Extra = v
	return s
}

func (s *GoodsSkuBatchSaveRequestBase) SetLogID(v string) *GoodsSkuBatchSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestBase) SetTrafficEnv(v *GoodsSkuBatchSaveRequestBaseTrafficEnv) *GoodsSkuBatchSaveRequestBase {
	s.TrafficEnv = v
	return s
}

type GoodsSkuBatchSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s GoodsSkuBatchSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestBaseTrafficEnv) SetEnv(v string) *GoodsSkuBatchSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestBaseTrafficEnv) SetOpen(v bool) *GoodsSkuBatchSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type GoodsSkuBatchSaveRequestSkusItem struct {
	Extra           *string                                 `json:"extra,omitempty" xml:"extra,omitempty"`
	BindSkus        []*string                               `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	ActualAmount    *int64                                  `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OutSkuId        *string                                 `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuName         *string                                 `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	Status          *int                                    `json:"status,omitempty" xml:"status,omitempty"`
	SkuId           *string                                 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuExt          *GoodsSkuBatchSaveRequestSkusItemSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	OriginAmount    *int64                                  `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	UpdateTime      *int64                                  `json:"update_time,omitempty" xml:"update_time,omitempty"`
	AttrKeyValueMap map[string]*string                      `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	CreateTime      *int64                                  `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Stock           *GoodsSkuBatchSaveRequestSkusItemStock  `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s GoodsSkuBatchSaveRequestSkusItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestSkusItem) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetExtra(v string) *GoodsSkuBatchSaveRequestSkusItem {
	s.Extra = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetBindSkus(v []*string) *GoodsSkuBatchSaveRequestSkusItem {
	s.BindSkus = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetActualAmount(v int64) *GoodsSkuBatchSaveRequestSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetOutSkuId(v string) *GoodsSkuBatchSaveRequestSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetSkuName(v string) *GoodsSkuBatchSaveRequestSkusItem {
	s.SkuName = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetStatus(v int) *GoodsSkuBatchSaveRequestSkusItem {
	s.Status = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetSkuId(v string) *GoodsSkuBatchSaveRequestSkusItem {
	s.SkuId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetSkuExt(v *GoodsSkuBatchSaveRequestSkusItemSkuExt) *GoodsSkuBatchSaveRequestSkusItem {
	s.SkuExt = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetOriginAmount(v int64) *GoodsSkuBatchSaveRequestSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetUpdateTime(v int64) *GoodsSkuBatchSaveRequestSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetAttrKeyValueMap(v map[string]*string) *GoodsSkuBatchSaveRequestSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetCreateTime(v int64) *GoodsSkuBatchSaveRequestSkusItem {
	s.CreateTime = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItem) SetStock(v *GoodsSkuBatchSaveRequestSkusItemStock) *GoodsSkuBatchSaveRequestSkusItem {
	s.Stock = v
	return s
}

type GoodsSkuBatchSaveRequestSkusItemSkuExt struct {
	OriginStockQty      *int64                                                     `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	RelRuleList         []*GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	BizId2cList         []*int64                                                   `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	BindSkuId           *int64                                                     `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	BindProductId       *int64                                                     `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	TakeawayPresaleInfo *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	BindSkus2c          map[int64][]*int64                                         `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	OriSkus             map[int64][]*int64                                         `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	LifeBizCode         *string                                                    `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExt) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExt) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetOriginStockQty(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetRelRuleList(v []*GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.RelRuleList = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetBizId2cList(v []*int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.BizId2cList = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetBindSkuId(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetBindProductId(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.BindProductId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetTakeawayPresaleInfo(v *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetBindSkus2c(v map[int64][]*int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetOriSkus(v map[int64][]*int64) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.OriSkus = v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExt) SetLifeBizCode(v string) *GoodsSkuBatchSaveRequestSkusItemSkuExt {
	s.LifeBizCode = &v
	return s
}

type GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem struct {
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetBizId(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetCoefficient(v string) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetConstantVal(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetPriceRel(v bool) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetSharedQty(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem) SetStockRel(v bool) *GoodsSkuBatchSaveRequestSkusItemSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

type GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *GoodsSkuBatchSaveRequestSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

type GoodsSkuBatchSaveRequestSkusItemStock struct {
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
}

func (s GoodsSkuBatchSaveRequestSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveRequestSkusItemStock) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetFrozenQty(v int64) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetLimitType(v int) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetSoldCount(v int64) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetSoldQty(v int64) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetStockQty(v int64) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *GoodsSkuBatchSaveRequestSkusItemStock) SetAvailQty(v int64) *GoodsSkuBatchSaveRequestSkusItemStock {
	s.AvailQty = &v
	return s
}

type GoodsSkuBatchSaveResponse struct {
	BaseResp *GoodsSkuBatchSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *GoodsSkuBatchSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *GoodsSkuBatchSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s GoodsSkuBatchSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveResponse) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveResponse) SetBaseResp(v *GoodsSkuBatchSaveResponseBaseResp) *GoodsSkuBatchSaveResponse {
	s.BaseResp = v
	return s
}

func (s *GoodsSkuBatchSaveResponse) SetData(v *GoodsSkuBatchSaveResponseData) *GoodsSkuBatchSaveResponse {
	s.Data = v
	return s
}

func (s *GoodsSkuBatchSaveResponse) SetExtra(v *GoodsSkuBatchSaveResponseExtra) *GoodsSkuBatchSaveResponse {
	s.Extra = v
	return s
}

type GoodsSkuBatchSaveResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GoodsSkuBatchSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveResponseBaseResp) SetStatusMessage(v string) *GoodsSkuBatchSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseBaseResp) SetExtra(v map[string]*string) *GoodsSkuBatchSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *GoodsSkuBatchSaveResponseBaseResp) SetStatusCode(v int32) *GoodsSkuBatchSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

type GoodsSkuBatchSaveResponseData struct {
	Description *string            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SkuIdList   []*int64           `json:"sku_id_list,omitempty" xml:"sku_id_list,omitempty" type:"Repeated"`
	SkuIds      map[string]*string `json:"sku_ids,omitempty" xml:"sku_ids,omitempty"`
}

func (s GoodsSkuBatchSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveResponseData) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveResponseData) SetDescription(v string) *GoodsSkuBatchSaveResponseData {
	s.Description = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseData) SetErrorCode(v int32) *GoodsSkuBatchSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseData) SetSkuIdList(v []*int64) *GoodsSkuBatchSaveResponseData {
	s.SkuIdList = v
	return s
}

func (s *GoodsSkuBatchSaveResponseData) SetSkuIds(v map[string]*string) *GoodsSkuBatchSaveResponseData {
	s.SkuIds = v
	return s
}

type GoodsSkuBatchSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s GoodsSkuBatchSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsSkuBatchSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsSkuBatchSaveResponseExtra) SetSubErrorCode(v int32) *GoodsSkuBatchSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseExtra) SetDescription(v string) *GoodsSkuBatchSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseExtra) SetErrorCode(v int32) *GoodsSkuBatchSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseExtra) SetLogid(v string) *GoodsSkuBatchSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseExtra) SetNow(v int64) *GoodsSkuBatchSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsSkuBatchSaveResponseExtra) SetSubDescription(v string) *GoodsSkuBatchSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type GoodsSpuOperateRequest struct {
	OutSpuId    *string            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	SpuId       *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	AccountId   *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Op          *int               `json:"op,omitempty" xml:"op,omitempty" require:"true"`
}

func (s GoodsSpuOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuOperateRequest) GoString() string {
	return s.String()
}

func (s *GoodsSpuOperateRequest) SetOutSpuId(v string) *GoodsSpuOperateRequest {
	s.OutSpuId = &v
	return s
}

func (s *GoodsSpuOperateRequest) SetHeader(v map[string]*string) *GoodsSpuOperateRequest {
	s.Header = v
	return s
}

func (s *GoodsSpuOperateRequest) SetAccessToken(v string) *GoodsSpuOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *GoodsSpuOperateRequest) SetSpuId(v int64) *GoodsSpuOperateRequest {
	s.SpuId = &v
	return s
}

func (s *GoodsSpuOperateRequest) SetAccountId(v int64) *GoodsSpuOperateRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsSpuOperateRequest) SetOp(v int) *GoodsSpuOperateRequest {
	s.Op = &v
	return s
}

type GoodsSpuOperateResponse struct {
	Extra *GoodsSpuOperateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *GoodsSpuOperateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GoodsSpuOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuOperateResponse) GoString() string {
	return s.String()
}

func (s *GoodsSpuOperateResponse) SetExtra(v *GoodsSpuOperateResponseExtra) *GoodsSpuOperateResponse {
	s.Extra = v
	return s
}

func (s *GoodsSpuOperateResponse) SetData(v *GoodsSpuOperateResponseData) *GoodsSpuOperateResponse {
	s.Data = v
	return s
}

type GoodsSpuOperateResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SpuId       *int64  `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsSpuOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuOperateResponseData) GoString() string {
	return s.String()
}

func (s *GoodsSpuOperateResponseData) SetErrorCode(v int32) *GoodsSpuOperateResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSpuOperateResponseData) SetSpuId(v int64) *GoodsSpuOperateResponseData {
	s.SpuId = &v
	return s
}

func (s *GoodsSpuOperateResponseData) SetDescription(v string) *GoodsSpuOperateResponseData {
	s.Description = &v
	return s
}

type GoodsSpuOperateResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s GoodsSpuOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsSpuOperateResponseExtra) SetNow(v int64) *GoodsSpuOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsSpuOperateResponseExtra) SetSubDescription(v string) *GoodsSpuOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsSpuOperateResponseExtra) SetSubErrorCode(v int32) *GoodsSpuOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsSpuOperateResponseExtra) SetDescription(v string) *GoodsSpuOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *GoodsSpuOperateResponseExtra) SetErrorCode(v int32) *GoodsSpuOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSpuOperateResponseExtra) SetLogid(v string) *GoodsSpuOperateResponseExtra {
	s.Logid = &v
	return s
}

type GoodsSpuSaveRequest struct {
	Spu         *GoodsSpuSaveRequestSpu `json:"spu,omitempty" xml:"spu,omitempty"`
	AccountId   *int64                  `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GoodsSpuSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuSaveRequest) GoString() string {
	return s.String()
}

func (s *GoodsSpuSaveRequest) SetSpu(v *GoodsSpuSaveRequestSpu) *GoodsSpuSaveRequest {
	s.Spu = v
	return s
}

func (s *GoodsSpuSaveRequest) SetAccountId(v int64) *GoodsSpuSaveRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsSpuSaveRequest) SetHeader(v map[string]*string) *GoodsSpuSaveRequest {
	s.Header = v
	return s
}

func (s *GoodsSpuSaveRequest) SetAccessToken(v string) *GoodsSpuSaveRequest {
	s.AccessToken = &v
	return s
}

type GoodsSpuSaveRequestSpu struct {
	StandardCategoryId *int64             `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty" require:"true"`
	AccountId          *int64             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutSpuId           *string            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty" require:"true"`
	SpuAttrMap         map[string]*string `json:"spu_attr_map,omitempty" xml:"spu_attr_map,omitempty"`
	SpuName            *string            `json:"spu_name,omitempty" xml:"spu_name,omitempty" require:"true"`
}

func (s GoodsSpuSaveRequestSpu) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuSaveRequestSpu) GoString() string {
	return s.String()
}

func (s *GoodsSpuSaveRequestSpu) SetStandardCategoryId(v int64) *GoodsSpuSaveRequestSpu {
	s.StandardCategoryId = &v
	return s
}

func (s *GoodsSpuSaveRequestSpu) SetAccountId(v int64) *GoodsSpuSaveRequestSpu {
	s.AccountId = &v
	return s
}

func (s *GoodsSpuSaveRequestSpu) SetOutSpuId(v string) *GoodsSpuSaveRequestSpu {
	s.OutSpuId = &v
	return s
}

func (s *GoodsSpuSaveRequestSpu) SetSpuAttrMap(v map[string]*string) *GoodsSpuSaveRequestSpu {
	s.SpuAttrMap = v
	return s
}

func (s *GoodsSpuSaveRequestSpu) SetSpuName(v string) *GoodsSpuSaveRequestSpu {
	s.SpuName = &v
	return s
}

type GoodsSpuSaveResponse struct {
	Extra *GoodsSpuSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *GoodsSpuSaveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GoodsSpuSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuSaveResponse) GoString() string {
	return s.String()
}

func (s *GoodsSpuSaveResponse) SetExtra(v *GoodsSpuSaveResponseExtra) *GoodsSpuSaveResponse {
	s.Extra = v
	return s
}

func (s *GoodsSpuSaveResponse) SetData(v *GoodsSpuSaveResponseData) *GoodsSpuSaveResponse {
	s.Data = v
	return s
}

type GoodsSpuSaveResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SpuId       *int64  `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsSpuSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuSaveResponseData) GoString() string {
	return s.String()
}

func (s *GoodsSpuSaveResponseData) SetErrorCode(v int32) *GoodsSpuSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSpuSaveResponseData) SetSpuId(v int64) *GoodsSpuSaveResponseData {
	s.SpuId = &v
	return s
}

func (s *GoodsSpuSaveResponseData) SetDescription(v string) *GoodsSpuSaveResponseData {
	s.Description = &v
	return s
}

type GoodsSpuSaveResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsSpuSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsSpuSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsSpuSaveResponseExtra) SetErrorCode(v int32) *GoodsSpuSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsSpuSaveResponseExtra) SetLogid(v string) *GoodsSpuSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *GoodsSpuSaveResponseExtra) SetNow(v int64) *GoodsSpuSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *GoodsSpuSaveResponseExtra) SetSubDescription(v string) *GoodsSpuSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsSpuSaveResponseExtra) SetSubErrorCode(v int32) *GoodsSpuSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsSpuSaveResponseExtra) SetDescription(v string) *GoodsSpuSaveResponseExtra {
	s.Description = &v
	return s
}

type GoodsTemplateGetRequest struct {
	OpenBizType    *int               `json:"open_biz_type,omitempty" xml:"open_biz_type,omitempty"`
	ProductSubType *int               `json:"product_sub_type,omitempty" xml:"product_sub_type,omitempty"`
	ProductType    *int               `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CategoryId     *string            `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
}

func (s GoodsTemplateGetRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetRequest) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetRequest) SetOpenBizType(v int) *GoodsTemplateGetRequest {
	s.OpenBizType = &v
	return s
}

func (s *GoodsTemplateGetRequest) SetProductSubType(v int) *GoodsTemplateGetRequest {
	s.ProductSubType = &v
	return s
}

func (s *GoodsTemplateGetRequest) SetProductType(v int) *GoodsTemplateGetRequest {
	s.ProductType = &v
	return s
}

func (s *GoodsTemplateGetRequest) SetHeader(v map[string]*string) *GoodsTemplateGetRequest {
	s.Header = v
	return s
}

func (s *GoodsTemplateGetRequest) SetAccessToken(v string) *GoodsTemplateGetRequest {
	s.AccessToken = &v
	return s
}

func (s *GoodsTemplateGetRequest) SetCategoryId(v string) *GoodsTemplateGetRequest {
	s.CategoryId = &v
	return s
}

type GoodsTemplateGetResponse struct {
	Data     *GoodsTemplateGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *GoodsTemplateGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *GoodsTemplateGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s GoodsTemplateGetResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponse) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponse) SetData(v *GoodsTemplateGetResponseData) *GoodsTemplateGetResponse {
	s.Data = v
	return s
}

func (s *GoodsTemplateGetResponse) SetExtra(v *GoodsTemplateGetResponseExtra) *GoodsTemplateGetResponse {
	s.Extra = v
	return s
}

func (s *GoodsTemplateGetResponse) SetBaseResp(v *GoodsTemplateGetResponseBaseResp) *GoodsTemplateGetResponse {
	s.BaseResp = v
	return s
}

type GoodsTemplateGetResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GoodsTemplateGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponseBaseResp) SetStatusMessage(v string) *GoodsTemplateGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *GoodsTemplateGetResponseBaseResp) SetExtra(v map[string]*string) *GoodsTemplateGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *GoodsTemplateGetResponseBaseResp) SetStatusCode(v int32) *GoodsTemplateGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

type GoodsTemplateGetResponseData struct {
	SkuAttrs     []*GoodsTemplateGetResponseDataSkuAttrsItem     `json:"sku_attrs,omitempty" xml:"sku_attrs,omitempty" type:"Repeated"`
	Description  *string                                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode    *int32                                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductAttrs []*GoodsTemplateGetResponseDataProductAttrsItem `json:"product_attrs,omitempty" xml:"product_attrs,omitempty" type:"Repeated"`
}

func (s GoodsTemplateGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponseData) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponseData) SetSkuAttrs(v []*GoodsTemplateGetResponseDataSkuAttrsItem) *GoodsTemplateGetResponseData {
	s.SkuAttrs = v
	return s
}

func (s *GoodsTemplateGetResponseData) SetDescription(v string) *GoodsTemplateGetResponseData {
	s.Description = &v
	return s
}

func (s *GoodsTemplateGetResponseData) SetErrorCode(v int32) *GoodsTemplateGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsTemplateGetResponseData) SetProductAttrs(v []*GoodsTemplateGetResponseDataProductAttrsItem) *GoodsTemplateGetResponseData {
	s.ProductAttrs = v
	return s
}

type GoodsTemplateGetResponseDataProductAttrsItem struct {
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	IsMulti    *bool   `json:"is_multi,omitempty" xml:"is_multi,omitempty"`
	IsRequired *bool   `json:"is_required,omitempty" xml:"is_required,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ValueDemo  *string `json:"value_demo,omitempty" xml:"value_demo,omitempty"`
	ValueType  *string `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s GoodsTemplateGetResponseDataProductAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponseDataProductAttrsItem) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetDesc(v string) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.Desc = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetIsMulti(v bool) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.IsMulti = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetIsRequired(v bool) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.IsRequired = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetKey(v string) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.Key = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetName(v string) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.Name = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetValueDemo(v string) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.ValueDemo = &v
	return s
}

func (s *GoodsTemplateGetResponseDataProductAttrsItem) SetValueType(v string) *GoodsTemplateGetResponseDataProductAttrsItem {
	s.ValueType = &v
	return s
}

type GoodsTemplateGetResponseDataSkuAttrsItem struct {
	ValueType  *string `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	IsMulti    *bool   `json:"is_multi,omitempty" xml:"is_multi,omitempty"`
	IsRequired *bool   `json:"is_required,omitempty" xml:"is_required,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ValueDemo  *string `json:"value_demo,omitempty" xml:"value_demo,omitempty"`
}

func (s GoodsTemplateGetResponseDataSkuAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponseDataSkuAttrsItem) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetValueType(v string) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.ValueType = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetDesc(v string) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.Desc = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetIsMulti(v bool) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.IsMulti = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetIsRequired(v bool) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.IsRequired = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetKey(v string) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.Key = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetName(v string) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.Name = &v
	return s
}

func (s *GoodsTemplateGetResponseDataSkuAttrsItem) SetValueDemo(v string) *GoodsTemplateGetResponseDataSkuAttrsItem {
	s.ValueDemo = &v
	return s
}

type GoodsTemplateGetResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s GoodsTemplateGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsTemplateGetResponseExtra) GoString() string {
	return s.String()
}

func (s *GoodsTemplateGetResponseExtra) SetSubDescription(v string) *GoodsTemplateGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GoodsTemplateGetResponseExtra) SetSubErrorCode(v int32) *GoodsTemplateGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GoodsTemplateGetResponseExtra) SetDescription(v string) *GoodsTemplateGetResponseExtra {
	s.Description = &v
	return s
}

func (s *GoodsTemplateGetResponseExtra) SetErrorCode(v int32) *GoodsTemplateGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GoodsTemplateGetResponseExtra) SetLogid(v string) *GoodsTemplateGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *GoodsTemplateGetResponseExtra) SetNow(v int64) *GoodsTemplateGetResponseExtra {
	s.Now = &v
	return s
}

type GroupCountRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s GroupCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCountRequest) GoString() string {
	return s.String()
}

func (s *GroupCountRequest) SetAccessToken(v string) *GroupCountRequest {
	s.AccessToken = &v
	return s
}

func (s *GroupCountRequest) SetOpenId(v string) *GroupCountRequest {
	s.OpenId = &v
	return s
}

func (s *GroupCountRequest) SetHeader(v map[string]*string) *GroupCountRequest {
	s.Header = v
	return s
}

type GroupCountResponse struct {
	GwExtra   *GroupCountResponseGwExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	LeftCount *int32                     `json:"left_count,omitempty" xml:"left_count,omitempty" require:"true"`
	Data      *GroupCountResponseData    `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GroupCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCountResponse) GoString() string {
	return s.String()
}

func (s *GroupCountResponse) SetGwExtra(v *GroupCountResponseGwExtra) *GroupCountResponse {
	s.GwExtra = v
	return s
}

func (s *GroupCountResponse) SetLeftCount(v int32) *GroupCountResponse {
	s.LeftCount = &v
	return s
}

func (s *GroupCountResponse) SetData(v *GroupCountResponseData) *GroupCountResponse {
	s.Data = v
	return s
}

type GroupCountResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GroupCountResponseData) String() string {
	return tea.Prettify(s)
}

func (s GroupCountResponseData) GoString() string {
	return s.String()
}

func (s *GroupCountResponseData) SetGwErrorCode(v int32) *GroupCountResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *GroupCountResponseData) SetGwDescription(v string) *GroupCountResponseData {
	s.GwDescription = &v
	return s
}

type GroupCountResponseGwExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GroupCountResponseGwExtra) String() string {
	return tea.Prettify(s)
}

func (s GroupCountResponseGwExtra) GoString() string {
	return s.String()
}

func (s *GroupCountResponseGwExtra) SetSubErrorCode(v int32) *GroupCountResponseGwExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GroupCountResponseGwExtra) SetSubDescription(v string) *GroupCountResponseGwExtra {
	s.SubDescription = &v
	return s
}

func (s *GroupCountResponseGwExtra) SetLogid(v string) *GroupCountResponseGwExtra {
	s.Logid = &v
	return s
}

func (s *GroupCountResponseGwExtra) SetNow(v int64) *GroupCountResponseGwExtra {
	s.Now = &v
	return s
}

func (s *GroupCountResponseGwExtra) SetErrorCode(v int32) *GroupCountResponseGwExtra {
	s.ErrorCode = &v
	return s
}

func (s *GroupCountResponseGwExtra) SetDescription(v string) *GroupCountResponseGwExtra {
	s.Description = &v
	return s
}

type HermesTradeOrderQueryRequest struct {
	PageNum              *int32             `json:"page_num,omitempty" xml:"page_num,omitempty"`
	UpdateOrderEndTime   *int64             `json:"update_order_end_time,omitempty" xml:"update_order_end_time,omitempty"`
	CreateOrderEndTime   *int64             `json:"create_order_end_time,omitempty" xml:"create_order_end_time,omitempty"`
	CreateOrderStartTime *int64             `json:"create_order_start_time,omitempty" xml:"create_order_start_time,omitempty"`
	GetSecretNumber      *bool              `json:"get_secret_number,omitempty" xml:"get_secret_number,omitempty"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId              *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	ExtOrderId           *string            `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	UpdateOrderStartTime *int64             `json:"update_order_start_time,omitempty" xml:"update_order_start_time,omitempty"`
	PageSize             *int32             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	OrderStatus          *int32             `json:"order_status,omitempty" xml:"order_status,omitempty"`
	AccountId            *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OpenId               *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s HermesTradeOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryRequest) SetPageNum(v int32) *HermesTradeOrderQueryRequest {
	s.PageNum = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetUpdateOrderEndTime(v int64) *HermesTradeOrderQueryRequest {
	s.UpdateOrderEndTime = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetCreateOrderEndTime(v int64) *HermesTradeOrderQueryRequest {
	s.CreateOrderEndTime = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetCreateOrderStartTime(v int64) *HermesTradeOrderQueryRequest {
	s.CreateOrderStartTime = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetGetSecretNumber(v bool) *HermesTradeOrderQueryRequest {
	s.GetSecretNumber = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetHeader(v map[string]*string) *HermesTradeOrderQueryRequest {
	s.Header = v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetAccessToken(v string) *HermesTradeOrderQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetOrderId(v string) *HermesTradeOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetExtOrderId(v string) *HermesTradeOrderQueryRequest {
	s.ExtOrderId = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetUpdateOrderStartTime(v int64) *HermesTradeOrderQueryRequest {
	s.UpdateOrderStartTime = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetPageSize(v int32) *HermesTradeOrderQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetOrderStatus(v int32) *HermesTradeOrderQueryRequest {
	s.OrderStatus = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetAccountId(v string) *HermesTradeOrderQueryRequest {
	s.AccountId = &v
	return s
}

func (s *HermesTradeOrderQueryRequest) SetOpenId(v string) *HermesTradeOrderQueryRequest {
	s.OpenId = &v
	return s
}

type HermesTradeOrderQueryResponse struct {
	Extra *HermesTradeOrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *HermesTradeOrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s HermesTradeOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponse) SetExtra(v *HermesTradeOrderQueryResponseExtra) *HermesTradeOrderQueryResponse {
	s.Extra = v
	return s
}

func (s *HermesTradeOrderQueryResponse) SetData(v *HermesTradeOrderQueryResponseData) *HermesTradeOrderQueryResponse {
	s.Data = v
	return s
}

type HermesTradeOrderQueryResponseData struct {
	Orders        []*HermesTradeOrderQueryResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
	Page          *HermesTradeOrderQueryResponseDataPage         `json:"page,omitempty" xml:"page,omitempty"`
	GwErrorCode   *int32                                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s HermesTradeOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseData) SetOrders(v []*HermesTradeOrderQueryResponseDataOrdersItem) *HermesTradeOrderQueryResponseData {
	s.Orders = v
	return s
}

func (s *HermesTradeOrderQueryResponseData) SetPage(v *HermesTradeOrderQueryResponseDataPage) *HermesTradeOrderQueryResponseData {
	s.Page = v
	return s
}

func (s *HermesTradeOrderQueryResponseData) SetGwErrorCode(v int32) *HermesTradeOrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *HermesTradeOrderQueryResponseData) SetGwDescription(v string) *HermesTradeOrderQueryResponseData {
	s.GwDescription = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItem struct {
	AmountInfo          *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo                `json:"amount_info,omitempty" xml:"amount_info,omitempty"`
	Products            []*HermesTradeOrderQueryResponseDataOrdersItemProductsItem            `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	DeliveryInfo        *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo              `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
	BuyerInfo           *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo                 `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
	Contacts            []*HermesTradeOrderQueryResponseDataOrdersItemContactsItem            `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	OrderId             *string                                                               `json:"order_id,omitempty" xml:"order_id,omitempty"`
	DiscountAmount      *int64                                                                `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	SubOrderAmountInfos []*HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem `json:"sub_order_amount_infos,omitempty" xml:"sub_order_amount_infos,omitempty" type:"Repeated"`
	PoiId               *string                                                               `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PayAmount           *int32                                                                `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	CreateOrderTime     *int64                                                                `json:"create_order_time,omitempty" xml:"create_order_time,omitempty"`
	PayTime             *int64                                                                `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	Poi                 *HermesTradeOrderQueryResponseDataOrdersItemPoi                       `json:"poi,omitempty" xml:"poi,omitempty"`
	UpdateOrderTime     *int64                                                                `json:"update_order_time,omitempty" xml:"update_order_time,omitempty"`
	Discounts           []*HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem           `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	IsDeliverLater      *bool                                                                 `json:"is_deliver_later,omitempty" xml:"is_deliver_later,omitempty"`
	SkuNum              *int32                                                                `json:"sku_num,omitempty" xml:"sku_num,omitempty"`
	ProductsAfterRefund []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem `json:"products_after_refund,omitempty" xml:"products_after_refund,omitempty" type:"Repeated"`
	MerchantInfo        *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo              `json:"merchant_info,omitempty" xml:"merchant_info,omitempty"`
	SkuTypeNum          *int32                                                                `json:"sku_type_num,omitempty" xml:"sku_type_num,omitempty"`
	ReceiverInfo        *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo              `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	OriginalAmount      *int32                                                                `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	RefundInfos         []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem         `json:"refund_infos,omitempty" xml:"refund_infos,omitempty" type:"Repeated"`
	OpenId              *string                                                               `json:"open_id,omitempty" xml:"open_id,omitempty"`
	PaymentDiscount     *int32                                                                `json:"payment_discount,omitempty" xml:"payment_discount,omitempty"`
	OrderType           *int32                                                                `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OrderStatus         *int32                                                                `json:"order_status,omitempty" xml:"order_status,omitempty"`
	Certificate         []*HermesTradeOrderQueryResponseDataOrdersItemCertificateItem         `json:"certificate,omitempty" xml:"certificate,omitempty" type:"Repeated"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetAmountInfo(v *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.AmountInfo = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetProducts(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.Products = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetDeliveryInfo(v *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.DeliveryInfo = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetBuyerInfo(v *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.BuyerInfo = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetContacts(v []*HermesTradeOrderQueryResponseDataOrdersItemContactsItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.Contacts = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetOrderId(v string) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.OrderId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetSubOrderAmountInfos(v []*HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.SubOrderAmountInfos = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetPoiId(v string) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.PoiId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetPayAmount(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.PayAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetCreateOrderTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.CreateOrderTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetPayTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.PayTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetPoi(v *HermesTradeOrderQueryResponseDataOrdersItemPoi) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.Poi = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetUpdateOrderTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.UpdateOrderTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetDiscounts(v []*HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.Discounts = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetIsDeliverLater(v bool) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.IsDeliverLater = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetSkuNum(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.SkuNum = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetProductsAfterRefund(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.ProductsAfterRefund = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetMerchantInfo(v *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.MerchantInfo = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetSkuTypeNum(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.SkuTypeNum = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetReceiverInfo(v *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.ReceiverInfo = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetOriginalAmount(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.OriginalAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetRefundInfos(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.RefundInfos = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetOpenId(v string) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.OpenId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetPaymentDiscount(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.PaymentDiscount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetOrderType(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.OrderType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetOrderStatus(v int32) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.OrderStatus = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItem) SetCertificate(v []*HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) *HermesTradeOrderQueryResponseDataOrdersItem {
	s.Certificate = v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemAmountInfo struct {
	CommissionAmount          *int64                                                                        `json:"commission_amount,omitempty" xml:"commission_amount,omitempty"`
	Discounts                 []*HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem         `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	PayDiscountAmount         *int64                                                                        `json:"pay_discount_amount,omitempty" xml:"pay_discount_amount,omitempty"`
	FreightPayAmount          *int64                                                                        `json:"freight_pay_amount,omitempty" xml:"freight_pay_amount,omitempty"`
	ActivitiesFeeAmount       *int64                                                                        `json:"activities_fee_amount,omitempty" xml:"activities_fee_amount,omitempty"`
	OriginAmount              *int64                                                                        `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	SalePrice                 *int64                                                                        `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	MerchantDiscountAmount    *int64                                                                        `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PackingPrice              *int64                                                                        `json:"packing_price,omitempty" xml:"packing_price,omitempty"`
	EstimatedOrderIncome      *int64                                                                        `json:"estimated_order_income,omitempty" xml:"estimated_order_income,omitempty"`
	PlatformDeliverFreightFee *int64                                                                        `json:"platform_deliver_freight_fee,omitempty" xml:"platform_deliver_freight_fee,omitempty"`
	DiscountAmount            *int64                                                                        `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DeductionInfoList         []*HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem `json:"deduction_info_list,omitempty" xml:"deduction_info_list,omitempty" type:"Repeated"`
	FreightActualPayAmount    *int64                                                                        `json:"freight_actual_pay_amount,omitempty" xml:"freight_actual_pay_amount,omitempty"`
	ProductOriginAmount       *int64                                                                        `json:"product_origin_amount,omitempty" xml:"product_origin_amount,omitempty"`
	MerchantDeliverFreightFee *int64                                                                        `json:"merchant_deliver_freight_fee,omitempty" xml:"merchant_deliver_freight_fee,omitempty"`
	OtherFeeAmount            *int64                                                                        `json:"other_fee_amount,omitempty" xml:"other_fee_amount,omitempty"`
	PlatformDiscountAmount    *int64                                                                        `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	PayAmount                 *int64                                                                        `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetCommissionAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.CommissionAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetDiscounts(v []*HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.Discounts = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetPayDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetFreightPayAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.FreightPayAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetActivitiesFeeAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.ActivitiesFeeAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.OriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetSalePrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.SalePrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetMerchantDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetPackingPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PackingPrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetEstimatedOrderIncome(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.EstimatedOrderIncome = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetPlatformDeliverFreightFee(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PlatformDeliverFreightFee = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetDeductionInfoList(v []*HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.DeductionInfoList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetFreightActualPayAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.FreightActualPayAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetProductOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.ProductOriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetMerchantDeliverFreightFee(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.MerchantDeliverFreightFee = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetOtherFeeAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.OtherFeeAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetPlatformDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo) SetPayAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayAmount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem struct {
	CertificateId          *int64 `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	DeductionAmount        *int64 `json:"deduction_amount,omitempty" xml:"deduction_amount,omitempty"`
	DeductionType          *int   `json:"deduction_type,omitempty" xml:"deduction_type,omitempty"`
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PayAmount              *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	SalePrice              *int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	SerialNum              *int64 `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetCertificateId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetDeductionAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.DeductionAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetDeductionType(v int) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.DeductionType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetMerchantDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetPayAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.PayAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetSalePrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.SalePrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem) SetSerialNum(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDeductionInfoListItem {
	s.SerialNum = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem struct {
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	DiscountAmount         *int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountType           *int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) SetMerchantDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) SetPlatformDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem) SetDiscountType(v int64) *HermesTradeOrderQueryResponseDataOrdersItemAmountInfoDiscountsItem {
	s.DiscountType = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo struct {
	BuyerSecretNumber *string `json:"buyer_secret_number,omitempty" xml:"buyer_secret_number,omitempty"`
	BuyerPhone        *string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	BuyerRealPhone    *string `json:"buyer_real_phone,omitempty" xml:"buyer_real_phone,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) SetBuyerSecretNumber(v string) *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo {
	s.BuyerSecretNumber = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) SetBuyerPhone(v string) *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo {
	s.BuyerPhone = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo) SetBuyerRealPhone(v string) *HermesTradeOrderQueryResponseDataOrdersItemBuyerInfo {
	s.BuyerRealPhone = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemCertificateItem struct {
	RefundAmount   *int32  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	RefundTime     *int64  `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	CertificateId  *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ItemStatus     *int32  `json:"item_status,omitempty" xml:"item_status,omitempty"`
	ItemUpdateTime *int64  `json:"item_update_time,omitempty" xml:"item_update_time,omitempty"`
	OrderItemId    *string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetRefundAmount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetRefundTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetCertificateId(v string) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.CertificateId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetItemStatus(v int32) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemStatus = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetItemUpdateTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemUpdateTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem) SetOrderItemId(v string) *HermesTradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.OrderItemId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemContactsItem struct {
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone        *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PhoneEncrypt *string `json:"phone_encrypt,omitempty" xml:"phone_encrypt,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemContactsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemContactsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemContactsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemContactsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemContactsItem) SetPhone(v string) *HermesTradeOrderQueryResponseDataOrdersItemContactsItem {
	s.Phone = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemContactsItem) SetPhoneEncrypt(v string) *HermesTradeOrderQueryResponseDataOrdersItemContactsItem {
	s.PhoneEncrypt = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo struct {
	IsBook         *bool   `json:"is_book,omitempty" xml:"is_book,omitempty"`
	Remark         *string `json:"remark,omitempty" xml:"remark,omitempty"`
	ShopNumber     *string `json:"shop_number,omitempty" xml:"shop_number,omitempty"`
	SysExpectTime  *string `json:"sys_expect_time,omitempty" xml:"sys_expect_time,omitempty"`
	TableWare      *string `json:"table_ware,omitempty" xml:"table_ware,omitempty"`
	UserExpectTime *string `json:"user_expect_time,omitempty" xml:"user_expect_time,omitempty"`
	DeliverModel   *int32  `json:"deliver_model,omitempty" xml:"deliver_model,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetIsBook(v bool) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.IsBook = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetRemark(v string) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.Remark = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetShopNumber(v string) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.ShopNumber = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetSysExpectTime(v string) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.SysExpectTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetTableWare(v string) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.TableWare = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetUserExpectTime(v string) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.UserExpectTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetDeliverModel(v int32) *HermesTradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.DeliverModel = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem struct {
	DiscountAmount         *int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountType           *int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountType(v int64) *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) SetMerchantDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem) SetPlatformDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo struct {
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountId(v string) *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountName(v string) *HermesTradeOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountName = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemPoi struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemPoi) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemPoi) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemPoi) SetPoiId(v string) *HermesTradeOrderQueryResponseDataOrdersItemPoi {
	s.PoiId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemPoi) SetPoiName(v string) *HermesTradeOrderQueryResponseDataOrdersItemPoi {
	s.PoiName = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem struct {
	Attributes          []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem  `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	ProductOriginAmount *int64                                                                               `json:"product_origin_amount,omitempty" xml:"product_origin_amount,omitempty"`
	OutId               *string                                                                              `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Ingredients         []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem `json:"ingredients,omitempty" xml:"ingredients,omitempty" type:"Repeated"`
	SkuId               *string                                                                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OriginAmount        *int64                                                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ProductName         *string                                                                              `json:"product_name,omitempty" xml:"product_name,omitempty"`
	Num                 *int32                                                                               `json:"num,omitempty" xml:"num,omitempty"`
	Specs               []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem       `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	ProductType         *int32                                                                               `json:"product_type,omitempty" xml:"product_type,omitempty"`
	Commodities         []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	ProductId           *string                                                                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetAttributes(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.Attributes = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetProductOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.ProductOriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetIngredients(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.Ingredients = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.OriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetProductName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.ProductName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetNum(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.Num = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetSpecs(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.Specs = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetProductType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.ProductType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetCommodities(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.Commodities = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemAttributesItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem struct {
	GroupName   *string                                                                                          `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList    []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	OptionCount *int32                                                                                           `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32                                                                                           `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) SetGroupName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem {
	s.GroupName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) SetItemList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem {
	s.ItemList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) SetOptionCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem {
	s.OptionCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem) SetTotalCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItem {
	s.TotalCount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem struct {
	WeightUnit        *string                                                                                                               `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	SkuAttrList       []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem       `json:"sku_attr_list,omitempty" xml:"sku_attr_list,omitempty" type:"Repeated"`
	OriginalSkuId     *string                                                                                                               `json:"original_sku_id,omitempty" xml:"original_sku_id,omitempty"`
	Ingredients       []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem       `json:"ingredients,omitempty" xml:"ingredients,omitempty" type:"Repeated"`
	AttrList          []*string                                                                                                             `json:"attr_list,omitempty" xml:"attr_list,omitempty" type:"Repeated"`
	SkuId             *string                                                                                                               `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Specs             []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem             `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	DeductionInfoList []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem `json:"deduction_info_list,omitempty" xml:"deduction_info_list,omitempty" type:"Repeated"`
	Unit              *string                                                                                                               `json:"unit,omitempty" xml:"unit,omitempty"`
	Name              *string                                                                                                               `json:"name,omitempty" xml:"name,omitempty"`
	WeightCount       *int64                                                                                                                `json:"weight_count,omitempty" xml:"weight_count,omitempty"`
	ItemTag           *string                                                                                                               `json:"item_tag,omitempty" xml:"item_tag,omitempty"`
	OutId             *string                                                                                                               `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Count             *int64                                                                                                                `json:"count,omitempty" xml:"count,omitempty"`
	OriginalPrice     *int64                                                                                                                `json:"original_price,omitempty" xml:"original_price,omitempty"`
	OriginalOutSkuId  *string                                                                                                               `json:"original_out_sku_id,omitempty" xml:"original_out_sku_id,omitempty"`
	ProductId         *string                                                                                                               `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Attributes        []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem        `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	Desc              *string                                                                                                               `json:"desc,omitempty" xml:"desc,omitempty"`
	OutSkuId          *string                                                                                                               `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price             *int64                                                                                                                `json:"price,omitempty" xml:"price,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetWeightUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.WeightUnit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetSkuAttrList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.SkuAttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetOriginalSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.OriginalSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetIngredients(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Ingredients = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetAttrList(v []*string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.AttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetSpecs(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Specs = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetDeductionInfoList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.DeductionInfoList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Unit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetWeightCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.WeightCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetItemTag(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.ItemTag = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetOriginalPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.OriginalPrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetOriginalOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.OriginalOutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetAttributes(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Attributes = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetDesc(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Desc = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem) SetPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItem {
	s.Price = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemAttributesItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem struct {
	VerifyId      *int64  `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	Amount        *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	CertificateId *int64  `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ThirdPartCode *string `json:"third_part_code,omitempty" xml:"third_part_code,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetVerifyId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.VerifyId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.Amount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetCertificateId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetThirdPartCode(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.ThirdPartCode = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem struct {
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Count     *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) SetCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem struct {
	UpgradePrice *int64  `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
	AttrName     *string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	Num          *int64  `json:"num,omitempty" xml:"num,omitempty"`
	OutId        *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId    *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetUpgradePrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.UpgradePrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetAttrName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.AttrName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetNum(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.Num = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem struct {
	ValueType   *int32  `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId    *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueMetric *string `json:"value_metric,omitempty" xml:"value_metric,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetValueType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.ValueType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.OutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem) SetValueMetric(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemCommoditiesItemItemListItemSpecsItem {
	s.ValueMetric = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem struct {
	Count     *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) SetCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemIngredientsItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem struct {
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueMetric *string `json:"value_metric,omitempty" xml:"value_metric,omitempty"`
	ValueType   *int32  `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId    *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetValueMetric(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.ValueMetric = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetValueType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.ValueType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsAfterRefundItemSpecsItem {
	s.OutSkuId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItem struct {
	ProductId           *string                                                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Num                 *int32                                                                    `json:"num,omitempty" xml:"num,omitempty"`
	ProductType         *int32                                                                    `json:"product_type,omitempty" xml:"product_type,omitempty"`
	Commodities         []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	ProductName         *string                                                                   `json:"product_name,omitempty" xml:"product_name,omitempty"`
	SkuId               *string                                                                   `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	ProductOriginAmount *int64                                                                    `json:"product_origin_amount,omitempty" xml:"product_origin_amount,omitempty"`
	OutId               *string                                                                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	OriginAmount        *int64                                                                    `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetNum(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.Num = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetProductType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetCommodities(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.Commodities = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetProductName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetProductOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductOriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItem) SetOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItem {
	s.OriginAmount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem struct {
	GroupName   *string                                                                               `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList    []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	OptionCount *int32                                                                                `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32                                                                                `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) SetGroupName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem {
	s.GroupName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) SetItemList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem {
	s.ItemList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) SetOptionCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem {
	s.OptionCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem) SetTotalCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItem {
	s.TotalCount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem struct {
	ProductId         *string                                                                                                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Desc              *string                                                                                                    `json:"desc,omitempty" xml:"desc,omitempty"`
	OriginalOutSkuId  *string                                                                                                    `json:"original_out_sku_id,omitempty" xml:"original_out_sku_id,omitempty"`
	Unit              *string                                                                                                    `json:"unit,omitempty" xml:"unit,omitempty"`
	OriginalSkuId     *string                                                                                                    `json:"original_sku_id,omitempty" xml:"original_sku_id,omitempty"`
	AttrList          []*string                                                                                                  `json:"attr_list,omitempty" xml:"attr_list,omitempty" type:"Repeated"`
	Price             *int64                                                                                                     `json:"price,omitempty" xml:"price,omitempty"`
	Name              *string                                                                                                    `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId          *string                                                                                                    `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	WeightCount       *int64                                                                                                     `json:"weight_count,omitempty" xml:"weight_count,omitempty"`
	OriginalPrice     *int64                                                                                                     `json:"original_price,omitempty" xml:"original_price,omitempty"`
	Specs             []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem             `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	WeightUnit        *string                                                                                                    `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	OutId             *string                                                                                                    `json:"out_id,omitempty" xml:"out_id,omitempty"`
	SkuAttrList       []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem       `json:"sku_attr_list,omitempty" xml:"sku_attr_list,omitempty" type:"Repeated"`
	DeductionInfoList []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem `json:"deduction_info_list,omitempty" xml:"deduction_info_list,omitempty" type:"Repeated"`
	ItemTag           *string                                                                                                    `json:"item_tag,omitempty" xml:"item_tag,omitempty"`
	Count             *int64                                                                                                     `json:"count,omitempty" xml:"count,omitempty"`
	SkuId             *string                                                                                                    `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Attributes        []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem        `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	Ingredients       []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem       `json:"ingredients,omitempty" xml:"ingredients,omitempty" type:"Repeated"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetDesc(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Desc = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetOriginalOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.OriginalOutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Unit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetOriginalSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.OriginalSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetAttrList(v []*string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.AttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Price = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetWeightCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.WeightCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetOriginalPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.OriginalPrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetSpecs(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Specs = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetWeightUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.WeightUnit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetSkuAttrList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.SkuAttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetDeductionInfoList(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.DeductionInfoList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetItemTag(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.ItemTag = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetAttributes(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Attributes = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem) SetIngredients(v []*HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItem {
	s.Ingredients = v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemAttributesItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem struct {
	CertificateId *int64  `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ThirdPartCode *string `json:"third_part_code,omitempty" xml:"third_part_code,omitempty"`
	VerifyId      *int64  `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	Amount        *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) SetCertificateId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) SetThirdPartCode(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.ThirdPartCode = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) SetVerifyId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.VerifyId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem) SetAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.Amount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem struct {
	Count     *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) SetCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemIngredientsItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem struct {
	AttrName     *string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	Num          *int64  `json:"num,omitempty" xml:"num,omitempty"`
	OutId        *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId    *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	UpgradePrice *int64  `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) SetAttrName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem {
	s.AttrName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) SetNum(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem {
	s.Num = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem) SetUpgradePrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSkuAttrListItem {
	s.UpgradePrice = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem struct {
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueMetric *string `json:"value_metric,omitempty" xml:"value_metric,omitempty"`
	ValueType   *int32  `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId    *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetValueMetric(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.ValueMetric = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetValueType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.ValueType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemProductsItemCommoditiesItemItemListItemSpecsItem {
	s.OutSkuId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo struct {
	Lat               *float64 `json:"lat,omitempty" xml:"lat,omitempty"`
	District          *string  `json:"district,omitempty" xml:"district,omitempty"`
	Lng               *float64 `json:"lng,omitempty" xml:"lng,omitempty"`
	DoorPlateNum      *string  `json:"door_plate_num,omitempty" xml:"door_plate_num,omitempty"`
	LocationAddress   *string  `json:"location_address,omitempty" xml:"location_address,omitempty"`
	ReceiverRealPhone *string  `json:"receiver_real_phone,omitempty" xml:"receiver_real_phone,omitempty"`
	Province          *string  `json:"province,omitempty" xml:"province,omitempty"`
	City              *string  `json:"city,omitempty" xml:"city,omitempty"`
	SecretNumber      *string  `json:"secret_number,omitempty" xml:"secret_number,omitempty"`
	LocationName      *string  `json:"location_name,omitempty" xml:"location_name,omitempty"`
	ReceiverPhone     *string  `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	Town              *string  `json:"town,omitempty" xml:"town,omitempty"`
	ReceiverName      *string  `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetLat(v float64) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.Lat = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetDistrict(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.District = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetLng(v float64) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.Lng = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetDoorPlateNum(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.DoorPlateNum = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetLocationAddress(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.LocationAddress = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetReceiverRealPhone(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.ReceiverRealPhone = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetProvince(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.Province = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetCity(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.City = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetSecretNumber(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.SecretNumber = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetLocationName(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.LocationName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetReceiverPhone(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.ReceiverPhone = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetTown(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.Town = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo) SetReceiverName(v string) *HermesTradeOrderQueryResponseDataOrdersItemReceiverInfo {
	s.ReceiverName = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem struct {
	RefundCompleteTime *int64                                                                          `json:"refund_complete_time,omitempty" xml:"refund_complete_time,omitempty"`
	UserRefundAmount   *int64                                                                          `json:"user_refund_amount,omitempty" xml:"user_refund_amount,omitempty"`
	AfterSaleId        *string                                                                         `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
	ProductsRefund     []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem `json:"products_refund,omitempty" xml:"products_refund,omitempty" type:"Repeated"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) SetRefundCompleteTime(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem {
	s.RefundCompleteTime = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) SetUserRefundAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem {
	s.UserRefundAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) SetAfterSaleId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem {
	s.AfterSaleId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem) SetProductsRefund(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItem {
	s.ProductsRefund = v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem struct {
	Specs               []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem       `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	OutId               *string                                                                                        `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Ingredients         []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem `json:"ingredients,omitempty" xml:"ingredients,omitempty" type:"Repeated"`
	ProductName         *string                                                                                        `json:"product_name,omitempty" xml:"product_name,omitempty"`
	Num                 *int32                                                                                         `json:"num,omitempty" xml:"num,omitempty"`
	Commodities         []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	ProductId           *string                                                                                        `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OriginAmount        *int64                                                                                         `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ProductType         *int32                                                                                         `json:"product_type,omitempty" xml:"product_type,omitempty"`
	SkuId               *string                                                                                        `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	ProductOriginAmount *int64                                                                                         `json:"product_origin_amount,omitempty" xml:"product_origin_amount,omitempty"`
	Attributes          []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem  `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetSpecs(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.Specs = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetIngredients(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.Ingredients = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetProductName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.ProductName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetNum(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.Num = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetCommodities(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.Commodities = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.OriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetProductType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.ProductType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetProductOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.ProductOriginAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem) SetAttributes(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItem {
	s.Attributes = v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemAttributesItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem struct {
	TotalCount  *int32                                                                                                     `json:"total_count,omitempty" xml:"total_count,omitempty"`
	GroupName   *string                                                                                                    `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList    []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	OptionCount *int32                                                                                                     `json:"option_count,omitempty" xml:"option_count,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) SetTotalCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem {
	s.TotalCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) SetGroupName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem {
	s.GroupName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) SetItemList(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem {
	s.ItemList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem) SetOptionCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItem {
	s.OptionCount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem struct {
	SkuId             *string                                                                                                                         `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Specs             []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem             `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	Desc              *string                                                                                                                         `json:"desc,omitempty" xml:"desc,omitempty"`
	ItemTag           *string                                                                                                                         `json:"item_tag,omitempty" xml:"item_tag,omitempty"`
	OriginalPrice     *int64                                                                                                                          `json:"original_price,omitempty" xml:"original_price,omitempty"`
	Price             *int64                                                                                                                          `json:"price,omitempty" xml:"price,omitempty"`
	Name              *string                                                                                                                         `json:"name,omitempty" xml:"name,omitempty"`
	Ingredients       []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem       `json:"ingredients,omitempty" xml:"ingredients,omitempty" type:"Repeated"`
	Attributes        []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem        `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	DeductionInfoList []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem `json:"deduction_info_list,omitempty" xml:"deduction_info_list,omitempty" type:"Repeated"`
	OutId             *string                                                                                                                         `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Unit              *string                                                                                                                         `json:"unit,omitempty" xml:"unit,omitempty"`
	WeightUnit        *string                                                                                                                         `json:"weight_unit,omitempty" xml:"weight_unit,omitempty"`
	OutSkuId          *string                                                                                                                         `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuAttrList       []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem       `json:"sku_attr_list,omitempty" xml:"sku_attr_list,omitempty" type:"Repeated"`
	OriginalOutSkuId  *string                                                                                                                         `json:"original_out_sku_id,omitempty" xml:"original_out_sku_id,omitempty"`
	AttrList          []*string                                                                                                                       `json:"attr_list,omitempty" xml:"attr_list,omitempty" type:"Repeated"`
	Count             *int64                                                                                                                          `json:"count,omitempty" xml:"count,omitempty"`
	OriginalSkuId     *string                                                                                                                         `json:"original_sku_id,omitempty" xml:"original_sku_id,omitempty"`
	WeightCount       *int64                                                                                                                          `json:"weight_count,omitempty" xml:"weight_count,omitempty"`
	ProductId         *string                                                                                                                         `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetSpecs(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Specs = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetDesc(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Desc = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetItemTag(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.ItemTag = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetOriginalPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.OriginalPrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetPrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Price = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetIngredients(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Ingredients = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetAttributes(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Attributes = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetDeductionInfoList(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.DeductionInfoList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Unit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetWeightUnit(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.WeightUnit = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetSkuAttrList(v []*HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.SkuAttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetOriginalOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.OriginalOutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetAttrList(v []*string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.AttrList = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetOriginalSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.OriginalSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetWeightCount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.WeightCount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemAttributesItem {
	s.Name = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem struct {
	CertificateId *int64  `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ThirdPartCode *string `json:"third_part_code,omitempty" xml:"third_part_code,omitempty"`
	VerifyId      *int64  `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	Amount        *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetCertificateId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetThirdPartCode(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.ThirdPartCode = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetVerifyId(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.VerifyId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem) SetAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemDeductionInfoListItem {
	s.Amount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Count     *int32  `json:"count,omitempty" xml:"count,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem) SetCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemIngredientsItem {
	s.Count = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem struct {
	OutId        *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId    *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	UpgradePrice *int64  `json:"upgrade_price,omitempty" xml:"upgrade_price,omitempty"`
	AttrName     *string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	Num          *int64  `json:"num,omitempty" xml:"num,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.ProductId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetUpgradePrice(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.UpgradePrice = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetAttrName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.AttrName = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem) SetNum(v int64) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSkuAttrListItem {
	s.Num = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem struct {
	ValueType   *int32  `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId    *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueMetric *string `json:"value_metric,omitempty" xml:"value_metric,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetValueType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.ValueType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.OutSkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem) SetValueMetric(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemCommoditiesItemItemListItemSpecsItem {
	s.ValueMetric = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem struct {
	Count     *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) SetCount(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem {
	s.Count = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) SetOutId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem {
	s.OutId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem) SetProductId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemIngredientsItem {
	s.ProductId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem struct {
	SkuId       *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Value       *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueMetric *string `json:"value_metric,omitempty" xml:"value_metric,omitempty"`
	ValueType   *int32  `json:"value_type,omitempty" xml:"value_type,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OutSkuId    *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.SkuId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetValue(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.Value = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetValueMetric(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.ValueMetric = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetValueType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.ValueType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetName(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.Name = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem) SetOutSkuId(v string) *HermesTradeOrderQueryResponseDataOrdersItemRefundInfosItemProductsRefundItemSpecsItem {
	s.OutSkuId = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem struct {
	PayAmount      *int64                                                                             `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	SubOrderId     *string                                                                            `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	SubOrderType   *int32                                                                             `json:"sub_order_type,omitempty" xml:"sub_order_type,omitempty"`
	DiscountAmount *int64                                                                             `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	Discounts      []*HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	OriginAmount   *int64                                                                             `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetPayAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.PayAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderId(v string) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderId = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderType(v int32) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderType = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscounts(v []*HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.Discounts = v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetOriginAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.OriginAmount = &v
	return s
}

type HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem struct {
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	DiscountAmount         *int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountType           *int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetMerchantDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetPlatformDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountAmount(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountType(v int64) *HermesTradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountType = &v
	return s
}

type HermesTradeOrderQueryResponseDataPage struct {
	PageNum  *int32 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Total    *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s HermesTradeOrderQueryResponseDataPage) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseDataPage) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseDataPage) SetPageNum(v int32) *HermesTradeOrderQueryResponseDataPage {
	s.PageNum = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataPage) SetPageSize(v int32) *HermesTradeOrderQueryResponseDataPage {
	s.PageSize = &v
	return s
}

func (s *HermesTradeOrderQueryResponseDataPage) SetTotal(v int64) *HermesTradeOrderQueryResponseDataPage {
	s.Total = &v
	return s
}

type HermesTradeOrderQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s HermesTradeOrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s HermesTradeOrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *HermesTradeOrderQueryResponseExtra) SetNow(v int64) *HermesTradeOrderQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *HermesTradeOrderQueryResponseExtra) SetSubDescription(v string) *HermesTradeOrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *HermesTradeOrderQueryResponseExtra) SetSubErrorCode(v int32) *HermesTradeOrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *HermesTradeOrderQueryResponseExtra) SetDescription(v string) *HermesTradeOrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *HermesTradeOrderQueryResponseExtra) SetErrorCode(v int32) *HermesTradeOrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *HermesTradeOrderQueryResponseExtra) SetLogid(v string) *HermesTradeOrderQueryResponseExtra {
	s.Logid = &v
	return s
}

type HotelOrderConfirmRequest struct {
	OrderId       *string                                `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ConfirmResult *HotelOrderConfirmRequestConfirmResult `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
	Header        map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s HotelOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderConfirmRequest) SetOrderId(v string) *HotelOrderConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *HotelOrderConfirmRequest) SetConfirmResult(v *HotelOrderConfirmRequestConfirmResult) *HotelOrderConfirmRequest {
	s.ConfirmResult = v
	return s
}

func (s *HotelOrderConfirmRequest) SetHeader(v map[string]*string) *HotelOrderConfirmRequest {
	s.Header = v
	return s
}

func (s *HotelOrderConfirmRequest) SetAccessToken(v string) *HotelOrderConfirmRequest {
	s.AccessToken = &v
	return s
}

type HotelOrderConfirmRequestConfirmResult struct {
	ConfirmNumber *string `json:"confirm_number,omitempty" xml:"confirm_number,omitempty"`
	ConfirmResult *int    `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
	RejectCode    *int    `json:"reject_code,omitempty" xml:"reject_code,omitempty"`
	RejectReason  *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s HotelOrderConfirmRequestConfirmResult) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderConfirmRequestConfirmResult) GoString() string {
	return s.String()
}

func (s *HotelOrderConfirmRequestConfirmResult) SetConfirmNumber(v string) *HotelOrderConfirmRequestConfirmResult {
	s.ConfirmNumber = &v
	return s
}

func (s *HotelOrderConfirmRequestConfirmResult) SetConfirmResult(v int) *HotelOrderConfirmRequestConfirmResult {
	s.ConfirmResult = &v
	return s
}

func (s *HotelOrderConfirmRequestConfirmResult) SetRejectCode(v int) *HotelOrderConfirmRequestConfirmResult {
	s.RejectCode = &v
	return s
}

func (s *HotelOrderConfirmRequestConfirmResult) SetRejectReason(v string) *HotelOrderConfirmRequestConfirmResult {
	s.RejectReason = &v
	return s
}

type HotelOrderConfirmResponse struct {
	Extra *HotelOrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *HotelOrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s HotelOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderConfirmResponse) SetExtra(v *HotelOrderConfirmResponseExtra) *HotelOrderConfirmResponse {
	s.Extra = v
	return s
}

func (s *HotelOrderConfirmResponse) SetData(v *HotelOrderConfirmResponseData) *HotelOrderConfirmResponse {
	s.Data = v
	return s
}

type HotelOrderConfirmResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s HotelOrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *HotelOrderConfirmResponseData) SetGwErrorCode(v int32) *HotelOrderConfirmResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *HotelOrderConfirmResponseData) SetGwDescription(v string) *HotelOrderConfirmResponseData {
	s.GwDescription = &v
	return s
}

type HotelOrderConfirmResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s HotelOrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s HotelOrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *HotelOrderConfirmResponseExtra) SetDescription(v string) *HotelOrderConfirmResponseExtra {
	s.Description = &v
	return s
}

func (s *HotelOrderConfirmResponseExtra) SetErrorCode(v int32) *HotelOrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *HotelOrderConfirmResponseExtra) SetLogid(v string) *HotelOrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *HotelOrderConfirmResponseExtra) SetNow(v int64) *HotelOrderConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *HotelOrderConfirmResponseExtra) SetSubDescription(v string) *HotelOrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *HotelOrderConfirmResponseExtra) SetSubErrorCode(v int32) *HotelOrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

type HotelPoiQueryRequest struct {
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PageIndex   *int32             `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s HotelPoiQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryRequest) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryRequest) SetAccountId(v string) *HotelPoiQueryRequest {
	s.AccountId = &v
	return s
}

func (s *HotelPoiQueryRequest) SetPageIndex(v int32) *HotelPoiQueryRequest {
	s.PageIndex = &v
	return s
}

func (s *HotelPoiQueryRequest) SetPageSize(v int32) *HotelPoiQueryRequest {
	s.PageSize = &v
	return s
}

func (s *HotelPoiQueryRequest) SetHeader(v map[string]*string) *HotelPoiQueryRequest {
	s.Header = v
	return s
}

func (s *HotelPoiQueryRequest) SetAccessToken(v string) *HotelPoiQueryRequest {
	s.AccessToken = &v
	return s
}

type HotelPoiQueryResponse struct {
	Extra *HotelPoiQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *HotelPoiQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s HotelPoiQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponse) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponse) SetExtra(v *HotelPoiQueryResponseExtra) *HotelPoiQueryResponse {
	s.Extra = v
	return s
}

func (s *HotelPoiQueryResponse) SetData(v *HotelPoiQueryResponseData) *HotelPoiQueryResponse {
	s.Data = v
	return s
}

type HotelPoiQueryResponseData struct {
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	HotelList     []*HotelPoiQueryResponseDataHotelListItem `json:"hotel_list,omitempty" xml:"hotel_list,omitempty" type:"Repeated"`
	Pagination    *HotelPoiQueryResponseDataPagination      `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

func (s HotelPoiQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponseData) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponseData) SetGwErrorCode(v int32) *HotelPoiQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *HotelPoiQueryResponseData) SetGwDescription(v string) *HotelPoiQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *HotelPoiQueryResponseData) SetHotelList(v []*HotelPoiQueryResponseDataHotelListItem) *HotelPoiQueryResponseData {
	s.HotelList = v
	return s
}

func (s *HotelPoiQueryResponseData) SetPagination(v *HotelPoiQueryResponseDataPagination) *HotelPoiQueryResponseData {
	s.Pagination = v
	return s
}

type HotelPoiQueryResponseDataHotelListItem struct {
	CategoryId *int64                                         `json:"category_id,omitempty" xml:"category_id,omitempty"`
	HotelId    *string                                        `json:"hotel_id,omitempty" xml:"hotel_id,omitempty" require:"true"`
	HotelName  *string                                        `json:"hotel_name,omitempty" xml:"hotel_name,omitempty" require:"true"`
	Latitude   *float64                                       `json:"latitude,omitempty" xml:"latitude,omitempty" require:"true"`
	Longitude  *float64                                       `json:"longitude,omitempty" xml:"longitude,omitempty" require:"true"`
	AccountId  *string                                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Address    *HotelPoiQueryResponseDataHotelListItemAddress `json:"address,omitempty" xml:"address,omitempty" require:"true"`
}

func (s HotelPoiQueryResponseDataHotelListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponseDataHotelListItem) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetCategoryId(v int64) *HotelPoiQueryResponseDataHotelListItem {
	s.CategoryId = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetHotelId(v string) *HotelPoiQueryResponseDataHotelListItem {
	s.HotelId = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetHotelName(v string) *HotelPoiQueryResponseDataHotelListItem {
	s.HotelName = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetLatitude(v float64) *HotelPoiQueryResponseDataHotelListItem {
	s.Latitude = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetLongitude(v float64) *HotelPoiQueryResponseDataHotelListItem {
	s.Longitude = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetAccountId(v string) *HotelPoiQueryResponseDataHotelListItem {
	s.AccountId = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItem) SetAddress(v *HotelPoiQueryResponseDataHotelListItemAddress) *HotelPoiQueryResponseDataHotelListItem {
	s.Address = v
	return s
}

type HotelPoiQueryResponseDataHotelListItemAddress struct {
	District      *string `json:"district,omitempty" xml:"district,omitempty"`
	Province      *string `json:"province,omitempty" xml:"province,omitempty"`
	City          *string `json:"city,omitempty" xml:"city,omitempty"`
	Country       *string `json:"country,omitempty" xml:"country,omitempty"`
	DetailAddress *string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
}

func (s HotelPoiQueryResponseDataHotelListItemAddress) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponseDataHotelListItemAddress) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponseDataHotelListItemAddress) SetDistrict(v string) *HotelPoiQueryResponseDataHotelListItemAddress {
	s.District = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItemAddress) SetProvince(v string) *HotelPoiQueryResponseDataHotelListItemAddress {
	s.Province = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItemAddress) SetCity(v string) *HotelPoiQueryResponseDataHotelListItemAddress {
	s.City = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItemAddress) SetCountry(v string) *HotelPoiQueryResponseDataHotelListItemAddress {
	s.Country = &v
	return s
}

func (s *HotelPoiQueryResponseDataHotelListItemAddress) SetDetailAddress(v string) *HotelPoiQueryResponseDataHotelListItemAddress {
	s.DetailAddress = &v
	return s
}

type HotelPoiQueryResponseDataPagination struct {
	HasMore    *bool  `json:"has_more,omitempty" xml:"has_more,omitempty"`
	PageCount  *int32 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	PageIndex  *int32 `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize   *int32 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
}

func (s HotelPoiQueryResponseDataPagination) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponseDataPagination) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponseDataPagination) SetHasMore(v bool) *HotelPoiQueryResponseDataPagination {
	s.HasMore = &v
	return s
}

func (s *HotelPoiQueryResponseDataPagination) SetPageCount(v int32) *HotelPoiQueryResponseDataPagination {
	s.PageCount = &v
	return s
}

func (s *HotelPoiQueryResponseDataPagination) SetPageIndex(v int32) *HotelPoiQueryResponseDataPagination {
	s.PageIndex = &v
	return s
}

func (s *HotelPoiQueryResponseDataPagination) SetPageSize(v int32) *HotelPoiQueryResponseDataPagination {
	s.PageSize = &v
	return s
}

func (s *HotelPoiQueryResponseDataPagination) SetTotalCount(v int32) *HotelPoiQueryResponseDataPagination {
	s.TotalCount = &v
	return s
}

type HotelPoiQueryResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s HotelPoiQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s HotelPoiQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *HotelPoiQueryResponseExtra) SetErrorCode(v int32) *HotelPoiQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *HotelPoiQueryResponseExtra) SetLogid(v string) *HotelPoiQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *HotelPoiQueryResponseExtra) SetNow(v int64) *HotelPoiQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *HotelPoiQueryResponseExtra) SetSubDescription(v string) *HotelPoiQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *HotelPoiQueryResponseExtra) SetSubErrorCode(v int32) *HotelPoiQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *HotelPoiQueryResponseExtra) SetDescription(v string) *HotelPoiQueryResponseExtra {
	s.Description = &v
	return s
}

type HotelSavepresaleRequest struct {
	PresaleInfo *HotelSavepresaleRequestPresaleInfo `json:"presale_info,omitempty" xml:"presale_info,omitempty" require:"true"`
	AccountId   *string                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s HotelSavepresaleRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequest) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequest) SetPresaleInfo(v *HotelSavepresaleRequestPresaleInfo) *HotelSavepresaleRequest {
	s.PresaleInfo = v
	return s
}

func (s *HotelSavepresaleRequest) SetAccountId(v string) *HotelSavepresaleRequest {
	s.AccountId = &v
	return s
}

func (s *HotelSavepresaleRequest) SetHeader(v map[string]*string) *HotelSavepresaleRequest {
	s.Header = v
	return s
}

func (s *HotelSavepresaleRequest) SetAccessToken(v string) *HotelSavepresaleRequest {
	s.AccessToken = &v
	return s
}

type HotelSavepresaleRequestPresaleInfo struct {
	OutId             *string                                              `json:"out_id,omitempty" xml:"out_id,omitempty" require:"true"`
	TradeInfo         *HotelSavepresaleRequestPresaleInfoTradeInfo         `json:"trade_info,omitempty" xml:"trade_info,omitempty" require:"true"`
	SaleInfo          *HotelSavepresaleRequestPresaleInfoSaleInfo          `json:"sale_info,omitempty" xml:"sale_info,omitempty" require:"true"`
	SettleType        *int                                                 `json:"settle_type,omitempty" xml:"settle_type,omitempty" require:"true"`
	PreSaleCouponId   *string                                              `json:"pre_sale_coupon_id,omitempty" xml:"pre_sale_coupon_id,omitempty"`
	MemberRequired    *int32                                               `json:"member_required,omitempty" xml:"member_required,omitempty"`
	CategoryId        *string                                              `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	Currency          *string                                              `json:"currency,omitempty" xml:"currency,omitempty"`
	Meals             []*HotelSavepresaleRequestPresaleInfoMealsItem       `json:"meals,omitempty" xml:"meals,omitempty" type:"Repeated"`
	PreSaleCouponInfo *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo `json:"pre_sale_coupon_info,omitempty" xml:"pre_sale_coupon_info,omitempty" require:"true"`
	NoteInfo          *HotelSavepresaleRequestPresaleInfoNoteInfo          `json:"note_info,omitempty" xml:"note_info,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfo) SetOutId(v string) *HotelSavepresaleRequestPresaleInfo {
	s.OutId = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetTradeInfo(v *HotelSavepresaleRequestPresaleInfoTradeInfo) *HotelSavepresaleRequestPresaleInfo {
	s.TradeInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetSaleInfo(v *HotelSavepresaleRequestPresaleInfoSaleInfo) *HotelSavepresaleRequestPresaleInfo {
	s.SaleInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetSettleType(v int) *HotelSavepresaleRequestPresaleInfo {
	s.SettleType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetPreSaleCouponId(v string) *HotelSavepresaleRequestPresaleInfo {
	s.PreSaleCouponId = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetMemberRequired(v int32) *HotelSavepresaleRequestPresaleInfo {
	s.MemberRequired = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetCategoryId(v string) *HotelSavepresaleRequestPresaleInfo {
	s.CategoryId = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetCurrency(v string) *HotelSavepresaleRequestPresaleInfo {
	s.Currency = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetMeals(v []*HotelSavepresaleRequestPresaleInfoMealsItem) *HotelSavepresaleRequestPresaleInfo {
	s.Meals = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetPreSaleCouponInfo(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) *HotelSavepresaleRequestPresaleInfo {
	s.PreSaleCouponInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfo) SetNoteInfo(v *HotelSavepresaleRequestPresaleInfoNoteInfo) *HotelSavepresaleRequestPresaleInfo {
	s.NoteInfo = v
	return s
}

type HotelSavepresaleRequestPresaleInfoMealsItem struct {
	Type *int   `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Num  *int64 `json:"num,omitempty" xml:"num,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoMealsItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoMealsItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoMealsItem) SetType(v int) *HotelSavepresaleRequestPresaleInfoMealsItem {
	s.Type = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoMealsItem) SetNum(v int64) *HotelSavepresaleRequestPresaleInfoMealsItem {
	s.Num = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoNoteInfo struct {
	OtherRemarkInfo       []*string                                                     `json:"other_remark_info,omitempty" xml:"other_remark_info,omitempty" type:"Repeated"`
	ServiceForForeign     *bool                                                         `json:"service_for_foreign,omitempty" xml:"service_for_foreign,omitempty" require:"true"`
	SuperimposedDiscounts *bool                                                         `json:"superimposed_discounts,omitempty" xml:"superimposed_discounts,omitempty" require:"true"`
	CheckTimeRange        *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange     `json:"check_time_range,omitempty" xml:"check_time_range,omitempty"`
	ExclusiveFee          []*HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem `json:"exclusive_fee,omitempty" xml:"exclusive_fee,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfo) SetOtherRemarkInfo(v []*string) *HotelSavepresaleRequestPresaleInfoNoteInfo {
	s.OtherRemarkInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfo) SetServiceForForeign(v bool) *HotelSavepresaleRequestPresaleInfoNoteInfo {
	s.ServiceForForeign = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfo) SetSuperimposedDiscounts(v bool) *HotelSavepresaleRequestPresaleInfoNoteInfo {
	s.SuperimposedDiscounts = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfo) SetCheckTimeRange(v *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange) *HotelSavepresaleRequestPresaleInfoNoteInfo {
	s.CheckTimeRange = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfo) SetExclusiveFee(v []*HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) *HotelSavepresaleRequestPresaleInfoNoteInfo {
	s.ExclusiveFee = v
	return s
}

type HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange struct {
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
	From *string `json:"from,omitempty" xml:"from,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange) SetTo(v string) *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange {
	s.To = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoNoteInfoCheckTimeRange {
	s.From = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem struct {
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Count    *int64  `json:"count,omitempty" xml:"count,omitempty"`
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) SetUnit(v string) *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem {
	s.Unit = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) SetCount(v int64) *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem {
	s.Count = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) SetImageUrl(v string) *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem {
	s.ImageUrl = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) SetName(v string) *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem {
	s.Name = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem) SetPrice(v int64) *HotelSavepresaleRequestPresaleInfoNoteInfoExclusiveFeeItem {
	s.Price = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo struct {
	ImangeList                []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem          `json:"imange_list,omitempty" xml:"imange_list,omitempty" require:"true" type:"Repeated"`
	CouponName                *string                                                                       `json:"coupon_name,omitempty" xml:"coupon_name,omitempty" require:"true"`
	Commodity                 []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem           `json:"commodity,omitempty" xml:"commodity,omitempty" type:"Repeated"`
	UsageDuration             *int64                                                                        `json:"usage_duration,omitempty" xml:"usage_duration,omitempty"`
	BindRatePlans             []*string                                                                     `json:"bind_rate_plans,omitempty" xml:"bind_rate_plans,omitempty" require:"true" type:"Repeated"`
	AppointmentAward          *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward          `json:"appointment_award,omitempty" xml:"appointment_award,omitempty"`
	MarkupType                *int64                                                                        `json:"markup_type,omitempty" xml:"markup_type,omitempty"`
	ApplyRoomNumber           *int64                                                                        `json:"apply_room_number,omitempty" xml:"apply_room_number,omitempty" require:"true"`
	CouponSeparate            *bool                                                                         `json:"coupon_separate,omitempty" xml:"coupon_separate,omitempty"`
	HotelCustomerReservedInfo *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo `json:"hotel_customer_reserved_info,omitempty" xml:"hotel_customer_reserved_info,omitempty"`
	SalesType                 *int                                                                          `json:"sales_type,omitempty" xml:"sales_type,omitempty"`
	MarkupInfo                []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem          `json:"markup_info,omitempty" xml:"markup_info,omitempty" type:"Repeated"`
	OriginalAmount            *int64                                                                        `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	ApplyNights               *int64                                                                        `json:"apply_nights,omitempty" xml:"apply_nights,omitempty" require:"true"`
	IsTestData                *bool                                                                         `json:"is_test_data,omitempty" xml:"is_test_data,omitempty"`
	ActualAmount              *int64                                                                        `json:"actual_amount,omitempty" xml:"actual_amount,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetImangeList(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.ImangeList = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetCouponName(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.CouponName = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetCommodity(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.Commodity = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetUsageDuration(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.UsageDuration = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetBindRatePlans(v []*string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.BindRatePlans = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetAppointmentAward(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.AppointmentAward = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetMarkupType(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.MarkupType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetApplyRoomNumber(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.ApplyRoomNumber = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetCouponSeparate(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.CouponSeparate = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetHotelCustomerReservedInfo(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.HotelCustomerReservedInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetSalesType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.SalesType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetMarkupInfo(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.MarkupInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetOriginalAmount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.OriginalAmount = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetApplyNights(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.ApplyNights = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetIsTestData(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.IsTestData = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo) SetActualAmount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfo {
	s.ActualAmount = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward struct {
	AwardType      *int                                                                          `json:"award_type,omitempty" xml:"award_type,omitempty"`
	Content        *string                                                                       `json:"content,omitempty" xml:"content,omitempty"`
	DateRange      *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange `json:"date_range,omitempty" xml:"date_range,omitempty"`
	IsAward        *bool                                                                         `json:"is_award,omitempty" xml:"is_award,omitempty"`
	IsAwardDisplay *bool                                                                         `json:"is_award_display,omitempty" xml:"is_award_display,omitempty"`
	TimeType       *int                                                                          `json:"time_type,omitempty" xml:"time_type,omitempty"`
	AfterPayDays   *int32                                                                        `json:"after_pay_days,omitempty" xml:"after_pay_days,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetAwardType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.AwardType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetContent(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.Content = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetDateRange(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.DateRange = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetIsAward(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.IsAward = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetIsAwardDisplay(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.IsAwardDisplay = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetTimeType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.TimeType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward) SetAfterPayDays(v int32) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAward {
	s.AfterPayDays = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange) SetEnd(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange {
	s.End = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange) SetStart(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoAppointmentAwardDateRange {
	s.Start = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem struct {
	GroupName     *string                                                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemInfo      []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem    `json:"item_info,omitempty" xml:"item_info,omitempty" type:"Repeated"`
	OptionalCount *int64                                                                             `json:"optional_count,omitempty" xml:"optional_count,omitempty"`
	TotalCount    *int64                                                                             `json:"total_count,omitempty" xml:"total_count,omitempty"`
	CommodityType *int                                                                               `json:"commodity_type,omitempty" xml:"commodity_type,omitempty"`
	EffectScope   []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem `json:"effect_scope,omitempty" xml:"effect_scope,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetGroupName(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.GroupName = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetItemInfo(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.ItemInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetOptionalCount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.OptionalCount = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetTotalCount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.TotalCount = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetCommodityType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.CommodityType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem) SetEffectScope(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItem {
	s.EffectScope = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem struct {
	EffectType   *int      `json:"effect_type,omitempty" xml:"effect_type,omitempty" require:"true"`
	EffectIdList []*string `json:"effect_id_list,omitempty" xml:"effect_id_list,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem) SetEffectType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem {
	s.EffectType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem) SetEffectIdList(v []*string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemEffectScopeItem {
	s.EffectIdList = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem struct {
	ProductInfo *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo `json:"product_info,omitempty" xml:"product_info,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem) SetProductInfo(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItem {
	s.ProductInfo = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo struct {
	Count    *int64  `json:"count,omitempty" xml:"count,omitempty"`
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) SetCount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo {
	s.Count = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) SetImageUrl(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo {
	s.ImageUrl = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) SetName(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo {
	s.Name = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) SetPrice(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo {
	s.Price = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo) SetUnit(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoCommodityItemItemInfoItemProductInfo {
	s.Unit = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo struct {
	ByMerchantCustomConfig     *bool                                                                                                   `json:"by_merchant_custom_config,omitempty" xml:"by_merchant_custom_config,omitempty"`
	RequireBookingInfo         *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo         `json:"require_booking_info,omitempty" xml:"require_booking_info,omitempty"`
	RequirePurchaseContactInfo *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo `json:"require_purchase_contact_info,omitempty" xml:"require_purchase_contact_info,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) SetByMerchantCustomConfig(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo {
	s.ByMerchantCustomConfig = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) SetRequireBookingInfo(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo {
	s.RequireBookingInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo) SetRequirePurchaseContactInfo(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfo {
	s.RequirePurchaseContactInfo = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo struct {
	IdentityTypeWithContentList []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem `json:"identity_type_with_content_list,omitempty" xml:"identity_type_with_content_list,omitempty" type:"Repeated"`
	Enable                      *bool                                                                                                                            `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo) SetIdentityTypeWithContentList(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo {
	s.IdentityTypeWithContentList = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo) SetEnable(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfo {
	s.Enable = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem struct {
	ContactIdentityContentList        []*int        `json:"contact_identity_content_list,omitempty" xml:"contact_identity_content_list,omitempty" type:"Repeated"`
	ContactIdentityTypeEnum           *int          `json:"contact_identity_type_enum,omitempty" xml:"contact_identity_type_enum,omitempty"`
	NeedConsistencyPurchaseAndBooking map[int]*bool `json:"need_consistency_purchase_and_booking,omitempty" xml:"need_consistency_purchase_and_booking,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) SetContactIdentityContentList(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem {
	s.ContactIdentityContentList = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) SetContactIdentityTypeEnum(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem {
	s.ContactIdentityTypeEnum = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem) SetNeedConsistencyPurchaseAndBooking(v map[int]*bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequireBookingInfoIdentityTypeWithContentListItem {
	s.NeedConsistencyPurchaseAndBooking = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo struct {
	Enable                      *bool                                                                                                                                    `json:"enable,omitempty" xml:"enable,omitempty"`
	IdentityTypeWithContentList []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem `json:"identity_type_with_content_list,omitempty" xml:"identity_type_with_content_list,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo) SetEnable(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo {
	s.Enable = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo) SetIdentityTypeWithContentList(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfo {
	s.IdentityTypeWithContentList = v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem struct {
	NeedConsistencyPurchaseAndBooking map[int]*bool `json:"need_consistency_purchase_and_booking,omitempty" xml:"need_consistency_purchase_and_booking,omitempty"`
	ContactIdentityContentList        []*int        `json:"contact_identity_content_list,omitempty" xml:"contact_identity_content_list,omitempty" type:"Repeated"`
	ContactIdentityTypeEnum           *int          `json:"contact_identity_type_enum,omitempty" xml:"contact_identity_type_enum,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) SetNeedConsistencyPurchaseAndBooking(v map[int]*bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem {
	s.NeedConsistencyPurchaseAndBooking = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) SetContactIdentityContentList(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem {
	s.ContactIdentityContentList = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem) SetContactIdentityTypeEnum(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoHotelCustomerReservedInfoRequirePurchaseContactInfoIdentityTypeWithContentListItem {
	s.ContactIdentityTypeEnum = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem struct {
	ImageUrl  *string `json:"image_url,omitempty" xml:"image_url,omitempty" require:"true"`
	ImageType *int    `json:"image_type,omitempty" xml:"image_type,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem) SetImageUrl(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem {
	s.ImageUrl = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem) SetImageType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoImangeListItem {
	s.ImageType = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem struct {
	MarkupDateType    *int                                                                                      `json:"markup_date_type,omitempty" xml:"markup_date_type,omitempty" require:"true"`
	MarkupDays        *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays              `json:"markup_days,omitempty" xml:"markup_days,omitempty"`
	MarkupDaysOfWeek  []*int                                                                                    `json:"markup_days_of_week,omitempty" xml:"markup_days_of_week,omitempty" type:"Repeated"`
	MarkupEffectScope []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem `json:"markup_effect_scope,omitempty" xml:"markup_effect_scope,omitempty" type:"Repeated"`
	MarkupHolidays    []*int                                                                                    `json:"markup_holidays,omitempty" xml:"markup_holidays,omitempty" type:"Repeated"`
	MarkupTime        []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem        `json:"markup_time,omitempty" xml:"markup_time,omitempty" type:"Repeated"`
	HolidaysYear      map[int]*string                                                                           `json:"holidays_year,omitempty" xml:"holidays_year,omitempty"`
	MarkupAmount      *int64                                                                                    `json:"markup_amount,omitempty" xml:"markup_amount,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupDateType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupDateType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupDays(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupDays = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupDaysOfWeek(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupDaysOfWeek = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupEffectScope(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupEffectScope = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupHolidays(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupHolidays = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupTime(v []*HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupTime = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetHolidaysYear(v map[int]*string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.HolidaysYear = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem) SetMarkupAmount(v int64) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItem {
	s.MarkupAmount = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays struct {
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
	From *string `json:"from,omitempty" xml:"from,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays) SetTo(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays {
	s.To = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupDays {
	s.From = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem struct {
	IsFullEffect *bool     `json:"is_full_effect,omitempty" xml:"is_full_effect,omitempty"`
	EffectIdList []*string `json:"effect_id_list,omitempty" xml:"effect_id_list,omitempty" type:"Repeated"`
	EffectType   *int      `json:"effect_type,omitempty" xml:"effect_type,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) SetIsFullEffect(v bool) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem {
	s.IsFullEffect = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) SetEffectIdList(v []*string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem {
	s.EffectIdList = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem) SetEffectType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupEffectScopeItem {
	s.EffectType = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem struct {
	MarkupDays       *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays `json:"markup_days,omitempty" xml:"markup_days,omitempty"`
	MarkupDaysOfWeek []*int                                                                                     `json:"markup_days_of_week,omitempty" xml:"markup_days_of_week,omitempty" type:"Repeated"`
	MarkupHolidays   []*int                                                                                     `json:"markup_holidays,omitempty" xml:"markup_holidays,omitempty" type:"Repeated"`
	HolidaysYear     map[int]*string                                                                            `json:"holidays_year,omitempty" xml:"holidays_year,omitempty"`
	MarkupDateType   *int                                                                                       `json:"markup_date_type,omitempty" xml:"markup_date_type,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) SetMarkupDays(v *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem {
	s.MarkupDays = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) SetMarkupDaysOfWeek(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem {
	s.MarkupDaysOfWeek = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) SetMarkupHolidays(v []*int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem {
	s.MarkupHolidays = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) SetHolidaysYear(v map[int]*string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem {
	s.HolidaysYear = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem) SetMarkupDateType(v int) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItem {
	s.MarkupDateType = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays {
	s.From = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays) SetTo(v string) *HotelSavepresaleRequestPresaleInfoPreSaleCouponInfoMarkupInfoItemMarkupTimeItemMarkupDays {
	s.To = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoSaleInfo struct {
	SaleDate        *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate      `json:"sale_date,omitempty" xml:"sale_date,omitempty" require:"true"`
	ShowChannel     *int                                                     `json:"show_channel,omitempty" xml:"show_channel,omitempty" require:"true"`
	BookDate        *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate      `json:"book_date,omitempty" xml:"book_date,omitempty" require:"true"`
	InventoryInfo   *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo `json:"inventory_info,omitempty" xml:"inventory_info,omitempty" require:"true"`
	IsAutoExtension *bool                                                    `json:"is_auto_extension,omitempty" xml:"is_auto_extension,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfo) SetSaleDate(v *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate) *HotelSavepresaleRequestPresaleInfoSaleInfo {
	s.SaleDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfo) SetShowChannel(v int) *HotelSavepresaleRequestPresaleInfoSaleInfo {
	s.ShowChannel = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfo) SetBookDate(v *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate) *HotelSavepresaleRequestPresaleInfoSaleInfo {
	s.BookDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfo) SetInventoryInfo(v *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo) *HotelSavepresaleRequestPresaleInfoSaleInfo {
	s.InventoryInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfo) SetIsAutoExtension(v bool) *HotelSavepresaleRequestPresaleInfoSaleInfo {
	s.IsAutoExtension = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoSaleInfoBookDate struct {
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
	From *string `json:"from,omitempty" xml:"from,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoBookDate) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoBookDate) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate) SetTo(v string) *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate {
	s.To = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoSaleInfoBookDate {
	s.From = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo struct {
	Num     *int64 `json:"num,omitempty" xml:"num,omitempty"`
	IsLimit *bool  `json:"is_limit,omitempty" xml:"is_limit,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo) SetNum(v int64) *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo {
	s.Num = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo) SetIsLimit(v bool) *HotelSavepresaleRequestPresaleInfoSaleInfoInventoryInfo {
	s.IsLimit = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate struct {
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
	From *string `json:"from,omitempty" xml:"from,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate) SetTo(v string) *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate {
	s.To = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoSaleInfoSaleDate {
	s.From = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfo struct {
	CustomerCanNoUseDate    *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate        `json:"customer_can_no_use_date,omitempty" xml:"customer_can_no_use_date,omitempty"`
	CustomerCanUseTime      *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime          `json:"customer_can_use_time,omitempty" xml:"customer_can_use_time,omitempty" require:"true"`
	LimtBuyRule             *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule                 `json:"limt_buy_rule,omitempty" xml:"limt_buy_rule,omitempty" require:"true"`
	BookRule                *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule                    `json:"book_rule,omitempty" xml:"book_rule,omitempty"`
	PartlyReserveRefundRule *int                                                                    `json:"partly_reserve_refund_rule,omitempty" xml:"partly_reserve_refund_rule,omitempty"`
	CancelBookingRule       *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule           `json:"cancel_booking_rule,omitempty" xml:"cancel_booking_rule,omitempty"`
	CustomerCanUseDate      *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate          `json:"customer_can_use_date,omitempty" xml:"customer_can_use_date,omitempty" require:"true"`
	InvoicInfo              *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo                  `json:"invoic_info,omitempty" xml:"invoic_info,omitempty" require:"true"`
	CancelBookingRuleList   []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem `json:"cancel_booking_rule_list,omitempty" xml:"cancel_booking_rule_list,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetCustomerCanNoUseDate(v *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.CustomerCanNoUseDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetCustomerCanUseTime(v *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.CustomerCanUseTime = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetLimtBuyRule(v *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.LimtBuyRule = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetBookRule(v *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.BookRule = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetPartlyReserveRefundRule(v int) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.PartlyReserveRefundRule = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetCancelBookingRule(v *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.CancelBookingRule = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetCustomerCanUseDate(v *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.CustomerCanUseDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetInvoicInfo(v *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.InvoicInfo = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfo) SetCancelBookingRuleList(v []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) *HotelSavepresaleRequestPresaleInfoTradeInfo {
	s.CancelBookingRuleList = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoBookRule struct {
	TimePreBook     *int64 `json:"time_pre_book,omitempty" xml:"time_pre_book,omitempty"`
	TimeUnit        *int   `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	EarliestBookDay *int64 `json:"earliest_book_day,omitempty" xml:"earliest_book_day,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) SetTimePreBook(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule {
	s.TimePreBook = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) SetTimeUnit(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule {
	s.TimeUnit = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule) SetEarliestBookDay(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoBookRule {
	s.EarliestBookDay = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule struct {
	CancelTimeType *int                                                                            `json:"cancel_time_type,omitempty" xml:"cancel_time_type,omitempty"`
	CancelType     *int                                                                            `json:"cancel_type,omitempty" xml:"cancel_type,omitempty" require:"true"`
	CancelOffset   []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem `json:"cancel_offset,omitempty" xml:"cancel_offset,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) SetCancelTimeType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule {
	s.CancelTimeType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) SetCancelType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule {
	s.CancelType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule) SetCancelOffset(v []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRule {
	s.CancelOffset = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem struct {
	TimeOffset *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset `json:"time_offset,omitempty" xml:"time_offset,omitempty"`
	CutType    *int                                                                                    `json:"cut_type,omitempty" xml:"cut_type,omitempty"`
	CutValue   *int64                                                                                  `json:"cut_value,omitempty" xml:"cut_value,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) SetTimeOffset(v *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem {
	s.TimeOffset = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) SetCutType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem {
	s.CutType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem) SetCutValue(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItem {
	s.CutValue = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset struct {
	Second *int32 `json:"second,omitempty" xml:"second,omitempty"`
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) SetSecond(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset {
	s.Second = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) SetDay(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset {
	s.Day = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) SetHour(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset {
	s.Hour = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset) SetMinute(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleCancelOffsetItemTimeOffset {
	s.Minute = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem struct {
	CancelTimeType *int                                                                                    `json:"cancel_time_type,omitempty" xml:"cancel_time_type,omitempty"`
	CancelType     *int                                                                                    `json:"cancel_type,omitempty" xml:"cancel_type,omitempty" require:"true"`
	CancelOffset   []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem `json:"cancel_offset,omitempty" xml:"cancel_offset,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) SetCancelTimeType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem {
	s.CancelTimeType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) SetCancelType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem {
	s.CancelType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem) SetCancelOffset(v []*HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItem {
	s.CancelOffset = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem struct {
	CutValue   *int64                                                                                          `json:"cut_value,omitempty" xml:"cut_value,omitempty"`
	TimeOffset *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset `json:"time_offset,omitempty" xml:"time_offset,omitempty"`
	CutType    *int                                                                                            `json:"cut_type,omitempty" xml:"cut_type,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) SetCutValue(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem {
	s.CutValue = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) SetTimeOffset(v *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem {
	s.TimeOffset = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem) SetCutType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItem {
	s.CutType = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset struct {
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
	Second *int32 `json:"second,omitempty" xml:"second,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) SetDay(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset {
	s.Day = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) SetHour(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset {
	s.Hour = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) SetMinute(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset {
	s.Minute = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset) SetSecond(v int32) *HotelSavepresaleRequestPresaleInfoTradeInfoCancelBookingRuleListItemCancelOffsetItemTimeOffset {
	s.Second = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate struct {
	CannotUseHolidays   []*int          `json:"cannot_use_holidays,omitempty" xml:"cannot_use_holidays,omitempty" type:"Repeated"`
	HolidaysYear        map[int]*string `json:"holidays_year,omitempty" xml:"holidays_year,omitempty"`
	CannotUseDate       []*string       `json:"cannot_use_date,omitempty" xml:"cannot_use_date,omitempty" type:"Repeated"`
	CannotUseDaysOfWeek []*int          `json:"cannot_use_days_of_week,omitempty" xml:"cannot_use_days_of_week,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) SetCannotUseHolidays(v []*int) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate {
	s.CannotUseHolidays = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) SetHolidaysYear(v map[int]*string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate {
	s.HolidaysYear = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) SetCannotUseDate(v []*string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate {
	s.CannotUseDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate) SetCannotUseDaysOfWeek(v []*int) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanNoUseDate {
	s.CannotUseDaysOfWeek = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate struct {
	DayDuration *int64                                                                `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
	UseDate     *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate `json:"use_date,omitempty" xml:"use_date,omitempty"`
	UseDateType *int                                                                  `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) SetDayDuration(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate {
	s.DayDuration = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) SetUseDate(v *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate {
	s.UseDate = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate) SetUseDateType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDate {
	s.UseDateType = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate {
	s.From = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate) SetTo(v string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseDateUseDate {
	s.To = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime struct {
	UseTimeType *int                                                                            `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
	UseTimeList []*HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem `json:"use_time_list,omitempty" xml:"use_time_list,omitempty" type:"Repeated"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime) SetUseTimeType(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime {
	s.UseTimeType = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime) SetUseTimeList(v []*HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTime {
	s.UseTimeList = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem struct {
	IsTimeNextday *bool                                                                                 `json:"is_time_nextday,omitempty" xml:"is_time_nextday,omitempty" require:"true"`
	TimeSpan      *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan `json:"time_span,omitempty" xml:"time_span,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem) SetIsTimeNextday(v bool) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem {
	s.IsTimeNextday = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem) SetTimeSpan(v *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItem {
	s.TimeSpan = v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan) SetFrom(v string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan {
	s.From = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan) SetTo(v string) *HotelSavepresaleRequestPresaleInfoTradeInfoCustomerCanUseTimeUseTimeListItemTimeSpan {
	s.To = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo struct {
	Subject      *string `json:"subject,omitempty" xml:"subject,omitempty"`
	InvoiceTypes []*int  `json:"invoice_types,omitempty" xml:"invoice_types,omitempty" type:"Repeated"`
	Provider     *int    `json:"provider,omitempty" xml:"provider,omitempty" require:"true"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) SetSubject(v string) *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo {
	s.Subject = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) SetInvoiceTypes(v []*int) *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo {
	s.InvoiceTypes = v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo) SetProvider(v int) *HotelSavepresaleRequestPresaleInfoTradeInfoInvoicInfo {
	s.Provider = &v
	return s
}

type HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule struct {
	EachOrderCanUseMax     *int64 `json:"each_order_can_use_max,omitempty" xml:"each_order_can_use_max,omitempty"`
	EachPersonEachOrderMax *int64 `json:"each_person_each_order_max,omitempty" xml:"each_person_each_order_max,omitempty" require:"true"`
	EachPersonMax          *int64 `json:"each_person_max,omitempty" xml:"each_person_max,omitempty"`
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) SetEachOrderCanUseMax(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule {
	s.EachOrderCanUseMax = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) SetEachPersonEachOrderMax(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule {
	s.EachPersonEachOrderMax = &v
	return s
}

func (s *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule) SetEachPersonMax(v int64) *HotelSavepresaleRequestPresaleInfoTradeInfoLimtBuyRule {
	s.EachPersonMax = &v
	return s
}

type HotelSavepresaleResponse struct {
	Data  *HotelSavepresaleResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *HotelSavepresaleResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s HotelSavepresaleResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleResponse) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleResponse) SetData(v *HotelSavepresaleResponseData) *HotelSavepresaleResponse {
	s.Data = v
	return s
}

func (s *HotelSavepresaleResponse) SetExtra(v *HotelSavepresaleResponseExtra) *HotelSavepresaleResponse {
	s.Extra = v
	return s
}

type HotelSavepresaleResponseData struct {
	OutId           *string `json:"out_id,omitempty" xml:"out_id,omitempty" require:"true"`
	PreSaleCouponId *string `json:"pre_sale_coupon_id,omitempty" xml:"pre_sale_coupon_id,omitempty"`
	GwErrorCode     *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription   *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s HotelSavepresaleResponseData) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleResponseData) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleResponseData) SetOutId(v string) *HotelSavepresaleResponseData {
	s.OutId = &v
	return s
}

func (s *HotelSavepresaleResponseData) SetPreSaleCouponId(v string) *HotelSavepresaleResponseData {
	s.PreSaleCouponId = &v
	return s
}

func (s *HotelSavepresaleResponseData) SetGwErrorCode(v int32) *HotelSavepresaleResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *HotelSavepresaleResponseData) SetGwDescription(v string) *HotelSavepresaleResponseData {
	s.GwDescription = &v
	return s
}

type HotelSavepresaleResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s HotelSavepresaleResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s HotelSavepresaleResponseExtra) GoString() string {
	return s.String()
}

func (s *HotelSavepresaleResponseExtra) SetSubDescription(v string) *HotelSavepresaleResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *HotelSavepresaleResponseExtra) SetSubErrorCode(v int32) *HotelSavepresaleResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *HotelSavepresaleResponseExtra) SetDescription(v string) *HotelSavepresaleResponseExtra {
	s.Description = &v
	return s
}

func (s *HotelSavepresaleResponseExtra) SetErrorCode(v int32) *HotelSavepresaleResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *HotelSavepresaleResponseExtra) SetLogid(v string) *HotelSavepresaleResponseExtra {
	s.Logid = &v
	return s
}

func (s *HotelSavepresaleResponseExtra) SetNow(v int64) *HotelSavepresaleResponseExtra {
	s.Now = &v
	return s
}

type HotelUserUpdateRequest struct {
	Stay             *int64             `json:"stay,omitempty" xml:"stay,omitempty"`
	IsNewMember      *int64             `json:"is_new_member,omitempty" xml:"is_new_member,omitempty"`
	PointsValue      *int64             `json:"points_value,omitempty" xml:"points_value,omitempty"`
	Phone            *string            `json:"phone,omitempty" xml:"phone,omitempty"`
	IdNumber         *string            `json:"id_number,omitempty" xml:"id_number,omitempty"`
	EffectEndTime    *int64             `json:"effect_end_time,omitempty" xml:"effect_end_time,omitempty"`
	MemberCardId     *string            `json:"member_card_id,omitempty" xml:"member_card_id,omitempty"`
	AccountId        *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PointsAmountCent *int64             `json:"points_amount_cent,omitempty" xml:"points_amount_cent,omitempty"`
	Email            *string            `json:"email,omitempty" xml:"email,omitempty"`
	RoomNights       *int64             `json:"room_nights,omitempty" xml:"room_nights,omitempty"`
	UserLevel        *int64             `json:"user_level,omitempty" xml:"user_level,omitempty"`
	Noshow           *int32             `json:"noshow,omitempty" xml:"noshow,omitempty"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	EffectStartTime  *int64             `json:"effect_start_time,omitempty" xml:"effect_start_time,omitempty"`
	LevelKeepNights  *int32             `json:"level_keep_nights,omitempty" xml:"level_keep_nights,omitempty"`
	LastNameCn       *string            `json:"last_name_cn,omitempty" xml:"last_name_cn,omitempty"`
	LevelUpNights    *int32             `json:"level_up_nights,omitempty" xml:"level_up_nights,omitempty"`
	Mobile           *string            `json:"mobile,omitempty" xml:"mobile,omitempty"`
	OpenId           *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Name             *string            `json:"name,omitempty" xml:"name,omitempty"`
	FirstNameCn      *string            `json:"first_name_cn,omitempty" xml:"first_name_cn,omitempty"`
	UpdateTime       *int64             `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
}

func (s HotelUserUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s HotelUserUpdateRequest) GoString() string {
	return s.String()
}

func (s *HotelUserUpdateRequest) SetStay(v int64) *HotelUserUpdateRequest {
	s.Stay = &v
	return s
}

func (s *HotelUserUpdateRequest) SetIsNewMember(v int64) *HotelUserUpdateRequest {
	s.IsNewMember = &v
	return s
}

func (s *HotelUserUpdateRequest) SetPointsValue(v int64) *HotelUserUpdateRequest {
	s.PointsValue = &v
	return s
}

func (s *HotelUserUpdateRequest) SetPhone(v string) *HotelUserUpdateRequest {
	s.Phone = &v
	return s
}

func (s *HotelUserUpdateRequest) SetIdNumber(v string) *HotelUserUpdateRequest {
	s.IdNumber = &v
	return s
}

func (s *HotelUserUpdateRequest) SetEffectEndTime(v int64) *HotelUserUpdateRequest {
	s.EffectEndTime = &v
	return s
}

func (s *HotelUserUpdateRequest) SetMemberCardId(v string) *HotelUserUpdateRequest {
	s.MemberCardId = &v
	return s
}

func (s *HotelUserUpdateRequest) SetAccountId(v string) *HotelUserUpdateRequest {
	s.AccountId = &v
	return s
}

func (s *HotelUserUpdateRequest) SetAccessToken(v string) *HotelUserUpdateRequest {
	s.AccessToken = &v
	return s
}

func (s *HotelUserUpdateRequest) SetPointsAmountCent(v int64) *HotelUserUpdateRequest {
	s.PointsAmountCent = &v
	return s
}

func (s *HotelUserUpdateRequest) SetEmail(v string) *HotelUserUpdateRequest {
	s.Email = &v
	return s
}

func (s *HotelUserUpdateRequest) SetRoomNights(v int64) *HotelUserUpdateRequest {
	s.RoomNights = &v
	return s
}

func (s *HotelUserUpdateRequest) SetUserLevel(v int64) *HotelUserUpdateRequest {
	s.UserLevel = &v
	return s
}

func (s *HotelUserUpdateRequest) SetNoshow(v int32) *HotelUserUpdateRequest {
	s.Noshow = &v
	return s
}

func (s *HotelUserUpdateRequest) SetHeader(v map[string]*string) *HotelUserUpdateRequest {
	s.Header = v
	return s
}

func (s *HotelUserUpdateRequest) SetEffectStartTime(v int64) *HotelUserUpdateRequest {
	s.EffectStartTime = &v
	return s
}

func (s *HotelUserUpdateRequest) SetLevelKeepNights(v int32) *HotelUserUpdateRequest {
	s.LevelKeepNights = &v
	return s
}

func (s *HotelUserUpdateRequest) SetLastNameCn(v string) *HotelUserUpdateRequest {
	s.LastNameCn = &v
	return s
}

func (s *HotelUserUpdateRequest) SetLevelUpNights(v int32) *HotelUserUpdateRequest {
	s.LevelUpNights = &v
	return s
}

func (s *HotelUserUpdateRequest) SetMobile(v string) *HotelUserUpdateRequest {
	s.Mobile = &v
	return s
}

func (s *HotelUserUpdateRequest) SetOpenId(v string) *HotelUserUpdateRequest {
	s.OpenId = &v
	return s
}

func (s *HotelUserUpdateRequest) SetName(v string) *HotelUserUpdateRequest {
	s.Name = &v
	return s
}

func (s *HotelUserUpdateRequest) SetFirstNameCn(v string) *HotelUserUpdateRequest {
	s.FirstNameCn = &v
	return s
}

func (s *HotelUserUpdateRequest) SetUpdateTime(v int64) *HotelUserUpdateRequest {
	s.UpdateTime = &v
	return s
}

type HotelUserUpdateResponse struct {
	Extra *HotelUserUpdateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *HotelUserUpdateResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s HotelUserUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s HotelUserUpdateResponse) GoString() string {
	return s.String()
}

func (s *HotelUserUpdateResponse) SetExtra(v *HotelUserUpdateResponseExtra) *HotelUserUpdateResponse {
	s.Extra = v
	return s
}

func (s *HotelUserUpdateResponse) SetData(v *HotelUserUpdateResponseData) *HotelUserUpdateResponse {
	s.Data = v
	return s
}

type HotelUserUpdateResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s HotelUserUpdateResponseData) String() string {
	return tea.Prettify(s)
}

func (s HotelUserUpdateResponseData) GoString() string {
	return s.String()
}

func (s *HotelUserUpdateResponseData) SetGwErrorCode(v int32) *HotelUserUpdateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *HotelUserUpdateResponseData) SetGwDescription(v string) *HotelUserUpdateResponseData {
	s.GwDescription = &v
	return s
}

type HotelUserUpdateResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s HotelUserUpdateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s HotelUserUpdateResponseExtra) GoString() string {
	return s.String()
}

func (s *HotelUserUpdateResponseExtra) SetNow(v int64) *HotelUserUpdateResponseExtra {
	s.Now = &v
	return s
}

func (s *HotelUserUpdateResponseExtra) SetSubDescription(v string) *HotelUserUpdateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *HotelUserUpdateResponseExtra) SetSubErrorCode(v int32) *HotelUserUpdateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *HotelUserUpdateResponseExtra) SetDescription(v string) *HotelUserUpdateResponseExtra {
	s.Description = &v
	return s
}

func (s *HotelUserUpdateResponseExtra) SetErrorCode(v int32) *HotelUserUpdateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *HotelUserUpdateResponseExtra) SetLogid(v string) *HotelUserUpdateResponseExtra {
	s.Logid = &v
	return s
}

type ImDelAppletTemplateRequest struct {
	CardTemplateId *string            `json:"card_template_id,omitempty" xml:"card_template_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ImDelAppletTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ImDelAppletTemplateRequest) GoString() string {
	return s.String()
}

func (s *ImDelAppletTemplateRequest) SetCardTemplateId(v string) *ImDelAppletTemplateRequest {
	s.CardTemplateId = &v
	return s
}

func (s *ImDelAppletTemplateRequest) SetHeader(v map[string]*string) *ImDelAppletTemplateRequest {
	s.Header = v
	return s
}

func (s *ImDelAppletTemplateRequest) SetAccessToken(v string) *ImDelAppletTemplateRequest {
	s.AccessToken = &v
	return s
}

type ImDelAppletTemplateResponse struct {
	Extra *ImDelAppletTemplateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ImDelAppletTemplateResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ImDelAppletTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ImDelAppletTemplateResponse) GoString() string {
	return s.String()
}

func (s *ImDelAppletTemplateResponse) SetExtra(v *ImDelAppletTemplateResponseExtra) *ImDelAppletTemplateResponse {
	s.Extra = v
	return s
}

func (s *ImDelAppletTemplateResponse) SetData(v *ImDelAppletTemplateResponseData) *ImDelAppletTemplateResponse {
	s.Data = v
	return s
}

type ImDelAppletTemplateResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ImDelAppletTemplateResponseData) String() string {
	return tea.Prettify(s)
}

func (s ImDelAppletTemplateResponseData) GoString() string {
	return s.String()
}

func (s *ImDelAppletTemplateResponseData) SetGwErrorCode(v int32) *ImDelAppletTemplateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ImDelAppletTemplateResponseData) SetGwDescription(v string) *ImDelAppletTemplateResponseData {
	s.GwDescription = &v
	return s
}

type ImDelAppletTemplateResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ImDelAppletTemplateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ImDelAppletTemplateResponseExtra) GoString() string {
	return s.String()
}

func (s *ImDelAppletTemplateResponseExtra) SetLogid(v string) *ImDelAppletTemplateResponseExtra {
	s.Logid = &v
	return s
}

func (s *ImDelAppletTemplateResponseExtra) SetNow(v int64) *ImDelAppletTemplateResponseExtra {
	s.Now = &v
	return s
}

func (s *ImDelAppletTemplateResponseExtra) SetSubDescription(v string) *ImDelAppletTemplateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ImDelAppletTemplateResponseExtra) SetSubErrorCode(v int32) *ImDelAppletTemplateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ImDelAppletTemplateResponseExtra) SetDescription(v string) *ImDelAppletTemplateResponseExtra {
	s.Description = &v
	return s
}

func (s *ImDelAppletTemplateResponseExtra) SetErrorCode(v int32) *ImDelAppletTemplateResponseExtra {
	s.ErrorCode = &v
	return s
}

type ImGetAppletTemplateRequest struct {
	Count          *int32             `json:"count,omitempty" xml:"count,omitempty"`
	Cursor         *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Status         *int32             `json:"status,omitempty" xml:"status,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CardTemplateId *string            `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
}

func (s ImGetAppletTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ImGetAppletTemplateRequest) GoString() string {
	return s.String()
}

func (s *ImGetAppletTemplateRequest) SetCount(v int32) *ImGetAppletTemplateRequest {
	s.Count = &v
	return s
}

func (s *ImGetAppletTemplateRequest) SetCursor(v int64) *ImGetAppletTemplateRequest {
	s.Cursor = &v
	return s
}

func (s *ImGetAppletTemplateRequest) SetStatus(v int32) *ImGetAppletTemplateRequest {
	s.Status = &v
	return s
}

func (s *ImGetAppletTemplateRequest) SetHeader(v map[string]*string) *ImGetAppletTemplateRequest {
	s.Header = v
	return s
}

func (s *ImGetAppletTemplateRequest) SetAccessToken(v string) *ImGetAppletTemplateRequest {
	s.AccessToken = &v
	return s
}

func (s *ImGetAppletTemplateRequest) SetCardTemplateId(v string) *ImGetAppletTemplateRequest {
	s.CardTemplateId = &v
	return s
}

type ImGetAppletTemplateResponse struct {
	Extra *ImGetAppletTemplateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ImGetAppletTemplateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ImGetAppletTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ImGetAppletTemplateResponse) GoString() string {
	return s.String()
}

func (s *ImGetAppletTemplateResponse) SetExtra(v *ImGetAppletTemplateResponseExtra) *ImGetAppletTemplateResponse {
	s.Extra = v
	return s
}

func (s *ImGetAppletTemplateResponse) SetData(v *ImGetAppletTemplateResponseData) *ImGetAppletTemplateResponse {
	s.Data = v
	return s
}

type ImGetAppletTemplateResponseData struct {
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor        *int64                                     `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                      `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	List          []*ImGetAppletTemplateResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ImGetAppletTemplateResponseData) String() string {
	return tea.Prettify(s)
}

func (s ImGetAppletTemplateResponseData) GoString() string {
	return s.String()
}

func (s *ImGetAppletTemplateResponseData) SetGwDescription(v string) *ImGetAppletTemplateResponseData {
	s.GwDescription = &v
	return s
}

func (s *ImGetAppletTemplateResponseData) SetCursor(v int64) *ImGetAppletTemplateResponseData {
	s.Cursor = &v
	return s
}

func (s *ImGetAppletTemplateResponseData) SetHasMore(v bool) *ImGetAppletTemplateResponseData {
	s.HasMore = &v
	return s
}

func (s *ImGetAppletTemplateResponseData) SetList(v []*ImGetAppletTemplateResponseDataListItem) *ImGetAppletTemplateResponseData {
	s.List = v
	return s
}

func (s *ImGetAppletTemplateResponseData) SetGwErrorCode(v int32) *ImGetAppletTemplateResponseData {
	s.GwErrorCode = &v
	return s
}

type ImGetAppletTemplateResponseDataListItem struct {
	RejectReasons  *string `json:"reject_reasons,omitempty" xml:"reject_reasons,omitempty"`
	UpdateTime     *int64  `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
	CreateTime     *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	Title          *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty" require:"true"`
	MediaId        *string `json:"media_id,omitempty" xml:"media_id,omitempty" require:"true"`
	AppIconUrl     *string `json:"app_icon_url,omitempty" xml:"app_icon_url,omitempty"`
	IconMediaId    *string `json:"icon_media_id,omitempty" xml:"icon_media_id,omitempty"`
	CardType       *int32  `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
}

func (s ImGetAppletTemplateResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s ImGetAppletTemplateResponseDataListItem) GoString() string {
	return s.String()
}

func (s *ImGetAppletTemplateResponseDataListItem) SetRejectReasons(v string) *ImGetAppletTemplateResponseDataListItem {
	s.RejectReasons = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetUpdateTime(v int64) *ImGetAppletTemplateResponseDataListItem {
	s.UpdateTime = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetCreateTime(v int64) *ImGetAppletTemplateResponseDataListItem {
	s.CreateTime = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetContent(v string) *ImGetAppletTemplateResponseDataListItem {
	s.Content = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetTitle(v string) *ImGetAppletTemplateResponseDataListItem {
	s.Title = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetName(v string) *ImGetAppletTemplateResponseDataListItem {
	s.Name = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetStatus(v int32) *ImGetAppletTemplateResponseDataListItem {
	s.Status = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetAppId(v string) *ImGetAppletTemplateResponseDataListItem {
	s.AppId = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetCardTemplateId(v string) *ImGetAppletTemplateResponseDataListItem {
	s.CardTemplateId = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetMediaId(v string) *ImGetAppletTemplateResponseDataListItem {
	s.MediaId = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetAppIconUrl(v string) *ImGetAppletTemplateResponseDataListItem {
	s.AppIconUrl = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetIconMediaId(v string) *ImGetAppletTemplateResponseDataListItem {
	s.IconMediaId = &v
	return s
}

func (s *ImGetAppletTemplateResponseDataListItem) SetCardType(v int32) *ImGetAppletTemplateResponseDataListItem {
	s.CardType = &v
	return s
}

type ImGetAppletTemplateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ImGetAppletTemplateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ImGetAppletTemplateResponseExtra) GoString() string {
	return s.String()
}

func (s *ImGetAppletTemplateResponseExtra) SetDescription(v string) *ImGetAppletTemplateResponseExtra {
	s.Description = &v
	return s
}

func (s *ImGetAppletTemplateResponseExtra) SetErrorCode(v int32) *ImGetAppletTemplateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ImGetAppletTemplateResponseExtra) SetLogid(v string) *ImGetAppletTemplateResponseExtra {
	s.Logid = &v
	return s
}

func (s *ImGetAppletTemplateResponseExtra) SetNow(v int64) *ImGetAppletTemplateResponseExtra {
	s.Now = &v
	return s
}

func (s *ImGetAppletTemplateResponseExtra) SetSubDescription(v string) *ImGetAppletTemplateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ImGetAppletTemplateResponseExtra) SetSubErrorCode(v int32) *ImGetAppletTemplateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ImSetAppletTemplateRequest struct {
	MediaId        *string            `json:"media_id,omitempty" xml:"media_id,omitempty"`
	Title          *string            `json:"title,omitempty" xml:"title,omitempty"`
	AppId          *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	CardTemplateId *string            `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	CardType       *int32             `json:"card_type,omitempty" xml:"card_type,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Content        *string            `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ImSetAppletTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ImSetAppletTemplateRequest) GoString() string {
	return s.String()
}

func (s *ImSetAppletTemplateRequest) SetMediaId(v string) *ImSetAppletTemplateRequest {
	s.MediaId = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetTitle(v string) *ImSetAppletTemplateRequest {
	s.Title = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetAppId(v string) *ImSetAppletTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetCardTemplateId(v string) *ImSetAppletTemplateRequest {
	s.CardTemplateId = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetCardType(v int32) *ImSetAppletTemplateRequest {
	s.CardType = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetHeader(v map[string]*string) *ImSetAppletTemplateRequest {
	s.Header = v
	return s
}

func (s *ImSetAppletTemplateRequest) SetAccessToken(v string) *ImSetAppletTemplateRequest {
	s.AccessToken = &v
	return s
}

func (s *ImSetAppletTemplateRequest) SetContent(v string) *ImSetAppletTemplateRequest {
	s.Content = &v
	return s
}

type ImSetAppletTemplateResponse struct {
	Data  *ImSetAppletTemplateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ImSetAppletTemplateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ImSetAppletTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ImSetAppletTemplateResponse) GoString() string {
	return s.String()
}

func (s *ImSetAppletTemplateResponse) SetData(v *ImSetAppletTemplateResponseData) *ImSetAppletTemplateResponse {
	s.Data = v
	return s
}

func (s *ImSetAppletTemplateResponse) SetExtra(v *ImSetAppletTemplateResponseExtra) *ImSetAppletTemplateResponse {
	s.Extra = v
	return s
}

type ImSetAppletTemplateResponseData struct {
	GwDescription  *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	GwErrorCode    *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ImSetAppletTemplateResponseData) String() string {
	return tea.Prettify(s)
}

func (s ImSetAppletTemplateResponseData) GoString() string {
	return s.String()
}

func (s *ImSetAppletTemplateResponseData) SetGwDescription(v string) *ImSetAppletTemplateResponseData {
	s.GwDescription = &v
	return s
}

func (s *ImSetAppletTemplateResponseData) SetCardTemplateId(v string) *ImSetAppletTemplateResponseData {
	s.CardTemplateId = &v
	return s
}

func (s *ImSetAppletTemplateResponseData) SetGwErrorCode(v int32) *ImSetAppletTemplateResponseData {
	s.GwErrorCode = &v
	return s
}

type ImSetAppletTemplateResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ImSetAppletTemplateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ImSetAppletTemplateResponseExtra) GoString() string {
	return s.String()
}

func (s *ImSetAppletTemplateResponseExtra) SetErrorCode(v int32) *ImSetAppletTemplateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ImSetAppletTemplateResponseExtra) SetLogid(v string) *ImSetAppletTemplateResponseExtra {
	s.Logid = &v
	return s
}

func (s *ImSetAppletTemplateResponseExtra) SetNow(v int64) *ImSetAppletTemplateResponseExtra {
	s.Now = &v
	return s
}

func (s *ImSetAppletTemplateResponseExtra) SetSubDescription(v string) *ImSetAppletTemplateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ImSetAppletTemplateResponseExtra) SetSubErrorCode(v int32) *ImSetAppletTemplateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ImSetAppletTemplateResponseExtra) SetDescription(v string) *ImSetAppletTemplateResponseExtra {
	s.Description = &v
	return s
}

type ImageMaterialUploadRequest struct {
	ImageMaterialUrl *string            `json:"image_material_url,omitempty" xml:"image_material_url,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ImageMaterialUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageMaterialUploadRequest) GoString() string {
	return s.String()
}

func (s *ImageMaterialUploadRequest) SetImageMaterialUrl(v string) *ImageMaterialUploadRequest {
	s.ImageMaterialUrl = &v
	return s
}

func (s *ImageMaterialUploadRequest) SetHeader(v map[string]*string) *ImageMaterialUploadRequest {
	s.Header = v
	return s
}

func (s *ImageMaterialUploadRequest) SetAccessToken(v string) *ImageMaterialUploadRequest {
	s.AccessToken = &v
	return s
}

type ImageMaterialUploadResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ImageMaterialUploadResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s ImageMaterialUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageMaterialUploadResponse) GoString() string {
	return s.String()
}

func (s *ImageMaterialUploadResponse) SetLogId(v string) *ImageMaterialUploadResponse {
	s.LogId = &v
	return s
}

func (s *ImageMaterialUploadResponse) SetData(v *ImageMaterialUploadResponseData) *ImageMaterialUploadResponse {
	s.Data = v
	return s
}

func (s *ImageMaterialUploadResponse) SetErrNo(v int32) *ImageMaterialUploadResponse {
	s.ErrNo = &v
	return s
}

func (s *ImageMaterialUploadResponse) SetErrMsg(v string) *ImageMaterialUploadResponse {
	s.ErrMsg = &v
	return s
}

type ImageMaterialUploadResponseData struct {
	ImageMaterialId *string `json:"image_material_id,omitempty" xml:"image_material_id,omitempty"`
}

func (s ImageMaterialUploadResponseData) String() string {
	return tea.Prettify(s)
}

func (s ImageMaterialUploadResponseData) GoString() string {
	return s.String()
}

func (s *ImageMaterialUploadResponseData) SetImageMaterialId(v string) *ImageMaterialUploadResponseData {
	s.ImageMaterialId = &v
	return s
}

type ImagexClientUploadRequest struct {
	Image       *util.FileField    `json:"image,omitempty" xml:"image,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ImagexClientUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s ImagexClientUploadRequest) GoString() string {
	return s.String()
}

func (s *ImagexClientUploadRequest) SetImage(v *util.FileField) *ImagexClientUploadRequest {
	s.Image = v
	return s
}

func (s *ImagexClientUploadRequest) SetHeader(v map[string]*string) *ImagexClientUploadRequest {
	s.Header = v
	return s
}

func (s *ImagexClientUploadRequest) SetAccessToken(v string) *ImagexClientUploadRequest {
	s.AccessToken = &v
	return s
}

type ImagexClientUploadResponse struct {
	Md5     *string                          `json:"md5,omitempty" xml:"md5,omitempty"`
	Data    *ImagexClientUploadResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Width   *int32                           `json:"width,omitempty" xml:"width,omitempty" require:"true"`
	Extra   *ImagexClientUploadResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Height  *int32                           `json:"height,omitempty" xml:"height,omitempty" require:"true"`
	ImageId *string                          `json:"image_id,omitempty" xml:"image_id,omitempty"`
}

func (s ImagexClientUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s ImagexClientUploadResponse) GoString() string {
	return s.String()
}

func (s *ImagexClientUploadResponse) SetMd5(v string) *ImagexClientUploadResponse {
	s.Md5 = &v
	return s
}

func (s *ImagexClientUploadResponse) SetData(v *ImagexClientUploadResponseData) *ImagexClientUploadResponse {
	s.Data = v
	return s
}

func (s *ImagexClientUploadResponse) SetWidth(v int32) *ImagexClientUploadResponse {
	s.Width = &v
	return s
}

func (s *ImagexClientUploadResponse) SetExtra(v *ImagexClientUploadResponseExtra) *ImagexClientUploadResponse {
	s.Extra = v
	return s
}

func (s *ImagexClientUploadResponse) SetHeight(v int32) *ImagexClientUploadResponse {
	s.Height = &v
	return s
}

func (s *ImagexClientUploadResponse) SetImageId(v string) *ImagexClientUploadResponse {
	s.ImageId = &v
	return s
}

type ImagexClientUploadResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ImagexClientUploadResponseData) String() string {
	return tea.Prettify(s)
}

func (s ImagexClientUploadResponseData) GoString() string {
	return s.String()
}

func (s *ImagexClientUploadResponseData) SetGwErrorCode(v int32) *ImagexClientUploadResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ImagexClientUploadResponseData) SetGwDescription(v string) *ImagexClientUploadResponseData {
	s.GwDescription = &v
	return s
}

type ImagexClientUploadResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ImagexClientUploadResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ImagexClientUploadResponseExtra) GoString() string {
	return s.String()
}

func (s *ImagexClientUploadResponseExtra) SetLogid(v string) *ImagexClientUploadResponseExtra {
	s.Logid = &v
	return s
}

func (s *ImagexClientUploadResponseExtra) SetNow(v int64) *ImagexClientUploadResponseExtra {
	s.Now = &v
	return s
}

func (s *ImagexClientUploadResponseExtra) SetSubDescription(v string) *ImagexClientUploadResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ImagexClientUploadResponseExtra) SetSubErrorCode(v int32) *ImagexClientUploadResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ImagexClientUploadResponseExtra) SetDescription(v string) *ImagexClientUploadResponseExtra {
	s.Description = &v
	return s
}

func (s *ImagexClientUploadResponseExtra) SetErrorCode(v int32) *ImagexClientUploadResponseExtra {
	s.ErrorCode = &v
	return s
}

type ImportCardAddRequest struct {
	TemplateCode *int32             `json:"template_code,omitempty" xml:"template_code,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ItemMap      map[string]*string `json:"item_map,omitempty" xml:"item_map,omitempty" require:"true"`
	StartPage    *string            `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
	AnchoridList []*int64           `json:"anchorid_list,omitempty" xml:"anchorid_list,omitempty" require:"true" type:"Repeated"`
}

func (s ImportCardAddRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCardAddRequest) GoString() string {
	return s.String()
}

func (s *ImportCardAddRequest) SetTemplateCode(v int32) *ImportCardAddRequest {
	s.TemplateCode = &v
	return s
}

func (s *ImportCardAddRequest) SetHeader(v map[string]*string) *ImportCardAddRequest {
	s.Header = v
	return s
}

func (s *ImportCardAddRequest) SetAccessToken(v string) *ImportCardAddRequest {
	s.AccessToken = &v
	return s
}

func (s *ImportCardAddRequest) SetItemMap(v map[string]*string) *ImportCardAddRequest {
	s.ItemMap = v
	return s
}

func (s *ImportCardAddRequest) SetStartPage(v string) *ImportCardAddRequest {
	s.StartPage = &v
	return s
}

func (s *ImportCardAddRequest) SetAnchoridList(v []*int64) *ImportCardAddRequest {
	s.AnchoridList = v
	return s
}

type ImportCardAddResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ImportCardAddResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportCardAddResponse) GoString() string {
	return s.String()
}

func (s *ImportCardAddResponse) SetErrNo(v int32) *ImportCardAddResponse {
	s.ErrNo = &v
	return s
}

func (s *ImportCardAddResponse) SetErrMsg(v string) *ImportCardAddResponse {
	s.ErrMsg = &v
	return s
}

func (s *ImportCardAddResponse) SetLogId(v string) *ImportCardAddResponse {
	s.LogId = &v
	return s
}

type ImportCardGetRequest struct {
	Anchorid    *int64             `json:"anchorid,omitempty" xml:"anchorid,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartPage   *string            `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
}

func (s ImportCardGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportCardGetRequest) GoString() string {
	return s.String()
}

func (s *ImportCardGetRequest) SetAnchorid(v int64) *ImportCardGetRequest {
	s.Anchorid = &v
	return s
}

func (s *ImportCardGetRequest) SetHeader(v map[string]*string) *ImportCardGetRequest {
	s.Header = v
	return s
}

func (s *ImportCardGetRequest) SetAccessToken(v string) *ImportCardGetRequest {
	s.AccessToken = &v
	return s
}

func (s *ImportCardGetRequest) SetStartPage(v string) *ImportCardGetRequest {
	s.StartPage = &v
	return s
}

type ImportCardGetResponse struct {
	AuditStatus  *int32             `json:"audit_status,omitempty" xml:"audit_status,omitempty" require:"true"`
	RejectReason *string            `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	ErrNo        *int32             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg       *string            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId        *string            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	TemplateCode *int32             `json:"template_code,omitempty" xml:"template_code,omitempty" require:"true"`
	ItemMap      map[string]*string `json:"item_map,omitempty" xml:"item_map,omitempty" require:"true"`
	Title        *string            `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s ImportCardGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportCardGetResponse) GoString() string {
	return s.String()
}

func (s *ImportCardGetResponse) SetAuditStatus(v int32) *ImportCardGetResponse {
	s.AuditStatus = &v
	return s
}

func (s *ImportCardGetResponse) SetRejectReason(v string) *ImportCardGetResponse {
	s.RejectReason = &v
	return s
}

func (s *ImportCardGetResponse) SetErrNo(v int32) *ImportCardGetResponse {
	s.ErrNo = &v
	return s
}

func (s *ImportCardGetResponse) SetErrMsg(v string) *ImportCardGetResponse {
	s.ErrMsg = &v
	return s
}

func (s *ImportCardGetResponse) SetLogId(v string) *ImportCardGetResponse {
	s.LogId = &v
	return s
}

func (s *ImportCardGetResponse) SetTemplateCode(v int32) *ImportCardGetResponse {
	s.TemplateCode = &v
	return s
}

func (s *ImportCardGetResponse) SetItemMap(v map[string]*string) *ImportCardGetResponse {
	s.ItemMap = v
	return s
}

func (s *ImportCardGetResponse) SetTitle(v string) *ImportCardGetResponse {
	s.Title = &v
	return s
}

type InfoMatchRequest struct {
	HotelImages     []*InfoMatchRequestHotelImagesItem     `json:"hotel_images,omitempty" xml:"hotel_images,omitempty" type:"Repeated"`
	HotelPolicy     *InfoMatchRequestHotelPolicy           `json:"hotel_policy,omitempty" xml:"hotel_policy,omitempty"`
	OutHotelId      *string                                `json:"out_hotel_id,omitempty" xml:"out_hotel_id,omitempty" require:"true"`
	HotelId         *string                                `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	AccountId       *string                                `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AccessToken     *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	HotelBaseInfo   *InfoMatchRequestHotelBaseInfo         `json:"hotel_base_info,omitempty" xml:"hotel_base_info,omitempty"`
	HotelFacilities []*InfoMatchRequestHotelFacilitiesItem `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty" type:"Repeated"`
	Active          *bool                                  `json:"active,omitempty" xml:"active,omitempty"`
	Header          map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
}

func (s InfoMatchRequest) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequest) GoString() string {
	return s.String()
}

func (s *InfoMatchRequest) SetHotelImages(v []*InfoMatchRequestHotelImagesItem) *InfoMatchRequest {
	s.HotelImages = v
	return s
}

func (s *InfoMatchRequest) SetHotelPolicy(v *InfoMatchRequestHotelPolicy) *InfoMatchRequest {
	s.HotelPolicy = v
	return s
}

func (s *InfoMatchRequest) SetOutHotelId(v string) *InfoMatchRequest {
	s.OutHotelId = &v
	return s
}

func (s *InfoMatchRequest) SetHotelId(v string) *InfoMatchRequest {
	s.HotelId = &v
	return s
}

func (s *InfoMatchRequest) SetAccountId(v string) *InfoMatchRequest {
	s.AccountId = &v
	return s
}

func (s *InfoMatchRequest) SetAccessToken(v string) *InfoMatchRequest {
	s.AccessToken = &v
	return s
}

func (s *InfoMatchRequest) SetHotelBaseInfo(v *InfoMatchRequestHotelBaseInfo) *InfoMatchRequest {
	s.HotelBaseInfo = v
	return s
}

func (s *InfoMatchRequest) SetHotelFacilities(v []*InfoMatchRequestHotelFacilitiesItem) *InfoMatchRequest {
	s.HotelFacilities = v
	return s
}

func (s *InfoMatchRequest) SetActive(v bool) *InfoMatchRequest {
	s.Active = &v
	return s
}

func (s *InfoMatchRequest) SetHeader(v map[string]*string) *InfoMatchRequest {
	s.Header = v
	return s
}

type InfoMatchRequestHotelBaseInfo struct {
	StarRating           *int                                                     `json:"star_rating,omitempty" xml:"star_rating,omitempty"`
	Currency             *string                                                  `json:"currency,omitempty" xml:"currency,omitempty"`
	HotelEnBrief         *string                                                  `json:"hotel_en_brief,omitempty" xml:"hotel_en_brief,omitempty"`
	HotelBuildingArea    *float64                                                 `json:"hotel_building_area,omitempty" xml:"hotel_building_area,omitempty"`
	Longitude            *float64                                                 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Phone                *InfoMatchRequestHotelBaseInfoPhone                      `json:"phone,omitempty" xml:"phone,omitempty"`
	HotelBrief           *string                                                  `json:"hotel_brief,omitempty" xml:"hotel_brief,omitempty"`
	Address              *InfoMatchRequestHotelBaseInfoAddress                    `json:"address,omitempty" xml:"address,omitempty"`
	TotalRoomQuantity    *int64                                                   `json:"total_room_quantity,omitempty" xml:"total_room_quantity,omitempty"`
	Latitude             *float64                                                 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	HotelName            *string                                                  `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	RecieveForeignGuest  *int                                                     `json:"recieve_foreign_guest,omitempty" xml:"recieve_foreign_guest,omitempty"`
	HotelEnName          *string                                                  `json:"hotel_en_name,omitempty" xml:"hotel_en_name,omitempty"`
	BuildTime            *string                                                  `json:"build_time,omitempty" xml:"build_time,omitempty"`
	LastRenovation       *string                                                  `json:"last_renovation,omitempty" xml:"last_renovation,omitempty"`
	ImportantInformation []*InfoMatchRequestHotelBaseInfoImportantInformationItem `json:"important_information,omitempty" xml:"important_information,omitempty" type:"Repeated"`
}

func (s InfoMatchRequestHotelBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelBaseInfo) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelBaseInfo) SetStarRating(v int) *InfoMatchRequestHotelBaseInfo {
	s.StarRating = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetCurrency(v string) *InfoMatchRequestHotelBaseInfo {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetHotelEnBrief(v string) *InfoMatchRequestHotelBaseInfo {
	s.HotelEnBrief = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetHotelBuildingArea(v float64) *InfoMatchRequestHotelBaseInfo {
	s.HotelBuildingArea = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetLongitude(v float64) *InfoMatchRequestHotelBaseInfo {
	s.Longitude = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetPhone(v *InfoMatchRequestHotelBaseInfoPhone) *InfoMatchRequestHotelBaseInfo {
	s.Phone = v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetHotelBrief(v string) *InfoMatchRequestHotelBaseInfo {
	s.HotelBrief = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetAddress(v *InfoMatchRequestHotelBaseInfoAddress) *InfoMatchRequestHotelBaseInfo {
	s.Address = v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetTotalRoomQuantity(v int64) *InfoMatchRequestHotelBaseInfo {
	s.TotalRoomQuantity = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetLatitude(v float64) *InfoMatchRequestHotelBaseInfo {
	s.Latitude = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetHotelName(v string) *InfoMatchRequestHotelBaseInfo {
	s.HotelName = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetRecieveForeignGuest(v int) *InfoMatchRequestHotelBaseInfo {
	s.RecieveForeignGuest = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetHotelEnName(v string) *InfoMatchRequestHotelBaseInfo {
	s.HotelEnName = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetBuildTime(v string) *InfoMatchRequestHotelBaseInfo {
	s.BuildTime = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetLastRenovation(v string) *InfoMatchRequestHotelBaseInfo {
	s.LastRenovation = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfo) SetImportantInformation(v []*InfoMatchRequestHotelBaseInfoImportantInformationItem) *InfoMatchRequestHotelBaseInfo {
	s.ImportantInformation = v
	return s
}

type InfoMatchRequestHotelBaseInfoAddress struct {
	Province      *string `json:"province,omitempty" xml:"province,omitempty"`
	City          *string `json:"city,omitempty" xml:"city,omitempty"`
	Country       *string `json:"country,omitempty" xml:"country,omitempty"`
	DetailAddress *string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	District      *string `json:"district,omitempty" xml:"district,omitempty"`
}

func (s InfoMatchRequestHotelBaseInfoAddress) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelBaseInfoAddress) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelBaseInfoAddress) SetProvince(v string) *InfoMatchRequestHotelBaseInfoAddress {
	s.Province = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoAddress) SetCity(v string) *InfoMatchRequestHotelBaseInfoAddress {
	s.City = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoAddress) SetCountry(v string) *InfoMatchRequestHotelBaseInfoAddress {
	s.Country = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoAddress) SetDetailAddress(v string) *InfoMatchRequestHotelBaseInfoAddress {
	s.DetailAddress = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoAddress) SetDistrict(v string) *InfoMatchRequestHotelBaseInfoAddress {
	s.District = &v
	return s
}

type InfoMatchRequestHotelBaseInfoImportantInformationItem struct {
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	EffectEndTime   *string `json:"effect_end_time,omitempty" xml:"effect_end_time,omitempty"`
	EffectStartTime *string `json:"effect_start_time,omitempty" xml:"effect_start_time,omitempty"`
}

func (s InfoMatchRequestHotelBaseInfoImportantInformationItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelBaseInfoImportantInformationItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelBaseInfoImportantInformationItem) SetContent(v string) *InfoMatchRequestHotelBaseInfoImportantInformationItem {
	s.Content = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoImportantInformationItem) SetEffectEndTime(v string) *InfoMatchRequestHotelBaseInfoImportantInformationItem {
	s.EffectEndTime = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoImportantInformationItem) SetEffectStartTime(v string) *InfoMatchRequestHotelBaseInfoImportantInformationItem {
	s.EffectStartTime = &v
	return s
}

type InfoMatchRequestHotelBaseInfoPhone struct {
	AreaCode    *string `json:"area_code,omitempty" xml:"area_code,omitempty"`
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	ExtCode     *string `json:"ext_code,omitempty" xml:"ext_code,omitempty"`
	MainCode    *string `json:"main_code,omitempty" xml:"main_code,omitempty"`
	PhoneType   *int    `json:"phone_type,omitempty" xml:"phone_type,omitempty"`
}

func (s InfoMatchRequestHotelBaseInfoPhone) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelBaseInfoPhone) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelBaseInfoPhone) SetAreaCode(v string) *InfoMatchRequestHotelBaseInfoPhone {
	s.AreaCode = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoPhone) SetCountryCode(v string) *InfoMatchRequestHotelBaseInfoPhone {
	s.CountryCode = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoPhone) SetExtCode(v string) *InfoMatchRequestHotelBaseInfoPhone {
	s.ExtCode = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoPhone) SetMainCode(v string) *InfoMatchRequestHotelBaseInfoPhone {
	s.MainCode = &v
	return s
}

func (s *InfoMatchRequestHotelBaseInfoPhone) SetPhoneType(v int) *InfoMatchRequestHotelBaseInfoPhone {
	s.PhoneType = &v
	return s
}

type InfoMatchRequestHotelFacilitiesItem struct {
	FacilityId       *int32 `json:"facility_id,omitempty" xml:"facility_id,omitempty"`
	FacilityCategory *int   `json:"facility_category,omitempty" xml:"facility_category,omitempty"`
}

func (s InfoMatchRequestHotelFacilitiesItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelFacilitiesItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelFacilitiesItem) SetFacilityId(v int32) *InfoMatchRequestHotelFacilitiesItem {
	s.FacilityId = &v
	return s
}

func (s *InfoMatchRequestHotelFacilitiesItem) SetFacilityCategory(v int) *InfoMatchRequestHotelFacilitiesItem {
	s.FacilityCategory = &v
	return s
}

type InfoMatchRequestHotelImagesItem struct {
	ImageUrl  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	ImageType *int    `json:"image_type,omitempty" xml:"image_type,omitempty"`
}

func (s InfoMatchRequestHotelImagesItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelImagesItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelImagesItem) SetImageUrl(v string) *InfoMatchRequestHotelImagesItem {
	s.ImageUrl = &v
	return s
}

func (s *InfoMatchRequestHotelImagesItem) SetImageType(v int) *InfoMatchRequestHotelImagesItem {
	s.ImageType = &v
	return s
}

type InfoMatchRequestHotelPolicy struct {
	ChildrenPolicy         *InfoMatchRequestHotelPolicyChildrenPolicy         `json:"children_policy,omitempty" xml:"children_policy,omitempty"`
	ExtraBedPolicy         *InfoMatchRequestHotelPolicyExtraBedPolicy         `json:"extra_bed_policy,omitempty" xml:"extra_bed_policy,omitempty"`
	ParkingPolicy          *InfoMatchRequestHotelPolicyParkingPolicy          `json:"parking_policy,omitempty" xml:"parking_policy,omitempty"`
	PetPolicy              *InfoMatchRequestHotelPolicyPetPolicy              `json:"pet_policy,omitempty" xml:"pet_policy,omitempty"`
	BreakfastPolicy        *InfoMatchRequestHotelPolicyBreakfastPolicy        `json:"breakfast_policy,omitempty" xml:"breakfast_policy,omitempty"`
	CheckInPolicy          *InfoMatchRequestHotelPolicyCheckInPolicy          `json:"check_in_policy,omitempty" xml:"check_in_policy,omitempty"`
	ChildBreakfastPolicy   *InfoMatchRequestHotelPolicyChildBreakfastPolicy   `json:"child_breakfast_policy,omitempty" xml:"child_breakfast_policy,omitempty"`
	ChildUseExistBedPolicy *InfoMatchRequestHotelPolicyChildUseExistBedPolicy `json:"child_use_exist_bed_policy,omitempty" xml:"child_use_exist_bed_policy,omitempty"`
}

func (s InfoMatchRequestHotelPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicy) SetChildrenPolicy(v *InfoMatchRequestHotelPolicyChildrenPolicy) *InfoMatchRequestHotelPolicy {
	s.ChildrenPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetExtraBedPolicy(v *InfoMatchRequestHotelPolicyExtraBedPolicy) *InfoMatchRequestHotelPolicy {
	s.ExtraBedPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetParkingPolicy(v *InfoMatchRequestHotelPolicyParkingPolicy) *InfoMatchRequestHotelPolicy {
	s.ParkingPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetPetPolicy(v *InfoMatchRequestHotelPolicyPetPolicy) *InfoMatchRequestHotelPolicy {
	s.PetPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetBreakfastPolicy(v *InfoMatchRequestHotelPolicyBreakfastPolicy) *InfoMatchRequestHotelPolicy {
	s.BreakfastPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetCheckInPolicy(v *InfoMatchRequestHotelPolicyCheckInPolicy) *InfoMatchRequestHotelPolicy {
	s.CheckInPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetChildBreakfastPolicy(v *InfoMatchRequestHotelPolicyChildBreakfastPolicy) *InfoMatchRequestHotelPolicy {
	s.ChildBreakfastPolicy = v
	return s
}

func (s *InfoMatchRequestHotelPolicy) SetChildUseExistBedPolicy(v *InfoMatchRequestHotelPolicyChildUseExistBedPolicy) *InfoMatchRequestHotelPolicy {
	s.ChildUseExistBedPolicy = v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicy struct {
	BreakfastForm     []*int                                              `json:"breakfast_form,omitempty" xml:"breakfast_form,omitempty" type:"Repeated"`
	BreakfastType     []*int                                              `json:"breakfast_type,omitempty" xml:"breakfast_type,omitempty" type:"Repeated"`
	OpenTime          *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime `json:"open_time,omitempty" xml:"open_time,omitempty"`
	SupportBreakfast  *bool                                               `json:"support_breakfast,omitempty" xml:"support_breakfast,omitempty"`
	BreakfastCategory []*int                                              `json:"breakfast_category,omitempty" xml:"breakfast_category,omitempty" type:"Repeated"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicy) SetBreakfastForm(v []*int) *InfoMatchRequestHotelPolicyBreakfastPolicy {
	s.BreakfastForm = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicy) SetBreakfastType(v []*int) *InfoMatchRequestHotelPolicyBreakfastPolicy {
	s.BreakfastType = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicy) SetOpenTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) *InfoMatchRequestHotelPolicyBreakfastPolicy {
	s.OpenTime = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicy) SetSupportBreakfast(v bool) *InfoMatchRequestHotelPolicyBreakfastPolicy {
	s.SupportBreakfast = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicy) SetBreakfastCategory(v []*int) *InfoMatchRequestHotelPolicyBreakfastPolicy {
	s.BreakfastCategory = v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime struct {
	WeekDayTimes *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes `json:"week_day_times,omitempty" xml:"week_day_times,omitempty"`
	EveryDayTime *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime `json:"every_day_time,omitempty" xml:"every_day_time,omitempty"`
	OpenTimeType *int                                                            `json:"open_time_type,omitempty" xml:"open_time_type,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) SetWeekDayTimes(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime {
	s.WeekDayTimes = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) SetEveryDayTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime {
	s.EveryDayTime = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime) SetOpenTimeType(v int) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTime {
	s.OpenTimeType = &v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime struct {
	EndTime   *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime   `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime) SetEndTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime {
	s.EndTime = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime) SetStartTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTime {
	s.StartTime = v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime struct {
	Hour        *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Mintue      *int32 `json:"mintue,omitempty" xml:"mintue,omitempty"`
	Second      *int32 `json:"second,omitempty" xml:"second,omitempty"`
	TimeLimited *bool  `json:"time_limited,omitempty" xml:"time_limited,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) SetHour(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime {
	s.Hour = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) SetMintue(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime {
	s.Mintue = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) SetSecond(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime {
	s.Second = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime) SetTimeLimited(v bool) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeEndTime {
	s.TimeLimited = &v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime struct {
	TimeLimited *bool  `json:"time_limited,omitempty" xml:"time_limited,omitempty"`
	Hour        *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Mintue      *int32 `json:"mintue,omitempty" xml:"mintue,omitempty"`
	Second      *int32 `json:"second,omitempty" xml:"second,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) SetTimeLimited(v bool) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime {
	s.TimeLimited = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) SetHour(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime {
	s.Hour = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) SetMintue(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime {
	s.Mintue = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime) SetSecond(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeEveryDayTimeStartTime {
	s.Second = &v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes struct {
	TimePeriods []*InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem `json:"time_periods,omitempty" xml:"time_periods,omitempty" type:"Repeated"`
	WeekDays    []*int                                                                           `json:"week_days,omitempty" xml:"week_days,omitempty" type:"Repeated"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes) SetTimePeriods(v []*InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes {
	s.TimePeriods = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes) SetWeekDays(v []*int) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimes {
	s.WeekDays = v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem struct {
	EndTime   *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime   `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem) SetEndTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem {
	s.EndTime = v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem) SetStartTime(v *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItem {
	s.StartTime = v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime struct {
	TimeLimited *bool  `json:"time_limited,omitempty" xml:"time_limited,omitempty"`
	Hour        *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Mintue      *int32 `json:"mintue,omitempty" xml:"mintue,omitempty"`
	Second      *int32 `json:"second,omitempty" xml:"second,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) SetTimeLimited(v bool) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime {
	s.TimeLimited = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) SetHour(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime {
	s.Hour = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) SetMintue(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime {
	s.Mintue = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime) SetSecond(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemEndTime {
	s.Second = &v
	return s
}

type InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime struct {
	Hour        *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Mintue      *int32 `json:"mintue,omitempty" xml:"mintue,omitempty"`
	Second      *int32 `json:"second,omitempty" xml:"second,omitempty"`
	TimeLimited *bool  `json:"time_limited,omitempty" xml:"time_limited,omitempty"`
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) SetHour(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime {
	s.Hour = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) SetMintue(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime {
	s.Mintue = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) SetSecond(v int32) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime {
	s.Second = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime) SetTimeLimited(v bool) *InfoMatchRequestHotelPolicyBreakfastPolicyOpenTimeWeekDayTimesTimePeriodsItemStartTime {
	s.TimeLimited = &v
	return s
}

type InfoMatchRequestHotelPolicyCheckInPolicy struct {
	CheckInTo      *string `json:"check_in_to,omitempty" xml:"check_in_to,omitempty"`
	CheckOutFrom   *string `json:"check_out_from,omitempty" xml:"check_out_from,omitempty"`
	CheckOutTo     *string `json:"check_out_to,omitempty" xml:"check_out_to,omitempty"`
	IsLimitArrival *bool   `json:"is_limit_arrival,omitempty" xml:"is_limit_arrival,omitempty"`
	CheckInFrom    *string `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
}

func (s InfoMatchRequestHotelPolicyCheckInPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyCheckInPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyCheckInPolicy) SetCheckInTo(v string) *InfoMatchRequestHotelPolicyCheckInPolicy {
	s.CheckInTo = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyCheckInPolicy) SetCheckOutFrom(v string) *InfoMatchRequestHotelPolicyCheckInPolicy {
	s.CheckOutFrom = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyCheckInPolicy) SetCheckOutTo(v string) *InfoMatchRequestHotelPolicyCheckInPolicy {
	s.CheckOutTo = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyCheckInPolicy) SetIsLimitArrival(v bool) *InfoMatchRequestHotelPolicyCheckInPolicy {
	s.IsLimitArrival = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyCheckInPolicy) SetCheckInFrom(v string) *InfoMatchRequestHotelPolicyCheckInPolicy {
	s.CheckInFrom = &v
	return s
}

type InfoMatchRequestHotelPolicyChildBreakfastPolicy struct {
	ChildBreakfastAge    []*InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem    `json:"child_breakfast_age,omitempty" xml:"child_breakfast_age,omitempty" type:"Repeated"`
	ChildBreakfastHeight []*InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem `json:"child_breakfast_height,omitempty" xml:"child_breakfast_height,omitempty" type:"Repeated"`
	AgeOrHeight          *int                                                                       `json:"age_or_height,omitempty" xml:"age_or_height,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicy) SetChildBreakfastAge(v []*InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) *InfoMatchRequestHotelPolicyChildBreakfastPolicy {
	s.ChildBreakfastAge = v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicy) SetChildBreakfastHeight(v []*InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) *InfoMatchRequestHotelPolicyChildBreakfastPolicy {
	s.ChildBreakfastHeight = v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicy) SetAgeOrHeight(v int) *InfoMatchRequestHotelPolicyChildBreakfastPolicy {
	s.AgeOrHeight = &v
	return s
}

type InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem struct {
	Start           *int32  `json:"start,omitempty" xml:"start,omitempty"`
	ChargeFrequency *string `json:"charge_frequency,omitempty" xml:"charge_frequency,omitempty"`
	ChargePercent   *int32  `json:"charge_percent,omitempty" xml:"charge_percent,omitempty"`
	ChargePrice     *int64  `json:"charge_price,omitempty" xml:"charge_price,omitempty"`
	Chargeable      *bool   `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string `json:"currency,omitempty" xml:"currency,omitempty"`
	End             *int32  `json:"end,omitempty" xml:"end,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetStart(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.Start = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetChargeFrequency(v string) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.ChargeFrequency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetChargePercent(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.ChargePercent = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetChargePrice(v int64) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.ChargePrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetChargeable(v bool) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetCurrency(v string) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem) SetEnd(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastAgeItem {
	s.End = &v
	return s
}

type InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem struct {
	Start           *int32  `json:"start,omitempty" xml:"start,omitempty"`
	ChargeFrequency *string `json:"charge_frequency,omitempty" xml:"charge_frequency,omitempty"`
	ChargePercent   *int32  `json:"charge_percent,omitempty" xml:"charge_percent,omitempty"`
	ChargePrice     *int64  `json:"charge_price,omitempty" xml:"charge_price,omitempty"`
	Chargeable      *bool   `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string `json:"currency,omitempty" xml:"currency,omitempty"`
	End             *int32  `json:"end,omitempty" xml:"end,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetStart(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.Start = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetChargeFrequency(v string) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.ChargeFrequency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetChargePercent(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.ChargePercent = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetChargePrice(v int64) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.ChargePrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetChargeable(v bool) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetCurrency(v string) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem) SetEnd(v int32) *InfoMatchRequestHotelPolicyChildBreakfastPolicyChildBreakfastHeightItem {
	s.End = &v
	return s
}

type InfoMatchRequestHotelPolicyChildUseExistBedPolicy struct {
	Charges       []*InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem `json:"charges,omitempty" xml:"charges,omitempty" type:"Repeated"`
	AllowChildNum *int64                                                          `json:"allow_child_num,omitempty" xml:"allow_child_num,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildUseExistBedPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildUseExistBedPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicy) SetCharges(v []*InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) *InfoMatchRequestHotelPolicyChildUseExistBedPolicy {
	s.Charges = v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicy) SetAllowChildNum(v int64) *InfoMatchRequestHotelPolicyChildUseExistBedPolicy {
	s.AllowChildNum = &v
	return s
}

type InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem struct {
	End             *int32  `json:"end,omitempty" xml:"end,omitempty"`
	Start           *int32  `json:"start,omitempty" xml:"start,omitempty"`
	ChargeFrequency *string `json:"charge_frequency,omitempty" xml:"charge_frequency,omitempty"`
	ChargePercent   *int32  `json:"charge_percent,omitempty" xml:"charge_percent,omitempty"`
	ChargePrice     *int64  `json:"charge_price,omitempty" xml:"charge_price,omitempty"`
	Chargeable      *bool   `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string `json:"currency,omitempty" xml:"currency,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetEnd(v int32) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.End = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetStart(v int32) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.Start = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetChargeFrequency(v string) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.ChargeFrequency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetChargePercent(v int32) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.ChargePercent = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetChargePrice(v int64) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.ChargePrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetChargeable(v bool) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem) SetCurrency(v string) *InfoMatchRequestHotelPolicyChildUseExistBedPolicyChargesItem {
	s.Currency = &v
	return s
}

type InfoMatchRequestHotelPolicyChildrenPolicy struct {
	MinChildrenAge    *int32 `json:"min_children_age,omitempty" xml:"min_children_age,omitempty"`
	IsAllowChildren   *bool  `json:"is_allow_children,omitempty" xml:"is_allow_children,omitempty"`
	IsAllowSharingBed *bool  `json:"is_allow_sharing_bed,omitempty" xml:"is_allow_sharing_bed,omitempty"`
	MaxChildrenAge    *int32 `json:"max_children_age,omitempty" xml:"max_children_age,omitempty"`
}

func (s InfoMatchRequestHotelPolicyChildrenPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyChildrenPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyChildrenPolicy) SetMinChildrenAge(v int32) *InfoMatchRequestHotelPolicyChildrenPolicy {
	s.MinChildrenAge = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildrenPolicy) SetIsAllowChildren(v bool) *InfoMatchRequestHotelPolicyChildrenPolicy {
	s.IsAllowChildren = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildrenPolicy) SetIsAllowSharingBed(v bool) *InfoMatchRequestHotelPolicyChildrenPolicy {
	s.IsAllowSharingBed = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyChildrenPolicy) SetMaxChildrenAge(v int32) *InfoMatchRequestHotelPolicyChildrenPolicy {
	s.MaxChildrenAge = &v
	return s
}

type InfoMatchRequestHotelPolicyExtraBedPolicy struct {
	ExtraChildBedCharges []*InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem `json:"extra_child_bed_charges,omitempty" xml:"extra_child_bed_charges,omitempty" type:"Repeated"`
	SupportExtraBed      *bool                                                                `json:"support_extra_bed,omitempty" xml:"support_extra_bed,omitempty"`
	AdultExtraBedPrice   *int64                                                               `json:"adult_extra_bed_price,omitempty" xml:"adult_extra_bed_price,omitempty"`
	Currency             *string                                                              `json:"currency,omitempty" xml:"currency,omitempty"`
	ExtraBabyBedCharges  []*InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem  `json:"extra_baby_bed_charges,omitempty" xml:"extra_baby_bed_charges,omitempty" type:"Repeated"`
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicy) SetExtraChildBedCharges(v []*InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) *InfoMatchRequestHotelPolicyExtraBedPolicy {
	s.ExtraChildBedCharges = v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicy) SetSupportExtraBed(v bool) *InfoMatchRequestHotelPolicyExtraBedPolicy {
	s.SupportExtraBed = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicy) SetAdultExtraBedPrice(v int64) *InfoMatchRequestHotelPolicyExtraBedPolicy {
	s.AdultExtraBedPrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicy) SetCurrency(v string) *InfoMatchRequestHotelPolicyExtraBedPolicy {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicy) SetExtraBabyBedCharges(v []*InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) *InfoMatchRequestHotelPolicyExtraBedPolicy {
	s.ExtraBabyBedCharges = v
	return s
}

type InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem struct {
	End             *int32  `json:"end,omitempty" xml:"end,omitempty"`
	Start           *int32  `json:"start,omitempty" xml:"start,omitempty"`
	ChargeFrequency *string `json:"charge_frequency,omitempty" xml:"charge_frequency,omitempty"`
	ChargePercent   *int32  `json:"charge_percent,omitempty" xml:"charge_percent,omitempty"`
	ChargePrice     *int64  `json:"charge_price,omitempty" xml:"charge_price,omitempty"`
	Chargeable      *bool   `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string `json:"currency,omitempty" xml:"currency,omitempty"`
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetEnd(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.End = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetStart(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.Start = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetChargeFrequency(v string) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.ChargeFrequency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetChargePercent(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.ChargePercent = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetChargePrice(v int64) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.ChargePrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetChargeable(v bool) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem) SetCurrency(v string) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraBabyBedChargesItem {
	s.Currency = &v
	return s
}

type InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem struct {
	ChargePrice     *int64  `json:"charge_price,omitempty" xml:"charge_price,omitempty"`
	Chargeable      *bool   `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string `json:"currency,omitempty" xml:"currency,omitempty"`
	End             *int32  `json:"end,omitempty" xml:"end,omitempty"`
	Start           *int32  `json:"start,omitempty" xml:"start,omitempty"`
	ChargeFrequency *string `json:"charge_frequency,omitempty" xml:"charge_frequency,omitempty"`
	ChargePercent   *int32  `json:"charge_percent,omitempty" xml:"charge_percent,omitempty"`
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetChargePrice(v int64) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.ChargePrice = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetChargeable(v bool) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetCurrency(v string) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetEnd(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.End = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetStart(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.Start = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetChargeFrequency(v string) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.ChargeFrequency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem) SetChargePercent(v int32) *InfoMatchRequestHotelPolicyExtraBedPolicyExtraChildBedChargesItem {
	s.ChargePercent = &v
	return s
}

type InfoMatchRequestHotelPolicyParkingPolicy struct {
	ParkingProvided       *bool                                                      `json:"parking_provided,omitempty" xml:"parking_provided,omitempty"`
	ChargePointDetail     *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail `json:"charge_point_detail,omitempty" xml:"charge_point_detail,omitempty"`
	ChargingPointProvided *bool                                                      `json:"charging_point_provided,omitempty" xml:"charging_point_provided,omitempty"`
	ParkingDetails        *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails    `json:"parking_details,omitempty" xml:"parking_details,omitempty"`
}

func (s InfoMatchRequestHotelPolicyParkingPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyParkingPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyParkingPolicy) SetParkingProvided(v bool) *InfoMatchRequestHotelPolicyParkingPolicy {
	s.ParkingProvided = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicy) SetChargePointDetail(v *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail) *InfoMatchRequestHotelPolicyParkingPolicy {
	s.ChargePointDetail = v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicy) SetChargingPointProvided(v bool) *InfoMatchRequestHotelPolicyParkingPolicy {
	s.ChargingPointProvided = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicy) SetParkingDetails(v *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) *InfoMatchRequestHotelPolicyParkingPolicy {
	s.ParkingDetails = v
	return s
}

type InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail struct {
	Type     *int `json:"type,omitempty" xml:"type,omitempty"`
	Location *int `json:"location,omitempty" xml:"location,omitempty"`
}

func (s InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail) SetType(v int) *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail {
	s.Type = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail) SetLocation(v int) *InfoMatchRequestHotelPolicyParkingPolicyChargePointDetail {
	s.Location = &v
	return s
}

type InfoMatchRequestHotelPolicyParkingPolicyParkingDetails struct {
	ReserveRequired *bool    `json:"reserve_required,omitempty" xml:"reserve_required,omitempty"`
	Type            *int     `json:"type,omitempty" xml:"type,omitempty"`
	Amount          *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	Chargeable      *bool    `json:"chargeable,omitempty" xml:"chargeable,omitempty"`
	Currency        *string  `json:"currency,omitempty" xml:"currency,omitempty"`
	Location        *int     `json:"location,omitempty" xml:"location,omitempty"`
}

func (s InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetReserveRequired(v bool) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.ReserveRequired = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetType(v int) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.Type = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetAmount(v float64) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.Amount = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetChargeable(v bool) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.Chargeable = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetCurrency(v string) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.Currency = &v
	return s
}

func (s *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails) SetLocation(v int) *InfoMatchRequestHotelPolicyParkingPolicyParkingDetails {
	s.Location = &v
	return s
}

type InfoMatchRequestHotelPolicyPetPolicy struct {
	IsAllowPet *bool `json:"is_allow_pet,omitempty" xml:"is_allow_pet,omitempty"`
}

func (s InfoMatchRequestHotelPolicyPetPolicy) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchRequestHotelPolicyPetPolicy) GoString() string {
	return s.String()
}

func (s *InfoMatchRequestHotelPolicyPetPolicy) SetIsAllowPet(v bool) *InfoMatchRequestHotelPolicyPetPolicy {
	s.IsAllowPet = &v
	return s
}

type InfoMatchResponse struct {
	Data  *InfoMatchResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *InfoMatchResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s InfoMatchResponse) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchResponse) GoString() string {
	return s.String()
}

func (s *InfoMatchResponse) SetData(v *InfoMatchResponseData) *InfoMatchResponse {
	s.Data = v
	return s
}

func (s *InfoMatchResponse) SetExtra(v *InfoMatchResponseExtra) *InfoMatchResponse {
	s.Extra = v
	return s
}

type InfoMatchResponseData struct {
	Status        *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CategoryId    *int64  `json:"category_id,omitempty" xml:"category_id,omitempty"`
	HotelId       *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Message       *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	OutHotelId    *string `json:"out_hotel_id,omitempty" xml:"out_hotel_id,omitempty" require:"true"`
}

func (s InfoMatchResponseData) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchResponseData) GoString() string {
	return s.String()
}

func (s *InfoMatchResponseData) SetStatus(v int) *InfoMatchResponseData {
	s.Status = &v
	return s
}

func (s *InfoMatchResponseData) SetGwErrorCode(v int32) *InfoMatchResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *InfoMatchResponseData) SetGwDescription(v string) *InfoMatchResponseData {
	s.GwDescription = &v
	return s
}

func (s *InfoMatchResponseData) SetCategoryId(v int64) *InfoMatchResponseData {
	s.CategoryId = &v
	return s
}

func (s *InfoMatchResponseData) SetHotelId(v string) *InfoMatchResponseData {
	s.HotelId = &v
	return s
}

func (s *InfoMatchResponseData) SetMessage(v string) *InfoMatchResponseData {
	s.Message = &v
	return s
}

func (s *InfoMatchResponseData) SetOutHotelId(v string) *InfoMatchResponseData {
	s.OutHotelId = &v
	return s
}

type InfoMatchResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s InfoMatchResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s InfoMatchResponseExtra) GoString() string {
	return s.String()
}

func (s *InfoMatchResponseExtra) SetErrorCode(v int32) *InfoMatchResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *InfoMatchResponseExtra) SetLogid(v string) *InfoMatchResponseExtra {
	s.Logid = &v
	return s
}

func (s *InfoMatchResponseExtra) SetNow(v int64) *InfoMatchResponseExtra {
	s.Now = &v
	return s
}

func (s *InfoMatchResponseExtra) SetSubDescription(v string) *InfoMatchResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *InfoMatchResponseExtra) SetSubErrorCode(v int32) *InfoMatchResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *InfoMatchResponseExtra) SetDescription(v string) *InfoMatchResponseExtra {
	s.Description = &v
	return s
}

type InitVideoPartUploadRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s InitVideoPartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s InitVideoPartUploadRequest) GoString() string {
	return s.String()
}

func (s *InitVideoPartUploadRequest) SetOpenId(v string) *InitVideoPartUploadRequest {
	s.OpenId = &v
	return s
}

func (s *InitVideoPartUploadRequest) SetHeader(v map[string]*string) *InitVideoPartUploadRequest {
	s.Header = v
	return s
}

func (s *InitVideoPartUploadRequest) SetAccessToken(v string) *InitVideoPartUploadRequest {
	s.AccessToken = &v
	return s
}

type InitVideoPartUploadResponse struct {
	Extra *InitVideoPartUploadResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *InitVideoPartUploadResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s InitVideoPartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s InitVideoPartUploadResponse) GoString() string {
	return s.String()
}

func (s *InitVideoPartUploadResponse) SetExtra(v *InitVideoPartUploadResponseExtra) *InitVideoPartUploadResponse {
	s.Extra = v
	return s
}

func (s *InitVideoPartUploadResponse) SetData(v *InitVideoPartUploadResponseData) *InitVideoPartUploadResponse {
	s.Data = v
	return s
}

type InitVideoPartUploadResponseData struct {
	UploadId      *string `json:"upload_id,omitempty" xml:"upload_id,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s InitVideoPartUploadResponseData) String() string {
	return tea.Prettify(s)
}

func (s InitVideoPartUploadResponseData) GoString() string {
	return s.String()
}

func (s *InitVideoPartUploadResponseData) SetUploadId(v string) *InitVideoPartUploadResponseData {
	s.UploadId = &v
	return s
}

func (s *InitVideoPartUploadResponseData) SetGwErrorCode(v int32) *InitVideoPartUploadResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *InitVideoPartUploadResponseData) SetGwDescription(v string) *InitVideoPartUploadResponseData {
	s.GwDescription = &v
	return s
}

type InitVideoPartUploadResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s InitVideoPartUploadResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s InitVideoPartUploadResponseExtra) GoString() string {
	return s.String()
}

func (s *InitVideoPartUploadResponseExtra) SetSubErrorCode(v int32) *InitVideoPartUploadResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *InitVideoPartUploadResponseExtra) SetDescription(v string) *InitVideoPartUploadResponseExtra {
	s.Description = &v
	return s
}

func (s *InitVideoPartUploadResponseExtra) SetErrorCode(v int32) *InitVideoPartUploadResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *InitVideoPartUploadResponseExtra) SetLogid(v string) *InitVideoPartUploadResponseExtra {
	s.Logid = &v
	return s
}

func (s *InitVideoPartUploadResponseExtra) SetNow(v int64) *InitVideoPartUploadResponseExtra {
	s.Now = &v
	return s
}

func (s *InitVideoPartUploadResponseExtra) SetSubDescription(v string) *InitVideoPartUploadResponseExtra {
	s.SubDescription = &v
	return s
}

type InventoryPushRequest struct {
	AccessToken        *string                                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PromotionId        *string                                 `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	AccountId          *string                                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ApplicableDate     *InventoryPushRequestApplicableDate     `json:"applicable_date,omitempty" xml:"applicable_date,omitempty" require:"true"`
	ApplicableResource *InventoryPushRequestApplicableResource `json:"applicable_resource,omitempty" xml:"applicable_resource,omitempty" require:"true"`
	Inventory          *int64                                  `json:"inventory,omitempty" xml:"inventory,omitempty" require:"true"`
	Header             map[string]*string                      `json:"header,omitempty" xml:"header,omitempty"`
}

func (s InventoryPushRequest) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushRequest) GoString() string {
	return s.String()
}

func (s *InventoryPushRequest) SetAccessToken(v string) *InventoryPushRequest {
	s.AccessToken = &v
	return s
}

func (s *InventoryPushRequest) SetPromotionId(v string) *InventoryPushRequest {
	s.PromotionId = &v
	return s
}

func (s *InventoryPushRequest) SetAccountId(v string) *InventoryPushRequest {
	s.AccountId = &v
	return s
}

func (s *InventoryPushRequest) SetApplicableDate(v *InventoryPushRequestApplicableDate) *InventoryPushRequest {
	s.ApplicableDate = v
	return s
}

func (s *InventoryPushRequest) SetApplicableResource(v *InventoryPushRequestApplicableResource) *InventoryPushRequest {
	s.ApplicableResource = v
	return s
}

func (s *InventoryPushRequest) SetInventory(v int64) *InventoryPushRequest {
	s.Inventory = &v
	return s
}

func (s *InventoryPushRequest) SetHeader(v map[string]*string) *InventoryPushRequest {
	s.Header = v
	return s
}

type InventoryPushRequestApplicableDate struct {
	StartDays  *string  `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	EndDays    *string  `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
}

func (s InventoryPushRequestApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushRequestApplicableDate) GoString() string {
	return s.String()
}

func (s *InventoryPushRequestApplicableDate) SetStartDays(v string) *InventoryPushRequestApplicableDate {
	s.StartDays = &v
	return s
}

func (s *InventoryPushRequestApplicableDate) SetDaysOfWeek(v []*int32) *InventoryPushRequestApplicableDate {
	s.DaysOfWeek = v
	return s
}

func (s *InventoryPushRequestApplicableDate) SetEndDays(v string) *InventoryPushRequestApplicableDate {
	s.EndDays = &v
	return s
}

type InventoryPushRequestApplicableResource struct {
	Resources []*InventoryPushRequestApplicableResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s InventoryPushRequestApplicableResource) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushRequestApplicableResource) GoString() string {
	return s.String()
}

func (s *InventoryPushRequestApplicableResource) SetResources(v []*InventoryPushRequestApplicableResourceResourcesItem) *InventoryPushRequestApplicableResource {
	s.Resources = v
	return s
}

type InventoryPushRequestApplicableResourceResourcesItem struct {
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
}

func (s InventoryPushRequestApplicableResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushRequestApplicableResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *InventoryPushRequestApplicableResourceResourcesItem) SetRatePlanId(v string) *InventoryPushRequestApplicableResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

func (s *InventoryPushRequestApplicableResourceResourcesItem) SetRatePlanStatus(v int32) *InventoryPushRequestApplicableResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

type InventoryPushResponse struct {
	Data  *InventoryPushResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *InventoryPushResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s InventoryPushResponse) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushResponse) GoString() string {
	return s.String()
}

func (s *InventoryPushResponse) SetData(v *InventoryPushResponseData) *InventoryPushResponse {
	s.Data = v
	return s
}

func (s *InventoryPushResponse) SetExtra(v *InventoryPushResponseExtra) *InventoryPushResponse {
	s.Extra = v
	return s
}

type InventoryPushResponseData struct {
	Status        *int                                          `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	MessageDetail []*InventoryPushResponseDataMessageDetailItem `json:"message_detail,omitempty" xml:"message_detail,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PromotionId   *string                                       `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
}

func (s InventoryPushResponseData) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushResponseData) GoString() string {
	return s.String()
}

func (s *InventoryPushResponseData) SetStatus(v int) *InventoryPushResponseData {
	s.Status = &v
	return s
}

func (s *InventoryPushResponseData) SetMessageDetail(v []*InventoryPushResponseDataMessageDetailItem) *InventoryPushResponseData {
	s.MessageDetail = v
	return s
}

func (s *InventoryPushResponseData) SetGwErrorCode(v int32) *InventoryPushResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *InventoryPushResponseData) SetGwDescription(v string) *InventoryPushResponseData {
	s.GwDescription = &v
	return s
}

func (s *InventoryPushResponseData) SetPromotionId(v string) *InventoryPushResponseData {
	s.PromotionId = &v
	return s
}

type InventoryPushResponseDataMessageDetailItem struct {
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
}

func (s InventoryPushResponseDataMessageDetailItem) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushResponseDataMessageDetailItem) GoString() string {
	return s.String()
}

func (s *InventoryPushResponseDataMessageDetailItem) SetErrorMessage(v string) *InventoryPushResponseDataMessageDetailItem {
	s.ErrorMessage = &v
	return s
}

func (s *InventoryPushResponseDataMessageDetailItem) SetRatePlanId(v string) *InventoryPushResponseDataMessageDetailItem {
	s.RatePlanId = &v
	return s
}

type InventoryPushResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s InventoryPushResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s InventoryPushResponseExtra) GoString() string {
	return s.String()
}

func (s *InventoryPushResponseExtra) SetErrorCode(v int32) *InventoryPushResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *InventoryPushResponseExtra) SetLogid(v string) *InventoryPushResponseExtra {
	s.Logid = &v
	return s
}

func (s *InventoryPushResponseExtra) SetNow(v int64) *InventoryPushResponseExtra {
	s.Now = &v
	return s
}

func (s *InventoryPushResponseExtra) SetSubDescription(v string) *InventoryPushResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *InventoryPushResponseExtra) SetSubErrorCode(v int32) *InventoryPushResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *InventoryPushResponseExtra) SetDescription(v string) *InventoryPushResponseExtra {
	s.Description = &v
	return s
}

type InvoiceNotifyRequest struct {
	InvoiceUrls   []*string          `json:"invoice_urls,omitempty" xml:"invoice_urls,omitempty" type:"Repeated"`
	Message       *string            `json:"message,omitempty" xml:"message,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ApplyId       *string            `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	InvoiceStatus *int               `json:"invoice_status,omitempty" xml:"invoice_status,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s InvoiceNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s InvoiceNotifyRequest) GoString() string {
	return s.String()
}

func (s *InvoiceNotifyRequest) SetInvoiceUrls(v []*string) *InvoiceNotifyRequest {
	s.InvoiceUrls = v
	return s
}

func (s *InvoiceNotifyRequest) SetMessage(v string) *InvoiceNotifyRequest {
	s.Message = &v
	return s
}

func (s *InvoiceNotifyRequest) SetAccountId(v string) *InvoiceNotifyRequest {
	s.AccountId = &v
	return s
}

func (s *InvoiceNotifyRequest) SetApplyId(v string) *InvoiceNotifyRequest {
	s.ApplyId = &v
	return s
}

func (s *InvoiceNotifyRequest) SetInvoiceStatus(v int) *InvoiceNotifyRequest {
	s.InvoiceStatus = &v
	return s
}

func (s *InvoiceNotifyRequest) SetHeader(v map[string]*string) *InvoiceNotifyRequest {
	s.Header = v
	return s
}

func (s *InvoiceNotifyRequest) SetAccessToken(v string) *InvoiceNotifyRequest {
	s.AccessToken = &v
	return s
}

type InvoiceNotifyResponse struct {
	Data  *InvoiceNotifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *InvoiceNotifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s InvoiceNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s InvoiceNotifyResponse) GoString() string {
	return s.String()
}

func (s *InvoiceNotifyResponse) SetData(v *InvoiceNotifyResponseData) *InvoiceNotifyResponse {
	s.Data = v
	return s
}

func (s *InvoiceNotifyResponse) SetExtra(v *InvoiceNotifyResponseExtra) *InvoiceNotifyResponse {
	s.Extra = v
	return s
}

type InvoiceNotifyResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s InvoiceNotifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s InvoiceNotifyResponseData) GoString() string {
	return s.String()
}

func (s *InvoiceNotifyResponseData) SetErrorCode(v int32) *InvoiceNotifyResponseData {
	s.ErrorCode = &v
	return s
}

func (s *InvoiceNotifyResponseData) SetDescription(v string) *InvoiceNotifyResponseData {
	s.Description = &v
	return s
}

type InvoiceNotifyResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s InvoiceNotifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s InvoiceNotifyResponseExtra) GoString() string {
	return s.String()
}

func (s *InvoiceNotifyResponseExtra) SetSubErrorCode(v int32) *InvoiceNotifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *InvoiceNotifyResponseExtra) SetDescription(v string) *InvoiceNotifyResponseExtra {
	s.Description = &v
	return s
}

func (s *InvoiceNotifyResponseExtra) SetErrorCode(v int32) *InvoiceNotifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *InvoiceNotifyResponseExtra) SetLogid(v string) *InvoiceNotifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *InvoiceNotifyResponseExtra) SetNow(v int64) *InvoiceNotifyResponseExtra {
	s.Now = &v
	return s
}

func (s *InvoiceNotifyResponseExtra) SetSubDescription(v string) *InvoiceNotifyResponseExtra {
	s.SubDescription = &v
	return s
}

type ItemBcGetBaseRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemBcGetBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseRequest) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseRequest) SetOpenId(v string) *ItemBcGetBaseRequest {
	s.OpenId = &v
	return s
}

func (s *ItemBcGetBaseRequest) SetItemId(v string) *ItemBcGetBaseRequest {
	s.ItemId = &v
	return s
}

func (s *ItemBcGetBaseRequest) SetHeader(v map[string]*string) *ItemBcGetBaseRequest {
	s.Header = v
	return s
}

func (s *ItemBcGetBaseRequest) SetAccessToken(v string) *ItemBcGetBaseRequest {
	s.AccessToken = &v
	return s
}

type ItemBcGetBaseResponse struct {
	Data   *ItemBcGetBaseResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ItemBcGetBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseResponse) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseResponse) SetData(v *ItemBcGetBaseResponseData) *ItemBcGetBaseResponse {
	s.Data = v
	return s
}

func (s *ItemBcGetBaseResponse) SetErrNo(v int32) *ItemBcGetBaseResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemBcGetBaseResponse) SetErrMsg(v string) *ItemBcGetBaseResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemBcGetBaseResponse) SetLogId(v string) *ItemBcGetBaseResponse {
	s.LogId = &v
	return s
}

type ItemBcGetBaseResponseData struct {
	Extra *ItemBcGetBaseResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemBcGetBaseResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemBcGetBaseResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseResponseData) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseResponseData) SetExtra(v *ItemBcGetBaseResponseDataExtra) *ItemBcGetBaseResponseData {
	s.Extra = v
	return s
}

func (s *ItemBcGetBaseResponseData) SetData(v *ItemBcGetBaseResponseDataData) *ItemBcGetBaseResponseData {
	s.Data = v
	return s
}

type ItemBcGetBaseResponseDataData struct {
	Result *ItemBcGetBaseResponseDataDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ItemBcGetBaseResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseResponseDataData) SetResult(v *ItemBcGetBaseResponseDataDataResult) *ItemBcGetBaseResponseDataData {
	s.Result = v
	return s
}

type ItemBcGetBaseResponseDataDataResult struct {
	TotalShare      *int64   `json:"total_share,omitempty" xml:"total_share,omitempty" require:"true"`
	AvgPlayDuration *float64 `json:"avg_play_duration,omitempty" xml:"avg_play_duration,omitempty" require:"true"`
	TotalPlay       *int64   `json:"total_play,omitempty" xml:"total_play,omitempty" require:"true"`
	TotalLike       *int64   `json:"total_like,omitempty" xml:"total_like,omitempty" require:"true"`
	TotalComment    *int64   `json:"total_comment,omitempty" xml:"total_comment,omitempty" require:"true"`
}

func (s ItemBcGetBaseResponseDataDataResult) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseResponseDataDataResult) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseResponseDataDataResult) SetTotalShare(v int64) *ItemBcGetBaseResponseDataDataResult {
	s.TotalShare = &v
	return s
}

func (s *ItemBcGetBaseResponseDataDataResult) SetAvgPlayDuration(v float64) *ItemBcGetBaseResponseDataDataResult {
	s.AvgPlayDuration = &v
	return s
}

func (s *ItemBcGetBaseResponseDataDataResult) SetTotalPlay(v int64) *ItemBcGetBaseResponseDataDataResult {
	s.TotalPlay = &v
	return s
}

func (s *ItemBcGetBaseResponseDataDataResult) SetTotalLike(v int64) *ItemBcGetBaseResponseDataDataResult {
	s.TotalLike = &v
	return s
}

func (s *ItemBcGetBaseResponseDataDataResult) SetTotalComment(v int64) *ItemBcGetBaseResponseDataDataResult {
	s.TotalComment = &v
	return s
}

type ItemBcGetBaseResponseDataExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ItemBcGetBaseResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetBaseResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemBcGetBaseResponseDataExtra) SetErrorCode(v int32) *ItemBcGetBaseResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemBcGetBaseResponseDataExtra) SetDescription(v string) *ItemBcGetBaseResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemBcGetBaseResponseDataExtra) SetSubErrorCode(v int32) *ItemBcGetBaseResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemBcGetBaseResponseDataExtra) SetSubDescription(v string) *ItemBcGetBaseResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemBcGetBaseResponseDataExtra) SetLogid(v string) *ItemBcGetBaseResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemBcGetBaseResponseDataExtra) SetNow(v int64) *ItemBcGetBaseResponseDataExtra {
	s.Now = &v
	return s
}

type ItemBcGetCommentRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
}

func (s ItemBcGetCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentRequest) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentRequest) SetOpenId(v string) *ItemBcGetCommentRequest {
	s.OpenId = &v
	return s
}

func (s *ItemBcGetCommentRequest) SetItemId(v string) *ItemBcGetCommentRequest {
	s.ItemId = &v
	return s
}

func (s *ItemBcGetCommentRequest) SetHeader(v map[string]*string) *ItemBcGetCommentRequest {
	s.Header = v
	return s
}

func (s *ItemBcGetCommentRequest) SetAccessToken(v string) *ItemBcGetCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemBcGetCommentRequest) SetDateType(v int64) *ItemBcGetCommentRequest {
	s.DateType = &v
	return s
}

type ItemBcGetCommentResponse struct {
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemBcGetCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ItemBcGetCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentResponse) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentResponse) SetErrMsg(v string) *ItemBcGetCommentResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemBcGetCommentResponse) SetLogId(v string) *ItemBcGetCommentResponse {
	s.LogId = &v
	return s
}

func (s *ItemBcGetCommentResponse) SetData(v *ItemBcGetCommentResponseData) *ItemBcGetCommentResponse {
	s.Data = v
	return s
}

func (s *ItemBcGetCommentResponse) SetErrNo(v int32) *ItemBcGetCommentResponse {
	s.ErrNo = &v
	return s
}

type ItemBcGetCommentResponseData struct {
	Extra *ItemBcGetCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemBcGetCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemBcGetCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentResponseData) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentResponseData) SetExtra(v *ItemBcGetCommentResponseDataExtra) *ItemBcGetCommentResponseData {
	s.Extra = v
	return s
}

func (s *ItemBcGetCommentResponseData) SetData(v *ItemBcGetCommentResponseDataData) *ItemBcGetCommentResponseData {
	s.Data = v
	return s
}

type ItemBcGetCommentResponseDataData struct {
	ResultList []*ItemBcGetCommentResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemBcGetCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentResponseDataData) SetResultList(v []*ItemBcGetCommentResponseDataDataResultListItem) *ItemBcGetCommentResponseDataData {
	s.ResultList = v
	return s
}

type ItemBcGetCommentResponseDataDataResultListItem struct {
	Date    *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Comment *int64  `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
}

func (s ItemBcGetCommentResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentResponseDataDataResultListItem) SetDate(v string) *ItemBcGetCommentResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *ItemBcGetCommentResponseDataDataResultListItem) SetComment(v int64) *ItemBcGetCommentResponseDataDataResultListItem {
	s.Comment = &v
	return s
}

type ItemBcGetCommentResponseDataExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ItemBcGetCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemBcGetCommentResponseDataExtra) SetErrorCode(v int32) *ItemBcGetCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemBcGetCommentResponseDataExtra) SetDescription(v string) *ItemBcGetCommentResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemBcGetCommentResponseDataExtra) SetSubErrorCode(v int32) *ItemBcGetCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemBcGetCommentResponseDataExtra) SetSubDescription(v string) *ItemBcGetCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemBcGetCommentResponseDataExtra) SetLogid(v string) *ItemBcGetCommentResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemBcGetCommentResponseDataExtra) SetNow(v int64) *ItemBcGetCommentResponseDataExtra {
	s.Now = &v
	return s
}

type ItemBcGetLikeRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemBcGetLikeRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeRequest) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeRequest) SetDateType(v int64) *ItemBcGetLikeRequest {
	s.DateType = &v
	return s
}

func (s *ItemBcGetLikeRequest) SetOpenId(v string) *ItemBcGetLikeRequest {
	s.OpenId = &v
	return s
}

func (s *ItemBcGetLikeRequest) SetItemId(v string) *ItemBcGetLikeRequest {
	s.ItemId = &v
	return s
}

func (s *ItemBcGetLikeRequest) SetHeader(v map[string]*string) *ItemBcGetLikeRequest {
	s.Header = v
	return s
}

func (s *ItemBcGetLikeRequest) SetAccessToken(v string) *ItemBcGetLikeRequest {
	s.AccessToken = &v
	return s
}

type ItemBcGetLikeResponse struct {
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemBcGetLikeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s ItemBcGetLikeResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeResponse) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeResponse) SetLogId(v string) *ItemBcGetLikeResponse {
	s.LogId = &v
	return s
}

func (s *ItemBcGetLikeResponse) SetData(v *ItemBcGetLikeResponseData) *ItemBcGetLikeResponse {
	s.Data = v
	return s
}

func (s *ItemBcGetLikeResponse) SetErrNo(v int32) *ItemBcGetLikeResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemBcGetLikeResponse) SetErrMsg(v string) *ItemBcGetLikeResponse {
	s.ErrMsg = &v
	return s
}

type ItemBcGetLikeResponseData struct {
	Data  *ItemBcGetLikeResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemBcGetLikeResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemBcGetLikeResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeResponseData) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeResponseData) SetData(v *ItemBcGetLikeResponseDataData) *ItemBcGetLikeResponseData {
	s.Data = v
	return s
}

func (s *ItemBcGetLikeResponseData) SetExtra(v *ItemBcGetLikeResponseDataExtra) *ItemBcGetLikeResponseData {
	s.Extra = v
	return s
}

type ItemBcGetLikeResponseDataData struct {
	ResultList []*ItemBcGetLikeResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemBcGetLikeResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeResponseDataData) SetResultList(v []*ItemBcGetLikeResponseDataDataResultListItem) *ItemBcGetLikeResponseDataData {
	s.ResultList = v
	return s
}

type ItemBcGetLikeResponseDataDataResultListItem struct {
	Like *int64  `json:"like,omitempty" xml:"like,omitempty" require:"true"`
	Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s ItemBcGetLikeResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeResponseDataDataResultListItem) SetLike(v int64) *ItemBcGetLikeResponseDataDataResultListItem {
	s.Like = &v
	return s
}

func (s *ItemBcGetLikeResponseDataDataResultListItem) SetDate(v string) *ItemBcGetLikeResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type ItemBcGetLikeResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ItemBcGetLikeResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetLikeResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemBcGetLikeResponseDataExtra) SetSubDescription(v string) *ItemBcGetLikeResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemBcGetLikeResponseDataExtra) SetLogid(v string) *ItemBcGetLikeResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemBcGetLikeResponseDataExtra) SetNow(v int64) *ItemBcGetLikeResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemBcGetLikeResponseDataExtra) SetErrorCode(v int32) *ItemBcGetLikeResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemBcGetLikeResponseDataExtra) SetDescription(v string) *ItemBcGetLikeResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemBcGetLikeResponseDataExtra) SetSubErrorCode(v int32) *ItemBcGetLikeResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type ItemBcGetPlayRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ItemBcGetPlayRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayRequest) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayRequest) SetAccessToken(v string) *ItemBcGetPlayRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemBcGetPlayRequest) SetOpenId(v string) *ItemBcGetPlayRequest {
	s.OpenId = &v
	return s
}

func (s *ItemBcGetPlayRequest) SetItemId(v string) *ItemBcGetPlayRequest {
	s.ItemId = &v
	return s
}

func (s *ItemBcGetPlayRequest) SetDateType(v int64) *ItemBcGetPlayRequest {
	s.DateType = &v
	return s
}

func (s *ItemBcGetPlayRequest) SetHeader(v map[string]*string) *ItemBcGetPlayRequest {
	s.Header = v
	return s
}

type ItemBcGetPlayResponse struct {
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemBcGetPlayResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ItemBcGetPlayResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayResponse) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayResponse) SetErrNo(v int32) *ItemBcGetPlayResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemBcGetPlayResponse) SetErrMsg(v string) *ItemBcGetPlayResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemBcGetPlayResponse) SetLogId(v string) *ItemBcGetPlayResponse {
	s.LogId = &v
	return s
}

func (s *ItemBcGetPlayResponse) SetData(v *ItemBcGetPlayResponseData) *ItemBcGetPlayResponse {
	s.Data = v
	return s
}

type ItemBcGetPlayResponseData struct {
	Extra *ItemBcGetPlayResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemBcGetPlayResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemBcGetPlayResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayResponseData) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayResponseData) SetExtra(v *ItemBcGetPlayResponseDataExtra) *ItemBcGetPlayResponseData {
	s.Extra = v
	return s
}

func (s *ItemBcGetPlayResponseData) SetData(v *ItemBcGetPlayResponseDataData) *ItemBcGetPlayResponseData {
	s.Data = v
	return s
}

type ItemBcGetPlayResponseDataData struct {
	ResultList []*ItemBcGetPlayResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemBcGetPlayResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayResponseDataData) SetResultList(v []*ItemBcGetPlayResponseDataDataResultListItem) *ItemBcGetPlayResponseDataData {
	s.ResultList = v
	return s
}

type ItemBcGetPlayResponseDataDataResultListItem struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Play *int64  `json:"play,omitempty" xml:"play,omitempty" require:"true"`
}

func (s ItemBcGetPlayResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayResponseDataDataResultListItem) SetDate(v string) *ItemBcGetPlayResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *ItemBcGetPlayResponseDataDataResultListItem) SetPlay(v int64) *ItemBcGetPlayResponseDataDataResultListItem {
	s.Play = &v
	return s
}

type ItemBcGetPlayResponseDataExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ItemBcGetPlayResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetPlayResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemBcGetPlayResponseDataExtra) SetDescription(v string) *ItemBcGetPlayResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemBcGetPlayResponseDataExtra) SetSubErrorCode(v int32) *ItemBcGetPlayResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemBcGetPlayResponseDataExtra) SetSubDescription(v string) *ItemBcGetPlayResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemBcGetPlayResponseDataExtra) SetLogid(v string) *ItemBcGetPlayResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemBcGetPlayResponseDataExtra) SetNow(v int64) *ItemBcGetPlayResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemBcGetPlayResponseDataExtra) SetErrorCode(v int32) *ItemBcGetPlayResponseDataExtra {
	s.ErrorCode = &v
	return s
}

type ItemBcGetShareRequest struct {
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemBcGetShareRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareRequest) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareRequest) SetItemId(v string) *ItemBcGetShareRequest {
	s.ItemId = &v
	return s
}

func (s *ItemBcGetShareRequest) SetDateType(v int64) *ItemBcGetShareRequest {
	s.DateType = &v
	return s
}

func (s *ItemBcGetShareRequest) SetOpenId(v string) *ItemBcGetShareRequest {
	s.OpenId = &v
	return s
}

func (s *ItemBcGetShareRequest) SetHeader(v map[string]*string) *ItemBcGetShareRequest {
	s.Header = v
	return s
}

func (s *ItemBcGetShareRequest) SetAccessToken(v string) *ItemBcGetShareRequest {
	s.AccessToken = &v
	return s
}

type ItemBcGetShareResponse struct {
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemBcGetShareResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s ItemBcGetShareResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareResponse) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareResponse) SetLogId(v string) *ItemBcGetShareResponse {
	s.LogId = &v
	return s
}

func (s *ItemBcGetShareResponse) SetData(v *ItemBcGetShareResponseData) *ItemBcGetShareResponse {
	s.Data = v
	return s
}

func (s *ItemBcGetShareResponse) SetErrNo(v int32) *ItemBcGetShareResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemBcGetShareResponse) SetErrMsg(v string) *ItemBcGetShareResponse {
	s.ErrMsg = &v
	return s
}

type ItemBcGetShareResponseData struct {
	Data  *ItemBcGetShareResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemBcGetShareResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemBcGetShareResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareResponseData) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareResponseData) SetData(v *ItemBcGetShareResponseDataData) *ItemBcGetShareResponseData {
	s.Data = v
	return s
}

func (s *ItemBcGetShareResponseData) SetExtra(v *ItemBcGetShareResponseDataExtra) *ItemBcGetShareResponseData {
	s.Extra = v
	return s
}

type ItemBcGetShareResponseDataData struct {
	ResultList []*ItemBcGetShareResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemBcGetShareResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareResponseDataData) SetResultList(v []*ItemBcGetShareResponseDataDataResultListItem) *ItemBcGetShareResponseDataData {
	s.ResultList = v
	return s
}

type ItemBcGetShareResponseDataDataResultListItem struct {
	Share *int64  `json:"share,omitempty" xml:"share,omitempty" require:"true"`
	Date  *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s ItemBcGetShareResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareResponseDataDataResultListItem) SetShare(v int64) *ItemBcGetShareResponseDataDataResultListItem {
	s.Share = &v
	return s
}

func (s *ItemBcGetShareResponseDataDataResultListItem) SetDate(v string) *ItemBcGetShareResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type ItemBcGetShareResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ItemBcGetShareResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemBcGetShareResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemBcGetShareResponseDataExtra) SetSubDescription(v string) *ItemBcGetShareResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemBcGetShareResponseDataExtra) SetLogid(v string) *ItemBcGetShareResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemBcGetShareResponseDataExtra) SetNow(v int64) *ItemBcGetShareResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemBcGetShareResponseDataExtra) SetErrorCode(v int32) *ItemBcGetShareResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemBcGetShareResponseDataExtra) SetDescription(v string) *ItemBcGetShareResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemBcGetShareResponseDataExtra) SetSubErrorCode(v int32) *ItemBcGetShareResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type ItemGetBaseRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
}

func (s ItemGetBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseRequest) GoString() string {
	return s.String()
}

func (s *ItemGetBaseRequest) SetOpenId(v string) *ItemGetBaseRequest {
	s.OpenId = &v
	return s
}

func (s *ItemGetBaseRequest) SetHeader(v map[string]*string) *ItemGetBaseRequest {
	s.Header = v
	return s
}

func (s *ItemGetBaseRequest) SetAccessToken(v string) *ItemGetBaseRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemGetBaseRequest) SetItemId(v string) *ItemGetBaseRequest {
	s.ItemId = &v
	return s
}

type ItemGetBaseResponse struct {
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemGetBaseResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s ItemGetBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseResponse) GoString() string {
	return s.String()
}

func (s *ItemGetBaseResponse) SetLogId(v string) *ItemGetBaseResponse {
	s.LogId = &v
	return s
}

func (s *ItemGetBaseResponse) SetData(v *ItemGetBaseResponseData) *ItemGetBaseResponse {
	s.Data = v
	return s
}

func (s *ItemGetBaseResponse) SetErrNo(v int32) *ItemGetBaseResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemGetBaseResponse) SetErrMsg(v string) *ItemGetBaseResponse {
	s.ErrMsg = &v
	return s
}

type ItemGetBaseResponseData struct {
	Data  *ItemGetBaseResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemGetBaseResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemGetBaseResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseResponseData) GoString() string {
	return s.String()
}

func (s *ItemGetBaseResponseData) SetData(v *ItemGetBaseResponseDataData) *ItemGetBaseResponseData {
	s.Data = v
	return s
}

func (s *ItemGetBaseResponseData) SetExtra(v *ItemGetBaseResponseDataExtra) *ItemGetBaseResponseData {
	s.Extra = v
	return s
}

type ItemGetBaseResponseDataData struct {
	Result *ItemGetBaseResponseDataDataResult `json:"result,omitempty" xml:"result,omitempty" require:"true"`
}

func (s ItemGetBaseResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemGetBaseResponseDataData) SetResult(v *ItemGetBaseResponseDataDataResult) *ItemGetBaseResponseDataData {
	s.Result = v
	return s
}

type ItemGetBaseResponseDataDataResult struct {
	TotalLike       *int64   `json:"total_like,omitempty" xml:"total_like,omitempty" require:"true"`
	TotalComment    *int64   `json:"total_comment,omitempty" xml:"total_comment,omitempty" require:"true"`
	TotalShare      *int64   `json:"total_share,omitempty" xml:"total_share,omitempty" require:"true"`
	AvgPlayDuration *float64 `json:"avg_play_duration,omitempty" xml:"avg_play_duration,omitempty" require:"true"`
	TotalPlay       *int64   `json:"total_play,omitempty" xml:"total_play,omitempty" require:"true"`
}

func (s ItemGetBaseResponseDataDataResult) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseResponseDataDataResult) GoString() string {
	return s.String()
}

func (s *ItemGetBaseResponseDataDataResult) SetTotalLike(v int64) *ItemGetBaseResponseDataDataResult {
	s.TotalLike = &v
	return s
}

func (s *ItemGetBaseResponseDataDataResult) SetTotalComment(v int64) *ItemGetBaseResponseDataDataResult {
	s.TotalComment = &v
	return s
}

func (s *ItemGetBaseResponseDataDataResult) SetTotalShare(v int64) *ItemGetBaseResponseDataDataResult {
	s.TotalShare = &v
	return s
}

func (s *ItemGetBaseResponseDataDataResult) SetAvgPlayDuration(v float64) *ItemGetBaseResponseDataDataResult {
	s.AvgPlayDuration = &v
	return s
}

func (s *ItemGetBaseResponseDataDataResult) SetTotalPlay(v int64) *ItemGetBaseResponseDataDataResult {
	s.TotalPlay = &v
	return s
}

type ItemGetBaseResponseDataExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ItemGetBaseResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemGetBaseResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemGetBaseResponseDataExtra) SetSubErrorCode(v int32) *ItemGetBaseResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemGetBaseResponseDataExtra) SetSubDescription(v string) *ItemGetBaseResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemGetBaseResponseDataExtra) SetLogid(v string) *ItemGetBaseResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemGetBaseResponseDataExtra) SetNow(v int64) *ItemGetBaseResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemGetBaseResponseDataExtra) SetErrorCode(v int32) *ItemGetBaseResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemGetBaseResponseDataExtra) SetDescription(v string) *ItemGetBaseResponseDataExtra {
	s.Description = &v
	return s
}

type ItemGetCommentRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
}

func (s ItemGetCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentRequest) GoString() string {
	return s.String()
}

func (s *ItemGetCommentRequest) SetDateType(v int64) *ItemGetCommentRequest {
	s.DateType = &v
	return s
}

func (s *ItemGetCommentRequest) SetHeader(v map[string]*string) *ItemGetCommentRequest {
	s.Header = v
	return s
}

func (s *ItemGetCommentRequest) SetAccessToken(v string) *ItemGetCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemGetCommentRequest) SetOpenId(v string) *ItemGetCommentRequest {
	s.OpenId = &v
	return s
}

func (s *ItemGetCommentRequest) SetItemId(v string) *ItemGetCommentRequest {
	s.ItemId = &v
	return s
}

type ItemGetCommentResponse struct {
	Data   *ItemGetCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ItemGetCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentResponse) GoString() string {
	return s.String()
}

func (s *ItemGetCommentResponse) SetData(v *ItemGetCommentResponseData) *ItemGetCommentResponse {
	s.Data = v
	return s
}

func (s *ItemGetCommentResponse) SetErrNo(v int32) *ItemGetCommentResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemGetCommentResponse) SetErrMsg(v string) *ItemGetCommentResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemGetCommentResponse) SetLogId(v string) *ItemGetCommentResponse {
	s.LogId = &v
	return s
}

type ItemGetCommentResponseData struct {
	Extra *ItemGetCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemGetCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemGetCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentResponseData) GoString() string {
	return s.String()
}

func (s *ItemGetCommentResponseData) SetExtra(v *ItemGetCommentResponseDataExtra) *ItemGetCommentResponseData {
	s.Extra = v
	return s
}

func (s *ItemGetCommentResponseData) SetData(v *ItemGetCommentResponseDataData) *ItemGetCommentResponseData {
	s.Data = v
	return s
}

type ItemGetCommentResponseDataData struct {
	ResultList []*ItemGetCommentResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemGetCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemGetCommentResponseDataData) SetResultList(v []*ItemGetCommentResponseDataDataResultListItem) *ItemGetCommentResponseDataData {
	s.ResultList = v
	return s
}

type ItemGetCommentResponseDataDataResultListItem struct {
	Comment *int64  `json:"comment,omitempty" xml:"comment,omitempty" require:"true"`
	Date    *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s ItemGetCommentResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemGetCommentResponseDataDataResultListItem) SetComment(v int64) *ItemGetCommentResponseDataDataResultListItem {
	s.Comment = &v
	return s
}

func (s *ItemGetCommentResponseDataDataResultListItem) SetDate(v string) *ItemGetCommentResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type ItemGetCommentResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ItemGetCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemGetCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemGetCommentResponseDataExtra) SetSubDescription(v string) *ItemGetCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemGetCommentResponseDataExtra) SetLogid(v string) *ItemGetCommentResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemGetCommentResponseDataExtra) SetNow(v int64) *ItemGetCommentResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemGetCommentResponseDataExtra) SetErrorCode(v int32) *ItemGetCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemGetCommentResponseDataExtra) SetDescription(v string) *ItemGetCommentResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemGetCommentResponseDataExtra) SetSubErrorCode(v int32) *ItemGetCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type ItemGetLikeRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemGetLikeRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeRequest) GoString() string {
	return s.String()
}

func (s *ItemGetLikeRequest) SetOpenId(v string) *ItemGetLikeRequest {
	s.OpenId = &v
	return s
}

func (s *ItemGetLikeRequest) SetItemId(v string) *ItemGetLikeRequest {
	s.ItemId = &v
	return s
}

func (s *ItemGetLikeRequest) SetDateType(v int64) *ItemGetLikeRequest {
	s.DateType = &v
	return s
}

func (s *ItemGetLikeRequest) SetHeader(v map[string]*string) *ItemGetLikeRequest {
	s.Header = v
	return s
}

func (s *ItemGetLikeRequest) SetAccessToken(v string) *ItemGetLikeRequest {
	s.AccessToken = &v
	return s
}

type ItemGetLikeResponse struct {
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemGetLikeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ItemGetLikeResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeResponse) GoString() string {
	return s.String()
}

func (s *ItemGetLikeResponse) SetErrNo(v int32) *ItemGetLikeResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemGetLikeResponse) SetErrMsg(v string) *ItemGetLikeResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemGetLikeResponse) SetLogId(v string) *ItemGetLikeResponse {
	s.LogId = &v
	return s
}

func (s *ItemGetLikeResponse) SetData(v *ItemGetLikeResponseData) *ItemGetLikeResponse {
	s.Data = v
	return s
}

type ItemGetLikeResponseData struct {
	Extra *ItemGetLikeResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemGetLikeResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemGetLikeResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeResponseData) GoString() string {
	return s.String()
}

func (s *ItemGetLikeResponseData) SetExtra(v *ItemGetLikeResponseDataExtra) *ItemGetLikeResponseData {
	s.Extra = v
	return s
}

func (s *ItemGetLikeResponseData) SetData(v *ItemGetLikeResponseDataData) *ItemGetLikeResponseData {
	s.Data = v
	return s
}

type ItemGetLikeResponseDataData struct {
	ResultList []*ItemGetLikeResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemGetLikeResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemGetLikeResponseDataData) SetResultList(v []*ItemGetLikeResponseDataDataResultListItem) *ItemGetLikeResponseDataData {
	s.ResultList = v
	return s
}

type ItemGetLikeResponseDataDataResultListItem struct {
	Like *int64  `json:"like,omitempty" xml:"like,omitempty" require:"true"`
	Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s ItemGetLikeResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemGetLikeResponseDataDataResultListItem) SetLike(v int64) *ItemGetLikeResponseDataDataResultListItem {
	s.Like = &v
	return s
}

func (s *ItemGetLikeResponseDataDataResultListItem) SetDate(v string) *ItemGetLikeResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type ItemGetLikeResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ItemGetLikeResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemGetLikeResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemGetLikeResponseDataExtra) SetNow(v int64) *ItemGetLikeResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemGetLikeResponseDataExtra) SetErrorCode(v int32) *ItemGetLikeResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemGetLikeResponseDataExtra) SetDescription(v string) *ItemGetLikeResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemGetLikeResponseDataExtra) SetSubErrorCode(v int32) *ItemGetLikeResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemGetLikeResponseDataExtra) SetSubDescription(v string) *ItemGetLikeResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemGetLikeResponseDataExtra) SetLogid(v string) *ItemGetLikeResponseDataExtra {
	s.Logid = &v
	return s
}

type ItemGetPlayRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemGetPlayRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayRequest) GoString() string {
	return s.String()
}

func (s *ItemGetPlayRequest) SetDateType(v int64) *ItemGetPlayRequest {
	s.DateType = &v
	return s
}

func (s *ItemGetPlayRequest) SetOpenId(v string) *ItemGetPlayRequest {
	s.OpenId = &v
	return s
}

func (s *ItemGetPlayRequest) SetItemId(v string) *ItemGetPlayRequest {
	s.ItemId = &v
	return s
}

func (s *ItemGetPlayRequest) SetHeader(v map[string]*string) *ItemGetPlayRequest {
	s.Header = v
	return s
}

func (s *ItemGetPlayRequest) SetAccessToken(v string) *ItemGetPlayRequest {
	s.AccessToken = &v
	return s
}

type ItemGetPlayResponse struct {
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemGetPlayResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ItemGetPlayResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayResponse) GoString() string {
	return s.String()
}

func (s *ItemGetPlayResponse) SetErrNo(v int32) *ItemGetPlayResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemGetPlayResponse) SetErrMsg(v string) *ItemGetPlayResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemGetPlayResponse) SetLogId(v string) *ItemGetPlayResponse {
	s.LogId = &v
	return s
}

func (s *ItemGetPlayResponse) SetData(v *ItemGetPlayResponseData) *ItemGetPlayResponse {
	s.Data = v
	return s
}

type ItemGetPlayResponseData struct {
	Data  *ItemGetPlayResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemGetPlayResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemGetPlayResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayResponseData) GoString() string {
	return s.String()
}

func (s *ItemGetPlayResponseData) SetData(v *ItemGetPlayResponseDataData) *ItemGetPlayResponseData {
	s.Data = v
	return s
}

func (s *ItemGetPlayResponseData) SetExtra(v *ItemGetPlayResponseDataExtra) *ItemGetPlayResponseData {
	s.Extra = v
	return s
}

type ItemGetPlayResponseDataData struct {
	ResultList []*ItemGetPlayResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemGetPlayResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemGetPlayResponseDataData) SetResultList(v []*ItemGetPlayResponseDataDataResultListItem) *ItemGetPlayResponseDataData {
	s.ResultList = v
	return s
}

type ItemGetPlayResponseDataDataResultListItem struct {
	Play *int64  `json:"play,omitempty" xml:"play,omitempty" require:"true"`
	Date *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s ItemGetPlayResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemGetPlayResponseDataDataResultListItem) SetPlay(v int64) *ItemGetPlayResponseDataDataResultListItem {
	s.Play = &v
	return s
}

func (s *ItemGetPlayResponseDataDataResultListItem) SetDate(v string) *ItemGetPlayResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type ItemGetPlayResponseDataExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ItemGetPlayResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemGetPlayResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemGetPlayResponseDataExtra) SetDescription(v string) *ItemGetPlayResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemGetPlayResponseDataExtra) SetSubErrorCode(v int32) *ItemGetPlayResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemGetPlayResponseDataExtra) SetSubDescription(v string) *ItemGetPlayResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemGetPlayResponseDataExtra) SetLogid(v string) *ItemGetPlayResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemGetPlayResponseDataExtra) SetNow(v int64) *ItemGetPlayResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemGetPlayResponseDataExtra) SetErrorCode(v int32) *ItemGetPlayResponseDataExtra {
	s.ErrorCode = &v
	return s
}

type ItemGetShareRequest struct {
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ItemGetShareRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareRequest) GoString() string {
	return s.String()
}

func (s *ItemGetShareRequest) SetItemId(v string) *ItemGetShareRequest {
	s.ItemId = &v
	return s
}

func (s *ItemGetShareRequest) SetDateType(v int64) *ItemGetShareRequest {
	s.DateType = &v
	return s
}

func (s *ItemGetShareRequest) SetOpenId(v string) *ItemGetShareRequest {
	s.OpenId = &v
	return s
}

func (s *ItemGetShareRequest) SetHeader(v map[string]*string) *ItemGetShareRequest {
	s.Header = v
	return s
}

func (s *ItemGetShareRequest) SetAccessToken(v string) *ItemGetShareRequest {
	s.AccessToken = &v
	return s
}

type ItemGetShareResponse struct {
	ErrNo  *int32                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemGetShareResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ItemGetShareResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareResponse) GoString() string {
	return s.String()
}

func (s *ItemGetShareResponse) SetErrNo(v int32) *ItemGetShareResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemGetShareResponse) SetErrMsg(v string) *ItemGetShareResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemGetShareResponse) SetLogId(v string) *ItemGetShareResponse {
	s.LogId = &v
	return s
}

func (s *ItemGetShareResponse) SetData(v *ItemGetShareResponseData) *ItemGetShareResponse {
	s.Data = v
	return s
}

type ItemGetShareResponseData struct {
	Data  *ItemGetShareResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemGetShareResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemGetShareResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareResponseData) GoString() string {
	return s.String()
}

func (s *ItemGetShareResponseData) SetData(v *ItemGetShareResponseDataData) *ItemGetShareResponseData {
	s.Data = v
	return s
}

func (s *ItemGetShareResponseData) SetExtra(v *ItemGetShareResponseDataExtra) *ItemGetShareResponseData {
	s.Extra = v
	return s
}

type ItemGetShareResponseDataData struct {
	ResultList []*ItemGetShareResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemGetShareResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemGetShareResponseDataData) SetResultList(v []*ItemGetShareResponseDataDataResultListItem) *ItemGetShareResponseDataData {
	s.ResultList = v
	return s
}

type ItemGetShareResponseDataDataResultListItem struct {
	Date  *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Share *int64  `json:"share,omitempty" xml:"share,omitempty" require:"true"`
}

func (s ItemGetShareResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *ItemGetShareResponseDataDataResultListItem) SetDate(v string) *ItemGetShareResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *ItemGetShareResponseDataDataResultListItem) SetShare(v int64) *ItemGetShareResponseDataDataResultListItem {
	s.Share = &v
	return s
}

type ItemGetShareResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ItemGetShareResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemGetShareResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemGetShareResponseDataExtra) SetSubDescription(v string) *ItemGetShareResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemGetShareResponseDataExtra) SetLogid(v string) *ItemGetShareResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemGetShareResponseDataExtra) SetNow(v int64) *ItemGetShareResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemGetShareResponseDataExtra) SetErrorCode(v int32) *ItemGetShareResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemGetShareResponseDataExtra) SetDescription(v string) *ItemGetShareResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemGetShareResponseDataExtra) SetSubErrorCode(v int32) *ItemGetShareResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type ItemListCommentReplyRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	CommentId   *string            `json:"comment_id,omitempty" xml:"comment_id,omitempty" require:"true"`
	SortType    *string            `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ItemListCommentReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyRequest) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyRequest) SetAccessToken(v string) *ItemListCommentReplyRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetItemId(v string) *ItemListCommentReplyRequest {
	s.ItemId = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetCommentId(v string) *ItemListCommentReplyRequest {
	s.CommentId = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetSortType(v string) *ItemListCommentReplyRequest {
	s.SortType = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetOpenId(v string) *ItemListCommentReplyRequest {
	s.OpenId = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetCursor(v int64) *ItemListCommentReplyRequest {
	s.Cursor = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetCount(v int32) *ItemListCommentReplyRequest {
	s.Count = &v
	return s
}

func (s *ItemListCommentReplyRequest) SetHeader(v map[string]*string) *ItemListCommentReplyRequest {
	s.Header = v
	return s
}

type ItemListCommentReplyResponse struct {
	Data   *ItemListCommentReplyResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ItemListCommentReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyResponse) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyResponse) SetData(v *ItemListCommentReplyResponseData) *ItemListCommentReplyResponse {
	s.Data = v
	return s
}

func (s *ItemListCommentReplyResponse) SetErrNo(v int32) *ItemListCommentReplyResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemListCommentReplyResponse) SetErrMsg(v string) *ItemListCommentReplyResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemListCommentReplyResponse) SetLogId(v string) *ItemListCommentReplyResponse {
	s.LogId = &v
	return s
}

type ItemListCommentReplyResponseData struct {
	Extra *ItemListCommentReplyResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ItemListCommentReplyResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ItemListCommentReplyResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyResponseData) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyResponseData) SetExtra(v *ItemListCommentReplyResponseDataExtra) *ItemListCommentReplyResponseData {
	s.Extra = v
	return s
}

func (s *ItemListCommentReplyResponseData) SetData(v *ItemListCommentReplyResponseDataData) *ItemListCommentReplyResponseData {
	s.Data = v
	return s
}

type ItemListCommentReplyResponseDataData struct {
	Cursor  *int64                                          `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore *bool                                           `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	List    []*ItemListCommentReplyResponseDataDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s ItemListCommentReplyResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyResponseDataData) SetCursor(v int64) *ItemListCommentReplyResponseDataData {
	s.Cursor = &v
	return s
}

func (s *ItemListCommentReplyResponseDataData) SetHasMore(v bool) *ItemListCommentReplyResponseDataData {
	s.HasMore = &v
	return s
}

func (s *ItemListCommentReplyResponseDataData) SetList(v []*ItemListCommentReplyResponseDataDataListItem) *ItemListCommentReplyResponseDataData {
	s.List = v
	return s
}

type ItemListCommentReplyResponseDataDataListItem struct {
	Content           *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	NickName          *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	DiggCount         *int32  `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	CommentId         *string `json:"comment_id,omitempty" xml:"comment_id,omitempty" require:"true"`
	CreateTime        *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	CommentUserId     *string `json:"comment_user_id,omitempty" xml:"comment_user_id,omitempty" require:"true"`
	ReplyCommentId    *string `json:"reply_comment_id,omitempty" xml:"reply_comment_id,omitempty"`
	Avatar            *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	ReplyCommentTotal *int32  `json:"reply_comment_total,omitempty" xml:"reply_comment_total,omitempty" require:"true"`
	Top               *bool   `json:"top,omitempty" xml:"top,omitempty" require:"true"`
}

func (s ItemListCommentReplyResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetContent(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.Content = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetNickName(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.NickName = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetDiggCount(v int32) *ItemListCommentReplyResponseDataDataListItem {
	s.DiggCount = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetCommentId(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.CommentId = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetCreateTime(v int64) *ItemListCommentReplyResponseDataDataListItem {
	s.CreateTime = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetCommentUserId(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.CommentUserId = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetReplyCommentId(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.ReplyCommentId = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetAvatar(v string) *ItemListCommentReplyResponseDataDataListItem {
	s.Avatar = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetReplyCommentTotal(v int32) *ItemListCommentReplyResponseDataDataListItem {
	s.ReplyCommentTotal = &v
	return s
}

func (s *ItemListCommentReplyResponseDataDataListItem) SetTop(v bool) *ItemListCommentReplyResponseDataDataListItem {
	s.Top = &v
	return s
}

type ItemListCommentReplyResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ItemListCommentReplyResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentReplyResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemListCommentReplyResponseDataExtra) SetNow(v int64) *ItemListCommentReplyResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemListCommentReplyResponseDataExtra) SetErrorCode(v int32) *ItemListCommentReplyResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemListCommentReplyResponseDataExtra) SetDescription(v string) *ItemListCommentReplyResponseDataExtra {
	s.Description = &v
	return s
}

func (s *ItemListCommentReplyResponseDataExtra) SetSubErrorCode(v int32) *ItemListCommentReplyResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemListCommentReplyResponseDataExtra) SetSubDescription(v string) *ItemListCommentReplyResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemListCommentReplyResponseDataExtra) SetLogid(v string) *ItemListCommentReplyResponseDataExtra {
	s.Logid = &v
	return s
}

type ItemListCommentRequest struct {
	SortType    *string            `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
}

func (s ItemListCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentRequest) GoString() string {
	return s.String()
}

func (s *ItemListCommentRequest) SetSortType(v string) *ItemListCommentRequest {
	s.SortType = &v
	return s
}

func (s *ItemListCommentRequest) SetOpenId(v string) *ItemListCommentRequest {
	s.OpenId = &v
	return s
}

func (s *ItemListCommentRequest) SetCursor(v int64) *ItemListCommentRequest {
	s.Cursor = &v
	return s
}

func (s *ItemListCommentRequest) SetHeader(v map[string]*string) *ItemListCommentRequest {
	s.Header = v
	return s
}

func (s *ItemListCommentRequest) SetAccessToken(v string) *ItemListCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemListCommentRequest) SetCount(v int32) *ItemListCommentRequest {
	s.Count = &v
	return s
}

func (s *ItemListCommentRequest) SetItemId(v string) *ItemListCommentRequest {
	s.ItemId = &v
	return s
}

type ItemListCommentResponse struct {
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ItemListCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ItemListCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentResponse) GoString() string {
	return s.String()
}

func (s *ItemListCommentResponse) SetErrNo(v int32) *ItemListCommentResponse {
	s.ErrNo = &v
	return s
}

func (s *ItemListCommentResponse) SetErrMsg(v string) *ItemListCommentResponse {
	s.ErrMsg = &v
	return s
}

func (s *ItemListCommentResponse) SetLogId(v string) *ItemListCommentResponse {
	s.LogId = &v
	return s
}

func (s *ItemListCommentResponse) SetData(v *ItemListCommentResponseData) *ItemListCommentResponse {
	s.Data = v
	return s
}

type ItemListCommentResponseData struct {
	Data  *ItemListCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ItemListCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ItemListCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentResponseData) GoString() string {
	return s.String()
}

func (s *ItemListCommentResponseData) SetData(v *ItemListCommentResponseDataData) *ItemListCommentResponseData {
	s.Data = v
	return s
}

func (s *ItemListCommentResponseData) SetExtra(v *ItemListCommentResponseDataExtra) *ItemListCommentResponseData {
	s.Extra = v
	return s
}

type ItemListCommentResponseDataData struct {
	HasMore *bool                                      `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	List    []*ItemListCommentResponseDataDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
	Cursor  *int64                                     `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
}

func (s ItemListCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *ItemListCommentResponseDataData) SetHasMore(v bool) *ItemListCommentResponseDataData {
	s.HasMore = &v
	return s
}

func (s *ItemListCommentResponseDataData) SetList(v []*ItemListCommentResponseDataDataListItem) *ItemListCommentResponseDataData {
	s.List = v
	return s
}

func (s *ItemListCommentResponseDataData) SetCursor(v int64) *ItemListCommentResponseDataData {
	s.Cursor = &v
	return s
}

type ItemListCommentResponseDataDataListItem struct {
	CommentUserId     *string `json:"comment_user_id,omitempty" xml:"comment_user_id,omitempty" require:"true"`
	CommentId         *string `json:"comment_id,omitempty" xml:"comment_id,omitempty" require:"true"`
	DiggCount         *int32  `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	Avatar            *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreateTime        *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	NickName          *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	Content           *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	ReplyCommentTotal *int32  `json:"reply_comment_total,omitempty" xml:"reply_comment_total,omitempty" require:"true"`
	ReplyCommentId    *string `json:"reply_comment_id,omitempty" xml:"reply_comment_id,omitempty"`
	Top               *bool   `json:"top,omitempty" xml:"top,omitempty" require:"true"`
}

func (s ItemListCommentResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *ItemListCommentResponseDataDataListItem) SetCommentUserId(v string) *ItemListCommentResponseDataDataListItem {
	s.CommentUserId = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetCommentId(v string) *ItemListCommentResponseDataDataListItem {
	s.CommentId = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetDiggCount(v int32) *ItemListCommentResponseDataDataListItem {
	s.DiggCount = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetAvatar(v string) *ItemListCommentResponseDataDataListItem {
	s.Avatar = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetCreateTime(v int64) *ItemListCommentResponseDataDataListItem {
	s.CreateTime = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetNickName(v string) *ItemListCommentResponseDataDataListItem {
	s.NickName = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetContent(v string) *ItemListCommentResponseDataDataListItem {
	s.Content = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetReplyCommentTotal(v int32) *ItemListCommentResponseDataDataListItem {
	s.ReplyCommentTotal = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetReplyCommentId(v string) *ItemListCommentResponseDataDataListItem {
	s.ReplyCommentId = &v
	return s
}

func (s *ItemListCommentResponseDataDataListItem) SetTop(v bool) *ItemListCommentResponseDataDataListItem {
	s.Top = &v
	return s
}

type ItemListCommentResponseDataExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ItemListCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s ItemListCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *ItemListCommentResponseDataExtra) SetSubErrorCode(v int32) *ItemListCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ItemListCommentResponseDataExtra) SetSubDescription(v string) *ItemListCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *ItemListCommentResponseDataExtra) SetLogid(v string) *ItemListCommentResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *ItemListCommentResponseDataExtra) SetNow(v int64) *ItemListCommentResponseDataExtra {
	s.Now = &v
	return s
}

func (s *ItemListCommentResponseDataExtra) SetErrorCode(v int32) *ItemListCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *ItemListCommentResponseDataExtra) SetDescription(v string) *ItemListCommentResponseDataExtra {
	s.Description = &v
	return s
}

type ItemReplyCommentRequest struct {
	OsVersion      *string            `json:"os_version,omitempty" xml:"os_version,omitempty"`
	Content        *string            `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	ItemId         *string            `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	OpenId         *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DevicePlatform *string            `json:"device_platform,omitempty" xml:"device_platform,omitempty"`
	Ip             *string            `json:"ip,omitempty" xml:"ip,omitempty"`
	SharkChannel   *string            `json:"shark_channel,omitempty" xml:"shark_channel,omitempty"`
	CommentId      *string            `json:"comment_id,omitempty" xml:"comment_id,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DeviceBrand    *string            `json:"device_brand,omitempty" xml:"device_brand,omitempty"`
	DeviceType     *string            `json:"device_type,omitempty" xml:"device_type,omitempty"`
}

func (s ItemReplyCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s ItemReplyCommentRequest) GoString() string {
	return s.String()
}

func (s *ItemReplyCommentRequest) SetOsVersion(v string) *ItemReplyCommentRequest {
	s.OsVersion = &v
	return s
}

func (s *ItemReplyCommentRequest) SetContent(v string) *ItemReplyCommentRequest {
	s.Content = &v
	return s
}

func (s *ItemReplyCommentRequest) SetItemId(v string) *ItemReplyCommentRequest {
	s.ItemId = &v
	return s
}

func (s *ItemReplyCommentRequest) SetHeader(v map[string]*string) *ItemReplyCommentRequest {
	s.Header = v
	return s
}

func (s *ItemReplyCommentRequest) SetOpenId(v string) *ItemReplyCommentRequest {
	s.OpenId = &v
	return s
}

func (s *ItemReplyCommentRequest) SetDevicePlatform(v string) *ItemReplyCommentRequest {
	s.DevicePlatform = &v
	return s
}

func (s *ItemReplyCommentRequest) SetIp(v string) *ItemReplyCommentRequest {
	s.Ip = &v
	return s
}

func (s *ItemReplyCommentRequest) SetSharkChannel(v string) *ItemReplyCommentRequest {
	s.SharkChannel = &v
	return s
}

func (s *ItemReplyCommentRequest) SetCommentId(v string) *ItemReplyCommentRequest {
	s.CommentId = &v
	return s
}

func (s *ItemReplyCommentRequest) SetAccessToken(v string) *ItemReplyCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *ItemReplyCommentRequest) SetDeviceBrand(v string) *ItemReplyCommentRequest {
	s.DeviceBrand = &v
	return s
}

func (s *ItemReplyCommentRequest) SetDeviceType(v string) *ItemReplyCommentRequest {
	s.DeviceType = &v
	return s
}
