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

func (client *Client) ActivityCreate(request *ActivityCreateRequest) (_result *ActivityCreateResponse, _err error) {
	return invoke[ActivityCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/dy_open_api/apps/v3/activity/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_name":                  tea.StringValue(request.ActivityName),
			"create_business_task_info_list": request.CreateBusinessTaskInfoList,
			"end_time":                       tea.Int64Value(request.EndTime),
			"start_time":                     tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ActivityCreatePromotionActivity(request *ActivityCreatePromotionActivityRequest) (_result *ActivityCreatePromotionActivityResponse, _err error) {
	return invoke[ActivityCreatePromotionActivityResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/activity/create_promotion_activity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"promotion_activity": request.PromotionActivity,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ActivityModify(request *ActivityModifyRequest) (_result *ActivityModifyResponse, _err error) {
	return invoke[ActivityModifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/dy_open_api/apps/v3/activity/modify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id":                    tea.Int64Value(request.ActivityId),
			"activity_name":                  tea.StringValue(request.ActivityName),
			"add_business_task_info_list":    request.AddBusinessTaskInfoList,
			"del_business_task_id_list":      tea.Int64ValueSlice(request.DelBusinessTaskIdList),
			"end_time":                       tea.Int64Value(request.EndTime),
			"modify_business_task_info_list": request.ModifyBusinessTaskInfoList,
			"start_time":                     tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ActivityModifyPromotionActivity(request *ActivityModifyPromotionActivityRequest) (_result *ActivityModifyPromotionActivityResponse, _err error) {
	return invoke[ActivityModifyPromotionActivityResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/activity/modify_promotion_activity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"biz_type":           tea.IntValue(request.BizType),
			"is_renewal":         tea.BoolValue(request.IsRenewal),
			"promotion_activity": request.PromotionActivity,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ActivityQueryBindedUser(request *ActivityQueryBindedUserRequest) (_result *ActivityQueryBindedUserResponse, _err error) {
	return invoke[ActivityQueryBindedUserResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/activity/query_binded_user/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id": tea.StringValue(request.ActivityId),
			"open_id":     tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ActivityQueryInfo(request *ActivityQueryInfoRequest) (_result *ActivityQueryInfoResponse, _err error) {
	return invoke[ActivityQueryInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/dy_open_api/apps/v3/activity/query_info/",
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

func (client *Client) ActivityQueryPromotionActivity(request *ActivityQueryPromotionActivityRequest) (_result *ActivityQueryPromotionActivityResponse, _err error) {
	return invoke[ActivityQueryPromotionActivityResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/activity/query_promotion_activity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id": tea.StringValue(request.ActivityId),
			"biz_type":    tea.IntValue(request.BizType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AdPlacementAdd(request *AdPlacementAddRequest) (_result *AdPlacementAddResponse, _err error) {
	return invoke[AdPlacementAddResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/ad_placement/add/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"ad_placement_name": tea.StringValue(request.AdPlacementName),
			"ad_placement_type": tea.Int64Value(request.AdPlacementType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AdPlacementQuery(request *AdPlacementQueryRequest) (_result *AdPlacementQueryResponse, _err error) {
	return invoke[AdPlacementQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/ad_placement/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AdPlacementUpdate(request *AdPlacementUpdateRequest) (_result *AdPlacementUpdateResponse, _err error) {
	return invoke[AdPlacementUpdateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/ad_placement/update/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"ad_placement_id": tea.StringValue(request.AdPlacementId),
			"status":          tea.Int64Value(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AddAppTestRelation(request *AddAppTestRelationRequest) (_result *AddAppTestRelationResponse, _err error) {
	return invoke[AddAppTestRelationResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/industry/v1/solution/add_app_test_relation/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"operator":    tea.StringValue(request.Operator),
			"ref_id_list": tea.StringSliceValue(request.RefIdList),
			"type":        tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AddAwemeVideoKeyword(request *AddAwemeVideoKeywordRequest) (_result *AddAwemeVideoKeywordResponse, _err error) {
	return invoke[AddAwemeVideoKeywordResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/add_aweme_video_keyword/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"keyword": tea.StringValue(request.Keyword),
			"reason":  tea.StringValue(request.Reason),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AddSimpleQrBind(request *AddSimpleQrBindRequest) (_result *AddSimpleQrBindResponse, _err error) {
	return invoke[AddSimpleQrBindResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/capacity/add_simple_qr_bind/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
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

func (client *Client) AffiliatedDetail(request *AffiliatedDetailRequest) (_result *AffiliatedDetailResponse, _err error) {
	return invoke[AffiliatedDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/affiliated/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":        tea.Int64Value(request.AccountId),
			"affiliated_id":     tea.Int64ValueSlice(request.AffiliatedId),
			"out_affiliated_id": tea.StringSliceValue(request.OutAffiliatedId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AffiliatedOperate(request *AffiliatedOperateRequest) (_result *AffiliatedOperateResponse, _err error) {
	return invoke[AffiliatedOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/affiliated/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":              request.Base,
			"account_id":        tea.Int64Value(request.AccountId),
			"op":                tea.IntValue(request.Op),
			"out_affiliated_id": tea.StringValue(request.OutAffiliatedId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AffiliatedPoiSellOut(request *AffiliatedPoiSellOutRequest) (_result *AffiliatedPoiSellOutResponse, _err error) {
	return invoke[AffiliatedPoiSellOutResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/affiliated/poi_sell_out/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":              request.Base,
			"account_id":        tea.Int64Value(request.AccountId),
			"ignore_fail_poi":   tea.BoolValue(request.IgnoreFailPoi),
			"out_affiliated_id": tea.StringValue(request.OutAffiliatedId),
			"poi_sell_out_rule": request.PoiSellOutRule,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AffiliatedSave(request *AffiliatedSaveRequest) (_result *AffiliatedSaveResponse, _err error) {
	return invoke[AffiliatedSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/affiliated/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":      tea.Int64Value(request.AccountId),
			"affiliated":      request.Affiliated,
			"ignore_fail_poi": tea.BoolValue(request.IgnoreFailPoi),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AfterSaleOrderAudit(request *AfterSaleOrderAuditRequest) (_result *AfterSaleOrderAuditResponse, _err error) {
	return invoke[AfterSaleOrderAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/after_sale/order/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"audit_status":   tea.BoolValue(request.AuditStatus),
			"certificate_id": tea.StringValue(request.CertificateId),
			"order_id":       tea.StringValue(request.OrderId),
			"reject_reason":  tea.StringValue(request.RejectReason),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AfterSaleOrderDetailGet(request *AfterSaleOrderDetailGetRequest) (_result *AfterSaleOrderDetailGetResponse, _err error) {
	return invoke[AfterSaleOrderDetailGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/after_sale/order_detail/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"certificate_id": tea.StringValue(request.CertificateId),
			"order_id":       tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AgencyQueryBillLink(request *AgencyQueryBillLinkRequest) (_result *AgencyQueryBillLinkResponse, _err error) {
	return invoke[AgencyQueryBillLinkResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/agency_query_bill_link/",
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

func (client *Client) AgencyQueryVideoSumData(request *AgencyQueryVideoSumDataRequest) (_result *AgencyQueryVideoSumDataResponse, _err error) {
	return invoke[AgencyQueryVideoSumDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/taskbox/agency_query_video_sum_data/",
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

func (client *Client) AkteAfterSaleOrderQuery(request *AkteAfterSaleOrderQueryRequest) (_result *AkteAfterSaleOrderQueryResponse, _err error) {
	return invoke[AkteAfterSaleOrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/after_sale/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"create_order_end_time":   tea.Int64Value(request.CreateOrderEndTime),
			"create_order_start_time": tea.Int64Value(request.CreateOrderStartTime),
			"cursor":                  tea.StringValue(request.Cursor),
			"page_size":               tea.Int32Value(request.PageSize),
			"refund_done_end_time":    tea.Int64Value(request.RefundDoneEndTime),
			"refund_done_start_time":  tea.Int64Value(request.RefundDoneStartTime),
			"refund_status":           tea.IntValue(request.RefundStatus),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AkteCommentReply(request *AkteCommentReplyRequest) (_result *AkteCommentReplyResponse, _err error) {
	return invoke[AkteCommentReplyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/comment/reply/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"poi_id":     tea.Int64Value(request.PoiId),
			"rate_id":    tea.Int64Value(request.RateId),
			"text":       tea.StringValue(request.Text),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AkteOrderQuery(request *AkteOrderQueryRequest) (_result *AkteOrderQueryResponse, _err error) {
	return invoke[AkteOrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"create_order_end_time":   tea.Int64Value(request.CreateOrderEndTime),
			"create_order_start_time": tea.Int64Value(request.CreateOrderStartTime),
			"cursor":                  tea.StringSliceValue(request.Cursor),
			"ext_order_id":            tea.StringValue(request.ExtOrderId),
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

func (client *Client) AliasCreateAlias(request *AliasCreateAliasRequest) (_result *AliasCreateAliasResponse, _err error) {
	return invoke[AliasCreateAliasResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/apps/v1/alias/create_alias/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"alias_word": tea.StringValue(request.AliasWord),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AliasListAlias(request *AliasListAliasRequest) (_result *AliasListAliasResponse, _err error) {
	return invoke[AliasListAliasResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/aweme/apps/v1/alias/list_alias/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AmusementNew(request *AmusementNewRequest) (_result *AmusementNewResponse, _err error) {
	return invoke[AmusementNewResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/amusement/new/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AmusementOverall(request *AmusementOverallRequest) (_result *AmusementOverallResponse, _err error) {
	return invoke[AmusementOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/amusement/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AnchorLinkmicCreate(request *AnchorLinkmicCreateRequest) (_result *AnchorLinkmicCreateResponse, _err error) {
	return invoke[AnchorLinkmicCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/anchor_linkmic/create",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":                tea.StringValue(request.AppId),
			"max_link_num_per_room": tea.Int64Value(request.MaxLinkNumPerRoom),
			"room_id":               tea.Int64Value(request.RoomId),
			"room_id_str":           tea.StringValue(request.RoomIdStr),
			"rooms":                 request.Rooms,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AnchorLinkmicPrepare(request *AnchorLinkmicPrepareRequest) (_result *AnchorLinkmicPrepareResponse, _err error) {
	return invoke[AnchorLinkmicPrepareResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/anchor_linkmic/prepare",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":      tea.StringValue(request.AppId),
			"room_id":     tea.Int64Value(request.RoomId),
			"room_id_str": tea.StringValue(request.RoomIdStr),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AnchorLinkmicQuery(request *AnchorLinkmicQueryRequest) (_result *AnchorLinkmicQueryResponse, _err error) {
	return invoke[AnchorLinkmicQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/anchor_linkmic/query",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":  tea.StringValue(request.AppId),
			"link_id": tea.Int64Value(request.LinkId),
			"room_id": tea.Int64Value(request.RoomId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AnchorLinkmicUpdateVoice(request *AnchorLinkmicUpdateVoiceRequest) (_result *AnchorLinkmicUpdateVoiceResponse, _err error) {
	return invoke[AnchorLinkmicUpdateVoiceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/anchor_linkmic/update_voice",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":     tea.StringValue(request.AppId),
			"link_id":    tea.Int64Value(request.LinkId),
			"user_infos": request.UserInfos,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppAddSubMerchant(request *AppAddSubMerchantRequest) (_result *AppAddSubMerchantResponse, _err error) {
	return invoke[AppAddSubMerchantResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/ecpay/v3/saas/app_add_sub_merchant/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":          tea.StringValue(request.AppId),
			"role":            tea.Int32Value(request.Role),
			"sub_merchant_id": tea.StringValue(request.SubMerchantId),
			"url_type":        tea.IntValue(request.UrlType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsCheckSessionKey(request *AppsCheckSessionKeyRequest) (_result *AppsCheckSessionKeyResponse, _err error) {
	return invoke[AppsCheckSessionKeyResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/mgplatform/api/apps/check_session_key",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"openid":     tea.StringValue(request.Openid),
			"sig_method": tea.StringValue(request.SigMethod),
			"signature":  tea.StringValue(request.Signature),
		},
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsJscode2session(request *AppsJscode2sessionRequest) (_result *AppsJscode2sessionResponse, _err error) {
	return invoke[AppsJscode2sessionResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/mgplatform/api/apps/jscode2session",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"anonymous_code": tea.StringValue(request.AnonymousCode),
			"appid":          tea.StringValue(request.Appid),
			"code":           tea.StringValue(request.Code),
			"secret":         tea.StringValue(request.Secret),
		},
	})
}

func (client *Client) AppsRemoveUserStorage(request *AppsRemoveUserStorageRequest) (_result *AppsRemoveUserStorageResponse, _err error) {
	return invoke[AppsRemoveUserStorageResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/remove_user_storage",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"openid":     tea.StringValue(request.Openid),
			"sig_method": tea.StringValue(request.SigMethod),
			"signature":  tea.StringValue(request.Signature),
		},
		Body: map[string]interface{}{
			"key": tea.StringSliceValue(request.Key),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsResetSessionKey(request *AppsResetSessionKeyRequest) (_result *AppsResetSessionKeyResponse, _err error) {
	return invoke[AppsResetSessionKeyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/reset_session_key",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_token": tea.StringValue(request.AccessToken),
			"openid":       tea.StringValue(request.Openid),
			"sig_method":   tea.StringValue(request.SigMethod),
			"signature":    tea.StringValue(request.Signature),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsSetUserStorage(request *AppsSetUserStorageRequest) (_result *AppsSetUserStorageResponse, _err error) {
	return invoke[AppsSetUserStorageResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/set_user_storage",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"openid":     tea.StringValue(request.Openid),
			"sig_method": tea.StringValue(request.SigMethod),
			"signature":  tea.StringValue(request.Signature),
		},
		Body: map[string]interface{}{
			"kv_list": request.KvList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsStableToken(request *AppsStableTokenRequest) (_result *AppsStableTokenResponse, _err error) {
	return invoke[AppsStableTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/stable_token",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":         tea.StringValue(request.Appid),
			"force_refresh": tea.BoolValue(request.ForceRefresh),
			"grant_type":    tea.StringValue(request.GrantType),
			"secret":        tea.StringValue(request.Secret),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) AppsUrlLinkGenerate(request *AppsUrlLinkGenerateRequest) (_result *AppsUrlLinkGenerateResponse, _err error) {
	return invoke[AppsUrlLinkGenerateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/url_link/generate",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_name":     tea.StringValue(request.AppName),
			"expire_time":  tea.Int64Value(request.ExpireTime),
			"query":        tea.StringValue(request.Query),
			"version_type": tea.StringValue(request.VersionType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsUrlLinkQueryInfo(request *AppsUrlLinkQueryInfoRequest) (_result *AppsUrlLinkQueryInfoResponse, _err error) {
	return invoke[AppsUrlLinkQueryInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/url_link/query_info",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"url_link": tea.StringValue(request.UrlLink),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsUrlLinkQueryQuota(request *AppsUrlLinkQueryQuotaRequest) (_result *AppsUrlLinkQueryQuotaResponse, _err error) {
	return invoke[AppsUrlLinkQueryQuotaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/url_link/query_quota",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AppsV2Token(request *AppsV2TokenRequest) (_result *AppsV2TokenResponse, _err error) {
	return invoke[AppsV2TokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/v2/token",
		Host:        client.GetHost("minigame.zijieapi.com"),
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

func (client *Client) AriNotify(request *AriNotifyRequest) (_result *AriNotifyResponse, _err error) {
	return invoke[AriNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/ari/notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":    tea.StringValue(request.AccountId),
			"date_range":    request.DateRange,
			"notify_scene":  tea.IntValueSlice(request.NotifyScene),
			"rate_plan_ids": tea.StringSliceValue(request.RatePlanIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AuditGet(request *AuditGetRequest) (_result *AuditGetResponse, _err error) {
	return invoke[AuditGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/im/group/enter/audit/get/",
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

func (client *Client) AuditNotify(request *AuditNotifyRequest) (_result *AuditNotifyResponse, _err error) {
	return invoke[AuditNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/after_sale/audit/notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"after_sale_id":     tea.StringValue(request.AfterSaleId),
			"is_approved":       tea.BoolValue(request.IsApproved),
			"out_after_sale_id": tea.StringValue(request.OutAfterSaleId),
			"reject_reason":     request.RejectReason,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AuditSet(request *AuditSetRequest) (_result *AuditSetResponse, _err error) {
	return invoke[AuditSetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/group/enter/audit/set/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"apply_id": tea.StringValue(request.ApplyId),
			"status":   tea.Int64Value(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AuthGetRelatedId(request *AuthGetRelatedIdRequest) (_result *AuthGetRelatedIdResponse, _err error) {
	return invoke[AuthGetRelatedIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/auth/get_related_id/",
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

func (client *Client) AuthorizeStatus(request *AuthorizeStatusRequest) (_result *AuthorizeStatusResponse, _err error) {
	return invoke[AuthorizeStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/authorize/status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"app_id":    tea.StringValue(request.AppId),
			"c_open_id": tea.StringValue(request.COpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) AuthorizeUserList(request *AuthorizeUserListRequest) (_result *AuthorizeUserListResponse, _err error) {
	return invoke[AuthorizeUserListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/authorize/user_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"cursor":    tea.StringValue(request.Cursor),
			"limit":     tea.Int64Value(request.Limit),
			"page_num":  tea.Int64Value(request.PageNum),
			"page_size": tea.Int64Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BatchRollbackConsumeCoupon(request *BatchRollbackConsumeCouponRequest) (_result *BatchRollbackConsumeCouponResponse, _err error) {
	return invoke[BatchRollbackConsumeCouponResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/batch_rollback_consume_coupon/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":                  tea.StringValue(request.AppId),
			"consume_out_no":          tea.StringValue(request.ConsumeOutNo),
			"coupon_id_list":          tea.StringSliceValue(request.CouponIdList),
			"open_id":                 tea.StringValue(request.OpenId),
			"order_id":                tea.StringValue(request.OrderId),
			"rollback_consume_out_no": tea.StringValue(request.RollbackConsumeOutNo),
			"rollback_consume_time":   tea.Int64Value(request.RollbackConsumeTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillCateringQuery(request *BillCateringQueryRequest) (_result *BillCateringQueryResponse, _err error) {
	return invoke[BillCateringQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/bill/catering_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"bill_date":       tea.StringValue(request.BillDate),
			"biz_type":        tea.IntValue(request.BizType),
			"cursor":          tea.StringValue(request.Cursor),
			"root_account_id": tea.StringValue(request.RootAccountId),
			"size":            tea.Int64Value(request.Size),
			"withdraw_id":     tea.StringValue(request.WithdrawId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillQueryLegerUrl(request *BillQueryLegerUrlRequest) (_result *BillQueryLegerUrlResponse, _err error) {
	return invoke[BillQueryLegerUrlResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/bill/query_leger_url/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":  tea.StringValue(request.AccountId),
			"launch_date": tea.StringValue(request.LaunchDate),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillQueryRebateUrl(request *BillQueryRebateUrlRequest) (_result *BillQueryRebateUrlResponse, _err error) {
	return invoke[BillQueryRebateUrlResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/bill/query_rebate_url/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"biz_month":  tea.StringValue(request.BizMonth),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillboardProp(request *BillboardPropRequest) (_result *BillboardPropResponse, _err error) {
	return invoke[BillboardPropResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/prop/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillboardStars(request *BillboardStarsRequest) (_result *BillboardStarsResponse, _err error) {
	return invoke[BillboardStarsResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/stars/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BillboardTopic(request *BillboardTopicRequest) (_result *BillboardTopicResponse, _err error) {
	return invoke[BillboardTopicResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/topic/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BindDetail(request *BindDetailRequest) (_result *BindDetailResponse, _err error) {
	return invoke[BindDetailResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/foodorder/couponproduct/bind/detail/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":                  request.Base,
			"account_id":            tea.Int64Value(request.AccountId),
			"coupon_product_id":     tea.Int64Value(request.CouponProductId),
			"out_coupon_product_id": tea.StringValue(request.OutCouponProductId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BindInfoAll(request *BindInfoAllRequest) (_result *BindInfoAllResponse, _err error) {
	return invoke[BindInfoAllResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v2/craftsman_openapi/merchat/craftsman/bind_info/all/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
			"cursor":     tea.Int64Value(request.Cursor),
			"size":       tea.Int64Value(request.Size),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BindInfoSingle(request *BindInfoSingleRequest) (_result *BindInfoSingleResponse, _err error) {
	return invoke[BindInfoSingleResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v2/craftsman_openapi/merchat/craftsman/bind_info/single/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":    tea.Int64Value(request.AccountId),
			"craftsman_uid": tea.StringValue(request.CraftsmanUid),
			"limit":         tea.Int64Value(request.Limit),
			"offset":        tea.Int64Value(request.Offset),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BindSave(request *BindSaveRequest) (_result *BindSaveResponse, _err error) {
	return invoke[BindSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/couponproduct/bind/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":            tea.Int64Value(request.AccountId),
			"coupon_product_id":     tea.Int64Value(request.CouponProductId),
			"fulfillment_type":      tea.IntValueSlice(request.FulfillmentType),
			"group_list":            request.GroupList,
			"out_coupon_product_id": tea.StringValue(request.OutCouponProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BookUserCancelBook(request *BookUserCancelBookRequest) (_result *BookUserCancelBookResponse, _err error) {
	return invoke[BookUserCancelBookResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/book/user_cancel_book/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"book_id":       tea.StringValue(request.BookId),
			"cancel_reason": tea.StringSliceValue(request.CancelReason),
			"open_id":       tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BookingAuditNotify(request *BookingAuditNotifyRequest) (_result *BookingAuditNotifyResponse, _err error) {
	return invoke[BookingAuditNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/hotel/booking/audit/notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"accommodation_status": tea.IntValue(request.AccommodationStatus),
			"order_id":             tea.StringValue(request.OrderId),
			"order_out_id":         tea.StringValue(request.OrderOutId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BookingConfig(request *BookingConfigRequest) (_result *BookingConfigResponse, _err error) {
	return invoke[BookingConfigResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/booking/config/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"config":     request.Config,
			"poi_ids":    tea.StringSliceValue(request.PoiIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BookingOrderFulfillment(request *BookingOrderFulfillmentRequest) (_result *BookingOrderFulfillmentResponse, _err error) {
	return invoke[BookingOrderFulfillmentResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/booking/order/fulfillment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":         tea.StringValue(request.AccountId),
			"ext_order_id":       tea.StringValue(request.ExtOrderId),
			"fulfillment_status": tea.IntValue(request.FulfillmentStatus),
			"order_id":           tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) BuyMerchantConfirmOrder(request *BuyMerchantConfirmOrderRequest) (_result *BuyMerchantConfirmOrderResponse, _err error) {
	return invoke[BuyMerchantConfirmOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trade/buy/merchant_confirm_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id": tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CalendarAri(request *CalendarAriRequest) (_result *CalendarAriResponse, _err error) {
	return invoke[CalendarAriResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/calendar_product/save/calendar/ari/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
		},
		Body: map[string]interface{}{
			"booking_calendar_stock": request.BookingCalendarStock,
			"out_id":                 tea.StringValue(request.OutId),
			"product_id":             tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CalendarOrderConfirm(request *CalendarOrderConfirmRequest) (_result *CalendarOrderConfirmResponse, _err error) {
	return invoke[CalendarOrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/calendar/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
		},
		Body: map[string]interface{}{
			"confirm_info": request.ConfirmInfo,
			"order_id":     tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CalendarProductListCalendarAri(request *CalendarProductListCalendarAriRequest) (_result *CalendarProductListCalendarAriResponse, _err error) {
	return invoke[CalendarProductListCalendarAriResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/calendar_product/list_calendar_ari/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
		},
		Body: map[string]interface{}{
			"end_date":   tea.StringValue(request.EndDate),
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.StringValue(request.ProductId),
			"start_date": tea.StringValue(request.StartDate),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CancelAudit(request *CancelAuditRequest) (_result *CancelAuditResponse, _err error) {
	return invoke[CancelAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/hotel/cancel/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"cancel_Id":     tea.StringValue(request.CancelId),
			"cancel_result": tea.IntValue(request.CancelResult),
			"cancel_type":   tea.IntValue(request.CancelType),
			"order_id":      tea.StringValue(request.OrderId),
			"reason":        tea.StringValue(request.Reason),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityApplyCapacity(request *CapacityApplyCapacityRequest) (_result *CapacityApplyCapacityResponse, _err error) {
	return invoke[CapacityApplyCapacityResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/apply_capacity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"apply_info":   request.ApplyInfo,
			"apply_reason": tea.StringValue(request.ApplyReason),
			"capacity_key": tea.StringValue(request.CapacityKey),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityBindAwemeRelation(request *CapacityBindAwemeRelationRequest) (_result *CapacityBindAwemeRelationResponse, _err error) {
	return invoke[CapacityBindAwemeRelationResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/bind_aweme_relation/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"audit_template_info": request.AuditTemplateInfo,
			"aweme_id":            tea.StringValue(request.AwemeId),
			"capacity_list":       tea.StringSliceValue(request.CapacityList),
			"co_subject":          tea.BoolValue(request.CoSubject),
			"cooperation_info":    request.CooperationInfo,
			"employee_info":       request.EmployeeInfo,
			"type":                tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityDeleteAlias(request *CapacityDeleteAliasRequest) (_result *CapacityDeleteAliasResponse, _err error) {
	return invoke[CapacityDeleteAliasResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/delete_alias/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"alias": tea.StringValue(request.Alias),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityModifyAlias(request *CapacityModifyAliasRequest) (_result *CapacityModifyAliasResponse, _err error) {
	return invoke[CapacityModifyAliasResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/modify_alias/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"after_alias":  tea.StringValue(request.AfterAlias),
			"before_alias": tea.StringValue(request.BeforeAlias),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityQueryAdIncome(request *CapacityQueryAdIncomeRequest) (_result *CapacityQueryAdIncomeResponse, _err error) {
	return invoke[CapacityQueryAdIncomeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v3/capacity/query_ad_income/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"end_date":   tea.StringValue(request.EndDate),
			"host_name":  tea.StringValue(request.HostName),
			"start_date": tea.StringValue(request.StartDate),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityQueryApplyStatus(request *CapacityQueryApplyStatusRequest) (_result *CapacityQueryApplyStatusResponse, _err error) {
	return invoke[CapacityQueryApplyStatusResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/query_apply_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"capacity_key": tea.StringValue(request.CapacityKey),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityQueryCapacityList(request *CapacityQueryCapacityListRequest) (_result *CapacityQueryCapacityListResponse, _err error) {
	return invoke[CapacityQueryCapacityListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/query_capacity_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacitySetSearchTag(request *CapacitySetSearchTagRequest) (_result *CapacitySetSearchTagResponse, _err error) {
	return invoke[CapacitySetSearchTagResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/set_search_tag/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"add_tag_list":    tea.StringSliceValue(request.AddTagList),
			"delete_tag_list": tea.StringSliceValue(request.DeleteTagList),
			"modify_tag_list": request.ModifyTagList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CapacityUnbindAwemeRelation(request *CapacityUnbindAwemeRelationRequest) (_result *CapacityUnbindAwemeRelationResponse, _err error) {
	return invoke[CapacityUnbindAwemeRelationResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/unbind_aweme_relation/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_id": tea.StringValue(request.AwemeId),
			"type":     tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CarComment(request *CarCommentRequest) (_result *CarCommentResponse, _err error) {
	return invoke[CarCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/car/comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CarDriver(request *CarDriverRequest) (_result *CarDriverResponse, _err error) {
	return invoke[CarDriverResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/car/driver/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CarOverall(request *CarOverallRequest) (_result *CarOverallResponse, _err error) {
	return invoke[CarOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/car/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CarPlay(request *CarPlayRequest) (_result *CarPlayResponse, _err error) {
	return invoke[CarPlayResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/car/play/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CarUse(request *CarUseRequest) (_result *CarUseResponse, _err error) {
	return invoke[CarUseResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/car/use/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardDelete(request *CardDeleteRequest) (_result *CardDeleteResponse, _err error) {
	return invoke[CardDeleteResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/delete",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"url":     tea.StringValue(request.Url),
			"user_id": request.UserId,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardGet(request *CardGetRequest) (_result *CardGetResponse, _err error) {
	return invoke[CardGetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/get",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"url":     tea.StringValue(request.Url),
			"user_id": request.UserId,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardImageDelete(request *CardImageDeleteRequest) (_result *CardImageDeleteResponse, _err error) {
	return invoke[CardImageDeleteResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/file/card_image/delete",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"image_ids": tea.StringValue(request.ImageIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardImageGet(request *CardImageGetRequest) (_result *CardImageGetResponse, _err error) {
	return invoke[CardImageGetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/file/card_image/get",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"image_ids": tea.StringValue(request.ImageIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardImageUpload(request *CardImageUploadRequest) (_result *CardImageUploadResponse, _err error) {
	return invoke[CardImageUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/file/card_image/upload",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardSet(request *CardSetRequest) (_result *CardSetResponse, _err error) {
	return invoke[CardSetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/set",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"card_type": request.CardType,
			"url":       tea.StringValue(request.Url),
			"user_id":   request.UserId,
			"value":     tea.StringValue(request.Value),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CardUpdate(request *CardUpdateRequest) (_result *CardUpdateResponse, _err error) {
	return invoke[CardUpdateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/update",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"card_type": request.CardType,
			"room_id":   tea.StringValue(request.RoomId),
			"value":     tea.StringValue(request.Value),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CategoryGetAuditCategories(request *CategoryGetAuditCategoriesRequest) (_result *CategoryGetAuditCategoriesResponse, _err error) {
	return invoke[CategoryGetAuditCategoriesResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/category/get_audit_categories/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CensorImage(request *CensorImageRequest) (_result *CensorImageResponse, _err error) {
	return invoke[CensorImageResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/censor/image/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":     tea.StringValue(request.AppId),
			"image":      tea.StringValue(request.Image),
			"image_data": tea.StringValue(request.ImageData),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertInfo(request *CertInfoRequest) (_result *CertInfoResponse, _err error) {
	return invoke[CertInfoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/poi/cert/info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"merchant_life_account_id": tea.Int64Value(request.MerchantLifeAccountId),
			"poi_id":                   tea.Int64Value(request.PoiId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateBatchVerify(request *CertificateBatchVerifyRequest) (_result *CertificateBatchVerifyResponse, _err error) {
	return invoke[CertificateBatchVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/certificate/batch_verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificate_batch_verify_list": request.CertificateBatchVerifyList,
			"poi_id":                        tea.StringValue(request.PoiId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateCallback(request *CertificateCallbackRequest) (_result *CertificateCallbackResponse, _err error) {
	return invoke[CertificateCallbackResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/scenic/groupon/certificate/callback/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificate_info_list": request.CertificateInfoList,
			"order_id":              tea.StringValue(request.OrderId),
			"result":                tea.IntValue(request.Result),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateCancel(request *CertificateCancelRequest) (_result *CertificateCancelResponse, _err error) {
	return invoke[CertificateCancelResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/certificate/cancel/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"batch_cancel_info":       request.BatchCancelInfo,
			"batch_cancel_info_list":  request.BatchCancelInfoList,
			"cancel_token":            tea.StringValue(request.CancelToken),
			"certificate_id":          tea.StringValue(request.CertificateId),
			"shop_order_id":           tea.StringValue(request.ShopOrderId),
			"times_card_cancel_count": tea.Int64Value(request.TimesCardCancelCount),
			"verify_id":               tea.StringValue(request.VerifyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateGet(request *CertificateGetRequest) (_result *CertificateGetResponse, _err error) {
	return invoke[CertificateGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/fulfilment/certificate/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"encrypted_code": tea.StringValue(request.EncryptedCode),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificatePrepare(request *CertificatePrepareRequest) (_result *CertificatePrepareResponse, _err error) {
	return invoke[CertificatePrepareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/fulfilment/certificate/prepare/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"code":           tea.StringValue(request.Code),
			"encrypted_data": tea.StringValue(request.EncryptedData),
			"poi_id":         tea.StringValue(request.PoiId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateQuery(request *CertificateQueryRequest) (_result *CertificateQueryResponse, _err error) {
	return invoke[CertificateQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/fulfilment/certificate/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"encrypted_code": tea.StringValue(request.EncryptedCode),
			"order_id":       tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CertificateVerify(request *CertificateVerifyRequest) (_result *CertificateVerifyResponse, _err error) {
	return invoke[CertificateVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/certificate/verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":          tea.StringValue(request.AccountId),
			"code_with_time_list": request.CodeWithTimeList,
			"codes":               tea.StringSliceValue(request.Codes),
			"encrypted_codes":     tea.StringSliceValue(request.EncryptedCodes),
			"order_id":            tea.StringValue(request.OrderId),
			"poi_id":              tea.StringValue(request.PoiId),
			"verify_extra":        request.VerifyExtra,
			"verify_sign_list":    tea.StringSliceValue(request.VerifySignList),
			"verify_token":        tea.StringValue(request.VerifyToken),
			"voucher":             request.Voucher,
			"vouchers":            request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ChangeUserBindAgent(request *ChangeUserBindAgentRequest) (_result *ChangeUserBindAgentResponse, _err error) {
	return invoke[ChangeUserBindAgentResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/change_user_bind_agent/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"douyin_id":    tea.StringValue(request.DouyinId),
			"new_agent_id": tea.Int64Value(request.NewAgentId),
			"old_agent_id": tea.Int64Value(request.OldAgentId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ClaimQuery(request *ClaimQueryRequest) (_result *ClaimQueryResponse, _err error) {
	return invoke[ClaimQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/merchant/claim/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"poi_ids":    tea.StringSliceValue(request.PoiIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ClueQuery(request *ClueQueryRequest) (_result *ClueQueryResponse, _err error) {
	return invoke[ClueQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/open_api/crm/clue/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"end_time":        tea.StringValue(request.EndTime),
			"life_account_id": tea.StringSliceValue(request.LifeAccountId),
			"open_id":         tea.StringValue(request.OpenId),
			"page":            tea.Int32Value(request.Page),
			"page_size":       tea.Int32Value(request.PageSize),
			"start_time":      tea.StringValue(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CoGameUploadUserData(request *CoGameUploadUserDataRequest) (_result *CoGameUploadUserDataResponse, _err error) {
	return invoke[CoGameUploadUserDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/round/co_game_upload_user_data",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchor_infos": request.AnchorInfos,
			"app_id":       tea.StringValue(request.AppId),
			"round_id":     tea.Int64Value(request.RoundId),
			"round_status": tea.Int64Value(request.RoundStatus),
			"user_list":    request.UserList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommentQuery(request *CommentQueryRequest) (_result *CommentQueryResponse, _err error) {
	return invoke[CommentQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/akte/comment/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"count":           tea.Int64Value(request.Count),
			"cursor":          tea.StringValue(request.Cursor),
			"end_time":        tea.Int64Value(request.EndTime),
			"poi_id_list":     tea.Int64ValueSlice(request.PoiIdList),
			"product_id_list": tea.Int64ValueSlice(request.ProductIdList),
			"start_time":      tea.Int64Value(request.StartTime),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRateOperate(request *CommissionRateOperateRequest) (_result *CommissionRateOperateResponse, _err error) {
	return invoke[CommissionRateOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/commission_rate/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"content":    request.Content,
			"op_type":    tea.IntValue(request.OpType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRatePlanIdQuery(request *CommissionRatePlanIdQueryRequest) (_result *CommissionRatePlanIdQueryResponse, _err error) {
	return invoke[CommissionRatePlanIdQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/commission_rate/plan_id_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"plan_ids":   tea.StringSliceValue(request.PlanIds),
			"plan_type":  tea.IntValue(request.PlanType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRateQuery(request *CommissionRateQueryRequest) (_result *CommissionRateQueryResponse, _err error) {
	return invoke[CommissionRateQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/commission_rate/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"dimension_id":   tea.StringValue(request.DimensionId),
			"dimension_type": tea.IntValue(request.DimensionType),
			"page_index":     tea.Int32Value(request.PageIndex),
			"page_size":      tea.Int32Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRateSave(request *CommissionRateSaveRequest) (_result *CommissionRateSaveResponse, _err error) {
	return invoke[CommissionRateSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/commission_rate/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":  tea.StringValue(request.AccountId),
			"plan_detail": request.PlanDetail,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRecordGet(request *CommissionRecordGetRequest) (_result *CommissionRecordGetResponse, _err error) {
	return invoke[CommissionRecordGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/commission_record/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"order_id":  tea.StringValue(request.OrderId),
			"record_id": tea.StringValue(request.RecordId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CommissionRecordQuery(request *CommissionRecordQueryRequest) (_result *CommissionRecordQueryResponse, _err error) {
	return invoke[CommissionRecordQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/commission_record/query/",
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

func (client *Client) CommonPlanSellDetail(request *CommonPlanSellDetailRequest) (_result *CommonPlanSellDetailResponse, _err error) {
	return invoke[CommonPlanSellDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/common_plan_sell_detail/",
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

func (client *Client) CommonPlanTalentDetail(request *CommonPlanTalentDetailRequest) (_result *CommonPlanTalentDetailResponse, _err error) {
	return invoke[CommonPlanTalentDetailResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/common_plan_talent_detail/",
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

func (client *Client) CommonPlanTalentMediaList(request *CommonPlanTalentMediaListRequest) (_result *CommonPlanTalentMediaListResponse, _err error) {
	return invoke[CommonPlanTalentMediaListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/common_plan_talent_media_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"content_type": tea.Int32Value(request.ContentType),
			"douyin_id":    tea.StringValue(request.DouyinId),
			"page_num":     tea.Int32Value(request.PageNum),
			"page_size":    tea.Int32Value(request.PageSize),
			"plan_id":      tea.Int64Value(request.PlanId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CompleteConfirm(request *CompleteConfirmRequest) (_result *CompleteConfirmResponse, _err error) {
	return invoke[CompleteConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/hermes/preparation/complete/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id": tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CompleteUploadUserResult(request *CompleteUploadUserResultRequest) (_result *CompleteUploadUserResultResponse, _err error) {
	return invoke[CompleteUploadUserResultResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/gaming_con/world_rank/complete_upload_user_result",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"complete_time":      tea.Int64Value(request.CompleteTime),
			"is_online_version":  tea.BoolValue(request.IsOnlineVersion),
			"world_rank_version": tea.StringValue(request.WorldRankVersion),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CompleteVideoPartUpload(request *CompleteVideoPartUploadRequest) (_result *CompleteVideoPartUploadResponse, _err error) {
	return invoke[CompleteVideoPartUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/complete_video_part_upload/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id":   tea.StringValue(request.OpenId),
			"upload_id": tea.StringValue(request.UploadId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ConfigLimitOpPoint(request *ConfigLimitOpPointRequest) (_result *ConfigLimitOpPointResponse, _err error) {
	return invoke[ConfigLimitOpPointResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/ecom/v1/config/limit_op_point/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":              tea.StringValue(request.AppId),
			"daily_add_count":     tea.Int64Value(request.DailyAddCount),
			"daily_deduct_count":  tea.Int64Value(request.DailyDeductCount),
			"op_type":             tea.IntValue(request.OpType),
			"single_add_limit":    tea.Int64Value(request.SingleAddLimit),
			"single_deduct_limit": tea.Int64Value(request.SingleDeductLimit),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ConfigRegisterMaApp(request *ConfigRegisterMaAppRequest) (_result *ConfigRegisterMaAppResponse, _err error) {
	return invoke[ConfigRegisterMaAppResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/ecom/v1/config/register_ma_app/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":             tea.StringValue(request.AppId),
			"callback_url":       tea.StringValue(request.CallbackUrl),
			"ecom_microapp_type": tea.IntValue(request.EcomMicroappType),
			"extra":              tea.StringValue(request.Extra),
			"preview_photo_url":  tea.StringValue(request.PreviewPhotoUrl),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ConvertVideoIdVideoIdToOpenItemId(request *ConvertVideoIdVideoIdToOpenItemIdRequest) (_result *ConvertVideoIdVideoIdToOpenItemIdResponse, _err error) {
	return invoke[ConvertVideoIdVideoIdToOpenItemIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/convert_video_id/video_id_to_open_item_id",
		Host:        client.GetHost("developer.toutiao.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_key":   tea.StringValue(request.AccessKey),
			"access_token": tea.StringValue(request.AccessToken),
			"app_id":       tea.StringValue(request.AppId),
			"video_ids":    tea.StringSliceValue(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaBrainCavity(request *CospaBrainCavityRequest) (_result *CospaBrainCavityResponse, _err error) {
	return invoke[CospaBrainCavityResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/brain_cavity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaNew(request *CospaNewRequest) (_result *CospaNewResponse, _err error) {
	return invoke[CospaNewResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/new/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaOutShot(request *CospaOutShotRequest) (_result *CospaOutShotResponse, _err error) {
	return invoke[CospaOutShotResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/out_shot/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaOverall(request *CospaOverallRequest) (_result *CospaOverallResponse, _err error) {
	return invoke[CospaOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaPainting(request *CospaPaintingRequest) (_result *CospaPaintingResponse, _err error) {
	return invoke[CospaPaintingResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/painting/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaQingMan(request *CospaQingManRequest) (_result *CospaQingManResponse, _err error) {
	return invoke[CospaQingManResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/qing_man/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CospaVoiceControl(request *CospaVoiceControlRequest) (_result *CospaVoiceControlResponse, _err error) {
	return invoke[CospaVoiceControlResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/cospa/voice_control/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponBatchConsumeCoupon(request *CouponBatchConsumeCouponRequest) (_result *CouponBatchConsumeCouponResponse, _err error) {
	return invoke[CouponBatchConsumeCouponResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/batch_consume_coupon/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":         tea.StringValue(request.AppId),
			"consume_out_no": tea.StringValue(request.ConsumeOutNo),
			"consume_time":   tea.Int64Value(request.ConsumeTime),
			"coupon_id_list": tea.StringSliceValue(request.CouponIdList),
			"open_id":        tea.StringValue(request.OpenId),
			"order_id":       tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponCreateDeveloperActivity(request *CouponCreateDeveloperActivityRequest) (_result *CouponCreateDeveloperActivityResponse, _err error) {
	return invoke[CouponCreateDeveloperActivityResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/create_developer_activity/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_name":        tea.StringValue(request.ActivityName),
			"app_id":               tea.StringValue(request.AppId),
			"coupon_meta_id":       tea.StringValue(request.CouponMetaId),
			"coupon_stock_number":  tea.Int64Value(request.CouponStockNumber),
			"merchant_activity_id": tea.StringValue(request.MerchantActivityId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponDeleteCouponMeta(request *CouponDeleteCouponMetaRequest) (_result *CouponDeleteCouponMetaResponse, _err error) {
	return invoke[CouponDeleteCouponMetaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/delete_coupon_meta/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":           tea.StringValue(request.AppId),
			"coupon_meta_id":   tea.StringValue(request.CouponMetaId),
			"merchant_meta_no": tea.StringValue(request.MerchantMetaNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponGetTalentCoupon(request *CouponGetTalentCouponRequest) (_result *CouponGetTalentCouponResponse, _err error) {
	return invoke[CouponGetTalentCouponResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/get_talent_coupon/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_type":   tea.Int32Value(request.AccountType),
			"app_id":         tea.StringValue(request.AppId),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"open_id":        tea.StringValue(request.OpenId),
			"talent_account": tea.StringValue(request.TalentAccount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponModifyCouponMeta(request *CouponModifyCouponMetaRequest) (_result *CouponModifyCouponMetaResponse, _err error) {
	return invoke[CouponModifyCouponMetaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/coupon/modify_coupon_meta/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"biz_type":    tea.IntValue(request.BizType),
			"coupon_meta": request.CouponMeta,
			"is_renewal":  tea.BoolValue(request.IsRenewal),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponPriceQuery(request *CouponPriceQueryRequest) (_result *CouponPriceQueryResponse, _err error) {
	return invoke[CouponPriceQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/product/coupon_price/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":              request.Base,
			"account_id":        tea.StringValue(request.AccountId),
			"order_create_time": tea.Int64Value(request.OrderCreateTime),
			"order_id":          tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponQueryCouponMeta(request *CouponQueryCouponMetaRequest) (_result *CouponQueryCouponMetaResponse, _err error) {
	return invoke[CouponQueryCouponMetaResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v2/coupon/query_coupon_meta/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"biz_type":       tea.IntValue(request.BizType),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CouponSetTalentCoupon(request *CouponSetTalentCouponRequest) (_result *CouponSetTalentCouponResponse, _err error) {
	return invoke[CouponSetTalentCouponResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/set_talent_coupon/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_type":      tea.Int32Value(request.AccountType),
			"app_id":            tea.StringValue(request.AppId),
			"award_stock_limit": tea.Int64Value(request.AwardStockLimit),
			"coupon_meta_id":    tea.StringValue(request.CouponMetaId),
			"open_id":           tea.StringValue(request.OpenId),
			"stock_limit":       tea.Int64Value(request.StockLimit),
			"talent_account":    tea.StringValue(request.TalentAccount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CreateCallback(request *CreateCallbackRequest) (_result *CreateCallbackResponse, _err error) {
	return invoke[CreateCallbackResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/fulfilment/create/callback/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"additional_info_list": request.AdditionalInfoList,
			"certificates":         request.Certificates,
			"codes":                tea.StringSliceValue(request.Codes),
			"fail_reason":          tea.StringValue(request.FailReason),
			"fail_reason_desc":     tea.StringValue(request.FailReasonDesc),
			"order_id":             tea.StringValue(request.OrderId),
			"result":               tea.Int64Value(request.Result),
			"third_order_id":       tea.StringValue(request.ThirdOrderId),
			"voucher":              request.Voucher,
			"vouchers":             request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CreateMaSubService(request *CreateMaSubServiceRequest) (_result *CreateMaSubServiceResponse, _err error) {
	return invoke[CreateMaSubServiceResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/create_ma_sub_service/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"search_key_word":  tea.StringSliceValue(request.SearchKeyWord),
			"start_page_url":   tea.StringValue(request.StartPageUrl),
			"sub_service_name": tea.StringValue(request.SubServiceName),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CrowdSave(request *CrowdSaveRequest) (_result *CrowdSaveResponse, _err error) {
	return invoke[CrowdSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/poi/crowd/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"crowds":     request.Crowds,
			"poi_id":     tea.StringValue(request.PoiId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) CustomizationQueryStatus(request *CustomizationQueryStatusRequest) (_result *CustomizationQueryStatusResponse, _err error) {
	return invoke[CustomizationQueryStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/ecom/v1/customization/query_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":   tea.StringValue(request.AppId),
			"open_id":  tea.StringValue(request.OpenId),
			"order_id": tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DataAnalysisQueryBehaviorData(request *DataAnalysisQueryBehaviorDataRequest) (_result *DataAnalysisQueryBehaviorDataResponse, _err error) {
	return invoke[DataAnalysisQueryBehaviorDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_behavior_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"host_name":    tea.StringValue(request.HostName),
			"os":           tea.StringValue(request.Os),
			"start_time":   tea.Int64Value(request.StartTime),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DataAnalysisQueryClientData(request *DataAnalysisQueryClientDataRequest) (_result *DataAnalysisQueryClientDataResponse, _err error) {
	return invoke[DataAnalysisQueryClientDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_client_data/",
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

func (client *Client) DataAnalysisQueryLiveRoom(request *DataAnalysisQueryLiveRoomRequest) (_result *DataAnalysisQueryLiveRoomResponse, _err error) {
	return invoke[DataAnalysisQueryLiveRoomResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_live_room/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"anchor_name": tea.StringValue(request.AnchorName),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DataAnalysisQueryPageData(request *DataAnalysisQueryPageDataRequest) (_result *DataAnalysisQueryPageDataResponse, _err error) {
	return invoke[DataAnalysisQueryPageDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_page_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"host_name":    tea.StringValue(request.HostName),
			"os":           tea.StringValue(request.Os),
			"start_time":   tea.Int64Value(request.StartTime),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DataAnalysisQueryRetentionData(request *DataAnalysisQueryRetentionDataRequest) (_result *DataAnalysisQueryRetentionDataResponse, _err error) {
	return invoke[DataAnalysisQueryRetentionDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_retention_data/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Query: map[string]interface{}{
			"end_time":     tea.Int64Value(request.EndTime),
			"host_name":    tea.StringValue(request.HostName),
			"os":           tea.StringValue(request.Os),
			"start_time":   tea.Int64Value(request.StartTime),
			"user_type":    tea.StringValue(request.UserType),
			"version_type": tea.StringValue(request.VersionType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DataAnalysisQuerySceneData(request *DataAnalysisQuerySceneDataRequest) (_result *DataAnalysisQuerySceneDataResponse, _err error) {
	return invoke[DataAnalysisQuerySceneDataResponse](client, request, &transport.Request{
		Method:  "GET",
		Path:    "/api/platform/v2/data_analysis/query_scene_data/",
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

func (client *Client) DataAnalysisQueryVideoData(request *DataAnalysisQueryVideoDataRequest) (_result *DataAnalysisQueryVideoDataResponse, _err error) {
	return invoke[DataAnalysisQueryVideoDataResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/platform/v2/data_analysis/query_video_data/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_short_id_list": tea.StringSliceValue(request.AwemeShortIdList),
			"end_time":            tea.Int64Value(request.EndTime),
			"host_name":           tea.StringValue(request.HostName),
			"item_id_list":        tea.StringSliceValue(request.ItemIdList),
			"open_item_id_list":   tea.StringSliceValue(request.OpenItemIdList),
			"query_bind_type":     tea.Int32Value(request.QueryBindType),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DefaultSet(request *DefaultSetRequest) (_result *DefaultSetResponse, _err error) {
	return invoke[DefaultSetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/default/set",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"card_type": request.CardType,
			"value":     tea.StringValue(request.Value),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DelRetainConsultCard(request *DelRetainConsultCardRequest) (_result *DelRetainConsultCardResponse, _err error) {
	return invoke[DelRetainConsultCardResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/im/del/retain_consult_card/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"card_id": tea.StringValue(request.CardId),
			"open_id": tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeleteAppTestRelation(request *DeleteAppTestRelationRequest) (_result *DeleteAppTestRelationResponse, _err error) {
	return invoke[DeleteAppTestRelationResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/industry/v1/solution/delete_app_test_relation/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"operator":    tea.StringValue(request.Operator),
			"ref_id_list": tea.StringSliceValue(request.RefIdList),
			"type":        tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeleteAwemeVideoKeyword(request *DeleteAwemeVideoKeywordRequest) (_result *DeleteAwemeVideoKeywordResponse, _err error) {
	return invoke[DeleteAwemeVideoKeywordResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/delete_aweme_video_keyword/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"keyword_id": tea.StringValue(request.KeywordId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeleteOrientedPlanTalent(request *DeleteOrientedPlanTalentRequest) (_result *DeleteOrientedPlanTalentResponse, _err error) {
	return invoke[DeleteOrientedPlanTalentResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/delete_oriented_plan_talent/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"douyin_id": tea.StringValue(request.DouyinId),
			"plan_id":   tea.Int64Value(request.PlanId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeletePurchaseInfo(request *DeletePurchaseInfoRequest) (_result *DeletePurchaseInfoResponse, _err error) {
	return invoke[DeletePurchaseInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/market/service/user/delete/purchase/info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
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

func (client *Client) DeleteSimpleQrBind(request *DeleteSimpleQrBindRequest) (_result *DeleteSimpleQrBindResponse, _err error) {
	return invoke[DeleteSimpleQrBindResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v2/capacity/delete_simple_qr_bind/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"qr_url": tea.StringValue(request.QrUrl),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DetailTrip(request *DetailTripRequest) (_result *DetailTripResponse, _err error) {
	return invoke[DetailTripResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/bill/detail/trip/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":      tea.StringValue(request.AccountId),
			"bill_date":       tea.StringValue(request.BillDate),
			"biz_type":        tea.IntValue(request.BizType),
			"cursor":          tea.StringValue(request.Cursor),
			"root_account_id": tea.StringValue(request.RootAccountId),
			"size":            tea.Int64Value(request.Size),
			"withdraw_id":     tea.StringValue(request.WithdrawId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DetailedQueryByOrder(request *DetailedQueryByOrderRequest) (_result *DetailedQueryByOrderResponse, _err error) {
	return invoke[DetailedQueryByOrderResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/ledger/detailed_query_by_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"order_ids":  tea.StringSliceValue(request.OrderIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCancelAuthOrder(request *DeveloperCancelAuthOrderRequest) (_result *DeveloperCancelAuthOrderResponse, _err error) {
	return invoke[DeveloperCancelAuthOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/cancel_auth_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id": tea.StringValue(request.AuthOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCheckDeductPeriod(request *DeveloperCheckDeductPeriodRequest) (_result *DeveloperCheckDeductPeriodResponse, _err error) {
	return invoke[DeveloperCheckDeductPeriodResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/check_deduct_period/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id":     tea.StringValue(request.AuthOrderId),
			"ext_auth_order_id": tea.StringValue(request.ExtAuthOrderId),
			"merchant_uid":      tea.StringValue(request.MerchantUid),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperClosePayOrder(request *DeveloperClosePayOrderRequest) (_result *DeveloperClosePayOrderResponse, _err error) {
	return invoke[DeveloperClosePayOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/close_pay_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"pay_order_id": tea.StringValue(request.PayOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCreateAuthOrder(request *DeveloperCreateAuthOrderRequest) (_result *DeveloperCreateAuthOrderResponse, _err error) {
	return invoke[DeveloperCreateAuthOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/create_auth_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"admissible_token":  tea.StringValue(request.AdmissibleToken),
			"notify_url":        tea.StringValue(request.NotifyUrl),
			"out_auth_order_no": tea.StringValue(request.OutAuthOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCreatePayOrder(request *DeveloperCreatePayOrderRequest) (_result *DeveloperCreatePayOrderResponse, _err error) {
	return invoke[DeveloperCreatePayOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/create_pay_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id":    tea.StringValue(request.AuthOrderId),
			"fee_detail_list":  request.FeeDetailList,
			"notify_url":       tea.StringValue(request.NotifyUrl),
			"out_pay_order_no": tea.StringValue(request.OutPayOrderNo),
			"total_amount":     tea.Int64Value(request.TotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCreateRefund(request *DeveloperCreateRefundRequest) (_result *DeveloperCreateRefundResponse, _err error) {
	return invoke[DeveloperCreateRefundResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/create_refund/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"fee_detail_list":     request.FeeDetailList,
			"notify_url":          tea.StringValue(request.NotifyUrl),
			"out_pay_refund_no":   tea.StringValue(request.OutPayRefundNo),
			"pay_order_id":        tea.StringValue(request.PayOrderId),
			"refund_reason":       tea.StringValue(request.RefundReason),
			"refund_total_amount": tea.Int64Value(request.RefundTotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCreateSignPay(request *DeveloperCreateSignPayRequest) (_result *DeveloperCreateSignPayResponse, _err error) {
	return invoke[DeveloperCreateSignPayResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/create_sign_pay/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id":       tea.StringValue(request.AuthOrderId),
			"expire_seconds":      tea.Int64Value(request.ExpireSeconds),
			"merchant_uid":        tea.StringValue(request.MerchantUid),
			"notify_url":          tea.StringValue(request.NotifyUrl),
			"out_pay_order_no":    tea.StringValue(request.OutPayOrderNo),
			"platform_pre_notify": tea.Int32Value(request.PlatformPreNotify),
			"total_amount":        tea.Int64Value(request.TotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperCreateSignRefund(request *DeveloperCreateSignRefundRequest) (_result *DeveloperCreateSignRefundResponse, _err error) {
	return invoke[DeveloperCreateSignRefundResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/create_sign_refund/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"notify_url":          tea.StringValue(request.NotifyUrl),
			"out_pay_refund_no":   tea.StringValue(request.OutPayRefundNo),
			"pay_order_id":        tea.StringValue(request.PayOrderId),
			"refund_reason":       tea.StringValue(request.RefundReason),
			"refund_total_amount": tea.Int64Value(request.RefundTotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperFinishAuthOrder(request *DeveloperFinishAuthOrderRequest) (_result *DeveloperFinishAuthOrderResponse, _err error) {
	return invoke[DeveloperFinishAuthOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/finish_auth_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id": tea.StringValue(request.AuthOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperFulfillPushStatus(request *DeveloperFulfillPushStatusRequest) (_result *DeveloperFulfillPushStatusResponse, _err error) {
	return invoke[DeveloperFulfillPushStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/fulfill_push_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"item_order_id_list": tea.StringSliceValue(request.ItemOrderIdList),
			"order_id":           tea.StringValue(request.OrderId),
			"to_status":          tea.StringValue(request.ToStatus),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperOrderQuery(request *DeveloperOrderQueryRequest) (_result *DeveloperOrderQueryResponse, _err error) {
	return invoke[DeveloperOrderQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/order_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":     tea.StringValue(request.OrderId),
			"out_order_no": tea.StringValue(request.OutOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperPreNotify(request *DeveloperPreNotifyRequest) (_result *DeveloperPreNotifyResponse, _err error) {
	return invoke[DeveloperPreNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/pre_notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"amount":            tea.Int64Value(request.Amount),
			"auth_order_id":     tea.StringValue(request.AuthOrderId),
			"ext_auth_order_id": tea.StringValue(request.ExtAuthOrderId),
			"merchant_uid":      tea.StringValue(request.MerchantUid),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQueryAdmissibleAuth(request *DeveloperQueryAdmissibleAuthRequest) (_result *DeveloperQueryAdmissibleAuthResponse, _err error) {
	return invoke[DeveloperQueryAdmissibleAuthResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_admissible_auth/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"merchant_uid": tea.StringValue(request.MerchantUid),
			"open_id":      tea.StringValue(request.OpenId),
			"scene":        tea.Int64Value(request.Scene),
			"service_id":   tea.StringValue(request.ServiceId),
			"total_amount": tea.Int64Value(request.TotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQueryAuthOrder(request *DeveloperQueryAuthOrderRequest) (_result *DeveloperQueryAuthOrderResponse, _err error) {
	return invoke[DeveloperQueryAuthOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_auth_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id":     tea.StringValue(request.AuthOrderId),
			"out_auth_order_no": tea.StringValue(request.OutAuthOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQueryCps(request *DeveloperQueryCpsRequest) (_result *DeveloperQueryCpsResponse, _err error) {
	return invoke[DeveloperQueryCpsResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/query_cps/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":     tea.StringValue(request.OrderId),
			"out_order_no": tea.StringValue(request.OutOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQueryPayOrder(request *DeveloperQueryPayOrderRequest) (_result *DeveloperQueryPayOrderResponse, _err error) {
	return invoke[DeveloperQueryPayOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_pay_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"out_pay_order_no": tea.StringValue(request.OutPayOrderNo),
			"pay_order_id":     tea.StringValue(request.PayOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQueryRefund(request *DeveloperQueryRefundRequest) (_result *DeveloperQueryRefundResponse, _err error) {
	return invoke[DeveloperQueryRefundResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_refund/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"out_pay_refund_no": tea.StringValue(request.OutPayRefundNo),
			"pay_refund_id":     tea.StringValue(request.PayRefundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQuerySignOrder(request *DeveloperQuerySignOrderRequest) (_result *DeveloperQuerySignOrderResponse, _err error) {
	return invoke[DeveloperQuerySignOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_sign_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id":     tea.StringValue(request.AuthOrderId),
			"out_auth_order_no": tea.StringValue(request.OutAuthOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQuerySignPay(request *DeveloperQuerySignPayRequest) (_result *DeveloperQuerySignPayResponse, _err error) {
	return invoke[DeveloperQuerySignPayResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_sign_pay/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"out_pay_order_no": tea.StringValue(request.OutPayOrderNo),
			"pay_order_id":     tea.StringValue(request.PayOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperQuerySignRefund(request *DeveloperQuerySignRefundRequest) (_result *DeveloperQuerySignRefundResponse, _err error) {
	return invoke[DeveloperQuerySignRefundResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/query_sign_refund/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"out_pay_refund_no": tea.StringValue(request.OutPayRefundNo),
			"pay_order_id":      tea.StringValue(request.PayOrderId),
			"pay_refund_id":     tea.StringValue(request.PayRefundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperRefundAuditCallback(request *DeveloperRefundAuditCallbackRequest) (_result *DeveloperRefundAuditCallbackResponse, _err error) {
	return invoke[DeveloperRefundAuditCallbackResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/refund_audit_callback/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"deny_message":        tea.StringValue(request.DenyMessage),
			"refund_audit_status": tea.Int32Value(request.RefundAuditStatus),
			"refund_id":           tea.StringValue(request.RefundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperRefundCreate(request *DeveloperRefundCreateRequest) (_result *DeveloperRefundCreateResponse, _err error) {
	return invoke[DeveloperRefundCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/refund_create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"cp_extra":            tea.StringValue(request.CpExtra),
			"item_order_detail":   request.ItemOrderDetail,
			"notify_url":          tea.StringValue(request.NotifyUrl),
			"order_entry_schema":  request.OrderEntrySchema,
			"order_id":            tea.StringValue(request.OrderId),
			"out_refund_no":       tea.StringValue(request.OutRefundNo),
			"refund_all":          tea.BoolValue(request.RefundAll),
			"refund_reason":       request.RefundReason,
			"refund_total_amount": tea.Int64Value(request.RefundTotalAmount),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperRefundQuery(request *DeveloperRefundQueryRequest) (_result *DeveloperRefundQueryResponse, _err error) {
	return invoke[DeveloperRefundQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/refund_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":      tea.StringValue(request.OrderId),
			"out_refund_no": tea.StringValue(request.OutRefundNo),
			"refund_id":     tea.StringValue(request.RefundId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperSettleCreate(request *DeveloperSettleCreateRequest) (_result *DeveloperSettleCreateResponse, _err error) {
	return invoke[DeveloperSettleCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/settle_create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":        tea.StringValue(request.AppId),
			"ext":           tea.StringValue(request.Ext),
			"item_order_id": tea.StringValue(request.ItemOrderId),
			"notify_url":    tea.StringValue(request.NotifyUrl),
			"out_order_no":  tea.StringValue(request.OutOrderNo),
			"out_settle_no": tea.StringValue(request.OutSettleNo),
			"settle_desc":   tea.StringValue(request.SettleDesc),
			"settle_params": tea.StringValue(request.SettleParams),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperSettleQuery(request *DeveloperSettleQueryRequest) (_result *DeveloperSettleQueryResponse, _err error) {
	return invoke[DeveloperSettleQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/settle_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":        tea.StringValue(request.AppId),
			"order_id":      tea.StringValue(request.OrderId),
			"out_order_no":  tea.StringValue(request.OutOrderNo),
			"out_settle_no": tea.StringValue(request.OutSettleNo),
			"settle_id":     tea.StringValue(request.SettleId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperTagQuery(request *DeveloperTagQueryRequest) (_result *DeveloperTagQueryResponse, _err error) {
	return invoke[DeveloperTagQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_basic/v1/developer/tag_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"goods_type": tea.IntValue(request.GoodsType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DeveloperTerminateSign(request *DeveloperTerminateSignRequest) (_result *DeveloperTerminateSignResponse, _err error) {
	return invoke[DeveloperTerminateSignResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade_auth/v1/developer/terminate_sign/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"auth_order_id": tea.StringValue(request.AuthOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DevtoolGetMountLegal(request *DevtoolGetMountLegalRequest) (_result *DevtoolGetMountLegalResponse, _err error) {
	return invoke[DevtoolGetMountLegalResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/devtool/get_mount_legal/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"micapp_id": tea.StringValue(request.MicappId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishBindGet(request *DishBindGetRequest) (_result *DishBindGetResponse, _err error) {
	return invoke[DishBindGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_dish/dish_bind/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"page_no":    tea.Int32Value(request.PageNo),
			"page_size":  tea.Int32Value(request.PageSize),
			"poi_id":     tea.Int64Value(request.PoiId),
			"product_id": tea.Int64Value(request.ProductId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishDraftGet(request *DishDraftGetRequest) (_result *DishDraftGetResponse, _err error) {
	return invoke[DishDraftGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_dish/dish/draft/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":        request.Base,
			"account_id":  tea.StringValue(request.AccountId),
			"dish_type":   tea.IntValue(request.DishType),
			"poi_id":      tea.StringValue(request.PoiId),
			"product_ids": tea.Int64ValueSlice(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishGroupQuery(request *DishGroupQueryRequest) (_result *DishGroupQueryResponse, _err error) {
	return invoke[DishGroupQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_dish/dish_group/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"group_type": tea.IntValue(request.GroupType),
			"page_no":    tea.Int32Value(request.PageNo),
			"page_size":  tea.Int32Value(request.PageSize),
			"poi_id":     tea.Int64Value(request.PoiId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishGroupSave(request *DishGroupSaveRequest) (_result *DishGroupSaveResponse, _err error) {
	return invoke[DishGroupSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish_group/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"group_info": request.GroupInfo,
			"group_type": tea.IntValue(request.GroupType),
			"poi_id":     tea.Int64Value(request.PoiId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishGroupSort(request *DishGroupSortRequest) (_result *DishGroupSortResponse, _err error) {
	return invoke[DishGroupSortResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish_group/sort/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"group_type": tea.IntValue(request.GroupType),
			"poi_id":     tea.Int64Value(request.PoiId),
			"sort_list":  tea.Int64ValueSlice(request.SortList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishOnlineGet(request *DishOnlineGetRequest) (_result *DishOnlineGetResponse, _err error) {
	return invoke[DishOnlineGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/dish/online/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":        request.Base,
			"account_id":  tea.StringValue(request.AccountId),
			"dish_type":   tea.IntValue(request.DishType),
			"poi_id":      tea.Int64Value(request.PoiId),
			"product_ids": tea.Int64ValueSlice(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishOperateStatus(request *DishOperateStatusRequest) (_result *DishOperateStatusResponse, _err error) {
	return invoke[DishOperateStatusResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/dish/operate_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":             request.Base,
			"account_id":       tea.StringValue(request.AccountId),
			"dy_poi_id":        tea.Int64Value(request.DyPoiId),
			"merchant_operate": tea.BoolValue(request.MerchantOperate),
			"operate_type":     tea.IntValue(request.OperateType),
			"product_ids":      tea.Int64ValueSlice(request.ProductIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishSaveOutId(request *DishSaveOutIdRequest) (_result *DishSaveOutIdResponse, _err error) {
	return invoke[DishSaveOutIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish/save_out_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":            request.Base,
			"account_id":      tea.StringValue(request.AccountId),
			"feed_out_id_map": request.FeedOutIdMap,
			"out_id":          tea.StringValue(request.OutId),
			"poi_id":          tea.Int64Value(request.PoiId),
			"product_id":      tea.Int64Value(request.ProductId),
			"sku_out_id_map":  request.SkuOutIdMap,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishSaveSubmit(request *DishSaveSubmitRequest) (_result *DishSaveSubmitResponse, _err error) {
	return invoke[DishSaveSubmitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish/save_submit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":             request.Base,
			"account_id":       tea.StringValue(request.AccountId),
			"add_dish_groups":  request.AddDishGroups,
			"apply_date":       request.ApplyDate,
			"attributes":       request.Attributes,
			"category_id":      tea.Int64Value(request.CategoryId),
			"delivery_method":  tea.IntValueSlice(request.DeliveryMethod),
			"dish_description": tea.StringValue(request.DishDescription),
			"dish_detail_info": request.DishDetailInfo,
			"dish_groups":      request.DishGroups,
			"image_list":       request.ImageList,
			"pois":             request.Pois,
			"product_id":       tea.Int64Value(request.ProductId),
			"product_name":     tea.StringValue(request.ProductName),
			"product_type":     tea.IntValue(request.ProductType),
			"settle_type":      tea.IntValue(request.SettleType),
			"skus":             request.Skus,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishSort(request *DishSortRequest) (_result *DishSortResponse, _err error) {
	return invoke[DishSortResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish/sort/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"group_id":   tea.Int64Value(request.GroupId),
			"group_type": tea.IntValue(request.GroupType),
			"poi_id":     tea.Int64Value(request.PoiId),
			"sort_list":  tea.Int64ValueSlice(request.SortList),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishSyncTaskCreate(request *DishSyncTaskCreateRequest) (_result *DishSyncTaskCreateResponse, _err error) {
	return invoke[DishSyncTaskCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_dish/dish_sync_task/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":        request.Base,
			"account_id":  tea.StringValue(request.AccountId),
			"poi_id":      tea.Int64ValueSlice(request.PoiId),
			"product_ids": tea.Int64ValueSlice(request.ProductIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishSyncTaskRecord(request *DishSyncTaskRecordRequest) (_result *DishSyncTaskRecordResponse, _err error) {
	return invoke[DishSyncTaskRecordResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_dish/dish_sync_task/record/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"page_no":    tea.Int32Value(request.PageNo),
			"page_size":  tea.Int32Value(request.PageSize),
			"task_id":    tea.Int64Value(request.TaskId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DishUpdateStock(request *DishUpdateStockRequest) (_result *DishUpdateStockResponse, _err error) {
	return invoke[DishUpdateStockResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/x_goods/dish/update_stock/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":             request.Base,
			"account_id":       tea.StringValue(request.AccountId),
			"dish_sku_stocks":  request.DishSkuStocks,
			"dy_poi_id":        tea.Int64Value(request.DyPoiId),
			"merchant_operate": tea.BoolValue(request.MerchantOperate),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DouyinCreateInteractTask(request *DouyinCreateInteractTaskRequest) (_result *DouyinCreateInteractTaskResponse, _err error) {
	return invoke[DouyinCreateInteractTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/douyin/create_interact_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":         tea.StringValue(request.AppId),
			"end_time":       tea.Int64Value(request.EndTime),
			"interact_rules": request.InteractRules,
			"max_count":      tea.Int64Value(request.MaxCount),
			"mount_link":     tea.StringValue(request.MountLink),
			"publish_type":   tea.Int32ValueSlice(request.PublishType),
			"start_time":     tea.Int64Value(request.StartTime),
			"tags":           tea.StringSliceValue(request.Tags),
			"task_type":      tea.Int32Value(request.TaskType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DouyinCreateTask(request *DouyinCreateTaskRequest) (_result *DouyinCreateTaskResponse, _err error) {
	return invoke[DouyinCreateTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/douyin/create_task/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":       tea.StringValue(request.AppId),
			"end_time":     tea.Int64Value(request.EndTime),
			"mount_link":   tea.StringValue(request.MountLink),
			"publish_type": tea.Int32ValueSlice(request.PublishType),
			"rule_type":    tea.Int32Value(request.RuleType),
			"start_time":   tea.Int64Value(request.StartTime),
			"tags":         tea.StringSliceValue(request.Tags),
			"target_count": tea.Int64Value(request.TargetCount),
			"task_type":    tea.Int32Value(request.TaskType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DouyinQueryUserTask(request *DouyinQueryUserTaskRequest) (_result *DouyinQueryUserTaskResponse, _err error) {
	return invoke[DouyinQueryUserTaskResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/douyin/query_user_task/",
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

func (client *Client) DraftGet(request *DraftGetRequest) (_result *DraftGetResponse, _err error) {
	return invoke[DraftGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/product/draft/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":           request.Base,
			"account_id":     tea.StringValue(request.AccountId),
			"get_draft_type": tea.IntValue(request.GetDraftType),
			"product_ids":    tea.StringSliceValue(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DraftList(request *DraftListRequest) (_result *DraftListResponse, _err error) {
	return invoke[DraftListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_dish/dish/draft/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"dish_type":  tea.IntValue(request.DishType),
			"page_no":    tea.Int32Value(request.PageNo),
			"page_size":  tea.Int32Value(request.PageSize),
			"poi_id":     tea.Int64Value(request.PoiId),
			"status":     tea.IntValueSlice(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DramaOverall(request *DramaOverallRequest) (_result *DramaOverallResponse, _err error) {
	return invoke[DramaOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/drama/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) DylkOpenIdGet(request *DylkOpenIdGetRequest) (_result *DylkOpenIdGetResponse, _err error) {
	return invoke[DylkOpenIdGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/open/common_biz/dylk_open_id/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":   tea.StringValue(request.AccountId),
			"open_id_list": tea.StringSliceValue(request.OpenIdList),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FailDataGet(request *FailDataGetRequest) (_result *FailDataGetResponse, _err error) {
	return invoke[FailDataGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/live_data/task/fail_data/get",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"appid":     tea.StringValue(request.Appid),
			"msg_type":  tea.StringValue(request.MsgType),
			"page_num":  tea.Int32Value(request.PageNum),
			"page_size": tea.Int32Value(request.PageSize),
			"roomid":    tea.StringValue(request.Roomid),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FansCheck(request *FansCheckRequest) (_result *FansCheckResponse, _err error) {
	return invoke[FansCheckResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/fans/check/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"follower_open_id": tea.StringValue(request.FollowerOpenId),
			"open_id":          tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FansClubGetInfo(request *FansClubGetInfoRequest) (_result *FansClubGetInfoResponse, _err error) {
	return invoke[FansClubGetInfoResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/live_data/fans_club/get_info",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"anchor_openid": tea.StringValue(request.AnchorOpenid),
			"roomid":        tea.Int64Value(request.Roomid),
			"user_openids":  tea.StringSliceValue(request.UserOpenids),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FansComment(request *FansCommentRequest) (_result *FansCommentResponse, _err error) {
	return invoke[FansCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/fans/comment/",
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

func (client *Client) FansCreate(request *FansCreateRequest) (_result *FansCreateResponse, _err error) {
	return invoke[FansCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/group/fans/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"active_fans":       tea.Int64Value(request.ActiveFans),
			"allow_invite":      tea.Int64Value(request.AllowInvite),
			"avatar_uri":        tea.StringValue(request.AvatarUri),
			"description":       tea.StringValue(request.Description),
			"fans_limit":        tea.Int64Value(request.FansLimit),
			"group_name":        tea.StringValue(request.GroupName),
			"group_type":        tea.Int64Value(request.GroupType),
			"item_auto_sync":    tea.Int64Value(request.ItemAutoSync),
			"live_auto_sync":    tea.Int64Value(request.LiveAutoSync),
			"open_audit_switch": tea.Int64Value(request.OpenAuditSwitch),
			"relation_type":     tea.Int64Value(request.RelationType),
			"show_at_profile":   tea.Int64Value(request.ShowAtProfile),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FansFavourite(request *FansFavouriteRequest) (_result *FansFavouriteResponse, _err error) {
	return invoke[FansFavouriteResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/fans/favourite/",
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

func (client *Client) FansList(request *FansListRequest) (_result *FansListResponse, _err error) {
	return invoke[FansListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/im/group/fans/list/",
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

func (client *Client) FansSource(request *FansSourceRequest) (_result *FansSourceResponse, _err error) {
	return invoke[FansSourceResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/fans/source/",
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

func (client *Client) FeedbackBatchReply(request *FeedbackBatchReplyRequest) (_result *FeedbackBatchReplyResponse, _err error) {
	return invoke[FeedbackBatchReplyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/feedback/batch_reply/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":      tea.StringValue(request.Appid),
			"reply_list": request.ReplyList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FeedbackGetFeedbackList(request *FeedbackGetFeedbackListRequest) (_result *FeedbackGetFeedbackListResponse, _err error) {
	return invoke[FeedbackGetFeedbackListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/feedback/get_feedback_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":         tea.StringValue(request.Appid),
			"current_page":  tea.Int32Value(request.CurrentPage),
			"dealt_status":  tea.IntValueSlice(request.DealtStatus),
			"end_time":      tea.Int64Value(request.EndTime),
			"feedback_id":   tea.Int64Value(request.FeedbackId),
			"feedback_type": tea.IntValue(request.FeedbackType),
			"open_id":       tea.StringValue(request.OpenId),
			"page_size":     tea.Int32Value(request.PageSize),
			"start_time":    tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FeedbackGetReplyResult(request *FeedbackGetReplyResultRequest) (_result *FeedbackGetReplyResultResponse, _err error) {
	return invoke[FeedbackGetReplyResultResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/feedback/get_reply_result/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"appid":       tea.StringValue(request.Appid),
			"feedback_id": tea.Int64Value(request.FeedbackId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FileUpload(request *FileUploadRequest) (_result *FileUploadResponse, _err error) {
	return invoke[FileUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/entrance/file/upload/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"file_content_base64": tea.StringValue(request.FileContentBase64),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodNew(request *FoodNewRequest) (_result *FoodNewResponse, _err error) {
	return invoke[FoodNewResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/food/new/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodOverall(request *FoodOverallRequest) (_result *FoodOverallResponse, _err error) {
	return invoke[FoodOverallResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/food/overall/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodShop(request *FoodShopRequest) (_result *FoodShopResponse, _err error) {
	return invoke[FoodShopResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/food/shop/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodTutorial(request *FoodTutorialRequest) (_result *FoodTutorialResponse, _err error) {
	return invoke[FoodTutorialResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/food/tutorial/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodorderPoistockSave(request *FoodorderPoistockSaveRequest) (_result *FoodorderPoistockSaveResponse, _err error) {
	return invoke[FoodorderPoistockSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/poistock/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":               request.Base,
			"account_id":         tea.Int64Value(request.AccountId),
			"ignore_fail_poi":    tea.BoolValue(request.IgnoreFailPoi),
			"out_product_id":     tea.StringValue(request.OutProductId),
			"poi_sku_stock_list": request.PoiSkuStockList,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodorderProductOperate(request *FoodorderProductOperateRequest) (_result *FoodorderProductOperateResponse, _err error) {
	return invoke[FoodorderProductOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/product/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.Int64Value(request.AccountId),
			"op":         tea.IntValue(request.Op),
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.Int64Value(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FoodorderProductSave(request *FoodorderProductSaveRequest) (_result *FoodorderProductSaveResponse, _err error) {
	return invoke[FoodorderProductSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/foodorder/product/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":      tea.Int64Value(request.AccountId),
			"ignore_fail_poi": tea.BoolValue(request.IgnoreFailPoi),
			"product":         request.Product,
			"supports_laike":  tea.BoolValue(request.SupportsLaike),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfillmentDeliveryVerify(request *FulfillmentDeliveryVerifyRequest) (_result *FulfillmentDeliveryVerifyResponse, _err error) {
	return invoke[FulfillmentDeliveryVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/fulfillment/delivery_verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificates":    request.Certificates,
			"delivery_extra":  request.DeliveryExtra,
			"encrypted_codes": tea.StringSliceValue(request.EncryptedCodes),
			"order_id":        tea.StringValue(request.OrderId),
			"poi_info":        tea.StringValue(request.PoiInfo),
			"verify_token":    tea.StringValue(request.VerifyToken),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfillmentOrderCanUse(request *FulfillmentOrderCanUseRequest) (_result *FulfillmentOrderCanUseResponse, _err error) {
	return invoke[FulfillmentOrderCanUseResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade/v2/fulfillment/order_can_use/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id": tea.StringValue(request.OrderId),
			"poi_id":   tea.StringValue(request.PoiId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfillmentPushDelivery(request *FulfillmentPushDeliveryRequest) (_result *FulfillmentPushDeliveryResponse, _err error) {
	return invoke[FulfillmentPushDeliveryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/fulfillment/push_delivery/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"book_id":         tea.StringValue(request.BookId),
			"delivery_extra":  request.DeliveryExtra,
			"delivery_status": tea.IntValue(request.DeliveryStatus),
			"item_book_list":  request.ItemBookList,
			"item_order_list": request.ItemOrderList,
			"out_order_no":    tea.StringValue(request.OutOrderNo),
			"poi_info":        tea.StringValue(request.PoiInfo),
			"token":           tea.StringValue(request.Token),
			"use_all":         tea.BoolValue(request.UseAll),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfillmentQueryUserCertificates(request *FulfillmentQueryUserCertificatesRequest) (_result *FulfillmentQueryUserCertificatesResponse, _err error) {
	return invoke[FulfillmentQueryUserCertificatesResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade/v2/fulfillment/query_user_certificates/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"biz_type":       tea.Int32Value(request.BizType),
			"open_id":        tea.StringValue(request.OpenId),
			"page":           tea.Int32Value(request.Page),
			"page_dimension": tea.IntValue(request.PageDimension),
			"page_size":      tea.Int32Value(request.PageSize),
			"poi_id":         tea.StringValue(request.PoiId),
			"time_range":     request.TimeRange,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfillmentVerifyCancel(request *FulfillmentVerifyCancelRequest) (_result *FulfillmentVerifyCancelResponse, _err error) {
	return invoke[FulfillmentVerifyCancelResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/trade/v2/fulfillment/verify_cancel/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"certificate_id": tea.StringValue(request.CertificateId),
			"order_id":       tea.StringValue(request.OrderId),
			"verify_id":      tea.StringValue(request.VerifyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfilmentCertificateCancel(request *FulfilmentCertificateCancelRequest) (_result *FulfilmentCertificateCancelResponse, _err error) {
	return invoke[FulfilmentCertificateCancelResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/compre_retail/fulfilment/certificate/cancel/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"batch_cancel_info":       request.BatchCancelInfo,
			"batch_cancel_info_list":  request.BatchCancelInfoList,
			"cancel_token":            tea.StringValue(request.CancelToken),
			"certificate_id":          tea.StringValue(request.CertificateId),
			"shop_order_id":           tea.StringValue(request.ShopOrderId),
			"times_card_cancel_count": tea.Int64Value(request.TimesCardCancelCount),
			"verify_id":               tea.StringValue(request.VerifyId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfilmentCertificatePrepare(request *FulfilmentCertificatePrepareRequest) (_result *FulfilmentCertificatePrepareResponse, _err error) {
	return invoke[FulfilmentCertificatePrepareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/compre_retail/fulfilment/certificate/prepare/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"code":           tea.StringValue(request.Code),
			"encrypted_data": tea.StringValue(request.EncryptedData),
			"poi_id":         tea.StringValue(request.PoiId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfilmentCertificateQuery(request *FulfilmentCertificateQueryRequest) (_result *FulfilmentCertificateQueryResponse, _err error) {
	return invoke[FulfilmentCertificateQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/compre_retail/fulfilment/certificate/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"encrypted_code": tea.StringValue(request.EncryptedCode),
			"order_id":       tea.StringValue(request.OrderId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfilmentCertificateVerify(request *FulfilmentCertificateVerifyRequest) (_result *FulfilmentCertificateVerifyResponse, _err error) {
	return invoke[FulfilmentCertificateVerifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/compre_retail/fulfilment/certificate/verify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":          tea.StringValue(request.AccountId),
			"code_with_time_list": request.CodeWithTimeList,
			"codes":               tea.StringSliceValue(request.Codes),
			"encrypted_codes":     tea.StringSliceValue(request.EncryptedCodes),
			"order_id":            tea.StringValue(request.OrderId),
			"poi_id":              tea.StringValue(request.PoiId),
			"verify_extra":        request.VerifyExtra,
			"verify_sign_list":    tea.StringSliceValue(request.VerifySignList),
			"verify_token":        tea.StringValue(request.VerifyToken),
			"voucher":             request.Voucher,
			"vouchers":            request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FulfilmentCreateCallback(request *FulfilmentCreateCallbackRequest) (_result *FulfilmentCreateCallbackResponse, _err error) {
	return invoke[FulfilmentCreateCallbackResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/compre_retail/fulfilment/create/callback/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"additional_info_list": request.AdditionalInfoList,
			"certificates":         request.Certificates,
			"codes":                tea.StringSliceValue(request.Codes),
			"fail_reason":          tea.StringValue(request.FailReason),
			"fail_reason_desc":     tea.StringValue(request.FailReasonDesc),
			"order_id":             tea.StringValue(request.OrderId),
			"result":               tea.Int64Value(request.Result),
			"third_order_id":       tea.StringValue(request.ThirdOrderId),
			"voucher":              request.Voucher,
			"vouchers":             request.Vouchers,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FunctionConfigQueryStatus(request *FunctionConfigQueryStatusRequest) (_result *FunctionConfigQueryStatusResponse, _err error) {
	return invoke[FunctionConfigQueryStatusResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/developer_toolbox/image_material/function_config/query_status/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"function_id": tea.StringSliceValue(request.FunctionId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) FundBills(request *FundBillsRequest) (_result *FundBillsResponse, _err error) {
	return invoke[FundBillsResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v3/fund/bills/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_type": tea.StringValue(request.AccountType),
			"app_id":       tea.StringValue(request.AppId),
			"bill_date":    tea.StringValue(request.BillDate),
			"merchant_id":  tea.StringValue(request.MerchantId),
			"payment_type": tea.StringValue(request.PaymentType),
			"trade_type":   tea.StringValue(request.TradeType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GameConsole(request *GameConsoleRequest) (_result *GameConsoleResponse, _err error) {
	return invoke[GameConsoleResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/game/console/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GameInf(request *GameInfRequest) (_result *GameInfResponse, _err error) {
	return invoke[GameInfResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/game/inf/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetAgencyUserBindRecord(request *GetAgencyUserBindRecordRequest) (_result *GetAgencyUserBindRecordResponse, _err error) {
	return invoke[GetAgencyUserBindRecordResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v1/taskbox/get_agency_user_bind_record/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"agency_talent_uid": tea.StringValue(request.AgencyTalentUid),
			"agent_id":          tea.Int64Value(request.AgentId),
			"douyin_id":         tea.StringValue(request.DouyinId),
			"page_no":           tea.Int32Value(request.PageNo),
			"page_size":         tea.Int32Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetAllCardsForUser(request *GetAllCardsForUserRequest) (_result *GetAllCardsForUserResponse, _err error) {
	return invoke[GetAllCardsForUserResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/card/get_all_cards_for_user",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetAwemeBindTemplateInfo(request *GetAwemeBindTemplateInfoRequest) (_result *GetAwemeBindTemplateInfoResponse, _err error) {
	return invoke[GetAwemeBindTemplateInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/get_aweme_bind_template_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_id":    tea.StringValue(request.AwemeId),
			"template_id": tea.Int64ValueSlice(request.TemplateId),
			"type":        tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetAwemeBindTemplateList(request *GetAwemeBindTemplateListRequest) (_result *GetAwemeBindTemplateListResponse, _err error) {
	return invoke[GetAwemeBindTemplateListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/capacity/get_aweme_bind_template_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_id":      tea.StringValue(request.AwemeId),
			"capacity_list": tea.StringSliceValue(request.CapacityList),
			"type":          tea.StringValue(request.Type),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetAwemeRelationBindQrcode(request *GetAwemeRelationBindQrcodeRequest) (_result *GetAwemeRelationBindQrcodeResponse, _err error) {
	return invoke[GetAwemeRelationBindQrcodeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/capacity/get_aweme_relation_bind_qrcode/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"capacity_list": tea.StringSliceValue(request.CapacityList),
			"co_subject":    tea.BoolValue(request.CoSubject),
			"type":          tea.StringValue(request.Type),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetBillDownloadUrl(request *GetBillDownloadUrlRequest) (_result *GetBillDownloadUrlResponse, _err error) {
	return invoke[GetBillDownloadUrlResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/get_bill_download_url/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":    tea.StringValue(request.AppId),
			"bill_date": tea.StringValue(request.BillDate),
			"bill_type": tea.StringValue(request.BillType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetCouponMetaStatistics(request *GetCouponMetaStatisticsRequest) (_result *GetCouponMetaStatisticsResponse, _err error) {
	return invoke[GetCouponMetaStatisticsResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/get_coupon_meta_statistics/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"activity_id":    tea.StringValue(request.ActivityId),
			"app_id":         tea.StringValue(request.AppId),
			"coupon_meta_id": tea.StringValue(request.CouponMetaId),
			"talent_account": tea.StringValue(request.TalentAccount),
			"talent_open_id": tea.StringValue(request.TalentOpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetCouponReceiveInfo(request *GetCouponReceiveInfoRequest) (_result *GetCouponReceiveInfoResponse, _err error) {
	return invoke[GetCouponReceiveInfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/promotion/v1/coupon/get_coupon_receive_info/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":        tea.StringValue(request.AppId),
			"coupon_id":     tea.StringValue(request.CouponId),
			"coupon_status": tea.IntValue(request.CouponStatus),
			"open_id":       tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetRetainConsultCard(request *GetRetainConsultCardRequest) (_result *GetRetainConsultCardResponse, _err error) {
	return invoke[GetRetainConsultCardResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/im/get/retain_consult_card/",
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

func (client *Client) GetUserEncryptKey(request *GetUserEncryptKeyRequest) (_result *GetUserEncryptKeyResponse, _err error) {
	return invoke[GetUserEncryptKeyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/get_user_encrypt_key/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":     tea.StringValue(request.AppId),
			"open_id":    tea.StringValue(request.OpenId),
			"sig_method": tea.StringValue(request.SigMethod),
			"signature":  tea.StringValue(request.Signature),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetUserGroupTag(request *GetUserGroupTagRequest) (_result *GetUserGroupTagResponse, _err error) {
	return invoke[GetUserGroupTagResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/apps/group_tag/get_user_group_tag",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"mp_id":   tea.StringValue(request.MpId),
			"open_id": tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GetUserVideoList(request *GetUserVideoListRequest) (_result *GetUserVideoListResponse, _err error) {
	return invoke[GetUserVideoListResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/v1/video/self/get_user_video_list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"aweme_id": tea.StringValue(request.AwemeId),
			"count":    tea.IntValue(request.Count),
			"cursor":   tea.Int64Value(request.Cursor),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GiftReceiveReward(request *GiftReceiveRewardRequest) (_result *GiftReceiveRewardResponse, _err error) {
	return invoke[GiftReceiveRewardResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/mgplatform/api/gift/receive_reward",
		Host:        client.GetHost("minigame.zijieapi.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"env_type":  tea.StringValue(request.EnvType),
			"gift_code": tea.StringValue(request.GiftCode),
			"open_id":   tea.StringValue(request.OpenId),
			"uuid":      tea.StringValue(request.Uuid),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsProductDraftGet(request *GoodsProductDraftGetRequest) (_result *GoodsProductDraftGetResponse, _err error) {
	return invoke[GoodsProductDraftGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/product/draft/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"get_draft_type": tea.IntValue(request.GetDraftType),
			"need_poi":       tea.BoolValue(request.NeedPoi),
			"out_ids":        tea.StringSliceValue(request.OutIds),
			"product_ids":    tea.StringSliceValue(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsProductFreeAudit(request *GoodsProductFreeAuditRequest) (_result *GoodsProductFreeAuditResponse, _err error) {
	return invoke[GoodsProductFreeAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/product/free_audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":          request.Base,
			"account_id":    tea.StringValue(request.AccountId),
			"actual_amount": tea.Int64Value(request.ActualAmount),
			"out_id":        tea.StringValue(request.OutId),
			"product_id":    tea.StringValue(request.ProductId),
			"sold_end_time": tea.Int64Value(request.SoldEndTime),
			"stock_qty":     tea.Int64Value(request.StockQty),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsProductOperate(request *GoodsProductOperateRequest) (_result *GoodsProductOperateResponse, _err error) {
	return invoke[GoodsProductOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/product/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"op_type":    tea.IntValue(request.OpType),
			"out_id":     tea.StringValue(request.OutId),
			"product_id": tea.StringValue(request.ProductId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsSkuBatchSave(request *GoodsSkuBatchSaveRequest) (_result *GoodsSkuBatchSaveResponse, _err error) {
	return invoke[GoodsSkuBatchSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/goods/sku/batch_save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"Base":           request.Base,
			"account_id":     tea.StringValue(request.AccountId),
			"product_id":     tea.StringValue(request.ProductId),
			"product_out_id": tea.StringValue(request.ProductOutId),
			"skus":           request.Skus,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsSpuOperate(request *GoodsSpuOperateRequest) (_result *GoodsSpuOperateResponse, _err error) {
	return invoke[GoodsSpuOperateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v2/goods/spu/operate/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
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

func (client *Client) GoodsSpuSave(request *GoodsSpuSaveRequest) (_result *GoodsSpuSaveResponse, _err error) {
	return invoke[GoodsSpuSaveResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v2/goods/spu/save/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.Int64Value(request.AccountId),
			"spu":        request.Spu,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GoodsTemplateGet(request *GoodsTemplateGetRequest) (_result *GoodsTemplateGetResponse, _err error) {
	return invoke[GoodsTemplateGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/template/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"category_id":      tea.StringValue(request.CategoryId),
			"open_biz_type":    tea.IntValue(request.OpenBizType),
			"product_sub_type": tea.IntValue(request.ProductSubType),
			"product_type":     tea.IntValue(request.ProductType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) GroupCount(request *GroupCountRequest) (_result *GroupCountResponse, _err error) {
	return invoke[GroupCountResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/im/group/count/",
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

func (client *Client) HermesTradeOrderQuery(request *HermesTradeOrderQueryRequest) (_result *HermesTradeOrderQueryResponse, _err error) {
	return invoke[HermesTradeOrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/hermes/trade/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":              tea.StringValue(request.AccountId),
			"create_order_end_time":   tea.Int64Value(request.CreateOrderEndTime),
			"create_order_start_time": tea.Int64Value(request.CreateOrderStartTime),
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

func (client *Client) HotelOrderConfirm(request *HotelOrderConfirmRequest) (_result *HotelOrderConfirmResponse, _err error) {
	return invoke[HotelOrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/hotel/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"confirm_result": request.ConfirmResult,
			"order_id":       tea.StringValue(request.OrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) HotelPoiQuery(request *HotelPoiQueryRequest) (_result *HotelPoiQueryResponse, _err error) {
	return invoke[HotelPoiQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/poi/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"page_index": tea.Int32Value(request.PageIndex),
			"page_size":  tea.Int32Value(request.PageSize),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) HotelSavepresale(request *HotelSavepresaleRequest) (_result *HotelSavepresaleResponse, _err error) {
	return invoke[HotelSavepresaleResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/savepresale/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":   tea.StringValue(request.AccountId),
			"presale_info": request.PresaleInfo,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) HotelUserUpdate(request *HotelUserUpdateRequest) (_result *HotelUserUpdateResponse, _err error) {
	return invoke[HotelUserUpdateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/member/hotel/user/update/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":         tea.StringValue(request.AccountId),
			"effect_end_time":    tea.Int64Value(request.EffectEndTime),
			"effect_start_time":  tea.Int64Value(request.EffectStartTime),
			"email":              tea.StringValue(request.Email),
			"first_name_cn":      tea.StringValue(request.FirstNameCn),
			"id_number":          tea.StringValue(request.IdNumber),
			"is_new_member":      tea.Int64Value(request.IsNewMember),
			"last_name_cn":       tea.StringValue(request.LastNameCn),
			"level_keep_nights":  tea.Int32Value(request.LevelKeepNights),
			"level_up_nights":    tea.Int32Value(request.LevelUpNights),
			"member_card_id":     tea.StringValue(request.MemberCardId),
			"mobile":             tea.StringValue(request.Mobile),
			"name":               tea.StringValue(request.Name),
			"noshow":             tea.Int32Value(request.Noshow),
			"open_id":            tea.StringValue(request.OpenId),
			"phone":              tea.StringValue(request.Phone),
			"points_amount_cent": tea.Int64Value(request.PointsAmountCent),
			"points_value":       tea.Int64Value(request.PointsValue),
			"room_nights":        tea.Int64Value(request.RoomNights),
			"stay":               tea.Int64Value(request.Stay),
			"update_time":        tea.Int64Value(request.UpdateTime),
			"user_level":         tea.Int64Value(request.UserLevel),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImDelAppletTemplate(request *ImDelAppletTemplateRequest) (_result *ImDelAppletTemplateResponse, _err error) {
	return invoke[ImDelAppletTemplateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/im/del_applet_template/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"card_template_id": tea.StringValue(request.CardTemplateId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImGetAppletTemplate(request *ImGetAppletTemplateRequest) (_result *ImGetAppletTemplateResponse, _err error) {
	return invoke[ImGetAppletTemplateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/im/get_applet_template/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"card_template_id": tea.StringValue(request.CardTemplateId),
			"count":            tea.Int32Value(request.Count),
			"cursor":           tea.Int64Value(request.Cursor),
			"status":           tea.Int32Value(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImSetAppletTemplate(request *ImSetAppletTemplateRequest) (_result *ImSetAppletTemplateResponse, _err error) {
	return invoke[ImSetAppletTemplateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/im/set_applet_template/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":           tea.StringValue(request.AppId),
			"card_template_id": tea.StringValue(request.CardTemplateId),
			"card_type":        tea.Int32Value(request.CardType),
			"content":          tea.StringValue(request.Content),
			"media_id":         tea.StringValue(request.MediaId),
			"title":            tea.StringValue(request.Title),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImageMaterialUpload(request *ImageMaterialUploadRequest) (_result *ImageMaterialUploadResponse, _err error) {
	return invoke[ImageMaterialUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/developer_toolbox/image_material/upload/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"image_material_url": tea.StringValue(request.ImageMaterialUrl),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImagexClientUpload(request *ImagexClientUploadRequest) (_result *ImagexClientUploadResponse, _err error) {
	return invoke[ImagexClientUploadResponse](client, request, &transport.Request{
		Method:  "POST",
		Path:    "/tool/imagex/client_upload/",
		Host:    client.GetHost("open.douyin.com"),
		Headers: request.Header,
		Body: map[string]interface{}{
			"image": request.Image,
		},
		Encoding:    transport.EncodingFileForm,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImportCardAdd(request *ImportCardAddRequest) (_result *ImportCardAddResponse, _err error) {
	return invoke[ImportCardAddResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/app/v1/import_card/add/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchorid_list": tea.Int64ValueSlice(request.AnchoridList),
			"item_map":      request.ItemMap,
			"start_page":    tea.StringValue(request.StartPage),
			"template_code": tea.Int32Value(request.TemplateCode),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ImportCardGet(request *ImportCardGetRequest) (_result *ImportCardGetResponse, _err error) {
	return invoke[ImportCardGetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/aweme/app/v1/import_card/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"anchorid":   tea.Int64Value(request.Anchorid),
			"start_page": tea.StringValue(request.StartPage),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) InfoMatch(request *InfoMatchRequest) (_result *InfoMatchResponse, _err error) {
	return invoke[InfoMatchResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/info/match/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":       tea.StringValue(request.AccountId),
			"active":           tea.BoolValue(request.Active),
			"hotel_base_info":  request.HotelBaseInfo,
			"hotel_facilities": request.HotelFacilities,
			"hotel_id":         tea.StringValue(request.HotelId),
			"hotel_images":     request.HotelImages,
			"hotel_policy":     request.HotelPolicy,
			"out_hotel_id":     tea.StringValue(request.OutHotelId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) InitVideoPartUpload(request *InitVideoPartUploadRequest) (_result *InitVideoPartUploadResponse, _err error) {
	return invoke[InitVideoPartUploadResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/douyin/v1/video/init_video_part_upload/",
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

func (client *Client) InventoryPush(request *InventoryPushRequest) (_result *InventoryPushResponse, _err error) {
	return invoke[InventoryPushResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/promotion/inventory/push/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":          tea.StringValue(request.AccountId),
			"applicable_date":     request.ApplicableDate,
			"applicable_resource": request.ApplicableResource,
			"inventory":           tea.Int64Value(request.Inventory),
			"promotion_id":        tea.StringValue(request.PromotionId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) InvoiceNotify(request *InvoiceNotifyRequest) (_result *InvoiceNotifyResponse, _err error) {
	return invoke[InvoiceNotifyResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/hotel/invoice/notify/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":     tea.StringValue(request.AccountId),
			"apply_id":       tea.StringValue(request.ApplyId),
			"invoice_status": tea.IntValue(request.InvoiceStatus),
			"invoice_urls":   tea.StringSliceValue(request.InvoiceUrls),
			"message":        tea.StringValue(request.Message),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemBcGetBase(request *ItemBcGetBaseRequest) (_result *ItemBcGetBaseResponse, _err error) {
	return invoke[ItemBcGetBaseResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item_bc/get_base/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"item_id": tea.StringValue(request.ItemId),
			"open_id": tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemBcGetComment(request *ItemBcGetCommentRequest) (_result *ItemBcGetCommentResponse, _err error) {
	return invoke[ItemBcGetCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item_bc/get_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemBcGetLike(request *ItemBcGetLikeRequest) (_result *ItemBcGetLikeResponse, _err error) {
	return invoke[ItemBcGetLikeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item_bc/get_like/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemBcGetPlay(request *ItemBcGetPlayRequest) (_result *ItemBcGetPlayResponse, _err error) {
	return invoke[ItemBcGetPlayResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item_bc/get_play/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemBcGetShare(request *ItemBcGetShareRequest) (_result *ItemBcGetShareResponse, _err error) {
	return invoke[ItemBcGetShareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item_bc/get_share/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemGetBase(request *ItemGetBaseRequest) (_result *ItemGetBaseResponse, _err error) {
	return invoke[ItemGetBaseResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/get_base/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"item_id": tea.StringValue(request.ItemId),
			"open_id": tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemGetComment(request *ItemGetCommentRequest) (_result *ItemGetCommentResponse, _err error) {
	return invoke[ItemGetCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/get_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemGetLike(request *ItemGetLikeRequest) (_result *ItemGetLikeResponse, _err error) {
	return invoke[ItemGetLikeResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/get_like/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemGetPlay(request *ItemGetPlayRequest) (_result *ItemGetPlayResponse, _err error) {
	return invoke[ItemGetPlayResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/get_play/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemGetShare(request *ItemGetShareRequest) (_result *ItemGetShareResponse, _err error) {
	return invoke[ItemGetShareResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/get_share/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"date_type": tea.Int64Value(request.DateType),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemListComment(request *ItemListCommentRequest) (_result *ItemListCommentResponse, _err error) {
	return invoke[ItemListCommentResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/list_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"count":     tea.Int32Value(request.Count),
			"cursor":    tea.Int64Value(request.Cursor),
			"item_id":   tea.StringValue(request.ItemId),
			"open_id":   tea.StringValue(request.OpenId),
			"sort_type": tea.StringValue(request.SortType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemListCommentReply(request *ItemListCommentReplyRequest) (_result *ItemListCommentReplyResponse, _err error) {
	return invoke[ItemListCommentReplyResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/item/list_comment_reply/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"comment_id": tea.StringValue(request.CommentId),
			"count":      tea.Int32Value(request.Count),
			"cursor":     tea.Int64Value(request.Cursor),
			"item_id":    tea.StringValue(request.ItemId),
			"open_id":    tea.StringValue(request.OpenId),
			"sort_type":  tea.StringValue(request.SortType),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ItemReplyComment(request *ItemReplyCommentRequest) (_result *ItemReplyCommentResponse, _err error) {
	return invoke[ItemReplyCommentResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/item/reply_comment/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"device_brand":    tea.StringValue(request.DeviceBrand),
			"device_platform": tea.StringValue(request.DevicePlatform),
			"device_type":     tea.StringValue(request.DeviceType),
			"ip":              tea.StringValue(request.Ip),
			"open_id":         tea.StringValue(request.OpenId),
			"os_version":      tea.StringValue(request.OsVersion),
			"shark_channel":   tea.StringValue(request.SharkChannel),
		},
		Body: map[string]interface{}{
			"comment_id": tea.StringValue(request.CommentId),
			"content":    tea.StringValue(request.Content),
			"item_id":    tea.StringValue(request.ItemId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) JsGetticket(request *JsGetticketRequest) (_result *JsGetticketResponse, _err error) {
	return invoke[JsGetticketResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/js/getticket/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LedgerDetailedQuery(request *LedgerDetailedQueryRequest) (_result *LedgerDetailedQueryResponse, _err error) {
	return invoke[LedgerDetailedQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/ledger/detailed_query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"bill_date":  tea.StringValue(request.BillDate),
			"cursor":     tea.StringValue(request.Cursor),
			"size":       tea.Int64Value(request.Size),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LedgerQuery(request *LedgerQueryRequest) (_result *LedgerQueryResponse, _err error) {
	return invoke[LedgerQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/ledger/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"bill_date":  tea.StringValue(request.BillDate),
			"cursor":     tea.StringValue(request.Cursor),
			"size":       tea.Int64Value(request.Size),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LedgerQueryByOrder(request *LedgerQueryByOrderRequest) (_result *LedgerQueryByOrderResponse, _err error) {
	return invoke[LedgerQueryByOrderResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/settle/ledger/query_by_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"order_ids":  tea.StringSliceValue(request.OrderIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LinkGet(request *LinkGetRequest) (_result *LinkGetResponse, _err error) {
	return invoke[LinkGetResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/entrance/h5/link/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"entrance_cert":        request.EntranceCert,
			"entrance_poi":         request.EntrancePoi,
			"external_merchant_id": tea.StringValue(request.ExternalMerchantId),
			"phone_no":             tea.StringValue(request.PhoneNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LinkmicQuery(request *LinkmicQueryRequest) (_result *LinkmicQueryResponse, _err error) {
	return invoke[LinkmicQueryResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/linkmic/query",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"app_id":  tea.StringValue(request.AppId),
			"room_id": tea.StringValue(request.RoomId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "x-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) ListPlanBySpuid(request *ListPlanBySpuidRequest) (_result *ListPlanBySpuidResponse, _err error) {
	return invoke[ListPlanBySpuidResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/match/v2/poi/list_plan_by_spuid/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"page_no":     tea.Int32Value(request.PageNo),
			"page_size":   tea.Int32Value(request.PageSize),
			"spu_id":      tea.Int64Value(request.SpuId),
			"spu_id_type": tea.IntValue(request.SpuIdType),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) LiveDataAck(request *LiveDataAckRequest) (_result *LiveDataAckResponse, _err error) {
	return invoke[LiveDataAckResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/live_data/ack",
		Host:        tea.String("webcast.bytedance.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"ack_type": tea.Int64Value(request.AckType),
			"app_id":   tea.StringValue(request.AppId),
			"data":     tea.StringValue(request.Data),
			"room_id":  tea.StringValue(request.RoomId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MatchTaskQuery(request *MatchTaskQueryRequest) (_result *MatchTaskQueryResponse, _err error) {
	return invoke[MatchTaskQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/poi/match/task/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"task_id": tea.Int64Value(request.TaskId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MerchantClaimSubmit(request *MerchantClaimSubmitRequest) (_result *MerchantClaimSubmitResponse, _err error) {
	return invoke[MerchantClaimSubmitResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/merchant/claim/submit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id": tea.StringValue(request.AccountId),
			"datas":      request.Datas,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MessageGetUserMessage(request *MessageGetUserMessageRequest) (_result *MessageGetUserMessageResponse, _err error) {
	return invoke[MessageGetUserMessageResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/apps/v1/message/get_user_message/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"end_time":   tea.Int64Value(request.EndTime),
			"page_num":   tea.Int64Value(request.PageNum),
			"page_size":  tea.Int64Value(request.PageSize),
			"start_time": tea.Int64Value(request.StartTime),
			"type":       tea.IntValue(request.Type),
			"username":   tea.StringValue(request.Username),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MessageResources(request *MessageResourcesRequest) (_result *MessageResourcesResponse, _err error) {
	return invoke[MessageResourcesResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/api/im/message/resources/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"conversation_id": tea.StringValue(request.ConversationId),
			"message_id":      tea.StringValue(request.MessageId),
			"open_id":         tea.StringValue(request.OpenId),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MsgGroup(request *MsgGroupRequest) (_result *MsgGroupResponse, _err error) {
	return invoke[MsgGroupResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/im/send/msg/group/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"open_id": tea.StringValue(request.OpenId),
		},
		Body: map[string]interface{}{
			"content": request.Content,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MusicHot(request *MusicHotRequest) (_result *MusicHotResponse, _err error) {
	return invoke[MusicHotResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/music/hot/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MusicOriginal(request *MusicOriginalRequest) (_result *MusicOriginalResponse, _err error) {
	return invoke[MusicOriginalResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/music/original/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) MusicSoar(request *MusicSoarRequest) (_result *MusicSoarResponse, _err error) {
	return invoke[MusicSoarResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/data/extern/billboard/music/soar/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OauthAccessToken(request *OauthAccessTokenRequest) (_result *OauthAccessTokenResponse, _err error) {
	return invoke[OauthAccessTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/access_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeForm,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
			"code":          tea.StringValue(request.Code),
			"grant_type":    tea.StringValue(request.GrantType),
		},
		Encoding: transport.EncodingForm,
	})
}

func (client *Client) OauthBusinessScopes(request *OauthBusinessScopesRequest) (_result *OauthBusinessScopesResponse, _err error) {
	return invoke[OauthBusinessScopesResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/business_scopes/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OauthBusinessToken(request *OauthBusinessTokenRequest) (_result *OauthBusinessTokenResponse, _err error) {
	return invoke[OauthBusinessTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/business_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
			"open_id":       tea.StringValue(request.OpenId),
			"scope":         tea.StringValue(request.Scope),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) OauthClientToken(request *OauthClientTokenRequest) (_result *OauthClientTokenResponse, _err error) {
	return invoke[OauthClientTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/client_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
			"grant_type":    tea.StringValue(request.GrantType),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) OauthRefreshBizToken(request *OauthRefreshBizTokenRequest) (_result *OauthRefreshBizTokenResponse, _err error) {
	return invoke[OauthRefreshBizTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/refresh_biz_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
			"refresh_token": tea.StringValue(request.RefreshToken),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) OauthRefreshToken(request *OauthRefreshTokenRequest) (_result *OauthRefreshTokenResponse, _err error) {
	return invoke[OauthRefreshTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/refresh_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeForm,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"grant_type":    tea.StringValue(request.GrantType),
			"refresh_token": tea.StringValue(request.RefreshToken),
		},
		Encoding: transport.EncodingForm,
	})
}

func (client *Client) OauthRenewRefreshToken(request *OauthRenewRefreshTokenRequest) (_result *OauthRenewRefreshTokenResponse, _err error) {
	return invoke[OauthRenewRefreshTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/renew_refresh_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeForm,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"refresh_token": tea.StringValue(request.RefreshToken),
		},
		Encoding: transport.EncodingForm,
	})
}

func (client *Client) OauthStableClientToken(request *OauthStableClientTokenRequest) (_result *OauthStableClientTokenResponse, _err error) {
	return invoke[OauthStableClientTokenResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/stable_client_token/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"client_key":    tea.StringValue(request.ClientKey),
			"client_secret": tea.StringValue(request.ClientSecret),
			"grant_type":    tea.StringValue(request.GrantType),
		},
		Encoding: transport.EncodingJSON,
	})
}

func (client *Client) OauthUserinfo(request *OauthUserinfoRequest) (_result *OauthUserinfoResponse, _err error) {
	return invoke[OauthUserinfoResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/oauth/userinfo/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_token": tea.StringValue(request.AccessToken),
			"open_id":      tea.StringValue(request.OpenId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access_token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OnlineGet(request *OnlineGetRequest) (_result *OnlineGetResponse, _err error) {
	return invoke[OnlineGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/goods/product/online/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"account_id":  tea.StringValue(request.AccountId),
			"out_ids":     tea.StringSliceValue(request.OutIds),
			"product_ids": tea.StringSliceValue(request.ProductIds),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OnlineList(request *OnlineListRequest) (_result *OnlineListResponse, _err error) {
	return invoke[OnlineListResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/x_goods/dish/online/list/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"Base":       request.Base,
			"account_id": tea.StringValue(request.AccountId),
			"dish_type":  tea.IntValue(request.DishType),
			"dy_poi_id":  tea.Int64Value(request.DyPoiId),
			"page_no":    tea.Int32Value(request.PageNo),
			"page_size":  tea.Int32Value(request.PageSize),
			"status":     tea.IntValueSlice(request.Status),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OpenGetticket(request *OpenGetticketRequest) (_result *OpenGetticketResponse, _err error) {
	return invoke[OpenGetticketResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/open/getticket/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OpenItemIdToEncryptId(request *OpenItemIdToEncryptIdRequest) (_result *OpenItemIdToEncryptIdResponse, _err error) {
	return invoke[OpenItemIdToEncryptIdResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/v1/convert_video_id/open_item_id_to_encrypt_id/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"access_key": tea.StringValue(request.AccessKey),
			"video_ids":  tea.StringSliceValue(request.VideoIds),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderApplyRefund(request *OrderApplyRefundRequest) (_result *OrderApplyRefundResponse, _err error) {
	return invoke[OrderApplyRefundResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/after_sale/order/apply_refund/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":          tea.StringValue(request.OrderId),
			"out_after_sale_id": tea.StringValue(request.OutAfterSaleId),
			"refund_reason":     request.RefundReason,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderAudit(request *OrderAuditRequest) (_result *OrderAuditResponse, _err error) {
	return invoke[OrderAuditResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/akte/booking/order/audit/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":         tea.StringValue(request.AccountId),
			"ext_order_id":       tea.StringValue(request.ExtOrderId),
			"order_id":           tea.StringValue(request.OrderId),
			"receipt_status":     tea.IntValue(request.ReceiptStatus),
			"refuse_reason":      tea.StringValue(request.RefuseReason),
			"refuse_reason_code": tea.IntValue(request.RefuseReasonCode),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderConfirm(request *OrderConfirmRequest) (_result *OrderConfirmResponse, _err error) {
	return invoke[OrderConfirmResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/trip/trade/travelagency/order/confirm/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"confirm_info":    request.ConfirmInfo,
			"order_id":        tea.StringValue(request.OrderId),
			"source_order_id": tea.StringValue(request.SourceOrderId),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderCreate(request *OrderCreateRequest) (_result *OrderCreateResponse, _err error) {
	return invoke[OrderCreateResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/partner/order/create/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":          tea.StringValue(request.AccountId),
			"charge_type":         tea.IntValue(request.ChargeType),
			"commission_ratio":    tea.StringValue(request.CommissionRatio),
			"cooperation_content": tea.IntValue(request.CooperationContent),
			"end_time":            tea.Int64Value(request.EndTime),
			"goods_channel":       tea.IntValue(request.GoodsChannel),
			"start_time":          tea.Int64Value(request.StartTime),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderCreateOrder(request *OrderCreateOrderRequest) (_result *OrderCreateOrderResponse, _err error) {
	return invoke[OrderCreateOrderResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/order/create_order/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"biz_line":                 tea.Int32Value(request.BizLine),
			"consume_shop_info":        request.ConsumeShopInfo,
			"contact_name":             tea.StringValue(request.ContactName),
			"cp_book_info":             request.CpBookInfo,
			"cp_extra":                 tea.StringValue(request.CpExtra),
			"delivery_info":            request.DeliveryInfo,
			"discount_amount":          tea.Int64Value(request.DiscountAmount),
			"extra":                    tea.StringValue(request.Extra),
			"fee_list":                 request.FeeList,
			"goods_list":               request.GoodsList,
			"limit_pay_way_list":       tea.Int64ValueSlice(request.LimitPayWayList),
			"open_id":                  tea.StringValue(request.OpenId),
			"order_entry_schema":       request.OrderEntrySchema,
			"out_order_no":             tea.StringValue(request.OutOrderNo),
			"pay_expire_seconds":       tea.Int64Value(request.PayExpireSeconds),
			"pay_notify_url":           tea.StringValue(request.PayNotifyUrl),
			"phone_num":                tea.StringValue(request.PhoneNum),
			"price_calculation_detail": request.PriceCalculationDetail,
			"reserve_type":             tea.Int32Value(request.ReserveType),
			"sku_list":                 request.SkuList,
			"total_amount":             tea.Int64Value(request.TotalAmount),
			"trade_option":             tea.StringValue(request.TradeOption),
			"user_address":             request.UserAddress,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderGet(request *OrderGetRequest) (_result *OrderGetResponse, _err error) {
	return invoke[OrderGetResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/order/get/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Query: map[string]interface{}{
			"order_id":              tea.StringValue(request.OrderId),
			"without_product_items": tea.StringValue(request.WithoutProductItems),
		},
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderMerchantReject(request *OrderMerchantRejectRequest) (_result *OrderMerchantRejectResponse, _err error) {
	return invoke[OrderMerchantRejectResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/goodlife/v1/after_sale/order/merchant_reject/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":      tea.StringValue(request.OrderId),
			"reject_reason": request.RejectReason,
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderQuery(request *OrderQueryRequest) (_result *OrderQueryResponse, _err error) {
	return invoke[OrderQueryResponse](client, request, &transport.Request{
		Method:      "GET",
		Path:        "/goodlife/v1/partner/order/query/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"account_id":           tea.StringValue(request.AccountId),
			"cooperation_contents": tea.IntValueSlice(request.CooperationContents),
			"end_time":             tea.Int64Value(request.EndTime),
			"goods_channel":        tea.IntValue(request.GoodsChannel),
			"is_asc":               tea.BoolValue(request.IsAsc),
			"page":                 tea.Int32Value(request.Page),
			"size":                 tea.Int32Value(request.Size),
			"start_time":           tea.Int64Value(request.StartTime),
			"status":               tea.IntValueSlice(request.Status),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}

func (client *Client) OrderQueryCps(request *OrderQueryCpsRequest) (_result *OrderQueryCpsResponse, _err error) {
	return invoke[OrderQueryCpsResponse](client, request, &transport.Request{
		Method:      "POST",
		Path:        "/api/apps/trade/v2/order/query_cps/",
		Host:        client.GetHost("open.douyin.com"),
		Headers:     request.Header,
		ContentType: transport.ContentTypeJSON,
		Body: map[string]interface{}{
			"order_id":     tea.StringValue(request.OrderId),
			"out_order_no": tea.StringValue(request.OutOrderNo),
		},
		Encoding:    transport.EncodingJSON,
		TokenHeader: "access-token",
		Token:       request.AccessToken,
	})
}
