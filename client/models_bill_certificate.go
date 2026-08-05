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

type BillQueryLegerUrlResponse struct {
	Data  *BillQueryLegerUrlResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *BillQueryLegerUrlResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Urls  []*string                       `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s BillQueryLegerUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s BillQueryLegerUrlResponse) GoString() string {
	return s.String()
}

func (s *BillQueryLegerUrlResponse) SetData(v *BillQueryLegerUrlResponseData) *BillQueryLegerUrlResponse {
	s.Data = v
	return s
}

func (s *BillQueryLegerUrlResponse) SetExtra(v *BillQueryLegerUrlResponseExtra) *BillQueryLegerUrlResponse {
	s.Extra = v
	return s
}

func (s *BillQueryLegerUrlResponse) SetUrls(v []*string) *BillQueryLegerUrlResponse {
	s.Urls = v
	return s
}

type BillQueryLegerUrlResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BillQueryLegerUrlResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillQueryLegerUrlResponseData) GoString() string {
	return s.String()
}

func (s *BillQueryLegerUrlResponseData) SetGwErrorCode(v int32) *BillQueryLegerUrlResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BillQueryLegerUrlResponseData) SetGwDescription(v string) *BillQueryLegerUrlResponseData {
	s.GwDescription = &v
	return s
}

type BillQueryLegerUrlResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s BillQueryLegerUrlResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillQueryLegerUrlResponseExtra) GoString() string {
	return s.String()
}

func (s *BillQueryLegerUrlResponseExtra) SetSubErrorCode(v int32) *BillQueryLegerUrlResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BillQueryLegerUrlResponseExtra) SetDescription(v string) *BillQueryLegerUrlResponseExtra {
	s.Description = &v
	return s
}

func (s *BillQueryLegerUrlResponseExtra) SetErrorCode(v int32) *BillQueryLegerUrlResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BillQueryLegerUrlResponseExtra) SetLogid(v string) *BillQueryLegerUrlResponseExtra {
	s.Logid = &v
	return s
}

func (s *BillQueryLegerUrlResponseExtra) SetNow(v int64) *BillQueryLegerUrlResponseExtra {
	s.Now = &v
	return s
}

func (s *BillQueryLegerUrlResponseExtra) SetSubDescription(v string) *BillQueryLegerUrlResponseExtra {
	s.SubDescription = &v
	return s
}

type BillQueryRebateUrlRequest struct {
	BizMonth    *string            `json:"biz_month,omitempty" xml:"biz_month,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BillQueryRebateUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s BillQueryRebateUrlRequest) GoString() string {
	return s.String()
}

func (s *BillQueryRebateUrlRequest) SetBizMonth(v string) *BillQueryRebateUrlRequest {
	s.BizMonth = &v
	return s
}

func (s *BillQueryRebateUrlRequest) SetAccountId(v string) *BillQueryRebateUrlRequest {
	s.AccountId = &v
	return s
}

func (s *BillQueryRebateUrlRequest) SetHeader(v map[string]*string) *BillQueryRebateUrlRequest {
	s.Header = v
	return s
}

func (s *BillQueryRebateUrlRequest) SetAccessToken(v string) *BillQueryRebateUrlRequest {
	s.AccessToken = &v
	return s
}

type BillQueryRebateUrlResponse struct {
	Urls  []*string                        `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	Data  *BillQueryRebateUrlResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *BillQueryRebateUrlResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s BillQueryRebateUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s BillQueryRebateUrlResponse) GoString() string {
	return s.String()
}

func (s *BillQueryRebateUrlResponse) SetUrls(v []*string) *BillQueryRebateUrlResponse {
	s.Urls = v
	return s
}

func (s *BillQueryRebateUrlResponse) SetData(v *BillQueryRebateUrlResponseData) *BillQueryRebateUrlResponse {
	s.Data = v
	return s
}

func (s *BillQueryRebateUrlResponse) SetExtra(v *BillQueryRebateUrlResponseExtra) *BillQueryRebateUrlResponse {
	s.Extra = v
	return s
}

type BillQueryRebateUrlResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BillQueryRebateUrlResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillQueryRebateUrlResponseData) GoString() string {
	return s.String()
}

func (s *BillQueryRebateUrlResponseData) SetGwDescription(v string) *BillQueryRebateUrlResponseData {
	s.GwDescription = &v
	return s
}

func (s *BillQueryRebateUrlResponseData) SetGwErrorCode(v int32) *BillQueryRebateUrlResponseData {
	s.GwErrorCode = &v
	return s
}

type BillQueryRebateUrlResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s BillQueryRebateUrlResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillQueryRebateUrlResponseExtra) GoString() string {
	return s.String()
}

func (s *BillQueryRebateUrlResponseExtra) SetSubDescription(v string) *BillQueryRebateUrlResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BillQueryRebateUrlResponseExtra) SetSubErrorCode(v int32) *BillQueryRebateUrlResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BillQueryRebateUrlResponseExtra) SetDescription(v string) *BillQueryRebateUrlResponseExtra {
	s.Description = &v
	return s
}

func (s *BillQueryRebateUrlResponseExtra) SetErrorCode(v int32) *BillQueryRebateUrlResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BillQueryRebateUrlResponseExtra) SetLogid(v string) *BillQueryRebateUrlResponseExtra {
	s.Logid = &v
	return s
}

func (s *BillQueryRebateUrlResponseExtra) SetNow(v int64) *BillQueryRebateUrlResponseExtra {
	s.Now = &v
	return s
}

type BillboardPropRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BillboardPropRequest) String() string {
	return tea.Prettify(s)
}

func (s BillboardPropRequest) GoString() string {
	return s.String()
}

func (s *BillboardPropRequest) SetHeader(v map[string]*string) *BillboardPropRequest {
	s.Header = v
	return s
}

func (s *BillboardPropRequest) SetAccessToken(v string) *BillboardPropRequest {
	s.AccessToken = &v
	return s
}

type BillboardPropResponse struct {
	Extra *BillboardPropResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *BillboardPropResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BillboardPropResponse) String() string {
	return tea.Prettify(s)
}

func (s BillboardPropResponse) GoString() string {
	return s.String()
}

func (s *BillboardPropResponse) SetExtra(v *BillboardPropResponseExtra) *BillboardPropResponse {
	s.Extra = v
	return s
}

func (s *BillboardPropResponse) SetData(v *BillboardPropResponseData) *BillboardPropResponse {
	s.Data = v
	return s
}

type BillboardPropResponseData struct {
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*BillboardPropResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BillboardPropResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillboardPropResponseData) GoString() string {
	return s.String()
}

func (s *BillboardPropResponseData) SetGwDescription(v string) *BillboardPropResponseData {
	s.GwDescription = &v
	return s
}

func (s *BillboardPropResponseData) SetList(v []*BillboardPropResponseDataListItem) *BillboardPropResponseData {
	s.List = v
	return s
}

func (s *BillboardPropResponseData) SetGwErrorCode(v int32) *BillboardPropResponseData {
	s.GwErrorCode = &v
	return s
}

type BillboardPropResponseDataListItem struct {
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	DailyCollectionCnt *float64 `json:"daily_collection_cnt,omitempty" xml:"daily_collection_cnt,omitempty" require:"true"`
	DailyIssueCnt      *float64 `json:"daily_issue_cnt,omitempty" xml:"daily_issue_cnt,omitempty" require:"true"`
	RankChange         *string  `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Rank               *int32   `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	ShootCnt           *float64 `json:"shoot_cnt,omitempty" xml:"shoot_cnt,omitempty" require:"true"`
	DailyIssuePercent  *string  `json:"daily_issue_percent,omitempty" xml:"daily_issue_percent,omitempty" require:"true"`
	ShowCnt            *float64 `json:"show_cnt,omitempty" xml:"show_cnt,omitempty" require:"true"`
	EffectValue        *float64 `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	DailyPlayCnt       *float64 `json:"daily_play_cnt,omitempty" xml:"daily_play_cnt,omitempty" require:"true"`
}

func (s BillboardPropResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s BillboardPropResponseDataListItem) GoString() string {
	return s.String()
}

func (s *BillboardPropResponseDataListItem) SetName(v string) *BillboardPropResponseDataListItem {
	s.Name = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetDailyCollectionCnt(v float64) *BillboardPropResponseDataListItem {
	s.DailyCollectionCnt = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetDailyIssueCnt(v float64) *BillboardPropResponseDataListItem {
	s.DailyIssueCnt = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetRankChange(v string) *BillboardPropResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetRank(v int32) *BillboardPropResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetShootCnt(v float64) *BillboardPropResponseDataListItem {
	s.ShootCnt = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetDailyIssuePercent(v string) *BillboardPropResponseDataListItem {
	s.DailyIssuePercent = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetShowCnt(v float64) *BillboardPropResponseDataListItem {
	s.ShowCnt = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetEffectValue(v float64) *BillboardPropResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *BillboardPropResponseDataListItem) SetDailyPlayCnt(v float64) *BillboardPropResponseDataListItem {
	s.DailyPlayCnt = &v
	return s
}

type BillboardPropResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s BillboardPropResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillboardPropResponseExtra) GoString() string {
	return s.String()
}

func (s *BillboardPropResponseExtra) SetSubDescription(v string) *BillboardPropResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BillboardPropResponseExtra) SetLogid(v string) *BillboardPropResponseExtra {
	s.Logid = &v
	return s
}

func (s *BillboardPropResponseExtra) SetNow(v int64) *BillboardPropResponseExtra {
	s.Now = &v
	return s
}

func (s *BillboardPropResponseExtra) SetErrorCode(v int32) *BillboardPropResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BillboardPropResponseExtra) SetDescription(v string) *BillboardPropResponseExtra {
	s.Description = &v
	return s
}

func (s *BillboardPropResponseExtra) SetSubErrorCode(v int32) *BillboardPropResponseExtra {
	s.SubErrorCode = &v
	return s
}

type BillboardStarsRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BillboardStarsRequest) String() string {
	return tea.Prettify(s)
}

func (s BillboardStarsRequest) GoString() string {
	return s.String()
}

func (s *BillboardStarsRequest) SetHeader(v map[string]*string) *BillboardStarsRequest {
	s.Header = v
	return s
}

func (s *BillboardStarsRequest) SetAccessToken(v string) *BillboardStarsRequest {
	s.AccessToken = &v
	return s
}

type BillboardStarsResponse struct {
	Data  *BillboardStarsResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *BillboardStarsResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s BillboardStarsResponse) String() string {
	return tea.Prettify(s)
}

func (s BillboardStarsResponse) GoString() string {
	return s.String()
}

func (s *BillboardStarsResponse) SetData(v *BillboardStarsResponseData) *BillboardStarsResponse {
	s.Data = v
	return s
}

func (s *BillboardStarsResponse) SetExtra(v *BillboardStarsResponseExtra) *BillboardStarsResponse {
	s.Extra = v
	return s
}

type BillboardStarsResponseData struct {
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*BillboardStarsResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s BillboardStarsResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillboardStarsResponseData) GoString() string {
	return s.String()
}

func (s *BillboardStarsResponseData) SetGwErrorCode(v int32) *BillboardStarsResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BillboardStarsResponseData) SetGwDescription(v string) *BillboardStarsResponseData {
	s.GwDescription = &v
	return s
}

func (s *BillboardStarsResponseData) SetList(v []*BillboardStarsResponseDataListItem) *BillboardStarsResponseData {
	s.List = v
	return s
}

type BillboardStarsResponseDataListItem struct {
	EffectValue *float64 `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	Rank        *int32   `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	Nickname    *string  `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar      *string  `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s BillboardStarsResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s BillboardStarsResponseDataListItem) GoString() string {
	return s.String()
}

func (s *BillboardStarsResponseDataListItem) SetEffectValue(v float64) *BillboardStarsResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *BillboardStarsResponseDataListItem) SetRank(v int32) *BillboardStarsResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *BillboardStarsResponseDataListItem) SetNickname(v string) *BillboardStarsResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *BillboardStarsResponseDataListItem) SetAvatar(v string) *BillboardStarsResponseDataListItem {
	s.Avatar = &v
	return s
}

type BillboardStarsResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BillboardStarsResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillboardStarsResponseExtra) GoString() string {
	return s.String()
}

func (s *BillboardStarsResponseExtra) SetDescription(v string) *BillboardStarsResponseExtra {
	s.Description = &v
	return s
}

func (s *BillboardStarsResponseExtra) SetSubErrorCode(v int32) *BillboardStarsResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BillboardStarsResponseExtra) SetSubDescription(v string) *BillboardStarsResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BillboardStarsResponseExtra) SetLogid(v string) *BillboardStarsResponseExtra {
	s.Logid = &v
	return s
}

func (s *BillboardStarsResponseExtra) SetNow(v int64) *BillboardStarsResponseExtra {
	s.Now = &v
	return s
}

func (s *BillboardStarsResponseExtra) SetErrorCode(v int32) *BillboardStarsResponseExtra {
	s.ErrorCode = &v
	return s
}

type BillboardTopicRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s BillboardTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s BillboardTopicRequest) GoString() string {
	return s.String()
}

func (s *BillboardTopicRequest) SetAccessToken(v string) *BillboardTopicRequest {
	s.AccessToken = &v
	return s
}

func (s *BillboardTopicRequest) SetHeader(v map[string]*string) *BillboardTopicRequest {
	s.Header = v
	return s
}

type BillboardTopicResponse struct {
	Extra *BillboardTopicResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *BillboardTopicResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BillboardTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s BillboardTopicResponse) GoString() string {
	return s.String()
}

func (s *BillboardTopicResponse) SetExtra(v *BillboardTopicResponseExtra) *BillboardTopicResponse {
	s.Extra = v
	return s
}

func (s *BillboardTopicResponse) SetData(v *BillboardTopicResponseData) *BillboardTopicResponse {
	s.Data = v
	return s
}

type BillboardTopicResponseData struct {
	List          []*BillboardTopicResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BillboardTopicResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillboardTopicResponseData) GoString() string {
	return s.String()
}

func (s *BillboardTopicResponseData) SetList(v []*BillboardTopicResponseDataListItem) *BillboardTopicResponseData {
	s.List = v
	return s
}

func (s *BillboardTopicResponseData) SetGwErrorCode(v int32) *BillboardTopicResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BillboardTopicResponseData) SetGwDescription(v string) *BillboardTopicResponseData {
	s.GwDescription = &v
	return s
}

type BillboardTopicResponseDataListItem struct {
	Rank        *int32   `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange  *string  `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Title       *string  `json:"title,omitempty" xml:"title,omitempty"`
	EffectValue *float64 `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
}

func (s BillboardTopicResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s BillboardTopicResponseDataListItem) GoString() string {
	return s.String()
}

func (s *BillboardTopicResponseDataListItem) SetRank(v int32) *BillboardTopicResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *BillboardTopicResponseDataListItem) SetRankChange(v string) *BillboardTopicResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *BillboardTopicResponseDataListItem) SetTitle(v string) *BillboardTopicResponseDataListItem {
	s.Title = &v
	return s
}

func (s *BillboardTopicResponseDataListItem) SetEffectValue(v float64) *BillboardTopicResponseDataListItem {
	s.EffectValue = &v
	return s
}

type BillboardTopicResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s BillboardTopicResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillboardTopicResponseExtra) GoString() string {
	return s.String()
}

func (s *BillboardTopicResponseExtra) SetNow(v int64) *BillboardTopicResponseExtra {
	s.Now = &v
	return s
}

func (s *BillboardTopicResponseExtra) SetErrorCode(v int32) *BillboardTopicResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BillboardTopicResponseExtra) SetDescription(v string) *BillboardTopicResponseExtra {
	s.Description = &v
	return s
}

func (s *BillboardTopicResponseExtra) SetSubErrorCode(v int32) *BillboardTopicResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BillboardTopicResponseExtra) SetSubDescription(v string) *BillboardTopicResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BillboardTopicResponseExtra) SetLogid(v string) *BillboardTopicResponseExtra {
	s.Logid = &v
	return s
}

type BindDetailRequest struct {
	Base               *BindDetailRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId          *int64                 `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	CouponProductId    *int64                 `json:"coupon_product_id,omitempty" xml:"coupon_product_id,omitempty"`
	OutCouponProductId *string                `json:"out_coupon_product_id,omitempty" xml:"out_coupon_product_id,omitempty"`
	Header             map[string]*string     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BindDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s BindDetailRequest) GoString() string {
	return s.String()
}

func (s *BindDetailRequest) SetBase(v *BindDetailRequestBase) *BindDetailRequest {
	s.Base = v
	return s
}

func (s *BindDetailRequest) SetAccountId(v int64) *BindDetailRequest {
	s.AccountId = &v
	return s
}

func (s *BindDetailRequest) SetCouponProductId(v int64) *BindDetailRequest {
	s.CouponProductId = &v
	return s
}

func (s *BindDetailRequest) SetOutCouponProductId(v string) *BindDetailRequest {
	s.OutCouponProductId = &v
	return s
}

func (s *BindDetailRequest) SetHeader(v map[string]*string) *BindDetailRequest {
	s.Header = v
	return s
}

func (s *BindDetailRequest) SetAccessToken(v string) *BindDetailRequest {
	s.AccessToken = &v
	return s
}

type BindDetailRequestBase struct {
	Extra      map[string]*string               `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                          `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *BindDetailRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                          `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                          `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                          `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s BindDetailRequestBase) String() string {
	return tea.Prettify(s)
}

func (s BindDetailRequestBase) GoString() string {
	return s.String()
}

func (s *BindDetailRequestBase) SetExtra(v map[string]*string) *BindDetailRequestBase {
	s.Extra = v
	return s
}

func (s *BindDetailRequestBase) SetLogID(v string) *BindDetailRequestBase {
	s.LogID = &v
	return s
}

func (s *BindDetailRequestBase) SetTrafficEnv(v *BindDetailRequestBaseTrafficEnv) *BindDetailRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *BindDetailRequestBase) SetAddr(v string) *BindDetailRequestBase {
	s.Addr = &v
	return s
}

func (s *BindDetailRequestBase) SetCaller(v string) *BindDetailRequestBase {
	s.Caller = &v
	return s
}

func (s *BindDetailRequestBase) SetClient(v string) *BindDetailRequestBase {
	s.Client = &v
	return s
}

type BindDetailRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s BindDetailRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s BindDetailRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *BindDetailRequestBaseTrafficEnv) SetOpen(v bool) *BindDetailRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *BindDetailRequestBaseTrafficEnv) SetEnv(v string) *BindDetailRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type BindDetailResponse struct {
	Data     *BindDetailResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *BindDetailResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *BindDetailResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s BindDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponse) GoString() string {
	return s.String()
}

func (s *BindDetailResponse) SetData(v *BindDetailResponseData) *BindDetailResponse {
	s.Data = v
	return s
}

func (s *BindDetailResponse) SetExtra(v *BindDetailResponseExtra) *BindDetailResponse {
	s.Extra = v
	return s
}

func (s *BindDetailResponse) SetBaseResp(v *BindDetailResponseBaseResp) *BindDetailResponse {
	s.BaseResp = v
	return s
}

type BindDetailResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s BindDetailResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseBaseResp) GoString() string {
	return s.String()
}

func (s *BindDetailResponseBaseResp) SetStatusMessage(v string) *BindDetailResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *BindDetailResponseBaseResp) SetExtra(v map[string]*string) *BindDetailResponseBaseResp {
	s.Extra = v
	return s
}

func (s *BindDetailResponseBaseResp) SetStatusCode(v int32) *BindDetailResponseBaseResp {
	s.StatusCode = &v
	return s
}

type BindDetailResponseData struct {
	Description     *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode       *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GroupList       []*BindDetailResponseDataGroupListItem `json:"group_list,omitempty" xml:"group_list,omitempty" type:"Repeated"`
	CouponProductId *int64                                 `json:"coupon_product_id,omitempty" xml:"coupon_product_id,omitempty"`
}

func (s BindDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseData) GoString() string {
	return s.String()
}

func (s *BindDetailResponseData) SetDescription(v string) *BindDetailResponseData {
	s.Description = &v
	return s
}

func (s *BindDetailResponseData) SetErrorCode(v int32) *BindDetailResponseData {
	s.ErrorCode = &v
	return s
}

func (s *BindDetailResponseData) SetGroupList(v []*BindDetailResponseDataGroupListItem) *BindDetailResponseData {
	s.GroupList = v
	return s
}

func (s *BindDetailResponseData) SetCouponProductId(v int64) *BindDetailResponseData {
	s.CouponProductId = &v
	return s
}

type BindDetailResponseDataGroupListItem struct {
	OptionCount *int32                                             `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32                                             `json:"total_count,omitempty" xml:"total_count,omitempty"`
	GroupName   *string                                            `json:"group_name,omitempty" xml:"group_name,omitempty" require:"true"`
	ItemList    []*BindDetailResponseDataGroupListItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s BindDetailResponseDataGroupListItem) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseDataGroupListItem) GoString() string {
	return s.String()
}

func (s *BindDetailResponseDataGroupListItem) SetOptionCount(v int32) *BindDetailResponseDataGroupListItem {
	s.OptionCount = &v
	return s
}

func (s *BindDetailResponseDataGroupListItem) SetTotalCount(v int32) *BindDetailResponseDataGroupListItem {
	s.TotalCount = &v
	return s
}

func (s *BindDetailResponseDataGroupListItem) SetGroupName(v string) *BindDetailResponseDataGroupListItem {
	s.GroupName = &v
	return s
}

func (s *BindDetailResponseDataGroupListItem) SetItemList(v []*BindDetailResponseDataGroupListItemItemListItem) *BindDetailResponseDataGroupListItem {
	s.ItemList = v
	return s
}

type BindDetailResponseDataGroupListItemItemListItem struct {
	Count       *int32                                                         `json:"count,omitempty" xml:"count,omitempty"`
	Price       *int64                                                         `json:"price,omitempty" xml:"price,omitempty"`
	SpuRelation *BindDetailResponseDataGroupListItemItemListItemSpuRelation    `json:"spu_relation,omitempty" xml:"spu_relation,omitempty"`
	Unit        *string                                                        `json:"unit,omitempty" xml:"unit,omitempty"`
	AttrList    []*BindDetailResponseDataGroupListItemItemListItemAttrListItem `json:"attr_list,omitempty" xml:"attr_list,omitempty" type:"Repeated"`
}

func (s BindDetailResponseDataGroupListItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseDataGroupListItemItemListItem) GoString() string {
	return s.String()
}

func (s *BindDetailResponseDataGroupListItemItemListItem) SetCount(v int32) *BindDetailResponseDataGroupListItemItemListItem {
	s.Count = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItem) SetPrice(v int64) *BindDetailResponseDataGroupListItemItemListItem {
	s.Price = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItem) SetSpuRelation(v *BindDetailResponseDataGroupListItemItemListItemSpuRelation) *BindDetailResponseDataGroupListItemItemListItem {
	s.SpuRelation = v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItem) SetUnit(v string) *BindDetailResponseDataGroupListItemItemListItem {
	s.Unit = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItem) SetAttrList(v []*BindDetailResponseDataGroupListItemItemListItemAttrListItem) *BindDetailResponseDataGroupListItemItemListItem {
	s.AttrList = v
	return s
}

type BindDetailResponseDataGroupListItemItemListItemAttrListItem struct {
	AttrKey   *string `json:"attr_key,omitempty" xml:"attr_key,omitempty" require:"true"`
	AttrValue *string `json:"attr_value,omitempty" xml:"attr_value,omitempty" require:"true"`
}

func (s BindDetailResponseDataGroupListItemItemListItemAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseDataGroupListItemItemListItemAttrListItem) GoString() string {
	return s.String()
}

func (s *BindDetailResponseDataGroupListItemItemListItemAttrListItem) SetAttrKey(v string) *BindDetailResponseDataGroupListItemItemListItemAttrListItem {
	s.AttrKey = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemAttrListItem) SetAttrValue(v string) *BindDetailResponseDataGroupListItemItemListItemAttrListItem {
	s.AttrValue = &v
	return s
}

type BindDetailResponseDataGroupListItemItemListItemSpuRelation struct {
	SpecList     []*BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem `json:"spec_list,omitempty" xml:"spec_list,omitempty" type:"Repeated"`
	SpuId        *int64                                                                    `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SpuPriceHigh *int64                                                                    `json:"spu_price_high,omitempty" xml:"spu_price_high,omitempty"`
	SpuPriceLow  *int64                                                                    `json:"spu_price_low,omitempty" xml:"spu_price_low,omitempty"`
	OutSpuId     *string                                                                   `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
}

func (s BindDetailResponseDataGroupListItemItemListItemSpuRelation) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseDataGroupListItemItemListItemSpuRelation) GoString() string {
	return s.String()
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelation) SetSpecList(v []*BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem) *BindDetailResponseDataGroupListItemItemListItemSpuRelation {
	s.SpecList = v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelation) SetSpuId(v int64) *BindDetailResponseDataGroupListItemItemListItemSpuRelation {
	s.SpuId = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelation) SetSpuPriceHigh(v int64) *BindDetailResponseDataGroupListItemItemListItemSpuRelation {
	s.SpuPriceHigh = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelation) SetSpuPriceLow(v int64) *BindDetailResponseDataGroupListItemItemListItemSpuRelation {
	s.SpuPriceLow = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelation) SetOutSpuId(v string) *BindDetailResponseDataGroupListItemItemListItemSpuRelation {
	s.OutSpuId = &v
	return s
}

type BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem struct {
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
}

func (s BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem) GoString() string {
	return s.String()
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem) SetItemKey(v string) *BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem {
	s.ItemKey = &v
	return s
}

func (s *BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem) SetGroupCode(v string) *BindDetailResponseDataGroupListItemItemListItemSpuRelationSpecListItem {
	s.GroupCode = &v
	return s
}

type BindDetailResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s BindDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BindDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *BindDetailResponseExtra) SetSubErrorCode(v int32) *BindDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BindDetailResponseExtra) SetDescription(v string) *BindDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *BindDetailResponseExtra) SetErrorCode(v int32) *BindDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BindDetailResponseExtra) SetLogid(v string) *BindDetailResponseExtra {
	s.Logid = &v
	return s
}

func (s *BindDetailResponseExtra) SetNow(v int64) *BindDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *BindDetailResponseExtra) SetSubDescription(v string) *BindDetailResponseExtra {
	s.SubDescription = &v
	return s
}

type BindInfoAllRequest struct {
	Size        *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
}

func (s BindInfoAllRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInfoAllRequest) GoString() string {
	return s.String()
}

func (s *BindInfoAllRequest) SetSize(v int64) *BindInfoAllRequest {
	s.Size = &v
	return s
}

func (s *BindInfoAllRequest) SetHeader(v map[string]*string) *BindInfoAllRequest {
	s.Header = v
	return s
}

func (s *BindInfoAllRequest) SetAccessToken(v string) *BindInfoAllRequest {
	s.AccessToken = &v
	return s
}

func (s *BindInfoAllRequest) SetAccountId(v int64) *BindInfoAllRequest {
	s.AccountId = &v
	return s
}

func (s *BindInfoAllRequest) SetCursor(v int64) *BindInfoAllRequest {
	s.Cursor = &v
	return s
}

type BindInfoAllResponse struct {
	Extra *BindInfoAllResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *BindInfoAllResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BindInfoAllResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInfoAllResponse) GoString() string {
	return s.String()
}

func (s *BindInfoAllResponse) SetExtra(v *BindInfoAllResponseExtra) *BindInfoAllResponse {
	s.Extra = v
	return s
}

func (s *BindInfoAllResponse) SetData(v *BindInfoAllResponseData) *BindInfoAllResponse {
	s.Data = v
	return s
}

type BindInfoAllResponseData struct {
	HasMore                     *bool                                                     `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	OpenapiMerchatCraftsmanInfo []*BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem `json:"openapi_merchat_craftsman_info,omitempty" xml:"openapi_merchat_craftsman_info,omitempty" type:"Repeated"`
	Cursor                      *int64                                                    `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	GwErrorCode                 *int32                                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription               *string                                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BindInfoAllResponseData) String() string {
	return tea.Prettify(s)
}

func (s BindInfoAllResponseData) GoString() string {
	return s.String()
}

func (s *BindInfoAllResponseData) SetHasMore(v bool) *BindInfoAllResponseData {
	s.HasMore = &v
	return s
}

func (s *BindInfoAllResponseData) SetOpenapiMerchatCraftsmanInfo(v []*BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) *BindInfoAllResponseData {
	s.OpenapiMerchatCraftsmanInfo = v
	return s
}

func (s *BindInfoAllResponseData) SetCursor(v int64) *BindInfoAllResponseData {
	s.Cursor = &v
	return s
}

func (s *BindInfoAllResponseData) SetGwErrorCode(v int32) *BindInfoAllResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BindInfoAllResponseData) SetGwDescription(v string) *BindInfoAllResponseData {
	s.GwDescription = &v
	return s
}

type BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem struct {
	BindStatus      *int    `json:"bind_status,omitempty" xml:"bind_status,omitempty"`
	PoiAccountName  *string `json:"poi_account_name,omitempty" xml:"poi_account_name,omitempty"`
	CooperationMode *int    `json:"cooperation_mode,omitempty" xml:"cooperation_mode,omitempty"`
	RealName        *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	AwemeShortId    *string `json:"aweme_short_id,omitempty" xml:"aweme_short_id,omitempty"`
	AccountName     *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	PoiId           *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Nickname        *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	CraftsmanUid    *string `json:"craftsman_uid,omitempty" xml:"craftsman_uid,omitempty"`
	BindStartTime   *int64  `json:"bind_start_time,omitempty" xml:"bind_start_time,omitempty"`
	AccountId       *int64  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PaymentStatus   *int    `json:"payment_status,omitempty" xml:"payment_status,omitempty"`
	BindEndTime     *int64  `json:"bind_end_time,omitempty" xml:"bind_end_time,omitempty"`
}

func (s BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) String() string {
	return tea.Prettify(s)
}

func (s BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) GoString() string {
	return s.String()
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindStatus(v int) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindStatus = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetPoiAccountName(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PoiAccountName = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetCooperationMode(v int) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.CooperationMode = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetRealName(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.RealName = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetAwemeShortId(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AwemeShortId = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetAccountName(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AccountName = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetPoiId(v int64) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PoiId = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetNickname(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.Nickname = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetCraftsmanUid(v string) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.CraftsmanUid = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindStartTime(v int64) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindStartTime = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetAccountId(v int64) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AccountId = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetPaymentStatus(v int) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PaymentStatus = &v
	return s
}

func (s *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindEndTime(v int64) *BindInfoAllResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindEndTime = &v
	return s
}

type BindInfoAllResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BindInfoAllResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BindInfoAllResponseExtra) GoString() string {
	return s.String()
}

func (s *BindInfoAllResponseExtra) SetLogid(v string) *BindInfoAllResponseExtra {
	s.Logid = &v
	return s
}

func (s *BindInfoAllResponseExtra) SetNow(v int64) *BindInfoAllResponseExtra {
	s.Now = &v
	return s
}

func (s *BindInfoAllResponseExtra) SetSubDescription(v string) *BindInfoAllResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BindInfoAllResponseExtra) SetSubErrorCode(v int32) *BindInfoAllResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BindInfoAllResponseExtra) SetDescription(v string) *BindInfoAllResponseExtra {
	s.Description = &v
	return s
}

func (s *BindInfoAllResponseExtra) SetErrorCode(v int32) *BindInfoAllResponseExtra {
	s.ErrorCode = &v
	return s
}

type BindInfoSingleRequest struct {
	CraftsmanUid *string            `json:"craftsman_uid,omitempty" xml:"craftsman_uid,omitempty" require:"true"`
	Limit        *int64             `json:"limit,omitempty" xml:"limit,omitempty" require:"true"`
	Offset       *int64             `json:"offset,omitempty" xml:"offset,omitempty" require:"true"`
	AccountId    *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BindInfoSingleRequest) String() string {
	return tea.Prettify(s)
}

func (s BindInfoSingleRequest) GoString() string {
	return s.String()
}

func (s *BindInfoSingleRequest) SetCraftsmanUid(v string) *BindInfoSingleRequest {
	s.CraftsmanUid = &v
	return s
}

func (s *BindInfoSingleRequest) SetLimit(v int64) *BindInfoSingleRequest {
	s.Limit = &v
	return s
}

func (s *BindInfoSingleRequest) SetOffset(v int64) *BindInfoSingleRequest {
	s.Offset = &v
	return s
}

func (s *BindInfoSingleRequest) SetAccountId(v int64) *BindInfoSingleRequest {
	s.AccountId = &v
	return s
}

func (s *BindInfoSingleRequest) SetHeader(v map[string]*string) *BindInfoSingleRequest {
	s.Header = v
	return s
}

func (s *BindInfoSingleRequest) SetAccessToken(v string) *BindInfoSingleRequest {
	s.AccessToken = &v
	return s
}

type BindInfoSingleResponse struct {
	Data  *BindInfoSingleResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *BindInfoSingleResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s BindInfoSingleResponse) String() string {
	return tea.Prettify(s)
}

func (s BindInfoSingleResponse) GoString() string {
	return s.String()
}

func (s *BindInfoSingleResponse) SetData(v *BindInfoSingleResponseData) *BindInfoSingleResponse {
	s.Data = v
	return s
}

func (s *BindInfoSingleResponse) SetExtra(v *BindInfoSingleResponseExtra) *BindInfoSingleResponse {
	s.Extra = v
	return s
}

type BindInfoSingleResponseData struct {
	GwDescription               *string                                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OpenapiMerchatCraftsmanInfo []*BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem `json:"openapi_merchat_craftsman_info,omitempty" xml:"openapi_merchat_craftsman_info,omitempty" type:"Repeated"`
	Total                       *int64                                                       `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	GwErrorCode                 *int32                                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BindInfoSingleResponseData) String() string {
	return tea.Prettify(s)
}

func (s BindInfoSingleResponseData) GoString() string {
	return s.String()
}

func (s *BindInfoSingleResponseData) SetGwDescription(v string) *BindInfoSingleResponseData {
	s.GwDescription = &v
	return s
}

func (s *BindInfoSingleResponseData) SetOpenapiMerchatCraftsmanInfo(v []*BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) *BindInfoSingleResponseData {
	s.OpenapiMerchatCraftsmanInfo = v
	return s
}

func (s *BindInfoSingleResponseData) SetTotal(v int64) *BindInfoSingleResponseData {
	s.Total = &v
	return s
}

func (s *BindInfoSingleResponseData) SetGwErrorCode(v int32) *BindInfoSingleResponseData {
	s.GwErrorCode = &v
	return s
}

type BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem struct {
	CooperationMode *int    `json:"cooperation_mode,omitempty" xml:"cooperation_mode,omitempty"`
	BindStatus      *int    `json:"bind_status,omitempty" xml:"bind_status,omitempty"`
	BindStartTime   *int64  `json:"bind_start_time,omitempty" xml:"bind_start_time,omitempty"`
	CraftsmanUid    *string `json:"craftsman_uid,omitempty" xml:"craftsman_uid,omitempty"`
	RealName        *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	Nickname        *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	BindEndTime     *int64  `json:"bind_end_time,omitempty" xml:"bind_end_time,omitempty"`
	PaymentStatus   *int    `json:"payment_status,omitempty" xml:"payment_status,omitempty"`
	AccountId       *int64  `json:"account_id,omitempty" xml:"account_id,omitempty"`
	PoiId           *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	AwemeShortId    *string `json:"aweme_short_id,omitempty" xml:"aweme_short_id,omitempty"`
	PoiAccountName  *string `json:"poi_account_name,omitempty" xml:"poi_account_name,omitempty"`
	AccountName     *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) String() string {
	return tea.Prettify(s)
}

func (s BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) GoString() string {
	return s.String()
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetCooperationMode(v int) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.CooperationMode = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindStatus(v int) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindStatus = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindStartTime(v int64) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindStartTime = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetCraftsmanUid(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.CraftsmanUid = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetRealName(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.RealName = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetNickname(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.Nickname = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetBindEndTime(v int64) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.BindEndTime = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetPaymentStatus(v int) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PaymentStatus = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetAccountId(v int64) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AccountId = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetPoiId(v int64) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PoiId = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetAwemeShortId(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AwemeShortId = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetPoiAccountName(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.PoiAccountName = &v
	return s
}

func (s *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem) SetAccountName(v string) *BindInfoSingleResponseDataOpenapiMerchatCraftsmanInfoItem {
	s.AccountName = &v
	return s
}

type BindInfoSingleResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BindInfoSingleResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BindInfoSingleResponseExtra) GoString() string {
	return s.String()
}

func (s *BindInfoSingleResponseExtra) SetErrorCode(v int32) *BindInfoSingleResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BindInfoSingleResponseExtra) SetLogid(v string) *BindInfoSingleResponseExtra {
	s.Logid = &v
	return s
}

func (s *BindInfoSingleResponseExtra) SetNow(v int64) *BindInfoSingleResponseExtra {
	s.Now = &v
	return s
}

func (s *BindInfoSingleResponseExtra) SetSubDescription(v string) *BindInfoSingleResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BindInfoSingleResponseExtra) SetSubErrorCode(v int32) *BindInfoSingleResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BindInfoSingleResponseExtra) SetDescription(v string) *BindInfoSingleResponseExtra {
	s.Description = &v
	return s
}

type BindSaveRequest struct {
	AccountId          *int64                          `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	CouponProductId    *int64                          `json:"coupon_product_id,omitempty" xml:"coupon_product_id,omitempty"`
	FulfillmentType    []*int                          `json:"fulfillment_type,omitempty" xml:"fulfillment_type,omitempty" type:"Repeated"`
	GroupList          []*BindSaveRequestGroupListItem `json:"group_list,omitempty" xml:"group_list,omitempty" require:"true" type:"Repeated"`
	Header             map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OutCouponProductId *string                         `json:"out_coupon_product_id,omitempty" xml:"out_coupon_product_id,omitempty"`
}

func (s BindSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequest) GoString() string {
	return s.String()
}

func (s *BindSaveRequest) SetAccountId(v int64) *BindSaveRequest {
	s.AccountId = &v
	return s
}

func (s *BindSaveRequest) SetCouponProductId(v int64) *BindSaveRequest {
	s.CouponProductId = &v
	return s
}

func (s *BindSaveRequest) SetFulfillmentType(v []*int) *BindSaveRequest {
	s.FulfillmentType = v
	return s
}

func (s *BindSaveRequest) SetGroupList(v []*BindSaveRequestGroupListItem) *BindSaveRequest {
	s.GroupList = v
	return s
}

func (s *BindSaveRequest) SetHeader(v map[string]*string) *BindSaveRequest {
	s.Header = v
	return s
}

func (s *BindSaveRequest) SetAccessToken(v string) *BindSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *BindSaveRequest) SetOutCouponProductId(v string) *BindSaveRequest {
	s.OutCouponProductId = &v
	return s
}

type BindSaveRequestGroupListItem struct {
	GroupName   *string                                     `json:"group_name,omitempty" xml:"group_name,omitempty" require:"true"`
	ItemList    []*BindSaveRequestGroupListItemItemListItem `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	OptionCount *int32                                      `json:"option_count,omitempty" xml:"option_count,omitempty"`
	TotalCount  *int32                                      `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s BindSaveRequestGroupListItem) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequestGroupListItem) GoString() string {
	return s.String()
}

func (s *BindSaveRequestGroupListItem) SetGroupName(v string) *BindSaveRequestGroupListItem {
	s.GroupName = &v
	return s
}

func (s *BindSaveRequestGroupListItem) SetItemList(v []*BindSaveRequestGroupListItemItemListItem) *BindSaveRequestGroupListItem {
	s.ItemList = v
	return s
}

func (s *BindSaveRequestGroupListItem) SetOptionCount(v int32) *BindSaveRequestGroupListItem {
	s.OptionCount = &v
	return s
}

func (s *BindSaveRequestGroupListItem) SetTotalCount(v int32) *BindSaveRequestGroupListItem {
	s.TotalCount = &v
	return s
}

type BindSaveRequestGroupListItemItemListItem struct {
	Price       *int64                                                  `json:"price,omitempty" xml:"price,omitempty"`
	SpuRelation *BindSaveRequestGroupListItemItemListItemSpuRelation    `json:"spu_relation,omitempty" xml:"spu_relation,omitempty"`
	Unit        *string                                                 `json:"unit,omitempty" xml:"unit,omitempty"`
	AttrList    []*BindSaveRequestGroupListItemItemListItemAttrListItem `json:"attr_list,omitempty" xml:"attr_list,omitempty" type:"Repeated"`
	Count       *int32                                                  `json:"count,omitempty" xml:"count,omitempty"`
}

func (s BindSaveRequestGroupListItemItemListItem) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequestGroupListItemItemListItem) GoString() string {
	return s.String()
}

func (s *BindSaveRequestGroupListItemItemListItem) SetPrice(v int64) *BindSaveRequestGroupListItemItemListItem {
	s.Price = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItem) SetSpuRelation(v *BindSaveRequestGroupListItemItemListItemSpuRelation) *BindSaveRequestGroupListItemItemListItem {
	s.SpuRelation = v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItem) SetUnit(v string) *BindSaveRequestGroupListItemItemListItem {
	s.Unit = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItem) SetAttrList(v []*BindSaveRequestGroupListItemItemListItemAttrListItem) *BindSaveRequestGroupListItemItemListItem {
	s.AttrList = v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItem) SetCount(v int32) *BindSaveRequestGroupListItemItemListItem {
	s.Count = &v
	return s
}

type BindSaveRequestGroupListItemItemListItemAttrListItem struct {
	AttrValue *string `json:"attr_value,omitempty" xml:"attr_value,omitempty" require:"true"`
	AttrKey   *string `json:"attr_key,omitempty" xml:"attr_key,omitempty" require:"true"`
}

func (s BindSaveRequestGroupListItemItemListItemAttrListItem) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequestGroupListItemItemListItemAttrListItem) GoString() string {
	return s.String()
}

func (s *BindSaveRequestGroupListItemItemListItemAttrListItem) SetAttrValue(v string) *BindSaveRequestGroupListItemItemListItemAttrListItem {
	s.AttrValue = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemAttrListItem) SetAttrKey(v string) *BindSaveRequestGroupListItemItemListItemAttrListItem {
	s.AttrKey = &v
	return s
}

type BindSaveRequestGroupListItemItemListItemSpuRelation struct {
	SpuPriceLow  *int64                                                             `json:"spu_price_low,omitempty" xml:"spu_price_low,omitempty"`
	CSpuId       *int64                                                             `json:"c_spu_id,omitempty" xml:"c_spu_id,omitempty"`
	OutSpuId     *string                                                            `json:"out_spu_id,omitempty" xml:"out_spu_id,omitempty"`
	SpecList     []*BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem `json:"spec_list,omitempty" xml:"spec_list,omitempty" type:"Repeated"`
	SpuId        *int64                                                             `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	SpuPriceHigh *int64                                                             `json:"spu_price_high,omitempty" xml:"spu_price_high,omitempty"`
}

func (s BindSaveRequestGroupListItemItemListItemSpuRelation) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequestGroupListItemItemListItemSpuRelation) GoString() string {
	return s.String()
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetSpuPriceLow(v int64) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.SpuPriceLow = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetCSpuId(v int64) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.CSpuId = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetOutSpuId(v string) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.OutSpuId = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetSpecList(v []*BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.SpecList = v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetSpuId(v int64) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.SpuId = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelation) SetSpuPriceHigh(v int64) *BindSaveRequestGroupListItemItemListItemSpuRelation {
	s.SpuPriceHigh = &v
	return s
}

type BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem struct {
	GroupCode *string `json:"group_code,omitempty" xml:"group_code,omitempty"`
	ItemKey   *string `json:"item_key,omitempty" xml:"item_key,omitempty"`
}

func (s BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem) String() string {
	return tea.Prettify(s)
}

func (s BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem) GoString() string {
	return s.String()
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem) SetGroupCode(v string) *BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem {
	s.GroupCode = &v
	return s
}

func (s *BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem) SetItemKey(v string) *BindSaveRequestGroupListItemItemListItemSpuRelationSpecListItem {
	s.ItemKey = &v
	return s
}

type BindSaveResponse struct {
	Extra    *BindSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *BindSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *BindSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BindSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s BindSaveResponse) GoString() string {
	return s.String()
}

func (s *BindSaveResponse) SetExtra(v *BindSaveResponseExtra) *BindSaveResponse {
	s.Extra = v
	return s
}

func (s *BindSaveResponse) SetBaseResp(v *BindSaveResponseBaseResp) *BindSaveResponse {
	s.BaseResp = v
	return s
}

func (s *BindSaveResponse) SetData(v *BindSaveResponseData) *BindSaveResponse {
	s.Data = v
	return s
}

type BindSaveResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s BindSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s BindSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *BindSaveResponseBaseResp) SetExtra(v map[string]*string) *BindSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *BindSaveResponseBaseResp) SetStatusCode(v int32) *BindSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *BindSaveResponseBaseResp) SetStatusMessage(v string) *BindSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type BindSaveResponseData struct {
	ErrorCode       *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	CouponProductId *int64  `json:"coupon_product_id,omitempty" xml:"coupon_product_id,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BindSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s BindSaveResponseData) GoString() string {
	return s.String()
}

func (s *BindSaveResponseData) SetErrorCode(v int32) *BindSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *BindSaveResponseData) SetCouponProductId(v int64) *BindSaveResponseData {
	s.CouponProductId = &v
	return s
}

func (s *BindSaveResponseData) SetDescription(v string) *BindSaveResponseData {
	s.Description = &v
	return s
}

type BindSaveResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s BindSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BindSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *BindSaveResponseExtra) SetSubErrorCode(v int32) *BindSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BindSaveResponseExtra) SetDescription(v string) *BindSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *BindSaveResponseExtra) SetErrorCode(v int32) *BindSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BindSaveResponseExtra) SetLogid(v string) *BindSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *BindSaveResponseExtra) SetNow(v int64) *BindSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *BindSaveResponseExtra) SetSubDescription(v string) *BindSaveResponseExtra {
	s.SubDescription = &v
	return s
}

type BookUserCancelBookRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BookId       *string            `json:"book_id,omitempty" xml:"book_id,omitempty" require:"true"`
	CancelReason []*string          `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty" type:"Repeated"`
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s BookUserCancelBookRequest) String() string {
	return tea.Prettify(s)
}

func (s BookUserCancelBookRequest) GoString() string {
	return s.String()
}

func (s *BookUserCancelBookRequest) SetHeader(v map[string]*string) *BookUserCancelBookRequest {
	s.Header = v
	return s
}

func (s *BookUserCancelBookRequest) SetAccessToken(v string) *BookUserCancelBookRequest {
	s.AccessToken = &v
	return s
}

func (s *BookUserCancelBookRequest) SetBookId(v string) *BookUserCancelBookRequest {
	s.BookId = &v
	return s
}

func (s *BookUserCancelBookRequest) SetCancelReason(v []*string) *BookUserCancelBookRequest {
	s.CancelReason = v
	return s
}

func (s *BookUserCancelBookRequest) SetOpenId(v string) *BookUserCancelBookRequest {
	s.OpenId = &v
	return s
}

type BookUserCancelBookResponse struct {
	Extra *BookUserCancelBookResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *BookUserCancelBookResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s BookUserCancelBookResponse) String() string {
	return tea.Prettify(s)
}

func (s BookUserCancelBookResponse) GoString() string {
	return s.String()
}

func (s *BookUserCancelBookResponse) SetExtra(v *BookUserCancelBookResponseExtra) *BookUserCancelBookResponse {
	s.Extra = v
	return s
}

func (s *BookUserCancelBookResponse) SetData(v *BookUserCancelBookResponseData) *BookUserCancelBookResponse {
	s.Data = v
	return s
}

type BookUserCancelBookResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BookUserCancelBookResponseData) String() string {
	return tea.Prettify(s)
}

func (s BookUserCancelBookResponseData) GoString() string {
	return s.String()
}

func (s *BookUserCancelBookResponseData) SetGwDescription(v string) *BookUserCancelBookResponseData {
	s.GwDescription = &v
	return s
}

func (s *BookUserCancelBookResponseData) SetGwErrorCode(v int32) *BookUserCancelBookResponseData {
	s.GwErrorCode = &v
	return s
}

type BookUserCancelBookResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s BookUserCancelBookResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BookUserCancelBookResponseExtra) GoString() string {
	return s.String()
}

func (s *BookUserCancelBookResponseExtra) SetNow(v int64) *BookUserCancelBookResponseExtra {
	s.Now = &v
	return s
}

func (s *BookUserCancelBookResponseExtra) SetSubDescription(v string) *BookUserCancelBookResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BookUserCancelBookResponseExtra) SetSubErrorCode(v int32) *BookUserCancelBookResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BookUserCancelBookResponseExtra) SetDescription(v string) *BookUserCancelBookResponseExtra {
	s.Description = &v
	return s
}

func (s *BookUserCancelBookResponseExtra) SetErrorCode(v int32) *BookUserCancelBookResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BookUserCancelBookResponseExtra) SetLogid(v string) *BookUserCancelBookResponseExtra {
	s.Logid = &v
	return s
}

type BookingAuditNotifyRequest struct {
	OrderOutId          *string            `json:"order_out_id,omitempty" xml:"order_out_id,omitempty" require:"true"`
	AccommodationStatus *int               `json:"accommodation_status,omitempty" xml:"accommodation_status,omitempty" require:"true"`
	OrderId             *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header              map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken         *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BookingAuditNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s BookingAuditNotifyRequest) GoString() string {
	return s.String()
}

func (s *BookingAuditNotifyRequest) SetOrderOutId(v string) *BookingAuditNotifyRequest {
	s.OrderOutId = &v
	return s
}

func (s *BookingAuditNotifyRequest) SetAccommodationStatus(v int) *BookingAuditNotifyRequest {
	s.AccommodationStatus = &v
	return s
}

func (s *BookingAuditNotifyRequest) SetOrderId(v string) *BookingAuditNotifyRequest {
	s.OrderId = &v
	return s
}

func (s *BookingAuditNotifyRequest) SetHeader(v map[string]*string) *BookingAuditNotifyRequest {
	s.Header = v
	return s
}

func (s *BookingAuditNotifyRequest) SetAccessToken(v string) *BookingAuditNotifyRequest {
	s.AccessToken = &v
	return s
}

type BookingAuditNotifyResponse struct {
	Extra *BookingAuditNotifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *BookingAuditNotifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s BookingAuditNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s BookingAuditNotifyResponse) GoString() string {
	return s.String()
}

func (s *BookingAuditNotifyResponse) SetExtra(v *BookingAuditNotifyResponseExtra) *BookingAuditNotifyResponse {
	s.Extra = v
	return s
}

func (s *BookingAuditNotifyResponse) SetData(v *BookingAuditNotifyResponseData) *BookingAuditNotifyResponse {
	s.Data = v
	return s
}

type BookingAuditNotifyResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BookingAuditNotifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s BookingAuditNotifyResponseData) GoString() string {
	return s.String()
}

func (s *BookingAuditNotifyResponseData) SetGwErrorCode(v int32) *BookingAuditNotifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BookingAuditNotifyResponseData) SetGwDescription(v string) *BookingAuditNotifyResponseData {
	s.GwDescription = &v
	return s
}

type BookingAuditNotifyResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s BookingAuditNotifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BookingAuditNotifyResponseExtra) GoString() string {
	return s.String()
}

func (s *BookingAuditNotifyResponseExtra) SetDescription(v string) *BookingAuditNotifyResponseExtra {
	s.Description = &v
	return s
}

func (s *BookingAuditNotifyResponseExtra) SetErrorCode(v int32) *BookingAuditNotifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BookingAuditNotifyResponseExtra) SetLogid(v string) *BookingAuditNotifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *BookingAuditNotifyResponseExtra) SetNow(v int64) *BookingAuditNotifyResponseExtra {
	s.Now = &v
	return s
}

func (s *BookingAuditNotifyResponseExtra) SetSubDescription(v string) *BookingAuditNotifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BookingAuditNotifyResponseExtra) SetSubErrorCode(v int32) *BookingAuditNotifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

type BookingConfigRequest struct {
	Config      *BookingConfigRequestConfig `json:"config,omitempty" xml:"config,omitempty" require:"true"`
	Header      map[string]*string          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiIds      []*string                   `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" require:"true" type:"Repeated"`
	AccountId   *string                     `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s BookingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s BookingConfigRequest) GoString() string {
	return s.String()
}

func (s *BookingConfigRequest) SetConfig(v *BookingConfigRequestConfig) *BookingConfigRequest {
	s.Config = v
	return s
}

func (s *BookingConfigRequest) SetHeader(v map[string]*string) *BookingConfigRequest {
	s.Header = v
	return s
}

func (s *BookingConfigRequest) SetAccessToken(v string) *BookingConfigRequest {
	s.AccessToken = &v
	return s
}

func (s *BookingConfigRequest) SetPoiIds(v []*string) *BookingConfigRequest {
	s.PoiIds = v
	return s
}

func (s *BookingConfigRequest) SetAccountId(v string) *BookingConfigRequest {
	s.AccountId = &v
	return s
}

type BookingConfigRequestConfig struct {
	BookingTip          *string   `json:"booking_tip,omitempty" xml:"booking_tip,omitempty"`
	EarliestBookingTime *int32    `json:"earliest_booking_time,omitempty" xml:"earliest_booking_time,omitempty" require:"true"`
	LatestBookingTime   *int32    `json:"latest_booking_time,omitempty" xml:"latest_booking_time,omitempty" require:"true"`
	OpenStatus          *int32    `json:"open_status,omitempty" xml:"open_status,omitempty" require:"true"`
	RemarkTags          []*string `json:"remark_tags,omitempty" xml:"remark_tags,omitempty" type:"Repeated"`
}

func (s BookingConfigRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s BookingConfigRequestConfig) GoString() string {
	return s.String()
}

func (s *BookingConfigRequestConfig) SetBookingTip(v string) *BookingConfigRequestConfig {
	s.BookingTip = &v
	return s
}

func (s *BookingConfigRequestConfig) SetEarliestBookingTime(v int32) *BookingConfigRequestConfig {
	s.EarliestBookingTime = &v
	return s
}

func (s *BookingConfigRequestConfig) SetLatestBookingTime(v int32) *BookingConfigRequestConfig {
	s.LatestBookingTime = &v
	return s
}

func (s *BookingConfigRequestConfig) SetOpenStatus(v int32) *BookingConfigRequestConfig {
	s.OpenStatus = &v
	return s
}

func (s *BookingConfigRequestConfig) SetRemarkTags(v []*string) *BookingConfigRequestConfig {
	s.RemarkTags = v
	return s
}

type BookingConfigResponse struct {
	Data  *BookingConfigResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *BookingConfigResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s BookingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s BookingConfigResponse) GoString() string {
	return s.String()
}

func (s *BookingConfigResponse) SetData(v *BookingConfigResponseData) *BookingConfigResponse {
	s.Data = v
	return s
}

func (s *BookingConfigResponse) SetExtra(v *BookingConfigResponseExtra) *BookingConfigResponse {
	s.Extra = v
	return s
}

type BookingConfigResponseData struct {
	TaskId      *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s BookingConfigResponseData) String() string {
	return tea.Prettify(s)
}

func (s BookingConfigResponseData) GoString() string {
	return s.String()
}

func (s *BookingConfigResponseData) SetTaskId(v string) *BookingConfigResponseData {
	s.TaskId = &v
	return s
}

func (s *BookingConfigResponseData) SetDescription(v string) *BookingConfigResponseData {
	s.Description = &v
	return s
}

func (s *BookingConfigResponseData) SetErrorCode(v int32) *BookingConfigResponseData {
	s.ErrorCode = &v
	return s
}

type BookingConfigResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s BookingConfigResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BookingConfigResponseExtra) GoString() string {
	return s.String()
}

func (s *BookingConfigResponseExtra) SetSubDescription(v string) *BookingConfigResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BookingConfigResponseExtra) SetSubErrorCode(v int32) *BookingConfigResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *BookingConfigResponseExtra) SetDescription(v string) *BookingConfigResponseExtra {
	s.Description = &v
	return s
}

func (s *BookingConfigResponseExtra) SetErrorCode(v int32) *BookingConfigResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BookingConfigResponseExtra) SetLogid(v string) *BookingConfigResponseExtra {
	s.Logid = &v
	return s
}

func (s *BookingConfigResponseExtra) SetNow(v int64) *BookingConfigResponseExtra {
	s.Now = &v
	return s
}

type BookingOrderFulfillmentRequest struct {
	AccountId         *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ExtOrderId        *string            `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty" require:"true"`
	FulfillmentStatus *int               `json:"fulfillment_status,omitempty" xml:"fulfillment_status,omitempty" require:"true"`
	OrderId           *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BookingOrderFulfillmentRequest) String() string {
	return tea.Prettify(s)
}

func (s BookingOrderFulfillmentRequest) GoString() string {
	return s.String()
}

func (s *BookingOrderFulfillmentRequest) SetAccountId(v string) *BookingOrderFulfillmentRequest {
	s.AccountId = &v
	return s
}

func (s *BookingOrderFulfillmentRequest) SetExtOrderId(v string) *BookingOrderFulfillmentRequest {
	s.ExtOrderId = &v
	return s
}

func (s *BookingOrderFulfillmentRequest) SetFulfillmentStatus(v int) *BookingOrderFulfillmentRequest {
	s.FulfillmentStatus = &v
	return s
}

func (s *BookingOrderFulfillmentRequest) SetOrderId(v string) *BookingOrderFulfillmentRequest {
	s.OrderId = &v
	return s
}

func (s *BookingOrderFulfillmentRequest) SetHeader(v map[string]*string) *BookingOrderFulfillmentRequest {
	s.Header = v
	return s
}

func (s *BookingOrderFulfillmentRequest) SetAccessToken(v string) *BookingOrderFulfillmentRequest {
	s.AccessToken = &v
	return s
}

type BookingOrderFulfillmentResponse struct {
	Extra *BookingOrderFulfillmentResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *BookingOrderFulfillmentResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BookingOrderFulfillmentResponse) String() string {
	return tea.Prettify(s)
}

func (s BookingOrderFulfillmentResponse) GoString() string {
	return s.String()
}

func (s *BookingOrderFulfillmentResponse) SetExtra(v *BookingOrderFulfillmentResponseExtra) *BookingOrderFulfillmentResponse {
	s.Extra = v
	return s
}

func (s *BookingOrderFulfillmentResponse) SetData(v *BookingOrderFulfillmentResponseData) *BookingOrderFulfillmentResponse {
	s.Data = v
	return s
}

type BookingOrderFulfillmentResponseData struct {
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BookingOrderFulfillmentResponseData) String() string {
	return tea.Prettify(s)
}

func (s BookingOrderFulfillmentResponseData) GoString() string {
	return s.String()
}

func (s *BookingOrderFulfillmentResponseData) SetErrorCode(v int32) *BookingOrderFulfillmentResponseData {
	s.ErrorCode = &v
	return s
}

func (s *BookingOrderFulfillmentResponseData) SetDescription(v string) *BookingOrderFulfillmentResponseData {
	s.Description = &v
	return s
}

type BookingOrderFulfillmentResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s BookingOrderFulfillmentResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BookingOrderFulfillmentResponseExtra) GoString() string {
	return s.String()
}

func (s *BookingOrderFulfillmentResponseExtra) SetDescription(v string) *BookingOrderFulfillmentResponseExtra {
	s.Description = &v
	return s
}

func (s *BookingOrderFulfillmentResponseExtra) SetErrorCode(v int32) *BookingOrderFulfillmentResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BookingOrderFulfillmentResponseExtra) SetLogid(v string) *BookingOrderFulfillmentResponseExtra {
	s.Logid = &v
	return s
}

func (s *BookingOrderFulfillmentResponseExtra) SetNow(v int64) *BookingOrderFulfillmentResponseExtra {
	s.Now = &v
	return s
}

func (s *BookingOrderFulfillmentResponseExtra) SetSubDescription(v string) *BookingOrderFulfillmentResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BookingOrderFulfillmentResponseExtra) SetSubErrorCode(v int32) *BookingOrderFulfillmentResponseExtra {
	s.SubErrorCode = &v
	return s
}

type BuyMerchantConfirmOrderRequest struct {
	OrderId     *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s BuyMerchantConfirmOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s BuyMerchantConfirmOrderRequest) GoString() string {
	return s.String()
}

func (s *BuyMerchantConfirmOrderRequest) SetOrderId(v string) *BuyMerchantConfirmOrderRequest {
	s.OrderId = &v
	return s
}

func (s *BuyMerchantConfirmOrderRequest) SetHeader(v map[string]*string) *BuyMerchantConfirmOrderRequest {
	s.Header = v
	return s
}

func (s *BuyMerchantConfirmOrderRequest) SetAccessToken(v string) *BuyMerchantConfirmOrderRequest {
	s.AccessToken = &v
	return s
}

type BuyMerchantConfirmOrderResponse struct {
	Data  *BuyMerchantConfirmOrderResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *BuyMerchantConfirmOrderResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s BuyMerchantConfirmOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s BuyMerchantConfirmOrderResponse) GoString() string {
	return s.String()
}

func (s *BuyMerchantConfirmOrderResponse) SetData(v *BuyMerchantConfirmOrderResponseData) *BuyMerchantConfirmOrderResponse {
	s.Data = v
	return s
}

func (s *BuyMerchantConfirmOrderResponse) SetExtra(v *BuyMerchantConfirmOrderResponseExtra) *BuyMerchantConfirmOrderResponse {
	s.Extra = v
	return s
}

type BuyMerchantConfirmOrderResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s BuyMerchantConfirmOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s BuyMerchantConfirmOrderResponseData) GoString() string {
	return s.String()
}

func (s *BuyMerchantConfirmOrderResponseData) SetGwErrorCode(v int32) *BuyMerchantConfirmOrderResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseData) SetGwDescription(v string) *BuyMerchantConfirmOrderResponseData {
	s.GwDescription = &v
	return s
}

type BuyMerchantConfirmOrderResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s BuyMerchantConfirmOrderResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BuyMerchantConfirmOrderResponseExtra) GoString() string {
	return s.String()
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetDescription(v string) *BuyMerchantConfirmOrderResponseExtra {
	s.Description = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetErrorCode(v int32) *BuyMerchantConfirmOrderResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetLogid(v string) *BuyMerchantConfirmOrderResponseExtra {
	s.Logid = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetNow(v int64) *BuyMerchantConfirmOrderResponseExtra {
	s.Now = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetSubDescription(v string) *BuyMerchantConfirmOrderResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BuyMerchantConfirmOrderResponseExtra) SetSubErrorCode(v int32) *BuyMerchantConfirmOrderResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CalendarAriRequest struct {
	AccountId            *string                                       `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ProductId            *string                                       `json:"product_id,omitempty" xml:"product_id,omitempty"`
	BookingCalendarStock []*CalendarAriRequestBookingCalendarStockItem `json:"booking_calendar_stock,omitempty" xml:"booking_calendar_stock,omitempty" type:"Repeated"`
	OutId                *string                                       `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Header               map[string]*string                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CalendarAriRequest) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriRequest) GoString() string {
	return s.String()
}

func (s *CalendarAriRequest) SetAccountId(v string) *CalendarAriRequest {
	s.AccountId = &v
	return s
}

func (s *CalendarAriRequest) SetProductId(v string) *CalendarAriRequest {
	s.ProductId = &v
	return s
}

func (s *CalendarAriRequest) SetBookingCalendarStock(v []*CalendarAriRequestBookingCalendarStockItem) *CalendarAriRequest {
	s.BookingCalendarStock = v
	return s
}

func (s *CalendarAriRequest) SetOutId(v string) *CalendarAriRequest {
	s.OutId = &v
	return s
}

func (s *CalendarAriRequest) SetHeader(v map[string]*string) *CalendarAriRequest {
	s.Header = v
	return s
}

func (s *CalendarAriRequest) SetAccessToken(v string) *CalendarAriRequest {
	s.AccessToken = &v
	return s
}

type CalendarAriRequestBookingCalendarStockItem struct {
	CalendarList []*CalendarAriRequestBookingCalendarStockItemCalendarListItem `json:"calendar_list,omitempty" xml:"calendar_list,omitempty" require:"true" type:"Repeated"`
}

func (s CalendarAriRequestBookingCalendarStockItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriRequestBookingCalendarStockItem) GoString() string {
	return s.String()
}

func (s *CalendarAriRequestBookingCalendarStockItem) SetCalendarList(v []*CalendarAriRequestBookingCalendarStockItemCalendarListItem) *CalendarAriRequestBookingCalendarStockItem {
	s.CalendarList = v
	return s
}

type CalendarAriRequestBookingCalendarStockItemCalendarListItem struct {
	AvailableQty   *int64                                                                   `json:"available_qty,omitempty" xml:"available_qty,omitempty"`
	CalendarStatus *int32                                                                   `json:"calendar_status,omitempty" xml:"calendar_status,omitempty" require:"true"`
	CalendarValue  *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue `json:"calendar_value,omitempty" xml:"calendar_value,omitempty" require:"true"`
	ActualAmount   *int64                                                                   `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s CalendarAriRequestBookingCalendarStockItemCalendarListItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriRequestBookingCalendarStockItemCalendarListItem) GoString() string {
	return s.String()
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItem) SetAvailableQty(v int64) *CalendarAriRequestBookingCalendarStockItemCalendarListItem {
	s.AvailableQty = &v
	return s
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItem) SetCalendarStatus(v int32) *CalendarAriRequestBookingCalendarStockItemCalendarListItem {
	s.CalendarStatus = &v
	return s
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItem) SetCalendarValue(v *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) *CalendarAriRequestBookingCalendarStockItemCalendarListItem {
	s.CalendarValue = v
	return s
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItem) SetActualAmount(v int64) *CalendarAriRequestBookingCalendarStockItemCalendarListItem {
	s.ActualAmount = &v
	return s
}

type CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue struct {
	EndDate              *string  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	StartDate            *string  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	AvailableWeekdayList []*int32 `json:"available_weekday_list,omitempty" xml:"available_weekday_list,omitempty" type:"Repeated"`
}

func (s CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) GoString() string {
	return s.String()
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) SetEndDate(v string) *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue {
	s.EndDate = &v
	return s
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) SetStartDate(v string) *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue {
	s.StartDate = &v
	return s
}

func (s *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue) SetAvailableWeekdayList(v []*int32) *CalendarAriRequestBookingCalendarStockItemCalendarListItemCalendarValue {
	s.AvailableWeekdayList = v
	return s
}

type CalendarAriResponse struct {
	Data  *CalendarAriResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CalendarAriResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CalendarAriResponse) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriResponse) GoString() string {
	return s.String()
}

func (s *CalendarAriResponse) SetData(v *CalendarAriResponseData) *CalendarAriResponse {
	s.Data = v
	return s
}

func (s *CalendarAriResponse) SetExtra(v *CalendarAriResponseExtra) *CalendarAriResponse {
	s.Extra = v
	return s
}

type CalendarAriResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CalendarAriResponseData) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriResponseData) GoString() string {
	return s.String()
}

func (s *CalendarAriResponseData) SetDescription(v string) *CalendarAriResponseData {
	s.Description = &v
	return s
}

func (s *CalendarAriResponseData) SetErrorCode(v int32) *CalendarAriResponseData {
	s.ErrorCode = &v
	return s
}

type CalendarAriResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CalendarAriResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CalendarAriResponseExtra) GoString() string {
	return s.String()
}

func (s *CalendarAriResponseExtra) SetLogid(v string) *CalendarAriResponseExtra {
	s.Logid = &v
	return s
}

func (s *CalendarAriResponseExtra) SetNow(v int64) *CalendarAriResponseExtra {
	s.Now = &v
	return s
}

func (s *CalendarAriResponseExtra) SetSubDescription(v string) *CalendarAriResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CalendarAriResponseExtra) SetSubErrorCode(v int32) *CalendarAriResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CalendarAriResponseExtra) SetDescription(v string) *CalendarAriResponseExtra {
	s.Description = &v
	return s
}

func (s *CalendarAriResponseExtra) SetErrorCode(v int32) *CalendarAriResponseExtra {
	s.ErrorCode = &v
	return s
}

type CalendarOrderConfirmRequest struct {
	AccountId   *string                                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OrderId     *string                                 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	ConfirmInfo *CalendarOrderConfirmRequestConfirmInfo `json:"confirm_info,omitempty" xml:"confirm_info,omitempty"`
	Header      map[string]*string                      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CalendarOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequest) SetAccountId(v string) *CalendarOrderConfirmRequest {
	s.AccountId = &v
	return s
}

func (s *CalendarOrderConfirmRequest) SetOrderId(v string) *CalendarOrderConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *CalendarOrderConfirmRequest) SetConfirmInfo(v *CalendarOrderConfirmRequestConfirmInfo) *CalendarOrderConfirmRequest {
	s.ConfirmInfo = v
	return s
}

func (s *CalendarOrderConfirmRequest) SetHeader(v map[string]*string) *CalendarOrderConfirmRequest {
	s.Header = v
	return s
}

func (s *CalendarOrderConfirmRequest) SetAccessToken(v string) *CalendarOrderConfirmRequest {
	s.AccessToken = &v
	return s
}

type CalendarOrderConfirmRequestConfirmInfo struct {
	RejectCode    *int32                                                `json:"reject_code,omitempty" xml:"reject_code,omitempty"`
	AttentionInfo *string                                               `json:"attention_info,omitempty" xml:"attention_info,omitempty"`
	BookInfo      []*CalendarOrderConfirmRequestConfirmInfoBookInfoItem `json:"book_info,omitempty" xml:"book_info,omitempty" type:"Repeated"`
	ConfirmResult *int32                                                `json:"confirm_result,omitempty" xml:"confirm_result,omitempty"`
}

func (s CalendarOrderConfirmRequestConfirmInfo) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequestConfirmInfo) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequestConfirmInfo) SetRejectCode(v int32) *CalendarOrderConfirmRequestConfirmInfo {
	s.RejectCode = &v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfo) SetAttentionInfo(v string) *CalendarOrderConfirmRequestConfirmInfo {
	s.AttentionInfo = &v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfo) SetBookInfo(v []*CalendarOrderConfirmRequestConfirmInfoBookInfoItem) *CalendarOrderConfirmRequestConfirmInfo {
	s.BookInfo = v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfo) SetConfirmResult(v int32) *CalendarOrderConfirmRequestConfirmInfo {
	s.ConfirmResult = &v
	return s
}

type CalendarOrderConfirmRequestConfirmInfoBookInfoItem struct {
	HotelInfo     *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo `json:"hotel_info,omitempty" xml:"hotel_info,omitempty"`
	AttentionInfo *string                                                      `json:"attention_info,omitempty" xml:"attention_info,omitempty"`
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItem) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItem) SetHotelInfo(v *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) *CalendarOrderConfirmRequestConfirmInfoBookInfoItem {
	s.HotelInfo = v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItem) SetAttentionInfo(v string) *CalendarOrderConfirmRequestConfirmInfoBookInfoItem {
	s.AttentionInfo = &v
	return s
}

type CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo struct {
	RoomItems      []*CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem `json:"room_items,omitempty" xml:"room_items,omitempty" type:"Repeated"`
	HotelConfirmNo *string                                                                     `json:"hotel_confirm_no,omitempty" xml:"hotel_confirm_no,omitempty"`
	PoiInfo        []*CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem   `json:"poi_info,omitempty" xml:"poi_info,omitempty" type:"Repeated"`
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) SetRoomItems(v []*CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo {
	s.RoomItems = v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) SetHotelConfirmNo(v string) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo {
	s.HotelConfirmNo = &v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo) SetPoiInfo(v []*CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfo {
	s.PoiInfo = v
	return s
}

type CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem) SetPoiId(v string) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem {
	s.PoiId = &v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem) SetPoiName(v string) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoPoiInfoItem {
	s.PoiName = &v
	return s
}

type CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem struct {
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
	Meals    *int32  `json:"meals,omitempty" xml:"meals,omitempty"`
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem) SetRoomType(v string) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem {
	s.RoomType = &v
	return s
}

func (s *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem) SetMeals(v int32) *CalendarOrderConfirmRequestConfirmInfoBookInfoItemHotelInfoRoomItemsItem {
	s.Meals = &v
	return s
}

type CalendarOrderConfirmResponse struct {
	Extra *CalendarOrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CalendarOrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CalendarOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmResponse) SetExtra(v *CalendarOrderConfirmResponseExtra) *CalendarOrderConfirmResponse {
	s.Extra = v
	return s
}

func (s *CalendarOrderConfirmResponse) SetData(v *CalendarOrderConfirmResponseData) *CalendarOrderConfirmResponse {
	s.Data = v
	return s
}

type CalendarOrderConfirmResponseData struct {
	ErrorCode   *string `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CalendarOrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmResponseData) SetErrorCode(v string) *CalendarOrderConfirmResponseData {
	s.ErrorCode = &v
	return s
}

func (s *CalendarOrderConfirmResponseData) SetDescription(v string) *CalendarOrderConfirmResponseData {
	s.Description = &v
	return s
}

type CalendarOrderConfirmResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s CalendarOrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CalendarOrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *CalendarOrderConfirmResponseExtra) SetSubErrorCode(v int32) *CalendarOrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CalendarOrderConfirmResponseExtra) SetDescription(v string) *CalendarOrderConfirmResponseExtra {
	s.Description = &v
	return s
}

func (s *CalendarOrderConfirmResponseExtra) SetErrorCode(v int32) *CalendarOrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CalendarOrderConfirmResponseExtra) SetLogid(v string) *CalendarOrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *CalendarOrderConfirmResponseExtra) SetNow(v int64) *CalendarOrderConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *CalendarOrderConfirmResponseExtra) SetSubDescription(v string) *CalendarOrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

type CalendarProductListCalendarAriRequest struct {
	StartDate   *string            `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate     *string            `json:"end_date,omitempty" xml:"end_date,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	OutId       *string            `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId   *string            `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s CalendarProductListCalendarAriRequest) String() string {
	return tea.Prettify(s)
}

func (s CalendarProductListCalendarAriRequest) GoString() string {
	return s.String()
}

func (s *CalendarProductListCalendarAriRequest) SetStartDate(v string) *CalendarProductListCalendarAriRequest {
	s.StartDate = &v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetEndDate(v string) *CalendarProductListCalendarAriRequest {
	s.EndDate = &v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetHeader(v map[string]*string) *CalendarProductListCalendarAriRequest {
	s.Header = v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetAccessToken(v string) *CalendarProductListCalendarAriRequest {
	s.AccessToken = &v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetAccountId(v string) *CalendarProductListCalendarAriRequest {
	s.AccountId = &v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetOutId(v string) *CalendarProductListCalendarAriRequest {
	s.OutId = &v
	return s
}

func (s *CalendarProductListCalendarAriRequest) SetProductId(v string) *CalendarProductListCalendarAriRequest {
	s.ProductId = &v
	return s
}

type CalendarProductListCalendarAriResponse struct {
	Data  *CalendarProductListCalendarAriResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CalendarProductListCalendarAriResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CalendarProductListCalendarAriResponse) String() string {
	return tea.Prettify(s)
}

func (s CalendarProductListCalendarAriResponse) GoString() string {
	return s.String()
}

func (s *CalendarProductListCalendarAriResponse) SetData(v *CalendarProductListCalendarAriResponseData) *CalendarProductListCalendarAriResponse {
	s.Data = v
	return s
}

func (s *CalendarProductListCalendarAriResponse) SetExtra(v *CalendarProductListCalendarAriResponseExtra) *CalendarProductListCalendarAriResponse {
	s.Extra = v
	return s
}

type CalendarProductListCalendarAriResponseData struct {
	CalendarAriList []*CalendarProductListCalendarAriResponseDataCalendarAriListItem `json:"calendar_ari_list,omitempty" xml:"calendar_ari_list,omitempty" type:"Repeated"`
	ProductId       *string                                                          `json:"product_id,omitempty" xml:"product_id,omitempty"`
	UpdateTime      *int64                                                           `json:"update_time,omitempty" xml:"update_time,omitempty"`
	CreateTime      *int64                                                           `json:"create_time,omitempty" xml:"create_time,omitempty"`
	GwDescription   *string                                                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	CategoryId      *int64                                                           `json:"category_id,omitempty" xml:"category_id,omitempty"`
	CategoryName    *string                                                          `json:"category_name,omitempty" xml:"category_name,omitempty"`
	GwErrorCode     *int32                                                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	MarketingAmount *int64                                                           `json:"marketing_amount,omitempty" xml:"marketing_amount,omitempty"`
	ProductName     *string                                                          `json:"product_name,omitempty" xml:"product_name,omitempty"`
	OutId           *string                                                          `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductStatus   *int64                                                           `json:"product_status,omitempty" xml:"product_status,omitempty"`
}

func (s CalendarProductListCalendarAriResponseData) String() string {
	return tea.Prettify(s)
}

func (s CalendarProductListCalendarAriResponseData) GoString() string {
	return s.String()
}

func (s *CalendarProductListCalendarAriResponseData) SetCalendarAriList(v []*CalendarProductListCalendarAriResponseDataCalendarAriListItem) *CalendarProductListCalendarAriResponseData {
	s.CalendarAriList = v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetProductId(v string) *CalendarProductListCalendarAriResponseData {
	s.ProductId = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetUpdateTime(v int64) *CalendarProductListCalendarAriResponseData {
	s.UpdateTime = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetCreateTime(v int64) *CalendarProductListCalendarAriResponseData {
	s.CreateTime = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetGwDescription(v string) *CalendarProductListCalendarAriResponseData {
	s.GwDescription = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetCategoryId(v int64) *CalendarProductListCalendarAriResponseData {
	s.CategoryId = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetCategoryName(v string) *CalendarProductListCalendarAriResponseData {
	s.CategoryName = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetGwErrorCode(v int32) *CalendarProductListCalendarAriResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetMarketingAmount(v int64) *CalendarProductListCalendarAriResponseData {
	s.MarketingAmount = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetProductName(v string) *CalendarProductListCalendarAriResponseData {
	s.ProductName = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetOutId(v string) *CalendarProductListCalendarAriResponseData {
	s.OutId = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseData) SetProductStatus(v int64) *CalendarProductListCalendarAriResponseData {
	s.ProductStatus = &v
	return s
}

type CalendarProductListCalendarAriResponseDataCalendarAriListItem struct {
	AvailableQty   *int64  `json:"available_qty,omitempty" xml:"available_qty,omitempty"`
	CalendarStatus *int64  `json:"calendar_status,omitempty" xml:"calendar_status,omitempty"`
	Date           *string `json:"date,omitempty" xml:"date,omitempty"`
	SoldQty        *int64  `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	TotalQty       *int64  `json:"total_qty,omitempty" xml:"total_qty,omitempty"`
	ActualAmount   *int64  `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s CalendarProductListCalendarAriResponseDataCalendarAriListItem) String() string {
	return tea.Prettify(s)
}

func (s CalendarProductListCalendarAriResponseDataCalendarAriListItem) GoString() string {
	return s.String()
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetAvailableQty(v int64) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.AvailableQty = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetCalendarStatus(v int64) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.CalendarStatus = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetDate(v string) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.Date = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetSoldQty(v int64) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.SoldQty = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetTotalQty(v int64) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.TotalQty = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseDataCalendarAriListItem) SetActualAmount(v int64) *CalendarProductListCalendarAriResponseDataCalendarAriListItem {
	s.ActualAmount = &v
	return s
}

type CalendarProductListCalendarAriResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CalendarProductListCalendarAriResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CalendarProductListCalendarAriResponseExtra) GoString() string {
	return s.String()
}

func (s *CalendarProductListCalendarAriResponseExtra) SetSubDescription(v string) *CalendarProductListCalendarAriResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseExtra) SetSubErrorCode(v int32) *CalendarProductListCalendarAriResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseExtra) SetDescription(v string) *CalendarProductListCalendarAriResponseExtra {
	s.Description = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseExtra) SetErrorCode(v int32) *CalendarProductListCalendarAriResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseExtra) SetLogid(v string) *CalendarProductListCalendarAriResponseExtra {
	s.Logid = &v
	return s
}

func (s *CalendarProductListCalendarAriResponseExtra) SetNow(v int64) *CalendarProductListCalendarAriResponseExtra {
	s.Now = &v
	return s
}

type CancelAuditRequest struct {
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OrderId      *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Reason       *string            `json:"reason,omitempty" xml:"reason,omitempty"`
	CancelId     *string            `json:"cancel_Id,omitempty" xml:"cancel_Id,omitempty" require:"true"`
	CancelResult *int               `json:"cancel_result,omitempty" xml:"cancel_result,omitempty" require:"true"`
	CancelType   *int               `json:"cancel_type,omitempty" xml:"cancel_type,omitempty" require:"true"`
}

func (s CancelAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAuditRequest) GoString() string {
	return s.String()
}

func (s *CancelAuditRequest) SetHeader(v map[string]*string) *CancelAuditRequest {
	s.Header = v
	return s
}

func (s *CancelAuditRequest) SetAccessToken(v string) *CancelAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *CancelAuditRequest) SetOrderId(v string) *CancelAuditRequest {
	s.OrderId = &v
	return s
}

func (s *CancelAuditRequest) SetReason(v string) *CancelAuditRequest {
	s.Reason = &v
	return s
}

func (s *CancelAuditRequest) SetCancelId(v string) *CancelAuditRequest {
	s.CancelId = &v
	return s
}

func (s *CancelAuditRequest) SetCancelResult(v int) *CancelAuditRequest {
	s.CancelResult = &v
	return s
}

func (s *CancelAuditRequest) SetCancelType(v int) *CancelAuditRequest {
	s.CancelType = &v
	return s
}

type CancelAuditResponse struct {
	Data  *CancelAuditResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *CancelAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CancelAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAuditResponse) GoString() string {
	return s.String()
}

func (s *CancelAuditResponse) SetData(v *CancelAuditResponseData) *CancelAuditResponse {
	s.Data = v
	return s
}

func (s *CancelAuditResponse) SetExtra(v *CancelAuditResponseExtra) *CancelAuditResponse {
	s.Extra = v
	return s
}

type CancelAuditResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CancelAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s CancelAuditResponseData) GoString() string {
	return s.String()
}

func (s *CancelAuditResponseData) SetGwDescription(v string) *CancelAuditResponseData {
	s.GwDescription = &v
	return s
}

func (s *CancelAuditResponseData) SetGwErrorCode(v int32) *CancelAuditResponseData {
	s.GwErrorCode = &v
	return s
}

type CancelAuditResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CancelAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CancelAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *CancelAuditResponseExtra) SetDescription(v string) *CancelAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *CancelAuditResponseExtra) SetErrorCode(v int32) *CancelAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CancelAuditResponseExtra) SetLogid(v string) *CancelAuditResponseExtra {
	s.Logid = &v
	return s
}

func (s *CancelAuditResponseExtra) SetNow(v int64) *CancelAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *CancelAuditResponseExtra) SetSubDescription(v string) *CancelAuditResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CancelAuditResponseExtra) SetSubErrorCode(v int32) *CancelAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CapacityApplyCapacityRequest struct {
	CapacityKey *string              `json:"capacity_key,omitempty" xml:"capacity_key,omitempty" require:"true"`
	ApplyInfo   map[string][]*string `json:"apply_info,omitempty" xml:"apply_info,omitempty" require:"true"`
	ApplyReason *string              `json:"apply_reason,omitempty" xml:"apply_reason,omitempty" require:"true"`
	Header      map[string]*string   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CapacityApplyCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityApplyCapacityRequest) GoString() string {
	return s.String()
}

func (s *CapacityApplyCapacityRequest) SetCapacityKey(v string) *CapacityApplyCapacityRequest {
	s.CapacityKey = &v
	return s
}

func (s *CapacityApplyCapacityRequest) SetApplyInfo(v map[string][]*string) *CapacityApplyCapacityRequest {
	s.ApplyInfo = v
	return s
}

func (s *CapacityApplyCapacityRequest) SetApplyReason(v string) *CapacityApplyCapacityRequest {
	s.ApplyReason = &v
	return s
}

func (s *CapacityApplyCapacityRequest) SetHeader(v map[string]*string) *CapacityApplyCapacityRequest {
	s.Header = v
	return s
}

func (s *CapacityApplyCapacityRequest) SetAccessToken(v string) *CapacityApplyCapacityRequest {
	s.AccessToken = &v
	return s
}

type CapacityApplyCapacityResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CapacityApplyCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityApplyCapacityResponse) GoString() string {
	return s.String()
}

func (s *CapacityApplyCapacityResponse) SetErrMsg(v string) *CapacityApplyCapacityResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityApplyCapacityResponse) SetErrNo(v int32) *CapacityApplyCapacityResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityApplyCapacityResponse) SetLogId(v string) *CapacityApplyCapacityResponse {
	s.LogId = &v
	return s
}

type CapacityBindAwemeRelationRequest struct {
	EmployeeInfo      *CapacityBindAwemeRelationRequestEmployeeInfo                      `json:"employee_info,omitempty" xml:"employee_info,omitempty"`
	Type              *string                                                            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	AwemeId           *string                                                            `json:"aweme_id,omitempty" xml:"aweme_id,omitempty" require:"true"`
	Header            map[string]*string                                                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CooperationInfo   *CapacityBindAwemeRelationRequestCooperationInfo                   `json:"cooperation_info,omitempty" xml:"cooperation_info,omitempty"`
	AuditTemplateInfo map[string]*CapacityBindAwemeRelationRequestAuditTemplateInfoValue `json:"audit_template_info,omitempty" xml:"audit_template_info,omitempty"`
	CapacityList      []*string                                                          `json:"capacity_list,omitempty" xml:"capacity_list,omitempty" require:"true" type:"Repeated"`
	CoSubject         *bool                                                              `json:"co_subject,omitempty" xml:"co_subject,omitempty"`
}

func (s CapacityBindAwemeRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequest) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequest) SetEmployeeInfo(v *CapacityBindAwemeRelationRequestEmployeeInfo) *CapacityBindAwemeRelationRequest {
	s.EmployeeInfo = v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetType(v string) *CapacityBindAwemeRelationRequest {
	s.Type = &v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetAwemeId(v string) *CapacityBindAwemeRelationRequest {
	s.AwemeId = &v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetHeader(v map[string]*string) *CapacityBindAwemeRelationRequest {
	s.Header = v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetAccessToken(v string) *CapacityBindAwemeRelationRequest {
	s.AccessToken = &v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetCooperationInfo(v *CapacityBindAwemeRelationRequestCooperationInfo) *CapacityBindAwemeRelationRequest {
	s.CooperationInfo = v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetAuditTemplateInfo(v map[string]*CapacityBindAwemeRelationRequestAuditTemplateInfoValue) *CapacityBindAwemeRelationRequest {
	s.AuditTemplateInfo = v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetCapacityList(v []*string) *CapacityBindAwemeRelationRequest {
	s.CapacityList = v
	return s
}

func (s *CapacityBindAwemeRelationRequest) SetCoSubject(v bool) *CapacityBindAwemeRelationRequest {
	s.CoSubject = &v
	return s
}

type CapacityBindAwemeRelationRequestAuditTemplateInfoValue struct {
	Title           *string                                                                      `json:"title,omitempty" xml:"title,omitempty"`
	TemplateContent []*CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem `json:"template_content,omitempty" xml:"template_content,omitempty" type:"Repeated"`
	Id              *int64                                                                       `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValue) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValue) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValue) SetTitle(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValue {
	s.Title = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValue) SetTemplateContent(v []*CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) *CapacityBindAwemeRelationRequestAuditTemplateInfoValue {
	s.TemplateContent = v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValue) SetId(v int64) *CapacityBindAwemeRelationRequestAuditTemplateInfoValue {
	s.Id = &v
	return s
}

type CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem struct {
	Name       *string                                                                                                  `json:"name,omitempty" xml:"name,omitempty"`
	Children   map[string][]*CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem `json:"children,omitempty" xml:"children,omitempty"`
	MaterielId *string                                                                                                  `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
	ValList    []*string                                                                                                `json:"val_list,omitempty" xml:"val_list,omitempty" type:"Repeated"`
	ValType    *int                                                                                                     `json:"val_type,omitempty" xml:"val_type,omitempty"`
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) SetName(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem {
	s.Name = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) SetChildren(v map[string][]*CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem {
	s.Children = v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) SetMaterielId(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem {
	s.MaterielId = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) SetValList(v []*string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem {
	s.ValList = v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem) SetValType(v int) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItem {
	s.ValType = &v
	return s
}

type CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem struct {
	ValExample   []*string `json:"val_example,omitempty" xml:"val_example,omitempty" type:"Repeated"`
	ValList      []*string `json:"val_list,omitempty" xml:"val_list,omitempty" type:"Repeated"`
	RejectReason *string   `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	MaterielId   *string   `json:"materiel_id,omitempty" xml:"materiel_id,omitempty"`
	Name         *string   `json:"name,omitempty" xml:"name,omitempty"`
	ValType      *int      `json:"val_type,omitempty" xml:"val_type,omitempty"`
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetValExample(v []*string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValExample = v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetValList(v []*string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValList = v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetRejectReason(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.RejectReason = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetMaterielId(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.MaterielId = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetName(v string) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.Name = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem) SetValType(v int) *CapacityBindAwemeRelationRequestAuditTemplateInfoValueTemplateContentItemChildrenValueItem {
	s.ValType = &v
	return s
}

type CapacityBindAwemeRelationRequestCooperationInfo struct {
	ContractImage       *string   `json:"contract_image,omitempty" xml:"contract_image,omitempty"`
	ContractImageList   []*string `json:"contract_image_list,omitempty" xml:"contract_image_list,omitempty" type:"Repeated"`
	RelationExpireDate  *int64    `json:"relation_expire_date,omitempty" xml:"relation_expire_date,omitempty" require:"true"`
	CooperationType     *string   `json:"cooperation_type,omitempty" xml:"cooperation_type,omitempty" require:"true"`
	CompanyName         *string   `json:"company_name,omitempty" xml:"company_name,omitempty"`
	BusinessLicenseCode *string   `json:"business_license_code,omitempty" xml:"business_license_code,omitempty"`
	RealName            *string   `json:"real_name,omitempty" xml:"real_name,omitempty"`
	IdentityNumber      *string   `json:"identity_number,omitempty" xml:"identity_number,omitempty"`
}

func (s CapacityBindAwemeRelationRequestCooperationInfo) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequestCooperationInfo) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetContractImage(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.ContractImage = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetContractImageList(v []*string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.ContractImageList = v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetRelationExpireDate(v int64) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.RelationExpireDate = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetCooperationType(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.CooperationType = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetCompanyName(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.CompanyName = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetBusinessLicenseCode(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.BusinessLicenseCode = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetRealName(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.RealName = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestCooperationInfo) SetIdentityNumber(v string) *CapacityBindAwemeRelationRequestCooperationInfo {
	s.IdentityNumber = &v
	return s
}

type CapacityBindAwemeRelationRequestEmployeeInfo struct {
	QualificationImageList []*string `json:"qualification_image_list,omitempty" xml:"qualification_image_list,omitempty" type:"Repeated"`
	RealName               *string   `json:"real_name,omitempty" xml:"real_name,omitempty" require:"true"`
	IdentityNumber         *string   `json:"identity_number,omitempty" xml:"identity_number,omitempty" require:"true"`
	ContractImage          *string   `json:"contract_image,omitempty" xml:"contract_image,omitempty"`
	RelationExpireDate     *int64    `json:"relation_expire_date,omitempty" xml:"relation_expire_date,omitempty" require:"true"`
	QualificationType      *string   `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
}

func (s CapacityBindAwemeRelationRequestEmployeeInfo) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationRequestEmployeeInfo) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetQualificationImageList(v []*string) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.QualificationImageList = v
	return s
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetRealName(v string) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.RealName = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetIdentityNumber(v string) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.IdentityNumber = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetContractImage(v string) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.ContractImage = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetRelationExpireDate(v int64) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.RelationExpireDate = &v
	return s
}

func (s *CapacityBindAwemeRelationRequestEmployeeInfo) SetQualificationType(v string) *CapacityBindAwemeRelationRequestEmployeeInfo {
	s.QualificationType = &v
	return s
}

type CapacityBindAwemeRelationResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CapacityBindAwemeRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityBindAwemeRelationResponse) GoString() string {
	return s.String()
}

func (s *CapacityBindAwemeRelationResponse) SetErrNo(v int32) *CapacityBindAwemeRelationResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityBindAwemeRelationResponse) SetErrMsg(v string) *CapacityBindAwemeRelationResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityBindAwemeRelationResponse) SetLogId(v string) *CapacityBindAwemeRelationResponse {
	s.LogId = &v
	return s
}

type CapacityDeleteAliasRequest struct {
	Alias       *string            `json:"alias,omitempty" xml:"alias,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CapacityDeleteAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityDeleteAliasRequest) GoString() string {
	return s.String()
}

func (s *CapacityDeleteAliasRequest) SetAlias(v string) *CapacityDeleteAliasRequest {
	s.Alias = &v
	return s
}

func (s *CapacityDeleteAliasRequest) SetHeader(v map[string]*string) *CapacityDeleteAliasRequest {
	s.Header = v
	return s
}

func (s *CapacityDeleteAliasRequest) SetAccessToken(v string) *CapacityDeleteAliasRequest {
	s.AccessToken = &v
	return s
}

type CapacityDeleteAliasResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CapacityDeleteAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityDeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *CapacityDeleteAliasResponse) SetErrNo(v int32) *CapacityDeleteAliasResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityDeleteAliasResponse) SetErrMsg(v string) *CapacityDeleteAliasResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityDeleteAliasResponse) SetLogId(v string) *CapacityDeleteAliasResponse {
	s.LogId = &v
	return s
}

type CapacityModifyAliasRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AfterAlias  *string            `json:"after_alias,omitempty" xml:"after_alias,omitempty" require:"true"`
	BeforeAlias *string            `json:"before_alias,omitempty" xml:"before_alias,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CapacityModifyAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityModifyAliasRequest) GoString() string {
	return s.String()
}

func (s *CapacityModifyAliasRequest) SetAccessToken(v string) *CapacityModifyAliasRequest {
	s.AccessToken = &v
	return s
}

func (s *CapacityModifyAliasRequest) SetAfterAlias(v string) *CapacityModifyAliasRequest {
	s.AfterAlias = &v
	return s
}

func (s *CapacityModifyAliasRequest) SetBeforeAlias(v string) *CapacityModifyAliasRequest {
	s.BeforeAlias = &v
	return s
}

func (s *CapacityModifyAliasRequest) SetHeader(v map[string]*string) *CapacityModifyAliasRequest {
	s.Header = v
	return s
}

type CapacityModifyAliasResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CapacityModifyAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityModifyAliasResponse) GoString() string {
	return s.String()
}

func (s *CapacityModifyAliasResponse) SetErrMsg(v string) *CapacityModifyAliasResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityModifyAliasResponse) SetLogId(v string) *CapacityModifyAliasResponse {
	s.LogId = &v
	return s
}

func (s *CapacityModifyAliasResponse) SetErrNo(v int32) *CapacityModifyAliasResponse {
	s.ErrNo = &v
	return s
}

type CapacityQueryAdIncomeRequest struct {
	HostName    *string            `json:"host_name,omitempty" xml:"host_name,omitempty"`
	StartDate   *string            `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
	EndDate     *string            `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CapacityQueryAdIncomeRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryAdIncomeRequest) GoString() string {
	return s.String()
}

func (s *CapacityQueryAdIncomeRequest) SetHostName(v string) *CapacityQueryAdIncomeRequest {
	s.HostName = &v
	return s
}

func (s *CapacityQueryAdIncomeRequest) SetStartDate(v string) *CapacityQueryAdIncomeRequest {
	s.StartDate = &v
	return s
}

func (s *CapacityQueryAdIncomeRequest) SetEndDate(v string) *CapacityQueryAdIncomeRequest {
	s.EndDate = &v
	return s
}

func (s *CapacityQueryAdIncomeRequest) SetHeader(v map[string]*string) *CapacityQueryAdIncomeRequest {
	s.Header = v
	return s
}

func (s *CapacityQueryAdIncomeRequest) SetAccessToken(v string) *CapacityQueryAdIncomeRequest {
	s.AccessToken = &v
	return s
}

type CapacityQueryAdIncomeResponse struct {
	Data   *CapacityQueryAdIncomeResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s CapacityQueryAdIncomeResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryAdIncomeResponse) GoString() string {
	return s.String()
}

func (s *CapacityQueryAdIncomeResponse) SetData(v *CapacityQueryAdIncomeResponseData) *CapacityQueryAdIncomeResponse {
	s.Data = v
	return s
}

func (s *CapacityQueryAdIncomeResponse) SetErrNo(v int32) *CapacityQueryAdIncomeResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityQueryAdIncomeResponse) SetErrMsg(v string) *CapacityQueryAdIncomeResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityQueryAdIncomeResponse) SetLogId(v string) *CapacityQueryAdIncomeResponse {
	s.LogId = &v
	return s
}

type CapacityQueryAdIncomeResponseData struct {
	IncomeList []*CapacityQueryAdIncomeResponseDataIncomeListItem `json:"income_list,omitempty" xml:"income_list,omitempty" require:"true" type:"Repeated"`
}

func (s CapacityQueryAdIncomeResponseData) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryAdIncomeResponseData) GoString() string {
	return s.String()
}

func (s *CapacityQueryAdIncomeResponseData) SetIncomeList(v []*CapacityQueryAdIncomeResponseDataIncomeListItem) *CapacityQueryAdIncomeResponseData {
	s.IncomeList = v
	return s
}

type CapacityQueryAdIncomeResponseDataIncomeListItem struct {
	Date              *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	Income            *int64  `json:"income,omitempty" xml:"income,omitempty" require:"true"`
	IncentiveFundCost *int64  `json:"incentive_fund_cost,omitempty" xml:"incentive_fund_cost,omitempty"`
}

func (s CapacityQueryAdIncomeResponseDataIncomeListItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryAdIncomeResponseDataIncomeListItem) GoString() string {
	return s.String()
}

func (s *CapacityQueryAdIncomeResponseDataIncomeListItem) SetDate(v string) *CapacityQueryAdIncomeResponseDataIncomeListItem {
	s.Date = &v
	return s
}

func (s *CapacityQueryAdIncomeResponseDataIncomeListItem) SetIncome(v int64) *CapacityQueryAdIncomeResponseDataIncomeListItem {
	s.Income = &v
	return s
}

func (s *CapacityQueryAdIncomeResponseDataIncomeListItem) SetIncentiveFundCost(v int64) *CapacityQueryAdIncomeResponseDataIncomeListItem {
	s.IncentiveFundCost = &v
	return s
}

type CapacityQueryApplyStatusRequest struct {
	CapacityKey *string            `json:"capacity_key,omitempty" xml:"capacity_key,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CapacityQueryApplyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryApplyStatusRequest) GoString() string {
	return s.String()
}

func (s *CapacityQueryApplyStatusRequest) SetCapacityKey(v string) *CapacityQueryApplyStatusRequest {
	s.CapacityKey = &v
	return s
}

func (s *CapacityQueryApplyStatusRequest) SetHeader(v map[string]*string) *CapacityQueryApplyStatusRequest {
	s.Header = v
	return s
}

func (s *CapacityQueryApplyStatusRequest) SetAccessToken(v string) *CapacityQueryApplyStatusRequest {
	s.AccessToken = &v
	return s
}

type CapacityQueryApplyStatusResponse struct {
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CapacityQueryApplyStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CapacityQueryApplyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryApplyStatusResponse) GoString() string {
	return s.String()
}

func (s *CapacityQueryApplyStatusResponse) SetErrMsg(v string) *CapacityQueryApplyStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacityQueryApplyStatusResponse) SetErrNo(v int32) *CapacityQueryApplyStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityQueryApplyStatusResponse) SetLogId(v string) *CapacityQueryApplyStatusResponse {
	s.LogId = &v
	return s
}

func (s *CapacityQueryApplyStatusResponse) SetData(v *CapacityQueryApplyStatusResponseData) *CapacityQueryApplyStatusResponse {
	s.Data = v
	return s
}

type CapacityQueryApplyStatusResponseData struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CapacityQueryApplyStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryApplyStatusResponseData) GoString() string {
	return s.String()
}

func (s *CapacityQueryApplyStatusResponseData) SetReason(v string) *CapacityQueryApplyStatusResponseData {
	s.Reason = &v
	return s
}

func (s *CapacityQueryApplyStatusResponseData) SetStatus(v int32) *CapacityQueryApplyStatusResponseData {
	s.Status = &v
	return s
}

type CapacityQueryCapacityListRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CapacityQueryCapacityListRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListRequest) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListRequest) SetHeader(v map[string]*string) *CapacityQueryCapacityListRequest {
	s.Header = v
	return s
}

func (s *CapacityQueryCapacityListRequest) SetAccessToken(v string) *CapacityQueryCapacityListRequest {
	s.AccessToken = &v
	return s
}

type CapacityQueryCapacityListResponse struct {
	ErrNo  *int32                                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CapacityQueryCapacityListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CapacityQueryCapacityListResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListResponse) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListResponse) SetErrNo(v int32) *CapacityQueryCapacityListResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityQueryCapacityListResponse) SetLogId(v string) *CapacityQueryCapacityListResponse {
	s.LogId = &v
	return s
}

func (s *CapacityQueryCapacityListResponse) SetData(v *CapacityQueryCapacityListResponseData) *CapacityQueryCapacityListResponse {
	s.Data = v
	return s
}

func (s *CapacityQueryCapacityListResponse) SetErrMsg(v string) *CapacityQueryCapacityListResponse {
	s.ErrMsg = &v
	return s
}

type CapacityQueryCapacityListResponseData struct {
	AppCapacityInfoList []*CapacityQueryCapacityListResponseDataAppCapacityInfoListItem `json:"app_capacity_info_list,omitempty" xml:"app_capacity_info_list,omitempty" type:"Repeated"`
}

func (s CapacityQueryCapacityListResponseData) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListResponseData) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListResponseData) SetAppCapacityInfoList(v []*CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) *CapacityQueryCapacityListResponseData {
	s.AppCapacityInfoList = v
	return s
}

type CapacityQueryCapacityListResponseDataAppCapacityInfoListItem struct {
	CapacityKey      *string                                                                                 `json:"capacity_key,omitempty" xml:"capacity_key,omitempty" require:"true"`
	IsAuthPermission *bool                                                                                   `json:"is_auth_permission,omitempty" xml:"is_auth_permission,omitempty"`
	HostApp          *string                                                                                 `json:"host_app,omitempty" xml:"host_app,omitempty"`
	LabDescription   *string                                                                                 `json:"lab_description,omitempty" xml:"lab_description,omitempty"`
	ApplyCondition   [][]*CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem `json:"apply_condition,omitempty" xml:"apply_condition,omitempty" type:"Repeated"`
	Description      *string                                                                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PermissionList   []*int32                                                                                `json:"permission_list,omitempty" xml:"permission_list,omitempty" type:"Repeated"`
	ApplyFieldList   []*CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem       `json:"apply_field_list,omitempty" xml:"apply_field_list,omitempty" type:"Repeated"`
	IsCanApply       *bool                                                                                   `json:"is_can_apply,omitempty" xml:"is_can_apply,omitempty"`
	IsLab            *int32                                                                                  `json:"is_lab,omitempty" xml:"is_lab,omitempty"`
	CapacityName     *string                                                                                 `json:"capacity_name,omitempty" xml:"capacity_name,omitempty" require:"true"`
	Status           *int32                                                                                  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetCapacityKey(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.CapacityKey = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetIsAuthPermission(v bool) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.IsAuthPermission = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetHostApp(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.HostApp = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetLabDescription(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.LabDescription = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetApplyCondition(v [][]*CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.ApplyCondition = v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetDescription(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.Description = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetPermissionList(v []*int32) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.PermissionList = v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetApplyFieldList(v []*CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.ApplyFieldList = v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetIsCanApply(v bool) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.IsCanApply = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetIsLab(v int32) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.IsLab = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetCapacityName(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.CapacityName = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem) SetStatus(v int32) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItem {
	s.Status = &v
	return s
}

type CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem struct {
	ConditionName        *string `json:"condition_name,omitempty" xml:"condition_name,omitempty"`
	Satisfied            *bool   `json:"satisfied,omitempty" xml:"satisfied,omitempty"`
	ConditionDescription *string `json:"condition_description,omitempty" xml:"condition_description,omitempty"`
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) SetConditionName(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem {
	s.ConditionName = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) SetSatisfied(v bool) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem {
	s.Satisfied = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem) SetConditionDescription(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyConditionItemItem {
	s.ConditionDescription = &v
	return s
}

type CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem struct {
	ApplyFieldValueType  *int32  `json:"apply_field_value_type,omitempty" xml:"apply_field_value_type,omitempty"`
	ApplyFieldDesc       *string `json:"apply_field_desc,omitempty" xml:"apply_field_desc,omitempty"`
	ApplyFieldIsRequired *bool   `json:"apply_field_is_required,omitempty" xml:"apply_field_is_required,omitempty"`
	ApplyFieldKey        *string `json:"apply_field_key,omitempty" xml:"apply_field_key,omitempty"`
	ApplyFieldName       *string `json:"apply_field_name,omitempty" xml:"apply_field_name,omitempty" require:"true"`
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) String() string {
	return tea.Prettify(s)
}

func (s CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) GoString() string {
	return s.String()
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) SetApplyFieldValueType(v int32) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem {
	s.ApplyFieldValueType = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) SetApplyFieldDesc(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem {
	s.ApplyFieldDesc = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) SetApplyFieldIsRequired(v bool) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem {
	s.ApplyFieldIsRequired = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) SetApplyFieldKey(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem {
	s.ApplyFieldKey = &v
	return s
}

func (s *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem) SetApplyFieldName(v string) *CapacityQueryCapacityListResponseDataAppCapacityInfoListItemApplyFieldListItem {
	s.ApplyFieldName = &v
	return s
}

type CapacitySetSearchTagRequest struct {
	AddTagList    []*string                                       `json:"add_tag_list,omitempty" xml:"add_tag_list,omitempty" type:"Repeated"`
	Header        map[string]*string                              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DeleteTagList []*string                                       `json:"delete_tag_list,omitempty" xml:"delete_tag_list,omitempty" type:"Repeated"`
	ModifyTagList []*CapacitySetSearchTagRequestModifyTagListItem `json:"modify_tag_list,omitempty" xml:"modify_tag_list,omitempty" type:"Repeated"`
}

func (s CapacitySetSearchTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacitySetSearchTagRequest) GoString() string {
	return s.String()
}

func (s *CapacitySetSearchTagRequest) SetAddTagList(v []*string) *CapacitySetSearchTagRequest {
	s.AddTagList = v
	return s
}

func (s *CapacitySetSearchTagRequest) SetHeader(v map[string]*string) *CapacitySetSearchTagRequest {
	s.Header = v
	return s
}

func (s *CapacitySetSearchTagRequest) SetAccessToken(v string) *CapacitySetSearchTagRequest {
	s.AccessToken = &v
	return s
}

func (s *CapacitySetSearchTagRequest) SetDeleteTagList(v []*string) *CapacitySetSearchTagRequest {
	s.DeleteTagList = v
	return s
}

func (s *CapacitySetSearchTagRequest) SetModifyTagList(v []*CapacitySetSearchTagRequestModifyTagListItem) *CapacitySetSearchTagRequest {
	s.ModifyTagList = v
	return s
}

type CapacitySetSearchTagRequestModifyTagListItem struct {
	SearchTag *string `json:"search_tag,omitempty" xml:"search_tag,omitempty" require:"true"`
	BeforeTag *string `json:"before_tag,omitempty" xml:"before_tag,omitempty" require:"true"`
}

func (s CapacitySetSearchTagRequestModifyTagListItem) String() string {
	return tea.Prettify(s)
}

func (s CapacitySetSearchTagRequestModifyTagListItem) GoString() string {
	return s.String()
}

func (s *CapacitySetSearchTagRequestModifyTagListItem) SetSearchTag(v string) *CapacitySetSearchTagRequestModifyTagListItem {
	s.SearchTag = &v
	return s
}

func (s *CapacitySetSearchTagRequestModifyTagListItem) SetBeforeTag(v string) *CapacitySetSearchTagRequestModifyTagListItem {
	s.BeforeTag = &v
	return s
}

type CapacitySetSearchTagResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CapacitySetSearchTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacitySetSearchTagResponse) GoString() string {
	return s.String()
}

func (s *CapacitySetSearchTagResponse) SetErrMsg(v string) *CapacitySetSearchTagResponse {
	s.ErrMsg = &v
	return s
}

func (s *CapacitySetSearchTagResponse) SetLogId(v string) *CapacitySetSearchTagResponse {
	s.LogId = &v
	return s
}

func (s *CapacitySetSearchTagResponse) SetErrNo(v int32) *CapacitySetSearchTagResponse {
	s.ErrNo = &v
	return s
}

type CapacityUnbindAwemeRelationRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AwemeId     *string            `json:"aweme_id,omitempty" xml:"aweme_id,omitempty" require:"true"`
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CapacityUnbindAwemeRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s CapacityUnbindAwemeRelationRequest) GoString() string {
	return s.String()
}

func (s *CapacityUnbindAwemeRelationRequest) SetAccessToken(v string) *CapacityUnbindAwemeRelationRequest {
	s.AccessToken = &v
	return s
}

func (s *CapacityUnbindAwemeRelationRequest) SetAwemeId(v string) *CapacityUnbindAwemeRelationRequest {
	s.AwemeId = &v
	return s
}

func (s *CapacityUnbindAwemeRelationRequest) SetType(v string) *CapacityUnbindAwemeRelationRequest {
	s.Type = &v
	return s
}

func (s *CapacityUnbindAwemeRelationRequest) SetHeader(v map[string]*string) *CapacityUnbindAwemeRelationRequest {
	s.Header = v
	return s
}

type CapacityUnbindAwemeRelationResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s CapacityUnbindAwemeRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s CapacityUnbindAwemeRelationResponse) GoString() string {
	return s.String()
}

func (s *CapacityUnbindAwemeRelationResponse) SetLogId(v string) *CapacityUnbindAwemeRelationResponse {
	s.LogId = &v
	return s
}

func (s *CapacityUnbindAwemeRelationResponse) SetErrNo(v int32) *CapacityUnbindAwemeRelationResponse {
	s.ErrNo = &v
	return s
}

func (s *CapacityUnbindAwemeRelationResponse) SetErrMsg(v string) *CapacityUnbindAwemeRelationResponse {
	s.ErrMsg = &v
	return s
}

type CarCommentRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CarCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s CarCommentRequest) GoString() string {
	return s.String()
}

func (s *CarCommentRequest) SetHeader(v map[string]*string) *CarCommentRequest {
	s.Header = v
	return s
}

func (s *CarCommentRequest) SetAccessToken(v string) *CarCommentRequest {
	s.AccessToken = &v
	return s
}

type CarCommentResponse struct {
	Extra *CarCommentResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CarCommentResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CarCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s CarCommentResponse) GoString() string {
	return s.String()
}

func (s *CarCommentResponse) SetExtra(v *CarCommentResponseExtra) *CarCommentResponse {
	s.Extra = v
	return s
}

func (s *CarCommentResponse) SetData(v *CarCommentResponseData) *CarCommentResponse {
	s.Data = v
	return s
}

type CarCommentResponseData struct {
	List          []*CarCommentResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CarCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s CarCommentResponseData) GoString() string {
	return s.String()
}

func (s *CarCommentResponseData) SetList(v []*CarCommentResponseDataListItem) *CarCommentResponseData {
	s.List = v
	return s
}

func (s *CarCommentResponseData) SetGwErrorCode(v int32) *CarCommentResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CarCommentResponseData) SetGwDescription(v string) *CarCommentResponseData {
	s.GwDescription = &v
	return s
}

type CarCommentResponseDataListItem struct {
	EffectValue      *float64                                       `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CarCommentResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                         `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                        `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                        `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                        `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                         `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                         `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
}

func (s CarCommentResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CarCommentResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CarCommentResponseDataListItem) SetEffectValue(v float64) *CarCommentResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetVideoList(v []*CarCommentResponseDataListItemVideoListItem) *CarCommentResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CarCommentResponseDataListItem) SetRank(v int32) *CarCommentResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetRankChange(v string) *CarCommentResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetNickname(v string) *CarCommentResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetAvatar(v string) *CarCommentResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetFollowerCount(v int64) *CarCommentResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CarCommentResponseDataListItem) SetOnbillbaordTimes(v int32) *CarCommentResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

type CarCommentResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CarCommentResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CarCommentResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CarCommentResponseDataListItemVideoListItem) SetItemCover(v string) *CarCommentResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CarCommentResponseDataListItemVideoListItem) SetShareUrl(v string) *CarCommentResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CarCommentResponseDataListItemVideoListItem) SetTitle(v string) *CarCommentResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type CarCommentResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CarCommentResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CarCommentResponseExtra) GoString() string {
	return s.String()
}

func (s *CarCommentResponseExtra) SetSubErrorCode(v int32) *CarCommentResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CarCommentResponseExtra) SetSubDescription(v string) *CarCommentResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CarCommentResponseExtra) SetLogid(v string) *CarCommentResponseExtra {
	s.Logid = &v
	return s
}

func (s *CarCommentResponseExtra) SetNow(v int64) *CarCommentResponseExtra {
	s.Now = &v
	return s
}

func (s *CarCommentResponseExtra) SetErrorCode(v int32) *CarCommentResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CarCommentResponseExtra) SetDescription(v string) *CarCommentResponseExtra {
	s.Description = &v
	return s
}

type CarDriverRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CarDriverRequest) String() string {
	return tea.Prettify(s)
}

func (s CarDriverRequest) GoString() string {
	return s.String()
}

func (s *CarDriverRequest) SetAccessToken(v string) *CarDriverRequest {
	s.AccessToken = &v
	return s
}

func (s *CarDriverRequest) SetHeader(v map[string]*string) *CarDriverRequest {
	s.Header = v
	return s
}

type CarDriverResponse struct {
	Data  *CarDriverResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CarDriverResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CarDriverResponse) String() string {
	return tea.Prettify(s)
}

func (s CarDriverResponse) GoString() string {
	return s.String()
}

func (s *CarDriverResponse) SetData(v *CarDriverResponseData) *CarDriverResponse {
	s.Data = v
	return s
}

func (s *CarDriverResponse) SetExtra(v *CarDriverResponseExtra) *CarDriverResponse {
	s.Extra = v
	return s
}

type CarDriverResponseData struct {
	GwDescription *string                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*CarDriverResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CarDriverResponseData) String() string {
	return tea.Prettify(s)
}

func (s CarDriverResponseData) GoString() string {
	return s.String()
}

func (s *CarDriverResponseData) SetGwDescription(v string) *CarDriverResponseData {
	s.GwDescription = &v
	return s
}

func (s *CarDriverResponseData) SetList(v []*CarDriverResponseDataListItem) *CarDriverResponseData {
	s.List = v
	return s
}

func (s *CarDriverResponseData) SetGwErrorCode(v int32) *CarDriverResponseData {
	s.GwErrorCode = &v
	return s
}

type CarDriverResponseDataListItem struct {
	Avatar           *string                                       `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                        `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                        `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                      `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CarDriverResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                        `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                       `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                       `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
}

func (s CarDriverResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CarDriverResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CarDriverResponseDataListItem) SetAvatar(v string) *CarDriverResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetFollowerCount(v int64) *CarDriverResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetOnbillbaordTimes(v int32) *CarDriverResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetEffectValue(v float64) *CarDriverResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetVideoList(v []*CarDriverResponseDataListItemVideoListItem) *CarDriverResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CarDriverResponseDataListItem) SetRank(v int32) *CarDriverResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetRankChange(v string) *CarDriverResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CarDriverResponseDataListItem) SetNickname(v string) *CarDriverResponseDataListItem {
	s.Nickname = &v
	return s
}

type CarDriverResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s CarDriverResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CarDriverResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CarDriverResponseDataListItemVideoListItem) SetTitle(v string) *CarDriverResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CarDriverResponseDataListItemVideoListItem) SetItemCover(v string) *CarDriverResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CarDriverResponseDataListItemVideoListItem) SetShareUrl(v string) *CarDriverResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type CarDriverResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CarDriverResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CarDriverResponseExtra) GoString() string {
	return s.String()
}

func (s *CarDriverResponseExtra) SetErrorCode(v int32) *CarDriverResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CarDriverResponseExtra) SetDescription(v string) *CarDriverResponseExtra {
	s.Description = &v
	return s
}

func (s *CarDriverResponseExtra) SetSubErrorCode(v int32) *CarDriverResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CarDriverResponseExtra) SetSubDescription(v string) *CarDriverResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CarDriverResponseExtra) SetLogid(v string) *CarDriverResponseExtra {
	s.Logid = &v
	return s
}

func (s *CarDriverResponseExtra) SetNow(v int64) *CarDriverResponseExtra {
	s.Now = &v
	return s
}

type CarOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CarOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s CarOverallRequest) GoString() string {
	return s.String()
}

func (s *CarOverallRequest) SetHeader(v map[string]*string) *CarOverallRequest {
	s.Header = v
	return s
}

func (s *CarOverallRequest) SetAccessToken(v string) *CarOverallRequest {
	s.AccessToken = &v
	return s
}

type CarOverallResponse struct {
	Data  *CarOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *CarOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CarOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s CarOverallResponse) GoString() string {
	return s.String()
}

func (s *CarOverallResponse) SetData(v *CarOverallResponseData) *CarOverallResponse {
	s.Data = v
	return s
}

func (s *CarOverallResponse) SetExtra(v *CarOverallResponseExtra) *CarOverallResponse {
	s.Extra = v
	return s
}

type CarOverallResponseData struct {
	GwErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*CarOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s CarOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s CarOverallResponseData) GoString() string {
	return s.String()
}

func (s *CarOverallResponseData) SetGwErrorCode(v int32) *CarOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CarOverallResponseData) SetGwDescription(v string) *CarOverallResponseData {
	s.GwDescription = &v
	return s
}

func (s *CarOverallResponseData) SetList(v []*CarOverallResponseDataListItem) *CarOverallResponseData {
	s.List = v
	return s
}

type CarOverallResponseDataListItem struct {
	RankChange       *string                                        `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                        `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                        `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                         `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                         `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                       `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CarOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                         `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s CarOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CarOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CarOverallResponseDataListItem) SetRankChange(v string) *CarOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetNickname(v string) *CarOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetAvatar(v string) *CarOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetFollowerCount(v int64) *CarOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *CarOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetEffectValue(v float64) *CarOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CarOverallResponseDataListItem) SetVideoList(v []*CarOverallResponseDataListItemVideoListItem) *CarOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CarOverallResponseDataListItem) SetRank(v int32) *CarOverallResponseDataListItem {
	s.Rank = &v
	return s
}

type CarOverallResponseDataListItemVideoListItem struct {
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
}

func (s CarOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CarOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CarOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *CarOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CarOverallResponseDataListItemVideoListItem) SetTitle(v string) *CarOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CarOverallResponseDataListItemVideoListItem) SetItemCover(v string) *CarOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

type CarOverallResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CarOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CarOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *CarOverallResponseExtra) SetSubErrorCode(v int32) *CarOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CarOverallResponseExtra) SetSubDescription(v string) *CarOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CarOverallResponseExtra) SetLogid(v string) *CarOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *CarOverallResponseExtra) SetNow(v int64) *CarOverallResponseExtra {
	s.Now = &v
	return s
}

func (s *CarOverallResponseExtra) SetErrorCode(v int32) *CarOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CarOverallResponseExtra) SetDescription(v string) *CarOverallResponseExtra {
	s.Description = &v
	return s
}

type CarPlayRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CarPlayRequest) String() string {
	return tea.Prettify(s)
}

func (s CarPlayRequest) GoString() string {
	return s.String()
}

func (s *CarPlayRequest) SetHeader(v map[string]*string) *CarPlayRequest {
	s.Header = v
	return s
}

func (s *CarPlayRequest) SetAccessToken(v string) *CarPlayRequest {
	s.AccessToken = &v
	return s
}

type CarPlayResponse struct {
	Extra *CarPlayResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CarPlayResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CarPlayResponse) String() string {
	return tea.Prettify(s)
}

func (s CarPlayResponse) GoString() string {
	return s.String()
}

func (s *CarPlayResponse) SetExtra(v *CarPlayResponseExtra) *CarPlayResponse {
	s.Extra = v
	return s
}

func (s *CarPlayResponse) SetData(v *CarPlayResponseData) *CarPlayResponse {
	s.Data = v
	return s
}

type CarPlayResponseData struct {
	GwDescription *string                        `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*CarPlayResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                         `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CarPlayResponseData) String() string {
	return tea.Prettify(s)
}

func (s CarPlayResponseData) GoString() string {
	return s.String()
}

func (s *CarPlayResponseData) SetGwDescription(v string) *CarPlayResponseData {
	s.GwDescription = &v
	return s
}

func (s *CarPlayResponseData) SetList(v []*CarPlayResponseDataListItem) *CarPlayResponseData {
	s.List = v
	return s
}

func (s *CarPlayResponseData) SetGwErrorCode(v int32) *CarPlayResponseData {
	s.GwErrorCode = &v
	return s
}

type CarPlayResponseDataListItem struct {
	RankChange       *string                                     `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                     `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                     `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                      `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                      `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                    `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CarPlayResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                      `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s CarPlayResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CarPlayResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CarPlayResponseDataListItem) SetRankChange(v string) *CarPlayResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetNickname(v string) *CarPlayResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetAvatar(v string) *CarPlayResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetFollowerCount(v int64) *CarPlayResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetOnbillbaordTimes(v int32) *CarPlayResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetEffectValue(v float64) *CarPlayResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CarPlayResponseDataListItem) SetVideoList(v []*CarPlayResponseDataListItemVideoListItem) *CarPlayResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CarPlayResponseDataListItem) SetRank(v int32) *CarPlayResponseDataListItem {
	s.Rank = &v
	return s
}

type CarPlayResponseDataListItemVideoListItem struct {
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CarPlayResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CarPlayResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CarPlayResponseDataListItemVideoListItem) SetItemCover(v string) *CarPlayResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CarPlayResponseDataListItemVideoListItem) SetShareUrl(v string) *CarPlayResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

func (s *CarPlayResponseDataListItemVideoListItem) SetTitle(v string) *CarPlayResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

type CarPlayResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s CarPlayResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CarPlayResponseExtra) GoString() string {
	return s.String()
}

func (s *CarPlayResponseExtra) SetErrorCode(v int32) *CarPlayResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CarPlayResponseExtra) SetDescription(v string) *CarPlayResponseExtra {
	s.Description = &v
	return s
}

func (s *CarPlayResponseExtra) SetSubErrorCode(v int32) *CarPlayResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CarPlayResponseExtra) SetSubDescription(v string) *CarPlayResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CarPlayResponseExtra) SetLogid(v string) *CarPlayResponseExtra {
	s.Logid = &v
	return s
}

func (s *CarPlayResponseExtra) SetNow(v int64) *CarPlayResponseExtra {
	s.Now = &v
	return s
}

type CarUseRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CarUseRequest) String() string {
	return tea.Prettify(s)
}

func (s CarUseRequest) GoString() string {
	return s.String()
}

func (s *CarUseRequest) SetHeader(v map[string]*string) *CarUseRequest {
	s.Header = v
	return s
}

func (s *CarUseRequest) SetAccessToken(v string) *CarUseRequest {
	s.AccessToken = &v
	return s
}

type CarUseResponse struct {
	Extra *CarUseResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CarUseResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CarUseResponse) String() string {
	return tea.Prettify(s)
}

func (s CarUseResponse) GoString() string {
	return s.String()
}

func (s *CarUseResponse) SetExtra(v *CarUseResponseExtra) *CarUseResponse {
	s.Extra = v
	return s
}

func (s *CarUseResponse) SetData(v *CarUseResponseData) *CarUseResponse {
	s.Data = v
	return s
}

type CarUseResponseData struct {
	List          []*CarUseResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CarUseResponseData) String() string {
	return tea.Prettify(s)
}

func (s CarUseResponseData) GoString() string {
	return s.String()
}

func (s *CarUseResponseData) SetList(v []*CarUseResponseDataListItem) *CarUseResponseData {
	s.List = v
	return s
}

func (s *CarUseResponseData) SetGwErrorCode(v int32) *CarUseResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CarUseResponseData) SetGwDescription(v string) *CarUseResponseData {
	s.GwDescription = &v
	return s
}

type CarUseResponseDataListItem struct {
	OnbillbaordTimes *int32                                     `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                   `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*CarUseResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                     `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                    `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                    `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                    `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                     `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
}

func (s CarUseResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s CarUseResponseDataListItem) GoString() string {
	return s.String()
}

func (s *CarUseResponseDataListItem) SetOnbillbaordTimes(v int32) *CarUseResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *CarUseResponseDataListItem) SetEffectValue(v float64) *CarUseResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *CarUseResponseDataListItem) SetVideoList(v []*CarUseResponseDataListItemVideoListItem) *CarUseResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *CarUseResponseDataListItem) SetRank(v int32) *CarUseResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *CarUseResponseDataListItem) SetRankChange(v string) *CarUseResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *CarUseResponseDataListItem) SetNickname(v string) *CarUseResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *CarUseResponseDataListItem) SetAvatar(v string) *CarUseResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *CarUseResponseDataListItem) SetFollowerCount(v int64) *CarUseResponseDataListItem {
	s.FollowerCount = &v
	return s
}

type CarUseResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s CarUseResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s CarUseResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *CarUseResponseDataListItemVideoListItem) SetTitle(v string) *CarUseResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *CarUseResponseDataListItemVideoListItem) SetItemCover(v string) *CarUseResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *CarUseResponseDataListItemVideoListItem) SetShareUrl(v string) *CarUseResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type CarUseResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CarUseResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CarUseResponseExtra) GoString() string {
	return s.String()
}

func (s *CarUseResponseExtra) SetSubDescription(v string) *CarUseResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CarUseResponseExtra) SetLogid(v string) *CarUseResponseExtra {
	s.Logid = &v
	return s
}

func (s *CarUseResponseExtra) SetNow(v int64) *CarUseResponseExtra {
	s.Now = &v
	return s
}

func (s *CarUseResponseExtra) SetErrorCode(v int32) *CarUseResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CarUseResponseExtra) SetDescription(v string) *CarUseResponseExtra {
	s.Description = &v
	return s
}

func (s *CarUseResponseExtra) SetSubErrorCode(v int32) *CarUseResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CardDeleteRequest struct {
	Url         *string                  `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	UserId      *CardDeleteRequestUserId `json:"user_id,omitempty" xml:"user_id,omitempty"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CardDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteRequest) GoString() string {
	return s.String()
}

func (s *CardDeleteRequest) SetUrl(v string) *CardDeleteRequest {
	s.Url = &v
	return s
}

func (s *CardDeleteRequest) SetUserId(v *CardDeleteRequestUserId) *CardDeleteRequest {
	s.UserId = v
	return s
}

func (s *CardDeleteRequest) SetHeader(v map[string]*string) *CardDeleteRequest {
	s.Header = v
	return s
}

func (s *CardDeleteRequest) SetAccessToken(v string) *CardDeleteRequest {
	s.AccessToken = &v
	return s
}

type CardDeleteRequestUserId struct {
}

func (s CardDeleteRequestUserId) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteRequestUserId) GoString() string {
	return s.String()
}

type CardDeleteResponse struct {
	Errcode *int32  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Status  *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s CardDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CardDeleteResponse) GoString() string {
	return s.String()
}

func (s *CardDeleteResponse) SetErrcode(v int32) *CardDeleteResponse {
	s.Errcode = &v
	return s
}

func (s *CardDeleteResponse) SetErrmsg(v string) *CardDeleteResponse {
	s.Errmsg = &v
	return s
}

func (s *CardDeleteResponse) SetStatus(v int) *CardDeleteResponse {
	s.Status = &v
	return s
}

type CardGetRequest struct {
	UserId      *CardGetRequestUserId `json:"user_id,omitempty" xml:"user_id,omitempty"`
	Header      map[string]*string    `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string               `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Url         *string               `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s CardGetRequest) String() string {
	return tea.Prettify(s)
}

func (s CardGetRequest) GoString() string {
	return s.String()
}

func (s *CardGetRequest) SetUserId(v *CardGetRequestUserId) *CardGetRequest {
	s.UserId = v
	return s
}

func (s *CardGetRequest) SetHeader(v map[string]*string) *CardGetRequest {
	s.Header = v
	return s
}

func (s *CardGetRequest) SetAccessToken(v string) *CardGetRequest {
	s.AccessToken = &v
	return s
}

func (s *CardGetRequest) SetUrl(v string) *CardGetRequest {
	s.Url = &v
	return s
}

type CardGetRequestUserId struct {
}

func (s CardGetRequestUserId) String() string {
	return tea.Prettify(s)
}

func (s CardGetRequestUserId) GoString() string {
	return s.String()
}

type CardGetResponse struct {
	CardType     *int    `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Status       *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AuditOpinion *string `json:"audit_opinion,omitempty" xml:"audit_opinion,omitempty"`
	Errcode      *int32  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg       *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s CardGetResponse) String() string {
	return tea.Prettify(s)
}

func (s CardGetResponse) GoString() string {
	return s.String()
}

func (s *CardGetResponse) SetCardType(v int) *CardGetResponse {
	s.CardType = &v
	return s
}

func (s *CardGetResponse) SetValue(v string) *CardGetResponse {
	s.Value = &v
	return s
}

func (s *CardGetResponse) SetStatus(v int) *CardGetResponse {
	s.Status = &v
	return s
}

func (s *CardGetResponse) SetAuditOpinion(v string) *CardGetResponse {
	s.AuditOpinion = &v
	return s
}

func (s *CardGetResponse) SetErrcode(v int32) *CardGetResponse {
	s.Errcode = &v
	return s
}

func (s *CardGetResponse) SetErrmsg(v string) *CardGetResponse {
	s.Errmsg = &v
	return s
}

type CardImageDeleteRequest struct {
	ImageIds    *string            `json:"image_ids,omitempty" xml:"image_ids,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CardImageDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CardImageDeleteRequest) GoString() string {
	return s.String()
}

func (s *CardImageDeleteRequest) SetImageIds(v string) *CardImageDeleteRequest {
	s.ImageIds = &v
	return s
}

func (s *CardImageDeleteRequest) SetHeader(v map[string]*string) *CardImageDeleteRequest {
	s.Header = v
	return s
}

func (s *CardImageDeleteRequest) SetAccessToken(v string) *CardImageDeleteRequest {
	s.AccessToken = &v
	return s
}

type CardImageDeleteResponse struct {
	Errmsg       *string                                    `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	ImagesStatus []*CardImageDeleteResponseImagesStatusItem `json:"images_status,omitempty" xml:"images_status,omitempty" require:"true" type:"Repeated"`
	Errcode      *int32                                     `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s CardImageDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CardImageDeleteResponse) GoString() string {
	return s.String()
}

func (s *CardImageDeleteResponse) SetErrmsg(v string) *CardImageDeleteResponse {
	s.Errmsg = &v
	return s
}

func (s *CardImageDeleteResponse) SetImagesStatus(v []*CardImageDeleteResponseImagesStatusItem) *CardImageDeleteResponse {
	s.ImagesStatus = v
	return s
}

func (s *CardImageDeleteResponse) SetErrcode(v int32) *CardImageDeleteResponse {
	s.Errcode = &v
	return s
}

type CardImageDeleteResponseImagesStatusItem struct {
	ImageId      *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	DeleteStatus *int    `json:"delete_status,omitempty" xml:"delete_status,omitempty"`
}

func (s CardImageDeleteResponseImagesStatusItem) String() string {
	return tea.Prettify(s)
}

func (s CardImageDeleteResponseImagesStatusItem) GoString() string {
	return s.String()
}

func (s *CardImageDeleteResponseImagesStatusItem) SetImageId(v string) *CardImageDeleteResponseImagesStatusItem {
	s.ImageId = &v
	return s
}

func (s *CardImageDeleteResponseImagesStatusItem) SetDeleteStatus(v int) *CardImageDeleteResponseImagesStatusItem {
	s.DeleteStatus = &v
	return s
}

type CardImageGetRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ImageIds    *string            `json:"image_ids,omitempty" xml:"image_ids,omitempty" require:"true"`
}

func (s CardImageGetRequest) String() string {
	return tea.Prettify(s)
}

func (s CardImageGetRequest) GoString() string {
	return s.String()
}

func (s *CardImageGetRequest) SetHeader(v map[string]*string) *CardImageGetRequest {
	s.Header = v
	return s
}

func (s *CardImageGetRequest) SetAccessToken(v string) *CardImageGetRequest {
	s.AccessToken = &v
	return s
}

func (s *CardImageGetRequest) SetImageIds(v string) *CardImageGetRequest {
	s.ImageIds = &v
	return s
}

type CardImageGetResponse struct {
	ImagesStatus []*CardImageGetResponseImagesStatusItem `json:"images_status,omitempty" xml:"images_status,omitempty" type:"Repeated"`
	Errcode      *int32                                  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg       *string                                 `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s CardImageGetResponse) String() string {
	return tea.Prettify(s)
}

func (s CardImageGetResponse) GoString() string {
	return s.String()
}

func (s *CardImageGetResponse) SetImagesStatus(v []*CardImageGetResponseImagesStatusItem) *CardImageGetResponse {
	s.ImagesStatus = v
	return s
}

func (s *CardImageGetResponse) SetErrcode(v int32) *CardImageGetResponse {
	s.Errcode = &v
	return s
}

func (s *CardImageGetResponse) SetErrmsg(v string) *CardImageGetResponse {
	s.Errmsg = &v
	return s
}

type CardImageGetResponseImagesStatusItem struct {
	ImageStatus  *int    `json:"image_status,omitempty" xml:"image_status,omitempty"`
	AuditOpinion *string `json:"audit_opinion,omitempty" xml:"audit_opinion,omitempty"`
	ImageId      *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
}

func (s CardImageGetResponseImagesStatusItem) String() string {
	return tea.Prettify(s)
}

func (s CardImageGetResponseImagesStatusItem) GoString() string {
	return s.String()
}

func (s *CardImageGetResponseImagesStatusItem) SetImageStatus(v int) *CardImageGetResponseImagesStatusItem {
	s.ImageStatus = &v
	return s
}

func (s *CardImageGetResponseImagesStatusItem) SetAuditOpinion(v string) *CardImageGetResponseImagesStatusItem {
	s.AuditOpinion = &v
	return s
}

func (s *CardImageGetResponseImagesStatusItem) SetImageId(v string) *CardImageGetResponseImagesStatusItem {
	s.ImageId = &v
	return s
}

type CardImageUploadRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CardImageUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s CardImageUploadRequest) GoString() string {
	return s.String()
}

func (s *CardImageUploadRequest) SetHeader(v map[string]*string) *CardImageUploadRequest {
	s.Header = v
	return s
}

func (s *CardImageUploadRequest) SetAccessToken(v string) *CardImageUploadRequest {
	s.AccessToken = &v
	return s
}

type CardImageUploadResponse struct {
	Status *int `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CardImageUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s CardImageUploadResponse) GoString() string {
	return s.String()
}

func (s *CardImageUploadResponse) SetStatus(v int) *CardImageUploadResponse {
	s.Status = &v
	return s
}

type CardSetRequest struct {
	Value       *string                 `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	UserId      *CardSetRequestUserId   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	Url         *string                 `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	CardType    *CardSetRequestCardType `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
	Header      map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CardSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CardSetRequest) GoString() string {
	return s.String()
}

func (s *CardSetRequest) SetValue(v string) *CardSetRequest {
	s.Value = &v
	return s
}

func (s *CardSetRequest) SetUserId(v *CardSetRequestUserId) *CardSetRequest {
	s.UserId = v
	return s
}

func (s *CardSetRequest) SetUrl(v string) *CardSetRequest {
	s.Url = &v
	return s
}

func (s *CardSetRequest) SetCardType(v *CardSetRequestCardType) *CardSetRequest {
	s.CardType = v
	return s
}

func (s *CardSetRequest) SetHeader(v map[string]*string) *CardSetRequest {
	s.Header = v
	return s
}

func (s *CardSetRequest) SetAccessToken(v string) *CardSetRequest {
	s.AccessToken = &v
	return s
}

type CardSetRequestCardType struct {
}

func (s CardSetRequestCardType) String() string {
	return tea.Prettify(s)
}

func (s CardSetRequestCardType) GoString() string {
	return s.String()
}

type CardSetRequestUserId struct {
}

func (s CardSetRequestUserId) String() string {
	return tea.Prettify(s)
}

func (s CardSetRequestUserId) GoString() string {
	return s.String()
}

type CardSetResponse struct {
	Feedback *string `json:"feedback,omitempty" xml:"feedback,omitempty"`
	Status   *int    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CardSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CardSetResponse) GoString() string {
	return s.String()
}

func (s *CardSetResponse) SetFeedback(v string) *CardSetResponse {
	s.Feedback = &v
	return s
}

func (s *CardSetResponse) SetStatus(v int) *CardSetResponse {
	s.Status = &v
	return s
}

type CardUpdateRequest struct {
	Value       *string                    `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RoomId      *string                    `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	CardType    *CardUpdateRequestCardType `json:"card_type,omitempty" xml:"card_type,omitempty" require:"true"`
}

func (s CardUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CardUpdateRequest) GoString() string {
	return s.String()
}

func (s *CardUpdateRequest) SetValue(v string) *CardUpdateRequest {
	s.Value = &v
	return s
}

func (s *CardUpdateRequest) SetHeader(v map[string]*string) *CardUpdateRequest {
	s.Header = v
	return s
}

func (s *CardUpdateRequest) SetAccessToken(v string) *CardUpdateRequest {
	s.AccessToken = &v
	return s
}

func (s *CardUpdateRequest) SetRoomId(v string) *CardUpdateRequest {
	s.RoomId = &v
	return s
}

func (s *CardUpdateRequest) SetCardType(v *CardUpdateRequestCardType) *CardUpdateRequest {
	s.CardType = v
	return s
}

type CardUpdateRequestCardType struct {
}

func (s CardUpdateRequestCardType) String() string {
	return tea.Prettify(s)
}

func (s CardUpdateRequestCardType) GoString() string {
	return s.String()
}

type CardUpdateResponse struct {
	Status       *int    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AuditOpinion *string `json:"audit_opinion,omitempty" xml:"audit_opinion,omitempty"`
}

func (s CardUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CardUpdateResponse) GoString() string {
	return s.String()
}

func (s *CardUpdateResponse) SetStatus(v int) *CardUpdateResponse {
	s.Status = &v
	return s
}

func (s *CardUpdateResponse) SetAuditOpinion(v string) *CardUpdateResponse {
	s.AuditOpinion = &v
	return s
}

type CategoryGetAuditCategoriesRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CategoryGetAuditCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s CategoryGetAuditCategoriesRequest) GoString() string {
	return s.String()
}

func (s *CategoryGetAuditCategoriesRequest) SetHeader(v map[string]*string) *CategoryGetAuditCategoriesRequest {
	s.Header = v
	return s
}

func (s *CategoryGetAuditCategoriesRequest) SetAccessToken(v string) *CategoryGetAuditCategoriesRequest {
	s.AccessToken = &v
	return s
}

type CategoryGetAuditCategoriesResponse struct {
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *CategoryGetAuditCategoriesResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s CategoryGetAuditCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s CategoryGetAuditCategoriesResponse) GoString() string {
	return s.String()
}

func (s *CategoryGetAuditCategoriesResponse) SetErrMsg(v string) *CategoryGetAuditCategoriesResponse {
	s.ErrMsg = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponse) SetLogId(v string) *CategoryGetAuditCategoriesResponse {
	s.LogId = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponse) SetData(v *CategoryGetAuditCategoriesResponseData) *CategoryGetAuditCategoriesResponse {
	s.Data = v
	return s
}

func (s *CategoryGetAuditCategoriesResponse) SetErrNo(v int32) *CategoryGetAuditCategoriesResponse {
	s.ErrNo = &v
	return s
}

type CategoryGetAuditCategoriesResponseData struct {
	CategoryLimit        *int64                                                            `json:"category_limit,omitempty" xml:"category_limit,omitempty" require:"true"`
	CategoryChangeLimit  *int64                                                            `json:"category_change_limit,omitempty" xml:"category_change_limit,omitempty" require:"true"`
	AppCategoryAuditInfo []*CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem `json:"app_category_audit_info,omitempty" xml:"app_category_audit_info,omitempty" require:"true" type:"Repeated"`
	RemainingTimes       *int64                                                            `json:"remaining_times,omitempty" xml:"remaining_times,omitempty" require:"true"`
}

func (s CategoryGetAuditCategoriesResponseData) String() string {
	return tea.Prettify(s)
}

func (s CategoryGetAuditCategoriesResponseData) GoString() string {
	return s.String()
}

func (s *CategoryGetAuditCategoriesResponseData) SetCategoryLimit(v int64) *CategoryGetAuditCategoriesResponseData {
	s.CategoryLimit = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseData) SetCategoryChangeLimit(v int64) *CategoryGetAuditCategoriesResponseData {
	s.CategoryChangeLimit = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseData) SetAppCategoryAuditInfo(v []*CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) *CategoryGetAuditCategoriesResponseData {
	s.AppCategoryAuditInfo = v
	return s
}

func (s *CategoryGetAuditCategoriesResponseData) SetRemainingTimes(v int64) *CategoryGetAuditCategoriesResponseData {
	s.RemainingTimes = &v
	return s
}

type CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem struct {
	Reason             *string                                                                           `json:"reason,omitempty" xml:"reason,omitempty" require:"true"`
	CategoryAuditState *int                                                                              `json:"category_audit_state,omitempty" xml:"category_audit_state,omitempty" require:"true"`
	Category           *string                                                                           `json:"category,omitempty" xml:"category,omitempty" require:"true"`
	CategoryCertIds    *string                                                                           `json:"category_cert_ids,omitempty" xml:"category_cert_ids,omitempty" require:"true"`
	CategoryName       *string                                                                           `json:"category_name,omitempty" xml:"category_name,omitempty" require:"true"`
	CategoryCert       []*CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem `json:"category_cert,omitempty" xml:"category_cert,omitempty" require:"true" type:"Repeated"`
}

func (s CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) String() string {
	return tea.Prettify(s)
}

func (s CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) GoString() string {
	return s.String()
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetReason(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.Reason = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetCategoryAuditState(v int) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.CategoryAuditState = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetCategory(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.Category = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetCategoryCertIds(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.CategoryCertIds = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetCategoryName(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.CategoryName = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem) SetCategoryCert(v []*CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItem {
	s.CategoryCert = v
	return s
}

type CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem struct {
	IsVideo  *int32    `json:"is_video,omitempty" xml:"is_video,omitempty"`
	Id       *string   `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	Name     *string   `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	PathList []*string `json:"path_list,omitempty" xml:"path_list,omitempty" type:"Repeated"`
}

func (s CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) String() string {
	return tea.Prettify(s)
}

func (s CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) GoString() string {
	return s.String()
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) SetIsVideo(v int32) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem {
	s.IsVideo = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) SetId(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem {
	s.Id = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) SetName(v string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem {
	s.Name = &v
	return s
}

func (s *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem) SetPathList(v []*string) *CategoryGetAuditCategoriesResponseDataAppCategoryAuditInfoItemCategoryCertItem {
	s.PathList = v
	return s
}

type CensorImageRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Image       *string            `json:"image,omitempty" xml:"image,omitempty"`
	ImageData   *string            `json:"image_data,omitempty" xml:"image_data,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CensorImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CensorImageRequest) GoString() string {
	return s.String()
}

func (s *CensorImageRequest) SetAccessToken(v string) *CensorImageRequest {
	s.AccessToken = &v
	return s
}

func (s *CensorImageRequest) SetImage(v string) *CensorImageRequest {
	s.Image = &v
	return s
}

func (s *CensorImageRequest) SetImageData(v string) *CensorImageRequest {
	s.ImageData = &v
	return s
}

func (s *CensorImageRequest) SetAppId(v string) *CensorImageRequest {
	s.AppId = &v
	return s
}

func (s *CensorImageRequest) SetHeader(v map[string]*string) *CensorImageRequest {
	s.Header = v
	return s
}

type CensorImageResponse struct {
	ErrMsg   *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId    *string                            `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Predicts []*CensorImageResponsePredictsItem `json:"predicts,omitempty" xml:"predicts,omitempty" type:"Repeated"`
	ErrNo    *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s CensorImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CensorImageResponse) GoString() string {
	return s.String()
}

func (s *CensorImageResponse) SetErrMsg(v string) *CensorImageResponse {
	s.ErrMsg = &v
	return s
}

func (s *CensorImageResponse) SetLogId(v string) *CensorImageResponse {
	s.LogId = &v
	return s
}

func (s *CensorImageResponse) SetPredicts(v []*CensorImageResponsePredictsItem) *CensorImageResponse {
	s.Predicts = v
	return s
}

func (s *CensorImageResponse) SetErrNo(v int32) *CensorImageResponse {
	s.ErrNo = &v
	return s
}

type CensorImageResponsePredictsItem struct {
	Hit       *bool   `json:"hit,omitempty" xml:"hit,omitempty"`
	ModelName *string `json:"model_name,omitempty" xml:"model_name,omitempty"`
}

func (s CensorImageResponsePredictsItem) String() string {
	return tea.Prettify(s)
}

func (s CensorImageResponsePredictsItem) GoString() string {
	return s.String()
}

func (s *CensorImageResponsePredictsItem) SetHit(v bool) *CensorImageResponsePredictsItem {
	s.Hit = &v
	return s
}

func (s *CensorImageResponsePredictsItem) SetModelName(v string) *CensorImageResponsePredictsItem {
	s.ModelName = &v
	return s
}

type CertInfoRequest struct {
	MerchantLifeAccountId *int64             `json:"merchant_life_account_id,omitempty" xml:"merchant_life_account_id,omitempty"`
	PoiId                 *int64             `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Header                map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CertInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CertInfoRequest) GoString() string {
	return s.String()
}

func (s *CertInfoRequest) SetMerchantLifeAccountId(v int64) *CertInfoRequest {
	s.MerchantLifeAccountId = &v
	return s
}

func (s *CertInfoRequest) SetPoiId(v int64) *CertInfoRequest {
	s.PoiId = &v
	return s
}

func (s *CertInfoRequest) SetHeader(v map[string]*string) *CertInfoRequest {
	s.Header = v
	return s
}

func (s *CertInfoRequest) SetAccessToken(v string) *CertInfoRequest {
	s.AccessToken = &v
	return s
}

type CertInfoResponse struct {
	Extra *CertInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *CertInfoResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s CertInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponse) GoString() string {
	return s.String()
}

func (s *CertInfoResponse) SetExtra(v *CertInfoResponseExtra) *CertInfoResponse {
	s.Extra = v
	return s
}

func (s *CertInfoResponse) SetData(v *CertInfoResponseData) *CertInfoResponse {
	s.Data = v
	return s
}

type CertInfoResponseData struct {
	Industry      *CertInfoResponseDataIndustry `json:"industry,omitempty" xml:"industry,omitempty"`
	PoiId         *int64                        `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Subject       *CertInfoResponseDataSubject  `json:"subject,omitempty" xml:"subject,omitempty"`
	GwErrorCode   *int32                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	AccountName   *string                       `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s CertInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponseData) GoString() string {
	return s.String()
}

func (s *CertInfoResponseData) SetIndustry(v *CertInfoResponseDataIndustry) *CertInfoResponseData {
	s.Industry = v
	return s
}

func (s *CertInfoResponseData) SetPoiId(v int64) *CertInfoResponseData {
	s.PoiId = &v
	return s
}

func (s *CertInfoResponseData) SetSubject(v *CertInfoResponseDataSubject) *CertInfoResponseData {
	s.Subject = v
	return s
}

func (s *CertInfoResponseData) SetGwErrorCode(v int32) *CertInfoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertInfoResponseData) SetGwDescription(v string) *CertInfoResponseData {
	s.GwDescription = &v
	return s
}

func (s *CertInfoResponseData) SetAccountName(v string) *CertInfoResponseData {
	s.AccountName = &v
	return s
}

type CertInfoResponseDataIndustry struct {
	MainCateroryCodes []*string                                `json:"main_caterory_codes,omitempty" xml:"main_caterory_codes,omitempty" type:"Repeated"`
	Quals             []*CertInfoResponseDataIndustryQualsItem `json:"quals,omitempty" xml:"quals,omitempty" type:"Repeated"`
	SubCateroryCodes  []*string                                `json:"sub_caterory_codes,omitempty" xml:"sub_caterory_codes,omitempty" type:"Repeated"`
}

func (s CertInfoResponseDataIndustry) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponseDataIndustry) GoString() string {
	return s.String()
}

func (s *CertInfoResponseDataIndustry) SetMainCateroryCodes(v []*string) *CertInfoResponseDataIndustry {
	s.MainCateroryCodes = v
	return s
}

func (s *CertInfoResponseDataIndustry) SetQuals(v []*CertInfoResponseDataIndustryQualsItem) *CertInfoResponseDataIndustry {
	s.Quals = v
	return s
}

func (s *CertInfoResponseDataIndustry) SetSubCateroryCodes(v []*string) *CertInfoResponseDataIndustry {
	s.SubCateroryCodes = v
	return s
}

type CertInfoResponseDataIndustryQualsItem struct {
	Urls           []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	ExpirationDate *string   `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
	Type           *int64    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CertInfoResponseDataIndustryQualsItem) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponseDataIndustryQualsItem) GoString() string {
	return s.String()
}

func (s *CertInfoResponseDataIndustryQualsItem) SetUrls(v []*string) *CertInfoResponseDataIndustryQualsItem {
	s.Urls = v
	return s
}

func (s *CertInfoResponseDataIndustryQualsItem) SetExpirationDate(v string) *CertInfoResponseDataIndustryQualsItem {
	s.ExpirationDate = &v
	return s
}

func (s *CertInfoResponseDataIndustryQualsItem) SetType(v int64) *CertInfoResponseDataIndustryQualsItem {
	s.Type = &v
	return s
}

type CertInfoResponseDataSubject struct {
	LegalPersonName *string   `json:"legal_person_name,omitempty" xml:"legal_person_name,omitempty"`
	LicenseId       *string   `json:"license_id,omitempty" xml:"license_id,omitempty"`
	LicenseUrls     []*string `json:"license_urls,omitempty" xml:"license_urls,omitempty" type:"Repeated"`
	QualType        *int64    `json:"qual_type,omitempty" xml:"qual_type,omitempty"`
	CompanyName     *string   `json:"company_name,omitempty" xml:"company_name,omitempty"`
	ExpirationDate  *string   `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

func (s CertInfoResponseDataSubject) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponseDataSubject) GoString() string {
	return s.String()
}

func (s *CertInfoResponseDataSubject) SetLegalPersonName(v string) *CertInfoResponseDataSubject {
	s.LegalPersonName = &v
	return s
}

func (s *CertInfoResponseDataSubject) SetLicenseId(v string) *CertInfoResponseDataSubject {
	s.LicenseId = &v
	return s
}

func (s *CertInfoResponseDataSubject) SetLicenseUrls(v []*string) *CertInfoResponseDataSubject {
	s.LicenseUrls = v
	return s
}

func (s *CertInfoResponseDataSubject) SetQualType(v int64) *CertInfoResponseDataSubject {
	s.QualType = &v
	return s
}

func (s *CertInfoResponseDataSubject) SetCompanyName(v string) *CertInfoResponseDataSubject {
	s.CompanyName = &v
	return s
}

func (s *CertInfoResponseDataSubject) SetExpirationDate(v string) *CertInfoResponseDataSubject {
	s.ExpirationDate = &v
	return s
}

type CertInfoResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CertInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *CertInfoResponseExtra) SetErrorCode(v int32) *CertInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertInfoResponseExtra) SetLogid(v string) *CertInfoResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertInfoResponseExtra) SetNow(v int64) *CertInfoResponseExtra {
	s.Now = &v
	return s
}

func (s *CertInfoResponseExtra) SetSubDescription(v string) *CertInfoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertInfoResponseExtra) SetSubErrorCode(v int32) *CertInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertInfoResponseExtra) SetDescription(v string) *CertInfoResponseExtra {
	s.Description = &v
	return s
}

type CertificateBatchVerifyRequest struct {
	CertificateBatchVerifyList []*CertificateBatchVerifyRequestCertificateBatchVerifyListItem `json:"certificate_batch_verify_list,omitempty" xml:"certificate_batch_verify_list,omitempty" require:"true" type:"Repeated"`
	PoiId                      *string                                                        `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	Header                     map[string]*string                                             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken                *string                                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CertificateBatchVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequest) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequest) SetCertificateBatchVerifyList(v []*CertificateBatchVerifyRequestCertificateBatchVerifyListItem) *CertificateBatchVerifyRequest {
	s.CertificateBatchVerifyList = v
	return s
}

func (s *CertificateBatchVerifyRequest) SetPoiId(v string) *CertificateBatchVerifyRequest {
	s.PoiId = &v
	return s
}

func (s *CertificateBatchVerifyRequest) SetHeader(v map[string]*string) *CertificateBatchVerifyRequest {
	s.Header = v
	return s
}

func (s *CertificateBatchVerifyRequest) SetAccessToken(v string) *CertificateBatchVerifyRequest {
	s.AccessToken = &v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItem struct {
	OrderId          *string                                                                            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyExtra      *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra            `json:"verify_extra,omitempty" xml:"verify_extra,omitempty"`
	VerifySignList   []*string                                                                          `json:"verify_sign_list,omitempty" xml:"verify_sign_list,omitempty" type:"Repeated"`
	VerifyToken      *string                                                                            `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	CodeWithTimeList []*CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem `json:"code_with_time_list,omitempty" xml:"code_with_time_list,omitempty" type:"Repeated"`
	EncryptedCodes   []*string                                                                          `json:"encrypted_codes,omitempty" xml:"encrypted_codes,omitempty" type:"Repeated"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItem) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetOrderId(v string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.OrderId = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetVerifyExtra(v *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.VerifyExtra = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetVerifySignList(v []*string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.VerifySignList = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetVerifyToken(v string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.VerifyToken = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetCodeWithTimeList(v []*CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.CodeWithTimeList = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItem) SetEncryptedCodes(v []*string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItem {
	s.EncryptedCodes = v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem struct {
	OuterNumb  *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb `json:"outer_numb,omitempty" xml:"outer_numb,omitempty"`
	SerialNum  *int32                                                                                    `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	VerifyTime *int64                                                                                    `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	Code       *string                                                                                   `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) SetOuterNumb(v *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem {
	s.OuterNumb = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) SetSerialNum(v int32) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem {
	s.SerialNum = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) SetVerifyTime(v int64) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem) SetCode(v string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItem {
	s.Code = &v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb struct {
	OrderNumber  *string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	CouponNumber *string `json:"coupon_number,omitempty" xml:"coupon_number,omitempty"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb) SetOrderNumber(v string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb {
	s.OrderNumber = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb) SetCouponNumber(v string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemCodeWithTimeListItemOuterNumb {
	s.CouponNumber = &v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra struct {
	DynamicCouponInfo   *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo   `json:"dynamic_coupon_info,omitempty" xml:"dynamic_coupon_info,omitempty"`
	OfflineAddPriceInfo *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo `json:"offline_add_price_info,omitempty" xml:"offline_add_price_info,omitempty"`
	OutGoodIds          []*string                                                                                  `json:"out_good_ids,omitempty" xml:"out_good_ids,omitempty" type:"Repeated"`
	TotalVerify         *bool                                                                                      `json:"total_verify,omitempty" xml:"total_verify,omitempty"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) SetDynamicCouponInfo(v *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra {
	s.DynamicCouponInfo = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) SetOfflineAddPriceInfo(v *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra {
	s.OfflineAddPriceInfo = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) SetOutGoodIds(v []*string) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra {
	s.OutGoodIds = v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra) SetTotalVerify(v bool) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtra {
	s.TotalVerify = &v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo struct {
	ActualDeductionAmount *int64 `json:"actual_deduction_amount,omitempty" xml:"actual_deduction_amount,omitempty"`
	BizTime               *int64 `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo) SetActualDeductionAmount(v int64) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo {
	s.ActualDeductionAmount = &v
	return s
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo) SetBizTime(v int64) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraDynamicCouponInfo {
	s.BizTime = &v
	return s
}

type CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo struct {
	IsOfflineAddPrice *bool `json:"is_offline_add_price,omitempty" xml:"is_offline_add_price,omitempty"`
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo) SetIsOfflineAddPrice(v bool) *CertificateBatchVerifyRequestCertificateBatchVerifyListItemVerifyExtraOfflineAddPriceInfo {
	s.IsOfflineAddPrice = &v
	return s
}

type CertificateBatchVerifyResponse struct {
	Extra                     *CertificateBatchVerifyResponseExtra                           `json:"extra,omitempty" xml:"extra,omitempty"`
	Data                      *CertificateBatchVerifyResponseData                            `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	CertificateVerifyRespList []*CertificateBatchVerifyResponseCertificateVerifyRespListItem `json:"certificate_verify_resp_list,omitempty" xml:"certificate_verify_resp_list,omitempty" require:"true" type:"Repeated"`
}

func (s CertificateBatchVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponse) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponse) SetExtra(v *CertificateBatchVerifyResponseExtra) *CertificateBatchVerifyResponse {
	s.Extra = v
	return s
}

func (s *CertificateBatchVerifyResponse) SetData(v *CertificateBatchVerifyResponseData) *CertificateBatchVerifyResponse {
	s.Data = v
	return s
}

func (s *CertificateBatchVerifyResponse) SetCertificateVerifyRespList(v []*CertificateBatchVerifyResponseCertificateVerifyRespListItem) *CertificateBatchVerifyResponse {
	s.CertificateVerifyRespList = v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItem struct {
	Data  *CertificateBatchVerifyResponseCertificateVerifyRespListItemData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItem) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItem) SetData(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemData) *CertificateBatchVerifyResponseCertificateVerifyRespListItem {
	s.Data = v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItem) SetExtra(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) *CertificateBatchVerifyResponseCertificateVerifyRespListItem {
	s.Extra = v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemData struct {
	VerifyResults []*CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemData) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemData) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemData) SetVerifyResults(v []*CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) *CertificateBatchVerifyResponseCertificateVerifyRespListItemData {
	s.VerifyResults = v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem struct {
	AccountId        *string                                                                                           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	VerifyAmountInfo *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo `json:"verify_amount_info,omitempty" xml:"verify_amount_info,omitempty"`
	Result           *int32                                                                                            `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	VerifyId         *string                                                                                           `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	OrderId          *string                                                                                           `json:"order_id,omitempty" xml:"order_id,omitempty"`
	IdCard           *string                                                                                           `json:"id_card,omitempty" xml:"id_card,omitempty"`
	OriginCode       *string                                                                                           `json:"origin_code,omitempty" xml:"origin_code,omitempty"`
	CertificateId    *string                                                                                           `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Msg              *string                                                                                           `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	CertificateNo    *string                                                                                           `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	ProductId        *string                                                                                           `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Qrcode           *string                                                                                           `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
	OutGoodId        *string                                                                                           `json:"out_good_id,omitempty" xml:"out_good_id,omitempty"`
	Code             *string                                                                                           `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetAccountId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.AccountId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetVerifyAmountInfo(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.VerifyAmountInfo = v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetResult(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.Result = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetVerifyId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetOrderId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.OrderId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetIdCard(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.IdCard = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetOriginCode(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.OriginCode = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetCertificateId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetMsg(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.Msg = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetCertificateNo(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.CertificateNo = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetProductId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.ProductId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetQrcode(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.Qrcode = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetOutGoodId(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.OutGoodId = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem) SetCode(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItem {
	s.Code = &v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo struct {
	TimesCardSerialAmount *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount `json:"times_card_serial_amount,omitempty" xml:"times_card_serial_amount,omitempty"`
	TimeCardAmount        *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount        `json:"time_card_amount,omitempty" xml:"time_card_amount,omitempty"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo) SetTimesCardSerialAmount(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo {
	s.TimesCardSerialAmount = v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo) SetTimeCardAmount(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfo {
	s.TimeCardAmount = v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount) SetAmount(v int64) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimeCardAmount {
	s.Amount = &v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount struct {
	SerialNumb *int32                                                                                                                       `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
	Amount     *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) SetSerialNumb(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount {
	s.SerialNumb = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount) SetAmount(v *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmount {
	s.Amount = v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount struct {
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetCouponPayAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPayAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalCurrency(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPaymentDiscountAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPlatformDiscountAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetMerchantTicketAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginListMarketAmount(v int64) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetListMarketAmount(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetBrandTicketAmount(v int64) *CertificateBatchVerifyResponseCertificateVerifyRespListItemDataVerifyResultsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.BrandTicketAmount = &v
	return s
}

type CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetSubErrorCode(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetDescription(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.Description = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetErrorCode(v int32) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetLogid(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.Logid = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetNow(v int64) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.Now = &v
	return s
}

func (s *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra) SetSubDescription(v string) *CertificateBatchVerifyResponseCertificateVerifyRespListItemExtra {
	s.SubDescription = &v
	return s
}

type CertificateBatchVerifyResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CertificateBatchVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseData) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseData) SetGwDescription(v string) *CertificateBatchVerifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *CertificateBatchVerifyResponseData) SetGwErrorCode(v int32) *CertificateBatchVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

type CertificateBatchVerifyResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s CertificateBatchVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateBatchVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateBatchVerifyResponseExtra) SetNow(v int64) *CertificateBatchVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificateBatchVerifyResponseExtra) SetSubDescription(v string) *CertificateBatchVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateBatchVerifyResponseExtra) SetSubErrorCode(v int32) *CertificateBatchVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificateBatchVerifyResponseExtra) SetDescription(v string) *CertificateBatchVerifyResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificateBatchVerifyResponseExtra) SetErrorCode(v int32) *CertificateBatchVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateBatchVerifyResponseExtra) SetLogid(v string) *CertificateBatchVerifyResponseExtra {
	s.Logid = &v
	return s
}

type CertificateCallbackRequest struct {
	AccessToken         *string                                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Result              *int                                                 `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	CertificateInfoList []*CertificateCallbackRequestCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	OrderId             *string                                              `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header              map[string]*string                                   `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CertificateCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackRequest) GoString() string {
	return s.String()
}

func (s *CertificateCallbackRequest) SetAccessToken(v string) *CertificateCallbackRequest {
	s.AccessToken = &v
	return s
}

func (s *CertificateCallbackRequest) SetResult(v int) *CertificateCallbackRequest {
	s.Result = &v
	return s
}

func (s *CertificateCallbackRequest) SetCertificateInfoList(v []*CertificateCallbackRequestCertificateInfoListItem) *CertificateCallbackRequest {
	s.CertificateInfoList = v
	return s
}

func (s *CertificateCallbackRequest) SetOrderId(v string) *CertificateCallbackRequest {
	s.OrderId = &v
	return s
}

func (s *CertificateCallbackRequest) SetHeader(v map[string]*string) *CertificateCallbackRequest {
	s.Header = v
	return s
}

type CertificateCallbackRequestCertificateInfoListItem struct {
	CertificateId *int64                                                              `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	ProjectList   []*CertificateCallbackRequestCertificateInfoListItemProjectListItem `json:"project_list,omitempty" xml:"project_list,omitempty" require:"true" type:"Repeated"`
}

func (s CertificateCallbackRequestCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackRequestCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateCallbackRequestCertificateInfoListItem) SetCertificateId(v int64) *CertificateCallbackRequestCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItem) SetProjectList(v []*CertificateCallbackRequestCertificateInfoListItemProjectListItem) *CertificateCallbackRequestCertificateInfoListItem {
	s.ProjectList = v
	return s
}

type CertificateCallbackRequestCertificateInfoListItemProjectListItem struct {
	CertificateList []*CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem `json:"certificate_list,omitempty" xml:"certificate_list,omitempty" require:"true" type:"Repeated"`
	CredentialList  []*CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem  `json:"credential_list,omitempty" xml:"credential_list,omitempty" type:"Repeated"`
	Name            *string                                                                                `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	ProjectId       *string                                                                                `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItem) GoString() string {
	return s.String()
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItem) SetCertificateList(v []*CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem) *CertificateCallbackRequestCertificateInfoListItemProjectListItem {
	s.CertificateList = v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItem) SetCredentialList(v []*CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem) *CertificateCallbackRequestCertificateInfoListItemProjectListItem {
	s.CredentialList = v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItem) SetName(v string) *CertificateCallbackRequestCertificateInfoListItemProjectListItem {
	s.Name = &v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItem) SetProjectId(v string) *CertificateCallbackRequestCertificateInfoListItemProjectListItem {
	s.ProjectId = &v
	return s
}

type CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem struct {
	CcertificateType *int    `json:"ccertificate_type,omitempty" xml:"ccertificate_type,omitempty" require:"true"`
	CertificateNo    *string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty" require:"true"`
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem) GoString() string {
	return s.String()
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem) SetCcertificateType(v int) *CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem {
	s.CcertificateType = &v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem) SetCertificateNo(v string) *CertificateCallbackRequestCertificateInfoListItemProjectListItemCertificateListItem {
	s.CertificateNo = &v
	return s
}

type CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem) GoString() string {
	return s.String()
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem) SetCredentialNo(v string) *CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem {
	s.CredentialNo = &v
	return s
}

func (s *CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem) SetCredentialType(v int) *CertificateCallbackRequestCertificateInfoListItemProjectListItemCredentialListItem {
	s.CredentialType = &v
	return s
}

type CertificateCallbackResponse struct {
	Extra               *CertificateCallbackResponseExtra                     `json:"extra,omitempty" xml:"extra,omitempty"`
	Data                *CertificateCallbackResponseData                      `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	CertificateInfoList []*CertificateCallbackResponseCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
}

func (s CertificateCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackResponse) GoString() string {
	return s.String()
}

func (s *CertificateCallbackResponse) SetExtra(v *CertificateCallbackResponseExtra) *CertificateCallbackResponse {
	s.Extra = v
	return s
}

func (s *CertificateCallbackResponse) SetData(v *CertificateCallbackResponseData) *CertificateCallbackResponse {
	s.Data = v
	return s
}

func (s *CertificateCallbackResponse) SetCertificateInfoList(v []*CertificateCallbackResponseCertificateInfoListItem) *CertificateCallbackResponse {
	s.CertificateInfoList = v
	return s
}

type CertificateCallbackResponseCertificateInfoListItem struct {
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
}

func (s CertificateCallbackResponseCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackResponseCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateCallbackResponseCertificateInfoListItem) SetCode(v string) *CertificateCallbackResponseCertificateInfoListItem {
	s.Code = &v
	return s
}

func (s *CertificateCallbackResponseCertificateInfoListItem) SetCertificateId(v string) *CertificateCallbackResponseCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

type CertificateCallbackResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CertificateCallbackResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackResponseData) GoString() string {
	return s.String()
}

func (s *CertificateCallbackResponseData) SetGwErrorCode(v int32) *CertificateCallbackResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertificateCallbackResponseData) SetGwDescription(v string) *CertificateCallbackResponseData {
	s.GwDescription = &v
	return s
}

type CertificateCallbackResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CertificateCallbackResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateCallbackResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateCallbackResponseExtra) SetDescription(v string) *CertificateCallbackResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificateCallbackResponseExtra) SetErrorCode(v int32) *CertificateCallbackResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateCallbackResponseExtra) SetLogid(v string) *CertificateCallbackResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertificateCallbackResponseExtra) SetNow(v int64) *CertificateCallbackResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificateCallbackResponseExtra) SetSubDescription(v string) *CertificateCallbackResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateCallbackResponseExtra) SetSubErrorCode(v int32) *CertificateCallbackResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CertificateCancelRequest struct {
	CancelToken          *string                                            `json:"cancel_token,omitempty" xml:"cancel_token,omitempty"`
	CertificateId        *string                                            `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Header               map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId            *string                                            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BatchCancelInfo      *CertificateCancelRequestBatchCancelInfo           `json:"batch_cancel_info,omitempty" xml:"batch_cancel_info,omitempty"`
	BatchCancelInfoList  []*CertificateCancelRequestBatchCancelInfoListItem `json:"batch_cancel_info_list,omitempty" xml:"batch_cancel_info_list,omitempty" type:"Repeated"`
	ShopOrderId          *string                                            `json:"shop_order_id,omitempty" xml:"shop_order_id,omitempty"`
	TimesCardCancelCount *int64                                             `json:"times_card_cancel_count,omitempty" xml:"times_card_cancel_count,omitempty"`
	VerifyId             *string                                            `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
}

func (s CertificateCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelRequest) GoString() string {
	return s.String()
}

func (s *CertificateCancelRequest) SetCancelToken(v string) *CertificateCancelRequest {
	s.CancelToken = &v
	return s
}

func (s *CertificateCancelRequest) SetCertificateId(v string) *CertificateCancelRequest {
	s.CertificateId = &v
	return s
}

func (s *CertificateCancelRequest) SetHeader(v map[string]*string) *CertificateCancelRequest {
	s.Header = v
	return s
}

func (s *CertificateCancelRequest) SetAccessToken(v string) *CertificateCancelRequest {
	s.AccessToken = &v
	return s
}

func (s *CertificateCancelRequest) SetAccountId(v string) *CertificateCancelRequest {
	s.AccountId = &v
	return s
}

func (s *CertificateCancelRequest) SetBatchCancelInfo(v *CertificateCancelRequestBatchCancelInfo) *CertificateCancelRequest {
	s.BatchCancelInfo = v
	return s
}

func (s *CertificateCancelRequest) SetBatchCancelInfoList(v []*CertificateCancelRequestBatchCancelInfoListItem) *CertificateCancelRequest {
	s.BatchCancelInfoList = v
	return s
}

func (s *CertificateCancelRequest) SetShopOrderId(v string) *CertificateCancelRequest {
	s.ShopOrderId = &v
	return s
}

func (s *CertificateCancelRequest) SetTimesCardCancelCount(v int64) *CertificateCancelRequest {
	s.TimesCardCancelCount = &v
	return s
}

func (s *CertificateCancelRequest) SetVerifyId(v string) *CertificateCancelRequest {
	s.VerifyId = &v
	return s
}

type CertificateCancelRequestBatchCancelInfo struct {
	OrderId      *string   `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyIdList []*string `json:"verify_id_list,omitempty" xml:"verify_id_list,omitempty" type:"Repeated"`
}

func (s CertificateCancelRequestBatchCancelInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelRequestBatchCancelInfo) GoString() string {
	return s.String()
}

func (s *CertificateCancelRequestBatchCancelInfo) SetOrderId(v string) *CertificateCancelRequestBatchCancelInfo {
	s.OrderId = &v
	return s
}

func (s *CertificateCancelRequestBatchCancelInfo) SetVerifyIdList(v []*string) *CertificateCancelRequestBatchCancelInfo {
	s.VerifyIdList = v
	return s
}

type CertificateCancelRequestBatchCancelInfoListItem struct {
	OrderId      *string   `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyIdList []*string `json:"verify_id_list,omitempty" xml:"verify_id_list,omitempty" type:"Repeated"`
}

func (s CertificateCancelRequestBatchCancelInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelRequestBatchCancelInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateCancelRequestBatchCancelInfoListItem) SetOrderId(v string) *CertificateCancelRequestBatchCancelInfoListItem {
	s.OrderId = &v
	return s
}

func (s *CertificateCancelRequestBatchCancelInfoListItem) SetVerifyIdList(v []*string) *CertificateCancelRequestBatchCancelInfoListItem {
	s.VerifyIdList = v
	return s
}

type CertificateCancelResponse struct {
	Extra *CertificateCancelResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *CertificateCancelResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CertificateCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelResponse) GoString() string {
	return s.String()
}

func (s *CertificateCancelResponse) SetExtra(v *CertificateCancelResponseExtra) *CertificateCancelResponse {
	s.Extra = v
	return s
}

func (s *CertificateCancelResponse) SetData(v *CertificateCancelResponseData) *CertificateCancelResponse {
	s.Data = v
	return s
}

type CertificateCancelResponseData struct {
	TransactionId *int64                                            `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`
	CancelResults []*CertificateCancelResponseDataCancelResultsItem `json:"cancel_results,omitempty" xml:"cancel_results,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CertificateCancelResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelResponseData) GoString() string {
	return s.String()
}

func (s *CertificateCancelResponseData) SetTransactionId(v int64) *CertificateCancelResponseData {
	s.TransactionId = &v
	return s
}

func (s *CertificateCancelResponseData) SetCancelResults(v []*CertificateCancelResponseDataCancelResultsItem) *CertificateCancelResponseData {
	s.CancelResults = v
	return s
}

func (s *CertificateCancelResponseData) SetGwErrorCode(v int32) *CertificateCancelResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertificateCancelResponseData) SetGwDescription(v string) *CertificateCancelResponseData {
	s.GwDescription = &v
	return s
}

type CertificateCancelResponseDataCancelResultsItem struct {
	ResultCode *int32  `json:"result_code,omitempty" xml:"result_code,omitempty" require:"true"`
	ResultMsg  *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	VerifyId   *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s CertificateCancelResponseDataCancelResultsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelResponseDataCancelResultsItem) GoString() string {
	return s.String()
}

func (s *CertificateCancelResponseDataCancelResultsItem) SetResultCode(v int32) *CertificateCancelResponseDataCancelResultsItem {
	s.ResultCode = &v
	return s
}

func (s *CertificateCancelResponseDataCancelResultsItem) SetResultMsg(v string) *CertificateCancelResponseDataCancelResultsItem {
	s.ResultMsg = &v
	return s
}

func (s *CertificateCancelResponseDataCancelResultsItem) SetVerifyId(v string) *CertificateCancelResponseDataCancelResultsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateCancelResponseDataCancelResultsItem) SetOrderId(v string) *CertificateCancelResponseDataCancelResultsItem {
	s.OrderId = &v
	return s
}

type CertificateCancelResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s CertificateCancelResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateCancelResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateCancelResponseExtra) SetLogid(v string) *CertificateCancelResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertificateCancelResponseExtra) SetNow(v int64) *CertificateCancelResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificateCancelResponseExtra) SetSubDescription(v string) *CertificateCancelResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateCancelResponseExtra) SetSubErrorCode(v int32) *CertificateCancelResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificateCancelResponseExtra) SetDescription(v string) *CertificateCancelResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificateCancelResponseExtra) SetErrorCode(v int32) *CertificateCancelResponseExtra {
	s.ErrorCode = &v
	return s
}

type CertificateGetRequest struct {
	EncryptedCode *string            `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CertificateGetRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetRequest) GoString() string {
	return s.String()
}

func (s *CertificateGetRequest) SetEncryptedCode(v string) *CertificateGetRequest {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateGetRequest) SetAccountId(v string) *CertificateGetRequest {
	s.AccountId = &v
	return s
}

func (s *CertificateGetRequest) SetHeader(v map[string]*string) *CertificateGetRequest {
	s.Header = v
	return s
}

func (s *CertificateGetRequest) SetAccessToken(v string) *CertificateGetRequest {
	s.AccessToken = &v
	return s
}

type CertificateGetResponse struct {
	Data  *CertificateGetResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *CertificateGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s CertificateGetResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponse) GoString() string {
	return s.String()
}

func (s *CertificateGetResponse) SetData(v *CertificateGetResponseData) *CertificateGetResponse {
	s.Data = v
	return s
}

func (s *CertificateGetResponse) SetExtra(v *CertificateGetResponseExtra) *CertificateGetResponse {
	s.Extra = v
	return s
}

type CertificateGetResponseData struct {
	Certificate   *CertificateGetResponseDataCertificate   `json:"certificate,omitempty" xml:"certificate,omitempty"`
	CertificateV2 *CertificateGetResponseDataCertificateV2 `json:"certificate_v2,omitempty" xml:"certificate_v2,omitempty"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s CertificateGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseData) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseData) SetCertificate(v *CertificateGetResponseDataCertificate) *CertificateGetResponseData {
	s.Certificate = v
	return s
}

func (s *CertificateGetResponseData) SetCertificateV2(v *CertificateGetResponseDataCertificateV2) *CertificateGetResponseData {
	s.CertificateV2 = v
	return s
}

func (s *CertificateGetResponseData) SetGwErrorCode(v int32) *CertificateGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertificateGetResponseData) SetGwDescription(v string) *CertificateGetResponseData {
	s.GwDescription = &v
	return s
}

type CertificateGetResponseDataCertificate struct {
	BookInfo             *CertificateGetResponseDataCertificateBookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	TimeCard             *CertificateGetResponseDataCertificateTimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	StartTime            *int64                                                     `json:"start_time,omitempty" xml:"start_time,omitempty"`
	PeriodCard           *CertificateGetResponseDataCertificatePeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	CertificateId        *int64                                                     `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Sku                  *CertificateGetResponseDataCertificateSku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	NotAvailablePoiList  []*string                                                  `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	UseTimeInfo          *CertificateGetResponseDataCertificateUseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	Status               *int                                                       `json:"status,omitempty" xml:"status,omitempty"`
	VerifyRecords        []*CertificateGetResponseDataCertificateVerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	ReserveInfo          *CertificateGetResponseDataCertificateReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	Code                 *string                                                    `json:"code,omitempty" xml:"code,omitempty"`
	OffPeakDiscountInfo  *CertificateGetResponseDataCertificateOffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
	EncryptedCode        *string                                                    `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	NotAvailableTimeInfo *CertificateGetResponseDataCertificateNotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	Amount               *CertificateGetResponseDataCertificateAmount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	AdditionalMap        map[int]*string                                            `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	UsedStatusType       *int                                                       `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
	ExpireTime           *int64                                                     `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Verify               *CertificateGetResponseDataCertificateVerify               `json:"verify,omitempty" xml:"verify,omitempty"`
}

func (s CertificateGetResponseDataCertificate) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificate) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificate) SetBookInfo(v *CertificateGetResponseDataCertificateBookInfo) *CertificateGetResponseDataCertificate {
	s.BookInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetTimeCard(v *CertificateGetResponseDataCertificateTimeCard) *CertificateGetResponseDataCertificate {
	s.TimeCard = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetStartTime(v int64) *CertificateGetResponseDataCertificate {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetPeriodCard(v *CertificateGetResponseDataCertificatePeriodCard) *CertificateGetResponseDataCertificate {
	s.PeriodCard = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetCertificateId(v int64) *CertificateGetResponseDataCertificate {
	s.CertificateId = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetSku(v *CertificateGetResponseDataCertificateSku) *CertificateGetResponseDataCertificate {
	s.Sku = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetNotAvailablePoiList(v []*string) *CertificateGetResponseDataCertificate {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetUseTimeInfo(v *CertificateGetResponseDataCertificateUseTimeInfo) *CertificateGetResponseDataCertificate {
	s.UseTimeInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetStatus(v int) *CertificateGetResponseDataCertificate {
	s.Status = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetVerifyRecords(v []*CertificateGetResponseDataCertificateVerifyRecordsItem) *CertificateGetResponseDataCertificate {
	s.VerifyRecords = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetReserveInfo(v *CertificateGetResponseDataCertificateReserveInfo) *CertificateGetResponseDataCertificate {
	s.ReserveInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetCode(v string) *CertificateGetResponseDataCertificate {
	s.Code = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetOffPeakDiscountInfo(v *CertificateGetResponseDataCertificateOffPeakDiscountInfo) *CertificateGetResponseDataCertificate {
	s.OffPeakDiscountInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetEncryptedCode(v string) *CertificateGetResponseDataCertificate {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetNotAvailableTimeInfo(v *CertificateGetResponseDataCertificateNotAvailableTimeInfo) *CertificateGetResponseDataCertificate {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetAmount(v *CertificateGetResponseDataCertificateAmount) *CertificateGetResponseDataCertificate {
	s.Amount = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetAdditionalMap(v map[int]*string) *CertificateGetResponseDataCertificate {
	s.AdditionalMap = v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetUsedStatusType(v int) *CertificateGetResponseDataCertificate {
	s.UsedStatusType = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetExpireTime(v int64) *CertificateGetResponseDataCertificate {
	s.ExpireTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificate) SetVerify(v *CertificateGetResponseDataCertificateVerify) *CertificateGetResponseDataCertificate {
	s.Verify = v
	return s
}

type CertificateGetResponseDataCertificateAmount struct {
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
}

func (s CertificateGetResponseDataCertificateAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateAmount) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateAmount) SetMerchantTicketAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetCouponPayAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetPayAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetPaymentDiscountAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetOriginalCurrency(v string) *CertificateGetResponseDataCertificateAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetListMarketAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetBrandTicketAmount(v int64) *CertificateGetResponseDataCertificateAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetOriginListMarketAmount(v int64) *CertificateGetResponseDataCertificateAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetOriginalAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateAmount) SetPlatformDiscountAmount(v int32) *CertificateGetResponseDataCertificateAmount {
	s.PlatformDiscountAmount = &v
	return s
}

type CertificateGetResponseDataCertificateBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
}

func (s CertificateGetResponseDataCertificateBookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateBookInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateBookInfo) SetBookPoiId(v string) *CertificateGetResponseDataCertificateBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateBookInfo) SetBookProductNumber(v int64) *CertificateGetResponseDataCertificateBookInfo {
	s.BookProductNumber = &v
	return s
}

type CertificateGetResponseDataCertificateNotAvailableTimeInfo struct {
	CanNoUseDate    []*CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
	CanNoUseWeekDay []*int64                                                                     `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
	FulfilEnable    *bool                                                                        `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
}

func (s CertificateGetResponseDataCertificateNotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateNotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateNotAvailableTimeInfo) SetCanNoUseDate(v []*CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem) *CertificateGetResponseDataCertificateNotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

func (s *CertificateGetResponseDataCertificateNotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificateGetResponseDataCertificateNotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

func (s *CertificateGetResponseDataCertificateNotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificateGetResponseDataCertificateNotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

type CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificateGetResponseDataCertificateNotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

type CertificateGetResponseDataCertificateOffPeakDiscountInfo struct {
	HasOffPeakDiscount *bool                                                                           `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
	IdleTimeLimitType  *int                                                                            `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	OffPeakTimeRange   []*CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificateGetResponseDataCertificateOffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificateGetResponseDataCertificateOffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) *CertificateGetResponseDataCertificateOffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

type CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem struct {
	EndTime            *int64                                                                                                `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

type CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificateGetResponseDataCertificateOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type CertificateGetResponseDataCertificatePeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificateGetResponseDataCertificatePeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificatePeriodCard) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificatePeriodCard) SetPeriodType(v int) *CertificateGetResponseDataCertificatePeriodCard {
	s.PeriodType = &v
	return s
}

type CertificateGetResponseDataCertificateReserveInfo struct {
	OrderReserveUserInfoList []*CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateReserveInfo) SetOrderReserveUserInfoList(v []*CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) *CertificateGetResponseDataCertificateReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem struct {
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}

func (s *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificateGetResponseDataCertificateReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

type CertificateGetResponseDataCertificateSku struct {
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
}

func (s CertificateGetResponseDataCertificateSku) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateSku) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateSku) SetTitle(v string) *CertificateGetResponseDataCertificateSku {
	s.Title = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetMarketPrice(v int64) *CertificateGetResponseDataCertificateSku {
	s.MarketPrice = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetProductId(v string) *CertificateGetResponseDataCertificateSku {
	s.ProductId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetAccountId(v string) *CertificateGetResponseDataCertificateSku {
	s.AccountId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetSuplierProductOutId(v string) *CertificateGetResponseDataCertificateSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetProductOutId(v string) *CertificateGetResponseDataCertificateSku {
	s.ProductOutId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetSkuId(v string) *CertificateGetResponseDataCertificateSku {
	s.SkuId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetThirdSkuId(v string) *CertificateGetResponseDataCertificateSku {
	s.ThirdSkuId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetGrouponType(v int) *CertificateGetResponseDataCertificateSku {
	s.GrouponType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetVoucherType(v int) *CertificateGetResponseDataCertificateSku {
	s.VoucherType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetSoldStartTime(v int64) *CertificateGetResponseDataCertificateSku {
	s.SoldStartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateSku) SetSkuOutId(v string) *CertificateGetResponseDataCertificateSku {
	s.SkuOutId = &v
	return s
}

type CertificateGetResponseDataCertificateTimeCard struct {
	TimesCount       *int32                                                               `json:"times_count,omitempty" xml:"times_count,omitempty"`
	TimesUsed        *int32                                                               `json:"times_used,omitempty" xml:"times_used,omitempty"`
	SerialAmountList []*CertificateGetResponseDataCertificateTimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
	TimeCardType     *int                                                                 `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
}

func (s CertificateGetResponseDataCertificateTimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateTimeCard) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateTimeCard) SetTimesCount(v int32) *CertificateGetResponseDataCertificateTimeCard {
	s.TimesCount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCard) SetTimesUsed(v int32) *CertificateGetResponseDataCertificateTimeCard {
	s.TimesUsed = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCard) SetSerialAmountList(v []*CertificateGetResponseDataCertificateTimeCardSerialAmountListItem) *CertificateGetResponseDataCertificateTimeCard {
	s.SerialAmountList = v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCard) SetTimeCardType(v int) *CertificateGetResponseDataCertificateTimeCard {
	s.TimeCardType = &v
	return s
}

type CertificateGetResponseDataCertificateTimeCardSerialAmountListItem struct {
	SerialNumb *int32                                                                   `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
	Amount     *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s CertificateGetResponseDataCertificateTimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateTimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItem) SetAmount(v *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

type CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount struct {
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
}

func (s CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificateGetResponseDataCertificateTimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

type CertificateGetResponseDataCertificateUseTimeInfo struct {
	UseTimeType    *int                                                                  `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
	TimePeriodList []*CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateUseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateUseTimeInfo) SetUseTimeType(v int) *CertificateGetResponseDataCertificateUseTimeInfo {
	s.UseTimeType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateUseTimeInfo) SetTimePeriodList(v []*CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) *CertificateGetResponseDataCertificateUseTimeInfo {
	s.TimePeriodList = v
	return s
}

type CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem struct {
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
}

func (s CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificateGetResponseDataCertificateUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

type CertificateGetResponseDataCertificateV2 struct {
	UseTimeInfo          *CertificateGetResponseDataCertificateV2UseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	Code                 *string                                                      `json:"code,omitempty" xml:"code,omitempty"`
	Sku                  *CertificateGetResponseDataCertificateV2Sku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	StartTime            *int64                                                       `json:"start_time,omitempty" xml:"start_time,omitempty"`
	UsedStatusType       *int                                                         `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
	ExpireTime           *int64                                                       `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Amount               *CertificateGetResponseDataCertificateV2Amount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	EncryptedCode        *string                                                      `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	Verify               *CertificateGetResponseDataCertificateV2Verify               `json:"verify,omitempty" xml:"verify,omitempty"`
	VerifyRecords        []*CertificateGetResponseDataCertificateV2VerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	TimeCard             *CertificateGetResponseDataCertificateV2TimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	ReserveInfo          *CertificateGetResponseDataCertificateV2ReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	AdditionalMap        map[int]*string                                              `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	BookInfo             *CertificateGetResponseDataCertificateV2BookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	NotAvailablePoiList  []*string                                                    `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	Status               *int                                                         `json:"status,omitempty" xml:"status,omitempty"`
	CertificateId        *int64                                                       `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	NotAvailableTimeInfo *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	PeriodCard           *CertificateGetResponseDataCertificateV2PeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	OffPeakDiscountInfo  *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2) SetUseTimeInfo(v *CertificateGetResponseDataCertificateV2UseTimeInfo) *CertificateGetResponseDataCertificateV2 {
	s.UseTimeInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetCode(v string) *CertificateGetResponseDataCertificateV2 {
	s.Code = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetSku(v *CertificateGetResponseDataCertificateV2Sku) *CertificateGetResponseDataCertificateV2 {
	s.Sku = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetStartTime(v int64) *CertificateGetResponseDataCertificateV2 {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetUsedStatusType(v int) *CertificateGetResponseDataCertificateV2 {
	s.UsedStatusType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetExpireTime(v int64) *CertificateGetResponseDataCertificateV2 {
	s.ExpireTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetAmount(v *CertificateGetResponseDataCertificateV2Amount) *CertificateGetResponseDataCertificateV2 {
	s.Amount = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetEncryptedCode(v string) *CertificateGetResponseDataCertificateV2 {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetVerify(v *CertificateGetResponseDataCertificateV2Verify) *CertificateGetResponseDataCertificateV2 {
	s.Verify = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetVerifyRecords(v []*CertificateGetResponseDataCertificateV2VerifyRecordsItem) *CertificateGetResponseDataCertificateV2 {
	s.VerifyRecords = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetTimeCard(v *CertificateGetResponseDataCertificateV2TimeCard) *CertificateGetResponseDataCertificateV2 {
	s.TimeCard = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetReserveInfo(v *CertificateGetResponseDataCertificateV2ReserveInfo) *CertificateGetResponseDataCertificateV2 {
	s.ReserveInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetAdditionalMap(v map[int]*string) *CertificateGetResponseDataCertificateV2 {
	s.AdditionalMap = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetBookInfo(v *CertificateGetResponseDataCertificateV2BookInfo) *CertificateGetResponseDataCertificateV2 {
	s.BookInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetNotAvailablePoiList(v []*string) *CertificateGetResponseDataCertificateV2 {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetStatus(v int) *CertificateGetResponseDataCertificateV2 {
	s.Status = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetCertificateId(v int64) *CertificateGetResponseDataCertificateV2 {
	s.CertificateId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetNotAvailableTimeInfo(v *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) *CertificateGetResponseDataCertificateV2 {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetPeriodCard(v *CertificateGetResponseDataCertificateV2PeriodCard) *CertificateGetResponseDataCertificateV2 {
	s.PeriodCard = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2) SetOffPeakDiscountInfo(v *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) *CertificateGetResponseDataCertificateV2 {
	s.OffPeakDiscountInfo = v
	return s
}

type CertificateGetResponseDataCertificateV2Amount struct {
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2Amount) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2Amount) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetBrandTicketAmount(v int64) *CertificateGetResponseDataCertificateV2Amount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetOriginalAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetMerchantTicketAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetPayAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.PayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetPlatformDiscountAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetPaymentDiscountAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetOriginalCurrency(v string) *CertificateGetResponseDataCertificateV2Amount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetOriginListMarketAmount(v int64) *CertificateGetResponseDataCertificateV2Amount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetCouponPayAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Amount) SetListMarketAmount(v int32) *CertificateGetResponseDataCertificateV2Amount {
	s.ListMarketAmount = &v
	return s
}

type CertificateGetResponseDataCertificateV2BookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2BookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2BookInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2BookInfo) SetBookPoiId(v string) *CertificateGetResponseDataCertificateV2BookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2BookInfo) SetBookProductNumber(v int64) *CertificateGetResponseDataCertificateV2BookInfo {
	s.BookProductNumber = &v
	return s
}

type CertificateGetResponseDataCertificateV2NotAvailableTimeInfo struct {
	FulfilEnable    *bool                                                                          `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
	CanNoUseDate    []*CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
	CanNoUseWeekDay []*int64                                                                       `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) SetCanNoUseDate(v []*CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem) *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificateGetResponseDataCertificateV2NotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

type CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem struct {
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificateGetResponseDataCertificateV2NotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

type CertificateGetResponseDataCertificateV2OffPeakDiscountInfo struct {
	IdleTimeLimitType  *int                                                                              `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	OffPeakTimeRange   []*CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
	HasOffPeakDiscount *bool                                                                             `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

type CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem struct {
	EndTime            *int64                                                                                                  `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                  `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                  `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

type CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificateGetResponseDataCertificateV2OffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type CertificateGetResponseDataCertificateV2PeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2PeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2PeriodCard) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2PeriodCard) SetPeriodType(v int) *CertificateGetResponseDataCertificateV2PeriodCard {
	s.PeriodType = &v
	return s
}

type CertificateGetResponseDataCertificateV2ReserveInfo struct {
	OrderReserveUserInfoList []*CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificateGetResponseDataCertificateV2ReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2ReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2ReserveInfo) SetOrderReserveUserInfoList(v []*CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) *CertificateGetResponseDataCertificateV2ReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem struct {
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificateGetResponseDataCertificateV2ReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

type CertificateGetResponseDataCertificateV2Sku struct {
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2Sku) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2Sku) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetAccountId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.AccountId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetGrouponType(v int) *CertificateGetResponseDataCertificateV2Sku {
	s.GrouponType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetProductOutId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.ProductOutId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetTitle(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.Title = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetProductId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.ProductId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetVoucherType(v int) *CertificateGetResponseDataCertificateV2Sku {
	s.VoucherType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetSkuId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.SkuId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetThirdSkuId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.ThirdSkuId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetMarketPrice(v int64) *CertificateGetResponseDataCertificateV2Sku {
	s.MarketPrice = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetSuplierProductOutId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetSkuOutId(v string) *CertificateGetResponseDataCertificateV2Sku {
	s.SkuOutId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Sku) SetSoldStartTime(v int64) *CertificateGetResponseDataCertificateV2Sku {
	s.SoldStartTime = &v
	return s
}

type CertificateGetResponseDataCertificateV2TimeCard struct {
	TimesCount       *int32                                                                 `json:"times_count,omitempty" xml:"times_count,omitempty"`
	TimesUsed        *int32                                                                 `json:"times_used,omitempty" xml:"times_used,omitempty"`
	SerialAmountList []*CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
	TimeCardType     *int                                                                   `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2TimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2TimeCard) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2TimeCard) SetTimesCount(v int32) *CertificateGetResponseDataCertificateV2TimeCard {
	s.TimesCount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCard) SetTimesUsed(v int32) *CertificateGetResponseDataCertificateV2TimeCard {
	s.TimesUsed = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCard) SetSerialAmountList(v []*CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem) *CertificateGetResponseDataCertificateV2TimeCard {
	s.SerialAmountList = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCard) SetTimeCardType(v int) *CertificateGetResponseDataCertificateV2TimeCard {
	s.TimeCardType = &v
	return s
}

type CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem struct {
	Amount     *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                     `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem) SetAmount(v *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

type CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount struct {
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
}

func (s CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificateGetResponseDataCertificateV2TimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

type CertificateGetResponseDataCertificateV2UseTimeInfo struct {
	TimePeriodList []*CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
	UseTimeType    *int                                                                    `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2UseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2UseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2UseTimeInfo) SetTimePeriodList(v []*CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) *CertificateGetResponseDataCertificateV2UseTimeInfo {
	s.TimePeriodList = v
	return s
}

func (s *CertificateGetResponseDataCertificateV2UseTimeInfo) SetUseTimeType(v int) *CertificateGetResponseDataCertificateV2UseTimeInfo {
	s.UseTimeType = &v
	return s
}

type CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificateGetResponseDataCertificateV2UseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

type CertificateGetResponseDataCertificateV2Verify struct {
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2Verify) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2Verify) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetCertificateId(v string) *CertificateGetResponseDataCertificateV2Verify {
	s.CertificateId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetPoiId(v int64) *CertificateGetResponseDataCertificateV2Verify {
	s.PoiId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetTimesCardSerialNum(v int32) *CertificateGetResponseDataCertificateV2Verify {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetVerifierUniqueId(v string) *CertificateGetResponseDataCertificateV2Verify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetVerifyId(v string) *CertificateGetResponseDataCertificateV2Verify {
	s.VerifyId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetVerifyTime(v int64) *CertificateGetResponseDataCertificateV2Verify {
	s.VerifyTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetVerifyType(v int) *CertificateGetResponseDataCertificateV2Verify {
	s.VerifyType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2Verify) SetCanCancel(v bool) *CertificateGetResponseDataCertificateV2Verify {
	s.CanCancel = &v
	return s
}

type CertificateGetResponseDataCertificateV2VerifyRecordsItem struct {
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
}

func (s CertificateGetResponseDataCertificateV2VerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateV2VerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetPoiId(v int64) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.PoiId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetVerifierUniqueId(v string) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetVerifyId(v string) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetVerifyTime(v int64) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetVerifyType(v int) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetCanCancel(v bool) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificateGetResponseDataCertificateV2VerifyRecordsItem) SetCertificateId(v string) *CertificateGetResponseDataCertificateV2VerifyRecordsItem {
	s.CertificateId = &v
	return s
}

type CertificateGetResponseDataCertificateVerify struct {
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
}

func (s CertificateGetResponseDataCertificateVerify) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateVerify) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateVerify) SetVerifierUniqueId(v string) *CertificateGetResponseDataCertificateVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetVerifyId(v string) *CertificateGetResponseDataCertificateVerify {
	s.VerifyId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetVerifyTime(v int64) *CertificateGetResponseDataCertificateVerify {
	s.VerifyTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetVerifyType(v int) *CertificateGetResponseDataCertificateVerify {
	s.VerifyType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetCanCancel(v bool) *CertificateGetResponseDataCertificateVerify {
	s.CanCancel = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetCertificateId(v string) *CertificateGetResponseDataCertificateVerify {
	s.CertificateId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetPoiId(v int64) *CertificateGetResponseDataCertificateVerify {
	s.PoiId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerify) SetTimesCardSerialNum(v int32) *CertificateGetResponseDataCertificateVerify {
	s.TimesCardSerialNum = &v
	return s
}

type CertificateGetResponseDataCertificateVerifyRecordsItem struct {
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s CertificateGetResponseDataCertificateVerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseDataCertificateVerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetVerifierUniqueId(v string) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetVerifyId(v string) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetVerifyTime(v int64) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetVerifyType(v int) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetCanCancel(v bool) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetCertificateId(v string) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateGetResponseDataCertificateVerifyRecordsItem) SetPoiId(v int64) *CertificateGetResponseDataCertificateVerifyRecordsItem {
	s.PoiId = &v
	return s
}

type CertificateGetResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s CertificateGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificateGetResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificateGetResponseExtra) SetDescription(v string) *CertificateGetResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificateGetResponseExtra) SetErrorCode(v int32) *CertificateGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificateGetResponseExtra) SetLogid(v string) *CertificateGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *CertificateGetResponseExtra) SetNow(v int64) *CertificateGetResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificateGetResponseExtra) SetSubDescription(v string) *CertificateGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificateGetResponseExtra) SetSubErrorCode(v int32) *CertificateGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

type CertificatePrepareRequest struct {
	PoiId         *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Code          *string            `json:"code,omitempty" xml:"code,omitempty"`
	EncryptedData *string            `json:"encrypted_data,omitempty" xml:"encrypted_data,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CertificatePrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareRequest) GoString() string {
	return s.String()
}

func (s *CertificatePrepareRequest) SetPoiId(v string) *CertificatePrepareRequest {
	s.PoiId = &v
	return s
}

func (s *CertificatePrepareRequest) SetAccountId(v string) *CertificatePrepareRequest {
	s.AccountId = &v
	return s
}

func (s *CertificatePrepareRequest) SetCode(v string) *CertificatePrepareRequest {
	s.Code = &v
	return s
}

func (s *CertificatePrepareRequest) SetEncryptedData(v string) *CertificatePrepareRequest {
	s.EncryptedData = &v
	return s
}

func (s *CertificatePrepareRequest) SetHeader(v map[string]*string) *CertificatePrepareRequest {
	s.Header = v
	return s
}

func (s *CertificatePrepareRequest) SetAccessToken(v string) *CertificatePrepareRequest {
	s.AccessToken = &v
	return s
}

type CertificatePrepareResponse struct {
	Data  *CertificatePrepareResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *CertificatePrepareResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s CertificatePrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponse) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponse) SetData(v *CertificatePrepareResponseData) *CertificatePrepareResponse {
	s.Data = v
	return s
}

func (s *CertificatePrepareResponse) SetExtra(v *CertificatePrepareResponseExtra) *CertificatePrepareResponse {
	s.Extra = v
	return s
}

type CertificatePrepareResponseData struct {
	VerifyToken    *string                                             `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	Certificates   []*CertificatePrepareResponseDataCertificatesItem   `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	CertificatesV2 []*CertificatePrepareResponseDataCertificatesV2Item `json:"certificates_v2,omitempty" xml:"certificates_v2,omitempty" type:"Repeated"`
	GwErrorCode    *int32                                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription  *string                                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	OrderId        *string                                             `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s CertificatePrepareResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseData) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseData) SetVerifyToken(v string) *CertificatePrepareResponseData {
	s.VerifyToken = &v
	return s
}

func (s *CertificatePrepareResponseData) SetCertificates(v []*CertificatePrepareResponseDataCertificatesItem) *CertificatePrepareResponseData {
	s.Certificates = v
	return s
}

func (s *CertificatePrepareResponseData) SetCertificatesV2(v []*CertificatePrepareResponseDataCertificatesV2Item) *CertificatePrepareResponseData {
	s.CertificatesV2 = v
	return s
}

func (s *CertificatePrepareResponseData) SetGwErrorCode(v int32) *CertificatePrepareResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertificatePrepareResponseData) SetGwDescription(v string) *CertificatePrepareResponseData {
	s.GwDescription = &v
	return s
}

func (s *CertificatePrepareResponseData) SetOrderId(v string) *CertificatePrepareResponseData {
	s.OrderId = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItem struct {
	OffPeakDiscountInfo  *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
	ReserveInfo          *CertificatePrepareResponseDataCertificatesItemReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	NotAvailablePoiList  []*string                                                           `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	NotAvailableTimeInfo *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	Code                 *string                                                             `json:"code,omitempty" xml:"code,omitempty"`
	VerifyRecords        []*CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	TimeCard             *CertificatePrepareResponseDataCertificatesItemTimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	Sku                  *CertificatePrepareResponseDataCertificatesItemSku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	ExpireTime           *int64                                                              `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	BookInfo             *CertificatePrepareResponseDataCertificatesItemBookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	UseTimeInfo          *CertificatePrepareResponseDataCertificatesItemUseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	UsedStatusType       *int                                                                `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
	AdditionalMap        map[int]*string                                                     `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	CertificateId        *int64                                                              `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Amount               *CertificatePrepareResponseDataCertificatesItemAmount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	StartTime            *int64                                                              `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Verify               *CertificatePrepareResponseDataCertificatesItemVerify               `json:"verify,omitempty" xml:"verify,omitempty"`
	PeriodCard           *CertificatePrepareResponseDataCertificatesItemPeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	EncryptedCode        *string                                                             `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	Status               *int                                                                `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetOffPeakDiscountInfo(v *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) *CertificatePrepareResponseDataCertificatesItem {
	s.OffPeakDiscountInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetReserveInfo(v *CertificatePrepareResponseDataCertificatesItemReserveInfo) *CertificatePrepareResponseDataCertificatesItem {
	s.ReserveInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetNotAvailablePoiList(v []*string) *CertificatePrepareResponseDataCertificatesItem {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetNotAvailableTimeInfo(v *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) *CertificatePrepareResponseDataCertificatesItem {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetCode(v string) *CertificatePrepareResponseDataCertificatesItem {
	s.Code = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetVerifyRecords(v []*CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) *CertificatePrepareResponseDataCertificatesItem {
	s.VerifyRecords = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetTimeCard(v *CertificatePrepareResponseDataCertificatesItemTimeCard) *CertificatePrepareResponseDataCertificatesItem {
	s.TimeCard = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetSku(v *CertificatePrepareResponseDataCertificatesItemSku) *CertificatePrepareResponseDataCertificatesItem {
	s.Sku = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetExpireTime(v int64) *CertificatePrepareResponseDataCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetBookInfo(v *CertificatePrepareResponseDataCertificatesItemBookInfo) *CertificatePrepareResponseDataCertificatesItem {
	s.BookInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetUseTimeInfo(v *CertificatePrepareResponseDataCertificatesItemUseTimeInfo) *CertificatePrepareResponseDataCertificatesItem {
	s.UseTimeInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetUsedStatusType(v int) *CertificatePrepareResponseDataCertificatesItem {
	s.UsedStatusType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetAdditionalMap(v map[int]*string) *CertificatePrepareResponseDataCertificatesItem {
	s.AdditionalMap = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetCertificateId(v int64) *CertificatePrepareResponseDataCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetAmount(v *CertificatePrepareResponseDataCertificatesItemAmount) *CertificatePrepareResponseDataCertificatesItem {
	s.Amount = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetVerify(v *CertificatePrepareResponseDataCertificatesItemVerify) *CertificatePrepareResponseDataCertificatesItem {
	s.Verify = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetPeriodCard(v *CertificatePrepareResponseDataCertificatesItemPeriodCard) *CertificatePrepareResponseDataCertificatesItem {
	s.PeriodCard = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetEncryptedCode(v string) *CertificatePrepareResponseDataCertificatesItem {
	s.EncryptedCode = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItem) SetStatus(v int) *CertificatePrepareResponseDataCertificatesItem {
	s.Status = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemAmount struct {
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemAmount) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetPayAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetOriginalCurrency(v string) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetListMarketAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetCouponPayAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetBrandTicketAmount(v int64) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetOriginListMarketAmount(v int64) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetPlatformDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetOriginalAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetPaymentDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemAmount) SetMerchantTicketAmount(v int32) *CertificatePrepareResponseDataCertificatesItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
	VerifyAmount      *int64  `json:"verify_amount,omitempty" xml:"verify_amount,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemBookInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemBookInfo) SetBookPoiId(v string) *CertificatePrepareResponseDataCertificatesItemBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemBookInfo) SetBookProductNumber(v int64) *CertificatePrepareResponseDataCertificatesItemBookInfo {
	s.BookProductNumber = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemBookInfo) SetVerifyAmount(v int64) *CertificatePrepareResponseDataCertificatesItemBookInfo {
	s.VerifyAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo struct {
	CanNoUseWeekDay []*int64                                                                              `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
	FulfilEnable    *bool                                                                                 `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
	CanNoUseDate    []*CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo) SetCanNoUseDate(v []*CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

type CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem struct {
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificatePrepareResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo struct {
	HasOffPeakDiscount *bool                                                                                    `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
	IdleTimeLimitType  *int                                                                                     `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	OffPeakTimeRange   []*CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

type CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem struct {
	EndTime            *int64                                                                                                         `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                         `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                         `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

type CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificatePrepareResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemPeriodCard) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemPeriodCard) SetPeriodType(v int) *CertificatePrepareResponseDataCertificatesItemPeriodCard {
	s.PeriodType = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemReserveInfo struct {
	OrderReserveUserInfoList []*CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesItemReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemReserveInfo) SetOrderReserveUserInfoList(v []*CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) *CertificatePrepareResponseDataCertificatesItemReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem struct {
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificatePrepareResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemSku struct {
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemSku) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemSku) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetMarketPrice(v int64) *CertificatePrepareResponseDataCertificatesItemSku {
	s.MarketPrice = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetProductOutId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.ProductOutId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetProductId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.ProductId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetSuplierProductOutId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetVoucherType(v int) *CertificatePrepareResponseDataCertificatesItemSku {
	s.VoucherType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetSoldStartTime(v int64) *CertificatePrepareResponseDataCertificatesItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetTitle(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.Title = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetGrouponType(v int) *CertificatePrepareResponseDataCertificatesItemSku {
	s.GrouponType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetSkuOutId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.SkuOutId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetSkuId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.SkuId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetAccountId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.AccountId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemSku) SetThirdSkuId(v string) *CertificatePrepareResponseDataCertificatesItemSku {
	s.ThirdSkuId = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemTimeCard struct {
	SerialAmountList []*CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
	TimeCardType     *int                                                                          `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
	TimesCount       *int32                                                                        `json:"times_count,omitempty" xml:"times_count,omitempty"`
	TimesUsed        *int32                                                                        `json:"times_used,omitempty" xml:"times_used,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCard) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCard) SetSerialAmountList(v []*CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem) *CertificatePrepareResponseDataCertificatesItemTimeCard {
	s.SerialAmountList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCard) SetTimeCardType(v int) *CertificatePrepareResponseDataCertificatesItemTimeCard {
	s.TimeCardType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCard) SetTimesCount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCard {
	s.TimesCount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCard) SetTimesUsed(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCard {
	s.TimesUsed = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem struct {
	Amount     *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                            `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem) SetAmount(v *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount struct {
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificatePrepareResponseDataCertificatesItemTimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemUseTimeInfo struct {
	UseTimeType    *int                                                                           `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
	TimePeriodList []*CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemUseTimeInfo) SetUseTimeType(v int) *CertificatePrepareResponseDataCertificatesItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemUseTimeInfo) SetTimePeriodList(v []*CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) *CertificatePrepareResponseDataCertificatesItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

type CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificatePrepareResponseDataCertificatesItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemVerify struct {
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesItemVerify) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemVerify) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetPoiId(v int64) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.PoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetTimesCardSerialNum(v int32) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetVerifierUniqueId(v string) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetVerifyId(v string) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.VerifyId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetVerifyTime(v int64) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.VerifyTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetVerifyType(v int) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.VerifyType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetCanCancel(v bool) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.CanCancel = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerify) SetCertificateId(v string) *CertificatePrepareResponseDataCertificatesItemVerify {
	s.CertificateId = &v
	return s
}

type CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem struct {
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetVerifierUniqueId(v string) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetVerifyId(v string) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetVerifyTime(v int64) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetVerifyType(v int) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetCanCancel(v bool) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetCertificateId(v string) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetPoiId(v int64) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.PoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificatePrepareResponseDataCertificatesItemVerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2Item struct {
	Amount               *CertificatePrepareResponseDataCertificatesV2ItemAmount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	PeriodCard           *CertificatePrepareResponseDataCertificatesV2ItemPeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	TimeCard             *CertificatePrepareResponseDataCertificatesV2ItemTimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	Sku                  *CertificatePrepareResponseDataCertificatesV2ItemSku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	Status               *int                                                                  `json:"status,omitempty" xml:"status,omitempty"`
	NotAvailableTimeInfo *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	UseTimeInfo          *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	UsedStatusType       *int                                                                  `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
	ReserveInfo          *CertificatePrepareResponseDataCertificatesV2ItemReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	NotAvailablePoiList  []*string                                                             `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	VerifyRecords        []*CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	Verify               *CertificatePrepareResponseDataCertificatesV2ItemVerify               `json:"verify,omitempty" xml:"verify,omitempty"`
	OffPeakDiscountInfo  *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
	AdditionalMap        map[int]*string                                                       `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	BookInfo             *CertificatePrepareResponseDataCertificatesV2ItemBookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	EncryptedCode        *string                                                               `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	StartTime            *int64                                                                `json:"start_time,omitempty" xml:"start_time,omitempty"`
	CertificateId        *int64                                                                `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Code                 *string                                                               `json:"code,omitempty" xml:"code,omitempty"`
	ExpireTime           *int64                                                                `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2Item) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2Item) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetAmount(v *CertificatePrepareResponseDataCertificatesV2ItemAmount) *CertificatePrepareResponseDataCertificatesV2Item {
	s.Amount = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetPeriodCard(v *CertificatePrepareResponseDataCertificatesV2ItemPeriodCard) *CertificatePrepareResponseDataCertificatesV2Item {
	s.PeriodCard = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetTimeCard(v *CertificatePrepareResponseDataCertificatesV2ItemTimeCard) *CertificatePrepareResponseDataCertificatesV2Item {
	s.TimeCard = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetSku(v *CertificatePrepareResponseDataCertificatesV2ItemSku) *CertificatePrepareResponseDataCertificatesV2Item {
	s.Sku = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetStatus(v int) *CertificatePrepareResponseDataCertificatesV2Item {
	s.Status = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetNotAvailableTimeInfo(v *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) *CertificatePrepareResponseDataCertificatesV2Item {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetUseTimeInfo(v *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo) *CertificatePrepareResponseDataCertificatesV2Item {
	s.UseTimeInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetUsedStatusType(v int) *CertificatePrepareResponseDataCertificatesV2Item {
	s.UsedStatusType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetReserveInfo(v *CertificatePrepareResponseDataCertificatesV2ItemReserveInfo) *CertificatePrepareResponseDataCertificatesV2Item {
	s.ReserveInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetNotAvailablePoiList(v []*string) *CertificatePrepareResponseDataCertificatesV2Item {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetVerifyRecords(v []*CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) *CertificatePrepareResponseDataCertificatesV2Item {
	s.VerifyRecords = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetVerify(v *CertificatePrepareResponseDataCertificatesV2ItemVerify) *CertificatePrepareResponseDataCertificatesV2Item {
	s.Verify = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetOffPeakDiscountInfo(v *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) *CertificatePrepareResponseDataCertificatesV2Item {
	s.OffPeakDiscountInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetAdditionalMap(v map[int]*string) *CertificatePrepareResponseDataCertificatesV2Item {
	s.AdditionalMap = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetBookInfo(v *CertificatePrepareResponseDataCertificatesV2ItemBookInfo) *CertificatePrepareResponseDataCertificatesV2Item {
	s.BookInfo = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetEncryptedCode(v string) *CertificatePrepareResponseDataCertificatesV2Item {
	s.EncryptedCode = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesV2Item {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetCertificateId(v int64) *CertificatePrepareResponseDataCertificatesV2Item {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetCode(v string) *CertificatePrepareResponseDataCertificatesV2Item {
	s.Code = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2Item) SetExpireTime(v int64) *CertificatePrepareResponseDataCertificatesV2Item {
	s.ExpireTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemAmount struct {
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemAmount) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetPlatformDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetMerchantTicketAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetOriginalCurrency(v string) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetCouponPayAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetOriginalAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetOriginListMarketAmount(v int64) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetListMarketAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetPaymentDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetPayAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemAmount) SetBrandTicketAmount(v int64) *CertificatePrepareResponseDataCertificatesV2ItemAmount {
	s.BrandTicketAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
	VerifyAmount      *int64  `json:"verify_amount,omitempty" xml:"verify_amount,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemBookInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemBookInfo) SetBookPoiId(v string) *CertificatePrepareResponseDataCertificatesV2ItemBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemBookInfo) SetBookProductNumber(v int64) *CertificatePrepareResponseDataCertificatesV2ItemBookInfo {
	s.BookProductNumber = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemBookInfo) SetVerifyAmount(v int64) *CertificatePrepareResponseDataCertificatesV2ItemBookInfo {
	s.VerifyAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo struct {
	CanNoUseWeekDay []*int64                                                                                `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
	FulfilEnable    *bool                                                                                   `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
	CanNoUseDate    []*CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo) SetCanNoUseDate(v []*CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem struct {
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemNotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo struct {
	IdleTimeLimitType  *int                                                                                       `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	OffPeakTimeRange   []*CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
	HasOffPeakDiscount *bool                                                                                      `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem struct {
	StartTime          *int64                                                                                                           `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                           `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
	EndTime            *int64                                                                                                           `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificatePrepareResponseDataCertificatesV2ItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemPeriodCard) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemPeriodCard) SetPeriodType(v int) *CertificatePrepareResponseDataCertificatesV2ItemPeriodCard {
	s.PeriodType = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemReserveInfo struct {
	OrderReserveUserInfoList []*CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemReserveInfo) SetOrderReserveUserInfoList(v []*CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) *CertificatePrepareResponseDataCertificatesV2ItemReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem struct {
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificatePrepareResponseDataCertificatesV2ItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemSku struct {
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemSku) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemSku) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetSoldStartTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetSkuId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.SkuId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetAccountId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.AccountId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetSuplierProductOutId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetSkuOutId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.SkuOutId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetThirdSkuId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetGrouponType(v int) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.GrouponType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetProductId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.ProductId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetVoucherType(v int) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.VoucherType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetTitle(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.Title = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetMarketPrice(v int64) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.MarketPrice = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemSku) SetProductOutId(v string) *CertificatePrepareResponseDataCertificatesV2ItemSku {
	s.ProductOutId = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemTimeCard struct {
	TimesUsed        *int32                                                                          `json:"times_used,omitempty" xml:"times_used,omitempty"`
	SerialAmountList []*CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem `json:"serial_amount_list,omitempty" xml:"serial_amount_list,omitempty" type:"Repeated"`
	TimeCardType     *int                                                                            `json:"time_card_type,omitempty" xml:"time_card_type,omitempty"`
	TimesCount       *int32                                                                          `json:"times_count,omitempty" xml:"times_count,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCard) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCard) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCard) SetTimesUsed(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCard {
	s.TimesUsed = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCard) SetSerialAmountList(v []*CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) *CertificatePrepareResponseDataCertificatesV2ItemTimeCard {
	s.SerialAmountList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCard) SetTimeCardType(v int) *CertificatePrepareResponseDataCertificatesV2ItemTimeCard {
	s.TimeCardType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCard) SetTimesCount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCard {
	s.TimesCount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem struct {
	Amount     *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount `json:"amount,omitempty" xml:"amount,omitempty"`
	SerialNumb *int32                                                                              `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) SetAmount(v *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem {
	s.Amount = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem) SetSerialNumb(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItem {
	s.SerialNumb = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount struct {
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetCouponPayAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPlatformDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPaymentDiscountAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetListMarketAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginListMarketAmount(v int64) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginalCurrency(v string) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetMerchantTicketAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetPayAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetOriginalAmount(v int32) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount) SetBrandTicketAmount(v int64) *CertificatePrepareResponseDataCertificatesV2ItemTimeCardSerialAmountListItemAmount {
	s.BrandTicketAmount = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo struct {
	TimePeriodList []*CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem `json:"time_period_list,omitempty" xml:"time_period_list,omitempty" type:"Repeated"`
	UseTimeType    *int                                                                             `json:"use_time_type,omitempty" xml:"use_time_type,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo) SetTimePeriodList(v []*CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo {
	s.TimePeriodList = v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo) SetUseTimeType(v int) *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfo {
	s.UseTimeType = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetEndTimeIsNextDay(v bool) *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetStartTime(v string) *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.StartTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem) SetEndTime(v string) *CertificatePrepareResponseDataCertificatesV2ItemUseTimeInfoTimePeriodListItem {
	s.EndTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemVerify struct {
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemVerify) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemVerify) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetVerifyType(v int) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.VerifyType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetCanCancel(v bool) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.CanCancel = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetCertificateId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetPoiId(v int64) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.PoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetTimesCardSerialNum(v int32) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetVerifierUniqueId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetVerifyId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.VerifyId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerify) SetVerifyTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemVerify {
	s.VerifyTime = &v
	return s
}

type CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem struct {
	VerifyTime         *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	VerifyType         *int    `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	CanCancel          *bool   `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	CertificateId      *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	PoiId              *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TimesCardSerialNum *int32  `json:"times_card_serial_num,omitempty" xml:"times_card_serial_num,omitempty"`
	VerifierUniqueId   *string `json:"verifier_unique_id,omitempty" xml:"verifier_unique_id,omitempty"`
	VerifyId           *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
}

func (s CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyTime(v int64) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyType(v int) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetCanCancel(v bool) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetCertificateId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetPoiId(v int64) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.PoiId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetTimesCardSerialNum(v int32) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.TimesCardSerialNum = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifierUniqueId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifierUniqueId = &v
	return s
}

func (s *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem) SetVerifyId(v string) *CertificatePrepareResponseDataCertificatesV2ItemVerifyRecordsItem {
	s.VerifyId = &v
	return s
}

type CertificatePrepareResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s CertificatePrepareResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s CertificatePrepareResponseExtra) GoString() string {
	return s.String()
}

func (s *CertificatePrepareResponseExtra) SetNow(v int64) *CertificatePrepareResponseExtra {
	s.Now = &v
	return s
}

func (s *CertificatePrepareResponseExtra) SetSubDescription(v string) *CertificatePrepareResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *CertificatePrepareResponseExtra) SetSubErrorCode(v int32) *CertificatePrepareResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *CertificatePrepareResponseExtra) SetDescription(v string) *CertificatePrepareResponseExtra {
	s.Description = &v
	return s
}

func (s *CertificatePrepareResponseExtra) SetErrorCode(v int32) *CertificatePrepareResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *CertificatePrepareResponseExtra) SetLogid(v string) *CertificatePrepareResponseExtra {
	s.Logid = &v
	return s
}

type CertificateQueryRequest struct {
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty"`
	EncryptedCode *string            `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty"`
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s CertificateQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryRequest) GoString() string {
	return s.String()
}

func (s *CertificateQueryRequest) SetAccountId(v string) *CertificateQueryRequest {
	s.AccountId = &v
	return s
}

func (s *CertificateQueryRequest) SetEncryptedCode(v string) *CertificateQueryRequest {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateQueryRequest) SetOrderId(v string) *CertificateQueryRequest {
	s.OrderId = &v
	return s
}

func (s *CertificateQueryRequest) SetHeader(v map[string]*string) *CertificateQueryRequest {
	s.Header = v
	return s
}

func (s *CertificateQueryRequest) SetAccessToken(v string) *CertificateQueryRequest {
	s.AccessToken = &v
	return s
}

type CertificateQueryResponse struct {
	Extra *CertificateQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *CertificateQueryResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s CertificateQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponse) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponse) SetExtra(v *CertificateQueryResponseExtra) *CertificateQueryResponse {
	s.Extra = v
	return s
}

func (s *CertificateQueryResponse) SetData(v *CertificateQueryResponseData) *CertificateQueryResponse {
	s.Data = v
	return s
}

type CertificateQueryResponseData struct {
	GwErrorCode    *int32                                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription  *string                                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Certificates   []*CertificateQueryResponseDataCertificatesItem   `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
	CertificatesV2 []*CertificateQueryResponseDataCertificatesV2Item `json:"certificates_v2,omitempty" xml:"certificates_v2,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseData) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseData) SetGwErrorCode(v int32) *CertificateQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *CertificateQueryResponseData) SetGwDescription(v string) *CertificateQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *CertificateQueryResponseData) SetCertificates(v []*CertificateQueryResponseDataCertificatesItem) *CertificateQueryResponseData {
	s.Certificates = v
	return s
}

func (s *CertificateQueryResponseData) SetCertificatesV2(v []*CertificateQueryResponseDataCertificatesV2Item) *CertificateQueryResponseData {
	s.CertificatesV2 = v
	return s
}

type CertificateQueryResponseDataCertificatesItem struct {
	Code                 *string                                                           `json:"code,omitempty" xml:"code,omitempty"`
	OffPeakDiscountInfo  *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo  `json:"off_peak_discount_info,omitempty" xml:"off_peak_discount_info,omitempty"`
	VerifyRecords        []*CertificateQueryResponseDataCertificatesItemVerifyRecordsItem  `json:"verify_records,omitempty" xml:"verify_records,omitempty" type:"Repeated"`
	Verify               *CertificateQueryResponseDataCertificatesItemVerify               `json:"verify,omitempty" xml:"verify,omitempty"`
	BookInfo             *CertificateQueryResponseDataCertificatesItemBookInfo             `json:"book_info,omitempty" xml:"book_info,omitempty"`
	Sku                  *CertificateQueryResponseDataCertificatesItemSku                  `json:"sku,omitempty" xml:"sku,omitempty" require:"true"`
	UseTimeInfo          *CertificateQueryResponseDataCertificatesItemUseTimeInfo          `json:"use_time_info,omitempty" xml:"use_time_info,omitempty"`
	TimeCard             *CertificateQueryResponseDataCertificatesItemTimeCard             `json:"time_card,omitempty" xml:"time_card,omitempty"`
	AdditionalMap        map[int]*string                                                   `json:"additional_map,omitempty" xml:"additional_map,omitempty"`
	StartTime            *int64                                                            `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ExpireTime           *int64                                                            `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Status               *int                                                              `json:"status,omitempty" xml:"status,omitempty"`
	EncryptedCode        *string                                                           `json:"encrypted_code,omitempty" xml:"encrypted_code,omitempty" require:"true"`
	Amount               *CertificateQueryResponseDataCertificatesItemAmount               `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	NotAvailablePoiList  []*string                                                         `json:"not_available_poi_list,omitempty" xml:"not_available_poi_list,omitempty" type:"Repeated"`
	ReserveInfo          *CertificateQueryResponseDataCertificatesItemReserveInfo          `json:"reserve_info,omitempty" xml:"reserve_info,omitempty"`
	CertificateId        *int64                                                            `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	NotAvailableTimeInfo *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo `json:"not_available_time_info,omitempty" xml:"not_available_time_info,omitempty"`
	PeriodCard           *CertificateQueryResponseDataCertificatesItemPeriodCard           `json:"period_card,omitempty" xml:"period_card,omitempty"`
	UsedStatusType       *int                                                              `json:"used_status_type,omitempty" xml:"used_status_type,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItem) SetCode(v string) *CertificateQueryResponseDataCertificatesItem {
	s.Code = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetOffPeakDiscountInfo(v *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) *CertificateQueryResponseDataCertificatesItem {
	s.OffPeakDiscountInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetVerifyRecords(v []*CertificateQueryResponseDataCertificatesItemVerifyRecordsItem) *CertificateQueryResponseDataCertificatesItem {
	s.VerifyRecords = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetVerify(v *CertificateQueryResponseDataCertificatesItemVerify) *CertificateQueryResponseDataCertificatesItem {
	s.Verify = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetBookInfo(v *CertificateQueryResponseDataCertificatesItemBookInfo) *CertificateQueryResponseDataCertificatesItem {
	s.BookInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetSku(v *CertificateQueryResponseDataCertificatesItemSku) *CertificateQueryResponseDataCertificatesItem {
	s.Sku = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetUseTimeInfo(v *CertificateQueryResponseDataCertificatesItemUseTimeInfo) *CertificateQueryResponseDataCertificatesItem {
	s.UseTimeInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetTimeCard(v *CertificateQueryResponseDataCertificatesItemTimeCard) *CertificateQueryResponseDataCertificatesItem {
	s.TimeCard = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetAdditionalMap(v map[int]*string) *CertificateQueryResponseDataCertificatesItem {
	s.AdditionalMap = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetExpireTime(v int64) *CertificateQueryResponseDataCertificatesItem {
	s.ExpireTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetStatus(v int) *CertificateQueryResponseDataCertificatesItem {
	s.Status = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetEncryptedCode(v string) *CertificateQueryResponseDataCertificatesItem {
	s.EncryptedCode = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetAmount(v *CertificateQueryResponseDataCertificatesItemAmount) *CertificateQueryResponseDataCertificatesItem {
	s.Amount = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetNotAvailablePoiList(v []*string) *CertificateQueryResponseDataCertificatesItem {
	s.NotAvailablePoiList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetReserveInfo(v *CertificateQueryResponseDataCertificatesItemReserveInfo) *CertificateQueryResponseDataCertificatesItem {
	s.ReserveInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetCertificateId(v int64) *CertificateQueryResponseDataCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetNotAvailableTimeInfo(v *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) *CertificateQueryResponseDataCertificatesItem {
	s.NotAvailableTimeInfo = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetPeriodCard(v *CertificateQueryResponseDataCertificatesItemPeriodCard) *CertificateQueryResponseDataCertificatesItem {
	s.PeriodCard = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItem) SetUsedStatusType(v int) *CertificateQueryResponseDataCertificatesItem {
	s.UsedStatusType = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemAmount struct {
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemAmount) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemAmount) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetOriginalCurrency(v string) *CertificateQueryResponseDataCertificatesItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetPayAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.PayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetBrandTicketAmount(v int64) *CertificateQueryResponseDataCertificatesItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetOriginalAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetMerchantTicketAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetOriginListMarketAmount(v int64) *CertificateQueryResponseDataCertificatesItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetListMarketAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetCouponPayAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetPaymentDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemAmount) SetPlatformDiscountAmount(v int32) *CertificateQueryResponseDataCertificatesItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemBookInfo struct {
	BookPoiId         *string `json:"book_poi_id,omitempty" xml:"book_poi_id,omitempty"`
	BookProductNumber *int64  `json:"book_product_number,omitempty" xml:"book_product_number,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemBookInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemBookInfo) SetBookPoiId(v string) *CertificateQueryResponseDataCertificatesItemBookInfo {
	s.BookPoiId = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemBookInfo) SetBookProductNumber(v int64) *CertificateQueryResponseDataCertificatesItemBookInfo {
	s.BookProductNumber = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo struct {
	CanNoUseWeekDay []*int64                                                                            `json:"can_no_use_week_day,omitempty" xml:"can_no_use_week_day,omitempty" type:"Repeated"`
	FulfilEnable    *bool                                                                               `json:"fulfil_enable,omitempty" xml:"fulfil_enable,omitempty"`
	CanNoUseDate    []*CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem `json:"can_no_use_date,omitempty" xml:"can_no_use_date,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) SetCanNoUseWeekDay(v []*int64) *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo {
	s.CanNoUseWeekDay = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) SetFulfilEnable(v bool) *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo {
	s.FulfilEnable = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo) SetCanNoUseDate(v []*CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfo {
	s.CanNoUseDate = v
	return s
}

type CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem struct {
	EndTime   *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) SetEndTime(v int64) *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesItemNotAvailableTimeInfoCanNoUseDateItem {
	s.StartTime = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo struct {
	OffPeakTimeRange   []*CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem `json:"off_peak_time_range,omitempty" xml:"off_peak_time_range,omitempty" type:"Repeated"`
	HasOffPeakDiscount *bool                                                                                  `json:"has_off_peak_discount,omitempty" xml:"has_off_peak_discount,omitempty" require:"true"`
	IdleTimeLimitType  *int                                                                                   `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) SetOffPeakTimeRange(v []*CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo {
	s.OffPeakTimeRange = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) SetHasOffPeakDiscount(v bool) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo {
	s.HasOffPeakDiscount = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo) SetIdleTimeLimitType(v int) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem struct {
	EndTime            *int64                                                                                                       `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime          *int64                                                                                                       `json:"start_time,omitempty" xml:"start_time,omitempty"`
	WeekDayList        []*int                                                                                                       `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRangeList []*CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem `json:"daily_time_range_list,omitempty" xml:"daily_time_range_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetEndTime(v int64) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetStartTime(v int64) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetWeekDayList(v []*int) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.WeekDayList = v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem) SetDailyTimeRangeList(v []*CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItem {
	s.DailyTimeRangeList = v
	return s
}

type CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	StartTime        *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime          *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTimeIsNextDay(v bool) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetStartTime(v string) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem) SetEndTime(v string) *CertificateQueryResponseDataCertificatesItemOffPeakDiscountInfoOffPeakTimeRangeItemDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemPeriodCard) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemPeriodCard) SetPeriodType(v int) *CertificateQueryResponseDataCertificatesItemPeriodCard {
	s.PeriodType = &v
	return s
}

type CertificateQueryResponseDataCertificatesItemReserveInfo struct {
	OrderReserveUserInfoList []*CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem `json:"order_reserve_user_info_list,omitempty" xml:"order_reserve_user_info_list,omitempty" type:"Repeated"`
}

func (s CertificateQueryResponseDataCertificatesItemReserveInfo) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemReserveInfo) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemReserveInfo) SetOrderReserveUserInfoList(v []*CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) *CertificateQueryResponseDataCertificatesItemReserveInfo {
	s.OrderReserveUserInfoList = v
	return s
}

type CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem struct {
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
	CredentialNumb *string `json:"credential_numb,omitempty" xml:"credential_numb,omitempty"`
}

func (s CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) GoString() string {
	return s.String()
}

func (s *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetCredentialType(v int) *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialType = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetName(v string) *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.Name = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetPhone(v string) *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.Phone = &v
	return s
}

func (s *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem) SetCredentialNumb(v string) *CertificateQueryResponseDataCertificatesItemReserveInfoOrderReserveUserInfoListItem {
	s.CredentialNumb = &v
	return s
}
