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

type CertificateQueryResponseDataCertificatesItemSku struct {
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemSku) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemSku) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetThirdSkuId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetMarketPrice(v int64) *CertificateQueryResponseDataCertificatesItemSku {
	s.MarketPrice = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetGrouponType(v int) *CertificateQueryResponseDataCertificatesItemSku {
	s.GrouponType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetSuplierProductOutId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetSkuId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.SkuId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetAccountId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.AccountId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetSkuOutId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.SkuOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetProductOutId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.ProductOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetVoucherType(v int) *CertificateQueryResponseDataCertificatesItemSku {
	s.VoucherType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetTitle(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.Title = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetSoldStartTime(v int64) *CertificateQueryResponseDataCertificatesItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemSku) SetProductId(v string) *CertificateQueryResponseDataCertificatesItemSku {
	s.ProductId = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemTimeCard struct {
	TimeCardType     *int                                                                        `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
	TimesCount       *int32                                                                      `json:"times_count,omitempty" xml:"times_count,omitempty"`
	TimesUsed        *int32                                                                      `json:"times_used,omitempty" xml:"times_used,omitempty"`
	SerialAmountList []*CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesItemTimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemTimeCard) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCard) SetTimeCardType(v int) *CertificateQueryResponseDataCertificatesItemTimeCard {
	s.TimeCardType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCard) SetTimesCount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCard {
	s.TimesCount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCard) SetTimesUsed(v int32) *CertificateQueryResponseDataCertificatesItemTimeCard {
	s.TimesUsed = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCard) SetSerialAmountList(v []*CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem) *CertificateQueryResponseDataCertificatesItemTimeCard {
	s.SerialAmountList = v
	return s
}

type CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem struct {
	Amount     *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                          `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem) SetAmount(v *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount struct {
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificateQueryResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemUseTimeInfo struct {
	UseTimeType    *int                                                                         `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
	TimePeriodList []*CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemUseTimeInfo) SetUseTimeType(v int) *CertificateQueryResponseDataCertificatesItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemUseTimeInfo) SetTimePeriodList(v []*CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) *CertificateQueryResponseDataCertificatesItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

type CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem struct {
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificateQueryResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemVerify struct {
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesItemVerify) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemVerify) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetCanCancel(v bool) *CertificateQueryResponseDataCertificatesItemVerify {
	s.CanCancel = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetCertificateId(v string) *CertificateQueryResponseDataCertificatesItemVerify {
	s.CertificateId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetPoiId(v int64) *CertificateQueryResponseDataCertificatesItemVerify {
	s.PoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetTimesCardSerialNum(v int32) *CertificateQueryResponseDataCertificatesItemVerify {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetVerifierUniqueId(v string) *CertificateQueryResponseDataCertificatesItemVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetVerifyId(v string) *CertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetVerifyTime(v int64) *CertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerify) SetVerifyType(v int) *CertificateQueryResponseDataCertificatesItemVerify {
	s.VerifyType = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemVerifyRecordsItem struct {
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetCanCancel(v bool) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetCertificateId(v string) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetPoiId(v int64) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.PoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetVerifierUniqueId(v string) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetVerifyId(v string) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetVerifyTime(v int64) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) SetVerifyType(v int) *CertificateQueryResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyType = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2Item struct {
	StartTime            *int64                                                              `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Status               *int                                                                `json:"status,omitempty" xml:"status,omitempty"`
	TimeCard             *CertificateQueryResponseDataCertificatesV2ItemTimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	VerifyRecords        []*CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	Amount               *CertificateQueryResponseDataCertificatesV2ItemAmount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	BookInfo             *CertificateQueryResponseDataCertificatesV2ItemBookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	UsedStatusType       *int                                                                `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
	ExpireTime           *int64                                                              `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	EncryptedCode        *string                                                             `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	NotAvailablePoiList  []*string                                                           `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	Sku                  *CertificateQueryResponseDataCertificatesV2ItemSku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	Code                 *string                                                             `json:"code,omitempty" xml:"code,omitempty"`
	AdditionalMap        map[int]*string                                                     `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	Verify               *CertificateQueryResponseDataCertificatesV2ItemVerify               `json:"verify,omitempty" xml:"verify,omitempty"`
	ReserveInfo          *CertificateQueryResponseDataCertificatesV2ItemReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	PeriodCard           *CertificateQueryResponseDataCertificatesV2ItemPeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	NotAvailableTimeInfo *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	CertificateId        *int64                                                              `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	UseTimeInfo          *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	OffPeakDiscountInfo  *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2Item) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2Item) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesV2Item {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetStatus(v int) *CertificateQueryResponseDataCertificatesV2Item {
	s.Status = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetTimeCard(v *CertificateQueryResponseDataCertificatesV2ItemTimeCard) *CertificateQueryResponseDataCertificatesV2Item {
	s.TimeCard = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetVerifyRecords(v []*CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) *CertificateQueryResponseDataCertificatesV2Item {
	s.VerifyRecords = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetAmount(v *CertificateQueryResponseDataCertificatesV2ItemAmount) *CertificateQueryResponseDataCertificatesV2Item {
	s.Amount = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetBookInfo(v *CertificateQueryResponseDataCertificatesV2ItemBookInfo) *CertificateQueryResponseDataCertificatesV2Item {
	s.BookInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetUsedStatusType(v int) *CertificateQueryResponseDataCertificatesV2Item {
	s.UsedStatusType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetExpireTime(v int64) *CertificateQueryResponseDataCertificatesV2Item {
	s.ExpireTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetEncryptedCode(v string) *CertificateQueryResponseDataCertificatesV2Item {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetNotAvailablePoiList(v []*string) *CertificateQueryResponseDataCertificatesV2Item {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetSku(v *CertificateQueryResponseDataCertificatesV2ItemSku) *CertificateQueryResponseDataCertificatesV2Item {
	s.Sku = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetCode(v string) *CertificateQueryResponseDataCertificatesV2Item {
	s.Code = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetAdditionalMap(v map[int]*string) *CertificateQueryResponseDataCertificatesV2Item {
	s.AdditionalMap = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetVerify(v *CertificateQueryResponseDataCertificatesV2ItemVerify) *CertificateQueryResponseDataCertificatesV2Item {
	s.Verify = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetReserveInfo(v *CertificateQueryResponseDataCertificatesV2ItemReserveInfo) *CertificateQueryResponseDataCertificatesV2Item {
	s.ReserveInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetPeriodCard(v *CertificateQueryResponseDataCertificatesV2ItemPeriodCard) *CertificateQueryResponseDataCertificatesV2Item {
	s.PeriodCard = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetNotAvailableTimeInfo(v *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) *CertificateQueryResponseDataCertificatesV2Item {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetCertificateId(v int64) *CertificateQueryResponseDataCertificatesV2Item {
	s.CertificateId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetUseTimeInfo(v *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo) *CertificateQueryResponseDataCertificatesV2Item {
	s.UseTimeInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2Item) SetOffPeakDiscountInfo(v *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) *CertificateQueryResponseDataCertificatesV2Item {
	s.OffPeakDiscountInfo = v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemAmount struct {
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetMerchantTicketAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetCouponPayAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetPlatformDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetListMarketAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetOriginListMarketAmount(v int64) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetPaymentDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetPayAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetBrandTicketAmount(v int64) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetOriginalAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemAmount) SetOriginalCurrency(v string) *CertificateQueryResponseDataCertificatesV2ItemAmount {
	s.OriginalCurrency = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemBookInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemBookInfo) SetBookPoiId(v string) *CertificateQueryResponseDataCertificatesV2ItemBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemBookInfo) SetBookProductNumber(v int64) *CertificateQueryResponseDataCertificatesV2ItemBookInfo {
	s.BookProductNumber = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo struct {
	FulfilEnable    *bool                                                                                 `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
	CanNoUseDate    []*CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
	CanNoUseWeekDay []*int64                                                                              `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetCanNoUseDate(v []*CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo struct {
	HasOffPeakDiscount *bool                                                                                    `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
	IdleTimeLimitType  *int                                                                                     `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	OffPeakTimeRange   []*CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem struct {
	EndTime            *int64                                                                                                         `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                         `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                         `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificateQueryResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemPeriodCard) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemPeriodCard) SetPeriodType(v int) *CertificateQueryResponseDataCertificatesV2ItemPeriodCard {
	s.PeriodType = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemReserveInfo struct {
	OrderReserveUserInfoList []*CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemReserveInfo) SetOrderReserveUserInfoList(v []*CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) *CertificateQueryResponseDataCertificatesV2ItemReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem struct {
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificateQueryResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemSku struct {
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemSku) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemSku) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetAccountId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.AccountId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetMarketPrice(v int64) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.MarketPrice = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetThirdSkuId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetSkuId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.SkuId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetSuplierProductOutId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetSkuOutId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.SkuOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetGrouponType(v int) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.GrouponType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetProductId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.ProductId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetVoucherType(v int) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.VoucherType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetProductOutId(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.ProductOutId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetSoldStartTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemSku) SetTitle(v string) *CertificateQueryResponseDataCertificatesV2ItemSku {
	s.Title = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemTimeCard struct {
	SerialAmountList []*CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
	TimeCardType     *int                                                                          `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
	TimesCount       *int32                                                                        `json:"times_count,omitempty" xml:"times_count,omitempty"`
	TimesUsed        *int32                                                                        `json:"times_used,omitempty" xml:"times_used,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCard) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCard) SetSerialAmountList(v []*CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) *CertificateQueryResponseDataCertificatesV2ItemTimeCard {
	s.SerialAmountList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCard) SetTimeCardType(v int) *CertificateQueryResponseDataCertificatesV2ItemTimeCard {
	s.TimeCardType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCard) SetTimesCount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCard {
	s.TimesCount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCard) SetTimesUsed(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCard {
	s.TimesUsed = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem struct {
	Amount     *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                            `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) SetAmount(v *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount struct {
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificateQueryResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo struct {
	TimePeriodList []*CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
	UseTimeType    *int                                                                           `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo) SetTimePeriodList(v []*CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo) SetUseTimeType(v int) *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem struct {
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificateQueryResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemVerify struct {
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemVerify) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemVerify) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetPoiId(v int64) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.PoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetTimesCardSerialNum(v int32) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetVerifierUniqueId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetVerifyId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.VerifyId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetVerifyTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.VerifyTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetVerifyType(v int) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.VerifyType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetCanCancel(v bool) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.CanCancel = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerify) SetCertificateId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerify {
	s.CertificateId = &v
	return s
}

type CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem struct {
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyTime(v int64) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyType(v int) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetCanCancel(v bool) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetCertificateId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetPoiId(v int64) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.PoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifierUniqueId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyId(v string) *CertificateQueryResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyId = &v
	return s
}

type CertificateQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CertificateQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseExtra) SetSubDescription(v string) *CertificateQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateQueryResponseExtra) SetSubErrorCode(v int32) *CertificateQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificateQueryResponseExtra) SetDescription(v string) *CertificateQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificateQueryResponseExtra) SetErrorCode(v int32) *CertificateQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateQueryResponseExtra) SetLogid(v string) *CertificateQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertificateQueryResponseExtra) SetNow(v int64) *CertificateQueryResponseExtra {
	s.Now = &v
	return s
}

type CertificateVerifyRequest struct {
	EncryptedCodes   []*string                                       `json:"encrypted_codes,omitempty" xml:"encrypted_codes,omitempty" type:"Repeated"`
	PoiId            *string                                         `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	Header           map[string]*string                              `json:"header,omitempty" xml:"header,omitempty"`
	VerifyExtra      *CertificateVerifyRequestVerifyExtra            `json:"verify_extra,omitempty" xml:"verify_extra,omitempty"`
	Codes            []*string                                       `json:"codes,omitempty" xml:"codes,omitempty" type:"Repeated"`
	Vouchers         []*CertificateVerifyRequestVouchersItem         `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	AccountId        *string                                         `json:"account_id,omitempty" xml:"account_id,omitempty"`
	VerifySignList   []*string                                       `json:"verify_sign_list,omitempty" xml:"verify_sign_list,omitempty" type:"Repeated"`
	VerifyToken      *string                                         `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	AccessToken      *string                                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Voucher          *CertificateVerifyRequestVoucher                `json:"voucher,omitempty" xml:"voucher,omitempty"`
	OrderId          *string                                         `json:"order_id,omitempty" xml:"order_id,omitempty"`
	CodeWithTimeList []*CertificateVerifyRequestCodeWithTimeListItem `json:"code_with_time_list,omitempty" xml:"code_with_time_list,omitempty" type:"Repeated"`
}

func (s CertificateVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequest) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequest) SetEncryptedCodes(v []*string) *CertificateVerifyRequest {
	s.EncryptedCodes = v
	return s
}

func (s *CertificateVerifyRequest) SetPoiId(v string) *CertificateVerifyRequest {
	s.PoiId = &v
	return s
}

func (s *CertificateVerifyRequest) SetHeader(v map[string]*string) *CertificateVerifyRequest {
	s.Header = v
	return s
}

func (s *CertificateVerifyRequest) SetVerifyExtra(v *CertificateVerifyRequestVerifyExtra) *CertificateVerifyRequest {
	s.VerifyExtra = v
	return s
}

func (s *CertificateVerifyRequest) SetCodes(v []*string) *CertificateVerifyRequest {
	s.Codes = v
	return s
}

func (s *CertificateVerifyRequest) SetVouchers(v []*CertificateVerifyRequestVouchersItem) *CertificateVerifyRequest {
	s.Vouchers = v
	return s
}

func (s *CertificateVerifyRequest) SetAccountId(v string) *CertificateVerifyRequest {
	s.AccountId = &v
	return s
}

func (s *CertificateVerifyRequest) SetVerifySignList(v []*string) *CertificateVerifyRequest {
	s.VerifySignList = v
	return s
}

func (s *CertificateVerifyRequest) SetVerifyToken(v string) *CertificateVerifyRequest {
	s.VerifyToken = &v
	return s
}

func (s *CertificateVerifyRequest) SetAccessToken(v string) *CertificateVerifyRequest {
	s.AccessToken = &v
	return s
}

func (s *CertificateVerifyRequest) SetVoucher(v *CertificateVerifyRequestVoucher) *CertificateVerifyRequest {
	s.Voucher = v
	return s
}

func (s *CertificateVerifyRequest) SetOrderId(v string) *CertificateVerifyRequest {
	s.OrderId = &v
	return s
}

func (s *CertificateVerifyRequest) SetCodeWithTimeList(v []*CertificateVerifyRequestCodeWithTimeListItem) *CertificateVerifyRequest {
	s.CodeWithTimeList = v
	return s
}

type CertificateVerifyRequestCodeWithTimeListItem struct {
	VerifyTime *int64                                                 `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	Code       *string                                                `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	OuterNumb  *CertificateVerifyRequestCodeWithTimeListItemOuterNumb `json:"outer_numb,omitempty" xml:"outer_numb,omitempty"`
	SerialNum  *int32                                                 `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
}

func (s CertificateVerifyRequestCodeWithTimeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestCodeWithTimeListItem) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestCodeWithTimeListItem) SetVerifyTime(v int64) *CertificateVerifyRequestCodeWithTimeListItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateVerifyRequestCodeWithTimeListItem) SetCode(v string) *CertificateVerifyRequestCodeWithTimeListItem {
	s.Code = &v
	return s
}

func (s *CertificateVerifyRequestCodeWithTimeListItem) SetOuterNumb(v *CertificateVerifyRequestCodeWithTimeListItemOuterNumb) *CertificateVerifyRequestCodeWithTimeListItem {
	s.OuterNumb = v
	return s
}

func (s *CertificateVerifyRequestCodeWithTimeListItem) SetSerialNum(v int32) *CertificateVerifyRequestCodeWithTimeListItem {
	s.SerialNum = &v
	return s
}

type CertificateVerifyRequestCodeWithTimeListItemOuterNumb struct {
	CouponNumber *string `json:"coupon_number,omitempty" xml:"coupon_number,omitempty"`
	OrderNumber  *string `json:"order_number,omitempty" xml:"order_number,omitempty"`
}

func (s CertificateVerifyRequestCodeWithTimeListItemOuterNumb) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestCodeWithTimeListItemOuterNumb) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestCodeWithTimeListItemOuterNumb) SetCouponNumber(v string) *CertificateVerifyRequestCodeWithTimeListItemOuterNumb {
	s.CouponNumber = &v
	return s
}

func (s *CertificateVerifyRequestCodeWithTimeListItemOuterNumb) SetOrderNumber(v string) *CertificateVerifyRequestCodeWithTimeListItemOuterNumb {
	s.OrderNumber = &v
	return s
}

type CertificateVerifyRequestVerifyExtra struct {
	TotalVerify         *bool                                                   `json:"total_verify,omitempty" xml:"total_verify,omitempty"`
	DynamicCouponInfo   *CertificateVerifyRequestVerifyExtraDynamicCouponInfo   `json:"dynamic_coupon_info,omitempty" xml:"dynamic_coupon_info,omitempty"`
	OfflineAddPriceInfo *CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo `json:"offline_add_price_info,omitempty" xml:"offline_add_price_info,omitempty"`
	OutGoodIds          []*string                                               `json:"out_good_ids,omitempty" xml:"out_good_ids,omitempty" type:"Repeated"`
}

func (s CertificateVerifyRequestVerifyExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestVerifyExtra) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestVerifyExtra) SetTotalVerify(v bool) *CertificateVerifyRequestVerifyExtra {
	s.TotalVerify = &v
	return s
}

func (s *CertificateVerifyRequestVerifyExtra) SetDynamicCouponInfo(v *CertificateVerifyRequestVerifyExtraDynamicCouponInfo) *CertificateVerifyRequestVerifyExtra {
	s.DynamicCouponInfo = v
	return s
}

func (s *CertificateVerifyRequestVerifyExtra) SetOfflineAddPriceInfo(v *CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) *CertificateVerifyRequestVerifyExtra {
	s.OfflineAddPriceInfo = v
	return s
}

func (s *CertificateVerifyRequestVerifyExtra) SetOutGoodIds(v []*string) *CertificateVerifyRequestVerifyExtra {
	s.OutGoodIds = v
	return s
}

type CertificateVerifyRequestVerifyExtraDynamicCouponInfo struct {
	ActualDeductionAmount *int64 `json:"actual_deduction_amount,omitempty" xml:"actual_deduction_amount,omitempty"`
	BizTime               *int64 `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
}

func (s CertificateVerifyRequestVerifyExtraDynamicCouponInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestVerifyExtraDynamicCouponInfo) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestVerifyExtraDynamicCouponInfo) SetActualDeductionAmount(v int64) *CertificateVerifyRequestVerifyExtraDynamicCouponInfo {
	s.ActualDeductionAmount = &v
	return s
}

func (s *CertificateVerifyRequestVerifyExtraDynamicCouponInfo) SetBizTime(v int64) *CertificateVerifyRequestVerifyExtraDynamicCouponInfo {
	s.BizTime = &v
	return s
}

type CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo struct {
	IsOfflineAddPrice *bool `json:"is_offline_add_price,omitempty" xml:"is_offline_add_price,omitempty"`
}

func (s CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo) SetIsOfflineAddPrice(v bool) *CertificateVerifyRequestVerifyExtraOfflineAddPriceInfo {
	s.IsOfflineAddPrice = &v
	return s
}

type CertificateVerifyRequestVoucher struct {
	ProjectId         *string   `json:"project_id,omitempty" xml:"project_id,omitempty"`
	QrcodeList        []*string `json:"qrcode_list,omitempty" xml:"qrcode_list,omitempty" type:"Repeated"`
	VerifyTime        *int64    `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	CertificateNoList []*string `json:"certificate_no_list,omitempty" xml:"certificate_no_list,omitempty" type:"Repeated"`
	IdCardList        []*string `json:"id_card_list,omitempty" xml:"id_card_list,omitempty" type:"Repeated"`
}

func (s CertificateVerifyRequestVoucher) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestVoucher) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestVoucher) SetProjectId(v string) *CertificateVerifyRequestVoucher {
	s.ProjectId = &v
	return s
}

func (s *CertificateVerifyRequestVoucher) SetQrcodeList(v []*string) *CertificateVerifyRequestVoucher {
	s.QrcodeList = v
	return s
}

func (s *CertificateVerifyRequestVoucher) SetVerifyTime(v int64) *CertificateVerifyRequestVoucher {
	s.VerifyTime = &v
	return s
}

func (s *CertificateVerifyRequestVoucher) SetCertificateNoList(v []*string) *CertificateVerifyRequestVoucher {
	s.CertificateNoList = v
	return s
}

func (s *CertificateVerifyRequestVoucher) SetIdCardList(v []*string) *CertificateVerifyRequestVoucher {
	s.IdCardList = v
	return s
}

type CertificateVerifyRequestVouchersItem struct {
	VerifyTime        *int64    `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	CertificateNoList []*string `json:"certificate_no_list,omitempty" xml:"certificate_no_list,omitempty" type:"Repeated"`
	IdCardList        []*string `json:"id_card_list,omitempty" xml:"id_card_list,omitempty" type:"Repeated"`
	ProjectId         *string   `json:"project_id,omitempty" xml:"project_id,omitempty"`
	QrcodeList        []*string `json:"qrcode_list,omitempty" xml:"qrcode_list,omitempty" type:"Repeated"`
}

func (s CertificateVerifyRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *CertificateVerifyRequestVouchersItem) SetVerifyTime(v int64) *CertificateVerifyRequestVouchersItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateVerifyRequestVouchersItem) SetCertificateNoList(v []*string) *CertificateVerifyRequestVouchersItem {
	s.CertificateNoList = v
	return s
}

func (s *CertificateVerifyRequestVouchersItem) SetIdCardList(v []*string) *CertificateVerifyRequestVouchersItem {
	s.IdCardList = v
	return s
}

func (s *CertificateVerifyRequestVouchersItem) SetProjectId(v string) *CertificateVerifyRequestVouchersItem {
	s.ProjectId = &v
	return s
}

func (s *CertificateVerifyRequestVouchersItem) SetQrcodeList(v []*string) *CertificateVerifyRequestVouchersItem {
	s.QrcodeList = v
	return s
}

type CertificateVerifyResponse struct {
	Extra *CertificateVerifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *CertificateVerifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CertificateVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponse) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponse) SetExtra(v *CertificateVerifyResponseExtra) *CertificateVerifyResponse {
	s.Extra = v
	return s
}

func (s *CertificateVerifyResponse) SetData(v *CertificateVerifyResponseData) *CertificateVerifyResponse {
	s.Data = v
	return s
}

type CertificateVerifyResponseData struct {
	GwDescription *string                                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	VerifyResults []*CertificateVerifyResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CertificateVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseData) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseData) SetGwDescription(v string) *CertificateVerifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *CertificateVerifyResponseData) SetVerifyResults(v []*CertificateVerifyResponseDataVerifyResultsItem) *CertificateVerifyResponseData {
	s.VerifyResults = v
	return s
}

func (s *CertificateVerifyResponseData) SetGwErrorCode(v int32) *CertificateVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

type CertificateVerifyResponseDataVerifyResultsItem struct {
	Qrcode           *string                                                         `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
	ProductId        *string                                                         `json:"product_id,omitempty" xml:"product_id,omitempty"`
	CertificateId    *string                                                         `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Result           *int32                                                          `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	VerifyAmountInfo *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo `json:"verify_amount_info,omitempty" xml:"verify_amount_info,omitempty"`
	OutGoodId        *string                                                         `json:"out_good_id,omitempty" xml:"out_good_id,omitempty"`
	OriginCode       *string                                                         `json:"origin_code,omitempty" xml:"origin_code,omitempty"`
	Msg              *string                                                         `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	VerifyId         *string                                                         `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	AccountId        *string                                                         `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OrderId          *string                                                         `json:"order_id,omitempty" xml:"order_id,omitempty"`
	CertificateNo    *string                                                         `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	IdCard           *string                                                         `json:"id_card,omitempty" xml:"id_card,omitempty"`
	Code             *string                                                         `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CertificateVerifyResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetQrcode(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.Qrcode = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetProductId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.ProductId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetCertificateId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetResult(v int32) *CertificateVerifyResponseDataVerifyResultsItem {
	s.Result = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetVerifyAmountInfo(v *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo) *CertificateVerifyResponseDataVerifyResultsItem {
	s.VerifyAmountInfo = v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetOutGoodId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.OutGoodId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetOriginCode(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.OriginCode = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetMsg(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.Msg = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetVerifyId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetAccountId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.AccountId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetOrderId(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.OrderId = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetCertificateNo(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.CertificateNo = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetIdCard(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.IdCard = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItem) SetCode(v string) *CertificateVerifyResponseDataVerifyResultsItem {
	s.Code = &v
	return s
}

type CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo struct {
	TimesCardSerialAmount *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount `json:"times_card_serial_amount,omitempty" xml:"times_card_serial_amount,omitempty"`
	TimeCardAmount        *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount        `json:"time_card_amount,omitempty" xml:"time_card_amount,omitempty"`
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo) SetTimesCardSerialAmount(v *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo {
	s.TimesCardSerialAmount = v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo) SetTimeCardAmount(v *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfo {
	s.TimeCardAmount = v
	return s
}

type CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) SetAmount(v int64) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimeCardAmount {
	s.Amount = &v
	return s
}

type CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount struct {
	Amount     *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                                     `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) SetAmount(v *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount {
	s.Amount = v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) SetSerialNumb(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount {
	s.SerialNumb = &v
	return s
}

type CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount struct {
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalCurrency(v string) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPaymentDiscountAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetCouponPayAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginListMarketAmount(v int64) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPlatformDiscountAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetListMarketAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetMerchantTicketAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPayAmount(v int32) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetBrandTicketAmount(v int64) *CertificateVerifyResponseDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.BrandTicketAmount = &v
	return s
}

type CertificateVerifyResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CertificateVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateVerifyResponseExtra) SetErrorCode(v int32) *CertificateVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateVerifyResponseExtra) SetLogid(v string) *CertificateVerifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertificateVerifyResponseExtra) SetNow(v int64) *CertificateVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificateVerifyResponseExtra) SetSubDescription(v string) *CertificateVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateVerifyResponseExtra) SetSubErrorCode(v int32) *CertificateVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificateVerifyResponseExtra) SetDescription(v string) *CertificateVerifyResponseExtra {
	s.Description = &v
	return s
}

type ChangeUserBindAgentRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DouyinId    *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	OldAgentId  *int64             `json:"old_agent_id,omitempty" xml:"old_agent_id,omitempty" require:"true"`
	NewAgentId  *int64             `json:"new_agent_id,omitempty" xml:"new_agent_id,omitempty" require:"true"`
}

func (s ChangeUserBindAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeUserBindAgentRequest) GoString() string {
	return s.String()
}

func (s *ChangeUserBindAgentRequest) SetHeader(v map[string]*string) *ChangeUserBindAgentRequest {
	s.Header = v
	return s
}

func (s *ChangeUserBindAgentRequest) SetAccessToken(v string) *ChangeUserBindAgentRequest {
	s.AccessToken = &v
	return s
}

func (s *ChangeUserBindAgentRequest) SetDouyinId(v string) *ChangeUserBindAgentRequest {
	s.DouyinId = &v
	return s
}

func (s *ChangeUserBindAgentRequest) SetOldAgentId(v int64) *ChangeUserBindAgentRequest {
	s.OldAgentId = &v
	return s
}

func (s *ChangeUserBindAgentRequest) SetNewAgentId(v int64) *ChangeUserBindAgentRequest {
	s.NewAgentId = &v
	return s
}

type ChangeUserBindAgentResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ChangeUserBindAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeUserBindAgentResponse) GoString() string {
	return s.String()
}

func (s *ChangeUserBindAgentResponse) SetErrMsg(v string) *ChangeUserBindAgentResponse {
	s.ErrMsg = &v
	return s
}

func (s *ChangeUserBindAgentResponse) SetLogId(v string) *ChangeUserBindAgentResponse {
	s.LogId = &v
	return s
}

func (s *ChangeUserBindAgentResponse) SetErrNo(v int32) *ChangeUserBindAgentResponse {
	s.ErrNo = &v
	return s
}

type ClaimQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiIds      []*string          `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" require:"true" type:"Repeated"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s ClaimQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryRequest) GoString() string {
	return s.String()
}

func (s *ClaimQueryRequest) SetHeader(v map[string]*string) *ClaimQueryRequest {
	s.Header = v
	return s
}

func (s *ClaimQueryRequest) SetAccessToken(v string) *ClaimQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *ClaimQueryRequest) SetPoiIds(v []*string) *ClaimQueryRequest {
	s.PoiIds = v
	return s
}

func (s *ClaimQueryRequest) SetAccountId(v string) *ClaimQueryRequest {
	s.AccountId = &v
	return s
}

type ClaimQueryResponse struct {
	Extra *ClaimQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *ClaimQueryResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ClaimQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponse) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponse) SetExtra(v *ClaimQueryResponseExtra) *ClaimQueryResponse {
	s.Extra = v
	return s
}

func (s *ClaimQueryResponse) SetData(v *ClaimQueryResponseData) *ClaimQueryResponse {
	s.Data = v
	return s
}

type ClaimQueryResponseData struct {
	TaskResults   []*ClaimQueryResponseDataTaskResultsItem `json:"task_results,omitempty" xml:"task_results,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ClaimQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseData) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseData) SetTaskResults(v []*ClaimQueryResponseDataTaskResultsItem) *ClaimQueryResponseData {
	s.TaskResults = v
	return s
}

func (s *ClaimQueryResponseData) SetGwErrorCode(v int32) *ClaimQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ClaimQueryResponseData) SetGwDescription(v string) *ClaimQueryResponseData {
	s.GwDescription = &v
	return s
}

type ClaimQueryResponseDataTaskResultsItem struct {
	ClaimPoiResult          *ClaimQueryResponseDataTaskResultsItemClaimPoiResult          `json:"claim_poi_result,omitempty" xml:"claim_poi_result,omitempty"`
	EditQualificationResult *ClaimQueryResponseDataTaskResultsItemEditQualificationResult `json:"edit_qualification_result,omitempty" xml:"edit_qualification_result,omitempty"`
	PoiId                   *string                                                       `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ClaimQueryResponseDataTaskResultsItem) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseDataTaskResultsItem) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseDataTaskResultsItem) SetClaimPoiResult(v *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) *ClaimQueryResponseDataTaskResultsItem {
	s.ClaimPoiResult = v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItem) SetEditQualificationResult(v *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) *ClaimQueryResponseDataTaskResultsItem {
	s.EditQualificationResult = v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItem) SetPoiId(v string) *ClaimQueryResponseDataTaskResultsItem {
	s.PoiId = &v
	return s
}

type ClaimQueryResponseDataTaskResultsItemClaimPoiResult struct {
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
	TaskStatus   *int    `json:"task_status,omitempty" xml:"task_status,omitempty"`
	AccountId    *int64  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId        *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ClaimQueryResponseDataTaskResultsItemClaimPoiResult) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseDataTaskResultsItemClaimPoiResult) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) SetRejectReason(v string) *ClaimQueryResponseDataTaskResultsItemClaimPoiResult {
	s.RejectReason = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) SetStatus(v int) *ClaimQueryResponseDataTaskResultsItemClaimPoiResult {
	s.Status = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) SetTaskStatus(v int) *ClaimQueryResponseDataTaskResultsItemClaimPoiResult {
	s.TaskStatus = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) SetAccountId(v int64) *ClaimQueryResponseDataTaskResultsItemClaimPoiResult {
	s.AccountId = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemClaimPoiResult) SetPoiId(v int64) *ClaimQueryResponseDataTaskResultsItemClaimPoiResult {
	s.PoiId = &v
	return s
}

type ClaimQueryResponseDataTaskResultsItemEditQualificationResult struct {
	SubjectLegalResult *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult `json:"subject_legal_result,omitempty" xml:"subject_legal_result,omitempty"`
	TaskStatus         *int                                                                            `json:"task_status,omitempty" xml:"task_status,omitempty"`
	AccountId          *int64                                                                          `json:"account_id,omitempty" xml:"account_id,omitempty"`
	IndustryResult     *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult     `json:"industry_result,omitempty" xml:"industry_result,omitempty"`
	PoiId              *int64                                                                          `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResult) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResult) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) SetSubjectLegalResult(v *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult) *ClaimQueryResponseDataTaskResultsItemEditQualificationResult {
	s.SubjectLegalResult = v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) SetTaskStatus(v int) *ClaimQueryResponseDataTaskResultsItemEditQualificationResult {
	s.TaskStatus = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) SetAccountId(v int64) *ClaimQueryResponseDataTaskResultsItemEditQualificationResult {
	s.AccountId = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) SetIndustryResult(v *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult) *ClaimQueryResponseDataTaskResultsItemEditQualificationResult {
	s.IndustryResult = v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResult) SetPoiId(v int64) *ClaimQueryResponseDataTaskResultsItemEditQualificationResult {
	s.PoiId = &v
	return s
}

type ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult struct {
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult) SetRejectReason(v string) *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult {
	s.RejectReason = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult) SetStatus(v int) *ClaimQueryResponseDataTaskResultsItemEditQualificationResultIndustryResult {
	s.Status = &v
	return s
}

type ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult struct {
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult) SetRejectReason(v string) *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult {
	s.RejectReason = &v
	return s
}

func (s *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult) SetStatus(v int) *ClaimQueryResponseDataTaskResultsItemEditQualificationResultSubjectLegalResult {
	s.Status = &v
	return s
}

type ClaimQueryResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s ClaimQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ClaimQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ClaimQueryResponseExtra) SetDescription(v string) *ClaimQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ClaimQueryResponseExtra) SetErrorCode(v int32) *ClaimQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ClaimQueryResponseExtra) SetLogid(v string) *ClaimQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ClaimQueryResponseExtra) SetNow(v int64) *ClaimQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *ClaimQueryResponseExtra) SetSubDescription(v string) *ClaimQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ClaimQueryResponseExtra) SetSubErrorCode(v int32) *ClaimQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ClueQueryRequest struct {
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	PageSize      *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	EndTime       *string            `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	Page          *int32             `json:"page,omitempty" xml:"page,omitempty" require:"true"`
	StartTime     *string            `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	LifeAccountId []*string          `json:"life_account_id,omitempty" xml:"life_account_id,omitempty" type:"Repeated"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ClueQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryRequest) GoString() string {
	return s.String()
}

func (s *ClueQueryRequest) SetOpenId(v string) *ClueQueryRequest {
	s.OpenId = &v
	return s
}

func (s *ClueQueryRequest) SetPageSize(v int32) *ClueQueryRequest {
	s.PageSize = &v
	return s
}

func (s *ClueQueryRequest) SetAccountId(v string) *ClueQueryRequest {
	s.AccountId = &v
	return s
}

func (s *ClueQueryRequest) SetEndTime(v string) *ClueQueryRequest {
	s.EndTime = &v
	return s
}

func (s *ClueQueryRequest) SetPage(v int32) *ClueQueryRequest {
	s.Page = &v
	return s
}

func (s *ClueQueryRequest) SetStartTime(v string) *ClueQueryRequest {
	s.StartTime = &v
	return s
}

func (s *ClueQueryRequest) SetLifeAccountId(v []*string) *ClueQueryRequest {
	s.LifeAccountId = v
	return s
}

func (s *ClueQueryRequest) SetHeader(v map[string]*string) *ClueQueryRequest {
	s.Header = v
	return s
}

func (s *ClueQueryRequest) SetAccessToken(v string) *ClueQueryRequest {
	s.AccessToken = &v
	return s
}

type ClueQueryResponse struct {
	Data  *ClueQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *ClueQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s ClueQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryResponse) GoString() string {
	return s.String()
}

func (s *ClueQueryResponse) SetData(v *ClueQueryResponseData) *ClueQueryResponse {
	s.Data = v
	return s
}

func (s *ClueQueryResponse) SetExtra(v *ClueQueryResponseExtra) *ClueQueryResponse {
	s.Extra = v
	return s
}

type ClueQueryResponseData struct {
	ClueData      []*ClueQueryResponseDataClueDataItem `json:"clue_data,omitempty" xml:"clue_data,omitempty" type:"Repeated"`
	Page          *ClueQueryResponseDataPage           `json:"page,omitempty" xml:"page,omitempty"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ClueQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryResponseData) GoString() string {
	return s.String()
}

func (s *ClueQueryResponseData) SetClueData(v []*ClueQueryResponseDataClueDataItem) *ClueQueryResponseData {
	s.ClueData = v
	return s
}

func (s *ClueQueryResponseData) SetPage(v *ClueQueryResponseDataPage) *ClueQueryResponseData {
	s.Page = v
	return s
}

func (s *ClueQueryResponseData) SetGwErrorCode(v int32) *ClueQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *ClueQueryResponseData) SetGwDescription(v string) *ClueQueryResponseData {
	s.GwDescription = &v
	return s
}

type ClueQueryResponseDataClueDataItem struct {
	ProductName              *string   `json:"product_name,omitempty" xml:"product_name,omitempty"`
	FlowEntrance             *int32    `json:"flow_entrance,omitempty" xml:"flow_entrance,omitempty"`
	EffectiveState           *int32    `json:"effective_state,omitempty" xml:"effective_state,omitempty"`
	AutoProvinceName         *string   `json:"auto_province_name,omitempty" xml:"auto_province_name,omitempty"`
	Business                 *string   `json:"business,omitempty" xml:"business,omitempty"`
	Telephone                *string   `json:"telephone,omitempty" xml:"telephone,omitempty"`
	StaffDouyinId            *string   `json:"staff_douyin_id,omitempty" xml:"staff_douyin_id,omitempty"`
	IntentionLifeAccountName *string   `json:"intention_life_account_name,omitempty" xml:"intention_life_account_name,omitempty"`
	FollowLifeAccountId      *string   `json:"follow_life_account_id,omitempty" xml:"follow_life_account_id,omitempty"`
	RootLifeAccountId        *string   `json:"root_life_account_id,omitempty" xml:"root_life_account_id,omitempty"`
	Age                      *int32    `json:"age,omitempty" xml:"age,omitempty"`
	AdType                   *int32    `json:"ad_type,omitempty" xml:"ad_type,omitempty"`
	OrderStatus              *string   `json:"order_status,omitempty" xml:"order_status,omitempty"`
	AutoCityName             *string   `json:"auto_city_name,omitempty" xml:"auto_city_name,omitempty"`
	SourceCraftsmanNickname  *string   `json:"source_craftsman_nickname,omitempty" xml:"source_craftsman_nickname,omitempty"`
	CityName                 *string   `json:"city_name,omitempty" xml:"city_name,omitempty"`
	SourceCraftsmanDouyinId  *string   `json:"source_craftsman_douyin_id,omitempty" xml:"source_craftsman_douyin_id,omitempty"`
	Remark                   *string   `json:"remark,omitempty" xml:"remark,omitempty"`
	ClueId                   *string   `json:"clue_id,omitempty" xml:"clue_id,omitempty"`
	ModifyTime               *string   `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	Name                     *string   `json:"name,omitempty" xml:"name,omitempty"`
	ImageId                  *int64    `json:"image_id,omitempty" xml:"image_id,omitempty"`
	FollowLifeAccountType    *int32    `json:"follow_life_account_type,omitempty" xml:"follow_life_account_type,omitempty"`
	ToolId                   *string   `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	VideoId                  *int64    `json:"video_id,omitempty" xml:"video_id,omitempty"`
	AuthorDouyinId           *string   `json:"author_douyin_id,omitempty" xml:"author_douyin_id,omitempty"`
	FollowPoiId              *string   `json:"follow_poi_id,omitempty" xml:"follow_poi_id,omitempty"`
	ProductType              *int32    `json:"product_type,omitempty" xml:"product_type,omitempty"`
	Gender                   *int32    `json:"gender,omitempty" xml:"gender,omitempty"`
	ClueType                 *int32    `json:"clue_type,omitempty" xml:"clue_type,omitempty"`
	StaffNickname            *string   `json:"staff_nickname,omitempty" xml:"staff_nickname,omitempty"`
	CountyName               *string   `json:"county_name,omitempty" xml:"county_name,omitempty"`
	ProvinceName             *string   `json:"province_name,omitempty" xml:"province_name,omitempty"`
	FollowStateName          *int32    `json:"follow_state_name,omitempty" xml:"follow_state_name,omitempty"`
	AdId                     *int64    `json:"ad_id,omitempty" xml:"ad_id,omitempty"`
	AuthorNickname           *string   `json:"author_nickname,omitempty" xml:"author_nickname,omitempty"`
	AdvertiserName           *string   `json:"advertiser_name,omitempty" xml:"advertiser_name,omitempty"`
	StaffCommerceNickname    *string   `json:"staff_commerce_nickname,omitempty" xml:"staff_commerce_nickname,omitempty"`
	ConvertStatus            *int32    `json:"convert_status,omitempty" xml:"convert_status,omitempty"`
	IntentionPoiId           *string   `json:"intention_poi_id,omitempty" xml:"intention_poi_id,omitempty"`
	IsPrivateClue            *int32    `json:"is_private_clue,omitempty" xml:"is_private_clue,omitempty"`
	TelAddr                  *string   `json:"tel_addr,omitempty" xml:"tel_addr,omitempty"`
	Tags                     []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	PromotionId              *int64    `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	Address                  *string   `json:"address,omitempty" xml:"address,omitempty"`
	RemarkDict               *string   `json:"remark_dict,omitempty" xml:"remark_dict,omitempty"`
	AdvertiserId             *string   `json:"advertiser_id,omitempty" xml:"advertiser_id,omitempty"`
	CreateTimeDetail         *string   `json:"create_time_detail,omitempty" xml:"create_time_detail,omitempty"`
	AllocationStatus         *int32    `json:"allocation_status,omitempty" xml:"allocation_status,omitempty"`
	OrderId                  *int64    `json:"order_id,omitempty" xml:"order_id,omitempty"`
	AuthorRole               *string   `json:"author_role,omitempty" xml:"author_role,omitempty"`
	LeadsPage                *int32    `json:"leads_page,omitempty" xml:"leads_page,omitempty"`
	FollowLifeAccountName    *string   `json:"follow_life_account_name,omitempty" xml:"follow_life_account_name,omitempty"`
	PromotionName            *string   `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	ClueOwnerName            *string   `json:"clue_owner_name,omitempty" xml:"clue_owner_name,omitempty"`
	ReqId                    *string   `json:"req_id,omitempty" xml:"req_id,omitempty"`
	ActionType               *int32    `json:"action_type,omitempty" xml:"action_type,omitempty"`
	FlowType                 *int32    `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	ContentId                *string   `json:"content_id,omitempty" xml:"content_id,omitempty"`
	Weixin                   *string   `json:"weixin,omitempty" xml:"weixin,omitempty"`
	ProductId                *string   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SystemTags               []*string `json:"system_tags,omitempty" xml:"system_tags,omitempty" type:"Repeated"`
	TitleId                  *int64    `json:"title_id,omitempty" xml:"title_id,omitempty"`
}

func (s ClueQueryResponseDataClueDataItem) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryResponseDataClueDataItem) GoString() string {
	return s.String()
}

func (s *ClueQueryResponseDataClueDataItem) SetProductName(v string) *ClueQueryResponseDataClueDataItem {
	s.ProductName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFlowEntrance(v int32) *ClueQueryResponseDataClueDataItem {
	s.FlowEntrance = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetEffectiveState(v int32) *ClueQueryResponseDataClueDataItem {
	s.EffectiveState = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAutoProvinceName(v string) *ClueQueryResponseDataClueDataItem {
	s.AutoProvinceName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetBusiness(v string) *ClueQueryResponseDataClueDataItem {
	s.Business = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetTelephone(v string) *ClueQueryResponseDataClueDataItem {
	s.Telephone = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetStaffDouyinId(v string) *ClueQueryResponseDataClueDataItem {
	s.StaffDouyinId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetIntentionLifeAccountName(v string) *ClueQueryResponseDataClueDataItem {
	s.IntentionLifeAccountName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFollowLifeAccountId(v string) *ClueQueryResponseDataClueDataItem {
	s.FollowLifeAccountId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetRootLifeAccountId(v string) *ClueQueryResponseDataClueDataItem {
	s.RootLifeAccountId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAge(v int32) *ClueQueryResponseDataClueDataItem {
	s.Age = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAdType(v int32) *ClueQueryResponseDataClueDataItem {
	s.AdType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetOrderStatus(v string) *ClueQueryResponseDataClueDataItem {
	s.OrderStatus = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAutoCityName(v string) *ClueQueryResponseDataClueDataItem {
	s.AutoCityName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetSourceCraftsmanNickname(v string) *ClueQueryResponseDataClueDataItem {
	s.SourceCraftsmanNickname = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetCityName(v string) *ClueQueryResponseDataClueDataItem {
	s.CityName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetSourceCraftsmanDouyinId(v string) *ClueQueryResponseDataClueDataItem {
	s.SourceCraftsmanDouyinId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetRemark(v string) *ClueQueryResponseDataClueDataItem {
	s.Remark = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetClueId(v string) *ClueQueryResponseDataClueDataItem {
	s.ClueId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetModifyTime(v string) *ClueQueryResponseDataClueDataItem {
	s.ModifyTime = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetName(v string) *ClueQueryResponseDataClueDataItem {
	s.Name = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetImageId(v int64) *ClueQueryResponseDataClueDataItem {
	s.ImageId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFollowLifeAccountType(v int32) *ClueQueryResponseDataClueDataItem {
	s.FollowLifeAccountType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetToolId(v string) *ClueQueryResponseDataClueDataItem {
	s.ToolId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetVideoId(v int64) *ClueQueryResponseDataClueDataItem {
	s.VideoId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAuthorDouyinId(v string) *ClueQueryResponseDataClueDataItem {
	s.AuthorDouyinId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFollowPoiId(v string) *ClueQueryResponseDataClueDataItem {
	s.FollowPoiId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetProductType(v int32) *ClueQueryResponseDataClueDataItem {
	s.ProductType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetGender(v int32) *ClueQueryResponseDataClueDataItem {
	s.Gender = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetClueType(v int32) *ClueQueryResponseDataClueDataItem {
	s.ClueType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetStaffNickname(v string) *ClueQueryResponseDataClueDataItem {
	s.StaffNickname = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetCountyName(v string) *ClueQueryResponseDataClueDataItem {
	s.CountyName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetProvinceName(v string) *ClueQueryResponseDataClueDataItem {
	s.ProvinceName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFollowStateName(v int32) *ClueQueryResponseDataClueDataItem {
	s.FollowStateName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAdId(v int64) *ClueQueryResponseDataClueDataItem {
	s.AdId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAuthorNickname(v string) *ClueQueryResponseDataClueDataItem {
	s.AuthorNickname = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAdvertiserName(v string) *ClueQueryResponseDataClueDataItem {
	s.AdvertiserName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetStaffCommerceNickname(v string) *ClueQueryResponseDataClueDataItem {
	s.StaffCommerceNickname = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetConvertStatus(v int32) *ClueQueryResponseDataClueDataItem {
	s.ConvertStatus = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetIntentionPoiId(v string) *ClueQueryResponseDataClueDataItem {
	s.IntentionPoiId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetIsPrivateClue(v int32) *ClueQueryResponseDataClueDataItem {
	s.IsPrivateClue = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetTelAddr(v string) *ClueQueryResponseDataClueDataItem {
	s.TelAddr = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetTags(v []*string) *ClueQueryResponseDataClueDataItem {
	s.Tags = v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetPromotionId(v int64) *ClueQueryResponseDataClueDataItem {
	s.PromotionId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAddress(v string) *ClueQueryResponseDataClueDataItem {
	s.Address = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetRemarkDict(v string) *ClueQueryResponseDataClueDataItem {
	s.RemarkDict = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAdvertiserId(v string) *ClueQueryResponseDataClueDataItem {
	s.AdvertiserId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetCreateTimeDetail(v string) *ClueQueryResponseDataClueDataItem {
	s.CreateTimeDetail = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAllocationStatus(v int32) *ClueQueryResponseDataClueDataItem {
	s.AllocationStatus = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetOrderId(v int64) *ClueQueryResponseDataClueDataItem {
	s.OrderId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetAuthorRole(v string) *ClueQueryResponseDataClueDataItem {
	s.AuthorRole = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetLeadsPage(v int32) *ClueQueryResponseDataClueDataItem {
	s.LeadsPage = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFollowLifeAccountName(v string) *ClueQueryResponseDataClueDataItem {
	s.FollowLifeAccountName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetPromotionName(v string) *ClueQueryResponseDataClueDataItem {
	s.PromotionName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetClueOwnerName(v string) *ClueQueryResponseDataClueDataItem {
	s.ClueOwnerName = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetReqId(v string) *ClueQueryResponseDataClueDataItem {
	s.ReqId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetActionType(v int32) *ClueQueryResponseDataClueDataItem {
	s.ActionType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetFlowType(v int32) *ClueQueryResponseDataClueDataItem {
	s.FlowType = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetContentId(v string) *ClueQueryResponseDataClueDataItem {
	s.ContentId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetWeixin(v string) *ClueQueryResponseDataClueDataItem {
	s.Weixin = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetProductId(v string) *ClueQueryResponseDataClueDataItem {
	s.ProductId = &v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetSystemTags(v []*string) *ClueQueryResponseDataClueDataItem {
	s.SystemTags = v
	return s
}

func (s *ClueQueryResponseDataClueDataItem) SetTitleId(v int64) *ClueQueryResponseDataClueDataItem {
	s.TitleId = &v
	return s
}

type ClueQueryResponseDataPage struct {
	PageSize   *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	PageTotal  *int32 `json:"page_total,omitempty" xml:"page_total,omitempty"`
	Total      *int32 `json:"total,omitempty" xml:"total,omitempty"`
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
}

func (s ClueQueryResponseDataPage) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryResponseDataPage) GoString() string {
	return s.String()
}

func (s *ClueQueryResponseDataPage) SetPageSize(v int32) *ClueQueryResponseDataPage {
	s.PageSize = &v
	return s
}

func (s *ClueQueryResponseDataPage) SetPageTotal(v int32) *ClueQueryResponseDataPage {
	s.PageTotal = &v
	return s
}

func (s *ClueQueryResponseDataPage) SetTotal(v int32) *ClueQueryResponseDataPage {
	s.Total = &v
	return s
}

func (s *ClueQueryResponseDataPage) SetPageNumber(v int32) *ClueQueryResponseDataPage {
	s.PageNumber = &v
	return s
}

type ClueQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s ClueQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s ClueQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *ClueQueryResponseExtra) SetSubDescription(v string) *ClueQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *ClueQueryResponseExtra) SetSubErrorCode(v int32) *ClueQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *ClueQueryResponseExtra) SetDescription(v string) *ClueQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *ClueQueryResponseExtra) SetErrorCode(v int32) *ClueQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *ClueQueryResponseExtra) SetLogid(v string) *ClueQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *ClueQueryResponseExtra) SetNow(v int64) *ClueQueryResponseExtra {
	s.Now = &v
	return s
}

type CoGameUploadUserDataRequest struct {
	RoundId     *int64                                        `json:"round_id,omitempty" xml:"round_id,omitempty"`
	Header      map[string]*string                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RoundStatus *int64                                        `json:"round_status,omitempty" xml:"round_status,omitempty"`
	AnchorInfos []*CoGameUploadUserDataRequestAnchorInfosItem `json:"anchor_infos,omitempty" xml:"anchor_infos,omitempty" type:"Repeated"`
	UserList    []*CoGameUploadUserDataRequestUserListItem    `json:"user_list,omitempty" xml:"user_list,omitempty" type:"Repeated"`
	AppId       *string                                       `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

func (s CoGameUploadUserDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CoGameUploadUserDataRequest) GoString() string {
	return s.String()
}

func (s *CoGameUploadUserDataRequest) SetRoundId(v int64) *CoGameUploadUserDataRequest {
	s.RoundId = &v
	return s
}

func (s *CoGameUploadUserDataRequest) SetHeader(v map[string]*string) *CoGameUploadUserDataRequest {
	s.Header = v
	return s
}

func (s *CoGameUploadUserDataRequest) SetAccessToken(v string) *CoGameUploadUserDataRequest {
	s.AccessToken = &v
	return s
}

func (s *CoGameUploadUserDataRequest) SetRoundStatus(v int64) *CoGameUploadUserDataRequest {
	s.RoundStatus = &v
	return s
}

func (s *CoGameUploadUserDataRequest) SetAnchorInfos(v []*CoGameUploadUserDataRequestAnchorInfosItem) *CoGameUploadUserDataRequest {
	s.AnchorInfos = v
	return s
}

func (s *CoGameUploadUserDataRequest) SetUserList(v []*CoGameUploadUserDataRequestUserListItem) *CoGameUploadUserDataRequest {
	s.UserList = v
	return s
}

func (s *CoGameUploadUserDataRequest) SetAppId(v string) *CoGameUploadUserDataRequest {
	s.AppId = &v
	return s
}

type CoGameUploadUserDataRequestAnchorInfosItem struct {
	RoomId       *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AnchorOpenId *string `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty"`
}

func (s CoGameUploadUserDataRequestAnchorInfosItem) String() string {
	return tea.Prettify(s)
}

func (s CoGameUploadUserDataRequestAnchorInfosItem) GoString() string {
	return s.String()
}

func (s *CoGameUploadUserDataRequestAnchorInfosItem) SetRoomId(v string) *CoGameUploadUserDataRequestAnchorInfosItem {
	s.RoomId = &v
	return s
}

func (s *CoGameUploadUserDataRequestAnchorInfosItem) SetAnchorOpenId(v string) *CoGameUploadUserDataRequestAnchorInfosItem {
	s.AnchorOpenId = &v
	return s
}

type CoGameUploadUserDataRequestUserListItem struct {
	OpenId *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	Score  *int64  `json:"score,omitempty" xml:"score,omitempty"`
}

func (s CoGameUploadUserDataRequestUserListItem) String() string {
	return tea.Prettify(s)
}

func (s CoGameUploadUserDataRequestUserListItem) GoString() string {
	return s.String()
}

func (s *CoGameUploadUserDataRequestUserListItem) SetOpenId(v string) *CoGameUploadUserDataRequestUserListItem {
	s.OpenId = &v
	return s
}

func (s *CoGameUploadUserDataRequestUserListItem) SetScore(v int64) *CoGameUploadUserDataRequestUserListItem {
	s.Score = &v
	return s
}

type CoGameUploadUserDataResponse struct {
	ErrCode      *int64                                          `json:"err_code,omitempty" xml:"err_code,omitempty"`
	ErrMsg       *string                                         `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	FailRoomList []*CoGameUploadUserDataResponseFailRoomListItem `json:"fail_room_list,omitempty" xml:"fail_room_list,omitempty" type:"Repeated"`
}

func (s CoGameUploadUserDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CoGameUploadUserDataResponse) GoString() string {
	return s.String()
}

func (s *CoGameUploadUserDataResponse) SetErrCode(v int64) *CoGameUploadUserDataResponse {
	s.ErrCode = &v
	return s
}

func (s *CoGameUploadUserDataResponse) SetErrMsg(v string) *CoGameUploadUserDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *CoGameUploadUserDataResponse) SetFailRoomList(v []*CoGameUploadUserDataResponseFailRoomListItem) *CoGameUploadUserDataResponse {
	s.FailRoomList = v
	return s
}

type CoGameUploadUserDataResponseFailRoomListItem struct {
	Errmsg    *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	RoomIdStr *string `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
	RoomId    *int64  `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Errcode   *int64  `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s CoGameUploadUserDataResponseFailRoomListItem) String() string {
	return tea.Prettify(s)
}

func (s CoGameUploadUserDataResponseFailRoomListItem) GoString() string {
	return s.String()
}

func (s *CoGameUploadUserDataResponseFailRoomListItem) SetErrmsg(v string) *CoGameUploadUserDataResponseFailRoomListItem {
	s.Errmsg = &v
	return s
}

func (s *CoGameUploadUserDataResponseFailRoomListItem) SetRoomIdStr(v string) *CoGameUploadUserDataResponseFailRoomListItem {
	s.RoomIdStr = &v
	return s
}

func (s *CoGameUploadUserDataResponseFailRoomListItem) SetRoomId(v int64) *CoGameUploadUserDataResponseFailRoomListItem {
	s.RoomId = &v
	return s
}

func (s *CoGameUploadUserDataResponseFailRoomListItem) SetErrcode(v int64) *CoGameUploadUserDataResponseFailRoomListItem {
	s.Errcode = &v
	return s
}

type CommentQueryRequest struct {
	StartTime     *int64             `json:"start_time,omitempty" xml:"start_time,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	EndTime       *int64             `json:"end_time,omitempty" xml:"end_time,omitempty"`
	PoiIdList     []*int64           `json:"poi_id_list,omitempty" xml:"poi_id_list,omitempty" type:"Repeated"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Count         *int64             `json:"count,omitempty" xml:"count,omitempty"`
	Cursor        *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	ProductIdList []*int64           `json:"product_id_list,omitempty" xml:"product_id_list,omitempty" type:"Repeated"`
}

func (s CommentQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryRequest) GoString() string {
	return s.String()
}

func (s *CommentQueryRequest) SetStartTime(v int64) *CommentQueryRequest {
	s.StartTime = &v
	return s
}

func (s *CommentQueryRequest) SetAccountId(v string) *CommentQueryRequest {
	s.AccountId = &v
	return s
}

func (s *CommentQueryRequest) SetEndTime(v int64) *CommentQueryRequest {
	s.EndTime = &v
	return s
}

func (s *CommentQueryRequest) SetPoiIdList(v []*int64) *CommentQueryRequest {
	s.PoiIdList = v
	return s
}

func (s *CommentQueryRequest) SetHeader(v map[string]*string) *CommentQueryRequest {
	s.Header = v
	return s
}

func (s *CommentQueryRequest) SetAccessToken(v string) *CommentQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *CommentQueryRequest) SetCount(v int64) *CommentQueryRequest {
	s.Count = &v
	return s
}

func (s *CommentQueryRequest) SetCursor(v string) *CommentQueryRequest {
	s.Cursor = &v
	return s
}

func (s *CommentQueryRequest) SetProductIdList(v []*int64) *CommentQueryRequest {
	s.ProductIdList = v
	return s
}

type CommentQueryResponse struct {
	Extra *CommentQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CommentQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CommentQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponse) GoString() string {
	return s.String()
}

func (s *CommentQueryResponse) SetExtra(v *CommentQueryResponseExtra) *CommentQueryResponse {
	s.Extra = v
	return s
}

func (s *CommentQueryResponse) SetData(v *CommentQueryResponseData) *CommentQueryResponse {
	s.Data = v
	return s
}

type CommentQueryResponseData struct {
	Cursor        *string                                 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	GwErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	HasMore       *bool                                   `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Comments      []*CommentQueryResponseDataCommentsItem `json:"comments,omitempty" xml:"comments,omitempty" type:"Repeated"`
}

func (s CommentQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponseData) GoString() string {
	return s.String()
}

func (s *CommentQueryResponseData) SetCursor(v string) *CommentQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *CommentQueryResponseData) SetGwErrorCode(v int32) *CommentQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CommentQueryResponseData) SetGwDescription(v string) *CommentQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *CommentQueryResponseData) SetHasMore(v bool) *CommentQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *CommentQueryResponseData) SetComments(v []*CommentQueryResponseDataCommentsItem) *CommentQueryResponseData {
	s.Comments = v
	return s
}

type CommentQueryResponseDataCommentsItem struct {
	PoiId       *int64                                           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ProductInfo *CommentQueryResponseDataCommentsItemProductInfo `json:"product_info,omitempty" xml:"product_info,omitempty"`
	CommentInfo *CommentQueryResponseDataCommentsItemCommentInfo `json:"comment_info,omitempty" xml:"comment_info,omitempty"`
}

func (s CommentQueryResponseDataCommentsItem) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponseDataCommentsItem) GoString() string {
	return s.String()
}

func (s *CommentQueryResponseDataCommentsItem) SetPoiId(v int64) *CommentQueryResponseDataCommentsItem {
	s.PoiId = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItem) SetProductInfo(v *CommentQueryResponseDataCommentsItemProductInfo) *CommentQueryResponseDataCommentsItem {
	s.ProductInfo = v
	return s
}

func (s *CommentQueryResponseDataCommentsItem) SetCommentInfo(v *CommentQueryResponseDataCommentsItemCommentInfo) *CommentQueryResponseDataCommentsItem {
	s.CommentInfo = v
	return s
}

type CommentQueryResponseDataCommentsItemCommentInfo struct {
	RateId        *int64    `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	RateImgList   []*string `json:"rate_img_list,omitempty" xml:"rate_img_list,omitempty" type:"Repeated"`
	RateScore     *int64    `json:"rate_score,omitempty" xml:"rate_score,omitempty"`
	RateText      *string   `json:"rate_text,omitempty" xml:"rate_text,omitempty"`
	CreateTime    *int64    `json:"create_time,omitempty" xml:"create_time,omitempty"`
	DiggCnt       *int64    `json:"digg_cnt,omitempty" xml:"digg_cnt,omitempty"`
	RateExporeCnt *int64    `json:"rate_expore_cnt,omitempty" xml:"rate_expore_cnt,omitempty"`
}

func (s CommentQueryResponseDataCommentsItemCommentInfo) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponseDataCommentsItemCommentInfo) GoString() string {
	return s.String()
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetRateId(v int64) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.RateId = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetRateImgList(v []*string) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.RateImgList = v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetRateScore(v int64) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.RateScore = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetRateText(v string) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.RateText = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetCreateTime(v int64) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.CreateTime = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetDiggCnt(v int64) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.DiggCnt = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemCommentInfo) SetRateExporeCnt(v int64) *CommentQueryResponseDataCommentsItemCommentInfo {
	s.RateExporeCnt = &v
	return s
}

type CommentQueryResponseDataCommentsItemProductInfo struct {
	ProductType *int32  `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ProductId   *int64  `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
}

func (s CommentQueryResponseDataCommentsItemProductInfo) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponseDataCommentsItemProductInfo) GoString() string {
	return s.String()
}

func (s *CommentQueryResponseDataCommentsItemProductInfo) SetProductType(v int32) *CommentQueryResponseDataCommentsItemProductInfo {
	s.ProductType = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemProductInfo) SetProductId(v int64) *CommentQueryResponseDataCommentsItemProductInfo {
	s.ProductId = &v
	return s
}

func (s *CommentQueryResponseDataCommentsItemProductInfo) SetProductName(v string) *CommentQueryResponseDataCommentsItemProductInfo {
	s.ProductName = &v
	return s
}

type CommentQueryResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CommentQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommentQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CommentQueryResponseExtra) SetErrorCode(v int32) *CommentQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CommentQueryResponseExtra) SetLogid(v string) *CommentQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommentQueryResponseExtra) SetNow(v int64) *CommentQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *CommentQueryResponseExtra) SetSubDescription(v string) *CommentQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommentQueryResponseExtra) SetSubErrorCode(v int32) *CommentQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CommentQueryResponseExtra) SetDescription(v string) *CommentQueryResponseExtra {
	s.Description = &v
	return s
}

type CommissionRateOperateRequest struct {
	Header      map[string]*string                         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string                                    `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Content     []*CommissionRateOperateRequestContentItem `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	OpType      *int                                       `json:"op_type,omitempty" xml:"op_type,omitempty" require:"true"`
}

func (s CommissionRateOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateOperateRequest) GoString() string {
	return s.String()
}

func (s *CommissionRateOperateRequest) SetHeader(v map[string]*string) *CommissionRateOperateRequest {
	s.Header = v
	return s
}

func (s *CommissionRateOperateRequest) SetAccessToken(v string) *CommissionRateOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *CommissionRateOperateRequest) SetAccountId(v string) *CommissionRateOperateRequest {
	s.AccountId = &v
	return s
}

func (s *CommissionRateOperateRequest) SetContent(v []*CommissionRateOperateRequestContentItem) *CommissionRateOperateRequest {
	s.Content = v
	return s
}

func (s *CommissionRateOperateRequest) SetOpType(v int) *CommissionRateOperateRequest {
	s.OpType = &v
	return s
}

type CommissionRateOperateRequestContentItem struct {
	TalentIds []*string `json:"talent_ids,omitempty" xml:"talent_ids,omitempty" type:"Repeated"`
	PlanId    *string   `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s CommissionRateOperateRequestContentItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateOperateRequestContentItem) GoString() string {
	return s.String()
}

func (s *CommissionRateOperateRequestContentItem) SetTalentIds(v []*string) *CommissionRateOperateRequestContentItem {
	s.TalentIds = v
	return s
}

func (s *CommissionRateOperateRequestContentItem) SetPlanId(v string) *CommissionRateOperateRequestContentItem {
	s.PlanId = &v
	return s
}

type CommissionRateOperateResponse struct {
	Data  []*CommissionRateOperateResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Extra *CommissionRateOperateResponseExtra      `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CommissionRateOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateOperateResponse) GoString() string {
	return s.String()
}

func (s *CommissionRateOperateResponse) SetData(v []*CommissionRateOperateResponseDataItem) *CommissionRateOperateResponse {
	s.Data = v
	return s
}

func (s *CommissionRateOperateResponse) SetExtra(v *CommissionRateOperateResponseExtra) *CommissionRateOperateResponse {
	s.Extra = v
	return s
}

type CommissionRateOperateResponseDataItem struct {
	PlanId          *string   `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	Code            *int32    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	FailedTalentIds []*string `json:"failed_talent_ids,omitempty" xml:"failed_talent_ids,omitempty" type:"Repeated"`
	Id              *string   `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	Message         *string   `json:"message,omitempty" xml:"message,omitempty" require:"true"`
}

func (s CommissionRateOperateResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateOperateResponseDataItem) GoString() string {
	return s.String()
}

func (s *CommissionRateOperateResponseDataItem) SetPlanId(v string) *CommissionRateOperateResponseDataItem {
	s.PlanId = &v
	return s
}

func (s *CommissionRateOperateResponseDataItem) SetCode(v int32) *CommissionRateOperateResponseDataItem {
	s.Code = &v
	return s
}

func (s *CommissionRateOperateResponseDataItem) SetFailedTalentIds(v []*string) *CommissionRateOperateResponseDataItem {
	s.FailedTalentIds = v
	return s
}

func (s *CommissionRateOperateResponseDataItem) SetId(v string) *CommissionRateOperateResponseDataItem {
	s.Id = &v
	return s
}

func (s *CommissionRateOperateResponseDataItem) SetMessage(v string) *CommissionRateOperateResponseDataItem {
	s.Message = &v
	return s
}

type CommissionRateOperateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CommissionRateOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRateOperateResponseExtra) SetDescription(v string) *CommissionRateOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRateOperateResponseExtra) SetErrorCode(v int32) *CommissionRateOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CommissionRateOperateResponseExtra) SetLogid(v string) *CommissionRateOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommissionRateOperateResponseExtra) SetNow(v int64) *CommissionRateOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRateOperateResponseExtra) SetSubDescription(v string) *CommissionRateOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRateOperateResponseExtra) SetSubErrorCode(v int32) *CommissionRateOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CommissionRatePlanIdQueryRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PlanType    *int               `json:"plan_type,omitempty" xml:"plan_type,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PlanIds     []*string          `json:"plan_ids,omitempty" xml:"plan_ids,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CommissionRatePlanIdQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryRequest) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryRequest) SetAccessToken(v string) *CommissionRatePlanIdQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *CommissionRatePlanIdQueryRequest) SetPlanType(v int) *CommissionRatePlanIdQueryRequest {
	s.PlanType = &v
	return s
}

func (s *CommissionRatePlanIdQueryRequest) SetAccountId(v string) *CommissionRatePlanIdQueryRequest {
	s.AccountId = &v
	return s
}

func (s *CommissionRatePlanIdQueryRequest) SetPlanIds(v []*string) *CommissionRatePlanIdQueryRequest {
	s.PlanIds = v
	return s
}

func (s *CommissionRatePlanIdQueryRequest) SetHeader(v map[string]*string) *CommissionRatePlanIdQueryRequest {
	s.Header = v
	return s
}

type CommissionRatePlanIdQueryResponse struct {
	Extra *CommissionRatePlanIdQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CommissionRatePlanIdQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CommissionRatePlanIdQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryResponse) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryResponse) SetExtra(v *CommissionRatePlanIdQueryResponseExtra) *CommissionRatePlanIdQueryResponse {
	s.Extra = v
	return s
}

func (s *CommissionRatePlanIdQueryResponse) SetData(v *CommissionRatePlanIdQueryResponseData) *CommissionRatePlanIdQueryResponse {
	s.Data = v
	return s
}

type CommissionRatePlanIdQueryResponseData struct {
	Pagination               *CommissionRatePlanIdQueryResponseDataPagination     `json:"pagination,omitempty" xml:"pagination,omitempty"`
	PlanList                 []*CommissionRatePlanIdQueryResponseDataPlanListItem `json:"plan_list,omitempty" xml:"plan_list,omitempty" type:"Repeated"`
	ShortVideoCommissionRate *int32                                               `json:"short_video_commission_rate,omitempty" xml:"short_video_commission_rate,omitempty"`
	GwErrorCode              *int32                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription            *string                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	LiveCommissionRate       *int32                                               `json:"live_commission_rate,omitempty" xml:"live_commission_rate,omitempty"`
}

func (s CommissionRatePlanIdQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryResponseData) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryResponseData) SetPagination(v *CommissionRatePlanIdQueryResponseDataPagination) *CommissionRatePlanIdQueryResponseData {
	s.Pagination = v
	return s
}

func (s *CommissionRatePlanIdQueryResponseData) SetPlanList(v []*CommissionRatePlanIdQueryResponseDataPlanListItem) *CommissionRatePlanIdQueryResponseData {
	s.PlanList = v
	return s
}

func (s *CommissionRatePlanIdQueryResponseData) SetShortVideoCommissionRate(v int32) *CommissionRatePlanIdQueryResponseData {
	s.ShortVideoCommissionRate = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseData) SetGwErrorCode(v int32) *CommissionRatePlanIdQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseData) SetGwDescription(v string) *CommissionRatePlanIdQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseData) SetLiveCommissionRate(v int32) *CommissionRatePlanIdQueryResponseData {
	s.LiveCommissionRate = &v
	return s
}

type CommissionRatePlanIdQueryResponseDataPagination struct {
	PageIndex  *int32 `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize   *int32 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	HasMore    *bool  `json:"has_more,omitempty" xml:"has_more,omitempty"`
	PageCount  *int32 `json:"page_count,omitempty" xml:"page_count,omitempty"`
}

func (s CommissionRatePlanIdQueryResponseDataPagination) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryResponseDataPagination) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryResponseDataPagination) SetPageIndex(v int32) *CommissionRatePlanIdQueryResponseDataPagination {
	s.PageIndex = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPagination) SetPageSize(v int32) *CommissionRatePlanIdQueryResponseDataPagination {
	s.PageSize = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPagination) SetTotalCount(v int32) *CommissionRatePlanIdQueryResponseDataPagination {
	s.TotalCount = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPagination) SetHasMore(v bool) *CommissionRatePlanIdQueryResponseDataPagination {
	s.HasMore = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPagination) SetPageCount(v int32) *CommissionRatePlanIdQueryResponseDataPagination {
	s.PageCount = &v
	return s
}

type CommissionRatePlanIdQueryResponseDataPlanListItem struct {
	ProductIds         []*string       `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	PlanContactPhone   *string         `json:"plan_contact_phone,omitempty" xml:"plan_contact_phone,omitempty"`
	UpdateTime         *int64          `json:"update_time,omitempty" xml:"update_time,omitempty"`
	PlanId             *string         `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	DimensionType      *int            `json:"dimension_type,omitempty" xml:"dimension_type,omitempty"`
	PlanType           *int            `json:"plan_type,omitempty" xml:"plan_type,omitempty"`
	TalentFulfill      map[string]*int `json:"talent_fulfill,omitempty" xml:"talent_fulfill,omitempty"`
	PlanName           *string         `json:"plan_name,omitempty" xml:"plan_name,omitempty"`
	CommissionDuration *int64          `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	EndTime            *int64          `json:"end_time,omitempty" xml:"end_time,omitempty"`
	CommissionRate     *int32          `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	TalentIds          []*string       `json:"talent_ids,omitempty" xml:"talent_ids,omitempty" type:"Repeated"`
	StartTime          *int64          `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Status             *int            `json:"status,omitempty" xml:"status,omitempty"`
	CreateTime         *int64          `json:"create_time,omitempty" xml:"create_time,omitempty"`
	SceneType          *int            `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	DimensionId        *string         `json:"dimension_id,omitempty" xml:"dimension_id,omitempty"`
}

func (s CommissionRatePlanIdQueryResponseDataPlanListItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryResponseDataPlanListItem) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetProductIds(v []*string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.ProductIds = v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetPlanContactPhone(v string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.PlanContactPhone = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetUpdateTime(v int64) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.UpdateTime = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetPlanId(v string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.PlanId = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetDimensionType(v int) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.DimensionType = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetPlanType(v int) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.PlanType = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetTalentFulfill(v map[string]*int) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.TalentFulfill = v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetPlanName(v string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.PlanName = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetCommissionDuration(v int64) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.CommissionDuration = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetEndTime(v int64) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.EndTime = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetCommissionRate(v int32) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.CommissionRate = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetTalentIds(v []*string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.TalentIds = v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetStartTime(v int64) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.StartTime = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetStatus(v int) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.Status = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetCreateTime(v int64) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.CreateTime = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetSceneType(v int) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.SceneType = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseDataPlanListItem) SetDimensionId(v string) *CommissionRatePlanIdQueryResponseDataPlanListItem {
	s.DimensionId = &v
	return s
}

type CommissionRatePlanIdQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CommissionRatePlanIdQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRatePlanIdQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetLogid(v string) *CommissionRatePlanIdQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetNow(v int64) *CommissionRatePlanIdQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetSubDescription(v string) *CommissionRatePlanIdQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetSubErrorCode(v int32) *CommissionRatePlanIdQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetDescription(v string) *CommissionRatePlanIdQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRatePlanIdQueryResponseExtra) SetErrorCode(v int32) *CommissionRatePlanIdQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type CommissionRateQueryRequest struct {
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DimensionId   *string            `json:"dimension_id,omitempty" xml:"dimension_id,omitempty" require:"true"`
	DimensionType *int               `json:"dimension_type,omitempty" xml:"dimension_type,omitempty" require:"true"`
	PageIndex     *int32             `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize      *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s CommissionRateQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryRequest) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryRequest) SetAccountId(v string) *CommissionRateQueryRequest {
	s.AccountId = &v
	return s
}

func (s *CommissionRateQueryRequest) SetHeader(v map[string]*string) *CommissionRateQueryRequest {
	s.Header = v
	return s
}

func (s *CommissionRateQueryRequest) SetAccessToken(v string) *CommissionRateQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *CommissionRateQueryRequest) SetDimensionId(v string) *CommissionRateQueryRequest {
	s.DimensionId = &v
	return s
}

func (s *CommissionRateQueryRequest) SetDimensionType(v int) *CommissionRateQueryRequest {
	s.DimensionType = &v
	return s
}

func (s *CommissionRateQueryRequest) SetPageIndex(v int32) *CommissionRateQueryRequest {
	s.PageIndex = &v
	return s
}

func (s *CommissionRateQueryRequest) SetPageSize(v int32) *CommissionRateQueryRequest {
	s.PageSize = &v
	return s
}

type CommissionRateQueryResponse struct {
	Data  *CommissionRateQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CommissionRateQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CommissionRateQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryResponse) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryResponse) SetData(v *CommissionRateQueryResponseData) *CommissionRateQueryResponse {
	s.Data = v
	return s
}

func (s *CommissionRateQueryResponse) SetExtra(v *CommissionRateQueryResponseExtra) *CommissionRateQueryResponse {
	s.Extra = v
	return s
}

type CommissionRateQueryResponseData struct {
	PlanList                 []*CommissionRateQueryResponseDataPlanListItem `json:"plan_list,omitempty" xml:"plan_list,omitempty" type:"Repeated"`
	GwErrorCode              *int32                                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription            *string                                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ShortVideoCommissionRate *int32                                         `json:"short_video_commission_rate,omitempty" xml:"short_video_commission_rate,omitempty"`
	LiveCommissionRate       *int32                                         `json:"live_commission_rate,omitempty" xml:"live_commission_rate,omitempty"`
	Pagination               *CommissionRateQueryResponseDataPagination     `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

func (s CommissionRateQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryResponseData) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryResponseData) SetPlanList(v []*CommissionRateQueryResponseDataPlanListItem) *CommissionRateQueryResponseData {
	s.PlanList = v
	return s
}

func (s *CommissionRateQueryResponseData) SetGwErrorCode(v int32) *CommissionRateQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CommissionRateQueryResponseData) SetGwDescription(v string) *CommissionRateQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *CommissionRateQueryResponseData) SetShortVideoCommissionRate(v int32) *CommissionRateQueryResponseData {
	s.ShortVideoCommissionRate = &v
	return s
}

func (s *CommissionRateQueryResponseData) SetLiveCommissionRate(v int32) *CommissionRateQueryResponseData {
	s.LiveCommissionRate = &v
	return s
}

func (s *CommissionRateQueryResponseData) SetPagination(v *CommissionRateQueryResponseDataPagination) *CommissionRateQueryResponseData {
	s.Pagination = v
	return s
}

type CommissionRateQueryResponseDataPagination struct {
	PageIndex  *int32 `json:"page_index,omitempty" xml:"page_index,omitempty" require:"true"`
	PageSize   *int32 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	HasMore    *bool  `json:"has_more,omitempty" xml:"has_more,omitempty"`
	PageCount  *int32 `json:"page_count,omitempty" xml:"page_count,omitempty"`
}

func (s CommissionRateQueryResponseDataPagination) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryResponseDataPagination) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryResponseDataPagination) SetPageIndex(v int32) *CommissionRateQueryResponseDataPagination {
	s.PageIndex = &v
	return s
}

func (s *CommissionRateQueryResponseDataPagination) SetPageSize(v int32) *CommissionRateQueryResponseDataPagination {
	s.PageSize = &v
	return s
}

func (s *CommissionRateQueryResponseDataPagination) SetTotalCount(v int32) *CommissionRateQueryResponseDataPagination {
	s.TotalCount = &v
	return s
}

func (s *CommissionRateQueryResponseDataPagination) SetHasMore(v bool) *CommissionRateQueryResponseDataPagination {
	s.HasMore = &v
	return s
}

func (s *CommissionRateQueryResponseDataPagination) SetPageCount(v int32) *CommissionRateQueryResponseDataPagination {
	s.PageCount = &v
	return s
}

type CommissionRateQueryResponseDataPlanListItem struct {
	PlanContactPhone   *string         `json:"plan_contact_phone,omitempty" xml:"plan_contact_phone,omitempty"`
	CommissionDuration *int64          `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	SceneType          *int            `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	TalentIds          []*string       `json:"talent_ids,omitempty" xml:"talent_ids,omitempty" type:"Repeated"`
	PlanId             *string         `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	Status             *int            `json:"status,omitempty" xml:"status,omitempty"`
	ProductIds         []*string       `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	PlanType           *int            `json:"plan_type,omitempty" xml:"plan_type,omitempty"`
	CommissionRate     *int32          `json:"commission_rate,omitempty" xml:"commission_rate,omitempty" require:"true"`
	UpdateTime         *int64          `json:"update_time,omitempty" xml:"update_time,omitempty"`
	TalentFulfill      map[string]*int `json:"talent_fulfill,omitempty" xml:"talent_fulfill,omitempty"`
	CreateTime         *int64          `json:"create_time,omitempty" xml:"create_time,omitempty"`
	PlanName           *string         `json:"plan_name,omitempty" xml:"plan_name,omitempty"`
	DimensionType      *int            `json:"dimension_type,omitempty" xml:"dimension_type,omitempty"`
	EndTime            *int64          `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64          `json:"start_time,omitempty" xml:"start_time,omitempty"`
	DimensionId        *string         `json:"dimension_id,omitempty" xml:"dimension_id,omitempty"`
}

func (s CommissionRateQueryResponseDataPlanListItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryResponseDataPlanListItem) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetPlanContactPhone(v string) *CommissionRateQueryResponseDataPlanListItem {
	s.PlanContactPhone = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetCommissionDuration(v int64) *CommissionRateQueryResponseDataPlanListItem {
	s.CommissionDuration = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetSceneType(v int) *CommissionRateQueryResponseDataPlanListItem {
	s.SceneType = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetTalentIds(v []*string) *CommissionRateQueryResponseDataPlanListItem {
	s.TalentIds = v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetPlanId(v string) *CommissionRateQueryResponseDataPlanListItem {
	s.PlanId = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetStatus(v int) *CommissionRateQueryResponseDataPlanListItem {
	s.Status = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetProductIds(v []*string) *CommissionRateQueryResponseDataPlanListItem {
	s.ProductIds = v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetPlanType(v int) *CommissionRateQueryResponseDataPlanListItem {
	s.PlanType = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetCommissionRate(v int32) *CommissionRateQueryResponseDataPlanListItem {
	s.CommissionRate = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetUpdateTime(v int64) *CommissionRateQueryResponseDataPlanListItem {
	s.UpdateTime = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetTalentFulfill(v map[string]*int) *CommissionRateQueryResponseDataPlanListItem {
	s.TalentFulfill = v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetCreateTime(v int64) *CommissionRateQueryResponseDataPlanListItem {
	s.CreateTime = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetPlanName(v string) *CommissionRateQueryResponseDataPlanListItem {
	s.PlanName = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetDimensionType(v int) *CommissionRateQueryResponseDataPlanListItem {
	s.DimensionType = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetEndTime(v int64) *CommissionRateQueryResponseDataPlanListItem {
	s.EndTime = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetStartTime(v int64) *CommissionRateQueryResponseDataPlanListItem {
	s.StartTime = &v
	return s
}

func (s *CommissionRateQueryResponseDataPlanListItem) SetDimensionId(v string) *CommissionRateQueryResponseDataPlanListItem {
	s.DimensionId = &v
	return s
}

type CommissionRateQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s CommissionRateQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRateQueryResponseExtra) SetNow(v int64) *CommissionRateQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRateQueryResponseExtra) SetSubDescription(v string) *CommissionRateQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRateQueryResponseExtra) SetSubErrorCode(v int32) *CommissionRateQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CommissionRateQueryResponseExtra) SetDescription(v string) *CommissionRateQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRateQueryResponseExtra) SetErrorCode(v int32) *CommissionRateQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CommissionRateQueryResponseExtra) SetLogid(v string) *CommissionRateQueryResponseExtra {
	s.Logid = &v
	return s
}

type CommissionRateSaveRequest struct {
	AccountId   *string                              `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	PlanDetail  *CommissionRateSaveRequestPlanDetail `json:"plan_detail,omitempty" xml:"plan_detail,omitempty"`
	Header      map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CommissionRateSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateSaveRequest) GoString() string {
	return s.String()
}

func (s *CommissionRateSaveRequest) SetAccountId(v string) *CommissionRateSaveRequest {
	s.AccountId = &v
	return s
}

func (s *CommissionRateSaveRequest) SetPlanDetail(v *CommissionRateSaveRequestPlanDetail) *CommissionRateSaveRequest {
	s.PlanDetail = v
	return s
}

func (s *CommissionRateSaveRequest) SetHeader(v map[string]*string) *CommissionRateSaveRequest {
	s.Header = v
	return s
}

func (s *CommissionRateSaveRequest) SetAccessToken(v string) *CommissionRateSaveRequest {
	s.AccessToken = &v
	return s
}

type CommissionRateSaveRequestPlanDetail struct {
	TalentIds          []*string `json:"talent_ids,omitempty" xml:"talent_ids,omitempty" type:"Repeated"`
	DimensionType      *int      `json:"dimension_type,omitempty" xml:"dimension_type,omitempty"`
	EndTime            *int64    `json:"end_time,omitempty" xml:"end_time,omitempty"`
	PlanType           *int      `json:"plan_type,omitempty" xml:"plan_type,omitempty"`
	ProductIds         []*string `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	DimensionIds       []*string `json:"dimension_ids,omitempty" xml:"dimension_ids,omitempty" type:"Repeated"`
	CommissionDuration *int64    `json:"commission_duration,omitempty" xml:"commission_duration,omitempty"`
	CommissionRate     *int32    `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	PlanName           *string   `json:"plan_name,omitempty" xml:"plan_name,omitempty"`
	PlanIds            []*string `json:"plan_ids,omitempty" xml:"plan_ids,omitempty" type:"Repeated"`
	SceneType          *int      `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
	StartTime          *int64    `json:"start_time,omitempty" xml:"start_time,omitempty"`
	PlanContactPhone   *string   `json:"plan_contact_phone,omitempty" xml:"plan_contact_phone,omitempty"`
}

func (s CommissionRateSaveRequestPlanDetail) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateSaveRequestPlanDetail) GoString() string {
	return s.String()
}

func (s *CommissionRateSaveRequestPlanDetail) SetTalentIds(v []*string) *CommissionRateSaveRequestPlanDetail {
	s.TalentIds = v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetDimensionType(v int) *CommissionRateSaveRequestPlanDetail {
	s.DimensionType = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetEndTime(v int64) *CommissionRateSaveRequestPlanDetail {
	s.EndTime = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetPlanType(v int) *CommissionRateSaveRequestPlanDetail {
	s.PlanType = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetProductIds(v []*string) *CommissionRateSaveRequestPlanDetail {
	s.ProductIds = v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetDimensionIds(v []*string) *CommissionRateSaveRequestPlanDetail {
	s.DimensionIds = v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetCommissionDuration(v int64) *CommissionRateSaveRequestPlanDetail {
	s.CommissionDuration = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetCommissionRate(v int32) *CommissionRateSaveRequestPlanDetail {
	s.CommissionRate = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetPlanName(v string) *CommissionRateSaveRequestPlanDetail {
	s.PlanName = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetPlanIds(v []*string) *CommissionRateSaveRequestPlanDetail {
	s.PlanIds = v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetSceneType(v int) *CommissionRateSaveRequestPlanDetail {
	s.SceneType = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetStartTime(v int64) *CommissionRateSaveRequestPlanDetail {
	s.StartTime = &v
	return s
}

func (s *CommissionRateSaveRequestPlanDetail) SetPlanContactPhone(v string) *CommissionRateSaveRequestPlanDetail {
	s.PlanContactPhone = &v
	return s
}

type CommissionRateSaveResponse struct {
	Data  []*CommissionRateSaveResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Extra *CommissionRateSaveResponseExtra      `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CommissionRateSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateSaveResponse) GoString() string {
	return s.String()
}

func (s *CommissionRateSaveResponse) SetData(v []*CommissionRateSaveResponseDataItem) *CommissionRateSaveResponse {
	s.Data = v
	return s
}

func (s *CommissionRateSaveResponse) SetExtra(v *CommissionRateSaveResponseExtra) *CommissionRateSaveResponse {
	s.Extra = v
	return s
}

type CommissionRateSaveResponseDataItem struct {
	FailedTalentIds []*string `json:"failed_talent_ids,omitempty" xml:"failed_talent_ids,omitempty" type:"Repeated"`
	Id              *string   `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	Message         *string   `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	PlanId          *string   `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	Code            *int32    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CommissionRateSaveResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateSaveResponseDataItem) GoString() string {
	return s.String()
}

func (s *CommissionRateSaveResponseDataItem) SetFailedTalentIds(v []*string) *CommissionRateSaveResponseDataItem {
	s.FailedTalentIds = v
	return s
}

func (s *CommissionRateSaveResponseDataItem) SetId(v string) *CommissionRateSaveResponseDataItem {
	s.Id = &v
	return s
}

func (s *CommissionRateSaveResponseDataItem) SetMessage(v string) *CommissionRateSaveResponseDataItem {
	s.Message = &v
	return s
}

func (s *CommissionRateSaveResponseDataItem) SetPlanId(v string) *CommissionRateSaveResponseDataItem {
	s.PlanId = &v
	return s
}

func (s *CommissionRateSaveResponseDataItem) SetCode(v int32) *CommissionRateSaveResponseDataItem {
	s.Code = &v
	return s
}

type CommissionRateSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CommissionRateSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRateSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRateSaveResponseExtra) SetDescription(v string) *CommissionRateSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRateSaveResponseExtra) SetErrorCode(v int32) *CommissionRateSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CommissionRateSaveResponseExtra) SetLogid(v string) *CommissionRateSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommissionRateSaveResponseExtra) SetNow(v int64) *CommissionRateSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRateSaveResponseExtra) SetSubDescription(v string) *CommissionRateSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRateSaveResponseExtra) SetSubErrorCode(v int32) *CommissionRateSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CommissionRecordGetRequest struct {
	RecordId    *string            `json:"record_id,omitempty" xml:"record_id,omitempty" require:"true"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CommissionRecordGetRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordGetRequest) GoString() string {
	return s.String()
}

func (s *CommissionRecordGetRequest) SetRecordId(v string) *CommissionRecordGetRequest {
	s.RecordId = &v
	return s
}

func (s *CommissionRecordGetRequest) SetOrderId(v string) *CommissionRecordGetRequest {
	s.OrderId = &v
	return s
}

func (s *CommissionRecordGetRequest) SetHeader(v map[string]*string) *CommissionRecordGetRequest {
	s.Header = v
	return s
}

func (s *CommissionRecordGetRequest) SetAccessToken(v string) *CommissionRecordGetRequest {
	s.AccessToken = &v
	return s
}

type CommissionRecordGetResponse struct {
	Extra *CommissionRecordGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *CommissionRecordGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CommissionRecordGetResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordGetResponse) GoString() string {
	return s.String()
}

func (s *CommissionRecordGetResponse) SetExtra(v *CommissionRecordGetResponseExtra) *CommissionRecordGetResponse {
	s.Extra = v
	return s
}

func (s *CommissionRecordGetResponse) SetData(v *CommissionRecordGetResponseData) *CommissionRecordGetResponse {
	s.Data = v
	return s
}

type CommissionRecordGetResponseData struct {
	GwErrorCode           *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription         *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CommissionRecordItems []*CommissionRecordGetResponseDataCommissionRecordItemsItem `json:"commission_record_items,omitempty" xml:"commission_record_items,omitempty" type:"Repeated"`
}

func (s CommissionRecordGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordGetResponseData) GoString() string {
	return s.String()
}

func (s *CommissionRecordGetResponseData) SetGwErrorCode(v int32) *CommissionRecordGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CommissionRecordGetResponseData) SetGwDescription(v string) *CommissionRecordGetResponseData {
	s.GwDescription = &v
	return s
}

func (s *CommissionRecordGetResponseData) SetCommissionRecordItems(v []*CommissionRecordGetResponseDataCommissionRecordItemsItem) *CommissionRecordGetResponseData {
	s.CommissionRecordItems = v
	return s
}

type CommissionRecordGetResponseDataCommissionRecordItemsItem struct {
	RejectReason          *string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	CommissionModeBefore  *int    `json:"commission_mode_before,omitempty" xml:"commission_mode_before,omitempty"`
	ProductId             *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	CommissionRatioBefore *string `json:"commission_ratio_before,omitempty" xml:"commission_ratio_before,omitempty"`
	AuditCommissionMode   *int    `json:"audit_commission_mode,omitempty" xml:"audit_commission_mode,omitempty"`
	MerchantAckTime       *int64  `json:"merchant_ack_time,omitempty" xml:"merchant_ack_time,omitempty"`
	CommissionAuditStatus *int    `json:"commission_audit_status,omitempty" xml:"commission_audit_status,omitempty" require:"true"`
	ItemId                *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	AuditCommissionRatio  *string `json:"audit_commission_ratio,omitempty" xml:"audit_commission_ratio,omitempty" require:"true"`
}

func (s CommissionRecordGetResponseDataCommissionRecordItemsItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordGetResponseDataCommissionRecordItemsItem) GoString() string {
	return s.String()
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetRejectReason(v string) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.RejectReason = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetCommissionModeBefore(v int) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.CommissionModeBefore = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetProductId(v string) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.ProductId = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetCommissionRatioBefore(v string) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.CommissionRatioBefore = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetAuditCommissionMode(v int) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.AuditCommissionMode = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetMerchantAckTime(v int64) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.MerchantAckTime = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetCommissionAuditStatus(v int) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.CommissionAuditStatus = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetItemId(v string) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.ItemId = &v
	return s
}

func (s *CommissionRecordGetResponseDataCommissionRecordItemsItem) SetAuditCommissionRatio(v string) *CommissionRecordGetResponseDataCommissionRecordItemsItem {
	s.AuditCommissionRatio = &v
	return s
}

type CommissionRecordGetResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CommissionRecordGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordGetResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRecordGetResponseExtra) SetDescription(v string) *CommissionRecordGetResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRecordGetResponseExtra) SetErrorCode(v int32) *CommissionRecordGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CommissionRecordGetResponseExtra) SetLogid(v string) *CommissionRecordGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommissionRecordGetResponseExtra) SetNow(v int64) *CommissionRecordGetResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRecordGetResponseExtra) SetSubDescription(v string) *CommissionRecordGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRecordGetResponseExtra) SetSubErrorCode(v int32) *CommissionRecordGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CommissionRecordQueryRequest struct {
	Size        *int32             `json:"size,omitempty" xml:"size,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Page        *int32             `json:"page,omitempty" xml:"page,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CommissionRecordQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordQueryRequest) GoString() string {
	return s.String()
}

func (s *CommissionRecordQueryRequest) SetSize(v int32) *CommissionRecordQueryRequest {
	s.Size = &v
	return s
}

func (s *CommissionRecordQueryRequest) SetOrderId(v string) *CommissionRecordQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CommissionRecordQueryRequest) SetPage(v int32) *CommissionRecordQueryRequest {
	s.Page = &v
	return s
}

func (s *CommissionRecordQueryRequest) SetHeader(v map[string]*string) *CommissionRecordQueryRequest {
	s.Header = v
	return s
}

func (s *CommissionRecordQueryRequest) SetAccessToken(v string) *CommissionRecordQueryRequest {
	s.AccessToken = &v
	return s
}

type CommissionRecordQueryResponse struct {
	Extra *CommissionRecordQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *CommissionRecordQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CommissionRecordQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordQueryResponse) GoString() string {
	return s.String()
}

func (s *CommissionRecordQueryResponse) SetExtra(v *CommissionRecordQueryResponseExtra) *CommissionRecordQueryResponse {
	s.Extra = v
	return s
}

func (s *CommissionRecordQueryResponse) SetData(v *CommissionRecordQueryResponseData) *CommissionRecordQueryResponse {
	s.Data = v
	return s
}

type CommissionRecordQueryResponseData struct {
	Total             *int32                                                    `json:"total,omitempty" xml:"total,omitempty"`
	CommissionRecords []*CommissionRecordQueryResponseDataCommissionRecordsItem `json:"commission_records,omitempty" xml:"commission_records,omitempty" type:"Repeated"`
	GwErrorCode       *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription     *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CommissionRecordQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordQueryResponseData) GoString() string {
	return s.String()
}

func (s *CommissionRecordQueryResponseData) SetTotal(v int32) *CommissionRecordQueryResponseData {
	s.Total = &v
	return s
}

func (s *CommissionRecordQueryResponseData) SetCommissionRecords(v []*CommissionRecordQueryResponseDataCommissionRecordsItem) *CommissionRecordQueryResponseData {
	s.CommissionRecords = v
	return s
}

func (s *CommissionRecordQueryResponseData) SetGwErrorCode(v int32) *CommissionRecordQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CommissionRecordQueryResponseData) SetGwDescription(v string) *CommissionRecordQueryResponseData {
	s.GwDescription = &v
	return s
}

type CommissionRecordQueryResponseDataCommissionRecordsItem struct {
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	RecordId   *string `json:"record_id,omitempty" xml:"record_id,omitempty" require:"true"`
}

func (s CommissionRecordQueryResponseDataCommissionRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordQueryResponseDataCommissionRecordsItem) GoString() string {
	return s.String()
}

func (s *CommissionRecordQueryResponseDataCommissionRecordsItem) SetCreateTime(v int64) *CommissionRecordQueryResponseDataCommissionRecordsItem {
	s.CreateTime = &v
	return s
}

func (s *CommissionRecordQueryResponseDataCommissionRecordsItem) SetRecordId(v string) *CommissionRecordQueryResponseDataCommissionRecordsItem {
	s.RecordId = &v
	return s
}

type CommissionRecordQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CommissionRecordQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CommissionRecordQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CommissionRecordQueryResponseExtra) SetLogid(v string) *CommissionRecordQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *CommissionRecordQueryResponseExtra) SetNow(v int64) *CommissionRecordQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *CommissionRecordQueryResponseExtra) SetSubDescription(v string) *CommissionRecordQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CommissionRecordQueryResponseExtra) SetSubErrorCode(v int32) *CommissionRecordQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CommissionRecordQueryResponseExtra) SetDescription(v string) *CommissionRecordQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *CommissionRecordQueryResponseExtra) SetErrorCode(v int32) *CommissionRecordQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type CommonPlanSellDetailRequest struct {
	PlanIdList  []*int64           `json:"plan_id_list,omitempty" xml:"plan_id_list,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CommonPlanSellDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanSellDetailRequest) GoString() string {
	return s.String()
}

func (s *CommonPlanSellDetailRequest) SetPlanIdList(v []*int64) *CommonPlanSellDetailRequest {
	s.PlanIdList = v
	return s
}

func (s *CommonPlanSellDetailRequest) SetHeader(v map[string]*string) *CommonPlanSellDetailRequest {
	s.Header = v
	return s
}

func (s *CommonPlanSellDetailRequest) SetAccessToken(v string) *CommonPlanSellDetailRequest {
	s.AccessToken = &v
	return s
}

type CommonPlanSellDetailResponse struct {
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CommonPlanSellDetailResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CommonPlanSellDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanSellDetailResponse) GoString() string {
	return s.String()
}

func (s *CommonPlanSellDetailResponse) SetErrNo(v int32) *CommonPlanSellDetailResponse {
	s.ErrNo = &v
	return s
}

func (s *CommonPlanSellDetailResponse) SetLogId(v string) *CommonPlanSellDetailResponse {
	s.LogId = &v
	return s
}

func (s *CommonPlanSellDetailResponse) SetData(v *CommonPlanSellDetailResponseData) *CommonPlanSellDetailResponse {
	s.Data = v
	return s
}

func (s *CommonPlanSellDetailResponse) SetErrMsg(v string) *CommonPlanSellDetailResponse {
	s.ErrMsg = &v
	return s
}

type CommonPlanSellDetailResponseData struct {
	Data map[string]*CommonPlanSellDetailResponseDataDataValue `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Date *string                                               `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s CommonPlanSellDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanSellDetailResponseData) GoString() string {
	return s.String()
}

func (s *CommonPlanSellDetailResponseData) SetData(v map[string]*CommonPlanSellDetailResponseDataDataValue) *CommonPlanSellDetailResponseData {
	s.Data = v
	return s
}

func (s *CommonPlanSellDetailResponseData) SetDate(v string) *CommonPlanSellDetailResponseData {
	s.Date = &v
	return s
}

type CommonPlanSellDetailResponseDataDataValue struct {
	UsedGmv          *int64 `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	Gmv              *int64 `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
	PlanId           *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	TalentCommission *int64 `json:"talent_commission,omitempty" xml:"talent_commission,omitempty" require:"true"`
}

func (s CommonPlanSellDetailResponseDataDataValue) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanSellDetailResponseDataDataValue) GoString() string {
	return s.String()
}

func (s *CommonPlanSellDetailResponseDataDataValue) SetUsedGmv(v int64) *CommonPlanSellDetailResponseDataDataValue {
	s.UsedGmv = &v
	return s
}

func (s *CommonPlanSellDetailResponseDataDataValue) SetGmv(v int64) *CommonPlanSellDetailResponseDataDataValue {
	s.Gmv = &v
	return s
}

func (s *CommonPlanSellDetailResponseDataDataValue) SetPlanId(v int64) *CommonPlanSellDetailResponseDataDataValue {
	s.PlanId = &v
	return s
}

func (s *CommonPlanSellDetailResponseDataDataValue) SetTalentCommission(v int64) *CommonPlanSellDetailResponseDataDataValue {
	s.TalentCommission = &v
	return s
}

type CommonPlanTalentDetailRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DouyinIdList []*string          `json:"douyin_id_list,omitempty" xml:"douyin_id_list,omitempty" require:"true" type:"Repeated"`
	PlanId       *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s CommonPlanTalentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentDetailRequest) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentDetailRequest) SetHeader(v map[string]*string) *CommonPlanTalentDetailRequest {
	s.Header = v
	return s
}

func (s *CommonPlanTalentDetailRequest) SetAccessToken(v string) *CommonPlanTalentDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *CommonPlanTalentDetailRequest) SetDouyinIdList(v []*string) *CommonPlanTalentDetailRequest {
	s.DouyinIdList = v
	return s
}

func (s *CommonPlanTalentDetailRequest) SetPlanId(v int64) *CommonPlanTalentDetailRequest {
	s.PlanId = &v
	return s
}

type CommonPlanTalentDetailResponse struct {
	Data   *CommonPlanTalentDetailResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CommonPlanTalentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentDetailResponse) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentDetailResponse) SetData(v *CommonPlanTalentDetailResponseData) *CommonPlanTalentDetailResponse {
	s.Data = v
	return s
}

func (s *CommonPlanTalentDetailResponse) SetErrMsg(v string) *CommonPlanTalentDetailResponse {
	s.ErrMsg = &v
	return s
}

func (s *CommonPlanTalentDetailResponse) SetErrNo(v int32) *CommonPlanTalentDetailResponse {
	s.ErrNo = &v
	return s
}

func (s *CommonPlanTalentDetailResponse) SetLogId(v string) *CommonPlanTalentDetailResponse {
	s.LogId = &v
	return s
}

type CommonPlanTalentDetailResponseData struct {
	Date *string                                                 `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Data map[string]*CommonPlanTalentDetailResponseDataDataValue `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CommonPlanTalentDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentDetailResponseData) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentDetailResponseData) SetDate(v string) *CommonPlanTalentDetailResponseData {
	s.Date = &v
	return s
}

func (s *CommonPlanTalentDetailResponseData) SetData(v map[string]*CommonPlanTalentDetailResponseDataDataValue) *CommonPlanTalentDetailResponseData {
	s.Data = v
	return s
}

type CommonPlanTalentDetailResponseDataDataValue struct {
	LiveCnt          *int32 `json:"live_cnt,omitempty" xml:"live_cnt,omitempty" require:"true"`
	ShortVideoCnt    *int32 `json:"short_video_cnt,omitempty" xml:"short_video_cnt,omitempty" require:"true"`
	TalentCommission *int64 `json:"talent_commission,omitempty" xml:"talent_commission,omitempty" require:"true"`
	UsedGmv          *int64 `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	Gmv              *int64 `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
}

func (s CommonPlanTalentDetailResponseDataDataValue) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentDetailResponseDataDataValue) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentDetailResponseDataDataValue) SetLiveCnt(v int32) *CommonPlanTalentDetailResponseDataDataValue {
	s.LiveCnt = &v
	return s
}

func (s *CommonPlanTalentDetailResponseDataDataValue) SetShortVideoCnt(v int32) *CommonPlanTalentDetailResponseDataDataValue {
	s.ShortVideoCnt = &v
	return s
}

func (s *CommonPlanTalentDetailResponseDataDataValue) SetTalentCommission(v int64) *CommonPlanTalentDetailResponseDataDataValue {
	s.TalentCommission = &v
	return s
}

func (s *CommonPlanTalentDetailResponseDataDataValue) SetUsedGmv(v int64) *CommonPlanTalentDetailResponseDataDataValue {
	s.UsedGmv = &v
	return s
}

func (s *CommonPlanTalentDetailResponseDataDataValue) SetGmv(v int64) *CommonPlanTalentDetailResponseDataDataValue {
	s.Gmv = &v
	return s
}

type CommonPlanTalentMediaListRequest struct {
	ContentType *int32             `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DouyinId    *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	PlanId      *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s CommonPlanTalentMediaListRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentMediaListRequest) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentMediaListRequest) SetContentType(v int32) *CommonPlanTalentMediaListRequest {
	s.ContentType = &v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetHeader(v map[string]*string) *CommonPlanTalentMediaListRequest {
	s.Header = v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetAccessToken(v string) *CommonPlanTalentMediaListRequest {
	s.AccessToken = &v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetDouyinId(v string) *CommonPlanTalentMediaListRequest {
	s.DouyinId = &v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetPageNum(v int32) *CommonPlanTalentMediaListRequest {
	s.PageNum = &v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetPageSize(v int32) *CommonPlanTalentMediaListRequest {
	s.PageSize = &v
	return s
}

func (s *CommonPlanTalentMediaListRequest) SetPlanId(v int64) *CommonPlanTalentMediaListRequest {
	s.PlanId = &v
	return s
}

type CommonPlanTalentMediaListResponse struct {
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CommonPlanTalentMediaListResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CommonPlanTalentMediaListResponse) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentMediaListResponse) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentMediaListResponse) SetErrMsg(v string) *CommonPlanTalentMediaListResponse {
	s.ErrMsg = &v
	return s
}

func (s *CommonPlanTalentMediaListResponse) SetErrNo(v int32) *CommonPlanTalentMediaListResponse {
	s.ErrNo = &v
	return s
}

func (s *CommonPlanTalentMediaListResponse) SetLogId(v string) *CommonPlanTalentMediaListResponse {
	s.LogId = &v
	return s
}

func (s *CommonPlanTalentMediaListResponse) SetData(v *CommonPlanTalentMediaListResponseData) *CommonPlanTalentMediaListResponse {
	s.Data = v
	return s
}

type CommonPlanTalentMediaListResponseData struct {
	TotalCount *int64                                           `json:"total_count,omitempty" xml:"total_count,omitempty" require:"true"`
	Data       []*CommonPlanTalentMediaListResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
	Date       *string                                          `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	PageCount  *int64                                           `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
}

func (s CommonPlanTalentMediaListResponseData) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentMediaListResponseData) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentMediaListResponseData) SetTotalCount(v int64) *CommonPlanTalentMediaListResponseData {
	s.TotalCount = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseData) SetData(v []*CommonPlanTalentMediaListResponseDataDataItem) *CommonPlanTalentMediaListResponseData {
	s.Data = v
	return s
}

func (s *CommonPlanTalentMediaListResponseData) SetDate(v string) *CommonPlanTalentMediaListResponseData {
	s.Date = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseData) SetPageCount(v int64) *CommonPlanTalentMediaListResponseData {
	s.PageCount = &v
	return s
}

type CommonPlanTalentMediaListResponseDataDataItem struct {
	Gmv              *int64  `json:"gmv,omitempty" xml:"gmv,omitempty" require:"true"`
	PlayCnt          *int64  `json:"play_cnt,omitempty" xml:"play_cnt,omitempty" require:"true"`
	TalentCommission *int64  `json:"talent_commission,omitempty" xml:"talent_commission,omitempty" require:"true"`
	UsedGmv          *int64  `json:"used_gmv,omitempty" xml:"used_gmv,omitempty" require:"true"`
	ContentId        *string `json:"content_id,omitempty" xml:"content_id,omitempty" require:"true"`
	ContentOpenId    *string `json:"content_open_id,omitempty" xml:"content_open_id,omitempty" require:"true"`
	ContentType      *int32  `json:"content_type,omitempty" xml:"content_type,omitempty" require:"true"`
}

func (s CommonPlanTalentMediaListResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s CommonPlanTalentMediaListResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetGmv(v int64) *CommonPlanTalentMediaListResponseDataDataItem {
	s.Gmv = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetPlayCnt(v int64) *CommonPlanTalentMediaListResponseDataDataItem {
	s.PlayCnt = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetTalentCommission(v int64) *CommonPlanTalentMediaListResponseDataDataItem {
	s.TalentCommission = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetUsedGmv(v int64) *CommonPlanTalentMediaListResponseDataDataItem {
	s.UsedGmv = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetContentId(v string) *CommonPlanTalentMediaListResponseDataDataItem {
	s.ContentId = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetContentOpenId(v string) *CommonPlanTalentMediaListResponseDataDataItem {
	s.ContentOpenId = &v
	return s
}

func (s *CommonPlanTalentMediaListResponseDataDataItem) SetContentType(v int32) *CommonPlanTalentMediaListResponseDataDataItem {
	s.ContentType = &v
	return s
}

type CommonRequest struct {
	Host   *string            `json:"Host,omitempty" xml:"Host,omitempty" require:"true"`
	Path   *string            `json:"Path,omitempty" xml:"Path,omitempty" require:"true"`
	Method *string            `json:"Method,omitempty" xml:"Method,omitempty" require:"true"`
	Header map[string]*string `json:"Header,omitempty" xml:"Header,omitempty" require:"true"`
	Query  map[string]*string `json:"Query,omitempty" xml:"Query,omitempty" require:"true"`
	Body   interface{}        `json:"Body,omitempty" xml:"Body,omitempty" require:"true"`
}

func (s CommonRequest) String() string {
	return tea.Prettify(s)
}

func (s CommonRequest) GoString() string {
	return s.String()
}

func (s *CommonRequest) SetHost(v string) *CommonRequest {
	s.Host = &v
	return s
}

func (s *CommonRequest) SetPath(v string) *CommonRequest {
	s.Path = &v
	return s
}

func (s *CommonRequest) SetMethod(v string) *CommonRequest {
	s.Method = &v
	return s
}

func (s *CommonRequest) SetHeader(v map[string]*string) *CommonRequest {
	s.Header = v
	return s
}

func (s *CommonRequest) SetQuery(v map[string]*string) *CommonRequest {
	s.Query = v
	return s
}

func (s *CommonRequest) SetBody(v interface{}) *CommonRequest {
	s.Body = v
	return s
}

type CompleteConfirmRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CompleteConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteConfirmRequest) GoString() string {
	return s.String()
}

func (s *CompleteConfirmRequest) SetAccessToken(v string) *CompleteConfirmRequest {
	s.AccessToken = &v
	return s
}

func (s *CompleteConfirmRequest) SetOrderId(v string) *CompleteConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *CompleteConfirmRequest) SetHeader(v map[string]*string) *CompleteConfirmRequest {
	s.Header = v
	return s
}

type CompleteConfirmResponse struct {
	Extra *CompleteConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CompleteConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CompleteConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteConfirmResponse) GoString() string {
	return s.String()
}

func (s *CompleteConfirmResponse) SetExtra(v *CompleteConfirmResponseExtra) *CompleteConfirmResponse {
	s.Extra = v
	return s
}

func (s *CompleteConfirmResponse) SetData(v *CompleteConfirmResponseData) *CompleteConfirmResponse {
	s.Data = v
	return s
}

type CompleteConfirmResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CompleteConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s CompleteConfirmResponseData) GoString() string {
	return s.String()
}

func (s *CompleteConfirmResponseData) SetGwDescription(v string) *CompleteConfirmResponseData {
	s.GwDescription = &v
	return s
}

func (s *CompleteConfirmResponseData) SetGwErrorCode(v int32) *CompleteConfirmResponseData {
	s.GwErrorCode = &v
	return s
}

type CompleteConfirmResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s CompleteConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CompleteConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *CompleteConfirmResponseExtra) SetNow(v int64) *CompleteConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *CompleteConfirmResponseExtra) SetSubDescription(v string) *CompleteConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CompleteConfirmResponseExtra) SetSubErrorCode(v int32) *CompleteConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CompleteConfirmResponseExtra) SetDescription(v string) *CompleteConfirmResponseExtra {
	s.Description = &v
	return s
}

func (s *CompleteConfirmResponseExtra) SetErrorCode(v int32) *CompleteConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CompleteConfirmResponseExtra) SetLogid(v string) *CompleteConfirmResponseExtra {
	s.Logid = &v
	return s
}

type CompleteUploadUserResultRequest struct {
	AppId            *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	WorldRankVersion *string            `json:"world_rank_version,omitempty" xml:"world_rank_version,omitempty" require:"true"`
	IsOnlineVersion  *bool              `json:"is_online_version,omitempty" xml:"is_online_version,omitempty" require:"true"`
	CompleteTime     *int64             `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CompleteUploadUserResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteUploadUserResultRequest) GoString() string {
	return s.String()
}

func (s *CompleteUploadUserResultRequest) SetAppId(v string) *CompleteUploadUserResultRequest {
	s.AppId = &v
	return s
}

func (s *CompleteUploadUserResultRequest) SetWorldRankVersion(v string) *CompleteUploadUserResultRequest {
	s.WorldRankVersion = &v
	return s
}

func (s *CompleteUploadUserResultRequest) SetIsOnlineVersion(v bool) *CompleteUploadUserResultRequest {
	s.IsOnlineVersion = &v
	return s
}

func (s *CompleteUploadUserResultRequest) SetCompleteTime(v int64) *CompleteUploadUserResultRequest {
	s.CompleteTime = &v
	return s
}

func (s *CompleteUploadUserResultRequest) SetHeader(v map[string]*string) *CompleteUploadUserResultRequest {
	s.Header = v
	return s
}

func (s *CompleteUploadUserResultRequest) SetAccessToken(v string) *CompleteUploadUserResultRequest {
	s.AccessToken = &v
	return s
}

type CompleteUploadUserResultResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CompleteUploadUserResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteUploadUserResultResponse) GoString() string {
	return s.String()
}

func (s *CompleteUploadUserResultResponse) SetErrNo(v int64) *CompleteUploadUserResultResponse {
	s.ErrNo = &v
	return s
}

func (s *CompleteUploadUserResultResponse) SetErrMsg(v string) *CompleteUploadUserResultResponse {
	s.ErrMsg = &v
	return s
}

type CompleteVideoPartUploadRequest struct {
	UploadId    *string            `json:"upload_id,omitempty" xml:"upload_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CompleteVideoPartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteVideoPartUploadRequest) GoString() string {
	return s.String()
}

func (s *CompleteVideoPartUploadRequest) SetUploadId(v string) *CompleteVideoPartUploadRequest {
	s.UploadId = &v
	return s
}

func (s *CompleteVideoPartUploadRequest) SetOpenId(v string) *CompleteVideoPartUploadRequest {
	s.OpenId = &v
	return s
}

func (s *CompleteVideoPartUploadRequest) SetHeader(v map[string]*string) *CompleteVideoPartUploadRequest {
	s.Header = v
	return s
}

func (s *CompleteVideoPartUploadRequest) SetAccessToken(v string) *CompleteVideoPartUploadRequest {
	s.AccessToken = &v
	return s
}

type CompleteVideoPartUploadResponse struct {
	Extra *CompleteVideoPartUploadResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CompleteVideoPartUploadResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CompleteVideoPartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteVideoPartUploadResponse) GoString() string {
	return s.String()
}

func (s *CompleteVideoPartUploadResponse) SetExtra(v *CompleteVideoPartUploadResponseExtra) *CompleteVideoPartUploadResponse {
	s.Extra = v
	return s
}

func (s *CompleteVideoPartUploadResponse) SetData(v *CompleteVideoPartUploadResponseData) *CompleteVideoPartUploadResponse {
	s.Data = v
	return s
}

type CompleteVideoPartUploadResponseData struct {
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Video         *CompleteVideoPartUploadResponseDataVideo `json:"video,omitempty" xml:"video,omitempty" require:"true"`
}

func (s CompleteVideoPartUploadResponseData) String() string {
	return tea.Prettify(s)
}

func (s CompleteVideoPartUploadResponseData) GoString() string {
	return s.String()
}

func (s *CompleteVideoPartUploadResponseData) SetGwErrorCode(v int32) *CompleteVideoPartUploadResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CompleteVideoPartUploadResponseData) SetGwDescription(v string) *CompleteVideoPartUploadResponseData {
	s.GwDescription = &v
	return s
}

func (s *CompleteVideoPartUploadResponseData) SetVideo(v *CompleteVideoPartUploadResponseDataVideo) *CompleteVideoPartUploadResponseData {
	s.Video = v
	return s
}

type CompleteVideoPartUploadResponseDataVideo struct {
	Height  *int32  `json:"height,omitempty" xml:"height,omitempty" require:"true"`
	VideoId *string `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	Width   *int32  `json:"width,omitempty" xml:"width,omitempty" require:"true"`
}

func (s CompleteVideoPartUploadResponseDataVideo) String() string {
	return tea.Prettify(s)
}

func (s CompleteVideoPartUploadResponseDataVideo) GoString() string {
	return s.String()
}

func (s *CompleteVideoPartUploadResponseDataVideo) SetHeight(v int32) *CompleteVideoPartUploadResponseDataVideo {
	s.Height = &v
	return s
}

func (s *CompleteVideoPartUploadResponseDataVideo) SetVideoId(v string) *CompleteVideoPartUploadResponseDataVideo {
	s.VideoId = &v
	return s
}

func (s *CompleteVideoPartUploadResponseDataVideo) SetWidth(v int32) *CompleteVideoPartUploadResponseDataVideo {
	s.Width = &v
	return s
}

type CompleteVideoPartUploadResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CompleteVideoPartUploadResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CompleteVideoPartUploadResponseExtra) GoString() string {
	return s.String()
}

func (s *CompleteVideoPartUploadResponseExtra) SetSubDescription(v string) *CompleteVideoPartUploadResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CompleteVideoPartUploadResponseExtra) SetSubErrorCode(v int32) *CompleteVideoPartUploadResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CompleteVideoPartUploadResponseExtra) SetDescription(v string) *CompleteVideoPartUploadResponseExtra {
	s.Description = &v
	return s
}

func (s *CompleteVideoPartUploadResponseExtra) SetErrorCode(v int32) *CompleteVideoPartUploadResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CompleteVideoPartUploadResponseExtra) SetLogid(v string) *CompleteVideoPartUploadResponseExtra {
	s.Logid = &v
	return s
}

func (s *CompleteVideoPartUploadResponseExtra) SetNow(v int64) *CompleteVideoPartUploadResponseExtra {
	s.Now = &v
	return s
}

type ConfigLimitOpPointRequest struct {
	AppId             *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	DailyAddCount     *int64             `json:"daily_add_count,omitempty" xml:"daily_add_count,omitempty"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DailyDeductCount  *int64             `json:"daily_deduct_count,omitempty" xml:"daily_deduct_count,omitempty"`
	OpType            *int               `json:"op_type,omitempty" xml:"op_type,omitempty" require:"true"`
	SingleAddLimit    *int64             `json:"single_add_limit,omitempty" xml:"single_add_limit,omitempty"`
	SingleDeductLimit *int64             `json:"single_deduct_limit,omitempty" xml:"single_deduct_limit,omitempty"`
}

func (s ConfigLimitOpPointRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigLimitOpPointRequest) GoString() string {
	return s.String()
}

func (s *ConfigLimitOpPointRequest) SetAppId(v string) *ConfigLimitOpPointRequest {
	s.AppId = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetDailyAddCount(v int64) *ConfigLimitOpPointRequest {
	s.DailyAddCount = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetHeader(v map[string]*string) *ConfigLimitOpPointRequest {
	s.Header = v
	return s
}

func (s *ConfigLimitOpPointRequest) SetAccessToken(v string) *ConfigLimitOpPointRequest {
	s.AccessToken = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetDailyDeductCount(v int64) *ConfigLimitOpPointRequest {
	s.DailyDeductCount = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetOpType(v int) *ConfigLimitOpPointRequest {
	s.OpType = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetSingleAddLimit(v int64) *ConfigLimitOpPointRequest {
	s.SingleAddLimit = &v
	return s
}

func (s *ConfigLimitOpPointRequest) SetSingleDeductLimit(v int64) *ConfigLimitOpPointRequest {
	s.SingleDeductLimit = &v
	return s
}

type ConfigLimitOpPointResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ConfigLimitOpPointResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigLimitOpPointResponse) GoString() string {
	return s.String()
}

func (s *ConfigLimitOpPointResponse) SetErrMsg(v string) *ConfigLimitOpPointResponse {
	s.ErrMsg = &v
	return s
}

func (s *ConfigLimitOpPointResponse) SetErrNo(v int32) *ConfigLimitOpPointResponse {
	s.ErrNo = &v
	return s
}

func (s *ConfigLimitOpPointResponse) SetLogId(v string) *ConfigLimitOpPointResponse {
	s.LogId = &v
	return s
}

type ConfigRegisterMaAppRequest struct {
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId            *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CallbackUrl      *string            `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	EcomMicroappType *int               `json:"ecom_microapp_type,omitempty" xml:"ecom_microapp_type,omitempty" require:"true"`
	Extra            *string            `json:"extra,omitempty" xml:"extra,omitempty"`
	PreviewPhotoUrl  *string            `json:"preview_photo_url,omitempty" xml:"preview_photo_url,omitempty"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ConfigRegisterMaAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigRegisterMaAppRequest) GoString() string {
	return s.String()
}

func (s *ConfigRegisterMaAppRequest) SetAccessToken(v string) *ConfigRegisterMaAppRequest {
	s.AccessToken = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetAppId(v string) *ConfigRegisterMaAppRequest {
	s.AppId = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetCallbackUrl(v string) *ConfigRegisterMaAppRequest {
	s.CallbackUrl = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetEcomMicroappType(v int) *ConfigRegisterMaAppRequest {
	s.EcomMicroappType = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetExtra(v string) *ConfigRegisterMaAppRequest {
	s.Extra = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetPreviewPhotoUrl(v string) *ConfigRegisterMaAppRequest {
	s.PreviewPhotoUrl = &v
	return s
}

func (s *ConfigRegisterMaAppRequest) SetHeader(v map[string]*string) *ConfigRegisterMaAppRequest {
	s.Header = v
	return s
}

type ConfigRegisterMaAppResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ConfigRegisterMaAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigRegisterMaAppResponse) GoString() string {
	return s.String()
}

func (s *ConfigRegisterMaAppResponse) SetLogId(v string) *ConfigRegisterMaAppResponse {
	s.LogId = &v
	return s
}

func (s *ConfigRegisterMaAppResponse) SetErrMsg(v string) *ConfigRegisterMaAppResponse {
	s.ErrMsg = &v
	return s
}

func (s *ConfigRegisterMaAppResponse) SetErrNo(v int32) *ConfigRegisterMaAppResponse {
	s.ErrNo = &v
	return s
}

type ConvertVideoIdVideoIdToOpenItemIdRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VideoIds    []*string          `json:"video_ids,omitempty" xml:"video_ids,omitempty" require:"true" type:"Repeated"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	AccessKey   *string            `json:"access_key,omitempty" xml:"access_key,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s ConvertVideoIdVideoIdToOpenItemIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertVideoIdVideoIdToOpenItemIdRequest) GoString() string {
	return s.String()
}

func (s *ConvertVideoIdVideoIdToOpenItemIdRequest) SetAccessToken(v string) *ConvertVideoIdVideoIdToOpenItemIdRequest {
	s.AccessToken = &v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdRequest) SetVideoIds(v []*string) *ConvertVideoIdVideoIdToOpenItemIdRequest {
	s.VideoIds = v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdRequest) SetAppId(v string) *ConvertVideoIdVideoIdToOpenItemIdRequest {
	s.AppId = &v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdRequest) SetAccessKey(v string) *ConvertVideoIdVideoIdToOpenItemIdRequest {
	s.AccessKey = &v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdRequest) SetHeader(v map[string]*string) *ConvertVideoIdVideoIdToOpenItemIdRequest {
	s.Header = v
	return s
}

type ConvertVideoIdVideoIdToOpenItemIdResponse struct {
	ErrTips *string                                        `json:"err_tips,omitempty" xml:"err_tips,omitempty" require:"true"`
	Data    *ConvertVideoIdVideoIdToOpenItemIdResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo   *int                                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponse) GoString() string {
	return s.String()
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponse) SetErrTips(v string) *ConvertVideoIdVideoIdToOpenItemIdResponse {
	s.ErrTips = &v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponse) SetData(v *ConvertVideoIdVideoIdToOpenItemIdResponseData) *ConvertVideoIdVideoIdToOpenItemIdResponse {
	s.Data = v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponse) SetErrNo(v int) *ConvertVideoIdVideoIdToOpenItemIdResponse {
	s.ErrNo = &v
	return s
}

type ConvertVideoIdVideoIdToOpenItemIdResponseData struct {
	ConvertResult *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult `json:"convert_result,omitempty" xml:"convert_result,omitempty" require:"true"`
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponseData) GoString() string {
	return s.String()
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponseData) SetConvertResult(v *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult) *ConvertVideoIdVideoIdToOpenItemIdResponseData {
	s.ConvertResult = v
	return s
}

type ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult struct {
	Key *string `json:"_key,omitempty" xml:"_key,omitempty" require:"true"`
	Val *string `json:"_val,omitempty" xml:"_val,omitempty" require:"true"`
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult) String() string {
	return tea.Prettify(s)
}

func (s ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult) GoString() string {
	return s.String()
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult) SetKey(v string) *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult {
	s.Key = &v
	return s
}

func (s *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult) SetVal(v string) *ConvertVideoIdVideoIdToOpenItemIdResponseDataConvertResult {
	s.Val = &v
	return s
}

type CospaBrainCavityRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CospaBrainCavityRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityRequest) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityRequest) SetHeader(v map[string]*string) *CospaBrainCavityRequest {
	s.Header = v
	return s
}

func (s *CospaBrainCavityRequest) SetAccessToken(v string) *CospaBrainCavityRequest {
	s.AccessToken = &v
	return s
}

type CospaBrainCavityResponse struct {
	Extra *CospaBrainCavityResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CospaBrainCavityResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CospaBrainCavityResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityResponse) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityResponse) SetExtra(v *CospaBrainCavityResponseExtra) *CospaBrainCavityResponse {
	s.Extra = v
	return s
}

func (s *CospaBrainCavityResponse) SetData(v *CospaBrainCavityResponseData) *CospaBrainCavityResponse {
	s.Data = v
	return s
}

type CospaBrainCavityResponseData struct {
	List          []*CospaBrainCavityResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaBrainCavityResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityResponseData) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityResponseData) SetList(v []*CospaBrainCavityResponseDataListItem) *CospaBrainCavityResponseData {
	s.List = v
	return s
}

func (s *CospaBrainCavityResponseData) SetGwErrorCode(v int32) *CospaBrainCavityResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaBrainCavityResponseData) SetGwDescription(v string) *CospaBrainCavityResponseData {
	s.GwDescription = &v
	return s
}

type CospaBrainCavityResponseDataListItem struct {
	Rank             *int32                                               `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                              `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                              `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                              `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                               `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                               `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                             `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaBrainCavityResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
}

func (s CospaBrainCavityResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityResponseDataListItem) SetRank(v int32) *CospaBrainCavityResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetRankChange(v string) *CospaBrainCavityResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetNickname(v string) *CospaBrainCavityResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetAvatar(v string) *CospaBrainCavityResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetFollowerCount(v int64) *CospaBrainCavityResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaBrainCavityResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetEffectValue(v float64) *CospaBrainCavityResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItem) SetVideoList(v []*CospaBrainCavityResponseDataListItemVideoListItem) *CospaBrainCavityResponseDataListItem {
	s.VideoList = v
	return s
}

type CospaBrainCavityResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s CospaBrainCavityResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityResponseDataListItemVideoListItem) SetTitle(v string) *CospaBrainCavityResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItemVideoListItem) SetItemCover(v string) *CospaBrainCavityResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CospaBrainCavityResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaBrainCavityResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type CospaBrainCavityResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaBrainCavityResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaBrainCavityResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaBrainCavityResponseExtra) SetSubErrorCode(v int32) *CospaBrainCavityResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CospaBrainCavityResponseExtra) SetSubDescription(v string) *CospaBrainCavityResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CospaBrainCavityResponseExtra) SetLogid(v string) *CospaBrainCavityResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaBrainCavityResponseExtra) SetNow(v int64) *CospaBrainCavityResponseExtra {
	s.Now = &v
	return s
}

func (s *CospaBrainCavityResponseExtra) SetErrorCode(v int32) *CospaBrainCavityResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaBrainCavityResponseExtra) SetDescription(v string) *CospaBrainCavityResponseExtra {
	s.Description = &v
	return s
}

type CospaNewRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CospaNewRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaNewRequest) GoString() string {
	return s.String()
}

func (s *CospaNewRequest) SetAccessToken(v string) *CospaNewRequest {
	s.AccessToken = &v
	return s
}

func (s *CospaNewRequest) SetHeader(v map[string]*string) *CospaNewRequest {
	s.Header = v
	return s
}

type CospaNewResponse struct {
	Data  *CospaNewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CospaNewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CospaNewResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaNewResponse) GoString() string {
	return s.String()
}

func (s *CospaNewResponse) SetData(v *CospaNewResponseData) *CospaNewResponse {
	s.Data = v
	return s
}

func (s *CospaNewResponse) SetExtra(v *CospaNewResponseExtra) *CospaNewResponse {
	s.Extra = v
	return s
}

type CospaNewResponseData struct {
	List          []*CospaNewResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                          `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                         `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaNewResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaNewResponseData) GoString() string {
	return s.String()
}

func (s *CospaNewResponseData) SetList(v []*CospaNewResponseDataListItem) *CospaNewResponseData {
	s.List = v
	return s
}

func (s *CospaNewResponseData) SetGwErrorCode(v int32) *CospaNewResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaNewResponseData) SetGwDescription(v string) *CospaNewResponseData {
	s.GwDescription = &v
	return s
}

type CospaNewResponseDataListItem struct {
	RankChange       *string                                      `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                      `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                      `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                       `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                       `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                     `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaNewResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                       `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s CospaNewResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaNewResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaNewResponseDataListItem) SetRankChange(v string) *CospaNewResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetNickname(v string) *CospaNewResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetAvatar(v string) *CospaNewResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetFollowerCount(v int64) *CospaNewResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaNewResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetEffectValue(v float64) *CospaNewResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaNewResponseDataListItem) SetVideoList(v []*CospaNewResponseDataListItemVideoListItem) *CospaNewResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CospaNewResponseDataListItem) SetRank(v int32) *CospaNewResponseDataListItem {
	s.Rank = &v
	return s
}

type CospaNewResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s CospaNewResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaNewResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaNewResponseDataListItemVideoListItem) SetTitle(v string) *CospaNewResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CospaNewResponseDataListItemVideoListItem) SetItemCover(v string) *CospaNewResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CospaNewResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaNewResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type CospaNewResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CospaNewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaNewResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaNewResponseExtra) SetSubDescription(v string) *CospaNewResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CospaNewResponseExtra) SetLogid(v string) *CospaNewResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaNewResponseExtra) SetNow(v int64) *CospaNewResponseExtra {
	s.Now = &v
	return s
}

func (s *CospaNewResponseExtra) SetErrorCode(v int32) *CospaNewResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaNewResponseExtra) SetDescription(v string) *CospaNewResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaNewResponseExtra) SetSubErrorCode(v int32) *CospaNewResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CospaOutShotRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CospaOutShotRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotRequest) GoString() string {
	return s.String()
}

func (s *CospaOutShotRequest) SetHeader(v map[string]*string) *CospaOutShotRequest {
	s.Header = v
	return s
}

func (s *CospaOutShotRequest) SetAccessToken(v string) *CospaOutShotRequest {
	s.AccessToken = &v
	return s
}

type CospaOutShotResponse struct {
	Data  *CospaOutShotResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CospaOutShotResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CospaOutShotResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotResponse) GoString() string {
	return s.String()
}

func (s *CospaOutShotResponse) SetData(v *CospaOutShotResponseData) *CospaOutShotResponse {
	s.Data = v
	return s
}

func (s *CospaOutShotResponse) SetExtra(v *CospaOutShotResponseExtra) *CospaOutShotResponse {
	s.Extra = v
	return s
}

type CospaOutShotResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*CospaOutShotResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s CospaOutShotResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotResponseData) GoString() string {
	return s.String()
}

func (s *CospaOutShotResponseData) SetGwErrorCode(v int32) *CospaOutShotResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaOutShotResponseData) SetGwDescription(v string) *CospaOutShotResponseData {
	s.GwDescription = &v
	return s
}

func (s *CospaOutShotResponseData) SetList(v []*CospaOutShotResponseDataListItem) *CospaOutShotResponseData {
	s.List = v
	return s
}

type CospaOutShotResponseDataListItem struct {
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaOutShotResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
}

func (s CospaOutShotResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaOutShotResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaOutShotResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetEffectValue(v float64) *CospaOutShotResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetVideoList(v []*CospaOutShotResponseDataListItemVideoListItem) *CospaOutShotResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetRank(v int32) *CospaOutShotResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetRankChange(v string) *CospaOutShotResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetNickname(v string) *CospaOutShotResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetAvatar(v string) *CospaOutShotResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaOutShotResponseDataListItem) SetFollowerCount(v int64) *CospaOutShotResponseDataListItem {
	s.FollowerCount = &v
	return s
}

type CospaOutShotResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s CospaOutShotResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaOutShotResponseDataListItemVideoListItem) SetTitle(v string) *CospaOutShotResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CospaOutShotResponseDataListItemVideoListItem) SetItemCover(v string) *CospaOutShotResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CospaOutShotResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaOutShotResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type CospaOutShotResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CospaOutShotResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaOutShotResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaOutShotResponseExtra) SetSubDescription(v string) *CospaOutShotResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CospaOutShotResponseExtra) SetLogid(v string) *CospaOutShotResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaOutShotResponseExtra) SetNow(v int64) *CospaOutShotResponseExtra {
	s.Now = &v
	return s
}

func (s *CospaOutShotResponseExtra) SetErrorCode(v int32) *CospaOutShotResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaOutShotResponseExtra) SetDescription(v string) *CospaOutShotResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaOutShotResponseExtra) SetSubErrorCode(v int32) *CospaOutShotResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CospaOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CospaOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallRequest) GoString() string {
	return s.String()
}

func (s *CospaOverallRequest) SetHeader(v map[string]*string) *CospaOverallRequest {
	s.Header = v
	return s
}

func (s *CospaOverallRequest) SetAccessToken(v string) *CospaOverallRequest {
	s.AccessToken = &v
	return s
}

type CospaOverallResponse struct {
	Extra *CospaOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CospaOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CospaOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallResponse) GoString() string {
	return s.String()
}

func (s *CospaOverallResponse) SetExtra(v *CospaOverallResponseExtra) *CospaOverallResponse {
	s.Extra = v
	return s
}

func (s *CospaOverallResponse) SetData(v *CospaOverallResponseData) *CospaOverallResponse {
	s.Data = v
	return s
}

type CospaOverallResponseData struct {
	List          []*CospaOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallResponseData) GoString() string {
	return s.String()
}

func (s *CospaOverallResponseData) SetList(v []*CospaOverallResponseDataListItem) *CospaOverallResponseData {
	s.List = v
	return s
}

func (s *CospaOverallResponseData) SetGwErrorCode(v int32) *CospaOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaOverallResponseData) SetGwDescription(v string) *CospaOverallResponseData {
	s.GwDescription = &v
	return s
}

type CospaOverallResponseDataListItem struct {
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
}

func (s CospaOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaOverallResponseDataListItem) SetRank(v int32) *CospaOverallResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetRankChange(v string) *CospaOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetNickname(v string) *CospaOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetAvatar(v string) *CospaOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetFollowerCount(v int64) *CospaOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetEffectValue(v float64) *CospaOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaOverallResponseDataListItem) SetVideoList(v []*CospaOverallResponseDataListItemVideoListItem) *CospaOverallResponseDataListItem {
	s.VideoList = v
	return s
}

type CospaOverallResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CospaOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaOverallResponseDataListItemVideoListItem) SetItemCover(v string) *CospaOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CospaOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CospaOverallResponseDataListItemVideoListItem) SetTitle(v string) *CospaOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type CospaOverallResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CospaOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaOverallResponseExtra) SetErrorCode(v int32) *CospaOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaOverallResponseExtra) SetDescription(v string) *CospaOverallResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaOverallResponseExtra) SetSubErrorCode(v int32) *CospaOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CospaOverallResponseExtra) SetSubDescription(v string) *CospaOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CospaOverallResponseExtra) SetLogid(v string) *CospaOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaOverallResponseExtra) SetNow(v int64) *CospaOverallResponseExtra {
	s.Now = &v
	return s
}

type CospaPaintingRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CospaPaintingRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingRequest) GoString() string {
	return s.String()
}

func (s *CospaPaintingRequest) SetHeader(v map[string]*string) *CospaPaintingRequest {
	s.Header = v
	return s
}

func (s *CospaPaintingRequest) SetAccessToken(v string) *CospaPaintingRequest {
	s.AccessToken = &v
	return s
}

type CospaPaintingResponse struct {
	Extra *CospaPaintingResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CospaPaintingResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CospaPaintingResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingResponse) GoString() string {
	return s.String()
}

func (s *CospaPaintingResponse) SetExtra(v *CospaPaintingResponseExtra) *CospaPaintingResponse {
	s.Extra = v
	return s
}

func (s *CospaPaintingResponse) SetData(v *CospaPaintingResponseData) *CospaPaintingResponse {
	s.Data = v
	return s
}

type CospaPaintingResponseData struct {
	List          []*CospaPaintingResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaPaintingResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingResponseData) GoString() string {
	return s.String()
}

func (s *CospaPaintingResponseData) SetList(v []*CospaPaintingResponseDataListItem) *CospaPaintingResponseData {
	s.List = v
	return s
}

func (s *CospaPaintingResponseData) SetGwErrorCode(v int32) *CospaPaintingResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaPaintingResponseData) SetGwDescription(v string) *CospaPaintingResponseData {
	s.GwDescription = &v
	return s
}

type CospaPaintingResponseDataListItem struct {
	FollowerCount    *int64                                            `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                            `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                          `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaPaintingResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                            `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                           `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                           `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                           `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s CospaPaintingResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaPaintingResponseDataListItem) SetFollowerCount(v int64) *CospaPaintingResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaPaintingResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetEffectValue(v float64) *CospaPaintingResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetVideoList(v []*CospaPaintingResponseDataListItemVideoListItem) *CospaPaintingResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetRank(v int32) *CospaPaintingResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetRankChange(v string) *CospaPaintingResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetNickname(v string) *CospaPaintingResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaPaintingResponseDataListItem) SetAvatar(v string) *CospaPaintingResponseDataListItem {
	s.Avatar = &v
	return s
}

type CospaPaintingResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CospaPaintingResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaPaintingResponseDataListItemVideoListItem) SetItemCover(v string) *CospaPaintingResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CospaPaintingResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaPaintingResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CospaPaintingResponseDataListItemVideoListItem) SetTitle(v string) *CospaPaintingResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type CospaPaintingResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s CospaPaintingResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaPaintingResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaPaintingResponseExtra) SetLogid(v string) *CospaPaintingResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaPaintingResponseExtra) SetNow(v int64) *CospaPaintingResponseExtra {
	s.Now = &v
	return s
}

func (s *CospaPaintingResponseExtra) SetErrorCode(v int32) *CospaPaintingResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaPaintingResponseExtra) SetDescription(v string) *CospaPaintingResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaPaintingResponseExtra) SetSubErrorCode(v int32) *CospaPaintingResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CospaPaintingResponseExtra) SetSubDescription(v string) *CospaPaintingResponseExtra {
	s.SubDescription = &v
	return s
}

type CospaQingManRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CospaQingManRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManRequest) GoString() string {
	return s.String()
}

func (s *CospaQingManRequest) SetHeader(v map[string]*string) *CospaQingManRequest {
	s.Header = v
	return s
}

func (s *CospaQingManRequest) SetAccessToken(v string) *CospaQingManRequest {
	s.AccessToken = &v
	return s
}

type CospaQingManResponse struct {
	Extra *CospaQingManResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CospaQingManResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CospaQingManResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManResponse) GoString() string {
	return s.String()
}

func (s *CospaQingManResponse) SetExtra(v *CospaQingManResponseExtra) *CospaQingManResponse {
	s.Extra = v
	return s
}

func (s *CospaQingManResponse) SetData(v *CospaQingManResponseData) *CospaQingManResponse {
	s.Data = v
	return s
}

type CospaQingManResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*CospaQingManResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s CospaQingManResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManResponseData) GoString() string {
	return s.String()
}

func (s *CospaQingManResponseData) SetGwErrorCode(v int32) *CospaQingManResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaQingManResponseData) SetGwDescription(v string) *CospaQingManResponseData {
	s.GwDescription = &v
	return s
}

func (s *CospaQingManResponseData) SetList(v []*CospaQingManResponseDataListItem) *CospaQingManResponseData {
	s.List = v
	return s
}

type CospaQingManResponseDataListItem struct {
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaQingManResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
}

func (s CospaQingManResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaQingManResponseDataListItem) SetNickname(v string) *CospaQingManResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetAvatar(v string) *CospaQingManResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetFollowerCount(v int64) *CospaQingManResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaQingManResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetEffectValue(v float64) *CospaQingManResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetVideoList(v []*CospaQingManResponseDataListItemVideoListItem) *CospaQingManResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CospaQingManResponseDataListItem) SetRank(v int32) *CospaQingManResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaQingManResponseDataListItem) SetRankChange(v string) *CospaQingManResponseDataListItem {
	s.RankChange = &v
	return s
}

type CospaQingManResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s CospaQingManResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaQingManResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaQingManResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CospaQingManResponseDataListItemVideoListItem) SetTitle(v string) *CospaQingManResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CospaQingManResponseDataListItemVideoListItem) SetItemCover(v string) *CospaQingManResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type CospaQingManResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s CospaQingManResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaQingManResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaQingManResponseExtra) SetLogid(v string) *CospaQingManResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaQingManResponseExtra) SetNow(v int64) *CospaQingManResponseExtra {
	s.Now = &v
	return s
}

func (s *CospaQingManResponseExtra) SetErrorCode(v int32) *CospaQingManResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaQingManResponseExtra) SetDescription(v string) *CospaQingManResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaQingManResponseExtra) SetSubErrorCode(v int32) *CospaQingManResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CospaQingManResponseExtra) SetSubDescription(v string) *CospaQingManResponseExtra {
	s.SubDescription = &v
	return s
}

type CospaVoiceControlRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CospaVoiceControlRequest) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlRequest) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlRequest) SetAccessToken(v string) *CospaVoiceControlRequest {
	s.AccessToken = &v
	return s
}

func (s *CospaVoiceControlRequest) SetHeader(v map[string]*string) *CospaVoiceControlRequest {
	s.Header = v
	return s
}

type CospaVoiceControlResponse struct {
	Extra *CospaVoiceControlResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CospaVoiceControlResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CospaVoiceControlResponse) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlResponse) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlResponse) SetExtra(v *CospaVoiceControlResponseExtra) *CospaVoiceControlResponse {
	s.Extra = v
	return s
}

func (s *CospaVoiceControlResponse) SetData(v *CospaVoiceControlResponseData) *CospaVoiceControlResponse {
	s.Data = v
	return s
}

type CospaVoiceControlResponseData struct {
	List          []*CospaVoiceControlResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CospaVoiceControlResponseData) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlResponseData) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlResponseData) SetList(v []*CospaVoiceControlResponseDataListItem) *CospaVoiceControlResponseData {
	s.List = v
	return s
}

func (s *CospaVoiceControlResponseData) SetGwErrorCode(v int32) *CospaVoiceControlResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CospaVoiceControlResponseData) SetGwDescription(v string) *CospaVoiceControlResponseData {
	s.GwDescription = &v
	return s
}

type CospaVoiceControlResponseDataListItem struct {
	Rank             *int32                                                `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                               `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                               `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                               `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                                `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                                `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                              `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CospaVoiceControlResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
}

func (s CospaVoiceControlResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlResponseDataListItem) SetRank(v int32) *CospaVoiceControlResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetRankChange(v string) *CospaVoiceControlResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetNickname(v string) *CospaVoiceControlResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetAvatar(v string) *CospaVoiceControlResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetFollowerCount(v int64) *CospaVoiceControlResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetOnbillbaordTimes(v int32) *CospaVoiceControlResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetEffectValue(v float64) *CospaVoiceControlResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItem) SetVideoList(v []*CospaVoiceControlResponseDataListItemVideoListItem) *CospaVoiceControlResponseDataListItem {
	s.VideoList = v
	return s
}

type CospaVoiceControlResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s CospaVoiceControlResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlResponseDataListItemVideoListItem) SetShareUrl(v string) *CospaVoiceControlResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItemVideoListItem) SetTitle(v string) *CospaVoiceControlResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CospaVoiceControlResponseDataListItemVideoListItem) SetItemCover(v string) *CospaVoiceControlResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type CospaVoiceControlResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CospaVoiceControlResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CospaVoiceControlResponseExtra) GoString() string {
	return s.String()
}

func (s *CospaVoiceControlResponseExtra) SetErrorCode(v int32) *CospaVoiceControlResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CospaVoiceControlResponseExtra) SetDescription(v string) *CospaVoiceControlResponseExtra {
	s.Description = &v
	return s
}

func (s *CospaVoiceControlResponseExtra) SetSubErrorCode(v int32) *CospaVoiceControlResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CospaVoiceControlResponseExtra) SetSubDescription(v string) *CospaVoiceControlResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CospaVoiceControlResponseExtra) SetLogid(v string) *CospaVoiceControlResponseExtra {
	s.Logid = &v
	return s
}

func (s *CospaVoiceControlResponseExtra) SetNow(v int64) *CospaVoiceControlResponseExtra {
	s.Now = &v
	return s
}

type CouponBatchConsumeCouponRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ConsumeTime  *int64             `json:"consume_time,omitempty" xml:"consume_time,omitempty" require:"true"`
	OrderId      *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	ConsumeOutNo *string            `json:"consume_out_no,omitempty" xml:"consume_out_no,omitempty" require:"true"`
	CouponIdList []*string          `json:"coupon_id_list,omitempty" xml:"coupon_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s CouponBatchConsumeCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponBatchConsumeCouponRequest) GoString() string {
	return s.String()
}

func (s *CouponBatchConsumeCouponRequest) SetHeader(v map[string]*string) *CouponBatchConsumeCouponRequest {
	s.Header = v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetAccessToken(v string) *CouponBatchConsumeCouponRequest {
	s.AccessToken = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetConsumeTime(v int64) *CouponBatchConsumeCouponRequest {
	s.ConsumeTime = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetOrderId(v string) *CouponBatchConsumeCouponRequest {
	s.OrderId = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetOpenId(v string) *CouponBatchConsumeCouponRequest {
	s.OpenId = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetAppId(v string) *CouponBatchConsumeCouponRequest {
	s.AppId = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetConsumeOutNo(v string) *CouponBatchConsumeCouponRequest {
	s.ConsumeOutNo = &v
	return s
}

func (s *CouponBatchConsumeCouponRequest) SetCouponIdList(v []*string) *CouponBatchConsumeCouponRequest {
	s.CouponIdList = v
	return s
}

type CouponBatchConsumeCouponResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CouponBatchConsumeCouponResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CouponBatchConsumeCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponBatchConsumeCouponResponse) GoString() string {
	return s.String()
}

func (s *CouponBatchConsumeCouponResponse) SetErrMsg(v string) *CouponBatchConsumeCouponResponse {
	s.ErrMsg = &v
	return s
}

func (s *CouponBatchConsumeCouponResponse) SetLogId(v string) *CouponBatchConsumeCouponResponse {
	s.LogId = &v
	return s
}

func (s *CouponBatchConsumeCouponResponse) SetData(v *CouponBatchConsumeCouponResponseData) *CouponBatchConsumeCouponResponse {
	s.Data = v
	return s
}

func (s *CouponBatchConsumeCouponResponse) SetErrNo(v int32) *CouponBatchConsumeCouponResponse {
	s.ErrNo = &v
	return s
}

type CouponBatchConsumeCouponResponseData struct {
	Results []*CouponBatchConsumeCouponResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s CouponBatchConsumeCouponResponseData) String() string {
	return tea.Prettify(s)
}

func (s CouponBatchConsumeCouponResponseData) GoString() string {
	return s.String()
}

func (s *CouponBatchConsumeCouponResponseData) SetResults(v []*CouponBatchConsumeCouponResponseDataResultsItem) *CouponBatchConsumeCouponResponseData {
	s.Results = v
	return s
}

type CouponBatchConsumeCouponResponseDataResultsItem struct {
	ErrMsg   *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty" require:"true"`
	ErrNo    *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CouponBatchConsumeCouponResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s CouponBatchConsumeCouponResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *CouponBatchConsumeCouponResponseDataResultsItem) SetErrMsg(v string) *CouponBatchConsumeCouponResponseDataResultsItem {
	s.ErrMsg = &v
	return s
}

func (s *CouponBatchConsumeCouponResponseDataResultsItem) SetCouponId(v string) *CouponBatchConsumeCouponResponseDataResultsItem {
	s.CouponId = &v
	return s
}

func (s *CouponBatchConsumeCouponResponseDataResultsItem) SetErrNo(v int32) *CouponBatchConsumeCouponResponseDataResultsItem {
	s.ErrNo = &v
	return s
}

type CouponCreateDeveloperActivityRequest struct {
	ActivityName       *string            `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	CouponMetaId       *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	AppId              *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	MerchantActivityId *string            `json:"merchant_activity_id,omitempty" xml:"merchant_activity_id,omitempty" require:"true"`
	CouponStockNumber  *int64             `json:"coupon_stock_number,omitempty" xml:"coupon_stock_number,omitempty" require:"true"`
	Header             map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CouponCreateDeveloperActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponCreateDeveloperActivityRequest) GoString() string {
	return s.String()
}

func (s *CouponCreateDeveloperActivityRequest) SetActivityName(v string) *CouponCreateDeveloperActivityRequest {
	s.ActivityName = &v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetCouponMetaId(v string) *CouponCreateDeveloperActivityRequest {
	s.CouponMetaId = &v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetAppId(v string) *CouponCreateDeveloperActivityRequest {
	s.AppId = &v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetMerchantActivityId(v string) *CouponCreateDeveloperActivityRequest {
	s.MerchantActivityId = &v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetCouponStockNumber(v int64) *CouponCreateDeveloperActivityRequest {
	s.CouponStockNumber = &v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetHeader(v map[string]*string) *CouponCreateDeveloperActivityRequest {
	s.Header = v
	return s
}

func (s *CouponCreateDeveloperActivityRequest) SetAccessToken(v string) *CouponCreateDeveloperActivityRequest {
	s.AccessToken = &v
	return s
}

type CouponCreateDeveloperActivityResponse struct {
	LogId  *string                                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CouponCreateDeveloperActivityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CouponCreateDeveloperActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponCreateDeveloperActivityResponse) GoString() string {
	return s.String()
}

func (s *CouponCreateDeveloperActivityResponse) SetLogId(v string) *CouponCreateDeveloperActivityResponse {
	s.LogId = &v
	return s
}

func (s *CouponCreateDeveloperActivityResponse) SetData(v *CouponCreateDeveloperActivityResponseData) *CouponCreateDeveloperActivityResponse {
	s.Data = v
	return s
}

func (s *CouponCreateDeveloperActivityResponse) SetErrNo(v int32) *CouponCreateDeveloperActivityResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponCreateDeveloperActivityResponse) SetErrMsg(v string) *CouponCreateDeveloperActivityResponse {
	s.ErrMsg = &v
	return s
}

type CouponCreateDeveloperActivityResponseData struct {
	ActivityId *string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

func (s CouponCreateDeveloperActivityResponseData) String() string {
	return tea.Prettify(s)
}

func (s CouponCreateDeveloperActivityResponseData) GoString() string {
	return s.String()
}

func (s *CouponCreateDeveloperActivityResponseData) SetActivityId(v string) *CouponCreateDeveloperActivityResponseData {
	s.ActivityId = &v
	return s
}

type CouponDeleteCouponMetaRequest struct {
	CouponMetaId   *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	MerchantMetaNo *string            `json:"merchant_meta_no,omitempty" xml:"merchant_meta_no,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId          *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s CouponDeleteCouponMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponDeleteCouponMetaRequest) GoString() string {
	return s.String()
}

func (s *CouponDeleteCouponMetaRequest) SetCouponMetaId(v string) *CouponDeleteCouponMetaRequest {
	s.CouponMetaId = &v
	return s
}

func (s *CouponDeleteCouponMetaRequest) SetMerchantMetaNo(v string) *CouponDeleteCouponMetaRequest {
	s.MerchantMetaNo = &v
	return s
}

func (s *CouponDeleteCouponMetaRequest) SetHeader(v map[string]*string) *CouponDeleteCouponMetaRequest {
	s.Header = v
	return s
}

func (s *CouponDeleteCouponMetaRequest) SetAccessToken(v string) *CouponDeleteCouponMetaRequest {
	s.AccessToken = &v
	return s
}

func (s *CouponDeleteCouponMetaRequest) SetAppId(v string) *CouponDeleteCouponMetaRequest {
	s.AppId = &v
	return s
}

type CouponDeleteCouponMetaResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CouponDeleteCouponMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponDeleteCouponMetaResponse) GoString() string {
	return s.String()
}

func (s *CouponDeleteCouponMetaResponse) SetErrNo(v int32) *CouponDeleteCouponMetaResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponDeleteCouponMetaResponse) SetErrMsg(v string) *CouponDeleteCouponMetaResponse {
	s.ErrMsg = &v
	return s
}

func (s *CouponDeleteCouponMetaResponse) SetLogId(v string) *CouponDeleteCouponMetaResponse {
	s.LogId = &v
	return s
}

type CouponGetTalentCouponRequest struct {
	AccountType   *int32             `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TalentAccount *string            `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CouponMetaId  *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
}

func (s CouponGetTalentCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponGetTalentCouponRequest) GoString() string {
	return s.String()
}

func (s *CouponGetTalentCouponRequest) SetAccountType(v int32) *CouponGetTalentCouponRequest {
	s.AccountType = &v
	return s
}

func (s *CouponGetTalentCouponRequest) SetHeader(v map[string]*string) *CouponGetTalentCouponRequest {
	s.Header = v
	return s
}

func (s *CouponGetTalentCouponRequest) SetAccessToken(v string) *CouponGetTalentCouponRequest {
	s.AccessToken = &v
	return s
}

func (s *CouponGetTalentCouponRequest) SetTalentAccount(v string) *CouponGetTalentCouponRequest {
	s.TalentAccount = &v
	return s
}

func (s *CouponGetTalentCouponRequest) SetOpenId(v string) *CouponGetTalentCouponRequest {
	s.OpenId = &v
	return s
}

func (s *CouponGetTalentCouponRequest) SetAppId(v string) *CouponGetTalentCouponRequest {
	s.AppId = &v
	return s
}

func (s *CouponGetTalentCouponRequest) SetCouponMetaId(v string) *CouponGetTalentCouponRequest {
	s.CouponMetaId = &v
	return s
}

type CouponGetTalentCouponResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CouponGetTalentCouponResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CouponGetTalentCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponGetTalentCouponResponse) GoString() string {
	return s.String()
}

func (s *CouponGetTalentCouponResponse) SetLogId(v string) *CouponGetTalentCouponResponse {
	s.LogId = &v
	return s
}

func (s *CouponGetTalentCouponResponse) SetData(v *CouponGetTalentCouponResponseData) *CouponGetTalentCouponResponse {
	s.Data = v
	return s
}

func (s *CouponGetTalentCouponResponse) SetErrNo(v int32) *CouponGetTalentCouponResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponGetTalentCouponResponse) SetErrMsg(v string) *CouponGetTalentCouponResponse {
	s.ErrMsg = &v
	return s
}

type CouponGetTalentCouponResponseData struct {
	TalentAccount    *string `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	Status           *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	StockLimit       *int64  `json:"stock_limit,omitempty" xml:"stock_limit,omitempty" require:"true"`
	AccountType      *int32  `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	AppId            *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CouponMetaId     *string `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	OpenId           *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	SendScene        *int    `json:"send_scene,omitempty" xml:"send_scene,omitempty"`
	AwardAssignedNum *int64  `json:"award_assigned_num,omitempty" xml:"award_assigned_num,omitempty"`
}

func (s CouponGetTalentCouponResponseData) String() string {
	return tea.Prettify(s)
}

func (s CouponGetTalentCouponResponseData) GoString() string {
	return s.String()
}

func (s *CouponGetTalentCouponResponseData) SetTalentAccount(v string) *CouponGetTalentCouponResponseData {
	s.TalentAccount = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetStatus(v int) *CouponGetTalentCouponResponseData {
	s.Status = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetStockLimit(v int64) *CouponGetTalentCouponResponseData {
	s.StockLimit = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetAccountType(v int32) *CouponGetTalentCouponResponseData {
	s.AccountType = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetAppId(v string) *CouponGetTalentCouponResponseData {
	s.AppId = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetCouponMetaId(v string) *CouponGetTalentCouponResponseData {
	s.CouponMetaId = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetOpenId(v string) *CouponGetTalentCouponResponseData {
	s.OpenId = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetSendScene(v int) *CouponGetTalentCouponResponseData {
	s.SendScene = &v
	return s
}

func (s *CouponGetTalentCouponResponseData) SetAwardAssignedNum(v int64) *CouponGetTalentCouponResponseData {
	s.AwardAssignedNum = &v
	return s
}

type CouponModifyCouponMetaRequest struct {
	CouponMeta  *CouponModifyCouponMetaRequestCouponMeta `json:"coupon_meta,omitempty" xml:"coupon_meta,omitempty" require:"true"`
	BizType     *int                                     `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	IsRenewal   *bool                                    `json:"is_renewal,omitempty" xml:"is_renewal,omitempty"`
	Header      map[string]*string                       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CouponModifyCouponMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponModifyCouponMetaRequest) GoString() string {
	return s.String()
}

func (s *CouponModifyCouponMetaRequest) SetCouponMeta(v *CouponModifyCouponMetaRequestCouponMeta) *CouponModifyCouponMetaRequest {
	s.CouponMeta = v
	return s
}

func (s *CouponModifyCouponMetaRequest) SetBizType(v int) *CouponModifyCouponMetaRequest {
	s.BizType = &v
	return s
}

func (s *CouponModifyCouponMetaRequest) SetIsRenewal(v bool) *CouponModifyCouponMetaRequest {
	s.IsRenewal = &v
	return s
}

func (s *CouponModifyCouponMetaRequest) SetHeader(v map[string]*string) *CouponModifyCouponMetaRequest {
	s.Header = v
	return s
}

func (s *CouponModifyCouponMetaRequest) SetAccessToken(v string) *CouponModifyCouponMetaRequest {
	s.AccessToken = &v
	return s
}

type CouponModifyCouponMetaRequestCouponMeta struct {
	ConsumeDesc         *string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty"`
	ValidBeginTime      *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	RemainedStockNumber *int64  `json:"remained_stock_number,omitempty" xml:"remained_stock_number,omitempty"`
	ReceiveDesc         *string `json:"receive_desc,omitempty" xml:"receive_desc,omitempty"`
	ValidType           *int    `json:"valid_type,omitempty" xml:"valid_type,omitempty"`
	ConsumePath         *string `json:"consume_path,omitempty" xml:"consume_path,omitempty"`
	CouponName          *string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	ReceiveEndTime      *int64  `json:"receive_end_time,omitempty" xml:"receive_end_time,omitempty"`
	ReceiveBeginTime    *int64  `json:"receive_begin_time,omitempty" xml:"receive_begin_time,omitempty"`
	FreeEpNumber        *int64  `json:"free_ep_number,omitempty" xml:"free_ep_number,omitempty"`
	ValidDuration       *int64  `json:"valid_duration,omitempty" xml:"valid_duration,omitempty"`
	CouponMetaId        *string `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	ValidEndTime        *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
}

func (s CouponModifyCouponMetaRequestCouponMeta) String() string {
	return tea.Prettify(s)
}

func (s CouponModifyCouponMetaRequestCouponMeta) GoString() string {
	return s.String()
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetConsumeDesc(v string) *CouponModifyCouponMetaRequestCouponMeta {
	s.ConsumeDesc = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetValidBeginTime(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.ValidBeginTime = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetRemainedStockNumber(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.RemainedStockNumber = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetReceiveDesc(v string) *CouponModifyCouponMetaRequestCouponMeta {
	s.ReceiveDesc = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetValidType(v int) *CouponModifyCouponMetaRequestCouponMeta {
	s.ValidType = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetConsumePath(v string) *CouponModifyCouponMetaRequestCouponMeta {
	s.ConsumePath = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetCouponName(v string) *CouponModifyCouponMetaRequestCouponMeta {
	s.CouponName = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetReceiveEndTime(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.ReceiveEndTime = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetReceiveBeginTime(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.ReceiveBeginTime = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetFreeEpNumber(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.FreeEpNumber = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetValidDuration(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.ValidDuration = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetCouponMetaId(v string) *CouponModifyCouponMetaRequestCouponMeta {
	s.CouponMetaId = &v
	return s
}

func (s *CouponModifyCouponMetaRequestCouponMeta) SetValidEndTime(v int64) *CouponModifyCouponMetaRequestCouponMeta {
	s.ValidEndTime = &v
	return s
}

type CouponModifyCouponMetaResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CouponModifyCouponMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponModifyCouponMetaResponse) GoString() string {
	return s.String()
}

func (s *CouponModifyCouponMetaResponse) SetLogId(v string) *CouponModifyCouponMetaResponse {
	s.LogId = &v
	return s
}

func (s *CouponModifyCouponMetaResponse) SetErrNo(v int32) *CouponModifyCouponMetaResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponModifyCouponMetaResponse) SetErrMsg(v string) *CouponModifyCouponMetaResponse {
	s.ErrMsg = &v
	return s
}

type CouponPriceQueryRequest struct {
	AccountId       *string                      `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OrderCreateTime *int64                       `json:"order_create_time,omitempty" xml:"order_create_time,omitempty" require:"true"`
	OrderId         *string                      `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Base            *CouponPriceQueryRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	Header          map[string]*string           `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                      `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CouponPriceQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryRequest) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryRequest) SetAccountId(v string) *CouponPriceQueryRequest {
	s.AccountId = &v
	return s
}

func (s *CouponPriceQueryRequest) SetOrderCreateTime(v int64) *CouponPriceQueryRequest {
	s.OrderCreateTime = &v
	return s
}

func (s *CouponPriceQueryRequest) SetOrderId(v string) *CouponPriceQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CouponPriceQueryRequest) SetBase(v *CouponPriceQueryRequestBase) *CouponPriceQueryRequest {
	s.Base = v
	return s
}

func (s *CouponPriceQueryRequest) SetHeader(v map[string]*string) *CouponPriceQueryRequest {
	s.Header = v
	return s
}

func (s *CouponPriceQueryRequest) SetAccessToken(v string) *CouponPriceQueryRequest {
	s.AccessToken = &v
	return s
}

type CouponPriceQueryRequestBase struct {
	Extra      map[string]*string                     `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *CouponPriceQueryRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s CouponPriceQueryRequestBase) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryRequestBase) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryRequestBase) SetExtra(v map[string]*string) *CouponPriceQueryRequestBase {
	s.Extra = v
	return s
}

func (s *CouponPriceQueryRequestBase) SetLogID(v string) *CouponPriceQueryRequestBase {
	s.LogID = &v
	return s
}

func (s *CouponPriceQueryRequestBase) SetTrafficEnv(v *CouponPriceQueryRequestBaseTrafficEnv) *CouponPriceQueryRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *CouponPriceQueryRequestBase) SetAddr(v string) *CouponPriceQueryRequestBase {
	s.Addr = &v
	return s
}

func (s *CouponPriceQueryRequestBase) SetCaller(v string) *CouponPriceQueryRequestBase {
	s.Caller = &v
	return s
}

func (s *CouponPriceQueryRequestBase) SetClient(v string) *CouponPriceQueryRequestBase {
	s.Client = &v
	return s
}

type CouponPriceQueryRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s CouponPriceQueryRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryRequestBaseTrafficEnv) SetEnv(v string) *CouponPriceQueryRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *CouponPriceQueryRequestBaseTrafficEnv) SetOpen(v bool) *CouponPriceQueryRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type CouponPriceQueryResponse struct {
	Extra    *CouponPriceQueryResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *CouponPriceQueryResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *CouponPriceQueryResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CouponPriceQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryResponse) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryResponse) SetExtra(v *CouponPriceQueryResponseExtra) *CouponPriceQueryResponse {
	s.Extra = v
	return s
}

func (s *CouponPriceQueryResponse) SetBaseResp(v *CouponPriceQueryResponseBaseResp) *CouponPriceQueryResponse {
	s.BaseResp = v
	return s
}

func (s *CouponPriceQueryResponse) SetData(v *CouponPriceQueryResponseData) *CouponPriceQueryResponse {
	s.Data = v
	return s
}

type CouponPriceQueryResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CouponPriceQueryResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryResponseBaseResp) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryResponseBaseResp) SetStatusMessage(v string) *CouponPriceQueryResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *CouponPriceQueryResponseBaseResp) SetExtra(v map[string]*string) *CouponPriceQueryResponseBaseResp {
	s.Extra = v
	return s
}

func (s *CouponPriceQueryResponseBaseResp) SetStatusCode(v int32) *CouponPriceQueryResponseBaseResp {
	s.StatusCode = &v
	return s
}

type CouponPriceQueryResponseData struct {
	OrderId      *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	DeductAmount *int32  `json:"deduct_amount,omitempty" xml:"deduct_amount,omitempty" require:"true"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode    *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CouponPriceQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryResponseData) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryResponseData) SetOrderId(v string) *CouponPriceQueryResponseData {
	s.OrderId = &v
	return s
}

func (s *CouponPriceQueryResponseData) SetDeductAmount(v int32) *CouponPriceQueryResponseData {
	s.DeductAmount = &v
	return s
}

func (s *CouponPriceQueryResponseData) SetDescription(v string) *CouponPriceQueryResponseData {
	s.Description = &v
	return s
}

func (s *CouponPriceQueryResponseData) SetErrorCode(v int32) *CouponPriceQueryResponseData {
	s.ErrorCode = &v
	return s
}

type CouponPriceQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s CouponPriceQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CouponPriceQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *CouponPriceQueryResponseExtra) SetSubErrorCode(v int32) *CouponPriceQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CouponPriceQueryResponseExtra) SetDescription(v string) *CouponPriceQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *CouponPriceQueryResponseExtra) SetErrorCode(v int32) *CouponPriceQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CouponPriceQueryResponseExtra) SetLogid(v string) *CouponPriceQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *CouponPriceQueryResponseExtra) SetNow(v int64) *CouponPriceQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *CouponPriceQueryResponseExtra) SetSubDescription(v string) *CouponPriceQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type CouponQueryCouponMetaRequest struct {
	BizType      *int               `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	CouponMetaId *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CouponQueryCouponMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponQueryCouponMetaRequest) GoString() string {
	return s.String()
}

func (s *CouponQueryCouponMetaRequest) SetBizType(v int) *CouponQueryCouponMetaRequest {
	s.BizType = &v
	return s
}

func (s *CouponQueryCouponMetaRequest) SetCouponMetaId(v string) *CouponQueryCouponMetaRequest {
	s.CouponMetaId = &v
	return s
}

func (s *CouponQueryCouponMetaRequest) SetHeader(v map[string]*string) *CouponQueryCouponMetaRequest {
	s.Header = v
	return s
}

func (s *CouponQueryCouponMetaRequest) SetAccessToken(v string) *CouponQueryCouponMetaRequest {
	s.AccessToken = &v
	return s
}

type CouponQueryCouponMetaResponse struct {
	CouponMeta *CouponQueryCouponMetaResponseCouponMeta `json:"coupon_meta,omitempty" xml:"coupon_meta,omitempty" require:"true"`
	ErrNo      *int32                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg     *string                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId      *string                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CouponQueryCouponMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponQueryCouponMetaResponse) GoString() string {
	return s.String()
}

func (s *CouponQueryCouponMetaResponse) SetCouponMeta(v *CouponQueryCouponMetaResponseCouponMeta) *CouponQueryCouponMetaResponse {
	s.CouponMeta = v
	return s
}

func (s *CouponQueryCouponMetaResponse) SetErrNo(v int32) *CouponQueryCouponMetaResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponQueryCouponMetaResponse) SetErrMsg(v string) *CouponQueryCouponMetaResponse {
	s.ErrMsg = &v
	return s
}

func (s *CouponQueryCouponMetaResponse) SetLogId(v string) *CouponQueryCouponMetaResponse {
	s.LogId = &v
	return s
}

type CouponQueryCouponMetaResponseCouponMeta struct {
	StockNumber      *int64  `json:"stock_number,omitempty" xml:"stock_number,omitempty" require:"true"`
	DiscountType     *int    `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	ConsumePath      *string `json:"consume_path,omitempty" xml:"consume_path,omitempty"`
	FreeEpNumber     *int64  `json:"free_ep_number,omitempty" xml:"free_ep_number,omitempty"`
	ConsumeDesc      *string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty"`
	Status           *int    `json:"status,omitempty" xml:"status,omitempty"`
	ReceiveBeginTime *int64  `json:"receive_begin_time,omitempty" xml:"receive_begin_time,omitempty" require:"true"`
	ReceiveDesc      *string `json:"receive_desc,omitempty" xml:"receive_desc,omitempty"`
	CallbackUrl      *string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	CouponMetaId     *string `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	RelatedType      *int    `json:"related_type,omitempty" xml:"related_type,omitempty"`
	OriginId         *string `json:"origin_id,omitempty" xml:"origin_id,omitempty"`
	ValidEndTime     *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	ValidType        *int    `json:"valid_type,omitempty" xml:"valid_type,omitempty" require:"true"`
	CouponName       *string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty" require:"true"`
	SecretSource     *int    `json:"secret_source,omitempty" xml:"secret_source,omitempty" require:"true"`
	ValidBeginTime   *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	MerchantMetaNo   *string `json:"merchant_meta_no,omitempty" xml:"merchant_meta_no,omitempty" require:"true"`
	DiscountAmount   *int64  `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	ValidDuration    *int64  `json:"valid_duration,omitempty" xml:"valid_duration,omitempty"`
	MinPayAmount     *int64  `json:"min_pay_amount,omitempty" xml:"min_pay_amount,omitempty"`
	ReceiveEndTime   *int64  `json:"receive_end_time,omitempty" xml:"receive_end_time,omitempty" require:"true"`
}

func (s CouponQueryCouponMetaResponseCouponMeta) String() string {
	return tea.Prettify(s)
}

func (s CouponQueryCouponMetaResponseCouponMeta) GoString() string {
	return s.String()
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetStockNumber(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.StockNumber = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetDiscountType(v int) *CouponQueryCouponMetaResponseCouponMeta {
	s.DiscountType = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetConsumePath(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.ConsumePath = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetFreeEpNumber(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.FreeEpNumber = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetConsumeDesc(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.ConsumeDesc = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetStatus(v int) *CouponQueryCouponMetaResponseCouponMeta {
	s.Status = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetReceiveBeginTime(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.ReceiveBeginTime = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetReceiveDesc(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.ReceiveDesc = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetCallbackUrl(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.CallbackUrl = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetCouponMetaId(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.CouponMetaId = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetRelatedType(v int) *CouponQueryCouponMetaResponseCouponMeta {
	s.RelatedType = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetOriginId(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.OriginId = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetValidEndTime(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.ValidEndTime = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetValidType(v int) *CouponQueryCouponMetaResponseCouponMeta {
	s.ValidType = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetCouponName(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.CouponName = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetSecretSource(v int) *CouponQueryCouponMetaResponseCouponMeta {
	s.SecretSource = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetValidBeginTime(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.ValidBeginTime = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetMerchantMetaNo(v string) *CouponQueryCouponMetaResponseCouponMeta {
	s.MerchantMetaNo = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetDiscountAmount(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.DiscountAmount = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetValidDuration(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.ValidDuration = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetMinPayAmount(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.MinPayAmount = &v
	return s
}

func (s *CouponQueryCouponMetaResponseCouponMeta) SetReceiveEndTime(v int64) *CouponQueryCouponMetaResponseCouponMeta {
	s.ReceiveEndTime = &v
	return s
}

type CouponSetTalentCouponRequest struct {
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TalentAccount   *string            `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	AwardStockLimit *int64             `json:"award_stock_limit,omitempty" xml:"award_stock_limit,omitempty"`
	AppId           *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	CouponMetaId    *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	AccountType     *int32             `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	StockLimit      *int64             `json:"stock_limit,omitempty" xml:"stock_limit,omitempty" require:"true"`
	OpenId          *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s CouponSetTalentCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s CouponSetTalentCouponRequest) GoString() string {
	return s.String()
}

func (s *CouponSetTalentCouponRequest) SetAccessToken(v string) *CouponSetTalentCouponRequest {
	s.AccessToken = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetTalentAccount(v string) *CouponSetTalentCouponRequest {
	s.TalentAccount = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetAwardStockLimit(v int64) *CouponSetTalentCouponRequest {
	s.AwardStockLimit = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetAppId(v string) *CouponSetTalentCouponRequest {
	s.AppId = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetHeader(v map[string]*string) *CouponSetTalentCouponRequest {
	s.Header = v
	return s
}

func (s *CouponSetTalentCouponRequest) SetCouponMetaId(v string) *CouponSetTalentCouponRequest {
	s.CouponMetaId = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetAccountType(v int32) *CouponSetTalentCouponRequest {
	s.AccountType = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetStockLimit(v int64) *CouponSetTalentCouponRequest {
	s.StockLimit = &v
	return s
}

func (s *CouponSetTalentCouponRequest) SetOpenId(v string) *CouponSetTalentCouponRequest {
	s.OpenId = &v
	return s
}

type CouponSetTalentCouponResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CouponSetTalentCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s CouponSetTalentCouponResponse) GoString() string {
	return s.String()
}

func (s *CouponSetTalentCouponResponse) SetLogId(v string) *CouponSetTalentCouponResponse {
	s.LogId = &v
	return s
}

func (s *CouponSetTalentCouponResponse) SetErrNo(v int32) *CouponSetTalentCouponResponse {
	s.ErrNo = &v
	return s
}

func (s *CouponSetTalentCouponResponse) SetErrMsg(v string) *CouponSetTalentCouponResponse {
	s.ErrMsg = &v
	return s
}

type CreateCallbackRequest struct {
	Voucher            *CreateCallbackRequestVoucher                  `json:"voucher,omitempty" xml:"voucher,omitempty"`
	Codes              []*string                                      `json:"codes,omitempty" xml:"codes,omitempty" type:"Repeated"`
	AdditionalInfoList []*CreateCallbackRequestAdditionalInfoListItem `json:"additional_info_list,omitempty" xml:"additional_info_list,omitempty" type:"Repeated"`
	OrderId            *string                                        `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	FailReasonDesc     *string                                        `json:"fail_reason_desc,omitempty" xml:"fail_reason_desc,omitempty"`
	Vouchers           []*CreateCallbackRequestVouchersItem           `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	ThirdOrderId       *string                                        `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
	Result             *int64                                         `json:"result,omitempty" xml:"result,omitempty"`
	FailReason         *string                                        `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	Certificates       []*CreateCallbackRequestCertificatesItem       `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	Header             map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CreateCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequest) SetVoucher(v *CreateCallbackRequestVoucher) *CreateCallbackRequest {
	s.Voucher = v
	return s
}

func (s *CreateCallbackRequest) SetCodes(v []*string) *CreateCallbackRequest {
	s.Codes = v
	return s
}

func (s *CreateCallbackRequest) SetAdditionalInfoList(v []*CreateCallbackRequestAdditionalInfoListItem) *CreateCallbackRequest {
	s.AdditionalInfoList = v
	return s
}

func (s *CreateCallbackRequest) SetOrderId(v string) *CreateCallbackRequest {
	s.OrderId = &v
	return s
}

func (s *CreateCallbackRequest) SetFailReasonDesc(v string) *CreateCallbackRequest {
	s.FailReasonDesc = &v
	return s
}

func (s *CreateCallbackRequest) SetVouchers(v []*CreateCallbackRequestVouchersItem) *CreateCallbackRequest {
	s.Vouchers = v
	return s
}

func (s *CreateCallbackRequest) SetThirdOrderId(v string) *CreateCallbackRequest {
	s.ThirdOrderId = &v
	return s
}

func (s *CreateCallbackRequest) SetResult(v int64) *CreateCallbackRequest {
	s.Result = &v
	return s
}

func (s *CreateCallbackRequest) SetFailReason(v string) *CreateCallbackRequest {
	s.FailReason = &v
	return s
}

func (s *CreateCallbackRequest) SetCertificates(v []*CreateCallbackRequestCertificatesItem) *CreateCallbackRequest {
	s.Certificates = v
	return s
}

func (s *CreateCallbackRequest) SetHeader(v map[string]*string) *CreateCallbackRequest {
	s.Header = v
	return s
}

func (s *CreateCallbackRequest) SetAccessToken(v string) *CreateCallbackRequest {
	s.AccessToken = &v
	return s
}

type CreateCallbackRequestAdditionalInfoListItem struct {
	Key           *string           `json:"key,omitempty" xml:"key,omitempty"`
	AdditionalMap map[int32]*string `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
}

func (s CreateCallbackRequestAdditionalInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestAdditionalInfoListItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestAdditionalInfoListItem) SetKey(v string) *CreateCallbackRequestAdditionalInfoListItem {
	s.Key = &v
	return s
}

func (s *CreateCallbackRequestAdditionalInfoListItem) SetAdditionalMap(v map[int32]*string) *CreateCallbackRequestAdditionalInfoListItem {
	s.AdditionalMap = v
	return s
}

type CreateCallbackRequestCertificatesItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CreateCallbackRequestCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestCertificatesItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestCertificatesItem) SetCertificateId(v string) *CreateCallbackRequestCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *CreateCallbackRequestCertificatesItem) SetCode(v string) *CreateCallbackRequestCertificatesItem {
	s.Code = &v
	return s
}

type CreateCallbackRequestVoucher struct {
	Projects []*CreateCallbackRequestVoucherProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	Entrance *CreateCallbackRequestVoucherEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
}

func (s CreateCallbackRequestVoucher) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVoucher) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVoucher) SetProjects(v []*CreateCallbackRequestVoucherProjectsItem) *CreateCallbackRequestVoucher {
	s.Projects = v
	return s
}

func (s *CreateCallbackRequestVoucher) SetEntrance(v *CreateCallbackRequestVoucherEntrance) *CreateCallbackRequestVoucher {
	s.Entrance = v
	return s
}

type CreateCallbackRequestVoucherEntrance struct {
	Credentials    []*CreateCallbackRequestVoucherEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                              `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                              `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	ProjectId      *string                                                `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                              `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                              `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                              `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
}

func (s CreateCallbackRequestVoucherEntrance) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVoucherEntrance) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVoucherEntrance) SetCredentials(v []*CreateCallbackRequestVoucherEntranceCredentialsItem) *CreateCallbackRequestVoucherEntrance {
	s.Credentials = v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetGmcodeImgs(v []*string) *CreateCallbackRequestVoucherEntrance {
	s.GmcodeImgs = v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetIdCards(v []*string) *CreateCallbackRequestVoucherEntrance {
	s.IdCards = v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetProjectId(v string) *CreateCallbackRequestVoucherEntrance {
	s.ProjectId = &v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetQrcodes(v []*string) *CreateCallbackRequestVoucherEntrance {
	s.Qrcodes = v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetUrls(v []*string) *CreateCallbackRequestVoucherEntrance {
	s.Urls = v
	return s
}

func (s *CreateCallbackRequestVoucherEntrance) SetCertificateNos(v []*string) *CreateCallbackRequestVoucherEntrance {
	s.CertificateNos = v
	return s
}

type CreateCallbackRequestVoucherEntranceCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CreateCallbackRequestVoucherEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVoucherEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVoucherEntranceCredentialsItem) SetCredentialNo(v string) *CreateCallbackRequestVoucherEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *CreateCallbackRequestVoucherEntranceCredentialsItem) SetCredentialType(v int64) *CreateCallbackRequestVoucherEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

type CreateCallbackRequestVoucherProjectsItem struct {
	IdCards        []*string                                                  `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                    `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	ProjectId      *string                                                    `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                  `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                  `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                  `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*CreateCallbackRequestVoucherProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                  `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
}

func (s CreateCallbackRequestVoucherProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVoucherProjectsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetIdCards(v []*string) *CreateCallbackRequestVoucherProjectsItem {
	s.IdCards = v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetName(v string) *CreateCallbackRequestVoucherProjectsItem {
	s.Name = &v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetProjectId(v string) *CreateCallbackRequestVoucherProjectsItem {
	s.ProjectId = &v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetQrcodes(v []*string) *CreateCallbackRequestVoucherProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetUrls(v []*string) *CreateCallbackRequestVoucherProjectsItem {
	s.Urls = v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetCertificateNos(v []*string) *CreateCallbackRequestVoucherProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetCredentials(v []*CreateCallbackRequestVoucherProjectsItemCredentialsItem) *CreateCallbackRequestVoucherProjectsItem {
	s.Credentials = v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItem) SetGmcodeImgs(v []*string) *CreateCallbackRequestVoucherProjectsItem {
	s.GmcodeImgs = v
	return s
}

type CreateCallbackRequestVoucherProjectsItemCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CreateCallbackRequestVoucherProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVoucherProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVoucherProjectsItemCredentialsItem) SetCredentialNo(v string) *CreateCallbackRequestVoucherProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *CreateCallbackRequestVoucherProjectsItemCredentialsItem) SetCredentialType(v int64) *CreateCallbackRequestVoucherProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

type CreateCallbackRequestVouchersItem struct {
	Entrance *CreateCallbackRequestVouchersItemEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
	Projects []*CreateCallbackRequestVouchersItemProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
}

func (s CreateCallbackRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVouchersItem) SetEntrance(v *CreateCallbackRequestVouchersItemEntrance) *CreateCallbackRequestVouchersItem {
	s.Entrance = v
	return s
}

func (s *CreateCallbackRequestVouchersItem) SetProjects(v []*CreateCallbackRequestVouchersItemProjectsItem) *CreateCallbackRequestVouchersItem {
	s.Projects = v
	return s
}

type CreateCallbackRequestVouchersItemEntrance struct {
	Credentials    []*CreateCallbackRequestVouchersItemEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                   `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                   `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	ProjectId      *string                                                     `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                   `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                   `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                   `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
}

func (s CreateCallbackRequestVouchersItemEntrance) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVouchersItemEntrance) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetCredentials(v []*CreateCallbackRequestVouchersItemEntranceCredentialsItem) *CreateCallbackRequestVouchersItemEntrance {
	s.Credentials = v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetGmcodeImgs(v []*string) *CreateCallbackRequestVouchersItemEntrance {
	s.GmcodeImgs = v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetIdCards(v []*string) *CreateCallbackRequestVouchersItemEntrance {
	s.IdCards = v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetProjectId(v string) *CreateCallbackRequestVouchersItemEntrance {
	s.ProjectId = &v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetQrcodes(v []*string) *CreateCallbackRequestVouchersItemEntrance {
	s.Qrcodes = v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetUrls(v []*string) *CreateCallbackRequestVouchersItemEntrance {
	s.Urls = v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntrance) SetCertificateNos(v []*string) *CreateCallbackRequestVouchersItemEntrance {
	s.CertificateNos = v
	return s
}

type CreateCallbackRequestVouchersItemEntranceCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CreateCallbackRequestVouchersItemEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVouchersItemEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVouchersItemEntranceCredentialsItem) SetCredentialNo(v string) *CreateCallbackRequestVouchersItemEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *CreateCallbackRequestVouchersItemEntranceCredentialsItem) SetCredentialType(v int64) *CreateCallbackRequestVouchersItemEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

type CreateCallbackRequestVouchersItemProjectsItem struct {
	CertificateNos []*string                                                       `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*CreateCallbackRequestVouchersItemProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                       `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                       `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                         `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	ProjectId      *string                                                         `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                       `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                       `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s CreateCallbackRequestVouchersItemProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVouchersItemProjectsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetCertificateNos(v []*string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetCredentials(v []*CreateCallbackRequestVouchersItemProjectsItemCredentialsItem) *CreateCallbackRequestVouchersItemProjectsItem {
	s.Credentials = v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetGmcodeImgs(v []*string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.GmcodeImgs = v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetIdCards(v []*string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.IdCards = v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetName(v string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.Name = &v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetProjectId(v string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.ProjectId = &v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetQrcodes(v []*string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItem) SetUrls(v []*string) *CreateCallbackRequestVouchersItemProjectsItem {
	s.Urls = v
	return s
}

type CreateCallbackRequestVouchersItemProjectsItemCredentialsItem struct {
	CredentialType *int64  `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s CreateCallbackRequestVouchersItemProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackRequestVouchersItemProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackRequestVouchersItemProjectsItemCredentialsItem) SetCredentialType(v int64) *CreateCallbackRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *CreateCallbackRequestVouchersItemProjectsItemCredentialsItem) SetCredentialNo(v string) *CreateCallbackRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

type CreateCallbackResponse struct {
	Data                *CreateCallbackResponseData                      `json:"data,omitempty" xml:"data,omitempty"`
	Extra               *CreateCallbackResponseExtra                     `json:"extra,omitempty" xml:"extra,omitempty"`
	CertificateInfoList []*CreateCallbackResponseCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
}

func (s CreateCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponse) SetData(v *CreateCallbackResponseData) *CreateCallbackResponse {
	s.Data = v
	return s
}

func (s *CreateCallbackResponse) SetExtra(v *CreateCallbackResponseExtra) *CreateCallbackResponse {
	s.Extra = v
	return s
}

func (s *CreateCallbackResponse) SetCertificateInfoList(v []*CreateCallbackResponseCertificateInfoListItem) *CreateCallbackResponse {
	s.CertificateInfoList = v
	return s
}

type CreateCallbackResponseCertificateInfoListItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CreateCallbackResponseCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponseCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseCertificateInfoListItem) SetCertificateId(v string) *CreateCallbackResponseCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *CreateCallbackResponseCertificateInfoListItem) SetCode(v string) *CreateCallbackResponseCertificateInfoListItem {
	s.Code = &v
	return s
}

type CreateCallbackResponseData struct {
	GwDescription       *string                                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CertificateInfoList []*CreateCallbackResponseDataCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	GwErrorCode         *int32                                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CreateCallbackResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponseData) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseData) SetGwDescription(v string) *CreateCallbackResponseData {
	s.GwDescription = &v
	return s
}

func (s *CreateCallbackResponseData) SetCertificateInfoList(v []*CreateCallbackResponseDataCertificateInfoListItem) *CreateCallbackResponseData {
	s.CertificateInfoList = v
	return s
}

func (s *CreateCallbackResponseData) SetGwErrorCode(v int32) *CreateCallbackResponseData {
	s.GwErrorCode = &v
	return s
}

type CreateCallbackResponseDataCertificateInfoListItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CreateCallbackResponseDataCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponseDataCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseDataCertificateInfoListItem) SetCertificateId(v string) *CreateCallbackResponseDataCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *CreateCallbackResponseDataCertificateInfoListItem) SetCode(v string) *CreateCallbackResponseDataCertificateInfoListItem {
	s.Code = &v
	return s
}

type CreateCallbackResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CreateCallbackResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CreateCallbackResponseExtra) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponseExtra) SetSubDescription(v string) *CreateCallbackResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CreateCallbackResponseExtra) SetSubErrorCode(v int32) *CreateCallbackResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CreateCallbackResponseExtra) SetDescription(v string) *CreateCallbackResponseExtra {
	s.Description = &v
	return s
}

func (s *CreateCallbackResponseExtra) SetErrorCode(v int32) *CreateCallbackResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CreateCallbackResponseExtra) SetLogid(v string) *CreateCallbackResponseExtra {
	s.Logid = &v
	return s
}

func (s *CreateCallbackResponseExtra) SetNow(v int64) *CreateCallbackResponseExtra {
	s.Now = &v
	return s
}

type CreateMaSubServiceRequest struct {
	SubServiceName *string            `json:"sub_service_name,omitempty" xml:"sub_service_name,omitempty" require:"true"`
	SearchKeyWord  []*string          `json:"search_key_word,omitempty" xml:"search_key_word,omitempty" require:"true" type:"Repeated"`
	StartPageUrl   *string            `json:"start_page_url,omitempty" xml:"start_page_url,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CreateMaSubServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMaSubServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateMaSubServiceRequest) SetSubServiceName(v string) *CreateMaSubServiceRequest {
	s.SubServiceName = &v
	return s
}

func (s *CreateMaSubServiceRequest) SetSearchKeyWord(v []*string) *CreateMaSubServiceRequest {
	s.SearchKeyWord = v
	return s
}

func (s *CreateMaSubServiceRequest) SetStartPageUrl(v string) *CreateMaSubServiceRequest {
	s.StartPageUrl = &v
	return s
}

func (s *CreateMaSubServiceRequest) SetHeader(v map[string]*string) *CreateMaSubServiceRequest {
	s.Header = v
	return s
}

func (s *CreateMaSubServiceRequest) SetAccessToken(v string) *CreateMaSubServiceRequest {
	s.AccessToken = &v
	return s
}

type CreateMaSubServiceResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CreateMaSubServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMaSubServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateMaSubServiceResponse) SetErrMsg(v string) *CreateMaSubServiceResponse {
	s.ErrMsg = &v
	return s
}

func (s *CreateMaSubServiceResponse) SetLogId(v string) *CreateMaSubServiceResponse {
	s.LogId = &v
	return s
}

func (s *CreateMaSubServiceResponse) SetErrNo(v int32) *CreateMaSubServiceResponse {
	s.ErrNo = &v
	return s
}

type CrowdSaveRequest struct {
	AccessToken *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Crowds      []*CrowdSaveRequestCrowdsItem `json:"crowds,omitempty" xml:"crowds,omitempty" type:"Repeated"`
	PoiId       *string                       `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Base        *CrowdSaveRequestBase         `json:"Base,omitempty" xml:"Base,omitempty"`
	Header      map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CrowdSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequest) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequest) SetAccessToken(v string) *CrowdSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *CrowdSaveRequest) SetAccountId(v string) *CrowdSaveRequest {
	s.AccountId = &v
	return s
}

func (s *CrowdSaveRequest) SetCrowds(v []*CrowdSaveRequestCrowdsItem) *CrowdSaveRequest {
	s.Crowds = v
	return s
}

func (s *CrowdSaveRequest) SetPoiId(v string) *CrowdSaveRequest {
	s.PoiId = &v
	return s
}

func (s *CrowdSaveRequest) SetBase(v *CrowdSaveRequestBase) *CrowdSaveRequest {
	s.Base = v
	return s
}

func (s *CrowdSaveRequest) SetHeader(v map[string]*string) *CrowdSaveRequest {
	s.Header = v
	return s
}

type CrowdSaveRequestBase struct {
	Addr       *string                         `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                         `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                         `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string              `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                         `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *CrowdSaveRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
}

func (s CrowdSaveRequestBase) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequestBase) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequestBase) SetAddr(v string) *CrowdSaveRequestBase {
	s.Addr = &v
	return s
}

func (s *CrowdSaveRequestBase) SetCaller(v string) *CrowdSaveRequestBase {
	s.Caller = &v
	return s
}

func (s *CrowdSaveRequestBase) SetClient(v string) *CrowdSaveRequestBase {
	s.Client = &v
	return s
}

func (s *CrowdSaveRequestBase) SetExtra(v map[string]*string) *CrowdSaveRequestBase {
	s.Extra = v
	return s
}

func (s *CrowdSaveRequestBase) SetLogID(v string) *CrowdSaveRequestBase {
	s.LogID = &v
	return s
}

func (s *CrowdSaveRequestBase) SetTrafficEnv(v *CrowdSaveRequestBaseTrafficEnv) *CrowdSaveRequestBase {
	s.TrafficEnv = v
	return s
}

type CrowdSaveRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s CrowdSaveRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequestBaseTrafficEnv) SetOpen(v bool) *CrowdSaveRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *CrowdSaveRequestBaseTrafficEnv) SetEnv(v string) *CrowdSaveRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type CrowdSaveRequestCrowdsItem struct {
	HeightCondition *CrowdSaveRequestCrowdsItemHeightCondition `json:"height_condition,omitempty" xml:"height_condition,omitempty"`
	Note            *string                                    `json:"note,omitempty" xml:"note,omitempty"`
	OpType          *int                                       `json:"op_type,omitempty" xml:"op_type,omitempty"`
	AgeCondition    *CrowdSaveRequestCrowdsItemAgeCondition    `json:"age_condition,omitempty" xml:"age_condition,omitempty"`
	CrowdType       *int                                       `json:"crowd_type,omitempty" xml:"crowd_type,omitempty" require:"true"`
}

func (s CrowdSaveRequestCrowdsItem) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequestCrowdsItem) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequestCrowdsItem) SetHeightCondition(v *CrowdSaveRequestCrowdsItemHeightCondition) *CrowdSaveRequestCrowdsItem {
	s.HeightCondition = v
	return s
}

func (s *CrowdSaveRequestCrowdsItem) SetNote(v string) *CrowdSaveRequestCrowdsItem {
	s.Note = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItem) SetOpType(v int) *CrowdSaveRequestCrowdsItem {
	s.OpType = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItem) SetAgeCondition(v *CrowdSaveRequestCrowdsItemAgeCondition) *CrowdSaveRequestCrowdsItem {
	s.AgeCondition = v
	return s
}

func (s *CrowdSaveRequestCrowdsItem) SetCrowdType(v int) *CrowdSaveRequestCrowdsItem {
	s.CrowdType = &v
	return s
}

type CrowdSaveRequestCrowdsItemAgeCondition struct {
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
}

func (s CrowdSaveRequestCrowdsItemAgeCondition) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequestCrowdsItemAgeCondition) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequestCrowdsItemAgeCondition) SetLowerBound(v int64) *CrowdSaveRequestCrowdsItemAgeCondition {
	s.LowerBound = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemAgeCondition) SetLowerBoundType(v int) *CrowdSaveRequestCrowdsItemAgeCondition {
	s.LowerBoundType = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemAgeCondition) SetUpperBound(v int64) *CrowdSaveRequestCrowdsItemAgeCondition {
	s.UpperBound = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemAgeCondition) SetUpperBoundType(v int) *CrowdSaveRequestCrowdsItemAgeCondition {
	s.UpperBoundType = &v
	return s
}

type CrowdSaveRequestCrowdsItemHeightCondition struct {
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
}

func (s CrowdSaveRequestCrowdsItemHeightCondition) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveRequestCrowdsItemHeightCondition) GoString() string {
	return s.String()
}

func (s *CrowdSaveRequestCrowdsItemHeightCondition) SetLowerBoundType(v int) *CrowdSaveRequestCrowdsItemHeightCondition {
	s.LowerBoundType = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemHeightCondition) SetUpperBound(v int64) *CrowdSaveRequestCrowdsItemHeightCondition {
	s.UpperBound = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemHeightCondition) SetUpperBoundType(v int) *CrowdSaveRequestCrowdsItemHeightCondition {
	s.UpperBoundType = &v
	return s
}

func (s *CrowdSaveRequestCrowdsItemHeightCondition) SetLowerBound(v int64) *CrowdSaveRequestCrowdsItemHeightCondition {
	s.LowerBound = &v
	return s
}

type CrowdSaveResponse struct {
	Extra    *CrowdSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	Data     *CrowdSaveResponseData     `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	BaseResp *CrowdSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty"`
}

func (s CrowdSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveResponse) GoString() string {
	return s.String()
}

func (s *CrowdSaveResponse) SetExtra(v *CrowdSaveResponseExtra) *CrowdSaveResponse {
	s.Extra = v
	return s
}

func (s *CrowdSaveResponse) SetData(v *CrowdSaveResponseData) *CrowdSaveResponse {
	s.Data = v
	return s
}

func (s *CrowdSaveResponse) SetBaseResp(v *CrowdSaveResponseBaseResp) *CrowdSaveResponse {
	s.BaseResp = v
	return s
}

type CrowdSaveResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s CrowdSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *CrowdSaveResponseBaseResp) SetExtra(v map[string]*string) *CrowdSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *CrowdSaveResponseBaseResp) SetStatusCode(v int32) *CrowdSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *CrowdSaveResponseBaseResp) SetStatusMessage(v string) *CrowdSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type CrowdSaveResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CrowdSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveResponseData) GoString() string {
	return s.String()
}

func (s *CrowdSaveResponseData) SetGwErrorCode(v int32) *CrowdSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CrowdSaveResponseData) SetGwDescription(v string) *CrowdSaveResponseData {
	s.GwDescription = &v
	return s
}

type CrowdSaveResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s CrowdSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CrowdSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *CrowdSaveResponseExtra) SetNow(v int64) *CrowdSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *CrowdSaveResponseExtra) SetSubDescription(v string) *CrowdSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CrowdSaveResponseExtra) SetSubErrorCode(v int32) *CrowdSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CrowdSaveResponseExtra) SetDescription(v string) *CrowdSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *CrowdSaveResponseExtra) SetErrorCode(v int32) *CrowdSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CrowdSaveResponseExtra) SetLogid(v string) *CrowdSaveResponseExtra {
	s.Logid = &v
	return s
}

type CustomizationQueryStatusRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s CustomizationQueryStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizationQueryStatusRequest) GoString() string {
	return s.String()
}

func (s *CustomizationQueryStatusRequest) SetOpenId(v string) *CustomizationQueryStatusRequest {
	s.OpenId = &v
	return s
}

func (s *CustomizationQueryStatusRequest) SetOrderId(v string) *CustomizationQueryStatusRequest {
	s.OrderId = &v
	return s
}

func (s *CustomizationQueryStatusRequest) SetHeader(v map[string]*string) *CustomizationQueryStatusRequest {
	s.Header = v
	return s
}

func (s *CustomizationQueryStatusRequest) SetAccessToken(v string) *CustomizationQueryStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *CustomizationQueryStatusRequest) SetAppId(v string) *CustomizationQueryStatusRequest {
	s.AppId = &v
	return s
}

type CustomizationQueryStatusResponse struct {
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CustomizationQueryStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CustomizationQueryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizationQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *CustomizationQueryStatusResponse) SetLogId(v string) *CustomizationQueryStatusResponse {
	s.LogId = &v
	return s
}

func (s *CustomizationQueryStatusResponse) SetData(v *CustomizationQueryStatusResponseData) *CustomizationQueryStatusResponse {
	s.Data = v
	return s
}

func (s *CustomizationQueryStatusResponse) SetErrMsg(v string) *CustomizationQueryStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *CustomizationQueryStatusResponse) SetErrNo(v int32) *CustomizationQueryStatusResponse {
	s.ErrNo = &v
	return s
}

type CustomizationQueryStatusResponseData struct {
	SuccessCode *int64 `json:"success_code,omitempty" xml:"success_code,omitempty" require:"true"`
}

func (s CustomizationQueryStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s CustomizationQueryStatusResponseData) GoString() string {
	return s.String()
}

func (s *CustomizationQueryStatusResponseData) SetSuccessCode(v int64) *CustomizationQueryStatusResponseData {
	s.SuccessCode = &v
	return s
}

type DataAnalysisQueryBehaviorDataRequest struct {
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Os          *string            `json:"os,omitempty" xml:"os,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DataAnalysisQueryBehaviorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryBehaviorDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetStartTime(v int64) *DataAnalysisQueryBehaviorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetEndTime(v int64) *DataAnalysisQueryBehaviorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetHostName(v string) *DataAnalysisQueryBehaviorDataRequest {
	s.HostName = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetOs(v string) *DataAnalysisQueryBehaviorDataRequest {
	s.Os = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetVersionType(v string) *DataAnalysisQueryBehaviorDataRequest {
	s.VersionType = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetHeader(v map[string]*string) *DataAnalysisQueryBehaviorDataRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQueryBehaviorDataRequest) SetAccessToken(v string) *DataAnalysisQueryBehaviorDataRequest {
	s.AccessToken = &v
	return s
}

type DataAnalysisQueryBehaviorDataResponse struct {
	ErrNo  *int32                                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DataAnalysisQueryBehaviorDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DataAnalysisQueryBehaviorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryBehaviorDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryBehaviorDataResponse) SetErrNo(v int32) *DataAnalysisQueryBehaviorDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponse) SetErrMsg(v string) *DataAnalysisQueryBehaviorDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponse) SetLogId(v string) *DataAnalysisQueryBehaviorDataResponse {
	s.LogId = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponse) SetData(v *DataAnalysisQueryBehaviorDataResponseData) *DataAnalysisQueryBehaviorDataResponse {
	s.Data = v
	return s
}

type DataAnalysisQueryBehaviorDataResponseData struct {
	Sum       *DataAnalysisQueryBehaviorDataResponseDataSum             `json:"sum,omitempty" xml:"sum,omitempty"`
	Behaviors []*DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem `json:"behaviors,omitempty" xml:"behaviors,omitempty" type:"Repeated"`
}

func (s DataAnalysisQueryBehaviorDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryBehaviorDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryBehaviorDataResponseData) SetSum(v *DataAnalysisQueryBehaviorDataResponseDataSum) *DataAnalysisQueryBehaviorDataResponseData {
	s.Sum = v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseData) SetBehaviors(v []*DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) *DataAnalysisQueryBehaviorDataResponseData {
	s.Behaviors = v
	return s
}

type DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem struct {
	TotalUserNum    *int64   `json:"total_user_num,omitempty" xml:"total_user_num,omitempty"`
	ShareTime       *int64   `json:"share_time,omitempty" xml:"share_time,omitempty"`
	NewUserNum      *int64   `json:"new_user_num,omitempty" xml:"new_user_num,omitempty"`
	TimeType        *int32   `json:"time_type,omitempty" xml:"time_type,omitempty"`
	ActiveUserNum   *int64   `json:"active_user_num,omitempty" xml:"active_user_num,omitempty"`
	Time            *string  `json:"time,omitempty" xml:"time,omitempty"`
	OpenTime        *int64   `json:"open_time,omitempty" xml:"open_time,omitempty"`
	PerUserOpenTime *float64 `json:"per_user_open_time,omitempty" xml:"per_user_open_time,omitempty"`
	PerUserStayTime *float64 `json:"per_user_stay_time,omitempty" xml:"per_user_stay_time,omitempty"`
	AvgStayTime     *float64 `json:"avg_stay_time,omitempty" xml:"avg_stay_time,omitempty"`
}

func (s DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetTotalUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.TotalUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetShareTime(v int64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.ShareTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetNewUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.NewUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetTimeType(v int32) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetActiveUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.ActiveUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetTime(v string) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetOpenTime(v int64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.OpenTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetPerUserOpenTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.PerUserOpenTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetPerUserStayTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.PerUserStayTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem) SetAvgStayTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataBehaviorsItem {
	s.AvgStayTime = &v
	return s
}

type DataAnalysisQueryBehaviorDataResponseDataSum struct {
	TotalUserNum    *int64   `json:"total_user_num,omitempty" xml:"total_user_num,omitempty"`
	NewUserNum      *int64   `json:"new_user_num,omitempty" xml:"new_user_num,omitempty"`
	AvgStayTime     *float64 `json:"avg_stay_time,omitempty" xml:"avg_stay_time,omitempty"`
	OpenTime        *int64   `json:"open_time,omitempty" xml:"open_time,omitempty"`
	ShareTime       *int64   `json:"share_time,omitempty" xml:"share_time,omitempty"`
	PerUserOpenTime *float64 `json:"per_user_open_time,omitempty" xml:"per_user_open_time,omitempty"`
	ActiveUserNum   *int64   `json:"active_user_num,omitempty" xml:"active_user_num,omitempty"`
	TimeType        *int32   `json:"time_type,omitempty" xml:"time_type,omitempty"`
	PerUserStayTime *float64 `json:"per_user_stay_time,omitempty" xml:"per_user_stay_time,omitempty"`
}

func (s DataAnalysisQueryBehaviorDataResponseDataSum) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryBehaviorDataResponseDataSum) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetTotalUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.TotalUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetNewUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.NewUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetAvgStayTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.AvgStayTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetOpenTime(v int64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.OpenTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetShareTime(v int64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.ShareTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetPerUserOpenTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.PerUserOpenTime = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetActiveUserNum(v int64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.ActiveUserNum = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetTimeType(v int32) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQueryBehaviorDataResponseDataSum) SetPerUserStayTime(v float64) *DataAnalysisQueryBehaviorDataResponseDataSum {
	s.PerUserStayTime = &v
	return s
}

type DataAnalysisQueryClientDataRequest struct {
	UserType    *string            `json:"user_type,omitempty" xml:"user_type,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataRequest) SetUserType(v string) *DataAnalysisQueryClientDataRequest {
	s.UserType = &v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetHostName(v string) *DataAnalysisQueryClientDataRequest {
	s.HostName = &v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetHeader(v map[string]*string) *DataAnalysisQueryClientDataRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetAccessToken(v string) *DataAnalysisQueryClientDataRequest {
	s.AccessToken = &v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetVersionType(v string) *DataAnalysisQueryClientDataRequest {
	s.VersionType = &v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetStartTime(v int64) *DataAnalysisQueryClientDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQueryClientDataRequest) SetEndTime(v int64) *DataAnalysisQueryClientDataRequest {
	s.EndTime = &v
	return s
}

type DataAnalysisQueryClientDataResponse struct {
	LogId  *string                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DataAnalysisQueryClientDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponse) SetLogId(v string) *DataAnalysisQueryClientDataResponse {
	s.LogId = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponse) SetData(v *DataAnalysisQueryClientDataResponseData) *DataAnalysisQueryClientDataResponse {
	s.Data = v
	return s
}

func (s *DataAnalysisQueryClientDataResponse) SetErrNo(v int32) *DataAnalysisQueryClientDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponse) SetErrMsg(v string) *DataAnalysisQueryClientDataResponse {
	s.ErrMsg = &v
	return s
}

type DataAnalysisQueryClientDataResponseData struct {
	LibVersion []*DataAnalysisQueryClientDataResponseDataLibVersionItem `json:"lib_version,omitempty" xml:"lib_version,omitempty" require:"true" type:"Repeated"`
	Model      []*DataAnalysisQueryClientDataResponseDataModelItem      `json:"model,omitempty" xml:"model,omitempty" require:"true" type:"Repeated"`
	AppVersion []*DataAnalysisQueryClientDataResponseDataAppVersionItem `json:"app_version,omitempty" xml:"app_version,omitempty" require:"true" type:"Repeated"`
	Brand      []*DataAnalysisQueryClientDataResponseDataBrandItem      `json:"brand,omitempty" xml:"brand,omitempty" require:"true" type:"Repeated"`
}

func (s DataAnalysisQueryClientDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponseData) SetLibVersion(v []*DataAnalysisQueryClientDataResponseDataLibVersionItem) *DataAnalysisQueryClientDataResponseData {
	s.LibVersion = v
	return s
}

func (s *DataAnalysisQueryClientDataResponseData) SetModel(v []*DataAnalysisQueryClientDataResponseDataModelItem) *DataAnalysisQueryClientDataResponseData {
	s.Model = v
	return s
}

func (s *DataAnalysisQueryClientDataResponseData) SetAppVersion(v []*DataAnalysisQueryClientDataResponseDataAppVersionItem) *DataAnalysisQueryClientDataResponseData {
	s.AppVersion = v
	return s
}

func (s *DataAnalysisQueryClientDataResponseData) SetBrand(v []*DataAnalysisQueryClientDataResponseDataBrandItem) *DataAnalysisQueryClientDataResponseData {
	s.Brand = v
	return s
}

type DataAnalysisQueryClientDataResponseDataAppVersionItem struct {
	Value *int32  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataResponseDataAppVersionItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponseDataAppVersionItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponseDataAppVersionItem) SetValue(v int32) *DataAnalysisQueryClientDataResponseDataAppVersionItem {
	s.Value = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponseDataAppVersionItem) SetName(v string) *DataAnalysisQueryClientDataResponseDataAppVersionItem {
	s.Name = &v
	return s
}

type DataAnalysisQueryClientDataResponseDataBrandItem struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataResponseDataBrandItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponseDataBrandItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponseDataBrandItem) SetName(v string) *DataAnalysisQueryClientDataResponseDataBrandItem {
	s.Name = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponseDataBrandItem) SetValue(v int32) *DataAnalysisQueryClientDataResponseDataBrandItem {
	s.Value = &v
	return s
}

type DataAnalysisQueryClientDataResponseDataLibVersionItem struct {
	Value *int32  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataResponseDataLibVersionItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponseDataLibVersionItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponseDataLibVersionItem) SetValue(v int32) *DataAnalysisQueryClientDataResponseDataLibVersionItem {
	s.Value = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponseDataLibVersionItem) SetName(v string) *DataAnalysisQueryClientDataResponseDataLibVersionItem {
	s.Name = &v
	return s
}

type DataAnalysisQueryClientDataResponseDataModelItem struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s DataAnalysisQueryClientDataResponseDataModelItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryClientDataResponseDataModelItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryClientDataResponseDataModelItem) SetName(v string) *DataAnalysisQueryClientDataResponseDataModelItem {
	s.Name = &v
	return s
}

func (s *DataAnalysisQueryClientDataResponseDataModelItem) SetValue(v int32) *DataAnalysisQueryClientDataResponseDataModelItem {
	s.Value = &v
	return s
}

type DataAnalysisQueryLiveRoomRequest struct {
	AnchorName  *string            `json:"anchor_name,omitempty" xml:"anchor_name,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DataAnalysisQueryLiveRoomRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryLiveRoomRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryLiveRoomRequest) SetAnchorName(v string) *DataAnalysisQueryLiveRoomRequest {
	s.AnchorName = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomRequest) SetHeader(v map[string]*string) *DataAnalysisQueryLiveRoomRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQueryLiveRoomRequest) SetAccessToken(v string) *DataAnalysisQueryLiveRoomRequest {
	s.AccessToken = &v
	return s
}

type DataAnalysisQueryLiveRoomResponse struct {
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DataAnalysisQueryLiveRoomResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DataAnalysisQueryLiveRoomResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryLiveRoomResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryLiveRoomResponse) SetErrMsg(v string) *DataAnalysisQueryLiveRoomResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponse) SetLogId(v string) *DataAnalysisQueryLiveRoomResponse {
	s.LogId = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponse) SetData(v *DataAnalysisQueryLiveRoomResponseData) *DataAnalysisQueryLiveRoomResponse {
	s.Data = v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponse) SetErrNo(v int32) *DataAnalysisQueryLiveRoomResponse {
	s.ErrNo = &v
	return s
}

type DataAnalysisQueryLiveRoomResponseData struct {
	HistoryLiveRoom []*DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem `json:"history_live_room,omitempty" xml:"history_live_room,omitempty" type:"Repeated"`
	CurrentLiveRoom []*DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem `json:"current_live_room,omitempty" xml:"current_live_room,omitempty" type:"Repeated"`
}

func (s DataAnalysisQueryLiveRoomResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryLiveRoomResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryLiveRoomResponseData) SetHistoryLiveRoom(v []*DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) *DataAnalysisQueryLiveRoomResponseData {
	s.HistoryLiveRoom = v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseData) SetCurrentLiveRoom(v []*DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) *DataAnalysisQueryLiveRoomResponseData {
	s.CurrentLiveRoom = v
	return s
}

type DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem struct {
	AnchorName   *string `json:"anchor_name,omitempty" xml:"anchor_name,omitempty"`
	LiveRoomId   *int64  `json:"live_room_id,omitempty" xml:"live_room_id,omitempty"`
	CreateTime   *int64  `json:"create_time,omitempty" xml:"create_time,omitempty"`
	LiveRoomName *string `json:"live_room_name,omitempty" xml:"live_room_name,omitempty"`
}

func (s DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) SetAnchorName(v string) *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem {
	s.AnchorName = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) SetLiveRoomId(v int64) *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem {
	s.LiveRoomId = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) SetCreateTime(v int64) *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem {
	s.CreateTime = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem) SetLiveRoomName(v string) *DataAnalysisQueryLiveRoomResponseDataCurrentLiveRoomItem {
	s.LiveRoomName = &v
	return s
}

type DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem struct {
	LiveRoomName *string `json:"live_room_name,omitempty" xml:"live_room_name,omitempty"`
	AnchorName   *string `json:"anchor_name,omitempty" xml:"anchor_name,omitempty"`
	LiveRoomId   *int64  `json:"live_room_id,omitempty" xml:"live_room_id,omitempty"`
	CreateTime   *int64  `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) SetLiveRoomName(v string) *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem {
	s.LiveRoomName = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) SetAnchorName(v string) *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem {
	s.AnchorName = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) SetLiveRoomId(v int64) *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem {
	s.LiveRoomId = &v
	return s
}

func (s *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem) SetCreateTime(v int64) *DataAnalysisQueryLiveRoomResponseDataHistoryLiveRoomItem {
	s.CreateTime = &v
	return s
}

type DataAnalysisQueryPageDataRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Os          *string            `json:"os,omitempty" xml:"os,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DataAnalysisQueryPageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryPageDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryPageDataRequest) SetAccessToken(v string) *DataAnalysisQueryPageDataRequest {
	s.AccessToken = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetStartTime(v int64) *DataAnalysisQueryPageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetEndTime(v int64) *DataAnalysisQueryPageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetHostName(v string) *DataAnalysisQueryPageDataRequest {
	s.HostName = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetOs(v string) *DataAnalysisQueryPageDataRequest {
	s.Os = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetVersionType(v string) *DataAnalysisQueryPageDataRequest {
	s.VersionType = &v
	return s
}

func (s *DataAnalysisQueryPageDataRequest) SetHeader(v map[string]*string) *DataAnalysisQueryPageDataRequest {
	s.Header = v
	return s
}

type DataAnalysisQueryPageDataResponse struct {
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DataAnalysisQueryPageDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DataAnalysisQueryPageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryPageDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryPageDataResponse) SetErrNo(v int32) *DataAnalysisQueryPageDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponse) SetErrMsg(v string) *DataAnalysisQueryPageDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponse) SetLogId(v string) *DataAnalysisQueryPageDataResponse {
	s.LogId = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponse) SetData(v *DataAnalysisQueryPageDataResponseData) *DataAnalysisQueryPageDataResponse {
	s.Data = v
	return s
}

type DataAnalysisQueryPageDataResponseData struct {
	PageList []*DataAnalysisQueryPageDataResponseDataPageListItem `json:"page_list,omitempty" xml:"page_list,omitempty" type:"Repeated"`
}

func (s DataAnalysisQueryPageDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryPageDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryPageDataResponseData) SetPageList(v []*DataAnalysisQueryPageDataResponseDataPageListItem) *DataAnalysisQueryPageDataResponseData {
	s.PageList = v
	return s
}

type DataAnalysisQueryPageDataResponseDataPageListItem struct {
	PageUv          *int32   `json:"page_uv,omitempty" xml:"page_uv,omitempty"`
	PageAvgStayTime *float64 `json:"page_avg_stay_time,omitempty" xml:"page_avg_stay_time,omitempty"`
	PageShareTime   *int32   `json:"page_share_time,omitempty" xml:"page_share_time,omitempty"`
	Page            *string  `json:"page,omitempty" xml:"page,omitempty"`
	PagePv          *int32   `json:"page_pv,omitempty" xml:"page_pv,omitempty"`
}

func (s DataAnalysisQueryPageDataResponseDataPageListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryPageDataResponseDataPageListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryPageDataResponseDataPageListItem) SetPageUv(v int32) *DataAnalysisQueryPageDataResponseDataPageListItem {
	s.PageUv = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponseDataPageListItem) SetPageAvgStayTime(v float64) *DataAnalysisQueryPageDataResponseDataPageListItem {
	s.PageAvgStayTime = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponseDataPageListItem) SetPageShareTime(v int32) *DataAnalysisQueryPageDataResponseDataPageListItem {
	s.PageShareTime = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponseDataPageListItem) SetPage(v string) *DataAnalysisQueryPageDataResponseDataPageListItem {
	s.Page = &v
	return s
}

func (s *DataAnalysisQueryPageDataResponseDataPageListItem) SetPagePv(v int32) *DataAnalysisQueryPageDataResponseDataPageListItem {
	s.PagePv = &v
	return s
}

type DataAnalysisQueryRetentionDataRequest struct {
	Os          *string            `json:"os,omitempty" xml:"os,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	UserType    *string            `json:"user_type,omitempty" xml:"user_type,omitempty" require:"true"`
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
}

func (s DataAnalysisQueryRetentionDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryRetentionDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryRetentionDataRequest) SetOs(v string) *DataAnalysisQueryRetentionDataRequest {
	s.Os = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetVersionType(v string) *DataAnalysisQueryRetentionDataRequest {
	s.VersionType = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetHeader(v map[string]*string) *DataAnalysisQueryRetentionDataRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetAccessToken(v string) *DataAnalysisQueryRetentionDataRequest {
	s.AccessToken = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetStartTime(v int64) *DataAnalysisQueryRetentionDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetEndTime(v int64) *DataAnalysisQueryRetentionDataRequest {
	s.EndTime = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetUserType(v string) *DataAnalysisQueryRetentionDataRequest {
	s.UserType = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataRequest) SetHostName(v string) *DataAnalysisQueryRetentionDataRequest {
	s.HostName = &v
	return s
}

type DataAnalysisQueryRetentionDataResponse struct {
	Data   *DataAnalysisQueryRetentionDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DataAnalysisQueryRetentionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryRetentionDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryRetentionDataResponse) SetData(v *DataAnalysisQueryRetentionDataResponseData) *DataAnalysisQueryRetentionDataResponse {
	s.Data = v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponse) SetErrNo(v int32) *DataAnalysisQueryRetentionDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponse) SetErrMsg(v string) *DataAnalysisQueryRetentionDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponse) SetLogId(v string) *DataAnalysisQueryRetentionDataResponse {
	s.LogId = &v
	return s
}

type DataAnalysisQueryRetentionDataResponseData struct {
	RetentionDataList []*DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem `json:"retention_data_list,omitempty" xml:"retention_data_list,omitempty" type:"Repeated"`
}

func (s DataAnalysisQueryRetentionDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryRetentionDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryRetentionDataResponseData) SetRetentionDataList(v []*DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) *DataAnalysisQueryRetentionDataResponseData {
	s.RetentionDataList = v
	return s
}

type DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem struct {
	RetentionRateList []*DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem `json:"retention_rate_list,omitempty" xml:"retention_rate_list,omitempty" type:"Repeated"`
	Time              *string                                                                                 `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	ActiveUserNum     *int64                                                                                  `json:"active_user_num,omitempty" xml:"active_user_num,omitempty" require:"true"`
}

func (s DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) SetRetentionRateList(v []*DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem) *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem {
	s.RetentionRateList = v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) SetTime(v string) *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem) SetActiveUserNum(v int64) *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItem {
	s.ActiveUserNum = &v
	return s
}

type DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem struct {
	Day  *int32   `json:"day,omitempty" xml:"day,omitempty" require:"true"`
	Rate *float64 `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
}

func (s DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem) SetDay(v int32) *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem {
	s.Day = &v
	return s
}

func (s *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem) SetRate(v float64) *DataAnalysisQueryRetentionDataResponseDataRetentionDataListItemRetentionRateListItem {
	s.Rate = &v
	return s
}

type DataAnalysisQuerySceneDataRequest struct {
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataRequest) SetHostName(v string) *DataAnalysisQuerySceneDataRequest {
	s.HostName = &v
	return s
}

func (s *DataAnalysisQuerySceneDataRequest) SetHeader(v map[string]*string) *DataAnalysisQuerySceneDataRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQuerySceneDataRequest) SetAccessToken(v string) *DataAnalysisQuerySceneDataRequest {
	s.AccessToken = &v
	return s
}

func (s *DataAnalysisQuerySceneDataRequest) SetVersionType(v string) *DataAnalysisQuerySceneDataRequest {
	s.VersionType = &v
	return s
}

func (s *DataAnalysisQuerySceneDataRequest) SetStartTime(v int64) *DataAnalysisQuerySceneDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQuerySceneDataRequest) SetEndTime(v int64) *DataAnalysisQuerySceneDataRequest {
	s.EndTime = &v
	return s
}

type DataAnalysisQuerySceneDataResponse struct {
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *DataAnalysisQuerySceneDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DataAnalysisQuerySceneDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponse) SetErrNo(v int32) *DataAnalysisQuerySceneDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponse) SetErrMsg(v string) *DataAnalysisQuerySceneDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponse) SetLogId(v string) *DataAnalysisQuerySceneDataResponse {
	s.LogId = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponse) SetData(v *DataAnalysisQuerySceneDataResponseData) *DataAnalysisQuerySceneDataResponse {
	s.Data = v
	return s
}

type DataAnalysisQuerySceneDataResponseData struct {
	SceneList []*DataAnalysisQuerySceneDataResponseDataSceneListItem `json:"scene_list,omitempty" xml:"scene_list,omitempty" type:"Repeated"`
}

func (s DataAnalysisQuerySceneDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseData) SetSceneList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItem) *DataAnalysisQuerySceneDataResponseData {
	s.SceneList = v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItem struct {
	ScenePvList          []*DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem          `json:"scene_pv_list,omitempty" xml:"scene_pv_list,omitempty" type:"Repeated"`
	SceneAvgStayTimeList []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem `json:"scene_avg_stay_time_list,omitempty" xml:"scene_avg_stay_time_list,omitempty" type:"Repeated"`
	SceneOpenPv          *int64                                                                         `json:"scene_open_pv,omitempty" xml:"scene_open_pv,omitempty" require:"true"`
	SceneOpenPvList      []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem      `json:"scene_open_pv_list,omitempty" xml:"scene_open_pv_list,omitempty" type:"Repeated"`
	SceneNewUvList       []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem       `json:"scene_new_uv_list,omitempty" xml:"scene_new_uv_list,omitempty" type:"Repeated"`
	SceneName            *string                                                                        `json:"scene_name,omitempty" xml:"scene_name,omitempty" require:"true"`
	SceneNewUv           *int64                                                                         `json:"scene_new_uv,omitempty" xml:"scene_new_uv,omitempty" require:"true"`
	SceneUvList          []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem          `json:"scene_uv_list,omitempty" xml:"scene_uv_list,omitempty" type:"Repeated"`
	SceneId              *string                                                                        `json:"scene_id,omitempty" xml:"scene_id,omitempty" require:"true"`
	SceneUv              *int64                                                                         `json:"scene_uv,omitempty" xml:"scene_uv,omitempty" require:"true"`
	ScenePv              *int64                                                                         `json:"scene_pv,omitempty" xml:"scene_pv,omitempty" require:"true"`
	SceneAvgStayTime     *float64                                                                       `json:"scene_avg_stay_time,omitempty" xml:"scene_avg_stay_time,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetScenePvList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.ScenePvList = v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneAvgStayTimeList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneAvgStayTimeList = v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneOpenPv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneOpenPv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneOpenPvList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneOpenPvList = v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneNewUvList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneNewUvList = v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneName(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneName = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneNewUv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneNewUv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneUvList(v []*DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneUvList = v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneId(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneId = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneUv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneUv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetScenePv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.ScenePv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItem) SetSceneAvgStayTime(v float64) *DataAnalysisQuerySceneDataResponseDataSceneListItem {
	s.SceneAvgStayTime = &v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem struct {
	SceneAvgStayTime *float64 `json:"scene_avg_stay_time,omitempty" xml:"scene_avg_stay_time,omitempty" require:"true"`
	TimeType         *int32   `json:"time_type,omitempty" xml:"time_type,omitempty" require:"true"`
	Time             *string  `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) SetSceneAvgStayTime(v float64) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem {
	s.SceneAvgStayTime = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) SetTimeType(v int32) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem) SetTime(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneAvgStayTimeListItem {
	s.Time = &v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem struct {
	TimeType   *int32  `json:"time_type,omitempty" xml:"time_type,omitempty" require:"true"`
	Time       *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	SceneNewUv *int64  `json:"scene_new_uv,omitempty" xml:"scene_new_uv,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) SetTimeType(v int32) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) SetTime(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem) SetSceneNewUv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneNewUvListItem {
	s.SceneNewUv = &v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem struct {
	SceneOpenPv *int64  `json:"scene_open_pv,omitempty" xml:"scene_open_pv,omitempty" require:"true"`
	TimeType    *int32  `json:"time_type,omitempty" xml:"time_type,omitempty" require:"true"`
	Time        *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) SetSceneOpenPv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem {
	s.SceneOpenPv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) SetTimeType(v int32) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem) SetTime(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneOpenPvListItem {
	s.Time = &v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem struct {
	TimeType *int32  `json:"time_type,omitempty" xml:"time_type,omitempty" require:"true"`
	Time     *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	ScenePv  *int64  `json:"scene_pv,omitempty" xml:"scene_pv,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) SetTimeType(v int32) *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem {
	s.TimeType = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) SetTime(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem) SetScenePv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItemScenePvListItem {
	s.ScenePv = &v
	return s
}

type DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem struct {
	Time     *string `json:"time,omitempty" xml:"time,omitempty" require:"true"`
	SceneUv  *int64  `json:"scene_uv,omitempty" xml:"scene_uv,omitempty" require:"true"`
	TimeType *int32  `json:"time_type,omitempty" xml:"time_type,omitempty" require:"true"`
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) SetTime(v string) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) SetSceneUv(v int64) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem {
	s.SceneUv = &v
	return s
}

func (s *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem) SetTimeType(v int32) *DataAnalysisQuerySceneDataResponseDataSceneListItemSceneUvListItem {
	s.TimeType = &v
	return s
}

type DataAnalysisQueryVideoDataRequest struct {
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartTime        *int64             `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	QueryBindType    *int32             `json:"query_bind_type,omitempty" xml:"query_bind_type,omitempty"`
	EndTime          *int64             `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	HostName         *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	AwemeShortIdList []*string          `json:"aweme_short_id_list,omitempty" xml:"aweme_short_id_list,omitempty" type:"Repeated"`
	ItemIdList       []*string          `json:"item_id_list,omitempty" xml:"item_id_list,omitempty" type:"Repeated"`
	OpenItemIdList   []*string          `json:"open_item_id_list,omitempty" xml:"open_item_id_list,omitempty" type:"Repeated"`
}

func (s DataAnalysisQueryVideoDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryVideoDataRequest) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryVideoDataRequest) SetHeader(v map[string]*string) *DataAnalysisQueryVideoDataRequest {
	s.Header = v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetAccessToken(v string) *DataAnalysisQueryVideoDataRequest {
	s.AccessToken = &v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetStartTime(v int64) *DataAnalysisQueryVideoDataRequest {
	s.StartTime = &v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetQueryBindType(v int32) *DataAnalysisQueryVideoDataRequest {
	s.QueryBindType = &v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetEndTime(v int64) *DataAnalysisQueryVideoDataRequest {
	s.EndTime = &v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetHostName(v string) *DataAnalysisQueryVideoDataRequest {
	s.HostName = &v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetAwemeShortIdList(v []*string) *DataAnalysisQueryVideoDataRequest {
	s.AwemeShortIdList = v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetItemIdList(v []*string) *DataAnalysisQueryVideoDataRequest {
	s.ItemIdList = v
	return s
}

func (s *DataAnalysisQueryVideoDataRequest) SetOpenItemIdList(v []*string) *DataAnalysisQueryVideoDataRequest {
	s.OpenItemIdList = v
	return s
}

type DataAnalysisQueryVideoDataResponse struct {
	Data   *DataAnalysisQueryVideoDataResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DataAnalysisQueryVideoDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryVideoDataResponse) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryVideoDataResponse) SetData(v *DataAnalysisQueryVideoDataResponseData) *DataAnalysisQueryVideoDataResponse {
	s.Data = v
	return s
}

func (s *DataAnalysisQueryVideoDataResponse) SetErrNo(v int32) *DataAnalysisQueryVideoDataResponse {
	s.ErrNo = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponse) SetErrMsg(v string) *DataAnalysisQueryVideoDataResponse {
	s.ErrMsg = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponse) SetLogId(v string) *DataAnalysisQueryVideoDataResponse {
	s.LogId = &v
	return s
}

type DataAnalysisQueryVideoDataResponseData struct {
	VideoDealDataList     []*DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem `json:"video_deal_data_list,omitempty" xml:"video_deal_data_list,omitempty" type:"Repeated"`
	VideoDealOverviewData *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData   `json:"video_deal_overview_data,omitempty" xml:"video_deal_overview_data,omitempty" require:"true"`
}

func (s DataAnalysisQueryVideoDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryVideoDataResponseData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryVideoDataResponseData) SetVideoDealDataList(v []*DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) *DataAnalysisQueryVideoDataResponseData {
	s.VideoDealDataList = v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseData) SetVideoDealOverviewData(v *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) *DataAnalysisQueryVideoDataResponseData {
	s.VideoDealOverviewData = v
	return s
}

type DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem struct {
	MpShowUv          *int64  `json:"mp_show_uv,omitempty" xml:"mp_show_uv,omitempty" require:"true"`
	RefundOrderCnt    *int64  `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
	RefundCustomerCnt *int64  `json:"refund_customer_cnt,omitempty" xml:"refund_customer_cnt,omitempty" require:"true"`
	RefundAmount      *int64  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
	MpDrainagePv      *int64  `json:"mp_drainage_pv,omitempty" xml:"mp_drainage_pv,omitempty" require:"true"`
	CreateOrderAmount *int64  `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty"`
	ValidClueCnt      *int64  `json:"valid_clue_cnt,omitempty" xml:"valid_clue_cnt,omitempty"`
	ClueCnt           *int64  `json:"clue_cnt,omitempty" xml:"clue_cnt,omitempty"`
	MpClickPv         *int64  `json:"mp_click_pv,omitempty" xml:"mp_click_pv,omitempty" require:"true"`
	Time              *string `json:"time,omitempty" xml:"time,omitempty"`
	CustomerOncePrice *int64  `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
	PayOrderAmount    *int64  `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	PayCustomerCnt    *int64  `json:"pay_customer_cnt,omitempty" xml:"pay_customer_cnt,omitempty" require:"true"`
	OrderOncePrice    *int64  `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	PayOrderCnt       *int64  `json:"pay_order_cnt,omitempty" xml:"pay_order_cnt,omitempty" require:"true"`
	VideoPlayUserCnt  *int64  `json:"video_play_user_cnt,omitempty" xml:"video_play_user_cnt,omitempty" require:"true"`
	MpDrainageUv      *int64  `json:"mp_drainage_uv,omitempty" xml:"mp_drainage_uv,omitempty" require:"true"`
	CreateOrderCnt    *int64  `json:"create_order_cnt,omitempty" xml:"create_order_cnt,omitempty" require:"true"`
	CreateCustomerCnt *int64  `json:"create_customer_cnt,omitempty" xml:"create_customer_cnt,omitempty"`
	MpClickUv         *int64  `json:"mp_click_uv,omitempty" xml:"mp_click_uv,omitempty" require:"true"`
	VideoPlayCnt      *int64  `json:"video_play_cnt,omitempty" xml:"video_play_cnt,omitempty" require:"true"`
	MpShowPv          *int64  `json:"mp_show_pv,omitempty" xml:"mp_show_pv,omitempty" require:"true"`
}

func (s DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpShowUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpShowUv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetRefundOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.RefundOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetRefundCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.RefundCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetRefundAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.RefundAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpDrainagePv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpDrainagePv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetCreateOrderAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.CreateOrderAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetValidClueCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.ValidClueCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetClueCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.ClueCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpClickPv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpClickPv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetTime(v string) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.Time = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetCustomerOncePrice(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.CustomerOncePrice = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetPayOrderAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.PayOrderAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetPayCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.PayCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetOrderOncePrice(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.OrderOncePrice = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetPayOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.PayOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetVideoPlayUserCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.VideoPlayUserCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpDrainageUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpDrainageUv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetCreateOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.CreateOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetCreateCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.CreateCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpClickUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpClickUv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetVideoPlayCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.VideoPlayCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem) SetMpShowPv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealDataListItem {
	s.MpShowPv = &v
	return s
}

type DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData struct {
	PayOrderCnt       *int64 `json:"pay_order_cnt,omitempty" xml:"pay_order_cnt,omitempty" require:"true"`
	CreateCustomerCnt *int64 `json:"create_customer_cnt,omitempty" xml:"create_customer_cnt,omitempty"`
	MpClickPv         *int64 `json:"mp_click_pv,omitempty" xml:"mp_click_pv,omitempty" require:"true"`
	CreateOrderCnt    *int64 `json:"create_order_cnt,omitempty" xml:"create_order_cnt,omitempty" require:"true"`
	MpDrainagePv      *int64 `json:"mp_drainage_pv,omitempty" xml:"mp_drainage_pv,omitempty" require:"true"`
	CreateOrderAmount *int64 `json:"create_order_amount,omitempty" xml:"create_order_amount,omitempty"`
	VideoPlayCnt      *int64 `json:"video_play_cnt,omitempty" xml:"video_play_cnt,omitempty" require:"true"`
	ValidClueCnt      *int64 `json:"valid_clue_cnt,omitempty" xml:"valid_clue_cnt,omitempty"`
	ClueCnt           *int64 `json:"clue_cnt,omitempty" xml:"clue_cnt,omitempty"`
	MpDrainageUv      *int64 `json:"mp_drainage_uv,omitempty" xml:"mp_drainage_uv,omitempty" require:"true"`
	PayOrderAmount    *int64 `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty" require:"true"`
	OrderOncePrice    *int64 `json:"order_once_price,omitempty" xml:"order_once_price,omitempty" require:"true"`
	RefundAmount      *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty" require:"true"`
	PayCustomerCnt    *int64 `json:"pay_customer_cnt,omitempty" xml:"pay_customer_cnt,omitempty" require:"true"`
	MpShowPv          *int64 `json:"mp_show_pv,omitempty" xml:"mp_show_pv,omitempty" require:"true"`
	MpClickUv         *int64 `json:"mp_click_uv,omitempty" xml:"mp_click_uv,omitempty" require:"true"`
	VideoPlayUserCnt  *int64 `json:"video_play_user_cnt,omitempty" xml:"video_play_user_cnt,omitempty" require:"true"`
	RefundOrderCnt    *int64 `json:"refund_order_cnt,omitempty" xml:"refund_order_cnt,omitempty" require:"true"`
	CustomerOncePrice *int64 `json:"customer_once_price,omitempty" xml:"customer_once_price,omitempty" require:"true"`
	RefundCustomerCnt *int64 `json:"refund_customer_cnt,omitempty" xml:"refund_customer_cnt,omitempty" require:"true"`
	MpShowUv          *int64 `json:"mp_show_uv,omitempty" xml:"mp_show_uv,omitempty" require:"true"`
}

func (s DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) String() string {
	return tea.Prettify(s)
}

func (s DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) GoString() string {
	return s.String()
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetPayOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.PayOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetCreateCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.CreateCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpClickPv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpClickPv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetCreateOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.CreateOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpDrainagePv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpDrainagePv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetCreateOrderAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.CreateOrderAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetVideoPlayCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.VideoPlayCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetValidClueCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.ValidClueCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetClueCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.ClueCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpDrainageUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpDrainageUv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetPayOrderAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.PayOrderAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetOrderOncePrice(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.OrderOncePrice = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetRefundAmount(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.RefundAmount = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetPayCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.PayCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpShowPv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpShowPv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpClickUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpClickUv = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetVideoPlayUserCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.VideoPlayUserCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetRefundOrderCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.RefundOrderCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetCustomerOncePrice(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.CustomerOncePrice = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetRefundCustomerCnt(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.RefundCustomerCnt = &v
	return s
}

func (s *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData) SetMpShowUv(v int64) *DataAnalysisQueryVideoDataResponseDataVideoDealOverviewData {
	s.MpShowUv = &v
	return s
}

type DefaultSetRequest struct {
	CardType    *DefaultSetRequestCardType `json:"card_type,omitempty" xml:"card_type,omitempty"`
	Value       *string                    `json:"value,omitempty" xml:"value,omitempty"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DefaultSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DefaultSetRequest) GoString() string {
	return s.String()
}

func (s *DefaultSetRequest) SetCardType(v *DefaultSetRequestCardType) *DefaultSetRequest {
	s.CardType = v
	return s
}

func (s *DefaultSetRequest) SetValue(v string) *DefaultSetRequest {
	s.Value = &v
	return s
}

func (s *DefaultSetRequest) SetHeader(v map[string]*string) *DefaultSetRequest {
	s.Header = v
	return s
}

func (s *DefaultSetRequest) SetAccessToken(v string) *DefaultSetRequest {
	s.AccessToken = &v
	return s
}

type DefaultSetRequestCardType struct {
}

func (s DefaultSetRequestCardType) String() string {
	return tea.Prettify(s)
}

func (s DefaultSetRequestCardType) GoString() string {
	return s.String()
}

type DefaultSetResponse struct {
	Status   *int    `json:"status,omitempty" xml:"status,omitempty"`
	Feedback *string `json:"feedback,omitempty" xml:"feedback,omitempty"`
}

func (s DefaultSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DefaultSetResponse) GoString() string {
	return s.String()
}

func (s *DefaultSetResponse) SetStatus(v int) *DefaultSetResponse {
	s.Status = &v
	return s
}

func (s *DefaultSetResponse) SetFeedback(v string) *DefaultSetResponse {
	s.Feedback = &v
	return s
}

type DelRetainConsultCardRequest struct {
	CardId      *string            `json:"card_id,omitempty" xml:"card_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DelRetainConsultCardRequest) String() string {
	return tea.Prettify(s)
}

func (s DelRetainConsultCardRequest) GoString() string {
	return s.String()
}

func (s *DelRetainConsultCardRequest) SetCardId(v string) *DelRetainConsultCardRequest {
	s.CardId = &v
	return s
}

func (s *DelRetainConsultCardRequest) SetOpenId(v string) *DelRetainConsultCardRequest {
	s.OpenId = &v
	return s
}

func (s *DelRetainConsultCardRequest) SetHeader(v map[string]*string) *DelRetainConsultCardRequest {
	s.Header = v
	return s
}

func (s *DelRetainConsultCardRequest) SetAccessToken(v string) *DelRetainConsultCardRequest {
	s.AccessToken = &v
	return s
}

type DelRetainConsultCardResponse struct {
	Extra *DelRetainConsultCardResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *DelRetainConsultCardResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DelRetainConsultCardResponse) String() string {
	return tea.Prettify(s)
}

func (s DelRetainConsultCardResponse) GoString() string {
	return s.String()
}

func (s *DelRetainConsultCardResponse) SetExtra(v *DelRetainConsultCardResponseExtra) *DelRetainConsultCardResponse {
	s.Extra = v
	return s
}

func (s *DelRetainConsultCardResponse) SetData(v *DelRetainConsultCardResponseData) *DelRetainConsultCardResponse {
	s.Data = v
	return s
}

type DelRetainConsultCardResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DelRetainConsultCardResponseData) String() string {
	return tea.Prettify(s)
}

func (s DelRetainConsultCardResponseData) GoString() string {
	return s.String()
}

func (s *DelRetainConsultCardResponseData) SetGwErrorCode(v int32) *DelRetainConsultCardResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DelRetainConsultCardResponseData) SetGwDescription(v string) *DelRetainConsultCardResponseData {
	s.GwDescription = &v
	return s
}

type DelRetainConsultCardResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DelRetainConsultCardResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DelRetainConsultCardResponseExtra) GoString() string {
	return s.String()
}

func (s *DelRetainConsultCardResponseExtra) SetErrorCode(v int32) *DelRetainConsultCardResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DelRetainConsultCardResponseExtra) SetLogid(v string) *DelRetainConsultCardResponseExtra {
	s.Logid = &v
	return s
}

func (s *DelRetainConsultCardResponseExtra) SetNow(v int64) *DelRetainConsultCardResponseExtra {
	s.Now = &v
	return s
}

func (s *DelRetainConsultCardResponseExtra) SetSubDescription(v string) *DelRetainConsultCardResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DelRetainConsultCardResponseExtra) SetSubErrorCode(v int32) *DelRetainConsultCardResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DelRetainConsultCardResponseExtra) SetDescription(v string) *DelRetainConsultCardResponseExtra {
	s.Description = &v
	return s
}

type DeleteAppTestRelationRequest struct {
	Operator    *string            `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	RefIdList   []*string          `json:"ref_id_list,omitempty" xml:"ref_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s DeleteAppTestRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTestRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppTestRelationRequest) SetOperator(v string) *DeleteAppTestRelationRequest {
	s.Operator = &v
	return s
}

func (s *DeleteAppTestRelationRequest) SetHeader(v map[string]*string) *DeleteAppTestRelationRequest {
	s.Header = v
	return s
}

func (s *DeleteAppTestRelationRequest) SetAccessToken(v string) *DeleteAppTestRelationRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteAppTestRelationRequest) SetType(v string) *DeleteAppTestRelationRequest {
	s.Type = &v
	return s
}

func (s *DeleteAppTestRelationRequest) SetRefIdList(v []*string) *DeleteAppTestRelationRequest {
	s.RefIdList = v
	return s
}

type DeleteAppTestRelationResponse struct {
	Extra *DeleteAppTestRelationResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *DeleteAppTestRelationResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeleteAppTestRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTestRelationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppTestRelationResponse) SetExtra(v *DeleteAppTestRelationResponseExtra) *DeleteAppTestRelationResponse {
	s.Extra = v
	return s
}

func (s *DeleteAppTestRelationResponse) SetData(v *DeleteAppTestRelationResponseData) *DeleteAppTestRelationResponse {
	s.Data = v
	return s
}

type DeleteAppTestRelationResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DeleteAppTestRelationResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTestRelationResponseData) GoString() string {
	return s.String()
}

func (s *DeleteAppTestRelationResponseData) SetGwErrorCode(v int32) *DeleteAppTestRelationResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DeleteAppTestRelationResponseData) SetGwDescription(v string) *DeleteAppTestRelationResponseData {
	s.GwDescription = &v
	return s
}

type DeleteAppTestRelationResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s DeleteAppTestRelationResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppTestRelationResponseExtra) GoString() string {
	return s.String()
}

func (s *DeleteAppTestRelationResponseExtra) SetNow(v int64) *DeleteAppTestRelationResponseExtra {
	s.Now = &v
	return s
}

func (s *DeleteAppTestRelationResponseExtra) SetErrorCode(v int32) *DeleteAppTestRelationResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAppTestRelationResponseExtra) SetDescription(v string) *DeleteAppTestRelationResponseExtra {
	s.Description = &v
	return s
}

func (s *DeleteAppTestRelationResponseExtra) SetSubErrorCode(v int32) *DeleteAppTestRelationResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DeleteAppTestRelationResponseExtra) SetSubDescription(v string) *DeleteAppTestRelationResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *DeleteAppTestRelationResponseExtra) SetLogid(v string) *DeleteAppTestRelationResponseExtra {
	s.Logid = &v
	return s
}

type DeleteAwemeVideoKeywordRequest struct {
	KeywordId   *string            `json:"keyword_id,omitempty" xml:"keyword_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeleteAwemeVideoKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAwemeVideoKeywordRequest) GoString() string {
	return s.String()
}

func (s *DeleteAwemeVideoKeywordRequest) SetKeywordId(v string) *DeleteAwemeVideoKeywordRequest {
	s.KeywordId = &v
	return s
}

func (s *DeleteAwemeVideoKeywordRequest) SetHeader(v map[string]*string) *DeleteAwemeVideoKeywordRequest {
	s.Header = v
	return s
}

func (s *DeleteAwemeVideoKeywordRequest) SetAccessToken(v string) *DeleteAwemeVideoKeywordRequest {
	s.AccessToken = &v
	return s
}

type DeleteAwemeVideoKeywordResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s DeleteAwemeVideoKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAwemeVideoKeywordResponse) GoString() string {
	return s.String()
}

func (s *DeleteAwemeVideoKeywordResponse) SetErrMsg(v string) *DeleteAwemeVideoKeywordResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeleteAwemeVideoKeywordResponse) SetLogId(v string) *DeleteAwemeVideoKeywordResponse {
	s.LogId = &v
	return s
}

func (s *DeleteAwemeVideoKeywordResponse) SetErrNo(v int32) *DeleteAwemeVideoKeywordResponse {
	s.ErrNo = &v
	return s
}

type DeleteOrientedPlanTalentRequest struct {
	PlanId      *int64             `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	DouyinId    *string            `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s DeleteOrientedPlanTalentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrientedPlanTalentRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrientedPlanTalentRequest) SetPlanId(v int64) *DeleteOrientedPlanTalentRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteOrientedPlanTalentRequest) SetDouyinId(v string) *DeleteOrientedPlanTalentRequest {
	s.DouyinId = &v
	return s
}

func (s *DeleteOrientedPlanTalentRequest) SetHeader(v map[string]*string) *DeleteOrientedPlanTalentRequest {
	s.Header = v
	return s
}

func (s *DeleteOrientedPlanTalentRequest) SetAccessToken(v string) *DeleteOrientedPlanTalentRequest {
	s.AccessToken = &v
	return s
}

type DeleteOrientedPlanTalentResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s DeleteOrientedPlanTalentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrientedPlanTalentResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrientedPlanTalentResponse) SetErrMsg(v string) *DeleteOrientedPlanTalentResponse {
	s.ErrMsg = &v
	return s
}

func (s *DeleteOrientedPlanTalentResponse) SetErrNo(v int32) *DeleteOrientedPlanTalentResponse {
	s.ErrNo = &v
	return s
}

func (s *DeleteOrientedPlanTalentResponse) SetLogId(v string) *DeleteOrientedPlanTalentResponse {
	s.LogId = &v
	return s
}

type DeletePurchaseInfoRequest struct {
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ServiceId     *string            `json:"service_id,omitempty" xml:"service_id,omitempty" require:"true"`
	ServiceModeId *string            `json:"service_mode_id,omitempty" xml:"service_mode_id,omitempty" require:"true"`
	OutTradeNo    *string            `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty" require:"true"`
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s DeletePurchaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchaseInfoRequest) GoString() string {
	return s.String()
}

func (s *DeletePurchaseInfoRequest) SetAccessToken(v string) *DeletePurchaseInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *DeletePurchaseInfoRequest) SetServiceId(v string) *DeletePurchaseInfoRequest {
	s.ServiceId = &v
	return s
}

func (s *DeletePurchaseInfoRequest) SetServiceModeId(v string) *DeletePurchaseInfoRequest {
	s.ServiceModeId = &v
	return s
}

func (s *DeletePurchaseInfoRequest) SetOutTradeNo(v string) *DeletePurchaseInfoRequest {
	s.OutTradeNo = &v
	return s
}

func (s *DeletePurchaseInfoRequest) SetOpenId(v string) *DeletePurchaseInfoRequest {
	s.OpenId = &v
	return s
}

func (s *DeletePurchaseInfoRequest) SetHeader(v map[string]*string) *DeletePurchaseInfoRequest {
	s.Header = v
	return s
}

type DeletePurchaseInfoResponse struct {
	Extra *DeletePurchaseInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *DeletePurchaseInfoResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s DeletePurchaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchaseInfoResponse) GoString() string {
	return s.String()
}

func (s *DeletePurchaseInfoResponse) SetExtra(v *DeletePurchaseInfoResponseExtra) *DeletePurchaseInfoResponse {
	s.Extra = v
	return s
}

func (s *DeletePurchaseInfoResponse) SetData(v *DeletePurchaseInfoResponseData) *DeletePurchaseInfoResponse {
	s.Data = v
	return s
}

type DeletePurchaseInfoResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s DeletePurchaseInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchaseInfoResponseData) GoString() string {
	return s.String()
}

func (s *DeletePurchaseInfoResponseData) SetGwErrorCode(v int32) *DeletePurchaseInfoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *DeletePurchaseInfoResponseData) SetGwDescription(v string) *DeletePurchaseInfoResponseData {
	s.GwDescription = &v
	return s
}

type DeletePurchaseInfoResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s DeletePurchaseInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s DeletePurchaseInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *DeletePurchaseInfoResponseExtra) SetLogid(v string) *DeletePurchaseInfoResponseExtra {
	s.Logid = &v
	return s
}

func (s *DeletePurchaseInfoResponseExtra) SetNow(v int64) *DeletePurchaseInfoResponseExtra {
	s.Now = &v
	return s
}

func (s *DeletePurchaseInfoResponseExtra) SetErrorCode(v int32) *DeletePurchaseInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *DeletePurchaseInfoResponseExtra) SetDescription(v string) *DeletePurchaseInfoResponseExtra {
	s.Description = &v
	return s
}

func (s *DeletePurchaseInfoResponseExtra) SetSubErrorCode(v int32) *DeletePurchaseInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *DeletePurchaseInfoResponseExtra) SetSubDescription(v string) *DeletePurchaseInfoResponseExtra {
	s.SubDescription = &v
	return s
}

type DeleteSimpleQrBindRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	QrUrl       *string            `json:"qr_url,omitempty" xml:"qr_url,omitempty" require:"true"`
}

func (s DeleteSimpleQrBindRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimpleQrBindRequest) GoString() string {
	return s.String()
}

func (s *DeleteSimpleQrBindRequest) SetHeader(v map[string]*string) *DeleteSimpleQrBindRequest {
	s.Header = v
	return s
}

func (s *DeleteSimpleQrBindRequest) SetAccessToken(v string) *DeleteSimpleQrBindRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteSimpleQrBindRequest) SetQrUrl(v string) *DeleteSimpleQrBindRequest {
	s.QrUrl = &v
	return s
}

type DeleteSimpleQrBindResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s DeleteSimpleQrBindResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimpleQrBindResponse) GoString() string {
	return s.String()
}

func (s *DeleteSimpleQrBindResponse) SetLogId(v string) *DeleteSimpleQrBindResponse {
	s.LogId = &v
	return s
}

func (s *DeleteSimpleQrBindResponse) SetErrNo(v int32) *DeleteSimpleQrBindResponse {
	s.ErrNo = &v
	return s
}

func (s *DeleteSimpleQrBindResponse) SetErrMsg(v string) *DeleteSimpleQrBindResponse {
	s.ErrMsg = &v
	return s
}

type DetailTripRequest struct {
	Cursor        *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RootAccountId *string            `json:"root_account_id,omitempty" xml:"root_account_id,omitempty"`
	Size          *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	WithdrawId    *string            `json:"withdraw_id,omitempty" xml:"withdraw_id,omitempty"`
	BillDate      *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
	BizType       *int               `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

func (s DetailTripRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailTripRequest) GoString() string {
	return s.String()
}

func (s *DetailTripRequest) SetCursor(v string) *DetailTripRequest {
	s.Cursor = &v
	return s
}

func (s *DetailTripRequest) SetHeader(v map[string]*string) *DetailTripRequest {
	s.Header = v
	return s
}

func (s *DetailTripRequest) SetAccessToken(v string) *DetailTripRequest {
	s.AccessToken = &v
	return s
}

func (s *DetailTripRequest) SetRootAccountId(v string) *DetailTripRequest {
	s.RootAccountId = &v
	return s
}

func (s *DetailTripRequest) SetSize(v int64) *DetailTripRequest {
	s.Size = &v
	return s
}

func (s *DetailTripRequest) SetAccountId(v string) *DetailTripRequest {
	s.AccountId = &v
	return s
}

func (s *DetailTripRequest) SetWithdrawId(v string) *DetailTripRequest {
	s.WithdrawId = &v
	return s
}

func (s *DetailTripRequest) SetBillDate(v string) *DetailTripRequest {
	s.BillDate = &v
	return s
}

func (s *DetailTripRequest) SetBizType(v int) *DetailTripRequest {
	s.BizType = &v
	return s
}

type DetailTripResponse struct {
	Data  *DetailTripResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *DetailTripResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s DetailTripResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailTripResponse) GoString() string {
	return s.String()
}

func (s *DetailTripResponse) SetData(v *DetailTripResponseData) *DetailTripResponse {
	s.Data = v
	return s
}

func (s *DetailTripResponse) SetExtra(v *DetailTripResponseExtra) *DetailTripResponse {
	s.Extra = v
	return s
}

type DetailTripResponseData struct {
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor        *string                                    `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                      `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	LedgerRecords []*DetailTripResponseDataLedgerRecordsItem `json:"ledger_records,omitempty" xml:"ledger_records,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s DetailTripResponseData) String() string {
	return tea.Prettify(s)
}

func (s DetailTripResponseData) GoString() string {
	return s.String()
}

func (s *DetailTripResponseData) SetGwDescription(v string) *DetailTripResponseData {
	s.GwDescription = &v
	return s
}

func (s *DetailTripResponseData) SetCursor(v string) *DetailTripResponseData {
	s.Cursor = &v
	return s
}

func (s *DetailTripResponseData) SetHasMore(v bool) *DetailTripResponseData {
	s.HasMore = &v
	return s
}

func (s *DetailTripResponseData) SetLedgerRecords(v []*DetailTripResponseDataLedgerRecordsItem) *DetailTripResponseData {
	s.LedgerRecords = v
	return s
}

func (s *DetailTripResponseData) SetGwErrorCode(v int32) *DetailTripResponseData {
	s.GwErrorCode = &v
	return s
}
