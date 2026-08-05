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

type DraftListResponseDataDishsItemImageListItem struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s DraftListResponseDataDishsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemImageListItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemImageListItem) SetUrl(v string) *DraftListResponseDataDishsItemImageListItem {
	s.Url = &v
	return s
}

type DraftListResponseDataDishsItemPoisItem struct {
	PoiId *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s DraftListResponseDataDishsItemPoisItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemPoisItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemPoisItem) SetPoiId(v string) *DraftListResponseDataDishsItemPoisItem {
	s.PoiId = &v
	return s
}

type DraftListResponseDataDishsItemProductSpecAttrsItem struct {
	ItemList  []*DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                           `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s DraftListResponseDataDishsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItem) SetItemList(v []*DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) *DraftListResponseDataDishsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItem) SetGroupName(v string) *DraftListResponseDataDishsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

type DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem struct {
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
}

func (s DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *DraftListResponseDataDishsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

type DraftListResponseDataDishsItemSkusItem struct {
	SkuName             *string                                                   `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuId               *int64                                                    `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	IsSetAutoComplement *bool                                                     `json:"is_set_auto_complement,omitempty" xml:"is_set_auto_complement,omitempty"`
	MaxStock            *int64                                                    `json:"max_stock,omitempty" xml:"max_stock,omitempty"`
	SkuSpecAttrs        []*DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	ResidueStock        *int64                                                    `json:"residue_stock,omitempty" xml:"residue_stock,omitempty"`
	PackFee             *DraftListResponseDataDishsItemSkusItemPackFee            `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	ActualAmount        *int64                                                    `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	IsSetSellOut        *bool                                                     `json:"is_set_sell_out,omitempty" xml:"is_set_sell_out,omitempty"`
	OutSkuId            *string                                                   `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s DraftListResponseDataDishsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemSkusItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemSkusItem) SetSkuName(v string) *DraftListResponseDataDishsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetSkuId(v int64) *DraftListResponseDataDishsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetIsSetAutoComplement(v bool) *DraftListResponseDataDishsItemSkusItem {
	s.IsSetAutoComplement = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetMaxStock(v int64) *DraftListResponseDataDishsItemSkusItem {
	s.MaxStock = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetSkuSpecAttrs(v []*DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem) *DraftListResponseDataDishsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetResidueStock(v int64) *DraftListResponseDataDishsItemSkusItem {
	s.ResidueStock = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetPackFee(v *DraftListResponseDataDishsItemSkusItemPackFee) *DraftListResponseDataDishsItemSkusItem {
	s.PackFee = v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetActualAmount(v int64) *DraftListResponseDataDishsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetIsSetSellOut(v bool) *DraftListResponseDataDishsItemSkusItem {
	s.IsSetSellOut = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItem) SetOutSkuId(v string) *DraftListResponseDataDishsItemSkusItem {
	s.OutSkuId = &v
	return s
}

type DraftListResponseDataDishsItemSkusItemPackFee struct {
	Step        *int32 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int32 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s DraftListResponseDataDishsItemSkusItemPackFee) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemSkusItemPackFee) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemSkusItemPackFee) SetStep(v int32) *DraftListResponseDataDishsItemSkusItemPackFee {
	s.Step = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemPackFee) SetPackFee(v int32) *DraftListResponseDataDishsItemSkusItemPackFee {
	s.PackFee = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemPackFee) SetPackFeeUnit(v int) *DraftListResponseDataDishsItemSkusItemPackFee {
	s.PackFeeUnit = &v
	return s
}

type DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem struct {
	GroupName *string                                                               `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

type DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *DraftListResponseDataDishsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

type DraftListResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s DraftListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DraftListResponseExtra) GoString() string {
	return s.String()
}

func (s *DraftListResponseExtra) SetDescription(v string) *DraftListResponseExtra {
	s.Description = &v
	return s
}

func (s *DraftListResponseExtra) SetErrorCode(v int32) *DraftListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DraftListResponseExtra) SetLogid(v string) *DraftListResponseExtra {
	s.Logid = &v
	return s
}

func (s *DraftListResponseExtra) SetNow(v int64) *DraftListResponseExtra {
	s.Now = &v
	return s
}

func (s *DraftListResponseExtra) SetSubDescription(v string) *DraftListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DraftListResponseExtra) SetSubErrorCode(v int32) *DraftListResponseExtra {
	s.SubErrorCode = &v
	return s
}

type DramaOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DramaOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallRequest) GoString() string {
	return s.String()
}

func (s *DramaOverallRequest) SetHeader(v map[string]*string) *DramaOverallRequest {
	s.Header = v
	return s
}

func (s *DramaOverallRequest) SetAccessToken(v string) *DramaOverallRequest {
	s.AccessToken = &v
	return s
}

type DramaOverallResponse struct {
	Extra *DramaOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *DramaOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DramaOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallResponse) GoString() string {
	return s.String()
}

func (s *DramaOverallResponse) SetExtra(v *DramaOverallResponseExtra) *DramaOverallResponse {
	s.Extra = v
	return s
}

func (s *DramaOverallResponse) SetData(v *DramaOverallResponseData) *DramaOverallResponse {
	s.Data = v
	return s
}

type DramaOverallResponseData struct {
	List          []*DramaOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DramaOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallResponseData) GoString() string {
	return s.String()
}

func (s *DramaOverallResponseData) SetList(v []*DramaOverallResponseDataListItem) *DramaOverallResponseData {
	s.List = v
	return s
}

func (s *DramaOverallResponseData) SetGwErrorCode(v int32) *DramaOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DramaOverallResponseData) SetGwDescription(v string) *DramaOverallResponseData {
	s.GwDescription = &v
	return s
}

type DramaOverallResponseDataListItem struct {
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*DramaOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
}

func (s DramaOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *DramaOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *DramaOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetEffectValue(v float64) *DramaOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetVideoList(v []*DramaOverallResponseDataListItemVideoListItem) *DramaOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *DramaOverallResponseDataListItem) SetRank(v int32) *DramaOverallResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetRankChange(v string) *DramaOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetNickname(v string) *DramaOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetAvatar(v string) *DramaOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *DramaOverallResponseDataListItem) SetFollowerCount(v int64) *DramaOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

type DramaOverallResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DramaOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *DramaOverallResponseDataListItemVideoListItem) SetItemCover(v string) *DramaOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *DramaOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *DramaOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *DramaOverallResponseDataListItemVideoListItem) SetTitle(v string) *DramaOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type DramaOverallResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DramaOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DramaOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *DramaOverallResponseExtra) SetDescription(v string) *DramaOverallResponseExtra {
	s.Description = &v
	return s
}

func (s *DramaOverallResponseExtra) SetSubErrorCode(v int32) *DramaOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DramaOverallResponseExtra) SetSubDescription(v string) *DramaOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DramaOverallResponseExtra) SetLogid(v string) *DramaOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *DramaOverallResponseExtra) SetNow(v int64) *DramaOverallResponseExtra {
	s.Now = &v
	return s
}

func (s *DramaOverallResponseExtra) SetErrorCode(v int32) *DramaOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

type DylkOpenIdGetRequest struct {
	OpenIdList  []*string          `json:"open_id_list,omitempty" xml:"open_id_list,omitempty" require:"true" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DylkOpenIdGetRequest) String() string {
	return tea.Prettify(s)
}

func (s DylkOpenIdGetRequest) GoString() string {
	return s.String()
}

func (s *DylkOpenIdGetRequest) SetOpenIdList(v []*string) *DylkOpenIdGetRequest {
	s.OpenIdList = v
	return s
}

func (s *DylkOpenIdGetRequest) SetAccountId(v string) *DylkOpenIdGetRequest {
	s.AccountId = &v
	return s
}

func (s *DylkOpenIdGetRequest) SetHeader(v map[string]*string) *DylkOpenIdGetRequest {
	s.Header = v
	return s
}

func (s *DylkOpenIdGetRequest) SetAccessToken(v string) *DylkOpenIdGetRequest {
	s.AccessToken = &v
	return s
}

type DylkOpenIdGetResponse struct {
	Data  *DylkOpenIdGetResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *DylkOpenIdGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s DylkOpenIdGetResponse) String() string {
	return tea.Prettify(s)
}

func (s DylkOpenIdGetResponse) GoString() string {
	return s.String()
}

func (s *DylkOpenIdGetResponse) SetData(v *DylkOpenIdGetResponseData) *DylkOpenIdGetResponse {
	s.Data = v
	return s
}

func (s *DylkOpenIdGetResponse) SetExtra(v *DylkOpenIdGetResponseExtra) *DylkOpenIdGetResponse {
	s.Extra = v
	return s
}

type DylkOpenIdGetResponseData struct {
	GwErrorCode   *int32             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	DylkOpenIdMap map[string]*string `json:"dylk_open_id_map,omitempty" xml:"dylk_open_id_map,omitempty" require:"true"`
}

func (s DylkOpenIdGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s DylkOpenIdGetResponseData) GoString() string {
	return s.String()
}

func (s *DylkOpenIdGetResponseData) SetGwErrorCode(v int32) *DylkOpenIdGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DylkOpenIdGetResponseData) SetGwDescription(v string) *DylkOpenIdGetResponseData {
	s.GwDescription = &v
	return s
}

func (s *DylkOpenIdGetResponseData) SetDylkOpenIdMap(v map[string]*string) *DylkOpenIdGetResponseData {
	s.DylkOpenIdMap = v
	return s
}

type DylkOpenIdGetResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s DylkOpenIdGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DylkOpenIdGetResponseExtra) GoString() string {
	return s.String()
}

func (s *DylkOpenIdGetResponseExtra) SetDescription(v string) *DylkOpenIdGetResponseExtra {
	s.Description = &v
	return s
}

func (s *DylkOpenIdGetResponseExtra) SetErrorCode(v int32) *DylkOpenIdGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DylkOpenIdGetResponseExtra) SetLogid(v string) *DylkOpenIdGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *DylkOpenIdGetResponseExtra) SetNow(v int64) *DylkOpenIdGetResponseExtra {
	s.Now = &v
	return s
}

func (s *DylkOpenIdGetResponseExtra) SetSubDescription(v string) *DylkOpenIdGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DylkOpenIdGetResponseExtra) SetSubErrorCode(v int32) *DylkOpenIdGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

type FailDataGetRequest struct {
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Roomid      *string            `json:"roomid,omitempty" xml:"roomid,omitempty" require:"true"`
	Appid       *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	MsgType     *string            `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
}

func (s FailDataGetRequest) String() string {
	return tea.Prettify(s)
}

func (s FailDataGetRequest) GoString() string {
	return s.String()
}

func (s *FailDataGetRequest) SetPageNum(v int32) *FailDataGetRequest {
	s.PageNum = &v
	return s
}

func (s *FailDataGetRequest) SetHeader(v map[string]*string) *FailDataGetRequest {
	s.Header = v
	return s
}

func (s *FailDataGetRequest) SetAccessToken(v string) *FailDataGetRequest {
	s.AccessToken = &v
	return s
}

func (s *FailDataGetRequest) SetPageSize(v int32) *FailDataGetRequest {
	s.PageSize = &v
	return s
}

func (s *FailDataGetRequest) SetRoomid(v string) *FailDataGetRequest {
	s.Roomid = &v
	return s
}

func (s *FailDataGetRequest) SetAppid(v string) *FailDataGetRequest {
	s.Appid = &v
	return s
}

func (s *FailDataGetRequest) SetMsgType(v string) *FailDataGetRequest {
	s.MsgType = &v
	return s
}

type FailDataGetResponse struct {
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	Data   *FailDataGetResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Logid  *string                  `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s FailDataGetResponse) String() string {
	return tea.Prettify(s)
}

func (s FailDataGetResponse) GoString() string {
	return s.String()
}

func (s *FailDataGetResponse) SetErrNo(v int32) *FailDataGetResponse {
	s.ErrNo = &v
	return s
}

func (s *FailDataGetResponse) SetErrMsg(v string) *FailDataGetResponse {
	s.ErrMsg = &v
	return s
}

func (s *FailDataGetResponse) SetData(v *FailDataGetResponseData) *FailDataGetResponse {
	s.Data = v
	return s
}

func (s *FailDataGetResponse) SetLogid(v string) *FailDataGetResponse {
	s.Logid = &v
	return s
}

type FailDataGetResponseData struct {
	PageNum    *int32                                 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	TotalCount *int32                                 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	DataList   []*FailDataGetResponseDataDataListItem `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
}

func (s FailDataGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s FailDataGetResponseData) GoString() string {
	return s.String()
}

func (s *FailDataGetResponseData) SetPageNum(v int32) *FailDataGetResponseData {
	s.PageNum = &v
	return s
}

func (s *FailDataGetResponseData) SetTotalCount(v int32) *FailDataGetResponseData {
	s.TotalCount = &v
	return s
}

func (s *FailDataGetResponseData) SetDataList(v []*FailDataGetResponseDataDataListItem) *FailDataGetResponseData {
	s.DataList = v
	return s
}

type FailDataGetResponseDataDataListItem struct {
	Payload *string `json:"payload,omitempty" xml:"payload,omitempty" require:"true"`
	RoomId  *string `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	MsgType *string `json:"msg_type,omitempty" xml:"msg_type,omitempty" require:"true"`
}

func (s FailDataGetResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FailDataGetResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *FailDataGetResponseDataDataListItem) SetPayload(v string) *FailDataGetResponseDataDataListItem {
	s.Payload = &v
	return s
}

func (s *FailDataGetResponseDataDataListItem) SetRoomId(v string) *FailDataGetResponseDataDataListItem {
	s.RoomId = &v
	return s
}

func (s *FailDataGetResponseDataDataListItem) SetMsgType(v string) *FailDataGetResponseDataDataListItem {
	s.MsgType = &v
	return s
}

type FansCheckRequest struct {
	FollowerOpenId *string            `json:"follower_open_id,omitempty" xml:"follower_open_id,omitempty" require:"true"`
	OpenId         *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FansCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s FansCheckRequest) GoString() string {
	return s.String()
}

func (s *FansCheckRequest) SetFollowerOpenId(v string) *FansCheckRequest {
	s.FollowerOpenId = &v
	return s
}

func (s *FansCheckRequest) SetOpenId(v string) *FansCheckRequest {
	s.OpenId = &v
	return s
}

func (s *FansCheckRequest) SetHeader(v map[string]*string) *FansCheckRequest {
	s.Header = v
	return s
}

func (s *FansCheckRequest) SetAccessToken(v string) *FansCheckRequest {
	s.AccessToken = &v
	return s
}

type FansCheckResponse struct {
	Data  *FansCheckResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *FansCheckResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FansCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s FansCheckResponse) GoString() string {
	return s.String()
}

func (s *FansCheckResponse) SetData(v *FansCheckResponseData) *FansCheckResponse {
	s.Data = v
	return s
}

func (s *FansCheckResponse) SetExtra(v *FansCheckResponseExtra) *FansCheckResponse {
	s.Extra = v
	return s
}

type FansCheckResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	FollowTime    *int64  `json:"follow_time,omitempty" xml:"follow_time,omitempty"`
	IsFollower    *bool   `json:"is_follower,omitempty" xml:"is_follower,omitempty" require:"true"`
}

func (s FansCheckResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansCheckResponseData) GoString() string {
	return s.String()
}

func (s *FansCheckResponseData) SetGwErrorCode(v int32) *FansCheckResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FansCheckResponseData) SetGwDescription(v string) *FansCheckResponseData {
	s.GwDescription = &v
	return s
}

func (s *FansCheckResponseData) SetFollowTime(v int64) *FansCheckResponseData {
	s.FollowTime = &v
	return s
}

func (s *FansCheckResponseData) SetIsFollower(v bool) *FansCheckResponseData {
	s.IsFollower = &v
	return s
}

type FansCheckResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FansCheckResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansCheckResponseExtra) GoString() string {
	return s.String()
}

func (s *FansCheckResponseExtra) SetErrorCode(v int32) *FansCheckResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FansCheckResponseExtra) SetLogid(v string) *FansCheckResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansCheckResponseExtra) SetNow(v int64) *FansCheckResponseExtra {
	s.Now = &v
	return s
}

func (s *FansCheckResponseExtra) SetSubDescription(v string) *FansCheckResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansCheckResponseExtra) SetSubErrorCode(v int32) *FansCheckResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FansCheckResponseExtra) SetDescription(v string) *FansCheckResponseExtra {
	s.Description = &v
	return s
}

type FansClubGetInfoRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AnchorOpenid *string            `json:"anchor_openid,omitempty" xml:"anchor_openid,omitempty" require:"true"`
	UserOpenids  []*string          `json:"user_openids,omitempty" xml:"user_openids,omitempty" require:"true" type:"Repeated"`
	Roomid       *int64             `json:"roomid,omitempty" xml:"roomid,omitempty" require:"true"`
}

func (s FansClubGetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s FansClubGetInfoRequest) GoString() string {
	return s.String()
}

func (s *FansClubGetInfoRequest) SetHeader(v map[string]*string) *FansClubGetInfoRequest {
	s.Header = v
	return s
}

func (s *FansClubGetInfoRequest) SetAccessToken(v string) *FansClubGetInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *FansClubGetInfoRequest) SetAnchorOpenid(v string) *FansClubGetInfoRequest {
	s.AnchorOpenid = &v
	return s
}

func (s *FansClubGetInfoRequest) SetUserOpenids(v []*string) *FansClubGetInfoRequest {
	s.UserOpenids = v
	return s
}

func (s *FansClubGetInfoRequest) SetRoomid(v int64) *FansClubGetInfoRequest {
	s.Roomid = &v
	return s
}

type FansClubGetInfoResponse struct {
	Data   *FansClubGetInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Logid  *string                      `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s FansClubGetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s FansClubGetInfoResponse) GoString() string {
	return s.String()
}

func (s *FansClubGetInfoResponse) SetData(v *FansClubGetInfoResponseData) *FansClubGetInfoResponse {
	s.Data = v
	return s
}

func (s *FansClubGetInfoResponse) SetLogid(v string) *FansClubGetInfoResponse {
	s.Logid = &v
	return s
}

func (s *FansClubGetInfoResponse) SetErrNo(v int32) *FansClubGetInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *FansClubGetInfoResponse) SetErrMsg(v string) *FansClubGetInfoResponse {
	s.ErrMsg = &v
	return s
}

type FansClubGetInfoResponseData struct {
	FansClubInfo map[string]*FansClubGetInfoResponseDataFansClubInfoValue `json:"fans_club_Info,omitempty" xml:"fans_club_Info,omitempty"`
}

func (s FansClubGetInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansClubGetInfoResponseData) GoString() string {
	return s.String()
}

func (s *FansClubGetInfoResponseData) SetFansClubInfo(v map[string]*FansClubGetInfoResponseDataFansClubInfoValue) *FansClubGetInfoResponseData {
	s.FansClubInfo = v
	return s
}

type FansClubGetInfoResponseDataFansClubInfoValue struct {
	LevelLayer      *int64 `json:"level_layer,omitempty" xml:"level_layer,omitempty"`
	ParticipateTime *int64 `json:"participate_time,omitempty" xml:"participate_time,omitempty"`
}

func (s FansClubGetInfoResponseDataFansClubInfoValue) String() string {
	return tea.Prettify(s)
}

func (s FansClubGetInfoResponseDataFansClubInfoValue) GoString() string {
	return s.String()
}

func (s *FansClubGetInfoResponseDataFansClubInfoValue) SetLevelLayer(v int64) *FansClubGetInfoResponseDataFansClubInfoValue {
	s.LevelLayer = &v
	return s
}

func (s *FansClubGetInfoResponseDataFansClubInfoValue) SetParticipateTime(v int64) *FansClubGetInfoResponseDataFansClubInfoValue {
	s.ParticipateTime = &v
	return s
}

type FansCommentRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FansCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s FansCommentRequest) GoString() string {
	return s.String()
}

func (s *FansCommentRequest) SetOpenId(v string) *FansCommentRequest {
	s.OpenId = &v
	return s
}

func (s *FansCommentRequest) SetHeader(v map[string]*string) *FansCommentRequest {
	s.Header = v
	return s
}

func (s *FansCommentRequest) SetAccessToken(v string) *FansCommentRequest {
	s.AccessToken = &v
	return s
}

type FansCommentResponse struct {
	Extra *FansCommentResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *FansCommentResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FansCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s FansCommentResponse) GoString() string {
	return s.String()
}

func (s *FansCommentResponse) SetExtra(v *FansCommentResponseExtra) *FansCommentResponse {
	s.Extra = v
	return s
}

func (s *FansCommentResponse) SetData(v *FansCommentResponseData) *FansCommentResponse {
	s.Data = v
	return s
}

type FansCommentResponseData struct {
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*FansCommentResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FansCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansCommentResponseData) GoString() string {
	return s.String()
}

func (s *FansCommentResponseData) SetGwDescription(v string) *FansCommentResponseData {
	s.GwDescription = &v
	return s
}

func (s *FansCommentResponseData) SetList(v []*FansCommentResponseDataListItem) *FansCommentResponseData {
	s.List = v
	return s
}

func (s *FansCommentResponseData) SetGwErrorCode(v int32) *FansCommentResponseData {
	s.GwErrorCode = &v
	return s
}

type FansCommentResponseDataListItem struct {
	HotValue *int64  `json:"hot_value,omitempty" xml:"hot_value,omitempty" require:"true"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
}

func (s FansCommentResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FansCommentResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FansCommentResponseDataListItem) SetHotValue(v int64) *FansCommentResponseDataListItem {
	s.HotValue = &v
	return s
}

func (s *FansCommentResponseDataListItem) SetKeyword(v string) *FansCommentResponseDataListItem {
	s.Keyword = &v
	return s
}

type FansCommentResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s FansCommentResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansCommentResponseExtra) GoString() string {
	return s.String()
}

func (s *FansCommentResponseExtra) SetDescription(v string) *FansCommentResponseExtra {
	s.Description = &v
	return s
}

func (s *FansCommentResponseExtra) SetErrorCode(v int32) *FansCommentResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FansCommentResponseExtra) SetLogid(v string) *FansCommentResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansCommentResponseExtra) SetNow(v int64) *FansCommentResponseExtra {
	s.Now = &v
	return s
}

func (s *FansCommentResponseExtra) SetSubDescription(v string) *FansCommentResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansCommentResponseExtra) SetSubErrorCode(v int32) *FansCommentResponseExtra {
	s.SubErrorCode = &v
	return s
}

type FansCreateRequest struct {
	GroupType       *int64             `json:"group_type,omitempty" xml:"group_type,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	RelationType    *int64             `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	AvatarUri       *string            `json:"avatar_uri,omitempty" xml:"avatar_uri,omitempty" require:"true"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId          *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	LiveAutoSync    *int64             `json:"live_auto_sync,omitempty" xml:"live_auto_sync,omitempty"`
	ActiveFans      *int64             `json:"active_fans,omitempty" xml:"active_fans,omitempty"`
	AllowInvite     *int64             `json:"allow_invite,omitempty" xml:"allow_invite,omitempty"`
	FansLimit       *int64             `json:"fans_limit,omitempty" xml:"fans_limit,omitempty"`
	ItemAutoSync    *int64             `json:"item_auto_sync,omitempty" xml:"item_auto_sync,omitempty"`
	ShowAtProfile   *int64             `json:"show_at_profile,omitempty" xml:"show_at_profile,omitempty"`
	GroupName       *string            `json:"group_name,omitempty" xml:"group_name,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	OpenAuditSwitch *int64             `json:"open_audit_switch,omitempty" xml:"open_audit_switch,omitempty"`
}

func (s FansCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s FansCreateRequest) GoString() string {
	return s.String()
}

func (s *FansCreateRequest) SetGroupType(v int64) *FansCreateRequest {
	s.GroupType = &v
	return s
}

func (s *FansCreateRequest) SetDescription(v string) *FansCreateRequest {
	s.Description = &v
	return s
}

func (s *FansCreateRequest) SetRelationType(v int64) *FansCreateRequest {
	s.RelationType = &v
	return s
}

func (s *FansCreateRequest) SetAvatarUri(v string) *FansCreateRequest {
	s.AvatarUri = &v
	return s
}

func (s *FansCreateRequest) SetAccessToken(v string) *FansCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *FansCreateRequest) SetOpenId(v string) *FansCreateRequest {
	s.OpenId = &v
	return s
}

func (s *FansCreateRequest) SetLiveAutoSync(v int64) *FansCreateRequest {
	s.LiveAutoSync = &v
	return s
}

func (s *FansCreateRequest) SetActiveFans(v int64) *FansCreateRequest {
	s.ActiveFans = &v
	return s
}

func (s *FansCreateRequest) SetAllowInvite(v int64) *FansCreateRequest {
	s.AllowInvite = &v
	return s
}

func (s *FansCreateRequest) SetFansLimit(v int64) *FansCreateRequest {
	s.FansLimit = &v
	return s
}

func (s *FansCreateRequest) SetItemAutoSync(v int64) *FansCreateRequest {
	s.ItemAutoSync = &v
	return s
}

func (s *FansCreateRequest) SetShowAtProfile(v int64) *FansCreateRequest {
	s.ShowAtProfile = &v
	return s
}

func (s *FansCreateRequest) SetGroupName(v string) *FansCreateRequest {
	s.GroupName = &v
	return s
}

func (s *FansCreateRequest) SetHeader(v map[string]*string) *FansCreateRequest {
	s.Header = v
	return s
}

func (s *FansCreateRequest) SetOpenAuditSwitch(v int64) *FansCreateRequest {
	s.OpenAuditSwitch = &v
	return s
}

type FansCreateResponse struct {
	GroupId *string                  `json:"group_id,omitempty" xml:"group_id,omitempty"`
	Extra   *FansCreateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data    *FansCreateResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FansCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s FansCreateResponse) GoString() string {
	return s.String()
}

func (s *FansCreateResponse) SetGroupId(v string) *FansCreateResponse {
	s.GroupId = &v
	return s
}

func (s *FansCreateResponse) SetExtra(v *FansCreateResponseExtra) *FansCreateResponse {
	s.Extra = v
	return s
}

func (s *FansCreateResponse) SetData(v *FansCreateResponseData) *FansCreateResponse {
	s.Data = v
	return s
}

type FansCreateResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FansCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansCreateResponseData) GoString() string {
	return s.String()
}

func (s *FansCreateResponseData) SetGwDescription(v string) *FansCreateResponseData {
	s.GwDescription = &v
	return s
}

func (s *FansCreateResponseData) SetGwErrorCode(v int32) *FansCreateResponseData {
	s.GwErrorCode = &v
	return s
}

type FansCreateResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s FansCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *FansCreateResponseExtra) SetSubDescription(v string) *FansCreateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansCreateResponseExtra) SetLogid(v string) *FansCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansCreateResponseExtra) SetNow(v int64) *FansCreateResponseExtra {
	s.Now = &v
	return s
}

func (s *FansCreateResponseExtra) SetErrorCode(v int32) *FansCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FansCreateResponseExtra) SetDescription(v string) *FansCreateResponseExtra {
	s.Description = &v
	return s
}

func (s *FansCreateResponseExtra) SetSubErrorCode(v int32) *FansCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type FansFavouriteRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FansFavouriteRequest) String() string {
	return tea.Prettify(s)
}

func (s FansFavouriteRequest) GoString() string {
	return s.String()
}

func (s *FansFavouriteRequest) SetOpenId(v string) *FansFavouriteRequest {
	s.OpenId = &v
	return s
}

func (s *FansFavouriteRequest) SetHeader(v map[string]*string) *FansFavouriteRequest {
	s.Header = v
	return s
}

func (s *FansFavouriteRequest) SetAccessToken(v string) *FansFavouriteRequest {
	s.AccessToken = &v
	return s
}

type FansFavouriteResponse struct {
	Extra *FansFavouriteResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *FansFavouriteResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FansFavouriteResponse) String() string {
	return tea.Prettify(s)
}

func (s FansFavouriteResponse) GoString() string {
	return s.String()
}

func (s *FansFavouriteResponse) SetExtra(v *FansFavouriteResponseExtra) *FansFavouriteResponse {
	s.Extra = v
	return s
}

func (s *FansFavouriteResponse) SetData(v *FansFavouriteResponseData) *FansFavouriteResponse {
	s.Data = v
	return s
}

type FansFavouriteResponseData struct {
	List          []*FansFavouriteResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FansFavouriteResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansFavouriteResponseData) GoString() string {
	return s.String()
}

func (s *FansFavouriteResponseData) SetList(v []*FansFavouriteResponseDataListItem) *FansFavouriteResponseData {
	s.List = v
	return s
}

func (s *FansFavouriteResponseData) SetGwErrorCode(v int32) *FansFavouriteResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FansFavouriteResponseData) SetGwDescription(v string) *FansFavouriteResponseData {
	s.GwDescription = &v
	return s
}

type FansFavouriteResponseDataListItem struct {
	Rank     *int32  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	HotValue *int64  `json:"hot_value,omitempty" xml:"hot_value,omitempty" require:"true"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
}

func (s FansFavouriteResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FansFavouriteResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FansFavouriteResponseDataListItem) SetRank(v int32) *FansFavouriteResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *FansFavouriteResponseDataListItem) SetHotValue(v int64) *FansFavouriteResponseDataListItem {
	s.HotValue = &v
	return s
}

func (s *FansFavouriteResponseDataListItem) SetKeyword(v string) *FansFavouriteResponseDataListItem {
	s.Keyword = &v
	return s
}

type FansFavouriteResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s FansFavouriteResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansFavouriteResponseExtra) GoString() string {
	return s.String()
}

func (s *FansFavouriteResponseExtra) SetSubDescription(v string) *FansFavouriteResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansFavouriteResponseExtra) SetSubErrorCode(v int32) *FansFavouriteResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FansFavouriteResponseExtra) SetDescription(v string) *FansFavouriteResponseExtra {
	s.Description = &v
	return s
}

func (s *FansFavouriteResponseExtra) SetErrorCode(v int32) *FansFavouriteResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FansFavouriteResponseExtra) SetLogid(v string) *FansFavouriteResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansFavouriteResponseExtra) SetNow(v int64) *FansFavouriteResponseExtra {
	s.Now = &v
	return s
}

type FansListRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FansListRequest) String() string {
	return tea.Prettify(s)
}

func (s FansListRequest) GoString() string {
	return s.String()
}

func (s *FansListRequest) SetOpenId(v string) *FansListRequest {
	s.OpenId = &v
	return s
}

func (s *FansListRequest) SetHeader(v map[string]*string) *FansListRequest {
	s.Header = v
	return s
}

func (s *FansListRequest) SetAccessToken(v string) *FansListRequest {
	s.AccessToken = &v
	return s
}

type FansListResponse struct {
	Extra     *FansListResponseExtra           `json:"extra,omitempty" xml:"extra,omitempty"`
	GroupList []*FansListResponseGroupListItem `json:"group_list,omitempty" xml:"group_list,omitempty" require:"true" type:"Repeated"`
	Data      *FansListResponseData            `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FansListResponse) String() string {
	return tea.Prettify(s)
}

func (s FansListResponse) GoString() string {
	return s.String()
}

func (s *FansListResponse) SetExtra(v *FansListResponseExtra) *FansListResponse {
	s.Extra = v
	return s
}

func (s *FansListResponse) SetGroupList(v []*FansListResponseGroupListItem) *FansListResponse {
	s.GroupList = v
	return s
}

func (s *FansListResponse) SetData(v *FansListResponseData) *FansListResponse {
	s.Data = v
	return s
}

type FansListResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FansListResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansListResponseData) GoString() string {
	return s.String()
}

func (s *FansListResponseData) SetGwErrorCode(v int32) *FansListResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FansListResponseData) SetGwDescription(v string) *FansListResponseData {
	s.GwDescription = &v
	return s
}

type FansListResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FansListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansListResponseExtra) GoString() string {
	return s.String()
}

func (s *FansListResponseExtra) SetSubErrorCode(v int32) *FansListResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FansListResponseExtra) SetSubDescription(v string) *FansListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansListResponseExtra) SetLogid(v string) *FansListResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansListResponseExtra) SetNow(v int64) *FansListResponseExtra {
	s.Now = &v
	return s
}

func (s *FansListResponseExtra) SetErrorCode(v int32) *FansListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FansListResponseExtra) SetDescription(v string) *FansListResponseExtra {
	s.Description = &v
	return s
}

type FansListResponseGroupListItem struct {
	MaxNum      *int64    `json:"max_num,omitempty" xml:"max_num,omitempty"`
	GroupName   *string   `json:"group_name,omitempty" xml:"group_name,omitempty" require:"true"`
	GroupType   *int32    `json:"group_type,omitempty" xml:"group_type,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	GroupId     *string   `json:"group_id,omitempty" xml:"group_id,omitempty" require:"true"`
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
	Tags        []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	EntryLimit  []*string `json:"entry_limit,omitempty" xml:"entry_limit,omitempty" type:"Repeated"`
	AvatarUri   *string   `json:"avatar_uri,omitempty" xml:"avatar_uri,omitempty"`
	ExistNum    *int64    `json:"exist_num,omitempty" xml:"exist_num,omitempty"`
}

func (s FansListResponseGroupListItem) String() string {
	return tea.Prettify(s)
}

func (s FansListResponseGroupListItem) GoString() string {
	return s.String()
}

func (s *FansListResponseGroupListItem) SetMaxNum(v int64) *FansListResponseGroupListItem {
	s.MaxNum = &v
	return s
}

func (s *FansListResponseGroupListItem) SetGroupName(v string) *FansListResponseGroupListItem {
	s.GroupName = &v
	return s
}

func (s *FansListResponseGroupListItem) SetGroupType(v int32) *FansListResponseGroupListItem {
	s.GroupType = &v
	return s
}

func (s *FansListResponseGroupListItem) SetDescription(v string) *FansListResponseGroupListItem {
	s.Description = &v
	return s
}

func (s *FansListResponseGroupListItem) SetGroupId(v string) *FansListResponseGroupListItem {
	s.GroupId = &v
	return s
}

func (s *FansListResponseGroupListItem) SetStatus(v string) *FansListResponseGroupListItem {
	s.Status = &v
	return s
}

func (s *FansListResponseGroupListItem) SetTags(v []*string) *FansListResponseGroupListItem {
	s.Tags = v
	return s
}

func (s *FansListResponseGroupListItem) SetEntryLimit(v []*string) *FansListResponseGroupListItem {
	s.EntryLimit = v
	return s
}

func (s *FansListResponseGroupListItem) SetAvatarUri(v string) *FansListResponseGroupListItem {
	s.AvatarUri = &v
	return s
}

func (s *FansListResponseGroupListItem) SetExistNum(v int64) *FansListResponseGroupListItem {
	s.ExistNum = &v
	return s
}

type FansSourceRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s FansSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s FansSourceRequest) GoString() string {
	return s.String()
}

func (s *FansSourceRequest) SetAccessToken(v string) *FansSourceRequest {
	s.AccessToken = &v
	return s
}

func (s *FansSourceRequest) SetOpenId(v string) *FansSourceRequest {
	s.OpenId = &v
	return s
}

func (s *FansSourceRequest) SetHeader(v map[string]*string) *FansSourceRequest {
	s.Header = v
	return s
}

type FansSourceResponse struct {
	Data  *FansSourceResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *FansSourceResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FansSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s FansSourceResponse) GoString() string {
	return s.String()
}

func (s *FansSourceResponse) SetData(v *FansSourceResponseData) *FansSourceResponse {
	s.Data = v
	return s
}

func (s *FansSourceResponse) SetExtra(v *FansSourceResponseExtra) *FansSourceResponse {
	s.Extra = v
	return s
}

type FansSourceResponseData struct {
	GwDescription *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*FansSourceResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FansSourceResponseData) String() string {
	return tea.Prettify(s)
}

func (s FansSourceResponseData) GoString() string {
	return s.String()
}

func (s *FansSourceResponseData) SetGwDescription(v string) *FansSourceResponseData {
	s.GwDescription = &v
	return s
}

func (s *FansSourceResponseData) SetList(v []*FansSourceResponseDataListItem) *FansSourceResponseData {
	s.List = v
	return s
}

func (s *FansSourceResponseData) SetGwErrorCode(v int32) *FansSourceResponseData {
	s.GwErrorCode = &v
	return s
}

type FansSourceResponseDataListItem struct {
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty" require:"true"`
	Source  *string `json:"source,omitempty" xml:"source,omitempty" require:"true"`
}

func (s FansSourceResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FansSourceResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FansSourceResponseDataListItem) SetPercent(v string) *FansSourceResponseDataListItem {
	s.Percent = &v
	return s
}

func (s *FansSourceResponseDataListItem) SetSource(v string) *FansSourceResponseDataListItem {
	s.Source = &v
	return s
}

type FansSourceResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FansSourceResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FansSourceResponseExtra) GoString() string {
	return s.String()
}

func (s *FansSourceResponseExtra) SetLogid(v string) *FansSourceResponseExtra {
	s.Logid = &v
	return s
}

func (s *FansSourceResponseExtra) SetNow(v int64) *FansSourceResponseExtra {
	s.Now = &v
	return s
}

func (s *FansSourceResponseExtra) SetSubDescription(v string) *FansSourceResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FansSourceResponseExtra) SetSubErrorCode(v int32) *FansSourceResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FansSourceResponseExtra) SetDescription(v string) *FansSourceResponseExtra {
	s.Description = &v
	return s
}

func (s *FansSourceResponseExtra) SetErrorCode(v int32) *FansSourceResponseExtra {
	s.ErrorCode = &v
	return s
}

type FeedbackBatchReplyRequest struct {
	Appid       *string                                   `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	ReplyList   []*FeedbackBatchReplyRequestReplyListItem `json:"reply_list,omitempty" xml:"reply_list,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string                        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FeedbackBatchReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackBatchReplyRequest) GoString() string {
	return s.String()
}

func (s *FeedbackBatchReplyRequest) SetAppid(v string) *FeedbackBatchReplyRequest {
	s.Appid = &v
	return s
}

func (s *FeedbackBatchReplyRequest) SetReplyList(v []*FeedbackBatchReplyRequestReplyListItem) *FeedbackBatchReplyRequest {
	s.ReplyList = v
	return s
}

func (s *FeedbackBatchReplyRequest) SetHeader(v map[string]*string) *FeedbackBatchReplyRequest {
	s.Header = v
	return s
}

func (s *FeedbackBatchReplyRequest) SetAccessToken(v string) *FeedbackBatchReplyRequest {
	s.AccessToken = &v
	return s
}

type FeedbackBatchReplyRequestReplyListItem struct {
	FeedbackId *int64  `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s FeedbackBatchReplyRequestReplyListItem) String() string {
	return tea.Prettify(s)
}

func (s FeedbackBatchReplyRequestReplyListItem) GoString() string {
	return s.String()
}

func (s *FeedbackBatchReplyRequestReplyListItem) SetFeedbackId(v int64) *FeedbackBatchReplyRequestReplyListItem {
	s.FeedbackId = &v
	return s
}

func (s *FeedbackBatchReplyRequestReplyListItem) SetContent(v string) *FeedbackBatchReplyRequestReplyListItem {
	s.Content = &v
	return s
}

type FeedbackBatchReplyResponse struct {
	Data   *FeedbackBatchReplyResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty"`
}

func (s FeedbackBatchReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackBatchReplyResponse) GoString() string {
	return s.String()
}

func (s *FeedbackBatchReplyResponse) SetData(v *FeedbackBatchReplyResponseData) *FeedbackBatchReplyResponse {
	s.Data = v
	return s
}

func (s *FeedbackBatchReplyResponse) SetErrNo(v int32) *FeedbackBatchReplyResponse {
	s.ErrNo = &v
	return s
}

func (s *FeedbackBatchReplyResponse) SetErrMsg(v string) *FeedbackBatchReplyResponse {
	s.ErrMsg = &v
	return s
}

func (s *FeedbackBatchReplyResponse) SetLogId(v string) *FeedbackBatchReplyResponse {
	s.LogId = &v
	return s
}

type FeedbackBatchReplyResponseData struct {
	FailReplyList []*FeedbackBatchReplyResponseDataFailReplyListItem `json:"fail_reply_list,omitempty" xml:"fail_reply_list,omitempty" type:"Repeated"`
}

func (s FeedbackBatchReplyResponseData) String() string {
	return tea.Prettify(s)
}

func (s FeedbackBatchReplyResponseData) GoString() string {
	return s.String()
}

func (s *FeedbackBatchReplyResponseData) SetFailReplyList(v []*FeedbackBatchReplyResponseDataFailReplyListItem) *FeedbackBatchReplyResponseData {
	s.FailReplyList = v
	return s
}

type FeedbackBatchReplyResponseDataFailReplyListItem struct {
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	FeedbackId *int64  `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
}

func (s FeedbackBatchReplyResponseDataFailReplyListItem) String() string {
	return tea.Prettify(s)
}

func (s FeedbackBatchReplyResponseDataFailReplyListItem) GoString() string {
	return s.String()
}

func (s *FeedbackBatchReplyResponseDataFailReplyListItem) SetFailReason(v string) *FeedbackBatchReplyResponseDataFailReplyListItem {
	s.FailReason = &v
	return s
}

func (s *FeedbackBatchReplyResponseDataFailReplyListItem) SetFeedbackId(v int64) *FeedbackBatchReplyResponseDataFailReplyListItem {
	s.FeedbackId = &v
	return s
}

type FeedbackGetFeedbackListRequest struct {
	Appid        *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	FeedbackId   *int64             `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	CurrentPage  *int32             `json:"current_page,omitempty" xml:"current_page,omitempty" require:"true"`
	FeedbackType *int               `json:"feedback_type,omitempty" xml:"feedback_type,omitempty"`
	StartTime    *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime      *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	DealtStatus  []*int             `json:"dealt_status,omitempty" xml:"dealt_status,omitempty" type:"Repeated"`
	PageSize     *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s FeedbackGetFeedbackListRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetFeedbackListRequest) GoString() string {
	return s.String()
}

func (s *FeedbackGetFeedbackListRequest) SetAppid(v string) *FeedbackGetFeedbackListRequest {
	s.Appid = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetFeedbackId(v int64) *FeedbackGetFeedbackListRequest {
	s.FeedbackId = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetCurrentPage(v int32) *FeedbackGetFeedbackListRequest {
	s.CurrentPage = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetFeedbackType(v int) *FeedbackGetFeedbackListRequest {
	s.FeedbackType = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetStartTime(v int64) *FeedbackGetFeedbackListRequest {
	s.StartTime = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetEndTime(v int64) *FeedbackGetFeedbackListRequest {
	s.EndTime = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetDealtStatus(v []*int) *FeedbackGetFeedbackListRequest {
	s.DealtStatus = v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetPageSize(v int32) *FeedbackGetFeedbackListRequest {
	s.PageSize = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetHeader(v map[string]*string) *FeedbackGetFeedbackListRequest {
	s.Header = v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetAccessToken(v string) *FeedbackGetFeedbackListRequest {
	s.AccessToken = &v
	return s
}

func (s *FeedbackGetFeedbackListRequest) SetOpenId(v string) *FeedbackGetFeedbackListRequest {
	s.OpenId = &v
	return s
}

type FeedbackGetFeedbackListResponse struct {
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *FeedbackGetFeedbackListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s FeedbackGetFeedbackListResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetFeedbackListResponse) GoString() string {
	return s.String()
}

func (s *FeedbackGetFeedbackListResponse) SetErrMsg(v string) *FeedbackGetFeedbackListResponse {
	s.ErrMsg = &v
	return s
}

func (s *FeedbackGetFeedbackListResponse) SetLogId(v string) *FeedbackGetFeedbackListResponse {
	s.LogId = &v
	return s
}

func (s *FeedbackGetFeedbackListResponse) SetData(v *FeedbackGetFeedbackListResponseData) *FeedbackGetFeedbackListResponse {
	s.Data = v
	return s
}

func (s *FeedbackGetFeedbackListResponse) SetErrNo(v int32) *FeedbackGetFeedbackListResponse {
	s.ErrNo = &v
	return s
}

type FeedbackGetFeedbackListResponseData struct {
	FeedbackList []*FeedbackGetFeedbackListResponseDataFeedbackListItem `json:"feedback_list,omitempty" xml:"feedback_list,omitempty" type:"Repeated"`
	Total        *int32                                                 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s FeedbackGetFeedbackListResponseData) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetFeedbackListResponseData) GoString() string {
	return s.String()
}

func (s *FeedbackGetFeedbackListResponseData) SetFeedbackList(v []*FeedbackGetFeedbackListResponseDataFeedbackListItem) *FeedbackGetFeedbackListResponseData {
	s.FeedbackList = v
	return s
}

func (s *FeedbackGetFeedbackListResponseData) SetTotal(v int32) *FeedbackGetFeedbackListResponseData {
	s.Total = &v
	return s
}

type FeedbackGetFeedbackListResponseDataFeedbackListItem struct {
	MicroappVersion *string                                                               `json:"microapp_version,omitempty" xml:"microapp_version,omitempty"`
	AppVersion      *string                                                               `json:"app_version,omitempty" xml:"app_version,omitempty"`
	Platform        *string                                                               `json:"platform,omitempty" xml:"platform,omitempty"`
	OrderId         *string                                                               `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Device          *string                                                               `json:"device,omitempty" xml:"device,omitempty"`
	Schema          *string                                                               `json:"schema,omitempty" xml:"schema,omitempty"`
	OsVersion       *string                                                               `json:"os_version,omitempty" xml:"os_version,omitempty"`
	OpenId          *string                                                               `json:"open_id,omitempty" xml:"open_id,omitempty"`
	DdlTime         *int64                                                                `json:"ddl_time,omitempty" xml:"ddl_time,omitempty"`
	ContentList     []*FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem `json:"content_list,omitempty" xml:"content_list,omitempty" type:"Repeated"`
	IsImportant     *int32                                                                `json:"is_important,omitempty" xml:"is_important,omitempty"`
	DealtStatus     *int                                                                  `json:"dealt_status,omitempty" xml:"dealt_status,omitempty"`
	FeedbackType    *int                                                                  `json:"FeedbackType,omitempty" xml:"FeedbackType,omitempty"`
	FeedbackId      *int64                                                                `json:"feedback_id,omitempty" xml:"feedback_id,omitempty"`
	Mobile          *string                                                               `json:"mobile,omitempty" xml:"mobile,omitempty"`
	LibraryVersion  *string                                                               `json:"library_version,omitempty" xml:"library_version,omitempty"`
}

func (s FeedbackGetFeedbackListResponseDataFeedbackListItem) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetFeedbackListResponseDataFeedbackListItem) GoString() string {
	return s.String()
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetMicroappVersion(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.MicroappVersion = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetAppVersion(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.AppVersion = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetPlatform(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.Platform = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetOrderId(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.OrderId = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetDevice(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.Device = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetSchema(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.Schema = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetOsVersion(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.OsVersion = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetOpenId(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.OpenId = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetDdlTime(v int64) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.DdlTime = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetContentList(v []*FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.ContentList = v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetIsImportant(v int32) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.IsImportant = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetDealtStatus(v int) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.DealtStatus = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetFeedbackType(v int) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.FeedbackType = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetFeedbackId(v int64) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.FeedbackId = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetMobile(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.Mobile = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItem) SetLibraryVersion(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItem {
	s.LibraryVersion = &v
	return s
}

type FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem) GoString() string {
	return s.String()
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem) SetContent(v string) *FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem {
	s.Content = &v
	return s
}

func (s *FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem) SetCreateTime(v int64) *FeedbackGetFeedbackListResponseDataFeedbackListItemContentListItem {
	s.CreateTime = &v
	return s
}

type FeedbackGetReplyResultRequest struct {
	Appid       *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	FeedbackId  *int64             `json:"feedback_id,omitempty" xml:"feedback_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FeedbackGetReplyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetReplyResultRequest) GoString() string {
	return s.String()
}

func (s *FeedbackGetReplyResultRequest) SetAppid(v string) *FeedbackGetReplyResultRequest {
	s.Appid = &v
	return s
}

func (s *FeedbackGetReplyResultRequest) SetFeedbackId(v int64) *FeedbackGetReplyResultRequest {
	s.FeedbackId = &v
	return s
}

func (s *FeedbackGetReplyResultRequest) SetHeader(v map[string]*string) *FeedbackGetReplyResultRequest {
	s.Header = v
	return s
}

func (s *FeedbackGetReplyResultRequest) SetAccessToken(v string) *FeedbackGetReplyResultRequest {
	s.AccessToken = &v
	return s
}

type FeedbackGetReplyResultResponse struct {
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *FeedbackGetReplyResultResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FeedbackGetReplyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetReplyResultResponse) GoString() string {
	return s.String()
}

func (s *FeedbackGetReplyResultResponse) SetErrNo(v int32) *FeedbackGetReplyResultResponse {
	s.ErrNo = &v
	return s
}

func (s *FeedbackGetReplyResultResponse) SetErrMsg(v string) *FeedbackGetReplyResultResponse {
	s.ErrMsg = &v
	return s
}

func (s *FeedbackGetReplyResultResponse) SetLogId(v string) *FeedbackGetReplyResultResponse {
	s.LogId = &v
	return s
}

func (s *FeedbackGetReplyResultResponse) SetData(v *FeedbackGetReplyResultResponseData) *FeedbackGetReplyResultResponse {
	s.Data = v
	return s
}

type FeedbackGetReplyResultResponseData struct {
	ReplyList []*FeedbackGetReplyResultResponseDataReplyListItem `json:"reply_list,omitempty" xml:"reply_list,omitempty" type:"Repeated"`
}

func (s FeedbackGetReplyResultResponseData) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetReplyResultResponseData) GoString() string {
	return s.String()
}

func (s *FeedbackGetReplyResultResponseData) SetReplyList(v []*FeedbackGetReplyResultResponseDataReplyListItem) *FeedbackGetReplyResultResponseData {
	s.ReplyList = v
	return s
}

type FeedbackGetReplyResultResponseDataReplyListItem struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	ReplyState *int    `json:"reply_state,omitempty" xml:"reply_state,omitempty" require:"true"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty" require:"true"`
}

func (s FeedbackGetReplyResultResponseDataReplyListItem) String() string {
	return tea.Prettify(s)
}

func (s FeedbackGetReplyResultResponseDataReplyListItem) GoString() string {
	return s.String()
}

func (s *FeedbackGetReplyResultResponseDataReplyListItem) SetContent(v string) *FeedbackGetReplyResultResponseDataReplyListItem {
	s.Content = &v
	return s
}

func (s *FeedbackGetReplyResultResponseDataReplyListItem) SetReplyState(v int) *FeedbackGetReplyResultResponseDataReplyListItem {
	s.ReplyState = &v
	return s
}

func (s *FeedbackGetReplyResultResponseDataReplyListItem) SetCreateTime(v int64) *FeedbackGetReplyResultResponseDataReplyListItem {
	s.CreateTime = &v
	return s
}

func (s *FeedbackGetReplyResultResponseDataReplyListItem) SetId(v int64) *FeedbackGetReplyResultResponseDataReplyListItem {
	s.Id = &v
	return s
}

type FileUploadRequest struct {
	FileContentBase64 *string            `json:"file_content_base64,omitempty" xml:"file_content_base64,omitempty"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FileUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s FileUploadRequest) GoString() string {
	return s.String()
}

func (s *FileUploadRequest) SetFileContentBase64(v string) *FileUploadRequest {
	s.FileContentBase64 = &v
	return s
}

func (s *FileUploadRequest) SetHeader(v map[string]*string) *FileUploadRequest {
	s.Header = v
	return s
}

func (s *FileUploadRequest) SetAccessToken(v string) *FileUploadRequest {
	s.AccessToken = &v
	return s
}

type FileUploadResponse struct {
	Extra *FileUploadResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *FileUploadResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FileUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponse) GoString() string {
	return s.String()
}

func (s *FileUploadResponse) SetExtra(v *FileUploadResponseExtra) *FileUploadResponse {
	s.Extra = v
	return s
}

func (s *FileUploadResponse) SetData(v *FileUploadResponseData) *FileUploadResponse {
	s.Data = v
	return s
}

type FileUploadResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Uri           *string `json:"uri,omitempty" xml:"uri,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FileUploadResponseData) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponseData) GoString() string {
	return s.String()
}

func (s *FileUploadResponseData) SetGwDescription(v string) *FileUploadResponseData {
	s.GwDescription = &v
	return s
}

func (s *FileUploadResponseData) SetUri(v string) *FileUploadResponseData {
	s.Uri = &v
	return s
}

func (s *FileUploadResponseData) SetGwErrorCode(v int32) *FileUploadResponseData {
	s.GwErrorCode = &v
	return s
}

type FileUploadResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s FileUploadResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponseExtra) GoString() string {
	return s.String()
}

func (s *FileUploadResponseExtra) SetSubErrorCode(v int32) *FileUploadResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FileUploadResponseExtra) SetDescription(v string) *FileUploadResponseExtra {
	s.Description = &v
	return s
}

func (s *FileUploadResponseExtra) SetErrorCode(v int32) *FileUploadResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FileUploadResponseExtra) SetLogid(v string) *FileUploadResponseExtra {
	s.Logid = &v
	return s
}

func (s *FileUploadResponseExtra) SetNow(v int64) *FileUploadResponseExtra {
	s.Now = &v
	return s
}

func (s *FileUploadResponseExtra) SetSubDescription(v string) *FileUploadResponseExtra {
	s.SubDescription = &v
	return s
}

type FoodNewRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodNewRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodNewRequest) GoString() string {
	return s.String()
}

func (s *FoodNewRequest) SetHeader(v map[string]*string) *FoodNewRequest {
	s.Header = v
	return s
}

func (s *FoodNewRequest) SetAccessToken(v string) *FoodNewRequest {
	s.AccessToken = &v
	return s
}

type FoodNewResponse struct {
	Data  *FoodNewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *FoodNewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FoodNewResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodNewResponse) GoString() string {
	return s.String()
}

func (s *FoodNewResponse) SetData(v *FoodNewResponseData) *FoodNewResponse {
	s.Data = v
	return s
}

func (s *FoodNewResponse) SetExtra(v *FoodNewResponseExtra) *FoodNewResponse {
	s.Extra = v
	return s
}

type FoodNewResponseData struct {
	List          []*FoodNewResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FoodNewResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodNewResponseData) GoString() string {
	return s.String()
}

func (s *FoodNewResponseData) SetList(v []*FoodNewResponseDataListItem) *FoodNewResponseData {
	s.List = v
	return s
}

func (s *FoodNewResponseData) SetGwErrorCode(v int32) *FoodNewResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FoodNewResponseData) SetGwDescription(v string) *FoodNewResponseData {
	s.GwDescription = &v
	return s
}

type FoodNewResponseDataListItem struct {
	Nickname         *string                                     `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                     `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                      `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                      `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                    `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*FoodNewResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                      `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                     `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
}

func (s FoodNewResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodNewResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FoodNewResponseDataListItem) SetNickname(v string) *FoodNewResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetAvatar(v string) *FoodNewResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetFollowerCount(v int64) *FoodNewResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetOnbillbaordTimes(v int32) *FoodNewResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetEffectValue(v float64) *FoodNewResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetVideoList(v []*FoodNewResponseDataListItemVideoListItem) *FoodNewResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *FoodNewResponseDataListItem) SetRank(v int32) *FoodNewResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *FoodNewResponseDataListItem) SetRankChange(v string) *FoodNewResponseDataListItem {
	s.RankChange = &v
	return s
}

type FoodNewResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s FoodNewResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodNewResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *FoodNewResponseDataListItemVideoListItem) SetTitle(v string) *FoodNewResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *FoodNewResponseDataListItemVideoListItem) SetItemCover(v string) *FoodNewResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *FoodNewResponseDataListItemVideoListItem) SetShareUrl(v string) *FoodNewResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type FoodNewResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s FoodNewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodNewResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodNewResponseExtra) SetErrorCode(v int32) *FoodNewResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodNewResponseExtra) SetDescription(v string) *FoodNewResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodNewResponseExtra) SetSubErrorCode(v int32) *FoodNewResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FoodNewResponseExtra) SetSubDescription(v string) *FoodNewResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodNewResponseExtra) SetLogid(v string) *FoodNewResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodNewResponseExtra) SetNow(v int64) *FoodNewResponseExtra {
	s.Now = &v
	return s
}

type FoodOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallRequest) GoString() string {
	return s.String()
}

func (s *FoodOverallRequest) SetHeader(v map[string]*string) *FoodOverallRequest {
	s.Header = v
	return s
}

func (s *FoodOverallRequest) SetAccessToken(v string) *FoodOverallRequest {
	s.AccessToken = &v
	return s
}

type FoodOverallResponse struct {
	Extra *FoodOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *FoodOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FoodOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallResponse) GoString() string {
	return s.String()
}

func (s *FoodOverallResponse) SetExtra(v *FoodOverallResponseExtra) *FoodOverallResponse {
	s.Extra = v
	return s
}

func (s *FoodOverallResponse) SetData(v *FoodOverallResponseData) *FoodOverallResponse {
	s.Data = v
	return s
}

type FoodOverallResponseData struct {
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*FoodOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FoodOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallResponseData) GoString() string {
	return s.String()
}

func (s *FoodOverallResponseData) SetGwDescription(v string) *FoodOverallResponseData {
	s.GwDescription = &v
	return s
}

func (s *FoodOverallResponseData) SetList(v []*FoodOverallResponseDataListItem) *FoodOverallResponseData {
	s.List = v
	return s
}

func (s *FoodOverallResponseData) SetGwErrorCode(v int32) *FoodOverallResponseData {
	s.GwErrorCode = &v
	return s
}

type FoodOverallResponseDataListItem struct {
	Avatar           *string                                         `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                          `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                          `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                        `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*FoodOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                          `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                         `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                         `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
}

func (s FoodOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FoodOverallResponseDataListItem) SetAvatar(v string) *FoodOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetFollowerCount(v int64) *FoodOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *FoodOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetEffectValue(v float64) *FoodOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetVideoList(v []*FoodOverallResponseDataListItemVideoListItem) *FoodOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *FoodOverallResponseDataListItem) SetRank(v int32) *FoodOverallResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetRankChange(v string) *FoodOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *FoodOverallResponseDataListItem) SetNickname(v string) *FoodOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

type FoodOverallResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s FoodOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *FoodOverallResponseDataListItemVideoListItem) SetTitle(v string) *FoodOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *FoodOverallResponseDataListItemVideoListItem) SetItemCover(v string) *FoodOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *FoodOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *FoodOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type FoodOverallResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s FoodOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodOverallResponseExtra) SetErrorCode(v int32) *FoodOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodOverallResponseExtra) SetDescription(v string) *FoodOverallResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodOverallResponseExtra) SetSubErrorCode(v int32) *FoodOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FoodOverallResponseExtra) SetSubDescription(v string) *FoodOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodOverallResponseExtra) SetLogid(v string) *FoodOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodOverallResponseExtra) SetNow(v int64) *FoodOverallResponseExtra {
	s.Now = &v
	return s
}

type FoodShopRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodShopRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodShopRequest) GoString() string {
	return s.String()
}

func (s *FoodShopRequest) SetHeader(v map[string]*string) *FoodShopRequest {
	s.Header = v
	return s
}

func (s *FoodShopRequest) SetAccessToken(v string) *FoodShopRequest {
	s.AccessToken = &v
	return s
}

type FoodShopResponse struct {
	Extra *FoodShopResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *FoodShopResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FoodShopResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodShopResponse) GoString() string {
	return s.String()
}

func (s *FoodShopResponse) SetExtra(v *FoodShopResponseExtra) *FoodShopResponse {
	s.Extra = v
	return s
}

func (s *FoodShopResponse) SetData(v *FoodShopResponseData) *FoodShopResponse {
	s.Data = v
	return s
}

type FoodShopResponseData struct {
	GwDescription *string                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*FoodShopResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FoodShopResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodShopResponseData) GoString() string {
	return s.String()
}

func (s *FoodShopResponseData) SetGwDescription(v string) *FoodShopResponseData {
	s.GwDescription = &v
	return s
}

func (s *FoodShopResponseData) SetList(v []*FoodShopResponseDataListItem) *FoodShopResponseData {
	s.List = v
	return s
}

func (s *FoodShopResponseData) SetGwErrorCode(v int32) *FoodShopResponseData {
	s.GwErrorCode = &v
	return s
}

type FoodShopResponseDataListItem struct {
	RankChange       *string                                      `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                      `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                      `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                       `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                       `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                     `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*FoodShopResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                       `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s FoodShopResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodShopResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FoodShopResponseDataListItem) SetRankChange(v string) *FoodShopResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetNickname(v string) *FoodShopResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetAvatar(v string) *FoodShopResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetFollowerCount(v int64) *FoodShopResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetOnbillbaordTimes(v int32) *FoodShopResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetEffectValue(v float64) *FoodShopResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *FoodShopResponseDataListItem) SetVideoList(v []*FoodShopResponseDataListItemVideoListItem) *FoodShopResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *FoodShopResponseDataListItem) SetRank(v int32) *FoodShopResponseDataListItem {
	s.Rank = &v
	return s
}

type FoodShopResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s FoodShopResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodShopResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *FoodShopResponseDataListItemVideoListItem) SetShareUrl(v string) *FoodShopResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *FoodShopResponseDataListItemVideoListItem) SetTitle(v string) *FoodShopResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *FoodShopResponseDataListItemVideoListItem) SetItemCover(v string) *FoodShopResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type FoodShopResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FoodShopResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodShopResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodShopResponseExtra) SetDescription(v string) *FoodShopResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodShopResponseExtra) SetSubErrorCode(v int32) *FoodShopResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FoodShopResponseExtra) SetSubDescription(v string) *FoodShopResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodShopResponseExtra) SetLogid(v string) *FoodShopResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodShopResponseExtra) SetNow(v int64) *FoodShopResponseExtra {
	s.Now = &v
	return s
}

func (s *FoodShopResponseExtra) SetErrorCode(v int32) *FoodShopResponseExtra {
	s.ErrorCode = &v
	return s
}

type FoodTutorialRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodTutorialRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialRequest) GoString() string {
	return s.String()
}

func (s *FoodTutorialRequest) SetHeader(v map[string]*string) *FoodTutorialRequest {
	s.Header = v
	return s
}

func (s *FoodTutorialRequest) SetAccessToken(v string) *FoodTutorialRequest {
	s.AccessToken = &v
	return s
}

type FoodTutorialResponse struct {
	Data  *FoodTutorialResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *FoodTutorialResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FoodTutorialResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialResponse) GoString() string {
	return s.String()
}

func (s *FoodTutorialResponse) SetData(v *FoodTutorialResponseData) *FoodTutorialResponse {
	s.Data = v
	return s
}

func (s *FoodTutorialResponse) SetExtra(v *FoodTutorialResponseExtra) *FoodTutorialResponse {
	s.Extra = v
	return s
}

type FoodTutorialResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*FoodTutorialResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s FoodTutorialResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialResponseData) GoString() string {
	return s.String()
}

func (s *FoodTutorialResponseData) SetGwErrorCode(v int32) *FoodTutorialResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FoodTutorialResponseData) SetGwDescription(v string) *FoodTutorialResponseData {
	s.GwDescription = &v
	return s
}

func (s *FoodTutorialResponseData) SetList(v []*FoodTutorialResponseDataListItem) *FoodTutorialResponseData {
	s.List = v
	return s
}

type FoodTutorialResponseDataListItem struct {
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*FoodTutorialResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
}

func (s FoodTutorialResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialResponseDataListItem) GoString() string {
	return s.String()
}

func (s *FoodTutorialResponseDataListItem) SetNickname(v string) *FoodTutorialResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetAvatar(v string) *FoodTutorialResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetFollowerCount(v int64) *FoodTutorialResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetOnbillbaordTimes(v int32) *FoodTutorialResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetEffectValue(v float64) *FoodTutorialResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetVideoList(v []*FoodTutorialResponseDataListItemVideoListItem) *FoodTutorialResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetRank(v int32) *FoodTutorialResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *FoodTutorialResponseDataListItem) SetRankChange(v string) *FoodTutorialResponseDataListItem {
	s.RankChange = &v
	return s
}

type FoodTutorialResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FoodTutorialResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *FoodTutorialResponseDataListItemVideoListItem) SetItemCover(v string) *FoodTutorialResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *FoodTutorialResponseDataListItemVideoListItem) SetShareUrl(v string) *FoodTutorialResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *FoodTutorialResponseDataListItemVideoListItem) SetTitle(v string) *FoodTutorialResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type FoodTutorialResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s FoodTutorialResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodTutorialResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodTutorialResponseExtra) SetSubDescription(v string) *FoodTutorialResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodTutorialResponseExtra) SetLogid(v string) *FoodTutorialResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodTutorialResponseExtra) SetNow(v int64) *FoodTutorialResponseExtra {
	s.Now = &v
	return s
}

func (s *FoodTutorialResponseExtra) SetErrorCode(v int32) *FoodTutorialResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodTutorialResponseExtra) SetDescription(v string) *FoodTutorialResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodTutorialResponseExtra) SetSubErrorCode(v int32) *FoodTutorialResponseExtra {
	s.SubErrorCode = &v
	return s
}

type FoodorderPoistockSaveRequest struct {
	PoiSkuStockList []*FoodorderPoistockSaveRequestPoiSkuStockListItem `json:"poi_sku_stock_list,omitempty" xml:"poi_sku_stock_list,omitempty" require:"true" type:"Repeated"`
	Base            *FoodorderPoistockSaveRequestBase                  `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *int64                                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	IgnoreFailPoi   *bool                                              `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
	OutProductId    *string                                            `json:"out_product_id,omitempty" xml:"out_product_id,omitempty" require:"true"`
	Header          map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodorderPoistockSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequest) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequest) SetPoiSkuStockList(v []*FoodorderPoistockSaveRequestPoiSkuStockListItem) *FoodorderPoistockSaveRequest {
	s.PoiSkuStockList = v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetBase(v *FoodorderPoistockSaveRequestBase) *FoodorderPoistockSaveRequest {
	s.Base = v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetAccountId(v int64) *FoodorderPoistockSaveRequest {
	s.AccountId = &v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetIgnoreFailPoi(v bool) *FoodorderPoistockSaveRequest {
	s.IgnoreFailPoi = &v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetOutProductId(v string) *FoodorderPoistockSaveRequest {
	s.OutProductId = &v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetHeader(v map[string]*string) *FoodorderPoistockSaveRequest {
	s.Header = v
	return s
}

func (s *FoodorderPoistockSaveRequest) SetAccessToken(v string) *FoodorderPoistockSaveRequest {
	s.AccessToken = &v
	return s
}

type FoodorderPoistockSaveRequestBase struct {
	Addr       *string                                     `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                     `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                     `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                          `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                     `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *FoodorderPoistockSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s FoodorderPoistockSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequestBase) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequestBase) SetAddr(v string) *FoodorderPoistockSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *FoodorderPoistockSaveRequestBase) SetCaller(v string) *FoodorderPoistockSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *FoodorderPoistockSaveRequestBase) SetClient(v string) *FoodorderPoistockSaveRequestBase {
	s.Client = &v
	return s
}

func (s *FoodorderPoistockSaveRequestBase) SetExtra(v map[string]*string) *FoodorderPoistockSaveRequestBase {
	s.Extra = v
	return s
}

func (s *FoodorderPoistockSaveRequestBase) SetLogID(v string) *FoodorderPoistockSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *FoodorderPoistockSaveRequestBase) SetTrafficEnv(v *FoodorderPoistockSaveRequestBaseTrafficEnv) *FoodorderPoistockSaveRequestBase {
	s.TrafficEnv = v
	return s
}

type FoodorderPoistockSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s FoodorderPoistockSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequestBaseTrafficEnv) SetEnv(v string) *FoodorderPoistockSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *FoodorderPoistockSaveRequestBaseTrafficEnv) SetOpen(v bool) *FoodorderPoistockSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type FoodorderPoistockSaveRequestPoiSkuStockListItem struct {
	Poi   *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi   `json:"poi,omitempty" xml:"poi,omitempty"`
	Stock *FoodorderPoistockSaveRequestPoiSkuStockListItemStock `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItem) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItem) SetPoi(v *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi) *FoodorderPoistockSaveRequestPoiSkuStockListItem {
	s.Poi = v
	return s
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItem) SetStock(v *FoodorderPoistockSaveRequestPoiSkuStockListItemStock) *FoodorderPoistockSaveRequestPoiSkuStockListItem {
	s.Stock = v
	return s
}

type FoodorderPoistockSaveRequestPoiSkuStockListItemPoi struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItemPoi) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi) SetExtId(v string) *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi {
	s.ExtId = &v
	return s
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi) SetPoiId(v int64) *FoodorderPoistockSaveRequestPoiSkuStockListItemPoi {
	s.PoiId = &v
	return s
}

type FoodorderPoistockSaveRequestPoiSkuStockListItemStock struct {
	Stock    *int64  `json:"stock,omitempty" xml:"stock,omitempty"`
	OutSkuId *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItemStock) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveRequestPoiSkuStockListItemStock) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItemStock) SetStock(v int64) *FoodorderPoistockSaveRequestPoiSkuStockListItemStock {
	s.Stock = &v
	return s
}

func (s *FoodorderPoistockSaveRequestPoiSkuStockListItemStock) SetOutSkuId(v string) *FoodorderPoistockSaveRequestPoiSkuStockListItemStock {
	s.OutSkuId = &v
	return s
}

type FoodorderPoistockSaveResponse struct {
	Extra    *FoodorderPoistockSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *FoodorderPoistockSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *FoodorderPoistockSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s FoodorderPoistockSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveResponse) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveResponse) SetExtra(v *FoodorderPoistockSaveResponseExtra) *FoodorderPoistockSaveResponse {
	s.Extra = v
	return s
}

func (s *FoodorderPoistockSaveResponse) SetBaseResp(v *FoodorderPoistockSaveResponseBaseResp) *FoodorderPoistockSaveResponse {
	s.BaseResp = v
	return s
}

func (s *FoodorderPoistockSaveResponse) SetData(v *FoodorderPoistockSaveResponseData) *FoodorderPoistockSaveResponse {
	s.Data = v
	return s
}

type FoodorderPoistockSaveResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s FoodorderPoistockSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveResponseBaseResp) SetExtra(v map[string]*string) *FoodorderPoistockSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *FoodorderPoistockSaveResponseBaseResp) SetStatusCode(v int32) *FoodorderPoistockSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *FoodorderPoistockSaveResponseBaseResp) SetStatusMessage(v string) *FoodorderPoistockSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type FoodorderPoistockSaveResponseData struct {
	ErrorCode   *int32                                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	FailPoiList []*FoodorderPoistockSaveResponseDataFailPoiListItem `json:"fail_poi_list,omitempty" xml:"fail_poi_list,omitempty" type:"Repeated"`
	Description *string                                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FoodorderPoistockSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveResponseData) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveResponseData) SetErrorCode(v int32) *FoodorderPoistockSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *FoodorderPoistockSaveResponseData) SetFailPoiList(v []*FoodorderPoistockSaveResponseDataFailPoiListItem) *FoodorderPoistockSaveResponseData {
	s.FailPoiList = v
	return s
}

func (s *FoodorderPoistockSaveResponseData) SetDescription(v string) *FoodorderPoistockSaveResponseData {
	s.Description = &v
	return s
}

type FoodorderPoistockSaveResponseDataFailPoiListItem struct {
	ExtId  *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId  *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s FoodorderPoistockSaveResponseDataFailPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveResponseDataFailPoiListItem) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveResponseDataFailPoiListItem) SetExtId(v string) *FoodorderPoistockSaveResponseDataFailPoiListItem {
	s.ExtId = &v
	return s
}

func (s *FoodorderPoistockSaveResponseDataFailPoiListItem) SetPoiId(v int64) *FoodorderPoistockSaveResponseDataFailPoiListItem {
	s.PoiId = &v
	return s
}

func (s *FoodorderPoistockSaveResponseDataFailPoiListItem) SetReason(v string) *FoodorderPoistockSaveResponseDataFailPoiListItem {
	s.Reason = &v
	return s
}

func (s *FoodorderPoistockSaveResponseDataFailPoiListItem) SetCode(v string) *FoodorderPoistockSaveResponseDataFailPoiListItem {
	s.Code = &v
	return s
}

type FoodorderPoistockSaveResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FoodorderPoistockSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodorderPoistockSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodorderPoistockSaveResponseExtra) SetErrorCode(v int32) *FoodorderPoistockSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodorderPoistockSaveResponseExtra) SetLogid(v string) *FoodorderPoistockSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodorderPoistockSaveResponseExtra) SetNow(v int64) *FoodorderPoistockSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *FoodorderPoistockSaveResponseExtra) SetSubDescription(v string) *FoodorderPoistockSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodorderPoistockSaveResponseExtra) SetSubErrorCode(v int32) *FoodorderPoistockSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FoodorderPoistockSaveResponseExtra) SetDescription(v string) *FoodorderPoistockSaveResponseExtra {
	s.Description = &v
	return s
}

type FoodorderProductOperateRequest struct {
	AccessToken *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OutId       *string                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId   *int64                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Base        *FoodorderProductOperateRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *int64                              `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Op          *int                                `json:"op,omitempty" xml:"op,omitempty" require:"true"`
	Header      map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
}

func (s FoodorderProductOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateRequest) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateRequest) SetAccessToken(v string) *FoodorderProductOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *FoodorderProductOperateRequest) SetOutId(v string) *FoodorderProductOperateRequest {
	s.OutId = &v
	return s
}

func (s *FoodorderProductOperateRequest) SetProductId(v int64) *FoodorderProductOperateRequest {
	s.ProductId = &v
	return s
}

func (s *FoodorderProductOperateRequest) SetBase(v *FoodorderProductOperateRequestBase) *FoodorderProductOperateRequest {
	s.Base = v
	return s
}

func (s *FoodorderProductOperateRequest) SetAccountId(v int64) *FoodorderProductOperateRequest {
	s.AccountId = &v
	return s
}

func (s *FoodorderProductOperateRequest) SetOp(v int) *FoodorderProductOperateRequest {
	s.Op = &v
	return s
}

func (s *FoodorderProductOperateRequest) SetHeader(v map[string]*string) *FoodorderProductOperateRequest {
	s.Header = v
	return s
}

type FoodorderProductOperateRequestBase struct {
	Extra      map[string]*string                            `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                       `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *FoodorderProductOperateRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                       `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                       `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                       `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s FoodorderProductOperateRequestBase) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateRequestBase) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateRequestBase) SetExtra(v map[string]*string) *FoodorderProductOperateRequestBase {
	s.Extra = v
	return s
}

func (s *FoodorderProductOperateRequestBase) SetLogID(v string) *FoodorderProductOperateRequestBase {
	s.LogID = &v
	return s
}

func (s *FoodorderProductOperateRequestBase) SetTrafficEnv(v *FoodorderProductOperateRequestBaseTrafficEnv) *FoodorderProductOperateRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *FoodorderProductOperateRequestBase) SetAddr(v string) *FoodorderProductOperateRequestBase {
	s.Addr = &v
	return s
}

func (s *FoodorderProductOperateRequestBase) SetCaller(v string) *FoodorderProductOperateRequestBase {
	s.Caller = &v
	return s
}

func (s *FoodorderProductOperateRequestBase) SetClient(v string) *FoodorderProductOperateRequestBase {
	s.Client = &v
	return s
}

type FoodorderProductOperateRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s FoodorderProductOperateRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateRequestBaseTrafficEnv) SetEnv(v string) *FoodorderProductOperateRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *FoodorderProductOperateRequestBaseTrafficEnv) SetOpen(v bool) *FoodorderProductOperateRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type FoodorderProductOperateResponse struct {
	BaseResp *FoodorderProductOperateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *FoodorderProductOperateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *FoodorderProductOperateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FoodorderProductOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateResponse) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateResponse) SetBaseResp(v *FoodorderProductOperateResponseBaseResp) *FoodorderProductOperateResponse {
	s.BaseResp = v
	return s
}

func (s *FoodorderProductOperateResponse) SetData(v *FoodorderProductOperateResponseData) *FoodorderProductOperateResponse {
	s.Data = v
	return s
}

func (s *FoodorderProductOperateResponse) SetExtra(v *FoodorderProductOperateResponseExtra) *FoodorderProductOperateResponse {
	s.Extra = v
	return s
}

type FoodorderProductOperateResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s FoodorderProductOperateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateResponseBaseResp) SetExtra(v map[string]*string) *FoodorderProductOperateResponseBaseResp {
	s.Extra = v
	return s
}

func (s *FoodorderProductOperateResponseBaseResp) SetStatusCode(v int32) *FoodorderProductOperateResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *FoodorderProductOperateResponseBaseResp) SetStatusMessage(v string) *FoodorderProductOperateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type FoodorderProductOperateResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s FoodorderProductOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateResponseData) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateResponseData) SetDescription(v string) *FoodorderProductOperateResponseData {
	s.Description = &v
	return s
}

func (s *FoodorderProductOperateResponseData) SetErrorCode(v int32) *FoodorderProductOperateResponseData {
	s.ErrorCode = &v
	return s
}

func (s *FoodorderProductOperateResponseData) SetProductId(v int64) *FoodorderProductOperateResponseData {
	s.ProductId = &v
	return s
}

type FoodorderProductOperateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s FoodorderProductOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodorderProductOperateResponseExtra) SetDescription(v string) *FoodorderProductOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodorderProductOperateResponseExtra) SetErrorCode(v int32) *FoodorderProductOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodorderProductOperateResponseExtra) SetLogid(v string) *FoodorderProductOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodorderProductOperateResponseExtra) SetNow(v int64) *FoodorderProductOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *FoodorderProductOperateResponseExtra) SetSubDescription(v string) *FoodorderProductOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodorderProductOperateResponseExtra) SetSubErrorCode(v int32) *FoodorderProductOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type FoodorderProductSaveRequest struct {
	Product       *FoodorderProductSaveRequestProduct `json:"product,omitempty" xml:"product,omitempty" require:"true"`
	SupportsLaike *bool                               `json:"supports_laike,omitempty" xml:"supports_laike,omitempty"`
	AccountId     *int64                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	IgnoreFailPoi *bool                               `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
	Header        map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FoodorderProductSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequest) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequest) SetProduct(v *FoodorderProductSaveRequestProduct) *FoodorderProductSaveRequest {
	s.Product = v
	return s
}

func (s *FoodorderProductSaveRequest) SetSupportsLaike(v bool) *FoodorderProductSaveRequest {
	s.SupportsLaike = &v
	return s
}

func (s *FoodorderProductSaveRequest) SetAccountId(v int64) *FoodorderProductSaveRequest {
	s.AccountId = &v
	return s
}

func (s *FoodorderProductSaveRequest) SetIgnoreFailPoi(v bool) *FoodorderProductSaveRequest {
	s.IgnoreFailPoi = &v
	return s
}

func (s *FoodorderProductSaveRequest) SetHeader(v map[string]*string) *FoodorderProductSaveRequest {
	s.Header = v
	return s
}

func (s *FoodorderProductSaveRequest) SetAccessToken(v string) *FoodorderProductSaveRequest {
	s.AccessToken = &v
	return s
}

type FoodorderProductSaveRequestProduct struct {
	AccountId         *int64                                                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Name              *string                                                   `json:"name,omitempty" xml:"name,omitempty"`
	OutSpuId          *string                                                   `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	Images            *FoodorderProductSaveRequestProductImages                 `json:"images,omitempty" xml:"images,omitempty"`
	FulfillmentMethod []*int                                                    `json:"fulfillment_method,omitempty" xml:"fulfillment_method,omitempty" type:"Repeated"`
	TestInfo          *FoodorderProductSaveRequestProductTestInfo               `json:"test_info,omitempty" xml:"test_info,omitempty"`
	SoldTime          *FoodorderProductSaveRequestProductSoldTime               `json:"sold_time,omitempty" xml:"sold_time,omitempty"`
	SkuList           []*FoodorderProductSaveRequestProductSkuListItem          `json:"sku_list,omitempty" xml:"sku_list,omitempty" type:"Repeated"`
	AffiliatedGroups  []*FoodorderProductSaveRequestProductAffiliatedGroupsItem `json:"affiliated_groups,omitempty" xml:"affiliated_groups,omitempty" type:"Repeated"`
	OutId             *string                                                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	OutUrl            *string                                                   `json:"out_url,omitempty" xml:"out_url,omitempty"`
	PackFee           *FoodorderProductSaveRequestProductPackFee                `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	SaleAttrGroups    []*FoodorderProductSaveRequestProductSaleAttrGroupsItem   `json:"sale_attr_groups,omitempty" xml:"sale_attr_groups,omitempty" type:"Repeated"`
	PoiList           []*FoodorderProductSaveRequestProductPoiListItem          `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	IsSinglePoiDish   *bool                                                     `json:"is_single_poi_dish,omitempty" xml:"is_single_poi_dish,omitempty"`
	SettleInfo        *FoodorderProductSaveRequestProductSettleInfo             `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	CategoryId        *int64                                                    `json:"category_id,omitempty" xml:"category_id,omitempty"`
	BizLine           *int                                                      `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
}

func (s FoodorderProductSaveRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProduct) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProduct) SetAccountId(v int64) *FoodorderProductSaveRequestProduct {
	s.AccountId = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetName(v string) *FoodorderProductSaveRequestProduct {
	s.Name = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetOutSpuId(v string) *FoodorderProductSaveRequestProduct {
	s.OutSpuId = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetImages(v *FoodorderProductSaveRequestProductImages) *FoodorderProductSaveRequestProduct {
	s.Images = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetFulfillmentMethod(v []*int) *FoodorderProductSaveRequestProduct {
	s.FulfillmentMethod = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetTestInfo(v *FoodorderProductSaveRequestProductTestInfo) *FoodorderProductSaveRequestProduct {
	s.TestInfo = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetSoldTime(v *FoodorderProductSaveRequestProductSoldTime) *FoodorderProductSaveRequestProduct {
	s.SoldTime = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetSkuList(v []*FoodorderProductSaveRequestProductSkuListItem) *FoodorderProductSaveRequestProduct {
	s.SkuList = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetAffiliatedGroups(v []*FoodorderProductSaveRequestProductAffiliatedGroupsItem) *FoodorderProductSaveRequestProduct {
	s.AffiliatedGroups = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetOutId(v string) *FoodorderProductSaveRequestProduct {
	s.OutId = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetOutUrl(v string) *FoodorderProductSaveRequestProduct {
	s.OutUrl = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetPackFee(v *FoodorderProductSaveRequestProductPackFee) *FoodorderProductSaveRequestProduct {
	s.PackFee = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetSaleAttrGroups(v []*FoodorderProductSaveRequestProductSaleAttrGroupsItem) *FoodorderProductSaveRequestProduct {
	s.SaleAttrGroups = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetPoiList(v []*FoodorderProductSaveRequestProductPoiListItem) *FoodorderProductSaveRequestProduct {
	s.PoiList = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetIsSinglePoiDish(v bool) *FoodorderProductSaveRequestProduct {
	s.IsSinglePoiDish = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetSettleInfo(v *FoodorderProductSaveRequestProductSettleInfo) *FoodorderProductSaveRequestProduct {
	s.SettleInfo = v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetCategoryId(v int64) *FoodorderProductSaveRequestProduct {
	s.CategoryId = &v
	return s
}

func (s *FoodorderProductSaveRequestProduct) SetBizLine(v int) *FoodorderProductSaveRequestProduct {
	s.BizLine = &v
	return s
}

type FoodorderProductSaveRequestProductAffiliatedGroupsItem struct {
	OutAffiliatedIds []*string `json:"out_affiliated_ids,omitempty" xml:"out_affiliated_ids,omitempty" type:"Repeated"`
}

func (s FoodorderProductSaveRequestProductAffiliatedGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductAffiliatedGroupsItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductAffiliatedGroupsItem) SetOutAffiliatedIds(v []*string) *FoodorderProductSaveRequestProductAffiliatedGroupsItem {
	s.OutAffiliatedIds = v
	return s
}

type FoodorderProductSaveRequestProductImages struct {
	HeadImage *FoodorderProductSaveRequestProductImagesHeadImage `json:"head_image,omitempty" xml:"head_image,omitempty"`
}

func (s FoodorderProductSaveRequestProductImages) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductImages) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductImages) SetHeadImage(v *FoodorderProductSaveRequestProductImagesHeadImage) *FoodorderProductSaveRequestProductImages {
	s.HeadImage = v
	return s
}

type FoodorderProductSaveRequestProductImagesHeadImage struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s FoodorderProductSaveRequestProductImagesHeadImage) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductImagesHeadImage) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductImagesHeadImage) SetUrl(v string) *FoodorderProductSaveRequestProductImagesHeadImage {
	s.Url = &v
	return s
}

type FoodorderProductSaveRequestProductPackFee struct {
	Step        *int64 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int64 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s FoodorderProductSaveRequestProductPackFee) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductPackFee) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductPackFee) SetStep(v int64) *FoodorderProductSaveRequestProductPackFee {
	s.Step = &v
	return s
}

func (s *FoodorderProductSaveRequestProductPackFee) SetPackFee(v int64) *FoodorderProductSaveRequestProductPackFee {
	s.PackFee = &v
	return s
}

func (s *FoodorderProductSaveRequestProductPackFee) SetPackFeeUnit(v int) *FoodorderProductSaveRequestProductPackFee {
	s.PackFeeUnit = &v
	return s
}

type FoodorderProductSaveRequestProductPoiListItem struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s FoodorderProductSaveRequestProductPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductPoiListItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductPoiListItem) SetPoiId(v int64) *FoodorderProductSaveRequestProductPoiListItem {
	s.PoiId = &v
	return s
}

func (s *FoodorderProductSaveRequestProductPoiListItem) SetExtId(v string) *FoodorderProductSaveRequestProductPoiListItem {
	s.ExtId = &v
	return s
}

type FoodorderProductSaveRequestProductSaleAttrGroupsItem struct {
	ItemList  []*FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	SpecType  *int                                                                `json:"spec_type,omitempty" xml:"spec_type,omitempty"`
	GroupCode *string                                                             `json:"group_code,omitempty" xml:"group_code,omitempty"`
	GroupName *string                                                             `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s FoodorderProductSaveRequestProductSaleAttrGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSaleAttrGroupsItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItem) SetItemList(v []*FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) *FoodorderProductSaveRequestProductSaleAttrGroupsItem {
	s.ItemList = v
	return s
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItem) SetSpecType(v int) *FoodorderProductSaveRequestProductSaleAttrGroupsItem {
	s.SpecType = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItem) SetGroupCode(v string) *FoodorderProductSaveRequestProductSaleAttrGroupsItem {
	s.GroupCode = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItem) SetGroupName(v string) *FoodorderProductSaveRequestProductSaleAttrGroupsItem {
	s.GroupName = &v
	return s
}

type FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem struct {
	ItemKey  *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	Checked  *bool   `json:"checked,omitempty" xml:"checked,omitempty"`
}

func (s FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) SetItemKey(v string) *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem {
	s.ItemKey = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) SetItemName(v string) *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem) SetChecked(v bool) *FoodorderProductSaveRequestProductSaleAttrGroupsItemItemListItem {
	s.Checked = &v
	return s
}

type FoodorderProductSaveRequestProductSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s FoodorderProductSaveRequestProductSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSettleInfo) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSettleInfo) SetSettleType(v int) *FoodorderProductSaveRequestProductSettleInfo {
	s.SettleType = &v
	return s
}

type FoodorderProductSaveRequestProductSkuListItem struct {
	OutSkuId     *string                                                          `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price        *FoodorderProductSaveRequestProductSkuListItemPrice              `json:"price,omitempty" xml:"price,omitempty"`
	SaleAttrList []*FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem `json:"sale_attr_list,omitempty" xml:"sale_attr_list,omitempty" type:"Repeated"`
}

func (s FoodorderProductSaveRequestProductSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSkuListItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSkuListItem) SetOutSkuId(v string) *FoodorderProductSaveRequestProductSkuListItem {
	s.OutSkuId = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSkuListItem) SetPrice(v *FoodorderProductSaveRequestProductSkuListItemPrice) *FoodorderProductSaveRequestProductSkuListItem {
	s.Price = v
	return s
}

func (s *FoodorderProductSaveRequestProductSkuListItem) SetSaleAttrList(v []*FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem) *FoodorderProductSaveRequestProductSkuListItem {
	s.SaleAttrList = v
	return s
}

type FoodorderProductSaveRequestProductSkuListItemPrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s FoodorderProductSaveRequestProductSkuListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSkuListItemPrice) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSkuListItemPrice) SetActualAmount(v int64) *FoodorderProductSaveRequestProductSkuListItemPrice {
	s.ActualAmount = &v
	return s
}

type FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem struct {
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
}

func (s FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem) SetGroupCode(v string) *FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem {
	s.GroupCode = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem) SetItemKey(v string) *FoodorderProductSaveRequestProductSkuListItemSaleAttrListItem {
	s.ItemKey = &v
	return s
}

type FoodorderProductSaveRequestProductSoldTime struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s FoodorderProductSaveRequestProductSoldTime) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductSoldTime) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductSoldTime) SetEndTime(v int64) *FoodorderProductSaveRequestProductSoldTime {
	s.EndTime = &v
	return s
}

func (s *FoodorderProductSaveRequestProductSoldTime) SetStartTime(v int64) *FoodorderProductSaveRequestProductSoldTime {
	s.StartTime = &v
	return s
}

type FoodorderProductSaveRequestProductTestInfo struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	UidList  []*string `json:"uid_list,omitempty" xml:"uid_list,omitempty" type:"Repeated"`
}

func (s FoodorderProductSaveRequestProductTestInfo) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveRequestProductTestInfo) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveRequestProductTestInfo) SetTestFlag(v bool) *FoodorderProductSaveRequestProductTestInfo {
	s.TestFlag = &v
	return s
}

func (s *FoodorderProductSaveRequestProductTestInfo) SetUidList(v []*string) *FoodorderProductSaveRequestProductTestInfo {
	s.UidList = v
	return s
}

type FoodorderProductSaveResponse struct {
	BaseResp *FoodorderProductSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *FoodorderProductSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *FoodorderProductSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s FoodorderProductSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveResponse) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveResponse) SetBaseResp(v *FoodorderProductSaveResponseBaseResp) *FoodorderProductSaveResponse {
	s.BaseResp = v
	return s
}

func (s *FoodorderProductSaveResponse) SetData(v *FoodorderProductSaveResponseData) *FoodorderProductSaveResponse {
	s.Data = v
	return s
}

func (s *FoodorderProductSaveResponse) SetExtra(v *FoodorderProductSaveResponseExtra) *FoodorderProductSaveResponse {
	s.Extra = v
	return s
}

type FoodorderProductSaveResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s FoodorderProductSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveResponseBaseResp) SetStatusMessage(v string) *FoodorderProductSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *FoodorderProductSaveResponseBaseResp) SetExtra(v map[string]*string) *FoodorderProductSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *FoodorderProductSaveResponseBaseResp) SetStatusCode(v int32) *FoodorderProductSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

type FoodorderProductSaveResponseData struct {
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FoodorderProductSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveResponseData) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveResponseData) SetProductId(v int64) *FoodorderProductSaveResponseData {
	s.ProductId = &v
	return s
}

func (s *FoodorderProductSaveResponseData) SetDescription(v string) *FoodorderProductSaveResponseData {
	s.Description = &v
	return s
}

func (s *FoodorderProductSaveResponseData) SetErrorCode(v int32) *FoodorderProductSaveResponseData {
	s.ErrorCode = &v
	return s
}

type FoodorderProductSaveResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s FoodorderProductSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FoodorderProductSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *FoodorderProductSaveResponseExtra) SetSubDescription(v string) *FoodorderProductSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FoodorderProductSaveResponseExtra) SetSubErrorCode(v int32) *FoodorderProductSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FoodorderProductSaveResponseExtra) SetDescription(v string) *FoodorderProductSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *FoodorderProductSaveResponseExtra) SetErrorCode(v int32) *FoodorderProductSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FoodorderProductSaveResponseExtra) SetLogid(v string) *FoodorderProductSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *FoodorderProductSaveResponseExtra) SetNow(v int64) *FoodorderProductSaveResponseExtra {
	s.Now = &v
	return s
}

type FulfillmentDeliveryVerifyRequest struct {
	DeliveryExtra  *FulfillmentDeliveryVerifyRequestDeliveryExtra      `json:"delivery_extra,omitempty" xml:"delivery_extra,omitempty"`
	EncryptedCodes []*string                                           `json:"encrypted_codes,omitempty" xml:"encrypted_codes,omitempty" type:"Repeated"`
	OrderId        *string                                             `json:"order_id,omitempty" xml:"order_id,omitempty"`
	PoiInfo        *string                                             `json:"poi_info,omitempty" xml:"poi_info,omitempty"`
	VerifyToken    *string                                             `json:"verify_token,omitempty" xml:"verify_token,omitempty"`
	Certificates   []*FulfillmentDeliveryVerifyRequestCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	Header         map[string]*string                                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string                                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfillmentDeliveryVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyRequest) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyRequest) SetDeliveryExtra(v *FulfillmentDeliveryVerifyRequestDeliveryExtra) *FulfillmentDeliveryVerifyRequest {
	s.DeliveryExtra = v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetEncryptedCodes(v []*string) *FulfillmentDeliveryVerifyRequest {
	s.EncryptedCodes = v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetOrderId(v string) *FulfillmentDeliveryVerifyRequest {
	s.OrderId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetPoiInfo(v string) *FulfillmentDeliveryVerifyRequest {
	s.PoiInfo = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetVerifyToken(v string) *FulfillmentDeliveryVerifyRequest {
	s.VerifyToken = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetCertificates(v []*FulfillmentDeliveryVerifyRequestCertificatesItem) *FulfillmentDeliveryVerifyRequest {
	s.Certificates = v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetHeader(v map[string]*string) *FulfillmentDeliveryVerifyRequest {
	s.Header = v
	return s
}

func (s *FulfillmentDeliveryVerifyRequest) SetAccessToken(v string) *FulfillmentDeliveryVerifyRequest {
	s.AccessToken = &v
	return s
}

type FulfillmentDeliveryVerifyRequestCertificatesItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	EncryptedCode *string `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty"`
	ItemOrderId   *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
}

func (s FulfillmentDeliveryVerifyRequestCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyRequestCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyRequestCertificatesItem) SetCertificateId(v string) *FulfillmentDeliveryVerifyRequestCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequestCertificatesItem) SetEncryptedCode(v string) *FulfillmentDeliveryVerifyRequestCertificatesItem {
	s.EncryptedCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequestCertificatesItem) SetItemOrderId(v string) *FulfillmentDeliveryVerifyRequestCertificatesItem {
	s.ItemOrderId = &v
	return s
}

type FulfillmentDeliveryVerifyRequestDeliveryExtra struct {
	ServerPhoneNum *string `json:"server_phone_num,omitempty" xml:"server_phone_num,omitempty"`
	OutShopId      *string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	OutShopName    *string `json:"out_shop_name,omitempty" xml:"out_shop_name,omitempty"`
	ServerName     *string `json:"server_name,omitempty" xml:"server_name,omitempty"`
}

func (s FulfillmentDeliveryVerifyRequestDeliveryExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyRequestDeliveryExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyRequestDeliveryExtra) SetServerPhoneNum(v string) *FulfillmentDeliveryVerifyRequestDeliveryExtra {
	s.ServerPhoneNum = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequestDeliveryExtra) SetOutShopId(v string) *FulfillmentDeliveryVerifyRequestDeliveryExtra {
	s.OutShopId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequestDeliveryExtra) SetOutShopName(v string) *FulfillmentDeliveryVerifyRequestDeliveryExtra {
	s.OutShopName = &v
	return s
}

func (s *FulfillmentDeliveryVerifyRequestDeliveryExtra) SetServerName(v string) *FulfillmentDeliveryVerifyRequestDeliveryExtra {
	s.ServerName = &v
	return s
}

type FulfillmentDeliveryVerifyResponse struct {
	Data  *FulfillmentDeliveryVerifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *FulfillmentDeliveryVerifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s FulfillmentDeliveryVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyResponse) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyResponse) SetData(v *FulfillmentDeliveryVerifyResponseData) *FulfillmentDeliveryVerifyResponse {
	s.Data = v
	return s
}

func (s *FulfillmentDeliveryVerifyResponse) SetExtra(v *FulfillmentDeliveryVerifyResponseExtra) *FulfillmentDeliveryVerifyResponse {
	s.Extra = v
	return s
}

type FulfillmentDeliveryVerifyResponseData struct {
	VerifyResults []*FulfillmentDeliveryVerifyResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfillmentDeliveryVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyResponseData) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyResponseData) SetVerifyResults(v []*FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) *FulfillmentDeliveryVerifyResponseData {
	s.VerifyResults = v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseData) SetGwErrorCode(v int32) *FulfillmentDeliveryVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseData) SetGwDescription(v string) *FulfillmentDeliveryVerifyResponseData {
	s.GwDescription = &v
	return s
}

type FulfillmentDeliveryVerifyResponseDataVerifyResultsItem struct {
	VerifyId        *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	VerifyTime      *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	CertificateCode *string `json:"certificate_code,omitempty" xml:"certificate_code,omitempty"`
	CertificateId   *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ItemOrderId     *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	ResultCode      *int32  `json:"result_code,omitempty" xml:"result_code,omitempty"`
	ResultMsg       *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetVerifyId(v string) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetVerifyTime(v int64) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.VerifyTime = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetCertificateCode(v string) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.CertificateCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetCertificateId(v string) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetItemOrderId(v string) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.ItemOrderId = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetResultCode(v int32) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.ResultCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem) SetResultMsg(v string) *FulfillmentDeliveryVerifyResponseDataVerifyResultsItem {
	s.ResultMsg = &v
	return s
}

type FulfillmentDeliveryVerifyResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s FulfillmentDeliveryVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentDeliveryVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetSubErrorCode(v int32) *FulfillmentDeliveryVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetDescription(v string) *FulfillmentDeliveryVerifyResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetErrorCode(v int32) *FulfillmentDeliveryVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetLogid(v string) *FulfillmentDeliveryVerifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetNow(v int64) *FulfillmentDeliveryVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfillmentDeliveryVerifyResponseExtra) SetSubDescription(v string) *FulfillmentDeliveryVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

type FulfillmentOrderCanUseRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	PoiId       *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s FulfillmentOrderCanUseRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseRequest) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseRequest) SetHeader(v map[string]*string) *FulfillmentOrderCanUseRequest {
	s.Header = v
	return s
}

func (s *FulfillmentOrderCanUseRequest) SetAccessToken(v string) *FulfillmentOrderCanUseRequest {
	s.AccessToken = &v
	return s
}

func (s *FulfillmentOrderCanUseRequest) SetOrderId(v string) *FulfillmentOrderCanUseRequest {
	s.OrderId = &v
	return s
}

func (s *FulfillmentOrderCanUseRequest) SetPoiId(v string) *FulfillmentOrderCanUseRequest {
	s.PoiId = &v
	return s
}

type FulfillmentOrderCanUseResponse struct {
	Extra *FulfillmentOrderCanUseResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *FulfillmentOrderCanUseResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FulfillmentOrderCanUseResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponse) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponse) SetExtra(v *FulfillmentOrderCanUseResponseExtra) *FulfillmentOrderCanUseResponse {
	s.Extra = v
	return s
}

func (s *FulfillmentOrderCanUseResponse) SetData(v *FulfillmentOrderCanUseResponseData) *FulfillmentOrderCanUseResponse {
	s.Data = v
	return s
}

type FulfillmentOrderCanUseResponseData struct {
	OrderId       *string                                               `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	CanUse        *bool                                                 `json:"can_use,omitempty" xml:"can_use,omitempty" require:"true"`
	Certificates  []*FulfillmentOrderCanUseResponseDataCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfillmentOrderCanUseResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseData) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseData) SetOrderId(v string) *FulfillmentOrderCanUseResponseData {
	s.OrderId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseData) SetCanUse(v bool) *FulfillmentOrderCanUseResponseData {
	s.CanUse = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseData) SetCertificates(v []*FulfillmentOrderCanUseResponseDataCertificatesItem) *FulfillmentOrderCanUseResponseData {
	s.Certificates = v
	return s
}

func (s *FulfillmentOrderCanUseResponseData) SetGwErrorCode(v int32) *FulfillmentOrderCanUseResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseData) SetGwDescription(v string) *FulfillmentOrderCanUseResponseData {
	s.GwDescription = &v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItem struct {
	StartTime             *int64                                                                   `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Code                  *string                                                                  `json:"code,omitempty" xml:"code,omitempty"`
	ExpireTime            *int64                                                                   `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	SkuInfo               *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo               `json:"sku_info,omitempty" xml:"sku_info,omitempty"`
	TimesCardInfo         *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo         `json:"times_card_info,omitempty" xml:"times_card_info,omitempty"`
	Amount                *FulfillmentOrderCanUseResponseDataCertificatesItemAmount                `json:"amount,omitempty" xml:"amount,omitempty"`
	CertificateId         *string                                                                  `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	OffPeakDiscountDetail *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail `json:"off_peak_discount_detail,omitempty" xml:"off_peak_discount_detail,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetStartTime(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetCode(v string) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.Code = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetExpireTime(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetSkuInfo(v *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.SkuInfo = v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetTimesCardInfo(v *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.TimesCardInfo = v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetAmount(v *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.Amount = v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetCertificateId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItem) SetOffPeakDiscountDetail(v *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) *FulfillmentOrderCanUseResponseDataCertificatesItem {
	s.OffPeakDiscountDetail = v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemAmount struct {
	OriginalAmount        *int32 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	PayAmount             *int32 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	PaymentDiscountAmount *int32 `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PlatformTicketAmount  *int32 `json:"platform_ticket_amount,omitempty" xml:"platform_ticket_amount,omitempty"`
	CouponPayAmount       *int32 `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	MerchantTicketAmount  *int32 `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemAmount) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemAmount) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetOriginalAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetPayAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.PayAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetPaymentDiscountAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetPlatformTicketAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.PlatformTicketAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetCouponPayAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemAmount) SetMerchantTicketAmount(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail struct {
	OffPeakTimeLimitType *int                                                                                `json:"off_peak_time_limit_type,omitempty" xml:"off_peak_time_limit_type,omitempty"`
	HasOffPeakDiscount   *bool                                                                               `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty"`
	OffPeakTime          *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime `json:"off_peak_time,omitempty" xml:"off_peak_time,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) SetOffPeakTimeLimitType(v int) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail {
	s.OffPeakTimeLimitType = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) SetHasOffPeakDiscount(v bool) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail) SetOffPeakTime(v *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetail {
	s.OffPeakTime = v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime struct {
	WeekDayList        []*int32                                                                                                    `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
	EndTime            *int64                                                                                                      `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                      `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) SetWeekDayList(v []*int32) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.WeekDayList = v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) SetDailyTimeRangeList(v []*FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.DailyTimeRangeList = v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) SetEndTime(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.EndTime = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime) SetStartTime(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.StartTime = &v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) SetEndTime(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) SetStartTime(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo struct {
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SpuId         *string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	GrouponType   *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	AccountId     *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	OutId         *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ThirdSkuId    *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetSkuId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.SkuId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetSpuId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.SpuId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetSoldStartTime(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.SoldStartTime = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetGrouponType(v int) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.GrouponType = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetAccountId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.AccountId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetMarketPrice(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.MarketPrice = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetOutId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.OutId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetThirdSkuId(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.ThirdSkuId = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo) SetTitle(v string) *FulfillmentOrderCanUseResponseDataCertificatesItemSkuInfo {
	s.Title = &v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo struct {
	TotalTimes          *int64                                                                                    `json:"total_times,omitempty" xml:"total_times,omitempty"`
	UsableTimes         *int64                                                                                    `json:"usable_times,omitempty" xml:"usable_times,omitempty"`
	LadderTimesCardInfo []*FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem `json:"ladder_times_card_info,omitempty" xml:"ladder_times_card_info,omitempty" type:"Repeated"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) SetTotalTimes(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo {
	s.TotalTimes = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) SetUsableTimes(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo {
	s.UsableTimes = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo) SetLadderTimesCardInfo(v []*FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfo {
	s.LadderTimesCardInfo = v
	return s
}

type FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem struct {
	Status         *int   `json:"status,omitempty" xml:"status,omitempty"`
	Ticket         *int64 `json:"ticket,omitempty" xml:"ticket,omitempty"`
	CrossedAmount  *int64 `json:"crossed_amount,omitempty" xml:"crossed_amount,omitempty"`
	MerchantTicket *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	OriginalAmount *int64 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	PayAmount      *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	Seq            *int32 `json:"seq,omitempty" xml:"seq,omitempty"`
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetStatus(v int) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Status = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetTicket(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Ticket = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetCrossedAmount(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.CrossedAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetMerchantTicket(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.MerchantTicket = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetOriginalAmount(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.OriginalAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetPayAmount(v int64) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.PayAmount = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetSeq(v int32) *FulfillmentOrderCanUseResponseDataCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Seq = &v
	return s
}

type FulfillmentOrderCanUseResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfillmentOrderCanUseResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentOrderCanUseResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentOrderCanUseResponseExtra) SetErrorCode(v int32) *FulfillmentOrderCanUseResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseExtra) SetLogid(v string) *FulfillmentOrderCanUseResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseExtra) SetNow(v int64) *FulfillmentOrderCanUseResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseExtra) SetSubDescription(v string) *FulfillmentOrderCanUseResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseExtra) SetSubErrorCode(v int32) *FulfillmentOrderCanUseResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfillmentOrderCanUseResponseExtra) SetDescription(v string) *FulfillmentOrderCanUseResponseExtra {
	s.Description = &v
	return s
}

type FulfillmentPushDeliveryRequest struct {
	DeliveryStatus *int                                               `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	BookId         *string                                            `json:"book_id,omitempty" xml:"book_id,omitempty"`
	DeliveryExtra  *FulfillmentPushDeliveryRequestDeliveryExtra       `json:"delivery_extra,omitempty" xml:"delivery_extra,omitempty"`
	UseAll         *bool                                              `json:"use_all,omitempty" xml:"use_all,omitempty"`
	OutOrderNo     *string                                            `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	Token          *string                                            `json:"token,omitempty" xml:"token,omitempty"`
	AccessToken    *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ItemBookList   []*FulfillmentPushDeliveryRequestItemBookListItem  `json:"item_book_list,omitempty" xml:"item_book_list,omitempty" type:"Repeated"`
	PoiInfo        *string                                            `json:"poi_info,omitempty" xml:"poi_info,omitempty"`
	ItemOrderList  []*FulfillmentPushDeliveryRequestItemOrderListItem `json:"item_order_list,omitempty" xml:"item_order_list,omitempty" type:"Repeated"`
	Header         map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
}

func (s FulfillmentPushDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryRequest) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryRequest) SetDeliveryStatus(v int) *FulfillmentPushDeliveryRequest {
	s.DeliveryStatus = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetBookId(v string) *FulfillmentPushDeliveryRequest {
	s.BookId = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetDeliveryExtra(v *FulfillmentPushDeliveryRequestDeliveryExtra) *FulfillmentPushDeliveryRequest {
	s.DeliveryExtra = v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetUseAll(v bool) *FulfillmentPushDeliveryRequest {
	s.UseAll = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetOutOrderNo(v string) *FulfillmentPushDeliveryRequest {
	s.OutOrderNo = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetToken(v string) *FulfillmentPushDeliveryRequest {
	s.Token = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetAccessToken(v string) *FulfillmentPushDeliveryRequest {
	s.AccessToken = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetItemBookList(v []*FulfillmentPushDeliveryRequestItemBookListItem) *FulfillmentPushDeliveryRequest {
	s.ItemBookList = v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetPoiInfo(v string) *FulfillmentPushDeliveryRequest {
	s.PoiInfo = &v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetItemOrderList(v []*FulfillmentPushDeliveryRequestItemOrderListItem) *FulfillmentPushDeliveryRequest {
	s.ItemOrderList = v
	return s
}

func (s *FulfillmentPushDeliveryRequest) SetHeader(v map[string]*string) *FulfillmentPushDeliveryRequest {
	s.Header = v
	return s
}

type FulfillmentPushDeliveryRequestDeliveryExtra struct {
	OutShopId      *string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	OutShopName    *string `json:"out_shop_name,omitempty" xml:"out_shop_name,omitempty"`
	ServerName     *string `json:"server_name,omitempty" xml:"server_name,omitempty"`
	ServerPhoneNum *string `json:"server_phone_num,omitempty" xml:"server_phone_num,omitempty"`
}

func (s FulfillmentPushDeliveryRequestDeliveryExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryRequestDeliveryExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryRequestDeliveryExtra) SetOutShopId(v string) *FulfillmentPushDeliveryRequestDeliveryExtra {
	s.OutShopId = &v
	return s
}

func (s *FulfillmentPushDeliveryRequestDeliveryExtra) SetOutShopName(v string) *FulfillmentPushDeliveryRequestDeliveryExtra {
	s.OutShopName = &v
	return s
}

func (s *FulfillmentPushDeliveryRequestDeliveryExtra) SetServerName(v string) *FulfillmentPushDeliveryRequestDeliveryExtra {
	s.ServerName = &v
	return s
}

func (s *FulfillmentPushDeliveryRequestDeliveryExtra) SetServerPhoneNum(v string) *FulfillmentPushDeliveryRequestDeliveryExtra {
	s.ServerPhoneNum = &v
	return s
}

type FulfillmentPushDeliveryRequestItemBookListItem struct {
	TimesNoList []*string `json:"times_no_list,omitempty" xml:"times_no_list,omitempty" type:"Repeated"`
	ItemBookId  *string   `json:"item_book_id,omitempty" xml:"item_book_id,omitempty"`
}

func (s FulfillmentPushDeliveryRequestItemBookListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryRequestItemBookListItem) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryRequestItemBookListItem) SetTimesNoList(v []*string) *FulfillmentPushDeliveryRequestItemBookListItem {
	s.TimesNoList = v
	return s
}

func (s *FulfillmentPushDeliveryRequestItemBookListItem) SetItemBookId(v string) *FulfillmentPushDeliveryRequestItemBookListItem {
	s.ItemBookId = &v
	return s
}

type FulfillmentPushDeliveryRequestItemOrderListItem struct {
	ExternalExtra []*FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem `json:"external_extra,omitempty" xml:"external_extra,omitempty" type:"Repeated"`
	ItemOrderId   *string                                                             `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	TimesNoList   []*string                                                           `json:"times_no_list,omitempty" xml:"times_no_list,omitempty" type:"Repeated"`
}

func (s FulfillmentPushDeliveryRequestItemOrderListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryRequestItemOrderListItem) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryRequestItemOrderListItem) SetExternalExtra(v []*FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem) *FulfillmentPushDeliveryRequestItemOrderListItem {
	s.ExternalExtra = v
	return s
}

func (s *FulfillmentPushDeliveryRequestItemOrderListItem) SetItemOrderId(v string) *FulfillmentPushDeliveryRequestItemOrderListItem {
	s.ItemOrderId = &v
	return s
}

func (s *FulfillmentPushDeliveryRequestItemOrderListItem) SetTimesNoList(v []*string) *FulfillmentPushDeliveryRequestItemOrderListItem {
	s.TimesNoList = v
	return s
}

type FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem struct {
	CertNo  *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem) SetCertNo(v string) *FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem {
	s.CertNo = &v
	return s
}

func (s *FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem) SetOrderId(v string) *FulfillmentPushDeliveryRequestItemOrderListItemExternalExtraItem {
	s.OrderId = &v
	return s
}

type FulfillmentPushDeliveryResponse struct {
	Data  *FulfillmentPushDeliveryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *FulfillmentPushDeliveryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s FulfillmentPushDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryResponse) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryResponse) SetData(v *FulfillmentPushDeliveryResponseData) *FulfillmentPushDeliveryResponse {
	s.Data = v
	return s
}

func (s *FulfillmentPushDeliveryResponse) SetExtra(v *FulfillmentPushDeliveryResponseExtra) *FulfillmentPushDeliveryResponse {
	s.Extra = v
	return s
}

type FulfillmentPushDeliveryResponseData struct {
	GwErrorCode   *int32                                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	DeliveryId    *string                                                 `json:"delivery_id,omitempty" xml:"delivery_id,omitempty"`
	VerifyResults []*FulfillmentPushDeliveryResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
}

func (s FulfillmentPushDeliveryResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryResponseData) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryResponseData) SetGwErrorCode(v int32) *FulfillmentPushDeliveryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseData) SetGwDescription(v string) *FulfillmentPushDeliveryResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseData) SetDeliveryId(v string) *FulfillmentPushDeliveryResponseData {
	s.DeliveryId = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseData) SetVerifyResults(v []*FulfillmentPushDeliveryResponseDataVerifyResultsItem) *FulfillmentPushDeliveryResponseData {
	s.VerifyResults = v
	return s
}

type FulfillmentPushDeliveryResponseDataVerifyResultsItem struct {
	VerifyTime    *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ItemBookId    *string `json:"item_book_id,omitempty" xml:"item_book_id,omitempty"`
	ItemOrderId   *string `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	VerifyId      *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
}

func (s FulfillmentPushDeliveryResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryResponseDataVerifyResultsItem) SetVerifyTime(v int64) *FulfillmentPushDeliveryResponseDataVerifyResultsItem {
	s.VerifyTime = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseDataVerifyResultsItem) SetCertificateId(v string) *FulfillmentPushDeliveryResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseDataVerifyResultsItem) SetItemBookId(v string) *FulfillmentPushDeliveryResponseDataVerifyResultsItem {
	s.ItemBookId = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseDataVerifyResultsItem) SetItemOrderId(v string) *FulfillmentPushDeliveryResponseDataVerifyResultsItem {
	s.ItemOrderId = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseDataVerifyResultsItem) SetVerifyId(v string) *FulfillmentPushDeliveryResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

type FulfillmentPushDeliveryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FulfillmentPushDeliveryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentPushDeliveryResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentPushDeliveryResponseExtra) SetLogid(v string) *FulfillmentPushDeliveryResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseExtra) SetNow(v int64) *FulfillmentPushDeliveryResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseExtra) SetSubDescription(v string) *FulfillmentPushDeliveryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseExtra) SetSubErrorCode(v int32) *FulfillmentPushDeliveryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseExtra) SetDescription(v string) *FulfillmentPushDeliveryResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfillmentPushDeliveryResponseExtra) SetErrorCode(v int32) *FulfillmentPushDeliveryResponseExtra {
	s.ErrorCode = &v
	return s
}

type FulfillmentQueryUserCertificatesRequest struct {
	Header        map[string]*string                                `json:"header,omitempty" xml:"header,omitempty"`
	OpenId        *string                                           `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Page          *int32                                            `json:"page,omitempty" xml:"page,omitempty"`
	PageDimension *int                                              `json:"page_dimension,omitempty" xml:"page_dimension,omitempty"`
	PageSize      *int32                                            `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PoiId         *string                                           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimeRange     *FulfillmentQueryUserCertificatesRequestTimeRange `json:"time_range,omitempty" xml:"time_range,omitempty"`
	BizType       *int32                                            `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	AccessToken   *string                                           `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId     *string                                           `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s FulfillmentQueryUserCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesRequest) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesRequest) SetHeader(v map[string]*string) *FulfillmentQueryUserCertificatesRequest {
	s.Header = v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetOpenId(v string) *FulfillmentQueryUserCertificatesRequest {
	s.OpenId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetPage(v int32) *FulfillmentQueryUserCertificatesRequest {
	s.Page = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetPageDimension(v int) *FulfillmentQueryUserCertificatesRequest {
	s.PageDimension = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetPageSize(v int32) *FulfillmentQueryUserCertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetPoiId(v string) *FulfillmentQueryUserCertificatesRequest {
	s.PoiId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetTimeRange(v *FulfillmentQueryUserCertificatesRequestTimeRange) *FulfillmentQueryUserCertificatesRequest {
	s.TimeRange = v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetBizType(v int32) *FulfillmentQueryUserCertificatesRequest {
	s.BizType = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetAccessToken(v string) *FulfillmentQueryUserCertificatesRequest {
	s.AccessToken = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequest) SetAccountId(v string) *FulfillmentQueryUserCertificatesRequest {
	s.AccountId = &v
	return s
}

type FulfillmentQueryUserCertificatesRequestTimeRange struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s FulfillmentQueryUserCertificatesRequestTimeRange) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesRequestTimeRange) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesRequestTimeRange) SetEndTime(v int64) *FulfillmentQueryUserCertificatesRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesRequestTimeRange) SetStartTime(v int64) *FulfillmentQueryUserCertificatesRequestTimeRange {
	s.StartTime = &v
	return s
}

type FulfillmentQueryUserCertificatesResponse struct {
	Data  *FulfillmentQueryUserCertificatesResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *FulfillmentQueryUserCertificatesResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s FulfillmentQueryUserCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponse) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponse) SetData(v *FulfillmentQueryUserCertificatesResponseData) *FulfillmentQueryUserCertificatesResponse {
	s.Data = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponse) SetExtra(v *FulfillmentQueryUserCertificatesResponseExtra) *FulfillmentQueryUserCertificatesResponse {
	s.Extra = v
	return s
}

type FulfillmentQueryUserCertificatesResponseData struct {
	GwDescription *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Orders        []*FulfillmentQueryUserCertificatesResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
	Total         *int64                                                    `json:"total,omitempty" xml:"total,omitempty"`
	GwErrorCode   *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FulfillmentQueryUserCertificatesResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseData) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseData) SetGwDescription(v string) *FulfillmentQueryUserCertificatesResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseData) SetOrders(v []*FulfillmentQueryUserCertificatesResponseDataOrdersItem) *FulfillmentQueryUserCertificatesResponseData {
	s.Orders = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseData) SetTotal(v int64) *FulfillmentQueryUserCertificatesResponseData {
	s.Total = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseData) SetGwErrorCode(v int32) *FulfillmentQueryUserCertificatesResponseData {
	s.GwErrorCode = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItem struct {
	Certificates []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	OrderId      *string                                                                   `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	CanUse       *bool                                                                     `json:"can_use,omitempty" xml:"can_use,omitempty" require:"true"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItem) SetCertificates(v []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) *FulfillmentQueryUserCertificatesResponseDataOrdersItem {
	s.Certificates = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItem) SetOrderId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItem {
	s.OrderId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItem) SetCanUse(v bool) *FulfillmentQueryUserCertificatesResponseDataOrdersItem {
	s.CanUse = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem struct {
	SkuInfo               *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo               `json:"sku_info,omitempty" xml:"sku_info,omitempty"`
	TimesCardInfo         *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo         `json:"times_card_info,omitempty" xml:"times_card_info,omitempty"`
	Amount                *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount                `json:"amount,omitempty" xml:"amount,omitempty"`
	CertificateId         *string                                                                                      `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	StartTime             *int64                                                                                       `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ExpireTime            *int64                                                                                       `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Code                  *string                                                                                      `json:"code,omitempty" xml:"code,omitempty"`
	OffPeakDiscountDetail *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail `json:"off_peak_discount_detail,omitempty" xml:"off_peak_discount_detail,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetSkuInfo(v *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.SkuInfo = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetTimesCardInfo(v *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.TimesCardInfo = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetAmount(v *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.Amount = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetCertificateId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetStartTime(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetExpireTime(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetCode(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.Code = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem) SetOffPeakDiscountDetail(v *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItem {
	s.OffPeakDiscountDetail = v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount struct {
	MerchantTicketAmount  *int32 `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginalAmount        *int32 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	PayAmount             *int32 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	PaymentDiscountAmount *int32 `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PlatformTicketAmount  *int32 `json:"platform_ticket_amount,omitempty" xml:"platform_ticket_amount,omitempty"`
	CouponPayAmount       *int32 `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetMerchantTicketAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetOriginalAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetPayAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.PayAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetPaymentDiscountAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetPlatformTicketAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.PlatformTicketAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount) SetCouponPayAmount(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemAmount {
	s.CouponPayAmount = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail struct {
	OffPeakTimeLimitType *int                                                                                                    `json:"off_peak_time_limit_type,omitempty" xml:"off_peak_time_limit_type,omitempty"`
	HasOffPeakDiscount   *bool                                                                                                   `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty"`
	OffPeakTime          *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime `json:"off_peak_time,omitempty" xml:"off_peak_time,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) SetOffPeakTimeLimitType(v int) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail {
	s.OffPeakTimeLimitType = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) SetHasOffPeakDiscount(v bool) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail) SetOffPeakTime(v *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetail {
	s.OffPeakTime = v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime struct {
	EndTime            *int64                                                                                                                          `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                                          `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int32                                                                                                                        `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) SetEndTime(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.EndTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) SetStartTime(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.StartTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) SetWeekDayList(v []*int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.WeekDayList = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime) SetDailyTimeRangeList(v []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTime {
	s.DailyTimeRangeList = v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem struct {
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) SetStartTime(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem) SetEndTime(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemOffPeakDiscountDetailOffPeakTimeDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo struct {
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	SpuId         *string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	ThirdSkuId    *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	GrouponType   *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OutId         *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	AccountId     *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetSoldStartTime(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.SoldStartTime = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetSpuId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.SpuId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetThirdSkuId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.ThirdSkuId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetTitle(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.Title = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetGrouponType(v int) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.GrouponType = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetSkuId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.SkuId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetOutId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.OutId = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetMarketPrice(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.MarketPrice = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo) SetAccountId(v string) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemSkuInfo {
	s.AccountId = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo struct {
	UsableTimes         *int64                                                                                                        `json:"usable_times,omitempty" xml:"usable_times,omitempty"`
	LadderTimesCardInfo []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem `json:"ladder_times_card_info,omitempty" xml:"ladder_times_card_info,omitempty" type:"Repeated"`
	TimesCardType       *int                                                                                                          `json:"times_card_type,omitempty" xml:"times_card_type,omitempty"`
	TotalTimes          *int64                                                                                                        `json:"total_times,omitempty" xml:"total_times,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) SetUsableTimes(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo {
	s.UsableTimes = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) SetLadderTimesCardInfo(v []*FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo {
	s.LadderTimesCardInfo = v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) SetTimesCardType(v int) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo {
	s.TimesCardType = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo) SetTotalTimes(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfo {
	s.TotalTimes = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem struct {
	MerchantTicket *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	OriginalAmount *int64 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	PayAmount      *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	Seq            *int32 `json:"seq,omitempty" xml:"seq,omitempty"`
	Status         *int   `json:"status,omitempty" xml:"status,omitempty"`
	Ticket         *int64 `json:"ticket,omitempty" xml:"ticket,omitempty"`
	CrossedAmount  *int64 `json:"crossed_amount,omitempty" xml:"crossed_amount,omitempty"`
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetMerchantTicket(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.MerchantTicket = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetOriginalAmount(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.OriginalAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetPayAmount(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.PayAmount = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetSeq(v int32) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Seq = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetStatus(v int) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Status = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetTicket(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.Ticket = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem) SetCrossedAmount(v int64) *FulfillmentQueryUserCertificatesResponseDataOrdersItemCertificatesItemTimesCardInfoLadderTimesCardInfoItem {
	s.CrossedAmount = &v
	return s
}

type FulfillmentQueryUserCertificatesResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s FulfillmentQueryUserCertificatesResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentQueryUserCertificatesResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetSubErrorCode(v int32) *FulfillmentQueryUserCertificatesResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetDescription(v string) *FulfillmentQueryUserCertificatesResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetErrorCode(v int32) *FulfillmentQueryUserCertificatesResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetLogid(v string) *FulfillmentQueryUserCertificatesResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetNow(v int64) *FulfillmentQueryUserCertificatesResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfillmentQueryUserCertificatesResponseExtra) SetSubDescription(v string) *FulfillmentQueryUserCertificatesResponseExtra {
	s.SubDescription = &v
	return s
}

type FulfillmentVerifyCancelRequest struct {
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyId      *string            `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	CertificateId *string            `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfillmentVerifyCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentVerifyCancelRequest) GoString() string {
	return s.String()
}

func (s *FulfillmentVerifyCancelRequest) SetOrderId(v string) *FulfillmentVerifyCancelRequest {
	s.OrderId = &v
	return s
}

func (s *FulfillmentVerifyCancelRequest) SetVerifyId(v string) *FulfillmentVerifyCancelRequest {
	s.VerifyId = &v
	return s
}

func (s *FulfillmentVerifyCancelRequest) SetCertificateId(v string) *FulfillmentVerifyCancelRequest {
	s.CertificateId = &v
	return s
}

func (s *FulfillmentVerifyCancelRequest) SetHeader(v map[string]*string) *FulfillmentVerifyCancelRequest {
	s.Header = v
	return s
}

func (s *FulfillmentVerifyCancelRequest) SetAccessToken(v string) *FulfillmentVerifyCancelRequest {
	s.AccessToken = &v
	return s
}

type FulfillmentVerifyCancelResponse struct {
	Extra *FulfillmentVerifyCancelResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *FulfillmentVerifyCancelResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FulfillmentVerifyCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentVerifyCancelResponse) GoString() string {
	return s.String()
}

func (s *FulfillmentVerifyCancelResponse) SetExtra(v *FulfillmentVerifyCancelResponseExtra) *FulfillmentVerifyCancelResponse {
	s.Extra = v
	return s
}

func (s *FulfillmentVerifyCancelResponse) SetData(v *FulfillmentVerifyCancelResponseData) *FulfillmentVerifyCancelResponse {
	s.Data = v
	return s
}

type FulfillmentVerifyCancelResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfillmentVerifyCancelResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentVerifyCancelResponseData) GoString() string {
	return s.String()
}

func (s *FulfillmentVerifyCancelResponseData) SetErrorCode(v int32) *FulfillmentVerifyCancelResponseData {
	s.ErrorCode = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseData) SetDescription(v string) *FulfillmentVerifyCancelResponseData {
	s.Description = &v
	return s
}

type FulfillmentVerifyCancelResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s FulfillmentVerifyCancelResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfillmentVerifyCancelResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfillmentVerifyCancelResponseExtra) SetSubDescription(v string) *FulfillmentVerifyCancelResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseExtra) SetSubErrorCode(v int32) *FulfillmentVerifyCancelResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseExtra) SetDescription(v string) *FulfillmentVerifyCancelResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseExtra) SetErrorCode(v int32) *FulfillmentVerifyCancelResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseExtra) SetLogid(v string) *FulfillmentVerifyCancelResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfillmentVerifyCancelResponseExtra) SetNow(v int64) *FulfillmentVerifyCancelResponseExtra {
	s.Now = &v
	return s
}

type FulfilmentCertificateCancelRequest struct {
	AccountId            *string                                                      `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BatchCancelInfoList  []*FulfilmentCertificateCancelRequestBatchCancelInfoListItem `json:"batch_cancel_info_list,omitempty" xml:"batch_cancel_info_list,omitempty" type:"Repeated"`
	ShopOrderId          *string                                                      `json:"shop_order_id,omitempty" xml:"shop_order_id,omitempty"`
	TimesCardCancelCount *int64                                                       `json:"times_card_cancel_count,omitempty" xml:"times_card_cancel_count,omitempty"`
	VerifyId             *string                                                      `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	BatchCancelInfo      *FulfilmentCertificateCancelRequestBatchCancelInfo           `json:"batch_cancel_info,omitempty" xml:"batch_cancel_info,omitempty"`
	CancelToken          *string                                                      `json:"cancel_token,omitempty" xml:"cancel_token,omitempty"`
	CertificateId        *string                                                      `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Header               map[string]*string                                           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string                                                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfilmentCertificateCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelRequest) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelRequest) SetAccountId(v string) *FulfilmentCertificateCancelRequest {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetBatchCancelInfoList(v []*FulfilmentCertificateCancelRequestBatchCancelInfoListItem) *FulfilmentCertificateCancelRequest {
	s.BatchCancelInfoList = v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetShopOrderId(v string) *FulfilmentCertificateCancelRequest {
	s.ShopOrderId = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetTimesCardCancelCount(v int64) *FulfilmentCertificateCancelRequest {
	s.TimesCardCancelCount = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetVerifyId(v string) *FulfilmentCertificateCancelRequest {
	s.VerifyId = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetBatchCancelInfo(v *FulfilmentCertificateCancelRequestBatchCancelInfo) *FulfilmentCertificateCancelRequest {
	s.BatchCancelInfo = v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetCancelToken(v string) *FulfilmentCertificateCancelRequest {
	s.CancelToken = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetCertificateId(v string) *FulfilmentCertificateCancelRequest {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetHeader(v map[string]*string) *FulfilmentCertificateCancelRequest {
	s.Header = v
	return s
}

func (s *FulfilmentCertificateCancelRequest) SetAccessToken(v string) *FulfilmentCertificateCancelRequest {
	s.AccessToken = &v
	return s
}

type FulfilmentCertificateCancelRequestBatchCancelInfo struct {
	VerifyIdList []*string `json:"verify_id_list,omitempty" xml:"verify_id_list,omitempty" type:"Repeated"`
	OrderId      *string   `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s FulfilmentCertificateCancelRequestBatchCancelInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelRequestBatchCancelInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelRequestBatchCancelInfo) SetVerifyIdList(v []*string) *FulfilmentCertificateCancelRequestBatchCancelInfo {
	s.VerifyIdList = v
	return s
}

func (s *FulfilmentCertificateCancelRequestBatchCancelInfo) SetOrderId(v string) *FulfilmentCertificateCancelRequestBatchCancelInfo {
	s.OrderId = &v
	return s
}

type FulfilmentCertificateCancelRequestBatchCancelInfoListItem struct {
	OrderId      *string   `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyIdList []*string `json:"verify_id_list,omitempty" xml:"verify_id_list,omitempty" type:"Repeated"`
}

func (s FulfilmentCertificateCancelRequestBatchCancelInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelRequestBatchCancelInfoListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelRequestBatchCancelInfoListItem) SetOrderId(v string) *FulfilmentCertificateCancelRequestBatchCancelInfoListItem {
	s.OrderId = &v
	return s
}

func (s *FulfilmentCertificateCancelRequestBatchCancelInfoListItem) SetVerifyIdList(v []*string) *FulfilmentCertificateCancelRequestBatchCancelInfoListItem {
	s.VerifyIdList = v
	return s
}

type FulfilmentCertificateCancelResponse struct {
	Data  *FulfilmentCertificateCancelResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *FulfilmentCertificateCancelResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FulfilmentCertificateCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelResponse) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelResponse) SetData(v *FulfilmentCertificateCancelResponseData) *FulfilmentCertificateCancelResponse {
	s.Data = v
	return s
}

func (s *FulfilmentCertificateCancelResponse) SetExtra(v *FulfilmentCertificateCancelResponseExtra) *FulfilmentCertificateCancelResponse {
	s.Extra = v
	return s
}

type FulfilmentCertificateCancelResponseData struct {
	GwErrorCode   *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CancelResults []*FulfilmentCertificateCancelResponseDataCancelResultsItem `json:"cancel_results,omitempty" xml:"cancel_results,omitempty" type:"Repeated"`
}

func (s FulfilmentCertificateCancelResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelResponseData) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelResponseData) SetGwErrorCode(v int32) *FulfilmentCertificateCancelResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseData) SetGwDescription(v string) *FulfilmentCertificateCancelResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseData) SetCancelResults(v []*FulfilmentCertificateCancelResponseDataCancelResultsItem) *FulfilmentCertificateCancelResponseData {
	s.CancelResults = v
	return s
}

type FulfilmentCertificateCancelResponseDataCancelResultsItem struct {
	ResultMsg  *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	VerifyId   *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ResultCode *int32  `json:"result_code,omitempty" xml:"result_code,omitempty" require:"true"`
}

func (s FulfilmentCertificateCancelResponseDataCancelResultsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelResponseDataCancelResultsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelResponseDataCancelResultsItem) SetResultMsg(v string) *FulfilmentCertificateCancelResponseDataCancelResultsItem {
	s.ResultMsg = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseDataCancelResultsItem) SetVerifyId(v string) *FulfilmentCertificateCancelResponseDataCancelResultsItem {
	s.VerifyId = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseDataCancelResultsItem) SetOrderId(v string) *FulfilmentCertificateCancelResponseDataCancelResultsItem {
	s.OrderId = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseDataCancelResultsItem) SetResultCode(v int32) *FulfilmentCertificateCancelResponseDataCancelResultsItem {
	s.ResultCode = &v
	return s
}

type FulfilmentCertificateCancelResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FulfilmentCertificateCancelResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateCancelResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateCancelResponseExtra) SetLogid(v string) *FulfilmentCertificateCancelResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseExtra) SetNow(v int64) *FulfilmentCertificateCancelResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseExtra) SetSubDescription(v string) *FulfilmentCertificateCancelResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseExtra) SetSubErrorCode(v int32) *FulfilmentCertificateCancelResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseExtra) SetDescription(v string) *FulfilmentCertificateCancelResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfilmentCertificateCancelResponseExtra) SetErrorCode(v int32) *FulfilmentCertificateCancelResponseExtra {
	s.ErrorCode = &v
	return s
}

type FulfilmentCertificatePrepareRequest struct {
	EncryptedData *string            `json:"encrypted_data,omitempty" xml:"encrypted_data,omitempty"`
	PoiId         *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Code          *string            `json:"code,omitempty" xml:"code,omitempty"`
}

func (s FulfilmentCertificatePrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareRequest) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareRequest) SetEncryptedData(v string) *FulfilmentCertificatePrepareRequest {
	s.EncryptedData = &v
	return s
}

func (s *FulfilmentCertificatePrepareRequest) SetPoiId(v string) *FulfilmentCertificatePrepareRequest {
	s.PoiId = &v
	return s
}

func (s *FulfilmentCertificatePrepareRequest) SetAccountId(v string) *FulfilmentCertificatePrepareRequest {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificatePrepareRequest) SetHeader(v map[string]*string) *FulfilmentCertificatePrepareRequest {
	s.Header = v
	return s
}

func (s *FulfilmentCertificatePrepareRequest) SetAccessToken(v string) *FulfilmentCertificatePrepareRequest {
	s.AccessToken = &v
	return s
}

func (s *FulfilmentCertificatePrepareRequest) SetCode(v string) *FulfilmentCertificatePrepareRequest {
	s.Code = &v
	return s
}

type FulfilmentCertificatePrepareResponse struct {
	Data  *FulfilmentCertificatePrepareResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *FulfilmentCertificatePrepareResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s FulfilmentCertificatePrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponse) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponse) SetData(v *FulfilmentCertificatePrepareResponseData) *FulfilmentCertificatePrepareResponse {
	s.Data = v
	return s
}

func (s *FulfilmentCertificatePrepareResponse) SetExtra(v *FulfilmentCertificatePrepareResponseExtra) *FulfilmentCertificatePrepareResponse {
	s.Extra = v
	return s
}

type FulfilmentCertificatePrepareResponseData struct {
	VerifyToken   *string                                                     `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	Certificates  []*FulfilmentCertificatePrepareResponseDataCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OrderId       *string                                                     `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s FulfilmentCertificatePrepareResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseData) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseData) SetVerifyToken(v string) *FulfilmentCertificatePrepareResponseData {
	s.VerifyToken = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseData) SetCertificates(v []*FulfilmentCertificatePrepareResponseDataCertificatesItem) *FulfilmentCertificatePrepareResponseData {
	s.Certificates = v
	return s
}

func (s *FulfilmentCertificatePrepareResponseData) SetGwErrorCode(v int32) *FulfilmentCertificatePrepareResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseData) SetGwDescription(v string) *FulfilmentCertificatePrepareResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseData) SetOrderId(v string) *FulfilmentCertificatePrepareResponseData {
	s.OrderId = &v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItem struct {
	CertificateId       *int64                                                               `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	ExpireTime          *int64                                                               `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	UseTimeInfo         *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	EncryptedCode       *string                                                              `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	Amount              *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount      `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	StartTime           *int64                                                               `json:"start_time,omitempty" xml:"start_time,omitempty"`
	NotAvailablePoiList []*string                                                            `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	BookInfo            *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo    `json:"book_info,omitempty" xml:"book_info,omitempty"`
	Sku                 *FulfilmentCertificatePrepareResponseDataCertificatesItemSku         `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetCertificateId(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetExpireTime(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetUseTimeInfo(v *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.UseTimeInfo = v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetEncryptedCode(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.EncryptedCode = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetAmount(v *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.Amount = v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetStartTime(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetNotAvailablePoiList(v []*string) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.NotAvailablePoiList = v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetBookInfo(v *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.BookInfo = v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItem) SetSku(v *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) *FulfilmentCertificatePrepareResponseDataCertificatesItem {
	s.Sku = v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItemAmount struct {
	PlatformDiscountAmount *int32 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	OriginalAmount         *int32 `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32 `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	ListMarketAmount       *int32 `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	MerchantTicketAmount   *int32 `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PayAmount              *int32 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetPlatformDiscountAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetOriginalAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetPaymentDiscountAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetListMarketAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetMerchantTicketAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount) SetPayAmount(v int32) *FulfilmentCertificatePrepareResponseDataCertificatesItemAmount {
	s.PayAmount = &v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo struct {
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
	VerifyAmount      *int64  `json:"verify_amount,omitempty" xml:"verify_amount,omitempty"`
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) SetBookProductNumber(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo {
	s.BookProductNumber = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) SetVerifyAmount(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo {
	s.VerifyAmount = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo) SetBookPoiId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemBookInfo {
	s.BookPoiId = &v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItemSku struct {
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemSku) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemSku) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetMarketPrice(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.MarketPrice = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetAccountId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetProductId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.ProductId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetSoldStartTime(v int64) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetTitle(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.Title = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetThirdSkuId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetSkuOutId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.SkuOutId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetVoucherType(v int) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.VoucherType = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetGrouponType(v int) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.GrouponType = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetSkuId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.SkuId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetProductOutId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.ProductOutId = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemSku) SetSuplierProductOutId(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemSku {
	s.SuplierProductOutId = &v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo struct {
	UseTimeType    *int                                                                                     `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
	TimePeriodList []*FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo) SetUseTimeType(v int) *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo) SetTimePeriodList(v []*FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

type FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem struct {
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *FulfilmentCertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

type FulfilmentCertificatePrepareResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s FulfilmentCertificatePrepareResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificatePrepareResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetLogid(v string) *FulfilmentCertificatePrepareResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetNow(v int64) *FulfilmentCertificatePrepareResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetSubDescription(v string) *FulfilmentCertificatePrepareResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetSubErrorCode(v int32) *FulfilmentCertificatePrepareResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetDescription(v string) *FulfilmentCertificatePrepareResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfilmentCertificatePrepareResponseExtra) SetErrorCode(v int32) *FulfilmentCertificatePrepareResponseExtra {
	s.ErrorCode = &v
	return s
}

type FulfilmentCertificateQueryRequest struct {
	EncryptedCode *string            `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty"`
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfilmentCertificateQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryRequest) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryRequest) SetEncryptedCode(v string) *FulfilmentCertificateQueryRequest {
	s.EncryptedCode = &v
	return s
}

func (s *FulfilmentCertificateQueryRequest) SetOrderId(v string) *FulfilmentCertificateQueryRequest {
	s.OrderId = &v
	return s
}

func (s *FulfilmentCertificateQueryRequest) SetHeader(v map[string]*string) *FulfilmentCertificateQueryRequest {
	s.Header = v
	return s
}

func (s *FulfilmentCertificateQueryRequest) SetAccessToken(v string) *FulfilmentCertificateQueryRequest {
	s.AccessToken = &v
	return s
}

type FulfilmentCertificateQueryResponse struct {
	Data  *FulfilmentCertificateQueryResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *FulfilmentCertificateQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FulfilmentCertificateQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponse) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponse) SetData(v *FulfilmentCertificateQueryResponseData) *FulfilmentCertificateQueryResponse {
	s.Data = v
	return s
}

func (s *FulfilmentCertificateQueryResponse) SetExtra(v *FulfilmentCertificateQueryResponseExtra) *FulfilmentCertificateQueryResponse {
	s.Extra = v
	return s
}

type FulfilmentCertificateQueryResponseData struct {
	GwErrorCode   *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Certificates  []*FulfilmentCertificateQueryResponseDataCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
}

func (s FulfilmentCertificateQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseData) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseData) SetGwErrorCode(v int32) *FulfilmentCertificateQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseData) SetGwDescription(v string) *FulfilmentCertificateQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseData) SetCertificates(v []*FulfilmentCertificateQueryResponseDataCertificatesItem) *FulfilmentCertificateQueryResponseData {
	s.Certificates = v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItem struct {
	ExpireTime    *int64                                                             `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Code          *string                                                            `json:"code,omitempty" xml:"code,omitempty"`
	EncryptedCode *string                                                            `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	Verify        *FulfilmentCertificateQueryResponseDataCertificatesItemVerify      `json:"verify,omitempty" xml:"verify,omitempty"`
	UseTimeInfo   *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	Status        *int                                                               `json:"status,omitempty" xml:"status,omitempty"`
	BookInfo      *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo    `json:"book_info,omitempty" xml:"book_info,omitempty"`
	StartTime     *int64                                                             `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Sku           *FulfilmentCertificateQueryResponseDataCertificatesItemSku         `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetExpireTime(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetCode(v string) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.Code = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetEncryptedCode(v string) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.EncryptedCode = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetVerify(v *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.Verify = v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetUseTimeInfo(v *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.UseTimeInfo = v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetStatus(v int) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.Status = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetBookInfo(v *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.BookInfo = v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetStartTime(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItem) SetSku(v *FulfilmentCertificateQueryResponseDataCertificatesItemSku) *FulfilmentCertificateQueryResponseDataCertificatesItem {
	s.Sku = v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo) SetBookPoiId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo) SetBookProductNumber(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItemBookInfo {
	s.BookProductNumber = &v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItemSku struct {
	SkuOutId      *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	AccountId     *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GrouponType   *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	ProductId     *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ThirdSkuId    *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	MarketPrice   *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SkuId         *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	VoucherType   *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	ProductOutId  *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	SoldStartTime *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemSku) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemSku) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetSkuOutId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.SkuOutId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetAccountId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetGrouponType(v int) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.GrouponType = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetProductId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.ProductId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetThirdSkuId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetMarketPrice(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.MarketPrice = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetSkuId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.SkuId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetTitle(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.Title = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetVoucherType(v int) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.VoucherType = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetProductOutId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.ProductOutId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemSku) SetSoldStartTime(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItemSku {
	s.SoldStartTime = &v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo struct {
	TimePeriodList []*FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
	UseTimeType    *int                                                                                   `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo) SetTimePeriodList(v []*FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo) SetUseTimeType(v int) *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem struct {
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

type FulfilmentCertificateQueryResponseDataCertificatesItemVerify struct {
	VerifyType       *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel        *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId    *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId            *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	VerifierUniqueId *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId         *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime       *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemVerify) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseDataCertificatesItemVerify) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetVerifyType(v int) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyType = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetCanCancel(v bool) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.CanCancel = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetCertificateId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetPoiId(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.PoiId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetVerifierUniqueId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetVerifyId(v string) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyId = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseDataCertificatesItemVerify) SetVerifyTime(v int64) *FulfilmentCertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyTime = &v
	return s
}

type FulfilmentCertificateQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s FulfilmentCertificateQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateQueryResponseExtra) SetSubErrorCode(v int32) *FulfilmentCertificateQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseExtra) SetDescription(v string) *FulfilmentCertificateQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseExtra) SetErrorCode(v int32) *FulfilmentCertificateQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseExtra) SetLogid(v string) *FulfilmentCertificateQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseExtra) SetNow(v int64) *FulfilmentCertificateQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfilmentCertificateQueryResponseExtra) SetSubDescription(v string) *FulfilmentCertificateQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type FulfilmentCertificateVerifyRequest struct {
	CodeWithTimeList []*FulfilmentCertificateVerifyRequestCodeWithTimeListItem `json:"code_with_time_list,omitempty" xml:"code_with_time_list,omitempty" type:"Repeated"`
	VerifyExtra      *FulfilmentCertificateVerifyRequestVerifyExtra            `json:"verify_extra,omitempty" xml:"verify_extra,omitempty"`
	Voucher          *FulfilmentCertificateVerifyRequestVoucher                `json:"voucher,omitempty" xml:"voucher,omitempty"`
	VerifySignList   []*string                                                 `json:"verify_sign_list,omitempty" xml:"verify_sign_list,omitempty" type:"Repeated"`
	AccountId        *string                                                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId            *string                                                   `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	EncryptedCodes   []*string                                                 `json:"encrypted_codes,omitempty" xml:"encrypted_codes,omitempty" type:"Repeated"`
	OrderId          *string                                                   `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Vouchers         []*FulfilmentCertificateVerifyRequestVouchersItem         `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	VerifyToken      *string                                                   `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	Codes            []*string                                                 `json:"codes,omitempty" xml:"codes,omitempty" type:"Repeated"`
	Header           map[string]*string                                        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string                                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfilmentCertificateVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequest) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequest) SetCodeWithTimeList(v []*FulfilmentCertificateVerifyRequestCodeWithTimeListItem) *FulfilmentCertificateVerifyRequest {
	s.CodeWithTimeList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetVerifyExtra(v *FulfilmentCertificateVerifyRequestVerifyExtra) *FulfilmentCertificateVerifyRequest {
	s.VerifyExtra = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetVoucher(v *FulfilmentCertificateVerifyRequestVoucher) *FulfilmentCertificateVerifyRequest {
	s.Voucher = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetVerifySignList(v []*string) *FulfilmentCertificateVerifyRequest {
	s.VerifySignList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetAccountId(v string) *FulfilmentCertificateVerifyRequest {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetPoiId(v string) *FulfilmentCertificateVerifyRequest {
	s.PoiId = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetEncryptedCodes(v []*string) *FulfilmentCertificateVerifyRequest {
	s.EncryptedCodes = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetOrderId(v string) *FulfilmentCertificateVerifyRequest {
	s.OrderId = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetVouchers(v []*FulfilmentCertificateVerifyRequestVouchersItem) *FulfilmentCertificateVerifyRequest {
	s.Vouchers = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetVerifyToken(v string) *FulfilmentCertificateVerifyRequest {
	s.VerifyToken = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetCodes(v []*string) *FulfilmentCertificateVerifyRequest {
	s.Codes = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetHeader(v map[string]*string) *FulfilmentCertificateVerifyRequest {
	s.Header = v
	return s
}

func (s *FulfilmentCertificateVerifyRequest) SetAccessToken(v string) *FulfilmentCertificateVerifyRequest {
	s.AccessToken = &v
	return s
}

type FulfilmentCertificateVerifyRequestCodeWithTimeListItem struct {
	OuterNumb  *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb `json:"outer_numb,omitempty" xml:"outer_numb,omitempty"`
	SerialNum  *int32                                                           `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	VerifyTime *int64                                                           `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	Code       *string                                                          `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s FulfilmentCertificateVerifyRequestCodeWithTimeListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestCodeWithTimeListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItem) SetOuterNumb(v *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb) *FulfilmentCertificateVerifyRequestCodeWithTimeListItem {
	s.OuterNumb = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItem) SetSerialNum(v int32) *FulfilmentCertificateVerifyRequestCodeWithTimeListItem {
	s.SerialNum = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItem) SetVerifyTime(v int64) *FulfilmentCertificateVerifyRequestCodeWithTimeListItem {
	s.VerifyTime = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItem) SetCode(v string) *FulfilmentCertificateVerifyRequestCodeWithTimeListItem {
	s.Code = &v
	return s
}

type FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb struct {
	CouponNumber *string `json:"coupon_number,omitempty" xml:"coupon_number,omitempty"`
	OrderNumber  *string `json:"order_number,omitempty" xml:"order_number,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb) SetCouponNumber(v string) *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb {
	s.CouponNumber = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb) SetOrderNumber(v string) *FulfilmentCertificateVerifyRequestCodeWithTimeListItemOuterNumb {
	s.OrderNumber = &v
	return s
}

type FulfilmentCertificateVerifyRequestVerifyExtra struct {
	OutGoodIds          []*string                                                         `json:"out_good_ids,omitempty" xml:"out_good_ids,omitempty" type:"Repeated"`
	TotalVerify         *bool                                                             `json:"total_verify,omitempty" xml:"total_verify,omitempty"`
	VerifyModel         *int                                                              `json:"verify_model,omitempty" xml:"verify_model,omitempty"`
	DynamicCouponInfo   *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo   `json:"dynamic_coupon_info,omitempty" xml:"dynamic_coupon_info,omitempty"`
	OfflineAddPriceInfo *FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo `json:"offline_add_price_info,omitempty" xml:"offline_add_price_info,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestVerifyExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestVerifyExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtra) SetOutGoodIds(v []*string) *FulfilmentCertificateVerifyRequestVerifyExtra {
	s.OutGoodIds = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtra) SetTotalVerify(v bool) *FulfilmentCertificateVerifyRequestVerifyExtra {
	s.TotalVerify = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtra) SetVerifyModel(v int) *FulfilmentCertificateVerifyRequestVerifyExtra {
	s.VerifyModel = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtra) SetDynamicCouponInfo(v *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo) *FulfilmentCertificateVerifyRequestVerifyExtra {
	s.DynamicCouponInfo = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtra) SetOfflineAddPriceInfo(v *FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) *FulfilmentCertificateVerifyRequestVerifyExtra {
	s.OfflineAddPriceInfo = v
	return s
}

type FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo struct {
	BizTime               *int64 `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
	ActualDeductionAmount *int64 `json:"actual_deduction_amount,omitempty" xml:"actual_deduction_amount,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo) SetBizTime(v int64) *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo {
	s.BizTime = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo) SetActualDeductionAmount(v int64) *FulfilmentCertificateVerifyRequestVerifyExtraDynamicCouponInfo {
	s.ActualDeductionAmount = &v
	return s
}

type FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo struct {
	IsOfflineAddPrice *bool `json:"is_offline_add_price,omitempty" xml:"is_offline_add_price,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) SetIsOfflineAddPrice(v bool) *FulfilmentCertificateVerifyRequestVerifyExtraOfflineAddPriceInfo {
	s.IsOfflineAddPrice = &v
	return s
}

type FulfilmentCertificateVerifyRequestVoucher struct {
	CertificateNoList []*string `json:"certificate_no_list,omitempty" xml:"certificate_no_list,omitempty" type:"Repeated"`
	IdCardList        []*string `json:"id_card_list,omitempty" xml:"id_card_list,omitempty" type:"Repeated"`
	ProjectId         *string   `json:"project_id,omitempty" xml:"project_id,omitempty"`
	QrcodeList        []*string `json:"qrcode_list,omitempty" xml:"qrcode_list,omitempty" type:"Repeated"`
	VerifyTime        *int64    `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestVoucher) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestVoucher) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestVoucher) SetCertificateNoList(v []*string) *FulfilmentCertificateVerifyRequestVoucher {
	s.CertificateNoList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVoucher) SetIdCardList(v []*string) *FulfilmentCertificateVerifyRequestVoucher {
	s.IdCardList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVoucher) SetProjectId(v string) *FulfilmentCertificateVerifyRequestVoucher {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVoucher) SetQrcodeList(v []*string) *FulfilmentCertificateVerifyRequestVoucher {
	s.QrcodeList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVoucher) SetVerifyTime(v int64) *FulfilmentCertificateVerifyRequestVoucher {
	s.VerifyTime = &v
	return s
}

type FulfilmentCertificateVerifyRequestVouchersItem struct {
	CertificateNoList []*string `json:"certificate_no_list,omitempty" xml:"certificate_no_list,omitempty" type:"Repeated"`
	IdCardList        []*string `json:"id_card_list,omitempty" xml:"id_card_list,omitempty" type:"Repeated"`
	ProjectId         *string   `json:"project_id,omitempty" xml:"project_id,omitempty"`
	QrcodeList        []*string `json:"qrcode_list,omitempty" xml:"qrcode_list,omitempty" type:"Repeated"`
	VerifyTime        *int64    `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
}

func (s FulfilmentCertificateVerifyRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyRequestVouchersItem) SetCertificateNoList(v []*string) *FulfilmentCertificateVerifyRequestVouchersItem {
	s.CertificateNoList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVouchersItem) SetIdCardList(v []*string) *FulfilmentCertificateVerifyRequestVouchersItem {
	s.IdCardList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVouchersItem) SetProjectId(v string) *FulfilmentCertificateVerifyRequestVouchersItem {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVouchersItem) SetQrcodeList(v []*string) *FulfilmentCertificateVerifyRequestVouchersItem {
	s.QrcodeList = v
	return s
}

func (s *FulfilmentCertificateVerifyRequestVouchersItem) SetVerifyTime(v int64) *FulfilmentCertificateVerifyRequestVouchersItem {
	s.VerifyTime = &v
	return s
}

type FulfilmentCertificateVerifyResponse struct {
	Extra *FulfilmentCertificateVerifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *FulfilmentCertificateVerifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FulfilmentCertificateVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyResponse) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyResponse) SetExtra(v *FulfilmentCertificateVerifyResponseExtra) *FulfilmentCertificateVerifyResponse {
	s.Extra = v
	return s
}

func (s *FulfilmentCertificateVerifyResponse) SetData(v *FulfilmentCertificateVerifyResponseData) *FulfilmentCertificateVerifyResponse {
	s.Data = v
	return s
}

type FulfilmentCertificateVerifyResponseData struct {
	GwErrorCode   *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	VerifyResults []*FulfilmentCertificateVerifyResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
}

func (s FulfilmentCertificateVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyResponseData) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyResponseData) SetGwErrorCode(v int32) *FulfilmentCertificateVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseData) SetGwDescription(v string) *FulfilmentCertificateVerifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseData) SetVerifyResults(v []*FulfilmentCertificateVerifyResponseDataVerifyResultsItem) *FulfilmentCertificateVerifyResponseData {
	s.VerifyResults = v
	return s
}

type FulfilmentCertificateVerifyResponseDataVerifyResultsItem struct {
	AccountId     *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Msg           *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	CertificateNo *string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	Result        *int32  `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	VerifyId      *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	OriginCode    *string `json:"origin_code,omitempty" xml:"origin_code,omitempty"`
	OutGoodId     *string `json:"out_good_id,omitempty" xml:"out_good_id,omitempty"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s FulfilmentCertificateVerifyResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetAccountId(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.AccountId = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetCode(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.Code = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetCertificateId(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetMsg(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.Msg = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetCertificateNo(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.CertificateNo = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetResult(v int32) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.Result = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetVerifyId(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetOriginCode(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.OriginCode = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetOutGoodId(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.OutGoodId = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseDataVerifyResultsItem) SetOrderId(v string) *FulfilmentCertificateVerifyResponseDataVerifyResultsItem {
	s.OrderId = &v
	return s
}

type FulfilmentCertificateVerifyResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfilmentCertificateVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCertificateVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetErrorCode(v int32) *FulfilmentCertificateVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetLogid(v string) *FulfilmentCertificateVerifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetNow(v int64) *FulfilmentCertificateVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetSubDescription(v string) *FulfilmentCertificateVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetSubErrorCode(v int32) *FulfilmentCertificateVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfilmentCertificateVerifyResponseExtra) SetDescription(v string) *FulfilmentCertificateVerifyResponseExtra {
	s.Description = &v
	return s
}

type FulfilmentCreateCallbackRequest struct {
	FailReason         *string                                                  `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	Header             map[string]*string                                       `json:"header,omitempty" xml:"header,omitempty"`
	FailReasonDesc     *string                                                  `json:"fail_reason_desc,omitempty" xml:"fail_reason_desc,omitempty"`
	Voucher            *FulfilmentCreateCallbackRequestVoucher                  `json:"voucher,omitempty" xml:"voucher,omitempty"`
	Vouchers           []*FulfilmentCreateCallbackRequestVouchersItem           `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	Result             *int64                                                   `json:"result,omitempty" xml:"result,omitempty"`
	Codes              []*string                                                `json:"codes,omitempty" xml:"codes,omitempty" type:"Repeated"`
	OrderId            *string                                                  `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ThirdOrderId       *string                                                  `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
	Certificates       []*FulfilmentCreateCallbackRequestCertificatesItem       `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	AdditionalInfoList []*FulfilmentCreateCallbackRequestAdditionalInfoListItem `json:"additional_info_list,omitempty" xml:"additional_info_list,omitempty" type:"Repeated"`
	AccessToken        *string                                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FulfilmentCreateCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequest) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequest) SetFailReason(v string) *FulfilmentCreateCallbackRequest {
	s.FailReason = &v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetHeader(v map[string]*string) *FulfilmentCreateCallbackRequest {
	s.Header = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetFailReasonDesc(v string) *FulfilmentCreateCallbackRequest {
	s.FailReasonDesc = &v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetVoucher(v *FulfilmentCreateCallbackRequestVoucher) *FulfilmentCreateCallbackRequest {
	s.Voucher = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetVouchers(v []*FulfilmentCreateCallbackRequestVouchersItem) *FulfilmentCreateCallbackRequest {
	s.Vouchers = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetResult(v int64) *FulfilmentCreateCallbackRequest {
	s.Result = &v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetCodes(v []*string) *FulfilmentCreateCallbackRequest {
	s.Codes = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetOrderId(v string) *FulfilmentCreateCallbackRequest {
	s.OrderId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetThirdOrderId(v string) *FulfilmentCreateCallbackRequest {
	s.ThirdOrderId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetCertificates(v []*FulfilmentCreateCallbackRequestCertificatesItem) *FulfilmentCreateCallbackRequest {
	s.Certificates = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetAdditionalInfoList(v []*FulfilmentCreateCallbackRequestAdditionalInfoListItem) *FulfilmentCreateCallbackRequest {
	s.AdditionalInfoList = v
	return s
}

func (s *FulfilmentCreateCallbackRequest) SetAccessToken(v string) *FulfilmentCreateCallbackRequest {
	s.AccessToken = &v
	return s
}

type FulfilmentCreateCallbackRequestAdditionalInfoListItem struct {
	AdditionalMap map[int32]*string `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	Key           *string           `json:"key,omitempty" xml:"key,omitempty"`
}

func (s FulfilmentCreateCallbackRequestAdditionalInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestAdditionalInfoListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestAdditionalInfoListItem) SetAdditionalMap(v map[int32]*string) *FulfilmentCreateCallbackRequestAdditionalInfoListItem {
	s.AdditionalMap = v
	return s
}

func (s *FulfilmentCreateCallbackRequestAdditionalInfoListItem) SetKey(v string) *FulfilmentCreateCallbackRequestAdditionalInfoListItem {
	s.Key = &v
	return s
}

type FulfilmentCreateCallbackRequestCertificatesItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s FulfilmentCreateCallbackRequestCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestCertificatesItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestCertificatesItem) SetCertificateId(v string) *FulfilmentCreateCallbackRequestCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestCertificatesItem) SetCode(v string) *FulfilmentCreateCallbackRequestCertificatesItem {
	s.Code = &v
	return s
}

type FulfilmentCreateCallbackRequestVoucher struct {
	Projects []*FulfilmentCreateCallbackRequestVoucherProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	Entrance *FulfilmentCreateCallbackRequestVoucherEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVoucher) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVoucher) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVoucher) SetProjects(v []*FulfilmentCreateCallbackRequestVoucherProjectsItem) *FulfilmentCreateCallbackRequestVoucher {
	s.Projects = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucher) SetEntrance(v *FulfilmentCreateCallbackRequestVoucherEntrance) *FulfilmentCreateCallbackRequestVoucher {
	s.Entrance = v
	return s
}

type FulfilmentCreateCallbackRequestVoucherEntrance struct {
	ProjectId      *string                                                          `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                        `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                        `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                        `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                        `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                        `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
}

func (s FulfilmentCreateCallbackRequestVoucherEntrance) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVoucherEntrance) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetProjectId(v string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetQrcodes(v []*string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.Qrcodes = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetUrls(v []*string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.Urls = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetCertificateNos(v []*string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.CertificateNos = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetCredentials(v []*FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.Credentials = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetGmcodeImgs(v []*string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.GmcodeImgs = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntrance) SetIdCards(v []*string) *FulfilmentCreateCallbackRequestVoucherEntrance {
	s.IdCards = v
	return s
}

type FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem) SetCredentialNo(v string) *FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem) SetCredentialType(v int64) *FulfilmentCreateCallbackRequestVoucherEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

type FulfilmentCreateCallbackRequestVoucherProjectsItem struct {
	CertificateNos []*string                                                            `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                            `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                            `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                              `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	ProjectId      *string                                                              `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                            `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                            `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s FulfilmentCreateCallbackRequestVoucherProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVoucherProjectsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetCertificateNos(v []*string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetCredentials(v []*FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.Credentials = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetGmcodeImgs(v []*string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.GmcodeImgs = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetIdCards(v []*string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.IdCards = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetName(v string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.Name = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetProjectId(v string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetQrcodes(v []*string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItem) SetUrls(v []*string) *FulfilmentCreateCallbackRequestVoucherProjectsItem {
	s.Urls = v
	return s
}

type FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem struct {
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem) SetCredentialType(v int64) *FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem) SetCredentialNo(v string) *FulfilmentCreateCallbackRequestVoucherProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

type FulfilmentCreateCallbackRequestVouchersItem struct {
	Projects []*FulfilmentCreateCallbackRequestVouchersItemProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	Entrance *FulfilmentCreateCallbackRequestVouchersItemEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVouchersItem) SetProjects(v []*FulfilmentCreateCallbackRequestVouchersItemProjectsItem) *FulfilmentCreateCallbackRequestVouchersItem {
	s.Projects = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItem) SetEntrance(v *FulfilmentCreateCallbackRequestVouchersItemEntrance) *FulfilmentCreateCallbackRequestVouchersItem {
	s.Entrance = v
	return s
}

type FulfilmentCreateCallbackRequestVouchersItemEntrance struct {
	GmcodeImgs     []*string                                                             `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                             `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	ProjectId      *string                                                               `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                             `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                             `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                             `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
}

func (s FulfilmentCreateCallbackRequestVouchersItemEntrance) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVouchersItemEntrance) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetGmcodeImgs(v []*string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.GmcodeImgs = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetIdCards(v []*string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.IdCards = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetProjectId(v string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetQrcodes(v []*string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.Qrcodes = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetUrls(v []*string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.Urls = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetCertificateNos(v []*string) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.CertificateNos = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntrance) SetCredentials(v []*FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem) *FulfilmentCreateCallbackRequestVouchersItemEntrance {
	s.Credentials = v
	return s
}

type FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem) SetCredentialNo(v string) *FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem) SetCredentialType(v int64) *FulfilmentCreateCallbackRequestVouchersItemEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

type FulfilmentCreateCallbackRequestVouchersItemProjectsItem struct {
	Name           *string                                                                   `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	ProjectId      *string                                                                   `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                                 `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                                 `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                                 `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                                 `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                                 `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
}

func (s FulfilmentCreateCallbackRequestVouchersItemProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVouchersItemProjectsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetName(v string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.Name = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetProjectId(v string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.ProjectId = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetQrcodes(v []*string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetUrls(v []*string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.Urls = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetCertificateNos(v []*string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetCredentials(v []*FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.Credentials = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetGmcodeImgs(v []*string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.GmcodeImgs = v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItem) SetIdCards(v []*string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItem {
	s.IdCards = v
	return s
}

type FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem struct {
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem) SetCredentialType(v int64) *FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem) SetCredentialNo(v string) *FulfilmentCreateCallbackRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

type FulfilmentCreateCallbackResponse struct {
	CertificateInfoList []*FulfilmentCreateCallbackResponseCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	Data                *FulfilmentCreateCallbackResponseData                      `json:"data,omitempty" xml:"data,omitempty"`
	Extra               *FulfilmentCreateCallbackResponseExtra                     `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s FulfilmentCreateCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackResponse) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackResponse) SetCertificateInfoList(v []*FulfilmentCreateCallbackResponseCertificateInfoListItem) *FulfilmentCreateCallbackResponse {
	s.CertificateInfoList = v
	return s
}

func (s *FulfilmentCreateCallbackResponse) SetData(v *FulfilmentCreateCallbackResponseData) *FulfilmentCreateCallbackResponse {
	s.Data = v
	return s
}

func (s *FulfilmentCreateCallbackResponse) SetExtra(v *FulfilmentCreateCallbackResponseExtra) *FulfilmentCreateCallbackResponse {
	s.Extra = v
	return s
}

type FulfilmentCreateCallbackResponseCertificateInfoListItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s FulfilmentCreateCallbackResponseCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackResponseCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackResponseCertificateInfoListItem) SetCertificateId(v string) *FulfilmentCreateCallbackResponseCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseCertificateInfoListItem) SetCode(v string) *FulfilmentCreateCallbackResponseCertificateInfoListItem {
	s.Code = &v
	return s
}

type FulfilmentCreateCallbackResponseData struct {
	CertificateInfoList []*FulfilmentCreateCallbackResponseDataCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	GwErrorCode         *int32                                                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription       *string                                                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfilmentCreateCallbackResponseData) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackResponseData) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackResponseData) SetCertificateInfoList(v []*FulfilmentCreateCallbackResponseDataCertificateInfoListItem) *FulfilmentCreateCallbackResponseData {
	s.CertificateInfoList = v
	return s
}

func (s *FulfilmentCreateCallbackResponseData) SetGwErrorCode(v int32) *FulfilmentCreateCallbackResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseData) SetGwDescription(v string) *FulfilmentCreateCallbackResponseData {
	s.GwDescription = &v
	return s
}

type FulfilmentCreateCallbackResponseDataCertificateInfoListItem struct {
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
}

func (s FulfilmentCreateCallbackResponseDataCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackResponseDataCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackResponseDataCertificateInfoListItem) SetCode(v string) *FulfilmentCreateCallbackResponseDataCertificateInfoListItem {
	s.Code = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseDataCertificateInfoListItem) SetCertificateId(v string) *FulfilmentCreateCallbackResponseDataCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

type FulfilmentCreateCallbackResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FulfilmentCreateCallbackResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s FulfilmentCreateCallbackResponseExtra) GoString() string {
	return s.String()
}

func (s *FulfilmentCreateCallbackResponseExtra) SetErrorCode(v int32) *FulfilmentCreateCallbackResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseExtra) SetLogid(v string) *FulfilmentCreateCallbackResponseExtra {
	s.Logid = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseExtra) SetNow(v int64) *FulfilmentCreateCallbackResponseExtra {
	s.Now = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseExtra) SetSubDescription(v string) *FulfilmentCreateCallbackResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseExtra) SetSubErrorCode(v int32) *FulfilmentCreateCallbackResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *FulfilmentCreateCallbackResponseExtra) SetDescription(v string) *FulfilmentCreateCallbackResponseExtra {
	s.Description = &v
	return s
}

type FunctionConfigQueryStatusRequest struct {
	FunctionId  []*string          `json:"function_id,omitempty" xml:"function_id,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s FunctionConfigQueryStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s FunctionConfigQueryStatusRequest) GoString() string {
	return s.String()
}

func (s *FunctionConfigQueryStatusRequest) SetFunctionId(v []*string) *FunctionConfigQueryStatusRequest {
	s.FunctionId = v
	return s
}

func (s *FunctionConfigQueryStatusRequest) SetHeader(v map[string]*string) *FunctionConfigQueryStatusRequest {
	s.Header = v
	return s
}

func (s *FunctionConfigQueryStatusRequest) SetAccessToken(v string) *FunctionConfigQueryStatusRequest {
	s.AccessToken = &v
	return s
}

type FunctionConfigQueryStatusResponse struct {
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *FunctionConfigQueryStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s FunctionConfigQueryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s FunctionConfigQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *FunctionConfigQueryStatusResponse) SetLogId(v string) *FunctionConfigQueryStatusResponse {
	s.LogId = &v
	return s
}

func (s *FunctionConfigQueryStatusResponse) SetData(v *FunctionConfigQueryStatusResponseData) *FunctionConfigQueryStatusResponse {
	s.Data = v
	return s
}

func (s *FunctionConfigQueryStatusResponse) SetErrNo(v int32) *FunctionConfigQueryStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *FunctionConfigQueryStatusResponse) SetErrMsg(v string) *FunctionConfigQueryStatusResponse {
	s.ErrMsg = &v
	return s
}

type FunctionConfigQueryStatusResponseData struct {
	FunctionConfigStatus map[string]*int32 `json:"function_config_status,omitempty" xml:"function_config_status,omitempty"`
}

func (s FunctionConfigQueryStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s FunctionConfigQueryStatusResponseData) GoString() string {
	return s.String()
}

func (s *FunctionConfigQueryStatusResponseData) SetFunctionConfigStatus(v map[string]*int32) *FunctionConfigQueryStatusResponseData {
	s.FunctionConfigStatus = v
	return s
}

type FundBillsRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	MerchantId  *string            `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	AccountType *string            `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	PaymentType *string            `json:"payment_type,omitempty" xml:"payment_type,omitempty" require:"true"`
	TradeType   *string            `json:"trade_type,omitempty" xml:"trade_type,omitempty" require:"true"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
}

func (s FundBillsRequest) String() string {
	return tea.Prettify(s)
}

func (s FundBillsRequest) GoString() string {
	return s.String()
}

func (s *FundBillsRequest) SetHeader(v map[string]*string) *FundBillsRequest {
	s.Header = v
	return s
}

func (s *FundBillsRequest) SetAccessToken(v string) *FundBillsRequest {
	s.AccessToken = &v
	return s
}

func (s *FundBillsRequest) SetAppId(v string) *FundBillsRequest {
	s.AppId = &v
	return s
}

func (s *FundBillsRequest) SetMerchantId(v string) *FundBillsRequest {
	s.MerchantId = &v
	return s
}

func (s *FundBillsRequest) SetAccountType(v string) *FundBillsRequest {
	s.AccountType = &v
	return s
}

func (s *FundBillsRequest) SetPaymentType(v string) *FundBillsRequest {
	s.PaymentType = &v
	return s
}

func (s *FundBillsRequest) SetTradeType(v string) *FundBillsRequest {
	s.TradeType = &v
	return s
}

func (s *FundBillsRequest) SetBillDate(v string) *FundBillsRequest {
	s.BillDate = &v
	return s
}

type FundBillsResponse struct {
	LogId  *string                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *FundBillsResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s FundBillsResponse) String() string {
	return tea.Prettify(s)
}

func (s FundBillsResponse) GoString() string {
	return s.String()
}

func (s *FundBillsResponse) SetLogId(v string) *FundBillsResponse {
	s.LogId = &v
	return s
}

func (s *FundBillsResponse) SetData(v *FundBillsResponseData) *FundBillsResponse {
	s.Data = v
	return s
}

func (s *FundBillsResponse) SetErrNo(v int32) *FundBillsResponse {
	s.ErrNo = &v
	return s
}

func (s *FundBillsResponse) SetErrMsg(v string) *FundBillsResponse {
	s.ErrMsg = &v
	return s
}

type FundBillsResponseData struct {
	FundBillList []*string `json:"fund_bill_list,omitempty" xml:"fund_bill_list,omitempty" require:"true" type:"Repeated"`
}

func (s FundBillsResponseData) String() string {
	return tea.Prettify(s)
}

func (s FundBillsResponseData) GoString() string {
	return s.String()
}

func (s *FundBillsResponseData) SetFundBillList(v []*string) *FundBillsResponseData {
	s.FundBillList = v
	return s
}

type GameConsoleRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GameConsoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleRequest) GoString() string {
	return s.String()
}

func (s *GameConsoleRequest) SetHeader(v map[string]*string) *GameConsoleRequest {
	s.Header = v
	return s
}

func (s *GameConsoleRequest) SetAccessToken(v string) *GameConsoleRequest {
	s.AccessToken = &v
	return s
}

type GameConsoleResponse struct {
	Extra *GameConsoleResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *GameConsoleResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GameConsoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleResponse) GoString() string {
	return s.String()
}

func (s *GameConsoleResponse) SetExtra(v *GameConsoleResponseExtra) *GameConsoleResponse {
	s.Extra = v
	return s
}

func (s *GameConsoleResponse) SetData(v *GameConsoleResponseData) *GameConsoleResponse {
	s.Data = v
	return s
}

type GameConsoleResponseData struct {
	List          []*GameConsoleResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GameConsoleResponseData) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleResponseData) GoString() string {
	return s.String()
}

func (s *GameConsoleResponseData) SetList(v []*GameConsoleResponseDataListItem) *GameConsoleResponseData {
	s.List = v
	return s
}

func (s *GameConsoleResponseData) SetGwErrorCode(v int32) *GameConsoleResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *GameConsoleResponseData) SetGwDescription(v string) *GameConsoleResponseData {
	s.GwDescription = &v
	return s
}

type GameConsoleResponseDataListItem struct {
	Rank             *int32                                          `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                         `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                         `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                         `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                          `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                          `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                        `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*GameConsoleResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
}

func (s GameConsoleResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleResponseDataListItem) GoString() string {
	return s.String()
}

func (s *GameConsoleResponseDataListItem) SetRank(v int32) *GameConsoleResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetRankChange(v string) *GameConsoleResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetNickname(v string) *GameConsoleResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetAvatar(v string) *GameConsoleResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetFollowerCount(v int64) *GameConsoleResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetOnbillbaordTimes(v int32) *GameConsoleResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetEffectValue(v float64) *GameConsoleResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *GameConsoleResponseDataListItem) SetVideoList(v []*GameConsoleResponseDataListItemVideoListItem) *GameConsoleResponseDataListItem {
	s.VideoList = v
	return s
}

type GameConsoleResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s GameConsoleResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *GameConsoleResponseDataListItemVideoListItem) SetShareUrl(v string) *GameConsoleResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *GameConsoleResponseDataListItemVideoListItem) SetTitle(v string) *GameConsoleResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *GameConsoleResponseDataListItemVideoListItem) SetItemCover(v string) *GameConsoleResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type GameConsoleResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s GameConsoleResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GameConsoleResponseExtra) GoString() string {
	return s.String()
}

func (s *GameConsoleResponseExtra) SetDescription(v string) *GameConsoleResponseExtra {
	s.Description = &v
	return s
}

func (s *GameConsoleResponseExtra) SetSubErrorCode(v int32) *GameConsoleResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GameConsoleResponseExtra) SetSubDescription(v string) *GameConsoleResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GameConsoleResponseExtra) SetLogid(v string) *GameConsoleResponseExtra {
	s.Logid = &v
	return s
}

func (s *GameConsoleResponseExtra) SetNow(v int64) *GameConsoleResponseExtra {
	s.Now = &v
	return s
}

func (s *GameConsoleResponseExtra) SetErrorCode(v int32) *GameConsoleResponseExtra {
	s.ErrorCode = &v
	return s
}

type GameInfRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GameInfRequest) String() string {
	return tea.Prettify(s)
}

func (s GameInfRequest) GoString() string {
	return s.String()
}

func (s *GameInfRequest) SetHeader(v map[string]*string) *GameInfRequest {
	s.Header = v
	return s
}

func (s *GameInfRequest) SetAccessToken(v string) *GameInfRequest {
	s.AccessToken = &v
	return s
}

type GameInfResponse struct {
	Extra *GameInfResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *GameInfResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GameInfResponse) String() string {
	return tea.Prettify(s)
}

func (s GameInfResponse) GoString() string {
	return s.String()
}

func (s *GameInfResponse) SetExtra(v *GameInfResponseExtra) *GameInfResponse {
	s.Extra = v
	return s
}

func (s *GameInfResponse) SetData(v *GameInfResponseData) *GameInfResponse {
	s.Data = v
	return s
}

type GameInfResponseData struct {
	List          []*GameInfResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GameInfResponseData) String() string {
	return tea.Prettify(s)
}

func (s GameInfResponseData) GoString() string {
	return s.String()
}

func (s *GameInfResponseData) SetList(v []*GameInfResponseDataListItem) *GameInfResponseData {
	s.List = v
	return s
}

func (s *GameInfResponseData) SetGwErrorCode(v int32) *GameInfResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *GameInfResponseData) SetGwDescription(v string) *GameInfResponseData {
	s.GwDescription = &v
	return s
}

type GameInfResponseDataListItem struct {
	VideoList        []*GameInfResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                      `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                     `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                     `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                     `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                      `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                      `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                    `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
}

func (s GameInfResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s GameInfResponseDataListItem) GoString() string {
	return s.String()
}

func (s *GameInfResponseDataListItem) SetVideoList(v []*GameInfResponseDataListItemVideoListItem) *GameInfResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *GameInfResponseDataListItem) SetRank(v int32) *GameInfResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *GameInfResponseDataListItem) SetRankChange(v string) *GameInfResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *GameInfResponseDataListItem) SetNickname(v string) *GameInfResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *GameInfResponseDataListItem) SetAvatar(v string) *GameInfResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *GameInfResponseDataListItem) SetFollowerCount(v int64) *GameInfResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *GameInfResponseDataListItem) SetOnbillbaordTimes(v int32) *GameInfResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *GameInfResponseDataListItem) SetEffectValue(v float64) *GameInfResponseDataListItem {
	s.EffectValue = &v
	return s
}

type GameInfResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s GameInfResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s GameInfResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *GameInfResponseDataListItemVideoListItem) SetTitle(v string) *GameInfResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *GameInfResponseDataListItemVideoListItem) SetItemCover(v string) *GameInfResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *GameInfResponseDataListItemVideoListItem) SetShareUrl(v string) *GameInfResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type GameInfResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s GameInfResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GameInfResponseExtra) GoString() string {
	return s.String()
}

func (s *GameInfResponseExtra) SetNow(v int64) *GameInfResponseExtra {
	s.Now = &v
	return s
}

func (s *GameInfResponseExtra) SetErrorCode(v int32) *GameInfResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GameInfResponseExtra) SetDescription(v string) *GameInfResponseExtra {
	s.Description = &v
	return s
}

func (s *GameInfResponseExtra) SetSubErrorCode(v int32) *GameInfResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GameInfResponseExtra) SetSubDescription(v string) *GameInfResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GameInfResponseExtra) SetLogid(v string) *GameInfResponseExtra {
	s.Logid = &v
	return s
}

type GetAgencyUserBindRecordRequest struct {
	PageSize        *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	AgentId         *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	DouyinId        *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty"`
	AgencyTalentUid *string            `json:"agency_talent_uid,omitempty" xml:"agency_talent_uid,omitempty"`
	PageNo          *int32             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetAgencyUserBindRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgencyUserBindRecordRequest) GoString() string {
	return s.String()
}

func (s *GetAgencyUserBindRecordRequest) SetPageSize(v int32) *GetAgencyUserBindRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetAgentId(v int64) *GetAgencyUserBindRecordRequest {
	s.AgentId = &v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetDouyinId(v string) *GetAgencyUserBindRecordRequest {
	s.DouyinId = &v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetAgencyTalentUid(v string) *GetAgencyUserBindRecordRequest {
	s.AgencyTalentUid = &v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetPageNo(v int32) *GetAgencyUserBindRecordRequest {
	s.PageNo = &v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetHeader(v map[string]*string) *GetAgencyUserBindRecordRequest {
	s.Header = v
	return s
}

func (s *GetAgencyUserBindRecordRequest) SetAccessToken(v string) *GetAgencyUserBindRecordRequest {
	s.AccessToken = &v
	return s
}

type GetAgencyUserBindRecordResponse struct {
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetAgencyUserBindRecordResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetAgencyUserBindRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgencyUserBindRecordResponse) GoString() string {
	return s.String()
}

func (s *GetAgencyUserBindRecordResponse) SetErrNo(v int32) *GetAgencyUserBindRecordResponse {
	s.ErrNo = &v
	return s
}

func (s *GetAgencyUserBindRecordResponse) SetErrMsg(v string) *GetAgencyUserBindRecordResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetAgencyUserBindRecordResponse) SetLogId(v string) *GetAgencyUserBindRecordResponse {
	s.LogId = &v
	return s
}

func (s *GetAgencyUserBindRecordResponse) SetData(v *GetAgencyUserBindRecordResponseData) *GetAgencyUserBindRecordResponse {
	s.Data = v
	return s
}

type GetAgencyUserBindRecordResponseData struct {
	TotalCount *int64                                            `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	Results    []*GetAgencyUserBindRecordResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s GetAgencyUserBindRecordResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAgencyUserBindRecordResponseData) GoString() string {
	return s.String()
}

func (s *GetAgencyUserBindRecordResponseData) SetTotalCount(v int64) *GetAgencyUserBindRecordResponseData {
	s.TotalCount = &v
	return s
}

func (s *GetAgencyUserBindRecordResponseData) SetResults(v []*GetAgencyUserBindRecordResponseDataResultsItem) *GetAgencyUserBindRecordResponseData {
	s.Results = v
	return s
}

type GetAgencyUserBindRecordResponseDataResultsItem struct {
	BindTime        *int64  `json:"bind_time,omitempty" xml:"bind_time,omitempty" require:"true"`
	UnbindTime      *int64  `json:"unbind_time,omitempty" xml:"unbind_time,omitempty"`
	AgencyTalentUid *string `json:"agency_talent_uid,omitempty" xml:"agency_talent_uid,omitempty"`
	AgentId         *int64  `json:"agent_id,omitempty" xml:"agent_id,omitempty" require:"true"`
	DouyinId        *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
}

func (s GetAgencyUserBindRecordResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s GetAgencyUserBindRecordResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *GetAgencyUserBindRecordResponseDataResultsItem) SetBindTime(v int64) *GetAgencyUserBindRecordResponseDataResultsItem {
	s.BindTime = &v
	return s
}

func (s *GetAgencyUserBindRecordResponseDataResultsItem) SetUnbindTime(v int64) *GetAgencyUserBindRecordResponseDataResultsItem {
	s.UnbindTime = &v
	return s
}

func (s *GetAgencyUserBindRecordResponseDataResultsItem) SetAgencyTalentUid(v string) *GetAgencyUserBindRecordResponseDataResultsItem {
	s.AgencyTalentUid = &v
	return s
}

func (s *GetAgencyUserBindRecordResponseDataResultsItem) SetAgentId(v int64) *GetAgencyUserBindRecordResponseDataResultsItem {
	s.AgentId = &v
	return s
}

func (s *GetAgencyUserBindRecordResponseDataResultsItem) SetDouyinId(v string) *GetAgencyUserBindRecordResponseDataResultsItem {
	s.DouyinId = &v
	return s
}

type GetAllCardsForUserRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetAllCardsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllCardsForUserRequest) GoString() string {
	return s.String()
}

func (s *GetAllCardsForUserRequest) SetHeader(v map[string]*string) *GetAllCardsForUserRequest {
	s.Header = v
	return s
}

func (s *GetAllCardsForUserRequest) SetAccessToken(v string) *GetAllCardsForUserRequest {
	s.AccessToken = &v
	return s
}

type GetAllCardsForUserResponse struct {
	Errmsg   *string                                   `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	CardList []*GetAllCardsForUserResponseCardListItem `json:"card_list,omitempty" xml:"card_list,omitempty" require:"true" type:"Repeated"`
	Errcode  *int32                                    `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s GetAllCardsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllCardsForUserResponse) GoString() string {
	return s.String()
}

func (s *GetAllCardsForUserResponse) SetErrmsg(v string) *GetAllCardsForUserResponse {
	s.Errmsg = &v
	return s
}

func (s *GetAllCardsForUserResponse) SetCardList(v []*GetAllCardsForUserResponseCardListItem) *GetAllCardsForUserResponse {
	s.CardList = v
	return s
}

func (s *GetAllCardsForUserResponse) SetErrcode(v int32) *GetAllCardsForUserResponse {
	s.Errcode = &v
	return s
}

type GetAllCardsForUserResponseCardListItem struct {
	UserId       *int64  `json:"user_id,omitempty" xml:"user_id,omitempty" require:"true"`
	CardType     *int    `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Url          *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AuditOpinion *string `json:"audit_opinion,omitempty" xml:"audit_opinion,omitempty"`
}

func (s GetAllCardsForUserResponseCardListItem) String() string {
	return tea.Prettify(s)
}

func (s GetAllCardsForUserResponseCardListItem) GoString() string {
	return s.String()
}

func (s *GetAllCardsForUserResponseCardListItem) SetUserId(v int64) *GetAllCardsForUserResponseCardListItem {
	s.UserId = &v
	return s
}

func (s *GetAllCardsForUserResponseCardListItem) SetCardType(v int) *GetAllCardsForUserResponseCardListItem {
	s.CardType = &v
	return s
}

func (s *GetAllCardsForUserResponseCardListItem) SetValue(v string) *GetAllCardsForUserResponseCardListItem {
	s.Value = &v
	return s
}

func (s *GetAllCardsForUserResponseCardListItem) SetUrl(v string) *GetAllCardsForUserResponseCardListItem {
	s.Url = &v
	return s
}

func (s *GetAllCardsForUserResponseCardListItem) SetStatus(v int) *GetAllCardsForUserResponseCardListItem {
	s.Status = &v
	return s
}

func (s *GetAllCardsForUserResponseCardListItem) SetAuditOpinion(v string) *GetAllCardsForUserResponseCardListItem {
	s.AuditOpinion = &v
	return s
}

type GetAwemeBindTemplateInfoRequest struct {
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	TemplateId  []*int64           `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true" type:"Repeated"`
	AwemeId     *string            `json:"aweme_id,omitempty" xml:"aweme_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetAwemeBindTemplateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoRequest) SetType(v string) *GetAwemeBindTemplateInfoRequest {
	s.Type = &v
	return s
}

func (s *GetAwemeBindTemplateInfoRequest) SetTemplateId(v []*int64) *GetAwemeBindTemplateInfoRequest {
	s.TemplateId = v
	return s
}

func (s *GetAwemeBindTemplateInfoRequest) SetAwemeId(v string) *GetAwemeBindTemplateInfoRequest {
	s.AwemeId = &v
	return s
}

func (s *GetAwemeBindTemplateInfoRequest) SetHeader(v map[string]*string) *GetAwemeBindTemplateInfoRequest {
	s.Header = v
	return s
}

func (s *GetAwemeBindTemplateInfoRequest) SetAccessToken(v string) *GetAwemeBindTemplateInfoRequest {
	s.AccessToken = &v
	return s
}

type GetAwemeBindTemplateInfoResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetAwemeBindTemplateInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s GetAwemeBindTemplateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoResponse) SetErrMsg(v string) *GetAwemeBindTemplateInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponse) SetLogId(v string) *GetAwemeBindTemplateInfoResponse {
	s.LogId = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponse) SetData(v *GetAwemeBindTemplateInfoResponseData) *GetAwemeBindTemplateInfoResponse {
	s.Data = v
	return s
}

func (s *GetAwemeBindTemplateInfoResponse) SetErrNo(v int32) *GetAwemeBindTemplateInfoResponse {
	s.ErrNo = &v
	return s
}

type GetAwemeBindTemplateInfoResponseData struct {
	TemplateInfo map[int64]*GetAwemeBindTemplateInfoResponseDataTemplateInfoValue `json:"template_info,omitempty" xml:"template_info,omitempty"`
}

func (s GetAwemeBindTemplateInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoResponseData) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoResponseData) SetTemplateInfo(v map[int64]*GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) *GetAwemeBindTemplateInfoResponseData {
	s.TemplateInfo = v
	return s
}

type GetAwemeBindTemplateInfoResponseDataTemplateInfoValue struct {
	Title           *string                                                                     `json:"title,omitempty" xml:"title,omitempty"`
	Desc            *string                                                                     `json:"desc,omitempty" xml:"desc,omitempty"`
	TemplateContent []*GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem `json:"template_content,omitempty" xml:"template_content,omitempty" type:"Repeated"`
	Id              *int64                                                                      `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) SetTitle(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue {
	s.Title = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) SetDesc(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue {
	s.Desc = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) SetTemplateContent(v []*GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue {
	s.TemplateContent = v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue) SetId(v int64) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValue {
	s.Id = &v
	return s
}

type GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem struct {
	Name       *string                                                                                                 `json:"name,omitempty" xml:"name,omitempty"`
	ValType    *int                                                                                                    `json:"val_type,omitempty" xml:"val_type,omitempty"`
	MaterielId *string                                                                                                 `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
	Children   map[string][]*GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem `json:"children,omitempty" xml:"children,omitempty"`
	Desc       *string                                                                                                 `json:"desc,omitempty" xml:"desc,omitempty"`
	MainTitle  *string                                                                                                 `json:"mainTitle,omitempty" xml:"mainTitle,omitempty"`
	ValExample []*string                                                                                               `json:"val_example,omitempty" xml:"val_example,omitempty" type:"Repeated"`
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetName(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.Name = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetValType(v int) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.ValType = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetMaterielId(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.MaterielId = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetChildren(v map[string][]*GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.Children = v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetDesc(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.Desc = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetMainTitle(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.MainTitle = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem) SetValExample(v []*string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItem {
	s.ValExample = v
	return s
}

type GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem struct {
	ValType      *int      `json:"val_type,omitempty" xml:"val_type,omitempty"`
	ValExample   []*string `json:"val_example,omitempty" xml:"val_example,omitempty" type:"Repeated"`
	ValList      []*string `json:"val_list,omitempty" xml:"val_list,omitempty" type:"Repeated"`
	RejectReason *string   `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	MaterielId   *string   `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
	Name         *string   `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetValType(v int) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValType = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetValExample(v []*string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValExample = v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetValList(v []*string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValList = v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetRejectReason(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.RejectReason = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetMaterielId(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.MaterielId = &v
	return s
}

func (s *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem) SetName(v string) *GetAwemeBindTemplateInfoResponseDataTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.Name = &v
	return s
}

type GetAwemeBindTemplateListRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AwemeId      *string            `json:"aweme_id,omitempty" xml:"aweme_id,omitempty" require:"true"`
	Type         *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	CapacityList []*string          `json:"capacity_list,omitempty" xml:"capacity_list,omitempty" require:"true" type:"Repeated"`
}

func (s GetAwemeBindTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateListRequest) SetHeader(v map[string]*string) *GetAwemeBindTemplateListRequest {
	s.Header = v
	return s
}

func (s *GetAwemeBindTemplateListRequest) SetAccessToken(v string) *GetAwemeBindTemplateListRequest {
	s.AccessToken = &v
	return s
}

func (s *GetAwemeBindTemplateListRequest) SetAwemeId(v string) *GetAwemeBindTemplateListRequest {
	s.AwemeId = &v
	return s
}

func (s *GetAwemeBindTemplateListRequest) SetType(v string) *GetAwemeBindTemplateListRequest {
	s.Type = &v
	return s
}

func (s *GetAwemeBindTemplateListRequest) SetCapacityList(v []*string) *GetAwemeBindTemplateListRequest {
	s.CapacityList = v
	return s
}

type GetAwemeBindTemplateListResponse struct {
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetAwemeBindTemplateListResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetAwemeBindTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateListResponse) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateListResponse) SetErrNo(v int32) *GetAwemeBindTemplateListResponse {
	s.ErrNo = &v
	return s
}

func (s *GetAwemeBindTemplateListResponse) SetErrMsg(v string) *GetAwemeBindTemplateListResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetAwemeBindTemplateListResponse) SetLogId(v string) *GetAwemeBindTemplateListResponse {
	s.LogId = &v
	return s
}

func (s *GetAwemeBindTemplateListResponse) SetData(v *GetAwemeBindTemplateListResponseData) *GetAwemeBindTemplateListResponse {
	s.Data = v
	return s
}

type GetAwemeBindTemplateListResponseData struct {
	TemplateMap map[string][]*GetAwemeBindTemplateListResponseDataTemplateMapValueItem `json:"template_map,omitempty" xml:"template_map,omitempty"`
}

func (s GetAwemeBindTemplateListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateListResponseData) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateListResponseData) SetTemplateMap(v map[string][]*GetAwemeBindTemplateListResponseDataTemplateMapValueItem) *GetAwemeBindTemplateListResponseData {
	s.TemplateMap = v
	return s
}

type GetAwemeBindTemplateListResponseDataTemplateMapValueItem struct {
	TemplateId   *int64  `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
	CategoryId   *string `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	CanApply     *bool   `json:"can_apply,omitempty" xml:"can_apply,omitempty" require:"true"`
	Reason       *string `json:"reason,omitempty" xml:"reason,omitempty"`
	CategoryName *string `json:"category_name,omitempty" xml:"category_name,omitempty" require:"true"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s GetAwemeBindTemplateListResponseDataTemplateMapValueItem) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeBindTemplateListResponseDataTemplateMapValueItem) GoString() string {
	return s.String()
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetTemplateId(v int64) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.TemplateId = &v
	return s
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetCategoryId(v string) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.CategoryId = &v
	return s
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetCanApply(v bool) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.CanApply = &v
	return s
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetReason(v string) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.Reason = &v
	return s
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetCategoryName(v string) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.CategoryName = &v
	return s
}

func (s *GetAwemeBindTemplateListResponseDataTemplateMapValueItem) SetTitle(v string) *GetAwemeBindTemplateListResponseDataTemplateMapValueItem {
	s.Title = &v
	return s
}

type GetAwemeRelationBindQrcodeRequest struct {
	CapacityList []*string          `json:"capacity_list,omitempty" xml:"capacity_list,omitempty" require:"true" type:"Repeated"`
	Type         *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CoSubject    *bool              `json:"co_subject,omitempty" xml:"co_subject,omitempty"`
}

func (s GetAwemeRelationBindQrcodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeRelationBindQrcodeRequest) GoString() string {
	return s.String()
}

func (s *GetAwemeRelationBindQrcodeRequest) SetCapacityList(v []*string) *GetAwemeRelationBindQrcodeRequest {
	s.CapacityList = v
	return s
}

func (s *GetAwemeRelationBindQrcodeRequest) SetType(v string) *GetAwemeRelationBindQrcodeRequest {
	s.Type = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeRequest) SetHeader(v map[string]*string) *GetAwemeRelationBindQrcodeRequest {
	s.Header = v
	return s
}

func (s *GetAwemeRelationBindQrcodeRequest) SetAccessToken(v string) *GetAwemeRelationBindQrcodeRequest {
	s.AccessToken = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeRequest) SetCoSubject(v bool) *GetAwemeRelationBindQrcodeRequest {
	s.CoSubject = &v
	return s
}

type GetAwemeRelationBindQrcodeResponse struct {
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetAwemeRelationBindQrcodeResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s GetAwemeRelationBindQrcodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeRelationBindQrcodeResponse) GoString() string {
	return s.String()
}

func (s *GetAwemeRelationBindQrcodeResponse) SetErrNo(v int32) *GetAwemeRelationBindQrcodeResponse {
	s.ErrNo = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeResponse) SetErrMsg(v string) *GetAwemeRelationBindQrcodeResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeResponse) SetLogId(v string) *GetAwemeRelationBindQrcodeResponse {
	s.LogId = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeResponse) SetData(v *GetAwemeRelationBindQrcodeResponseData) *GetAwemeRelationBindQrcodeResponse {
	s.Data = v
	return s
}

type GetAwemeRelationBindQrcodeResponseData struct {
	QrcodeParseContent *string `json:"qrcode_parse_content,omitempty" xml:"qrcode_parse_content,omitempty"`
	QrcodeUrl          *string `json:"qrcode_url,omitempty" xml:"qrcode_url,omitempty"`
}

func (s GetAwemeRelationBindQrcodeResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetAwemeRelationBindQrcodeResponseData) GoString() string {
	return s.String()
}

func (s *GetAwemeRelationBindQrcodeResponseData) SetQrcodeParseContent(v string) *GetAwemeRelationBindQrcodeResponseData {
	s.QrcodeParseContent = &v
	return s
}

func (s *GetAwemeRelationBindQrcodeResponseData) SetQrcodeUrl(v string) *GetAwemeRelationBindQrcodeResponseData {
	s.QrcodeUrl = &v
	return s
}

type GetBillDownloadUrlRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	BillType    *string            `json:"bill_type,omitempty" xml:"bill_type,omitempty" require:"true"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
}

func (s GetBillDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBillDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetBillDownloadUrlRequest) SetHeader(v map[string]*string) *GetBillDownloadUrlRequest {
	s.Header = v
	return s
}

func (s *GetBillDownloadUrlRequest) SetAccessToken(v string) *GetBillDownloadUrlRequest {
	s.AccessToken = &v
	return s
}

func (s *GetBillDownloadUrlRequest) SetAppId(v string) *GetBillDownloadUrlRequest {
	s.AppId = &v
	return s
}

func (s *GetBillDownloadUrlRequest) SetBillType(v string) *GetBillDownloadUrlRequest {
	s.BillType = &v
	return s
}

func (s *GetBillDownloadUrlRequest) SetBillDate(v string) *GetBillDownloadUrlRequest {
	s.BillDate = &v
	return s
}

type GetBillDownloadUrlResponse struct {
	IsGenerated *bool                           `json:"is_generated,omitempty" xml:"is_generated,omitempty" require:"true"`
	ErrNo       *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg      *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId       *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data        *GetBillDownloadUrlResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetBillDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBillDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetBillDownloadUrlResponse) SetIsGenerated(v bool) *GetBillDownloadUrlResponse {
	s.IsGenerated = &v
	return s
}

func (s *GetBillDownloadUrlResponse) SetErrNo(v int32) *GetBillDownloadUrlResponse {
	s.ErrNo = &v
	return s
}

func (s *GetBillDownloadUrlResponse) SetErrMsg(v string) *GetBillDownloadUrlResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetBillDownloadUrlResponse) SetLogId(v string) *GetBillDownloadUrlResponse {
	s.LogId = &v
	return s
}

func (s *GetBillDownloadUrlResponse) SetData(v *GetBillDownloadUrlResponseData) *GetBillDownloadUrlResponse {
	s.Data = v
	return s
}

type GetBillDownloadUrlResponseData struct {
	BillDownloadUrl *string `json:"bill_download_url,omitempty" xml:"bill_download_url,omitempty" require:"true"`
}

func (s GetBillDownloadUrlResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetBillDownloadUrlResponseData) GoString() string {
	return s.String()
}

func (s *GetBillDownloadUrlResponseData) SetBillDownloadUrl(v string) *GetBillDownloadUrlResponseData {
	s.BillDownloadUrl = &v
	return s
}

type GetCouponMetaStatisticsRequest struct {
	CouponMetaId  *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	ActivityId    *string            `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	TalentAccount *string            `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TalentOpenId  *string            `json:"talent_open_id,omitempty" xml:"talent_open_id,omitempty"`
}

func (s GetCouponMetaStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCouponMetaStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetCouponMetaStatisticsRequest) SetCouponMetaId(v string) *GetCouponMetaStatisticsRequest {
	s.CouponMetaId = &v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetAppId(v string) *GetCouponMetaStatisticsRequest {
	s.AppId = &v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetActivityId(v string) *GetCouponMetaStatisticsRequest {
	s.ActivityId = &v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetTalentAccount(v string) *GetCouponMetaStatisticsRequest {
	s.TalentAccount = &v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetHeader(v map[string]*string) *GetCouponMetaStatisticsRequest {
	s.Header = v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetAccessToken(v string) *GetCouponMetaStatisticsRequest {
	s.AccessToken = &v
	return s
}

func (s *GetCouponMetaStatisticsRequest) SetTalentOpenId(v string) *GetCouponMetaStatisticsRequest {
	s.TalentOpenId = &v
	return s
}

type GetCouponMetaStatisticsResponse struct {
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetCouponMetaStatisticsResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s GetCouponMetaStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCouponMetaStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCouponMetaStatisticsResponse) SetErrMsg(v string) *GetCouponMetaStatisticsResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetCouponMetaStatisticsResponse) SetLogId(v string) *GetCouponMetaStatisticsResponse {
	s.LogId = &v
	return s
}

func (s *GetCouponMetaStatisticsResponse) SetData(v *GetCouponMetaStatisticsResponseData) *GetCouponMetaStatisticsResponse {
	s.Data = v
	return s
}

func (s *GetCouponMetaStatisticsResponse) SetErrNo(v int32) *GetCouponMetaStatisticsResponse {
	s.ErrNo = &v
	return s
}

type GetCouponMetaStatisticsResponseData struct {
	TalentStats   []*GetCouponMetaStatisticsResponseDataTalentStatsItem   `json:"talent_stats,omitempty" xml:"talent_stats,omitempty" type:"Repeated"`
	ActivityStats []*GetCouponMetaStatisticsResponseDataActivityStatsItem `json:"activity_stats,omitempty" xml:"activity_stats,omitempty" type:"Repeated"`
}

func (s GetCouponMetaStatisticsResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCouponMetaStatisticsResponseData) GoString() string {
	return s.String()
}

func (s *GetCouponMetaStatisticsResponseData) SetTalentStats(v []*GetCouponMetaStatisticsResponseDataTalentStatsItem) *GetCouponMetaStatisticsResponseData {
	s.TalentStats = v
	return s
}

func (s *GetCouponMetaStatisticsResponseData) SetActivityStats(v []*GetCouponMetaStatisticsResponseDataActivityStatsItem) *GetCouponMetaStatisticsResponseData {
	s.ActivityStats = v
	return s
}

type GetCouponMetaStatisticsResponseDataActivityStatsItem struct {
	ActivityId  *string `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	ReceivedNum *int64  `json:"received_num,omitempty" xml:"received_num,omitempty" require:"true"`
	ConsumedNum *int64  `json:"consumed_num,omitempty" xml:"consumed_num,omitempty" require:"true"`
}

func (s GetCouponMetaStatisticsResponseDataActivityStatsItem) String() string {
	return tea.Prettify(s)
}

func (s GetCouponMetaStatisticsResponseDataActivityStatsItem) GoString() string {
	return s.String()
}

func (s *GetCouponMetaStatisticsResponseDataActivityStatsItem) SetActivityId(v string) *GetCouponMetaStatisticsResponseDataActivityStatsItem {
	s.ActivityId = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataActivityStatsItem) SetReceivedNum(v int64) *GetCouponMetaStatisticsResponseDataActivityStatsItem {
	s.ReceivedNum = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataActivityStatsItem) SetConsumedNum(v int64) *GetCouponMetaStatisticsResponseDataActivityStatsItem {
	s.ConsumedNum = &v
	return s
}

type GetCouponMetaStatisticsResponseDataTalentStatsItem struct {
	ExposedNum    *int64  `json:"exposed_num,omitempty" xml:"exposed_num,omitempty" require:"true"`
	ReceivedNum   *int64  `json:"received_num,omitempty" xml:"received_num,omitempty" require:"true"`
	ConsumedNum   *int64  `json:"consumed_num,omitempty" xml:"consumed_num,omitempty" require:"true"`
	TalentAccount *string `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	TalentOpenId  *string `json:"talent_open_id,omitempty" xml:"talent_open_id,omitempty" require:"true"`
}

func (s GetCouponMetaStatisticsResponseDataTalentStatsItem) String() string {
	return tea.Prettify(s)
}

func (s GetCouponMetaStatisticsResponseDataTalentStatsItem) GoString() string {
	return s.String()
}

func (s *GetCouponMetaStatisticsResponseDataTalentStatsItem) SetExposedNum(v int64) *GetCouponMetaStatisticsResponseDataTalentStatsItem {
	s.ExposedNum = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataTalentStatsItem) SetReceivedNum(v int64) *GetCouponMetaStatisticsResponseDataTalentStatsItem {
	s.ReceivedNum = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataTalentStatsItem) SetConsumedNum(v int64) *GetCouponMetaStatisticsResponseDataTalentStatsItem {
	s.ConsumedNum = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataTalentStatsItem) SetTalentAccount(v string) *GetCouponMetaStatisticsResponseDataTalentStatsItem {
	s.TalentAccount = &v
	return s
}

func (s *GetCouponMetaStatisticsResponseDataTalentStatsItem) SetTalentOpenId(v string) *GetCouponMetaStatisticsResponseDataTalentStatsItem {
	s.TalentOpenId = &v
	return s
}

type GetCouponReceiveInfoRequest struct {
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	CouponId     *string            `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	CouponStatus *int               `json:"coupon_status,omitempty" xml:"coupon_status,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetCouponReceiveInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCouponReceiveInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCouponReceiveInfoRequest) SetAppId(v string) *GetCouponReceiveInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetCouponReceiveInfoRequest) SetOpenId(v string) *GetCouponReceiveInfoRequest {
	s.OpenId = &v
	return s
}

func (s *GetCouponReceiveInfoRequest) SetCouponId(v string) *GetCouponReceiveInfoRequest {
	s.CouponId = &v
	return s
}

func (s *GetCouponReceiveInfoRequest) SetCouponStatus(v int) *GetCouponReceiveInfoRequest {
	s.CouponStatus = &v
	return s
}

func (s *GetCouponReceiveInfoRequest) SetHeader(v map[string]*string) *GetCouponReceiveInfoRequest {
	s.Header = v
	return s
}

func (s *GetCouponReceiveInfoRequest) SetAccessToken(v string) *GetCouponReceiveInfoRequest {
	s.AccessToken = &v
	return s
}

type GetCouponReceiveInfoResponse struct {
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetCouponReceiveInfoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s GetCouponReceiveInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCouponReceiveInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCouponReceiveInfoResponse) SetErrNo(v int32) *GetCouponReceiveInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *GetCouponReceiveInfoResponse) SetErrMsg(v string) *GetCouponReceiveInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetCouponReceiveInfoResponse) SetLogId(v string) *GetCouponReceiveInfoResponse {
	s.LogId = &v
	return s
}

func (s *GetCouponReceiveInfoResponse) SetData(v *GetCouponReceiveInfoResponseData) *GetCouponReceiveInfoResponse {
	s.Data = v
	return s
}

type GetCouponReceiveInfoResponseData struct {
	CouponReceiveList []*GetCouponReceiveInfoResponseDataCouponReceiveListItem `json:"coupon_receive_list,omitempty" xml:"coupon_receive_list,omitempty" require:"true" type:"Repeated"`
}

func (s GetCouponReceiveInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetCouponReceiveInfoResponseData) GoString() string {
	return s.String()
}

func (s *GetCouponReceiveInfoResponseData) SetCouponReceiveList(v []*GetCouponReceiveInfoResponseDataCouponReceiveListItem) *GetCouponReceiveInfoResponseData {
	s.CouponReceiveList = v
	return s
}

type GetCouponReceiveInfoResponseDataCouponReceiveListItem struct {
	ValidBeginTime *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty" require:"true"`
	MerchantMetaNo *string `json:"merchant_meta_no,omitempty" xml:"merchant_meta_no,omitempty" require:"true"`
	TalentOpenId   *string `json:"talent_open_id,omitempty" xml:"talent_open_id,omitempty"`
	DiscountType   *int32  `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	CouponId       *string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty" require:"true"`
	DiscountAmount *int64  `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	MinPayAmount   *int64  `json:"min_pay_amount,omitempty" xml:"min_pay_amount,omitempty"`
	ReceiveDesc    *string `json:"receive_desc,omitempty" xml:"receive_desc,omitempty"`
	ValidEndTime   *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty" require:"true"`
	CouponName     *string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty" require:"true"`
	CouponStatus   *int    `json:"coupon_status,omitempty" xml:"coupon_status,omitempty" require:"true"`
	ReceiveTime    *int64  `json:"receive_time,omitempty" xml:"receive_time,omitempty" require:"true"`
	ConsumePath    *string `json:"consume_path,omitempty" xml:"consume_path,omitempty"`
	TalentAccount  *string `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	ConsumeDesc    *string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty"`
}

func (s GetCouponReceiveInfoResponseDataCouponReceiveListItem) String() string {
	return tea.Prettify(s)
}

func (s GetCouponReceiveInfoResponseDataCouponReceiveListItem) GoString() string {
	return s.String()
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetValidBeginTime(v int64) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ValidBeginTime = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetMerchantMetaNo(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.MerchantMetaNo = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetTalentOpenId(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.TalentOpenId = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetDiscountType(v int32) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.DiscountType = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetCouponId(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.CouponId = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetDiscountAmount(v int64) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.DiscountAmount = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetMinPayAmount(v int64) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.MinPayAmount = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetReceiveDesc(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ReceiveDesc = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetValidEndTime(v int64) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ValidEndTime = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetCouponName(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.CouponName = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetCouponStatus(v int) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.CouponStatus = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetReceiveTime(v int64) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ReceiveTime = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetConsumePath(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ConsumePath = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetTalentAccount(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.TalentAccount = &v
	return s
}

func (s *GetCouponReceiveInfoResponseDataCouponReceiveListItem) SetConsumeDesc(v string) *GetCouponReceiveInfoResponseDataCouponReceiveListItem {
	s.ConsumeDesc = &v
	return s
}

type GetRetainConsultCardRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetRetainConsultCardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRetainConsultCardRequest) GoString() string {
	return s.String()
}

func (s *GetRetainConsultCardRequest) SetOpenId(v string) *GetRetainConsultCardRequest {
	s.OpenId = &v
	return s
}

func (s *GetRetainConsultCardRequest) SetHeader(v map[string]*string) *GetRetainConsultCardRequest {
	s.Header = v
	return s
}

func (s *GetRetainConsultCardRequest) SetAccessToken(v string) *GetRetainConsultCardRequest {
	s.AccessToken = &v
	return s
}

type GetRetainConsultCardResponse struct {
	Extra *GetRetainConsultCardResponseExtra       `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *GetRetainConsultCardResponseData        `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Cards []*GetRetainConsultCardResponseCardsItem `json:"cards,omitempty" xml:"cards,omitempty" require:"true" type:"Repeated"`
}

func (s GetRetainConsultCardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRetainConsultCardResponse) GoString() string {
	return s.String()
}

func (s *GetRetainConsultCardResponse) SetExtra(v *GetRetainConsultCardResponseExtra) *GetRetainConsultCardResponse {
	s.Extra = v
	return s
}

func (s *GetRetainConsultCardResponse) SetData(v *GetRetainConsultCardResponseData) *GetRetainConsultCardResponse {
	s.Data = v
	return s
}

func (s *GetRetainConsultCardResponse) SetCards(v []*GetRetainConsultCardResponseCardsItem) *GetRetainConsultCardResponse {
	s.Cards = v
	return s
}

type GetRetainConsultCardResponseCardsItem struct {
	Components []*int  `json:"components,omitempty" xml:"components,omitempty" require:"true" type:"Repeated"`
	MediaId    *string `json:"media_id,omitempty" xml:"media_id,omitempty" require:"true"`
	Status     *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	CardId     *string `json:"card_id,omitempty" xml:"card_id,omitempty" require:"true"`
}

func (s GetRetainConsultCardResponseCardsItem) String() string {
	return tea.Prettify(s)
}

func (s GetRetainConsultCardResponseCardsItem) GoString() string {
	return s.String()
}

func (s *GetRetainConsultCardResponseCardsItem) SetComponents(v []*int) *GetRetainConsultCardResponseCardsItem {
	s.Components = v
	return s
}

func (s *GetRetainConsultCardResponseCardsItem) SetMediaId(v string) *GetRetainConsultCardResponseCardsItem {
	s.MediaId = &v
	return s
}

func (s *GetRetainConsultCardResponseCardsItem) SetStatus(v int) *GetRetainConsultCardResponseCardsItem {
	s.Status = &v
	return s
}

func (s *GetRetainConsultCardResponseCardsItem) SetTitle(v string) *GetRetainConsultCardResponseCardsItem {
	s.Title = &v
	return s
}

func (s *GetRetainConsultCardResponseCardsItem) SetCardId(v string) *GetRetainConsultCardResponseCardsItem {
	s.CardId = &v
	return s
}

type GetRetainConsultCardResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s GetRetainConsultCardResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetRetainConsultCardResponseData) GoString() string {
	return s.String()
}

func (s *GetRetainConsultCardResponseData) SetGwDescription(v string) *GetRetainConsultCardResponseData {
	s.GwDescription = &v
	return s
}

func (s *GetRetainConsultCardResponseData) SetGwErrorCode(v int32) *GetRetainConsultCardResponseData {
	s.GwErrorCode = &v
	return s
}

type GetRetainConsultCardResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GetRetainConsultCardResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s GetRetainConsultCardResponseExtra) GoString() string {
	return s.String()
}

func (s *GetRetainConsultCardResponseExtra) SetErrorCode(v int32) *GetRetainConsultCardResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *GetRetainConsultCardResponseExtra) SetLogid(v string) *GetRetainConsultCardResponseExtra {
	s.Logid = &v
	return s
}

func (s *GetRetainConsultCardResponseExtra) SetNow(v int64) *GetRetainConsultCardResponseExtra {
	s.Now = &v
	return s
}

func (s *GetRetainConsultCardResponseExtra) SetSubDescription(v string) *GetRetainConsultCardResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *GetRetainConsultCardResponseExtra) SetSubErrorCode(v int32) *GetRetainConsultCardResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *GetRetainConsultCardResponseExtra) SetDescription(v string) *GetRetainConsultCardResponseExtra {
	s.Description = &v
	return s
}

type GetUserEncryptKeyRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Signature   *string            `json:"signature,omitempty" xml:"signature,omitempty" require:"true"`
	SigMethod   *string            `json:"sig_method,omitempty" xml:"sig_method,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s GetUserEncryptKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserEncryptKeyRequest) GoString() string {
	return s.String()
}

func (s *GetUserEncryptKeyRequest) SetHeader(v map[string]*string) *GetUserEncryptKeyRequest {
	s.Header = v
	return s
}

func (s *GetUserEncryptKeyRequest) SetAccessToken(v string) *GetUserEncryptKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *GetUserEncryptKeyRequest) SetOpenId(v string) *GetUserEncryptKeyRequest {
	s.OpenId = &v
	return s
}

func (s *GetUserEncryptKeyRequest) SetSignature(v string) *GetUserEncryptKeyRequest {
	s.Signature = &v
	return s
}

func (s *GetUserEncryptKeyRequest) SetSigMethod(v string) *GetUserEncryptKeyRequest {
	s.SigMethod = &v
	return s
}

func (s *GetUserEncryptKeyRequest) SetAppId(v string) *GetUserEncryptKeyRequest {
	s.AppId = &v
	return s
}

type GetUserEncryptKeyResponse struct {
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   []*GetUserEncryptKeyResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s GetUserEncryptKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserEncryptKeyResponse) GoString() string {
	return s.String()
}

func (s *GetUserEncryptKeyResponse) SetErrMsg(v string) *GetUserEncryptKeyResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetUserEncryptKeyResponse) SetLogId(v string) *GetUserEncryptKeyResponse {
	s.LogId = &v
	return s
}

func (s *GetUserEncryptKeyResponse) SetData(v []*GetUserEncryptKeyResponseDataItem) *GetUserEncryptKeyResponse {
	s.Data = v
	return s
}

func (s *GetUserEncryptKeyResponse) SetErrNo(v int32) *GetUserEncryptKeyResponse {
	s.ErrNo = &v
	return s
}

type GetUserEncryptKeyResponseDataItem struct {
	ExpireAt   *int64  `json:"expire_at,omitempty" xml:"expire_at,omitempty" require:"true"`
	Iv         *string `json:"iv,omitempty" xml:"iv,omitempty" require:"true"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	EncryptKey *string `json:"encrypt_key,omitempty" xml:"encrypt_key,omitempty" require:"true"`
	Version    *int32  `json:"version,omitempty" xml:"version,omitempty" require:"true"`
}

func (s GetUserEncryptKeyResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s GetUserEncryptKeyResponseDataItem) GoString() string {
	return s.String()
}

func (s *GetUserEncryptKeyResponseDataItem) SetExpireAt(v int64) *GetUserEncryptKeyResponseDataItem {
	s.ExpireAt = &v
	return s
}

func (s *GetUserEncryptKeyResponseDataItem) SetIv(v string) *GetUserEncryptKeyResponseDataItem {
	s.Iv = &v
	return s
}

func (s *GetUserEncryptKeyResponseDataItem) SetCreateTime(v int64) *GetUserEncryptKeyResponseDataItem {
	s.CreateTime = &v
	return s
}

func (s *GetUserEncryptKeyResponseDataItem) SetEncryptKey(v string) *GetUserEncryptKeyResponseDataItem {
	s.EncryptKey = &v
	return s
}

func (s *GetUserEncryptKeyResponseDataItem) SetVersion(v int32) *GetUserEncryptKeyResponseDataItem {
	s.Version = &v
	return s
}

type GetUserGroupTagRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	MpId        *string            `json:"mp_id,omitempty" xml:"mp_id,omitempty" require:"true"`
}

func (s GetUserGroupTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupTagRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupTagRequest) SetOpenId(v string) *GetUserGroupTagRequest {
	s.OpenId = &v
	return s
}

func (s *GetUserGroupTagRequest) SetHeader(v map[string]*string) *GetUserGroupTagRequest {
	s.Header = v
	return s
}

func (s *GetUserGroupTagRequest) SetAccessToken(v string) *GetUserGroupTagRequest {
	s.AccessToken = &v
	return s
}

func (s *GetUserGroupTagRequest) SetMpId(v string) *GetUserGroupTagRequest {
	s.MpId = &v
	return s
}

type GetUserGroupTagResponse struct {
	UserGroupTagList []*GetUserGroupTagResponseUserGroupTagListItem `json:"user_group_tag_list,omitempty" xml:"user_group_tag_list,omitempty" require:"true" type:"Repeated"`
	ErrNo            *int32                                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg           *string                                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId            *string                                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s GetUserGroupTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupTagResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupTagResponse) SetUserGroupTagList(v []*GetUserGroupTagResponseUserGroupTagListItem) *GetUserGroupTagResponse {
	s.UserGroupTagList = v
	return s
}

func (s *GetUserGroupTagResponse) SetErrNo(v int32) *GetUserGroupTagResponse {
	s.ErrNo = &v
	return s
}

func (s *GetUserGroupTagResponse) SetErrMsg(v string) *GetUserGroupTagResponse {
	s.ErrMsg = &v
	return s
}

func (s *GetUserGroupTagResponse) SetLogId(v string) *GetUserGroupTagResponse {
	s.LogId = &v
	return s
}

type GetUserGroupTagResponseUserGroupTagListItem struct {
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty" require:"true"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	TagId      *string `json:"tag_id,omitempty" xml:"tag_id,omitempty" require:"true"`
}

func (s GetUserGroupTagResponseUserGroupTagListItem) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupTagResponseUserGroupTagListItem) GoString() string {
	return s.String()
}

func (s *GetUserGroupTagResponseUserGroupTagListItem) SetTemplateId(v string) *GetUserGroupTagResponseUserGroupTagListItem {
	s.TemplateId = &v
	return s
}

func (s *GetUserGroupTagResponseUserGroupTagListItem) SetStatus(v int32) *GetUserGroupTagResponseUserGroupTagListItem {
	s.Status = &v
	return s
}

func (s *GetUserGroupTagResponseUserGroupTagListItem) SetTagId(v string) *GetUserGroupTagResponseUserGroupTagListItem {
	s.TagId = &v
	return s
}

type GetUserVideoListRequest struct {
	Count       *int               `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	AwemeId     *string            `json:"aweme_id,omitempty" xml:"aweme_id,omitempty" require:"true"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GetUserVideoListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserVideoListRequest) GoString() string {
	return s.String()
}

func (s *GetUserVideoListRequest) SetCount(v int) *GetUserVideoListRequest {
	s.Count = &v
	return s
}

func (s *GetUserVideoListRequest) SetAwemeId(v string) *GetUserVideoListRequest {
	s.AwemeId = &v
	return s
}

func (s *GetUserVideoListRequest) SetCursor(v int64) *GetUserVideoListRequest {
	s.Cursor = &v
	return s
}

func (s *GetUserVideoListRequest) SetHeader(v map[string]*string) *GetUserVideoListRequest {
	s.Header = v
	return s
}

func (s *GetUserVideoListRequest) SetAccessToken(v string) *GetUserVideoListRequest {
	s.AccessToken = &v
	return s
}

type GetUserVideoListResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *GetUserVideoListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s GetUserVideoListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserVideoListResponse) GoString() string {
	return s.String()
}

func (s *GetUserVideoListResponse) SetLogId(v string) *GetUserVideoListResponse {
	s.LogId = &v
	return s
}

func (s *GetUserVideoListResponse) SetData(v *GetUserVideoListResponseData) *GetUserVideoListResponse {
	s.Data = v
	return s
}

func (s *GetUserVideoListResponse) SetErrNo(v int32) *GetUserVideoListResponse {
	s.ErrNo = &v
	return s
}

func (s *GetUserVideoListResponse) SetErrMsg(v string) *GetUserVideoListResponse {
	s.ErrMsg = &v
	return s
}

type GetUserVideoListResponseData struct {
	UserPublishItemList []*GetUserVideoListResponseDataUserPublishItemListItem `json:"user_publish_item_list,omitempty" xml:"user_publish_item_list,omitempty" require:"true" type:"Repeated"`
	MaxCursor           *int64                                                 `json:"max_cursor,omitempty" xml:"max_cursor,omitempty" require:"true"`
	HasMore             *bool                                                  `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
}

func (s GetUserVideoListResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetUserVideoListResponseData) GoString() string {
	return s.String()
}

func (s *GetUserVideoListResponseData) SetUserPublishItemList(v []*GetUserVideoListResponseDataUserPublishItemListItem) *GetUserVideoListResponseData {
	s.UserPublishItemList = v
	return s
}

func (s *GetUserVideoListResponseData) SetMaxCursor(v int64) *GetUserVideoListResponseData {
	s.MaxCursor = &v
	return s
}

func (s *GetUserVideoListResponseData) SetHasMore(v bool) *GetUserVideoListResponseData {
	s.HasMore = &v
	return s
}

type GetUserVideoListResponseDataUserPublishItemListItem struct {
	VideoLink  *string `json:"video_link,omitempty" xml:"video_link,omitempty" require:"true"`
	ItemId     *int64  `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	CreateTime *string `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
}

func (s GetUserVideoListResponseDataUserPublishItemListItem) String() string {
	return tea.Prettify(s)
}

func (s GetUserVideoListResponseDataUserPublishItemListItem) GoString() string {
	return s.String()
}

func (s *GetUserVideoListResponseDataUserPublishItemListItem) SetVideoLink(v string) *GetUserVideoListResponseDataUserPublishItemListItem {
	s.VideoLink = &v
	return s
}

func (s *GetUserVideoListResponseDataUserPublishItemListItem) SetItemId(v int64) *GetUserVideoListResponseDataUserPublishItemListItem {
	s.ItemId = &v
	return s
}

func (s *GetUserVideoListResponseDataUserPublishItemListItem) SetCreateTime(v string) *GetUserVideoListResponseDataUserPublishItemListItem {
	s.CreateTime = &v
	return s
}

type GiftReceiveRewardRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	GiftCode    *string            `json:"gift_code,omitempty" xml:"gift_code,omitempty" require:"true"`
	EnvType     *string            `json:"env_type,omitempty" xml:"env_type,omitempty"`
	Uuid        *string            `json:"uuid,omitempty" xml:"uuid,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s GiftReceiveRewardRequest) String() string {
	return tea.Prettify(s)
}

func (s GiftReceiveRewardRequest) GoString() string {
	return s.String()
}

func (s *GiftReceiveRewardRequest) SetAccessToken(v string) *GiftReceiveRewardRequest {
	s.AccessToken = &v
	return s
}

func (s *GiftReceiveRewardRequest) SetOpenId(v string) *GiftReceiveRewardRequest {
	s.OpenId = &v
	return s
}

func (s *GiftReceiveRewardRequest) SetGiftCode(v string) *GiftReceiveRewardRequest {
	s.GiftCode = &v
	return s
}

func (s *GiftReceiveRewardRequest) SetEnvType(v string) *GiftReceiveRewardRequest {
	s.EnvType = &v
	return s
}

func (s *GiftReceiveRewardRequest) SetUuid(v string) *GiftReceiveRewardRequest {
	s.Uuid = &v
	return s
}

func (s *GiftReceiveRewardRequest) SetHeader(v map[string]*string) *GiftReceiveRewardRequest {
	s.Header = v
	return s
}

type GiftReceiveRewardResponse struct {
	ErrMsg   *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	GiftInfo *GiftReceiveRewardResponseGiftInfo `json:"gift_info,omitempty" xml:"gift_info,omitempty"`
	ErrNo    *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s GiftReceiveRewardResponse) String() string {
	return tea.Prettify(s)
}

func (s GiftReceiveRewardResponse) GoString() string {
	return s.String()
}

func (s *GiftReceiveRewardResponse) SetErrMsg(v string) *GiftReceiveRewardResponse {
	s.ErrMsg = &v
	return s
}

func (s *GiftReceiveRewardResponse) SetGiftInfo(v *GiftReceiveRewardResponseGiftInfo) *GiftReceiveRewardResponse {
	s.GiftInfo = v
	return s
}

func (s *GiftReceiveRewardResponse) SetErrNo(v int32) *GiftReceiveRewardResponse {
	s.ErrNo = &v
	return s
}

type GiftReceiveRewardResponseGiftInfo struct {
	Name                   *string                                          `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	PlayType               *int                                             `json:"play_type,omitempty" xml:"play_type,omitempty" require:"true"`
	IconUrl                *string                                          `json:"icon_url,omitempty" xml:"icon_url,omitempty" require:"true"`
	GiftId                 *string                                          `json:"gift_id,omitempty" xml:"gift_id,omitempty" require:"true"`
	UserReceiveGuide       []*string                                        `json:"user_receive_guide,omitempty" xml:"user_receive_guide,omitempty" require:"true" type:"Repeated"`
	GiftEffectiveStartTime *int64                                           `json:"gift_effective_start_time,omitempty" xml:"gift_effective_start_time,omitempty" require:"true"`
	PropList               []*GiftReceiveRewardResponseGiftInfoPropListItem `json:"prop_list,omitempty" xml:"prop_list,omitempty" require:"true" type:"Repeated"`
	GiftEffectiveEndTime   *int64                                           `json:"gift_effective_end_time,omitempty" xml:"gift_effective_end_time,omitempty" require:"true"`
}

func (s GiftReceiveRewardResponseGiftInfo) String() string {
	return tea.Prettify(s)
}

func (s GiftReceiveRewardResponseGiftInfo) GoString() string {
	return s.String()
}

func (s *GiftReceiveRewardResponseGiftInfo) SetName(v string) *GiftReceiveRewardResponseGiftInfo {
	s.Name = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetPlayType(v int) *GiftReceiveRewardResponseGiftInfo {
	s.PlayType = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetIconUrl(v string) *GiftReceiveRewardResponseGiftInfo {
	s.IconUrl = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetGiftId(v string) *GiftReceiveRewardResponseGiftInfo {
	s.GiftId = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetUserReceiveGuide(v []*string) *GiftReceiveRewardResponseGiftInfo {
	s.UserReceiveGuide = v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetGiftEffectiveStartTime(v int64) *GiftReceiveRewardResponseGiftInfo {
	s.GiftEffectiveStartTime = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetPropList(v []*GiftReceiveRewardResponseGiftInfoPropListItem) *GiftReceiveRewardResponseGiftInfo {
	s.PropList = v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfo) SetGiftEffectiveEndTime(v int64) *GiftReceiveRewardResponseGiftInfo {
	s.GiftEffectiveEndTime = &v
	return s
}

type GiftReceiveRewardResponseGiftInfoPropListItem struct {
	Count  *int32  `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	PropId *string `json:"prop_id,omitempty" xml:"prop_id,omitempty" require:"true"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Icon   *string `json:"icon,omitempty" xml:"icon,omitempty" require:"true"`
}

func (s GiftReceiveRewardResponseGiftInfoPropListItem) String() string {
	return tea.Prettify(s)
}

func (s GiftReceiveRewardResponseGiftInfoPropListItem) GoString() string {
	return s.String()
}

func (s *GiftReceiveRewardResponseGiftInfoPropListItem) SetCount(v int32) *GiftReceiveRewardResponseGiftInfoPropListItem {
	s.Count = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfoPropListItem) SetPropId(v string) *GiftReceiveRewardResponseGiftInfoPropListItem {
	s.PropId = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfoPropListItem) SetName(v string) *GiftReceiveRewardResponseGiftInfoPropListItem {
	s.Name = &v
	return s
}

func (s *GiftReceiveRewardResponseGiftInfoPropListItem) SetIcon(v string) *GiftReceiveRewardResponseGiftInfoPropListItem {
	s.Icon = &v
	return s
}

type GoodsProductDraftGetRequest struct {
	NeedPoi      *bool              `json:"need_poi,omitempty" xml:"need_poi,omitempty"`
	OutIds       []*string          `json:"out_ids,omitempty" xml:"out_ids,omitempty" type:"Repeated"`
	ProductIds   []*string          `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	AccountId    *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GetDraftType *int               `json:"get_draft_type,omitempty" xml:"get_draft_type,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s GoodsProductDraftGetRequest) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetRequest) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetRequest) SetNeedPoi(v bool) *GoodsProductDraftGetRequest {
	s.NeedPoi = &v
	return s
}

func (s *GoodsProductDraftGetRequest) SetOutIds(v []*string) *GoodsProductDraftGetRequest {
	s.OutIds = v
	return s
}

func (s *GoodsProductDraftGetRequest) SetProductIds(v []*string) *GoodsProductDraftGetRequest {
	s.ProductIds = v
	return s
}

func (s *GoodsProductDraftGetRequest) SetAccountId(v string) *GoodsProductDraftGetRequest {
	s.AccountId = &v
	return s
}

func (s *GoodsProductDraftGetRequest) SetGetDraftType(v int) *GoodsProductDraftGetRequest {
	s.GetDraftType = &v
	return s
}

func (s *GoodsProductDraftGetRequest) SetHeader(v map[string]*string) *GoodsProductDraftGetRequest {
	s.Header = v
	return s
}

func (s *GoodsProductDraftGetRequest) SetAccessToken(v string) *GoodsProductDraftGetRequest {
	s.AccessToken = &v
	return s
}

type GoodsProductDraftGetResponse struct {
	Data     *GoodsProductDraftGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *GoodsProductDraftGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *GoodsProductDraftGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponse) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponse) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponse) SetData(v *GoodsProductDraftGetResponseData) *GoodsProductDraftGetResponse {
	s.Data = v
	return s
}

func (s *GoodsProductDraftGetResponse) SetExtra(v *GoodsProductDraftGetResponseExtra) *GoodsProductDraftGetResponse {
	s.Extra = v
	return s
}

func (s *GoodsProductDraftGetResponse) SetBaseResp(v *GoodsProductDraftGetResponseBaseResp) *GoodsProductDraftGetResponse {
	s.BaseResp = v
	return s
}

type GoodsProductDraftGetResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s GoodsProductDraftGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseBaseResp) SetExtra(v map[string]*string) *GoodsProductDraftGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *GoodsProductDraftGetResponseBaseResp) SetStatusCode(v int32) *GoodsProductDraftGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseBaseResp) SetStatusMessage(v string) *GoodsProductDraftGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type GoodsProductDraftGetResponseData struct {
	ErrorCode     *int32                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductDrafts []*GoodsProductDraftGetResponseDataProductDraftsItem `json:"product_drafts,omitempty" xml:"product_drafts,omitempty" type:"Repeated"`
	Description   *string                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseData) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseData) SetErrorCode(v int32) *GoodsProductDraftGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseData) SetProductDrafts(v []*GoodsProductDraftGetResponseDataProductDraftsItem) *GoodsProductDraftGetResponseData {
	s.ProductDrafts = v
	return s
}

func (s *GoodsProductDraftGetResponseData) SetDescription(v string) *GoodsProductDraftGetResponseData {
	s.Description = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItem struct {
	DraftStatus     *int                                                         `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	BanOnlineStatus *bool                                                        `json:"ban_online_status,omitempty" xml:"ban_online_status,omitempty"`
	AuditMsg        *string                                                      `json:"audit_msg,omitempty" xml:"audit_msg,omitempty"`
	Sku             *GoodsProductDraftGetResponseDataProductDraftsItemSku        `json:"sku,omitempty" xml:"sku,omitempty"`
	Extra           *string                                                      `json:"extra,omitempty" xml:"extra,omitempty"`
	Product         *GoodsProductDraftGetResponseDataProductDraftsItemProduct    `json:"product,omitempty" xml:"product,omitempty"`
	Skus            []*GoodsProductDraftGetResponseDataProductDraftsItemSkusItem `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	BanOnlineCode   *int64                                                       `json:"ban_online_code,omitempty" xml:"ban_online_code,omitempty"`
	BanOnlineReason *string                                                      `json:"ban_online_reason,omitempty" xml:"ban_online_reason,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItem) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetDraftStatus(v int) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.DraftStatus = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetBanOnlineStatus(v bool) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.BanOnlineStatus = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetAuditMsg(v string) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.AuditMsg = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetSku(v *GoodsProductDraftGetResponseDataProductDraftsItemSku) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.Sku = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetExtra(v string) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.Extra = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetProduct(v *GoodsProductDraftGetResponseDataProductDraftsItemProduct) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.Product = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetSkus(v []*GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.Skus = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetBanOnlineCode(v int64) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.BanOnlineCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItem) SetBanOnlineReason(v string) *GoodsProductDraftGetResponseDataProductDraftsItem {
	s.BanOnlineReason = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemProduct struct {
	UpdateTime       *int64                                                              `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductSubType   *int                                                                `json:"product_sub_type,omitempty" xml:"product_sub_type,omitempty"`
	CategoryId       *int64                                                              `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	Extra            *string                                                             `json:"extra,omitempty" xml:"extra,omitempty"`
	ProductId        *string                                                             `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SoldEndTime      *int64                                                              `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	SoldStartTime    *int64                                                              `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	OpenBizType      *int                                                                `json:"open_biz_type,omitempty" xml:"open_biz_type,omitempty"`
	ProductType      *int                                                                `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	AccountName      *string                                                             `json:"account_name,omitempty" xml:"account_name,omitempty"`
	OutId            *string                                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Telephone        []*string                                                           `json:"telephone,omitempty" xml:"telephone,omitempty" type:"Repeated"`
	SpuId            *string                                                             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	CreateTime       *int64                                                              `json:"create_time,omitempty" xml:"create_time,omitempty"`
	CategoryFullName *string                                                             `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	BizLine          *int                                                                `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	ContactName      *string                                                             `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	ProductName      *string                                                             `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	OutUrl           *string                                                             `json:"out_url,omitempty" xml:"out_url,omitempty"`
	CreatorAccountId *int64                                                              `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	AttrKeyValueMap  map[string]*string                                                  `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	Desc             *string                                                             `json:"desc,omitempty" xml:"desc,omitempty"`
	Pois             []*GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ProductExt       *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	Version          *int64                                                              `json:"version,omitempty" xml:"version,omitempty"`
	OwnerAccountId   *int64                                                              `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProduct) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProduct) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetUpdateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.UpdateTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetProductSubType(v int) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ProductSubType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetCategoryId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.CategoryId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetExtra(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.Extra = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetProductId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ProductId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetSoldEndTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.SoldEndTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetSoldStartTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.SoldStartTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetOpenBizType(v int) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.OpenBizType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetProductType(v int) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ProductType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetAccountName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.AccountName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetOutId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.OutId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetTelephone(v []*string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.Telephone = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetSpuId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.SpuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetCreateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.CreateTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetCategoryFullName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.CategoryFullName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetBizLine(v int) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.BizLine = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetContactName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ContactName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetProductName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ProductName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetOutUrl(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.OutUrl = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetCreatorAccountId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetAttrKeyValueMap(v map[string]*string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetDesc(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.Desc = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetPois(v []*GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.Pois = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetProductExt(v *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.ProductExt = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetVersion(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.Version = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProduct) SetOwnerAccountId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProduct {
	s.OwnerAccountId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem struct {
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) SetPoiId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem {
	s.PoiId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) SetSupplierExtId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem) SetSupplierId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProductPoisItem {
	s.SupplierId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt struct {
	ElemeBizCode      *string                                                                         `json:"eleme_biz_code,omitempty" xml:"eleme_biz_code,omitempty"`
	DisplayPrice      *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice `json:"display_price,omitempty" xml:"display_price,omitempty"`
	AllSkuSellOut     *bool                                                                           `json:"all_sku_sell_out,omitempty" xml:"all_sku_sell_out,omitempty"`
	IsBindClueElement *bool                                                                           `json:"is_bind_clue_element,omitempty" xml:"is_bind_clue_element,omitempty"`
	CategoryOutId     *string                                                                         `json:"category_out_id,omitempty" xml:"category_out_id,omitempty"`
	TestExtra         *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra    `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
	AgencyRate        *int64                                                                          `json:"agency_rate,omitempty" xml:"agency_rate,omitempty"`
	AutoOnline        *bool                                                                           `json:"auto_online,omitempty" xml:"auto_online,omitempty"`
	ElemeExtraInfo    *string                                                                         `json:"eleme_extra_info,omitempty" xml:"eleme_extra_info,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetElemeBizCode(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.ElemeBizCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetDisplayPrice(v *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.DisplayPrice = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetAllSkuSellOut(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.AllSkuSellOut = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetIsBindClueElement(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.IsBindClueElement = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetCategoryOutId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.CategoryOutId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetTestExtra(v *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.TestExtra = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetAgencyRate(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.AgencyRate = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetAutoOnline(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.AutoOnline = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt) SetElemeExtraInfo(v string) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExt {
	s.ElemeExtraInfo = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice struct {
	HighPrice *int64 `json:"high_price,omitempty" xml:"high_price,omitempty"`
	LowPrice  *int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice) SetHighPrice(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice {
	s.HighPrice = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice) SetLowPrice(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtDisplayPrice {
	s.LowPrice = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra) SetTestFlag(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra {
	s.TestFlag = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra) SetUids(v []*string) *GoodsProductDraftGetResponseDataProductDraftsItemProductProductExtTestExtra {
	s.Uids = v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSku struct {
	Stock           *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	SkuName         *string                                                     `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	UpdateTime      *int64                                                      `json:"update_time,omitempty" xml:"update_time,omitempty"`
	BindSkus        []*string                                                   `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	OriginAmount    *int64                                                      `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	SkuId           *string                                                     `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Extra           *string                                                     `json:"extra,omitempty" xml:"extra,omitempty"`
	ActualAmount    *int64                                                      `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	AttrKeyValueMap map[string]*string                                          `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	Status          *int                                                        `json:"status,omitempty" xml:"status,omitempty"`
	SkuExt          *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	OutSkuId        *string                                                     `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	CreateTime      *int64                                                      `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSku) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSku) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetStock(v *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.Stock = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetSkuName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.SkuName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetUpdateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.UpdateTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetBindSkus(v []*string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.BindSkus = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetOriginAmount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.OriginAmount = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetSkuId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.SkuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetExtra(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.Extra = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetActualAmount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.ActualAmount = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetAttrKeyValueMap(v map[string]*string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetStatus(v int) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.Status = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetSkuExt(v *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.SkuExt = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetOutSkuId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.OutSkuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSku) SetCreateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSku {
	s.CreateTime = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt struct {
	CMspuId             *int64                                                                         `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	OriSkus             map[int64][]*int64                                                             `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	BindSkus2c          map[int64][]*int64                                                             `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	DiscountPromo       *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	TakeawayPresaleInfo *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	BizId2cList         []*int64                                                                       `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	BindProductId       *int64                                                                         `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	SettleType          *int64                                                                         `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	RelRuleList         []*GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	UseSubRelStock      *bool                                                                          `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
	LifeBizCode         *string                                                                        `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	AccountSettle       *bool                                                                          `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
	OriginStockQty      *int64                                                                         `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	BindSkuId           *int64                                                                         `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetCMspuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.CMspuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetOriSkus(v map[int64][]*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.OriSkus = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetBindSkus2c(v map[int64][]*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetDiscountPromo(v *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetTakeawayPresaleInfo(v *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetBizId2cList(v []*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.BizId2cList = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetBindProductId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.BindProductId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetSettleType(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.SettleType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetRelRuleList(v []*GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.RelRuleList = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetUseSubRelStock(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.UseSubRelStock = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetLifeBizCode(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetAccountSettle(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.AccountSettle = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetOriginStockQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt) SetBindSkuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExt {
	s.BindSkuId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo struct {
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) SetBrandActivityId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) SetPlanId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo) SetPromoId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem struct {
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetStockRel(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetBizId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetCoefficient(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetConstantVal(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetPriceRel(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem) SetSharedQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkuStock struct {
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetSoldQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.SoldQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetStockQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.StockQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetAvailQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.AvailQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetFrozenQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.FrozenQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetLimitType(v int) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.LimitType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock) SetSoldCount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkuStock {
	s.SoldCount = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItem struct {
	Status          *int                                                             `json:"status,omitempty" xml:"status,omitempty"`
	Extra           *string                                                          `json:"extra,omitempty" xml:"extra,omitempty"`
	Stock           *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	UpdateTime      *int64                                                           `json:"update_time,omitempty" xml:"update_time,omitempty"`
	SkuName         *string                                                          `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	AttrKeyValueMap map[string]*string                                               `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	ActualAmount    *int64                                                           `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	CreateTime      *int64                                                           `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OutSkuId        *string                                                          `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	BindSkus        []*string                                                        `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuExt          *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	SkuId           *string                                                          `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OriginAmount    *int64                                                           `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetStatus(v int) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.Status = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetExtra(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.Extra = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetStock(v *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemStock) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.Stock = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetUpdateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetSkuName(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetAttrKeyValueMap(v map[string]*string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetActualAmount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetCreateTime(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.CreateTime = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetOutSkuId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetBindSkus(v []*string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.BindSkus = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetSkuExt(v *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.SkuExt = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetSkuId(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem) SetOriginAmount(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItem {
	s.OriginAmount = &v
	return s
}

type GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt struct {
	SettleType          *int64                                                                              `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	RelRuleList         []*GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	UseSubRelStock      *bool                                                                               `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
	OriSkus             map[int64][]*int64                                                                  `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	BindProductId       *int64                                                                              `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	OriginStockQty      *int64                                                                              `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	TakeawayPresaleInfo *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	CMspuId             *int64                                                                              `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	LifeBizCode         *string                                                                             `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	BindSkus2c          map[int64][]*int64                                                                  `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	BindSkuId           *int64                                                                              `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	BizId2cList         []*int64                                                                            `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	DiscountPromo       *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	AccountSettle       *bool                                                                               `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) String() string {
	return tea.Prettify(s)
}

func (s GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) GoString() string {
	return s.String()
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetSettleType(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.SettleType = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetRelRuleList(v []*GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtRelRuleListItem) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.RelRuleList = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetUseSubRelStock(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.UseSubRelStock = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetOriSkus(v map[int64][]*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.OriSkus = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetBindProductId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.BindProductId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetOriginStockQty(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetTakeawayPresaleInfo(v *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtTakeawayPresaleInfo) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetCMspuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.CMspuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetLifeBizCode(v string) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetBindSkus2c(v map[int64][]*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetBindSkuId(v int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetBizId2cList(v []*int64) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.BizId2cList = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetDiscountPromo(v *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExtDiscountPromo) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt) SetAccountSettle(v bool) *GoodsProductDraftGetResponseDataProductDraftsItemSkusItemSkuExt {
	s.AccountSettle = &v
	return s
}
