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

type QueryCreatedTplListResponseDataTemplateListItem struct {
	Classification *int      `json:"classification,omitempty" xml:"classification,omitempty" require:"true"`
	HostList       []*string `json:"host_list,omitempty" xml:"host_list,omitempty" require:"true" type:"Repeated"`
	OperatingTime  *int64    `json:"operating_time,omitempty" xml:"operating_time,omitempty" require:"true"`
	Status         *int      `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	FailReason     *string   `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	Title          *string   `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	KeywordList    []*string `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCreatedTplListResponseDataTemplateListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryCreatedTplListResponseDataTemplateListItem) GoString() string {
	return s.String()
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetClassification(v int) *QueryCreatedTplListResponseDataTemplateListItem {
	s.Classification = &v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetHostList(v []*string) *QueryCreatedTplListResponseDataTemplateListItem {
	s.HostList = v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetOperatingTime(v int64) *QueryCreatedTplListResponseDataTemplateListItem {
	s.OperatingTime = &v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetStatus(v int) *QueryCreatedTplListResponseDataTemplateListItem {
	s.Status = &v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetFailReason(v string) *QueryCreatedTplListResponseDataTemplateListItem {
	s.FailReason = &v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetTitle(v string) *QueryCreatedTplListResponseDataTemplateListItem {
	s.Title = &v
	return s
}

func (s *QueryCreatedTplListResponseDataTemplateListItem) SetKeywordList(v []*string) *QueryCreatedTplListResponseDataTemplateListItem {
	s.KeywordList = v
	return s
}

type QueryDealDataWithConversionRequest struct {
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ScenesList  []*string          `json:"scenes_list,omitempty" xml:"scenes_list,omitempty" type:"Repeated"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s QueryDealDataWithConversionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDealDataWithConversionRequest) GoString() string {
	return s.String()
}

func (s *QueryDealDataWithConversionRequest) SetEndTime(v int64) *QueryDealDataWithConversionRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDealDataWithConversionRequest) SetHostName(v string) *QueryDealDataWithConversionRequest {
	s.HostName = &v
	return s
}

func (s *QueryDealDataWithConversionRequest) SetHeader(v map[string]*string) *QueryDealDataWithConversionRequest {
	s.Header = v
	return s
}

func (s *QueryDealDataWithConversionRequest) SetAccessToken(v string) *QueryDealDataWithConversionRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryDealDataWithConversionRequest) SetScenesList(v []*string) *QueryDealDataWithConversionRequest {
	s.ScenesList = v
	return s
}

func (s *QueryDealDataWithConversionRequest) SetStartTime(v int64) *QueryDealDataWithConversionRequest {
	s.StartTime = &v
	return s
}

type QueryDealDataWithConversionResponse struct {
	Data   *QueryDealDataWithConversionResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryDealDataWithConversionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDealDataWithConversionResponse) GoString() string {
	return s.String()
}

func (s *QueryDealDataWithConversionResponse) SetData(v *QueryDealDataWithConversionResponseData) *QueryDealDataWithConversionResponse {
	s.Data = v
	return s
}

func (s *QueryDealDataWithConversionResponse) SetErrNo(v int32) *QueryDealDataWithConversionResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryDealDataWithConversionResponse) SetErrMsg(v string) *QueryDealDataWithConversionResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryDealDataWithConversionResponse) SetLogId(v string) *QueryDealDataWithConversionResponse {
	s.LogId = &v
	return s
}

type QueryDealDataWithConversionResponseData struct {
	Data *QueryDealDataWithConversionResponseDataData `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryDealDataWithConversionResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDealDataWithConversionResponseData) GoString() string {
	return s.String()
}

func (s *QueryDealDataWithConversionResponseData) SetData(v *QueryDealDataWithConversionResponseDataData) *QueryDealDataWithConversionResponseData {
	s.Data = v
	return s
}

type QueryDealDataWithConversionResponseDataData struct {
	MpShowPv          *int64 `json:"MpShowPv,omitempty" xml:"MpShowPv,omitempty"`
	MpDrainagePv      *int64 `json:"MpDrainagePv,omitempty" xml:"MpDrainagePv,omitempty"`
	RefundPeopleCount *int64 `json:"RefundPeopleCount,omitempty" xml:"RefundPeopleCount,omitempty"`
	RefundOrderCount  *int64 `json:"RefundOrderCount,omitempty" xml:"RefundOrderCount,omitempty"`
	PayOrderCount     *int64 `json:"PayOrderCount,omitempty" xml:"PayOrderCount,omitempty"`
	MpDrainageUv      *int64 `json:"MpDrainageUv,omitempty" xml:"MpDrainageUv,omitempty"`
	MpShowUv          *int64 `json:"MpShowUv,omitempty" xml:"MpShowUv,omitempty"`
	CreateOrderCount  *int64 `json:"CreateOrderCount,omitempty" xml:"CreateOrderCount,omitempty"`
	CreateUserCount   *int64 `json:"CreateUserCount,omitempty" xml:"CreateUserCount,omitempty"`
	PayPeopleCount    *int64 `json:"PayPeopleCount,omitempty" xml:"PayPeopleCount,omitempty"`
}

func (s QueryDealDataWithConversionResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryDealDataWithConversionResponseDataData) GoString() string {
	return s.String()
}

func (s *QueryDealDataWithConversionResponseDataData) SetMpShowPv(v int64) *QueryDealDataWithConversionResponseDataData {
	s.MpShowPv = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetMpDrainagePv(v int64) *QueryDealDataWithConversionResponseDataData {
	s.MpDrainagePv = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetRefundPeopleCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.RefundPeopleCount = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetRefundOrderCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.RefundOrderCount = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetPayOrderCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.PayOrderCount = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetMpDrainageUv(v int64) *QueryDealDataWithConversionResponseDataData {
	s.MpDrainageUv = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetMpShowUv(v int64) *QueryDealDataWithConversionResponseDataData {
	s.MpShowUv = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetCreateOrderCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.CreateOrderCount = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetCreateUserCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.CreateUserCount = &v
	return s
}

func (s *QueryDealDataWithConversionResponseDataData) SetPayPeopleCount(v int64) *QueryDealDataWithConversionResponseDataData {
	s.PayPeopleCount = &v
	return s
}

type QueryDealOverviewDataRequest struct {
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryDealOverviewDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDealOverviewDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDealOverviewDataRequest) SetVersionType(v string) *QueryDealOverviewDataRequest {
	s.VersionType = &v
	return s
}

func (s *QueryDealOverviewDataRequest) SetStartTime(v int64) *QueryDealOverviewDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryDealOverviewDataRequest) SetEndTime(v int64) *QueryDealOverviewDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryDealOverviewDataRequest) SetHostName(v string) *QueryDealOverviewDataRequest {
	s.HostName = &v
	return s
}

func (s *QueryDealOverviewDataRequest) SetHeader(v map[string]*string) *QueryDealOverviewDataRequest {
	s.Header = v
	return s
}

func (s *QueryDealOverviewDataRequest) SetAccessToken(v string) *QueryDealOverviewDataRequest {
	s.AccessToken = &v
	return s
}

type QueryDealOverviewDataResponse struct {
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryDealOverviewDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryDealOverviewDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDealOverviewDataResponse) GoString() string {
	return s.String()
}

func (s *QueryDealOverviewDataResponse) SetErrNo(v int32) *QueryDealOverviewDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryDealOverviewDataResponse) SetErrMsg(v string) *QueryDealOverviewDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryDealOverviewDataResponse) SetLogId(v string) *QueryDealOverviewDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryDealOverviewDataResponse) SetData(v *QueryDealOverviewDataResponseData) *QueryDealOverviewDataResponse {
	s.Data = v
	return s
}

type QueryDealOverviewDataResponseData struct {
	DealOverviewData *QueryDealOverviewDataResponseDataDealOverviewData   `json:"deal_overview_data,omitempty" xml:"deal_overview_data,omitempty" require:"true"`
	DealDataList     []*QueryDealOverviewDataResponseDataDealDataListItem `json:"deal_data_list,omitempty" xml:"deal_data_list,omitempty" type:"Repeated"`
}

func (s QueryDealOverviewDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryDealOverviewDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryDealOverviewDataResponseData) SetDealOverviewData(v *QueryDealOverviewDataResponseDataDealOverviewData) *QueryDealOverviewDataResponseData {
	s.DealOverviewData = v
	return s
}

func (s *QueryDealOverviewDataResponseData) SetDealDataList(v []*QueryDealOverviewDataResponseDataDealDataListItem) *QueryDealOverviewDataResponseData {
	s.DealDataList = v
	return s
}

type QueryDealOverviewDataResponseDataDealDataListItem struct {
	SuccessUserCnt           *int64  `json:"success_user_cnt,omitempty" xml:"success_user_cnt,omitempty"`
	SaleProductCnt           *int64  `json:"sale_product_cnt,omitempty" xml:"sale_product_cnt,omitempty"`
	Time                     *string `json:"time,omitempty" xml:"time,omitempty"`
	PayOrderAmount           *int64  `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty"`
	PayOrderCount            *int64  `json:"pay_order_count,omitempty" xml:"pay_order_count,omitempty"`
	RefundSuccessOrderAmount *int64  `json:"refund_success_order_amount,omitempty" xml:"refund_success_order_amount,omitempty"`
	CreateOrderCount         *int64  `json:"create_order_count,omitempty" xml:"create_order_count,omitempty"`
	RefundSuccessUserCnt     *int64  `json:"refund_success_user_cnt,omitempty" xml:"refund_success_user_cnt,omitempty"`
	RefundSuccessOrderCnt    *int64  `json:"refund_success_order_cnt,omitempty" xml:"refund_success_order_cnt,omitempty"`
	CreateOrderAmount        *int64  `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty"`
}

func (s QueryDealOverviewDataResponseDataDealDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryDealOverviewDataResponseDataDealDataListItem) GoString() string {
	return s.String()
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetSuccessUserCnt(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.SuccessUserCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetSaleProductCnt(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.SaleProductCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetTime(v string) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.Time = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetPayOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetPayOrderCount(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.PayOrderCount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetRefundSuccessOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.RefundSuccessOrderAmount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetCreateOrderCount(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.CreateOrderCount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetRefundSuccessUserCnt(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.RefundSuccessUserCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetRefundSuccessOrderCnt(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.RefundSuccessOrderCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealDataListItem) SetCreateOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealDataListItem {
	s.CreateOrderAmount = &v
	return s
}

type QueryDealOverviewDataResponseDataDealOverviewData struct {
	RefundSuccessUserCnt     *int64  `json:"refund_success_user_cnt,omitempty" xml:"refund_success_user_cnt,omitempty"`
	CreateOrderAmount        *int64  `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty"`
	SaleProductCnt           *int64  `json:"sale_product_cnt,omitempty" xml:"sale_product_cnt,omitempty"`
	PayOrderCount            *int64  `json:"pay_order_count,omitempty" xml:"pay_order_count,omitempty"`
	RefundSuccessOrderAmount *int64  `json:"refund_success_order_amount,omitempty" xml:"refund_success_order_amount,omitempty"`
	Time                     *string `json:"time,omitempty" xml:"time,omitempty"`
	PayOrderAmount           *int64  `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty"`
	SuccessUserCnt           *int64  `json:"success_user_cnt,omitempty" xml:"success_user_cnt,omitempty"`
	RefundSuccessOrderCnt    *int64  `json:"refund_success_order_cnt,omitempty" xml:"refund_success_order_cnt,omitempty"`
	CreateOrderCount         *int64  `json:"create_order_count,omitempty" xml:"create_order_count,omitempty"`
}

func (s QueryDealOverviewDataResponseDataDealOverviewData) String() string {
	return tea.Prettify(s)
}

func (s QueryDealOverviewDataResponseDataDealOverviewData) GoString() string {
	return s.String()
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetRefundSuccessUserCnt(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.RefundSuccessUserCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetCreateOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.CreateOrderAmount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetSaleProductCnt(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.SaleProductCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetPayOrderCount(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.PayOrderCount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetRefundSuccessOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.RefundSuccessOrderAmount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetTime(v string) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.Time = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetPayOrderAmount(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetSuccessUserCnt(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.SuccessUserCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetRefundSuccessOrderCnt(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.RefundSuccessOrderCnt = &v
	return s
}

func (s *QueryDealOverviewDataResponseDataDealOverviewData) SetCreateOrderCount(v int64) *QueryDealOverviewDataResponseDataDealOverviewData {
	s.CreateOrderCount = &v
	return s
}

type QueryItemOrderInfoRequest struct {
	OrderId         *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ItemOrderIdList []*string          `json:"item_order_id_list,omitempty" xml:"item_order_id_list,omitempty" type:"Repeated"`
}

func (s QueryItemOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoRequest) SetOrderId(v string) *QueryItemOrderInfoRequest {
	s.OrderId = &v
	return s
}

func (s *QueryItemOrderInfoRequest) SetHeader(v map[string]*string) *QueryItemOrderInfoRequest {
	s.Header = v
	return s
}

func (s *QueryItemOrderInfoRequest) SetAccessToken(v string) *QueryItemOrderInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryItemOrderInfoRequest) SetItemOrderIdList(v []*string) *QueryItemOrderInfoRequest {
	s.ItemOrderIdList = v
	return s
}

type QueryItemOrderInfoResponse struct {
	Extra *QueryItemOrderInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *QueryItemOrderInfoResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryItemOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponse) SetExtra(v *QueryItemOrderInfoResponseExtra) *QueryItemOrderInfoResponse {
	s.Extra = v
	return s
}

func (s *QueryItemOrderInfoResponse) SetData(v *QueryItemOrderInfoResponseData) *QueryItemOrderInfoResponse {
	s.Data = v
	return s
}

type QueryItemOrderInfoResponseData struct {
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ItemList      []*QueryItemOrderInfoResponseDataItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	OrderSource   *string                                       `json:"order_source,omitempty" xml:"order_source,omitempty"`
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s QueryItemOrderInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseData) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseData) SetGwDescription(v string) *QueryItemOrderInfoResponseData {
	s.GwDescription = &v
	return s
}

func (s *QueryItemOrderInfoResponseData) SetItemList(v []*QueryItemOrderInfoResponseDataItemListItem) *QueryItemOrderInfoResponseData {
	s.ItemList = v
	return s
}

func (s *QueryItemOrderInfoResponseData) SetOrderSource(v string) *QueryItemOrderInfoResponseData {
	s.OrderSource = &v
	return s
}

func (s *QueryItemOrderInfoResponseData) SetGwErrorCode(v int32) *QueryItemOrderInfoResponseData {
	s.GwErrorCode = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItem struct {
	ItemOrderId     *string                                                          `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	ItemOrderStatus *int32                                                           `json:"item_order_status,omitempty" xml:"item_order_status,omitempty"`
	ItemVerifyInfo  *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo        `json:"item_verify_info,omitempty" xml:"item_verify_info,omitempty"`
	SubItemInfoList []*QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem `json:"sub_item_info_list,omitempty" xml:"sub_item_info_list,omitempty" type:"Repeated"`
	TimesCardInfo   *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo         `json:"times_card_info,omitempty" xml:"times_card_info,omitempty"`
	ValidEndTime    *int64                                                           `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	ValidStartTime  *int64                                                           `json:"valid_start_time,omitempty" xml:"valid_start_time,omitempty"`
	DeliveryTime    *int64                                                           `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}

func (s QueryItemOrderInfoResponseDataItemListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItem) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetItemOrderId(v string) *QueryItemOrderInfoResponseDataItemListItem {
	s.ItemOrderId = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetItemOrderStatus(v int32) *QueryItemOrderInfoResponseDataItemListItem {
	s.ItemOrderStatus = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetItemVerifyInfo(v *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) *QueryItemOrderInfoResponseDataItemListItem {
	s.ItemVerifyInfo = v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetSubItemInfoList(v []*QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) *QueryItemOrderInfoResponseDataItemListItem {
	s.SubItemInfoList = v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetTimesCardInfo(v *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) *QueryItemOrderInfoResponseDataItemListItem {
	s.TimesCardInfo = v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetValidEndTime(v int64) *QueryItemOrderInfoResponseDataItemListItem {
	s.ValidEndTime = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetValidStartTime(v int64) *QueryItemOrderInfoResponseDataItemListItem {
	s.ValidStartTime = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItem) SetDeliveryTime(v int64) *QueryItemOrderInfoResponseDataItemListItem {
	s.DeliveryTime = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo struct {
	DeductionOrderId   *string                                                                     `json:"deduction_order_id,omitempty" xml:"deduction_order_id,omitempty"`
	PoiId              *string                                                                     `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	VerifyCancelResult *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult `json:"verify_cancel_result,omitempty" xml:"verify_cancel_result,omitempty"`
	VerifyResult       *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult       `json:"verify_result,omitempty" xml:"verify_result,omitempty"`
	DeductCode         *string                                                                     `json:"deduct_code,omitempty" xml:"deduct_code,omitempty"`
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) SetDeductionOrderId(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo {
	s.DeductionOrderId = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) SetPoiId(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo {
	s.PoiId = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) SetVerifyCancelResult(v *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo {
	s.VerifyCancelResult = v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) SetVerifyResult(v *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo {
	s.VerifyResult = v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo) SetDeductCode(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfo {
	s.DeductCode = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult struct {
	CancelMsg  *string `json:"cancel_msg,omitempty" xml:"cancel_msg,omitempty"`
	CancelSuc  *bool   `json:"cancel_suc,omitempty" xml:"cancel_suc,omitempty"`
	CancelTime *int64  `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) SetCancelMsg(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult {
	s.CancelMsg = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) SetCancelSuc(v bool) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult {
	s.CancelSuc = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult) SetCancelTime(v int64) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyCancelResult {
	s.CancelTime = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	ResultCode *int32  `json:"result_code,omitempty" xml:"result_code,omitempty"`
	ResultMsg  *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	VerifyId   *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	VerifyTime *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) SetCode(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult {
	s.Code = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) SetResultCode(v int32) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult {
	s.ResultCode = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) SetResultMsg(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) SetVerifyId(v string) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult {
	s.VerifyId = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult) SetVerifyTime(v int64) *QueryItemOrderInfoResponseDataItemListItemItemVerifyInfoVerifyResult {
	s.VerifyTime = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem struct {
	BookEndTime   *int64  `json:"book_end_time,omitempty" xml:"book_end_time,omitempty" require:"true"`
	BookStartTime *int64  `json:"book_start_time,omitempty" xml:"book_start_time,omitempty" require:"true"`
	OriginAmount  *int64  `json:"origin_amount,omitempty" xml:"origin_amount,omitempty" require:"true"`
	SubItemId     *string `json:"sub_item_id,omitempty" xml:"sub_item_id,omitempty" require:"true"`
}

func (s QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) SetBookEndTime(v int64) *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem {
	s.BookEndTime = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) SetBookStartTime(v int64) *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem {
	s.BookStartTime = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) SetOriginAmount(v int64) *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem {
	s.OriginAmount = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem) SetSubItemId(v string) *QueryItemOrderInfoResponseDataItemListItemSubItemInfoListItem {
	s.SubItemId = &v
	return s
}

type QueryItemOrderInfoResponseDataItemListItemTimesCardInfo struct {
	UsableTimes      *int64 `json:"usable_times,omitempty" xml:"usable_times,omitempty" require:"true"`
	ActualAmountOnce *int64 `json:"actual_amount_once,omitempty" xml:"actual_amount_once,omitempty" require:"true"`
	RefundTimes      *int64 `json:"refund_times,omitempty" xml:"refund_times,omitempty" require:"true"`
	TotalTimes       *int64 `json:"total_times,omitempty" xml:"total_times,omitempty" require:"true"`
}

func (s QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) SetUsableTimes(v int64) *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo {
	s.UsableTimes = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) SetActualAmountOnce(v int64) *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo {
	s.ActualAmountOnce = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) SetRefundTimes(v int64) *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo {
	s.RefundTimes = &v
	return s
}

func (s *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo) SetTotalTimes(v int64) *QueryItemOrderInfoResponseDataItemListItemTimesCardInfo {
	s.TotalTimes = &v
	return s
}

type QueryItemOrderInfoResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s QueryItemOrderInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryItemOrderInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryItemOrderInfoResponseExtra) SetSubErrorCode(v int32) *QueryItemOrderInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryItemOrderInfoResponseExtra) SetDescription(v string) *QueryItemOrderInfoResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryItemOrderInfoResponseExtra) SetErrorCode(v int32) *QueryItemOrderInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QueryItemOrderInfoResponseExtra) SetLogid(v string) *QueryItemOrderInfoResponseExtra {
	s.Logid = &v
	return s
}

func (s *QueryItemOrderInfoResponseExtra) SetNow(v int64) *QueryItemOrderInfoResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryItemOrderInfoResponseExtra) SetSubDescription(v string) *QueryItemOrderInfoResponseExtra {
	s.SubDescription = &v
	return s
}

type QueryLiveDealDataRequest struct {
	LiveRoomId  *int64             `json:"live_room_id,omitempty" xml:"live_room_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryLiveDealDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveDealDataRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveDealDataRequest) SetLiveRoomId(v int64) *QueryLiveDealDataRequest {
	s.LiveRoomId = &v
	return s
}

func (s *QueryLiveDealDataRequest) SetHeader(v map[string]*string) *QueryLiveDealDataRequest {
	s.Header = v
	return s
}

func (s *QueryLiveDealDataRequest) SetAccessToken(v string) *QueryLiveDealDataRequest {
	s.AccessToken = &v
	return s
}

type QueryLiveDealDataResponse struct {
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryLiveDealDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryLiveDealDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveDealDataResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveDealDataResponse) SetLogId(v string) *QueryLiveDealDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryLiveDealDataResponse) SetData(v *QueryLiveDealDataResponseData) *QueryLiveDealDataResponse {
	s.Data = v
	return s
}

func (s *QueryLiveDealDataResponse) SetErrNo(v int32) *QueryLiveDealDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryLiveDealDataResponse) SetErrMsg(v string) *QueryLiveDealDataResponse {
	s.ErrMsg = &v
	return s
}

type QueryLiveDealDataResponseData struct {
	DealDataList     []*QueryLiveDealDataResponseDataDealDataListItem `json:"deal_data_list,omitempty" xml:"deal_data_list,omitempty" type:"Repeated"`
	DealOverviewData *QueryLiveDealDataResponseDataDealOverviewData   `json:"deal_overview_data,omitempty" xml:"deal_overview_data,omitempty"`
}

func (s QueryLiveDealDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveDealDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryLiveDealDataResponseData) SetDealDataList(v []*QueryLiveDealDataResponseDataDealDataListItem) *QueryLiveDealDataResponseData {
	s.DealDataList = v
	return s
}

func (s *QueryLiveDealDataResponseData) SetDealOverviewData(v *QueryLiveDealDataResponseDataDealOverviewData) *QueryLiveDealDataResponseData {
	s.DealOverviewData = v
	return s
}

type QueryLiveDealDataResponseDataDealDataListItem struct {
	PayOrderAmount           *int64  `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	ShowCount                *int64  `json:"show_count,omitempty" xml:"show_count,omitempty" require:"true"`
	PayOrderCount            *int64  `json:"pay_order_count,omitempty" xml:"pay_order_count,omitempty" require:"true"`
	OrderOncePrice           *int64  `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	CustomerOncePrice        *int64  `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
	ClickCount               *int64  `json:"click_count,omitempty" xml:"click_count,omitempty" require:"true"`
	CreateOrderAmount        *int64  `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty" require:"true"`
	RefundOrderAmount        *int64  `json:"refund_order_amount,omitempty" xml:"refund_order_amount,omitempty" require:"true"`
	CreateOrderCount         *int64  `json:"create_order_count,omitempty" xml:"create_order_count,omitempty" require:"true"`
	RefundOrderCnt           *int64  `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
	CreateUserCount          *int64  `json:"create_user_count,omitempty" xml:"create_user_count,omitempty" require:"true"`
	RefundSuccessUserCnt     *int64  `json:"refund_success_user_cnt,omitempty" xml:"refund_success_user_cnt,omitempty" require:"true"`
	SuccessUserCnt           *int64  `json:"success_user_cnt,omitempty" xml:"success_user_cnt,omitempty" require:"true"`
	Time                     *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	SaleProductCnt           *int64  `json:"sale_product_cnt,omitempty" xml:"sale_product_cnt,omitempty" require:"true"`
	RefundSuccessOrderCnt    *int64  `json:"refund_success_order_cnt,omitempty" xml:"refund_success_order_cnt,omitempty" require:"true"`
	RefundSuccessOrderAmount *int64  `json:"refund_success_order_amount,omitempty" xml:"refund_success_order_amount,omitempty" require:"true"`
}

func (s QueryLiveDealDataResponseDataDealDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveDealDataResponseDataDealDataListItem) GoString() string {
	return s.String()
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetPayOrderAmount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetShowCount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.ShowCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetPayOrderCount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.PayOrderCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetOrderOncePrice(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.OrderOncePrice = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetCustomerOncePrice(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.CustomerOncePrice = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetClickCount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.ClickCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetCreateOrderAmount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.CreateOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetRefundOrderAmount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.RefundOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetCreateOrderCount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.CreateOrderCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetRefundOrderCnt(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.RefundOrderCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetCreateUserCount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.CreateUserCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetRefundSuccessUserCnt(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.RefundSuccessUserCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetSuccessUserCnt(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.SuccessUserCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetTime(v string) *QueryLiveDealDataResponseDataDealDataListItem {
	s.Time = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetSaleProductCnt(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.SaleProductCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetRefundSuccessOrderCnt(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.RefundSuccessOrderCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealDataListItem) SetRefundSuccessOrderAmount(v int64) *QueryLiveDealDataResponseDataDealDataListItem {
	s.RefundSuccessOrderAmount = &v
	return s
}

type QueryLiveDealDataResponseDataDealOverviewData struct {
	RefundOrderCnt           *int64 `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
	CreateUserCount          *int64 `json:"create_user_count,omitempty" xml:"create_user_count,omitempty" require:"true"`
	RefundSuccessOrderAmount *int64 `json:"refund_success_order_amount,omitempty" xml:"refund_success_order_amount,omitempty" require:"true"`
	PayOrderAmount           *int64 `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	CreateOrderAmount        *int64 `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty" require:"true"`
	RefundSuccessOrderCnt    *int64 `json:"refund_success_order_cnt,omitempty" xml:"refund_success_order_cnt,omitempty" require:"true"`
	ShowCount                *int64 `json:"show_count,omitempty" xml:"show_count,omitempty" require:"true"`
	RefundOrderAmount        *int64 `json:"refund_order_amount,omitempty" xml:"refund_order_amount,omitempty" require:"true"`
	SaleProductCnt           *int64 `json:"sale_product_cnt,omitempty" xml:"sale_product_cnt,omitempty" require:"true"`
	CustomerOncePrice        *int64 `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
	OrderOncePrice           *int64 `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	RefundSuccessUserCnt     *int64 `json:"refund_success_user_cnt,omitempty" xml:"refund_success_user_cnt,omitempty" require:"true"`
	ClickCount               *int64 `json:"click_count,omitempty" xml:"click_count,omitempty" require:"true"`
	SuccessUserCnt           *int64 `json:"success_user_cnt,omitempty" xml:"success_user_cnt,omitempty" require:"true"`
	CreateOrderCount         *int64 `json:"create_order_count,omitempty" xml:"create_order_count,omitempty" require:"true"`
	PayOrderCount            *int64 `json:"pay_order_count,omitempty" xml:"pay_order_count,omitempty" require:"true"`
}

func (s QueryLiveDealDataResponseDataDealOverviewData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveDealDataResponseDataDealOverviewData) GoString() string {
	return s.String()
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetRefundOrderCnt(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.RefundOrderCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetCreateUserCount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.CreateUserCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetRefundSuccessOrderAmount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.RefundSuccessOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetPayOrderAmount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetCreateOrderAmount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.CreateOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetRefundSuccessOrderCnt(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.RefundSuccessOrderCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetShowCount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.ShowCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetRefundOrderAmount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.RefundOrderAmount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetSaleProductCnt(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.SaleProductCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetCustomerOncePrice(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.CustomerOncePrice = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetOrderOncePrice(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.OrderOncePrice = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetRefundSuccessUserCnt(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.RefundSuccessUserCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetClickCount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.ClickCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetSuccessUserCnt(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.SuccessUserCnt = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetCreateOrderCount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.CreateOrderCount = &v
	return s
}

func (s *QueryLiveDealDataResponseDataDealOverviewData) SetPayOrderCount(v int64) *QueryLiveDealDataResponseDataDealOverviewData {
	s.PayOrderCount = &v
	return s
}

type QueryLiveRoomDataRequest struct {
	LiveRoomId  *int64             `json:"live_room_id,omitempty" xml:"live_room_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryLiveRoomDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveRoomDataRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveRoomDataRequest) SetLiveRoomId(v int64) *QueryLiveRoomDataRequest {
	s.LiveRoomId = &v
	return s
}

func (s *QueryLiveRoomDataRequest) SetHeader(v map[string]*string) *QueryLiveRoomDataRequest {
	s.Header = v
	return s
}

func (s *QueryLiveRoomDataRequest) SetAccessToken(v string) *QueryLiveRoomDataRequest {
	s.AccessToken = &v
	return s
}

type QueryLiveRoomDataResponse struct {
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryLiveRoomDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryLiveRoomDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveRoomDataResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveRoomDataResponse) SetErrMsg(v string) *QueryLiveRoomDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryLiveRoomDataResponse) SetLogId(v string) *QueryLiveRoomDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryLiveRoomDataResponse) SetData(v *QueryLiveRoomDataResponseData) *QueryLiveRoomDataResponse {
	s.Data = v
	return s
}

func (s *QueryLiveRoomDataResponse) SetErrNo(v int32) *QueryLiveRoomDataResponse {
	s.ErrNo = &v
	return s
}

type QueryLiveRoomDataResponseData struct {
	LiveDataList     []*QueryLiveRoomDataResponseDataLiveDataListItem `json:"live_data_list,omitempty" xml:"live_data_list,omitempty" type:"Repeated"`
	LiveDataOverview *QueryLiveRoomDataResponseDataLiveDataOverview   `json:"live_data_overview,omitempty" xml:"live_data_overview,omitempty"`
}

func (s QueryLiveRoomDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveRoomDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryLiveRoomDataResponseData) SetLiveDataList(v []*QueryLiveRoomDataResponseDataLiveDataListItem) *QueryLiveRoomDataResponseData {
	s.LiveDataList = v
	return s
}

func (s *QueryLiveRoomDataResponseData) SetLiveDataOverview(v *QueryLiveRoomDataResponseDataLiveDataOverview) *QueryLiveRoomDataResponseData {
	s.LiveDataOverview = v
	return s
}

type QueryLiveRoomDataResponseDataLiveDataListItem struct {
	CumulativeAudienceCount *int64  `json:"cumulative_audience_count,omitempty" xml:"cumulative_audience_count,omitempty" require:"true"`
	CommentCount            *int64  `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
	ShareCount              *int64  `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	LikeTimes               *int64  `json:"like_times,omitempty" xml:"like_times,omitempty" require:"true"`
	Time                    *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s QueryLiveRoomDataResponseDataLiveDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveRoomDataResponseDataLiveDataListItem) GoString() string {
	return s.String()
}

func (s *QueryLiveRoomDataResponseDataLiveDataListItem) SetCumulativeAudienceCount(v int64) *QueryLiveRoomDataResponseDataLiveDataListItem {
	s.CumulativeAudienceCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataListItem) SetCommentCount(v int64) *QueryLiveRoomDataResponseDataLiveDataListItem {
	s.CommentCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataListItem) SetShareCount(v int64) *QueryLiveRoomDataResponseDataLiveDataListItem {
	s.ShareCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataListItem) SetLikeTimes(v int64) *QueryLiveRoomDataResponseDataLiveDataListItem {
	s.LikeTimes = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataListItem) SetTime(v string) *QueryLiveRoomDataResponseDataLiveDataListItem {
	s.Time = &v
	return s
}

type QueryLiveRoomDataResponseDataLiveDataOverview struct {
	ShareCount              *int64 `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	LikeTimes               *int64 `json:"like_times,omitempty" xml:"like_times,omitempty" require:"true"`
	CumulativeAudienceCount *int64 `json:"cumulative_audience_count,omitempty" xml:"cumulative_audience_count,omitempty" require:"true"`
	OnlineUserCount         *int64 `json:"online_user_count,omitempty" xml:"online_user_count,omitempty" require:"true"`
	PerCapitaTime           *int64 `json:"per_capita_time,omitempty" xml:"per_capita_time,omitempty" require:"true"`
	IncreasedFansCount      *int64 `json:"increased_fans_count,omitempty" xml:"increased_fans_count,omitempty" require:"true"`
	CommentCount            *int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
}

func (s QueryLiveRoomDataResponseDataLiveDataOverview) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveRoomDataResponseDataLiveDataOverview) GoString() string {
	return s.String()
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetShareCount(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.ShareCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetLikeTimes(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.LikeTimes = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetCumulativeAudienceCount(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.CumulativeAudienceCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetOnlineUserCount(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.OnlineUserCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetPerCapitaTime(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.PerCapitaTime = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetIncreasedFansCount(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.IncreasedFansCount = &v
	return s
}

func (s *QueryLiveRoomDataResponseDataLiveDataOverview) SetCommentCount(v int64) *QueryLiveRoomDataResponseDataLiveDataOverview {
	s.CommentCount = &v
	return s
}

type QueryLiveWithShortIdRequest struct {
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageNo           *int64             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	PageSize         *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	StartTime        *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName         *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	AwemeShortIdList []*string          `json:"aweme_short_id_list,omitempty" xml:"aweme_short_id_list,omitempty" type:"Repeated"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryLiveWithShortIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdRequest) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdRequest) SetAccessToken(v string) *QueryLiveWithShortIdRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetPageNo(v int64) *QueryLiveWithShortIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetPageSize(v int64) *QueryLiveWithShortIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetStartTime(v int64) *QueryLiveWithShortIdRequest {
	s.StartTime = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetEndTime(v int64) *QueryLiveWithShortIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetHostName(v string) *QueryLiveWithShortIdRequest {
	s.HostName = &v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetAwemeShortIdList(v []*string) *QueryLiveWithShortIdRequest {
	s.AwemeShortIdList = v
	return s
}

func (s *QueryLiveWithShortIdRequest) SetHeader(v map[string]*string) *QueryLiveWithShortIdRequest {
	s.Header = v
	return s
}

type QueryLiveWithShortIdResponse struct {
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryLiveWithShortIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryLiveWithShortIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdResponse) SetErrMsg(v string) *QueryLiveWithShortIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryLiveWithShortIdResponse) SetLogId(v string) *QueryLiveWithShortIdResponse {
	s.LogId = &v
	return s
}

func (s *QueryLiveWithShortIdResponse) SetData(v *QueryLiveWithShortIdResponseData) *QueryLiveWithShortIdResponse {
	s.Data = v
	return s
}

func (s *QueryLiveWithShortIdResponse) SetErrNo(v int32) *QueryLiveWithShortIdResponse {
	s.ErrNo = &v
	return s
}

type QueryLiveWithShortIdResponseData struct {
	ShortDataList []*QueryLiveWithShortIdResponseDataShortDataListItem `json:"ShortDataList,omitempty" xml:"ShortDataList,omitempty" type:"Repeated"`
	Total         *int64                                               `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
}

func (s QueryLiveWithShortIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdResponseData) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdResponseData) SetShortDataList(v []*QueryLiveWithShortIdResponseDataShortDataListItem) *QueryLiveWithShortIdResponseData {
	s.ShortDataList = v
	return s
}

func (s *QueryLiveWithShortIdResponseData) SetTotal(v int64) *QueryLiveWithShortIdResponseData {
	s.Total = &v
	return s
}

type QueryLiveWithShortIdResponseDataShortDataListItem struct {
	ShortInfo *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo `json:"ShortInfo,omitempty" xml:"ShortInfo,omitempty"`
	Data      *QueryLiveWithShortIdResponseDataShortDataListItemData      `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryLiveWithShortIdResponseDataShortDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdResponseDataShortDataListItem) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItem) SetShortInfo(v *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) *QueryLiveWithShortIdResponseDataShortDataListItem {
	s.ShortInfo = v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItem) SetData(v *QueryLiveWithShortIdResponseDataShortDataListItemData) *QueryLiveWithShortIdResponseDataShortDataListItem {
	s.Data = v
	return s
}

type QueryLiveWithShortIdResponseDataShortDataListItemData struct {
	PayCustomerCnt         *int64 `json:"PayCustomerCnt,omitempty" xml:"PayCustomerCnt,omitempty" require:"true"`
	PayOrderAmount         *int64 `json:"PayOrderAmount,omitempty" xml:"PayOrderAmount,omitempty" require:"true"`
	CustomerOncePrice      *int64 `json:"CustomerOncePrice,omitempty" xml:"CustomerOncePrice,omitempty" require:"true"`
	RoomWatchRate          *int64 `json:"RoomWatchRate,omitempty" xml:"RoomWatchRate,omitempty" require:"true"`
	RoomNewFans            *int64 `json:"RoomNewFans,omitempty" xml:"RoomNewFans,omitempty" require:"true"`
	RoomDuration           *int64 `json:"RoomDuration,omitempty" xml:"RoomDuration,omitempty" require:"true"`
	RoomSessionCnt         *int64 `json:"RoomSessionCnt,omitempty" xml:"RoomSessionCnt,omitempty" require:"true"`
	RoomHighestOnlineCount *int64 `json:"RoomHighestOnlineCount,omitempty" xml:"RoomHighestOnlineCount,omitempty" require:"true"`
}

func (s QueryLiveWithShortIdResponseDataShortDataListItemData) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdResponseDataShortDataListItemData) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetPayCustomerCnt(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.PayCustomerCnt = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetPayOrderAmount(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetCustomerOncePrice(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.CustomerOncePrice = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetRoomWatchRate(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.RoomWatchRate = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetRoomNewFans(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.RoomNewFans = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetRoomDuration(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.RoomDuration = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetRoomSessionCnt(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.RoomSessionCnt = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemData) SetRoomHighestOnlineCount(v int64) *QueryLiveWithShortIdResponseDataShortDataListItemData {
	s.RoomHighestOnlineCount = &v
	return s
}

type QueryLiveWithShortIdResponseDataShortDataListItemShortInfo struct {
	Nickname     *string `json:"Nickname,omitempty" xml:"Nickname,omitempty" require:"true"`
	Avat         *string `json:"Avat,omitempty" xml:"Avat,omitempty" require:"true"`
	AwemeShortID *string `json:"AwemeShortID,omitempty" xml:"AwemeShortID,omitempty" require:"true"`
	AccountType  *int    `json:"AccountType,omitempty" xml:"AccountType,omitempty" require:"true"`
}

func (s QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) GoString() string {
	return s.String()
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) SetNickname(v string) *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo {
	s.Nickname = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) SetAvat(v string) *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo {
	s.Avat = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) SetAwemeShortID(v string) *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo {
	s.AwemeShortID = &v
	return s
}

func (s *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo) SetAccountType(v int) *QueryLiveWithShortIdResponseDataShortDataListItemShortInfo {
	s.AccountType = &v
	return s
}

type QueryMaSubServiceRequest struct {
	PageSize      *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	ApprovalState *int64             `json:"approval_state,omitempty" xml:"approval_state,omitempty"`
	PageNo        *int64             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryMaSubServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMaSubServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryMaSubServiceRequest) SetPageSize(v int64) *QueryMaSubServiceRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMaSubServiceRequest) SetApprovalState(v int64) *QueryMaSubServiceRequest {
	s.ApprovalState = &v
	return s
}

func (s *QueryMaSubServiceRequest) SetPageNo(v int64) *QueryMaSubServiceRequest {
	s.PageNo = &v
	return s
}

func (s *QueryMaSubServiceRequest) SetHeader(v map[string]*string) *QueryMaSubServiceRequest {
	s.Header = v
	return s
}

func (s *QueryMaSubServiceRequest) SetAccessToken(v string) *QueryMaSubServiceRequest {
	s.AccessToken = &v
	return s
}

type QueryMaSubServiceResponse struct {
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryMaSubServiceResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryMaSubServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMaSubServiceResponse) GoString() string {
	return s.String()
}

func (s *QueryMaSubServiceResponse) SetErrMsg(v string) *QueryMaSubServiceResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryMaSubServiceResponse) SetLogId(v string) *QueryMaSubServiceResponse {
	s.LogId = &v
	return s
}

func (s *QueryMaSubServiceResponse) SetData(v *QueryMaSubServiceResponseData) *QueryMaSubServiceResponse {
	s.Data = v
	return s
}

func (s *QueryMaSubServiceResponse) SetErrNo(v int32) *QueryMaSubServiceResponse {
	s.ErrNo = &v
	return s
}

type QueryMaSubServiceResponseData struct {
	Total                *int64                                                   `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	MaSubServiceInfoList []*QueryMaSubServiceResponseDataMaSubServiceInfoListItem `json:"ma_sub_service_info_list,omitempty" xml:"ma_sub_service_info_list,omitempty" type:"Repeated"`
	HasMore              *bool                                                    `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
}

func (s QueryMaSubServiceResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryMaSubServiceResponseData) GoString() string {
	return s.String()
}

func (s *QueryMaSubServiceResponseData) SetTotal(v int64) *QueryMaSubServiceResponseData {
	s.Total = &v
	return s
}

func (s *QueryMaSubServiceResponseData) SetMaSubServiceInfoList(v []*QueryMaSubServiceResponseDataMaSubServiceInfoListItem) *QueryMaSubServiceResponseData {
	s.MaSubServiceInfoList = v
	return s
}

func (s *QueryMaSubServiceResponseData) SetHasMore(v bool) *QueryMaSubServiceResponseData {
	s.HasMore = &v
	return s
}

type QueryMaSubServiceResponseDataMaSubServiceInfoListItem struct {
	MicroAppIcon   *string                                                             `json:"micro_app_icon,omitempty" xml:"micro_app_icon,omitempty" require:"true"`
	MicroAppId     *string                                                             `json:"micro_app_id,omitempty" xml:"micro_app_id,omitempty" require:"true"`
	SubServiceId   *string                                                             `json:"sub_service_id,omitempty" xml:"sub_service_id,omitempty" require:"true"`
	SearchKeyWord  []*string                                                           `json:"search_key_word,omitempty" xml:"search_key_word,omitempty" require:"true" type:"Repeated"`
	SubServiceName *string                                                             `json:"sub_service_name,omitempty" xml:"sub_service_name,omitempty" require:"true"`
	CreateTime     *string                                                             `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	Reasons        []*QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem `json:"reasons,omitempty" xml:"reasons,omitempty" type:"Repeated"`
	StartPageUrl   *string                                                             `json:"start_page_url,omitempty" xml:"start_page_url,omitempty" require:"true"`
	ApprovalState  *int64                                                              `json:"approval_state,omitempty" xml:"approval_state,omitempty" require:"true"`
}

func (s QueryMaSubServiceResponseDataMaSubServiceInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryMaSubServiceResponseDataMaSubServiceInfoListItem) GoString() string {
	return s.String()
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetMicroAppIcon(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.MicroAppIcon = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetMicroAppId(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.MicroAppId = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetSubServiceId(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.SubServiceId = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetSearchKeyWord(v []*string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.SearchKeyWord = v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetSubServiceName(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.SubServiceName = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetCreateTime(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.CreateTime = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetReasons(v []*QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.Reasons = v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetStartPageUrl(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.StartPageUrl = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItem) SetApprovalState(v int64) *QueryMaSubServiceResponseDataMaSubServiceInfoListItem {
	s.ApprovalState = &v
	return s
}

type QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem struct {
	AuditField *string `json:"audit_field,omitempty" xml:"audit_field,omitempty" require:"true"`
	Reason     *string `json:"reason,omitempty" xml:"reason,omitempty" require:"true"`
}

func (s QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem) GoString() string {
	return s.String()
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem) SetAuditField(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem {
	s.AuditField = &v
	return s
}

func (s *QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem) SetReason(v string) *QueryMaSubServiceResponseDataMaSubServiceInfoListItemReasonsItem {
	s.Reason = &v
	return s
}

type QueryOrderRequest struct {
	PoiId       *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	BookId      *string            `json:"book_id,omitempty" xml:"book_id,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s QueryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) SetPoiId(v string) *QueryOrderRequest {
	s.PoiId = &v
	return s
}

func (s *QueryOrderRequest) SetAccountId(v string) *QueryOrderRequest {
	s.AccountId = &v
	return s
}

func (s *QueryOrderRequest) SetBookId(v string) *QueryOrderRequest {
	s.BookId = &v
	return s
}

func (s *QueryOrderRequest) SetAccessToken(v string) *QueryOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryOrderRequest) SetHeader(v map[string]*string) *QueryOrderRequest {
	s.Header = v
	return s
}

func (s *QueryOrderRequest) SetEndTime(v int64) *QueryOrderRequest {
	s.EndTime = &v
	return s
}

func (s *QueryOrderRequest) SetPageNum(v int32) *QueryOrderRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOrderRequest) SetPageSize(v int32) *QueryOrderRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrderRequest) SetStartTime(v int64) *QueryOrderRequest {
	s.StartTime = &v
	return s
}

type QueryOrderResponse struct {
	Extra *QueryOrderResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *QueryOrderResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) SetExtra(v *QueryOrderResponseExtra) *QueryOrderResponse {
	s.Extra = v
	return s
}

func (s *QueryOrderResponse) SetData(v *QueryOrderResponseData) *QueryOrderResponse {
	s.Data = v
	return s
}

type QueryOrderResponseData struct {
	OrderList     []*QueryOrderResponseDataOrderListItem `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
	Total         *int32                                 `json:"total,omitempty" xml:"total,omitempty"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s QueryOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseData) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseData) SetOrderList(v []*QueryOrderResponseDataOrderListItem) *QueryOrderResponseData {
	s.OrderList = v
	return s
}

func (s *QueryOrderResponseData) SetTotal(v int32) *QueryOrderResponseData {
	s.Total = &v
	return s
}

func (s *QueryOrderResponseData) SetGwErrorCode(v int32) *QueryOrderResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QueryOrderResponseData) SetGwDescription(v string) *QueryOrderResponseData {
	s.GwDescription = &v
	return s
}

type QueryOrderResponseDataOrderListItem struct {
	Status           *int                                                       `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	UserInfos        []*QueryOrderResponseDataOrderListItemUserInfosItem        `json:"user_infos,omitempty" xml:"user_infos,omitempty" type:"Repeated"`
	BookId           *string                                                    `json:"book_id,omitempty" xml:"book_id,omitempty" require:"true"`
	BookTakingTime   *int64                                                     `json:"book_taking_time,omitempty" xml:"book_taking_time,omitempty" require:"true"`
	BookTime         *int64                                                     `json:"book_time,omitempty" xml:"book_time,omitempty" require:"true"`
	BookingInfo      *QueryOrderResponseDataOrderListItemBookingInfo            `json:"booking_info,omitempty" xml:"booking_info,omitempty"`
	IndustryInfoList []*QueryOrderResponseDataOrderListItemIndustryInfoListItem `json:"industry_info_list,omitempty" xml:"industry_info_list,omitempty" type:"Repeated"`
	Products         []*QueryOrderResponseDataOrderListItemProductsItem         `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
}

func (s QueryOrderResponseDataOrderListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseDataOrderListItem) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseDataOrderListItem) SetStatus(v int) *QueryOrderResponseDataOrderListItem {
	s.Status = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetUserInfos(v []*QueryOrderResponseDataOrderListItemUserInfosItem) *QueryOrderResponseDataOrderListItem {
	s.UserInfos = v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetBookId(v string) *QueryOrderResponseDataOrderListItem {
	s.BookId = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetBookTakingTime(v int64) *QueryOrderResponseDataOrderListItem {
	s.BookTakingTime = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetBookTime(v int64) *QueryOrderResponseDataOrderListItem {
	s.BookTime = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetBookingInfo(v *QueryOrderResponseDataOrderListItemBookingInfo) *QueryOrderResponseDataOrderListItem {
	s.BookingInfo = v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetIndustryInfoList(v []*QueryOrderResponseDataOrderListItemIndustryInfoListItem) *QueryOrderResponseDataOrderListItem {
	s.IndustryInfoList = v
	return s
}

func (s *QueryOrderResponseDataOrderListItem) SetProducts(v []*QueryOrderResponseDataOrderListItemProductsItem) *QueryOrderResponseDataOrderListItem {
	s.Products = v
	return s
}

type QueryOrderResponseDataOrderListItemBookingInfo struct {
	UserNotes *string   `json:"user_notes,omitempty" xml:"user_notes,omitempty"`
	Count     *int32    `json:"count,omitempty" xml:"count,omitempty"`
	ItemIds   []*string `json:"item_ids,omitempty" xml:"item_ids,omitempty" type:"Repeated"`
	PoiId     *string   `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s QueryOrderResponseDataOrderListItemBookingInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseDataOrderListItemBookingInfo) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseDataOrderListItemBookingInfo) SetUserNotes(v string) *QueryOrderResponseDataOrderListItemBookingInfo {
	s.UserNotes = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemBookingInfo) SetCount(v int32) *QueryOrderResponseDataOrderListItemBookingInfo {
	s.Count = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemBookingInfo) SetItemIds(v []*string) *QueryOrderResponseDataOrderListItemBookingInfo {
	s.ItemIds = v
	return s
}

func (s *QueryOrderResponseDataOrderListItemBookingInfo) SetPoiId(v string) *QueryOrderResponseDataOrderListItemBookingInfo {
	s.PoiId = &v
	return s
}

type QueryOrderResponseDataOrderListItemIndustryInfoListItem struct {
	AttrValue   *string `json:"attr_value,omitempty" xml:"attr_value,omitempty" require:"true"`
	AttrKey     *string `json:"attr_key,omitempty" xml:"attr_key,omitempty" require:"true"`
	AttrKeyName *string `json:"attr_key_name,omitempty" xml:"attr_key_name,omitempty"`
}

func (s QueryOrderResponseDataOrderListItemIndustryInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseDataOrderListItemIndustryInfoListItem) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseDataOrderListItemIndustryInfoListItem) SetAttrValue(v string) *QueryOrderResponseDataOrderListItemIndustryInfoListItem {
	s.AttrValue = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemIndustryInfoListItem) SetAttrKey(v string) *QueryOrderResponseDataOrderListItemIndustryInfoListItem {
	s.AttrKey = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemIndustryInfoListItem) SetAttrKeyName(v string) *QueryOrderResponseDataOrderListItemIndustryInfoListItem {
	s.AttrKeyName = &v
	return s
}

type QueryOrderResponseDataOrderListItemProductsItem struct {
	SkuId          *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OriginalAmount *string `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	SkuAme         *string `json:"sku_ame,omitempty" xml:"sku_ame,omitempty"`
}

func (s QueryOrderResponseDataOrderListItemProductsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseDataOrderListItemProductsItem) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseDataOrderListItemProductsItem) SetSkuId(v string) *QueryOrderResponseDataOrderListItemProductsItem {
	s.SkuId = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemProductsItem) SetOriginalAmount(v string) *QueryOrderResponseDataOrderListItemProductsItem {
	s.OriginalAmount = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemProductsItem) SetSkuAme(v string) *QueryOrderResponseDataOrderListItemProductsItem {
	s.SkuAme = &v
	return s
}

type QueryOrderResponseDataOrderListItemUserInfosItem struct {
	Phone         *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Sex           *int32  `json:"sex,omitempty" xml:"sex,omitempty"`
	Age           *int32  `json:"age,omitempty" xml:"age,omitempty"`
	LicenseNumber *string `json:"license_number,omitempty" xml:"license_number,omitempty"`
	LicenseType   *int32  `json:"license_type,omitempty" xml:"license_type,omitempty"`
	MaritalStatus *int32  `json:"marital_status,omitempty" xml:"marital_status,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryOrderResponseDataOrderListItemUserInfosItem) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseDataOrderListItemUserInfosItem) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetPhone(v string) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.Phone = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetSex(v int32) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.Sex = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetAge(v int32) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.Age = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetLicenseNumber(v string) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.LicenseNumber = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetLicenseType(v int32) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.LicenseType = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetMaritalStatus(v int32) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.MaritalStatus = &v
	return s
}

func (s *QueryOrderResponseDataOrderListItemUserInfosItem) SetName(v string) *QueryOrderResponseDataOrderListItemUserInfosItem {
	s.Name = &v
	return s
}

type QueryOrderResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s QueryOrderResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseExtra) SetLogid(v string) *QueryOrderResponseExtra {
	s.Logid = &v
	return s
}

func (s *QueryOrderResponseExtra) SetNow(v int64) *QueryOrderResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryOrderResponseExtra) SetSubDescription(v string) *QueryOrderResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryOrderResponseExtra) SetSubErrorCode(v int32) *QueryOrderResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryOrderResponseExtra) SetDescription(v string) *QueryOrderResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryOrderResponseExtra) SetErrorCode(v int32) *QueryOrderResponseExtra {
	s.ErrorCode = &v
	return s
}

type QueryProductDealDataRequest struct {
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryProductDealDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductDealDataRequest) GoString() string {
	return s.String()
}

func (s *QueryProductDealDataRequest) SetEndTime(v int64) *QueryProductDealDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryProductDealDataRequest) SetHostName(v string) *QueryProductDealDataRequest {
	s.HostName = &v
	return s
}

func (s *QueryProductDealDataRequest) SetPageNum(v int64) *QueryProductDealDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryProductDealDataRequest) SetPageSize(v int64) *QueryProductDealDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryProductDealDataRequest) SetStartTime(v int64) *QueryProductDealDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryProductDealDataRequest) SetHeader(v map[string]*string) *QueryProductDealDataRequest {
	s.Header = v
	return s
}

func (s *QueryProductDealDataRequest) SetAccessToken(v string) *QueryProductDealDataRequest {
	s.AccessToken = &v
	return s
}

type QueryProductDealDataResponse struct {
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryProductDealDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryProductDealDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductDealDataResponse) GoString() string {
	return s.String()
}

func (s *QueryProductDealDataResponse) SetLogId(v string) *QueryProductDealDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryProductDealDataResponse) SetData(v *QueryProductDealDataResponseData) *QueryProductDealDataResponse {
	s.Data = v
	return s
}

func (s *QueryProductDealDataResponse) SetErrNo(v int32) *QueryProductDealDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryProductDealDataResponse) SetErrMsg(v string) *QueryProductDealDataResponse {
	s.ErrMsg = &v
	return s
}

type QueryProductDealDataResponseData struct {
	Total        *int64                                              `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	DealDataList []*QueryProductDealDataResponseDataDealDataListItem `json:"deal_data_list,omitempty" xml:"deal_data_list,omitempty" type:"Repeated"`
}

func (s QueryProductDealDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductDealDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryProductDealDataResponseData) SetTotal(v int64) *QueryProductDealDataResponseData {
	s.Total = &v
	return s
}

func (s *QueryProductDealDataResponseData) SetDealDataList(v []*QueryProductDealDataResponseDataDealDataListItem) *QueryProductDealDataResponseData {
	s.DealDataList = v
	return s
}

type QueryProductDealDataResponseDataDealDataListItem struct {
	ProductName   *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	ProductId     *int64  `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	PayAmount     *int64  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PayProductCnt *int64  `json:"pay_product_cnt,omitempty" xml:"pay_product_cnt,omitempty" require:"true"`
	PayUserCnt    *int64  `json:"pay_user_cnt,omitempty" xml:"pay_user_cnt,omitempty" require:"true"`
}

func (s QueryProductDealDataResponseDataDealDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryProductDealDataResponseDataDealDataListItem) GoString() string {
	return s.String()
}

func (s *QueryProductDealDataResponseDataDealDataListItem) SetProductName(v string) *QueryProductDealDataResponseDataDealDataListItem {
	s.ProductName = &v
	return s
}

func (s *QueryProductDealDataResponseDataDealDataListItem) SetProductId(v int64) *QueryProductDealDataResponseDataDealDataListItem {
	s.ProductId = &v
	return s
}

func (s *QueryProductDealDataResponseDataDealDataListItem) SetPayAmount(v int64) *QueryProductDealDataResponseDataDealDataListItem {
	s.PayAmount = &v
	return s
}

func (s *QueryProductDealDataResponseDataDealDataListItem) SetPayProductCnt(v int64) *QueryProductDealDataResponseDataDealDataListItem {
	s.PayProductCnt = &v
	return s
}

func (s *QueryProductDealDataResponseDataDealDataListItem) SetPayUserCnt(v int64) *QueryProductDealDataResponseDataDealDataListItem {
	s.PayUserCnt = &v
	return s
}

type QueryRealTimeUserDataRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
}

func (s QueryRealTimeUserDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRealTimeUserDataRequest) GoString() string {
	return s.String()
}

func (s *QueryRealTimeUserDataRequest) SetHeader(v map[string]*string) *QueryRealTimeUserDataRequest {
	s.Header = v
	return s
}

func (s *QueryRealTimeUserDataRequest) SetAccessToken(v string) *QueryRealTimeUserDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryRealTimeUserDataRequest) SetHostName(v string) *QueryRealTimeUserDataRequest {
	s.HostName = &v
	return s
}

func (s *QueryRealTimeUserDataRequest) SetVersionType(v string) *QueryRealTimeUserDataRequest {
	s.VersionType = &v
	return s
}

type QueryRealTimeUserDataResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryRealTimeUserDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryRealTimeUserDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRealTimeUserDataResponse) GoString() string {
	return s.String()
}

func (s *QueryRealTimeUserDataResponse) SetLogId(v string) *QueryRealTimeUserDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryRealTimeUserDataResponse) SetData(v *QueryRealTimeUserDataResponseData) *QueryRealTimeUserDataResponse {
	s.Data = v
	return s
}

func (s *QueryRealTimeUserDataResponse) SetErrNo(v int32) *QueryRealTimeUserDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryRealTimeUserDataResponse) SetErrMsg(v string) *QueryRealTimeUserDataResponse {
	s.ErrMsg = &v
	return s
}

type QueryRealTimeUserDataResponseData struct {
	Behaviors []*QueryRealTimeUserDataResponseDataBehaviorsItem `json:"behaviors,omitempty" xml:"behaviors,omitempty" type:"Repeated"`
	Sum       *QueryRealTimeUserDataResponseDataSum             `json:"sum,omitempty" xml:"sum,omitempty"`
}

func (s QueryRealTimeUserDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryRealTimeUserDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryRealTimeUserDataResponseData) SetBehaviors(v []*QueryRealTimeUserDataResponseDataBehaviorsItem) *QueryRealTimeUserDataResponseData {
	s.Behaviors = v
	return s
}

func (s *QueryRealTimeUserDataResponseData) SetSum(v *QueryRealTimeUserDataResponseDataSum) *QueryRealTimeUserDataResponseData {
	s.Sum = v
	return s
}

type QueryRealTimeUserDataResponseDataBehaviorsItem struct {
	ActiveUserNum *int64  `json:"active_user_num,omitempty" xml:"active_user_num,omitempty" require:"true"`
	OpenTime      *int64  `json:"open_time,omitempty" xml:"open_time,omitempty"`
	Time          *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s QueryRealTimeUserDataResponseDataBehaviorsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryRealTimeUserDataResponseDataBehaviorsItem) GoString() string {
	return s.String()
}

func (s *QueryRealTimeUserDataResponseDataBehaviorsItem) SetActiveUserNum(v int64) *QueryRealTimeUserDataResponseDataBehaviorsItem {
	s.ActiveUserNum = &v
	return s
}

func (s *QueryRealTimeUserDataResponseDataBehaviorsItem) SetOpenTime(v int64) *QueryRealTimeUserDataResponseDataBehaviorsItem {
	s.OpenTime = &v
	return s
}

func (s *QueryRealTimeUserDataResponseDataBehaviorsItem) SetTime(v string) *QueryRealTimeUserDataResponseDataBehaviorsItem {
	s.Time = &v
	return s
}

type QueryRealTimeUserDataResponseDataSum struct {
	ActiveUserNum *int64 `json:"active_user_num,omitempty" xml:"active_user_num,omitempty"`
	OpenTime      *int64 `json:"open_time,omitempty" xml:"open_time,omitempty"`
}

func (s QueryRealTimeUserDataResponseDataSum) String() string {
	return tea.Prettify(s)
}

func (s QueryRealTimeUserDataResponseDataSum) GoString() string {
	return s.String()
}

func (s *QueryRealTimeUserDataResponseDataSum) SetActiveUserNum(v int64) *QueryRealTimeUserDataResponseDataSum {
	s.ActiveUserNum = &v
	return s
}

func (s *QueryRealTimeUserDataResponseDataSum) SetOpenTime(v int64) *QueryRealTimeUserDataResponseDataSum {
	s.OpenTime = &v
	return s
}

type QueryRecordByCertRequest struct {
	CertificateIds []*string          `json:"certificate_ids,omitempty" xml:"certificate_ids,omitempty" require:"true" type:"Repeated"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryRecordByCertRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertRequest) SetCertificateIds(v []*string) *QueryRecordByCertRequest {
	s.CertificateIds = v
	return s
}

func (s *QueryRecordByCertRequest) SetHeader(v map[string]*string) *QueryRecordByCertRequest {
	s.Header = v
	return s
}

func (s *QueryRecordByCertRequest) SetAccessToken(v string) *QueryRecordByCertRequest {
	s.AccessToken = &v
	return s
}

type QueryRecordByCertResponse struct {
	Data  *QueryRecordByCertResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *QueryRecordByCertResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s QueryRecordByCertResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertResponse) SetData(v *QueryRecordByCertResponseData) *QueryRecordByCertResponse {
	s.Data = v
	return s
}

func (s *QueryRecordByCertResponse) SetExtra(v *QueryRecordByCertResponseExtra) *QueryRecordByCertResponse {
	s.Extra = v
	return s
}

type QueryRecordByCertResponseData struct {
	GwErrorCode   *int32                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Records       []*QueryRecordByCertResponseDataRecordsItem `json:"records,omitempty" xml:"records,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRecordByCertResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertResponseData) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertResponseData) SetGwErrorCode(v int32) *QueryRecordByCertResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QueryRecordByCertResponseData) SetGwDescription(v string) *QueryRecordByCertResponseData {
	s.GwDescription = &v
	return s
}

func (s *QueryRecordByCertResponseData) SetRecords(v []*QueryRecordByCertResponseDataRecordsItem) *QueryRecordByCertResponseData {
	s.Records = v
	return s
}

type QueryRecordByCertResponseDataRecordsItem struct {
	OrderId       *string                                      `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Amount        map[string]*int64                            `json:"amount,omitempty" xml:"amount,omitempty"`
	VerifyId      *string                                      `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	FundAmount    map[string]*int64                            `json:"fund_amount,omitempty" xml:"fund_amount,omitempty"`
	VerifyTime    *int64                                       `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	IsMallStore   *bool                                        `json:"is_mall_store,omitempty" xml:"is_mall_store,omitempty"`
	Code          *string                                      `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Status        *int                                         `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CertificateId *string                                      `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Cursor        *string                                      `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	LedgerId      *string                                      `json:"ledger_id,omitempty" xml:"ledger_id,omitempty" require:"true"`
	Sku           *QueryRecordByCertResponseDataRecordsItemSku `json:"sku,omitempty" xml:"sku,omitempty"`
}

func (s QueryRecordByCertResponseDataRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertResponseDataRecordsItem) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetOrderId(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.OrderId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetAmount(v map[string]*int64) *QueryRecordByCertResponseDataRecordsItem {
	s.Amount = v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetVerifyId(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetFundAmount(v map[string]*int64) *QueryRecordByCertResponseDataRecordsItem {
	s.FundAmount = v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetVerifyTime(v int64) *QueryRecordByCertResponseDataRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetIsMallStore(v bool) *QueryRecordByCertResponseDataRecordsItem {
	s.IsMallStore = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetCode(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.Code = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetStatus(v int) *QueryRecordByCertResponseDataRecordsItem {
	s.Status = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetCertificateId(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetCursor(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.Cursor = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetLedgerId(v string) *QueryRecordByCertResponseDataRecordsItem {
	s.LedgerId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItem) SetSku(v *QueryRecordByCertResponseDataRecordsItemSku) *QueryRecordByCertResponseDataRecordsItem {
	s.Sku = v
	return s
}

type QueryRecordByCertResponseDataRecordsItemSku struct {
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	ThirdSkuId    *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryRecordByCertResponseDataRecordsItemSku) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertResponseDataRecordsItemSku) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertResponseDataRecordsItemSku) SetMarketPrice(v int64) *QueryRecordByCertResponseDataRecordsItemSku {
	s.MarketPrice = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItemSku) SetSkuId(v string) *QueryRecordByCertResponseDataRecordsItemSku {
	s.SkuId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItemSku) SetSoldStartTime(v int64) *QueryRecordByCertResponseDataRecordsItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItemSku) SetThirdSkuId(v string) *QueryRecordByCertResponseDataRecordsItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *QueryRecordByCertResponseDataRecordsItemSku) SetTitle(v string) *QueryRecordByCertResponseDataRecordsItemSku {
	s.Title = &v
	return s
}

type QueryRecordByCertResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s QueryRecordByCertResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordByCertResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryRecordByCertResponseExtra) SetNow(v int64) *QueryRecordByCertResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryRecordByCertResponseExtra) SetSubDescription(v string) *QueryRecordByCertResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryRecordByCertResponseExtra) SetSubErrorCode(v int32) *QueryRecordByCertResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryRecordByCertResponseExtra) SetDescription(v string) *QueryRecordByCertResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryRecordByCertResponseExtra) SetErrorCode(v int32) *QueryRecordByCertResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QueryRecordByCertResponseExtra) SetLogid(v string) *QueryRecordByCertResponseExtra {
	s.Logid = &v
	return s
}

type QueryRecordBySubfulfilRequest struct {
	SubFulfilId []*string          `json:"sub_fulfil_id,omitempty" xml:"sub_fulfil_id,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryRecordBySubfulfilRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilRequest) SetSubFulfilId(v []*string) *QueryRecordBySubfulfilRequest {
	s.SubFulfilId = v
	return s
}

func (s *QueryRecordBySubfulfilRequest) SetHeader(v map[string]*string) *QueryRecordBySubfulfilRequest {
	s.Header = v
	return s
}

func (s *QueryRecordBySubfulfilRequest) SetAccessToken(v string) *QueryRecordBySubfulfilRequest {
	s.AccessToken = &v
	return s
}

type QueryRecordBySubfulfilResponse struct {
	Extra *QueryRecordBySubfulfilResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *QueryRecordBySubfulfilResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryRecordBySubfulfilResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilResponse) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilResponse) SetExtra(v *QueryRecordBySubfulfilResponseExtra) *QueryRecordBySubfulfilResponse {
	s.Extra = v
	return s
}

func (s *QueryRecordBySubfulfilResponse) SetData(v *QueryRecordBySubfulfilResponseData) *QueryRecordBySubfulfilResponse {
	s.Data = v
	return s
}

type QueryRecordBySubfulfilResponseData struct {
	GwErrorCode   *int32                                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Records       []*QueryRecordBySubfulfilResponseDataRecordsItem `json:"records,omitempty" xml:"records,omitempty" require:"true" type:"Repeated"`
}

func (s QueryRecordBySubfulfilResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilResponseData) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilResponseData) SetGwErrorCode(v int32) *QueryRecordBySubfulfilResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseData) SetGwDescription(v string) *QueryRecordBySubfulfilResponseData {
	s.GwDescription = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseData) SetRecords(v []*QueryRecordBySubfulfilResponseDataRecordsItem) *QueryRecordBySubfulfilResponseData {
	s.Records = v
	return s
}

type QueryRecordBySubfulfilResponseDataRecordsItem struct {
	BizTime              *int64                                                                        `json:"biz_time,omitempty" xml:"biz_time,omitempty" require:"true"`
	CommissionBaseAmount *int64                                                                        `json:"commission_base_amount,omitempty" xml:"commission_base_amount,omitempty" require:"true"`
	PayAmount            *int64                                                                        `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	AccountName          *string                                                                       `json:"account_name,omitempty" xml:"account_name,omitempty" require:"true"`
	LedgerDatailMap      map[string]*QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue `json:"ledger_datail_map,omitempty" xml:"ledger_datail_map,omitempty" require:"true"`
	SubFulfilId          *string                                                                       `json:"sub_fulfil_id,omitempty" xml:"sub_fulfil_id,omitempty" require:"true"`
	LedgerId             *string                                                                       `json:"ledger_id,omitempty" xml:"ledger_id,omitempty" require:"true"`
	SkuType              *int32                                                                        `json:"sku_type,omitempty" xml:"sku_type,omitempty" require:"true"`
	ShopOrderId          *string                                                                       `json:"shop_order_id,omitempty" xml:"shop_order_id,omitempty" require:"true"`
	WithdrawId           *string                                                                       `json:"withdraw_id,omitempty" xml:"withdraw_id,omitempty" require:"true"`
	LedgerPlatformTicket *int64                                                                        `json:"ledger_platform_ticket,omitempty" xml:"ledger_platform_ticket,omitempty"`
	ItemOrderId          *string                                                                       `json:"item_order_id,omitempty" xml:"item_order_id,omitempty" require:"true"`
	AccountId            *string                                                                       `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	SkuId                *string                                                                       `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
}

func (s QueryRecordBySubfulfilResponseDataRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilResponseDataRecordsItem) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetBizTime(v int64) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.BizTime = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetCommissionBaseAmount(v int64) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.CommissionBaseAmount = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetPayAmount(v int64) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.PayAmount = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetAccountName(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.AccountName = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetLedgerDatailMap(v map[string]*QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.LedgerDatailMap = v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetSubFulfilId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.SubFulfilId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetLedgerId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.LedgerId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetSkuType(v int32) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.SkuType = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetShopOrderId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.ShopOrderId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetWithdrawId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.WithdrawId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetLedgerPlatformTicket(v int64) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.LedgerPlatformTicket = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetItemOrderId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.ItemOrderId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetAccountId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.AccountId = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItem) SetSkuId(v string) *QueryRecordBySubfulfilResponseDataRecordsItem {
	s.SkuId = &v
	return s
}

type QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue struct {
	Amount          *int64 `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	CommissionRatio *int64 `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
}

func (s QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue) SetAmount(v int64) *QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue {
	s.Amount = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue) SetCommissionRatio(v int64) *QueryRecordBySubfulfilResponseDataRecordsItemLedgerDatailMapValue {
	s.CommissionRatio = &v
	return s
}

type QueryRecordBySubfulfilResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s QueryRecordBySubfulfilResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryRecordBySubfulfilResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryRecordBySubfulfilResponseExtra) SetNow(v int64) *QueryRecordBySubfulfilResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseExtra) SetSubDescription(v string) *QueryRecordBySubfulfilResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseExtra) SetSubErrorCode(v int32) *QueryRecordBySubfulfilResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseExtra) SetDescription(v string) *QueryRecordBySubfulfilResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseExtra) SetErrorCode(v int32) *QueryRecordBySubfulfilResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QueryRecordBySubfulfilResponseExtra) SetLogid(v string) *QueryRecordBySubfulfilResponseExtra {
	s.Logid = &v
	return s
}

type QuerySearchTagListRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QuerySearchTagListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySearchTagListRequest) GoString() string {
	return s.String()
}

func (s *QuerySearchTagListRequest) SetAccessToken(v string) *QuerySearchTagListRequest {
	s.AccessToken = &v
	return s
}

func (s *QuerySearchTagListRequest) SetHeader(v map[string]*string) *QuerySearchTagListRequest {
	s.Header = v
	return s
}

type QuerySearchTagListResponse struct {
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QuerySearchTagListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QuerySearchTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySearchTagListResponse) GoString() string {
	return s.String()
}

func (s *QuerySearchTagListResponse) SetErrNo(v int32) *QuerySearchTagListResponse {
	s.ErrNo = &v
	return s
}

func (s *QuerySearchTagListResponse) SetErrMsg(v string) *QuerySearchTagListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QuerySearchTagListResponse) SetLogId(v string) *QuerySearchTagListResponse {
	s.LogId = &v
	return s
}

func (s *QuerySearchTagListResponse) SetData(v *QuerySearchTagListResponseData) *QuerySearchTagListResponse {
	s.Data = v
	return s
}

type QuerySearchTagListResponseData struct {
	MonthAvailableTimes *int64                                             `json:"month_available_times,omitempty" xml:"month_available_times,omitempty" require:"true"`
	MonthTotalTimes     *int64                                             `json:"month_total_times,omitempty" xml:"month_total_times,omitempty" require:"true"`
	SearchTagList       []*QuerySearchTagListResponseDataSearchTagListItem `json:"search_tag_list,omitempty" xml:"search_tag_list,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySearchTagListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuerySearchTagListResponseData) GoString() string {
	return s.String()
}

func (s *QuerySearchTagListResponseData) SetMonthAvailableTimes(v int64) *QuerySearchTagListResponseData {
	s.MonthAvailableTimes = &v
	return s
}

func (s *QuerySearchTagListResponseData) SetMonthTotalTimes(v int64) *QuerySearchTagListResponseData {
	s.MonthTotalTimes = &v
	return s
}

func (s *QuerySearchTagListResponseData) SetSearchTagList(v []*QuerySearchTagListResponseDataSearchTagListItem) *QuerySearchTagListResponseData {
	s.SearchTagList = v
	return s
}

type QuerySearchTagListResponseDataSearchTagListItem struct {
	SearchTag   *string `json:"search_tag,omitempty" xml:"search_tag,omitempty" require:"true"`
	AuditStatus *int32  `json:"audit_status,omitempty" xml:"audit_status,omitempty" require:"true"`
	AuditReason *string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty" require:"true"`
}

func (s QuerySearchTagListResponseDataSearchTagListItem) String() string {
	return tea.Prettify(s)
}

func (s QuerySearchTagListResponseDataSearchTagListItem) GoString() string {
	return s.String()
}

func (s *QuerySearchTagListResponseDataSearchTagListItem) SetSearchTag(v string) *QuerySearchTagListResponseDataSearchTagListItem {
	s.SearchTag = &v
	return s
}

func (s *QuerySearchTagListResponseDataSearchTagListItem) SetAuditStatus(v int32) *QuerySearchTagListResponseDataSearchTagListItem {
	s.AuditStatus = &v
	return s
}

func (s *QuerySearchTagListResponseDataSearchTagListItem) SetAuditReason(v string) *QuerySearchTagListResponseDataSearchTagListItem {
	s.AuditReason = &v
	return s
}

type QueryShortLiveDataWithIdRequest struct {
	RoomIdList       []*string          `json:"room_id_list,omitempty" xml:"room_id_list,omitempty" type:"Repeated"`
	EndTime          *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	PageSize         *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AwemeShortIdList []*string          `json:"aweme_short_id_list,omitempty" xml:"aweme_short_id_list,omitempty" type:"Repeated"`
	StartTime        *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	ItemIdList       []*string          `json:"item_id_list,omitempty" xml:"item_id_list,omitempty" type:"Repeated"`
	QueryBindType    *int32             `json:"query_bind_type,omitempty" xml:"query_bind_type,omitempty" require:"true"`
	PageNo           *int64             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	QueryDataType    *int32             `json:"query_data_type,omitempty" xml:"query_data_type,omitempty" require:"true"`
	HostName         *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	OpenItemIdList   []*string          `json:"open_item_id_list,omitempty" xml:"open_item_id_list,omitempty" type:"Repeated"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryShortLiveDataWithIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveDataWithIdRequest) GoString() string {
	return s.String()
}

func (s *QueryShortLiveDataWithIdRequest) SetRoomIdList(v []*string) *QueryShortLiveDataWithIdRequest {
	s.RoomIdList = v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetEndTime(v int64) *QueryShortLiveDataWithIdRequest {
	s.EndTime = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetPageSize(v int64) *QueryShortLiveDataWithIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetAccessToken(v string) *QueryShortLiveDataWithIdRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetAwemeShortIdList(v []*string) *QueryShortLiveDataWithIdRequest {
	s.AwemeShortIdList = v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetStartTime(v int64) *QueryShortLiveDataWithIdRequest {
	s.StartTime = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetItemIdList(v []*string) *QueryShortLiveDataWithIdRequest {
	s.ItemIdList = v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetQueryBindType(v int32) *QueryShortLiveDataWithIdRequest {
	s.QueryBindType = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetPageNo(v int64) *QueryShortLiveDataWithIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetQueryDataType(v int32) *QueryShortLiveDataWithIdRequest {
	s.QueryDataType = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetHostName(v string) *QueryShortLiveDataWithIdRequest {
	s.HostName = &v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetOpenItemIdList(v []*string) *QueryShortLiveDataWithIdRequest {
	s.OpenItemIdList = v
	return s
}

func (s *QueryShortLiveDataWithIdRequest) SetHeader(v map[string]*string) *QueryShortLiveDataWithIdRequest {
	s.Header = v
	return s
}

type QueryShortLiveDataWithIdResponse struct {
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryShortLiveDataWithIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryShortLiveDataWithIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveDataWithIdResponse) GoString() string {
	return s.String()
}

func (s *QueryShortLiveDataWithIdResponse) SetLogId(v string) *QueryShortLiveDataWithIdResponse {
	s.LogId = &v
	return s
}

func (s *QueryShortLiveDataWithIdResponse) SetData(v *QueryShortLiveDataWithIdResponseData) *QueryShortLiveDataWithIdResponse {
	s.Data = v
	return s
}

func (s *QueryShortLiveDataWithIdResponse) SetErrNo(v int32) *QueryShortLiveDataWithIdResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryShortLiveDataWithIdResponse) SetErrMsg(v string) *QueryShortLiveDataWithIdResponse {
	s.ErrMsg = &v
	return s
}

type QueryShortLiveDataWithIdResponseData struct {
	DetailData map[int64]map[string]*string `json:"detail_data,omitempty" xml:"detail_data,omitempty" require:"true"`
	DataList   []map[string]*string         `json:"data_list,omitempty" xml:"data_list,omitempty" require:"true" type:"Repeated"`
	Sum        *int64                       `json:"sum,omitempty" xml:"sum,omitempty" require:"true"`
}

func (s QueryShortLiveDataWithIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveDataWithIdResponseData) GoString() string {
	return s.String()
}

func (s *QueryShortLiveDataWithIdResponseData) SetDetailData(v map[int64]map[string]*string) *QueryShortLiveDataWithIdResponseData {
	s.DetailData = v
	return s
}

func (s *QueryShortLiveDataWithIdResponseData) SetDataList(v []map[string]*string) *QueryShortLiveDataWithIdResponseData {
	s.DataList = v
	return s
}

func (s *QueryShortLiveDataWithIdResponseData) SetSum(v int64) *QueryShortLiveDataWithIdResponseData {
	s.Sum = &v
	return s
}

type QueryShortLiveIdWithAwemeidRequest struct {
	AwemeShortIdList []*string          `json:"aweme_short_id_list,omitempty" xml:"aweme_short_id_list,omitempty" require:"true" type:"Repeated"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	QueryDataType    *int32             `json:"query_data_type,omitempty" xml:"query_data_type,omitempty"`
	StartTime        *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName         *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	QueryBindType    *int32             `json:"query_bind_type,omitempty" xml:"query_bind_type,omitempty" require:"true"`
}

func (s QueryShortLiveIdWithAwemeidRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveIdWithAwemeidRequest) GoString() string {
	return s.String()
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetAwemeShortIdList(v []*string) *QueryShortLiveIdWithAwemeidRequest {
	s.AwemeShortIdList = v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetHeader(v map[string]*string) *QueryShortLiveIdWithAwemeidRequest {
	s.Header = v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetAccessToken(v string) *QueryShortLiveIdWithAwemeidRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetQueryDataType(v int32) *QueryShortLiveIdWithAwemeidRequest {
	s.QueryDataType = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetStartTime(v int64) *QueryShortLiveIdWithAwemeidRequest {
	s.StartTime = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetEndTime(v int64) *QueryShortLiveIdWithAwemeidRequest {
	s.EndTime = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetHostName(v string) *QueryShortLiveIdWithAwemeidRequest {
	s.HostName = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidRequest) SetQueryBindType(v int32) *QueryShortLiveIdWithAwemeidRequest {
	s.QueryBindType = &v
	return s
}

type QueryShortLiveIdWithAwemeidResponse struct {
	LogId  *string                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryShortLiveIdWithAwemeidResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryShortLiveIdWithAwemeidResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveIdWithAwemeidResponse) GoString() string {
	return s.String()
}

func (s *QueryShortLiveIdWithAwemeidResponse) SetLogId(v string) *QueryShortLiveIdWithAwemeidResponse {
	s.LogId = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidResponse) SetData(v *QueryShortLiveIdWithAwemeidResponseData) *QueryShortLiveIdWithAwemeidResponse {
	s.Data = v
	return s
}

func (s *QueryShortLiveIdWithAwemeidResponse) SetErrNo(v int32) *QueryShortLiveIdWithAwemeidResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryShortLiveIdWithAwemeidResponse) SetErrMsg(v string) *QueryShortLiveIdWithAwemeidResponse {
	s.ErrMsg = &v
	return s
}

type QueryShortLiveIdWithAwemeidResponseData struct {
	OverViewData map[string]*string   `json:"over_view_data,omitempty" xml:"over_view_data,omitempty" require:"true"`
	ItemData     map[string][]*string `json:"item_data,omitempty" xml:"item_data,omitempty" require:"true"`
	LiveData     map[string][]*string `json:"live_data,omitempty" xml:"live_data,omitempty" require:"true"`
}

func (s QueryShortLiveIdWithAwemeidResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryShortLiveIdWithAwemeidResponseData) GoString() string {
	return s.String()
}

func (s *QueryShortLiveIdWithAwemeidResponseData) SetOverViewData(v map[string]*string) *QueryShortLiveIdWithAwemeidResponseData {
	s.OverViewData = v
	return s
}

func (s *QueryShortLiveIdWithAwemeidResponseData) SetItemData(v map[string][]*string) *QueryShortLiveIdWithAwemeidResponseData {
	s.ItemData = v
	return s
}

func (s *QueryShortLiveIdWithAwemeidResponseData) SetLiveData(v map[string][]*string) *QueryShortLiveIdWithAwemeidResponseData {
	s.LiveData = v
	return s
}

type QuerySimpleQrBindListRequest struct {
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QuerySimpleQrBindListRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySimpleQrBindListRequest) GoString() string {
	return s.String()
}

func (s *QuerySimpleQrBindListRequest) SetPageNum(v int64) *QuerySimpleQrBindListRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySimpleQrBindListRequest) SetPageSize(v int64) *QuerySimpleQrBindListRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySimpleQrBindListRequest) SetHeader(v map[string]*string) *QuerySimpleQrBindListRequest {
	s.Header = v
	return s
}

func (s *QuerySimpleQrBindListRequest) SetAccessToken(v string) *QuerySimpleQrBindListRequest {
	s.AccessToken = &v
	return s
}

type QuerySimpleQrBindListResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QuerySimpleQrBindListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QuerySimpleQrBindListResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySimpleQrBindListResponse) GoString() string {
	return s.String()
}

func (s *QuerySimpleQrBindListResponse) SetErrMsg(v string) *QuerySimpleQrBindListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QuerySimpleQrBindListResponse) SetLogId(v string) *QuerySimpleQrBindListResponse {
	s.LogId = &v
	return s
}

func (s *QuerySimpleQrBindListResponse) SetData(v *QuerySimpleQrBindListResponseData) *QuerySimpleQrBindListResponse {
	s.Data = v
	return s
}

func (s *QuerySimpleQrBindListResponse) SetErrNo(v int32) *QuerySimpleQrBindListResponse {
	s.ErrNo = &v
	return s
}

type QuerySimpleQrBindListResponseData struct {
	TotalCount        *int32                                         `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	BindUpperLimit    *int32                                         `json:"bind_upper_limit,omitempty" xml:"bind_upper_limit,omitempty" require:"true"`
	BindRestCount     *int32                                         `json:"bind_rest_count,omitempty" xml:"bind_rest_count,omitempty" require:"true"`
	ReleaseUpperLimit *int32                                         `json:"release_upper_limit,omitempty" xml:"release_upper_limit,omitempty" require:"true"`
	ReleaseRestCount  *int32                                         `json:"release_rest_count,omitempty" xml:"release_rest_count,omitempty" require:"true"`
	QrList            []*QuerySimpleQrBindListResponseDataQrListItem `json:"qr_list,omitempty" xml:"qr_list,omitempty" type:"Repeated"`
}

func (s QuerySimpleQrBindListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuerySimpleQrBindListResponseData) GoString() string {
	return s.String()
}

func (s *QuerySimpleQrBindListResponseData) SetTotalCount(v int32) *QuerySimpleQrBindListResponseData {
	s.TotalCount = &v
	return s
}

func (s *QuerySimpleQrBindListResponseData) SetBindUpperLimit(v int32) *QuerySimpleQrBindListResponseData {
	s.BindUpperLimit = &v
	return s
}

func (s *QuerySimpleQrBindListResponseData) SetBindRestCount(v int32) *QuerySimpleQrBindListResponseData {
	s.BindRestCount = &v
	return s
}

func (s *QuerySimpleQrBindListResponseData) SetReleaseUpperLimit(v int32) *QuerySimpleQrBindListResponseData {
	s.ReleaseUpperLimit = &v
	return s
}

func (s *QuerySimpleQrBindListResponseData) SetReleaseRestCount(v int32) *QuerySimpleQrBindListResponseData {
	s.ReleaseRestCount = &v
	return s
}

func (s *QuerySimpleQrBindListResponseData) SetQrList(v []*QuerySimpleQrBindListResponseDataQrListItem) *QuerySimpleQrBindListResponseData {
	s.QrList = v
	return s
}

type QuerySimpleQrBindListResponseDataQrListItem struct {
	Status               *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	ExclusiveQrUrlPrefix *int32  `json:"exclusive_qr_url_prefix,omitempty" xml:"exclusive_qr_url_prefix,omitempty" require:"true"`
	QrUrl                *string `json:"qr_url,omitempty" xml:"qr_url,omitempty" require:"true"`
	LoadPath             *string `json:"load_path,omitempty" xml:"load_path,omitempty" require:"true"`
	Stage                *string `json:"stage,omitempty" xml:"stage,omitempty" require:"true"`
}

func (s QuerySimpleQrBindListResponseDataQrListItem) String() string {
	return tea.Prettify(s)
}

func (s QuerySimpleQrBindListResponseDataQrListItem) GoString() string {
	return s.String()
}

func (s *QuerySimpleQrBindListResponseDataQrListItem) SetStatus(v int32) *QuerySimpleQrBindListResponseDataQrListItem {
	s.Status = &v
	return s
}

func (s *QuerySimpleQrBindListResponseDataQrListItem) SetExclusiveQrUrlPrefix(v int32) *QuerySimpleQrBindListResponseDataQrListItem {
	s.ExclusiveQrUrlPrefix = &v
	return s
}

func (s *QuerySimpleQrBindListResponseDataQrListItem) SetQrUrl(v string) *QuerySimpleQrBindListResponseDataQrListItem {
	s.QrUrl = &v
	return s
}

func (s *QuerySimpleQrBindListResponseDataQrListItem) SetLoadPath(v string) *QuerySimpleQrBindListResponseDataQrListItem {
	s.LoadPath = &v
	return s
}

func (s *QuerySimpleQrBindListResponseDataQrListItem) SetStage(v string) *QuerySimpleQrBindListResponseDataQrListItem {
	s.Stage = &v
	return s
}

type QuerySmallHomeOrderDataRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
}

func (s QuerySmallHomeOrderDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOrderDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOrderDataRequest) SetHeader(v map[string]*string) *QuerySmallHomeOrderDataRequest {
	s.Header = v
	return s
}

func (s *QuerySmallHomeOrderDataRequest) SetAccessToken(v string) *QuerySmallHomeOrderDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QuerySmallHomeOrderDataRequest) SetPageSize(v int64) *QuerySmallHomeOrderDataRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySmallHomeOrderDataRequest) SetStartTime(v int64) *QuerySmallHomeOrderDataRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySmallHomeOrderDataRequest) SetEndTime(v int64) *QuerySmallHomeOrderDataRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySmallHomeOrderDataRequest) SetPageNum(v int64) *QuerySmallHomeOrderDataRequest {
	s.PageNum = &v
	return s
}

type QuerySmallHomeOrderDataResponse struct {
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QuerySmallHomeOrderDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QuerySmallHomeOrderDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOrderDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOrderDataResponse) SetErrNo(v int32) *QuerySmallHomeOrderDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponse) SetErrMsg(v string) *QuerySmallHomeOrderDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponse) SetLogId(v string) *QuerySmallHomeOrderDataResponse {
	s.LogId = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponse) SetData(v *QuerySmallHomeOrderDataResponseData) *QuerySmallHomeOrderDataResponse {
	s.Data = v
	return s
}

type QuerySmallHomeOrderDataResponseData struct {
	OrderData map[string]map[string]*string `json:"order_data,omitempty" xml:"order_data,omitempty"`
	TotalNum  *int64                        `json:"total_num,omitempty" xml:"total_num,omitempty"`
	HasMore   *bool                         `json:"has_more,omitempty" xml:"has_more,omitempty"`
	AppId     *string                       `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	AppName   *string                       `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
}

func (s QuerySmallHomeOrderDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOrderDataResponseData) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOrderDataResponseData) SetOrderData(v map[string]map[string]*string) *QuerySmallHomeOrderDataResponseData {
	s.OrderData = v
	return s
}

func (s *QuerySmallHomeOrderDataResponseData) SetTotalNum(v int64) *QuerySmallHomeOrderDataResponseData {
	s.TotalNum = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponseData) SetHasMore(v bool) *QuerySmallHomeOrderDataResponseData {
	s.HasMore = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponseData) SetAppId(v string) *QuerySmallHomeOrderDataResponseData {
	s.AppId = &v
	return s
}

func (s *QuerySmallHomeOrderDataResponseData) SetAppName(v string) *QuerySmallHomeOrderDataResponseData {
	s.AppName = &v
	return s
}

type QuerySmallHomeOverviewDataRequest struct {
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QuerySmallHomeOverviewDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOverviewDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOverviewDataRequest) SetStartTime(v int64) *QuerySmallHomeOverviewDataRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySmallHomeOverviewDataRequest) SetEndTime(v int64) *QuerySmallHomeOverviewDataRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySmallHomeOverviewDataRequest) SetHeader(v map[string]*string) *QuerySmallHomeOverviewDataRequest {
	s.Header = v
	return s
}

func (s *QuerySmallHomeOverviewDataRequest) SetAccessToken(v string) *QuerySmallHomeOverviewDataRequest {
	s.AccessToken = &v
	return s
}

type QuerySmallHomeOverviewDataResponse struct {
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QuerySmallHomeOverviewDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QuerySmallHomeOverviewDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOverviewDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOverviewDataResponse) SetLogId(v string) *QuerySmallHomeOverviewDataResponse {
	s.LogId = &v
	return s
}

func (s *QuerySmallHomeOverviewDataResponse) SetData(v *QuerySmallHomeOverviewDataResponseData) *QuerySmallHomeOverviewDataResponse {
	s.Data = v
	return s
}

func (s *QuerySmallHomeOverviewDataResponse) SetErrNo(v int32) *QuerySmallHomeOverviewDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QuerySmallHomeOverviewDataResponse) SetErrMsg(v string) *QuerySmallHomeOverviewDataResponse {
	s.ErrMsg = &v
	return s
}

type QuerySmallHomeOverviewDataResponseData struct {
	AppName      *string            `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	OverViewData map[string]*string `json:"over_view_data,omitempty" xml:"over_view_data,omitempty"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s QuerySmallHomeOverviewDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeOverviewDataResponseData) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeOverviewDataResponseData) SetAppName(v string) *QuerySmallHomeOverviewDataResponseData {
	s.AppName = &v
	return s
}

func (s *QuerySmallHomeOverviewDataResponseData) SetOverViewData(v map[string]*string) *QuerySmallHomeOverviewDataResponseData {
	s.OverViewData = v
	return s
}

func (s *QuerySmallHomeOverviewDataResponseData) SetAppId(v string) *QuerySmallHomeOverviewDataResponseData {
	s.AppId = &v
	return s
}

type QuerySmallHomeRoomDataRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s QuerySmallHomeRoomDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeRoomDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeRoomDataRequest) SetHeader(v map[string]*string) *QuerySmallHomeRoomDataRequest {
	s.Header = v
	return s
}

func (s *QuerySmallHomeRoomDataRequest) SetAccessToken(v string) *QuerySmallHomeRoomDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QuerySmallHomeRoomDataRequest) SetStartTime(v int64) *QuerySmallHomeRoomDataRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySmallHomeRoomDataRequest) SetEndTime(v int64) *QuerySmallHomeRoomDataRequest {
	s.EndTime = &v
	return s
}

type QuerySmallHomeRoomDataResponse struct {
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QuerySmallHomeRoomDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QuerySmallHomeRoomDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeRoomDataResponse) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeRoomDataResponse) SetErrMsg(v string) *QuerySmallHomeRoomDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QuerySmallHomeRoomDataResponse) SetLogId(v string) *QuerySmallHomeRoomDataResponse {
	s.LogId = &v
	return s
}

func (s *QuerySmallHomeRoomDataResponse) SetData(v *QuerySmallHomeRoomDataResponseData) *QuerySmallHomeRoomDataResponse {
	s.Data = v
	return s
}

func (s *QuerySmallHomeRoomDataResponse) SetErrNo(v int32) *QuerySmallHomeRoomDataResponse {
	s.ErrNo = &v
	return s
}

type QuerySmallHomeRoomDataResponseData struct {
	RoomData map[string]map[string]*string `json:"room_data,omitempty" xml:"room_data,omitempty"`
	TotalNum *int64                        `json:"total_num,omitempty" xml:"total_num,omitempty"`
	HasMore  *bool                         `json:"has_more,omitempty" xml:"has_more,omitempty"`
	AppId    *string                       `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	AppName  *string                       `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
}

func (s QuerySmallHomeRoomDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuerySmallHomeRoomDataResponseData) GoString() string {
	return s.String()
}

func (s *QuerySmallHomeRoomDataResponseData) SetRoomData(v map[string]map[string]*string) *QuerySmallHomeRoomDataResponseData {
	s.RoomData = v
	return s
}

func (s *QuerySmallHomeRoomDataResponseData) SetTotalNum(v int64) *QuerySmallHomeRoomDataResponseData {
	s.TotalNum = &v
	return s
}

func (s *QuerySmallHomeRoomDataResponseData) SetHasMore(v bool) *QuerySmallHomeRoomDataResponseData {
	s.HasMore = &v
	return s
}

func (s *QuerySmallHomeRoomDataResponseData) SetAppId(v string) *QuerySmallHomeRoomDataResponseData {
	s.AppId = &v
	return s
}

func (s *QuerySmallHomeRoomDataResponseData) SetAppName(v string) *QuerySmallHomeRoomDataResponseData {
	s.AppName = &v
	return s
}

type QueryTaskVideoDailyDataRequest struct {
	PageNum               *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	VideoPublishEndTime   *int64             `json:"video_publish_end_time,omitempty" xml:"video_publish_end_time,omitempty" require:"true"`
	VideoPublishStartTime *int64             `json:"video_publish_start_time,omitempty" xml:"video_publish_start_time,omitempty" require:"true"`
	DouyinId              *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty"`
	PageSize              *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	BillingDate           *string            `json:"billing_date,omitempty" xml:"billing_date,omitempty" require:"true"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AgencyClientKey       *string            `json:"agency_client_key,omitempty" xml:"agency_client_key,omitempty"`
}

func (s QueryTaskVideoDailyDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDailyDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDailyDataRequest) SetPageNum(v int32) *QueryTaskVideoDailyDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetVideoPublishEndTime(v int64) *QueryTaskVideoDailyDataRequest {
	s.VideoPublishEndTime = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetVideoPublishStartTime(v int64) *QueryTaskVideoDailyDataRequest {
	s.VideoPublishStartTime = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetDouyinId(v string) *QueryTaskVideoDailyDataRequest {
	s.DouyinId = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetPageSize(v int32) *QueryTaskVideoDailyDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetBillingDate(v string) *QueryTaskVideoDailyDataRequest {
	s.BillingDate = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetHeader(v map[string]*string) *QueryTaskVideoDailyDataRequest {
	s.Header = v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetAccessToken(v string) *QueryTaskVideoDailyDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryTaskVideoDailyDataRequest) SetAgencyClientKey(v string) *QueryTaskVideoDailyDataRequest {
	s.AgencyClientKey = &v
	return s
}

type QueryTaskVideoDailyDataResponse struct {
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryTaskVideoDailyDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryTaskVideoDailyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDailyDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDailyDataResponse) SetLogId(v string) *QueryTaskVideoDailyDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponse) SetData(v *QueryTaskVideoDailyDataResponseData) *QueryTaskVideoDailyDataResponse {
	s.Data = v
	return s
}

func (s *QueryTaskVideoDailyDataResponse) SetErrNo(v int32) *QueryTaskVideoDailyDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponse) SetErrMsg(v string) *QueryTaskVideoDailyDataResponse {
	s.ErrMsg = &v
	return s
}

type QueryTaskVideoDailyDataResponseData struct {
	VideoDailyDataList []*QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem `json:"video_daily_data_list,omitempty" xml:"video_daily_data_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTaskVideoDailyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDailyDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDailyDataResponseData) SetVideoDailyDataList(v []*QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) *QueryTaskVideoDailyDataResponseData {
	s.VideoDailyDataList = v
	return s
}

type QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem struct {
	Author             *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	ExpectedProfit     *int64  `json:"expected_profit,omitempty" xml:"expected_profit,omitempty"`
	TaskId             *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	RefundGmv1d        *int64  `json:"refund_gmv_1d,omitempty" xml:"refund_gmv_1d,omitempty"`
	Date               *string `json:"date,omitempty" xml:"date,omitempty"`
	BillingGmv1d       *int64  `json:"billing_gmv_1d,omitempty" xml:"billing_gmv_1d,omitempty"`
	AdShareCost1d      *int64  `json:"ad_share_cost_1d,omitempty" xml:"ad_share_cost_1d,omitempty"`
	MicroAppTitle      *string `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty" require:"true"`
	VideoTitle         *string `json:"video_title,omitempty" xml:"video_title,omitempty" require:"true"`
	VideoLink          *string `json:"video_link,omitempty" xml:"video_link,omitempty" require:"true"`
	Gmv1d              *int64  `json:"gmv_1d,omitempty" xml:"gmv_1d,omitempty"`
	FeedAdShareCost1d  *int64  `json:"feed_ad_share_cost_1d,omitempty" xml:"feed_ad_share_cost_1d,omitempty"`
	PublishTime        *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
	ActiveCnt1d        *int64  `json:"active_cnt_1d,omitempty" xml:"active_cnt_1d,omitempty"`
	ClientName         *string `json:"client_name,omitempty" xml:"client_name,omitempty"`
	DouyinId           *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	TaskName           *string `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	VideoId            *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	BillingRefundGmv1d *int64  `json:"billing_refund_gmv_1d,omitempty" xml:"billing_refund_gmv_1d,omitempty"`
}

func (s QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetAuthor(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.Author = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetExpectedProfit(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.ExpectedProfit = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetTaskId(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.TaskId = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetRefundGmv1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.RefundGmv1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetDate(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.Date = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetBillingGmv1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.BillingGmv1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetAdShareCost1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.AdShareCost1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetMicroAppTitle(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.MicroAppTitle = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetVideoTitle(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.VideoTitle = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetVideoLink(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.VideoLink = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetGmv1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.Gmv1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetFeedAdShareCost1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.FeedAdShareCost1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetPublishTime(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.PublishTime = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetActiveCnt1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.ActiveCnt1d = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetClientName(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.ClientName = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetDouyinId(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.DouyinId = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetTaskName(v string) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.TaskName = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetVideoId(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.VideoId = &v
	return s
}

func (s *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem) SetBillingRefundGmv1d(v int64) *QueryTaskVideoDailyDataResponseDataVideoDailyDataListItem {
	s.BillingRefundGmv1d = &v
	return s
}

type QueryTaskVideoDataRequest struct {
	PageSize              *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	PageNum               *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskIds               []*int64           `json:"task_ids,omitempty" xml:"task_ids,omitempty" require:"true" type:"Repeated"`
	VideoPublishStartTime *int64             `json:"video_publish_start_time,omitempty" xml:"video_publish_start_time,omitempty" require:"true"`
	VideoPublishEndTime   *int64             `json:"video_publish_end_time,omitempty" xml:"video_publish_end_time,omitempty" require:"true"`
}

func (s QueryTaskVideoDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDataRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDataRequest) SetPageSize(v int32) *QueryTaskVideoDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskVideoDataRequest) SetPageNum(v int32) *QueryTaskVideoDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskVideoDataRequest) SetHeader(v map[string]*string) *QueryTaskVideoDataRequest {
	s.Header = v
	return s
}

func (s *QueryTaskVideoDataRequest) SetAccessToken(v string) *QueryTaskVideoDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryTaskVideoDataRequest) SetTaskIds(v []*int64) *QueryTaskVideoDataRequest {
	s.TaskIds = v
	return s
}

func (s *QueryTaskVideoDataRequest) SetVideoPublishStartTime(v int64) *QueryTaskVideoDataRequest {
	s.VideoPublishStartTime = &v
	return s
}

func (s *QueryTaskVideoDataRequest) SetVideoPublishEndTime(v int64) *QueryTaskVideoDataRequest {
	s.VideoPublishEndTime = &v
	return s
}

type QueryTaskVideoDataResponse struct {
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryTaskVideoDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryTaskVideoDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDataResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDataResponse) SetErrMsg(v string) *QueryTaskVideoDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryTaskVideoDataResponse) SetLogId(v string) *QueryTaskVideoDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryTaskVideoDataResponse) SetData(v *QueryTaskVideoDataResponseData) *QueryTaskVideoDataResponse {
	s.Data = v
	return s
}

func (s *QueryTaskVideoDataResponse) SetErrNo(v int32) *QueryTaskVideoDataResponse {
	s.ErrNo = &v
	return s
}

type QueryTaskVideoDataResponseData struct {
	VideoSumDataList []*QueryTaskVideoDataResponseDataVideoSumDataListItem `json:"video_sum_data_list,omitempty" xml:"video_sum_data_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTaskVideoDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDataResponseData) SetVideoSumDataList(v []*QueryTaskVideoDataResponseDataVideoSumDataListItem) *QueryTaskVideoDataResponseData {
	s.VideoSumDataList = v
	return s
}

type QueryTaskVideoDataResponseDataVideoSumDataListItem struct {
	Author             *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	MicroAppTitle      *string `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty" require:"true"`
	ClientName         *string `json:"client_name,omitempty" xml:"client_name,omitempty"`
	Likes              *int64  `json:"likes,omitempty" xml:"likes,omitempty" require:"true"`
	Date               *string `json:"date,omitempty" xml:"date,omitempty"`
	RefundGmvTd        *int64  `json:"refund_gmv_td,omitempty" xml:"refund_gmv_td,omitempty"`
	GmvTd              *int64  `json:"gmv_td,omitempty" xml:"gmv_td,omitempty"`
	TaskName           *string `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	VideoTitle         *string `json:"video_title,omitempty" xml:"video_title,omitempty" require:"true"`
	BillingRefundGmvTd *int64  `json:"billing_refund_gmv_td,omitempty" xml:"billing_refund_gmv_td,omitempty"`
	BillingGmvTd       *int64  `json:"billing_gmv_td,omitempty" xml:"billing_gmv_td,omitempty"`
	ActiveCntTd        *int64  `json:"active_cnt_td,omitempty" xml:"active_cnt_td,omitempty"`
	Shares             *int64  `json:"shares,omitempty" xml:"shares,omitempty" require:"true"`
	AdShareCostTd      *int64  `json:"ad_share_cost_td,omitempty" xml:"ad_share_cost_td,omitempty"`
	FeedAdShareCostTd  *int64  `json:"feed_ad_share_cost_td,omitempty" xml:"feed_ad_share_cost_td,omitempty"`
	VideoViews         *int64  `json:"video_views,omitempty" xml:"video_views,omitempty" require:"true"`
	DouyinId           *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	VideoLink          *string `json:"video_link,omitempty" xml:"video_link,omitempty" require:"true"`
	Comments           *int64  `json:"comments,omitempty" xml:"comments,omitempty" require:"true"`
	VideoId            *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	TaskId             *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	PublishTime        *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
	Clicks             *int64  `json:"clicks,omitempty" xml:"clicks,omitempty"`
	ExpectedProfit     *int64  `json:"expected_profit,omitempty" xml:"expected_profit,omitempty"`
}

func (s QueryTaskVideoDataResponseDataVideoSumDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoDataResponseDataVideoSumDataListItem) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetAuthor(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Author = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetMicroAppTitle(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.MicroAppTitle = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetClientName(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.ClientName = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetLikes(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Likes = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetDate(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Date = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetRefundGmvTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.RefundGmvTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetGmvTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.GmvTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetTaskName(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.TaskName = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetVideoTitle(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.VideoTitle = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetBillingRefundGmvTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.BillingRefundGmvTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetBillingGmvTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.BillingGmvTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetActiveCntTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.ActiveCntTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetShares(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Shares = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetAdShareCostTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.AdShareCostTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetFeedAdShareCostTd(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.FeedAdShareCostTd = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetVideoViews(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.VideoViews = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetDouyinId(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.DouyinId = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetVideoLink(v string) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.VideoLink = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetComments(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Comments = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetVideoId(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.VideoId = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetTaskId(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.TaskId = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetPublishTime(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.PublishTime = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetClicks(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.Clicks = &v
	return s
}

func (s *QueryTaskVideoDataResponseDataVideoSumDataListItem) SetExpectedProfit(v int64) *QueryTaskVideoDataResponseDataVideoSumDataListItem {
	s.ExpectedProfit = &v
	return s
}

type QueryTaskVideoStatusRequest struct {
	PageNum               *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	AppId                 *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	VideoPublishEndTime   *int64             `json:"video_publish_end_time,omitempty" xml:"video_publish_end_time,omitempty" require:"true"`
	PageSize              *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	DouyinId              *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty"`
	AgentId               *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	VideoPublishStartTime *int64             `json:"video_publish_start_time,omitempty" xml:"video_publish_start_time,omitempty" require:"true"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryTaskVideoStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoStatusRequest) SetPageNum(v int32) *QueryTaskVideoStatusRequest {
	s.PageNum = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetAppId(v string) *QueryTaskVideoStatusRequest {
	s.AppId = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetHeader(v map[string]*string) *QueryTaskVideoStatusRequest {
	s.Header = v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetVideoPublishEndTime(v int64) *QueryTaskVideoStatusRequest {
	s.VideoPublishEndTime = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetPageSize(v int32) *QueryTaskVideoStatusRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetDouyinId(v string) *QueryTaskVideoStatusRequest {
	s.DouyinId = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetAgentId(v int64) *QueryTaskVideoStatusRequest {
	s.AgentId = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetVideoPublishStartTime(v int64) *QueryTaskVideoStatusRequest {
	s.VideoPublishStartTime = &v
	return s
}

func (s *QueryTaskVideoStatusRequest) SetAccessToken(v string) *QueryTaskVideoStatusRequest {
	s.AccessToken = &v
	return s
}

type QueryTaskVideoStatusResponse struct {
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryTaskVideoStatusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryTaskVideoStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoStatusResponse) SetErrMsg(v string) *QueryTaskVideoStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryTaskVideoStatusResponse) SetLogId(v string) *QueryTaskVideoStatusResponse {
	s.LogId = &v
	return s
}

func (s *QueryTaskVideoStatusResponse) SetData(v *QueryTaskVideoStatusResponseData) *QueryTaskVideoStatusResponse {
	s.Data = v
	return s
}

func (s *QueryTaskVideoStatusResponse) SetErrNo(v int32) *QueryTaskVideoStatusResponse {
	s.ErrNo = &v
	return s
}

type QueryTaskVideoStatusResponseData struct {
	Results []*QueryTaskVideoStatusResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTaskVideoStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoStatusResponseData) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoStatusResponseData) SetResults(v []*QueryTaskVideoStatusResponseDataResultsItem) *QueryTaskVideoStatusResponseData {
	s.Results = v
	return s
}

type QueryTaskVideoStatusResponseDataResultsItem struct {
	UnbindReason *string `json:"unbind_reason,omitempty" xml:"unbind_reason,omitempty"`
	VideoId      *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	VideoStatus  *int    `json:"video_status,omitempty" xml:"video_status,omitempty" require:"true"`
}

func (s QueryTaskVideoStatusResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryTaskVideoStatusResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *QueryTaskVideoStatusResponseDataResultsItem) SetUnbindReason(v string) *QueryTaskVideoStatusResponseDataResultsItem {
	s.UnbindReason = &v
	return s
}

func (s *QueryTaskVideoStatusResponseDataResultsItem) SetVideoId(v int64) *QueryTaskVideoStatusResponseDataResultsItem {
	s.VideoId = &v
	return s
}

func (s *QueryTaskVideoStatusResponseDataResultsItem) SetVideoStatus(v int) *QueryTaskVideoStatusResponseDataResultsItem {
	s.VideoStatus = &v
	return s
}

type QueryUserInteractTaskRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TaskId      []*string          `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryUserInteractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskRequest) SetAppId(v string) *QueryUserInteractTaskRequest {
	s.AppId = &v
	return s
}

func (s *QueryUserInteractTaskRequest) SetOpenId(v string) *QueryUserInteractTaskRequest {
	s.OpenId = &v
	return s
}

func (s *QueryUserInteractTaskRequest) SetTaskId(v []*string) *QueryUserInteractTaskRequest {
	s.TaskId = v
	return s
}

func (s *QueryUserInteractTaskRequest) SetHeader(v map[string]*string) *QueryUserInteractTaskRequest {
	s.Header = v
	return s
}

func (s *QueryUserInteractTaskRequest) SetAccessToken(v string) *QueryUserInteractTaskRequest {
	s.AccessToken = &v
	return s
}

type QueryUserInteractTaskResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryUserInteractTaskResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryUserInteractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskResponse) SetLogId(v string) *QueryUserInteractTaskResponse {
	s.LogId = &v
	return s
}

func (s *QueryUserInteractTaskResponse) SetData(v *QueryUserInteractTaskResponseData) *QueryUserInteractTaskResponse {
	s.Data = v
	return s
}

func (s *QueryUserInteractTaskResponse) SetErrNo(v int32) *QueryUserInteractTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryUserInteractTaskResponse) SetErrMsg(v string) *QueryUserInteractTaskResponse {
	s.ErrMsg = &v
	return s
}

type QueryUserInteractTaskResponseData struct {
	TaskInfoList map[string]*QueryUserInteractTaskResponseDataTaskInfoListValue `json:"task_info_list,omitempty" xml:"task_info_list,omitempty" require:"true"`
}

func (s QueryUserInteractTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskResponseData) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskResponseData) SetTaskInfoList(v map[string]*QueryUserInteractTaskResponseDataTaskInfoListValue) *QueryUserInteractTaskResponseData {
	s.TaskInfoList = v
	return s
}

type QueryUserInteractTaskResponseDataTaskInfoListValue struct {
	MaxCount     *int64                                                             `json:"max_count,omitempty" xml:"max_count,omitempty" require:"true"`
	ErrNo        *int32                                                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg       *string                                                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	TaskId       *string                                                            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	VideoInfo    []*QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem `json:"video_info,omitempty" xml:"video_info,omitempty" require:"true" type:"Repeated"`
	IsValid      *bool                                                              `json:"is_valid,omitempty" xml:"is_valid,omitempty" require:"true"`
	SuccessCount *int64                                                             `json:"success_count,omitempty" xml:"success_count,omitempty" require:"true"`
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValue) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValue) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetMaxCount(v int64) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.MaxCount = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetErrNo(v int32) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.ErrNo = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetErrMsg(v string) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.ErrMsg = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetTaskId(v string) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.TaskId = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetVideoInfo(v []*QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.VideoInfo = v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetIsValid(v bool) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.IsValid = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValue) SetSuccessCount(v int64) *QueryUserInteractTaskResponseDataTaskInfoListValue {
	s.SuccessCount = &v
	return s
}

type QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem struct {
	InteractInfo map[string]*QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue `json:"interact_info,omitempty" xml:"interact_info,omitempty" require:"true"`
	VideoId      *string                                                                                      `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	VideoStatus  *int32                                                                                       `json:"video_status,omitempty" xml:"video_status,omitempty" require:"true"`
	Completed    *bool                                                                                        `json:"completed,omitempty" xml:"completed,omitempty" require:"true"`
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) SetInteractInfo(v map[string]*QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.InteractInfo = v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) SetVideoId(v string) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.VideoId = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) SetVideoStatus(v int32) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.VideoStatus = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem) SetCompleted(v bool) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItem {
	s.Completed = &v
	return s
}

type QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue struct {
	Completed *bool  `json:"completed,omitempty" xml:"completed,omitempty" require:"true"`
	Stage     *int32 `json:"stage,omitempty" xml:"stage,omitempty" require:"true"`
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue) GoString() string {
	return s.String()
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue) SetCompleted(v bool) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue {
	s.Completed = &v
	return s
}

func (s *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue) SetStage(v int32) *QueryUserInteractTaskResponseDataTaskInfoListValueVideoInfoItemInteractInfoValue {
	s.Stage = &v
	return s
}

type QueryUserPortraitDataRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UserType    *string            `json:"user_type,omitempty" xml:"user_type,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty" require:"true"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryUserPortraitDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataRequest) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataRequest) SetAccessToken(v string) *QueryUserPortraitDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetUserType(v string) *QueryUserPortraitDataRequest {
	s.UserType = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetHostName(v string) *QueryUserPortraitDataRequest {
	s.HostName = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetVersionType(v string) *QueryUserPortraitDataRequest {
	s.VersionType = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetStartTime(v int64) *QueryUserPortraitDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetEndTime(v int64) *QueryUserPortraitDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUserPortraitDataRequest) SetHeader(v map[string]*string) *QueryUserPortraitDataRequest {
	s.Header = v
	return s
}

type QueryUserPortraitDataResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryUserPortraitDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryUserPortraitDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponse) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponse) SetErrMsg(v string) *QueryUserPortraitDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryUserPortraitDataResponse) SetLogId(v string) *QueryUserPortraitDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryUserPortraitDataResponse) SetData(v *QueryUserPortraitDataResponseData) *QueryUserPortraitDataResponse {
	s.Data = v
	return s
}

func (s *QueryUserPortraitDataResponse) SetErrNo(v int32) *QueryUserPortraitDataResponse {
	s.ErrNo = &v
	return s
}

type QueryUserPortraitDataResponseData struct {
	TotalUser *int64                                           `json:"total_user,omitempty" xml:"total_user,omitempty" require:"true"`
	Province  []*QueryUserPortraitDataResponseDataProvinceItem `json:"province,omitempty" xml:"province,omitempty" type:"Repeated"`
	City      []*QueryUserPortraitDataResponseDataCityItem     `json:"city,omitempty" xml:"city,omitempty" type:"Repeated"`
	Gender    []*QueryUserPortraitDataResponseDataGenderItem   `json:"gender,omitempty" xml:"gender,omitempty" type:"Repeated"`
	Age       []*QueryUserPortraitDataResponseDataAgeItem      `json:"age,omitempty" xml:"age,omitempty" type:"Repeated"`
}

func (s QueryUserPortraitDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponseData) SetTotalUser(v int64) *QueryUserPortraitDataResponseData {
	s.TotalUser = &v
	return s
}

func (s *QueryUserPortraitDataResponseData) SetProvince(v []*QueryUserPortraitDataResponseDataProvinceItem) *QueryUserPortraitDataResponseData {
	s.Province = v
	return s
}

func (s *QueryUserPortraitDataResponseData) SetCity(v []*QueryUserPortraitDataResponseDataCityItem) *QueryUserPortraitDataResponseData {
	s.City = v
	return s
}

func (s *QueryUserPortraitDataResponseData) SetGender(v []*QueryUserPortraitDataResponseDataGenderItem) *QueryUserPortraitDataResponseData {
	s.Gender = v
	return s
}

func (s *QueryUserPortraitDataResponseData) SetAge(v []*QueryUserPortraitDataResponseDataAgeItem) *QueryUserPortraitDataResponseData {
	s.Age = v
	return s
}

type QueryUserPortraitDataResponseDataAgeItem struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryUserPortraitDataResponseDataAgeItem) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponseDataAgeItem) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponseDataAgeItem) SetName(v string) *QueryUserPortraitDataResponseDataAgeItem {
	s.Name = &v
	return s
}

func (s *QueryUserPortraitDataResponseDataAgeItem) SetValue(v int32) *QueryUserPortraitDataResponseDataAgeItem {
	s.Value = &v
	return s
}

type QueryUserPortraitDataResponseDataCityItem struct {
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryUserPortraitDataResponseDataCityItem) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponseDataCityItem) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponseDataCityItem) SetValue(v int32) *QueryUserPortraitDataResponseDataCityItem {
	s.Value = &v
	return s
}

func (s *QueryUserPortraitDataResponseDataCityItem) SetName(v string) *QueryUserPortraitDataResponseDataCityItem {
	s.Name = &v
	return s
}

type QueryUserPortraitDataResponseDataGenderItem struct {
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryUserPortraitDataResponseDataGenderItem) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponseDataGenderItem) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponseDataGenderItem) SetValue(v int32) *QueryUserPortraitDataResponseDataGenderItem {
	s.Value = &v
	return s
}

func (s *QueryUserPortraitDataResponseDataGenderItem) SetName(v string) *QueryUserPortraitDataResponseDataGenderItem {
	s.Name = &v
	return s
}

type QueryUserPortraitDataResponseDataProvinceItem struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryUserPortraitDataResponseDataProvinceItem) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPortraitDataResponseDataProvinceItem) GoString() string {
	return s.String()
}

func (s *QueryUserPortraitDataResponseDataProvinceItem) SetName(v string) *QueryUserPortraitDataResponseDataProvinceItem {
	s.Name = &v
	return s
}

func (s *QueryUserPortraitDataResponseDataProvinceItem) SetValue(v int32) *QueryUserPortraitDataResponseDataProvinceItem {
	s.Value = &v
	return s
}

type QueryVideoDealDataRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s QueryVideoDealDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoDealDataRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoDealDataRequest) SetHeader(v map[string]*string) *QueryVideoDealDataRequest {
	s.Header = v
	return s
}

func (s *QueryVideoDealDataRequest) SetAccessToken(v string) *QueryVideoDealDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryVideoDealDataRequest) SetEndTime(v int64) *QueryVideoDealDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryVideoDealDataRequest) SetHostName(v string) *QueryVideoDealDataRequest {
	s.HostName = &v
	return s
}

func (s *QueryVideoDealDataRequest) SetStartTime(v int64) *QueryVideoDealDataRequest {
	s.StartTime = &v
	return s
}

type QueryVideoDealDataResponse struct {
	Data   *QueryVideoDealDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryVideoDealDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoDealDataResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoDealDataResponse) SetData(v *QueryVideoDealDataResponseData) *QueryVideoDealDataResponse {
	s.Data = v
	return s
}

func (s *QueryVideoDealDataResponse) SetErrNo(v int32) *QueryVideoDealDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryVideoDealDataResponse) SetErrMsg(v string) *QueryVideoDealDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryVideoDealDataResponse) SetLogId(v string) *QueryVideoDealDataResponse {
	s.LogId = &v
	return s
}

type QueryVideoDealDataResponseData struct {
	VideoDealOverviewData *QueryVideoDealDataResponseDataVideoDealOverviewData   `json:"video_deal_overview_data,omitempty" xml:"video_deal_overview_data,omitempty" require:"true"`
	VideoDealDataList     []*QueryVideoDealDataResponseDataVideoDealDataListItem `json:"video_deal_data_list,omitempty" xml:"video_deal_data_list,omitempty" type:"Repeated"`
}

func (s QueryVideoDealDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoDealDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryVideoDealDataResponseData) SetVideoDealOverviewData(v *QueryVideoDealDataResponseDataVideoDealOverviewData) *QueryVideoDealDataResponseData {
	s.VideoDealOverviewData = v
	return s
}

func (s *QueryVideoDealDataResponseData) SetVideoDealDataList(v []*QueryVideoDealDataResponseDataVideoDealDataListItem) *QueryVideoDealDataResponseData {
	s.VideoDealDataList = v
	return s
}

type QueryVideoDealDataResponseDataVideoDealDataListItem struct {
	PayOrderCnt       *int64  `json:"pay_order_cnt,omitempty" xml:"pay_order_cnt,omitempty" require:"true"`
	RefundOrderCnt    *int64  `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
	OrderOncePrice    *int64  `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	MpShowUv          *int64  `json:"mp_show_uv,omitempty" xml:"mp_show_uv,omitempty" require:"true"`
	ValidClueCnt      *int64  `json:"valid_clue_cnt,omitempty" xml:"valid_clue_cnt,omitempty"`
	ValidClueUcnt     *int64  `json:"valid_clue_ucnt,omitempty" xml:"valid_clue_ucnt,omitempty"`
	CreateOrderCnt    *int64  `json:"create_order_cnt,omitempty" xml:"create_order_cnt,omitempty" require:"true"`
	PayCustomerCnt    *int64  `json:"pay_customer_cnt,omitempty" xml:"pay_customer_cnt,omitempty" require:"true"`
	ClueCnt           *int64  `json:"clue_cnt,omitempty" xml:"clue_cnt,omitempty"`
	ClueUcnt          *int64  `json:"clue_ucnt,omitempty" xml:"clue_ucnt,omitempty"`
	CreateCustomerCnt *int64  `json:"create_customer_cnt,omitempty" xml:"create_customer_cnt,omitempty"`
	PayOrderAmount    *int64  `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	MpClickPv         *int64  `json:"mp_click_pv,omitempty" xml:"mp_click_pv,omitempty" require:"true"`
	Time              *string `json:"time,omitempty" xml:"time,omitempty"`
	MpShowPv          *int64  `json:"mp_show_pv,omitempty" xml:"mp_show_pv,omitempty" require:"true"`
	ItemVv            *int64  `json:"item_vv,omitempty" xml:"item_vv,omitempty"`
	MpDrainageUv      *int64  `json:"mp_drainage_uv,omitempty" xml:"mp_drainage_uv,omitempty" require:"true"`
	MpDrainagePv      *int64  `json:"mp_drainage_pv,omitempty" xml:"mp_drainage_pv,omitempty" require:"true"`
	RefundAmount      *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
	VideoPlayCnt      *int64  `json:"video_play_cnt,omitempty" xml:"video_play_cnt,omitempty" require:"true"`
	RefundCustomerCnt *int64  `json:"refund_customer_cnt,omitempty" xml:"refund_customer_cnt,omitempty" require:"true"`
	ItemUv            *int64  `json:"item_uv,omitempty" xml:"item_uv,omitempty"`
	CustomerOncePrice *int64  `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
}

func (s QueryVideoDealDataResponseDataVideoDealDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoDealDataResponseDataVideoDealDataListItem) GoString() string {
	return s.String()
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetPayOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.PayOrderCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetRefundOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.RefundOrderCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetOrderOncePrice(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.OrderOncePrice = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetMpShowUv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.MpShowUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetValidClueCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ValidClueCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetValidClueUcnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ValidClueUcnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetCreateOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.CreateOrderCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetPayCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.PayCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetClueCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ClueCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetClueUcnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ClueUcnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetCreateCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.CreateCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetPayOrderAmount(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetMpClickPv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.MpClickPv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetTime(v string) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.Time = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetMpShowPv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.MpShowPv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetItemVv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ItemVv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetMpDrainageUv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.MpDrainageUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetMpDrainagePv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.MpDrainagePv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetRefundAmount(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.RefundAmount = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetVideoPlayCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.VideoPlayCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetRefundCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.RefundCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetItemUv(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.ItemUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealDataListItem) SetCustomerOncePrice(v int64) *QueryVideoDealDataResponseDataVideoDealDataListItem {
	s.CustomerOncePrice = &v
	return s
}

type QueryVideoDealDataResponseDataVideoDealOverviewData struct {
	OrderOncePrice    *int64 `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	PayOrderCnt       *int64 `json:"pay_order_cnt,omitempty" xml:"pay_order_cnt,omitempty" require:"true"`
	ClueUcnt          *int64 `json:"clue_ucnt,omitempty" xml:"clue_ucnt,omitempty"`
	RefundCustomerCnt *int64 `json:"refund_customer_cnt,omitempty" xml:"refund_customer_cnt,omitempty" require:"true"`
	CreateCustomerCnt *int64 `json:"create_customer_cnt,omitempty" xml:"create_customer_cnt,omitempty"`
	MpShowUv          *int64 `json:"mp_show_uv,omitempty" xml:"mp_show_uv,omitempty" require:"true"`
	MpShowPv          *int64 `json:"mp_show_pv,omitempty" xml:"mp_show_pv,omitempty" require:"true"`
	PayOrderAmount    *int64 `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	PayCustomerCnt    *int64 `json:"pay_customer_cnt,omitempty" xml:"pay_customer_cnt,omitempty" require:"true"`
	CustomerOncePrice *int64 `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
	ClueCnt           *int64 `json:"clue_cnt,omitempty" xml:"clue_cnt,omitempty"`
	MpClickPv         *int64 `json:"mp_click_pv,omitempty" xml:"mp_click_pv,omitempty" require:"true"`
	VideoPlayCnt      *int64 `json:"video_play_cnt,omitempty" xml:"video_play_cnt,omitempty" require:"true"`
	ItemVv            *int64 `json:"item_vv,omitempty" xml:"item_vv,omitempty"`
	RefundAmount      *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
	MpDrainagePv      *int64 `json:"mp_drainage_pv,omitempty" xml:"mp_drainage_pv,omitempty" require:"true"`
	ValidClueCnt      *int64 `json:"valid_clue_cnt,omitempty" xml:"valid_clue_cnt,omitempty"`
	CreateOrderCnt    *int64 `json:"create_order_cnt,omitempty" xml:"create_order_cnt,omitempty" require:"true"`
	MpDrainageUv      *int64 `json:"mp_drainage_uv,omitempty" xml:"mp_drainage_uv,omitempty" require:"true"`
	ItemUv            *int64 `json:"item_uv,omitempty" xml:"item_uv,omitempty"`
	ValidClueUcnt     *int64 `json:"valid_clue_ucnt,omitempty" xml:"valid_clue_ucnt,omitempty"`
	RefundOrderCnt    *int64 `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
}

func (s QueryVideoDealDataResponseDataVideoDealOverviewData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoDealDataResponseDataVideoDealOverviewData) GoString() string {
	return s.String()
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetOrderOncePrice(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.OrderOncePrice = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetPayOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.PayOrderCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetClueUcnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ClueUcnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetRefundCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.RefundCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetCreateCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.CreateCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetMpShowUv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.MpShowUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetMpShowPv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.MpShowPv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetPayOrderAmount(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.PayOrderAmount = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetPayCustomerCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.PayCustomerCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetCustomerOncePrice(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.CustomerOncePrice = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetClueCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ClueCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetMpClickPv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.MpClickPv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetVideoPlayCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.VideoPlayCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetItemVv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ItemVv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetRefundAmount(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.RefundAmount = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetMpDrainagePv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.MpDrainagePv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetValidClueCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ValidClueCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetCreateOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.CreateOrderCnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetMpDrainageUv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.MpDrainageUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetItemUv(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ItemUv = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetValidClueUcnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.ValidClueUcnt = &v
	return s
}

func (s *QueryVideoDealDataResponseDataVideoDealOverviewData) SetRefundOrderCnt(v int64) *QueryVideoDealDataResponseDataVideoDealOverviewData {
	s.RefundOrderCnt = &v
	return s
}

type QueryVideoSumDataRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VideoIds    []*int64           `json:"video_ids,omitempty" xml:"video_ids,omitempty" require:"true" type:"Repeated"`
}

func (s QueryVideoSumDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoSumDataRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoSumDataRequest) SetHeader(v map[string]*string) *QueryVideoSumDataRequest {
	s.Header = v
	return s
}

func (s *QueryVideoSumDataRequest) SetAccessToken(v string) *QueryVideoSumDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryVideoSumDataRequest) SetVideoIds(v []*int64) *QueryVideoSumDataRequest {
	s.VideoIds = v
	return s
}

type QueryVideoSumDataResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryVideoSumDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryVideoSumDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoSumDataResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoSumDataResponse) SetErrNo(v int32) *QueryVideoSumDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryVideoSumDataResponse) SetErrMsg(v string) *QueryVideoSumDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryVideoSumDataResponse) SetLogId(v string) *QueryVideoSumDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryVideoSumDataResponse) SetData(v *QueryVideoSumDataResponseData) *QueryVideoSumDataResponse {
	s.Data = v
	return s
}

type QueryVideoSumDataResponseData struct {
	Data []*QueryVideoSumDataResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryVideoSumDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoSumDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryVideoSumDataResponseData) SetData(v []*QueryVideoSumDataResponseDataDataItem) *QueryVideoSumDataResponseData {
	s.Data = v
	return s
}

type QueryVideoSumDataResponseDataDataItem struct {
	StarServiceProvider *string `json:"star_service_provider,omitempty" xml:"star_service_provider,omitempty"`
	CapitalAccount      *int    `json:"capital_account,omitempty" xml:"capital_account,omitempty"`
	VideoId             *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	TaskId              *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	AgentId             *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	PublishTime         *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
	BillingGmvTd        *int64  `json:"billing_gmv_td,omitempty" xml:"billing_gmv_td,omitempty" require:"true"`
	BilingRefundGmvTd   *int64  `json:"biling_refund_gmv_td,omitempty" xml:"biling_refund_gmv_td,omitempty" require:"true"`
}

func (s QueryVideoSumDataResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoSumDataResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *QueryVideoSumDataResponseDataDataItem) SetStarServiceProvider(v string) *QueryVideoSumDataResponseDataDataItem {
	s.StarServiceProvider = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetCapitalAccount(v int) *QueryVideoSumDataResponseDataDataItem {
	s.CapitalAccount = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetVideoId(v int64) *QueryVideoSumDataResponseDataDataItem {
	s.VideoId = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetTaskId(v int64) *QueryVideoSumDataResponseDataDataItem {
	s.TaskId = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetAgentId(v string) *QueryVideoSumDataResponseDataDataItem {
	s.AgentId = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetPublishTime(v int64) *QueryVideoSumDataResponseDataDataItem {
	s.PublishTime = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetBillingGmvTd(v int64) *QueryVideoSumDataResponseDataDataItem {
	s.BillingGmvTd = &v
	return s
}

func (s *QueryVideoSumDataResponseDataDataItem) SetBilingRefundGmvTd(v int64) *QueryVideoSumDataResponseDataDataItem {
	s.BilingRefundGmvTd = &v
	return s
}

type QueryVideoWithSourceRequest struct {
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime        *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName         *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	AwemeShortIdList []*string          `json:"aweme_short_id_list,omitempty" xml:"aweme_short_id_list,omitempty" type:"Repeated"`
}

func (s QueryVideoWithSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoWithSourceRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoWithSourceRequest) SetHeader(v map[string]*string) *QueryVideoWithSourceRequest {
	s.Header = v
	return s
}

func (s *QueryVideoWithSourceRequest) SetAccessToken(v string) *QueryVideoWithSourceRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryVideoWithSourceRequest) SetStartTime(v int64) *QueryVideoWithSourceRequest {
	s.StartTime = &v
	return s
}

func (s *QueryVideoWithSourceRequest) SetEndTime(v int64) *QueryVideoWithSourceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryVideoWithSourceRequest) SetHostName(v string) *QueryVideoWithSourceRequest {
	s.HostName = &v
	return s
}

func (s *QueryVideoWithSourceRequest) SetAwemeShortIdList(v []*string) *QueryVideoWithSourceRequest {
	s.AwemeShortIdList = v
	return s
}

type QueryVideoWithSourceResponse struct {
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryVideoWithSourceResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryVideoWithSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoWithSourceResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoWithSourceResponse) SetLogId(v string) *QueryVideoWithSourceResponse {
	s.LogId = &v
	return s
}

func (s *QueryVideoWithSourceResponse) SetData(v *QueryVideoWithSourceResponseData) *QueryVideoWithSourceResponse {
	s.Data = v
	return s
}

func (s *QueryVideoWithSourceResponse) SetErrNo(v int32) *QueryVideoWithSourceResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryVideoWithSourceResponse) SetErrMsg(v string) *QueryVideoWithSourceResponse {
	s.ErrMsg = &v
	return s
}

type QueryVideoWithSourceResponseData struct {
	DataList []*QueryVideoWithSourceResponseDataDataListItem `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s QueryVideoWithSourceResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoWithSourceResponseData) GoString() string {
	return s.String()
}

func (s *QueryVideoWithSourceResponseData) SetDataList(v []*QueryVideoWithSourceResponseDataDataListItem) *QueryVideoWithSourceResponseData {
	s.DataList = v
	return s
}

type QueryVideoWithSourceResponseDataDataListItem struct {
	CommonData *QueryVideoWithSourceResponseDataDataListItemCommonData `json:"CommonData,omitempty" xml:"CommonData,omitempty"`
	Scenes     *string                                                 `json:"Scenes,omitempty" xml:"Scenes,omitempty" require:"true"`
	ScenesName *string                                                 `json:"ScenesName,omitempty" xml:"ScenesName,omitempty" require:"true"`
}

func (s QueryVideoWithSourceResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoWithSourceResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *QueryVideoWithSourceResponseDataDataListItem) SetCommonData(v *QueryVideoWithSourceResponseDataDataListItemCommonData) *QueryVideoWithSourceResponseDataDataListItem {
	s.CommonData = v
	return s
}

func (s *QueryVideoWithSourceResponseDataDataListItem) SetScenes(v string) *QueryVideoWithSourceResponseDataDataListItem {
	s.Scenes = &v
	return s
}

func (s *QueryVideoWithSourceResponseDataDataListItem) SetScenesName(v string) *QueryVideoWithSourceResponseDataDataListItem {
	s.ScenesName = &v
	return s
}

type QueryVideoWithSourceResponseDataDataListItemCommonData struct {
	MpDrainageUv *int64 `json:"MpDrainageUv,omitempty" xml:"MpDrainageUv,omitempty" require:"true"`
	MpShowPv     *int64 `json:"MpShowPv,omitempty" xml:"MpShowPv,omitempty" require:"true"`
	MpShowUv     *int64 `json:"MpShowUv,omitempty" xml:"MpShowUv,omitempty" require:"true"`
	MpDrainagePv *int64 `json:"MpDrainagePv,omitempty" xml:"MpDrainagePv,omitempty" require:"true"`
}

func (s QueryVideoWithSourceResponseDataDataListItemCommonData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoWithSourceResponseDataDataListItemCommonData) GoString() string {
	return s.String()
}

func (s *QueryVideoWithSourceResponseDataDataListItemCommonData) SetMpDrainageUv(v int64) *QueryVideoWithSourceResponseDataDataListItemCommonData {
	s.MpDrainageUv = &v
	return s
}

func (s *QueryVideoWithSourceResponseDataDataListItemCommonData) SetMpShowPv(v int64) *QueryVideoWithSourceResponseDataDataListItemCommonData {
	s.MpShowPv = &v
	return s
}

func (s *QueryVideoWithSourceResponseDataDataListItemCommonData) SetMpShowUv(v int64) *QueryVideoWithSourceResponseDataDataListItemCommonData {
	s.MpShowUv = &v
	return s
}

func (s *QueryVideoWithSourceResponseDataDataListItemCommonData) SetMpDrainagePv(v int64) *QueryVideoWithSourceResponseDataDataListItemCommonData {
	s.MpDrainagePv = &v
	return s
}

type QueryViolateTalentListRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryViolateTalentListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryViolateTalentListRequest) GoString() string {
	return s.String()
}

func (s *QueryViolateTalentListRequest) SetHeader(v map[string]*string) *QueryViolateTalentListRequest {
	s.Header = v
	return s
}

func (s *QueryViolateTalentListRequest) SetAccessToken(v string) *QueryViolateTalentListRequest {
	s.AccessToken = &v
	return s
}

type QueryViolateTalentListResponse struct {
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryViolateTalentListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryViolateTalentListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryViolateTalentListResponse) GoString() string {
	return s.String()
}

func (s *QueryViolateTalentListResponse) SetErrMsg(v string) *QueryViolateTalentListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryViolateTalentListResponse) SetLogId(v string) *QueryViolateTalentListResponse {
	s.LogId = &v
	return s
}

func (s *QueryViolateTalentListResponse) SetData(v *QueryViolateTalentListResponseData) *QueryViolateTalentListResponse {
	s.Data = v
	return s
}

func (s *QueryViolateTalentListResponse) SetErrNo(v int32) *QueryViolateTalentListResponse {
	s.ErrNo = &v
	return s
}

type QueryViolateTalentListResponseData struct {
	ViolateTalentsUrl *string `json:"violate_talents_url,omitempty" xml:"violate_talents_url,omitempty" require:"true"`
}

func (s QueryViolateTalentListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryViolateTalentListResponseData) GoString() string {
	return s.String()
}

func (s *QueryViolateTalentListResponseData) SetViolateTalentsUrl(v string) *QueryViolateTalentListResponseData {
	s.ViolateTalentsUrl = &v
	return s
}

type QuizGetRequest struct {
	Type        *int32             `json:"type,omitempty" xml:"type,omitempty"`
	Level       *int32             `json:"level,omitempty" xml:"level,omitempty"`
	Num         *int32             `json:"num,omitempty" xml:"num,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QuizGetRequest) String() string {
	return tea.Prettify(s)
}

func (s QuizGetRequest) GoString() string {
	return s.String()
}

func (s *QuizGetRequest) SetType(v int32) *QuizGetRequest {
	s.Type = &v
	return s
}

func (s *QuizGetRequest) SetLevel(v int32) *QuizGetRequest {
	s.Level = &v
	return s
}

func (s *QuizGetRequest) SetNum(v int32) *QuizGetRequest {
	s.Num = &v
	return s
}

func (s *QuizGetRequest) SetHeader(v map[string]*string) *QuizGetRequest {
	s.Header = v
	return s
}

func (s *QuizGetRequest) SetAccessToken(v string) *QuizGetRequest {
	s.AccessToken = &v
	return s
}

type QuizGetResponse struct {
	Data *QuizGetResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QuizGetResponse) String() string {
	return tea.Prettify(s)
}

func (s QuizGetResponse) GoString() string {
	return s.String()
}

func (s *QuizGetResponse) SetData(v *QuizGetResponseData) *QuizGetResponse {
	s.Data = v
	return s
}

type QuizGetResponseData struct {
	List []*QuizGetResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QuizGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s QuizGetResponseData) GoString() string {
	return s.String()
}

func (s *QuizGetResponseData) SetList(v []*QuizGetResponseDataListItem) *QuizGetResponseData {
	s.List = v
	return s
}

type QuizGetResponseDataListItem struct {
	Options     []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Answer      *int32    `json:"answer,omitempty" xml:"answer,omitempty"`
	ResourceUrl *string   `json:"resource_url,omitempty" xml:"resource_url,omitempty"`
	Id          *string   `json:"id,omitempty" xml:"id,omitempty"`
	Title       *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QuizGetResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QuizGetResponseDataListItem) GoString() string {
	return s.String()
}

func (s *QuizGetResponseDataListItem) SetOptions(v []*string) *QuizGetResponseDataListItem {
	s.Options = v
	return s
}

func (s *QuizGetResponseDataListItem) SetAnswer(v int32) *QuizGetResponseDataListItem {
	s.Answer = &v
	return s
}

func (s *QuizGetResponseDataListItem) SetResourceUrl(v string) *QuizGetResponseDataListItem {
	s.ResourceUrl = &v
	return s
}

func (s *QuizGetResponseDataListItem) SetId(v string) *QuizGetResponseDataListItem {
	s.Id = &v
	return s
}

func (s *QuizGetResponseDataListItem) SetTitle(v string) *QuizGetResponseDataListItem {
	s.Title = &v
	return s
}

type RateplanSaveRequest struct {
	AccountId   *string                      `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	RatePlan    *RateplanSaveRequestRatePlan `json:"rate_plan,omitempty" xml:"rate_plan,omitempty" require:"true"`
	Header      map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RateplanSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequest) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequest) SetAccountId(v string) *RateplanSaveRequest {
	s.AccountId = &v
	return s
}

func (s *RateplanSaveRequest) SetRatePlan(v *RateplanSaveRequestRatePlan) *RateplanSaveRequest {
	s.RatePlan = v
	return s
}

func (s *RateplanSaveRequest) SetHeader(v map[string]*string) *RateplanSaveRequest {
	s.Header = v
	return s
}

func (s *RateplanSaveRequest) SetAccessToken(v string) *RateplanSaveRequest {
	s.AccessToken = &v
	return s
}

type RateplanSaveRequestRatePlan struct {
	HotelId *string                                 `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Rooms   []*RateplanSaveRequestRatePlanRoomsItem `json:"rooms,omitempty" xml:"rooms,omitempty" require:"true" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlan) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlan) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlan) SetHotelId(v string) *RateplanSaveRequestRatePlan {
	s.HotelId = &v
	return s
}

func (s *RateplanSaveRequestRatePlan) SetRooms(v []*RateplanSaveRequestRatePlanRoomsItem) *RateplanSaveRequestRatePlan {
	s.Rooms = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItem struct {
	RoomId    *string                                              `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	RatePlans []*RateplanSaveRequestRatePlanRoomsItemRatePlansItem `json:"rate_plans,omitempty" xml:"rate_plans,omitempty" require:"true" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlanRoomsItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItem) SetRoomId(v string) *RateplanSaveRequestRatePlanRoomsItem {
	s.RoomId = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItem) SetRatePlans(v []*RateplanSaveRequestRatePlanRoomsItemRatePlansItem) *RateplanSaveRequestRatePlanRoomsItem {
	s.RatePlans = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItem struct {
	ConfirmImmediately *bool                                                               `json:"confirm_immediately,omitempty" xml:"confirm_immediately,omitempty"`
	OutRatePlanId      *string                                                             `json:"out_rate_plan_id,omitempty" xml:"out_rate_plan_id,omitempty" require:"true"`
	Invoice            *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice           `json:"invoice,omitempty" xml:"invoice,omitempty"`
	BookRules          *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules         `json:"book_rules,omitempty" xml:"book_rules,omitempty"`
	SalesType          *int                                                                `json:"sales_type,omitempty" xml:"sales_type,omitempty"`
	Meals              []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem       `json:"meals,omitempty" xml:"meals,omitempty" type:"Repeated"`
	SettleType         *int                                                                `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	Member             *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember            `json:"member,omitempty" xml:"member,omitempty"`
	RatePlanId         *string                                                             `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	CancelRules        []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem `json:"cancel_rules,omitempty" xml:"cancel_rules,omitempty" type:"Repeated"`
	BookTimeRules      *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules     `json:"book_time_rules,omitempty" xml:"book_time_rules,omitempty"`
	HourlyRoomDetail   *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail  `json:"hourly_room_detail,omitempty" xml:"hourly_room_detail,omitempty"`
	Policy             *int32                                                              `json:"policy,omitempty" xml:"policy,omitempty"`
	CrowdConfig        []*int                                                              `json:"crowd_config,omitempty" xml:"crowd_config,omitempty" type:"Repeated"`
	StayRules          *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules         `json:"stay_rules,omitempty" xml:"stay_rules,omitempty"`
	SalesTag           *int                                                                `json:"sales_tag,omitempty" xml:"sales_tag,omitempty"`
	RatePlanName       *string                                                             `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty" require:"true"`
	Currency           *string                                                             `json:"currency,omitempty" xml:"currency,omitempty"`
	InvoiceProvider    *int                                                                `json:"invoice_provider,omitempty" xml:"invoice_provider,omitempty"`
	Active             *bool                                                               `json:"active,omitempty" xml:"active,omitempty"`
	Packages           []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem    `json:"packages,omitempty" xml:"packages,omitempty" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetConfirmImmediately(v bool) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.ConfirmImmediately = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetOutRatePlanId(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.OutRatePlanId = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetInvoice(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Invoice = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetBookRules(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.BookRules = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetSalesType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.SalesType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetMeals(v []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Meals = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetSettleType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.SettleType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetMember(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Member = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetRatePlanId(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.RatePlanId = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetCancelRules(v []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.CancelRules = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetBookTimeRules(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.BookTimeRules = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetHourlyRoomDetail(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.HourlyRoomDetail = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetPolicy(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Policy = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetCrowdConfig(v []*int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.CrowdConfig = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetStayRules(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.StayRules = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetSalesTag(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.SalesTag = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetRatePlanName(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.RatePlanName = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetCurrency(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Currency = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetInvoiceProvider(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.InvoiceProvider = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetActive(v bool) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Active = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetPackages(v []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) *RateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Packages = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules struct {
	CheckInFrom      *string                                                                   `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
	CheckInTo        *string                                                                   `json:"check_in_to,omitempty" xml:"check_in_to,omitempty"`
	MinAdvanceTime   *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime `json:"min_advance_time,omitempty" xml:"min_advance_time,omitempty"`
	CheckOutTo       *string                                                                   `json:"check_out_to,omitempty" xml:"check_out_to,omitempty"`
	MaxAdvanceTime   *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime `json:"max_advance_time,omitempty" xml:"max_advance_time,omitempty"`
	MaxQuantityLimt  *int64                                                                    `json:"max_quantity_limt,omitempty" xml:"max_quantity_limt,omitempty"`
	MidnightRoom     *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom   `json:"midnight_room,omitempty" xml:"midnight_room,omitempty"`
	ApplicablePeople []*int                                                                    `json:"applicable_people,omitempty" xml:"applicable_people,omitempty" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetCheckInFrom(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.CheckInFrom = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetCheckInTo(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.CheckInTo = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetMinAdvanceTime(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.MinAdvanceTime = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetCheckOutTo(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.CheckOutTo = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetMaxAdvanceTime(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.MaxAdvanceTime = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetMaxQuantityLimt(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.MaxQuantityLimt = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetMidnightRoom(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.MidnightRoom = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules) SetApplicablePeople(v []*int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRules {
	s.ApplicablePeople = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime struct {
	Day *int32 `json:"day,omitempty" xml:"day,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime) SetDay(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMaxAdvanceTime {
	s.Day = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom struct {
	IsMidnightRoom    *bool  `json:"is_midnight_room,omitempty" xml:"is_midnight_room,omitempty"`
	LatestBookingTime *int64 `json:"latest_booking_time,omitempty" xml:"latest_booking_time,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom) SetIsMidnightRoom(v bool) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom {
	s.IsMidnightRoom = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom) SetLatestBookingTime(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMidnightRoom {
	s.LatestBookingTime = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime struct {
	Day *int32 `json:"day,omitempty" xml:"day,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime) SetDay(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookRulesMinAdvanceTime {
	s.Day = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules struct {
	IsTimeNextday *bool                                                                   `json:"is_time_nextday,omitempty" xml:"is_time_nextday,omitempty" require:"true"`
	TimeSpan      *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan `json:"time_span,omitempty" xml:"time_span,omitempty" require:"true"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules) SetIsTimeNextday(v bool) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules {
	s.IsTimeNextday = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules) SetTimeSpan(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRules {
	s.TimeSpan = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan) SetFrom(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan {
	s.From = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan) SetTo(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemBookTimeRulesTimeSpan {
	s.To = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem struct {
	CancelTimeType   *int                                                                              `json:"cancel_time_type,omitempty" xml:"cancel_time_type,omitempty"`
	CancelType       *int                                                                              `json:"cancel_type,omitempty" xml:"cancel_type,omitempty" require:"true"`
	CutType          *int                                                                              `json:"cut_type,omitempty" xml:"cut_type,omitempty"`
	CutValue         *int64                                                                            `json:"cut_value,omitempty" xml:"cut_value,omitempty"`
	CancelOffsetTime *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime `json:"cancel_offset_time,omitempty" xml:"cancel_offset_time,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) SetCancelTimeType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem {
	s.CancelTimeType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) SetCancelType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem {
	s.CancelType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) SetCutType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem {
	s.CutType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) SetCutValue(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem {
	s.CutValue = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem) SetCancelOffsetTime(v *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItem {
	s.CancelOffsetTime = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime struct {
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) SetDay(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime {
	s.Day = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) SetHour(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime {
	s.Hour = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime) SetMinute(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemCancelRulesItemCancelOffsetTime {
	s.Minute = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail struct {
	LatestCheckOut  *string `json:"latest_check_out,omitempty" xml:"latest_check_out,omitempty" require:"true"`
	UsageDuration   *int64  `json:"usage_duration,omitempty" xml:"usage_duration,omitempty" require:"true"`
	EarliestCheckIn *string `json:"earliest_check_in,omitempty" xml:"earliest_check_in,omitempty" require:"true"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetLatestCheckOut(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.LatestCheckOut = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetUsageDuration(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.UsageDuration = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetEarliestCheckIn(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.EarliestCheckIn = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice struct {
	InvoiceTypes []*int  `json:"invoice_types,omitempty" xml:"invoice_types,omitempty" type:"Repeated"`
	Subject      *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice) SetInvoiceTypes(v []*int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice {
	s.InvoiceTypes = v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice) SetSubject(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemInvoice {
	s.Subject = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem struct {
	Num  *int64 `json:"num,omitempty" xml:"num,omitempty" require:"true"`
	Type *int   `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem) SetNum(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem {
	s.Num = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem) SetType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMealsItem {
	s.Type = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember struct {
	IsMember     *bool    `json:"is_member,omitempty" xml:"is_member,omitempty" require:"true"`
	MemberLevels []*int32 `json:"member_levels,omitempty" xml:"member_levels,omitempty" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember) SetIsMember(v bool) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember {
	s.IsMember = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember) SetMemberLevels(v []*int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemMember {
	s.MemberLevels = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem struct {
	OptionalCount *int32                                                                        `json:"optional_count,omitempty" xml:"optional_count,omitempty"`
	PackageName   *string                                                                       `json:"package_name,omitempty" xml:"package_name,omitempty"`
	PackageType   *int                                                                          `json:"package_type,omitempty" xml:"package_type,omitempty"`
	Policy        *string                                                                       `json:"policy,omitempty" xml:"policy,omitempty"`
	ItemInfos     []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem `json:"item_infos,omitempty" xml:"item_infos,omitempty" type:"Repeated"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) SetOptionalCount(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem {
	s.OptionalCount = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) SetPackageName(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem {
	s.PackageName = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) SetPackageType(v int) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem {
	s.PackageType = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) SetPolicy(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem {
	s.Policy = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem) SetItemInfos(v []*RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItem {
	s.ItemInfos = v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem struct {
	Value       *int64  `json:"value,omitempty" xml:"value,omitempty"`
	Count       *int32  `json:"count,omitempty" xml:"count,omitempty"`
	ImageUrl    *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductInfo *string `json:"product_info,omitempty" xml:"product_info,omitempty"`
	Unit        *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetValue(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.Value = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetCount(v int32) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.Count = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetImageUrl(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.ImageUrl = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetName(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.Name = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetProductInfo(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.ProductInfo = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem) SetUnit(v string) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemPackagesItemItemInfosItem {
	s.Unit = &v
	return s
}

type RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules struct {
	MinLos *int64 `json:"min_los,omitempty" xml:"min_los,omitempty"`
	MaxLos *int64 `json:"max_los,omitempty" xml:"max_los,omitempty"`
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules) GoString() string {
	return s.String()
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules) SetMinLos(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules {
	s.MinLos = &v
	return s
}

func (s *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules) SetMaxLos(v int64) *RateplanSaveRequestRatePlanRoomsItemRatePlansItemStayRules {
	s.MaxLos = &v
	return s
}

type RateplanSaveResponse struct {
	Data  *RateplanSaveResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *RateplanSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s RateplanSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveResponse) GoString() string {
	return s.String()
}

func (s *RateplanSaveResponse) SetData(v *RateplanSaveResponseData) *RateplanSaveResponse {
	s.Data = v
	return s
}

func (s *RateplanSaveResponse) SetExtra(v *RateplanSaveResponseExtra) *RateplanSaveResponse {
	s.Extra = v
	return s
}

type RateplanSaveResponseData struct {
	RatePlanMap   []*RateplanSaveResponseDataRatePlanMapItem `json:"rate_plan_map,omitempty" xml:"rate_plan_map,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RateplanSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveResponseData) GoString() string {
	return s.String()
}

func (s *RateplanSaveResponseData) SetRatePlanMap(v []*RateplanSaveResponseDataRatePlanMapItem) *RateplanSaveResponseData {
	s.RatePlanMap = v
	return s
}

func (s *RateplanSaveResponseData) SetGwErrorCode(v int32) *RateplanSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RateplanSaveResponseData) SetGwDescription(v string) *RateplanSaveResponseData {
	s.GwDescription = &v
	return s
}

type RateplanSaveResponseDataRatePlanMapItem struct {
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message       *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	OutRatePlanId *string `json:"out_rate_plan_id,omitempty" xml:"out_rate_plan_id,omitempty" require:"true"`
	RatePlanId    *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
}

func (s RateplanSaveResponseDataRatePlanMapItem) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveResponseDataRatePlanMapItem) GoString() string {
	return s.String()
}

func (s *RateplanSaveResponseDataRatePlanMapItem) SetCode(v string) *RateplanSaveResponseDataRatePlanMapItem {
	s.Code = &v
	return s
}

func (s *RateplanSaveResponseDataRatePlanMapItem) SetMessage(v string) *RateplanSaveResponseDataRatePlanMapItem {
	s.Message = &v
	return s
}

func (s *RateplanSaveResponseDataRatePlanMapItem) SetOutRatePlanId(v string) *RateplanSaveResponseDataRatePlanMapItem {
	s.OutRatePlanId = &v
	return s
}

func (s *RateplanSaveResponseDataRatePlanMapItem) SetRatePlanId(v string) *RateplanSaveResponseDataRatePlanMapItem {
	s.RatePlanId = &v
	return s
}

type RateplanSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s RateplanSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RateplanSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *RateplanSaveResponseExtra) SetDescription(v string) *RateplanSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *RateplanSaveResponseExtra) SetErrorCode(v int32) *RateplanSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RateplanSaveResponseExtra) SetLogid(v string) *RateplanSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *RateplanSaveResponseExtra) SetNow(v int64) *RateplanSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *RateplanSaveResponseExtra) SetSubDescription(v string) *RateplanSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *RateplanSaveResponseExtra) SetSubErrorCode(v int32) *RateplanSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type RecallMsgRequest struct {
	MsgId            *string            `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId           *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ConversationId   *string            `json:"conversation_id,omitempty" xml:"conversation_id,omitempty" require:"true"`
	ConversationType *int32             `json:"conversation_type,omitempty" xml:"conversation_type,omitempty" require:"true"`
}

func (s RecallMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallMsgRequest) GoString() string {
	return s.String()
}

func (s *RecallMsgRequest) SetMsgId(v string) *RecallMsgRequest {
	s.MsgId = &v
	return s
}

func (s *RecallMsgRequest) SetHeader(v map[string]*string) *RecallMsgRequest {
	s.Header = v
	return s
}

func (s *RecallMsgRequest) SetAccessToken(v string) *RecallMsgRequest {
	s.AccessToken = &v
	return s
}

func (s *RecallMsgRequest) SetOpenId(v string) *RecallMsgRequest {
	s.OpenId = &v
	return s
}

func (s *RecallMsgRequest) SetConversationId(v string) *RecallMsgRequest {
	s.ConversationId = &v
	return s
}

func (s *RecallMsgRequest) SetConversationType(v int32) *RecallMsgRequest {
	s.ConversationType = &v
	return s
}

type RecallMsgResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s RecallMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallMsgResponse) GoString() string {
	return s.String()
}

func (s *RecallMsgResponse) SetLogId(v string) *RecallMsgResponse {
	s.LogId = &v
	return s
}

func (s *RecallMsgResponse) SetErrMsg(v string) *RecallMsgResponse {
	s.ErrMsg = &v
	return s
}

func (s *RecallMsgResponse) SetErrNo(v int32) *RecallMsgResponse {
	s.ErrNo = &v
	return s
}

type RefundApplyRequest struct {
	OperatorId          *int64                                       `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	OrderOutAfterSaleId *string                                      `json:"order_out_after_sale_id,omitempty" xml:"order_out_after_sale_id,omitempty"`
	ApplyRefundItemList []*RefundApplyRequestApplyRefundItemListItem `json:"apply_refund_item_list,omitempty" xml:"apply_refund_item_list,omitempty" type:"Repeated"`
	MerchantDesc        *string                                      `json:"merchant_desc,omitempty" xml:"merchant_desc,omitempty"`
	AccessToken         *string                                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OperatorName        *string                                      `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	OrderId             *string                                      `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RefundReason        *int                                         `json:"refund_reason,omitempty" xml:"refund_reason,omitempty" require:"true"`
	Header              map[string]*string                           `json:"header,omitempty" xml:"header,omitempty"`
}

func (s RefundApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequest) GoString() string {
	return s.String()
}

func (s *RefundApplyRequest) SetOperatorId(v int64) *RefundApplyRequest {
	s.OperatorId = &v
	return s
}

func (s *RefundApplyRequest) SetOrderOutAfterSaleId(v string) *RefundApplyRequest {
	s.OrderOutAfterSaleId = &v
	return s
}

func (s *RefundApplyRequest) SetApplyRefundItemList(v []*RefundApplyRequestApplyRefundItemListItem) *RefundApplyRequest {
	s.ApplyRefundItemList = v
	return s
}

func (s *RefundApplyRequest) SetMerchantDesc(v string) *RefundApplyRequest {
	s.MerchantDesc = &v
	return s
}

func (s *RefundApplyRequest) SetAccessToken(v string) *RefundApplyRequest {
	s.AccessToken = &v
	return s
}

func (s *RefundApplyRequest) SetOperatorName(v string) *RefundApplyRequest {
	s.OperatorName = &v
	return s
}

func (s *RefundApplyRequest) SetOrderId(v string) *RefundApplyRequest {
	s.OrderId = &v
	return s
}

func (s *RefundApplyRequest) SetRefundReason(v int) *RefundApplyRequest {
	s.RefundReason = &v
	return s
}

func (s *RefundApplyRequest) SetHeader(v map[string]*string) *RefundApplyRequest {
	s.Header = v
	return s
}

type RefundApplyRequestApplyRefundItemListItem struct {
	IsNeedRefundItem         *bool                                                                    `json:"is_need_refund_item,omitempty" xml:"is_need_refund_item,omitempty"`
	ItemId                   *string                                                                  `json:"item_id,omitempty" xml:"item_id,omitempty"`
	AffiliatedRefundItemList []*RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem `json:"affiliated_refund_item_list,omitempty" xml:"affiliated_refund_item_list,omitempty" type:"Repeated"`
}

func (s RefundApplyRequestApplyRefundItemListItem) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestApplyRefundItemListItem) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestApplyRefundItemListItem) SetIsNeedRefundItem(v bool) *RefundApplyRequestApplyRefundItemListItem {
	s.IsNeedRefundItem = &v
	return s
}

func (s *RefundApplyRequestApplyRefundItemListItem) SetItemId(v string) *RefundApplyRequestApplyRefundItemListItem {
	s.ItemId = &v
	return s
}

func (s *RefundApplyRequestApplyRefundItemListItem) SetAffiliatedRefundItemList(v []*RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem) *RefundApplyRequestApplyRefundItemListItem {
	s.AffiliatedRefundItemList = v
	return s
}

type RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem struct {
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem) SetItemId(v string) *RefundApplyRequestApplyRefundItemListItemAffiliatedRefundItemListItem {
	s.ItemId = &v
	return s
}

type RefundApplyResponse struct {
	Data  *RefundApplyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *RefundApplyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s RefundApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponse) GoString() string {
	return s.String()
}

func (s *RefundApplyResponse) SetData(v *RefundApplyResponseData) *RefundApplyResponse {
	s.Data = v
	return s
}

func (s *RefundApplyResponse) SetExtra(v *RefundApplyResponseExtra) *RefundApplyResponse {
	s.Extra = v
	return s
}

type RefundApplyResponseData struct {
	Idempotent    *bool   `json:"idempotent,omitempty" xml:"idempotent,omitempty"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty" require:"true"`
	AfterSaleId   *string `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
}

func (s RefundApplyResponseData) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseData) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseData) SetIdempotent(v bool) *RefundApplyResponseData {
	s.Idempotent = &v
	return s
}

func (s *RefundApplyResponseData) SetOrderId(v string) *RefundApplyResponseData {
	s.OrderId = &v
	return s
}

func (s *RefundApplyResponseData) SetGwErrorCode(v int32) *RefundApplyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RefundApplyResponseData) SetGwDescription(v string) *RefundApplyResponseData {
	s.GwDescription = &v
	return s
}

func (s *RefundApplyResponseData) SetSuccess(v bool) *RefundApplyResponseData {
	s.Success = &v
	return s
}

func (s *RefundApplyResponseData) SetAfterSaleId(v string) *RefundApplyResponseData {
	s.AfterSaleId = &v
	return s
}

type RefundApplyResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s RefundApplyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseExtra) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseExtra) SetSubErrorCode(v int32) *RefundApplyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RefundApplyResponseExtra) SetDescription(v string) *RefundApplyResponseExtra {
	s.Description = &v
	return s
}

func (s *RefundApplyResponseExtra) SetErrorCode(v int32) *RefundApplyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RefundApplyResponseExtra) SetLogid(v string) *RefundApplyResponseExtra {
	s.Logid = &v
	return s
}

func (s *RefundApplyResponseExtra) SetNow(v int64) *RefundApplyResponseExtra {
	s.Now = &v
	return s
}

func (s *RefundApplyResponseExtra) SetSubDescription(v string) *RefundApplyResponseExtra {
	s.SubDescription = &v
	return s
}

type RefundAuditRequest struct {
	Reason        *string                              `json:"reason,omitempty" xml:"reason,omitempty"`
	Result        *int                                 `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	Voucher       *string                              `json:"voucher,omitempty" xml:"voucher,omitempty"`
	AfterSaleId   *string                              `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
	Certificate   []*RefundAuditRequestCertificateItem `json:"certificate,omitempty" xml:"certificate,omitempty" type:"Repeated"`
	Header        map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Code          *string                              `json:"code,omitempty" xml:"code,omitempty"`
	CertificateId *string                              `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
}

func (s RefundAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditRequest) GoString() string {
	return s.String()
}

func (s *RefundAuditRequest) SetReason(v string) *RefundAuditRequest {
	s.Reason = &v
	return s
}

func (s *RefundAuditRequest) SetResult(v int) *RefundAuditRequest {
	s.Result = &v
	return s
}

func (s *RefundAuditRequest) SetVoucher(v string) *RefundAuditRequest {
	s.Voucher = &v
	return s
}

func (s *RefundAuditRequest) SetAfterSaleId(v string) *RefundAuditRequest {
	s.AfterSaleId = &v
	return s
}

func (s *RefundAuditRequest) SetCertificate(v []*RefundAuditRequestCertificateItem) *RefundAuditRequest {
	s.Certificate = v
	return s
}

func (s *RefundAuditRequest) SetHeader(v map[string]*string) *RefundAuditRequest {
	s.Header = v
	return s
}

func (s *RefundAuditRequest) SetAccessToken(v string) *RefundAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *RefundAuditRequest) SetCode(v string) *RefundAuditRequest {
	s.Code = &v
	return s
}

func (s *RefundAuditRequest) SetCertificateId(v string) *RefundAuditRequest {
	s.CertificateId = &v
	return s
}

type RefundAuditRequestCertificateItem struct {
	Voucher       *string                                     `json:"voucher,omitempty" xml:"voucher,omitempty"`
	CertificateId *string                                     `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Code          *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	PatchInfo     *RefundAuditRequestCertificateItemPatchInfo `json:"patch_info,omitempty" xml:"patch_info,omitempty"`
}

func (s RefundAuditRequestCertificateItem) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditRequestCertificateItem) GoString() string {
	return s.String()
}

func (s *RefundAuditRequestCertificateItem) SetVoucher(v string) *RefundAuditRequestCertificateItem {
	s.Voucher = &v
	return s
}

func (s *RefundAuditRequestCertificateItem) SetCertificateId(v string) *RefundAuditRequestCertificateItem {
	s.CertificateId = &v
	return s
}

func (s *RefundAuditRequestCertificateItem) SetCode(v string) *RefundAuditRequestCertificateItem {
	s.Code = &v
	return s
}

func (s *RefundAuditRequestCertificateItem) SetPatchInfo(v *RefundAuditRequestCertificateItemPatchInfo) *RefundAuditRequestCertificateItem {
	s.PatchInfo = v
	return s
}

type RefundAuditRequestCertificateItemPatchInfo struct {
	Content   *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	PatchType *int64  `json:"patch_type,omitempty" xml:"patch_type,omitempty" require:"true"`
}

func (s RefundAuditRequestCertificateItemPatchInfo) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditRequestCertificateItemPatchInfo) GoString() string {
	return s.String()
}

func (s *RefundAuditRequestCertificateItemPatchInfo) SetContent(v string) *RefundAuditRequestCertificateItemPatchInfo {
	s.Content = &v
	return s
}

func (s *RefundAuditRequestCertificateItemPatchInfo) SetPatchType(v int64) *RefundAuditRequestCertificateItemPatchInfo {
	s.PatchType = &v
	return s
}

type RefundAuditResponse struct {
	Extra *RefundAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *RefundAuditResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s RefundAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditResponse) GoString() string {
	return s.String()
}

func (s *RefundAuditResponse) SetExtra(v *RefundAuditResponseExtra) *RefundAuditResponse {
	s.Extra = v
	return s
}

func (s *RefundAuditResponse) SetData(v *RefundAuditResponseData) *RefundAuditResponse {
	s.Data = v
	return s
}

type RefundAuditResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RefundAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditResponseData) GoString() string {
	return s.String()
}

func (s *RefundAuditResponseData) SetGwErrorCode(v int32) *RefundAuditResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RefundAuditResponseData) SetGwDescription(v string) *RefundAuditResponseData {
	s.GwDescription = &v
	return s
}

type RefundAuditResponseExtra struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid       *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s RefundAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RefundAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *RefundAuditResponseExtra) SetDescription(v string) *RefundAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *RefundAuditResponseExtra) SetErrorCode(v int32) *RefundAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RefundAuditResponseExtra) SetLogid(v string) *RefundAuditResponseExtra {
	s.Logid = &v
	return s
}

type RelationQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExtIds      []*string          `json:"ext_ids,omitempty" xml:"ext_ids,omitempty" type:"Repeated"`
}

func (s RelationQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s RelationQueryRequest) GoString() string {
	return s.String()
}

func (s *RelationQueryRequest) SetHeader(v map[string]*string) *RelationQueryRequest {
	s.Header = v
	return s
}

func (s *RelationQueryRequest) SetAccessToken(v string) *RelationQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *RelationQueryRequest) SetExtIds(v []*string) *RelationQueryRequest {
	s.ExtIds = v
	return s
}

type RelationQueryResponse struct {
	Extra *RelationQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *RelationQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s RelationQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RelationQueryResponse) GoString() string {
	return s.String()
}

func (s *RelationQueryResponse) SetExtra(v *RelationQueryResponseExtra) *RelationQueryResponse {
	s.Extra = v
	return s
}

func (s *RelationQueryResponse) SetData(v *RelationQueryResponseData) *RelationQueryResponse {
	s.Data = v
	return s
}

type RelationQueryResponseData struct {
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Relations     []*RelationQueryResponseDataRelationsItem `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s RelationQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s RelationQueryResponseData) GoString() string {
	return s.String()
}

func (s *RelationQueryResponseData) SetGwDescription(v string) *RelationQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *RelationQueryResponseData) SetRelations(v []*RelationQueryResponseDataRelationsItem) *RelationQueryResponseData {
	s.Relations = v
	return s
}

func (s *RelationQueryResponseData) SetGwErrorCode(v int32) *RelationQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type RelationQueryResponseDataRelationsItem struct {
	MatchRelationStatus *int64  `json:"match_relation_status,omitempty" xml:"match_relation_status,omitempty"`
	ExtId               *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	MatchPoiId          *string `json:"match_poi_id,omitempty" xml:"match_poi_id,omitempty"`
}

func (s RelationQueryResponseDataRelationsItem) String() string {
	return tea.Prettify(s)
}

func (s RelationQueryResponseDataRelationsItem) GoString() string {
	return s.String()
}

func (s *RelationQueryResponseDataRelationsItem) SetMatchRelationStatus(v int64) *RelationQueryResponseDataRelationsItem {
	s.MatchRelationStatus = &v
	return s
}

func (s *RelationQueryResponseDataRelationsItem) SetExtId(v string) *RelationQueryResponseDataRelationsItem {
	s.ExtId = &v
	return s
}

func (s *RelationQueryResponseDataRelationsItem) SetMatchPoiId(v string) *RelationQueryResponseDataRelationsItem {
	s.MatchPoiId = &v
	return s
}

type RelationQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s RelationQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RelationQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *RelationQueryResponseExtra) SetSubDescription(v string) *RelationQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *RelationQueryResponseExtra) SetSubErrorCode(v int32) *RelationQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RelationQueryResponseExtra) SetDescription(v string) *RelationQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *RelationQueryResponseExtra) SetErrorCode(v int32) *RelationQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RelationQueryResponseExtra) SetLogid(v string) *RelationQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *RelationQueryResponseExtra) SetNow(v int64) *RelationQueryResponseExtra {
	s.Now = &v
	return s
}

type RemaintimesDecrRequest struct {
	ServiceId     *string            `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	ServiceModeId *string            `json:"service_mode_id,omitempty" xml:"service_mode_id,omitempty" require:"true"`
	OutTradeNo    *string            `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty" require:"true"`
	Count         *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	IsTestEnv     *bool              `json:"is_test_env,omitempty" xml:"is_test_env,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s RemaintimesDecrRequest) String() string {
	return tea.Prettify(s)
}

func (s RemaintimesDecrRequest) GoString() string {
	return s.String()
}

func (s *RemaintimesDecrRequest) SetServiceId(v string) *RemaintimesDecrRequest {
	s.ServiceId = &v
	return s
}

func (s *RemaintimesDecrRequest) SetServiceModeId(v string) *RemaintimesDecrRequest {
	s.ServiceModeId = &v
	return s
}

func (s *RemaintimesDecrRequest) SetOutTradeNo(v string) *RemaintimesDecrRequest {
	s.OutTradeNo = &v
	return s
}

func (s *RemaintimesDecrRequest) SetCount(v int32) *RemaintimesDecrRequest {
	s.Count = &v
	return s
}

func (s *RemaintimesDecrRequest) SetIsTestEnv(v bool) *RemaintimesDecrRequest {
	s.IsTestEnv = &v
	return s
}

func (s *RemaintimesDecrRequest) SetHeader(v map[string]*string) *RemaintimesDecrRequest {
	s.Header = v
	return s
}

func (s *RemaintimesDecrRequest) SetAccessToken(v string) *RemaintimesDecrRequest {
	s.AccessToken = &v
	return s
}

func (s *RemaintimesDecrRequest) SetOpenId(v string) *RemaintimesDecrRequest {
	s.OpenId = &v
	return s
}

type RemaintimesDecrResponse struct {
	Data  *RemaintimesDecrResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *RemaintimesDecrResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s RemaintimesDecrResponse) String() string {
	return tea.Prettify(s)
}

func (s RemaintimesDecrResponse) GoString() string {
	return s.String()
}

func (s *RemaintimesDecrResponse) SetData(v *RemaintimesDecrResponseData) *RemaintimesDecrResponse {
	s.Data = v
	return s
}

func (s *RemaintimesDecrResponse) SetExtra(v *RemaintimesDecrResponseExtra) *RemaintimesDecrResponse {
	s.Extra = v
	return s
}

type RemaintimesDecrResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RemaintimesDecrResponseData) String() string {
	return tea.Prettify(s)
}

func (s RemaintimesDecrResponseData) GoString() string {
	return s.String()
}

func (s *RemaintimesDecrResponseData) SetGwErrorCode(v int32) *RemaintimesDecrResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RemaintimesDecrResponseData) SetGwDescription(v string) *RemaintimesDecrResponseData {
	s.GwDescription = &v
	return s
}

type RemaintimesDecrResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RemaintimesDecrResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RemaintimesDecrResponseExtra) GoString() string {
	return s.String()
}

func (s *RemaintimesDecrResponseExtra) SetSubErrorCode(v int32) *RemaintimesDecrResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RemaintimesDecrResponseExtra) SetSubDescription(v string) *RemaintimesDecrResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *RemaintimesDecrResponseExtra) SetLogid(v string) *RemaintimesDecrResponseExtra {
	s.Logid = &v
	return s
}

func (s *RemaintimesDecrResponseExtra) SetNow(v int64) *RemaintimesDecrResponseExtra {
	s.Now = &v
	return s
}

func (s *RemaintimesDecrResponseExtra) SetErrorCode(v int32) *RemaintimesDecrResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RemaintimesDecrResponseExtra) SetDescription(v string) *RemaintimesDecrResponseExtra {
	s.Description = &v
	return s
}

type ReplyReplyUserTextRequest struct {
	ConversationId *string            `json:"conversation_id,omitempty" xml:"conversation_id,omitempty" require:"true"`
	MsgId          *string            `json:"msg_id,omitempty" xml:"msg_id,omitempty" require:"true"`
	SenderName     *string            `json:"sender_name,omitempty" xml:"sender_name,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CreateTime     *int64             `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	MsgType        *string            `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
	Content        *string            `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	MicroGameId    *string            `json:"micro_game_id,omitempty" xml:"micro_game_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ReplyReplyUserTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplyReplyUserTextRequest) GoString() string {
	return s.String()
}

func (s *ReplyReplyUserTextRequest) SetConversationId(v string) *ReplyReplyUserTextRequest {
	s.ConversationId = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetMsgId(v string) *ReplyReplyUserTextRequest {
	s.MsgId = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetSenderName(v string) *ReplyReplyUserTextRequest {
	s.SenderName = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetAccessToken(v string) *ReplyReplyUserTextRequest {
	s.AccessToken = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetCreateTime(v int64) *ReplyReplyUserTextRequest {
	s.CreateTime = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetMsgType(v string) *ReplyReplyUserTextRequest {
	s.MsgType = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetContent(v string) *ReplyReplyUserTextRequest {
	s.Content = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetMicroGameId(v string) *ReplyReplyUserTextRequest {
	s.MicroGameId = &v
	return s
}

func (s *ReplyReplyUserTextRequest) SetHeader(v map[string]*string) *ReplyReplyUserTextRequest {
	s.Header = v
	return s
}

type ReplyReplyUserTextResponse struct {
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	Data   *ReplyReplyUserTextResponseData `json:"data,omitempty" xml:"data,omitempty"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ReplyReplyUserTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplyReplyUserTextResponse) GoString() string {
	return s.String()
}

func (s *ReplyReplyUserTextResponse) SetErrMsg(v string) *ReplyReplyUserTextResponse {
	s.ErrMsg = &v
	return s
}

func (s *ReplyReplyUserTextResponse) SetData(v *ReplyReplyUserTextResponseData) *ReplyReplyUserTextResponse {
	s.Data = v
	return s
}

func (s *ReplyReplyUserTextResponse) SetLogId(v string) *ReplyReplyUserTextResponse {
	s.LogId = &v
	return s
}

func (s *ReplyReplyUserTextResponse) SetErrNo(v int32) *ReplyReplyUserTextResponse {
	s.ErrNo = &v
	return s
}

type ReplyReplyUserTextResponseData struct {
	FilterReason *int32  `json:"filter_reason,omitempty" xml:"filter_reason,omitempty"`
	ImMsgId      *string `json:"im_msg_id,omitempty" xml:"im_msg_id,omitempty"`
}

func (s ReplyReplyUserTextResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReplyReplyUserTextResponseData) GoString() string {
	return s.String()
}

func (s *ReplyReplyUserTextResponseData) SetFilterReason(v int32) *ReplyReplyUserTextResponseData {
	s.FilterReason = &v
	return s
}

func (s *ReplyReplyUserTextResponseData) SetImMsgId(v string) *ReplyReplyUserTextResponseData {
	s.ImMsgId = &v
	return s
}

type ReportTaskCreateRequest struct {
	PoiId               *string             `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	Longitude           *string             `json:"longitude,omitempty" xml:"longitude,omitempty"`
	AccountId           *string             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header              map[string]*string  `json:"header,omitempty" xml:"header,omitempty"`
	Address             *string             `json:"address,omitempty" xml:"address,omitempty"`
	TypeCode            *string             `json:"type_code,omitempty" xml:"type_code,omitempty"`
	City                *string             `json:"city,omitempty" xml:"city,omitempty"`
	TelList             []*string           `json:"tel_list,omitempty" xml:"tel_list,omitempty" type:"Repeated"`
	Latitude            *string             `json:"latitude,omitempty" xml:"latitude,omitempty"`
	PoiName             *string             `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	AdditionalImageUrls []*string           `json:"additional_image_urls,omitempty" xml:"additional_image_urls,omitempty" type:"Repeated"`
	OpenStatus          *int32              `json:"open_status,omitempty" xml:"open_status,omitempty"`
	Province            *string             `json:"province,omitempty" xml:"province,omitempty"`
	OpenTimes           map[int32][]*string `json:"open_times,omitempty" xml:"open_times,omitempty"`
	AdditionalInfo      *string             `json:"additional_info,omitempty" xml:"additional_info,omitempty"`
	AccessToken         *string             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ReportTaskCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskCreateRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskCreateRequest) SetPoiId(v string) *ReportTaskCreateRequest {
	s.PoiId = &v
	return s
}

func (s *ReportTaskCreateRequest) SetLongitude(v string) *ReportTaskCreateRequest {
	s.Longitude = &v
	return s
}

func (s *ReportTaskCreateRequest) SetAccountId(v string) *ReportTaskCreateRequest {
	s.AccountId = &v
	return s
}

func (s *ReportTaskCreateRequest) SetHeader(v map[string]*string) *ReportTaskCreateRequest {
	s.Header = v
	return s
}

func (s *ReportTaskCreateRequest) SetAddress(v string) *ReportTaskCreateRequest {
	s.Address = &v
	return s
}

func (s *ReportTaskCreateRequest) SetTypeCode(v string) *ReportTaskCreateRequest {
	s.TypeCode = &v
	return s
}

func (s *ReportTaskCreateRequest) SetCity(v string) *ReportTaskCreateRequest {
	s.City = &v
	return s
}

func (s *ReportTaskCreateRequest) SetTelList(v []*string) *ReportTaskCreateRequest {
	s.TelList = v
	return s
}

func (s *ReportTaskCreateRequest) SetLatitude(v string) *ReportTaskCreateRequest {
	s.Latitude = &v
	return s
}

func (s *ReportTaskCreateRequest) SetPoiName(v string) *ReportTaskCreateRequest {
	s.PoiName = &v
	return s
}

func (s *ReportTaskCreateRequest) SetAdditionalImageUrls(v []*string) *ReportTaskCreateRequest {
	s.AdditionalImageUrls = v
	return s
}

func (s *ReportTaskCreateRequest) SetOpenStatus(v int32) *ReportTaskCreateRequest {
	s.OpenStatus = &v
	return s
}

func (s *ReportTaskCreateRequest) SetProvince(v string) *ReportTaskCreateRequest {
	s.Province = &v
	return s
}

func (s *ReportTaskCreateRequest) SetOpenTimes(v map[int32][]*string) *ReportTaskCreateRequest {
	s.OpenTimes = v
	return s
}

func (s *ReportTaskCreateRequest) SetAdditionalInfo(v string) *ReportTaskCreateRequest {
	s.AdditionalInfo = &v
	return s
}

func (s *ReportTaskCreateRequest) SetAccessToken(v string) *ReportTaskCreateRequest {
	s.AccessToken = &v
	return s
}

type ReportTaskCreateResponse struct {
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	Extra  *ReportTaskCreateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ReportTaskCreateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ReportTaskCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskCreateResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskCreateResponse) SetErrMsg(v string) *ReportTaskCreateResponse {
	s.ErrMsg = &v
	return s
}

func (s *ReportTaskCreateResponse) SetErrNo(v int32) *ReportTaskCreateResponse {
	s.ErrNo = &v
	return s
}

func (s *ReportTaskCreateResponse) SetExtra(v *ReportTaskCreateResponseExtra) *ReportTaskCreateResponse {
	s.Extra = v
	return s
}

func (s *ReportTaskCreateResponse) SetLogId(v string) *ReportTaskCreateResponse {
	s.LogId = &v
	return s
}

func (s *ReportTaskCreateResponse) SetData(v *ReportTaskCreateResponseData) *ReportTaskCreateResponse {
	s.Data = v
	return s
}

type ReportTaskCreateResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	FailMessage   *string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	TaskId        *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ReportTaskCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskCreateResponseData) GoString() string {
	return s.String()
}

func (s *ReportTaskCreateResponseData) SetGwErrorCode(v int32) *ReportTaskCreateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ReportTaskCreateResponseData) SetGwDescription(v string) *ReportTaskCreateResponseData {
	s.GwDescription = &v
	return s
}

func (s *ReportTaskCreateResponseData) SetFailMessage(v string) *ReportTaskCreateResponseData {
	s.FailMessage = &v
	return s
}

func (s *ReportTaskCreateResponseData) SetTaskId(v string) *ReportTaskCreateResponseData {
	s.TaskId = &v
	return s
}

type ReportTaskCreateResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ReportTaskCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *ReportTaskCreateResponseExtra) SetSubErrorCode(v int32) *ReportTaskCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ReportTaskCreateResponseExtra) SetDescription(v string) *ReportTaskCreateResponseExtra {
	s.Description = &v
	return s
}

func (s *ReportTaskCreateResponseExtra) SetErrorCode(v int32) *ReportTaskCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ReportTaskCreateResponseExtra) SetLogid(v string) *ReportTaskCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *ReportTaskCreateResponseExtra) SetNow(v int64) *ReportTaskCreateResponseExtra {
	s.Now = &v
	return s
}

func (s *ReportTaskCreateResponseExtra) SetSubDescription(v string) *ReportTaskCreateResponseExtra {
	s.SubDescription = &v
	return s
}

type ReportTaskViewRequest struct {
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PoiId       *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ReportTaskViewRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskViewRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskViewRequest) SetTaskId(v string) *ReportTaskViewRequest {
	s.TaskId = &v
	return s
}

func (s *ReportTaskViewRequest) SetAccountId(v string) *ReportTaskViewRequest {
	s.AccountId = &v
	return s
}

func (s *ReportTaskViewRequest) SetPoiId(v string) *ReportTaskViewRequest {
	s.PoiId = &v
	return s
}

func (s *ReportTaskViewRequest) SetHeader(v map[string]*string) *ReportTaskViewRequest {
	s.Header = v
	return s
}

func (s *ReportTaskViewRequest) SetAccessToken(v string) *ReportTaskViewRequest {
	s.AccessToken = &v
	return s
}

type ReportTaskViewResponse struct {
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	Extra  *ReportTaskViewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ReportTaskViewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ReportTaskViewResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskViewResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskViewResponse) SetErrMsg(v string) *ReportTaskViewResponse {
	s.ErrMsg = &v
	return s
}

func (s *ReportTaskViewResponse) SetErrNo(v int32) *ReportTaskViewResponse {
	s.ErrNo = &v
	return s
}

func (s *ReportTaskViewResponse) SetExtra(v *ReportTaskViewResponseExtra) *ReportTaskViewResponse {
	s.Extra = v
	return s
}

func (s *ReportTaskViewResponse) SetLogId(v string) *ReportTaskViewResponse {
	s.LogId = &v
	return s
}

func (s *ReportTaskViewResponse) SetData(v *ReportTaskViewResponseData) *ReportTaskViewResponse {
	s.Data = v
	return s
}

type ReportTaskViewResponseData struct {
	RejectReason  *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Status        *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReportTaskViewResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskViewResponseData) GoString() string {
	return s.String()
}

func (s *ReportTaskViewResponseData) SetRejectReason(v string) *ReportTaskViewResponseData {
	s.RejectReason = &v
	return s
}

func (s *ReportTaskViewResponseData) SetGwErrorCode(v int32) *ReportTaskViewResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ReportTaskViewResponseData) SetGwDescription(v string) *ReportTaskViewResponseData {
	s.GwDescription = &v
	return s
}

func (s *ReportTaskViewResponseData) SetStatus(v int32) *ReportTaskViewResponseData {
	s.Status = &v
	return s
}

type ReportTaskViewResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ReportTaskViewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskViewResponseExtra) GoString() string {
	return s.String()
}

func (s *ReportTaskViewResponseExtra) SetSubErrorCode(v int32) *ReportTaskViewResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ReportTaskViewResponseExtra) SetDescription(v string) *ReportTaskViewResponseExtra {
	s.Description = &v
	return s
}

func (s *ReportTaskViewResponseExtra) SetErrorCode(v int32) *ReportTaskViewResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ReportTaskViewResponseExtra) SetLogid(v string) *ReportTaskViewResponseExtra {
	s.Logid = &v
	return s
}

func (s *ReportTaskViewResponseExtra) SetNow(v int64) *ReportTaskViewResponseExtra {
	s.Now = &v
	return s
}

func (s *ReportTaskViewResponseExtra) SetSubDescription(v string) *ReportTaskViewResponseExtra {
	s.SubDescription = &v
	return s
}

type ReserveCodeBatchImportRequest struct {
	SkuId       *string            `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	ThirdSkuId  *string            `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Codes       []*string          `json:"codes,omitempty" xml:"codes,omitempty" require:"true" type:"Repeated"`
	ExpiredTime *int64             `json:"expired_time,omitempty" xml:"expired_time,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ReserveCodeBatchImportRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBatchImportRequest) GoString() string {
	return s.String()
}

func (s *ReserveCodeBatchImportRequest) SetSkuId(v string) *ReserveCodeBatchImportRequest {
	s.SkuId = &v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetThirdSkuId(v string) *ReserveCodeBatchImportRequest {
	s.ThirdSkuId = &v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetAccountId(v string) *ReserveCodeBatchImportRequest {
	s.AccountId = &v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetCodes(v []*string) *ReserveCodeBatchImportRequest {
	s.Codes = v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetExpiredTime(v int64) *ReserveCodeBatchImportRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetHeader(v map[string]*string) *ReserveCodeBatchImportRequest {
	s.Header = v
	return s
}

func (s *ReserveCodeBatchImportRequest) SetAccessToken(v string) *ReserveCodeBatchImportRequest {
	s.AccessToken = &v
	return s
}

type ReserveCodeBatchImportResponse struct {
	Extra *ReserveCodeBatchImportResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ReserveCodeBatchImportResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ReserveCodeBatchImportResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBatchImportResponse) GoString() string {
	return s.String()
}

func (s *ReserveCodeBatchImportResponse) SetExtra(v *ReserveCodeBatchImportResponseExtra) *ReserveCodeBatchImportResponse {
	s.Extra = v
	return s
}

func (s *ReserveCodeBatchImportResponse) SetData(v *ReserveCodeBatchImportResponseData) *ReserveCodeBatchImportResponse {
	s.Data = v
	return s
}

type ReserveCodeBatchImportResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ReserveCodeBatchImportResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBatchImportResponseData) GoString() string {
	return s.String()
}

func (s *ReserveCodeBatchImportResponseData) SetGwErrorCode(v int32) *ReserveCodeBatchImportResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ReserveCodeBatchImportResponseData) SetGwDescription(v string) *ReserveCodeBatchImportResponseData {
	s.GwDescription = &v
	return s
}

type ReserveCodeBatchImportResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ReserveCodeBatchImportResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBatchImportResponseExtra) GoString() string {
	return s.String()
}

func (s *ReserveCodeBatchImportResponseExtra) SetLogid(v string) *ReserveCodeBatchImportResponseExtra {
	s.Logid = &v
	return s
}

func (s *ReserveCodeBatchImportResponseExtra) SetNow(v int64) *ReserveCodeBatchImportResponseExtra {
	s.Now = &v
	return s
}

func (s *ReserveCodeBatchImportResponseExtra) SetSubDescription(v string) *ReserveCodeBatchImportResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ReserveCodeBatchImportResponseExtra) SetSubErrorCode(v int32) *ReserveCodeBatchImportResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ReserveCodeBatchImportResponseExtra) SetDescription(v string) *ReserveCodeBatchImportResponseExtra {
	s.Description = &v
	return s
}

func (s *ReserveCodeBatchImportResponseExtra) SetErrorCode(v int32) *ReserveCodeBatchImportResponseExtra {
	s.ErrorCode = &v
	return s
}

type ReserveCodeBindOrderInfoRequest struct {
	SkuId       *string            `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Code        *string            `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s ReserveCodeBindOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoRequest) SetSkuId(v string) *ReserveCodeBindOrderInfoRequest {
	s.SkuId = &v
	return s
}

func (s *ReserveCodeBindOrderInfoRequest) SetAccountId(v string) *ReserveCodeBindOrderInfoRequest {
	s.AccountId = &v
	return s
}

func (s *ReserveCodeBindOrderInfoRequest) SetHeader(v map[string]*string) *ReserveCodeBindOrderInfoRequest {
	s.Header = v
	return s
}

func (s *ReserveCodeBindOrderInfoRequest) SetAccessToken(v string) *ReserveCodeBindOrderInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *ReserveCodeBindOrderInfoRequest) SetCode(v string) *ReserveCodeBindOrderInfoRequest {
	s.Code = &v
	return s
}

type ReserveCodeBindOrderInfoResponse struct {
	Extra *ReserveCodeBindOrderInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ReserveCodeBindOrderInfoResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ReserveCodeBindOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponse) SetExtra(v *ReserveCodeBindOrderInfoResponseExtra) *ReserveCodeBindOrderInfoResponse {
	s.Extra = v
	return s
}

func (s *ReserveCodeBindOrderInfoResponse) SetData(v *ReserveCodeBindOrderInfoResponseData) *ReserveCodeBindOrderInfoResponse {
	s.Data = v
	return s
}

type ReserveCodeBindOrderInfoResponseData struct {
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Order         *ReserveCodeBindOrderInfoResponseDataOrder `json:"order,omitempty" xml:"order,omitempty"`
}

func (s ReserveCodeBindOrderInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponseData) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponseData) SetGwErrorCode(v int32) *ReserveCodeBindOrderInfoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseData) SetGwDescription(v string) *ReserveCodeBindOrderInfoResponseData {
	s.GwDescription = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseData) SetOrder(v *ReserveCodeBindOrderInfoResponseDataOrder) *ReserveCodeBindOrderInfoResponseData {
	s.Order = v
	return s
}

type ReserveCodeBindOrderInfoResponseDataOrder struct {
	CreateTime *int64                                                   `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	OrderId    *string                                                  `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	SkuId      *string                                                  `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	ThirdSkuId *string                                                  `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Tourists   []*ReserveCodeBindOrderInfoResponseDataOrderTouristsItem `json:"tourists,omitempty" xml:"tourists,omitempty" type:"Repeated"`
	Amount     *ReserveCodeBindOrderInfoResponseDataOrderAmount         `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	Codes      []*string                                                `json:"codes,omitempty" xml:"codes,omitempty" require:"true" type:"Repeated"`
}

func (s ReserveCodeBindOrderInfoResponseDataOrder) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponseDataOrder) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetCreateTime(v int64) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.CreateTime = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetOrderId(v string) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.OrderId = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetSkuId(v string) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.SkuId = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetThirdSkuId(v string) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.ThirdSkuId = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetTourists(v []*ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.Tourists = v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetAmount(v *ReserveCodeBindOrderInfoResponseDataOrderAmount) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.Amount = v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrder) SetCodes(v []*string) *ReserveCodeBindOrderInfoResponseDataOrder {
	s.Codes = v
	return s
}

type ReserveCodeBindOrderInfoResponseDataOrderAmount struct {
	PlatformTicketAmount  *int32 `json:"platform_ticket_amount,omitempty" xml:"platform_ticket_amount,omitempty"`
	FeeAmount             *int32 `json:"fee_amount,omitempty" xml:"fee_amount,omitempty"`
	MerchantTicketAmount  *int32 `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginalAmount        *int32 `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PayAmount             *int32 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PaymentDiscountAmount *int32 `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
}

func (s ReserveCodeBindOrderInfoResponseDataOrderAmount) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponseDataOrderAmount) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetPlatformTicketAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.PlatformTicketAmount = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetFeeAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.FeeAmount = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetMerchantTicketAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetOriginalAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.OriginalAmount = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetPayAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.PayAmount = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderAmount) SetPaymentDiscountAmount(v int32) *ReserveCodeBindOrderInfoResponseDataOrderAmount {
	s.PaymentDiscountAmount = &v
	return s
}

type ReserveCodeBindOrderInfoResponseDataOrderTouristsItem struct {
	IdCard *string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone  *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) SetIdCard(v string) *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem {
	s.IdCard = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) SetName(v string) *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem {
	s.Name = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem) SetPhone(v string) *ReserveCodeBindOrderInfoResponseDataOrderTouristsItem {
	s.Phone = &v
	return s
}

type ReserveCodeBindOrderInfoResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ReserveCodeBindOrderInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ReserveCodeBindOrderInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetNow(v int64) *ReserveCodeBindOrderInfoResponseExtra {
	s.Now = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetSubDescription(v string) *ReserveCodeBindOrderInfoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetSubErrorCode(v int32) *ReserveCodeBindOrderInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetDescription(v string) *ReserveCodeBindOrderInfoResponseExtra {
	s.Description = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetErrorCode(v int32) *ReserveCodeBindOrderInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ReserveCodeBindOrderInfoResponseExtra) SetLogid(v string) *ReserveCodeBindOrderInfoResponseExtra {
	s.Logid = &v
	return s
}

type RetailOrderConfirmRequest struct {
	OrderId       *string                                `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	PickUpCode    *string                                `json:"pick_up_code,omitempty" xml:"pick_up_code,omitempty"`
	PickUpTime    *int64                                 `json:"pick_up_time,omitempty" xml:"pick_up_time,omitempty"`
	RejectReason  *RetailOrderConfirmRequestRejectReason `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Header        map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ConfirmResult *int                                   `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
}

func (s RetailOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmRequest) SetOrderId(v string) *RetailOrderConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *RetailOrderConfirmRequest) SetPickUpCode(v string) *RetailOrderConfirmRequest {
	s.PickUpCode = &v
	return s
}

func (s *RetailOrderConfirmRequest) SetPickUpTime(v int64) *RetailOrderConfirmRequest {
	s.PickUpTime = &v
	return s
}

func (s *RetailOrderConfirmRequest) SetRejectReason(v *RetailOrderConfirmRequestRejectReason) *RetailOrderConfirmRequest {
	s.RejectReason = v
	return s
}

func (s *RetailOrderConfirmRequest) SetHeader(v map[string]*string) *RetailOrderConfirmRequest {
	s.Header = v
	return s
}

func (s *RetailOrderConfirmRequest) SetAccessToken(v string) *RetailOrderConfirmRequest {
	s.AccessToken = &v
	return s
}

func (s *RetailOrderConfirmRequest) SetConfirmResult(v int) *RetailOrderConfirmRequest {
	s.ConfirmResult = &v
	return s
}

type RetailOrderConfirmRequestRejectReason struct {
	ShowReason    []*RetailOrderConfirmRequestRejectReasonShowReasonItem `json:"show_reason,omitempty" xml:"show_reason,omitempty" type:"Repeated"`
	UserMediaList []*string                                              `json:"user_media_list,omitempty" xml:"user_media_list,omitempty" type:"Repeated"`
	Desc          *string                                                `json:"desc,omitempty" xml:"desc,omitempty"`
	ReasonCode    []*int64                                               `json:"reason_code,omitempty" xml:"reason_code,omitempty" require:"true" type:"Repeated"`
}

func (s RetailOrderConfirmRequestRejectReason) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmRequestRejectReason) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmRequestRejectReason) SetShowReason(v []*RetailOrderConfirmRequestRejectReasonShowReasonItem) *RetailOrderConfirmRequestRejectReason {
	s.ShowReason = v
	return s
}

func (s *RetailOrderConfirmRequestRejectReason) SetUserMediaList(v []*string) *RetailOrderConfirmRequestRejectReason {
	s.UserMediaList = v
	return s
}

func (s *RetailOrderConfirmRequestRejectReason) SetDesc(v string) *RetailOrderConfirmRequestRejectReason {
	s.Desc = &v
	return s
}

func (s *RetailOrderConfirmRequestRejectReason) SetReasonCode(v []*int64) *RetailOrderConfirmRequestRejectReason {
	s.ReasonCode = v
	return s
}

type RetailOrderConfirmRequestRejectReasonShowReasonItem struct {
	ReasonCode *int64  `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	Msg        *string `json:"msg,omitempty" xml:"msg,omitempty"`
}

func (s RetailOrderConfirmRequestRejectReasonShowReasonItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmRequestRejectReasonShowReasonItem) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmRequestRejectReasonShowReasonItem) SetReasonCode(v int64) *RetailOrderConfirmRequestRejectReasonShowReasonItem {
	s.ReasonCode = &v
	return s
}

func (s *RetailOrderConfirmRequestRejectReasonShowReasonItem) SetMsg(v string) *RetailOrderConfirmRequestRejectReasonShowReasonItem {
	s.Msg = &v
	return s
}

type RetailOrderConfirmResponse struct {
	Data  *RetailOrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *RetailOrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s RetailOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmResponse) SetData(v *RetailOrderConfirmResponseData) *RetailOrderConfirmResponse {
	s.Data = v
	return s
}

func (s *RetailOrderConfirmResponse) SetExtra(v *RetailOrderConfirmResponseExtra) *RetailOrderConfirmResponse {
	s.Extra = v
	return s
}

type RetailOrderConfirmResponseData struct {
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s RetailOrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmResponseData) SetSuccess(v bool) *RetailOrderConfirmResponseData {
	s.Success = &v
	return s
}

func (s *RetailOrderConfirmResponseData) SetGwErrorCode(v int32) *RetailOrderConfirmResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RetailOrderConfirmResponseData) SetGwDescription(v string) *RetailOrderConfirmResponseData {
	s.GwDescription = &v
	return s
}

func (s *RetailOrderConfirmResponseData) SetOrderId(v string) *RetailOrderConfirmResponseData {
	s.OrderId = &v
	return s
}

type RetailOrderConfirmResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RetailOrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *RetailOrderConfirmResponseExtra) SetErrorCode(v int32) *RetailOrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RetailOrderConfirmResponseExtra) SetLogid(v string) *RetailOrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *RetailOrderConfirmResponseExtra) SetNow(v int64) *RetailOrderConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *RetailOrderConfirmResponseExtra) SetSubDescription(v string) *RetailOrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *RetailOrderConfirmResponseExtra) SetSubErrorCode(v int32) *RetailOrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RetailOrderConfirmResponseExtra) SetDescription(v string) *RetailOrderConfirmResponseExtra {
	s.Description = &v
	return s
}

type RetailOrderQueryRequest struct {
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RetailOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryRequest) SetOrderId(v string) *RetailOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *RetailOrderQueryRequest) SetAccountId(v string) *RetailOrderQueryRequest {
	s.AccountId = &v
	return s
}

func (s *RetailOrderQueryRequest) SetHeader(v map[string]*string) *RetailOrderQueryRequest {
	s.Header = v
	return s
}

func (s *RetailOrderQueryRequest) SetAccessToken(v string) *RetailOrderQueryRequest {
	s.AccessToken = &v
	return s
}

type RetailOrderQueryResponse struct {
	Data  *RetailOrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *RetailOrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s RetailOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponse) SetData(v *RetailOrderQueryResponseData) *RetailOrderQueryResponse {
	s.Data = v
	return s
}

func (s *RetailOrderQueryResponse) SetExtra(v *RetailOrderQueryResponseExtra) *RetailOrderQueryResponse {
	s.Extra = v
	return s
}

type RetailOrderQueryResponseData struct {
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Orders        []*RetailOrderQueryResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
}

func (s RetailOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseData) SetGwErrorCode(v int32) *RetailOrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RetailOrderQueryResponseData) SetGwDescription(v string) *RetailOrderQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *RetailOrderQueryResponseData) SetOrders(v []*RetailOrderQueryResponseDataOrdersItem) *RetailOrderQueryResponseData {
	s.Orders = v
	return s
}

type RetailOrderQueryResponseDataOrdersItem struct {
	RetailOrderBase *RetailOrderQueryResponseDataOrdersItemRetailOrderBase     `json:"retail_order_base,omitempty" xml:"retail_order_base,omitempty"`
	RetailSkuList   []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItem `json:"retail_sku_list,omitempty" xml:"retail_sku_list,omitempty" type:"Repeated"`
	RetailAmount    *RetailOrderQueryResponseDataOrdersItemRetailAmount        `json:"retail_amount,omitempty" xml:"retail_amount,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItem) SetRetailOrderBase(v *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) *RetailOrderQueryResponseDataOrdersItem {
	s.RetailOrderBase = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItem) SetRetailSkuList(v []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) *RetailOrderQueryResponseDataOrdersItem {
	s.RetailSkuList = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItem) SetRetailAmount(v *RetailOrderQueryResponseDataOrdersItemRetailAmount) *RetailOrderQueryResponseDataOrdersItem {
	s.RetailAmount = v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailAmount struct {
	DeductAmount           *int64 `json:"deduct_amount,omitempty" xml:"deduct_amount,omitempty"`
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	OriginAmount           *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	PayAmount              *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailAmount) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailAmount) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailAmount) SetDeductAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailAmount {
	s.DeductAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailAmount) SetMerchantDiscountAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailAmount {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailAmount) SetOriginAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailAmount {
	s.OriginAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailAmount) SetPayAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailAmount {
	s.PayAmount = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailOrderBase struct {
	PoiId             *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	RetailOrderStatus *int    `json:"retail_order_status,omitempty" xml:"retail_order_status,omitempty"`
	RetailOrderType   *int    `json:"retail_order_type,omitempty" xml:"retail_order_type,omitempty"`
	AccountId         *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OpenId            *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	OrderId           *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderOutId        *string `json:"order_out_id,omitempty" xml:"order_out_id,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailOrderBase) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailOrderBase) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetPoiId(v string) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.PoiId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetRetailOrderStatus(v int) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.RetailOrderStatus = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetRetailOrderType(v int) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.RetailOrderType = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetAccountId(v string) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.AccountId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetOpenId(v string) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.OpenId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetOrderId(v string) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.OrderId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailOrderBase) SetOrderOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailOrderBase {
	s.OrderOutId = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItem struct {
	ProductId              *string                                                                `json:"product_id,omitempty" xml:"product_id,omitempty"`
	MerchantDiscountAmount *int64                                                                 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	SkuAttr                []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem  `json:"sku_attr,omitempty" xml:"sku_attr,omitempty" type:"Repeated"`
	ProductName            *string                                                                `json:"product_name,omitempty" xml:"product_name,omitempty"`
	SkuOutId               *string                                                                `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	ProductOutId           *string                                                                `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	ItemList               []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	SkuName                *string                                                                `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuId                  *string                                                                `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	UnitAmount             *int64                                                                 `json:"unit_amount,omitempty" xml:"unit_amount,omitempty"`
	Count                  *int32                                                                 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetProductId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.ProductId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetMerchantDiscountAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetSkuAttr(v []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.SkuAttr = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetProductName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.ProductName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetSkuOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.SkuOutId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetProductOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.ProductOutId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetItemList(v []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.ItemList = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetSkuName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.SkuName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetSkuId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.SkuId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetUnitAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.UnitAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem) SetCount(v int32) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItem {
	s.Count = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem struct {
	AffiliatedSkuList      []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem `json:"affiliated_sku_list,omitempty" xml:"affiliated_sku_list,omitempty" type:"Repeated"`
	CommodityWeight        *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight         `json:"commodity_weight,omitempty" xml:"commodity_weight,omitempty"`
	DeductInfo             *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo              `json:"deduct_info,omitempty" xml:"deduct_info,omitempty"`
	ItemOrderId            *string                                                                                     `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	MerchantDiscountAmount *int64                                                                                      `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) SetAffiliatedSkuList(v []*RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem {
	s.AffiliatedSkuList = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) SetCommodityWeight(v *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem {
	s.CommodityWeight = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) SetDeductInfo(v *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem {
	s.DeductInfo = v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) SetItemOrderId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem {
	s.ItemOrderId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem) SetMerchantDiscountAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItem {
	s.MerchantDiscountAmount = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem struct {
	TotalAmount  *int64  `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	Count        *int32  `json:"count,omitempty" xml:"count,omitempty"`
	ProductId    *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName  *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductOutId *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	SkuId        *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName      *string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuOutId     *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetTotalAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.TotalAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetCount(v int32) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.Count = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetProductId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.ProductId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetProductName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.ProductName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetProductOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.ProductOutId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetSkuId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.SkuId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetSkuName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.SkuName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem) SetSkuOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemAffiliatedSkuListItem {
	s.SkuOutId = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight struct {
	Count *int32  `json:"count,omitempty" xml:"count,omitempty"`
	Unit  *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight) SetCount(v int32) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight {
	s.Count = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight) SetUnit(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemCommodityWeight {
	s.Unit = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo struct {
	SourceCertificateId          *string `json:"source_certificate_id,omitempty" xml:"source_certificate_id,omitempty"`
	SourceOriginAmount           *int64  `json:"source_origin_amount,omitempty" xml:"source_origin_amount,omitempty"`
	SourceProductOutId           *string `json:"source_product_out_id,omitempty" xml:"source_product_out_id,omitempty"`
	SourceSkuId                  *int64  `json:"source_sku_id,omitempty" xml:"source_sku_id,omitempty"`
	GrouponType                  *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SourceProductId              *int64  `json:"source_product_id,omitempty" xml:"source_product_id,omitempty"`
	SourceProductName            *string `json:"source_product_name,omitempty" xml:"source_product_name,omitempty"`
	SourceThirdPartCode          *string `json:"source_third_part_code,omitempty" xml:"source_third_part_code,omitempty"`
	SourceMerchantDiscountAmount *int64  `json:"source_merchant_discount_amount,omitempty" xml:"source_merchant_discount_amount,omitempty"`
	SourceSkuName                *string `json:"source_sku_name,omitempty" xml:"source_sku_name,omitempty"`
	DeductAmount                 *int64  `json:"deduct_amount,omitempty" xml:"deduct_amount,omitempty"`
	SerialNum                    *int64  `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	VerifyId                     *int64  `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	SourceSkuOutId               *string `json:"source_sku_out_id,omitempty" xml:"source_sku_out_id,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceCertificateId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceCertificateId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceOriginAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceOriginAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceProductOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceProductOutId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceSkuId(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceSkuId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetGrouponType(v int) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.GrouponType = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceProductId(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceProductId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceProductName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceProductName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceThirdPartCode(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceThirdPartCode = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceMerchantDiscountAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceMerchantDiscountAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceSkuName(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceSkuName = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetDeductAmount(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.DeductAmount = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSerialNum(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SerialNum = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetVerifyId(v int64) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.VerifyId = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo) SetSourceSkuOutId(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemItemListItemDeductInfo {
	s.SourceSkuOutId = &v
	return s
}

type RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem struct {
	AttrValue *string `json:"attr_value,omitempty" xml:"attr_value,omitempty"`
	AttrKey   *string `json:"attr_key,omitempty" xml:"attr_key,omitempty"`
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem) SetAttrValue(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem {
	s.AttrValue = &v
	return s
}

func (s *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem) SetAttrKey(v string) *RetailOrderQueryResponseDataOrdersItemRetailSkuListItemSkuAttrItem {
	s.AttrKey = &v
	return s
}

type RetailOrderQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s RetailOrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *RetailOrderQueryResponseExtra) SetNow(v int64) *RetailOrderQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *RetailOrderQueryResponseExtra) SetSubDescription(v string) *RetailOrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *RetailOrderQueryResponseExtra) SetSubErrorCode(v int32) *RetailOrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RetailOrderQueryResponseExtra) SetDescription(v string) *RetailOrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *RetailOrderQueryResponseExtra) SetErrorCode(v int32) *RetailOrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RetailOrderQueryResponseExtra) SetLogid(v string) *RetailOrderQueryResponseExtra {
	s.Logid = &v
	return s
}

type RetailOrderRefundAuditRequest struct {
	OrderId        *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutAfterSaleId *string            `json:"out_after_sale_id,omitempty" xml:"out_after_sale_id,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ReasonCode     *int64             `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
	ReasonMsg      *string            `json:"reason_msg,omitempty" xml:"reason_msg,omitempty"`
	AfterSaleId    *string            `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
	AuditResult    *int               `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
}

func (s RetailOrderRefundAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderRefundAuditRequest) GoString() string {
	return s.String()
}

func (s *RetailOrderRefundAuditRequest) SetOrderId(v string) *RetailOrderRefundAuditRequest {
	s.OrderId = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetOutAfterSaleId(v string) *RetailOrderRefundAuditRequest {
	s.OutAfterSaleId = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetHeader(v map[string]*string) *RetailOrderRefundAuditRequest {
	s.Header = v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetAccessToken(v string) *RetailOrderRefundAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetReasonCode(v int64) *RetailOrderRefundAuditRequest {
	s.ReasonCode = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetReasonMsg(v string) *RetailOrderRefundAuditRequest {
	s.ReasonMsg = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetAfterSaleId(v string) *RetailOrderRefundAuditRequest {
	s.AfterSaleId = &v
	return s
}

func (s *RetailOrderRefundAuditRequest) SetAuditResult(v int) *RetailOrderRefundAuditRequest {
	s.AuditResult = &v
	return s
}

type RetailOrderRefundAuditResponse struct {
	Extra *RetailOrderRefundAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *RetailOrderRefundAuditResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s RetailOrderRefundAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderRefundAuditResponse) GoString() string {
	return s.String()
}

func (s *RetailOrderRefundAuditResponse) SetExtra(v *RetailOrderRefundAuditResponseExtra) *RetailOrderRefundAuditResponse {
	s.Extra = v
	return s
}

func (s *RetailOrderRefundAuditResponse) SetData(v *RetailOrderRefundAuditResponseData) *RetailOrderRefundAuditResponse {
	s.Data = v
	return s
}

type RetailOrderRefundAuditResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RetailOrderRefundAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderRefundAuditResponseData) GoString() string {
	return s.String()
}

func (s *RetailOrderRefundAuditResponseData) SetGwErrorCode(v int32) *RetailOrderRefundAuditResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RetailOrderRefundAuditResponseData) SetGwDescription(v string) *RetailOrderRefundAuditResponseData {
	s.GwDescription = &v
	return s
}

type RetailOrderRefundAuditResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s RetailOrderRefundAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RetailOrderRefundAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *RetailOrderRefundAuditResponseExtra) SetSubErrorCode(v int32) *RetailOrderRefundAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RetailOrderRefundAuditResponseExtra) SetDescription(v string) *RetailOrderRefundAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *RetailOrderRefundAuditResponseExtra) SetErrorCode(v int32) *RetailOrderRefundAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RetailOrderRefundAuditResponseExtra) SetLogid(v string) *RetailOrderRefundAuditResponseExtra {
	s.Logid = &v
	return s
}

func (s *RetailOrderRefundAuditResponseExtra) SetNow(v int64) *RetailOrderRefundAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *RetailOrderRefundAuditResponseExtra) SetSubDescription(v string) *RetailOrderRefundAuditResponseExtra {
	s.SubDescription = &v
	return s
}

type RoundCompleteUploadUserResultRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AnchorOpenId *string            `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty" require:"true"`
	RoomId       *string            `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	CompleteTime *int64             `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	RoundId      *int64             `json:"round_id,omitempty" xml:"round_id,omitempty" require:"true"`
}

func (s RoundCompleteUploadUserResultRequest) String() string {
	return tea.Prettify(s)
}

func (s RoundCompleteUploadUserResultRequest) GoString() string {
	return s.String()
}

func (s *RoundCompleteUploadUserResultRequest) SetHeader(v map[string]*string) *RoundCompleteUploadUserResultRequest {
	s.Header = v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetAccessToken(v string) *RoundCompleteUploadUserResultRequest {
	s.AccessToken = &v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetAnchorOpenId(v string) *RoundCompleteUploadUserResultRequest {
	s.AnchorOpenId = &v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetRoomId(v string) *RoundCompleteUploadUserResultRequest {
	s.RoomId = &v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetCompleteTime(v int64) *RoundCompleteUploadUserResultRequest {
	s.CompleteTime = &v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetAppId(v string) *RoundCompleteUploadUserResultRequest {
	s.AppId = &v
	return s
}

func (s *RoundCompleteUploadUserResultRequest) SetRoundId(v int64) *RoundCompleteUploadUserResultRequest {
	s.RoundId = &v
	return s
}

type RoundCompleteUploadUserResultResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s RoundCompleteUploadUserResultResponse) String() string {
	return tea.Prettify(s)
}

func (s RoundCompleteUploadUserResultResponse) GoString() string {
	return s.String()
}

func (s *RoundCompleteUploadUserResultResponse) SetErrMsg(v string) *RoundCompleteUploadUserResultResponse {
	s.ErrMsg = &v
	return s
}

func (s *RoundCompleteUploadUserResultResponse) SetErrNo(v int64) *RoundCompleteUploadUserResultResponse {
	s.ErrNo = &v
	return s
}

type RoundInfoRequest struct {
	RoomId        *string                              `json:"room_id,omitempty" xml:"room_id,omitempty"`
	EndStatus     *int64                               `json:"end_status,omitempty" xml:"end_status,omitempty"`
	AppId         *string                              `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Header        map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	StartTime     *int64                               `json:"start_time,omitempty" xml:"start_time,omitempty"`
	AnchorOpenId  *string                              `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty"`
	EndTime       *int64                               `json:"end_time,omitempty" xml:"end_time,omitempty"`
	MvpList       []*RoundInfoRequestMvpListItem       `json:"mvp_list,omitempty" xml:"mvp_list,omitempty" type:"Repeated"`
	EndUserList   []*RoundInfoRequestEndUserListItem   `json:"end_user_list,omitempty" xml:"end_user_list,omitempty" type:"Repeated"`
	StartUserList []*RoundInfoRequestStartUserListItem `json:"start_user_list,omitempty" xml:"start_user_list,omitempty" type:"Repeated"`
	RoundId       *string                              `json:"round_id,omitempty" xml:"round_id,omitempty"`
	AccessToken   *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RoundInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s RoundInfoRequest) GoString() string {
	return s.String()
}

func (s *RoundInfoRequest) SetRoomId(v string) *RoundInfoRequest {
	s.RoomId = &v
	return s
}

func (s *RoundInfoRequest) SetEndStatus(v int64) *RoundInfoRequest {
	s.EndStatus = &v
	return s
}

func (s *RoundInfoRequest) SetAppId(v string) *RoundInfoRequest {
	s.AppId = &v
	return s
}

func (s *RoundInfoRequest) SetHeader(v map[string]*string) *RoundInfoRequest {
	s.Header = v
	return s
}

func (s *RoundInfoRequest) SetStartTime(v int64) *RoundInfoRequest {
	s.StartTime = &v
	return s
}

func (s *RoundInfoRequest) SetAnchorOpenId(v string) *RoundInfoRequest {
	s.AnchorOpenId = &v
	return s
}

func (s *RoundInfoRequest) SetEndTime(v int64) *RoundInfoRequest {
	s.EndTime = &v
	return s
}

func (s *RoundInfoRequest) SetMvpList(v []*RoundInfoRequestMvpListItem) *RoundInfoRequest {
	s.MvpList = v
	return s
}

func (s *RoundInfoRequest) SetEndUserList(v []*RoundInfoRequestEndUserListItem) *RoundInfoRequest {
	s.EndUserList = v
	return s
}

func (s *RoundInfoRequest) SetStartUserList(v []*RoundInfoRequestStartUserListItem) *RoundInfoRequest {
	s.StartUserList = v
	return s
}

func (s *RoundInfoRequest) SetRoundId(v string) *RoundInfoRequest {
	s.RoundId = &v
	return s
}

func (s *RoundInfoRequest) SetAccessToken(v string) *RoundInfoRequest {
	s.AccessToken = &v
	return s
}

type RoundInfoRequestEndUserListItem struct {
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserOpenId *string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	Diamonds   *int64  `json:"diamonds,omitempty" xml:"diamonds,omitempty"`
}

func (s RoundInfoRequestEndUserListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundInfoRequestEndUserListItem) GoString() string {
	return s.String()
}

func (s *RoundInfoRequestEndUserListItem) SetOrderId(v string) *RoundInfoRequestEndUserListItem {
	s.OrderId = &v
	return s
}

func (s *RoundInfoRequestEndUserListItem) SetUserOpenId(v string) *RoundInfoRequestEndUserListItem {
	s.UserOpenId = &v
	return s
}

func (s *RoundInfoRequestEndUserListItem) SetDiamonds(v int64) *RoundInfoRequestEndUserListItem {
	s.Diamonds = &v
	return s
}

type RoundInfoRequestMvpListItem struct {
	UserOpenId *string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	Rank       *int64  `json:"rank,omitempty" xml:"rank,omitempty"`
}

func (s RoundInfoRequestMvpListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundInfoRequestMvpListItem) GoString() string {
	return s.String()
}

func (s *RoundInfoRequestMvpListItem) SetUserOpenId(v string) *RoundInfoRequestMvpListItem {
	s.UserOpenId = &v
	return s
}

func (s *RoundInfoRequestMvpListItem) SetRank(v int64) *RoundInfoRequestMvpListItem {
	s.Rank = &v
	return s
}

type RoundInfoRequestStartUserListItem struct {
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UserOpenId *string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
	Diamonds   *int64  `json:"diamonds,omitempty" xml:"diamonds,omitempty"`
}

func (s RoundInfoRequestStartUserListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundInfoRequestStartUserListItem) GoString() string {
	return s.String()
}

func (s *RoundInfoRequestStartUserListItem) SetOrderId(v string) *RoundInfoRequestStartUserListItem {
	s.OrderId = &v
	return s
}

func (s *RoundInfoRequestStartUserListItem) SetUserOpenId(v string) *RoundInfoRequestStartUserListItem {
	s.UserOpenId = &v
	return s
}

func (s *RoundInfoRequestStartUserListItem) SetDiamonds(v int64) *RoundInfoRequestStartUserListItem {
	s.Diamonds = &v
	return s
}

type RoundInfoResponse struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RoundInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RoundInfoResponse) GoString() string {
	return s.String()
}

func (s *RoundInfoResponse) SetStatus(v string) *RoundInfoResponse {
	s.Status = &v
	return s
}

type RoundSyncStatusRequest struct {
	EndTime         *int64                                       `json:"end_time,omitempty" xml:"end_time,omitempty"`
	GroupResultList []*RoundSyncStatusRequestGroupResultListItem `json:"group_result_list,omitempty" xml:"group_result_list,omitempty" type:"Repeated"`
	RoundId         *int64                                       `json:"round_id,omitempty" xml:"round_id,omitempty" require:"true"`
	RoomId          *string                                      `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	AnchorOpenId    *string                                      `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty" require:"true"`
	StartTime       *int64                                       `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	Status          *int                                         `json:"status,omitempty" xml:"status,omitempty"`
	AppId           *string                                      `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header          map[string]*string                           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RoundSyncStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s RoundSyncStatusRequest) GoString() string {
	return s.String()
}

func (s *RoundSyncStatusRequest) SetEndTime(v int64) *RoundSyncStatusRequest {
	s.EndTime = &v
	return s
}

func (s *RoundSyncStatusRequest) SetGroupResultList(v []*RoundSyncStatusRequestGroupResultListItem) *RoundSyncStatusRequest {
	s.GroupResultList = v
	return s
}

func (s *RoundSyncStatusRequest) SetRoundId(v int64) *RoundSyncStatusRequest {
	s.RoundId = &v
	return s
}

func (s *RoundSyncStatusRequest) SetRoomId(v string) *RoundSyncStatusRequest {
	s.RoomId = &v
	return s
}

func (s *RoundSyncStatusRequest) SetAnchorOpenId(v string) *RoundSyncStatusRequest {
	s.AnchorOpenId = &v
	return s
}

func (s *RoundSyncStatusRequest) SetStartTime(v int64) *RoundSyncStatusRequest {
	s.StartTime = &v
	return s
}

func (s *RoundSyncStatusRequest) SetStatus(v int) *RoundSyncStatusRequest {
	s.Status = &v
	return s
}

func (s *RoundSyncStatusRequest) SetAppId(v string) *RoundSyncStatusRequest {
	s.AppId = &v
	return s
}

func (s *RoundSyncStatusRequest) SetHeader(v map[string]*string) *RoundSyncStatusRequest {
	s.Header = v
	return s
}

func (s *RoundSyncStatusRequest) SetAccessToken(v string) *RoundSyncStatusRequest {
	s.AccessToken = &v
	return s
}

type RoundSyncStatusRequestGroupResultListItem struct {
	Result  *int    `json:"result,omitempty" xml:"result,omitempty"`
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s RoundSyncStatusRequestGroupResultListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundSyncStatusRequestGroupResultListItem) GoString() string {
	return s.String()
}

func (s *RoundSyncStatusRequestGroupResultListItem) SetResult(v int) *RoundSyncStatusRequestGroupResultListItem {
	s.Result = &v
	return s
}

func (s *RoundSyncStatusRequestGroupResultListItem) SetGroupId(v string) *RoundSyncStatusRequestGroupResultListItem {
	s.GroupId = &v
	return s
}

type RoundSyncStatusResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s RoundSyncStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s RoundSyncStatusResponse) GoString() string {
	return s.String()
}

func (s *RoundSyncStatusResponse) SetErrNo(v int64) *RoundSyncStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *RoundSyncStatusResponse) SetErrMsg(v string) *RoundSyncStatusResponse {
	s.ErrMsg = &v
	return s
}

type RoundUploadRankListRequest struct {
	RoomId       *string                                   `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	RankList     []*RoundUploadRankListRequestRankListItem `json:"rank_list,omitempty" xml:"rank_list,omitempty" require:"true" type:"Repeated"`
	AppId        *string                                   `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	RoundId      *int64                                    `json:"round_id,omitempty" xml:"round_id,omitempty" require:"true"`
	AnchorOpenId *string                                   `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty" require:"true"`
	Header       map[string]*string                        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RoundUploadRankListRequest) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadRankListRequest) GoString() string {
	return s.String()
}

func (s *RoundUploadRankListRequest) SetRoomId(v string) *RoundUploadRankListRequest {
	s.RoomId = &v
	return s
}

func (s *RoundUploadRankListRequest) SetRankList(v []*RoundUploadRankListRequestRankListItem) *RoundUploadRankListRequest {
	s.RankList = v
	return s
}

func (s *RoundUploadRankListRequest) SetAppId(v string) *RoundUploadRankListRequest {
	s.AppId = &v
	return s
}

func (s *RoundUploadRankListRequest) SetRoundId(v int64) *RoundUploadRankListRequest {
	s.RoundId = &v
	return s
}

func (s *RoundUploadRankListRequest) SetAnchorOpenId(v string) *RoundUploadRankListRequest {
	s.AnchorOpenId = &v
	return s
}

func (s *RoundUploadRankListRequest) SetHeader(v map[string]*string) *RoundUploadRankListRequest {
	s.Header = v
	return s
}

func (s *RoundUploadRankListRequest) SetAccessToken(v string) *RoundUploadRankListRequest {
	s.AccessToken = &v
	return s
}

type RoundUploadRankListRequestRankListItem struct {
	WinningPoints      *int64  `json:"winning_points,omitempty" xml:"winning_points,omitempty"`
	RoundResult        *int    `json:"round_result,omitempty" xml:"round_result,omitempty" require:"true"`
	OpenId             *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Rank               *int64  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	Score              *int64  `json:"score,omitempty" xml:"score,omitempty" require:"true"`
	WinningStreakCount *int64  `json:"winning_streak_count,omitempty" xml:"winning_streak_count,omitempty"`
}

func (s RoundUploadRankListRequestRankListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadRankListRequestRankListItem) GoString() string {
	return s.String()
}

func (s *RoundUploadRankListRequestRankListItem) SetWinningPoints(v int64) *RoundUploadRankListRequestRankListItem {
	s.WinningPoints = &v
	return s
}

func (s *RoundUploadRankListRequestRankListItem) SetRoundResult(v int) *RoundUploadRankListRequestRankListItem {
	s.RoundResult = &v
	return s
}

func (s *RoundUploadRankListRequestRankListItem) SetOpenId(v string) *RoundUploadRankListRequestRankListItem {
	s.OpenId = &v
	return s
}

func (s *RoundUploadRankListRequestRankListItem) SetRank(v int64) *RoundUploadRankListRequestRankListItem {
	s.Rank = &v
	return s
}

func (s *RoundUploadRankListRequestRankListItem) SetScore(v int64) *RoundUploadRankListRequestRankListItem {
	s.Score = &v
	return s
}

func (s *RoundUploadRankListRequestRankListItem) SetWinningStreakCount(v int64) *RoundUploadRankListRequestRankListItem {
	s.WinningStreakCount = &v
	return s
}

type RoundUploadRankListResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s RoundUploadRankListResponse) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadRankListResponse) GoString() string {
	return s.String()
}

func (s *RoundUploadRankListResponse) SetErrMsg(v string) *RoundUploadRankListResponse {
	s.ErrMsg = &v
	return s
}

func (s *RoundUploadRankListResponse) SetErrNo(v int64) *RoundUploadRankListResponse {
	s.ErrNo = &v
	return s
}

type RoundUploadUserResultRequest struct {
	AppId        *string                                     `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	RoundId      *int64                                      `json:"round_id,omitempty" xml:"round_id,omitempty" require:"true"`
	AnchorOpenId *string                                     `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty" require:"true"`
	RoomId       *string                                     `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	Header       map[string]*string                          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UserList     []*RoundUploadUserResultRequestUserListItem `json:"user_list,omitempty" xml:"user_list,omitempty" require:"true" type:"Repeated"`
}

func (s RoundUploadUserResultRequest) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadUserResultRequest) GoString() string {
	return s.String()
}

func (s *RoundUploadUserResultRequest) SetAppId(v string) *RoundUploadUserResultRequest {
	s.AppId = &v
	return s
}

func (s *RoundUploadUserResultRequest) SetRoundId(v int64) *RoundUploadUserResultRequest {
	s.RoundId = &v
	return s
}

func (s *RoundUploadUserResultRequest) SetAnchorOpenId(v string) *RoundUploadUserResultRequest {
	s.AnchorOpenId = &v
	return s
}

func (s *RoundUploadUserResultRequest) SetRoomId(v string) *RoundUploadUserResultRequest {
	s.RoomId = &v
	return s
}

func (s *RoundUploadUserResultRequest) SetHeader(v map[string]*string) *RoundUploadUserResultRequest {
	s.Header = v
	return s
}

func (s *RoundUploadUserResultRequest) SetAccessToken(v string) *RoundUploadUserResultRequest {
	s.AccessToken = &v
	return s
}

func (s *RoundUploadUserResultRequest) SetUserList(v []*RoundUploadUserResultRequestUserListItem) *RoundUploadUserResultRequest {
	s.UserList = v
	return s
}

type RoundUploadUserResultRequestUserListItem struct {
	WinningPoints      *int64  `json:"winning_points,omitempty" xml:"winning_points,omitempty"`
	RoundResult        *int    `json:"round_result,omitempty" xml:"round_result,omitempty" require:"true"`
	OpenId             *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Rank               *int64  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	Score              *int64  `json:"score,omitempty" xml:"score,omitempty" require:"true"`
	WinningStreakCount *int64  `json:"winning_streak_count,omitempty" xml:"winning_streak_count,omitempty"`
}

func (s RoundUploadUserResultRequestUserListItem) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadUserResultRequestUserListItem) GoString() string {
	return s.String()
}

func (s *RoundUploadUserResultRequestUserListItem) SetWinningPoints(v int64) *RoundUploadUserResultRequestUserListItem {
	s.WinningPoints = &v
	return s
}

func (s *RoundUploadUserResultRequestUserListItem) SetRoundResult(v int) *RoundUploadUserResultRequestUserListItem {
	s.RoundResult = &v
	return s
}

func (s *RoundUploadUserResultRequestUserListItem) SetOpenId(v string) *RoundUploadUserResultRequestUserListItem {
	s.OpenId = &v
	return s
}

func (s *RoundUploadUserResultRequestUserListItem) SetRank(v int64) *RoundUploadUserResultRequestUserListItem {
	s.Rank = &v
	return s
}

func (s *RoundUploadUserResultRequestUserListItem) SetScore(v int64) *RoundUploadUserResultRequestUserListItem {
	s.Score = &v
	return s
}

func (s *RoundUploadUserResultRequestUserListItem) SetWinningStreakCount(v int64) *RoundUploadUserResultRequestUserListItem {
	s.WinningStreakCount = &v
	return s
}

type RoundUploadUserResultResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s RoundUploadUserResultResponse) String() string {
	return tea.Prettify(s)
}

func (s RoundUploadUserResultResponse) GoString() string {
	return s.String()
}

func (s *RoundUploadUserResultResponse) SetErrNo(v int64) *RoundUploadUserResultResponse {
	s.ErrNo = &v
	return s
}

func (s *RoundUploadUserResultResponse) SetErrMsg(v string) *RoundUploadUserResultResponse {
	s.ErrMsg = &v
	return s
}

type RtEcpmQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DateHour    *string            `json:"date_hour,omitempty" xml:"date_hour,omitempty"`
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	StartDate   *string            `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate     *string            `json:"end_date,omitempty" xml:"end_date,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s RtEcpmQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s RtEcpmQueryRequest) GoString() string {
	return s.String()
}

func (s *RtEcpmQueryRequest) SetHeader(v map[string]*string) *RtEcpmQueryRequest {
	s.Header = v
	return s
}

func (s *RtEcpmQueryRequest) SetAccessToken(v string) *RtEcpmQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *RtEcpmQueryRequest) SetDateHour(v string) *RtEcpmQueryRequest {
	s.DateHour = &v
	return s
}

func (s *RtEcpmQueryRequest) SetCursor(v string) *RtEcpmQueryRequest {
	s.Cursor = &v
	return s
}

func (s *RtEcpmQueryRequest) SetPageSize(v int32) *RtEcpmQueryRequest {
	s.PageSize = &v
	return s
}

func (s *RtEcpmQueryRequest) SetStartDate(v string) *RtEcpmQueryRequest {
	s.StartDate = &v
	return s
}

func (s *RtEcpmQueryRequest) SetEndDate(v string) *RtEcpmQueryRequest {
	s.EndDate = &v
	return s
}

func (s *RtEcpmQueryRequest) SetOpenId(v string) *RtEcpmQueryRequest {
	s.OpenId = &v
	return s
}

type RtEcpmQueryResponse struct {
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *RtEcpmQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s RtEcpmQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RtEcpmQueryResponse) GoString() string {
	return s.String()
}

func (s *RtEcpmQueryResponse) SetErrMsg(v string) *RtEcpmQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *RtEcpmQueryResponse) SetLogId(v string) *RtEcpmQueryResponse {
	s.LogId = &v
	return s
}

func (s *RtEcpmQueryResponse) SetData(v *RtEcpmQueryResponseData) *RtEcpmQueryResponse {
	s.Data = v
	return s
}

func (s *RtEcpmQueryResponse) SetErrNo(v int32) *RtEcpmQueryResponse {
	s.ErrNo = &v
	return s
}

type RtEcpmQueryResponseData struct {
	Records    []*RtEcpmQueryResponseDataRecordsItem `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	NextCursor *string                               `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
}

func (s RtEcpmQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s RtEcpmQueryResponseData) GoString() string {
	return s.String()
}

func (s *RtEcpmQueryResponseData) SetRecords(v []*RtEcpmQueryResponseDataRecordsItem) *RtEcpmQueryResponseData {
	s.Records = v
	return s
}

func (s *RtEcpmQueryResponseData) SetNextCursor(v string) *RtEcpmQueryResponseData {
	s.NextCursor = &v
	return s
}

type RtEcpmQueryResponseDataRecordsItem struct {
	Id        *string `json:"id,omitempty" xml:"id,omitempty"`
	MpId      *string `json:"mp_id,omitempty" xml:"mp_id,omitempty"`
	Cost      *string `json:"cost,omitempty" xml:"cost,omitempty"`
	OpenId    *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	EventTime *string `json:"event_time,omitempty" xml:"event_time,omitempty"`
	AdType    *string `json:"ad_type,omitempty" xml:"ad_type,omitempty"`
	ReqId     *string `json:"req_id,omitempty" xml:"req_id,omitempty"`
}

func (s RtEcpmQueryResponseDataRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s RtEcpmQueryResponseDataRecordsItem) GoString() string {
	return s.String()
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetId(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.Id = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetMpId(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.MpId = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetCost(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.Cost = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetOpenId(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.OpenId = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetEventTime(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.EventTime = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetAdType(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.AdType = &v
	return s
}

func (s *RtEcpmQueryResponseDataRecordsItem) SetReqId(v string) *RtEcpmQueryResponseDataRecordsItem {
	s.ReqId = &v
	return s
}

type RulePushRequest struct {
	IsAutoExtension      *bool                                `json:"is_auto_extension,omitempty" xml:"is_auto_extension,omitempty" require:"true"`
	PromotionBasicInfo   *RulePushRequestPromotionBasicInfo   `json:"promotion_basic_info,omitempty" xml:"promotion_basic_info,omitempty" require:"true"`
	FullPatternPromotion *RulePushRequestFullPatternPromotion `json:"full_pattern_promotion,omitempty" xml:"full_pattern_promotion,omitempty"`
	ApplicableDate       *RulePushRequestApplicableDate       `json:"applicable_date,omitempty" xml:"applicable_date,omitempty" require:"true"`
	AccountId            *string                              `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Active               *bool                                `json:"active,omitempty" xml:"active,omitempty"`
	PromotionId          *string                              `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	ApplicableResource   *RulePushRequestApplicableResource   `json:"applicable_resource,omitempty" xml:"applicable_resource,omitempty" require:"true"`
	UnapplicableDate     *RulePushRequestUnapplicableDate     `json:"unapplicable_date,omitempty" xml:"unapplicable_date,omitempty"`
	Header               map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s RulePushRequest) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequest) GoString() string {
	return s.String()
}

func (s *RulePushRequest) SetIsAutoExtension(v bool) *RulePushRequest {
	s.IsAutoExtension = &v
	return s
}

func (s *RulePushRequest) SetPromotionBasicInfo(v *RulePushRequestPromotionBasicInfo) *RulePushRequest {
	s.PromotionBasicInfo = v
	return s
}

func (s *RulePushRequest) SetFullPatternPromotion(v *RulePushRequestFullPatternPromotion) *RulePushRequest {
	s.FullPatternPromotion = v
	return s
}

func (s *RulePushRequest) SetApplicableDate(v *RulePushRequestApplicableDate) *RulePushRequest {
	s.ApplicableDate = v
	return s
}

func (s *RulePushRequest) SetAccountId(v string) *RulePushRequest {
	s.AccountId = &v
	return s
}

func (s *RulePushRequest) SetActive(v bool) *RulePushRequest {
	s.Active = &v
	return s
}

func (s *RulePushRequest) SetPromotionId(v string) *RulePushRequest {
	s.PromotionId = &v
	return s
}

func (s *RulePushRequest) SetApplicableResource(v *RulePushRequestApplicableResource) *RulePushRequest {
	s.ApplicableResource = v
	return s
}

func (s *RulePushRequest) SetUnapplicableDate(v *RulePushRequestUnapplicableDate) *RulePushRequest {
	s.UnapplicableDate = v
	return s
}

func (s *RulePushRequest) SetHeader(v map[string]*string) *RulePushRequest {
	s.Header = v
	return s
}

func (s *RulePushRequest) SetAccessToken(v string) *RulePushRequest {
	s.AccessToken = &v
	return s
}

type RulePushRequestApplicableDate struct {
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" require:"true" type:"Repeated"`
	EndDays    *string  `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
	StartDays  *string  `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
}

func (s RulePushRequestApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestApplicableDate) GoString() string {
	return s.String()
}

func (s *RulePushRequestApplicableDate) SetDaysOfWeek(v []*int32) *RulePushRequestApplicableDate {
	s.DaysOfWeek = v
	return s
}

func (s *RulePushRequestApplicableDate) SetEndDays(v string) *RulePushRequestApplicableDate {
	s.EndDays = &v
	return s
}

func (s *RulePushRequestApplicableDate) SetStartDays(v string) *RulePushRequestApplicableDate {
	s.StartDays = &v
	return s
}

type RulePushRequestApplicableResource struct {
	Resources []*RulePushRequestApplicableResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s RulePushRequestApplicableResource) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestApplicableResource) GoString() string {
	return s.String()
}

func (s *RulePushRequestApplicableResource) SetResources(v []*RulePushRequestApplicableResourceResourcesItem) *RulePushRequestApplicableResource {
	s.Resources = v
	return s
}

type RulePushRequestApplicableResourceResourcesItem struct {
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
}

func (s RulePushRequestApplicableResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestApplicableResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *RulePushRequestApplicableResourceResourcesItem) SetRatePlanId(v string) *RulePushRequestApplicableResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

func (s *RulePushRequestApplicableResourceResourcesItem) SetRatePlanStatus(v int32) *RulePushRequestApplicableResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

type RulePushRequestFullPatternPromotion struct {
	FullPatternPromotionDetail *RulePushRequestFullPatternPromotionFullPatternPromotionDetail `json:"full_pattern_promotion_detail,omitempty" xml:"full_pattern_promotion_detail,omitempty"`
}

func (s RulePushRequestFullPatternPromotion) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestFullPatternPromotion) GoString() string {
	return s.String()
}

func (s *RulePushRequestFullPatternPromotion) SetFullPatternPromotionDetail(v *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) *RulePushRequestFullPatternPromotion {
	s.FullPatternPromotionDetail = v
	return s
}

type RulePushRequestFullPatternPromotionFullPatternPromotionDetail struct {
	DiscountType       *int     `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	EndHour            *int64   `json:"end_hour,omitempty" xml:"end_hour,omitempty"`
	IsHotelNewCustomer *bool    `json:"is_hotel_new_customer,omitempty" xml:"is_hotel_new_customer,omitempty"`
	MemberPromotion    []*int64 `json:"member_promotion,omitempty" xml:"member_promotion,omitempty" type:"Repeated"`
	StartHour          *int64   `json:"start_hour,omitempty" xml:"start_hour,omitempty"`
	AdvanceDays        *int64   `json:"advance_days,omitempty" xml:"advance_days,omitempty"`
	ConsecutiveDays    *int64   `json:"consecutive_days,omitempty" xml:"consecutive_days,omitempty"`
}

func (s RulePushRequestFullPatternPromotionFullPatternPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestFullPatternPromotionFullPatternPromotionDetail) GoString() string {
	return s.String()
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetDiscountType(v int) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.DiscountType = &v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetEndHour(v int64) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.EndHour = &v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetIsHotelNewCustomer(v bool) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.IsHotelNewCustomer = &v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetMemberPromotion(v []*int64) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.MemberPromotion = v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetStartHour(v int64) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.StartHour = &v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetAdvanceDays(v int64) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.AdvanceDays = &v
	return s
}

func (s *RulePushRequestFullPatternPromotionFullPatternPromotionDetail) SetConsecutiveDays(v int64) *RulePushRequestFullPatternPromotionFullPatternPromotionDetail {
	s.ConsecutiveDays = &v
	return s
}

type RulePushRequestPromotionBasicInfo struct {
	SubTypeCode     *string `json:"sub_type_code,omitempty" xml:"sub_type_code,omitempty"`
	SubTypeCodeName *string `json:"sub_type_code_name,omitempty" xml:"sub_type_code_name,omitempty"`
	TypeCode        *string `json:"type_code,omitempty" xml:"type_code,omitempty" require:"true"`
	TypeName        *string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	OutType         *string `json:"out_type,omitempty" xml:"out_type,omitempty"`
}

func (s RulePushRequestPromotionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestPromotionBasicInfo) GoString() string {
	return s.String()
}

func (s *RulePushRequestPromotionBasicInfo) SetSubTypeCode(v string) *RulePushRequestPromotionBasicInfo {
	s.SubTypeCode = &v
	return s
}

func (s *RulePushRequestPromotionBasicInfo) SetSubTypeCodeName(v string) *RulePushRequestPromotionBasicInfo {
	s.SubTypeCodeName = &v
	return s
}

func (s *RulePushRequestPromotionBasicInfo) SetTypeCode(v string) *RulePushRequestPromotionBasicInfo {
	s.TypeCode = &v
	return s
}

func (s *RulePushRequestPromotionBasicInfo) SetTypeName(v string) *RulePushRequestPromotionBasicInfo {
	s.TypeName = &v
	return s
}

func (s *RulePushRequestPromotionBasicInfo) SetOutType(v string) *RulePushRequestPromotionBasicInfo {
	s.OutType = &v
	return s
}

type RulePushRequestUnapplicableDate struct {
	UnapplicableDateRange []*RulePushRequestUnapplicableDateUnapplicableDateRangeItem `json:"unapplicable_date_range,omitempty" xml:"unapplicable_date_range,omitempty" type:"Repeated"`
}

func (s RulePushRequestUnapplicableDate) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestUnapplicableDate) GoString() string {
	return s.String()
}

func (s *RulePushRequestUnapplicableDate) SetUnapplicableDateRange(v []*RulePushRequestUnapplicableDateUnapplicableDateRangeItem) *RulePushRequestUnapplicableDate {
	s.UnapplicableDateRange = v
	return s
}

type RulePushRequestUnapplicableDateUnapplicableDateRangeItem struct {
	EndDays   *string `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
	StartDays *string `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
}

func (s RulePushRequestUnapplicableDateUnapplicableDateRangeItem) String() string {
	return tea.Prettify(s)
}

func (s RulePushRequestUnapplicableDateUnapplicableDateRangeItem) GoString() string {
	return s.String()
}

func (s *RulePushRequestUnapplicableDateUnapplicableDateRangeItem) SetEndDays(v string) *RulePushRequestUnapplicableDateUnapplicableDateRangeItem {
	s.EndDays = &v
	return s
}

func (s *RulePushRequestUnapplicableDateUnapplicableDateRangeItem) SetStartDays(v string) *RulePushRequestUnapplicableDateUnapplicableDateRangeItem {
	s.StartDays = &v
	return s
}

type RulePushResponse struct {
	Data  *RulePushResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *RulePushResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s RulePushResponse) String() string {
	return tea.Prettify(s)
}

func (s RulePushResponse) GoString() string {
	return s.String()
}

func (s *RulePushResponse) SetData(v *RulePushResponseData) *RulePushResponse {
	s.Data = v
	return s
}

func (s *RulePushResponse) SetExtra(v *RulePushResponseExtra) *RulePushResponse {
	s.Extra = v
	return s
}

type RulePushResponseData struct {
	Status        *int                                     `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	MessageDetail []*RulePushResponseDataMessageDetailItem `json:"message_detail,omitempty" xml:"message_detail,omitempty" type:"Repeated"`
	PromotionId   *string                                  `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s RulePushResponseData) String() string {
	return tea.Prettify(s)
}

func (s RulePushResponseData) GoString() string {
	return s.String()
}

func (s *RulePushResponseData) SetStatus(v int) *RulePushResponseData {
	s.Status = &v
	return s
}

func (s *RulePushResponseData) SetMessageDetail(v []*RulePushResponseDataMessageDetailItem) *RulePushResponseData {
	s.MessageDetail = v
	return s
}

func (s *RulePushResponseData) SetPromotionId(v string) *RulePushResponseData {
	s.PromotionId = &v
	return s
}

func (s *RulePushResponseData) SetGwErrorCode(v int32) *RulePushResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *RulePushResponseData) SetGwDescription(v string) *RulePushResponseData {
	s.GwDescription = &v
	return s
}

type RulePushResponseDataMessageDetailItem struct {
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
}

func (s RulePushResponseDataMessageDetailItem) String() string {
	return tea.Prettify(s)
}

func (s RulePushResponseDataMessageDetailItem) GoString() string {
	return s.String()
}

func (s *RulePushResponseDataMessageDetailItem) SetErrorMessage(v string) *RulePushResponseDataMessageDetailItem {
	s.ErrorMessage = &v
	return s
}

func (s *RulePushResponseDataMessageDetailItem) SetRatePlanId(v string) *RulePushResponseDataMessageDetailItem {
	s.RatePlanId = &v
	return s
}

type RulePushResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s RulePushResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s RulePushResponseExtra) GoString() string {
	return s.String()
}

func (s *RulePushResponseExtra) SetSubErrorCode(v int32) *RulePushResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *RulePushResponseExtra) SetDescription(v string) *RulePushResponseExtra {
	s.Description = &v
	return s
}

func (s *RulePushResponseExtra) SetErrorCode(v int32) *RulePushResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *RulePushResponseExtra) SetLogid(v string) *RulePushResponseExtra {
	s.Logid = &v
	return s
}

func (s *RulePushResponseExtra) SetNow(v int64) *RulePushResponseExtra {
	s.Now = &v
	return s
}

func (s *RulePushResponseExtra) SetSubDescription(v string) *RulePushResponseExtra {
	s.SubDescription = &v
	return s
}

type RuntimeOptions struct {
	Autoretry      *bool `json:"autoretry,omitempty" xml:"autoretry,omitempty" require:"true"`
	IgnoreSSL      *bool `json:"ignoreSSL,omitempty" xml:"ignoreSSL,omitempty" require:"true"`
	MaxAttempts    *int  `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty" require:"true"`
	ReadTimeout    *int  `json:"readTimeout,omitempty" xml:"readTimeout,omitempty" require:"true"`
	ConnectTimeout *int  `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty" require:"true"`
}

func (s RuntimeOptions) String() string {
	return tea.Prettify(s)
}

func (s RuntimeOptions) GoString() string {
	return s.String()
}

func (s *RuntimeOptions) SetAutoretry(v bool) *RuntimeOptions {
	s.Autoretry = &v
	return s
}

func (s *RuntimeOptions) SetIgnoreSSL(v bool) *RuntimeOptions {
	s.IgnoreSSL = &v
	return s
}

func (s *RuntimeOptions) SetMaxAttempts(v int) *RuntimeOptions {
	s.MaxAttempts = &v
	return s
}

func (s *RuntimeOptions) SetReadTimeout(v int) *RuntimeOptions {
	s.ReadTimeout = &v
	return s
}

func (s *RuntimeOptions) SetConnectTimeout(v int) *RuntimeOptions {
	s.ConnectTimeout = &v
	return s
}

type SaasAddMerchantRequest struct {
	ThirdpartyComponentId *string            `json:"thirdparty_component_id,omitempty" xml:"thirdparty_component_id,omitempty" require:"true"`
	ProdId                *int32             `json:"prod_id,omitempty" xml:"prod_id,omitempty"`
	UrlType               *int               `json:"url_type,omitempty" xml:"url_type,omitempty" require:"true"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s SaasAddMerchantRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasAddMerchantRequest) GoString() string {
	return s.String()
}

func (s *SaasAddMerchantRequest) SetThirdpartyComponentId(v string) *SaasAddMerchantRequest {
	s.ThirdpartyComponentId = &v
	return s
}

func (s *SaasAddMerchantRequest) SetProdId(v int32) *SaasAddMerchantRequest {
	s.ProdId = &v
	return s
}

func (s *SaasAddMerchantRequest) SetUrlType(v int) *SaasAddMerchantRequest {
	s.UrlType = &v
	return s
}

func (s *SaasAddMerchantRequest) SetHeader(v map[string]*string) *SaasAddMerchantRequest {
	s.Header = v
	return s
}

func (s *SaasAddMerchantRequest) SetAccessToken(v string) *SaasAddMerchantRequest {
	s.AccessToken = &v
	return s
}

type SaasAddMerchantResponse struct {
	Data   *SaasAddMerchantResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SaasAddMerchantResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasAddMerchantResponse) GoString() string {
	return s.String()
}

func (s *SaasAddMerchantResponse) SetData(v *SaasAddMerchantResponseData) *SaasAddMerchantResponse {
	s.Data = v
	return s
}

func (s *SaasAddMerchantResponse) SetErrNo(v int32) *SaasAddMerchantResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasAddMerchantResponse) SetErrMsg(v string) *SaasAddMerchantResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasAddMerchantResponse) SetLogId(v string) *SaasAddMerchantResponse {
	s.LogId = &v
	return s
}

type SaasAddMerchantResponseData struct {
	MerchantId *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s SaasAddMerchantResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasAddMerchantResponseData) GoString() string {
	return s.String()
}

func (s *SaasAddMerchantResponseData) SetMerchantId(v string) *SaasAddMerchantResponseData {
	s.MerchantId = &v
	return s
}

func (s *SaasAddMerchantResponseData) SetUrl(v string) *SaasAddMerchantResponseData {
	s.Url = &v
	return s
}

type SaasAddSubMerchantRequest struct {
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	SubMerchantId *string            `json:"sub_merchant_id,omitempty" xml:"sub_merchant_id,omitempty" require:"true"`
	UrlType       *int               `json:"url_type,omitempty" xml:"url_type,omitempty" require:"true"`
	ProdId        *int32             `json:"prod_id,omitempty" xml:"prod_id,omitempty"`
	Role          *int32             `json:"role,omitempty" xml:"role,omitempty"`
	ThirdpartyId  *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s SaasAddSubMerchantRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasAddSubMerchantRequest) GoString() string {
	return s.String()
}

func (s *SaasAddSubMerchantRequest) SetAccessToken(v string) *SaasAddSubMerchantRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetSubMerchantId(v string) *SaasAddSubMerchantRequest {
	s.SubMerchantId = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetUrlType(v int) *SaasAddSubMerchantRequest {
	s.UrlType = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetProdId(v int32) *SaasAddSubMerchantRequest {
	s.ProdId = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetRole(v int32) *SaasAddSubMerchantRequest {
	s.Role = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetThirdpartyId(v string) *SaasAddSubMerchantRequest {
	s.ThirdpartyId = &v
	return s
}

func (s *SaasAddSubMerchantRequest) SetHeader(v map[string]*string) *SaasAddSubMerchantRequest {
	s.Header = v
	return s
}

type SaasAddSubMerchantResponse struct {
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaasAddSubMerchantResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s SaasAddSubMerchantResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasAddSubMerchantResponse) GoString() string {
	return s.String()
}

func (s *SaasAddSubMerchantResponse) SetLogId(v string) *SaasAddSubMerchantResponse {
	s.LogId = &v
	return s
}

func (s *SaasAddSubMerchantResponse) SetData(v *SaasAddSubMerchantResponseData) *SaasAddSubMerchantResponse {
	s.Data = v
	return s
}

func (s *SaasAddSubMerchantResponse) SetErrNo(v int32) *SaasAddSubMerchantResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasAddSubMerchantResponse) SetErrMsg(v string) *SaasAddSubMerchantResponse {
	s.ErrMsg = &v
	return s
}

type SaasAddSubMerchantResponseData struct {
	MerchantId *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s SaasAddSubMerchantResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasAddSubMerchantResponseData) GoString() string {
	return s.String()
}

func (s *SaasAddSubMerchantResponseData) SetMerchantId(v string) *SaasAddSubMerchantResponseData {
	s.MerchantId = &v
	return s
}

func (s *SaasAddSubMerchantResponseData) SetUrl(v string) *SaasAddSubMerchantResponseData {
	s.Url = &v
	return s
}

type SaasCreateMerchantRequest struct {
	BusinessLicense       *SaasCreateMerchantRequestBusinessLicense           `json:"business_license,omitempty" xml:"business_license,omitempty" require:"true"`
	Beneficiary           *SaasCreateMerchantRequestBeneficiary               `json:"beneficiary,omitempty" xml:"beneficiary,omitempty"`
	CallbackUrl           *string                                             `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	CreateName            *string                                             `json:"create_name,omitempty" xml:"create_name,omitempty" require:"true"`
	MerchantName          *string                                             `json:"merchant_name,omitempty" xml:"merchant_name,omitempty" require:"true"`
	IndustryCode          []*string                                           `json:"industry_code,omitempty" xml:"industry_code,omitempty" require:"true" type:"Repeated"`
	AccessToken           *string                                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ThirdpartyId          *string                                             `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty"`
	MerchantCardInfo      *SaasCreateMerchantRequestMerchantCardInfo          `json:"merchant_card_info,omitempty" xml:"merchant_card_info,omitempty" require:"true"`
	Type                  *int64                                              `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Header                map[string]*string                                  `json:"header,omitempty" xml:"header,omitempty"`
	IndustryInfoPicUrls   []*SaasCreateMerchantRequestIndustryInfoPicUrlsItem `json:"industry_info_pic_urls,omitempty" xml:"industry_info_pic_urls,omitempty" type:"Repeated"`
	SubMerchantId         *string                                             `json:"sub_merchant_id,omitempty" xml:"sub_merchant_id,omitempty"`
	ExtEvidences          []*SaasCreateMerchantRequestExtEvidencesItem        `json:"ext_evidences,omitempty" xml:"ext_evidences,omitempty" type:"Repeated"`
	LegalPerson           *SaasCreateMerchantRequestLegalPerson               `json:"legal_person,omitempty" xml:"legal_person,omitempty" require:"true"`
	DistrictCode          *string                                             `json:"district_code,omitempty" xml:"district_code,omitempty" require:"true"`
	Channels              []*string                                           `json:"channels,omitempty" xml:"channels,omitempty" require:"true" type:"Repeated"`
	OutOrderId            *string                                             `json:"out_order_id,omitempty" xml:"out_order_id,omitempty" require:"true"`
	CityCode              *string                                             `json:"city_code,omitempty" xml:"city_code,omitempty" require:"true"`
	AppId                 *string                                             `json:"app_id,omitempty" xml:"app_id,omitempty"`
	BeneficiaryType       *string                                             `json:"beneficiary_type,omitempty" xml:"beneficiary_type,omitempty"`
	ProvinceCode          *string                                             `json:"province_code,omitempty" xml:"province_code,omitempty" require:"true"`
	MerchantOperationInfo *SaasCreateMerchantRequestMerchantOperationInfo     `json:"merchant_operation_info,omitempty" xml:"merchant_operation_info,omitempty" require:"true"`
	RegisteredAddr        *string                                             `json:"registered_addr,omitempty" xml:"registered_addr,omitempty" require:"true"`
	MerchantType          *int64                                              `json:"merchant_type,omitempty" xml:"merchant_type,omitempty" require:"true"`
	MerchantShortName     *string                                             `json:"merchant_short_name,omitempty" xml:"merchant_short_name,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequest) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequest) SetBusinessLicense(v *SaasCreateMerchantRequestBusinessLicense) *SaasCreateMerchantRequest {
	s.BusinessLicense = v
	return s
}

func (s *SaasCreateMerchantRequest) SetBeneficiary(v *SaasCreateMerchantRequestBeneficiary) *SaasCreateMerchantRequest {
	s.Beneficiary = v
	return s
}

func (s *SaasCreateMerchantRequest) SetCallbackUrl(v string) *SaasCreateMerchantRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetCreateName(v string) *SaasCreateMerchantRequest {
	s.CreateName = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetMerchantName(v string) *SaasCreateMerchantRequest {
	s.MerchantName = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetIndustryCode(v []*string) *SaasCreateMerchantRequest {
	s.IndustryCode = v
	return s
}

func (s *SaasCreateMerchantRequest) SetAccessToken(v string) *SaasCreateMerchantRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetThirdpartyId(v string) *SaasCreateMerchantRequest {
	s.ThirdpartyId = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetMerchantCardInfo(v *SaasCreateMerchantRequestMerchantCardInfo) *SaasCreateMerchantRequest {
	s.MerchantCardInfo = v
	return s
}

func (s *SaasCreateMerchantRequest) SetType(v int64) *SaasCreateMerchantRequest {
	s.Type = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetHeader(v map[string]*string) *SaasCreateMerchantRequest {
	s.Header = v
	return s
}

func (s *SaasCreateMerchantRequest) SetIndustryInfoPicUrls(v []*SaasCreateMerchantRequestIndustryInfoPicUrlsItem) *SaasCreateMerchantRequest {
	s.IndustryInfoPicUrls = v
	return s
}

func (s *SaasCreateMerchantRequest) SetSubMerchantId(v string) *SaasCreateMerchantRequest {
	s.SubMerchantId = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetExtEvidences(v []*SaasCreateMerchantRequestExtEvidencesItem) *SaasCreateMerchantRequest {
	s.ExtEvidences = v
	return s
}

func (s *SaasCreateMerchantRequest) SetLegalPerson(v *SaasCreateMerchantRequestLegalPerson) *SaasCreateMerchantRequest {
	s.LegalPerson = v
	return s
}

func (s *SaasCreateMerchantRequest) SetDistrictCode(v string) *SaasCreateMerchantRequest {
	s.DistrictCode = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetChannels(v []*string) *SaasCreateMerchantRequest {
	s.Channels = v
	return s
}

func (s *SaasCreateMerchantRequest) SetOutOrderId(v string) *SaasCreateMerchantRequest {
	s.OutOrderId = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetCityCode(v string) *SaasCreateMerchantRequest {
	s.CityCode = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetAppId(v string) *SaasCreateMerchantRequest {
	s.AppId = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetBeneficiaryType(v string) *SaasCreateMerchantRequest {
	s.BeneficiaryType = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetProvinceCode(v string) *SaasCreateMerchantRequest {
	s.ProvinceCode = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetMerchantOperationInfo(v *SaasCreateMerchantRequestMerchantOperationInfo) *SaasCreateMerchantRequest {
	s.MerchantOperationInfo = v
	return s
}

func (s *SaasCreateMerchantRequest) SetRegisteredAddr(v string) *SaasCreateMerchantRequest {
	s.RegisteredAddr = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetMerchantType(v int64) *SaasCreateMerchantRequest {
	s.MerchantType = &v
	return s
}

func (s *SaasCreateMerchantRequest) SetMerchantShortName(v string) *SaasCreateMerchantRequest {
	s.MerchantShortName = &v
	return s
}

type SaasCreateMerchantRequestBeneficiary struct {
	BeginDate   *string `json:"begin_date,omitempty" xml:"begin_date,omitempty" require:"true"`
	ExpDate     *string `json:"exp_date,omitempty" xml:"exp_date,omitempty" require:"true"`
	Address     *string `json:"address,omitempty" xml:"address,omitempty"`
	IdType      *int32  `json:"id_type,omitempty" xml:"id_type,omitempty" require:"true"`
	IdNo        *string `json:"id_no,omitempty" xml:"id_no,omitempty" require:"true"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	FrontPicUrl *string `json:"front_pic_url,omitempty" xml:"front_pic_url,omitempty" require:"true"`
	BackPicUrl  *string `json:"back_pic_url,omitempty" xml:"back_pic_url,omitempty"`
}

func (s SaasCreateMerchantRequestBeneficiary) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestBeneficiary) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestBeneficiary) SetBeginDate(v string) *SaasCreateMerchantRequestBeneficiary {
	s.BeginDate = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetExpDate(v string) *SaasCreateMerchantRequestBeneficiary {
	s.ExpDate = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetAddress(v string) *SaasCreateMerchantRequestBeneficiary {
	s.Address = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetIdType(v int32) *SaasCreateMerchantRequestBeneficiary {
	s.IdType = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetIdNo(v string) *SaasCreateMerchantRequestBeneficiary {
	s.IdNo = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetName(v string) *SaasCreateMerchantRequestBeneficiary {
	s.Name = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetFrontPicUrl(v string) *SaasCreateMerchantRequestBeneficiary {
	s.FrontPicUrl = &v
	return s
}

func (s *SaasCreateMerchantRequestBeneficiary) SetBackPicUrl(v string) *SaasCreateMerchantRequestBeneficiary {
	s.BackPicUrl = &v
	return s
}

type SaasCreateMerchantRequestBusinessLicense struct {
	EndDate                   *string                                                                  `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	BusinessLicenseType       *int32                                                                   `json:"business_license_type,omitempty" xml:"business_license_type,omitempty" require:"true"`
	BusinessLicenseCode       *string                                                                  `json:"business_license_code,omitempty" xml:"business_license_code,omitempty" require:"true"`
	BusinessLicensePicurl     []*SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem     `json:"business_license_picurl,omitempty" xml:"business_license_picurl,omitempty" require:"true" type:"Repeated"`
	BusinessLicenseBackPicurl []*SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem `json:"business_license_back_picurl,omitempty" xml:"business_license_back_picurl,omitempty" type:"Repeated"`
	BeginDate                 *string                                                                  `json:"begin_date,omitempty" xml:"begin_date,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestBusinessLicense) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestBusinessLicense) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetEndDate(v string) *SaasCreateMerchantRequestBusinessLicense {
	s.EndDate = &v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetBusinessLicenseType(v int32) *SaasCreateMerchantRequestBusinessLicense {
	s.BusinessLicenseType = &v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetBusinessLicenseCode(v string) *SaasCreateMerchantRequestBusinessLicense {
	s.BusinessLicenseCode = &v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetBusinessLicensePicurl(v []*SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem) *SaasCreateMerchantRequestBusinessLicense {
	s.BusinessLicensePicurl = v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetBusinessLicenseBackPicurl(v []*SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem) *SaasCreateMerchantRequestBusinessLicense {
	s.BusinessLicenseBackPicurl = v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicense) SetBeginDate(v string) *SaasCreateMerchantRequestBusinessLicense {
	s.BeginDate = &v
	return s
}

type SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem struct {
	Url     *string `json:"url,omitempty" xml:"url,omitempty"`
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
}

func (s SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem) SetUrl(v string) *SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem {
	s.Url = &v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem) SetChannel(v string) *SaasCreateMerchantRequestBusinessLicenseBusinessLicenseBackPicurlItem {
	s.Channel = &v
	return s
}

type SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
	Url     *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem) SetChannel(v string) *SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem {
	s.Channel = &v
	return s
}

func (s *SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem) SetUrl(v string) *SaasCreateMerchantRequestBusinessLicenseBusinessLicensePicurlItem {
	s.Url = &v
	return s
}

type SaasCreateMerchantRequestExtEvidencesItem struct {
	Url     *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestExtEvidencesItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestExtEvidencesItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestExtEvidencesItem) SetUrl(v string) *SaasCreateMerchantRequestExtEvidencesItem {
	s.Url = &v
	return s
}

func (s *SaasCreateMerchantRequestExtEvidencesItem) SetChannel(v string) *SaasCreateMerchantRequestExtEvidencesItem {
	s.Channel = &v
	return s
}

type SaasCreateMerchantRequestIndustryInfoPicUrlsItem struct {
	Url     *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestIndustryInfoPicUrlsItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestIndustryInfoPicUrlsItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestIndustryInfoPicUrlsItem) SetUrl(v string) *SaasCreateMerchantRequestIndustryInfoPicUrlsItem {
	s.Url = &v
	return s
}

func (s *SaasCreateMerchantRequestIndustryInfoPicUrlsItem) SetChannel(v string) *SaasCreateMerchantRequestIndustryInfoPicUrlsItem {
	s.Channel = &v
	return s
}

type SaasCreateMerchantRequestLegalPerson struct {
	FrontPicUrl []*SaasCreateMerchantRequestLegalPersonFrontPicUrlItem `json:"front_pic_url,omitempty" xml:"front_pic_url,omitempty" require:"true" type:"Repeated"`
	BackPicUrl  []*SaasCreateMerchantRequestLegalPersonBackPicUrlItem  `json:"back_pic_url,omitempty" xml:"back_pic_url,omitempty" type:"Repeated"`
	BeginDate   *string                                                `json:"begin_date,omitempty" xml:"begin_date,omitempty" require:"true"`
	ExpDate     *string                                                `json:"exp_date,omitempty" xml:"exp_date,omitempty" require:"true"`
	Address     *string                                                `json:"address,omitempty" xml:"address,omitempty"`
	IdType      *int32                                                 `json:"id_type,omitempty" xml:"id_type,omitempty" require:"true"`
	IdNo        *string                                                `json:"id_no,omitempty" xml:"id_no,omitempty" require:"true"`
	Name        *string                                                `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestLegalPerson) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestLegalPerson) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestLegalPerson) SetFrontPicUrl(v []*SaasCreateMerchantRequestLegalPersonFrontPicUrlItem) *SaasCreateMerchantRequestLegalPerson {
	s.FrontPicUrl = v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetBackPicUrl(v []*SaasCreateMerchantRequestLegalPersonBackPicUrlItem) *SaasCreateMerchantRequestLegalPerson {
	s.BackPicUrl = v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetBeginDate(v string) *SaasCreateMerchantRequestLegalPerson {
	s.BeginDate = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetExpDate(v string) *SaasCreateMerchantRequestLegalPerson {
	s.ExpDate = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetAddress(v string) *SaasCreateMerchantRequestLegalPerson {
	s.Address = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetIdType(v int32) *SaasCreateMerchantRequestLegalPerson {
	s.IdType = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetIdNo(v string) *SaasCreateMerchantRequestLegalPerson {
	s.IdNo = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPerson) SetName(v string) *SaasCreateMerchantRequestLegalPerson {
	s.Name = &v
	return s
}

type SaasCreateMerchantRequestLegalPersonBackPicUrlItem struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
	Url     *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestLegalPersonBackPicUrlItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestLegalPersonBackPicUrlItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestLegalPersonBackPicUrlItem) SetChannel(v string) *SaasCreateMerchantRequestLegalPersonBackPicUrlItem {
	s.Channel = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPersonBackPicUrlItem) SetUrl(v string) *SaasCreateMerchantRequestLegalPersonBackPicUrlItem {
	s.Url = &v
	return s
}

type SaasCreateMerchantRequestLegalPersonFrontPicUrlItem struct {
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
	Url     *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestLegalPersonFrontPicUrlItem) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestLegalPersonFrontPicUrlItem) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestLegalPersonFrontPicUrlItem) SetChannel(v string) *SaasCreateMerchantRequestLegalPersonFrontPicUrlItem {
	s.Channel = &v
	return s
}

func (s *SaasCreateMerchantRequestLegalPersonFrontPicUrlItem) SetUrl(v string) *SaasCreateMerchantRequestLegalPersonFrontPicUrlItem {
	s.Url = &v
	return s
}

type SaasCreateMerchantRequestMerchantCardInfo struct {
	AlipaySettleType *int32  `json:"alipay_settle_type,omitempty" xml:"alipay_settle_type,omitempty" require:"true"`
	AccountNo        *string `json:"account_no,omitempty" xml:"account_no,omitempty" require:"true"`
	AlipayAccountNo  *string `json:"alipay_account_no,omitempty" xml:"alipay_account_no,omitempty"`
	CardType         *string `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
	SettleType       *int32  `json:"settle_type,omitempty" xml:"settle_type,omitempty" require:"true"`
	AccountName      *string `json:"account_name,omitempty" xml:"account_name,omitempty" require:"true"`
	BankFullName     *string `json:"bank_full_name,omitempty" xml:"bank_full_name,omitempty" require:"true"`
}

func (s SaasCreateMerchantRequestMerchantCardInfo) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestMerchantCardInfo) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetAlipaySettleType(v int32) *SaasCreateMerchantRequestMerchantCardInfo {
	s.AlipaySettleType = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetAccountNo(v string) *SaasCreateMerchantRequestMerchantCardInfo {
	s.AccountNo = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetAlipayAccountNo(v string) *SaasCreateMerchantRequestMerchantCardInfo {
	s.AlipayAccountNo = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetCardType(v string) *SaasCreateMerchantRequestMerchantCardInfo {
	s.CardType = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetSettleType(v int32) *SaasCreateMerchantRequestMerchantCardInfo {
	s.SettleType = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetAccountName(v string) *SaasCreateMerchantRequestMerchantCardInfo {
	s.AccountName = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantCardInfo) SetBankFullName(v string) *SaasCreateMerchantRequestMerchantCardInfo {
	s.BankFullName = &v
	return s
}

type SaasCreateMerchantRequestMerchantOperationInfo struct {
	ManageEmail                 *string `json:"manage_email,omitempty" xml:"manage_email,omitempty" require:"true"`
	AlipayAccountNo             *string `json:"alipay_account_no,omitempty" xml:"alipay_account_no,omitempty"`
	ShopName                    *string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	IdBeginDate                 *string `json:"id_begin_date,omitempty" xml:"id_begin_date,omitempty"`
	IdFrontPicUrl               *string `json:"id_front_pic_url,omitempty" xml:"id_front_pic_url,omitempty"`
	IdExpDate                   *string `json:"id_exp_date,omitempty" xml:"id_exp_date,omitempty"`
	ManageName                  *string `json:"manage_name,omitempty" xml:"manage_name,omitempty" require:"true"`
	IdType                      *int32  `json:"id_type,omitempty" xml:"id_type,omitempty"`
	ShopUrl                     *string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	BusinessAuthorizationLetter *string `json:"business_authorization_letter,omitempty" xml:"business_authorization_letter,omitempty"`
	ManageIdNo                  *string `json:"manage_id_no,omitempty" xml:"manage_id_no,omitempty" require:"true"`
	ManagePersonType            *int32  `json:"manage_person_type,omitempty" xml:"manage_person_type,omitempty" require:"true"`
	ManageMobile                *string `json:"manage_mobile,omitempty" xml:"manage_mobile,omitempty" require:"true"`
	IdBackPicUrl                *string `json:"id_back_pic_url,omitempty" xml:"id_back_pic_url,omitempty"`
}

func (s SaasCreateMerchantRequestMerchantOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantRequestMerchantOperationInfo) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetManageEmail(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ManageEmail = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetAlipayAccountNo(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.AlipayAccountNo = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetShopName(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ShopName = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetIdBeginDate(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.IdBeginDate = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetIdFrontPicUrl(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.IdFrontPicUrl = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetIdExpDate(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.IdExpDate = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetManageName(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ManageName = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetIdType(v int32) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.IdType = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetShopUrl(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ShopUrl = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetBusinessAuthorizationLetter(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.BusinessAuthorizationLetter = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetManageIdNo(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ManageIdNo = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetManagePersonType(v int32) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ManagePersonType = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetManageMobile(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.ManageMobile = &v
	return s
}

func (s *SaasCreateMerchantRequestMerchantOperationInfo) SetIdBackPicUrl(v string) *SaasCreateMerchantRequestMerchantOperationInfo {
	s.IdBackPicUrl = &v
	return s
}

type SaasCreateMerchantResponse struct {
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaasCreateMerchantResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SaasCreateMerchantResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantResponse) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantResponse) SetErrMsg(v string) *SaasCreateMerchantResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasCreateMerchantResponse) SetLogId(v string) *SaasCreateMerchantResponse {
	s.LogId = &v
	return s
}

func (s *SaasCreateMerchantResponse) SetData(v *SaasCreateMerchantResponseData) *SaasCreateMerchantResponse {
	s.Data = v
	return s
}

func (s *SaasCreateMerchantResponse) SetErrNo(v int32) *SaasCreateMerchantResponse {
	s.ErrNo = &v
	return s
}

type SaasCreateMerchantResponseData struct {
	ApplyId    *string `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	MerchantId *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
}

func (s SaasCreateMerchantResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasCreateMerchantResponseData) GoString() string {
	return s.String()
}

func (s *SaasCreateMerchantResponseData) SetApplyId(v string) *SaasCreateMerchantResponseData {
	s.ApplyId = &v
	return s
}

func (s *SaasCreateMerchantResponseData) SetMerchantId(v string) *SaasCreateMerchantResponseData {
	s.MerchantId = &v
	return s
}

type SaasGetAppMerchantRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ThirdpartyId *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty" require:"true"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	UrlType      *int               `json:"url_type,omitempty" xml:"url_type,omitempty" require:"true"`
}

func (s SaasGetAppMerchantRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasGetAppMerchantRequest) GoString() string {
	return s.String()
}

func (s *SaasGetAppMerchantRequest) SetHeader(v map[string]*string) *SaasGetAppMerchantRequest {
	s.Header = v
	return s
}

func (s *SaasGetAppMerchantRequest) SetAccessToken(v string) *SaasGetAppMerchantRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasGetAppMerchantRequest) SetThirdpartyId(v string) *SaasGetAppMerchantRequest {
	s.ThirdpartyId = &v
	return s
}

func (s *SaasGetAppMerchantRequest) SetAppId(v string) *SaasGetAppMerchantRequest {
	s.AppId = &v
	return s
}

func (s *SaasGetAppMerchantRequest) SetUrlType(v int) *SaasGetAppMerchantRequest {
	s.UrlType = &v
	return s
}

type SaasGetAppMerchantResponse struct {
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaasGetAppMerchantResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s SaasGetAppMerchantResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasGetAppMerchantResponse) GoString() string {
	return s.String()
}

func (s *SaasGetAppMerchantResponse) SetErrMsg(v string) *SaasGetAppMerchantResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasGetAppMerchantResponse) SetLogId(v string) *SaasGetAppMerchantResponse {
	s.LogId = &v
	return s
}

func (s *SaasGetAppMerchantResponse) SetData(v *SaasGetAppMerchantResponseData) *SaasGetAppMerchantResponse {
	s.Data = v
	return s
}

func (s *SaasGetAppMerchantResponse) SetErrNo(v int32) *SaasGetAppMerchantResponse {
	s.ErrNo = &v
	return s
}

type SaasGetAppMerchantResponseData struct {
	Url        *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	MerchantId *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
}

func (s SaasGetAppMerchantResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasGetAppMerchantResponseData) GoString() string {
	return s.String()
}

func (s *SaasGetAppMerchantResponseData) SetUrl(v string) *SaasGetAppMerchantResponseData {
	s.Url = &v
	return s
}

func (s *SaasGetAppMerchantResponseData) SetMerchantId(v string) *SaasGetAppMerchantResponseData {
	s.MerchantId = &v
	return s
}

type SaasMerchantWithdrawRequest struct {
	OutOrderId     *string            `json:"out_order_id,omitempty" xml:"out_order_id,omitempty" require:"true"`
	ThirdpartyId   *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty"`
	Callback       *string            `json:"callback,omitempty" xml:"callback,omitempty"`
	WithdrawAmount *int64             `json:"withdraw_amount,omitempty" xml:"withdraw_amount,omitempty" require:"true"`
	CpExtra        *string            `json:"cp_extra,omitempty" xml:"cp_extra,omitempty"`
	MerchantEntity *int32             `json:"merchant_entity,omitempty" xml:"merchant_entity,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId          *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	ChannelType    *string            `json:"channel_type,omitempty" xml:"channel_type,omitempty" require:"true"`
	MerchantUid    *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty" require:"true"`
}

func (s SaasMerchantWithdrawRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasMerchantWithdrawRequest) GoString() string {
	return s.String()
}

func (s *SaasMerchantWithdrawRequest) SetOutOrderId(v string) *SaasMerchantWithdrawRequest {
	s.OutOrderId = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetThirdpartyId(v string) *SaasMerchantWithdrawRequest {
	s.ThirdpartyId = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetCallback(v string) *SaasMerchantWithdrawRequest {
	s.Callback = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetWithdrawAmount(v int64) *SaasMerchantWithdrawRequest {
	s.WithdrawAmount = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetCpExtra(v string) *SaasMerchantWithdrawRequest {
	s.CpExtra = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetMerchantEntity(v int32) *SaasMerchantWithdrawRequest {
	s.MerchantEntity = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetHeader(v map[string]*string) *SaasMerchantWithdrawRequest {
	s.Header = v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetAccessToken(v string) *SaasMerchantWithdrawRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetAppId(v string) *SaasMerchantWithdrawRequest {
	s.AppId = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetChannelType(v string) *SaasMerchantWithdrawRequest {
	s.ChannelType = &v
	return s
}

func (s *SaasMerchantWithdrawRequest) SetMerchantUid(v string) *SaasMerchantWithdrawRequest {
	s.MerchantUid = &v
	return s
}

type SaasMerchantWithdrawResponse struct {
	Data   *SaasMerchantWithdrawResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SaasMerchantWithdrawResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasMerchantWithdrawResponse) GoString() string {
	return s.String()
}

func (s *SaasMerchantWithdrawResponse) SetData(v *SaasMerchantWithdrawResponseData) *SaasMerchantWithdrawResponse {
	s.Data = v
	return s
}

func (s *SaasMerchantWithdrawResponse) SetErrNo(v int32) *SaasMerchantWithdrawResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasMerchantWithdrawResponse) SetErrMsg(v string) *SaasMerchantWithdrawResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasMerchantWithdrawResponse) SetLogId(v string) *SaasMerchantWithdrawResponse {
	s.LogId = &v
	return s
}

type SaasMerchantWithdrawResponseData struct {
	MerchantEntity *int32  `json:"merchant_entity,omitempty" xml:"merchant_entity,omitempty" require:"true"`
	OrderId        *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s SaasMerchantWithdrawResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasMerchantWithdrawResponseData) GoString() string {
	return s.String()
}

func (s *SaasMerchantWithdrawResponseData) SetMerchantEntity(v int32) *SaasMerchantWithdrawResponseData {
	s.MerchantEntity = &v
	return s
}

func (s *SaasMerchantWithdrawResponseData) SetOrderId(v string) *SaasMerchantWithdrawResponseData {
	s.OrderId = &v
	return s
}

type SaasQueryMerchantBalanceRequest struct {
	MerchantUid    *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty" require:"true"`
	ChannelType    *string            `json:"channel_type,omitempty" xml:"channel_type,omitempty" require:"true"`
	MerchantEntity *int32             `json:"merchant_entity,omitempty" xml:"merchant_entity,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId          *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	ThirdpartyId   *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty"`
}

func (s SaasQueryMerchantBalanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantBalanceRequest) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantBalanceRequest) SetMerchantUid(v string) *SaasQueryMerchantBalanceRequest {
	s.MerchantUid = &v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetChannelType(v string) *SaasQueryMerchantBalanceRequest {
	s.ChannelType = &v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetMerchantEntity(v int32) *SaasQueryMerchantBalanceRequest {
	s.MerchantEntity = &v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetHeader(v map[string]*string) *SaasQueryMerchantBalanceRequest {
	s.Header = v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetAccessToken(v string) *SaasQueryMerchantBalanceRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetAppId(v string) *SaasQueryMerchantBalanceRequest {
	s.AppId = &v
	return s
}

func (s *SaasQueryMerchantBalanceRequest) SetThirdpartyId(v string) *SaasQueryMerchantBalanceRequest {
	s.ThirdpartyId = &v
	return s
}

type SaasQueryMerchantBalanceResponse struct {
	Data   *SaasQueryMerchantBalanceResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s SaasQueryMerchantBalanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantBalanceResponse) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantBalanceResponse) SetData(v *SaasQueryMerchantBalanceResponseData) *SaasQueryMerchantBalanceResponse {
	s.Data = v
	return s
}

func (s *SaasQueryMerchantBalanceResponse) SetErrNo(v int32) *SaasQueryMerchantBalanceResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponse) SetErrMsg(v string) *SaasQueryMerchantBalanceResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponse) SetLogId(v string) *SaasQueryMerchantBalanceResponse {
	s.LogId = &v
	return s
}

type SaasQueryMerchantBalanceResponseData struct {
	MerchantEntity *int32                                           `json:"merchant_entity,omitempty" xml:"merchant_entity,omitempty" require:"true"`
	AccountInfo    *SaasQueryMerchantBalanceResponseDataAccountInfo `json:"account_info,omitempty" xml:"account_info,omitempty" require:"true"`
	SettleInfo     *SaasQueryMerchantBalanceResponseDataSettleInfo  `json:"settle_info,omitempty" xml:"settle_info,omitempty" require:"true"`
}

func (s SaasQueryMerchantBalanceResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantBalanceResponseData) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantBalanceResponseData) SetMerchantEntity(v int32) *SaasQueryMerchantBalanceResponseData {
	s.MerchantEntity = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseData) SetAccountInfo(v *SaasQueryMerchantBalanceResponseDataAccountInfo) *SaasQueryMerchantBalanceResponseData {
	s.AccountInfo = v
	return s
}

func (s *SaasQueryMerchantBalanceResponseData) SetSettleInfo(v *SaasQueryMerchantBalanceResponseDataSettleInfo) *SaasQueryMerchantBalanceResponseData {
	s.SettleInfo = v
	return s
}

type SaasQueryMerchantBalanceResponseDataAccountInfo struct {
	OnlineBalance       *int64 `json:"online_balance,omitempty" xml:"online_balance,omitempty" require:"true"`
	WithdrawableBalacne *int64 `json:"withdrawable_balacne,omitempty" xml:"withdrawable_balacne,omitempty" require:"true"`
	FreezeBalance       *int64 `json:"freeze_balance,omitempty" xml:"freeze_balance,omitempty" require:"true"`
}

func (s SaasQueryMerchantBalanceResponseDataAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantBalanceResponseDataAccountInfo) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantBalanceResponseDataAccountInfo) SetOnlineBalance(v int64) *SaasQueryMerchantBalanceResponseDataAccountInfo {
	s.OnlineBalance = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseDataAccountInfo) SetWithdrawableBalacne(v int64) *SaasQueryMerchantBalanceResponseDataAccountInfo {
	s.WithdrawableBalacne = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseDataAccountInfo) SetFreezeBalance(v int64) *SaasQueryMerchantBalanceResponseDataAccountInfo {
	s.FreezeBalance = &v
	return s
}

type SaasQueryMerchantBalanceResponseDataSettleInfo struct {
	BankName      *string `json:"bank_name,omitempty" xml:"bank_name,omitempty" require:"true"`
	SettleType    *int32  `json:"settle_type,omitempty" xml:"settle_type,omitempty" require:"true"`
	SettleAccount *string `json:"settle_account,omitempty" xml:"settle_account,omitempty" require:"true"`
	BankcardNo    *string `json:"bankcard_no,omitempty" xml:"bankcard_no,omitempty" require:"true"`
}

func (s SaasQueryMerchantBalanceResponseDataSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantBalanceResponseDataSettleInfo) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantBalanceResponseDataSettleInfo) SetBankName(v string) *SaasQueryMerchantBalanceResponseDataSettleInfo {
	s.BankName = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseDataSettleInfo) SetSettleType(v int32) *SaasQueryMerchantBalanceResponseDataSettleInfo {
	s.SettleType = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseDataSettleInfo) SetSettleAccount(v string) *SaasQueryMerchantBalanceResponseDataSettleInfo {
	s.SettleAccount = &v
	return s
}

func (s *SaasQueryMerchantBalanceResponseDataSettleInfo) SetBankcardNo(v string) *SaasQueryMerchantBalanceResponseDataSettleInfo {
	s.BankcardNo = &v
	return s
}

type SaasQueryMerchantStatusRequest struct {
	SubMerchantId *string            `json:"sub_merchant_id,omitempty" xml:"sub_merchant_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	ThirdpartyId  *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty"`
	MerchantId    *string            `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
}

func (s SaasQueryMerchantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusRequest) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusRequest) SetSubMerchantId(v string) *SaasQueryMerchantStatusRequest {
	s.SubMerchantId = &v
	return s
}

func (s *SaasQueryMerchantStatusRequest) SetHeader(v map[string]*string) *SaasQueryMerchantStatusRequest {
	s.Header = v
	return s
}

func (s *SaasQueryMerchantStatusRequest) SetAccessToken(v string) *SaasQueryMerchantStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasQueryMerchantStatusRequest) SetAppId(v string) *SaasQueryMerchantStatusRequest {
	s.AppId = &v
	return s
}

func (s *SaasQueryMerchantStatusRequest) SetThirdpartyId(v string) *SaasQueryMerchantStatusRequest {
	s.ThirdpartyId = &v
	return s
}

func (s *SaasQueryMerchantStatusRequest) SetMerchantId(v string) *SaasQueryMerchantStatusRequest {
	s.MerchantId = &v
	return s
}

type SaasQueryMerchantStatusResponse struct {
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *SaasQueryMerchantStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SaasQueryMerchantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusResponse) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusResponse) SetErrNo(v int32) *SaasQueryMerchantStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *SaasQueryMerchantStatusResponse) SetErrMsg(v string) *SaasQueryMerchantStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *SaasQueryMerchantStatusResponse) SetLogId(v string) *SaasQueryMerchantStatusResponse {
	s.LogId = &v
	return s
}

func (s *SaasQueryMerchantStatusResponse) SetData(v *SaasQueryMerchantStatusResponseData) *SaasQueryMerchantStatusResponse {
	s.Data = v
	return s
}

type SaasQueryMerchantStatusResponseData struct {
	PayStatusInfo         map[string]*SaasQueryMerchantStatusResponseDataPayStatusInfoValue         `json:"pay_status_info,omitempty" xml:"pay_status_info,omitempty"`
	MerchantStatusInfo    map[string]*SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue    `json:"merchant_status_info,omitempty" xml:"merchant_status_info,omitempty" require:"true"`
	NewMerchantStatusInfo map[string]*SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue `json:"new_merchant_status_info,omitempty" xml:"new_merchant_status_info,omitempty" require:"true"`
}

func (s SaasQueryMerchantStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusResponseData) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusResponseData) SetPayStatusInfo(v map[string]*SaasQueryMerchantStatusResponseDataPayStatusInfoValue) *SaasQueryMerchantStatusResponseData {
	s.PayStatusInfo = v
	return s
}

func (s *SaasQueryMerchantStatusResponseData) SetMerchantStatusInfo(v map[string]*SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) *SaasQueryMerchantStatusResponseData {
	s.MerchantStatusInfo = v
	return s
}

func (s *SaasQueryMerchantStatusResponseData) SetNewMerchantStatusInfo(v map[string]*SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) *SaasQueryMerchantStatusResponseData {
	s.NewMerchantStatusInfo = v
	return s
}

type SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue struct {
	RejectReason       *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty" require:"true"`
	LegalValidationUrl *string `json:"legal_validation_url,omitempty" xml:"legal_validation_url,omitempty" require:"true"`
	LegalSignUrl       *string `json:"legal_sign_url,omitempty" xml:"legal_sign_url,omitempty" require:"true"`
	MerchantId         *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) SetRejectReason(v string) *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue {
	s.RejectReason = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) SetLegalValidationUrl(v string) *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue {
	s.LegalValidationUrl = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) SetLegalSignUrl(v string) *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue {
	s.LegalSignUrl = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) SetMerchantId(v string) *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue {
	s.MerchantId = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue) SetStatus(v string) *SaasQueryMerchantStatusResponseDataMerchantStatusInfoValue {
	s.Status = &v
	return s
}

type SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue struct {
	Status             *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	RejectReason       *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty" require:"true"`
	LegalValidationUrl *string `json:"legal_validation_url,omitempty" xml:"legal_validation_url,omitempty" require:"true"`
	LegalSignUrl       *string `json:"legal_sign_url,omitempty" xml:"legal_sign_url,omitempty" require:"true"`
	MerchantId         *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
}

func (s SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) SetStatus(v string) *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue {
	s.Status = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) SetRejectReason(v string) *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue {
	s.RejectReason = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) SetLegalValidationUrl(v string) *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue {
	s.LegalValidationUrl = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) SetLegalSignUrl(v string) *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue {
	s.LegalSignUrl = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue) SetMerchantId(v string) *SaasQueryMerchantStatusResponseDataNewMerchantStatusInfoValue {
	s.MerchantId = &v
	return s
}

type SaasQueryMerchantStatusResponseDataPayStatusInfoValue struct {
	LegalSignUrl       *string `json:"legal_sign_url,omitempty" xml:"legal_sign_url,omitempty" require:"true"`
	MerchantId         *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	RejectReason       *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty" require:"true"`
	LegalValidationUrl *string `json:"legal_validation_url,omitempty" xml:"legal_validation_url,omitempty" require:"true"`
}

func (s SaasQueryMerchantStatusResponseDataPayStatusInfoValue) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryMerchantStatusResponseDataPayStatusInfoValue) GoString() string {
	return s.String()
}

func (s *SaasQueryMerchantStatusResponseDataPayStatusInfoValue) SetLegalSignUrl(v string) *SaasQueryMerchantStatusResponseDataPayStatusInfoValue {
	s.LegalSignUrl = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataPayStatusInfoValue) SetMerchantId(v string) *SaasQueryMerchantStatusResponseDataPayStatusInfoValue {
	s.MerchantId = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataPayStatusInfoValue) SetStatus(v string) *SaasQueryMerchantStatusResponseDataPayStatusInfoValue {
	s.Status = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataPayStatusInfoValue) SetRejectReason(v string) *SaasQueryMerchantStatusResponseDataPayStatusInfoValue {
	s.RejectReason = &v
	return s
}

func (s *SaasQueryMerchantStatusResponseDataPayStatusInfoValue) SetLegalValidationUrl(v string) *SaasQueryMerchantStatusResponseDataPayStatusInfoValue {
	s.LegalValidationUrl = &v
	return s
}

type SaasQueryWithdrawOrderRequest struct {
	MerchantUid  *string            `json:"merchant_uid,omitempty" xml:"merchant_uid,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OutOrderId   *string            `json:"out_order_id,omitempty" xml:"out_order_id,omitempty" require:"true"`
	ChannelType  *string            `json:"channel_type,omitempty" xml:"channel_type,omitempty" require:"true"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	ThirdpartyId *string            `json:"thirdparty_id,omitempty" xml:"thirdparty_id,omitempty"`
}

func (s SaasQueryWithdrawOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaasQueryWithdrawOrderRequest) GoString() string {
	return s.String()
}

func (s *SaasQueryWithdrawOrderRequest) SetMerchantUid(v string) *SaasQueryWithdrawOrderRequest {
	s.MerchantUid = &v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetHeader(v map[string]*string) *SaasQueryWithdrawOrderRequest {
	s.Header = v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetAccessToken(v string) *SaasQueryWithdrawOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetOutOrderId(v string) *SaasQueryWithdrawOrderRequest {
	s.OutOrderId = &v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetChannelType(v string) *SaasQueryWithdrawOrderRequest {
	s.ChannelType = &v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetAppId(v string) *SaasQueryWithdrawOrderRequest {
	s.AppId = &v
	return s
}

func (s *SaasQueryWithdrawOrderRequest) SetThirdpartyId(v string) *SaasQueryWithdrawOrderRequest {
	s.ThirdpartyId = &v
	return s
}
