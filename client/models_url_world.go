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

type UrlLinkQueryQuotaRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UrlLinkQueryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryQuotaRequest) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryQuotaRequest) SetAppId(v string) *UrlLinkQueryQuotaRequest {
	s.AppId = &v
	return s
}

func (s *UrlLinkQueryQuotaRequest) SetHeader(v map[string]*string) *UrlLinkQueryQuotaRequest {
	s.Header = v
	return s
}

func (s *UrlLinkQueryQuotaRequest) SetAccessToken(v string) *UrlLinkQueryQuotaRequest {
	s.AccessToken = &v
	return s
}

type UrlLinkQueryQuotaResponse struct {
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UrlLinkQueryQuotaResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s UrlLinkQueryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryQuotaResponse) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryQuotaResponse) SetErrMsg(v string) *UrlLinkQueryQuotaResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlLinkQueryQuotaResponse) SetLogId(v string) *UrlLinkQueryQuotaResponse {
	s.LogId = &v
	return s
}

func (s *UrlLinkQueryQuotaResponse) SetData(v *UrlLinkQueryQuotaResponseData) *UrlLinkQueryQuotaResponse {
	s.Data = v
	return s
}

func (s *UrlLinkQueryQuotaResponse) SetErrNo(v int32) *UrlLinkQueryQuotaResponse {
	s.ErrNo = &v
	return s
}

type UrlLinkQueryQuotaResponseData struct {
	UrlLinkUsed  *int32 `json:"url_link_used,omitempty" xml:"url_link_used,omitempty"`
	UrlLinkLimit *int32 `json:"url_link_limit,omitempty" xml:"url_link_limit,omitempty"`
}

func (s UrlLinkQueryQuotaResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryQuotaResponseData) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryQuotaResponseData) SetUrlLinkUsed(v int32) *UrlLinkQueryQuotaResponseData {
	s.UrlLinkUsed = &v
	return s
}

func (s *UrlLinkQueryQuotaResponseData) SetUrlLinkLimit(v int32) *UrlLinkQueryQuotaResponseData {
	s.UrlLinkLimit = &v
	return s
}

type UrlQuerySchemaQuotaRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UrlQuerySchemaQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaQuotaRequest) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaQuotaRequest) SetAppId(v string) *UrlQuerySchemaQuotaRequest {
	s.AppId = &v
	return s
}

func (s *UrlQuerySchemaQuotaRequest) SetHeader(v map[string]*string) *UrlQuerySchemaQuotaRequest {
	s.Header = v
	return s
}

func (s *UrlQuerySchemaQuotaRequest) SetAccessToken(v string) *UrlQuerySchemaQuotaRequest {
	s.AccessToken = &v
	return s
}

type UrlQuerySchemaQuotaResponse struct {
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *UrlQuerySchemaQuotaResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s UrlQuerySchemaQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaQuotaResponse) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaQuotaResponse) SetErrMsg(v string) *UrlQuerySchemaQuotaResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlQuerySchemaQuotaResponse) SetLogId(v string) *UrlQuerySchemaQuotaResponse {
	s.LogId = &v
	return s
}

func (s *UrlQuerySchemaQuotaResponse) SetData(v *UrlQuerySchemaQuotaResponseData) *UrlQuerySchemaQuotaResponse {
	s.Data = v
	return s
}

func (s *UrlQuerySchemaQuotaResponse) SetErrNo(v int32) *UrlQuerySchemaQuotaResponse {
	s.ErrNo = &v
	return s
}

type UrlQuerySchemaQuotaResponseData struct {
	ShortTermSchemaQuota *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota `json:"short_term_schema_quota,omitempty" xml:"short_term_schema_quota,omitempty"`
	LongTermSchemaQuota  *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota  `json:"long_term_schema_quota,omitempty" xml:"long_term_schema_quota,omitempty"`
}

func (s UrlQuerySchemaQuotaResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaQuotaResponseData) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaQuotaResponseData) SetShortTermSchemaQuota(v *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota) *UrlQuerySchemaQuotaResponseData {
	s.ShortTermSchemaQuota = v
	return s
}

func (s *UrlQuerySchemaQuotaResponseData) SetLongTermSchemaQuota(v *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota) *UrlQuerySchemaQuotaResponseData {
	s.LongTermSchemaQuota = v
	return s
}

type UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota struct {
	SchemaLimit *int32 `json:"schema_limit,omitempty" xml:"schema_limit,omitempty"`
	SchemaUsed  *int32 `json:"schema_used,omitempty" xml:"schema_used,omitempty"`
}

func (s UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota) SetSchemaLimit(v int32) *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota {
	s.SchemaLimit = &v
	return s
}

func (s *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota) SetSchemaUsed(v int32) *UrlQuerySchemaQuotaResponseDataLongTermSchemaQuota {
	s.SchemaUsed = &v
	return s
}

type UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota struct {
	SchemaUsed  *int32 `json:"schema_used,omitempty" xml:"schema_used,omitempty"`
	SchemaLimit *int32 `json:"schema_limit,omitempty" xml:"schema_limit,omitempty"`
}

func (s UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota) SetSchemaUsed(v int32) *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota {
	s.SchemaUsed = &v
	return s
}

func (s *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota) SetSchemaLimit(v int32) *UrlQuerySchemaQuotaResponseDataShortTermSchemaQuota {
	s.SchemaLimit = &v
	return s
}

type UrlQuerySchemaRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Schema      *string            `json:"schema,omitempty" xml:"schema,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UrlQuerySchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaRequest) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaRequest) SetAppId(v string) *UrlQuerySchemaRequest {
	s.AppId = &v
	return s
}

func (s *UrlQuerySchemaRequest) SetSchema(v string) *UrlQuerySchemaRequest {
	s.Schema = &v
	return s
}

func (s *UrlQuerySchemaRequest) SetHeader(v map[string]*string) *UrlQuerySchemaRequest {
	s.Header = v
	return s
}

func (s *UrlQuerySchemaRequest) SetAccessToken(v string) *UrlQuerySchemaRequest {
	s.AccessToken = &v
	return s
}

type UrlQuerySchemaResponse struct {
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *UrlQuerySchemaResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s UrlQuerySchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaResponse) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaResponse) SetErrMsg(v string) *UrlQuerySchemaResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlQuerySchemaResponse) SetLogId(v string) *UrlQuerySchemaResponse {
	s.LogId = &v
	return s
}

func (s *UrlQuerySchemaResponse) SetData(v *UrlQuerySchemaResponseData) *UrlQuerySchemaResponse {
	s.Data = v
	return s
}

func (s *UrlQuerySchemaResponse) SetErrNo(v int32) *UrlQuerySchemaResponse {
	s.ErrNo = &v
	return s
}

type UrlQuerySchemaResponseData struct {
	ExpireTime *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	AppId      *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Path       *string `json:"path,omitempty" xml:"path,omitempty"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

func (s UrlQuerySchemaResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlQuerySchemaResponseData) GoString() string {
	return s.String()
}

func (s *UrlQuerySchemaResponseData) SetExpireTime(v int64) *UrlQuerySchemaResponseData {
	s.ExpireTime = &v
	return s
}

func (s *UrlQuerySchemaResponseData) SetAppId(v string) *UrlQuerySchemaResponseData {
	s.AppId = &v
	return s
}

func (s *UrlQuerySchemaResponseData) SetPath(v string) *UrlQuerySchemaResponseData {
	s.Path = &v
	return s
}

func (s *UrlQuerySchemaResponseData) SetQuery(v string) *UrlQuerySchemaResponseData {
	s.Query = &v
	return s
}

func (s *UrlQuerySchemaResponseData) SetCreateTime(v int64) *UrlQuerySchemaResponseData {
	s.CreateTime = &v
	return s
}

type UserBcGetCommentRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
}

func (s UserBcGetCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentRequest) SetOpenId(v string) *UserBcGetCommentRequest {
	s.OpenId = &v
	return s
}

func (s *UserBcGetCommentRequest) SetHeader(v map[string]*string) *UserBcGetCommentRequest {
	s.Header = v
	return s
}

func (s *UserBcGetCommentRequest) SetAccessToken(v string) *UserBcGetCommentRequest {
	s.AccessToken = &v
	return s
}

func (s *UserBcGetCommentRequest) SetDateType(v int64) *UserBcGetCommentRequest {
	s.DateType = &v
	return s
}

type UserBcGetCommentResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserBcGetCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UserBcGetCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentResponse) SetLogId(v string) *UserBcGetCommentResponse {
	s.LogId = &v
	return s
}

func (s *UserBcGetCommentResponse) SetData(v *UserBcGetCommentResponseData) *UserBcGetCommentResponse {
	s.Data = v
	return s
}

func (s *UserBcGetCommentResponse) SetErrNo(v int32) *UserBcGetCommentResponse {
	s.ErrNo = &v
	return s
}

func (s *UserBcGetCommentResponse) SetErrMsg(v string) *UserBcGetCommentResponse {
	s.ErrMsg = &v
	return s
}

type UserBcGetCommentResponseData struct {
	Extra *UserBcGetCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserBcGetCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserBcGetCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentResponseData) SetExtra(v *UserBcGetCommentResponseDataExtra) *UserBcGetCommentResponseData {
	s.Extra = v
	return s
}

func (s *UserBcGetCommentResponseData) SetData(v *UserBcGetCommentResponseDataData) *UserBcGetCommentResponseData {
	s.Data = v
	return s
}

type UserBcGetCommentResponseDataData struct {
	ResultList []*UserBcGetCommentResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentResponseDataData) SetResultList(v []*UserBcGetCommentResponseDataDataResultListItem) *UserBcGetCommentResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetCommentResponseDataDataResultListItem struct {
	NewComment *int64  `json:"new_comment,omitempty" xml:"new_comment,omitempty"`
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserBcGetCommentResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentResponseDataDataResultListItem) SetNewComment(v int64) *UserBcGetCommentResponseDataDataResultListItem {
	s.NewComment = &v
	return s
}

func (s *UserBcGetCommentResponseDataDataResultListItem) SetDate(v string) *UserBcGetCommentResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserBcGetCommentResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s UserBcGetCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetCommentResponseDataExtra) SetNow(v int64) *UserBcGetCommentResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetCommentResponseDataExtra) SetErrorCode(v int32) *UserBcGetCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserBcGetCommentResponseDataExtra) SetDescription(v string) *UserBcGetCommentResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetCommentResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserBcGetCommentResponseDataExtra) SetSubDescription(v string) *UserBcGetCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserBcGetCommentResponseDataExtra) SetLogid(v string) *UserBcGetCommentResponseDataExtra {
	s.Logid = &v
	return s
}

type UserBcGetFansRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s UserBcGetFansRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetFansRequest) SetAccessToken(v string) *UserBcGetFansRequest {
	s.AccessToken = &v
	return s
}

func (s *UserBcGetFansRequest) SetOpenId(v string) *UserBcGetFansRequest {
	s.OpenId = &v
	return s
}

func (s *UserBcGetFansRequest) SetDateType(v int64) *UserBcGetFansRequest {
	s.DateType = &v
	return s
}

func (s *UserBcGetFansRequest) SetHeader(v map[string]*string) *UserBcGetFansRequest {
	s.Header = v
	return s
}

type UserBcGetFansResponse struct {
	Data   *UserBcGetFansResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UserBcGetFansResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetFansResponse) SetData(v *UserBcGetFansResponseData) *UserBcGetFansResponse {
	s.Data = v
	return s
}

func (s *UserBcGetFansResponse) SetErrNo(v int32) *UserBcGetFansResponse {
	s.ErrNo = &v
	return s
}

func (s *UserBcGetFansResponse) SetErrMsg(v string) *UserBcGetFansResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserBcGetFansResponse) SetLogId(v string) *UserBcGetFansResponse {
	s.LogId = &v
	return s
}

type UserBcGetFansResponseData struct {
	Data  *UserBcGetFansResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserBcGetFansResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserBcGetFansResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetFansResponseData) SetData(v *UserBcGetFansResponseDataData) *UserBcGetFansResponseData {
	s.Data = v
	return s
}

func (s *UserBcGetFansResponseData) SetExtra(v *UserBcGetFansResponseDataExtra) *UserBcGetFansResponseData {
	s.Extra = v
	return s
}

type UserBcGetFansResponseDataData struct {
	ResultList []*UserBcGetFansResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetFansResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetFansResponseDataData) SetResultList(v []*UserBcGetFansResponseDataDataResultListItem) *UserBcGetFansResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetFansResponseDataDataResultListItem struct {
	NewFans   *int64  `json:"new_fans,omitempty" xml:"new_fans,omitempty"`
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	TotalFans *int64  `json:"total_fans,omitempty" xml:"total_fans,omitempty"`
}

func (s UserBcGetFansResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetFansResponseDataDataResultListItem) SetNewFans(v int64) *UserBcGetFansResponseDataDataResultListItem {
	s.NewFans = &v
	return s
}

func (s *UserBcGetFansResponseDataDataResultListItem) SetDate(v string) *UserBcGetFansResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserBcGetFansResponseDataDataResultListItem) SetTotalFans(v int64) *UserBcGetFansResponseDataDataResultListItem {
	s.TotalFans = &v
	return s
}

type UserBcGetFansResponseDataExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserBcGetFansResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetFansResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetFansResponseDataExtra) SetDescription(v string) *UserBcGetFansResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetFansResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetFansResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserBcGetFansResponseDataExtra) SetSubDescription(v string) *UserBcGetFansResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserBcGetFansResponseDataExtra) SetLogid(v string) *UserBcGetFansResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserBcGetFansResponseDataExtra) SetNow(v int64) *UserBcGetFansResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetFansResponseDataExtra) SetErrorCode(v int32) *UserBcGetFansResponseDataExtra {
	s.ErrorCode = &v
	return s
}

type UserBcGetItemRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s UserBcGetItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetItemRequest) SetDateType(v int64) *UserBcGetItemRequest {
	s.DateType = &v
	return s
}

func (s *UserBcGetItemRequest) SetHeader(v map[string]*string) *UserBcGetItemRequest {
	s.Header = v
	return s
}

func (s *UserBcGetItemRequest) SetAccessToken(v string) *UserBcGetItemRequest {
	s.AccessToken = &v
	return s
}

func (s *UserBcGetItemRequest) SetOpenId(v string) *UserBcGetItemRequest {
	s.OpenId = &v
	return s
}

type UserBcGetItemResponse struct {
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserBcGetItemResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UserBcGetItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetItemResponse) SetErrNo(v int32) *UserBcGetItemResponse {
	s.ErrNo = &v
	return s
}

func (s *UserBcGetItemResponse) SetErrMsg(v string) *UserBcGetItemResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserBcGetItemResponse) SetLogId(v string) *UserBcGetItemResponse {
	s.LogId = &v
	return s
}

func (s *UserBcGetItemResponse) SetData(v *UserBcGetItemResponseData) *UserBcGetItemResponse {
	s.Data = v
	return s
}

type UserBcGetItemResponseData struct {
	Data  *UserBcGetItemResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserBcGetItemResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserBcGetItemResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetItemResponseData) SetData(v *UserBcGetItemResponseDataData) *UserBcGetItemResponseData {
	s.Data = v
	return s
}

func (s *UserBcGetItemResponseData) SetExtra(v *UserBcGetItemResponseDataExtra) *UserBcGetItemResponseData {
	s.Extra = v
	return s
}

type UserBcGetItemResponseDataData struct {
	ResultList []*UserBcGetItemResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetItemResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetItemResponseDataData) SetResultList(v []*UserBcGetItemResponseDataDataResultListItem) *UserBcGetItemResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetItemResponseDataDataResultListItem struct {
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewPlay    *int64  `json:"new_play,omitempty" xml:"new_play,omitempty"`
	NewIssue   *int64  `json:"new_issue,omitempty" xml:"new_issue,omitempty"`
	TotalIssue *int64  `json:"total_issue,omitempty" xml:"total_issue,omitempty"`
}

func (s UserBcGetItemResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetItemResponseDataDataResultListItem) SetDate(v string) *UserBcGetItemResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserBcGetItemResponseDataDataResultListItem) SetNewPlay(v int64) *UserBcGetItemResponseDataDataResultListItem {
	s.NewPlay = &v
	return s
}

func (s *UserBcGetItemResponseDataDataResultListItem) SetNewIssue(v int64) *UserBcGetItemResponseDataDataResultListItem {
	s.NewIssue = &v
	return s
}

func (s *UserBcGetItemResponseDataDataResultListItem) SetTotalIssue(v int64) *UserBcGetItemResponseDataDataResultListItem {
	s.TotalIssue = &v
	return s
}

type UserBcGetItemResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s UserBcGetItemResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetItemResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetItemResponseDataExtra) SetNow(v int64) *UserBcGetItemResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetItemResponseDataExtra) SetErrorCode(v int32) *UserBcGetItemResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserBcGetItemResponseDataExtra) SetDescription(v string) *UserBcGetItemResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetItemResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetItemResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserBcGetItemResponseDataExtra) SetSubDescription(v string) *UserBcGetItemResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserBcGetItemResponseDataExtra) SetLogid(v string) *UserBcGetItemResponseDataExtra {
	s.Logid = &v
	return s
}

type UserBcGetLikeRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserBcGetLikeRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeRequest) SetOpenId(v string) *UserBcGetLikeRequest {
	s.OpenId = &v
	return s
}

func (s *UserBcGetLikeRequest) SetDateType(v int64) *UserBcGetLikeRequest {
	s.DateType = &v
	return s
}

func (s *UserBcGetLikeRequest) SetHeader(v map[string]*string) *UserBcGetLikeRequest {
	s.Header = v
	return s
}

func (s *UserBcGetLikeRequest) SetAccessToken(v string) *UserBcGetLikeRequest {
	s.AccessToken = &v
	return s
}

type UserBcGetLikeResponse struct {
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserBcGetLikeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UserBcGetLikeResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeResponse) SetErrNo(v int32) *UserBcGetLikeResponse {
	s.ErrNo = &v
	return s
}

func (s *UserBcGetLikeResponse) SetErrMsg(v string) *UserBcGetLikeResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserBcGetLikeResponse) SetLogId(v string) *UserBcGetLikeResponse {
	s.LogId = &v
	return s
}

func (s *UserBcGetLikeResponse) SetData(v *UserBcGetLikeResponseData) *UserBcGetLikeResponse {
	s.Data = v
	return s
}

type UserBcGetLikeResponseData struct {
	Extra *UserBcGetLikeResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserBcGetLikeResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserBcGetLikeResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeResponseData) SetExtra(v *UserBcGetLikeResponseDataExtra) *UserBcGetLikeResponseData {
	s.Extra = v
	return s
}

func (s *UserBcGetLikeResponseData) SetData(v *UserBcGetLikeResponseDataData) *UserBcGetLikeResponseData {
	s.Data = v
	return s
}

type UserBcGetLikeResponseDataData struct {
	ResultList []*UserBcGetLikeResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetLikeResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeResponseDataData) SetResultList(v []*UserBcGetLikeResponseDataDataResultListItem) *UserBcGetLikeResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetLikeResponseDataDataResultListItem struct {
	Date    *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewLike *int64  `json:"new_like,omitempty" xml:"new_like,omitempty"`
}

func (s UserBcGetLikeResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeResponseDataDataResultListItem) SetDate(v string) *UserBcGetLikeResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserBcGetLikeResponseDataDataResultListItem) SetNewLike(v int64) *UserBcGetLikeResponseDataDataResultListItem {
	s.NewLike = &v
	return s
}

type UserBcGetLikeResponseDataExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s UserBcGetLikeResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetLikeResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetLikeResponseDataExtra) SetLogid(v string) *UserBcGetLikeResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserBcGetLikeResponseDataExtra) SetNow(v int64) *UserBcGetLikeResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetLikeResponseDataExtra) SetErrorCode(v int32) *UserBcGetLikeResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserBcGetLikeResponseDataExtra) SetDescription(v string) *UserBcGetLikeResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetLikeResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetLikeResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserBcGetLikeResponseDataExtra) SetSubDescription(v string) *UserBcGetLikeResponseDataExtra {
	s.SubDescription = &v
	return s
}

type UserBcGetProfileRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserBcGetProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileRequest) SetOpenId(v string) *UserBcGetProfileRequest {
	s.OpenId = &v
	return s
}

func (s *UserBcGetProfileRequest) SetDateType(v int64) *UserBcGetProfileRequest {
	s.DateType = &v
	return s
}

func (s *UserBcGetProfileRequest) SetHeader(v map[string]*string) *UserBcGetProfileRequest {
	s.Header = v
	return s
}

func (s *UserBcGetProfileRequest) SetAccessToken(v string) *UserBcGetProfileRequest {
	s.AccessToken = &v
	return s
}

type UserBcGetProfileResponse struct {
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserBcGetProfileResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s UserBcGetProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileResponse) SetErrMsg(v string) *UserBcGetProfileResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserBcGetProfileResponse) SetLogId(v string) *UserBcGetProfileResponse {
	s.LogId = &v
	return s
}

func (s *UserBcGetProfileResponse) SetData(v *UserBcGetProfileResponseData) *UserBcGetProfileResponse {
	s.Data = v
	return s
}

func (s *UserBcGetProfileResponse) SetErrNo(v int32) *UserBcGetProfileResponse {
	s.ErrNo = &v
	return s
}

type UserBcGetProfileResponseData struct {
	Extra *UserBcGetProfileResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserBcGetProfileResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserBcGetProfileResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileResponseData) SetExtra(v *UserBcGetProfileResponseDataExtra) *UserBcGetProfileResponseData {
	s.Extra = v
	return s
}

func (s *UserBcGetProfileResponseData) SetData(v *UserBcGetProfileResponseDataData) *UserBcGetProfileResponseData {
	s.Data = v
	return s
}

type UserBcGetProfileResponseDataData struct {
	ResultList []*UserBcGetProfileResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetProfileResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileResponseDataData) SetResultList(v []*UserBcGetProfileResponseDataDataResultListItem) *UserBcGetProfileResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetProfileResponseDataDataResultListItem struct {
	ProfileUv *int64  `json:"profile_uv,omitempty" xml:"profile_uv,omitempty"`
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserBcGetProfileResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileResponseDataDataResultListItem) SetProfileUv(v int64) *UserBcGetProfileResponseDataDataResultListItem {
	s.ProfileUv = &v
	return s
}

func (s *UserBcGetProfileResponseDataDataResultListItem) SetDate(v string) *UserBcGetProfileResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserBcGetProfileResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s UserBcGetProfileResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetProfileResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetProfileResponseDataExtra) SetSubDescription(v string) *UserBcGetProfileResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserBcGetProfileResponseDataExtra) SetLogid(v string) *UserBcGetProfileResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserBcGetProfileResponseDataExtra) SetNow(v int64) *UserBcGetProfileResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetProfileResponseDataExtra) SetErrorCode(v int32) *UserBcGetProfileResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserBcGetProfileResponseDataExtra) SetDescription(v string) *UserBcGetProfileResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetProfileResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetProfileResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type UserBcGetShareRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserBcGetShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareRequest) GoString() string {
	return s.String()
}

func (s *UserBcGetShareRequest) SetDateType(v int64) *UserBcGetShareRequest {
	s.DateType = &v
	return s
}

func (s *UserBcGetShareRequest) SetOpenId(v string) *UserBcGetShareRequest {
	s.OpenId = &v
	return s
}

func (s *UserBcGetShareRequest) SetHeader(v map[string]*string) *UserBcGetShareRequest {
	s.Header = v
	return s
}

func (s *UserBcGetShareRequest) SetAccessToken(v string) *UserBcGetShareRequest {
	s.AccessToken = &v
	return s
}

type UserBcGetShareResponse struct {
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserBcGetShareResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UserBcGetShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareResponse) GoString() string {
	return s.String()
}

func (s *UserBcGetShareResponse) SetErrNo(v int32) *UserBcGetShareResponse {
	s.ErrNo = &v
	return s
}

func (s *UserBcGetShareResponse) SetErrMsg(v string) *UserBcGetShareResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserBcGetShareResponse) SetLogId(v string) *UserBcGetShareResponse {
	s.LogId = &v
	return s
}

func (s *UserBcGetShareResponse) SetData(v *UserBcGetShareResponseData) *UserBcGetShareResponse {
	s.Data = v
	return s
}

type UserBcGetShareResponseData struct {
	Extra *UserBcGetShareResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserBcGetShareResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserBcGetShareResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareResponseData) GoString() string {
	return s.String()
}

func (s *UserBcGetShareResponseData) SetExtra(v *UserBcGetShareResponseDataExtra) *UserBcGetShareResponseData {
	s.Extra = v
	return s
}

func (s *UserBcGetShareResponseData) SetData(v *UserBcGetShareResponseDataData) *UserBcGetShareResponseData {
	s.Data = v
	return s
}

type UserBcGetShareResponseDataData struct {
	ResultList []*UserBcGetShareResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserBcGetShareResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareResponseDataData) GoString() string {
	return s.String()
}

func (s *UserBcGetShareResponseDataData) SetResultList(v []*UserBcGetShareResponseDataDataResultListItem) *UserBcGetShareResponseDataData {
	s.ResultList = v
	return s
}

type UserBcGetShareResponseDataDataResultListItem struct {
	NewShare *int64  `json:"new_share,omitempty" xml:"new_share,omitempty"`
	Date     *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserBcGetShareResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserBcGetShareResponseDataDataResultListItem) SetNewShare(v int64) *UserBcGetShareResponseDataDataResultListItem {
	s.NewShare = &v
	return s
}

func (s *UserBcGetShareResponseDataDataResultListItem) SetDate(v string) *UserBcGetShareResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserBcGetShareResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s UserBcGetShareResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserBcGetShareResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserBcGetShareResponseDataExtra) SetNow(v int64) *UserBcGetShareResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserBcGetShareResponseDataExtra) SetErrorCode(v int32) *UserBcGetShareResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserBcGetShareResponseDataExtra) SetDescription(v string) *UserBcGetShareResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserBcGetShareResponseDataExtra) SetSubErrorCode(v int32) *UserBcGetShareResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserBcGetShareResponseDataExtra) SetSubDescription(v string) *UserBcGetShareResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserBcGetShareResponseDataExtra) SetLogid(v string) *UserBcGetShareResponseDataExtra {
	s.Logid = &v
	return s
}

type UserChangeRequest struct {
	OpenId       *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	AccountId    *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	MemberCardId *string            `json:"member_card_id,omitempty" xml:"member_card_id,omitempty"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s UserChangeRequest) GoString() string {
	return s.String()
}

func (s *UserChangeRequest) SetOpenId(v string) *UserChangeRequest {
	s.OpenId = &v
	return s
}

func (s *UserChangeRequest) SetAccountId(v string) *UserChangeRequest {
	s.AccountId = &v
	return s
}

func (s *UserChangeRequest) SetMemberCardId(v string) *UserChangeRequest {
	s.MemberCardId = &v
	return s
}

func (s *UserChangeRequest) SetHeader(v map[string]*string) *UserChangeRequest {
	s.Header = v
	return s
}

func (s *UserChangeRequest) SetAccessToken(v string) *UserChangeRequest {
	s.AccessToken = &v
	return s
}

type UserChangeResponse struct {
	Extra *UserChangeResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserChangeResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UserChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s UserChangeResponse) GoString() string {
	return s.String()
}

func (s *UserChangeResponse) SetExtra(v *UserChangeResponseExtra) *UserChangeResponse {
	s.Extra = v
	return s
}

func (s *UserChangeResponse) SetData(v *UserChangeResponseData) *UserChangeResponse {
	s.Data = v
	return s
}

type UserChangeResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserChangeResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserChangeResponseData) GoString() string {
	return s.String()
}

func (s *UserChangeResponseData) SetGwErrorCode(v int32) *UserChangeResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserChangeResponseData) SetGwDescription(v string) *UserChangeResponseData {
	s.GwDescription = &v
	return s
}

type UserChangeResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserChangeResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserChangeResponseExtra) GoString() string {
	return s.String()
}

func (s *UserChangeResponseExtra) SetLogid(v string) *UserChangeResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserChangeResponseExtra) SetNow(v int64) *UserChangeResponseExtra {
	s.Now = &v
	return s
}

func (s *UserChangeResponseExtra) SetSubDescription(v string) *UserChangeResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserChangeResponseExtra) SetSubErrorCode(v int32) *UserChangeResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserChangeResponseExtra) SetDescription(v string) *UserChangeResponseExtra {
	s.Description = &v
	return s
}

func (s *UserChangeResponseExtra) SetErrorCode(v int32) *UserChangeResponseExtra {
	s.ErrorCode = &v
	return s
}

type UserCommentRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s UserCommentRequest) GoString() string {
	return s.String()
}

func (s *UserCommentRequest) SetDateType(v int64) *UserCommentRequest {
	s.DateType = &v
	return s
}

func (s *UserCommentRequest) SetOpenId(v string) *UserCommentRequest {
	s.OpenId = &v
	return s
}

func (s *UserCommentRequest) SetHeader(v map[string]*string) *UserCommentRequest {
	s.Header = v
	return s
}

func (s *UserCommentRequest) SetAccessToken(v string) *UserCommentRequest {
	s.AccessToken = &v
	return s
}

type UserCommentResponse struct {
	Extra *UserCommentResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserCommentResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s UserCommentResponse) GoString() string {
	return s.String()
}

func (s *UserCommentResponse) SetExtra(v *UserCommentResponseExtra) *UserCommentResponse {
	s.Extra = v
	return s
}

func (s *UserCommentResponse) SetData(v *UserCommentResponseData) *UserCommentResponse {
	s.Data = v
	return s
}

type UserCommentResponseData struct {
	ResultList    []*UserCommentResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserCommentResponseData) GoString() string {
	return s.String()
}

func (s *UserCommentResponseData) SetResultList(v []*UserCommentResponseDataResultListItem) *UserCommentResponseData {
	s.ResultList = v
	return s
}

func (s *UserCommentResponseData) SetGwErrorCode(v int32) *UserCommentResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserCommentResponseData) SetGwDescription(v string) *UserCommentResponseData {
	s.GwDescription = &v
	return s
}

type UserCommentResponseDataResultListItem struct {
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewComment *int64  `json:"new_comment,omitempty" xml:"new_comment,omitempty"`
}

func (s UserCommentResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserCommentResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserCommentResponseDataResultListItem) SetDate(v string) *UserCommentResponseDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserCommentResponseDataResultListItem) SetNewComment(v int64) *UserCommentResponseDataResultListItem {
	s.NewComment = &v
	return s
}

type UserCommentResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserCommentResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserCommentResponseExtra) GoString() string {
	return s.String()
}

func (s *UserCommentResponseExtra) SetSubErrorCode(v int32) *UserCommentResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserCommentResponseExtra) SetSubDescription(v string) *UserCommentResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserCommentResponseExtra) SetLogid(v string) *UserCommentResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserCommentResponseExtra) SetNow(v int64) *UserCommentResponseExtra {
	s.Now = &v
	return s
}

func (s *UserCommentResponseExtra) SetErrorCode(v int32) *UserCommentResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserCommentResponseExtra) SetDescription(v string) *UserCommentResponseExtra {
	s.Description = &v
	return s
}

type UserFansDataRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserFansDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataRequest) GoString() string {
	return s.String()
}

func (s *UserFansDataRequest) SetOpenId(v string) *UserFansDataRequest {
	s.OpenId = &v
	return s
}

func (s *UserFansDataRequest) SetHeader(v map[string]*string) *UserFansDataRequest {
	s.Header = v
	return s
}

func (s *UserFansDataRequest) SetAccessToken(v string) *UserFansDataRequest {
	s.AccessToken = &v
	return s
}

type UserFansDataResponse struct {
	Extra *UserFansDataResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserFansDataResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserFansDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponse) GoString() string {
	return s.String()
}

func (s *UserFansDataResponse) SetExtra(v *UserFansDataResponseExtra) *UserFansDataResponse {
	s.Extra = v
	return s
}

func (s *UserFansDataResponse) SetData(v *UserFansDataResponseData) *UserFansDataResponse {
	s.Data = v
	return s
}

type UserFansDataResponseData struct {
	GwDescription *string                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	FansData      *UserFansDataResponseDataFansData `json:"fans_data,omitempty" xml:"fans_data,omitempty" require:"true"`
	GwErrorCode   *int32                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserFansDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseData) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseData) SetGwDescription(v string) *UserFansDataResponseData {
	s.GwDescription = &v
	return s
}

func (s *UserFansDataResponseData) SetFansData(v *UserFansDataResponseDataFansData) *UserFansDataResponseData {
	s.FansData = v
	return s
}

func (s *UserFansDataResponseData) SetGwErrorCode(v int32) *UserFansDataResponseData {
	s.GwErrorCode = &v
	return s
}

type UserFansDataResponseDataFansData struct {
	FlowContributions         []*UserFansDataResponseDataFansDataFlowContributionsItem         `json:"flow_contributions,omitempty" xml:"flow_contributions,omitempty" type:"Repeated"`
	GenderDistributions       []*UserFansDataResponseDataFansDataGenderDistributionsItem       `json:"gender_distributions,omitempty" xml:"gender_distributions,omitempty" type:"Repeated"`
	GeographicalDistributions []*UserFansDataResponseDataFansDataGeographicalDistributionsItem `json:"geographical_distributions,omitempty" xml:"geographical_distributions,omitempty" type:"Repeated"`
	InterestDistributions     []*UserFansDataResponseDataFansDataInterestDistributionsItem     `json:"interest_distributions,omitempty" xml:"interest_distributions,omitempty" type:"Repeated"`
	ActiveDaysDistributions   []*UserFansDataResponseDataFansDataActiveDaysDistributionsItem   `json:"active_days_distributions,omitempty" xml:"active_days_distributions,omitempty" type:"Repeated"`
	AgeDistributions          []*UserFansDataResponseDataFansDataAgeDistributionsItem          `json:"age_distributions,omitempty" xml:"age_distributions,omitempty" type:"Repeated"`
	AllFansNum                *int32                                                           `json:"all_fans_num,omitempty" xml:"all_fans_num,omitempty" require:"true"`
	DeviceDistributions       []*UserFansDataResponseDataFansDataDeviceDistributionsItem       `json:"device_distributions,omitempty" xml:"device_distributions,omitempty" type:"Repeated"`
}

func (s UserFansDataResponseDataFansData) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansData) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansData) SetFlowContributions(v []*UserFansDataResponseDataFansDataFlowContributionsItem) *UserFansDataResponseDataFansData {
	s.FlowContributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetGenderDistributions(v []*UserFansDataResponseDataFansDataGenderDistributionsItem) *UserFansDataResponseDataFansData {
	s.GenderDistributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetGeographicalDistributions(v []*UserFansDataResponseDataFansDataGeographicalDistributionsItem) *UserFansDataResponseDataFansData {
	s.GeographicalDistributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetInterestDistributions(v []*UserFansDataResponseDataFansDataInterestDistributionsItem) *UserFansDataResponseDataFansData {
	s.InterestDistributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetActiveDaysDistributions(v []*UserFansDataResponseDataFansDataActiveDaysDistributionsItem) *UserFansDataResponseDataFansData {
	s.ActiveDaysDistributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetAgeDistributions(v []*UserFansDataResponseDataFansDataAgeDistributionsItem) *UserFansDataResponseDataFansData {
	s.AgeDistributions = v
	return s
}

func (s *UserFansDataResponseDataFansData) SetAllFansNum(v int32) *UserFansDataResponseDataFansData {
	s.AllFansNum = &v
	return s
}

func (s *UserFansDataResponseDataFansData) SetDeviceDistributions(v []*UserFansDataResponseDataFansDataDeviceDistributionsItem) *UserFansDataResponseDataFansData {
	s.DeviceDistributions = v
	return s
}

type UserFansDataResponseDataFansDataActiveDaysDistributionsItem struct {
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataActiveDaysDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataActiveDaysDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataActiveDaysDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataActiveDaysDistributionsItem {
	s.Item = &v
	return s
}

func (s *UserFansDataResponseDataFansDataActiveDaysDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataActiveDaysDistributionsItem {
	s.Value = &v
	return s
}

type UserFansDataResponseDataFansDataAgeDistributionsItem struct {
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataAgeDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataAgeDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataAgeDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataAgeDistributionsItem {
	s.Value = &v
	return s
}

func (s *UserFansDataResponseDataFansDataAgeDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataAgeDistributionsItem {
	s.Item = &v
	return s
}

type UserFansDataResponseDataFansDataDeviceDistributionsItem struct {
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataDeviceDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataDeviceDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataDeviceDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataDeviceDistributionsItem {
	s.Item = &v
	return s
}

func (s *UserFansDataResponseDataFansDataDeviceDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataDeviceDistributionsItem {
	s.Value = &v
	return s
}

type UserFansDataResponseDataFansDataFlowContributionsItem struct {
	Flow    *string `json:"flow,omitempty" xml:"flow,omitempty" require:"true"`
	AllSum  *int64  `json:"all_sum,omitempty" xml:"all_sum,omitempty" require:"true"`
	FansSum *int64  `json:"fans_sum,omitempty" xml:"fans_sum,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataFlowContributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataFlowContributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataFlowContributionsItem) SetFlow(v string) *UserFansDataResponseDataFansDataFlowContributionsItem {
	s.Flow = &v
	return s
}

func (s *UserFansDataResponseDataFansDataFlowContributionsItem) SetAllSum(v int64) *UserFansDataResponseDataFansDataFlowContributionsItem {
	s.AllSum = &v
	return s
}

func (s *UserFansDataResponseDataFansDataFlowContributionsItem) SetFansSum(v int64) *UserFansDataResponseDataFansDataFlowContributionsItem {
	s.FansSum = &v
	return s
}

type UserFansDataResponseDataFansDataGenderDistributionsItem struct {
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataGenderDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataGenderDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataGenderDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataGenderDistributionsItem {
	s.Value = &v
	return s
}

func (s *UserFansDataResponseDataFansDataGenderDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataGenderDistributionsItem {
	s.Item = &v
	return s
}

type UserFansDataResponseDataFansDataGeographicalDistributionsItem struct {
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataGeographicalDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataGeographicalDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataGeographicalDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataGeographicalDistributionsItem {
	s.Item = &v
	return s
}

func (s *UserFansDataResponseDataFansDataGeographicalDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataGeographicalDistributionsItem {
	s.Value = &v
	return s
}

type UserFansDataResponseDataFansDataInterestDistributionsItem struct {
	Value *int64  `json:"value,omitempty" xml:"value,omitempty" require:"true"`
	Item  *string `json:"item,omitempty" xml:"item,omitempty" require:"true"`
}

func (s UserFansDataResponseDataFansDataInterestDistributionsItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseDataFansDataInterestDistributionsItem) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseDataFansDataInterestDistributionsItem) SetValue(v int64) *UserFansDataResponseDataFansDataInterestDistributionsItem {
	s.Value = &v
	return s
}

func (s *UserFansDataResponseDataFansDataInterestDistributionsItem) SetItem(v string) *UserFansDataResponseDataFansDataInterestDistributionsItem {
	s.Item = &v
	return s
}

type UserFansDataResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s UserFansDataResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserFansDataResponseExtra) GoString() string {
	return s.String()
}

func (s *UserFansDataResponseExtra) SetDescription(v string) *UserFansDataResponseExtra {
	s.Description = &v
	return s
}

func (s *UserFansDataResponseExtra) SetErrorCode(v int32) *UserFansDataResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserFansDataResponseExtra) SetLogid(v string) *UserFansDataResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserFansDataResponseExtra) SetNow(v int64) *UserFansDataResponseExtra {
	s.Now = &v
	return s
}

func (s *UserFansDataResponseExtra) SetSubDescription(v string) *UserFansDataResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserFansDataResponseExtra) SetSubErrorCode(v int32) *UserFansDataResponseExtra {
	s.SubErrorCode = &v
	return s
}

type UserFansRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserFansRequest) String() string {
	return tea.Prettify(s)
}

func (s UserFansRequest) GoString() string {
	return s.String()
}

func (s *UserFansRequest) SetOpenId(v string) *UserFansRequest {
	s.OpenId = &v
	return s
}

func (s *UserFansRequest) SetDateType(v int64) *UserFansRequest {
	s.DateType = &v
	return s
}

func (s *UserFansRequest) SetHeader(v map[string]*string) *UserFansRequest {
	s.Header = v
	return s
}

func (s *UserFansRequest) SetAccessToken(v string) *UserFansRequest {
	s.AccessToken = &v
	return s
}

type UserFansResponse struct {
	Extra *UserFansResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserFansResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserFansResponse) String() string {
	return tea.Prettify(s)
}

func (s UserFansResponse) GoString() string {
	return s.String()
}

func (s *UserFansResponse) SetExtra(v *UserFansResponseExtra) *UserFansResponse {
	s.Extra = v
	return s
}

func (s *UserFansResponse) SetData(v *UserFansResponseData) *UserFansResponse {
	s.Data = v
	return s
}

type UserFansResponseData struct {
	ResultList    []*UserFansResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserFansResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserFansResponseData) GoString() string {
	return s.String()
}

func (s *UserFansResponseData) SetResultList(v []*UserFansResponseDataResultListItem) *UserFansResponseData {
	s.ResultList = v
	return s
}

func (s *UserFansResponseData) SetGwErrorCode(v int32) *UserFansResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserFansResponseData) SetGwDescription(v string) *UserFansResponseData {
	s.GwDescription = &v
	return s
}

type UserFansResponseDataResultListItem struct {
	NewFans   *int64  `json:"new_fans,omitempty" xml:"new_fans,omitempty"`
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	TotalFans *int64  `json:"total_fans,omitempty" xml:"total_fans,omitempty"`
}

func (s UserFansResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserFansResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserFansResponseDataResultListItem) SetNewFans(v int64) *UserFansResponseDataResultListItem {
	s.NewFans = &v
	return s
}

func (s *UserFansResponseDataResultListItem) SetDate(v string) *UserFansResponseDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserFansResponseDataResultListItem) SetTotalFans(v int64) *UserFansResponseDataResultListItem {
	s.TotalFans = &v
	return s
}

type UserFansResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserFansResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserFansResponseExtra) GoString() string {
	return s.String()
}

func (s *UserFansResponseExtra) SetDescription(v string) *UserFansResponseExtra {
	s.Description = &v
	return s
}

func (s *UserFansResponseExtra) SetSubErrorCode(v int32) *UserFansResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserFansResponseExtra) SetSubDescription(v string) *UserFansResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserFansResponseExtra) SetLogid(v string) *UserFansResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserFansResponseExtra) SetNow(v int64) *UserFansResponseExtra {
	s.Now = &v
	return s
}

func (s *UserFansResponseExtra) SetErrorCode(v int32) *UserFansResponseExtra {
	s.ErrorCode = &v
	return s
}

type UserGetCommentRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserGetCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentRequest) GoString() string {
	return s.String()
}

func (s *UserGetCommentRequest) SetOpenId(v string) *UserGetCommentRequest {
	s.OpenId = &v
	return s
}

func (s *UserGetCommentRequest) SetDateType(v int64) *UserGetCommentRequest {
	s.DateType = &v
	return s
}

func (s *UserGetCommentRequest) SetHeader(v map[string]*string) *UserGetCommentRequest {
	s.Header = v
	return s
}

func (s *UserGetCommentRequest) SetAccessToken(v string) *UserGetCommentRequest {
	s.AccessToken = &v
	return s
}

type UserGetCommentResponse struct {
	Data   *UserGetCommentResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UserGetCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentResponse) GoString() string {
	return s.String()
}

func (s *UserGetCommentResponse) SetData(v *UserGetCommentResponseData) *UserGetCommentResponse {
	s.Data = v
	return s
}

func (s *UserGetCommentResponse) SetErrNo(v int32) *UserGetCommentResponse {
	s.ErrNo = &v
	return s
}

func (s *UserGetCommentResponse) SetErrMsg(v string) *UserGetCommentResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserGetCommentResponse) SetLogId(v string) *UserGetCommentResponse {
	s.LogId = &v
	return s
}

type UserGetCommentResponseData struct {
	Data  *UserGetCommentResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserGetCommentResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserGetCommentResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentResponseData) GoString() string {
	return s.String()
}

func (s *UserGetCommentResponseData) SetData(v *UserGetCommentResponseDataData) *UserGetCommentResponseData {
	s.Data = v
	return s
}

func (s *UserGetCommentResponseData) SetExtra(v *UserGetCommentResponseDataExtra) *UserGetCommentResponseData {
	s.Extra = v
	return s
}

type UserGetCommentResponseDataData struct {
	ResultList []*UserGetCommentResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetCommentResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetCommentResponseDataData) SetResultList(v []*UserGetCommentResponseDataDataResultListItem) *UserGetCommentResponseDataData {
	s.ResultList = v
	return s
}

type UserGetCommentResponseDataDataResultListItem struct {
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewComment *int64  `json:"new_comment,omitempty" xml:"new_comment,omitempty"`
}

func (s UserGetCommentResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetCommentResponseDataDataResultListItem) SetDate(v string) *UserGetCommentResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserGetCommentResponseDataDataResultListItem) SetNewComment(v int64) *UserGetCommentResponseDataDataResultListItem {
	s.NewComment = &v
	return s
}

type UserGetCommentResponseDataExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s UserGetCommentResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetCommentResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetCommentResponseDataExtra) SetErrorCode(v int32) *UserGetCommentResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserGetCommentResponseDataExtra) SetDescription(v string) *UserGetCommentResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserGetCommentResponseDataExtra) SetSubErrorCode(v int32) *UserGetCommentResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetCommentResponseDataExtra) SetSubDescription(v string) *UserGetCommentResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserGetCommentResponseDataExtra) SetLogid(v string) *UserGetCommentResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetCommentResponseDataExtra) SetNow(v int64) *UserGetCommentResponseDataExtra {
	s.Now = &v
	return s
}

type UserGetFansRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserGetFansRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansRequest) GoString() string {
	return s.String()
}

func (s *UserGetFansRequest) SetOpenId(v string) *UserGetFansRequest {
	s.OpenId = &v
	return s
}

func (s *UserGetFansRequest) SetDateType(v int64) *UserGetFansRequest {
	s.DateType = &v
	return s
}

func (s *UserGetFansRequest) SetHeader(v map[string]*string) *UserGetFansRequest {
	s.Header = v
	return s
}

func (s *UserGetFansRequest) SetAccessToken(v string) *UserGetFansRequest {
	s.AccessToken = &v
	return s
}

type UserGetFansResponse struct {
	Data   *UserGetFansResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UserGetFansResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansResponse) GoString() string {
	return s.String()
}

func (s *UserGetFansResponse) SetData(v *UserGetFansResponseData) *UserGetFansResponse {
	s.Data = v
	return s
}

func (s *UserGetFansResponse) SetErrNo(v int32) *UserGetFansResponse {
	s.ErrNo = &v
	return s
}

func (s *UserGetFansResponse) SetErrMsg(v string) *UserGetFansResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserGetFansResponse) SetLogId(v string) *UserGetFansResponse {
	s.LogId = &v
	return s
}

type UserGetFansResponseData struct {
	Data  *UserGetFansResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserGetFansResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserGetFansResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansResponseData) GoString() string {
	return s.String()
}

func (s *UserGetFansResponseData) SetData(v *UserGetFansResponseDataData) *UserGetFansResponseData {
	s.Data = v
	return s
}

func (s *UserGetFansResponseData) SetExtra(v *UserGetFansResponseDataExtra) *UserGetFansResponseData {
	s.Extra = v
	return s
}

type UserGetFansResponseDataData struct {
	ResultList []*UserGetFansResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetFansResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetFansResponseDataData) SetResultList(v []*UserGetFansResponseDataDataResultListItem) *UserGetFansResponseDataData {
	s.ResultList = v
	return s
}

type UserGetFansResponseDataDataResultListItem struct {
	TotalFans *int64  `json:"total_fans,omitempty" xml:"total_fans,omitempty"`
	NewFans   *int64  `json:"new_fans,omitempty" xml:"new_fans,omitempty"`
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserGetFansResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetFansResponseDataDataResultListItem) SetTotalFans(v int64) *UserGetFansResponseDataDataResultListItem {
	s.TotalFans = &v
	return s
}

func (s *UserGetFansResponseDataDataResultListItem) SetNewFans(v int64) *UserGetFansResponseDataDataResultListItem {
	s.NewFans = &v
	return s
}

func (s *UserGetFansResponseDataDataResultListItem) SetDate(v string) *UserGetFansResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserGetFansResponseDataExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s UserGetFansResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetFansResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetFansResponseDataExtra) SetLogid(v string) *UserGetFansResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetFansResponseDataExtra) SetNow(v int64) *UserGetFansResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserGetFansResponseDataExtra) SetErrorCode(v int32) *UserGetFansResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserGetFansResponseDataExtra) SetDescription(v string) *UserGetFansResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserGetFansResponseDataExtra) SetSubErrorCode(v int32) *UserGetFansResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetFansResponseDataExtra) SetSubDescription(v string) *UserGetFansResponseDataExtra {
	s.SubDescription = &v
	return s
}

type UserGetItemRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserGetItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemRequest) GoString() string {
	return s.String()
}

func (s *UserGetItemRequest) SetOpenId(v string) *UserGetItemRequest {
	s.OpenId = &v
	return s
}

func (s *UserGetItemRequest) SetDateType(v int64) *UserGetItemRequest {
	s.DateType = &v
	return s
}

func (s *UserGetItemRequest) SetHeader(v map[string]*string) *UserGetItemRequest {
	s.Header = v
	return s
}

func (s *UserGetItemRequest) SetAccessToken(v string) *UserGetItemRequest {
	s.AccessToken = &v
	return s
}

type UserGetItemResponse struct {
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserGetItemResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UserGetItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemResponse) GoString() string {
	return s.String()
}

func (s *UserGetItemResponse) SetLogId(v string) *UserGetItemResponse {
	s.LogId = &v
	return s
}

func (s *UserGetItemResponse) SetData(v *UserGetItemResponseData) *UserGetItemResponse {
	s.Data = v
	return s
}

func (s *UserGetItemResponse) SetErrNo(v int32) *UserGetItemResponse {
	s.ErrNo = &v
	return s
}

func (s *UserGetItemResponse) SetErrMsg(v string) *UserGetItemResponse {
	s.ErrMsg = &v
	return s
}

type UserGetItemResponseData struct {
	Extra *UserGetItemResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserGetItemResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserGetItemResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemResponseData) GoString() string {
	return s.String()
}

func (s *UserGetItemResponseData) SetExtra(v *UserGetItemResponseDataExtra) *UserGetItemResponseData {
	s.Extra = v
	return s
}

func (s *UserGetItemResponseData) SetData(v *UserGetItemResponseDataData) *UserGetItemResponseData {
	s.Data = v
	return s
}

type UserGetItemResponseDataData struct {
	ResultList []*UserGetItemResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetItemResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetItemResponseDataData) SetResultList(v []*UserGetItemResponseDataDataResultListItem) *UserGetItemResponseDataData {
	s.ResultList = v
	return s
}

type UserGetItemResponseDataDataResultListItem struct {
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewPlay    *int64  `json:"new_play,omitempty" xml:"new_play,omitempty"`
	NewIssue   *int64  `json:"new_issue,omitempty" xml:"new_issue,omitempty"`
	TotalIssue *int64  `json:"total_issue,omitempty" xml:"total_issue,omitempty"`
}

func (s UserGetItemResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetItemResponseDataDataResultListItem) SetDate(v string) *UserGetItemResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserGetItemResponseDataDataResultListItem) SetNewPlay(v int64) *UserGetItemResponseDataDataResultListItem {
	s.NewPlay = &v
	return s
}

func (s *UserGetItemResponseDataDataResultListItem) SetNewIssue(v int64) *UserGetItemResponseDataDataResultListItem {
	s.NewIssue = &v
	return s
}

func (s *UserGetItemResponseDataDataResultListItem) SetTotalIssue(v int64) *UserGetItemResponseDataDataResultListItem {
	s.TotalIssue = &v
	return s
}

type UserGetItemResponseDataExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserGetItemResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetItemResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetItemResponseDataExtra) SetSubErrorCode(v int32) *UserGetItemResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetItemResponseDataExtra) SetSubDescription(v string) *UserGetItemResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserGetItemResponseDataExtra) SetLogid(v string) *UserGetItemResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetItemResponseDataExtra) SetNow(v int64) *UserGetItemResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserGetItemResponseDataExtra) SetErrorCode(v int32) *UserGetItemResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserGetItemResponseDataExtra) SetDescription(v string) *UserGetItemResponseDataExtra {
	s.Description = &v
	return s
}

type UserGetLikeRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s UserGetLikeRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeRequest) GoString() string {
	return s.String()
}

func (s *UserGetLikeRequest) SetDateType(v int64) *UserGetLikeRequest {
	s.DateType = &v
	return s
}

func (s *UserGetLikeRequest) SetHeader(v map[string]*string) *UserGetLikeRequest {
	s.Header = v
	return s
}

func (s *UserGetLikeRequest) SetAccessToken(v string) *UserGetLikeRequest {
	s.AccessToken = &v
	return s
}

func (s *UserGetLikeRequest) SetOpenId(v string) *UserGetLikeRequest {
	s.OpenId = &v
	return s
}

type UserGetLikeResponse struct {
	ErrMsg *string                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserGetLikeResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s UserGetLikeResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeResponse) GoString() string {
	return s.String()
}

func (s *UserGetLikeResponse) SetErrMsg(v string) *UserGetLikeResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserGetLikeResponse) SetLogId(v string) *UserGetLikeResponse {
	s.LogId = &v
	return s
}

func (s *UserGetLikeResponse) SetData(v *UserGetLikeResponseData) *UserGetLikeResponse {
	s.Data = v
	return s
}

func (s *UserGetLikeResponse) SetErrNo(v int32) *UserGetLikeResponse {
	s.ErrNo = &v
	return s
}

type UserGetLikeResponseData struct {
	Data  *UserGetLikeResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserGetLikeResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserGetLikeResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeResponseData) GoString() string {
	return s.String()
}

func (s *UserGetLikeResponseData) SetData(v *UserGetLikeResponseDataData) *UserGetLikeResponseData {
	s.Data = v
	return s
}

func (s *UserGetLikeResponseData) SetExtra(v *UserGetLikeResponseDataExtra) *UserGetLikeResponseData {
	s.Extra = v
	return s
}

type UserGetLikeResponseDataData struct {
	ResultList []*UserGetLikeResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetLikeResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetLikeResponseDataData) SetResultList(v []*UserGetLikeResponseDataDataResultListItem) *UserGetLikeResponseDataData {
	s.ResultList = v
	return s
}

type UserGetLikeResponseDataDataResultListItem struct {
	NewLike *int64  `json:"new_like,omitempty" xml:"new_like,omitempty"`
	Date    *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserGetLikeResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetLikeResponseDataDataResultListItem) SetNewLike(v int64) *UserGetLikeResponseDataDataResultListItem {
	s.NewLike = &v
	return s
}

func (s *UserGetLikeResponseDataDataResultListItem) SetDate(v string) *UserGetLikeResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserGetLikeResponseDataExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s UserGetLikeResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetLikeResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetLikeResponseDataExtra) SetLogid(v string) *UserGetLikeResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetLikeResponseDataExtra) SetNow(v int64) *UserGetLikeResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserGetLikeResponseDataExtra) SetErrorCode(v int32) *UserGetLikeResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserGetLikeResponseDataExtra) SetDescription(v string) *UserGetLikeResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserGetLikeResponseDataExtra) SetSubErrorCode(v int32) *UserGetLikeResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetLikeResponseDataExtra) SetSubDescription(v string) *UserGetLikeResponseDataExtra {
	s.SubDescription = &v
	return s
}

type UserGetProfileRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s UserGetProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileRequest) GoString() string {
	return s.String()
}

func (s *UserGetProfileRequest) SetAccessToken(v string) *UserGetProfileRequest {
	s.AccessToken = &v
	return s
}

func (s *UserGetProfileRequest) SetOpenId(v string) *UserGetProfileRequest {
	s.OpenId = &v
	return s
}

func (s *UserGetProfileRequest) SetDateType(v int64) *UserGetProfileRequest {
	s.DateType = &v
	return s
}

func (s *UserGetProfileRequest) SetHeader(v map[string]*string) *UserGetProfileRequest {
	s.Header = v
	return s
}

type UserGetProfileResponse struct {
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UserGetProfileResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UserGetProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileResponse) GoString() string {
	return s.String()
}

func (s *UserGetProfileResponse) SetErrNo(v int32) *UserGetProfileResponse {
	s.ErrNo = &v
	return s
}

func (s *UserGetProfileResponse) SetErrMsg(v string) *UserGetProfileResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserGetProfileResponse) SetLogId(v string) *UserGetProfileResponse {
	s.LogId = &v
	return s
}

func (s *UserGetProfileResponse) SetData(v *UserGetProfileResponseData) *UserGetProfileResponse {
	s.Data = v
	return s
}

type UserGetProfileResponseData struct {
	Data  *UserGetProfileResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserGetProfileResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserGetProfileResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileResponseData) GoString() string {
	return s.String()
}

func (s *UserGetProfileResponseData) SetData(v *UserGetProfileResponseDataData) *UserGetProfileResponseData {
	s.Data = v
	return s
}

func (s *UserGetProfileResponseData) SetExtra(v *UserGetProfileResponseDataExtra) *UserGetProfileResponseData {
	s.Extra = v
	return s
}

type UserGetProfileResponseDataData struct {
	ResultList []*UserGetProfileResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetProfileResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetProfileResponseDataData) SetResultList(v []*UserGetProfileResponseDataDataResultListItem) *UserGetProfileResponseDataData {
	s.ResultList = v
	return s
}

type UserGetProfileResponseDataDataResultListItem struct {
	ProfileUv *int64  `json:"profile_uv,omitempty" xml:"profile_uv,omitempty"`
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserGetProfileResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetProfileResponseDataDataResultListItem) SetProfileUv(v int64) *UserGetProfileResponseDataDataResultListItem {
	s.ProfileUv = &v
	return s
}

func (s *UserGetProfileResponseDataDataResultListItem) SetDate(v string) *UserGetProfileResponseDataDataResultListItem {
	s.Date = &v
	return s
}

type UserGetProfileResponseDataExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserGetProfileResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetProfileResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetProfileResponseDataExtra) SetDescription(v string) *UserGetProfileResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserGetProfileResponseDataExtra) SetSubErrorCode(v int32) *UserGetProfileResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetProfileResponseDataExtra) SetSubDescription(v string) *UserGetProfileResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserGetProfileResponseDataExtra) SetLogid(v string) *UserGetProfileResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetProfileResponseDataExtra) SetNow(v int64) *UserGetProfileResponseDataExtra {
	s.Now = &v
	return s
}

func (s *UserGetProfileResponseDataExtra) SetErrorCode(v int32) *UserGetProfileResponseDataExtra {
	s.ErrorCode = &v
	return s
}

type UserGetShareRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserGetShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareRequest) GoString() string {
	return s.String()
}

func (s *UserGetShareRequest) SetDateType(v int64) *UserGetShareRequest {
	s.DateType = &v
	return s
}

func (s *UserGetShareRequest) SetOpenId(v string) *UserGetShareRequest {
	s.OpenId = &v
	return s
}

func (s *UserGetShareRequest) SetHeader(v map[string]*string) *UserGetShareRequest {
	s.Header = v
	return s
}

func (s *UserGetShareRequest) SetAccessToken(v string) *UserGetShareRequest {
	s.AccessToken = &v
	return s
}

type UserGetShareResponse struct {
	Data   *UserGetShareResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UserGetShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareResponse) GoString() string {
	return s.String()
}

func (s *UserGetShareResponse) SetData(v *UserGetShareResponseData) *UserGetShareResponse {
	s.Data = v
	return s
}

func (s *UserGetShareResponse) SetErrNo(v int32) *UserGetShareResponse {
	s.ErrNo = &v
	return s
}

func (s *UserGetShareResponse) SetErrMsg(v string) *UserGetShareResponse {
	s.ErrMsg = &v
	return s
}

func (s *UserGetShareResponse) SetLogId(v string) *UserGetShareResponse {
	s.LogId = &v
	return s
}

type UserGetShareResponseData struct {
	Data  *UserGetShareResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserGetShareResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserGetShareResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareResponseData) GoString() string {
	return s.String()
}

func (s *UserGetShareResponseData) SetData(v *UserGetShareResponseDataData) *UserGetShareResponseData {
	s.Data = v
	return s
}

func (s *UserGetShareResponseData) SetExtra(v *UserGetShareResponseDataExtra) *UserGetShareResponseData {
	s.Extra = v
	return s
}

type UserGetShareResponseDataData struct {
	ResultList []*UserGetShareResponseDataDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserGetShareResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareResponseDataData) GoString() string {
	return s.String()
}

func (s *UserGetShareResponseDataData) SetResultList(v []*UserGetShareResponseDataDataResultListItem) *UserGetShareResponseDataData {
	s.ResultList = v
	return s
}

type UserGetShareResponseDataDataResultListItem struct {
	Date     *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	NewShare *int64  `json:"new_share,omitempty" xml:"new_share,omitempty"`
}

func (s UserGetShareResponseDataDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareResponseDataDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserGetShareResponseDataDataResultListItem) SetDate(v string) *UserGetShareResponseDataDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserGetShareResponseDataDataResultListItem) SetNewShare(v int64) *UserGetShareResponseDataDataResultListItem {
	s.NewShare = &v
	return s
}

type UserGetShareResponseDataExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s UserGetShareResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s UserGetShareResponseDataExtra) GoString() string {
	return s.String()
}

func (s *UserGetShareResponseDataExtra) SetErrorCode(v int32) *UserGetShareResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserGetShareResponseDataExtra) SetDescription(v string) *UserGetShareResponseDataExtra {
	s.Description = &v
	return s
}

func (s *UserGetShareResponseDataExtra) SetSubErrorCode(v int32) *UserGetShareResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserGetShareResponseDataExtra) SetSubDescription(v string) *UserGetShareResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *UserGetShareResponseDataExtra) SetLogid(v string) *UserGetShareResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *UserGetShareResponseDataExtra) SetNow(v int64) *UserGetShareResponseDataExtra {
	s.Now = &v
	return s
}

type UserItemRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UserItemRequest) GoString() string {
	return s.String()
}

func (s *UserItemRequest) SetOpenId(v string) *UserItemRequest {
	s.OpenId = &v
	return s
}

func (s *UserItemRequest) SetDateType(v int64) *UserItemRequest {
	s.DateType = &v
	return s
}

func (s *UserItemRequest) SetHeader(v map[string]*string) *UserItemRequest {
	s.Header = v
	return s
}

func (s *UserItemRequest) SetAccessToken(v string) *UserItemRequest {
	s.AccessToken = &v
	return s
}

type UserItemResponse struct {
	Extra *UserItemResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserItemResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UserItemResponse) GoString() string {
	return s.String()
}

func (s *UserItemResponse) SetExtra(v *UserItemResponseExtra) *UserItemResponse {
	s.Extra = v
	return s
}

func (s *UserItemResponse) SetData(v *UserItemResponseData) *UserItemResponse {
	s.Data = v
	return s
}

type UserItemResponseData struct {
	ResultList    []*UserItemResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserItemResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserItemResponseData) GoString() string {
	return s.String()
}

func (s *UserItemResponseData) SetResultList(v []*UserItemResponseDataResultListItem) *UserItemResponseData {
	s.ResultList = v
	return s
}

func (s *UserItemResponseData) SetGwErrorCode(v int32) *UserItemResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserItemResponseData) SetGwDescription(v string) *UserItemResponseData {
	s.GwDescription = &v
	return s
}

type UserItemResponseDataResultListItem struct {
	NewPlay    *int64  `json:"new_play,omitempty" xml:"new_play,omitempty"`
	NewIssue   *int64  `json:"new_issue,omitempty" xml:"new_issue,omitempty"`
	TotalIssue *int64  `json:"total_issue,omitempty" xml:"total_issue,omitempty"`
	Date       *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserItemResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserItemResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserItemResponseDataResultListItem) SetNewPlay(v int64) *UserItemResponseDataResultListItem {
	s.NewPlay = &v
	return s
}

func (s *UserItemResponseDataResultListItem) SetNewIssue(v int64) *UserItemResponseDataResultListItem {
	s.NewIssue = &v
	return s
}

func (s *UserItemResponseDataResultListItem) SetTotalIssue(v int64) *UserItemResponseDataResultListItem {
	s.TotalIssue = &v
	return s
}

func (s *UserItemResponseDataResultListItem) SetDate(v string) *UserItemResponseDataResultListItem {
	s.Date = &v
	return s
}

type UserItemResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserItemResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserItemResponseExtra) GoString() string {
	return s.String()
}

func (s *UserItemResponseExtra) SetSubErrorCode(v int32) *UserItemResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserItemResponseExtra) SetSubDescription(v string) *UserItemResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserItemResponseExtra) SetLogid(v string) *UserItemResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserItemResponseExtra) SetNow(v int64) *UserItemResponseExtra {
	s.Now = &v
	return s
}

func (s *UserItemResponseExtra) SetErrorCode(v int32) *UserItemResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserItemResponseExtra) SetDescription(v string) *UserItemResponseExtra {
	s.Description = &v
	return s
}

type UserLikeRequest struct {
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s UserLikeRequest) String() string {
	return tea.Prettify(s)
}

func (s UserLikeRequest) GoString() string {
	return s.String()
}

func (s *UserLikeRequest) SetDateType(v int64) *UserLikeRequest {
	s.DateType = &v
	return s
}

func (s *UserLikeRequest) SetHeader(v map[string]*string) *UserLikeRequest {
	s.Header = v
	return s
}

func (s *UserLikeRequest) SetAccessToken(v string) *UserLikeRequest {
	s.AccessToken = &v
	return s
}

func (s *UserLikeRequest) SetOpenId(v string) *UserLikeRequest {
	s.OpenId = &v
	return s
}

type UserLikeResponse struct {
	Data  *UserLikeResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UserLikeResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserLikeResponse) String() string {
	return tea.Prettify(s)
}

func (s UserLikeResponse) GoString() string {
	return s.String()
}

func (s *UserLikeResponse) SetData(v *UserLikeResponseData) *UserLikeResponse {
	s.Data = v
	return s
}

func (s *UserLikeResponse) SetExtra(v *UserLikeResponseExtra) *UserLikeResponse {
	s.Extra = v
	return s
}

type UserLikeResponseData struct {
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ResultList    []*UserLikeResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
}

func (s UserLikeResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserLikeResponseData) GoString() string {
	return s.String()
}

func (s *UserLikeResponseData) SetGwErrorCode(v int32) *UserLikeResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserLikeResponseData) SetGwDescription(v string) *UserLikeResponseData {
	s.GwDescription = &v
	return s
}

func (s *UserLikeResponseData) SetResultList(v []*UserLikeResponseDataResultListItem) *UserLikeResponseData {
	s.ResultList = v
	return s
}

type UserLikeResponseDataResultListItem struct {
	NewLike *int64  `json:"new_like,omitempty" xml:"new_like,omitempty"`
	Date    *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserLikeResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserLikeResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserLikeResponseDataResultListItem) SetNewLike(v int64) *UserLikeResponseDataResultListItem {
	s.NewLike = &v
	return s
}

func (s *UserLikeResponseDataResultListItem) SetDate(v string) *UserLikeResponseDataResultListItem {
	s.Date = &v
	return s
}

type UserLikeResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s UserLikeResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserLikeResponseExtra) GoString() string {
	return s.String()
}

func (s *UserLikeResponseExtra) SetLogid(v string) *UserLikeResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserLikeResponseExtra) SetNow(v int64) *UserLikeResponseExtra {
	s.Now = &v
	return s
}

func (s *UserLikeResponseExtra) SetErrorCode(v int32) *UserLikeResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserLikeResponseExtra) SetDescription(v string) *UserLikeResponseExtra {
	s.Description = &v
	return s
}

func (s *UserLikeResponseExtra) SetSubErrorCode(v int32) *UserLikeResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserLikeResponseExtra) SetSubDescription(v string) *UserLikeResponseExtra {
	s.SubDescription = &v
	return s
}

type UserLogoutRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	UpdateTime  *int64             `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
}

func (s UserLogoutRequest) String() string {
	return tea.Prettify(s)
}

func (s UserLogoutRequest) GoString() string {
	return s.String()
}

func (s *UserLogoutRequest) SetHeader(v map[string]*string) *UserLogoutRequest {
	s.Header = v
	return s
}

func (s *UserLogoutRequest) SetAccessToken(v string) *UserLogoutRequest {
	s.AccessToken = &v
	return s
}

func (s *UserLogoutRequest) SetAccountId(v string) *UserLogoutRequest {
	s.AccountId = &v
	return s
}

func (s *UserLogoutRequest) SetOpenId(v string) *UserLogoutRequest {
	s.OpenId = &v
	return s
}

func (s *UserLogoutRequest) SetUpdateTime(v int64) *UserLogoutRequest {
	s.UpdateTime = &v
	return s
}

type UserLogoutResponse struct {
	Data  *UserLogoutResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *UserLogoutResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UserLogoutResponse) String() string {
	return tea.Prettify(s)
}

func (s UserLogoutResponse) GoString() string {
	return s.String()
}

func (s *UserLogoutResponse) SetData(v *UserLogoutResponseData) *UserLogoutResponse {
	s.Data = v
	return s
}

func (s *UserLogoutResponse) SetExtra(v *UserLogoutResponseExtra) *UserLogoutResponse {
	s.Extra = v
	return s
}

type UserLogoutResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserLogoutResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserLogoutResponseData) GoString() string {
	return s.String()
}

func (s *UserLogoutResponseData) SetGwDescription(v string) *UserLogoutResponseData {
	s.GwDescription = &v
	return s
}

func (s *UserLogoutResponseData) SetGwErrorCode(v int32) *UserLogoutResponseData {
	s.GwErrorCode = &v
	return s
}

type UserLogoutResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s UserLogoutResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserLogoutResponseExtra) GoString() string {
	return s.String()
}

func (s *UserLogoutResponseExtra) SetNow(v int64) *UserLogoutResponseExtra {
	s.Now = &v
	return s
}

func (s *UserLogoutResponseExtra) SetSubDescription(v string) *UserLogoutResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserLogoutResponseExtra) SetSubErrorCode(v int32) *UserLogoutResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserLogoutResponseExtra) SetDescription(v string) *UserLogoutResponseExtra {
	s.Description = &v
	return s
}

func (s *UserLogoutResponseExtra) SetErrorCode(v int32) *UserLogoutResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserLogoutResponseExtra) SetLogid(v string) *UserLogoutResponseExtra {
	s.Logid = &v
	return s
}

type UserProfileRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
}

func (s UserProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UserProfileRequest) GoString() string {
	return s.String()
}

func (s *UserProfileRequest) SetHeader(v map[string]*string) *UserProfileRequest {
	s.Header = v
	return s
}

func (s *UserProfileRequest) SetAccessToken(v string) *UserProfileRequest {
	s.AccessToken = &v
	return s
}

func (s *UserProfileRequest) SetOpenId(v string) *UserProfileRequest {
	s.OpenId = &v
	return s
}

func (s *UserProfileRequest) SetDateType(v int64) *UserProfileRequest {
	s.DateType = &v
	return s
}

type UserProfileResponse struct {
	Extra *UserProfileResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserProfileResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UserProfileResponse) GoString() string {
	return s.String()
}

func (s *UserProfileResponse) SetExtra(v *UserProfileResponseExtra) *UserProfileResponse {
	s.Extra = v
	return s
}

func (s *UserProfileResponse) SetData(v *UserProfileResponseData) *UserProfileResponse {
	s.Data = v
	return s
}

type UserProfileResponseData struct {
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ResultList    []*UserProfileResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserProfileResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserProfileResponseData) GoString() string {
	return s.String()
}

func (s *UserProfileResponseData) SetGwDescription(v string) *UserProfileResponseData {
	s.GwDescription = &v
	return s
}

func (s *UserProfileResponseData) SetResultList(v []*UserProfileResponseDataResultListItem) *UserProfileResponseData {
	s.ResultList = v
	return s
}

func (s *UserProfileResponseData) SetGwErrorCode(v int32) *UserProfileResponseData {
	s.GwErrorCode = &v
	return s
}

type UserProfileResponseDataResultListItem struct {
	Date      *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
	ProfileUv *int64  `json:"profile_uv,omitempty" xml:"profile_uv,omitempty"`
}

func (s UserProfileResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserProfileResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserProfileResponseDataResultListItem) SetDate(v string) *UserProfileResponseDataResultListItem {
	s.Date = &v
	return s
}

func (s *UserProfileResponseDataResultListItem) SetProfileUv(v int64) *UserProfileResponseDataResultListItem {
	s.ProfileUv = &v
	return s
}

type UserProfileResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserProfileResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserProfileResponseExtra) GoString() string {
	return s.String()
}

func (s *UserProfileResponseExtra) SetSubErrorCode(v int32) *UserProfileResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserProfileResponseExtra) SetSubDescription(v string) *UserProfileResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserProfileResponseExtra) SetLogid(v string) *UserProfileResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserProfileResponseExtra) SetNow(v int64) *UserProfileResponseExtra {
	s.Now = &v
	return s
}

func (s *UserProfileResponseExtra) SetErrorCode(v int32) *UserProfileResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserProfileResponseExtra) SetDescription(v string) *UserProfileResponseExtra {
	s.Description = &v
	return s
}

type UserShareRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	DateType    *int64             `json:"date_type,omitempty" xml:"date_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserShareRequest) String() string {
	return tea.Prettify(s)
}

func (s UserShareRequest) GoString() string {
	return s.String()
}

func (s *UserShareRequest) SetOpenId(v string) *UserShareRequest {
	s.OpenId = &v
	return s
}

func (s *UserShareRequest) SetDateType(v int64) *UserShareRequest {
	s.DateType = &v
	return s
}

func (s *UserShareRequest) SetHeader(v map[string]*string) *UserShareRequest {
	s.Header = v
	return s
}

func (s *UserShareRequest) SetAccessToken(v string) *UserShareRequest {
	s.AccessToken = &v
	return s
}

type UserShareResponse struct {
	Extra *UserShareResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *UserShareResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UserShareResponse) String() string {
	return tea.Prettify(s)
}

func (s UserShareResponse) GoString() string {
	return s.String()
}

func (s *UserShareResponse) SetExtra(v *UserShareResponseExtra) *UserShareResponse {
	s.Extra = v
	return s
}

func (s *UserShareResponse) SetData(v *UserShareResponseData) *UserShareResponse {
	s.Data = v
	return s
}

type UserShareResponseData struct {
	ResultList    []*UserShareResponseDataResultListItem `json:"result_list,omitempty" xml:"result_list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserShareResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserShareResponseData) GoString() string {
	return s.String()
}

func (s *UserShareResponseData) SetResultList(v []*UserShareResponseDataResultListItem) *UserShareResponseData {
	s.ResultList = v
	return s
}

func (s *UserShareResponseData) SetGwErrorCode(v int32) *UserShareResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserShareResponseData) SetGwDescription(v string) *UserShareResponseData {
	s.GwDescription = &v
	return s
}

type UserShareResponseDataResultListItem struct {
	NewShare *int64  `json:"new_share,omitempty" xml:"new_share,omitempty"`
	Date     *string `json:"date,omitempty" xml:"date,omitempty" require:"true"`
}

func (s UserShareResponseDataResultListItem) String() string {
	return tea.Prettify(s)
}

func (s UserShareResponseDataResultListItem) GoString() string {
	return s.String()
}

func (s *UserShareResponseDataResultListItem) SetNewShare(v int64) *UserShareResponseDataResultListItem {
	s.NewShare = &v
	return s
}

func (s *UserShareResponseDataResultListItem) SetDate(v string) *UserShareResponseDataResultListItem {
	s.Date = &v
	return s
}

type UserShareResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s UserShareResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserShareResponseExtra) GoString() string {
	return s.String()
}

func (s *UserShareResponseExtra) SetDescription(v string) *UserShareResponseExtra {
	s.Description = &v
	return s
}

func (s *UserShareResponseExtra) SetSubErrorCode(v int32) *UserShareResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *UserShareResponseExtra) SetSubDescription(v string) *UserShareResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserShareResponseExtra) SetLogid(v string) *UserShareResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserShareResponseExtra) SetNow(v int64) *UserShareResponseExtra {
	s.Now = &v
	return s
}

func (s *UserShareResponseExtra) SetErrorCode(v int32) *UserShareResponseExtra {
	s.ErrorCode = &v
	return s
}

type UserUpdateRequest struct {
	Members     []*UserUpdateRequestMembersItem `json:"members,omitempty" xml:"members,omitempty" require:"true" type:"Repeated"`
	AccountId   *string                         `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UserUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s UserUpdateRequest) GoString() string {
	return s.String()
}

func (s *UserUpdateRequest) SetMembers(v []*UserUpdateRequestMembersItem) *UserUpdateRequest {
	s.Members = v
	return s
}

func (s *UserUpdateRequest) SetAccountId(v string) *UserUpdateRequest {
	s.AccountId = &v
	return s
}

func (s *UserUpdateRequest) SetHeader(v map[string]*string) *UserUpdateRequest {
	s.Header = v
	return s
}

func (s *UserUpdateRequest) SetAccessToken(v string) *UserUpdateRequest {
	s.AccessToken = &v
	return s
}

type UserUpdateRequestMembersItem struct {
	UserLevel        *int32  `json:"user_level,omitempty" xml:"user_level,omitempty" require:"true"`
	NewMobile        *string `json:"new_mobile,omitempty" xml:"new_mobile,omitempty"`
	OpenId           *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	PointsAmountCent *int64  `json:"points_amount_cent,omitempty" xml:"points_amount_cent,omitempty" require:"true"`
	PointsUpdateCent *int64  `json:"points_update_cent,omitempty" xml:"points_update_cent,omitempty"`
	Unbind           *bool   `json:"unbind,omitempty" xml:"unbind,omitempty"`
	UpdateTime       *int64  `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
}

func (s UserUpdateRequestMembersItem) String() string {
	return tea.Prettify(s)
}

func (s UserUpdateRequestMembersItem) GoString() string {
	return s.String()
}

func (s *UserUpdateRequestMembersItem) SetUserLevel(v int32) *UserUpdateRequestMembersItem {
	s.UserLevel = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetNewMobile(v string) *UserUpdateRequestMembersItem {
	s.NewMobile = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetOpenId(v string) *UserUpdateRequestMembersItem {
	s.OpenId = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetPointsAmountCent(v int64) *UserUpdateRequestMembersItem {
	s.PointsAmountCent = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetPointsUpdateCent(v int64) *UserUpdateRequestMembersItem {
	s.PointsUpdateCent = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetUnbind(v bool) *UserUpdateRequestMembersItem {
	s.Unbind = &v
	return s
}

func (s *UserUpdateRequestMembersItem) SetUpdateTime(v int64) *UserUpdateRequestMembersItem {
	s.UpdateTime = &v
	return s
}

type UserUpdateResponse struct {
	Data       *UserUpdateResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Now        *string                  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	StatusCode *int32                   `json:"status_code,omitempty" xml:"status_code,omitempty" require:"true"`
	StatusMsg  *string                  `json:"status_msg,omitempty" xml:"status_msg,omitempty" require:"true"`
	Extra      *UserUpdateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	LogId      *string                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UserUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s UserUpdateResponse) GoString() string {
	return s.String()
}

func (s *UserUpdateResponse) SetData(v *UserUpdateResponseData) *UserUpdateResponse {
	s.Data = v
	return s
}

func (s *UserUpdateResponse) SetNow(v string) *UserUpdateResponse {
	s.Now = &v
	return s
}

func (s *UserUpdateResponse) SetStatusCode(v int32) *UserUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *UserUpdateResponse) SetStatusMsg(v string) *UserUpdateResponse {
	s.StatusMsg = &v
	return s
}

func (s *UserUpdateResponse) SetExtra(v *UserUpdateResponseExtra) *UserUpdateResponse {
	s.Extra = v
	return s
}

func (s *UserUpdateResponse) SetLogId(v string) *UserUpdateResponse {
	s.LogId = &v
	return s
}

type UserUpdateResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UserUpdateResponseData) String() string {
	return tea.Prettify(s)
}

func (s UserUpdateResponseData) GoString() string {
	return s.String()
}

func (s *UserUpdateResponseData) SetGwErrorCode(v int32) *UserUpdateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UserUpdateResponseData) SetGwDescription(v string) *UserUpdateResponseData {
	s.GwDescription = &v
	return s
}

type UserUpdateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s UserUpdateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UserUpdateResponseExtra) GoString() string {
	return s.String()
}

func (s *UserUpdateResponseExtra) SetDescription(v string) *UserUpdateResponseExtra {
	s.Description = &v
	return s
}

func (s *UserUpdateResponseExtra) SetErrorCode(v int32) *UserUpdateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UserUpdateResponseExtra) SetLogid(v string) *UserUpdateResponseExtra {
	s.Logid = &v
	return s
}

func (s *UserUpdateResponseExtra) SetNow(v int64) *UserUpdateResponseExtra {
	s.Now = &v
	return s
}

func (s *UserUpdateResponseExtra) SetSubDescription(v string) *UserUpdateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UserUpdateResponseExtra) SetSubErrorCode(v int32) *UserUpdateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type V1AppRegionAddRequest struct {
	RegionId    *string            `json:"region_id,omitempty" xml:"region_id,omitempty" require:"true"`
	Path        *string            `json:"path,omitempty" xml:"path,omitempty"`
	BindType    *int               `json:"bind_type,omitempty" xml:"bind_type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V1AppRegionAddRequest) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionAddRequest) GoString() string {
	return s.String()
}

func (s *V1AppRegionAddRequest) SetRegionId(v string) *V1AppRegionAddRequest {
	s.RegionId = &v
	return s
}

func (s *V1AppRegionAddRequest) SetPath(v string) *V1AppRegionAddRequest {
	s.Path = &v
	return s
}

func (s *V1AppRegionAddRequest) SetBindType(v int) *V1AppRegionAddRequest {
	s.BindType = &v
	return s
}

func (s *V1AppRegionAddRequest) SetHeader(v map[string]*string) *V1AppRegionAddRequest {
	s.Header = v
	return s
}

func (s *V1AppRegionAddRequest) SetAccessToken(v string) *V1AppRegionAddRequest {
	s.AccessToken = &v
	return s
}

type V1AppRegionAddResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	DataId *string `json:"data_id,omitempty" xml:"data_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s V1AppRegionAddResponse) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionAddResponse) GoString() string {
	return s.String()
}

func (s *V1AppRegionAddResponse) SetLogId(v string) *V1AppRegionAddResponse {
	s.LogId = &v
	return s
}

func (s *V1AppRegionAddResponse) SetDataId(v string) *V1AppRegionAddResponse {
	s.DataId = &v
	return s
}

func (s *V1AppRegionAddResponse) SetErrNo(v int32) *V1AppRegionAddResponse {
	s.ErrNo = &v
	return s
}

func (s *V1AppRegionAddResponse) SetErrMsg(v string) *V1AppRegionAddResponse {
	s.ErrMsg = &v
	return s
}

type V1AppRegionDeleteRequest struct {
	RegionId    *string            `json:"region_id,omitempty" xml:"region_id,omitempty" require:"true"`
	Path        *string            `json:"path,omitempty" xml:"path,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V1AppRegionDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionDeleteRequest) GoString() string {
	return s.String()
}

func (s *V1AppRegionDeleteRequest) SetRegionId(v string) *V1AppRegionDeleteRequest {
	s.RegionId = &v
	return s
}

func (s *V1AppRegionDeleteRequest) SetPath(v string) *V1AppRegionDeleteRequest {
	s.Path = &v
	return s
}

func (s *V1AppRegionDeleteRequest) SetHeader(v map[string]*string) *V1AppRegionDeleteRequest {
	s.Header = v
	return s
}

func (s *V1AppRegionDeleteRequest) SetAccessToken(v string) *V1AppRegionDeleteRequest {
	s.AccessToken = &v
	return s
}

type V1AppRegionDeleteResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s V1AppRegionDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionDeleteResponse) GoString() string {
	return s.String()
}

func (s *V1AppRegionDeleteResponse) SetErrMsg(v string) *V1AppRegionDeleteResponse {
	s.ErrMsg = &v
	return s
}

func (s *V1AppRegionDeleteResponse) SetLogId(v string) *V1AppRegionDeleteResponse {
	s.LogId = &v
	return s
}

func (s *V1AppRegionDeleteResponse) SetErrNo(v int32) *V1AppRegionDeleteResponse {
	s.ErrNo = &v
	return s
}

type V1AppRegionModifyRequest struct {
	Path          *string            `json:"path,omitempty" xml:"path,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OriginRegion  *string            `json:"origin_region,omitempty" xml:"origin_region,omitempty" require:"true"`
	CurrentRegion *string            `json:"current_region,omitempty" xml:"current_region,omitempty" require:"true"`
}

func (s V1AppRegionModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionModifyRequest) GoString() string {
	return s.String()
}

func (s *V1AppRegionModifyRequest) SetPath(v string) *V1AppRegionModifyRequest {
	s.Path = &v
	return s
}

func (s *V1AppRegionModifyRequest) SetHeader(v map[string]*string) *V1AppRegionModifyRequest {
	s.Header = v
	return s
}

func (s *V1AppRegionModifyRequest) SetAccessToken(v string) *V1AppRegionModifyRequest {
	s.AccessToken = &v
	return s
}

func (s *V1AppRegionModifyRequest) SetOriginRegion(v string) *V1AppRegionModifyRequest {
	s.OriginRegion = &v
	return s
}

func (s *V1AppRegionModifyRequest) SetCurrentRegion(v string) *V1AppRegionModifyRequest {
	s.CurrentRegion = &v
	return s
}

type V1AppRegionModifyResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s V1AppRegionModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s V1AppRegionModifyResponse) GoString() string {
	return s.String()
}

func (s *V1AppRegionModifyResponse) SetErrNo(v int32) *V1AppRegionModifyResponse {
	s.ErrNo = &v
	return s
}

func (s *V1AppRegionModifyResponse) SetErrMsg(v string) *V1AppRegionModifyResponse {
	s.ErrMsg = &v
	return s
}

func (s *V1AppRegionModifyResponse) SetLogId(v string) *V1AppRegionModifyResponse {
	s.LogId = &v
	return s
}

type V1AuthGetRelatedIdRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V1AuthGetRelatedIdRequest) String() string {
	return tea.Prettify(s)
}

func (s V1AuthGetRelatedIdRequest) GoString() string {
	return s.String()
}

func (s *V1AuthGetRelatedIdRequest) SetOpenId(v string) *V1AuthGetRelatedIdRequest {
	s.OpenId = &v
	return s
}

func (s *V1AuthGetRelatedIdRequest) SetHeader(v map[string]*string) *V1AuthGetRelatedIdRequest {
	s.Header = v
	return s
}

func (s *V1AuthGetRelatedIdRequest) SetAccessToken(v string) *V1AuthGetRelatedIdRequest {
	s.AccessToken = &v
	return s
}

type V1AuthGetRelatedIdResponse struct {
	Data   *V1AuthGetRelatedIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty"`
}

func (s V1AuthGetRelatedIdResponse) String() string {
	return tea.Prettify(s)
}

func (s V1AuthGetRelatedIdResponse) GoString() string {
	return s.String()
}

func (s *V1AuthGetRelatedIdResponse) SetData(v *V1AuthGetRelatedIdResponseData) *V1AuthGetRelatedIdResponse {
	s.Data = v
	return s
}

func (s *V1AuthGetRelatedIdResponse) SetErrMsg(v string) *V1AuthGetRelatedIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *V1AuthGetRelatedIdResponse) SetErrNo(v int32) *V1AuthGetRelatedIdResponse {
	s.ErrNo = &v
	return s
}

func (s *V1AuthGetRelatedIdResponse) SetLogId(v string) *V1AuthGetRelatedIdResponse {
	s.LogId = &v
	return s
}

type V1AuthGetRelatedIdResponseData struct {
	AlliedId *string `json:"allied_id,omitempty" xml:"allied_id,omitempty"`
}

func (s V1AuthGetRelatedIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s V1AuthGetRelatedIdResponseData) GoString() string {
	return s.String()
}

func (s *V1AuthGetRelatedIdResponseData) SetAlliedId(v string) *V1AuthGetRelatedIdResponseData {
	s.AlliedId = &v
	return s
}

type V1GetPhonenumberInfoRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Code        *string            `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s V1GetPhonenumberInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s V1GetPhonenumberInfoRequest) GoString() string {
	return s.String()
}

func (s *V1GetPhonenumberInfoRequest) SetHeader(v map[string]*string) *V1GetPhonenumberInfoRequest {
	s.Header = v
	return s
}

func (s *V1GetPhonenumberInfoRequest) SetAccessToken(v string) *V1GetPhonenumberInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *V1GetPhonenumberInfoRequest) SetCode(v string) *V1GetPhonenumberInfoRequest {
	s.Code = &v
	return s
}

type V1GetPhonenumberInfoResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s V1GetPhonenumberInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s V1GetPhonenumberInfoResponse) GoString() string {
	return s.String()
}

func (s *V1GetPhonenumberInfoResponse) SetLogId(v string) *V1GetPhonenumberInfoResponse {
	s.LogId = &v
	return s
}

func (s *V1GetPhonenumberInfoResponse) SetData(v string) *V1GetPhonenumberInfoResponse {
	s.Data = &v
	return s
}

func (s *V1GetPhonenumberInfoResponse) SetErrNo(v int32) *V1GetPhonenumberInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *V1GetPhonenumberInfoResponse) SetErrMsg(v string) *V1GetPhonenumberInfoResponse {
	s.ErrMsg = &v
	return s
}

type V1GoodsProductSaveRequest struct {
	Skus        []*V1GoodsProductSaveRequestSkusItem `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	Ability     *V1GoodsProductSaveRequestAbility    `json:"ability,omitempty" xml:"ability,omitempty"`
	Header      map[string]*string                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string                              `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Product     *V1GoodsProductSaveRequestProduct    `json:"product,omitempty" xml:"product,omitempty" require:"true"`
	Sku         *V1GoodsProductSaveRequestSku        `json:"sku,omitempty" xml:"sku,omitempty"`
}

func (s V1GoodsProductSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequest) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequest) SetSkus(v []*V1GoodsProductSaveRequestSkusItem) *V1GoodsProductSaveRequest {
	s.Skus = v
	return s
}

func (s *V1GoodsProductSaveRequest) SetAbility(v *V1GoodsProductSaveRequestAbility) *V1GoodsProductSaveRequest {
	s.Ability = v
	return s
}

func (s *V1GoodsProductSaveRequest) SetHeader(v map[string]*string) *V1GoodsProductSaveRequest {
	s.Header = v
	return s
}

func (s *V1GoodsProductSaveRequest) SetAccessToken(v string) *V1GoodsProductSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *V1GoodsProductSaveRequest) SetAccountId(v string) *V1GoodsProductSaveRequest {
	s.AccountId = &v
	return s
}

func (s *V1GoodsProductSaveRequest) SetProduct(v *V1GoodsProductSaveRequestProduct) *V1GoodsProductSaveRequest {
	s.Product = v
	return s
}

func (s *V1GoodsProductSaveRequest) SetSku(v *V1GoodsProductSaveRequestSku) *V1GoodsProductSaveRequest {
	s.Sku = v
	return s
}

type V1GoodsProductSaveRequestAbility struct {
	IgnoreInapplicablePoi *bool `json:"ignore_inapplicable_poi,omitempty" xml:"ignore_inapplicable_poi,omitempty"`
}

func (s V1GoodsProductSaveRequestAbility) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestAbility) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestAbility) SetIgnoreInapplicablePoi(v bool) *V1GoodsProductSaveRequestAbility {
	s.IgnoreInapplicablePoi = &v
	return s
}

type V1GoodsProductSaveRequestProduct struct {
	SoldStartTime    *int64                                      `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	CategoryId       *int64                                      `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	UpdateTime       *int64                                      `json:"update_time,omitempty" xml:"update_time,omitempty"`
	OwnerAccountId   *int64                                      `json:"owner_account_id,omitempty" xml:"owner_account_id,omitempty"`
	BizLine          *int                                        `json:"biz_line,omitempty" xml:"biz_line,omitempty" require:"true"`
	ProductId        *string                                     `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SoldEndTime      *int64                                      `json:"sold_end_time,omitempty" xml:"sold_end_time,omitempty"`
	ContactName      *string                                     `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	OutId            *string                                     `json:"out_id,omitempty" xml:"out_id,omitempty"`
	CreatorAccountId *int64                                      `json:"creator_account_id,omitempty" xml:"creator_account_id,omitempty"`
	CategoryFullName *string                                     `json:"category_full_name,omitempty" xml:"category_full_name,omitempty"`
	Extra            *string                                     `json:"extra,omitempty" xml:"extra,omitempty"`
	ProductExt       *V1GoodsProductSaveRequestProductProductExt `json:"product_ext,omitempty" xml:"product_ext,omitempty"`
	SpuId            *string                                     `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Pois             []*V1GoodsProductSaveRequestProductPoisItem `json:"pois,omitempty" xml:"pois,omitempty" type:"Repeated"`
	AttrKeyValueMap  map[string]*string                          `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	OpenBizType      *int                                        `json:"open_biz_type,omitempty" xml:"open_biz_type,omitempty"`
	Desc             *string                                     `json:"desc,omitempty" xml:"desc,omitempty"`
	ProductName      *string                                     `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	AccountName      *string                                     `json:"account_name,omitempty" xml:"account_name,omitempty"`
	ProductType      *int                                        `json:"product_type,omitempty" xml:"product_type,omitempty" require:"true"`
	CreateTime       *int64                                      `json:"create_time,omitempty" xml:"create_time,omitempty"`
	ProductSubType   *int                                        `json:"product_sub_type,omitempty" xml:"product_sub_type,omitempty"`
	Telephone        []*string                                   `json:"telephone,omitempty" xml:"telephone,omitempty" type:"Repeated"`
	OutUrl           *string                                     `json:"out_url,omitempty" xml:"out_url,omitempty"`
	Version          *int64                                      `json:"version,omitempty" xml:"version,omitempty"`
}

func (s V1GoodsProductSaveRequestProduct) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestProduct) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestProduct) SetSoldStartTime(v int64) *V1GoodsProductSaveRequestProduct {
	s.SoldStartTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetCategoryId(v int64) *V1GoodsProductSaveRequestProduct {
	s.CategoryId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetUpdateTime(v int64) *V1GoodsProductSaveRequestProduct {
	s.UpdateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetOwnerAccountId(v int64) *V1GoodsProductSaveRequestProduct {
	s.OwnerAccountId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetBizLine(v int) *V1GoodsProductSaveRequestProduct {
	s.BizLine = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetProductId(v string) *V1GoodsProductSaveRequestProduct {
	s.ProductId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetSoldEndTime(v int64) *V1GoodsProductSaveRequestProduct {
	s.SoldEndTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetContactName(v string) *V1GoodsProductSaveRequestProduct {
	s.ContactName = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetOutId(v string) *V1GoodsProductSaveRequestProduct {
	s.OutId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetCreatorAccountId(v int64) *V1GoodsProductSaveRequestProduct {
	s.CreatorAccountId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetCategoryFullName(v string) *V1GoodsProductSaveRequestProduct {
	s.CategoryFullName = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetExtra(v string) *V1GoodsProductSaveRequestProduct {
	s.Extra = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetProductExt(v *V1GoodsProductSaveRequestProductProductExt) *V1GoodsProductSaveRequestProduct {
	s.ProductExt = v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetSpuId(v string) *V1GoodsProductSaveRequestProduct {
	s.SpuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetPois(v []*V1GoodsProductSaveRequestProductPoisItem) *V1GoodsProductSaveRequestProduct {
	s.Pois = v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetAttrKeyValueMap(v map[string]*string) *V1GoodsProductSaveRequestProduct {
	s.AttrKeyValueMap = v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetOpenBizType(v int) *V1GoodsProductSaveRequestProduct {
	s.OpenBizType = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetDesc(v string) *V1GoodsProductSaveRequestProduct {
	s.Desc = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetProductName(v string) *V1GoodsProductSaveRequestProduct {
	s.ProductName = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetAccountName(v string) *V1GoodsProductSaveRequestProduct {
	s.AccountName = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetProductType(v int) *V1GoodsProductSaveRequestProduct {
	s.ProductType = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetCreateTime(v int64) *V1GoodsProductSaveRequestProduct {
	s.CreateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetProductSubType(v int) *V1GoodsProductSaveRequestProduct {
	s.ProductSubType = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetTelephone(v []*string) *V1GoodsProductSaveRequestProduct {
	s.Telephone = v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetOutUrl(v string) *V1GoodsProductSaveRequestProduct {
	s.OutUrl = &v
	return s
}

func (s *V1GoodsProductSaveRequestProduct) SetVersion(v int64) *V1GoodsProductSaveRequestProduct {
	s.Version = &v
	return s
}

type V1GoodsProductSaveRequestProductPoisItem struct {
	PoiId         *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SupplierExtId *string `json:"supplier_ext_id,omitempty" xml:"supplier_ext_id,omitempty"`
	SupplierId    *int64  `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

func (s V1GoodsProductSaveRequestProductPoisItem) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestProductPoisItem) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestProductPoisItem) SetPoiId(v string) *V1GoodsProductSaveRequestProductPoisItem {
	s.PoiId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductPoisItem) SetSupplierExtId(v string) *V1GoodsProductSaveRequestProductPoisItem {
	s.SupplierExtId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductPoisItem) SetSupplierId(v int64) *V1GoodsProductSaveRequestProductPoisItem {
	s.SupplierId = &v
	return s
}

type V1GoodsProductSaveRequestProductProductExt struct {
	AutoOnline        *bool                                                `json:"auto_online,omitempty" xml:"auto_online,omitempty"`
	ElemeExtraInfo    *string                                              `json:"eleme_extra_info,omitempty" xml:"eleme_extra_info,omitempty"`
	ElemeBizCode      *string                                              `json:"eleme_biz_code,omitempty" xml:"eleme_biz_code,omitempty"`
	IsBindClueElement *bool                                                `json:"is_bind_clue_element,omitempty" xml:"is_bind_clue_element,omitempty"`
	AgencyRate        *int64                                               `json:"agency_rate,omitempty" xml:"agency_rate,omitempty"`
	AllSkuSellOut     *bool                                                `json:"all_sku_sell_out,omitempty" xml:"all_sku_sell_out,omitempty"`
	CategoryOutId     *string                                              `json:"category_out_id,omitempty" xml:"category_out_id,omitempty"`
	TestExtra         *V1GoodsProductSaveRequestProductProductExtTestExtra `json:"test_extra,omitempty" xml:"test_extra,omitempty"`
}

func (s V1GoodsProductSaveRequestProductProductExt) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestProductProductExt) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetAutoOnline(v bool) *V1GoodsProductSaveRequestProductProductExt {
	s.AutoOnline = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetElemeExtraInfo(v string) *V1GoodsProductSaveRequestProductProductExt {
	s.ElemeExtraInfo = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetElemeBizCode(v string) *V1GoodsProductSaveRequestProductProductExt {
	s.ElemeBizCode = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetIsBindClueElement(v bool) *V1GoodsProductSaveRequestProductProductExt {
	s.IsBindClueElement = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetAgencyRate(v int64) *V1GoodsProductSaveRequestProductProductExt {
	s.AgencyRate = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetAllSkuSellOut(v bool) *V1GoodsProductSaveRequestProductProductExt {
	s.AllSkuSellOut = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetCategoryOutId(v string) *V1GoodsProductSaveRequestProductProductExt {
	s.CategoryOutId = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExt) SetTestExtra(v *V1GoodsProductSaveRequestProductProductExtTestExtra) *V1GoodsProductSaveRequestProductProductExt {
	s.TestExtra = v
	return s
}

type V1GoodsProductSaveRequestProductProductExtTestExtra struct {
	TestFlag *bool     `json:"test_flag,omitempty" xml:"test_flag,omitempty"`
	Uids     []*string `json:"uids,omitempty" xml:"uids,omitempty" type:"Repeated"`
}

func (s V1GoodsProductSaveRequestProductProductExtTestExtra) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestProductProductExtTestExtra) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestProductProductExtTestExtra) SetTestFlag(v bool) *V1GoodsProductSaveRequestProductProductExtTestExtra {
	s.TestFlag = &v
	return s
}

func (s *V1GoodsProductSaveRequestProductProductExtTestExtra) SetUids(v []*string) *V1GoodsProductSaveRequestProductProductExtTestExtra {
	s.Uids = v
	return s
}

type V1GoodsProductSaveRequestSku struct {
	SkuName         *string                             `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	OutSkuId        *string                             `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	AttrKeyValueMap map[string]*string                  `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	OriginAmount    *int64                              `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ActualAmount    *int64                              `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	CreateTime      *int64                              `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Extra           *string                             `json:"extra,omitempty" xml:"extra,omitempty"`
	SkuId           *string                             `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	Stock           *V1GoodsProductSaveRequestSkuStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	UpdateTime      *int64                              `json:"update_time,omitempty" xml:"update_time,omitempty"`
	Status          *int                                `json:"status,omitempty" xml:"status,omitempty"`
	BindSkus        []*string                           `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	SkuExt          *V1GoodsProductSaveRequestSkuSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
}

func (s V1GoodsProductSaveRequestSku) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSku) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSku) SetSkuName(v string) *V1GoodsProductSaveRequestSku {
	s.SkuName = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetOutSkuId(v string) *V1GoodsProductSaveRequestSku {
	s.OutSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetAttrKeyValueMap(v map[string]*string) *V1GoodsProductSaveRequestSku {
	s.AttrKeyValueMap = v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetOriginAmount(v int64) *V1GoodsProductSaveRequestSku {
	s.OriginAmount = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetActualAmount(v int64) *V1GoodsProductSaveRequestSku {
	s.ActualAmount = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetCreateTime(v int64) *V1GoodsProductSaveRequestSku {
	s.CreateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetExtra(v string) *V1GoodsProductSaveRequestSku {
	s.Extra = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetSkuId(v string) *V1GoodsProductSaveRequestSku {
	s.SkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetStock(v *V1GoodsProductSaveRequestSkuStock) *V1GoodsProductSaveRequestSku {
	s.Stock = v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetUpdateTime(v int64) *V1GoodsProductSaveRequestSku {
	s.UpdateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetStatus(v int) *V1GoodsProductSaveRequestSku {
	s.Status = &v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetBindSkus(v []*string) *V1GoodsProductSaveRequestSku {
	s.BindSkus = v
	return s
}

func (s *V1GoodsProductSaveRequestSku) SetSkuExt(v *V1GoodsProductSaveRequestSkuSkuExt) *V1GoodsProductSaveRequestSku {
	s.SkuExt = v
	return s
}

type V1GoodsProductSaveRequestSkuSkuExt struct {
	SettleType          *int64                                                 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	OriSkus             map[int64][]*int64                                     `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	BindSkus2c          map[int64][]*int64                                     `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	RelRuleList         []*V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	BizId2cList         []*int64                                               `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	AccountSettle       *bool                                                  `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
	OriginStockQty      *int64                                                 `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	TakeawayPresaleInfo *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	CMspuId             *int64                                                 `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	BindProductId       *int64                                                 `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
	LifeBizCode         *string                                                `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	BindSkuId           *int64                                                 `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	DiscountPromo       *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	UseSubRelStock      *bool                                                  `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
}

func (s V1GoodsProductSaveRequestSkuSkuExt) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkuSkuExt) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetSettleType(v int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.SettleType = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetOriSkus(v map[int64][]*int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.OriSkus = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetBindSkus2c(v map[int64][]*int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetRelRuleList(v []*V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) *V1GoodsProductSaveRequestSkuSkuExt {
	s.RelRuleList = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetBizId2cList(v []*int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.BizId2cList = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetAccountSettle(v bool) *V1GoodsProductSaveRequestSkuSkuExt {
	s.AccountSettle = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetOriginStockQty(v int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetTakeawayPresaleInfo(v *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo) *V1GoodsProductSaveRequestSkuSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetCMspuId(v int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.CMspuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetBindProductId(v int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.BindProductId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetLifeBizCode(v string) *V1GoodsProductSaveRequestSkuSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetBindSkuId(v int64) *V1GoodsProductSaveRequestSkuSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetDiscountPromo(v *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) *V1GoodsProductSaveRequestSkuSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExt) SetUseSubRelStock(v bool) *V1GoodsProductSaveRequestSkuSkuExt {
	s.UseSubRelStock = &v
	return s
}

type V1GoodsProductSaveRequestSkuSkuExtDiscountPromo struct {
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
}

func (s V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) SetBrandActivityId(v int64) *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) SetPlanId(v int64) *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo) SetPromoId(v int64) *V1GoodsProductSaveRequestSkuSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

type V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem struct {
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
}

func (s V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetConstantVal(v int64) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetPriceRel(v bool) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetSharedQty(v int64) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetStockRel(v bool) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetBizId(v int64) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem) SetCoefficient(v string) *V1GoodsProductSaveRequestSkuSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

type V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
}

func (s V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *V1GoodsProductSaveRequestSkuSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

type V1GoodsProductSaveRequestSkuStock struct {
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
}

func (s V1GoodsProductSaveRequestSkuStock) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkuStock) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkuStock) SetSoldQty(v int64) *V1GoodsProductSaveRequestSkuStock {
	s.SoldQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuStock) SetStockQty(v int64) *V1GoodsProductSaveRequestSkuStock {
	s.StockQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuStock) SetAvailQty(v int64) *V1GoodsProductSaveRequestSkuStock {
	s.AvailQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuStock) SetFrozenQty(v int64) *V1GoodsProductSaveRequestSkuStock {
	s.FrozenQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuStock) SetLimitType(v int) *V1GoodsProductSaveRequestSkuStock {
	s.LimitType = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkuStock) SetSoldCount(v int64) *V1GoodsProductSaveRequestSkuStock {
	s.SoldCount = &v
	return s
}

type V1GoodsProductSaveRequestSkusItem struct {
	ActualAmount    *int64                                   `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	UpdateTime      *int64                                   `json:"update_time,omitempty" xml:"update_time,omitempty"`
	OutSkuId        *string                                  `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	AttrKeyValueMap map[string]*string                       `json:"attr_key_value_map,omitempty" xml:"attr_key_value_map,omitempty"`
	BindSkus        []*string                                `json:"bind_skus,omitempty" xml:"bind_skus,omitempty" type:"Repeated"`
	Extra           *string                                  `json:"extra,omitempty" xml:"extra,omitempty"`
	Stock           *V1GoodsProductSaveRequestSkusItemStock  `json:"stock,omitempty" xml:"stock,omitempty"`
	SkuExt          *V1GoodsProductSaveRequestSkusItemSkuExt `json:"sku_ext,omitempty" xml:"sku_ext,omitempty"`
	SkuName         *string                                  `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	Status          *int                                     `json:"status,omitempty" xml:"status,omitempty"`
	CreateTime      *int64                                   `json:"create_time,omitempty" xml:"create_time,omitempty"`
	OriginAmount    *int64                                   `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	SkuId           *string                                  `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s V1GoodsProductSaveRequestSkusItem) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItem) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItem) SetActualAmount(v int64) *V1GoodsProductSaveRequestSkusItem {
	s.ActualAmount = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetUpdateTime(v int64) *V1GoodsProductSaveRequestSkusItem {
	s.UpdateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetOutSkuId(v string) *V1GoodsProductSaveRequestSkusItem {
	s.OutSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetAttrKeyValueMap(v map[string]*string) *V1GoodsProductSaveRequestSkusItem {
	s.AttrKeyValueMap = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetBindSkus(v []*string) *V1GoodsProductSaveRequestSkusItem {
	s.BindSkus = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetExtra(v string) *V1GoodsProductSaveRequestSkusItem {
	s.Extra = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetStock(v *V1GoodsProductSaveRequestSkusItemStock) *V1GoodsProductSaveRequestSkusItem {
	s.Stock = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetSkuExt(v *V1GoodsProductSaveRequestSkusItemSkuExt) *V1GoodsProductSaveRequestSkusItem {
	s.SkuExt = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetSkuName(v string) *V1GoodsProductSaveRequestSkusItem {
	s.SkuName = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetStatus(v int) *V1GoodsProductSaveRequestSkusItem {
	s.Status = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetCreateTime(v int64) *V1GoodsProductSaveRequestSkusItem {
	s.CreateTime = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetOriginAmount(v int64) *V1GoodsProductSaveRequestSkusItem {
	s.OriginAmount = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItem) SetSkuId(v string) *V1GoodsProductSaveRequestSkusItem {
	s.SkuId = &v
	return s
}

type V1GoodsProductSaveRequestSkusItemSkuExt struct {
	UseSubRelStock      *bool                                                       `json:"use_sub_rel_stock,omitempty" xml:"use_sub_rel_stock,omitempty"`
	RelRuleList         []*V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem   `json:"rel_rule_list,omitempty" xml:"rel_rule_list,omitempty" type:"Repeated"`
	SettleType          *int64                                                      `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	BizId2cList         []*int64                                                    `json:"biz_id_2c_list,omitempty" xml:"biz_id_2c_list,omitempty" type:"Repeated"`
	LifeBizCode         *string                                                     `json:"life_biz_code,omitempty" xml:"life_biz_code,omitempty"`
	TakeawayPresaleInfo *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo `json:"takeaway_presale_info,omitempty" xml:"takeaway_presale_info,omitempty"`
	DiscountPromo       *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo       `json:"discount_promo,omitempty" xml:"discount_promo,omitempty"`
	CMspuId             *int64                                                      `json:"c_mspu_id,omitempty" xml:"c_mspu_id,omitempty"`
	BindSkus2c          map[int64][]*int64                                          `json:"bind_skus_2c,omitempty" xml:"bind_skus_2c,omitempty"`
	AccountSettle       *bool                                                       `json:"account_settle,omitempty" xml:"account_settle,omitempty"`
	BindSkuId           *int64                                                      `json:"bind_sku_id,omitempty" xml:"bind_sku_id,omitempty"`
	OriSkus             map[int64][]*int64                                          `json:"ori_skus,omitempty" xml:"ori_skus,omitempty"`
	OriginStockQty      *int64                                                      `json:"origin_stock_qty,omitempty" xml:"origin_stock_qty,omitempty"`
	BindProductId       *int64                                                      `json:"bind_product_id,omitempty" xml:"bind_product_id,omitempty"`
}

func (s V1GoodsProductSaveRequestSkusItemSkuExt) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItemSkuExt) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetUseSubRelStock(v bool) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.UseSubRelStock = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetRelRuleList(v []*V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.RelRuleList = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetSettleType(v int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.SettleType = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetBizId2cList(v []*int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.BizId2cList = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetLifeBizCode(v string) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.LifeBizCode = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetTakeawayPresaleInfo(v *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.TakeawayPresaleInfo = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetDiscountPromo(v *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.DiscountPromo = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetCMspuId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.CMspuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetBindSkus2c(v map[int64][]*int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.BindSkus2c = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetAccountSettle(v bool) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.AccountSettle = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetBindSkuId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.BindSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetOriSkus(v map[int64][]*int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.OriSkus = v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetOriginStockQty(v int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.OriginStockQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExt) SetBindProductId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExt {
	s.BindProductId = &v
	return s
}

type V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo struct {
	PromoId         *int64 `json:"PromoId,omitempty" xml:"PromoId,omitempty"`
	BrandActivityId *int64 `json:"BrandActivityId,omitempty" xml:"BrandActivityId,omitempty"`
	PlanId          *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) SetPromoId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo {
	s.PromoId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) SetBrandActivityId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo {
	s.BrandActivityId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo) SetPlanId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtDiscountPromo {
	s.PlanId = &v
	return s
}

type V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem struct {
	SharedQty   *int64  `json:"SharedQty,omitempty" xml:"SharedQty,omitempty"`
	StockRel    *bool   `json:"StockRel,omitempty" xml:"StockRel,omitempty"`
	BizId       *int64  `json:"BizId,omitempty" xml:"BizId,omitempty" require:"true"`
	Coefficient *string `json:"Coefficient,omitempty" xml:"Coefficient,omitempty"`
	ConstantVal *int64  `json:"ConstantVal,omitempty" xml:"ConstantVal,omitempty"`
	PriceRel    *bool   `json:"PriceRel,omitempty" xml:"PriceRel,omitempty"`
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetSharedQty(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.SharedQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetStockRel(v bool) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.StockRel = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetBizId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.BizId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetCoefficient(v string) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.Coefficient = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetConstantVal(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.ConstantVal = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem) SetPriceRel(v bool) *V1GoodsProductSaveRequestSkusItemSkuExtRelRuleListItem {
	s.PriceRel = &v
	return s
}

type V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo struct {
	TakeawayPresaleSkuId     *int64 `json:"takeaway_presale_sku_id,omitempty" xml:"takeaway_presale_sku_id,omitempty" require:"true"`
	TakeawayPresaleProductId *int64 `json:"takeaway_presale_product_id,omitempty" xml:"takeaway_presale_product_id,omitempty" require:"true"`
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleSkuId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleSkuId = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo) SetTakeawayPresaleProductId(v int64) *V1GoodsProductSaveRequestSkusItemSkuExtTakeawayPresaleInfo {
	s.TakeawayPresaleProductId = &v
	return s
}

type V1GoodsProductSaveRequestSkusItemStock struct {
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
}

func (s V1GoodsProductSaveRequestSkusItemStock) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveRequestSkusItemStock) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetLimitType(v int) *V1GoodsProductSaveRequestSkusItemStock {
	s.LimitType = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetSoldCount(v int64) *V1GoodsProductSaveRequestSkusItemStock {
	s.SoldCount = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetSoldQty(v int64) *V1GoodsProductSaveRequestSkusItemStock {
	s.SoldQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetStockQty(v int64) *V1GoodsProductSaveRequestSkusItemStock {
	s.StockQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetAvailQty(v int64) *V1GoodsProductSaveRequestSkusItemStock {
	s.AvailQty = &v
	return s
}

func (s *V1GoodsProductSaveRequestSkusItemStock) SetFrozenQty(v int64) *V1GoodsProductSaveRequestSkusItemStock {
	s.FrozenQty = &v
	return s
}

type V1GoodsProductSaveResponse struct {
	BaseResp *V1GoodsProductSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *V1GoodsProductSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *V1GoodsProductSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s V1GoodsProductSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveResponse) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveResponse) SetBaseResp(v *V1GoodsProductSaveResponseBaseResp) *V1GoodsProductSaveResponse {
	s.BaseResp = v
	return s
}

func (s *V1GoodsProductSaveResponse) SetData(v *V1GoodsProductSaveResponseData) *V1GoodsProductSaveResponse {
	s.Data = v
	return s
}

func (s *V1GoodsProductSaveResponse) SetExtra(v *V1GoodsProductSaveResponseExtra) *V1GoodsProductSaveResponse {
	s.Extra = v
	return s
}

type V1GoodsProductSaveResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s V1GoodsProductSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveResponseBaseResp) SetExtra(v map[string]*string) *V1GoodsProductSaveResponseBaseResp {
	s.Extra = v
	return s
}

func (s *V1GoodsProductSaveResponseBaseResp) SetStatusCode(v int32) *V1GoodsProductSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *V1GoodsProductSaveResponseBaseResp) SetStatusMessage(v string) *V1GoodsProductSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type V1GoodsProductSaveResponseData struct {
	ExtraMap    map[string]*string `json:"extra_map,omitempty" xml:"extra_map,omitempty"`
	ProductId   *string            `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SkuIds      map[string]*string `json:"sku_ids,omitempty" xml:"sku_ids,omitempty"`
	Description *string            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s V1GoodsProductSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveResponseData) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveResponseData) SetExtraMap(v map[string]*string) *V1GoodsProductSaveResponseData {
	s.ExtraMap = v
	return s
}

func (s *V1GoodsProductSaveResponseData) SetProductId(v string) *V1GoodsProductSaveResponseData {
	s.ProductId = &v
	return s
}

func (s *V1GoodsProductSaveResponseData) SetSkuIds(v map[string]*string) *V1GoodsProductSaveResponseData {
	s.SkuIds = v
	return s
}

func (s *V1GoodsProductSaveResponseData) SetDescription(v string) *V1GoodsProductSaveResponseData {
	s.Description = &v
	return s
}

func (s *V1GoodsProductSaveResponseData) SetErrorCode(v int32) *V1GoodsProductSaveResponseData {
	s.ErrorCode = &v
	return s
}

type V1GoodsProductSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s V1GoodsProductSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsProductSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *V1GoodsProductSaveResponseExtra) SetDescription(v string) *V1GoodsProductSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *V1GoodsProductSaveResponseExtra) SetErrorCode(v int32) *V1GoodsProductSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *V1GoodsProductSaveResponseExtra) SetLogid(v string) *V1GoodsProductSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *V1GoodsProductSaveResponseExtra) SetNow(v int64) *V1GoodsProductSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *V1GoodsProductSaveResponseExtra) SetSubDescription(v string) *V1GoodsProductSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *V1GoodsProductSaveResponseExtra) SetSubErrorCode(v int32) *V1GoodsProductSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type V1GoodsStockSyncRequest struct {
	Stock       *V1GoodsStockSyncRequestStock `json:"stock,omitempty" xml:"stock,omitempty" require:"true"`
	Base        *V1GoodsStockSyncRequestBase  `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId   *string                       `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header      map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OutId       *string                       `json:"out_id,omitempty" xml:"out_id,omitempty"`
	ProductId   *string                       `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s V1GoodsStockSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncRequest) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncRequest) SetStock(v *V1GoodsStockSyncRequestStock) *V1GoodsStockSyncRequest {
	s.Stock = v
	return s
}

func (s *V1GoodsStockSyncRequest) SetBase(v *V1GoodsStockSyncRequestBase) *V1GoodsStockSyncRequest {
	s.Base = v
	return s
}

func (s *V1GoodsStockSyncRequest) SetAccountId(v string) *V1GoodsStockSyncRequest {
	s.AccountId = &v
	return s
}

func (s *V1GoodsStockSyncRequest) SetHeader(v map[string]*string) *V1GoodsStockSyncRequest {
	s.Header = v
	return s
}

func (s *V1GoodsStockSyncRequest) SetAccessToken(v string) *V1GoodsStockSyncRequest {
	s.AccessToken = &v
	return s
}

func (s *V1GoodsStockSyncRequest) SetOutId(v string) *V1GoodsStockSyncRequest {
	s.OutId = &v
	return s
}

func (s *V1GoodsStockSyncRequest) SetProductId(v string) *V1GoodsStockSyncRequest {
	s.ProductId = &v
	return s
}

type V1GoodsStockSyncRequestBase struct {
	TrafficEnv *V1GoodsStockSyncRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                     `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                `json:"LogID,omitempty" xml:"LogID,omitempty"`
}

func (s V1GoodsStockSyncRequestBase) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncRequestBase) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncRequestBase) SetTrafficEnv(v *V1GoodsStockSyncRequestBaseTrafficEnv) *V1GoodsStockSyncRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *V1GoodsStockSyncRequestBase) SetAddr(v string) *V1GoodsStockSyncRequestBase {
	s.Addr = &v
	return s
}

func (s *V1GoodsStockSyncRequestBase) SetCaller(v string) *V1GoodsStockSyncRequestBase {
	s.Caller = &v
	return s
}

func (s *V1GoodsStockSyncRequestBase) SetClient(v string) *V1GoodsStockSyncRequestBase {
	s.Client = &v
	return s
}

func (s *V1GoodsStockSyncRequestBase) SetExtra(v map[string]*string) *V1GoodsStockSyncRequestBase {
	s.Extra = v
	return s
}

func (s *V1GoodsStockSyncRequestBase) SetLogID(v string) *V1GoodsStockSyncRequestBase {
	s.LogID = &v
	return s
}

type V1GoodsStockSyncRequestBaseTrafficEnv struct {
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s V1GoodsStockSyncRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncRequestBaseTrafficEnv) SetOpen(v bool) *V1GoodsStockSyncRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

func (s *V1GoodsStockSyncRequestBaseTrafficEnv) SetEnv(v string) *V1GoodsStockSyncRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

type V1GoodsStockSyncRequestStock struct {
	SoldCount *int64 `json:"sold_count,omitempty" xml:"sold_count,omitempty"`
	SoldQty   *int64 `json:"sold_qty,omitempty" xml:"sold_qty,omitempty"`
	StockQty  *int64 `json:"stock_qty,omitempty" xml:"stock_qty,omitempty"`
	AvailQty  *int64 `json:"avail_qty,omitempty" xml:"avail_qty,omitempty"`
	FrozenQty *int64 `json:"frozen_qty,omitempty" xml:"frozen_qty,omitempty"`
	LimitType *int   `json:"limit_type,omitempty" xml:"limit_type,omitempty" require:"true"`
}

func (s V1GoodsStockSyncRequestStock) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncRequestStock) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncRequestStock) SetSoldCount(v int64) *V1GoodsStockSyncRequestStock {
	s.SoldCount = &v
	return s
}

func (s *V1GoodsStockSyncRequestStock) SetSoldQty(v int64) *V1GoodsStockSyncRequestStock {
	s.SoldQty = &v
	return s
}

func (s *V1GoodsStockSyncRequestStock) SetStockQty(v int64) *V1GoodsStockSyncRequestStock {
	s.StockQty = &v
	return s
}

func (s *V1GoodsStockSyncRequestStock) SetAvailQty(v int64) *V1GoodsStockSyncRequestStock {
	s.AvailQty = &v
	return s
}

func (s *V1GoodsStockSyncRequestStock) SetFrozenQty(v int64) *V1GoodsStockSyncRequestStock {
	s.FrozenQty = &v
	return s
}

func (s *V1GoodsStockSyncRequestStock) SetLimitType(v int) *V1GoodsStockSyncRequestStock {
	s.LimitType = &v
	return s
}

type V1GoodsStockSyncResponse struct {
	BaseResp *V1GoodsStockSyncResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *V1GoodsStockSyncResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *V1GoodsStockSyncResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s V1GoodsStockSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncResponse) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncResponse) SetBaseResp(v *V1GoodsStockSyncResponseBaseResp) *V1GoodsStockSyncResponse {
	s.BaseResp = v
	return s
}

func (s *V1GoodsStockSyncResponse) SetData(v *V1GoodsStockSyncResponseData) *V1GoodsStockSyncResponse {
	s.Data = v
	return s
}

func (s *V1GoodsStockSyncResponse) SetExtra(v *V1GoodsStockSyncResponseExtra) *V1GoodsStockSyncResponse {
	s.Extra = v
	return s
}

type V1GoodsStockSyncResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s V1GoodsStockSyncResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncResponseBaseResp) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncResponseBaseResp) SetStatusMessage(v string) *V1GoodsStockSyncResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *V1GoodsStockSyncResponseBaseResp) SetExtra(v map[string]*string) *V1GoodsStockSyncResponseBaseResp {
	s.Extra = v
	return s
}

func (s *V1GoodsStockSyncResponseBaseResp) SetStatusCode(v int32) *V1GoodsStockSyncResponseBaseResp {
	s.StatusCode = &v
	return s
}

type V1GoodsStockSyncResponseData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s V1GoodsStockSyncResponseData) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncResponseData) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncResponseData) SetDescription(v string) *V1GoodsStockSyncResponseData {
	s.Description = &v
	return s
}

func (s *V1GoodsStockSyncResponseData) SetErrorCode(v int32) *V1GoodsStockSyncResponseData {
	s.ErrorCode = &v
	return s
}

type V1GoodsStockSyncResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s V1GoodsStockSyncResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s V1GoodsStockSyncResponseExtra) GoString() string {
	return s.String()
}

func (s *V1GoodsStockSyncResponseExtra) SetErrorCode(v int32) *V1GoodsStockSyncResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *V1GoodsStockSyncResponseExtra) SetLogid(v string) *V1GoodsStockSyncResponseExtra {
	s.Logid = &v
	return s
}

func (s *V1GoodsStockSyncResponseExtra) SetNow(v int64) *V1GoodsStockSyncResponseExtra {
	s.Now = &v
	return s
}

func (s *V1GoodsStockSyncResponseExtra) SetSubDescription(v string) *V1GoodsStockSyncResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *V1GoodsStockSyncResponseExtra) SetSubErrorCode(v int32) *V1GoodsStockSyncResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *V1GoodsStockSyncResponseExtra) SetDescription(v string) *V1GoodsStockSyncResponseExtra {
	s.Description = &v
	return s
}

type V1ListKolOrderRequest struct {
	CreateTimeEnd   *int64             `json:"create_time_end,omitempty" xml:"create_time_end,omitempty"`
	Cursor          *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageSize        *int64             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	KolId           *string            `json:"kol_id,omitempty" xml:"kol_id,omitempty" require:"true"`
	OpenId          *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	CreateTimeStart *int64             `json:"create_time_start,omitempty" xml:"create_time_start,omitempty"`
}

func (s V1ListKolOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s V1ListKolOrderRequest) GoString() string {
	return s.String()
}

func (s *V1ListKolOrderRequest) SetCreateTimeEnd(v int64) *V1ListKolOrderRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *V1ListKolOrderRequest) SetCursor(v int64) *V1ListKolOrderRequest {
	s.Cursor = &v
	return s
}

func (s *V1ListKolOrderRequest) SetHeader(v map[string]*string) *V1ListKolOrderRequest {
	s.Header = v
	return s
}

func (s *V1ListKolOrderRequest) SetAccessToken(v string) *V1ListKolOrderRequest {
	s.AccessToken = &v
	return s
}

func (s *V1ListKolOrderRequest) SetPageSize(v int64) *V1ListKolOrderRequest {
	s.PageSize = &v
	return s
}

func (s *V1ListKolOrderRequest) SetKolId(v string) *V1ListKolOrderRequest {
	s.KolId = &v
	return s
}

func (s *V1ListKolOrderRequest) SetOpenId(v string) *V1ListKolOrderRequest {
	s.OpenId = &v
	return s
}

func (s *V1ListKolOrderRequest) SetCreateTimeStart(v int64) *V1ListKolOrderRequest {
	s.CreateTimeStart = &v
	return s
}

type V1ListKolOrderResponse struct {
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *V1ListKolOrderResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s V1ListKolOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s V1ListKolOrderResponse) GoString() string {
	return s.String()
}

func (s *V1ListKolOrderResponse) SetLogId(v string) *V1ListKolOrderResponse {
	s.LogId = &v
	return s
}

func (s *V1ListKolOrderResponse) SetData(v *V1ListKolOrderResponseData) *V1ListKolOrderResponse {
	s.Data = v
	return s
}

func (s *V1ListKolOrderResponse) SetErrNo(v int32) *V1ListKolOrderResponse {
	s.ErrNo = &v
	return s
}

func (s *V1ListKolOrderResponse) SetErrMsg(v string) *V1ListKolOrderResponse {
	s.ErrMsg = &v
	return s
}

type V1ListKolOrderResponseData struct {
	Orders []*V1ListKolOrderResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
	Cursor *int64                                  `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s V1ListKolOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s V1ListKolOrderResponseData) GoString() string {
	return s.String()
}

func (s *V1ListKolOrderResponseData) SetOrders(v []*V1ListKolOrderResponseDataOrdersItem) *V1ListKolOrderResponseData {
	s.Orders = v
	return s
}

func (s *V1ListKolOrderResponseData) SetCursor(v int64) *V1ListKolOrderResponseData {
	s.Cursor = &v
	return s
}

type V1ListKolOrderResponseDataOrdersItem struct {
	OrderId        *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VirtualOrderId *string `json:"virtual_order_id,omitempty" xml:"virtual_order_id,omitempty"`
	PayTime        *int64  `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	TotalPayAmount *int32  `json:"total_pay_amount,omitempty" xml:"total_pay_amount,omitempty"`
	CreateTime     *int64  `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s V1ListKolOrderResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s V1ListKolOrderResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetOrderId(v string) *V1ListKolOrderResponseDataOrdersItem {
	s.OrderId = &v
	return s
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetVirtualOrderId(v string) *V1ListKolOrderResponseDataOrdersItem {
	s.VirtualOrderId = &v
	return s
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetPayTime(v int64) *V1ListKolOrderResponseDataOrdersItem {
	s.PayTime = &v
	return s
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetTotalPayAmount(v int32) *V1ListKolOrderResponseDataOrdersItem {
	s.TotalPayAmount = &v
	return s
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetCreateTime(v int64) *V1ListKolOrderResponseDataOrdersItem {
	s.CreateTime = &v
	return s
}

func (s *V1ListKolOrderResponseDataOrdersItem) SetStatus(v string) *V1ListKolOrderResponseDataOrdersItem {
	s.Status = &v
	return s
}

type V1NotifyRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Page        *string            `json:"page,omitempty" xml:"page,omitempty"`
	Data        map[string]*string `json:"data,omitempty" xml:"data,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	TplId       *string            `json:"tpl_id,omitempty" xml:"tpl_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V1NotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s V1NotifyRequest) GoString() string {
	return s.String()
}

func (s *V1NotifyRequest) SetOpenId(v string) *V1NotifyRequest {
	s.OpenId = &v
	return s
}

func (s *V1NotifyRequest) SetPage(v string) *V1NotifyRequest {
	s.Page = &v
	return s
}

func (s *V1NotifyRequest) SetData(v map[string]*string) *V1NotifyRequest {
	s.Data = v
	return s
}

func (s *V1NotifyRequest) SetAppId(v string) *V1NotifyRequest {
	s.AppId = &v
	return s
}

func (s *V1NotifyRequest) SetTplId(v string) *V1NotifyRequest {
	s.TplId = &v
	return s
}

func (s *V1NotifyRequest) SetHeader(v map[string]*string) *V1NotifyRequest {
	s.Header = v
	return s
}

func (s *V1NotifyRequest) SetAccessToken(v string) *V1NotifyRequest {
	s.AccessToken = &v
	return s
}

type V1NotifyResponse struct {
	ErrNo   *int64  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrTips *string `json:"err_tips,omitempty" xml:"err_tips,omitempty"`
}

func (s V1NotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s V1NotifyResponse) GoString() string {
	return s.String()
}

func (s *V1NotifyResponse) SetErrNo(v int64) *V1NotifyResponse {
	s.ErrNo = &v
	return s
}

func (s *V1NotifyResponse) SetErrTips(v string) *V1NotifyResponse {
	s.ErrTips = &v
	return s
}

type V2CouponCreateCouponMetaRequest struct {
	BizType     *int                                       `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	Header      map[string]*string                         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CouponMeta  *V2CouponCreateCouponMetaRequestCouponMeta `json:"coupon_meta,omitempty" xml:"coupon_meta,omitempty" require:"true"`
}

func (s V2CouponCreateCouponMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s V2CouponCreateCouponMetaRequest) GoString() string {
	return s.String()
}

func (s *V2CouponCreateCouponMetaRequest) SetBizType(v int) *V2CouponCreateCouponMetaRequest {
	s.BizType = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequest) SetHeader(v map[string]*string) *V2CouponCreateCouponMetaRequest {
	s.Header = v
	return s
}

func (s *V2CouponCreateCouponMetaRequest) SetAccessToken(v string) *V2CouponCreateCouponMetaRequest {
	s.AccessToken = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequest) SetCouponMeta(v *V2CouponCreateCouponMetaRequestCouponMeta) *V2CouponCreateCouponMetaRequest {
	s.CouponMeta = v
	return s
}

type V2CouponCreateCouponMetaRequestCouponMeta struct {
	ReceiveBeginTime *int64  `json:"receive_begin_time,omitempty" xml:"receive_begin_time,omitempty" require:"true"`
	FreeEpNumber     *int64  `json:"free_ep_number,omitempty" xml:"free_ep_number,omitempty"`
	ConsumeDesc      *string `json:"consume_desc,omitempty" xml:"consume_desc,omitempty" require:"true"`
	ConsumePath      *string `json:"consume_path,omitempty" xml:"consume_path,omitempty" require:"true"`
	DiscountAmount   *int64  `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	ValidBeginTime   *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	ReceiveDesc      *string `json:"receive_desc,omitempty" xml:"receive_desc,omitempty"`
	ValidDuration    *int64  `json:"valid_duration,omitempty" xml:"valid_duration,omitempty"`
	CouponName       *string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty" require:"true"`
	CallbackUrl      *string `json:"callback_url,omitempty" xml:"callback_url,omitempty"`
	OriginId         *string `json:"origin_id,omitempty" xml:"origin_id,omitempty"`
	StockNumber      *int64  `json:"stock_number,omitempty" xml:"stock_number,omitempty" require:"true"`
	ValidType        *int    `json:"valid_type,omitempty" xml:"valid_type,omitempty" require:"true"`
	RelatedType      *int    `json:"related_type,omitempty" xml:"related_type,omitempty"`
	ValidEndTime     *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	DiscountType     *int    `json:"discount_type,omitempty" xml:"discount_type,omitempty" require:"true"`
	MinPayAmount     *int64  `json:"min_pay_amount,omitempty" xml:"min_pay_amount,omitempty"`
	SecretSource     *int    `json:"secret_source,omitempty" xml:"secret_source,omitempty"`
	ReceiveEndTime   *int64  `json:"receive_end_time,omitempty" xml:"receive_end_time,omitempty" require:"true"`
	MerchantMetaNo   *string `json:"merchant_meta_no,omitempty" xml:"merchant_meta_no,omitempty" require:"true"`
}

func (s V2CouponCreateCouponMetaRequestCouponMeta) String() string {
	return tea.Prettify(s)
}

func (s V2CouponCreateCouponMetaRequestCouponMeta) GoString() string {
	return s.String()
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetReceiveBeginTime(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ReceiveBeginTime = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetFreeEpNumber(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.FreeEpNumber = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetConsumeDesc(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ConsumeDesc = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetConsumePath(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ConsumePath = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetDiscountAmount(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.DiscountAmount = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetValidBeginTime(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ValidBeginTime = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetReceiveDesc(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ReceiveDesc = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetValidDuration(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ValidDuration = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetCouponName(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.CouponName = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetCallbackUrl(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.CallbackUrl = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetOriginId(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.OriginId = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetStockNumber(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.StockNumber = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetValidType(v int) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ValidType = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetRelatedType(v int) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.RelatedType = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetValidEndTime(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ValidEndTime = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetDiscountType(v int) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.DiscountType = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetMinPayAmount(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.MinPayAmount = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetSecretSource(v int) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.SecretSource = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetReceiveEndTime(v int64) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.ReceiveEndTime = &v
	return s
}

func (s *V2CouponCreateCouponMetaRequestCouponMeta) SetMerchantMetaNo(v string) *V2CouponCreateCouponMetaRequestCouponMeta {
	s.MerchantMetaNo = &v
	return s
}

type V2CouponCreateCouponMetaResponse struct {
	Data   *V2CouponCreateCouponMetaResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s V2CouponCreateCouponMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s V2CouponCreateCouponMetaResponse) GoString() string {
	return s.String()
}

func (s *V2CouponCreateCouponMetaResponse) SetData(v *V2CouponCreateCouponMetaResponseData) *V2CouponCreateCouponMetaResponse {
	s.Data = v
	return s
}

func (s *V2CouponCreateCouponMetaResponse) SetErrNo(v int32) *V2CouponCreateCouponMetaResponse {
	s.ErrNo = &v
	return s
}

func (s *V2CouponCreateCouponMetaResponse) SetErrMsg(v string) *V2CouponCreateCouponMetaResponse {
	s.ErrMsg = &v
	return s
}

func (s *V2CouponCreateCouponMetaResponse) SetLogId(v string) *V2CouponCreateCouponMetaResponse {
	s.LogId = &v
	return s
}

type V2CouponCreateCouponMetaResponseData struct {
	CouponMetaId *string `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
}

func (s V2CouponCreateCouponMetaResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2CouponCreateCouponMetaResponseData) GoString() string {
	return s.String()
}

func (s *V2CouponCreateCouponMetaResponseData) SetCouponMetaId(v string) *V2CouponCreateCouponMetaResponseData {
	s.CouponMetaId = &v
	return s
}

type V2FileUploadMaterialRequest struct {
	MaterialType *int32             `json:"material_type,omitempty" xml:"material_type,omitempty" require:"true"`
	MaterialFile *util.FileField    `json:"material_file,omitempty" xml:"material_file,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V2FileUploadMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s V2FileUploadMaterialRequest) GoString() string {
	return s.String()
}

func (s *V2FileUploadMaterialRequest) SetMaterialType(v int32) *V2FileUploadMaterialRequest {
	s.MaterialType = &v
	return s
}

func (s *V2FileUploadMaterialRequest) SetMaterialFile(v *util.FileField) *V2FileUploadMaterialRequest {
	s.MaterialFile = v
	return s
}

func (s *V2FileUploadMaterialRequest) SetHeader(v map[string]*string) *V2FileUploadMaterialRequest {
	s.Header = v
	return s
}

func (s *V2FileUploadMaterialRequest) SetAccessToken(v string) *V2FileUploadMaterialRequest {
	s.AccessToken = &v
	return s
}

type V2FileUploadMaterialResponse struct {
	Data   *V2FileUploadMaterialResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s V2FileUploadMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s V2FileUploadMaterialResponse) GoString() string {
	return s.String()
}

func (s *V2FileUploadMaterialResponse) SetData(v *V2FileUploadMaterialResponseData) *V2FileUploadMaterialResponse {
	s.Data = v
	return s
}

func (s *V2FileUploadMaterialResponse) SetErrNo(v int32) *V2FileUploadMaterialResponse {
	s.ErrNo = &v
	return s
}

func (s *V2FileUploadMaterialResponse) SetErrMsg(v string) *V2FileUploadMaterialResponse {
	s.ErrMsg = &v
	return s
}

func (s *V2FileUploadMaterialResponse) SetLogId(v string) *V2FileUploadMaterialResponse {
	s.LogId = &v
	return s
}

type V2FileUploadMaterialResponseData struct {
	Path *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s V2FileUploadMaterialResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2FileUploadMaterialResponseData) GoString() string {
	return s.String()
}

func (s *V2FileUploadMaterialResponseData) SetPath(v string) *V2FileUploadMaterialResponseData {
	s.Path = &v
	return s
}

type V2Jscode2sessionRequest struct {
	Code          *string            `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	AnonymousCode *string            `json:"anonymous_code,omitempty" xml:"anonymous_code,omitempty"`
	Appid         *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Secret        *string            `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s V2Jscode2sessionRequest) String() string {
	return tea.Prettify(s)
}

func (s V2Jscode2sessionRequest) GoString() string {
	return s.String()
}

func (s *V2Jscode2sessionRequest) SetCode(v string) *V2Jscode2sessionRequest {
	s.Code = &v
	return s
}

func (s *V2Jscode2sessionRequest) SetAnonymousCode(v string) *V2Jscode2sessionRequest {
	s.AnonymousCode = &v
	return s
}

func (s *V2Jscode2sessionRequest) SetAppid(v string) *V2Jscode2sessionRequest {
	s.Appid = &v
	return s
}

func (s *V2Jscode2sessionRequest) SetSecret(v string) *V2Jscode2sessionRequest {
	s.Secret = &v
	return s
}

func (s *V2Jscode2sessionRequest) SetHeader(v map[string]*string) *V2Jscode2sessionRequest {
	s.Header = v
	return s
}

type V2Jscode2sessionResponse struct {
	ErrTips *string                       `json:"err_tips,omitempty" xml:"err_tips,omitempty"`
	Data    *V2Jscode2sessionResponseData `json:"data,omitempty" xml:"data,omitempty"`
	LogId   *string                       `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ErrNo   *int64                        `json:"err_no,omitempty" xml:"err_no,omitempty"`
}

func (s V2Jscode2sessionResponse) String() string {
	return tea.Prettify(s)
}

func (s V2Jscode2sessionResponse) GoString() string {
	return s.String()
}

func (s *V2Jscode2sessionResponse) SetErrTips(v string) *V2Jscode2sessionResponse {
	s.ErrTips = &v
	return s
}

func (s *V2Jscode2sessionResponse) SetData(v *V2Jscode2sessionResponseData) *V2Jscode2sessionResponse {
	s.Data = v
	return s
}

func (s *V2Jscode2sessionResponse) SetLogId(v string) *V2Jscode2sessionResponse {
	s.LogId = &v
	return s
}

func (s *V2Jscode2sessionResponse) SetErrNo(v int64) *V2Jscode2sessionResponse {
	s.ErrNo = &v
	return s
}

type V2Jscode2sessionResponseData struct {
	AnonymousOpenid *string `json:"anonymous_openid,omitempty" xml:"anonymous_openid,omitempty"`
	Unionid         *string `json:"unionid,omitempty" xml:"unionid,omitempty"`
	SessionKey      *string `json:"session_key,omitempty" xml:"session_key,omitempty"`
	Openid          *string `json:"openid,omitempty" xml:"openid,omitempty"`
}

func (s V2Jscode2sessionResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2Jscode2sessionResponseData) GoString() string {
	return s.String()
}

func (s *V2Jscode2sessionResponseData) SetAnonymousOpenid(v string) *V2Jscode2sessionResponseData {
	s.AnonymousOpenid = &v
	return s
}

func (s *V2Jscode2sessionResponseData) SetUnionid(v string) *V2Jscode2sessionResponseData {
	s.Unionid = &v
	return s
}

func (s *V2Jscode2sessionResponseData) SetSessionKey(v string) *V2Jscode2sessionResponseData {
	s.SessionKey = &v
	return s
}

func (s *V2Jscode2sessionResponseData) SetOpenid(v string) *V2Jscode2sessionResponseData {
	s.Openid = &v
	return s
}

type V2SearchVideoRequest struct {
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	DeviceId    *int64             `json:"device_id,omitempty" xml:"device_id,omitempty" require:"true"`
	Keyword     *string            `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	PublishTime *int32             `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	SearchId    *string            `json:"search_id,omitempty" xml:"search_id,omitempty"`
	SortType    *int32             `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V2SearchVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s V2SearchVideoRequest) GoString() string {
	return s.String()
}

func (s *V2SearchVideoRequest) SetCursor(v int64) *V2SearchVideoRequest {
	s.Cursor = &v
	return s
}

func (s *V2SearchVideoRequest) SetDeviceId(v int64) *V2SearchVideoRequest {
	s.DeviceId = &v
	return s
}

func (s *V2SearchVideoRequest) SetKeyword(v string) *V2SearchVideoRequest {
	s.Keyword = &v
	return s
}

func (s *V2SearchVideoRequest) SetOpenId(v string) *V2SearchVideoRequest {
	s.OpenId = &v
	return s
}

func (s *V2SearchVideoRequest) SetPublishTime(v int32) *V2SearchVideoRequest {
	s.PublishTime = &v
	return s
}

func (s *V2SearchVideoRequest) SetSearchId(v string) *V2SearchVideoRequest {
	s.SearchId = &v
	return s
}

func (s *V2SearchVideoRequest) SetSortType(v int32) *V2SearchVideoRequest {
	s.SortType = &v
	return s
}

func (s *V2SearchVideoRequest) SetCount(v int32) *V2SearchVideoRequest {
	s.Count = &v
	return s
}

func (s *V2SearchVideoRequest) SetHeader(v map[string]*string) *V2SearchVideoRequest {
	s.Header = v
	return s
}

func (s *V2SearchVideoRequest) SetAccessToken(v string) *V2SearchVideoRequest {
	s.AccessToken = &v
	return s
}

type V2SearchVideoResponse struct {
	Data   *V2SearchVideoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                     `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                    `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                    `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s V2SearchVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s V2SearchVideoResponse) GoString() string {
	return s.String()
}

func (s *V2SearchVideoResponse) SetData(v *V2SearchVideoResponseData) *V2SearchVideoResponse {
	s.Data = v
	return s
}

func (s *V2SearchVideoResponse) SetErrNo(v int32) *V2SearchVideoResponse {
	s.ErrNo = &v
	return s
}

func (s *V2SearchVideoResponse) SetErrMsg(v string) *V2SearchVideoResponse {
	s.ErrMsg = &v
	return s
}

func (s *V2SearchVideoResponse) SetLogId(v string) *V2SearchVideoResponse {
	s.LogId = &v
	return s
}

type V2SearchVideoResponseData struct {
	HasMore   *bool                                     `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	SearchId  *string                                   `json:"search_id,omitempty" xml:"search_id,omitempty" require:"true"`
	VideoList []*V2SearchVideoResponseDataVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" require:"true" type:"Repeated"`
	Cursor    *int64                                    `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
}

func (s V2SearchVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2SearchVideoResponseData) GoString() string {
	return s.String()
}

func (s *V2SearchVideoResponseData) SetHasMore(v bool) *V2SearchVideoResponseData {
	s.HasMore = &v
	return s
}

func (s *V2SearchVideoResponseData) SetSearchId(v string) *V2SearchVideoResponseData {
	s.SearchId = &v
	return s
}

func (s *V2SearchVideoResponseData) SetVideoList(v []*V2SearchVideoResponseDataVideoListItem) *V2SearchVideoResponseData {
	s.VideoList = v
	return s
}

func (s *V2SearchVideoResponseData) SetCursor(v int64) *V2SearchVideoResponseData {
	s.Cursor = &v
	return s
}

type V2SearchVideoResponseDataVideoListItem struct {
	Nickname    *string                                           `json:"nickname,omitempty" xml:"nickname,omitempty"`
	Duration    *int64                                            `json:"duration,omitempty" xml:"duration,omitempty"`
	Link        *string                                           `json:"link,omitempty" xml:"link,omitempty"`
	ItemId      *string                                           `json:"item_id,omitempty" xml:"item_id,omitempty"`
	CreateTime  *int64                                            `json:"create_time,omitempty" xml:"create_time,omitempty"`
	Title       *string                                           `json:"title,omitempty" xml:"title,omitempty"`
	Statistics  *V2SearchVideoResponseDataVideoListItemStatistics `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Cover       *string                                           `json:"cover,omitempty" xml:"cover,omitempty"`
	CoverWidth  *int32                                            `json:"cover_width,omitempty" xml:"cover_width,omitempty"`
	Avatar      *string                                           `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CoverHeight *int32                                            `json:"cover_height,omitempty" xml:"cover_height,omitempty"`
}

func (s V2SearchVideoResponseDataVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s V2SearchVideoResponseDataVideoListItem) GoString() string {
	return s.String()
}

func (s *V2SearchVideoResponseDataVideoListItem) SetNickname(v string) *V2SearchVideoResponseDataVideoListItem {
	s.Nickname = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetDuration(v int64) *V2SearchVideoResponseDataVideoListItem {
	s.Duration = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetLink(v string) *V2SearchVideoResponseDataVideoListItem {
	s.Link = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetItemId(v string) *V2SearchVideoResponseDataVideoListItem {
	s.ItemId = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetCreateTime(v int64) *V2SearchVideoResponseDataVideoListItem {
	s.CreateTime = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetTitle(v string) *V2SearchVideoResponseDataVideoListItem {
	s.Title = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetStatistics(v *V2SearchVideoResponseDataVideoListItemStatistics) *V2SearchVideoResponseDataVideoListItem {
	s.Statistics = v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetCover(v string) *V2SearchVideoResponseDataVideoListItem {
	s.Cover = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetCoverWidth(v int32) *V2SearchVideoResponseDataVideoListItem {
	s.CoverWidth = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetAvatar(v string) *V2SearchVideoResponseDataVideoListItem {
	s.Avatar = &v
	return s
}

func (s *V2SearchVideoResponseDataVideoListItem) SetCoverHeight(v int32) *V2SearchVideoResponseDataVideoListItem {
	s.CoverHeight = &v
	return s
}

type V2SearchVideoResponseDataVideoListItemStatistics struct {
	DiggCount *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
}

func (s V2SearchVideoResponseDataVideoListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s V2SearchVideoResponseDataVideoListItemStatistics) GoString() string {
	return s.String()
}

func (s *V2SearchVideoResponseDataVideoListItemStatistics) SetDiggCount(v int32) *V2SearchVideoResponseDataVideoListItemStatistics {
	s.DiggCount = &v
	return s
}

type V2TaskboxQueryTaskInfoRequest struct {
	TaskCategory       *int               `json:"task_category,omitempty" xml:"task_category,omitempty"`
	QueryParamsType    *int               `json:"query_params_type,omitempty" xml:"query_params_type,omitempty" require:"true"`
	QueryParamsContent *string            `json:"query_params_content,omitempty" xml:"query_params_content,omitempty" require:"true"`
	PageNum            *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	PageSize           *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Header             map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s V2TaskboxQueryTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoRequest) SetTaskCategory(v int) *V2TaskboxQueryTaskInfoRequest {
	s.TaskCategory = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetQueryParamsType(v int) *V2TaskboxQueryTaskInfoRequest {
	s.QueryParamsType = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetQueryParamsContent(v string) *V2TaskboxQueryTaskInfoRequest {
	s.QueryParamsContent = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetPageNum(v int32) *V2TaskboxQueryTaskInfoRequest {
	s.PageNum = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetPageSize(v int32) *V2TaskboxQueryTaskInfoRequest {
	s.PageSize = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetHeader(v map[string]*string) *V2TaskboxQueryTaskInfoRequest {
	s.Header = v
	return s
}

func (s *V2TaskboxQueryTaskInfoRequest) SetAccessToken(v string) *V2TaskboxQueryTaskInfoRequest {
	s.AccessToken = &v
	return s
}

type V2TaskboxQueryTaskInfoResponse struct {
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *V2TaskboxQueryTaskInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s V2TaskboxQueryTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoResponse) SetErrMsg(v string) *V2TaskboxQueryTaskInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponse) SetLogId(v string) *V2TaskboxQueryTaskInfoResponse {
	s.LogId = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponse) SetData(v *V2TaskboxQueryTaskInfoResponseData) *V2TaskboxQueryTaskInfoResponse {
	s.Data = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponse) SetErrNo(v int32) *V2TaskboxQueryTaskInfoResponse {
	s.ErrNo = &v
	return s
}

type V2TaskboxQueryTaskInfoResponseData struct {
	Total     *int64                                         `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	AppId     *string                                        `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Tasks     []*V2TaskboxQueryTaskInfoResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
	PageCount *int64                                         `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
}

func (s V2TaskboxQueryTaskInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoResponseData) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoResponseData) SetTotal(v int64) *V2TaskboxQueryTaskInfoResponseData {
	s.Total = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseData) SetAppId(v string) *V2TaskboxQueryTaskInfoResponseData {
	s.AppId = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseData) SetTasks(v []*V2TaskboxQueryTaskInfoResponseDataTasksItem) *V2TaskboxQueryTaskInfoResponseData {
	s.Tasks = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseData) SetPageCount(v int64) *V2TaskboxQueryTaskInfoResponseData {
	s.PageCount = &v
	return s
}

type V2TaskboxQueryTaskInfoResponseDataTasksItem struct {
	TalentMixPaymentAllocateRatio map[int32]*string                                                       `json:"talent_mix_payment_allocate_ratio,omitempty" xml:"talent_mix_payment_allocate_ratio,omitempty"`
	StartPage                     *string                                                                 `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
	TalentPaymentAllocateRatio    *int32                                                                  `json:"talent_payment_allocate_ratio,omitempty" xml:"talent_payment_allocate_ratio,omitempty"`
	TaskId                        *int64                                                                  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	ServiceRate                   *int32                                                                  `json:"service_rate,omitempty" xml:"service_rate,omitempty"`
	TaskType                      *int                                                                    `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	PageType                      *int                                                                    `json:"page_type,omitempty" xml:"page_type,omitempty" require:"true"`
	ReferVideoCaptures            []*string                                                               `json:"refer_video_captures,omitempty" xml:"refer_video_captures,omitempty" type:"Repeated"`
	CpaParam                      *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam                    `json:"cpa_param,omitempty" xml:"cpa_param,omitempty"`
	OrientedTalentRelList         []*V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem `json:"oriented_talent_rel_list,omitempty" xml:"oriented_talent_rel_list,omitempty" type:"Repeated"`
	TaskRefundPeriod              *int32                                                                  `json:"task_refund_period,omitempty" xml:"task_refund_period,omitempty"`
	RejectReason                  *string                                                                 `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	ReferMaCaptures               []*string                                                               `json:"refer_ma_captures,omitempty" xml:"refer_ma_captures,omitempty" type:"Repeated"`
	TaskSettleType                *int32                                                                  `json:"task_settle_type,omitempty" xml:"task_settle_type,omitempty" require:"true"`
	TaskStartTime                 *int64                                                                  `json:"task_start_time,omitempty" xml:"task_start_time,omitempty" require:"true"`
	MixPaymentAllocateRatio       map[int32]*string                                                       `json:"mix_payment_allocate_ratio,omitempty" xml:"mix_payment_allocate_ratio,omitempty"`
	PlatformAddressWeb            *string                                                                 `json:"platform_address_web,omitempty" xml:"platform_address_web,omitempty"`
	TaskTags                      []*string                                                               `json:"task_tags,omitempty" xml:"task_tags,omitempty" require:"true" type:"Repeated"`
	TaskName                      *string                                                                 `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	PaymentAllocateRatio          *int32                                                                  `json:"payment_allocate_ratio,omitempty" xml:"payment_allocate_ratio,omitempty"`
	TaskIcon                      *string                                                                 `json:"task_icon,omitempty" xml:"task_icon,omitempty" require:"true"`
	Status                        *int32                                                                  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	PlatformAddressApp            *string                                                                 `json:"platform_address_app,omitempty" xml:"platform_address_app,omitempty"`
	ReferGids                     []*int64                                                                `json:"refer_gids,omitempty" xml:"refer_gids,omitempty" type:"Repeated"`
	TaskEndTime                   *int64                                                                  `json:"task_end_time,omitempty" xml:"task_end_time,omitempty" require:"true"`
	CapitalAccount                *int                                                                    `json:"capital_account,omitempty" xml:"capital_account,omitempty"`
	ReferVideos                   []*string                                                               `json:"refer_videos,omitempty" xml:"refer_videos,omitempty" type:"Repeated"`
	TaskDesc                      *string                                                                 `json:"task_desc,omitempty" xml:"task_desc,omitempty" require:"true"`
	Appid                         *string                                                                 `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	AnchorTitle                   *string                                                                 `json:"anchor_title,omitempty" xml:"anchor_title,omitempty" require:"true"`
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTalentMixPaymentAllocateRatio(v map[int32]*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TalentMixPaymentAllocateRatio = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetStartPage(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.StartPage = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTalentPaymentAllocateRatio(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TalentPaymentAllocateRatio = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskId(v int64) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskId = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetServiceRate(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ServiceRate = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskType(v int) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskType = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetPageType(v int) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PageType = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetReferVideoCaptures(v []*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferVideoCaptures = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetCpaParam(v *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.CpaParam = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetOrientedTalentRelList(v []*V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.OrientedTalentRelList = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskRefundPeriod(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskRefundPeriod = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetRejectReason(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.RejectReason = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetReferMaCaptures(v []*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferMaCaptures = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskSettleType(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskSettleType = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskStartTime(v int64) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskStartTime = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetMixPaymentAllocateRatio(v map[int32]*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.MixPaymentAllocateRatio = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetPlatformAddressWeb(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PlatformAddressWeb = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskTags(v []*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskTags = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskName(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskName = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetPaymentAllocateRatio(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PaymentAllocateRatio = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskIcon(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskIcon = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetStatus(v int32) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.Status = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetPlatformAddressApp(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PlatformAddressApp = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetReferGids(v []*int64) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferGids = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskEndTime(v int64) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskEndTime = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetCapitalAccount(v int) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.CapitalAccount = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetReferVideos(v []*string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferVideos = v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskDesc(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskDesc = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetAppid(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.Appid = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItem) SetAnchorTitle(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItem {
	s.AnchorTitle = &v
	return s
}

type V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam struct {
	TalentSinglePrice          *int64 `json:"talent_single_price,omitempty" xml:"talent_single_price,omitempty"`
	SinglePriceWithServiceRate *int64 `json:"single_price_with_service_rate,omitempty" xml:"single_price_with_service_rate,omitempty"`
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam) SetTalentSinglePrice(v int64) *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam {
	s.TalentSinglePrice = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam) SetSinglePriceWithServiceRate(v int64) *V2TaskboxQueryTaskInfoResponseDataTasksItemCpaParam {
	s.SinglePriceWithServiceRate = &v
	return s
}

type V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem struct {
	DouyinId         *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	CooperationState *int    `json:"cooperation_state,omitempty" xml:"cooperation_state,omitempty" require:"true"`
	CancelOperator   *int    `json:"cancel_operator,omitempty" xml:"cancel_operator,omitempty" require:"true"`
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) String() string {
	return tea.Prettify(s)
}

func (s V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) GoString() string {
	return s.String()
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetDouyinId(v string) *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.DouyinId = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetCooperationState(v int) *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.CooperationState = &v
	return s
}

func (s *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetCancelOperator(v int) *V2TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.CancelOperator = &v
	return s
}

type V2TokenRequest struct {
	GrantType *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	Appid     *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Secret    *string            `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	Header    map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s V2TokenRequest) String() string {
	return tea.Prettify(s)
}

func (s V2TokenRequest) GoString() string {
	return s.String()
}

func (s *V2TokenRequest) SetGrantType(v string) *V2TokenRequest {
	s.GrantType = &v
	return s
}

func (s *V2TokenRequest) SetAppid(v string) *V2TokenRequest {
	s.Appid = &v
	return s
}

func (s *V2TokenRequest) SetSecret(v string) *V2TokenRequest {
	s.Secret = &v
	return s
}

func (s *V2TokenRequest) SetHeader(v map[string]*string) *V2TokenRequest {
	s.Header = v
	return s
}

type V2TokenResponse struct {
	Data    *V2TokenResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo   *int                 `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrTips *string              `json:"err_tips,omitempty" xml:"err_tips,omitempty"`
}

func (s V2TokenResponse) String() string {
	return tea.Prettify(s)
}

func (s V2TokenResponse) GoString() string {
	return s.String()
}

func (s *V2TokenResponse) SetData(v *V2TokenResponseData) *V2TokenResponse {
	s.Data = v
	return s
}

func (s *V2TokenResponse) SetErrNo(v int) *V2TokenResponse {
	s.ErrNo = &v
	return s
}

func (s *V2TokenResponse) SetErrTips(v string) *V2TokenResponse {
	s.ErrTips = &v
	return s
}

type V2TokenResponseData struct {
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

func (s V2TokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s V2TokenResponseData) GoString() string {
	return s.String()
}

func (s *V2TokenResponseData) SetAccessToken(v string) *V2TokenResponseData {
	s.AccessToken = &v
	return s
}

func (s *V2TokenResponseData) SetExpiresIn(v int64) *V2TokenResponseData {
	s.ExpiresIn = &v
	return s
}

type V3BillsRequest struct {
	BillType    *string            `json:"bill_type,omitempty" xml:"bill_type,omitempty" require:"true"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	MerchantId  *string            `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
}

func (s V3BillsRequest) String() string {
	return tea.Prettify(s)
}

func (s V3BillsRequest) GoString() string {
	return s.String()
}

func (s *V3BillsRequest) SetBillType(v string) *V3BillsRequest {
	s.BillType = &v
	return s
}

func (s *V3BillsRequest) SetBillDate(v string) *V3BillsRequest {
	s.BillDate = &v
	return s
}

func (s *V3BillsRequest) SetHeader(v map[string]*string) *V3BillsRequest {
	s.Header = v
	return s
}

func (s *V3BillsRequest) SetAccessToken(v string) *V3BillsRequest {
	s.AccessToken = &v
	return s
}

func (s *V3BillsRequest) SetAppId(v string) *V3BillsRequest {
	s.AppId = &v
	return s
}

func (s *V3BillsRequest) SetMerchantId(v string) *V3BillsRequest {
	s.MerchantId = &v
	return s
}

type V3BillsResponse struct {
	LogId  *string              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *V3BillsResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s V3BillsResponse) String() string {
	return tea.Prettify(s)
}

func (s V3BillsResponse) GoString() string {
	return s.String()
}

func (s *V3BillsResponse) SetLogId(v string) *V3BillsResponse {
	s.LogId = &v
	return s
}

func (s *V3BillsResponse) SetData(v *V3BillsResponseData) *V3BillsResponse {
	s.Data = v
	return s
}

func (s *V3BillsResponse) SetErrNo(v int32) *V3BillsResponse {
	s.ErrNo = &v
	return s
}

func (s *V3BillsResponse) SetErrMsg(v string) *V3BillsResponse {
	s.ErrMsg = &v
	return s
}

type V3BillsResponseData struct {
	BillList []*string `json:"bill_list,omitempty" xml:"bill_list,omitempty" require:"true" type:"Repeated"`
}

func (s V3BillsResponseData) String() string {
	return tea.Prettify(s)
}

func (s V3BillsResponseData) GoString() string {
	return s.String()
}

func (s *V3BillsResponseData) SetBillList(v []*string) *V3BillsResponseData {
	s.BillList = v
	return s
}

type VerifyRecordQueryRequest struct {
	StartTime   *int64             `json:"start_time,omitempty" xml:"start_time,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	EndTime     *int64             `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PoiIds      []*string          `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" type:"Repeated"`
	Size        *int32             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
}

func (s VerifyRecordQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryRequest) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryRequest) SetStartTime(v int64) *VerifyRecordQueryRequest {
	s.StartTime = &v
	return s
}

func (s *VerifyRecordQueryRequest) SetAccountId(v string) *VerifyRecordQueryRequest {
	s.AccountId = &v
	return s
}

func (s *VerifyRecordQueryRequest) SetCursor(v string) *VerifyRecordQueryRequest {
	s.Cursor = &v
	return s
}

func (s *VerifyRecordQueryRequest) SetEndTime(v int64) *VerifyRecordQueryRequest {
	s.EndTime = &v
	return s
}

func (s *VerifyRecordQueryRequest) SetHeader(v map[string]*string) *VerifyRecordQueryRequest {
	s.Header = v
	return s
}

func (s *VerifyRecordQueryRequest) SetAccessToken(v string) *VerifyRecordQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *VerifyRecordQueryRequest) SetPoiIds(v []*string) *VerifyRecordQueryRequest {
	s.PoiIds = v
	return s
}

func (s *VerifyRecordQueryRequest) SetSize(v int32) *VerifyRecordQueryRequest {
	s.Size = &v
	return s
}

type VerifyRecordQueryResponse struct {
	Data  *VerifyRecordQueryResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *VerifyRecordQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s VerifyRecordQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponse) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponse) SetData(v *VerifyRecordQueryResponseData) *VerifyRecordQueryResponse {
	s.Data = v
	return s
}

func (s *VerifyRecordQueryResponse) SetExtra(v *VerifyRecordQueryResponseExtra) *VerifyRecordQueryResponse {
	s.Extra = v
	return s
}

type VerifyRecordQueryResponseData struct {
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Records       []*VerifyRecordQueryResponseDataRecordsItem   `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	RecordsV2     []*VerifyRecordQueryResponseDataRecordsV2Item `json:"records_v2,omitempty" xml:"records_v2,omitempty" type:"Repeated"`
	Total         *int64                                        `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s VerifyRecordQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseData) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseData) SetGwDescription(v string) *VerifyRecordQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *VerifyRecordQueryResponseData) SetRecords(v []*VerifyRecordQueryResponseDataRecordsItem) *VerifyRecordQueryResponseData {
	s.Records = v
	return s
}

func (s *VerifyRecordQueryResponseData) SetRecordsV2(v []*VerifyRecordQueryResponseDataRecordsV2Item) *VerifyRecordQueryResponseData {
	s.RecordsV2 = v
	return s
}

func (s *VerifyRecordQueryResponseData) SetTotal(v int64) *VerifyRecordQueryResponseData {
	s.Total = &v
	return s
}

func (s *VerifyRecordQueryResponseData) SetGwErrorCode(v int32) *VerifyRecordQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItem struct {
	Cursor            *string                                                   `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	Status            *int                                                      `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Amount            *VerifyRecordQueryResponseDataRecordsItemAmount           `json:"amount,omitempty" xml:"amount,omitempty"`
	CancelTime        *int64                                                    `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	VerifyAmountInfo  *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo `json:"verify_amount_info,omitempty" xml:"verify_amount_info,omitempty"`
	Sku               *VerifyRecordQueryResponseDataRecordsItemSku              `json:"sku,omitempty" xml:"sku,omitempty"`
	PeriodCard        *VerifyRecordQueryResponseDataRecordsItemPeriodCard       `json:"period_card,omitempty" xml:"period_card,omitempty"`
	CertificateStatus *int                                                      `json:"certificate_status,omitempty" xml:"certificate_status,omitempty"`
	CanCancel         *bool                                                     `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	VerifyTime        *int64                                                    `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	CertificateId     *string                                                   `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	VerifyId          *string                                                   `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyType        *int                                                      `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	Code              *string                                                   `json:"code,omitempty" xml:"code,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItem) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCursor(v string) *VerifyRecordQueryResponseDataRecordsItem {
	s.Cursor = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetStatus(v int) *VerifyRecordQueryResponseDataRecordsItem {
	s.Status = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetAmount(v *VerifyRecordQueryResponseDataRecordsItemAmount) *VerifyRecordQueryResponseDataRecordsItem {
	s.Amount = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCancelTime(v int64) *VerifyRecordQueryResponseDataRecordsItem {
	s.CancelTime = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetVerifyAmountInfo(v *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo) *VerifyRecordQueryResponseDataRecordsItem {
	s.VerifyAmountInfo = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetSku(v *VerifyRecordQueryResponseDataRecordsItemSku) *VerifyRecordQueryResponseDataRecordsItem {
	s.Sku = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetPeriodCard(v *VerifyRecordQueryResponseDataRecordsItemPeriodCard) *VerifyRecordQueryResponseDataRecordsItem {
	s.PeriodCard = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCertificateStatus(v int) *VerifyRecordQueryResponseDataRecordsItem {
	s.CertificateStatus = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCanCancel(v bool) *VerifyRecordQueryResponseDataRecordsItem {
	s.CanCancel = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetVerifyTime(v int64) *VerifyRecordQueryResponseDataRecordsItem {
	s.VerifyTime = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCertificateId(v string) *VerifyRecordQueryResponseDataRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetVerifyId(v string) *VerifyRecordQueryResponseDataRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetVerifyType(v int) *VerifyRecordQueryResponseDataRecordsItem {
	s.VerifyType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItem) SetCode(v string) *VerifyRecordQueryResponseDataRecordsItem {
	s.Code = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemAmount struct {
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetPlatformDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetOriginListMarketAmount(v int64) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetCouponPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetOriginalCurrency(v string) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetBrandTicketAmount(v int64) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetOriginalAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.OriginalAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetPaymentDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.PayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetListMarketAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemAmount) SetMerchantTicketAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemPeriodCard) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemPeriodCard) SetPeriodType(v int) *VerifyRecordQueryResponseDataRecordsItemPeriodCard {
	s.PeriodType = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemSku struct {
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemSku) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemSku) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetVoucherType(v int) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.VoucherType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetGrouponType(v int) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.GrouponType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetSkuId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.SkuId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetSoldStartTime(v int64) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetSkuOutId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.SkuOutId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetAccountId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.AccountId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetThirdSkuId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetTitle(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.Title = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetProductOutId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.ProductOutId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetMarketPrice(v int64) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.MarketPrice = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetProductId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.ProductId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemSku) SetSuplierProductOutId(v string) *VerifyRecordQueryResponseDataRecordsItemSku {
	s.SuplierProductOutId = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo struct {
	TimesCardSerialAmount *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount `json:"times_card_serial_amount,omitempty" xml:"times_card_serial_amount,omitempty"`
	TimeCardAmount        *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount        `json:"time_card_amount,omitempty" xml:"time_card_amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo) SetTimesCardSerialAmount(v *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo {
	s.TimesCardSerialAmount = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo) SetTimeCardAmount(v *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfo {
	s.TimeCardAmount = v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount) SetAmount(v int64) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimeCardAmount {
	s.Amount = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount struct {
	SerialNumb *int32                                                                               `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
	Amount     *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount) SetSerialNumb(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount {
	s.SerialNumb = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount) SetAmount(v *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmount {
	s.Amount = v
	return s
}

type VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount struct {
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetBrandTicketAmount(v int64) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPlatformDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetListMarketAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginListMarketAmount(v int64) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPaymentDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetCouponPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalCurrency(v string) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount) SetMerchantTicketAmount(v int32) *VerifyRecordQueryResponseDataRecordsItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.MerchantTicketAmount = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2Item struct {
	CanCancel         *bool                                                       `json:"can_cancel,omitempty" xml:"can_cancel,omitempty" require:"true"`
	Amount            *VerifyRecordQueryResponseDataRecordsV2ItemAmount           `json:"amount,omitempty" xml:"amount,omitempty"`
	VerifyType        *int                                                        `json:"verify_type,omitempty" xml:"verify_type,omitempty" require:"true"`
	VerifyAmountInfo  *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo `json:"verify_amount_info,omitempty" xml:"verify_amount_info,omitempty"`
	VerifyTime        *int64                                                      `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	Sku               *VerifyRecordQueryResponseDataRecordsV2ItemSku              `json:"sku,omitempty" xml:"sku,omitempty"`
	Cursor            *string                                                     `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	PeriodCard        *VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard       `json:"period_card,omitempty" xml:"period_card,omitempty"`
	VerifyId          *string                                                     `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	CertificateId     *string                                                     `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code              *string                                                     `json:"code,omitempty" xml:"code,omitempty"`
	CertificateStatus *int                                                        `json:"certificate_status,omitempty" xml:"certificate_status,omitempty"`
	Status            *int                                                        `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	CancelTime        *int64                                                      `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2Item) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2Item) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCanCancel(v bool) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.CanCancel = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetAmount(v *VerifyRecordQueryResponseDataRecordsV2ItemAmount) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.Amount = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetVerifyType(v int) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.VerifyType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetVerifyAmountInfo(v *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.VerifyAmountInfo = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetVerifyTime(v int64) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.VerifyTime = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetSku(v *VerifyRecordQueryResponseDataRecordsV2ItemSku) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.Sku = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCursor(v string) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.Cursor = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetPeriodCard(v *VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.PeriodCard = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetVerifyId(v string) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.VerifyId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCertificateId(v string) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.CertificateId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCode(v string) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.Code = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCertificateStatus(v int) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.CertificateStatus = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetStatus(v int) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.Status = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2Item) SetCancelTime(v int64) *VerifyRecordQueryResponseDataRecordsV2Item {
	s.CancelTime = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemAmount struct {
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.PayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetOriginalCurrency(v string) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetOriginListMarketAmount(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetMerchantTicketAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetPlatformDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetBrandTicketAmount(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetListMarketAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.ListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetCouponPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetPaymentDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemAmount) SetOriginalAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemAmount {
	s.OriginalAmount = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard struct {
	PeriodType *int `json:"period_type,omitempty" xml:"period_type,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard) SetPeriodType(v int) *VerifyRecordQueryResponseDataRecordsV2ItemPeriodCard {
	s.PeriodType = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemSku struct {
	SoldStartTime       *int64  `json:"sold_start_time,omitempty" xml:"sold_start_time,omitempty"`
	AccountId           *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	ProductOutId        *string `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	ThirdSkuId          *string `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	VoucherType         *int    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	MarketPrice         *int64  `json:"market_price,omitempty" xml:"market_price,omitempty"`
	SuplierProductOutId *string `json:"suplier_product_out_id,omitempty" xml:"suplier_product_out_id,omitempty"`
	SkuOutId            *string `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	GrouponType         *int    `json:"groupon_type,omitempty" xml:"groupon_type,omitempty"`
	ProductId           *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemSku) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemSku) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetSoldStartTime(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.SoldStartTime = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetAccountId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.AccountId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetTitle(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.Title = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetProductOutId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.ProductOutId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetThirdSkuId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.ThirdSkuId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetSkuId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.SkuId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetVoucherType(v int) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.VoucherType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetMarketPrice(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.MarketPrice = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetSuplierProductOutId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.SuplierProductOutId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetSkuOutId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.SkuOutId = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetGrouponType(v int) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.GrouponType = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemSku) SetProductId(v string) *VerifyRecordQueryResponseDataRecordsV2ItemSku {
	s.ProductId = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo struct {
	TimeCardAmount        *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount        `json:"time_card_amount,omitempty" xml:"time_card_amount,omitempty"`
	TimesCardSerialAmount *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount `json:"times_card_serial_amount,omitempty" xml:"times_card_serial_amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo) SetTimeCardAmount(v *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo {
	s.TimeCardAmount = v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo) SetTimesCardSerialAmount(v *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfo {
	s.TimesCardSerialAmount = v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount struct {
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount) SetAmount(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimeCardAmount {
	s.Amount = &v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount struct {
	SerialNumb *int32                                                                                 `json:"serial_numb,omitempty" xml:"serial_numb,omitempty" require:"true"`
	Amount     *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount) SetSerialNumb(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount {
	s.SerialNumb = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount) SetAmount(v *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmount {
	s.Amount = v
	return s
}

type VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount struct {
	MerchantTicketAmount   *int32  `json:"merchant_ticket_amount,omitempty" xml:"merchant_ticket_amount,omitempty"`
	CouponPayAmount        *int32  `json:"coupon_pay_amount,omitempty" xml:"coupon_pay_amount,omitempty"`
	PaymentDiscountAmount  *int32  `json:"payment_discount_amount,omitempty" xml:"payment_discount_amount,omitempty"`
	BrandTicketAmount      *int64  `json:"brand_ticket_amount,omitempty" xml:"brand_ticket_amount,omitempty"`
	PayAmount              *int32  `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	OriginalCurrency       *string `json:"original_currency,omitempty" xml:"original_currency,omitempty"`
	PlatformDiscountAmount *int32  `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	OriginalAmount         *int32  `json:"original_amount,omitempty" xml:"original_amount,omitempty" require:"true"`
	OriginListMarketAmount *int64  `json:"origin_list_market_amount,omitempty" xml:"origin_list_market_amount,omitempty"`
	ListMarketAmount       *int32  `json:"list_market_amount,omitempty" xml:"list_market_amount,omitempty"`
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetMerchantTicketAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.MerchantTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetCouponPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.CouponPayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPaymentDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PaymentDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetBrandTicketAmount(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.BrandTicketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPayAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PayAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalCurrency(v string) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalCurrency = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetPlatformDiscountAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginalAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginalAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetOriginListMarketAmount(v int64) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.OriginListMarketAmount = &v
	return s
}

func (s *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount) SetListMarketAmount(v int32) *VerifyRecordQueryResponseDataRecordsV2ItemVerifyAmountInfoTimesCardSerialAmountAmount {
	s.ListMarketAmount = &v
	return s
}

type VerifyRecordQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s VerifyRecordQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VerifyRecordQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *VerifyRecordQueryResponseExtra) SetNow(v int64) *VerifyRecordQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *VerifyRecordQueryResponseExtra) SetSubDescription(v string) *VerifyRecordQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VerifyRecordQueryResponseExtra) SetSubErrorCode(v int32) *VerifyRecordQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VerifyRecordQueryResponseExtra) SetDescription(v string) *VerifyRecordQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *VerifyRecordQueryResponseExtra) SetErrorCode(v int32) *VerifyRecordQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VerifyRecordQueryResponseExtra) SetLogid(v string) *VerifyRecordQueryResponseExtra {
	s.Logid = &v
	return s
}

type VideoBcQueryRequest struct {
	ItemIds     []*string          `json:"item_ids,omitempty" xml:"item_ids,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s VideoBcQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryRequest) GoString() string {
	return s.String()
}

func (s *VideoBcQueryRequest) SetItemIds(v []*string) *VideoBcQueryRequest {
	s.ItemIds = v
	return s
}

func (s *VideoBcQueryRequest) SetHeader(v map[string]*string) *VideoBcQueryRequest {
	s.Header = v
	return s
}

func (s *VideoBcQueryRequest) SetAccessToken(v string) *VideoBcQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoBcQueryRequest) SetOpenId(v string) *VideoBcQueryRequest {
	s.OpenId = &v
	return s
}

type VideoBcQueryResponse struct {
	ErrNo  *int32                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *VideoBcQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoBcQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponse) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponse) SetErrNo(v int32) *VideoBcQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoBcQueryResponse) SetErrMsg(v string) *VideoBcQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *VideoBcQueryResponse) SetLogId(v string) *VideoBcQueryResponse {
	s.LogId = &v
	return s
}

func (s *VideoBcQueryResponse) SetData(v *VideoBcQueryResponseData) *VideoBcQueryResponse {
	s.Data = v
	return s
}

type VideoBcQueryResponseData struct {
	Extra *VideoBcQueryResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *VideoBcQueryResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s VideoBcQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseData) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseData) SetExtra(v *VideoBcQueryResponseDataExtra) *VideoBcQueryResponseData {
	s.Extra = v
	return s
}

func (s *VideoBcQueryResponseData) SetData(v *VideoBcQueryResponseDataData) *VideoBcQueryResponseData {
	s.Data = v
	return s
}

type VideoBcQueryResponseDataData struct {
	List []*VideoBcQueryResponseDataDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s VideoBcQueryResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseDataData) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseDataData) SetList(v []*VideoBcQueryResponseDataDataListItem) *VideoBcQueryResponseDataData {
	s.List = v
	return s
}

type VideoBcQueryResponseDataDataListItem struct {
	ItemId      *string                                          `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	Cover       *string                                          `json:"cover,omitempty" xml:"cover,omitempty" require:"true"`
	IsReviewed  *bool                                            `json:"is_reviewed,omitempty" xml:"is_reviewed,omitempty" require:"true"`
	MediaType   *int                                             `json:"media_type,omitempty" xml:"media_type,omitempty"`
	ShareUrl    *string                                          `json:"share_url,omitempty" xml:"share_url,omitempty" require:"true"`
	IsTop       *bool                                            `json:"is_top,omitempty" xml:"is_top,omitempty" require:"true"`
	VideoStatus *int32                                           `json:"video_status,omitempty" xml:"video_status,omitempty" require:"true"`
	Statistics  *VideoBcQueryResponseDataDataListItemStatistics  `json:"statistics,omitempty" xml:"statistics,omitempty"`
	CreateTime  *int64                                           `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	VideoAnchor *VideoBcQueryResponseDataDataListItemVideoAnchor `json:"video_anchor,omitempty" xml:"video_anchor,omitempty"`
	Title       *string                                          `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	VideoId     *string                                          `json:"video_id,omitempty" xml:"video_id,omitempty"`
}

func (s VideoBcQueryResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseDataDataListItem) SetItemId(v string) *VideoBcQueryResponseDataDataListItem {
	s.ItemId = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetCover(v string) *VideoBcQueryResponseDataDataListItem {
	s.Cover = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetIsReviewed(v bool) *VideoBcQueryResponseDataDataListItem {
	s.IsReviewed = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetMediaType(v int) *VideoBcQueryResponseDataDataListItem {
	s.MediaType = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetShareUrl(v string) *VideoBcQueryResponseDataDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetIsTop(v bool) *VideoBcQueryResponseDataDataListItem {
	s.IsTop = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetVideoStatus(v int32) *VideoBcQueryResponseDataDataListItem {
	s.VideoStatus = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetStatistics(v *VideoBcQueryResponseDataDataListItemStatistics) *VideoBcQueryResponseDataDataListItem {
	s.Statistics = v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetCreateTime(v int64) *VideoBcQueryResponseDataDataListItem {
	s.CreateTime = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetVideoAnchor(v *VideoBcQueryResponseDataDataListItemVideoAnchor) *VideoBcQueryResponseDataDataListItem {
	s.VideoAnchor = v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetTitle(v string) *VideoBcQueryResponseDataDataListItem {
	s.Title = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItem) SetVideoId(v string) *VideoBcQueryResponseDataDataListItem {
	s.VideoId = &v
	return s
}

type VideoBcQueryResponseDataDataListItemStatistics struct {
	CommentCount  *int32 `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
	DiggCount     *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	DownloadCount *int32 `json:"download_count,omitempty" xml:"download_count,omitempty" require:"true"`
	PlayCount     *int32 `json:"play_count,omitempty" xml:"play_count,omitempty" require:"true"`
	ShareCount    *int32 `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	ForwardCount  *int32 `json:"forward_count,omitempty" xml:"forward_count,omitempty" require:"true"`
}

func (s VideoBcQueryResponseDataDataListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseDataDataListItemStatistics) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetCommentCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.CommentCount = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetDiggCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.DiggCount = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetDownloadCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.DownloadCount = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetPlayCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.PlayCount = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetShareCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.ShareCount = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemStatistics) SetForwardCount(v int32) *VideoBcQueryResponseDataDataListItemStatistics {
	s.ForwardCount = &v
	return s
}

type VideoBcQueryResponseDataDataListItemVideoAnchor struct {
	AnchorId   *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	AnchorType *int    `json:"anchor_type,omitempty" xml:"anchor_type,omitempty"`
}

func (s VideoBcQueryResponseDataDataListItemVideoAnchor) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseDataDataListItemVideoAnchor) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseDataDataListItemVideoAnchor) SetAnchorId(v string) *VideoBcQueryResponseDataDataListItemVideoAnchor {
	s.AnchorId = &v
	return s
}

func (s *VideoBcQueryResponseDataDataListItemVideoAnchor) SetAnchorType(v int) *VideoBcQueryResponseDataDataListItemVideoAnchor {
	s.AnchorType = &v
	return s
}

type VideoBcQueryResponseDataExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s VideoBcQueryResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoBcQueryResponseDataExtra) GoString() string {
	return s.String()
}

func (s *VideoBcQueryResponseDataExtra) SetSubDescription(v string) *VideoBcQueryResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoBcQueryResponseDataExtra) SetLogid(v string) *VideoBcQueryResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *VideoBcQueryResponseDataExtra) SetNow(v int64) *VideoBcQueryResponseDataExtra {
	s.Now = &v
	return s
}

func (s *VideoBcQueryResponseDataExtra) SetErrorCode(v int32) *VideoBcQueryResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoBcQueryResponseDataExtra) SetDescription(v string) *VideoBcQueryResponseDataExtra {
	s.Description = &v
	return s
}

func (s *VideoBcQueryResponseDataExtra) SetSubErrorCode(v int32) *VideoBcQueryResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

type VideoCreateImageTextRequest struct {
	PoiId          *string            `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	MicroAppTitle  *string            `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty"`
	AtUsers        []*string          `json:"at_users,omitempty" xml:"at_users,omitempty" type:"Repeated"`
	MusicId        *int64             `json:"music_id,omitempty" xml:"music_id,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	MicroAppUrl    *string            `json:"micro_app_url,omitempty" xml:"micro_app_url,omitempty"`
	DownloadType   *int32             `json:"download_type,omitempty" xml:"download_type,omitempty"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	PrivateStatus  *int32             `json:"private_status,omitempty" xml:"private_status,omitempty"`
	AgentClientKey *string            `json:"agent_client_key,omitempty" xml:"agent_client_key,omitempty"`
	ImageList      []*string          `json:"image_list,omitempty" xml:"image_list,omitempty" require:"true" type:"Repeated"`
	TaskId         *int64             `json:"task_id,omitempty" xml:"task_id,omitempty"`
	Text           *string            `json:"text,omitempty" xml:"text,omitempty"`
	OpenId         *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	PoiCommerce    *bool              `json:"poi_commerce,omitempty" xml:"poi_commerce,omitempty"`
	MicroAppId     *string            `json:"micro_app_id,omitempty" xml:"micro_app_id,omitempty"`
}

func (s VideoCreateImageTextRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateImageTextRequest) GoString() string {
	return s.String()
}

func (s *VideoCreateImageTextRequest) SetPoiId(v string) *VideoCreateImageTextRequest {
	s.PoiId = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetMicroAppTitle(v string) *VideoCreateImageTextRequest {
	s.MicroAppTitle = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetAtUsers(v []*string) *VideoCreateImageTextRequest {
	s.AtUsers = v
	return s
}

func (s *VideoCreateImageTextRequest) SetMusicId(v int64) *VideoCreateImageTextRequest {
	s.MusicId = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetAccessToken(v string) *VideoCreateImageTextRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetMicroAppUrl(v string) *VideoCreateImageTextRequest {
	s.MicroAppUrl = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetDownloadType(v int32) *VideoCreateImageTextRequest {
	s.DownloadType = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetHeader(v map[string]*string) *VideoCreateImageTextRequest {
	s.Header = v
	return s
}

func (s *VideoCreateImageTextRequest) SetPrivateStatus(v int32) *VideoCreateImageTextRequest {
	s.PrivateStatus = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetAgentClientKey(v string) *VideoCreateImageTextRequest {
	s.AgentClientKey = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetImageList(v []*string) *VideoCreateImageTextRequest {
	s.ImageList = v
	return s
}

func (s *VideoCreateImageTextRequest) SetTaskId(v int64) *VideoCreateImageTextRequest {
	s.TaskId = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetText(v string) *VideoCreateImageTextRequest {
	s.Text = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetOpenId(v string) *VideoCreateImageTextRequest {
	s.OpenId = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetPoiCommerce(v bool) *VideoCreateImageTextRequest {
	s.PoiCommerce = &v
	return s
}

func (s *VideoCreateImageTextRequest) SetMicroAppId(v string) *VideoCreateImageTextRequest {
	s.MicroAppId = &v
	return s
}

type VideoCreateImageTextResponse struct {
	Extra *VideoCreateImageTextResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *VideoCreateImageTextResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s VideoCreateImageTextResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateImageTextResponse) GoString() string {
	return s.String()
}

func (s *VideoCreateImageTextResponse) SetExtra(v *VideoCreateImageTextResponseExtra) *VideoCreateImageTextResponse {
	s.Extra = v
	return s
}

func (s *VideoCreateImageTextResponse) SetData(v *VideoCreateImageTextResponseData) *VideoCreateImageTextResponse {
	s.Data = v
	return s
}

type VideoCreateImageTextResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	VideoId       *string `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	ItemId        *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s VideoCreateImageTextResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateImageTextResponseData) GoString() string {
	return s.String()
}

func (s *VideoCreateImageTextResponseData) SetGwDescription(v string) *VideoCreateImageTextResponseData {
	s.GwDescription = &v
	return s
}

func (s *VideoCreateImageTextResponseData) SetVideoId(v string) *VideoCreateImageTextResponseData {
	s.VideoId = &v
	return s
}

func (s *VideoCreateImageTextResponseData) SetItemId(v string) *VideoCreateImageTextResponseData {
	s.ItemId = &v
	return s
}

func (s *VideoCreateImageTextResponseData) SetGwErrorCode(v int32) *VideoCreateImageTextResponseData {
	s.GwErrorCode = &v
	return s
}

type VideoCreateImageTextResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s VideoCreateImageTextResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateImageTextResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoCreateImageTextResponseExtra) SetNow(v int64) *VideoCreateImageTextResponseExtra {
	s.Now = &v
	return s
}

func (s *VideoCreateImageTextResponseExtra) SetSubDescription(v string) *VideoCreateImageTextResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoCreateImageTextResponseExtra) SetSubErrorCode(v int32) *VideoCreateImageTextResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoCreateImageTextResponseExtra) SetDescription(v string) *VideoCreateImageTextResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoCreateImageTextResponseExtra) SetErrorCode(v int32) *VideoCreateImageTextResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoCreateImageTextResponseExtra) SetLogid(v string) *VideoCreateImageTextResponseExtra {
	s.Logid = &v
	return s
}

type VideoCreateVideoRequest struct {
	OpenId                 *string                                        `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Text                   *string                                        `json:"text,omitempty" xml:"text,omitempty"`
	CustomCoverImageUrl    *string                                        `json:"custom_cover_image_url,omitempty" xml:"custom_cover_image_url,omitempty"`
	MicroAppTitle          *string                                        `json:"micro_app_title,omitempty" xml:"micro_app_title,omitempty"`
	PoiCommerce            *bool                                          `json:"poi_commerce,omitempty" xml:"poi_commerce,omitempty"`
	VideoId                *string                                        `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	VrTranscodeExtraParams *VideoCreateVideoRequestVrTranscodeExtraParams `json:"vr_transcode_extra_params,omitempty" xml:"vr_transcode_extra_params,omitempty"`
	MicroAppUrl            *string                                        `json:"micro_app_url,omitempty" xml:"micro_app_url,omitempty"`
	DownloadType           *int32                                         `json:"download_type,omitempty" xml:"download_type,omitempty"`
	Header                 map[string]*string                             `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken            *string                                        `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PrivateStatus          *int32                                         `json:"private_status,omitempty" xml:"private_status,omitempty"`
	PoiId                  *string                                        `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	AtUsers                []*string                                      `json:"at_users,omitempty" xml:"at_users,omitempty" type:"Repeated"`
	AgentClientKey         *string                                        `json:"agent_client_key,omitempty" xml:"agent_client_key,omitempty"`
	CoverTsp               *float64                                       `json:"cover_tsp,omitempty" xml:"cover_tsp,omitempty"`
	MicroAppId             *string                                        `json:"micro_app_id,omitempty" xml:"micro_app_id,omitempty"`
	TaskId                 *int64                                         `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s VideoCreateVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoRequest) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoRequest) SetOpenId(v string) *VideoCreateVideoRequest {
	s.OpenId = &v
	return s
}

func (s *VideoCreateVideoRequest) SetText(v string) *VideoCreateVideoRequest {
	s.Text = &v
	return s
}

func (s *VideoCreateVideoRequest) SetCustomCoverImageUrl(v string) *VideoCreateVideoRequest {
	s.CustomCoverImageUrl = &v
	return s
}

func (s *VideoCreateVideoRequest) SetMicroAppTitle(v string) *VideoCreateVideoRequest {
	s.MicroAppTitle = &v
	return s
}

func (s *VideoCreateVideoRequest) SetPoiCommerce(v bool) *VideoCreateVideoRequest {
	s.PoiCommerce = &v
	return s
}

func (s *VideoCreateVideoRequest) SetVideoId(v string) *VideoCreateVideoRequest {
	s.VideoId = &v
	return s
}

func (s *VideoCreateVideoRequest) SetVrTranscodeExtraParams(v *VideoCreateVideoRequestVrTranscodeExtraParams) *VideoCreateVideoRequest {
	s.VrTranscodeExtraParams = v
	return s
}

func (s *VideoCreateVideoRequest) SetMicroAppUrl(v string) *VideoCreateVideoRequest {
	s.MicroAppUrl = &v
	return s
}

func (s *VideoCreateVideoRequest) SetDownloadType(v int32) *VideoCreateVideoRequest {
	s.DownloadType = &v
	return s
}

func (s *VideoCreateVideoRequest) SetHeader(v map[string]*string) *VideoCreateVideoRequest {
	s.Header = v
	return s
}

func (s *VideoCreateVideoRequest) SetAccessToken(v string) *VideoCreateVideoRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoCreateVideoRequest) SetPrivateStatus(v int32) *VideoCreateVideoRequest {
	s.PrivateStatus = &v
	return s
}

func (s *VideoCreateVideoRequest) SetPoiId(v string) *VideoCreateVideoRequest {
	s.PoiId = &v
	return s
}

func (s *VideoCreateVideoRequest) SetAtUsers(v []*string) *VideoCreateVideoRequest {
	s.AtUsers = v
	return s
}

func (s *VideoCreateVideoRequest) SetAgentClientKey(v string) *VideoCreateVideoRequest {
	s.AgentClientKey = &v
	return s
}

func (s *VideoCreateVideoRequest) SetCoverTsp(v float64) *VideoCreateVideoRequest {
	s.CoverTsp = &v
	return s
}

func (s *VideoCreateVideoRequest) SetMicroAppId(v string) *VideoCreateVideoRequest {
	s.MicroAppId = &v
	return s
}

func (s *VideoCreateVideoRequest) SetTaskId(v int64) *VideoCreateVideoRequest {
	s.TaskId = &v
	return s
}

type VideoCreateVideoRequestVrTranscodeExtraParams struct {
	VideoStyle *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle `json:"video_style,omitempty" xml:"video_style,omitempty" require:"true"`
}

func (s VideoCreateVideoRequestVrTranscodeExtraParams) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoRequestVrTranscodeExtraParams) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoRequestVrTranscodeExtraParams) SetVideoStyle(v *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) *VideoCreateVideoRequestVrTranscodeExtraParams {
	s.VideoStyle = v
	return s
}

type VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle struct {
	Vstyle          *int `json:"vstyle,omitempty" xml:"vstyle,omitempty" require:"true"`
	Dimention       *int `json:"dimention,omitempty" xml:"dimention,omitempty" require:"true"`
	ProjectionModel *int `json:"projection_model,omitempty" xml:"projection_model,omitempty" require:"true"`
	ViewSize        *int `json:"view_size,omitempty" xml:"view_size,omitempty" require:"true"`
}

func (s VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) SetVstyle(v int) *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle {
	s.Vstyle = &v
	return s
}

func (s *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) SetDimention(v int) *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle {
	s.Dimention = &v
	return s
}

func (s *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) SetProjectionModel(v int) *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle {
	s.ProjectionModel = &v
	return s
}

func (s *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle) SetViewSize(v int) *VideoCreateVideoRequestVrTranscodeExtraParamsVideoStyle {
	s.ViewSize = &v
	return s
}

type VideoCreateVideoResponse struct {
	Data  *VideoCreateVideoResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *VideoCreateVideoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s VideoCreateVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoResponse) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoResponse) SetData(v *VideoCreateVideoResponseData) *VideoCreateVideoResponse {
	s.Data = v
	return s
}

func (s *VideoCreateVideoResponse) SetExtra(v *VideoCreateVideoResponseExtra) *VideoCreateVideoResponse {
	s.Extra = v
	return s
}

type VideoCreateVideoResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ItemId        *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	VideoId       *string `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
}

func (s VideoCreateVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoResponseData) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoResponseData) SetGwErrorCode(v int32) *VideoCreateVideoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *VideoCreateVideoResponseData) SetGwDescription(v string) *VideoCreateVideoResponseData {
	s.GwDescription = &v
	return s
}

func (s *VideoCreateVideoResponseData) SetItemId(v string) *VideoCreateVideoResponseData {
	s.ItemId = &v
	return s
}

func (s *VideoCreateVideoResponseData) SetVideoId(v string) *VideoCreateVideoResponseData {
	s.VideoId = &v
	return s
}

type VideoCreateVideoResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s VideoCreateVideoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoCreateVideoResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoCreateVideoResponseExtra) SetLogid(v string) *VideoCreateVideoResponseExtra {
	s.Logid = &v
	return s
}

func (s *VideoCreateVideoResponseExtra) SetNow(v int64) *VideoCreateVideoResponseExtra {
	s.Now = &v
	return s
}

func (s *VideoCreateVideoResponseExtra) SetSubDescription(v string) *VideoCreateVideoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoCreateVideoResponseExtra) SetSubErrorCode(v int32) *VideoCreateVideoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoCreateVideoResponseExtra) SetDescription(v string) *VideoCreateVideoResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoCreateVideoResponseExtra) SetErrorCode(v int32) *VideoCreateVideoResponseExtra {
	s.ErrorCode = &v
	return s
}

type VideoIdToOpenItemIdRequest struct {
	AccessKey   *string            `json:"access_key,omitempty" xml:"access_key,omitempty" require:"true"`
	VideoIds    []*string          `json:"video_ids,omitempty" xml:"video_ids,omitempty" type:"Repeated"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoIdToOpenItemIdRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoIdToOpenItemIdRequest) GoString() string {
	return s.String()
}

func (s *VideoIdToOpenItemIdRequest) SetAccessKey(v string) *VideoIdToOpenItemIdRequest {
	s.AccessKey = &v
	return s
}

func (s *VideoIdToOpenItemIdRequest) SetVideoIds(v []*string) *VideoIdToOpenItemIdRequest {
	s.VideoIds = v
	return s
}

func (s *VideoIdToOpenItemIdRequest) SetAppId(v string) *VideoIdToOpenItemIdRequest {
	s.AppId = &v
	return s
}

func (s *VideoIdToOpenItemIdRequest) SetHeader(v map[string]*string) *VideoIdToOpenItemIdRequest {
	s.Header = v
	return s
}

func (s *VideoIdToOpenItemIdRequest) SetAccessToken(v string) *VideoIdToOpenItemIdRequest {
	s.AccessToken = &v
	return s
}

type VideoIdToOpenItemIdResponse struct {
	Data   *VideoIdToOpenItemIdResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s VideoIdToOpenItemIdResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoIdToOpenItemIdResponse) GoString() string {
	return s.String()
}

func (s *VideoIdToOpenItemIdResponse) SetData(v *VideoIdToOpenItemIdResponseData) *VideoIdToOpenItemIdResponse {
	s.Data = v
	return s
}

func (s *VideoIdToOpenItemIdResponse) SetErrNo(v int32) *VideoIdToOpenItemIdResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoIdToOpenItemIdResponse) SetErrMsg(v string) *VideoIdToOpenItemIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *VideoIdToOpenItemIdResponse) SetLogId(v string) *VideoIdToOpenItemIdResponse {
	s.LogId = &v
	return s
}

type VideoIdToOpenItemIdResponseData struct {
	ConvertResult map[string]*string `json:"convert_result,omitempty" xml:"convert_result,omitempty" require:"true"`
}

func (s VideoIdToOpenItemIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoIdToOpenItemIdResponseData) GoString() string {
	return s.String()
}

func (s *VideoIdToOpenItemIdResponseData) SetConvertResult(v map[string]*string) *VideoIdToOpenItemIdResponseData {
	s.ConvertResult = v
	return s
}

type VideoMountApplyPermissionRequest struct {
	AnchorText     *string            `json:"anchor_text,omitempty" xml:"anchor_text,omitempty" require:"true"`
	ImagePath      *string            `json:"image_path,omitempty" xml:"image_path,omitempty" require:"true"`
	Header         map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ComponentAppid *string            `json:"component_appid,omitempty" xml:"component_appid,omitempty" require:"true"`
	Intro          *string            `json:"intro,omitempty" xml:"intro,omitempty" require:"true"`
}

func (s VideoMountApplyPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoMountApplyPermissionRequest) GoString() string {
	return s.String()
}

func (s *VideoMountApplyPermissionRequest) SetAnchorText(v string) *VideoMountApplyPermissionRequest {
	s.AnchorText = &v
	return s
}

func (s *VideoMountApplyPermissionRequest) SetImagePath(v string) *VideoMountApplyPermissionRequest {
	s.ImagePath = &v
	return s
}

func (s *VideoMountApplyPermissionRequest) SetHeader(v map[string]*string) *VideoMountApplyPermissionRequest {
	s.Header = v
	return s
}

func (s *VideoMountApplyPermissionRequest) SetAccessToken(v string) *VideoMountApplyPermissionRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoMountApplyPermissionRequest) SetComponentAppid(v string) *VideoMountApplyPermissionRequest {
	s.ComponentAppid = &v
	return s
}

func (s *VideoMountApplyPermissionRequest) SetIntro(v string) *VideoMountApplyPermissionRequest {
	s.Intro = &v
	return s
}

type VideoMountApplyPermissionResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s VideoMountApplyPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoMountApplyPermissionResponse) GoString() string {
	return s.String()
}

func (s *VideoMountApplyPermissionResponse) SetLogId(v string) *VideoMountApplyPermissionResponse {
	s.LogId = &v
	return s
}

func (s *VideoMountApplyPermissionResponse) SetErrNo(v int32) *VideoMountApplyPermissionResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoMountApplyPermissionResponse) SetErrMsg(v string) *VideoMountApplyPermissionResponse {
	s.ErrMsg = &v
	return s
}

type VideoMountModifyConfigRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ConfigKey   *string            `json:"config_key,omitempty" xml:"config_key,omitempty" require:"true"`
	ConfigValue *string            `json:"config_value,omitempty" xml:"config_value,omitempty" require:"true"`
}

func (s VideoMountModifyConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoMountModifyConfigRequest) GoString() string {
	return s.String()
}

func (s *VideoMountModifyConfigRequest) SetHeader(v map[string]*string) *VideoMountModifyConfigRequest {
	s.Header = v
	return s
}

func (s *VideoMountModifyConfigRequest) SetAccessToken(v string) *VideoMountModifyConfigRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoMountModifyConfigRequest) SetConfigKey(v string) *VideoMountModifyConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *VideoMountModifyConfigRequest) SetConfigValue(v string) *VideoMountModifyConfigRequest {
	s.ConfigValue = &v
	return s
}

type VideoMountModifyConfigResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s VideoMountModifyConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoMountModifyConfigResponse) GoString() string {
	return s.String()
}

func (s *VideoMountModifyConfigResponse) SetLogId(v string) *VideoMountModifyConfigResponse {
	s.LogId = &v
	return s
}

func (s *VideoMountModifyConfigResponse) SetErrNo(v int32) *VideoMountModifyConfigResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoMountModifyConfigResponse) SetErrMsg(v string) *VideoMountModifyConfigResponse {
	s.ErrMsg = &v
	return s
}

type VideoMountQueryConfigHistoryRequest struct {
	PageSize    *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	PageNum     *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoMountQueryConfigHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigHistoryRequest) SetPageSize(v int32) *VideoMountQueryConfigHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *VideoMountQueryConfigHistoryRequest) SetPageNum(v int32) *VideoMountQueryConfigHistoryRequest {
	s.PageNum = &v
	return s
}

func (s *VideoMountQueryConfigHistoryRequest) SetHeader(v map[string]*string) *VideoMountQueryConfigHistoryRequest {
	s.Header = v
	return s
}

func (s *VideoMountQueryConfigHistoryRequest) SetAccessToken(v string) *VideoMountQueryConfigHistoryRequest {
	s.AccessToken = &v
	return s
}

type VideoMountQueryConfigHistoryResponse struct {
	LogId  *string                                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *VideoMountQueryConfigHistoryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s VideoMountQueryConfigHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigHistoryResponse) SetLogId(v string) *VideoMountQueryConfigHistoryResponse {
	s.LogId = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponse) SetData(v *VideoMountQueryConfigHistoryResponseData) *VideoMountQueryConfigHistoryResponse {
	s.Data = v
	return s
}

func (s *VideoMountQueryConfigHistoryResponse) SetErrNo(v int32) *VideoMountQueryConfigHistoryResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponse) SetErrMsg(v string) *VideoMountQueryConfigHistoryResponse {
	s.ErrMsg = &v
	return s
}

type VideoMountQueryConfigHistoryResponseData struct {
	Total             *int64                                                           `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	ChangeHistoryList []*VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem `json:"change_history_list,omitempty" xml:"change_history_list,omitempty" require:"true" type:"Repeated"`
}

func (s VideoMountQueryConfigHistoryResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigHistoryResponseData) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigHistoryResponseData) SetTotal(v int64) *VideoMountQueryConfigHistoryResponseData {
	s.Total = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseData) SetChangeHistoryList(v []*VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) *VideoMountQueryConfigHistoryResponseData {
	s.ChangeHistoryList = v
	return s
}

type VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem struct {
	ReasonList []*string `json:"reason_list,omitempty" xml:"reason_list,omitempty" require:"true" type:"Repeated"`
	ChangeTime *int32    `json:"change_time,omitempty" xml:"change_time,omitempty" require:"true"`
	ConfigKey  *string   `json:"config_key,omitempty" xml:"config_key,omitempty" require:"true"`
	OldValue   *string   `json:"old_value,omitempty" xml:"old_value,omitempty" require:"true"`
	NewValue   *string   `json:"new_value,omitempty" xml:"new_value,omitempty" require:"true"`
	Status     *int32    `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetReasonList(v []*string) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.ReasonList = v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetChangeTime(v int32) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.ChangeTime = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetConfigKey(v string) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.ConfigKey = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetOldValue(v string) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.OldValue = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetNewValue(v string) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.NewValue = &v
	return s
}

func (s *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem) SetStatus(v int32) *VideoMountQueryConfigHistoryResponseDataChangeHistoryListItem {
	s.Status = &v
	return s
}

type VideoMountQueryConfigRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoMountQueryConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigRequest) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigRequest) SetHeader(v map[string]*string) *VideoMountQueryConfigRequest {
	s.Header = v
	return s
}

func (s *VideoMountQueryConfigRequest) SetAccessToken(v string) *VideoMountQueryConfigRequest {
	s.AccessToken = &v
	return s
}

type VideoMountQueryConfigResponse struct {
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *VideoMountQueryConfigResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s VideoMountQueryConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigResponse) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigResponse) SetLogId(v string) *VideoMountQueryConfigResponse {
	s.LogId = &v
	return s
}

func (s *VideoMountQueryConfigResponse) SetData(v *VideoMountQueryConfigResponseData) *VideoMountQueryConfigResponse {
	s.Data = v
	return s
}

func (s *VideoMountQueryConfigResponse) SetErrNo(v int32) *VideoMountQueryConfigResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoMountQueryConfigResponse) SetErrMsg(v string) *VideoMountQueryConfigResponse {
	s.ErrMsg = &v
	return s
}

type VideoMountQueryConfigResponseData struct {
	Title    *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	IconType *string `json:"icon_type,omitempty" xml:"icon_type,omitempty" require:"true"`
	DescType *string `json:"desc_type,omitempty" xml:"desc_type,omitempty" require:"true"`
}

func (s VideoMountQueryConfigResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoMountQueryConfigResponseData) GoString() string {
	return s.String()
}

func (s *VideoMountQueryConfigResponseData) SetTitle(v string) *VideoMountQueryConfigResponseData {
	s.Title = &v
	return s
}

func (s *VideoMountQueryConfigResponseData) SetIconType(v string) *VideoMountQueryConfigResponseData {
	s.IconType = &v
	return s
}

func (s *VideoMountQueryConfigResponseData) SetDescType(v string) *VideoMountQueryConfigResponseData {
	s.DescType = &v
	return s
}

type VideoQueryRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemIds     []*string          `json:"item_ids,omitempty" xml:"item_ids,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryRequest) GoString() string {
	return s.String()
}

func (s *VideoQueryRequest) SetOpenId(v string) *VideoQueryRequest {
	s.OpenId = &v
	return s
}

func (s *VideoQueryRequest) SetItemIds(v []*string) *VideoQueryRequest {
	s.ItemIds = v
	return s
}

func (s *VideoQueryRequest) SetHeader(v map[string]*string) *VideoQueryRequest {
	s.Header = v
	return s
}

func (s *VideoQueryRequest) SetAccessToken(v string) *VideoQueryRequest {
	s.AccessToken = &v
	return s
}

type VideoQueryResponse struct {
	ErrNo  *int32                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *VideoQueryResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponse) GoString() string {
	return s.String()
}

func (s *VideoQueryResponse) SetErrNo(v int32) *VideoQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *VideoQueryResponse) SetErrMsg(v string) *VideoQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *VideoQueryResponse) SetLogId(v string) *VideoQueryResponse {
	s.LogId = &v
	return s
}

func (s *VideoQueryResponse) SetData(v *VideoQueryResponseData) *VideoQueryResponse {
	s.Data = v
	return s
}

type VideoQueryResponseData struct {
	Data  *VideoQueryResponseDataData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *VideoQueryResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s VideoQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseData) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseData) SetData(v *VideoQueryResponseDataData) *VideoQueryResponseData {
	s.Data = v
	return s
}

func (s *VideoQueryResponseData) SetExtra(v *VideoQueryResponseDataExtra) *VideoQueryResponseData {
	s.Extra = v
	return s
}

type VideoQueryResponseDataData struct {
	List []*VideoQueryResponseDataDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
}

func (s VideoQueryResponseDataData) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseDataData) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseDataData) SetList(v []*VideoQueryResponseDataDataListItem) *VideoQueryResponseDataData {
	s.List = v
	return s
}

type VideoQueryResponseDataDataListItem struct {
	ItemId      *string                                        `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	IsTop       *bool                                          `json:"is_top,omitempty" xml:"is_top,omitempty" require:"true"`
	IsReviewed  *bool                                          `json:"is_reviewed,omitempty" xml:"is_reviewed,omitempty" require:"true"`
	MediaType   *int                                           `json:"media_type,omitempty" xml:"media_type,omitempty"`
	Cover       *string                                        `json:"cover,omitempty" xml:"cover,omitempty" require:"true"`
	Statistics  *VideoQueryResponseDataDataListItemStatistics  `json:"statistics,omitempty" xml:"statistics,omitempty"`
	VideoId     *string                                        `json:"video_id,omitempty" xml:"video_id,omitempty"`
	CreateTime  *int64                                         `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	Title       *string                                        `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	VideoStatus *int32                                         `json:"video_status,omitempty" xml:"video_status,omitempty" require:"true"`
	ShareUrl    *string                                        `json:"share_url,omitempty" xml:"share_url,omitempty" require:"true"`
	VideoAnchor *VideoQueryResponseDataDataListItemVideoAnchor `json:"video_anchor,omitempty" xml:"video_anchor,omitempty"`
}

func (s VideoQueryResponseDataDataListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseDataDataListItem) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseDataDataListItem) SetItemId(v string) *VideoQueryResponseDataDataListItem {
	s.ItemId = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetIsTop(v bool) *VideoQueryResponseDataDataListItem {
	s.IsTop = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetIsReviewed(v bool) *VideoQueryResponseDataDataListItem {
	s.IsReviewed = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetMediaType(v int) *VideoQueryResponseDataDataListItem {
	s.MediaType = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetCover(v string) *VideoQueryResponseDataDataListItem {
	s.Cover = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetStatistics(v *VideoQueryResponseDataDataListItemStatistics) *VideoQueryResponseDataDataListItem {
	s.Statistics = v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetVideoId(v string) *VideoQueryResponseDataDataListItem {
	s.VideoId = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetCreateTime(v int64) *VideoQueryResponseDataDataListItem {
	s.CreateTime = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetTitle(v string) *VideoQueryResponseDataDataListItem {
	s.Title = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetVideoStatus(v int32) *VideoQueryResponseDataDataListItem {
	s.VideoStatus = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetShareUrl(v string) *VideoQueryResponseDataDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *VideoQueryResponseDataDataListItem) SetVideoAnchor(v *VideoQueryResponseDataDataListItemVideoAnchor) *VideoQueryResponseDataDataListItem {
	s.VideoAnchor = v
	return s
}

type VideoQueryResponseDataDataListItemStatistics struct {
	PlayCount     *int32 `json:"play_count,omitempty" xml:"play_count,omitempty" require:"true"`
	ShareCount    *int32 `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	ForwardCount  *int32 `json:"forward_count,omitempty" xml:"forward_count,omitempty" require:"true"`
	CommentCount  *int32 `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
	DiggCount     *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	DownloadCount *int32 `json:"download_count,omitempty" xml:"download_count,omitempty" require:"true"`
}

func (s VideoQueryResponseDataDataListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseDataDataListItemStatistics) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetPlayCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.PlayCount = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetShareCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.ShareCount = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetForwardCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.ForwardCount = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetCommentCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.CommentCount = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetDiggCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.DiggCount = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemStatistics) SetDownloadCount(v int32) *VideoQueryResponseDataDataListItemStatistics {
	s.DownloadCount = &v
	return s
}

type VideoQueryResponseDataDataListItemVideoAnchor struct {
	AnchorType *int    `json:"anchor_type,omitempty" xml:"anchor_type,omitempty"`
	AnchorId   *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s VideoQueryResponseDataDataListItemVideoAnchor) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseDataDataListItemVideoAnchor) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseDataDataListItemVideoAnchor) SetAnchorType(v int) *VideoQueryResponseDataDataListItemVideoAnchor {
	s.AnchorType = &v
	return s
}

func (s *VideoQueryResponseDataDataListItemVideoAnchor) SetAnchorId(v string) *VideoQueryResponseDataDataListItemVideoAnchor {
	s.AnchorId = &v
	return s
}

type VideoQueryResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s VideoQueryResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoQueryResponseDataExtra) GoString() string {
	return s.String()
}

func (s *VideoQueryResponseDataExtra) SetNow(v int64) *VideoQueryResponseDataExtra {
	s.Now = &v
	return s
}

func (s *VideoQueryResponseDataExtra) SetErrorCode(v int32) *VideoQueryResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoQueryResponseDataExtra) SetDescription(v string) *VideoQueryResponseDataExtra {
	s.Description = &v
	return s
}

func (s *VideoQueryResponseDataExtra) SetSubErrorCode(v int32) *VideoQueryResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoQueryResponseDataExtra) SetSubDescription(v string) *VideoQueryResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoQueryResponseDataExtra) SetLogid(v string) *VideoQueryResponseDataExtra {
	s.Logid = &v
	return s
}

type VideoUploadVideoPartRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UploadId    *string            `json:"upload_id,omitempty" xml:"upload_id,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	PartNumber  *int64             `json:"part_number,omitempty" xml:"part_number,omitempty" require:"true"`
	Video       *util.FileField    `json:"video,omitempty" xml:"video,omitempty" require:"true"`
}

func (s VideoUploadVideoPartRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoPartRequest) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoPartRequest) SetHeader(v map[string]*string) *VideoUploadVideoPartRequest {
	s.Header = v
	return s
}

func (s *VideoUploadVideoPartRequest) SetAccessToken(v string) *VideoUploadVideoPartRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoUploadVideoPartRequest) SetUploadId(v string) *VideoUploadVideoPartRequest {
	s.UploadId = &v
	return s
}

func (s *VideoUploadVideoPartRequest) SetOpenId(v string) *VideoUploadVideoPartRequest {
	s.OpenId = &v
	return s
}

func (s *VideoUploadVideoPartRequest) SetPartNumber(v int64) *VideoUploadVideoPartRequest {
	s.PartNumber = &v
	return s
}

func (s *VideoUploadVideoPartRequest) SetVideo(v *util.FileField) *VideoUploadVideoPartRequest {
	s.Video = v
	return s
}

type VideoUploadVideoPartResponse struct {
	Extra *VideoUploadVideoPartResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *VideoUploadVideoPartResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s VideoUploadVideoPartResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoPartResponse) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoPartResponse) SetExtra(v *VideoUploadVideoPartResponseExtra) *VideoUploadVideoPartResponse {
	s.Extra = v
	return s
}

func (s *VideoUploadVideoPartResponse) SetData(v *VideoUploadVideoPartResponseData) *VideoUploadVideoPartResponse {
	s.Data = v
	return s
}

type VideoUploadVideoPartResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s VideoUploadVideoPartResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoPartResponseData) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoPartResponseData) SetGwErrorCode(v int32) *VideoUploadVideoPartResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *VideoUploadVideoPartResponseData) SetGwDescription(v string) *VideoUploadVideoPartResponseData {
	s.GwDescription = &v
	return s
}

type VideoUploadVideoPartResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s VideoUploadVideoPartResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoPartResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoPartResponseExtra) SetSubDescription(v string) *VideoUploadVideoPartResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoUploadVideoPartResponseExtra) SetSubErrorCode(v int32) *VideoUploadVideoPartResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoUploadVideoPartResponseExtra) SetDescription(v string) *VideoUploadVideoPartResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoUploadVideoPartResponseExtra) SetErrorCode(v int32) *VideoUploadVideoPartResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoUploadVideoPartResponseExtra) SetLogid(v string) *VideoUploadVideoPartResponseExtra {
	s.Logid = &v
	return s
}

func (s *VideoUploadVideoPartResponseExtra) SetNow(v int64) *VideoUploadVideoPartResponseExtra {
	s.Now = &v
	return s
}

type VideoUploadVideoRequest struct {
	Video       *util.FileField    `json:"video,omitempty" xml:"video,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s VideoUploadVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoRequest) SetVideo(v *util.FileField) *VideoUploadVideoRequest {
	s.Video = v
	return s
}

func (s *VideoUploadVideoRequest) SetHeader(v map[string]*string) *VideoUploadVideoRequest {
	s.Header = v
	return s
}

func (s *VideoUploadVideoRequest) SetAccessToken(v string) *VideoUploadVideoRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoUploadVideoRequest) SetOpenId(v string) *VideoUploadVideoRequest {
	s.OpenId = &v
	return s
}

type VideoUploadVideoResponse struct {
	Data  *VideoUploadVideoResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *VideoUploadVideoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s VideoUploadVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoResponse) SetData(v *VideoUploadVideoResponseData) *VideoUploadVideoResponse {
	s.Data = v
	return s
}

func (s *VideoUploadVideoResponse) SetExtra(v *VideoUploadVideoResponseExtra) *VideoUploadVideoResponse {
	s.Extra = v
	return s
}

type VideoUploadVideoResponseData struct {
	GwErrorCode   *int32                             `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                            `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Video         *VideoUploadVideoResponseDataVideo `json:"video,omitempty" xml:"video,omitempty" require:"true"`
}

func (s VideoUploadVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoResponseData) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoResponseData) SetGwErrorCode(v int32) *VideoUploadVideoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *VideoUploadVideoResponseData) SetGwDescription(v string) *VideoUploadVideoResponseData {
	s.GwDescription = &v
	return s
}

func (s *VideoUploadVideoResponseData) SetVideo(v *VideoUploadVideoResponseDataVideo) *VideoUploadVideoResponseData {
	s.Video = v
	return s
}

type VideoUploadVideoResponseDataVideo struct {
	VideoId *string `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	Width   *int32  `json:"width,omitempty" xml:"width,omitempty" require:"true"`
	Height  *int32  `json:"height,omitempty" xml:"height,omitempty" require:"true"`
}

func (s VideoUploadVideoResponseDataVideo) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoResponseDataVideo) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoResponseDataVideo) SetVideoId(v string) *VideoUploadVideoResponseDataVideo {
	s.VideoId = &v
	return s
}

func (s *VideoUploadVideoResponseDataVideo) SetWidth(v int32) *VideoUploadVideoResponseDataVideo {
	s.Width = &v
	return s
}

func (s *VideoUploadVideoResponseDataVideo) SetHeight(v int32) *VideoUploadVideoResponseDataVideo {
	s.Height = &v
	return s
}

type VideoUploadVideoResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s VideoUploadVideoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoUploadVideoResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoUploadVideoResponseExtra) SetLogid(v string) *VideoUploadVideoResponseExtra {
	s.Logid = &v
	return s
}

func (s *VideoUploadVideoResponseExtra) SetNow(v int64) *VideoUploadVideoResponseExtra {
	s.Now = &v
	return s
}

func (s *VideoUploadVideoResponseExtra) SetSubDescription(v string) *VideoUploadVideoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoUploadVideoResponseExtra) SetSubErrorCode(v int32) *VideoUploadVideoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoUploadVideoResponseExtra) SetDescription(v string) *VideoUploadVideoResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoUploadVideoResponseExtra) SetErrorCode(v int32) *VideoUploadVideoResponseExtra {
	s.ErrorCode = &v
	return s
}

type VideoVideoBasicInfoRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ItemIds     []*string          `json:"item_ids,omitempty" xml:"item_ids,omitempty" type:"Repeated"`
	VideoIds    []*string          `json:"video_ids,omitempty" xml:"video_ids,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s VideoVideoBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *VideoVideoBasicInfoRequest) SetAccessToken(v string) *VideoVideoBasicInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *VideoVideoBasicInfoRequest) SetOpenId(v string) *VideoVideoBasicInfoRequest {
	s.OpenId = &v
	return s
}

func (s *VideoVideoBasicInfoRequest) SetItemIds(v []*string) *VideoVideoBasicInfoRequest {
	s.ItemIds = v
	return s
}

func (s *VideoVideoBasicInfoRequest) SetVideoIds(v []*string) *VideoVideoBasicInfoRequest {
	s.VideoIds = v
	return s
}

func (s *VideoVideoBasicInfoRequest) SetHeader(v map[string]*string) *VideoVideoBasicInfoRequest {
	s.Header = v
	return s
}

type VideoVideoBasicInfoResponse struct {
	Extra *VideoVideoBasicInfoResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *VideoVideoBasicInfoResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s VideoVideoBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *VideoVideoBasicInfoResponse) SetExtra(v *VideoVideoBasicInfoResponseExtra) *VideoVideoBasicInfoResponse {
	s.Extra = v
	return s
}

func (s *VideoVideoBasicInfoResponse) SetData(v *VideoVideoBasicInfoResponseData) *VideoVideoBasicInfoResponse {
	s.Data = v
	return s
}

type VideoVideoBasicInfoResponseData struct {
	List          []*VideoVideoBasicInfoResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s VideoVideoBasicInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoBasicInfoResponseData) GoString() string {
	return s.String()
}

func (s *VideoVideoBasicInfoResponseData) SetList(v []*VideoVideoBasicInfoResponseDataListItem) *VideoVideoBasicInfoResponseData {
	s.List = v
	return s
}

func (s *VideoVideoBasicInfoResponseData) SetGwErrorCode(v int32) *VideoVideoBasicInfoResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *VideoVideoBasicInfoResponseData) SetGwDescription(v string) *VideoVideoBasicInfoResponseData {
	s.GwDescription = &v
	return s
}

type VideoVideoBasicInfoResponseDataListItem struct {
	Title      *string `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	Cover      *string `json:"cover,omitempty" xml:"cover,omitempty" require:"true"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	MediaType  *int    `json:"media_type,omitempty" xml:"media_type,omitempty" require:"true"`
	VideoId    *string `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	ItemId     *string `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
}

func (s VideoVideoBasicInfoResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoBasicInfoResponseDataListItem) GoString() string {
	return s.String()
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetTitle(v string) *VideoVideoBasicInfoResponseDataListItem {
	s.Title = &v
	return s
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetCover(v string) *VideoVideoBasicInfoResponseDataListItem {
	s.Cover = &v
	return s
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetCreateTime(v int64) *VideoVideoBasicInfoResponseDataListItem {
	s.CreateTime = &v
	return s
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetMediaType(v int) *VideoVideoBasicInfoResponseDataListItem {
	s.MediaType = &v
	return s
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetVideoId(v string) *VideoVideoBasicInfoResponseDataListItem {
	s.VideoId = &v
	return s
}

func (s *VideoVideoBasicInfoResponseDataListItem) SetItemId(v string) *VideoVideoBasicInfoResponseDataListItem {
	s.ItemId = &v
	return s
}

type VideoVideoBasicInfoResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s VideoVideoBasicInfoResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoBasicInfoResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoVideoBasicInfoResponseExtra) SetErrorCode(v int32) *VideoVideoBasicInfoResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoVideoBasicInfoResponseExtra) SetDescription(v string) *VideoVideoBasicInfoResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoVideoBasicInfoResponseExtra) SetSubErrorCode(v int32) *VideoVideoBasicInfoResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoVideoBasicInfoResponseExtra) SetSubDescription(v string) *VideoVideoBasicInfoResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoVideoBasicInfoResponseExtra) SetLogid(v string) *VideoVideoBasicInfoResponseExtra {
	s.Logid = &v
	return s
}

func (s *VideoVideoBasicInfoResponseExtra) SetNow(v int64) *VideoVideoBasicInfoResponseExtra {
	s.Now = &v
	return s
}

type VideoVideoDataRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	VideoIds    []*string          `json:"video_ids,omitempty" xml:"video_ids,omitempty" type:"Repeated"`
	ItemIds     []*string          `json:"item_ids,omitempty" xml:"item_ids,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoVideoDataRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataRequest) GoString() string {
	return s.String()
}

func (s *VideoVideoDataRequest) SetOpenId(v string) *VideoVideoDataRequest {
	s.OpenId = &v
	return s
}

func (s *VideoVideoDataRequest) SetVideoIds(v []*string) *VideoVideoDataRequest {
	s.VideoIds = v
	return s
}

func (s *VideoVideoDataRequest) SetItemIds(v []*string) *VideoVideoDataRequest {
	s.ItemIds = v
	return s
}

func (s *VideoVideoDataRequest) SetHeader(v map[string]*string) *VideoVideoDataRequest {
	s.Header = v
	return s
}

func (s *VideoVideoDataRequest) SetAccessToken(v string) *VideoVideoDataRequest {
	s.AccessToken = &v
	return s
}

type VideoVideoDataResponse struct {
	Extra *VideoVideoDataResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *VideoVideoDataResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s VideoVideoDataResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataResponse) GoString() string {
	return s.String()
}

func (s *VideoVideoDataResponse) SetExtra(v *VideoVideoDataResponseExtra) *VideoVideoDataResponse {
	s.Extra = v
	return s
}

func (s *VideoVideoDataResponse) SetData(v *VideoVideoDataResponseData) *VideoVideoDataResponse {
	s.Data = v
	return s
}

type VideoVideoDataResponseData struct {
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*VideoVideoDataResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s VideoVideoDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataResponseData) GoString() string {
	return s.String()
}

func (s *VideoVideoDataResponseData) SetGwDescription(v string) *VideoVideoDataResponseData {
	s.GwDescription = &v
	return s
}

func (s *VideoVideoDataResponseData) SetList(v []*VideoVideoDataResponseDataListItem) *VideoVideoDataResponseData {
	s.List = v
	return s
}

func (s *VideoVideoDataResponseData) SetGwErrorCode(v int32) *VideoVideoDataResponseData {
	s.GwErrorCode = &v
	return s
}

type VideoVideoDataResponseDataListItem struct {
	CreateTime *int64                                        `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	Cover      *string                                       `json:"cover,omitempty" xml:"cover,omitempty" require:"true"`
	ShareUrl   *string                                       `json:"share_url,omitempty" xml:"share_url,omitempty" require:"true"`
	ItemId     *string                                       `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	MediaType  *int                                          `json:"media_type,omitempty" xml:"media_type,omitempty"`
	IsTop      *bool                                         `json:"is_top,omitempty" xml:"is_top,omitempty" require:"true"`
	Statistics *VideoVideoDataResponseDataListItemStatistics `json:"statistics,omitempty" xml:"statistics,omitempty"`
	VideoId    *string                                       `json:"video_id,omitempty" xml:"video_id,omitempty"`
	Title      *string                                       `json:"title,omitempty" xml:"title,omitempty" require:"true"`
}

func (s VideoVideoDataResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataResponseDataListItem) GoString() string {
	return s.String()
}

func (s *VideoVideoDataResponseDataListItem) SetCreateTime(v int64) *VideoVideoDataResponseDataListItem {
	s.CreateTime = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetCover(v string) *VideoVideoDataResponseDataListItem {
	s.Cover = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetShareUrl(v string) *VideoVideoDataResponseDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetItemId(v string) *VideoVideoDataResponseDataListItem {
	s.ItemId = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetMediaType(v int) *VideoVideoDataResponseDataListItem {
	s.MediaType = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetIsTop(v bool) *VideoVideoDataResponseDataListItem {
	s.IsTop = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetStatistics(v *VideoVideoDataResponseDataListItemStatistics) *VideoVideoDataResponseDataListItem {
	s.Statistics = v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetVideoId(v string) *VideoVideoDataResponseDataListItem {
	s.VideoId = &v
	return s
}

func (s *VideoVideoDataResponseDataListItem) SetTitle(v string) *VideoVideoDataResponseDataListItem {
	s.Title = &v
	return s
}

type VideoVideoDataResponseDataListItemStatistics struct {
	PlayCount     *int32 `json:"play_count,omitempty" xml:"play_count,omitempty" require:"true"`
	ShareCount    *int32 `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	ForwardCount  *int32 `json:"forward_count,omitempty" xml:"forward_count,omitempty" require:"true"`
	CommentCount  *int32 `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
	DiggCount     *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	DownloadCount *int32 `json:"download_count,omitempty" xml:"download_count,omitempty" require:"true"`
}

func (s VideoVideoDataResponseDataListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataResponseDataListItemStatistics) GoString() string {
	return s.String()
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetPlayCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.PlayCount = &v
	return s
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetShareCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.ShareCount = &v
	return s
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetForwardCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.ForwardCount = &v
	return s
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetCommentCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.CommentCount = &v
	return s
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetDiggCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.DiggCount = &v
	return s
}

func (s *VideoVideoDataResponseDataListItemStatistics) SetDownloadCount(v int32) *VideoVideoDataResponseDataListItemStatistics {
	s.DownloadCount = &v
	return s
}

type VideoVideoDataResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s VideoVideoDataResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoDataResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoVideoDataResponseExtra) SetLogid(v string) *VideoVideoDataResponseExtra {
	s.Logid = &v
	return s
}

func (s *VideoVideoDataResponseExtra) SetNow(v int64) *VideoVideoDataResponseExtra {
	s.Now = &v
	return s
}

func (s *VideoVideoDataResponseExtra) SetErrorCode(v int32) *VideoVideoDataResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoVideoDataResponseExtra) SetDescription(v string) *VideoVideoDataResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoVideoDataResponseExtra) SetSubErrorCode(v int32) *VideoVideoDataResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoVideoDataResponseExtra) SetSubDescription(v string) *VideoVideoDataResponseExtra {
	s.SubDescription = &v
	return s
}

type VideoVideoListRequest struct {
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s VideoVideoListRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListRequest) GoString() string {
	return s.String()
}

func (s *VideoVideoListRequest) SetCursor(v int64) *VideoVideoListRequest {
	s.Cursor = &v
	return s
}

func (s *VideoVideoListRequest) SetOpenId(v string) *VideoVideoListRequest {
	s.OpenId = &v
	return s
}

func (s *VideoVideoListRequest) SetCount(v int32) *VideoVideoListRequest {
	s.Count = &v
	return s
}

func (s *VideoVideoListRequest) SetHeader(v map[string]*string) *VideoVideoListRequest {
	s.Header = v
	return s
}

func (s *VideoVideoListRequest) SetAccessToken(v string) *VideoVideoListRequest {
	s.AccessToken = &v
	return s
}

type VideoVideoListResponse struct {
	Data  *VideoVideoListResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *VideoVideoListResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s VideoVideoListResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListResponse) GoString() string {
	return s.String()
}

func (s *VideoVideoListResponse) SetData(v *VideoVideoListResponseData) *VideoVideoListResponse {
	s.Data = v
	return s
}

func (s *VideoVideoListResponse) SetExtra(v *VideoVideoListResponseExtra) *VideoVideoListResponse {
	s.Extra = v
	return s
}

type VideoVideoListResponseData struct {
	List          []*VideoVideoListResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor        *int64                                `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                 `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
}

func (s VideoVideoListResponseData) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListResponseData) GoString() string {
	return s.String()
}

func (s *VideoVideoListResponseData) SetList(v []*VideoVideoListResponseDataListItem) *VideoVideoListResponseData {
	s.List = v
	return s
}

func (s *VideoVideoListResponseData) SetGwErrorCode(v int32) *VideoVideoListResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *VideoVideoListResponseData) SetGwDescription(v string) *VideoVideoListResponseData {
	s.GwDescription = &v
	return s
}

func (s *VideoVideoListResponseData) SetCursor(v int64) *VideoVideoListResponseData {
	s.Cursor = &v
	return s
}

func (s *VideoVideoListResponseData) SetHasMore(v bool) *VideoVideoListResponseData {
	s.HasMore = &v
	return s
}

type VideoVideoListResponseDataListItem struct {
	ShareUrl    *string                                       `json:"share_url,omitempty" xml:"share_url,omitempty" require:"true"`
	IsTop       *bool                                         `json:"is_top,omitempty" xml:"is_top,omitempty" require:"true"`
	Cover       *string                                       `json:"cover,omitempty" xml:"cover,omitempty" require:"true"`
	Title       *string                                       `json:"title,omitempty" xml:"title,omitempty" require:"true"`
	CreateTime  *int64                                        `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	IsReviewed  *bool                                         `json:"is_reviewed,omitempty" xml:"is_reviewed,omitempty" require:"true"`
	Statistics  *VideoVideoListResponseDataListItemStatistics `json:"statistics,omitempty" xml:"statistics,omitempty"`
	MediaType   *int                                          `json:"media_type,omitempty" xml:"media_type,omitempty"`
	VideoStatus *int32                                        `json:"video_status,omitempty" xml:"video_status,omitempty" require:"true"`
	ItemId      *string                                       `json:"item_id,omitempty" xml:"item_id,omitempty" require:"true"`
	VideoId     *string                                       `json:"video_id,omitempty" xml:"video_id,omitempty"`
}

func (s VideoVideoListResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListResponseDataListItem) GoString() string {
	return s.String()
}

func (s *VideoVideoListResponseDataListItem) SetShareUrl(v string) *VideoVideoListResponseDataListItem {
	s.ShareUrl = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetIsTop(v bool) *VideoVideoListResponseDataListItem {
	s.IsTop = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetCover(v string) *VideoVideoListResponseDataListItem {
	s.Cover = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetTitle(v string) *VideoVideoListResponseDataListItem {
	s.Title = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetCreateTime(v int64) *VideoVideoListResponseDataListItem {
	s.CreateTime = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetIsReviewed(v bool) *VideoVideoListResponseDataListItem {
	s.IsReviewed = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetStatistics(v *VideoVideoListResponseDataListItemStatistics) *VideoVideoListResponseDataListItem {
	s.Statistics = v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetMediaType(v int) *VideoVideoListResponseDataListItem {
	s.MediaType = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetVideoStatus(v int32) *VideoVideoListResponseDataListItem {
	s.VideoStatus = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetItemId(v string) *VideoVideoListResponseDataListItem {
	s.ItemId = &v
	return s
}

func (s *VideoVideoListResponseDataListItem) SetVideoId(v string) *VideoVideoListResponseDataListItem {
	s.VideoId = &v
	return s
}

type VideoVideoListResponseDataListItemStatistics struct {
	ForwardCount  *int32 `json:"forward_count,omitempty" xml:"forward_count,omitempty" require:"true"`
	PlayCount     *int32 `json:"play_count,omitempty" xml:"play_count,omitempty" require:"true"`
	ShareCount    *int32 `json:"share_count,omitempty" xml:"share_count,omitempty" require:"true"`
	CommentCount  *int32 `json:"comment_count,omitempty" xml:"comment_count,omitempty" require:"true"`
	DiggCount     *int32 `json:"digg_count,omitempty" xml:"digg_count,omitempty" require:"true"`
	DownloadCount *int32 `json:"download_count,omitempty" xml:"download_count,omitempty" require:"true"`
}

func (s VideoVideoListResponseDataListItemStatistics) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListResponseDataListItemStatistics) GoString() string {
	return s.String()
}

func (s *VideoVideoListResponseDataListItemStatistics) SetForwardCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.ForwardCount = &v
	return s
}

func (s *VideoVideoListResponseDataListItemStatistics) SetPlayCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.PlayCount = &v
	return s
}

func (s *VideoVideoListResponseDataListItemStatistics) SetShareCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.ShareCount = &v
	return s
}

func (s *VideoVideoListResponseDataListItemStatistics) SetCommentCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.CommentCount = &v
	return s
}

func (s *VideoVideoListResponseDataListItemStatistics) SetDiggCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.DiggCount = &v
	return s
}

func (s *VideoVideoListResponseDataListItemStatistics) SetDownloadCount(v int32) *VideoVideoListResponseDataListItemStatistics {
	s.DownloadCount = &v
	return s
}

type VideoVideoListResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s VideoVideoListResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s VideoVideoListResponseExtra) GoString() string {
	return s.String()
}

func (s *VideoVideoListResponseExtra) SetNow(v int64) *VideoVideoListResponseExtra {
	s.Now = &v
	return s
}

func (s *VideoVideoListResponseExtra) SetSubDescription(v string) *VideoVideoListResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *VideoVideoListResponseExtra) SetSubErrorCode(v int32) *VideoVideoListResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *VideoVideoListResponseExtra) SetDescription(v string) *VideoVideoListResponseExtra {
	s.Description = &v
	return s
}

func (s *VideoVideoListResponseExtra) SetErrorCode(v int32) *VideoVideoListResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *VideoVideoListResponseExtra) SetLogid(v string) *VideoVideoListResponseExtra {
	s.Logid = &v
	return s
}

type WebcastmateInfoRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Token       *string            `json:"token,omitempty" xml:"token,omitempty" require:"true"`
}

func (s WebcastmateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoRequest) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoRequest) SetHeader(v map[string]*string) *WebcastmateInfoRequest {
	s.Header = v
	return s
}

func (s *WebcastmateInfoRequest) SetAccessToken(v string) *WebcastmateInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *WebcastmateInfoRequest) SetToken(v string) *WebcastmateInfoRequest {
	s.Token = &v
	return s
}

type WebcastmateInfoResponse struct {
	Data *WebcastmateInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s WebcastmateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponse) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponse) SetData(v *WebcastmateInfoResponseData) *WebcastmateInfoResponse {
	s.Data = v
	return s
}

type WebcastmateInfoResponseData struct {
	AckCfg     []*WebcastmateInfoResponseDataAckCfgItem `json:"ack_cfg,omitempty" xml:"ack_cfg,omitempty" type:"Repeated"`
	LinkerInfo *WebcastmateInfoResponseDataLinkerInfo   `json:"linker_info,omitempty" xml:"linker_info,omitempty"`
	Info       *WebcastmateInfoResponseDataInfo         `json:"info,omitempty" xml:"info,omitempty"`
}

func (s WebcastmateInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponseData) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponseData) SetAckCfg(v []*WebcastmateInfoResponseDataAckCfgItem) *WebcastmateInfoResponseData {
	s.AckCfg = v
	return s
}

func (s *WebcastmateInfoResponseData) SetLinkerInfo(v *WebcastmateInfoResponseDataLinkerInfo) *WebcastmateInfoResponseData {
	s.LinkerInfo = v
	return s
}

func (s *WebcastmateInfoResponseData) SetInfo(v *WebcastmateInfoResponseDataInfo) *WebcastmateInfoResponseData {
	s.Info = v
	return s
}

type WebcastmateInfoResponseDataAckCfgItem struct {
	MsgType       *string `json:"msg_type,omitempty" xml:"msg_type,omitempty"`
	AckType       *int64  `json:"ack_type,omitempty" xml:"ack_type,omitempty"`
	BatchInterval *int64  `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	BatchMaxNum   *int64  `json:"batch_max_num,omitempty" xml:"batch_max_num,omitempty"`
}

func (s WebcastmateInfoResponseDataAckCfgItem) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponseDataAckCfgItem) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponseDataAckCfgItem) SetMsgType(v string) *WebcastmateInfoResponseDataAckCfgItem {
	s.MsgType = &v
	return s
}

func (s *WebcastmateInfoResponseDataAckCfgItem) SetAckType(v int64) *WebcastmateInfoResponseDataAckCfgItem {
	s.AckType = &v
	return s
}

func (s *WebcastmateInfoResponseDataAckCfgItem) SetBatchInterval(v int64) *WebcastmateInfoResponseDataAckCfgItem {
	s.BatchInterval = &v
	return s
}

func (s *WebcastmateInfoResponseDataAckCfgItem) SetBatchMaxNum(v int64) *WebcastmateInfoResponseDataAckCfgItem {
	s.BatchMaxNum = &v
	return s
}

type WebcastmateInfoResponseDataInfo struct {
	JoinGameUserRole    *int64   `json:"join_game_user_role,omitempty" xml:"join_game_user_role,omitempty"`
	RoomId              *int64   `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AnchorOpenId        *string  `json:"anchor_open_id,omitempty" xml:"anchor_open_id,omitempty"`
	AvatarUrl           *string  `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	NickName            *string  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	SupportedGameScenes []*int64 `json:"supported_game_scenes,omitempty" xml:"supported_game_scenes,omitempty" type:"Repeated"`
	JoinGameUserOpenId  *string  `json:"join_game_user_open_id,omitempty" xml:"join_game_user_open_id,omitempty"`
}

func (s WebcastmateInfoResponseDataInfo) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponseDataInfo) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponseDataInfo) SetJoinGameUserRole(v int64) *WebcastmateInfoResponseDataInfo {
	s.JoinGameUserRole = &v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetRoomId(v int64) *WebcastmateInfoResponseDataInfo {
	s.RoomId = &v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetAnchorOpenId(v string) *WebcastmateInfoResponseDataInfo {
	s.AnchorOpenId = &v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetAvatarUrl(v string) *WebcastmateInfoResponseDataInfo {
	s.AvatarUrl = &v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetNickName(v string) *WebcastmateInfoResponseDataInfo {
	s.NickName = &v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetSupportedGameScenes(v []*int64) *WebcastmateInfoResponseDataInfo {
	s.SupportedGameScenes = v
	return s
}

func (s *WebcastmateInfoResponseDataInfo) SetJoinGameUserOpenId(v string) *WebcastmateInfoResponseDataInfo {
	s.JoinGameUserOpenId = &v
	return s
}

type WebcastmateInfoResponseDataLinkerInfo struct {
	RewardInfo           []*string                                            `json:"reward_info,omitempty" xml:"reward_info,omitempty" type:"Repeated"`
	CurReward            *string                                              `json:"cur_reward,omitempty" xml:"cur_reward,omitempty"`
	MasterStatus         *int64                                               `json:"master_status,omitempty" xml:"master_status,omitempty"`
	LinkerActivityMaster *string                                              `json:"linker_activity_master,omitempty" xml:"linker_activity_master,omitempty"`
	UserInfo             []*WebcastmateInfoResponseDataLinkerInfoUserInfoItem `json:"user_info,omitempty" xml:"user_info,omitempty" type:"Repeated"`
	LinkerId             *int64                                               `json:"linker_id,omitempty" xml:"linker_id,omitempty"`
	LinkerStatus         *int64                                               `json:"linker_status,omitempty" xml:"linker_status,omitempty"`
	ActivityMasterOpenid *string                                              `json:"activity_master_openid,omitempty" xml:"activity_master_openid,omitempty"`
}

func (s WebcastmateInfoResponseDataLinkerInfo) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponseDataLinkerInfo) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetRewardInfo(v []*string) *WebcastmateInfoResponseDataLinkerInfo {
	s.RewardInfo = v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetCurReward(v string) *WebcastmateInfoResponseDataLinkerInfo {
	s.CurReward = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetMasterStatus(v int64) *WebcastmateInfoResponseDataLinkerInfo {
	s.MasterStatus = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetLinkerActivityMaster(v string) *WebcastmateInfoResponseDataLinkerInfo {
	s.LinkerActivityMaster = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetUserInfo(v []*WebcastmateInfoResponseDataLinkerInfoUserInfoItem) *WebcastmateInfoResponseDataLinkerInfo {
	s.UserInfo = v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetLinkerId(v int64) *WebcastmateInfoResponseDataLinkerInfo {
	s.LinkerId = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetLinkerStatus(v int64) *WebcastmateInfoResponseDataLinkerInfo {
	s.LinkerStatus = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfo) SetActivityMasterOpenid(v string) *WebcastmateInfoResponseDataLinkerInfo {
	s.ActivityMasterOpenid = &v
	return s
}

type WebcastmateInfoResponseDataLinkerInfoUserInfoItem struct {
	NickName  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	OpenId    *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
}

func (s WebcastmateInfoResponseDataLinkerInfoUserInfoItem) String() string {
	return tea.Prettify(s)
}

func (s WebcastmateInfoResponseDataLinkerInfoUserInfoItem) GoString() string {
	return s.String()
}

func (s *WebcastmateInfoResponseDataLinkerInfoUserInfoItem) SetNickName(v string) *WebcastmateInfoResponseDataLinkerInfoUserInfoItem {
	s.NickName = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfoUserInfoItem) SetOpenId(v string) *WebcastmateInfoResponseDataLinkerInfoUserInfoItem {
	s.OpenId = &v
	return s
}

func (s *WebcastmateInfoResponseDataLinkerInfoUserInfoItem) SetAvatarUrl(v string) *WebcastmateInfoResponseDataLinkerInfoUserInfoItem {
	s.AvatarUrl = &v
	return s
}

type WithdrawCateringQueryRequest struct {
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	EndDate     *string            `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	Size        *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	StartDate   *string            `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s WithdrawCateringQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCateringQueryRequest) GoString() string {
	return s.String()
}

func (s *WithdrawCateringQueryRequest) SetCursor(v string) *WithdrawCateringQueryRequest {
	s.Cursor = &v
	return s
}

func (s *WithdrawCateringQueryRequest) SetEndDate(v string) *WithdrawCateringQueryRequest {
	s.EndDate = &v
	return s
}

func (s *WithdrawCateringQueryRequest) SetSize(v int64) *WithdrawCateringQueryRequest {
	s.Size = &v
	return s
}

func (s *WithdrawCateringQueryRequest) SetStartDate(v string) *WithdrawCateringQueryRequest {
	s.StartDate = &v
	return s
}

func (s *WithdrawCateringQueryRequest) SetHeader(v map[string]*string) *WithdrawCateringQueryRequest {
	s.Header = v
	return s
}

func (s *WithdrawCateringQueryRequest) SetAccessToken(v string) *WithdrawCateringQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *WithdrawCateringQueryRequest) SetAccountId(v string) *WithdrawCateringQueryRequest {
	s.AccountId = &v
	return s
}

type WithdrawCateringQueryResponse struct {
	Data  *WithdrawCateringQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *WithdrawCateringQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s WithdrawCateringQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCateringQueryResponse) GoString() string {
	return s.String()
}

func (s *WithdrawCateringQueryResponse) SetData(v *WithdrawCateringQueryResponseData) *WithdrawCateringQueryResponse {
	s.Data = v
	return s
}

func (s *WithdrawCateringQueryResponse) SetExtra(v *WithdrawCateringQueryResponseExtra) *WithdrawCateringQueryResponse {
	s.Extra = v
	return s
}

type WithdrawCateringQueryResponseData struct {
	GwDescription       *string                                                     `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	BillWithdrawRecords []*WithdrawCateringQueryResponseDataBillWithdrawRecordsItem `json:"bill_withdraw_records,omitempty" xml:"bill_withdraw_records,omitempty" require:"true" type:"Repeated"`
	Cursor              *string                                                     `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore             *bool                                                       `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	GwErrorCode         *int32                                                      `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s WithdrawCateringQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCateringQueryResponseData) GoString() string {
	return s.String()
}

func (s *WithdrawCateringQueryResponseData) SetGwDescription(v string) *WithdrawCateringQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *WithdrawCateringQueryResponseData) SetBillWithdrawRecords(v []*WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) *WithdrawCateringQueryResponseData {
	s.BillWithdrawRecords = v
	return s
}

func (s *WithdrawCateringQueryResponseData) SetCursor(v string) *WithdrawCateringQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *WithdrawCateringQueryResponseData) SetHasMore(v bool) *WithdrawCateringQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *WithdrawCateringQueryResponseData) SetGwErrorCode(v int32) *WithdrawCateringQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type WithdrawCateringQueryResponseDataBillWithdrawRecordsItem struct {
	WithdrawMsg   *string `json:"withdraw_msg,omitempty" xml:"withdraw_msg,omitempty"`
	WithdrawId    *string `json:"withdraw_id,omitempty" xml:"withdraw_id,omitempty"`
	WithdrawTime  *int64  `json:"withdraw_time,omitempty" xml:"withdraw_time,omitempty"`
	WithdrawMode  *int    `json:"withdraw_mode,omitempty" xml:"withdraw_mode,omitempty"`
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
	SettleAccount *string `json:"settle_account,omitempty" xml:"settle_account,omitempty"`
	AccountName   *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	Amount        *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	LaunchTime    *int64  `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
}

func (s WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) GoString() string {
	return s.String()
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetWithdrawMsg(v string) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawMsg = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetWithdrawId(v string) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawId = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetWithdrawTime(v int64) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawTime = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetWithdrawMode(v int) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawMode = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetStatus(v int64) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.Status = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetSettleAccount(v string) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.SettleAccount = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetAccountName(v string) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.AccountName = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetAmount(v int64) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.Amount = &v
	return s
}

func (s *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem) SetLaunchTime(v int64) *WithdrawCateringQueryResponseDataBillWithdrawRecordsItem {
	s.LaunchTime = &v
	return s
}

type WithdrawCateringQueryResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s WithdrawCateringQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCateringQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *WithdrawCateringQueryResponseExtra) SetLogid(v string) *WithdrawCateringQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *WithdrawCateringQueryResponseExtra) SetNow(v int64) *WithdrawCateringQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *WithdrawCateringQueryResponseExtra) SetSubDescription(v string) *WithdrawCateringQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *WithdrawCateringQueryResponseExtra) SetSubErrorCode(v int32) *WithdrawCateringQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *WithdrawCateringQueryResponseExtra) SetDescription(v string) *WithdrawCateringQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *WithdrawCateringQueryResponseExtra) SetErrorCode(v int32) *WithdrawCateringQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

type WithdrawCompositeQueryRequest struct {
	StartDate   *string            `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	EndDate     *string            `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	Size        *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s WithdrawCompositeQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCompositeQueryRequest) GoString() string {
	return s.String()
}

func (s *WithdrawCompositeQueryRequest) SetStartDate(v string) *WithdrawCompositeQueryRequest {
	s.StartDate = &v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetAccountId(v string) *WithdrawCompositeQueryRequest {
	s.AccountId = &v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetCursor(v string) *WithdrawCompositeQueryRequest {
	s.Cursor = &v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetEndDate(v string) *WithdrawCompositeQueryRequest {
	s.EndDate = &v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetSize(v int64) *WithdrawCompositeQueryRequest {
	s.Size = &v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetHeader(v map[string]*string) *WithdrawCompositeQueryRequest {
	s.Header = v
	return s
}

func (s *WithdrawCompositeQueryRequest) SetAccessToken(v string) *WithdrawCompositeQueryRequest {
	s.AccessToken = &v
	return s
}

type WithdrawCompositeQueryResponse struct {
	Data  *WithdrawCompositeQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *WithdrawCompositeQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s WithdrawCompositeQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCompositeQueryResponse) GoString() string {
	return s.String()
}

func (s *WithdrawCompositeQueryResponse) SetData(v *WithdrawCompositeQueryResponseData) *WithdrawCompositeQueryResponse {
	s.Data = v
	return s
}

func (s *WithdrawCompositeQueryResponse) SetExtra(v *WithdrawCompositeQueryResponseExtra) *WithdrawCompositeQueryResponse {
	s.Extra = v
	return s
}

type WithdrawCompositeQueryResponseData struct {
	GwDescription       *string                                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor              *string                                                      `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore             *bool                                                        `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	BillWithdrawRecords []*WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem `json:"bill_withdraw_records,omitempty" xml:"bill_withdraw_records,omitempty" require:"true" type:"Repeated"`
	GwErrorCode         *int32                                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s WithdrawCompositeQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCompositeQueryResponseData) GoString() string {
	return s.String()
}

func (s *WithdrawCompositeQueryResponseData) SetGwDescription(v string) *WithdrawCompositeQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *WithdrawCompositeQueryResponseData) SetCursor(v string) *WithdrawCompositeQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *WithdrawCompositeQueryResponseData) SetHasMore(v bool) *WithdrawCompositeQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *WithdrawCompositeQueryResponseData) SetBillWithdrawRecords(v []*WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) *WithdrawCompositeQueryResponseData {
	s.BillWithdrawRecords = v
	return s
}

func (s *WithdrawCompositeQueryResponseData) SetGwErrorCode(v int32) *WithdrawCompositeQueryResponseData {
	s.GwErrorCode = &v
	return s
}

type WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem struct {
	Status        *int64  `json:"status,omitempty" xml:"status,omitempty"`
	WithdrawTime  *int64  `json:"withdraw_time,omitempty" xml:"withdraw_time,omitempty"`
	WithdrawId    *string `json:"withdraw_id,omitempty" xml:"withdraw_id,omitempty"`
	SettleAccount *string `json:"settle_account,omitempty" xml:"settle_account,omitempty"`
	Amount        *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	LaunchTime    *int64  `json:"launch_time,omitempty" xml:"launch_time,omitempty"`
	WithdrawMode  *int    `json:"withdraw_mode,omitempty" xml:"withdraw_mode,omitempty"`
	AccountName   *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	WithdrawMsg   *string `json:"withdraw_msg,omitempty" xml:"withdraw_msg,omitempty"`
}

func (s WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) GoString() string {
	return s.String()
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetStatus(v int64) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.Status = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetWithdrawTime(v int64) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawTime = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetWithdrawId(v string) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawId = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetSettleAccount(v string) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.SettleAccount = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetAmount(v int64) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.Amount = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetLaunchTime(v int64) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.LaunchTime = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetWithdrawMode(v int) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawMode = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetAccountName(v string) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.AccountName = &v
	return s
}

func (s *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem) SetWithdrawMsg(v string) *WithdrawCompositeQueryResponseDataBillWithdrawRecordsItem {
	s.WithdrawMsg = &v
	return s
}

type WithdrawCompositeQueryResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s WithdrawCompositeQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s WithdrawCompositeQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *WithdrawCompositeQueryResponseExtra) SetSubErrorCode(v int32) *WithdrawCompositeQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *WithdrawCompositeQueryResponseExtra) SetDescription(v string) *WithdrawCompositeQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *WithdrawCompositeQueryResponseExtra) SetErrorCode(v int32) *WithdrawCompositeQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *WithdrawCompositeQueryResponseExtra) SetLogid(v string) *WithdrawCompositeQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *WithdrawCompositeQueryResponseExtra) SetNow(v int64) *WithdrawCompositeQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *WithdrawCompositeQueryResponseExtra) SetSubDescription(v string) *WithdrawCompositeQueryResponseExtra {
	s.SubDescription = &v
	return s
}

type WorldRankSetValidVersionRequest struct {
	AppId            *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	WorldRankVersion *string            `json:"world_rank_version,omitempty" xml:"world_rank_version,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	IsOnlineVersion  *bool              `json:"is_online_version,omitempty" xml:"is_online_version,omitempty" require:"true"`
}

func (s WorldRankSetValidVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s WorldRankSetValidVersionRequest) GoString() string {
	return s.String()
}

func (s *WorldRankSetValidVersionRequest) SetAppId(v string) *WorldRankSetValidVersionRequest {
	s.AppId = &v
	return s
}

func (s *WorldRankSetValidVersionRequest) SetWorldRankVersion(v string) *WorldRankSetValidVersionRequest {
	s.WorldRankVersion = &v
	return s
}

func (s *WorldRankSetValidVersionRequest) SetHeader(v map[string]*string) *WorldRankSetValidVersionRequest {
	s.Header = v
	return s
}

func (s *WorldRankSetValidVersionRequest) SetAccessToken(v string) *WorldRankSetValidVersionRequest {
	s.AccessToken = &v
	return s
}

func (s *WorldRankSetValidVersionRequest) SetIsOnlineVersion(v bool) *WorldRankSetValidVersionRequest {
	s.IsOnlineVersion = &v
	return s
}

type WorldRankSetValidVersionResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s WorldRankSetValidVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s WorldRankSetValidVersionResponse) GoString() string {
	return s.String()
}

func (s *WorldRankSetValidVersionResponse) SetErrMsg(v string) *WorldRankSetValidVersionResponse {
	s.ErrMsg = &v
	return s
}

func (s *WorldRankSetValidVersionResponse) SetErrNo(v int64) *WorldRankSetValidVersionResponse {
	s.ErrNo = &v
	return s
}

type WorldRankUploadRankListRequest struct {
	AppId            *string                                       `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	WorldRankVersion *string                                       `json:"world_rank_version,omitempty" xml:"world_rank_version,omitempty" require:"true"`
	RankList         []*WorldRankUploadRankListRequestRankListItem `json:"rank_list,omitempty" xml:"rank_list,omitempty" require:"true" type:"Repeated"`
	Header           map[string]*string                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	IsOnlineVersion  *bool                                         `json:"is_online_version,omitempty" xml:"is_online_version,omitempty"`
}

func (s WorldRankUploadRankListRequest) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadRankListRequest) GoString() string {
	return s.String()
}

func (s *WorldRankUploadRankListRequest) SetAppId(v string) *WorldRankUploadRankListRequest {
	s.AppId = &v
	return s
}

func (s *WorldRankUploadRankListRequest) SetWorldRankVersion(v string) *WorldRankUploadRankListRequest {
	s.WorldRankVersion = &v
	return s
}

func (s *WorldRankUploadRankListRequest) SetRankList(v []*WorldRankUploadRankListRequestRankListItem) *WorldRankUploadRankListRequest {
	s.RankList = v
	return s
}

func (s *WorldRankUploadRankListRequest) SetHeader(v map[string]*string) *WorldRankUploadRankListRequest {
	s.Header = v
	return s
}

func (s *WorldRankUploadRankListRequest) SetAccessToken(v string) *WorldRankUploadRankListRequest {
	s.AccessToken = &v
	return s
}

func (s *WorldRankUploadRankListRequest) SetIsOnlineVersion(v bool) *WorldRankUploadRankListRequest {
	s.IsOnlineVersion = &v
	return s
}

type WorldRankUploadRankListRequestRankListItem struct {
	WinningStreakCount *int64  `json:"winning_streak_count,omitempty" xml:"winning_streak_count,omitempty" require:"true"`
	Rank               *int64  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	OpenId             *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	WinningPoints      *int64  `json:"winning_points,omitempty" xml:"winning_points,omitempty"`
	Score              *int64  `json:"score,omitempty" xml:"score,omitempty" require:"true"`
}

func (s WorldRankUploadRankListRequestRankListItem) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadRankListRequestRankListItem) GoString() string {
	return s.String()
}

func (s *WorldRankUploadRankListRequestRankListItem) SetWinningStreakCount(v int64) *WorldRankUploadRankListRequestRankListItem {
	s.WinningStreakCount = &v
	return s
}

func (s *WorldRankUploadRankListRequestRankListItem) SetRank(v int64) *WorldRankUploadRankListRequestRankListItem {
	s.Rank = &v
	return s
}

func (s *WorldRankUploadRankListRequestRankListItem) SetOpenId(v string) *WorldRankUploadRankListRequestRankListItem {
	s.OpenId = &v
	return s
}

func (s *WorldRankUploadRankListRequestRankListItem) SetWinningPoints(v int64) *WorldRankUploadRankListRequestRankListItem {
	s.WinningPoints = &v
	return s
}

func (s *WorldRankUploadRankListRequestRankListItem) SetScore(v int64) *WorldRankUploadRankListRequestRankListItem {
	s.Score = &v
	return s
}

type WorldRankUploadRankListResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s WorldRankUploadRankListResponse) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadRankListResponse) GoString() string {
	return s.String()
}

func (s *WorldRankUploadRankListResponse) SetErrNo(v int64) *WorldRankUploadRankListResponse {
	s.ErrNo = &v
	return s
}

func (s *WorldRankUploadRankListResponse) SetErrMsg(v string) *WorldRankUploadRankListResponse {
	s.ErrMsg = &v
	return s
}

type WorldRankUploadUserResultRequest struct {
	AppId            *string                                         `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	WorldRankVersion *string                                         `json:"world_rank_version,omitempty" xml:"world_rank_version,omitempty" require:"true"`
	UserList         []*WorldRankUploadUserResultRequestUserListItem `json:"user_list,omitempty" xml:"user_list,omitempty" require:"true" type:"Repeated"`
	IsOnlineVersion  *bool                                           `json:"is_online_version,omitempty" xml:"is_online_version,omitempty" require:"true"`
	Header           map[string]*string                              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string                                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s WorldRankUploadUserResultRequest) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadUserResultRequest) GoString() string {
	return s.String()
}

func (s *WorldRankUploadUserResultRequest) SetAppId(v string) *WorldRankUploadUserResultRequest {
	s.AppId = &v
	return s
}

func (s *WorldRankUploadUserResultRequest) SetWorldRankVersion(v string) *WorldRankUploadUserResultRequest {
	s.WorldRankVersion = &v
	return s
}

func (s *WorldRankUploadUserResultRequest) SetUserList(v []*WorldRankUploadUserResultRequestUserListItem) *WorldRankUploadUserResultRequest {
	s.UserList = v
	return s
}

func (s *WorldRankUploadUserResultRequest) SetIsOnlineVersion(v bool) *WorldRankUploadUserResultRequest {
	s.IsOnlineVersion = &v
	return s
}

func (s *WorldRankUploadUserResultRequest) SetHeader(v map[string]*string) *WorldRankUploadUserResultRequest {
	s.Header = v
	return s
}

func (s *WorldRankUploadUserResultRequest) SetAccessToken(v string) *WorldRankUploadUserResultRequest {
	s.AccessToken = &v
	return s
}

type WorldRankUploadUserResultRequestUserListItem struct {
	OpenId             *string `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	WinningPoints      *int64  `json:"winning_points,omitempty" xml:"winning_points,omitempty"`
	Score              *int64  `json:"score,omitempty" xml:"score,omitempty" require:"true"`
	WinningStreakCount *int64  `json:"winning_streak_count,omitempty" xml:"winning_streak_count,omitempty"`
	Rank               *int64  `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s WorldRankUploadUserResultRequestUserListItem) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadUserResultRequestUserListItem) GoString() string {
	return s.String()
}

func (s *WorldRankUploadUserResultRequestUserListItem) SetOpenId(v string) *WorldRankUploadUserResultRequestUserListItem {
	s.OpenId = &v
	return s
}

func (s *WorldRankUploadUserResultRequestUserListItem) SetWinningPoints(v int64) *WorldRankUploadUserResultRequestUserListItem {
	s.WinningPoints = &v
	return s
}

func (s *WorldRankUploadUserResultRequestUserListItem) SetScore(v int64) *WorldRankUploadUserResultRequestUserListItem {
	s.Score = &v
	return s
}

func (s *WorldRankUploadUserResultRequestUserListItem) SetWinningStreakCount(v int64) *WorldRankUploadUserResultRequestUserListItem {
	s.WinningStreakCount = &v
	return s
}

func (s *WorldRankUploadUserResultRequestUserListItem) SetRank(v int64) *WorldRankUploadUserResultRequestUserListItem {
	s.Rank = &v
	return s
}

type WorldRankUploadUserResultResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s WorldRankUploadUserResultResponse) String() string {
	return tea.Prettify(s)
}

func (s WorldRankUploadUserResultResponse) GoString() string {
	return s.String()
}

func (s *WorldRankUploadUserResultResponse) SetErrNo(v int64) *WorldRankUploadUserResultResponse {
	s.ErrNo = &v
	return s
}

func (s *WorldRankUploadUserResultResponse) SetErrMsg(v string) *WorldRankUploadUserResultResponse {
	s.ErrMsg = &v
	return s
}
