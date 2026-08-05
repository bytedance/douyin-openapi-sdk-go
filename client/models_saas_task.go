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

type SaasQueryWithdrawOrderResponse struct {
	Data   *SaasQueryWithdrawOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SaasQueryWithdrawOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryWithdrawOrderResponse) GoString() string {
	return s.String()
}

func (s *SaasQueryWithdrawOrderResponse) SetData(v *SaasQueryWithdrawOrderResponseData) *SaasQueryWithdrawOrderResponse {
	s.Data = v
	return s
}

func (s *SaasQueryWithdrawOrderResponse) SetErrNo(v int32) *SaasQueryWithdrawOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasQueryWithdrawOrderResponse) SetErrMsg(v string) *SaasQueryWithdrawOrderResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasQueryWithdrawOrderResponse) SetLogId(v string) *SaasQueryWithdrawOrderResponse {
	s.LogId = &v
	return s
}

type SaasQueryWithdrawOrderResponseData struct {
	Status    *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	StatusMsg *string `json:"status_msg,omitempty" xml:"status_msg,omitempty" require:"true"`
}

func (s SaasQueryWithdrawOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryWithdrawOrderResponseData) GoString() string {
	return s.String()
}

func (s *SaasQueryWithdrawOrderResponseData) SetStatus(v string) *SaasQueryWithdrawOrderResponseData {
	s.Status = &v
	return s
}

func (s *SaasQueryWithdrawOrderResponseData) SetStatusMsg(v string) *SaasQueryWithdrawOrderResponseData {
	s.StatusMsg = &v
	return s
}

type SaveLiveOrientedPlanRequest struct {
	DouyinIdList  []*string                                     `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" type:"Repeated"`
	MerchantPhone *string                                       `json:"merchant_phone,omitempty" xml:"merchant_phone,omitempty"`
	Header        map[string]*string                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PlanId        *int64                                        `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	PlanName      *string                                       `json:"plan_name,omitempty" xml:"plan_name,omitempty"`
	ProductList   []*SaveLiveOrientedPlanRequestProductListItem `json:"product_list,omitempty" xml:"product_list,omitempty" type:"Repeated"`
}

func (s SaveLiveOrientedPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveLiveOrientedPlanRequest) GoString() string {
	return s.String()
}

func (s *SaveLiveOrientedPlanRequest) SetDouyinIdList(v []*string) *SaveLiveOrientedPlanRequest {
	s.DouyinIdList = v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetMerchantPhone(v string) *SaveLiveOrientedPlanRequest {
	s.MerchantPhone = &v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetHeader(v map[string]*string) *SaveLiveOrientedPlanRequest {
	s.Header = v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetAccessToken(v string) *SaveLiveOrientedPlanRequest {
	s.AccessToken = &v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetPlanId(v int64) *SaveLiveOrientedPlanRequest {
	s.PlanId = &v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetPlanName(v string) *SaveLiveOrientedPlanRequest {
	s.PlanName = &v
	return s
}

func (s *SaveLiveOrientedPlanRequest) SetProductList(v []*SaveLiveOrientedPlanRequestProductListItem) *SaveLiveOrientedPlanRequest {
	s.ProductList = v
	return s
}

type SaveLiveOrientedPlanRequestProductListItem struct {
	ProductId      *int64 `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	CommissionRate *int64 `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
}

func (s SaveLiveOrientedPlanRequestProductListItem) String() string {
	return tea.Prettify(s)
}

func (s SaveLiveOrientedPlanRequestProductListItem) GoString() string {
	return s.String()
}

func (s *SaveLiveOrientedPlanRequestProductListItem) SetProductId(v int64) *SaveLiveOrientedPlanRequestProductListItem {
	s.ProductId = &v
	return s
}

func (s *SaveLiveOrientedPlanRequestProductListItem) SetCommissionRate(v int64) *SaveLiveOrientedPlanRequestProductListItem {
	s.CommissionRate = &v
	return s
}

type SaveLiveOrientedPlanResponse struct {
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaveLiveOrientedPlanResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SaveLiveOrientedPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveLiveOrientedPlanResponse) GoString() string {
	return s.String()
}

func (s *SaveLiveOrientedPlanResponse) SetLogId(v string) *SaveLiveOrientedPlanResponse {
	s.LogId = &v
	return s
}

func (s *SaveLiveOrientedPlanResponse) SetData(v *SaveLiveOrientedPlanResponseData) *SaveLiveOrientedPlanResponse {
	s.Data = v
	return s
}

func (s *SaveLiveOrientedPlanResponse) SetErrMsg(v string) *SaveLiveOrientedPlanResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaveLiveOrientedPlanResponse) SetErrNo(v int32) *SaveLiveOrientedPlanResponse {
	s.ErrNo = &v
	return s
}

type SaveLiveOrientedPlanResponseData struct {
	PlanId *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s SaveLiveOrientedPlanResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaveLiveOrientedPlanResponseData) GoString() string {
	return s.String()
}

func (s *SaveLiveOrientedPlanResponseData) SetPlanId(v int64) *SaveLiveOrientedPlanResponseData {
	s.PlanId = &v
	return s
}

type SavePresaleAriRequest struct {
	ProductId          *string                                  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	RoomAddPriceRule   *SavePresaleAriRequestRoomAddPriceRule   `json:"room_add_price_rule,omitempty" xml:"room_add_price_rule,omitempty"`
	MarketingAmount    *int64                                   `json:"marketing_amount,omitempty" xml:"marketing_amount,omitempty"`
	ChildrenDefinition *SavePresaleAriRequestChildrenDefinition `json:"children_definition,omitempty" xml:"children_definition,omitempty"`
	Header             map[string]*string                       `json:"header,omitempty" xml:"header,omitempty"`
	StockQtyLimitType  *int32                                   `json:"stock_qty_limit_type,omitempty" xml:"stock_qty_limit_type,omitempty"`
	ActualAmount       *int64                                   `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount       *int64                                   `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	UserAddPriceRule   *SavePresaleAriRequestUserAddPriceRule   `json:"user_add_price_rule,omitempty" xml:"user_add_price_rule,omitempty"`
	AccessToken        *string                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DateAddPriceRule   *SavePresaleAriRequestDateAddPriceRule   `json:"date_add_price_rule,omitempty" xml:"date_add_price_rule,omitempty"`
	AccountId          *string                                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	TotalStockQty      *int64                                   `json:"total_stock_qty,omitempty" xml:"total_stock_qty,omitempty"`
}

func (s SavePresaleAriRequest) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequest) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequest) SetProductId(v string) *SavePresaleAriRequest {
	s.ProductId = &v
	return s
}

func (s *SavePresaleAriRequest) SetRoomAddPriceRule(v *SavePresaleAriRequestRoomAddPriceRule) *SavePresaleAriRequest {
	s.RoomAddPriceRule = v
	return s
}

func (s *SavePresaleAriRequest) SetMarketingAmount(v int64) *SavePresaleAriRequest {
	s.MarketingAmount = &v
	return s
}

func (s *SavePresaleAriRequest) SetChildrenDefinition(v *SavePresaleAriRequestChildrenDefinition) *SavePresaleAriRequest {
	s.ChildrenDefinition = v
	return s
}

func (s *SavePresaleAriRequest) SetHeader(v map[string]*string) *SavePresaleAriRequest {
	s.Header = v
	return s
}

func (s *SavePresaleAriRequest) SetStockQtyLimitType(v int32) *SavePresaleAriRequest {
	s.StockQtyLimitType = &v
	return s
}

func (s *SavePresaleAriRequest) SetActualAmount(v int64) *SavePresaleAriRequest {
	s.ActualAmount = &v
	return s
}

func (s *SavePresaleAriRequest) SetOriginAmount(v int64) *SavePresaleAriRequest {
	s.OriginAmount = &v
	return s
}

func (s *SavePresaleAriRequest) SetUserAddPriceRule(v *SavePresaleAriRequestUserAddPriceRule) *SavePresaleAriRequest {
	s.UserAddPriceRule = v
	return s
}

func (s *SavePresaleAriRequest) SetAccessToken(v string) *SavePresaleAriRequest {
	s.AccessToken = &v
	return s
}

func (s *SavePresaleAriRequest) SetDateAddPriceRule(v *SavePresaleAriRequestDateAddPriceRule) *SavePresaleAriRequest {
	s.DateAddPriceRule = v
	return s
}

func (s *SavePresaleAriRequest) SetAccountId(v string) *SavePresaleAriRequest {
	s.AccountId = &v
	return s
}

func (s *SavePresaleAriRequest) SetTotalStockQty(v int64) *SavePresaleAriRequest {
	s.TotalStockQty = &v
	return s
}

type SavePresaleAriRequestChildrenDefinition struct {
	MaxAge    *int32 `json:"max_age,omitempty" xml:"max_age,omitempty"`
	MaxHeight *int32 `json:"max_height,omitempty" xml:"max_height,omitempty"`
	MinAge    *int32 `json:"min_age,omitempty" xml:"min_age,omitempty"`
	MinHeight *int32 `json:"min_height,omitempty" xml:"min_height,omitempty"`
}

func (s SavePresaleAriRequestChildrenDefinition) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestChildrenDefinition) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestChildrenDefinition) SetMaxAge(v int32) *SavePresaleAriRequestChildrenDefinition {
	s.MaxAge = &v
	return s
}

func (s *SavePresaleAriRequestChildrenDefinition) SetMaxHeight(v int32) *SavePresaleAriRequestChildrenDefinition {
	s.MaxHeight = &v
	return s
}

func (s *SavePresaleAriRequestChildrenDefinition) SetMinAge(v int32) *SavePresaleAriRequestChildrenDefinition {
	s.MinAge = &v
	return s
}

func (s *SavePresaleAriRequestChildrenDefinition) SetMinHeight(v int32) *SavePresaleAriRequestChildrenDefinition {
	s.MinHeight = &v
	return s
}

type SavePresaleAriRequestDateAddPriceRule struct {
	Enable           *bool                                                        `json:"enable,omitempty" xml:"enable,omitempty"`
	HotDayMarkupType *int                                                         `json:"hot_day_markup_type,omitempty" xml:"hot_day_markup_type,omitempty"`
	AddPriceRuleList []*SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
}

func (s SavePresaleAriRequestDateAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestDateAddPriceRule) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestDateAddPriceRule) SetEnable(v bool) *SavePresaleAriRequestDateAddPriceRule {
	s.Enable = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRule) SetHotDayMarkupType(v int) *SavePresaleAriRequestDateAddPriceRule {
	s.HotDayMarkupType = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRule) SetAddPriceRuleList(v []*SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) *SavePresaleAriRequestDateAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

type SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem struct {
	DaysOfWeek []*int32                                                                  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	EndDate    *string                                                                   `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Price      *int64                                                                    `json:"price,omitempty" xml:"price,omitempty"`
	AdultPrice *int64                                                                    `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	DatesList  []*SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem `json:"dates_list,omitempty" xml:"dates_list,omitempty" type:"Repeated"`
	ChildPrice *int64                                                                    `json:"child_price,omitempty" xml:"child_price,omitempty"`
	Holidays   []*int32                                                                  `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	StartDate  *string                                                                   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	DateType   *int32                                                                    `json:"date_type,omitempty" xml:"date_type,omitempty"`
}

func (s SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetDatesList(v []*SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.DatesList = v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

type SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem struct {
	DateType   *int    `json:"date_type,omitempty" xml:"date_type,omitempty"`
	DaysOfWeek []*int  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	EndDate    *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate  *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}

func (s SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) SetDateType(v int) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DateType = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) SetDaysOfWeek(v []*int) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DaysOfWeek = v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) SetEndDate(v string) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem) SetStartDate(v string) *SavePresaleAriRequestDateAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.StartDate = &v
	return s
}

type SavePresaleAriRequestRoomAddPriceRule struct {
	AddPriceRuleList []*SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
	Enable           *bool                                                        `json:"enable,omitempty" xml:"enable,omitempty"`
	HotDayMarkupType *int                                                         `json:"hot_day_markup_type,omitempty" xml:"hot_day_markup_type,omitempty"`
}

func (s SavePresaleAriRequestRoomAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestRoomAddPriceRule) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestRoomAddPriceRule) SetAddPriceRuleList(v []*SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) *SavePresaleAriRequestRoomAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRule) SetEnable(v bool) *SavePresaleAriRequestRoomAddPriceRule {
	s.Enable = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRule) SetHotDayMarkupType(v int) *SavePresaleAriRequestRoomAddPriceRule {
	s.HotDayMarkupType = &v
	return s
}

type SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem struct {
	Holidays   []*int32                                                                  `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	Price      *int64                                                                    `json:"price,omitempty" xml:"price,omitempty"`
	DateType   *int32                                                                    `json:"date_type,omitempty" xml:"date_type,omitempty"`
	EndDate    *string                                                                   `json:"end_date,omitempty" xml:"end_date,omitempty"`
	AdultPrice *int64                                                                    `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	ChildPrice *int64                                                                    `json:"child_price,omitempty" xml:"child_price,omitempty"`
	StartDate  *string                                                                   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	DatesList  []*SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem `json:"dates_list,omitempty" xml:"dates_list,omitempty" type:"Repeated"`
	DaysOfWeek []*int32                                                                  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
}

func (s SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetDatesList(v []*SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.DatesList = v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

type SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem struct {
	DaysOfWeek []*int  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	EndDate    *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate  *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	DateType   *int    `json:"date_type,omitempty" xml:"date_type,omitempty"`
}

func (s SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) SetDaysOfWeek(v []*int) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DaysOfWeek = v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) SetEndDate(v string) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) SetStartDate(v string) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.StartDate = &v
	return s
}

func (s *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem) SetDateType(v int) *SavePresaleAriRequestRoomAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DateType = &v
	return s
}

type SavePresaleAriRequestUserAddPriceRule struct {
	HotDayMarkupType *int                                                         `json:"hot_day_markup_type,omitempty" xml:"hot_day_markup_type,omitempty"`
	AddPriceRuleList []*SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
	Enable           *bool                                                        `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s SavePresaleAriRequestUserAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestUserAddPriceRule) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestUserAddPriceRule) SetHotDayMarkupType(v int) *SavePresaleAriRequestUserAddPriceRule {
	s.HotDayMarkupType = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRule) SetAddPriceRuleList(v []*SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) *SavePresaleAriRequestUserAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRule) SetEnable(v bool) *SavePresaleAriRequestUserAddPriceRule {
	s.Enable = &v
	return s
}

type SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem struct {
	DateType   *int32                                                                    `json:"date_type,omitempty" xml:"date_type,omitempty"`
	AdultPrice *int64                                                                    `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	EndDate    *string                                                                   `json:"end_date,omitempty" xml:"end_date,omitempty"`
	DatesList  []*SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem `json:"dates_list,omitempty" xml:"dates_list,omitempty" type:"Repeated"`
	ChildPrice *int64                                                                    `json:"child_price,omitempty" xml:"child_price,omitempty"`
	DaysOfWeek []*int32                                                                  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	Holidays   []*int32                                                                  `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	StartDate  *string                                                                   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	Price      *int64                                                                    `json:"price,omitempty" xml:"price,omitempty"`
}

func (s SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetDatesList(v []*SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.DatesList = v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

type SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem struct {
	EndDate    *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate  *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	DateType   *int    `json:"date_type,omitempty" xml:"date_type,omitempty"`
	DaysOfWeek []*int  `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
}

func (s SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) GoString() string {
	return s.String()
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) SetEndDate(v string) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.EndDate = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) SetStartDate(v string) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.StartDate = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) SetDateType(v int) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DateType = &v
	return s
}

func (s *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem) SetDaysOfWeek(v []*int) *SavePresaleAriRequestUserAddPriceRuleAddPriceRuleListItemDatesListItem {
	s.DaysOfWeek = v
	return s
}

type SavePresaleAriResponse struct {
	Data  *SavePresaleAriResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SavePresaleAriResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s SavePresaleAriResponse) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriResponse) GoString() string {
	return s.String()
}

func (s *SavePresaleAriResponse) SetData(v *SavePresaleAriResponseData) *SavePresaleAriResponse {
	s.Data = v
	return s
}

func (s *SavePresaleAriResponse) SetExtra(v *SavePresaleAriResponseExtra) *SavePresaleAriResponse {
	s.Extra = v
	return s
}

type SavePresaleAriResponseData struct {
	Descripiton   *string `json:"descripiton,omitempty" xml:"descripiton,omitempty"`
	ErrorCode     *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SavePresaleAriResponseData) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriResponseData) GoString() string {
	return s.String()
}

func (s *SavePresaleAriResponseData) SetDescripiton(v string) *SavePresaleAriResponseData {
	s.Descripiton = &v
	return s
}

func (s *SavePresaleAriResponseData) SetErrorCode(v string) *SavePresaleAriResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SavePresaleAriResponseData) SetGwDescription(v string) *SavePresaleAriResponseData {
	s.GwDescription = &v
	return s
}

type SavePresaleAriResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s SavePresaleAriResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SavePresaleAriResponseExtra) GoString() string {
	return s.String()
}

func (s *SavePresaleAriResponseExtra) SetSubDescription(v string) *SavePresaleAriResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SavePresaleAriResponseExtra) SetSubErrorCode(v int32) *SavePresaleAriResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SavePresaleAriResponseExtra) SetDescription(v string) *SavePresaleAriResponseExtra {
	s.Description = &v
	return s
}

func (s *SavePresaleAriResponseExtra) SetErrorCode(v int32) *SavePresaleAriResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SavePresaleAriResponseExtra) SetLogid(v string) *SavePresaleAriResponseExtra {
	s.Logid = &v
	return s
}

func (s *SavePresaleAriResponseExtra) SetNow(v int64) *SavePresaleAriResponseExtra {
	s.Now = &v
	return s
}

type SaveRetainConsultCardRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Components  []*int             `json:"components,omitempty" xml:"components,omitempty" require:"true" type:"Repeated"`
	MediaId     *string            `json:"media_id,omitempty" xml:"media_id,omitempty" require:"true"`
	Title       *string            `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	CardId      *string            `json:"card_id,omitempty" xml:"card_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SaveRetainConsultCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveRetainConsultCardRequest) GoString() string {
	return s.String()
}

func (s *SaveRetainConsultCardRequest) SetAccessToken(v string) *SaveRetainConsultCardRequest {
	s.AccessToken = &v
	return s
}

func (s *SaveRetainConsultCardRequest) SetOpenId(v string) *SaveRetainConsultCardRequest {
	s.OpenId = &v
	return s
}

func (s *SaveRetainConsultCardRequest) SetComponents(v []*int) *SaveRetainConsultCardRequest {
	s.Components = v
	return s
}

func (s *SaveRetainConsultCardRequest) SetMediaId(v string) *SaveRetainConsultCardRequest {
	s.MediaId = &v
	return s
}

func (s *SaveRetainConsultCardRequest) SetTitle(v string) *SaveRetainConsultCardRequest {
	s.Title = &v
	return s
}

func (s *SaveRetainConsultCardRequest) SetCardId(v string) *SaveRetainConsultCardRequest {
	s.CardId = &v
	return s
}

func (s *SaveRetainConsultCardRequest) SetHeader(v map[string]*string) *SaveRetainConsultCardRequest {
	s.Header = v
	return s
}

type SaveRetainConsultCardResponse struct {
	CardId *string                             `json:"card_id,omitempty" xml:"card_id,omitempty"`
	Data   *SaveRetainConsultCardResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra  *SaveRetainConsultCardResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SaveRetainConsultCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveRetainConsultCardResponse) GoString() string {
	return s.String()
}

func (s *SaveRetainConsultCardResponse) SetCardId(v string) *SaveRetainConsultCardResponse {
	s.CardId = &v
	return s
}

func (s *SaveRetainConsultCardResponse) SetData(v *SaveRetainConsultCardResponseData) *SaveRetainConsultCardResponse {
	s.Data = v
	return s
}

func (s *SaveRetainConsultCardResponse) SetExtra(v *SaveRetainConsultCardResponseExtra) *SaveRetainConsultCardResponse {
	s.Extra = v
	return s
}

type SaveRetainConsultCardResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SaveRetainConsultCardResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaveRetainConsultCardResponseData) GoString() string {
	return s.String()
}

func (s *SaveRetainConsultCardResponseData) SetGwErrorCode(v int32) *SaveRetainConsultCardResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SaveRetainConsultCardResponseData) SetGwDescription(v string) *SaveRetainConsultCardResponseData {
	s.GwDescription = &v
	return s
}

type SaveRetainConsultCardResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s SaveRetainConsultCardResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SaveRetainConsultCardResponseExtra) GoString() string {
	return s.String()
}

func (s *SaveRetainConsultCardResponseExtra) SetNow(v int64) *SaveRetainConsultCardResponseExtra {
	s.Now = &v
	return s
}

func (s *SaveRetainConsultCardResponseExtra) SetSubDescription(v string) *SaveRetainConsultCardResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SaveRetainConsultCardResponseExtra) SetSubErrorCode(v int32) *SaveRetainConsultCardResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SaveRetainConsultCardResponseExtra) SetDescription(v string) *SaveRetainConsultCardResponseExtra {
	s.Description = &v
	return s
}

func (s *SaveRetainConsultCardResponseExtra) SetErrorCode(v int32) *SaveRetainConsultCardResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SaveRetainConsultCardResponseExtra) SetLogid(v string) *SaveRetainConsultCardResponseExtra {
	s.Logid = &v
	return s
}

type SaveVideoOrientedPlanRequest struct {
	PlanName           *string                                        `json:"plan_name,omitempty" xml:"plan_name,omitempty"`
	ProductList        []*SaveVideoOrientedPlanRequestProductListItem `json:"product_list,omitempty" xml:"product_list,omitempty" type:"Repeated"`
	CommissionDuration *int64                                         `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	DouyinIdList       []*string                                      `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" type:"Repeated"`
	EndTime            *int64                                         `json:"end_time,omitempty" xml:"end_time,omitempty"`
	PlanId             *int64                                         `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	StartTime          *int64                                         `json:"start_time,omitempty" xml:"start_time,omitempty"`
	MerchantPhone      *string                                        `json:"merchant_phone,omitempty" xml:"merchant_phone,omitempty"`
	Header             map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SaveVideoOrientedPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoOrientedPlanRequest) GoString() string {
	return s.String()
}

func (s *SaveVideoOrientedPlanRequest) SetPlanName(v string) *SaveVideoOrientedPlanRequest {
	s.PlanName = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetProductList(v []*SaveVideoOrientedPlanRequestProductListItem) *SaveVideoOrientedPlanRequest {
	s.ProductList = v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetCommissionDuration(v int64) *SaveVideoOrientedPlanRequest {
	s.CommissionDuration = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetDouyinIdList(v []*string) *SaveVideoOrientedPlanRequest {
	s.DouyinIdList = v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetEndTime(v int64) *SaveVideoOrientedPlanRequest {
	s.EndTime = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetPlanId(v int64) *SaveVideoOrientedPlanRequest {
	s.PlanId = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetStartTime(v int64) *SaveVideoOrientedPlanRequest {
	s.StartTime = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetMerchantPhone(v string) *SaveVideoOrientedPlanRequest {
	s.MerchantPhone = &v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetHeader(v map[string]*string) *SaveVideoOrientedPlanRequest {
	s.Header = v
	return s
}

func (s *SaveVideoOrientedPlanRequest) SetAccessToken(v string) *SaveVideoOrientedPlanRequest {
	s.AccessToken = &v
	return s
}

type SaveVideoOrientedPlanRequestProductListItem struct {
	CommissionRate *int64 `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	ProductId      *int64 `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s SaveVideoOrientedPlanRequestProductListItem) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoOrientedPlanRequestProductListItem) GoString() string {
	return s.String()
}

func (s *SaveVideoOrientedPlanRequestProductListItem) SetCommissionRate(v int64) *SaveVideoOrientedPlanRequestProductListItem {
	s.CommissionRate = &v
	return s
}

func (s *SaveVideoOrientedPlanRequestProductListItem) SetProductId(v int64) *SaveVideoOrientedPlanRequestProductListItem {
	s.ProductId = &v
	return s
}

type SaveVideoOrientedPlanResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaveVideoOrientedPlanResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SaveVideoOrientedPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoOrientedPlanResponse) GoString() string {
	return s.String()
}

func (s *SaveVideoOrientedPlanResponse) SetErrMsg(v string) *SaveVideoOrientedPlanResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaveVideoOrientedPlanResponse) SetErrNo(v int32) *SaveVideoOrientedPlanResponse {
	s.ErrNo = &v
	return s
}

func (s *SaveVideoOrientedPlanResponse) SetLogId(v string) *SaveVideoOrientedPlanResponse {
	s.LogId = &v
	return s
}

func (s *SaveVideoOrientedPlanResponse) SetData(v *SaveVideoOrientedPlanResponseData) *SaveVideoOrientedPlanResponse {
	s.Data = v
	return s
}

type SaveVideoOrientedPlanResponseData struct {
	PlanId *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s SaveVideoOrientedPlanResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaveVideoOrientedPlanResponseData) GoString() string {
	return s.String()
}

func (s *SaveVideoOrientedPlanResponseData) SetPlanId(v int64) *SaveVideoOrientedPlanResponseData {
	s.PlanId = &v
	return s
}

type SchemaGenerateRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Query       *string            `json:"query,omitempty" xml:"query,omitempty"`
	ExpireTime  *int64             `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	NoExpire    *bool              `json:"no_expire,omitempty" xml:"no_expire,omitempty" require:"true"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
}

func (s SchemaGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s SchemaGenerateRequest) GoString() string {
	return s.String()
}

func (s *SchemaGenerateRequest) SetHeader(v map[string]*string) *SchemaGenerateRequest {
	s.Header = v
	return s
}

func (s *SchemaGenerateRequest) SetAccessToken(v string) *SchemaGenerateRequest {
	s.AccessToken = &v
	return s
}

func (s *SchemaGenerateRequest) SetQuery(v string) *SchemaGenerateRequest {
	s.Query = &v
	return s
}

func (s *SchemaGenerateRequest) SetExpireTime(v int64) *SchemaGenerateRequest {
	s.ExpireTime = &v
	return s
}

func (s *SchemaGenerateRequest) SetNoExpire(v bool) *SchemaGenerateRequest {
	s.NoExpire = &v
	return s
}

func (s *SchemaGenerateRequest) SetVersionType(v string) *SchemaGenerateRequest {
	s.VersionType = &v
	return s
}

type SchemaGenerateResponse struct {
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SchemaGenerateResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SchemaGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s SchemaGenerateResponse) GoString() string {
	return s.String()
}

func (s *SchemaGenerateResponse) SetErrMsg(v string) *SchemaGenerateResponse {
	s.ErrMsg = &v
	return s
}

func (s *SchemaGenerateResponse) SetLogId(v string) *SchemaGenerateResponse {
	s.LogId = &v
	return s
}

func (s *SchemaGenerateResponse) SetData(v *SchemaGenerateResponseData) *SchemaGenerateResponse {
	s.Data = v
	return s
}

func (s *SchemaGenerateResponse) SetErrNo(v int32) *SchemaGenerateResponse {
	s.ErrNo = &v
	return s
}

type SchemaGenerateResponseData struct {
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty" require:"true"`
}

func (s SchemaGenerateResponseData) String() string {
	return tea.Prettify(s)
}

func (s SchemaGenerateResponseData) GoString() string {
	return s.String()
}

func (s *SchemaGenerateResponseData) SetSchema(v string) *SchemaGenerateResponseData {
	s.Schema = &v
	return s
}

type SchemaGetItemInfoRequest struct {
	VideoId     *int64             `json:"video_id,omitempty" xml:"video_id,omitempty"`
	ExpireAt    *int64             `json:"expire_at,omitempty" xml:"expire_at,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SchemaGetItemInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SchemaGetItemInfoRequest) GoString() string {
	return s.String()
}

func (s *SchemaGetItemInfoRequest) SetVideoId(v int64) *SchemaGetItemInfoRequest {
	s.VideoId = &v
	return s
}

func (s *SchemaGetItemInfoRequest) SetExpireAt(v int64) *SchemaGetItemInfoRequest {
	s.ExpireAt = &v
	return s
}

func (s *SchemaGetItemInfoRequest) SetItemId(v string) *SchemaGetItemInfoRequest {
	s.ItemId = &v
	return s
}

func (s *SchemaGetItemInfoRequest) SetHeader(v map[string]*string) *SchemaGetItemInfoRequest {
	s.Header = v
	return s
}

func (s *SchemaGetItemInfoRequest) SetAccessToken(v string) *SchemaGetItemInfoRequest {
	s.AccessToken = &v
	return s
}

type SchemaGetItemInfoResponse struct {
	Data   *SchemaGetItemInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int64                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SchemaGetItemInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SchemaGetItemInfoResponse) GoString() string {
	return s.String()
}

func (s *SchemaGetItemInfoResponse) SetData(v *SchemaGetItemInfoResponseData) *SchemaGetItemInfoResponse {
	s.Data = v
	return s
}

func (s *SchemaGetItemInfoResponse) SetErrMsg(v string) *SchemaGetItemInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *SchemaGetItemInfoResponse) SetErrNo(v int64) *SchemaGetItemInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *SchemaGetItemInfoResponse) SetLogId(v string) *SchemaGetItemInfoResponse {
	s.LogId = &v
	return s
}

type SchemaGetItemInfoResponseData struct {
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s SchemaGetItemInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s SchemaGetItemInfoResponseData) GoString() string {
	return s.String()
}

func (s *SchemaGetItemInfoResponseData) SetSchema(v string) *SchemaGetItemInfoResponseData {
	s.Schema = &v
	return s
}

type SchemaQueryInfoRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Schema      *string            `json:"schema,omitempty" xml:"schema,omitempty" require:"true"`
}

func (s SchemaQueryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryInfoRequest) GoString() string {
	return s.String()
}

func (s *SchemaQueryInfoRequest) SetHeader(v map[string]*string) *SchemaQueryInfoRequest {
	s.Header = v
	return s
}

func (s *SchemaQueryInfoRequest) SetAccessToken(v string) *SchemaQueryInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *SchemaQueryInfoRequest) SetSchema(v string) *SchemaQueryInfoRequest {
	s.Schema = &v
	return s
}

type SchemaQueryInfoResponse struct {
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SchemaQueryInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SchemaQueryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryInfoResponse) GoString() string {
	return s.String()
}

func (s *SchemaQueryInfoResponse) SetErrMsg(v string) *SchemaQueryInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *SchemaQueryInfoResponse) SetLogId(v string) *SchemaQueryInfoResponse {
	s.LogId = &v
	return s
}

func (s *SchemaQueryInfoResponse) SetData(v *SchemaQueryInfoResponseData) *SchemaQueryInfoResponse {
	s.Data = v
	return s
}

func (s *SchemaQueryInfoResponse) SetErrNo(v int32) *SchemaQueryInfoResponse {
	s.ErrNo = &v
	return s
}

type SchemaQueryInfoResponseData struct {
	Query       *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	CreateTime  *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	ExpireTime  *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	VersionType *string `json:"version_type,omitempty" xml:"version_type,omitempty" require:"true"`
	AppId       *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Path        *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s SchemaQueryInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryInfoResponseData) GoString() string {
	return s.String()
}

func (s *SchemaQueryInfoResponseData) SetQuery(v string) *SchemaQueryInfoResponseData {
	s.Query = &v
	return s
}

func (s *SchemaQueryInfoResponseData) SetCreateTime(v int64) *SchemaQueryInfoResponseData {
	s.CreateTime = &v
	return s
}

func (s *SchemaQueryInfoResponseData) SetExpireTime(v int64) *SchemaQueryInfoResponseData {
	s.ExpireTime = &v
	return s
}

func (s *SchemaQueryInfoResponseData) SetVersionType(v string) *SchemaQueryInfoResponseData {
	s.VersionType = &v
	return s
}

func (s *SchemaQueryInfoResponseData) SetAppId(v string) *SchemaQueryInfoResponseData {
	s.AppId = &v
	return s
}

func (s *SchemaQueryInfoResponseData) SetPath(v string) *SchemaQueryInfoResponseData {
	s.Path = &v
	return s
}

type SchemaQueryQuotaRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SchemaQueryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryQuotaRequest) GoString() string {
	return s.String()
}

func (s *SchemaQueryQuotaRequest) SetHeader(v map[string]*string) *SchemaQueryQuotaRequest {
	s.Header = v
	return s
}

func (s *SchemaQueryQuotaRequest) SetAccessToken(v string) *SchemaQueryQuotaRequest {
	s.AccessToken = &v
	return s
}

type SchemaQueryQuotaResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SchemaQueryQuotaResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s SchemaQueryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryQuotaResponse) GoString() string {
	return s.String()
}

func (s *SchemaQueryQuotaResponse) SetLogId(v string) *SchemaQueryQuotaResponse {
	s.LogId = &v
	return s
}

func (s *SchemaQueryQuotaResponse) SetData(v *SchemaQueryQuotaResponseData) *SchemaQueryQuotaResponse {
	s.Data = v
	return s
}

func (s *SchemaQueryQuotaResponse) SetErrNo(v int32) *SchemaQueryQuotaResponse {
	s.ErrNo = &v
	return s
}

func (s *SchemaQueryQuotaResponse) SetErrMsg(v string) *SchemaQueryQuotaResponse {
	s.ErrMsg = &v
	return s
}

type SchemaQueryQuotaResponseData struct {
	ShortTermSchemaQuota *SchemaQueryQuotaResponseDataShortTermSchemaQuota `json:"short_term_schema_quota,omitempty" xml:"short_term_schema_quota,omitempty" require:"true"`
	LongTermSchemaQuota  *SchemaQueryQuotaResponseDataLongTermSchemaQuota  `json:"long_term_schema_quota,omitempty" xml:"long_term_schema_quota,omitempty" require:"true"`
}

func (s SchemaQueryQuotaResponseData) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryQuotaResponseData) GoString() string {
	return s.String()
}

func (s *SchemaQueryQuotaResponseData) SetShortTermSchemaQuota(v *SchemaQueryQuotaResponseDataShortTermSchemaQuota) *SchemaQueryQuotaResponseData {
	s.ShortTermSchemaQuota = v
	return s
}

func (s *SchemaQueryQuotaResponseData) SetLongTermSchemaQuota(v *SchemaQueryQuotaResponseDataLongTermSchemaQuota) *SchemaQueryQuotaResponseData {
	s.LongTermSchemaQuota = v
	return s
}

type SchemaQueryQuotaResponseDataLongTermSchemaQuota struct {
	SchemaLimit *int32 `json:"schema_limit,omitempty" xml:"schema_limit,omitempty" require:"true"`
	SchemaUsed  *int32 `json:"schema_used,omitempty" xml:"schema_used,omitempty" require:"true"`
}

func (s SchemaQueryQuotaResponseDataLongTermSchemaQuota) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryQuotaResponseDataLongTermSchemaQuota) GoString() string {
	return s.String()
}

func (s *SchemaQueryQuotaResponseDataLongTermSchemaQuota) SetSchemaLimit(v int32) *SchemaQueryQuotaResponseDataLongTermSchemaQuota {
	s.SchemaLimit = &v
	return s
}

func (s *SchemaQueryQuotaResponseDataLongTermSchemaQuota) SetSchemaUsed(v int32) *SchemaQueryQuotaResponseDataLongTermSchemaQuota {
	s.SchemaUsed = &v
	return s
}

type SchemaQueryQuotaResponseDataShortTermSchemaQuota struct {
	SchemaLimit *int32 `json:"schema_limit,omitempty" xml:"schema_limit,omitempty" require:"true"`
	SchemaUsed  *int32 `json:"schema_used,omitempty" xml:"schema_used,omitempty" require:"true"`
}

func (s SchemaQueryQuotaResponseDataShortTermSchemaQuota) String() string {
	return tea.Prettify(s)
}

func (s SchemaQueryQuotaResponseDataShortTermSchemaQuota) GoString() string {
	return s.String()
}

func (s *SchemaQueryQuotaResponseDataShortTermSchemaQuota) SetSchemaLimit(v int32) *SchemaQueryQuotaResponseDataShortTermSchemaQuota {
	s.SchemaLimit = &v
	return s
}

func (s *SchemaQueryQuotaResponseDataShortTermSchemaQuota) SetSchemaUsed(v int32) *SchemaQueryQuotaResponseDataShortTermSchemaQuota {
	s.SchemaUsed = &v
	return s
}

type ScoreQueryRequest struct {
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PoiIdList   []*int64           `json:"poi_id_list,omitempty" xml:"poi_id_list,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ScoreQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ScoreQueryRequest) GoString() string {
	return s.String()
}

func (s *ScoreQueryRequest) SetAccountId(v string) *ScoreQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ScoreQueryRequest) SetPoiIdList(v []*int64) *ScoreQueryRequest {
	s.PoiIdList = v
	return s
}

func (s *ScoreQueryRequest) SetHeader(v map[string]*string) *ScoreQueryRequest {
	s.Header = v
	return s
}

func (s *ScoreQueryRequest) SetAccessToken(v string) *ScoreQueryRequest {
	s.AccessToken = &v
	return s
}

type ScoreQueryResponse struct {
	Data  *ScoreQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ScoreQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ScoreQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ScoreQueryResponse) GoString() string {
	return s.String()
}

func (s *ScoreQueryResponse) SetData(v *ScoreQueryResponseData) *ScoreQueryResponse {
	s.Data = v
	return s
}

func (s *ScoreQueryResponse) SetExtra(v *ScoreQueryResponseExtra) *ScoreQueryResponse {
	s.Extra = v
	return s
}

type ScoreQueryResponseData struct {
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ScoreInfoList []*ScoreQueryResponseDataScoreInfoListItem `json:"score_info_list,omitempty" xml:"score_info_list,omitempty" type:"Repeated"`
}

func (s ScoreQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ScoreQueryResponseData) GoString() string {
	return s.String()
}

func (s *ScoreQueryResponseData) SetGwErrorCode(v int32) *ScoreQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ScoreQueryResponseData) SetGwDescription(v string) *ScoreQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *ScoreQueryResponseData) SetScoreInfoList(v []*ScoreQueryResponseDataScoreInfoListItem) *ScoreQueryResponseData {
	s.ScoreInfoList = v
	return s
}

type ScoreQueryResponseDataScoreInfoListItem struct {
	PoiId          *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	QueryTime      *int64  `json:"query_time,omitempty" xml:"query_time,omitempty"`
	Score          *string `json:"score,omitempty" xml:"score,omitempty"`
	YesterdayScore *string `json:"yesterday_score,omitempty" xml:"yesterday_score,omitempty"`
}

func (s ScoreQueryResponseDataScoreInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ScoreQueryResponseDataScoreInfoListItem) GoString() string {
	return s.String()
}

func (s *ScoreQueryResponseDataScoreInfoListItem) SetPoiId(v int64) *ScoreQueryResponseDataScoreInfoListItem {
	s.PoiId = &v
	return s
}

func (s *ScoreQueryResponseDataScoreInfoListItem) SetQueryTime(v int64) *ScoreQueryResponseDataScoreInfoListItem {
	s.QueryTime = &v
	return s
}

func (s *ScoreQueryResponseDataScoreInfoListItem) SetScore(v string) *ScoreQueryResponseDataScoreInfoListItem {
	s.Score = &v
	return s
}

func (s *ScoreQueryResponseDataScoreInfoListItem) SetYesterdayScore(v string) *ScoreQueryResponseDataScoreInfoListItem {
	s.YesterdayScore = &v
	return s
}

type ScoreQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ScoreQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ScoreQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ScoreQueryResponseExtra) SetSubDescription(v string) *ScoreQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ScoreQueryResponseExtra) SetSubErrorCode(v int32) *ScoreQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ScoreQueryResponseExtra) SetDescription(v string) *ScoreQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ScoreQueryResponseExtra) SetErrorCode(v int32) *ScoreQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ScoreQueryResponseExtra) SetLogid(v string) *ScoreQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ScoreQueryResponseExtra) SetNow(v int64) *ScoreQueryResponseExtra {
	s.Now = &v
	return s
}

type SearchCheckSubServiceRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SearchCheckSubServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchCheckSubServiceRequest) GoString() string {
	return s.String()
}

func (s *SearchCheckSubServiceRequest) SetHeader(v map[string]*string) *SearchCheckSubServiceRequest {
	s.Header = v
	return s
}

func (s *SearchCheckSubServiceRequest) SetAccessToken(v string) *SearchCheckSubServiceRequest {
	s.AccessToken = &v
	return s
}

type SearchCheckSubServiceResponse struct {
	Data   *SearchCheckSubServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SearchCheckSubServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchCheckSubServiceResponse) GoString() string {
	return s.String()
}

func (s *SearchCheckSubServiceResponse) SetData(v *SearchCheckSubServiceResponseData) *SearchCheckSubServiceResponse {
	s.Data = v
	return s
}

func (s *SearchCheckSubServiceResponse) SetErrNo(v int32) *SearchCheckSubServiceResponse {
	s.ErrNo = &v
	return s
}

func (s *SearchCheckSubServiceResponse) SetErrMsg(v string) *SearchCheckSubServiceResponse {
	s.ErrMsg = &v
	return s
}

func (s *SearchCheckSubServiceResponse) SetLogId(v string) *SearchCheckSubServiceResponse {
	s.LogId = &v
	return s
}

type SearchCheckSubServiceResponseData struct {
	HasPermission *bool `json:"has_permission,omitempty" xml:"has_permission,omitempty" require:"true"`
}

func (s SearchCheckSubServiceResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchCheckSubServiceResponseData) GoString() string {
	return s.String()
}

func (s *SearchCheckSubServiceResponseData) SetHasPermission(v bool) *SearchCheckSubServiceResponseData {
	s.HasPermission = &v
	return s
}

type SearchDeleteIndexRequest struct {
	PathList    []*string          `json:"path_list,omitempty" xml:"path_list,omitempty" require:"true" type:"Repeated"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Name        *string            `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SearchDeleteIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *SearchDeleteIndexRequest) SetPathList(v []*string) *SearchDeleteIndexRequest {
	s.PathList = v
	return s
}

func (s *SearchDeleteIndexRequest) SetAppId(v string) *SearchDeleteIndexRequest {
	s.AppId = &v
	return s
}

func (s *SearchDeleteIndexRequest) SetName(v string) *SearchDeleteIndexRequest {
	s.Name = &v
	return s
}

func (s *SearchDeleteIndexRequest) SetHeader(v map[string]*string) *SearchDeleteIndexRequest {
	s.Header = v
	return s
}

func (s *SearchDeleteIndexRequest) SetAccessToken(v string) *SearchDeleteIndexRequest {
	s.AccessToken = &v
	return s
}

type SearchDeleteIndexResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SearchDeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *SearchDeleteIndexResponse) SetErrMsg(v string) *SearchDeleteIndexResponse {
	s.ErrMsg = &v
	return s
}

func (s *SearchDeleteIndexResponse) SetLogId(v string) *SearchDeleteIndexResponse {
	s.LogId = &v
	return s
}

func (s *SearchDeleteIndexResponse) SetErrNo(v int32) *SearchDeleteIndexResponse {
	s.ErrNo = &v
	return s
}

type SearchDeleteSubServiceRequest struct {
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	SubServiceId *string            `json:"sub_service_id,omitempty" xml:"sub_service_id,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SearchDeleteSubServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDeleteSubServiceRequest) GoString() string {
	return s.String()
}

func (s *SearchDeleteSubServiceRequest) SetAccessToken(v string) *SearchDeleteSubServiceRequest {
	s.AccessToken = &v
	return s
}

func (s *SearchDeleteSubServiceRequest) SetSubServiceId(v string) *SearchDeleteSubServiceRequest {
	s.SubServiceId = &v
	return s
}

func (s *SearchDeleteSubServiceRequest) SetHeader(v map[string]*string) *SearchDeleteSubServiceRequest {
	s.Header = v
	return s
}

type SearchDeleteSubServiceResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SearchDeleteSubServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDeleteSubServiceResponse) GoString() string {
	return s.String()
}

func (s *SearchDeleteSubServiceResponse) SetErrMsg(v string) *SearchDeleteSubServiceResponse {
	s.ErrMsg = &v
	return s
}

func (s *SearchDeleteSubServiceResponse) SetLogId(v string) *SearchDeleteSubServiceResponse {
	s.LogId = &v
	return s
}

func (s *SearchDeleteSubServiceResponse) SetErrNo(v int32) *SearchDeleteSubServiceResponse {
	s.ErrNo = &v
	return s
}

type SearchExperienceRequest struct {
	SortType    *int32             `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	ContentType *int32             `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	DeviceId    *int64             `json:"device_id,omitempty" xml:"device_id,omitempty"`
	Keyword     *string            `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
	SearchId    *string            `json:"search_id,omitempty" xml:"search_id,omitempty"`
	Uid         *string            `json:"uid,omitempty" xml:"uid,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SearchExperienceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchExperienceRequest) GoString() string {
	return s.String()
}

func (s *SearchExperienceRequest) SetSortType(v int32) *SearchExperienceRequest {
	s.SortType = &v
	return s
}

func (s *SearchExperienceRequest) SetContentType(v int32) *SearchExperienceRequest {
	s.ContentType = &v
	return s
}

func (s *SearchExperienceRequest) SetCursor(v int64) *SearchExperienceRequest {
	s.Cursor = &v
	return s
}

func (s *SearchExperienceRequest) SetHeader(v map[string]*string) *SearchExperienceRequest {
	s.Header = v
	return s
}

func (s *SearchExperienceRequest) SetDeviceId(v int64) *SearchExperienceRequest {
	s.DeviceId = &v
	return s
}

func (s *SearchExperienceRequest) SetKeyword(v string) *SearchExperienceRequest {
	s.Keyword = &v
	return s
}

func (s *SearchExperienceRequest) SetSearchId(v string) *SearchExperienceRequest {
	s.SearchId = &v
	return s
}

func (s *SearchExperienceRequest) SetUid(v string) *SearchExperienceRequest {
	s.Uid = &v
	return s
}

func (s *SearchExperienceRequest) SetCount(v int32) *SearchExperienceRequest {
	s.Count = &v
	return s
}

func (s *SearchExperienceRequest) SetAccessToken(v string) *SearchExperienceRequest {
	s.AccessToken = &v
	return s
}

type SearchExperienceResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SearchExperienceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s SearchExperienceResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchExperienceResponse) GoString() string {
	return s.String()
}

func (s *SearchExperienceResponse) SetLogId(v string) *SearchExperienceResponse {
	s.LogId = &v
	return s
}

func (s *SearchExperienceResponse) SetData(v *SearchExperienceResponseData) *SearchExperienceResponse {
	s.Data = v
	return s
}

func (s *SearchExperienceResponse) SetErrNo(v int32) *SearchExperienceResponse {
	s.ErrNo = &v
	return s
}

func (s *SearchExperienceResponse) SetErrMsg(v string) *SearchExperienceResponse {
	s.ErrMsg = &v
	return s
}

type SearchExperienceResponseData struct {
	Cursor   *int64                                  `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Data     []*SearchExperienceResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	HasMore  *bool                                   `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	SearchId *string                                 `json:"search_id,omitempty" xml:"search_id,omitempty" require:"true"`
}

func (s SearchExperienceResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchExperienceResponseData) GoString() string {
	return s.String()
}

func (s *SearchExperienceResponseData) SetCursor(v int64) *SearchExperienceResponseData {
	s.Cursor = &v
	return s
}

func (s *SearchExperienceResponseData) SetData(v []*SearchExperienceResponseDataDataItem) *SearchExperienceResponseData {
	s.Data = v
	return s
}

func (s *SearchExperienceResponseData) SetHasMore(v bool) *SearchExperienceResponseData {
	s.HasMore = &v
	return s
}

func (s *SearchExperienceResponseData) SetSearchId(v string) *SearchExperienceResponseData {
	s.SearchId = &v
	return s
}

type SearchExperienceResponseDataDataItem struct {
	Nickname        *string                                    `json:"nickname,omitempty" xml:"nickname,omitempty"`
	ShareUrl        *string                                    `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title           *string                                    `json:"title,omitempty" xml:"title,omitempty"`
	ContentType     *int32                                     `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Cover           *SearchExperienceResponseDataDataItemCover `json:"cover,omitempty" xml:"cover,omitempty"`
	Duration        *int64                                     `json:"duration,omitempty" xml:"duration,omitempty"`
	HighQualityText *string                                    `json:"high_quality_text,omitempty" xml:"high_quality_text,omitempty"`
	ItemId          *string                                    `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s SearchExperienceResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s SearchExperienceResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *SearchExperienceResponseDataDataItem) SetNickname(v string) *SearchExperienceResponseDataDataItem {
	s.Nickname = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetShareUrl(v string) *SearchExperienceResponseDataDataItem {
	s.ShareUrl = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetTitle(v string) *SearchExperienceResponseDataDataItem {
	s.Title = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetContentType(v int32) *SearchExperienceResponseDataDataItem {
	s.ContentType = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetCover(v *SearchExperienceResponseDataDataItemCover) *SearchExperienceResponseDataDataItem {
	s.Cover = v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetDuration(v int64) *SearchExperienceResponseDataDataItem {
	s.Duration = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetHighQualityText(v string) *SearchExperienceResponseDataDataItem {
	s.HighQualityText = &v
	return s
}

func (s *SearchExperienceResponseDataDataItem) SetItemId(v string) *SearchExperienceResponseDataDataItem {
	s.ItemId = &v
	return s
}

type SearchExperienceResponseDataDataItemCover struct {
	UrlList []*string `json:"url_list,omitempty" xml:"url_list,omitempty" type:"Repeated"`
	Width   *int32    `json:"width,omitempty" xml:"width,omitempty"`
	Height  *int32    `json:"height,omitempty" xml:"height,omitempty"`
}

func (s SearchExperienceResponseDataDataItemCover) String() string {
	return tea.Prettify(s)
}

func (s SearchExperienceResponseDataDataItemCover) GoString() string {
	return s.String()
}

func (s *SearchExperienceResponseDataDataItemCover) SetUrlList(v []*string) *SearchExperienceResponseDataDataItemCover {
	s.UrlList = v
	return s
}

func (s *SearchExperienceResponseDataDataItemCover) SetWidth(v int32) *SearchExperienceResponseDataDataItemCover {
	s.Width = &v
	return s
}

func (s *SearchExperienceResponseDataDataItemCover) SetHeight(v int32) *SearchExperienceResponseDataDataItemCover {
	s.Height = &v
	return s
}

type SearchKeywordRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Count       *int64             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	Keyword     *string            `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
	City        *string            `json:"city,omitempty" xml:"city,omitempty"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SearchKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordRequest) GoString() string {
	return s.String()
}

func (s *SearchKeywordRequest) SetAccessToken(v string) *SearchKeywordRequest {
	s.AccessToken = &v
	return s
}

func (s *SearchKeywordRequest) SetCount(v int64) *SearchKeywordRequest {
	s.Count = &v
	return s
}

func (s *SearchKeywordRequest) SetKeyword(v string) *SearchKeywordRequest {
	s.Keyword = &v
	return s
}

func (s *SearchKeywordRequest) SetCity(v string) *SearchKeywordRequest {
	s.City = &v
	return s
}

func (s *SearchKeywordRequest) SetCursor(v int64) *SearchKeywordRequest {
	s.Cursor = &v
	return s
}

func (s *SearchKeywordRequest) SetHeader(v map[string]*string) *SearchKeywordRequest {
	s.Header = v
	return s
}

type SearchKeywordResponse struct {
	Data      *SearchKeywordResponseData      `json:"data,omitempty" xml:"data,omitempty"`
	GwExtra   *SearchKeywordResponseGwExtra   `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	ExtraBody *SearchKeywordResponseExtraBody `json:"extra_body,omitempty" xml:"extra_body,omitempty"`
}

func (s SearchKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordResponse) GoString() string {
	return s.String()
}

func (s *SearchKeywordResponse) SetData(v *SearchKeywordResponseData) *SearchKeywordResponse {
	s.Data = v
	return s
}

func (s *SearchKeywordResponse) SetGwExtra(v *SearchKeywordResponseGwExtra) *SearchKeywordResponse {
	s.GwExtra = v
	return s
}

func (s *SearchKeywordResponse) SetExtraBody(v *SearchKeywordResponseExtraBody) *SearchKeywordResponse {
	s.ExtraBody = v
	return s
}

type SearchKeywordResponseData struct {
	HasMore     *bool                                `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Pois        []*SearchKeywordResponseDataPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ErrorCode   *int64                               `json:"error_code,omitempty" xml:"error_code,omitempty"`
	Description *string                              `json:"description,omitempty" xml:"description,omitempty"`
	Cursor      *int64                               `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s SearchKeywordResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordResponseData) GoString() string {
	return s.String()
}

func (s *SearchKeywordResponseData) SetHasMore(v bool) *SearchKeywordResponseData {
	s.HasMore = &v
	return s
}

func (s *SearchKeywordResponseData) SetPois(v []*SearchKeywordResponseDataPoisItem) *SearchKeywordResponseData {
	s.Pois = v
	return s
}

func (s *SearchKeywordResponseData) SetErrorCode(v int64) *SearchKeywordResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SearchKeywordResponseData) SetDescription(v string) *SearchKeywordResponseData {
	s.Description = &v
	return s
}

func (s *SearchKeywordResponseData) SetCursor(v int64) *SearchKeywordResponseData {
	s.Cursor = &v
	return s
}

type SearchKeywordResponseDataPoisItem struct {
	CityCode    *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	City        *string `json:"city,omitempty" xml:"city,omitempty"`
	Address     *string `json:"address,omitempty" xml:"address,omitempty"`
	PoiId       *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	District    *string `json:"district,omitempty" xml:"district,omitempty"`
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	Country     *string `json:"country,omitempty" xml:"country,omitempty"`
	Location    *string `json:"location,omitempty" xml:"location,omitempty"`
	PoiName     *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	Province    *string `json:"province,omitempty" xml:"province,omitempty"`
}

func (s SearchKeywordResponseDataPoisItem) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordResponseDataPoisItem) GoString() string {
	return s.String()
}

func (s *SearchKeywordResponseDataPoisItem) SetCityCode(v string) *SearchKeywordResponseDataPoisItem {
	s.CityCode = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetCity(v string) *SearchKeywordResponseDataPoisItem {
	s.City = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetAddress(v string) *SearchKeywordResponseDataPoisItem {
	s.Address = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetPoiId(v string) *SearchKeywordResponseDataPoisItem {
	s.PoiId = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetDistrict(v string) *SearchKeywordResponseDataPoisItem {
	s.District = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetCountryCode(v string) *SearchKeywordResponseDataPoisItem {
	s.CountryCode = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetCountry(v string) *SearchKeywordResponseDataPoisItem {
	s.Country = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetLocation(v string) *SearchKeywordResponseDataPoisItem {
	s.Location = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetPoiName(v string) *SearchKeywordResponseDataPoisItem {
	s.PoiName = &v
	return s
}

func (s *SearchKeywordResponseDataPoisItem) SetProvince(v string) *SearchKeywordResponseDataPoisItem {
	s.Province = &v
	return s
}

type SearchKeywordResponseExtraBody struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty"`
	ErrorCode      *int64  `json:"error_code,omitempty" xml:"error_code,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	SubErrorCode   *int64  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty"`
}

func (s SearchKeywordResponseExtraBody) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordResponseExtraBody) GoString() string {
	return s.String()
}

func (s *SearchKeywordResponseExtraBody) SetNow(v int64) *SearchKeywordResponseExtraBody {
	s.Now = &v
	return s
}

func (s *SearchKeywordResponseExtraBody) SetErrorCode(v int64) *SearchKeywordResponseExtraBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchKeywordResponseExtraBody) SetDescription(v string) *SearchKeywordResponseExtraBody {
	s.Description = &v
	return s
}

func (s *SearchKeywordResponseExtraBody) SetSubErrorCode(v int64) *SearchKeywordResponseExtraBody {
	s.SubErrorCode = &v
	return s
}

func (s *SearchKeywordResponseExtraBody) SetSubDescription(v string) *SearchKeywordResponseExtraBody {
	s.SubDescription = &v
	return s
}

func (s *SearchKeywordResponseExtraBody) SetLogid(v string) *SearchKeywordResponseExtraBody {
	s.Logid = &v
	return s
}

type SearchKeywordResponseGwExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SearchKeywordResponseGwExtra) String() string {
	return tea.Prettify(s)
}

func (s SearchKeywordResponseGwExtra) GoString() string {
	return s.String()
}

func (s *SearchKeywordResponseGwExtra) SetSubErrorCode(v int32) *SearchKeywordResponseGwExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SearchKeywordResponseGwExtra) SetSubDescription(v string) *SearchKeywordResponseGwExtra {
	s.SubDescription = &v
	return s
}

func (s *SearchKeywordResponseGwExtra) SetLogid(v string) *SearchKeywordResponseGwExtra {
	s.Logid = &v
	return s
}

func (s *SearchKeywordResponseGwExtra) SetNow(v int64) *SearchKeywordResponseGwExtra {
	s.Now = &v
	return s
}

func (s *SearchKeywordResponseGwExtra) SetErrorCode(v int32) *SearchKeywordResponseGwExtra {
	s.ErrorCode = &v
	return s
}

func (s *SearchKeywordResponseGwExtra) SetDescription(v string) *SearchKeywordResponseGwExtra {
	s.Description = &v
	return s
}

type SearchUploadSitemapRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PagePaths   []*string          `json:"page_paths,omitempty" xml:"page_paths,omitempty" require:"true" type:"Repeated"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SearchUploadSitemapRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchUploadSitemapRequest) GoString() string {
	return s.String()
}

func (s *SearchUploadSitemapRequest) SetAccessToken(v string) *SearchUploadSitemapRequest {
	s.AccessToken = &v
	return s
}

func (s *SearchUploadSitemapRequest) SetPagePaths(v []*string) *SearchUploadSitemapRequest {
	s.PagePaths = v
	return s
}

func (s *SearchUploadSitemapRequest) SetAppId(v string) *SearchUploadSitemapRequest {
	s.AppId = &v
	return s
}

func (s *SearchUploadSitemapRequest) SetHeader(v map[string]*string) *SearchUploadSitemapRequest {
	s.Header = v
	return s
}

type SearchUploadSitemapResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SearchUploadSitemapResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchUploadSitemapResponse) GoString() string {
	return s.String()
}

func (s *SearchUploadSitemapResponse) SetErrNo(v int32) *SearchUploadSitemapResponse {
	s.ErrNo = &v
	return s
}

func (s *SearchUploadSitemapResponse) SetErrMsg(v string) *SearchUploadSitemapResponse {
	s.ErrMsg = &v
	return s
}

func (s *SearchUploadSitemapResponse) SetLogId(v string) *SearchUploadSitemapResponse {
	s.LogId = &v
	return s
}

type SearchVideoRequest struct {
	Keyword     *string            `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
	SearchId    *string            `json:"search_id,omitempty" xml:"search_id,omitempty"`
	SortType    *int32             `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	PublishTime *int32             `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	DeviceId    *int64             `json:"device_id,omitempty" xml:"device_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s SearchVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchVideoRequest) GoString() string {
	return s.String()
}

func (s *SearchVideoRequest) SetKeyword(v string) *SearchVideoRequest {
	s.Keyword = &v
	return s
}

func (s *SearchVideoRequest) SetSearchId(v string) *SearchVideoRequest {
	s.SearchId = &v
	return s
}

func (s *SearchVideoRequest) SetSortType(v int32) *SearchVideoRequest {
	s.SortType = &v
	return s
}

func (s *SearchVideoRequest) SetPublishTime(v int32) *SearchVideoRequest {
	s.PublishTime = &v
	return s
}

func (s *SearchVideoRequest) SetHeader(v map[string]*string) *SearchVideoRequest {
	s.Header = v
	return s
}

func (s *SearchVideoRequest) SetAccessToken(v string) *SearchVideoRequest {
	s.AccessToken = &v
	return s
}

func (s *SearchVideoRequest) SetCount(v int32) *SearchVideoRequest {
	s.Count = &v
	return s
}

func (s *SearchVideoRequest) SetCursor(v int64) *SearchVideoRequest {
	s.Cursor = &v
	return s
}

func (s *SearchVideoRequest) SetDeviceId(v int64) *SearchVideoRequest {
	s.DeviceId = &v
	return s
}

func (s *SearchVideoRequest) SetOpenId(v string) *SearchVideoRequest {
	s.OpenId = &v
	return s
}

type SearchVideoResponse struct {
	Data   *SearchVideoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SearchVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchVideoResponse) GoString() string {
	return s.String()
}

func (s *SearchVideoResponse) SetData(v *SearchVideoResponseData) *SearchVideoResponse {
	s.Data = v
	return s
}

func (s *SearchVideoResponse) SetErrNo(v int32) *SearchVideoResponse {
	s.ErrNo = &v
	return s
}

func (s *SearchVideoResponse) SetErrMsg(v string) *SearchVideoResponse {
	s.ErrMsg = &v
	return s
}

func (s *SearchVideoResponse) SetLogId(v string) *SearchVideoResponse {
	s.LogId = &v
	return s
}

type SearchVideoResponseData struct {
	Cursor    *int64                                  `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore   *bool                                   `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	SearchId  *string                                 `json:"search_id,omitempty" xml:"search_id,omitempty" require:"true"`
	VideoList []*SearchVideoResponseDataVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" require:"true" type:"Repeated"`
}

func (s SearchVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s SearchVideoResponseData) GoString() string {
	return s.String()
}

func (s *SearchVideoResponseData) SetCursor(v int64) *SearchVideoResponseData {
	s.Cursor = &v
	return s
}

func (s *SearchVideoResponseData) SetHasMore(v bool) *SearchVideoResponseData {
	s.HasMore = &v
	return s
}

func (s *SearchVideoResponseData) SetSearchId(v string) *SearchVideoResponseData {
	s.SearchId = &v
	return s
}

func (s *SearchVideoResponseData) SetVideoList(v []*SearchVideoResponseDataVideoListItem) *SearchVideoResponseData {
	s.VideoList = v
	return s
}

type SearchVideoResponseDataVideoListItem struct {
	Avatar      *string                                         `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Link        *string                                         `json:"link,omitempty" xml:"link,omitempty"`
	Cover       *string                                         `json:"cover,omitempty" xml:"cover,omitempty"`
	CoverHeight *int32                                          `json:"cover_height,omitempty" xml:"cover_height,omitempty"`
	Statistics  *SearchVideoResponseDataVideoListItemStatistics `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Duration    *int64                                          `json:"duration,omitempty" xml:"duration,omitempty"`
	CoverWidth  *int32                                          `json:"cover_width,omitempty" xml:"cover_width,omitempty"`
	Nickname    *string                                         `json:"nickname,omitempty" xml:"nickname,omitempty"`
	ItemId      *string                                         `json:"item_id,omitempty" xml:"item_id,omitempty"`
	Title       *string                                         `json:"title,omitempty" xml:"title,omitempty"`
	CreateTime  *int64                                          `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s SearchVideoResponseDataVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SearchVideoResponseDataVideoListItem) GoString() string {
	return s.String()
}

func (s *SearchVideoResponseDataVideoListItem) SetAvatar(v string) *SearchVideoResponseDataVideoListItem {
	s.Avatar = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetLink(v string) *SearchVideoResponseDataVideoListItem {
	s.Link = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetCover(v string) *SearchVideoResponseDataVideoListItem {
	s.Cover = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetCoverHeight(v int32) *SearchVideoResponseDataVideoListItem {
	s.CoverHeight = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetStatistics(v *SearchVideoResponseDataVideoListItemStatistics) *SearchVideoResponseDataVideoListItem {
	s.Statistics = v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetDuration(v int64) *SearchVideoResponseDataVideoListItem {
	s.Duration = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetCoverWidth(v int32) *SearchVideoResponseDataVideoListItem {
	s.CoverWidth = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetNickname(v string) *SearchVideoResponseDataVideoListItem {
	s.Nickname = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetItemId(v string) *SearchVideoResponseDataVideoListItem {
	s.ItemId = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetTitle(v string) *SearchVideoResponseDataVideoListItem {
	s.Title = &v
	return s
}

func (s *SearchVideoResponseDataVideoListItem) SetCreateTime(v int64) *SearchVideoResponseDataVideoListItem {
	s.CreateTime = &v
	return s
}

type SearchVideoResponseDataVideoListItemStatistics struct {
	DiggCount *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
}

func (s SearchVideoResponseDataVideoListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s SearchVideoResponseDataVideoListItemStatistics) GoString() string {
	return s.String()
}

func (s *SearchVideoResponseDataVideoListItemStatistics) SetDiggCount(v int32) *SearchVideoResponseDataVideoListItemStatistics {
	s.DiggCount = &v
	return s
}

type SendCouponToDesignatedUserRequest struct {
	ActivityId       *string            `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	OpenId           *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	AppId            *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	MerchantCouponId *string            `json:"merchant_coupon_id,omitempty" xml:"merchant_coupon_id,omitempty"`
}

func (s SendCouponToDesignatedUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCouponToDesignatedUserRequest) GoString() string {
	return s.String()
}

func (s *SendCouponToDesignatedUserRequest) SetActivityId(v string) *SendCouponToDesignatedUserRequest {
	s.ActivityId = &v
	return s
}

func (s *SendCouponToDesignatedUserRequest) SetOpenId(v string) *SendCouponToDesignatedUserRequest {
	s.OpenId = &v
	return s
}

func (s *SendCouponToDesignatedUserRequest) SetAppId(v string) *SendCouponToDesignatedUserRequest {
	s.AppId = &v
	return s
}

func (s *SendCouponToDesignatedUserRequest) SetHeader(v map[string]*string) *SendCouponToDesignatedUserRequest {
	s.Header = v
	return s
}

func (s *SendCouponToDesignatedUserRequest) SetAccessToken(v string) *SendCouponToDesignatedUserRequest {
	s.AccessToken = &v
	return s
}

func (s *SendCouponToDesignatedUserRequest) SetMerchantCouponId(v string) *SendCouponToDesignatedUserRequest {
	s.MerchantCouponId = &v
	return s
}

type SendCouponToDesignatedUserResponse struct {
	Data   *SendCouponToDesignatedUserResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SendCouponToDesignatedUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCouponToDesignatedUserResponse) GoString() string {
	return s.String()
}

func (s *SendCouponToDesignatedUserResponse) SetData(v *SendCouponToDesignatedUserResponseData) *SendCouponToDesignatedUserResponse {
	s.Data = v
	return s
}

func (s *SendCouponToDesignatedUserResponse) SetErrNo(v int32) *SendCouponToDesignatedUserResponse {
	s.ErrNo = &v
	return s
}

func (s *SendCouponToDesignatedUserResponse) SetErrMsg(v string) *SendCouponToDesignatedUserResponse {
	s.ErrMsg = &v
	return s
}

func (s *SendCouponToDesignatedUserResponse) SetLogId(v string) *SendCouponToDesignatedUserResponse {
	s.LogId = &v
	return s
}

type SendCouponToDesignatedUserResponseData struct {
	CouponInfo *SendCouponToDesignatedUserResponseDataCouponInfo `json:"coupon_info,omitempty" xml:"coupon_info,omitempty" require:"true"`
}

func (s SendCouponToDesignatedUserResponseData) String() string {
	return tea.Prettify(s)
}

func (s SendCouponToDesignatedUserResponseData) GoString() string {
	return s.String()
}

func (s *SendCouponToDesignatedUserResponseData) SetCouponInfo(v *SendCouponToDesignatedUserResponseDataCouponInfo) *SendCouponToDesignatedUserResponseData {
	s.CouponInfo = v
	return s
}

type SendCouponToDesignatedUserResponseDataCouponInfo struct {
	DiscountType   *int32  `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	ValidEndTime   *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty" require:"true"`
	CouponName     *string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty" require:"true"`
	DiscountAmount *int64  `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	ValidBeginTime *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty" require:"true"`
	CouponId       *string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty" require:"true"`
	CouponStatus   *int    `json:"coupon_status,omitempty" xml:"coupon_status,omitempty" require:"true"`
	MinPayAmount   *int64  `json:"min_pay_amount,omitempty" xml:"min_pay_amount,omitempty"`
	MerchantMetaNo *string `json:"merchant_meta_no,omitempty" xml:"merchant_meta_no,omitempty" require:"true"`
	ReceiveTime    *int64  `json:"receive_time,omitempty" xml:"receive_time,omitempty" require:"true"`
}

func (s SendCouponToDesignatedUserResponseDataCouponInfo) String() string {
	return tea.Prettify(s)
}

func (s SendCouponToDesignatedUserResponseDataCouponInfo) GoString() string {
	return s.String()
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetDiscountType(v int32) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.DiscountType = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetValidEndTime(v int64) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.ValidEndTime = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetCouponName(v string) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.CouponName = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetDiscountAmount(v int64) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.DiscountAmount = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetValidBeginTime(v int64) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.ValidBeginTime = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetCouponId(v string) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.CouponId = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetCouponStatus(v int) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.CouponStatus = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetMinPayAmount(v int64) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.MinPayAmount = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetMerchantMetaNo(v string) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.MerchantMetaNo = &v
	return s
}

func (s *SendCouponToDesignatedUserResponseDataCouponInfo) SetReceiveTime(v int64) *SendCouponToDesignatedUserResponseDataCouponInfo {
	s.ReceiveTime = &v
	return s
}

type SendMsgRequest struct {
	Scene          *string                          `json:"scene,omitempty" xml:"scene,omitempty"`
	ContentList    []*SendMsgRequestContentListItem `json:"content_list,omitempty" xml:"content_list,omitempty" type:"Repeated"`
	ConversationId *string                          `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
	Header         map[string]*string               `json:"header,omitempty" xml:"header,omitempty"`
	OpenId         *string                          `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ConversionId   *string                          `json:"conversion_id,omitempty" xml:"conversion_id,omitempty"`
	Content        *SendMsgRequestContent           `json:"content,omitempty" xml:"content,omitempty"`
	AccessToken    *string                          `json:"access_token,omitempty" xml:"access_token,omitempty"`
	MsgId          *string                          `json:"msg_id,omitempty" xml:"msg_id,omitempty"`
	ToUserId       *string                          `json:"to_user_id,omitempty" xml:"to_user_id,omitempty" require:"true"`
}

func (s SendMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequest) GoString() string {
	return s.String()
}

func (s *SendMsgRequest) SetScene(v string) *SendMsgRequest {
	s.Scene = &v
	return s
}

func (s *SendMsgRequest) SetContentList(v []*SendMsgRequestContentListItem) *SendMsgRequest {
	s.ContentList = v
	return s
}

func (s *SendMsgRequest) SetConversationId(v string) *SendMsgRequest {
	s.ConversationId = &v
	return s
}

func (s *SendMsgRequest) SetHeader(v map[string]*string) *SendMsgRequest {
	s.Header = v
	return s
}

func (s *SendMsgRequest) SetOpenId(v string) *SendMsgRequest {
	s.OpenId = &v
	return s
}

func (s *SendMsgRequest) SetConversionId(v string) *SendMsgRequest {
	s.ConversionId = &v
	return s
}

func (s *SendMsgRequest) SetContent(v *SendMsgRequestContent) *SendMsgRequest {
	s.Content = v
	return s
}

func (s *SendMsgRequest) SetAccessToken(v string) *SendMsgRequest {
	s.AccessToken = &v
	return s
}

func (s *SendMsgRequest) SetMsgId(v string) *SendMsgRequest {
	s.MsgId = &v
	return s
}

func (s *SendMsgRequest) SetToUserId(v string) *SendMsgRequest {
	s.ToUserId = &v
	return s
}

type SendMsgRequestContent struct {
	RetainConsultCard      *SendMsgRequestContentRetainConsultCard      `json:"retain_consult_card,omitempty" xml:"retain_consult_card,omitempty"`
	GroupInvitation        *SendMsgRequestContentGroupInvitation        `json:"group_invitation,omitempty" xml:"group_invitation,omitempty"`
	AppletCoupon           *SendMsgRequestContentAppletCoupon           `json:"applet_coupon,omitempty" xml:"applet_coupon,omitempty"`
	MsgType                *int                                         `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	AuthPrivateMessageCard *SendMsgRequestContentAuthPrivateMessageCard `json:"auth_private_message_card,omitempty" xml:"auth_private_message_card,omitempty"`
	Image                  *SendMsgRequestContentImage                  `json:"image,omitempty" xml:"image,omitempty"`
	Text                   *SendMsgRequestContentText                   `json:"text,omitempty" xml:"text,omitempty"`
	Video                  *SendMsgRequestContentVideo                  `json:"video,omitempty" xml:"video,omitempty"`
	AppletCard             *SendMsgRequestContentAppletCard             `json:"applet_card,omitempty" xml:"applet_card,omitempty"`
}

func (s SendMsgRequestContent) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContent) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContent) SetRetainConsultCard(v *SendMsgRequestContentRetainConsultCard) *SendMsgRequestContent {
	s.RetainConsultCard = v
	return s
}

func (s *SendMsgRequestContent) SetGroupInvitation(v *SendMsgRequestContentGroupInvitation) *SendMsgRequestContent {
	s.GroupInvitation = v
	return s
}

func (s *SendMsgRequestContent) SetAppletCoupon(v *SendMsgRequestContentAppletCoupon) *SendMsgRequestContent {
	s.AppletCoupon = v
	return s
}

func (s *SendMsgRequestContent) SetMsgType(v int) *SendMsgRequestContent {
	s.MsgType = &v
	return s
}

func (s *SendMsgRequestContent) SetAuthPrivateMessageCard(v *SendMsgRequestContentAuthPrivateMessageCard) *SendMsgRequestContent {
	s.AuthPrivateMessageCard = v
	return s
}

func (s *SendMsgRequestContent) SetImage(v *SendMsgRequestContentImage) *SendMsgRequestContent {
	s.Image = v
	return s
}

func (s *SendMsgRequestContent) SetText(v *SendMsgRequestContentText) *SendMsgRequestContent {
	s.Text = v
	return s
}

func (s *SendMsgRequestContent) SetVideo(v *SendMsgRequestContentVideo) *SendMsgRequestContent {
	s.Video = v
	return s
}

func (s *SendMsgRequestContent) SetAppletCard(v *SendMsgRequestContentAppletCard) *SendMsgRequestContent {
	s.AppletCard = v
	return s
}

type SendMsgRequestContentAppletCard struct {
	CardId         *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	Schema         *string `json:"schema,omitempty" xml:"schema,omitempty"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

func (s SendMsgRequestContentAppletCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentAppletCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentAppletCard) SetCardId(v string) *SendMsgRequestContentAppletCard {
	s.CardId = &v
	return s
}

func (s *SendMsgRequestContentAppletCard) SetCardTemplateId(v string) *SendMsgRequestContentAppletCard {
	s.CardTemplateId = &v
	return s
}

func (s *SendMsgRequestContentAppletCard) SetPath(v string) *SendMsgRequestContentAppletCard {
	s.Path = &v
	return s
}

func (s *SendMsgRequestContentAppletCard) SetQuery(v string) *SendMsgRequestContentAppletCard {
	s.Query = &v
	return s
}

func (s *SendMsgRequestContentAppletCard) SetSchema(v string) *SendMsgRequestContentAppletCard {
	s.Schema = &v
	return s
}

func (s *SendMsgRequestContentAppletCard) SetAppId(v string) *SendMsgRequestContentAppletCard {
	s.AppId = &v
	return s
}

type SendMsgRequestContentAppletCoupon struct {
	CouponMetaId *int64 `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	ActivityId   *int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

func (s SendMsgRequestContentAppletCoupon) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentAppletCoupon) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentAppletCoupon) SetCouponMetaId(v int64) *SendMsgRequestContentAppletCoupon {
	s.CouponMetaId = &v
	return s
}

func (s *SendMsgRequestContentAppletCoupon) SetActivityId(v int64) *SendMsgRequestContentAppletCoupon {
	s.ActivityId = &v
	return s
}

type SendMsgRequestContentAuthPrivateMessageCard struct {
	AppInfo    *SendMsgRequestContentAuthPrivateMessageCardAppInfo    `json:"app_info,omitempty" xml:"app_info,omitempty"`
	ToUserInfo *SendMsgRequestContentAuthPrivateMessageCardToUserInfo `json:"to_user_info,omitempty" xml:"to_user_info,omitempty"`
	Content    *string                                                `json:"Content,omitempty" xml:"Content,omitempty"`
	JumpUrl    *string                                                `json:"JumpUrl,omitempty" xml:"JumpUrl,omitempty"`
	Title      *string                                                `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SendMsgRequestContentAuthPrivateMessageCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentAuthPrivateMessageCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentAuthPrivateMessageCard) SetAppInfo(v *SendMsgRequestContentAuthPrivateMessageCardAppInfo) *SendMsgRequestContentAuthPrivateMessageCard {
	s.AppInfo = v
	return s
}

func (s *SendMsgRequestContentAuthPrivateMessageCard) SetToUserInfo(v *SendMsgRequestContentAuthPrivateMessageCardToUserInfo) *SendMsgRequestContentAuthPrivateMessageCard {
	s.ToUserInfo = v
	return s
}

func (s *SendMsgRequestContentAuthPrivateMessageCard) SetContent(v string) *SendMsgRequestContentAuthPrivateMessageCard {
	s.Content = &v
	return s
}

func (s *SendMsgRequestContentAuthPrivateMessageCard) SetJumpUrl(v string) *SendMsgRequestContentAuthPrivateMessageCard {
	s.JumpUrl = &v
	return s
}

func (s *SendMsgRequestContentAuthPrivateMessageCard) SetTitle(v string) *SendMsgRequestContentAuthPrivateMessageCard {
	s.Title = &v
	return s
}

type SendMsgRequestContentAuthPrivateMessageCardAppInfo struct {
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s SendMsgRequestContentAuthPrivateMessageCardAppInfo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentAuthPrivateMessageCardAppInfo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentAuthPrivateMessageCardAppInfo) SetAppId(v string) *SendMsgRequestContentAuthPrivateMessageCardAppInfo {
	s.AppId = &v
	return s
}

type SendMsgRequestContentAuthPrivateMessageCardToUserInfo struct {
	OpenId *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AppId  *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

func (s SendMsgRequestContentAuthPrivateMessageCardToUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentAuthPrivateMessageCardToUserInfo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentAuthPrivateMessageCardToUserInfo) SetOpenId(v string) *SendMsgRequestContentAuthPrivateMessageCardToUserInfo {
	s.OpenId = &v
	return s
}

func (s *SendMsgRequestContentAuthPrivateMessageCardToUserInfo) SetAppId(v string) *SendMsgRequestContentAuthPrivateMessageCardToUserInfo {
	s.AppId = &v
	return s
}

type SendMsgRequestContentGroupInvitation struct {
	GroupToken *string `json:"group_token,omitempty" xml:"group_token,omitempty"`
	GroupId    *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s SendMsgRequestContentGroupInvitation) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentGroupInvitation) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentGroupInvitation) SetGroupToken(v string) *SendMsgRequestContentGroupInvitation {
	s.GroupToken = &v
	return s
}

func (s *SendMsgRequestContentGroupInvitation) SetGroupId(v string) *SendMsgRequestContentGroupInvitation {
	s.GroupId = &v
	return s
}

type SendMsgRequestContentImage struct {
	MediaId *string `json:"media_id,omitempty" xml:"media_id,omitempty"`
}

func (s SendMsgRequestContentImage) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentImage) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentImage) SetMediaId(v string) *SendMsgRequestContentImage {
	s.MediaId = &v
	return s
}

type SendMsgRequestContentListItem struct {
	GroupInvitation        *SendMsgRequestContentListItemGroupInvitation        `json:"group_invitation,omitempty" xml:"group_invitation,omitempty"`
	AppletCoupon           *SendMsgRequestContentListItemAppletCoupon           `json:"applet_coupon,omitempty" xml:"applet_coupon,omitempty"`
	RetainConsultCard      *SendMsgRequestContentListItemRetainConsultCard      `json:"retain_consult_card,omitempty" xml:"retain_consult_card,omitempty"`
	Video                  *SendMsgRequestContentListItemVideo                  `json:"video,omitempty" xml:"video,omitempty"`
	MsgType                *int                                                 `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	Text                   *SendMsgRequestContentListItemText                   `json:"text,omitempty" xml:"text,omitempty"`
	Image                  *SendMsgRequestContentListItemImage                  `json:"image,omitempty" xml:"image,omitempty"`
	AppletCard             *SendMsgRequestContentListItemAppletCard             `json:"applet_card,omitempty" xml:"applet_card,omitempty"`
	AuthPrivateMessageCard *SendMsgRequestContentListItemAuthPrivateMessageCard `json:"auth_private_message_card,omitempty" xml:"auth_private_message_card,omitempty"`
}

func (s SendMsgRequestContentListItem) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItem) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItem) SetGroupInvitation(v *SendMsgRequestContentListItemGroupInvitation) *SendMsgRequestContentListItem {
	s.GroupInvitation = v
	return s
}

func (s *SendMsgRequestContentListItem) SetAppletCoupon(v *SendMsgRequestContentListItemAppletCoupon) *SendMsgRequestContentListItem {
	s.AppletCoupon = v
	return s
}

func (s *SendMsgRequestContentListItem) SetRetainConsultCard(v *SendMsgRequestContentListItemRetainConsultCard) *SendMsgRequestContentListItem {
	s.RetainConsultCard = v
	return s
}

func (s *SendMsgRequestContentListItem) SetVideo(v *SendMsgRequestContentListItemVideo) *SendMsgRequestContentListItem {
	s.Video = v
	return s
}

func (s *SendMsgRequestContentListItem) SetMsgType(v int) *SendMsgRequestContentListItem {
	s.MsgType = &v
	return s
}

func (s *SendMsgRequestContentListItem) SetText(v *SendMsgRequestContentListItemText) *SendMsgRequestContentListItem {
	s.Text = v
	return s
}

func (s *SendMsgRequestContentListItem) SetImage(v *SendMsgRequestContentListItemImage) *SendMsgRequestContentListItem {
	s.Image = v
	return s
}

func (s *SendMsgRequestContentListItem) SetAppletCard(v *SendMsgRequestContentListItemAppletCard) *SendMsgRequestContentListItem {
	s.AppletCard = v
	return s
}

func (s *SendMsgRequestContentListItem) SetAuthPrivateMessageCard(v *SendMsgRequestContentListItemAuthPrivateMessageCard) *SendMsgRequestContentListItem {
	s.AuthPrivateMessageCard = v
	return s
}

type SendMsgRequestContentListItemAppletCard struct {
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	Schema         *string `json:"schema,omitempty" xml:"schema,omitempty"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	CardId         *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s SendMsgRequestContentListItemAppletCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemAppletCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemAppletCard) SetQuery(v string) *SendMsgRequestContentListItemAppletCard {
	s.Query = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCard) SetSchema(v string) *SendMsgRequestContentListItemAppletCard {
	s.Schema = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCard) SetAppId(v string) *SendMsgRequestContentListItemAppletCard {
	s.AppId = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCard) SetCardId(v string) *SendMsgRequestContentListItemAppletCard {
	s.CardId = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCard) SetCardTemplateId(v string) *SendMsgRequestContentListItemAppletCard {
	s.CardTemplateId = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCard) SetPath(v string) *SendMsgRequestContentListItemAppletCard {
	s.Path = &v
	return s
}

type SendMsgRequestContentListItemAppletCoupon struct {
	CouponMetaId *int64 `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	ActivityId   *int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

func (s SendMsgRequestContentListItemAppletCoupon) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemAppletCoupon) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemAppletCoupon) SetCouponMetaId(v int64) *SendMsgRequestContentListItemAppletCoupon {
	s.CouponMetaId = &v
	return s
}

func (s *SendMsgRequestContentListItemAppletCoupon) SetActivityId(v int64) *SendMsgRequestContentListItemAppletCoupon {
	s.ActivityId = &v
	return s
}

type SendMsgRequestContentListItemAuthPrivateMessageCard struct {
	Content    *string                                                        `json:"Content,omitempty" xml:"Content,omitempty"`
	JumpUrl    *string                                                        `json:"JumpUrl,omitempty" xml:"JumpUrl,omitempty"`
	Title      *string                                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	AppInfo    *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo    `json:"app_info,omitempty" xml:"app_info,omitempty"`
	ToUserInfo *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo `json:"to_user_info,omitempty" xml:"to_user_info,omitempty"`
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCard) SetContent(v string) *SendMsgRequestContentListItemAuthPrivateMessageCard {
	s.Content = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCard) SetJumpUrl(v string) *SendMsgRequestContentListItemAuthPrivateMessageCard {
	s.JumpUrl = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCard) SetTitle(v string) *SendMsgRequestContentListItemAuthPrivateMessageCard {
	s.Title = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCard) SetAppInfo(v *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) *SendMsgRequestContentListItemAuthPrivateMessageCard {
	s.AppInfo = v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCard) SetToUserInfo(v *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) *SendMsgRequestContentListItemAuthPrivateMessageCard {
	s.ToUserInfo = v
	return s
}

type SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo struct {
	AppId   *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) SetAppId(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo {
	s.AppId = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) SetAppIcon(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo {
	s.AppIcon = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo) SetAppName(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardAppInfo {
	s.AppName = &v
	return s
}

type SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo struct {
	Avatar   *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	UID      *int64  `json:"UID,omitempty" xml:"UID,omitempty"`
	AppId    *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	OpenId   *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) SetAvatar(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo {
	s.Avatar = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) SetNickName(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo {
	s.NickName = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) SetUID(v int64) *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo {
	s.UID = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) SetAppId(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo {
	s.AppId = &v
	return s
}

func (s *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo) SetOpenId(v string) *SendMsgRequestContentListItemAuthPrivateMessageCardToUserInfo {
	s.OpenId = &v
	return s
}

type SendMsgRequestContentListItemGroupInvitation struct {
	GroupId    *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupToken *string `json:"group_token,omitempty" xml:"group_token,omitempty"`
}

func (s SendMsgRequestContentListItemGroupInvitation) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemGroupInvitation) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemGroupInvitation) SetGroupId(v string) *SendMsgRequestContentListItemGroupInvitation {
	s.GroupId = &v
	return s
}

func (s *SendMsgRequestContentListItemGroupInvitation) SetGroupToken(v string) *SendMsgRequestContentListItemGroupInvitation {
	s.GroupToken = &v
	return s
}

type SendMsgRequestContentListItemImage struct {
	MediaId *string `json:"media_id,omitempty" xml:"media_id,omitempty"`
}

func (s SendMsgRequestContentListItemImage) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemImage) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemImage) SetMediaId(v string) *SendMsgRequestContentListItemImage {
	s.MediaId = &v
	return s
}

type SendMsgRequestContentListItemRetainConsultCard struct {
	CardId *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
}

func (s SendMsgRequestContentListItemRetainConsultCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemRetainConsultCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemRetainConsultCard) SetCardId(v string) *SendMsgRequestContentListItemRetainConsultCard {
	s.CardId = &v
	return s
}

type SendMsgRequestContentListItemText struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SendMsgRequestContentListItemText) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemText) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemText) SetText(v string) *SendMsgRequestContentListItemText {
	s.Text = &v
	return s
}

type SendMsgRequestContentListItemVideo struct {
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s SendMsgRequestContentListItemVideo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentListItemVideo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentListItemVideo) SetItemId(v string) *SendMsgRequestContentListItemVideo {
	s.ItemId = &v
	return s
}

type SendMsgRequestContentRetainConsultCard struct {
	CardId *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
}

func (s SendMsgRequestContentRetainConsultCard) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentRetainConsultCard) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentRetainConsultCard) SetCardId(v string) *SendMsgRequestContentRetainConsultCard {
	s.CardId = &v
	return s
}

type SendMsgRequestContentText struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SendMsgRequestContentText) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentText) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentText) SetText(v string) *SendMsgRequestContentText {
	s.Text = &v
	return s
}

type SendMsgRequestContentVideo struct {
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s SendMsgRequestContentVideo) String() string {
	return tea.Prettify(s)
}

func (s SendMsgRequestContentVideo) GoString() string {
	return s.String()
}

func (s *SendMsgRequestContentVideo) SetItemId(v string) *SendMsgRequestContentVideo {
	s.ItemId = &v
	return s
}

type SendMsgResponse struct {
	Extra   *SendMsgResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	MsgId   *string               `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	MsgList []*string             `json:"msg_list,omitempty" xml:"msg_list,omitempty" require:"true" type:"Repeated"`
	Data    *SendMsgResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s SendMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMsgResponse) GoString() string {
	return s.String()
}

func (s *SendMsgResponse) SetExtra(v *SendMsgResponseExtra) *SendMsgResponse {
	s.Extra = v
	return s
}

func (s *SendMsgResponse) SetMsgId(v string) *SendMsgResponse {
	s.MsgId = &v
	return s
}

func (s *SendMsgResponse) SetMsgList(v []*string) *SendMsgResponse {
	s.MsgList = v
	return s
}

func (s *SendMsgResponse) SetData(v *SendMsgResponseData) *SendMsgResponse {
	s.Data = v
	return s
}

type SendMsgResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SendMsgResponseData) String() string {
	return tea.Prettify(s)
}

func (s SendMsgResponseData) GoString() string {
	return s.String()
}

func (s *SendMsgResponseData) SetGwErrorCode(v int32) *SendMsgResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SendMsgResponseData) SetGwDescription(v string) *SendMsgResponseData {
	s.GwDescription = &v
	return s
}

type SendMsgResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s SendMsgResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SendMsgResponseExtra) GoString() string {
	return s.String()
}

func (s *SendMsgResponseExtra) SetSubDescription(v string) *SendMsgResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SendMsgResponseExtra) SetSubErrorCode(v int32) *SendMsgResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SendMsgResponseExtra) SetDescription(v string) *SendMsgResponseExtra {
	s.Description = &v
	return s
}

func (s *SendMsgResponseExtra) SetErrorCode(v int32) *SendMsgResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SendMsgResponseExtra) SetLogid(v string) *SendMsgResponseExtra {
	s.Logid = &v
	return s
}

func (s *SendMsgResponseExtra) SetNow(v int64) *SendMsgResponseExtra {
	s.Now = &v
	return s
}

type ServeQueryRequest struct {
	AccessToken *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *int64                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Param       *ServeQueryRequestParam `json:"param,omitempty" xml:"param,omitempty"`
	Header      map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ServeQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryRequest) GoString() string {
	return s.String()
}

func (s *ServeQueryRequest) SetAccessToken(v string) *ServeQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *ServeQueryRequest) SetAccountId(v int64) *ServeQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ServeQueryRequest) SetParam(v *ServeQueryRequestParam) *ServeQueryRequest {
	s.Param = v
	return s
}

func (s *ServeQueryRequest) SetHeader(v map[string]*string) *ServeQueryRequest {
	s.Header = v
	return s
}

type ServeQueryRequestParam struct {
	PoiIdList         []*string `json:"poi_id_list,omitempty" xml:"poi_id_list,omitempty" type:"Repeated"`
	MerchantAccountId *string   `json:"merchant_account_id,omitempty" xml:"merchant_account_id,omitempty"`
}

func (s ServeQueryRequestParam) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryRequestParam) GoString() string {
	return s.String()
}

func (s *ServeQueryRequestParam) SetPoiIdList(v []*string) *ServeQueryRequestParam {
	s.PoiIdList = v
	return s
}

func (s *ServeQueryRequestParam) SetMerchantAccountId(v string) *ServeQueryRequestParam {
	s.MerchantAccountId = &v
	return s
}

type ServeQueryResponse struct {
	Data  *ServeQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ServeQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ServeQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponse) GoString() string {
	return s.String()
}

func (s *ServeQueryResponse) SetData(v *ServeQueryResponseData) *ServeQueryResponse {
	s.Data = v
	return s
}

func (s *ServeQueryResponse) SetExtra(v *ServeQueryResponseExtra) *ServeQueryResponse {
	s.Extra = v
	return s
}

type ServeQueryResponseData struct {
	StoreOrderServeInfo map[string]*ServeQueryResponseDataStoreOrderServeInfoValue `json:"store_order_serve_info,omitempty" xml:"store_order_serve_info,omitempty"`
	GwErrorCode         *int32                                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription       *string                                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ServeQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponseData) GoString() string {
	return s.String()
}

func (s *ServeQueryResponseData) SetStoreOrderServeInfo(v map[string]*ServeQueryResponseDataStoreOrderServeInfoValue) *ServeQueryResponseData {
	s.StoreOrderServeInfo = v
	return s
}

func (s *ServeQueryResponseData) SetGwErrorCode(v int32) *ServeQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ServeQueryResponseData) SetGwDescription(v string) *ServeQueryResponseData {
	s.GwDescription = &v
	return s
}

type ServeQueryResponseDataStoreOrderServeInfoValue struct {
	ServeStatus *int                                                                           `json:"serve_status,omitempty" xml:"serve_status,omitempty"`
	ServeWeek   map[int32][]*ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem  `json:"serve_week,omitempty" xml:"serve_week,omitempty"`
	ServeDate   map[string][]*ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem `json:"serve_date,omitempty" xml:"serve_date,omitempty"`
}

func (s ServeQueryResponseDataStoreOrderServeInfoValue) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponseDataStoreOrderServeInfoValue) GoString() string {
	return s.String()
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValue) SetServeStatus(v int) *ServeQueryResponseDataStoreOrderServeInfoValue {
	s.ServeStatus = &v
	return s
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValue) SetServeWeek(v map[int32][]*ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem) *ServeQueryResponseDataStoreOrderServeInfoValue {
	s.ServeWeek = v
	return s
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValue) SetServeDate(v map[string][]*ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem) *ServeQueryResponseDataStoreOrderServeInfoValue {
	s.ServeDate = v
	return s
}

type ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem struct {
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem) GoString() string {
	return s.String()
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem) SetStartTime(v string) *ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem {
	s.StartTime = &v
	return s
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem) SetEndTime(v string) *ServeQueryResponseDataStoreOrderServeInfoValueServeDateValueItem {
	s.EndTime = &v
	return s
}

type ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem) GoString() string {
	return s.String()
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem) SetEndTime(v string) *ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem {
	s.EndTime = &v
	return s
}

func (s *ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem) SetStartTime(v string) *ServeQueryResponseDataStoreOrderServeInfoValueServeWeekValueItem {
	s.StartTime = &v
	return s
}

type ServeQueryResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ServeQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ServeQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ServeQueryResponseExtra) SetDescription(v string) *ServeQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ServeQueryResponseExtra) SetErrorCode(v int32) *ServeQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ServeQueryResponseExtra) SetLogid(v string) *ServeQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ServeQueryResponseExtra) SetNow(v int64) *ServeQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *ServeQueryResponseExtra) SetSubDescription(v string) *ServeQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ServeQueryResponseExtra) SetSubErrorCode(v int32) *ServeQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ServeSubmitRequest struct {
	AccountId   *int64                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Param       *ServeSubmitRequestParam `json:"param,omitempty" xml:"param,omitempty"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ServeSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitRequest) GoString() string {
	return s.String()
}

func (s *ServeSubmitRequest) SetAccountId(v int64) *ServeSubmitRequest {
	s.AccountId = &v
	return s
}

func (s *ServeSubmitRequest) SetParam(v *ServeSubmitRequestParam) *ServeSubmitRequest {
	s.Param = v
	return s
}

func (s *ServeSubmitRequest) SetHeader(v map[string]*string) *ServeSubmitRequest {
	s.Header = v
	return s
}

func (s *ServeSubmitRequest) SetAccessToken(v string) *ServeSubmitRequest {
	s.AccessToken = &v
	return s
}

type ServeSubmitRequestParam struct {
	MerchantAccountId *string                                                 `json:"merchant_account_id,omitempty" xml:"merchant_account_id,omitempty"`
	PoiIdList         []*string                                               `json:"poi_id_list,omitempty" xml:"poi_id_list,omitempty" type:"Repeated"`
	ServeDate         map[string][]*ServeSubmitRequestParamServeDateValueItem `json:"serve_date,omitempty" xml:"serve_date,omitempty"`
	ServeStatus       *int                                                    `json:"serve_status,omitempty" xml:"serve_status,omitempty"`
	ServeWeek         map[int32][]*ServeSubmitRequestParamServeWeekValueItem  `json:"serve_week,omitempty" xml:"serve_week,omitempty"`
}

func (s ServeSubmitRequestParam) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitRequestParam) GoString() string {
	return s.String()
}

func (s *ServeSubmitRequestParam) SetMerchantAccountId(v string) *ServeSubmitRequestParam {
	s.MerchantAccountId = &v
	return s
}

func (s *ServeSubmitRequestParam) SetPoiIdList(v []*string) *ServeSubmitRequestParam {
	s.PoiIdList = v
	return s
}

func (s *ServeSubmitRequestParam) SetServeDate(v map[string][]*ServeSubmitRequestParamServeDateValueItem) *ServeSubmitRequestParam {
	s.ServeDate = v
	return s
}

func (s *ServeSubmitRequestParam) SetServeStatus(v int) *ServeSubmitRequestParam {
	s.ServeStatus = &v
	return s
}

func (s *ServeSubmitRequestParam) SetServeWeek(v map[int32][]*ServeSubmitRequestParamServeWeekValueItem) *ServeSubmitRequestParam {
	s.ServeWeek = v
	return s
}

type ServeSubmitRequestParamServeDateValueItem struct {
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s ServeSubmitRequestParamServeDateValueItem) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitRequestParamServeDateValueItem) GoString() string {
	return s.String()
}

func (s *ServeSubmitRequestParamServeDateValueItem) SetStartTime(v string) *ServeSubmitRequestParamServeDateValueItem {
	s.StartTime = &v
	return s
}

func (s *ServeSubmitRequestParamServeDateValueItem) SetEndTime(v string) *ServeSubmitRequestParamServeDateValueItem {
	s.EndTime = &v
	return s
}

type ServeSubmitRequestParamServeWeekValueItem struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ServeSubmitRequestParamServeWeekValueItem) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitRequestParamServeWeekValueItem) GoString() string {
	return s.String()
}

func (s *ServeSubmitRequestParamServeWeekValueItem) SetEndTime(v string) *ServeSubmitRequestParamServeWeekValueItem {
	s.EndTime = &v
	return s
}

func (s *ServeSubmitRequestParamServeWeekValueItem) SetStartTime(v string) *ServeSubmitRequestParamServeWeekValueItem {
	s.StartTime = &v
	return s
}

type ServeSubmitResponse struct {
	Extra *ServeSubmitResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ServeSubmitResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ServeSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitResponse) GoString() string {
	return s.String()
}

func (s *ServeSubmitResponse) SetExtra(v *ServeSubmitResponseExtra) *ServeSubmitResponse {
	s.Extra = v
	return s
}

func (s *ServeSubmitResponse) SetData(v *ServeSubmitResponseData) *ServeSubmitResponse {
	s.Data = v
	return s
}

type ServeSubmitResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ServeSubmitResponseData) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitResponseData) GoString() string {
	return s.String()
}

func (s *ServeSubmitResponseData) SetGwErrorCode(v int32) *ServeSubmitResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ServeSubmitResponseData) SetGwDescription(v string) *ServeSubmitResponseData {
	s.GwDescription = &v
	return s
}

type ServeSubmitResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ServeSubmitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ServeSubmitResponseExtra) GoString() string {
	return s.String()
}

func (s *ServeSubmitResponseExtra) SetDescription(v string) *ServeSubmitResponseExtra {
	s.Description = &v
	return s
}

func (s *ServeSubmitResponseExtra) SetErrorCode(v int32) *ServeSubmitResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ServeSubmitResponseExtra) SetLogid(v string) *ServeSubmitResponseExtra {
	s.Logid = &v
	return s
}

func (s *ServeSubmitResponseExtra) SetNow(v int64) *ServeSubmitResponseExtra {
	s.Now = &v
	return s
}

func (s *ServeSubmitResponseExtra) SetSubDescription(v string) *ServeSubmitResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ServeSubmitResponseExtra) SetSubErrorCode(v int32) *ServeSubmitResponseExtra {
	s.SubErrorCode = &v
	return s
}

type SetUserGroupTagRequest struct {
	Status      *int32             `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	MpId        *string            `json:"mp_id,omitempty" xml:"mp_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TagId       *string            `json:"tag_id,omitempty" xml:"tag_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SetUserGroupTagRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserGroupTagRequest) GoString() string {
	return s.String()
}

func (s *SetUserGroupTagRequest) SetStatus(v int32) *SetUserGroupTagRequest {
	s.Status = &v
	return s
}

func (s *SetUserGroupTagRequest) SetMpId(v string) *SetUserGroupTagRequest {
	s.MpId = &v
	return s
}

func (s *SetUserGroupTagRequest) SetOpenId(v string) *SetUserGroupTagRequest {
	s.OpenId = &v
	return s
}

func (s *SetUserGroupTagRequest) SetTagId(v string) *SetUserGroupTagRequest {
	s.TagId = &v
	return s
}

func (s *SetUserGroupTagRequest) SetHeader(v map[string]*string) *SetUserGroupTagRequest {
	s.Header = v
	return s
}

func (s *SetUserGroupTagRequest) SetAccessToken(v string) *SetUserGroupTagRequest {
	s.AccessToken = &v
	return s
}

type SetUserGroupTagResponse struct {
	ErrMsg           *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId            *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	IsSandboxRequest *bool   `json:"is_sandbox_request,omitempty" xml:"is_sandbox_request,omitempty" require:"true"`
	ErrNo            *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SetUserGroupTagResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserGroupTagResponse) GoString() string {
	return s.String()
}

func (s *SetUserGroupTagResponse) SetErrMsg(v string) *SetUserGroupTagResponse {
	s.ErrMsg = &v
	return s
}

func (s *SetUserGroupTagResponse) SetLogId(v string) *SetUserGroupTagResponse {
	s.LogId = &v
	return s
}

func (s *SetUserGroupTagResponse) SetIsSandboxRequest(v bool) *SetUserGroupTagResponse {
	s.IsSandboxRequest = &v
	return s
}

func (s *SetUserGroupTagResponse) SetErrNo(v int32) *SetUserGroupTagResponse {
	s.ErrNo = &v
	return s
}

type SettingDisableRequest struct {
	GroupId          *string            `json:"group_id,omitempty" xml:"group_id,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId           *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	GroupSettingType *int               `json:"group_setting_type,omitempty" xml:"group_setting_type,omitempty" require:"true"`
}

func (s SettingDisableRequest) String() string {
	return tea.Prettify(s)
}

func (s SettingDisableRequest) GoString() string {
	return s.String()
}

func (s *SettingDisableRequest) SetGroupId(v string) *SettingDisableRequest {
	s.GroupId = &v
	return s
}

func (s *SettingDisableRequest) SetHeader(v map[string]*string) *SettingDisableRequest {
	s.Header = v
	return s
}

func (s *SettingDisableRequest) SetAccessToken(v string) *SettingDisableRequest {
	s.AccessToken = &v
	return s
}

func (s *SettingDisableRequest) SetOpenId(v string) *SettingDisableRequest {
	s.OpenId = &v
	return s
}

func (s *SettingDisableRequest) SetGroupSettingType(v int) *SettingDisableRequest {
	s.GroupSettingType = &v
	return s
}

type SettingDisableResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SettingDisableResponse) String() string {
	return tea.Prettify(s)
}

func (s SettingDisableResponse) GoString() string {
	return s.String()
}

func (s *SettingDisableResponse) SetErrMsg(v string) *SettingDisableResponse {
	s.ErrMsg = &v
	return s
}

func (s *SettingDisableResponse) SetLogId(v string) *SettingDisableResponse {
	s.LogId = &v
	return s
}

func (s *SettingDisableResponse) SetErrNo(v int32) *SettingDisableResponse {
	s.ErrNo = &v
	return s
}

type SettingSetRequest struct {
	Header           map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId           *string                         `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	MsgList          []*SettingSetRequestMsgListItem `json:"msg_list,omitempty" xml:"msg_list,omitempty" require:"true" type:"Repeated"`
	GroupSettingType *int                            `json:"group_setting_type,omitempty" xml:"group_setting_type,omitempty" require:"true"`
}

func (s SettingSetRequest) String() string {
	return tea.Prettify(s)
}

func (s SettingSetRequest) GoString() string {
	return s.String()
}

func (s *SettingSetRequest) SetHeader(v map[string]*string) *SettingSetRequest {
	s.Header = v
	return s
}

func (s *SettingSetRequest) SetAccessToken(v string) *SettingSetRequest {
	s.AccessToken = &v
	return s
}

func (s *SettingSetRequest) SetOpenId(v string) *SettingSetRequest {
	s.OpenId = &v
	return s
}

func (s *SettingSetRequest) SetMsgList(v []*SettingSetRequestMsgListItem) *SettingSetRequest {
	s.MsgList = v
	return s
}

func (s *SettingSetRequest) SetGroupSettingType(v int) *SettingSetRequest {
	s.GroupSettingType = &v
	return s
}

type SettingSetRequestMsgListItem struct {
	AppletCoupon *SettingSetRequestMsgListItemAppletCoupon `json:"applet_coupon,omitempty" xml:"applet_coupon,omitempty"`
	MsgType      *int                                      `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	Text         *SettingSetRequestMsgListItemText         `json:"text,omitempty" xml:"text,omitempty"`
	AppletCard   *SettingSetRequestMsgListItemAppletCard   `json:"applet_card,omitempty" xml:"applet_card,omitempty"`
}

func (s SettingSetRequestMsgListItem) String() string {
	return tea.Prettify(s)
}

func (s SettingSetRequestMsgListItem) GoString() string {
	return s.String()
}

func (s *SettingSetRequestMsgListItem) SetAppletCoupon(v *SettingSetRequestMsgListItemAppletCoupon) *SettingSetRequestMsgListItem {
	s.AppletCoupon = v
	return s
}

func (s *SettingSetRequestMsgListItem) SetMsgType(v int) *SettingSetRequestMsgListItem {
	s.MsgType = &v
	return s
}

func (s *SettingSetRequestMsgListItem) SetText(v *SettingSetRequestMsgListItemText) *SettingSetRequestMsgListItem {
	s.Text = v
	return s
}

func (s *SettingSetRequestMsgListItem) SetAppletCard(v *SettingSetRequestMsgListItemAppletCard) *SettingSetRequestMsgListItem {
	s.AppletCard = v
	return s
}

type SettingSetRequestMsgListItemAppletCard struct {
	Query          *string `json:"query,omitempty" xml:"query,omitempty"`
	Path           *string `json:"path,omitempty" xml:"path,omitempty"`
	AppId          *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	CardId         *string `json:"card_id,omitempty" xml:"card_id,omitempty"`
	CardTemplateId *string `json:"card_template_id,omitempty" xml:"card_template_id,omitempty"`
}

func (s SettingSetRequestMsgListItemAppletCard) String() string {
	return tea.Prettify(s)
}

func (s SettingSetRequestMsgListItemAppletCard) GoString() string {
	return s.String()
}

func (s *SettingSetRequestMsgListItemAppletCard) SetQuery(v string) *SettingSetRequestMsgListItemAppletCard {
	s.Query = &v
	return s
}

func (s *SettingSetRequestMsgListItemAppletCard) SetPath(v string) *SettingSetRequestMsgListItemAppletCard {
	s.Path = &v
	return s
}

func (s *SettingSetRequestMsgListItemAppletCard) SetAppId(v string) *SettingSetRequestMsgListItemAppletCard {
	s.AppId = &v
	return s
}

func (s *SettingSetRequestMsgListItemAppletCard) SetCardId(v string) *SettingSetRequestMsgListItemAppletCard {
	s.CardId = &v
	return s
}

func (s *SettingSetRequestMsgListItemAppletCard) SetCardTemplateId(v string) *SettingSetRequestMsgListItemAppletCard {
	s.CardTemplateId = &v
	return s
}

type SettingSetRequestMsgListItemAppletCoupon struct {
	CouponMetaId *int64 `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	ActivityId   *int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

func (s SettingSetRequestMsgListItemAppletCoupon) String() string {
	return tea.Prettify(s)
}

func (s SettingSetRequestMsgListItemAppletCoupon) GoString() string {
	return s.String()
}

func (s *SettingSetRequestMsgListItemAppletCoupon) SetCouponMetaId(v int64) *SettingSetRequestMsgListItemAppletCoupon {
	s.CouponMetaId = &v
	return s
}

func (s *SettingSetRequestMsgListItemAppletCoupon) SetActivityId(v int64) *SettingSetRequestMsgListItemAppletCoupon {
	s.ActivityId = &v
	return s
}

type SettingSetRequestMsgListItemText struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SettingSetRequestMsgListItemText) String() string {
	return tea.Prettify(s)
}

func (s SettingSetRequestMsgListItemText) GoString() string {
	return s.String()
}

func (s *SettingSetRequestMsgListItemText) SetText(v string) *SettingSetRequestMsgListItemText {
	s.Text = &v
	return s
}

type SettingSetResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SettingSetResponse) String() string {
	return tea.Prettify(s)
}

func (s SettingSetResponse) GoString() string {
	return s.String()
}

func (s *SettingSetResponse) SetErrNo(v int32) *SettingSetResponse {
	s.ErrNo = &v
	return s
}

func (s *SettingSetResponse) SetErrMsg(v string) *SettingSetResponse {
	s.ErrMsg = &v
	return s
}

func (s *SettingSetResponse) SetLogId(v string) *SettingSetResponse {
	s.LogId = &v
	return s
}

type ShareCreateActivityIdRequest struct {
	Base        *ShareCreateActivityIdRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string                `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                           `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ShareCreateActivityIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdRequest) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdRequest) SetBase(v *ShareCreateActivityIdRequestBase) *ShareCreateActivityIdRequest {
	s.Base = v
	return s
}

func (s *ShareCreateActivityIdRequest) SetHeader(v map[string]*string) *ShareCreateActivityIdRequest {
	s.Header = v
	return s
}

func (s *ShareCreateActivityIdRequest) SetAccessToken(v string) *ShareCreateActivityIdRequest {
	s.AccessToken = &v
	return s
}

type ShareCreateActivityIdRequestBase struct {
	Addr       *string                                     `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Client     *string                                     `json:"Client,omitempty" xml:"Client,omitempty"`
	TrafficEnv *ShareCreateActivityIdRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Extra      map[string]*string                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                     `json:"LogID,omitempty" xml:"LogID,omitempty"`
	Caller     *string                                     `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s ShareCreateActivityIdRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdRequestBase) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdRequestBase) SetAddr(v string) *ShareCreateActivityIdRequestBase {
	s.Addr = &v
	return s
}

func (s *ShareCreateActivityIdRequestBase) SetClient(v string) *ShareCreateActivityIdRequestBase {
	s.Client = &v
	return s
}

func (s *ShareCreateActivityIdRequestBase) SetTrafficEnv(v *ShareCreateActivityIdRequestBaseTrafficEnv) *ShareCreateActivityIdRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ShareCreateActivityIdRequestBase) SetExtra(v map[string]*string) *ShareCreateActivityIdRequestBase {
	s.Extra = v
	return s
}

func (s *ShareCreateActivityIdRequestBase) SetLogID(v string) *ShareCreateActivityIdRequestBase {
	s.LogID = &v
	return s
}

func (s *ShareCreateActivityIdRequestBase) SetCaller(v string) *ShareCreateActivityIdRequestBase {
	s.Caller = &v
	return s
}

type ShareCreateActivityIdRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ShareCreateActivityIdRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdRequestBaseTrafficEnv) SetEnv(v string) *ShareCreateActivityIdRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ShareCreateActivityIdRequestBaseTrafficEnv) SetOpen(v bool) *ShareCreateActivityIdRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ShareCreateActivityIdResponse struct {
	BaseResp *ShareCreateActivityIdResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty"`
	ErrNo    *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg   *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId    *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data     *ShareCreateActivityIdResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ShareCreateActivityIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdResponse) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdResponse) SetBaseResp(v *ShareCreateActivityIdResponseBaseResp) *ShareCreateActivityIdResponse {
	s.BaseResp = v
	return s
}

func (s *ShareCreateActivityIdResponse) SetErrNo(v int32) *ShareCreateActivityIdResponse {
	s.ErrNo = &v
	return s
}

func (s *ShareCreateActivityIdResponse) SetErrMsg(v string) *ShareCreateActivityIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShareCreateActivityIdResponse) SetLogId(v string) *ShareCreateActivityIdResponse {
	s.LogId = &v
	return s
}

func (s *ShareCreateActivityIdResponse) SetData(v *ShareCreateActivityIdResponseData) *ShareCreateActivityIdResponse {
	s.Data = v
	return s
}

type ShareCreateActivityIdResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ShareCreateActivityIdResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdResponseBaseResp) SetStatusMessage(v string) *ShareCreateActivityIdResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ShareCreateActivityIdResponseBaseResp) SetStatusCode(v int32) *ShareCreateActivityIdResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ShareCreateActivityIdResponseBaseResp) SetExtra(v map[string]*string) *ShareCreateActivityIdResponseBaseResp {
	s.Extra = v
	return s
}

type ShareCreateActivityIdResponseData struct {
	ExpireTime *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	ActivityId *string `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ShareCreateActivityIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateActivityIdResponseData) GoString() string {
	return s.String()
}

func (s *ShareCreateActivityIdResponseData) SetExpireTime(v int64) *ShareCreateActivityIdResponseData {
	s.ExpireTime = &v
	return s
}

func (s *ShareCreateActivityIdResponseData) SetActivityId(v string) *ShareCreateActivityIdResponseData {
	s.ActivityId = &v
	return s
}

type ShareCreateTaskRequest struct {
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskType    *int32             `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	TargetCount *int64             `json:"target_count,omitempty" xml:"target_count,omitempty" require:"true"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s ShareCreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateTaskRequest) GoString() string {
	return s.String()
}

func (s *ShareCreateTaskRequest) SetEndTime(v int64) *ShareCreateTaskRequest {
	s.EndTime = &v
	return s
}

func (s *ShareCreateTaskRequest) SetHeader(v map[string]*string) *ShareCreateTaskRequest {
	s.Header = v
	return s
}

func (s *ShareCreateTaskRequest) SetAccessToken(v string) *ShareCreateTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *ShareCreateTaskRequest) SetTaskType(v int32) *ShareCreateTaskRequest {
	s.TaskType = &v
	return s
}

func (s *ShareCreateTaskRequest) SetTargetCount(v int64) *ShareCreateTaskRequest {
	s.TargetCount = &v
	return s
}

func (s *ShareCreateTaskRequest) SetStartTime(v int64) *ShareCreateTaskRequest {
	s.StartTime = &v
	return s
}

type ShareCreateTaskResponse struct {
	Data   *ShareCreateTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ShareCreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateTaskResponse) GoString() string {
	return s.String()
}

func (s *ShareCreateTaskResponse) SetData(v *ShareCreateTaskResponseData) *ShareCreateTaskResponse {
	s.Data = v
	return s
}

func (s *ShareCreateTaskResponse) SetErrNo(v int32) *ShareCreateTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *ShareCreateTaskResponse) SetErrMsg(v string) *ShareCreateTaskResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShareCreateTaskResponse) SetLogId(v string) *ShareCreateTaskResponse {
	s.LogId = &v
	return s
}

type ShareCreateTaskResponseData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ShareCreateTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s ShareCreateTaskResponseData) GoString() string {
	return s.String()
}

func (s *ShareCreateTaskResponseData) SetTaskId(v string) *ShareCreateTaskResponseData {
	s.TaskId = &v
	return s
}

type ShareQueryUserTaskRequest struct {
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ShareQueryUserTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareQueryUserTaskRequest) GoString() string {
	return s.String()
}

func (s *ShareQueryUserTaskRequest) SetTaskId(v string) *ShareQueryUserTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ShareQueryUserTaskRequest) SetOpenId(v string) *ShareQueryUserTaskRequest {
	s.OpenId = &v
	return s
}

func (s *ShareQueryUserTaskRequest) SetHeader(v map[string]*string) *ShareQueryUserTaskRequest {
	s.Header = v
	return s
}

func (s *ShareQueryUserTaskRequest) SetAccessToken(v string) *ShareQueryUserTaskRequest {
	s.AccessToken = &v
	return s
}

type ShareQueryUserTaskResponse struct {
	Data   *ShareQueryUserTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ShareQueryUserTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareQueryUserTaskResponse) GoString() string {
	return s.String()
}

func (s *ShareQueryUserTaskResponse) SetData(v *ShareQueryUserTaskResponseData) *ShareQueryUserTaskResponse {
	s.Data = v
	return s
}

func (s *ShareQueryUserTaskResponse) SetErrNo(v int32) *ShareQueryUserTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *ShareQueryUserTaskResponse) SetErrMsg(v string) *ShareQueryUserTaskResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShareQueryUserTaskResponse) SetLogId(v string) *ShareQueryUserTaskResponse {
	s.LogId = &v
	return s
}

type ShareQueryUserTaskResponseData struct {
	IsValid      *bool   `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	SuccessCount *int64  `json:"success_count,omitempty" xml:"success_count,omitempty"`
	TargetCount  *int64  `json:"target_count,omitempty" xml:"target_count,omitempty"`
	TaskId       *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	Completed    *bool   `json:"completed,omitempty" xml:"completed,omitempty"`
}

func (s ShareQueryUserTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s ShareQueryUserTaskResponseData) GoString() string {
	return s.String()
}

func (s *ShareQueryUserTaskResponseData) SetIsValid(v bool) *ShareQueryUserTaskResponseData {
	s.IsValid = &v
	return s
}

func (s *ShareQueryUserTaskResponseData) SetSuccessCount(v int64) *ShareQueryUserTaskResponseData {
	s.SuccessCount = &v
	return s
}

func (s *ShareQueryUserTaskResponseData) SetTargetCount(v int64) *ShareQueryUserTaskResponseData {
	s.TargetCount = &v
	return s
}

func (s *ShareQueryUserTaskResponseData) SetTaskId(v string) *ShareQueryUserTaskResponseData {
	s.TaskId = &v
	return s
}

func (s *ShareQueryUserTaskResponseData) SetCompleted(v bool) *ShareQueryUserTaskResponseData {
	s.Completed = &v
	return s
}

type ShareUnbindUnionGroupRequest struct {
	MpId        *string            `json:"mp_id,omitempty" xml:"mp_id,omitempty" require:"true"`
	UnionId     *string            `json:"union_id,omitempty" xml:"union_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ShareUnbindUnionGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareUnbindUnionGroupRequest) GoString() string {
	return s.String()
}

func (s *ShareUnbindUnionGroupRequest) SetMpId(v string) *ShareUnbindUnionGroupRequest {
	s.MpId = &v
	return s
}

func (s *ShareUnbindUnionGroupRequest) SetUnionId(v string) *ShareUnbindUnionGroupRequest {
	s.UnionId = &v
	return s
}

func (s *ShareUnbindUnionGroupRequest) SetHeader(v map[string]*string) *ShareUnbindUnionGroupRequest {
	s.Header = v
	return s
}

func (s *ShareUnbindUnionGroupRequest) SetAccessToken(v string) *ShareUnbindUnionGroupRequest {
	s.AccessToken = &v
	return s
}

type ShareUnbindUnionGroupResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ShareUnbindUnionGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareUnbindUnionGroupResponse) GoString() string {
	return s.String()
}

func (s *ShareUnbindUnionGroupResponse) SetErrNo(v int32) *ShareUnbindUnionGroupResponse {
	s.ErrNo = &v
	return s
}

func (s *ShareUnbindUnionGroupResponse) SetErrMsg(v string) *ShareUnbindUnionGroupResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShareUnbindUnionGroupResponse) SetLogId(v string) *ShareUnbindUnionGroupResponse {
	s.LogId = &v
	return s
}

type ShareUpdateDynamicMessageRequest struct {
	TargetState  *int32                                              `json:"target_state,omitempty" xml:"target_state,omitempty" require:"true"`
	Header       map[string]*string                                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Query        *string                                             `json:"query,omitempty" xml:"query,omitempty"`
	VersionType  *string                                             `json:"version_type,omitempty" xml:"version_type,omitempty"`
	TemplateInfo []*ShareUpdateDynamicMessageRequestTemplateInfoItem `json:"template_info,omitempty" xml:"template_info,omitempty" type:"Repeated"`
	ActivityId   *string                                             `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ShareUpdateDynamicMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareUpdateDynamicMessageRequest) GoString() string {
	return s.String()
}

func (s *ShareUpdateDynamicMessageRequest) SetTargetState(v int32) *ShareUpdateDynamicMessageRequest {
	s.TargetState = &v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetHeader(v map[string]*string) *ShareUpdateDynamicMessageRequest {
	s.Header = v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetAccessToken(v string) *ShareUpdateDynamicMessageRequest {
	s.AccessToken = &v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetQuery(v string) *ShareUpdateDynamicMessageRequest {
	s.Query = &v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetVersionType(v string) *ShareUpdateDynamicMessageRequest {
	s.VersionType = &v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetTemplateInfo(v []*ShareUpdateDynamicMessageRequestTemplateInfoItem) *ShareUpdateDynamicMessageRequest {
	s.TemplateInfo = v
	return s
}

func (s *ShareUpdateDynamicMessageRequest) SetActivityId(v string) *ShareUpdateDynamicMessageRequest {
	s.ActivityId = &v
	return s
}

type ShareUpdateDynamicMessageRequestTemplateInfoItem struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s ShareUpdateDynamicMessageRequestTemplateInfoItem) String() string {
	return tea.Prettify(s)
}

func (s ShareUpdateDynamicMessageRequestTemplateInfoItem) GoString() string {
	return s.String()
}

func (s *ShareUpdateDynamicMessageRequestTemplateInfoItem) SetValue(v string) *ShareUpdateDynamicMessageRequestTemplateInfoItem {
	s.Value = &v
	return s
}

func (s *ShareUpdateDynamicMessageRequestTemplateInfoItem) SetName(v string) *ShareUpdateDynamicMessageRequestTemplateInfoItem {
	s.Name = &v
	return s
}

type ShareUpdateDynamicMessageResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ShareUpdateDynamicMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareUpdateDynamicMessageResponse) GoString() string {
	return s.String()
}

func (s *ShareUpdateDynamicMessageResponse) SetErrMsg(v string) *ShareUpdateDynamicMessageResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShareUpdateDynamicMessageResponse) SetLogId(v string) *ShareUpdateDynamicMessageResponse {
	s.LogId = &v
	return s
}

func (s *ShareUpdateDynamicMessageResponse) SetErrNo(v int32) *ShareUpdateDynamicMessageResponse {
	s.ErrNo = &v
	return s
}

type ShareidRequest struct {
	NeedCallback   *bool              `json:"need_callback,omitempty" xml:"need_callback,omitempty"`
	SourceStyleId  *string            `json:"source_style_id,omitempty" xml:"source_style_id,omitempty"`
	DefaultHashtag *string            `json:"default_hashtag,omitempty" xml:"default_hashtag,omitempty"`
	LinkParam      *string            `json:"link_param,omitempty" xml:"link_param,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ShareidRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareidRequest) GoString() string {
	return s.String()
}

func (s *ShareidRequest) SetNeedCallback(v bool) *ShareidRequest {
	s.NeedCallback = &v
	return s
}

func (s *ShareidRequest) SetSourceStyleId(v string) *ShareidRequest {
	s.SourceStyleId = &v
	return s
}

func (s *ShareidRequest) SetDefaultHashtag(v string) *ShareidRequest {
	s.DefaultHashtag = &v
	return s
}

func (s *ShareidRequest) SetLinkParam(v string) *ShareidRequest {
	s.LinkParam = &v
	return s
}

func (s *ShareidRequest) SetHeader(v map[string]*string) *ShareidRequest {
	s.Header = v
	return s
}

func (s *ShareidRequest) SetAccessToken(v string) *ShareidRequest {
	s.AccessToken = &v
	return s
}

type ShareidResponse struct {
	Data  *ShareidResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ShareidResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ShareidResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareidResponse) GoString() string {
	return s.String()
}

func (s *ShareidResponse) SetData(v *ShareidResponseData) *ShareidResponse {
	s.Data = v
	return s
}

func (s *ShareidResponse) SetExtra(v *ShareidResponseExtra) *ShareidResponse {
	s.Extra = v
	return s
}

type ShareidResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int64  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ShareId     *string `json:"share_id,omitempty" xml:"share_id,omitempty" require:"true"`
}

func (s ShareidResponseData) String() string {
	return tea.Prettify(s)
}

func (s ShareidResponseData) GoString() string {
	return s.String()
}

func (s *ShareidResponseData) SetDescription(v string) *ShareidResponseData {
	s.Description = &v
	return s
}

func (s *ShareidResponseData) SetErrorCode(v int64) *ShareidResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ShareidResponseData) SetShareId(v string) *ShareidResponseData {
	s.ShareId = &v
	return s
}

type ShareidResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ShareidResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ShareidResponseExtra) GoString() string {
	return s.String()
}

func (s *ShareidResponseExtra) SetSubDescription(v string) *ShareidResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ShareidResponseExtra) SetSubErrorCode(v int32) *ShareidResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ShareidResponseExtra) SetDescription(v string) *ShareidResponseExtra {
	s.Description = &v
	return s
}

func (s *ShareidResponseExtra) SetErrorCode(v int32) *ShareidResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ShareidResponseExtra) SetLogid(v string) *ShareidResponseExtra {
	s.Logid = &v
	return s
}

func (s *ShareidResponseExtra) SetNow(v int64) *ShareidResponseExtra {
	s.Now = &v
	return s
}

type ShopMemberLeaveRequest struct {
	ShopId      *int64             `json:"shop_id,omitempty" xml:"shop_id,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ShopMemberLeaveRequest) String() string {
	return tea.Prettify(s)
}

func (s ShopMemberLeaveRequest) GoString() string {
	return s.String()
}

func (s *ShopMemberLeaveRequest) SetShopId(v int64) *ShopMemberLeaveRequest {
	s.ShopId = &v
	return s
}

func (s *ShopMemberLeaveRequest) SetAppId(v string) *ShopMemberLeaveRequest {
	s.AppId = &v
	return s
}

func (s *ShopMemberLeaveRequest) SetOpenId(v string) *ShopMemberLeaveRequest {
	s.OpenId = &v
	return s
}

func (s *ShopMemberLeaveRequest) SetHeader(v map[string]*string) *ShopMemberLeaveRequest {
	s.Header = v
	return s
}

func (s *ShopMemberLeaveRequest) SetAccessToken(v string) *ShopMemberLeaveRequest {
	s.AccessToken = &v
	return s
}

type ShopMemberLeaveResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ShopMemberLeaveResponse) String() string {
	return tea.Prettify(s)
}

func (s ShopMemberLeaveResponse) GoString() string {
	return s.String()
}

func (s *ShopMemberLeaveResponse) SetErrMsg(v string) *ShopMemberLeaveResponse {
	s.ErrMsg = &v
	return s
}

func (s *ShopMemberLeaveResponse) SetErrNo(v int32) *ShopMemberLeaveResponse {
	s.ErrNo = &v
	return s
}

func (s *ShopMemberLeaveResponse) SetLogId(v string) *ShopMemberLeaveResponse {
	s.LogId = &v
	return s
}

type ShopPoiQueryRequest struct {
	PoiId        *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	RelationType *int               `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	Size         *int32             `json:"size,omitempty" xml:"size,omitempty"`
	ThirdId      *string            `json:"third_id,omitempty" xml:"third_id,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId    *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Page         *int32             `json:"page,omitempty" xml:"page,omitempty"`
}

func (s ShopPoiQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryRequest) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryRequest) SetPoiId(v string) *ShopPoiQueryRequest {
	s.PoiId = &v
	return s
}

func (s *ShopPoiQueryRequest) SetRelationType(v int) *ShopPoiQueryRequest {
	s.RelationType = &v
	return s
}

func (s *ShopPoiQueryRequest) SetSize(v int32) *ShopPoiQueryRequest {
	s.Size = &v
	return s
}

func (s *ShopPoiQueryRequest) SetThirdId(v string) *ShopPoiQueryRequest {
	s.ThirdId = &v
	return s
}

func (s *ShopPoiQueryRequest) SetHeader(v map[string]*string) *ShopPoiQueryRequest {
	s.Header = v
	return s
}

func (s *ShopPoiQueryRequest) SetAccessToken(v string) *ShopPoiQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *ShopPoiQueryRequest) SetAccountId(v string) *ShopPoiQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ShopPoiQueryRequest) SetPage(v int32) *ShopPoiQueryRequest {
	s.Page = &v
	return s
}

type ShopPoiQueryResponse struct {
	Data  *ShopPoiQueryResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *ShopPoiQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ShopPoiQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponse) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponse) SetData(v *ShopPoiQueryResponseData) *ShopPoiQueryResponse {
	s.Data = v
	return s
}

func (s *ShopPoiQueryResponse) SetExtra(v *ShopPoiQueryResponseExtra) *ShopPoiQueryResponse {
	s.Extra = v
	return s
}

type ShopPoiQueryResponseData struct {
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Pois          []*ShopPoiQueryResponseDataPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	Total         *int64                              `json:"total,omitempty" xml:"total,omitempty"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ShopPoiQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseData) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseData) SetGwDescription(v string) *ShopPoiQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *ShopPoiQueryResponseData) SetPois(v []*ShopPoiQueryResponseDataPoisItem) *ShopPoiQueryResponseData {
	s.Pois = v
	return s
}

func (s *ShopPoiQueryResponseData) SetTotal(v int64) *ShopPoiQueryResponseData {
	s.Total = &v
	return s
}

func (s *ShopPoiQueryResponseData) SetGwErrorCode(v int32) *ShopPoiQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type ShopPoiQueryResponseDataPoisItem struct {
	RootAccount *ShopPoiQueryResponseDataPoisItemRootAccount `json:"root_account,omitempty" xml:"root_account,omitempty"`
	Account     *ShopPoiQueryResponseDataPoisItemAccount     `json:"account,omitempty" xml:"account,omitempty"`
	Poi         *ShopPoiQueryResponseDataPoisItemPoi         `json:"poi,omitempty" xml:"poi,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItem) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItem) SetRootAccount(v *ShopPoiQueryResponseDataPoisItemRootAccount) *ShopPoiQueryResponseDataPoisItem {
	s.RootAccount = v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItem) SetAccount(v *ShopPoiQueryResponseDataPoisItemAccount) *ShopPoiQueryResponseDataPoisItem {
	s.Account = v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItem) SetPoi(v *ShopPoiQueryResponseDataPoisItemPoi) *ShopPoiQueryResponseDataPoisItem {
	s.Poi = v
	return s
}

type ShopPoiQueryResponseDataPoisItemAccount struct {
	PoiAccount    *ShopPoiQueryResponseDataPoisItemAccountPoiAccount    `json:"poi_account,omitempty" xml:"poi_account,omitempty"`
	RootAccount   *ShopPoiQueryResponseDataPoisItemAccountRootAccount   `json:"root_account,omitempty" xml:"root_account,omitempty"`
	ParentAccount *ShopPoiQueryResponseDataPoisItemAccountParentAccount `json:"parent_account,omitempty" xml:"parent_account,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemAccount) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemAccount) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemAccount) SetPoiAccount(v *ShopPoiQueryResponseDataPoisItemAccountPoiAccount) *ShopPoiQueryResponseDataPoisItemAccount {
	s.PoiAccount = v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccount) SetRootAccount(v *ShopPoiQueryResponseDataPoisItemAccountRootAccount) *ShopPoiQueryResponseDataPoisItemAccount {
	s.RootAccount = v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccount) SetParentAccount(v *ShopPoiQueryResponseDataPoisItemAccountParentAccount) *ShopPoiQueryResponseDataPoisItemAccount {
	s.ParentAccount = v
	return s
}

type ShopPoiQueryResponseDataPoisItemAccountParentAccount struct {
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemAccountParentAccount) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemAccountParentAccount) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemAccountParentAccount) SetAccountType(v string) *ShopPoiQueryResponseDataPoisItemAccountParentAccount {
	s.AccountType = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountParentAccount) SetAccountId(v string) *ShopPoiQueryResponseDataPoisItemAccountParentAccount {
	s.AccountId = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountParentAccount) SetAccountName(v string) *ShopPoiQueryResponseDataPoisItemAccountParentAccount {
	s.AccountName = &v
	return s
}

type ShopPoiQueryResponseDataPoisItemAccountPoiAccount struct {
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemAccountPoiAccount) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemAccountPoiAccount) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemAccountPoiAccount) SetAccountId(v string) *ShopPoiQueryResponseDataPoisItemAccountPoiAccount {
	s.AccountId = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountPoiAccount) SetAccountName(v string) *ShopPoiQueryResponseDataPoisItemAccountPoiAccount {
	s.AccountName = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountPoiAccount) SetAccountType(v string) *ShopPoiQueryResponseDataPoisItemAccountPoiAccount {
	s.AccountType = &v
	return s
}

type ShopPoiQueryResponseDataPoisItemAccountRootAccount struct {
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemAccountRootAccount) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemAccountRootAccount) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemAccountRootAccount) SetAccountId(v string) *ShopPoiQueryResponseDataPoisItemAccountRootAccount {
	s.AccountId = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountRootAccount) SetAccountName(v string) *ShopPoiQueryResponseDataPoisItemAccountRootAccount {
	s.AccountName = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemAccountRootAccount) SetAccountType(v string) *ShopPoiQueryResponseDataPoisItemAccountRootAccount {
	s.AccountType = &v
	return s
}

type ShopPoiQueryResponseDataPoisItemPoi struct {
	Address   *string  `json:"address,omitempty" xml:"address,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	PoiId     *string  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName   *string  `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemPoi) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemPoi) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemPoi) SetAddress(v string) *ShopPoiQueryResponseDataPoisItemPoi {
	s.Address = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemPoi) SetLatitude(v float64) *ShopPoiQueryResponseDataPoisItemPoi {
	s.Latitude = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemPoi) SetLongitude(v float64) *ShopPoiQueryResponseDataPoisItemPoi {
	s.Longitude = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemPoi) SetPoiId(v string) *ShopPoiQueryResponseDataPoisItemPoi {
	s.PoiId = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemPoi) SetPoiName(v string) *ShopPoiQueryResponseDataPoisItemPoi {
	s.PoiName = &v
	return s
}

type ShopPoiQueryResponseDataPoisItemRootAccount struct {
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty"`
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s ShopPoiQueryResponseDataPoisItemRootAccount) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseDataPoisItemRootAccount) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseDataPoisItemRootAccount) SetAccountName(v string) *ShopPoiQueryResponseDataPoisItemRootAccount {
	s.AccountName = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemRootAccount) SetAccountType(v string) *ShopPoiQueryResponseDataPoisItemRootAccount {
	s.AccountType = &v
	return s
}

func (s *ShopPoiQueryResponseDataPoisItemRootAccount) SetAccountId(v string) *ShopPoiQueryResponseDataPoisItemRootAccount {
	s.AccountId = &v
	return s
}

type ShopPoiQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ShopPoiQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ShopPoiQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ShopPoiQueryResponseExtra) SetLogid(v string) *ShopPoiQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ShopPoiQueryResponseExtra) SetNow(v int64) *ShopPoiQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *ShopPoiQueryResponseExtra) SetSubDescription(v string) *ShopPoiQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ShopPoiQueryResponseExtra) SetSubErrorCode(v int32) *ShopPoiQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ShopPoiQueryResponseExtra) SetDescription(v string) *ShopPoiQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ShopPoiQueryResponseExtra) SetErrorCode(v int32) *ShopPoiQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type SkuGetRequest struct {
	ProductOutId *string            `json:"product_out_id,omitempty" xml:"product_out_id,omitempty" require:"true"`
	SkuIds       []*string          `json:"sku_ids,omitempty" xml:"sku_ids,omitempty" type:"Repeated"`
	Base         *SkuGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId    *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutSkuIds    []*string          `json:"out_sku_ids,omitempty" xml:"out_sku_ids,omitempty" type:"Repeated"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SkuGetRequest) String() string {
	return tea.Prettify(s)
}

func (s SkuGetRequest) GoString() string {
	return s.String()
}

func (s *SkuGetRequest) SetProductOutId(v string) *SkuGetRequest {
	s.ProductOutId = &v
	return s
}

func (s *SkuGetRequest) SetSkuIds(v []*string) *SkuGetRequest {
	s.SkuIds = v
	return s
}

func (s *SkuGetRequest) SetBase(v *SkuGetRequestBase) *SkuGetRequest {
	s.Base = v
	return s
}

func (s *SkuGetRequest) SetAccountId(v string) *SkuGetRequest {
	s.AccountId = &v
	return s
}

func (s *SkuGetRequest) SetOutSkuIds(v []*string) *SkuGetRequest {
	s.OutSkuIds = v
	return s
}

func (s *SkuGetRequest) SetHeader(v map[string]*string) *SkuGetRequest {
	s.Header = v
	return s
}

func (s *SkuGetRequest) SetAccessToken(v string) *SkuGetRequest {
	s.AccessToken = &v
	return s
}

type SkuGetRequestBase struct {
	Client     *string                      `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string           `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                      `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *SkuGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                      `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                      `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s SkuGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s SkuGetRequestBase) GoString() string {
	return s.String()
}

func (s *SkuGetRequestBase) SetClient(v string) *SkuGetRequestBase {
	s.Client = &v
	return s
}

func (s *SkuGetRequestBase) SetExtra(v map[string]*string) *SkuGetRequestBase {
	s.Extra = v
	return s
}

func (s *SkuGetRequestBase) SetLogID(v string) *SkuGetRequestBase {
	s.LogID = &v
	return s
}

func (s *SkuGetRequestBase) SetTrafficEnv(v *SkuGetRequestBaseTrafficEnv) *SkuGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *SkuGetRequestBase) SetAddr(v string) *SkuGetRequestBase {
	s.Addr = &v
	return s
}

func (s *SkuGetRequestBase) SetCaller(v string) *SkuGetRequestBase {
	s.Caller = &v
	return s
}

type SkuGetRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s SkuGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s SkuGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *SkuGetRequestBaseTrafficEnv) SetEnv(v string) *SkuGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *SkuGetRequestBaseTrafficEnv) SetOpen(v bool) *SkuGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type SkuGetResponse struct {
	Extra    *SkuGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *SkuGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *SkuGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SkuGetResponse) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponse) GoString() string {
	return s.String()
}

func (s *SkuGetResponse) SetExtra(v *SkuGetResponseExtra) *SkuGetResponse {
	s.Extra = v
	return s
}

func (s *SkuGetResponse) SetBaseResp(v *SkuGetResponseBaseResp) *SkuGetResponse {
	s.BaseResp = v
	return s
}

func (s *SkuGetResponse) SetData(v *SkuGetResponseData) *SkuGetResponse {
	s.Data = v
	return s
}

type SkuGetResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SkuGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *SkuGetResponseBaseResp) SetStatusMessage(v string) *SkuGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *SkuGetResponseBaseResp) SetExtra(v map[string]*string) *SkuGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *SkuGetResponseBaseResp) SetStatusCode(v int32) *SkuGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

type SkuGetResponseData struct {
	Description *string                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Skus        []*SkuGetResponseDataSkusItem `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
}

func (s SkuGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseData) GoString() string {
	return s.String()
}

func (s *SkuGetResponseData) SetDescription(v string) *SkuGetResponseData {
	s.Description = &v
	return s
}

func (s *SkuGetResponseData) SetErrorCode(v int32) *SkuGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SkuGetResponseData) SetSkus(v []*SkuGetResponseDataSkusItem) *SkuGetResponseData {
	s.Skus = v
	return s
}

type SkuGetResponseDataSkusItem struct {
	OriginAmount    *int64                            `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	AttrKeyValueMap map[string]*string                `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	BindSkus        []*string                         `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuName         *string                           `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	CreateTime      *int64                            `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OutSkuId        *string                           `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	UpdateTime      *int64                            `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ActualAmount    *int64                            `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	SkuId           *string                           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Extra           *string                           `json:"extra,omitempty" xml:"extra,omitempty"`
	Stock           *SkuGetResponseDataSkusItemStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	Status          *int                              `json:"status,omitempty" xml:"status,omitempty"`
	SkuExt          *SkuGetResponseDataSkusItemSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
}

func (s SkuGetResponseDataSkusItem) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseDataSkusItem) GoString() string {
	return s.String()
}

func (s *SkuGetResponseDataSkusItem) SetOriginAmount(v int64) *SkuGetResponseDataSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetAttrKeyValueMap(v map[string]*string) *SkuGetResponseDataSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetBindSkus(v []*string) *SkuGetResponseDataSkusItem {
	s.BindSkus = v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetSkuName(v string) *SkuGetResponseDataSkusItem {
	s.SkuName = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetCreateTime(v int64) *SkuGetResponseDataSkusItem {
	s.CreateTime = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetOutSkuId(v string) *SkuGetResponseDataSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetUpdateTime(v int64) *SkuGetResponseDataSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetActualAmount(v int64) *SkuGetResponseDataSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetSkuId(v string) *SkuGetResponseDataSkusItem {
	s.SkuId = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetExtra(v string) *SkuGetResponseDataSkusItem {
	s.Extra = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetStock(v *SkuGetResponseDataSkusItemStock) *SkuGetResponseDataSkusItem {
	s.Stock = v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetStatus(v int) *SkuGetResponseDataSkusItem {
	s.Status = &v
	return s
}

func (s *SkuGetResponseDataSkusItem) SetSkuExt(v *SkuGetResponseDataSkusItemSkuExt) *SkuGetResponseDataSkusItem {
	s.SkuExt = v
	return s
}

type SkuGetResponseDataSkusItemSkuExt struct {
	BindSkuId           *int64                                               `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	TakeawayPresaleInfo *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	BindSkus2c          map[int64][]*int64                                   `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	OriSkus             map[int64][]*int64                                   `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	RelRuleList         []*SkuGetResponseDataSkusItemSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	BizId2cList         []*int64                                             `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	BindProductId       *int64                                               `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	LifeBizCode         *string                                              `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	OriginStockQty      *int64                                               `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
}

func (s SkuGetResponseDataSkusItemSkuExt) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseDataSkusItemSkuExt) GoString() string {
	return s.String()
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetBindSkuId(v int64) *SkuGetResponseDataSkusItemSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetTakeawayPresaleInfo(v *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo) *SkuGetResponseDataSkusItemSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetBindSkus2c(v map[int64][]*int64) *SkuGetResponseDataSkusItemSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetOriSkus(v map[int64][]*int64) *SkuGetResponseDataSkusItemSkuExt {
	s.OriSkus = v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetRelRuleList(v []*SkuGetResponseDataSkusItemSkuExtRelRuleListItem) *SkuGetResponseDataSkusItemSkuExt {
	s.RelRuleList = v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetBizId2cList(v []*int64) *SkuGetResponseDataSkusItemSkuExt {
	s.BizId2cList = v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetBindProductId(v int64) *SkuGetResponseDataSkusItemSkuExt {
	s.BindProductId = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetLifeBizCode(v string) *SkuGetResponseDataSkusItemSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExt) SetOriginStockQty(v int64) *SkuGetResponseDataSkusItemSkuExt {
	s.OriginStockQty = &v
	return s
}

type SkuGetResponseDataSkusItemSkuExtRelRuleListItem struct {
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
}

func (s SkuGetResponseDataSkusItemSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseDataSkusItemSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetConstantVal(v int64) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetPriceRel(v bool) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetSharedQty(v int64) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetStockRel(v bool) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetBizId(v int64) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtRelRuleListItem) SetCoefficient(v string) *SkuGetResponseDataSkusItemSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

type SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
}

func (s SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

func (s *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *SkuGetResponseDataSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

type SkuGetResponseDataSkusItemStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s SkuGetResponseDataSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseDataSkusItemStock) GoString() string {
	return s.String()
}

func (s *SkuGetResponseDataSkusItemStock) SetSoldCount(v int64) *SkuGetResponseDataSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *SkuGetResponseDataSkusItemStock) SetSoldQty(v int64) *SkuGetResponseDataSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *SkuGetResponseDataSkusItemStock) SetStockQty(v int64) *SkuGetResponseDataSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *SkuGetResponseDataSkusItemStock) SetAvailQty(v int64) *SkuGetResponseDataSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *SkuGetResponseDataSkusItemStock) SetFrozenQty(v int64) *SkuGetResponseDataSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *SkuGetResponseDataSkusItemStock) SetLimitType(v int) *SkuGetResponseDataSkusItemStock {
	s.LimitType = &v
	return s
}

type SkuGetResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s SkuGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SkuGetResponseExtra) GoString() string {
	return s.String()
}

func (s *SkuGetResponseExtra) SetSubErrorCode(v int32) *SkuGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SkuGetResponseExtra) SetDescription(v string) *SkuGetResponseExtra {
	s.Description = &v
	return s
}

func (s *SkuGetResponseExtra) SetErrorCode(v int32) *SkuGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SkuGetResponseExtra) SetLogid(v string) *SkuGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *SkuGetResponseExtra) SetNow(v int64) *SkuGetResponseExtra {
	s.Now = &v
	return s
}

func (s *SkuGetResponseExtra) SetSubDescription(v string) *SkuGetResponseExtra {
	s.SubDescription = &v
	return s
}

type SkuUpsertRequest struct {
	AccountId   *string                            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId       *string                            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SkuInfoList []*SkuUpsertRequestSkuInfoListItem `json:"sku_info_list,omitempty" xml:"sku_info_list,omitempty" type:"Repeated"`
	TimeSlot    *int64                             `json:"time_slot,omitempty" xml:"time_slot,omitempty"`
	Header      map[string]*string                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SkuUpsertRequest) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertRequest) GoString() string {
	return s.String()
}

func (s *SkuUpsertRequest) SetAccountId(v string) *SkuUpsertRequest {
	s.AccountId = &v
	return s
}

func (s *SkuUpsertRequest) SetPoiId(v string) *SkuUpsertRequest {
	s.PoiId = &v
	return s
}

func (s *SkuUpsertRequest) SetSkuInfoList(v []*SkuUpsertRequestSkuInfoListItem) *SkuUpsertRequest {
	s.SkuInfoList = v
	return s
}

func (s *SkuUpsertRequest) SetTimeSlot(v int64) *SkuUpsertRequest {
	s.TimeSlot = &v
	return s
}

func (s *SkuUpsertRequest) SetHeader(v map[string]*string) *SkuUpsertRequest {
	s.Header = v
	return s
}

func (s *SkuUpsertRequest) SetAccessToken(v string) *SkuUpsertRequest {
	s.AccessToken = &v
	return s
}

type SkuUpsertRequestSkuInfoListItem struct {
	SkuOperateType *int    `json:"sku_operate_type,omitempty" xml:"sku_operate_type,omitempty"`
	SkuOutId       *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	MaxPersonNum   *int64  `json:"max_person_num,omitempty" xml:"max_person_num,omitempty"`
	MinPersonNum   *int64  `json:"min_person_num,omitempty" xml:"min_person_num,omitempty"`
	SkuId          *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName        *string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
}

func (s SkuUpsertRequestSkuInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertRequestSkuInfoListItem) GoString() string {
	return s.String()
}

func (s *SkuUpsertRequestSkuInfoListItem) SetSkuOperateType(v int) *SkuUpsertRequestSkuInfoListItem {
	s.SkuOperateType = &v
	return s
}

func (s *SkuUpsertRequestSkuInfoListItem) SetSkuOutId(v string) *SkuUpsertRequestSkuInfoListItem {
	s.SkuOutId = &v
	return s
}

func (s *SkuUpsertRequestSkuInfoListItem) SetMaxPersonNum(v int64) *SkuUpsertRequestSkuInfoListItem {
	s.MaxPersonNum = &v
	return s
}

func (s *SkuUpsertRequestSkuInfoListItem) SetMinPersonNum(v int64) *SkuUpsertRequestSkuInfoListItem {
	s.MinPersonNum = &v
	return s
}

func (s *SkuUpsertRequestSkuInfoListItem) SetSkuId(v string) *SkuUpsertRequestSkuInfoListItem {
	s.SkuId = &v
	return s
}

func (s *SkuUpsertRequestSkuInfoListItem) SetSkuName(v string) *SkuUpsertRequestSkuInfoListItem {
	s.SkuName = &v
	return s
}

type SkuUpsertResponse struct {
	Data  *SkuUpsertResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SkuUpsertResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SkuUpsertResponse) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertResponse) GoString() string {
	return s.String()
}

func (s *SkuUpsertResponse) SetData(v *SkuUpsertResponseData) *SkuUpsertResponse {
	s.Data = v
	return s
}

func (s *SkuUpsertResponse) SetExtra(v *SkuUpsertResponseExtra) *SkuUpsertResponse {
	s.Extra = v
	return s
}

type SkuUpsertResponseData struct {
	Description *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SkuInfoList []*SkuUpsertResponseDataSkuInfoListItem `json:"sku_info_list,omitempty" xml:"sku_info_list,omitempty" type:"Repeated"`
}

func (s SkuUpsertResponseData) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertResponseData) GoString() string {
	return s.String()
}

func (s *SkuUpsertResponseData) SetDescription(v string) *SkuUpsertResponseData {
	s.Description = &v
	return s
}

func (s *SkuUpsertResponseData) SetErrorCode(v int32) *SkuUpsertResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SkuUpsertResponseData) SetSkuInfoList(v []*SkuUpsertResponseDataSkuInfoListItem) *SkuUpsertResponseData {
	s.SkuInfoList = v
	return s
}

type SkuUpsertResponseDataSkuInfoListItem struct {
	SkuId    *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuOutId *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
}

func (s SkuUpsertResponseDataSkuInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertResponseDataSkuInfoListItem) GoString() string {
	return s.String()
}

func (s *SkuUpsertResponseDataSkuInfoListItem) SetSkuId(v string) *SkuUpsertResponseDataSkuInfoListItem {
	s.SkuId = &v
	return s
}

func (s *SkuUpsertResponseDataSkuInfoListItem) SetSkuOutId(v string) *SkuUpsertResponseDataSkuInfoListItem {
	s.SkuOutId = &v
	return s
}

type SkuUpsertResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SkuUpsertResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SkuUpsertResponseExtra) GoString() string {
	return s.String()
}

func (s *SkuUpsertResponseExtra) SetErrorCode(v int32) *SkuUpsertResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SkuUpsertResponseExtra) SetLogid(v string) *SkuUpsertResponseExtra {
	s.Logid = &v
	return s
}

func (s *SkuUpsertResponseExtra) SetNow(v int64) *SkuUpsertResponseExtra {
	s.Now = &v
	return s
}

func (s *SkuUpsertResponseExtra) SetSubDescription(v string) *SkuUpsertResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SkuUpsertResponseExtra) SetSubErrorCode(v int32) *SkuUpsertResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SkuUpsertResponseExtra) SetDescription(v string) *SkuUpsertResponseExtra {
	s.Description = &v
	return s
}

type SportBasketballRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportBasketballRequest) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballRequest) GoString() string {
	return s.String()
}

func (s *SportBasketballRequest) SetHeader(v map[string]*string) *SportBasketballRequest {
	s.Header = v
	return s
}

func (s *SportBasketballRequest) SetAccessToken(v string) *SportBasketballRequest {
	s.AccessToken = &v
	return s
}

type SportBasketballResponse struct {
	Data  *SportBasketballResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SportBasketballResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SportBasketballResponse) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballResponse) GoString() string {
	return s.String()
}

func (s *SportBasketballResponse) SetData(v *SportBasketballResponseData) *SportBasketballResponse {
	s.Data = v
	return s
}

func (s *SportBasketballResponse) SetExtra(v *SportBasketballResponseExtra) *SportBasketballResponse {
	s.Extra = v
	return s
}

type SportBasketballResponseData struct {
	List          []*SportBasketballResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportBasketballResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballResponseData) GoString() string {
	return s.String()
}

func (s *SportBasketballResponseData) SetList(v []*SportBasketballResponseDataListItem) *SportBasketballResponseData {
	s.List = v
	return s
}

func (s *SportBasketballResponseData) SetGwErrorCode(v int32) *SportBasketballResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportBasketballResponseData) SetGwDescription(v string) *SportBasketballResponseData {
	s.GwDescription = &v
	return s
}

type SportBasketballResponseDataListItem struct {
	RankChange       *string                                             `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                             `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                             `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                              `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                              `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                            `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportBasketballResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                              `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s SportBasketballResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportBasketballResponseDataListItem) SetRankChange(v string) *SportBasketballResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetNickname(v string) *SportBasketballResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetAvatar(v string) *SportBasketballResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetFollowerCount(v int64) *SportBasketballResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetOnbillbaordTimes(v int32) *SportBasketballResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetEffectValue(v float64) *SportBasketballResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportBasketballResponseDataListItem) SetVideoList(v []*SportBasketballResponseDataListItemVideoListItem) *SportBasketballResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportBasketballResponseDataListItem) SetRank(v int32) *SportBasketballResponseDataListItem {
	s.Rank = &v
	return s
}

type SportBasketballResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SportBasketballResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportBasketballResponseDataListItemVideoListItem) SetItemCover(v string) *SportBasketballResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *SportBasketballResponseDataListItemVideoListItem) SetShareUrl(v string) *SportBasketballResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportBasketballResponseDataListItemVideoListItem) SetTitle(v string) *SportBasketballResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type SportBasketballResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s SportBasketballResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportBasketballResponseExtra) GoString() string {
	return s.String()
}

func (s *SportBasketballResponseExtra) SetErrorCode(v int32) *SportBasketballResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportBasketballResponseExtra) SetDescription(v string) *SportBasketballResponseExtra {
	s.Description = &v
	return s
}

func (s *SportBasketballResponseExtra) SetSubErrorCode(v int32) *SportBasketballResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportBasketballResponseExtra) SetSubDescription(v string) *SportBasketballResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportBasketballResponseExtra) SetLogid(v string) *SportBasketballResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportBasketballResponseExtra) SetNow(v int64) *SportBasketballResponseExtra {
	s.Now = &v
	return s
}

type SportComprehensiveRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SportComprehensiveRequest) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveRequest) GoString() string {
	return s.String()
}

func (s *SportComprehensiveRequest) SetAccessToken(v string) *SportComprehensiveRequest {
	s.AccessToken = &v
	return s
}

func (s *SportComprehensiveRequest) SetHeader(v map[string]*string) *SportComprehensiveRequest {
	s.Header = v
	return s
}

type SportComprehensiveResponse struct {
	Data  *SportComprehensiveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SportComprehensiveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SportComprehensiveResponse) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveResponse) GoString() string {
	return s.String()
}

func (s *SportComprehensiveResponse) SetData(v *SportComprehensiveResponseData) *SportComprehensiveResponse {
	s.Data = v
	return s
}

func (s *SportComprehensiveResponse) SetExtra(v *SportComprehensiveResponseExtra) *SportComprehensiveResponse {
	s.Extra = v
	return s
}

type SportComprehensiveResponseData struct {
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*SportComprehensiveResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s SportComprehensiveResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveResponseData) GoString() string {
	return s.String()
}

func (s *SportComprehensiveResponseData) SetGwDescription(v string) *SportComprehensiveResponseData {
	s.GwDescription = &v
	return s
}

func (s *SportComprehensiveResponseData) SetList(v []*SportComprehensiveResponseDataListItem) *SportComprehensiveResponseData {
	s.List = v
	return s
}

func (s *SportComprehensiveResponseData) SetGwErrorCode(v int32) *SportComprehensiveResponseData {
	s.GwErrorCode = &v
	return s
}

type SportComprehensiveResponseDataListItem struct {
	RankChange       *string                                                `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                                `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                                `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                                 `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                                 `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                               `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportComprehensiveResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                                 `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s SportComprehensiveResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportComprehensiveResponseDataListItem) SetRankChange(v string) *SportComprehensiveResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetNickname(v string) *SportComprehensiveResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetAvatar(v string) *SportComprehensiveResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetFollowerCount(v int64) *SportComprehensiveResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetOnbillbaordTimes(v int32) *SportComprehensiveResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetEffectValue(v float64) *SportComprehensiveResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetVideoList(v []*SportComprehensiveResponseDataListItemVideoListItem) *SportComprehensiveResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportComprehensiveResponseDataListItem) SetRank(v int32) *SportComprehensiveResponseDataListItem {
	s.Rank = &v
	return s
}

type SportComprehensiveResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s SportComprehensiveResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportComprehensiveResponseDataListItemVideoListItem) SetShareUrl(v string) *SportComprehensiveResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportComprehensiveResponseDataListItemVideoListItem) SetTitle(v string) *SportComprehensiveResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *SportComprehensiveResponseDataListItemVideoListItem) SetItemCover(v string) *SportComprehensiveResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type SportComprehensiveResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s SportComprehensiveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportComprehensiveResponseExtra) GoString() string {
	return s.String()
}

func (s *SportComprehensiveResponseExtra) SetErrorCode(v int32) *SportComprehensiveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportComprehensiveResponseExtra) SetDescription(v string) *SportComprehensiveResponseExtra {
	s.Description = &v
	return s
}

func (s *SportComprehensiveResponseExtra) SetSubErrorCode(v int32) *SportComprehensiveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportComprehensiveResponseExtra) SetSubDescription(v string) *SportComprehensiveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportComprehensiveResponseExtra) SetLogid(v string) *SportComprehensiveResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportComprehensiveResponseExtra) SetNow(v int64) *SportComprehensiveResponseExtra {
	s.Now = &v
	return s
}

type SportCultureRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportCultureRequest) String() string {
	return tea.Prettify(s)
}

func (s SportCultureRequest) GoString() string {
	return s.String()
}

func (s *SportCultureRequest) SetHeader(v map[string]*string) *SportCultureRequest {
	s.Header = v
	return s
}

func (s *SportCultureRequest) SetAccessToken(v string) *SportCultureRequest {
	s.AccessToken = &v
	return s
}

type SportCultureResponse struct {
	Data  *SportCultureResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SportCultureResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SportCultureResponse) String() string {
	return tea.Prettify(s)
}

func (s SportCultureResponse) GoString() string {
	return s.String()
}

func (s *SportCultureResponse) SetData(v *SportCultureResponseData) *SportCultureResponse {
	s.Data = v
	return s
}

func (s *SportCultureResponse) SetExtra(v *SportCultureResponseExtra) *SportCultureResponse {
	s.Extra = v
	return s
}

type SportCultureResponseData struct {
	List          []*SportCultureResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportCultureResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportCultureResponseData) GoString() string {
	return s.String()
}

func (s *SportCultureResponseData) SetList(v []*SportCultureResponseDataListItem) *SportCultureResponseData {
	s.List = v
	return s
}

func (s *SportCultureResponseData) SetGwErrorCode(v int32) *SportCultureResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportCultureResponseData) SetGwDescription(v string) *SportCultureResponseData {
	s.GwDescription = &v
	return s
}

type SportCultureResponseDataListItem struct {
	VideoList        []*SportCultureResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
}

func (s SportCultureResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportCultureResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportCultureResponseDataListItem) SetVideoList(v []*SportCultureResponseDataListItemVideoListItem) *SportCultureResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportCultureResponseDataListItem) SetRank(v int32) *SportCultureResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetRankChange(v string) *SportCultureResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetNickname(v string) *SportCultureResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetAvatar(v string) *SportCultureResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetFollowerCount(v int64) *SportCultureResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetOnbillbaordTimes(v int32) *SportCultureResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportCultureResponseDataListItem) SetEffectValue(v float64) *SportCultureResponseDataListItem {
	s.EffectValue = &v
	return s
}

type SportCultureResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s SportCultureResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportCultureResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportCultureResponseDataListItemVideoListItem) SetTitle(v string) *SportCultureResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *SportCultureResponseDataListItemVideoListItem) SetItemCover(v string) *SportCultureResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *SportCultureResponseDataListItemVideoListItem) SetShareUrl(v string) *SportCultureResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type SportCultureResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportCultureResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportCultureResponseExtra) GoString() string {
	return s.String()
}

func (s *SportCultureResponseExtra) SetSubErrorCode(v int32) *SportCultureResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportCultureResponseExtra) SetSubDescription(v string) *SportCultureResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportCultureResponseExtra) SetLogid(v string) *SportCultureResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportCultureResponseExtra) SetNow(v int64) *SportCultureResponseExtra {
	s.Now = &v
	return s
}

func (s *SportCultureResponseExtra) SetErrorCode(v int32) *SportCultureResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportCultureResponseExtra) SetDescription(v string) *SportCultureResponseExtra {
	s.Description = &v
	return s
}

type SportFitnessRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportFitnessRequest) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessRequest) GoString() string {
	return s.String()
}

func (s *SportFitnessRequest) SetHeader(v map[string]*string) *SportFitnessRequest {
	s.Header = v
	return s
}

func (s *SportFitnessRequest) SetAccessToken(v string) *SportFitnessRequest {
	s.AccessToken = &v
	return s
}

type SportFitnessResponse struct {
	Data  *SportFitnessResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SportFitnessResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SportFitnessResponse) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessResponse) GoString() string {
	return s.String()
}

func (s *SportFitnessResponse) SetData(v *SportFitnessResponseData) *SportFitnessResponse {
	s.Data = v
	return s
}

func (s *SportFitnessResponse) SetExtra(v *SportFitnessResponseExtra) *SportFitnessResponse {
	s.Extra = v
	return s
}

type SportFitnessResponseData struct {
	List          []*SportFitnessResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportFitnessResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessResponseData) GoString() string {
	return s.String()
}

func (s *SportFitnessResponseData) SetList(v []*SportFitnessResponseDataListItem) *SportFitnessResponseData {
	s.List = v
	return s
}

func (s *SportFitnessResponseData) SetGwErrorCode(v int32) *SportFitnessResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportFitnessResponseData) SetGwDescription(v string) *SportFitnessResponseData {
	s.GwDescription = &v
	return s
}

type SportFitnessResponseDataListItem struct {
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportFitnessResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s SportFitnessResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportFitnessResponseDataListItem) SetFollowerCount(v int64) *SportFitnessResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetOnbillbaordTimes(v int32) *SportFitnessResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetEffectValue(v float64) *SportFitnessResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetVideoList(v []*SportFitnessResponseDataListItemVideoListItem) *SportFitnessResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportFitnessResponseDataListItem) SetRank(v int32) *SportFitnessResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetRankChange(v string) *SportFitnessResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetNickname(v string) *SportFitnessResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportFitnessResponseDataListItem) SetAvatar(v string) *SportFitnessResponseDataListItem {
	s.Avatar = &v
	return s
}

type SportFitnessResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s SportFitnessResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportFitnessResponseDataListItemVideoListItem) SetTitle(v string) *SportFitnessResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *SportFitnessResponseDataListItemVideoListItem) SetItemCover(v string) *SportFitnessResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *SportFitnessResponseDataListItemVideoListItem) SetShareUrl(v string) *SportFitnessResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type SportFitnessResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s SportFitnessResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportFitnessResponseExtra) GoString() string {
	return s.String()
}

func (s *SportFitnessResponseExtra) SetDescription(v string) *SportFitnessResponseExtra {
	s.Description = &v
	return s
}

func (s *SportFitnessResponseExtra) SetSubErrorCode(v int32) *SportFitnessResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportFitnessResponseExtra) SetSubDescription(v string) *SportFitnessResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportFitnessResponseExtra) SetLogid(v string) *SportFitnessResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportFitnessResponseExtra) SetNow(v int64) *SportFitnessResponseExtra {
	s.Now = &v
	return s
}

func (s *SportFitnessResponseExtra) SetErrorCode(v int32) *SportFitnessResponseExtra {
	s.ErrorCode = &v
	return s
}

type SportOutdoorsRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportOutdoorsRequest) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsRequest) GoString() string {
	return s.String()
}

func (s *SportOutdoorsRequest) SetHeader(v map[string]*string) *SportOutdoorsRequest {
	s.Header = v
	return s
}

func (s *SportOutdoorsRequest) SetAccessToken(v string) *SportOutdoorsRequest {
	s.AccessToken = &v
	return s
}

type SportOutdoorsResponse struct {
	Data  *SportOutdoorsResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SportOutdoorsResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SportOutdoorsResponse) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsResponse) GoString() string {
	return s.String()
}

func (s *SportOutdoorsResponse) SetData(v *SportOutdoorsResponseData) *SportOutdoorsResponse {
	s.Data = v
	return s
}

func (s *SportOutdoorsResponse) SetExtra(v *SportOutdoorsResponseExtra) *SportOutdoorsResponse {
	s.Extra = v
	return s
}

type SportOutdoorsResponseData struct {
	List          []*SportOutdoorsResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportOutdoorsResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsResponseData) GoString() string {
	return s.String()
}

func (s *SportOutdoorsResponseData) SetList(v []*SportOutdoorsResponseDataListItem) *SportOutdoorsResponseData {
	s.List = v
	return s
}

func (s *SportOutdoorsResponseData) SetGwErrorCode(v int32) *SportOutdoorsResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportOutdoorsResponseData) SetGwDescription(v string) *SportOutdoorsResponseData {
	s.GwDescription = &v
	return s
}

type SportOutdoorsResponseDataListItem struct {
	FollowerCount    *int64                                            `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                            `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                          `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportOutdoorsResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                            `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                           `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                           `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                           `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s SportOutdoorsResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportOutdoorsResponseDataListItem) SetFollowerCount(v int64) *SportOutdoorsResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetOnbillbaordTimes(v int32) *SportOutdoorsResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetEffectValue(v float64) *SportOutdoorsResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetVideoList(v []*SportOutdoorsResponseDataListItemVideoListItem) *SportOutdoorsResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetRank(v int32) *SportOutdoorsResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetRankChange(v string) *SportOutdoorsResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetNickname(v string) *SportOutdoorsResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportOutdoorsResponseDataListItem) SetAvatar(v string) *SportOutdoorsResponseDataListItem {
	s.Avatar = &v
	return s
}

type SportOutdoorsResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s SportOutdoorsResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportOutdoorsResponseDataListItemVideoListItem) SetShareUrl(v string) *SportOutdoorsResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportOutdoorsResponseDataListItemVideoListItem) SetTitle(v string) *SportOutdoorsResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *SportOutdoorsResponseDataListItemVideoListItem) SetItemCover(v string) *SportOutdoorsResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type SportOutdoorsResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s SportOutdoorsResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportOutdoorsResponseExtra) GoString() string {
	return s.String()
}

func (s *SportOutdoorsResponseExtra) SetLogid(v string) *SportOutdoorsResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportOutdoorsResponseExtra) SetNow(v int64) *SportOutdoorsResponseExtra {
	s.Now = &v
	return s
}

func (s *SportOutdoorsResponseExtra) SetErrorCode(v int32) *SportOutdoorsResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportOutdoorsResponseExtra) SetDescription(v string) *SportOutdoorsResponseExtra {
	s.Description = &v
	return s
}

func (s *SportOutdoorsResponseExtra) SetSubErrorCode(v int32) *SportOutdoorsResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportOutdoorsResponseExtra) SetSubDescription(v string) *SportOutdoorsResponseExtra {
	s.SubDescription = &v
	return s
}

type SportOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s SportOverallRequest) GoString() string {
	return s.String()
}

func (s *SportOverallRequest) SetHeader(v map[string]*string) *SportOverallRequest {
	s.Header = v
	return s
}

func (s *SportOverallRequest) SetAccessToken(v string) *SportOverallRequest {
	s.AccessToken = &v
	return s
}

type SportOverallResponse struct {
	Extra *SportOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *SportOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SportOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s SportOverallResponse) GoString() string {
	return s.String()
}

func (s *SportOverallResponse) SetExtra(v *SportOverallResponseExtra) *SportOverallResponse {
	s.Extra = v
	return s
}

func (s *SportOverallResponse) SetData(v *SportOverallResponseData) *SportOverallResponse {
	s.Data = v
	return s
}

type SportOverallResponseData struct {
	List          []*SportOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportOverallResponseData) GoString() string {
	return s.String()
}

func (s *SportOverallResponseData) SetList(v []*SportOverallResponseDataListItem) *SportOverallResponseData {
	s.List = v
	return s
}

func (s *SportOverallResponseData) SetGwErrorCode(v int32) *SportOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportOverallResponseData) SetGwDescription(v string) *SportOverallResponseData {
	s.GwDescription = &v
	return s
}

type SportOverallResponseDataListItem struct {
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s SportOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportOverallResponseDataListItem) SetRankChange(v string) *SportOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetNickname(v string) *SportOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetAvatar(v string) *SportOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetFollowerCount(v int64) *SportOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *SportOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetEffectValue(v float64) *SportOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportOverallResponseDataListItem) SetVideoList(v []*SportOverallResponseDataListItemVideoListItem) *SportOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportOverallResponseDataListItem) SetRank(v int32) *SportOverallResponseDataListItem {
	s.Rank = &v
	return s
}

type SportOverallResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s SportOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *SportOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportOverallResponseDataListItemVideoListItem) SetTitle(v string) *SportOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *SportOverallResponseDataListItemVideoListItem) SetItemCover(v string) *SportOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type SportOverallResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *SportOverallResponseExtra) SetSubErrorCode(v int32) *SportOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportOverallResponseExtra) SetSubDescription(v string) *SportOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportOverallResponseExtra) SetLogid(v string) *SportOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportOverallResponseExtra) SetNow(v int64) *SportOverallResponseExtra {
	s.Now = &v
	return s
}

func (s *SportOverallResponseExtra) SetErrorCode(v int32) *SportOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportOverallResponseExtra) SetDescription(v string) *SportOverallResponseExtra {
	s.Description = &v
	return s
}

type SportSoccerRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportSoccerRequest) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerRequest) GoString() string {
	return s.String()
}

func (s *SportSoccerRequest) SetHeader(v map[string]*string) *SportSoccerRequest {
	s.Header = v
	return s
}

func (s *SportSoccerRequest) SetAccessToken(v string) *SportSoccerRequest {
	s.AccessToken = &v
	return s
}

type SportSoccerResponse struct {
	Extra *SportSoccerResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *SportSoccerResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SportSoccerResponse) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerResponse) GoString() string {
	return s.String()
}

func (s *SportSoccerResponse) SetExtra(v *SportSoccerResponseExtra) *SportSoccerResponse {
	s.Extra = v
	return s
}

func (s *SportSoccerResponse) SetData(v *SportSoccerResponseData) *SportSoccerResponse {
	s.Data = v
	return s
}

type SportSoccerResponseData struct {
	List          []*SportSoccerResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportSoccerResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerResponseData) GoString() string {
	return s.String()
}

func (s *SportSoccerResponseData) SetList(v []*SportSoccerResponseDataListItem) *SportSoccerResponseData {
	s.List = v
	return s
}

func (s *SportSoccerResponseData) SetGwErrorCode(v int32) *SportSoccerResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportSoccerResponseData) SetGwDescription(v string) *SportSoccerResponseData {
	s.GwDescription = &v
	return s
}

type SportSoccerResponseDataListItem struct {
	VideoList        []*SportSoccerResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                          `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                         `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                         `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                         `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                          `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                          `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                        `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
}

func (s SportSoccerResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportSoccerResponseDataListItem) SetVideoList(v []*SportSoccerResponseDataListItemVideoListItem) *SportSoccerResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportSoccerResponseDataListItem) SetRank(v int32) *SportSoccerResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetRankChange(v string) *SportSoccerResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetNickname(v string) *SportSoccerResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetAvatar(v string) *SportSoccerResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetFollowerCount(v int64) *SportSoccerResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetOnbillbaordTimes(v int32) *SportSoccerResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportSoccerResponseDataListItem) SetEffectValue(v float64) *SportSoccerResponseDataListItem {
	s.EffectValue = &v
	return s
}

type SportSoccerResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SportSoccerResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportSoccerResponseDataListItemVideoListItem) SetItemCover(v string) *SportSoccerResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *SportSoccerResponseDataListItemVideoListItem) SetShareUrl(v string) *SportSoccerResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportSoccerResponseDataListItemVideoListItem) SetTitle(v string) *SportSoccerResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type SportSoccerResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportSoccerResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportSoccerResponseExtra) GoString() string {
	return s.String()
}

func (s *SportSoccerResponseExtra) SetSubErrorCode(v int32) *SportSoccerResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportSoccerResponseExtra) SetSubDescription(v string) *SportSoccerResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportSoccerResponseExtra) SetLogid(v string) *SportSoccerResponseExtra {
	s.Logid = &v
	return s
}

func (s *SportSoccerResponseExtra) SetNow(v int64) *SportSoccerResponseExtra {
	s.Now = &v
	return s
}

func (s *SportSoccerResponseExtra) SetErrorCode(v int32) *SportSoccerResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportSoccerResponseExtra) SetDescription(v string) *SportSoccerResponseExtra {
	s.Description = &v
	return s
}

type SportTableTennisRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SportTableTennisRequest) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisRequest) GoString() string {
	return s.String()
}

func (s *SportTableTennisRequest) SetHeader(v map[string]*string) *SportTableTennisRequest {
	s.Header = v
	return s
}

func (s *SportTableTennisRequest) SetAccessToken(v string) *SportTableTennisRequest {
	s.AccessToken = &v
	return s
}

type SportTableTennisResponse struct {
	Extra *SportTableTennisResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *SportTableTennisResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SportTableTennisResponse) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisResponse) GoString() string {
	return s.String()
}

func (s *SportTableTennisResponse) SetExtra(v *SportTableTennisResponseExtra) *SportTableTennisResponse {
	s.Extra = v
	return s
}

func (s *SportTableTennisResponse) SetData(v *SportTableTennisResponseData) *SportTableTennisResponse {
	s.Data = v
	return s
}

type SportTableTennisResponseData struct {
	List          []*SportTableTennisResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SportTableTennisResponseData) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisResponseData) GoString() string {
	return s.String()
}

func (s *SportTableTennisResponseData) SetList(v []*SportTableTennisResponseDataListItem) *SportTableTennisResponseData {
	s.List = v
	return s
}

func (s *SportTableTennisResponseData) SetGwErrorCode(v int32) *SportTableTennisResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *SportTableTennisResponseData) SetGwDescription(v string) *SportTableTennisResponseData {
	s.GwDescription = &v
	return s
}

type SportTableTennisResponseDataListItem struct {
	FollowerCount    *int64                                               `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                               `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                             `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*SportTableTennisResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                               `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                              `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                              `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                              `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s SportTableTennisResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisResponseDataListItem) GoString() string {
	return s.String()
}

func (s *SportTableTennisResponseDataListItem) SetFollowerCount(v int64) *SportTableTennisResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetOnbillbaordTimes(v int32) *SportTableTennisResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetEffectValue(v float64) *SportTableTennisResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetVideoList(v []*SportTableTennisResponseDataListItemVideoListItem) *SportTableTennisResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetRank(v int32) *SportTableTennisResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetRankChange(v string) *SportTableTennisResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetNickname(v string) *SportTableTennisResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *SportTableTennisResponseDataListItem) SetAvatar(v string) *SportTableTennisResponseDataListItem {
	s.Avatar = &v
	return s
}

type SportTableTennisResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SportTableTennisResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *SportTableTennisResponseDataListItemVideoListItem) SetItemCover(v string) *SportTableTennisResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *SportTableTennisResponseDataListItemVideoListItem) SetShareUrl(v string) *SportTableTennisResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *SportTableTennisResponseDataListItemVideoListItem) SetTitle(v string) *SportTableTennisResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type SportTableTennisResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s SportTableTennisResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SportTableTennisResponseExtra) GoString() string {
	return s.String()
}

func (s *SportTableTennisResponseExtra) SetNow(v int64) *SportTableTennisResponseExtra {
	s.Now = &v
	return s
}

func (s *SportTableTennisResponseExtra) SetErrorCode(v int32) *SportTableTennisResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SportTableTennisResponseExtra) SetDescription(v string) *SportTableTennisResponseExtra {
	s.Description = &v
	return s
}

func (s *SportTableTennisResponseExtra) SetSubErrorCode(v int32) *SportTableTennisResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SportTableTennisResponseExtra) SetSubDescription(v string) *SportTableTennisResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SportTableTennisResponseExtra) SetLogid(v string) *SportTableTennisResponseExtra {
	s.Logid = &v
	return s
}

type SpuDetailRequest struct {
	AccessToken *string               `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OutSpuId    *string               `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpuId       *int64                `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Base        *SpuDetailRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *int64                `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string    `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SpuDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailRequest) GoString() string {
	return s.String()
}

func (s *SpuDetailRequest) SetAccessToken(v string) *SpuDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *SpuDetailRequest) SetOutSpuId(v string) *SpuDetailRequest {
	s.OutSpuId = &v
	return s
}

func (s *SpuDetailRequest) SetSpuId(v int64) *SpuDetailRequest {
	s.SpuId = &v
	return s
}

func (s *SpuDetailRequest) SetBase(v *SpuDetailRequestBase) *SpuDetailRequest {
	s.Base = v
	return s
}

func (s *SpuDetailRequest) SetAccountId(v int64) *SpuDetailRequest {
	s.AccountId = &v
	return s
}

func (s *SpuDetailRequest) SetHeader(v map[string]*string) *SpuDetailRequest {
	s.Header = v
	return s
}

type SpuDetailRequestBase struct {
	Client     *string                         `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string              `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                         `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *SpuDetailRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                         `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                         `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s SpuDetailRequestBase) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailRequestBase) GoString() string {
	return s.String()
}

func (s *SpuDetailRequestBase) SetClient(v string) *SpuDetailRequestBase {
	s.Client = &v
	return s
}

func (s *SpuDetailRequestBase) SetExtra(v map[string]*string) *SpuDetailRequestBase {
	s.Extra = v
	return s
}

func (s *SpuDetailRequestBase) SetLogID(v string) *SpuDetailRequestBase {
	s.LogID = &v
	return s
}

func (s *SpuDetailRequestBase) SetTrafficEnv(v *SpuDetailRequestBaseTrafficEnv) *SpuDetailRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *SpuDetailRequestBase) SetAddr(v string) *SpuDetailRequestBase {
	s.Addr = &v
	return s
}

func (s *SpuDetailRequestBase) SetCaller(v string) *SpuDetailRequestBase {
	s.Caller = &v
	return s
}

type SpuDetailRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s SpuDetailRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *SpuDetailRequestBaseTrafficEnv) SetOpen(v bool) *SpuDetailRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *SpuDetailRequestBaseTrafficEnv) SetEnv(v string) *SpuDetailRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type SpuDetailResponse struct {
	Extra    *SpuDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *SpuDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *SpuDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SpuDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponse) GoString() string {
	return s.String()
}

func (s *SpuDetailResponse) SetExtra(v *SpuDetailResponseExtra) *SpuDetailResponse {
	s.Extra = v
	return s
}

func (s *SpuDetailResponse) SetBaseResp(v *SpuDetailResponseBaseResp) *SpuDetailResponse {
	s.BaseResp = v
	return s
}

func (s *SpuDetailResponse) SetData(v *SpuDetailResponseData) *SpuDetailResponse {
	s.Data = v
	return s
}

type SpuDetailResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SpuDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseBaseResp) SetStatusMessage(v string) *SpuDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *SpuDetailResponseBaseResp) SetExtra(v map[string]*string) *SpuDetailResponseBaseResp {
	s.Extra = v
	return s
}

func (s *SpuDetailResponseBaseResp) SetStatusCode(v int32) *SpuDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

type SpuDetailResponseData struct {
	Description *string                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SpuOnline   *SpuDetailResponseDataSpuOnline `json:"spu_online,omitempty" xml:"spu_online,omitempty"`
}

func (s SpuDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseData) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseData) SetDescription(v string) *SpuDetailResponseData {
	s.Description = &v
	return s
}

func (s *SpuDetailResponseData) SetErrorCode(v int32) *SpuDetailResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SpuDetailResponseData) SetSpuOnline(v *SpuDetailResponseDataSpuOnline) *SpuDetailResponseData {
	s.SpuOnline = v
	return s
}

type SpuDetailResponseDataSpuOnline struct {
	Spu    *SpuDetailResponseDataSpuOnlineSpu `json:"spu,omitempty" xml:"spu,omitempty"`
	Status *int                               `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SpuDetailResponseDataSpuOnline) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnline) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnline) SetSpu(v *SpuDetailResponseDataSpuOnlineSpu) *SpuDetailResponseDataSpuOnline {
	s.Spu = v
	return s
}

func (s *SpuDetailResponseDataSpuOnline) SetStatus(v int) *SpuDetailResponseDataSpuOnline {
	s.Status = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpu struct {
	BizLine        *int                                                   `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	SpuName        *string                                                `json:"spu_name,omitempty" xml:"spu_name,omitempty"`
	Description    *string                                                `json:"description,omitempty" xml:"description,omitempty"`
	AccountId      *int64                                                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	SpuId          *int64                                                 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	OutSpuId       *string                                                `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	CategoryId     *int64                                                 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	SaleAttrGroups []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem `json:"sale_attr_groups,omitempty" xml:"sale_attr_groups,omitempty" type:"Repeated"`
	Images         *SpuDetailResponseDataSpuOnlineSpuImages               `json:"images,omitempty" xml:"images,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpu) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpu) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetBizLine(v int) *SpuDetailResponseDataSpuOnlineSpu {
	s.BizLine = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetSpuName(v string) *SpuDetailResponseDataSpuOnlineSpu {
	s.SpuName = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetDescription(v string) *SpuDetailResponseDataSpuOnlineSpu {
	s.Description = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetAccountId(v int64) *SpuDetailResponseDataSpuOnlineSpu {
	s.AccountId = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetSpuId(v int64) *SpuDetailResponseDataSpuOnlineSpu {
	s.SpuId = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetOutSpuId(v string) *SpuDetailResponseDataSpuOnlineSpu {
	s.OutSpuId = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetCategoryId(v int64) *SpuDetailResponseDataSpuOnlineSpu {
	s.CategoryId = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetSaleAttrGroups(v []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) *SpuDetailResponseDataSpuOnlineSpu {
	s.SaleAttrGroups = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpu) SetImages(v *SpuDetailResponseDataSpuOnlineSpuImages) *SpuDetailResponseDataSpuOnlineSpu {
	s.Images = v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuImages struct {
	DescriptionImage *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage `json:"description_image,omitempty" xml:"description_image,omitempty"`
	HeadImage        *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage        `json:"head_image,omitempty" xml:"head_image,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuImages) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuImages) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuImages) SetDescriptionImage(v *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage) *SpuDetailResponseDataSpuOnlineSpuImages {
	s.DescriptionImage = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuImages) SetHeadImage(v *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage) *SpuDetailResponseDataSpuOnlineSpuImages {
	s.HeadImage = v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage struct {
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage) SetUri(v string) *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage {
	s.Uri = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage) SetUrl(v string) *SpuDetailResponseDataSpuOnlineSpuImagesDescriptionImage {
	s.Url = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuImagesHeadImage struct {
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuImagesHeadImage) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuImagesHeadImage) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage) SetUri(v string) *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage {
	s.Uri = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage) SetUrl(v string) *SpuDetailResponseDataSpuOnlineSpuImagesHeadImage {
	s.Url = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem struct {
	ItemList  []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupCode *string                                                            `json:"group_code,omitempty" xml:"group_code,omitempty"`
	GroupName *string                                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) SetItemList(v []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem {
	s.ItemList = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) SetGroupCode(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem {
	s.GroupCode = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem) SetGroupName(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItem {
	s.GroupName = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem struct {
	Image     *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage          `json:"image,omitempty" xml:"image,omitempty"`
	IsDefault *bool                                                                          `json:"is_default,omitempty" xml:"is_default,omitempty"`
	ItemKey   *string                                                                        `json:"item_key,omitempty" xml:"item_key,omitempty"`
	ItemName  *string                                                                        `json:"item_name,omitempty" xml:"item_name,omitempty"`
	AddPrice  *int32                                                                         `json:"add_price,omitempty" xml:"add_price,omitempty"`
	DescList  []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem `json:"desc_list,omitempty" xml:"desc_list,omitempty" type:"Repeated"`
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetImage(v *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.Image = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetIsDefault(v bool) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.IsDefault = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetItemKey(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.ItemKey = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetItemName(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetAddPrice(v int32) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.AddPrice = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem) SetDescList(v []*SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItem {
	s.DescList = v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem struct {
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	Val *string `json:"val,omitempty" xml:"val,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem) SetKey(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem {
	s.Key = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem) SetVal(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemDescListItem {
	s.Val = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage struct {
	ImageTag  []*string                                                                     `json:"image_tag,omitempty" xml:"image_tag,omitempty" type:"Repeated"`
	ImageType *int                                                                          `json:"image_type,omitempty" xml:"image_type,omitempty"`
	ItemInfo  *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo `json:"item_info,omitempty" xml:"item_info,omitempty"`
	Name      *string                                                                       `json:"name,omitempty" xml:"name,omitempty"`
	Uri       *string                                                                       `json:"uri,omitempty" xml:"uri,omitempty" require:"true"`
	Url       *string                                                                       `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetImageTag(v []*string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.ImageTag = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetImageType(v int) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.ImageType = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetItemInfo(v *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.ItemInfo = v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetName(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.Name = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetUri(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.Uri = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage) SetUrl(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImage {
	s.Url = &v
	return s
}

type SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo struct {
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo) SetItemId(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo {
	s.ItemId = &v
	return s
}

func (s *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo) SetItemName(v string) *SpuDetailResponseDataSpuOnlineSpuSaleAttrGroupsItemItemListItemImageItemInfo {
	s.ItemName = &v
	return s
}

type SpuDetailResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s SpuDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SpuDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *SpuDetailResponseExtra) SetNow(v int64) *SpuDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *SpuDetailResponseExtra) SetSubDescription(v string) *SpuDetailResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SpuDetailResponseExtra) SetSubErrorCode(v int32) *SpuDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SpuDetailResponseExtra) SetDescription(v string) *SpuDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *SpuDetailResponseExtra) SetErrorCode(v int32) *SpuDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SpuDetailResponseExtra) SetLogid(v string) *SpuDetailResponseExtra {
	s.Logid = &v
	return s
}

type SpuGetRequest struct {
	AccountId   *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OutSpuId    *string            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpuId       *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SpuGetRequest) String() string {
	return tea.Prettify(s)
}

func (s SpuGetRequest) GoString() string {
	return s.String()
}

func (s *SpuGetRequest) SetAccountId(v int64) *SpuGetRequest {
	s.AccountId = &v
	return s
}

func (s *SpuGetRequest) SetOutSpuId(v string) *SpuGetRequest {
	s.OutSpuId = &v
	return s
}

func (s *SpuGetRequest) SetSpuId(v int64) *SpuGetRequest {
	s.SpuId = &v
	return s
}

func (s *SpuGetRequest) SetHeader(v map[string]*string) *SpuGetRequest {
	s.Header = v
	return s
}

func (s *SpuGetRequest) SetAccessToken(v string) *SpuGetRequest {
	s.AccessToken = &v
	return s
}

type SpuGetResponse struct {
	Data  *SpuGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *SpuGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SpuGetResponse) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponse) GoString() string {
	return s.String()
}

func (s *SpuGetResponse) SetData(v *SpuGetResponseData) *SpuGetResponse {
	s.Data = v
	return s
}

func (s *SpuGetResponse) SetExtra(v *SpuGetResponseExtra) *SpuGetResponse {
	s.Extra = v
	return s
}

type SpuGetResponseData struct {
	SpuDraft    *SpuGetResponseDataSpuDraft  `json:"spu_draft,omitempty" xml:"spu_draft,omitempty"`
	SpuOnline   *SpuGetResponseDataSpuOnline `json:"spu_online,omitempty" xml:"spu_online,omitempty"`
	Description *string                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s SpuGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseData) GoString() string {
	return s.String()
}

func (s *SpuGetResponseData) SetSpuDraft(v *SpuGetResponseDataSpuDraft) *SpuGetResponseData {
	s.SpuDraft = v
	return s
}

func (s *SpuGetResponseData) SetSpuOnline(v *SpuGetResponseDataSpuOnline) *SpuGetResponseData {
	s.SpuOnline = v
	return s
}

func (s *SpuGetResponseData) SetDescription(v string) *SpuGetResponseData {
	s.Description = &v
	return s
}

func (s *SpuGetResponseData) SetErrorCode(v int32) *SpuGetResponseData {
	s.ErrorCode = &v
	return s
}

type SpuGetResponseDataSpuDraft struct {
	Spu         *SpuGetResponseDataSpuDraftSpu       `json:"spu,omitempty" xml:"spu,omitempty"`
	AuditInfo   *SpuGetResponseDataSpuDraftAuditInfo `json:"audit_info,omitempty" xml:"audit_info,omitempty"`
	DraftStatus *int                                 `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
}

func (s SpuGetResponseDataSpuDraft) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseDataSpuDraft) GoString() string {
	return s.String()
}

func (s *SpuGetResponseDataSpuDraft) SetSpu(v *SpuGetResponseDataSpuDraftSpu) *SpuGetResponseDataSpuDraft {
	s.Spu = v
	return s
}

func (s *SpuGetResponseDataSpuDraft) SetAuditInfo(v *SpuGetResponseDataSpuDraftAuditInfo) *SpuGetResponseDataSpuDraft {
	s.AuditInfo = v
	return s
}

func (s *SpuGetResponseDataSpuDraft) SetDraftStatus(v int) *SpuGetResponseDataSpuDraft {
	s.DraftStatus = &v
	return s
}

type SpuGetResponseDataSpuDraftAuditInfo struct {
	EndTime   *int64  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
	StartTime *int64  `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s SpuGetResponseDataSpuDraftAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseDataSpuDraftAuditInfo) GoString() string {
	return s.String()
}

func (s *SpuGetResponseDataSpuDraftAuditInfo) SetEndTime(v int64) *SpuGetResponseDataSpuDraftAuditInfo {
	s.EndTime = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftAuditInfo) SetReason(v string) *SpuGetResponseDataSpuDraftAuditInfo {
	s.Reason = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftAuditInfo) SetStartTime(v int64) *SpuGetResponseDataSpuDraftAuditInfo {
	s.StartTime = &v
	return s
}

type SpuGetResponseDataSpuDraftSpu struct {
	StandardCategoryId *int64             `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty"`
	UpdateTime         *int64             `json:"update_time,omitempty" xml:"update_time,omitempty"`
	AccountId          *int64             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	CreateTime         *int64             `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OutSpuId           *string            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpuAttrMap         map[string]*string `json:"spu_attr_map,omitempty" xml:"spu_attr_map,omitempty"`
	SpuId              *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SpuName            *string            `json:"spu_name,omitempty" xml:"spu_name,omitempty"`
}

func (s SpuGetResponseDataSpuDraftSpu) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseDataSpuDraftSpu) GoString() string {
	return s.String()
}

func (s *SpuGetResponseDataSpuDraftSpu) SetStandardCategoryId(v int64) *SpuGetResponseDataSpuDraftSpu {
	s.StandardCategoryId = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetUpdateTime(v int64) *SpuGetResponseDataSpuDraftSpu {
	s.UpdateTime = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetAccountId(v int64) *SpuGetResponseDataSpuDraftSpu {
	s.AccountId = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetCreateTime(v int64) *SpuGetResponseDataSpuDraftSpu {
	s.CreateTime = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetOutSpuId(v string) *SpuGetResponseDataSpuDraftSpu {
	s.OutSpuId = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetSpuAttrMap(v map[string]*string) *SpuGetResponseDataSpuDraftSpu {
	s.SpuAttrMap = v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetSpuId(v int64) *SpuGetResponseDataSpuDraftSpu {
	s.SpuId = &v
	return s
}

func (s *SpuGetResponseDataSpuDraftSpu) SetSpuName(v string) *SpuGetResponseDataSpuDraftSpu {
	s.SpuName = &v
	return s
}

type SpuGetResponseDataSpuOnline struct {
	OnlineStatus *int                            `json:"online_status,omitempty" xml:"online_status,omitempty"`
	Spu          *SpuGetResponseDataSpuOnlineSpu `json:"spu,omitempty" xml:"spu,omitempty"`
}

func (s SpuGetResponseDataSpuOnline) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseDataSpuOnline) GoString() string {
	return s.String()
}

func (s *SpuGetResponseDataSpuOnline) SetOnlineStatus(v int) *SpuGetResponseDataSpuOnline {
	s.OnlineStatus = &v
	return s
}

func (s *SpuGetResponseDataSpuOnline) SetSpu(v *SpuGetResponseDataSpuOnlineSpu) *SpuGetResponseDataSpuOnline {
	s.Spu = v
	return s
}

type SpuGetResponseDataSpuOnlineSpu struct {
	UpdateTime         *int64             `json:"update_time,omitempty" xml:"update_time,omitempty"`
	AccountId          *int64             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	CreateTime         *int64             `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OutSpuId           *string            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpuAttrMap         map[string]*string `json:"spu_attr_map,omitempty" xml:"spu_attr_map,omitempty"`
	SpuId              *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SpuName            *string            `json:"spu_name,omitempty" xml:"spu_name,omitempty"`
	StandardCategoryId *int64             `json:"standard_category_id,omitempty" xml:"standard_category_id,omitempty"`
}

func (s SpuGetResponseDataSpuOnlineSpu) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseDataSpuOnlineSpu) GoString() string {
	return s.String()
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetUpdateTime(v int64) *SpuGetResponseDataSpuOnlineSpu {
	s.UpdateTime = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetAccountId(v int64) *SpuGetResponseDataSpuOnlineSpu {
	s.AccountId = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetCreateTime(v int64) *SpuGetResponseDataSpuOnlineSpu {
	s.CreateTime = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetOutSpuId(v string) *SpuGetResponseDataSpuOnlineSpu {
	s.OutSpuId = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetSpuAttrMap(v map[string]*string) *SpuGetResponseDataSpuOnlineSpu {
	s.SpuAttrMap = v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetSpuId(v int64) *SpuGetResponseDataSpuOnlineSpu {
	s.SpuId = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetSpuName(v string) *SpuGetResponseDataSpuOnlineSpu {
	s.SpuName = &v
	return s
}

func (s *SpuGetResponseDataSpuOnlineSpu) SetStandardCategoryId(v int64) *SpuGetResponseDataSpuOnlineSpu {
	s.StandardCategoryId = &v
	return s
}

type SpuGetResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SpuGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SpuGetResponseExtra) GoString() string {
	return s.String()
}

func (s *SpuGetResponseExtra) SetErrorCode(v int32) *SpuGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SpuGetResponseExtra) SetLogid(v string) *SpuGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *SpuGetResponseExtra) SetNow(v int64) *SpuGetResponseExtra {
	s.Now = &v
	return s
}

func (s *SpuGetResponseExtra) SetSubDescription(v string) *SpuGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SpuGetResponseExtra) SetSubErrorCode(v int32) *SpuGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SpuGetResponseExtra) SetDescription(v string) *SpuGetResponseExtra {
	s.Description = &v
	return s
}

type SpuOperateRequest struct {
	SpuId       *int64                 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Base        *SpuOperateRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *int64                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Op          *int                   `json:"op,omitempty" xml:"op,omitempty" require:"true"`
	OutSpuId    *string                `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
}

func (s SpuOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateRequest) GoString() string {
	return s.String()
}

func (s *SpuOperateRequest) SetSpuId(v int64) *SpuOperateRequest {
	s.SpuId = &v
	return s
}

func (s *SpuOperateRequest) SetBase(v *SpuOperateRequestBase) *SpuOperateRequest {
	s.Base = v
	return s
}

func (s *SpuOperateRequest) SetAccountId(v int64) *SpuOperateRequest {
	s.AccountId = &v
	return s
}

func (s *SpuOperateRequest) SetHeader(v map[string]*string) *SpuOperateRequest {
	s.Header = v
	return s
}

func (s *SpuOperateRequest) SetAccessToken(v string) *SpuOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *SpuOperateRequest) SetOp(v int) *SpuOperateRequest {
	s.Op = &v
	return s
}

func (s *SpuOperateRequest) SetOutSpuId(v string) *SpuOperateRequest {
	s.OutSpuId = &v
	return s
}

type SpuOperateRequestBase struct {
	Caller     *string                          `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                          `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string               `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                          `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *SpuOperateRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                          `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s SpuOperateRequestBase) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateRequestBase) GoString() string {
	return s.String()
}

func (s *SpuOperateRequestBase) SetCaller(v string) *SpuOperateRequestBase {
	s.Caller = &v
	return s
}

func (s *SpuOperateRequestBase) SetClient(v string) *SpuOperateRequestBase {
	s.Client = &v
	return s
}

func (s *SpuOperateRequestBase) SetExtra(v map[string]*string) *SpuOperateRequestBase {
	s.Extra = v
	return s
}

func (s *SpuOperateRequestBase) SetLogID(v string) *SpuOperateRequestBase {
	s.LogID = &v
	return s
}

func (s *SpuOperateRequestBase) SetTrafficEnv(v *SpuOperateRequestBaseTrafficEnv) *SpuOperateRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *SpuOperateRequestBase) SetAddr(v string) *SpuOperateRequestBase {
	s.Addr = &v
	return s
}

type SpuOperateRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s SpuOperateRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *SpuOperateRequestBaseTrafficEnv) SetOpen(v bool) *SpuOperateRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *SpuOperateRequestBaseTrafficEnv) SetEnv(v string) *SpuOperateRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type SpuOperateResponse struct {
	BaseResp *SpuOperateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *SpuOperateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *SpuOperateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s SpuOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateResponse) GoString() string {
	return s.String()
}

func (s *SpuOperateResponse) SetBaseResp(v *SpuOperateResponseBaseResp) *SpuOperateResponse {
	s.BaseResp = v
	return s
}

func (s *SpuOperateResponse) SetData(v *SpuOperateResponseData) *SpuOperateResponse {
	s.Data = v
	return s
}

func (s *SpuOperateResponse) SetExtra(v *SpuOperateResponseExtra) *SpuOperateResponse {
	s.Extra = v
	return s
}

type SpuOperateResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s SpuOperateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *SpuOperateResponseBaseResp) SetExtra(v map[string]*string) *SpuOperateResponseBaseResp {
	s.Extra = v
	return s
}

func (s *SpuOperateResponseBaseResp) SetStatusCode(v int32) *SpuOperateResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *SpuOperateResponseBaseResp) SetStatusMessage(v string) *SpuOperateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type SpuOperateResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SpuId       *int64  `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
}

func (s SpuOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateResponseData) GoString() string {
	return s.String()
}

func (s *SpuOperateResponseData) SetDescription(v string) *SpuOperateResponseData {
	s.Description = &v
	return s
}

func (s *SpuOperateResponseData) SetErrorCode(v int32) *SpuOperateResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SpuOperateResponseData) SetSpuId(v int64) *SpuOperateResponseData {
	s.SpuId = &v
	return s
}

type SpuOperateResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SpuOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SpuOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *SpuOperateResponseExtra) SetErrorCode(v int32) *SpuOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SpuOperateResponseExtra) SetLogid(v string) *SpuOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *SpuOperateResponseExtra) SetNow(v int64) *SpuOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *SpuOperateResponseExtra) SetSubDescription(v string) *SpuOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SpuOperateResponseExtra) SetSubErrorCode(v int32) *SpuOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SpuOperateResponseExtra) SetDescription(v string) *SpuOperateResponseExtra {
	s.Description = &v
	return s
}

type SpuSaveRequest struct {
	AccountId   *int64              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Spu         *SpuSaveRequestSpu  `json:"spu,omitempty" xml:"spu,omitempty" require:"true"`
	Base        *SpuSaveRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SpuSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequest) GoString() string {
	return s.String()
}

func (s *SpuSaveRequest) SetAccountId(v int64) *SpuSaveRequest {
	s.AccountId = &v
	return s
}

func (s *SpuSaveRequest) SetSpu(v *SpuSaveRequestSpu) *SpuSaveRequest {
	s.Spu = v
	return s
}

func (s *SpuSaveRequest) SetBase(v *SpuSaveRequestBase) *SpuSaveRequest {
	s.Base = v
	return s
}

func (s *SpuSaveRequest) SetHeader(v map[string]*string) *SpuSaveRequest {
	s.Header = v
	return s
}

func (s *SpuSaveRequest) SetAccessToken(v string) *SpuSaveRequest {
	s.AccessToken = &v
	return s
}

type SpuSaveRequestBase struct {
	Extra      map[string]*string            `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                       `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *SpuSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                       `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                       `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                       `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s SpuSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequestBase) GoString() string {
	return s.String()
}

func (s *SpuSaveRequestBase) SetExtra(v map[string]*string) *SpuSaveRequestBase {
	s.Extra = v
	return s
}

func (s *SpuSaveRequestBase) SetLogID(v string) *SpuSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *SpuSaveRequestBase) SetTrafficEnv(v *SpuSaveRequestBaseTrafficEnv) *SpuSaveRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *SpuSaveRequestBase) SetAddr(v string) *SpuSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *SpuSaveRequestBase) SetCaller(v string) *SpuSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *SpuSaveRequestBase) SetClient(v string) *SpuSaveRequestBase {
	s.Client = &v
	return s
}

type SpuSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s SpuSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *SpuSaveRequestBaseTrafficEnv) SetEnv(v string) *SpuSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *SpuSaveRequestBaseTrafficEnv) SetOpen(v bool) *SpuSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type SpuSaveRequestSpu struct {
	OutSpuId       *string                                `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpuName        *string                                `json:"spu_name,omitempty" xml:"spu_name,omitempty"`
	SaleAttrGroups []*SpuSaveRequestSpuSaleAttrGroupsItem `json:"sale_attr_groups,omitempty" xml:"sale_attr_groups,omitempty" type:"Repeated"`
	AccountId      *int64                                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BizLine        *int                                   `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	CategoryId     *int64                                 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

func (s SpuSaveRequestSpu) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequestSpu) GoString() string {
	return s.String()
}

func (s *SpuSaveRequestSpu) SetOutSpuId(v string) *SpuSaveRequestSpu {
	s.OutSpuId = &v
	return s
}

func (s *SpuSaveRequestSpu) SetSpuName(v string) *SpuSaveRequestSpu {
	s.SpuName = &v
	return s
}

func (s *SpuSaveRequestSpu) SetSaleAttrGroups(v []*SpuSaveRequestSpuSaleAttrGroupsItem) *SpuSaveRequestSpu {
	s.SaleAttrGroups = v
	return s
}

func (s *SpuSaveRequestSpu) SetAccountId(v int64) *SpuSaveRequestSpu {
	s.AccountId = &v
	return s
}

func (s *SpuSaveRequestSpu) SetBizLine(v int) *SpuSaveRequestSpu {
	s.BizLine = &v
	return s
}

func (s *SpuSaveRequestSpu) SetCategoryId(v int64) *SpuSaveRequestSpu {
	s.CategoryId = &v
	return s
}

type SpuSaveRequestSpuSaleAttrGroupsItem struct {
	GroupCode *string                                            `json:"group_code,omitempty" xml:"group_code,omitempty"`
	GroupName *string                                            `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*SpuSaveRequestSpuSaleAttrGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s SpuSaveRequestSpuSaleAttrGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequestSpuSaleAttrGroupsItem) GoString() string {
	return s.String()
}

func (s *SpuSaveRequestSpuSaleAttrGroupsItem) SetGroupCode(v string) *SpuSaveRequestSpuSaleAttrGroupsItem {
	s.GroupCode = &v
	return s
}

func (s *SpuSaveRequestSpuSaleAttrGroupsItem) SetGroupName(v string) *SpuSaveRequestSpuSaleAttrGroupsItem {
	s.GroupName = &v
	return s
}

func (s *SpuSaveRequestSpuSaleAttrGroupsItem) SetItemList(v []*SpuSaveRequestSpuSaleAttrGroupsItemItemListItem) *SpuSaveRequestSpuSaleAttrGroupsItem {
	s.ItemList = v
	return s
}

type SpuSaveRequestSpuSaleAttrGroupsItemItemListItem struct {
	ItemKey  *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}

func (s SpuSaveRequestSpuSaleAttrGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveRequestSpuSaleAttrGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *SpuSaveRequestSpuSaleAttrGroupsItemItemListItem) SetItemKey(v string) *SpuSaveRequestSpuSaleAttrGroupsItemItemListItem {
	s.ItemKey = &v
	return s
}

func (s *SpuSaveRequestSpuSaleAttrGroupsItemItemListItem) SetItemName(v string) *SpuSaveRequestSpuSaleAttrGroupsItemItemListItem {
	s.ItemName = &v
	return s
}

type SpuSaveResponse struct {
	Extra    *SpuSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *SpuSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *SpuSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SpuSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveResponse) GoString() string {
	return s.String()
}

func (s *SpuSaveResponse) SetExtra(v *SpuSaveResponseExtra) *SpuSaveResponse {
	s.Extra = v
	return s
}

func (s *SpuSaveResponse) SetBaseResp(v *SpuSaveResponseBaseResp) *SpuSaveResponse {
	s.BaseResp = v
	return s
}

func (s *SpuSaveResponse) SetData(v *SpuSaveResponseData) *SpuSaveResponse {
	s.Data = v
	return s
}

type SpuSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s SpuSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *SpuSaveResponseBaseResp) SetStatusCode(v int32) *SpuSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *SpuSaveResponseBaseResp) SetStatusMessage(v string) *SpuSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *SpuSaveResponseBaseResp) SetExtra(v map[string]*string) *SpuSaveResponseBaseResp {
	s.Extra = v
	return s
}

type SpuSaveResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SpuId       *int64  `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s SpuSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveResponseData) GoString() string {
	return s.String()
}

func (s *SpuSaveResponseData) SetErrorCode(v int32) *SpuSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *SpuSaveResponseData) SetSpuId(v int64) *SpuSaveResponseData {
	s.SpuId = &v
	return s
}

func (s *SpuSaveResponseData) SetDescription(v string) *SpuSaveResponseData {
	s.Description = &v
	return s
}

type SpuSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s SpuSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SpuSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *SpuSaveResponseExtra) SetSubErrorCode(v int32) *SpuSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SpuSaveResponseExtra) SetDescription(v string) *SpuSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *SpuSaveResponseExtra) SetErrorCode(v int32) *SpuSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *SpuSaveResponseExtra) SetLogid(v string) *SpuSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *SpuSaveResponseExtra) SetNow(v int64) *SpuSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *SpuSaveResponseExtra) SetSubDescription(v string) *SpuSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type StatusQueryRequest struct {
	AccountId      *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	HotelIdList    []*string          `json:"hotel_id_list,omitempty" xml:"hotel_id_list,omitempty" type:"Repeated"`
	OutHotelIdList []*string          `json:"out_hotel_id_list,omitempty" xml:"out_hotel_id_list,omitempty" type:"Repeated"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s StatusQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s StatusQueryRequest) GoString() string {
	return s.String()
}

func (s *StatusQueryRequest) SetAccountId(v string) *StatusQueryRequest {
	s.AccountId = &v
	return s
}

func (s *StatusQueryRequest) SetHotelIdList(v []*string) *StatusQueryRequest {
	s.HotelIdList = v
	return s
}

func (s *StatusQueryRequest) SetOutHotelIdList(v []*string) *StatusQueryRequest {
	s.OutHotelIdList = v
	return s
}

func (s *StatusQueryRequest) SetHeader(v map[string]*string) *StatusQueryRequest {
	s.Header = v
	return s
}

func (s *StatusQueryRequest) SetAccessToken(v string) *StatusQueryRequest {
	s.AccessToken = &v
	return s
}

type StatusQueryResponse struct {
	Data  []*StatusQueryResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Extra *StatusQueryResponseExtra      `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s StatusQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s StatusQueryResponse) GoString() string {
	return s.String()
}

func (s *StatusQueryResponse) SetData(v []*StatusQueryResponseDataItem) *StatusQueryResponse {
	s.Data = v
	return s
}

func (s *StatusQueryResponse) SetExtra(v *StatusQueryResponseExtra) *StatusQueryResponse {
	s.Extra = v
	return s
}

type StatusQueryResponseDataItem struct {
	HotelId    *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	OutHotelId *string `json:"out_hotel_id,omitempty" xml:"out_hotel_id,omitempty" require:"true"`
	Status     *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CategoryId *int64  `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

func (s StatusQueryResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s StatusQueryResponseDataItem) GoString() string {
	return s.String()
}

func (s *StatusQueryResponseDataItem) SetHotelId(v string) *StatusQueryResponseDataItem {
	s.HotelId = &v
	return s
}

func (s *StatusQueryResponseDataItem) SetMessage(v string) *StatusQueryResponseDataItem {
	s.Message = &v
	return s
}

func (s *StatusQueryResponseDataItem) SetOutHotelId(v string) *StatusQueryResponseDataItem {
	s.OutHotelId = &v
	return s
}

func (s *StatusQueryResponseDataItem) SetStatus(v int) *StatusQueryResponseDataItem {
	s.Status = &v
	return s
}

func (s *StatusQueryResponseDataItem) SetCategoryId(v int64) *StatusQueryResponseDataItem {
	s.CategoryId = &v
	return s
}

type StatusQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s StatusQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s StatusQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *StatusQueryResponseExtra) SetSubDescription(v string) *StatusQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *StatusQueryResponseExtra) SetSubErrorCode(v int32) *StatusQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *StatusQueryResponseExtra) SetDescription(v string) *StatusQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *StatusQueryResponseExtra) SetErrorCode(v int32) *StatusQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *StatusQueryResponseExtra) SetLogid(v string) *StatusQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *StatusQueryResponseExtra) SetNow(v int64) *StatusQueryResponseExtra {
	s.Now = &v
	return s
}

type StockSaveRequest struct {
	AccountId   *string                     `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Aris        []*StockSaveRequestArisItem `json:"aris,omitempty" xml:"aris,omitempty" require:"true" type:"Repeated"`
	HotelId     *string                     `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Header      map[string]*string          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s StockSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s StockSaveRequest) GoString() string {
	return s.String()
}

func (s *StockSaveRequest) SetAccountId(v string) *StockSaveRequest {
	s.AccountId = &v
	return s
}

func (s *StockSaveRequest) SetAris(v []*StockSaveRequestArisItem) *StockSaveRequest {
	s.Aris = v
	return s
}

func (s *StockSaveRequest) SetHotelId(v string) *StockSaveRequest {
	s.HotelId = &v
	return s
}

func (s *StockSaveRequest) SetHeader(v map[string]*string) *StockSaveRequest {
	s.Header = v
	return s
}

func (s *StockSaveRequest) SetAccessToken(v string) *StockSaveRequest {
	s.AccessToken = &v
	return s
}

type StockSaveRequestArisItem struct {
	LengthOfStay *int64                             `json:"length_of_stay,omitempty" xml:"length_of_stay,omitempty"`
	Fplos        *string                            `json:"fplos,omitempty" xml:"fplos,omitempty"`
	Priority     *int                               `json:"priority,omitempty" xml:"priority,omitempty"`
	RoomId       *string                            `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Timerange    *StockSaveRequestArisItemTimerange `json:"timerange,omitempty" xml:"timerange,omitempty" require:"true"`
	Inventory    *int64                             `json:"inventory,omitempty" xml:"inventory,omitempty"`
	Available    *bool                              `json:"available,omitempty" xml:"available,omitempty"`
	DaysOfWeek   []*int                             `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	RatePlanId   *string                            `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
}

func (s StockSaveRequestArisItem) String() string {
	return tea.Prettify(s)
}

func (s StockSaveRequestArisItem) GoString() string {
	return s.String()
}

func (s *StockSaveRequestArisItem) SetLengthOfStay(v int64) *StockSaveRequestArisItem {
	s.LengthOfStay = &v
	return s
}

func (s *StockSaveRequestArisItem) SetFplos(v string) *StockSaveRequestArisItem {
	s.Fplos = &v
	return s
}

func (s *StockSaveRequestArisItem) SetPriority(v int) *StockSaveRequestArisItem {
	s.Priority = &v
	return s
}

func (s *StockSaveRequestArisItem) SetRoomId(v string) *StockSaveRequestArisItem {
	s.RoomId = &v
	return s
}

func (s *StockSaveRequestArisItem) SetTimerange(v *StockSaveRequestArisItemTimerange) *StockSaveRequestArisItem {
	s.Timerange = v
	return s
}

func (s *StockSaveRequestArisItem) SetInventory(v int64) *StockSaveRequestArisItem {
	s.Inventory = &v
	return s
}

func (s *StockSaveRequestArisItem) SetAvailable(v bool) *StockSaveRequestArisItem {
	s.Available = &v
	return s
}

func (s *StockSaveRequestArisItem) SetDaysOfWeek(v []*int) *StockSaveRequestArisItem {
	s.DaysOfWeek = v
	return s
}

func (s *StockSaveRequestArisItem) SetRatePlanId(v string) *StockSaveRequestArisItem {
	s.RatePlanId = &v
	return s
}

type StockSaveRequestArisItemTimerange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s StockSaveRequestArisItemTimerange) String() string {
	return tea.Prettify(s)
}

func (s StockSaveRequestArisItemTimerange) GoString() string {
	return s.String()
}

func (s *StockSaveRequestArisItemTimerange) SetEnd(v string) *StockSaveRequestArisItemTimerange {
	s.End = &v
	return s
}

func (s *StockSaveRequestArisItemTimerange) SetStart(v string) *StockSaveRequestArisItemTimerange {
	s.Start = &v
	return s
}

type StockSaveResponse struct {
	Data  *StockSaveResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *StockSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s StockSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s StockSaveResponse) GoString() string {
	return s.String()
}

func (s *StockSaveResponse) SetData(v *StockSaveResponseData) *StockSaveResponse {
	s.Data = v
	return s
}

func (s *StockSaveResponse) SetExtra(v *StockSaveResponseExtra) *StockSaveResponse {
	s.Extra = v
	return s
}

type StockSaveResponseData struct {
	SaveResult    []*StockSaveResponseDataSaveResultItem `json:"save_result,omitempty" xml:"save_result,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s StockSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s StockSaveResponseData) GoString() string {
	return s.String()
}

func (s *StockSaveResponseData) SetSaveResult(v []*StockSaveResponseDataSaveResultItem) *StockSaveResponseData {
	s.SaveResult = v
	return s
}

func (s *StockSaveResponseData) SetGwErrorCode(v int32) *StockSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *StockSaveResponseData) SetGwDescription(v string) *StockSaveResponseData {
	s.GwDescription = &v
	return s
}

type StockSaveResponseDataSaveResultItem struct {
	Message    *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	RatePlanId *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	Code       *int64  `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s StockSaveResponseDataSaveResultItem) String() string {
	return tea.Prettify(s)
}

func (s StockSaveResponseDataSaveResultItem) GoString() string {
	return s.String()
}

func (s *StockSaveResponseDataSaveResultItem) SetMessage(v string) *StockSaveResponseDataSaveResultItem {
	s.Message = &v
	return s
}

func (s *StockSaveResponseDataSaveResultItem) SetRatePlanId(v string) *StockSaveResponseDataSaveResultItem {
	s.RatePlanId = &v
	return s
}

func (s *StockSaveResponseDataSaveResultItem) SetCode(v int64) *StockSaveResponseDataSaveResultItem {
	s.Code = &v
	return s
}

type StockSaveResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s StockSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s StockSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *StockSaveResponseExtra) SetLogid(v string) *StockSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *StockSaveResponseExtra) SetNow(v int64) *StockSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *StockSaveResponseExtra) SetSubDescription(v string) *StockSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *StockSaveResponseExtra) SetSubErrorCode(v int32) *StockSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *StockSaveResponseExtra) SetDescription(v string) *StockSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *StockSaveResponseExtra) SetErrorCode(v int32) *StockSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

type StockSyncRequest struct {
	Stock       *StockSyncRequestStock `json:"stock,omitempty" xml:"stock,omitempty" require:"true"`
	Header      map[string]*string     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductId   *string                `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	OutId       *string                `json:"out_id,omitempty" xml:"out_id,omitempty" require:"true"`
}

func (s StockSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s StockSyncRequest) GoString() string {
	return s.String()
}

func (s *StockSyncRequest) SetStock(v *StockSyncRequestStock) *StockSyncRequest {
	s.Stock = v
	return s
}

func (s *StockSyncRequest) SetHeader(v map[string]*string) *StockSyncRequest {
	s.Header = v
	return s
}

func (s *StockSyncRequest) SetAccessToken(v string) *StockSyncRequest {
	s.AccessToken = &v
	return s
}

func (s *StockSyncRequest) SetProductId(v string) *StockSyncRequest {
	s.ProductId = &v
	return s
}

func (s *StockSyncRequest) SetOutId(v string) *StockSyncRequest {
	s.OutId = &v
	return s
}

type StockSyncRequestStock struct {
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty" require:"true"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s StockSyncRequestStock) String() string {
	return tea.Prettify(s)
}

func (s StockSyncRequestStock) GoString() string {
	return s.String()
}

func (s *StockSyncRequestStock) SetStockQty(v int64) *StockSyncRequestStock {
	s.StockQty = &v
	return s
}

func (s *StockSyncRequestStock) SetLimitType(v int) *StockSyncRequestStock {
	s.LimitType = &v
	return s
}

type StockSyncResponse struct {
	BaseResp *StockSyncResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s StockSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s StockSyncResponse) GoString() string {
	return s.String()
}

func (s *StockSyncResponse) SetBaseResp(v *StockSyncResponseBaseResp) *StockSyncResponse {
	s.BaseResp = v
	return s
}

type StockSyncResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty" require:"true"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty" require:"true"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty" require:"true"`
}

func (s StockSyncResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s StockSyncResponseBaseResp) GoString() string {
	return s.String()
}

func (s *StockSyncResponseBaseResp) SetStatusCode(v int32) *StockSyncResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *StockSyncResponseBaseResp) SetExtra(v map[string]*string) *StockSyncResponseBaseResp {
	s.Extra = v
	return s
}

func (s *StockSyncResponseBaseResp) SetStatusMessage(v string) *StockSyncResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type StockUpdateNotifyRequest struct {
	Date        *string            `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	PoiId       *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s StockUpdateNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s StockUpdateNotifyRequest) GoString() string {
	return s.String()
}

func (s *StockUpdateNotifyRequest) SetDate(v string) *StockUpdateNotifyRequest {
	s.Date = &v
	return s
}

func (s *StockUpdateNotifyRequest) SetPoiId(v string) *StockUpdateNotifyRequest {
	s.PoiId = &v
	return s
}

func (s *StockUpdateNotifyRequest) SetHeader(v map[string]*string) *StockUpdateNotifyRequest {
	s.Header = v
	return s
}

func (s *StockUpdateNotifyRequest) SetAccessToken(v string) *StockUpdateNotifyRequest {
	s.AccessToken = &v
	return s
}

func (s *StockUpdateNotifyRequest) SetAccountId(v string) *StockUpdateNotifyRequest {
	s.AccountId = &v
	return s
}

type StockUpdateNotifyResponse struct {
	BaseResp *StockUpdateNotifyResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty"`
	Data     *StockUpdateNotifyResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *StockUpdateNotifyResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s StockUpdateNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s StockUpdateNotifyResponse) GoString() string {
	return s.String()
}

func (s *StockUpdateNotifyResponse) SetBaseResp(v *StockUpdateNotifyResponseBaseResp) *StockUpdateNotifyResponse {
	s.BaseResp = v
	return s
}

func (s *StockUpdateNotifyResponse) SetData(v *StockUpdateNotifyResponseData) *StockUpdateNotifyResponse {
	s.Data = v
	return s
}

func (s *StockUpdateNotifyResponse) SetExtra(v *StockUpdateNotifyResponseExtra) *StockUpdateNotifyResponse {
	s.Extra = v
	return s
}

type StockUpdateNotifyResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s StockUpdateNotifyResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s StockUpdateNotifyResponseBaseResp) GoString() string {
	return s.String()
}

func (s *StockUpdateNotifyResponseBaseResp) SetStatusCode(v int32) *StockUpdateNotifyResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *StockUpdateNotifyResponseBaseResp) SetStatusMessage(v string) *StockUpdateNotifyResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *StockUpdateNotifyResponseBaseResp) SetExtra(v map[string]*string) *StockUpdateNotifyResponseBaseResp {
	s.Extra = v
	return s
}

type StockUpdateNotifyResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s StockUpdateNotifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s StockUpdateNotifyResponseData) GoString() string {
	return s.String()
}

func (s *StockUpdateNotifyResponseData) SetDescription(v string) *StockUpdateNotifyResponseData {
	s.Description = &v
	return s
}

func (s *StockUpdateNotifyResponseData) SetErrorCode(v int32) *StockUpdateNotifyResponseData {
	s.ErrorCode = &v
	return s
}

type StockUpdateNotifyResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s StockUpdateNotifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s StockUpdateNotifyResponseExtra) GoString() string {
	return s.String()
}

func (s *StockUpdateNotifyResponseExtra) SetNow(v int64) *StockUpdateNotifyResponseExtra {
	s.Now = &v
	return s
}

func (s *StockUpdateNotifyResponseExtra) SetSubDescription(v string) *StockUpdateNotifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *StockUpdateNotifyResponseExtra) SetSubErrorCode(v int32) *StockUpdateNotifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *StockUpdateNotifyResponseExtra) SetDescription(v string) *StockUpdateNotifyResponseExtra {
	s.Description = &v
	return s
}

func (s *StockUpdateNotifyResponseExtra) SetErrorCode(v int32) *StockUpdateNotifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *StockUpdateNotifyResponseExtra) SetLogid(v string) *StockUpdateNotifyResponseExtra {
	s.Logid = &v
	return s
}

type SubscriptionAddAppTplRequest struct {
	TemplateId   *int64             `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
	KeywordList  []*string          `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" require:"true" type:"Repeated"`
	ISVClientKey *string            `json:"ISVClientKey,omitempty" xml:"ISVClientKey,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SubscriptionAddAppTplRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionAddAppTplRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionAddAppTplRequest) SetTemplateId(v int64) *SubscriptionAddAppTplRequest {
	s.TemplateId = &v
	return s
}

func (s *SubscriptionAddAppTplRequest) SetKeywordList(v []*string) *SubscriptionAddAppTplRequest {
	s.KeywordList = v
	return s
}

func (s *SubscriptionAddAppTplRequest) SetISVClientKey(v string) *SubscriptionAddAppTplRequest {
	s.ISVClientKey = &v
	return s
}

func (s *SubscriptionAddAppTplRequest) SetHeader(v map[string]*string) *SubscriptionAddAppTplRequest {
	s.Header = v
	return s
}

func (s *SubscriptionAddAppTplRequest) SetAccessToken(v string) *SubscriptionAddAppTplRequest {
	s.AccessToken = &v
	return s
}

type SubscriptionAddAppTplResponse struct {
	Data   *SubscriptionAddAppTplResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SubscriptionAddAppTplResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionAddAppTplResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionAddAppTplResponse) SetData(v *SubscriptionAddAppTplResponseData) *SubscriptionAddAppTplResponse {
	s.Data = v
	return s
}

func (s *SubscriptionAddAppTplResponse) SetErrNo(v int32) *SubscriptionAddAppTplResponse {
	s.ErrNo = &v
	return s
}

func (s *SubscriptionAddAppTplResponse) SetErrMsg(v string) *SubscriptionAddAppTplResponse {
	s.ErrMsg = &v
	return s
}

func (s *SubscriptionAddAppTplResponse) SetLogId(v string) *SubscriptionAddAppTplResponse {
	s.LogId = &v
	return s
}

type SubscriptionAddAppTplResponseData struct {
	MsgId *string `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
}

func (s SubscriptionAddAppTplResponseData) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionAddAppTplResponseData) GoString() string {
	return s.String()
}

func (s *SubscriptionAddAppTplResponseData) SetMsgId(v string) *SubscriptionAddAppTplResponseData {
	s.MsgId = &v
	return s
}

type SubscriptionCreateTplRequest struct {
	Title          *string            `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	KeywordList    []*string          `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" require:"true" type:"Repeated"`
	CategoryIds    *string            `json:"category_ids,omitempty" xml:"category_ids,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Classification *int               `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	HostList       []*string          `json:"host_list,omitempty" xml:"host_list,omitempty" require:"true" type:"Repeated"`
}

func (s SubscriptionCreateTplRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionCreateTplRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionCreateTplRequest) SetTitle(v string) *SubscriptionCreateTplRequest {
	s.Title = &v
	return s
}

func (s *SubscriptionCreateTplRequest) SetKeywordList(v []*string) *SubscriptionCreateTplRequest {
	s.KeywordList = v
	return s
}

func (s *SubscriptionCreateTplRequest) SetCategoryIds(v string) *SubscriptionCreateTplRequest {
	s.CategoryIds = &v
	return s
}

func (s *SubscriptionCreateTplRequest) SetHeader(v map[string]*string) *SubscriptionCreateTplRequest {
	s.Header = v
	return s
}

func (s *SubscriptionCreateTplRequest) SetAccessToken(v string) *SubscriptionCreateTplRequest {
	s.AccessToken = &v
	return s
}

func (s *SubscriptionCreateTplRequest) SetClassification(v int) *SubscriptionCreateTplRequest {
	s.Classification = &v
	return s
}

func (s *SubscriptionCreateTplRequest) SetHostList(v []*string) *SubscriptionCreateTplRequest {
	s.HostList = v
	return s
}

type SubscriptionCreateTplResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s SubscriptionCreateTplResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionCreateTplResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionCreateTplResponse) SetLogId(v string) *SubscriptionCreateTplResponse {
	s.LogId = &v
	return s
}

func (s *SubscriptionCreateTplResponse) SetErrNo(v int32) *SubscriptionCreateTplResponse {
	s.ErrNo = &v
	return s
}

func (s *SubscriptionCreateTplResponse) SetErrMsg(v string) *SubscriptionCreateTplResponse {
	s.ErrMsg = &v
	return s
}

type SubscriptionDeleteAppTplRequest struct {
	MsgId       *string            `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SubscriptionDeleteAppTplRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionDeleteAppTplRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionDeleteAppTplRequest) SetMsgId(v string) *SubscriptionDeleteAppTplRequest {
	s.MsgId = &v
	return s
}

func (s *SubscriptionDeleteAppTplRequest) SetHeader(v map[string]*string) *SubscriptionDeleteAppTplRequest {
	s.Header = v
	return s
}

func (s *SubscriptionDeleteAppTplRequest) SetAccessToken(v string) *SubscriptionDeleteAppTplRequest {
	s.AccessToken = &v
	return s
}

type SubscriptionDeleteAppTplResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SubscriptionDeleteAppTplResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionDeleteAppTplResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionDeleteAppTplResponse) SetErrMsg(v string) *SubscriptionDeleteAppTplResponse {
	s.ErrMsg = &v
	return s
}

func (s *SubscriptionDeleteAppTplResponse) SetLogId(v string) *SubscriptionDeleteAppTplResponse {
	s.LogId = &v
	return s
}

func (s *SubscriptionDeleteAppTplResponse) SetErrNo(v int32) *SubscriptionDeleteAppTplResponse {
	s.ErrNo = &v
	return s
}

type SubscriptionNotifyUserRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Page        *string            `json:"page,omitempty" xml:"page,omitempty"`
	Data        map[string]*string `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	NotifyType  []*int             `json:"notify_type,omitempty" xml:"notify_type,omitempty" type:"Repeated"`
	MsgId       *string            `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
}

func (s SubscriptionNotifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionNotifyUserRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionNotifyUserRequest) SetHeader(v map[string]*string) *SubscriptionNotifyUserRequest {
	s.Header = v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetAccessToken(v string) *SubscriptionNotifyUserRequest {
	s.AccessToken = &v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetOpenId(v string) *SubscriptionNotifyUserRequest {
	s.OpenId = &v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetPage(v string) *SubscriptionNotifyUserRequest {
	s.Page = &v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetData(v map[string]*string) *SubscriptionNotifyUserRequest {
	s.Data = v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetNotifyType(v []*int) *SubscriptionNotifyUserRequest {
	s.NotifyType = v
	return s
}

func (s *SubscriptionNotifyUserRequest) SetMsgId(v string) *SubscriptionNotifyUserRequest {
	s.MsgId = &v
	return s
}

type SubscriptionNotifyUserResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SubscriptionNotifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionNotifyUserResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionNotifyUserResponse) SetErrNo(v int32) *SubscriptionNotifyUserResponse {
	s.ErrNo = &v
	return s
}

func (s *SubscriptionNotifyUserResponse) SetErrMsg(v string) *SubscriptionNotifyUserResponse {
	s.ErrMsg = &v
	return s
}

func (s *SubscriptionNotifyUserResponse) SetLogId(v string) *SubscriptionNotifyUserResponse {
	s.LogId = &v
	return s
}

type SubscriptionQueryAppTplRequest struct {
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageNum        *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize       *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Classification *int               `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	CategoryIds    *string            `json:"category_ids,omitempty" xml:"category_ids,omitempty"`
}

func (s SubscriptionQueryAppTplRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryAppTplRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryAppTplRequest) SetHeader(v map[string]*string) *SubscriptionQueryAppTplRequest {
	s.Header = v
	return s
}

func (s *SubscriptionQueryAppTplRequest) SetAccessToken(v string) *SubscriptionQueryAppTplRequest {
	s.AccessToken = &v
	return s
}

func (s *SubscriptionQueryAppTplRequest) SetPageNum(v int32) *SubscriptionQueryAppTplRequest {
	s.PageNum = &v
	return s
}

func (s *SubscriptionQueryAppTplRequest) SetPageSize(v int32) *SubscriptionQueryAppTplRequest {
	s.PageSize = &v
	return s
}

func (s *SubscriptionQueryAppTplRequest) SetClassification(v int) *SubscriptionQueryAppTplRequest {
	s.Classification = &v
	return s
}

func (s *SubscriptionQueryAppTplRequest) SetCategoryIds(v string) *SubscriptionQueryAppTplRequest {
	s.CategoryIds = &v
	return s
}

type SubscriptionQueryAppTplResponse struct {
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SubscriptionQueryAppTplResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s SubscriptionQueryAppTplResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryAppTplResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryAppTplResponse) SetLogId(v string) *SubscriptionQueryAppTplResponse {
	s.LogId = &v
	return s
}

func (s *SubscriptionQueryAppTplResponse) SetData(v *SubscriptionQueryAppTplResponseData) *SubscriptionQueryAppTplResponse {
	s.Data = v
	return s
}

func (s *SubscriptionQueryAppTplResponse) SetErrNo(v int32) *SubscriptionQueryAppTplResponse {
	s.ErrNo = &v
	return s
}

func (s *SubscriptionQueryAppTplResponse) SetErrMsg(v string) *SubscriptionQueryAppTplResponse {
	s.ErrMsg = &v
	return s
}

type SubscriptionQueryAppTplResponseData struct {
	TemplateList []*SubscriptionQueryAppTplResponseDataTemplateListItem `json:"template_list,omitempty" xml:"template_list,omitempty" require:"true" type:"Repeated"`
	TotalCount   *int64                                                 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
}

func (s SubscriptionQueryAppTplResponseData) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryAppTplResponseData) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryAppTplResponseData) SetTemplateList(v []*SubscriptionQueryAppTplResponseDataTemplateListItem) *SubscriptionQueryAppTplResponseData {
	s.TemplateList = v
	return s
}

func (s *SubscriptionQueryAppTplResponseData) SetTotalCount(v int64) *SubscriptionQueryAppTplResponseData {
	s.TotalCount = &v
	return s
}

type SubscriptionQueryAppTplResponseDataTemplateListItem struct {
	Classification *int      `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	KeywordList    []*string `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" require:"true" type:"Repeated"`
	CategoryName   *string   `json:"category_name,omitempty" xml:"category_name,omitempty" require:"true"`
	HostList       []*string `json:"host_list,omitempty" xml:"host_list,omitempty" require:"true" type:"Repeated"`
	MsgId          *string   `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	TemplateId     *int64    `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
}

func (s SubscriptionQueryAppTplResponseDataTemplateListItem) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryAppTplResponseDataTemplateListItem) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetClassification(v int) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.Classification = &v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetTitle(v string) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.Title = &v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetKeywordList(v []*string) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.KeywordList = v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetCategoryName(v string) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.CategoryName = &v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetHostList(v []*string) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.HostList = v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetMsgId(v string) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.MsgId = &v
	return s
}

func (s *SubscriptionQueryAppTplResponseDataTemplateListItem) SetTemplateId(v int64) *SubscriptionQueryAppTplResponseDataTemplateListItem {
	s.TemplateId = &v
	return s
}

type SubscriptionQueryTplListRequest struct {
	Classification *int               `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TemplateType   *int               `json:"template_type,omitempty" xml:"template_type,omitempty" require:"true"`
	CategoryIds    *string            `json:"category_ids,omitempty" xml:"category_ids,omitempty"`
	Keyword        *string            `json:"keyword,omitempty" xml:"keyword,omitempty"`
	PageNum        *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize       *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s SubscriptionQueryTplListRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryTplListRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryTplListRequest) SetClassification(v int) *SubscriptionQueryTplListRequest {
	s.Classification = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetHeader(v map[string]*string) *SubscriptionQueryTplListRequest {
	s.Header = v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetAccessToken(v string) *SubscriptionQueryTplListRequest {
	s.AccessToken = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetTemplateType(v int) *SubscriptionQueryTplListRequest {
	s.TemplateType = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetCategoryIds(v string) *SubscriptionQueryTplListRequest {
	s.CategoryIds = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetKeyword(v string) *SubscriptionQueryTplListRequest {
	s.Keyword = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetPageNum(v int32) *SubscriptionQueryTplListRequest {
	s.PageNum = &v
	return s
}

func (s *SubscriptionQueryTplListRequest) SetPageSize(v int32) *SubscriptionQueryTplListRequest {
	s.PageSize = &v
	return s
}

type SubscriptionQueryTplListResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SubscriptionQueryTplListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SubscriptionQueryTplListResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryTplListResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryTplListResponse) SetErrMsg(v string) *SubscriptionQueryTplListResponse {
	s.ErrMsg = &v
	return s
}

func (s *SubscriptionQueryTplListResponse) SetLogId(v string) *SubscriptionQueryTplListResponse {
	s.LogId = &v
	return s
}

func (s *SubscriptionQueryTplListResponse) SetData(v *SubscriptionQueryTplListResponseData) *SubscriptionQueryTplListResponse {
	s.Data = v
	return s
}

func (s *SubscriptionQueryTplListResponse) SetErrNo(v int32) *SubscriptionQueryTplListResponse {
	s.ErrNo = &v
	return s
}

type SubscriptionQueryTplListResponseData struct {
	TotalCount   *int64                                                  `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	TemplateList []*SubscriptionQueryTplListResponseDataTemplateListItem `json:"template_list,omitempty" xml:"template_list,omitempty" require:"true" type:"Repeated"`
}

func (s SubscriptionQueryTplListResponseData) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryTplListResponseData) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryTplListResponseData) SetTotalCount(v int64) *SubscriptionQueryTplListResponseData {
	s.TotalCount = &v
	return s
}

func (s *SubscriptionQueryTplListResponseData) SetTemplateList(v []*SubscriptionQueryTplListResponseDataTemplateListItem) *SubscriptionQueryTplListResponseData {
	s.TemplateList = v
	return s
}

type SubscriptionQueryTplListResponseDataTemplateListItem struct {
	KeywordList    []*string `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" require:"true" type:"Repeated"`
	CategoryName   *string   `json:"category_name,omitempty" xml:"category_name,omitempty" require:"true"`
	HostList       []*string `json:"host_list,omitempty" xml:"host_list,omitempty" require:"true" type:"Repeated"`
	TemplateId     *int64    `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
	Classification *int      `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s SubscriptionQueryTplListResponseDataTemplateListItem) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryTplListResponseDataTemplateListItem) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetKeywordList(v []*string) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.KeywordList = v
	return s
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetCategoryName(v string) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.CategoryName = &v
	return s
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetHostList(v []*string) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.HostList = v
	return s
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetTemplateId(v int64) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.TemplateId = &v
	return s
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetClassification(v int) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.Classification = &v
	return s
}

func (s *SubscriptionQueryTplListResponseDataTemplateListItem) SetTitle(v string) *SubscriptionQueryTplListResponseDataTemplateListItem {
	s.Title = &v
	return s
}

type SubscriptionQueryUserSubscribeRequest struct {
	MsgId       *string            `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SubscriptionQueryUserSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryUserSubscribeRequest) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryUserSubscribeRequest) SetMsgId(v string) *SubscriptionQueryUserSubscribeRequest {
	s.MsgId = &v
	return s
}

func (s *SubscriptionQueryUserSubscribeRequest) SetOpenId(v string) *SubscriptionQueryUserSubscribeRequest {
	s.OpenId = &v
	return s
}

func (s *SubscriptionQueryUserSubscribeRequest) SetHeader(v map[string]*string) *SubscriptionQueryUserSubscribeRequest {
	s.Header = v
	return s
}

func (s *SubscriptionQueryUserSubscribeRequest) SetAccessToken(v string) *SubscriptionQueryUserSubscribeRequest {
	s.AccessToken = &v
	return s
}

type SubscriptionQueryUserSubscribeResponse struct {
	Data   *SubscriptionQueryUserSubscribeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SubscriptionQueryUserSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryUserSubscribeResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryUserSubscribeResponse) SetData(v *SubscriptionQueryUserSubscribeResponseData) *SubscriptionQueryUserSubscribeResponse {
	s.Data = v
	return s
}

func (s *SubscriptionQueryUserSubscribeResponse) SetErrNo(v int32) *SubscriptionQueryUserSubscribeResponse {
	s.ErrNo = &v
	return s
}

func (s *SubscriptionQueryUserSubscribeResponse) SetErrMsg(v string) *SubscriptionQueryUserSubscribeResponse {
	s.ErrMsg = &v
	return s
}

func (s *SubscriptionQueryUserSubscribeResponse) SetLogId(v string) *SubscriptionQueryUserSubscribeResponse {
	s.LogId = &v
	return s
}

type SubscriptionQueryUserSubscribeResponseData struct {
	UserAllowNotification *bool  `json:"user_allow_notification,omitempty" xml:"user_allow_notification,omitempty"`
	SubscribeStatus       *int32 `json:"subscribe_status,omitempty" xml:"subscribe_status,omitempty"`
}

func (s SubscriptionQueryUserSubscribeResponseData) String() string {
	return tea.Prettify(s)
}

func (s SubscriptionQueryUserSubscribeResponseData) GoString() string {
	return s.String()
}

func (s *SubscriptionQueryUserSubscribeResponseData) SetUserAllowNotification(v bool) *SubscriptionQueryUserSubscribeResponseData {
	s.UserAllowNotification = &v
	return s
}

func (s *SubscriptionQueryUserSubscribeResponseData) SetSubscribeStatus(v int32) *SubscriptionQueryUserSubscribeResponseData {
	s.SubscribeStatus = &v
	return s
}

type SyncStatusRequest struct {
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId        *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	OrderOutStatus *int               `json:"order_out_status,omitempty" xml:"order_out_status,omitempty" require:"true"`
}

func (s SyncStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncStatusRequest) GoString() string {
	return s.String()
}

func (s *SyncStatusRequest) SetHeader(v map[string]*string) *SyncStatusRequest {
	s.Header = v
	return s
}

func (s *SyncStatusRequest) SetAccessToken(v string) *SyncStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *SyncStatusRequest) SetOrderId(v string) *SyncStatusRequest {
	s.OrderId = &v
	return s
}

func (s *SyncStatusRequest) SetOrderOutStatus(v int) *SyncStatusRequest {
	s.OrderOutStatus = &v
	return s
}

type SyncStatusResponse struct {
	Extra *SyncStatusResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *SyncStatusResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SyncStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncStatusResponse) GoString() string {
	return s.String()
}

func (s *SyncStatusResponse) SetExtra(v *SyncStatusResponseExtra) *SyncStatusResponse {
	s.Extra = v
	return s
}

func (s *SyncStatusResponse) SetData(v *SyncStatusResponseData) *SyncStatusResponse {
	s.Data = v
	return s
}

type SyncStatusResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s SyncStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s SyncStatusResponseData) GoString() string {
	return s.String()
}

func (s *SyncStatusResponseData) SetGwDescription(v string) *SyncStatusResponseData {
	s.GwDescription = &v
	return s
}

func (s *SyncStatusResponseData) SetOrderId(v string) *SyncStatusResponseData {
	s.OrderId = &v
	return s
}

func (s *SyncStatusResponseData) SetSuccess(v bool) *SyncStatusResponseData {
	s.Success = &v
	return s
}

func (s *SyncStatusResponseData) SetGwErrorCode(v int32) *SyncStatusResponseData {
	s.GwErrorCode = &v
	return s
}

type SyncStatusResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s SyncStatusResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s SyncStatusResponseExtra) GoString() string {
	return s.String()
}

func (s *SyncStatusResponseExtra) SetLogid(v string) *SyncStatusResponseExtra {
	s.Logid = &v
	return s
}

func (s *SyncStatusResponseExtra) SetNow(v int64) *SyncStatusResponseExtra {
	s.Now = &v
	return s
}

func (s *SyncStatusResponseExtra) SetSubDescription(v string) *SyncStatusResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *SyncStatusResponseExtra) SetSubErrorCode(v int32) *SyncStatusResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *SyncStatusResponseExtra) SetDescription(v string) *SyncStatusResponseExtra {
	s.Description = &v
	return s
}

func (s *SyncStatusResponseExtra) SetErrorCode(v int32) *SyncStatusResponseExtra {
	s.ErrorCode = &v
	return s
}

type TaskCreateRequest struct {
	Latitude            *string             `json:"latitude,omitempty" xml:"latitude,omitempty" require:"true"`
	AccountId           *string             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AdditionalImageUrls []*string           `json:"additional_image_urls,omitempty" xml:"additional_image_urls,omitempty" type:"Repeated"`
	OpenStatus          *int32              `json:"open_status,omitempty" xml:"open_status,omitempty" require:"true"`
	HeadImageUrls       []*string           `json:"head_image_urls,omitempty" xml:"head_image_urls,omitempty" type:"Repeated"`
	Longitude           *string             `json:"longitude,omitempty" xml:"longitude,omitempty" require:"true"`
	Header              map[string]*string  `json:"header,omitempty" xml:"header,omitempty"`
	TypeCode            *string             `json:"type_code,omitempty" xml:"type_code,omitempty" require:"true"`
	Province            *string             `json:"province,omitempty" xml:"province,omitempty" require:"true"`
	ExtId               *string             `json:"ext_id,omitempty" xml:"ext_id,omitempty" require:"true"`
	PoiName             *string             `json:"poi_name,omitempty" xml:"poi_name,omitempty" require:"true"`
	Address             *string             `json:"address,omitempty" xml:"address,omitempty" require:"true"`
	TelList             []*string           `json:"tel_list,omitempty" xml:"tel_list,omitempty" type:"Repeated"`
	LastTaskId          *string             `json:"last_task_id,omitempty" xml:"last_task_id,omitempty"`
	OpenTimes           map[int32][]*string `json:"open_times,omitempty" xml:"open_times,omitempty"`
	AdditionalInfo      *string             `json:"additional_info,omitempty" xml:"additional_info,omitempty"`
	City                *string             `json:"city,omitempty" xml:"city,omitempty" require:"true"`
	AccessToken         *string             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateRequest) GoString() string {
	return s.String()
}

func (s *TaskCreateRequest) SetLatitude(v string) *TaskCreateRequest {
	s.Latitude = &v
	return s
}

func (s *TaskCreateRequest) SetAccountId(v string) *TaskCreateRequest {
	s.AccountId = &v
	return s
}

func (s *TaskCreateRequest) SetAdditionalImageUrls(v []*string) *TaskCreateRequest {
	s.AdditionalImageUrls = v
	return s
}

func (s *TaskCreateRequest) SetOpenStatus(v int32) *TaskCreateRequest {
	s.OpenStatus = &v
	return s
}

func (s *TaskCreateRequest) SetHeadImageUrls(v []*string) *TaskCreateRequest {
	s.HeadImageUrls = v
	return s
}

func (s *TaskCreateRequest) SetLongitude(v string) *TaskCreateRequest {
	s.Longitude = &v
	return s
}

func (s *TaskCreateRequest) SetHeader(v map[string]*string) *TaskCreateRequest {
	s.Header = v
	return s
}

func (s *TaskCreateRequest) SetTypeCode(v string) *TaskCreateRequest {
	s.TypeCode = &v
	return s
}

func (s *TaskCreateRequest) SetProvince(v string) *TaskCreateRequest {
	s.Province = &v
	return s
}

func (s *TaskCreateRequest) SetExtId(v string) *TaskCreateRequest {
	s.ExtId = &v
	return s
}

func (s *TaskCreateRequest) SetPoiName(v string) *TaskCreateRequest {
	s.PoiName = &v
	return s
}

func (s *TaskCreateRequest) SetAddress(v string) *TaskCreateRequest {
	s.Address = &v
	return s
}

func (s *TaskCreateRequest) SetTelList(v []*string) *TaskCreateRequest {
	s.TelList = v
	return s
}

func (s *TaskCreateRequest) SetLastTaskId(v string) *TaskCreateRequest {
	s.LastTaskId = &v
	return s
}

func (s *TaskCreateRequest) SetOpenTimes(v map[int32][]*string) *TaskCreateRequest {
	s.OpenTimes = v
	return s
}

func (s *TaskCreateRequest) SetAdditionalInfo(v string) *TaskCreateRequest {
	s.AdditionalInfo = &v
	return s
}

func (s *TaskCreateRequest) SetCity(v string) *TaskCreateRequest {
	s.City = &v
	return s
}

func (s *TaskCreateRequest) SetAccessToken(v string) *TaskCreateRequest {
	s.AccessToken = &v
	return s
}

type TaskCreateResponse struct {
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	Extra  *TaskCreateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskCreateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateResponse) GoString() string {
	return s.String()
}

func (s *TaskCreateResponse) SetErrNo(v int32) *TaskCreateResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskCreateResponse) SetExtra(v *TaskCreateResponseExtra) *TaskCreateResponse {
	s.Extra = v
	return s
}

func (s *TaskCreateResponse) SetLogId(v string) *TaskCreateResponse {
	s.LogId = &v
	return s
}

func (s *TaskCreateResponse) SetData(v *TaskCreateResponseData) *TaskCreateResponse {
	s.Data = v
	return s
}

func (s *TaskCreateResponse) SetErrMsg(v string) *TaskCreateResponse {
	s.ErrMsg = &v
	return s
}

type TaskCreateResponseData struct {
	FailMessage   *string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	TaskId        *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TaskCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateResponseData) GoString() string {
	return s.String()
}

func (s *TaskCreateResponseData) SetFailMessage(v string) *TaskCreateResponseData {
	s.FailMessage = &v
	return s
}

func (s *TaskCreateResponseData) SetTaskId(v string) *TaskCreateResponseData {
	s.TaskId = &v
	return s
}

func (s *TaskCreateResponseData) SetGwErrorCode(v int32) *TaskCreateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TaskCreateResponseData) SetGwDescription(v string) *TaskCreateResponseData {
	s.GwDescription = &v
	return s
}

type TaskCreateResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TaskCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *TaskCreateResponseExtra) SetLogid(v string) *TaskCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *TaskCreateResponseExtra) SetNow(v int64) *TaskCreateResponseExtra {
	s.Now = &v
	return s
}

func (s *TaskCreateResponseExtra) SetSubDescription(v string) *TaskCreateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskCreateResponseExtra) SetSubErrorCode(v int32) *TaskCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskCreateResponseExtra) SetDescription(v string) *TaskCreateResponseExtra {
	s.Description = &v
	return s
}

func (s *TaskCreateResponseExtra) SetErrorCode(v int32) *TaskCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

type TaskCreateVideoRequest struct {
	Conditions  []*string          `json:"conditions,omitempty" xml:"conditions,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	ItemId      *string            `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl    *string            `json:"video_url,omitempty" xml:"video_url,omitempty"`
	TaskName    *string            `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
}

func (s TaskCreateVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateVideoRequest) GoString() string {
	return s.String()
}

func (s *TaskCreateVideoRequest) SetConditions(v []*string) *TaskCreateVideoRequest {
	s.Conditions = v
	return s
}

func (s *TaskCreateVideoRequest) SetHeader(v map[string]*string) *TaskCreateVideoRequest {
	s.Header = v
	return s
}

func (s *TaskCreateVideoRequest) SetAccessToken(v string) *TaskCreateVideoRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskCreateVideoRequest) SetStartTime(v int64) *TaskCreateVideoRequest {
	s.StartTime = &v
	return s
}

func (s *TaskCreateVideoRequest) SetEndTime(v int64) *TaskCreateVideoRequest {
	s.EndTime = &v
	return s
}

func (s *TaskCreateVideoRequest) SetItemId(v string) *TaskCreateVideoRequest {
	s.ItemId = &v
	return s
}

func (s *TaskCreateVideoRequest) SetVideoUrl(v string) *TaskCreateVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *TaskCreateVideoRequest) SetTaskName(v string) *TaskCreateVideoRequest {
	s.TaskName = &v
	return s
}

type TaskCreateVideoResponse struct {
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskCreateVideoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s TaskCreateVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateVideoResponse) GoString() string {
	return s.String()
}

func (s *TaskCreateVideoResponse) SetErrMsg(v string) *TaskCreateVideoResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskCreateVideoResponse) SetLogId(v string) *TaskCreateVideoResponse {
	s.LogId = &v
	return s
}

func (s *TaskCreateVideoResponse) SetData(v *TaskCreateVideoResponseData) *TaskCreateVideoResponse {
	s.Data = v
	return s
}

func (s *TaskCreateVideoResponse) SetErrNo(v int32) *TaskCreateVideoResponse {
	s.ErrNo = &v
	return s
}

type TaskCreateVideoResponseData struct {
	ItemId     *string                           `json:"item_id,omitempty" xml:"item_id,omitempty"`
	Extra      *TaskCreateVideoResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	TaskId     *string                           `json:"task_id,omitempty" xml:"task_id,omitempty"`
	TaskStatus *int32                            `json:"task_status,omitempty" xml:"task_status,omitempty"`
}

func (s TaskCreateVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateVideoResponseData) GoString() string {
	return s.String()
}

func (s *TaskCreateVideoResponseData) SetItemId(v string) *TaskCreateVideoResponseData {
	s.ItemId = &v
	return s
}

func (s *TaskCreateVideoResponseData) SetExtra(v *TaskCreateVideoResponseDataExtra) *TaskCreateVideoResponseData {
	s.Extra = v
	return s
}

func (s *TaskCreateVideoResponseData) SetTaskId(v string) *TaskCreateVideoResponseData {
	s.TaskId = &v
	return s
}

func (s *TaskCreateVideoResponseData) SetTaskStatus(v int32) *TaskCreateVideoResponseData {
	s.TaskStatus = &v
	return s
}

type TaskCreateVideoResponseDataExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s TaskCreateVideoResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskCreateVideoResponseDataExtra) GoString() string {
	return s.String()
}

func (s *TaskCreateVideoResponseDataExtra) SetErrorCode(v int32) *TaskCreateVideoResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *TaskCreateVideoResponseDataExtra) SetDescription(v string) *TaskCreateVideoResponseDataExtra {
	s.Description = &v
	return s
}

func (s *TaskCreateVideoResponseDataExtra) SetSubErrorCode(v int32) *TaskCreateVideoResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskCreateVideoResponseDataExtra) SetSubDescription(v string) *TaskCreateVideoResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskCreateVideoResponseDataExtra) SetLogid(v string) *TaskCreateVideoResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *TaskCreateVideoResponseDataExtra) SetNow(v int64) *TaskCreateVideoResponseDataExtra {
	s.Now = &v
	return s
}

type TaskGetRequest struct {
	Roomid      *string            `json:"roomid,omitempty" xml:"roomid,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Appid       *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	MsgType     *string            `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
}

func (s TaskGetRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskGetRequest) GoString() string {
	return s.String()
}

func (s *TaskGetRequest) SetRoomid(v string) *TaskGetRequest {
	s.Roomid = &v
	return s
}

func (s *TaskGetRequest) SetHeader(v map[string]*string) *TaskGetRequest {
	s.Header = v
	return s
}

func (s *TaskGetRequest) SetAccessToken(v string) *TaskGetRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskGetRequest) SetAppid(v string) *TaskGetRequest {
	s.Appid = &v
	return s
}

func (s *TaskGetRequest) SetMsgType(v string) *TaskGetRequest {
	s.MsgType = &v
	return s
}

type TaskGetResponse struct {
	Logid  *string              `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	ErrNo  *int32               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	Data   *TaskGetResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TaskGetResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskGetResponse) GoString() string {
	return s.String()
}

func (s *TaskGetResponse) SetLogid(v string) *TaskGetResponse {
	s.Logid = &v
	return s
}

func (s *TaskGetResponse) SetErrNo(v int32) *TaskGetResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskGetResponse) SetErrMsg(v string) *TaskGetResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskGetResponse) SetData(v *TaskGetResponseData) *TaskGetResponse {
	s.Data = v
	return s
}

type TaskGetResponseData struct {
	Status *int `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TaskGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskGetResponseData) GoString() string {
	return s.String()
}

func (s *TaskGetResponseData) SetStatus(v int) *TaskGetResponseData {
	s.Status = &v
	return s
}

type TaskQueryRequest struct {
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskQueryRequest) GoString() string {
	return s.String()
}

func (s *TaskQueryRequest) SetTaskId(v string) *TaskQueryRequest {
	s.TaskId = &v
	return s
}

func (s *TaskQueryRequest) SetAccountId(v string) *TaskQueryRequest {
	s.AccountId = &v
	return s
}

func (s *TaskQueryRequest) SetHeader(v map[string]*string) *TaskQueryRequest {
	s.Header = v
	return s
}

func (s *TaskQueryRequest) SetAccessToken(v string) *TaskQueryRequest {
	s.AccessToken = &v
	return s
}

type TaskQueryResponse struct {
	Data  *TaskQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TaskQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TaskQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskQueryResponse) GoString() string {
	return s.String()
}

func (s *TaskQueryResponse) SetData(v *TaskQueryResponseData) *TaskQueryResponse {
	s.Data = v
	return s
}

func (s *TaskQueryResponse) SetExtra(v *TaskQueryResponseExtra) *TaskQueryResponse {
	s.Extra = v
	return s
}

type TaskQueryResponseData struct {
	TaskId        *string                                 `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	ErrorCode     *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SuccessPoiCnt *int32                                  `json:"success_poi_cnt,omitempty" xml:"success_poi_cnt,omitempty"`
	AccountId     *string                                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	TotalPoiCnt   *int32                                  `json:"total_poi_cnt,omitempty" xml:"total_poi_cnt,omitempty"`
	PoiIds        []*string                               `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
	TaskContent   *string                                 `json:"task_content,omitempty" xml:"task_content,omitempty"`
	TaskResults   []*TaskQueryResponseDataTaskResultsItem `json:"task_results,omitempty" xml:"task_results,omitempty" type:"Repeated"`
	Description   *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	TaskStatus    *int32                                  `json:"task_status,omitempty" xml:"task_status,omitempty" require:"true"`
}

func (s TaskQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskQueryResponseData) GoString() string {
	return s.String()
}

func (s *TaskQueryResponseData) SetTaskId(v string) *TaskQueryResponseData {
	s.TaskId = &v
	return s
}

func (s *TaskQueryResponseData) SetErrorCode(v int32) *TaskQueryResponseData {
	s.ErrorCode = &v
	return s
}

func (s *TaskQueryResponseData) SetSuccessPoiCnt(v int32) *TaskQueryResponseData {
	s.SuccessPoiCnt = &v
	return s
}

func (s *TaskQueryResponseData) SetAccountId(v string) *TaskQueryResponseData {
	s.AccountId = &v
	return s
}

func (s *TaskQueryResponseData) SetTotalPoiCnt(v int32) *TaskQueryResponseData {
	s.TotalPoiCnt = &v
	return s
}

func (s *TaskQueryResponseData) SetPoiIds(v []*string) *TaskQueryResponseData {
	s.PoiIds = v
	return s
}

func (s *TaskQueryResponseData) SetTaskContent(v string) *TaskQueryResponseData {
	s.TaskContent = &v
	return s
}

func (s *TaskQueryResponseData) SetTaskResults(v []*TaskQueryResponseDataTaskResultsItem) *TaskQueryResponseData {
	s.TaskResults = v
	return s
}

func (s *TaskQueryResponseData) SetDescription(v string) *TaskQueryResponseData {
	s.Description = &v
	return s
}

func (s *TaskQueryResponseData) SetTaskStatus(v int32) *TaskQueryResponseData {
	s.TaskStatus = &v
	return s
}

type TaskQueryResponseDataTaskResultsItem struct {
	Status   *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	PoiId    *string `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
}

func (s TaskQueryResponseDataTaskResultsItem) String() string {
	return tea.Prettify(s)
}

func (s TaskQueryResponseDataTaskResultsItem) GoString() string {
	return s.String()
}

func (s *TaskQueryResponseDataTaskResultsItem) SetStatus(v int32) *TaskQueryResponseDataTaskResultsItem {
	s.Status = &v
	return s
}

func (s *TaskQueryResponseDataTaskResultsItem) SetErrorMsg(v string) *TaskQueryResponseDataTaskResultsItem {
	s.ErrorMsg = &v
	return s
}

func (s *TaskQueryResponseDataTaskResultsItem) SetPoiId(v string) *TaskQueryResponseDataTaskResultsItem {
	s.PoiId = &v
	return s
}

type TaskQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TaskQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *TaskQueryResponseExtra) SetNow(v int64) *TaskQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *TaskQueryResponseExtra) SetSubDescription(v string) *TaskQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskQueryResponseExtra) SetSubErrorCode(v int32) *TaskQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskQueryResponseExtra) SetDescription(v string) *TaskQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *TaskQueryResponseExtra) SetErrorCode(v int32) *TaskQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TaskQueryResponseExtra) SetLogid(v string) *TaskQueryResponseExtra {
	s.Logid = &v
	return s
}

type TaskStartRequest struct {
	Appid       *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	MsgType     *string            `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	Roomid      *string            `json:"roomid,omitempty" xml:"roomid,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskStartRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskStartRequest) GoString() string {
	return s.String()
}

func (s *TaskStartRequest) SetAppid(v string) *TaskStartRequest {
	s.Appid = &v
	return s
}

func (s *TaskStartRequest) SetMsgType(v string) *TaskStartRequest {
	s.MsgType = &v
	return s
}

func (s *TaskStartRequest) SetRoomid(v string) *TaskStartRequest {
	s.Roomid = &v
	return s
}

func (s *TaskStartRequest) SetHeader(v map[string]*string) *TaskStartRequest {
	s.Header = v
	return s
}

func (s *TaskStartRequest) SetAccessToken(v string) *TaskStartRequest {
	s.AccessToken = &v
	return s
}

type TaskStartResponse struct {
	Data   *TaskStartResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Logid  *string                `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	ErrNo  *int32                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskStartResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskStartResponse) GoString() string {
	return s.String()
}

func (s *TaskStartResponse) SetData(v *TaskStartResponseData) *TaskStartResponse {
	s.Data = v
	return s
}

func (s *TaskStartResponse) SetLogid(v string) *TaskStartResponse {
	s.Logid = &v
	return s
}

func (s *TaskStartResponse) SetErrNo(v int32) *TaskStartResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskStartResponse) SetErrMsg(v string) *TaskStartResponse {
	s.ErrMsg = &v
	return s
}

type TaskStartResponseData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s TaskStartResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskStartResponseData) GoString() string {
	return s.String()
}

func (s *TaskStartResponseData) SetTaskId(v string) *TaskStartResponseData {
	s.TaskId = &v
	return s
}

type TaskSubmitRequest struct {
	Datas       []*TaskSubmitRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" type:"Repeated"`
	Header      map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskSubmitRequest) GoString() string {
	return s.String()
}

func (s *TaskSubmitRequest) SetDatas(v []*TaskSubmitRequestDatasItem) *TaskSubmitRequest {
	s.Datas = v
	return s
}

func (s *TaskSubmitRequest) SetHeader(v map[string]*string) *TaskSubmitRequest {
	s.Header = v
	return s
}

func (s *TaskSubmitRequest) SetAccessToken(v string) *TaskSubmitRequest {
	s.AccessToken = &v
	return s
}

type TaskSubmitRequestDatasItem struct {
	ContactTel    *string           `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
	OpenTimes     map[int][]*string `json:"open_times,omitempty" xml:"open_times,omitempty"`
	ContactPhone  *string           `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	PoiId         *string           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	City          *string           `json:"city,omitempty" xml:"city,omitempty"`
	Longitude     *string           `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Province      *string           `json:"province,omitempty" xml:"province,omitempty"`
	PoiName       *string           `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	Address       *string           `json:"address,omitempty" xml:"address,omitempty"`
	Latitude      *string           `json:"latitude,omitempty" xml:"latitude,omitempty"`
	AMapId        *string           `json:"a_map_id,omitempty" xml:"a_map_id,omitempty"`
	HeadImageUrls []*string         `json:"head_image_urls,omitempty" xml:"head_image_urls,omitempty" type:"Repeated"`
	IndustryCode  *string           `json:"industry_code,omitempty" xml:"industry_code,omitempty"`
	ExtId         *string           `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s TaskSubmitRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s TaskSubmitRequestDatasItem) GoString() string {
	return s.String()
}

func (s *TaskSubmitRequestDatasItem) SetContactTel(v string) *TaskSubmitRequestDatasItem {
	s.ContactTel = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetOpenTimes(v map[int][]*string) *TaskSubmitRequestDatasItem {
	s.OpenTimes = v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetContactPhone(v string) *TaskSubmitRequestDatasItem {
	s.ContactPhone = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetPoiId(v string) *TaskSubmitRequestDatasItem {
	s.PoiId = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetCity(v string) *TaskSubmitRequestDatasItem {
	s.City = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetLongitude(v string) *TaskSubmitRequestDatasItem {
	s.Longitude = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetProvince(v string) *TaskSubmitRequestDatasItem {
	s.Province = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetPoiName(v string) *TaskSubmitRequestDatasItem {
	s.PoiName = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetAddress(v string) *TaskSubmitRequestDatasItem {
	s.Address = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetLatitude(v string) *TaskSubmitRequestDatasItem {
	s.Latitude = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetAMapId(v string) *TaskSubmitRequestDatasItem {
	s.AMapId = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetHeadImageUrls(v []*string) *TaskSubmitRequestDatasItem {
	s.HeadImageUrls = v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetIndustryCode(v string) *TaskSubmitRequestDatasItem {
	s.IndustryCode = &v
	return s
}

func (s *TaskSubmitRequestDatasItem) SetExtId(v string) *TaskSubmitRequestDatasItem {
	s.ExtId = &v
	return s
}

type TaskSubmitResponse struct {
	Data  *TaskSubmitResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TaskSubmitResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TaskSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskSubmitResponse) GoString() string {
	return s.String()
}

func (s *TaskSubmitResponse) SetData(v *TaskSubmitResponseData) *TaskSubmitResponse {
	s.Data = v
	return s
}

func (s *TaskSubmitResponse) SetExtra(v *TaskSubmitResponseExtra) *TaskSubmitResponse {
	s.Extra = v
	return s
}

type TaskSubmitResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	TaskId        *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TaskSubmitResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskSubmitResponseData) GoString() string {
	return s.String()
}

func (s *TaskSubmitResponseData) SetGwDescription(v string) *TaskSubmitResponseData {
	s.GwDescription = &v
	return s
}

func (s *TaskSubmitResponseData) SetTaskId(v int64) *TaskSubmitResponseData {
	s.TaskId = &v
	return s
}

func (s *TaskSubmitResponseData) SetGwErrorCode(v int32) *TaskSubmitResponseData {
	s.GwErrorCode = &v
	return s
}
