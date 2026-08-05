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

	"github.com/bytedance/douyin-openapi-sdk-go/internal/transport"
)

func (client *Client) OrderSyncStatus(request *OrderSyncStatusRequest) (_result *OrderSyncStatusResponse, _err error) {
	return invoke[OrderSyncStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/distribution/order/sync_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"accept_time":       tea.Int64Value(request.AcceptTime),
			"behavior":          tea.IntValue(request.Behavior),
			"order_id":          tea.StringValue(request.OrderId),
			"rider_flow_time":   tea.Int64Value(request.RiderFlowTime),
			"rider_lat":         tea.Float64Value(request.RiderLat),
			"rider_lng":         tea.Float64Value(request.RiderLng),
			"rider_name":        tea.StringValue(request.RiderName),
			"rider_phone":       tea.StringValue(request.RiderPhone),
			"rider_phone_type":  tea.Int64Value(request.RiderPhoneType),
			"three_delivery_id": tea.StringValue(request.ThreeDeliveryId),
			"three_source":      tea.IntValue(request.ThreeSource),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrientedPlanTalentDetail(request *OrientedPlanTalentDetailRequest) (_result *OrientedPlanTalentDetailResponse, _err error) {
	return invoke[OrientedPlanTalentDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/oriented_plan_talent_detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"douyin_id_list": tea.StringSliceValue(request.DouyinIdList),
			"plan_id":        tea.Int64Value(request.PlanId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PassportOpenGetAccessCode(request *PassportOpenGetAccessCodeRequest) (_result *PassportOpenGetAccessCodeResponse, _err error) {
	return invoke[PassportOpenGetAccessCodeResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/passport/open/get_access_code",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeForm,
		Body: map[string]interface{}{
			"access_token": tea.StringValue(request.AccessToken),
		},
		Encoding:    transport.EncodingForm,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PassportOpenGetClientCode(request *PassportOpenGetClientCodeRequest) (_result *PassportOpenGetClientCodeResponse, _err error) {
	return invoke[PassportOpenGetClientCodeResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/passport/open/get_client_code",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeForm,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
		},
		Encoding: transport.EncodingForm,
	})
}

func (client *Client) PhysicalRoomOperate(request *PhysicalRoomOperateRequest) (_result *PhysicalRoomOperateResponse, _err error) {
	return invoke[PhysicalRoomOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/physical_room/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"op_type":    tea.IntValue(request.OpType),
			"room_ids":   tea.StringSliceValue(request.RoomIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PhysicalRoomQuery(request *PhysicalRoomQueryRequest) (_result *PhysicalRoomQueryResponse, _err error) {
	return invoke[PhysicalRoomQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/physical_room/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"need_rate_plan": tea.BoolValue(request.NeedRatePlan),
			"room_ids":       tea.StringSliceValue(request.RoomIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PhysicalRoomSave(request *PhysicalRoomSaveRequest) (_result *PhysicalRoomSaveResponse, _err error) {
	return invoke[PhysicalRoomSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/physical_room/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"by_fileds":  tea.StringSliceValue(request.ByFileds),
			"poi_id":     tea.StringValue(request.PoiId),
			"room_info":  request.RoomInfo,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PhysicalRoomSearch(request *PhysicalRoomSearchRequest) (_result *PhysicalRoomSearchResponse, _err error) {
	return invoke[PhysicalRoomSearchResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/physical_room/search/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"poi_ids":    tea.StringSliceValue(request.PoiIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PlayletBusinessUpload(request *PlayletBusinessUploadRequest) (_result *PlayletBusinessUploadResponse, _err error) {
	return invoke[PlayletBusinessUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/playlet_business/upload/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"context":    request.Context,
			"event_type": tea.StringValue(request.EventType),
			"properties": tea.StringValue(request.Properties),
			"timestamp":  tea.Int64Value(request.Timestamp),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiClaim(request *PoiClaimRequest) (_result *PoiClaimResponse, _err error) {
	return invoke[PoiClaimResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/poi/claim/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"datas":       request.Datas,
			"target_type": tea.IntValue(request.TargetType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiDecorate(request *PoiDecorateRequest) (_result *PoiDecorateResponse, _err error) {
	return invoke[PoiDecorateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/poi/decorate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"datas": request.Datas,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiOrientedPlanDetail(request *PoiOrientedPlanDetailRequest) (_result *PoiOrientedPlanDetailResponse, _err error) {
	return invoke[PoiOrientedPlanDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/oriented_plan_detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"plan_id_list": tea.Int64ValueSlice(request.PlanIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiOrientedPlanList(request *PoiOrientedPlanListRequest) (_result *PoiOrientedPlanListResponse, _err error) {
	return invoke[PoiOrientedPlanListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/oriented_plan_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"spu_id_list": tea.Int64ValueSlice(request.SpuIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiQuery(request *PoiQueryRequest) (_result *PoiQueryResponse, _err error) {
	return invoke[PoiQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/order/poi/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"order_id":   tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiSaveCommonPlan(request *PoiSaveCommonPlanRequest) (_result *PoiSaveCommonPlanResponse, _err error) {
	return invoke[PoiSaveCommonPlanResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/save_common_plan/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"commission_rate": tea.Int64Value(request.CommissionRate),
			"content_type":    tea.Int32Value(request.ContentType),
			"plan_id":         tea.Int64Value(request.PlanId),
			"spu_id":          tea.Int64Value(request.SpuId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiScroll(request *PoiScrollRequest) (_result *PoiScrollResponse, _err error) {
	return invoke[PoiScrollResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/poi/scroll/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"page_size":  tea.Int32Value(request.PageSize),
			"scroll_id":  tea.StringValue(request.ScrollId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiSellOutDetail(request *PoiSellOutDetailRequest) (_result *PoiSellOutDetailResponse, _err error) {
	return invoke[PoiSellOutDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/affiliated/poi_sell_out/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":              request.Base,
			"account_id":        tea.Int64Value(request.AccountId),
			"affiliated_id":     tea.Int64Value(request.AffiliatedId),
			"out_affiliated_id": tea.StringValue(request.OutAffiliatedId),
			"poi_ids":           tea.Int64ValueSlice(request.PoiIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiSellOutSave(request *PoiSellOutSaveRequest) (_result *PoiSellOutSaveResponse, _err error) {
	return invoke[PoiSellOutSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/product/poi_sell_out/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":              request.Base,
			"account_id":        tea.StringValue(request.AccountId),
			"poi_sell_out_rule": request.PoiSellOutRule,
			"product_id":        tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiSync(request *PoiSyncRequest) (_result *PoiSyncResponse, _err error) {
	return invoke[PoiSyncResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/poi/sync/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"datas":       request.Datas,
			"target_type": tea.IntValue(request.TargetType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiTaskQuery(request *PoiTaskQueryRequest) (_result *PoiTaskQueryResponse, _err error) {
	return invoke[PoiTaskQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/poi/task/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"task_ids": tea.Int64ValueSlice(request.TaskIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiUpdate(request *PoiUpdateRequest) (_result *PoiUpdateResponse, _err error) {
	return invoke[PoiUpdateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/poi/update/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"datas": request.Datas,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoipriceSave(request *PoipriceSaveRequest) (_result *PoipriceSaveResponse, _err error) {
	return invoke[PoipriceSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/poiprice/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":               request.Base,
			"account_id":         tea.Int64Value(request.AccountId),
			"ignore_fail_poi":    tea.BoolValue(request.IgnoreFailPoi),
			"out_product_id":     tea.StringValue(request.OutProductId),
			"poi_sku_price_list": request.PoiSkuPriceList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiproductDetail(request *PoiproductDetailRequest) (_result *PoiproductDetailResponse, _err error) {
	return invoke[PoiproductDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/poiproduct/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.Int64Value(request.AccountId),
			"out_product_id":  tea.StringValue(request.OutProductId),
			"out_sku_id_list": tea.StringSliceValue(request.OutSkuIdList),
			"poi_ext_ids":     tea.StringSliceValue(request.PoiExtIds),
			"poi_ids":         tea.Int64ValueSlice(request.PoiIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoiproductOperate(request *PoiproductOperateRequest) (_result *PoiproductOperateResponse, _err error) {
	return invoke[PoiproductOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/poiproduct/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":        tea.Int64Value(request.AccountId),
			"ignore_fail_poi":   tea.BoolValue(request.IgnoreFailPoi),
			"operate_info_list": request.OperateInfoList,
			"out_product_id":    tea.StringValue(request.OutProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoistockDetail(request *PoistockDetailRequest) (_result *PoistockDetailResponse, _err error) {
	return invoke[PoistockDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/pickupcoupon/poistock/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":           tea.Int64Value(request.AccountId),
			"root_life_account_id": tea.Int64Value(request.RootLifeAccountId),
			"sku_poi_list":         request.SkuPoiList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PoistockSave(request *PoistockSaveRequest) (_result *PoistockSaveResponse, _err error) {
	return invoke[PoistockSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/pickupcoupon/poistock/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":           tea.Int64Value(request.AccountId),
			"poi_stock_list":       request.PoiStockList,
			"root_life_account_id": tea.Int64Value(request.RootLifeAccountId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PostingBindVideo(request *PostingBindVideoRequest) (_result *PostingBindVideoResponse, _err error) {
	return invoke[PostingBindVideoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/task/posting/bind_video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"task_id":  tea.StringValue(request.TaskId),
			"video_id": tea.StringValue(request.VideoId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PostingCreate(request *PostingCreateRequest) (_result *PostingCreateResponse, _err error) {
	return invoke[PostingCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/task/posting/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"end_time":       tea.Int64Value(request.EndTime),
			"start_time":     tea.Int64Value(request.StartTime),
			"task_condition": request.TaskCondition,
			"task_name":      tea.StringValue(request.TaskName),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PostingUser(request *PostingUserRequest) (_result *PostingUserResponse, _err error) {
	return invoke[PostingUserResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/task/posting/user/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"target_open_id": tea.StringValue(request.TargetOpenId),
			"task_id":        tea.StringValue(request.TaskId),
			"video_id":       tea.StringValue(request.VideoId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PresaleCouponSaveBookCalendarStock(request *PresaleCouponSaveBookCalendarStockRequest) (_result *PresaleCouponSaveBookCalendarStockResponse, _err error) {
	return invoke[PresaleCouponSaveBookCalendarStockResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/presale_coupon/save/book_calendar_stock/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":                 tea.StringValue(request.AccountId),
			"book_calendar_stock_config": request.BookCalendarStockConfig,
			"product_id":                 tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PresaleRateplanSave(request *PresaleRateplanSaveRequest) (_result *PresaleRateplanSaveResponse, _err error) {
	return invoke[PresaleRateplanSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/presale/rateplan/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"rate_plan":  request.RatePlan,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PricePush(request *PricePushRequest) (_result *PricePushResponse, _err error) {
	return invoke[PricePushResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion/price/push/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":          tea.StringValue(request.AccountId),
			"applicable_date":     request.ApplicableDate,
			"applicable_resource": request.ApplicableResource,
			"price_member_levels": request.PriceMemberLevels,
			"promotion_id":        tea.StringValue(request.PromotionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PriceSave(request *PriceSaveRequest) (_result *PriceSaveResponse, _err error) {
	return invoke[PriceSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/price/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"aris":       request.Aris,
			"hotel_id":   tea.StringValue(request.HotelId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PrivacySettingAdd(request *PrivacySettingAddRequest) (_result *PrivacySettingAddResponse, _err error) {
	return invoke[PrivacySettingAddResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/privacy_setting/add/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"contact_way":       tea.IntValue(request.ContactWay),
			"email":             tea.StringValue(request.Email),
			"is_privacy_config": tea.BoolValue(request.IsPrivacyConfig),
			"land_line":         tea.StringValue(request.LandLine),
			"phone":             tea.StringValue(request.Phone),
			"privacy_item_list": request.PrivacyItemList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PrivacySettingQuery(request *PrivacySettingQueryRequest) (_result *PrivacySettingQueryResponse, _err error) {
	return invoke[PrivacySettingQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/privacy_setting/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductCommissionGet(request *ProductCommissionGetRequest) (_result *ProductCommissionGetResponse, _err error) {
	return invoke[ProductCommissionGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/product_commission/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"order_id": tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductCommissionMget(request *ProductCommissionMgetRequest) (_result *ProductCommissionMgetResponse, _err error) {
	return invoke[ProductCommissionMgetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/partner/product_commission/mget/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":    tea.StringValue(request.OrderId),
			"product_ids": tea.StringSliceValue(request.ProductIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductCommissionQuery(request *ProductCommissionQueryRequest) (_result *ProductCommissionQueryResponse, _err error) {
	return invoke[ProductCommissionQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/product_commission/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"order_id": tea.StringValue(request.OrderId),
			"page":     tea.Int32Value(request.Page),
			"size":     tea.Int32Value(request.Size),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductCommissionSave(request *ProductCommissionSaveRequest) (_result *ProductCommissionSaveResponse, _err error) {
	return invoke[ProductCommissionSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/partner/product_commission/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":      tea.StringValue(request.OrderId),
			"product_items": request.ProductItems,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductDetail(request *ProductDetailRequest) (_result *ProductDetailResponse, _err error) {
	return invoke[ProductDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/product/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.Int64Value(request.ProductId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductDraftList(request *ProductDraftListRequest) (_result *ProductDraftListResponse, _err error) {
	return invoke[ProductDraftListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/product/draft/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"count":      tea.Int32Value(request.Count),
			"cursor":     tea.Int32Value(request.Cursor),
			"dy_poi_id":  tea.Int64Value(request.DyPoiId),
			"status":     tea.IntValueSlice(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductDraftQuery(request *ProductDraftQueryRequest) (_result *ProductDraftQueryResponse, _err error) {
	return invoke[ProductDraftQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/product/draft/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"count":      tea.Int64Value(request.Count),
			"cursor":     tea.StringValue(request.Cursor),
			"status":     tea.IntValue(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductGetPois(request *ProductGetPoisRequest) (_result *ProductGetPoisResponse, _err error) {
	return invoke[ProductGetPoisResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/product/get_pois/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":               request.Base,
			"account_id":         tea.StringValue(request.AccountId),
			"need_sell_out_info": tea.BoolValue(request.NeedSellOutInfo),
			"product_ids":        tea.Int64ValueSlice(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductOnlineGet(request *ProductOnlineGetRequest) (_result *ProductOnlineGetResponse, _err error) {
	return invoke[ProductOnlineGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v2/x_goods/product/online/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":        request.Base,
			"account_id":  tea.StringValue(request.AccountId),
			"product_ids": tea.Int64ValueSlice(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductOnlineList(request *ProductOnlineListRequest) (_result *ProductOnlineListResponse, _err error) {
	return invoke[ProductOnlineListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v2/x_goods/product/online/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":           request.Base,
			"account_id":     tea.StringValue(request.AccountId),
			"count":          tea.Int32Value(request.Count),
			"cursor":         tea.Int32Value(request.Cursor),
			"dy_poi_id":      tea.Int64Value(request.DyPoiId),
			"page_no":        tea.Int32Value(request.PageNo),
			"page_size":      tea.Int32Value(request.PageSize),
			"search_by_page": tea.BoolValue(request.SearchByPage),
			"status":         tea.IntValueSlice(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductOnlineQuery(request *ProductOnlineQueryRequest) (_result *ProductOnlineQueryResponse, _err error) {
	return invoke[ProductOnlineQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/product/online/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":         tea.StringValue(request.AccountId),
			"count":              tea.Int64Value(request.Count),
			"cursor":             tea.StringValue(request.Cursor),
			"ext_ids":            tea.Int64ValueSlice(request.ExtIds),
			"goods_creator_type": tea.IntValue(request.GoodsCreatorType),
			"goods_query_type":   tea.IntValue(request.GoodsQueryType),
			"poi_ids":            tea.Int64ValueSlice(request.PoiIds),
			"product_name":       tea.StringValue(request.ProductName),
			"query_all_poi":      tea.BoolValue(request.QueryAllPoi),
			"status":             tea.IntValue(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductSave(request *ProductSaveRequest) (_result *ProductSaveResponse, _err error) {
	return invoke[ProductSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/product/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"product":    request.Product,
			"sku":        request.Sku,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductUpdatePrice(request *ProductUpdatePriceRequest) (_result *ProductUpdatePriceResponse, _err error) {
	return invoke[ProductUpdatePriceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/product/update_price/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":          request.Base,
			"account_id":    tea.StringValue(request.AccountId),
			"actual_amount": tea.Int64Value(request.ActualAmount),
			"product_id":    tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductUpdateStatus(request *ProductUpdateStatusRequest) (_result *ProductUpdateStatusResponse, _err error) {
	return invoke[ProductUpdateStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/product/update_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":         request.Base,
			"account_id":   tea.StringValue(request.AccountId),
			"operate_type": tea.IntValue(request.OperateType),
			"product_ids":  tea.StringSliceValue(request.ProductIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ProductUpdateStock(request *ProductUpdateStockRequest) (_result *ProductUpdateStockResponse, _err error) {
	return invoke[ProductUpdateStockResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/product/update_stock/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"limit_type": tea.IntValue(request.LimitType),
			"out_sku_id": tea.StringValue(request.OutSkuId),
			"pois":       request.Pois,
			"product_id": tea.StringValue(request.ProductId),
			"sku_id":     tea.StringValue(request.SkuId),
			"stock_qty":  tea.Int64Value(request.StockQty),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PromotionExit(request *PromotionExitRequest) (_result *PromotionExitResponse, _err error) {
	return invoke[PromotionExitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion/exit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":    tea.StringValue(request.AccountId),
			"promotion_id":  tea.StringValue(request.PromotionId),
			"quit_resource": request.QuitResource,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PromotionInfoQuery(request *PromotionInfoQueryRequest) (_result *PromotionInfoQueryResponse, _err error) {
	return invoke[PromotionInfoQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion_info/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":   tea.StringValue(request.AccountId),
			"hotel_id":     tea.StringSliceValue(request.HotelId),
			"page":         tea.Int32Value(request.Page),
			"page_size":    tea.Int32Value(request.PageSize),
			"promotion_id": tea.StringSliceValue(request.PromotionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PromotionPush(request *PromotionPushRequest) (_result *PromotionPushResponse, _err error) {
	return invoke[PromotionPushResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion/push/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":                   tea.StringValue(request.AccountId),
			"active":                       tea.BoolValue(request.Active),
			"applicable_date":              request.ApplicableDate,
			"applicable_resource":          request.ApplicableResource,
			"daily_sale_promotion":         request.DailySalePromotion,
			"early_bird_promotion":         request.EarlyBirdPromotion,
			"flash_sale_promotion":         request.FlashSalePromotion,
			"hotel_new_customer_promotion": request.HotelNewCustomerPromotion,
			"is_auto_extension":            tea.BoolValue(request.IsAutoExtension),
			"night_sale_promotion":         request.NightSalePromotion,
			"promotion_basic_info":         request.PromotionBasicInfo,
			"promotion_id":                 tea.StringValue(request.PromotionId),
			"succession_promotion":         request.SuccessionPromotion,
			"unapplicable_date":            request.UnapplicableDate,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PurchaseInfo(request *PurchaseInfoRequest) (_result *PurchaseInfoResponse, _err error) {
	return invoke[PurchaseInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/market/service/user/insert/purchase/info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"duration":        tea.Int64Value(request.Duration),
			"open_id":         tea.StringValue(request.OpenId),
			"out_trade_no":    tea.StringValue(request.OutTradeNo),
			"period_type":     tea.IntValue(request.PeriodType),
			"purchase_time":   tea.Int64Value(request.PurchaseTime),
			"service_id":      tea.StringValue(request.ServiceId),
			"service_mode_id": tea.StringValue(request.ServiceModeId),
			"usage_times":     tea.Int64Value(request.UsageTimes),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) PurchaseList(request *PurchaseListRequest) (_result *PurchaseListResponse, _err error) {
	return invoke[PurchaseListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/market/service/user/purchase/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"is_test_env": tea.BoolValue(request.IsTestEnv),
			"open_id":     tea.StringValue(request.OpenId),
			"service_id":  tea.StringValue(request.ServiceId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QaOpenapiSdkPostCommon(request *QaOpenapiSdkPostCommonRequest) (_result *QaOpenapiSdkPostCommonResponse, _err error) {
	return invoke[QaOpenapiSdkPostCommonResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/qa_openapi_test/v1/qa_openapi_sdk/post_common/",
		Host:        tea.String("open-platform-qa.byted.org"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"default_query":  tea.StringValue(request.DefaultQuery),
			"optional_query": tea.StringValue(request.OptionalQuery),
			"required_query": tea.StringValue(request.RequiredQuery),
		},
		Body: map[string]interface{}{
			"default_list":    tea.StringSliceValue(request.DefaultList),
			"default_map":     request.DefaultMap,
			"default_struct":  request.DefaultStruct,
			"optional_list":   tea.StringSliceValue(request.OptionalList),
			"optional_map":    request.OptionalMap,
			"optional_struct": request.OptionalStruct,
			"qa_extra":        request.QaExtra,
			"required_list":   tea.StringSliceValue(request.RequiredList),
			"required_map":    request.RequiredMap,
			"required_struct": request.RequiredStruct,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QrcodeCreate(request *QrcodeCreateRequest) (_result *QrcodeCreateResponse, _err error) {
	return invoke[QrcodeCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/qrcode/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_name":       tea.StringValue(request.AppName),
			"appid":          tea.StringValue(request.Appid),
			"background":     request.Background,
			"is_circle_code": tea.BoolValue(request.IsCircleCode),
			"line_color":     request.LineColor,
			"path":           tea.StringValue(request.Path),
			"set_icon":       tea.BoolValue(request.SetIcon),
			"width":          tea.Int32Value(request.Width),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QualSearch(request *QualSearchRequest) (_result *QualSearchResponse, _err error) {
	return invoke[QualSearchResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/account/qual/search/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
		},
		Body: map[string]interface{}{
			"data": request.Data,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryActivityMetaData(request *QueryActivityMetaDataRequest) (_result *QueryActivityMetaDataResponse, _err error) {
	return invoke[QueryActivityMetaDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/query_activity_meta_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"coupon_name":    tea.StringValue(request.CouponName),
			"discount_type":  tea.IntValue(request.DiscountType),
			"open_id":        tea.StringValue(request.OpenId),
			"page_num":       tea.Int32Value(request.PageNum),
			"page_size":      tea.Int32Value(request.PageSize),
			"send_scene":     tea.IntValue(request.SendScene),
			"status":         tea.IntValue(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryActivityUserCompletionStatus(request *QueryActivityUserCompletionStatusRequest) (_result *QueryActivityUserCompletionStatusResponse, _err error) {
	return invoke[QueryActivityUserCompletionStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/dy_open_api/apps/v3/activity/query_activity_user_completion_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id":  tea.Int64Value(request.ActivityId),
			"task_id_list": tea.Int64ValueSlice(request.TaskIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAdSettlementList(request *QueryAdSettlementListRequest) (_result *QueryAdSettlementListResponse, _err error) {
	return invoke[QueryAdSettlementListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v3/capacity/query_ad_settlement_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"month":  tea.StringValue(request.Month),
			"status": tea.Int32Value(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAgencyVideoDailyData(request *QueryAgencyVideoDailyDataRequest) (_result *QueryAgencyVideoDailyDataResponse, _err error) {
	return invoke[QueryAgencyVideoDailyDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_agency_video_daily_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agent_id":                 tea.Int64Value(request.AgentId),
			"app_id":                   tea.StringValue(request.AppId),
			"billing_date":             tea.StringValue(request.BillingDate),
			"douyin_id":                tea.StringValue(request.DouyinId),
			"page_num":                 tea.Int32Value(request.PageNum),
			"page_size":                tea.Int32Value(request.PageSize),
			"video_publish_end_time":   tea.Int64Value(request.VideoPublishEndTime),
			"video_publish_start_time": tea.Int64Value(request.VideoPublishStartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAgencyVideoSumData(request *QueryAgencyVideoSumDataRequest) (_result *QueryAgencyVideoSumDataResponse, _err error) {
	return invoke[QueryAgencyVideoSumDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_agency_video_sum_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agent_id":                 tea.Int64Value(request.AgentId),
			"app_id":                   tea.StringValue(request.AppId),
			"douyin_id":                tea.StringValue(request.DouyinId),
			"page_num":                 tea.Int32Value(request.PageNum),
			"page_size":                tea.Int32Value(request.PageSize),
			"video_publish_end_time":   tea.Int64Value(request.VideoPublishEndTime),
			"video_publish_start_time": tea.Int64Value(request.VideoPublishStartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAppTaskId(request *QueryAppTaskIdRequest) (_result *QueryAppTaskIdResponse, _err error) {
	return invoke[QueryAppTaskIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/query_app_task_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":             tea.StringValue(request.Appid),
			"create_end_time":   tea.Int64Value(request.CreateEndTime),
			"create_start_time": tea.Int64Value(request.CreateStartTime),
			"task_category":     tea.IntValue(request.TaskCategory),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAppTestRelation(request *QueryAppTestRelationRequest) (_result *QueryAppTestRelationResponse, _err error) {
	return invoke[QueryAppTestRelationResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/industry/v1/solution/query_app_test_relation/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"type": tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryApplyPermissionStatus(request *QueryApplyPermissionStatusRequest) (_result *QueryApplyPermissionStatusResponse, _err error) {
	return invoke[QueryApplyPermissionStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/video_mount/query_apply_permission_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"component_appid": tea.StringValue(request.ComponentAppid),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAri(request *QueryAriRequest) (_result *QueryAriResponse, _err error) {
	return invoke[QueryAriResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/presale_coupon/query/ari/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":               tea.StringValue(request.AccountId),
			"booking_stock_end_date":   tea.StringValue(request.BookingStockEndDate),
			"booking_stock_start_date": tea.StringValue(request.BookingStockStartDate),
			"product_id":               tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAwemeRelationList(request *QueryAwemeRelationListRequest) (_result *QueryAwemeRelationListResponse, _err error) {
	return invoke[QueryAwemeRelationListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/query_aweme_relation_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"page_num":  tea.Int64Value(request.PageNum),
			"page_size": tea.Int64Value(request.PageSize),
			"type":      tea.StringValue(request.Type),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryAwemeVideoKeywordList(request *QueryAwemeVideoKeywordListRequest) (_result *QueryAwemeVideoKeywordListResponse, _err error) {
	return invoke[QueryAwemeVideoKeywordListResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/apps/v1/capacity/query_aweme_video_keyword_list/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"page_num":  tea.Int32Value(request.PageNum),
			"page_size": tea.Int32Value(request.PageSize),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryCommonPlanTalentList(request *QueryCommonPlanTalentListRequest) (_result *QueryCommonPlanTalentListResponse, _err error) {
	return invoke[QueryCommonPlanTalentListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/query_common_plan_talent_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"page_num":  tea.Int32Value(request.PageNum),
			"page_size": tea.Int32Value(request.PageSize),
			"plan_id":   tea.Int64Value(request.PlanId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryComponentWithData(request *QueryComponentWithDataRequest) (_result *QueryComponentWithDataResponse, _err error) {
	return invoke[QueryComponentWithDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_component_with_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"componentId_list": tea.StringSliceValue(request.ComponentIdList),
			"end_time":         tea.Int64Value(request.EndTime),
			"page_no":          tea.Int64Value(request.PageNo),
			"page_size":        tea.Int64Value(request.PageSize),
			"start_time":       tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryComponentWithDetail(request *QueryComponentWithDetailRequest) (_result *QueryComponentWithDetailResponse, _err error) {
	return invoke[QueryComponentWithDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_component_with_detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"componentId_list": tea.StringSliceValue(request.ComponentIdList),
			"end_time":         tea.Int64Value(request.EndTime),
			"is_query_live":    tea.BoolValue(request.IsQueryLive),
			"is_query_video":   tea.BoolValue(request.IsQueryVideo),
			"page_no":          tea.Int64Value(request.PageNo),
			"page_size":        tea.Int64Value(request.PageSize),
			"start_time":       tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryComponentWithOverview(request *QueryComponentWithOverviewRequest) (_result *QueryComponentWithOverviewResponse, _err error) {
	return invoke[QueryComponentWithOverviewResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_component_with_overview/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"componentId_list": tea.StringSliceValue(request.ComponentIdList),
			"end_time":         tea.Int64Value(request.EndTime),
			"start_time":       tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryComponentWithSource(request *QueryComponentWithSourceRequest) (_result *QueryComponentWithSourceResponse, _err error) {
	return invoke[QueryComponentWithSourceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_component_with_source/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"componentId_list": tea.StringSliceValue(request.ComponentIdList),
			"end_time":         tea.Int64Value(request.EndTime),
			"start_time":       tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryContact(request *QueryContactRequest) (_result *QueryContactResponse, _err error) {
	return invoke[QueryContactResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/retail/order/query/contact/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":            tea.StringValue(request.OrderId),
			"source_phone_number": tea.StringValue(request.SourcePhoneNumber),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryCreatedTplList(request *QueryCreatedTplListRequest) (_result *QueryCreatedTplListResponse, _err error) {
	return invoke[QueryCreatedTplListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/notification/v2/subscription/query_created_tpl_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"page_num":  tea.Int32Value(request.PageNum),
			"page_size": tea.Int32Value(request.PageSize),
			"status":    tea.IntValue(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryDealDataWithConversion(request *QueryDealDataWithConversionRequest) (_result *QueryDealDataWithConversionResponse, _err error) {
	return invoke[QueryDealDataWithConversionResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_deal_data_with_conversion/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"end_time":    tea.Int64Value(request.EndTime),
			"host_name":   tea.StringValue(request.HostName),
			"scenes_list": tea.StringSliceValue(request.ScenesList),
			"start_time":  tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryDealOverviewData(request *QueryDealOverviewDataRequest) (_result *QueryDealOverviewDataResponse, _err error) {
	return invoke[QueryDealOverviewDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_deal_overview_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"host_name":    tea.StringValue(request.HostName),
			"start_time":   tea.Int64Value(request.StartTime),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryItemOrderInfo(request *QueryItemOrderInfoRequest) (_result *QueryItemOrderInfoResponse, _err error) {
	return invoke[QueryItemOrderInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/order/query_item_order_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"item_order_id_list": tea.StringSliceValue(request.ItemOrderIdList),
			"order_id":           tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryLiveDealData(request *QueryLiveDealDataRequest) (_result *QueryLiveDealDataResponse, _err error) {
	return invoke[QueryLiveDealDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_live_deal_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"live_room_id": tea.Int64Value(request.LiveRoomId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryLiveRoomData(request *QueryLiveRoomDataRequest) (_result *QueryLiveRoomDataResponse, _err error) {
	return invoke[QueryLiveRoomDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_live_room_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"live_room_id": tea.Int64Value(request.LiveRoomId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryLiveWithShortId(request *QueryLiveWithShortIdRequest) (_result *QueryLiveWithShortIdResponse, _err error) {
	return invoke[QueryLiveWithShortIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_live_with_short_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_short_id_list": tea.StringSliceValue(request.AwemeShortIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"host_name":           tea.StringValue(request.HostName),
			"page_no":             tea.Int64Value(request.PageNo),
			"page_size":           tea.Int64Value(request.PageSize),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryMaSubService(request *QueryMaSubServiceRequest) (_result *QueryMaSubServiceResponse, _err error) {
	return invoke[QueryMaSubServiceResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/query_ma_sub_service/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"approval_state": tea.Int64Value(request.ApprovalState),
			"page_no":        tea.Int64Value(request.PageNo),
			"page_size":      tea.Int64Value(request.PageSize),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryOrder(request *QueryOrderRequest) (_result *QueryOrderResponse, _err error) {
	return invoke[QueryOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/comprehensive/trade/query/order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"book_id":    tea.StringValue(request.BookId),
			"end_time":   tea.Int64Value(request.EndTime),
			"page_num":   tea.Int32Value(request.PageNum),
			"page_size":  tea.Int32Value(request.PageSize),
			"poi_id":     tea.StringValue(request.PoiId),
			"start_time": tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryProductDealData(request *QueryProductDealDataRequest) (_result *QueryProductDealDataResponse, _err error) {
	return invoke[QueryProductDealDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_product_deal_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"host_name":  tea.StringValue(request.HostName),
			"page_num":   tea.Int64Value(request.PageNum),
			"page_size":  tea.Int64Value(request.PageSize),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryRealTimeUserData(request *QueryRealTimeUserDataRequest) (_result *QueryRealTimeUserDataResponse, _err error) {
	return invoke[QueryRealTimeUserDataResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/platform/v2/data_analysis/query_real_time_user_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"host_name":    tea.StringValue(request.HostName),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryRecordByCert(request *QueryRecordByCertRequest) (_result *QueryRecordByCertResponse, _err error) {
	return invoke[QueryRecordByCertResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/ledger/query_record_by_cert/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"certificate_ids": tea.StringSliceValue(request.CertificateIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryRecordBySubfulfil(request *QueryRecordBySubfulfilRequest) (_result *QueryRecordBySubfulfilResponse, _err error) {
	return invoke[QueryRecordBySubfulfilResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/bill/query_record_by_subfulfil/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"sub_fulfil_id": tea.StringSliceValue(request.SubFulfilId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuerySearchTagList(request *QuerySearchTagListRequest) (_result *QuerySearchTagListResponse, _err error) {
	return invoke[QuerySearchTagListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/query_search_tag_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryShortLiveDataWithId(request *QueryShortLiveDataWithIdRequest) (_result *QueryShortLiveDataWithIdResponse, _err error) {
	return invoke[QueryShortLiveDataWithIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_short_live_data_with_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_short_id_list": tea.StringSliceValue(request.AwemeShortIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"host_name":           tea.StringValue(request.HostName),
			"item_id_list":        tea.StringSliceValue(request.ItemIdList),
			"open_item_id_list":   tea.StringSliceValue(request.OpenItemIdList),
			"page_no":             tea.Int64Value(request.PageNo),
			"page_size":           tea.Int64Value(request.PageSize),
			"query_bind_type":     tea.Int32Value(request.QueryBindType),
			"query_data_type":     tea.Int32Value(request.QueryDataType),
			"room_id_list":        tea.StringSliceValue(request.RoomIdList),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryShortLiveIdWithAwemeid(request *QueryShortLiveIdWithAwemeidRequest) (_result *QueryShortLiveIdWithAwemeidResponse, _err error) {
	return invoke[QueryShortLiveIdWithAwemeidResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_short_live_id_with_awemeid/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_short_id_list": tea.StringSliceValue(request.AwemeShortIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"host_name":           tea.StringValue(request.HostName),
			"query_bind_type":     tea.Int32Value(request.QueryBindType),
			"query_data_type":     tea.Int32Value(request.QueryDataType),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuerySimpleQrBindList(request *QuerySimpleQrBindListRequest) (_result *QuerySimpleQrBindListResponse, _err error) {
	return invoke[QuerySimpleQrBindListResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/apps/v2/capacity/query_simple_qr_bind_list/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"page_num":  tea.Int64Value(request.PageNum),
			"page_size": tea.Int64Value(request.PageSize),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuerySmallHomeOrderData(request *QuerySmallHomeOrderDataRequest) (_result *QuerySmallHomeOrderDataResponse, _err error) {
	return invoke[QuerySmallHomeOrderDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_small_home_order_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"page_num":   tea.Int64Value(request.PageNum),
			"page_size":  tea.Int64Value(request.PageSize),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuerySmallHomeOverviewData(request *QuerySmallHomeOverviewDataRequest) (_result *QuerySmallHomeOverviewDataResponse, _err error) {
	return invoke[QuerySmallHomeOverviewDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_small_home_overview_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuerySmallHomeRoomData(request *QuerySmallHomeRoomDataRequest) (_result *QuerySmallHomeRoomDataResponse, _err error) {
	return invoke[QuerySmallHomeRoomDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_small_home_room_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryTaskVideoDailyData(request *QueryTaskVideoDailyDataRequest) (_result *QueryTaskVideoDailyDataResponse, _err error) {
	return invoke[QueryTaskVideoDailyDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_task_video_daily_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agency_client_key":        tea.StringValue(request.AgencyClientKey),
			"billing_date":             tea.StringValue(request.BillingDate),
			"douyin_id":                tea.StringValue(request.DouyinId),
			"page_num":                 tea.Int32Value(request.PageNum),
			"page_size":                tea.Int32Value(request.PageSize),
			"video_publish_end_time":   tea.Int64Value(request.VideoPublishEndTime),
			"video_publish_start_time": tea.Int64Value(request.VideoPublishStartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryTaskVideoData(request *QueryTaskVideoDataRequest) (_result *QueryTaskVideoDataResponse, _err error) {
	return invoke[QueryTaskVideoDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_task_video_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"page_num":                 tea.Int32Value(request.PageNum),
			"page_size":                tea.Int32Value(request.PageSize),
			"task_ids":                 tea.Int64ValueSlice(request.TaskIds),
			"video_publish_end_time":   tea.Int64Value(request.VideoPublishEndTime),
			"video_publish_start_time": tea.Int64Value(request.VideoPublishStartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryTaskVideoStatus(request *QueryTaskVideoStatusRequest) (_result *QueryTaskVideoStatusResponse, _err error) {
	return invoke[QueryTaskVideoStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_task_video_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agent_id":                 tea.Int64Value(request.AgentId),
			"app_id":                   tea.StringValue(request.AppId),
			"douyin_id":                tea.StringValue(request.DouyinId),
			"page_num":                 tea.Int32Value(request.PageNum),
			"page_size":                tea.Int32Value(request.PageSize),
			"video_publish_end_time":   tea.Int64Value(request.VideoPublishEndTime),
			"video_publish_start_time": tea.Int64Value(request.VideoPublishStartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryUserInteractTask(request *QueryUserInteractTaskRequest) (_result *QueryUserInteractTaskResponse, _err error) {
	return invoke[QueryUserInteractTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/douyin/query_user_interact_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":  tea.StringValue(request.AppId),
			"open_id": tea.StringValue(request.OpenId),
			"task_id": tea.StringSliceValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryUserPortraitData(request *QueryUserPortraitDataRequest) (_result *QueryUserPortraitDataResponse, _err error) {
	return invoke[QueryUserPortraitDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_user_portrait_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"host_name":    tea.StringValue(request.HostName),
			"start_time":   tea.Int64Value(request.StartTime),
			"user_type":    tea.StringValue(request.UserType),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryVideoDealData(request *QueryVideoDealDataRequest) (_result *QueryVideoDealDataResponse, _err error) {
	return invoke[QueryVideoDealDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_video_deal_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"host_name":  tea.StringValue(request.HostName),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryVideoSumData(request *QueryVideoSumDataRequest) (_result *QueryVideoSumDataResponse, _err error) {
	return invoke[QueryVideoSumDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_video_sum_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"video_ids": tea.Int64ValueSlice(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryVideoWithSource(request *QueryVideoWithSourceRequest) (_result *QueryVideoWithSourceResponse, _err error) {
	return invoke[QueryVideoWithSourceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_video_with_source/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_short_id_list": tea.StringSliceValue(request.AwemeShortIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"host_name":           tea.StringValue(request.HostName),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QueryViolateTalentList(request *QueryViolateTalentListRequest) (_result *QueryViolateTalentListResponse, _err error) {
	return invoke[QueryViolateTalentListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_violate_talent_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) QuizGet(request *QuizGetRequest) (_result *QuizGetResponse, _err error) {
	return invoke[QuizGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/quiz/get",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"level": tea.Int32Value(request.Level),
			"num":   tea.Int32Value(request.Num),
			"type":  tea.Int32Value(request.Type),
		},
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RateplanSave(request *RateplanSaveRequest) (_result *RateplanSaveResponse, _err error) {
	return invoke[RateplanSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/rateplan/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"rate_plan":  request.RatePlan,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RecallMsg(request *RecallMsgRequest) (_result *RecallMsgResponse, _err error) {
	return invoke[RecallMsgResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/recall/msg/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"conversation_id":   tea.StringValue(request.ConversationId),
			"conversation_type": tea.Int32Value(request.ConversationType),
			"msg_id":            tea.StringValue(request.MsgId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RefundApply(request *RefundApplyRequest) (_result *RefundApplyResponse, _err error) {
	return invoke[RefundApplyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/retail/order/refund/apply/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"apply_refund_item_list":  request.ApplyRefundItemList,
			"merchant_desc":           tea.StringValue(request.MerchantDesc),
			"operator_id":             tea.Int64Value(request.OperatorId),
			"operator_name":           tea.StringValue(request.OperatorName),
			"order_id":                tea.StringValue(request.OrderId),
			"order_out_after_sale_id": tea.StringValue(request.OrderOutAfterSaleId),
			"refund_reason":           tea.IntValue(request.RefundReason),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RefundAudit(request *RefundAuditRequest) (_result *RefundAuditResponse, _err error) {
	return invoke[RefundAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/refund/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"after_sale_id":  tea.StringValue(request.AfterSaleId),
			"certificate":    request.Certificate,
			"certificate_id": tea.StringValue(request.CertificateId),
			"code":           tea.StringValue(request.Code),
			"reason":         tea.StringValue(request.Reason),
			"result":         tea.IntValue(request.Result),
			"voucher":        tea.StringValue(request.Voucher),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RelationQuery(request *RelationQueryRequest) (_result *RelationQueryResponse, _err error) {
	return invoke[RelationQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/poi/match/relation/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"ext_ids": tea.StringSliceValue(request.ExtIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RemaintimesDecr(request *RemaintimesDecrRequest) (_result *RemaintimesDecrResponse, _err error) {
	return invoke[RemaintimesDecrResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/market/service/user/remaintimes/decr/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"count":           tea.Int32Value(request.Count),
			"is_test_env":     tea.BoolValue(request.IsTestEnv),
			"open_id":         tea.StringValue(request.OpenId),
			"out_trade_no":    tea.StringValue(request.OutTradeNo),
			"service_id":      tea.StringValue(request.ServiceId),
			"service_mode_id": tea.StringValue(request.ServiceModeId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ReplyReplyUserText(request *ReplyReplyUserTextRequest) (_result *ReplyReplyUserTextResponse, _err error) {
	return invoke[ReplyReplyUserTextResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/reply/reply_user_text",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"content":         tea.StringValue(request.Content),
			"conversation_id": tea.StringValue(request.ConversationId),
			"create_time":     tea.Int64Value(request.CreateTime),
			"micro_game_id":   tea.StringValue(request.MicroGameId),
			"msg_id":          tea.StringValue(request.MsgId),
			"msg_type":        tea.StringValue(request.MsgType),
			"sender_name":     tea.StringValue(request.SenderName),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ReportTaskCreate(request *ReportTaskCreateRequest) (_result *ReportTaskCreateResponse, _err error) {
	return invoke[ReportTaskCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/report/task/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":            tea.StringValue(request.AccountId),
			"additional_image_urls": tea.StringSliceValue(request.AdditionalImageUrls),
			"additional_info":       tea.StringValue(request.AdditionalInfo),
			"address":               tea.StringValue(request.Address),
			"city":                  tea.StringValue(request.City),
			"latitude":              tea.StringValue(request.Latitude),
			"longitude":             tea.StringValue(request.Longitude),
			"open_status":           tea.Int32Value(request.OpenStatus),
			"open_times":            request.OpenTimes,
			"poi_id":                tea.StringValue(request.PoiId),
			"poi_name":              tea.StringValue(request.PoiName),
			"province":              tea.StringValue(request.Province),
			"tel_list":              tea.StringSliceValue(request.TelList),
			"type_code":             tea.StringValue(request.TypeCode),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ReportTaskView(request *ReportTaskViewRequest) (_result *ReportTaskViewResponse, _err error) {
	return invoke[ReportTaskViewResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/report/task/view/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"poi_id":     tea.StringValue(request.PoiId),
			"task_id":    tea.StringValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ReserveCodeBatchImport(request *ReserveCodeBatchImportRequest) (_result *ReserveCodeBatchImportResponse, _err error) {
	return invoke[ReserveCodeBatchImportResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/reserve_code/batch_import/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":   tea.StringValue(request.AccountId),
			"codes":        tea.StringSliceValue(request.Codes),
			"expired_time": tea.Int64Value(request.ExpiredTime),
			"sku_id":       tea.StringValue(request.SkuId),
			"third_sku_id": tea.StringValue(request.ThirdSkuId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ReserveCodeBindOrderInfo(request *ReserveCodeBindOrderInfoRequest) (_result *ReserveCodeBindOrderInfoResponse, _err error) {
	return invoke[ReserveCodeBindOrderInfoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/fulfilment/reserve_code/bind_order_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"code":       tea.StringValue(request.Code),
			"sku_id":     tea.StringValue(request.SkuId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RetailOrderConfirm(request *RetailOrderConfirmRequest) (_result *RetailOrderConfirmResponse, _err error) {
	return invoke[RetailOrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/retail/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"confirm_result": tea.IntValue(request.ConfirmResult),
			"order_id":       tea.StringValue(request.OrderId),
			"pick_up_code":   tea.StringValue(request.PickUpCode),
			"pick_up_time":   tea.Int64Value(request.PickUpTime),
			"reject_reason":  request.RejectReason,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RetailOrderQuery(request *RetailOrderQueryRequest) (_result *RetailOrderQueryResponse, _err error) {
	return invoke[RetailOrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/retail/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"order_id":   tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RetailOrderRefundAudit(request *RetailOrderRefundAuditRequest) (_result *RetailOrderRefundAuditResponse, _err error) {
	return invoke[RetailOrderRefundAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/retail/order/refund/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"after_sale_id":     tea.StringValue(request.AfterSaleId),
			"audit_result":      tea.IntValue(request.AuditResult),
			"order_id":          tea.StringValue(request.OrderId),
			"out_after_sale_id": tea.StringValue(request.OutAfterSaleId),
			"reason_code":       tea.Int64Value(request.ReasonCode),
			"reason_msg":        tea.StringValue(request.ReasonMsg),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RoundCompleteUploadUserResult(request *RoundCompleteUploadUserResultRequest) (_result *RoundCompleteUploadUserResultResponse, _err error) {
	return invoke[RoundCompleteUploadUserResultResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/complete_upload_user_result",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_open_id": tea.StringValue(request.AnchorOpenId),
			"app_id":         tea.StringValue(request.AppId),
			"complete_time":  tea.Int64Value(request.CompleteTime),
			"room_id":        tea.StringValue(request.RoomId),
			"round_id":       tea.Int64Value(request.RoundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RoundInfo(request *RoundInfoRequest) (_result *RoundInfoResponse, _err error) {
	return invoke[RoundInfoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/business/round/info",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"anchor_open_id":  tea.StringValue(request.AnchorOpenId),
			"app_id":          tea.StringValue(request.AppId),
			"end_status":      tea.Int64Value(request.EndStatus),
			"end_time":        tea.Int64Value(request.EndTime),
			"end_user_list":   request.EndUserList,
			"mvp_list":        request.MvpList,
			"room_id":         tea.StringValue(request.RoomId),
			"round_id":        tea.StringValue(request.RoundId),
			"start_time":      tea.Int64Value(request.StartTime),
			"start_user_list": request.StartUserList,
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RoundSyncStatus(request *RoundSyncStatusRequest) (_result *RoundSyncStatusResponse, _err error) {
	return invoke[RoundSyncStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/sync_status",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_open_id":    tea.StringValue(request.AnchorOpenId),
			"app_id":            tea.StringValue(request.AppId),
			"end_time":          tea.Int64Value(request.EndTime),
			"group_result_list": request.GroupResultList,
			"room_id":           tea.StringValue(request.RoomId),
			"round_id":          tea.Int64Value(request.RoundId),
			"start_time":        tea.Int64Value(request.StartTime),
			"status":            tea.IntValue(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RoundUploadRankList(request *RoundUploadRankListRequest) (_result *RoundUploadRankListResponse, _err error) {
	return invoke[RoundUploadRankListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/upload_rank_list",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_open_id": tea.StringValue(request.AnchorOpenId),
			"app_id":         tea.StringValue(request.AppId),
			"rank_list":      request.RankList,
			"room_id":        tea.StringValue(request.RoomId),
			"round_id":       tea.Int64Value(request.RoundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RoundUploadUserResult(request *RoundUploadUserResultRequest) (_result *RoundUploadUserResultResponse, _err error) {
	return invoke[RoundUploadUserResultResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/upload_user_result",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_open_id": tea.StringValue(request.AnchorOpenId),
			"app_id":         tea.StringValue(request.AppId),
			"room_id":        tea.StringValue(request.RoomId),
			"round_id":       tea.Int64Value(request.RoundId),
			"user_list":      request.UserList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RtEcpmQuery(request *RtEcpmQueryRequest) (_result *RtEcpmQueryResponse, _err error) {
	return invoke[RtEcpmQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/traffic/v2/rt_ecpm/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"cursor":     tea.StringValue(request.Cursor),
			"date_hour":  tea.StringValue(request.DateHour),
			"end_date":   tea.StringValue(request.EndDate),
			"open_id":    tea.StringValue(request.OpenId),
			"page_size":  tea.Int32Value(request.PageSize),
			"start_date": tea.StringValue(request.StartDate),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) RulePush(request *RulePushRequest) (_result *RulePushResponse, _err error) {
	return invoke[RulePushResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion/price/rule/push/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":             tea.StringValue(request.AccountId),
			"active":                 tea.BoolValue(request.Active),
			"applicable_date":        request.ApplicableDate,
			"applicable_resource":    request.ApplicableResource,
			"full_pattern_promotion": request.FullPatternPromotion,
			"is_auto_extension":      tea.BoolValue(request.IsAutoExtension),
			"promotion_basic_info":   request.PromotionBasicInfo,
			"promotion_id":           tea.StringValue(request.PromotionId),
			"unapplicable_date":      request.UnapplicableDate,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasAddMerchant(request *SaasAddMerchantRequest) (_result *SaasAddMerchantResponse, _err error) {
	return invoke[SaasAddMerchantResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/add_merchant/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"prod_id":                 tea.Int32Value(request.ProdId),
			"thirdparty_component_id": tea.StringValue(request.ThirdpartyComponentId),
			"url_type":                tea.IntValue(request.UrlType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasAddSubMerchant(request *SaasAddSubMerchantRequest) (_result *SaasAddSubMerchantResponse, _err error) {
	return invoke[SaasAddSubMerchantResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/add_sub_merchant/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"prod_id":         tea.Int32Value(request.ProdId),
			"role":            tea.Int32Value(request.Role),
			"sub_merchant_id": tea.StringValue(request.SubMerchantId),
			"thirdparty_id":   tea.StringValue(request.ThirdpartyId),
			"url_type":        tea.IntValue(request.UrlType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasCreateMerchant(request *SaasCreateMerchantRequest) (_result *SaasCreateMerchantResponse, _err error) {
	return invoke[SaasCreateMerchantResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/create_merchant/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":                  tea.StringValue(request.AppId),
			"beneficiary":             request.Beneficiary,
			"beneficiary_type":        tea.StringValue(request.BeneficiaryType),
			"business_license":        request.BusinessLicense,
			"callback_url":            tea.StringValue(request.CallbackUrl),
			"channels":                tea.StringSliceValue(request.Channels),
			"city_code":               tea.StringValue(request.CityCode),
			"create_name":             tea.StringValue(request.CreateName),
			"district_code":           tea.StringValue(request.DistrictCode),
			"ext_evidences":           request.ExtEvidences,
			"industry_code":           tea.StringSliceValue(request.IndustryCode),
			"industry_info_pic_urls":  request.IndustryInfoPicUrls,
			"legal_person":            request.LegalPerson,
			"merchant_card_info":      request.MerchantCardInfo,
			"merchant_name":           tea.StringValue(request.MerchantName),
			"merchant_operation_info": request.MerchantOperationInfo,
			"merchant_short_name":     tea.StringValue(request.MerchantShortName),
			"merchant_type":           tea.Int64Value(request.MerchantType),
			"out_order_id":            tea.StringValue(request.OutOrderId),
			"province_code":           tea.StringValue(request.ProvinceCode),
			"registered_addr":         tea.StringValue(request.RegisteredAddr),
			"sub_merchant_id":         tea.StringValue(request.SubMerchantId),
			"thirdparty_id":           tea.StringValue(request.ThirdpartyId),
			"type":                    tea.Int64Value(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasGetAppMerchant(request *SaasGetAppMerchantRequest) (_result *SaasGetAppMerchantResponse, _err error) {
	return invoke[SaasGetAppMerchantResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/get_app_merchant/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":        tea.StringValue(request.AppId),
			"thirdparty_id": tea.StringValue(request.ThirdpartyId),
			"url_type":      tea.IntValue(request.UrlType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasMerchantWithdraw(request *SaasMerchantWithdrawRequest) (_result *SaasMerchantWithdrawResponse, _err error) {
	return invoke[SaasMerchantWithdrawResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/saas/merchant_withdraw/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":          tea.StringValue(request.AppId),
			"callback":        tea.StringValue(request.Callback),
			"channel_type":    tea.StringValue(request.ChannelType),
			"cp_extra":        tea.StringValue(request.CpExtra),
			"merchant_entity": tea.Int32Value(request.MerchantEntity),
			"merchant_uid":    tea.StringValue(request.MerchantUid),
			"out_order_id":    tea.StringValue(request.OutOrderId),
			"thirdparty_id":   tea.StringValue(request.ThirdpartyId),
			"withdraw_amount": tea.Int64Value(request.WithdrawAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasQueryMerchantBalance(request *SaasQueryMerchantBalanceRequest) (_result *SaasQueryMerchantBalanceResponse, _err error) {
	return invoke[SaasQueryMerchantBalanceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/saas/query_merchant_balance/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":          tea.StringValue(request.AppId),
			"channel_type":    tea.StringValue(request.ChannelType),
			"merchant_entity": tea.Int32Value(request.MerchantEntity),
			"merchant_uid":    tea.StringValue(request.MerchantUid),
			"thirdparty_id":   tea.StringValue(request.ThirdpartyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasQueryMerchantStatus(request *SaasQueryMerchantStatusRequest) (_result *SaasQueryMerchantStatusResponse, _err error) {
	return invoke[SaasQueryMerchantStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/query_merchant_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":          tea.StringValue(request.AppId),
			"merchant_id":     tea.StringValue(request.MerchantId),
			"sub_merchant_id": tea.StringValue(request.SubMerchantId),
			"thirdparty_id":   tea.StringValue(request.ThirdpartyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaasQueryWithdrawOrder(request *SaasQueryWithdrawOrderRequest) (_result *SaasQueryWithdrawOrderResponse, _err error) {
	return invoke[SaasQueryWithdrawOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/saas/query_withdraw_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":        tea.StringValue(request.AppId),
			"channel_type":  tea.StringValue(request.ChannelType),
			"merchant_uid":  tea.StringValue(request.MerchantUid),
			"out_order_id":  tea.StringValue(request.OutOrderId),
			"thirdparty_id": tea.StringValue(request.ThirdpartyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaveLiveOrientedPlan(request *SaveLiveOrientedPlanRequest) (_result *SaveLiveOrientedPlanResponse, _err error) {
	return invoke[SaveLiveOrientedPlanResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/save_live_oriented_plan/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"douyin_id_list": tea.StringSliceValue(request.DouyinIdList),
			"merchant_phone": tea.StringValue(request.MerchantPhone),
			"plan_id":        tea.Int64Value(request.PlanId),
			"plan_name":      tea.StringValue(request.PlanName),
			"product_list":   request.ProductList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SavePresaleAri(request *SavePresaleAriRequest) (_result *SavePresaleAriResponse, _err error) {
	return invoke[SavePresaleAriResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/presale_coupon/save/presale/ari/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":           tea.StringValue(request.AccountId),
			"actual_amount":        tea.Int64Value(request.ActualAmount),
			"children_definition":  request.ChildrenDefinition,
			"date_add_price_rule":  request.DateAddPriceRule,
			"marketing_amount":     tea.Int64Value(request.MarketingAmount),
			"origin_amount":        tea.Int64Value(request.OriginAmount),
			"product_id":           tea.StringValue(request.ProductId),
			"room_add_price_rule":  request.RoomAddPriceRule,
			"stock_qty_limit_type": tea.Int32Value(request.StockQtyLimitType),
			"total_stock_qty":      tea.Int64Value(request.TotalStockQty),
			"user_add_price_rule":  request.UserAddPriceRule,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaveRetainConsultCard(request *SaveRetainConsultCardRequest) (_result *SaveRetainConsultCardResponse, _err error) {
	return invoke[SaveRetainConsultCardResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/save/retain_consult_card/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"card_id":    tea.StringValue(request.CardId),
			"components": tea.IntValueSlice(request.Components),
			"media_id":   tea.StringValue(request.MediaId),
			"title":      tea.StringValue(request.Title),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SaveVideoOrientedPlan(request *SaveVideoOrientedPlanRequest) (_result *SaveVideoOrientedPlanResponse, _err error) {
	return invoke[SaveVideoOrientedPlanResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/save_video_oriented_plan/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"commission_duration": tea.Int64Value(request.CommissionDuration),
			"douyin_id_list":      tea.StringSliceValue(request.DouyinIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"merchant_phone":      tea.StringValue(request.MerchantPhone),
			"plan_id":             tea.Int64Value(request.PlanId),
			"plan_name":           tea.StringValue(request.PlanName),
			"product_list":        request.ProductList,
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SchemaGenerate(request *SchemaGenerateRequest) (_result *SchemaGenerateResponse, _err error) {
	return invoke[SchemaGenerateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/schema/generate",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"expire_time":  tea.Int64Value(request.ExpireTime),
			"no_expire":    tea.BoolValue(request.NoExpire),
			"query":        tea.StringValue(request.Query),
			"version_type": tea.StringValue(request.VersionType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SchemaGetItemInfo(request *SchemaGetItemInfoRequest) (_result *SchemaGetItemInfoResponse, _err error) {
	return invoke[SchemaGetItemInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/schema/get_item_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"expire_at": tea.Int64Value(request.ExpireAt),
			"item_id":   tea.StringValue(request.ItemId),
			"video_id":  tea.Int64Value(request.VideoId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SchemaQueryInfo(request *SchemaQueryInfoRequest) (_result *SchemaQueryInfoResponse, _err error) {
	return invoke[SchemaQueryInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/schema/query_info",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"schema": tea.StringValue(request.Schema),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SchemaQueryQuota(request *SchemaQueryQuotaRequest) (_result *SchemaQueryQuotaResponse, _err error) {
	return invoke[SchemaQueryQuotaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/schema/query_quota",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ScoreQuery(request *ScoreQueryRequest) (_result *ScoreQueryResponse, _err error) {
	return invoke[ScoreQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/comment/score/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":  tea.StringValue(request.AccountId),
			"poi_id_list": tea.Int64ValueSlice(request.PoiIdList),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchCheckSubService(request *SearchCheckSubServiceRequest) (_result *SearchCheckSubServiceResponse, _err error) {
	return invoke[SearchCheckSubServiceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/search/check_sub_service/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchDeleteIndex(request *SearchDeleteIndexRequest) (_result *SearchDeleteIndexResponse, _err error) {
	return invoke[SearchDeleteIndexResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/search/delete_index/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":    tea.StringValue(request.AppId),
			"name":      tea.StringValue(request.Name),
			"path_list": tea.StringSliceValue(request.PathList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchDeleteSubService(request *SearchDeleteSubServiceRequest) (_result *SearchDeleteSubServiceResponse, _err error) {
	return invoke[SearchDeleteSubServiceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/search/delete_sub_service/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"sub_service_id": tea.StringValue(request.SubServiceId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchExperience(request *SearchExperienceRequest) (_result *SearchExperienceResponse, _err error) {
	return invoke[SearchExperienceResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/dy_open_api/v2/search/experience/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"content_type": tea.Int32Value(request.ContentType),
			"count":        tea.Int32Value(request.Count),
			"cursor":       tea.Int64Value(request.Cursor),
			"device_id":    tea.Int64Value(request.DeviceId),
			"keyword":      tea.StringValue(request.Keyword),
			"search_id":    tea.StringValue(request.SearchId),
			"sort_type":    tea.Int32Value(request.SortType),
			"uid":          tea.StringValue(request.Uid),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchKeyword(request *SearchKeywordRequest) (_result *SearchKeywordResponse, _err error) {
	return invoke[SearchKeywordResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/poi/search/keyword/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"city":    tea.StringValue(request.City),
			"count":   tea.Int64Value(request.Count),
			"cursor":  tea.Int64Value(request.Cursor),
			"keyword": tea.StringValue(request.Keyword),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchUploadSitemap(request *SearchUploadSitemapRequest) (_result *SearchUploadSitemapResponse, _err error) {
	return invoke[SearchUploadSitemapResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/search/upload_sitemap/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":     tea.StringValue(request.AppId),
			"page_paths": tea.StringSliceValue(request.PagePaths),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SearchVideo(request *SearchVideoRequest) (_result *SearchVideoResponse, _err error) {
	return invoke[SearchVideoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/dy_open_api/v2/search/video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"count":        tea.Int32Value(request.Count),
			"cursor":       tea.Int64Value(request.Cursor),
			"device_id":    tea.Int64Value(request.DeviceId),
			"keyword":      tea.StringValue(request.Keyword),
			"open_id":      tea.StringValue(request.OpenId),
			"publish_time": tea.Int32Value(request.PublishTime),
			"search_id":    tea.StringValue(request.SearchId),
			"sort_type":    tea.Int32Value(request.SortType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SendCouponToDesignatedUser(request *SendCouponToDesignatedUserRequest) (_result *SendCouponToDesignatedUserResponse, _err error) {
	return invoke[SendCouponToDesignatedUserResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/send_coupon_to_designated_user/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id":        tea.StringValue(request.ActivityId),
			"app_id":             tea.StringValue(request.AppId),
			"merchant_coupon_id": tea.StringValue(request.MerchantCouponId),
			"open_id":            tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SendMsg(request *SendMsgRequest) (_result *SendMsgResponse, _err error) {
	return invoke[SendMsgResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/send/msg/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"content":         request.Content,
			"content_list":    request.ContentList,
			"conversation_id": tea.StringValue(request.ConversationId),
			"conversion_id":   tea.StringValue(request.ConversionId),
			"msg_id":          tea.StringValue(request.MsgId),
			"scene":           tea.StringValue(request.Scene),
			"to_user_id":      tea.StringValue(request.ToUserId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ServeQuery(request *ServeQueryRequest) (_result *ServeQueryResponse, _err error) {
	return invoke[ServeQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/order/serve/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
		},
		Body: map[string]interface{}{
			"param": request.Param,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ServeSubmit(request *ServeSubmitRequest) (_result *ServeSubmitResponse, _err error) {
	return invoke[ServeSubmitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/order/serve/submit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
		},
		Body: map[string]interface{}{
			"param": request.Param,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SetUserGroupTag(request *SetUserGroupTagRequest) (_result *SetUserGroupTagResponse, _err error) {
	return invoke[SetUserGroupTagResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/group_tag/set_user_group_tag",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"mp_id":   tea.StringValue(request.MpId),
			"open_id": tea.StringValue(request.OpenId),
			"status":  tea.Int32Value(request.Status),
			"tag_id":  tea.StringValue(request.TagId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SettingDisable(request *SettingDisableRequest) (_result *SettingDisableResponse, _err error) {
	return invoke[SettingDisableResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/group/setting/disable/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"group_id":           tea.StringValue(request.GroupId),
			"group_setting_type": tea.IntValue(request.GroupSettingType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SettingSet(request *SettingSetRequest) (_result *SettingSetResponse, _err error) {
	return invoke[SettingSetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/group/setting/set/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"group_setting_type": tea.IntValue(request.GroupSettingType),
			"msg_list":           request.MsgList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShareCreateActivityId(request *ShareCreateActivityIdRequest) (_result *ShareCreateActivityIdResponse, _err error) {
	return invoke[ShareCreateActivityIdResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/mgplatform/api/apps/share/create_activity_id",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base": request.Base,
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShareCreateTask(request *ShareCreateTaskRequest) (_result *ShareCreateTaskResponse, _err error) {
	return invoke[ShareCreateTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/share/create_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"start_time":   tea.Int64Value(request.StartTime),
			"target_count": tea.Int64Value(request.TargetCount),
			"task_type":    tea.Int32Value(request.TaskType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShareQueryUserTask(request *ShareQueryUserTaskRequest) (_result *ShareQueryUserTaskResponse, _err error) {
	return invoke[ShareQueryUserTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/share/query_user_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
			"task_id": tea.StringValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShareUnbindUnionGroup(request *ShareUnbindUnionGroupRequest) (_result *ShareUnbindUnionGroupResponse, _err error) {
	return invoke[ShareUnbindUnionGroupResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/share/unbind_union_group",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"mp_id":    tea.StringValue(request.MpId),
			"union_id": tea.StringValue(request.UnionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShareUpdateDynamicMessage(request *ShareUpdateDynamicMessageRequest) (_result *ShareUpdateDynamicMessageResponse, _err error) {
	return invoke[ShareUpdateDynamicMessageResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/share/update_dynamic_message",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id":   tea.StringValue(request.ActivityId),
			"query":         tea.StringValue(request.Query),
			"target_state":  tea.Int32Value(request.TargetState),
			"template_info": request.TemplateInfo,
			"version_type":  tea.StringValue(request.VersionType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) Shareid(request *ShareidRequest) (_result *ShareidResponse, _err error) {
	return invoke[ShareidResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/share-id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"default_hashtag": tea.StringValue(request.DefaultHashtag),
			"link_param":      tea.StringValue(request.LinkParam),
			"need_callback":   tea.BoolValue(request.NeedCallback),
			"source_style_id": tea.StringValue(request.SourceStyleId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShopMemberLeave(request *ShopMemberLeaveRequest) (_result *ShopMemberLeaveResponse, _err error) {
	return invoke[ShopMemberLeaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/ecom/v1/shop_member/leave/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":  tea.StringValue(request.AppId),
			"open_id": tea.StringValue(request.OpenId),
			"shop_id": tea.Int64Value(request.ShopId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ShopPoiQuery(request *ShopPoiQueryRequest) (_result *ShopPoiQueryResponse, _err error) {
	return invoke[ShopPoiQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/shop/poi/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":    tea.StringValue(request.AccountId),
			"page":          tea.Int32Value(request.Page),
			"poi_id":        tea.StringValue(request.PoiId),
			"relation_type": tea.IntValue(request.RelationType),
			"size":          tea.Int32Value(request.Size),
			"third_id":      tea.StringValue(request.ThirdId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SkuGet(request *SkuGetRequest) (_result *SkuGetResponse, _err error) {
	return invoke[SkuGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/sku/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":           request.Base,
			"account_id":     tea.StringValue(request.AccountId),
			"out_sku_ids":    tea.StringSliceValue(request.OutSkuIds),
			"product_out_id": tea.StringValue(request.ProductOutId),
			"sku_ids":        tea.StringSliceValue(request.SkuIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SkuUpsert(request *SkuUpsertRequest) (_result *SkuUpsertResponse, _err error) {
	return invoke[SkuUpsertResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/comprehensive/reception/stock/sku/upsert/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
		},
		Body: map[string]interface{}{
			"poi_id":        tea.StringValue(request.PoiId),
			"sku_info_list": request.SkuInfoList,
			"time_slot":     tea.Int64Value(request.TimeSlot),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportBasketball(request *SportBasketballRequest) (_result *SportBasketballResponse, _err error) {
	return invoke[SportBasketballResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/basketball/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportComprehensive(request *SportComprehensiveRequest) (_result *SportComprehensiveResponse, _err error) {
	return invoke[SportComprehensiveResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/comprehensive/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportCulture(request *SportCultureRequest) (_result *SportCultureResponse, _err error) {
	return invoke[SportCultureResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/culture/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportFitness(request *SportFitnessRequest) (_result *SportFitnessResponse, _err error) {
	return invoke[SportFitnessResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/fitness/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportOutdoors(request *SportOutdoorsRequest) (_result *SportOutdoorsResponse, _err error) {
	return invoke[SportOutdoorsResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/outdoors/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportOverall(request *SportOverallRequest) (_result *SportOverallResponse, _err error) {
	return invoke[SportOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportSoccer(request *SportSoccerRequest) (_result *SportSoccerResponse, _err error) {
	return invoke[SportSoccerResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/soccer/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SportTableTennis(request *SportTableTennisRequest) (_result *SportTableTennisResponse, _err error) {
	return invoke[SportTableTennisResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/sport/table_tennis/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SpuDetail(request *SpuDetailRequest) (_result *SpuDetailResponse, _err error) {
	return invoke[SpuDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/spu/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.Int64Value(request.AccountId),
			"out_spu_id": tea.StringValue(request.OutSpuId),
			"spu_id":     tea.Int64Value(request.SpuId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SpuGet(request *SpuGetRequest) (_result *SpuGetResponse, _err error) {
	return invoke[SpuGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v2/goods/spu/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
			"out_spu_id": tea.StringValue(request.OutSpuId),
			"spu_id":     tea.Int64Value(request.SpuId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SpuOperate(request *SpuOperateRequest) (_result *SpuOperateResponse, _err error) {
	return invoke[SpuOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/spu/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.Int64Value(request.AccountId),
			"op":         tea.IntValue(request.Op),
			"out_spu_id": tea.StringValue(request.OutSpuId),
			"spu_id":     tea.Int64Value(request.SpuId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SpuSave(request *SpuSaveRequest) (_result *SpuSaveResponse, _err error) {
	return invoke[SpuSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/spu/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.Int64Value(request.AccountId),
			"spu":        request.Spu,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) StatusQuery(request *StatusQueryRequest) (_result *StatusQueryResponse, _err error) {
	return invoke[StatusQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/status/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":        tea.StringValue(request.AccountId),
			"hotel_id_list":     tea.StringSliceValue(request.HotelIdList),
			"out_hotel_id_list": tea.StringSliceValue(request.OutHotelIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) StockSave(request *StockSaveRequest) (_result *StockSaveResponse, _err error) {
	return invoke[StockSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/stock/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"aris":       request.Aris,
			"hotel_id":   tea.StringValue(request.HotelId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) StockSync(request *StockSyncRequest) (_result *StockSyncResponse, _err error) {
	return invoke[StockSyncResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/life/goods/stock/sync/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.StringValue(request.ProductId),
			"stock":      request.Stock,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) StockUpdateNotify(request *StockUpdateNotifyRequest) (_result *StockUpdateNotifyResponse, _err error) {
	return invoke[StockUpdateNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/booking/stock_update/notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"date":       tea.StringValue(request.Date),
			"poi_id":     tea.StringValue(request.PoiId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionAddAppTpl(request *SubscriptionAddAppTplRequest) (_result *SubscriptionAddAppTplResponse, _err error) {
	return invoke[SubscriptionAddAppTplResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/notification/v2/subscription/add_app_tpl/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"ISVClientKey": tea.StringValue(request.ISVClientKey),
			"keyword_list": tea.StringSliceValue(request.KeywordList),
			"template_id":  tea.Int64Value(request.TemplateId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionCreateTpl(request *SubscriptionCreateTplRequest) (_result *SubscriptionCreateTplResponse, _err error) {
	return invoke[SubscriptionCreateTplResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/notification/v2/subscription/create_tpl/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"category_ids":   tea.StringValue(request.CategoryIds),
			"classification": tea.IntValue(request.Classification),
			"host_list":      tea.StringSliceValue(request.HostList),
			"keyword_list":   tea.StringSliceValue(request.KeywordList),
			"title":          tea.StringValue(request.Title),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionDeleteAppTpl(request *SubscriptionDeleteAppTplRequest) (_result *SubscriptionDeleteAppTplResponse, _err error) {
	return invoke[SubscriptionDeleteAppTplResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/notification/v2/subscription/delete_app_tpl/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"msg_id": tea.StringValue(request.MsgId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionNotifyUser(request *SubscriptionNotifyUserRequest) (_result *SubscriptionNotifyUserResponse, _err error) {
	return invoke[SubscriptionNotifyUserResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/notification/v2/subscription/notify_user/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"data":        request.Data,
			"msg_id":      tea.StringValue(request.MsgId),
			"notify_type": tea.IntValueSlice(request.NotifyType),
			"open_id":     tea.StringValue(request.OpenId),
			"page":        tea.StringValue(request.Page),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionQueryAppTpl(request *SubscriptionQueryAppTplRequest) (_result *SubscriptionQueryAppTplResponse, _err error) {
	return invoke[SubscriptionQueryAppTplResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/notification/v2/subscription/query_app_tpl/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"category_ids":   tea.StringValue(request.CategoryIds),
			"classification": tea.IntValue(request.Classification),
			"page_num":       tea.Int32Value(request.PageNum),
			"page_size":      tea.Int32Value(request.PageSize),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionQueryTplList(request *SubscriptionQueryTplListRequest) (_result *SubscriptionQueryTplListResponse, _err error) {
	return invoke[SubscriptionQueryTplListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/notification/v2/subscription/query_tpl_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"category_ids":   tea.StringValue(request.CategoryIds),
			"classification": tea.IntValue(request.Classification),
			"keyword":        tea.StringValue(request.Keyword),
			"page_num":       tea.Int32Value(request.PageNum),
			"page_size":      tea.Int32Value(request.PageSize),
			"template_type":  tea.IntValue(request.TemplateType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SubscriptionQueryUserSubscribe(request *SubscriptionQueryUserSubscribeRequest) (_result *SubscriptionQueryUserSubscribeResponse, _err error) {
	return invoke[SubscriptionQueryUserSubscribeResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/notification/v2/subscription/query_user_subscribe/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"msg_id":  tea.StringValue(request.MsgId),
			"open_id": tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) SyncStatus(request *SyncStatusRequest) (_result *SyncStatusResponse, _err error) {
	return invoke[SyncStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/retail/order/sync/status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":         tea.StringValue(request.OrderId),
			"order_out_status": tea.IntValue(request.OrderOutStatus),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskCreate(request *TaskCreateRequest) (_result *TaskCreateResponse, _err error) {
	return invoke[TaskCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/match/task/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":            tea.StringValue(request.AccountId),
			"additional_image_urls": tea.StringSliceValue(request.AdditionalImageUrls),
			"additional_info":       tea.StringValue(request.AdditionalInfo),
			"address":               tea.StringValue(request.Address),
			"city":                  tea.StringValue(request.City),
			"ext_id":                tea.StringValue(request.ExtId),
			"head_image_urls":       tea.StringSliceValue(request.HeadImageUrls),
			"last_task_id":          tea.StringValue(request.LastTaskId),
			"latitude":              tea.StringValue(request.Latitude),
			"longitude":             tea.StringValue(request.Longitude),
			"open_status":           tea.Int32Value(request.OpenStatus),
			"open_times":            request.OpenTimes,
			"poi_name":              tea.StringValue(request.PoiName),
			"province":              tea.StringValue(request.Province),
			"tel_list":              tea.StringSliceValue(request.TelList),
			"type_code":             tea.StringValue(request.TypeCode),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskCreateVideo(request *TaskCreateVideoRequest) (_result *TaskCreateVideoResponse, _err error) {
	return invoke[TaskCreateVideoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/task/create_video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"conditions": tea.StringSliceValue(request.Conditions),
			"end_time":   tea.Int64Value(request.EndTime),
			"item_id":    tea.StringValue(request.ItemId),
			"start_time": tea.Int64Value(request.StartTime),
			"task_name":  tea.StringValue(request.TaskName),
			"video_url":  tea.StringValue(request.VideoUrl),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskGet(request *TaskGetRequest) (_result *TaskGetResponse, _err error) {
	return invoke[TaskGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/live_data/task/get",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"appid":    tea.StringValue(request.Appid),
			"msg_type": tea.StringValue(request.MsgType),
			"roomid":   tea.StringValue(request.Roomid),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskQuery(request *TaskQueryRequest) (_result *TaskQueryResponse, _err error) {
	return invoke[TaskQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/booking/config/task/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"task_id":    tea.StringValue(request.TaskId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskStart(request *TaskStartRequest) (_result *TaskStartResponse, _err error) {
	return invoke[TaskStartResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/live_data/task/start",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":    tea.StringValue(request.Appid),
			"msg_type": tea.StringValue(request.MsgType),
			"roomid":   tea.StringValue(request.Roomid),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskSubmit(request *TaskSubmitRequest) (_result *TaskSubmitResponse, _err error) {
	return invoke[TaskSubmitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/match/task/submit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"datas": request.Datas,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskView(request *TaskViewRequest) (_result *TaskViewResponse, _err error) {
	return invoke[TaskViewResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/match/task/view/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"ext_id":     tea.StringValue(request.ExtId),
			"task_id":    tea.StringValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskWriteoffLive(request *TaskWriteoffLiveRequest) (_result *TaskWriteoffLiveResponse, _err error) {
	return invoke[TaskWriteoffLiveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/task/writeoff_live/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
			"task_id": tea.StringValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskWriteoffVideo(request *TaskWriteoffVideoRequest) (_result *TaskWriteoffVideoResponse, _err error) {
	return invoke[TaskWriteoffVideoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/task/writeoff_video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
			"task_id": tea.StringValue(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxAddRoomTask(request *TaskboxAddRoomTaskRequest) (_result *TaskboxAddRoomTaskResponse, _err error) {
	return invoke[TaskboxAddRoomTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/add_room_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"content_tag":       tea.StringValue(request.ContentTag),
			"form_tag":          tea.StringValue(request.FormTag),
			"refer_ma_captures": tea.StringSliceValue(request.ReferMaCaptures),
			"refer_videos":      tea.StringSliceValue(request.ReferVideos),
			"room_title":        tea.StringValue(request.RoomTitle),
			"start_page":        tea.StringValue(request.StartPage),
			"task_desc":         tea.StringValue(request.TaskDesc),
			"task_end_time":     tea.Int64Value(request.TaskEndTime),
			"task_icon":         tea.StringValue(request.TaskIcon),
			"task_name":         tea.StringValue(request.TaskName),
			"task_settle_type":  tea.Int32Value(request.TaskSettleType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxAddTask(request *TaskboxAddTaskRequest) (_result *TaskboxAddTaskResponse, _err error) {
	return invoke[TaskboxAddTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/add_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_title":                      tea.StringValue(request.AnchorTitle),
			"douyin_ids":                        tea.StringSliceValue(request.DouyinIds),
			"mix_payment_allocate_ratio":        request.MixPaymentAllocateRatio,
			"page_type":                         tea.IntValue(request.PageType),
			"payment_allocate_ratio":            tea.Int32Value(request.PaymentAllocateRatio),
			"refer_gids":                        tea.Int64ValueSlice(request.ReferGids),
			"refer_ma_captures":                 tea.StringSliceValue(request.ReferMaCaptures),
			"refer_videos":                      tea.StringSliceValue(request.ReferVideos),
			"start_page":                        tea.StringValue(request.StartPage),
			"talent_mix_payment_allocate_ratio": request.TalentMixPaymentAllocateRatio,
			"talent_payment_allocate_ratio":     tea.Int32Value(request.TalentPaymentAllocateRatio),
			"task_desc":                         tea.StringValue(request.TaskDesc),
			"task_end_time":                     tea.Int64Value(request.TaskEndTime),
			"task_icon":                         tea.StringValue(request.TaskIcon),
			"task_name":                         tea.StringValue(request.TaskName),
			"task_refund_period":                tea.Int32Value(request.TaskRefundPeriod),
			"task_settle_type":                  tea.Int32Value(request.TaskSettleType),
			"task_start_time":                   tea.Int64Value(request.TaskStartTime),
			"task_tags":                         tea.StringSliceValue(request.TaskTags),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxGenAgentLink(request *TaskboxGenAgentLinkRequest) (_result *TaskboxGenAgentLinkResponse, _err error) {
	return invoke[TaskboxGenAgentLinkResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/gen_agent_link/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agency_talent_uid": tea.StringValue(request.AgencyTalentUid),
			"agent_id":          tea.Int64Value(request.AgentId),
			"app_id":            tea.StringValue(request.AppId),
			"task_category":     tea.IntValue(request.TaskCategory),
			"task_id":           tea.Int64Value(request.TaskId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxQueryAppTaskId(request *TaskboxQueryAppTaskIdRequest) (_result *TaskboxQueryAppTaskIdResponse, _err error) {
	return invoke[TaskboxQueryAppTaskIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_app_task_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"create_end_time":   tea.Int64Value(request.CreateEndTime),
			"create_start_time": tea.Int64Value(request.CreateStartTime),
			"task_category":     tea.IntValue(request.TaskCategory),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxQueryBillLink(request *TaskboxQueryBillLinkRequest) (_result *TaskboxQueryBillLinkResponse, _err error) {
	return invoke[TaskboxQueryBillLinkResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_bill_link/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"bill_date": tea.StringValue(request.BillDate),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxQueryTaskInfo(request *TaskboxQueryTaskInfoRequest) (_result *TaskboxQueryTaskInfoResponse, _err error) {
	return invoke[TaskboxQueryTaskInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/query_task_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":                tea.StringValue(request.Appid),
			"page_no":              tea.Int32Value(request.PageNo),
			"page_size":            tea.Int32Value(request.PageSize),
			"query_params_content": tea.StringValue(request.QueryParamsContent),
			"query_params_type":    tea.IntValue(request.QueryParamsType),
			"task_category":        tea.IntValue(request.TaskCategory),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxSaveAgent(request *TaskboxSaveAgentRequest) (_result *TaskboxSaveAgentResponse, _err error) {
	return invoke[TaskboxSaveAgentResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/save_agent/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agent_id":       tea.Int64Value(request.AgentId),
			"agent_nickname": tea.StringValue(request.AgentNickname),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxUpdateStatus(request *TaskboxUpdateStatusRequest) (_result *TaskboxUpdateStatusResponse, _err error) {
	return invoke[TaskboxUpdateStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/update_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"task_id":     tea.Int64Value(request.TaskId),
			"task_status": tea.Int64Value(request.TaskStatus),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TaskboxUpdateTask(request *TaskboxUpdateTaskRequest) (_result *TaskboxUpdateTaskResponse, _err error) {
	return invoke[TaskboxUpdateTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/update_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_title":      tea.StringValue(request.AnchorTitle),
			"page_type":         tea.IntValue(request.PageType),
			"refer_gids":        tea.Int64ValueSlice(request.ReferGids),
			"refer_ma_captures": tea.StringSliceValue(request.ReferMaCaptures),
			"refer_videos":      tea.StringSliceValue(request.ReferVideos),
			"start_page":        tea.StringValue(request.StartPage),
			"task_desc":         tea.StringValue(request.TaskDesc),
			"task_end_time":     tea.Int64Value(request.TaskEndTime),
			"task_icon":         tea.StringValue(request.TaskIcon),
			"task_id":           tea.Int64Value(request.TaskId),
			"task_name":         tea.StringValue(request.TaskName),
			"task_settle_type":  tea.Int32Value(request.TaskSettleType),
			"task_start_time":   tea.Int64Value(request.TaskStartTime),
			"task_tags":         tea.StringSliceValue(request.TaskTags),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TemplateGet(request *TemplateGetRequest) (_result *TemplateGetResponse, _err error) {
	return invoke[TemplateGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/template/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":                request.Base,
			"account_id":          tea.StringValue(request.AccountId),
			"category_id":         tea.StringValue(request.CategoryId),
			"goods_template_type": tea.IntValue(request.GoodsTemplateType),
			"product_type":        tea.IntValue(request.ProductType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TemplateList(request *TemplateListRequest) (_result *TemplateListResponse, _err error) {
	return invoke[TemplateListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/template/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":        request.Base,
			"page_number": tea.Int64Value(request.PageNumber),
			"page_size":   tea.Int64Value(request.PageSize),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TicketCalendarBatchSave(request *TicketCalendarBatchSaveRequest) (_result *TicketCalendarBatchSaveResponse, _err error) {
	return invoke[TicketCalendarBatchSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/ticket_calendar/batch_save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"product_id":              tea.StringValue(request.ProductId),
			"product_out_id":          tea.StringValue(request.ProductOutId),
			"ticket_sku_and_calendar": request.TicketSkuAndCalendar,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TicketCalendarSave(request *TicketCalendarSaveRequest) (_result *TicketCalendarSaveResponse, _err error) {
	return invoke[TicketCalendarSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/ticket_calendar/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":           tea.StringValue(request.AccountId),
			"calendars":            request.Calendars,
			"product_id":           tea.StringValue(request.ProductId),
			"product_out_id":       tea.StringValue(request.ProductOutId),
			"ticket_specification": request.TicketSpecification,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TicketQuery(request *TicketQueryRequest) (_result *TicketQueryResponse, _err error) {
	return invoke[TicketQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/claim/appeal/ticket/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
		},
		Body: map[string]interface{}{
			"param": request.Param,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TicketSave(request *TicketSaveRequest) (_result *TicketSaveResponse, _err error) {
	return invoke[TicketSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/ticket/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"ticket":     request.Ticket,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TicketSubmit(request *TicketSubmitRequest) (_result *TicketSubmitResponse, _err error) {
	return invoke[TicketSubmitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/claim/appeal/ticket/submit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
		},
		Body: map[string]interface{}{
			"data": request.Data,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitButtonWhiteSetting(request *ToolkitButtonWhiteSettingRequest) (_result *ToolkitButtonWhiteSettingResponse, _err error) {
	return invoke[ToolkitButtonWhiteSettingResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/button_white_setting/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"button_type": tea.IntValue(request.ButtonType),
			"gray_rate":   tea.Int64Value(request.GrayRate),
			"open_all":    tea.BoolValue(request.OpenAll),
			"uid_list":    tea.Int64ValueSlice(request.UidList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitChangeLockStatus(request *ToolkitChangeLockStatusRequest) (_result *ToolkitChangeLockStatusResponse, _err error) {
	return invoke[ToolkitChangeLockStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/change_lock_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"operation_type":  tea.IntValue(request.OperationType),
			"order_info_list": request.OrderInfoList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitPushServiceDone(request *ToolkitPushServiceDoneRequest) (_result *ToolkitPushServiceDoneResponse, _err error) {
	return invoke[ToolkitPushServiceDoneResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/push_service_done/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":       tea.StringValue(request.OrderId),
			"verify_id_list": tea.StringSliceValue(request.VerifyIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitQueryCertificateInfo(request *ToolkitQueryCertificateInfoRequest) (_result *ToolkitQueryCertificateInfoResponse, _err error) {
	return invoke[ToolkitQueryCertificateInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/query_certificate_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificate_id_list": tea.StringSliceValue(request.CertificateIdList),
			"order_id_list":       tea.StringSliceValue(request.OrderIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitQueryText(request *ToolkitQueryTextRequest) (_result *ToolkitQueryTextResponse, _err error) {
	return invoke[ToolkitQueryTextResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/query_text/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"text_type": tea.IntValue(request.TextType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitUpdateMerchantConf(request *ToolkitUpdateMerchantConfRequest) (_result *ToolkitUpdateMerchantConfResponse, _err error) {
	return invoke[ToolkitUpdateMerchantConfResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/update_merchant_conf/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":               tea.StringValue(request.AccountId),
			"bind_biz_type":            tea.IntValue(request.BindBizType),
			"delivery_app_info":        request.DeliveryAppInfo,
			"product_double_open_info": request.ProductDoubleOpenInfo,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitUpdateMerchantPath(request *ToolkitUpdateMerchantPathRequest) (_result *ToolkitUpdateMerchantPathResponse, _err error) {
	return invoke[ToolkitUpdateMerchantPathResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/update_merchant_path/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"bind_biz_type":  tea.IntValue(request.BindBizType),
			"path_data_list": request.PathDataList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ToolkitVerifyLocalCertificates(request *ToolkitVerifyLocalCertificatesRequest) (_result *ToolkitVerifyLocalCertificatesResponse, _err error) {
	return invoke[ToolkitVerifyLocalCertificatesResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/toolkit/verify_local_certificates/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"open_id":            tea.StringValue(request.OpenId),
			"order_entry_schema": request.OrderEntrySchema,
			"order_info_list":    request.OrderInfoList,
			"poi_id":             tea.StringValue(request.PoiId),
			"verify_token":       tea.StringValue(request.VerifyToken),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TradeOrderConfirm(request *TradeOrderConfirmRequest) (_result *TradeOrderConfirmResponse, _err error) {
	return invoke[TradeOrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/comprehensive/trade/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"book_id":        tea.StringValue(request.BookId),
			"confirm_result": tea.IntValue(request.ConfirmResult),
			"fulfil_type":    tea.IntValue(request.FulfilType),
			"merchant_notes": tea.StringValue(request.MerchantNotes),
			"reason":         tea.StringValue(request.Reason),
			"reject_code":    tea.Int32Value(request.RejectCode),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TradeOrderQuery(request *TradeOrderQueryRequest) (_result *TradeOrderQueryResponse, _err error) {
	return invoke[TradeOrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/trade/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"create_order_end_time":   tea.Int64Value(request.CreateOrderEndTime),
			"create_order_start_time": tea.Int64Value(request.CreateOrderStartTime),
			"cursor":                  tea.StringSliceValue(request.Cursor),
			"ext_order_id":            tea.StringValue(request.ExtOrderId),
			"get_secret_number":       tea.BoolValue(request.GetSecretNumber),
			"open_id":                 tea.StringValue(request.OpenId),
			"order_id":                tea.StringValue(request.OrderId),
			"order_status":            tea.Int32Value(request.OrderStatus),
			"page_num":                tea.Int32Value(request.PageNum),
			"page_size":               tea.Int32Value(request.PageSize),
			"update_order_end_time":   tea.Int64Value(request.UpdateOrderEndTime),
			"update_order_start_time": tea.Int64Value(request.UpdateOrderStartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TrafficPermissionOpen(request *TrafficPermissionOpenRequest) (_result *TrafficPermissionOpenResponse, _err error) {
	return invoke[TrafficPermissionOpenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/traffic_permission/open/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"bank_account_number": tea.StringValue(request.BankAccountNumber),
			"bank_branch":         tea.StringValue(request.BankBranch),
			"bank_name":           tea.StringValue(request.BankName),
			"city":                tea.StringValue(request.City),
			"phone_number":        tea.StringValue(request.PhoneNumber),
			"province":            tea.StringValue(request.Province),
			"tax_nature":          tea.IntValue(request.TaxNature),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TrafficPermissionQuery(request *TrafficPermissionQueryRequest) (_result *TrafficPermissionQueryResponse, _err error) {
	return invoke[TrafficPermissionQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/traffic_permission/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TrafficVerify(request *TrafficVerifyRequest) (_result *TrafficVerifyResponse, _err error) {
	return invoke[TrafficVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/certificate/traffic/verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":     tea.StringValue(request.OrderId),
			"verify_mode":  tea.IntValue(request.VerifyMode),
			"verify_token": tea.StringValue(request.VerifyToken),
			"vouchers":     request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TransferCallback(request *TransferCallbackRequest) (_result *TransferCallbackResponse, _err error) {
	return invoke[TransferCallbackResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/transfer/callback/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificates": request.Certificates,
			"fail_reason":  tea.StringValue(request.FailReason),
			"order_id":     tea.StringValue(request.OrderId),
			"request_id":   tea.StringValue(request.RequestId),
			"result":       tea.Int64Value(request.Result),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TravelNew(request *TravelNewRequest) (_result *TravelNewResponse, _err error) {
	return invoke[TravelNewResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/travel/new/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TravelOverall(request *TravelOverallRequest) (_result *TravelOverallResponse, _err error) {
	return invoke[TravelOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/travel/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TravelagencyProductOperate(request *TravelagencyProductOperateRequest) (_result *TravelagencyProductOperateResponse, _err error) {
	return invoke[TravelagencyProductOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/product_operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
		},
		Body: map[string]interface{}{
			"opt_type":          tea.Int32Value(request.OptType),
			"product_info_list": request.ProductInfoList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TripCertificateVerify(request *TripCertificateVerifyRequest) (_result *TripCertificateVerifyResponse, _err error) {
	return invoke[TripCertificateVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/certificate/verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":     tea.StringValue(request.OrderId),
			"poi_id":       tea.StringValue(request.PoiId),
			"verify_time":  tea.Int64Value(request.VerifyTime),
			"verify_token": tea.StringValue(request.VerifyToken),
			"vouchers":     request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TripOrderConfirm(request *TripOrderConfirmRequest) (_result *TripOrderConfirmResponse, _err error) {
	return invoke[TripOrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"confirm_result": tea.IntValue(request.ConfirmResult),
			"order_id":       tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TripProductOperate(request *TripProductOperateRequest) (_result *TripProductOperateResponse, _err error) {
	return invoke[TripProductOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/product/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"op_type":         tea.IntValue(request.OpType),
			"product_id_list": tea.StringSliceValue(request.ProductIdList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TripRefundAudit(request *TripRefundAuditRequest) (_result *TripRefundAuditResponse, _err error) {
	return invoke[TripRefundAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/refund/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"audit_result":      tea.IntValue(request.AuditResult),
			"biz_uniq_key":      tea.StringValue(request.BizUniqKey),
			"order_id":          tea.StringValue(request.OrderId),
			"refund_fee_amount": tea.Int64Value(request.RefundFeeAmount),
			"vouchers":          request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) TripTicketQuery(request *TripTicketQueryRequest) (_result *TripTicketQueryResponse, _err error) {
	return invoke[TripTicketQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/trip/ticket/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"page":            tea.Int32Value(request.Page),
			"page_size":       tea.Int32Value(request.PageSize),
			"product_ids":     tea.StringSliceValue(request.ProductIds),
			"product_out_ids": tea.StringSliceValue(request.ProductOutIds),
			"status":          tea.IntValue(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateCommonPlanStatus(request *UpdateCommonPlanStatusRequest) (_result *UpdateCommonPlanStatusResponse, _err error) {
	return invoke[UpdateCommonPlanStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/update_common_plan_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"plan_update_list": request.PlanUpdateList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateCouponMetaStatus(request *UpdateCouponMetaStatusRequest) (_result *UpdateCouponMetaStatusResponse, _err error) {
	return invoke[UpdateCouponMetaStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/update_coupon_meta_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"coupon_meta_id":     tea.StringValue(request.CouponMetaId),
			"coupon_meta_status": tea.IntValue(request.CouponMetaStatus),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateCouponMetaStock(request *UpdateCouponMetaStockRequest) (_result *UpdateCouponMetaStockResponse, _err error) {
	return invoke[UpdateCouponMetaStockResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/update_coupon_meta_stock/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"action":         tea.StringValue(request.Action),
			"app_id":         tea.StringValue(request.AppId),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"number":         tea.Int64Value(request.Number),
			"unique_key":     tea.StringValue(request.UniqueKey),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateOrientedPlanStatus(request *UpdateOrientedPlanStatusRequest) (_result *UpdateOrientedPlanStatusResponse, _err error) {
	return invoke[UpdateOrientedPlanStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/update_oriented_plan_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"plan_update_list": request.PlanUpdateList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdatePromotionActivityStatus(request *UpdatePromotionActivityStatusRequest) (_result *UpdatePromotionActivityStatusResponse, _err error) {
	return invoke[UpdatePromotionActivityStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/activity/update_promotion_activity_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id": tea.StringValue(request.ActivityId),
			"biz_type":    tea.IntValue(request.BizType),
			"status":      tea.IntValue(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateSimpleQrBind(request *UpdateSimpleQrBindRequest) (_result *UpdateSimpleQrBindResponse, _err error) {
	return invoke[UpdateSimpleQrBindResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/capacity/update_simple_qr_bind/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"before_qr_url":           tea.StringValue(request.BeforeQrUrl),
			"exclusive_qr_url_prefix": tea.Int32Value(request.ExclusiveQrUrlPrefix),
			"load_path":               tea.StringValue(request.LoadPath),
			"qr_url":                  tea.StringValue(request.QrUrl),
			"stage":                   tea.StringValue(request.Stage),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateSimpleQrBindStatus(request *UpdateSimpleQrBindStatusRequest) (_result *UpdateSimpleQrBindStatusResponse, _err error) {
	return invoke[UpdateSimpleQrBindStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/capacity/update_simple_qr_bind_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"qr_url": tea.StringValue(request.QrUrl),
			"status": tea.Int32Value(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateTalentCouponStatus(request *UpdateTalentCouponStatusRequest) (_result *UpdateTalentCouponStatusResponse, _err error) {
	return invoke[UpdateTalentCouponStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/update_talent_coupon_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_type":   tea.Int32Value(request.AccountType),
			"app_id":         tea.StringValue(request.AppId),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"open_id":        tea.StringValue(request.OpenId),
			"status":         tea.IntValue(request.Status),
			"talent_account": tea.StringValue(request.TalentAccount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UpdateTalentCouponStock(request *UpdateTalentCouponStockRequest) (_result *UpdateTalentCouponStockResponse, _err error) {
	return invoke[UpdateTalentCouponStockResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/update_talent_coupon_stock/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_type":   tea.Int32Value(request.AccountType),
			"action":         tea.StringValue(request.Action),
			"app_id":         tea.StringValue(request.AppId),
			"award_number":   tea.Int64Value(request.AwardNumber),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"number":         tea.Int64Value(request.Number),
			"open_id":        tea.StringValue(request.OpenId),
			"talent_account": tea.StringValue(request.TalentAccount),
			"unique_key":     tea.StringValue(request.UniqueKey),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UploadImage(request *UploadImageRequest) (_result *UploadImageResponse, _err error) {
	return invoke[UploadImageResponse](client, request, &transport.Request{
		Method:  "POST",
		Path:    "/api/douyin/v1/video/upload_image/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"image": request.Image,
		},
		Encoding:    transport.EncodingFileForm,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UploadUserGroupInfo(request *UploadUserGroupInfoRequest) (_result *UploadUserGroupInfoResponse, _err error) {
	return invoke[UploadUserGroupInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/upload_user_group_info",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":   tea.StringValue(request.AppId),
			"group_id": tea.StringValue(request.GroupId),
			"open_id":  tea.StringValue(request.OpenId),
			"room_id":  tea.StringValue(request.RoomId),
			"round_id": tea.Int64Value(request.RoundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlGenerateSchema(request *UrlGenerateSchemaRequest) (_result *UrlGenerateSchemaResponse, _err error) {
	return invoke[UrlGenerateSchemaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url/generate_schema/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":      tea.StringValue(request.AppId),
			"expire_time": tea.Int64Value(request.ExpireTime),
			"no_expire":   tea.BoolValue(request.NoExpire),
			"path":        tea.StringValue(request.Path),
			"query":       tea.StringValue(request.Query),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlLinkGenerate(request *UrlLinkGenerateRequest) (_result *UrlLinkGenerateResponse, _err error) {
	return invoke[UrlLinkGenerateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url_link/generate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":      tea.StringValue(request.AppId),
			"app_name":    tea.StringValue(request.AppName),
			"expire_time": tea.Int64Value(request.ExpireTime),
			"path":        tea.StringValue(request.Path),
			"query":       tea.StringValue(request.Query),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlLinkQueryInfo(request *UrlLinkQueryInfoRequest) (_result *UrlLinkQueryInfoResponse, _err error) {
	return invoke[UrlLinkQueryInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url_link/query_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":   tea.StringValue(request.AppId),
			"url_link": tea.StringValue(request.UrlLink),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlLinkQueryQuota(request *UrlLinkQueryQuotaRequest) (_result *UrlLinkQueryQuotaResponse, _err error) {
	return invoke[UrlLinkQueryQuotaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url_link/query_quota/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id": tea.StringValue(request.AppId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlQuerySchema(request *UrlQuerySchemaRequest) (_result *UrlQuerySchemaResponse, _err error) {
	return invoke[UrlQuerySchemaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url/query_schema/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id": tea.StringValue(request.AppId),
			"schema": tea.StringValue(request.Schema),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UrlQuerySchemaQuota(request *UrlQuerySchemaQuotaRequest) (_result *UrlQuerySchemaQuotaResponse, _err error) {
	return invoke[UrlQuerySchemaQuotaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/url/query_schema_quota/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id": tea.StringValue(request.AppId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetComment(request *UserBcGetCommentRequest) (_result *UserBcGetCommentResponse, _err error) {
	return invoke[UserBcGetCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetFans(request *UserBcGetFansRequest) (_result *UserBcGetFansResponse, _err error) {
	return invoke[UserBcGetFansResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_fans/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetItem(request *UserBcGetItemRequest) (_result *UserBcGetItemResponse, _err error) {
	return invoke[UserBcGetItemResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_item/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetLike(request *UserBcGetLikeRequest) (_result *UserBcGetLikeResponse, _err error) {
	return invoke[UserBcGetLikeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_like/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetProfile(request *UserBcGetProfileRequest) (_result *UserBcGetProfileResponse, _err error) {
	return invoke[UserBcGetProfileResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_profile/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserBcGetShare(request *UserBcGetShareRequest) (_result *UserBcGetShareResponse, _err error) {
	return invoke[UserBcGetShareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user_bc/get_share/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserChange(request *UserChangeRequest) (_result *UserChangeResponse, _err error) {
	return invoke[UserChangeResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/member/hotel/user/change/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"member_card_id": tea.StringValue(request.MemberCardId),
			"open_id":        tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserComment(request *UserCommentRequest) (_result *UserCommentResponse, _err error) {
	return invoke[UserCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserFans(request *UserFansRequest) (_result *UserFansResponse, _err error) {
	return invoke[UserFansResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/fans/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserFansData(request *UserFansDataRequest) (_result *UserFansDataResponse, _err error) {
	return invoke[UserFansDataResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/douyin/v1/user/fans_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetComment(request *UserGetCommentRequest) (_result *UserGetCommentResponse, _err error) {
	return invoke[UserGetCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetFans(request *UserGetFansRequest) (_result *UserGetFansResponse, _err error) {
	return invoke[UserGetFansResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_fans/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetItem(request *UserGetItemRequest) (_result *UserGetItemResponse, _err error) {
	return invoke[UserGetItemResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_item/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetLike(request *UserGetLikeRequest) (_result *UserGetLikeResponse, _err error) {
	return invoke[UserGetLikeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_like/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetProfile(request *UserGetProfileRequest) (_result *UserGetProfileResponse, _err error) {
	return invoke[UserGetProfileResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_profile/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserGetShare(request *UserGetShareRequest) (_result *UserGetShareResponse, _err error) {
	return invoke[UserGetShareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/user/get_share/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserItem(request *UserItemRequest) (_result *UserItemResponse, _err error) {
	return invoke[UserItemResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/item/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserLike(request *UserLikeRequest) (_result *UserLikeResponse, _err error) {
	return invoke[UserLikeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/like/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserLogout(request *UserLogoutRequest) (_result *UserLogoutResponse, _err error) {
	return invoke[UserLogoutResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/member/hotel/user/logout/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":  tea.StringValue(request.AccountId),
			"open_id":     tea.StringValue(request.OpenId),
			"update_time": tea.Int64Value(request.UpdateTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserProfile(request *UserProfileRequest) (_result *UserProfileResponse, _err error) {
	return invoke[UserProfileResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/profile/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserShare(request *UserShareRequest) (_result *UserShareResponse, _err error) {
	return invoke[UserShareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/external/user/share/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) UserUpdate(request *UserUpdateRequest) (_result *UserUpdateResponse, _err error) {
	return invoke[UserUpdateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/member/user/update/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"members":    request.Members,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1AppRegionAdd(request *V1AppRegionAddRequest) (_result *V1AppRegionAddResponse, _err error) {
	return invoke[V1AppRegionAddResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/product/v1/app_region_add/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"bind_type": tea.IntValue(request.BindType),
			"path":      tea.StringValue(request.Path),
			"region_id": tea.StringValue(request.RegionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1AppRegionDelete(request *V1AppRegionDeleteRequest) (_result *V1AppRegionDeleteResponse, _err error) {
	return invoke[V1AppRegionDeleteResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/product/v1/app_region_delete/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"path":      tea.StringValue(request.Path),
			"region_id": tea.StringValue(request.RegionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1AppRegionModify(request *V1AppRegionModifyRequest) (_result *V1AppRegionModifyResponse, _err error) {
	return invoke[V1AppRegionModifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/product/v1/app_region_modify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"current_region": tea.StringValue(request.CurrentRegion),
			"origin_region":  tea.StringValue(request.OriginRegion),
			"path":           tea.StringValue(request.Path),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1AuthGetRelatedId(request *V1AuthGetRelatedIdRequest) (_result *V1AuthGetRelatedIdResponse, _err error) {
	return invoke[V1AuthGetRelatedIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/auth/get_related_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1GetPhonenumberInfo(request *V1GetPhonenumberInfoRequest) (_result *V1GetPhonenumberInfoResponse, _err error) {
	return invoke[V1GetPhonenumberInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/get_phonenumber_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"code": tea.StringValue(request.Code),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1GoodsProductSave(request *V1GoodsProductSaveRequest) (_result *V1GoodsProductSaveResponse, _err error) {
	return invoke[V1GoodsProductSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/product/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"ability":    request.Ability,
			"account_id": tea.StringValue(request.AccountId),
			"product":    request.Product,
			"sku":        request.Sku,
			"skus":       request.Skus,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1GoodsStockSync(request *V1GoodsStockSyncRequest) (_result *V1GoodsStockSyncResponse, _err error) {
	return invoke[V1GoodsStockSyncResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/stock/sync/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.StringValue(request.ProductId),
			"stock":      request.Stock,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1ListKolOrder(request *V1ListKolOrderRequest) (_result *V1ListKolOrderResponse, _err error) {
	return invoke[V1ListKolOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/solution/v1/list_kol_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"create_time_end":   tea.Int64Value(request.CreateTimeEnd),
			"create_time_start": tea.Int64Value(request.CreateTimeStart),
			"cursor":            tea.Int64Value(request.Cursor),
			"kol_id":            tea.StringValue(request.KolId),
			"open_id":           tea.StringValue(request.OpenId),
			"page_size":         tea.Int64Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V1Notify(request *V1NotifyRequest) (_result *V1NotifyResponse, _err error) {
	return invoke[V1NotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/subscribe_notification/developer/v1/notify",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_token": tea.StringValue(request.AccessToken),
			"app_id":       tea.StringValue(request.AppId),
			"data":         request.Data,
			"open_id":      tea.StringValue(request.OpenId),
			"page":         tea.StringValue(request.Page),
			"tpl_id":       tea.StringValue(request.TplId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V2CouponCreateCouponMeta(request *V2CouponCreateCouponMetaRequest) (_result *V2CouponCreateCouponMetaResponse, _err error) {
	return invoke[V2CouponCreateCouponMetaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/coupon/create_coupon_meta/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"biz_type":    tea.IntValue(request.BizType),
			"coupon_meta": request.CouponMeta,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V2FileUploadMaterial(request *V2FileUploadMaterialRequest) (_result *V2FileUploadMaterialResponse, _err error) {
	return invoke[V2FileUploadMaterialResponse](client, request, &transport.Request{
		Method:  "POST",
		Path:    "/api/apps/v2/file/upload_material/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Body: map[string]interface{}{
			"material_file": request.MaterialFile,
			"material_type": tea.Int32Value(request.MaterialType),
		},
		Encoding:    transport.EncodingFileForm,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V2Jscode2session(request *V2Jscode2sessionRequest) (_result *V2Jscode2sessionResponse, _err error) {
	return invoke[V2Jscode2sessionResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/jscode2session",
		Host:        client.GetHost("developer.toutiao.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anonymous_code": tea.StringValue(request.AnonymousCode),
			"appid":          tea.StringValue(request.Appid),
			"code":           tea.StringValue(request.Code),
			"secret":         tea.StringValue(request.Secret),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) V2SearchVideo(request *V2SearchVideoRequest) (_result *V2SearchVideoResponse, _err error) {
	return invoke[V2SearchVideoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/dy_open_api/v2/search/video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"count":        tea.Int32Value(request.Count),
			"cursor":       tea.Int64Value(request.Cursor),
			"device_id":    tea.Int64Value(request.DeviceId),
			"keyword":      tea.StringValue(request.Keyword),
			"open_id":      tea.StringValue(request.OpenId),
			"publish_time": tea.Int32Value(request.PublishTime),
			"search_id":    tea.StringValue(request.SearchId),
			"sort_type":    tea.Int32Value(request.SortType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V2TaskboxQueryTaskInfo(request *V2TaskboxQueryTaskInfoRequest) (_result *V2TaskboxQueryTaskInfoResponse, _err error) {
	return invoke[V2TaskboxQueryTaskInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/query_task_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"page_num":             tea.Int32Value(request.PageNum),
			"page_size":            tea.Int32Value(request.PageSize),
			"query_params_content": tea.StringValue(request.QueryParamsContent),
			"query_params_type":    tea.IntValue(request.QueryParamsType),
			"task_category":        tea.IntValue(request.TaskCategory),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) V2Token(request *V2TokenRequest) (_result *V2TokenResponse, _err error) {
	return invoke[V2TokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/token",
		Host:        client.GetHost("developer.toutiao.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":      tea.StringValue(request.Appid),
			"grant_type": tea.StringValue(request.GrantType),
			"secret":     tea.StringValue(request.Secret),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) V3Bills(request *V3BillsRequest) (_result *V3BillsResponse, _err error) {
	return invoke[V3BillsResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v3/bills/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":      tea.StringValue(request.AppId),
			"bill_date":   tea.StringValue(request.BillDate),
			"bill_type":   tea.StringValue(request.BillType),
			"merchant_id": tea.StringValue(request.MerchantId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VerifyRecordQuery(request *VerifyRecordQueryRequest) (_result *VerifyRecordQueryResponse, _err error) {
	return invoke[VerifyRecordQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/fulfilment/certificate/verify_record/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"cursor":     tea.StringValue(request.Cursor),
			"end_time":   tea.Int64Value(request.EndTime),
			"poi_ids":    tea.StringSliceValue(request.PoiIds),
			"size":       tea.Int32Value(request.Size),
			"start_time": tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoBcQuery(request *VideoBcQueryRequest) (_result *VideoBcQueryResponse, _err error) {
	return invoke[VideoBcQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/video_bc/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"item_ids": tea.StringSliceValue(request.ItemIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoCreateImageText(request *VideoCreateImageTextRequest) (_result *VideoCreateImageTextResponse, _err error) {
	return invoke[VideoCreateImageTextResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/create_image_text/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"agent_client_key": tea.StringValue(request.AgentClientKey),
			"at_users":         tea.StringSliceValue(request.AtUsers),
			"download_type":    tea.Int32Value(request.DownloadType),
			"image_list":       tea.StringSliceValue(request.ImageList),
			"micro_app_id":     tea.StringValue(request.MicroAppId),
			"micro_app_title":  tea.StringValue(request.MicroAppTitle),
			"micro_app_url":    tea.StringValue(request.MicroAppUrl),
			"music_id":         tea.Int64Value(request.MusicId),
			"poi_commerce":     tea.BoolValue(request.PoiCommerce),
			"poi_id":           tea.StringValue(request.PoiId),
			"private_status":   tea.Int32Value(request.PrivateStatus),
			"task_id":          tea.Int64Value(request.TaskId),
			"text":             tea.StringValue(request.Text),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoCreateVideo(request *VideoCreateVideoRequest) (_result *VideoCreateVideoResponse, _err error) {
	return invoke[VideoCreateVideoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/create_video/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"agent_client_key":          tea.StringValue(request.AgentClientKey),
			"at_users":                  tea.StringSliceValue(request.AtUsers),
			"cover_tsp":                 tea.Float64Value(request.CoverTsp),
			"custom_cover_image_url":    tea.StringValue(request.CustomCoverImageUrl),
			"download_type":             tea.Int32Value(request.DownloadType),
			"micro_app_id":              tea.StringValue(request.MicroAppId),
			"micro_app_title":           tea.StringValue(request.MicroAppTitle),
			"micro_app_url":             tea.StringValue(request.MicroAppUrl),
			"poi_commerce":              tea.BoolValue(request.PoiCommerce),
			"poi_id":                    tea.StringValue(request.PoiId),
			"private_status":            tea.Int32Value(request.PrivateStatus),
			"task_id":                   tea.Int64Value(request.TaskId),
			"text":                      tea.StringValue(request.Text),
			"video_id":                  tea.StringValue(request.VideoId),
			"vr_transcode_extra_params": request.VrTranscodeExtraParams,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoIdToOpenItemId(request *VideoIdToOpenItemIdRequest) (_result *VideoIdToOpenItemIdResponse, _err error) {
	return invoke[VideoIdToOpenItemIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/convert_video_id/video_id_to_open_item_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_key": tea.StringValue(request.AccessKey),
			"app_id":     tea.StringValue(request.AppId),
			"video_ids":  tea.StringSliceValue(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoMountApplyPermission(request *VideoMountApplyPermissionRequest) (_result *VideoMountApplyPermissionResponse, _err error) {
	return invoke[VideoMountApplyPermissionResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/video_mount/apply_permission/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_text":     tea.StringValue(request.AnchorText),
			"component_appid": tea.StringValue(request.ComponentAppid),
			"image_path":      tea.StringValue(request.ImagePath),
			"intro":           tea.StringValue(request.Intro),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoMountModifyConfig(request *VideoMountModifyConfigRequest) (_result *VideoMountModifyConfigResponse, _err error) {
	return invoke[VideoMountModifyConfigResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/video_mount/modify_config/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"config_key":   tea.StringValue(request.ConfigKey),
			"config_value": tea.StringValue(request.ConfigValue),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoMountQueryConfig(request *VideoMountQueryConfigRequest) (_result *VideoMountQueryConfigResponse, _err error) {
	return invoke[VideoMountQueryConfigResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/video_mount/query_config/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoMountQueryConfigHistory(request *VideoMountQueryConfigHistoryRequest) (_result *VideoMountQueryConfigHistoryResponse, _err error) {
	return invoke[VideoMountQueryConfigHistoryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/video_mount/query_config_history/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"page_num":  tea.Int32Value(request.PageNum),
			"page_size": tea.Int32Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoQuery(request *VideoQueryRequest) (_result *VideoQueryResponse, _err error) {
	return invoke[VideoQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/video/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"item_ids": tea.StringSliceValue(request.ItemIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoUploadVideo(request *VideoUploadVideoRequest) (_result *VideoUploadVideoResponse, _err error) {
	return invoke[VideoUploadVideoResponse](client, request, &transport.Request{
		Method:  "POST",
		Path:    "/api/douyin/v1/video/upload_video/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"video": request.Video,
		},
		Encoding:    transport.EncodingFileForm,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoUploadVideoPart(request *VideoUploadVideoPartRequest) (_result *VideoUploadVideoPartResponse, _err error) {
	return invoke[VideoUploadVideoPartResponse](client, request, &transport.Request{
		Method:  "POST",
		Path:    "/api/douyin/v1/video/upload_video_part/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"open_id":     tea.StringValue(request.OpenId),
			"part_number": tea.Int64Value(request.PartNumber),
			"upload_id":   tea.StringValue(request.UploadId),
		},
		Body: map[string]interface{}{
			"video": request.Video,
		},
		Encoding:    transport.EncodingFileForm,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoVideoBasicInfo(request *VideoVideoBasicInfoRequest) (_result *VideoVideoBasicInfoResponse, _err error) {
	return invoke[VideoVideoBasicInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/video_basic_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"item_ids":  tea.StringSliceValue(request.ItemIds),
			"video_ids": tea.StringSliceValue(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoVideoData(request *VideoVideoDataRequest) (_result *VideoVideoDataResponse, _err error) {
	return invoke[VideoVideoDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/video_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"item_ids":  tea.StringSliceValue(request.ItemIds),
			"video_ids": tea.StringSliceValue(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) VideoVideoList(request *VideoVideoListRequest) (_result *VideoVideoListResponse, _err error) {
	return invoke[VideoVideoListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/douyin/v1/video/video_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"count":   tea.Int32Value(request.Count),
			"cursor":  tea.Int64Value(request.Cursor),
			"open_id": tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WebcastmateInfo(request *WebcastmateInfoRequest) (_result *WebcastmateInfoResponse, _err error) {
	return invoke[WebcastmateInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/webcastmate/info",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"token": tea.StringValue(request.Token),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WithdrawCateringQuery(request *WithdrawCateringQueryRequest) (_result *WithdrawCateringQueryResponse, _err error) {
	return invoke[WithdrawCateringQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/withdraw/catering_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"cursor":     tea.StringValue(request.Cursor),
			"end_date":   tea.StringValue(request.EndDate),
			"size":       tea.Int64Value(request.Size),
			"start_date": tea.StringValue(request.StartDate),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WithdrawCompositeQuery(request *WithdrawCompositeQueryRequest) (_result *WithdrawCompositeQueryResponse, _err error) {
	return invoke[WithdrawCompositeQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/withdraw/composite_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"cursor":     tea.StringValue(request.Cursor),
			"end_date":   tea.StringValue(request.EndDate),
			"size":       tea.Int64Value(request.Size),
			"start_date": tea.StringValue(request.StartDate),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WorldRankSetValidVersion(request *WorldRankSetValidVersionRequest) (_result *WorldRankSetValidVersionResponse, _err error) {
	return invoke[WorldRankSetValidVersionResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/world_rank/set_valid_version",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"is_online_version":  tea.BoolValue(request.IsOnlineVersion),
			"world_rank_version": tea.StringValue(request.WorldRankVersion),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WorldRankUploadRankList(request *WorldRankUploadRankListRequest) (_result *WorldRankUploadRankListResponse, _err error) {
	return invoke[WorldRankUploadRankListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/world_rank/upload_rank_list",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"is_online_version":  tea.BoolValue(request.IsOnlineVersion),
			"rank_list":          request.RankList,
			"world_rank_version": tea.StringValue(request.WorldRankVersion),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) WorldRankUploadUserResult(request *WorldRankUploadUserResultRequest) (_result *WorldRankUploadUserResultResponse, _err error) {
	return invoke[WorldRankUploadUserResultResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/world_rank/upload_user_result",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"is_online_version":  tea.BoolValue(request.IsOnlineVersion),
			"user_list":          request.UserList,
			"world_rank_version": tea.StringValue(request.WorldRankVersion),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}
