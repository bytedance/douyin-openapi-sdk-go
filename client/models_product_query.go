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

type ProductOnlineGetRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductOnlineGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetRequestBaseTrafficEnv) SetEnv(v string) *ProductOnlineGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductOnlineGetRequestBaseTrafficEnv) SetOpen(v bool) *ProductOnlineGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductOnlineGetResponse struct {
	Data     *ProductOnlineGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductOnlineGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *ProductOnlineGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s ProductOnlineGetResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponse) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponse) SetData(v *ProductOnlineGetResponseData) *ProductOnlineGetResponse {
	s.Data = v
	return s
}

func (s *ProductOnlineGetResponse) SetExtra(v *ProductOnlineGetResponseExtra) *ProductOnlineGetResponse {
	s.Extra = v
	return s
}

func (s *ProductOnlineGetResponse) SetBaseResp(v *ProductOnlineGetResponseBaseResp) *ProductOnlineGetResponse {
	s.BaseResp = v
	return s
}

type ProductOnlineGetResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductOnlineGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseBaseResp) SetStatusCode(v int32) *ProductOnlineGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductOnlineGetResponseBaseResp) SetStatusMessage(v string) *ProductOnlineGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductOnlineGetResponseBaseResp) SetExtra(v map[string]*string) *ProductOnlineGetResponseBaseResp {
	s.Extra = v
	return s
}

type ProductOnlineGetResponseData struct {
	ErrorCode   *int32                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Products    []*ProductOnlineGetResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Description *string                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductOnlineGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseData) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseData) SetErrorCode(v int32) *ProductOnlineGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductOnlineGetResponseData) SetProducts(v []*ProductOnlineGetResponseDataProductsItem) *ProductOnlineGetResponseData {
	s.Products = v
	return s
}

func (s *ProductOnlineGetResponseData) SetDescription(v string) *ProductOnlineGetResponseData {
	s.Description = &v
	return s
}

type ProductOnlineGetResponseDataProductsItem struct {
	CreateTime            *int64                                                              `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Attributes            []*ProductOnlineGetResponseDataProductsItemAttributesItem           `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	IsSuperSkuSub         *bool                                                               `json:"is_super_sku_sub,omitempty" xml:"is_super_sku_sub,omitempty"`
	LimitBuyRule          *ProductOnlineGetResponseDataProductsItemLimitBuyRule               `json:"limit_buy_rule,omitempty" xml:"limit_buy_rule,omitempty"`
	Skus                  []*ProductOnlineGetResponseDataProductsItemSkusItem                 `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	PrimaryMaterials      []*ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem     `json:"primary_materials,omitempty" xml:"primary_materials,omitempty" type:"Repeated"`
	PoiCount              *int64                                                              `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	BizLine               *int                                                                `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	SoldEndTime           *int64                                                              `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	DishesImageList       []*ProductOnlineGetResponseDataProductsItemDishesImageListItem      `json:"dishes_image_list,omitempty" xml:"dishes_image_list,omitempty" type:"Repeated"`
	CategoryFullName      *string                                                             `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	NotIndependentSale    *bool                                                               `json:"not_independent_sale,omitempty" xml:"not_independent_sale,omitempty"`
	SoldStartTime         *int64                                                              `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	RecPersonNum          *int64                                                              `json:"rec_person_num,omitempty" xml:"rec_person_num,omitempty"`
	ApplyDate             *ProductOnlineGetResponseDataProductsItemApplyDate                  `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	SettleType            *int64                                                              `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	ShowChannel           *int                                                                `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	SuperimposedDiscounts *bool                                                               `json:"superimposed_discounts,omitempty" xml:"superimposed_discounts,omitempty"`
	ComboGroups           []*ProductOnlineGetResponseDataProductsItemComboGroupsItem          `json:"combo_groups,omitempty" xml:"combo_groups,omitempty" type:"Repeated"`
	DeliveryMethod        []*int                                                              `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	EnvironmentImageList  []*ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem `json:"environment_image_list,omitempty" xml:"environment_image_list,omitempty" type:"Repeated"`
	OutId                 *string                                                             `json:"out_id,omitempty" xml:"out_id,omitempty"`
	AutoRenew             *bool                                                               `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	Commodity             *string                                                             `json:"commodity,omitempty" xml:"commodity,omitempty"`
	ProductSpecAttrs      []*ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem     `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	AddDishGroups         []*ProductOnlineGetResponseDataProductsItemAddDishGroupsItem        `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	CategoryId            *int64                                                              `json:"category_id,omitempty" xml:"category_id,omitempty"`
	AccountId             *string                                                             `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName           *string                                                             `json:"account_name,omitempty" xml:"account_name,omitempty"`
	UpdateTime            *int64                                                              `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductExt            *ProductOnlineGetResponseDataProductsItemProductExt                 `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	PatternType           *int64                                                              `json:"pattern_type,omitempty" xml:"pattern_type,omitempty"`
	DescriptionRichText   []*string                                                           `json:"description_rich_text,omitempty" xml:"description_rich_text,omitempty" type:"Repeated"`
	OnlineStatus          *int                                                                `json:"online_status,omitempty" xml:"online_status,omitempty"`
	ProductName           *string                                                             `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductType           *int                                                                `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ImageList             []*ProductOnlineGetResponseDataProductsItemImageListItem            `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	ProductId             *int64                                                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	RecPersonNumMax       *int64                                                              `json:"rec_person_num_max,omitempty" xml:"rec_person_num_max,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItem) SetCreateTime(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.CreateTime = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetAttributes(v []*ProductOnlineGetResponseDataProductsItemAttributesItem) *ProductOnlineGetResponseDataProductsItem {
	s.Attributes = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetIsSuperSkuSub(v bool) *ProductOnlineGetResponseDataProductsItem {
	s.IsSuperSkuSub = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetLimitBuyRule(v *ProductOnlineGetResponseDataProductsItemLimitBuyRule) *ProductOnlineGetResponseDataProductsItem {
	s.LimitBuyRule = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetSkus(v []*ProductOnlineGetResponseDataProductsItemSkusItem) *ProductOnlineGetResponseDataProductsItem {
	s.Skus = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetPrimaryMaterials(v []*ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem) *ProductOnlineGetResponseDataProductsItem {
	s.PrimaryMaterials = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetPoiCount(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.PoiCount = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetBizLine(v int) *ProductOnlineGetResponseDataProductsItem {
	s.BizLine = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetSoldEndTime(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.SoldEndTime = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetDishesImageList(v []*ProductOnlineGetResponseDataProductsItemDishesImageListItem) *ProductOnlineGetResponseDataProductsItem {
	s.DishesImageList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetCategoryFullName(v string) *ProductOnlineGetResponseDataProductsItem {
	s.CategoryFullName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetNotIndependentSale(v bool) *ProductOnlineGetResponseDataProductsItem {
	s.NotIndependentSale = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetSoldStartTime(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.SoldStartTime = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetRecPersonNum(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.RecPersonNum = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetApplyDate(v *ProductOnlineGetResponseDataProductsItemApplyDate) *ProductOnlineGetResponseDataProductsItem {
	s.ApplyDate = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetSettleType(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.SettleType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetShowChannel(v int) *ProductOnlineGetResponseDataProductsItem {
	s.ShowChannel = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetSuperimposedDiscounts(v bool) *ProductOnlineGetResponseDataProductsItem {
	s.SuperimposedDiscounts = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetComboGroups(v []*ProductOnlineGetResponseDataProductsItemComboGroupsItem) *ProductOnlineGetResponseDataProductsItem {
	s.ComboGroups = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetDeliveryMethod(v []*int) *ProductOnlineGetResponseDataProductsItem {
	s.DeliveryMethod = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetEnvironmentImageList(v []*ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem) *ProductOnlineGetResponseDataProductsItem {
	s.EnvironmentImageList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetOutId(v string) *ProductOnlineGetResponseDataProductsItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetAutoRenew(v bool) *ProductOnlineGetResponseDataProductsItem {
	s.AutoRenew = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetCommodity(v string) *ProductOnlineGetResponseDataProductsItem {
	s.Commodity = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetProductSpecAttrs(v []*ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem) *ProductOnlineGetResponseDataProductsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetAddDishGroups(v []*ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) *ProductOnlineGetResponseDataProductsItem {
	s.AddDishGroups = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetCategoryId(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.CategoryId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetAccountId(v string) *ProductOnlineGetResponseDataProductsItem {
	s.AccountId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetAccountName(v string) *ProductOnlineGetResponseDataProductsItem {
	s.AccountName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetUpdateTime(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.UpdateTime = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetProductExt(v *ProductOnlineGetResponseDataProductsItemProductExt) *ProductOnlineGetResponseDataProductsItem {
	s.ProductExt = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetPatternType(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.PatternType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetDescriptionRichText(v []*string) *ProductOnlineGetResponseDataProductsItem {
	s.DescriptionRichText = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetOnlineStatus(v int) *ProductOnlineGetResponseDataProductsItem {
	s.OnlineStatus = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetProductName(v string) *ProductOnlineGetResponseDataProductsItem {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetProductType(v int) *ProductOnlineGetResponseDataProductsItem {
	s.ProductType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetImageList(v []*ProductOnlineGetResponseDataProductsItemImageListItem) *ProductOnlineGetResponseDataProductsItem {
	s.ImageList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetProductId(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItem) SetRecPersonNumMax(v int64) *ProductOnlineGetResponseDataProductsItem {
	s.RecPersonNumMax = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemAddDishGroupsItem struct {
	GroupId   *int64                                                                   `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                                  `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) SetGroupId(v int64) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) SetGroupName(v string) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem) SetItemList(v []*ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

type ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem struct {
	OutId       *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) SetOutId(v string) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem) SetProductName(v string) *ProductOnlineGetResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemApplyDate struct {
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemApplyDate) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemApplyDate) SetDayDuration(v int32) *ProductOnlineGetResponseDataProductsItemApplyDate {
	s.DayDuration = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemApplyDate) SetUseDateType(v int) *ProductOnlineGetResponseDataProductsItemApplyDate {
	s.UseDateType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemApplyDate) SetUseEndDate(v string) *ProductOnlineGetResponseDataProductsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemApplyDate) SetUseStartDate(v string) *ProductOnlineGetResponseDataProductsItemApplyDate {
	s.UseStartDate = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemAttributesItem struct {
	GroupName *string                                                               `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineGetResponseDataProductsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemAttributesItem) SetGroupName(v string) *ProductOnlineGetResponseDataProductsItemAttributesItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemAttributesItem) SetItemList(v []*ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem) *ProductOnlineGetResponseDataProductsItemAttributesItem {
	s.ItemList = v
	return s
}

type ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem) SetName(v string) *ProductOnlineGetResponseDataProductsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemComboGroupsItem struct {
	ItemList  []*ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupId   *int64                                                                 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                                `json:"group_name,omitempty" xml:"group_name,omitempty"`
	GroupRule *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule      `json:"group_rule,omitempty" xml:"group_rule,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItem) SetItemList(v []*ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) *ProductOnlineGetResponseDataProductsItemComboGroupsItem {
	s.ItemList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItem) SetGroupId(v int64) *ProductOnlineGetResponseDataProductsItemComboGroupsItem {
	s.GroupId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItem) SetGroupName(v string) *ProductOnlineGetResponseDataProductsItemComboGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItem) SetGroupRule(v *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule) *ProductOnlineGetResponseDataProductsItemComboGroupsItem {
	s.GroupRule = v
	return s
}

type ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule struct {
	OptionCount *int32 `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule) SetOptionCount(v int32) *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule {
	s.OptionCount = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule) SetTotalCount(v int32) *ProductOnlineGetResponseDataProductsItemComboGroupsItemGroupRule {
	s.TotalCount = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem struct {
	Unit        *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Count       *int64  `json:"count,omitempty" xml:"count,omitempty"`
	OutId       *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetUnit(v string) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetCount(v int64) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.Count = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetOutId(v string) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetPrice(v int64) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetProductId(v int64) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem) SetProductName(v string) *ProductOnlineGetResponseDataProductsItemComboGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemDishesImageListItem struct {
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemDishesImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemDishesImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemDishesImageListItem) SetUrl(v string) *ProductOnlineGetResponseDataProductsItemDishesImageListItem {
	s.Url = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemDishesImageListItem) SetName(v string) *ProductOnlineGetResponseDataProductsItemDishesImageListItem {
	s.Name = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem) SetName(v string) *ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem {
	s.Name = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem) SetUrl(v string) *ProductOnlineGetResponseDataProductsItemEnvironmentImageListItem {
	s.Url = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemImageListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemImageListItem) SetName(v string) *ProductOnlineGetResponseDataProductsItemImageListItem {
	s.Name = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemImageListItem) SetUrl(v string) *ProductOnlineGetResponseDataProductsItemImageListItem {
	s.Url = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemLimitBuyRule struct {
	RuleList []*ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem `json:"rule_list,omitempty" xml:"rule_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineGetResponseDataProductsItemLimitBuyRule) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemLimitBuyRule) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemLimitBuyRule) SetRuleList(v []*ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) *ProductOnlineGetResponseDataProductsItemLimitBuyRule {
	s.RuleList = v
	return s
}

type ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem struct {
	RangeType   *int   `json:"range_type,omitempty" xml:"range_type,omitempty" require:"true"`
	SubjectType *int   `json:"subject_type,omitempty" xml:"subject_type,omitempty" require:"true"`
	LimitNum    *int32 `json:"limit_num,omitempty" xml:"limit_num,omitempty" require:"true"`
}

func (s ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) SetRangeType(v int) *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.RangeType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) SetSubjectType(v int) *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.SubjectType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem) SetLimitNum(v int32) *ProductOnlineGetResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.LimitNum = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem) SetName(v string) *ProductOnlineGetResponseDataProductsItemPrimaryMaterialsItem {
	s.Name = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemProductExt struct {
	TestExtra *ProductOnlineGetResponseDataProductsItemProductExtTestExtra `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemProductExt) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemProductExt) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemProductExt) SetTestExtra(v *ProductOnlineGetResponseDataProductsItemProductExtTestExtra) *ProductOnlineGetResponseDataProductsItemProductExt {
	s.TestExtra = v
	return s
}

type ProductOnlineGetResponseDataProductsItemProductExtTestExtra struct {
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemProductExtTestExtra) SetUids(v []*string) *ProductOnlineGetResponseDataProductsItemProductExtTestExtra {
	s.Uids = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemProductExtTestExtra) SetTestFlag(v bool) *ProductOnlineGetResponseDataProductsItemProductExtTestExtra {
	s.TestFlag = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem struct {
	ItemList  []*ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                     `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem) SetItemList(v []*ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem) SetGroupName(v string) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem struct {
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *ProductOnlineGetResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemSkusItem struct {
	SkuSpecAttrs []*ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	Status       *int                                                                `json:"status,omitempty" xml:"status,omitempty"`
	Stock        *ProductOnlineGetResponseDataProductsItemSkusItemStock              `json:"stock,omitempty" xml:"stock,omitempty"`
	ActualAmount *int64                                                              `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64                                                              `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	OutSkuId     *string                                                             `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId        *int64                                                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName      *string                                                             `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemSkusItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetSkuSpecAttrs(v []*ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetStatus(v int) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.Status = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetStock(v *ProductOnlineGetResponseDataProductsItemSkusItemStock) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.Stock = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetActualAmount(v int64) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetOriginAmount(v int64) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetOutSkuId(v string) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetSkuId(v int64) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItem) SetSkuName(v string) *ProductOnlineGetResponseDataProductsItemSkusItem {
	s.SkuName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem struct {
	ItemList  []*ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                         `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *ProductOnlineGetResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

type ProductOnlineGetResponseDataProductsItemSkusItemStock struct {
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseDataProductsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetSoldQty(v int64) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetStockQty(v int64) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetAvailQty(v int64) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetFrozenQty(v int64) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetLimitType(v int) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *ProductOnlineGetResponseDataProductsItemSkusItemStock) SetSoldCount(v int64) *ProductOnlineGetResponseDataProductsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

type ProductOnlineGetResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ProductOnlineGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineGetResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineGetResponseExtra) SetSubErrorCode(v int32) *ProductOnlineGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductOnlineGetResponseExtra) SetDescription(v string) *ProductOnlineGetResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductOnlineGetResponseExtra) SetErrorCode(v int32) *ProductOnlineGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductOnlineGetResponseExtra) SetLogid(v string) *ProductOnlineGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductOnlineGetResponseExtra) SetNow(v int64) *ProductOnlineGetResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductOnlineGetResponseExtra) SetSubDescription(v string) *ProductOnlineGetResponseExtra {
	s.SubDescription = &v
	return s
}

type ProductOnlineListRequest struct {
	Status       []*int                        `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
	PageNo       *int32                        `json:"page_no,omitempty" xml:"page_no,omitempty"`
	SearchByPage *bool                         `json:"search_by_page,omitempty" xml:"search_by_page,omitempty"`
	AccountId    *string                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Cursor       *int32                        `json:"cursor,omitempty" xml:"cursor,omitempty"`
	DyPoiId      *int64                        `json:"dy_poi_id,omitempty" xml:"dy_poi_id,omitempty"`
	PageSize     *int32                        `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Count        *int32                        `json:"count,omitempty" xml:"count,omitempty"`
	Base         *ProductOnlineListRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header       map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ProductOnlineListRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListRequest) GoString() string {
	return s.String()
}

func (s *ProductOnlineListRequest) SetStatus(v []*int) *ProductOnlineListRequest {
	s.Status = v
	return s
}

func (s *ProductOnlineListRequest) SetPageNo(v int32) *ProductOnlineListRequest {
	s.PageNo = &v
	return s
}

func (s *ProductOnlineListRequest) SetSearchByPage(v bool) *ProductOnlineListRequest {
	s.SearchByPage = &v
	return s
}

func (s *ProductOnlineListRequest) SetAccountId(v string) *ProductOnlineListRequest {
	s.AccountId = &v
	return s
}

func (s *ProductOnlineListRequest) SetCursor(v int32) *ProductOnlineListRequest {
	s.Cursor = &v
	return s
}

func (s *ProductOnlineListRequest) SetDyPoiId(v int64) *ProductOnlineListRequest {
	s.DyPoiId = &v
	return s
}

func (s *ProductOnlineListRequest) SetPageSize(v int32) *ProductOnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *ProductOnlineListRequest) SetCount(v int32) *ProductOnlineListRequest {
	s.Count = &v
	return s
}

func (s *ProductOnlineListRequest) SetBase(v *ProductOnlineListRequestBase) *ProductOnlineListRequest {
	s.Base = v
	return s
}

func (s *ProductOnlineListRequest) SetHeader(v map[string]*string) *ProductOnlineListRequest {
	s.Header = v
	return s
}

func (s *ProductOnlineListRequest) SetAccessToken(v string) *ProductOnlineListRequest {
	s.AccessToken = &v
	return s
}

type ProductOnlineListRequestBase struct {
	Caller     *string                                 `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                 `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                 `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductOnlineListRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                 `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s ProductOnlineListRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListRequestBase) GoString() string {
	return s.String()
}

func (s *ProductOnlineListRequestBase) SetCaller(v string) *ProductOnlineListRequestBase {
	s.Caller = &v
	return s
}

func (s *ProductOnlineListRequestBase) SetClient(v string) *ProductOnlineListRequestBase {
	s.Client = &v
	return s
}

func (s *ProductOnlineListRequestBase) SetExtra(v map[string]*string) *ProductOnlineListRequestBase {
	s.Extra = v
	return s
}

func (s *ProductOnlineListRequestBase) SetLogID(v string) *ProductOnlineListRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductOnlineListRequestBase) SetTrafficEnv(v *ProductOnlineListRequestBaseTrafficEnv) *ProductOnlineListRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductOnlineListRequestBase) SetAddr(v string) *ProductOnlineListRequestBase {
	s.Addr = &v
	return s
}

type ProductOnlineListRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s ProductOnlineListRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductOnlineListRequestBaseTrafficEnv) SetOpen(v bool) *ProductOnlineListRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *ProductOnlineListRequestBaseTrafficEnv) SetEnv(v string) *ProductOnlineListRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type ProductOnlineListResponse struct {
	Data     *ProductOnlineListResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductOnlineListResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *ProductOnlineListResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s ProductOnlineListResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponse) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponse) SetData(v *ProductOnlineListResponseData) *ProductOnlineListResponse {
	s.Data = v
	return s
}

func (s *ProductOnlineListResponse) SetExtra(v *ProductOnlineListResponseExtra) *ProductOnlineListResponse {
	s.Extra = v
	return s
}

func (s *ProductOnlineListResponse) SetBaseResp(v *ProductOnlineListResponseBaseResp) *ProductOnlineListResponse {
	s.BaseResp = v
	return s
}

type ProductOnlineListResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductOnlineListResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseBaseResp) SetStatusCode(v int32) *ProductOnlineListResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductOnlineListResponseBaseResp) SetStatusMessage(v string) *ProductOnlineListResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductOnlineListResponseBaseResp) SetExtra(v map[string]*string) *ProductOnlineListResponseBaseResp {
	s.Extra = v
	return s
}

type ProductOnlineListResponseData struct {
	Products    []*ProductOnlineListResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Total       *int32                                       `json:"total,omitempty" xml:"total,omitempty"`
	Description *string                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	HasMore     *bool                                        `json:"has_more,omitempty" xml:"has_more,omitempty"`
	NextCursor  *int32                                       `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
}

func (s ProductOnlineListResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseData) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseData) SetProducts(v []*ProductOnlineListResponseDataProductsItem) *ProductOnlineListResponseData {
	s.Products = v
	return s
}

func (s *ProductOnlineListResponseData) SetTotal(v int32) *ProductOnlineListResponseData {
	s.Total = &v
	return s
}

func (s *ProductOnlineListResponseData) SetDescription(v string) *ProductOnlineListResponseData {
	s.Description = &v
	return s
}

func (s *ProductOnlineListResponseData) SetErrorCode(v int32) *ProductOnlineListResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductOnlineListResponseData) SetHasMore(v bool) *ProductOnlineListResponseData {
	s.HasMore = &v
	return s
}

func (s *ProductOnlineListResponseData) SetNextCursor(v int32) *ProductOnlineListResponseData {
	s.NextCursor = &v
	return s
}

type ProductOnlineListResponseDataProductsItem struct {
	ImageList             []*ProductOnlineListResponseDataProductsItemImageListItem            `json:"image_list,omitempty" xml:"image_list,omitempty" type:"Repeated"`
	ProductExt            *ProductOnlineListResponseDataProductsItemProductExt                 `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	AccountId             *string                                                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ProductName           *string                                                              `json:"product_name,omitempty" xml:"product_name,omitempty"`
	SettleType            *int64                                                               `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	PatternType           *int64                                                               `json:"pattern_type,omitempty" xml:"pattern_type,omitempty"`
	CreateTime            *int64                                                               `json:"create_time,omitempty" xml:"create_time,omitempty"`
	NotIndependentSale    *bool                                                                `json:"not_independent_sale,omitempty" xml:"not_independent_sale,omitempty"`
	ShowChannel           *int                                                                 `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	ProductId             *int64                                                               `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Skus                  []*ProductOnlineListResponseDataProductsItemSkusItem                 `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	Commodity             *string                                                              `json:"commodity,omitempty" xml:"commodity,omitempty"`
	SoldStartTime         *int64                                                               `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	ComboGroups           []*ProductOnlineListResponseDataProductsItemComboGroupsItem          `json:"combo_groups,omitempty" xml:"combo_groups,omitempty" type:"Repeated"`
	AutoRenew             *bool                                                                `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	SuperimposedDiscounts *bool                                                                `json:"superimposed_discounts,omitempty" xml:"superimposed_discounts,omitempty"`
	RecPersonNum          *int64                                                               `json:"rec_person_num,omitempty" xml:"rec_person_num,omitempty"`
	OutId                 *string                                                              `json:"out_id,omitempty" xml:"out_id,omitempty"`
	AddDishGroups         []*ProductOnlineListResponseDataProductsItemAddDishGroupsItem        `json:"add_dish_groups,omitempty" xml:"add_dish_groups,omitempty" type:"Repeated"`
	ProductType           *int                                                                 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	RecPersonNumMax       *int64                                                               `json:"rec_person_num_max,omitempty" xml:"rec_person_num_max,omitempty"`
	BizLine               *int                                                                 `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	ApplyDate             *ProductOnlineListResponseDataProductsItemApplyDate                  `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	CategoryId            *int64                                                               `json:"category_id,omitempty" xml:"category_id,omitempty"`
	EnvironmentImageList  []*ProductOnlineListResponseDataProductsItemEnvironmentImageListItem `json:"environment_image_list,omitempty" xml:"environment_image_list,omitempty" type:"Repeated"`
	UpdateTime            *int64                                                               `json:"update_time,omitempty" xml:"update_time,omitempty"`
	SoldEndTime           *int64                                                               `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	ProductSpecAttrs      []*ProductOnlineListResponseDataProductsItemProductSpecAttrsItem     `json:"product_spec_attrs,omitempty" xml:"product_spec_attrs,omitempty" type:"Repeated"`
	IsSuperSkuSub         *bool                                                                `json:"is_super_sku_sub,omitempty" xml:"is_super_sku_sub,omitempty"`
	Attributes            []*ProductOnlineListResponseDataProductsItemAttributesItem           `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	AccountName           *string                                                              `json:"account_name,omitempty" xml:"account_name,omitempty"`
	PoiCount              *int64                                                               `json:"poi_count,omitempty" xml:"poi_count,omitempty"`
	DescriptionRichText   []*string                                                            `json:"description_rich_text,omitempty" xml:"description_rich_text,omitempty" type:"Repeated"`
	LimitBuyRule          *ProductOnlineListResponseDataProductsItemLimitBuyRule               `json:"limit_buy_rule,omitempty" xml:"limit_buy_rule,omitempty"`
	OnlineStatus          *int                                                                 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	PrimaryMaterials      []*ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem     `json:"primary_materials,omitempty" xml:"primary_materials,omitempty" type:"Repeated"`
	DishesImageList       []*ProductOnlineListResponseDataProductsItemDishesImageListItem      `json:"dishes_image_list,omitempty" xml:"dishes_image_list,omitempty" type:"Repeated"`
	DeliveryMethod        []*int                                                               `json:"delivery_method,omitempty" xml:"delivery_method,omitempty" type:"Repeated"`
	CategoryFullName      *string                                                              `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItem) SetImageList(v []*ProductOnlineListResponseDataProductsItemImageListItem) *ProductOnlineListResponseDataProductsItem {
	s.ImageList = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetProductExt(v *ProductOnlineListResponseDataProductsItemProductExt) *ProductOnlineListResponseDataProductsItem {
	s.ProductExt = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetAccountId(v string) *ProductOnlineListResponseDataProductsItem {
	s.AccountId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetProductName(v string) *ProductOnlineListResponseDataProductsItem {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetSettleType(v int64) *ProductOnlineListResponseDataProductsItem {
	s.SettleType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetPatternType(v int64) *ProductOnlineListResponseDataProductsItem {
	s.PatternType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetCreateTime(v int64) *ProductOnlineListResponseDataProductsItem {
	s.CreateTime = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetNotIndependentSale(v bool) *ProductOnlineListResponseDataProductsItem {
	s.NotIndependentSale = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetShowChannel(v int) *ProductOnlineListResponseDataProductsItem {
	s.ShowChannel = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetProductId(v int64) *ProductOnlineListResponseDataProductsItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetSkus(v []*ProductOnlineListResponseDataProductsItemSkusItem) *ProductOnlineListResponseDataProductsItem {
	s.Skus = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetCommodity(v string) *ProductOnlineListResponseDataProductsItem {
	s.Commodity = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetSoldStartTime(v int64) *ProductOnlineListResponseDataProductsItem {
	s.SoldStartTime = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetComboGroups(v []*ProductOnlineListResponseDataProductsItemComboGroupsItem) *ProductOnlineListResponseDataProductsItem {
	s.ComboGroups = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetAutoRenew(v bool) *ProductOnlineListResponseDataProductsItem {
	s.AutoRenew = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetSuperimposedDiscounts(v bool) *ProductOnlineListResponseDataProductsItem {
	s.SuperimposedDiscounts = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetRecPersonNum(v int64) *ProductOnlineListResponseDataProductsItem {
	s.RecPersonNum = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetOutId(v string) *ProductOnlineListResponseDataProductsItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetAddDishGroups(v []*ProductOnlineListResponseDataProductsItemAddDishGroupsItem) *ProductOnlineListResponseDataProductsItem {
	s.AddDishGroups = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetProductType(v int) *ProductOnlineListResponseDataProductsItem {
	s.ProductType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetRecPersonNumMax(v int64) *ProductOnlineListResponseDataProductsItem {
	s.RecPersonNumMax = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetBizLine(v int) *ProductOnlineListResponseDataProductsItem {
	s.BizLine = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetApplyDate(v *ProductOnlineListResponseDataProductsItemApplyDate) *ProductOnlineListResponseDataProductsItem {
	s.ApplyDate = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetCategoryId(v int64) *ProductOnlineListResponseDataProductsItem {
	s.CategoryId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetEnvironmentImageList(v []*ProductOnlineListResponseDataProductsItemEnvironmentImageListItem) *ProductOnlineListResponseDataProductsItem {
	s.EnvironmentImageList = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetUpdateTime(v int64) *ProductOnlineListResponseDataProductsItem {
	s.UpdateTime = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetSoldEndTime(v int64) *ProductOnlineListResponseDataProductsItem {
	s.SoldEndTime = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetProductSpecAttrs(v []*ProductOnlineListResponseDataProductsItemProductSpecAttrsItem) *ProductOnlineListResponseDataProductsItem {
	s.ProductSpecAttrs = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetIsSuperSkuSub(v bool) *ProductOnlineListResponseDataProductsItem {
	s.IsSuperSkuSub = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetAttributes(v []*ProductOnlineListResponseDataProductsItemAttributesItem) *ProductOnlineListResponseDataProductsItem {
	s.Attributes = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetAccountName(v string) *ProductOnlineListResponseDataProductsItem {
	s.AccountName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetPoiCount(v int64) *ProductOnlineListResponseDataProductsItem {
	s.PoiCount = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetDescriptionRichText(v []*string) *ProductOnlineListResponseDataProductsItem {
	s.DescriptionRichText = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetLimitBuyRule(v *ProductOnlineListResponseDataProductsItemLimitBuyRule) *ProductOnlineListResponseDataProductsItem {
	s.LimitBuyRule = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetOnlineStatus(v int) *ProductOnlineListResponseDataProductsItem {
	s.OnlineStatus = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetPrimaryMaterials(v []*ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem) *ProductOnlineListResponseDataProductsItem {
	s.PrimaryMaterials = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetDishesImageList(v []*ProductOnlineListResponseDataProductsItemDishesImageListItem) *ProductOnlineListResponseDataProductsItem {
	s.DishesImageList = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetDeliveryMethod(v []*int) *ProductOnlineListResponseDataProductsItem {
	s.DeliveryMethod = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItem) SetCategoryFullName(v string) *ProductOnlineListResponseDataProductsItem {
	s.CategoryFullName = &v
	return s
}

type ProductOnlineListResponseDataProductsItemAddDishGroupsItem struct {
	GroupId   *int64                                                                    `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName *string                                                                   `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineListResponseDataProductsItemAddDishGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemAddDishGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItem) SetGroupId(v int64) *ProductOnlineListResponseDataProductsItemAddDishGroupsItem {
	s.GroupId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItem) SetGroupName(v string) *ProductOnlineListResponseDataProductsItemAddDishGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItem) SetItemList(v []*ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) *ProductOnlineListResponseDataProductsItemAddDishGroupsItem {
	s.ItemList = v
	return s
}

type ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem struct {
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	OutId       *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) SetProductId(v int64) *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) SetProductName(v string) *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) SetOutId(v string) *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem) SetPrice(v int64) *ProductOnlineListResponseDataProductsItemAddDishGroupsItemItemListItem {
	s.Price = &v
	return s
}

type ProductOnlineListResponseDataProductsItemApplyDate struct {
	UseDateType  *int    `json:"use_date_type,omitempty" xml:"use_date_type,omitempty" require:"true"`
	UseEndDate   *string `json:"use_end_date,omitempty" xml:"use_end_date,omitempty"`
	UseStartDate *string `json:"use_start_date,omitempty" xml:"use_start_date,omitempty"`
	DayDuration  *int32  `json:"day_duration,omitempty" xml:"day_duration,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemApplyDate) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemApplyDate) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemApplyDate) SetUseDateType(v int) *ProductOnlineListResponseDataProductsItemApplyDate {
	s.UseDateType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemApplyDate) SetUseEndDate(v string) *ProductOnlineListResponseDataProductsItemApplyDate {
	s.UseEndDate = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemApplyDate) SetUseStartDate(v string) *ProductOnlineListResponseDataProductsItemApplyDate {
	s.UseStartDate = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemApplyDate) SetDayDuration(v int32) *ProductOnlineListResponseDataProductsItemApplyDate {
	s.DayDuration = &v
	return s
}

type ProductOnlineListResponseDataProductsItemAttributesItem struct {
	GroupName *string                                                                `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductOnlineListResponseDataProductsItemAttributesItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineListResponseDataProductsItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemAttributesItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemAttributesItem) SetGroupName(v string) *ProductOnlineListResponseDataProductsItemAttributesItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemAttributesItem) SetItemList(v []*ProductOnlineListResponseDataProductsItemAttributesItemItemListItem) *ProductOnlineListResponseDataProductsItemAttributesItem {
	s.ItemList = v
	return s
}

type ProductOnlineListResponseDataProductsItemAttributesItemItemListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemAttributesItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemAttributesItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemAttributesItemItemListItem) SetName(v string) *ProductOnlineListResponseDataProductsItemAttributesItemItemListItem {
	s.Name = &v
	return s
}

type ProductOnlineListResponseDataProductsItemComboGroupsItem struct {
	GroupName *string                                                                 `json:"group_name,omitempty" xml:"group_name,omitempty"`
	GroupRule *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule      `json:"group_rule,omitempty" xml:"group_rule,omitempty"`
	ItemList  []*ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupId   *int64                                                                  `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItem) SetGroupName(v string) *ProductOnlineListResponseDataProductsItemComboGroupsItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItem) SetGroupRule(v *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule) *ProductOnlineListResponseDataProductsItemComboGroupsItem {
	s.GroupRule = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItem) SetItemList(v []*ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) *ProductOnlineListResponseDataProductsItemComboGroupsItem {
	s.ItemList = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItem) SetGroupId(v int64) *ProductOnlineListResponseDataProductsItemComboGroupsItem {
	s.GroupId = &v
	return s
}

type ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule struct {
	OptionCount *int32 `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule) SetOptionCount(v int32) *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule {
	s.OptionCount = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule) SetTotalCount(v int32) *ProductOnlineListResponseDataProductsItemComboGroupsItemGroupRule {
	s.TotalCount = &v
	return s
}

type ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem struct {
	Count       *int64  `json:"count,omitempty" xml:"count,omitempty"`
	OutId       *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Price       *int64  `json:"price,omitempty" xml:"price,omitempty"`
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	Unit        *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetCount(v int64) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.Count = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetOutId(v string) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.OutId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetPrice(v int64) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetProductId(v int64) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetProductName(v string) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem) SetUnit(v string) *ProductOnlineListResponseDataProductsItemComboGroupsItemItemListItem {
	s.Unit = &v
	return s
}

type ProductOnlineListResponseDataProductsItemDishesImageListItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemDishesImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemDishesImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemDishesImageListItem) SetName(v string) *ProductOnlineListResponseDataProductsItemDishesImageListItem {
	s.Name = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemDishesImageListItem) SetUrl(v string) *ProductOnlineListResponseDataProductsItemDishesImageListItem {
	s.Url = &v
	return s
}

type ProductOnlineListResponseDataProductsItemEnvironmentImageListItem struct {
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemEnvironmentImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemEnvironmentImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemEnvironmentImageListItem) SetUrl(v string) *ProductOnlineListResponseDataProductsItemEnvironmentImageListItem {
	s.Url = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemEnvironmentImageListItem) SetName(v string) *ProductOnlineListResponseDataProductsItemEnvironmentImageListItem {
	s.Name = &v
	return s
}

type ProductOnlineListResponseDataProductsItemImageListItem struct {
	Url  *string `json:"url,omitempty" xml:"url,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemImageListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemImageListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemImageListItem) SetUrl(v string) *ProductOnlineListResponseDataProductsItemImageListItem {
	s.Url = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemImageListItem) SetName(v string) *ProductOnlineListResponseDataProductsItemImageListItem {
	s.Name = &v
	return s
}

type ProductOnlineListResponseDataProductsItemLimitBuyRule struct {
	RuleList []*ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem `json:"rule_list,omitempty" xml:"rule_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineListResponseDataProductsItemLimitBuyRule) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemLimitBuyRule) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemLimitBuyRule) SetRuleList(v []*ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) *ProductOnlineListResponseDataProductsItemLimitBuyRule {
	s.RuleList = v
	return s
}

type ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem struct {
	RangeType   *int   `json:"range_type,omitempty" xml:"range_type,omitempty" require:"true"`
	SubjectType *int   `json:"subject_type,omitempty" xml:"subject_type,omitempty" require:"true"`
	LimitNum    *int32 `json:"limit_num,omitempty" xml:"limit_num,omitempty" require:"true"`
}

func (s ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) SetRangeType(v int) *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.RangeType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) SetSubjectType(v int) *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.SubjectType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem) SetLimitNum(v int32) *ProductOnlineListResponseDataProductsItemLimitBuyRuleRuleListItem {
	s.LimitNum = &v
	return s
}

type ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem) SetName(v string) *ProductOnlineListResponseDataProductsItemPrimaryMaterialsItem {
	s.Name = &v
	return s
}

type ProductOnlineListResponseDataProductsItemProductExt struct {
	TestExtra *ProductOnlineListResponseDataProductsItemProductExtTestExtra `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemProductExt) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemProductExt) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemProductExt) SetTestExtra(v *ProductOnlineListResponseDataProductsItemProductExtTestExtra) *ProductOnlineListResponseDataProductsItemProductExt {
	s.TestExtra = v
	return s
}

type ProductOnlineListResponseDataProductsItemProductExtTestExtra struct {
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemProductExtTestExtra) SetUids(v []*string) *ProductOnlineListResponseDataProductsItemProductExtTestExtra {
	s.Uids = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemProductExtTestExtra) SetTestFlag(v bool) *ProductOnlineListResponseDataProductsItemProductExtTestExtra {
	s.TestFlag = &v
	return s
}

type ProductOnlineListResponseDataProductsItemProductSpecAttrsItem struct {
	GroupName *string                                                                      `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductOnlineListResponseDataProductsItemProductSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemProductSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItem) SetGroupName(v string) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItem {
	s.GroupName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItem) SetItemList(v []*ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItem {
	s.ItemList = v
	return s
}

type ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem struct {
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) SetWeight(v string) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) SetPrice(v int32) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) SetSpecName(v string) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem) SetUnit(v string) *ProductOnlineListResponseDataProductsItemProductSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

type ProductOnlineListResponseDataProductsItemSkusItem struct {
	SkuSpecAttrs []*ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem `json:"sku_spec_attrs,omitempty" xml:"sku_spec_attrs,omitempty" type:"Repeated"`
	Status       *int                                                                 `json:"status,omitempty" xml:"status,omitempty"`
	Stock        *ProductOnlineListResponseDataProductsItemSkusItemStock              `json:"stock,omitempty" xml:"stock,omitempty"`
	ActualAmount *int64                                                               `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64                                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	OutSkuId     *string                                                              `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId        *int64                                                               `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName      *string                                                              `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemSkusItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetSkuSpecAttrs(v []*ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.SkuSpecAttrs = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetStatus(v int) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.Status = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetStock(v *ProductOnlineListResponseDataProductsItemSkusItemStock) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.Stock = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetActualAmount(v int64) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetOriginAmount(v int64) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetOutSkuId(v string) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetSkuId(v int64) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItem) SetSkuName(v string) *ProductOnlineListResponseDataProductsItemSkusItem {
	s.SkuName = &v
	return s
}

type ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem struct {
	ItemList  []*ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	GroupName *string                                                                          `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem) SetItemList(v []*ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem {
	s.ItemList = v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem) SetGroupName(v string) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItem {
	s.GroupName = &v
	return s
}

type ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem struct {
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Weight   *string `json:"weight,omitempty" xml:"weight,omitempty"`
	Price    *int32  `json:"price,omitempty" xml:"price,omitempty"`
	SpecName *string `json:"spec_name,omitempty" xml:"spec_name,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetUnit(v string) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Unit = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetWeight(v string) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetPrice(v int32) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.Price = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem) SetSpecName(v string) *ProductOnlineListResponseDataProductsItemSkusItemSkuSpecAttrsItemItemListItem {
	s.SpecName = &v
	return s
}

type ProductOnlineListResponseDataProductsItemSkusItemStock struct {
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
}

func (s ProductOnlineListResponseDataProductsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseDataProductsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetLimitType(v int) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetSoldCount(v int64) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetSoldQty(v int64) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetStockQty(v int64) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetAvailQty(v int64) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *ProductOnlineListResponseDataProductsItemSkusItemStock) SetFrozenQty(v int64) *ProductOnlineListResponseDataProductsItemSkusItemStock {
	s.FrozenQty = &v
	return s
}

type ProductOnlineListResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ProductOnlineListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineListResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineListResponseExtra) SetDescription(v string) *ProductOnlineListResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductOnlineListResponseExtra) SetErrorCode(v int32) *ProductOnlineListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductOnlineListResponseExtra) SetLogid(v string) *ProductOnlineListResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductOnlineListResponseExtra) SetNow(v int64) *ProductOnlineListResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductOnlineListResponseExtra) SetSubDescription(v string) *ProductOnlineListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductOnlineListResponseExtra) SetSubErrorCode(v int32) *ProductOnlineListResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ProductOnlineQueryRequest struct {
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiIds           []*int64           `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
	Cursor           *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	ProductName      *string            `json:"product_name,omitempty" xml:"product_name,omitempty"`
	AccountId        *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	QueryAllPoi      *bool              `json:"query_all_poi,omitempty" xml:"query_all_poi,omitempty"`
	GoodsQueryType   *int               `json:"goods_query_type,omitempty" xml:"goods_query_type,omitempty"`
	Status           *int               `json:"status,omitempty" xml:"status,omitempty"`
	Count            *int64             `json:"count,omitempty" xml:"count,omitempty"`
	GoodsCreatorType *int               `json:"goods_creator_type,omitempty" xml:"goods_creator_type,omitempty"`
	ExtIds           []*int64           `json:"ext_ids,omitempty" xml:"ext_ids,omitempty" type:"Repeated"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ProductOnlineQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryRequest) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryRequest) SetAccessToken(v string) *ProductOnlineQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetPoiIds(v []*int64) *ProductOnlineQueryRequest {
	s.PoiIds = v
	return s
}

func (s *ProductOnlineQueryRequest) SetCursor(v string) *ProductOnlineQueryRequest {
	s.Cursor = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetProductName(v string) *ProductOnlineQueryRequest {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetAccountId(v string) *ProductOnlineQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetQueryAllPoi(v bool) *ProductOnlineQueryRequest {
	s.QueryAllPoi = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetGoodsQueryType(v int) *ProductOnlineQueryRequest {
	s.GoodsQueryType = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetStatus(v int) *ProductOnlineQueryRequest {
	s.Status = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetCount(v int64) *ProductOnlineQueryRequest {
	s.Count = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetGoodsCreatorType(v int) *ProductOnlineQueryRequest {
	s.GoodsCreatorType = &v
	return s
}

func (s *ProductOnlineQueryRequest) SetExtIds(v []*int64) *ProductOnlineQueryRequest {
	s.ExtIds = v
	return s
}

func (s *ProductOnlineQueryRequest) SetHeader(v map[string]*string) *ProductOnlineQueryRequest {
	s.Header = v
	return s
}

type ProductOnlineQueryResponse struct {
	Data     *ProductOnlineQueryResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductOnlineQueryResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *ProductOnlineQueryResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponse) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponse) SetData(v *ProductOnlineQueryResponseData) *ProductOnlineQueryResponse {
	s.Data = v
	return s
}

func (s *ProductOnlineQueryResponse) SetExtra(v *ProductOnlineQueryResponseExtra) *ProductOnlineQueryResponse {
	s.Extra = v
	return s
}

func (s *ProductOnlineQueryResponse) SetBaseResp(v *ProductOnlineQueryResponseBaseResp) *ProductOnlineQueryResponse {
	s.BaseResp = v
	return s
}

type ProductOnlineQueryResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ProductOnlineQueryResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseBaseResp) SetExtra(v map[string]*string) *ProductOnlineQueryResponseBaseResp {
	s.Extra = v
	return s
}

func (s *ProductOnlineQueryResponseBaseResp) SetStatusCode(v int32) *ProductOnlineQueryResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductOnlineQueryResponseBaseResp) SetStatusMessage(v string) *ProductOnlineQueryResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type ProductOnlineQueryResponseData struct {
	HasMore     *bool                                         `json:"has_more,omitempty" xml:"has_more,omitempty"`
	NextCursor  *string                                       `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	Products    []*ProductOnlineQueryResponseDataProductsItem `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	Description *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseData) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseData) SetHasMore(v bool) *ProductOnlineQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *ProductOnlineQueryResponseData) SetNextCursor(v string) *ProductOnlineQueryResponseData {
	s.NextCursor = &v
	return s
}

func (s *ProductOnlineQueryResponseData) SetProducts(v []*ProductOnlineQueryResponseDataProductsItem) *ProductOnlineQueryResponseData {
	s.Products = v
	return s
}

func (s *ProductOnlineQueryResponseData) SetDescription(v string) *ProductOnlineQueryResponseData {
	s.Description = &v
	return s
}

func (s *ProductOnlineQueryResponseData) SetErrorCode(v int32) *ProductOnlineQueryResponseData {
	s.ErrorCode = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItem struct {
	Product          *ProductOnlineQueryResponseDataProductsItemProduct                `json:"product,omitempty" xml:"product,omitempty"`
	Sku              *ProductOnlineQueryResponseDataProductsItemSku                    `json:"sku,omitempty" xml:"sku,omitempty"`
	Skus             []*ProductOnlineQueryResponseDataProductsItemSkusItem             `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	CommissionInfo   *ProductOnlineQueryResponseDataProductsItemCommissionInfo         `json:"commission_info,omitempty" xml:"commission_info,omitempty"`
	IsSellOut        *bool                                                             `json:"is_sell_out,omitempty" xml:"is_sell_out,omitempty"`
	OnlineStatus     *int                                                              `json:"online_status,omitempty" xml:"online_status,omitempty"`
	PoiSellOutStatus []*ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem `json:"poi_sell_out_status,omitempty" xml:"poi_sell_out_status,omitempty" type:"Repeated"`
}

func (s ProductOnlineQueryResponseDataProductsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetProduct(v *ProductOnlineQueryResponseDataProductsItemProduct) *ProductOnlineQueryResponseDataProductsItem {
	s.Product = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetSku(v *ProductOnlineQueryResponseDataProductsItemSku) *ProductOnlineQueryResponseDataProductsItem {
	s.Sku = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetSkus(v []*ProductOnlineQueryResponseDataProductsItemSkusItem) *ProductOnlineQueryResponseDataProductsItem {
	s.Skus = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetCommissionInfo(v *ProductOnlineQueryResponseDataProductsItemCommissionInfo) *ProductOnlineQueryResponseDataProductsItem {
	s.CommissionInfo = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetIsSellOut(v bool) *ProductOnlineQueryResponseDataProductsItem {
	s.IsSellOut = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetOnlineStatus(v int) *ProductOnlineQueryResponseDataProductsItem {
	s.OnlineStatus = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItem) SetPoiSellOutStatus(v []*ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) *ProductOnlineQueryResponseDataProductsItem {
	s.PoiSellOutStatus = v
	return s
}

type ProductOnlineQueryResponseDataProductsItemCommissionInfo struct {
	PlatformTakeRate *int64 `json:"platform_take_rate,omitempty" xml:"platform_take_rate,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemCommissionInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemCommissionInfo) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemCommissionInfo) SetPlatformTakeRate(v int64) *ProductOnlineQueryResponseDataProductsItemCommissionInfo {
	s.PlatformTakeRate = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem struct {
	SellOutEndTime   *int64  `json:"sell_out_end_time,omitempty" xml:"sell_out_end_time,omitempty"`
	SellOutStartTime *int64  `json:"sell_out_start_time,omitempty" xml:"sell_out_start_time,omitempty"`
	SellOutStatus    *int    `json:"sell_out_status,omitempty" xml:"sell_out_status,omitempty"`
	PoiId            *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) SetSellOutEndTime(v int64) *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem {
	s.SellOutEndTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) SetSellOutStartTime(v int64) *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem {
	s.SellOutStartTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) SetSellOutStatus(v int) *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem {
	s.SellOutStatus = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem) SetPoiId(v string) *ProductOnlineQueryResponseDataProductsItemPoiSellOutStatusItem {
	s.PoiId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemProduct struct {
	Desc             *string                                                      `json:"desc,omitempty" xml:"desc,omitempty"`
	BizLine          *int                                                         `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	OwnerAccountId   *int64                                                       `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	SpuId            *string                                                      `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	CreatorAccountId *int64                                                       `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	ProductName      *string                                                      `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	CreateTime       *int64                                                       `json:"create_time,omitempty" xml:"create_time,omitempty"`
	UpdateTime       *int64                                                       `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ContactName      *string                                                      `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	ProductSubType   *int                                                         `json:"product_sub_type,omitempty" xml:"product_sub_type,omitempty"`
	Version          *int64                                                       `json:"version,omitempty" xml:"version,omitempty"`
	CategoryId       *int64                                                       `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	OutUrl           *string                                                      `json:"out_url,omitempty" xml:"out_url,omitempty"`
	Telephone        []*string                                                    `json:"telephone,omitempty" xml:"telephone,omitempty" type:"Repeated"`
	ProductId        *string                                                      `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductType      *int                                                         `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	AccountName      *string                                                      `json:"account_name,omitempty" xml:"account_name,omitempty"`
	SoldStartTime    *int64                                                       `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	Extra            *string                                                      `json:"extra,omitempty" xml:"extra,omitempty"`
	AttrKeyValueMap  map[string]*string                                           `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	CategoryFullName *string                                                      `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	Pois             []*ProductOnlineQueryResponseDataProductsItemProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ProductExt       *ProductOnlineQueryResponseDataProductsItemProductProductExt `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	OpenBizType      *int                                                         `json:"open_biz_type,omitempty" xml:"open_biz_type,omitempty"`
	OutId            *string                                                      `json:"out_id,omitempty" xml:"out_id,omitempty"`
	SoldEndTime      *int64                                                       `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemProduct) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetDesc(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.Desc = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetBizLine(v int) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.BizLine = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetOwnerAccountId(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetSpuId(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.SpuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetCreatorAccountId(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetProductName(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ProductName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetCreateTime(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.CreateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetUpdateTime(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.UpdateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetContactName(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ContactName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetProductSubType(v int) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ProductSubType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetVersion(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.Version = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetCategoryId(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.CategoryId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetOutUrl(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.OutUrl = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetTelephone(v []*string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.Telephone = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetProductId(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ProductId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetProductType(v int) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ProductType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetAccountName(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.AccountName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetSoldStartTime(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.SoldStartTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetExtra(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.Extra = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetAttrKeyValueMap(v map[string]*string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetCategoryFullName(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.CategoryFullName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetPois(v []*ProductOnlineQueryResponseDataProductsItemProductPoisItem) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.Pois = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetProductExt(v *ProductOnlineQueryResponseDataProductsItemProductProductExt) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.ProductExt = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetOpenBizType(v int) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.OpenBizType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetOutId(v string) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.OutId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProduct) SetSoldEndTime(v int64) *ProductOnlineQueryResponseDataProductsItemProduct {
	s.SoldEndTime = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemProductPoisItem struct {
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemProductPoisItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemProductPoisItem) SetSupplierExtId(v string) *ProductOnlineQueryResponseDataProductsItemProductPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductPoisItem) SetSupplierId(v int64) *ProductOnlineQueryResponseDataProductsItemProductPoisItem {
	s.SupplierId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductPoisItem) SetPoiId(v string) *ProductOnlineQueryResponseDataProductsItemProductPoisItem {
	s.PoiId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemProductProductExt struct {
	TestExtra         *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra    `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
	ElemeExtraInfo    *string                                                                  `json:"eleme_extra_info,omitempty" xml:"eleme_extra_info,omitempty"`
	DisplayPrice      *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice `json:"display_price,omitempty" xml:"display_price,omitempty"`
	AllSkuSellOut     *bool                                                                    `json:"all_sku_sell_out,omitempty" xml:"all_sku_sell_out,omitempty"`
	CategoryOutId     *string                                                                  `json:"category_out_id,omitempty" xml:"category_out_id,omitempty"`
	AgencyRate        *int64                                                                   `json:"agency_rate,omitempty" xml:"agency_rate,omitempty"`
	AutoOnline        *bool                                                                    `json:"auto_online,omitempty" xml:"auto_online,omitempty"`
	IsBindClueElement *bool                                                                    `json:"is_bind_clue_element,omitempty" xml:"is_bind_clue_element,omitempty"`
	ElemeBizCode      *string                                                                  `json:"eleme_biz_code,omitempty" xml:"eleme_biz_code,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExt) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExt) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetTestExtra(v *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.TestExtra = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetElemeExtraInfo(v string) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.ElemeExtraInfo = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetDisplayPrice(v *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.DisplayPrice = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetAllSkuSellOut(v bool) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.AllSkuSellOut = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetCategoryOutId(v string) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.CategoryOutId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetAgencyRate(v int64) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.AgencyRate = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetAutoOnline(v bool) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.AutoOnline = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetIsBindClueElement(v bool) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.IsBindClueElement = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExt) SetElemeBizCode(v string) *ProductOnlineQueryResponseDataProductsItemProductProductExt {
	s.ElemeBizCode = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice struct {
	HighPrice *int64 `json:"high_price,omitempty" xml:"high_price,omitempty"`
	LowPrice  *int64 `json:"low_price,omitempty" xml:"low_price,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice) SetHighPrice(v int64) *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice {
	s.HighPrice = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice) SetLowPrice(v int64) *ProductOnlineQueryResponseDataProductsItemProductProductExtDisplayPrice {
	s.LowPrice = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra) SetTestFlag(v bool) *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra {
	s.TestFlag = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra) SetUids(v []*string) *ProductOnlineQueryResponseDataProductsItemProductProductExtTestExtra {
	s.Uids = v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSku struct {
	CreateTime      *int64                                               `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OriginAmount    *int64                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ActualAmount    *int64                                               `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	Status          *int                                                 `json:"status,omitempty" xml:"status,omitempty"`
	SkuExt          *ProductOnlineQueryResponseDataProductsItemSkuSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	SkuName         *string                                              `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	AttrKeyValueMap map[string]*string                                   `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	Stock           *ProductOnlineQueryResponseDataProductsItemSkuStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	OutSkuId        *string                                              `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	UpdateTime      *int64                                               `json:"update_time,omitempty" xml:"update_time,omitempty"`
	BindSkus        []*string                                            `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	Extra           *string                                              `json:"extra,omitempty" xml:"extra,omitempty"`
	SkuId           *string                                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSku) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSku) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetCreateTime(v int64) *ProductOnlineQueryResponseDataProductsItemSku {
	s.CreateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetOriginAmount(v int64) *ProductOnlineQueryResponseDataProductsItemSku {
	s.OriginAmount = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetActualAmount(v int64) *ProductOnlineQueryResponseDataProductsItemSku {
	s.ActualAmount = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetStatus(v int) *ProductOnlineQueryResponseDataProductsItemSku {
	s.Status = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetSkuExt(v *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) *ProductOnlineQueryResponseDataProductsItemSku {
	s.SkuExt = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetSkuName(v string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.SkuName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetAttrKeyValueMap(v map[string]*string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetStock(v *ProductOnlineQueryResponseDataProductsItemSkuStock) *ProductOnlineQueryResponseDataProductsItemSku {
	s.Stock = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetOutSkuId(v string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.OutSkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetUpdateTime(v int64) *ProductOnlineQueryResponseDataProductsItemSku {
	s.UpdateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetBindSkus(v []*string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.BindSkus = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetExtra(v string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.Extra = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSku) SetSkuId(v string) *ProductOnlineQueryResponseDataProductsItemSku {
	s.SkuId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkuSkuExt struct {
	UseSubRelStock      *bool                                                                   `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
	DiscountPromo       *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	RelRuleList         []*ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	AccountSettle       *bool                                                                   `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
	BindProductId       *int64                                                                  `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	BindSkus2c          map[int64][]*int64                                                      `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	BindSkuId           *int64                                                                  `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	CMspuId             *int64                                                                  `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	OriginStockQty      *int64                                                                  `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	SettleType          *int64                                                                  `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	LifeBizCode         *string                                                                 `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	BizId2cList         []*int64                                                                `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	OriSkus             map[int64][]*int64                                                      `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	TakeawayPresaleInfo *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExt) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExt) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetUseSubRelStock(v bool) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.UseSubRelStock = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetDiscountPromo(v *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetRelRuleList(v []*ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.RelRuleList = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetAccountSettle(v bool) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.AccountSettle = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetBindProductId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.BindProductId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetBindSkus2c(v map[int64][]*int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetBindSkuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetCMspuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.CMspuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetOriginStockQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetSettleType(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.SettleType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetLifeBizCode(v string) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetBizId2cList(v []*int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.BizId2cList = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetOriSkus(v map[int64][]*int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.OriSkus = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExt) SetTakeawayPresaleInfo(v *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) *ProductOnlineQueryResponseDataProductsItemSkuSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo struct {
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) SetBrandActivityId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) SetPlanId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo) SetPromoId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem struct {
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetCoefficient(v string) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetConstantVal(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetPriceRel(v bool) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetSharedQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetStockRel(v bool) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem) SetBizId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *ProductOnlineQueryResponseDataProductsItemSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkuStock struct {
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkuStock) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkuStock) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuStock) SetStockQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkuStock {
	s.StockQty = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkuStock) SetLimitType(v int) *ProductOnlineQueryResponseDataProductsItemSkuStock {
	s.LimitType = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItem struct {
	OriginAmount    *int64                                                    `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	UpdateTime      *int64                                                    `json:"update_time,omitempty" xml:"update_time,omitempty"`
	CreateTime      *int64                                                    `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Extra           *string                                                   `json:"extra,omitempty" xml:"extra,omitempty"`
	Stock           *ProductOnlineQueryResponseDataProductsItemSkusItemStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	SkuId           *string                                                   `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	AttrKeyValueMap map[string]*string                                        `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	SkuName         *string                                                   `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuExt          *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	BindSkus        []*string                                                 `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	Status          *int                                                      `json:"status,omitempty" xml:"status,omitempty"`
	ActualAmount    *int64                                                    `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OutSkuId        *string                                                   `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetOriginAmount(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetUpdateTime(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetCreateTime(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.CreateTime = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetExtra(v string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.Extra = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetStock(v *ProductOnlineQueryResponseDataProductsItemSkusItemStock) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.Stock = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetSkuId(v string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.SkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetAttrKeyValueMap(v map[string]*string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetSkuName(v string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.SkuName = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetSkuExt(v *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.SkuExt = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetBindSkus(v []*string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.BindSkus = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetStatus(v int) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.Status = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetActualAmount(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItem) SetOutSkuId(v string) *ProductOnlineQueryResponseDataProductsItemSkusItem {
	s.OutSkuId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt struct {
	AccountSettle       *bool                                                                        `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
	TakeawayPresaleInfo *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	CMspuId             *int64                                                                       `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	OriginStockQty      *int64                                                                       `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	BindProductId       *int64                                                                       `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	BindSkus2c          map[int64][]*int64                                                           `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	UseSubRelStock      *bool                                                                        `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
	BindSkuId           *int64                                                                       `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	RelRuleList         []*ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	DiscountPromo       *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	BizId2cList         []*int64                                                                     `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	SettleType          *int64                                                                       `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	LifeBizCode         *string                                                                      `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	OriSkus             map[int64][]*int64                                                           `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetAccountSettle(v bool) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.AccountSettle = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetTakeawayPresaleInfo(v *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetCMspuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.CMspuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetOriginStockQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetBindProductId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.BindProductId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetBindSkus2c(v map[int64][]*int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetUseSubRelStock(v bool) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.UseSubRelStock = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetBindSkuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetRelRuleList(v []*ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.RelRuleList = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetDiscountPromo(v *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetBizId2cList(v []*int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.BizId2cList = v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetSettleType(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.SettleType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetLifeBizCode(v string) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt) SetOriSkus(v map[int64][]*int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExt {
	s.OriSkus = v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo struct {
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) SetPromoId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) SetBrandActivityId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo) SetPlanId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem struct {
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetBizId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetCoefficient(v string) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetConstantVal(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetPriceRel(v bool) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetSharedQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem) SetStockRel(v bool) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

type ProductOnlineQueryResponseDataProductsItemSkusItemStock struct {
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseDataProductsItemSkusItemStock) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemStock) SetLimitType(v int) *ProductOnlineQueryResponseDataProductsItemSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *ProductOnlineQueryResponseDataProductsItemSkusItemStock) SetStockQty(v int64) *ProductOnlineQueryResponseDataProductsItemSkusItemStock {
	s.StockQty = &v
	return s
}

type ProductOnlineQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s ProductOnlineQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductOnlineQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductOnlineQueryResponseExtra) SetLogid(v string) *ProductOnlineQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductOnlineQueryResponseExtra) SetNow(v int64) *ProductOnlineQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductOnlineQueryResponseExtra) SetSubDescription(v string) *ProductOnlineQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductOnlineQueryResponseExtra) SetSubErrorCode(v int32) *ProductOnlineQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductOnlineQueryResponseExtra) SetDescription(v string) *ProductOnlineQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductOnlineQueryResponseExtra) SetErrorCode(v int32) *ProductOnlineQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type ProductSaveRequest struct {
	Base        *ProductSaveRequestBase    `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Product     *ProductSaveRequestProduct `json:"product,omitempty" xml:"product,omitempty" require:"true"`
	Sku         *ProductSaveRequestSku     `json:"sku,omitempty" xml:"sku,omitempty"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ProductSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequest) GoString() string {
	return s.String()
}

func (s *ProductSaveRequest) SetBase(v *ProductSaveRequestBase) *ProductSaveRequest {
	s.Base = v
	return s
}

func (s *ProductSaveRequest) SetAccountId(v string) *ProductSaveRequest {
	s.AccountId = &v
	return s
}

func (s *ProductSaveRequest) SetProduct(v *ProductSaveRequestProduct) *ProductSaveRequest {
	s.Product = v
	return s
}

func (s *ProductSaveRequest) SetSku(v *ProductSaveRequestSku) *ProductSaveRequest {
	s.Sku = v
	return s
}

func (s *ProductSaveRequest) SetHeader(v map[string]*string) *ProductSaveRequest {
	s.Header = v
	return s
}

func (s *ProductSaveRequest) SetAccessToken(v string) *ProductSaveRequest {
	s.AccessToken = &v
	return s
}

type ProductSaveRequestBase struct {
	Client     *string                           `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                           `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                           `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                           `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s ProductSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestBase) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestBase) SetClient(v string) *ProductSaveRequestBase {
	s.Client = &v
	return s
}

func (s *ProductSaveRequestBase) SetExtra(v map[string]*string) *ProductSaveRequestBase {
	s.Extra = v
	return s
}

func (s *ProductSaveRequestBase) SetLogID(v string) *ProductSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductSaveRequestBase) SetTrafficEnv(v *ProductSaveRequestBaseTrafficEnv) *ProductSaveRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductSaveRequestBase) SetAddr(v string) *ProductSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductSaveRequestBase) SetCaller(v string) *ProductSaveRequestBase {
	s.Caller = &v
	return s
}

type ProductSaveRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestBaseTrafficEnv) SetEnv(v string) *ProductSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductSaveRequestBaseTrafficEnv) SetOpen(v bool) *ProductSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductSaveRequestProduct struct {
	CreatorAccountId *int64                               `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	CategoryFullName *string                              `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	SoldStartTime    *int64                               `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	Extra            *string                              `json:"extra,omitempty" xml:"extra,omitempty"`
	Version          *int64                               `json:"version,omitempty" xml:"version,omitempty"`
	CategoryId       *int64                               `json:"category_id,omitempty" xml:"category_id,omitempty"`
	OwnerAccountId   *int64                               `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	AccountId        *string                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName      *string                              `json:"account_name,omitempty" xml:"account_name,omitempty"`
	CreateTime       *int64                               `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OutId            *string                              `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductExt       *ProductSaveRequestProductProductExt `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	UpdateTime       *int64                               `json:"update_time,omitempty" xml:"update_time,omitempty"`
	ProductType      *int                                 `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ProductName      *string                              `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductId        *string                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OutUrl           *string                              `json:"out_url,omitempty" xml:"out_url,omitempty"`
	BizLine          *int                                 `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	AttrKeyValueMap  map[string]*string                   `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	Pois             []*ProductSaveRequestProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	Desc             *string                              `json:"desc,omitempty" xml:"desc,omitempty"`
	SoldEndTime      *int64                               `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	ProductSpec      *int                                 `json:"product_spec,omitempty" xml:"product_spec,omitempty"`
}

func (s ProductSaveRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestProduct) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestProduct) SetCreatorAccountId(v int64) *ProductSaveRequestProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetCategoryFullName(v string) *ProductSaveRequestProduct {
	s.CategoryFullName = &v
	return s
}

func (s *ProductSaveRequestProduct) SetSoldStartTime(v int64) *ProductSaveRequestProduct {
	s.SoldStartTime = &v
	return s
}

func (s *ProductSaveRequestProduct) SetExtra(v string) *ProductSaveRequestProduct {
	s.Extra = &v
	return s
}

func (s *ProductSaveRequestProduct) SetVersion(v int64) *ProductSaveRequestProduct {
	s.Version = &v
	return s
}

func (s *ProductSaveRequestProduct) SetCategoryId(v int64) *ProductSaveRequestProduct {
	s.CategoryId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetOwnerAccountId(v int64) *ProductSaveRequestProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetAccountId(v string) *ProductSaveRequestProduct {
	s.AccountId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetAccountName(v string) *ProductSaveRequestProduct {
	s.AccountName = &v
	return s
}

func (s *ProductSaveRequestProduct) SetCreateTime(v int64) *ProductSaveRequestProduct {
	s.CreateTime = &v
	return s
}

func (s *ProductSaveRequestProduct) SetOutId(v string) *ProductSaveRequestProduct {
	s.OutId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetProductExt(v *ProductSaveRequestProductProductExt) *ProductSaveRequestProduct {
	s.ProductExt = v
	return s
}

func (s *ProductSaveRequestProduct) SetUpdateTime(v int64) *ProductSaveRequestProduct {
	s.UpdateTime = &v
	return s
}

func (s *ProductSaveRequestProduct) SetProductType(v int) *ProductSaveRequestProduct {
	s.ProductType = &v
	return s
}

func (s *ProductSaveRequestProduct) SetProductName(v string) *ProductSaveRequestProduct {
	s.ProductName = &v
	return s
}

func (s *ProductSaveRequestProduct) SetProductId(v string) *ProductSaveRequestProduct {
	s.ProductId = &v
	return s
}

func (s *ProductSaveRequestProduct) SetOutUrl(v string) *ProductSaveRequestProduct {
	s.OutUrl = &v
	return s
}

func (s *ProductSaveRequestProduct) SetBizLine(v int) *ProductSaveRequestProduct {
	s.BizLine = &v
	return s
}

func (s *ProductSaveRequestProduct) SetAttrKeyValueMap(v map[string]*string) *ProductSaveRequestProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductSaveRequestProduct) SetPois(v []*ProductSaveRequestProductPoisItem) *ProductSaveRequestProduct {
	s.Pois = v
	return s
}

func (s *ProductSaveRequestProduct) SetDesc(v string) *ProductSaveRequestProduct {
	s.Desc = &v
	return s
}

func (s *ProductSaveRequestProduct) SetSoldEndTime(v int64) *ProductSaveRequestProduct {
	s.SoldEndTime = &v
	return s
}

func (s *ProductSaveRequestProduct) SetProductSpec(v int) *ProductSaveRequestProduct {
	s.ProductSpec = &v
	return s
}

type ProductSaveRequestProductPoisItem struct {
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ProductSaveRequestProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestProductPoisItem) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestProductPoisItem) SetSupplierExtId(v string) *ProductSaveRequestProductPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *ProductSaveRequestProductPoisItem) SetSupplierId(v int64) *ProductSaveRequestProductPoisItem {
	s.SupplierId = &v
	return s
}

func (s *ProductSaveRequestProductPoisItem) SetPoiId(v string) *ProductSaveRequestProductPoisItem {
	s.PoiId = &v
	return s
}

type ProductSaveRequestProductProductExt struct {
	TestExtra *ProductSaveRequestProductProductExtTestExtra `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
}

func (s ProductSaveRequestProductProductExt) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestProductProductExt) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestProductProductExt) SetTestExtra(v *ProductSaveRequestProductProductExtTestExtra) *ProductSaveRequestProductProductExt {
	s.TestExtra = v
	return s
}

type ProductSaveRequestProductProductExtTestExtra struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
}

func (s ProductSaveRequestProductProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestProductProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestProductProductExtTestExtra) SetTestFlag(v bool) *ProductSaveRequestProductProductExtTestExtra {
	s.TestFlag = &v
	return s
}

func (s *ProductSaveRequestProductProductExtTestExtra) SetUids(v []*string) *ProductSaveRequestProductProductExtTestExtra {
	s.Uids = v
	return s
}

type ProductSaveRequestSku struct {
	AttrKeyValueMap map[string]*string                `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	Status          *int                              `json:"status,omitempty" xml:"status,omitempty"`
	BindSkus        []*string                         `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	OriginAmount    *int64                            `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	CreateTime      *int64                            `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Stock           *ProductSaveRequestSkuStock       `json:"stock,omitempty" xml:"stock,omitempty"`
	SkuId           *string                           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OutSkuId        *string                           `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Extra           *string                           `json:"extra,omitempty" xml:"extra,omitempty"`
	Specs           []*ProductSaveRequestSkuSpecsItem `json:"specs,omitempty" xml:"specs,omitempty" type:"Repeated"`
	SkuName         *string                           `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	ActualAmount    *int64                            `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	UpdateTime      *int64                            `json:"update_time,omitempty" xml:"update_time,omitempty"`
}

func (s ProductSaveRequestSku) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestSku) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestSku) SetAttrKeyValueMap(v map[string]*string) *ProductSaveRequestSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *ProductSaveRequestSku) SetStatus(v int) *ProductSaveRequestSku {
	s.Status = &v
	return s
}

func (s *ProductSaveRequestSku) SetBindSkus(v []*string) *ProductSaveRequestSku {
	s.BindSkus = v
	return s
}

func (s *ProductSaveRequestSku) SetOriginAmount(v int64) *ProductSaveRequestSku {
	s.OriginAmount = &v
	return s
}

func (s *ProductSaveRequestSku) SetCreateTime(v int64) *ProductSaveRequestSku {
	s.CreateTime = &v
	return s
}

func (s *ProductSaveRequestSku) SetStock(v *ProductSaveRequestSkuStock) *ProductSaveRequestSku {
	s.Stock = v
	return s
}

func (s *ProductSaveRequestSku) SetSkuId(v string) *ProductSaveRequestSku {
	s.SkuId = &v
	return s
}

func (s *ProductSaveRequestSku) SetOutSkuId(v string) *ProductSaveRequestSku {
	s.OutSkuId = &v
	return s
}

func (s *ProductSaveRequestSku) SetExtra(v string) *ProductSaveRequestSku {
	s.Extra = &v
	return s
}

func (s *ProductSaveRequestSku) SetSpecs(v []*ProductSaveRequestSkuSpecsItem) *ProductSaveRequestSku {
	s.Specs = v
	return s
}

func (s *ProductSaveRequestSku) SetSkuName(v string) *ProductSaveRequestSku {
	s.SkuName = &v
	return s
}

func (s *ProductSaveRequestSku) SetActualAmount(v int64) *ProductSaveRequestSku {
	s.ActualAmount = &v
	return s
}

func (s *ProductSaveRequestSku) SetUpdateTime(v int64) *ProductSaveRequestSku {
	s.UpdateTime = &v
	return s
}

type ProductSaveRequestSkuSpecsItem struct {
	GroupName *string                                       `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ItemList  []*ProductSaveRequestSkuSpecsItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s ProductSaveRequestSkuSpecsItem) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestSkuSpecsItem) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestSkuSpecsItem) SetGroupName(v string) *ProductSaveRequestSkuSpecsItem {
	s.GroupName = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItem) SetItemList(v []*ProductSaveRequestSkuSpecsItemItemListItem) *ProductSaveRequestSkuSpecsItem {
	s.ItemList = v
	return s
}

type ProductSaveRequestSkuSpecsItemItemListItem struct {
	Weight    *string `json:"weight,omitempty" xml:"weight,omitempty"`
	AddPrice  *int32  `json:"add_price,omitempty" xml:"add_price,omitempty"`
	IsDefault *bool   `json:"is_default,omitempty" xml:"is_default,omitempty"`
	ItemName  *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	OutSkuId  *string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	SkuId     *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Unit      *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ProductSaveRequestSkuSpecsItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestSkuSpecsItemItemListItem) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetWeight(v string) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.Weight = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetAddPrice(v int32) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.AddPrice = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetIsDefault(v bool) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.IsDefault = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetItemName(v string) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.ItemName = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetOutSkuId(v string) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.OutSkuId = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetSkuId(v string) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.SkuId = &v
	return s
}

func (s *ProductSaveRequestSkuSpecsItemItemListItem) SetUnit(v string) *ProductSaveRequestSkuSpecsItemItemListItem {
	s.Unit = &v
	return s
}

type ProductSaveRequestSkuStock struct {
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
}

func (s ProductSaveRequestSkuStock) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveRequestSkuStock) GoString() string {
	return s.String()
}

func (s *ProductSaveRequestSkuStock) SetLimitType(v int) *ProductSaveRequestSkuStock {
	s.LimitType = &v
	return s
}

func (s *ProductSaveRequestSkuStock) SetSoldCount(v int64) *ProductSaveRequestSkuStock {
	s.SoldCount = &v
	return s
}

func (s *ProductSaveRequestSkuStock) SetSoldQty(v int64) *ProductSaveRequestSkuStock {
	s.SoldQty = &v
	return s
}

func (s *ProductSaveRequestSkuStock) SetStockQty(v int64) *ProductSaveRequestSkuStock {
	s.StockQty = &v
	return s
}

func (s *ProductSaveRequestSkuStock) SetAvailQty(v int64) *ProductSaveRequestSkuStock {
	s.AvailQty = &v
	return s
}

func (s *ProductSaveRequestSkuStock) SetFrozenQty(v int64) *ProductSaveRequestSkuStock {
	s.FrozenQty = &v
	return s
}

type ProductSaveResponse struct {
	BaseResp *ProductSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ProductSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveResponse) GoString() string {
	return s.String()
}

func (s *ProductSaveResponse) SetBaseResp(v *ProductSaveResponseBaseResp) *ProductSaveResponse {
	s.BaseResp = v
	return s
}

func (s *ProductSaveResponse) SetData(v *ProductSaveResponseData) *ProductSaveResponse {
	s.Data = v
	return s
}

func (s *ProductSaveResponse) SetExtra(v *ProductSaveResponseExtra) *ProductSaveResponse {
	s.Extra = v
	return s
}

type ProductSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s ProductSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductSaveResponseBaseResp) SetStatusCode(v int32) *ProductSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductSaveResponseBaseResp) SetStatusMessage(v string) *ProductSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductSaveResponseBaseResp) SetExtra(v map[string]*string) *ProductSaveResponseBaseResp {
	s.Extra = v
	return s
}

type ProductSaveResponseData struct {
	SkuIds      []*string `json:"sku_ids,omitempty" xml:"sku_ids,omitempty" type:"Repeated"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductId   *string   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuId       *string   `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s ProductSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveResponseData) GoString() string {
	return s.String()
}

func (s *ProductSaveResponseData) SetSkuIds(v []*string) *ProductSaveResponseData {
	s.SkuIds = v
	return s
}

func (s *ProductSaveResponseData) SetDescription(v string) *ProductSaveResponseData {
	s.Description = &v
	return s
}

func (s *ProductSaveResponseData) SetErrorCode(v int32) *ProductSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductSaveResponseData) SetProductId(v string) *ProductSaveResponseData {
	s.ProductId = &v
	return s
}

func (s *ProductSaveResponseData) SetSkuId(v string) *ProductSaveResponseData {
	s.SkuId = &v
	return s
}

type ProductSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ProductSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductSaveResponseExtra) SetDescription(v string) *ProductSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductSaveResponseExtra) SetErrorCode(v int32) *ProductSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductSaveResponseExtra) SetLogid(v string) *ProductSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductSaveResponseExtra) SetNow(v int64) *ProductSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductSaveResponseExtra) SetSubDescription(v string) *ProductSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductSaveResponseExtra) SetSubErrorCode(v int32) *ProductSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ProductUpdatePriceRequest struct {
	Header       map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductId    *string                        `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Base         *ProductUpdatePriceRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId    *string                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ActualAmount *int64                         `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s ProductUpdatePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceRequest) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceRequest) SetHeader(v map[string]*string) *ProductUpdatePriceRequest {
	s.Header = v
	return s
}

func (s *ProductUpdatePriceRequest) SetAccessToken(v string) *ProductUpdatePriceRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductUpdatePriceRequest) SetProductId(v string) *ProductUpdatePriceRequest {
	s.ProductId = &v
	return s
}

func (s *ProductUpdatePriceRequest) SetBase(v *ProductUpdatePriceRequestBase) *ProductUpdatePriceRequest {
	s.Base = v
	return s
}

func (s *ProductUpdatePriceRequest) SetAccountId(v string) *ProductUpdatePriceRequest {
	s.AccountId = &v
	return s
}

func (s *ProductUpdatePriceRequest) SetActualAmount(v int64) *ProductUpdatePriceRequest {
	s.ActualAmount = &v
	return s
}

type ProductUpdatePriceRequestBase struct {
	Addr       *string                                  `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                  `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                  `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                       `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                  `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductUpdatePriceRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s ProductUpdatePriceRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceRequestBase) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceRequestBase) SetAddr(v string) *ProductUpdatePriceRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductUpdatePriceRequestBase) SetCaller(v string) *ProductUpdatePriceRequestBase {
	s.Caller = &v
	return s
}

func (s *ProductUpdatePriceRequestBase) SetClient(v string) *ProductUpdatePriceRequestBase {
	s.Client = &v
	return s
}

func (s *ProductUpdatePriceRequestBase) SetExtra(v map[string]*string) *ProductUpdatePriceRequestBase {
	s.Extra = v
	return s
}

func (s *ProductUpdatePriceRequestBase) SetLogID(v string) *ProductUpdatePriceRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductUpdatePriceRequestBase) SetTrafficEnv(v *ProductUpdatePriceRequestBaseTrafficEnv) *ProductUpdatePriceRequestBase {
	s.TrafficEnv = v
	return s
}

type ProductUpdatePriceRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductUpdatePriceRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceRequestBaseTrafficEnv) SetEnv(v string) *ProductUpdatePriceRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductUpdatePriceRequestBaseTrafficEnv) SetOpen(v bool) *ProductUpdatePriceRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductUpdatePriceResponse struct {
	BaseResp *ProductUpdatePriceResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductUpdatePriceResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductUpdatePriceResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ProductUpdatePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceResponse) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceResponse) SetBaseResp(v *ProductUpdatePriceResponseBaseResp) *ProductUpdatePriceResponse {
	s.BaseResp = v
	return s
}

func (s *ProductUpdatePriceResponse) SetData(v *ProductUpdatePriceResponseData) *ProductUpdatePriceResponse {
	s.Data = v
	return s
}

func (s *ProductUpdatePriceResponse) SetExtra(v *ProductUpdatePriceResponseExtra) *ProductUpdatePriceResponse {
	s.Extra = v
	return s
}

type ProductUpdatePriceResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ProductUpdatePriceResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceResponseBaseResp) SetExtra(v map[string]*string) *ProductUpdatePriceResponseBaseResp {
	s.Extra = v
	return s
}

func (s *ProductUpdatePriceResponseBaseResp) SetStatusCode(v int32) *ProductUpdatePriceResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductUpdatePriceResponseBaseResp) SetStatusMessage(v string) *ProductUpdatePriceResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type ProductUpdatePriceResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductUpdatePriceResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceResponseData) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceResponseData) SetErrorCode(v int32) *ProductUpdatePriceResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdatePriceResponseData) SetDescription(v string) *ProductUpdatePriceResponseData {
	s.Description = &v
	return s
}

type ProductUpdatePriceResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s ProductUpdatePriceResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdatePriceResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductUpdatePriceResponseExtra) SetNow(v int64) *ProductUpdatePriceResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductUpdatePriceResponseExtra) SetSubDescription(v string) *ProductUpdatePriceResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductUpdatePriceResponseExtra) SetSubErrorCode(v int32) *ProductUpdatePriceResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductUpdatePriceResponseExtra) SetDescription(v string) *ProductUpdatePriceResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductUpdatePriceResponseExtra) SetErrorCode(v int32) *ProductUpdatePriceResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdatePriceResponseExtra) SetLogid(v string) *ProductUpdatePriceResponseExtra {
	s.Logid = &v
	return s
}

type ProductUpdateStatusRequest struct {
	OperateType *int                            `json:"operate_type,omitempty" xml:"operate_type,omitempty" require:"true"`
	Header      map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductIds  []*string                       `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Base        *ProductUpdateStatusRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                         `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s ProductUpdateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusRequest) SetOperateType(v int) *ProductUpdateStatusRequest {
	s.OperateType = &v
	return s
}

func (s *ProductUpdateStatusRequest) SetHeader(v map[string]*string) *ProductUpdateStatusRequest {
	s.Header = v
	return s
}

func (s *ProductUpdateStatusRequest) SetAccessToken(v string) *ProductUpdateStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *ProductUpdateStatusRequest) SetProductIds(v []*string) *ProductUpdateStatusRequest {
	s.ProductIds = v
	return s
}

func (s *ProductUpdateStatusRequest) SetBase(v *ProductUpdateStatusRequestBase) *ProductUpdateStatusRequest {
	s.Base = v
	return s
}

func (s *ProductUpdateStatusRequest) SetAccountId(v string) *ProductUpdateStatusRequest {
	s.AccountId = &v
	return s
}

type ProductUpdateStatusRequestBase struct {
	Client     *string                                   `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                        `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                   `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductUpdateStatusRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                   `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                   `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s ProductUpdateStatusRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusRequestBase) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusRequestBase) SetClient(v string) *ProductUpdateStatusRequestBase {
	s.Client = &v
	return s
}

func (s *ProductUpdateStatusRequestBase) SetExtra(v map[string]*string) *ProductUpdateStatusRequestBase {
	s.Extra = v
	return s
}

func (s *ProductUpdateStatusRequestBase) SetLogID(v string) *ProductUpdateStatusRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductUpdateStatusRequestBase) SetTrafficEnv(v *ProductUpdateStatusRequestBaseTrafficEnv) *ProductUpdateStatusRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductUpdateStatusRequestBase) SetAddr(v string) *ProductUpdateStatusRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductUpdateStatusRequestBase) SetCaller(v string) *ProductUpdateStatusRequestBase {
	s.Caller = &v
	return s
}

type ProductUpdateStatusRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductUpdateStatusRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusRequestBaseTrafficEnv) SetEnv(v string) *ProductUpdateStatusRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductUpdateStatusRequestBaseTrafficEnv) SetOpen(v bool) *ProductUpdateStatusRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductUpdateStatusResponse struct {
	BaseResp *ProductUpdateStatusResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductUpdateStatusResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *ProductUpdateStatusResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s ProductUpdateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusResponse) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusResponse) SetBaseResp(v *ProductUpdateStatusResponseBaseResp) *ProductUpdateStatusResponse {
	s.BaseResp = v
	return s
}

func (s *ProductUpdateStatusResponse) SetData(v *ProductUpdateStatusResponseData) *ProductUpdateStatusResponse {
	s.Data = v
	return s
}

func (s *ProductUpdateStatusResponse) SetExtra(v *ProductUpdateStatusResponseExtra) *ProductUpdateStatusResponse {
	s.Extra = v
	return s
}

type ProductUpdateStatusResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ProductUpdateStatusResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusResponseBaseResp) SetStatusMessage(v string) *ProductUpdateStatusResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *ProductUpdateStatusResponseBaseResp) SetExtra(v map[string]*string) *ProductUpdateStatusResponseBaseResp {
	s.Extra = v
	return s
}

func (s *ProductUpdateStatusResponseBaseResp) SetStatusCode(v int32) *ProductUpdateStatusResponseBaseResp {
	s.StatusCode = &v
	return s
}

type ProductUpdateStatusResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductUpdateStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusResponseData) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusResponseData) SetErrorCode(v int32) *ProductUpdateStatusResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdateStatusResponseData) SetDescription(v string) *ProductUpdateStatusResponseData {
	s.Description = &v
	return s
}

type ProductUpdateStatusResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ProductUpdateStatusResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStatusResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductUpdateStatusResponseExtra) SetSubDescription(v string) *ProductUpdateStatusResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ProductUpdateStatusResponseExtra) SetSubErrorCode(v int32) *ProductUpdateStatusResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductUpdateStatusResponseExtra) SetDescription(v string) *ProductUpdateStatusResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductUpdateStatusResponseExtra) SetErrorCode(v int32) *ProductUpdateStatusResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdateStatusResponseExtra) SetLogid(v string) *ProductUpdateStatusResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductUpdateStatusResponseExtra) SetNow(v int64) *ProductUpdateStatusResponseExtra {
	s.Now = &v
	return s
}

type ProductUpdateStockRequest struct {
	SkuId       *string                              `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	StockQty    *int64                               `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	LimitType   *int                                 `json:"limit_type,omitempty" xml:"limit_type,omitempty"`
	OutSkuId    *string                              `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	Pois        []*ProductUpdateStockRequestPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	ProductId   *string                              `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Header      map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	Base        *ProductUpdateStockRequestBase       `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccessToken *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ProductUpdateStockRequest) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockRequest) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockRequest) SetSkuId(v string) *ProductUpdateStockRequest {
	s.SkuId = &v
	return s
}

func (s *ProductUpdateStockRequest) SetStockQty(v int64) *ProductUpdateStockRequest {
	s.StockQty = &v
	return s
}

func (s *ProductUpdateStockRequest) SetLimitType(v int) *ProductUpdateStockRequest {
	s.LimitType = &v
	return s
}

func (s *ProductUpdateStockRequest) SetOutSkuId(v string) *ProductUpdateStockRequest {
	s.OutSkuId = &v
	return s
}

func (s *ProductUpdateStockRequest) SetPois(v []*ProductUpdateStockRequestPoisItem) *ProductUpdateStockRequest {
	s.Pois = v
	return s
}

func (s *ProductUpdateStockRequest) SetProductId(v string) *ProductUpdateStockRequest {
	s.ProductId = &v
	return s
}

func (s *ProductUpdateStockRequest) SetHeader(v map[string]*string) *ProductUpdateStockRequest {
	s.Header = v
	return s
}

func (s *ProductUpdateStockRequest) SetBase(v *ProductUpdateStockRequestBase) *ProductUpdateStockRequest {
	s.Base = v
	return s
}

func (s *ProductUpdateStockRequest) SetAccountId(v string) *ProductUpdateStockRequest {
	s.AccountId = &v
	return s
}

func (s *ProductUpdateStockRequest) SetAccessToken(v string) *ProductUpdateStockRequest {
	s.AccessToken = &v
	return s
}

type ProductUpdateStockRequestBase struct {
	Client     *string                                  `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                       `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                  `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *ProductUpdateStockRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                  `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                  `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s ProductUpdateStockRequestBase) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockRequestBase) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockRequestBase) SetClient(v string) *ProductUpdateStockRequestBase {
	s.Client = &v
	return s
}

func (s *ProductUpdateStockRequestBase) SetExtra(v map[string]*string) *ProductUpdateStockRequestBase {
	s.Extra = v
	return s
}

func (s *ProductUpdateStockRequestBase) SetLogID(v string) *ProductUpdateStockRequestBase {
	s.LogID = &v
	return s
}

func (s *ProductUpdateStockRequestBase) SetTrafficEnv(v *ProductUpdateStockRequestBaseTrafficEnv) *ProductUpdateStockRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *ProductUpdateStockRequestBase) SetAddr(v string) *ProductUpdateStockRequestBase {
	s.Addr = &v
	return s
}

func (s *ProductUpdateStockRequestBase) SetCaller(v string) *ProductUpdateStockRequestBase {
	s.Caller = &v
	return s
}

type ProductUpdateStockRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s ProductUpdateStockRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockRequestBaseTrafficEnv) SetEnv(v string) *ProductUpdateStockRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *ProductUpdateStockRequestBaseTrafficEnv) SetOpen(v bool) *ProductUpdateStockRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type ProductUpdateStockRequestPoisItem struct {
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
}

func (s ProductUpdateStockRequestPoisItem) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockRequestPoisItem) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockRequestPoisItem) SetSupplierId(v int64) *ProductUpdateStockRequestPoisItem {
	s.SupplierId = &v
	return s
}

func (s *ProductUpdateStockRequestPoisItem) SetPoiId(v string) *ProductUpdateStockRequestPoisItem {
	s.PoiId = &v
	return s
}

func (s *ProductUpdateStockRequestPoisItem) SetSupplierExtId(v string) *ProductUpdateStockRequestPoisItem {
	s.SupplierExtId = &v
	return s
}

type ProductUpdateStockResponse struct {
	Extra    *ProductUpdateStockResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *ProductUpdateStockResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *ProductUpdateStockResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s ProductUpdateStockResponse) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockResponse) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockResponse) SetExtra(v *ProductUpdateStockResponseExtra) *ProductUpdateStockResponse {
	s.Extra = v
	return s
}

func (s *ProductUpdateStockResponse) SetBaseResp(v *ProductUpdateStockResponseBaseResp) *ProductUpdateStockResponse {
	s.BaseResp = v
	return s
}

func (s *ProductUpdateStockResponse) SetData(v *ProductUpdateStockResponseData) *ProductUpdateStockResponse {
	s.Data = v
	return s
}

type ProductUpdateStockResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ProductUpdateStockResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockResponseBaseResp) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockResponseBaseResp) SetExtra(v map[string]*string) *ProductUpdateStockResponseBaseResp {
	s.Extra = v
	return s
}

func (s *ProductUpdateStockResponseBaseResp) SetStatusCode(v int32) *ProductUpdateStockResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *ProductUpdateStockResponseBaseResp) SetStatusMessage(v string) *ProductUpdateStockResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type ProductUpdateStockResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ProductUpdateStockResponseData) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockResponseData) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockResponseData) SetErrorCode(v int32) *ProductUpdateStockResponseData {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdateStockResponseData) SetDescription(v string) *ProductUpdateStockResponseData {
	s.Description = &v
	return s
}

type ProductUpdateStockResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s ProductUpdateStockResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ProductUpdateStockResponseExtra) GoString() string {
	return s.String()
}

func (s *ProductUpdateStockResponseExtra) SetSubErrorCode(v int32) *ProductUpdateStockResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ProductUpdateStockResponseExtra) SetDescription(v string) *ProductUpdateStockResponseExtra {
	s.Description = &v
	return s
}

func (s *ProductUpdateStockResponseExtra) SetErrorCode(v int32) *ProductUpdateStockResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ProductUpdateStockResponseExtra) SetLogid(v string) *ProductUpdateStockResponseExtra {
	s.Logid = &v
	return s
}

func (s *ProductUpdateStockResponseExtra) SetNow(v int64) *ProductUpdateStockResponseExtra {
	s.Now = &v
	return s
}

func (s *ProductUpdateStockResponseExtra) SetSubDescription(v string) *ProductUpdateStockResponseExtra {
	s.SubDescription = &v
	return s
}

type PromotionExitRequest struct {
	AccountId    *string                           `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header       map[string]*string                `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                           `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PromotionId  *string                           `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	QuitResource *PromotionExitRequestQuitResource `json:"quit_resource,omitempty" xml:"quit_resource,omitempty" require:"true"`
}

func (s PromotionExitRequest) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitRequest) GoString() string {
	return s.String()
}

func (s *PromotionExitRequest) SetAccountId(v string) *PromotionExitRequest {
	s.AccountId = &v
	return s
}

func (s *PromotionExitRequest) SetHeader(v map[string]*string) *PromotionExitRequest {
	s.Header = v
	return s
}

func (s *PromotionExitRequest) SetAccessToken(v string) *PromotionExitRequest {
	s.AccessToken = &v
	return s
}

func (s *PromotionExitRequest) SetPromotionId(v string) *PromotionExitRequest {
	s.PromotionId = &v
	return s
}

func (s *PromotionExitRequest) SetQuitResource(v *PromotionExitRequestQuitResource) *PromotionExitRequest {
	s.QuitResource = v
	return s
}

type PromotionExitRequestQuitResource struct {
	Resources []*PromotionExitRequestQuitResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s PromotionExitRequestQuitResource) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitRequestQuitResource) GoString() string {
	return s.String()
}

func (s *PromotionExitRequestQuitResource) SetResources(v []*PromotionExitRequestQuitResourceResourcesItem) *PromotionExitRequestQuitResource {
	s.Resources = v
	return s
}

type PromotionExitRequestQuitResourceResourcesItem struct {
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
}

func (s PromotionExitRequestQuitResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitRequestQuitResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *PromotionExitRequestQuitResourceResourcesItem) SetRatePlanId(v string) *PromotionExitRequestQuitResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

func (s *PromotionExitRequestQuitResourceResourcesItem) SetRatePlanStatus(v int32) *PromotionExitRequestQuitResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

type PromotionExitResponse struct {
	Data  *PromotionExitResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PromotionExitResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PromotionExitResponse) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitResponse) GoString() string {
	return s.String()
}

func (s *PromotionExitResponse) SetData(v *PromotionExitResponseData) *PromotionExitResponse {
	s.Data = v
	return s
}

func (s *PromotionExitResponse) SetExtra(v *PromotionExitResponseExtra) *PromotionExitResponse {
	s.Extra = v
	return s
}

type PromotionExitResponseData struct {
	MessageDetail []*PromotionExitResponseDataMessageDetailItem `json:"message_detail,omitempty" xml:"message_detail,omitempty" type:"Repeated"`
	Status        *int                                          `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PromotionExitResponseData) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitResponseData) GoString() string {
	return s.String()
}

func (s *PromotionExitResponseData) SetMessageDetail(v []*PromotionExitResponseDataMessageDetailItem) *PromotionExitResponseData {
	s.MessageDetail = v
	return s
}

func (s *PromotionExitResponseData) SetStatus(v int) *PromotionExitResponseData {
	s.Status = &v
	return s
}

func (s *PromotionExitResponseData) SetGwErrorCode(v int32) *PromotionExitResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PromotionExitResponseData) SetGwDescription(v string) *PromotionExitResponseData {
	s.GwDescription = &v
	return s
}

type PromotionExitResponseDataMessageDetailItem struct {
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	PromotionId  *string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
}

func (s PromotionExitResponseDataMessageDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitResponseDataMessageDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionExitResponseDataMessageDetailItem) SetRatePlanId(v string) *PromotionExitResponseDataMessageDetailItem {
	s.RatePlanId = &v
	return s
}

func (s *PromotionExitResponseDataMessageDetailItem) SetErrorMessage(v string) *PromotionExitResponseDataMessageDetailItem {
	s.ErrorMessage = &v
	return s
}

func (s *PromotionExitResponseDataMessageDetailItem) SetPromotionId(v string) *PromotionExitResponseDataMessageDetailItem {
	s.PromotionId = &v
	return s
}

type PromotionExitResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s PromotionExitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PromotionExitResponseExtra) GoString() string {
	return s.String()
}

func (s *PromotionExitResponseExtra) SetNow(v int64) *PromotionExitResponseExtra {
	s.Now = &v
	return s
}

func (s *PromotionExitResponseExtra) SetSubDescription(v string) *PromotionExitResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PromotionExitResponseExtra) SetSubErrorCode(v int32) *PromotionExitResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PromotionExitResponseExtra) SetDescription(v string) *PromotionExitResponseExtra {
	s.Description = &v
	return s
}

func (s *PromotionExitResponseExtra) SetErrorCode(v int32) *PromotionExitResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PromotionExitResponseExtra) SetLogid(v string) *PromotionExitResponseExtra {
	s.Logid = &v
	return s
}

type PromotionInfoQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Page        *int32             `json:"page,omitempty" xml:"page,omitempty"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PromotionId []*string          `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	HotelId     []*string          `json:"hotel_id,omitempty" xml:"hotel_id,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s PromotionInfoQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryRequest) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryRequest) SetAccessToken(v string) *PromotionInfoQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *PromotionInfoQueryRequest) SetPage(v int32) *PromotionInfoQueryRequest {
	s.Page = &v
	return s
}

func (s *PromotionInfoQueryRequest) SetPageSize(v int32) *PromotionInfoQueryRequest {
	s.PageSize = &v
	return s
}

func (s *PromotionInfoQueryRequest) SetPromotionId(v []*string) *PromotionInfoQueryRequest {
	s.PromotionId = v
	return s
}

func (s *PromotionInfoQueryRequest) SetAccountId(v string) *PromotionInfoQueryRequest {
	s.AccountId = &v
	return s
}

func (s *PromotionInfoQueryRequest) SetHotelId(v []*string) *PromotionInfoQueryRequest {
	s.HotelId = v
	return s
}

func (s *PromotionInfoQueryRequest) SetHeader(v map[string]*string) *PromotionInfoQueryRequest {
	s.Header = v
	return s
}

type PromotionInfoQueryResponse struct {
	Extra *PromotionInfoQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PromotionInfoQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PromotionInfoQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponse) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponse) SetExtra(v *PromotionInfoQueryResponseExtra) *PromotionInfoQueryResponse {
	s.Extra = v
	return s
}

func (s *PromotionInfoQueryResponse) SetData(v *PromotionInfoQueryResponseData) *PromotionInfoQueryResponse {
	s.Data = v
	return s
}

type PromotionInfoQueryResponseData struct {
	Pagination    *PromotionInfoQueryResponseDataPagination       `json:"pagination,omitempty" xml:"pagination,omitempty"`
	Promotions    []*PromotionInfoQueryResponseDataPromotionsItem `json:"promotions,omitempty" xml:"promotions,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseData) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseData) SetPagination(v *PromotionInfoQueryResponseDataPagination) *PromotionInfoQueryResponseData {
	s.Pagination = v
	return s
}

func (s *PromotionInfoQueryResponseData) SetPromotions(v []*PromotionInfoQueryResponseDataPromotionsItem) *PromotionInfoQueryResponseData {
	s.Promotions = v
	return s
}

func (s *PromotionInfoQueryResponseData) SetGwErrorCode(v int32) *PromotionInfoQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PromotionInfoQueryResponseData) SetGwDescription(v string) *PromotionInfoQueryResponseData {
	s.GwDescription = &v
	return s
}

type PromotionInfoQueryResponseDataPagination struct {
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	HasMore    *bool  `json:"has_more,omitempty" xml:"has_more,omitempty"`
	PageCount  *int32 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	PageIndex  *int32 `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize   *int32 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPagination) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPagination) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPagination) SetTotalCount(v int32) *PromotionInfoQueryResponseDataPagination {
	s.TotalCount = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPagination) SetHasMore(v bool) *PromotionInfoQueryResponseDataPagination {
	s.HasMore = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPagination) SetPageCount(v int32) *PromotionInfoQueryResponseDataPagination {
	s.PageCount = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPagination) SetPageIndex(v int32) *PromotionInfoQueryResponseDataPagination {
	s.PageIndex = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPagination) SetPageSize(v int32) *PromotionInfoQueryResponseDataPagination {
	s.PageSize = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItem struct {
	FullPatternPromotion      *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion      `json:"full_pattern_promotion,omitempty" xml:"full_pattern_promotion,omitempty"`
	UnApplicableDate          *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate          `json:"un_applicable_date,omitempty" xml:"un_applicable_date,omitempty"`
	HotelNewCustomerPromotion *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion `json:"hotel_new_customer_promotion,omitempty" xml:"hotel_new_customer_promotion,omitempty"`
	SuccessionPromotion       *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion       `json:"succession_promotion,omitempty" xml:"succession_promotion,omitempty"`
	IsAutoExtension           *bool                                                                  `json:"is_auto_extension,omitempty" xml:"is_auto_extension,omitempty" require:"true"`
	FlashSalePromotion        *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion        `json:"flash_sale_promotion,omitempty" xml:"flash_sale_promotion,omitempty"`
	ApplicableDate            *PromotionInfoQueryResponseDataPromotionsItemApplicableDate            `json:"applicable_date,omitempty" xml:"applicable_date,omitempty" require:"true"`
	ApplicableResource        *PromotionInfoQueryResponseDataPromotionsItemApplicableResource        `json:"applicable_resource,omitempty" xml:"applicable_resource,omitempty" require:"true"`
	PromotionId               *string                                                                `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	PromotionBasicInfo        *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo        `json:"promotion_basic_info,omitempty" xml:"promotion_basic_info,omitempty" require:"true"`
	NightSalePromotion        *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion        `json:"night_sale_promotion,omitempty" xml:"night_sale_promotion,omitempty"`
	DailySalePromotion        *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion        `json:"daily_sale_promotion,omitempty" xml:"daily_sale_promotion,omitempty"`
	EarlyBirdPromotion        *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion        `json:"early_bird_promotion,omitempty" xml:"early_bird_promotion,omitempty"`
}

func (s PromotionInfoQueryResponseDataPromotionsItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetFullPatternPromotion(v *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.FullPatternPromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetUnApplicableDate(v *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate) *PromotionInfoQueryResponseDataPromotionsItem {
	s.UnApplicableDate = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetHotelNewCustomerPromotion(v *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.HotelNewCustomerPromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetSuccessionPromotion(v *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.SuccessionPromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetIsAutoExtension(v bool) *PromotionInfoQueryResponseDataPromotionsItem {
	s.IsAutoExtension = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetFlashSalePromotion(v *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.FlashSalePromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetApplicableDate(v *PromotionInfoQueryResponseDataPromotionsItemApplicableDate) *PromotionInfoQueryResponseDataPromotionsItem {
	s.ApplicableDate = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetApplicableResource(v *PromotionInfoQueryResponseDataPromotionsItemApplicableResource) *PromotionInfoQueryResponseDataPromotionsItem {
	s.ApplicableResource = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetPromotionId(v string) *PromotionInfoQueryResponseDataPromotionsItem {
	s.PromotionId = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetPromotionBasicInfo(v *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) *PromotionInfoQueryResponseDataPromotionsItem {
	s.PromotionBasicInfo = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetNightSalePromotion(v *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.NightSalePromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetDailySalePromotion(v *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.DailySalePromotion = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItem) SetEarlyBirdPromotion(v *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion) *PromotionInfoQueryResponseDataPromotionsItem {
	s.EarlyBirdPromotion = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemApplicableDate struct {
	EndDays    *string  `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
	StartDays  *string  `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" require:"true" type:"Repeated"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableDate) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableDate) SetEndDays(v string) *PromotionInfoQueryResponseDataPromotionsItemApplicableDate {
	s.EndDays = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableDate) SetStartDays(v string) *PromotionInfoQueryResponseDataPromotionsItemApplicableDate {
	s.StartDays = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableDate) SetDaysOfWeek(v []*int32) *PromotionInfoQueryResponseDataPromotionsItemApplicableDate {
	s.DaysOfWeek = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemApplicableResource struct {
	Resources []*PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableResource) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableResource) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableResource) SetResources(v []*PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem) *PromotionInfoQueryResponseDataPromotionsItemApplicableResource {
	s.Resources = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem struct {
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem) SetRatePlanId(v string) *PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem) SetRatePlanStatus(v int32) *PromotionInfoQueryResponseDataPromotionsItemApplicableResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion struct {
	DailySalePromotionDetail *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail `json:"daily_sale_promotion_detail,omitempty" xml:"daily_sale_promotion_detail,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion) SetDailySalePromotionDetail(v *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail) *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotion {
	s.DailySalePromotionDetail = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail struct {
	DiscountType  *int   `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail {
	s.DiscountType = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemDailySalePromotionDailySalePromotionDetail {
	s.DiscountValue = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion struct {
	EarlyBirdPromotionDetail []*PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem `json:"early_bird_promotion_detail,omitempty" xml:"early_bird_promotion_detail,omitempty" require:"true" type:"Repeated"`
	DiscountType             *int                                                                                          `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion) SetEarlyBirdPromotionDetail(v []*PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem) *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion {
	s.EarlyBirdPromotionDetail = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotion {
	s.DiscountType = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem struct {
	AdvanceDay    *int64 `json:"advance_day,omitempty" xml:"advance_day,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem) SetAdvanceDay(v int64) *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem {
	s.AdvanceDay = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemEarlyBirdPromotionEarlyBirdPromotionDetailItem {
	s.DiscountValue = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion struct {
	FlashSaleDetail []*PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem `json:"flash_sale_detail,omitempty" xml:"flash_sale_detail,omitempty" require:"true" type:"Repeated"`
	DiscountType    *int                                                                                 `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion) SetFlashSaleDetail(v []*PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion {
	s.FlashSaleDetail = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotion {
	s.DiscountType = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem struct {
	EndHour       *int   `json:"end_hour,omitempty" xml:"end_hour,omitempty" require:"true"`
	StartHour     *int   `json:"start_hour,omitempty" xml:"start_hour,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) SetEndHour(v int) *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem {
	s.EndHour = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) SetStartHour(v int) *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem {
	s.StartHour = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemFlashSalePromotionFlashSaleDetailItem {
	s.DiscountValue = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion struct {
	FullPatternPromotionDetail *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail `json:"full_pattern_promotion_detail,omitempty" xml:"full_pattern_promotion_detail,omitempty"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion) SetFullPatternPromotionDetail(v *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotion {
	s.FullPatternPromotionDetail = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail struct {
	StartHour          *int64   `json:"start_hour,omitempty" xml:"start_hour,omitempty"`
	AdvanceDays        *int64   `json:"advance_days,omitempty" xml:"advance_days,omitempty"`
	ConsecutiveDays    *int64   `json:"consecutive_days,omitempty" xml:"consecutive_days,omitempty"`
	DiscountType       *int     `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	EndHour            *int64   `json:"end_hour,omitempty" xml:"end_hour,omitempty"`
	IsHotelNewCustomer *bool    `json:"is_hotel_new_customer,omitempty" xml:"is_hotel_new_customer,omitempty"`
	MemberPromotion    []*int64 `json:"member_promotion,omitempty" xml:"member_promotion,omitempty" type:"Repeated"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetStartHour(v int64) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.StartHour = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetAdvanceDays(v int64) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.AdvanceDays = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetConsecutiveDays(v int64) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.ConsecutiveDays = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.DiscountType = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetEndHour(v int64) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.EndHour = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetIsHotelNewCustomer(v bool) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.IsHotelNewCustomer = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail) SetMemberPromotion(v []*int64) *PromotionInfoQueryResponseDataPromotionsItemFullPatternPromotionFullPatternPromotionDetail {
	s.MemberPromotion = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion struct {
	HotelNewCustomerPromotionDetail *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail `json:"hotel_new_customer_promotion_detail,omitempty" xml:"hotel_new_customer_promotion_detail,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion) SetHotelNewCustomerPromotionDetail(v *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotion {
	s.HotelNewCustomerPromotionDetail = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail struct {
	DiscountType  *int   `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail {
	s.DiscountType = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemHotelNewCustomerPromotionHotelNewCustomerPromotionDetail {
	s.DiscountValue = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion struct {
	NightSaleDetail []*PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem `json:"night_saleDetail,omitempty" xml:"night_saleDetail,omitempty" require:"true" type:"Repeated"`
	DiscountType    *int                                                                                 `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion) SetNightSaleDetail(v []*PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion {
	s.NightSaleDetail = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotion {
	s.DiscountType = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem struct {
	StartHour     *int   `json:"start_hour,omitempty" xml:"start_hour,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
	EndHour       *int   `json:"end_hour,omitempty" xml:"end_hour,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) SetStartHour(v int) *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem {
	s.StartHour = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem {
	s.DiscountValue = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem) SetEndHour(v int) *PromotionInfoQueryResponseDataPromotionsItemNightSalePromotionNightSaleDetailItem {
	s.EndHour = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo struct {
	OutType         *string `json:"out_type,omitempty" xml:"out_type,omitempty"`
	SubTypeCode     *string `json:"sub_type_code,omitempty" xml:"sub_type_code,omitempty"`
	SubTypeCodeName *string `json:"sub_type_code_name,omitempty" xml:"sub_type_code_name,omitempty"`
	TypeCode        *string `json:"type_code,omitempty" xml:"type_code,omitempty" require:"true"`
	TypeName        *string `json:"type_name,omitempty" xml:"type_name,omitempty"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) SetOutType(v string) *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo {
	s.OutType = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) SetSubTypeCode(v string) *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo {
	s.SubTypeCode = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) SetSubTypeCodeName(v string) *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo {
	s.SubTypeCodeName = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) SetTypeCode(v string) *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo {
	s.TypeCode = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo) SetTypeName(v string) *PromotionInfoQueryResponseDataPromotionsItemPromotionBasicInfo {
	s.TypeName = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion struct {
	SuccessionPromotionDetail []*PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem `json:"succession_promotion_detail,omitempty" xml:"succession_promotion_detail,omitempty" require:"true" type:"Repeated"`
	DiscountType              *int                                                                                            `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion) SetSuccessionPromotionDetail(v []*PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem) *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion {
	s.SuccessionPromotionDetail = v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion) SetDiscountType(v int) *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotion {
	s.DiscountType = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem struct {
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
	StayLength    *int64 `json:"stay_length,omitempty" xml:"stay_length,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem) SetDiscountValue(v int64) *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem {
	s.DiscountValue = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem) SetStayLength(v int64) *PromotionInfoQueryResponseDataPromotionsItemSuccessionPromotionSuccessionPromotionDetailItem {
	s.StayLength = &v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate struct {
	UnapplicableDateRange []*PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem `json:"unapplicable_date_range,omitempty" xml:"unapplicable_date_range,omitempty" type:"Repeated"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate) SetUnapplicableDateRange(v []*PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem) *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDate {
	s.UnapplicableDateRange = v
	return s
}

type PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem struct {
	StartDays *string `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
	EndDays   *string `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem) SetStartDays(v string) *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem {
	s.StartDays = &v
	return s
}

func (s *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem) SetEndDays(v string) *PromotionInfoQueryResponseDataPromotionsItemUnApplicableDateUnapplicableDateRangeItem {
	s.EndDays = &v
	return s
}

type PromotionInfoQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PromotionInfoQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PromotionInfoQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *PromotionInfoQueryResponseExtra) SetSubErrorCode(v int32) *PromotionInfoQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PromotionInfoQueryResponseExtra) SetDescription(v string) *PromotionInfoQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *PromotionInfoQueryResponseExtra) SetErrorCode(v int32) *PromotionInfoQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PromotionInfoQueryResponseExtra) SetLogid(v string) *PromotionInfoQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *PromotionInfoQueryResponseExtra) SetNow(v int64) *PromotionInfoQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *PromotionInfoQueryResponseExtra) SetSubDescription(v string) *PromotionInfoQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type PromotionPushRequest struct {
	ApplicableDate            *PromotionPushRequestApplicableDate            `json:"applicable_date,omitempty" xml:"applicable_date,omitempty" require:"true"`
	HotelNewCustomerPromotion *PromotionPushRequestHotelNewCustomerPromotion `json:"hotel_new_customer_promotion,omitempty" xml:"hotel_new_customer_promotion,omitempty"`
	PromotionBasicInfo        *PromotionPushRequestPromotionBasicInfo        `json:"promotion_basic_info,omitempty" xml:"promotion_basic_info,omitempty" require:"true"`
	IsAutoExtension           *bool                                          `json:"is_auto_extension,omitempty" xml:"is_auto_extension,omitempty" require:"true"`
	AccessToken               *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Active                    *bool                                          `json:"active,omitempty" xml:"active,omitempty"`
	FlashSalePromotion        *PromotionPushRequestFlashSalePromotion        `json:"flash_sale_promotion,omitempty" xml:"flash_sale_promotion,omitempty"`
	UnapplicableDate          *PromotionPushRequestUnapplicableDate          `json:"unapplicable_date,omitempty" xml:"unapplicable_date,omitempty"`
	PromotionId               *string                                        `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	Header                    map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	AccountId                 *string                                        `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	EarlyBirdPromotion        *PromotionPushRequestEarlyBirdPromotion        `json:"early_bird_promotion,omitempty" xml:"early_bird_promotion,omitempty"`
	SuccessionPromotion       *PromotionPushRequestSuccessionPromotion       `json:"succession_promotion,omitempty" xml:"succession_promotion,omitempty"`
	NightSalePromotion        *PromotionPushRequestNightSalePromotion        `json:"night_sale_promotion,omitempty" xml:"night_sale_promotion,omitempty"`
	ApplicableResource        *PromotionPushRequestApplicableResource        `json:"applicable_resource,omitempty" xml:"applicable_resource,omitempty" require:"true"`
	DailySalePromotion        *PromotionPushRequestDailySalePromotion        `json:"daily_sale_promotion,omitempty" xml:"daily_sale_promotion,omitempty"`
}

func (s PromotionPushRequest) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequest) GoString() string {
	return s.String()
}

func (s *PromotionPushRequest) SetApplicableDate(v *PromotionPushRequestApplicableDate) *PromotionPushRequest {
	s.ApplicableDate = v
	return s
}

func (s *PromotionPushRequest) SetHotelNewCustomerPromotion(v *PromotionPushRequestHotelNewCustomerPromotion) *PromotionPushRequest {
	s.HotelNewCustomerPromotion = v
	return s
}

func (s *PromotionPushRequest) SetPromotionBasicInfo(v *PromotionPushRequestPromotionBasicInfo) *PromotionPushRequest {
	s.PromotionBasicInfo = v
	return s
}

func (s *PromotionPushRequest) SetIsAutoExtension(v bool) *PromotionPushRequest {
	s.IsAutoExtension = &v
	return s
}

func (s *PromotionPushRequest) SetAccessToken(v string) *PromotionPushRequest {
	s.AccessToken = &v
	return s
}

func (s *PromotionPushRequest) SetActive(v bool) *PromotionPushRequest {
	s.Active = &v
	return s
}

func (s *PromotionPushRequest) SetFlashSalePromotion(v *PromotionPushRequestFlashSalePromotion) *PromotionPushRequest {
	s.FlashSalePromotion = v
	return s
}

func (s *PromotionPushRequest) SetUnapplicableDate(v *PromotionPushRequestUnapplicableDate) *PromotionPushRequest {
	s.UnapplicableDate = v
	return s
}

func (s *PromotionPushRequest) SetPromotionId(v string) *PromotionPushRequest {
	s.PromotionId = &v
	return s
}

func (s *PromotionPushRequest) SetHeader(v map[string]*string) *PromotionPushRequest {
	s.Header = v
	return s
}

func (s *PromotionPushRequest) SetAccountId(v string) *PromotionPushRequest {
	s.AccountId = &v
	return s
}

func (s *PromotionPushRequest) SetEarlyBirdPromotion(v *PromotionPushRequestEarlyBirdPromotion) *PromotionPushRequest {
	s.EarlyBirdPromotion = v
	return s
}

func (s *PromotionPushRequest) SetSuccessionPromotion(v *PromotionPushRequestSuccessionPromotion) *PromotionPushRequest {
	s.SuccessionPromotion = v
	return s
}

func (s *PromotionPushRequest) SetNightSalePromotion(v *PromotionPushRequestNightSalePromotion) *PromotionPushRequest {
	s.NightSalePromotion = v
	return s
}

func (s *PromotionPushRequest) SetApplicableResource(v *PromotionPushRequestApplicableResource) *PromotionPushRequest {
	s.ApplicableResource = v
	return s
}

func (s *PromotionPushRequest) SetDailySalePromotion(v *PromotionPushRequestDailySalePromotion) *PromotionPushRequest {
	s.DailySalePromotion = v
	return s
}

type PromotionPushRequestApplicableDate struct {
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" require:"true" type:"Repeated"`
	EndDays    *string  `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
	StartDays  *string  `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
}

func (s PromotionPushRequestApplicableDate) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestApplicableDate) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestApplicableDate) SetDaysOfWeek(v []*int32) *PromotionPushRequestApplicableDate {
	s.DaysOfWeek = v
	return s
}

func (s *PromotionPushRequestApplicableDate) SetEndDays(v string) *PromotionPushRequestApplicableDate {
	s.EndDays = &v
	return s
}

func (s *PromotionPushRequestApplicableDate) SetStartDays(v string) *PromotionPushRequestApplicableDate {
	s.StartDays = &v
	return s
}

type PromotionPushRequestApplicableResource struct {
	Resources []*PromotionPushRequestApplicableResourceResourcesItem `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s PromotionPushRequestApplicableResource) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestApplicableResource) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestApplicableResource) SetResources(v []*PromotionPushRequestApplicableResourceResourcesItem) *PromotionPushRequestApplicableResource {
	s.Resources = v
	return s
}

type PromotionPushRequestApplicableResourceResourcesItem struct {
	RatePlanStatus *int32  `json:"rate_plan_status,omitempty" xml:"rate_plan_status,omitempty"`
	RatePlanId     *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
}

func (s PromotionPushRequestApplicableResourceResourcesItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestApplicableResourceResourcesItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestApplicableResourceResourcesItem) SetRatePlanStatus(v int32) *PromotionPushRequestApplicableResourceResourcesItem {
	s.RatePlanStatus = &v
	return s
}

func (s *PromotionPushRequestApplicableResourceResourcesItem) SetRatePlanId(v string) *PromotionPushRequestApplicableResourceResourcesItem {
	s.RatePlanId = &v
	return s
}

type PromotionPushRequestDailySalePromotion struct {
	DailySalePromotionDetail *PromotionPushRequestDailySalePromotionDailySalePromotionDetail `json:"daily_sale_promotion_detail,omitempty" xml:"daily_sale_promotion_detail,omitempty" require:"true"`
}

func (s PromotionPushRequestDailySalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestDailySalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestDailySalePromotion) SetDailySalePromotionDetail(v *PromotionPushRequestDailySalePromotionDailySalePromotionDetail) *PromotionPushRequestDailySalePromotion {
	s.DailySalePromotionDetail = v
	return s
}

type PromotionPushRequestDailySalePromotionDailySalePromotionDetail struct {
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
	DiscountType  *int   `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionPushRequestDailySalePromotionDailySalePromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestDailySalePromotionDailySalePromotionDetail) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestDailySalePromotionDailySalePromotionDetail) SetDiscountValue(v int64) *PromotionPushRequestDailySalePromotionDailySalePromotionDetail {
	s.DiscountValue = &v
	return s
}

func (s *PromotionPushRequestDailySalePromotionDailySalePromotionDetail) SetDiscountType(v int) *PromotionPushRequestDailySalePromotionDailySalePromotionDetail {
	s.DiscountType = &v
	return s
}

type PromotionPushRequestEarlyBirdPromotion struct {
	DiscountType             *int                                                                  `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	EarlyBirdPromotionDetail []*PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem `json:"early_bird_promotion_detail,omitempty" xml:"early_bird_promotion_detail,omitempty" require:"true" type:"Repeated"`
}

func (s PromotionPushRequestEarlyBirdPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestEarlyBirdPromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestEarlyBirdPromotion) SetDiscountType(v int) *PromotionPushRequestEarlyBirdPromotion {
	s.DiscountType = &v
	return s
}

func (s *PromotionPushRequestEarlyBirdPromotion) SetEarlyBirdPromotionDetail(v []*PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem) *PromotionPushRequestEarlyBirdPromotion {
	s.EarlyBirdPromotionDetail = v
	return s
}

type PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem struct {
	AdvanceDay    *int64 `json:"advance_day,omitempty" xml:"advance_day,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem) SetAdvanceDay(v int64) *PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem {
	s.AdvanceDay = &v
	return s
}

func (s *PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem) SetDiscountValue(v int64) *PromotionPushRequestEarlyBirdPromotionEarlyBirdPromotionDetailItem {
	s.DiscountValue = &v
	return s
}

type PromotionPushRequestFlashSalePromotion struct {
	DiscountType    *int                                                         `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	FlashSaleDetail []*PromotionPushRequestFlashSalePromotionFlashSaleDetailItem `json:"flash_sale_detail,omitempty" xml:"flash_sale_detail,omitempty" require:"true" type:"Repeated"`
}

func (s PromotionPushRequestFlashSalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestFlashSalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestFlashSalePromotion) SetDiscountType(v int) *PromotionPushRequestFlashSalePromotion {
	s.DiscountType = &v
	return s
}

func (s *PromotionPushRequestFlashSalePromotion) SetFlashSaleDetail(v []*PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) *PromotionPushRequestFlashSalePromotion {
	s.FlashSaleDetail = v
	return s
}

type PromotionPushRequestFlashSalePromotionFlashSaleDetailItem struct {
	EndHour       *int   `json:"end_hour,omitempty" xml:"end_hour,omitempty" require:"true"`
	StartHour     *int   `json:"start_hour,omitempty" xml:"start_hour,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) SetEndHour(v int) *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem {
	s.EndHour = &v
	return s
}

func (s *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) SetStartHour(v int) *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem {
	s.StartHour = &v
	return s
}

func (s *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem) SetDiscountValue(v int64) *PromotionPushRequestFlashSalePromotionFlashSaleDetailItem {
	s.DiscountValue = &v
	return s
}

type PromotionPushRequestHotelNewCustomerPromotion struct {
	HotelNewCustomerPromotionDetail *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail `json:"hotel_new_customer_promotion_detail,omitempty" xml:"hotel_new_customer_promotion_detail,omitempty" require:"true"`
}

func (s PromotionPushRequestHotelNewCustomerPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestHotelNewCustomerPromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestHotelNewCustomerPromotion) SetHotelNewCustomerPromotionDetail(v *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) *PromotionPushRequestHotelNewCustomerPromotion {
	s.HotelNewCustomerPromotionDetail = v
	return s
}

type PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail struct {
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
	DiscountType  *int   `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) SetDiscountValue(v int64) *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail {
	s.DiscountValue = &v
	return s
}

func (s *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail) SetDiscountType(v int) *PromotionPushRequestHotelNewCustomerPromotionHotelNewCustomerPromotionDetail {
	s.DiscountType = &v
	return s
}

type PromotionPushRequestNightSalePromotion struct {
	DiscountType    *int                                                         `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	NightSaleDetail []*PromotionPushRequestNightSalePromotionNightSaleDetailItem `json:"night_saleDetail,omitempty" xml:"night_saleDetail,omitempty" require:"true" type:"Repeated"`
}

func (s PromotionPushRequestNightSalePromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestNightSalePromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestNightSalePromotion) SetDiscountType(v int) *PromotionPushRequestNightSalePromotion {
	s.DiscountType = &v
	return s
}

func (s *PromotionPushRequestNightSalePromotion) SetNightSaleDetail(v []*PromotionPushRequestNightSalePromotionNightSaleDetailItem) *PromotionPushRequestNightSalePromotion {
	s.NightSaleDetail = v
	return s
}

type PromotionPushRequestNightSalePromotionNightSaleDetailItem struct {
	EndHour       *int   `json:"end_hour,omitempty" xml:"end_hour,omitempty" require:"true"`
	StartHour     *int   `json:"start_hour,omitempty" xml:"start_hour,omitempty" require:"true"`
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
}

func (s PromotionPushRequestNightSalePromotionNightSaleDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestNightSalePromotionNightSaleDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestNightSalePromotionNightSaleDetailItem) SetEndHour(v int) *PromotionPushRequestNightSalePromotionNightSaleDetailItem {
	s.EndHour = &v
	return s
}

func (s *PromotionPushRequestNightSalePromotionNightSaleDetailItem) SetStartHour(v int) *PromotionPushRequestNightSalePromotionNightSaleDetailItem {
	s.StartHour = &v
	return s
}

func (s *PromotionPushRequestNightSalePromotionNightSaleDetailItem) SetDiscountValue(v int64) *PromotionPushRequestNightSalePromotionNightSaleDetailItem {
	s.DiscountValue = &v
	return s
}

type PromotionPushRequestPromotionBasicInfo struct {
	TypeName        *string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	OutType         *string `json:"out_type,omitempty" xml:"out_type,omitempty"`
	SubTypeCode     *string `json:"sub_type_code,omitempty" xml:"sub_type_code,omitempty"`
	SubTypeCodeName *string `json:"sub_type_code_name,omitempty" xml:"sub_type_code_name,omitempty"`
	TypeCode        *string `json:"type_code,omitempty" xml:"type_code,omitempty" require:"true"`
}

func (s PromotionPushRequestPromotionBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestPromotionBasicInfo) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestPromotionBasicInfo) SetTypeName(v string) *PromotionPushRequestPromotionBasicInfo {
	s.TypeName = &v
	return s
}

func (s *PromotionPushRequestPromotionBasicInfo) SetOutType(v string) *PromotionPushRequestPromotionBasicInfo {
	s.OutType = &v
	return s
}

func (s *PromotionPushRequestPromotionBasicInfo) SetSubTypeCode(v string) *PromotionPushRequestPromotionBasicInfo {
	s.SubTypeCode = &v
	return s
}

func (s *PromotionPushRequestPromotionBasicInfo) SetSubTypeCodeName(v string) *PromotionPushRequestPromotionBasicInfo {
	s.SubTypeCodeName = &v
	return s
}

func (s *PromotionPushRequestPromotionBasicInfo) SetTypeCode(v string) *PromotionPushRequestPromotionBasicInfo {
	s.TypeCode = &v
	return s
}

type PromotionPushRequestSuccessionPromotion struct {
	DiscountType              *int                                                                    `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	SuccessionPromotionDetail []*PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem `json:"succession_promotion_detail,omitempty" xml:"succession_promotion_detail,omitempty" require:"true" type:"Repeated"`
}

func (s PromotionPushRequestSuccessionPromotion) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestSuccessionPromotion) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestSuccessionPromotion) SetDiscountType(v int) *PromotionPushRequestSuccessionPromotion {
	s.DiscountType = &v
	return s
}

func (s *PromotionPushRequestSuccessionPromotion) SetSuccessionPromotionDetail(v []*PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem) *PromotionPushRequestSuccessionPromotion {
	s.SuccessionPromotionDetail = v
	return s
}

type PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem struct {
	DiscountValue *int64 `json:"discount_value,omitempty" xml:"discount_value,omitempty" require:"true"`
	StayLength    *int64 `json:"stay_length,omitempty" xml:"stay_length,omitempty" require:"true"`
}

func (s PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem) SetDiscountValue(v int64) *PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem {
	s.DiscountValue = &v
	return s
}

func (s *PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem) SetStayLength(v int64) *PromotionPushRequestSuccessionPromotionSuccessionPromotionDetailItem {
	s.StayLength = &v
	return s
}

type PromotionPushRequestUnapplicableDate struct {
	UnapplicableDateRange []*PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem `json:"unapplicable_date_range,omitempty" xml:"unapplicable_date_range,omitempty" type:"Repeated"`
}

func (s PromotionPushRequestUnapplicableDate) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestUnapplicableDate) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestUnapplicableDate) SetUnapplicableDateRange(v []*PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem) *PromotionPushRequestUnapplicableDate {
	s.UnapplicableDateRange = v
	return s
}

type PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem struct {
	StartDays *string `json:"start_days,omitempty" xml:"start_days,omitempty" require:"true"`
	EndDays   *string `json:"end_days,omitempty" xml:"end_days,omitempty" require:"true"`
}

func (s PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem) GoString() string {
	return s.String()
}

func (s *PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem) SetStartDays(v string) *PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem {
	s.StartDays = &v
	return s
}

func (s *PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem) SetEndDays(v string) *PromotionPushRequestUnapplicableDateUnapplicableDateRangeItem {
	s.EndDays = &v
	return s
}

type PromotionPushResponse struct {
	Data  *PromotionPushResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *PromotionPushResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PromotionPushResponse) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushResponse) GoString() string {
	return s.String()
}

func (s *PromotionPushResponse) SetData(v *PromotionPushResponseData) *PromotionPushResponse {
	s.Data = v
	return s
}

func (s *PromotionPushResponse) SetExtra(v *PromotionPushResponseExtra) *PromotionPushResponse {
	s.Extra = v
	return s
}

type PromotionPushResponseData struct {
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PromotionId   *string                                       `json:"promotion_id,omitempty" xml:"promotion_id,omitempty" require:"true"`
	Status        *int                                          `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	MessageDetail []*PromotionPushResponseDataMessageDetailItem `json:"message_detail,omitempty" xml:"message_detail,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PromotionPushResponseData) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushResponseData) GoString() string {
	return s.String()
}

func (s *PromotionPushResponseData) SetGwDescription(v string) *PromotionPushResponseData {
	s.GwDescription = &v
	return s
}

func (s *PromotionPushResponseData) SetPromotionId(v string) *PromotionPushResponseData {
	s.PromotionId = &v
	return s
}

func (s *PromotionPushResponseData) SetStatus(v int) *PromotionPushResponseData {
	s.Status = &v
	return s
}

func (s *PromotionPushResponseData) SetMessageDetail(v []*PromotionPushResponseDataMessageDetailItem) *PromotionPushResponseData {
	s.MessageDetail = v
	return s
}

func (s *PromotionPushResponseData) SetGwErrorCode(v int32) *PromotionPushResponseData {
	s.GwErrorCode = &v
	return s
}

type PromotionPushResponseDataMessageDetailItem struct {
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	RatePlanId   *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
}

func (s PromotionPushResponseDataMessageDetailItem) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushResponseDataMessageDetailItem) GoString() string {
	return s.String()
}

func (s *PromotionPushResponseDataMessageDetailItem) SetErrorMessage(v string) *PromotionPushResponseDataMessageDetailItem {
	s.ErrorMessage = &v
	return s
}

func (s *PromotionPushResponseDataMessageDetailItem) SetRatePlanId(v string) *PromotionPushResponseDataMessageDetailItem {
	s.RatePlanId = &v
	return s
}

type PromotionPushResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s PromotionPushResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PromotionPushResponseExtra) GoString() string {
	return s.String()
}

func (s *PromotionPushResponseExtra) SetDescription(v string) *PromotionPushResponseExtra {
	s.Description = &v
	return s
}

func (s *PromotionPushResponseExtra) SetErrorCode(v int32) *PromotionPushResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PromotionPushResponseExtra) SetLogid(v string) *PromotionPushResponseExtra {
	s.Logid = &v
	return s
}

func (s *PromotionPushResponseExtra) SetNow(v int64) *PromotionPushResponseExtra {
	s.Now = &v
	return s
}

func (s *PromotionPushResponseExtra) SetSubDescription(v string) *PromotionPushResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PromotionPushResponseExtra) SetSubErrorCode(v int32) *PromotionPushResponseExtra {
	s.SubErrorCode = &v
	return s
}

type PurchaseInfoRequest struct {
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ServiceId     *string            `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	ServiceModeId *string            `json:"service_mode_id,omitempty" xml:"service_mode_id,omitempty" require:"true"`
	PurchaseTime  *int64             `json:"purchase_time,omitempty" xml:"purchase_time,omitempty" require:"true"`
	OutTradeNo    *string            `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty" require:"true"`
	PeriodType    *int               `json:"period_type,omitempty" xml:"period_type,omitempty" require:"true"`
	Duration      *int64             `json:"duration,omitempty" xml:"duration,omitempty"`
	UsageTimes    *int64             `json:"usage_times,omitempty" xml:"usage_times,omitempty"`
}

func (s PurchaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s PurchaseInfoRequest) GoString() string {
	return s.String()
}

func (s *PurchaseInfoRequest) SetOpenId(v string) *PurchaseInfoRequest {
	s.OpenId = &v
	return s
}

func (s *PurchaseInfoRequest) SetHeader(v map[string]*string) *PurchaseInfoRequest {
	s.Header = v
	return s
}

func (s *PurchaseInfoRequest) SetAccessToken(v string) *PurchaseInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *PurchaseInfoRequest) SetServiceId(v string) *PurchaseInfoRequest {
	s.ServiceId = &v
	return s
}

func (s *PurchaseInfoRequest) SetServiceModeId(v string) *PurchaseInfoRequest {
	s.ServiceModeId = &v
	return s
}

func (s *PurchaseInfoRequest) SetPurchaseTime(v int64) *PurchaseInfoRequest {
	s.PurchaseTime = &v
	return s
}

func (s *PurchaseInfoRequest) SetOutTradeNo(v string) *PurchaseInfoRequest {
	s.OutTradeNo = &v
	return s
}

func (s *PurchaseInfoRequest) SetPeriodType(v int) *PurchaseInfoRequest {
	s.PeriodType = &v
	return s
}

func (s *PurchaseInfoRequest) SetDuration(v int64) *PurchaseInfoRequest {
	s.Duration = &v
	return s
}

func (s *PurchaseInfoRequest) SetUsageTimes(v int64) *PurchaseInfoRequest {
	s.UsageTimes = &v
	return s
}

type PurchaseInfoResponse struct {
	Data  *PurchaseInfoResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *PurchaseInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s PurchaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PurchaseInfoResponse) GoString() string {
	return s.String()
}

func (s *PurchaseInfoResponse) SetData(v *PurchaseInfoResponseData) *PurchaseInfoResponse {
	s.Data = v
	return s
}

func (s *PurchaseInfoResponse) SetExtra(v *PurchaseInfoResponseExtra) *PurchaseInfoResponse {
	s.Extra = v
	return s
}

type PurchaseInfoResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s PurchaseInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s PurchaseInfoResponseData) GoString() string {
	return s.String()
}

func (s *PurchaseInfoResponseData) SetGwErrorCode(v int32) *PurchaseInfoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *PurchaseInfoResponseData) SetGwDescription(v string) *PurchaseInfoResponseData {
	s.GwDescription = &v
	return s
}

type PurchaseInfoResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s PurchaseInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PurchaseInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *PurchaseInfoResponseExtra) SetErrorCode(v int32) *PurchaseInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PurchaseInfoResponseExtra) SetDescription(v string) *PurchaseInfoResponseExtra {
	s.Description = &v
	return s
}

func (s *PurchaseInfoResponseExtra) SetSubErrorCode(v int32) *PurchaseInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PurchaseInfoResponseExtra) SetSubDescription(v string) *PurchaseInfoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *PurchaseInfoResponseExtra) SetLogid(v string) *PurchaseInfoResponseExtra {
	s.Logid = &v
	return s
}

func (s *PurchaseInfoResponseExtra) SetNow(v int64) *PurchaseInfoResponseExtra {
	s.Now = &v
	return s
}

type PurchaseListRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ServiceId   *string            `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	IsTestEnv   *bool              `json:"is_test_env,omitempty" xml:"is_test_env,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s PurchaseListRequest) String() string {
	return tea.Prettify(s)
}

func (s PurchaseListRequest) GoString() string {
	return s.String()
}

func (s *PurchaseListRequest) SetOpenId(v string) *PurchaseListRequest {
	s.OpenId = &v
	return s
}

func (s *PurchaseListRequest) SetServiceId(v string) *PurchaseListRequest {
	s.ServiceId = &v
	return s
}

func (s *PurchaseListRequest) SetIsTestEnv(v bool) *PurchaseListRequest {
	s.IsTestEnv = &v
	return s
}

func (s *PurchaseListRequest) SetHeader(v map[string]*string) *PurchaseListRequest {
	s.Header = v
	return s
}

func (s *PurchaseListRequest) SetAccessToken(v string) *PurchaseListRequest {
	s.AccessToken = &v
	return s
}

type PurchaseListResponse struct {
	Extra *PurchaseListResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *PurchaseListResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PurchaseListResponse) String() string {
	return tea.Prettify(s)
}

func (s PurchaseListResponse) GoString() string {
	return s.String()
}

func (s *PurchaseListResponse) SetExtra(v *PurchaseListResponseExtra) *PurchaseListResponse {
	s.Extra = v
	return s
}

func (s *PurchaseListResponse) SetData(v *PurchaseListResponseData) *PurchaseListResponse {
	s.Data = v
	return s
}

type PurchaseListResponseData struct {
	GwDescription    *string                                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PurchaseInfoList []*PurchaseListResponseDataPurchaseInfoListItem `json:"purchase_info_list,omitempty" xml:"purchase_info_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode      *int32                                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s PurchaseListResponseData) String() string {
	return tea.Prettify(s)
}

func (s PurchaseListResponseData) GoString() string {
	return s.String()
}

func (s *PurchaseListResponseData) SetGwDescription(v string) *PurchaseListResponseData {
	s.GwDescription = &v
	return s
}

func (s *PurchaseListResponseData) SetPurchaseInfoList(v []*PurchaseListResponseDataPurchaseInfoListItem) *PurchaseListResponseData {
	s.PurchaseInfoList = v
	return s
}

func (s *PurchaseListResponseData) SetGwErrorCode(v int32) *PurchaseListResponseData {
	s.GwErrorCode = &v
	return s
}

type PurchaseListResponseDataPurchaseInfoListItem struct {
	ServiceModeId      *string `json:"service_mode_id,omitempty" xml:"service_mode_id,omitempty"`
	ServiceStatus      *int    `json:"service_status,omitempty" xml:"service_status,omitempty" require:"true"`
	RemainServiceTimes *int64  `json:"remain_service_times,omitempty" xml:"remain_service_times,omitempty"`
	ServiceName        *string `json:"service_name,omitempty" xml:"service_name,omitempty" require:"true"`
	ServiceId          *string `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	EffectiveTime      *int64  `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	ExpireTime         *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	SpecificationType  *int    `json:"specification_type,omitempty" xml:"specification_type,omitempty"`
	SpecificationTitle *string `json:"specification_title,omitempty" xml:"specification_title,omitempty"`
}

func (s PurchaseListResponseDataPurchaseInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s PurchaseListResponseDataPurchaseInfoListItem) GoString() string {
	return s.String()
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetServiceModeId(v string) *PurchaseListResponseDataPurchaseInfoListItem {
	s.ServiceModeId = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetServiceStatus(v int) *PurchaseListResponseDataPurchaseInfoListItem {
	s.ServiceStatus = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetRemainServiceTimes(v int64) *PurchaseListResponseDataPurchaseInfoListItem {
	s.RemainServiceTimes = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetServiceName(v string) *PurchaseListResponseDataPurchaseInfoListItem {
	s.ServiceName = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetServiceId(v string) *PurchaseListResponseDataPurchaseInfoListItem {
	s.ServiceId = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetEffectiveTime(v int64) *PurchaseListResponseDataPurchaseInfoListItem {
	s.EffectiveTime = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetExpireTime(v int64) *PurchaseListResponseDataPurchaseInfoListItem {
	s.ExpireTime = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetSpecificationType(v int) *PurchaseListResponseDataPurchaseInfoListItem {
	s.SpecificationType = &v
	return s
}

func (s *PurchaseListResponseDataPurchaseInfoListItem) SetSpecificationTitle(v string) *PurchaseListResponseDataPurchaseInfoListItem {
	s.SpecificationTitle = &v
	return s
}

type PurchaseListResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s PurchaseListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s PurchaseListResponseExtra) GoString() string {
	return s.String()
}

func (s *PurchaseListResponseExtra) SetLogid(v string) *PurchaseListResponseExtra {
	s.Logid = &v
	return s
}

func (s *PurchaseListResponseExtra) SetNow(v int64) *PurchaseListResponseExtra {
	s.Now = &v
	return s
}

func (s *PurchaseListResponseExtra) SetErrorCode(v int32) *PurchaseListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *PurchaseListResponseExtra) SetDescription(v string) *PurchaseListResponseExtra {
	s.Description = &v
	return s
}

func (s *PurchaseListResponseExtra) SetSubErrorCode(v int32) *PurchaseListResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *PurchaseListResponseExtra) SetSubDescription(v string) *PurchaseListResponseExtra {
	s.SubDescription = &v
	return s
}

type QaOpenapiSdkPostCommonRequest struct {
	DefaultQuery   *string                                      `json:"default_query,omitempty" xml:"default_query,omitempty"`
	QaExtra        *QaOpenapiSdkPostCommonRequestQaExtra        `json:"qa_extra,omitempty" xml:"qa_extra,omitempty" require:"true"`
	OptionalStruct *QaOpenapiSdkPostCommonRequestOptionalStruct `json:"optional_struct,omitempty" xml:"optional_struct,omitempty"`
	DefaultList    []*string                                    `json:"default_list,omitempty" xml:"default_list,omitempty" type:"Repeated"`
	OptionalQuery  *string                                      `json:"optional_query,omitempty" xml:"optional_query,omitempty"`
	RequiredQuery  *string                                      `json:"required_query,omitempty" xml:"required_query,omitempty" require:"true"`
	OptionalList   []*string                                    `json:"optional_list,omitempty" xml:"optional_list,omitempty" type:"Repeated"`
	RequiredList   []*string                                    `json:"required_list,omitempty" xml:"required_list,omitempty" require:"true" type:"Repeated"`
	RequiredStruct *QaOpenapiSdkPostCommonRequestRequiredStruct `json:"required_struct,omitempty" xml:"required_struct,omitempty" require:"true"`
	DefaultMap     map[string]*string                           `json:"default_map,omitempty" xml:"default_map,omitempty"`
	Header         map[string]*string                           `json:"header,omitempty" xml:"header,omitempty"`
	DefaultStruct  *QaOpenapiSdkPostCommonRequestDefaultStruct  `json:"default_struct,omitempty" xml:"default_struct,omitempty"`
	RequiredMap    map[string]*string                           `json:"required_map,omitempty" xml:"required_map,omitempty" require:"true"`
	OptionalMap    map[string]*string                           `json:"optional_map,omitempty" xml:"optional_map,omitempty"`
	AccessToken    *string                                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QaOpenapiSdkPostCommonRequest) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonRequest) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonRequest) SetDefaultQuery(v string) *QaOpenapiSdkPostCommonRequest {
	s.DefaultQuery = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetQaExtra(v *QaOpenapiSdkPostCommonRequestQaExtra) *QaOpenapiSdkPostCommonRequest {
	s.QaExtra = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetOptionalStruct(v *QaOpenapiSdkPostCommonRequestOptionalStruct) *QaOpenapiSdkPostCommonRequest {
	s.OptionalStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetDefaultList(v []*string) *QaOpenapiSdkPostCommonRequest {
	s.DefaultList = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetOptionalQuery(v string) *QaOpenapiSdkPostCommonRequest {
	s.OptionalQuery = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetRequiredQuery(v string) *QaOpenapiSdkPostCommonRequest {
	s.RequiredQuery = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetOptionalList(v []*string) *QaOpenapiSdkPostCommonRequest {
	s.OptionalList = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetRequiredList(v []*string) *QaOpenapiSdkPostCommonRequest {
	s.RequiredList = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetRequiredStruct(v *QaOpenapiSdkPostCommonRequestRequiredStruct) *QaOpenapiSdkPostCommonRequest {
	s.RequiredStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetDefaultMap(v map[string]*string) *QaOpenapiSdkPostCommonRequest {
	s.DefaultMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetHeader(v map[string]*string) *QaOpenapiSdkPostCommonRequest {
	s.Header = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetDefaultStruct(v *QaOpenapiSdkPostCommonRequestDefaultStruct) *QaOpenapiSdkPostCommonRequest {
	s.DefaultStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetRequiredMap(v map[string]*string) *QaOpenapiSdkPostCommonRequest {
	s.RequiredMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetOptionalMap(v map[string]*string) *QaOpenapiSdkPostCommonRequest {
	s.OptionalMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequest) SetAccessToken(v string) *QaOpenapiSdkPostCommonRequest {
	s.AccessToken = &v
	return s
}

type QaOpenapiSdkPostCommonRequestDefaultStruct struct {
	DefaultFloat64         *float64 `json:"default_float64,omitempty" xml:"default_float64,omitempty"`
	DefaultString          *string  `json:"default_string,omitempty" xml:"default_string,omitempty"`
	DefaultByteSliceString []byte   `json:"default_byte_slice_string,omitempty" xml:"default_byte_slice_string,omitempty"`
	DefaultBool            *bool    `json:"default_bool,omitempty" xml:"default_bool,omitempty"`
	DefaultInt8            *int     `json:"default_int8,omitempty" xml:"default_int8,omitempty"`
	DefaultInt16           *int     `json:"default_int16,omitempty" xml:"default_int16,omitempty"`
	DefaultInt32           *int32   `json:"default_int32,omitempty" xml:"default_int32,omitempty"`
	DefaultInt64           *int64   `json:"default_int64,omitempty" xml:"default_int64,omitempty"`
}

func (s QaOpenapiSdkPostCommonRequestDefaultStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonRequestDefaultStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultFloat64(v float64) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultFloat64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultString(v string) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultByteSliceString(v []byte) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultBool(v bool) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultInt8(v int) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultInt8 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultInt16(v int) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultInt32(v int32) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultInt32 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestDefaultStruct) SetDefaultInt64(v int64) *QaOpenapiSdkPostCommonRequestDefaultStruct {
	s.DefaultInt64 = &v
	return s
}

type QaOpenapiSdkPostCommonRequestOptionalStruct struct {
	OptionalInt64           *int64   `json:"optional_int64,omitempty" xml:"optional_int64,omitempty"`
	OptionalFloat64         *float64 `json:"optional_float64,omitempty" xml:"optional_float64,omitempty"`
	OptionalString          *string  `json:"optional_string,omitempty" xml:"optional_string,omitempty"`
	OptionalByteSliceString []byte   `json:"optional_byte_slice_string,omitempty" xml:"optional_byte_slice_string,omitempty"`
	OptionalBool            *bool    `json:"optional_bool,omitempty" xml:"optional_bool,omitempty"`
	OptionalInt8            *int     `json:"optional_int8,omitempty" xml:"optional_int8,omitempty"`
	OptionalInt16           *int     `json:"optional_int16,omitempty" xml:"optional_int16,omitempty"`
	OptionalInt32           *int32   `json:"optional_int32,omitempty" xml:"optional_int32,omitempty"`
}

func (s QaOpenapiSdkPostCommonRequestOptionalStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonRequestOptionalStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalInt64(v int64) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalInt64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalFloat64(v float64) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalFloat64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalString(v string) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalByteSliceString(v []byte) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalBool(v bool) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalInt8(v int) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalInt8 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalInt16(v int) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestOptionalStruct) SetOptionalInt32(v int32) *QaOpenapiSdkPostCommonRequestOptionalStruct {
	s.OptionalInt32 = &v
	return s
}

type QaOpenapiSdkPostCommonRequestQaExtra struct {
	IsPanic          *bool   `json:"is_panic,omitempty" xml:"is_panic,omitempty" require:"true"`
	ProductSource    *int    `json:"product_source,omitempty" xml:"product_source,omitempty" require:"true"`
	ErrNo            *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	IsDeleteRespData *bool   `json:"is_delete_resp_data,omitempty" xml:"is_delete_resp_data,omitempty" require:"true"`
	SceneDesc        *string `json:"scene_desc,omitempty" xml:"scene_desc,omitempty" require:"true"`
	TimeOutSecond    *int64  `json:"time_out_second,omitempty" xml:"time_out_second,omitempty" require:"true"`
}

func (s QaOpenapiSdkPostCommonRequestQaExtra) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonRequestQaExtra) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetIsPanic(v bool) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.IsPanic = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetProductSource(v int) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.ProductSource = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetErrNo(v int32) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.ErrNo = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetIsDeleteRespData(v bool) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.IsDeleteRespData = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetSceneDesc(v string) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.SceneDesc = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestQaExtra) SetTimeOutSecond(v int64) *QaOpenapiSdkPostCommonRequestQaExtra {
	s.TimeOutSecond = &v
	return s
}

type QaOpenapiSdkPostCommonRequestRequiredStruct struct {
	RequiredInt16           *int     `json:"required_int16,omitempty" xml:"required_int16,omitempty" require:"true"`
	RequiredInt32           *int32   `json:"required_int32,omitempty" xml:"required_int32,omitempty" require:"true"`
	RequiredInt64           *int64   `json:"required_int64,omitempty" xml:"required_int64,omitempty" require:"true"`
	RequiredFloat64         *float64 `json:"required_float64,omitempty" xml:"required_float64,omitempty" require:"true"`
	RequiredString          *string  `json:"required_string,omitempty" xml:"required_string,omitempty" require:"true"`
	RequiredByteSliceString []byte   `json:"required_byte_slice_string,omitempty" xml:"required_byte_slice_string,omitempty" require:"true"`
	RequiredBool            *bool    `json:"required_bool,omitempty" xml:"required_bool,omitempty" require:"true"`
	RequiredInt8            *int     `json:"required_int8,omitempty" xml:"required_int8,omitempty" require:"true"`
}

func (s QaOpenapiSdkPostCommonRequestRequiredStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonRequestRequiredStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredInt16(v int) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredInt32(v int32) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredInt32 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredInt64(v int64) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredInt64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredFloat64(v float64) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredFloat64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredString(v string) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredByteSliceString(v []byte) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredBool(v bool) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonRequestRequiredStruct) SetRequiredInt8(v int) *QaOpenapiSdkPostCommonRequestRequiredStruct {
	s.RequiredInt8 = &v
	return s
}

type QaOpenapiSdkPostCommonResponse struct {
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QaOpenapiSdkPostCommonResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QaOpenapiSdkPostCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponse) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponse) SetErrMsg(v string) *QaOpenapiSdkPostCommonResponse {
	s.ErrMsg = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponse) SetLogId(v string) *QaOpenapiSdkPostCommonResponse {
	s.LogId = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponse) SetData(v *QaOpenapiSdkPostCommonResponseData) *QaOpenapiSdkPostCommonResponse {
	s.Data = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponse) SetErrNo(v int32) *QaOpenapiSdkPostCommonResponse {
	s.ErrNo = &v
	return s
}

type QaOpenapiSdkPostCommonResponseData struct {
	V3GwClientKey            *string                                           `json:"V3GwClientKey,omitempty" xml:"V3GwClientKey,omitempty"`
	V3GwAccessTokenType      *int32                                            `json:"V3GwAccessTokenType,omitempty" xml:"V3GwAccessTokenType,omitempty"`
	RequiredQuery            *string                                           `json:"required_query,omitempty" xml:"required_query,omitempty" require:"true"`
	V3GwClientInfo           *string                                           `json:"V3GwClientInfo,omitempty" xml:"V3GwClientInfo,omitempty"`
	RequiredMap              map[string]*string                                `json:"required_map,omitempty" xml:"required_map,omitempty" require:"true"`
	V3GwAppType              *int64                                            `json:"V3GwAppType,omitempty" xml:"V3GwAppType,omitempty"`
	V3GwISVClientKey         *string                                           `json:"V3GwISVClientKey,omitempty" xml:"V3GwISVClientKey,omitempty"`
	V3GwUserScopes           []*string                                         `json:"V3GwUserScopes,omitempty" xml:"V3GwUserScopes,omitempty" type:"Repeated"`
	V3GwClientScopes         []*string                                         `json:"V3GwClientScopes,omitempty" xml:"V3GwClientScopes,omitempty" type:"Repeated"`
	V3GwAccessTokenEncrypted *string                                           `json:"V3GwAccessTokenEncrypted,omitempty" xml:"V3GwAccessTokenEncrypted,omitempty"`
	RequiredList             []*string                                         `json:"required_list,omitempty" xml:"required_list,omitempty" require:"true" type:"Repeated"`
	QaExtra                  *QaOpenapiSdkPostCommonResponseDataQaExtra        `json:"qa_extra,omitempty" xml:"qa_extra,omitempty" require:"true"`
	V3GwUserId               *string                                           `json:"V3GwUserId,omitempty" xml:"V3GwUserId,omitempty"`
	DefaultStruct            *QaOpenapiSdkPostCommonResponseDataDefaultStruct  `json:"default_struct,omitempty" xml:"default_struct,omitempty"`
	V3GwClientIP             *string                                           `json:"V3GwClientIP,omitempty" xml:"V3GwClientIP,omitempty"`
	DefaultMap               map[string]*string                                `json:"default_map,omitempty" xml:"default_map,omitempty"`
	OptionalQuery            *string                                           `json:"optional_query,omitempty" xml:"optional_query,omitempty"`
	OptionalStruct           *QaOpenapiSdkPostCommonResponseDataOptionalStruct `json:"optional_struct,omitempty" xml:"optional_struct,omitempty"`
	OptionalList             []*string                                         `json:"optional_list,omitempty" xml:"optional_list,omitempty" type:"Repeated"`
	DefaultList              []*string                                         `json:"default_list,omitempty" xml:"default_list,omitempty" type:"Repeated"`
	V3GwOpenId               *string                                           `json:"V3GwOpenId,omitempty" xml:"V3GwOpenId,omitempty"`
	RawUri                   *string                                           `json:"raw_uri,omitempty" xml:"raw_uri,omitempty"`
	RequiredStruct           *QaOpenapiSdkPostCommonResponseDataRequiredStruct `json:"required_struct,omitempty" xml:"required_struct,omitempty" require:"true"`
	OptionalMap              map[string]*string                                `json:"optional_map,omitempty" xml:"optional_map,omitempty"`
	DefaultQuery             *string                                           `json:"default_query,omitempty" xml:"default_query,omitempty"`
}

func (s QaOpenapiSdkPostCommonResponseData) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponseData) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwClientKey(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwClientKey = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwAccessTokenType(v int32) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwAccessTokenType = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetRequiredQuery(v string) *QaOpenapiSdkPostCommonResponseData {
	s.RequiredQuery = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwClientInfo(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwClientInfo = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetRequiredMap(v map[string]*string) *QaOpenapiSdkPostCommonResponseData {
	s.RequiredMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwAppType(v int64) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwAppType = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwISVClientKey(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwISVClientKey = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwUserScopes(v []*string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwUserScopes = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwClientScopes(v []*string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwClientScopes = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwAccessTokenEncrypted(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwAccessTokenEncrypted = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetRequiredList(v []*string) *QaOpenapiSdkPostCommonResponseData {
	s.RequiredList = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetQaExtra(v *QaOpenapiSdkPostCommonResponseDataQaExtra) *QaOpenapiSdkPostCommonResponseData {
	s.QaExtra = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwUserId(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwUserId = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetDefaultStruct(v *QaOpenapiSdkPostCommonResponseDataDefaultStruct) *QaOpenapiSdkPostCommonResponseData {
	s.DefaultStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwClientIP(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwClientIP = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetDefaultMap(v map[string]*string) *QaOpenapiSdkPostCommonResponseData {
	s.DefaultMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetOptionalQuery(v string) *QaOpenapiSdkPostCommonResponseData {
	s.OptionalQuery = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetOptionalStruct(v *QaOpenapiSdkPostCommonResponseDataOptionalStruct) *QaOpenapiSdkPostCommonResponseData {
	s.OptionalStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetOptionalList(v []*string) *QaOpenapiSdkPostCommonResponseData {
	s.OptionalList = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetDefaultList(v []*string) *QaOpenapiSdkPostCommonResponseData {
	s.DefaultList = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetV3GwOpenId(v string) *QaOpenapiSdkPostCommonResponseData {
	s.V3GwOpenId = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetRawUri(v string) *QaOpenapiSdkPostCommonResponseData {
	s.RawUri = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetRequiredStruct(v *QaOpenapiSdkPostCommonResponseDataRequiredStruct) *QaOpenapiSdkPostCommonResponseData {
	s.RequiredStruct = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetOptionalMap(v map[string]*string) *QaOpenapiSdkPostCommonResponseData {
	s.OptionalMap = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseData) SetDefaultQuery(v string) *QaOpenapiSdkPostCommonResponseData {
	s.DefaultQuery = &v
	return s
}

type QaOpenapiSdkPostCommonResponseDataDefaultStruct struct {
	DefaultInt16           *int     `json:"default_int16,omitempty" xml:"default_int16,omitempty"`
	DefaultInt32           *int32   `json:"default_int32,omitempty" xml:"default_int32,omitempty"`
	DefaultInt64           *int64   `json:"default_int64,omitempty" xml:"default_int64,omitempty"`
	DefaultFloat64         *float64 `json:"default_float64,omitempty" xml:"default_float64,omitempty"`
	DefaultString          *string  `json:"default_string,omitempty" xml:"default_string,omitempty"`
	DefaultByteSliceString []byte   `json:"default_byte_slice_string,omitempty" xml:"default_byte_slice_string,omitempty"`
	DefaultBool            *bool    `json:"default_bool,omitempty" xml:"default_bool,omitempty"`
	DefaultInt8            *int     `json:"default_int8,omitempty" xml:"default_int8,omitempty"`
}

func (s QaOpenapiSdkPostCommonResponseDataDefaultStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponseDataDefaultStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultInt16(v int) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultInt32(v int32) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultInt32 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultInt64(v int64) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultInt64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultFloat64(v float64) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultFloat64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultString(v string) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultByteSliceString(v []byte) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultBool(v bool) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataDefaultStruct) SetDefaultInt8(v int) *QaOpenapiSdkPostCommonResponseDataDefaultStruct {
	s.DefaultInt8 = &v
	return s
}

type QaOpenapiSdkPostCommonResponseDataOptionalStruct struct {
	OptionalInt16           *int     `json:"optional_int16,omitempty" xml:"optional_int16,omitempty"`
	OptionalInt32           *int32   `json:"optional_int32,omitempty" xml:"optional_int32,omitempty"`
	OptionalInt64           *int64   `json:"optional_int64,omitempty" xml:"optional_int64,omitempty"`
	OptionalFloat64         *float64 `json:"optional_float64,omitempty" xml:"optional_float64,omitempty"`
	OptionalString          *string  `json:"optional_string,omitempty" xml:"optional_string,omitempty"`
	OptionalByteSliceString []byte   `json:"optional_byte_slice_string,omitempty" xml:"optional_byte_slice_string,omitempty"`
	OptionalBool            *bool    `json:"optional_bool,omitempty" xml:"optional_bool,omitempty"`
	OptionalInt8            *int     `json:"optional_int8,omitempty" xml:"optional_int8,omitempty"`
}

func (s QaOpenapiSdkPostCommonResponseDataOptionalStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponseDataOptionalStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalInt16(v int) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalInt32(v int32) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalInt32 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalInt64(v int64) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalInt64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalFloat64(v float64) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalFloat64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalString(v string) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalByteSliceString(v []byte) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalBool(v bool) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataOptionalStruct) SetOptionalInt8(v int) *QaOpenapiSdkPostCommonResponseDataOptionalStruct {
	s.OptionalInt8 = &v
	return s
}

type QaOpenapiSdkPostCommonResponseDataQaExtra struct {
	TimeOutSecond    *int64  `json:"time_out_second,omitempty" xml:"time_out_second,omitempty" require:"true"`
	IsPanic          *bool   `json:"is_panic,omitempty" xml:"is_panic,omitempty" require:"true"`
	ProductSource    *int    `json:"product_source,omitempty" xml:"product_source,omitempty" require:"true"`
	ErrNo            *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	IsDeleteRespData *bool   `json:"is_delete_resp_data,omitempty" xml:"is_delete_resp_data,omitempty" require:"true"`
	SceneDesc        *string `json:"scene_desc,omitempty" xml:"scene_desc,omitempty" require:"true"`
}

func (s QaOpenapiSdkPostCommonResponseDataQaExtra) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponseDataQaExtra) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetTimeOutSecond(v int64) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.TimeOutSecond = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetIsPanic(v bool) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.IsPanic = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetProductSource(v int) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.ProductSource = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetErrNo(v int32) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.ErrNo = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetIsDeleteRespData(v bool) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.IsDeleteRespData = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataQaExtra) SetSceneDesc(v string) *QaOpenapiSdkPostCommonResponseDataQaExtra {
	s.SceneDesc = &v
	return s
}

type QaOpenapiSdkPostCommonResponseDataRequiredStruct struct {
	RequiredString          *string  `json:"required_string,omitempty" xml:"required_string,omitempty" require:"true"`
	RequiredByteSliceString []byte   `json:"required_byte_slice_string,omitempty" xml:"required_byte_slice_string,omitempty" require:"true"`
	RequiredBool            *bool    `json:"required_bool,omitempty" xml:"required_bool,omitempty" require:"true"`
	RequiredInt8            *int     `json:"required_int8,omitempty" xml:"required_int8,omitempty" require:"true"`
	RequiredInt16           *int     `json:"required_int16,omitempty" xml:"required_int16,omitempty" require:"true"`
	RequiredInt32           *int32   `json:"required_int32,omitempty" xml:"required_int32,omitempty" require:"true"`
	RequiredInt64           *int64   `json:"required_int64,omitempty" xml:"required_int64,omitempty" require:"true"`
	RequiredFloat64         *float64 `json:"required_float64,omitempty" xml:"required_float64,omitempty" require:"true"`
}

func (s QaOpenapiSdkPostCommonResponseDataRequiredStruct) String() string {
	return tea.Prettify(s)
}

func (s QaOpenapiSdkPostCommonResponseDataRequiredStruct) GoString() string {
	return s.String()
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredString(v string) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredString = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredByteSliceString(v []byte) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredByteSliceString = v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredBool(v bool) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredBool = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredInt8(v int) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredInt8 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredInt16(v int) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredInt16 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredInt32(v int32) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredInt32 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredInt64(v int64) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredInt64 = &v
	return s
}

func (s *QaOpenapiSdkPostCommonResponseDataRequiredStruct) SetRequiredFloat64(v float64) *QaOpenapiSdkPostCommonResponseDataRequiredStruct {
	s.RequiredFloat64 = &v
	return s
}

type QrcodeCreateRequest struct {
	AccessToken  *string                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppName      *string                        `json:"app_name,omitempty" xml:"app_name,omitempty"`
	Background   *QrcodeCreateRequestBackground `json:"background,omitempty" xml:"background,omitempty"`
	Path         *string                        `json:"path,omitempty" xml:"path,omitempty"`
	LineColor    *QrcodeCreateRequestLineColor  `json:"line_color,omitempty" xml:"line_color,omitempty"`
	Header       map[string]*string             `json:"header,omitempty" xml:"header,omitempty"`
	Appid        *string                        `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Width        *int32                         `json:"width,omitempty" xml:"width,omitempty"`
	SetIcon      *bool                          `json:"set_icon,omitempty" xml:"set_icon,omitempty"`
	IsCircleCode *bool                          `json:"is_circle_code,omitempty" xml:"is_circle_code,omitempty"`
}

func (s QrcodeCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s QrcodeCreateRequest) GoString() string {
	return s.String()
}

func (s *QrcodeCreateRequest) SetAccessToken(v string) *QrcodeCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *QrcodeCreateRequest) SetAppName(v string) *QrcodeCreateRequest {
	s.AppName = &v
	return s
}

func (s *QrcodeCreateRequest) SetBackground(v *QrcodeCreateRequestBackground) *QrcodeCreateRequest {
	s.Background = v
	return s
}

func (s *QrcodeCreateRequest) SetPath(v string) *QrcodeCreateRequest {
	s.Path = &v
	return s
}

func (s *QrcodeCreateRequest) SetLineColor(v *QrcodeCreateRequestLineColor) *QrcodeCreateRequest {
	s.LineColor = v
	return s
}

func (s *QrcodeCreateRequest) SetHeader(v map[string]*string) *QrcodeCreateRequest {
	s.Header = v
	return s
}

func (s *QrcodeCreateRequest) SetAppid(v string) *QrcodeCreateRequest {
	s.Appid = &v
	return s
}

func (s *QrcodeCreateRequest) SetWidth(v int32) *QrcodeCreateRequest {
	s.Width = &v
	return s
}

func (s *QrcodeCreateRequest) SetSetIcon(v bool) *QrcodeCreateRequest {
	s.SetIcon = &v
	return s
}

func (s *QrcodeCreateRequest) SetIsCircleCode(v bool) *QrcodeCreateRequest {
	s.IsCircleCode = &v
	return s
}

type QrcodeCreateRequestBackground struct {
	B *int32 `json:"b,omitempty" xml:"b,omitempty"`
	R *int32 `json:"r,omitempty" xml:"r,omitempty"`
	G *int32 `json:"g,omitempty" xml:"g,omitempty"`
}

func (s QrcodeCreateRequestBackground) String() string {
	return tea.Prettify(s)
}

func (s QrcodeCreateRequestBackground) GoString() string {
	return s.String()
}

func (s *QrcodeCreateRequestBackground) SetB(v int32) *QrcodeCreateRequestBackground {
	s.B = &v
	return s
}

func (s *QrcodeCreateRequestBackground) SetR(v int32) *QrcodeCreateRequestBackground {
	s.R = &v
	return s
}

func (s *QrcodeCreateRequestBackground) SetG(v int32) *QrcodeCreateRequestBackground {
	s.G = &v
	return s
}

type QrcodeCreateRequestLineColor struct {
	G *int32 `json:"g,omitempty" xml:"g,omitempty"`
	B *int32 `json:"b,omitempty" xml:"b,omitempty"`
	R *int32 `json:"r,omitempty" xml:"r,omitempty"`
}

func (s QrcodeCreateRequestLineColor) String() string {
	return tea.Prettify(s)
}

func (s QrcodeCreateRequestLineColor) GoString() string {
	return s.String()
}

func (s *QrcodeCreateRequestLineColor) SetG(v int32) *QrcodeCreateRequestLineColor {
	s.G = &v
	return s
}

func (s *QrcodeCreateRequestLineColor) SetB(v int32) *QrcodeCreateRequestLineColor {
	s.B = &v
	return s
}

func (s *QrcodeCreateRequestLineColor) SetR(v int32) *QrcodeCreateRequestLineColor {
	s.R = &v
	return s
}

type QrcodeCreateResponse struct {
	ErrNo  *int32                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QrcodeCreateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QrcodeCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s QrcodeCreateResponse) GoString() string {
	return s.String()
}

func (s *QrcodeCreateResponse) SetErrNo(v int32) *QrcodeCreateResponse {
	s.ErrNo = &v
	return s
}

func (s *QrcodeCreateResponse) SetErrMsg(v string) *QrcodeCreateResponse {
	s.ErrMsg = &v
	return s
}

func (s *QrcodeCreateResponse) SetLogId(v string) *QrcodeCreateResponse {
	s.LogId = &v
	return s
}

func (s *QrcodeCreateResponse) SetData(v *QrcodeCreateResponseData) *QrcodeCreateResponse {
	s.Data = v
	return s
}

type QrcodeCreateResponseData struct {
	Img *string `json:"img,omitempty" xml:"img,omitempty"`
}

func (s QrcodeCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s QrcodeCreateResponseData) GoString() string {
	return s.String()
}

func (s *QrcodeCreateResponseData) SetImg(v string) *QrcodeCreateResponseData {
	s.Img = &v
	return s
}

type QualSearchRequest struct {
	AccessToken *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *int64                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Data        *QualSearchRequestData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Header      map[string]*string     `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QualSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s QualSearchRequest) GoString() string {
	return s.String()
}

func (s *QualSearchRequest) SetAccessToken(v string) *QualSearchRequest {
	s.AccessToken = &v
	return s
}

func (s *QualSearchRequest) SetAccountId(v int64) *QualSearchRequest {
	s.AccountId = &v
	return s
}

func (s *QualSearchRequest) SetData(v *QualSearchRequestData) *QualSearchRequest {
	s.Data = v
	return s
}

func (s *QualSearchRequest) SetHeader(v map[string]*string) *QualSearchRequest {
	s.Header = v
	return s
}

type QualSearchRequestData struct {
	PoiIds                 []*int64                         `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
	QualIds                []*int64                         `json:"qual_ids,omitempty" xml:"qual_ids,omitempty" type:"Repeated"`
	QualName               *string                          `json:"qual_name,omitempty" xml:"qual_name,omitempty"`
	AccurateLifeAccountIds []*int64                         `json:"accurate_life_account_ids,omitempty" xml:"accurate_life_account_ids,omitempty" type:"Repeated"`
	DataAccess             *QualSearchRequestDataDataAccess `json:"data_access,omitempty" xml:"data_access,omitempty"`
	PageIndex              *int32                           `json:"page_index,omitempty" xml:"page_index,omitempty"`
	PageSize               *int32                           `json:"page_size,omitempty" xml:"page_size,omitempty"`
	ParentLifeAccountIds   []*int64                         `json:"parent_life_account_ids,omitempty" xml:"parent_life_account_ids,omitempty" type:"Repeated"`
}

func (s QualSearchRequestData) String() string {
	return tea.Prettify(s)
}

func (s QualSearchRequestData) GoString() string {
	return s.String()
}

func (s *QualSearchRequestData) SetPoiIds(v []*int64) *QualSearchRequestData {
	s.PoiIds = v
	return s
}

func (s *QualSearchRequestData) SetQualIds(v []*int64) *QualSearchRequestData {
	s.QualIds = v
	return s
}

func (s *QualSearchRequestData) SetQualName(v string) *QualSearchRequestData {
	s.QualName = &v
	return s
}

func (s *QualSearchRequestData) SetAccurateLifeAccountIds(v []*int64) *QualSearchRequestData {
	s.AccurateLifeAccountIds = v
	return s
}

func (s *QualSearchRequestData) SetDataAccess(v *QualSearchRequestDataDataAccess) *QualSearchRequestData {
	s.DataAccess = v
	return s
}

func (s *QualSearchRequestData) SetPageIndex(v int32) *QualSearchRequestData {
	s.PageIndex = &v
	return s
}

func (s *QualSearchRequestData) SetPageSize(v int32) *QualSearchRequestData {
	s.PageSize = &v
	return s
}

func (s *QualSearchRequestData) SetParentLifeAccountIds(v []*int64) *QualSearchRequestData {
	s.ParentLifeAccountIds = v
	return s
}

type QualSearchRequestDataDataAccess struct {
	NeedEffectiveQual    *bool `json:"need_effective_qual,omitempty" xml:"need_effective_qual,omitempty"`
	OnlyStoreAccountQual *bool `json:"only_store_account_qual,omitempty" xml:"only_store_account_qual,omitempty"`
}

func (s QualSearchRequestDataDataAccess) String() string {
	return tea.Prettify(s)
}

func (s QualSearchRequestDataDataAccess) GoString() string {
	return s.String()
}

func (s *QualSearchRequestDataDataAccess) SetNeedEffectiveQual(v bool) *QualSearchRequestDataDataAccess {
	s.NeedEffectiveQual = &v
	return s
}

func (s *QualSearchRequestDataDataAccess) SetOnlyStoreAccountQual(v bool) *QualSearchRequestDataDataAccess {
	s.OnlyStoreAccountQual = &v
	return s
}

type QualSearchResponse struct {
	Data  *QualSearchResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *QualSearchResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s QualSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s QualSearchResponse) GoString() string {
	return s.String()
}

func (s *QualSearchResponse) SetData(v *QualSearchResponseData) *QualSearchResponse {
	s.Data = v
	return s
}

func (s *QualSearchResponse) SetExtra(v *QualSearchResponseExtra) *QualSearchResponse {
	s.Extra = v
	return s
}

type QualSearchResponseData struct {
	Total         *int32                                   `json:"total,omitempty" xml:"total,omitempty"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Page          *int32                                   `json:"page,omitempty" xml:"page,omitempty"`
	PageSize      *int32                                   `json:"page_size,omitempty" xml:"page_size,omitempty"`
	QualDtoList   []*QualSearchResponseDataQualDtoListItem `json:"qual_dto_list,omitempty" xml:"qual_dto_list,omitempty" type:"Repeated"`
}

func (s QualSearchResponseData) String() string {
	return tea.Prettify(s)
}

func (s QualSearchResponseData) GoString() string {
	return s.String()
}

func (s *QualSearchResponseData) SetTotal(v int32) *QualSearchResponseData {
	s.Total = &v
	return s
}

func (s *QualSearchResponseData) SetGwErrorCode(v int32) *QualSearchResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QualSearchResponseData) SetGwDescription(v string) *QualSearchResponseData {
	s.GwDescription = &v
	return s
}

func (s *QualSearchResponseData) SetPage(v int32) *QualSearchResponseData {
	s.Page = &v
	return s
}

func (s *QualSearchResponseData) SetPageSize(v int32) *QualSearchResponseData {
	s.PageSize = &v
	return s
}

func (s *QualSearchResponseData) SetQualDtoList(v []*QualSearchResponseDataQualDtoListItem) *QualSearchResponseData {
	s.QualDtoList = v
	return s
}

type QualSearchResponseDataQualDtoListItem struct {
	EffectiveDate     *string                                                `json:"effective_date,omitempty" xml:"effective_date,omitempty" require:"true"`
	Status            *int                                                   `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	QualificationType *int32                                                 `json:"qualification_type,omitempty" xml:"qualification_type,omitempty" require:"true"`
	QualCategory      *int                                                   `json:"qual_category,omitempty" xml:"qual_category,omitempty"`
	CreateTime        *string                                                `json:"create_time,omitempty" xml:"create_time,omitempty"`
	StatusChangedTime *string                                                `json:"status_changed_time,omitempty" xml:"status_changed_time,omitempty" require:"true"`
	LifeAccountId     *int64                                                 `json:"life_account_id,omitempty" xml:"life_account_id,omitempty"`
	QualificationId   *int64                                                 `json:"qualification_id,omitempty" xml:"qualification_id,omitempty" require:"true"`
	QualTypeName      *string                                                `json:"qual_type_name,omitempty" xml:"qual_type_name,omitempty"`
	RejectReason      *string                                                `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Attributes        []*QualSearchResponseDataQualDtoListItemAttributesItem `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
}

func (s QualSearchResponseDataQualDtoListItem) String() string {
	return tea.Prettify(s)
}

func (s QualSearchResponseDataQualDtoListItem) GoString() string {
	return s.String()
}

func (s *QualSearchResponseDataQualDtoListItem) SetEffectiveDate(v string) *QualSearchResponseDataQualDtoListItem {
	s.EffectiveDate = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetStatus(v int) *QualSearchResponseDataQualDtoListItem {
	s.Status = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetQualificationType(v int32) *QualSearchResponseDataQualDtoListItem {
	s.QualificationType = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetQualCategory(v int) *QualSearchResponseDataQualDtoListItem {
	s.QualCategory = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetCreateTime(v string) *QualSearchResponseDataQualDtoListItem {
	s.CreateTime = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetStatusChangedTime(v string) *QualSearchResponseDataQualDtoListItem {
	s.StatusChangedTime = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetLifeAccountId(v int64) *QualSearchResponseDataQualDtoListItem {
	s.LifeAccountId = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetQualificationId(v int64) *QualSearchResponseDataQualDtoListItem {
	s.QualificationId = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetQualTypeName(v string) *QualSearchResponseDataQualDtoListItem {
	s.QualTypeName = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetRejectReason(v string) *QualSearchResponseDataQualDtoListItem {
	s.RejectReason = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItem) SetAttributes(v []*QualSearchResponseDataQualDtoListItemAttributesItem) *QualSearchResponseDataQualDtoListItem {
	s.Attributes = v
	return s
}

type QualSearchResponseDataQualDtoListItemAttributesItem struct {
	AttributeKey    *string   `json:"attribute_key,omitempty" xml:"attribute_key,omitempty" require:"true"`
	AttributeValues []*string `json:"attribute_values,omitempty" xml:"attribute_values,omitempty" require:"true" type:"Repeated"`
}

func (s QualSearchResponseDataQualDtoListItemAttributesItem) String() string {
	return tea.Prettify(s)
}

func (s QualSearchResponseDataQualDtoListItemAttributesItem) GoString() string {
	return s.String()
}

func (s *QualSearchResponseDataQualDtoListItemAttributesItem) SetAttributeKey(v string) *QualSearchResponseDataQualDtoListItemAttributesItem {
	s.AttributeKey = &v
	return s
}

func (s *QualSearchResponseDataQualDtoListItemAttributesItem) SetAttributeValues(v []*string) *QualSearchResponseDataQualDtoListItemAttributesItem {
	s.AttributeValues = v
	return s
}

type QualSearchResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s QualSearchResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QualSearchResponseExtra) GoString() string {
	return s.String()
}

func (s *QualSearchResponseExtra) SetNow(v int64) *QualSearchResponseExtra {
	s.Now = &v
	return s
}

func (s *QualSearchResponseExtra) SetSubDescription(v string) *QualSearchResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QualSearchResponseExtra) SetSubErrorCode(v int32) *QualSearchResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QualSearchResponseExtra) SetDescription(v string) *QualSearchResponseExtra {
	s.Description = &v
	return s
}

func (s *QualSearchResponseExtra) SetErrorCode(v int32) *QualSearchResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QualSearchResponseExtra) SetLogid(v string) *QualSearchResponseExtra {
	s.Logid = &v
	return s
}

type QueryActivityMetaDataRequest struct {
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	CouponMetaId *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CouponName   *string            `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	PageNum      *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize     *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	SendScene    *int               `json:"send_scene,omitempty" xml:"send_scene,omitempty" require:"true"`
	DiscountType *int               `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	Status       *int               `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s QueryActivityMetaDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityMetaDataRequest) GoString() string {
	return s.String()
}

func (s *QueryActivityMetaDataRequest) SetOpenId(v string) *QueryActivityMetaDataRequest {
	s.OpenId = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetCouponMetaId(v string) *QueryActivityMetaDataRequest {
	s.CouponMetaId = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetHeader(v map[string]*string) *QueryActivityMetaDataRequest {
	s.Header = v
	return s
}

func (s *QueryActivityMetaDataRequest) SetAccessToken(v string) *QueryActivityMetaDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetCouponName(v string) *QueryActivityMetaDataRequest {
	s.CouponName = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetPageNum(v int32) *QueryActivityMetaDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetPageSize(v int32) *QueryActivityMetaDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetSendScene(v int) *QueryActivityMetaDataRequest {
	s.SendScene = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetDiscountType(v int) *QueryActivityMetaDataRequest {
	s.DiscountType = &v
	return s
}

func (s *QueryActivityMetaDataRequest) SetStatus(v int) *QueryActivityMetaDataRequest {
	s.Status = &v
	return s
}

type QueryActivityMetaDataResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryActivityMetaDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryActivityMetaDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityMetaDataResponse) GoString() string {
	return s.String()
}

func (s *QueryActivityMetaDataResponse) SetErrMsg(v string) *QueryActivityMetaDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryActivityMetaDataResponse) SetLogId(v string) *QueryActivityMetaDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryActivityMetaDataResponse) SetData(v *QueryActivityMetaDataResponseData) *QueryActivityMetaDataResponse {
	s.Data = v
	return s
}

func (s *QueryActivityMetaDataResponse) SetErrNo(v int32) *QueryActivityMetaDataResponse {
	s.ErrNo = &v
	return s
}

type QueryActivityMetaDataResponseData struct {
	ActivityList []*QueryActivityMetaDataResponseDataActivityListItem `json:"activity_list,omitempty" xml:"activity_list,omitempty" require:"true" type:"Repeated"`
	Total        *int64                                               `json:"total,omitempty" xml:"total,omitempty" require:"true"`
}

func (s QueryActivityMetaDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityMetaDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryActivityMetaDataResponseData) SetActivityList(v []*QueryActivityMetaDataResponseDataActivityListItem) *QueryActivityMetaDataResponseData {
	s.ActivityList = v
	return s
}

func (s *QueryActivityMetaDataResponseData) SetTotal(v int64) *QueryActivityMetaDataResponseData {
	s.Total = &v
	return s
}

type QueryActivityMetaDataResponseDataActivityListItem struct {
	AppIcon             *string `json:"app_icon,omitempty" xml:"app_icon,omitempty" require:"true"`
	CouponMetaName      *string `json:"coupon_meta_name,omitempty" xml:"coupon_meta_name,omitempty" require:"true"`
	ReceiveBeginTime    *int64  `json:"receive_begin_time,omitempty" xml:"receive_begin_time,omitempty" require:"true"`
	ConsumeDesc         *string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty" require:"true"`
	ActivityStatus      *int    `json:"activity_status,omitempty" xml:"activity_status,omitempty" require:"true"`
	ShareFissionDesc    *string `json:"share_fission_desc,omitempty" xml:"share_fission_desc,omitempty" require:"true"`
	AppId               *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CouponMetaId        *string `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	AppName             *string `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	DiscountAmount      *int64  `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	ActivityId          *string `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	TotalStock          *int64  `json:"total_stock,omitempty" xml:"total_stock,omitempty" require:"true"`
	DiscountTypeName    *string `json:"discount_type_name,omitempty" xml:"discount_type_name,omitempty" require:"true"`
	ConsumeTimeDesc     *string `json:"consume_time_desc,omitempty" xml:"consume_time_desc,omitempty" require:"true"`
	SupportShareFission *bool   `json:"support_share_fission,omitempty" xml:"support_share_fission,omitempty" require:"true"`
	RemainStock         *int64  `json:"remain_stock,omitempty" xml:"remain_stock,omitempty" require:"true"`
	ReceiveEndTime      *int64  `json:"receive_end_time,omitempty" xml:"receive_end_time,omitempty" require:"true"`
	CouponIcon          *string `json:"coupon_icon,omitempty" xml:"coupon_icon,omitempty" require:"true"`
	DiscountType        *int    `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
}

func (s QueryActivityMetaDataResponseDataActivityListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityMetaDataResponseDataActivityListItem) GoString() string {
	return s.String()
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetAppIcon(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.AppIcon = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetCouponMetaName(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.CouponMetaName = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetReceiveBeginTime(v int64) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ReceiveBeginTime = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetConsumeDesc(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ConsumeDesc = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetActivityStatus(v int) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ActivityStatus = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetShareFissionDesc(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ShareFissionDesc = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetAppId(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.AppId = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetCouponMetaId(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.CouponMetaId = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetAppName(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.AppName = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetDiscountAmount(v int64) *QueryActivityMetaDataResponseDataActivityListItem {
	s.DiscountAmount = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetActivityId(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ActivityId = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetTotalStock(v int64) *QueryActivityMetaDataResponseDataActivityListItem {
	s.TotalStock = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetDiscountTypeName(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.DiscountTypeName = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetConsumeTimeDesc(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ConsumeTimeDesc = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetSupportShareFission(v bool) *QueryActivityMetaDataResponseDataActivityListItem {
	s.SupportShareFission = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetRemainStock(v int64) *QueryActivityMetaDataResponseDataActivityListItem {
	s.RemainStock = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetReceiveEndTime(v int64) *QueryActivityMetaDataResponseDataActivityListItem {
	s.ReceiveEndTime = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetCouponIcon(v string) *QueryActivityMetaDataResponseDataActivityListItem {
	s.CouponIcon = &v
	return s
}

func (s *QueryActivityMetaDataResponseDataActivityListItem) SetDiscountType(v int) *QueryActivityMetaDataResponseDataActivityListItem {
	s.DiscountType = &v
	return s
}

type QueryActivityUserCompletionStatusRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ActivityId  *int64             `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	TaskIdList  []*int64           `json:"task_id_list,omitempty" xml:"task_id_list,omitempty" type:"Repeated"`
}

func (s QueryActivityUserCompletionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityUserCompletionStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryActivityUserCompletionStatusRequest) SetHeader(v map[string]*string) *QueryActivityUserCompletionStatusRequest {
	s.Header = v
	return s
}

func (s *QueryActivityUserCompletionStatusRequest) SetAccessToken(v string) *QueryActivityUserCompletionStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryActivityUserCompletionStatusRequest) SetActivityId(v int64) *QueryActivityUserCompletionStatusRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryActivityUserCompletionStatusRequest) SetTaskIdList(v []*int64) *QueryActivityUserCompletionStatusRequest {
	s.TaskIdList = v
	return s
}

type QueryActivityUserCompletionStatusResponse struct {
	PostingVideoBindCompeleteInfoMap map[int64]*QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue `json:"posting_video_bind_compelete_info_map,omitempty" xml:"posting_video_bind_compelete_info_map,omitempty"`
	TaskCompleteStatusMap            map[int64]*bool                                                                           `json:"task_complete_status_map,omitempty" xml:"task_complete_status_map,omitempty" require:"true"`
}

func (s QueryActivityUserCompletionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityUserCompletionStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryActivityUserCompletionStatusResponse) SetPostingVideoBindCompeleteInfoMap(v map[int64]*QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) *QueryActivityUserCompletionStatusResponse {
	s.PostingVideoBindCompeleteInfoMap = v
	return s
}

func (s *QueryActivityUserCompletionStatusResponse) SetTaskCompleteStatusMap(v map[int64]*bool) *QueryActivityUserCompletionStatusResponse {
	s.TaskCompleteStatusMap = v
	return s
}

type QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue struct {
	PostingNumStageMap map[int64]*bool `json:"posting_num_stage_map,omitempty" xml:"posting_num_stage_map,omitempty"`
	CommentStageMap    map[int64]*bool `json:"comment_stage_map,omitempty" xml:"comment_stage_map,omitempty"`
	DiggStageMap       map[int64]*bool `json:"digg_stage_map,omitempty" xml:"digg_stage_map,omitempty"`
	PlayVideoStageMap  map[int64]*bool `json:"play_video_stage_map,omitempty" xml:"play_video_stage_map,omitempty"`
}

func (s QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) GoString() string {
	return s.String()
}

func (s *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) SetPostingNumStageMap(v map[int64]*bool) *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue {
	s.PostingNumStageMap = v
	return s
}

func (s *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) SetCommentStageMap(v map[int64]*bool) *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue {
	s.CommentStageMap = v
	return s
}

func (s *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) SetDiggStageMap(v map[int64]*bool) *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue {
	s.DiggStageMap = v
	return s
}

func (s *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue) SetPlayVideoStageMap(v map[int64]*bool) *QueryActivityUserCompletionStatusResponsePostingVideoBindCompeleteInfoMapValue {
	s.PlayVideoStageMap = v
	return s
}

type QueryAdSettlementListRequest struct {
	Month       *string            `json:"month,omitempty" xml:"month,omitempty"`
	Status      *int32             `json:"status,omitempty" xml:"status,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryAdSettlementListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdSettlementListRequest) GoString() string {
	return s.String()
}

func (s *QueryAdSettlementListRequest) SetMonth(v string) *QueryAdSettlementListRequest {
	s.Month = &v
	return s
}

func (s *QueryAdSettlementListRequest) SetStatus(v int32) *QueryAdSettlementListRequest {
	s.Status = &v
	return s
}

func (s *QueryAdSettlementListRequest) SetHeader(v map[string]*string) *QueryAdSettlementListRequest {
	s.Header = v
	return s
}

func (s *QueryAdSettlementListRequest) SetAccessToken(v string) *QueryAdSettlementListRequest {
	s.AccessToken = &v
	return s
}

type QueryAdSettlementListResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryAdSettlementListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryAdSettlementListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdSettlementListResponse) GoString() string {
	return s.String()
}

func (s *QueryAdSettlementListResponse) SetLogId(v string) *QueryAdSettlementListResponse {
	s.LogId = &v
	return s
}

func (s *QueryAdSettlementListResponse) SetData(v *QueryAdSettlementListResponseData) *QueryAdSettlementListResponse {
	s.Data = v
	return s
}

func (s *QueryAdSettlementListResponse) SetErrNo(v int32) *QueryAdSettlementListResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryAdSettlementListResponse) SetErrMsg(v string) *QueryAdSettlementListResponse {
	s.ErrMsg = &v
	return s
}

type QueryAdSettlementListResponseData struct {
	SettlementList []*QueryAdSettlementListResponseDataSettlementListItem `json:"settlement_list,omitempty" xml:"settlement_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAdSettlementListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAdSettlementListResponseData) GoString() string {
	return s.String()
}

func (s *QueryAdSettlementListResponseData) SetSettlementList(v []*QueryAdSettlementListResponseDataSettlementListItem) *QueryAdSettlementListResponseData {
	s.SettlementList = v
	return s
}

type QueryAdSettlementListResponseDataSettlementListItem struct {
	SettlementTotalAmount *int64   `json:"settlement_total_amount,omitempty" xml:"settlement_total_amount,omitempty" require:"true"`
	Status                *int32   `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	SettlementSerial      *string  `json:"settlement_serial,omitempty" xml:"settlement_serial,omitempty" require:"true"`
	SettlementName        *string  `json:"settlement_name,omitempty" xml:"settlement_name,omitempty" require:"true"`
	SettlementPeriod      *string  `json:"settlement_period,omitempty" xml:"settlement_period,omitempty" require:"true"`
	TaxRate               *float64 `json:"tax_rate,omitempty" xml:"tax_rate,omitempty" require:"true"`
}

func (s QueryAdSettlementListResponseDataSettlementListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAdSettlementListResponseDataSettlementListItem) GoString() string {
	return s.String()
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetSettlementTotalAmount(v int64) *QueryAdSettlementListResponseDataSettlementListItem {
	s.SettlementTotalAmount = &v
	return s
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetStatus(v int32) *QueryAdSettlementListResponseDataSettlementListItem {
	s.Status = &v
	return s
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetSettlementSerial(v string) *QueryAdSettlementListResponseDataSettlementListItem {
	s.SettlementSerial = &v
	return s
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetSettlementName(v string) *QueryAdSettlementListResponseDataSettlementListItem {
	s.SettlementName = &v
	return s
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetSettlementPeriod(v string) *QueryAdSettlementListResponseDataSettlementListItem {
	s.SettlementPeriod = &v
	return s
}

func (s *QueryAdSettlementListResponseDataSettlementListItem) SetTaxRate(v float64) *QueryAdSettlementListResponseDataSettlementListItem {
	s.TaxRate = &v
	return s
}

type QueryAgencyVideoDailyDataRequest struct {
	PageSize              *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	AppId                 *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	BillingDate           *string            `json:"billing_date,omitempty" xml:"billing_date,omitempty" require:"true"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AgentId               *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	VideoPublishEndTime   *int64             `json:"video_publish_end_time,omitempty" xml:"video_publish_end_time,omitempty" require:"true"`
	PageNum               *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	VideoPublishStartTime *int64             `json:"video_publish_start_time,omitempty" xml:"video_publish_start_time,omitempty" require:"true"`
	DouyinId              *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty"`
}

func (s QueryAgencyVideoDailyDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoDailyDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoDailyDataRequest) SetPageSize(v int32) *QueryAgencyVideoDailyDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetAppId(v string) *QueryAgencyVideoDailyDataRequest {
	s.AppId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetBillingDate(v string) *QueryAgencyVideoDailyDataRequest {
	s.BillingDate = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetAccessToken(v string) *QueryAgencyVideoDailyDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetAgentId(v int64) *QueryAgencyVideoDailyDataRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetHeader(v map[string]*string) *QueryAgencyVideoDailyDataRequest {
	s.Header = v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetVideoPublishEndTime(v int64) *QueryAgencyVideoDailyDataRequest {
	s.VideoPublishEndTime = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetPageNum(v int32) *QueryAgencyVideoDailyDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetVideoPublishStartTime(v int64) *QueryAgencyVideoDailyDataRequest {
	s.VideoPublishStartTime = &v
	return s
}

func (s *QueryAgencyVideoDailyDataRequest) SetDouyinId(v string) *QueryAgencyVideoDailyDataRequest {
	s.DouyinId = &v
	return s
}

type QueryAgencyVideoDailyDataResponse struct {
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryAgencyVideoDailyDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryAgencyVideoDailyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoDailyDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoDailyDataResponse) SetErrNo(v int32) *QueryAgencyVideoDailyDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponse) SetErrMsg(v string) *QueryAgencyVideoDailyDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponse) SetLogId(v string) *QueryAgencyVideoDailyDataResponse {
	s.LogId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponse) SetData(v *QueryAgencyVideoDailyDataResponseData) *QueryAgencyVideoDailyDataResponse {
	s.Data = v
	return s
}

type QueryAgencyVideoDailyDataResponseData struct {
	Results []*QueryAgencyVideoDailyDataResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAgencyVideoDailyDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoDailyDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoDailyDataResponseData) SetResults(v []*QueryAgencyVideoDailyDataResponseDataResultsItem) *QueryAgencyVideoDailyDataResponseData {
	s.Results = v
	return s
}

type QueryAgencyVideoDailyDataResponseDataResultsItem struct {
	TaskName           *string `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	MicroAppTitle      *string `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty" require:"true"`
	AgentId            *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Gmv1d              *int64  `json:"gmv_1d,omitempty" xml:"gmv_1d,omitempty"`
	BillingRefundGmv1d *int64  `json:"billing_refund_gmv_1d,omitempty" xml:"billing_refund_gmv_1d,omitempty"`
	FeedAdShareCost1d  *int64  `json:"feed_ad_share_cost_1d,omitempty" xml:"feed_ad_share_cost_1d,omitempty"`
	Author             *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	Date               *string `json:"date,omitempty" xml:"date,omitempty"`
	ActiveCnt1d        *int64  `json:"active_cnt_1d,omitempty" xml:"active_cnt_1d,omitempty"`
	RefundGmv1d        *int64  `json:"refund_gmv_1d,omitempty" xml:"refund_gmv_1d,omitempty"`
	VideoTitle         *string `json:"video_title,omitempty" xml:"video_title,omitempty" require:"true"`
	BillingGmv1d       *int64  `json:"billing_gmv_1d,omitempty" xml:"billing_gmv_1d,omitempty"`
	VideoLink          *string `json:"video_link,omitempty" xml:"video_link,omitempty" require:"true"`
	DouyinId           *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	AdShareCost1d      *int64  `json:"ad_share_cost_1d,omitempty" xml:"ad_share_cost_1d,omitempty"`
	VideoId            *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	TalentProfit1d     *int64  `json:"talent_profit_1d,omitempty" xml:"talent_profit_1d,omitempty"`
	TaskId             *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	PublishTime        *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
}

func (s QueryAgencyVideoDailyDataResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoDailyDataResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetTaskName(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.TaskName = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetMicroAppTitle(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.MicroAppTitle = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetAgentId(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.AgentId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetGmv1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.Gmv1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetBillingRefundGmv1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.BillingRefundGmv1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetFeedAdShareCost1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.FeedAdShareCost1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetAuthor(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.Author = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetDate(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.Date = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetActiveCnt1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.ActiveCnt1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetRefundGmv1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.RefundGmv1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetVideoTitle(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.VideoTitle = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetBillingGmv1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.BillingGmv1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetVideoLink(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.VideoLink = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetDouyinId(v string) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.DouyinId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetAdShareCost1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.AdShareCost1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetVideoId(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.VideoId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetTalentProfit1d(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.TalentProfit1d = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetTaskId(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.TaskId = &v
	return s
}

func (s *QueryAgencyVideoDailyDataResponseDataResultsItem) SetPublishTime(v int64) *QueryAgencyVideoDailyDataResponseDataResultsItem {
	s.PublishTime = &v
	return s
}

type QueryAgencyVideoSumDataRequest struct {
	DouyinId              *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty"`
	AppId                 *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	AgentId               *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	PageNum               *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VideoPublishEndTime   *int64             `json:"video_publish_end_time,omitempty" xml:"video_publish_end_time,omitempty" require:"true"`
	VideoPublishStartTime *int64             `json:"video_publish_start_time,omitempty" xml:"video_publish_start_time,omitempty" require:"true"`
	PageSize              *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s QueryAgencyVideoSumDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoSumDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoSumDataRequest) SetDouyinId(v string) *QueryAgencyVideoSumDataRequest {
	s.DouyinId = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetAppId(v string) *QueryAgencyVideoSumDataRequest {
	s.AppId = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetAgentId(v int64) *QueryAgencyVideoSumDataRequest {
	s.AgentId = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetPageNum(v int32) *QueryAgencyVideoSumDataRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetHeader(v map[string]*string) *QueryAgencyVideoSumDataRequest {
	s.Header = v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetAccessToken(v string) *QueryAgencyVideoSumDataRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetVideoPublishEndTime(v int64) *QueryAgencyVideoSumDataRequest {
	s.VideoPublishEndTime = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetVideoPublishStartTime(v int64) *QueryAgencyVideoSumDataRequest {
	s.VideoPublishStartTime = &v
	return s
}

func (s *QueryAgencyVideoSumDataRequest) SetPageSize(v int32) *QueryAgencyVideoSumDataRequest {
	s.PageSize = &v
	return s
}

type QueryAgencyVideoSumDataResponse struct {
	Data   *QueryAgencyVideoSumDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryAgencyVideoSumDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoSumDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoSumDataResponse) SetData(v *QueryAgencyVideoSumDataResponseData) *QueryAgencyVideoSumDataResponse {
	s.Data = v
	return s
}

func (s *QueryAgencyVideoSumDataResponse) SetErrNo(v int32) *QueryAgencyVideoSumDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponse) SetErrMsg(v string) *QueryAgencyVideoSumDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponse) SetLogId(v string) *QueryAgencyVideoSumDataResponse {
	s.LogId = &v
	return s
}

type QueryAgencyVideoSumDataResponseData struct {
	Results []*QueryAgencyVideoSumDataResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAgencyVideoSumDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoSumDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoSumDataResponseData) SetResults(v []*QueryAgencyVideoSumDataResponseDataResultsItem) *QueryAgencyVideoSumDataResponseData {
	s.Results = v
	return s
}

type QueryAgencyVideoSumDataResponseDataResultsItem struct {
	GmvTd              *int64  `json:"gmv_td,omitempty" xml:"gmv_td,omitempty"`
	TalentProfitTd     *int64  `json:"talent_profit_td,omitempty" xml:"talent_profit_td,omitempty"`
	RefundGmvTd        *int64  `json:"refund_gmv_td,omitempty" xml:"refund_gmv_td,omitempty"`
	AgentId            *string `json:"agent_id,omitempty" xml:"agent_id,omitempty" require:"true"`
	Shares             *int64  `json:"shares,omitempty" xml:"shares,omitempty" require:"true"`
	Author             *string `json:"author,omitempty" xml:"author,omitempty" require:"true"`
	Likes              *int64  `json:"likes,omitempty" xml:"likes,omitempty" require:"true"`
	TaskId             *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	Clicks             *int64  `json:"clicks,omitempty" xml:"clicks,omitempty"`
	FeedAdShareCostTd  *int64  `json:"feed_ad_share_cost_td,omitempty" xml:"feed_ad_share_cost_td,omitempty"`
	AdShareCostTd      *int64  `json:"ad_share_cost_td,omitempty" xml:"ad_share_cost_td,omitempty"`
	VideoLink          *string `json:"video_link,omitempty" xml:"video_link,omitempty" require:"true"`
	TaskName           *string `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	VideoId            *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	BillingRefundGmvTd *int64  `json:"billing_refund_gmv_td,omitempty" xml:"billing_refund_gmv_td,omitempty"`
	BillingGmvTd       *int64  `json:"billing_gmv_td,omitempty" xml:"billing_gmv_td,omitempty"`
	ActiveCntTd        *int64  `json:"active_cnt_td,omitempty" xml:"active_cnt_td,omitempty"`
	DouyinId           *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	VideoViews         *int64  `json:"video_views,omitempty" xml:"video_views,omitempty" require:"true"`
	PublishTime        *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
	Comments           *int64  `json:"comments,omitempty" xml:"comments,omitempty" require:"true"`
	Date               *string `json:"date,omitempty" xml:"date,omitempty"`
	MicroAppTitle      *string `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty" require:"true"`
	VideoTitle         *string `json:"video_title,omitempty" xml:"video_title,omitempty" require:"true"`
}

func (s QueryAgencyVideoSumDataResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAgencyVideoSumDataResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetGmvTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.GmvTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetTalentProfitTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.TalentProfitTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetRefundGmvTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.RefundGmvTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetAgentId(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.AgentId = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetShares(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Shares = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetAuthor(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Author = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetLikes(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Likes = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetTaskId(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.TaskId = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetClicks(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Clicks = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetFeedAdShareCostTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.FeedAdShareCostTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetAdShareCostTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.AdShareCostTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetVideoLink(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.VideoLink = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetTaskName(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.TaskName = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetVideoId(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.VideoId = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetBillingRefundGmvTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.BillingRefundGmvTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetBillingGmvTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.BillingGmvTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetActiveCntTd(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.ActiveCntTd = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetDouyinId(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.DouyinId = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetVideoViews(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.VideoViews = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetPublishTime(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.PublishTime = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetComments(v int64) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Comments = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetDate(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.Date = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetMicroAppTitle(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.MicroAppTitle = &v
	return s
}

func (s *QueryAgencyVideoSumDataResponseDataResultsItem) SetVideoTitle(v string) *QueryAgencyVideoSumDataResponseDataResultsItem {
	s.VideoTitle = &v
	return s
}

type QueryAppTaskIdRequest struct {
	CreateStartTime *int64             `json:"create_start_time,omitempty" xml:"create_start_time,omitempty" require:"true"`
	CreateEndTime   *int64             `json:"create_end_time,omitempty" xml:"create_end_time,omitempty" require:"true"`
	TaskCategory    *int               `json:"task_category,omitempty" xml:"task_category,omitempty"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Appid           *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
}

func (s QueryAppTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTaskIdRequest) GoString() string {
	return s.String()
}

func (s *QueryAppTaskIdRequest) SetCreateStartTime(v int64) *QueryAppTaskIdRequest {
	s.CreateStartTime = &v
	return s
}

func (s *QueryAppTaskIdRequest) SetCreateEndTime(v int64) *QueryAppTaskIdRequest {
	s.CreateEndTime = &v
	return s
}

func (s *QueryAppTaskIdRequest) SetTaskCategory(v int) *QueryAppTaskIdRequest {
	s.TaskCategory = &v
	return s
}

func (s *QueryAppTaskIdRequest) SetHeader(v map[string]*string) *QueryAppTaskIdRequest {
	s.Header = v
	return s
}

func (s *QueryAppTaskIdRequest) SetAccessToken(v string) *QueryAppTaskIdRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryAppTaskIdRequest) SetAppid(v string) *QueryAppTaskIdRequest {
	s.Appid = &v
	return s
}

type QueryAppTaskIdResponse struct {
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryAppTaskIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s QueryAppTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTaskIdResponse) GoString() string {
	return s.String()
}

func (s *QueryAppTaskIdResponse) SetLogId(v string) *QueryAppTaskIdResponse {
	s.LogId = &v
	return s
}

func (s *QueryAppTaskIdResponse) SetData(v *QueryAppTaskIdResponseData) *QueryAppTaskIdResponse {
	s.Data = v
	return s
}

func (s *QueryAppTaskIdResponse) SetErrNo(v int32) *QueryAppTaskIdResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryAppTaskIdResponse) SetErrMsg(v string) *QueryAppTaskIdResponse {
	s.ErrMsg = &v
	return s
}

type QueryAppTaskIdResponseData struct {
	TaskIds []*int64 `json:"task_ids,omitempty" xml:"task_ids,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppTaskIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTaskIdResponseData) GoString() string {
	return s.String()
}

func (s *QueryAppTaskIdResponseData) SetTaskIds(v []*int64) *QueryAppTaskIdResponseData {
	s.TaskIds = v
	return s
}

type QueryAppTestRelationRequest struct {
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryAppTestRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTestRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryAppTestRelationRequest) SetType(v string) *QueryAppTestRelationRequest {
	s.Type = &v
	return s
}

func (s *QueryAppTestRelationRequest) SetHeader(v map[string]*string) *QueryAppTestRelationRequest {
	s.Header = v
	return s
}

func (s *QueryAppTestRelationRequest) SetAccessToken(v string) *QueryAppTestRelationRequest {
	s.AccessToken = &v
	return s
}

type QueryAppTestRelationResponse struct {
	Data      *QueryAppTestRelationResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra     *QueryAppTestRelationResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	RefIdList []*string                          `json:"ref_id_list,omitempty" xml:"ref_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryAppTestRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTestRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryAppTestRelationResponse) SetData(v *QueryAppTestRelationResponseData) *QueryAppTestRelationResponse {
	s.Data = v
	return s
}

func (s *QueryAppTestRelationResponse) SetExtra(v *QueryAppTestRelationResponseExtra) *QueryAppTestRelationResponse {
	s.Extra = v
	return s
}

func (s *QueryAppTestRelationResponse) SetRefIdList(v []*string) *QueryAppTestRelationResponse {
	s.RefIdList = v
	return s
}

type QueryAppTestRelationResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s QueryAppTestRelationResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTestRelationResponseData) GoString() string {
	return s.String()
}

func (s *QueryAppTestRelationResponseData) SetGwErrorCode(v int32) *QueryAppTestRelationResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QueryAppTestRelationResponseData) SetGwDescription(v string) *QueryAppTestRelationResponseData {
	s.GwDescription = &v
	return s
}

type QueryAppTestRelationResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s QueryAppTestRelationResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryAppTestRelationResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryAppTestRelationResponseExtra) SetErrorCode(v int32) *QueryAppTestRelationResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QueryAppTestRelationResponseExtra) SetDescription(v string) *QueryAppTestRelationResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryAppTestRelationResponseExtra) SetSubErrorCode(v int32) *QueryAppTestRelationResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryAppTestRelationResponseExtra) SetSubDescription(v string) *QueryAppTestRelationResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryAppTestRelationResponseExtra) SetLogid(v string) *QueryAppTestRelationResponseExtra {
	s.Logid = &v
	return s
}

func (s *QueryAppTestRelationResponseExtra) SetNow(v int64) *QueryAppTestRelationResponseExtra {
	s.Now = &v
	return s
}

type QueryApplyPermissionStatusRequest struct {
	ComponentAppid *string            `json:"component_appid,omitempty" xml:"component_appid,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryApplyPermissionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryApplyPermissionStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryApplyPermissionStatusRequest) SetComponentAppid(v string) *QueryApplyPermissionStatusRequest {
	s.ComponentAppid = &v
	return s
}

func (s *QueryApplyPermissionStatusRequest) SetHeader(v map[string]*string) *QueryApplyPermissionStatusRequest {
	s.Header = v
	return s
}

func (s *QueryApplyPermissionStatusRequest) SetAccessToken(v string) *QueryApplyPermissionStatusRequest {
	s.AccessToken = &v
	return s
}

type QueryApplyPermissionStatusResponse struct {
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryApplyPermissionStatusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s QueryApplyPermissionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryApplyPermissionStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryApplyPermissionStatusResponse) SetErrNo(v int32) *QueryApplyPermissionStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryApplyPermissionStatusResponse) SetErrMsg(v string) *QueryApplyPermissionStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryApplyPermissionStatusResponse) SetLogId(v string) *QueryApplyPermissionStatusResponse {
	s.LogId = &v
	return s
}

func (s *QueryApplyPermissionStatusResponse) SetData(v *QueryApplyPermissionStatusResponseData) *QueryApplyPermissionStatusResponse {
	s.Data = v
	return s
}

type QueryApplyPermissionStatusResponseData struct {
	Status *int32  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty" require:"true"`
}

func (s QueryApplyPermissionStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryApplyPermissionStatusResponseData) GoString() string {
	return s.String()
}

func (s *QueryApplyPermissionStatusResponseData) SetStatus(v int32) *QueryApplyPermissionStatusResponseData {
	s.Status = &v
	return s
}

func (s *QueryApplyPermissionStatusResponseData) SetReason(v string) *QueryApplyPermissionStatusResponseData {
	s.Reason = &v
	return s
}

type QueryAriRequest struct {
	BookingStockEndDate   *string            `json:"booking_stock_end_date,omitempty" xml:"booking_stock_end_date,omitempty"`
	ProductId             *string            `json:"product_id,omitempty" xml:"product_id,omitempty"`
	AccountId             *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BookingStockStartDate *string            `json:"booking_stock_start_date,omitempty" xml:"booking_stock_start_date,omitempty"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryAriRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAriRequest) GoString() string {
	return s.String()
}

func (s *QueryAriRequest) SetBookingStockEndDate(v string) *QueryAriRequest {
	s.BookingStockEndDate = &v
	return s
}

func (s *QueryAriRequest) SetProductId(v string) *QueryAriRequest {
	s.ProductId = &v
	return s
}

func (s *QueryAriRequest) SetAccountId(v string) *QueryAriRequest {
	s.AccountId = &v
	return s
}

func (s *QueryAriRequest) SetBookingStockStartDate(v string) *QueryAriRequest {
	s.BookingStockStartDate = &v
	return s
}

func (s *QueryAriRequest) SetHeader(v map[string]*string) *QueryAriRequest {
	s.Header = v
	return s
}

func (s *QueryAriRequest) SetAccessToken(v string) *QueryAriRequest {
	s.AccessToken = &v
	return s
}

type QueryAriResponse struct {
	Data  *QueryAriResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *QueryAriResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s QueryAriResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponse) GoString() string {
	return s.String()
}

func (s *QueryAriResponse) SetData(v *QueryAriResponseData) *QueryAriResponse {
	s.Data = v
	return s
}

func (s *QueryAriResponse) SetExtra(v *QueryAriResponseExtra) *QueryAriResponse {
	s.Extra = v
	return s
}

type QueryAriResponseData struct {
	RoomAddPriceRule     *QueryAriResponseDataRoomAddPriceRule     `json:"room_add_price_rule,omitempty" xml:"room_add_price_rule,omitempty"`
	ProductStatus        *int32                                    `json:"product_status,omitempty" xml:"product_status,omitempty"`
	ChildrenDefinition   *QueryAriResponseDataChildrenDefinition   `json:"children_definition,omitempty" xml:"children_definition,omitempty"`
	BookingCalendarStock *QueryAriResponseDataBookingCalendarStock `json:"booking_calendar_stock,omitempty" xml:"booking_calendar_stock,omitempty"`
	ErrorCode            *string                                   `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ProductName          *string                                   `json:"product_name,omitempty" xml:"product_name,omitempty"`
	UserAddPriceRule     *QueryAriResponseDataUserAddPriceRule     `json:"user_add_price_rule,omitempty" xml:"user_add_price_rule,omitempty"`
	DateAddPriceRule     *QueryAriResponseDataDateAddPriceRule     `json:"date_add_price_rule,omitempty" xml:"date_add_price_rule,omitempty"`
	StockQtyLimitType    *int32                                    `json:"stock_qty_limit_type,omitempty" xml:"stock_qty_limit_type,omitempty"`
	Descripiton          *string                                   `json:"descripiton,omitempty" xml:"descripiton,omitempty"`
	CreateTime           *int64                                    `json:"create_time,omitempty" xml:"create_time,omitempty"`
	CategoryId           *string                                   `json:"category_id,omitempty" xml:"category_id,omitempty"`
	ActualAmount         *int64                                    `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	SoldCount            *int64                                    `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	MarketingAmount      *int64                                    `json:"marketing_amount,omitempty" xml:"marketing_amount,omitempty"`
	TotalStockQty        *int64                                    `json:"total_stock_qty,omitempty" xml:"total_stock_qty,omitempty"`
	GwDescription        *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ProductId            *string                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	UpdateTime           *int64                                    `json:"update_time,omitempty" xml:"update_time,omitempty"`
	OriginalAmount       *int64                                    `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	CategoryName         *string                                   `json:"category_name,omitempty" xml:"category_name,omitempty"`
}

func (s QueryAriResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseData) GoString() string {
	return s.String()
}

func (s *QueryAriResponseData) SetRoomAddPriceRule(v *QueryAriResponseDataRoomAddPriceRule) *QueryAriResponseData {
	s.RoomAddPriceRule = v
	return s
}

func (s *QueryAriResponseData) SetProductStatus(v int32) *QueryAriResponseData {
	s.ProductStatus = &v
	return s
}

func (s *QueryAriResponseData) SetChildrenDefinition(v *QueryAriResponseDataChildrenDefinition) *QueryAriResponseData {
	s.ChildrenDefinition = v
	return s
}

func (s *QueryAriResponseData) SetBookingCalendarStock(v *QueryAriResponseDataBookingCalendarStock) *QueryAriResponseData {
	s.BookingCalendarStock = v
	return s
}

func (s *QueryAriResponseData) SetErrorCode(v string) *QueryAriResponseData {
	s.ErrorCode = &v
	return s
}

func (s *QueryAriResponseData) SetProductName(v string) *QueryAriResponseData {
	s.ProductName = &v
	return s
}

func (s *QueryAriResponseData) SetUserAddPriceRule(v *QueryAriResponseDataUserAddPriceRule) *QueryAriResponseData {
	s.UserAddPriceRule = v
	return s
}

func (s *QueryAriResponseData) SetDateAddPriceRule(v *QueryAriResponseDataDateAddPriceRule) *QueryAriResponseData {
	s.DateAddPriceRule = v
	return s
}

func (s *QueryAriResponseData) SetStockQtyLimitType(v int32) *QueryAriResponseData {
	s.StockQtyLimitType = &v
	return s
}

func (s *QueryAriResponseData) SetDescripiton(v string) *QueryAriResponseData {
	s.Descripiton = &v
	return s
}

func (s *QueryAriResponseData) SetCreateTime(v int64) *QueryAriResponseData {
	s.CreateTime = &v
	return s
}

func (s *QueryAriResponseData) SetCategoryId(v string) *QueryAriResponseData {
	s.CategoryId = &v
	return s
}

func (s *QueryAriResponseData) SetActualAmount(v int64) *QueryAriResponseData {
	s.ActualAmount = &v
	return s
}

func (s *QueryAriResponseData) SetSoldCount(v int64) *QueryAriResponseData {
	s.SoldCount = &v
	return s
}

func (s *QueryAriResponseData) SetMarketingAmount(v int64) *QueryAriResponseData {
	s.MarketingAmount = &v
	return s
}

func (s *QueryAriResponseData) SetTotalStockQty(v int64) *QueryAriResponseData {
	s.TotalStockQty = &v
	return s
}

func (s *QueryAriResponseData) SetGwDescription(v string) *QueryAriResponseData {
	s.GwDescription = &v
	return s
}

func (s *QueryAriResponseData) SetProductId(v string) *QueryAriResponseData {
	s.ProductId = &v
	return s
}

func (s *QueryAriResponseData) SetUpdateTime(v int64) *QueryAriResponseData {
	s.UpdateTime = &v
	return s
}

func (s *QueryAriResponseData) SetOriginalAmount(v int64) *QueryAriResponseData {
	s.OriginalAmount = &v
	return s
}

func (s *QueryAriResponseData) SetCategoryName(v string) *QueryAriResponseData {
	s.CategoryName = &v
	return s
}

type QueryAriResponseDataBookingCalendarStock struct {
	StockList []*QueryAriResponseDataBookingCalendarStockStockListItem `json:"stock_list,omitempty" xml:"stock_list,omitempty" type:"Repeated"`
}

func (s QueryAriResponseDataBookingCalendarStock) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataBookingCalendarStock) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataBookingCalendarStock) SetStockList(v []*QueryAriResponseDataBookingCalendarStockStockListItem) *QueryAriResponseDataBookingCalendarStock {
	s.StockList = v
	return s
}

type QueryAriResponseDataBookingCalendarStockStockListItem struct {
	CalendarStatus    *int32  `json:"calendar_status,omitempty" xml:"calendar_status,omitempty"`
	StockQtyLimitType *int32  `json:"stock_qty_limit_type,omitempty" xml:"stock_qty_limit_type,omitempty"`
	SoldQty           *int64  `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	TotalQty          *int64  `json:"total_qty,omitempty" xml:"total_qty,omitempty"`
	AvailableQty      *int64  `json:"available_qty,omitempty" xml:"available_qty,omitempty"`
	Date              *string `json:"date,omitempty" xml:"date,omitempty"`
}

func (s QueryAriResponseDataBookingCalendarStockStockListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataBookingCalendarStockStockListItem) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetCalendarStatus(v int32) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.CalendarStatus = &v
	return s
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetStockQtyLimitType(v int32) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.StockQtyLimitType = &v
	return s
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetSoldQty(v int64) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.SoldQty = &v
	return s
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetTotalQty(v int64) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.TotalQty = &v
	return s
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetAvailableQty(v int64) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.AvailableQty = &v
	return s
}

func (s *QueryAriResponseDataBookingCalendarStockStockListItem) SetDate(v string) *QueryAriResponseDataBookingCalendarStockStockListItem {
	s.Date = &v
	return s
}

type QueryAriResponseDataChildrenDefinition struct {
	MinAge    *int32 `json:"min_age,omitempty" xml:"min_age,omitempty"`
	MaxAge    *int32 `json:"max_age,omitempty" xml:"max_age,omitempty"`
	MinHeight *int32 `json:"min_height,omitempty" xml:"min_height,omitempty"`
	MaxHeight *int32 `json:"max_height,omitempty" xml:"max_height,omitempty"`
}

func (s QueryAriResponseDataChildrenDefinition) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataChildrenDefinition) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataChildrenDefinition) SetMinAge(v int32) *QueryAriResponseDataChildrenDefinition {
	s.MinAge = &v
	return s
}

func (s *QueryAriResponseDataChildrenDefinition) SetMaxAge(v int32) *QueryAriResponseDataChildrenDefinition {
	s.MaxAge = &v
	return s
}

func (s *QueryAriResponseDataChildrenDefinition) SetMinHeight(v int32) *QueryAriResponseDataChildrenDefinition {
	s.MinHeight = &v
	return s
}

func (s *QueryAriResponseDataChildrenDefinition) SetMaxHeight(v int32) *QueryAriResponseDataChildrenDefinition {
	s.MaxHeight = &v
	return s
}

type QueryAriResponseDataDateAddPriceRule struct {
	Enable           *bool                                                       `json:"enable,omitempty" xml:"enable,omitempty"`
	AddPriceRuleList []*QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
}

func (s QueryAriResponseDataDateAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataDateAddPriceRule) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataDateAddPriceRule) SetEnable(v bool) *QueryAriResponseDataDateAddPriceRule {
	s.Enable = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRule) SetAddPriceRuleList(v []*QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) *QueryAriResponseDataDateAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

type QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem struct {
	Holidays   []*int32 `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	StartDate  *string  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate    *string  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	ChildPrice *int64   `json:"child_price,omitempty" xml:"child_price,omitempty"`
	AdultPrice *int64   `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	Price      *int64   `json:"price,omitempty" xml:"price,omitempty"`
	DateType   *int32   `json:"date_type,omitempty" xml:"date_type,omitempty"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
}

func (s QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

func (s *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *QueryAriResponseDataDateAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

type QueryAriResponseDataRoomAddPriceRule struct {
	AddPriceRuleList []*QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
	Enable           *bool                                                       `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s QueryAriResponseDataRoomAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataRoomAddPriceRule) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataRoomAddPriceRule) SetAddPriceRuleList(v []*QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) *QueryAriResponseDataRoomAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRule) SetEnable(v bool) *QueryAriResponseDataRoomAddPriceRule {
	s.Enable = &v
	return s
}

type QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem struct {
	DateType   *int32   `json:"date_type,omitempty" xml:"date_type,omitempty"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	Holidays   []*int32 `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	StartDate  *string  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate    *string  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	ChildPrice *int64   `json:"child_price,omitempty" xml:"child_price,omitempty"`
	AdultPrice *int64   `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	Price      *int64   `json:"price,omitempty" xml:"price,omitempty"`
}

func (s QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *QueryAriResponseDataRoomAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

type QueryAriResponseDataUserAddPriceRule struct {
	Enable           *bool                                                       `json:"enable,omitempty" xml:"enable,omitempty"`
	AddPriceRuleList []*QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem `json:"add_price_rule_list,omitempty" xml:"add_price_rule_list,omitempty" type:"Repeated"`
}

func (s QueryAriResponseDataUserAddPriceRule) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataUserAddPriceRule) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataUserAddPriceRule) SetEnable(v bool) *QueryAriResponseDataUserAddPriceRule {
	s.Enable = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRule) SetAddPriceRuleList(v []*QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) *QueryAriResponseDataUserAddPriceRule {
	s.AddPriceRuleList = v
	return s
}

type QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem struct {
	AdultPrice *int64   `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	Price      *int64   `json:"price,omitempty" xml:"price,omitempty"`
	DateType   *int32   `json:"date_type,omitempty" xml:"date_type,omitempty"`
	DaysOfWeek []*int32 `json:"days_of_week,omitempty" xml:"days_of_week,omitempty" type:"Repeated"`
	Holidays   []*int32 `json:"holidays,omitempty" xml:"holidays,omitempty" type:"Repeated"`
	StartDate  *string  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate    *string  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	ChildPrice *int64   `json:"child_price,omitempty" xml:"child_price,omitempty"`
}

func (s QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) GoString() string {
	return s.String()
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetAdultPrice(v int64) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.AdultPrice = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetPrice(v int64) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.Price = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetDateType(v int32) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.DateType = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetDaysOfWeek(v []*int32) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.DaysOfWeek = v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetHolidays(v []*int32) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.Holidays = v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetStartDate(v string) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.StartDate = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetEndDate(v string) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.EndDate = &v
	return s
}

func (s *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem) SetChildPrice(v int64) *QueryAriResponseDataUserAddPriceRuleAddPriceRuleListItem {
	s.ChildPrice = &v
	return s
}

type QueryAriResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s QueryAriResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryAriResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryAriResponseExtra) SetNow(v int64) *QueryAriResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryAriResponseExtra) SetErrorCode(v int32) *QueryAriResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *QueryAriResponseExtra) SetDescription(v string) *QueryAriResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryAriResponseExtra) SetSubErrorCode(v int32) *QueryAriResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryAriResponseExtra) SetSubDescription(v string) *QueryAriResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryAriResponseExtra) SetLogid(v string) *QueryAriResponseExtra {
	s.Logid = &v
	return s
}

type QueryAwemeRelationListRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryAwemeRelationListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListRequest) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListRequest) SetAccessToken(v string) *QueryAwemeRelationListRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryAwemeRelationListRequest) SetPageSize(v int64) *QueryAwemeRelationListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAwemeRelationListRequest) SetType(v string) *QueryAwemeRelationListRequest {
	s.Type = &v
	return s
}

func (s *QueryAwemeRelationListRequest) SetPageNum(v int64) *QueryAwemeRelationListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAwemeRelationListRequest) SetHeader(v map[string]*string) *QueryAwemeRelationListRequest {
	s.Header = v
	return s
}

type QueryAwemeRelationListResponse struct {
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryAwemeRelationListResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryAwemeRelationListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponse) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponse) SetErrNo(v int32) *QueryAwemeRelationListResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryAwemeRelationListResponse) SetErrMsg(v string) *QueryAwemeRelationListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryAwemeRelationListResponse) SetLogId(v string) *QueryAwemeRelationListResponse {
	s.LogId = &v
	return s
}

func (s *QueryAwemeRelationListResponse) SetData(v *QueryAwemeRelationListResponseData) *QueryAwemeRelationListResponse {
	s.Data = v
	return s
}

type QueryAwemeRelationListResponseData struct {
	TotalCount *int64                                        `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	List       []*QueryAwemeRelationListResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryAwemeRelationListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseData) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseData) SetTotalCount(v int64) *QueryAwemeRelationListResponseData {
	s.TotalCount = &v
	return s
}

func (s *QueryAwemeRelationListResponseData) SetList(v []*QueryAwemeRelationListResponseDataListItem) *QueryAwemeRelationListResponseData {
	s.List = v
	return s
}

type QueryAwemeRelationListResponseDataListItem struct {
	AwemeId             *string                                                        `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	Reason              *string                                                        `json:"reason,omitempty" xml:"reason,omitempty"`
	UserName            *string                                                        `json:"user_name,omitempty" xml:"user_name,omitempty"`
	UnbindFrequencyInfo *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo `json:"unbind_frequency_info,omitempty" xml:"unbind_frequency_info,omitempty"`
	CapacityList        []*QueryAwemeRelationListResponseDataListItemCapacityListItem  `json:"capacity_list,omitempty" xml:"capacity_list,omitempty" type:"Repeated"`
	AccountType         *int                                                           `json:"account_type,omitempty" xml:"account_type,omitempty"`
	BindStatus          *int                                                           `json:"bind_status,omitempty" xml:"bind_status,omitempty"`
}

func (s QueryAwemeRelationListResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItem) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItem) SetAwemeId(v string) *QueryAwemeRelationListResponseDataListItem {
	s.AwemeId = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetReason(v string) *QueryAwemeRelationListResponseDataListItem {
	s.Reason = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetUserName(v string) *QueryAwemeRelationListResponseDataListItem {
	s.UserName = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetUnbindFrequencyInfo(v *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) *QueryAwemeRelationListResponseDataListItem {
	s.UnbindFrequencyInfo = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetCapacityList(v []*QueryAwemeRelationListResponseDataListItemCapacityListItem) *QueryAwemeRelationListResponseDataListItem {
	s.CapacityList = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetAccountType(v int) *QueryAwemeRelationListResponseDataListItem {
	s.AccountType = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItem) SetBindStatus(v int) *QueryAwemeRelationListResponseDataListItem {
	s.BindStatus = &v
	return s
}

type QueryAwemeRelationListResponseDataListItemCapacityListItem struct {
	Status            *int                                                                         `json:"status,omitempty" xml:"status,omitempty"`
	FailReason        *string                                                                      `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	AuditTemplateInfo *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo `json:"audit_template_info,omitempty" xml:"audit_template_info,omitempty"`
	CapacityKey       *string                                                                      `json:"capacity_key,omitempty" xml:"capacity_key,omitempty"`
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItem) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItem) SetStatus(v int) *QueryAwemeRelationListResponseDataListItemCapacityListItem {
	s.Status = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItem) SetFailReason(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItem {
	s.FailReason = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItem) SetAuditTemplateInfo(v *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) *QueryAwemeRelationListResponseDataListItemCapacityListItem {
	s.AuditTemplateInfo = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItem) SetCapacityKey(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItem {
	s.CapacityKey = &v
	return s
}

type QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo struct {
	Hints           []*string                                                                                         `json:"hints,omitempty" xml:"hints,omitempty" type:"Repeated"`
	Id              *int64                                                                                            `json:"id,omitempty" xml:"id,omitempty"`
	Title           *string                                                                                           `json:"title,omitempty" xml:"title,omitempty"`
	Desc            *string                                                                                           `json:"desc,omitempty" xml:"desc,omitempty"`
	TemplateContent []*QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem `json:"template_content,omitempty" xml:"template_content,omitempty" type:"Repeated"`
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) SetHints(v []*string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo {
	s.Hints = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) SetId(v int64) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo {
	s.Id = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) SetTitle(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo {
	s.Title = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) SetDesc(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo {
	s.Desc = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo) SetTemplateContent(v []*QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfo {
	s.TemplateContent = v
	return s
}

type QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem struct {
	Desc         *string                                                                                                                       `json:"desc,omitempty" xml:"desc,omitempty"`
	RejectReason *string                                                                                                                       `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	ValList      []*string                                                                                                                     `json:"val_list,omitempty" xml:"val_list,omitempty" type:"Repeated"`
	Children     map[string][]*QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem `json:"children,omitempty" xml:"children,omitempty"`
	Name         *string                                                                                                                       `json:"name,omitempty" xml:"name,omitempty"`
	ValType      *int                                                                                                                          `json:"val_type,omitempty" xml:"val_type,omitempty"`
	MaterielId   *string                                                                                                                       `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetDesc(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.Desc = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetRejectReason(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.RejectReason = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetValList(v []*string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.ValList = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetChildren(v map[string][]*QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.Children = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetName(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.Name = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetValType(v int) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.ValType = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem) SetMaterielId(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItem {
	s.MaterielId = &v
	return s
}

type QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem struct {
	MaterielId   *string   `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
	Name         *string   `json:"name,omitempty" xml:"name,omitempty"`
	ValType      *int      `json:"val_type,omitempty" xml:"val_type,omitempty"`
	ValExample   []*string `json:"val_example,omitempty" xml:"val_example,omitempty" type:"Repeated"`
	ValList      []*string `json:"val_list,omitempty" xml:"val_list,omitempty" type:"Repeated"`
	RejectReason *string   `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetMaterielId(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.MaterielId = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetName(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.Name = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetValType(v int) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.ValType = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetValExample(v []*string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.ValExample = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetValList(v []*string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.ValList = v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem) SetRejectReason(v string) *QueryAwemeRelationListResponseDataListItemCapacityListItemAuditTemplateInfoTemplateContentItemChildrenValueItem {
	s.RejectReason = &v
	return s
}

type QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo struct {
	LimitFrequency  *int64 `json:"limit_frequency,omitempty" xml:"limit_frequency,omitempty"`
	UnboundCount    *int64 `json:"unbound_count,omitempty" xml:"unbound_count,omitempty"`
	FirstUnbindTime *int64 `json:"first_unbind_time,omitempty" xml:"first_unbind_time,omitempty"`
	LimitPeriod     *int64 `json:"limit_period,omitempty" xml:"limit_period,omitempty"`
}

func (s QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) GoString() string {
	return s.String()
}

func (s *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) SetLimitFrequency(v int64) *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo {
	s.LimitFrequency = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) SetUnboundCount(v int64) *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo {
	s.UnboundCount = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) SetFirstUnbindTime(v int64) *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo {
	s.FirstUnbindTime = &v
	return s
}

func (s *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo) SetLimitPeriod(v int64) *QueryAwemeRelationListResponseDataListItemUnbindFrequencyInfo {
	s.LimitPeriod = &v
	return s
}

type QueryAwemeVideoKeywordListRequest struct {
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryAwemeVideoKeywordListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeVideoKeywordListRequest) GoString() string {
	return s.String()
}

func (s *QueryAwemeVideoKeywordListRequest) SetPageNum(v int32) *QueryAwemeVideoKeywordListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAwemeVideoKeywordListRequest) SetPageSize(v int32) *QueryAwemeVideoKeywordListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAwemeVideoKeywordListRequest) SetHeader(v map[string]*string) *QueryAwemeVideoKeywordListRequest {
	s.Header = v
	return s
}

func (s *QueryAwemeVideoKeywordListRequest) SetAccessToken(v string) *QueryAwemeVideoKeywordListRequest {
	s.AccessToken = &v
	return s
}

type QueryAwemeVideoKeywordListResponse struct {
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryAwemeVideoKeywordListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryAwemeVideoKeywordListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeVideoKeywordListResponse) GoString() string {
	return s.String()
}

func (s *QueryAwemeVideoKeywordListResponse) SetErrMsg(v string) *QueryAwemeVideoKeywordListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryAwemeVideoKeywordListResponse) SetLogId(v string) *QueryAwemeVideoKeywordListResponse {
	s.LogId = &v
	return s
}

func (s *QueryAwemeVideoKeywordListResponse) SetData(v *QueryAwemeVideoKeywordListResponseData) *QueryAwemeVideoKeywordListResponse {
	s.Data = v
	return s
}

func (s *QueryAwemeVideoKeywordListResponse) SetErrNo(v int32) *QueryAwemeVideoKeywordListResponse {
	s.ErrNo = &v
	return s
}

type QueryAwemeVideoKeywordListResponseData struct {
	KeywordList []*QueryAwemeVideoKeywordListResponseDataKeywordListItem `json:"keyword_list,omitempty" xml:"keyword_list,omitempty" type:"Repeated"`
	TotalCount  *int64                                                   `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
}

func (s QueryAwemeVideoKeywordListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeVideoKeywordListResponseData) GoString() string {
	return s.String()
}

func (s *QueryAwemeVideoKeywordListResponseData) SetKeywordList(v []*QueryAwemeVideoKeywordListResponseDataKeywordListItem) *QueryAwemeVideoKeywordListResponseData {
	s.KeywordList = v
	return s
}

func (s *QueryAwemeVideoKeywordListResponseData) SetTotalCount(v int64) *QueryAwemeVideoKeywordListResponseData {
	s.TotalCount = &v
	return s
}

type QueryAwemeVideoKeywordListResponseDataKeywordListItem struct {
	RejectReasonList []*string `json:"reject_reason_list,omitempty" xml:"reject_reason_list,omitempty" type:"Repeated"`
	KeywordId        *string   `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
	Keyword          *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Status           *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAwemeVideoKeywordListResponseDataKeywordListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAwemeVideoKeywordListResponseDataKeywordListItem) GoString() string {
	return s.String()
}

func (s *QueryAwemeVideoKeywordListResponseDataKeywordListItem) SetRejectReasonList(v []*string) *QueryAwemeVideoKeywordListResponseDataKeywordListItem {
	s.RejectReasonList = v
	return s
}

func (s *QueryAwemeVideoKeywordListResponseDataKeywordListItem) SetKeywordId(v string) *QueryAwemeVideoKeywordListResponseDataKeywordListItem {
	s.KeywordId = &v
	return s
}

func (s *QueryAwemeVideoKeywordListResponseDataKeywordListItem) SetKeyword(v string) *QueryAwemeVideoKeywordListResponseDataKeywordListItem {
	s.Keyword = &v
	return s
}

func (s *QueryAwemeVideoKeywordListResponseDataKeywordListItem) SetStatus(v int32) *QueryAwemeVideoKeywordListResponseDataKeywordListItem {
	s.Status = &v
	return s
}

type QueryCommonPlanTalentListRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PlanId      *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s QueryCommonPlanTalentListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCommonPlanTalentListRequest) GoString() string {
	return s.String()
}

func (s *QueryCommonPlanTalentListRequest) SetHeader(v map[string]*string) *QueryCommonPlanTalentListRequest {
	s.Header = v
	return s
}

func (s *QueryCommonPlanTalentListRequest) SetAccessToken(v string) *QueryCommonPlanTalentListRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryCommonPlanTalentListRequest) SetPlanId(v int64) *QueryCommonPlanTalentListRequest {
	s.PlanId = &v
	return s
}

func (s *QueryCommonPlanTalentListRequest) SetPageNum(v int32) *QueryCommonPlanTalentListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCommonPlanTalentListRequest) SetPageSize(v int32) *QueryCommonPlanTalentListRequest {
	s.PageSize = &v
	return s
}

type QueryCommonPlanTalentListResponse struct {
	Data   *QueryCommonPlanTalentListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryCommonPlanTalentListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCommonPlanTalentListResponse) GoString() string {
	return s.String()
}

func (s *QueryCommonPlanTalentListResponse) SetData(v *QueryCommonPlanTalentListResponseData) *QueryCommonPlanTalentListResponse {
	s.Data = v
	return s
}

func (s *QueryCommonPlanTalentListResponse) SetErrMsg(v string) *QueryCommonPlanTalentListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryCommonPlanTalentListResponse) SetErrNo(v int32) *QueryCommonPlanTalentListResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryCommonPlanTalentListResponse) SetLogId(v string) *QueryCommonPlanTalentListResponse {
	s.LogId = &v
	return s
}

type QueryCommonPlanTalentListResponseData struct {
	TotalCount *int64                                           `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	Data       []*QueryCommonPlanTalentListResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	Date       *string                                          `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	PageCount  *int64                                           `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
}

func (s QueryCommonPlanTalentListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryCommonPlanTalentListResponseData) GoString() string {
	return s.String()
}

func (s *QueryCommonPlanTalentListResponseData) SetTotalCount(v int64) *QueryCommonPlanTalentListResponseData {
	s.TotalCount = &v
	return s
}

func (s *QueryCommonPlanTalentListResponseData) SetData(v []*QueryCommonPlanTalentListResponseDataDataItem) *QueryCommonPlanTalentListResponseData {
	s.Data = v
	return s
}

func (s *QueryCommonPlanTalentListResponseData) SetDate(v string) *QueryCommonPlanTalentListResponseData {
	s.Date = &v
	return s
}

func (s *QueryCommonPlanTalentListResponseData) SetPageCount(v int64) *QueryCommonPlanTalentListResponseData {
	s.PageCount = &v
	return s
}

type QueryCommonPlanTalentListResponseDataDataItem struct {
	ContentType *int32  `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	DouyinId    *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	Gmv         *int64  `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
	Nickname    *string `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
}

func (s QueryCommonPlanTalentListResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s QueryCommonPlanTalentListResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *QueryCommonPlanTalentListResponseDataDataItem) SetContentType(v int32) *QueryCommonPlanTalentListResponseDataDataItem {
	s.ContentType = &v
	return s
}

func (s *QueryCommonPlanTalentListResponseDataDataItem) SetDouyinId(v string) *QueryCommonPlanTalentListResponseDataDataItem {
	s.DouyinId = &v
	return s
}

func (s *QueryCommonPlanTalentListResponseDataDataItem) SetGmv(v int64) *QueryCommonPlanTalentListResponseDataDataItem {
	s.Gmv = &v
	return s
}

func (s *QueryCommonPlanTalentListResponseDataDataItem) SetNickname(v string) *QueryCommonPlanTalentListResponseDataDataItem {
	s.Nickname = &v
	return s
}

type QueryComponentWithDataRequest struct {
	PageNo          *int64             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	PageSize        *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	StartTime       *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime         *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	ComponentIdList []*string          `json:"componentId_list,omitempty" xml:"componentId_list,omitempty" type:"Repeated"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryComponentWithDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDataRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDataRequest) SetPageNo(v int64) *QueryComponentWithDataRequest {
	s.PageNo = &v
	return s
}

func (s *QueryComponentWithDataRequest) SetPageSize(v int64) *QueryComponentWithDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryComponentWithDataRequest) SetStartTime(v int64) *QueryComponentWithDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryComponentWithDataRequest) SetEndTime(v int64) *QueryComponentWithDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryComponentWithDataRequest) SetComponentIdList(v []*string) *QueryComponentWithDataRequest {
	s.ComponentIdList = v
	return s
}

func (s *QueryComponentWithDataRequest) SetHeader(v map[string]*string) *QueryComponentWithDataRequest {
	s.Header = v
	return s
}

func (s *QueryComponentWithDataRequest) SetAccessToken(v string) *QueryComponentWithDataRequest {
	s.AccessToken = &v
	return s
}

type QueryComponentWithDataResponse struct {
	Data   *QueryComponentWithDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryComponentWithDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDataResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDataResponse) SetData(v *QueryComponentWithDataResponseData) *QueryComponentWithDataResponse {
	s.Data = v
	return s
}

func (s *QueryComponentWithDataResponse) SetErrNo(v int32) *QueryComponentWithDataResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryComponentWithDataResponse) SetErrMsg(v string) *QueryComponentWithDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryComponentWithDataResponse) SetLogId(v string) *QueryComponentWithDataResponse {
	s.LogId = &v
	return s
}

type QueryComponentWithDataResponseData struct {
	Total    *int64                                            `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	DataList []*QueryComponentWithDataResponseDataDataListItem `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
}

func (s QueryComponentWithDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDataResponseData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDataResponseData) SetTotal(v int64) *QueryComponentWithDataResponseData {
	s.Total = &v
	return s
}

func (s *QueryComponentWithDataResponseData) SetDataList(v []*QueryComponentWithDataResponseDataDataListItem) *QueryComponentWithDataResponseData {
	s.DataList = v
	return s
}

type QueryComponentWithDataResponseDataDataListItem struct {
	ComponentReportUv *int64  `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentName     *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ComponentShowUv   *int64  `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ComponentId       *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	ComponentClickPv  *int64  `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
	ComponentShowPv   *int64  `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
	ComponentClickUv  *int64  `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
	ComponentReportPv *int64  `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
}

func (s QueryComponentWithDataResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDataResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentReportUv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentName(v string) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentName = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentShowUv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentId(v string) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentId = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentClickPv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentClickPv = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentShowPv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentShowPv = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentClickUv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentClickUv = &v
	return s
}

func (s *QueryComponentWithDataResponseDataDataListItem) SetComponentReportPv(v int64) *QueryComponentWithDataResponseDataDataListItem {
	s.ComponentReportPv = &v
	return s
}

type QueryComponentWithDetailRequest struct {
	PageNo          *int64             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	IsQueryLive     *bool              `json:"is_query_live,omitempty" xml:"is_query_live,omitempty" require:"true"`
	StartTime       *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	ComponentIdList []*string          `json:"componentId_list,omitempty" xml:"componentId_list,omitempty" type:"Repeated"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	PageSize        *int64             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	IsQueryVideo    *bool              `json:"is_query_video,omitempty" xml:"is_query_video,omitempty" require:"true"`
	EndTime         *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryComponentWithDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailRequest) SetPageNo(v int64) *QueryComponentWithDetailRequest {
	s.PageNo = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetIsQueryLive(v bool) *QueryComponentWithDetailRequest {
	s.IsQueryLive = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetStartTime(v int64) *QueryComponentWithDetailRequest {
	s.StartTime = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetComponentIdList(v []*string) *QueryComponentWithDetailRequest {
	s.ComponentIdList = v
	return s
}

func (s *QueryComponentWithDetailRequest) SetHeader(v map[string]*string) *QueryComponentWithDetailRequest {
	s.Header = v
	return s
}

func (s *QueryComponentWithDetailRequest) SetPageSize(v int64) *QueryComponentWithDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetIsQueryVideo(v bool) *QueryComponentWithDetailRequest {
	s.IsQueryVideo = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetEndTime(v int64) *QueryComponentWithDetailRequest {
	s.EndTime = &v
	return s
}

func (s *QueryComponentWithDetailRequest) SetAccessToken(v string) *QueryComponentWithDetailRequest {
	s.AccessToken = &v
	return s
}

type QueryComponentWithDetailResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *QueryComponentWithDetailResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponse) SetErrMsg(v string) *QueryComponentWithDetailResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryComponentWithDetailResponse) SetLogId(v string) *QueryComponentWithDetailResponse {
	s.LogId = &v
	return s
}

func (s *QueryComponentWithDetailResponse) SetData(v *QueryComponentWithDetailResponseData) *QueryComponentWithDetailResponse {
	s.Data = v
	return s
}

func (s *QueryComponentWithDetailResponse) SetErrNo(v int32) *QueryComponentWithDetailResponse {
	s.ErrNo = &v
	return s
}

type QueryComponentWithDetailResponseData struct {
	VideoData *QueryComponentWithDetailResponseDataVideoData `json:"VideoData,omitempty" xml:"VideoData,omitempty"`
	LiveData  *QueryComponentWithDetailResponseDataLiveData  `json:"LiveData,omitempty" xml:"LiveData,omitempty"`
}

func (s QueryComponentWithDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseData) SetVideoData(v *QueryComponentWithDetailResponseDataVideoData) *QueryComponentWithDetailResponseData {
	s.VideoData = v
	return s
}

func (s *QueryComponentWithDetailResponseData) SetLiveData(v *QueryComponentWithDetailResponseDataLiveData) *QueryComponentWithDetailResponseData {
	s.LiveData = v
	return s
}

type QueryComponentWithDetailResponseDataLiveData struct {
	DataList []*QueryComponentWithDetailResponseDataLiveDataDataListItem `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Total    *int64                                                      `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataLiveData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataLiveData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataLiveData) SetDataList(v []*QueryComponentWithDetailResponseDataLiveDataDataListItem) *QueryComponentWithDetailResponseDataLiveData {
	s.DataList = v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveData) SetTotal(v int64) *QueryComponentWithDetailResponseDataLiveData {
	s.Total = &v
	return s
}

type QueryComponentWithDetailResponseDataLiveDataDataListItem struct {
	RoomId         *int64                                                             `json:"RoomId,omitempty" xml:"RoomId,omitempty" require:"true"`
	RoomCreateTime *int64                                                             `json:"RoomCreateTime,omitempty" xml:"RoomCreateTime,omitempty" require:"true"`
	RoomDuration   *int64                                                             `json:"RoomDuration,omitempty" xml:"RoomDuration,omitempty" require:"true"`
	ShortInfo      *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo `json:"ShortInfo,omitempty" xml:"ShortInfo,omitempty"`
	Data           *QueryComponentWithDetailResponseDataLiveDataDataListItemData      `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItem) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItem) SetRoomId(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItem {
	s.RoomId = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItem) SetRoomCreateTime(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItem {
	s.RoomCreateTime = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItem) SetRoomDuration(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItem {
	s.RoomDuration = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItem) SetShortInfo(v *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) *QueryComponentWithDetailResponseDataLiveDataDataListItem {
	s.ShortInfo = v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItem) SetData(v *QueryComponentWithDetailResponseDataLiveDataDataListItemData) *QueryComponentWithDetailResponseDataLiveDataDataListItem {
	s.Data = v
	return s
}

type QueryComponentWithDetailResponseDataLiveDataDataListItemData struct {
	ComponentReportUv *int64 `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentShowPv   *int64 `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
	ComponentShowUv   *int64 `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ComponentReportPv *int64 `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
	ComponentClickPv  *int64 `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
	ComponentClickUv  *int64 `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItemData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItemData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentReportUv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentShowPv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentShowPv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentShowUv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentReportPv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentReportPv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentClickPv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentClickPv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemData) SetComponentClickUv(v int64) *QueryComponentWithDetailResponseDataLiveDataDataListItemData {
	s.ComponentClickUv = &v
	return s
}

type QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo struct {
	AccountType  *int    `json:"AccountType,omitempty" xml:"AccountType,omitempty" require:"true"`
	Nickname     *string `json:"Nickname,omitempty" xml:"Nickname,omitempty" require:"true"`
	Avat         *string `json:"Avat,omitempty" xml:"Avat,omitempty" require:"true"`
	AwemeShortID *string `json:"AwemeShortID,omitempty" xml:"AwemeShortID,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) SetAccountType(v int) *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo {
	s.AccountType = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) SetNickname(v string) *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo {
	s.Nickname = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) SetAvat(v string) *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo {
	s.Avat = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo) SetAwemeShortID(v string) *QueryComponentWithDetailResponseDataLiveDataDataListItemShortInfo {
	s.AwemeShortID = &v
	return s
}

type QueryComponentWithDetailResponseDataVideoData struct {
	DataList []*QueryComponentWithDetailResponseDataVideoDataDataListItem `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Total    *int64                                                       `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataVideoData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataVideoData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataVideoData) SetDataList(v []*QueryComponentWithDetailResponseDataVideoDataDataListItem) *QueryComponentWithDetailResponseDataVideoData {
	s.DataList = v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoData) SetTotal(v int64) *QueryComponentWithDetailResponseDataVideoData {
	s.Total = &v
	return s
}

type QueryComponentWithDetailResponseDataVideoDataDataListItem struct {
	ItemDuration   *int64                                                              `json:"ItemDuration,omitempty" xml:"ItemDuration,omitempty" require:"true"`
	ItemAddr       *string                                                             `json:"ItemAddr,omitempty" xml:"ItemAddr,omitempty" require:"true"`
	ShortInfo      *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo `json:"ShortInfo,omitempty" xml:"ShortInfo,omitempty"`
	ItemId         *int64                                                              `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	ItemVv         *int64                                                              `json:"ItemVv,omitempty" xml:"ItemVv,omitempty" require:"true"`
	Data           *QueryComponentWithDetailResponseDataVideoDataDataListItemData      `json:"Data,omitempty" xml:"Data,omitempty"`
	ItemCover      *string                                                             `json:"ItemCover,omitempty" xml:"ItemCover,omitempty" require:"true"`
	ItemCreateTime *int64                                                              `json:"ItemCreateTime,omitempty" xml:"ItemCreateTime,omitempty" require:"true"`
	ItemTitle      *string                                                             `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty" require:"true"`
	CompletionRate *int64                                                              `json:"CompletionRate,omitempty" xml:"CompletionRate,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItem) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemDuration(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemDuration = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemAddr(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemAddr = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetShortInfo(v *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ShortInfo = v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemId(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemId = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemVv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemVv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetData(v *QueryComponentWithDetailResponseDataVideoDataDataListItemData) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.Data = v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemCover(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemCover = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemCreateTime(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemCreateTime = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetItemTitle(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.ItemTitle = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItem) SetCompletionRate(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItem {
	s.CompletionRate = &v
	return s
}

type QueryComponentWithDetailResponseDataVideoDataDataListItemData struct {
	ComponentReportPv *int64 `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
	ComponentClickUv  *int64 `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
	ComponentShowUv   *int64 `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ComponentReportUv *int64 `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentShowPv   *int64 `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
	ComponentClickPv  *int64 `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItemData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItemData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentReportPv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentReportPv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentClickUv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentClickUv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentShowUv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentReportUv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentShowPv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentShowPv = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemData) SetComponentClickPv(v int64) *QueryComponentWithDetailResponseDataVideoDataDataListItemData {
	s.ComponentClickPv = &v
	return s
}

type QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo struct {
	Avat         *string `json:"Avat,omitempty" xml:"Avat,omitempty" require:"true"`
	AwemeShortID *string `json:"AwemeShortID,omitempty" xml:"AwemeShortID,omitempty" require:"true"`
	AccountType  *int    `json:"AccountType,omitempty" xml:"AccountType,omitempty" require:"true"`
	Nickname     *string `json:"Nickname,omitempty" xml:"Nickname,omitempty" require:"true"`
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) GoString() string {
	return s.String()
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) SetAvat(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo {
	s.Avat = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) SetAwemeShortID(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo {
	s.AwemeShortID = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) SetAccountType(v int) *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo {
	s.AccountType = &v
	return s
}

func (s *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo) SetNickname(v string) *QueryComponentWithDetailResponseDataVideoDataDataListItemShortInfo {
	s.Nickname = &v
	return s
}

type QueryComponentWithOverviewRequest struct {
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	EndTime         *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	ComponentIdList []*string          `json:"componentId_list,omitempty" xml:"componentId_list,omitempty" type:"Repeated"`
	StartTime       *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryComponentWithOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithOverviewRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentWithOverviewRequest) SetAccessToken(v string) *QueryComponentWithOverviewRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryComponentWithOverviewRequest) SetEndTime(v int64) *QueryComponentWithOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *QueryComponentWithOverviewRequest) SetComponentIdList(v []*string) *QueryComponentWithOverviewRequest {
	s.ComponentIdList = v
	return s
}

func (s *QueryComponentWithOverviewRequest) SetStartTime(v int64) *QueryComponentWithOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *QueryComponentWithOverviewRequest) SetHeader(v map[string]*string) *QueryComponentWithOverviewRequest {
	s.Header = v
	return s
}

type QueryComponentWithOverviewResponse struct {
	Data   *QueryComponentWithOverviewResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryComponentWithOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithOverviewResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentWithOverviewResponse) SetData(v *QueryComponentWithOverviewResponseData) *QueryComponentWithOverviewResponse {
	s.Data = v
	return s
}

func (s *QueryComponentWithOverviewResponse) SetErrNo(v int32) *QueryComponentWithOverviewResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryComponentWithOverviewResponse) SetErrMsg(v string) *QueryComponentWithOverviewResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryComponentWithOverviewResponse) SetLogId(v string) *QueryComponentWithOverviewResponse {
	s.LogId = &v
	return s
}

type QueryComponentWithOverviewResponseData struct {
	ComponentOverviewData *QueryComponentWithOverviewResponseDataComponentOverviewData       `json:"ComponentOverviewData,omitempty" xml:"ComponentOverviewData,omitempty" require:"true"`
	ComponentDataInfoList []*QueryComponentWithOverviewResponseDataComponentDataInfoListItem `json:"ComponentDataInfoList,omitempty" xml:"ComponentDataInfoList,omitempty" type:"Repeated"`
}

func (s QueryComponentWithOverviewResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithOverviewResponseData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithOverviewResponseData) SetComponentOverviewData(v *QueryComponentWithOverviewResponseDataComponentOverviewData) *QueryComponentWithOverviewResponseData {
	s.ComponentOverviewData = v
	return s
}

func (s *QueryComponentWithOverviewResponseData) SetComponentDataInfoList(v []*QueryComponentWithOverviewResponseDataComponentDataInfoListItem) *QueryComponentWithOverviewResponseData {
	s.ComponentDataInfoList = v
	return s
}

type QueryComponentWithOverviewResponseDataComponentDataInfoListItem struct {
	Date              *string `json:"Date,omitempty" xml:"Date,omitempty"`
	Hour              *string `json:"Hour,omitempty" xml:"Hour,omitempty"`
	ValidClueUcnt     *int64  `json:"ValidClueUcnt,omitempty" xml:"ValidClueUcnt,omitempty"`
	ValidClueCnt      *int64  `json:"ValidClueCnt,omitempty" xml:"ValidClueCnt,omitempty"`
	ClueCnt           *int64  `json:"ClueCnt,omitempty" xml:"ClueCnt,omitempty"`
	ComponentReportPv *int64  `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
	ComponentClickUv  *int64  `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
	ComponentShowPv   *int64  `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
	ComponentShowUv   *int64  `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ClueUcnt          *int64  `json:"ClueUcnt,omitempty" xml:"ClueUcnt,omitempty"`
	ComponentReportUv *int64  `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentClickPv  *int64  `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
}

func (s QueryComponentWithOverviewResponseDataComponentDataInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithOverviewResponseDataComponentDataInfoListItem) GoString() string {
	return s.String()
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetDate(v string) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.Date = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetHour(v string) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.Hour = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetValidClueUcnt(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ValidClueUcnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetValidClueCnt(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ValidClueCnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetClueCnt(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ClueCnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentReportPv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentReportPv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentClickUv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentClickUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentShowPv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentShowPv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentShowUv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetClueUcnt(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ClueUcnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentReportUv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentDataInfoListItem) SetComponentClickPv(v int64) *QueryComponentWithOverviewResponseDataComponentDataInfoListItem {
	s.ComponentClickPv = &v
	return s
}

type QueryComponentWithOverviewResponseDataComponentOverviewData struct {
	ComponentClickPv  *int64 `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
	ValidClueCnt      *int64 `json:"ValidClueCnt,omitempty" xml:"ValidClueCnt,omitempty"`
	ComponentReportUv *int64 `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentShowUv   *int64 `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ComponentClickUv  *int64 `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
	ClueCnt           *int64 `json:"ClueCnt,omitempty" xml:"ClueCnt,omitempty"`
	ClueUcnt          *int64 `json:"ClueUcnt,omitempty" xml:"ClueUcnt,omitempty"`
	ComponentReportPv *int64 `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
	ValidClueUcnt     *int64 `json:"ValidClueUcnt,omitempty" xml:"ValidClueUcnt,omitempty"`
	ComponentShowPv   *int64 `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
}

func (s QueryComponentWithOverviewResponseDataComponentOverviewData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithOverviewResponseDataComponentOverviewData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentClickPv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentClickPv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetValidClueCnt(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ValidClueCnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentReportUv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentShowUv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentClickUv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentClickUv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetClueCnt(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ClueCnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetClueUcnt(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ClueUcnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentReportPv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentReportPv = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetValidClueUcnt(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ValidClueUcnt = &v
	return s
}

func (s *QueryComponentWithOverviewResponseDataComponentOverviewData) SetComponentShowPv(v int64) *QueryComponentWithOverviewResponseDataComponentOverviewData {
	s.ComponentShowPv = &v
	return s
}

type QueryComponentWithSourceRequest struct {
	StartTime       *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime         *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ComponentIdList []*string          `json:"componentId_list,omitempty" xml:"componentId_list,omitempty" type:"Repeated"`
}

func (s QueryComponentWithSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithSourceRequest) GoString() string {
	return s.String()
}

func (s *QueryComponentWithSourceRequest) SetStartTime(v int64) *QueryComponentWithSourceRequest {
	s.StartTime = &v
	return s
}

func (s *QueryComponentWithSourceRequest) SetEndTime(v int64) *QueryComponentWithSourceRequest {
	s.EndTime = &v
	return s
}

func (s *QueryComponentWithSourceRequest) SetHeader(v map[string]*string) *QueryComponentWithSourceRequest {
	s.Header = v
	return s
}

func (s *QueryComponentWithSourceRequest) SetAccessToken(v string) *QueryComponentWithSourceRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryComponentWithSourceRequest) SetComponentIdList(v []*string) *QueryComponentWithSourceRequest {
	s.ComponentIdList = v
	return s
}

type QueryComponentWithSourceResponse struct {
	Data   *QueryComponentWithSourceResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryComponentWithSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithSourceResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentWithSourceResponse) SetData(v *QueryComponentWithSourceResponseData) *QueryComponentWithSourceResponse {
	s.Data = v
	return s
}

func (s *QueryComponentWithSourceResponse) SetErrNo(v int32) *QueryComponentWithSourceResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryComponentWithSourceResponse) SetErrMsg(v string) *QueryComponentWithSourceResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryComponentWithSourceResponse) SetLogId(v string) *QueryComponentWithSourceResponse {
	s.LogId = &v
	return s
}

type QueryComponentWithSourceResponseData struct {
	ScenesDataList []*QueryComponentWithSourceResponseDataScenesDataListItem `json:"ScenesDataList,omitempty" xml:"ScenesDataList,omitempty" type:"Repeated"`
}

func (s QueryComponentWithSourceResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithSourceResponseData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithSourceResponseData) SetScenesDataList(v []*QueryComponentWithSourceResponseDataScenesDataListItem) *QueryComponentWithSourceResponseData {
	s.ScenesDataList = v
	return s
}

type QueryComponentWithSourceResponseDataScenesDataListItem struct {
	Data       *QueryComponentWithSourceResponseDataScenesDataListItemData `json:"Data,omitempty" xml:"Data,omitempty"`
	Scenes     *string                                                     `json:"Scenes,omitempty" xml:"Scenes,omitempty" require:"true"`
	ScenesName *string                                                     `json:"ScenesName,omitempty" xml:"ScenesName,omitempty" require:"true"`
}

func (s QueryComponentWithSourceResponseDataScenesDataListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithSourceResponseDataScenesDataListItem) GoString() string {
	return s.String()
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItem) SetData(v *QueryComponentWithSourceResponseDataScenesDataListItemData) *QueryComponentWithSourceResponseDataScenesDataListItem {
	s.Data = v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItem) SetScenes(v string) *QueryComponentWithSourceResponseDataScenesDataListItem {
	s.Scenes = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItem) SetScenesName(v string) *QueryComponentWithSourceResponseDataScenesDataListItem {
	s.ScenesName = &v
	return s
}

type QueryComponentWithSourceResponseDataScenesDataListItemData struct {
	ComponentClickUv  *int64 `json:"ComponentClickUv,omitempty" xml:"ComponentClickUv,omitempty" require:"true"`
	ComponentReportUv *int64 `json:"ComponentReportUv,omitempty" xml:"ComponentReportUv,omitempty" require:"true"`
	ComponentShowUv   *int64 `json:"ComponentShowUv,omitempty" xml:"ComponentShowUv,omitempty" require:"true"`
	ComponentClickPv  *int64 `json:"ComponentClickPv,omitempty" xml:"ComponentClickPv,omitempty" require:"true"`
	ComponentShowPv   *int64 `json:"ComponentShowPv,omitempty" xml:"ComponentShowPv,omitempty" require:"true"`
	ComponentReportPv *int64 `json:"ComponentReportPv,omitempty" xml:"ComponentReportPv,omitempty" require:"true"`
}

func (s QueryComponentWithSourceResponseDataScenesDataListItemData) String() string {
	return tea.Prettify(s)
}

func (s QueryComponentWithSourceResponseDataScenesDataListItemData) GoString() string {
	return s.String()
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentClickUv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentClickUv = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentReportUv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentReportUv = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentShowUv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentShowUv = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentClickPv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentClickPv = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentShowPv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentShowPv = &v
	return s
}

func (s *QueryComponentWithSourceResponseDataScenesDataListItemData) SetComponentReportPv(v int64) *QueryComponentWithSourceResponseDataScenesDataListItemData {
	s.ComponentReportPv = &v
	return s
}

type QueryContactRequest struct {
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	SourcePhoneNumber *string            `json:"source_phone_number,omitempty" xml:"source_phone_number,omitempty"`
	OrderId           *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s QueryContactRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryContactRequest) GoString() string {
	return s.String()
}

func (s *QueryContactRequest) SetAccessToken(v string) *QueryContactRequest {
	s.AccessToken = &v
	return s
}

func (s *QueryContactRequest) SetSourcePhoneNumber(v string) *QueryContactRequest {
	s.SourcePhoneNumber = &v
	return s
}

func (s *QueryContactRequest) SetOrderId(v string) *QueryContactRequest {
	s.OrderId = &v
	return s
}

func (s *QueryContactRequest) SetHeader(v map[string]*string) *QueryContactRequest {
	s.Header = v
	return s
}

type QueryContactResponse struct {
	Extra *QueryContactResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *QueryContactResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryContactResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryContactResponse) GoString() string {
	return s.String()
}

func (s *QueryContactResponse) SetExtra(v *QueryContactResponseExtra) *QueryContactResponse {
	s.Extra = v
	return s
}

func (s *QueryContactResponse) SetData(v *QueryContactResponseData) *QueryContactResponse {
	s.Data = v
	return s
}

type QueryContactResponseData struct {
	GwErrorCode       *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription     *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SourcePhoneNumber *string `json:"source_phone_number,omitempty" xml:"source_phone_number,omitempty"`
	SecretNumber      *string `json:"secret_number,omitempty" xml:"secret_number,omitempty"`
}

func (s QueryContactResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryContactResponseData) GoString() string {
	return s.String()
}

func (s *QueryContactResponseData) SetGwErrorCode(v int32) *QueryContactResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *QueryContactResponseData) SetGwDescription(v string) *QueryContactResponseData {
	s.GwDescription = &v
	return s
}

func (s *QueryContactResponseData) SetSourcePhoneNumber(v string) *QueryContactResponseData {
	s.SourcePhoneNumber = &v
	return s
}

func (s *QueryContactResponseData) SetSecretNumber(v string) *QueryContactResponseData {
	s.SecretNumber = &v
	return s
}

type QueryContactResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s QueryContactResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s QueryContactResponseExtra) GoString() string {
	return s.String()
}

func (s *QueryContactResponseExtra) SetLogid(v string) *QueryContactResponseExtra {
	s.Logid = &v
	return s
}

func (s *QueryContactResponseExtra) SetNow(v int64) *QueryContactResponseExtra {
	s.Now = &v
	return s
}

func (s *QueryContactResponseExtra) SetSubDescription(v string) *QueryContactResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *QueryContactResponseExtra) SetSubErrorCode(v int32) *QueryContactResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *QueryContactResponseExtra) SetDescription(v string) *QueryContactResponseExtra {
	s.Description = &v
	return s
}

func (s *QueryContactResponseExtra) SetErrorCode(v int32) *QueryContactResponseExtra {
	s.ErrorCode = &v
	return s
}

type QueryCreatedTplListRequest struct {
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Status      *int               `json:"status,omitempty" xml:"status,omitempty"`
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s QueryCreatedTplListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreatedTplListRequest) GoString() string {
	return s.String()
}

func (s *QueryCreatedTplListRequest) SetPageSize(v int32) *QueryCreatedTplListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCreatedTplListRequest) SetStatus(v int) *QueryCreatedTplListRequest {
	s.Status = &v
	return s
}

func (s *QueryCreatedTplListRequest) SetPageNum(v int32) *QueryCreatedTplListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCreatedTplListRequest) SetHeader(v map[string]*string) *QueryCreatedTplListRequest {
	s.Header = v
	return s
}

func (s *QueryCreatedTplListRequest) SetAccessToken(v string) *QueryCreatedTplListRequest {
	s.AccessToken = &v
	return s
}

type QueryCreatedTplListResponse struct {
	Data   *QueryCreatedTplListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s QueryCreatedTplListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreatedTplListResponse) GoString() string {
	return s.String()
}

func (s *QueryCreatedTplListResponse) SetData(v *QueryCreatedTplListResponseData) *QueryCreatedTplListResponse {
	s.Data = v
	return s
}

func (s *QueryCreatedTplListResponse) SetErrNo(v int32) *QueryCreatedTplListResponse {
	s.ErrNo = &v
	return s
}

func (s *QueryCreatedTplListResponse) SetErrMsg(v string) *QueryCreatedTplListResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryCreatedTplListResponse) SetLogId(v string) *QueryCreatedTplListResponse {
	s.LogId = &v
	return s
}

type QueryCreatedTplListResponseData struct {
	TotalCount   *int64                                             `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	TemplateList []*QueryCreatedTplListResponseDataTemplateListItem `json:"template_list,omitempty" xml:"template_list,omitempty" require:"true" type:"Repeated"`
}

func (s QueryCreatedTplListResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryCreatedTplListResponseData) GoString() string {
	return s.String()
}

func (s *QueryCreatedTplListResponseData) SetTotalCount(v int64) *QueryCreatedTplListResponseData {
	s.TotalCount = &v
	return s
}

func (s *QueryCreatedTplListResponseData) SetTemplateList(v []*QueryCreatedTplListResponseDataTemplateListItem) *QueryCreatedTplListResponseData {
	s.TemplateList = v
	return s
}
