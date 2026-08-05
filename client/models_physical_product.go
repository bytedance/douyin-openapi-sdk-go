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

type PhysicalRoomSaveResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PhysicalRoomSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSaveResponseExtra) SetErrorCode(v int32) *PhysicalRoomSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PhysicalRoomSaveResponseExtra) SetLogid(v string) *PhysicalRoomSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PhysicalRoomSaveResponseExtra) SetNow(v int64) *PhysicalRoomSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PhysicalRoomSaveResponseExtra) SetSubDescription(v string) *PhysicalRoomSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PhysicalRoomSaveResponseExtra) SetSubErrorCode(v int32) *PhysicalRoomSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PhysicalRoomSaveResponseExtra) SetDescription(v string) *PhysicalRoomSaveResponseExtra {
	s.Description = &v
	return s
}

type PhysicalRoomSearchRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PoiIds      []*string          `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
}

func (s PhysicalRoomSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchRequest) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchRequest) SetHeader(v map[string]*string) *PhysicalRoomSearchRequest {
	s.Header = v
	return s
}

func (s *PhysicalRoomSearchRequest) SetAccessToken(v string) *PhysicalRoomSearchRequest {
	s.AccessToken = &v
	return s
}

func (s *PhysicalRoomSearchRequest) SetAccountId(v string) *PhysicalRoomSearchRequest {
	s.AccountId = &v
	return s
}

func (s *PhysicalRoomSearchRequest) SetPoiIds(v []*string) *PhysicalRoomSearchRequest {
	s.PoiIds = v
	return s
}

type PhysicalRoomSearchResponse struct {
	Extra *PhysicalRoomSearchResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PhysicalRoomSearchResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PhysicalRoomSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponse) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponse) SetExtra(v *PhysicalRoomSearchResponseExtra) *PhysicalRoomSearchResponse {
	s.Extra = v
	return s
}

func (s *PhysicalRoomSearchResponse) SetData(v *PhysicalRoomSearchResponseData) *PhysicalRoomSearchResponse {
	s.Data = v
	return s
}

type PhysicalRoomSearchResponseData struct {
	RoomLists     map[string][]*PhysicalRoomSearchResponseDataRoomListsValueItem `json:"room_lists,omitempty" xml:"room_lists,omitempty"`
	GwErrorCode   *int32                                                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PhysicalRoomSearchResponseData) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseData) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseData) SetRoomLists(v map[string][]*PhysicalRoomSearchResponseDataRoomListsValueItem) *PhysicalRoomSearchResponseData {
	s.RoomLists = v
	return s
}

func (s *PhysicalRoomSearchResponseData) SetGwErrorCode(v int32) *PhysicalRoomSearchResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PhysicalRoomSearchResponseData) SetGwDescription(v string) *PhysicalRoomSearchResponseData {
	s.GwDescription = &v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItem struct {
	Images       []*PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem                `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	CnName       *string                                                                      `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	Area         *PhysicalRoomSearchResponseDataRoomListsValueItemArea                        `json:"area,omitempty" xml:"area,omitempty"`
	Status       *int                                                                         `json:"status,omitempty" xml:"status,omitempty"`
	Smoking      *int                                                                         `json:"smoking,omitempty" xml:"smoking,omitempty"`
	Floor        []*PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem                 `json:"floor,omitempty" xml:"floor,omitempty" type:"Repeated"`
	MaxOccupancy *int32                                                                       `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	EnName       *string                                                                      `json:"en_name,omitempty" xml:"en_name,omitempty"`
	BedGroups    []*PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem             `json:"bed_groups,omitempty" xml:"bed_groups,omitempty" type:"Repeated"`
	RoomId       *string                                                                      `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Window       *int                                                                         `json:"window,omitempty" xml:"window,omitempty"`
	Internets    []*int                                                                       `json:"internets,omitempty" xml:"internets,omitempty" type:"Repeated"`
	CategoryId   *int64                                                                       `json:"category_id,omitempty" xml:"category_id,omitempty"`
	FacilityMap  map[string]*PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue `json:"facility_map,omitempty" xml:"facility_map,omitempty"`
	Descriptions []*string                                                                    `json:"descriptions,omitempty" xml:"descriptions,omitempty" type:"Repeated"`
	RoomNum      *int32                                                                       `json:"room_num,omitempty" xml:"room_num,omitempty"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetImages(v []*PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Images = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetCnName(v string) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.CnName = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetArea(v *PhysicalRoomSearchResponseDataRoomListsValueItemArea) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Area = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetStatus(v int) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Status = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetSmoking(v int) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Smoking = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetFloor(v []*PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Floor = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetMaxOccupancy(v int32) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.MaxOccupancy = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetEnName(v string) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.EnName = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetBedGroups(v []*PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.BedGroups = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetRoomId(v string) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.RoomId = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetWindow(v int) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Window = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetInternets(v []*int) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Internets = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetCategoryId(v int64) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.CategoryId = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetFacilityMap(v map[string]*PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.FacilityMap = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetDescriptions(v []*string) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.Descriptions = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItem) SetRoomNum(v int32) *PhysicalRoomSearchResponseDataRoomListsValueItem {
	s.RoomNum = &v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemArea struct {
	ValueList []*int32 `json:"value_list,omitempty" xml:"value_list,omitempty" type:"Repeated"`
	ValueType *int     `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemArea) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemArea) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemArea) SetValueList(v []*int32) *PhysicalRoomSearchResponseDataRoomListsValueItemArea {
	s.ValueList = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemArea) SetValueType(v int) *PhysicalRoomSearchResponseDataRoomListsValueItemArea {
	s.ValueType = &v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem struct {
	BedItems []*PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem `json:"bed_items,omitempty" xml:"bed_items,omitempty" require:"true" type:"Repeated"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem) SetBedItems(v []*PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItem {
	s.BedItems = v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem struct {
	BedType *int   `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	BedNum  *int32 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
	BedSize *int32 `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) SetBedType(v int) *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem {
	s.BedType = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) SetBedNum(v int32) *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem {
	s.BedNum = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem) SetBedSize(v int32) *PhysicalRoomSearchResponseDataRoomListsValueItemBedGroupsItemBedItemsItem {
	s.BedSize = &v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue struct {
	FacilityIds []*int64 `json:"facility_ids,omitempty" xml:"facility_ids,omitempty" require:"true" type:"Repeated"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue) SetFacilityIds(v []*int64) *PhysicalRoomSearchResponseDataRoomListsValueItemFacilityMapValue {
	s.FacilityIds = v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem struct {
	ValueList []*int32 `json:"value_list,omitempty" xml:"value_list,omitempty" type:"Repeated"`
	ValueType *int     `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem) SetValueList(v []*int32) *PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem {
	s.ValueList = v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem) SetValueType(v int) *PhysicalRoomSearchResponseDataRoomListsValueItemFloorItem {
	s.ValueType = &v
	return s
}

type PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem struct {
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty" require:"true"`
	ImageUri *string `json:"image_uri,omitempty" xml:"image_uri,omitempty"`
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem) SetImageUrl(v string) *PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem {
	s.ImageUrl = &v
	return s
}

func (s *PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem) SetImageUri(v string) *PhysicalRoomSearchResponseDataRoomListsValueItemImagesItem {
	s.ImageUri = &v
	return s
}

type PhysicalRoomSearchResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PhysicalRoomSearchResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PhysicalRoomSearchResponseExtra) GoString() string {
	return s.String()
}

func (s *PhysicalRoomSearchResponseExtra) SetErrorCode(v int32) *PhysicalRoomSearchResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PhysicalRoomSearchResponseExtra) SetLogid(v string) *PhysicalRoomSearchResponseExtra {
	s.Logid = &v
	return s
}

func (s *PhysicalRoomSearchResponseExtra) SetNow(v int64) *PhysicalRoomSearchResponseExtra {
	s.Now = &v
	return s
}

func (s *PhysicalRoomSearchResponseExtra) SetSubDescription(v string) *PhysicalRoomSearchResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PhysicalRoomSearchResponseExtra) SetSubErrorCode(v int32) *PhysicalRoomSearchResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PhysicalRoomSearchResponseExtra) SetDescription(v string) *PhysicalRoomSearchResponseExtra {
	s.Description = &v
	return s
}

type PlayletBusinessUploadRequest struct {
	Properties  *string                              `json:"properties,omitempty" xml:"properties,omitempty" require:"true"`
	Header      map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	EventType   *string                              `json:"event_type,omitempty" xml:"event_type,omitempty" require:"true"`
	Context     *PlayletBusinessUploadRequestContext `json:"context,omitempty" xml:"context,omitempty" require:"true"`
	Timestamp   *int64                               `json:"timestamp,omitempty" xml:"timestamp,omitempty" require:"true"`
}

func (s PlayletBusinessUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayletBusinessUploadRequest) GoString() string {
	return s.String()
}

func (s *PlayletBusinessUploadRequest) SetProperties(v string) *PlayletBusinessUploadRequest {
	s.Properties = &v
	return s
}

func (s *PlayletBusinessUploadRequest) SetHeader(v map[string]*string) *PlayletBusinessUploadRequest {
	s.Header = v
	return s
}

func (s *PlayletBusinessUploadRequest) SetAccessToken(v string) *PlayletBusinessUploadRequest {
	s.AccessToken = &v
	return s
}

func (s *PlayletBusinessUploadRequest) SetEventType(v string) *PlayletBusinessUploadRequest {
	s.EventType = &v
	return s
}

func (s *PlayletBusinessUploadRequest) SetContext(v *PlayletBusinessUploadRequestContext) *PlayletBusinessUploadRequest {
	s.Context = v
	return s
}

func (s *PlayletBusinessUploadRequest) SetTimestamp(v int64) *PlayletBusinessUploadRequest {
	s.Timestamp = &v
	return s
}

type PlayletBusinessUploadRequestContext struct {
	Device *PlayletBusinessUploadRequestContextDevice `json:"device,omitempty" xml:"device,omitempty" require:"true"`
	Ad     *PlayletBusinessUploadRequestContextAd     `json:"ad,omitempty" xml:"ad,omitempty"`
}

func (s PlayletBusinessUploadRequestContext) String() string {
	return tea.Prettify(s)
}

func (s PlayletBusinessUploadRequestContext) GoString() string {
	return s.String()
}

func (s *PlayletBusinessUploadRequestContext) SetDevice(v *PlayletBusinessUploadRequestContextDevice) *PlayletBusinessUploadRequestContext {
	s.Device = v
	return s
}

func (s *PlayletBusinessUploadRequestContext) SetAd(v *PlayletBusinessUploadRequestContextAd) *PlayletBusinessUploadRequestContext {
	s.Ad = v
	return s
}

type PlayletBusinessUploadRequestContextAd struct {
	Callback *string `json:"callback,omitempty" xml:"callback,omitempty"`
}

func (s PlayletBusinessUploadRequestContextAd) String() string {
	return tea.Prettify(s)
}

func (s PlayletBusinessUploadRequestContextAd) GoString() string {
	return s.String()
}

func (s *PlayletBusinessUploadRequestContextAd) SetCallback(v string) *PlayletBusinessUploadRequestContextAd {
	s.Callback = &v
	return s
}

type PlayletBusinessUploadRequestContextDevice struct {
	OpenId *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s PlayletBusinessUploadRequestContextDevice) String() string {
	return tea.Prettify(s)
}

func (s PlayletBusinessUploadRequestContextDevice) GoString() string {
	return s.String()
}

func (s *PlayletBusinessUploadRequestContextDevice) SetOpenId(v string) *PlayletBusinessUploadRequestContextDevice {
	s.OpenId = &v
	return s
}

type PlayletBusinessUploadResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s PlayletBusinessUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s PlayletBusinessUploadResponse) GoString() string {
	return s.String()
}

func (s *PlayletBusinessUploadResponse) SetLogId(v string) *PlayletBusinessUploadResponse {
	s.LogId = &v
	return s
}

func (s *PlayletBusinessUploadResponse) SetErrNo(v int32) *PlayletBusinessUploadResponse {
	s.ErrNo = &v
	return s
}

func (s *PlayletBusinessUploadResponse) SetErrMsg(v string) *PlayletBusinessUploadResponse {
	s.ErrMsg = &v
	return s
}

type PoiClaimRequest struct {
	Datas       []*PoiClaimRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" require:"true" type:"Repeated"`
	TargetType  *int                        `json:"target_type,omitempty" xml:"target_type,omitempty" require:"true"`
	Header      map[string]*string          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequest) GoString() string {
	return s.String()
}

func (s *PoiClaimRequest) SetDatas(v []*PoiClaimRequestDatasItem) *PoiClaimRequest {
	s.Datas = v
	return s
}

func (s *PoiClaimRequest) SetTargetType(v int) *PoiClaimRequest {
	s.TargetType = &v
	return s
}

func (s *PoiClaimRequest) SetHeader(v map[string]*string) *PoiClaimRequest {
	s.Header = v
	return s
}

func (s *PoiClaimRequest) SetAccessToken(v string) *PoiClaimRequest {
	s.AccessToken = &v
	return s
}

type PoiClaimRequestDatasItem struct {
	LegalPerson           *PoiClaimRequestDatasItemLegalPerson           `json:"legal_person,omitempty" xml:"legal_person,omitempty"`
	PoiId                 *string                                        `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	BusinessPartnership   *PoiClaimRequestDatasItemBusinessPartnership   `json:"business_partnership,omitempty" xml:"business_partnership,omitempty"`
	Industry              *PoiClaimRequestDatasItemIndustry              `json:"industry,omitempty" xml:"industry,omitempty"`
	License               *PoiClaimRequestDatasItemLicense               `json:"license,omitempty" xml:"license,omitempty"`
	ThirdId               *string                                        `json:"third_id,omitempty" xml:"third_id,omitempty"`
	PoiClaimAuthorization *PoiClaimRequestDatasItemPoiClaimAuthorization `json:"poi_claim_authorization,omitempty" xml:"poi_claim_authorization,omitempty"`
	AccountId             *string                                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Owner                 *PoiClaimRequestDatasItemOwner                 `json:"owner,omitempty" xml:"owner,omitempty"`
}

func (s PoiClaimRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItem) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItem) SetLegalPerson(v *PoiClaimRequestDatasItemLegalPerson) *PoiClaimRequestDatasItem {
	s.LegalPerson = v
	return s
}

func (s *PoiClaimRequestDatasItem) SetPoiId(v string) *PoiClaimRequestDatasItem {
	s.PoiId = &v
	return s
}

func (s *PoiClaimRequestDatasItem) SetBusinessPartnership(v *PoiClaimRequestDatasItemBusinessPartnership) *PoiClaimRequestDatasItem {
	s.BusinessPartnership = v
	return s
}

func (s *PoiClaimRequestDatasItem) SetIndustry(v *PoiClaimRequestDatasItemIndustry) *PoiClaimRequestDatasItem {
	s.Industry = v
	return s
}

func (s *PoiClaimRequestDatasItem) SetLicense(v *PoiClaimRequestDatasItemLicense) *PoiClaimRequestDatasItem {
	s.License = v
	return s
}

func (s *PoiClaimRequestDatasItem) SetThirdId(v string) *PoiClaimRequestDatasItem {
	s.ThirdId = &v
	return s
}

func (s *PoiClaimRequestDatasItem) SetPoiClaimAuthorization(v *PoiClaimRequestDatasItemPoiClaimAuthorization) *PoiClaimRequestDatasItem {
	s.PoiClaimAuthorization = v
	return s
}

func (s *PoiClaimRequestDatasItem) SetAccountId(v string) *PoiClaimRequestDatasItem {
	s.AccountId = &v
	return s
}

func (s *PoiClaimRequestDatasItem) SetOwner(v *PoiClaimRequestDatasItemOwner) *PoiClaimRequestDatasItem {
	s.Owner = v
	return s
}

type PoiClaimRequestDatasItemBusinessPartnership struct {
	Authorization      *PoiClaimRequestDatasItemBusinessPartnershipAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty"`
	PartnerAccountType *int                                                      `json:"partner_account_type,omitempty" xml:"partner_account_type,omitempty"`
}

func (s PoiClaimRequestDatasItemBusinessPartnership) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemBusinessPartnership) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemBusinessPartnership) SetAuthorization(v *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) *PoiClaimRequestDatasItemBusinessPartnership {
	s.Authorization = v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnership) SetPartnerAccountType(v int) *PoiClaimRequestDatasItemBusinessPartnership {
	s.PartnerAccountType = &v
	return s
}

type PoiClaimRequestDatasItemBusinessPartnershipAuthorization struct {
	Expiration    *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	Id            *string   `json:"id,omitempty" xml:"id,omitempty"`
	Type          *int      `json:"type,omitempty" xml:"type,omitempty"`
	Urls          []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	Agreement     *int      `json:"agreement,omitempty" xml:"agreement,omitempty"`
	AppIds        []*string `json:"app_ids,omitempty" xml:"app_ids,omitempty" type:"Repeated"`
	EffectiveTime *string   `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
}

func (s PoiClaimRequestDatasItemBusinessPartnershipAuthorization) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemBusinessPartnershipAuthorization) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetExpiration(v string) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.Expiration = &v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetId(v string) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.Id = &v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetType(v int) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.Type = &v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetUrls(v []*string) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.Urls = v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetAgreement(v int) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.Agreement = &v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetAppIds(v []*string) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.AppIds = v
	return s
}

func (s *PoiClaimRequestDatasItemBusinessPartnershipAuthorization) SetEffectiveTime(v string) *PoiClaimRequestDatasItemBusinessPartnershipAuthorization {
	s.EffectiveTime = &v
	return s
}

type PoiClaimRequestDatasItemIndustry struct {
	MajorIndustryCode  *string                                               `json:"major_industry_code,omitempty" xml:"major_industry_code,omitempty"`
	MinorIndustryCodes []*string                                             `json:"minor_industry_codes,omitempty" xml:"minor_industry_codes,omitempty" type:"Repeated"`
	Qualifications     []*PoiClaimRequestDatasItemIndustryQualificationsItem `json:"qualifications,omitempty" xml:"qualifications,omitempty" type:"Repeated"`
}

func (s PoiClaimRequestDatasItemIndustry) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemIndustry) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemIndustry) SetMajorIndustryCode(v string) *PoiClaimRequestDatasItemIndustry {
	s.MajorIndustryCode = &v
	return s
}

func (s *PoiClaimRequestDatasItemIndustry) SetMinorIndustryCodes(v []*string) *PoiClaimRequestDatasItemIndustry {
	s.MinorIndustryCodes = v
	return s
}

func (s *PoiClaimRequestDatasItemIndustry) SetQualifications(v []*PoiClaimRequestDatasItemIndustryQualificationsItem) *PoiClaimRequestDatasItemIndustry {
	s.Qualifications = v
	return s
}

type PoiClaimRequestDatasItemIndustryQualificationsItem struct {
	QualificationExpiration *string   `json:"qualification_expiration,omitempty" xml:"qualification_expiration,omitempty"`
	QualificationId         *string   `json:"qualification_id,omitempty" xml:"qualification_id,omitempty"`
	QualificationType       *int64    `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
	QualificationUrls       []*string `json:"qualification_urls,omitempty" xml:"qualification_urls,omitempty" type:"Repeated"`
}

func (s PoiClaimRequestDatasItemIndustryQualificationsItem) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemIndustryQualificationsItem) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemIndustryQualificationsItem) SetQualificationExpiration(v string) *PoiClaimRequestDatasItemIndustryQualificationsItem {
	s.QualificationExpiration = &v
	return s
}

func (s *PoiClaimRequestDatasItemIndustryQualificationsItem) SetQualificationId(v string) *PoiClaimRequestDatasItemIndustryQualificationsItem {
	s.QualificationId = &v
	return s
}

func (s *PoiClaimRequestDatasItemIndustryQualificationsItem) SetQualificationType(v int64) *PoiClaimRequestDatasItemIndustryQualificationsItem {
	s.QualificationType = &v
	return s
}

func (s *PoiClaimRequestDatasItemIndustryQualificationsItem) SetQualificationUrls(v []*string) *PoiClaimRequestDatasItemIndustryQualificationsItem {
	s.QualificationUrls = v
	return s
}

type PoiClaimRequestDatasItemLegalPerson struct {
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	QualificationType *int64  `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
	UseOcr            *bool   `json:"use_ocr,omitempty" xml:"use_ocr,omitempty"`
	IdCardBackUrl     *string `json:"id_card_back_url,omitempty" xml:"id_card_back_url,omitempty"`
	IdCardExpiration  *string `json:"id_card_expiration,omitempty" xml:"id_card_expiration,omitempty"`
	IdCardFrontUrl    *string `json:"id_card_front_url,omitempty" xml:"id_card_front_url,omitempty"`
	IdCardNo          *string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
}

func (s PoiClaimRequestDatasItemLegalPerson) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemLegalPerson) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetName(v string) *PoiClaimRequestDatasItemLegalPerson {
	s.Name = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetQualificationType(v int64) *PoiClaimRequestDatasItemLegalPerson {
	s.QualificationType = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetUseOcr(v bool) *PoiClaimRequestDatasItemLegalPerson {
	s.UseOcr = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetIdCardBackUrl(v string) *PoiClaimRequestDatasItemLegalPerson {
	s.IdCardBackUrl = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetIdCardExpiration(v string) *PoiClaimRequestDatasItemLegalPerson {
	s.IdCardExpiration = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetIdCardFrontUrl(v string) *PoiClaimRequestDatasItemLegalPerson {
	s.IdCardFrontUrl = &v
	return s
}

func (s *PoiClaimRequestDatasItemLegalPerson) SetIdCardNo(v string) *PoiClaimRequestDatasItemLegalPerson {
	s.IdCardNo = &v
	return s
}

type PoiClaimRequestDatasItemLicense struct {
	Province        *string   `json:"province,omitempty" xml:"province,omitempty"`
	UseOcr          *bool     `json:"use_ocr,omitempty" xml:"use_ocr,omitempty"`
	LicenseType     *int64    `json:"license_type,omitempty" xml:"license_type,omitempty"`
	LicenseId       *string   `json:"license_id,omitempty" xml:"license_id,omitempty"`
	City            *string   `json:"city,omitempty" xml:"city,omitempty"`
	SalesRange      *string   `json:"sales_range,omitempty" xml:"sales_range,omitempty"`
	Expiration      *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	CompanyName     *string   `json:"company_name,omitempty" xml:"company_name,omitempty"`
	LicenseUrls     []*string `json:"license_urls,omitempty" xml:"license_urls,omitempty" type:"Repeated"`
	LegalPersonName *string   `json:"legal_person_name,omitempty" xml:"legal_person_name,omitempty"`
	Address         *string   `json:"address,omitempty" xml:"address,omitempty"`
}

func (s PoiClaimRequestDatasItemLicense) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemLicense) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemLicense) SetProvince(v string) *PoiClaimRequestDatasItemLicense {
	s.Province = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetUseOcr(v bool) *PoiClaimRequestDatasItemLicense {
	s.UseOcr = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetLicenseType(v int64) *PoiClaimRequestDatasItemLicense {
	s.LicenseType = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetLicenseId(v string) *PoiClaimRequestDatasItemLicense {
	s.LicenseId = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetCity(v string) *PoiClaimRequestDatasItemLicense {
	s.City = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetSalesRange(v string) *PoiClaimRequestDatasItemLicense {
	s.SalesRange = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetExpiration(v string) *PoiClaimRequestDatasItemLicense {
	s.Expiration = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetCompanyName(v string) *PoiClaimRequestDatasItemLicense {
	s.CompanyName = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetLicenseUrls(v []*string) *PoiClaimRequestDatasItemLicense {
	s.LicenseUrls = v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetLegalPersonName(v string) *PoiClaimRequestDatasItemLicense {
	s.LegalPersonName = &v
	return s
}

func (s *PoiClaimRequestDatasItemLicense) SetAddress(v string) *PoiClaimRequestDatasItemLicense {
	s.Address = &v
	return s
}

type PoiClaimRequestDatasItemOwner struct {
	Role  *string `json:"role,omitempty" xml:"role,omitempty"`
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s PoiClaimRequestDatasItemOwner) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemOwner) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemOwner) SetRole(v string) *PoiClaimRequestDatasItemOwner {
	s.Role = &v
	return s
}

func (s *PoiClaimRequestDatasItemOwner) SetEmail(v string) *PoiClaimRequestDatasItemOwner {
	s.Email = &v
	return s
}

func (s *PoiClaimRequestDatasItemOwner) SetName(v string) *PoiClaimRequestDatasItemOwner {
	s.Name = &v
	return s
}

func (s *PoiClaimRequestDatasItemOwner) SetPhone(v string) *PoiClaimRequestDatasItemOwner {
	s.Phone = &v
	return s
}

type PoiClaimRequestDatasItemPoiClaimAuthorization struct {
	Urls          []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	EffectiveTime *string   `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
}

func (s PoiClaimRequestDatasItemPoiClaimAuthorization) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimRequestDatasItemPoiClaimAuthorization) GoString() string {
	return s.String()
}

func (s *PoiClaimRequestDatasItemPoiClaimAuthorization) SetUrls(v []*string) *PoiClaimRequestDatasItemPoiClaimAuthorization {
	s.Urls = v
	return s
}

func (s *PoiClaimRequestDatasItemPoiClaimAuthorization) SetEffectiveTime(v string) *PoiClaimRequestDatasItemPoiClaimAuthorization {
	s.EffectiveTime = &v
	return s
}

type PoiClaimResponse struct {
	Extra *PoiClaimResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PoiClaimResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoiClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimResponse) GoString() string {
	return s.String()
}

func (s *PoiClaimResponse) SetExtra(v *PoiClaimResponseExtra) *PoiClaimResponse {
	s.Extra = v
	return s
}

func (s *PoiClaimResponse) SetData(v *PoiClaimResponseData) *PoiClaimResponse {
	s.Data = v
	return s
}

type PoiClaimResponseData struct {
	Tasks         []*PoiClaimResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	GwErrorCode   *int32                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiClaimResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimResponseData) GoString() string {
	return s.String()
}

func (s *PoiClaimResponseData) SetTasks(v []*PoiClaimResponseDataTasksItem) *PoiClaimResponseData {
	s.Tasks = v
	return s
}

func (s *PoiClaimResponseData) SetGwErrorCode(v int32) *PoiClaimResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiClaimResponseData) SetGwDescription(v string) *PoiClaimResponseData {
	s.GwDescription = &v
	return s
}

type PoiClaimResponseDataTasksItem struct {
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId     *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TaskId    *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
	ThirdId   *string `json:"third_id,omitempty" xml:"third_id,omitempty"`
}

func (s PoiClaimResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *PoiClaimResponseDataTasksItem) SetAccountId(v string) *PoiClaimResponseDataTasksItem {
	s.AccountId = &v
	return s
}

func (s *PoiClaimResponseDataTasksItem) SetPoiId(v string) *PoiClaimResponseDataTasksItem {
	s.PoiId = &v
	return s
}

func (s *PoiClaimResponseDataTasksItem) SetTaskId(v int64) *PoiClaimResponseDataTasksItem {
	s.TaskId = &v
	return s
}

func (s *PoiClaimResponseDataTasksItem) SetThirdId(v string) *PoiClaimResponseDataTasksItem {
	s.ThirdId = &v
	return s
}

type PoiClaimResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiClaimResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiClaimResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiClaimResponseExtra) SetErrorCode(v int32) *PoiClaimResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiClaimResponseExtra) SetLogid(v string) *PoiClaimResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiClaimResponseExtra) SetNow(v int64) *PoiClaimResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiClaimResponseExtra) SetSubDescription(v string) *PoiClaimResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiClaimResponseExtra) SetSubErrorCode(v int32) *PoiClaimResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiClaimResponseExtra) SetDescription(v string) *PoiClaimResponseExtra {
	s.Description = &v
	return s
}

type PoiDecorateRequest struct {
	AccessToken *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Datas       []*PoiDecorateRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PoiDecorateRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequest) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequest) SetAccessToken(v string) *PoiDecorateRequest {
	s.AccessToken = &v
	return s
}

func (s *PoiDecorateRequest) SetDatas(v []*PoiDecorateRequestDatasItem) *PoiDecorateRequest {
	s.Datas = v
	return s
}

func (s *PoiDecorateRequest) SetHeader(v map[string]*string) *PoiDecorateRequest {
	s.Header = v
	return s
}

type PoiDecorateRequestDatasItem struct {
	Decoration *PoiDecorateRequestDatasItemDecoration `json:"decoration,omitempty" xml:"decoration,omitempty"`
	PoiId      *string                                `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ThirdId    *string                                `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId  *string                                `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s PoiDecorateRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItem) SetDecoration(v *PoiDecorateRequestDatasItemDecoration) *PoiDecorateRequestDatasItem {
	s.Decoration = v
	return s
}

func (s *PoiDecorateRequestDatasItem) SetPoiId(v string) *PoiDecorateRequestDatasItem {
	s.PoiId = &v
	return s
}

func (s *PoiDecorateRequestDatasItem) SetThirdId(v string) *PoiDecorateRequestDatasItem {
	s.ThirdId = &v
	return s
}

func (s *PoiDecorateRequestDatasItem) SetAccountId(v string) *PoiDecorateRequestDatasItem {
	s.AccountId = &v
	return s
}

type PoiDecorateRequestDatasItemDecoration struct {
	Tags         []*string                                               `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Announcement *PoiDecorateRequestDatasItemDecorationAnnouncement      `json:"announcement,omitempty" xml:"announcement,omitempty"`
	CoverImages  []*PoiDecorateRequestDatasItemDecorationCoverImagesItem `json:"cover_images,omitempty" xml:"cover_images,omitempty" type:"Repeated"`
	HeadImages   []*PoiDecorateRequestDatasItemDecorationHeadImagesItem  `json:"head_images,omitempty" xml:"head_images,omitempty" type:"Repeated"`
	Recommends   []*PoiDecorateRequestDatasItemDecorationRecommendsItem  `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
}

func (s PoiDecorateRequestDatasItemDecoration) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecoration) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecoration) SetTags(v []*string) *PoiDecorateRequestDatasItemDecoration {
	s.Tags = v
	return s
}

func (s *PoiDecorateRequestDatasItemDecoration) SetAnnouncement(v *PoiDecorateRequestDatasItemDecorationAnnouncement) *PoiDecorateRequestDatasItemDecoration {
	s.Announcement = v
	return s
}

func (s *PoiDecorateRequestDatasItemDecoration) SetCoverImages(v []*PoiDecorateRequestDatasItemDecorationCoverImagesItem) *PoiDecorateRequestDatasItemDecoration {
	s.CoverImages = v
	return s
}

func (s *PoiDecorateRequestDatasItemDecoration) SetHeadImages(v []*PoiDecorateRequestDatasItemDecorationHeadImagesItem) *PoiDecorateRequestDatasItemDecoration {
	s.HeadImages = v
	return s
}

func (s *PoiDecorateRequestDatasItemDecoration) SetRecommends(v []*PoiDecorateRequestDatasItemDecorationRecommendsItem) *PoiDecorateRequestDatasItemDecoration {
	s.Recommends = v
	return s
}

type PoiDecorateRequestDatasItemDecorationAnnouncement struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Enable      *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PoiDecorateRequestDatasItemDecorationAnnouncement) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecorationAnnouncement) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecorationAnnouncement) SetDescription(v string) *PoiDecorateRequestDatasItemDecorationAnnouncement {
	s.Description = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationAnnouncement) SetEnable(v bool) *PoiDecorateRequestDatasItemDecorationAnnouncement {
	s.Enable = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationAnnouncement) SetTitle(v string) *PoiDecorateRequestDatasItemDecorationAnnouncement {
	s.Title = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationAnnouncement) SetUrl(v string) *PoiDecorateRequestDatasItemDecorationAnnouncement {
	s.Url = &v
	return s
}

type PoiDecorateRequestDatasItemDecorationCoverImagesItem struct {
	SortValue *int64  `json:"sort_value,omitempty" xml:"sort_value,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PoiDecorateRequestDatasItemDecorationCoverImagesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecorationCoverImagesItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecorationCoverImagesItem) SetSortValue(v int64) *PoiDecorateRequestDatasItemDecorationCoverImagesItem {
	s.SortValue = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationCoverImagesItem) SetUrl(v string) *PoiDecorateRequestDatasItemDecorationCoverImagesItem {
	s.Url = &v
	return s
}

type PoiDecorateRequestDatasItemDecorationHeadImagesItem struct {
	SortValue *int64  `json:"sort_value,omitempty" xml:"sort_value,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PoiDecorateRequestDatasItemDecorationHeadImagesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecorationHeadImagesItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecorationHeadImagesItem) SetSortValue(v int64) *PoiDecorateRequestDatasItemDecorationHeadImagesItem {
	s.SortValue = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationHeadImagesItem) SetUrl(v string) *PoiDecorateRequestDatasItemDecorationHeadImagesItem {
	s.Url = &v
	return s
}

type PoiDecorateRequestDatasItemDecorationRecommendsItem struct {
	Items []*PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Title *string                                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PoiDecorateRequestDatasItemDecorationRecommendsItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecorationRecommendsItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecorationRecommendsItem) SetItems(v []*PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) *PoiDecorateRequestDatasItemDecorationRecommendsItem {
	s.Items = v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationRecommendsItem) SetTitle(v string) *PoiDecorateRequestDatasItemDecorationRecommendsItem {
	s.Title = &v
	return s
}

type PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem struct {
	SortValue *int64  `json:"sort_value,omitempty" xml:"sort_value,omitempty"`
	ImageName *string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	ImageUrl  *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

func (s PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) SetSortValue(v int64) *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem {
	s.SortValue = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) SetImageName(v string) *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem {
	s.ImageName = &v
	return s
}

func (s *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem) SetImageUrl(v string) *PoiDecorateRequestDatasItemDecorationRecommendsItemItemsItem {
	s.ImageUrl = &v
	return s
}

type PoiDecorateResponse struct {
	Extra *PoiDecorateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PoiDecorateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoiDecorateResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateResponse) GoString() string {
	return s.String()
}

func (s *PoiDecorateResponse) SetExtra(v *PoiDecorateResponseExtra) *PoiDecorateResponse {
	s.Extra = v
	return s
}

func (s *PoiDecorateResponse) SetData(v *PoiDecorateResponseData) *PoiDecorateResponse {
	s.Data = v
	return s
}

type PoiDecorateResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Tasks         []*PoiDecorateResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s PoiDecorateResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateResponseData) GoString() string {
	return s.String()
}

func (s *PoiDecorateResponseData) SetGwErrorCode(v int32) *PoiDecorateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiDecorateResponseData) SetGwDescription(v string) *PoiDecorateResponseData {
	s.GwDescription = &v
	return s
}

func (s *PoiDecorateResponseData) SetTasks(v []*PoiDecorateResponseDataTasksItem) *PoiDecorateResponseData {
	s.Tasks = v
	return s
}

type PoiDecorateResponseDataTasksItem struct {
	PoiId     *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TaskId    *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
	ThirdId   *string `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s PoiDecorateResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *PoiDecorateResponseDataTasksItem) SetPoiId(v string) *PoiDecorateResponseDataTasksItem {
	s.PoiId = &v
	return s
}

func (s *PoiDecorateResponseDataTasksItem) SetTaskId(v int64) *PoiDecorateResponseDataTasksItem {
	s.TaskId = &v
	return s
}

func (s *PoiDecorateResponseDataTasksItem) SetThirdId(v string) *PoiDecorateResponseDataTasksItem {
	s.ThirdId = &v
	return s
}

func (s *PoiDecorateResponseDataTasksItem) SetAccountId(v string) *PoiDecorateResponseDataTasksItem {
	s.AccountId = &v
	return s
}

type PoiDecorateResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiDecorateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiDecorateResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiDecorateResponseExtra) SetErrorCode(v int32) *PoiDecorateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiDecorateResponseExtra) SetLogid(v string) *PoiDecorateResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiDecorateResponseExtra) SetNow(v int64) *PoiDecorateResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiDecorateResponseExtra) SetSubDescription(v string) *PoiDecorateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiDecorateResponseExtra) SetSubErrorCode(v int32) *PoiDecorateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiDecorateResponseExtra) SetDescription(v string) *PoiDecorateResponseExtra {
	s.Description = &v
	return s
}

type PoiOrientedPlanDetailRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PlanIdList  []*int64           `json:"plan_id_list,omitempty" xml:"plan_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s PoiOrientedPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailRequest) SetHeader(v map[string]*string) *PoiOrientedPlanDetailRequest {
	s.Header = v
	return s
}

func (s *PoiOrientedPlanDetailRequest) SetAccessToken(v string) *PoiOrientedPlanDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *PoiOrientedPlanDetailRequest) SetPlanIdList(v []*int64) *PoiOrientedPlanDetailRequest {
	s.PlanIdList = v
	return s
}

type PoiOrientedPlanDetailResponse struct {
	Data   *PoiOrientedPlanDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s PoiOrientedPlanDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailResponse) SetData(v *PoiOrientedPlanDetailResponseData) *PoiOrientedPlanDetailResponse {
	s.Data = v
	return s
}

func (s *PoiOrientedPlanDetailResponse) SetErrMsg(v string) *PoiOrientedPlanDetailResponse {
	s.ErrMsg = &v
	return s
}

func (s *PoiOrientedPlanDetailResponse) SetErrNo(v int32) *PoiOrientedPlanDetailResponse {
	s.ErrNo = &v
	return s
}

func (s *PoiOrientedPlanDetailResponse) SetLogId(v string) *PoiOrientedPlanDetailResponse {
	s.LogId = &v
	return s
}

type PoiOrientedPlanDetailResponseData struct {
	Date *string                                                `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Data map[string]*PoiOrientedPlanDetailResponseDataDataValue `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s PoiOrientedPlanDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailResponseData) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailResponseData) SetDate(v string) *PoiOrientedPlanDetailResponseData {
	s.Date = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseData) SetData(v map[string]*PoiOrientedPlanDetailResponseDataDataValue) *PoiOrientedPlanDetailResponseData {
	s.Data = v
	return s
}

type PoiOrientedPlanDetailResponseDataDataValue struct {
	MediaCnt *int32                                              `json:"media_cnt,omitempty" xml:"media_cnt,omitempty" require:"true"`
	PlanInfo *PoiOrientedPlanDetailResponseDataDataValuePlanInfo `json:"plan_info,omitempty" xml:"plan_info,omitempty" require:"true"`
	UsedGmv  *int64                                              `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	Gmv      *int64                                              `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
}

func (s PoiOrientedPlanDetailResponseDataDataValue) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailResponseDataDataValue) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailResponseDataDataValue) SetMediaCnt(v int32) *PoiOrientedPlanDetailResponseDataDataValue {
	s.MediaCnt = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValue) SetPlanInfo(v *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) *PoiOrientedPlanDetailResponseDataDataValue {
	s.PlanInfo = v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValue) SetUsedGmv(v int64) *PoiOrientedPlanDetailResponseDataDataValue {
	s.UsedGmv = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValue) SetGmv(v int64) *PoiOrientedPlanDetailResponseDataDataValue {
	s.Gmv = &v
	return s
}

type PoiOrientedPlanDetailResponseDataDataValuePlanInfo struct {
	PlanId             *int64                                                               `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	DouyinIdList       []*string                                                            `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" require:"true" type:"Repeated"`
	ProductList        []*PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem `json:"product_list,omitempty" xml:"product_list,omitempty" require:"true" type:"Repeated"`
	CreateTime         *string                                                              `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	StartTime          *int64                                                               `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	PlanName           *string                                                              `json:"plan_name,omitempty" xml:"plan_name,omitempty" require:"true"`
	Status             *int32                                                               `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	TalentStatusMap    map[string]*int32                                                    `json:"talent_status_map,omitempty" xml:"talent_status_map,omitempty"`
	CommissionDuration *int64                                                               `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	UpdateTime         *string                                                              `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
	ContentType        *int32                                                               `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	EndTime            *int64                                                               `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s PoiOrientedPlanDetailResponseDataDataValuePlanInfo) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailResponseDataDataValuePlanInfo) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetPlanId(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.PlanId = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetDouyinIdList(v []*string) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.DouyinIdList = v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetProductList(v []*PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.ProductList = v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetCreateTime(v string) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.CreateTime = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetStartTime(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.StartTime = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetPlanName(v string) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.PlanName = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetStatus(v int32) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.Status = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetTalentStatusMap(v map[string]*int32) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.TalentStatusMap = v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetCommissionDuration(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.CommissionDuration = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetUpdateTime(v string) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.UpdateTime = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetContentType(v int32) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.ContentType = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfo) SetEndTime(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfo {
	s.EndTime = &v
	return s
}

type PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem struct {
	CommissionRate *int64 `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	ProductId      *int64 `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem) SetCommissionRate(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem {
	s.CommissionRate = &v
	return s
}

func (s *PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem) SetProductId(v int64) *PoiOrientedPlanDetailResponseDataDataValuePlanInfoProductListItem {
	s.ProductId = &v
	return s
}

type PoiOrientedPlanListRequest struct {
	SpuIdList   []*int64           `json:"spu_id_list,omitempty" xml:"spu_id_list,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiOrientedPlanListRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanListRequest) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanListRequest) SetSpuIdList(v []*int64) *PoiOrientedPlanListRequest {
	s.SpuIdList = v
	return s
}

func (s *PoiOrientedPlanListRequest) SetHeader(v map[string]*string) *PoiOrientedPlanListRequest {
	s.Header = v
	return s
}

func (s *PoiOrientedPlanListRequest) SetAccessToken(v string) *PoiOrientedPlanListRequest {
	s.AccessToken = &v
	return s
}

type PoiOrientedPlanListResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *PoiOrientedPlanListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s PoiOrientedPlanListResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanListResponse) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanListResponse) SetLogId(v string) *PoiOrientedPlanListResponse {
	s.LogId = &v
	return s
}

func (s *PoiOrientedPlanListResponse) SetData(v *PoiOrientedPlanListResponseData) *PoiOrientedPlanListResponse {
	s.Data = v
	return s
}

func (s *PoiOrientedPlanListResponse) SetErrMsg(v string) *PoiOrientedPlanListResponse {
	s.ErrMsg = &v
	return s
}

func (s *PoiOrientedPlanListResponse) SetErrNo(v int32) *PoiOrientedPlanListResponse {
	s.ErrNo = &v
	return s
}

type PoiOrientedPlanListResponseData struct {
	Data       []*PoiOrientedPlanListResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	PageCount  *int64                                     `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
	TotalCount *int64                                     `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
}

func (s PoiOrientedPlanListResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanListResponseData) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanListResponseData) SetData(v []*PoiOrientedPlanListResponseDataDataItem) *PoiOrientedPlanListResponseData {
	s.Data = v
	return s
}

func (s *PoiOrientedPlanListResponseData) SetPageCount(v int64) *PoiOrientedPlanListResponseData {
	s.PageCount = &v
	return s
}

func (s *PoiOrientedPlanListResponseData) SetTotalCount(v int64) *PoiOrientedPlanListResponseData {
	s.TotalCount = &v
	return s
}

type PoiOrientedPlanListResponseDataDataItem struct {
	ContentType        *int32                                                    `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	CommissionDuration *int64                                                    `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	PlanId             *int64                                                    `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	Status             *int32                                                    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	EndTime            *int64                                                    `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	UpdateTime         *string                                                   `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
	ProductList        []*PoiOrientedPlanListResponseDataDataItemProductListItem `json:"product_list,omitempty" xml:"product_list,omitempty" require:"true" type:"Repeated"`
	DouyinIdList       []*string                                                 `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" require:"true" type:"Repeated"`
	TalentStatusMap    map[string]*int32                                         `json:"talent_status_map,omitempty" xml:"talent_status_map,omitempty"`
	StartTime          *int64                                                    `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	CreateTimeUnix     *int64                                                    `json:"create_time_unix,omitempty" xml:"create_time_unix,omitempty"`
	PlanName           *string                                                   `json:"plan_name,omitempty" xml:"plan_name,omitempty" require:"true"`
	CreateTime         *string                                                   `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	UpdateTimeUnix     *int64                                                    `json:"update_time_unix,omitempty" xml:"update_time_unix,omitempty"`
}

func (s PoiOrientedPlanListResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanListResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetContentType(v int32) *PoiOrientedPlanListResponseDataDataItem {
	s.ContentType = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetCommissionDuration(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.CommissionDuration = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetPlanId(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.PlanId = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetStatus(v int32) *PoiOrientedPlanListResponseDataDataItem {
	s.Status = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetEndTime(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.EndTime = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetUpdateTime(v string) *PoiOrientedPlanListResponseDataDataItem {
	s.UpdateTime = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetProductList(v []*PoiOrientedPlanListResponseDataDataItemProductListItem) *PoiOrientedPlanListResponseDataDataItem {
	s.ProductList = v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetDouyinIdList(v []*string) *PoiOrientedPlanListResponseDataDataItem {
	s.DouyinIdList = v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetTalentStatusMap(v map[string]*int32) *PoiOrientedPlanListResponseDataDataItem {
	s.TalentStatusMap = v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetStartTime(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.StartTime = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetCreateTimeUnix(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.CreateTimeUnix = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetPlanName(v string) *PoiOrientedPlanListResponseDataDataItem {
	s.PlanName = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetCreateTime(v string) *PoiOrientedPlanListResponseDataDataItem {
	s.CreateTime = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItem) SetUpdateTimeUnix(v int64) *PoiOrientedPlanListResponseDataDataItem {
	s.UpdateTimeUnix = &v
	return s
}

type PoiOrientedPlanListResponseDataDataItemProductListItem struct {
	CommissionRate *int64 `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	ProductId      *int64 `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s PoiOrientedPlanListResponseDataDataItemProductListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiOrientedPlanListResponseDataDataItemProductListItem) GoString() string {
	return s.String()
}

func (s *PoiOrientedPlanListResponseDataDataItemProductListItem) SetCommissionRate(v int64) *PoiOrientedPlanListResponseDataDataItemProductListItem {
	s.CommissionRate = &v
	return s
}

func (s *PoiOrientedPlanListResponseDataDataItemProductListItem) SetProductId(v int64) *PoiOrientedPlanListResponseDataDataItemProductListItem {
	s.ProductId = &v
	return s
}

type PoiQueryRequest struct {
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiQueryRequest) GoString() string {
	return s.String()
}

func (s *PoiQueryRequest) SetAccountId(v string) *PoiQueryRequest {
	s.AccountId = &v
	return s
}

func (s *PoiQueryRequest) SetOrderId(v string) *PoiQueryRequest {
	s.OrderId = &v
	return s
}

func (s *PoiQueryRequest) SetHeader(v map[string]*string) *PoiQueryRequest {
	s.Header = v
	return s
}

func (s *PoiQueryRequest) SetAccessToken(v string) *PoiQueryRequest {
	s.AccessToken = &v
	return s
}

type PoiQueryResponse struct {
	Data  *PoiQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PoiQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoiQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiQueryResponse) GoString() string {
	return s.String()
}

func (s *PoiQueryResponse) SetData(v *PoiQueryResponseData) *PoiQueryResponse {
	s.Data = v
	return s
}

func (s *PoiQueryResponse) SetExtra(v *PoiQueryResponseExtra) *PoiQueryResponse {
	s.Extra = v
	return s
}

type PoiQueryResponseData struct {
	PoiList       []*PoiQueryResponseDataPoiListItem `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiQueryResponseData) GoString() string {
	return s.String()
}

func (s *PoiQueryResponseData) SetPoiList(v []*PoiQueryResponseDataPoiListItem) *PoiQueryResponseData {
	s.PoiList = v
	return s
}

func (s *PoiQueryResponseData) SetGwErrorCode(v int32) *PoiQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiQueryResponseData) SetGwDescription(v string) *PoiQueryResponseData {
	s.GwDescription = &v
	return s
}

type PoiQueryResponseDataPoiListItem struct {
	PoiId        *string  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	CityCode     *string  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	ProvinceName *string  `json:"province_name,omitempty" xml:"province_name,omitempty"`
	Latitude     *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	CountryCode  *string  `json:"country_code,omitempty" xml:"country_code,omitempty"`
	Address      *string  `json:"address,omitempty" xml:"address,omitempty"`
	ProvinceCode *string  `json:"province_code,omitempty" xml:"province_code,omitempty"`
	CountryName  *string  `json:"country_name,omitempty" xml:"country_name,omitempty"`
	PoiName      *string  `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	Longitude    *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
	CityName     *string  `json:"city_name,omitempty" xml:"city_name,omitempty"`
}

func (s PoiQueryResponseDataPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiQueryResponseDataPoiListItem) GoString() string {
	return s.String()
}

func (s *PoiQueryResponseDataPoiListItem) SetPoiId(v string) *PoiQueryResponseDataPoiListItem {
	s.PoiId = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetCityCode(v string) *PoiQueryResponseDataPoiListItem {
	s.CityCode = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetProvinceName(v string) *PoiQueryResponseDataPoiListItem {
	s.ProvinceName = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetLatitude(v float64) *PoiQueryResponseDataPoiListItem {
	s.Latitude = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetCountryCode(v string) *PoiQueryResponseDataPoiListItem {
	s.CountryCode = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetAddress(v string) *PoiQueryResponseDataPoiListItem {
	s.Address = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetProvinceCode(v string) *PoiQueryResponseDataPoiListItem {
	s.ProvinceCode = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetCountryName(v string) *PoiQueryResponseDataPoiListItem {
	s.CountryName = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetPoiName(v string) *PoiQueryResponseDataPoiListItem {
	s.PoiName = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetLongitude(v float64) *PoiQueryResponseDataPoiListItem {
	s.Longitude = &v
	return s
}

func (s *PoiQueryResponseDataPoiListItem) SetCityName(v string) *PoiQueryResponseDataPoiListItem {
	s.CityName = &v
	return s
}

type PoiQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoiQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiQueryResponseExtra) SetSubErrorCode(v int32) *PoiQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiQueryResponseExtra) SetDescription(v string) *PoiQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiQueryResponseExtra) SetErrorCode(v int32) *PoiQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiQueryResponseExtra) SetLogid(v string) *PoiQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiQueryResponseExtra) SetNow(v int64) *PoiQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiQueryResponseExtra) SetSubDescription(v string) *PoiQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type PoiSaveCommonPlanRequest struct {
	PlanId         *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	SpuId          *int64             `json:"spu_id,omitempty" xml:"spu_id,omitempty" require:"true"`
	CommissionRate *int64             `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	ContentType    *int32             `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiSaveCommonPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiSaveCommonPlanRequest) GoString() string {
	return s.String()
}

func (s *PoiSaveCommonPlanRequest) SetPlanId(v int64) *PoiSaveCommonPlanRequest {
	s.PlanId = &v
	return s
}

func (s *PoiSaveCommonPlanRequest) SetSpuId(v int64) *PoiSaveCommonPlanRequest {
	s.SpuId = &v
	return s
}

func (s *PoiSaveCommonPlanRequest) SetCommissionRate(v int64) *PoiSaveCommonPlanRequest {
	s.CommissionRate = &v
	return s
}

func (s *PoiSaveCommonPlanRequest) SetContentType(v int32) *PoiSaveCommonPlanRequest {
	s.ContentType = &v
	return s
}

func (s *PoiSaveCommonPlanRequest) SetHeader(v map[string]*string) *PoiSaveCommonPlanRequest {
	s.Header = v
	return s
}

func (s *PoiSaveCommonPlanRequest) SetAccessToken(v string) *PoiSaveCommonPlanRequest {
	s.AccessToken = &v
	return s
}

type PoiSaveCommonPlanResponse struct {
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *PoiSaveCommonPlanResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s PoiSaveCommonPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiSaveCommonPlanResponse) GoString() string {
	return s.String()
}

func (s *PoiSaveCommonPlanResponse) SetLogId(v string) *PoiSaveCommonPlanResponse {
	s.LogId = &v
	return s
}

func (s *PoiSaveCommonPlanResponse) SetData(v *PoiSaveCommonPlanResponseData) *PoiSaveCommonPlanResponse {
	s.Data = v
	return s
}

func (s *PoiSaveCommonPlanResponse) SetErrMsg(v string) *PoiSaveCommonPlanResponse {
	s.ErrMsg = &v
	return s
}

func (s *PoiSaveCommonPlanResponse) SetErrNo(v int32) *PoiSaveCommonPlanResponse {
	s.ErrNo = &v
	return s
}

type PoiSaveCommonPlanResponseData struct {
	PlanId *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s PoiSaveCommonPlanResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiSaveCommonPlanResponseData) GoString() string {
	return s.String()
}

func (s *PoiSaveCommonPlanResponseData) SetPlanId(v int64) *PoiSaveCommonPlanResponseData {
	s.PlanId = &v
	return s
}

type PoiScrollRequest struct {
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	ScrollId    *string            `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiScrollRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiScrollRequest) GoString() string {
	return s.String()
}

func (s *PoiScrollRequest) SetAccountId(v string) *PoiScrollRequest {
	s.AccountId = &v
	return s
}

func (s *PoiScrollRequest) SetPageSize(v int32) *PoiScrollRequest {
	s.PageSize = &v
	return s
}

func (s *PoiScrollRequest) SetScrollId(v string) *PoiScrollRequest {
	s.ScrollId = &v
	return s
}

func (s *PoiScrollRequest) SetHeader(v map[string]*string) *PoiScrollRequest {
	s.Header = v
	return s
}

func (s *PoiScrollRequest) SetAccessToken(v string) *PoiScrollRequest {
	s.AccessToken = &v
	return s
}

type PoiScrollResponse struct {
	Data  *string                 `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PoiScrollResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoiScrollResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiScrollResponse) GoString() string {
	return s.String()
}

func (s *PoiScrollResponse) SetData(v string) *PoiScrollResponse {
	s.Data = &v
	return s
}

func (s *PoiScrollResponse) SetExtra(v *PoiScrollResponseExtra) *PoiScrollResponse {
	s.Extra = v
	return s
}

type PoiScrollResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s PoiScrollResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiScrollResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiScrollResponseExtra) SetDescription(v string) *PoiScrollResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiScrollResponseExtra) SetErrorCode(v int32) *PoiScrollResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiScrollResponseExtra) SetLogid(v string) *PoiScrollResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiScrollResponseExtra) SetNow(v int64) *PoiScrollResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiScrollResponseExtra) SetSubDescription(v string) *PoiScrollResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiScrollResponseExtra) SetSubErrorCode(v int32) *PoiScrollResponseExtra {
	s.SubErrorCode = &v
	return s
}

type PoiSellOutDetailRequest struct {
	OutAffiliatedId *string                      `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	PoiIds          []*int64                     `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" require:"true" type:"Repeated"`
	Header          map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base            *PoiSellOutDetailRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *int64                       `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AffiliatedId    *int64                       `json:"affiliated_id,omitempty" xml:"affiliated_id,omitempty"`
}

func (s PoiSellOutDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailRequest) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailRequest) SetOutAffiliatedId(v string) *PoiSellOutDetailRequest {
	s.OutAffiliatedId = &v
	return s
}

func (s *PoiSellOutDetailRequest) SetPoiIds(v []*int64) *PoiSellOutDetailRequest {
	s.PoiIds = v
	return s
}

func (s *PoiSellOutDetailRequest) SetHeader(v map[string]*string) *PoiSellOutDetailRequest {
	s.Header = v
	return s
}

func (s *PoiSellOutDetailRequest) SetAccessToken(v string) *PoiSellOutDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *PoiSellOutDetailRequest) SetBase(v *PoiSellOutDetailRequestBase) *PoiSellOutDetailRequest {
	s.Base = v
	return s
}

func (s *PoiSellOutDetailRequest) SetAccountId(v int64) *PoiSellOutDetailRequest {
	s.AccountId = &v
	return s
}

func (s *PoiSellOutDetailRequest) SetAffiliatedId(v int64) *PoiSellOutDetailRequest {
	s.AffiliatedId = &v
	return s
}

type PoiSellOutDetailRequestBase struct {
	Addr       *string                                `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                     `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *PoiSellOutDetailRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s PoiSellOutDetailRequestBase) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailRequestBase) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailRequestBase) SetAddr(v string) *PoiSellOutDetailRequestBase {
	s.Addr = &v
	return s
}

func (s *PoiSellOutDetailRequestBase) SetCaller(v string) *PoiSellOutDetailRequestBase {
	s.Caller = &v
	return s
}

func (s *PoiSellOutDetailRequestBase) SetClient(v string) *PoiSellOutDetailRequestBase {
	s.Client = &v
	return s
}

func (s *PoiSellOutDetailRequestBase) SetExtra(v map[string]*string) *PoiSellOutDetailRequestBase {
	s.Extra = v
	return s
}

func (s *PoiSellOutDetailRequestBase) SetLogID(v string) *PoiSellOutDetailRequestBase {
	s.LogID = &v
	return s
}

func (s *PoiSellOutDetailRequestBase) SetTrafficEnv(v *PoiSellOutDetailRequestBaseTrafficEnv) *PoiSellOutDetailRequestBase {
	s.TrafficEnv = v
	return s
}

type PoiSellOutDetailRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s PoiSellOutDetailRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailRequestBaseTrafficEnv) SetEnv(v string) *PoiSellOutDetailRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *PoiSellOutDetailRequestBaseTrafficEnv) SetOpen(v bool) *PoiSellOutDetailRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type PoiSellOutDetailResponse struct {
	BaseResp *PoiSellOutDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *PoiSellOutDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoiSellOutDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoiSellOutDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponse) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponse) SetBaseResp(v *PoiSellOutDetailResponseBaseResp) *PoiSellOutDetailResponse {
	s.BaseResp = v
	return s
}

func (s *PoiSellOutDetailResponse) SetData(v *PoiSellOutDetailResponseData) *PoiSellOutDetailResponse {
	s.Data = v
	return s
}

func (s *PoiSellOutDetailResponse) SetExtra(v *PoiSellOutDetailResponseExtra) *PoiSellOutDetailResponse {
	s.Extra = v
	return s
}

type PoiSellOutDetailResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s PoiSellOutDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponseBaseResp) SetStatusCode(v int32) *PoiSellOutDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoiSellOutDetailResponseBaseResp) SetStatusMessage(v string) *PoiSellOutDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoiSellOutDetailResponseBaseResp) SetExtra(v map[string]*string) *PoiSellOutDetailResponseBaseResp {
	s.Extra = v
	return s
}

type PoiSellOutDetailResponseData struct {
	AffiliatedPoiSellOutMap map[int64]*PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue `json:"affiliated_poi_sell_out_map,omitempty" xml:"affiliated_poi_sell_out_map,omitempty"`
	Description             *string                                                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode               *int32                                                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PoiSellOutDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponseData) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponseData) SetAffiliatedPoiSellOutMap(v map[int64]*PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue) *PoiSellOutDetailResponseData {
	s.AffiliatedPoiSellOutMap = v
	return s
}

func (s *PoiSellOutDetailResponseData) SetDescription(v string) *PoiSellOutDetailResponseData {
	s.Description = &v
	return s
}

func (s *PoiSellOutDetailResponseData) SetErrorCode(v int32) *PoiSellOutDetailResponseData {
	s.ErrorCode = &v
	return s
}

type PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue struct {
	SellOutRule *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule `json:"sell_out_rule,omitempty" xml:"sell_out_rule,omitempty"`
	PoiId       *string                                                              `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue) SetSellOutRule(v *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue {
	s.SellOutRule = v
	return s
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue) SetPoiId(v string) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValue {
	s.PoiId = &v
	return s
}

type PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule struct {
	LongTermSellOutStatus *int    `json:"long_term_sell_out_status,omitempty" xml:"long_term_sell_out_status,omitempty"`
	SellOutType           *int    `json:"sell_out_type,omitempty" xml:"sell_out_type,omitempty"`
	StartTime             *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime               *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) SetLongTermSellOutStatus(v int) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule {
	s.LongTermSellOutStatus = &v
	return s
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) SetSellOutType(v int) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule {
	s.SellOutType = &v
	return s
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) SetStartTime(v string) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule {
	s.StartTime = &v
	return s
}

func (s *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule) SetEndTime(v string) *PoiSellOutDetailResponseDataAffiliatedPoiSellOutMapValueSellOutRule {
	s.EndTime = &v
	return s
}

type PoiSellOutDetailResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s PoiSellOutDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiSellOutDetailResponseExtra) SetNow(v int64) *PoiSellOutDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiSellOutDetailResponseExtra) SetSubDescription(v string) *PoiSellOutDetailResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiSellOutDetailResponseExtra) SetSubErrorCode(v int32) *PoiSellOutDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiSellOutDetailResponseExtra) SetDescription(v string) *PoiSellOutDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiSellOutDetailResponseExtra) SetErrorCode(v int32) *PoiSellOutDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiSellOutDetailResponseExtra) SetLogid(v string) *PoiSellOutDetailResponseExtra {
	s.Logid = &v
	return s
}

type PoiSellOutSaveRequest struct {
	AccountId      *string                                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiSellOutRule []*PoiSellOutSaveRequestPoiSellOutRuleItem `json:"poi_sell_out_rule,omitempty" xml:"poi_sell_out_rule,omitempty" require:"true" type:"Repeated"`
	Header         map[string]*string                         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string                                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductId      *string                                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Base           *PoiSellOutSaveRequestBase                 `json:"Base,omitempty" xml:"Base,omitempty"`
}

func (s PoiSellOutSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveRequest) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveRequest) SetAccountId(v string) *PoiSellOutSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PoiSellOutSaveRequest) SetPoiSellOutRule(v []*PoiSellOutSaveRequestPoiSellOutRuleItem) *PoiSellOutSaveRequest {
	s.PoiSellOutRule = v
	return s
}

func (s *PoiSellOutSaveRequest) SetHeader(v map[string]*string) *PoiSellOutSaveRequest {
	s.Header = v
	return s
}

func (s *PoiSellOutSaveRequest) SetAccessToken(v string) *PoiSellOutSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *PoiSellOutSaveRequest) SetProductId(v string) *PoiSellOutSaveRequest {
	s.ProductId = &v
	return s
}

func (s *PoiSellOutSaveRequest) SetBase(v *PoiSellOutSaveRequestBase) *PoiSellOutSaveRequest {
	s.Base = v
	return s
}

type PoiSellOutSaveRequestBase struct {
	TrafficEnv *PoiSellOutSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                              `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                              `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                              `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                   `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                              `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s PoiSellOutSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveRequestBase) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveRequestBase) SetTrafficEnv(v *PoiSellOutSaveRequestBaseTrafficEnv) *PoiSellOutSaveRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *PoiSellOutSaveRequestBase) SetAddr(v string) *PoiSellOutSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *PoiSellOutSaveRequestBase) SetCaller(v string) *PoiSellOutSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *PoiSellOutSaveRequestBase) SetClient(v string) *PoiSellOutSaveRequestBase {
	s.Client = &v
	return s
}

func (s *PoiSellOutSaveRequestBase) SetExtra(v map[string]*string) *PoiSellOutSaveRequestBase {
	s.Extra = v
	return s
}

func (s *PoiSellOutSaveRequestBase) SetLogID(v string) *PoiSellOutSaveRequestBase {
	s.LogID = &v
	return s
}

type PoiSellOutSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s PoiSellOutSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveRequestBaseTrafficEnv) SetEnv(v string) *PoiSellOutSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *PoiSellOutSaveRequestBaseTrafficEnv) SetOpen(v bool) *PoiSellOutSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type PoiSellOutSaveRequestPoiSellOutRuleItem struct {
	PoiId       *string                                             `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SellOutRule *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule `json:"sell_out_rule,omitempty" xml:"sell_out_rule,omitempty"`
}

func (s PoiSellOutSaveRequestPoiSellOutRuleItem) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveRequestPoiSellOutRuleItem) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItem) SetPoiId(v string) *PoiSellOutSaveRequestPoiSellOutRuleItem {
	s.PoiId = &v
	return s
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItem) SetSellOutRule(v *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) *PoiSellOutSaveRequestPoiSellOutRuleItem {
	s.SellOutRule = v
	return s
}

type PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule struct {
	StartTime             *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime               *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	LongTermSellOutStatus *int    `json:"long_term_sell_out_status,omitempty" xml:"long_term_sell_out_status,omitempty"`
	SellOutType           *int    `json:"sell_out_type,omitempty" xml:"sell_out_type,omitempty"`
}

func (s PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) SetStartTime(v string) *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule {
	s.StartTime = &v
	return s
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) SetEndTime(v string) *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule {
	s.EndTime = &v
	return s
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) SetLongTermSellOutStatus(v int) *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule {
	s.LongTermSellOutStatus = &v
	return s
}

func (s *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule) SetSellOutType(v int) *PoiSellOutSaveRequestPoiSellOutRuleItemSellOutRule {
	s.SellOutType = &v
	return s
}

type PoiSellOutSaveResponse struct {
	Extra    *PoiSellOutSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *PoiSellOutSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *PoiSellOutSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoiSellOutSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveResponse) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveResponse) SetExtra(v *PoiSellOutSaveResponseExtra) *PoiSellOutSaveResponse {
	s.Extra = v
	return s
}

func (s *PoiSellOutSaveResponse) SetBaseResp(v *PoiSellOutSaveResponseBaseResp) *PoiSellOutSaveResponse {
	s.BaseResp = v
	return s
}

func (s *PoiSellOutSaveResponse) SetData(v *PoiSellOutSaveResponseData) *PoiSellOutSaveResponse {
	s.Data = v
	return s
}

type PoiSellOutSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s PoiSellOutSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveResponseBaseResp) SetStatusCode(v int32) *PoiSellOutSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoiSellOutSaveResponseBaseResp) SetStatusMessage(v string) *PoiSellOutSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoiSellOutSaveResponseBaseResp) SetExtra(v map[string]*string) *PoiSellOutSaveResponseBaseResp {
	s.Extra = v
	return s
}

type PoiSellOutSaveResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiSellOutSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveResponseData) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveResponseData) SetErrorCode(v int32) *PoiSellOutSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PoiSellOutSaveResponseData) SetDescription(v string) *PoiSellOutSaveResponseData {
	s.Description = &v
	return s
}

type PoiSellOutSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoiSellOutSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiSellOutSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiSellOutSaveResponseExtra) SetSubErrorCode(v int32) *PoiSellOutSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiSellOutSaveResponseExtra) SetDescription(v string) *PoiSellOutSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiSellOutSaveResponseExtra) SetErrorCode(v int32) *PoiSellOutSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiSellOutSaveResponseExtra) SetLogid(v string) *PoiSellOutSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiSellOutSaveResponseExtra) SetNow(v int64) *PoiSellOutSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiSellOutSaveResponseExtra) SetSubDescription(v string) *PoiSellOutSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type PoiSyncRequest struct {
	Datas       []*PoiSyncRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" require:"true" type:"Repeated"`
	TargetType  *int                       `json:"target_type,omitempty" xml:"target_type,omitempty" require:"true"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequest) GoString() string {
	return s.String()
}

func (s *PoiSyncRequest) SetDatas(v []*PoiSyncRequestDatasItem) *PoiSyncRequest {
	s.Datas = v
	return s
}

func (s *PoiSyncRequest) SetTargetType(v int) *PoiSyncRequest {
	s.TargetType = &v
	return s
}

func (s *PoiSyncRequest) SetHeader(v map[string]*string) *PoiSyncRequest {
	s.Header = v
	return s
}

func (s *PoiSyncRequest) SetAccessToken(v string) *PoiSyncRequest {
	s.AccessToken = &v
	return s
}

type PoiSyncRequestDatasItem struct {
	ThirdId               *string                                       `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId             *string                                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Industry              *PoiSyncRequestDatasItemIndustry              `json:"industry,omitempty" xml:"industry,omitempty"`
	Owner                 *PoiSyncRequestDatasItemOwner                 `json:"owner,omitempty" xml:"owner,omitempty"`
	LegalPerson           *PoiSyncRequestDatasItemLegalPerson           `json:"legal_person,omitempty" xml:"legal_person,omitempty"`
	PoiId                 *string                                       `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	BusinessPartnership   *PoiSyncRequestDatasItemBusinessPartnership   `json:"business_partnership,omitempty" xml:"business_partnership,omitempty"`
	PoiClaimAuthorization *PoiSyncRequestDatasItemPoiClaimAuthorization `json:"poi_claim_authorization,omitempty" xml:"poi_claim_authorization,omitempty"`
	License               *PoiSyncRequestDatasItemLicense               `json:"license,omitempty" xml:"license,omitempty"`
}

func (s PoiSyncRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItem) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItem) SetThirdId(v string) *PoiSyncRequestDatasItem {
	s.ThirdId = &v
	return s
}

func (s *PoiSyncRequestDatasItem) SetAccountId(v string) *PoiSyncRequestDatasItem {
	s.AccountId = &v
	return s
}

func (s *PoiSyncRequestDatasItem) SetIndustry(v *PoiSyncRequestDatasItemIndustry) *PoiSyncRequestDatasItem {
	s.Industry = v
	return s
}

func (s *PoiSyncRequestDatasItem) SetOwner(v *PoiSyncRequestDatasItemOwner) *PoiSyncRequestDatasItem {
	s.Owner = v
	return s
}

func (s *PoiSyncRequestDatasItem) SetLegalPerson(v *PoiSyncRequestDatasItemLegalPerson) *PoiSyncRequestDatasItem {
	s.LegalPerson = v
	return s
}

func (s *PoiSyncRequestDatasItem) SetPoiId(v string) *PoiSyncRequestDatasItem {
	s.PoiId = &v
	return s
}

func (s *PoiSyncRequestDatasItem) SetBusinessPartnership(v *PoiSyncRequestDatasItemBusinessPartnership) *PoiSyncRequestDatasItem {
	s.BusinessPartnership = v
	return s
}

func (s *PoiSyncRequestDatasItem) SetPoiClaimAuthorization(v *PoiSyncRequestDatasItemPoiClaimAuthorization) *PoiSyncRequestDatasItem {
	s.PoiClaimAuthorization = v
	return s
}

func (s *PoiSyncRequestDatasItem) SetLicense(v *PoiSyncRequestDatasItemLicense) *PoiSyncRequestDatasItem {
	s.License = v
	return s
}

type PoiSyncRequestDatasItemBusinessPartnership struct {
	Authorization      *PoiSyncRequestDatasItemBusinessPartnershipAuthorization `json:"authorization,omitempty" xml:"authorization,omitempty"`
	PartnerAccountType *int                                                     `json:"partner_account_type,omitempty" xml:"partner_account_type,omitempty"`
}

func (s PoiSyncRequestDatasItemBusinessPartnership) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemBusinessPartnership) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemBusinessPartnership) SetAuthorization(v *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) *PoiSyncRequestDatasItemBusinessPartnership {
	s.Authorization = v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnership) SetPartnerAccountType(v int) *PoiSyncRequestDatasItemBusinessPartnership {
	s.PartnerAccountType = &v
	return s
}

type PoiSyncRequestDatasItemBusinessPartnershipAuthorization struct {
	EffectiveTime *string   `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	Expiration    *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	Id            *string   `json:"id,omitempty" xml:"id,omitempty"`
	Type          *int      `json:"type,omitempty" xml:"type,omitempty"`
	Urls          []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	Agreement     *int      `json:"agreement,omitempty" xml:"agreement,omitempty"`
	AppIds        []*string `json:"app_ids,omitempty" xml:"app_ids,omitempty" type:"Repeated"`
}

func (s PoiSyncRequestDatasItemBusinessPartnershipAuthorization) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemBusinessPartnershipAuthorization) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetEffectiveTime(v string) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.EffectiveTime = &v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetExpiration(v string) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.Expiration = &v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetId(v string) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.Id = &v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetType(v int) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.Type = &v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetUrls(v []*string) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.Urls = v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetAgreement(v int) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.Agreement = &v
	return s
}

func (s *PoiSyncRequestDatasItemBusinessPartnershipAuthorization) SetAppIds(v []*string) *PoiSyncRequestDatasItemBusinessPartnershipAuthorization {
	s.AppIds = v
	return s
}

type PoiSyncRequestDatasItemIndustry struct {
	MajorIndustryCode  *string                                              `json:"major_industry_code,omitempty" xml:"major_industry_code,omitempty"`
	MinorIndustryCodes []*string                                            `json:"minor_industry_codes,omitempty" xml:"minor_industry_codes,omitempty" type:"Repeated"`
	Qualifications     []*PoiSyncRequestDatasItemIndustryQualificationsItem `json:"qualifications,omitempty" xml:"qualifications,omitempty" type:"Repeated"`
}

func (s PoiSyncRequestDatasItemIndustry) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemIndustry) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemIndustry) SetMajorIndustryCode(v string) *PoiSyncRequestDatasItemIndustry {
	s.MajorIndustryCode = &v
	return s
}

func (s *PoiSyncRequestDatasItemIndustry) SetMinorIndustryCodes(v []*string) *PoiSyncRequestDatasItemIndustry {
	s.MinorIndustryCodes = v
	return s
}

func (s *PoiSyncRequestDatasItemIndustry) SetQualifications(v []*PoiSyncRequestDatasItemIndustryQualificationsItem) *PoiSyncRequestDatasItemIndustry {
	s.Qualifications = v
	return s
}

type PoiSyncRequestDatasItemIndustryQualificationsItem struct {
	QualificationType       *int64    `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
	QualificationUrls       []*string `json:"qualification_urls,omitempty" xml:"qualification_urls,omitempty" type:"Repeated"`
	QualificationExpiration *string   `json:"qualification_expiration,omitempty" xml:"qualification_expiration,omitempty"`
	QualificationId         *string   `json:"qualification_id,omitempty" xml:"qualification_id,omitempty"`
}

func (s PoiSyncRequestDatasItemIndustryQualificationsItem) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemIndustryQualificationsItem) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemIndustryQualificationsItem) SetQualificationType(v int64) *PoiSyncRequestDatasItemIndustryQualificationsItem {
	s.QualificationType = &v
	return s
}

func (s *PoiSyncRequestDatasItemIndustryQualificationsItem) SetQualificationUrls(v []*string) *PoiSyncRequestDatasItemIndustryQualificationsItem {
	s.QualificationUrls = v
	return s
}

func (s *PoiSyncRequestDatasItemIndustryQualificationsItem) SetQualificationExpiration(v string) *PoiSyncRequestDatasItemIndustryQualificationsItem {
	s.QualificationExpiration = &v
	return s
}

func (s *PoiSyncRequestDatasItemIndustryQualificationsItem) SetQualificationId(v string) *PoiSyncRequestDatasItemIndustryQualificationsItem {
	s.QualificationId = &v
	return s
}

type PoiSyncRequestDatasItemLegalPerson struct {
	IdCardFrontUrl    *string `json:"id_card_front_url,omitempty" xml:"id_card_front_url,omitempty"`
	IdCardNo          *string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	QualificationType *int64  `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
	UseOcr            *bool   `json:"use_ocr,omitempty" xml:"use_ocr,omitempty"`
	IdCardBackUrl     *string `json:"id_card_back_url,omitempty" xml:"id_card_back_url,omitempty"`
	IdCardExpiration  *string `json:"id_card_expiration,omitempty" xml:"id_card_expiration,omitempty"`
}

func (s PoiSyncRequestDatasItemLegalPerson) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemLegalPerson) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetIdCardFrontUrl(v string) *PoiSyncRequestDatasItemLegalPerson {
	s.IdCardFrontUrl = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetIdCardNo(v string) *PoiSyncRequestDatasItemLegalPerson {
	s.IdCardNo = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetName(v string) *PoiSyncRequestDatasItemLegalPerson {
	s.Name = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetQualificationType(v int64) *PoiSyncRequestDatasItemLegalPerson {
	s.QualificationType = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetUseOcr(v bool) *PoiSyncRequestDatasItemLegalPerson {
	s.UseOcr = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetIdCardBackUrl(v string) *PoiSyncRequestDatasItemLegalPerson {
	s.IdCardBackUrl = &v
	return s
}

func (s *PoiSyncRequestDatasItemLegalPerson) SetIdCardExpiration(v string) *PoiSyncRequestDatasItemLegalPerson {
	s.IdCardExpiration = &v
	return s
}

type PoiSyncRequestDatasItemLicense struct {
	SalesRange      *string   `json:"sales_range,omitempty" xml:"sales_range,omitempty"`
	LicenseId       *string   `json:"license_id,omitempty" xml:"license_id,omitempty"`
	LegalPersonName *string   `json:"legal_person_name,omitempty" xml:"legal_person_name,omitempty"`
	CompanyName     *string   `json:"company_name,omitempty" xml:"company_name,omitempty"`
	LicenseUrls     []*string `json:"license_urls,omitempty" xml:"license_urls,omitempty" type:"Repeated"`
	Province        *string   `json:"province,omitempty" xml:"province,omitempty"`
	UseOcr          *bool     `json:"use_ocr,omitempty" xml:"use_ocr,omitempty"`
	City            *string   `json:"city,omitempty" xml:"city,omitempty"`
	Expiration      *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	Address         *string   `json:"address,omitempty" xml:"address,omitempty"`
	LicenseType     *int64    `json:"license_type,omitempty" xml:"license_type,omitempty"`
}

func (s PoiSyncRequestDatasItemLicense) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemLicense) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemLicense) SetSalesRange(v string) *PoiSyncRequestDatasItemLicense {
	s.SalesRange = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetLicenseId(v string) *PoiSyncRequestDatasItemLicense {
	s.LicenseId = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetLegalPersonName(v string) *PoiSyncRequestDatasItemLicense {
	s.LegalPersonName = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetCompanyName(v string) *PoiSyncRequestDatasItemLicense {
	s.CompanyName = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetLicenseUrls(v []*string) *PoiSyncRequestDatasItemLicense {
	s.LicenseUrls = v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetProvince(v string) *PoiSyncRequestDatasItemLicense {
	s.Province = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetUseOcr(v bool) *PoiSyncRequestDatasItemLicense {
	s.UseOcr = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetCity(v string) *PoiSyncRequestDatasItemLicense {
	s.City = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetExpiration(v string) *PoiSyncRequestDatasItemLicense {
	s.Expiration = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetAddress(v string) *PoiSyncRequestDatasItemLicense {
	s.Address = &v
	return s
}

func (s *PoiSyncRequestDatasItemLicense) SetLicenseType(v int64) *PoiSyncRequestDatasItemLicense {
	s.LicenseType = &v
	return s
}

type PoiSyncRequestDatasItemOwner struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Role  *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s PoiSyncRequestDatasItemOwner) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemOwner) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemOwner) SetEmail(v string) *PoiSyncRequestDatasItemOwner {
	s.Email = &v
	return s
}

func (s *PoiSyncRequestDatasItemOwner) SetName(v string) *PoiSyncRequestDatasItemOwner {
	s.Name = &v
	return s
}

func (s *PoiSyncRequestDatasItemOwner) SetPhone(v string) *PoiSyncRequestDatasItemOwner {
	s.Phone = &v
	return s
}

func (s *PoiSyncRequestDatasItemOwner) SetRole(v string) *PoiSyncRequestDatasItemOwner {
	s.Role = &v
	return s
}

type PoiSyncRequestDatasItemPoiClaimAuthorization struct {
	EffectiveTime *string   `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	Urls          []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s PoiSyncRequestDatasItemPoiClaimAuthorization) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncRequestDatasItemPoiClaimAuthorization) GoString() string {
	return s.String()
}

func (s *PoiSyncRequestDatasItemPoiClaimAuthorization) SetEffectiveTime(v string) *PoiSyncRequestDatasItemPoiClaimAuthorization {
	s.EffectiveTime = &v
	return s
}

func (s *PoiSyncRequestDatasItemPoiClaimAuthorization) SetUrls(v []*string) *PoiSyncRequestDatasItemPoiClaimAuthorization {
	s.Urls = v
	return s
}

type PoiSyncResponse struct {
	Data  *PoiSyncResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PoiSyncResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoiSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncResponse) GoString() string {
	return s.String()
}

func (s *PoiSyncResponse) SetData(v *PoiSyncResponseData) *PoiSyncResponse {
	s.Data = v
	return s
}

func (s *PoiSyncResponse) SetExtra(v *PoiSyncResponseExtra) *PoiSyncResponse {
	s.Extra = v
	return s
}

type PoiSyncResponseData struct {
	Tasks         []*PoiSyncResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	GwErrorCode   *int32                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiSyncResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncResponseData) GoString() string {
	return s.String()
}

func (s *PoiSyncResponseData) SetTasks(v []*PoiSyncResponseDataTasksItem) *PoiSyncResponseData {
	s.Tasks = v
	return s
}

func (s *PoiSyncResponseData) SetGwErrorCode(v int32) *PoiSyncResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiSyncResponseData) SetGwDescription(v string) *PoiSyncResponseData {
	s.GwDescription = &v
	return s
}

type PoiSyncResponseDataTasksItem struct {
	ThirdId   *string `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId     *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TaskId    *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s PoiSyncResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *PoiSyncResponseDataTasksItem) SetThirdId(v string) *PoiSyncResponseDataTasksItem {
	s.ThirdId = &v
	return s
}

func (s *PoiSyncResponseDataTasksItem) SetAccountId(v string) *PoiSyncResponseDataTasksItem {
	s.AccountId = &v
	return s
}

func (s *PoiSyncResponseDataTasksItem) SetPoiId(v string) *PoiSyncResponseDataTasksItem {
	s.PoiId = &v
	return s
}

func (s *PoiSyncResponseDataTasksItem) SetTaskId(v int64) *PoiSyncResponseDataTasksItem {
	s.TaskId = &v
	return s
}

type PoiSyncResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PoiSyncResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiSyncResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiSyncResponseExtra) SetLogid(v string) *PoiSyncResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiSyncResponseExtra) SetNow(v int64) *PoiSyncResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiSyncResponseExtra) SetSubDescription(v string) *PoiSyncResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiSyncResponseExtra) SetSubErrorCode(v int32) *PoiSyncResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiSyncResponseExtra) SetDescription(v string) *PoiSyncResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiSyncResponseExtra) SetErrorCode(v int32) *PoiSyncResponseExtra {
	s.ErrorCode = &v
	return s
}

type PoiTaskQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskIds     []*int64           `json:"task_ids,omitempty" xml:"task_ids,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PoiTaskQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryRequest) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryRequest) SetAccessToken(v string) *PoiTaskQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *PoiTaskQueryRequest) SetTaskIds(v []*int64) *PoiTaskQueryRequest {
	s.TaskIds = v
	return s
}

func (s *PoiTaskQueryRequest) SetHeader(v map[string]*string) *PoiTaskQueryRequest {
	s.Header = v
	return s
}

type PoiTaskQueryResponse struct {
	Extra *PoiTaskQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PoiTaskQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoiTaskQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponse) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponse) SetExtra(v *PoiTaskQueryResponseExtra) *PoiTaskQueryResponse {
	s.Extra = v
	return s
}

func (s *PoiTaskQueryResponse) SetData(v *PoiTaskQueryResponseData) *PoiTaskQueryResponse {
	s.Data = v
	return s
}

type PoiTaskQueryResponseData struct {
	TaskResults   []*PoiTaskQueryResponseDataTaskResultsItem `json:"task_results,omitempty" xml:"task_results,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiTaskQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseData) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseData) SetTaskResults(v []*PoiTaskQueryResponseDataTaskResultsItem) *PoiTaskQueryResponseData {
	s.TaskResults = v
	return s
}

func (s *PoiTaskQueryResponseData) SetGwErrorCode(v int32) *PoiTaskQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiTaskQueryResponseData) SetGwDescription(v string) *PoiTaskQueryResponseData {
	s.GwDescription = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItem struct {
	BusinessPartnershipResult *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult `json:"business_partnership_result,omitempty" xml:"business_partnership_result,omitempty"`
	ClaimResult               *PoiTaskQueryResponseDataTaskResultsItemClaimResult               `json:"claim_result,omitempty" xml:"claim_result,omitempty"`
	PoiDecorationResult       *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult       `json:"poi_decoration_result,omitempty" xml:"poi_decoration_result,omitempty"`
	PoiId                     *string                                                           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiResult                 *PoiTaskQueryResponseDataTaskResultsItemPoiResult                 `json:"poi_result,omitempty" xml:"poi_result,omitempty"`
	TaskId                    *int64                                                            `json:"task_id,omitempty" xml:"task_id,omitempty"`
	ThirdId                   *string                                                           `json:"third_id,omitempty" xml:"third_id,omitempty"`
	TrademarkResult           *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult           `json:"trademark_result,omitempty" xml:"trademark_result,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItem) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItem) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetBusinessPartnershipResult(v *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) *PoiTaskQueryResponseDataTaskResultsItem {
	s.BusinessPartnershipResult = v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetClaimResult(v *PoiTaskQueryResponseDataTaskResultsItemClaimResult) *PoiTaskQueryResponseDataTaskResultsItem {
	s.ClaimResult = v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetPoiDecorationResult(v *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) *PoiTaskQueryResponseDataTaskResultsItem {
	s.PoiDecorationResult = v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetPoiId(v string) *PoiTaskQueryResponseDataTaskResultsItem {
	s.PoiId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetPoiResult(v *PoiTaskQueryResponseDataTaskResultsItemPoiResult) *PoiTaskQueryResponseDataTaskResultsItem {
	s.PoiResult = v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetTaskId(v int64) *PoiTaskQueryResponseDataTaskResultsItem {
	s.TaskId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetThirdId(v string) *PoiTaskQueryResponseDataTaskResultsItem {
	s.ThirdId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItem) SetTrademarkResult(v *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) *PoiTaskQueryResponseDataTaskResultsItem {
	s.TrademarkResult = v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult struct {
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
	PartnershipId *int64  `json:"partnership_id,omitempty" xml:"partnership_id,omitempty"`
	RejectReason  *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) SetStatus(v int64) *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) SetPartnershipId(v int64) *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult {
	s.PartnershipId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemBusinessPartnershipResult {
	s.RejectReason = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemClaimResult struct {
	RejectReason       *string                                                                     `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status             *int64                                                                      `json:"status,omitempty" xml:"status,omitempty"`
	AccountId          *string                                                                     `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ExistLicenseId     *string                                                                     `json:"exist_license_id,omitempty" xml:"exist_license_id,omitempty"`
	KeyAuditResultList []*PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem `json:"key_audit_result_list,omitempty" xml:"key_audit_result_list,omitempty" type:"Repeated"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemClaimResult) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemClaimResult) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResult) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemClaimResult {
	s.RejectReason = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResult) SetStatus(v int64) *PoiTaskQueryResponseDataTaskResultsItemClaimResult {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResult) SetAccountId(v string) *PoiTaskQueryResponseDataTaskResultsItemClaimResult {
	s.AccountId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResult) SetExistLicenseId(v string) *PoiTaskQueryResponseDataTaskResultsItemClaimResult {
	s.ExistLicenseId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResult) SetKeyAuditResultList(v []*PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) *PoiTaskQueryResponseDataTaskResultsItemClaimResult {
	s.KeyAuditResultList = v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem struct {
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) SetStatus(v int) *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) SetKey(v string) *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem {
	s.Key = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemClaimResultKeyAuditResultListItem {
	s.RejectReason = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult struct {
	RejectReason       *string                                                                             `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status             *int                                                                                `json:"status,omitempty" xml:"status,omitempty"`
	KeyAuditResultList []*PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem `json:"key_audit_result_list,omitempty" xml:"key_audit_result_list,omitempty" type:"Repeated"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult {
	s.RejectReason = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) SetStatus(v int) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult) SetKeyAuditResultList(v []*PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResult {
	s.KeyAuditResultList = v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem struct {
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem {
	s.RejectReason = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) SetStatus(v int) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem) SetKey(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiDecorationResultKeyAuditResultListItem {
	s.Key = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemPoiResult struct {
	KeyAuditResultList []*PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem `json:"key_audit_result_list,omitempty" xml:"key_audit_result_list,omitempty" type:"Repeated"`
	RejectReason       *string                                                                   `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status             *int                                                                      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiResult) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiResult) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResult) SetKeyAuditResultList(v []*PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) *PoiTaskQueryResponseDataTaskResultsItemPoiResult {
	s.KeyAuditResultList = v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResult) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiResult {
	s.RejectReason = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResult) SetStatus(v int) *PoiTaskQueryResponseDataTaskResultsItemPoiResult {
	s.Status = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem struct {
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) SetKey(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem {
	s.Key = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem {
	s.RejectReason = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem) SetStatus(v int) *PoiTaskQueryResponseDataTaskResultsItemPoiResultKeyAuditResultListItem {
	s.Status = &v
	return s
}

type PoiTaskQueryResponseDataTaskResultsItemTrademarkResult struct {
	Status       *int64  `json:"status,omitempty" xml:"status,omitempty"`
	AccountId    *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) SetStatus(v int64) *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult {
	s.Status = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) SetAccountId(v string) *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult {
	s.AccountId = &v
	return s
}

func (s *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult) SetRejectReason(v string) *PoiTaskQueryResponseDataTaskResultsItemTrademarkResult {
	s.RejectReason = &v
	return s
}

type PoiTaskQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s PoiTaskQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiTaskQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiTaskQueryResponseExtra) SetSubDescription(v string) *PoiTaskQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiTaskQueryResponseExtra) SetSubErrorCode(v int32) *PoiTaskQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiTaskQueryResponseExtra) SetDescription(v string) *PoiTaskQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiTaskQueryResponseExtra) SetErrorCode(v int32) *PoiTaskQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiTaskQueryResponseExtra) SetLogid(v string) *PoiTaskQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiTaskQueryResponseExtra) SetNow(v int64) *PoiTaskQueryResponseExtra {
	s.Now = &v
	return s
}

type PoiUpdateRequest struct {
	Datas       []*PoiUpdateRequestDatasItem `json:"datas,omitempty" xml:"datas,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequest) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequest) SetDatas(v []*PoiUpdateRequestDatasItem) *PoiUpdateRequest {
	s.Datas = v
	return s
}

func (s *PoiUpdateRequest) SetHeader(v map[string]*string) *PoiUpdateRequest {
	s.Header = v
	return s
}

func (s *PoiUpdateRequest) SetAccessToken(v string) *PoiUpdateRequest {
	s.AccessToken = &v
	return s
}

type PoiUpdateRequestDatasItem struct {
	PoiId     *string                                  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiType   *int64                                   `json:"poi_type,omitempty" xml:"poi_type,omitempty"`
	Services  []*PoiUpdateRequestDatasItemServicesItem `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
	ThirdId   *string                                  `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId *string                                  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Delivery  *PoiUpdateRequestDatasItemDelivery       `json:"delivery,omitempty" xml:"delivery,omitempty"`
	Poi       *PoiUpdateRequestDatasItemPoi            `json:"poi,omitempty" xml:"poi,omitempty"`
}

func (s PoiUpdateRequestDatasItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItem) SetPoiId(v string) *PoiUpdateRequestDatasItem {
	s.PoiId = &v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetPoiType(v int64) *PoiUpdateRequestDatasItem {
	s.PoiType = &v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetServices(v []*PoiUpdateRequestDatasItemServicesItem) *PoiUpdateRequestDatasItem {
	s.Services = v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetThirdId(v string) *PoiUpdateRequestDatasItem {
	s.ThirdId = &v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetAccountId(v string) *PoiUpdateRequestDatasItem {
	s.AccountId = &v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetDelivery(v *PoiUpdateRequestDatasItemDelivery) *PoiUpdateRequestDatasItem {
	s.Delivery = v
	return s
}

func (s *PoiUpdateRequestDatasItem) SetPoi(v *PoiUpdateRequestDatasItemPoi) *PoiUpdateRequestDatasItem {
	s.Poi = v
	return s
}

type PoiUpdateRequestDatasItemDelivery struct {
	ServiceType      *int                                                   `json:"service_type,omitempty" xml:"service_type,omitempty"`
	DeliveryStrategy *string                                                `json:"delivery_strategy,omitempty" xml:"delivery_strategy,omitempty"`
	DeliveryStatus   *int64                                                 `json:"delivery_status,omitempty" xml:"delivery_status,omitempty"`
	BrandName        *string                                                `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	DeliveryType     *string                                                `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	AppId            *string                                                `json:"app_id,omitempty" xml:"app_id,omitempty"`
	OrderTimes       []*PoiUpdateRequestDatasItemDeliveryOrderTimesItem     `json:"order_times,omitempty" xml:"order_times,omitempty" type:"Repeated"`
	OrderStatus      *int                                                   `json:"order_status,omitempty" xml:"order_status,omitempty"`
	Subtitle         *string                                                `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	DeliveryScopes   []*PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem `json:"delivery_scopes,omitempty" xml:"delivery_scopes,omitempty" type:"Repeated"`
	DeliveryService  *string                                                `json:"delivery_service,omitempty" xml:"delivery_service,omitempty"`
	OpenAdvanceOrder *bool                                                  `json:"open_advance_order,omitempty" xml:"open_advance_order,omitempty"`
	BannerUrl        *string                                                `json:"banner_url,omitempty" xml:"banner_url,omitempty"`
}

func (s PoiUpdateRequestDatasItemDelivery) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemDelivery) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemDelivery) SetServiceType(v int) *PoiUpdateRequestDatasItemDelivery {
	s.ServiceType = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetDeliveryStrategy(v string) *PoiUpdateRequestDatasItemDelivery {
	s.DeliveryStrategy = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetDeliveryStatus(v int64) *PoiUpdateRequestDatasItemDelivery {
	s.DeliveryStatus = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetBrandName(v string) *PoiUpdateRequestDatasItemDelivery {
	s.BrandName = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetDeliveryType(v string) *PoiUpdateRequestDatasItemDelivery {
	s.DeliveryType = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetAppId(v string) *PoiUpdateRequestDatasItemDelivery {
	s.AppId = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetOrderTimes(v []*PoiUpdateRequestDatasItemDeliveryOrderTimesItem) *PoiUpdateRequestDatasItemDelivery {
	s.OrderTimes = v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetOrderStatus(v int) *PoiUpdateRequestDatasItemDelivery {
	s.OrderStatus = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetSubtitle(v string) *PoiUpdateRequestDatasItemDelivery {
	s.Subtitle = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetDeliveryScopes(v []*PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) *PoiUpdateRequestDatasItemDelivery {
	s.DeliveryScopes = v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetDeliveryService(v string) *PoiUpdateRequestDatasItemDelivery {
	s.DeliveryService = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetOpenAdvanceOrder(v bool) *PoiUpdateRequestDatasItemDelivery {
	s.OpenAdvanceOrder = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDelivery) SetBannerUrl(v string) *PoiUpdateRequestDatasItemDelivery {
	s.BannerUrl = &v
	return s
}

type PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem struct {
	MinPrice      *int64                                                           `json:"min_price,omitempty" xml:"min_price,omitempty"`
	Scopes        []*PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	DeliveryPrice *int64                                                           `json:"delivery_price,omitempty" xml:"delivery_price,omitempty"`
}

func (s PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) SetMinPrice(v int64) *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem {
	s.MinPrice = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) SetScopes(v []*PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem) *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem {
	s.Scopes = v
	return s
}

func (s *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem) SetDeliveryPrice(v int64) *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItem {
	s.DeliveryPrice = &v
	return s
}

type PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem struct {
	Latitude  *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem) SetLatitude(v string) *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem {
	s.Latitude = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem) SetLongitude(v string) *PoiUpdateRequestDatasItemDeliveryDeliveryScopesItemScopesItem {
	s.Longitude = &v
	return s
}

type PoiUpdateRequestDatasItemDeliveryOrderTimesItem struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Weeks     []*int  `json:"weeks,omitempty" xml:"weeks,omitempty" type:"Repeated"`
}

func (s PoiUpdateRequestDatasItemDeliveryOrderTimesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemDeliveryOrderTimesItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemDeliveryOrderTimesItem) SetEndTime(v string) *PoiUpdateRequestDatasItemDeliveryOrderTimesItem {
	s.EndTime = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDeliveryOrderTimesItem) SetStartTime(v string) *PoiUpdateRequestDatasItemDeliveryOrderTimesItem {
	s.StartTime = &v
	return s
}

func (s *PoiUpdateRequestDatasItemDeliveryOrderTimesItem) SetWeeks(v []*int) *PoiUpdateRequestDatasItemDeliveryOrderTimesItem {
	s.Weeks = v
	return s
}

type PoiUpdateRequestDatasItemPoi struct {
	Address       *string                                        `json:"address,omitempty" xml:"address,omitempty"`
	Longitude     *string                                        `json:"longitude,omitempty" xml:"longitude,omitempty"`
	OpenTimes     []*string                                      `json:"open_times,omitempty" xml:"open_times,omitempty" type:"Repeated"`
	Latitude      *string                                        `json:"latitude,omitempty" xml:"latitude,omitempty"`
	IndustryCode  *string                                        `json:"industry_code,omitempty" xml:"industry_code,omitempty"`
	ContactTel    *string                                        `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
	PoiName       *string                                        `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	HeadImageUrls []*string                                      `json:"head_image_urls,omitempty" xml:"head_image_urls,omitempty" type:"Repeated"`
	ContactPhone  *string                                        `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	OpenTimesV2   []*PoiUpdateRequestDatasItemPoiOpenTimesV2Item `json:"open_times_v2,omitempty" xml:"open_times_v2,omitempty" type:"Repeated"`
	OpenStatus    *int32                                         `json:"open_status,omitempty" xml:"open_status,omitempty"`
}

func (s PoiUpdateRequestDatasItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemPoi) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemPoi) SetAddress(v string) *PoiUpdateRequestDatasItemPoi {
	s.Address = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetLongitude(v string) *PoiUpdateRequestDatasItemPoi {
	s.Longitude = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetOpenTimes(v []*string) *PoiUpdateRequestDatasItemPoi {
	s.OpenTimes = v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetLatitude(v string) *PoiUpdateRequestDatasItemPoi {
	s.Latitude = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetIndustryCode(v string) *PoiUpdateRequestDatasItemPoi {
	s.IndustryCode = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetContactTel(v string) *PoiUpdateRequestDatasItemPoi {
	s.ContactTel = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetPoiName(v string) *PoiUpdateRequestDatasItemPoi {
	s.PoiName = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetHeadImageUrls(v []*string) *PoiUpdateRequestDatasItemPoi {
	s.HeadImageUrls = v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetContactPhone(v string) *PoiUpdateRequestDatasItemPoi {
	s.ContactPhone = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetOpenTimesV2(v []*PoiUpdateRequestDatasItemPoiOpenTimesV2Item) *PoiUpdateRequestDatasItemPoi {
	s.OpenTimesV2 = v
	return s
}

func (s *PoiUpdateRequestDatasItemPoi) SetOpenStatus(v int32) *PoiUpdateRequestDatasItemPoi {
	s.OpenStatus = &v
	return s
}

type PoiUpdateRequestDatasItemPoiOpenTimesV2Item struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Weeks     []*int  `json:"weeks,omitempty" xml:"weeks,omitempty" type:"Repeated"`
}

func (s PoiUpdateRequestDatasItemPoiOpenTimesV2Item) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemPoiOpenTimesV2Item) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemPoiOpenTimesV2Item) SetEndTime(v string) *PoiUpdateRequestDatasItemPoiOpenTimesV2Item {
	s.EndTime = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoiOpenTimesV2Item) SetStartTime(v string) *PoiUpdateRequestDatasItemPoiOpenTimesV2Item {
	s.StartTime = &v
	return s
}

func (s *PoiUpdateRequestDatasItemPoiOpenTimesV2Item) SetWeeks(v []*int) *PoiUpdateRequestDatasItemPoiOpenTimesV2Item {
	s.Weeks = v
	return s
}

type PoiUpdateRequestDatasItemServicesItem struct {
	Entry       *PoiUpdateRequestDatasItemServicesItemEntry `json:"entry,omitempty" xml:"entry,omitempty"`
	ServiceType *int                                        `json:"service_type,omitempty" xml:"service_type,omitempty"`
	Enable      *int                                        `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s PoiUpdateRequestDatasItemServicesItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemServicesItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemServicesItem) SetEntry(v *PoiUpdateRequestDatasItemServicesItemEntry) *PoiUpdateRequestDatasItemServicesItem {
	s.Entry = v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItem) SetServiceType(v int) *PoiUpdateRequestDatasItemServicesItem {
	s.ServiceType = &v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItem) SetEnable(v int) *PoiUpdateRequestDatasItemServicesItem {
	s.Enable = &v
	return s
}

type PoiUpdateRequestDatasItemServicesItemEntry struct {
	EntryMiniApp *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp `json:"entry_mini_app,omitempty" xml:"entry_mini_app,omitempty"`
	EntryType    *int                                                    `json:"entry_type,omitempty" xml:"entry_type,omitempty"`
}

func (s PoiUpdateRequestDatasItemServicesItemEntry) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemServicesItemEntry) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemServicesItemEntry) SetEntryMiniApp(v *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) *PoiUpdateRequestDatasItemServicesItemEntry {
	s.EntryMiniApp = v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItemEntry) SetEntryType(v int) *PoiUpdateRequestDatasItemServicesItemEntry {
	s.EntryType = &v
	return s
}

type PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp struct {
	IsTest *int64  `json:"is_test,omitempty" xml:"is_test,omitempty"`
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
	AppId  *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

func (s PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) GoString() string {
	return s.String()
}

func (s *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) SetIsTest(v int64) *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp {
	s.IsTest = &v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) SetParams(v string) *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp {
	s.Params = &v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) SetPath(v string) *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp {
	s.Path = &v
	return s
}

func (s *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp) SetAppId(v string) *PoiUpdateRequestDatasItemServicesItemEntryEntryMiniApp {
	s.AppId = &v
	return s
}

type PoiUpdateResponse struct {
	Extra *PoiUpdateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *PoiUpdateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PoiUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateResponse) GoString() string {
	return s.String()
}

func (s *PoiUpdateResponse) SetExtra(v *PoiUpdateResponseExtra) *PoiUpdateResponse {
	s.Extra = v
	return s
}

func (s *PoiUpdateResponse) SetData(v *PoiUpdateResponseData) *PoiUpdateResponse {
	s.Data = v
	return s
}

type PoiUpdateResponseData struct {
	Tasks         []*PoiUpdateResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	GwErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiUpdateResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateResponseData) GoString() string {
	return s.String()
}

func (s *PoiUpdateResponseData) SetTasks(v []*PoiUpdateResponseDataTasksItem) *PoiUpdateResponseData {
	s.Tasks = v
	return s
}

func (s *PoiUpdateResponseData) SetGwErrorCode(v int32) *PoiUpdateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PoiUpdateResponseData) SetGwDescription(v string) *PoiUpdateResponseData {
	s.GwDescription = &v
	return s
}

type PoiUpdateResponseDataTasksItem struct {
	TaskId    *int64  `json:"task_id,omitempty" xml:"task_id,omitempty"`
	ThirdId   *string `json:"third_id,omitempty" xml:"third_id,omitempty"`
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId     *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoiUpdateResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *PoiUpdateResponseDataTasksItem) SetTaskId(v int64) *PoiUpdateResponseDataTasksItem {
	s.TaskId = &v
	return s
}

func (s *PoiUpdateResponseDataTasksItem) SetThirdId(v string) *PoiUpdateResponseDataTasksItem {
	s.ThirdId = &v
	return s
}

func (s *PoiUpdateResponseDataTasksItem) SetAccountId(v string) *PoiUpdateResponseDataTasksItem {
	s.AccountId = &v
	return s
}

func (s *PoiUpdateResponseDataTasksItem) SetPoiId(v string) *PoiUpdateResponseDataTasksItem {
	s.PoiId = &v
	return s
}

type PoiUpdateResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s PoiUpdateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiUpdateResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiUpdateResponseExtra) SetSubDescription(v string) *PoiUpdateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiUpdateResponseExtra) SetSubErrorCode(v int32) *PoiUpdateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiUpdateResponseExtra) SetDescription(v string) *PoiUpdateResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiUpdateResponseExtra) SetErrorCode(v int32) *PoiUpdateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiUpdateResponseExtra) SetLogid(v string) *PoiUpdateResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiUpdateResponseExtra) SetNow(v int64) *PoiUpdateResponseExtra {
	s.Now = &v
	return s
}

type PoipriceSaveRequest struct {
	PoiSkuPriceList []*PoipriceSaveRequestPoiSkuPriceListItem `json:"poi_sku_price_list,omitempty" xml:"poi_sku_price_list,omitempty" require:"true" type:"Repeated"`
	Base            *PoipriceSaveRequestBase                  `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *int64                                    `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header          map[string]*string                        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	IgnoreFailPoi   *bool                                     `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
	OutProductId    *string                                   `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
}

func (s PoipriceSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequest) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequest) SetPoiSkuPriceList(v []*PoipriceSaveRequestPoiSkuPriceListItem) *PoipriceSaveRequest {
	s.PoiSkuPriceList = v
	return s
}

func (s *PoipriceSaveRequest) SetBase(v *PoipriceSaveRequestBase) *PoipriceSaveRequest {
	s.Base = v
	return s
}

func (s *PoipriceSaveRequest) SetAccountId(v int64) *PoipriceSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PoipriceSaveRequest) SetHeader(v map[string]*string) *PoipriceSaveRequest {
	s.Header = v
	return s
}

func (s *PoipriceSaveRequest) SetAccessToken(v string) *PoipriceSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *PoipriceSaveRequest) SetIgnoreFailPoi(v bool) *PoipriceSaveRequest {
	s.IgnoreFailPoi = &v
	return s
}

func (s *PoipriceSaveRequest) SetOutProductId(v string) *PoipriceSaveRequest {
	s.OutProductId = &v
	return s
}

type PoipriceSaveRequestBase struct {
	Client     *string                            `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                 `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                            `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *PoipriceSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                            `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                            `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s PoipriceSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestBase) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestBase) SetClient(v string) *PoipriceSaveRequestBase {
	s.Client = &v
	return s
}

func (s *PoipriceSaveRequestBase) SetExtra(v map[string]*string) *PoipriceSaveRequestBase {
	s.Extra = v
	return s
}

func (s *PoipriceSaveRequestBase) SetLogID(v string) *PoipriceSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *PoipriceSaveRequestBase) SetTrafficEnv(v *PoipriceSaveRequestBaseTrafficEnv) *PoipriceSaveRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *PoipriceSaveRequestBase) SetAddr(v string) *PoipriceSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *PoipriceSaveRequestBase) SetCaller(v string) *PoipriceSaveRequestBase {
	s.Caller = &v
	return s
}

type PoipriceSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s PoipriceSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestBaseTrafficEnv) SetEnv(v string) *PoipriceSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *PoipriceSaveRequestBaseTrafficEnv) SetOpen(v bool) *PoipriceSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type PoipriceSaveRequestPoiSkuPriceListItem struct {
	Price *PoipriceSaveRequestPoiSkuPriceListItemPrice `json:"price,omitempty" xml:"price,omitempty"`
	Poi   *PoipriceSaveRequestPoiSkuPriceListItemPoi   `json:"poi,omitempty" xml:"poi,omitempty"`
}

func (s PoipriceSaveRequestPoiSkuPriceListItem) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestPoiSkuPriceListItem) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestPoiSkuPriceListItem) SetPrice(v *PoipriceSaveRequestPoiSkuPriceListItemPrice) *PoipriceSaveRequestPoiSkuPriceListItem {
	s.Price = v
	return s
}

func (s *PoipriceSaveRequestPoiSkuPriceListItem) SetPoi(v *PoipriceSaveRequestPoiSkuPriceListItemPoi) *PoipriceSaveRequestPoiSkuPriceListItem {
	s.Poi = v
	return s
}

type PoipriceSaveRequestPoiSkuPriceListItemPoi struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPoi) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestPoiSkuPriceListItemPoi) SetExtId(v string) *PoipriceSaveRequestPoiSkuPriceListItemPoi {
	s.ExtId = &v
	return s
}

func (s *PoipriceSaveRequestPoiSkuPriceListItemPoi) SetPoiId(v int64) *PoipriceSaveRequestPoiSkuPriceListItemPoi {
	s.PoiId = &v
	return s
}

type PoipriceSaveRequestPoiSkuPriceListItemPrice struct {
	OutSkuId *string                                           `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price    *PoipriceSaveRequestPoiSkuPriceListItemPricePrice `json:"price,omitempty" xml:"price,omitempty"`
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPrice) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestPoiSkuPriceListItemPrice) SetOutSkuId(v string) *PoipriceSaveRequestPoiSkuPriceListItemPrice {
	s.OutSkuId = &v
	return s
}

func (s *PoipriceSaveRequestPoiSkuPriceListItemPrice) SetPrice(v *PoipriceSaveRequestPoiSkuPriceListItemPricePrice) *PoipriceSaveRequestPoiSkuPriceListItemPrice {
	s.Price = v
	return s
}

type PoipriceSaveRequestPoiSkuPriceListItemPricePrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPricePrice) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveRequestPoiSkuPriceListItemPricePrice) GoString() string {
	return s.String()
}

func (s *PoipriceSaveRequestPoiSkuPriceListItemPricePrice) SetActualAmount(v int64) *PoipriceSaveRequestPoiSkuPriceListItemPricePrice {
	s.ActualAmount = &v
	return s
}

type PoipriceSaveResponse struct {
	Data     *PoipriceSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoipriceSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *PoipriceSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s PoipriceSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveResponse) GoString() string {
	return s.String()
}

func (s *PoipriceSaveResponse) SetData(v *PoipriceSaveResponseData) *PoipriceSaveResponse {
	s.Data = v
	return s
}

func (s *PoipriceSaveResponse) SetExtra(v *PoipriceSaveResponseExtra) *PoipriceSaveResponse {
	s.Extra = v
	return s
}

func (s *PoipriceSaveResponse) SetBaseResp(v *PoipriceSaveResponseBaseResp) *PoipriceSaveResponse {
	s.BaseResp = v
	return s
}

type PoipriceSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s PoipriceSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoipriceSaveResponseBaseResp) SetStatusCode(v int32) *PoipriceSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoipriceSaveResponseBaseResp) SetStatusMessage(v string) *PoipriceSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoipriceSaveResponseBaseResp) SetExtra(v map[string]*string) *PoipriceSaveResponseBaseResp {
	s.Extra = v
	return s
}

type PoipriceSaveResponseData struct {
	ErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	FailPoiList []*PoipriceSaveResponseDataFailPoiListItem `json:"fail_poi_list,omitempty" xml:"fail_poi_list,omitempty" type:"Repeated"`
	Description *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoipriceSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveResponseData) GoString() string {
	return s.String()
}

func (s *PoipriceSaveResponseData) SetErrorCode(v int32) *PoipriceSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PoipriceSaveResponseData) SetFailPoiList(v []*PoipriceSaveResponseDataFailPoiListItem) *PoipriceSaveResponseData {
	s.FailPoiList = v
	return s
}

func (s *PoipriceSaveResponseData) SetDescription(v string) *PoipriceSaveResponseData {
	s.Description = &v
	return s
}

type PoipriceSaveResponseDataFailPoiListItem struct {
	ExtId  *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId  *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s PoipriceSaveResponseDataFailPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveResponseDataFailPoiListItem) GoString() string {
	return s.String()
}

func (s *PoipriceSaveResponseDataFailPoiListItem) SetExtId(v string) *PoipriceSaveResponseDataFailPoiListItem {
	s.ExtId = &v
	return s
}

func (s *PoipriceSaveResponseDataFailPoiListItem) SetPoiId(v int64) *PoipriceSaveResponseDataFailPoiListItem {
	s.PoiId = &v
	return s
}

func (s *PoipriceSaveResponseDataFailPoiListItem) SetReason(v string) *PoipriceSaveResponseDataFailPoiListItem {
	s.Reason = &v
	return s
}

func (s *PoipriceSaveResponseDataFailPoiListItem) SetCode(v string) *PoipriceSaveResponseDataFailPoiListItem {
	s.Code = &v
	return s
}

type PoipriceSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoipriceSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoipriceSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PoipriceSaveResponseExtra) SetSubErrorCode(v int32) *PoipriceSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoipriceSaveResponseExtra) SetDescription(v string) *PoipriceSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *PoipriceSaveResponseExtra) SetErrorCode(v int32) *PoipriceSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoipriceSaveResponseExtra) SetLogid(v string) *PoipriceSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoipriceSaveResponseExtra) SetNow(v int64) *PoipriceSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PoipriceSaveResponseExtra) SetSubDescription(v string) *PoipriceSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type PoiproductDetailRequest struct {
	PoiExtIds    []*string          `json:"poi_ext_ids,omitempty" xml:"poi_ext_ids,omitempty" type:"Repeated"`
	PoiIds       []*int64           `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
	AccountId    *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OutProductId *string            `json:"out_product_id,omitempty" xml:"out_product_id,omitempty" require:"true"`
	OutSkuIdList []*string          `json:"out_sku_id_list,omitempty" xml:"out_sku_id_list,omitempty" type:"Repeated"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiproductDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailRequest) GoString() string {
	return s.String()
}

func (s *PoiproductDetailRequest) SetPoiExtIds(v []*string) *PoiproductDetailRequest {
	s.PoiExtIds = v
	return s
}

func (s *PoiproductDetailRequest) SetPoiIds(v []*int64) *PoiproductDetailRequest {
	s.PoiIds = v
	return s
}

func (s *PoiproductDetailRequest) SetAccountId(v int64) *PoiproductDetailRequest {
	s.AccountId = &v
	return s
}

func (s *PoiproductDetailRequest) SetOutProductId(v string) *PoiproductDetailRequest {
	s.OutProductId = &v
	return s
}

func (s *PoiproductDetailRequest) SetOutSkuIdList(v []*string) *PoiproductDetailRequest {
	s.OutSkuIdList = v
	return s
}

func (s *PoiproductDetailRequest) SetHeader(v map[string]*string) *PoiproductDetailRequest {
	s.Header = v
	return s
}

func (s *PoiproductDetailRequest) SetAccessToken(v string) *PoiproductDetailRequest {
	s.AccessToken = &v
	return s
}

type PoiproductDetailResponse struct {
	Data     *PoiproductDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoiproductDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *PoiproductDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s PoiproductDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponse) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponse) SetData(v *PoiproductDetailResponseData) *PoiproductDetailResponse {
	s.Data = v
	return s
}

func (s *PoiproductDetailResponse) SetExtra(v *PoiproductDetailResponseExtra) *PoiproductDetailResponse {
	s.Extra = v
	return s
}

func (s *PoiproductDetailResponse) SetBaseResp(v *PoiproductDetailResponseBaseResp) *PoiproductDetailResponse {
	s.BaseResp = v
	return s
}

type PoiproductDetailResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s PoiproductDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseBaseResp) SetStatusCode(v int32) *PoiproductDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoiproductDetailResponseBaseResp) SetStatusMessage(v string) *PoiproductDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoiproductDetailResponseBaseResp) SetExtra(v map[string]*string) *PoiproductDetailResponseBaseResp {
	s.Extra = v
	return s
}

type PoiproductDetailResponseData struct {
	DetailList            []*PoiproductDetailResponseDataDetailListItem            `json:"detail_list,omitempty" xml:"detail_list,omitempty" type:"Repeated"`
	ErrorCode             *int32                                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	FailPoiList           []*PoiproductDetailResponseDataFailPoiListItem           `json:"fail_poi_list,omitempty" xml:"fail_poi_list,omitempty" type:"Repeated"`
	QueryPriceFailSkuList []*PoiproductDetailResponseDataQueryPriceFailSkuListItem `json:"query_price_fail_sku_list,omitempty" xml:"query_price_fail_sku_list,omitempty" type:"Repeated"`
	QueryStockFailSkuList []*PoiproductDetailResponseDataQueryStockFailSkuListItem `json:"query_stock_fail_sku_list,omitempty" xml:"query_stock_fail_sku_list,omitempty" type:"Repeated"`
	Description           *string                                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoiproductDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseData) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseData) SetDetailList(v []*PoiproductDetailResponseDataDetailListItem) *PoiproductDetailResponseData {
	s.DetailList = v
	return s
}

func (s *PoiproductDetailResponseData) SetErrorCode(v int32) *PoiproductDetailResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PoiproductDetailResponseData) SetFailPoiList(v []*PoiproductDetailResponseDataFailPoiListItem) *PoiproductDetailResponseData {
	s.FailPoiList = v
	return s
}

func (s *PoiproductDetailResponseData) SetQueryPriceFailSkuList(v []*PoiproductDetailResponseDataQueryPriceFailSkuListItem) *PoiproductDetailResponseData {
	s.QueryPriceFailSkuList = v
	return s
}

func (s *PoiproductDetailResponseData) SetQueryStockFailSkuList(v []*PoiproductDetailResponseDataQueryStockFailSkuListItem) *PoiproductDetailResponseData {
	s.QueryStockFailSkuList = v
	return s
}

func (s *PoiproductDetailResponseData) SetDescription(v string) *PoiproductDetailResponseData {
	s.Description = &v
	return s
}

type PoiproductDetailResponseDataDetailListItem struct {
	Status             *int                                                                `json:"status,omitempty" xml:"status,omitempty"`
	OutId              *string                                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Poi                *PoiproductDetailResponseDataDetailListItemPoi                      `json:"poi,omitempty" xml:"poi,omitempty"`
	ProductId          *int64                                                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SaleAttrStatusList []*PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem `json:"sale_attr_status_list,omitempty" xml:"sale_attr_status_list,omitempty" type:"Repeated"`
	SkuPriceList       []*PoiproductDetailResponseDataDetailListItemSkuPriceListItem       `json:"sku_price_list,omitempty" xml:"sku_price_list,omitempty" type:"Repeated"`
	SkuStatusList      []*PoiproductDetailResponseDataDetailListItemSkuStatusListItem      `json:"sku_status_list,omitempty" xml:"sku_status_list,omitempty" type:"Repeated"`
	SkuStockList       []*PoiproductDetailResponseDataDetailListItemSkuStockListItem       `json:"sku_stock_list,omitempty" xml:"sku_stock_list,omitempty" type:"Repeated"`
}

func (s PoiproductDetailResponseDataDetailListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItem) SetStatus(v int) *PoiproductDetailResponseDataDetailListItem {
	s.Status = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetOutId(v string) *PoiproductDetailResponseDataDetailListItem {
	s.OutId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetPoi(v *PoiproductDetailResponseDataDetailListItemPoi) *PoiproductDetailResponseDataDetailListItem {
	s.Poi = v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetProductId(v int64) *PoiproductDetailResponseDataDetailListItem {
	s.ProductId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetSaleAttrStatusList(v []*PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) *PoiproductDetailResponseDataDetailListItem {
	s.SaleAttrStatusList = v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetSkuPriceList(v []*PoiproductDetailResponseDataDetailListItemSkuPriceListItem) *PoiproductDetailResponseDataDetailListItem {
	s.SkuPriceList = v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetSkuStatusList(v []*PoiproductDetailResponseDataDetailListItemSkuStatusListItem) *PoiproductDetailResponseDataDetailListItem {
	s.SkuStatusList = v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItem) SetSkuStockList(v []*PoiproductDetailResponseDataDetailListItemSkuStockListItem) *PoiproductDetailResponseDataDetailListItem {
	s.SkuStockList = v
	return s
}

type PoiproductDetailResponseDataDetailListItemPoi struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoiproductDetailResponseDataDetailListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemPoi) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemPoi) SetExtId(v string) *PoiproductDetailResponseDataDetailListItemPoi {
	s.ExtId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemPoi) SetPoiId(v int64) *PoiproductDetailResponseDataDetailListItemPoi {
	s.PoiId = &v
	return s
}

type PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem struct {
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty" require:"true"`
	Status    *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) SetGroupCode(v string) *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem {
	s.GroupCode = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) SetItemKey(v string) *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem {
	s.ItemKey = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem) SetStatus(v int32) *PoiproductDetailResponseDataDetailListItemSaleAttrStatusListItem {
	s.Status = &v
	return s
}

type PoiproductDetailResponseDataDetailListItemSkuPriceListItem struct {
	OutSkuId *string                                                          `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price    *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice `json:"price,omitempty" xml:"price,omitempty"`
	SkuId    *int64                                                           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s PoiproductDetailResponseDataDetailListItemSkuPriceListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemSkuPriceListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemSkuPriceListItem) SetOutSkuId(v string) *PoiproductDetailResponseDataDetailListItemSkuPriceListItem {
	s.OutSkuId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuPriceListItem) SetPrice(v *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice) *PoiproductDetailResponseDataDetailListItemSkuPriceListItem {
	s.Price = v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuPriceListItem) SetSkuId(v int64) *PoiproductDetailResponseDataDetailListItemSkuPriceListItem {
	s.SkuId = &v
	return s
}

type PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice) SetActualAmount(v int64) *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice {
	s.ActualAmount = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice) SetOriginAmount(v int64) *PoiproductDetailResponseDataDetailListItemSkuPriceListItemPrice {
	s.OriginAmount = &v
	return s
}

type PoiproductDetailResponseDataDetailListItemSkuStatusListItem struct {
	SkuId    *int64  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Status   *int    `json:"status,omitempty" xml:"status,omitempty"`
	OutSkuId *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s PoiproductDetailResponseDataDetailListItemSkuStatusListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemSkuStatusListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStatusListItem) SetSkuId(v int64) *PoiproductDetailResponseDataDetailListItemSkuStatusListItem {
	s.SkuId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStatusListItem) SetStatus(v int) *PoiproductDetailResponseDataDetailListItemSkuStatusListItem {
	s.Status = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStatusListItem) SetOutSkuId(v string) *PoiproductDetailResponseDataDetailListItemSkuStatusListItem {
	s.OutSkuId = &v
	return s
}

type PoiproductDetailResponseDataDetailListItemSkuStockListItem struct {
	SkuId    *int64  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Stock    *int64  `json:"stock,omitempty" xml:"stock,omitempty"`
	OutSkuId *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s PoiproductDetailResponseDataDetailListItemSkuStockListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataDetailListItemSkuStockListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStockListItem) SetSkuId(v int64) *PoiproductDetailResponseDataDetailListItemSkuStockListItem {
	s.SkuId = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStockListItem) SetStock(v int64) *PoiproductDetailResponseDataDetailListItemSkuStockListItem {
	s.Stock = &v
	return s
}

func (s *PoiproductDetailResponseDataDetailListItemSkuStockListItem) SetOutSkuId(v string) *PoiproductDetailResponseDataDetailListItemSkuStockListItem {
	s.OutSkuId = &v
	return s
}

type PoiproductDetailResponseDataFailPoiListItem struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	ExtId  *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId  *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoiproductDetailResponseDataFailPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataFailPoiListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataFailPoiListItem) SetReason(v string) *PoiproductDetailResponseDataFailPoiListItem {
	s.Reason = &v
	return s
}

func (s *PoiproductDetailResponseDataFailPoiListItem) SetCode(v string) *PoiproductDetailResponseDataFailPoiListItem {
	s.Code = &v
	return s
}

func (s *PoiproductDetailResponseDataFailPoiListItem) SetExtId(v string) *PoiproductDetailResponseDataFailPoiListItem {
	s.ExtId = &v
	return s
}

func (s *PoiproductDetailResponseDataFailPoiListItem) SetPoiId(v int64) *PoiproductDetailResponseDataFailPoiListItem {
	s.PoiId = &v
	return s
}

type PoiproductDetailResponseDataQueryPriceFailSkuListItem struct {
	SkuId  *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s PoiproductDetailResponseDataQueryPriceFailSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataQueryPriceFailSkuListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataQueryPriceFailSkuListItem) SetSkuId(v int64) *PoiproductDetailResponseDataQueryPriceFailSkuListItem {
	s.SkuId = &v
	return s
}

func (s *PoiproductDetailResponseDataQueryPriceFailSkuListItem) SetCode(v string) *PoiproductDetailResponseDataQueryPriceFailSkuListItem {
	s.Code = &v
	return s
}

func (s *PoiproductDetailResponseDataQueryPriceFailSkuListItem) SetReason(v string) *PoiproductDetailResponseDataQueryPriceFailSkuListItem {
	s.Reason = &v
	return s
}

type PoiproductDetailResponseDataQueryStockFailSkuListItem struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	SkuId  *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s PoiproductDetailResponseDataQueryStockFailSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseDataQueryStockFailSkuListItem) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseDataQueryStockFailSkuListItem) SetCode(v string) *PoiproductDetailResponseDataQueryStockFailSkuListItem {
	s.Code = &v
	return s
}

func (s *PoiproductDetailResponseDataQueryStockFailSkuListItem) SetReason(v string) *PoiproductDetailResponseDataQueryStockFailSkuListItem {
	s.Reason = &v
	return s
}

func (s *PoiproductDetailResponseDataQueryStockFailSkuListItem) SetSkuId(v int64) *PoiproductDetailResponseDataQueryStockFailSkuListItem {
	s.SkuId = &v
	return s
}

type PoiproductDetailResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s PoiproductDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiproductDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiproductDetailResponseExtra) SetNow(v int64) *PoiproductDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiproductDetailResponseExtra) SetSubDescription(v string) *PoiproductDetailResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PoiproductDetailResponseExtra) SetSubErrorCode(v int32) *PoiproductDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiproductDetailResponseExtra) SetDescription(v string) *PoiproductDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiproductDetailResponseExtra) SetErrorCode(v int32) *PoiproductDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiproductDetailResponseExtra) SetLogid(v string) *PoiproductDetailResponseExtra {
	s.Logid = &v
	return s
}

type PoiproductOperateRequest struct {
	IgnoreFailPoi   *bool                                          `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
	OperateInfoList []*PoiproductOperateRequestOperateInfoListItem `json:"operate_info_list,omitempty" xml:"operate_info_list,omitempty" require:"true" type:"Repeated"`
	OutProductId    *string                                        `json:"out_product_id,omitempty" xml:"out_product_id,omitempty" require:"true"`
	AccountId       *int64                                         `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header          map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoiproductOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateRequest) GoString() string {
	return s.String()
}

func (s *PoiproductOperateRequest) SetIgnoreFailPoi(v bool) *PoiproductOperateRequest {
	s.IgnoreFailPoi = &v
	return s
}

func (s *PoiproductOperateRequest) SetOperateInfoList(v []*PoiproductOperateRequestOperateInfoListItem) *PoiproductOperateRequest {
	s.OperateInfoList = v
	return s
}

func (s *PoiproductOperateRequest) SetOutProductId(v string) *PoiproductOperateRequest {
	s.OutProductId = &v
	return s
}

func (s *PoiproductOperateRequest) SetAccountId(v int64) *PoiproductOperateRequest {
	s.AccountId = &v
	return s
}

func (s *PoiproductOperateRequest) SetHeader(v map[string]*string) *PoiproductOperateRequest {
	s.Header = v
	return s
}

func (s *PoiproductOperateRequest) SetAccessToken(v string) *PoiproductOperateRequest {
	s.AccessToken = &v
	return s
}

type PoiproductOperateRequestOperateInfoListItem struct {
	SaleAttrStatusList []*PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem `json:"sale_attr_status_list,omitempty" xml:"sale_attr_status_list,omitempty" type:"Repeated"`
	SkuStatusList      []*PoiproductOperateRequestOperateInfoListItemSkuStatusListItem      `json:"sku_status_list,omitempty" xml:"sku_status_list,omitempty" type:"Repeated"`
	Poi                *PoiproductOperateRequestOperateInfoListItemPoi                      `json:"poi,omitempty" xml:"poi,omitempty" require:"true"`
	ProductStatus      *int                                                                 `json:"product_status,omitempty" xml:"product_status,omitempty"`
}

func (s PoiproductOperateRequestOperateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateRequestOperateInfoListItem) GoString() string {
	return s.String()
}

func (s *PoiproductOperateRequestOperateInfoListItem) SetSaleAttrStatusList(v []*PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) *PoiproductOperateRequestOperateInfoListItem {
	s.SaleAttrStatusList = v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItem) SetSkuStatusList(v []*PoiproductOperateRequestOperateInfoListItemSkuStatusListItem) *PoiproductOperateRequestOperateInfoListItem {
	s.SkuStatusList = v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItem) SetPoi(v *PoiproductOperateRequestOperateInfoListItemPoi) *PoiproductOperateRequestOperateInfoListItem {
	s.Poi = v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItem) SetProductStatus(v int) *PoiproductOperateRequestOperateInfoListItem {
	s.ProductStatus = &v
	return s
}

type PoiproductOperateRequestOperateInfoListItemPoi struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoiproductOperateRequestOperateInfoListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateRequestOperateInfoListItemPoi) GoString() string {
	return s.String()
}

func (s *PoiproductOperateRequestOperateInfoListItemPoi) SetExtId(v string) *PoiproductOperateRequestOperateInfoListItemPoi {
	s.ExtId = &v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItemPoi) SetPoiId(v int64) *PoiproductOperateRequestOperateInfoListItemPoi {
	s.PoiId = &v
	return s
}

type PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem struct {
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty" require:"true"`
	Status    *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) GoString() string {
	return s.String()
}

func (s *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) SetItemKey(v string) *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem {
	s.ItemKey = &v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) SetStatus(v int32) *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem {
	s.Status = &v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem) SetGroupCode(v string) *PoiproductOperateRequestOperateInfoListItemSaleAttrStatusListItem {
	s.GroupCode = &v
	return s
}

type PoiproductOperateRequestOperateInfoListItemSkuStatusListItem struct {
	OutSkuId *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Status   *int    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PoiproductOperateRequestOperateInfoListItemSkuStatusListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateRequestOperateInfoListItemSkuStatusListItem) GoString() string {
	return s.String()
}

func (s *PoiproductOperateRequestOperateInfoListItemSkuStatusListItem) SetOutSkuId(v string) *PoiproductOperateRequestOperateInfoListItemSkuStatusListItem {
	s.OutSkuId = &v
	return s
}

func (s *PoiproductOperateRequestOperateInfoListItemSkuStatusListItem) SetStatus(v int) *PoiproductOperateRequestOperateInfoListItemSkuStatusListItem {
	s.Status = &v
	return s
}

type PoiproductOperateResponse struct {
	Data     *PoiproductOperateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoiproductOperateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *PoiproductOperateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s PoiproductOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateResponse) GoString() string {
	return s.String()
}

func (s *PoiproductOperateResponse) SetData(v *PoiproductOperateResponseData) *PoiproductOperateResponse {
	s.Data = v
	return s
}

func (s *PoiproductOperateResponse) SetExtra(v *PoiproductOperateResponseExtra) *PoiproductOperateResponse {
	s.Extra = v
	return s
}

func (s *PoiproductOperateResponse) SetBaseResp(v *PoiproductOperateResponseBaseResp) *PoiproductOperateResponse {
	s.BaseResp = v
	return s
}

type PoiproductOperateResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PoiproductOperateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoiproductOperateResponseBaseResp) SetStatusMessage(v string) *PoiproductOperateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoiproductOperateResponseBaseResp) SetExtra(v map[string]*string) *PoiproductOperateResponseBaseResp {
	s.Extra = v
	return s
}

func (s *PoiproductOperateResponseBaseResp) SetStatusCode(v int32) *PoiproductOperateResponseBaseResp {
	s.StatusCode = &v
	return s
}

type PoiproductOperateResponseData struct {
	FailPoiList []*PoiproductOperateResponseDataFailPoiListItem `json:"fail_poi_list,omitempty" xml:"fail_poi_list,omitempty" type:"Repeated"`
	Description *string                                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PoiproductOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateResponseData) GoString() string {
	return s.String()
}

func (s *PoiproductOperateResponseData) SetFailPoiList(v []*PoiproductOperateResponseDataFailPoiListItem) *PoiproductOperateResponseData {
	s.FailPoiList = v
	return s
}

func (s *PoiproductOperateResponseData) SetDescription(v string) *PoiproductOperateResponseData {
	s.Description = &v
	return s
}

func (s *PoiproductOperateResponseData) SetErrorCode(v int32) *PoiproductOperateResponseData {
	s.ErrorCode = &v
	return s
}

type PoiproductOperateResponseDataFailPoiListItem struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	ExtId  *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId  *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s PoiproductOperateResponseDataFailPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateResponseDataFailPoiListItem) GoString() string {
	return s.String()
}

func (s *PoiproductOperateResponseDataFailPoiListItem) SetCode(v string) *PoiproductOperateResponseDataFailPoiListItem {
	s.Code = &v
	return s
}

func (s *PoiproductOperateResponseDataFailPoiListItem) SetExtId(v string) *PoiproductOperateResponseDataFailPoiListItem {
	s.ExtId = &v
	return s
}

func (s *PoiproductOperateResponseDataFailPoiListItem) SetPoiId(v int64) *PoiproductOperateResponseDataFailPoiListItem {
	s.PoiId = &v
	return s
}

func (s *PoiproductOperateResponseDataFailPoiListItem) SetReason(v string) *PoiproductOperateResponseDataFailPoiListItem {
	s.Reason = &v
	return s
}

type PoiproductOperateResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoiproductOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoiproductOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *PoiproductOperateResponseExtra) SetSubErrorCode(v int32) *PoiproductOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoiproductOperateResponseExtra) SetDescription(v string) *PoiproductOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *PoiproductOperateResponseExtra) SetErrorCode(v int32) *PoiproductOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoiproductOperateResponseExtra) SetLogid(v string) *PoiproductOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoiproductOperateResponseExtra) SetNow(v int64) *PoiproductOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *PoiproductOperateResponseExtra) SetSubDescription(v string) *PoiproductOperateResponseExtra {
	s.SubDescription = &v
	return s
}

type PoistockDetailRequest struct {
	RootLifeAccountId *int64                                 `json:"root_life_account_id,omitempty" xml:"root_life_account_id,omitempty" require:"true"`
	SkuPoiList        []*PoistockDetailRequestSkuPoiListItem `json:"sku_poi_list,omitempty" xml:"sku_poi_list,omitempty" require:"true" type:"Repeated"`
	Header            map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId         *int64                                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s PoistockDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailRequest) GoString() string {
	return s.String()
}

func (s *PoistockDetailRequest) SetRootLifeAccountId(v int64) *PoistockDetailRequest {
	s.RootLifeAccountId = &v
	return s
}

func (s *PoistockDetailRequest) SetSkuPoiList(v []*PoistockDetailRequestSkuPoiListItem) *PoistockDetailRequest {
	s.SkuPoiList = v
	return s
}

func (s *PoistockDetailRequest) SetHeader(v map[string]*string) *PoistockDetailRequest {
	s.Header = v
	return s
}

func (s *PoistockDetailRequest) SetAccessToken(v string) *PoistockDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *PoistockDetailRequest) SetAccountId(v int64) *PoistockDetailRequest {
	s.AccountId = &v
	return s
}

type PoistockDetailRequestSkuPoiListItem struct {
	PoiList   []*PoistockDetailRequestSkuPoiListItemPoiListItem `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	ProductId *int64                                            `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuId     *int64                                            `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s PoistockDetailRequestSkuPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailRequestSkuPoiListItem) GoString() string {
	return s.String()
}

func (s *PoistockDetailRequestSkuPoiListItem) SetPoiList(v []*PoistockDetailRequestSkuPoiListItemPoiListItem) *PoistockDetailRequestSkuPoiListItem {
	s.PoiList = v
	return s
}

func (s *PoistockDetailRequestSkuPoiListItem) SetProductId(v int64) *PoistockDetailRequestSkuPoiListItem {
	s.ProductId = &v
	return s
}

func (s *PoistockDetailRequestSkuPoiListItem) SetSkuId(v int64) *PoistockDetailRequestSkuPoiListItem {
	s.SkuId = &v
	return s
}

type PoistockDetailRequestSkuPoiListItemPoiListItem struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s PoistockDetailRequestSkuPoiListItemPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailRequestSkuPoiListItemPoiListItem) GoString() string {
	return s.String()
}

func (s *PoistockDetailRequestSkuPoiListItemPoiListItem) SetPoiId(v int64) *PoistockDetailRequestSkuPoiListItemPoiListItem {
	s.PoiId = &v
	return s
}

func (s *PoistockDetailRequestSkuPoiListItemPoiListItem) SetExtId(v string) *PoistockDetailRequestSkuPoiListItemPoiListItem {
	s.ExtId = &v
	return s
}

type PoistockDetailResponse struct {
	BaseResp *PoistockDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *PoistockDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoistockDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoistockDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponse) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponse) SetBaseResp(v *PoistockDetailResponseBaseResp) *PoistockDetailResponse {
	s.BaseResp = v
	return s
}

func (s *PoistockDetailResponse) SetData(v *PoistockDetailResponseData) *PoistockDetailResponse {
	s.Data = v
	return s
}

func (s *PoistockDetailResponse) SetExtra(v *PoistockDetailResponseExtra) *PoistockDetailResponse {
	s.Extra = v
	return s
}

type PoistockDetailResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s PoistockDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponseBaseResp) SetExtra(v map[string]*string) *PoistockDetailResponseBaseResp {
	s.Extra = v
	return s
}

func (s *PoistockDetailResponseBaseResp) SetStatusCode(v int32) *PoistockDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoistockDetailResponseBaseResp) SetStatusMessage(v string) *PoistockDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type PoistockDetailResponseData struct {
	ErrorCode            *int32                                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	SkuIdPoiStockListMap map[int64][]*PoistockDetailResponseDataSkuIdPoiStockListMapValueItem `json:"sku_id_poi_stock_list_map,omitempty" xml:"sku_id_poi_stock_list_map,omitempty" require:"true"`
	Description          *string                                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoistockDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponseData) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponseData) SetErrorCode(v int32) *PoistockDetailResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PoistockDetailResponseData) SetSkuIdPoiStockListMap(v map[int64][]*PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) *PoistockDetailResponseData {
	s.SkuIdPoiStockListMap = v
	return s
}

func (s *PoistockDetailResponseData) SetDescription(v string) *PoistockDetailResponseData {
	s.Description = &v
	return s
}

type PoistockDetailResponseDataSkuIdPoiStockListMapValueItem struct {
	AvailQty *int64                                                      `json:"AvailQty,omitempty" xml:"AvailQty,omitempty" require:"true"`
	Poi      *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi `json:"Poi,omitempty" xml:"Poi,omitempty" require:"true"`
	PoiName  *string                                                     `json:"PoiName,omitempty" xml:"PoiName,omitempty" require:"true"`
	StockQty *int64                                                      `json:"StockQty,omitempty" xml:"StockQty,omitempty" require:"true"`
}

func (s PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) SetAvailQty(v int64) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem {
	s.AvailQty = &v
	return s
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) SetPoi(v *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem {
	s.Poi = v
	return s
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) SetPoiName(v string) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem {
	s.PoiName = &v
	return s
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem) SetStockQty(v int64) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItem {
	s.StockQty = &v
	return s
}

type PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi) SetPoiId(v int64) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi {
	s.PoiId = &v
	return s
}

func (s *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi) SetExtId(v string) *PoistockDetailResponseDataSkuIdPoiStockListMapValueItemPoi {
	s.ExtId = &v
	return s
}

type PoistockDetailResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoistockDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoistockDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *PoistockDetailResponseExtra) SetSubErrorCode(v int32) *PoistockDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoistockDetailResponseExtra) SetDescription(v string) *PoistockDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *PoistockDetailResponseExtra) SetErrorCode(v int32) *PoistockDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoistockDetailResponseExtra) SetLogid(v string) *PoistockDetailResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoistockDetailResponseExtra) SetNow(v int64) *PoistockDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *PoistockDetailResponseExtra) SetSubDescription(v string) *PoistockDetailResponseExtra {
	s.SubDescription = &v
	return s
}

type PoistockSaveRequest struct {
	PoiStockList      []*PoistockSaveRequestPoiStockListItem `json:"poi_stock_list,omitempty" xml:"poi_stock_list,omitempty" require:"true" type:"Repeated"`
	RootLifeAccountId *int64                                 `json:"root_life_account_id,omitempty" xml:"root_life_account_id,omitempty" require:"true"`
	AccountId         *int64                                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header            map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PoistockSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveRequest) GoString() string {
	return s.String()
}

func (s *PoistockSaveRequest) SetPoiStockList(v []*PoistockSaveRequestPoiStockListItem) *PoistockSaveRequest {
	s.PoiStockList = v
	return s
}

func (s *PoistockSaveRequest) SetRootLifeAccountId(v int64) *PoistockSaveRequest {
	s.RootLifeAccountId = &v
	return s
}

func (s *PoistockSaveRequest) SetAccountId(v int64) *PoistockSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PoistockSaveRequest) SetHeader(v map[string]*string) *PoistockSaveRequest {
	s.Header = v
	return s
}

func (s *PoistockSaveRequest) SetAccessToken(v string) *PoistockSaveRequest {
	s.AccessToken = &v
	return s
}

type PoistockSaveRequestPoiStockListItem struct {
	StockQty  *int64                                  `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	Poi       *PoistockSaveRequestPoiStockListItemPoi `json:"poi,omitempty" xml:"poi,omitempty"`
	ProductId *int64                                  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuId     *int64                                  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s PoistockSaveRequestPoiStockListItem) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveRequestPoiStockListItem) GoString() string {
	return s.String()
}

func (s *PoistockSaveRequestPoiStockListItem) SetStockQty(v int64) *PoistockSaveRequestPoiStockListItem {
	s.StockQty = &v
	return s
}

func (s *PoistockSaveRequestPoiStockListItem) SetPoi(v *PoistockSaveRequestPoiStockListItemPoi) *PoistockSaveRequestPoiStockListItem {
	s.Poi = v
	return s
}

func (s *PoistockSaveRequestPoiStockListItem) SetProductId(v int64) *PoistockSaveRequestPoiStockListItem {
	s.ProductId = &v
	return s
}

func (s *PoistockSaveRequestPoiStockListItem) SetSkuId(v int64) *PoistockSaveRequestPoiStockListItem {
	s.SkuId = &v
	return s
}

type PoistockSaveRequestPoiStockListItemPoi struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s PoistockSaveRequestPoiStockListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveRequestPoiStockListItemPoi) GoString() string {
	return s.String()
}

func (s *PoistockSaveRequestPoiStockListItemPoi) SetExtId(v string) *PoistockSaveRequestPoiStockListItemPoi {
	s.ExtId = &v
	return s
}

func (s *PoistockSaveRequestPoiStockListItemPoi) SetPoiId(v int64) *PoistockSaveRequestPoiStockListItemPoi {
	s.PoiId = &v
	return s
}

type PoistockSaveResponse struct {
	BaseResp *PoistockSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *PoistockSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *PoistockSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PoistockSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponse) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponse) SetBaseResp(v *PoistockSaveResponseBaseResp) *PoistockSaveResponse {
	s.BaseResp = v
	return s
}

func (s *PoistockSaveResponse) SetData(v *PoistockSaveResponseData) *PoistockSaveResponse {
	s.Data = v
	return s
}

func (s *PoistockSaveResponse) SetExtra(v *PoistockSaveResponseExtra) *PoistockSaveResponse {
	s.Extra = v
	return s
}

type PoistockSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s PoistockSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponseBaseResp) SetStatusCode(v int32) *PoistockSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *PoistockSaveResponseBaseResp) SetStatusMessage(v string) *PoistockSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *PoistockSaveResponseBaseResp) SetExtra(v map[string]*string) *PoistockSaveResponseBaseResp {
	s.Extra = v
	return s
}

type PoistockSaveResponseData struct {
	ErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	FailPoiList []*PoistockSaveResponseDataFailPoiListItem `json:"fail_poi_list,omitempty" xml:"fail_poi_list,omitempty" type:"Repeated"`
	Description *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PoistockSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponseData) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponseData) SetErrorCode(v int32) *PoistockSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PoistockSaveResponseData) SetFailPoiList(v []*PoistockSaveResponseDataFailPoiListItem) *PoistockSaveResponseData {
	s.FailPoiList = v
	return s
}

func (s *PoistockSaveResponseData) SetDescription(v string) *PoistockSaveResponseData {
	s.Description = &v
	return s
}

type PoistockSaveResponseDataFailPoiListItem struct {
	Poi       *PoistockSaveResponseDataFailPoiListItemPoi `json:"poi,omitempty" xml:"poi,omitempty"`
	ProductId *int64                                      `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Reason    *string                                     `json:"reason,omitempty" xml:"reason,omitempty"`
	SkuId     *int64                                      `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Code      *string                                     `json:"code,omitempty" xml:"code,omitempty"`
}

func (s PoistockSaveResponseDataFailPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponseDataFailPoiListItem) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponseDataFailPoiListItem) SetPoi(v *PoistockSaveResponseDataFailPoiListItemPoi) *PoistockSaveResponseDataFailPoiListItem {
	s.Poi = v
	return s
}

func (s *PoistockSaveResponseDataFailPoiListItem) SetProductId(v int64) *PoistockSaveResponseDataFailPoiListItem {
	s.ProductId = &v
	return s
}

func (s *PoistockSaveResponseDataFailPoiListItem) SetReason(v string) *PoistockSaveResponseDataFailPoiListItem {
	s.Reason = &v
	return s
}

func (s *PoistockSaveResponseDataFailPoiListItem) SetSkuId(v int64) *PoistockSaveResponseDataFailPoiListItem {
	s.SkuId = &v
	return s
}

func (s *PoistockSaveResponseDataFailPoiListItem) SetCode(v string) *PoistockSaveResponseDataFailPoiListItem {
	s.Code = &v
	return s
}

type PoistockSaveResponseDataFailPoiListItemPoi struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s PoistockSaveResponseDataFailPoiListItemPoi) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponseDataFailPoiListItemPoi) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponseDataFailPoiListItemPoi) SetPoiId(v int64) *PoistockSaveResponseDataFailPoiListItemPoi {
	s.PoiId = &v
	return s
}

func (s *PoistockSaveResponseDataFailPoiListItemPoi) SetExtId(v string) *PoistockSaveResponseDataFailPoiListItemPoi {
	s.ExtId = &v
	return s
}

type PoistockSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PoistockSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PoistockSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PoistockSaveResponseExtra) SetSubErrorCode(v int32) *PoistockSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PoistockSaveResponseExtra) SetDescription(v string) *PoistockSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *PoistockSaveResponseExtra) SetErrorCode(v int32) *PoistockSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PoistockSaveResponseExtra) SetLogid(v string) *PoistockSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PoistockSaveResponseExtra) SetNow(v int64) *PoistockSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PoistockSaveResponseExtra) SetSubDescription(v string) *PoistockSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type PostingBindVideoRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty"`
	VideoId     *string            `json:"video_id,omitempty" xml:"video_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PostingBindVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s PostingBindVideoRequest) GoString() string {
	return s.String()
}

func (s *PostingBindVideoRequest) SetAccessToken(v string) *PostingBindVideoRequest {
	s.AccessToken = &v
	return s
}

func (s *PostingBindVideoRequest) SetOpenId(v string) *PostingBindVideoRequest {
	s.OpenId = &v
	return s
}

func (s *PostingBindVideoRequest) SetTaskId(v string) *PostingBindVideoRequest {
	s.TaskId = &v
	return s
}

func (s *PostingBindVideoRequest) SetVideoId(v string) *PostingBindVideoRequest {
	s.VideoId = &v
	return s
}

func (s *PostingBindVideoRequest) SetHeader(v map[string]*string) *PostingBindVideoRequest {
	s.Header = v
	return s
}

type PostingBindVideoResponse struct {
	Data  *PostingBindVideoResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *PostingBindVideoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PostingBindVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s PostingBindVideoResponse) GoString() string {
	return s.String()
}

func (s *PostingBindVideoResponse) SetData(v *PostingBindVideoResponseData) *PostingBindVideoResponse {
	s.Data = v
	return s
}

func (s *PostingBindVideoResponse) SetExtra(v *PostingBindVideoResponseExtra) *PostingBindVideoResponse {
	s.Extra = v
	return s
}

type PostingBindVideoResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PostingBindVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s PostingBindVideoResponseData) GoString() string {
	return s.String()
}

func (s *PostingBindVideoResponseData) SetGwErrorCode(v int32) *PostingBindVideoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PostingBindVideoResponseData) SetGwDescription(v string) *PostingBindVideoResponseData {
	s.GwDescription = &v
	return s
}

type PostingBindVideoResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PostingBindVideoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PostingBindVideoResponseExtra) GoString() string {
	return s.String()
}

func (s *PostingBindVideoResponseExtra) SetDescription(v string) *PostingBindVideoResponseExtra {
	s.Description = &v
	return s
}

func (s *PostingBindVideoResponseExtra) SetSubErrorCode(v int32) *PostingBindVideoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PostingBindVideoResponseExtra) SetSubDescription(v string) *PostingBindVideoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PostingBindVideoResponseExtra) SetLogid(v string) *PostingBindVideoResponseExtra {
	s.Logid = &v
	return s
}

func (s *PostingBindVideoResponseExtra) SetNow(v int64) *PostingBindVideoResponseExtra {
	s.Now = &v
	return s
}

func (s *PostingBindVideoResponseExtra) SetErrorCode(v int32) *PostingBindVideoResponseExtra {
	s.ErrorCode = &v
	return s
}

type PostingCreateRequest struct {
	TaskName      *string                            `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	TaskCondition *PostingCreateRequestTaskCondition `json:"task_condition,omitempty" xml:"task_condition,omitempty" require:"true"`
	StartTime     *int64                             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	Header        map[string]*string                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	EndTime       *int64                             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s PostingCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s PostingCreateRequest) GoString() string {
	return s.String()
}

func (s *PostingCreateRequest) SetTaskName(v string) *PostingCreateRequest {
	s.TaskName = &v
	return s
}

func (s *PostingCreateRequest) SetTaskCondition(v *PostingCreateRequestTaskCondition) *PostingCreateRequest {
	s.TaskCondition = v
	return s
}

func (s *PostingCreateRequest) SetStartTime(v int64) *PostingCreateRequest {
	s.StartTime = &v
	return s
}

func (s *PostingCreateRequest) SetHeader(v map[string]*string) *PostingCreateRequest {
	s.Header = v
	return s
}

func (s *PostingCreateRequest) SetAccessToken(v string) *PostingCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *PostingCreateRequest) SetEndTime(v int64) *PostingCreateRequest {
	s.EndTime = &v
	return s
}

type PostingCreateRequestTaskCondition struct {
	MinValue  *int64  `json:"min_value,omitempty" xml:"min_value,omitempty" require:"true"`
	MaxValue  *int64  `json:"max_value,omitempty" xml:"max_value,omitempty" require:"true"`
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty" require:"true"`
}

func (s PostingCreateRequestTaskCondition) String() string {
	return tea.Prettify(s)
}

func (s PostingCreateRequestTaskCondition) GoString() string {
	return s.String()
}

func (s *PostingCreateRequestTaskCondition) SetMinValue(v int64) *PostingCreateRequestTaskCondition {
	s.MinValue = &v
	return s
}

func (s *PostingCreateRequestTaskCondition) SetMaxValue(v int64) *PostingCreateRequestTaskCondition {
	s.MaxValue = &v
	return s
}

func (s *PostingCreateRequestTaskCondition) SetCondition(v string) *PostingCreateRequestTaskCondition {
	s.Condition = &v
	return s
}

type PostingCreateResponse struct {
	Extra *PostingCreateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PostingCreateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PostingCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s PostingCreateResponse) GoString() string {
	return s.String()
}

func (s *PostingCreateResponse) SetExtra(v *PostingCreateResponseExtra) *PostingCreateResponse {
	s.Extra = v
	return s
}

func (s *PostingCreateResponse) SetData(v *PostingCreateResponseData) *PostingCreateResponse {
	s.Data = v
	return s
}

type PostingCreateResponseData struct {
	TaskId        *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	TaskStatus    *int32  `json:"task_status,omitempty" xml:"task_status,omitempty"`
}

func (s PostingCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s PostingCreateResponseData) GoString() string {
	return s.String()
}

func (s *PostingCreateResponseData) SetTaskId(v string) *PostingCreateResponseData {
	s.TaskId = &v
	return s
}

func (s *PostingCreateResponseData) SetGwErrorCode(v int32) *PostingCreateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PostingCreateResponseData) SetGwDescription(v string) *PostingCreateResponseData {
	s.GwDescription = &v
	return s
}

func (s *PostingCreateResponseData) SetTaskStatus(v int32) *PostingCreateResponseData {
	s.TaskStatus = &v
	return s
}

type PostingCreateResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PostingCreateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PostingCreateResponseExtra) GoString() string {
	return s.String()
}

func (s *PostingCreateResponseExtra) SetSubErrorCode(v int32) *PostingCreateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PostingCreateResponseExtra) SetSubDescription(v string) *PostingCreateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PostingCreateResponseExtra) SetLogid(v string) *PostingCreateResponseExtra {
	s.Logid = &v
	return s
}

func (s *PostingCreateResponseExtra) SetNow(v int64) *PostingCreateResponseExtra {
	s.Now = &v
	return s
}

func (s *PostingCreateResponseExtra) SetErrorCode(v int32) *PostingCreateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PostingCreateResponseExtra) SetDescription(v string) *PostingCreateResponseExtra {
	s.Description = &v
	return s
}

type PostingUserRequest struct {
	VideoId      *string            `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	TaskId       *string            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	TargetOpenId *string            `json:"target_open_id,omitempty" xml:"target_open_id,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PostingUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PostingUserRequest) GoString() string {
	return s.String()
}

func (s *PostingUserRequest) SetVideoId(v string) *PostingUserRequest {
	s.VideoId = &v
	return s
}

func (s *PostingUserRequest) SetTaskId(v string) *PostingUserRequest {
	s.TaskId = &v
	return s
}

func (s *PostingUserRequest) SetTargetOpenId(v string) *PostingUserRequest {
	s.TargetOpenId = &v
	return s
}

func (s *PostingUserRequest) SetHeader(v map[string]*string) *PostingUserRequest {
	s.Header = v
	return s
}

func (s *PostingUserRequest) SetAccessToken(v string) *PostingUserRequest {
	s.AccessToken = &v
	return s
}

type PostingUserResponse struct {
	Data  *PostingUserResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PostingUserResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PostingUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PostingUserResponse) GoString() string {
	return s.String()
}

func (s *PostingUserResponse) SetData(v *PostingUserResponseData) *PostingUserResponse {
	s.Data = v
	return s
}

func (s *PostingUserResponse) SetExtra(v *PostingUserResponseExtra) *PostingUserResponse {
	s.Extra = v
	return s
}

type PostingUserResponseData struct {
	Result        *bool   `json:"result,omitempty" xml:"result,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PostingUserResponseData) String() string {
	return tea.Prettify(s)
}

func (s PostingUserResponseData) GoString() string {
	return s.String()
}

func (s *PostingUserResponseData) SetResult(v bool) *PostingUserResponseData {
	s.Result = &v
	return s
}

func (s *PostingUserResponseData) SetGwErrorCode(v int32) *PostingUserResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PostingUserResponseData) SetGwDescription(v string) *PostingUserResponseData {
	s.GwDescription = &v
	return s
}

type PostingUserResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s PostingUserResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PostingUserResponseExtra) GoString() string {
	return s.String()
}

func (s *PostingUserResponseExtra) SetNow(v int64) *PostingUserResponseExtra {
	s.Now = &v
	return s
}

func (s *PostingUserResponseExtra) SetErrorCode(v int32) *PostingUserResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PostingUserResponseExtra) SetDescription(v string) *PostingUserResponseExtra {
	s.Description = &v
	return s
}

func (s *PostingUserResponseExtra) SetSubErrorCode(v int32) *PostingUserResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PostingUserResponseExtra) SetSubDescription(v string) *PostingUserResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PostingUserResponseExtra) SetLogid(v string) *PostingUserResponseExtra {
	s.Logid = &v
	return s
}

type PresaleCouponSaveBookCalendarStockRequest struct {
	AccountId               *string                                                           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BookCalendarStockConfig *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig `json:"book_calendar_stock_config,omitempty" xml:"book_calendar_stock_config,omitempty"`
	ProductId               *string                                                           `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Header                  map[string]*string                                                `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken             *string                                                           `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PresaleCouponSaveBookCalendarStockRequest) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockRequest) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockRequest) SetAccountId(v string) *PresaleCouponSaveBookCalendarStockRequest {
	s.AccountId = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequest) SetBookCalendarStockConfig(v *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) *PresaleCouponSaveBookCalendarStockRequest {
	s.BookCalendarStockConfig = v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequest) SetProductId(v string) *PresaleCouponSaveBookCalendarStockRequest {
	s.ProductId = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequest) SetHeader(v map[string]*string) *PresaleCouponSaveBookCalendarStockRequest {
	s.Header = v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequest) SetAccessToken(v string) *PresaleCouponSaveBookCalendarStockRequest {
	s.AccessToken = &v
	return s
}

type PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig struct {
	DatePeriodList    []*PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem `json:"date_period_list,omitempty" xml:"date_period_list,omitempty" type:"Repeated"`
	StockQtyLimitType *int32                                                                                `json:"stock_qty_limit_type,omitempty" xml:"stock_qty_limit_type,omitempty"`
	AvailableStock    *int64                                                                                `json:"available_stock,omitempty" xml:"available_stock,omitempty"`
	CalendarStatus    *int32                                                                                `json:"calendar_status,omitempty" xml:"calendar_status,omitempty"`
}

func (s PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) SetDatePeriodList(v []*PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig {
	s.DatePeriodList = v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) SetStockQtyLimitType(v int32) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig {
	s.StockQtyLimitType = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) SetAvailableStock(v int64) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig {
	s.AvailableStock = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig) SetCalendarStatus(v int32) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfig {
	s.CalendarStatus = &v
	return s
}

type PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem struct {
	EndDate    *string  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate  *string  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
}

func (s PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) SetEndDate(v string) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem {
	s.EndDate = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) SetStartDate(v string) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem {
	s.StartDate = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem) SetDaysOfWeek(v []*int32) *PresaleCouponSaveBookCalendarStockRequestBookCalendarStockConfigDatePeriodListItem {
	s.DaysOfWeek = v
	return s
}

type PresaleCouponSaveBookCalendarStockResponse struct {
	Extra *PresaleCouponSaveBookCalendarStockResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PresaleCouponSaveBookCalendarStockResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PresaleCouponSaveBookCalendarStockResponse) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockResponse) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockResponse) SetExtra(v *PresaleCouponSaveBookCalendarStockResponseExtra) *PresaleCouponSaveBookCalendarStockResponse {
	s.Extra = v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponse) SetData(v *PresaleCouponSaveBookCalendarStockResponseData) *PresaleCouponSaveBookCalendarStockResponse {
	s.Data = v
	return s
}

type PresaleCouponSaveBookCalendarStockResponseData struct {
	Descripiton   *string `json:"descripiton,omitempty" xml:"descripiton,omitempty"`
	ErrorCode     *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PresaleCouponSaveBookCalendarStockResponseData) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockResponseData) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockResponseData) SetDescripiton(v string) *PresaleCouponSaveBookCalendarStockResponseData {
	s.Descripiton = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseData) SetErrorCode(v string) *PresaleCouponSaveBookCalendarStockResponseData {
	s.ErrorCode = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseData) SetGwDescription(v string) *PresaleCouponSaveBookCalendarStockResponseData {
	s.GwDescription = &v
	return s
}

type PresaleCouponSaveBookCalendarStockResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PresaleCouponSaveBookCalendarStockResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PresaleCouponSaveBookCalendarStockResponseExtra) GoString() string {
	return s.String()
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetLogid(v string) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.Logid = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetNow(v int64) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.Now = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetSubDescription(v string) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetSubErrorCode(v int32) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetDescription(v string) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.Description = &v
	return s
}

func (s *PresaleCouponSaveBookCalendarStockResponseExtra) SetErrorCode(v int32) *PresaleCouponSaveBookCalendarStockResponseExtra {
	s.ErrorCode = &v
	return s
}

type PresaleRateplanSaveRequest struct {
	AccessToken *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RatePlan    *PresaleRateplanSaveRequestRatePlan `json:"rate_plan,omitempty" xml:"rate_plan,omitempty" require:"true"`
	AccountId   *string                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PresaleRateplanSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveRequest) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveRequest) SetAccessToken(v string) *PresaleRateplanSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *PresaleRateplanSaveRequest) SetRatePlan(v *PresaleRateplanSaveRequestRatePlan) *PresaleRateplanSaveRequest {
	s.RatePlan = v
	return s
}

func (s *PresaleRateplanSaveRequest) SetAccountId(v string) *PresaleRateplanSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PresaleRateplanSaveRequest) SetHeader(v map[string]*string) *PresaleRateplanSaveRequest {
	s.Header = v
	return s
}

type PresaleRateplanSaveRequestRatePlan struct {
	HotelId *string                                        `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Rooms   []*PresaleRateplanSaveRequestRatePlanRoomsItem `json:"rooms,omitempty" xml:"rooms,omitempty" require:"true" type:"Repeated"`
}

func (s PresaleRateplanSaveRequestRatePlan) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveRequestRatePlan) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveRequestRatePlan) SetHotelId(v string) *PresaleRateplanSaveRequestRatePlan {
	s.HotelId = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlan) SetRooms(v []*PresaleRateplanSaveRequestRatePlanRoomsItem) *PresaleRateplanSaveRequestRatePlan {
	s.Rooms = v
	return s
}

type PresaleRateplanSaveRequestRatePlanRoomsItem struct {
	RoomId    *string                                                     `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	RatePlans []*PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem `json:"rate_plans,omitempty" xml:"rate_plans,omitempty" require:"true" type:"Repeated"`
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItem) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItem) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItem) SetRoomId(v string) *PresaleRateplanSaveRequestRatePlanRoomsItem {
	s.RoomId = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItem) SetRatePlans(v []*PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) *PresaleRateplanSaveRequestRatePlanRoomsItem {
	s.RatePlans = v
	return s
}

type PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem struct {
	HourlyRoomDetail *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail `json:"hourly_room_detail,omitempty" xml:"hourly_room_detail,omitempty"`
	SalesType        *int                                                                      `json:"sales_type,omitempty" xml:"sales_type,omitempty"`
	Currency         *string                                                                   `json:"currency,omitempty" xml:"currency,omitempty"`
	RatePlanId       *string                                                                   `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	Active           *bool                                                                     `json:"active,omitempty" xml:"active,omitempty"`
	OutRatePlanId    *string                                                                   `json:"out_rate_plan_id,omitempty" xml:"out_rate_plan_id,omitempty" require:"true"`
	RatePlanName     *string                                                                   `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty" require:"true"`
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetHourlyRoomDetail(v *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.HourlyRoomDetail = v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetSalesType(v int) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.SalesType = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetCurrency(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Currency = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetRatePlanId(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.RatePlanId = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetActive(v bool) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.Active = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetOutRatePlanId(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.OutRatePlanId = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem) SetRatePlanName(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItem {
	s.RatePlanName = &v
	return s
}

type PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail struct {
	UsageDuration   *int64  `json:"usage_duration,omitempty" xml:"usage_duration,omitempty" require:"true"`
	EarliestCheckIn *string `json:"earliest_check_in,omitempty" xml:"earliest_check_in,omitempty" require:"true"`
	LatestCheckOut  *string `json:"latest_check_out,omitempty" xml:"latest_check_out,omitempty" require:"true"`
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetUsageDuration(v int64) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.UsageDuration = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetEarliestCheckIn(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.EarliestCheckIn = &v
	return s
}

func (s *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail) SetLatestCheckOut(v string) *PresaleRateplanSaveRequestRatePlanRoomsItemRatePlansItemHourlyRoomDetail {
	s.LatestCheckOut = &v
	return s
}

type PresaleRateplanSaveResponse struct {
	Data  *PresaleRateplanSaveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PresaleRateplanSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PresaleRateplanSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveResponse) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveResponse) SetData(v *PresaleRateplanSaveResponseData) *PresaleRateplanSaveResponse {
	s.Data = v
	return s
}

func (s *PresaleRateplanSaveResponse) SetExtra(v *PresaleRateplanSaveResponseExtra) *PresaleRateplanSaveResponse {
	s.Extra = v
	return s
}

type PresaleRateplanSaveResponseData struct {
	RatePlanMap   []*PresaleRateplanSaveResponseDataRatePlanMapItem `json:"rate_plan_map,omitempty" xml:"rate_plan_map,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PresaleRateplanSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveResponseData) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveResponseData) SetRatePlanMap(v []*PresaleRateplanSaveResponseDataRatePlanMapItem) *PresaleRateplanSaveResponseData {
	s.RatePlanMap = v
	return s
}

func (s *PresaleRateplanSaveResponseData) SetGwErrorCode(v int32) *PresaleRateplanSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PresaleRateplanSaveResponseData) SetGwDescription(v string) *PresaleRateplanSaveResponseData {
	s.GwDescription = &v
	return s
}

type PresaleRateplanSaveResponseDataRatePlanMapItem struct {
	OutRatePlanId *string `json:"out_rate_plan_id,omitempty" xml:"out_rate_plan_id,omitempty" require:"true"`
	RatePlanId    *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message       *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s PresaleRateplanSaveResponseDataRatePlanMapItem) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveResponseDataRatePlanMapItem) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveResponseDataRatePlanMapItem) SetOutRatePlanId(v string) *PresaleRateplanSaveResponseDataRatePlanMapItem {
	s.OutRatePlanId = &v
	return s
}

func (s *PresaleRateplanSaveResponseDataRatePlanMapItem) SetRatePlanId(v string) *PresaleRateplanSaveResponseDataRatePlanMapItem {
	s.RatePlanId = &v
	return s
}

func (s *PresaleRateplanSaveResponseDataRatePlanMapItem) SetCode(v string) *PresaleRateplanSaveResponseDataRatePlanMapItem {
	s.Code = &v
	return s
}

func (s *PresaleRateplanSaveResponseDataRatePlanMapItem) SetMessage(v string) *PresaleRateplanSaveResponseDataRatePlanMapItem {
	s.Message = &v
	return s
}

type PresaleRateplanSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s PresaleRateplanSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PresaleRateplanSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PresaleRateplanSaveResponseExtra) SetDescription(v string) *PresaleRateplanSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *PresaleRateplanSaveResponseExtra) SetErrorCode(v int32) *PresaleRateplanSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PresaleRateplanSaveResponseExtra) SetLogid(v string) *PresaleRateplanSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PresaleRateplanSaveResponseExtra) SetNow(v int64) *PresaleRateplanSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PresaleRateplanSaveResponseExtra) SetSubDescription(v string) *PresaleRateplanSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PresaleRateplanSaveResponseExtra) SetSubErrorCode(v int32) *PresaleRateplanSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type PricePushRequest struct {
	ApplicableDate     *PricePushRequestApplicableDate          `json:"applicable_date,omitempty" xml:"applicable_date,omitempty" require:"true"`
	ApplicableResource *PricePushRequestApplicableResource      `json:"applicable_resource,omitempty" xml:"applicable_resource,omitempty" require:"true"`
	Header             map[string]*string                       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PriceMemberLevels  []*PricePushRequestPriceMemberLevelsItem `json:"price_member_levels,omitempty" xml:"price_member_levels,omitempty" require:"true" type:"Repeated"`
	PromotionId        *string                                  `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	AccountId          *string                                  `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s PricePushRequest) String() string {
	return tea.Prettify(s)
}

func (s PricePushRequest) GoString() string {
	return s.String()
}

func (s *PricePushRequest) SetApplicableDate(v *PricePushRequestApplicableDate) *PricePushRequest {
	s.ApplicableDate = v
	return s
}

func (s *PricePushRequest) SetApplicableResource(v *PricePushRequestApplicableResource) *PricePushRequest {
	s.ApplicableResource = v
	return s
}

func (s *PricePushRequest) SetHeader(v map[string]*string) *PricePushRequest {
	s.Header = v
	return s
}

func (s *PricePushRequest) SetAccessToken(v string) *PricePushRequest {
	s.AccessToken = &v
	return s
}

func (s *PricePushRequest) SetPriceMemberLevels(v []*PricePushRequestPriceMemberLevelsItem) *PricePushRequest {
	s.PriceMemberLevels = v
	return s
}

func (s *PricePushRequest) SetPromotionId(v string) *PricePushRequest {
	s.PromotionId = &v
	return s
}

func (s *PricePushRequest) SetAccountId(v string) *PricePushRequest {
	s.AccountId = &v
	return s
}

type PricePushRequestApplicableDate struct {
	EndDays    *string  `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
	StartDays  *string  `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
}

func (s PricePushRequestApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s PricePushRequestApplicableDate) GoString() string {
	return s.String()
}

func (s *PricePushRequestApplicableDate) SetEndDays(v string) *PricePushRequestApplicableDate {
	s.EndDays = &v
	return s
}

func (s *PricePushRequestApplicableDate) SetStartDays(v string) *PricePushRequestApplicableDate {
	s.StartDays = &v
	return s
}

func (s *PricePushRequestApplicableDate) SetDaysOfWeek(v []*int32) *PricePushRequestApplicableDate {
	s.DaysOfWeek = v
	return s
}

type PricePushRequestApplicableResource struct {
	Resources []*PricePushRequestApplicableResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s PricePushRequestApplicableResource) String() string {
	return tea.Prettify(s)
}

func (s PricePushRequestApplicableResource) GoString() string {
	return s.String()
}

func (s *PricePushRequestApplicableResource) SetResources(v []*PricePushRequestApplicableResourceResourcesItem) *PricePushRequestApplicableResource {
	s.Resources = v
	return s
}

type PricePushRequestApplicableResourceResourcesItem struct {
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
}

func (s PricePushRequestApplicableResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s PricePushRequestApplicableResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *PricePushRequestApplicableResourceResourcesItem) SetRatePlanId(v string) *PricePushRequestApplicableResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

func (s *PricePushRequestApplicableResourceResourcesItem) SetRatePlanStatus(v int32) *PricePushRequestApplicableResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

type PricePushRequestPriceMemberLevelsItem struct {
	Price           *int64 `json:"price,omitempty" xml:"price,omitempty" require:"true"`
	BindMemberLevel *int64 `json:"bind_member_level,omitempty" xml:"bind_member_level,omitempty"`
}

func (s PricePushRequestPriceMemberLevelsItem) String() string {
	return tea.Prettify(s)
}

func (s PricePushRequestPriceMemberLevelsItem) GoString() string {
	return s.String()
}

func (s *PricePushRequestPriceMemberLevelsItem) SetPrice(v int64) *PricePushRequestPriceMemberLevelsItem {
	s.Price = &v
	return s
}

func (s *PricePushRequestPriceMemberLevelsItem) SetBindMemberLevel(v int64) *PricePushRequestPriceMemberLevelsItem {
	s.BindMemberLevel = &v
	return s
}

type PricePushResponse struct {
	Data  *PricePushResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PricePushResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PricePushResponse) String() string {
	return tea.Prettify(s)
}

func (s PricePushResponse) GoString() string {
	return s.String()
}

func (s *PricePushResponse) SetData(v *PricePushResponseData) *PricePushResponse {
	s.Data = v
	return s
}

func (s *PricePushResponse) SetExtra(v *PricePushResponseExtra) *PricePushResponse {
	s.Extra = v
	return s
}

type PricePushResponseData struct {
	Status        *int                                      `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	MessageDetail []*PricePushResponseDataMessageDetailItem `json:"message_detail,omitempty" xml:"message_detail,omitempty" type:"Repeated"`
	PromotionId   *string                                   `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PricePushResponseData) String() string {
	return tea.Prettify(s)
}

func (s PricePushResponseData) GoString() string {
	return s.String()
}

func (s *PricePushResponseData) SetStatus(v int) *PricePushResponseData {
	s.Status = &v
	return s
}

func (s *PricePushResponseData) SetMessageDetail(v []*PricePushResponseDataMessageDetailItem) *PricePushResponseData {
	s.MessageDetail = v
	return s
}

func (s *PricePushResponseData) SetPromotionId(v string) *PricePushResponseData {
	s.PromotionId = &v
	return s
}

func (s *PricePushResponseData) SetGwErrorCode(v int32) *PricePushResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PricePushResponseData) SetGwDescription(v string) *PricePushResponseData {
	s.GwDescription = &v
	return s
}

type PricePushResponseDataMessageDetailItem struct {
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}

func (s PricePushResponseDataMessageDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PricePushResponseDataMessageDetailItem) GoString() string {
	return s.String()
}

func (s *PricePushResponseDataMessageDetailItem) SetRatePlanId(v string) *PricePushResponseDataMessageDetailItem {
	s.RatePlanId = &v
	return s
}

func (s *PricePushResponseDataMessageDetailItem) SetErrorMessage(v string) *PricePushResponseDataMessageDetailItem {
	s.ErrorMessage = &v
	return s
}

type PricePushResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PricePushResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PricePushResponseExtra) GoString() string {
	return s.String()
}

func (s *PricePushResponseExtra) SetSubErrorCode(v int32) *PricePushResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PricePushResponseExtra) SetDescription(v string) *PricePushResponseExtra {
	s.Description = &v
	return s
}

func (s *PricePushResponseExtra) SetErrorCode(v int32) *PricePushResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PricePushResponseExtra) SetLogid(v string) *PricePushResponseExtra {
	s.Logid = &v
	return s
}

func (s *PricePushResponseExtra) SetNow(v int64) *PricePushResponseExtra {
	s.Now = &v
	return s
}

func (s *PricePushResponseExtra) SetSubDescription(v string) *PricePushResponseExtra {
	s.SubDescription = &v
	return s
}

type PriceSaveRequest struct {
	HotelId     *string                     `json:"hotel_id,omitempty" xml:"hotel_id,omitempty" require:"true"`
	AccountId   *string                     `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Aris        []*PriceSaveRequestArisItem `json:"aris,omitempty" xml:"aris,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PriceSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequest) GoString() string {
	return s.String()
}

func (s *PriceSaveRequest) SetHotelId(v string) *PriceSaveRequest {
	s.HotelId = &v
	return s
}

func (s *PriceSaveRequest) SetAccountId(v string) *PriceSaveRequest {
	s.AccountId = &v
	return s
}

func (s *PriceSaveRequest) SetAris(v []*PriceSaveRequestArisItem) *PriceSaveRequest {
	s.Aris = v
	return s
}

func (s *PriceSaveRequest) SetHeader(v map[string]*string) *PriceSaveRequest {
	s.Header = v
	return s
}

func (s *PriceSaveRequest) SetAccessToken(v string) *PriceSaveRequest {
	s.AccessToken = &v
	return s
}

type PriceSaveRequestArisItem struct {
	Timerange        *PriceSaveRequestArisItemTimerange              `json:"timerange,omitempty" xml:"timerange,omitempty" require:"true"`
	BookTimeRules    *PriceSaveRequestArisItemBookTimeRules          `json:"book_time_rules,omitempty" xml:"book_time_rules,omitempty"`
	Meals            []*PriceSaveRequestArisItemMealsItem            `json:"meals,omitempty" xml:"meals,omitempty" type:"Repeated"`
	LengthOfStay     *int64                                          `json:"length_of_stay,omitempty" xml:"length_of_stay,omitempty"`
	RoomId           *string                                         `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Priority         *int                                            `json:"priority,omitempty" xml:"priority,omitempty"`
	DaysOfWeek       []*int                                          `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	CancelRules      []*PriceSaveRequestArisItemCancelRulesItem      `json:"cancel_rules,omitempty" xml:"cancel_rules,omitempty" type:"Repeated"`
	BookRules        *PriceSaveRequestArisItemBookRules              `json:"book_rules,omitempty" xml:"book_rules,omitempty"`
	LosRateBreakDown []*PriceSaveRequestArisItemLosRateBreakDownItem `json:"los_rate_break_down,omitempty" xml:"los_rate_break_down,omitempty" type:"Repeated"`
	RetailAmount     *int64                                          `json:"retail_amount,omitempty" xml:"retail_amount,omitempty"`
	OriginalAmount   *int64                                          `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	AmountBeforeTax  *int64                                          `json:"amount_before_tax,omitempty" xml:"amount_before_tax,omitempty"`
	StayRules        *PriceSaveRequestArisItemStayRules              `json:"stay_rules,omitempty" xml:"stay_rules,omitempty"`
	RatePlanId       *string                                         `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
}

func (s PriceSaveRequestArisItem) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItem) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItem) SetTimerange(v *PriceSaveRequestArisItemTimerange) *PriceSaveRequestArisItem {
	s.Timerange = v
	return s
}

func (s *PriceSaveRequestArisItem) SetBookTimeRules(v *PriceSaveRequestArisItemBookTimeRules) *PriceSaveRequestArisItem {
	s.BookTimeRules = v
	return s
}

func (s *PriceSaveRequestArisItem) SetMeals(v []*PriceSaveRequestArisItemMealsItem) *PriceSaveRequestArisItem {
	s.Meals = v
	return s
}

func (s *PriceSaveRequestArisItem) SetLengthOfStay(v int64) *PriceSaveRequestArisItem {
	s.LengthOfStay = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetRoomId(v string) *PriceSaveRequestArisItem {
	s.RoomId = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetPriority(v int) *PriceSaveRequestArisItem {
	s.Priority = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetDaysOfWeek(v []*int) *PriceSaveRequestArisItem {
	s.DaysOfWeek = v
	return s
}

func (s *PriceSaveRequestArisItem) SetCancelRules(v []*PriceSaveRequestArisItemCancelRulesItem) *PriceSaveRequestArisItem {
	s.CancelRules = v
	return s
}

func (s *PriceSaveRequestArisItem) SetBookRules(v *PriceSaveRequestArisItemBookRules) *PriceSaveRequestArisItem {
	s.BookRules = v
	return s
}

func (s *PriceSaveRequestArisItem) SetLosRateBreakDown(v []*PriceSaveRequestArisItemLosRateBreakDownItem) *PriceSaveRequestArisItem {
	s.LosRateBreakDown = v
	return s
}

func (s *PriceSaveRequestArisItem) SetRetailAmount(v int64) *PriceSaveRequestArisItem {
	s.RetailAmount = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetOriginalAmount(v int64) *PriceSaveRequestArisItem {
	s.OriginalAmount = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetAmountBeforeTax(v int64) *PriceSaveRequestArisItem {
	s.AmountBeforeTax = &v
	return s
}

func (s *PriceSaveRequestArisItem) SetStayRules(v *PriceSaveRequestArisItemStayRules) *PriceSaveRequestArisItem {
	s.StayRules = v
	return s
}

func (s *PriceSaveRequestArisItem) SetRatePlanId(v string) *PriceSaveRequestArisItem {
	s.RatePlanId = &v
	return s
}

type PriceSaveRequestArisItemBookRules struct {
	MidnightRoom     *PriceSaveRequestArisItemBookRulesMidnightRoom   `json:"midnight_room,omitempty" xml:"midnight_room,omitempty"`
	CheckInFrom      *string                                          `json:"check_in_from,omitempty" xml:"check_in_from,omitempty"`
	MaxLosDays       *int64                                           `json:"max_los_days,omitempty" xml:"max_los_days,omitempty"`
	ApplicablePeople []*int                                           `json:"applicable_people,omitempty" xml:"applicable_people,omitempty" type:"Repeated"`
	CheckOutTo       *string                                          `json:"check_out_to,omitempty" xml:"check_out_to,omitempty"`
	MaxQuantityLimt  *int64                                           `json:"max_quantity_limt,omitempty" xml:"max_quantity_limt,omitempty"`
	MinLosDays       *int64                                           `json:"min_los_days,omitempty" xml:"min_los_days,omitempty"`
	CheckInTo        *string                                          `json:"check_in_to,omitempty" xml:"check_in_to,omitempty"`
	MinAdvanceTime   *PriceSaveRequestArisItemBookRulesMinAdvanceTime `json:"min_advance_time,omitempty" xml:"min_advance_time,omitempty"`
	MaxAdvanceTime   *PriceSaveRequestArisItemBookRulesMaxAdvanceTime `json:"max_advance_time,omitempty" xml:"max_advance_time,omitempty"`
}

func (s PriceSaveRequestArisItemBookRules) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookRules) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookRules) SetMidnightRoom(v *PriceSaveRequestArisItemBookRulesMidnightRoom) *PriceSaveRequestArisItemBookRules {
	s.MidnightRoom = v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetCheckInFrom(v string) *PriceSaveRequestArisItemBookRules {
	s.CheckInFrom = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetMaxLosDays(v int64) *PriceSaveRequestArisItemBookRules {
	s.MaxLosDays = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetApplicablePeople(v []*int) *PriceSaveRequestArisItemBookRules {
	s.ApplicablePeople = v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetCheckOutTo(v string) *PriceSaveRequestArisItemBookRules {
	s.CheckOutTo = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetMaxQuantityLimt(v int64) *PriceSaveRequestArisItemBookRules {
	s.MaxQuantityLimt = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetMinLosDays(v int64) *PriceSaveRequestArisItemBookRules {
	s.MinLosDays = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetCheckInTo(v string) *PriceSaveRequestArisItemBookRules {
	s.CheckInTo = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetMinAdvanceTime(v *PriceSaveRequestArisItemBookRulesMinAdvanceTime) *PriceSaveRequestArisItemBookRules {
	s.MinAdvanceTime = v
	return s
}

func (s *PriceSaveRequestArisItemBookRules) SetMaxAdvanceTime(v *PriceSaveRequestArisItemBookRulesMaxAdvanceTime) *PriceSaveRequestArisItemBookRules {
	s.MaxAdvanceTime = v
	return s
}

type PriceSaveRequestArisItemBookRulesMaxAdvanceTime struct {
	Second *int32 `json:"second,omitempty" xml:"second,omitempty"`
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
}

func (s PriceSaveRequestArisItemBookRulesMaxAdvanceTime) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookRulesMaxAdvanceTime) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookRulesMaxAdvanceTime) SetSecond(v int32) *PriceSaveRequestArisItemBookRulesMaxAdvanceTime {
	s.Second = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMaxAdvanceTime) SetDay(v int32) *PriceSaveRequestArisItemBookRulesMaxAdvanceTime {
	s.Day = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMaxAdvanceTime) SetHour(v int32) *PriceSaveRequestArisItemBookRulesMaxAdvanceTime {
	s.Hour = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMaxAdvanceTime) SetMinute(v int32) *PriceSaveRequestArisItemBookRulesMaxAdvanceTime {
	s.Minute = &v
	return s
}

type PriceSaveRequestArisItemBookRulesMidnightRoom struct {
	IsMidnightRoom    *bool  `json:"is_midnight_room,omitempty" xml:"is_midnight_room,omitempty"`
	LatestBookingTime *int64 `json:"latest_booking_time,omitempty" xml:"latest_booking_time,omitempty"`
}

func (s PriceSaveRequestArisItemBookRulesMidnightRoom) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookRulesMidnightRoom) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookRulesMidnightRoom) SetIsMidnightRoom(v bool) *PriceSaveRequestArisItemBookRulesMidnightRoom {
	s.IsMidnightRoom = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMidnightRoom) SetLatestBookingTime(v int64) *PriceSaveRequestArisItemBookRulesMidnightRoom {
	s.LatestBookingTime = &v
	return s
}

type PriceSaveRequestArisItemBookRulesMinAdvanceTime struct {
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
	Second *int32 `json:"second,omitempty" xml:"second,omitempty"`
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
}

func (s PriceSaveRequestArisItemBookRulesMinAdvanceTime) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookRulesMinAdvanceTime) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookRulesMinAdvanceTime) SetHour(v int32) *PriceSaveRequestArisItemBookRulesMinAdvanceTime {
	s.Hour = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMinAdvanceTime) SetMinute(v int32) *PriceSaveRequestArisItemBookRulesMinAdvanceTime {
	s.Minute = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMinAdvanceTime) SetSecond(v int32) *PriceSaveRequestArisItemBookRulesMinAdvanceTime {
	s.Second = &v
	return s
}

func (s *PriceSaveRequestArisItemBookRulesMinAdvanceTime) SetDay(v int32) *PriceSaveRequestArisItemBookRulesMinAdvanceTime {
	s.Day = &v
	return s
}

type PriceSaveRequestArisItemBookTimeRules struct {
	IsTimeNextday *bool                                          `json:"is_time_nextday,omitempty" xml:"is_time_nextday,omitempty" require:"true"`
	TimeSpan      *PriceSaveRequestArisItemBookTimeRulesTimeSpan `json:"time_span,omitempty" xml:"time_span,omitempty" require:"true"`
}

func (s PriceSaveRequestArisItemBookTimeRules) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookTimeRules) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookTimeRules) SetIsTimeNextday(v bool) *PriceSaveRequestArisItemBookTimeRules {
	s.IsTimeNextday = &v
	return s
}

func (s *PriceSaveRequestArisItemBookTimeRules) SetTimeSpan(v *PriceSaveRequestArisItemBookTimeRulesTimeSpan) *PriceSaveRequestArisItemBookTimeRules {
	s.TimeSpan = v
	return s
}

type PriceSaveRequestArisItemBookTimeRulesTimeSpan struct {
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	To   *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s PriceSaveRequestArisItemBookTimeRulesTimeSpan) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemBookTimeRulesTimeSpan) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemBookTimeRulesTimeSpan) SetFrom(v string) *PriceSaveRequestArisItemBookTimeRulesTimeSpan {
	s.From = &v
	return s
}

func (s *PriceSaveRequestArisItemBookTimeRulesTimeSpan) SetTo(v string) *PriceSaveRequestArisItemBookTimeRulesTimeSpan {
	s.To = &v
	return s
}

type PriceSaveRequestArisItemCancelRulesItem struct {
	CancelOffsetTime *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime `json:"cancel_offset_time,omitempty" xml:"cancel_offset_time,omitempty"`
	CancelTimeType   *int                                                     `json:"cancel_time_type,omitempty" xml:"cancel_time_type,omitempty"`
	CancelType       *int                                                     `json:"cancel_type,omitempty" xml:"cancel_type,omitempty" require:"true"`
	CutType          *int                                                     `json:"cut_type,omitempty" xml:"cut_type,omitempty"`
	CutValue         *int64                                                   `json:"cut_value,omitempty" xml:"cut_value,omitempty"`
}

func (s PriceSaveRequestArisItemCancelRulesItem) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemCancelRulesItem) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemCancelRulesItem) SetCancelOffsetTime(v *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) *PriceSaveRequestArisItemCancelRulesItem {
	s.CancelOffsetTime = v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItem) SetCancelTimeType(v int) *PriceSaveRequestArisItemCancelRulesItem {
	s.CancelTimeType = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItem) SetCancelType(v int) *PriceSaveRequestArisItemCancelRulesItem {
	s.CancelType = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItem) SetCutType(v int) *PriceSaveRequestArisItemCancelRulesItem {
	s.CutType = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItem) SetCutValue(v int64) *PriceSaveRequestArisItemCancelRulesItem {
	s.CutValue = &v
	return s
}

type PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime struct {
	Hour   *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	Minute *int32 `json:"minute,omitempty" xml:"minute,omitempty"`
	Second *int32 `json:"second,omitempty" xml:"second,omitempty"`
	Day    *int32 `json:"day,omitempty" xml:"day,omitempty"`
}

func (s PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) SetHour(v int32) *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime {
	s.Hour = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) SetMinute(v int32) *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime {
	s.Minute = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) SetSecond(v int32) *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime {
	s.Second = &v
	return s
}

func (s *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime) SetDay(v int32) *PriceSaveRequestArisItemCancelRulesItemCancelOffsetTime {
	s.Day = &v
	return s
}

type PriceSaveRequestArisItemLosRateBreakDownItem struct {
	DayAmountBeforeTax *int64 `json:"day_amount_before_tax,omitempty" xml:"day_amount_before_tax,omitempty"`
	DayOriginalAmount  *int64 `json:"day_original_amount,omitempty" xml:"day_original_amount,omitempty"`
	DayRetailAmount    *int64 `json:"day_retail_amount,omitempty" xml:"day_retail_amount,omitempty"`
}

func (s PriceSaveRequestArisItemLosRateBreakDownItem) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemLosRateBreakDownItem) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemLosRateBreakDownItem) SetDayAmountBeforeTax(v int64) *PriceSaveRequestArisItemLosRateBreakDownItem {
	s.DayAmountBeforeTax = &v
	return s
}

func (s *PriceSaveRequestArisItemLosRateBreakDownItem) SetDayOriginalAmount(v int64) *PriceSaveRequestArisItemLosRateBreakDownItem {
	s.DayOriginalAmount = &v
	return s
}

func (s *PriceSaveRequestArisItemLosRateBreakDownItem) SetDayRetailAmount(v int64) *PriceSaveRequestArisItemLosRateBreakDownItem {
	s.DayRetailAmount = &v
	return s
}

type PriceSaveRequestArisItemMealsItem struct {
	Num  *int64 `json:"num,omitempty" xml:"num,omitempty" require:"true"`
	Type *int   `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s PriceSaveRequestArisItemMealsItem) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemMealsItem) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemMealsItem) SetNum(v int64) *PriceSaveRequestArisItemMealsItem {
	s.Num = &v
	return s
}

func (s *PriceSaveRequestArisItemMealsItem) SetType(v int) *PriceSaveRequestArisItemMealsItem {
	s.Type = &v
	return s
}

type PriceSaveRequestArisItemStayRules struct {
	MaxLos *int64 `json:"max_los,omitempty" xml:"max_los,omitempty"`
	MinLos *int64 `json:"min_los,omitempty" xml:"min_los,omitempty"`
}

func (s PriceSaveRequestArisItemStayRules) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemStayRules) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemStayRules) SetMaxLos(v int64) *PriceSaveRequestArisItemStayRules {
	s.MaxLos = &v
	return s
}

func (s *PriceSaveRequestArisItemStayRules) SetMinLos(v int64) *PriceSaveRequestArisItemStayRules {
	s.MinLos = &v
	return s
}

type PriceSaveRequestArisItemTimerange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s PriceSaveRequestArisItemTimerange) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveRequestArisItemTimerange) GoString() string {
	return s.String()
}

func (s *PriceSaveRequestArisItemTimerange) SetEnd(v string) *PriceSaveRequestArisItemTimerange {
	s.End = &v
	return s
}

func (s *PriceSaveRequestArisItemTimerange) SetStart(v string) *PriceSaveRequestArisItemTimerange {
	s.Start = &v
	return s
}

type PriceSaveResponse struct {
	Data  *PriceSaveResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *PriceSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s PriceSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveResponse) GoString() string {
	return s.String()
}

func (s *PriceSaveResponse) SetData(v *PriceSaveResponseData) *PriceSaveResponse {
	s.Data = v
	return s
}

func (s *PriceSaveResponse) SetExtra(v *PriceSaveResponseExtra) *PriceSaveResponse {
	s.Extra = v
	return s
}

type PriceSaveResponseData struct {
	SaveResult    []*PriceSaveResponseDataSaveResultItem `json:"save_result,omitempty" xml:"save_result,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PriceSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveResponseData) GoString() string {
	return s.String()
}

func (s *PriceSaveResponseData) SetSaveResult(v []*PriceSaveResponseDataSaveResultItem) *PriceSaveResponseData {
	s.SaveResult = v
	return s
}

func (s *PriceSaveResponseData) SetGwErrorCode(v int32) *PriceSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PriceSaveResponseData) SetGwDescription(v string) *PriceSaveResponseData {
	s.GwDescription = &v
	return s
}

type PriceSaveResponseDataSaveResultItem struct {
	Message    *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	RatePlanId *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	Code       *int64  `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s PriceSaveResponseDataSaveResultItem) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveResponseDataSaveResultItem) GoString() string {
	return s.String()
}

func (s *PriceSaveResponseDataSaveResultItem) SetMessage(v string) *PriceSaveResponseDataSaveResultItem {
	s.Message = &v
	return s
}

func (s *PriceSaveResponseDataSaveResultItem) SetRatePlanId(v string) *PriceSaveResponseDataSaveResultItem {
	s.RatePlanId = &v
	return s
}

func (s *PriceSaveResponseDataSaveResultItem) SetCode(v int64) *PriceSaveResponseDataSaveResultItem {
	s.Code = &v
	return s
}

type PriceSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s PriceSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PriceSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *PriceSaveResponseExtra) SetDescription(v string) *PriceSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *PriceSaveResponseExtra) SetErrorCode(v int32) *PriceSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PriceSaveResponseExtra) SetLogid(v string) *PriceSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *PriceSaveResponseExtra) SetNow(v int64) *PriceSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *PriceSaveResponseExtra) SetSubDescription(v string) *PriceSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PriceSaveResponseExtra) SetSubErrorCode(v int32) *PriceSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type PrivacySettingAddRequest struct {
	AccessToken     *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ContactWay      *int                                           `json:"contact_way,omitempty" xml:"contact_way,omitempty"`
	Email           *string                                        `json:"email,omitempty" xml:"email,omitempty"`
	Phone           *string                                        `json:"phone,omitempty" xml:"phone,omitempty"`
	LandLine        *string                                        `json:"land_line,omitempty" xml:"land_line,omitempty"`
	IsPrivacyConfig *bool                                          `json:"is_privacy_config,omitempty" xml:"is_privacy_config,omitempty" require:"true"`
	PrivacyItemList []*PrivacySettingAddRequestPrivacyItemListItem `json:"privacy_item_list,omitempty" xml:"privacy_item_list,omitempty" type:"Repeated"`
	Header          map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PrivacySettingAddRequest) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingAddRequest) GoString() string {
	return s.String()
}

func (s *PrivacySettingAddRequest) SetAccessToken(v string) *PrivacySettingAddRequest {
	s.AccessToken = &v
	return s
}

func (s *PrivacySettingAddRequest) SetContactWay(v int) *PrivacySettingAddRequest {
	s.ContactWay = &v
	return s
}

func (s *PrivacySettingAddRequest) SetEmail(v string) *PrivacySettingAddRequest {
	s.Email = &v
	return s
}

func (s *PrivacySettingAddRequest) SetPhone(v string) *PrivacySettingAddRequest {
	s.Phone = &v
	return s
}

func (s *PrivacySettingAddRequest) SetLandLine(v string) *PrivacySettingAddRequest {
	s.LandLine = &v
	return s
}

func (s *PrivacySettingAddRequest) SetIsPrivacyConfig(v bool) *PrivacySettingAddRequest {
	s.IsPrivacyConfig = &v
	return s
}

func (s *PrivacySettingAddRequest) SetPrivacyItemList(v []*PrivacySettingAddRequestPrivacyItemListItem) *PrivacySettingAddRequest {
	s.PrivacyItemList = v
	return s
}

func (s *PrivacySettingAddRequest) SetHeader(v map[string]*string) *PrivacySettingAddRequest {
	s.Header = v
	return s
}

type PrivacySettingAddRequestPrivacyItemListItem struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ApplyDesc *string `json:"apply_desc,omitempty" xml:"apply_desc,omitempty"`
}

func (s PrivacySettingAddRequestPrivacyItemListItem) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingAddRequestPrivacyItemListItem) GoString() string {
	return s.String()
}

func (s *PrivacySettingAddRequestPrivacyItemListItem) SetId(v int64) *PrivacySettingAddRequestPrivacyItemListItem {
	s.Id = &v
	return s
}

func (s *PrivacySettingAddRequestPrivacyItemListItem) SetApplyDesc(v string) *PrivacySettingAddRequestPrivacyItemListItem {
	s.ApplyDesc = &v
	return s
}

type PrivacySettingAddResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s PrivacySettingAddResponse) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingAddResponse) GoString() string {
	return s.String()
}

func (s *PrivacySettingAddResponse) SetErrMsg(v string) *PrivacySettingAddResponse {
	s.ErrMsg = &v
	return s
}

func (s *PrivacySettingAddResponse) SetLogId(v string) *PrivacySettingAddResponse {
	s.LogId = &v
	return s
}

func (s *PrivacySettingAddResponse) SetErrNo(v int32) *PrivacySettingAddResponse {
	s.ErrNo = &v
	return s
}

type PrivacySettingQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PrivacySettingQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingQueryRequest) GoString() string {
	return s.String()
}

func (s *PrivacySettingQueryRequest) SetHeader(v map[string]*string) *PrivacySettingQueryRequest {
	s.Header = v
	return s
}

func (s *PrivacySettingQueryRequest) SetAccessToken(v string) *PrivacySettingQueryRequest {
	s.AccessToken = &v
	return s
}

type PrivacySettingQueryResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *PrivacySettingQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s PrivacySettingQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingQueryResponse) GoString() string {
	return s.String()
}

func (s *PrivacySettingQueryResponse) SetLogId(v string) *PrivacySettingQueryResponse {
	s.LogId = &v
	return s
}

func (s *PrivacySettingQueryResponse) SetData(v *PrivacySettingQueryResponseData) *PrivacySettingQueryResponse {
	s.Data = v
	return s
}

func (s *PrivacySettingQueryResponse) SetErrNo(v int32) *PrivacySettingQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *PrivacySettingQueryResponse) SetErrMsg(v string) *PrivacySettingQueryResponse {
	s.ErrMsg = &v
	return s
}

type PrivacySettingQueryResponseData struct {
	IsConfigure *bool `json:"is_configure,omitempty" xml:"is_configure,omitempty"`
}

func (s PrivacySettingQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s PrivacySettingQueryResponseData) GoString() string {
	return s.String()
}

func (s *PrivacySettingQueryResponseData) SetIsConfigure(v bool) *PrivacySettingQueryResponseData {
	s.IsConfigure = &v
	return s
}

type ProductCommissionGetRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ProductCommissionGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetRequest) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetRequest) SetAccessToken(v string) *ProductCommissionGetRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductCommissionGetRequest) SetOrderId(v string) *ProductCommissionGetRequest {
	s.OrderId = &v
	return s
}

func (s *ProductCommissionGetRequest) SetHeader(v map[string]*string) *ProductCommissionGetRequest {
	s.Header = v
	return s
}

type ProductCommissionGetResponse struct {
	Data  *ProductCommissionGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ProductCommissionGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s ProductCommissionGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetResponse) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetResponse) SetData(v *ProductCommissionGetResponseData) *ProductCommissionGetResponse {
	s.Data = v
	return s
}

func (s *ProductCommissionGetResponse) SetExtra(v *ProductCommissionGetResponseExtra) *ProductCommissionGetResponse {
	s.Extra = v
	return s
}

type ProductCommissionGetResponseData struct {
	CommissionRecords []*ProductCommissionGetResponseDataCommissionRecordsItem `json:"commission_records,omitempty" xml:"commission_records,omitempty" type:"Repeated"`
	GwErrorCode       *int32                                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription     *string                                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductCommissionGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetResponseData) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetResponseData) SetCommissionRecords(v []*ProductCommissionGetResponseDataCommissionRecordsItem) *ProductCommissionGetResponseData {
	s.CommissionRecords = v
	return s
}

func (s *ProductCommissionGetResponseData) SetGwErrorCode(v int32) *ProductCommissionGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ProductCommissionGetResponseData) SetGwDescription(v string) *ProductCommissionGetResponseData {
	s.GwDescription = &v
	return s
}

type ProductCommissionGetResponseDataCommissionRecordsItem struct {
	IsEffective     *bool                                                                      `json:"is_effective,omitempty" xml:"is_effective,omitempty" require:"true"`
	MerchantAckTime *int64                                                                     `json:"merchant_ack_time,omitempty" xml:"merchant_ack_time,omitempty" require:"true"`
	ProductDetails  []*ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem `json:"product_details,omitempty" xml:"product_details,omitempty" type:"Repeated"`
	RecordId        *string                                                                    `json:"record_id,omitempty" xml:"record_id,omitempty" require:"true"`
	Status          *int                                                                       `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	UpdateTime      *int64                                                                     `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
	CreateTime      *int64                                                                     `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
}

func (s ProductCommissionGetResponseDataCommissionRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetResponseDataCommissionRecordsItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetIsEffective(v bool) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.IsEffective = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetMerchantAckTime(v int64) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.MerchantAckTime = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetProductDetails(v []*ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.ProductDetails = v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetRecordId(v string) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.RecordId = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetStatus(v int) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.Status = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetUpdateTime(v int64) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.UpdateTime = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItem) SetCreateTime(v int64) *ProductCommissionGetResponseDataCommissionRecordsItem {
	s.CreateTime = &v
	return s
}

type ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem struct {
	ProductStatus   *int    `json:"product_status,omitempty" xml:"product_status,omitempty" require:"true"`
	ActualPrice     *string `json:"actual_price,omitempty" xml:"actual_price,omitempty" require:"true"`
	CommissionRatio *string `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty" require:"true"`
	ProductId       *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) SetProductStatus(v int) *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem {
	s.ProductStatus = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) SetActualPrice(v string) *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem {
	s.ActualPrice = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) SetCommissionRatio(v string) *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem {
	s.CommissionRatio = &v
	return s
}

func (s *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem) SetProductId(v string) *ProductCommissionGetResponseDataCommissionRecordsItemProductDetailsItem {
	s.ProductId = &v
	return s
}

type ProductCommissionGetResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ProductCommissionGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionGetResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductCommissionGetResponseExtra) SetNow(v int64) *ProductCommissionGetResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductCommissionGetResponseExtra) SetSubDescription(v string) *ProductCommissionGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductCommissionGetResponseExtra) SetSubErrorCode(v int32) *ProductCommissionGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductCommissionGetResponseExtra) SetDescription(v string) *ProductCommissionGetResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductCommissionGetResponseExtra) SetErrorCode(v int32) *ProductCommissionGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductCommissionGetResponseExtra) SetLogid(v string) *ProductCommissionGetResponseExtra {
	s.Logid = &v
	return s
}

type ProductCommissionMgetRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	ProductIds  []*string          `json:"product_ids,omitempty" xml:"product_ids,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ProductCommissionMgetRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetRequest) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetRequest) SetAccessToken(v string) *ProductCommissionMgetRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductCommissionMgetRequest) SetOrderId(v string) *ProductCommissionMgetRequest {
	s.OrderId = &v
	return s
}

func (s *ProductCommissionMgetRequest) SetProductIds(v []*string) *ProductCommissionMgetRequest {
	s.ProductIds = v
	return s
}

func (s *ProductCommissionMgetRequest) SetHeader(v map[string]*string) *ProductCommissionMgetRequest {
	s.Header = v
	return s
}

type ProductCommissionMgetResponse struct {
	Data  *ProductCommissionMgetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ProductCommissionMgetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s ProductCommissionMgetResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponse) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponse) SetData(v *ProductCommissionMgetResponseData) *ProductCommissionMgetResponse {
	s.Data = v
	return s
}

func (s *ProductCommissionMgetResponse) SetExtra(v *ProductCommissionMgetResponseExtra) *ProductCommissionMgetResponse {
	s.Extra = v
	return s
}

type ProductCommissionMgetResponseData struct {
	ProductCommissions map[string]*ProductCommissionMgetResponseDataProductCommissionsValue `json:"product_commissions,omitempty" xml:"product_commissions,omitempty"`
	GwErrorCode        *int32                                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription      *string                                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductCommissionMgetResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponseData) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponseData) SetProductCommissions(v map[string]*ProductCommissionMgetResponseDataProductCommissionsValue) *ProductCommissionMgetResponseData {
	s.ProductCommissions = v
	return s
}

func (s *ProductCommissionMgetResponseData) SetGwErrorCode(v int32) *ProductCommissionMgetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ProductCommissionMgetResponseData) SetGwDescription(v string) *ProductCommissionMgetResponseData {
	s.GwDescription = &v
	return s
}

type ProductCommissionMgetResponseDataProductCommissionsValue struct {
	Product                    *ProductCommissionMgetResponseDataProductCommissionsValueProduct                    `json:"product,omitempty" xml:"product,omitempty"`
	LatestCommissionRecordItem *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem `json:"latest_commission_record_item,omitempty" xml:"latest_commission_record_item,omitempty"`
}

func (s ProductCommissionMgetResponseDataProductCommissionsValue) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponseDataProductCommissionsValue) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValue) SetProduct(v *ProductCommissionMgetResponseDataProductCommissionsValueProduct) *ProductCommissionMgetResponseDataProductCommissionsValue {
	s.Product = v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValue) SetLatestCommissionRecordItem(v *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) *ProductCommissionMgetResponseDataProductCommissionsValue {
	s.LatestCommissionRecordItem = v
	return s
}

type ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem struct {
	AuditCommissionRatio  *string `json:"audit_commission_ratio,omitempty" xml:"audit_commission_ratio,omitempty" require:"true"`
	CommissionRatioBefore *string `json:"commission_ratio_before,omitempty" xml:"commission_ratio_before,omitempty"`
	CommissionModeBefore  *int    `json:"commission_mode_before,omitempty" xml:"commission_mode_before,omitempty"`
	CommissionAuditStatus *int    `json:"commission_audit_status,omitempty" xml:"commission_audit_status,omitempty" require:"true"`
	RejectReason          *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	AuditCommissionMode   *int    `json:"audit_commission_mode,omitempty" xml:"audit_commission_mode,omitempty"`
	ItemId                *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	ProductId             *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	MerchantAckTime       *int64  `json:"merchant_ack_time,omitempty" xml:"merchant_ack_time,omitempty"`
}

func (s ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetAuditCommissionRatio(v string) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.AuditCommissionRatio = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetCommissionRatioBefore(v string) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.CommissionRatioBefore = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetCommissionModeBefore(v int) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.CommissionModeBefore = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetCommissionAuditStatus(v int) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.CommissionAuditStatus = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetRejectReason(v string) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.RejectReason = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetAuditCommissionMode(v int) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.AuditCommissionMode = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetItemId(v string) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.ItemId = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetProductId(v string) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.ProductId = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem) SetMerchantAckTime(v int64) *ProductCommissionMgetResponseDataProductCommissionsValueLatestCommissionRecordItem {
	s.MerchantAckTime = &v
	return s
}

type ProductCommissionMgetResponseDataProductCommissionsValueProduct struct {
	ProductId                *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	ProductName              *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	ProductStatus            *int    `json:"product_status,omitempty" xml:"product_status,omitempty" require:"true"`
	ActualPrice              *string `json:"actual_price,omitempty" xml:"actual_price,omitempty" require:"true"`
	EffectiveCommissionMode  *int    `json:"effective_commission_mode,omitempty" xml:"effective_commission_mode,omitempty"`
	EffectiveCommissionRatio *string `json:"effective_commission_ratio,omitempty" xml:"effective_commission_ratio,omitempty"`
	ProductCommissionStatus  *int    `json:"product_commission_status,omitempty" xml:"product_commission_status,omitempty" require:"true"`
}

func (s ProductCommissionMgetResponseDataProductCommissionsValueProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponseDataProductCommissionsValueProduct) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetProductId(v string) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.ProductId = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetProductName(v string) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.ProductName = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetProductStatus(v int) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.ProductStatus = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetActualPrice(v string) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.ActualPrice = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetEffectiveCommissionMode(v int) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.EffectiveCommissionMode = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetEffectiveCommissionRatio(v string) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.EffectiveCommissionRatio = &v
	return s
}

func (s *ProductCommissionMgetResponseDataProductCommissionsValueProduct) SetProductCommissionStatus(v int) *ProductCommissionMgetResponseDataProductCommissionsValueProduct {
	s.ProductCommissionStatus = &v
	return s
}

type ProductCommissionMgetResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ProductCommissionMgetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionMgetResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductCommissionMgetResponseExtra) SetSubErrorCode(v int32) *ProductCommissionMgetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductCommissionMgetResponseExtra) SetDescription(v string) *ProductCommissionMgetResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductCommissionMgetResponseExtra) SetErrorCode(v int32) *ProductCommissionMgetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductCommissionMgetResponseExtra) SetLogid(v string) *ProductCommissionMgetResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductCommissionMgetResponseExtra) SetNow(v int64) *ProductCommissionMgetResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductCommissionMgetResponseExtra) SetSubDescription(v string) *ProductCommissionMgetResponseExtra {
	s.SubDescription = &v
	return s
}

type ProductCommissionQueryRequest struct {
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Page        *int32             `json:"page,omitempty" xml:"page,omitempty"`
	Size        *int32             `json:"size,omitempty" xml:"size,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ProductCommissionQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryRequest) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryRequest) SetOrderId(v string) *ProductCommissionQueryRequest {
	s.OrderId = &v
	return s
}

func (s *ProductCommissionQueryRequest) SetPage(v int32) *ProductCommissionQueryRequest {
	s.Page = &v
	return s
}

func (s *ProductCommissionQueryRequest) SetSize(v int32) *ProductCommissionQueryRequest {
	s.Size = &v
	return s
}

func (s *ProductCommissionQueryRequest) SetHeader(v map[string]*string) *ProductCommissionQueryRequest {
	s.Header = v
	return s
}

func (s *ProductCommissionQueryRequest) SetAccessToken(v string) *ProductCommissionQueryRequest {
	s.AccessToken = &v
	return s
}

type ProductCommissionQueryResponse struct {
	Extra *ProductCommissionQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *ProductCommissionQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ProductCommissionQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponse) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponse) SetExtra(v *ProductCommissionQueryResponseExtra) *ProductCommissionQueryResponse {
	s.Extra = v
	return s
}

func (s *ProductCommissionQueryResponse) SetData(v *ProductCommissionQueryResponseData) *ProductCommissionQueryResponse {
	s.Data = v
	return s
}

type ProductCommissionQueryResponseData struct {
	ProductCommissions []*ProductCommissionQueryResponseDataProductCommissionsItem `json:"product_commissions,omitempty" xml:"product_commissions,omitempty" type:"Repeated"`
	GwErrorCode        *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription      *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Total              *int32                                                      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ProductCommissionQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponseData) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponseData) SetProductCommissions(v []*ProductCommissionQueryResponseDataProductCommissionsItem) *ProductCommissionQueryResponseData {
	s.ProductCommissions = v
	return s
}

func (s *ProductCommissionQueryResponseData) SetGwErrorCode(v int32) *ProductCommissionQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ProductCommissionQueryResponseData) SetGwDescription(v string) *ProductCommissionQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *ProductCommissionQueryResponseData) SetTotal(v int32) *ProductCommissionQueryResponseData {
	s.Total = &v
	return s
}

type ProductCommissionQueryResponseDataProductCommissionsItem struct {
	LatestCommissionRecordItem *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem `json:"latest_commission_record_item,omitempty" xml:"latest_commission_record_item,omitempty"`
	Product                    *ProductCommissionQueryResponseDataProductCommissionsItemProduct                    `json:"product,omitempty" xml:"product,omitempty"`
}

func (s ProductCommissionQueryResponseDataProductCommissionsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponseDataProductCommissionsItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItem) SetLatestCommissionRecordItem(v *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) *ProductCommissionQueryResponseDataProductCommissionsItem {
	s.LatestCommissionRecordItem = v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItem) SetProduct(v *ProductCommissionQueryResponseDataProductCommissionsItemProduct) *ProductCommissionQueryResponseDataProductCommissionsItem {
	s.Product = v
	return s
}

type ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem struct {
	CommissionAuditStatus *int    `json:"commission_audit_status,omitempty" xml:"commission_audit_status,omitempty" require:"true"`
	AuditCommissionRatio  *string `json:"audit_commission_ratio,omitempty" xml:"audit_commission_ratio,omitempty" require:"true"`
	ItemId                *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	MerchantAckTime       *int64  `json:"merchant_ack_time,omitempty" xml:"merchant_ack_time,omitempty"`
	ProductId             *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	CommissionRatioBefore *string `json:"commission_ratio_before,omitempty" xml:"commission_ratio_before,omitempty"`
	CommissionModeBefore  *int    `json:"commission_mode_before,omitempty" xml:"commission_mode_before,omitempty"`
	RejectReason          *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	AuditCommissionMode   *int    `json:"audit_commission_mode,omitempty" xml:"audit_commission_mode,omitempty"`
}

func (s ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetCommissionAuditStatus(v int) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.CommissionAuditStatus = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetAuditCommissionRatio(v string) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.AuditCommissionRatio = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetItemId(v string) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.ItemId = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetMerchantAckTime(v int64) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.MerchantAckTime = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetProductId(v string) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.ProductId = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetCommissionRatioBefore(v string) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.CommissionRatioBefore = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetCommissionModeBefore(v int) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.CommissionModeBefore = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetRejectReason(v string) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.RejectReason = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem) SetAuditCommissionMode(v int) *ProductCommissionQueryResponseDataProductCommissionsItemLatestCommissionRecordItem {
	s.AuditCommissionMode = &v
	return s
}

type ProductCommissionQueryResponseDataProductCommissionsItemProduct struct {
	EffectiveCommissionMode  *int    `json:"effective_commission_mode,omitempty" xml:"effective_commission_mode,omitempty"`
	EffectiveCommissionRatio *string `json:"effective_commission_ratio,omitempty" xml:"effective_commission_ratio,omitempty"`
	ProductCommissionStatus  *int    `json:"product_commission_status,omitempty" xml:"product_commission_status,omitempty" require:"true"`
	ProductId                *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	ProductName              *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	ProductStatus            *int    `json:"product_status,omitempty" xml:"product_status,omitempty" require:"true"`
	ActualPrice              *string `json:"actual_price,omitempty" xml:"actual_price,omitempty" require:"true"`
}

func (s ProductCommissionQueryResponseDataProductCommissionsItemProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponseDataProductCommissionsItemProduct) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetEffectiveCommissionMode(v int) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.EffectiveCommissionMode = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetEffectiveCommissionRatio(v string) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.EffectiveCommissionRatio = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetProductCommissionStatus(v int) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.ProductCommissionStatus = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetProductId(v string) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.ProductId = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetProductName(v string) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.ProductName = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetProductStatus(v int) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.ProductStatus = &v
	return s
}

func (s *ProductCommissionQueryResponseDataProductCommissionsItemProduct) SetActualPrice(v string) *ProductCommissionQueryResponseDataProductCommissionsItemProduct {
	s.ActualPrice = &v
	return s
}

type ProductCommissionQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ProductCommissionQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductCommissionQueryResponseExtra) SetSubDescription(v string) *ProductCommissionQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductCommissionQueryResponseExtra) SetSubErrorCode(v int32) *ProductCommissionQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductCommissionQueryResponseExtra) SetDescription(v string) *ProductCommissionQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductCommissionQueryResponseExtra) SetErrorCode(v int32) *ProductCommissionQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductCommissionQueryResponseExtra) SetLogid(v string) *ProductCommissionQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductCommissionQueryResponseExtra) SetNow(v int64) *ProductCommissionQueryResponseExtra {
	s.Now = &v
	return s
}

type ProductCommissionSaveRequest struct {
	ProductItems []*ProductCommissionSaveRequestProductItemsItem `json:"product_items,omitempty" xml:"product_items,omitempty" type:"Repeated"`
	Header       map[string]*string                              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId      *string                                         `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s ProductCommissionSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveRequest) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveRequest) SetProductItems(v []*ProductCommissionSaveRequestProductItemsItem) *ProductCommissionSaveRequest {
	s.ProductItems = v
	return s
}

func (s *ProductCommissionSaveRequest) SetHeader(v map[string]*string) *ProductCommissionSaveRequest {
	s.Header = v
	return s
}

func (s *ProductCommissionSaveRequest) SetAccessToken(v string) *ProductCommissionSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductCommissionSaveRequest) SetOrderId(v string) *ProductCommissionSaveRequest {
	s.OrderId = &v
	return s
}

type ProductCommissionSaveRequestProductItemsItem struct {
	AllianceCommissionRatio *string `json:"alliance_commission_ratio,omitempty" xml:"alliance_commission_ratio,omitempty"`
	CommissionMode          *int    `json:"commission_mode,omitempty" xml:"commission_mode,omitempty"`
	CommissionRatio         *string `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
	ProductId               *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s ProductCommissionSaveRequestProductItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveRequestProductItemsItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveRequestProductItemsItem) SetAllianceCommissionRatio(v string) *ProductCommissionSaveRequestProductItemsItem {
	s.AllianceCommissionRatio = &v
	return s
}

func (s *ProductCommissionSaveRequestProductItemsItem) SetCommissionMode(v int) *ProductCommissionSaveRequestProductItemsItem {
	s.CommissionMode = &v
	return s
}

func (s *ProductCommissionSaveRequestProductItemsItem) SetCommissionRatio(v string) *ProductCommissionSaveRequestProductItemsItem {
	s.CommissionRatio = &v
	return s
}

func (s *ProductCommissionSaveRequestProductItemsItem) SetProductId(v string) *ProductCommissionSaveRequestProductItemsItem {
	s.ProductId = &v
	return s
}

type ProductCommissionSaveResponse struct {
	Extra *ProductCommissionSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *ProductCommissionSaveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ProductCommissionSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveResponse) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveResponse) SetExtra(v *ProductCommissionSaveResponseExtra) *ProductCommissionSaveResponse {
	s.Extra = v
	return s
}

func (s *ProductCommissionSaveResponse) SetData(v *ProductCommissionSaveResponseData) *ProductCommissionSaveResponse {
	s.Data = v
	return s
}

type ProductCommissionSaveResponseData struct {
	GwDescription                *string                                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	InvalidCommissionProductList []*ProductCommissionSaveResponseDataInvalidCommissionProductListItem `json:"invalid_commission_product_list,omitempty" xml:"invalid_commission_product_list,omitempty" type:"Repeated"`
	RecordId                     *string                                                              `json:"record_id,omitempty" xml:"record_id,omitempty"`
	GwErrorCode                  *int32                                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ProductCommissionSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveResponseData) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveResponseData) SetGwDescription(v string) *ProductCommissionSaveResponseData {
	s.GwDescription = &v
	return s
}

func (s *ProductCommissionSaveResponseData) SetInvalidCommissionProductList(v []*ProductCommissionSaveResponseDataInvalidCommissionProductListItem) *ProductCommissionSaveResponseData {
	s.InvalidCommissionProductList = v
	return s
}

func (s *ProductCommissionSaveResponseData) SetRecordId(v string) *ProductCommissionSaveResponseData {
	s.RecordId = &v
	return s
}

func (s *ProductCommissionSaveResponseData) SetGwErrorCode(v int32) *ProductCommissionSaveResponseData {
	s.GwErrorCode = &v
	return s
}

type ProductCommissionSaveResponseDataInvalidCommissionProductListItem struct {
	AllianceCommissionRatio *string `json:"alliance_commission_ratio,omitempty" xml:"alliance_commission_ratio,omitempty"`
	CommissionMode          *int    `json:"commission_mode,omitempty" xml:"commission_mode,omitempty"`
	CommissionRatio         *string `json:"commission_ratio,omitempty" xml:"commission_ratio,omitempty"`
	ProductId               *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s ProductCommissionSaveResponseDataInvalidCommissionProductListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveResponseDataInvalidCommissionProductListItem) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveResponseDataInvalidCommissionProductListItem) SetAllianceCommissionRatio(v string) *ProductCommissionSaveResponseDataInvalidCommissionProductListItem {
	s.AllianceCommissionRatio = &v
	return s
}

func (s *ProductCommissionSaveResponseDataInvalidCommissionProductListItem) SetCommissionMode(v int) *ProductCommissionSaveResponseDataInvalidCommissionProductListItem {
	s.CommissionMode = &v
	return s
}

func (s *ProductCommissionSaveResponseDataInvalidCommissionProductListItem) SetCommissionRatio(v string) *ProductCommissionSaveResponseDataInvalidCommissionProductListItem {
	s.CommissionRatio = &v
	return s
}

func (s *ProductCommissionSaveResponseDataInvalidCommissionProductListItem) SetProductId(v string) *ProductCommissionSaveResponseDataInvalidCommissionProductListItem {
	s.ProductId = &v
	return s
}

type ProductCommissionSaveResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ProductCommissionSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductCommissionSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductCommissionSaveResponseExtra) SetSubDescription(v string) *ProductCommissionSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductCommissionSaveResponseExtra) SetSubErrorCode(v int32) *ProductCommissionSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductCommissionSaveResponseExtra) SetDescription(v string) *ProductCommissionSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductCommissionSaveResponseExtra) SetErrorCode(v int32) *ProductCommissionSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductCommissionSaveResponseExtra) SetLogid(v string) *ProductCommissionSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductCommissionSaveResponseExtra) SetNow(v int64) *ProductCommissionSaveResponseExtra {
	s.Now = &v
	return s
}

type ProductDetailRequest struct {
	AccountId   *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OutId       *string            `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId   *int64             `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ProductDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailRequest) GoString() string {
	return s.String()
}

func (s *ProductDetailRequest) SetAccountId(v int64) *ProductDetailRequest {
	s.AccountId = &v
	return s
}

func (s *ProductDetailRequest) SetOutId(v string) *ProductDetailRequest {
	s.OutId = &v
	return s
}

func (s *ProductDetailRequest) SetProductId(v int64) *ProductDetailRequest {
	s.ProductId = &v
	return s
}

func (s *ProductDetailRequest) SetHeader(v map[string]*string) *ProductDetailRequest {
	s.Header = v
	return s
}

func (s *ProductDetailRequest) SetAccessToken(v string) *ProductDetailRequest {
	s.AccessToken = &v
	return s
}

type ProductDetailResponse struct {
	Extra    *ProductDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *ProductDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ProductDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponse) GoString() string {
	return s.String()
}

func (s *ProductDetailResponse) SetExtra(v *ProductDetailResponseExtra) *ProductDetailResponse {
	s.Extra = v
	return s
}

func (s *ProductDetailResponse) SetBaseResp(v *ProductDetailResponseBaseResp) *ProductDetailResponse {
	s.BaseResp = v
	return s
}

func (s *ProductDetailResponse) SetData(v *ProductDetailResponseData) *ProductDetailResponse {
	s.Data = v
	return s
}

type ProductDetailResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseBaseResp) SetStatusCode(v int32) *ProductDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductDetailResponseBaseResp) SetStatusMessage(v string) *ProductDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductDetailResponseBaseResp) SetExtra(v map[string]*string) *ProductDetailResponseBaseResp {
	s.Extra = v
	return s
}

type ProductDetailResponseData struct {
	ErrorCode     *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductDraft  *ProductDetailResponseDataProductDraft  `json:"product_draft,omitempty" xml:"product_draft,omitempty"`
	ProductOnline *ProductDetailResponseDataProductOnline `json:"product_online,omitempty" xml:"product_online,omitempty"`
	Description   *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseData) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseData) SetErrorCode(v int32) *ProductDetailResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductDetailResponseData) SetProductDraft(v *ProductDetailResponseDataProductDraft) *ProductDetailResponseData {
	s.ProductDraft = v
	return s
}

func (s *ProductDetailResponseData) SetProductOnline(v *ProductDetailResponseDataProductOnline) *ProductDetailResponseData {
	s.ProductOnline = v
	return s
}

func (s *ProductDetailResponseData) SetDescription(v string) *ProductDetailResponseData {
	s.Description = &v
	return s
}

type ProductDetailResponseDataProductDraft struct {
	Product     *ProductDetailResponseDataProductDraftProduct `json:"product,omitempty" xml:"product,omitempty"`
	DraftStatus *int                                          `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
}

func (s ProductDetailResponseDataProductDraft) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraft) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraft) SetProduct(v *ProductDetailResponseDataProductDraftProduct) *ProductDetailResponseDataProductDraft {
	s.Product = v
	return s
}

func (s *ProductDetailResponseDataProductDraft) SetDraftStatus(v int) *ProductDetailResponseDataProductDraft {
	s.DraftStatus = &v
	return s
}

type ProductDetailResponseDataProductDraftProduct struct {
	FulfillmentMethod []*int                                                              `json:"fulfillment_method,omitempty" xml:"fulfillment_method,omitempty" type:"Repeated"`
	SkuList           []*ProductDetailResponseDataProductDraftProductSkuListItem          `json:"sku_list,omitempty" xml:"sku_list,omitempty" type:"Repeated"`
	CategoryId        *int64                                                              `json:"category_id,omitempty" xml:"category_id,omitempty"`
	SpuId             *int64                                                              `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	AccountId         *int64                                                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PackFee           *ProductDetailResponseDataProductDraftProductPackFee                `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	SaleAttrGroups    []*ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem   `json:"sale_attr_groups,omitempty" xml:"sale_attr_groups,omitempty" type:"Repeated"`
	SettleInfo        *ProductDetailResponseDataProductDraftProductSettleInfo             `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	OutUrl            *string                                                             `json:"out_url,omitempty" xml:"out_url,omitempty"`
	PoiList           []*ProductDetailResponseDataProductDraftProductPoiListItem          `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	Images            *ProductDetailResponseDataProductDraftProductImages                 `json:"images,omitempty" xml:"images,omitempty"`
	BizLine           *int                                                                `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	OutSpuId          *string                                                             `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	OutId             *string                                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Name              *string                                                             `json:"name,omitempty" xml:"name,omitempty"`
	ProductId         *int64                                                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	AffiliatedGroups  []*ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem `json:"affiliated_groups,omitempty" xml:"affiliated_groups,omitempty" type:"Repeated"`
	TestInfo          *ProductDetailResponseDataProductDraftProductTestInfo               `json:"test_info,omitempty" xml:"test_info,omitempty"`
	SoldTime          *ProductDetailResponseDataProductDraftProductSoldTime               `json:"sold_time,omitempty" xml:"sold_time,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProduct) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProduct) SetFulfillmentMethod(v []*int) *ProductDetailResponseDataProductDraftProduct {
	s.FulfillmentMethod = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetSkuList(v []*ProductDetailResponseDataProductDraftProductSkuListItem) *ProductDetailResponseDataProductDraftProduct {
	s.SkuList = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetCategoryId(v int64) *ProductDetailResponseDataProductDraftProduct {
	s.CategoryId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetSpuId(v int64) *ProductDetailResponseDataProductDraftProduct {
	s.SpuId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetAccountId(v int64) *ProductDetailResponseDataProductDraftProduct {
	s.AccountId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetPackFee(v *ProductDetailResponseDataProductDraftProductPackFee) *ProductDetailResponseDataProductDraftProduct {
	s.PackFee = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetSaleAttrGroups(v []*ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) *ProductDetailResponseDataProductDraftProduct {
	s.SaleAttrGroups = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetSettleInfo(v *ProductDetailResponseDataProductDraftProductSettleInfo) *ProductDetailResponseDataProductDraftProduct {
	s.SettleInfo = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetOutUrl(v string) *ProductDetailResponseDataProductDraftProduct {
	s.OutUrl = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetPoiList(v []*ProductDetailResponseDataProductDraftProductPoiListItem) *ProductDetailResponseDataProductDraftProduct {
	s.PoiList = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetImages(v *ProductDetailResponseDataProductDraftProductImages) *ProductDetailResponseDataProductDraftProduct {
	s.Images = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetBizLine(v int) *ProductDetailResponseDataProductDraftProduct {
	s.BizLine = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetOutSpuId(v string) *ProductDetailResponseDataProductDraftProduct {
	s.OutSpuId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetOutId(v string) *ProductDetailResponseDataProductDraftProduct {
	s.OutId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetName(v string) *ProductDetailResponseDataProductDraftProduct {
	s.Name = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetProductId(v int64) *ProductDetailResponseDataProductDraftProduct {
	s.ProductId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetAffiliatedGroups(v []*ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) *ProductDetailResponseDataProductDraftProduct {
	s.AffiliatedGroups = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetTestInfo(v *ProductDetailResponseDataProductDraftProductTestInfo) *ProductDetailResponseDataProductDraftProduct {
	s.TestInfo = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProduct) SetSoldTime(v *ProductDetailResponseDataProductDraftProductSoldTime) *ProductDetailResponseDataProductDraftProduct {
	s.SoldTime = v
	return s
}

type ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem struct {
	GroupName        *string   `json:"group_name,omitempty" xml:"group_name,omitempty"`
	OutAffiliatedIds []*string `json:"out_affiliated_ids,omitempty" xml:"out_affiliated_ids,omitempty" type:"Repeated"`
	GroupCode        *string   `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) SetGroupName(v string) *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) SetOutAffiliatedIds(v []*string) *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem {
	s.OutAffiliatedIds = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem) SetGroupCode(v string) *ProductDetailResponseDataProductDraftProductAffiliatedGroupsItem {
	s.GroupCode = &v
	return s
}

type ProductDetailResponseDataProductDraftProductImages struct {
	HeadImage *ProductDetailResponseDataProductDraftProductImagesHeadImage `json:"head_image,omitempty" xml:"head_image,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductImages) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductImages) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductImages) SetHeadImage(v *ProductDetailResponseDataProductDraftProductImagesHeadImage) *ProductDetailResponseDataProductDraftProductImages {
	s.HeadImage = v
	return s
}

type ProductDetailResponseDataProductDraftProductImagesHeadImage struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductImagesHeadImage) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductImagesHeadImage) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductImagesHeadImage) SetUrl(v string) *ProductDetailResponseDataProductDraftProductImagesHeadImage {
	s.Url = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductImagesHeadImage) SetUri(v string) *ProductDetailResponseDataProductDraftProductImagesHeadImage {
	s.Uri = &v
	return s
}

type ProductDetailResponseDataProductDraftProductPackFee struct {
	Step        *int64 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int64 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductPackFee) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductPackFee) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductPackFee) SetStep(v int64) *ProductDetailResponseDataProductDraftProductPackFee {
	s.Step = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductPackFee) SetPackFee(v int64) *ProductDetailResponseDataProductDraftProductPackFee {
	s.PackFee = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductPackFee) SetPackFeeUnit(v int) *ProductDetailResponseDataProductDraftProductPackFee {
	s.PackFeeUnit = &v
	return s
}

type ProductDetailResponseDataProductDraftProductPoiListItem struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductPoiListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductPoiListItem) SetPoiId(v int64) *ProductDetailResponseDataProductDraftProductPoiListItem {
	s.PoiId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductPoiListItem) SetExtId(v string) *ProductDetailResponseDataProductDraftProductPoiListItem {
	s.ExtId = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem struct {
	GroupName *string                                                                       `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	SpecType  *int                                                                          `json:"spec_type,omitempty" xml:"spec_type,omitempty"`
	GroupCode *string                                                                       `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) SetGroupName(v string) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) SetItemList(v []*ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem {
	s.ItemList = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) SetSpecType(v int) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem {
	s.SpecType = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem) SetGroupCode(v string) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItem {
	s.GroupCode = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem struct {
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	Checked  *bool   `json:"checked,omitempty" xml:"checked,omitempty"`
	ItemKey  *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) SetItemName(v string) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) SetChecked(v bool) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem {
	s.Checked = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem) SetItemKey(v string) *ProductDetailResponseDataProductDraftProductSaleAttrGroupsItemItemListItem {
	s.ItemKey = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSettleInfo) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSettleInfo) SetSettleType(v int) *ProductDetailResponseDataProductDraftProductSettleInfo {
	s.SettleType = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSkuListItem struct {
	OutSkuId     *string                                                                    `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price        *ProductDetailResponseDataProductDraftProductSkuListItemPrice              `json:"price,omitempty" xml:"price,omitempty"`
	SaleAttrList []*ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem `json:"sale_attr_list,omitempty" xml:"sale_attr_list,omitempty" type:"Repeated"`
	SkuId        *int64                                                                     `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Status       *int                                                                       `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSkuListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItem) SetOutSkuId(v string) *ProductDetailResponseDataProductDraftProductSkuListItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItem) SetPrice(v *ProductDetailResponseDataProductDraftProductSkuListItemPrice) *ProductDetailResponseDataProductDraftProductSkuListItem {
	s.Price = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItem) SetSaleAttrList(v []*ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem) *ProductDetailResponseDataProductDraftProductSkuListItem {
	s.SaleAttrList = v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItem) SetSkuId(v int64) *ProductDetailResponseDataProductDraftProductSkuListItem {
	s.SkuId = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItem) SetStatus(v int) *ProductDetailResponseDataProductDraftProductSkuListItem {
	s.Status = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSkuListItemPrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSkuListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSkuListItemPrice) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItemPrice) SetActualAmount(v int64) *ProductDetailResponseDataProductDraftProductSkuListItemPrice {
	s.ActualAmount = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItemPrice) SetOriginAmount(v int64) *ProductDetailResponseDataProductDraftProductSkuListItemPrice {
	s.OriginAmount = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem struct {
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem) SetItemKey(v string) *ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem {
	s.ItemKey = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem) SetGroupCode(v string) *ProductDetailResponseDataProductDraftProductSkuListItemSaleAttrListItem {
	s.GroupCode = &v
	return s
}

type ProductDetailResponseDataProductDraftProductSoldTime struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ProductDetailResponseDataProductDraftProductSoldTime) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductSoldTime) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductSoldTime) SetEndTime(v int64) *ProductDetailResponseDataProductDraftProductSoldTime {
	s.EndTime = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductSoldTime) SetStartTime(v int64) *ProductDetailResponseDataProductDraftProductSoldTime {
	s.StartTime = &v
	return s
}

type ProductDetailResponseDataProductDraftProductTestInfo struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	UidList  []*string `json:"uid_list,omitempty" xml:"uid_list,omitempty" type:"Repeated"`
}

func (s ProductDetailResponseDataProductDraftProductTestInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductDraftProductTestInfo) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductDraftProductTestInfo) SetTestFlag(v bool) *ProductDetailResponseDataProductDraftProductTestInfo {
	s.TestFlag = &v
	return s
}

func (s *ProductDetailResponseDataProductDraftProductTestInfo) SetUidList(v []*string) *ProductDetailResponseDataProductDraftProductTestInfo {
	s.UidList = v
	return s
}

type ProductDetailResponseDataProductOnline struct {
	Status  *int                                           `json:"status,omitempty" xml:"status,omitempty"`
	Product *ProductDetailResponseDataProductOnlineProduct `json:"product,omitempty" xml:"product,omitempty"`
}

func (s ProductDetailResponseDataProductOnline) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnline) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnline) SetStatus(v int) *ProductDetailResponseDataProductOnline {
	s.Status = &v
	return s
}

func (s *ProductDetailResponseDataProductOnline) SetProduct(v *ProductDetailResponseDataProductOnlineProduct) *ProductDetailResponseDataProductOnline {
	s.Product = v
	return s
}

type ProductDetailResponseDataProductOnlineProduct struct {
	OutUrl            *string                                                              `json:"out_url,omitempty" xml:"out_url,omitempty"`
	Name              *string                                                              `json:"name,omitempty" xml:"name,omitempty"`
	SaleAttrGroups    []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem   `json:"sale_attr_groups,omitempty" xml:"sale_attr_groups,omitempty" type:"Repeated"`
	FulfillmentMethod []*int                                                               `json:"fulfillment_method,omitempty" xml:"fulfillment_method,omitempty" type:"Repeated"`
	CategoryId        *int64                                                               `json:"category_id,omitempty" xml:"category_id,omitempty"`
	PackFee           *ProductDetailResponseDataProductOnlineProductPackFee                `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	AffiliatedGroups  []*ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem `json:"affiliated_groups,omitempty" xml:"affiliated_groups,omitempty" type:"Repeated"`
	SpuId             *int64                                                               `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Images            *ProductDetailResponseDataProductOnlineProductImages                 `json:"images,omitempty" xml:"images,omitempty"`
	AccountId         *int64                                                               `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutId             *string                                                              `json:"out_id,omitempty" xml:"out_id,omitempty"`
	SoldTime          *ProductDetailResponseDataProductOnlineProductSoldTime               `json:"sold_time,omitempty" xml:"sold_time,omitempty"`
	ProductId         *int64                                                               `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SettleInfo        *ProductDetailResponseDataProductOnlineProductSettleInfo             `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	BizLine           *int                                                                 `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	TestInfo          *ProductDetailResponseDataProductOnlineProductTestInfo               `json:"test_info,omitempty" xml:"test_info,omitempty"`
	PoiList           []*ProductDetailResponseDataProductOnlineProductPoiListItem          `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	SkuList           []*ProductDetailResponseDataProductOnlineProductSkuListItem          `json:"sku_list,omitempty" xml:"sku_list,omitempty" type:"Repeated"`
	OutSpuId          *string                                                              `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProduct) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetOutUrl(v string) *ProductDetailResponseDataProductOnlineProduct {
	s.OutUrl = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetName(v string) *ProductDetailResponseDataProductOnlineProduct {
	s.Name = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetSaleAttrGroups(v []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) *ProductDetailResponseDataProductOnlineProduct {
	s.SaleAttrGroups = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetFulfillmentMethod(v []*int) *ProductDetailResponseDataProductOnlineProduct {
	s.FulfillmentMethod = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetCategoryId(v int64) *ProductDetailResponseDataProductOnlineProduct {
	s.CategoryId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetPackFee(v *ProductDetailResponseDataProductOnlineProductPackFee) *ProductDetailResponseDataProductOnlineProduct {
	s.PackFee = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetAffiliatedGroups(v []*ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) *ProductDetailResponseDataProductOnlineProduct {
	s.AffiliatedGroups = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetSpuId(v int64) *ProductDetailResponseDataProductOnlineProduct {
	s.SpuId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetImages(v *ProductDetailResponseDataProductOnlineProductImages) *ProductDetailResponseDataProductOnlineProduct {
	s.Images = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetAccountId(v int64) *ProductDetailResponseDataProductOnlineProduct {
	s.AccountId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetOutId(v string) *ProductDetailResponseDataProductOnlineProduct {
	s.OutId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetSoldTime(v *ProductDetailResponseDataProductOnlineProductSoldTime) *ProductDetailResponseDataProductOnlineProduct {
	s.SoldTime = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetProductId(v int64) *ProductDetailResponseDataProductOnlineProduct {
	s.ProductId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetSettleInfo(v *ProductDetailResponseDataProductOnlineProductSettleInfo) *ProductDetailResponseDataProductOnlineProduct {
	s.SettleInfo = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetBizLine(v int) *ProductDetailResponseDataProductOnlineProduct {
	s.BizLine = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetTestInfo(v *ProductDetailResponseDataProductOnlineProductTestInfo) *ProductDetailResponseDataProductOnlineProduct {
	s.TestInfo = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetPoiList(v []*ProductDetailResponseDataProductOnlineProductPoiListItem) *ProductDetailResponseDataProductOnlineProduct {
	s.PoiList = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetSkuList(v []*ProductDetailResponseDataProductOnlineProductSkuListItem) *ProductDetailResponseDataProductOnlineProduct {
	s.SkuList = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProduct) SetOutSpuId(v string) *ProductDetailResponseDataProductOnlineProduct {
	s.OutSpuId = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem struct {
	GroupName        *string   `json:"group_name,omitempty" xml:"group_name,omitempty"`
	OutAffiliatedIds []*string `json:"out_affiliated_ids,omitempty" xml:"out_affiliated_ids,omitempty" type:"Repeated"`
	GroupCode        *string   `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) SetGroupName(v string) *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) SetOutAffiliatedIds(v []*string) *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem {
	s.OutAffiliatedIds = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem) SetGroupCode(v string) *ProductDetailResponseDataProductOnlineProductAffiliatedGroupsItem {
	s.GroupCode = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductImages struct {
	HeadImage *ProductDetailResponseDataProductOnlineProductImagesHeadImage `json:"head_image,omitempty" xml:"head_image,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductImages) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductImages) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductImages) SetHeadImage(v *ProductDetailResponseDataProductOnlineProductImagesHeadImage) *ProductDetailResponseDataProductOnlineProductImages {
	s.HeadImage = v
	return s
}

type ProductDetailResponseDataProductOnlineProductImagesHeadImage struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	Uri *string `json:"uri,omitempty" xml:"uri,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductImagesHeadImage) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductImagesHeadImage) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductImagesHeadImage) SetUrl(v string) *ProductDetailResponseDataProductOnlineProductImagesHeadImage {
	s.Url = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductImagesHeadImage) SetUri(v string) *ProductDetailResponseDataProductOnlineProductImagesHeadImage {
	s.Uri = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductPackFee struct {
	Step        *int64 `json:"step,omitempty" xml:"step,omitempty"`
	PackFee     *int64 `json:"pack_fee,omitempty" xml:"pack_fee,omitempty"`
	PackFeeUnit *int   `json:"pack_fee_unit,omitempty" xml:"pack_fee_unit,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductPackFee) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductPackFee) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductPackFee) SetStep(v int64) *ProductDetailResponseDataProductOnlineProductPackFee {
	s.Step = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductPackFee) SetPackFee(v int64) *ProductDetailResponseDataProductOnlineProductPackFee {
	s.PackFee = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductPackFee) SetPackFeeUnit(v int) *ProductDetailResponseDataProductOnlineProductPackFee {
	s.PackFeeUnit = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductPoiListItem struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductPoiListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductPoiListItem) SetExtId(v string) *ProductDetailResponseDataProductOnlineProductPoiListItem {
	s.ExtId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductPoiListItem) SetPoiId(v int64) *ProductDetailResponseDataProductOnlineProductPoiListItem {
	s.PoiId = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem struct {
	SpecType  *int                                                                           `json:"spec_type,omitempty" xml:"spec_type,omitempty"`
	GroupCode *string                                                                        `json:"group_code,omitempty" xml:"group_code,omitempty"`
	GroupName *string                                                                        `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) SetSpecType(v int) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem {
	s.SpecType = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) SetGroupCode(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem {
	s.GroupCode = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) SetGroupName(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem) SetItemList(v []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItem {
	s.ItemList = v
	return s
}

type ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem struct {
	DescList []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem `json:"desc_list,omitempty" xml:"desc_list,omitempty" type:"Repeated"`
	ItemKey  *string                                                                                    `json:"item_key,omitempty" xml:"item_key,omitempty"`
	ItemName *string                                                                                    `json:"item_name,omitempty" xml:"item_name,omitempty"`
	Checked  *bool                                                                                      `json:"checked,omitempty" xml:"checked,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) SetDescList(v []*ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem {
	s.DescList = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) SetItemKey(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem {
	s.ItemKey = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) SetItemName(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem) SetChecked(v bool) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItem {
	s.Checked = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem struct {
	Val *string `json:"val,omitempty" xml:"val,omitempty"`
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem) SetVal(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem {
	s.Val = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem) SetKey(v string) *ProductDetailResponseDataProductOnlineProductSaleAttrGroupsItemItemListItemDescListItem {
	s.Key = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSettleInfo) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSettleInfo) SetSettleType(v int) *ProductDetailResponseDataProductOnlineProductSettleInfo {
	s.SettleType = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSkuListItem struct {
	Status       *int                                                                        `json:"status,omitempty" xml:"status,omitempty"`
	OutSkuId     *string                                                                     `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Price        *ProductDetailResponseDataProductOnlineProductSkuListItemPrice              `json:"price,omitempty" xml:"price,omitempty"`
	SaleAttrList []*ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem `json:"sale_attr_list,omitempty" xml:"sale_attr_list,omitempty" type:"Repeated"`
	SkuId        *int64                                                                      `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItem) SetStatus(v int) *ProductDetailResponseDataProductOnlineProductSkuListItem {
	s.Status = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItem) SetOutSkuId(v string) *ProductDetailResponseDataProductOnlineProductSkuListItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItem) SetPrice(v *ProductDetailResponseDataProductOnlineProductSkuListItemPrice) *ProductDetailResponseDataProductOnlineProductSkuListItem {
	s.Price = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItem) SetSaleAttrList(v []*ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem) *ProductDetailResponseDataProductOnlineProductSkuListItem {
	s.SaleAttrList = v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItem) SetSkuId(v int64) *ProductDetailResponseDataProductOnlineProductSkuListItem {
	s.SkuId = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSkuListItemPrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItemPrice) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItemPrice) SetActualAmount(v int64) *ProductDetailResponseDataProductOnlineProductSkuListItemPrice {
	s.ActualAmount = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItemPrice) SetOriginAmount(v int64) *ProductDetailResponseDataProductOnlineProductSkuListItemPrice {
	s.OriginAmount = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem struct {
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem) SetItemKey(v string) *ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem {
	s.ItemKey = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem) SetGroupCode(v string) *ProductDetailResponseDataProductOnlineProductSkuListItemSaleAttrListItem {
	s.GroupCode = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductSoldTime struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s ProductDetailResponseDataProductOnlineProductSoldTime) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductSoldTime) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductSoldTime) SetEndTime(v int64) *ProductDetailResponseDataProductOnlineProductSoldTime {
	s.EndTime = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductSoldTime) SetStartTime(v int64) *ProductDetailResponseDataProductOnlineProductSoldTime {
	s.StartTime = &v
	return s
}

type ProductDetailResponseDataProductOnlineProductTestInfo struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	UidList  []*string `json:"uid_list,omitempty" xml:"uid_list,omitempty" type:"Repeated"`
}

func (s ProductDetailResponseDataProductOnlineProductTestInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseDataProductOnlineProductTestInfo) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseDataProductOnlineProductTestInfo) SetTestFlag(v bool) *ProductDetailResponseDataProductOnlineProductTestInfo {
	s.TestFlag = &v
	return s
}

func (s *ProductDetailResponseDataProductOnlineProductTestInfo) SetUidList(v []*string) *ProductDetailResponseDataProductOnlineProductTestInfo {
	s.UidList = v
	return s
}

type ProductDetailResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ProductDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductDetailResponseExtra) SetDescription(v string) *ProductDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductDetailResponseExtra) SetErrorCode(v int32) *ProductDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductDetailResponseExtra) SetLogid(v string) *ProductDetailResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductDetailResponseExtra) SetNow(v int64) *ProductDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductDetailResponseExtra) SetSubDescription(v string) *ProductDetailResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductDetailResponseExtra) SetSubErrorCode(v int32) *ProductDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ProductDraftListRequest struct {
	Cursor      *int32                       `json:"cursor,omitempty" xml:"cursor,omitempty"`
	DyPoiId     *int64                       `json:"dy_poi_id,omitempty" xml:"dy_poi_id,omitempty"`
	Status      []*int                       `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	Base        *ProductDraftListRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                      `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Count       *int32                       `json:"count,omitempty" xml:"count,omitempty"`
}

func (s ProductDraftListRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListRequest) GoString() string {
	return s.String()
}

func (s *ProductDraftListRequest) SetCursor(v int32) *ProductDraftListRequest {
	s.Cursor = &v
	return s
}

func (s *ProductDraftListRequest) SetDyPoiId(v int64) *ProductDraftListRequest {
	s.DyPoiId = &v
	return s
}

func (s *ProductDraftListRequest) SetStatus(v []*int) *ProductDraftListRequest {
	s.Status = v
	return s
}

func (s *ProductDraftListRequest) SetBase(v *ProductDraftListRequestBase) *ProductDraftListRequest {
	s.Base = v
	return s
}

func (s *ProductDraftListRequest) SetAccountId(v string) *ProductDraftListRequest {
	s.AccountId = &v
	return s
}

func (s *ProductDraftListRequest) SetHeader(v map[string]*string) *ProductDraftListRequest {
	s.Header = v
	return s
}

func (s *ProductDraftListRequest) SetAccessToken(v string) *ProductDraftListRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductDraftListRequest) SetCount(v int32) *ProductDraftListRequest {
	s.Count = &v
	return s
}

type ProductDraftListRequestBase struct {
	Caller     *string                                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                     `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductDraftListRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s ProductDraftListRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListRequestBase) GoString() string {
	return s.String()
}

func (s *ProductDraftListRequestBase) SetCaller(v string) *ProductDraftListRequestBase {
	s.Caller = &v
	return s
}

func (s *ProductDraftListRequestBase) SetClient(v string) *ProductDraftListRequestBase {
	s.Client = &v
	return s
}

func (s *ProductDraftListRequestBase) SetExtra(v map[string]*string) *ProductDraftListRequestBase {
	s.Extra = v
	return s
}

func (s *ProductDraftListRequestBase) SetLogID(v string) *ProductDraftListRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductDraftListRequestBase) SetTrafficEnv(v *ProductDraftListRequestBaseTrafficEnv) *ProductDraftListRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductDraftListRequestBase) SetAddr(v string) *ProductDraftListRequestBase {
	s.Addr = &v
	return s
}

type ProductDraftListRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s ProductDraftListRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductDraftListRequestBaseTrafficEnv) SetOpen(v bool) *ProductDraftListRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *ProductDraftListRequestBaseTrafficEnv) SetEnv(v string) *ProductDraftListRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type ProductDraftListResponse struct {
	Data     *ProductDraftListResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductDraftListResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *ProductDraftListResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s ProductDraftListResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponse) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponse) SetData(v *ProductDraftListResponseData) *ProductDraftListResponse {
	s.Data = v
	return s
}

func (s *ProductDraftListResponse) SetExtra(v *ProductDraftListResponseExtra) *ProductDraftListResponse {
	s.Extra = v
	return s
}

func (s *ProductDraftListResponse) SetBaseResp(v *ProductDraftListResponseBaseResp) *ProductDraftListResponse {
	s.BaseResp = v
	return s
}

type ProductDraftListResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductDraftListResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseBaseResp) SetStatusCode(v int32) *ProductDraftListResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductDraftListResponseBaseResp) SetStatusMessage(v string) *ProductDraftListResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductDraftListResponseBaseResp) SetExtra(v map[string]*string) *ProductDraftListResponseBaseResp {
	s.Extra = v
	return s
}

type ProductDraftListResponseData struct {
	Description *string                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore     *bool                                       `json:"has_more,omitempty" xml:"has_more,omitempty"`
	NextCursor  *int32                                      `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	Products    []*ProductDraftListResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Total       *int32                                      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ProductDraftListResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseData) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseData) SetDescription(v string) *ProductDraftListResponseData {
	s.Description = &v
	return s
}

func (s *ProductDraftListResponseData) SetErrorCode(v int32) *ProductDraftListResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductDraftListResponseData) SetHasMore(v bool) *ProductDraftListResponseData {
	s.HasMore = &v
	return s
}

func (s *ProductDraftListResponseData) SetNextCursor(v int32) *ProductDraftListResponseData {
	s.NextCursor = &v
	return s
}

func (s *ProductDraftListResponseData) SetProducts(v []*ProductDraftListResponseDataProductsItem) *ProductDraftListResponseData {
	s.Products = v
	return s
}

func (s *ProductDraftListResponseData) SetTotal(v int32) *ProductDraftListResponseData {
	s.Total = &v
	return s
}

type ProductDraftListResponseDataProductsItem struct {
	OutId              *string                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductType        *int                                                `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	ProductName        *string                                             `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	SkuName            *string                                             `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	AccountId          *string                                             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	CategoryId         *int64                                              `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	SkuId              *string                                             `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	IsSuperSkuSub      *bool                                               `json:"is_super_sku_sub,omitempty" xml:"is_super_sku_sub,omitempty"`
	OutSkuId           *string                                             `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	ActualAmount       *int64                                              `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	PoiCount           *int64                                              `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	DraftStatus        *int                                                `json:"draft_status,omitempty" xml:"draft_status,omitempty" require:"true"`
	AuditMsg           *string                                             `json:"audit_msg,omitempty" xml:"audit_msg,omitempty"`
	Pois               []*ProductDraftListResponseDataProductsItemPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	CategoryFullName   *string                                             `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	OriginAmount       *int64                                              `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	NotIndependentSale *bool                                               `json:"not_independent_sale,omitempty" xml:"not_independent_sale,omitempty"`
	ProductId          *string                                             `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Stock              *ProductDraftListResponseDataProductsItemStock      `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s ProductDraftListResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseDataProductsItem) SetOutId(v string) *ProductDraftListResponseDataProductsItem {
	s.OutId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetProductType(v int) *ProductDraftListResponseDataProductsItem {
	s.ProductType = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetProductName(v string) *ProductDraftListResponseDataProductsItem {
	s.ProductName = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetSkuName(v string) *ProductDraftListResponseDataProductsItem {
	s.SkuName = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetAccountId(v string) *ProductDraftListResponseDataProductsItem {
	s.AccountId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetCategoryId(v int64) *ProductDraftListResponseDataProductsItem {
	s.CategoryId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetSkuId(v string) *ProductDraftListResponseDataProductsItem {
	s.SkuId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetIsSuperSkuSub(v bool) *ProductDraftListResponseDataProductsItem {
	s.IsSuperSkuSub = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetOutSkuId(v string) *ProductDraftListResponseDataProductsItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetActualAmount(v int64) *ProductDraftListResponseDataProductsItem {
	s.ActualAmount = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetPoiCount(v int64) *ProductDraftListResponseDataProductsItem {
	s.PoiCount = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetDraftStatus(v int) *ProductDraftListResponseDataProductsItem {
	s.DraftStatus = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetAuditMsg(v string) *ProductDraftListResponseDataProductsItem {
	s.AuditMsg = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetPois(v []*ProductDraftListResponseDataProductsItemPoisItem) *ProductDraftListResponseDataProductsItem {
	s.Pois = v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetCategoryFullName(v string) *ProductDraftListResponseDataProductsItem {
	s.CategoryFullName = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetOriginAmount(v int64) *ProductDraftListResponseDataProductsItem {
	s.OriginAmount = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetNotIndependentSale(v bool) *ProductDraftListResponseDataProductsItem {
	s.NotIndependentSale = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetProductId(v string) *ProductDraftListResponseDataProductsItem {
	s.ProductId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItem) SetStock(v *ProductDraftListResponseDataProductsItemStock) *ProductDraftListResponseDataProductsItem {
	s.Stock = v
	return s
}

type ProductDraftListResponseDataProductsItemPoisItem struct {
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ProductDraftListResponseDataProductsItemPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseDataProductsItemPoisItem) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseDataProductsItemPoisItem) SetSupplierExtId(v string) *ProductDraftListResponseDataProductsItemPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemPoisItem) SetSupplierId(v int64) *ProductDraftListResponseDataProductsItemPoisItem {
	s.SupplierId = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemPoisItem) SetPoiId(v string) *ProductDraftListResponseDataProductsItemPoisItem {
	s.PoiId = &v
	return s
}

type ProductDraftListResponseDataProductsItemStock struct {
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
}

func (s ProductDraftListResponseDataProductsItemStock) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseDataProductsItemStock) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseDataProductsItemStock) SetFrozenQty(v int64) *ProductDraftListResponseDataProductsItemStock {
	s.FrozenQty = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemStock) SetLimitType(v int) *ProductDraftListResponseDataProductsItemStock {
	s.LimitType = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemStock) SetSoldCount(v int64) *ProductDraftListResponseDataProductsItemStock {
	s.SoldCount = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemStock) SetSoldQty(v int64) *ProductDraftListResponseDataProductsItemStock {
	s.SoldQty = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemStock) SetStockQty(v int64) *ProductDraftListResponseDataProductsItemStock {
	s.StockQty = &v
	return s
}

func (s *ProductDraftListResponseDataProductsItemStock) SetAvailQty(v int64) *ProductDraftListResponseDataProductsItemStock {
	s.AvailQty = &v
	return s
}

type ProductDraftListResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ProductDraftListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftListResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductDraftListResponseExtra) SetNow(v int64) *ProductDraftListResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductDraftListResponseExtra) SetSubDescription(v string) *ProductDraftListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductDraftListResponseExtra) SetSubErrorCode(v int32) *ProductDraftListResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductDraftListResponseExtra) SetDescription(v string) *ProductDraftListResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductDraftListResponseExtra) SetErrorCode(v int32) *ProductDraftListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductDraftListResponseExtra) SetLogid(v string) *ProductDraftListResponseExtra {
	s.Logid = &v
	return s
}

type ProductDraftQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Status      *int               `json:"status,omitempty" xml:"status,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Count       *int64             `json:"count,omitempty" xml:"count,omitempty"`
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ProductDraftQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryRequest) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryRequest) SetAccessToken(v string) *ProductDraftQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductDraftQueryRequest) SetStatus(v int) *ProductDraftQueryRequest {
	s.Status = &v
	return s
}

func (s *ProductDraftQueryRequest) SetAccountId(v string) *ProductDraftQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ProductDraftQueryRequest) SetCount(v int64) *ProductDraftQueryRequest {
	s.Count = &v
	return s
}

func (s *ProductDraftQueryRequest) SetCursor(v string) *ProductDraftQueryRequest {
	s.Cursor = &v
	return s
}

func (s *ProductDraftQueryRequest) SetHeader(v map[string]*string) *ProductDraftQueryRequest {
	s.Header = v
	return s
}

type ProductDraftQueryResponse struct {
	BaseResp *ProductDraftQueryResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductDraftQueryResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductDraftQueryResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ProductDraftQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponse) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponse) SetBaseResp(v *ProductDraftQueryResponseBaseResp) *ProductDraftQueryResponse {
	s.BaseResp = v
	return s
}

func (s *ProductDraftQueryResponse) SetData(v *ProductDraftQueryResponseData) *ProductDraftQueryResponse {
	s.Data = v
	return s
}

func (s *ProductDraftQueryResponse) SetExtra(v *ProductDraftQueryResponseExtra) *ProductDraftQueryResponse {
	s.Extra = v
	return s
}

type ProductDraftQueryResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ProductDraftQueryResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseBaseResp) SetStatusMessage(v string) *ProductDraftQueryResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductDraftQueryResponseBaseResp) SetExtra(v map[string]*string) *ProductDraftQueryResponseBaseResp {
	s.Extra = v
	return s
}

func (s *ProductDraftQueryResponseBaseResp) SetStatusCode(v int32) *ProductDraftQueryResponseBaseResp {
	s.StatusCode = &v
	return s
}

type ProductDraftQueryResponseData struct {
	Products    []*ProductDraftQueryResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore     *bool                                        `json:"has_more,omitempty" xml:"has_more,omitempty"`
	NextCursor  *string                                      `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
}

func (s ProductDraftQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseData) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseData) SetProducts(v []*ProductDraftQueryResponseDataProductsItem) *ProductDraftQueryResponseData {
	s.Products = v
	return s
}

func (s *ProductDraftQueryResponseData) SetDescription(v string) *ProductDraftQueryResponseData {
	s.Description = &v
	return s
}

func (s *ProductDraftQueryResponseData) SetErrorCode(v int32) *ProductDraftQueryResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductDraftQueryResponseData) SetHasMore(v bool) *ProductDraftQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *ProductDraftQueryResponseData) SetNextCursor(v string) *ProductDraftQueryResponseData {
	s.NextCursor = &v
	return s
}

type ProductDraftQueryResponseDataProductsItem struct {
	DraftStatus *int                                              `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	Extra       *string                                           `json:"extra,omitempty" xml:"extra,omitempty"`
	Sku         *ProductDraftQueryResponseDataProductsItemSku     `json:"sku,omitempty" xml:"sku,omitempty"`
	AuditMsg    *string                                           `json:"audit_msg,omitempty" xml:"audit_msg,omitempty"`
	Product     *ProductDraftQueryResponseDataProductsItemProduct `json:"product,omitempty" xml:"product,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItem) SetDraftStatus(v int) *ProductDraftQueryResponseDataProductsItem {
	s.DraftStatus = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItem) SetExtra(v string) *ProductDraftQueryResponseDataProductsItem {
	s.Extra = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItem) SetSku(v *ProductDraftQueryResponseDataProductsItemSku) *ProductDraftQueryResponseDataProductsItem {
	s.Sku = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItem) SetAuditMsg(v string) *ProductDraftQueryResponseDataProductsItem {
	s.AuditMsg = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItem) SetProduct(v *ProductDraftQueryResponseDataProductsItemProduct) *ProductDraftQueryResponseDataProductsItem {
	s.Product = v
	return s
}

type ProductDraftQueryResponseDataProductsItemProduct struct {
	OutUrl           *string                                                     `json:"out_url,omitempty" xml:"out_url,omitempty"`
	Telephone        []*string                                                   `json:"telephone,omitempty" xml:"telephone,omitempty" type:"Repeated"`
	Pois             []*ProductDraftQueryResponseDataProductsItemProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	AttrKeyValueMap  map[string]*string                                          `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	ProductName      *string                                                     `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	Desc             *string                                                     `json:"desc,omitempty" xml:"desc,omitempty"`
	CategoryFullName *string                                                     `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	SpuId            *string                                                     `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SoldEndTime      *int64                                                      `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	ProductSubType   *int                                                        `json:"product_sub_type,omitempty" xml:"product_sub_type,omitempty"`
	OwnerAccountId   *int64                                                      `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	ProductId        *string                                                     `json:"product_id,omitempty" xml:"product_id,omitempty"`
	BizLine          *int                                                        `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	Extra            *string                                                     `json:"extra,omitempty" xml:"extra,omitempty"`
	Version          *int64                                                      `json:"version,omitempty" xml:"version,omitempty"`
	ContactName      *string                                                     `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CreateTime       *int64                                                      `json:"create_time,omitempty" xml:"create_time,omitempty"`
	CategoryId       *int64                                                      `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	OutId            *string                                                     `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductType      *int                                                        `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	AccountName      *string                                                     `json:"account_name,omitempty" xml:"account_name,omitempty"`
	CreatorAccountId *int64                                                      `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	SoldStartTime    *int64                                                      `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	UpdateTime       *int64                                                      `json:"update_time,omitempty" xml:"update_time,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItemProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemProduct) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetOutUrl(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.OutUrl = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetTelephone(v []*string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.Telephone = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetPois(v []*ProductDraftQueryResponseDataProductsItemProductPoisItem) *ProductDraftQueryResponseDataProductsItemProduct {
	s.Pois = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetAttrKeyValueMap(v map[string]*string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetProductName(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.ProductName = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetDesc(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.Desc = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetCategoryFullName(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.CategoryFullName = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetSpuId(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.SpuId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetSoldEndTime(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.SoldEndTime = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetProductSubType(v int) *ProductDraftQueryResponseDataProductsItemProduct {
	s.ProductSubType = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetOwnerAccountId(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetProductId(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.ProductId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetBizLine(v int) *ProductDraftQueryResponseDataProductsItemProduct {
	s.BizLine = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetExtra(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.Extra = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetVersion(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.Version = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetContactName(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.ContactName = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetCreateTime(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.CreateTime = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetCategoryId(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.CategoryId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetOutId(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.OutId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetProductType(v int) *ProductDraftQueryResponseDataProductsItemProduct {
	s.ProductType = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetAccountName(v string) *ProductDraftQueryResponseDataProductsItemProduct {
	s.AccountName = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetCreatorAccountId(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetSoldStartTime(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.SoldStartTime = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProduct) SetUpdateTime(v int64) *ProductDraftQueryResponseDataProductsItemProduct {
	s.UpdateTime = &v
	return s
}

type ProductDraftQueryResponseDataProductsItemProductPoisItem struct {
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItemProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemProductPoisItem) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemProductPoisItem) SetPoiId(v string) *ProductDraftQueryResponseDataProductsItemProductPoisItem {
	s.PoiId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProductPoisItem) SetSupplierExtId(v string) *ProductDraftQueryResponseDataProductsItemProductPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemProductPoisItem) SetSupplierId(v int64) *ProductDraftQueryResponseDataProductsItemProductPoisItem {
	s.SupplierId = &v
	return s
}

type ProductDraftQueryResponseDataProductsItemSku struct {
	CreateTime      *int64                                              `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Extra           *string                                             `json:"extra,omitempty" xml:"extra,omitempty"`
	OutSkuId        *string                                             `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	OriginAmount    *int64                                              `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	Status          *int                                                `json:"status,omitempty" xml:"status,omitempty"`
	Stock           *ProductDraftQueryResponseDataProductsItemSkuStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	AttrKeyValueMap map[string]*string                                  `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	UpdateTime      *int64                                              `json:"update_time,omitempty" xml:"update_time,omitempty"`
	SkuName         *string                                             `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuId           *string                                             `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	ActualAmount    *int64                                              `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	BindSkus        []*string                                           `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuExt          *ProductDraftQueryResponseDataProductsItemSkuSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItemSku) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemSku) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetCreateTime(v int64) *ProductDraftQueryResponseDataProductsItemSku {
	s.CreateTime = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetExtra(v string) *ProductDraftQueryResponseDataProductsItemSku {
	s.Extra = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetOutSkuId(v string) *ProductDraftQueryResponseDataProductsItemSku {
	s.OutSkuId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetOriginAmount(v int64) *ProductDraftQueryResponseDataProductsItemSku {
	s.OriginAmount = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetStatus(v int) *ProductDraftQueryResponseDataProductsItemSku {
	s.Status = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetStock(v *ProductDraftQueryResponseDataProductsItemSkuStock) *ProductDraftQueryResponseDataProductsItemSku {
	s.Stock = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetAttrKeyValueMap(v map[string]*string) *ProductDraftQueryResponseDataProductsItemSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetUpdateTime(v int64) *ProductDraftQueryResponseDataProductsItemSku {
	s.UpdateTime = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetSkuName(v string) *ProductDraftQueryResponseDataProductsItemSku {
	s.SkuName = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetSkuId(v string) *ProductDraftQueryResponseDataProductsItemSku {
	s.SkuId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetActualAmount(v int64) *ProductDraftQueryResponseDataProductsItemSku {
	s.ActualAmount = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetBindSkus(v []*string) *ProductDraftQueryResponseDataProductsItemSku {
	s.BindSkus = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSku) SetSkuExt(v *ProductDraftQueryResponseDataProductsItemSkuSkuExt) *ProductDraftQueryResponseDataProductsItemSku {
	s.SkuExt = v
	return s
}

type ProductDraftQueryResponseDataProductsItemSkuSkuExt struct {
	TakeawayPresaleInfo *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	RelRuleList         []*ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	LifeBizCode         *string                                                                `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	OriginStockQty      *int64                                                                 `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	BindSkus2c          map[int64][]*int64                                                     `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	BizId2cList         []*int64                                                               `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	OriSkus             map[int64][]*int64                                                     `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	BindSkuId           *int64                                                                 `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	BindProductId       *int64                                                                 `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExt) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExt) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetTakeawayPresaleInfo(v *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetRelRuleList(v []*ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.RelRuleList = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetLifeBizCode(v string) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetOriginStockQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetBindSkus2c(v map[int64][]*int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetBizId2cList(v []*int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.BizId2cList = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetOriSkus(v map[int64][]*int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.OriSkus = v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetBindSkuId(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExt) SetBindProductId(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExt {
	s.BindProductId = &v
	return s
}

type ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem struct {
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetCoefficient(v string) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetConstantVal(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetPriceRel(v bool) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetSharedQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetStockRel(v bool) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetBizId(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

type ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *ProductDraftQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

type ProductDraftQueryResponseDataProductsItemSkuStock struct {
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
}

func (s ProductDraftQueryResponseDataProductsItemSkuStock) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseDataProductsItemSkuStock) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetAvailQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.AvailQty = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetFrozenQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.FrozenQty = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetLimitType(v int) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.LimitType = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetSoldCount(v int64) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.SoldCount = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetSoldQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.SoldQty = &v
	return s
}

func (s *ProductDraftQueryResponseDataProductsItemSkuStock) SetStockQty(v int64) *ProductDraftQueryResponseDataProductsItemSkuStock {
	s.StockQty = &v
	return s
}

type ProductDraftQueryResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ProductDraftQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductDraftQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductDraftQueryResponseExtra) SetDescription(v string) *ProductDraftQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductDraftQueryResponseExtra) SetErrorCode(v int32) *ProductDraftQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductDraftQueryResponseExtra) SetLogid(v string) *ProductDraftQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductDraftQueryResponseExtra) SetNow(v int64) *ProductDraftQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductDraftQueryResponseExtra) SetSubDescription(v string) *ProductDraftQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductDraftQueryResponseExtra) SetSubErrorCode(v int32) *ProductDraftQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ProductGetPoisRequest struct {
	Header          map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base            *ProductGetPoisRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *string                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	NeedSellOutInfo *bool                      `json:"need_sell_out_info,omitempty" xml:"need_sell_out_info,omitempty"`
	ProductIds      []*int64                   `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
}

func (s ProductGetPoisRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisRequest) GoString() string {
	return s.String()
}

func (s *ProductGetPoisRequest) SetHeader(v map[string]*string) *ProductGetPoisRequest {
	s.Header = v
	return s
}

func (s *ProductGetPoisRequest) SetAccessToken(v string) *ProductGetPoisRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductGetPoisRequest) SetBase(v *ProductGetPoisRequestBase) *ProductGetPoisRequest {
	s.Base = v
	return s
}

func (s *ProductGetPoisRequest) SetAccountId(v string) *ProductGetPoisRequest {
	s.AccountId = &v
	return s
}

func (s *ProductGetPoisRequest) SetNeedSellOutInfo(v bool) *ProductGetPoisRequest {
	s.NeedSellOutInfo = &v
	return s
}

func (s *ProductGetPoisRequest) SetProductIds(v []*int64) *ProductGetPoisRequest {
	s.ProductIds = v
	return s
}

type ProductGetPoisRequestBase struct {
	Client     *string                              `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                   `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                              `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductGetPoisRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                              `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                              `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s ProductGetPoisRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisRequestBase) GoString() string {
	return s.String()
}

func (s *ProductGetPoisRequestBase) SetClient(v string) *ProductGetPoisRequestBase {
	s.Client = &v
	return s
}

func (s *ProductGetPoisRequestBase) SetExtra(v map[string]*string) *ProductGetPoisRequestBase {
	s.Extra = v
	return s
}

func (s *ProductGetPoisRequestBase) SetLogID(v string) *ProductGetPoisRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductGetPoisRequestBase) SetTrafficEnv(v *ProductGetPoisRequestBaseTrafficEnv) *ProductGetPoisRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductGetPoisRequestBase) SetAddr(v string) *ProductGetPoisRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductGetPoisRequestBase) SetCaller(v string) *ProductGetPoisRequestBase {
	s.Caller = &v
	return s
}

type ProductGetPoisRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductGetPoisRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductGetPoisRequestBaseTrafficEnv) SetEnv(v string) *ProductGetPoisRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductGetPoisRequestBaseTrafficEnv) SetOpen(v bool) *ProductGetPoisRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductGetPoisResponse struct {
	Data     *ProductGetPoisResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductGetPoisResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *ProductGetPoisResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s ProductGetPoisResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponse) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponse) SetData(v *ProductGetPoisResponseData) *ProductGetPoisResponse {
	s.Data = v
	return s
}

func (s *ProductGetPoisResponse) SetExtra(v *ProductGetPoisResponseExtra) *ProductGetPoisResponse {
	s.Extra = v
	return s
}

func (s *ProductGetPoisResponse) SetBaseResp(v *ProductGetPoisResponseBaseResp) *ProductGetPoisResponse {
	s.BaseResp = v
	return s
}

type ProductGetPoisResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductGetPoisResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponseBaseResp) SetStatusCode(v int32) *ProductGetPoisResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductGetPoisResponseBaseResp) SetStatusMessage(v string) *ProductGetPoisResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductGetPoisResponseBaseResp) SetExtra(v map[string]*string) *ProductGetPoisResponseBaseResp {
	s.Extra = v
	return s
}

type ProductGetPoisResponseData struct {
	ErrorCode     *int32                                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductPoiMap map[int64][]*ProductGetPoisResponseDataProductPoiMapValueItem `json:"product_poi_map,omitempty" xml:"product_poi_map,omitempty"`
	Description   *string                                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductGetPoisResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponseData) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponseData) SetErrorCode(v int32) *ProductGetPoisResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductGetPoisResponseData) SetProductPoiMap(v map[int64][]*ProductGetPoisResponseDataProductPoiMapValueItem) *ProductGetPoisResponseData {
	s.ProductPoiMap = v
	return s
}

func (s *ProductGetPoisResponseData) SetDescription(v string) *ProductGetPoisResponseData {
	s.Description = &v
	return s
}

type ProductGetPoisResponseDataProductPoiMapValueItem struct {
	SellOutInfo *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo `json:"sell_out_info,omitempty" xml:"sell_out_info,omitempty"`
	PoiId       *int64                                                       `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
}

func (s ProductGetPoisResponseDataProductPoiMapValueItem) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponseDataProductPoiMapValueItem) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponseDataProductPoiMapValueItem) SetSellOutInfo(v *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) *ProductGetPoisResponseDataProductPoiMapValueItem {
	s.SellOutInfo = v
	return s
}

func (s *ProductGetPoisResponseDataProductPoiMapValueItem) SetPoiId(v int64) *ProductGetPoisResponseDataProductPoiMapValueItem {
	s.PoiId = &v
	return s
}

type ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo struct {
	SellOutEndTime   *int64 `json:"sell_out_end_time,omitempty" xml:"sell_out_end_time,omitempty"`
	SellOutStartTime *int64 `json:"sell_out_start_time,omitempty" xml:"sell_out_start_time,omitempty"`
	SellOutStatus    *int64 `json:"sell_out_status,omitempty" xml:"sell_out_status,omitempty"`
}

func (s ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) SetSellOutEndTime(v int64) *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo {
	s.SellOutEndTime = &v
	return s
}

func (s *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) SetSellOutStartTime(v int64) *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo {
	s.SellOutStartTime = &v
	return s
}

func (s *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo) SetSellOutStatus(v int64) *ProductGetPoisResponseDataProductPoiMapValueItemSellOutInfo {
	s.SellOutStatus = &v
	return s
}

type ProductGetPoisResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ProductGetPoisResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductGetPoisResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductGetPoisResponseExtra) SetSubErrorCode(v int32) *ProductGetPoisResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductGetPoisResponseExtra) SetDescription(v string) *ProductGetPoisResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductGetPoisResponseExtra) SetErrorCode(v int32) *ProductGetPoisResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductGetPoisResponseExtra) SetLogid(v string) *ProductGetPoisResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductGetPoisResponseExtra) SetNow(v int64) *ProductGetPoisResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductGetPoisResponseExtra) SetSubDescription(v string) *ProductGetPoisResponseExtra {
	s.SubDescription = &v
	return s
}

type ProductOnlineGetRequest struct {
	AccessToken *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductIds  []*int64                     `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base        *ProductOnlineGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                      `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ProductOnlineGetRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetRequest) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetRequest) SetAccessToken(v string) *ProductOnlineGetRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductOnlineGetRequest) SetProductIds(v []*int64) *ProductOnlineGetRequest {
	s.ProductIds = v
	return s
}

func (s *ProductOnlineGetRequest) SetBase(v *ProductOnlineGetRequestBase) *ProductOnlineGetRequest {
	s.Base = v
	return s
}

func (s *ProductOnlineGetRequest) SetAccountId(v string) *ProductOnlineGetRequest {
	s.AccountId = &v
	return s
}

func (s *ProductOnlineGetRequest) SetHeader(v map[string]*string) *ProductOnlineGetRequest {
	s.Header = v
	return s
}

type ProductOnlineGetRequestBase struct {
	TrafficEnv *ProductOnlineGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                     `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s ProductOnlineGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetRequestBase) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetRequestBase) SetTrafficEnv(v *ProductOnlineGetRequestBaseTrafficEnv) *ProductOnlineGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductOnlineGetRequestBase) SetAddr(v string) *ProductOnlineGetRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductOnlineGetRequestBase) SetCaller(v string) *ProductOnlineGetRequestBase {
	s.Caller = &v
	return s
}

func (s *ProductOnlineGetRequestBase) SetClient(v string) *ProductOnlineGetRequestBase {
	s.Client = &v
	return s
}

func (s *ProductOnlineGetRequestBase) SetExtra(v map[string]*string) *ProductOnlineGetRequestBase {
	s.Extra = v
	return s
}

func (s *ProductOnlineGetRequestBase) SetLogID(v string) *ProductOnlineGetRequestBase {
	s.LogID = &v
	return s
}
