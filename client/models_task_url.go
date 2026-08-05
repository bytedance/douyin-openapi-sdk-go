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

type TaskSubmitResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TaskSubmitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskSubmitResponseExtra) GoString() string {
	return s.String()
}

func (s *TaskSubmitResponseExtra) SetLogid(v string) *TaskSubmitResponseExtra {
	s.Logid = &v
	return s
}

func (s *TaskSubmitResponseExtra) SetNow(v int64) *TaskSubmitResponseExtra {
	s.Now = &v
	return s
}

func (s *TaskSubmitResponseExtra) SetSubDescription(v string) *TaskSubmitResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskSubmitResponseExtra) SetSubErrorCode(v int32) *TaskSubmitResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskSubmitResponseExtra) SetDescription(v string) *TaskSubmitResponseExtra {
	s.Description = &v
	return s
}

func (s *TaskSubmitResponseExtra) SetErrorCode(v int32) *TaskSubmitResponseExtra {
	s.ErrorCode = &v
	return s
}

type TaskViewRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ExtId       *string            `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TaskViewRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskViewRequest) GoString() string {
	return s.String()
}

func (s *TaskViewRequest) SetAccessToken(v string) *TaskViewRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskViewRequest) SetTaskId(v string) *TaskViewRequest {
	s.TaskId = &v
	return s
}

func (s *TaskViewRequest) SetAccountId(v string) *TaskViewRequest {
	s.AccountId = &v
	return s
}

func (s *TaskViewRequest) SetExtId(v string) *TaskViewRequest {
	s.ExtId = &v
	return s
}

func (s *TaskViewRequest) SetHeader(v map[string]*string) *TaskViewRequest {
	s.Header = v
	return s
}

type TaskViewResponse struct {
	Data   *TaskViewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                 `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	Extra  *TaskViewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	LogId  *string                `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s TaskViewResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskViewResponse) GoString() string {
	return s.String()
}

func (s *TaskViewResponse) SetData(v *TaskViewResponseData) *TaskViewResponse {
	s.Data = v
	return s
}

func (s *TaskViewResponse) SetErrMsg(v string) *TaskViewResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskViewResponse) SetErrNo(v int32) *TaskViewResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskViewResponse) SetExtra(v *TaskViewResponseExtra) *TaskViewResponse {
	s.Extra = v
	return s
}

func (s *TaskViewResponse) SetLogId(v string) *TaskViewResponse {
	s.LogId = &v
	return s
}

type TaskViewResponseData struct {
	AddStatus          *int32   `json:"add_status,omitempty" xml:"add_status,omitempty"`
	PoiStatus          *int32   `json:"poi_status,omitempty" xml:"poi_status,omitempty"`
	MatchStatus        *int32   `json:"match_status,omitempty" xml:"match_status,omitempty" require:"true"`
	AddRejectReason    *string  `json:"add_reject_reason,omitempty" xml:"add_reject_reason,omitempty"`
	OpenTimes          *string  `json:"open_times,omitempty" xml:"open_times,omitempty"`
	TypeCode           *string  `json:"type_code,omitempty" xml:"type_code,omitempty"`
	GwErrorCode        *int32   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	PoiName            *string  `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	GwDescription      *string  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	MatchFailureReason *string  `json:"match_failure_reason,omitempty" xml:"match_failure_reason,omitempty"`
	Region             *string  `json:"region,omitempty" xml:"region,omitempty"`
	Message            *string  `json:"message,omitempty" xml:"message,omitempty"`
	OpenStatus         *int32   `json:"open_status,omitempty" xml:"open_status,omitempty"`
	PoiId              *string  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Latitude           *float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Address            *string  `json:"address,omitempty" xml:"address,omitempty"`
	Longitude          *float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s TaskViewResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskViewResponseData) GoString() string {
	return s.String()
}

func (s *TaskViewResponseData) SetAddStatus(v int32) *TaskViewResponseData {
	s.AddStatus = &v
	return s
}

func (s *TaskViewResponseData) SetPoiStatus(v int32) *TaskViewResponseData {
	s.PoiStatus = &v
	return s
}

func (s *TaskViewResponseData) SetMatchStatus(v int32) *TaskViewResponseData {
	s.MatchStatus = &v
	return s
}

func (s *TaskViewResponseData) SetAddRejectReason(v string) *TaskViewResponseData {
	s.AddRejectReason = &v
	return s
}

func (s *TaskViewResponseData) SetOpenTimes(v string) *TaskViewResponseData {
	s.OpenTimes = &v
	return s
}

func (s *TaskViewResponseData) SetTypeCode(v string) *TaskViewResponseData {
	s.TypeCode = &v
	return s
}

func (s *TaskViewResponseData) SetGwErrorCode(v int32) *TaskViewResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TaskViewResponseData) SetPoiName(v string) *TaskViewResponseData {
	s.PoiName = &v
	return s
}

func (s *TaskViewResponseData) SetGwDescription(v string) *TaskViewResponseData {
	s.GwDescription = &v
	return s
}

func (s *TaskViewResponseData) SetMatchFailureReason(v string) *TaskViewResponseData {
	s.MatchFailureReason = &v
	return s
}

func (s *TaskViewResponseData) SetRegion(v string) *TaskViewResponseData {
	s.Region = &v
	return s
}

func (s *TaskViewResponseData) SetMessage(v string) *TaskViewResponseData {
	s.Message = &v
	return s
}

func (s *TaskViewResponseData) SetOpenStatus(v int32) *TaskViewResponseData {
	s.OpenStatus = &v
	return s
}

func (s *TaskViewResponseData) SetPoiId(v string) *TaskViewResponseData {
	s.PoiId = &v
	return s
}

func (s *TaskViewResponseData) SetLatitude(v float64) *TaskViewResponseData {
	s.Latitude = &v
	return s
}

func (s *TaskViewResponseData) SetAddress(v string) *TaskViewResponseData {
	s.Address = &v
	return s
}

func (s *TaskViewResponseData) SetLongitude(v float64) *TaskViewResponseData {
	s.Longitude = &v
	return s
}

type TaskViewResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TaskViewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskViewResponseExtra) GoString() string {
	return s.String()
}

func (s *TaskViewResponseExtra) SetErrorCode(v int32) *TaskViewResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TaskViewResponseExtra) SetLogid(v string) *TaskViewResponseExtra {
	s.Logid = &v
	return s
}

func (s *TaskViewResponseExtra) SetNow(v int64) *TaskViewResponseExtra {
	s.Now = &v
	return s
}

func (s *TaskViewResponseExtra) SetSubDescription(v string) *TaskViewResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskViewResponseExtra) SetSubErrorCode(v int32) *TaskViewResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskViewResponseExtra) SetDescription(v string) *TaskViewResponseExtra {
	s.Description = &v
	return s
}

type TaskWriteoffLiveRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskWriteoffLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffLiveRequest) GoString() string {
	return s.String()
}

func (s *TaskWriteoffLiveRequest) SetOpenId(v string) *TaskWriteoffLiveRequest {
	s.OpenId = &v
	return s
}

func (s *TaskWriteoffLiveRequest) SetTaskId(v string) *TaskWriteoffLiveRequest {
	s.TaskId = &v
	return s
}

func (s *TaskWriteoffLiveRequest) SetHeader(v map[string]*string) *TaskWriteoffLiveRequest {
	s.Header = v
	return s
}

func (s *TaskWriteoffLiveRequest) SetAccessToken(v string) *TaskWriteoffLiveRequest {
	s.AccessToken = &v
	return s
}

type TaskWriteoffLiveResponse struct {
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskWriteoffLiveResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TaskWriteoffLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffLiveResponse) GoString() string {
	return s.String()
}

func (s *TaskWriteoffLiveResponse) SetErrNo(v int32) *TaskWriteoffLiveResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskWriteoffLiveResponse) SetErrMsg(v string) *TaskWriteoffLiveResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskWriteoffLiveResponse) SetLogId(v string) *TaskWriteoffLiveResponse {
	s.LogId = &v
	return s
}

func (s *TaskWriteoffLiveResponse) SetData(v *TaskWriteoffLiveResponseData) *TaskWriteoffLiveResponse {
	s.Data = v
	return s
}

type TaskWriteoffLiveResponseData struct {
	Extra  *TaskWriteoffLiveResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Result *bool                              `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TaskWriteoffLiveResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffLiveResponseData) GoString() string {
	return s.String()
}

func (s *TaskWriteoffLiveResponseData) SetExtra(v *TaskWriteoffLiveResponseDataExtra) *TaskWriteoffLiveResponseData {
	s.Extra = v
	return s
}

func (s *TaskWriteoffLiveResponseData) SetResult(v bool) *TaskWriteoffLiveResponseData {
	s.Result = &v
	return s
}

type TaskWriteoffLiveResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TaskWriteoffLiveResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffLiveResponseDataExtra) GoString() string {
	return s.String()
}

func (s *TaskWriteoffLiveResponseDataExtra) SetNow(v int64) *TaskWriteoffLiveResponseDataExtra {
	s.Now = &v
	return s
}

func (s *TaskWriteoffLiveResponseDataExtra) SetErrorCode(v int32) *TaskWriteoffLiveResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *TaskWriteoffLiveResponseDataExtra) SetDescription(v string) *TaskWriteoffLiveResponseDataExtra {
	s.Description = &v
	return s
}

func (s *TaskWriteoffLiveResponseDataExtra) SetSubErrorCode(v int32) *TaskWriteoffLiveResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskWriteoffLiveResponseDataExtra) SetSubDescription(v string) *TaskWriteoffLiveResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskWriteoffLiveResponseDataExtra) SetLogid(v string) *TaskWriteoffLiveResponseDataExtra {
	s.Logid = &v
	return s
}

type TaskWriteoffVideoRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	TaskId      *string            `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskWriteoffVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffVideoRequest) GoString() string {
	return s.String()
}

func (s *TaskWriteoffVideoRequest) SetOpenId(v string) *TaskWriteoffVideoRequest {
	s.OpenId = &v
	return s
}

func (s *TaskWriteoffVideoRequest) SetTaskId(v string) *TaskWriteoffVideoRequest {
	s.TaskId = &v
	return s
}

func (s *TaskWriteoffVideoRequest) SetHeader(v map[string]*string) *TaskWriteoffVideoRequest {
	s.Header = v
	return s
}

func (s *TaskWriteoffVideoRequest) SetAccessToken(v string) *TaskWriteoffVideoRequest {
	s.AccessToken = &v
	return s
}

type TaskWriteoffVideoResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskWriteoffVideoResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TaskWriteoffVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffVideoResponse) GoString() string {
	return s.String()
}

func (s *TaskWriteoffVideoResponse) SetErrNo(v int32) *TaskWriteoffVideoResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskWriteoffVideoResponse) SetErrMsg(v string) *TaskWriteoffVideoResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskWriteoffVideoResponse) SetLogId(v string) *TaskWriteoffVideoResponse {
	s.LogId = &v
	return s
}

func (s *TaskWriteoffVideoResponse) SetData(v *TaskWriteoffVideoResponseData) *TaskWriteoffVideoResponse {
	s.Data = v
	return s
}

type TaskWriteoffVideoResponseData struct {
	Extra  *TaskWriteoffVideoResponseDataExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Result *bool                               `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TaskWriteoffVideoResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffVideoResponseData) GoString() string {
	return s.String()
}

func (s *TaskWriteoffVideoResponseData) SetExtra(v *TaskWriteoffVideoResponseDataExtra) *TaskWriteoffVideoResponseData {
	s.Extra = v
	return s
}

func (s *TaskWriteoffVideoResponseData) SetResult(v bool) *TaskWriteoffVideoResponseData {
	s.Result = &v
	return s
}

type TaskWriteoffVideoResponseDataExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TaskWriteoffVideoResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s TaskWriteoffVideoResponseDataExtra) GoString() string {
	return s.String()
}

func (s *TaskWriteoffVideoResponseDataExtra) SetNow(v int64) *TaskWriteoffVideoResponseDataExtra {
	s.Now = &v
	return s
}

func (s *TaskWriteoffVideoResponseDataExtra) SetErrorCode(v int32) *TaskWriteoffVideoResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *TaskWriteoffVideoResponseDataExtra) SetDescription(v string) *TaskWriteoffVideoResponseDataExtra {
	s.Description = &v
	return s
}

func (s *TaskWriteoffVideoResponseDataExtra) SetSubErrorCode(v int32) *TaskWriteoffVideoResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TaskWriteoffVideoResponseDataExtra) SetSubDescription(v string) *TaskWriteoffVideoResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *TaskWriteoffVideoResponseDataExtra) SetLogid(v string) *TaskWriteoffVideoResponseDataExtra {
	s.Logid = &v
	return s
}

type TaskboxAddRoomTaskRequest struct {
	TaskIcon        *string            `json:"task_icon,omitempty" xml:"task_icon,omitempty" require:"true"`
	StartPage       *string            `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
	ContentTag      *string            `json:"content_tag,omitempty" xml:"content_tag,omitempty" require:"true"`
	TaskDesc        *string            `json:"task_desc,omitempty" xml:"task_desc,omitempty" require:"true"`
	RoomTitle       *string            `json:"room_title,omitempty" xml:"room_title,omitempty" require:"true"`
	TaskEndTime     *int64             `json:"task_end_time,omitempty" xml:"task_end_time,omitempty" require:"true"`
	ReferMaCaptures []*string          `json:"refer_ma_captures,omitempty" xml:"refer_ma_captures,omitempty" require:"true" type:"Repeated"`
	ReferVideos     []*string          `json:"refer_videos,omitempty" xml:"refer_videos,omitempty" type:"Repeated"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	TaskSettleType  *int32             `json:"task_settle_type,omitempty" xml:"task_settle_type,omitempty" require:"true"`
	FormTag         *string            `json:"form_tag,omitempty" xml:"form_tag,omitempty" require:"true"`
	TaskName        *string            `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
}

func (s TaskboxAddRoomTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddRoomTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskboxAddRoomTaskRequest) SetTaskIcon(v string) *TaskboxAddRoomTaskRequest {
	s.TaskIcon = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetStartPage(v string) *TaskboxAddRoomTaskRequest {
	s.StartPage = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetContentTag(v string) *TaskboxAddRoomTaskRequest {
	s.ContentTag = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetTaskDesc(v string) *TaskboxAddRoomTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetRoomTitle(v string) *TaskboxAddRoomTaskRequest {
	s.RoomTitle = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetTaskEndTime(v int64) *TaskboxAddRoomTaskRequest {
	s.TaskEndTime = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetReferMaCaptures(v []*string) *TaskboxAddRoomTaskRequest {
	s.ReferMaCaptures = v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetReferVideos(v []*string) *TaskboxAddRoomTaskRequest {
	s.ReferVideos = v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetAccessToken(v string) *TaskboxAddRoomTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetHeader(v map[string]*string) *TaskboxAddRoomTaskRequest {
	s.Header = v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetTaskSettleType(v int32) *TaskboxAddRoomTaskRequest {
	s.TaskSettleType = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetFormTag(v string) *TaskboxAddRoomTaskRequest {
	s.FormTag = &v
	return s
}

func (s *TaskboxAddRoomTaskRequest) SetTaskName(v string) *TaskboxAddRoomTaskRequest {
	s.TaskName = &v
	return s
}

type TaskboxAddRoomTaskResponse struct {
	Data   *TaskboxAddRoomTaskResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                          `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                         `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                         `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s TaskboxAddRoomTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddRoomTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskboxAddRoomTaskResponse) SetData(v *TaskboxAddRoomTaskResponseData) *TaskboxAddRoomTaskResponse {
	s.Data = v
	return s
}

func (s *TaskboxAddRoomTaskResponse) SetErrNo(v int32) *TaskboxAddRoomTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxAddRoomTaskResponse) SetErrMsg(v string) *TaskboxAddRoomTaskResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskboxAddRoomTaskResponse) SetLogId(v string) *TaskboxAddRoomTaskResponse {
	s.LogId = &v
	return s
}

type TaskboxAddRoomTaskResponseData struct {
	TaskUrl *string `json:"task_url,omitempty" xml:"task_url,omitempty" require:"true"`
	TaskId  *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s TaskboxAddRoomTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddRoomTaskResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxAddRoomTaskResponseData) SetTaskUrl(v string) *TaskboxAddRoomTaskResponseData {
	s.TaskUrl = &v
	return s
}

func (s *TaskboxAddRoomTaskResponseData) SetTaskId(v int64) *TaskboxAddRoomTaskResponseData {
	s.TaskId = &v
	return s
}

type TaskboxAddTaskRequest struct {
	TaskStartTime                 *int64             `json:"task_start_time,omitempty" xml:"task_start_time,omitempty" require:"true"`
	MixPaymentAllocateRatio       map[int32]*string  `json:"mix_payment_allocate_ratio,omitempty" xml:"mix_payment_allocate_ratio,omitempty"`
	AccessToken                   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TalentPaymentAllocateRatio    *int32             `json:"talent_payment_allocate_ratio,omitempty" xml:"talent_payment_allocate_ratio,omitempty"`
	PaymentAllocateRatio          *int32             `json:"payment_allocate_ratio,omitempty" xml:"payment_allocate_ratio,omitempty"`
	ReferGids                     []*int64           `json:"refer_gids,omitempty" xml:"refer_gids,omitempty" type:"Repeated"`
	TaskEndTime                   *int64             `json:"task_end_time,omitempty" xml:"task_end_time,omitempty" require:"true"`
	StartPage                     *string            `json:"start_page,omitempty" xml:"start_page,omitempty"`
	PageType                      *int               `json:"page_type,omitempty" xml:"page_type,omitempty"`
	TaskRefundPeriod              *int32             `json:"task_refund_period,omitempty" xml:"task_refund_period,omitempty"`
	TaskSettleType                *int32             `json:"task_settle_type,omitempty" xml:"task_settle_type,omitempty" require:"true"`
	ReferMaCaptures               []*string          `json:"refer_ma_captures,omitempty" xml:"refer_ma_captures,omitempty" type:"Repeated"`
	AnchorTitle                   *string            `json:"anchor_title,omitempty" xml:"anchor_title,omitempty"`
	TalentMixPaymentAllocateRatio map[int32]*string  `json:"talent_mix_payment_allocate_ratio,omitempty" xml:"talent_mix_payment_allocate_ratio,omitempty"`
	ReferVideos                   []*string          `json:"refer_videos,omitempty" xml:"refer_videos,omitempty" type:"Repeated"`
	TaskTags                      []*string          `json:"task_tags,omitempty" xml:"task_tags,omitempty" require:"true" type:"Repeated"`
	TaskName                      *string            `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	DouyinIds                     []*string          `json:"douyin_ids,omitempty" xml:"douyin_ids,omitempty" type:"Repeated"`
	TaskIcon                      *string            `json:"task_icon,omitempty" xml:"task_icon,omitempty" require:"true"`
	TaskDesc                      *string            `json:"task_desc,omitempty" xml:"task_desc,omitempty" require:"true"`
	Header                        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TaskboxAddTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskboxAddTaskRequest) SetTaskStartTime(v int64) *TaskboxAddTaskRequest {
	s.TaskStartTime = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetMixPaymentAllocateRatio(v map[int32]*string) *TaskboxAddTaskRequest {
	s.MixPaymentAllocateRatio = v
	return s
}

func (s *TaskboxAddTaskRequest) SetAccessToken(v string) *TaskboxAddTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetTalentPaymentAllocateRatio(v int32) *TaskboxAddTaskRequest {
	s.TalentPaymentAllocateRatio = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetPaymentAllocateRatio(v int32) *TaskboxAddTaskRequest {
	s.PaymentAllocateRatio = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetReferGids(v []*int64) *TaskboxAddTaskRequest {
	s.ReferGids = v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskEndTime(v int64) *TaskboxAddTaskRequest {
	s.TaskEndTime = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetStartPage(v string) *TaskboxAddTaskRequest {
	s.StartPage = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetPageType(v int) *TaskboxAddTaskRequest {
	s.PageType = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskRefundPeriod(v int32) *TaskboxAddTaskRequest {
	s.TaskRefundPeriod = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskSettleType(v int32) *TaskboxAddTaskRequest {
	s.TaskSettleType = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetReferMaCaptures(v []*string) *TaskboxAddTaskRequest {
	s.ReferMaCaptures = v
	return s
}

func (s *TaskboxAddTaskRequest) SetAnchorTitle(v string) *TaskboxAddTaskRequest {
	s.AnchorTitle = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetTalentMixPaymentAllocateRatio(v map[int32]*string) *TaskboxAddTaskRequest {
	s.TalentMixPaymentAllocateRatio = v
	return s
}

func (s *TaskboxAddTaskRequest) SetReferVideos(v []*string) *TaskboxAddTaskRequest {
	s.ReferVideos = v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskTags(v []*string) *TaskboxAddTaskRequest {
	s.TaskTags = v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskName(v string) *TaskboxAddTaskRequest {
	s.TaskName = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetDouyinIds(v []*string) *TaskboxAddTaskRequest {
	s.DouyinIds = v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskIcon(v string) *TaskboxAddTaskRequest {
	s.TaskIcon = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetTaskDesc(v string) *TaskboxAddTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *TaskboxAddTaskRequest) SetHeader(v map[string]*string) *TaskboxAddTaskRequest {
	s.Header = v
	return s
}

type TaskboxAddTaskResponse struct {
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxAddTaskResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskboxAddTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskboxAddTaskResponse) SetLogId(v string) *TaskboxAddTaskResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxAddTaskResponse) SetData(v *TaskboxAddTaskResponseData) *TaskboxAddTaskResponse {
	s.Data = v
	return s
}

func (s *TaskboxAddTaskResponse) SetErrNo(v int32) *TaskboxAddTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxAddTaskResponse) SetErrMsg(v string) *TaskboxAddTaskResponse {
	s.ErrMsg = &v
	return s
}

type TaskboxAddTaskResponseData struct {
	TaskId *int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s TaskboxAddTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxAddTaskResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxAddTaskResponseData) SetTaskId(v int64) *TaskboxAddTaskResponseData {
	s.TaskId = &v
	return s
}

type TaskboxGenAgentLinkRequest struct {
	AgentId         *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty" require:"true"`
	AppId           *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TaskId          *int64             `json:"task_id,omitempty" xml:"task_id,omitempty"`
	AgencyTalentUid *string            `json:"agency_talent_uid,omitempty" xml:"agency_talent_uid,omitempty"`
	TaskCategory    *int               `json:"task_category,omitempty" xml:"task_category,omitempty"`
}

func (s TaskboxGenAgentLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxGenAgentLinkRequest) GoString() string {
	return s.String()
}

func (s *TaskboxGenAgentLinkRequest) SetAgentId(v int64) *TaskboxGenAgentLinkRequest {
	s.AgentId = &v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetAppId(v string) *TaskboxGenAgentLinkRequest {
	s.AppId = &v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetHeader(v map[string]*string) *TaskboxGenAgentLinkRequest {
	s.Header = v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetAccessToken(v string) *TaskboxGenAgentLinkRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetTaskId(v int64) *TaskboxGenAgentLinkRequest {
	s.TaskId = &v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetAgencyTalentUid(v string) *TaskboxGenAgentLinkRequest {
	s.AgencyTalentUid = &v
	return s
}

func (s *TaskboxGenAgentLinkRequest) SetTaskCategory(v int) *TaskboxGenAgentLinkRequest {
	s.TaskCategory = &v
	return s
}

type TaskboxGenAgentLinkResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxGenAgentLinkResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskboxGenAgentLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxGenAgentLinkResponse) GoString() string {
	return s.String()
}

func (s *TaskboxGenAgentLinkResponse) SetLogId(v string) *TaskboxGenAgentLinkResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxGenAgentLinkResponse) SetData(v *TaskboxGenAgentLinkResponseData) *TaskboxGenAgentLinkResponse {
	s.Data = v
	return s
}

func (s *TaskboxGenAgentLinkResponse) SetErrNo(v int32) *TaskboxGenAgentLinkResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxGenAgentLinkResponse) SetErrMsg(v string) *TaskboxGenAgentLinkResponse {
	s.ErrMsg = &v
	return s
}

type TaskboxGenAgentLinkResponseData struct {
	WebLink *string `json:"web_link,omitempty" xml:"web_link,omitempty" require:"true"`
	AppLink *string `json:"app_link,omitempty" xml:"app_link,omitempty" require:"true"`
}

func (s TaskboxGenAgentLinkResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxGenAgentLinkResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxGenAgentLinkResponseData) SetWebLink(v string) *TaskboxGenAgentLinkResponseData {
	s.WebLink = &v
	return s
}

func (s *TaskboxGenAgentLinkResponseData) SetAppLink(v string) *TaskboxGenAgentLinkResponseData {
	s.AppLink = &v
	return s
}

type TaskboxQueryAppTaskIdRequest struct {
	TaskCategory    *int               `json:"task_category,omitempty" xml:"task_category,omitempty"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CreateStartTime *int64             `json:"create_start_time,omitempty" xml:"create_start_time,omitempty" require:"true"`
	CreateEndTime   *int64             `json:"create_end_time,omitempty" xml:"create_end_time,omitempty" require:"true"`
}

func (s TaskboxQueryAppTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryAppTaskIdRequest) GoString() string {
	return s.String()
}

func (s *TaskboxQueryAppTaskIdRequest) SetTaskCategory(v int) *TaskboxQueryAppTaskIdRequest {
	s.TaskCategory = &v
	return s
}

func (s *TaskboxQueryAppTaskIdRequest) SetHeader(v map[string]*string) *TaskboxQueryAppTaskIdRequest {
	s.Header = v
	return s
}

func (s *TaskboxQueryAppTaskIdRequest) SetAccessToken(v string) *TaskboxQueryAppTaskIdRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxQueryAppTaskIdRequest) SetCreateStartTime(v int64) *TaskboxQueryAppTaskIdRequest {
	s.CreateStartTime = &v
	return s
}

func (s *TaskboxQueryAppTaskIdRequest) SetCreateEndTime(v int64) *TaskboxQueryAppTaskIdRequest {
	s.CreateEndTime = &v
	return s
}

type TaskboxQueryAppTaskIdResponse struct {
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxQueryAppTaskIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s TaskboxQueryAppTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryAppTaskIdResponse) GoString() string {
	return s.String()
}

func (s *TaskboxQueryAppTaskIdResponse) SetErrMsg(v string) *TaskboxQueryAppTaskIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskboxQueryAppTaskIdResponse) SetLogId(v string) *TaskboxQueryAppTaskIdResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxQueryAppTaskIdResponse) SetData(v *TaskboxQueryAppTaskIdResponseData) *TaskboxQueryAppTaskIdResponse {
	s.Data = v
	return s
}

func (s *TaskboxQueryAppTaskIdResponse) SetErrNo(v int32) *TaskboxQueryAppTaskIdResponse {
	s.ErrNo = &v
	return s
}

type TaskboxQueryAppTaskIdResponseData struct {
	TaskIds []*int64 `json:"task_ids,omitempty" xml:"task_ids,omitempty" require:"true" type:"Repeated"`
}

func (s TaskboxQueryAppTaskIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryAppTaskIdResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxQueryAppTaskIdResponseData) SetTaskIds(v []*int64) *TaskboxQueryAppTaskIdResponseData {
	s.TaskIds = v
	return s
}

type TaskboxQueryBillLinkRequest struct {
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskboxQueryBillLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryBillLinkRequest) GoString() string {
	return s.String()
}

func (s *TaskboxQueryBillLinkRequest) SetBillDate(v string) *TaskboxQueryBillLinkRequest {
	s.BillDate = &v
	return s
}

func (s *TaskboxQueryBillLinkRequest) SetHeader(v map[string]*string) *TaskboxQueryBillLinkRequest {
	s.Header = v
	return s
}

func (s *TaskboxQueryBillLinkRequest) SetAccessToken(v string) *TaskboxQueryBillLinkRequest {
	s.AccessToken = &v
	return s
}

type TaskboxQueryBillLinkResponse struct {
	Data   *TaskboxQueryBillLinkResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s TaskboxQueryBillLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryBillLinkResponse) GoString() string {
	return s.String()
}

func (s *TaskboxQueryBillLinkResponse) SetData(v *TaskboxQueryBillLinkResponseData) *TaskboxQueryBillLinkResponse {
	s.Data = v
	return s
}

func (s *TaskboxQueryBillLinkResponse) SetErrNo(v int32) *TaskboxQueryBillLinkResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxQueryBillLinkResponse) SetErrMsg(v string) *TaskboxQueryBillLinkResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskboxQueryBillLinkResponse) SetLogId(v string) *TaskboxQueryBillLinkResponse {
	s.LogId = &v
	return s
}

type TaskboxQueryBillLinkResponseData struct {
	LiveBillLink *string `json:"live_bill_link,omitempty" xml:"live_bill_link,omitempty"`
	BillLink     *string `json:"bill_link,omitempty" xml:"bill_link,omitempty"`
}

func (s TaskboxQueryBillLinkResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryBillLinkResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxQueryBillLinkResponseData) SetLiveBillLink(v string) *TaskboxQueryBillLinkResponseData {
	s.LiveBillLink = &v
	return s
}

func (s *TaskboxQueryBillLinkResponseData) SetBillLink(v string) *TaskboxQueryBillLinkResponseData {
	s.BillLink = &v
	return s
}

type TaskboxQueryTaskInfoRequest struct {
	Appid              *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Header             map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	QueryParamsType    *int               `json:"query_params_type,omitempty" xml:"query_params_type,omitempty" require:"true"`
	QueryParamsContent *string            `json:"query_params_content,omitempty" xml:"query_params_content,omitempty" require:"true"`
	PageNo             *int32             `json:"page_no,omitempty" xml:"page_no,omitempty" require:"true"`
	PageSize           *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	TaskCategory       *int               `json:"task_category,omitempty" xml:"task_category,omitempty"`
}

func (s TaskboxQueryTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *TaskboxQueryTaskInfoRequest) SetAppid(v string) *TaskboxQueryTaskInfoRequest {
	s.Appid = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetHeader(v map[string]*string) *TaskboxQueryTaskInfoRequest {
	s.Header = v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetAccessToken(v string) *TaskboxQueryTaskInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetQueryParamsType(v int) *TaskboxQueryTaskInfoRequest {
	s.QueryParamsType = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetQueryParamsContent(v string) *TaskboxQueryTaskInfoRequest {
	s.QueryParamsContent = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetPageNo(v int32) *TaskboxQueryTaskInfoRequest {
	s.PageNo = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetPageSize(v int32) *TaskboxQueryTaskInfoRequest {
	s.PageSize = &v
	return s
}

func (s *TaskboxQueryTaskInfoRequest) SetTaskCategory(v int) *TaskboxQueryTaskInfoRequest {
	s.TaskCategory = &v
	return s
}

type TaskboxQueryTaskInfoResponse struct {
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxQueryTaskInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskboxQueryTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *TaskboxQueryTaskInfoResponse) SetLogId(v string) *TaskboxQueryTaskInfoResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponse) SetData(v *TaskboxQueryTaskInfoResponseData) *TaskboxQueryTaskInfoResponse {
	s.Data = v
	return s
}

func (s *TaskboxQueryTaskInfoResponse) SetErrNo(v int32) *TaskboxQueryTaskInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponse) SetErrMsg(v string) *TaskboxQueryTaskInfoResponse {
	s.ErrMsg = &v
	return s
}

type TaskboxQueryTaskInfoResponseData struct {
	Tasks     []*TaskboxQueryTaskInfoResponseDataTasksItem `json:"tasks,omitempty" xml:"tasks,omitempty" require:"true" type:"Repeated"`
	PageCount *int64                                       `json:"page_count,omitempty" xml:"page_count,omitempty" require:"true"`
	Total     *int64                                       `json:"total,omitempty" xml:"total,omitempty" require:"true"`
	AppId     *string                                      `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s TaskboxQueryTaskInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryTaskInfoResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxQueryTaskInfoResponseData) SetTasks(v []*TaskboxQueryTaskInfoResponseDataTasksItem) *TaskboxQueryTaskInfoResponseData {
	s.Tasks = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseData) SetPageCount(v int64) *TaskboxQueryTaskInfoResponseData {
	s.PageCount = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseData) SetTotal(v int64) *TaskboxQueryTaskInfoResponseData {
	s.Total = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseData) SetAppId(v string) *TaskboxQueryTaskInfoResponseData {
	s.AppId = &v
	return s
}

type TaskboxQueryTaskInfoResponseDataTasksItem struct {
	ReferVideoCaptures            []*string                                                             `json:"refer_video_captures,omitempty" xml:"refer_video_captures,omitempty" type:"Repeated"`
	TalentMixPaymentAllocateRatio map[int32]*string                                                     `json:"talent_mix_payment_allocate_ratio,omitempty" xml:"talent_mix_payment_allocate_ratio,omitempty"`
	PlatformAddressApp            *string                                                               `json:"platform_address_app,omitempty" xml:"platform_address_app,omitempty"`
	StartPage                     *string                                                               `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
	PaymentAllocateRatio          *int32                                                                `json:"payment_allocate_ratio,omitempty" xml:"payment_allocate_ratio,omitempty"`
	CapitalAccount                *int                                                                  `json:"capital_account,omitempty" xml:"capital_account,omitempty"`
	TaskSettleType                *int32                                                                `json:"task_settle_type,omitempty" xml:"task_settle_type,omitempty" require:"true"`
	ReferVideos                   []*string                                                             `json:"refer_videos,omitempty" xml:"refer_videos,omitempty" type:"Repeated"`
	ReferGids                     []*int64                                                              `json:"refer_gids,omitempty" xml:"refer_gids,omitempty" type:"Repeated"`
	ReferMaCaptures               []*string                                                             `json:"refer_ma_captures,omitempty" xml:"refer_ma_captures,omitempty" type:"Repeated"`
	PageType                      *int                                                                  `json:"page_type,omitempty" xml:"page_type,omitempty" require:"true"`
	TaskId                        *int64                                                                `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	OrientedTalentRelList         []*TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem `json:"oriented_talent_rel_list,omitempty" xml:"oriented_talent_rel_list,omitempty" type:"Repeated"`
	TaskName                      *string                                                               `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	RejectReason                  *string                                                               `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	TaskRefundPeriod              *int32                                                                `json:"task_refund_period,omitempty" xml:"task_refund_period,omitempty"`
	TaskIcon                      *string                                                               `json:"task_icon,omitempty" xml:"task_icon,omitempty" require:"true"`
	TaskDesc                      *string                                                               `json:"task_desc,omitempty" xml:"task_desc,omitempty" require:"true"`
	PlatformAddressWeb            *string                                                               `json:"platform_address_web,omitempty" xml:"platform_address_web,omitempty"`
	AnchorTitle                   *string                                                               `json:"anchor_title,omitempty" xml:"anchor_title,omitempty" require:"true"`
	Appid                         *string                                                               `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	TaskType                      *int                                                                  `json:"task_type,omitempty" xml:"task_type,omitempty" require:"true"`
	Status                        *int32                                                                `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	TaskEndTime                   *int64                                                                `json:"task_end_time,omitempty" xml:"task_end_time,omitempty" require:"true"`
	TaskStartTime                 *int64                                                                `json:"task_start_time,omitempty" xml:"task_start_time,omitempty" require:"true"`
	TalentPaymentAllocateRatio    *int32                                                                `json:"talent_payment_allocate_ratio,omitempty" xml:"talent_payment_allocate_ratio,omitempty"`
	TaskTags                      []*string                                                             `json:"task_tags,omitempty" xml:"task_tags,omitempty" require:"true" type:"Repeated"`
}

func (s TaskboxQueryTaskInfoResponseDataTasksItem) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryTaskInfoResponseDataTasksItem) GoString() string {
	return s.String()
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetReferVideoCaptures(v []*string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferVideoCaptures = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTalentMixPaymentAllocateRatio(v map[int32]*string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TalentMixPaymentAllocateRatio = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetPlatformAddressApp(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PlatformAddressApp = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetStartPage(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.StartPage = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetPaymentAllocateRatio(v int32) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PaymentAllocateRatio = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetCapitalAccount(v int) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.CapitalAccount = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskSettleType(v int32) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskSettleType = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetReferVideos(v []*string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferVideos = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetReferGids(v []*int64) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferGids = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetReferMaCaptures(v []*string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.ReferMaCaptures = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetPageType(v int) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PageType = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskId(v int64) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskId = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetOrientedTalentRelList(v []*TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.OrientedTalentRelList = v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskName(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskName = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetRejectReason(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.RejectReason = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskRefundPeriod(v int32) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskRefundPeriod = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskIcon(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskIcon = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskDesc(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskDesc = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetPlatformAddressWeb(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.PlatformAddressWeb = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetAnchorTitle(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.AnchorTitle = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetAppid(v string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.Appid = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskType(v int) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskType = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetStatus(v int32) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.Status = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskEndTime(v int64) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskEndTime = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskStartTime(v int64) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskStartTime = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTalentPaymentAllocateRatio(v int32) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TalentPaymentAllocateRatio = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItem) SetTaskTags(v []*string) *TaskboxQueryTaskInfoResponseDataTasksItem {
	s.TaskTags = v
	return s
}

type TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem struct {
	CancelOperator   *int    `json:"cancel_operator,omitempty" xml:"cancel_operator,omitempty"`
	DouyinId         *string `json:"douyin_id,omitempty" xml:"douyin_id,omitempty" require:"true"`
	CooperationState *int    `json:"cooperation_state,omitempty" xml:"cooperation_state,omitempty" require:"true"`
}

func (s TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) String() string {
	return tea.Prettify(s)
}

func (s TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) GoString() string {
	return s.String()
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetCancelOperator(v int) *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.CancelOperator = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetDouyinId(v string) *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.DouyinId = &v
	return s
}

func (s *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem) SetCooperationState(v int) *TaskboxQueryTaskInfoResponseDataTasksItemOrientedTalentRelListItem {
	s.CooperationState = &v
	return s
}

type TaskboxSaveAgentRequest struct {
	AgentId       *int64             `json:"agent_id,omitempty" xml:"agent_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AgentNickname *string            `json:"agent_nickname,omitempty" xml:"agent_nickname,omitempty" require:"true"`
}

func (s TaskboxSaveAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxSaveAgentRequest) GoString() string {
	return s.String()
}

func (s *TaskboxSaveAgentRequest) SetAgentId(v int64) *TaskboxSaveAgentRequest {
	s.AgentId = &v
	return s
}

func (s *TaskboxSaveAgentRequest) SetHeader(v map[string]*string) *TaskboxSaveAgentRequest {
	s.Header = v
	return s
}

func (s *TaskboxSaveAgentRequest) SetAccessToken(v string) *TaskboxSaveAgentRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxSaveAgentRequest) SetAgentNickname(v string) *TaskboxSaveAgentRequest {
	s.AgentNickname = &v
	return s
}

type TaskboxSaveAgentResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxSaveAgentResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskboxSaveAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxSaveAgentResponse) GoString() string {
	return s.String()
}

func (s *TaskboxSaveAgentResponse) SetLogId(v string) *TaskboxSaveAgentResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxSaveAgentResponse) SetData(v *TaskboxSaveAgentResponseData) *TaskboxSaveAgentResponse {
	s.Data = v
	return s
}

func (s *TaskboxSaveAgentResponse) SetErrNo(v int32) *TaskboxSaveAgentResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxSaveAgentResponse) SetErrMsg(v string) *TaskboxSaveAgentResponse {
	s.ErrMsg = &v
	return s
}

type TaskboxSaveAgentResponseData struct {
	AgentId *int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty" require:"true"`
}

func (s TaskboxSaveAgentResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxSaveAgentResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxSaveAgentResponseData) SetAgentId(v int64) *TaskboxSaveAgentResponseData {
	s.AgentId = &v
	return s
}

type TaskboxUpdateStatusRequest struct {
	TaskStatus  *int64             `json:"task_status,omitempty" xml:"task_status,omitempty" require:"true"`
	TaskId      *int64             `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TaskboxUpdateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateStatusRequest) SetTaskStatus(v int64) *TaskboxUpdateStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *TaskboxUpdateStatusRequest) SetTaskId(v int64) *TaskboxUpdateStatusRequest {
	s.TaskId = &v
	return s
}

func (s *TaskboxUpdateStatusRequest) SetHeader(v map[string]*string) *TaskboxUpdateStatusRequest {
	s.Header = v
	return s
}

func (s *TaskboxUpdateStatusRequest) SetAccessToken(v string) *TaskboxUpdateStatusRequest {
	s.AccessToken = &v
	return s
}

type TaskboxUpdateStatusResponse struct {
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxUpdateStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s TaskboxUpdateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateStatusResponse) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateStatusResponse) SetLogId(v string) *TaskboxUpdateStatusResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxUpdateStatusResponse) SetData(v *TaskboxUpdateStatusResponseData) *TaskboxUpdateStatusResponse {
	s.Data = v
	return s
}

func (s *TaskboxUpdateStatusResponse) SetErrNo(v int32) *TaskboxUpdateStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxUpdateStatusResponse) SetErrMsg(v string) *TaskboxUpdateStatusResponse {
	s.ErrMsg = &v
	return s
}

type TaskboxUpdateStatusResponseData struct {
	TaskId *int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s TaskboxUpdateStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateStatusResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateStatusResponseData) SetTaskId(v int64) *TaskboxUpdateStatusResponseData {
	s.TaskId = &v
	return s
}

type TaskboxUpdateTaskRequest struct {
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	PageType        *int               `json:"page_type,omitempty" xml:"page_type,omitempty"`
	TaskTags        []*string          `json:"task_tags,omitempty" xml:"task_tags,omitempty" require:"true" type:"Repeated"`
	ReferVideos     []*string          `json:"refer_videos,omitempty" xml:"refer_videos,omitempty" type:"Repeated"`
	TaskStartTime   *int64             `json:"task_start_time,omitempty" xml:"task_start_time,omitempty" require:"true"`
	TaskName        *string            `json:"task_name,omitempty" xml:"task_name,omitempty" require:"true"`
	AnchorTitle     *string            `json:"anchor_title,omitempty" xml:"anchor_title,omitempty"`
	ReferGids       []*int64           `json:"refer_gids,omitempty" xml:"refer_gids,omitempty" type:"Repeated"`
	TaskId          *int64             `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	TaskIcon        *string            `json:"task_icon,omitempty" xml:"task_icon,omitempty" require:"true"`
	TaskSettleType  *int32             `json:"task_settle_type,omitempty" xml:"task_settle_type,omitempty" require:"true"`
	ReferMaCaptures []*string          `json:"refer_ma_captures,omitempty" xml:"refer_ma_captures,omitempty" type:"Repeated"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	StartPage       *string            `json:"start_page,omitempty" xml:"start_page,omitempty" require:"true"`
	TaskDesc        *string            `json:"task_desc,omitempty" xml:"task_desc,omitempty" require:"true"`
	TaskEndTime     *int64             `json:"task_end_time,omitempty" xml:"task_end_time,omitempty" require:"true"`
}

func (s TaskboxUpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateTaskRequest) SetHeader(v map[string]*string) *TaskboxUpdateTaskRequest {
	s.Header = v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetPageType(v int) *TaskboxUpdateTaskRequest {
	s.PageType = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskTags(v []*string) *TaskboxUpdateTaskRequest {
	s.TaskTags = v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetReferVideos(v []*string) *TaskboxUpdateTaskRequest {
	s.ReferVideos = v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskStartTime(v int64) *TaskboxUpdateTaskRequest {
	s.TaskStartTime = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskName(v string) *TaskboxUpdateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetAnchorTitle(v string) *TaskboxUpdateTaskRequest {
	s.AnchorTitle = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetReferGids(v []*int64) *TaskboxUpdateTaskRequest {
	s.ReferGids = v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskId(v int64) *TaskboxUpdateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskIcon(v string) *TaskboxUpdateTaskRequest {
	s.TaskIcon = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskSettleType(v int32) *TaskboxUpdateTaskRequest {
	s.TaskSettleType = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetReferMaCaptures(v []*string) *TaskboxUpdateTaskRequest {
	s.ReferMaCaptures = v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetAccessToken(v string) *TaskboxUpdateTaskRequest {
	s.AccessToken = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetStartPage(v string) *TaskboxUpdateTaskRequest {
	s.StartPage = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskDesc(v string) *TaskboxUpdateTaskRequest {
	s.TaskDesc = &v
	return s
}

func (s *TaskboxUpdateTaskRequest) SetTaskEndTime(v int64) *TaskboxUpdateTaskRequest {
	s.TaskEndTime = &v
	return s
}

type TaskboxUpdateTaskResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TaskboxUpdateTaskResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TaskboxUpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateTaskResponse) SetErrNo(v int32) *TaskboxUpdateTaskResponse {
	s.ErrNo = &v
	return s
}

func (s *TaskboxUpdateTaskResponse) SetErrMsg(v string) *TaskboxUpdateTaskResponse {
	s.ErrMsg = &v
	return s
}

func (s *TaskboxUpdateTaskResponse) SetLogId(v string) *TaskboxUpdateTaskResponse {
	s.LogId = &v
	return s
}

func (s *TaskboxUpdateTaskResponse) SetData(v *TaskboxUpdateTaskResponseData) *TaskboxUpdateTaskResponse {
	s.Data = v
	return s
}

type TaskboxUpdateTaskResponseData struct {
	TaskId *int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s TaskboxUpdateTaskResponseData) String() string {
	return tea.Prettify(s)
}

func (s TaskboxUpdateTaskResponseData) GoString() string {
	return s.String()
}

func (s *TaskboxUpdateTaskResponseData) SetTaskId(v int64) *TaskboxUpdateTaskResponseData {
	s.TaskId = &v
	return s
}

type TemplateGetRequest struct {
	AccessToken       *string                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
	GoodsTemplateType *int                    `json:"goods_template_type,omitempty" xml:"goods_template_type,omitempty"`
	ProductType       *int                    `json:"product_type,omitempty" xml:"product_type,omitempty"`
	Base              *TemplateGetRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId         *string                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	CategoryId        *string                 `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	Header            map[string]*string      `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TemplateGetRequest) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetRequest) GoString() string {
	return s.String()
}

func (s *TemplateGetRequest) SetAccessToken(v string) *TemplateGetRequest {
	s.AccessToken = &v
	return s
}

func (s *TemplateGetRequest) SetGoodsTemplateType(v int) *TemplateGetRequest {
	s.GoodsTemplateType = &v
	return s
}

func (s *TemplateGetRequest) SetProductType(v int) *TemplateGetRequest {
	s.ProductType = &v
	return s
}

func (s *TemplateGetRequest) SetBase(v *TemplateGetRequestBase) *TemplateGetRequest {
	s.Base = v
	return s
}

func (s *TemplateGetRequest) SetAccountId(v string) *TemplateGetRequest {
	s.AccountId = &v
	return s
}

func (s *TemplateGetRequest) SetCategoryId(v string) *TemplateGetRequest {
	s.CategoryId = &v
	return s
}

func (s *TemplateGetRequest) SetHeader(v map[string]*string) *TemplateGetRequest {
	s.Header = v
	return s
}

type TemplateGetRequestBase struct {
	Caller     *string                           `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                           `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                           `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *TemplateGetRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                           `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s TemplateGetRequestBase) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetRequestBase) GoString() string {
	return s.String()
}

func (s *TemplateGetRequestBase) SetCaller(v string) *TemplateGetRequestBase {
	s.Caller = &v
	return s
}

func (s *TemplateGetRequestBase) SetClient(v string) *TemplateGetRequestBase {
	s.Client = &v
	return s
}

func (s *TemplateGetRequestBase) SetExtra(v map[string]*string) *TemplateGetRequestBase {
	s.Extra = v
	return s
}

func (s *TemplateGetRequestBase) SetLogID(v string) *TemplateGetRequestBase {
	s.LogID = &v
	return s
}

func (s *TemplateGetRequestBase) SetTrafficEnv(v *TemplateGetRequestBaseTrafficEnv) *TemplateGetRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *TemplateGetRequestBase) SetAddr(v string) *TemplateGetRequestBase {
	s.Addr = &v
	return s
}

type TemplateGetRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s TemplateGetRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *TemplateGetRequestBaseTrafficEnv) SetEnv(v string) *TemplateGetRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *TemplateGetRequestBaseTrafficEnv) SetOpen(v bool) *TemplateGetRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type TemplateGetResponse struct {
	Extra    *TemplateGetResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *TemplateGetResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *TemplateGetResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TemplateGetResponse) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponse) GoString() string {
	return s.String()
}

func (s *TemplateGetResponse) SetExtra(v *TemplateGetResponseExtra) *TemplateGetResponse {
	s.Extra = v
	return s
}

func (s *TemplateGetResponse) SetBaseResp(v *TemplateGetResponseBaseResp) *TemplateGetResponse {
	s.BaseResp = v
	return s
}

func (s *TemplateGetResponse) SetData(v *TemplateGetResponseData) *TemplateGetResponse {
	s.Data = v
	return s
}

type TemplateGetResponseBaseResp struct {
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s TemplateGetResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponseBaseResp) GoString() string {
	return s.String()
}

func (s *TemplateGetResponseBaseResp) SetStatusMessage(v string) *TemplateGetResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *TemplateGetResponseBaseResp) SetExtra(v map[string]*string) *TemplateGetResponseBaseResp {
	s.Extra = v
	return s
}

func (s *TemplateGetResponseBaseResp) SetStatusCode(v int32) *TemplateGetResponseBaseResp {
	s.StatusCode = &v
	return s
}

type TemplateGetResponseData struct {
	Description  *string                                    `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode    *int32                                     `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	ProductAttrs []*TemplateGetResponseDataProductAttrsItem `json:"product_attrs,omitempty" xml:"product_attrs,omitempty" type:"Repeated"`
	SkuAttrs     []*TemplateGetResponseDataSkuAttrsItem     `json:"sku_attrs,omitempty" xml:"sku_attrs,omitempty" type:"Repeated"`
}

func (s TemplateGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponseData) GoString() string {
	return s.String()
}

func (s *TemplateGetResponseData) SetDescription(v string) *TemplateGetResponseData {
	s.Description = &v
	return s
}

func (s *TemplateGetResponseData) SetErrorCode(v int32) *TemplateGetResponseData {
	s.ErrorCode = &v
	return s
}

func (s *TemplateGetResponseData) SetProductAttrs(v []*TemplateGetResponseDataProductAttrsItem) *TemplateGetResponseData {
	s.ProductAttrs = v
	return s
}

func (s *TemplateGetResponseData) SetSkuAttrs(v []*TemplateGetResponseDataSkuAttrsItem) *TemplateGetResponseData {
	s.SkuAttrs = v
	return s
}

type TemplateGetResponseDataProductAttrsItem struct {
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	IsMulti    *bool   `json:"is_multi,omitempty" xml:"is_multi,omitempty"`
	IsRequired *bool   `json:"is_required,omitempty" xml:"is_required,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ValueDemo  *string `json:"value_demo,omitempty" xml:"value_demo,omitempty"`
	ValueType  *string `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s TemplateGetResponseDataProductAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponseDataProductAttrsItem) GoString() string {
	return s.String()
}

func (s *TemplateGetResponseDataProductAttrsItem) SetDesc(v string) *TemplateGetResponseDataProductAttrsItem {
	s.Desc = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetIsMulti(v bool) *TemplateGetResponseDataProductAttrsItem {
	s.IsMulti = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetIsRequired(v bool) *TemplateGetResponseDataProductAttrsItem {
	s.IsRequired = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetKey(v string) *TemplateGetResponseDataProductAttrsItem {
	s.Key = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetName(v string) *TemplateGetResponseDataProductAttrsItem {
	s.Name = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetValueDemo(v string) *TemplateGetResponseDataProductAttrsItem {
	s.ValueDemo = &v
	return s
}

func (s *TemplateGetResponseDataProductAttrsItem) SetValueType(v string) *TemplateGetResponseDataProductAttrsItem {
	s.ValueType = &v
	return s
}

type TemplateGetResponseDataSkuAttrsItem struct {
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	IsMulti    *bool   `json:"is_multi,omitempty" xml:"is_multi,omitempty"`
	IsRequired *bool   `json:"is_required,omitempty" xml:"is_required,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ValueDemo  *string `json:"value_demo,omitempty" xml:"value_demo,omitempty"`
	ValueType  *string `json:"value_type,omitempty" xml:"value_type,omitempty"`
}

func (s TemplateGetResponseDataSkuAttrsItem) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponseDataSkuAttrsItem) GoString() string {
	return s.String()
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetDesc(v string) *TemplateGetResponseDataSkuAttrsItem {
	s.Desc = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetIsMulti(v bool) *TemplateGetResponseDataSkuAttrsItem {
	s.IsMulti = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetIsRequired(v bool) *TemplateGetResponseDataSkuAttrsItem {
	s.IsRequired = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetKey(v string) *TemplateGetResponseDataSkuAttrsItem {
	s.Key = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetName(v string) *TemplateGetResponseDataSkuAttrsItem {
	s.Name = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetValueDemo(v string) *TemplateGetResponseDataSkuAttrsItem {
	s.ValueDemo = &v
	return s
}

func (s *TemplateGetResponseDataSkuAttrsItem) SetValueType(v string) *TemplateGetResponseDataSkuAttrsItem {
	s.ValueType = &v
	return s
}

type TemplateGetResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TemplateGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TemplateGetResponseExtra) GoString() string {
	return s.String()
}

func (s *TemplateGetResponseExtra) SetErrorCode(v int32) *TemplateGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TemplateGetResponseExtra) SetLogid(v string) *TemplateGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *TemplateGetResponseExtra) SetNow(v int64) *TemplateGetResponseExtra {
	s.Now = &v
	return s
}

func (s *TemplateGetResponseExtra) SetSubDescription(v string) *TemplateGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TemplateGetResponseExtra) SetSubErrorCode(v int32) *TemplateGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TemplateGetResponseExtra) SetDescription(v string) *TemplateGetResponseExtra {
	s.Description = &v
	return s
}

type TemplateListRequest struct {
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PageNumber  *int64                   `json:"page_number,omitempty" xml:"page_number,omitempty" require:"true"`
	PageSize    *int64                   `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	Base        *TemplateListRequestBase `json:"Base,omitempty" xml:"Base,omitempty" require:"true"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s TemplateListRequest) GoString() string {
	return s.String()
}

func (s *TemplateListRequest) SetAccessToken(v string) *TemplateListRequest {
	s.AccessToken = &v
	return s
}

func (s *TemplateListRequest) SetPageNumber(v int64) *TemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *TemplateListRequest) SetPageSize(v int64) *TemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *TemplateListRequest) SetBase(v *TemplateListRequestBase) *TemplateListRequest {
	s.Base = v
	return s
}

func (s *TemplateListRequest) SetHeader(v map[string]*string) *TemplateListRequest {
	s.Header = v
	return s
}

type TemplateListRequestBase struct {
	Client     *string                            `json:"Client,omitempty" xml:"Client,omitempty"`
	TrafficEnv *TemplateListRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Extra      map[string]*string                 `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                            `json:"LogID,omitempty" xml:"LogID,omitempty"`
	Caller     *string                            `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Addr       *string                            `json:"Addr,omitempty" xml:"Addr,omitempty"`
}

func (s TemplateListRequestBase) String() string {
	return tea.Prettify(s)
}

func (s TemplateListRequestBase) GoString() string {
	return s.String()
}

func (s *TemplateListRequestBase) SetClient(v string) *TemplateListRequestBase {
	s.Client = &v
	return s
}

func (s *TemplateListRequestBase) SetTrafficEnv(v *TemplateListRequestBaseTrafficEnv) *TemplateListRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *TemplateListRequestBase) SetExtra(v map[string]*string) *TemplateListRequestBase {
	s.Extra = v
	return s
}

func (s *TemplateListRequestBase) SetLogID(v string) *TemplateListRequestBase {
	s.LogID = &v
	return s
}

func (s *TemplateListRequestBase) SetCaller(v string) *TemplateListRequestBase {
	s.Caller = &v
	return s
}

func (s *TemplateListRequestBase) SetAddr(v string) *TemplateListRequestBase {
	s.Addr = &v
	return s
}

type TemplateListRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s TemplateListRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s TemplateListRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *TemplateListRequestBaseTrafficEnv) SetEnv(v string) *TemplateListRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *TemplateListRequestBaseTrafficEnv) SetOpen(v bool) *TemplateListRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type TemplateListResponse struct {
	ErrMsg *string                   `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                   `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TemplateListResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                    `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s TemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponse) GoString() string {
	return s.String()
}

func (s *TemplateListResponse) SetErrMsg(v string) *TemplateListResponse {
	s.ErrMsg = &v
	return s
}

func (s *TemplateListResponse) SetLogId(v string) *TemplateListResponse {
	s.LogId = &v
	return s
}

func (s *TemplateListResponse) SetData(v *TemplateListResponseData) *TemplateListResponse {
	s.Data = v
	return s
}

func (s *TemplateListResponse) SetErrNo(v int32) *TemplateListResponse {
	s.ErrNo = &v
	return s
}

type TemplateListResponseData struct {
	IsCapacityOpen *int64                                      `json:"is_capacity_open,omitempty" xml:"is_capacity_open,omitempty" require:"true"`
	Extra          *TemplateListResponseDataExtra              `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp       *TemplateListResponseDataBaseResp           `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	TemplateList   []*TemplateListResponseDataTemplateListItem `json:"template_list,omitempty" xml:"template_list,omitempty" type:"Repeated"`
	TotalCount     *int64                                      `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s TemplateListResponseData) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseData) GoString() string {
	return s.String()
}

func (s *TemplateListResponseData) SetIsCapacityOpen(v int64) *TemplateListResponseData {
	s.IsCapacityOpen = &v
	return s
}

func (s *TemplateListResponseData) SetExtra(v *TemplateListResponseDataExtra) *TemplateListResponseData {
	s.Extra = v
	return s
}

func (s *TemplateListResponseData) SetBaseResp(v *TemplateListResponseDataBaseResp) *TemplateListResponseData {
	s.BaseResp = v
	return s
}

func (s *TemplateListResponseData) SetTemplateList(v []*TemplateListResponseDataTemplateListItem) *TemplateListResponseData {
	s.TemplateList = v
	return s
}

func (s *TemplateListResponseData) SetTotalCount(v int64) *TemplateListResponseData {
	s.TotalCount = &v
	return s
}

type TemplateListResponseDataBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s TemplateListResponseDataBaseResp) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseDataBaseResp) GoString() string {
	return s.String()
}

func (s *TemplateListResponseDataBaseResp) SetStatusCode(v int32) *TemplateListResponseDataBaseResp {
	s.StatusCode = &v
	return s
}

func (s *TemplateListResponseDataBaseResp) SetExtra(v map[string]*string) *TemplateListResponseDataBaseResp {
	s.Extra = v
	return s
}

func (s *TemplateListResponseDataBaseResp) SetStatusMessage(v string) *TemplateListResponseDataBaseResp {
	s.StatusMessage = &v
	return s
}

type TemplateListResponseDataExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TemplateListResponseDataExtra) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseDataExtra) GoString() string {
	return s.String()
}

func (s *TemplateListResponseDataExtra) SetSubErrorCode(v int32) *TemplateListResponseDataExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TemplateListResponseDataExtra) SetSubDescription(v string) *TemplateListResponseDataExtra {
	s.SubDescription = &v
	return s
}

func (s *TemplateListResponseDataExtra) SetLogid(v string) *TemplateListResponseDataExtra {
	s.Logid = &v
	return s
}

func (s *TemplateListResponseDataExtra) SetNow(v int64) *TemplateListResponseDataExtra {
	s.Now = &v
	return s
}

func (s *TemplateListResponseDataExtra) SetErrorCode(v int32) *TemplateListResponseDataExtra {
	s.ErrorCode = &v
	return s
}

func (s *TemplateListResponseDataExtra) SetDescription(v string) *TemplateListResponseDataExtra {
	s.Description = &v
	return s
}

type TemplateListResponseDataTemplateListItem struct {
	ReviewStatus  *int                                                  `json:"review_status,omitempty" xml:"review_status,omitempty"`
	Ctime         *int64                                                `json:"ctime,omitempty" xml:"ctime,omitempty"`
	TpType        *int                                                  `json:"tp_type,omitempty" xml:"tp_type,omitempty"`
	Uri           *string                                               `json:"uri,omitempty" xml:"uri,omitempty"`
	VideoUrl      *string                                               `json:"video_url,omitempty" xml:"video_url,omitempty"`
	ImageUrl      *string                                               `json:"image_url,omitempty" xml:"image_url,omitempty"`
	TemplateId    *int64                                                `json:"template_id,omitempty" xml:"template_id,omitempty"`
	Desc          *string                                               `json:"desc,omitempty" xml:"desc,omitempty"`
	Status        *int                                                  `json:"status,omitempty" xml:"status,omitempty"`
	Reason        []*TemplateListResponseDataTemplateListItemReasonItem `json:"reason,omitempty" xml:"reason,omitempty" type:"Repeated"`
	SlotMapLength *int32                                                `json:"slot_map_length,omitempty" xml:"slot_map_length,omitempty"`
	MaxDuration   *int64                                                `json:"max_duration,omitempty" xml:"max_duration,omitempty"`
	Title         *string                                               `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TemplateListResponseDataTemplateListItem) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseDataTemplateListItem) GoString() string {
	return s.String()
}

func (s *TemplateListResponseDataTemplateListItem) SetReviewStatus(v int) *TemplateListResponseDataTemplateListItem {
	s.ReviewStatus = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetCtime(v int64) *TemplateListResponseDataTemplateListItem {
	s.Ctime = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetTpType(v int) *TemplateListResponseDataTemplateListItem {
	s.TpType = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetUri(v string) *TemplateListResponseDataTemplateListItem {
	s.Uri = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetVideoUrl(v string) *TemplateListResponseDataTemplateListItem {
	s.VideoUrl = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetImageUrl(v string) *TemplateListResponseDataTemplateListItem {
	s.ImageUrl = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetTemplateId(v int64) *TemplateListResponseDataTemplateListItem {
	s.TemplateId = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetDesc(v string) *TemplateListResponseDataTemplateListItem {
	s.Desc = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetStatus(v int) *TemplateListResponseDataTemplateListItem {
	s.Status = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetReason(v []*TemplateListResponseDataTemplateListItemReasonItem) *TemplateListResponseDataTemplateListItem {
	s.Reason = v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetSlotMapLength(v int32) *TemplateListResponseDataTemplateListItem {
	s.SlotMapLength = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetMaxDuration(v int64) *TemplateListResponseDataTemplateListItem {
	s.MaxDuration = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItem) SetTitle(v string) *TemplateListResponseDataTemplateListItem {
	s.Title = &v
	return s
}

type TemplateListResponseDataTemplateListItemReasonItem struct {
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	AuditAdvice *string `json:"AuditAdvice,omitempty" xml:"AuditAdvice,omitempty"`
}

func (s TemplateListResponseDataTemplateListItemReasonItem) String() string {
	return tea.Prettify(s)
}

func (s TemplateListResponseDataTemplateListItemReasonItem) GoString() string {
	return s.String()
}

func (s *TemplateListResponseDataTemplateListItemReasonItem) SetAuditReason(v string) *TemplateListResponseDataTemplateListItemReasonItem {
	s.AuditReason = &v
	return s
}

func (s *TemplateListResponseDataTemplateListItemReasonItem) SetAuditAdvice(v string) *TemplateListResponseDataTemplateListItemReasonItem {
	s.AuditAdvice = &v
	return s
}

type TicketCalendarBatchSaveRequest struct {
	AccessToken          *string                                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	TicketSkuAndCalendar []*TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem `json:"ticket_sku_and_calendar,omitempty" xml:"ticket_sku_and_calendar,omitempty" require:"true" type:"Repeated"`
	AccountId            *string                                                   `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	ProductId            *string                                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductOutId         *string                                                   `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	Header               map[string]*string                                        `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TicketCalendarBatchSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequest) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequest) SetAccessToken(v string) *TicketCalendarBatchSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *TicketCalendarBatchSaveRequest) SetTicketSkuAndCalendar(v []*TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem) *TicketCalendarBatchSaveRequest {
	s.TicketSkuAndCalendar = v
	return s
}

func (s *TicketCalendarBatchSaveRequest) SetAccountId(v string) *TicketCalendarBatchSaveRequest {
	s.AccountId = &v
	return s
}

func (s *TicketCalendarBatchSaveRequest) SetProductId(v string) *TicketCalendarBatchSaveRequest {
	s.ProductId = &v
	return s
}

func (s *TicketCalendarBatchSaveRequest) SetProductOutId(v string) *TicketCalendarBatchSaveRequest {
	s.ProductOutId = &v
	return s
}

func (s *TicketCalendarBatchSaveRequest) SetHeader(v map[string]*string) *TicketCalendarBatchSaveRequest {
	s.Header = v
	return s
}

type TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem struct {
	CalendarList        []*TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem  `json:"calendar_list,omitempty" xml:"calendar_list,omitempty" type:"Repeated"`
	TicketSpecification *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification `json:"ticket_specification,omitempty" xml:"ticket_specification,omitempty" require:"true"`
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem) SetCalendarList(v []*TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem {
	s.CalendarList = v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem) SetTicketSpecification(v *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItem {
	s.TicketSpecification = v
	return s
}

type TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem struct {
	CalendarValue *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue `json:"calendar_value,omitempty" xml:"calendar_value,omitempty" require:"true"`
	OriginAmount  *int64                                                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty" require:"true"`
	SettleAmount  *int64                                                                               `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	Status        *int                                                                                 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	StockQty      *int64                                                                               `json:"stock_qty,omitempty" xml:"stock_qty,omitempty" require:"true"`
	ActualAmount  *int64                                                                               `json:"actual_amount,omitempty" xml:"actual_amount,omitempty" require:"true"`
	CalendarType  *int                                                                                 `json:"calendar_type,omitempty" xml:"calendar_type,omitempty" require:"true"`
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetCalendarValue(v *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.CalendarValue = v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetOriginAmount(v int64) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.OriginAmount = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetSettleAmount(v int64) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.SettleAmount = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetStatus(v int) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.Status = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetStockQty(v int64) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.StockQty = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetActualAmount(v int64) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.ActualAmount = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem) SetCalendarType(v int) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItem {
	s.CalendarType = &v
	return s
}

type TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue struct {
	ValueList       []*string `json:"value_list,omitempty" xml:"value_list,omitempty" require:"true" type:"Repeated"`
	EndDate         *string   `json:"end_date,omitempty" xml:"end_date,omitempty"`
	ExcludeDateList []*string `json:"exclude_date_list,omitempty" xml:"exclude_date_list,omitempty" type:"Repeated"`
	StartDate       *string   `json:"start_date,omitempty" xml:"start_date,omitempty"`
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) SetValueList(v []*string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue {
	s.ValueList = v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) SetEndDate(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue {
	s.EndDate = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) SetExcludeDateList(v []*string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue {
	s.ExcludeDateList = v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue) SetStartDate(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemCalendarListItemCalendarValue {
	s.StartDate = &v
	return s
}

type TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification struct {
	SkuName       *string                                                                                 `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuOutId      *string                                                                                 `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	TicketArea    *string                                                                                 `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
	TicketSeat    *string                                                                                 `json:"ticket_seat,omitempty" xml:"ticket_seat,omitempty"`
	TicketSession *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession `json:"ticket_session,omitempty" xml:"ticket_session,omitempty"`
	FreeMerchant  *bool                                                                                   `json:"free_merchant,omitempty" xml:"free_merchant,omitempty"`
	SettleType    *int                                                                                    `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	SkuId         *string                                                                                 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetSkuName(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.SkuName = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetSkuOutId(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.SkuOutId = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetTicketArea(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.TicketArea = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetTicketSeat(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.TicketSeat = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetTicketSession(v *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.TicketSession = v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetFreeMerchant(v bool) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.FreeMerchant = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetSettleType(v int) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.SettleType = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification) SetSkuId(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecification {
	s.SkuId = &v
	return s
}

type TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession struct {
	TicketSessionName *string `json:"ticket_session_name,omitempty" xml:"ticket_session_name,omitempty" require:"true"`
	TicketSessionTime *string `json:"ticket_session_time,omitempty" xml:"ticket_session_time,omitempty"`
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession) SetTicketSessionName(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession {
	s.TicketSessionName = &v
	return s
}

func (s *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession) SetTicketSessionTime(v string) *TicketCalendarBatchSaveRequestTicketSkuAndCalendarItemTicketSpecificationTicketSession {
	s.TicketSessionTime = &v
	return s
}

type TicketCalendarBatchSaveResponse struct {
	Extra *TicketCalendarBatchSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TicketCalendarBatchSaveResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TicketCalendarBatchSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveResponse) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveResponse) SetExtra(v *TicketCalendarBatchSaveResponseExtra) *TicketCalendarBatchSaveResponse {
	s.Extra = v
	return s
}

func (s *TicketCalendarBatchSaveResponse) SetData(v *TicketCalendarBatchSaveResponseData) *TicketCalendarBatchSaveResponse {
	s.Data = v
	return s
}

type TicketCalendarBatchSaveResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TicketCalendarBatchSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveResponseData) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveResponseData) SetGwErrorCode(v int32) *TicketCalendarBatchSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseData) SetGwDescription(v string) *TicketCalendarBatchSaveResponseData {
	s.GwDescription = &v
	return s
}

type TicketCalendarBatchSaveResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TicketCalendarBatchSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarBatchSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *TicketCalendarBatchSaveResponseExtra) SetNow(v int64) *TicketCalendarBatchSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseExtra) SetSubDescription(v string) *TicketCalendarBatchSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseExtra) SetSubErrorCode(v int32) *TicketCalendarBatchSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseExtra) SetDescription(v string) *TicketCalendarBatchSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseExtra) SetErrorCode(v int32) *TicketCalendarBatchSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TicketCalendarBatchSaveResponseExtra) SetLogid(v string) *TicketCalendarBatchSaveResponseExtra {
	s.Logid = &v
	return s
}

type TicketCalendarSaveRequest struct {
	TicketSpecification *TicketCalendarSaveRequestTicketSpecification `json:"ticket_specification,omitempty" xml:"ticket_specification,omitempty" require:"true"`
	AccountId           *string                                       `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header              map[string]*string                            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken         *string                                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Calendars           []*TicketCalendarSaveRequestCalendarsItem     `json:"calendars,omitempty" xml:"calendars,omitempty" require:"true" type:"Repeated"`
	ProductId           *string                                       `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductOutId        *string                                       `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
}

func (s TicketCalendarSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveRequest) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveRequest) SetTicketSpecification(v *TicketCalendarSaveRequestTicketSpecification) *TicketCalendarSaveRequest {
	s.TicketSpecification = v
	return s
}

func (s *TicketCalendarSaveRequest) SetAccountId(v string) *TicketCalendarSaveRequest {
	s.AccountId = &v
	return s
}

func (s *TicketCalendarSaveRequest) SetHeader(v map[string]*string) *TicketCalendarSaveRequest {
	s.Header = v
	return s
}

func (s *TicketCalendarSaveRequest) SetAccessToken(v string) *TicketCalendarSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *TicketCalendarSaveRequest) SetCalendars(v []*TicketCalendarSaveRequestCalendarsItem) *TicketCalendarSaveRequest {
	s.Calendars = v
	return s
}

func (s *TicketCalendarSaveRequest) SetProductId(v string) *TicketCalendarSaveRequest {
	s.ProductId = &v
	return s
}

func (s *TicketCalendarSaveRequest) SetProductOutId(v string) *TicketCalendarSaveRequest {
	s.ProductOutId = &v
	return s
}

type TicketCalendarSaveRequestCalendarsItem struct {
	SettleAmount  *int64                                               `json:"settle_amount,omitempty" xml:"settle_amount,omitempty"`
	Status        *int                                                 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	StockQty      *int64                                               `json:"stock_qty,omitempty" xml:"stock_qty,omitempty" require:"true"`
	ActualAmount  *int64                                               `json:"actual_amount,omitempty" xml:"actual_amount,omitempty" require:"true"`
	CalendarType  *int                                                 `json:"calendar_type,omitempty" xml:"calendar_type,omitempty" require:"true"`
	CalendarValue *TicketCalendarSaveRequestCalendarsItemCalendarValue `json:"calendar_value,omitempty" xml:"calendar_value,omitempty" require:"true"`
	OriginAmount  *int64                                               `json:"origin_amount,omitempty" xml:"origin_amount,omitempty" require:"true"`
}

func (s TicketCalendarSaveRequestCalendarsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveRequestCalendarsItem) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetSettleAmount(v int64) *TicketCalendarSaveRequestCalendarsItem {
	s.SettleAmount = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetStatus(v int) *TicketCalendarSaveRequestCalendarsItem {
	s.Status = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetStockQty(v int64) *TicketCalendarSaveRequestCalendarsItem {
	s.StockQty = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetActualAmount(v int64) *TicketCalendarSaveRequestCalendarsItem {
	s.ActualAmount = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetCalendarType(v int) *TicketCalendarSaveRequestCalendarsItem {
	s.CalendarType = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetCalendarValue(v *TicketCalendarSaveRequestCalendarsItemCalendarValue) *TicketCalendarSaveRequestCalendarsItem {
	s.CalendarValue = v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItem) SetOriginAmount(v int64) *TicketCalendarSaveRequestCalendarsItem {
	s.OriginAmount = &v
	return s
}

type TicketCalendarSaveRequestCalendarsItemCalendarValue struct {
	EndDate         *string   `json:"end_date,omitempty" xml:"end_date,omitempty"`
	ExcludeDateList []*string `json:"exclude_date_list,omitempty" xml:"exclude_date_list,omitempty" type:"Repeated"`
	StartDate       *string   `json:"start_date,omitempty" xml:"start_date,omitempty"`
	ValueList       []*string `json:"value_list,omitempty" xml:"value_list,omitempty" require:"true" type:"Repeated"`
}

func (s TicketCalendarSaveRequestCalendarsItemCalendarValue) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveRequestCalendarsItemCalendarValue) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveRequestCalendarsItemCalendarValue) SetEndDate(v string) *TicketCalendarSaveRequestCalendarsItemCalendarValue {
	s.EndDate = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItemCalendarValue) SetExcludeDateList(v []*string) *TicketCalendarSaveRequestCalendarsItemCalendarValue {
	s.ExcludeDateList = v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItemCalendarValue) SetStartDate(v string) *TicketCalendarSaveRequestCalendarsItemCalendarValue {
	s.StartDate = &v
	return s
}

func (s *TicketCalendarSaveRequestCalendarsItemCalendarValue) SetValueList(v []*string) *TicketCalendarSaveRequestCalendarsItemCalendarValue {
	s.ValueList = v
	return s
}

type TicketCalendarSaveRequestTicketSpecification struct {
	SkuName       *string                                                    `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuOutId      *string                                                    `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	TicketArea    *string                                                    `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
	TicketSeat    *string                                                    `json:"ticket_seat,omitempty" xml:"ticket_seat,omitempty"`
	TicketSession *TicketCalendarSaveRequestTicketSpecificationTicketSession `json:"ticket_session,omitempty" xml:"ticket_session,omitempty"`
	FreeMerchant  *bool                                                      `json:"free_merchant,omitempty" xml:"free_merchant,omitempty"`
	SettleType    *int                                                       `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	SkuId         *string                                                    `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s TicketCalendarSaveRequestTicketSpecification) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveRequestTicketSpecification) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetSkuName(v string) *TicketCalendarSaveRequestTicketSpecification {
	s.SkuName = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetSkuOutId(v string) *TicketCalendarSaveRequestTicketSpecification {
	s.SkuOutId = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetTicketArea(v string) *TicketCalendarSaveRequestTicketSpecification {
	s.TicketArea = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetTicketSeat(v string) *TicketCalendarSaveRequestTicketSpecification {
	s.TicketSeat = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetTicketSession(v *TicketCalendarSaveRequestTicketSpecificationTicketSession) *TicketCalendarSaveRequestTicketSpecification {
	s.TicketSession = v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetFreeMerchant(v bool) *TicketCalendarSaveRequestTicketSpecification {
	s.FreeMerchant = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetSettleType(v int) *TicketCalendarSaveRequestTicketSpecification {
	s.SettleType = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecification) SetSkuId(v string) *TicketCalendarSaveRequestTicketSpecification {
	s.SkuId = &v
	return s
}

type TicketCalendarSaveRequestTicketSpecificationTicketSession struct {
	TicketSessionTime *string `json:"ticket_session_time,omitempty" xml:"ticket_session_time,omitempty"`
	TicketSessionName *string `json:"ticket_session_name,omitempty" xml:"ticket_session_name,omitempty" require:"true"`
}

func (s TicketCalendarSaveRequestTicketSpecificationTicketSession) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveRequestTicketSpecificationTicketSession) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveRequestTicketSpecificationTicketSession) SetTicketSessionTime(v string) *TicketCalendarSaveRequestTicketSpecificationTicketSession {
	s.TicketSessionTime = &v
	return s
}

func (s *TicketCalendarSaveRequestTicketSpecificationTicketSession) SetTicketSessionName(v string) *TicketCalendarSaveRequestTicketSpecificationTicketSession {
	s.TicketSessionName = &v
	return s
}

type TicketCalendarSaveResponse struct {
	Data  *TicketCalendarSaveResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *TicketCalendarSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TicketCalendarSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveResponse) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveResponse) SetData(v *TicketCalendarSaveResponseData) *TicketCalendarSaveResponse {
	s.Data = v
	return s
}

func (s *TicketCalendarSaveResponse) SetExtra(v *TicketCalendarSaveResponseExtra) *TicketCalendarSaveResponse {
	s.Extra = v
	return s
}

type TicketCalendarSaveResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TicketCalendarSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveResponseData) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveResponseData) SetGwErrorCode(v int32) *TicketCalendarSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TicketCalendarSaveResponseData) SetGwDescription(v string) *TicketCalendarSaveResponseData {
	s.GwDescription = &v
	return s
}

type TicketCalendarSaveResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s TicketCalendarSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TicketCalendarSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *TicketCalendarSaveResponseExtra) SetDescription(v string) *TicketCalendarSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *TicketCalendarSaveResponseExtra) SetErrorCode(v int32) *TicketCalendarSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TicketCalendarSaveResponseExtra) SetLogid(v string) *TicketCalendarSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *TicketCalendarSaveResponseExtra) SetNow(v int64) *TicketCalendarSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *TicketCalendarSaveResponseExtra) SetSubDescription(v string) *TicketCalendarSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TicketCalendarSaveResponseExtra) SetSubErrorCode(v int32) *TicketCalendarSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

type TicketQueryRequest struct {
	AccountId   *int64                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Param       *TicketQueryRequestParam `json:"param,omitempty" xml:"param,omitempty"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TicketQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketQueryRequest) GoString() string {
	return s.String()
}

func (s *TicketQueryRequest) SetAccountId(v int64) *TicketQueryRequest {
	s.AccountId = &v
	return s
}

func (s *TicketQueryRequest) SetParam(v *TicketQueryRequestParam) *TicketQueryRequest {
	s.Param = v
	return s
}

func (s *TicketQueryRequest) SetHeader(v map[string]*string) *TicketQueryRequest {
	s.Header = v
	return s
}

func (s *TicketQueryRequest) SetAccessToken(v string) *TicketQueryRequest {
	s.AccessToken = &v
	return s
}

type TicketQueryRequestParam struct {
	PoiId    *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TicketId *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

func (s TicketQueryRequestParam) String() string {
	return tea.Prettify(s)
}

func (s TicketQueryRequestParam) GoString() string {
	return s.String()
}

func (s *TicketQueryRequestParam) SetPoiId(v int64) *TicketQueryRequestParam {
	s.PoiId = &v
	return s
}

func (s *TicketQueryRequestParam) SetTicketId(v string) *TicketQueryRequestParam {
	s.TicketId = &v
	return s
}

type TicketQueryResponse struct {
	Data  *TicketQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TicketQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TicketQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketQueryResponse) GoString() string {
	return s.String()
}

func (s *TicketQueryResponse) SetData(v *TicketQueryResponseData) *TicketQueryResponse {
	s.Data = v
	return s
}

func (s *TicketQueryResponse) SetExtra(v *TicketQueryResponseExtra) *TicketQueryResponse {
	s.Extra = v
	return s
}

type TicketQueryResponseData struct {
	AppealStatus  *int    `json:"appeal_status,omitempty" xml:"appeal_status,omitempty"`
	PoiId         *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ResultDesc    *string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	TicketId      *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TicketQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s TicketQueryResponseData) GoString() string {
	return s.String()
}

func (s *TicketQueryResponseData) SetAppealStatus(v int) *TicketQueryResponseData {
	s.AppealStatus = &v
	return s
}

func (s *TicketQueryResponseData) SetPoiId(v int64) *TicketQueryResponseData {
	s.PoiId = &v
	return s
}

func (s *TicketQueryResponseData) SetResultDesc(v string) *TicketQueryResponseData {
	s.ResultDesc = &v
	return s
}

func (s *TicketQueryResponseData) SetTicketId(v string) *TicketQueryResponseData {
	s.TicketId = &v
	return s
}

func (s *TicketQueryResponseData) SetGwErrorCode(v int32) *TicketQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TicketQueryResponseData) SetGwDescription(v string) *TicketQueryResponseData {
	s.GwDescription = &v
	return s
}

type TicketQueryResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TicketQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TicketQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *TicketQueryResponseExtra) SetErrorCode(v int32) *TicketQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryResponseExtra) SetLogid(v string) *TicketQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *TicketQueryResponseExtra) SetNow(v int64) *TicketQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *TicketQueryResponseExtra) SetSubDescription(v string) *TicketQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TicketQueryResponseExtra) SetSubErrorCode(v int32) *TicketQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TicketQueryResponseExtra) SetDescription(v string) *TicketQueryResponseExtra {
	s.Description = &v
	return s
}

type TicketSaveRequest struct {
	AccountId   *string                  `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Ticket      *TicketSaveRequestTicket `json:"ticket,omitempty" xml:"ticket,omitempty" require:"true"`
}

func (s TicketSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequest) GoString() string {
	return s.String()
}

func (s *TicketSaveRequest) SetAccountId(v string) *TicketSaveRequest {
	s.AccountId = &v
	return s
}

func (s *TicketSaveRequest) SetHeader(v map[string]*string) *TicketSaveRequest {
	s.Header = v
	return s
}

func (s *TicketSaveRequest) SetAccessToken(v string) *TicketSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *TicketSaveRequest) SetTicket(v *TicketSaveRequestTicket) *TicketSaveRequest {
	s.Ticket = v
	return s
}

type TicketSaveRequestTicket struct {
	PoiIds               []*string                                          `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" require:"true" type:"Repeated"`
	BuyLimit             *int32                                             `json:"buy_limit,omitempty" xml:"buy_limit,omitempty"`
	TicketRule           *TicketSaveRequestTicketTicketRule                 `json:"ticket_rule,omitempty" xml:"ticket_rule,omitempty" require:"true"`
	IsTestData           *bool                                              `json:"is_test_data,omitempty" xml:"is_test_data,omitempty"`
	TestUid              []*string                                          `json:"test_uid,omitempty" xml:"test_uid,omitempty" type:"Repeated"`
	UseTicketAddress     *string                                            `json:"use_ticket_address,omitempty" xml:"use_ticket_address,omitempty"`
	ProductId            *string                                            `json:"product_id,omitempty" xml:"product_id,omitempty"`
	TicketTypeName       *string                                            `json:"ticket_type_name,omitempty" xml:"ticket_type_name,omitempty" require:"true"`
	AdmissionTime        *int32                                             `json:"admission_time,omitempty" xml:"admission_time,omitempty"`
	CrowdOpType          *int                                               `json:"crowd_op_type,omitempty" xml:"crowd_op_type,omitempty"`
	TicketTypeId         *string                                            `json:"ticket_type_id,omitempty" xml:"ticket_type_id,omitempty"`
	IsTestProduct        *bool                                              `json:"is_test_product,omitempty" xml:"is_test_product,omitempty"`
	Crowds               []*TicketSaveRequestTicketCrowdsItem               `json:"crowds,omitempty" xml:"crowds,omitempty" require:"true" type:"Repeated"`
	OpeningTime          *TicketSaveRequestTicketOpeningTime                `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	TicketPeriod         *int32                                             `json:"ticket_period,omitempty" xml:"ticket_period,omitempty" require:"true"`
	AuditResult          *TicketSaveRequestTicketAuditResult                `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	DraftStatus          *int                                               `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	CustomerReservedInfo *TicketSaveRequestTicketCustomerReservedInfo       `json:"customer_reserved_info,omitempty" xml:"customer_reserved_info,omitempty" require:"true"`
	Status               *int                                               `json:"status,omitempty" xml:"status,omitempty"`
	ProductName          *string                                            `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	ProductOutId         *string                                            `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	PreOrderTime         *TicketSaveRequestTicketPreOrderTime               `json:"pre_order_time,omitempty" xml:"pre_order_time,omitempty" require:"true"`
	TicketSpecifications []*TicketSaveRequestTicketTicketSpecificationsItem `json:"ticket_specifications,omitempty" xml:"ticket_specifications,omitempty" require:"true" type:"Repeated"`
	TicketTypeOutId      *string                                            `json:"ticket_type_out_id,omitempty" xml:"ticket_type_out_id,omitempty"`
	ShowChannel          *int32                                             `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	OtherNote            *string                                            `json:"other_note,omitempty" xml:"other_note,omitempty"`
	CategoryId           *int64                                             `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	FeeNotInclude        *string                                            `json:"fee_not_include,omitempty" xml:"fee_not_include,omitempty"`
	UseTimes             *int32                                             `json:"use_times,omitempty" xml:"use_times,omitempty" require:"true"`
	StockSyncRule        *TicketSaveRequestTicketStockSyncRule              `json:"stock_sync_rule,omitempty" xml:"stock_sync_rule,omitempty"`
	ShowDate             *int32                                             `json:"show_date,omitempty" xml:"show_date,omitempty"`
	RefundRule           *TicketSaveRequestTicketRefundRule                 `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" require:"true"`
	RequireIdCard        *bool                                              `json:"require_id_card,omitempty" xml:"require_id_card,omitempty"`
	MiniProgram          *TicketSaveRequestTicketMiniProgram                `json:"mini_program,omitempty" xml:"mini_program,omitempty"`
	NeedTicket           *bool                                              `json:"need_ticket,omitempty" xml:"need_ticket,omitempty" require:"true"`
	ImgUrls              []*string                                          `json:"img_urls,omitempty" xml:"img_urls,omitempty" require:"true" type:"Repeated"`
	BuyLimitRule         []*TicketSaveRequestTicketBuyLimitRuleItem         `json:"buy_limit_rule,omitempty" xml:"buy_limit_rule,omitempty" type:"Repeated"`
	FeeInclude           *string                                            `json:"fee_include,omitempty" xml:"fee_include,omitempty" require:"true"`
	Region               *TicketSaveRequestTicketRegion                     `json:"region,omitempty" xml:"region,omitempty"`
}

func (s TicketSaveRequestTicket) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicket) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicket) SetPoiIds(v []*string) *TicketSaveRequestTicket {
	s.PoiIds = v
	return s
}

func (s *TicketSaveRequestTicket) SetBuyLimit(v int32) *TicketSaveRequestTicket {
	s.BuyLimit = &v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketRule(v *TicketSaveRequestTicketTicketRule) *TicketSaveRequestTicket {
	s.TicketRule = v
	return s
}

func (s *TicketSaveRequestTicket) SetIsTestData(v bool) *TicketSaveRequestTicket {
	s.IsTestData = &v
	return s
}

func (s *TicketSaveRequestTicket) SetTestUid(v []*string) *TicketSaveRequestTicket {
	s.TestUid = v
	return s
}

func (s *TicketSaveRequestTicket) SetUseTicketAddress(v string) *TicketSaveRequestTicket {
	s.UseTicketAddress = &v
	return s
}

func (s *TicketSaveRequestTicket) SetProductId(v string) *TicketSaveRequestTicket {
	s.ProductId = &v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketTypeName(v string) *TicketSaveRequestTicket {
	s.TicketTypeName = &v
	return s
}

func (s *TicketSaveRequestTicket) SetAdmissionTime(v int32) *TicketSaveRequestTicket {
	s.AdmissionTime = &v
	return s
}

func (s *TicketSaveRequestTicket) SetCrowdOpType(v int) *TicketSaveRequestTicket {
	s.CrowdOpType = &v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketTypeId(v string) *TicketSaveRequestTicket {
	s.TicketTypeId = &v
	return s
}

func (s *TicketSaveRequestTicket) SetIsTestProduct(v bool) *TicketSaveRequestTicket {
	s.IsTestProduct = &v
	return s
}

func (s *TicketSaveRequestTicket) SetCrowds(v []*TicketSaveRequestTicketCrowdsItem) *TicketSaveRequestTicket {
	s.Crowds = v
	return s
}

func (s *TicketSaveRequestTicket) SetOpeningTime(v *TicketSaveRequestTicketOpeningTime) *TicketSaveRequestTicket {
	s.OpeningTime = v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketPeriod(v int32) *TicketSaveRequestTicket {
	s.TicketPeriod = &v
	return s
}

func (s *TicketSaveRequestTicket) SetAuditResult(v *TicketSaveRequestTicketAuditResult) *TicketSaveRequestTicket {
	s.AuditResult = v
	return s
}

func (s *TicketSaveRequestTicket) SetDraftStatus(v int) *TicketSaveRequestTicket {
	s.DraftStatus = &v
	return s
}

func (s *TicketSaveRequestTicket) SetCustomerReservedInfo(v *TicketSaveRequestTicketCustomerReservedInfo) *TicketSaveRequestTicket {
	s.CustomerReservedInfo = v
	return s
}

func (s *TicketSaveRequestTicket) SetStatus(v int) *TicketSaveRequestTicket {
	s.Status = &v
	return s
}

func (s *TicketSaveRequestTicket) SetProductName(v string) *TicketSaveRequestTicket {
	s.ProductName = &v
	return s
}

func (s *TicketSaveRequestTicket) SetProductOutId(v string) *TicketSaveRequestTicket {
	s.ProductOutId = &v
	return s
}

func (s *TicketSaveRequestTicket) SetPreOrderTime(v *TicketSaveRequestTicketPreOrderTime) *TicketSaveRequestTicket {
	s.PreOrderTime = v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketSpecifications(v []*TicketSaveRequestTicketTicketSpecificationsItem) *TicketSaveRequestTicket {
	s.TicketSpecifications = v
	return s
}

func (s *TicketSaveRequestTicket) SetTicketTypeOutId(v string) *TicketSaveRequestTicket {
	s.TicketTypeOutId = &v
	return s
}

func (s *TicketSaveRequestTicket) SetShowChannel(v int32) *TicketSaveRequestTicket {
	s.ShowChannel = &v
	return s
}

func (s *TicketSaveRequestTicket) SetOtherNote(v string) *TicketSaveRequestTicket {
	s.OtherNote = &v
	return s
}

func (s *TicketSaveRequestTicket) SetCategoryId(v int64) *TicketSaveRequestTicket {
	s.CategoryId = &v
	return s
}

func (s *TicketSaveRequestTicket) SetFeeNotInclude(v string) *TicketSaveRequestTicket {
	s.FeeNotInclude = &v
	return s
}

func (s *TicketSaveRequestTicket) SetUseTimes(v int32) *TicketSaveRequestTicket {
	s.UseTimes = &v
	return s
}

func (s *TicketSaveRequestTicket) SetStockSyncRule(v *TicketSaveRequestTicketStockSyncRule) *TicketSaveRequestTicket {
	s.StockSyncRule = v
	return s
}

func (s *TicketSaveRequestTicket) SetShowDate(v int32) *TicketSaveRequestTicket {
	s.ShowDate = &v
	return s
}

func (s *TicketSaveRequestTicket) SetRefundRule(v *TicketSaveRequestTicketRefundRule) *TicketSaveRequestTicket {
	s.RefundRule = v
	return s
}

func (s *TicketSaveRequestTicket) SetRequireIdCard(v bool) *TicketSaveRequestTicket {
	s.RequireIdCard = &v
	return s
}

func (s *TicketSaveRequestTicket) SetMiniProgram(v *TicketSaveRequestTicketMiniProgram) *TicketSaveRequestTicket {
	s.MiniProgram = v
	return s
}

func (s *TicketSaveRequestTicket) SetNeedTicket(v bool) *TicketSaveRequestTicket {
	s.NeedTicket = &v
	return s
}

func (s *TicketSaveRequestTicket) SetImgUrls(v []*string) *TicketSaveRequestTicket {
	s.ImgUrls = v
	return s
}

func (s *TicketSaveRequestTicket) SetBuyLimitRule(v []*TicketSaveRequestTicketBuyLimitRuleItem) *TicketSaveRequestTicket {
	s.BuyLimitRule = v
	return s
}

func (s *TicketSaveRequestTicket) SetFeeInclude(v string) *TicketSaveRequestTicket {
	s.FeeInclude = &v
	return s
}

func (s *TicketSaveRequestTicket) SetRegion(v *TicketSaveRequestTicketRegion) *TicketSaveRequestTicket {
	s.Region = v
	return s
}

type TicketSaveRequestTicketAuditResult struct {
	IsTicketTypePass      *bool   `json:"is_ticket_type_pass,omitempty" xml:"is_ticket_type_pass,omitempty" require:"true"`
	ProductAuditReason    *string `json:"product_audit_reason,omitempty" xml:"product_audit_reason,omitempty"`
	TicketTypeAuditReason *string `json:"ticket_type_audit_reason,omitempty" xml:"ticket_type_audit_reason,omitempty"`
	IsProductPass         *bool   `json:"is_product_pass,omitempty" xml:"is_product_pass,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketAuditResult) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketAuditResult) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketAuditResult) SetIsTicketTypePass(v bool) *TicketSaveRequestTicketAuditResult {
	s.IsTicketTypePass = &v
	return s
}

func (s *TicketSaveRequestTicketAuditResult) SetProductAuditReason(v string) *TicketSaveRequestTicketAuditResult {
	s.ProductAuditReason = &v
	return s
}

func (s *TicketSaveRequestTicketAuditResult) SetTicketTypeAuditReason(v string) *TicketSaveRequestTicketAuditResult {
	s.TicketTypeAuditReason = &v
	return s
}

func (s *TicketSaveRequestTicketAuditResult) SetIsProductPass(v bool) *TicketSaveRequestTicketAuditResult {
	s.IsProductPass = &v
	return s
}

type TicketSaveRequestTicketBuyLimitRuleItem struct {
	LimitRange   *int      `json:"limit_range,omitempty" xml:"limit_range,omitempty" require:"true"`
	LimitSubject *int      `json:"limit_subject,omitempty" xml:"limit_subject,omitempty" require:"true"`
	TimeLength   *int32    `json:"time_length,omitempty" xml:"time_length,omitempty"`
	TimePeriod   []*string `json:"time_period,omitempty" xml:"time_period,omitempty" type:"Repeated"`
	LimitNum     *int32    `json:"limit_num,omitempty" xml:"limit_num,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketBuyLimitRuleItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketBuyLimitRuleItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketBuyLimitRuleItem) SetLimitRange(v int) *TicketSaveRequestTicketBuyLimitRuleItem {
	s.LimitRange = &v
	return s
}

func (s *TicketSaveRequestTicketBuyLimitRuleItem) SetLimitSubject(v int) *TicketSaveRequestTicketBuyLimitRuleItem {
	s.LimitSubject = &v
	return s
}

func (s *TicketSaveRequestTicketBuyLimitRuleItem) SetTimeLength(v int32) *TicketSaveRequestTicketBuyLimitRuleItem {
	s.TimeLength = &v
	return s
}

func (s *TicketSaveRequestTicketBuyLimitRuleItem) SetTimePeriod(v []*string) *TicketSaveRequestTicketBuyLimitRuleItem {
	s.TimePeriod = v
	return s
}

func (s *TicketSaveRequestTicketBuyLimitRuleItem) SetLimitNum(v int32) *TicketSaveRequestTicketBuyLimitRuleItem {
	s.LimitNum = &v
	return s
}

type TicketSaveRequestTicketCrowdsItem struct {
	CrowdCondition *TicketSaveRequestTicketCrowdsItemCrowdCondition `json:"crowd_condition,omitempty" xml:"crowd_condition,omitempty"`
	CrowdNum       *int32                                           `json:"crowd_num,omitempty" xml:"crowd_num,omitempty" require:"true"`
	CrowdType      *int                                             `json:"crowd_type,omitempty" xml:"crowd_type,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketCrowdsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCrowdsItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCrowdsItem) SetCrowdCondition(v *TicketSaveRequestTicketCrowdsItemCrowdCondition) *TicketSaveRequestTicketCrowdsItem {
	s.CrowdCondition = v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItem) SetCrowdNum(v int32) *TicketSaveRequestTicketCrowdsItem {
	s.CrowdNum = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItem) SetCrowdType(v int) *TicketSaveRequestTicketCrowdsItem {
	s.CrowdType = &v
	return s
}

type TicketSaveRequestTicketCrowdsItemCrowdCondition struct {
	AgeCondition         *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition    `json:"age_condition,omitempty" xml:"age_condition,omitempty"`
	AgeHeightOpType      *int                                                            `json:"age_height_op_type,omitempty" xml:"age_height_op_type,omitempty"`
	CertificateCondition []*int                                                          `json:"certificate_condition,omitempty" xml:"certificate_condition,omitempty" type:"Repeated"`
	HeightCondition      *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition `json:"height_condition,omitempty" xml:"height_condition,omitempty"`
	Note                 *string                                                         `json:"note,omitempty" xml:"note,omitempty"`
	ScenicDefaultCrowd   *bool                                                           `json:"scenic_default_crowd,omitempty" xml:"scenic_default_crowd,omitempty"`
	WeightCondition      *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition `json:"weight_condition,omitempty" xml:"weight_condition,omitempty"`
	AccompanyCondition   *int                                                            `json:"accompany_condition,omitempty" xml:"accompany_condition,omitempty"`
}

func (s TicketSaveRequestTicketCrowdsItemCrowdCondition) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCrowdsItemCrowdCondition) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetAgeCondition(v *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.AgeCondition = v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetAgeHeightOpType(v int) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.AgeHeightOpType = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetCertificateCondition(v []*int) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.CertificateCondition = v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetHeightCondition(v *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.HeightCondition = v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetNote(v string) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.Note = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetScenicDefaultCrowd(v bool) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.ScenicDefaultCrowd = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetWeightCondition(v *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.WeightCondition = v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdCondition) SetAccompanyCondition(v int) *TicketSaveRequestTicketCrowdsItemCrowdCondition {
	s.AccompanyCondition = &v
	return s
}

type TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition struct {
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) SetLowerBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition {
	s.LowerBoundType = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) SetUpperBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition {
	s.UpperBound = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) SetUpperBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition {
	s.UpperBoundType = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition) SetLowerBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionAgeCondition {
	s.LowerBound = &v
	return s
}

type TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition struct {
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) SetUpperBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition {
	s.UpperBound = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) SetUpperBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition {
	s.UpperBoundType = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) SetLowerBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition {
	s.LowerBound = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition) SetLowerBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionHeightCondition {
	s.LowerBoundType = &v
	return s
}

type TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition struct {
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) SetUpperBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition {
	s.UpperBound = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) SetUpperBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition {
	s.UpperBoundType = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) SetLowerBound(v int64) *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition {
	s.LowerBound = &v
	return s
}

func (s *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition) SetLowerBoundType(v int) *TicketSaveRequestTicketCrowdsItemCrowdConditionWeightCondition {
	s.LowerBoundType = &v
	return s
}

type TicketSaveRequestTicketCustomerReservedInfo struct {
	ReservedCrowdNumber *int32 `json:"reserved_crowd_number,omitempty" xml:"reserved_crowd_number,omitempty"`
	ReservedRules       []*int `json:"reserved_rules,omitempty" xml:"reserved_rules,omitempty" type:"Repeated"`
	ReservedType        *int   `json:"reserved_type,omitempty" xml:"reserved_type,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketCustomerReservedInfo) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketCustomerReservedInfo) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketCustomerReservedInfo) SetReservedCrowdNumber(v int32) *TicketSaveRequestTicketCustomerReservedInfo {
	s.ReservedCrowdNumber = &v
	return s
}

func (s *TicketSaveRequestTicketCustomerReservedInfo) SetReservedRules(v []*int) *TicketSaveRequestTicketCustomerReservedInfo {
	s.ReservedRules = v
	return s
}

func (s *TicketSaveRequestTicketCustomerReservedInfo) SetReservedType(v int) *TicketSaveRequestTicketCustomerReservedInfo {
	s.ReservedType = &v
	return s
}

type TicketSaveRequestTicketMiniProgram struct {
	EntryType          *int32  `json:"entry_type,omitempty" xml:"entry_type,omitempty"`
	SaleUnitDetailPath *string `json:"sale_unit_detail_path,omitempty" xml:"sale_unit_detail_path,omitempty"`
	TradeUrl           *string `json:"trade_url,omitempty" xml:"trade_url,omitempty"`
	AppId              *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketMiniProgram) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketMiniProgram) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketMiniProgram) SetEntryType(v int32) *TicketSaveRequestTicketMiniProgram {
	s.EntryType = &v
	return s
}

func (s *TicketSaveRequestTicketMiniProgram) SetSaleUnitDetailPath(v string) *TicketSaveRequestTicketMiniProgram {
	s.SaleUnitDetailPath = &v
	return s
}

func (s *TicketSaveRequestTicketMiniProgram) SetTradeUrl(v string) *TicketSaveRequestTicketMiniProgram {
	s.TradeUrl = &v
	return s
}

func (s *TicketSaveRequestTicketMiniProgram) SetAppId(v string) *TicketSaveRequestTicketMiniProgram {
	s.AppId = &v
	return s
}

type TicketSaveRequestTicketOpeningTime struct {
	EarliestTime *int64 `json:"earliest_time,omitempty" xml:"earliest_time,omitempty"`
	LatestTime   *int64 `json:"latest_time,omitempty" xml:"latest_time,omitempty"`
}

func (s TicketSaveRequestTicketOpeningTime) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketOpeningTime) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketOpeningTime) SetEarliestTime(v int64) *TicketSaveRequestTicketOpeningTime {
	s.EarliestTime = &v
	return s
}

func (s *TicketSaveRequestTicketOpeningTime) SetLatestTime(v int64) *TicketSaveRequestTicketOpeningTime {
	s.LatestTime = &v
	return s
}

type TicketSaveRequestTicketPreOrderTime struct {
	PreTime *int64 `json:"pre_time,omitempty" xml:"pre_time,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketPreOrderTime) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketPreOrderTime) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketPreOrderTime) SetPreTime(v int64) *TicketSaveRequestTicketPreOrderTime {
	s.PreTime = &v
	return s
}

type TicketSaveRequestTicketRefundRule struct {
	RefundDetails  []*TicketSaveRequestTicketRefundRuleRefundDetailsItem `json:"refund_details,omitempty" xml:"refund_details,omitempty" type:"Repeated"`
	RefundPart     *bool                                                 `json:"refund_part,omitempty" xml:"refund_part,omitempty"`
	RefundType     *int                                                  `json:"refund_type,omitempty" xml:"refund_type,omitempty" require:"true"`
	AutoRefundTime *int64                                                `json:"auto_refund_time,omitempty" xml:"auto_refund_time,omitempty"`
	AutoVerifyTime *int64                                                `json:"auto_verify_time,omitempty" xml:"auto_verify_time,omitempty"`
}

func (s TicketSaveRequestTicketRefundRule) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketRefundRule) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketRefundRule) SetRefundDetails(v []*TicketSaveRequestTicketRefundRuleRefundDetailsItem) *TicketSaveRequestTicketRefundRule {
	s.RefundDetails = v
	return s
}

func (s *TicketSaveRequestTicketRefundRule) SetRefundPart(v bool) *TicketSaveRequestTicketRefundRule {
	s.RefundPart = &v
	return s
}

func (s *TicketSaveRequestTicketRefundRule) SetRefundType(v int) *TicketSaveRequestTicketRefundRule {
	s.RefundType = &v
	return s
}

func (s *TicketSaveRequestTicketRefundRule) SetAutoRefundTime(v int64) *TicketSaveRequestTicketRefundRule {
	s.AutoRefundTime = &v
	return s
}

func (s *TicketSaveRequestTicketRefundRule) SetAutoVerifyTime(v int64) *TicketSaveRequestTicketRefundRule {
	s.AutoVerifyTime = &v
	return s
}

type TicketSaveRequestTicketRefundRuleRefundDetailsItem struct {
	RefundFee     *int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty" require:"true"`
	RefundFeeType *int   `json:"refund_fee_type,omitempty" xml:"refund_fee_type,omitempty" require:"true"`
	RefundTime    *int64 `json:"refund_time,omitempty" xml:"refund_time,omitempty" require:"true"`
}

func (s TicketSaveRequestTicketRefundRuleRefundDetailsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketRefundRuleRefundDetailsItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketRefundRuleRefundDetailsItem) SetRefundFee(v int64) *TicketSaveRequestTicketRefundRuleRefundDetailsItem {
	s.RefundFee = &v
	return s
}

func (s *TicketSaveRequestTicketRefundRuleRefundDetailsItem) SetRefundFeeType(v int) *TicketSaveRequestTicketRefundRuleRefundDetailsItem {
	s.RefundFeeType = &v
	return s
}

func (s *TicketSaveRequestTicketRefundRuleRefundDetailsItem) SetRefundTime(v int64) *TicketSaveRequestTicketRefundRuleRefundDetailsItem {
	s.RefundTime = &v
	return s
}

type TicketSaveRequestTicketRegion struct {
	RestrictFlag      *bool                                                 `json:"restrict_flag,omitempty" xml:"restrict_flag,omitempty"`
	RestrictDistricts []*TicketSaveRequestTicketRegionRestrictDistrictsItem `json:"restrict_districts,omitempty" xml:"restrict_districts,omitempty" type:"Repeated"`
}

func (s TicketSaveRequestTicketRegion) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketRegion) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketRegion) SetRestrictFlag(v bool) *TicketSaveRequestTicketRegion {
	s.RestrictFlag = &v
	return s
}

func (s *TicketSaveRequestTicketRegion) SetRestrictDistricts(v []*TicketSaveRequestTicketRegionRestrictDistrictsItem) *TicketSaveRequestTicketRegion {
	s.RestrictDistricts = v
	return s
}

type TicketSaveRequestTicketRegionRestrictDistrictsItem struct {
	DistrictType     *int                                                                      `json:"district_type,omitempty" xml:"district_type,omitempty"`
	ProvinceInfoList []*TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem `json:"province_info_list,omitempty" xml:"province_info_list,omitempty" type:"Repeated"`
}

func (s TicketSaveRequestTicketRegionRestrictDistrictsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketRegionRestrictDistrictsItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketRegionRestrictDistrictsItem) SetDistrictType(v int) *TicketSaveRequestTicketRegionRestrictDistrictsItem {
	s.DistrictType = &v
	return s
}

func (s *TicketSaveRequestTicketRegionRestrictDistrictsItem) SetProvinceInfoList(v []*TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem) *TicketSaveRequestTicketRegionRestrictDistrictsItem {
	s.ProvinceInfoList = v
	return s
}

type TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem struct {
	CityCodeList []*string `json:"city_code_list,omitempty" xml:"city_code_list,omitempty" type:"Repeated"`
	ProvinceCode *string   `json:"province_code,omitempty" xml:"province_code,omitempty"`
}

func (s TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem) SetCityCodeList(v []*string) *TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem {
	s.CityCodeList = v
	return s
}

func (s *TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem) SetProvinceCode(v string) *TicketSaveRequestTicketRegionRestrictDistrictsItemProvinceInfoListItem {
	s.ProvinceCode = &v
	return s
}

type TicketSaveRequestTicketStockSyncRule struct {
	SyncOn *bool `json:"sync_on,omitempty" xml:"sync_on,omitempty"`
}

func (s TicketSaveRequestTicketStockSyncRule) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketStockSyncRule) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketStockSyncRule) SetSyncOn(v bool) *TicketSaveRequestTicketStockSyncRule {
	s.SyncOn = &v
	return s
}

type TicketSaveRequestTicketTicketRule struct {
	ChangeLatestTime    *string `json:"change_latest_time,omitempty" xml:"change_latest_time,omitempty"`
	ChangeTicketAddress *string `json:"change_ticket_address,omitempty" xml:"change_ticket_address,omitempty"`
	CodeExtraNote       *string `json:"code_extra_note,omitempty" xml:"code_extra_note,omitempty"`
	CodeNote            *string `json:"code_note,omitempty" xml:"code_note,omitempty"`
	CodeSendingInfo     []*int  `json:"code_sending_info,omitempty" xml:"code_sending_info,omitempty" require:"true" type:"Repeated"`
	CodeType            *int    `json:"code_type,omitempty" xml:"code_type,omitempty"`
	UrlType             *int    `json:"url_type,omitempty" xml:"url_type,omitempty"`
	ChangeEarliestTime  *string `json:"change_earliest_time,omitempty" xml:"change_earliest_time,omitempty"`
}

func (s TicketSaveRequestTicketTicketRule) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketTicketRule) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketTicketRule) SetChangeLatestTime(v string) *TicketSaveRequestTicketTicketRule {
	s.ChangeLatestTime = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetChangeTicketAddress(v string) *TicketSaveRequestTicketTicketRule {
	s.ChangeTicketAddress = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetCodeExtraNote(v string) *TicketSaveRequestTicketTicketRule {
	s.CodeExtraNote = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetCodeNote(v string) *TicketSaveRequestTicketTicketRule {
	s.CodeNote = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetCodeSendingInfo(v []*int) *TicketSaveRequestTicketTicketRule {
	s.CodeSendingInfo = v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetCodeType(v int) *TicketSaveRequestTicketTicketRule {
	s.CodeType = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetUrlType(v int) *TicketSaveRequestTicketTicketRule {
	s.UrlType = &v
	return s
}

func (s *TicketSaveRequestTicketTicketRule) SetChangeEarliestTime(v string) *TicketSaveRequestTicketTicketRule {
	s.ChangeEarliestTime = &v
	return s
}

type TicketSaveRequestTicketTicketSpecificationsItem struct {
	TicketSession *TicketSaveRequestTicketTicketSpecificationsItemTicketSession `json:"ticket_session,omitempty" xml:"ticket_session,omitempty"`
	FreeMerchant  *bool                                                         `json:"free_merchant,omitempty" xml:"free_merchant,omitempty"`
	SettleType    *int                                                          `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	SkuName       *string                                                       `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuOutId      *string                                                       `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	TicketArea    *string                                                       `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
	TicketSeat    *string                                                       `json:"ticket_seat,omitempty" xml:"ticket_seat,omitempty"`
}

func (s TicketSaveRequestTicketTicketSpecificationsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketTicketSpecificationsItem) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetTicketSession(v *TicketSaveRequestTicketTicketSpecificationsItemTicketSession) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.TicketSession = v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetFreeMerchant(v bool) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.FreeMerchant = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetSettleType(v int) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.SettleType = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetSkuName(v string) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.SkuName = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetSkuOutId(v string) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.SkuOutId = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetTicketArea(v string) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.TicketArea = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItem) SetTicketSeat(v string) *TicketSaveRequestTicketTicketSpecificationsItem {
	s.TicketSeat = &v
	return s
}

type TicketSaveRequestTicketTicketSpecificationsItemTicketSession struct {
	TicketSessionName *string `json:"ticket_session_name,omitempty" xml:"ticket_session_name,omitempty" require:"true"`
	TicketSessionTime *string `json:"ticket_session_time,omitempty" xml:"ticket_session_time,omitempty"`
}

func (s TicketSaveRequestTicketTicketSpecificationsItemTicketSession) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveRequestTicketTicketSpecificationsItemTicketSession) GoString() string {
	return s.String()
}

func (s *TicketSaveRequestTicketTicketSpecificationsItemTicketSession) SetTicketSessionName(v string) *TicketSaveRequestTicketTicketSpecificationsItemTicketSession {
	s.TicketSessionName = &v
	return s
}

func (s *TicketSaveRequestTicketTicketSpecificationsItemTicketSession) SetTicketSessionTime(v string) *TicketSaveRequestTicketTicketSpecificationsItemTicketSession {
	s.TicketSessionTime = &v
	return s
}

type TicketSaveResponse struct {
	Data  *TicketSaveResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TicketSaveResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TicketSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveResponse) GoString() string {
	return s.String()
}

func (s *TicketSaveResponse) SetData(v *TicketSaveResponseData) *TicketSaveResponse {
	s.Data = v
	return s
}

func (s *TicketSaveResponse) SetExtra(v *TicketSaveResponseExtra) *TicketSaveResponse {
	s.Extra = v
	return s
}

type TicketSaveResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ProductId     *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	TicketTypeId  *string `json:"ticket_type_id,omitempty" xml:"ticket_type_id,omitempty" require:"true"`
}

func (s TicketSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveResponseData) GoString() string {
	return s.String()
}

func (s *TicketSaveResponseData) SetGwErrorCode(v int32) *TicketSaveResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TicketSaveResponseData) SetGwDescription(v string) *TicketSaveResponseData {
	s.GwDescription = &v
	return s
}

func (s *TicketSaveResponseData) SetProductId(v string) *TicketSaveResponseData {
	s.ProductId = &v
	return s
}

func (s *TicketSaveResponseData) SetTicketTypeId(v string) *TicketSaveResponseData {
	s.TicketTypeId = &v
	return s
}

type TicketSaveResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TicketSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TicketSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *TicketSaveResponseExtra) SetLogid(v string) *TicketSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *TicketSaveResponseExtra) SetNow(v int64) *TicketSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *TicketSaveResponseExtra) SetSubDescription(v string) *TicketSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TicketSaveResponseExtra) SetSubErrorCode(v int32) *TicketSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TicketSaveResponseExtra) SetDescription(v string) *TicketSaveResponseExtra {
	s.Description = &v
	return s
}

func (s *TicketSaveResponseExtra) SetErrorCode(v int32) *TicketSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

type TicketSubmitRequest struct {
	Header      map[string]*string       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId   *int64                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Data        *TicketSubmitRequestData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TicketSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequest) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequest) SetHeader(v map[string]*string) *TicketSubmitRequest {
	s.Header = v
	return s
}

func (s *TicketSubmitRequest) SetAccessToken(v string) *TicketSubmitRequest {
	s.AccessToken = &v
	return s
}

func (s *TicketSubmitRequest) SetAccountId(v int64) *TicketSubmitRequest {
	s.AccountId = &v
	return s
}

func (s *TicketSubmitRequest) SetData(v *TicketSubmitRequestData) *TicketSubmitRequest {
	s.Data = v
	return s
}

type TicketSubmitRequestData struct {
	Others       []*TicketSubmitRequestDataOthersItem       `json:"others,omitempty" xml:"others,omitempty" type:"Repeated"`
	PoiId        *int64                                     `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Subject      *TicketSubmitRequestDataSubject            `json:"subject,omitempty" xml:"subject,omitempty"`
	Applications []*TicketSubmitRequestDataApplicationsItem `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	Legal        *TicketSubmitRequestDataLegal              `json:"legal,omitempty" xml:"legal,omitempty"`
}

func (s TicketSubmitRequestData) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequestData) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequestData) SetOthers(v []*TicketSubmitRequestDataOthersItem) *TicketSubmitRequestData {
	s.Others = v
	return s
}

func (s *TicketSubmitRequestData) SetPoiId(v int64) *TicketSubmitRequestData {
	s.PoiId = &v
	return s
}

func (s *TicketSubmitRequestData) SetSubject(v *TicketSubmitRequestDataSubject) *TicketSubmitRequestData {
	s.Subject = v
	return s
}

func (s *TicketSubmitRequestData) SetApplications(v []*TicketSubmitRequestDataApplicationsItem) *TicketSubmitRequestData {
	s.Applications = v
	return s
}

func (s *TicketSubmitRequestData) SetLegal(v *TicketSubmitRequestDataLegal) *TicketSubmitRequestData {
	s.Legal = v
	return s
}

type TicketSubmitRequestDataApplicationsItem struct {
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

func (s TicketSubmitRequestDataApplicationsItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequestDataApplicationsItem) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequestDataApplicationsItem) SetImageUrl(v string) *TicketSubmitRequestDataApplicationsItem {
	s.ImageUrl = &v
	return s
}

type TicketSubmitRequestDataLegal struct {
	IdCardNo       *string `json:"id_card_no,omitempty" xml:"id_card_no,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	IdCardBackUrl  *string `json:"id_card_back_url,omitempty" xml:"id_card_back_url,omitempty"`
	IdCardFrontUrl *string `json:"id_card_front_url,omitempty" xml:"id_card_front_url,omitempty"`
}

func (s TicketSubmitRequestDataLegal) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequestDataLegal) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequestDataLegal) SetIdCardNo(v string) *TicketSubmitRequestDataLegal {
	s.IdCardNo = &v
	return s
}

func (s *TicketSubmitRequestDataLegal) SetName(v string) *TicketSubmitRequestDataLegal {
	s.Name = &v
	return s
}

func (s *TicketSubmitRequestDataLegal) SetIdCardBackUrl(v string) *TicketSubmitRequestDataLegal {
	s.IdCardBackUrl = &v
	return s
}

func (s *TicketSubmitRequestDataLegal) SetIdCardFrontUrl(v string) *TicketSubmitRequestDataLegal {
	s.IdCardFrontUrl = &v
	return s
}

type TicketSubmitRequestDataOthersItem struct {
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
}

func (s TicketSubmitRequestDataOthersItem) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequestDataOthersItem) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequestDataOthersItem) SetImageUrl(v string) *TicketSubmitRequestDataOthersItem {
	s.ImageUrl = &v
	return s
}

type TicketSubmitRequestDataSubject struct {
	Expiration      *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	LegalPersonName *string   `json:"legal_person_name,omitempty" xml:"legal_person_name,omitempty"`
	LicenseId       *string   `json:"license_id,omitempty" xml:"license_id,omitempty"`
	LicenseType     *int64    `json:"license_type,omitempty" xml:"license_type,omitempty"`
	LicenseUrls     []*string `json:"license_urls,omitempty" xml:"license_urls,omitempty" type:"Repeated"`
	CompanyName     *string   `json:"company_name,omitempty" xml:"company_name,omitempty"`
}

func (s TicketSubmitRequestDataSubject) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitRequestDataSubject) GoString() string {
	return s.String()
}

func (s *TicketSubmitRequestDataSubject) SetExpiration(v string) *TicketSubmitRequestDataSubject {
	s.Expiration = &v
	return s
}

func (s *TicketSubmitRequestDataSubject) SetLegalPersonName(v string) *TicketSubmitRequestDataSubject {
	s.LegalPersonName = &v
	return s
}

func (s *TicketSubmitRequestDataSubject) SetLicenseId(v string) *TicketSubmitRequestDataSubject {
	s.LicenseId = &v
	return s
}

func (s *TicketSubmitRequestDataSubject) SetLicenseType(v int64) *TicketSubmitRequestDataSubject {
	s.LicenseType = &v
	return s
}

func (s *TicketSubmitRequestDataSubject) SetLicenseUrls(v []*string) *TicketSubmitRequestDataSubject {
	s.LicenseUrls = v
	return s
}

func (s *TicketSubmitRequestDataSubject) SetCompanyName(v string) *TicketSubmitRequestDataSubject {
	s.CompanyName = &v
	return s
}

type TicketSubmitResponse struct {
	Data  *TicketSubmitResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TicketSubmitResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TicketSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitResponse) GoString() string {
	return s.String()
}

func (s *TicketSubmitResponse) SetData(v *TicketSubmitResponseData) *TicketSubmitResponse {
	s.Data = v
	return s
}

func (s *TicketSubmitResponse) SetExtra(v *TicketSubmitResponseExtra) *TicketSubmitResponse {
	s.Extra = v
	return s
}

type TicketSubmitResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PoiId         *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	TicketId      *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TicketSubmitResponseData) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitResponseData) GoString() string {
	return s.String()
}

func (s *TicketSubmitResponseData) SetGwDescription(v string) *TicketSubmitResponseData {
	s.GwDescription = &v
	return s
}

func (s *TicketSubmitResponseData) SetPoiId(v int64) *TicketSubmitResponseData {
	s.PoiId = &v
	return s
}

func (s *TicketSubmitResponseData) SetTicketId(v string) *TicketSubmitResponseData {
	s.TicketId = &v
	return s
}

func (s *TicketSubmitResponseData) SetGwErrorCode(v int32) *TicketSubmitResponseData {
	s.GwErrorCode = &v
	return s
}

type TicketSubmitResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s TicketSubmitResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TicketSubmitResponseExtra) GoString() string {
	return s.String()
}

func (s *TicketSubmitResponseExtra) SetDescription(v string) *TicketSubmitResponseExtra {
	s.Description = &v
	return s
}

func (s *TicketSubmitResponseExtra) SetErrorCode(v int32) *TicketSubmitResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TicketSubmitResponseExtra) SetLogid(v string) *TicketSubmitResponseExtra {
	s.Logid = &v
	return s
}

func (s *TicketSubmitResponseExtra) SetNow(v int64) *TicketSubmitResponseExtra {
	s.Now = &v
	return s
}

func (s *TicketSubmitResponseExtra) SetSubDescription(v string) *TicketSubmitResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TicketSubmitResponseExtra) SetSubErrorCode(v int32) *TicketSubmitResponseExtra {
	s.SubErrorCode = &v
	return s
}

type ToolkitButtonWhiteSettingRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ButtonType  *int               `json:"button_type,omitempty" xml:"button_type,omitempty" require:"true"`
	GrayRate    *int64             `json:"gray_rate,omitempty" xml:"gray_rate,omitempty"`
	OpenAll     *bool              `json:"open_all,omitempty" xml:"open_all,omitempty"`
	UidList     []*int64           `json:"uid_list,omitempty" xml:"uid_list,omitempty" type:"Repeated"`
}

func (s ToolkitButtonWhiteSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitButtonWhiteSettingRequest) GoString() string {
	return s.String()
}

func (s *ToolkitButtonWhiteSettingRequest) SetHeader(v map[string]*string) *ToolkitButtonWhiteSettingRequest {
	s.Header = v
	return s
}

func (s *ToolkitButtonWhiteSettingRequest) SetAccessToken(v string) *ToolkitButtonWhiteSettingRequest {
	s.AccessToken = &v
	return s
}

func (s *ToolkitButtonWhiteSettingRequest) SetButtonType(v int) *ToolkitButtonWhiteSettingRequest {
	s.ButtonType = &v
	return s
}

func (s *ToolkitButtonWhiteSettingRequest) SetGrayRate(v int64) *ToolkitButtonWhiteSettingRequest {
	s.GrayRate = &v
	return s
}

func (s *ToolkitButtonWhiteSettingRequest) SetOpenAll(v bool) *ToolkitButtonWhiteSettingRequest {
	s.OpenAll = &v
	return s
}

func (s *ToolkitButtonWhiteSettingRequest) SetUidList(v []*int64) *ToolkitButtonWhiteSettingRequest {
	s.UidList = v
	return s
}

type ToolkitButtonWhiteSettingResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ToolkitButtonWhiteSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitButtonWhiteSettingResponse) GoString() string {
	return s.String()
}

func (s *ToolkitButtonWhiteSettingResponse) SetErrMsg(v string) *ToolkitButtonWhiteSettingResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitButtonWhiteSettingResponse) SetErrNo(v int32) *ToolkitButtonWhiteSettingResponse {
	s.ErrNo = &v
	return s
}

func (s *ToolkitButtonWhiteSettingResponse) SetLogId(v string) *ToolkitButtonWhiteSettingResponse {
	s.LogId = &v
	return s
}

type ToolkitChangeLockStatusRequest struct {
	OperationType *int                                               `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	OrderInfoList []*ToolkitChangeLockStatusRequestOrderInfoListItem `json:"order_info_list,omitempty" xml:"order_info_list,omitempty" type:"Repeated"`
	Header        map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ToolkitChangeLockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusRequest) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusRequest) SetOperationType(v int) *ToolkitChangeLockStatusRequest {
	s.OperationType = &v
	return s
}

func (s *ToolkitChangeLockStatusRequest) SetOrderInfoList(v []*ToolkitChangeLockStatusRequestOrderInfoListItem) *ToolkitChangeLockStatusRequest {
	s.OrderInfoList = v
	return s
}

func (s *ToolkitChangeLockStatusRequest) SetHeader(v map[string]*string) *ToolkitChangeLockStatusRequest {
	s.Header = v
	return s
}

func (s *ToolkitChangeLockStatusRequest) SetAccessToken(v string) *ToolkitChangeLockStatusRequest {
	s.AccessToken = &v
	return s
}

type ToolkitChangeLockStatusRequestOrderInfoListItem struct {
	CertificateInfoList []*ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	LockKey             *string                                                                   `json:"lock_key,omitempty" xml:"lock_key,omitempty"`
	OrderId             *string                                                                   `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItem) SetCertificateInfoList(v []*ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) *ToolkitChangeLockStatusRequestOrderInfoListItem {
	s.CertificateInfoList = v
	return s
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItem) SetLockKey(v string) *ToolkitChangeLockStatusRequestOrderInfoListItem {
	s.LockKey = &v
	return s
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItem) SetOrderId(v string) *ToolkitChangeLockStatusRequestOrderInfoListItem {
	s.OrderId = &v
	return s
}

type ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem struct {
	BookInfo          *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo          `json:"book_info,omitempty" xml:"book_info,omitempty"`
	CertificateId     *string                                                                                  `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	TimesCardLockInfo *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo `json:"times_card_lock_info,omitempty" xml:"times_card_lock_info,omitempty"`
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) SetBookInfo(v *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem {
	s.BookInfo = v
	return s
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) SetCertificateId(v string) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem) SetTimesCardLockInfo(v *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItem {
	s.TimesCardLockInfo = v
	return s
}

type ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo struct {
	BookEndTime   *int64 `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
	BookStartTime *int64 `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo) SetBookEndTime(v int64) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo {
	s.BookEndTime = &v
	return s
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo) SetBookStartTime(v int64) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemBookInfo {
	s.BookStartTime = &v
	return s
}

type ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo struct {
	LockCount *int32 `json:"lock_count,omitempty" xml:"lock_count,omitempty"`
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo) SetLockCount(v int32) *ToolkitChangeLockStatusRequestOrderInfoListItemCertificateInfoListItemTimesCardLockInfo {
	s.LockCount = &v
	return s
}

type ToolkitChangeLockStatusResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ToolkitChangeLockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitChangeLockStatusResponse) GoString() string {
	return s.String()
}

func (s *ToolkitChangeLockStatusResponse) SetErrMsg(v string) *ToolkitChangeLockStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitChangeLockStatusResponse) SetErrNo(v int32) *ToolkitChangeLockStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *ToolkitChangeLockStatusResponse) SetLogId(v string) *ToolkitChangeLockStatusResponse {
	s.LogId = &v
	return s
}

type ToolkitPushServiceDoneRequest struct {
	OrderId      *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	VerifyIdList []*string          `json:"verify_id_list,omitempty" xml:"verify_id_list,omitempty" type:"Repeated"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ToolkitPushServiceDoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitPushServiceDoneRequest) GoString() string {
	return s.String()
}

func (s *ToolkitPushServiceDoneRequest) SetOrderId(v string) *ToolkitPushServiceDoneRequest {
	s.OrderId = &v
	return s
}

func (s *ToolkitPushServiceDoneRequest) SetVerifyIdList(v []*string) *ToolkitPushServiceDoneRequest {
	s.VerifyIdList = v
	return s
}

func (s *ToolkitPushServiceDoneRequest) SetHeader(v map[string]*string) *ToolkitPushServiceDoneRequest {
	s.Header = v
	return s
}

func (s *ToolkitPushServiceDoneRequest) SetAccessToken(v string) *ToolkitPushServiceDoneRequest {
	s.AccessToken = &v
	return s
}

type ToolkitPushServiceDoneResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ToolkitPushServiceDoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitPushServiceDoneResponse) GoString() string {
	return s.String()
}

func (s *ToolkitPushServiceDoneResponse) SetLogId(v string) *ToolkitPushServiceDoneResponse {
	s.LogId = &v
	return s
}

func (s *ToolkitPushServiceDoneResponse) SetErrMsg(v string) *ToolkitPushServiceDoneResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitPushServiceDoneResponse) SetErrNo(v int32) *ToolkitPushServiceDoneResponse {
	s.ErrNo = &v
	return s
}

type ToolkitQueryCertificateInfoRequest struct {
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	CertificateIdList []*string          `json:"certificate_id_list,omitempty" xml:"certificate_id_list,omitempty" type:"Repeated"`
	OrderIdList       []*string          `json:"order_id_list,omitempty" xml:"order_id_list,omitempty" type:"Repeated"`
}

func (s ToolkitQueryCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoRequest) SetHeader(v map[string]*string) *ToolkitQueryCertificateInfoRequest {
	s.Header = v
	return s
}

func (s *ToolkitQueryCertificateInfoRequest) SetAccessToken(v string) *ToolkitQueryCertificateInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *ToolkitQueryCertificateInfoRequest) SetCertificateIdList(v []*string) *ToolkitQueryCertificateInfoRequest {
	s.CertificateIdList = v
	return s
}

func (s *ToolkitQueryCertificateInfoRequest) SetOrderIdList(v []*string) *ToolkitQueryCertificateInfoRequest {
	s.OrderIdList = v
	return s
}

type ToolkitQueryCertificateInfoResponse struct {
	ErrNo  *int32                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ToolkitQueryCertificateInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s ToolkitQueryCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponse) SetErrNo(v int32) *ToolkitQueryCertificateInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponse) SetLogId(v string) *ToolkitQueryCertificateInfoResponse {
	s.LogId = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponse) SetData(v *ToolkitQueryCertificateInfoResponseData) *ToolkitQueryCertificateInfoResponse {
	s.Data = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponse) SetErrMsg(v string) *ToolkitQueryCertificateInfoResponse {
	s.ErrMsg = &v
	return s
}

type ToolkitQueryCertificateInfoResponseData struct {
	CertificateInfoList []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
}

func (s ToolkitQueryCertificateInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseData) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseData) SetCertificateInfoList(v []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) *ToolkitQueryCertificateInfoResponseData {
	s.CertificateInfoList = v
	return s
}

type ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem struct {
	CertificateId  *string                                                                             `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	LockInfoList   []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem   `json:"lock_info_list,omitempty" xml:"lock_info_list,omitempty" type:"Repeated"`
	OrderId        *string                                                                             `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Status         *int32                                                                              `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	TimesCardInfo  *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo        `json:"times_card_info,omitempty" xml:"times_card_info,omitempty"`
	VerifyInfoList []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem `json:"verify_info_list,omitempty" xml:"verify_info_list,omitempty" type:"Repeated"`
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetCertificateId(v string) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetLockInfoList(v []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.LockInfoList = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetOrderId(v string) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.OrderId = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetStatus(v int32) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.Status = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetTimesCardInfo(v *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.TimesCardInfo = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem) SetVerifyInfoList(v []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItem {
	s.VerifyInfoList = v
	return s
}

type ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem struct {
	LockKey *string `json:"lock_key,omitempty" xml:"lock_key,omitempty"`
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem) SetLockKey(v string) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemLockInfoListItem {
	s.LockKey = &v
	return s
}

type ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo struct {
	LadderTimesCardInfo   []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem `json:"ladder_times_card_info,omitempty" xml:"ladder_times_card_info,omitempty" type:"Repeated"`
	LockedTimes           *int64                                                                                                `json:"locked_times,omitempty" xml:"locked_times,omitempty" require:"true"`
	LockedTimesSerialNums []*int32                                                                                              `json:"locked_times_serial_nums,omitempty" xml:"locked_times_serial_nums,omitempty" type:"Repeated"`
	TotalTimes            *int64                                                                                                `json:"total_times,omitempty" xml:"total_times,omitempty" require:"true"`
	TotalTimesSerialNums  []*int32                                                                                              `json:"total_times_serial_nums,omitempty" xml:"total_times_serial_nums,omitempty" type:"Repeated"`
	UsedTimes             *int64                                                                                                `json:"used_times,omitempty" xml:"used_times,omitempty" require:"true"`
	UsedTimesSerialNums   []*int32                                                                                              `json:"used_times_serial_nums,omitempty" xml:"used_times_serial_nums,omitempty" type:"Repeated"`
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetLadderTimesCardInfo(v []*ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.LadderTimesCardInfo = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetLockedTimes(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.LockedTimes = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetLockedTimesSerialNums(v []*int32) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.LockedTimesSerialNums = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetTotalTimes(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.TotalTimes = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetTotalTimesSerialNums(v []*int32) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.TotalTimesSerialNums = v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetUsedTimes(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.UsedTimes = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo) SetUsedTimesSerialNums(v []*int32) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfo {
	s.UsedTimesSerialNums = v
	return s
}

type ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem struct {
	PayAmount      *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	Seq            *int32 `json:"seq,omitempty" xml:"seq,omitempty"`
	Status         *int   `json:"status,omitempty" xml:"status,omitempty"`
	Ticket         *int64 `json:"ticket,omitempty" xml:"ticket,omitempty"`
	CrossedAmount  *int64 `json:"crossed_amount,omitempty" xml:"crossed_amount,omitempty"`
	MerchantTicket *int64 `json:"merchant_ticket,omitempty" xml:"merchant_ticket,omitempty"`
	OriginalAmount *int64 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetPayAmount(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.PayAmount = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetSeq(v int32) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.Seq = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetStatus(v int) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.Status = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetTicket(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.Ticket = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetCrossedAmount(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.CrossedAmount = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetMerchantTicket(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.MerchantTicket = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem) SetOriginalAmount(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemTimesCardInfoLadderTimesCardInfoItem {
	s.OriginalAmount = &v
	return s
}

type ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem struct {
	VerifyId   *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
	VerifyTime *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem) SetVerifyId(v string) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem {
	s.VerifyId = &v
	return s
}

func (s *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem) SetVerifyTime(v int64) *ToolkitQueryCertificateInfoResponseDataCertificateInfoListItemVerifyInfoListItem {
	s.VerifyTime = &v
	return s
}

type ToolkitQueryTextRequest struct {
	TextType    *int               `json:"text_type,omitempty" xml:"text_type,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ToolkitQueryTextRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryTextRequest) GoString() string {
	return s.String()
}

func (s *ToolkitQueryTextRequest) SetTextType(v int) *ToolkitQueryTextRequest {
	s.TextType = &v
	return s
}

func (s *ToolkitQueryTextRequest) SetHeader(v map[string]*string) *ToolkitQueryTextRequest {
	s.Header = v
	return s
}

func (s *ToolkitQueryTextRequest) SetAccessToken(v string) *ToolkitQueryTextRequest {
	s.AccessToken = &v
	return s
}

type ToolkitQueryTextResponse struct {
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ToolkitQueryTextResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ToolkitQueryTextResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryTextResponse) GoString() string {
	return s.String()
}

func (s *ToolkitQueryTextResponse) SetErrNo(v int32) *ToolkitQueryTextResponse {
	s.ErrNo = &v
	return s
}

func (s *ToolkitQueryTextResponse) SetErrMsg(v string) *ToolkitQueryTextResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitQueryTextResponse) SetLogId(v string) *ToolkitQueryTextResponse {
	s.LogId = &v
	return s
}

func (s *ToolkitQueryTextResponse) SetData(v *ToolkitQueryTextResponseData) *ToolkitQueryTextResponse {
	s.Data = v
	return s
}

type ToolkitQueryTextResponseData struct {
	TextList []*ToolkitQueryTextResponseDataTextListItem `json:"text_list,omitempty" xml:"text_list,omitempty" type:"Repeated"`
}

func (s ToolkitQueryTextResponseData) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryTextResponseData) GoString() string {
	return s.String()
}

func (s *ToolkitQueryTextResponseData) SetTextList(v []*ToolkitQueryTextResponseDataTextListItem) *ToolkitQueryTextResponseData {
	s.TextList = v
	return s
}

type ToolkitQueryTextResponseDataTextListItem struct {
	TextContentList []*ToolkitQueryTextResponseDataTextListItemTextContentListItem `json:"text_content_list,omitempty" xml:"text_content_list,omitempty" type:"Repeated"`
	TextType        *int                                                           `json:"text_type,omitempty" xml:"text_type,omitempty" require:"true"`
}

func (s ToolkitQueryTextResponseDataTextListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryTextResponseDataTextListItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryTextResponseDataTextListItem) SetTextContentList(v []*ToolkitQueryTextResponseDataTextListItemTextContentListItem) *ToolkitQueryTextResponseDataTextListItem {
	s.TextContentList = v
	return s
}

func (s *ToolkitQueryTextResponseDataTextListItem) SetTextType(v int) *ToolkitQueryTextResponseDataTextListItem {
	s.TextType = &v
	return s
}

type ToolkitQueryTextResponseDataTextListItemTextContentListItem struct {
	TextContent *string `json:"text_content,omitempty" xml:"text_content,omitempty" require:"true"`
	TextId      *string `json:"text_id,omitempty" xml:"text_id,omitempty" require:"true"`
}

func (s ToolkitQueryTextResponseDataTextListItemTextContentListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitQueryTextResponseDataTextListItemTextContentListItem) GoString() string {
	return s.String()
}

func (s *ToolkitQueryTextResponseDataTextListItemTextContentListItem) SetTextContent(v string) *ToolkitQueryTextResponseDataTextListItemTextContentListItem {
	s.TextContent = &v
	return s
}

func (s *ToolkitQueryTextResponseDataTextListItemTextContentListItem) SetTextId(v string) *ToolkitQueryTextResponseDataTextListItemTextContentListItem {
	s.TextId = &v
	return s
}

type ToolkitUpdateMerchantConfRequest struct {
	BindBizType           *int                                                   `json:"bind_biz_type,omitempty" xml:"bind_biz_type,omitempty" require:"true"`
	DeliveryAppInfo       *ToolkitUpdateMerchantConfRequestDeliveryAppInfo       `json:"delivery_app_info,omitempty" xml:"delivery_app_info,omitempty"`
	ProductDoubleOpenInfo *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo `json:"product_double_open_info,omitempty" xml:"product_double_open_info,omitempty"`
	Header                map[string]*string                                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken           *string                                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId             *string                                                `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s ToolkitUpdateMerchantConfRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantConfRequest) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantConfRequest) SetBindBizType(v int) *ToolkitUpdateMerchantConfRequest {
	s.BindBizType = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequest) SetDeliveryAppInfo(v *ToolkitUpdateMerchantConfRequestDeliveryAppInfo) *ToolkitUpdateMerchantConfRequest {
	s.DeliveryAppInfo = v
	return s
}

func (s *ToolkitUpdateMerchantConfRequest) SetProductDoubleOpenInfo(v *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) *ToolkitUpdateMerchantConfRequest {
	s.ProductDoubleOpenInfo = v
	return s
}

func (s *ToolkitUpdateMerchantConfRequest) SetHeader(v map[string]*string) *ToolkitUpdateMerchantConfRequest {
	s.Header = v
	return s
}

func (s *ToolkitUpdateMerchantConfRequest) SetAccessToken(v string) *ToolkitUpdateMerchantConfRequest {
	s.AccessToken = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequest) SetAccountId(v string) *ToolkitUpdateMerchantConfRequest {
	s.AccountId = &v
	return s
}

type ToolkitUpdateMerchantConfRequestDeliveryAppInfo struct {
	DisplayMode      *int    `json:"display_mode,omitempty" xml:"display_mode,omitempty"`
	GuidanceTextId   *string `json:"guidance_text_id,omitempty" xml:"guidance_text_id,omitempty"`
	ButtonTextId     *string `json:"button_text_id,omitempty" xml:"button_text_id,omitempty"`
	DetailPageTextId *string `json:"detail_page_text_id,omitempty" xml:"detail_page_text_id,omitempty"`
}

func (s ToolkitUpdateMerchantConfRequestDeliveryAppInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantConfRequestDeliveryAppInfo) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantConfRequestDeliveryAppInfo) SetDisplayMode(v int) *ToolkitUpdateMerchantConfRequestDeliveryAppInfo {
	s.DisplayMode = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestDeliveryAppInfo) SetGuidanceTextId(v string) *ToolkitUpdateMerchantConfRequestDeliveryAppInfo {
	s.GuidanceTextId = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestDeliveryAppInfo) SetButtonTextId(v string) *ToolkitUpdateMerchantConfRequestDeliveryAppInfo {
	s.ButtonTextId = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestDeliveryAppInfo) SetDetailPageTextId(v string) *ToolkitUpdateMerchantConfRequestDeliveryAppInfo {
	s.DetailPageTextId = &v
	return s
}

type ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo struct {
	ArrivalStoreButtonTextId *string `json:"arrival_store_button_text_id,omitempty" xml:"arrival_store_button_text_id,omitempty"`
	DisplayMode              *int    `json:"display_mode,omitempty" xml:"display_mode,omitempty"`
	GuidanceTextId           *string `json:"guidance_text_id,omitempty" xml:"guidance_text_id,omitempty"`
	ArrivalHomeButtonTextId  *string `json:"arrival_home_button_text_id,omitempty" xml:"arrival_home_button_text_id,omitempty"`
}

func (s ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) SetArrivalStoreButtonTextId(v string) *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo {
	s.ArrivalStoreButtonTextId = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) SetDisplayMode(v int) *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo {
	s.DisplayMode = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) SetGuidanceTextId(v string) *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo {
	s.GuidanceTextId = &v
	return s
}

func (s *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo) SetArrivalHomeButtonTextId(v string) *ToolkitUpdateMerchantConfRequestProductDoubleOpenInfo {
	s.ArrivalHomeButtonTextId = &v
	return s
}

type ToolkitUpdateMerchantConfResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ToolkitUpdateMerchantConfResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantConfResponse) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantConfResponse) SetLogId(v string) *ToolkitUpdateMerchantConfResponse {
	s.LogId = &v
	return s
}

func (s *ToolkitUpdateMerchantConfResponse) SetErrMsg(v string) *ToolkitUpdateMerchantConfResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitUpdateMerchantConfResponse) SetErrNo(v int32) *ToolkitUpdateMerchantConfResponse {
	s.ErrNo = &v
	return s
}

type ToolkitUpdateMerchantPathRequest struct {
	AccountId    *string                                             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	BindBizType  *int                                                `json:"bind_biz_type,omitempty" xml:"bind_biz_type,omitempty" require:"true"`
	PathDataList []*ToolkitUpdateMerchantPathRequestPathDataListItem `json:"path_data_list,omitempty" xml:"path_data_list,omitempty" type:"Repeated"`
	Header       map[string]*string                                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ToolkitUpdateMerchantPathRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantPathRequest) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantPathRequest) SetAccountId(v string) *ToolkitUpdateMerchantPathRequest {
	s.AccountId = &v
	return s
}

func (s *ToolkitUpdateMerchantPathRequest) SetBindBizType(v int) *ToolkitUpdateMerchantPathRequest {
	s.BindBizType = &v
	return s
}

func (s *ToolkitUpdateMerchantPathRequest) SetPathDataList(v []*ToolkitUpdateMerchantPathRequestPathDataListItem) *ToolkitUpdateMerchantPathRequest {
	s.PathDataList = v
	return s
}

func (s *ToolkitUpdateMerchantPathRequest) SetHeader(v map[string]*string) *ToolkitUpdateMerchantPathRequest {
	s.Header = v
	return s
}

func (s *ToolkitUpdateMerchantPathRequest) SetAccessToken(v string) *ToolkitUpdateMerchantPathRequest {
	s.AccessToken = &v
	return s
}

type ToolkitUpdateMerchantPathRequestPathDataListItem struct {
	PathType *int    `json:"path_type,omitempty" xml:"path_type,omitempty" require:"true"`
	Path     *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
}

func (s ToolkitUpdateMerchantPathRequestPathDataListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantPathRequestPathDataListItem) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantPathRequestPathDataListItem) SetPathType(v int) *ToolkitUpdateMerchantPathRequestPathDataListItem {
	s.PathType = &v
	return s
}

func (s *ToolkitUpdateMerchantPathRequestPathDataListItem) SetPath(v string) *ToolkitUpdateMerchantPathRequestPathDataListItem {
	s.Path = &v
	return s
}

type ToolkitUpdateMerchantPathResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ToolkitUpdateMerchantPathResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitUpdateMerchantPathResponse) GoString() string {
	return s.String()
}

func (s *ToolkitUpdateMerchantPathResponse) SetLogId(v string) *ToolkitUpdateMerchantPathResponse {
	s.LogId = &v
	return s
}

func (s *ToolkitUpdateMerchantPathResponse) SetErrMsg(v string) *ToolkitUpdateMerchantPathResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitUpdateMerchantPathResponse) SetErrNo(v int32) *ToolkitUpdateMerchantPathResponse {
	s.ErrNo = &v
	return s
}

type ToolkitVerifyLocalCertificatesRequest struct {
	PoiId            *string                                                   `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	VerifyToken      *string                                                   `json:"verify_token,omitempty" xml:"verify_token,omitempty"`
	Header           map[string]*string                                        `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string                                                   `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId           *string                                                   `json:"open_id,omitempty" xml:"open_id,omitempty"`
	OrderEntrySchema *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema    `json:"order_entry_schema,omitempty" xml:"order_entry_schema,omitempty"`
	OrderInfoList    []*ToolkitVerifyLocalCertificatesRequestOrderInfoListItem `json:"order_info_list,omitempty" xml:"order_info_list,omitempty" type:"Repeated"`
}

func (s ToolkitVerifyLocalCertificatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetPoiId(v string) *ToolkitVerifyLocalCertificatesRequest {
	s.PoiId = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetVerifyToken(v string) *ToolkitVerifyLocalCertificatesRequest {
	s.VerifyToken = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetHeader(v map[string]*string) *ToolkitVerifyLocalCertificatesRequest {
	s.Header = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetAccessToken(v string) *ToolkitVerifyLocalCertificatesRequest {
	s.AccessToken = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetOpenId(v string) *ToolkitVerifyLocalCertificatesRequest {
	s.OpenId = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetOrderEntrySchema(v *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema) *ToolkitVerifyLocalCertificatesRequest {
	s.OrderEntrySchema = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequest) SetOrderInfoList(v []*ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) *ToolkitVerifyLocalCertificatesRequest {
	s.OrderInfoList = v
	return s
}

type ToolkitVerifyLocalCertificatesRequestOrderEntrySchema struct {
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ToolkitVerifyLocalCertificatesRequestOrderEntrySchema) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesRequestOrderEntrySchema) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema) SetParams(v string) *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema {
	s.Params = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema) SetPath(v string) *ToolkitVerifyLocalCertificatesRequestOrderEntrySchema {
	s.Path = &v
	return s
}

type ToolkitVerifyLocalCertificatesRequestOrderInfoListItem struct {
	OrderId             *string                                                                          `json:"order_id,omitempty" xml:"order_id,omitempty"`
	CertificateInfoList []*ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem `json:"certificate_info_list,omitempty" xml:"certificate_info_list,omitempty" type:"Repeated"`
	LockKey             *string                                                                          `json:"lock_key,omitempty" xml:"lock_key,omitempty"`
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) SetOrderId(v string) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem {
	s.OrderId = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) SetCertificateInfoList(v []*ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem {
	s.CertificateInfoList = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem) SetLockKey(v string) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItem {
	s.LockKey = &v
	return s
}

type ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem struct {
	BookInfo      *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo `json:"book_info,omitempty" xml:"book_info,omitempty"`
	CertificateId *string                                                                                `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem) SetBookInfo(v *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem {
	s.BookInfo = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem) SetCertificateId(v string) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItem {
	s.CertificateId = &v
	return s
}

type ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo struct {
	BookEndTime   *int64 `json:"book_end_time,omitempty" xml:"book_end_time,omitempty"`
	BookStartTime *int64 `json:"book_start_time,omitempty" xml:"book_start_time,omitempty"`
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo) SetBookEndTime(v int64) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo {
	s.BookEndTime = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo) SetBookStartTime(v int64) *ToolkitVerifyLocalCertificatesRequestOrderInfoListItemCertificateInfoListItemBookInfo {
	s.BookStartTime = &v
	return s
}

type ToolkitVerifyLocalCertificatesResponse struct {
	Data   *ToolkitVerifyLocalCertificatesResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ToolkitVerifyLocalCertificatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesResponse) SetData(v *ToolkitVerifyLocalCertificatesResponseData) *ToolkitVerifyLocalCertificatesResponse {
	s.Data = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponse) SetErrMsg(v string) *ToolkitVerifyLocalCertificatesResponse {
	s.ErrMsg = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponse) SetErrNo(v int32) *ToolkitVerifyLocalCertificatesResponse {
	s.ErrNo = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponse) SetLogId(v string) *ToolkitVerifyLocalCertificatesResponse {
	s.LogId = &v
	return s
}

type ToolkitVerifyLocalCertificatesResponseData struct {
	OrderVerifyResults []*ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem `json:"order_verify_results,omitempty" xml:"order_verify_results,omitempty" type:"Repeated"`
}

func (s ToolkitVerifyLocalCertificatesResponseData) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesResponseData) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesResponseData) SetOrderVerifyResults(v []*ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem) *ToolkitVerifyLocalCertificatesResponseData {
	s.OrderVerifyResults = v
	return s
}

type ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem struct {
	CertificateVerifyResults []*ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem `json:"certificate_verify_results,omitempty" xml:"certificate_verify_results,omitempty" require:"true" type:"Repeated"`
	OrderId                  *string                                                                                         `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem) SetCertificateVerifyResults(v []*ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem {
	s.CertificateVerifyResults = v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem) SetOrderId(v string) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItem {
	s.OrderId = &v
	return s
}

type ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem struct {
	VerifyTime    *int64  `json:"verify_time,omitempty" xml:"verify_time,omitempty" require:"true"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	ResultCode    *int32  `json:"result_code,omitempty" xml:"result_code,omitempty" require:"true"`
	ResultMsg     *string `json:"result_msg,omitempty" xml:"result_msg,omitempty" require:"true"`
	VerifyId      *string `json:"verify_id,omitempty" xml:"verify_id,omitempty" require:"true"`
}

func (s ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetVerifyTime(v int64) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.VerifyTime = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetCertificateId(v string) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetCode(v string) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.Code = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetResultCode(v int32) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.ResultCode = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetResultMsg(v string) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.ResultMsg = &v
	return s
}

func (s *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem) SetVerifyId(v string) *ToolkitVerifyLocalCertificatesResponseDataOrderVerifyResultsItemCertificateVerifyResultsItem {
	s.VerifyId = &v
	return s
}

type TradeOrderConfirmRequest struct {
	FulfilType    *int               `json:"fulfil_type,omitempty" xml:"fulfil_type,omitempty"`
	MerchantNotes *string            `json:"merchant_notes,omitempty" xml:"merchant_notes,omitempty"`
	Reason        *string            `json:"reason,omitempty" xml:"reason,omitempty"`
	RejectCode    *int32             `json:"reject_code,omitempty" xml:"reject_code,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BookId        *string            `json:"book_id,omitempty" xml:"book_id,omitempty" require:"true"`
	ConfirmResult *int               `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
}

func (s TradeOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *TradeOrderConfirmRequest) SetFulfilType(v int) *TradeOrderConfirmRequest {
	s.FulfilType = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetMerchantNotes(v string) *TradeOrderConfirmRequest {
	s.MerchantNotes = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetReason(v string) *TradeOrderConfirmRequest {
	s.Reason = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetRejectCode(v int32) *TradeOrderConfirmRequest {
	s.RejectCode = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetHeader(v map[string]*string) *TradeOrderConfirmRequest {
	s.Header = v
	return s
}

func (s *TradeOrderConfirmRequest) SetAccessToken(v string) *TradeOrderConfirmRequest {
	s.AccessToken = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetBookId(v string) *TradeOrderConfirmRequest {
	s.BookId = &v
	return s
}

func (s *TradeOrderConfirmRequest) SetConfirmResult(v int) *TradeOrderConfirmRequest {
	s.ConfirmResult = &v
	return s
}

type TradeOrderConfirmResponse struct {
	Extra *TradeOrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TradeOrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TradeOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *TradeOrderConfirmResponse) SetExtra(v *TradeOrderConfirmResponseExtra) *TradeOrderConfirmResponse {
	s.Extra = v
	return s
}

func (s *TradeOrderConfirmResponse) SetData(v *TradeOrderConfirmResponseData) *TradeOrderConfirmResponse {
	s.Data = v
	return s
}

type TradeOrderConfirmResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TradeOrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *TradeOrderConfirmResponseData) SetGwDescription(v string) *TradeOrderConfirmResponseData {
	s.GwDescription = &v
	return s
}

func (s *TradeOrderConfirmResponseData) SetGwErrorCode(v int32) *TradeOrderConfirmResponseData {
	s.GwErrorCode = &v
	return s
}

type TradeOrderConfirmResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s TradeOrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *TradeOrderConfirmResponseExtra) SetDescription(v string) *TradeOrderConfirmResponseExtra {
	s.Description = &v
	return s
}

func (s *TradeOrderConfirmResponseExtra) SetErrorCode(v int32) *TradeOrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TradeOrderConfirmResponseExtra) SetLogid(v string) *TradeOrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *TradeOrderConfirmResponseExtra) SetNow(v int64) *TradeOrderConfirmResponseExtra {
	s.Now = &v
	return s
}

func (s *TradeOrderConfirmResponseExtra) SetSubDescription(v string) *TradeOrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TradeOrderConfirmResponseExtra) SetSubErrorCode(v int32) *TradeOrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

type TradeOrderQueryRequest struct {
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UpdateOrderEndTime   *int64             `json:"update_order_end_time,omitempty" xml:"update_order_end_time,omitempty"`
	Cursor               []*string          `json:"cursor,omitempty" xml:"cursor,omitempty" type:"Repeated"`
	PageSize             *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	OrderStatus          *int32             `json:"order_status,omitempty" xml:"order_status,omitempty"`
	CreateOrderEndTime   *int64             `json:"create_order_end_time,omitempty" xml:"create_order_end_time,omitempty"`
	UpdateOrderStartTime *int64             `json:"update_order_start_time,omitempty" xml:"update_order_start_time,omitempty"`
	CreateOrderStartTime *int64             `json:"create_order_start_time,omitempty" xml:"create_order_start_time,omitempty"`
	ExtOrderId           *string            `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	OrderId              *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OpenId               *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	GetSecretNumber      *bool              `json:"get_secret_number,omitempty" xml:"get_secret_number,omitempty"`
	PageNum              *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	AccountId            *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s TradeOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryRequest) SetHeader(v map[string]*string) *TradeOrderQueryRequest {
	s.Header = v
	return s
}

func (s *TradeOrderQueryRequest) SetAccessToken(v string) *TradeOrderQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *TradeOrderQueryRequest) SetUpdateOrderEndTime(v int64) *TradeOrderQueryRequest {
	s.UpdateOrderEndTime = &v
	return s
}

func (s *TradeOrderQueryRequest) SetCursor(v []*string) *TradeOrderQueryRequest {
	s.Cursor = v
	return s
}

func (s *TradeOrderQueryRequest) SetPageSize(v int32) *TradeOrderQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TradeOrderQueryRequest) SetOrderStatus(v int32) *TradeOrderQueryRequest {
	s.OrderStatus = &v
	return s
}

func (s *TradeOrderQueryRequest) SetCreateOrderEndTime(v int64) *TradeOrderQueryRequest {
	s.CreateOrderEndTime = &v
	return s
}

func (s *TradeOrderQueryRequest) SetUpdateOrderStartTime(v int64) *TradeOrderQueryRequest {
	s.UpdateOrderStartTime = &v
	return s
}

func (s *TradeOrderQueryRequest) SetCreateOrderStartTime(v int64) *TradeOrderQueryRequest {
	s.CreateOrderStartTime = &v
	return s
}

func (s *TradeOrderQueryRequest) SetExtOrderId(v string) *TradeOrderQueryRequest {
	s.ExtOrderId = &v
	return s
}

func (s *TradeOrderQueryRequest) SetOrderId(v string) *TradeOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *TradeOrderQueryRequest) SetOpenId(v string) *TradeOrderQueryRequest {
	s.OpenId = &v
	return s
}

func (s *TradeOrderQueryRequest) SetGetSecretNumber(v bool) *TradeOrderQueryRequest {
	s.GetSecretNumber = &v
	return s
}

func (s *TradeOrderQueryRequest) SetPageNum(v int32) *TradeOrderQueryRequest {
	s.PageNum = &v
	return s
}

func (s *TradeOrderQueryRequest) SetAccountId(v string) *TradeOrderQueryRequest {
	s.AccountId = &v
	return s
}

type TradeOrderQueryResponse struct {
	Data  *TradeOrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TradeOrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s TradeOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponse) SetData(v *TradeOrderQueryResponseData) *TradeOrderQueryResponse {
	s.Data = v
	return s
}

func (s *TradeOrderQueryResponse) SetExtra(v *TradeOrderQueryResponseExtra) *TradeOrderQueryResponse {
	s.Extra = v
	return s
}

type TradeOrderQueryResponseData struct {
	Orders        []*TradeOrderQueryResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
	Page          *TradeOrderQueryResponseDataPage         `json:"page,omitempty" xml:"page,omitempty"`
	SearchAfter   *TradeOrderQueryResponseDataSearchAfter  `json:"search_after,omitempty" xml:"search_after,omitempty"`
	GwErrorCode   *int32                                   `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                  `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseData) SetOrders(v []*TradeOrderQueryResponseDataOrdersItem) *TradeOrderQueryResponseData {
	s.Orders = v
	return s
}

func (s *TradeOrderQueryResponseData) SetPage(v *TradeOrderQueryResponseDataPage) *TradeOrderQueryResponseData {
	s.Page = v
	return s
}

func (s *TradeOrderQueryResponseData) SetSearchAfter(v *TradeOrderQueryResponseDataSearchAfter) *TradeOrderQueryResponseData {
	s.SearchAfter = v
	return s
}

func (s *TradeOrderQueryResponseData) SetGwErrorCode(v int32) *TradeOrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TradeOrderQueryResponseData) SetGwDescription(v string) *TradeOrderQueryResponseData {
	s.GwDescription = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItem struct {
	Certificate         []*TradeOrderQueryResponseDataOrdersItemCertificateItem         `json:"certificate,omitempty" xml:"certificate,omitempty" type:"Repeated"`
	SkuId               *string                                                         `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	DeliveryInfo        *TradeOrderQueryResponseDataOrdersItemDeliveryInfo              `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
	IntentionPoiId      *string                                                         `json:"intention_poi_id,omitempty" xml:"intention_poi_id,omitempty"`
	DiscountAmount      *int64                                                          `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	PayTime             *int64                                                          `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	PaymentDiscount     *int32                                                          `json:"payment_discount,omitempty" xml:"payment_discount,omitempty"`
	AnchorId            *int64                                                          `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	Products            []*TradeOrderQueryResponseDataOrdersItemProductsItem            `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	MerchantInfo        *TradeOrderQueryResponseDataOrdersItemMerchantInfo              `json:"merchant_info,omitempty" xml:"merchant_info,omitempty"`
	OrderId             *string                                                         `json:"order_id,omitempty" xml:"order_id,omitempty"`
	ReceiptAmount       *int64                                                          `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	Contacts            []*TradeOrderQueryResponseDataOrdersItemContactsItem            `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	Count               *int32                                                          `json:"count,omitempty" xml:"count,omitempty"`
	OrderStatus         *int32                                                          `json:"order_status,omitempty" xml:"order_status,omitempty"`
	SubOrderAmountInfos []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem `json:"sub_order_amount_infos,omitempty" xml:"sub_order_amount_infos,omitempty" type:"Repeated"`
	OpenId              *string                                                         `json:"open_id,omitempty" xml:"open_id,omitempty"`
	Discounts           []*TradeOrderQueryResponseDataOrdersItemDiscountsItem           `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	CreateOrderTime     *int64                                                          `json:"create_order_time,omitempty" xml:"create_order_time,omitempty"`
	SourceOrderId       *string                                                         `json:"source_order_id,omitempty" xml:"source_order_id,omitempty"`
	AmountInfo          *TradeOrderQueryResponseDataOrdersItemAmountInfo                `json:"amount_info,omitempty" xml:"amount_info,omitempty"`
	UpdateOrderTime     *int64                                                          `json:"update_order_time,omitempty" xml:"update_order_time,omitempty"`
	CraftsmanUid        *string                                                         `json:"craftsman_uid,omitempty" xml:"craftsman_uid,omitempty"`
	SkuName             *string                                                         `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	OriginalAmount      *int32                                                          `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	ThirdSkuId          *string                                                         `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	PoiId               *string                                                         `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PayAmount           *int32                                                          `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	RoomId              *string                                                         `json:"room_id,omitempty" xml:"room_id,omitempty"`
	OrderType           *int32                                                          `json:"order_type,omitempty" xml:"order_type,omitempty"`
	Poi                 *TradeOrderQueryResponseDataOrdersItemPoi                       `json:"poi,omitempty" xml:"poi,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetCertificate(v []*TradeOrderQueryResponseDataOrdersItemCertificateItem) *TradeOrderQueryResponseDataOrdersItem {
	s.Certificate = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetSkuId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.SkuId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetDeliveryInfo(v *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) *TradeOrderQueryResponseDataOrdersItem {
	s.DeliveryInfo = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetIntentionPoiId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.IntentionPoiId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.DiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetPayTime(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.PayTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetPaymentDiscount(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.PaymentDiscount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetAnchorId(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.AnchorId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetProducts(v []*TradeOrderQueryResponseDataOrdersItemProductsItem) *TradeOrderQueryResponseDataOrdersItem {
	s.Products = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetMerchantInfo(v *TradeOrderQueryResponseDataOrdersItemMerchantInfo) *TradeOrderQueryResponseDataOrdersItem {
	s.MerchantInfo = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetOrderId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.OrderId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetReceiptAmount(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.ReceiptAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetContacts(v []*TradeOrderQueryResponseDataOrdersItemContactsItem) *TradeOrderQueryResponseDataOrdersItem {
	s.Contacts = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetCount(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.Count = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetOrderStatus(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.OrderStatus = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetSubOrderAmountInfos(v []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) *TradeOrderQueryResponseDataOrdersItem {
	s.SubOrderAmountInfos = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetOpenId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.OpenId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetDiscounts(v []*TradeOrderQueryResponseDataOrdersItemDiscountsItem) *TradeOrderQueryResponseDataOrdersItem {
	s.Discounts = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetCreateOrderTime(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.CreateOrderTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetSourceOrderId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.SourceOrderId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetAmountInfo(v *TradeOrderQueryResponseDataOrdersItemAmountInfo) *TradeOrderQueryResponseDataOrdersItem {
	s.AmountInfo = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetUpdateOrderTime(v int64) *TradeOrderQueryResponseDataOrdersItem {
	s.UpdateOrderTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetCraftsmanUid(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.CraftsmanUid = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetSkuName(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.SkuName = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetOriginalAmount(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.OriginalAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetThirdSkuId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.ThirdSkuId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetPoiId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.PoiId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetPayAmount(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.PayAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetRoomId(v string) *TradeOrderQueryResponseDataOrdersItem {
	s.RoomId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetOrderType(v int32) *TradeOrderQueryResponseDataOrdersItem {
	s.OrderType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItem) SetPoi(v *TradeOrderQueryResponseDataOrdersItemPoi) *TradeOrderQueryResponseDataOrdersItem {
	s.Poi = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemAmountInfo struct {
	MerchantDiscountAmount    *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	ProductOriginAmount       *int64 `json:"product_origin_amount,omitempty" xml:"product_origin_amount,omitempty"`
	CommissionAmount          *int64 `json:"commission_amount,omitempty" xml:"commission_amount,omitempty"`
	OriginAmount              *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	PlatformDiscountAmount    *int64 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	FreightPayAmount          *int64 `json:"freight_pay_amount,omitempty" xml:"freight_pay_amount,omitempty"`
	ActivitiesFeeAmount       *int64 `json:"activities_fee_amount,omitempty" xml:"activities_fee_amount,omitempty"`
	ProviderDiscountAmount    *int64 `json:"provider_discount_amount,omitempty" xml:"provider_discount_amount,omitempty"`
	PayAmount                 *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	BrandDiscountAmount       *int64 `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
	SalePrice                 *int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	MerchantDeliverFreightFee *int64 `json:"merchant_deliver_freight_fee,omitempty" xml:"merchant_deliver_freight_fee,omitempty"`
	EstimatedOrderIncome      *int64 `json:"estimated_order_income,omitempty" xml:"estimated_order_income,omitempty"`
	PayDiscountAmount         *int64 `json:"pay_discount_amount,omitempty" xml:"pay_discount_amount,omitempty"`
	PlatformDeliverFreightFee *int64 `json:"platform_deliver_freight_fee,omitempty" xml:"platform_deliver_freight_fee,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemAmountInfo) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetMerchantDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetProductOriginAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.ProductOriginAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetCommissionAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.CommissionAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetOriginAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.OriginAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetPlatformDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetFreightPayAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.FreightPayAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetActivitiesFeeAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.ActivitiesFeeAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetProviderDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.ProviderDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetPayAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetBrandDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.BrandDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetSalePrice(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.SalePrice = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetMerchantDeliverFreightFee(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.MerchantDeliverFreightFee = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetEstimatedOrderIncome(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.EstimatedOrderIncome = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetPayDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemAmountInfo) SetPlatformDeliverFreightFee(v int64) *TradeOrderQueryResponseDataOrdersItemAmountInfo {
	s.PlatformDeliverFreightFee = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemCertificateItem struct {
	OrderItemId    *string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	RefundAmount   *int32  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	RefundTime     *int64  `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	CertificateId  *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	CombinationId  *string `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
	ItemStatus     *int32  `json:"item_status,omitempty" xml:"item_status,omitempty"`
	ItemUpdateTime *int64  `json:"item_update_time,omitempty" xml:"item_update_time,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemCertificateItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemCertificateItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetOrderItemId(v string) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.OrderItemId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetRefundAmount(v int32) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetRefundTime(v int64) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetCertificateId(v string) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.CertificateId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetCombinationId(v string) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.CombinationId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetItemStatus(v int32) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemStatus = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemCertificateItem) SetItemUpdateTime(v int64) *TradeOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemUpdateTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemContactsItem struct {
	Phone        *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PhoneEncrypt *string `json:"phone_encrypt,omitempty" xml:"phone_encrypt,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemContactsItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemContactsItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemContactsItem) SetPhone(v string) *TradeOrderQueryResponseDataOrdersItemContactsItem {
	s.Phone = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemContactsItem) SetPhoneEncrypt(v string) *TradeOrderQueryResponseDataOrdersItemContactsItem {
	s.PhoneEncrypt = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemContactsItem) SetName(v string) *TradeOrderQueryResponseDataOrdersItemContactsItem {
	s.Name = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDeliveryInfo struct {
	DeliverModel   *int32  `json:"deliver_model,omitempty" xml:"deliver_model,omitempty"`
	IsBook         *bool   `json:"is_book,omitempty" xml:"is_book,omitempty"`
	Remark         *string `json:"remark,omitempty" xml:"remark,omitempty"`
	ShopNumber     *string `json:"shop_number,omitempty" xml:"shop_number,omitempty"`
	SysExpectTime  *string `json:"sys_expect_time,omitempty" xml:"sys_expect_time,omitempty"`
	TableWare      *string `json:"table_ware,omitempty" xml:"table_ware,omitempty"`
	UserExpectTime *string `json:"user_expect_time,omitempty" xml:"user_expect_time,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDeliveryInfo) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetDeliverModel(v int32) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.DeliverModel = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetIsBook(v bool) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.IsBook = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetRemark(v string) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.Remark = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetShopNumber(v string) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.ShopNumber = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetSysExpectTime(v string) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.SysExpectTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetTableWare(v string) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.TableWare = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDeliveryInfo) SetUserExpectTime(v string) *TradeOrderQueryResponseDataOrdersItemDeliveryInfo {
	s.UserExpectTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDiscountsItem struct {
	DiscountAmount         *int64                                                          `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountType           *int64                                                          `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	IdleTimeInfo           *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo `json:"idle_time_info,omitempty" xml:"idle_time_info,omitempty"`
	MerchantDiscountAmount *int64                                                          `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64                                                          `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ProviderDiscountAmount *int64                                                          `json:"provider_discount_amount,omitempty" xml:"provider_discount_amount,omitempty"`
	BrandDiscountAmount    *int64                                                          `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountType(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetIdleTimeInfo(v *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.IdleTimeInfo = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetMerchantDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetPlatformDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetProviderDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.ProviderDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItem) SetBrandDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItem {
	s.BrandDiscountAmount = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo struct {
	IdleTimeLimitType *int                                                                             `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	WriteOffTimeLimit *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit `json:"WriteOffTimeLimit,omitempty" xml:"WriteOffTimeLimit,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo) SetIdleTimeLimitType(v int) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo) SetWriteOffTimeLimit(v *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfo {
	s.WriteOffTimeLimit = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit struct {
	WeekDayList        []*int                                                                                                   `json:"WeekDayList,omitempty" xml:"WeekDayList,omitempty" type:"Repeated"`
	WriteOffEndTime    *int64                                                                                                   `json:"WriteOffEndTime,omitempty" xml:"WriteOffEndTime,omitempty"`
	WriteOffStartTime  *int64                                                                                                   `json:"WriteOffStartTime,omitempty" xml:"WriteOffStartTime,omitempty"`
	DailyTimeRangeList []*TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem `json:"DailyTimeRangeList,omitempty" xml:"DailyTimeRangeList,omitempty" type:"Repeated"`
	RelativeTime       *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime             `json:"RelativeTime,omitempty" xml:"RelativeTime,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWeekDayList(v []*int) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WeekDayList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWriteOffEndTime(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WriteOffEndTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWriteOffStartTime(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WriteOffStartTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetDailyTimeRangeList(v []*TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.DailyTimeRangeList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetRelativeTime(v *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.RelativeTime = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) SetStartTime(v string) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) SetEndTime(v string) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime struct {
	RelativeTime     *int64 `json:"RelativeTime,omitempty" xml:"RelativeTime,omitempty"`
	RelativeTimeType *int   `json:"RelativeTimeType,omitempty" xml:"RelativeTimeType,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) SetRelativeTime(v int64) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime {
	s.RelativeTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) SetRelativeTimeType(v int) *TradeOrderQueryResponseDataOrdersItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime {
	s.RelativeTimeType = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemMerchantInfo struct {
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemMerchantInfo) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemMerchantInfo) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountId(v string) *TradeOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountName(v string) *TradeOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountName = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemPoi struct {
	PoiId   *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	PoiName *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemPoi) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemPoi) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemPoi) SetPoiId(v string) *TradeOrderQueryResponseDataOrdersItemPoi {
	s.PoiId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemPoi) SetPoiName(v string) *TradeOrderQueryResponseDataOrdersItemPoi {
	s.PoiName = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItem struct {
	SnapshotProduct *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct `json:"snapshot_product,omitempty" xml:"snapshot_product,omitempty"`
	Num             *int32                                                            `json:"num,omitempty" xml:"num,omitempty"`
	ProductName     *string                                                           `json:"product_name,omitempty" xml:"product_name,omitempty"`
	SkuId           *string                                                           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	OriginAmount    *int64                                                            `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ProductId       *string                                                           `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetSnapshotProduct(v *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.SnapshotProduct = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetNum(v int32) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.Num = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetProductName(v string) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductName = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetSkuId(v string) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.SkuId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetOriginAmount(v int64) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.OriginAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItem) SetProductId(v string) *TradeOrderQueryResponseDataOrdersItemProductsItem {
	s.ProductId = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct struct {
	TagInfoList []*TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem `json:"tag_info_list,omitempty" xml:"tag_info_list,omitempty" type:"Repeated"`
	ProductAttr *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr       `json:"product_attr,omitempty" xml:"product_attr,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct) SetTagInfoList(v []*TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct {
	s.TagInfoList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct) SetProductAttr(v *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProduct {
	s.ProductAttr = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr struct {
	DynamicParRule *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule `json:"dynamic_par_rule,omitempty" xml:"dynamic_par_rule,omitempty"`
	ShowChannel    *int                                                                                       `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	VoucherParType *int                                                                                       `json:"voucher_par_type,omitempty" xml:"voucher_par_type,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) SetDynamicParRule(v *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr {
	s.DynamicParRule = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) SetShowChannel(v int) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr {
	s.ShowChannel = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr) SetVoucherParType(v int) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttr {
	s.VoucherParType = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule struct {
	DynamicParMinAmount    *int64                                                                                                                 `json:"dynamic_par_min_amount,omitempty" xml:"dynamic_par_min_amount,omitempty" require:"true"`
	DynamicParRuleDesc     *string                                                                                                                `json:"dynamic_par_rule_desc,omitempty" xml:"dynamic_par_rule_desc,omitempty" require:"true"`
	DynamicParRuleItemList []*TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem `json:"dynamic_par_rule_item_list,omitempty" xml:"dynamic_par_rule_item_list,omitempty" require:"true" type:"Repeated"`
	DynamicParMaxAmount    *int64                                                                                                                 `json:"dynamic_par_max_amount,omitempty" xml:"dynamic_par_max_amount,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) SetDynamicParMinAmount(v int64) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule {
	s.DynamicParMinAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) SetDynamicParRuleDesc(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule {
	s.DynamicParRuleDesc = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) SetDynamicParRuleItemList(v []*TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule {
	s.DynamicParRuleItemList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule) SetDynamicParMaxAmount(v int64) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRule {
	s.DynamicParMaxAmount = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem struct {
	DeductibleAmount *int64                                                                                                                         `json:"deductible_amount,omitempty" xml:"deductible_amount,omitempty" require:"true"`
	TimePeriod       *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod `json:"time_period,omitempty" xml:"time_period,omitempty" require:"true"`
	DayOfWeek        []*int                                                                                                                         `json:"day_of_week,omitempty" xml:"day_of_week,omitempty" require:"true" type:"Repeated"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) SetDeductibleAmount(v int64) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem {
	s.DeductibleAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) SetTimePeriod(v *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem {
	s.TimePeriod = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem) SetDayOfWeek(v []*int) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItem {
	s.DayOfWeek = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod struct {
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
	UseEndTime       *string `json:"use_end_time,omitempty" xml:"use_end_time,omitempty" require:"true"`
	UseStartTime     *string `json:"use_start_time,omitempty" xml:"use_start_time,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetEndTimeIsNextDay(v bool) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.EndTimeIsNextDay = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetUseEndTime(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.UseEndTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetUseStartTime(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductProductAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.UseStartTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) SetTagKey(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem {
	s.TagKey = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) SetTagName(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem {
	s.TagName = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem) SetTagValue(v string) *TradeOrderQueryResponseDataOrdersItemProductsItemSnapshotProductTagInfoListItem {
	s.TagValue = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem struct {
	DiscountAmount *int64                                                                       `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	Discounts      []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	OriginAmount   *int64                                                                       `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	PayAmount      *int64                                                                       `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	ReceiptAmount  *int64                                                                       `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	SubOrderId     *string                                                                      `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	SubOrderType   *int32                                                                       `json:"sub_order_type,omitempty" xml:"sub_order_type,omitempty"`
	CombinationId  *string                                                                      `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.DiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscounts(v []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.Discounts = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetOriginAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.OriginAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetPayAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.PayAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetReceiptAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.ReceiptAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderId(v string) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderId = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderType(v int32) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetCombinationId(v string) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.CombinationId = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem struct {
	MerchantDiscountAmount *int64                                                                                 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64                                                                                 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	ProviderDiscountAmount *int64                                                                                 `json:"provider_discount_amount,omitempty" xml:"provider_discount_amount,omitempty"`
	BrandDiscountAmount    *int64                                                                                 `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
	DiscountAmount         *int64                                                                                 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountType           *int64                                                                                 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	IdleTimeInfo           *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo `json:"idle_time_info,omitempty" xml:"idle_time_info,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetMerchantDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetPlatformDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetProviderDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.ProviderDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetBrandDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.BrandDiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountAmount(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountType(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetIdleTimeInfo(v *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.IdleTimeInfo = v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo struct {
	WriteOffTimeLimit *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit `json:"WriteOffTimeLimit,omitempty" xml:"WriteOffTimeLimit,omitempty"`
	IdleTimeLimitType *int                                                                                                    `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo) SetWriteOffTimeLimit(v *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo {
	s.WriteOffTimeLimit = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo) SetIdleTimeLimitType(v int) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfo {
	s.IdleTimeLimitType = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit struct {
	WriteOffStartTime  *int64                                                                                                                          `json:"WriteOffStartTime,omitempty" xml:"WriteOffStartTime,omitempty"`
	DailyTimeRangeList []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem `json:"DailyTimeRangeList,omitempty" xml:"DailyTimeRangeList,omitempty" type:"Repeated"`
	RelativeTime       *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime             `json:"RelativeTime,omitempty" xml:"RelativeTime,omitempty"`
	WeekDayList        []*int                                                                                                                          `json:"WeekDayList,omitempty" xml:"WeekDayList,omitempty" type:"Repeated"`
	WriteOffEndTime    *int64                                                                                                                          `json:"WriteOffEndTime,omitempty" xml:"WriteOffEndTime,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWriteOffStartTime(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WriteOffStartTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetDailyTimeRangeList(v []*TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.DailyTimeRangeList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetRelativeTime(v *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.RelativeTime = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWeekDayList(v []*int) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WeekDayList = v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit) SetWriteOffEndTime(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimit {
	s.WriteOffEndTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) SetEndTime(v string) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem {
	s.EndTime = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem) SetStartTime(v string) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitDailyTimeRangeListItem {
	s.StartTime = &v
	return s
}

type TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime struct {
	RelativeTimeType *int   `json:"RelativeTimeType,omitempty" xml:"RelativeTimeType,omitempty" require:"true"`
	RelativeTime     *int64 `json:"RelativeTime,omitempty" xml:"RelativeTime,omitempty"`
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) SetRelativeTimeType(v int) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime {
	s.RelativeTimeType = &v
	return s
}

func (s *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime) SetRelativeTime(v int64) *TradeOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemIdleTimeInfoWriteOffTimeLimitRelativeTime {
	s.RelativeTime = &v
	return s
}

type TradeOrderQueryResponseDataPage struct {
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Total    *int64 `json:"total,omitempty" xml:"total,omitempty"`
	PageNum  *int32 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

func (s TradeOrderQueryResponseDataPage) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataPage) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataPage) SetPageSize(v int32) *TradeOrderQueryResponseDataPage {
	s.PageSize = &v
	return s
}

func (s *TradeOrderQueryResponseDataPage) SetTotal(v int64) *TradeOrderQueryResponseDataPage {
	s.Total = &v
	return s
}

func (s *TradeOrderQueryResponseDataPage) SetPageNum(v int32) *TradeOrderQueryResponseDataPage {
	s.PageNum = &v
	return s
}

type TradeOrderQueryResponseDataSearchAfter struct {
	CursorValue    []*string   `json:"CursorValue,omitempty" xml:"CursorValue,omitempty" type:"Repeated"`
	Size           *int32      `json:"Size,omitempty" xml:"Size,omitempty"`
	AllCursorValue [][]*string `json:"AllCursorValue,omitempty" xml:"AllCursorValue,omitempty" type:"Repeated"`
}

func (s TradeOrderQueryResponseDataSearchAfter) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseDataSearchAfter) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseDataSearchAfter) SetCursorValue(v []*string) *TradeOrderQueryResponseDataSearchAfter {
	s.CursorValue = v
	return s
}

func (s *TradeOrderQueryResponseDataSearchAfter) SetSize(v int32) *TradeOrderQueryResponseDataSearchAfter {
	s.Size = &v
	return s
}

func (s *TradeOrderQueryResponseDataSearchAfter) SetAllCursorValue(v [][]*string) *TradeOrderQueryResponseDataSearchAfter {
	s.AllCursorValue = v
	return s
}

type TradeOrderQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TradeOrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TradeOrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *TradeOrderQueryResponseExtra) SetNow(v int64) *TradeOrderQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *TradeOrderQueryResponseExtra) SetSubDescription(v string) *TradeOrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TradeOrderQueryResponseExtra) SetSubErrorCode(v int32) *TradeOrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TradeOrderQueryResponseExtra) SetDescription(v string) *TradeOrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *TradeOrderQueryResponseExtra) SetErrorCode(v int32) *TradeOrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TradeOrderQueryResponseExtra) SetLogid(v string) *TradeOrderQueryResponseExtra {
	s.Logid = &v
	return s
}

type TrafficPermissionOpenRequest struct {
	TaxNature         *int               `json:"tax_nature,omitempty" xml:"tax_nature,omitempty" require:"true"`
	BankAccountNumber *string            `json:"bank_account_number,omitempty" xml:"bank_account_number,omitempty" require:"true"`
	BankBranch        *string            `json:"bank_branch,omitempty" xml:"bank_branch,omitempty" require:"true"`
	Province          *string            `json:"province,omitempty" xml:"province,omitempty" require:"true"`
	BankName          *string            `json:"bank_name,omitempty" xml:"bank_name,omitempty" require:"true"`
	PhoneNumber       *string            `json:"phone_number,omitempty" xml:"phone_number,omitempty" require:"true"`
	Header            map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	City              *string            `json:"city,omitempty" xml:"city,omitempty" require:"true"`
}

func (s TrafficPermissionOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s TrafficPermissionOpenRequest) GoString() string {
	return s.String()
}

func (s *TrafficPermissionOpenRequest) SetTaxNature(v int) *TrafficPermissionOpenRequest {
	s.TaxNature = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetBankAccountNumber(v string) *TrafficPermissionOpenRequest {
	s.BankAccountNumber = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetBankBranch(v string) *TrafficPermissionOpenRequest {
	s.BankBranch = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetProvince(v string) *TrafficPermissionOpenRequest {
	s.Province = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetBankName(v string) *TrafficPermissionOpenRequest {
	s.BankName = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetPhoneNumber(v string) *TrafficPermissionOpenRequest {
	s.PhoneNumber = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetHeader(v map[string]*string) *TrafficPermissionOpenRequest {
	s.Header = v
	return s
}

func (s *TrafficPermissionOpenRequest) SetAccessToken(v string) *TrafficPermissionOpenRequest {
	s.AccessToken = &v
	return s
}

func (s *TrafficPermissionOpenRequest) SetCity(v string) *TrafficPermissionOpenRequest {
	s.City = &v
	return s
}

type TrafficPermissionOpenResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s TrafficPermissionOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s TrafficPermissionOpenResponse) GoString() string {
	return s.String()
}

func (s *TrafficPermissionOpenResponse) SetErrMsg(v string) *TrafficPermissionOpenResponse {
	s.ErrMsg = &v
	return s
}

func (s *TrafficPermissionOpenResponse) SetLogId(v string) *TrafficPermissionOpenResponse {
	s.LogId = &v
	return s
}

func (s *TrafficPermissionOpenResponse) SetErrNo(v int32) *TrafficPermissionOpenResponse {
	s.ErrNo = &v
	return s
}

type TrafficPermissionQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TrafficPermissionQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TrafficPermissionQueryRequest) GoString() string {
	return s.String()
}

func (s *TrafficPermissionQueryRequest) SetHeader(v map[string]*string) *TrafficPermissionQueryRequest {
	s.Header = v
	return s
}

func (s *TrafficPermissionQueryRequest) SetAccessToken(v string) *TrafficPermissionQueryRequest {
	s.AccessToken = &v
	return s
}

type TrafficPermissionQueryResponse struct {
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *TrafficPermissionQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TrafficPermissionQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TrafficPermissionQueryResponse) GoString() string {
	return s.String()
}

func (s *TrafficPermissionQueryResponse) SetErrNo(v int32) *TrafficPermissionQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *TrafficPermissionQueryResponse) SetErrMsg(v string) *TrafficPermissionQueryResponse {
	s.ErrMsg = &v
	return s
}

func (s *TrafficPermissionQueryResponse) SetLogId(v string) *TrafficPermissionQueryResponse {
	s.LogId = &v
	return s
}

func (s *TrafficPermissionQueryResponse) SetData(v *TrafficPermissionQueryResponseData) *TrafficPermissionQueryResponse {
	s.Data = v
	return s
}

type TrafficPermissionQueryResponseData struct {
	CanOpen *int64 `json:"can_open,omitempty" xml:"can_open,omitempty"`
	Status  *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TrafficPermissionQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s TrafficPermissionQueryResponseData) GoString() string {
	return s.String()
}

func (s *TrafficPermissionQueryResponseData) SetCanOpen(v int64) *TrafficPermissionQueryResponseData {
	s.CanOpen = &v
	return s
}

func (s *TrafficPermissionQueryResponseData) SetStatus(v int64) *TrafficPermissionQueryResponseData {
	s.Status = &v
	return s
}

type TrafficVerifyRequest struct {
	Header      map[string]*string                  `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                             `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Vouchers    []*TrafficVerifyRequestVouchersItem `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	OrderId     *string                             `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	VerifyMode  *int                                `json:"verify_mode,omitempty" xml:"verify_mode,omitempty"`
	VerifyToken *string                             `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
}

func (s TrafficVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyRequest) GoString() string {
	return s.String()
}

func (s *TrafficVerifyRequest) SetHeader(v map[string]*string) *TrafficVerifyRequest {
	s.Header = v
	return s
}

func (s *TrafficVerifyRequest) SetAccessToken(v string) *TrafficVerifyRequest {
	s.AccessToken = &v
	return s
}

func (s *TrafficVerifyRequest) SetVouchers(v []*TrafficVerifyRequestVouchersItem) *TrafficVerifyRequest {
	s.Vouchers = v
	return s
}

func (s *TrafficVerifyRequest) SetOrderId(v string) *TrafficVerifyRequest {
	s.OrderId = &v
	return s
}

func (s *TrafficVerifyRequest) SetVerifyMode(v int) *TrafficVerifyRequest {
	s.VerifyMode = &v
	return s
}

func (s *TrafficVerifyRequest) SetVerifyToken(v string) *TrafficVerifyRequest {
	s.VerifyToken = &v
	return s
}

type TrafficVerifyRequestVouchersItem struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s TrafficVerifyRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *TrafficVerifyRequestVouchersItem) SetCode(v string) *TrafficVerifyRequestVouchersItem {
	s.Code = &v
	return s
}

type TrafficVerifyResponse struct {
	Data  *TrafficVerifyResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TrafficVerifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TrafficVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyResponse) GoString() string {
	return s.String()
}

func (s *TrafficVerifyResponse) SetData(v *TrafficVerifyResponseData) *TrafficVerifyResponse {
	s.Data = v
	return s
}

func (s *TrafficVerifyResponse) SetExtra(v *TrafficVerifyResponseExtra) *TrafficVerifyResponse {
	s.Extra = v
	return s
}

type TrafficVerifyResponseData struct {
	GwErrorCode   *int32                                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	VerifyResults []*TrafficVerifyResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
}

func (s TrafficVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyResponseData) GoString() string {
	return s.String()
}

func (s *TrafficVerifyResponseData) SetGwErrorCode(v int32) *TrafficVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TrafficVerifyResponseData) SetGwDescription(v string) *TrafficVerifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *TrafficVerifyResponseData) SetVerifyResults(v []*TrafficVerifyResponseDataVerifyResultsItem) *TrafficVerifyResponseData {
	s.VerifyResults = v
	return s
}

type TrafficVerifyResponseDataVerifyResultsItem struct {
	Result        *int32  `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	VerifyId      *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	Msg           *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
}

func (s TrafficVerifyResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *TrafficVerifyResponseDataVerifyResultsItem) SetResult(v int32) *TrafficVerifyResponseDataVerifyResultsItem {
	s.Result = &v
	return s
}

func (s *TrafficVerifyResponseDataVerifyResultsItem) SetVerifyId(v string) *TrafficVerifyResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *TrafficVerifyResponseDataVerifyResultsItem) SetCertificateId(v string) *TrafficVerifyResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

func (s *TrafficVerifyResponseDataVerifyResultsItem) SetMsg(v string) *TrafficVerifyResponseDataVerifyResultsItem {
	s.Msg = &v
	return s
}

type TrafficVerifyResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TrafficVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TrafficVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *TrafficVerifyResponseExtra) SetNow(v int64) *TrafficVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *TrafficVerifyResponseExtra) SetSubDescription(v string) *TrafficVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TrafficVerifyResponseExtra) SetSubErrorCode(v int32) *TrafficVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TrafficVerifyResponseExtra) SetDescription(v string) *TrafficVerifyResponseExtra {
	s.Description = &v
	return s
}

func (s *TrafficVerifyResponseExtra) SetErrorCode(v int32) *TrafficVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TrafficVerifyResponseExtra) SetLogid(v string) *TrafficVerifyResponseExtra {
	s.Logid = &v
	return s
}

type TransferCallbackRequest struct {
	FailReason   *string                                    `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	OrderId      *string                                    `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RequestId    *string                                    `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	Result       *int64                                     `json:"result,omitempty" xml:"result,omitempty" require:"true"`
	Header       map[string]*string                         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string                                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Certificates []*TransferCallbackRequestCertificatesItem `json:"certificates,omitempty" xml:"certificates,omitempty" type:"Repeated"`
}

func (s TransferCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallbackRequest) GoString() string {
	return s.String()
}

func (s *TransferCallbackRequest) SetFailReason(v string) *TransferCallbackRequest {
	s.FailReason = &v
	return s
}

func (s *TransferCallbackRequest) SetOrderId(v string) *TransferCallbackRequest {
	s.OrderId = &v
	return s
}

func (s *TransferCallbackRequest) SetRequestId(v string) *TransferCallbackRequest {
	s.RequestId = &v
	return s
}

func (s *TransferCallbackRequest) SetResult(v int64) *TransferCallbackRequest {
	s.Result = &v
	return s
}

func (s *TransferCallbackRequest) SetHeader(v map[string]*string) *TransferCallbackRequest {
	s.Header = v
	return s
}

func (s *TransferCallbackRequest) SetAccessToken(v string) *TransferCallbackRequest {
	s.AccessToken = &v
	return s
}

func (s *TransferCallbackRequest) SetCertificates(v []*TransferCallbackRequestCertificatesItem) *TransferCallbackRequest {
	s.Certificates = v
	return s
}

type TransferCallbackRequestCertificatesItem struct {
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	NewCode       *string `json:"new_code,omitempty" xml:"new_code,omitempty"`
}

func (s TransferCallbackRequestCertificatesItem) String() string {
	return tea.Prettify(s)
}

func (s TransferCallbackRequestCertificatesItem) GoString() string {
	return s.String()
}

func (s *TransferCallbackRequestCertificatesItem) SetCertificateId(v string) *TransferCallbackRequestCertificatesItem {
	s.CertificateId = &v
	return s
}

func (s *TransferCallbackRequestCertificatesItem) SetCode(v string) *TransferCallbackRequestCertificatesItem {
	s.Code = &v
	return s
}

func (s *TransferCallbackRequestCertificatesItem) SetNewCode(v string) *TransferCallbackRequestCertificatesItem {
	s.NewCode = &v
	return s
}

type TransferCallbackResponse struct {
	Extra *TransferCallbackResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TransferCallbackResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TransferCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallbackResponse) GoString() string {
	return s.String()
}

func (s *TransferCallbackResponse) SetExtra(v *TransferCallbackResponseExtra) *TransferCallbackResponse {
	s.Extra = v
	return s
}

func (s *TransferCallbackResponse) SetData(v *TransferCallbackResponseData) *TransferCallbackResponse {
	s.Data = v
	return s
}

type TransferCallbackResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TransferCallbackResponseData) String() string {
	return tea.Prettify(s)
}

func (s TransferCallbackResponseData) GoString() string {
	return s.String()
}

func (s *TransferCallbackResponseData) SetGwErrorCode(v int32) *TransferCallbackResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TransferCallbackResponseData) SetGwDescription(v string) *TransferCallbackResponseData {
	s.GwDescription = &v
	return s
}

type TransferCallbackResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TransferCallbackResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TransferCallbackResponseExtra) GoString() string {
	return s.String()
}

func (s *TransferCallbackResponseExtra) SetErrorCode(v int32) *TransferCallbackResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TransferCallbackResponseExtra) SetLogid(v string) *TransferCallbackResponseExtra {
	s.Logid = &v
	return s
}

func (s *TransferCallbackResponseExtra) SetNow(v int64) *TransferCallbackResponseExtra {
	s.Now = &v
	return s
}

func (s *TransferCallbackResponseExtra) SetSubDescription(v string) *TransferCallbackResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TransferCallbackResponseExtra) SetSubErrorCode(v int32) *TransferCallbackResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TransferCallbackResponseExtra) SetDescription(v string) *TransferCallbackResponseExtra {
	s.Description = &v
	return s
}

type TravelNewRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TravelNewRequest) String() string {
	return tea.Prettify(s)
}

func (s TravelNewRequest) GoString() string {
	return s.String()
}

func (s *TravelNewRequest) SetAccessToken(v string) *TravelNewRequest {
	s.AccessToken = &v
	return s
}

func (s *TravelNewRequest) SetHeader(v map[string]*string) *TravelNewRequest {
	s.Header = v
	return s
}

type TravelNewResponse struct {
	Extra *TravelNewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TravelNewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TravelNewResponse) String() string {
	return tea.Prettify(s)
}

func (s TravelNewResponse) GoString() string {
	return s.String()
}

func (s *TravelNewResponse) SetExtra(v *TravelNewResponseExtra) *TravelNewResponse {
	s.Extra = v
	return s
}

func (s *TravelNewResponse) SetData(v *TravelNewResponseData) *TravelNewResponse {
	s.Data = v
	return s
}

type TravelNewResponseData struct {
	GwDescription *string                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*TravelNewResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TravelNewResponseData) String() string {
	return tea.Prettify(s)
}

func (s TravelNewResponseData) GoString() string {
	return s.String()
}

func (s *TravelNewResponseData) SetGwDescription(v string) *TravelNewResponseData {
	s.GwDescription = &v
	return s
}

func (s *TravelNewResponseData) SetList(v []*TravelNewResponseDataListItem) *TravelNewResponseData {
	s.List = v
	return s
}

func (s *TravelNewResponseData) SetGwErrorCode(v int32) *TravelNewResponseData {
	s.GwErrorCode = &v
	return s
}

type TravelNewResponseDataListItem struct {
	RankChange       *string                                       `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                       `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                       `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                        `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                        `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                      `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*TravelNewResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                        `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
}

func (s TravelNewResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelNewResponseDataListItem) GoString() string {
	return s.String()
}

func (s *TravelNewResponseDataListItem) SetRankChange(v string) *TravelNewResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetNickname(v string) *TravelNewResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetAvatar(v string) *TravelNewResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetFollowerCount(v int64) *TravelNewResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetOnbillbaordTimes(v int32) *TravelNewResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetEffectValue(v float64) *TravelNewResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *TravelNewResponseDataListItem) SetVideoList(v []*TravelNewResponseDataListItemVideoListItem) *TravelNewResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *TravelNewResponseDataListItem) SetRank(v int32) *TravelNewResponseDataListItem {
	s.Rank = &v
	return s
}

type TravelNewResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s TravelNewResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelNewResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *TravelNewResponseDataListItemVideoListItem) SetTitle(v string) *TravelNewResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *TravelNewResponseDataListItemVideoListItem) SetItemCover(v string) *TravelNewResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *TravelNewResponseDataListItemVideoListItem) SetShareUrl(v string) *TravelNewResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type TravelNewResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TravelNewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TravelNewResponseExtra) GoString() string {
	return s.String()
}

func (s *TravelNewResponseExtra) SetSubErrorCode(v int32) *TravelNewResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TravelNewResponseExtra) SetSubDescription(v string) *TravelNewResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TravelNewResponseExtra) SetLogid(v string) *TravelNewResponseExtra {
	s.Logid = &v
	return s
}

func (s *TravelNewResponseExtra) SetNow(v int64) *TravelNewResponseExtra {
	s.Now = &v
	return s
}

func (s *TravelNewResponseExtra) SetErrorCode(v int32) *TravelNewResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TravelNewResponseExtra) SetDescription(v string) *TravelNewResponseExtra {
	s.Description = &v
	return s
}

type TravelOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TravelOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallRequest) GoString() string {
	return s.String()
}

func (s *TravelOverallRequest) SetHeader(v map[string]*string) *TravelOverallRequest {
	s.Header = v
	return s
}

func (s *TravelOverallRequest) SetAccessToken(v string) *TravelOverallRequest {
	s.AccessToken = &v
	return s
}

type TravelOverallResponse struct {
	Extra *TravelOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TravelOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TravelOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallResponse) GoString() string {
	return s.String()
}

func (s *TravelOverallResponse) SetExtra(v *TravelOverallResponseExtra) *TravelOverallResponse {
	s.Extra = v
	return s
}

func (s *TravelOverallResponse) SetData(v *TravelOverallResponseData) *TravelOverallResponse {
	s.Data = v
	return s
}

type TravelOverallResponseData struct {
	List          []*TravelOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                               `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                              `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TravelOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallResponseData) GoString() string {
	return s.String()
}

func (s *TravelOverallResponseData) SetList(v []*TravelOverallResponseDataListItem) *TravelOverallResponseData {
	s.List = v
	return s
}

func (s *TravelOverallResponseData) SetGwErrorCode(v int32) *TravelOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TravelOverallResponseData) SetGwDescription(v string) *TravelOverallResponseData {
	s.GwDescription = &v
	return s
}

type TravelOverallResponseDataListItem struct {
	Nickname         *string                                           `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                           `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                            `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                            `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                          `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*TravelOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                            `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                           `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
}

func (s TravelOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *TravelOverallResponseDataListItem) SetNickname(v string) *TravelOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetAvatar(v string) *TravelOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetFollowerCount(v int64) *TravelOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *TravelOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetEffectValue(v float64) *TravelOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetVideoList(v []*TravelOverallResponseDataListItemVideoListItem) *TravelOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *TravelOverallResponseDataListItem) SetRank(v int32) *TravelOverallResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *TravelOverallResponseDataListItem) SetRankChange(v string) *TravelOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

type TravelOverallResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s TravelOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *TravelOverallResponseDataListItemVideoListItem) SetTitle(v string) *TravelOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *TravelOverallResponseDataListItemVideoListItem) SetItemCover(v string) *TravelOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *TravelOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *TravelOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type TravelOverallResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TravelOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TravelOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *TravelOverallResponseExtra) SetNow(v int64) *TravelOverallResponseExtra {
	s.Now = &v
	return s
}

func (s *TravelOverallResponseExtra) SetErrorCode(v int32) *TravelOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TravelOverallResponseExtra) SetDescription(v string) *TravelOverallResponseExtra {
	s.Description = &v
	return s
}

func (s *TravelOverallResponseExtra) SetSubErrorCode(v int32) *TravelOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TravelOverallResponseExtra) SetSubDescription(v string) *TravelOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TravelOverallResponseExtra) SetLogid(v string) *TravelOverallResponseExtra {
	s.Logid = &v
	return s
}

type TravelagencyProductOperateRequest struct {
	AccountId       *string                                                 `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ProductInfoList []*TravelagencyProductOperateRequestProductInfoListItem `json:"product_info_list,omitempty" xml:"product_info_list,omitempty" type:"Repeated"`
	OptType         *int32                                                  `json:"opt_type,omitempty" xml:"opt_type,omitempty"`
	Header          map[string]*string                                      `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                                 `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s TravelagencyProductOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateRequest) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateRequest) SetAccountId(v string) *TravelagencyProductOperateRequest {
	s.AccountId = &v
	return s
}

func (s *TravelagencyProductOperateRequest) SetProductInfoList(v []*TravelagencyProductOperateRequestProductInfoListItem) *TravelagencyProductOperateRequest {
	s.ProductInfoList = v
	return s
}

func (s *TravelagencyProductOperateRequest) SetOptType(v int32) *TravelagencyProductOperateRequest {
	s.OptType = &v
	return s
}

func (s *TravelagencyProductOperateRequest) SetHeader(v map[string]*string) *TravelagencyProductOperateRequest {
	s.Header = v
	return s
}

func (s *TravelagencyProductOperateRequest) SetAccessToken(v string) *TravelagencyProductOperateRequest {
	s.AccessToken = &v
	return s
}

type TravelagencyProductOperateRequestProductInfoListItem struct {
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
}

func (s TravelagencyProductOperateRequestProductInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateRequestProductInfoListItem) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateRequestProductInfoListItem) SetProductId(v string) *TravelagencyProductOperateRequestProductInfoListItem {
	s.ProductId = &v
	return s
}

func (s *TravelagencyProductOperateRequestProductInfoListItem) SetOutId(v string) *TravelagencyProductOperateRequestProductInfoListItem {
	s.OutId = &v
	return s
}

type TravelagencyProductOperateResponse struct {
	Extra *TravelagencyProductOperateResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TravelagencyProductOperateResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TravelagencyProductOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateResponse) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateResponse) SetExtra(v *TravelagencyProductOperateResponseExtra) *TravelagencyProductOperateResponse {
	s.Extra = v
	return s
}

func (s *TravelagencyProductOperateResponse) SetData(v *TravelagencyProductOperateResponseData) *TravelagencyProductOperateResponse {
	s.Data = v
	return s
}

type TravelagencyProductOperateResponseData struct {
	GwErrorCode         *int32                                                           `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription       *string                                                          `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ProductInfoFailList []*TravelagencyProductOperateResponseDataProductInfoFailListItem `json:"product_info_fail_list,omitempty" xml:"product_info_fail_list,omitempty" type:"Repeated"`
}

func (s TravelagencyProductOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateResponseData) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateResponseData) SetGwErrorCode(v int32) *TravelagencyProductOperateResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TravelagencyProductOperateResponseData) SetGwDescription(v string) *TravelagencyProductOperateResponseData {
	s.GwDescription = &v
	return s
}

func (s *TravelagencyProductOperateResponseData) SetProductInfoFailList(v []*TravelagencyProductOperateResponseDataProductInfoFailListItem) *TravelagencyProductOperateResponseData {
	s.ProductInfoFailList = v
	return s
}

type TravelagencyProductOperateResponseDataProductInfoFailListItem struct {
	FailReason  *string                                                                   `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	ProductInfo *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo `json:"product_info,omitempty" xml:"product_info,omitempty"`
}

func (s TravelagencyProductOperateResponseDataProductInfoFailListItem) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateResponseDataProductInfoFailListItem) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateResponseDataProductInfoFailListItem) SetFailReason(v string) *TravelagencyProductOperateResponseDataProductInfoFailListItem {
	s.FailReason = &v
	return s
}

func (s *TravelagencyProductOperateResponseDataProductInfoFailListItem) SetProductInfo(v *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo) *TravelagencyProductOperateResponseDataProductInfoFailListItem {
	s.ProductInfo = v
	return s
}

type TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo struct {
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	OutId     *string `json:"out_id,omitempty" xml:"out_id,omitempty"`
}

func (s TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo) SetProductId(v string) *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo {
	s.ProductId = &v
	return s
}

func (s *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo) SetOutId(v string) *TravelagencyProductOperateResponseDataProductInfoFailListItemProductInfo {
	s.OutId = &v
	return s
}

type TravelagencyProductOperateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s TravelagencyProductOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TravelagencyProductOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *TravelagencyProductOperateResponseExtra) SetDescription(v string) *TravelagencyProductOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *TravelagencyProductOperateResponseExtra) SetErrorCode(v int32) *TravelagencyProductOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TravelagencyProductOperateResponseExtra) SetLogid(v string) *TravelagencyProductOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *TravelagencyProductOperateResponseExtra) SetNow(v int64) *TravelagencyProductOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *TravelagencyProductOperateResponseExtra) SetSubDescription(v string) *TravelagencyProductOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TravelagencyProductOperateResponseExtra) SetSubErrorCode(v int32) *TravelagencyProductOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type TripCertificateVerifyRequest struct {
	VerifyToken *string                                     `json:"verify_token,omitempty" xml:"verify_token,omitempty" require:"true"`
	Vouchers    []*TripCertificateVerifyRequestVouchersItem `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	OrderId     *string                                     `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	PoiId       *string                                     `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	Header      map[string]*string                          `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                     `json:"access_token,omitempty" xml:"access_token,omitempty"`
	VerifyTime  *int64                                      `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
}

func (s TripCertificateVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequest) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequest) SetVerifyToken(v string) *TripCertificateVerifyRequest {
	s.VerifyToken = &v
	return s
}

func (s *TripCertificateVerifyRequest) SetVouchers(v []*TripCertificateVerifyRequestVouchersItem) *TripCertificateVerifyRequest {
	s.Vouchers = v
	return s
}

func (s *TripCertificateVerifyRequest) SetOrderId(v string) *TripCertificateVerifyRequest {
	s.OrderId = &v
	return s
}

func (s *TripCertificateVerifyRequest) SetPoiId(v string) *TripCertificateVerifyRequest {
	s.PoiId = &v
	return s
}

func (s *TripCertificateVerifyRequest) SetHeader(v map[string]*string) *TripCertificateVerifyRequest {
	s.Header = v
	return s
}

func (s *TripCertificateVerifyRequest) SetAccessToken(v string) *TripCertificateVerifyRequest {
	s.AccessToken = &v
	return s
}

func (s *TripCertificateVerifyRequest) SetVerifyTime(v int64) *TripCertificateVerifyRequest {
	s.VerifyTime = &v
	return s
}

type TripCertificateVerifyRequestVouchersItem struct {
	Projects   []*TripCertificateVerifyRequestVouchersItemProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	VerifyTime *int64                                                  `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	Entrance   *TripCertificateVerifyRequestVouchersItemEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
}

func (s TripCertificateVerifyRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequestVouchersItem) SetProjects(v []*TripCertificateVerifyRequestVouchersItemProjectsItem) *TripCertificateVerifyRequestVouchersItem {
	s.Projects = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItem) SetVerifyTime(v int64) *TripCertificateVerifyRequestVouchersItem {
	s.VerifyTime = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItem) SetEntrance(v *TripCertificateVerifyRequestVouchersItemEntrance) *TripCertificateVerifyRequestVouchersItem {
	s.Entrance = v
	return s
}

type TripCertificateVerifyRequestVouchersItemEntrance struct {
	Urls           []*string                                                          `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                          `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                          `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
	IdCards        []*string                                                          `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                            `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId      *string                                                            `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                          `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
}

func (s TripCertificateVerifyRequestVouchersItemEntrance) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequestVouchersItemEntrance) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetUrls(v []*string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.Urls = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetCertificateNos(v []*string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.CertificateNos = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetCredentials(v []*TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.Credentials = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetGmcodeImgs(v []*string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.GmcodeImgs = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetIdCards(v []*string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.IdCards = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetName(v string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.Name = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetProjectId(v string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.ProjectId = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntrance) SetQrcodes(v []*string) *TripCertificateVerifyRequestVouchersItemEntrance {
	s.Qrcodes = v
	return s
}

type TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem struct {
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem) SetCredentialType(v int) *TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem) SetCredentialNo(v string) *TripCertificateVerifyRequestVouchersItemEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

type TripCertificateVerifyRequestVouchersItemProjectsItem struct {
	IdCards        []*string                                                              `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                                `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId      *string                                                                `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                              `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                              `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                              `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	GmcodeImgs     []*string                                                              `json:"gmcode_imgs,omitempty" xml:"gmcode_imgs,omitempty" type:"Repeated"`
}

func (s TripCertificateVerifyRequestVouchersItemProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequestVouchersItemProjectsItem) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetIdCards(v []*string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.IdCards = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetName(v string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.Name = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetProjectId(v string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.ProjectId = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetQrcodes(v []*string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetUrls(v []*string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.Urls = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetCertificateNos(v []*string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetCredentials(v []*TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.Credentials = v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItem) SetGmcodeImgs(v []*string) *TripCertificateVerifyRequestVouchersItemProjectsItem {
	s.GmcodeImgs = v
	return s
}

type TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem struct {
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
}

func (s TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem) SetCredentialNo(v string) *TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

func (s *TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem) SetCredentialType(v int) *TripCertificateVerifyRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

type TripCertificateVerifyResponse struct {
	Extra *TripCertificateVerifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TripCertificateVerifyResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s TripCertificateVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyResponse) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyResponse) SetExtra(v *TripCertificateVerifyResponseExtra) *TripCertificateVerifyResponse {
	s.Extra = v
	return s
}

func (s *TripCertificateVerifyResponse) SetData(v *TripCertificateVerifyResponseData) *TripCertificateVerifyResponse {
	s.Data = v
	return s
}

type TripCertificateVerifyResponseData struct {
	GwErrorCode   *int32                                                `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                               `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	VerifyResults []*TripCertificateVerifyResponseDataVerifyResultsItem `json:"verify_results,omitempty" xml:"verify_results,omitempty" type:"Repeated"`
}

func (s TripCertificateVerifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyResponseData) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyResponseData) SetGwErrorCode(v int32) *TripCertificateVerifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TripCertificateVerifyResponseData) SetGwDescription(v string) *TripCertificateVerifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *TripCertificateVerifyResponseData) SetVerifyResults(v []*TripCertificateVerifyResponseDataVerifyResultsItem) *TripCertificateVerifyResponseData {
	s.VerifyResults = v
	return s
}

type TripCertificateVerifyResponseDataVerifyResultsItem struct {
	Code          *int32  `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg           *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	VerifyId      *string `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	CertificateId *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
}

func (s TripCertificateVerifyResponseDataVerifyResultsItem) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyResponseDataVerifyResultsItem) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyResponseDataVerifyResultsItem) SetCode(v int32) *TripCertificateVerifyResponseDataVerifyResultsItem {
	s.Code = &v
	return s
}

func (s *TripCertificateVerifyResponseDataVerifyResultsItem) SetMsg(v string) *TripCertificateVerifyResponseDataVerifyResultsItem {
	s.Msg = &v
	return s
}

func (s *TripCertificateVerifyResponseDataVerifyResultsItem) SetVerifyId(v string) *TripCertificateVerifyResponseDataVerifyResultsItem {
	s.VerifyId = &v
	return s
}

func (s *TripCertificateVerifyResponseDataVerifyResultsItem) SetCertificateId(v string) *TripCertificateVerifyResponseDataVerifyResultsItem {
	s.CertificateId = &v
	return s
}

type TripCertificateVerifyResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TripCertificateVerifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TripCertificateVerifyResponseExtra) GoString() string {
	return s.String()
}

func (s *TripCertificateVerifyResponseExtra) SetNow(v int64) *TripCertificateVerifyResponseExtra {
	s.Now = &v
	return s
}

func (s *TripCertificateVerifyResponseExtra) SetSubDescription(v string) *TripCertificateVerifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TripCertificateVerifyResponseExtra) SetSubErrorCode(v int32) *TripCertificateVerifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TripCertificateVerifyResponseExtra) SetDescription(v string) *TripCertificateVerifyResponseExtra {
	s.Description = &v
	return s
}

func (s *TripCertificateVerifyResponseExtra) SetErrorCode(v int32) *TripCertificateVerifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TripCertificateVerifyResponseExtra) SetLogid(v string) *TripCertificateVerifyResponseExtra {
	s.Logid = &v
	return s
}

type TripOrderConfirmRequest struct {
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ConfirmResult *int               `json:"confirm_result,omitempty" xml:"confirm_result,omitempty" require:"true"`
}

func (s TripOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s TripOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *TripOrderConfirmRequest) SetOrderId(v string) *TripOrderConfirmRequest {
	s.OrderId = &v
	return s
}

func (s *TripOrderConfirmRequest) SetHeader(v map[string]*string) *TripOrderConfirmRequest {
	s.Header = v
	return s
}

func (s *TripOrderConfirmRequest) SetAccessToken(v string) *TripOrderConfirmRequest {
	s.AccessToken = &v
	return s
}

func (s *TripOrderConfirmRequest) SetConfirmResult(v int) *TripOrderConfirmRequest {
	s.ConfirmResult = &v
	return s
}

type TripOrderConfirmResponse struct {
	Extra *TripOrderConfirmResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *TripOrderConfirmResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TripOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s TripOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *TripOrderConfirmResponse) SetExtra(v *TripOrderConfirmResponseExtra) *TripOrderConfirmResponse {
	s.Extra = v
	return s
}

func (s *TripOrderConfirmResponse) SetData(v *TripOrderConfirmResponseData) *TripOrderConfirmResponse {
	s.Data = v
	return s
}

type TripOrderConfirmResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TripOrderConfirmResponseData) String() string {
	return tea.Prettify(s)
}

func (s TripOrderConfirmResponseData) GoString() string {
	return s.String()
}

func (s *TripOrderConfirmResponseData) SetGwDescription(v string) *TripOrderConfirmResponseData {
	s.GwDescription = &v
	return s
}

func (s *TripOrderConfirmResponseData) SetGwErrorCode(v int32) *TripOrderConfirmResponseData {
	s.GwErrorCode = &v
	return s
}

type TripOrderConfirmResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s TripOrderConfirmResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TripOrderConfirmResponseExtra) GoString() string {
	return s.String()
}

func (s *TripOrderConfirmResponseExtra) SetSubDescription(v string) *TripOrderConfirmResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TripOrderConfirmResponseExtra) SetSubErrorCode(v int32) *TripOrderConfirmResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TripOrderConfirmResponseExtra) SetDescription(v string) *TripOrderConfirmResponseExtra {
	s.Description = &v
	return s
}

func (s *TripOrderConfirmResponseExtra) SetErrorCode(v int32) *TripOrderConfirmResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TripOrderConfirmResponseExtra) SetLogid(v string) *TripOrderConfirmResponseExtra {
	s.Logid = &v
	return s
}

func (s *TripOrderConfirmResponseExtra) SetNow(v int64) *TripOrderConfirmResponseExtra {
	s.Now = &v
	return s
}

type TripProductOperateRequest struct {
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpType        *int               `json:"op_type,omitempty" xml:"op_type,omitempty" require:"true"`
	ProductIdList []*string          `json:"product_id_list,omitempty" xml:"product_id_list,omitempty" type:"Repeated"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s TripProductOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s TripProductOperateRequest) GoString() string {
	return s.String()
}

func (s *TripProductOperateRequest) SetHeader(v map[string]*string) *TripProductOperateRequest {
	s.Header = v
	return s
}

func (s *TripProductOperateRequest) SetAccessToken(v string) *TripProductOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *TripProductOperateRequest) SetOpType(v int) *TripProductOperateRequest {
	s.OpType = &v
	return s
}

func (s *TripProductOperateRequest) SetProductIdList(v []*string) *TripProductOperateRequest {
	s.ProductIdList = v
	return s
}

func (s *TripProductOperateRequest) SetAccountId(v string) *TripProductOperateRequest {
	s.AccountId = &v
	return s
}

type TripProductOperateResponse struct {
	Data  []*TripProductOperateResponseDataItem `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Extra *TripProductOperateResponseExtra      `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TripProductOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s TripProductOperateResponse) GoString() string {
	return s.String()
}

func (s *TripProductOperateResponse) SetData(v []*TripProductOperateResponseDataItem) *TripProductOperateResponse {
	s.Data = v
	return s
}

func (s *TripProductOperateResponse) SetExtra(v *TripProductOperateResponseExtra) *TripProductOperateResponse {
	s.Extra = v
	return s
}

type TripProductOperateResponseDataItem struct {
	Message   *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	Code      *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s TripProductOperateResponseDataItem) String() string {
	return tea.Prettify(s)
}

func (s TripProductOperateResponseDataItem) GoString() string {
	return s.String()
}

func (s *TripProductOperateResponseDataItem) SetMessage(v string) *TripProductOperateResponseDataItem {
	s.Message = &v
	return s
}

func (s *TripProductOperateResponseDataItem) SetProductId(v string) *TripProductOperateResponseDataItem {
	s.ProductId = &v
	return s
}

func (s *TripProductOperateResponseDataItem) SetCode(v string) *TripProductOperateResponseDataItem {
	s.Code = &v
	return s
}

type TripProductOperateResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TripProductOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TripProductOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *TripProductOperateResponseExtra) SetNow(v int64) *TripProductOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *TripProductOperateResponseExtra) SetSubDescription(v string) *TripProductOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TripProductOperateResponseExtra) SetSubErrorCode(v int32) *TripProductOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TripProductOperateResponseExtra) SetDescription(v string) *TripProductOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *TripProductOperateResponseExtra) SetErrorCode(v int32) *TripProductOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TripProductOperateResponseExtra) SetLogid(v string) *TripProductOperateResponseExtra {
	s.Logid = &v
	return s
}

type TripRefundAuditRequest struct {
	BizUniqKey      *string                               `json:"biz_uniq_key,omitempty" xml:"biz_uniq_key,omitempty" require:"true"`
	OrderId         *string                               `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	RefundFeeAmount *int64                                `json:"refund_fee_amount,omitempty" xml:"refund_fee_amount,omitempty"`
	Vouchers        []*TripRefundAuditRequestVouchersItem `json:"vouchers,omitempty" xml:"vouchers,omitempty" type:"Repeated"`
	Header          map[string]*string                    `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                               `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AuditResult     *int                                  `json:"audit_result,omitempty" xml:"audit_result,omitempty" require:"true"`
}

func (s TripRefundAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequest) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequest) SetBizUniqKey(v string) *TripRefundAuditRequest {
	s.BizUniqKey = &v
	return s
}

func (s *TripRefundAuditRequest) SetOrderId(v string) *TripRefundAuditRequest {
	s.OrderId = &v
	return s
}

func (s *TripRefundAuditRequest) SetRefundFeeAmount(v int64) *TripRefundAuditRequest {
	s.RefundFeeAmount = &v
	return s
}

func (s *TripRefundAuditRequest) SetVouchers(v []*TripRefundAuditRequestVouchersItem) *TripRefundAuditRequest {
	s.Vouchers = v
	return s
}

func (s *TripRefundAuditRequest) SetHeader(v map[string]*string) *TripRefundAuditRequest {
	s.Header = v
	return s
}

func (s *TripRefundAuditRequest) SetAccessToken(v string) *TripRefundAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *TripRefundAuditRequest) SetAuditResult(v int) *TripRefundAuditRequest {
	s.AuditResult = &v
	return s
}

type TripRefundAuditRequestVouchersItem struct {
	Projects   []*TripRefundAuditRequestVouchersItemProjectsItem `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	VerifyTime *int64                                            `json:"verify_time,omitempty" xml:"verify_time,omitempty"`
	Entrance   *TripRefundAuditRequestVouchersItemEntrance       `json:"entrance,omitempty" xml:"entrance,omitempty"`
}

func (s TripRefundAuditRequestVouchersItem) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequestVouchersItem) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequestVouchersItem) SetProjects(v []*TripRefundAuditRequestVouchersItemProjectsItem) *TripRefundAuditRequestVouchersItem {
	s.Projects = v
	return s
}

func (s *TripRefundAuditRequestVouchersItem) SetVerifyTime(v int64) *TripRefundAuditRequestVouchersItem {
	s.VerifyTime = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItem) SetEntrance(v *TripRefundAuditRequestVouchersItemEntrance) *TripRefundAuditRequestVouchersItem {
	s.Entrance = v
	return s
}

type TripRefundAuditRequestVouchersItemEntrance struct {
	IdCards        []*string                                                    `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                      `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId      *string                                                      `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
	Qrcodes        []*string                                                    `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                    `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                    `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*TripRefundAuditRequestVouchersItemEntranceCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
}

func (s TripRefundAuditRequestVouchersItemEntrance) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequestVouchersItemEntrance) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetIdCards(v []*string) *TripRefundAuditRequestVouchersItemEntrance {
	s.IdCards = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetName(v string) *TripRefundAuditRequestVouchersItemEntrance {
	s.Name = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetProjectId(v string) *TripRefundAuditRequestVouchersItemEntrance {
	s.ProjectId = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetQrcodes(v []*string) *TripRefundAuditRequestVouchersItemEntrance {
	s.Qrcodes = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetUrls(v []*string) *TripRefundAuditRequestVouchersItemEntrance {
	s.Urls = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetCertificateNos(v []*string) *TripRefundAuditRequestVouchersItemEntrance {
	s.CertificateNos = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntrance) SetCredentials(v []*TripRefundAuditRequestVouchersItemEntranceCredentialsItem) *TripRefundAuditRequestVouchersItemEntrance {
	s.Credentials = v
	return s
}

type TripRefundAuditRequestVouchersItemEntranceCredentialsItem struct {
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s TripRefundAuditRequestVouchersItemEntranceCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequestVouchersItemEntranceCredentialsItem) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequestVouchersItemEntranceCredentialsItem) SetCredentialType(v int) *TripRefundAuditRequestVouchersItemEntranceCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItemEntranceCredentialsItem) SetCredentialNo(v string) *TripRefundAuditRequestVouchersItemEntranceCredentialsItem {
	s.CredentialNo = &v
	return s
}

type TripRefundAuditRequestVouchersItemProjectsItem struct {
	Qrcodes        []*string                                                        `json:"qrcodes,omitempty" xml:"qrcodes,omitempty" type:"Repeated"`
	Urls           []*string                                                        `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
	CertificateNos []*string                                                        `json:"certificate_nos,omitempty" xml:"certificate_nos,omitempty" type:"Repeated"`
	Credentials    []*TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem `json:"credentials,omitempty" xml:"credentials,omitempty" type:"Repeated"`
	IdCards        []*string                                                        `json:"id_cards,omitempty" xml:"id_cards,omitempty" type:"Repeated"`
	Name           *string                                                          `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId      *string                                                          `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true"`
}

func (s TripRefundAuditRequestVouchersItemProjectsItem) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequestVouchersItemProjectsItem) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetQrcodes(v []*string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.Qrcodes = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetUrls(v []*string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.Urls = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetCertificateNos(v []*string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.CertificateNos = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetCredentials(v []*TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.Credentials = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetIdCards(v []*string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.IdCards = v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetName(v string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.Name = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItem) SetProjectId(v string) *TripRefundAuditRequestVouchersItemProjectsItem {
	s.ProjectId = &v
	return s
}

type TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem struct {
	CredentialType *int    `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	CredentialNo   *string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
}

func (s TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem) GoString() string {
	return s.String()
}

func (s *TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem) SetCredentialType(v int) *TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialType = &v
	return s
}

func (s *TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem) SetCredentialNo(v string) *TripRefundAuditRequestVouchersItemProjectsItemCredentialsItem {
	s.CredentialNo = &v
	return s
}

type TripRefundAuditResponse struct {
	Data  *TripRefundAuditResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *TripRefundAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TripRefundAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditResponse) GoString() string {
	return s.String()
}

func (s *TripRefundAuditResponse) SetData(v *TripRefundAuditResponseData) *TripRefundAuditResponse {
	s.Data = v
	return s
}

func (s *TripRefundAuditResponse) SetExtra(v *TripRefundAuditResponseExtra) *TripRefundAuditResponse {
	s.Extra = v
	return s
}

type TripRefundAuditResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s TripRefundAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditResponseData) GoString() string {
	return s.String()
}

func (s *TripRefundAuditResponseData) SetGwDescription(v string) *TripRefundAuditResponseData {
	s.GwDescription = &v
	return s
}

func (s *TripRefundAuditResponseData) SetGwErrorCode(v int32) *TripRefundAuditResponseData {
	s.GwErrorCode = &v
	return s
}

type TripRefundAuditResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s TripRefundAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TripRefundAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *TripRefundAuditResponseExtra) SetSubErrorCode(v int32) *TripRefundAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TripRefundAuditResponseExtra) SetDescription(v string) *TripRefundAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *TripRefundAuditResponseExtra) SetErrorCode(v int32) *TripRefundAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TripRefundAuditResponseExtra) SetLogid(v string) *TripRefundAuditResponseExtra {
	s.Logid = &v
	return s
}

func (s *TripRefundAuditResponseExtra) SetNow(v int64) *TripRefundAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *TripRefundAuditResponseExtra) SetSubDescription(v string) *TripRefundAuditResponseExtra {
	s.SubDescription = &v
	return s
}

type TripTicketQueryRequest struct {
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ProductOutIds []*string          `json:"product_out_ids,omitempty" xml:"product_out_ids,omitempty" type:"Repeated"`
	Status        *int               `json:"status,omitempty" xml:"status,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Page          *int32             `json:"page,omitempty" xml:"page,omitempty"`
	PageSize      *int32             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	ProductIds    []*string          `json:"product_ids,omitempty" xml:"product_ids,omitempty" type:"Repeated"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s TripTicketQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryRequest) GoString() string {
	return s.String()
}

func (s *TripTicketQueryRequest) SetAccessToken(v string) *TripTicketQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *TripTicketQueryRequest) SetProductOutIds(v []*string) *TripTicketQueryRequest {
	s.ProductOutIds = v
	return s
}

func (s *TripTicketQueryRequest) SetStatus(v int) *TripTicketQueryRequest {
	s.Status = &v
	return s
}

func (s *TripTicketQueryRequest) SetAccountId(v string) *TripTicketQueryRequest {
	s.AccountId = &v
	return s
}

func (s *TripTicketQueryRequest) SetPage(v int32) *TripTicketQueryRequest {
	s.Page = &v
	return s
}

func (s *TripTicketQueryRequest) SetPageSize(v int32) *TripTicketQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TripTicketQueryRequest) SetProductIds(v []*string) *TripTicketQueryRequest {
	s.ProductIds = v
	return s
}

func (s *TripTicketQueryRequest) SetHeader(v map[string]*string) *TripTicketQueryRequest {
	s.Header = v
	return s
}

type TripTicketQueryResponse struct {
	Data  *TripTicketQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *TripTicketQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s TripTicketQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponse) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponse) SetData(v *TripTicketQueryResponseData) *TripTicketQueryResponse {
	s.Data = v
	return s
}

func (s *TripTicketQueryResponse) SetExtra(v *TripTicketQueryResponseExtra) *TripTicketQueryResponse {
	s.Extra = v
	return s
}

type TripTicketQueryResponseData struct {
	GwErrorCode   *int32                                    `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                   `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	PageInfo      *TripTicketQueryResponseDataPageInfo      `json:"page_info,omitempty" xml:"page_info,omitempty" require:"true"`
	Tickets       []*TripTicketQueryResponseDataTicketsItem `json:"tickets,omitempty" xml:"tickets,omitempty" require:"true" type:"Repeated"`
}

func (s TripTicketQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseData) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseData) SetGwErrorCode(v int32) *TripTicketQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *TripTicketQueryResponseData) SetGwDescription(v string) *TripTicketQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *TripTicketQueryResponseData) SetPageInfo(v *TripTicketQueryResponseDataPageInfo) *TripTicketQueryResponseData {
	s.PageInfo = v
	return s
}

func (s *TripTicketQueryResponseData) SetTickets(v []*TripTicketQueryResponseDataTicketsItem) *TripTicketQueryResponseData {
	s.Tickets = v
	return s
}

type TripTicketQueryResponseDataPageInfo struct {
	Page         *int32 `json:"page,omitempty" xml:"page,omitempty"`
	PageSize     *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TotalPage    *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	TotalRecords *int32 `json:"total_records,omitempty" xml:"total_records,omitempty"`
}

func (s TripTicketQueryResponseDataPageInfo) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataPageInfo) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataPageInfo) SetPage(v int32) *TripTicketQueryResponseDataPageInfo {
	s.Page = &v
	return s
}

func (s *TripTicketQueryResponseDataPageInfo) SetPageSize(v int32) *TripTicketQueryResponseDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *TripTicketQueryResponseDataPageInfo) SetTotalPage(v int32) *TripTicketQueryResponseDataPageInfo {
	s.TotalPage = &v
	return s
}

func (s *TripTicketQueryResponseDataPageInfo) SetTotalRecords(v int32) *TripTicketQueryResponseDataPageInfo {
	s.TotalRecords = &v
	return s
}

type TripTicketQueryResponseDataTicketsItem struct {
	BuyLimit             *int32                                                            `json:"buy_limit,omitempty" xml:"buy_limit,omitempty"`
	TicketTypeId         *string                                                           `json:"ticket_type_id,omitempty" xml:"ticket_type_id,omitempty"`
	AuditResult          *TripTicketQueryResponseDataTicketsItemAuditResult                `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	BuyLimitRule         []*TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem         `json:"buy_limit_rule,omitempty" xml:"buy_limit_rule,omitempty" type:"Repeated"`
	Region               *TripTicketQueryResponseDataTicketsItemRegion                     `json:"region,omitempty" xml:"region,omitempty"`
	ProductOutId         *string                                                           `json:"product_out_id,omitempty" xml:"product_out_id,omitempty"`
	UseTicketAddress     *string                                                           `json:"use_ticket_address,omitempty" xml:"use_ticket_address,omitempty"`
	NeedTicket           *bool                                                             `json:"need_ticket,omitempty" xml:"need_ticket,omitempty" require:"true"`
	FeeInclude           *string                                                           `json:"fee_include,omitempty" xml:"fee_include,omitempty" require:"true"`
	CustomerReservedInfo *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo       `json:"customer_reserved_info,omitempty" xml:"customer_reserved_info,omitempty" require:"true"`
	DraftStatus          *int                                                              `json:"draft_status,omitempty" xml:"draft_status,omitempty"`
	PoiIds               []*string                                                         `json:"poi_ids,omitempty" xml:"poi_ids,omitempty" require:"true" type:"Repeated"`
	CategoryId           *int64                                                            `json:"category_id,omitempty" xml:"category_id,omitempty" require:"true"`
	ShowDate             *int32                                                            `json:"show_date,omitempty" xml:"show_date,omitempty"`
	Crowds               []*TripTicketQueryResponseDataTicketsItemCrowdsItem               `json:"crowds,omitempty" xml:"crowds,omitempty" require:"true" type:"Repeated"`
	CrowdOpType          *int                                                              `json:"crowd_op_type,omitempty" xml:"crowd_op_type,omitempty"`
	TestUid              []*string                                                         `json:"test_uid,omitempty" xml:"test_uid,omitempty" type:"Repeated"`
	TicketSpecifications []*TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem `json:"ticket_specifications,omitempty" xml:"ticket_specifications,omitempty" require:"true" type:"Repeated"`
	TicketTypeOutId      *string                                                           `json:"ticket_type_out_id,omitempty" xml:"ticket_type_out_id,omitempty"`
	IsTestData           *bool                                                             `json:"is_test_data,omitempty" xml:"is_test_data,omitempty"`
	OtherNote            *string                                                           `json:"other_note,omitempty" xml:"other_note,omitempty"`
	OpeningTime          *TripTicketQueryResponseDataTicketsItemOpeningTime                `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	ProductName          *string                                                           `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true"`
	FeeNotInclude        *string                                                           `json:"fee_not_include,omitempty" xml:"fee_not_include,omitempty"`
	Status               *int                                                              `json:"status,omitempty" xml:"status,omitempty"`
	StockSyncRule        *TripTicketQueryResponseDataTicketsItemStockSyncRule              `json:"stock_sync_rule,omitempty" xml:"stock_sync_rule,omitempty"`
	ProductId            *string                                                           `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ImgUrls              []*string                                                         `json:"img_urls,omitempty" xml:"img_urls,omitempty" require:"true" type:"Repeated"`
	TicketTypeName       *string                                                           `json:"ticket_type_name,omitempty" xml:"ticket_type_name,omitempty" require:"true"`
	UseTimes             *int32                                                            `json:"use_times,omitempty" xml:"use_times,omitempty" require:"true"`
	TicketRule           *TripTicketQueryResponseDataTicketsItemTicketRule                 `json:"ticket_rule,omitempty" xml:"ticket_rule,omitempty" require:"true"`
	RequireIdCard        *bool                                                             `json:"require_id_card,omitempty" xml:"require_id_card,omitempty"`
	IsTestProduct        *bool                                                             `json:"is_test_product,omitempty" xml:"is_test_product,omitempty"`
	TicketPeriod         *int32                                                            `json:"ticket_period,omitempty" xml:"ticket_period,omitempty" require:"true"`
	PreOrderTime         *TripTicketQueryResponseDataTicketsItemPreOrderTime               `json:"pre_order_time,omitempty" xml:"pre_order_time,omitempty" require:"true"`
	ShowChannel          *int32                                                            `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	MiniProgram          *TripTicketQueryResponseDataTicketsItemMiniProgram                `json:"mini_program,omitempty" xml:"mini_program,omitempty"`
	RefundRule           *TripTicketQueryResponseDataTicketsItemRefundRule                 `json:"refund_rule,omitempty" xml:"refund_rule,omitempty" require:"true"`
	AdmissionTime        *int32                                                            `json:"admission_time,omitempty" xml:"admission_time,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItem) SetBuyLimit(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.BuyLimit = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketTypeId(v string) *TripTicketQueryResponseDataTicketsItem {
	s.TicketTypeId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetAuditResult(v *TripTicketQueryResponseDataTicketsItemAuditResult) *TripTicketQueryResponseDataTicketsItem {
	s.AuditResult = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetBuyLimitRule(v []*TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) *TripTicketQueryResponseDataTicketsItem {
	s.BuyLimitRule = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetRegion(v *TripTicketQueryResponseDataTicketsItemRegion) *TripTicketQueryResponseDataTicketsItem {
	s.Region = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetProductOutId(v string) *TripTicketQueryResponseDataTicketsItem {
	s.ProductOutId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetUseTicketAddress(v string) *TripTicketQueryResponseDataTicketsItem {
	s.UseTicketAddress = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetNeedTicket(v bool) *TripTicketQueryResponseDataTicketsItem {
	s.NeedTicket = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetFeeInclude(v string) *TripTicketQueryResponseDataTicketsItem {
	s.FeeInclude = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetCustomerReservedInfo(v *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) *TripTicketQueryResponseDataTicketsItem {
	s.CustomerReservedInfo = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetDraftStatus(v int) *TripTicketQueryResponseDataTicketsItem {
	s.DraftStatus = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetPoiIds(v []*string) *TripTicketQueryResponseDataTicketsItem {
	s.PoiIds = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetCategoryId(v int64) *TripTicketQueryResponseDataTicketsItem {
	s.CategoryId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetShowDate(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.ShowDate = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetCrowds(v []*TripTicketQueryResponseDataTicketsItemCrowdsItem) *TripTicketQueryResponseDataTicketsItem {
	s.Crowds = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetCrowdOpType(v int) *TripTicketQueryResponseDataTicketsItem {
	s.CrowdOpType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTestUid(v []*string) *TripTicketQueryResponseDataTicketsItem {
	s.TestUid = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketSpecifications(v []*TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) *TripTicketQueryResponseDataTicketsItem {
	s.TicketSpecifications = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketTypeOutId(v string) *TripTicketQueryResponseDataTicketsItem {
	s.TicketTypeOutId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetIsTestData(v bool) *TripTicketQueryResponseDataTicketsItem {
	s.IsTestData = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetOtherNote(v string) *TripTicketQueryResponseDataTicketsItem {
	s.OtherNote = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetOpeningTime(v *TripTicketQueryResponseDataTicketsItemOpeningTime) *TripTicketQueryResponseDataTicketsItem {
	s.OpeningTime = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetProductName(v string) *TripTicketQueryResponseDataTicketsItem {
	s.ProductName = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetFeeNotInclude(v string) *TripTicketQueryResponseDataTicketsItem {
	s.FeeNotInclude = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetStatus(v int) *TripTicketQueryResponseDataTicketsItem {
	s.Status = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetStockSyncRule(v *TripTicketQueryResponseDataTicketsItemStockSyncRule) *TripTicketQueryResponseDataTicketsItem {
	s.StockSyncRule = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetProductId(v string) *TripTicketQueryResponseDataTicketsItem {
	s.ProductId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetImgUrls(v []*string) *TripTicketQueryResponseDataTicketsItem {
	s.ImgUrls = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketTypeName(v string) *TripTicketQueryResponseDataTicketsItem {
	s.TicketTypeName = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetUseTimes(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.UseTimes = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketRule(v *TripTicketQueryResponseDataTicketsItemTicketRule) *TripTicketQueryResponseDataTicketsItem {
	s.TicketRule = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetRequireIdCard(v bool) *TripTicketQueryResponseDataTicketsItem {
	s.RequireIdCard = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetIsTestProduct(v bool) *TripTicketQueryResponseDataTicketsItem {
	s.IsTestProduct = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetTicketPeriod(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.TicketPeriod = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetPreOrderTime(v *TripTicketQueryResponseDataTicketsItemPreOrderTime) *TripTicketQueryResponseDataTicketsItem {
	s.PreOrderTime = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetShowChannel(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.ShowChannel = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetMiniProgram(v *TripTicketQueryResponseDataTicketsItemMiniProgram) *TripTicketQueryResponseDataTicketsItem {
	s.MiniProgram = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetRefundRule(v *TripTicketQueryResponseDataTicketsItemRefundRule) *TripTicketQueryResponseDataTicketsItem {
	s.RefundRule = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItem) SetAdmissionTime(v int32) *TripTicketQueryResponseDataTicketsItem {
	s.AdmissionTime = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemAuditResult struct {
	ProductAuditReason    *string `json:"product_audit_reason,omitempty" xml:"product_audit_reason,omitempty"`
	TicketTypeAuditReason *string `json:"ticket_type_audit_reason,omitempty" xml:"ticket_type_audit_reason,omitempty"`
	IsProductPass         *bool   `json:"is_product_pass,omitempty" xml:"is_product_pass,omitempty" require:"true"`
	IsTicketTypePass      *bool   `json:"is_ticket_type_pass,omitempty" xml:"is_ticket_type_pass,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemAuditResult) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemAuditResult) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemAuditResult) SetProductAuditReason(v string) *TripTicketQueryResponseDataTicketsItemAuditResult {
	s.ProductAuditReason = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemAuditResult) SetTicketTypeAuditReason(v string) *TripTicketQueryResponseDataTicketsItemAuditResult {
	s.TicketTypeAuditReason = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemAuditResult) SetIsProductPass(v bool) *TripTicketQueryResponseDataTicketsItemAuditResult {
	s.IsProductPass = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemAuditResult) SetIsTicketTypePass(v bool) *TripTicketQueryResponseDataTicketsItemAuditResult {
	s.IsTicketTypePass = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem struct {
	TimePeriod   []*string `json:"time_period,omitempty" xml:"time_period,omitempty" type:"Repeated"`
	LimitNum     *int32    `json:"limit_num,omitempty" xml:"limit_num,omitempty" require:"true"`
	LimitRange   *int      `json:"limit_range,omitempty" xml:"limit_range,omitempty" require:"true"`
	LimitSubject *int      `json:"limit_subject,omitempty" xml:"limit_subject,omitempty" require:"true"`
	TimeLength   *int32    `json:"time_length,omitempty" xml:"time_length,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) SetTimePeriod(v []*string) *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem {
	s.TimePeriod = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) SetLimitNum(v int32) *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem {
	s.LimitNum = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) SetLimitRange(v int) *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem {
	s.LimitRange = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) SetLimitSubject(v int) *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem {
	s.LimitSubject = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem) SetTimeLength(v int32) *TripTicketQueryResponseDataTicketsItemBuyLimitRuleItem {
	s.TimeLength = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemCrowdsItem struct {
	CrowdType      *int                                                            `json:"crowd_type,omitempty" xml:"crowd_type,omitempty" require:"true"`
	CrowdCondition *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition `json:"crowd_condition,omitempty" xml:"crowd_condition,omitempty"`
	CrowdNum       *int32                                                          `json:"crowd_num,omitempty" xml:"crowd_num,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItem) SetCrowdType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItem {
	s.CrowdType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItem) SetCrowdCondition(v *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) *TripTicketQueryResponseDataTicketsItemCrowdsItem {
	s.CrowdCondition = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItem) SetCrowdNum(v int32) *TripTicketQueryResponseDataTicketsItemCrowdsItem {
	s.CrowdNum = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition struct {
	AgeHeightOpType      *int                                                                           `json:"age_height_op_type,omitempty" xml:"age_height_op_type,omitempty"`
	CertificateCondition []*int                                                                         `json:"certificate_condition,omitempty" xml:"certificate_condition,omitempty" type:"Repeated"`
	HeightCondition      *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition `json:"height_condition,omitempty" xml:"height_condition,omitempty"`
	Note                 *string                                                                        `json:"note,omitempty" xml:"note,omitempty"`
	ScenicDefaultCrowd   *bool                                                                          `json:"scenic_default_crowd,omitempty" xml:"scenic_default_crowd,omitempty"`
	WeightCondition      *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition `json:"weight_condition,omitempty" xml:"weight_condition,omitempty"`
	AccompanyCondition   *int                                                                           `json:"accompany_condition,omitempty" xml:"accompany_condition,omitempty"`
	AgeCondition         *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition    `json:"age_condition,omitempty" xml:"age_condition,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetAgeHeightOpType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.AgeHeightOpType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetCertificateCondition(v []*int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.CertificateCondition = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetHeightCondition(v *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.HeightCondition = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetNote(v string) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.Note = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetScenicDefaultCrowd(v bool) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.ScenicDefaultCrowd = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetWeightCondition(v *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.WeightCondition = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetAccompanyCondition(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.AccompanyCondition = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition) SetAgeCondition(v *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdCondition {
	s.AgeCondition = v
	return s
}

type TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition struct {
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) SetLowerBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition {
	s.LowerBoundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) SetUpperBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition {
	s.UpperBound = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) SetUpperBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition {
	s.UpperBoundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition) SetLowerBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionAgeCondition {
	s.LowerBound = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition struct {
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) SetUpperBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition {
	s.UpperBoundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) SetLowerBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition {
	s.LowerBound = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) SetLowerBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition {
	s.LowerBoundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition) SetUpperBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionHeightCondition {
	s.UpperBound = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition struct {
	LowerBound     *int64 `json:"lower_bound,omitempty" xml:"lower_bound,omitempty"`
	LowerBoundType *int   `json:"lower_bound_type,omitempty" xml:"lower_bound_type,omitempty"`
	UpperBound     *int64 `json:"upper_bound,omitempty" xml:"upper_bound,omitempty"`
	UpperBoundType *int   `json:"upper_bound_type,omitempty" xml:"upper_bound_type,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) SetLowerBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition {
	s.LowerBound = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) SetLowerBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition {
	s.LowerBoundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) SetUpperBound(v int64) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition {
	s.UpperBound = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition) SetUpperBoundType(v int) *TripTicketQueryResponseDataTicketsItemCrowdsItemCrowdConditionWeightCondition {
	s.UpperBoundType = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemCustomerReservedInfo struct {
	ReservedCrowdNumber *int32 `json:"reserved_crowd_number,omitempty" xml:"reserved_crowd_number,omitempty"`
	ReservedRules       []*int `json:"reserved_rules,omitempty" xml:"reserved_rules,omitempty" type:"Repeated"`
	ReservedType        *int   `json:"reserved_type,omitempty" xml:"reserved_type,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) SetReservedCrowdNumber(v int32) *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo {
	s.ReservedCrowdNumber = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) SetReservedRules(v []*int) *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo {
	s.ReservedRules = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo) SetReservedType(v int) *TripTicketQueryResponseDataTicketsItemCustomerReservedInfo {
	s.ReservedType = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemMiniProgram struct {
	EntryType          *int32  `json:"entry_type,omitempty" xml:"entry_type,omitempty"`
	SaleUnitDetailPath *string `json:"sale_unit_detail_path,omitempty" xml:"sale_unit_detail_path,omitempty"`
	TradeUrl           *string `json:"trade_url,omitempty" xml:"trade_url,omitempty"`
	AppId              *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemMiniProgram) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemMiniProgram) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemMiniProgram) SetEntryType(v int32) *TripTicketQueryResponseDataTicketsItemMiniProgram {
	s.EntryType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemMiniProgram) SetSaleUnitDetailPath(v string) *TripTicketQueryResponseDataTicketsItemMiniProgram {
	s.SaleUnitDetailPath = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemMiniProgram) SetTradeUrl(v string) *TripTicketQueryResponseDataTicketsItemMiniProgram {
	s.TradeUrl = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemMiniProgram) SetAppId(v string) *TripTicketQueryResponseDataTicketsItemMiniProgram {
	s.AppId = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemOpeningTime struct {
	LatestTime   *int64 `json:"latest_time,omitempty" xml:"latest_time,omitempty"`
	EarliestTime *int64 `json:"earliest_time,omitempty" xml:"earliest_time,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemOpeningTime) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemOpeningTime) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemOpeningTime) SetLatestTime(v int64) *TripTicketQueryResponseDataTicketsItemOpeningTime {
	s.LatestTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemOpeningTime) SetEarliestTime(v int64) *TripTicketQueryResponseDataTicketsItemOpeningTime {
	s.EarliestTime = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemPreOrderTime struct {
	EnterTime *int64 `json:"enter_time,omitempty" xml:"enter_time,omitempty"`
	PreTime   *int64 `json:"pre_time,omitempty" xml:"pre_time,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemPreOrderTime) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemPreOrderTime) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemPreOrderTime) SetEnterTime(v int64) *TripTicketQueryResponseDataTicketsItemPreOrderTime {
	s.EnterTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemPreOrderTime) SetPreTime(v int64) *TripTicketQueryResponseDataTicketsItemPreOrderTime {
	s.PreTime = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemRefundRule struct {
	RefundDetails  []*TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem `json:"refund_details,omitempty" xml:"refund_details,omitempty" type:"Repeated"`
	RefundPart     *bool                                                                `json:"refund_part,omitempty" xml:"refund_part,omitempty"`
	RefundType     *int                                                                 `json:"refund_type,omitempty" xml:"refund_type,omitempty" require:"true"`
	AutoRefundTime *int64                                                               `json:"auto_refund_time,omitempty" xml:"auto_refund_time,omitempty"`
	AutoVerifyTime *int64                                                               `json:"auto_verify_time,omitempty" xml:"auto_verify_time,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemRefundRule) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemRefundRule) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRule) SetRefundDetails(v []*TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) *TripTicketQueryResponseDataTicketsItemRefundRule {
	s.RefundDetails = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRule) SetRefundPart(v bool) *TripTicketQueryResponseDataTicketsItemRefundRule {
	s.RefundPart = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRule) SetRefundType(v int) *TripTicketQueryResponseDataTicketsItemRefundRule {
	s.RefundType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRule) SetAutoRefundTime(v int64) *TripTicketQueryResponseDataTicketsItemRefundRule {
	s.AutoRefundTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRule) SetAutoVerifyTime(v int64) *TripTicketQueryResponseDataTicketsItemRefundRule {
	s.AutoVerifyTime = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem struct {
	RefundTime    *int64 `json:"refund_time,omitempty" xml:"refund_time,omitempty" require:"true"`
	RefundFee     *int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty" require:"true"`
	RefundFeeType *int   `json:"refund_fee_type,omitempty" xml:"refund_fee_type,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) SetRefundTime(v int64) *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem {
	s.RefundTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) SetRefundFee(v int64) *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem {
	s.RefundFee = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem) SetRefundFeeType(v int) *TripTicketQueryResponseDataTicketsItemRefundRuleRefundDetailsItem {
	s.RefundFeeType = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemRegion struct {
	RestrictDistricts []*TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem `json:"restrict_districts,omitempty" xml:"restrict_districts,omitempty" type:"Repeated"`
	RestrictFlag      *bool                                                                `json:"restrict_flag,omitempty" xml:"restrict_flag,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemRegion) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemRegion) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemRegion) SetRestrictDistricts(v []*TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem) *TripTicketQueryResponseDataTicketsItemRegion {
	s.RestrictDistricts = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRegion) SetRestrictFlag(v bool) *TripTicketQueryResponseDataTicketsItemRegion {
	s.RestrictFlag = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem struct {
	ProvinceInfoList []*TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem `json:"province_info_list,omitempty" xml:"province_info_list,omitempty" type:"Repeated"`
	DistrictType     *int                                                                                     `json:"district_type,omitempty" xml:"district_type,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem) SetProvinceInfoList(v []*TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem) *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem {
	s.ProvinceInfoList = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem) SetDistrictType(v int) *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItem {
	s.DistrictType = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem struct {
	CityCodeList []*string `json:"city_code_list,omitempty" xml:"city_code_list,omitempty" type:"Repeated"`
	ProvinceCode *string   `json:"province_code,omitempty" xml:"province_code,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem) SetCityCodeList(v []*string) *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem {
	s.CityCodeList = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem) SetProvinceCode(v string) *TripTicketQueryResponseDataTicketsItemRegionRestrictDistrictsItemProvinceInfoListItem {
	s.ProvinceCode = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemStockSyncRule struct {
	SyncOn *bool `json:"sync_on,omitempty" xml:"sync_on,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemStockSyncRule) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemStockSyncRule) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemStockSyncRule) SetSyncOn(v bool) *TripTicketQueryResponseDataTicketsItemStockSyncRule {
	s.SyncOn = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemTicketRule struct {
	CodeExtraNote       *string `json:"code_extra_note,omitempty" xml:"code_extra_note,omitempty"`
	CodeNote            *string `json:"code_note,omitempty" xml:"code_note,omitempty"`
	CodeSendingInfo     []*int  `json:"code_sending_info,omitempty" xml:"code_sending_info,omitempty" require:"true" type:"Repeated"`
	CodeType            *int    `json:"code_type,omitempty" xml:"code_type,omitempty"`
	UrlType             *int    `json:"url_type,omitempty" xml:"url_type,omitempty"`
	ChangeEarliestTime  *string `json:"change_earliest_time,omitempty" xml:"change_earliest_time,omitempty"`
	ChangeLatestTime    *string `json:"change_latest_time,omitempty" xml:"change_latest_time,omitempty"`
	ChangeTicketAddress *string `json:"change_ticket_address,omitempty" xml:"change_ticket_address,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemTicketRule) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemTicketRule) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetCodeExtraNote(v string) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.CodeExtraNote = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetCodeNote(v string) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.CodeNote = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetCodeSendingInfo(v []*int) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.CodeSendingInfo = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetCodeType(v int) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.CodeType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetUrlType(v int) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.UrlType = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetChangeEarliestTime(v string) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.ChangeEarliestTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetChangeLatestTime(v string) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.ChangeLatestTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketRule) SetChangeTicketAddress(v string) *TripTicketQueryResponseDataTicketsItemTicketRule {
	s.ChangeTicketAddress = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem struct {
	SkuId         *string                                                                      `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	SkuName       *string                                                                      `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	SkuOutId      *string                                                                      `json:"sku_out_id,omitempty" xml:"sku_out_id,omitempty"`
	TicketArea    *string                                                                      `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
	TicketSeat    *string                                                                      `json:"ticket_seat,omitempty" xml:"ticket_seat,omitempty"`
	TicketSession *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession `json:"ticket_session,omitempty" xml:"ticket_session,omitempty"`
	FreeMerchant  *bool                                                                        `json:"free_merchant,omitempty" xml:"free_merchant,omitempty"`
	SettleType    *int                                                                         `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetSkuId(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.SkuId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetSkuName(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.SkuName = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetSkuOutId(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.SkuOutId = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetTicketArea(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.TicketArea = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetTicketSeat(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.TicketSeat = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetTicketSession(v *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.TicketSession = v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetFreeMerchant(v bool) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.FreeMerchant = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem) SetSettleType(v int) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItem {
	s.SettleType = &v
	return s
}

type TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession struct {
	TicketSessionTime *string `json:"ticket_session_time,omitempty" xml:"ticket_session_time,omitempty"`
	TicketSessionName *string `json:"ticket_session_name,omitempty" xml:"ticket_session_name,omitempty" require:"true"`
}

func (s TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession) SetTicketSessionTime(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession {
	s.TicketSessionTime = &v
	return s
}

func (s *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession) SetTicketSessionName(v string) *TripTicketQueryResponseDataTicketsItemTicketSpecificationsItemTicketSession {
	s.TicketSessionName = &v
	return s
}

type TripTicketQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s TripTicketQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s TripTicketQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *TripTicketQueryResponseExtra) SetNow(v int64) *TripTicketQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *TripTicketQueryResponseExtra) SetSubDescription(v string) *TripTicketQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *TripTicketQueryResponseExtra) SetSubErrorCode(v int32) *TripTicketQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *TripTicketQueryResponseExtra) SetDescription(v string) *TripTicketQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *TripTicketQueryResponseExtra) SetErrorCode(v int32) *TripTicketQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *TripTicketQueryResponseExtra) SetLogid(v string) *TripTicketQueryResponseExtra {
	s.Logid = &v
	return s
}

type UpdateCommonPlanStatusRequest struct {
	AccessToken    *string                                            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	PlanUpdateList []*UpdateCommonPlanStatusRequestPlanUpdateListItem `json:"plan_update_list,omitempty" xml:"plan_update_list,omitempty" require:"true" type:"Repeated"`
	Header         map[string]*string                                 `json:"header,omitempty" xml:"header,omitempty"`
}

func (s UpdateCommonPlanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonPlanStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommonPlanStatusRequest) SetAccessToken(v string) *UpdateCommonPlanStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateCommonPlanStatusRequest) SetPlanUpdateList(v []*UpdateCommonPlanStatusRequestPlanUpdateListItem) *UpdateCommonPlanStatusRequest {
	s.PlanUpdateList = v
	return s
}

func (s *UpdateCommonPlanStatusRequest) SetHeader(v map[string]*string) *UpdateCommonPlanStatusRequest {
	s.Header = v
	return s
}

type UpdateCommonPlanStatusRequestPlanUpdateListItem struct {
	Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	PlanId *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
}

func (s UpdateCommonPlanStatusRequestPlanUpdateListItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonPlanStatusRequestPlanUpdateListItem) GoString() string {
	return s.String()
}

func (s *UpdateCommonPlanStatusRequestPlanUpdateListItem) SetStatus(v int32) *UpdateCommonPlanStatusRequestPlanUpdateListItem {
	s.Status = &v
	return s
}

func (s *UpdateCommonPlanStatusRequestPlanUpdateListItem) SetPlanId(v int64) *UpdateCommonPlanStatusRequestPlanUpdateListItem {
	s.PlanId = &v
	return s
}

type UpdateCommonPlanStatusResponse struct {
	Data   *UpdateCommonPlanStatusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UpdateCommonPlanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonPlanStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommonPlanStatusResponse) SetData(v *UpdateCommonPlanStatusResponseData) *UpdateCommonPlanStatusResponse {
	s.Data = v
	return s
}

func (s *UpdateCommonPlanStatusResponse) SetErrMsg(v string) *UpdateCommonPlanStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdateCommonPlanStatusResponse) SetErrNo(v int32) *UpdateCommonPlanStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateCommonPlanStatusResponse) SetLogId(v string) *UpdateCommonPlanStatusResponse {
	s.LogId = &v
	return s
}

type UpdateCommonPlanStatusResponseData struct {
	FailPlanIdList []*int64 `json:"fail_plan_id_list,omitempty" xml:"fail_plan_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateCommonPlanStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonPlanStatusResponseData) GoString() string {
	return s.String()
}

func (s *UpdateCommonPlanStatusResponseData) SetFailPlanIdList(v []*int64) *UpdateCommonPlanStatusResponseData {
	s.FailPlanIdList = v
	return s
}

type UpdateCouponMetaStatusRequest struct {
	CouponMetaStatus *int               `json:"coupon_meta_status,omitempty" xml:"coupon_meta_status,omitempty" require:"true"`
	AppId            *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CouponMetaId     *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	Header           map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken      *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UpdateCouponMetaStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCouponMetaStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCouponMetaStatusRequest) SetCouponMetaStatus(v int) *UpdateCouponMetaStatusRequest {
	s.CouponMetaStatus = &v
	return s
}

func (s *UpdateCouponMetaStatusRequest) SetAppId(v string) *UpdateCouponMetaStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCouponMetaStatusRequest) SetCouponMetaId(v string) *UpdateCouponMetaStatusRequest {
	s.CouponMetaId = &v
	return s
}

func (s *UpdateCouponMetaStatusRequest) SetHeader(v map[string]*string) *UpdateCouponMetaStatusRequest {
	s.Header = v
	return s
}

func (s *UpdateCouponMetaStatusRequest) SetAccessToken(v string) *UpdateCouponMetaStatusRequest {
	s.AccessToken = &v
	return s
}

type UpdateCouponMetaStatusResponse struct {
	LogId  *string                             `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UpdateCouponMetaStatusResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                              `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                             `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UpdateCouponMetaStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCouponMetaStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCouponMetaStatusResponse) SetLogId(v string) *UpdateCouponMetaStatusResponse {
	s.LogId = &v
	return s
}

func (s *UpdateCouponMetaStatusResponse) SetData(v *UpdateCouponMetaStatusResponseData) *UpdateCouponMetaStatusResponse {
	s.Data = v
	return s
}

func (s *UpdateCouponMetaStatusResponse) SetErrNo(v int32) *UpdateCouponMetaStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateCouponMetaStatusResponse) SetErrMsg(v string) *UpdateCouponMetaStatusResponse {
	s.ErrMsg = &v
	return s
}

type UpdateCouponMetaStatusResponseData struct {
	ErrNo  *int32  `json:"ErrNo,omitempty" xml:"ErrNo,omitempty" require:"true"`
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty" require:"true"`
	LogID  *string `json:"LogID,omitempty" xml:"LogID,omitempty" require:"true"`
}

func (s UpdateCouponMetaStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCouponMetaStatusResponseData) GoString() string {
	return s.String()
}

func (s *UpdateCouponMetaStatusResponseData) SetErrNo(v int32) *UpdateCouponMetaStatusResponseData {
	s.ErrNo = &v
	return s
}

func (s *UpdateCouponMetaStatusResponseData) SetErrMsg(v string) *UpdateCouponMetaStatusResponseData {
	s.ErrMsg = &v
	return s
}

func (s *UpdateCouponMetaStatusResponseData) SetLogID(v string) *UpdateCouponMetaStatusResponseData {
	s.LogID = &v
	return s
}

type UpdateCouponMetaStockRequest struct {
	CouponMetaId *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	UniqueKey    *string            `json:"unique_key,omitempty" xml:"unique_key,omitempty" require:"true"`
	Number       *int64             `json:"number,omitempty" xml:"number,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken  *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Action       *string            `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	AppId        *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s UpdateCouponMetaStockRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCouponMetaStockRequest) GoString() string {
	return s.String()
}

func (s *UpdateCouponMetaStockRequest) SetCouponMetaId(v string) *UpdateCouponMetaStockRequest {
	s.CouponMetaId = &v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetUniqueKey(v string) *UpdateCouponMetaStockRequest {
	s.UniqueKey = &v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetNumber(v int64) *UpdateCouponMetaStockRequest {
	s.Number = &v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetHeader(v map[string]*string) *UpdateCouponMetaStockRequest {
	s.Header = v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetAccessToken(v string) *UpdateCouponMetaStockRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetAction(v string) *UpdateCouponMetaStockRequest {
	s.Action = &v
	return s
}

func (s *UpdateCouponMetaStockRequest) SetAppId(v string) *UpdateCouponMetaStockRequest {
	s.AppId = &v
	return s
}

type UpdateCouponMetaStockResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UpdateCouponMetaStockResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCouponMetaStockResponse) GoString() string {
	return s.String()
}

func (s *UpdateCouponMetaStockResponse) SetErrNo(v int32) *UpdateCouponMetaStockResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateCouponMetaStockResponse) SetErrMsg(v string) *UpdateCouponMetaStockResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdateCouponMetaStockResponse) SetLogId(v string) *UpdateCouponMetaStockResponse {
	s.LogId = &v
	return s
}

type UpdateOrientedPlanStatusRequest struct {
	PlanUpdateList []*UpdateOrientedPlanStatusRequestPlanUpdateListItem `json:"plan_update_list,omitempty" xml:"plan_update_list,omitempty" require:"true" type:"Repeated"`
	Header         map[string]*string                                   `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string                                              `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UpdateOrientedPlanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrientedPlanStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrientedPlanStatusRequest) SetPlanUpdateList(v []*UpdateOrientedPlanStatusRequestPlanUpdateListItem) *UpdateOrientedPlanStatusRequest {
	s.PlanUpdateList = v
	return s
}

func (s *UpdateOrientedPlanStatusRequest) SetHeader(v map[string]*string) *UpdateOrientedPlanStatusRequest {
	s.Header = v
	return s
}

func (s *UpdateOrientedPlanStatusRequest) SetAccessToken(v string) *UpdateOrientedPlanStatusRequest {
	s.AccessToken = &v
	return s
}

type UpdateOrientedPlanStatusRequestPlanUpdateListItem struct {
	PlanId *int64 `json:"plan_id,omitempty" xml:"plan_id,omitempty" require:"true"`
	Status *int32 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s UpdateOrientedPlanStatusRequestPlanUpdateListItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrientedPlanStatusRequestPlanUpdateListItem) GoString() string {
	return s.String()
}

func (s *UpdateOrientedPlanStatusRequestPlanUpdateListItem) SetPlanId(v int64) *UpdateOrientedPlanStatusRequestPlanUpdateListItem {
	s.PlanId = &v
	return s
}

func (s *UpdateOrientedPlanStatusRequestPlanUpdateListItem) SetStatus(v int32) *UpdateOrientedPlanStatusRequestPlanUpdateListItem {
	s.Status = &v
	return s
}

type UpdateOrientedPlanStatusResponse struct {
	ErrNo  *int32                                `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                               `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UpdateOrientedPlanStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                               `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UpdateOrientedPlanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrientedPlanStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrientedPlanStatusResponse) SetErrNo(v int32) *UpdateOrientedPlanStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateOrientedPlanStatusResponse) SetLogId(v string) *UpdateOrientedPlanStatusResponse {
	s.LogId = &v
	return s
}

func (s *UpdateOrientedPlanStatusResponse) SetData(v *UpdateOrientedPlanStatusResponseData) *UpdateOrientedPlanStatusResponse {
	s.Data = v
	return s
}

func (s *UpdateOrientedPlanStatusResponse) SetErrMsg(v string) *UpdateOrientedPlanStatusResponse {
	s.ErrMsg = &v
	return s
}

type UpdateOrientedPlanStatusResponseData struct {
	FailPlanIdList []*int64 `json:"fail_plan_id_list,omitempty" xml:"fail_plan_id_list,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateOrientedPlanStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrientedPlanStatusResponseData) GoString() string {
	return s.String()
}

func (s *UpdateOrientedPlanStatusResponseData) SetFailPlanIdList(v []*int64) *UpdateOrientedPlanStatusResponseData {
	s.FailPlanIdList = v
	return s
}

type UpdatePromotionActivityStatusRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BizType     *int               `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	ActivityId  *string            `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	Status      *int               `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s UpdatePromotionActivityStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePromotionActivityStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdatePromotionActivityStatusRequest) SetAccessToken(v string) *UpdatePromotionActivityStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdatePromotionActivityStatusRequest) SetBizType(v int) *UpdatePromotionActivityStatusRequest {
	s.BizType = &v
	return s
}

func (s *UpdatePromotionActivityStatusRequest) SetActivityId(v string) *UpdatePromotionActivityStatusRequest {
	s.ActivityId = &v
	return s
}

func (s *UpdatePromotionActivityStatusRequest) SetStatus(v int) *UpdatePromotionActivityStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdatePromotionActivityStatusRequest) SetHeader(v map[string]*string) *UpdatePromotionActivityStatusRequest {
	s.Header = v
	return s
}

type UpdatePromotionActivityStatusResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UpdatePromotionActivityStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePromotionActivityStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePromotionActivityStatusResponse) SetErrNo(v int32) *UpdatePromotionActivityStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdatePromotionActivityStatusResponse) SetErrMsg(v string) *UpdatePromotionActivityStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdatePromotionActivityStatusResponse) SetLogId(v string) *UpdatePromotionActivityStatusResponse {
	s.LogId = &v
	return s
}

type UpdateSimpleQrBindRequest struct {
	Stage                *string            `json:"stage,omitempty" xml:"stage,omitempty" require:"true"`
	ExclusiveQrUrlPrefix *int32             `json:"exclusive_qr_url_prefix,omitempty" xml:"exclusive_qr_url_prefix,omitempty" require:"true"`
	BeforeQrUrl          *string            `json:"before_qr_url,omitempty" xml:"before_qr_url,omitempty" require:"true"`
	QrUrl                *string            `json:"qr_url,omitempty" xml:"qr_url,omitempty" require:"true"`
	LoadPath             *string            `json:"load_path,omitempty" xml:"load_path,omitempty" require:"true"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UpdateSimpleQrBindRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimpleQrBindRequest) GoString() string {
	return s.String()
}

func (s *UpdateSimpleQrBindRequest) SetStage(v string) *UpdateSimpleQrBindRequest {
	s.Stage = &v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetExclusiveQrUrlPrefix(v int32) *UpdateSimpleQrBindRequest {
	s.ExclusiveQrUrlPrefix = &v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetBeforeQrUrl(v string) *UpdateSimpleQrBindRequest {
	s.BeforeQrUrl = &v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetQrUrl(v string) *UpdateSimpleQrBindRequest {
	s.QrUrl = &v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetLoadPath(v string) *UpdateSimpleQrBindRequest {
	s.LoadPath = &v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetHeader(v map[string]*string) *UpdateSimpleQrBindRequest {
	s.Header = v
	return s
}

func (s *UpdateSimpleQrBindRequest) SetAccessToken(v string) *UpdateSimpleQrBindRequest {
	s.AccessToken = &v
	return s
}

type UpdateSimpleQrBindResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UpdateSimpleQrBindResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimpleQrBindResponse) GoString() string {
	return s.String()
}

func (s *UpdateSimpleQrBindResponse) SetErrNo(v int32) *UpdateSimpleQrBindResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateSimpleQrBindResponse) SetErrMsg(v string) *UpdateSimpleQrBindResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdateSimpleQrBindResponse) SetLogId(v string) *UpdateSimpleQrBindResponse {
	s.LogId = &v
	return s
}

type UpdateSimpleQrBindStatusRequest struct {
	Status      *int32             `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	QrUrl       *string            `json:"qr_url,omitempty" xml:"qr_url,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UpdateSimpleQrBindStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimpleQrBindStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSimpleQrBindStatusRequest) SetStatus(v int32) *UpdateSimpleQrBindStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateSimpleQrBindStatusRequest) SetQrUrl(v string) *UpdateSimpleQrBindStatusRequest {
	s.QrUrl = &v
	return s
}

func (s *UpdateSimpleQrBindStatusRequest) SetHeader(v map[string]*string) *UpdateSimpleQrBindStatusRequest {
	s.Header = v
	return s
}

func (s *UpdateSimpleQrBindStatusRequest) SetAccessToken(v string) *UpdateSimpleQrBindStatusRequest {
	s.AccessToken = &v
	return s
}

type UpdateSimpleQrBindStatusResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s UpdateSimpleQrBindStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSimpleQrBindStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSimpleQrBindStatusResponse) SetErrNo(v int32) *UpdateSimpleQrBindStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateSimpleQrBindStatusResponse) SetErrMsg(v string) *UpdateSimpleQrBindStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdateSimpleQrBindStatusResponse) SetLogId(v string) *UpdateSimpleQrBindStatusResponse {
	s.LogId = &v
	return s
}

type UpdateTalentCouponStatusRequest struct {
	TalentAccount *string            `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	CouponMetaId  *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Status        *int               `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	AccountType   *int32             `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
}

func (s UpdateTalentCouponStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTalentCouponStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTalentCouponStatusRequest) SetTalentAccount(v string) *UpdateTalentCouponStatusRequest {
	s.TalentAccount = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetOpenId(v string) *UpdateTalentCouponStatusRequest {
	s.OpenId = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetAppId(v string) *UpdateTalentCouponStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetCouponMetaId(v string) *UpdateTalentCouponStatusRequest {
	s.CouponMetaId = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetHeader(v map[string]*string) *UpdateTalentCouponStatusRequest {
	s.Header = v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetAccessToken(v string) *UpdateTalentCouponStatusRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetStatus(v int) *UpdateTalentCouponStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateTalentCouponStatusRequest) SetAccountType(v int32) *UpdateTalentCouponStatusRequest {
	s.AccountType = &v
	return s
}

type UpdateTalentCouponStatusResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s UpdateTalentCouponStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTalentCouponStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateTalentCouponStatusResponse) SetErrMsg(v string) *UpdateTalentCouponStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *UpdateTalentCouponStatusResponse) SetLogId(v string) *UpdateTalentCouponStatusResponse {
	s.LogId = &v
	return s
}

func (s *UpdateTalentCouponStatusResponse) SetErrNo(v int32) *UpdateTalentCouponStatusResponse {
	s.ErrNo = &v
	return s
}

type UpdateTalentCouponStockRequest struct {
	Number        *int64             `json:"number,omitempty" xml:"number,omitempty" require:"true"`
	AwardNumber   *int64             `json:"award_number,omitempty" xml:"award_number,omitempty"`
	CouponMetaId  *string            `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	AccountType   *int32             `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	Action        *string            `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	TalentAccount *string            `json:"talent_account,omitempty" xml:"talent_account,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UniqueKey     *string            `json:"unique_key,omitempty" xml:"unique_key,omitempty" require:"true"`
	OpenId        *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s UpdateTalentCouponStockRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTalentCouponStockRequest) GoString() string {
	return s.String()
}

func (s *UpdateTalentCouponStockRequest) SetNumber(v int64) *UpdateTalentCouponStockRequest {
	s.Number = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetAwardNumber(v int64) *UpdateTalentCouponStockRequest {
	s.AwardNumber = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetCouponMetaId(v string) *UpdateTalentCouponStockRequest {
	s.CouponMetaId = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetAccountType(v int32) *UpdateTalentCouponStockRequest {
	s.AccountType = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetAction(v string) *UpdateTalentCouponStockRequest {
	s.Action = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetHeader(v map[string]*string) *UpdateTalentCouponStockRequest {
	s.Header = v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetAppId(v string) *UpdateTalentCouponStockRequest {
	s.AppId = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetTalentAccount(v string) *UpdateTalentCouponStockRequest {
	s.TalentAccount = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetAccessToken(v string) *UpdateTalentCouponStockRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetUniqueKey(v string) *UpdateTalentCouponStockRequest {
	s.UniqueKey = &v
	return s
}

func (s *UpdateTalentCouponStockRequest) SetOpenId(v string) *UpdateTalentCouponStockRequest {
	s.OpenId = &v
	return s
}

type UpdateTalentCouponStockResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UpdateTalentCouponStockResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTalentCouponStockResponse) GoString() string {
	return s.String()
}

func (s *UpdateTalentCouponStockResponse) SetLogId(v string) *UpdateTalentCouponStockResponse {
	s.LogId = &v
	return s
}

func (s *UpdateTalentCouponStockResponse) SetErrNo(v int32) *UpdateTalentCouponStockResponse {
	s.ErrNo = &v
	return s
}

func (s *UpdateTalentCouponStockResponse) SetErrMsg(v string) *UpdateTalentCouponStockResponse {
	s.ErrMsg = &v
	return s
}

type UploadImageRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Image       *util.FileField    `json:"image,omitempty" xml:"image,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UploadImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadImageRequest) GoString() string {
	return s.String()
}

func (s *UploadImageRequest) SetOpenId(v string) *UploadImageRequest {
	s.OpenId = &v
	return s
}

func (s *UploadImageRequest) SetImage(v *util.FileField) *UploadImageRequest {
	s.Image = v
	return s
}

func (s *UploadImageRequest) SetHeader(v map[string]*string) *UploadImageRequest {
	s.Header = v
	return s
}

func (s *UploadImageRequest) SetAccessToken(v string) *UploadImageRequest {
	s.AccessToken = &v
	return s
}

type UploadImageResponse struct {
	Data  *UploadImageResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *UploadImageResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s UploadImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponse) GoString() string {
	return s.String()
}

func (s *UploadImageResponse) SetData(v *UploadImageResponseData) *UploadImageResponse {
	s.Data = v
	return s
}

func (s *UploadImageResponse) SetExtra(v *UploadImageResponseExtra) *UploadImageResponse {
	s.Extra = v
	return s
}

type UploadImageResponseData struct {
	Image         *UploadImageResponseDataImage `json:"image,omitempty" xml:"image,omitempty" require:"true"`
	GwErrorCode   *int32                        `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                       `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UploadImageResponseData) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponseData) GoString() string {
	return s.String()
}

func (s *UploadImageResponseData) SetImage(v *UploadImageResponseDataImage) *UploadImageResponseData {
	s.Image = v
	return s
}

func (s *UploadImageResponseData) SetGwErrorCode(v int32) *UploadImageResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *UploadImageResponseData) SetGwDescription(v string) *UploadImageResponseData {
	s.GwDescription = &v
	return s
}

type UploadImageResponseDataImage struct {
	Height  *int32  `json:"height,omitempty" xml:"height,omitempty" require:"true"`
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty" require:"true"`
	Width   *int32  `json:"width,omitempty" xml:"width,omitempty" require:"true"`
}

func (s UploadImageResponseDataImage) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponseDataImage) GoString() string {
	return s.String()
}

func (s *UploadImageResponseDataImage) SetHeight(v int32) *UploadImageResponseDataImage {
	s.Height = &v
	return s
}

func (s *UploadImageResponseDataImage) SetImageId(v string) *UploadImageResponseDataImage {
	s.ImageId = &v
	return s
}

func (s *UploadImageResponseDataImage) SetWidth(v int32) *UploadImageResponseDataImage {
	s.Width = &v
	return s
}

type UploadImageResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s UploadImageResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s UploadImageResponseExtra) GoString() string {
	return s.String()
}

func (s *UploadImageResponseExtra) SetDescription(v string) *UploadImageResponseExtra {
	s.Description = &v
	return s
}

func (s *UploadImageResponseExtra) SetErrorCode(v int32) *UploadImageResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *UploadImageResponseExtra) SetLogid(v string) *UploadImageResponseExtra {
	s.Logid = &v
	return s
}

func (s *UploadImageResponseExtra) SetNow(v int64) *UploadImageResponseExtra {
	s.Now = &v
	return s
}

func (s *UploadImageResponseExtra) SetSubDescription(v string) *UploadImageResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *UploadImageResponseExtra) SetSubErrorCode(v int32) *UploadImageResponseExtra {
	s.SubErrorCode = &v
	return s
}

type UploadUserGroupInfoRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	GroupId     *string            `json:"group_id,omitempty" xml:"group_id,omitempty" require:"true"`
	RoundId     *int64             `json:"round_id,omitempty" xml:"round_id,omitempty" require:"true"`
	RoomId      *string            `json:"room_id,omitempty" xml:"room_id,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UploadUserGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadUserGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *UploadUserGroupInfoRequest) SetOpenId(v string) *UploadUserGroupInfoRequest {
	s.OpenId = &v
	return s
}

func (s *UploadUserGroupInfoRequest) SetGroupId(v string) *UploadUserGroupInfoRequest {
	s.GroupId = &v
	return s
}

func (s *UploadUserGroupInfoRequest) SetRoundId(v int64) *UploadUserGroupInfoRequest {
	s.RoundId = &v
	return s
}

func (s *UploadUserGroupInfoRequest) SetRoomId(v string) *UploadUserGroupInfoRequest {
	s.RoomId = &v
	return s
}

func (s *UploadUserGroupInfoRequest) SetAppId(v string) *UploadUserGroupInfoRequest {
	s.AppId = &v
	return s
}

func (s *UploadUserGroupInfoRequest) SetHeader(v map[string]*string) *UploadUserGroupInfoRequest {
	s.Header = v
	return s
}

func (s *UploadUserGroupInfoRequest) SetAccessToken(v string) *UploadUserGroupInfoRequest {
	s.AccessToken = &v
	return s
}

type UploadUserGroupInfoResponse struct {
	ErrNo  *int64  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s UploadUserGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadUserGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *UploadUserGroupInfoResponse) SetErrNo(v int64) *UploadUserGroupInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *UploadUserGroupInfoResponse) SetErrMsg(v string) *UploadUserGroupInfoResponse {
	s.ErrMsg = &v
	return s
}

type UrlGenerateSchemaRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Path        *string            `json:"path,omitempty" xml:"path,omitempty"`
	Query       *string            `json:"query,omitempty" xml:"query,omitempty"`
	ExpireTime  *int64             `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	NoExpire    *bool              `json:"no_expire,omitempty" xml:"no_expire,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UrlGenerateSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlGenerateSchemaRequest) GoString() string {
	return s.String()
}

func (s *UrlGenerateSchemaRequest) SetAppId(v string) *UrlGenerateSchemaRequest {
	s.AppId = &v
	return s
}

func (s *UrlGenerateSchemaRequest) SetPath(v string) *UrlGenerateSchemaRequest {
	s.Path = &v
	return s
}

func (s *UrlGenerateSchemaRequest) SetQuery(v string) *UrlGenerateSchemaRequest {
	s.Query = &v
	return s
}

func (s *UrlGenerateSchemaRequest) SetExpireTime(v int64) *UrlGenerateSchemaRequest {
	s.ExpireTime = &v
	return s
}

func (s *UrlGenerateSchemaRequest) SetNoExpire(v bool) *UrlGenerateSchemaRequest {
	s.NoExpire = &v
	return s
}

func (s *UrlGenerateSchemaRequest) SetHeader(v map[string]*string) *UrlGenerateSchemaRequest {
	s.Header = v
	return s
}

func (s *UrlGenerateSchemaRequest) SetAccessToken(v string) *UrlGenerateSchemaRequest {
	s.AccessToken = &v
	return s
}

type UrlGenerateSchemaResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty"`
	Data   *UrlGenerateSchemaResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UrlGenerateSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlGenerateSchemaResponse) GoString() string {
	return s.String()
}

func (s *UrlGenerateSchemaResponse) SetErrNo(v int32) *UrlGenerateSchemaResponse {
	s.ErrNo = &v
	return s
}

func (s *UrlGenerateSchemaResponse) SetErrMsg(v string) *UrlGenerateSchemaResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlGenerateSchemaResponse) SetLogId(v string) *UrlGenerateSchemaResponse {
	s.LogId = &v
	return s
}

func (s *UrlGenerateSchemaResponse) SetData(v *UrlGenerateSchemaResponseData) *UrlGenerateSchemaResponse {
	s.Data = v
	return s
}

type UrlGenerateSchemaResponseData struct {
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s UrlGenerateSchemaResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlGenerateSchemaResponseData) GoString() string {
	return s.String()
}

func (s *UrlGenerateSchemaResponseData) SetSchema(v string) *UrlGenerateSchemaResponseData {
	s.Schema = &v
	return s
}

type UrlLinkGenerateRequest struct {
	Path        *string            `json:"path,omitempty" xml:"path,omitempty"`
	Query       *string            `json:"query,omitempty" xml:"query,omitempty"`
	ExpireTime  *int64             `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppName     *string            `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s UrlLinkGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkGenerateRequest) GoString() string {
	return s.String()
}

func (s *UrlLinkGenerateRequest) SetPath(v string) *UrlLinkGenerateRequest {
	s.Path = &v
	return s
}

func (s *UrlLinkGenerateRequest) SetQuery(v string) *UrlLinkGenerateRequest {
	s.Query = &v
	return s
}

func (s *UrlLinkGenerateRequest) SetExpireTime(v int64) *UrlLinkGenerateRequest {
	s.ExpireTime = &v
	return s
}

func (s *UrlLinkGenerateRequest) SetHeader(v map[string]*string) *UrlLinkGenerateRequest {
	s.Header = v
	return s
}

func (s *UrlLinkGenerateRequest) SetAccessToken(v string) *UrlLinkGenerateRequest {
	s.AccessToken = &v
	return s
}

func (s *UrlLinkGenerateRequest) SetAppName(v string) *UrlLinkGenerateRequest {
	s.AppName = &v
	return s
}

func (s *UrlLinkGenerateRequest) SetAppId(v string) *UrlLinkGenerateRequest {
	s.AppId = &v
	return s
}

type UrlLinkGenerateResponse struct {
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UrlLinkGenerateResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UrlLinkGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkGenerateResponse) GoString() string {
	return s.String()
}

func (s *UrlLinkGenerateResponse) SetErrNo(v int32) *UrlLinkGenerateResponse {
	s.ErrNo = &v
	return s
}

func (s *UrlLinkGenerateResponse) SetErrMsg(v string) *UrlLinkGenerateResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlLinkGenerateResponse) SetLogId(v string) *UrlLinkGenerateResponse {
	s.LogId = &v
	return s
}

func (s *UrlLinkGenerateResponse) SetData(v *UrlLinkGenerateResponseData) *UrlLinkGenerateResponse {
	s.Data = v
	return s
}

type UrlLinkGenerateResponseData struct {
	UrlLink *string `json:"url_link,omitempty" xml:"url_link,omitempty" require:"true"`
}

func (s UrlLinkGenerateResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkGenerateResponseData) GoString() string {
	return s.String()
}

func (s *UrlLinkGenerateResponseData) SetUrlLink(v string) *UrlLinkGenerateResponseData {
	s.UrlLink = &v
	return s
}

type UrlLinkQueryInfoRequest struct {
	UrlLink     *string            `json:"url_link,omitempty" xml:"url_link,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s UrlLinkQueryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryInfoRequest) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryInfoRequest) SetUrlLink(v string) *UrlLinkQueryInfoRequest {
	s.UrlLink = &v
	return s
}

func (s *UrlLinkQueryInfoRequest) SetAppId(v string) *UrlLinkQueryInfoRequest {
	s.AppId = &v
	return s
}

func (s *UrlLinkQueryInfoRequest) SetHeader(v map[string]*string) *UrlLinkQueryInfoRequest {
	s.Header = v
	return s
}

func (s *UrlLinkQueryInfoRequest) SetAccessToken(v string) *UrlLinkQueryInfoRequest {
	s.AccessToken = &v
	return s
}

type UrlLinkQueryInfoResponse struct {
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *UrlLinkQueryInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s UrlLinkQueryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryInfoResponse) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryInfoResponse) SetErrMsg(v string) *UrlLinkQueryInfoResponse {
	s.ErrMsg = &v
	return s
}

func (s *UrlLinkQueryInfoResponse) SetLogId(v string) *UrlLinkQueryInfoResponse {
	s.LogId = &v
	return s
}

func (s *UrlLinkQueryInfoResponse) SetData(v *UrlLinkQueryInfoResponseData) *UrlLinkQueryInfoResponse {
	s.Data = v
	return s
}

func (s *UrlLinkQueryInfoResponse) SetErrNo(v int32) *UrlLinkQueryInfoResponse {
	s.ErrNo = &v
	return s
}

type UrlLinkQueryInfoResponseData struct {
	Path       *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
	Query      *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	CreateTime *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	ExpireTime *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	AppName    *string `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	AppId      *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s UrlLinkQueryInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s UrlLinkQueryInfoResponseData) GoString() string {
	return s.String()
}

func (s *UrlLinkQueryInfoResponseData) SetPath(v string) *UrlLinkQueryInfoResponseData {
	s.Path = &v
	return s
}

func (s *UrlLinkQueryInfoResponseData) SetQuery(v string) *UrlLinkQueryInfoResponseData {
	s.Query = &v
	return s
}

func (s *UrlLinkQueryInfoResponseData) SetCreateTime(v int64) *UrlLinkQueryInfoResponseData {
	s.CreateTime = &v
	return s
}

func (s *UrlLinkQueryInfoResponseData) SetExpireTime(v int64) *UrlLinkQueryInfoResponseData {
	s.ExpireTime = &v
	return s
}

func (s *UrlLinkQueryInfoResponseData) SetAppName(v string) *UrlLinkQueryInfoResponseData {
	s.AppName = &v
	return s
}

func (s *UrlLinkQueryInfoResponseData) SetAppId(v string) *UrlLinkQueryInfoResponseData {
	s.AppId = &v
	return s
}
