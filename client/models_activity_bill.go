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

type ActivityCreatePromotionActivityRequest struct {
	PromotionActivity *ActivityCreatePromotionActivityRequestPromotionActivity `json:"promotion_activity,omitempty" xml:"promotion_activity,omitempty" require:"true"`
	Header            map[string]*string                                       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ActivityCreatePromotionActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequest) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequest) SetPromotionActivity(v *ActivityCreatePromotionActivityRequestPromotionActivity) *ActivityCreatePromotionActivityRequest {
	s.PromotionActivity = v
	return s
}

func (s *ActivityCreatePromotionActivityRequest) SetHeader(v map[string]*string) *ActivityCreatePromotionActivityRequest {
	s.Header = v
	return s
}

func (s *ActivityCreatePromotionActivityRequest) SetAccessToken(v string) *ActivityCreatePromotionActivityRequest {
	s.AccessToken = &v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivity struct {
	LiveActivity    *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity    `json:"live_activity,omitempty" xml:"live_activity,omitempty"`
	FeedActivity    *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity    `json:"feed_activity,omitempty" xml:"feed_activity,omitempty"`
	CouponMetaId    *string                                                                 `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty" require:"true"`
	SendScene       *int                                                                    `json:"send_scene,omitempty" xml:"send_scene,omitempty" require:"true"`
	ImActivity      *ActivityCreatePromotionActivityRequestPromotionActivityImActivity      `json:"im_activity,omitempty" xml:"im_activity,omitempty"`
	ActivityName    *string                                                                 `json:"activity_name,omitempty" xml:"activity_name,omitempty" require:"true"`
	ReceiveLimit    *int64                                                                  `json:"receive_limit,omitempty" xml:"receive_limit,omitempty" require:"true"`
	SidebarActivity *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity `json:"sidebar_activity,omitempty" xml:"sidebar_activity,omitempty"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivity) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetLiveActivity(v *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.LiveActivity = v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetFeedActivity(v *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.FeedActivity = v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetCouponMetaId(v string) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.CouponMetaId = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetSendScene(v int) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.SendScene = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetImActivity(v *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.ImActivity = v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetActivityName(v string) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.ActivityName = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetReceiveLimit(v int64) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.ReceiveLimit = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivity) SetSidebarActivity(v *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) *ActivityCreatePromotionActivityRequestPromotionActivity {
	s.SidebarActivity = v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity struct {
	IsRecommendedStrategy *bool  `json:"is_recommended_strategy,omitempty" xml:"is_recommended_strategy,omitempty"`
	CanReceiveNum         *int64 `json:"can_receive_num,omitempty" xml:"can_receive_num,omitempty" require:"true"`
	ValidBeginTime        *int64 `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty" require:"true"`
	ValidEndTime          *int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty" require:"true"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) SetIsRecommendedStrategy(v bool) *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity {
	s.IsRecommendedStrategy = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) SetCanReceiveNum(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity {
	s.CanReceiveNum = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) SetValidBeginTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity) SetValidEndTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityFeedActivity {
	s.ValidEndTime = &v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivityImActivity struct {
	ValidEndTime       *int64                                                                               `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty" require:"true"`
	ShareFissionConfig *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig `json:"share_fission_config,omitempty" xml:"share_fission_config,omitempty"`
	CouponTaskType     *int                                                                                 `json:"coupon_task_type,omitempty" xml:"coupon_task_type,omitempty"`
	NoticeText         *string                                                                              `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
	ValidBeginTime     *int64                                                                               `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty" require:"true"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityImActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityImActivity) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) SetValidEndTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityImActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) SetShareFissionConfig(v *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) *ActivityCreatePromotionActivityRequestPromotionActivityImActivity {
	s.ShareFissionConfig = v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) SetCouponTaskType(v int) *ActivityCreatePromotionActivityRequestPromotionActivityImActivity {
	s.CouponTaskType = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) SetNoticeText(v string) *ActivityCreatePromotionActivityRequestPromotionActivityImActivity {
	s.NoticeText = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivity) SetValidBeginTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityImActivity {
	s.ValidBeginTime = &v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig struct {
	AwardType      *int   `json:"award_type,omitempty" xml:"award_type,omitempty" require:"true"`
	AwardCouponId  *int64 `json:"award_coupon_id,omitempty" xml:"award_coupon_id,omitempty"`
	AwardCouponNum *int32 `json:"award_coupon_num,omitempty" xml:"award_coupon_num,omitempty"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) SetAwardType(v int) *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig {
	s.AwardType = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) SetAwardCouponId(v int64) *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig {
	s.AwardCouponId = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig) SetAwardCouponNum(v int32) *ActivityCreatePromotionActivityRequestPromotionActivityImActivityShareFissionConfig {
	s.AwardCouponNum = &v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity struct {
	CouponTaskType  *int    `json:"coupon_task_type,omitempty" xml:"coupon_task_type,omitempty"`
	ClueComponentId *string `json:"clue_component_id,omitempty" xml:"clue_component_id,omitempty"`
	NoticeText      *string `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) SetCouponTaskType(v int) *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity {
	s.CouponTaskType = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) SetClueComponentId(v string) *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity {
	s.ClueComponentId = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity) SetNoticeText(v string) *ActivityCreatePromotionActivityRequestPromotionActivityLiveActivity {
	s.NoticeText = &v
	return s
}

type ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity struct {
	CompleteTimes          *int32  `json:"complete_times,omitempty" xml:"complete_times,omitempty"`
	ActionTrigger          *int    `json:"action_trigger,omitempty" xml:"action_trigger,omitempty" require:"true"`
	JumpText               *string `json:"jump_text,omitempty" xml:"jump_text,omitempty" require:"true"`
	ValidEndTime           *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty" require:"true"`
	NeedBind               *bool   `json:"need_bind,omitempty" xml:"need_bind,omitempty"`
	IsCustomEducationPopup *bool   `json:"is_custom_education_popup,omitempty" xml:"is_custom_education_popup,omitempty"`
	ValidBeginTime         *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty" require:"true"`
	IsCustomAwardPopup     *bool   `json:"is_custom_award_popup,omitempty" xml:"is_custom_award_popup,omitempty"`
	ShortTitle             *string `json:"short_title,omitempty" xml:"short_title,omitempty" require:"true"`
	HighValueContent       *string `json:"high_value_content,omitempty" xml:"high_value_content,omitempty" require:"true"`
	RecentBubbleText       *string `json:"recent_bubble_text,omitempty" xml:"recent_bubble_text,omitempty" require:"true"`
}

func (s ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetCompleteTimes(v int32) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.CompleteTimes = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetActionTrigger(v int) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.ActionTrigger = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetJumpText(v string) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.JumpText = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetValidEndTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetNeedBind(v bool) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.NeedBind = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetIsCustomEducationPopup(v bool) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.IsCustomEducationPopup = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetValidBeginTime(v int64) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetIsCustomAwardPopup(v bool) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.IsCustomAwardPopup = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetShortTitle(v string) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.ShortTitle = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetHighValueContent(v string) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.HighValueContent = &v
	return s
}

func (s *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity) SetRecentBubbleText(v string) *ActivityCreatePromotionActivityRequestPromotionActivitySidebarActivity {
	s.RecentBubbleText = &v
	return s
}

type ActivityCreatePromotionActivityResponse struct {
	ErrMsg *string                                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *ActivityCreatePromotionActivityResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s ActivityCreatePromotionActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityResponse) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityResponse) SetErrMsg(v string) *ActivityCreatePromotionActivityResponse {
	s.ErrMsg = &v
	return s
}

func (s *ActivityCreatePromotionActivityResponse) SetLogId(v string) *ActivityCreatePromotionActivityResponse {
	s.LogId = &v
	return s
}

func (s *ActivityCreatePromotionActivityResponse) SetData(v *ActivityCreatePromotionActivityResponseData) *ActivityCreatePromotionActivityResponse {
	s.Data = v
	return s
}

func (s *ActivityCreatePromotionActivityResponse) SetErrNo(v int32) *ActivityCreatePromotionActivityResponse {
	s.ErrNo = &v
	return s
}

type ActivityCreatePromotionActivityResponseData struct {
	ActivityId *string `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ActivityCreatePromotionActivityResponseData) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreatePromotionActivityResponseData) GoString() string {
	return s.String()
}

func (s *ActivityCreatePromotionActivityResponseData) SetActivityId(v string) *ActivityCreatePromotionActivityResponseData {
	s.ActivityId = &v
	return s
}

type ActivityCreateRequest struct {
	EndTime                    *int64                                                 `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	StartTime                  *int64                                                 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	ActivityName               *string                                                `json:"activity_name,omitempty" xml:"activity_name,omitempty" require:"true"`
	CreateBusinessTaskInfoList []*ActivityCreateRequestCreateBusinessTaskInfoListItem `json:"create_business_task_info_list,omitempty" xml:"create_business_task_info_list,omitempty" require:"true" type:"Repeated"`
	Header                     map[string]*string                                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken                *string                                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s ActivityCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequest) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequest) SetEndTime(v int64) *ActivityCreateRequest {
	s.EndTime = &v
	return s
}

func (s *ActivityCreateRequest) SetStartTime(v int64) *ActivityCreateRequest {
	s.StartTime = &v
	return s
}

func (s *ActivityCreateRequest) SetActivityName(v string) *ActivityCreateRequest {
	s.ActivityName = &v
	return s
}

func (s *ActivityCreateRequest) SetCreateBusinessTaskInfoList(v []*ActivityCreateRequestCreateBusinessTaskInfoListItem) *ActivityCreateRequest {
	s.CreateBusinessTaskInfoList = v
	return s
}

func (s *ActivityCreateRequest) SetHeader(v map[string]*string) *ActivityCreateRequest {
	s.Header = v
	return s
}

func (s *ActivityCreateRequest) SetAccessToken(v string) *ActivityCreateRequest {
	s.AccessToken = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItem struct {
	EndTime       *int64                                                            `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime     *int64                                                            `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TaskEventInfo *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo `json:"task_event_info,omitempty" xml:"task_event_info,omitempty"`
	TaskName      *string                                                           `json:"task_name,omitempty" xml:"task_name,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItem) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItem) SetEndTime(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItem {
	s.EndTime = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItem) SetStartTime(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItem {
	s.StartTime = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItem) SetTaskEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItem {
	s.TaskEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItem) SetTaskName(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItem {
	s.TaskName = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo struct {
	PostingVideEventInfo             *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo             `json:"posting_vide_event_info,omitempty" xml:"posting_vide_event_info,omitempty"`
	ShortVideoShareEventInfo         *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo         `json:"short_video_share_event_info,omitempty" xml:"short_video_share_event_info,omitempty"`
	LiveDiggEventInfo                *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo                `json:"live_digg_event_info,omitempty" xml:"live_digg_event_info,omitempty"`
	LiveCommentEventInfo             *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo             `json:"live_comment_event_info,omitempty" xml:"live_comment_event_info,omitempty"`
	LiveShareEventInfo               *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo               `json:"live_share_event_info,omitempty" xml:"live_share_event_info,omitempty"`
	LiveWatchEventInfo               *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo               `json:"live_watch_event_info,omitempty" xml:"live_watch_event_info,omitempty"`
	ShortVideoCommentEventInfo       *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo       `json:"short_video_comment_event_info,omitempty" xml:"short_video_comment_event_info,omitempty"`
	ShortVideoCollectionEventInfo    *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo    `json:"short_video_collection_event_info,omitempty" xml:"short_video_collection_event_info,omitempty"`
	AccountFollowEventInfo           *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo           `json:"account_follow_event_info,omitempty" xml:"account_follow_event_info,omitempty"`
	ShortVideoDiggEventInfo          *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo          `json:"short_video_digg_event_info,omitempty" xml:"short_video_digg_event_info,omitempty"`
	ShortVideoFinishPlayingEventInfo *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo `json:"short_video_finish_playing_event_info,omitempty" xml:"short_video_finish_playing_event_info,omitempty"`
	CommonEventInfo                  *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo                  `json:"common_event_info,omitempty" xml:"common_event_info,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetPostingVideEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.PostingVideEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetShortVideoShareEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoShareEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetLiveDiggEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.LiveDiggEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetLiveCommentEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.LiveCommentEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetLiveShareEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.LiveShareEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetLiveWatchEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.LiveWatchEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCommentEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCommentEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCollectionEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCollectionEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetAccountFollowEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.AccountFollowEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetShortVideoDiggEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoDiggEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetShortVideoFinishPlayingEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoFinishPlayingEventInfo = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo) SetCommonEventInfo(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfo {
	s.CommonEventInfo = v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) SetAwemeId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo struct {
	TaskDependence     *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence           `json:"task_dependence,omitempty" xml:"task_dependence,omitempty" require:"true"`
	TaskTypeEnum       *int                                                                                                     `json:"task_type_enum,omitempty" xml:"task_type_enum,omitempty" require:"true"`
	CompleteConditions []*ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem `json:"complete_conditions,omitempty" xml:"complete_conditions,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskDependence(v *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskDependence = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskTypeEnum(v int) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskTypeEnum = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetCompleteConditions(v []*ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.CompleteConditions = v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem struct {
	ConditionType              *int     `json:"condition_type,omitempty" xml:"condition_type,omitempty" require:"true"`
	TaskCompleteConditionStage []*int64 `json:"task_complete_condition_stage,omitempty" xml:"task_complete_condition_stage,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetConditionType(v int) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.ConditionType = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetTaskCompleteConditionStage(v []*int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.TaskCompleteConditionStage = v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence struct {
	DependenceList []*ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem `json:"dependence_list,omitempty" xml:"dependence_list,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) SetDependenceList(v []*ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence {
	s.DependenceList = v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem struct {
	DependenceTypeEnum *int    `json:"dependence_type_enum,omitempty" xml:"dependence_type_enum,omitempty" require:"true"`
	DependenceValue    *string `json:"dependence_value,omitempty" xml:"dependence_value,omitempty" require:"true"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceTypeEnum(v int) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceTypeEnum = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceValue(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceValue = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetAwemeId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.GameId = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo struct {
	AwemeId   *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetAwemeId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetDiggCount(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.DiggCount = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.GameId = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo struct {
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetAwemeId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo struct {
	AwemeId       *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId        *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	WatchDuration *int64  `json:"watch_duration,omitempty" xml:"watch_duration,omitempty" require:"true"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetAwemeId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetWatchDuration(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.WatchDuration = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo struct {
	PlayVideoStage []*int64 `json:"play_video_stage,omitempty" xml:"play_video_stage,omitempty" type:"Repeated"`
	CommentStage   []*int64 `json:"comment_stage,omitempty" xml:"comment_stage,omitempty" type:"Repeated"`
	DiggStage      []*int64 `json:"digg_stage,omitempty" xml:"digg_stage,omitempty" type:"Repeated"`
	NumStage       []*int64 `json:"num_stage,omitempty" xml:"num_stage,omitempty" type:"Repeated"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetPlayVideoStage(v []*int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.PlayVideoStage = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetCommentStage(v []*int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.CommentStage = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetDiggStage(v []*int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.DiggStage = v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetNumStage(v []*int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.NumStage = v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo struct {
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetItemId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetVideoUrl(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo struct {
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetVideoUrl(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetAnchorId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetItemId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.ItemId = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo struct {
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId    *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl  *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId  *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetItemId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetVideoUrl(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetAnchorId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetDiggCount(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.DiggCount = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo struct {
	AnchorId             *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	FinishPlayVidoeCount *int64  `json:"finish_play_vidoe_count,omitempty" xml:"finish_play_vidoe_count,omitempty"`
	GameId               *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId               *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl             *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetAnchorId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetFinishPlayVidoeCount(v int64) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.FinishPlayVidoeCount = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetItemId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetVideoUrl(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo struct {
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetGameId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetItemId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetVideoUrl(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetAnchorId(v string) *ActivityCreateRequestCreateBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityCreateResponse struct {
	ActivityId         *int64   `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	BusinessTaskIdList []*int64 `json:"business_task_id_list,omitempty" xml:"business_task_id_list,omitempty" type:"Repeated"`
}

func (s ActivityCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityCreateResponse) GoString() string {
	return s.String()
}

func (s *ActivityCreateResponse) SetActivityId(v int64) *ActivityCreateResponse {
	s.ActivityId = &v
	return s
}

func (s *ActivityCreateResponse) SetBusinessTaskIdList(v []*int64) *ActivityCreateResponse {
	s.BusinessTaskIdList = v
	return s
}

type ActivityModifyPromotionActivityRequest struct {
	Header            map[string]*string                                       `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                                  `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BizType           *int                                                     `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	IsRenewal         *bool                                                    `json:"is_renewal,omitempty" xml:"is_renewal,omitempty"`
	PromotionActivity *ActivityModifyPromotionActivityRequestPromotionActivity `json:"promotion_activity,omitempty" xml:"promotion_activity,omitempty" require:"true"`
}

func (s ActivityModifyPromotionActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequest) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequest) SetHeader(v map[string]*string) *ActivityModifyPromotionActivityRequest {
	s.Header = v
	return s
}

func (s *ActivityModifyPromotionActivityRequest) SetAccessToken(v string) *ActivityModifyPromotionActivityRequest {
	s.AccessToken = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequest) SetBizType(v int) *ActivityModifyPromotionActivityRequest {
	s.BizType = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequest) SetIsRenewal(v bool) *ActivityModifyPromotionActivityRequest {
	s.IsRenewal = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequest) SetPromotionActivity(v *ActivityModifyPromotionActivityRequestPromotionActivity) *ActivityModifyPromotionActivityRequest {
	s.PromotionActivity = v
	return s
}

type ActivityModifyPromotionActivityRequestPromotionActivity struct {
	ActivityId      *string                                                                 `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	SidebarActivity *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity `json:"sidebar_activity,omitempty" xml:"sidebar_activity,omitempty"`
	ActivityName    *string                                                                 `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	ImActivity      *ActivityModifyPromotionActivityRequestPromotionActivityImActivity      `json:"im_activity,omitempty" xml:"im_activity,omitempty"`
	LiveActivity    *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity    `json:"live_activity,omitempty" xml:"live_activity,omitempty"`
	FeedActivity    *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity    `json:"feed_activity,omitempty" xml:"feed_activity,omitempty"`
}

func (s ActivityModifyPromotionActivityRequestPromotionActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequestPromotionActivity) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetActivityId(v string) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.ActivityId = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetSidebarActivity(v *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.SidebarActivity = v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetActivityName(v string) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.ActivityName = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetImActivity(v *ActivityModifyPromotionActivityRequestPromotionActivityImActivity) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.ImActivity = v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetLiveActivity(v *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.LiveActivity = v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivity) SetFeedActivity(v *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) *ActivityModifyPromotionActivityRequestPromotionActivity {
	s.FeedActivity = v
	return s
}

type ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity struct {
	ValidBeginTime        *int64 `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	ValidEndTime          *int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	IsRecommendedStrategy *bool  `json:"is_recommended_strategy,omitempty" xml:"is_recommended_strategy,omitempty"`
	CanReceiveNum         *int64 `json:"can_receive_num,omitempty" xml:"can_receive_num,omitempty"`
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) SetValidBeginTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) SetValidEndTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) SetIsRecommendedStrategy(v bool) *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity {
	s.IsRecommendedStrategy = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity) SetCanReceiveNum(v int64) *ActivityModifyPromotionActivityRequestPromotionActivityFeedActivity {
	s.CanReceiveNum = &v
	return s
}

type ActivityModifyPromotionActivityRequestPromotionActivityImActivity struct {
	ValidBeginTime *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	ValidEndTime   *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	NoticeText     *string `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityImActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityImActivity) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityImActivity) SetValidBeginTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivityImActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityImActivity) SetValidEndTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivityImActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityImActivity) SetNoticeText(v string) *ActivityModifyPromotionActivityRequestPromotionActivityImActivity {
	s.NoticeText = &v
	return s
}

type ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity struct {
	NoticeText      *string `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
	ClueComponentId *string `json:"clue_component_id,omitempty" xml:"clue_component_id,omitempty"`
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity) SetNoticeText(v string) *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity {
	s.NoticeText = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity) SetClueComponentId(v string) *ActivityModifyPromotionActivityRequestPromotionActivityLiveActivity {
	s.ClueComponentId = &v
	return s
}

type ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity struct {
	IsCustomAwardPopup     *bool   `json:"is_custom_award_popup,omitempty" xml:"is_custom_award_popup,omitempty"`
	ValidEndTime           *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	HighValueContent       *string `json:"high_value_content,omitempty" xml:"high_value_content,omitempty"`
	ValidBeginTime         *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	IsCustomEducationPopup *bool   `json:"is_custom_education_popup,omitempty" xml:"is_custom_education_popup,omitempty"`
	JumpText               *string `json:"jump_text,omitempty" xml:"jump_text,omitempty"`
	ShortTitle             *string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	RecentBubbleText       *string `json:"recent_bubble_text,omitempty" xml:"recent_bubble_text,omitempty"`
}

func (s ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetIsCustomAwardPopup(v bool) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.IsCustomAwardPopup = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetValidEndTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetHighValueContent(v string) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.HighValueContent = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetValidBeginTime(v int64) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetIsCustomEducationPopup(v bool) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.IsCustomEducationPopup = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetJumpText(v string) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.JumpText = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetShortTitle(v string) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.ShortTitle = &v
	return s
}

func (s *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity) SetRecentBubbleText(v string) *ActivityModifyPromotionActivityRequestPromotionActivitySidebarActivity {
	s.RecentBubbleText = &v
	return s
}

type ActivityModifyPromotionActivityResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ActivityModifyPromotionActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyPromotionActivityResponse) GoString() string {
	return s.String()
}

func (s *ActivityModifyPromotionActivityResponse) SetErrNo(v int32) *ActivityModifyPromotionActivityResponse {
	s.ErrNo = &v
	return s
}

func (s *ActivityModifyPromotionActivityResponse) SetErrMsg(v string) *ActivityModifyPromotionActivityResponse {
	s.ErrMsg = &v
	return s
}

func (s *ActivityModifyPromotionActivityResponse) SetLogId(v string) *ActivityModifyPromotionActivityResponse {
	s.LogId = &v
	return s
}

type ActivityModifyRequest struct {
	EndTime                    *int64                                                 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	AccessToken                *string                                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ModifyBusinessTaskInfoList []*ActivityModifyRequestModifyBusinessTaskInfoListItem `json:"modify_business_task_info_list,omitempty" xml:"modify_business_task_info_list,omitempty" type:"Repeated"`
	ActivityId                 *int64                                                 `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
	DelBusinessTaskIdList      []*int64                                               `json:"del_business_task_id_list,omitempty" xml:"del_business_task_id_list,omitempty" type:"Repeated"`
	Header                     map[string]*string                                     `json:"header,omitempty" xml:"header,omitempty"`
	StartTime                  *int64                                                 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	ActivityName               *string                                                `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	AddBusinessTaskInfoList    []*ActivityModifyRequestAddBusinessTaskInfoListItem    `json:"add_business_task_info_list,omitempty" xml:"add_business_task_info_list,omitempty" type:"Repeated"`
}

func (s ActivityModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequest) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequest) SetEndTime(v int64) *ActivityModifyRequest {
	s.EndTime = &v
	return s
}

func (s *ActivityModifyRequest) SetAccessToken(v string) *ActivityModifyRequest {
	s.AccessToken = &v
	return s
}

func (s *ActivityModifyRequest) SetModifyBusinessTaskInfoList(v []*ActivityModifyRequestModifyBusinessTaskInfoListItem) *ActivityModifyRequest {
	s.ModifyBusinessTaskInfoList = v
	return s
}

func (s *ActivityModifyRequest) SetActivityId(v int64) *ActivityModifyRequest {
	s.ActivityId = &v
	return s
}

func (s *ActivityModifyRequest) SetDelBusinessTaskIdList(v []*int64) *ActivityModifyRequest {
	s.DelBusinessTaskIdList = v
	return s
}

func (s *ActivityModifyRequest) SetHeader(v map[string]*string) *ActivityModifyRequest {
	s.Header = v
	return s
}

func (s *ActivityModifyRequest) SetStartTime(v int64) *ActivityModifyRequest {
	s.StartTime = &v
	return s
}

func (s *ActivityModifyRequest) SetActivityName(v string) *ActivityModifyRequest {
	s.ActivityName = &v
	return s
}

func (s *ActivityModifyRequest) SetAddBusinessTaskInfoList(v []*ActivityModifyRequestAddBusinessTaskInfoListItem) *ActivityModifyRequest {
	s.AddBusinessTaskInfoList = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItem struct {
	TaskId        *int64                                                         `json:"task_id,omitempty" xml:"task_id,omitempty"`
	TaskName      *string                                                        `json:"task_name,omitempty" xml:"task_name,omitempty"`
	EndTime       *int64                                                         `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime     *int64                                                         `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TaskEventInfo *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo `json:"task_event_info,omitempty" xml:"task_event_info,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItem) SetTaskId(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItem {
	s.TaskId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItem) SetTaskName(v string) *ActivityModifyRequestAddBusinessTaskInfoListItem {
	s.TaskName = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItem) SetEndTime(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItem {
	s.EndTime = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItem) SetStartTime(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItem {
	s.StartTime = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItem) SetTaskEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItem {
	s.TaskEventInfo = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo struct {
	AccountFollowEventInfo           *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo           `json:"account_follow_event_info,omitempty" xml:"account_follow_event_info,omitempty"`
	LiveCommentEventInfo             *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo             `json:"live_comment_event_info,omitempty" xml:"live_comment_event_info,omitempty"`
	ShortVideoCollectionEventInfo    *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo    `json:"short_video_collection_event_info,omitempty" xml:"short_video_collection_event_info,omitempty"`
	LiveShareEventInfo               *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo               `json:"live_share_event_info,omitempty" xml:"live_share_event_info,omitempty"`
	LiveWatchEventInfo               *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo               `json:"live_watch_event_info,omitempty" xml:"live_watch_event_info,omitempty"`
	ShortVideoCommentEventInfo       *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo       `json:"short_video_comment_event_info,omitempty" xml:"short_video_comment_event_info,omitempty"`
	LiveDiggEventInfo                *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo                `json:"live_digg_event_info,omitempty" xml:"live_digg_event_info,omitempty"`
	PostingVideEventInfo             *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo             `json:"posting_vide_event_info,omitempty" xml:"posting_vide_event_info,omitempty"`
	ShortVideoDiggEventInfo          *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo          `json:"short_video_digg_event_info,omitempty" xml:"short_video_digg_event_info,omitempty"`
	ShortVideoFinishPlayingEventInfo *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo `json:"short_video_finish_playing_event_info,omitempty" xml:"short_video_finish_playing_event_info,omitempty"`
	CommonEventInfo                  *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo                  `json:"common_event_info,omitempty" xml:"common_event_info,omitempty"`
	ShortVideoShareEventInfo         *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo         `json:"short_video_share_event_info,omitempty" xml:"short_video_share_event_info,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetAccountFollowEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.AccountFollowEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetLiveCommentEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.LiveCommentEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCollectionEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCollectionEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetLiveShareEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.LiveShareEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetLiveWatchEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.LiveWatchEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCommentEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCommentEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetLiveDiggEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.LiveDiggEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetPostingVideEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.PostingVideEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetShortVideoDiggEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoDiggEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetShortVideoFinishPlayingEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoFinishPlayingEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetCommonEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.CommonEventInfo = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo) SetShortVideoShareEventInfo(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoShareEventInfo = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) SetAwemeId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo struct {
	TaskDependence     *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence           `json:"task_dependence,omitempty" xml:"task_dependence,omitempty" require:"true"`
	TaskTypeEnum       *int                                                                                                  `json:"task_type_enum,omitempty" xml:"task_type_enum,omitempty" require:"true"`
	CompleteConditions []*ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem `json:"complete_conditions,omitempty" xml:"complete_conditions,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskDependence(v *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskDependence = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskTypeEnum(v int) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskTypeEnum = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetCompleteConditions(v []*ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.CompleteConditions = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem struct {
	ConditionType              *int     `json:"condition_type,omitempty" xml:"condition_type,omitempty" require:"true"`
	TaskCompleteConditionStage []*int64 `json:"task_complete_condition_stage,omitempty" xml:"task_complete_condition_stage,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetConditionType(v int) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.ConditionType = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetTaskCompleteConditionStage(v []*int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.TaskCompleteConditionStage = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence struct {
	DependenceList []*ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem `json:"dependence_list,omitempty" xml:"dependence_list,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) SetDependenceList(v []*ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence {
	s.DependenceList = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem struct {
	DependenceValue    *string `json:"dependence_value,omitempty" xml:"dependence_value,omitempty" require:"true"`
	DependenceTypeEnum *int    `json:"dependence_type_enum,omitempty" xml:"dependence_type_enum,omitempty" require:"true"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceValue(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceValue = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceTypeEnum(v int) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceTypeEnum = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetAwemeId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.GameId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo struct {
	AwemeId   *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetAwemeId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetDiggCount(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.DiggCount = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.GameId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetAwemeId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.GameId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo struct {
	AwemeId       *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId        *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	WatchDuration *int64  `json:"watch_duration,omitempty" xml:"watch_duration,omitempty" require:"true"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetAwemeId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetWatchDuration(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.WatchDuration = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo struct {
	DiggStage      []*int64 `json:"digg_stage,omitempty" xml:"digg_stage,omitempty" type:"Repeated"`
	NumStage       []*int64 `json:"num_stage,omitempty" xml:"num_stage,omitempty" type:"Repeated"`
	PlayVideoStage []*int64 `json:"play_video_stage,omitempty" xml:"play_video_stage,omitempty" type:"Repeated"`
	CommentStage   []*int64 `json:"comment_stage,omitempty" xml:"comment_stage,omitempty" type:"Repeated"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetDiggStage(v []*int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.DiggStage = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetNumStage(v []*int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.NumStage = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetPlayVideoStage(v []*int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.PlayVideoStage = v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetCommentStage(v []*int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.CommentStage = v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo struct {
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetItemId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetVideoUrl(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo struct {
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetItemId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetVideoUrl(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetAnchorId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo struct {
	AnchorId  *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId    *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl  *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetAnchorId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetDiggCount(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.DiggCount = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetItemId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetVideoUrl(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo struct {
	FinishPlayVidoeCount *int64  `json:"finish_play_vidoe_count,omitempty" xml:"finish_play_vidoe_count,omitempty"`
	GameId               *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId               *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl             *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId             *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetFinishPlayVidoeCount(v int64) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.FinishPlayVidoeCount = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetItemId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetVideoUrl(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetAnchorId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo struct {
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetVideoUrl(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetAnchorId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetGameId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetItemId(v string) *ActivityModifyRequestAddBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.ItemId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItem struct {
	TaskName      *string                                                           `json:"task_name,omitempty" xml:"task_name,omitempty"`
	EndTime       *int64                                                            `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime     *int64                                                            `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TaskEventInfo *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo `json:"task_event_info,omitempty" xml:"task_event_info,omitempty"`
	TaskId        *int64                                                            `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItem) SetTaskName(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItem {
	s.TaskName = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItem) SetEndTime(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItem {
	s.EndTime = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItem) SetStartTime(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItem {
	s.StartTime = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItem) SetTaskEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItem {
	s.TaskEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItem) SetTaskId(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItem {
	s.TaskId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo struct {
	ShortVideoShareEventInfo         *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo         `json:"short_video_share_event_info,omitempty" xml:"short_video_share_event_info,omitempty"`
	LiveWatchEventInfo               *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo               `json:"live_watch_event_info,omitempty" xml:"live_watch_event_info,omitempty"`
	AccountFollowEventInfo           *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo           `json:"account_follow_event_info,omitempty" xml:"account_follow_event_info,omitempty"`
	PostingVideEventInfo             *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo             `json:"posting_vide_event_info,omitempty" xml:"posting_vide_event_info,omitempty"`
	LiveDiggEventInfo                *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo                `json:"live_digg_event_info,omitempty" xml:"live_digg_event_info,omitempty"`
	ShortVideoCommentEventInfo       *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo       `json:"short_video_comment_event_info,omitempty" xml:"short_video_comment_event_info,omitempty"`
	ShortVideoDiggEventInfo          *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo          `json:"short_video_digg_event_info,omitempty" xml:"short_video_digg_event_info,omitempty"`
	LiveCommentEventInfo             *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo             `json:"live_comment_event_info,omitempty" xml:"live_comment_event_info,omitempty"`
	ShortVideoFinishPlayingEventInfo *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo `json:"short_video_finish_playing_event_info,omitempty" xml:"short_video_finish_playing_event_info,omitempty"`
	ShortVideoCollectionEventInfo    *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo    `json:"short_video_collection_event_info,omitempty" xml:"short_video_collection_event_info,omitempty"`
	CommonEventInfo                  *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo                  `json:"common_event_info,omitempty" xml:"common_event_info,omitempty"`
	LiveShareEventInfo               *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo               `json:"live_share_event_info,omitempty" xml:"live_share_event_info,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetShortVideoShareEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoShareEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetLiveWatchEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.LiveWatchEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetAccountFollowEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.AccountFollowEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetPostingVideEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.PostingVideEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetLiveDiggEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.LiveDiggEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCommentEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCommentEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetShortVideoDiggEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoDiggEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetLiveCommentEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.LiveCommentEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetShortVideoFinishPlayingEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoFinishPlayingEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCollectionEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCollectionEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetCommonEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.CommonEventInfo = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo) SetLiveShareEventInfo(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfo {
	s.LiveShareEventInfo = v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) SetAwemeId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo struct {
	TaskTypeEnum       *int                                                                                                     `json:"task_type_enum,omitempty" xml:"task_type_enum,omitempty" require:"true"`
	CompleteConditions []*ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem `json:"complete_conditions,omitempty" xml:"complete_conditions,omitempty" require:"true" type:"Repeated"`
	TaskDependence     *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence           `json:"task_dependence,omitempty" xml:"task_dependence,omitempty" require:"true"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskTypeEnum(v int) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskTypeEnum = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetCompleteConditions(v []*ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.CompleteConditions = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskDependence(v *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskDependence = v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem struct {
	TaskCompleteConditionStage []*int64 `json:"task_complete_condition_stage,omitempty" xml:"task_complete_condition_stage,omitempty" require:"true" type:"Repeated"`
	ConditionType              *int     `json:"condition_type,omitempty" xml:"condition_type,omitempty" require:"true"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetTaskCompleteConditionStage(v []*int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.TaskCompleteConditionStage = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetConditionType(v int) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.ConditionType = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence struct {
	DependenceList []*ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem `json:"dependence_list,omitempty" xml:"dependence_list,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) SetDependenceList(v []*ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence {
	s.DependenceList = v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem struct {
	DependenceTypeEnum *int    `json:"dependence_type_enum,omitempty" xml:"dependence_type_enum,omitempty" require:"true"`
	DependenceValue    *string `json:"dependence_value,omitempty" xml:"dependence_value,omitempty" require:"true"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceTypeEnum(v int) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceTypeEnum = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceValue(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceValue = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo struct {
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetAwemeId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo struct {
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	AwemeId   *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetAwemeId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetDiggCount(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.DiggCount = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo struct {
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetAwemeId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo struct {
	AwemeId       *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId        *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	WatchDuration *int64  `json:"watch_duration,omitempty" xml:"watch_duration,omitempty" require:"true"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetAwemeId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetWatchDuration(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.WatchDuration = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo struct {
	CommentStage   []*int64 `json:"comment_stage,omitempty" xml:"comment_stage,omitempty" type:"Repeated"`
	DiggStage      []*int64 `json:"digg_stage,omitempty" xml:"digg_stage,omitempty" type:"Repeated"`
	NumStage       []*int64 `json:"num_stage,omitempty" xml:"num_stage,omitempty" type:"Repeated"`
	PlayVideoStage []*int64 `json:"play_video_stage,omitempty" xml:"play_video_stage,omitempty" type:"Repeated"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetCommentStage(v []*int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.CommentStage = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetDiggStage(v []*int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.DiggStage = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetNumStage(v []*int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.NumStage = v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetPlayVideoStage(v []*int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.PlayVideoStage = v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo struct {
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetItemId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetVideoUrl(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo struct {
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetItemId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetVideoUrl(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetAnchorId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo struct {
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId    *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl  *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId  *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetItemId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetVideoUrl(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetAnchorId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetDiggCount(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.DiggCount = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo struct {
	VideoUrl             *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId             *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	FinishPlayVidoeCount *int64  `json:"finish_play_vidoe_count,omitempty" xml:"finish_play_vidoe_count,omitempty"`
	GameId               *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId               *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetVideoUrl(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetAnchorId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetFinishPlayVidoeCount(v int64) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.FinishPlayVidoeCount = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetItemId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.ItemId = &v
	return s
}

type ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo struct {
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetGameId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetItemId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetVideoUrl(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetAnchorId(v string) *ActivityModifyRequestModifyBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityModifyResponse struct {
	ModifyBusinessTaskIdList []*int64 `json:"modify_business_task_id_list,omitempty" xml:"modify_business_task_id_list,omitempty" type:"Repeated"`
	ActivityId               *int64   `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	AddBusinessTaskIdList    []*int64 `json:"add_business_task_id_list,omitempty" xml:"add_business_task_id_list,omitempty" type:"Repeated"`
	DelBusinessTaskIdList    []*int64 `json:"del_business_task_id_list,omitempty" xml:"del_business_task_id_list,omitempty" type:"Repeated"`
}

func (s ActivityModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityModifyResponse) GoString() string {
	return s.String()
}

func (s *ActivityModifyResponse) SetModifyBusinessTaskIdList(v []*int64) *ActivityModifyResponse {
	s.ModifyBusinessTaskIdList = v
	return s
}

func (s *ActivityModifyResponse) SetActivityId(v int64) *ActivityModifyResponse {
	s.ActivityId = &v
	return s
}

func (s *ActivityModifyResponse) SetAddBusinessTaskIdList(v []*int64) *ActivityModifyResponse {
	s.AddBusinessTaskIdList = v
	return s
}

func (s *ActivityModifyResponse) SetDelBusinessTaskIdList(v []*int64) *ActivityModifyResponse {
	s.DelBusinessTaskIdList = v
	return s
}

type ActivityQueryBindedUserRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	ActivityId  *string            `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ActivityQueryBindedUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryBindedUserRequest) GoString() string {
	return s.String()
}

func (s *ActivityQueryBindedUserRequest) SetHeader(v map[string]*string) *ActivityQueryBindedUserRequest {
	s.Header = v
	return s
}

func (s *ActivityQueryBindedUserRequest) SetAccessToken(v string) *ActivityQueryBindedUserRequest {
	s.AccessToken = &v
	return s
}

func (s *ActivityQueryBindedUserRequest) SetOpenId(v string) *ActivityQueryBindedUserRequest {
	s.OpenId = &v
	return s
}

func (s *ActivityQueryBindedUserRequest) SetActivityId(v string) *ActivityQueryBindedUserRequest {
	s.ActivityId = &v
	return s
}

type ActivityQueryBindedUserResponse struct {
	Data   *ActivityQueryBindedUserResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s ActivityQueryBindedUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryBindedUserResponse) GoString() string {
	return s.String()
}

func (s *ActivityQueryBindedUserResponse) SetData(v *ActivityQueryBindedUserResponseData) *ActivityQueryBindedUserResponse {
	s.Data = v
	return s
}

func (s *ActivityQueryBindedUserResponse) SetErrNo(v int32) *ActivityQueryBindedUserResponse {
	s.ErrNo = &v
	return s
}

func (s *ActivityQueryBindedUserResponse) SetErrMsg(v string) *ActivityQueryBindedUserResponse {
	s.ErrMsg = &v
	return s
}

func (s *ActivityQueryBindedUserResponse) SetLogId(v string) *ActivityQueryBindedUserResponse {
	s.LogId = &v
	return s
}

type ActivityQueryBindedUserResponseData struct {
	IsBinded *bool `json:"is_binded,omitempty" xml:"is_binded,omitempty" require:"true"`
}

func (s ActivityQueryBindedUserResponseData) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryBindedUserResponseData) GoString() string {
	return s.String()
}

func (s *ActivityQueryBindedUserResponseData) SetIsBinded(v bool) *ActivityQueryBindedUserResponseData {
	s.IsBinded = &v
	return s
}

type ActivityQueryInfoRequest struct {
	TaskIdList  []*int64           `json:"task_id_list,omitempty" xml:"task_id_list,omitempty" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	ActivityId  *int64             `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ActivityQueryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoRequest) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoRequest) SetTaskIdList(v []*int64) *ActivityQueryInfoRequest {
	s.TaskIdList = v
	return s
}

func (s *ActivityQueryInfoRequest) SetHeader(v map[string]*string) *ActivityQueryInfoRequest {
	s.Header = v
	return s
}

func (s *ActivityQueryInfoRequest) SetAccessToken(v string) *ActivityQueryInfoRequest {
	s.AccessToken = &v
	return s
}

func (s *ActivityQueryInfoRequest) SetActivityId(v int64) *ActivityQueryInfoRequest {
	s.ActivityId = &v
	return s
}

type ActivityQueryInfoResponse struct {
	ActivityInfo         *ActivityQueryInfoResponseActivityInfo               `json:"ActivityInfo,omitempty" xml:"ActivityInfo,omitempty" require:"true"`
	BusinessTaskInfoList []*ActivityQueryInfoResponseBusinessTaskInfoListItem `json:"business_task_info_list,omitempty" xml:"business_task_info_list,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityQueryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponse) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponse) SetActivityInfo(v *ActivityQueryInfoResponseActivityInfo) *ActivityQueryInfoResponse {
	s.ActivityInfo = v
	return s
}

func (s *ActivityQueryInfoResponse) SetBusinessTaskInfoList(v []*ActivityQueryInfoResponseBusinessTaskInfoListItem) *ActivityQueryInfoResponse {
	s.BusinessTaskInfoList = v
	return s
}

type ActivityQueryInfoResponseActivityInfo struct {
	ActivityName *string `json:"activity_name,omitempty" xml:"activity_name,omitempty" require:"true"`
	EndTime      *int64  `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	StartTime    *int64  `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s ActivityQueryInfoResponseActivityInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseActivityInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseActivityInfo) SetActivityName(v string) *ActivityQueryInfoResponseActivityInfo {
	s.ActivityName = &v
	return s
}

func (s *ActivityQueryInfoResponseActivityInfo) SetEndTime(v int64) *ActivityQueryInfoResponseActivityInfo {
	s.EndTime = &v
	return s
}

func (s *ActivityQueryInfoResponseActivityInfo) SetStartTime(v int64) *ActivityQueryInfoResponseActivityInfo {
	s.StartTime = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItem struct {
	EndTime       *int64                                                          `json:"end_time,omitempty" xml:"end_time,omitempty"`
	StartTime     *int64                                                          `json:"start_time,omitempty" xml:"start_time,omitempty"`
	TaskEventInfo *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo `json:"task_event_info,omitempty" xml:"task_event_info,omitempty"`
	TaskId        *int64                                                          `json:"task_id,omitempty" xml:"task_id,omitempty"`
	TaskName      *string                                                         `json:"task_name,omitempty" xml:"task_name,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItem) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItem) SetEndTime(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItem {
	s.EndTime = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItem) SetStartTime(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItem {
	s.StartTime = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItem) SetTaskEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItem {
	s.TaskEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItem) SetTaskId(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItem {
	s.TaskId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItem) SetTaskName(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItem {
	s.TaskName = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo struct {
	ShortVideoCommentEventInfo       *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo       `json:"short_video_comment_event_info,omitempty" xml:"short_video_comment_event_info,omitempty"`
	LiveShareEventInfo               *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo               `json:"live_share_event_info,omitempty" xml:"live_share_event_info,omitempty"`
	PostingVideEventInfo             *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo             `json:"posting_vide_event_info,omitempty" xml:"posting_vide_event_info,omitempty"`
	LiveDiggEventInfo                *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo                `json:"live_digg_event_info,omitempty" xml:"live_digg_event_info,omitempty"`
	LiveWatchEventInfo               *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo               `json:"live_watch_event_info,omitempty" xml:"live_watch_event_info,omitempty"`
	AccountFollowEventInfo           *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo           `json:"account_follow_event_info,omitempty" xml:"account_follow_event_info,omitempty"`
	ShortVideoShareEventInfo         *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo         `json:"short_video_share_event_info,omitempty" xml:"short_video_share_event_info,omitempty"`
	ShortVideoDiggEventInfo          *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo          `json:"short_video_digg_event_info,omitempty" xml:"short_video_digg_event_info,omitempty"`
	CommonEventInfo                  *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo                  `json:"common_event_info,omitempty" xml:"common_event_info,omitempty"`
	ShortVideoCollectionEventInfo    *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo    `json:"short_video_collection_event_info,omitempty" xml:"short_video_collection_event_info,omitempty"`
	ShortVideoFinishPlayingEventInfo *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo `json:"short_video_finish_playing_event_info,omitempty" xml:"short_video_finish_playing_event_info,omitempty"`
	LiveCommentEventInfo             *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo             `json:"live_comment_event_info,omitempty" xml:"live_comment_event_info,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCommentEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCommentEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetLiveShareEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.LiveShareEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetPostingVideEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.PostingVideEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetLiveDiggEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.LiveDiggEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetLiveWatchEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.LiveWatchEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetAccountFollowEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.AccountFollowEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetShortVideoShareEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoShareEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetShortVideoDiggEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoDiggEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetCommonEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.CommonEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetShortVideoCollectionEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoCollectionEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetShortVideoFinishPlayingEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.ShortVideoFinishPlayingEventInfo = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo) SetLiveCommentEventInfo(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfo {
	s.LiveCommentEventInfo = v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo) SetAwemeId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoAccountFollowEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo struct {
	TaskTypeEnum       *int                                                                                                   `json:"task_type_enum,omitempty" xml:"task_type_enum,omitempty" require:"true"`
	CompleteConditions []*ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem `json:"complete_conditions,omitempty" xml:"complete_conditions,omitempty" require:"true" type:"Repeated"`
	TaskDependence     *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence           `json:"task_dependence,omitempty" xml:"task_dependence,omitempty" require:"true"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskTypeEnum(v int) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskTypeEnum = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetCompleteConditions(v []*ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.CompleteConditions = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo) SetTaskDependence(v *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfo {
	s.TaskDependence = v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem struct {
	TaskCompleteConditionStage []*int64 `json:"task_complete_condition_stage,omitempty" xml:"task_complete_condition_stage,omitempty" require:"true" type:"Repeated"`
	ConditionType              *int     `json:"condition_type,omitempty" xml:"condition_type,omitempty" require:"true"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetTaskCompleteConditionStage(v []*int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.TaskCompleteConditionStage = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem) SetConditionType(v int) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoCompleteConditionsItem {
	s.ConditionType = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence struct {
	DependenceList []*ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem `json:"dependence_list,omitempty" xml:"dependence_list,omitempty" require:"true" type:"Repeated"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence) SetDependenceList(v []*ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependence {
	s.DependenceList = v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem struct {
	DependenceTypeEnum *int    `json:"dependence_type_enum,omitempty" xml:"dependence_type_enum,omitempty" require:"true"`
	DependenceValue    *string `json:"dependence_value,omitempty" xml:"dependence_value,omitempty" require:"true"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceTypeEnum(v int) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceTypeEnum = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem) SetDependenceValue(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoCommonEventInfoTaskDependenceDependenceListItem {
	s.DependenceValue = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetAwemeId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveCommentEventInfo {
	s.GameId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo struct {
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	AwemeId   *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetDiggCount(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.DiggCount = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo) SetAwemeId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveDiggEventInfo {
	s.AwemeId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo struct {
	AwemeId *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId  *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetAwemeId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveShareEventInfo {
	s.GameId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo struct {
	WatchDuration *int64  `json:"watch_duration,omitempty" xml:"watch_duration,omitempty" require:"true"`
	AwemeId       *string `json:"aweme_id,omitempty" xml:"aweme_id,omitempty"`
	GameId        *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetWatchDuration(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.WatchDuration = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetAwemeId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.AwemeId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoLiveWatchEventInfo {
	s.GameId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo struct {
	DiggStage      []*int64 `json:"digg_stage,omitempty" xml:"digg_stage,omitempty" type:"Repeated"`
	NumStage       []*int64 `json:"num_stage,omitempty" xml:"num_stage,omitempty" type:"Repeated"`
	PlayVideoStage []*int64 `json:"play_video_stage,omitempty" xml:"play_video_stage,omitempty" type:"Repeated"`
	CommentStage   []*int64 `json:"comment_stage,omitempty" xml:"comment_stage,omitempty" type:"Repeated"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetDiggStage(v []*int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.DiggStage = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetNumStage(v []*int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.NumStage = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetPlayVideoStage(v []*int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.PlayVideoStage = v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo) SetCommentStage(v []*int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoPostingVideEventInfo {
	s.CommentStage = v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo struct {
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetVideoUrl(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo) SetItemId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCollectionEventInfo {
	s.ItemId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo struct {
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetItemId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetVideoUrl(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetAnchorId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoCommentEventInfo {
	s.GameId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo struct {
	AnchorId  *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	DiggCount *int64  `json:"digg_count,omitempty" xml:"digg_count,omitempty"`
	GameId    *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId    *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl  *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetAnchorId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetDiggCount(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.DiggCount = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetItemId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo) SetVideoUrl(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoDiggEventInfo {
	s.VideoUrl = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo struct {
	FinishPlayVidoeCount *int64  `json:"finish_play_vidoe_count,omitempty" xml:"finish_play_vidoe_count,omitempty"`
	GameId               *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId               *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	VideoUrl             *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId             *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetFinishPlayVidoeCount(v int64) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.FinishPlayVidoeCount = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetItemId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.ItemId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetVideoUrl(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo) SetAnchorId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoFinishPlayingEventInfo {
	s.AnchorId = &v
	return s
}

type ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo struct {
	VideoUrl *string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	AnchorId *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	GameId   *string `json:"game_id,omitempty" xml:"game_id,omitempty"`
	ItemId   *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) GoString() string {
	return s.String()
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetVideoUrl(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.VideoUrl = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetAnchorId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.AnchorId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetGameId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.GameId = &v
	return s
}

func (s *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo) SetItemId(v string) *ActivityQueryInfoResponseBusinessTaskInfoListItemTaskEventInfoShortVideoShareEventInfo {
	s.ItemId = &v
	return s
}

type ActivityQueryPromotionActivityRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BizType     *int               `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	ActivityId  *string            `json:"activity_id,omitempty" xml:"activity_id,omitempty" require:"true"`
}

func (s ActivityQueryPromotionActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityRequest) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityRequest) SetHeader(v map[string]*string) *ActivityQueryPromotionActivityRequest {
	s.Header = v
	return s
}

func (s *ActivityQueryPromotionActivityRequest) SetAccessToken(v string) *ActivityQueryPromotionActivityRequest {
	s.AccessToken = &v
	return s
}

func (s *ActivityQueryPromotionActivityRequest) SetBizType(v int) *ActivityQueryPromotionActivityRequest {
	s.BizType = &v
	return s
}

func (s *ActivityQueryPromotionActivityRequest) SetActivityId(v string) *ActivityQueryPromotionActivityRequest {
	s.ActivityId = &v
	return s
}

type ActivityQueryPromotionActivityResponse struct {
	ErrNo             *int32                                                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg            *string                                                  `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId             *string                                                  `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	PromotionActivity *ActivityQueryPromotionActivityResponsePromotionActivity `json:"promotion_activity,omitempty" xml:"promotion_activity,omitempty" require:"true"`
}

func (s ActivityQueryPromotionActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponse) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponse) SetErrNo(v int32) *ActivityQueryPromotionActivityResponse {
	s.ErrNo = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponse) SetErrMsg(v string) *ActivityQueryPromotionActivityResponse {
	s.ErrMsg = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponse) SetLogId(v string) *ActivityQueryPromotionActivityResponse {
	s.LogId = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponse) SetPromotionActivity(v *ActivityQueryPromotionActivityResponsePromotionActivity) *ActivityQueryPromotionActivityResponse {
	s.PromotionActivity = v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivity struct {
	ImActivity      *ActivityQueryPromotionActivityResponsePromotionActivityImActivity      `json:"im_activity,omitempty" xml:"im_activity,omitempty"`
	LiveActivity    *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity    `json:"live_activity,omitempty" xml:"live_activity,omitempty"`
	ActivityName    *string                                                                 `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	ReceiveLimit    *int64                                                                  `json:"receive_limit,omitempty" xml:"receive_limit,omitempty"`
	SendScene       *int                                                                    `json:"send_scene,omitempty" xml:"send_scene,omitempty"`
	SidebarActivity *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity `json:"sidebar_activity,omitempty" xml:"sidebar_activity,omitempty"`
	CouponMetaId    *string                                                                 `json:"coupon_meta_id,omitempty" xml:"coupon_meta_id,omitempty"`
	ActivityId      *string                                                                 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	Status          *int                                                                    `json:"status,omitempty" xml:"status,omitempty"`
	FeedActivity    *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity    `json:"feed_activity,omitempty" xml:"feed_activity,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivity) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetImActivity(v *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.ImActivity = v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetLiveActivity(v *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.LiveActivity = v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetActivityName(v string) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.ActivityName = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetReceiveLimit(v int64) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.ReceiveLimit = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetSendScene(v int) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.SendScene = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetSidebarActivity(v *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.SidebarActivity = v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetCouponMetaId(v string) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.CouponMetaId = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetActivityId(v string) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.ActivityId = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetStatus(v int) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.Status = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivity) SetFeedActivity(v *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) *ActivityQueryPromotionActivityResponsePromotionActivity {
	s.FeedActivity = v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity struct {
	CanReceiveNum         *int64 `json:"can_receive_num,omitempty" xml:"can_receive_num,omitempty"`
	ValidBeginTime        *int64 `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	ValidEndTime          *int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	IsRecommendedStrategy *bool  `json:"is_recommended_strategy,omitempty" xml:"is_recommended_strategy,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) SetCanReceiveNum(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity {
	s.CanReceiveNum = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) SetValidBeginTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) SetValidEndTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity) SetIsRecommendedStrategy(v bool) *ActivityQueryPromotionActivityResponsePromotionActivityFeedActivity {
	s.IsRecommendedStrategy = &v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivityImActivity struct {
	NoticeText         *string                                                                              `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
	ValidBeginTime     *int64                                                                               `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	ValidEndTime       *int64                                                                               `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	ShareFissionConfig *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig `json:"share_fission_config,omitempty" xml:"share_fission_config,omitempty"`
	CouponTaskType     *int                                                                                 `json:"coupon_task_type,omitempty" xml:"coupon_task_type,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityImActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityImActivity) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) SetNoticeText(v string) *ActivityQueryPromotionActivityResponsePromotionActivityImActivity {
	s.NoticeText = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) SetValidBeginTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityImActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) SetValidEndTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityImActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) SetShareFissionConfig(v *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) *ActivityQueryPromotionActivityResponsePromotionActivityImActivity {
	s.ShareFissionConfig = v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivity) SetCouponTaskType(v int) *ActivityQueryPromotionActivityResponsePromotionActivityImActivity {
	s.CouponTaskType = &v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig struct {
	AwardCouponNum  *int32  `json:"award_coupon_num,omitempty" xml:"award_coupon_num,omitempty"`
	AwardType       *int    `json:"award_type,omitempty" xml:"award_type,omitempty" require:"true"`
	AwardCouponId   *int64  `json:"award_coupon_id,omitempty" xml:"award_coupon_id,omitempty"`
	AwardCouponName *string `json:"award_coupon_name,omitempty" xml:"award_coupon_name,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) SetAwardCouponNum(v int32) *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig {
	s.AwardCouponNum = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) SetAwardType(v int) *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig {
	s.AwardType = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) SetAwardCouponId(v int64) *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig {
	s.AwardCouponId = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig) SetAwardCouponName(v string) *ActivityQueryPromotionActivityResponsePromotionActivityImActivityShareFissionConfig {
	s.AwardCouponName = &v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity struct {
	CouponTaskType  *int    `json:"coupon_task_type,omitempty" xml:"coupon_task_type,omitempty"`
	ClueComponentId *string `json:"clue_component_id,omitempty" xml:"clue_component_id,omitempty"`
	NoticeText      *string `json:"notice_text,omitempty" xml:"notice_text,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) SetCouponTaskType(v int) *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity {
	s.CouponTaskType = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) SetClueComponentId(v string) *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity {
	s.ClueComponentId = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity) SetNoticeText(v string) *ActivityQueryPromotionActivityResponsePromotionActivityLiveActivity {
	s.NoticeText = &v
	return s
}

type ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity struct {
	JumpText               *string `json:"jump_text,omitempty" xml:"jump_text,omitempty"`
	ValidBeginTime         *int64  `json:"valid_begin_time,omitempty" xml:"valid_begin_time,omitempty"`
	CompleteTimes          *int32  `json:"complete_times,omitempty" xml:"complete_times,omitempty"`
	IsCustomEducationPopup *bool   `json:"is_custom_education_popup,omitempty" xml:"is_custom_education_popup,omitempty"`
	ValidEndTime           *int64  `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	ShortTitle             *string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	HighValueContent       *string `json:"high_value_content,omitempty" xml:"high_value_content,omitempty"`
	NeedBind               *bool   `json:"need_bind,omitempty" xml:"need_bind,omitempty"`
	ActionTrigger          *int    `json:"action_trigger,omitempty" xml:"action_trigger,omitempty"`
	IsCustomAwardPopup     *bool   `json:"is_custom_award_popup,omitempty" xml:"is_custom_award_popup,omitempty"`
	RecentBubbleText       *string `json:"recent_bubble_text,omitempty" xml:"recent_bubble_text,omitempty"`
}

func (s ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) String() string {
	return tea.Prettify(s)
}

func (s ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) GoString() string {
	return s.String()
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetJumpText(v string) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.JumpText = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetValidBeginTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.ValidBeginTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetCompleteTimes(v int32) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.CompleteTimes = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetIsCustomEducationPopup(v bool) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.IsCustomEducationPopup = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetValidEndTime(v int64) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.ValidEndTime = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetShortTitle(v string) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.ShortTitle = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetHighValueContent(v string) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.HighValueContent = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetNeedBind(v bool) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.NeedBind = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetActionTrigger(v int) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.ActionTrigger = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetIsCustomAwardPopup(v bool) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.IsCustomAwardPopup = &v
	return s
}

func (s *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity) SetRecentBubbleText(v string) *ActivityQueryPromotionActivityResponsePromotionActivitySidebarActivity {
	s.RecentBubbleText = &v
	return s
}

type AdPlacementAddRequest struct {
	AdPlacementType *int64             `json:"ad_placement_type,omitempty" xml:"ad_placement_type,omitempty" require:"true"`
	AdPlacementName *string            `json:"ad_placement_name,omitempty" xml:"ad_placement_name,omitempty" require:"true"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AdPlacementAddRequest) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementAddRequest) GoString() string {
	return s.String()
}

func (s *AdPlacementAddRequest) SetAdPlacementType(v int64) *AdPlacementAddRequest {
	s.AdPlacementType = &v
	return s
}

func (s *AdPlacementAddRequest) SetAdPlacementName(v string) *AdPlacementAddRequest {
	s.AdPlacementName = &v
	return s
}

func (s *AdPlacementAddRequest) SetHeader(v map[string]*string) *AdPlacementAddRequest {
	s.Header = v
	return s
}

func (s *AdPlacementAddRequest) SetAccessToken(v string) *AdPlacementAddRequest {
	s.AccessToken = &v
	return s
}

type AdPlacementAddResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s AdPlacementAddResponse) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementAddResponse) GoString() string {
	return s.String()
}

func (s *AdPlacementAddResponse) SetErrMsg(v string) *AdPlacementAddResponse {
	s.ErrMsg = &v
	return s
}

func (s *AdPlacementAddResponse) SetLogId(v string) *AdPlacementAddResponse {
	s.LogId = &v
	return s
}

func (s *AdPlacementAddResponse) SetErrNo(v int32) *AdPlacementAddResponse {
	s.ErrNo = &v
	return s
}

type AdPlacementQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AdPlacementQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementQueryRequest) GoString() string {
	return s.String()
}

func (s *AdPlacementQueryRequest) SetHeader(v map[string]*string) *AdPlacementQueryRequest {
	s.Header = v
	return s
}

func (s *AdPlacementQueryRequest) SetAccessToken(v string) *AdPlacementQueryRequest {
	s.AccessToken = &v
	return s
}

type AdPlacementQueryResponse struct {
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AdPlacementQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s AdPlacementQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementQueryResponse) GoString() string {
	return s.String()
}

func (s *AdPlacementQueryResponse) SetLogId(v string) *AdPlacementQueryResponse {
	s.LogId = &v
	return s
}

func (s *AdPlacementQueryResponse) SetData(v *AdPlacementQueryResponseData) *AdPlacementQueryResponse {
	s.Data = v
	return s
}

func (s *AdPlacementQueryResponse) SetErrNo(v int32) *AdPlacementQueryResponse {
	s.ErrNo = &v
	return s
}

func (s *AdPlacementQueryResponse) SetErrMsg(v string) *AdPlacementQueryResponse {
	s.ErrMsg = &v
	return s
}

type AdPlacementQueryResponseData struct {
	AdPlacementList []*AdPlacementQueryResponseDataAdPlacementListItem `json:"ad_placement_list,omitempty" xml:"ad_placement_list,omitempty" type:"Repeated"`
}

func (s AdPlacementQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementQueryResponseData) GoString() string {
	return s.String()
}

func (s *AdPlacementQueryResponseData) SetAdPlacementList(v []*AdPlacementQueryResponseDataAdPlacementListItem) *AdPlacementQueryResponseData {
	s.AdPlacementList = v
	return s
}

type AdPlacementQueryResponseDataAdPlacementListItem struct {
	AdPlacementType *int64  `json:"ad_placement_type,omitempty" xml:"ad_placement_type,omitempty"`
	Status          *int64  `json:"status,omitempty" xml:"status,omitempty"`
	AdPlacementId   *string `json:"ad_placement_id,omitempty" xml:"ad_placement_id,omitempty"`
	AdPlacementName *string `json:"ad_placement_name,omitempty" xml:"ad_placement_name,omitempty"`
}

func (s AdPlacementQueryResponseDataAdPlacementListItem) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementQueryResponseDataAdPlacementListItem) GoString() string {
	return s.String()
}

func (s *AdPlacementQueryResponseDataAdPlacementListItem) SetAdPlacementType(v int64) *AdPlacementQueryResponseDataAdPlacementListItem {
	s.AdPlacementType = &v
	return s
}

func (s *AdPlacementQueryResponseDataAdPlacementListItem) SetStatus(v int64) *AdPlacementQueryResponseDataAdPlacementListItem {
	s.Status = &v
	return s
}

func (s *AdPlacementQueryResponseDataAdPlacementListItem) SetAdPlacementId(v string) *AdPlacementQueryResponseDataAdPlacementListItem {
	s.AdPlacementId = &v
	return s
}

func (s *AdPlacementQueryResponseDataAdPlacementListItem) SetAdPlacementName(v string) *AdPlacementQueryResponseDataAdPlacementListItem {
	s.AdPlacementName = &v
	return s
}

type AdPlacementUpdateRequest struct {
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AdPlacementId *string            `json:"ad_placement_id,omitempty" xml:"ad_placement_id,omitempty" require:"true"`
	Status        *int64             `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AdPlacementUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementUpdateRequest) GoString() string {
	return s.String()
}

func (s *AdPlacementUpdateRequest) SetAccessToken(v string) *AdPlacementUpdateRequest {
	s.AccessToken = &v
	return s
}

func (s *AdPlacementUpdateRequest) SetAdPlacementId(v string) *AdPlacementUpdateRequest {
	s.AdPlacementId = &v
	return s
}

func (s *AdPlacementUpdateRequest) SetStatus(v int64) *AdPlacementUpdateRequest {
	s.Status = &v
	return s
}

func (s *AdPlacementUpdateRequest) SetHeader(v map[string]*string) *AdPlacementUpdateRequest {
	s.Header = v
	return s
}

type AdPlacementUpdateResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

func (s AdPlacementUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s AdPlacementUpdateResponse) GoString() string {
	return s.String()
}

func (s *AdPlacementUpdateResponse) SetLogId(v string) *AdPlacementUpdateResponse {
	s.LogId = &v
	return s
}

func (s *AdPlacementUpdateResponse) SetErrNo(v int32) *AdPlacementUpdateResponse {
	s.ErrNo = &v
	return s
}

func (s *AdPlacementUpdateResponse) SetErrMsg(v string) *AdPlacementUpdateResponse {
	s.ErrMsg = &v
	return s
}

type AddAppTestRelationRequest struct {
	Type        *string            `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RefIdList   []*string          `json:"ref_id_list,omitempty" xml:"ref_id_list,omitempty" require:"true" type:"Repeated"`
	Operator    *string            `json:"operator,omitempty" xml:"operator,omitempty" require:"true"`
}

func (s AddAppTestRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppTestRelationRequest) GoString() string {
	return s.String()
}

func (s *AddAppTestRelationRequest) SetType(v string) *AddAppTestRelationRequest {
	s.Type = &v
	return s
}

func (s *AddAppTestRelationRequest) SetHeader(v map[string]*string) *AddAppTestRelationRequest {
	s.Header = v
	return s
}

func (s *AddAppTestRelationRequest) SetAccessToken(v string) *AddAppTestRelationRequest {
	s.AccessToken = &v
	return s
}

func (s *AddAppTestRelationRequest) SetRefIdList(v []*string) *AddAppTestRelationRequest {
	s.RefIdList = v
	return s
}

func (s *AddAppTestRelationRequest) SetOperator(v string) *AddAppTestRelationRequest {
	s.Operator = &v
	return s
}

type AddAppTestRelationResponse struct {
	Data  *AddAppTestRelationResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *AddAppTestRelationResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s AddAppTestRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppTestRelationResponse) GoString() string {
	return s.String()
}

func (s *AddAppTestRelationResponse) SetData(v *AddAppTestRelationResponseData) *AddAppTestRelationResponse {
	s.Data = v
	return s
}

func (s *AddAppTestRelationResponse) SetExtra(v *AddAppTestRelationResponseExtra) *AddAppTestRelationResponse {
	s.Extra = v
	return s
}

type AddAppTestRelationResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AddAppTestRelationResponseData) String() string {
	return tea.Prettify(s)
}

func (s AddAppTestRelationResponseData) GoString() string {
	return s.String()
}

func (s *AddAppTestRelationResponseData) SetGwDescription(v string) *AddAppTestRelationResponseData {
	s.GwDescription = &v
	return s
}

func (s *AddAppTestRelationResponseData) SetGwErrorCode(v int32) *AddAppTestRelationResponseData {
	s.GwErrorCode = &v
	return s
}

type AddAppTestRelationResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AddAppTestRelationResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AddAppTestRelationResponseExtra) GoString() string {
	return s.String()
}

func (s *AddAppTestRelationResponseExtra) SetSubErrorCode(v int32) *AddAppTestRelationResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AddAppTestRelationResponseExtra) SetSubDescription(v string) *AddAppTestRelationResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AddAppTestRelationResponseExtra) SetLogid(v string) *AddAppTestRelationResponseExtra {
	s.Logid = &v
	return s
}

func (s *AddAppTestRelationResponseExtra) SetNow(v int64) *AddAppTestRelationResponseExtra {
	s.Now = &v
	return s
}

func (s *AddAppTestRelationResponseExtra) SetErrorCode(v int32) *AddAppTestRelationResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AddAppTestRelationResponseExtra) SetDescription(v string) *AddAppTestRelationResponseExtra {
	s.Description = &v
	return s
}

type AddAwemeVideoKeywordRequest struct {
	Keyword     *string            `json:"keyword,omitempty" xml:"keyword,omitempty" require:"true"`
	Reason      *string            `json:"reason,omitempty" xml:"reason,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AddAwemeVideoKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAwemeVideoKeywordRequest) GoString() string {
	return s.String()
}

func (s *AddAwemeVideoKeywordRequest) SetKeyword(v string) *AddAwemeVideoKeywordRequest {
	s.Keyword = &v
	return s
}

func (s *AddAwemeVideoKeywordRequest) SetReason(v string) *AddAwemeVideoKeywordRequest {
	s.Reason = &v
	return s
}

func (s *AddAwemeVideoKeywordRequest) SetHeader(v map[string]*string) *AddAwemeVideoKeywordRequest {
	s.Header = v
	return s
}

func (s *AddAwemeVideoKeywordRequest) SetAccessToken(v string) *AddAwemeVideoKeywordRequest {
	s.AccessToken = &v
	return s
}

type AddAwemeVideoKeywordResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s AddAwemeVideoKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAwemeVideoKeywordResponse) GoString() string {
	return s.String()
}

func (s *AddAwemeVideoKeywordResponse) SetErrMsg(v string) *AddAwemeVideoKeywordResponse {
	s.ErrMsg = &v
	return s
}

func (s *AddAwemeVideoKeywordResponse) SetLogId(v string) *AddAwemeVideoKeywordResponse {
	s.LogId = &v
	return s
}

func (s *AddAwemeVideoKeywordResponse) SetErrNo(v int32) *AddAwemeVideoKeywordResponse {
	s.ErrNo = &v
	return s
}

type AddSimpleQrBindRequest struct {
	ExclusiveQrUrlPrefix *int32             `json:"exclusive_qr_url_prefix,omitempty" xml:"exclusive_qr_url_prefix,omitempty" require:"true"`
	QrUrl                *string            `json:"qr_url,omitempty" xml:"qr_url,omitempty" require:"true"`
	LoadPath             *string            `json:"load_path,omitempty" xml:"load_path,omitempty" require:"true"`
	Stage                *string            `json:"stage,omitempty" xml:"stage,omitempty" require:"true"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AddSimpleQrBindRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSimpleQrBindRequest) GoString() string {
	return s.String()
}

func (s *AddSimpleQrBindRequest) SetExclusiveQrUrlPrefix(v int32) *AddSimpleQrBindRequest {
	s.ExclusiveQrUrlPrefix = &v
	return s
}

func (s *AddSimpleQrBindRequest) SetQrUrl(v string) *AddSimpleQrBindRequest {
	s.QrUrl = &v
	return s
}

func (s *AddSimpleQrBindRequest) SetLoadPath(v string) *AddSimpleQrBindRequest {
	s.LoadPath = &v
	return s
}

func (s *AddSimpleQrBindRequest) SetStage(v string) *AddSimpleQrBindRequest {
	s.Stage = &v
	return s
}

func (s *AddSimpleQrBindRequest) SetHeader(v map[string]*string) *AddSimpleQrBindRequest {
	s.Header = v
	return s
}

func (s *AddSimpleQrBindRequest) SetAccessToken(v string) *AddSimpleQrBindRequest {
	s.AccessToken = &v
	return s
}

type AddSimpleQrBindResponse struct {
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s AddSimpleQrBindResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSimpleQrBindResponse) GoString() string {
	return s.String()
}

func (s *AddSimpleQrBindResponse) SetLogId(v string) *AddSimpleQrBindResponse {
	s.LogId = &v
	return s
}

func (s *AddSimpleQrBindResponse) SetErrNo(v int32) *AddSimpleQrBindResponse {
	s.ErrNo = &v
	return s
}

func (s *AddSimpleQrBindResponse) SetErrMsg(v string) *AddSimpleQrBindResponse {
	s.ErrMsg = &v
	return s
}

type AffiliatedDetailRequest struct {
	AccountId       *int64             `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AffiliatedId    []*int64           `json:"affiliated_id,omitempty" xml:"affiliated_id,omitempty" type:"Repeated"`
	OutAffiliatedId []*string          `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty" type:"Repeated"`
	Header          map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AffiliatedDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailRequest) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailRequest) SetAccountId(v int64) *AffiliatedDetailRequest {
	s.AccountId = &v
	return s
}

func (s *AffiliatedDetailRequest) SetAffiliatedId(v []*int64) *AffiliatedDetailRequest {
	s.AffiliatedId = v
	return s
}

func (s *AffiliatedDetailRequest) SetOutAffiliatedId(v []*string) *AffiliatedDetailRequest {
	s.OutAffiliatedId = v
	return s
}

func (s *AffiliatedDetailRequest) SetHeader(v map[string]*string) *AffiliatedDetailRequest {
	s.Header = v
	return s
}

func (s *AffiliatedDetailRequest) SetAccessToken(v string) *AffiliatedDetailRequest {
	s.AccessToken = &v
	return s
}

type AffiliatedDetailResponse struct {
	Extra *AffiliatedDetailResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *AffiliatedDetailResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AffiliatedDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponse) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponse) SetExtra(v *AffiliatedDetailResponseExtra) *AffiliatedDetailResponse {
	s.Extra = v
	return s
}

func (s *AffiliatedDetailResponse) SetData(v *AffiliatedDetailResponseData) *AffiliatedDetailResponse {
	s.Data = v
	return s
}

type AffiliatedDetailResponseData struct {
	Affiliated      map[string]*AffiliatedDetailResponseDataAffiliatedValue      `json:"affiliated,omitempty" xml:"affiliated,omitempty"`
	AffiliatedDraft map[string]*AffiliatedDetailResponseDataAffiliatedDraftValue `json:"affiliated_draft,omitempty" xml:"affiliated_draft,omitempty"`
	Description     *string                                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode       *int32                                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AffiliatedDetailResponseData) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseData) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseData) SetAffiliated(v map[string]*AffiliatedDetailResponseDataAffiliatedValue) *AffiliatedDetailResponseData {
	s.Affiliated = v
	return s
}

func (s *AffiliatedDetailResponseData) SetAffiliatedDraft(v map[string]*AffiliatedDetailResponseDataAffiliatedDraftValue) *AffiliatedDetailResponseData {
	s.AffiliatedDraft = v
	return s
}

func (s *AffiliatedDetailResponseData) SetDescription(v string) *AffiliatedDetailResponseData {
	s.Description = &v
	return s
}

func (s *AffiliatedDetailResponseData) SetErrorCode(v int32) *AffiliatedDetailResponseData {
	s.ErrorCode = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedDraftValue struct {
	Affiliated         *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated `json:"affiliated,omitempty" xml:"affiliated,omitempty"`
	ProductDraftStatus *int                                                        `json:"product_draft_status,omitempty" xml:"product_draft_status,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValue) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValue) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValue) SetAffiliated(v *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) *AffiliatedDetailResponseDataAffiliatedDraftValue {
	s.Affiliated = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValue) SetProductDraftStatus(v int) *AffiliatedDetailResponseDataAffiliatedDraftValue {
	s.ProductDraftStatus = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated struct {
	AccountId       *int64                                                                   `json:"account_id,omitempty" xml:"account_id,omitempty"`
	BizLine         *int                                                                     `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	OutUrl          *string                                                                  `json:"out_url,omitempty" xml:"out_url,omitempty"`
	OutId           *string                                                                  `json:"out_id,omitempty" xml:"out_id,omitempty"`
	Stock           *int64                                                                   `json:"stock,omitempty" xml:"stock,omitempty"`
	Name            *string                                                                  `json:"name,omitempty" xml:"name,omitempty"`
	PoiList         []*AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	FulfillmentType []*int                                                                   `json:"fulfillment_type,omitempty" xml:"fulfillment_type,omitempty" type:"Repeated"`
	CategoryId      *int64                                                                   `json:"category_id,omitempty" xml:"category_id,omitempty"`
	ProductId       *int64                                                                   `json:"product_id,omitempty" xml:"product_id,omitempty"`
	SettleInfo      *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo    `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	Price           *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice         `json:"price,omitempty" xml:"price,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetAccountId(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.AccountId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetBizLine(v int) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.BizLine = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetOutUrl(v string) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.OutUrl = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetOutId(v string) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.OutId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetStock(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.Stock = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetName(v string) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.Name = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetPoiList(v []*AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.PoiList = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetFulfillmentType(v []*int) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.FulfillmentType = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetCategoryId(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.CategoryId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetProductId(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.ProductId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetSettleInfo(v *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.SettleInfo = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated) SetPrice(v *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliated {
	s.Price = v
	return s
}

type AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem struct {
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem) SetExtId(v string) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem {
	s.ExtId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem) SetPoiId(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPoiListItem {
	s.PoiId = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice) SetActualAmount(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice {
	s.ActualAmount = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice) SetOriginAmount(v int64) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedPrice {
	s.OriginAmount = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo) SetSettleType(v int) *AffiliatedDetailResponseDataAffiliatedDraftValueAffiliatedSettleInfo {
	s.SettleType = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedValue struct {
	ProductId       *int64                                                    `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Name            *string                                                   `json:"name,omitempty" xml:"name,omitempty"`
	Price           *AffiliatedDetailResponseDataAffiliatedValuePrice         `json:"price,omitempty" xml:"price,omitempty"`
	Stock           *int64                                                    `json:"stock,omitempty" xml:"stock,omitempty"`
	OutId           *string                                                   `json:"out_id,omitempty" xml:"out_id,omitempty"`
	OutUrl          *string                                                   `json:"out_url,omitempty" xml:"out_url,omitempty"`
	CategoryId      *int64                                                    `json:"category_id,omitempty" xml:"category_id,omitempty"`
	SettleInfo      *AffiliatedDetailResponseDataAffiliatedValueSettleInfo    `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	BizLine         *int                                                      `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	FulfillmentType []*int                                                    `json:"fulfillment_type,omitempty" xml:"fulfillment_type,omitempty" type:"Repeated"`
	PoiList         []*AffiliatedDetailResponseDataAffiliatedValuePoiListItem `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	AccountId       *int64                                                    `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedValue) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedValue) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetProductId(v int64) *AffiliatedDetailResponseDataAffiliatedValue {
	s.ProductId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetName(v string) *AffiliatedDetailResponseDataAffiliatedValue {
	s.Name = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetPrice(v *AffiliatedDetailResponseDataAffiliatedValuePrice) *AffiliatedDetailResponseDataAffiliatedValue {
	s.Price = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetStock(v int64) *AffiliatedDetailResponseDataAffiliatedValue {
	s.Stock = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetOutId(v string) *AffiliatedDetailResponseDataAffiliatedValue {
	s.OutId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetOutUrl(v string) *AffiliatedDetailResponseDataAffiliatedValue {
	s.OutUrl = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetCategoryId(v int64) *AffiliatedDetailResponseDataAffiliatedValue {
	s.CategoryId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetSettleInfo(v *AffiliatedDetailResponseDataAffiliatedValueSettleInfo) *AffiliatedDetailResponseDataAffiliatedValue {
	s.SettleInfo = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetBizLine(v int) *AffiliatedDetailResponseDataAffiliatedValue {
	s.BizLine = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetFulfillmentType(v []*int) *AffiliatedDetailResponseDataAffiliatedValue {
	s.FulfillmentType = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetPoiList(v []*AffiliatedDetailResponseDataAffiliatedValuePoiListItem) *AffiliatedDetailResponseDataAffiliatedValue {
	s.PoiList = v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValue) SetAccountId(v int64) *AffiliatedDetailResponseDataAffiliatedValue {
	s.AccountId = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedValuePoiListItem struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedValuePoiListItem) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedValuePoiListItem) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedValuePoiListItem) SetPoiId(v int64) *AffiliatedDetailResponseDataAffiliatedValuePoiListItem {
	s.PoiId = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValuePoiListItem) SetExtId(v string) *AffiliatedDetailResponseDataAffiliatedValuePoiListItem {
	s.ExtId = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedValuePrice struct {
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedValuePrice) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedValuePrice) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedValuePrice) SetActualAmount(v int64) *AffiliatedDetailResponseDataAffiliatedValuePrice {
	s.ActualAmount = &v
	return s
}

func (s *AffiliatedDetailResponseDataAffiliatedValuePrice) SetOriginAmount(v int64) *AffiliatedDetailResponseDataAffiliatedValuePrice {
	s.OriginAmount = &v
	return s
}

type AffiliatedDetailResponseDataAffiliatedValueSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s AffiliatedDetailResponseDataAffiliatedValueSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseDataAffiliatedValueSettleInfo) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseDataAffiliatedValueSettleInfo) SetSettleType(v int) *AffiliatedDetailResponseDataAffiliatedValueSettleInfo {
	s.SettleType = &v
	return s
}

type AffiliatedDetailResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s AffiliatedDetailResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedDetailResponseExtra) GoString() string {
	return s.String()
}

func (s *AffiliatedDetailResponseExtra) SetSubErrorCode(v int32) *AffiliatedDetailResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AffiliatedDetailResponseExtra) SetDescription(v string) *AffiliatedDetailResponseExtra {
	s.Description = &v
	return s
}

func (s *AffiliatedDetailResponseExtra) SetErrorCode(v int32) *AffiliatedDetailResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedDetailResponseExtra) SetLogid(v string) *AffiliatedDetailResponseExtra {
	s.Logid = &v
	return s
}

func (s *AffiliatedDetailResponseExtra) SetNow(v int64) *AffiliatedDetailResponseExtra {
	s.Now = &v
	return s
}

func (s *AffiliatedDetailResponseExtra) SetSubDescription(v string) *AffiliatedDetailResponseExtra {
	s.SubDescription = &v
	return s
}

type AffiliatedOperateRequest struct {
	Op              *int                          `json:"op,omitempty" xml:"op,omitempty" require:"true"`
	OutAffiliatedId *string                       `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	Header          map[string]*string            `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                       `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base            *AffiliatedOperateRequestBase `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *int64                        `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s AffiliatedOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateRequest) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateRequest) SetOp(v int) *AffiliatedOperateRequest {
	s.Op = &v
	return s
}

func (s *AffiliatedOperateRequest) SetOutAffiliatedId(v string) *AffiliatedOperateRequest {
	s.OutAffiliatedId = &v
	return s
}

func (s *AffiliatedOperateRequest) SetHeader(v map[string]*string) *AffiliatedOperateRequest {
	s.Header = v
	return s
}

func (s *AffiliatedOperateRequest) SetAccessToken(v string) *AffiliatedOperateRequest {
	s.AccessToken = &v
	return s
}

func (s *AffiliatedOperateRequest) SetBase(v *AffiliatedOperateRequestBase) *AffiliatedOperateRequest {
	s.Base = v
	return s
}

func (s *AffiliatedOperateRequest) SetAccountId(v int64) *AffiliatedOperateRequest {
	s.AccountId = &v
	return s
}

type AffiliatedOperateRequestBase struct {
	Extra      map[string]*string                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LogID      *string                                 `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *AffiliatedOperateRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                 `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                 `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                 `json:"Client,omitempty" xml:"Client,omitempty"`
}

func (s AffiliatedOperateRequestBase) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateRequestBase) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateRequestBase) SetExtra(v map[string]*string) *AffiliatedOperateRequestBase {
	s.Extra = v
	return s
}

func (s *AffiliatedOperateRequestBase) SetLogID(v string) *AffiliatedOperateRequestBase {
	s.LogID = &v
	return s
}

func (s *AffiliatedOperateRequestBase) SetTrafficEnv(v *AffiliatedOperateRequestBaseTrafficEnv) *AffiliatedOperateRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *AffiliatedOperateRequestBase) SetAddr(v string) *AffiliatedOperateRequestBase {
	s.Addr = &v
	return s
}

func (s *AffiliatedOperateRequestBase) SetCaller(v string) *AffiliatedOperateRequestBase {
	s.Caller = &v
	return s
}

func (s *AffiliatedOperateRequestBase) SetClient(v string) *AffiliatedOperateRequestBase {
	s.Client = &v
	return s
}

type AffiliatedOperateRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s AffiliatedOperateRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateRequestBaseTrafficEnv) SetEnv(v string) *AffiliatedOperateRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *AffiliatedOperateRequestBaseTrafficEnv) SetOpen(v bool) *AffiliatedOperateRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type AffiliatedOperateResponse struct {
	Extra    *AffiliatedOperateResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
	BaseResp *AffiliatedOperateResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *AffiliatedOperateResponseData     `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AffiliatedOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateResponse) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateResponse) SetExtra(v *AffiliatedOperateResponseExtra) *AffiliatedOperateResponse {
	s.Extra = v
	return s
}

func (s *AffiliatedOperateResponse) SetBaseResp(v *AffiliatedOperateResponseBaseResp) *AffiliatedOperateResponse {
	s.BaseResp = v
	return s
}

func (s *AffiliatedOperateResponse) SetData(v *AffiliatedOperateResponseData) *AffiliatedOperateResponse {
	s.Data = v
	return s
}

type AffiliatedOperateResponseBaseResp struct {
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s AffiliatedOperateResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateResponseBaseResp) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateResponseBaseResp) SetExtra(v map[string]*string) *AffiliatedOperateResponseBaseResp {
	s.Extra = v
	return s
}

func (s *AffiliatedOperateResponseBaseResp) SetStatusCode(v int32) *AffiliatedOperateResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *AffiliatedOperateResponseBaseResp) SetStatusMessage(v string) *AffiliatedOperateResponseBaseResp {
	s.StatusMessage = &v
	return s
}

type AffiliatedOperateResponseData struct {
	OutAffiliatedId *string `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	AffiliatedId    *int64  `json:"affiliated_id,omitempty" xml:"affiliated_id,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode       *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AffiliatedOperateResponseData) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateResponseData) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateResponseData) SetOutAffiliatedId(v string) *AffiliatedOperateResponseData {
	s.OutAffiliatedId = &v
	return s
}

func (s *AffiliatedOperateResponseData) SetAffiliatedId(v int64) *AffiliatedOperateResponseData {
	s.AffiliatedId = &v
	return s
}

func (s *AffiliatedOperateResponseData) SetDescription(v string) *AffiliatedOperateResponseData {
	s.Description = &v
	return s
}

func (s *AffiliatedOperateResponseData) SetErrorCode(v int32) *AffiliatedOperateResponseData {
	s.ErrorCode = &v
	return s
}

type AffiliatedOperateResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s AffiliatedOperateResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedOperateResponseExtra) GoString() string {
	return s.String()
}

func (s *AffiliatedOperateResponseExtra) SetDescription(v string) *AffiliatedOperateResponseExtra {
	s.Description = &v
	return s
}

func (s *AffiliatedOperateResponseExtra) SetErrorCode(v int32) *AffiliatedOperateResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedOperateResponseExtra) SetLogid(v string) *AffiliatedOperateResponseExtra {
	s.Logid = &v
	return s
}

func (s *AffiliatedOperateResponseExtra) SetNow(v int64) *AffiliatedOperateResponseExtra {
	s.Now = &v
	return s
}

func (s *AffiliatedOperateResponseExtra) SetSubDescription(v string) *AffiliatedOperateResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AffiliatedOperateResponseExtra) SetSubErrorCode(v int32) *AffiliatedOperateResponseExtra {
	s.SubErrorCode = &v
	return s
}

type AffiliatedPoiSellOutRequest struct {
	Header          map[string]*string                               `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken     *string                                          `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Base            *AffiliatedPoiSellOutRequestBase                 `json:"Base,omitempty" xml:"Base,omitempty"`
	AccountId       *int64                                           `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	IgnoreFailPoi   *bool                                            `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
	OutAffiliatedId *string                                          `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	PoiSellOutRule  []*AffiliatedPoiSellOutRequestPoiSellOutRuleItem `json:"poi_sell_out_rule,omitempty" xml:"poi_sell_out_rule,omitempty" require:"true" type:"Repeated"`
}

func (s AffiliatedPoiSellOutRequest) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutRequest) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutRequest) SetHeader(v map[string]*string) *AffiliatedPoiSellOutRequest {
	s.Header = v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetAccessToken(v string) *AffiliatedPoiSellOutRequest {
	s.AccessToken = &v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetBase(v *AffiliatedPoiSellOutRequestBase) *AffiliatedPoiSellOutRequest {
	s.Base = v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetAccountId(v int64) *AffiliatedPoiSellOutRequest {
	s.AccountId = &v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetIgnoreFailPoi(v bool) *AffiliatedPoiSellOutRequest {
	s.IgnoreFailPoi = &v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetOutAffiliatedId(v string) *AffiliatedPoiSellOutRequest {
	s.OutAffiliatedId = &v
	return s
}

func (s *AffiliatedPoiSellOutRequest) SetPoiSellOutRule(v []*AffiliatedPoiSellOutRequestPoiSellOutRuleItem) *AffiliatedPoiSellOutRequest {
	s.PoiSellOutRule = v
	return s
}

type AffiliatedPoiSellOutRequestBase struct {
	LogID      *string                                    `json:"LogID,omitempty" xml:"LogID,omitempty"`
	TrafficEnv *AffiliatedPoiSellOutRequestBaseTrafficEnv `json:"TrafficEnv,omitempty" xml:"TrafficEnv,omitempty"`
	Addr       *string                                    `json:"Addr,omitempty" xml:"Addr,omitempty"`
	Caller     *string                                    `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Client     *string                                    `json:"Client,omitempty" xml:"Client,omitempty"`
	Extra      map[string]*string                         `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s AffiliatedPoiSellOutRequestBase) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutRequestBase) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutRequestBase) SetLogID(v string) *AffiliatedPoiSellOutRequestBase {
	s.LogID = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestBase) SetTrafficEnv(v *AffiliatedPoiSellOutRequestBaseTrafficEnv) *AffiliatedPoiSellOutRequestBase {
	s.TrafficEnv = v
	return s
}

func (s *AffiliatedPoiSellOutRequestBase) SetAddr(v string) *AffiliatedPoiSellOutRequestBase {
	s.Addr = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestBase) SetCaller(v string) *AffiliatedPoiSellOutRequestBase {
	s.Caller = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestBase) SetClient(v string) *AffiliatedPoiSellOutRequestBase {
	s.Client = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestBase) SetExtra(v map[string]*string) *AffiliatedPoiSellOutRequestBase {
	s.Extra = v
	return s
}

type AffiliatedPoiSellOutRequestBaseTrafficEnv struct {
	Env  *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Open *bool   `json:"Open,omitempty" xml:"Open,omitempty"`
}

func (s AffiliatedPoiSellOutRequestBaseTrafficEnv) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutRequestBaseTrafficEnv) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutRequestBaseTrafficEnv) SetEnv(v string) *AffiliatedPoiSellOutRequestBaseTrafficEnv {
	s.Env = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestBaseTrafficEnv) SetOpen(v bool) *AffiliatedPoiSellOutRequestBaseTrafficEnv {
	s.Open = &v
	return s
}

type AffiliatedPoiSellOutRequestPoiSellOutRuleItem struct {
	SellOutRule *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule `json:"sell_out_rule,omitempty" xml:"sell_out_rule,omitempty"`
	PoiId       *string                                                   `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s AffiliatedPoiSellOutRequestPoiSellOutRuleItem) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutRequestPoiSellOutRuleItem) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutRequestPoiSellOutRuleItem) SetSellOutRule(v *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule) *AffiliatedPoiSellOutRequestPoiSellOutRuleItem {
	s.SellOutRule = v
	return s
}

func (s *AffiliatedPoiSellOutRequestPoiSellOutRuleItem) SetPoiId(v string) *AffiliatedPoiSellOutRequestPoiSellOutRuleItem {
	s.PoiId = &v
	return s
}

type AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule struct {
	LongTermSellOutStatus *int `json:"long_term_sell_out_status,omitempty" xml:"long_term_sell_out_status,omitempty"`
	SellOutType           *int `json:"sell_out_type,omitempty" xml:"sell_out_type,omitempty"`
}

func (s AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule) SetLongTermSellOutStatus(v int) *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule {
	s.LongTermSellOutStatus = &v
	return s
}

func (s *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule) SetSellOutType(v int) *AffiliatedPoiSellOutRequestPoiSellOutRuleItemSellOutRule {
	s.SellOutType = &v
	return s
}

type AffiliatedPoiSellOutResponse struct {
	BaseResp *AffiliatedPoiSellOutResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
	Data     *AffiliatedPoiSellOutResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *AffiliatedPoiSellOutResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AffiliatedPoiSellOutResponse) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutResponse) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutResponse) SetBaseResp(v *AffiliatedPoiSellOutResponseBaseResp) *AffiliatedPoiSellOutResponse {
	s.BaseResp = v
	return s
}

func (s *AffiliatedPoiSellOutResponse) SetData(v *AffiliatedPoiSellOutResponseData) *AffiliatedPoiSellOutResponse {
	s.Data = v
	return s
}

func (s *AffiliatedPoiSellOutResponse) SetExtra(v *AffiliatedPoiSellOutResponseExtra) *AffiliatedPoiSellOutResponse {
	s.Extra = v
	return s
}

type AffiliatedPoiSellOutResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s AffiliatedPoiSellOutResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutResponseBaseResp) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutResponseBaseResp) SetStatusCode(v int32) *AffiliatedPoiSellOutResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseBaseResp) SetStatusMessage(v string) *AffiliatedPoiSellOutResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseBaseResp) SetExtra(v map[string]*string) *AffiliatedPoiSellOutResponseBaseResp {
	s.Extra = v
	return s
}

type AffiliatedPoiSellOutResponseData struct {
	Description     *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode       *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	OutAffiliatedId *string `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	AffiliatedId    *int64  `json:"affiliated_id,omitempty" xml:"affiliated_id,omitempty"`
}

func (s AffiliatedPoiSellOutResponseData) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutResponseData) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutResponseData) SetDescription(v string) *AffiliatedPoiSellOutResponseData {
	s.Description = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseData) SetErrorCode(v int32) *AffiliatedPoiSellOutResponseData {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseData) SetOutAffiliatedId(v string) *AffiliatedPoiSellOutResponseData {
	s.OutAffiliatedId = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseData) SetAffiliatedId(v int64) *AffiliatedPoiSellOutResponseData {
	s.AffiliatedId = &v
	return s
}

type AffiliatedPoiSellOutResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s AffiliatedPoiSellOutResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedPoiSellOutResponseExtra) GoString() string {
	return s.String()
}

func (s *AffiliatedPoiSellOutResponseExtra) SetDescription(v string) *AffiliatedPoiSellOutResponseExtra {
	s.Description = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseExtra) SetErrorCode(v int32) *AffiliatedPoiSellOutResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseExtra) SetLogid(v string) *AffiliatedPoiSellOutResponseExtra {
	s.Logid = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseExtra) SetNow(v int64) *AffiliatedPoiSellOutResponseExtra {
	s.Now = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseExtra) SetSubDescription(v string) *AffiliatedPoiSellOutResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AffiliatedPoiSellOutResponseExtra) SetSubErrorCode(v int32) *AffiliatedPoiSellOutResponseExtra {
	s.SubErrorCode = &v
	return s
}

type AffiliatedSaveRequest struct {
	AccountId     *int64                           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	Header        map[string]*string               `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string                          `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Affiliated    *AffiliatedSaveRequestAffiliated `json:"affiliated,omitempty" xml:"affiliated,omitempty" require:"true"`
	IgnoreFailPoi *bool                            `json:"ignore_fail_poi,omitempty" xml:"ignore_fail_poi,omitempty"`
}

func (s AffiliatedSaveRequest) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveRequest) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveRequest) SetAccountId(v int64) *AffiliatedSaveRequest {
	s.AccountId = &v
	return s
}

func (s *AffiliatedSaveRequest) SetHeader(v map[string]*string) *AffiliatedSaveRequest {
	s.Header = v
	return s
}

func (s *AffiliatedSaveRequest) SetAccessToken(v string) *AffiliatedSaveRequest {
	s.AccessToken = &v
	return s
}

func (s *AffiliatedSaveRequest) SetAffiliated(v *AffiliatedSaveRequestAffiliated) *AffiliatedSaveRequest {
	s.Affiliated = v
	return s
}

func (s *AffiliatedSaveRequest) SetIgnoreFailPoi(v bool) *AffiliatedSaveRequest {
	s.IgnoreFailPoi = &v
	return s
}

type AffiliatedSaveRequestAffiliated struct {
	OutUrl          *string                                       `json:"out_url,omitempty" xml:"out_url,omitempty"`
	Name            *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	CategoryId      *int64                                        `json:"category_id,omitempty" xml:"category_id,omitempty"`
	SettleInfo      *AffiliatedSaveRequestAffiliatedSettleInfo    `json:"settle_info,omitempty" xml:"settle_info,omitempty"`
	BizLine         *int                                          `json:"biz_line,omitempty" xml:"biz_line,omitempty"`
	FulfillmentType []*int                                        `json:"fulfillment_type,omitempty" xml:"fulfillment_type,omitempty" type:"Repeated"`
	OutId           *string                                       `json:"out_id,omitempty" xml:"out_id,omitempty"`
	PoiList         []*AffiliatedSaveRequestAffiliatedPoiListItem `json:"poi_list,omitempty" xml:"poi_list,omitempty" type:"Repeated"`
	ProductId       *int64                                        `json:"product_id,omitempty" xml:"product_id,omitempty"`
	Price           *AffiliatedSaveRequestAffiliatedPrice         `json:"price,omitempty" xml:"price,omitempty"`
	AccountId       *int64                                        `json:"account_id,omitempty" xml:"account_id,omitempty"`
}

func (s AffiliatedSaveRequestAffiliated) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveRequestAffiliated) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveRequestAffiliated) SetOutUrl(v string) *AffiliatedSaveRequestAffiliated {
	s.OutUrl = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetName(v string) *AffiliatedSaveRequestAffiliated {
	s.Name = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetCategoryId(v int64) *AffiliatedSaveRequestAffiliated {
	s.CategoryId = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetSettleInfo(v *AffiliatedSaveRequestAffiliatedSettleInfo) *AffiliatedSaveRequestAffiliated {
	s.SettleInfo = v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetBizLine(v int) *AffiliatedSaveRequestAffiliated {
	s.BizLine = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetFulfillmentType(v []*int) *AffiliatedSaveRequestAffiliated {
	s.FulfillmentType = v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetOutId(v string) *AffiliatedSaveRequestAffiliated {
	s.OutId = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetPoiList(v []*AffiliatedSaveRequestAffiliatedPoiListItem) *AffiliatedSaveRequestAffiliated {
	s.PoiList = v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetProductId(v int64) *AffiliatedSaveRequestAffiliated {
	s.ProductId = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetPrice(v *AffiliatedSaveRequestAffiliatedPrice) *AffiliatedSaveRequestAffiliated {
	s.Price = v
	return s
}

func (s *AffiliatedSaveRequestAffiliated) SetAccountId(v int64) *AffiliatedSaveRequestAffiliated {
	s.AccountId = &v
	return s
}

type AffiliatedSaveRequestAffiliatedPoiListItem struct {
	PoiId *int64  `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	ExtId *string `json:"ext_id,omitempty" xml:"ext_id,omitempty"`
}

func (s AffiliatedSaveRequestAffiliatedPoiListItem) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveRequestAffiliatedPoiListItem) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveRequestAffiliatedPoiListItem) SetPoiId(v int64) *AffiliatedSaveRequestAffiliatedPoiListItem {
	s.PoiId = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliatedPoiListItem) SetExtId(v string) *AffiliatedSaveRequestAffiliatedPoiListItem {
	s.ExtId = &v
	return s
}

type AffiliatedSaveRequestAffiliatedPrice struct {
	OriginAmount *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	ActualAmount *int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
}

func (s AffiliatedSaveRequestAffiliatedPrice) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveRequestAffiliatedPrice) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveRequestAffiliatedPrice) SetOriginAmount(v int64) *AffiliatedSaveRequestAffiliatedPrice {
	s.OriginAmount = &v
	return s
}

func (s *AffiliatedSaveRequestAffiliatedPrice) SetActualAmount(v int64) *AffiliatedSaveRequestAffiliatedPrice {
	s.ActualAmount = &v
	return s
}

type AffiliatedSaveRequestAffiliatedSettleInfo struct {
	SettleType *int `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
}

func (s AffiliatedSaveRequestAffiliatedSettleInfo) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveRequestAffiliatedSettleInfo) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveRequestAffiliatedSettleInfo) SetSettleType(v int) *AffiliatedSaveRequestAffiliatedSettleInfo {
	s.SettleType = &v
	return s
}

type AffiliatedSaveResponse struct {
	Data     *AffiliatedSaveResponseData     `json:"data,omitempty" xml:"data,omitempty"`
	Extra    *AffiliatedSaveResponseExtra    `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	BaseResp *AffiliatedSaveResponseBaseResp `json:"BaseResp,omitempty" xml:"BaseResp,omitempty" require:"true"`
}

func (s AffiliatedSaveResponse) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveResponse) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveResponse) SetData(v *AffiliatedSaveResponseData) *AffiliatedSaveResponse {
	s.Data = v
	return s
}

func (s *AffiliatedSaveResponse) SetExtra(v *AffiliatedSaveResponseExtra) *AffiliatedSaveResponse {
	s.Extra = v
	return s
}

func (s *AffiliatedSaveResponse) SetBaseResp(v *AffiliatedSaveResponseBaseResp) *AffiliatedSaveResponse {
	s.BaseResp = v
	return s
}

type AffiliatedSaveResponseBaseResp struct {
	StatusCode    *int32             `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage *string            `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Extra         map[string]*string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s AffiliatedSaveResponseBaseResp) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveResponseBaseResp) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveResponseBaseResp) SetStatusCode(v int32) *AffiliatedSaveResponseBaseResp {
	s.StatusCode = &v
	return s
}

func (s *AffiliatedSaveResponseBaseResp) SetStatusMessage(v string) *AffiliatedSaveResponseBaseResp {
	s.StatusMessage = &v
	return s
}

func (s *AffiliatedSaveResponseBaseResp) SetExtra(v map[string]*string) *AffiliatedSaveResponseBaseResp {
	s.Extra = v
	return s
}

type AffiliatedSaveResponseData struct {
	ErrorCode       *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	OutAffiliatedId *string `json:"out_affiliated_id,omitempty" xml:"out_affiliated_id,omitempty"`
	ProductId       *int64  `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AffiliatedSaveResponseData) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveResponseData) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveResponseData) SetErrorCode(v int32) *AffiliatedSaveResponseData {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedSaveResponseData) SetOutAffiliatedId(v string) *AffiliatedSaveResponseData {
	s.OutAffiliatedId = &v
	return s
}

func (s *AffiliatedSaveResponseData) SetProductId(v int64) *AffiliatedSaveResponseData {
	s.ProductId = &v
	return s
}

func (s *AffiliatedSaveResponseData) SetDescription(v string) *AffiliatedSaveResponseData {
	s.Description = &v
	return s
}

type AffiliatedSaveResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AffiliatedSaveResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AffiliatedSaveResponseExtra) GoString() string {
	return s.String()
}

func (s *AffiliatedSaveResponseExtra) SetErrorCode(v int32) *AffiliatedSaveResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AffiliatedSaveResponseExtra) SetLogid(v string) *AffiliatedSaveResponseExtra {
	s.Logid = &v
	return s
}

func (s *AffiliatedSaveResponseExtra) SetNow(v int64) *AffiliatedSaveResponseExtra {
	s.Now = &v
	return s
}

func (s *AffiliatedSaveResponseExtra) SetSubDescription(v string) *AffiliatedSaveResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AffiliatedSaveResponseExtra) SetSubErrorCode(v int32) *AffiliatedSaveResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AffiliatedSaveResponseExtra) SetDescription(v string) *AffiliatedSaveResponseExtra {
	s.Description = &v
	return s
}

type AfterSaleOrderAuditRequest struct {
	CertificateId *string            `json:"certificate_id,omitempty" xml:"certificate_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AuditStatus   *bool              `json:"audit_status,omitempty" xml:"audit_status,omitempty" require:"true"`
	RejectReason  *string            `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
}

func (s AfterSaleOrderAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderAuditRequest) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderAuditRequest) SetCertificateId(v string) *AfterSaleOrderAuditRequest {
	s.CertificateId = &v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetHeader(v map[string]*string) *AfterSaleOrderAuditRequest {
	s.Header = v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetAccessToken(v string) *AfterSaleOrderAuditRequest {
	s.AccessToken = &v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetAuditStatus(v bool) *AfterSaleOrderAuditRequest {
	s.AuditStatus = &v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetRejectReason(v string) *AfterSaleOrderAuditRequest {
	s.RejectReason = &v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetAccountId(v string) *AfterSaleOrderAuditRequest {
	s.AccountId = &v
	return s
}

func (s *AfterSaleOrderAuditRequest) SetOrderId(v string) *AfterSaleOrderAuditRequest {
	s.OrderId = &v
	return s
}

type AfterSaleOrderAuditResponse struct {
	Extra *AfterSaleOrderAuditResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *AfterSaleOrderAuditResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AfterSaleOrderAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderAuditResponse) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderAuditResponse) SetExtra(v *AfterSaleOrderAuditResponseExtra) *AfterSaleOrderAuditResponse {
	s.Extra = v
	return s
}

func (s *AfterSaleOrderAuditResponse) SetData(v *AfterSaleOrderAuditResponseData) *AfterSaleOrderAuditResponse {
	s.Data = v
	return s
}

type AfterSaleOrderAuditResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Reentry       *bool   `json:"reentry,omitempty" xml:"reentry,omitempty"`
}

func (s AfterSaleOrderAuditResponseData) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderAuditResponseData) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderAuditResponseData) SetGwErrorCode(v int32) *AfterSaleOrderAuditResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AfterSaleOrderAuditResponseData) SetGwDescription(v string) *AfterSaleOrderAuditResponseData {
	s.GwDescription = &v
	return s
}

func (s *AfterSaleOrderAuditResponseData) SetReentry(v bool) *AfterSaleOrderAuditResponseData {
	s.Reentry = &v
	return s
}

type AfterSaleOrderAuditResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s AfterSaleOrderAuditResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderAuditResponseExtra) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderAuditResponseExtra) SetSubDescription(v string) *AfterSaleOrderAuditResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AfterSaleOrderAuditResponseExtra) SetLogid(v string) *AfterSaleOrderAuditResponseExtra {
	s.Logid = &v
	return s
}

func (s *AfterSaleOrderAuditResponseExtra) SetNow(v int64) *AfterSaleOrderAuditResponseExtra {
	s.Now = &v
	return s
}

func (s *AfterSaleOrderAuditResponseExtra) SetErrorCode(v int32) *AfterSaleOrderAuditResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AfterSaleOrderAuditResponseExtra) SetDescription(v string) *AfterSaleOrderAuditResponseExtra {
	s.Description = &v
	return s
}

func (s *AfterSaleOrderAuditResponseExtra) SetSubErrorCode(v int32) *AfterSaleOrderAuditResponseExtra {
	s.SubErrorCode = &v
	return s
}

type AfterSaleOrderDetailGetRequest struct {
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	CertificateId *string            `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	OrderId       *string            `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AfterSaleOrderDetailGetRequest) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetRequest) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetRequest) SetAccountId(v string) *AfterSaleOrderDetailGetRequest {
	s.AccountId = &v
	return s
}

func (s *AfterSaleOrderDetailGetRequest) SetCertificateId(v string) *AfterSaleOrderDetailGetRequest {
	s.CertificateId = &v
	return s
}

func (s *AfterSaleOrderDetailGetRequest) SetOrderId(v string) *AfterSaleOrderDetailGetRequest {
	s.OrderId = &v
	return s
}

func (s *AfterSaleOrderDetailGetRequest) SetHeader(v map[string]*string) *AfterSaleOrderDetailGetRequest {
	s.Header = v
	return s
}

func (s *AfterSaleOrderDetailGetRequest) SetAccessToken(v string) *AfterSaleOrderDetailGetRequest {
	s.AccessToken = &v
	return s
}

type AfterSaleOrderDetailGetResponse struct {
	Extra *AfterSaleOrderDetailGetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *AfterSaleOrderDetailGetResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AfterSaleOrderDetailGetResponse) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponse) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponse) SetExtra(v *AfterSaleOrderDetailGetResponseExtra) *AfterSaleOrderDetailGetResponse {
	s.Extra = v
	return s
}

func (s *AfterSaleOrderDetailGetResponse) SetData(v *AfterSaleOrderDetailGetResponseData) *AfterSaleOrderDetailGetResponse {
	s.Data = v
	return s
}

type AfterSaleOrderDetailGetResponseData struct {
	AfterSaleOrderList []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem `json:"after_sale_order_list,omitempty" xml:"after_sale_order_list,omitempty" type:"Repeated"`
	Cursor             *string                                                      `json:"cursor,omitempty" xml:"cursor,omitempty"`
	HasMore            *bool                                                        `json:"has_more,omitempty" xml:"has_more,omitempty"`
	GwErrorCode        *int32                                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription      *string                                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AfterSaleOrderDetailGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseData) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseData) SetAfterSaleOrderList(v []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) *AfterSaleOrderDetailGetResponseData {
	s.AfterSaleOrderList = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseData) SetCursor(v string) *AfterSaleOrderDetailGetResponseData {
	s.Cursor = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseData) SetHasMore(v bool) *AfterSaleOrderDetailGetResponseData {
	s.HasMore = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseData) SetGwErrorCode(v int32) *AfterSaleOrderDetailGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseData) SetGwDescription(v string) *AfterSaleOrderDetailGetResponseData {
	s.GwDescription = &v
	return s
}

type AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem struct {
	MarketRefundAmount    *int64                                                                         `json:"market_refund_amount,omitempty" xml:"market_refund_amount,omitempty"`
	Status                *int                                                                           `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	UserRefundAmount      *int64                                                                         `json:"user_refund_amount,omitempty" xml:"user_refund_amount,omitempty"`
	UserDeductFeeAmount   *int64                                                                         `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
	AuditTime             *int64                                                                         `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
	AuditResult           *string                                                                        `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	RefundType            *int64                                                                         `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	RefundAmount          *int64                                                                         `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	MerchantAccountId     *int64                                                                         `json:"merchant_account_id,omitempty" xml:"merchant_account_id,omitempty"`
	OrderType             *int64                                                                         `json:"order_type,omitempty" xml:"order_type,omitempty"`
	RefundInfoList        []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem `json:"refund_info_list,omitempty" xml:"refund_info_list,omitempty" type:"Repeated"`
	CompleteTime          *int64                                                                         `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
	MarketDeductFeeAmount *int64                                                                         `json:"market_deduct_fee_amount,omitempty" xml:"market_deduct_fee_amount,omitempty"`
	AfterSaleId           *string                                                                        `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty" require:"true"`
	OutRefundPaymentId    *string                                                                        `json:"out_refund_payment_id,omitempty" xml:"out_refund_payment_id,omitempty"`
	RealRefundAmount      *int64                                                                         `json:"real_refund_amount,omitempty" xml:"real_refund_amount,omitempty"`
	CreateTime            *int64                                                                         `json:"create_time,omitempty" xml:"create_time,omitempty"`
	TotalRefundAmount     *int64                                                                         `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
	DeductFeeAmount       *int64                                                                         `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
	OrderId               *string                                                                        `json:"order_id,omitempty" xml:"order_id,omitempty"`
	UpdateTime            *int64                                                                         `json:"update_time,omitempty" xml:"update_time,omitempty"`
	Reason                *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason               `json:"reason,omitempty" xml:"reason,omitempty"`
	RejectReason          *string                                                                        `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	OutBizAfterSaleId     *string                                                                        `json:"out_biz_after_sale_id,omitempty" xml:"out_biz_after_sale_id,omitempty"`
	TradeType             *int64                                                                         `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetMarketRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.MarketRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetStatus(v int) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.Status = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetUserRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.UserRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetUserDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.UserDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetAuditTime(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.AuditTime = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetAuditResult(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.AuditResult = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetRefundType(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.RefundType = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.RefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetMerchantAccountId(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.MerchantAccountId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetOrderType(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.OrderType = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetRefundInfoList(v []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.RefundInfoList = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetCompleteTime(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.CompleteTime = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetMarketDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.MarketDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetAfterSaleId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.AfterSaleId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetOutRefundPaymentId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.OutRefundPaymentId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetRealRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.RealRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetCreateTime(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.CreateTime = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetTotalRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.TotalRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.DeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetOrderId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.OrderId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetUpdateTime(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.UpdateTime = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetReason(v *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.Reason = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetRejectReason(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.RejectReason = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetOutBizAfterSaleId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.OutBizAfterSaleId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem) SetTradeType(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItem {
	s.TradeType = &v
	return s
}

type AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason struct {
	ReasonCode []*int64                                                                         `json:"reason_code,omitempty" xml:"reason_code,omitempty" type:"Repeated"`
	ShowReason []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem `json:"show_reason,omitempty" xml:"show_reason,omitempty" type:"Repeated"`
	Desc       *string                                                                          `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) SetReasonCode(v []*int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason {
	s.ReasonCode = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) SetShowReason(v []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason {
	s.ShowReason = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason) SetDesc(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReason {
	s.Desc = &v
	return s
}

type AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem struct {
	Msg        *string `json:"msg,omitempty" xml:"msg,omitempty"`
	ReasonCode *int64  `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem) SetMsg(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem {
	s.Msg = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem) SetReasonCode(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemReasonShowReasonItem {
	s.ReasonCode = &v
	return s
}

type AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem struct {
	CertificateId                 *string                                                                                                   `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	RealRefundAmount              *int64                                                                                                    `json:"real_refund_amount,omitempty" xml:"real_refund_amount,omitempty"`
	TimesCardRefundInfoList       []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem `json:"times_card_refund_info_list,omitempty" xml:"times_card_refund_info_list,omitempty" type:"Repeated"`
	PlatformMarketDeductFeeAmount *int64                                                                                                    `json:"platform_market_deduct_fee_amount,omitempty" xml:"platform_market_deduct_fee_amount,omitempty"`
	MarketRefundAmount            *int64                                                                                                    `json:"market_refund_amount,omitempty" xml:"market_refund_amount,omitempty"`
	UserDeductFeeAmount           *int64                                                                                                    `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
	DeductFeeAmount               *int64                                                                                                    `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
	RefundId                      *string                                                                                                   `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	RefundStatus                  *int                                                                                                      `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	MarketDeductFeeAmount         *int64                                                                                                    `json:"market_deduct_fee_amount,omitempty" xml:"market_deduct_fee_amount,omitempty"`
	TotalRefundAmount             *int64                                                                                                    `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
	PlatformMarketRefundAmount    *int64                                                                                                    `json:"platform_market_refund_amount,omitempty" xml:"platform_market_refund_amount,omitempty"`
	Code                          *string                                                                                                   `json:"code,omitempty" xml:"code,omitempty"`
	UserRefundAmount              *int64                                                                                                    `json:"user_refund_amount,omitempty" xml:"user_refund_amount,omitempty"`
	OrderItemId                   *string                                                                                                   `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetCertificateId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRealRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RealRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetTimesCardRefundInfoList(v []*AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.TimesCardRefundInfoList = v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetPlatformMarketDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.PlatformMarketDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetMarketRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.MarketRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetUserDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.UserDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.DeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRefundId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RefundId = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRefundStatus(v int) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RefundStatus = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetMarketDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.MarketDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetTotalRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.TotalRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetPlatformMarketRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.PlatformMarketRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetCode(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.Code = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetUserRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.UserRefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem) SetOrderItemId(v string) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.OrderItemId = &v
	return s
}

type AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem struct {
	RefundAmount                  *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	DeductFeeAmount               *int64 `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
	MarketingAmount               *int64 `json:"marketing_amount,omitempty" xml:"marketing_amount,omitempty"`
	PlatformMarketAmount          *int64 `json:"platform_market_amount,omitempty" xml:"platform_market_amount,omitempty"`
	PlatformMarketDeductFeeAmount *int64 `json:"platform_market_deduct_fee_amount,omitempty" xml:"platform_market_deduct_fee_amount,omitempty"`
	OriginAmount                  *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	SerialNumber                  *int32 `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	UserAmount                    *int64 `json:"user_amount,omitempty" xml:"user_amount,omitempty"`
	UserDeductFeeAmount           *int64 `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetRefundAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.RefundAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.DeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetMarketingAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.MarketingAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetPlatformMarketAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.PlatformMarketAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetPlatformMarketDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.PlatformMarketDeductFeeAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetOriginAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.OriginAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetSerialNumber(v int32) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.SerialNumber = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetUserAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.UserAmount = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetUserDeductFeeAmount(v int64) *AfterSaleOrderDetailGetResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.UserDeductFeeAmount = &v
	return s
}

type AfterSaleOrderDetailGetResponseExtra struct {
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s AfterSaleOrderDetailGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AfterSaleOrderDetailGetResponseExtra) GoString() string {
	return s.String()
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetSubErrorCode(v int32) *AfterSaleOrderDetailGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetDescription(v string) *AfterSaleOrderDetailGetResponseExtra {
	s.Description = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetErrorCode(v int32) *AfterSaleOrderDetailGetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetLogid(v string) *AfterSaleOrderDetailGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetNow(v int64) *AfterSaleOrderDetailGetResponseExtra {
	s.Now = &v
	return s
}

func (s *AfterSaleOrderDetailGetResponseExtra) SetSubDescription(v string) *AfterSaleOrderDetailGetResponseExtra {
	s.SubDescription = &v
	return s
}

type AgencyQueryBillLinkRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	BillDate    *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
}

func (s AgencyQueryBillLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryBillLinkRequest) GoString() string {
	return s.String()
}

func (s *AgencyQueryBillLinkRequest) SetHeader(v map[string]*string) *AgencyQueryBillLinkRequest {
	s.Header = v
	return s
}

func (s *AgencyQueryBillLinkRequest) SetAccessToken(v string) *AgencyQueryBillLinkRequest {
	s.AccessToken = &v
	return s
}

func (s *AgencyQueryBillLinkRequest) SetBillDate(v string) *AgencyQueryBillLinkRequest {
	s.BillDate = &v
	return s
}

type AgencyQueryBillLinkResponse struct {
	Data   *AgencyQueryBillLinkResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s AgencyQueryBillLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryBillLinkResponse) GoString() string {
	return s.String()
}

func (s *AgencyQueryBillLinkResponse) SetData(v *AgencyQueryBillLinkResponseData) *AgencyQueryBillLinkResponse {
	s.Data = v
	return s
}

func (s *AgencyQueryBillLinkResponse) SetErrNo(v int32) *AgencyQueryBillLinkResponse {
	s.ErrNo = &v
	return s
}

func (s *AgencyQueryBillLinkResponse) SetErrMsg(v string) *AgencyQueryBillLinkResponse {
	s.ErrMsg = &v
	return s
}

func (s *AgencyQueryBillLinkResponse) SetLogId(v string) *AgencyQueryBillLinkResponse {
	s.LogId = &v
	return s
}

type AgencyQueryBillLinkResponseData struct {
	LiveBillLink *string `json:"live_bill_link,omitempty" xml:"live_bill_link,omitempty"`
	BillLink     *string `json:"bill_link,omitempty" xml:"bill_link,omitempty"`
}

func (s AgencyQueryBillLinkResponseData) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryBillLinkResponseData) GoString() string {
	return s.String()
}

func (s *AgencyQueryBillLinkResponseData) SetLiveBillLink(v string) *AgencyQueryBillLinkResponseData {
	s.LiveBillLink = &v
	return s
}

func (s *AgencyQueryBillLinkResponseData) SetBillLink(v string) *AgencyQueryBillLinkResponseData {
	s.BillLink = &v
	return s
}

type AgencyQueryVideoSumDataRequest struct {
	VideoIds    []*int64           `json:"video_ids,omitempty" xml:"video_ids,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AgencyQueryVideoSumDataRequest) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryVideoSumDataRequest) GoString() string {
	return s.String()
}

func (s *AgencyQueryVideoSumDataRequest) SetVideoIds(v []*int64) *AgencyQueryVideoSumDataRequest {
	s.VideoIds = v
	return s
}

func (s *AgencyQueryVideoSumDataRequest) SetHeader(v map[string]*string) *AgencyQueryVideoSumDataRequest {
	s.Header = v
	return s
}

func (s *AgencyQueryVideoSumDataRequest) SetAccessToken(v string) *AgencyQueryVideoSumDataRequest {
	s.AccessToken = &v
	return s
}

type AgencyQueryVideoSumDataResponse struct {
	LogId  *string                              `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AgencyQueryVideoSumDataResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                               `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                              `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s AgencyQueryVideoSumDataResponse) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryVideoSumDataResponse) GoString() string {
	return s.String()
}

func (s *AgencyQueryVideoSumDataResponse) SetLogId(v string) *AgencyQueryVideoSumDataResponse {
	s.LogId = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponse) SetData(v *AgencyQueryVideoSumDataResponseData) *AgencyQueryVideoSumDataResponse {
	s.Data = v
	return s
}

func (s *AgencyQueryVideoSumDataResponse) SetErrNo(v int32) *AgencyQueryVideoSumDataResponse {
	s.ErrNo = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponse) SetErrMsg(v string) *AgencyQueryVideoSumDataResponse {
	s.ErrMsg = &v
	return s
}

type AgencyQueryVideoSumDataResponseData struct {
	Data []*AgencyQueryVideoSumDataResponseDataDataItem `json:"data,omitempty" xml:"data,omitempty" require:"true" type:"Repeated"`
}

func (s AgencyQueryVideoSumDataResponseData) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryVideoSumDataResponseData) GoString() string {
	return s.String()
}

func (s *AgencyQueryVideoSumDataResponseData) SetData(v []*AgencyQueryVideoSumDataResponseDataDataItem) *AgencyQueryVideoSumDataResponseData {
	s.Data = v
	return s
}

type AgencyQueryVideoSumDataResponseDataDataItem struct {
	AgentId             *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	PublishTime         *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty" require:"true"`
	BillingGmvTd        *int64  `json:"billing_gmv_td,omitempty" xml:"billing_gmv_td,omitempty" require:"true"`
	BilingRefundGmvTd   *int64  `json:"biling_refund_gmv_td,omitempty" xml:"biling_refund_gmv_td,omitempty" require:"true"`
	StarServiceProvider *string `json:"star_service_provider,omitempty" xml:"star_service_provider,omitempty"`
	CapitalAccount      *int    `json:"capital_account,omitempty" xml:"capital_account,omitempty"`
	VideoId             *int64  `json:"video_id,omitempty" xml:"video_id,omitempty" require:"true"`
	TaskId              *int64  `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
}

func (s AgencyQueryVideoSumDataResponseDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s AgencyQueryVideoSumDataResponseDataDataItem) GoString() string {
	return s.String()
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetAgentId(v string) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.AgentId = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetPublishTime(v int64) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.PublishTime = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetBillingGmvTd(v int64) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.BillingGmvTd = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetBilingRefundGmvTd(v int64) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.BilingRefundGmvTd = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetStarServiceProvider(v string) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.StarServiceProvider = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetCapitalAccount(v int) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.CapitalAccount = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetVideoId(v int64) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.VideoId = &v
	return s
}

func (s *AgencyQueryVideoSumDataResponseDataDataItem) SetTaskId(v int64) *AgencyQueryVideoSumDataResponseDataDataItem {
	s.TaskId = &v
	return s
}

type AkteAfterSaleOrderQueryRequest struct {
	RefundDoneEndTime    *int64             `json:"refund_done_end_time,omitempty" xml:"refund_done_end_time,omitempty"`
	RefundDoneStartTime  *int64             `json:"refund_done_start_time,omitempty" xml:"refund_done_start_time,omitempty"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AccountId            *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	CreateOrderEndTime   *int64             `json:"create_order_end_time,omitempty" xml:"create_order_end_time,omitempty"`
	CreateOrderStartTime *int64             `json:"create_order_start_time,omitempty" xml:"create_order_start_time,omitempty"`
	Cursor               *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	RefundStatus         *int               `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	PageSize             *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
}

func (s AkteAfterSaleOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryRequest) SetRefundDoneEndTime(v int64) *AkteAfterSaleOrderQueryRequest {
	s.RefundDoneEndTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetRefundDoneStartTime(v int64) *AkteAfterSaleOrderQueryRequest {
	s.RefundDoneStartTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetHeader(v map[string]*string) *AkteAfterSaleOrderQueryRequest {
	s.Header = v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetAccessToken(v string) *AkteAfterSaleOrderQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetAccountId(v string) *AkteAfterSaleOrderQueryRequest {
	s.AccountId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetCreateOrderEndTime(v int64) *AkteAfterSaleOrderQueryRequest {
	s.CreateOrderEndTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetCreateOrderStartTime(v int64) *AkteAfterSaleOrderQueryRequest {
	s.CreateOrderStartTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetCursor(v string) *AkteAfterSaleOrderQueryRequest {
	s.Cursor = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetRefundStatus(v int) *AkteAfterSaleOrderQueryRequest {
	s.RefundStatus = &v
	return s
}

func (s *AkteAfterSaleOrderQueryRequest) SetPageSize(v int32) *AkteAfterSaleOrderQueryRequest {
	s.PageSize = &v
	return s
}

type AkteAfterSaleOrderQueryResponse struct {
	Extra *AkteAfterSaleOrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *AkteAfterSaleOrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponse) SetExtra(v *AkteAfterSaleOrderQueryResponseExtra) *AkteAfterSaleOrderQueryResponse {
	s.Extra = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponse) SetData(v *AkteAfterSaleOrderQueryResponseData) *AkteAfterSaleOrderQueryResponse {
	s.Data = v
	return s
}

type AkteAfterSaleOrderQueryResponseData struct {
	HasMore            *bool                                                        `json:"has_more,omitempty" xml:"has_more,omitempty"`
	GwErrorCode        *int32                                                       `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription      *string                                                      `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	AfterSaleOrderList []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem `json:"after_sale_order_list,omitempty" xml:"after_sale_order_list,omitempty" type:"Repeated"`
	Cursor             *string                                                      `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseData) SetHasMore(v bool) *AkteAfterSaleOrderQueryResponseData {
	s.HasMore = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseData) SetGwErrorCode(v int32) *AkteAfterSaleOrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseData) SetGwDescription(v string) *AkteAfterSaleOrderQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseData) SetAfterSaleOrderList(v []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) *AkteAfterSaleOrderQueryResponseData {
	s.AfterSaleOrderList = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseData) SetCursor(v string) *AkteAfterSaleOrderQueryResponseData {
	s.Cursor = &v
	return s
}

type AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem struct {
	RefundType            *int64                                                                         `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	MarketRefundAmount    *int64                                                                         `json:"market_refund_amount,omitempty" xml:"market_refund_amount,omitempty"`
	OutBizAfterSaleId     *string                                                                        `json:"out_biz_after_sale_id,omitempty" xml:"out_biz_after_sale_id,omitempty"`
	CreateTime            *int64                                                                         `json:"create_time,omitempty" xml:"create_time,omitempty"`
	TradeType             *int64                                                                         `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	RefundInfoList        []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem `json:"refund_info_list,omitempty" xml:"refund_info_list,omitempty" type:"Repeated"`
	RefundAmount          *int64                                                                         `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	AuditTime             *int64                                                                         `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
	AfterSaleId           *string                                                                        `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty" require:"true"`
	OutRefundPaymentId    *string                                                                        `json:"out_refund_payment_id,omitempty" xml:"out_refund_payment_id,omitempty"`
	MerchantAccountId     *int64                                                                         `json:"merchant_account_id,omitempty" xml:"merchant_account_id,omitempty"`
	Status                *int                                                                           `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	TotalRefundAmount     *int64                                                                         `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
	UserRefundAmount      *int64                                                                         `json:"user_refund_amount,omitempty" xml:"user_refund_amount,omitempty"`
	OrderId               *string                                                                        `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Reason                *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason               `json:"reason,omitempty" xml:"reason,omitempty"`
	RealRefundAmount      *int64                                                                         `json:"real_refund_amount,omitempty" xml:"real_refund_amount,omitempty"`
	OrderType             *int64                                                                         `json:"order_type,omitempty" xml:"order_type,omitempty"`
	RejectReason          *string                                                                        `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	MarketDeductFeeAmount *int64                                                                         `json:"market_deduct_fee_amount,omitempty" xml:"market_deduct_fee_amount,omitempty"`
	CompleteTime          *int64                                                                         `json:"complete_time,omitempty" xml:"complete_time,omitempty"`
	UserDeductFeeAmount   *int64                                                                         `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
	AuditResult           *string                                                                        `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	UpdateTime            *int64                                                                         `json:"update_time,omitempty" xml:"update_time,omitempty"`
	DeductFeeAmount       *int64                                                                         `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetRefundType(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.RefundType = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetMarketRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.MarketRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetOutBizAfterSaleId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.OutBizAfterSaleId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetCreateTime(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.CreateTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetTradeType(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.TradeType = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetRefundInfoList(v []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.RefundInfoList = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.RefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetAuditTime(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.AuditTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetAfterSaleId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.AfterSaleId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetOutRefundPaymentId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.OutRefundPaymentId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetMerchantAccountId(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.MerchantAccountId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetStatus(v int) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.Status = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetTotalRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.TotalRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetUserRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.UserRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetOrderId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.OrderId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetReason(v *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.Reason = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetRealRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.RealRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetOrderType(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.OrderType = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetRejectReason(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.RejectReason = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetMarketDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.MarketDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetCompleteTime(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.CompleteTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetUserDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.UserDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetAuditResult(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.AuditResult = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetUpdateTime(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.UpdateTime = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem) SetDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItem {
	s.DeductFeeAmount = &v
	return s
}

type AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason struct {
	ReasonCode []*int64                                                                         `json:"reason_code,omitempty" xml:"reason_code,omitempty" type:"Repeated"`
	ShowReason []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem `json:"show_reason,omitempty" xml:"show_reason,omitempty" type:"Repeated"`
	Desc       *string                                                                          `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) SetReasonCode(v []*int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason {
	s.ReasonCode = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) SetShowReason(v []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason {
	s.ShowReason = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason) SetDesc(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReason {
	s.Desc = &v
	return s
}

type AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem struct {
	Msg        *string `json:"msg,omitempty" xml:"msg,omitempty"`
	ReasonCode *int64  `json:"reason_code,omitempty" xml:"reason_code,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem) SetMsg(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem {
	s.Msg = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem) SetReasonCode(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemReasonShowReasonItem {
	s.ReasonCode = &v
	return s
}

type AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem struct {
	TotalRefundAmount             *int64                                                                                                    `json:"total_refund_amount,omitempty" xml:"total_refund_amount,omitempty"`
	DeductFeeAmount               *int64                                                                                                    `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
	Code                          *string                                                                                                   `json:"code,omitempty" xml:"code,omitempty"`
	RealRefundAmount              *int64                                                                                                    `json:"real_refund_amount,omitempty" xml:"real_refund_amount,omitempty"`
	RefundStatus                  *int                                                                                                      `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	TimesCardRefundInfoList       []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem `json:"times_card_refund_info_list,omitempty" xml:"times_card_refund_info_list,omitempty" type:"Repeated"`
	CertificateId                 *string                                                                                                   `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	PlatformMarketRefundAmount    *int64                                                                                                    `json:"platform_market_refund_amount,omitempty" xml:"platform_market_refund_amount,omitempty"`
	RefundId                      *string                                                                                                   `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	PlatformMarketDeductFeeAmount *int64                                                                                                    `json:"platform_market_deduct_fee_amount,omitempty" xml:"platform_market_deduct_fee_amount,omitempty"`
	UserRefundAmount              *int64                                                                                                    `json:"user_refund_amount,omitempty" xml:"user_refund_amount,omitempty"`
	MarketRefundAmount            *int64                                                                                                    `json:"market_refund_amount,omitempty" xml:"market_refund_amount,omitempty"`
	OrderItemId                   *string                                                                                                   `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	MarketDeductFeeAmount         *int64                                                                                                    `json:"market_deduct_fee_amount,omitempty" xml:"market_deduct_fee_amount,omitempty"`
	UserDeductFeeAmount           *int64                                                                                                    `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetTotalRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.TotalRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.DeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetCode(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.Code = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRealRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RealRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRefundStatus(v int) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RefundStatus = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetTimesCardRefundInfoList(v []*AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.TimesCardRefundInfoList = v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetCertificateId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.CertificateId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetPlatformMarketRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.PlatformMarketRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetRefundId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.RefundId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetPlatformMarketDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.PlatformMarketDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetUserRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.UserRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetMarketRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.MarketRefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetOrderItemId(v string) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.OrderItemId = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetMarketDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.MarketDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem) SetUserDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItem {
	s.UserDeductFeeAmount = &v
	return s
}

type AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem struct {
	PlatformMarketAmount          *int64 `json:"platform_market_amount,omitempty" xml:"platform_market_amount,omitempty"`
	SerialNumber                  *int32 `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	OriginAmount                  *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	UserAmount                    *int64 `json:"user_amount,omitempty" xml:"user_amount,omitempty"`
	PlatformMarketDeductFeeAmount *int64 `json:"platform_market_deduct_fee_amount,omitempty" xml:"platform_market_deduct_fee_amount,omitempty"`
	MarketingAmount               *int64 `json:"marketing_amount,omitempty" xml:"marketing_amount,omitempty"`
	UserDeductFeeAmount           *int64 `json:"user_deduct_fee_amount,omitempty" xml:"user_deduct_fee_amount,omitempty"`
	RefundAmount                  *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	DeductFeeAmount               *int64 `json:"deduct_fee_amount,omitempty" xml:"deduct_fee_amount,omitempty"`
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetPlatformMarketAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.PlatformMarketAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetSerialNumber(v int32) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.SerialNumber = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetOriginAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.OriginAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetUserAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.UserAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetPlatformMarketDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.PlatformMarketDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetMarketingAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.MarketingAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetUserDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.UserDeductFeeAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetRefundAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.RefundAmount = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem) SetDeductFeeAmount(v int64) *AkteAfterSaleOrderQueryResponseDataAfterSaleOrderListItemRefundInfoListItemTimesCardRefundInfoListItem {
	s.DeductFeeAmount = &v
	return s
}

type AkteAfterSaleOrderQueryResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
}

func (s AkteAfterSaleOrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AkteAfterSaleOrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetSubDescription(v string) *AkteAfterSaleOrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetSubErrorCode(v int32) *AkteAfterSaleOrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetDescription(v string) *AkteAfterSaleOrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetErrorCode(v int32) *AkteAfterSaleOrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetLogid(v string) *AkteAfterSaleOrderQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *AkteAfterSaleOrderQueryResponseExtra) SetNow(v int64) *AkteAfterSaleOrderQueryResponseExtra {
	s.Now = &v
	return s
}

type AkteCommentReplyRequest struct {
	PoiId       *int64             `json:"poi_id,omitempty" xml:"poi_id,omitempty" require:"true"`
	RateId      *int64             `json:"rate_id,omitempty" xml:"rate_id,omitempty" require:"true"`
	Text        *string            `json:"text,omitempty" xml:"text,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AkteCommentReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s AkteCommentReplyRequest) GoString() string {
	return s.String()
}

func (s *AkteCommentReplyRequest) SetPoiId(v int64) *AkteCommentReplyRequest {
	s.PoiId = &v
	return s
}

func (s *AkteCommentReplyRequest) SetRateId(v int64) *AkteCommentReplyRequest {
	s.RateId = &v
	return s
}

func (s *AkteCommentReplyRequest) SetText(v string) *AkteCommentReplyRequest {
	s.Text = &v
	return s
}

func (s *AkteCommentReplyRequest) SetAccountId(v string) *AkteCommentReplyRequest {
	s.AccountId = &v
	return s
}

func (s *AkteCommentReplyRequest) SetHeader(v map[string]*string) *AkteCommentReplyRequest {
	s.Header = v
	return s
}

func (s *AkteCommentReplyRequest) SetAccessToken(v string) *AkteCommentReplyRequest {
	s.AccessToken = &v
	return s
}

type AkteCommentReplyResponse struct {
	Extra *AkteCommentReplyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *AkteCommentReplyResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AkteCommentReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s AkteCommentReplyResponse) GoString() string {
	return s.String()
}

func (s *AkteCommentReplyResponse) SetExtra(v *AkteCommentReplyResponseExtra) *AkteCommentReplyResponse {
	s.Extra = v
	return s
}

func (s *AkteCommentReplyResponse) SetData(v *AkteCommentReplyResponseData) *AkteCommentReplyResponse {
	s.Data = v
	return s
}

type AkteCommentReplyResponseData struct {
	ReplyId       *int64  `json:"reply_id,omitempty" xml:"reply_id,omitempty"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AkteCommentReplyResponseData) String() string {
	return tea.Prettify(s)
}

func (s AkteCommentReplyResponseData) GoString() string {
	return s.String()
}

func (s *AkteCommentReplyResponseData) SetReplyId(v int64) *AkteCommentReplyResponseData {
	s.ReplyId = &v
	return s
}

func (s *AkteCommentReplyResponseData) SetGwErrorCode(v int32) *AkteCommentReplyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AkteCommentReplyResponseData) SetGwDescription(v string) *AkteCommentReplyResponseData {
	s.GwDescription = &v
	return s
}

type AkteCommentReplyResponseExtra struct {
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AkteCommentReplyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AkteCommentReplyResponseExtra) GoString() string {
	return s.String()
}

func (s *AkteCommentReplyResponseExtra) SetErrorCode(v int32) *AkteCommentReplyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AkteCommentReplyResponseExtra) SetLogid(v string) *AkteCommentReplyResponseExtra {
	s.Logid = &v
	return s
}

func (s *AkteCommentReplyResponseExtra) SetNow(v int64) *AkteCommentReplyResponseExtra {
	s.Now = &v
	return s
}

func (s *AkteCommentReplyResponseExtra) SetSubDescription(v string) *AkteCommentReplyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AkteCommentReplyResponseExtra) SetSubErrorCode(v int32) *AkteCommentReplyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AkteCommentReplyResponseExtra) SetDescription(v string) *AkteCommentReplyResponseExtra {
	s.Description = &v
	return s
}

type AkteOrderQueryRequest struct {
	OrderId              *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	ExtOrderId           *string            `json:"ext_order_id,omitempty" xml:"ext_order_id,omitempty"`
	CreateOrderEndTime   *int64             `json:"create_order_end_time,omitempty" xml:"create_order_end_time,omitempty"`
	Cursor               []*string          `json:"cursor,omitempty" xml:"cursor,omitempty" type:"Repeated"`
	UpdateOrderStartTime *int64             `json:"update_order_start_time,omitempty" xml:"update_order_start_time,omitempty"`
	CreateOrderStartTime *int64             `json:"create_order_start_time,omitempty" xml:"create_order_start_time,omitempty"`
	OrderStatus          *int32             `json:"order_status,omitempty" xml:"order_status,omitempty"`
	PageSize             *int32             `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	OpenId               *string            `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AccountId            *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	UpdateOrderEndTime   *int64             `json:"update_order_end_time,omitempty" xml:"update_order_end_time,omitempty"`
	PageNum              *int32             `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AkteOrderQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryRequest) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryRequest) SetOrderId(v string) *AkteOrderQueryRequest {
	s.OrderId = &v
	return s
}

func (s *AkteOrderQueryRequest) SetExtOrderId(v string) *AkteOrderQueryRequest {
	s.ExtOrderId = &v
	return s
}

func (s *AkteOrderQueryRequest) SetCreateOrderEndTime(v int64) *AkteOrderQueryRequest {
	s.CreateOrderEndTime = &v
	return s
}

func (s *AkteOrderQueryRequest) SetCursor(v []*string) *AkteOrderQueryRequest {
	s.Cursor = v
	return s
}

func (s *AkteOrderQueryRequest) SetUpdateOrderStartTime(v int64) *AkteOrderQueryRequest {
	s.UpdateOrderStartTime = &v
	return s
}

func (s *AkteOrderQueryRequest) SetCreateOrderStartTime(v int64) *AkteOrderQueryRequest {
	s.CreateOrderStartTime = &v
	return s
}

func (s *AkteOrderQueryRequest) SetOrderStatus(v int32) *AkteOrderQueryRequest {
	s.OrderStatus = &v
	return s
}

func (s *AkteOrderQueryRequest) SetPageSize(v int32) *AkteOrderQueryRequest {
	s.PageSize = &v
	return s
}

func (s *AkteOrderQueryRequest) SetOpenId(v string) *AkteOrderQueryRequest {
	s.OpenId = &v
	return s
}

func (s *AkteOrderQueryRequest) SetAccountId(v string) *AkteOrderQueryRequest {
	s.AccountId = &v
	return s
}

func (s *AkteOrderQueryRequest) SetAccessToken(v string) *AkteOrderQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *AkteOrderQueryRequest) SetUpdateOrderEndTime(v int64) *AkteOrderQueryRequest {
	s.UpdateOrderEndTime = &v
	return s
}

func (s *AkteOrderQueryRequest) SetPageNum(v int32) *AkteOrderQueryRequest {
	s.PageNum = &v
	return s
}

func (s *AkteOrderQueryRequest) SetHeader(v map[string]*string) *AkteOrderQueryRequest {
	s.Header = v
	return s
}

type AkteOrderQueryResponse struct {
	Data  *AkteOrderQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *AkteOrderQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AkteOrderQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponse) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponse) SetData(v *AkteOrderQueryResponseData) *AkteOrderQueryResponse {
	s.Data = v
	return s
}

func (s *AkteOrderQueryResponse) SetExtra(v *AkteOrderQueryResponseExtra) *AkteOrderQueryResponse {
	s.Extra = v
	return s
}

type AkteOrderQueryResponseData struct {
	Page          *AkteOrderQueryResponseDataPage         `json:"page,omitempty" xml:"page,omitempty"`
	SearchAfter   *AkteOrderQueryResponseDataSearchAfter  `json:"search_after,omitempty" xml:"search_after,omitempty"`
	GwErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Orders        []*AkteOrderQueryResponseDataOrdersItem `json:"orders,omitempty" xml:"orders,omitempty" type:"Repeated"`
}

func (s AkteOrderQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseData) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseData) SetPage(v *AkteOrderQueryResponseDataPage) *AkteOrderQueryResponseData {
	s.Page = v
	return s
}

func (s *AkteOrderQueryResponseData) SetSearchAfter(v *AkteOrderQueryResponseDataSearchAfter) *AkteOrderQueryResponseData {
	s.SearchAfter = v
	return s
}

func (s *AkteOrderQueryResponseData) SetGwErrorCode(v int32) *AkteOrderQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AkteOrderQueryResponseData) SetGwDescription(v string) *AkteOrderQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *AkteOrderQueryResponseData) SetOrders(v []*AkteOrderQueryResponseDataOrdersItem) *AkteOrderQueryResponseData {
	s.Orders = v
	return s
}

type AkteOrderQueryResponseDataOrdersItem struct {
	DiscountAmount      *int64                                                         `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	Certificate         []*AkteOrderQueryResponseDataOrdersItemCertificateItem         `json:"certificate,omitempty" xml:"certificate,omitempty" type:"Repeated"`
	AttributeInfo       *AkteOrderQueryResponseDataOrdersItemAttributeInfo             `json:"attribute_info,omitempty" xml:"attribute_info,omitempty"`
	Poi                 *AkteOrderQueryResponseDataOrdersItemPoi                       `json:"poi,omitempty" xml:"poi,omitempty"`
	MerchantInfo        *AkteOrderQueryResponseDataOrdersItemMerchantInfo              `json:"merchant_info,omitempty" xml:"merchant_info,omitempty"`
	OrderStatus         *int32                                                         `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OpenId              *string                                                        `json:"open_id,omitempty" xml:"open_id,omitempty"`
	PaymentDiscount     *int32                                                         `json:"payment_discount,omitempty" xml:"payment_discount,omitempty"`
	HasDonatedItems     *bool                                                          `json:"has_donated_items,omitempty" xml:"has_donated_items,omitempty"`
	UpdateOrderTime     *int64                                                         `json:"update_order_time,omitempty" xml:"update_order_time,omitempty"`
	Contacts            []*AkteOrderQueryResponseDataOrdersItemContactsItem            `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	Goods               []*AkteOrderQueryResponseDataOrdersItemGoodsItem               `json:"goods,omitempty" xml:"goods,omitempty" type:"Repeated"`
	OrderType           *int32                                                         `json:"order_type,omitempty" xml:"order_type,omitempty"`
	PayAmount           *int32                                                         `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	OrderId             *string                                                        `json:"order_id,omitempty" xml:"order_id,omitempty" require:"true"`
	CreateOrderTime     *int64                                                         `json:"create_order_time,omitempty" xml:"create_order_time,omitempty"`
	OrderPayAmount      *int32                                                         `json:"order_pay_amount,omitempty" xml:"order_pay_amount,omitempty"`
	AmountInfo          *AkteOrderQueryResponseDataOrdersItemAmountInfo                `json:"amount_info,omitempty" xml:"amount_info,omitempty"`
	OriginalAmount      *int32                                                         `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	PayTime             *int64                                                         `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	SubOrderAmountInfos []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem `json:"sub_order_amount_infos,omitempty" xml:"sub_order_amount_infos,omitempty" type:"Repeated"`
	Discounts           []*AkteOrderQueryResponseDataOrdersItemDiscountsItem           `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
}

func (s AkteOrderQueryResponseDataOrdersItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItem {
	s.DiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetCertificate(v []*AkteOrderQueryResponseDataOrdersItemCertificateItem) *AkteOrderQueryResponseDataOrdersItem {
	s.Certificate = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetAttributeInfo(v *AkteOrderQueryResponseDataOrdersItemAttributeInfo) *AkteOrderQueryResponseDataOrdersItem {
	s.AttributeInfo = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetPoi(v *AkteOrderQueryResponseDataOrdersItemPoi) *AkteOrderQueryResponseDataOrdersItem {
	s.Poi = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetMerchantInfo(v *AkteOrderQueryResponseDataOrdersItemMerchantInfo) *AkteOrderQueryResponseDataOrdersItem {
	s.MerchantInfo = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOrderStatus(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.OrderStatus = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOpenId(v string) *AkteOrderQueryResponseDataOrdersItem {
	s.OpenId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetPaymentDiscount(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.PaymentDiscount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetHasDonatedItems(v bool) *AkteOrderQueryResponseDataOrdersItem {
	s.HasDonatedItems = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetUpdateOrderTime(v int64) *AkteOrderQueryResponseDataOrdersItem {
	s.UpdateOrderTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetContacts(v []*AkteOrderQueryResponseDataOrdersItemContactsItem) *AkteOrderQueryResponseDataOrdersItem {
	s.Contacts = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetGoods(v []*AkteOrderQueryResponseDataOrdersItemGoodsItem) *AkteOrderQueryResponseDataOrdersItem {
	s.Goods = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOrderType(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.OrderType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetPayAmount(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.PayAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOrderId(v string) *AkteOrderQueryResponseDataOrdersItem {
	s.OrderId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetCreateOrderTime(v int64) *AkteOrderQueryResponseDataOrdersItem {
	s.CreateOrderTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOrderPayAmount(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.OrderPayAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetAmountInfo(v *AkteOrderQueryResponseDataOrdersItemAmountInfo) *AkteOrderQueryResponseDataOrdersItem {
	s.AmountInfo = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetOriginalAmount(v int32) *AkteOrderQueryResponseDataOrdersItem {
	s.OriginalAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetPayTime(v int64) *AkteOrderQueryResponseDataOrdersItem {
	s.PayTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetSubOrderAmountInfos(v []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) *AkteOrderQueryResponseDataOrdersItem {
	s.SubOrderAmountInfos = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItem) SetDiscounts(v []*AkteOrderQueryResponseDataOrdersItemDiscountsItem) *AkteOrderQueryResponseDataOrdersItem {
	s.Discounts = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemAmountInfo struct {
	PayDiscountAmount      *int64 `json:"pay_discount_amount,omitempty" xml:"pay_discount_amount,omitempty"`
	PlatformDiscountAmount *int64 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	BrandDiscountAmount    *int64 `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
	MerchantDiscountAmount *int64 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	OrderPayAmount         *int64 `json:"order_pay_amount,omitempty" xml:"order_pay_amount,omitempty"`
	OriginAmount           *int64 `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	PayAmount              *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemAmountInfo) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemAmountInfo) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetPayDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetPlatformDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetBrandDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.BrandDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetMerchantDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetOrderPayAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.OrderPayAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetOriginAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.OriginAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAmountInfo) SetPayAmount(v int64) *AkteOrderQueryResponseDataOrdersItemAmountInfo {
	s.PayAmount = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemAttributeInfo struct {
	RoomId   *int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AnchorId *int64 `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemAttributeInfo) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemAttributeInfo) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemAttributeInfo) SetRoomId(v int64) *AkteOrderQueryResponseDataOrdersItemAttributeInfo {
	s.RoomId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemAttributeInfo) SetAnchorId(v int64) *AkteOrderQueryResponseDataOrdersItemAttributeInfo {
	s.AnchorId = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemCertificateItem struct {
	ItemStatus          *int32  `json:"item_status,omitempty" xml:"item_status,omitempty"`
	OriginCertificateId *string `json:"origin_certificate_id,omitempty" xml:"origin_certificate_id,omitempty"`
	IsDonatedItem       *bool   `json:"is_donated_item,omitempty" xml:"is_donated_item,omitempty"`
	RefundTime          *int64  `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	OrderItemId         *string `json:"order_item_id,omitempty" xml:"order_item_id,omitempty"`
	SubSkuId            *string `json:"sub_sku_id,omitempty" xml:"sub_sku_id,omitempty"`
	ItemUpdateTime      *int64  `json:"item_update_time,omitempty" xml:"item_update_time,omitempty"`
	SkuId               *string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	CertificateId       *string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	CombinationId       *string `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
	RefundAmount        *int32  `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemCertificateItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemCertificateItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetItemStatus(v int32) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemStatus = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetOriginCertificateId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.OriginCertificateId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetIsDonatedItem(v bool) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.IsDonatedItem = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetRefundTime(v int64) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetOrderItemId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.OrderItemId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetSubSkuId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.SubSkuId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetItemUpdateTime(v int64) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.ItemUpdateTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetSkuId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.SkuId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetCertificateId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.CertificateId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetCombinationId(v string) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.CombinationId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemCertificateItem) SetRefundAmount(v int32) *AkteOrderQueryResponseDataOrdersItemCertificateItem {
	s.RefundAmount = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemContactsItem struct {
	PhoneEncrypt *string `json:"phone_encrypt,omitempty" xml:"phone_encrypt,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone        *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemContactsItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemContactsItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemContactsItem) SetPhoneEncrypt(v string) *AkteOrderQueryResponseDataOrdersItemContactsItem {
	s.PhoneEncrypt = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemContactsItem) SetName(v string) *AkteOrderQueryResponseDataOrdersItemContactsItem {
	s.Name = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemContactsItem) SetPhone(v string) *AkteOrderQueryResponseDataOrdersItemContactsItem {
	s.Phone = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemDiscountsItem struct {
	DiscountType           *int64                                                          `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	MerchantDiscountAmount *int64                                                          `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64                                                          `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	BrandDiscountAmount    *int64                                                          `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
	DiscountAmount         *int64                                                          `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountExtra          *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra `json:"discount_extra,omitempty" xml:"discount_extra,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountType(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetMerchantDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetPlatformDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetBrandDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.BrandDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItem) SetDiscountExtra(v *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra) *AkteOrderQueryResponseDataOrdersItemDiscountsItem {
	s.DiscountExtra = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra struct {
	IdleTimeDiscountInfo *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo `json:"idle_time_discount_info,omitempty" xml:"idle_time_discount_info,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra) SetIdleTimeDiscountInfo(v *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtra {
	s.IdleTimeDiscountInfo = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo struct {
	WeekDayList       []*int                                                                                                  `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
	DailyTimeRange    []*AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem `json:"daily_time_range,omitempty" xml:"daily_time_range,omitempty" type:"Repeated"`
	DiscountEndTime   *int64                                                                                                  `json:"discount_end_time,omitempty" xml:"discount_end_time,omitempty"`
	DiscountStartTime *int64                                                                                                  `json:"discount_start_time,omitempty" xml:"discount_start_time,omitempty"`
	IdleTimeLimitType *int                                                                                                    `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetWeekDayList(v []*int) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.WeekDayList = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDailyTimeRange(v []*AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DailyTimeRange = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDiscountEndTime(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DiscountEndTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDiscountStartTime(v int64) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DiscountStartTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetIdleTimeLimitType(v int) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem struct {
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) SetStartTime(v string) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem {
	s.StartTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) SetEndTime(v string) *AkteOrderQueryResponseDataOrdersItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem {
	s.EndTime = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItem struct {
	SkuName     *string                                                   `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	ThirdSkuId  *string                                                   `json:"third_sku_id,omitempty" xml:"third_sku_id,omitempty"`
	Count       *int32                                                    `json:"count,omitempty" xml:"count,omitempty"`
	GoodsDetail *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail `json:"goods_detail,omitempty" xml:"goods_detail,omitempty"`
	SkuId       *string                                                   `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItem) SetSkuName(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItem {
	s.SkuName = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItem) SetThirdSkuId(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItem {
	s.ThirdSkuId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItem) SetCount(v int32) *AkteOrderQueryResponseDataOrdersItemGoodsItem {
	s.Count = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItem) SetGoodsDetail(v *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) *AkteOrderQueryResponseDataOrdersItemGoodsItem {
	s.GoodsDetail = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItem) SetSkuId(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItem {
	s.SkuId = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail struct {
	GoodsAttr   *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr `json:"goods_attr,omitempty" xml:"goods_attr,omitempty"`
	ProductAttr map[string]*string                                                 `json:"product_attr,omitempty" xml:"product_attr,omitempty"`
	ProductType *int64                                                             `json:"product_type,omitempty" xml:"product_type,omitempty"`
	ShowChannel *int32                                                             `json:"show_channel,omitempty" xml:"show_channel,omitempty"`
	SkuAttr     map[string]*string                                                 `json:"sku_attr,omitempty" xml:"sku_attr,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) SetGoodsAttr(v *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail {
	s.GoodsAttr = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) SetProductAttr(v map[string]*string) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail {
	s.ProductAttr = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) SetProductType(v int64) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail {
	s.ProductType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) SetShowChannel(v int32) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail {
	s.ShowChannel = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail) SetSkuAttr(v map[string]*string) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetail {
	s.SkuAttr = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr struct {
	DynamicParRule *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule `json:"dynamic_par_rule,omitempty" xml:"dynamic_par_rule,omitempty"`
	VoucherParType *int                                                                             `json:"voucher_par_type,omitempty" xml:"voucher_par_type,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr) SetDynamicParRule(v *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr {
	s.DynamicParRule = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr) SetVoucherParType(v int) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttr {
	s.VoucherParType = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule struct {
	DynamicParRuleDesc     *string                                                                                                      `json:"dynamic_par_rule_desc,omitempty" xml:"dynamic_par_rule_desc,omitempty" require:"true"`
	DynamicParRuleItemList []*AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem `json:"dynamic_par_rule_item_list,omitempty" xml:"dynamic_par_rule_item_list,omitempty" require:"true" type:"Repeated"`
	DynamicParMaxAmount    *int64                                                                                                       `json:"dynamic_par_max_amount,omitempty" xml:"dynamic_par_max_amount,omitempty" require:"true"`
	DynamicParMinAmount    *int64                                                                                                       `json:"dynamic_par_min_amount,omitempty" xml:"dynamic_par_min_amount,omitempty" require:"true"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) SetDynamicParRuleDesc(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule {
	s.DynamicParRuleDesc = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) SetDynamicParRuleItemList(v []*AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule {
	s.DynamicParRuleItemList = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) SetDynamicParMaxAmount(v int64) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule {
	s.DynamicParMaxAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule) SetDynamicParMinAmount(v int64) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRule {
	s.DynamicParMinAmount = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem struct {
	DeductibleAmount *int64                                                                                                               `json:"deductible_amount,omitempty" xml:"deductible_amount,omitempty" require:"true"`
	TimePeriod       *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod `json:"time_period,omitempty" xml:"time_period,omitempty" require:"true"`
	DayOfWeek        []*int                                                                                                               `json:"day_of_week,omitempty" xml:"day_of_week,omitempty" require:"true" type:"Repeated"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) SetDeductibleAmount(v int64) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem {
	s.DeductibleAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) SetTimePeriod(v *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem {
	s.TimePeriod = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem) SetDayOfWeek(v []*int) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItem {
	s.DayOfWeek = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod struct {
	UseEndTime       *string `json:"use_end_time,omitempty" xml:"use_end_time,omitempty" require:"true"`
	UseStartTime     *string `json:"use_start_time,omitempty" xml:"use_start_time,omitempty" require:"true"`
	EndTimeIsNextDay *bool   `json:"end_time_is_next_day,omitempty" xml:"end_time_is_next_day,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetUseEndTime(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.UseEndTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetUseStartTime(v string) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.UseStartTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod) SetEndTimeIsNextDay(v bool) *AkteOrderQueryResponseDataOrdersItemGoodsItemGoodsDetailGoodsAttrDynamicParRuleDynamicParRuleItemListItemTimePeriod {
	s.EndTimeIsNextDay = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemMerchantInfo struct {
	AccountId   *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty" xml:"account_name,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemMerchantInfo) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemMerchantInfo) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountId(v string) *AkteOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemMerchantInfo) SetAccountName(v string) *AkteOrderQueryResponseDataOrdersItemMerchantInfo {
	s.AccountName = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemPoi struct {
	PoiName          *string `json:"poi_name,omitempty" xml:"poi_name,omitempty"`
	IntentionPoiId   *string `json:"intention_poi_id,omitempty" xml:"intention_poi_id,omitempty"`
	IntentionPoiName *string `json:"intention_poi_name,omitempty" xml:"intention_poi_name,omitempty"`
	PoiId            *string `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemPoi) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemPoi) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemPoi) SetPoiName(v string) *AkteOrderQueryResponseDataOrdersItemPoi {
	s.PoiName = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemPoi) SetIntentionPoiId(v string) *AkteOrderQueryResponseDataOrdersItemPoi {
	s.IntentionPoiId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemPoi) SetIntentionPoiName(v string) *AkteOrderQueryResponseDataOrdersItemPoi {
	s.IntentionPoiName = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemPoi) SetPoiId(v string) *AkteOrderQueryResponseDataOrdersItemPoi {
	s.PoiId = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem struct {
	SubSkuId           *string                                                                     `json:"sub_sku_id,omitempty" xml:"sub_sku_id,omitempty"`
	SkuId              *string                                                                     `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	CombinationId      *string                                                                     `json:"combination_id,omitempty" xml:"combination_id,omitempty"`
	DiscountAmount     *int64                                                                      `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	SubOrderType       *int32                                                                      `json:"sub_order_type,omitempty" xml:"sub_order_type,omitempty"`
	OriginAmount       *int64                                                                      `json:"origin_amount,omitempty" xml:"origin_amount,omitempty"`
	SubCouponPayAmount *int64                                                                      `json:"sub_coupon_pay_amount,omitempty" xml:"sub_coupon_pay_amount,omitempty"`
	Discounts          []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem `json:"discounts,omitempty" xml:"discounts,omitempty" type:"Repeated"`
	PayAmount          *int64                                                                      `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	SubOrderId         *string                                                                     `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubSkuId(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubSkuId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSkuId(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SkuId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetCombinationId(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.CombinationId = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.DiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderType(v int32) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetOriginAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.OriginAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubCouponPayAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubCouponPayAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetDiscounts(v []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.Discounts = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetPayAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.PayAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem) SetSubOrderId(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItem {
	s.SubOrderId = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem struct {
	DiscountAmount         *int64                                                                                 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	DiscountExtra          *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra `json:"discount_extra,omitempty" xml:"discount_extra,omitempty"`
	DiscountType           *int64                                                                                 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	MerchantDiscountAmount *int64                                                                                 `json:"merchant_discount_amount,omitempty" xml:"merchant_discount_amount,omitempty"`
	PlatformDiscountAmount *int64                                                                                 `json:"platform_discount_amount,omitempty" xml:"platform_discount_amount,omitempty"`
	BrandDiscountAmount    *int64                                                                                 `json:"brand_discount_amount,omitempty" xml:"brand_discount_amount,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountExtra(v *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountExtra = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetDiscountType(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.DiscountType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetMerchantDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.MerchantDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetPlatformDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.PlatformDiscountAmount = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem) SetBrandDiscountAmount(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItem {
	s.BrandDiscountAmount = &v
	return s
}

type AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra struct {
	IdleTimeDiscountInfo *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo `json:"idle_time_discount_info,omitempty" xml:"idle_time_discount_info,omitempty"`
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra) SetIdleTimeDiscountInfo(v *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtra {
	s.IdleTimeDiscountInfo = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo struct {
	DailyTimeRange    []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem `json:"daily_time_range,omitempty" xml:"daily_time_range,omitempty" type:"Repeated"`
	DiscountEndTime   *int64                                                                                                                         `json:"discount_end_time,omitempty" xml:"discount_end_time,omitempty"`
	DiscountStartTime *int64                                                                                                                         `json:"discount_start_time,omitempty" xml:"discount_start_time,omitempty"`
	IdleTimeLimitType *int                                                                                                                           `json:"idle_time_limit_type,omitempty" xml:"idle_time_limit_type,omitempty"`
	WeekDayList       []*int                                                                                                                         `json:"week_day_list,omitempty" xml:"week_day_list,omitempty" type:"Repeated"`
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDailyTimeRange(v []*AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DailyTimeRange = v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDiscountEndTime(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DiscountEndTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetDiscountStartTime(v int64) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.DiscountStartTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetIdleTimeLimitType(v int) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.IdleTimeLimitType = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo) SetWeekDayList(v []*int) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfo {
	s.WeekDayList = v
	return s
}

type AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem struct {
	EndTime   *string `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) SetEndTime(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem {
	s.EndTime = &v
	return s
}

func (s *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem) SetStartTime(v string) *AkteOrderQueryResponseDataOrdersItemSubOrderAmountInfosItemDiscountsItemDiscountExtraIdleTimeDiscountInfoDailyTimeRangeItem {
	s.StartTime = &v
	return s
}

type AkteOrderQueryResponseDataPage struct {
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Total    *int64 `json:"total,omitempty" xml:"total,omitempty"`
	PageNum  *int32 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

func (s AkteOrderQueryResponseDataPage) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataPage) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataPage) SetPageSize(v int32) *AkteOrderQueryResponseDataPage {
	s.PageSize = &v
	return s
}

func (s *AkteOrderQueryResponseDataPage) SetTotal(v int64) *AkteOrderQueryResponseDataPage {
	s.Total = &v
	return s
}

func (s *AkteOrderQueryResponseDataPage) SetPageNum(v int32) *AkteOrderQueryResponseDataPage {
	s.PageNum = &v
	return s
}

type AkteOrderQueryResponseDataSearchAfter struct {
	AllCursor [][]*string `json:"all_cursor,omitempty" xml:"all_cursor,omitempty" type:"Repeated"`
	Cursor    []*string   `json:"cursor,omitempty" xml:"cursor,omitempty" type:"Repeated"`
	Size      *int32      `json:"size,omitempty" xml:"size,omitempty"`
}

func (s AkteOrderQueryResponseDataSearchAfter) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseDataSearchAfter) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseDataSearchAfter) SetAllCursor(v [][]*string) *AkteOrderQueryResponseDataSearchAfter {
	s.AllCursor = v
	return s
}

func (s *AkteOrderQueryResponseDataSearchAfter) SetCursor(v []*string) *AkteOrderQueryResponseDataSearchAfter {
	s.Cursor = v
	return s
}

func (s *AkteOrderQueryResponseDataSearchAfter) SetSize(v int32) *AkteOrderQueryResponseDataSearchAfter {
	s.Size = &v
	return s
}

type AkteOrderQueryResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s AkteOrderQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AkteOrderQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *AkteOrderQueryResponseExtra) SetNow(v int64) *AkteOrderQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *AkteOrderQueryResponseExtra) SetSubDescription(v string) *AkteOrderQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AkteOrderQueryResponseExtra) SetSubErrorCode(v int32) *AkteOrderQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AkteOrderQueryResponseExtra) SetDescription(v string) *AkteOrderQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *AkteOrderQueryResponseExtra) SetErrorCode(v int32) *AkteOrderQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AkteOrderQueryResponseExtra) SetLogid(v string) *AkteOrderQueryResponseExtra {
	s.Logid = &v
	return s
}

type AliasCreateAliasRequest struct {
	AliasWord   *string            `json:"alias_word,omitempty" xml:"alias_word,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AliasCreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s AliasCreateAliasRequest) GoString() string {
	return s.String()
}

func (s *AliasCreateAliasRequest) SetAliasWord(v string) *AliasCreateAliasRequest {
	s.AliasWord = &v
	return s
}

func (s *AliasCreateAliasRequest) SetHeader(v map[string]*string) *AliasCreateAliasRequest {
	s.Header = v
	return s
}

func (s *AliasCreateAliasRequest) SetAccessToken(v string) *AliasCreateAliasRequest {
	s.AccessToken = &v
	return s
}

type AliasCreateAliasResponse struct {
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s AliasCreateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s AliasCreateAliasResponse) GoString() string {
	return s.String()
}

func (s *AliasCreateAliasResponse) SetErrNo(v int32) *AliasCreateAliasResponse {
	s.ErrNo = &v
	return s
}

func (s *AliasCreateAliasResponse) SetLogId(v string) *AliasCreateAliasResponse {
	s.LogId = &v
	return s
}

func (s *AliasCreateAliasResponse) SetErrMsg(v string) *AliasCreateAliasResponse {
	s.ErrMsg = &v
	return s
}

type AliasListAliasRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AliasListAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s AliasListAliasRequest) GoString() string {
	return s.String()
}

func (s *AliasListAliasRequest) SetAccessToken(v string) *AliasListAliasRequest {
	s.AccessToken = &v
	return s
}

func (s *AliasListAliasRequest) SetHeader(v map[string]*string) *AliasListAliasRequest {
	s.Header = v
	return s
}

type AliasListAliasResponse struct {
	ErrNo  *int32                      `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                     `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                     `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AliasListAliasResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s AliasListAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s AliasListAliasResponse) GoString() string {
	return s.String()
}

func (s *AliasListAliasResponse) SetErrNo(v int32) *AliasListAliasResponse {
	s.ErrNo = &v
	return s
}

func (s *AliasListAliasResponse) SetErrMsg(v string) *AliasListAliasResponse {
	s.ErrMsg = &v
	return s
}

func (s *AliasListAliasResponse) SetLogId(v string) *AliasListAliasResponse {
	s.LogId = &v
	return s
}

func (s *AliasListAliasResponse) SetData(v *AliasListAliasResponseData) *AliasListAliasResponse {
	s.Data = v
	return s
}

type AliasListAliasResponseData struct {
	WeekAvailableTimes *int32                                     `json:"week_available_times,omitempty" xml:"week_available_times,omitempty" require:"true"`
	WeekTotalTimes     *int32                                     `json:"week_total_times,omitempty" xml:"week_total_times,omitempty" require:"true"`
	AliasList          []*AliasListAliasResponseDataAliasListItem `json:"alias_list,omitempty" xml:"alias_list,omitempty" require:"true" type:"Repeated"`
}

func (s AliasListAliasResponseData) String() string {
	return tea.Prettify(s)
}

func (s AliasListAliasResponseData) GoString() string {
	return s.String()
}

func (s *AliasListAliasResponseData) SetWeekAvailableTimes(v int32) *AliasListAliasResponseData {
	s.WeekAvailableTimes = &v
	return s
}

func (s *AliasListAliasResponseData) SetWeekTotalTimes(v int32) *AliasListAliasResponseData {
	s.WeekTotalTimes = &v
	return s
}

func (s *AliasListAliasResponseData) SetAliasList(v []*AliasListAliasResponseDataAliasListItem) *AliasListAliasResponseData {
	s.AliasList = v
	return s
}

type AliasListAliasResponseDataAliasListItem struct {
	Ctime       *int64  `json:"ctime,omitempty" xml:"ctime,omitempty" require:"true"`
	Utime       *int64  `json:"utime,omitempty" xml:"utime,omitempty" require:"true"`
	AliasWord   *string `json:"alias_word,omitempty" xml:"alias_word,omitempty" require:"true"`
	AuditStatus *int64  `json:"audit_status,omitempty" xml:"audit_status,omitempty" require:"true"`
	AuditReason *string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty" require:"true"`
}

func (s AliasListAliasResponseDataAliasListItem) String() string {
	return tea.Prettify(s)
}

func (s AliasListAliasResponseDataAliasListItem) GoString() string {
	return s.String()
}

func (s *AliasListAliasResponseDataAliasListItem) SetCtime(v int64) *AliasListAliasResponseDataAliasListItem {
	s.Ctime = &v
	return s
}

func (s *AliasListAliasResponseDataAliasListItem) SetUtime(v int64) *AliasListAliasResponseDataAliasListItem {
	s.Utime = &v
	return s
}

func (s *AliasListAliasResponseDataAliasListItem) SetAliasWord(v string) *AliasListAliasResponseDataAliasListItem {
	s.AliasWord = &v
	return s
}

func (s *AliasListAliasResponseDataAliasListItem) SetAuditStatus(v int64) *AliasListAliasResponseDataAliasListItem {
	s.AuditStatus = &v
	return s
}

func (s *AliasListAliasResponseDataAliasListItem) SetAuditReason(v string) *AliasListAliasResponseDataAliasListItem {
	s.AuditReason = &v
	return s
}

type AmusementNewRequest struct {
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AmusementNewRequest) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewRequest) GoString() string {
	return s.String()
}

func (s *AmusementNewRequest) SetAccessToken(v string) *AmusementNewRequest {
	s.AccessToken = &v
	return s
}

func (s *AmusementNewRequest) SetHeader(v map[string]*string) *AmusementNewRequest {
	s.Header = v
	return s
}

type AmusementNewResponse struct {
	Extra *AmusementNewResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data  *AmusementNewResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AmusementNewResponse) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewResponse) GoString() string {
	return s.String()
}

func (s *AmusementNewResponse) SetExtra(v *AmusementNewResponseExtra) *AmusementNewResponse {
	s.Extra = v
	return s
}

func (s *AmusementNewResponse) SetData(v *AmusementNewResponseData) *AmusementNewResponse {
	s.Data = v
	return s
}

type AmusementNewResponseData struct {
	GwErrorCode   *int32                              `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                             `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	List          []*AmusementNewResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s AmusementNewResponseData) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewResponseData) GoString() string {
	return s.String()
}

func (s *AmusementNewResponseData) SetGwErrorCode(v int32) *AmusementNewResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AmusementNewResponseData) SetGwDescription(v string) *AmusementNewResponseData {
	s.GwDescription = &v
	return s
}

func (s *AmusementNewResponseData) SetList(v []*AmusementNewResponseDataListItem) *AmusementNewResponseData {
	s.List = v
	return s
}

type AmusementNewResponseDataListItem struct {
	VideoList        []*AmusementNewResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                           `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                          `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                          `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                          `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
	FollowerCount    *int64                                           `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                           `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                         `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
}

func (s AmusementNewResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewResponseDataListItem) GoString() string {
	return s.String()
}

func (s *AmusementNewResponseDataListItem) SetVideoList(v []*AmusementNewResponseDataListItemVideoListItem) *AmusementNewResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *AmusementNewResponseDataListItem) SetRank(v int32) *AmusementNewResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetRankChange(v string) *AmusementNewResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetNickname(v string) *AmusementNewResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetAvatar(v string) *AmusementNewResponseDataListItem {
	s.Avatar = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetFollowerCount(v int64) *AmusementNewResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetOnbillbaordTimes(v int32) *AmusementNewResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *AmusementNewResponseDataListItem) SetEffectValue(v float64) *AmusementNewResponseDataListItem {
	s.EffectValue = &v
	return s
}

type AmusementNewResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s AmusementNewResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *AmusementNewResponseDataListItemVideoListItem) SetTitle(v string) *AmusementNewResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *AmusementNewResponseDataListItemVideoListItem) SetItemCover(v string) *AmusementNewResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *AmusementNewResponseDataListItemVideoListItem) SetShareUrl(v string) *AmusementNewResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type AmusementNewResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AmusementNewResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AmusementNewResponseExtra) GoString() string {
	return s.String()
}

func (s *AmusementNewResponseExtra) SetDescription(v string) *AmusementNewResponseExtra {
	s.Description = &v
	return s
}

func (s *AmusementNewResponseExtra) SetSubErrorCode(v int32) *AmusementNewResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AmusementNewResponseExtra) SetSubDescription(v string) *AmusementNewResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AmusementNewResponseExtra) SetLogid(v string) *AmusementNewResponseExtra {
	s.Logid = &v
	return s
}

func (s *AmusementNewResponseExtra) SetNow(v int64) *AmusementNewResponseExtra {
	s.Now = &v
	return s
}

func (s *AmusementNewResponseExtra) SetErrorCode(v int32) *AmusementNewResponseExtra {
	s.ErrorCode = &v
	return s
}

type AmusementOverallRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AmusementOverallRequest) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallRequest) GoString() string {
	return s.String()
}

func (s *AmusementOverallRequest) SetHeader(v map[string]*string) *AmusementOverallRequest {
	s.Header = v
	return s
}

func (s *AmusementOverallRequest) SetAccessToken(v string) *AmusementOverallRequest {
	s.AccessToken = &v
	return s
}

type AmusementOverallResponse struct {
	Data  *AmusementOverallResponseData  `json:"data,omitempty" xml:"data,omitempty"`
	Extra *AmusementOverallResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AmusementOverallResponse) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallResponse) GoString() string {
	return s.String()
}

func (s *AmusementOverallResponse) SetData(v *AmusementOverallResponseData) *AmusementOverallResponse {
	s.Data = v
	return s
}

func (s *AmusementOverallResponse) SetExtra(v *AmusementOverallResponseExtra) *AmusementOverallResponse {
	s.Extra = v
	return s
}

type AmusementOverallResponseData struct {
	List          []*AmusementOverallResponseDataListItem `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                 `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AmusementOverallResponseData) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallResponseData) GoString() string {
	return s.String()
}

func (s *AmusementOverallResponseData) SetList(v []*AmusementOverallResponseDataListItem) *AmusementOverallResponseData {
	s.List = v
	return s
}

func (s *AmusementOverallResponseData) SetGwErrorCode(v int32) *AmusementOverallResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AmusementOverallResponseData) SetGwDescription(v string) *AmusementOverallResponseData {
	s.GwDescription = &v
	return s
}

type AmusementOverallResponseDataListItem struct {
	FollowerCount    *int64                                               `json:"follower_count,omitempty" xml:"follower_count,omitempty" require:"true"`
	OnbillbaordTimes *int32                                               `json:"onbillbaord_times,omitempty" xml:"onbillbaord_times,omitempty" require:"true"`
	EffectValue      *float64                                             `json:"effect_value,omitempty" xml:"effect_value,omitempty" require:"true"`
	VideoList        []*AmusementOverallResponseDataListItemVideoListItem `json:"video_list,omitempty" xml:"video_list,omitempty" type:"Repeated"`
	Rank             *int32                                               `json:"rank,omitempty" xml:"rank,omitempty" require:"true"`
	RankChange       *string                                              `json:"rank_change,omitempty" xml:"rank_change,omitempty" require:"true"`
	Nickname         *string                                              `json:"nickname,omitempty" xml:"nickname,omitempty" require:"true"`
	Avatar           *string                                              `json:"avatar,omitempty" xml:"avatar,omitempty" require:"true"`
}

func (s AmusementOverallResponseDataListItem) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallResponseDataListItem) GoString() string {
	return s.String()
}

func (s *AmusementOverallResponseDataListItem) SetFollowerCount(v int64) *AmusementOverallResponseDataListItem {
	s.FollowerCount = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetOnbillbaordTimes(v int32) *AmusementOverallResponseDataListItem {
	s.OnbillbaordTimes = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetEffectValue(v float64) *AmusementOverallResponseDataListItem {
	s.EffectValue = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetVideoList(v []*AmusementOverallResponseDataListItemVideoListItem) *AmusementOverallResponseDataListItem {
	s.VideoList = v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetRank(v int32) *AmusementOverallResponseDataListItem {
	s.Rank = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetRankChange(v string) *AmusementOverallResponseDataListItem {
	s.RankChange = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetNickname(v string) *AmusementOverallResponseDataListItem {
	s.Nickname = &v
	return s
}

func (s *AmusementOverallResponseDataListItem) SetAvatar(v string) *AmusementOverallResponseDataListItem {
	s.Avatar = &v
	return s
}

type AmusementOverallResponseDataListItemVideoListItem struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ItemCover *string `json:"item_cover,omitempty" xml:"item_cover,omitempty" require:"true"`
	ShareUrl  *string `json:"share_url,omitempty" xml:"share_url,omitempty"`
}

func (s AmusementOverallResponseDataListItemVideoListItem) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallResponseDataListItemVideoListItem) GoString() string {
	return s.String()
}

func (s *AmusementOverallResponseDataListItemVideoListItem) SetTitle(v string) *AmusementOverallResponseDataListItemVideoListItem {
	s.Title = &v
	return s
}

func (s *AmusementOverallResponseDataListItemVideoListItem) SetItemCover(v string) *AmusementOverallResponseDataListItemVideoListItem {
	s.ItemCover = &v
	return s
}

func (s *AmusementOverallResponseDataListItemVideoListItem) SetShareUrl(v string) *AmusementOverallResponseDataListItemVideoListItem {
	s.ShareUrl = &v
	return s
}

type AmusementOverallResponseExtra struct {
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s AmusementOverallResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AmusementOverallResponseExtra) GoString() string {
	return s.String()
}

func (s *AmusementOverallResponseExtra) SetSubDescription(v string) *AmusementOverallResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AmusementOverallResponseExtra) SetLogid(v string) *AmusementOverallResponseExtra {
	s.Logid = &v
	return s
}

func (s *AmusementOverallResponseExtra) SetNow(v int64) *AmusementOverallResponseExtra {
	s.Now = &v
	return s
}

func (s *AmusementOverallResponseExtra) SetErrorCode(v int32) *AmusementOverallResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AmusementOverallResponseExtra) SetDescription(v string) *AmusementOverallResponseExtra {
	s.Description = &v
	return s
}

func (s *AmusementOverallResponseExtra) SetSubErrorCode(v int32) *AmusementOverallResponseExtra {
	s.SubErrorCode = &v
	return s
}

type AnchorLinkmicCreateRequest struct {
	RoomId            *int64                                 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	MaxLinkNumPerRoom *int64                                 `json:"max_link_num_per_room,omitempty" xml:"max_link_num_per_room,omitempty"`
	Header            map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken       *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RoomIdStr         *string                                `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
	AppId             *string                                `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Rooms             []*AnchorLinkmicCreateRequestRoomsItem `json:"rooms,omitempty" xml:"rooms,omitempty" type:"Repeated"`
}

func (s AnchorLinkmicCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateRequest) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateRequest) SetRoomId(v int64) *AnchorLinkmicCreateRequest {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetMaxLinkNumPerRoom(v int64) *AnchorLinkmicCreateRequest {
	s.MaxLinkNumPerRoom = &v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetHeader(v map[string]*string) *AnchorLinkmicCreateRequest {
	s.Header = v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetAccessToken(v string) *AnchorLinkmicCreateRequest {
	s.AccessToken = &v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetRoomIdStr(v string) *AnchorLinkmicCreateRequest {
	s.RoomIdStr = &v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetAppId(v string) *AnchorLinkmicCreateRequest {
	s.AppId = &v
	return s
}

func (s *AnchorLinkmicCreateRequest) SetRooms(v []*AnchorLinkmicCreateRequestRoomsItem) *AnchorLinkmicCreateRequest {
	s.Rooms = v
	return s
}

type AnchorLinkmicCreateRequestRoomsItem struct {
	RoomId          *int64    `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AudienceOpenIds []*string `json:"audience_open_ids,omitempty" xml:"audience_open_ids,omitempty" type:"Repeated"`
	RoomIdStr       *string   `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
}

func (s AnchorLinkmicCreateRequestRoomsItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateRequestRoomsItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateRequestRoomsItem) SetRoomId(v int64) *AnchorLinkmicCreateRequestRoomsItem {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicCreateRequestRoomsItem) SetAudienceOpenIds(v []*string) *AnchorLinkmicCreateRequestRoomsItem {
	s.AudienceOpenIds = v
	return s
}

func (s *AnchorLinkmicCreateRequestRoomsItem) SetRoomIdStr(v string) *AnchorLinkmicCreateRequestRoomsItem {
	s.RoomIdStr = &v
	return s
}

type AnchorLinkmicCreateResponse struct {
	Errmsg  *string                          `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data    *AnchorLinkmicCreateResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Errcode *int64                           `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s AnchorLinkmicCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponse) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponse) SetErrmsg(v string) *AnchorLinkmicCreateResponse {
	s.Errmsg = &v
	return s
}

func (s *AnchorLinkmicCreateResponse) SetData(v *AnchorLinkmicCreateResponseData) *AnchorLinkmicCreateResponse {
	s.Data = v
	return s
}

func (s *AnchorLinkmicCreateResponse) SetErrcode(v int64) *AnchorLinkmicCreateResponse {
	s.Errcode = &v
	return s
}

type AnchorLinkmicCreateResponseData struct {
	FilterRoomList             []*AnchorLinkmicCreateResponseDataFilterRoomListItem             `json:"filter_room_list,omitempty" xml:"filter_room_list,omitempty" type:"Repeated"`
	JustFilterAudienceRoomList []*AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem `json:"just_filter_audience_room_list,omitempty" xml:"just_filter_audience_room_list,omitempty" type:"Repeated"`
	LinkIdStr                  *string                                                          `json:"link_id_str,omitempty" xml:"link_id_str,omitempty"`
	LinkId                     *int64                                                           `json:"link_id,omitempty" xml:"link_id,omitempty"`
}

func (s AnchorLinkmicCreateResponseData) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponseData) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponseData) SetFilterRoomList(v []*AnchorLinkmicCreateResponseDataFilterRoomListItem) *AnchorLinkmicCreateResponseData {
	s.FilterRoomList = v
	return s
}

func (s *AnchorLinkmicCreateResponseData) SetJustFilterAudienceRoomList(v []*AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) *AnchorLinkmicCreateResponseData {
	s.JustFilterAudienceRoomList = v
	return s
}

func (s *AnchorLinkmicCreateResponseData) SetLinkIdStr(v string) *AnchorLinkmicCreateResponseData {
	s.LinkIdStr = &v
	return s
}

func (s *AnchorLinkmicCreateResponseData) SetLinkId(v int64) *AnchorLinkmicCreateResponseData {
	s.LinkId = &v
	return s
}

type AnchorLinkmicCreateResponseDataFilterRoomListItem struct {
	FilterReason       *int64                                                                     `json:"filterReason,omitempty" xml:"filterReason,omitempty"`
	FilterAudienceList []*AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem `json:"filter_audience_list,omitempty" xml:"filter_audience_list,omitempty" type:"Repeated"`
	RoomIdStr          *string                                                                    `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
	RoomId             *int64                                                                     `json:"room_id,omitempty" xml:"room_id,omitempty"`
	IsFilter           *bool                                                                      `json:"is_filter,omitempty" xml:"is_filter,omitempty"`
}

func (s AnchorLinkmicCreateResponseDataFilterRoomListItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponseDataFilterRoomListItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItem) SetFilterReason(v int64) *AnchorLinkmicCreateResponseDataFilterRoomListItem {
	s.FilterReason = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItem) SetFilterAudienceList(v []*AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem) *AnchorLinkmicCreateResponseDataFilterRoomListItem {
	s.FilterAudienceList = v
	return s
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItem) SetRoomIdStr(v string) *AnchorLinkmicCreateResponseDataFilterRoomListItem {
	s.RoomIdStr = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItem) SetRoomId(v int64) *AnchorLinkmicCreateResponseDataFilterRoomListItem {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItem) SetIsFilter(v bool) *AnchorLinkmicCreateResponseDataFilterRoomListItem {
	s.IsFilter = &v
	return s
}

type AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem struct {
	OpenId       *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	FilterReason *int64  `json:"filterReason,omitempty" xml:"filterReason,omitempty"`
}

func (s AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem) SetOpenId(v string) *AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem {
	s.OpenId = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem) SetFilterReason(v int64) *AnchorLinkmicCreateResponseDataFilterRoomListItemFilterAudienceListItem {
	s.FilterReason = &v
	return s
}

type AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem struct {
	FilterAudienceList []*AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem `json:"filter_audience_list,omitempty" xml:"filter_audience_list,omitempty" type:"Repeated"`
	RoomIdStr          *string                                                                                `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
	RoomId             *int64                                                                                 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	IsFilter           *bool                                                                                  `json:"is_filter,omitempty" xml:"is_filter,omitempty"`
	FilterReason       *int64                                                                                 `json:"filterReason,omitempty" xml:"filterReason,omitempty"`
}

func (s AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) SetFilterAudienceList(v []*AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem {
	s.FilterAudienceList = v
	return s
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) SetRoomIdStr(v string) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem {
	s.RoomIdStr = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) SetRoomId(v int64) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) SetIsFilter(v bool) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem {
	s.IsFilter = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem) SetFilterReason(v int64) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItem {
	s.FilterReason = &v
	return s
}

type AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem struct {
	OpenId       *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	FilterReason *int64  `json:"filterReason,omitempty" xml:"filterReason,omitempty"`
}

func (s AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem) SetOpenId(v string) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem {
	s.OpenId = &v
	return s
}

func (s *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem) SetFilterReason(v int64) *AnchorLinkmicCreateResponseDataJustFilterAudienceRoomListItemFilterAudienceListItem {
	s.FilterReason = &v
	return s
}

type AnchorLinkmicPrepareRequest struct {
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	RoomIdStr   *string            `json:"room_id_str,omitempty" xml:"room_id_str,omitempty"`
	RoomId      *int64             `json:"room_id,omitempty" xml:"room_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AnchorLinkmicPrepareRequest) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicPrepareRequest) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicPrepareRequest) SetAppId(v string) *AnchorLinkmicPrepareRequest {
	s.AppId = &v
	return s
}

func (s *AnchorLinkmicPrepareRequest) SetRoomIdStr(v string) *AnchorLinkmicPrepareRequest {
	s.RoomIdStr = &v
	return s
}

func (s *AnchorLinkmicPrepareRequest) SetRoomId(v int64) *AnchorLinkmicPrepareRequest {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicPrepareRequest) SetHeader(v map[string]*string) *AnchorLinkmicPrepareRequest {
	s.Header = v
	return s
}

func (s *AnchorLinkmicPrepareRequest) SetAccessToken(v string) *AnchorLinkmicPrepareRequest {
	s.AccessToken = &v
	return s
}

type AnchorLinkmicPrepareResponse struct {
	Errcode *int64  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s AnchorLinkmicPrepareResponse) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicPrepareResponse) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicPrepareResponse) SetErrcode(v int64) *AnchorLinkmicPrepareResponse {
	s.Errcode = &v
	return s
}

func (s *AnchorLinkmicPrepareResponse) SetErrmsg(v string) *AnchorLinkmicPrepareResponse {
	s.Errmsg = &v
	return s
}

type AnchorLinkmicQueryRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	LinkId      *int64             `json:"link_id,omitempty" xml:"link_id,omitempty"`
	RoomId      *int64             `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
}

func (s AnchorLinkmicQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryRequest) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryRequest) SetHeader(v map[string]*string) *AnchorLinkmicQueryRequest {
	s.Header = v
	return s
}

func (s *AnchorLinkmicQueryRequest) SetAccessToken(v string) *AnchorLinkmicQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *AnchorLinkmicQueryRequest) SetLinkId(v int64) *AnchorLinkmicQueryRequest {
	s.LinkId = &v
	return s
}

func (s *AnchorLinkmicQueryRequest) SetRoomId(v int64) *AnchorLinkmicQueryRequest {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicQueryRequest) SetAppId(v string) *AnchorLinkmicQueryRequest {
	s.AppId = &v
	return s
}

type AnchorLinkmicQueryResponse struct {
	Errmsg  *string                         `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data    *AnchorLinkmicQueryResponseData `json:"data,omitempty" xml:"data,omitempty"`
	Errcode *int64                          `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s AnchorLinkmicQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryResponse) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryResponse) SetErrmsg(v string) *AnchorLinkmicQueryResponse {
	s.Errmsg = &v
	return s
}

func (s *AnchorLinkmicQueryResponse) SetData(v *AnchorLinkmicQueryResponseData) *AnchorLinkmicQueryResponse {
	s.Data = v
	return s
}

func (s *AnchorLinkmicQueryResponse) SetErrcode(v int64) *AnchorLinkmicQueryResponse {
	s.Errcode = &v
	return s
}

type AnchorLinkmicQueryResponseData struct {
	Rooms []*AnchorLinkmicQueryResponseDataRoomsItem `json:"rooms,omitempty" xml:"rooms,omitempty" type:"Repeated"`
}

func (s AnchorLinkmicQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryResponseData) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryResponseData) SetRooms(v []*AnchorLinkmicQueryResponseDataRoomsItem) *AnchorLinkmicQueryResponseData {
	s.Rooms = v
	return s
}

type AnchorLinkmicQueryResponseDataRoomsItem struct {
	Anchor             *AnchorLinkmicQueryResponseDataRoomsItemAnchor                   `json:"anchor,omitempty" xml:"anchor,omitempty"`
	RoomId             *int64                                                           `json:"room_id,omitempty" xml:"room_id,omitempty"`
	LinkedAudienceList []*AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem `json:"linked_audience_list,omitempty" xml:"linked_audience_list,omitempty" type:"Repeated"`
}

func (s AnchorLinkmicQueryResponseDataRoomsItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryResponseDataRoomsItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryResponseDataRoomsItem) SetAnchor(v *AnchorLinkmicQueryResponseDataRoomsItemAnchor) *AnchorLinkmicQueryResponseDataRoomsItem {
	s.Anchor = v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItem) SetRoomId(v int64) *AnchorLinkmicQueryResponseDataRoomsItem {
	s.RoomId = &v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItem) SetLinkedAudienceList(v []*AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) *AnchorLinkmicQueryResponseDataRoomsItem {
	s.LinkedAudienceList = v
	return s
}

type AnchorLinkmicQueryResponseDataRoomsItemAnchor struct {
	NickName  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	OpenId    *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
}

func (s AnchorLinkmicQueryResponseDataRoomsItemAnchor) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryResponseDataRoomsItemAnchor) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemAnchor) SetNickName(v string) *AnchorLinkmicQueryResponseDataRoomsItemAnchor {
	s.NickName = &v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemAnchor) SetOpenId(v string) *AnchorLinkmicQueryResponseDataRoomsItemAnchor {
	s.OpenId = &v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemAnchor) SetAvatarUrl(v string) *AnchorLinkmicQueryResponseDataRoomsItemAnchor {
	s.AvatarUrl = &v
	return s
}

type AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem struct {
	AvatarUrl *string `json:"avatar_url,omitempty" xml:"avatar_url,omitempty"`
	NickName  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	OpenId    *string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) SetAvatarUrl(v string) *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem {
	s.AvatarUrl = &v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) SetNickName(v string) *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem {
	s.NickName = &v
	return s
}

func (s *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem) SetOpenId(v string) *AnchorLinkmicQueryResponseDataRoomsItemLinkedAudienceListItem {
	s.OpenId = &v
	return s
}

type AnchorLinkmicUpdateVoiceRequest struct {
	LinkId      *int64                                          `json:"link_id,omitempty" xml:"link_id,omitempty"`
	UserInfos   []*AnchorLinkmicUpdateVoiceRequestUserInfosItem `json:"user_infos,omitempty" xml:"user_infos,omitempty" type:"Repeated"`
	AppId       *string                                         `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Header      map[string]*string                              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AnchorLinkmicUpdateVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicUpdateVoiceRequest) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicUpdateVoiceRequest) SetLinkId(v int64) *AnchorLinkmicUpdateVoiceRequest {
	s.LinkId = &v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequest) SetUserInfos(v []*AnchorLinkmicUpdateVoiceRequestUserInfosItem) *AnchorLinkmicUpdateVoiceRequest {
	s.UserInfos = v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequest) SetAppId(v string) *AnchorLinkmicUpdateVoiceRequest {
	s.AppId = &v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequest) SetHeader(v map[string]*string) *AnchorLinkmicUpdateVoiceRequest {
	s.Header = v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequest) SetAccessToken(v string) *AnchorLinkmicUpdateVoiceRequest {
	s.AccessToken = &v
	return s
}

type AnchorLinkmicUpdateVoiceRequestUserInfosItem struct {
	VoiceInfos []*AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem `json:"voice_infos,omitempty" xml:"voice_infos,omitempty" type:"Repeated"`
	OpenId     *string                                                       `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

func (s AnchorLinkmicUpdateVoiceRequestUserInfosItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicUpdateVoiceRequestUserInfosItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicUpdateVoiceRequestUserInfosItem) SetVoiceInfos(v []*AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem) *AnchorLinkmicUpdateVoiceRequestUserInfosItem {
	s.VoiceInfos = v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequestUserInfosItem) SetOpenId(v string) *AnchorLinkmicUpdateVoiceRequestUserInfosItem {
	s.OpenId = &v
	return s
}

type AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem struct {
	NeighborOpenId *string `json:"neighbor_open_id,omitempty" xml:"neighbor_open_id,omitempty"`
	Voice          *int64  `json:"voice,omitempty" xml:"voice,omitempty"`
}

func (s AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem) SetNeighborOpenId(v string) *AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem {
	s.NeighborOpenId = &v
	return s
}

func (s *AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem) SetVoice(v int64) *AnchorLinkmicUpdateVoiceRequestUserInfosItemVoiceInfosItem {
	s.Voice = &v
	return s
}

type AnchorLinkmicUpdateVoiceResponse struct {
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Errcode *int64  `json:"errcode,omitempty" xml:"errcode,omitempty"`
}

func (s AnchorLinkmicUpdateVoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s AnchorLinkmicUpdateVoiceResponse) GoString() string {
	return s.String()
}

func (s *AnchorLinkmicUpdateVoiceResponse) SetErrmsg(v string) *AnchorLinkmicUpdateVoiceResponse {
	s.Errmsg = &v
	return s
}

func (s *AnchorLinkmicUpdateVoiceResponse) SetErrcode(v int64) *AnchorLinkmicUpdateVoiceResponse {
	s.Errcode = &v
	return s
}

type AppAddSubMerchantRequest struct {
	SubMerchantId *string            `json:"sub_merchant_id,omitempty" xml:"sub_merchant_id,omitempty" require:"true"`
	UrlType       *int               `json:"url_type,omitempty" xml:"url_type,omitempty"`
	Role          *int32             `json:"role,omitempty" xml:"role,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	AppId         *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
}

func (s AppAddSubMerchantRequest) String() string {
	return tea.Prettify(s)
}

func (s AppAddSubMerchantRequest) GoString() string {
	return s.String()
}

func (s *AppAddSubMerchantRequest) SetSubMerchantId(v string) *AppAddSubMerchantRequest {
	s.SubMerchantId = &v
	return s
}

func (s *AppAddSubMerchantRequest) SetUrlType(v int) *AppAddSubMerchantRequest {
	s.UrlType = &v
	return s
}

func (s *AppAddSubMerchantRequest) SetRole(v int32) *AppAddSubMerchantRequest {
	s.Role = &v
	return s
}

func (s *AppAddSubMerchantRequest) SetHeader(v map[string]*string) *AppAddSubMerchantRequest {
	s.Header = v
	return s
}

func (s *AppAddSubMerchantRequest) SetAccessToken(v string) *AppAddSubMerchantRequest {
	s.AccessToken = &v
	return s
}

func (s *AppAddSubMerchantRequest) SetAppId(v string) *AppAddSubMerchantRequest {
	s.AppId = &v
	return s
}

type AppAddSubMerchantResponse struct {
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AppAddSubMerchantResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AppAddSubMerchantResponse) String() string {
	return tea.Prettify(s)
}

func (s AppAddSubMerchantResponse) GoString() string {
	return s.String()
}

func (s *AppAddSubMerchantResponse) SetErrNo(v int32) *AppAddSubMerchantResponse {
	s.ErrNo = &v
	return s
}

func (s *AppAddSubMerchantResponse) SetErrMsg(v string) *AppAddSubMerchantResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppAddSubMerchantResponse) SetLogId(v string) *AppAddSubMerchantResponse {
	s.LogId = &v
	return s
}

func (s *AppAddSubMerchantResponse) SetData(v *AppAddSubMerchantResponseData) *AppAddSubMerchantResponse {
	s.Data = v
	return s
}

type AppAddSubMerchantResponseData struct {
	MerchantId *string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty" require:"true"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
}

func (s AppAddSubMerchantResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppAddSubMerchantResponseData) GoString() string {
	return s.String()
}

func (s *AppAddSubMerchantResponseData) SetMerchantId(v string) *AppAddSubMerchantResponseData {
	s.MerchantId = &v
	return s
}

func (s *AppAddSubMerchantResponseData) SetUrl(v string) *AppAddSubMerchantResponseData {
	s.Url = &v
	return s
}

type AppsCheckSessionKeyRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Signature   *string            `json:"signature,omitempty" xml:"signature,omitempty" require:"true"`
	SigMethod   *string            `json:"sig_method,omitempty" xml:"sig_method,omitempty" require:"true"`
	Openid      *string            `json:"openid,omitempty" xml:"openid,omitempty" require:"true"`
}

func (s AppsCheckSessionKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsCheckSessionKeyRequest) GoString() string {
	return s.String()
}

func (s *AppsCheckSessionKeyRequest) SetHeader(v map[string]*string) *AppsCheckSessionKeyRequest {
	s.Header = v
	return s
}

func (s *AppsCheckSessionKeyRequest) SetAccessToken(v string) *AppsCheckSessionKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *AppsCheckSessionKeyRequest) SetSignature(v string) *AppsCheckSessionKeyRequest {
	s.Signature = &v
	return s
}

func (s *AppsCheckSessionKeyRequest) SetSigMethod(v string) *AppsCheckSessionKeyRequest {
	s.SigMethod = &v
	return s
}

func (s *AppsCheckSessionKeyRequest) SetOpenid(v string) *AppsCheckSessionKeyRequest {
	s.Openid = &v
	return s
}

type AppsCheckSessionKeyResponse struct {
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	ErrNo  *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s AppsCheckSessionKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsCheckSessionKeyResponse) GoString() string {
	return s.String()
}

func (s *AppsCheckSessionKeyResponse) SetErrMsg(v string) *AppsCheckSessionKeyResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppsCheckSessionKeyResponse) SetLogId(v string) *AppsCheckSessionKeyResponse {
	s.LogId = &v
	return s
}

func (s *AppsCheckSessionKeyResponse) SetErrNo(v int32) *AppsCheckSessionKeyResponse {
	s.ErrNo = &v
	return s
}

type AppsJscode2sessionRequest struct {
	Secret        *string            `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	Code          *string            `json:"code,omitempty" xml:"code,omitempty"`
	AnonymousCode *string            `json:"anonymous_code,omitempty" xml:"anonymous_code,omitempty"`
	Appid         *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AppsJscode2sessionRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsJscode2sessionRequest) GoString() string {
	return s.String()
}

func (s *AppsJscode2sessionRequest) SetSecret(v string) *AppsJscode2sessionRequest {
	s.Secret = &v
	return s
}

func (s *AppsJscode2sessionRequest) SetCode(v string) *AppsJscode2sessionRequest {
	s.Code = &v
	return s
}

func (s *AppsJscode2sessionRequest) SetAnonymousCode(v string) *AppsJscode2sessionRequest {
	s.AnonymousCode = &v
	return s
}

func (s *AppsJscode2sessionRequest) SetAppid(v string) *AppsJscode2sessionRequest {
	s.Appid = &v
	return s
}

func (s *AppsJscode2sessionRequest) SetHeader(v map[string]*string) *AppsJscode2sessionRequest {
	s.Header = v
	return s
}

type AppsJscode2sessionResponse struct {
	Openid          *string `json:"openid,omitempty" xml:"openid,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	Error           *int64  `json:"error,omitempty" xml:"error,omitempty" require:"true"`
	Unionid         *string `json:"unionid,omitempty" xml:"unionid,omitempty"`
	AnonymousOpenid *string `json:"anonymous_openid,omitempty" xml:"anonymous_openid,omitempty"`
	Errcode         *int64  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	SessionKey      *string `json:"session_key,omitempty" xml:"session_key,omitempty"`
	Errmsg          *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s AppsJscode2sessionResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsJscode2sessionResponse) GoString() string {
	return s.String()
}

func (s *AppsJscode2sessionResponse) SetOpenid(v string) *AppsJscode2sessionResponse {
	s.Openid = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetMessage(v string) *AppsJscode2sessionResponse {
	s.Message = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetError(v int64) *AppsJscode2sessionResponse {
	s.Error = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetUnionid(v string) *AppsJscode2sessionResponse {
	s.Unionid = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetAnonymousOpenid(v string) *AppsJscode2sessionResponse {
	s.AnonymousOpenid = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetErrcode(v int64) *AppsJscode2sessionResponse {
	s.Errcode = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetSessionKey(v string) *AppsJscode2sessionResponse {
	s.SessionKey = &v
	return s
}

func (s *AppsJscode2sessionResponse) SetErrmsg(v string) *AppsJscode2sessionResponse {
	s.Errmsg = &v
	return s
}

type AppsRemoveUserStorageRequest struct {
	Signature   *string            `json:"signature,omitempty" xml:"signature,omitempty" require:"true"`
	SigMethod   *string            `json:"sig_method,omitempty" xml:"sig_method,omitempty" require:"true"`
	Openid      *string            `json:"openid,omitempty" xml:"openid,omitempty" require:"true"`
	Key         []*string          `json:"key,omitempty" xml:"key,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsRemoveUserStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsRemoveUserStorageRequest) GoString() string {
	return s.String()
}

func (s *AppsRemoveUserStorageRequest) SetSignature(v string) *AppsRemoveUserStorageRequest {
	s.Signature = &v
	return s
}

func (s *AppsRemoveUserStorageRequest) SetSigMethod(v string) *AppsRemoveUserStorageRequest {
	s.SigMethod = &v
	return s
}

func (s *AppsRemoveUserStorageRequest) SetOpenid(v string) *AppsRemoveUserStorageRequest {
	s.Openid = &v
	return s
}

func (s *AppsRemoveUserStorageRequest) SetKey(v []*string) *AppsRemoveUserStorageRequest {
	s.Key = v
	return s
}

func (s *AppsRemoveUserStorageRequest) SetHeader(v map[string]*string) *AppsRemoveUserStorageRequest {
	s.Header = v
	return s
}

func (s *AppsRemoveUserStorageRequest) SetAccessToken(v string) *AppsRemoveUserStorageRequest {
	s.AccessToken = &v
	return s
}

type AppsRemoveUserStorageResponse struct {
	Error   *int32  `json:"error,omitempty" xml:"error,omitempty" require:"true"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Errcode *int32  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s AppsRemoveUserStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsRemoveUserStorageResponse) GoString() string {
	return s.String()
}

func (s *AppsRemoveUserStorageResponse) SetError(v int32) *AppsRemoveUserStorageResponse {
	s.Error = &v
	return s
}

func (s *AppsRemoveUserStorageResponse) SetMessage(v string) *AppsRemoveUserStorageResponse {
	s.Message = &v
	return s
}

func (s *AppsRemoveUserStorageResponse) SetErrcode(v int32) *AppsRemoveUserStorageResponse {
	s.Errcode = &v
	return s
}

func (s *AppsRemoveUserStorageResponse) SetErrmsg(v string) *AppsRemoveUserStorageResponse {
	s.Errmsg = &v
	return s
}

type AppsResetSessionKeyRequest struct {
	Openid      *string            `json:"openid,omitempty" xml:"openid,omitempty" require:"true"`
	Signature   *string            `json:"signature,omitempty" xml:"signature,omitempty" require:"true"`
	SigMethod   *string            `json:"sig_method,omitempty" xml:"sig_method,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsResetSessionKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsResetSessionKeyRequest) GoString() string {
	return s.String()
}

func (s *AppsResetSessionKeyRequest) SetOpenid(v string) *AppsResetSessionKeyRequest {
	s.Openid = &v
	return s
}

func (s *AppsResetSessionKeyRequest) SetSignature(v string) *AppsResetSessionKeyRequest {
	s.Signature = &v
	return s
}

func (s *AppsResetSessionKeyRequest) SetSigMethod(v string) *AppsResetSessionKeyRequest {
	s.SigMethod = &v
	return s
}

func (s *AppsResetSessionKeyRequest) SetHeader(v map[string]*string) *AppsResetSessionKeyRequest {
	s.Header = v
	return s
}

func (s *AppsResetSessionKeyRequest) SetAccessToken(v string) *AppsResetSessionKeyRequest {
	s.AccessToken = &v
	return s
}

type AppsResetSessionKeyResponse struct {
	SessionKey *string `json:"session_key,omitempty" xml:"session_key,omitempty" require:"true"`
	ErrNo      *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg     *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId      *string `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Openid     *string `json:"openid,omitempty" xml:"openid,omitempty" require:"true"`
}

func (s AppsResetSessionKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsResetSessionKeyResponse) GoString() string {
	return s.String()
}

func (s *AppsResetSessionKeyResponse) SetSessionKey(v string) *AppsResetSessionKeyResponse {
	s.SessionKey = &v
	return s
}

func (s *AppsResetSessionKeyResponse) SetErrNo(v int32) *AppsResetSessionKeyResponse {
	s.ErrNo = &v
	return s
}

func (s *AppsResetSessionKeyResponse) SetErrMsg(v string) *AppsResetSessionKeyResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppsResetSessionKeyResponse) SetLogId(v string) *AppsResetSessionKeyResponse {
	s.LogId = &v
	return s
}

func (s *AppsResetSessionKeyResponse) SetOpenid(v string) *AppsResetSessionKeyResponse {
	s.Openid = &v
	return s
}

type AppsSetUserStorageRequest struct {
	AccessToken *string                                `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Signature   *string                                `json:"signature,omitempty" xml:"signature,omitempty" require:"true"`
	SigMethod   *string                                `json:"sig_method,omitempty" xml:"sig_method,omitempty" require:"true"`
	Openid      *string                                `json:"openid,omitempty" xml:"openid,omitempty" require:"true"`
	KvList      []*AppsSetUserStorageRequestKvListItem `json:"kv_list,omitempty" xml:"kv_list,omitempty" require:"true" type:"Repeated"`
	Header      map[string]*string                     `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AppsSetUserStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsSetUserStorageRequest) GoString() string {
	return s.String()
}

func (s *AppsSetUserStorageRequest) SetAccessToken(v string) *AppsSetUserStorageRequest {
	s.AccessToken = &v
	return s
}

func (s *AppsSetUserStorageRequest) SetSignature(v string) *AppsSetUserStorageRequest {
	s.Signature = &v
	return s
}

func (s *AppsSetUserStorageRequest) SetSigMethod(v string) *AppsSetUserStorageRequest {
	s.SigMethod = &v
	return s
}

func (s *AppsSetUserStorageRequest) SetOpenid(v string) *AppsSetUserStorageRequest {
	s.Openid = &v
	return s
}

func (s *AppsSetUserStorageRequest) SetKvList(v []*AppsSetUserStorageRequestKvListItem) *AppsSetUserStorageRequest {
	s.KvList = v
	return s
}

func (s *AppsSetUserStorageRequest) SetHeader(v map[string]*string) *AppsSetUserStorageRequest {
	s.Header = v
	return s
}

type AppsSetUserStorageRequestKvListItem struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
	Value *string `json:"value,omitempty" xml:"value,omitempty" require:"true"`
}

func (s AppsSetUserStorageRequestKvListItem) String() string {
	return tea.Prettify(s)
}

func (s AppsSetUserStorageRequestKvListItem) GoString() string {
	return s.String()
}

func (s *AppsSetUserStorageRequestKvListItem) SetKey(v string) *AppsSetUserStorageRequestKvListItem {
	s.Key = &v
	return s
}

func (s *AppsSetUserStorageRequestKvListItem) SetValue(v string) *AppsSetUserStorageRequestKvListItem {
	s.Value = &v
	return s
}

type AppsSetUserStorageResponse struct {
	Error   *int32  `json:"error,omitempty" xml:"error,omitempty" require:"true"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Errcode *int32  `json:"errcode,omitempty" xml:"errcode,omitempty"`
	Errmsg  *string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

func (s AppsSetUserStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsSetUserStorageResponse) GoString() string {
	return s.String()
}

func (s *AppsSetUserStorageResponse) SetError(v int32) *AppsSetUserStorageResponse {
	s.Error = &v
	return s
}

func (s *AppsSetUserStorageResponse) SetMessage(v string) *AppsSetUserStorageResponse {
	s.Message = &v
	return s
}

func (s *AppsSetUserStorageResponse) SetErrcode(v int32) *AppsSetUserStorageResponse {
	s.Errcode = &v
	return s
}

func (s *AppsSetUserStorageResponse) SetErrmsg(v string) *AppsSetUserStorageResponse {
	s.Errmsg = &v
	return s
}

type AppsStableTokenRequest struct {
	Secret       *string            `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	GrantType    *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	ForceRefresh *bool              `json:"force_refresh,omitempty" xml:"force_refresh,omitempty"`
	Appid        *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Header       map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AppsStableTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsStableTokenRequest) GoString() string {
	return s.String()
}

func (s *AppsStableTokenRequest) SetSecret(v string) *AppsStableTokenRequest {
	s.Secret = &v
	return s
}

func (s *AppsStableTokenRequest) SetGrantType(v string) *AppsStableTokenRequest {
	s.GrantType = &v
	return s
}

func (s *AppsStableTokenRequest) SetForceRefresh(v bool) *AppsStableTokenRequest {
	s.ForceRefresh = &v
	return s
}

func (s *AppsStableTokenRequest) SetAppid(v string) *AppsStableTokenRequest {
	s.Appid = &v
	return s
}

func (s *AppsStableTokenRequest) SetHeader(v map[string]*string) *AppsStableTokenRequest {
	s.Header = v
	return s
}

type AppsStableTokenResponse struct {
	Data   *AppsStableTokenResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s AppsStableTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsStableTokenResponse) GoString() string {
	return s.String()
}

func (s *AppsStableTokenResponse) SetData(v *AppsStableTokenResponseData) *AppsStableTokenResponse {
	s.Data = v
	return s
}

func (s *AppsStableTokenResponse) SetErrNo(v int32) *AppsStableTokenResponse {
	s.ErrNo = &v
	return s
}

func (s *AppsStableTokenResponse) SetErrMsg(v string) *AppsStableTokenResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppsStableTokenResponse) SetLogId(v string) *AppsStableTokenResponse {
	s.LogId = &v
	return s
}

type AppsStableTokenResponseData struct {
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsStableTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppsStableTokenResponseData) GoString() string {
	return s.String()
}

func (s *AppsStableTokenResponseData) SetExpiresIn(v int64) *AppsStableTokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *AppsStableTokenResponseData) SetAccessToken(v string) *AppsStableTokenResponseData {
	s.AccessToken = &v
	return s
}

type AppsUrlLinkGenerateRequest struct {
	VersionType *string            `json:"version_type,omitempty" xml:"version_type,omitempty"`
	AppName     *string            `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	Query       *string            `json:"query,omitempty" xml:"query,omitempty"`
	ExpireTime  *int64             `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsUrlLinkGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkGenerateRequest) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkGenerateRequest) SetVersionType(v string) *AppsUrlLinkGenerateRequest {
	s.VersionType = &v
	return s
}

func (s *AppsUrlLinkGenerateRequest) SetAppName(v string) *AppsUrlLinkGenerateRequest {
	s.AppName = &v
	return s
}

func (s *AppsUrlLinkGenerateRequest) SetQuery(v string) *AppsUrlLinkGenerateRequest {
	s.Query = &v
	return s
}

func (s *AppsUrlLinkGenerateRequest) SetExpireTime(v int64) *AppsUrlLinkGenerateRequest {
	s.ExpireTime = &v
	return s
}

func (s *AppsUrlLinkGenerateRequest) SetHeader(v map[string]*string) *AppsUrlLinkGenerateRequest {
	s.Header = v
	return s
}

func (s *AppsUrlLinkGenerateRequest) SetAccessToken(v string) *AppsUrlLinkGenerateRequest {
	s.AccessToken = &v
	return s
}

type AppsUrlLinkGenerateResponse struct {
	ErrNo  *int32                           `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                          `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                          `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AppsUrlLinkGenerateResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AppsUrlLinkGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkGenerateResponse) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkGenerateResponse) SetErrNo(v int32) *AppsUrlLinkGenerateResponse {
	s.ErrNo = &v
	return s
}

func (s *AppsUrlLinkGenerateResponse) SetErrMsg(v string) *AppsUrlLinkGenerateResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppsUrlLinkGenerateResponse) SetLogId(v string) *AppsUrlLinkGenerateResponse {
	s.LogId = &v
	return s
}

func (s *AppsUrlLinkGenerateResponse) SetData(v *AppsUrlLinkGenerateResponseData) *AppsUrlLinkGenerateResponse {
	s.Data = v
	return s
}

type AppsUrlLinkGenerateResponseData struct {
	UrlLink *string `json:"url_link,omitempty" xml:"url_link,omitempty" require:"true"`
}

func (s AppsUrlLinkGenerateResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkGenerateResponseData) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkGenerateResponseData) SetUrlLink(v string) *AppsUrlLinkGenerateResponseData {
	s.UrlLink = &v
	return s
}

type AppsUrlLinkQueryInfoRequest struct {
	UrlLink     *string            `json:"url_link,omitempty" xml:"url_link,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsUrlLinkQueryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryInfoRequest) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryInfoRequest) SetUrlLink(v string) *AppsUrlLinkQueryInfoRequest {
	s.UrlLink = &v
	return s
}

func (s *AppsUrlLinkQueryInfoRequest) SetHeader(v map[string]*string) *AppsUrlLinkQueryInfoRequest {
	s.Header = v
	return s
}

func (s *AppsUrlLinkQueryInfoRequest) SetAccessToken(v string) *AppsUrlLinkQueryInfoRequest {
	s.AccessToken = &v
	return s
}

type AppsUrlLinkQueryInfoResponse struct {
	LogId  *string                           `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AppsUrlLinkQueryInfoResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrNo  *int32                            `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                           `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s AppsUrlLinkQueryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryInfoResponse) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryInfoResponse) SetLogId(v string) *AppsUrlLinkQueryInfoResponse {
	s.LogId = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponse) SetData(v *AppsUrlLinkQueryInfoResponseData) *AppsUrlLinkQueryInfoResponse {
	s.Data = v
	return s
}

func (s *AppsUrlLinkQueryInfoResponse) SetErrNo(v int32) *AppsUrlLinkQueryInfoResponse {
	s.ErrNo = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponse) SetErrMsg(v string) *AppsUrlLinkQueryInfoResponse {
	s.ErrMsg = &v
	return s
}

type AppsUrlLinkQueryInfoResponseData struct {
	AppName     *string `json:"app_name,omitempty" xml:"app_name,omitempty" require:"true"`
	AppId       *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	Path        *string `json:"path,omitempty" xml:"path,omitempty" require:"true"`
	Query       *string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	CreateTime  *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	ExpireTime  *int64  `json:"expire_time,omitempty" xml:"expire_time,omitempty" require:"true"`
	VersionType *string `json:"version_type,omitempty" xml:"version_type,omitempty" require:"true"`
}

func (s AppsUrlLinkQueryInfoResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryInfoResponseData) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryInfoResponseData) SetAppName(v string) *AppsUrlLinkQueryInfoResponseData {
	s.AppName = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetAppId(v string) *AppsUrlLinkQueryInfoResponseData {
	s.AppId = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetPath(v string) *AppsUrlLinkQueryInfoResponseData {
	s.Path = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetQuery(v string) *AppsUrlLinkQueryInfoResponseData {
	s.Query = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetCreateTime(v int64) *AppsUrlLinkQueryInfoResponseData {
	s.CreateTime = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetExpireTime(v int64) *AppsUrlLinkQueryInfoResponseData {
	s.ExpireTime = &v
	return s
}

func (s *AppsUrlLinkQueryInfoResponseData) SetVersionType(v string) *AppsUrlLinkQueryInfoResponseData {
	s.VersionType = &v
	return s
}

type AppsUrlLinkQueryQuotaRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AppsUrlLinkQueryQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryQuotaRequest) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryQuotaRequest) SetHeader(v map[string]*string) *AppsUrlLinkQueryQuotaRequest {
	s.Header = v
	return s
}

func (s *AppsUrlLinkQueryQuotaRequest) SetAccessToken(v string) *AppsUrlLinkQueryQuotaRequest {
	s.AccessToken = &v
	return s
}

type AppsUrlLinkQueryQuotaResponse struct {
	ErrNo  *int32                             `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	ErrMsg *string                            `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                            `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AppsUrlLinkQueryQuotaResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AppsUrlLinkQueryQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryQuotaResponse) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryQuotaResponse) SetErrNo(v int32) *AppsUrlLinkQueryQuotaResponse {
	s.ErrNo = &v
	return s
}

func (s *AppsUrlLinkQueryQuotaResponse) SetErrMsg(v string) *AppsUrlLinkQueryQuotaResponse {
	s.ErrMsg = &v
	return s
}

func (s *AppsUrlLinkQueryQuotaResponse) SetLogId(v string) *AppsUrlLinkQueryQuotaResponse {
	s.LogId = &v
	return s
}

func (s *AppsUrlLinkQueryQuotaResponse) SetData(v *AppsUrlLinkQueryQuotaResponseData) *AppsUrlLinkQueryQuotaResponse {
	s.Data = v
	return s
}

type AppsUrlLinkQueryQuotaResponseData struct {
	UrlLinkUsed  *int32 `json:"url_link_used,omitempty" xml:"url_link_used,omitempty" require:"true"`
	UrlLinkLimit *int32 `json:"url_link_limit,omitempty" xml:"url_link_limit,omitempty" require:"true"`
}

func (s AppsUrlLinkQueryQuotaResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppsUrlLinkQueryQuotaResponseData) GoString() string {
	return s.String()
}

func (s *AppsUrlLinkQueryQuotaResponseData) SetUrlLinkUsed(v int32) *AppsUrlLinkQueryQuotaResponseData {
	s.UrlLinkUsed = &v
	return s
}

func (s *AppsUrlLinkQueryQuotaResponseData) SetUrlLinkLimit(v int32) *AppsUrlLinkQueryQuotaResponseData {
	s.UrlLinkLimit = &v
	return s
}

type AppsV2TokenRequest struct {
	Appid     *string            `json:"appid,omitempty" xml:"appid,omitempty" require:"true"`
	Secret    *string            `json:"secret,omitempty" xml:"secret,omitempty" require:"true"`
	GrantType *string            `json:"grant_type,omitempty" xml:"grant_type,omitempty" require:"true"`
	Header    map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s AppsV2TokenRequest) String() string {
	return tea.Prettify(s)
}

func (s AppsV2TokenRequest) GoString() string {
	return s.String()
}

func (s *AppsV2TokenRequest) SetAppid(v string) *AppsV2TokenRequest {
	s.Appid = &v
	return s
}

func (s *AppsV2TokenRequest) SetSecret(v string) *AppsV2TokenRequest {
	s.Secret = &v
	return s
}

func (s *AppsV2TokenRequest) SetGrantType(v string) *AppsV2TokenRequest {
	s.GrantType = &v
	return s
}

func (s *AppsV2TokenRequest) SetHeader(v map[string]*string) *AppsV2TokenRequest {
	s.Header = v
	return s
}

type AppsV2TokenResponse struct {
	ErrTips *string                  `json:"err_tips,omitempty" xml:"err_tips,omitempty" require:"true"`
	Data    *AppsV2TokenResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo   *int64                   `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s AppsV2TokenResponse) String() string {
	return tea.Prettify(s)
}

func (s AppsV2TokenResponse) GoString() string {
	return s.String()
}

func (s *AppsV2TokenResponse) SetErrTips(v string) *AppsV2TokenResponse {
	s.ErrTips = &v
	return s
}

func (s *AppsV2TokenResponse) SetData(v *AppsV2TokenResponseData) *AppsV2TokenResponse {
	s.Data = v
	return s
}

func (s *AppsV2TokenResponse) SetErrNo(v int64) *AppsV2TokenResponse {
	s.ErrNo = &v
	return s
}

type AppsV2TokenResponseData struct {
	ExpiresIn   *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty" require:"true"`
	ExpiresAt   *int64  `json:"expires_at,omitempty" xml:"expires_at,omitempty" require:"true"`
	AccessToken *string `json:"access_token,omitempty" xml:"access_token,omitempty" require:"true"`
}

func (s AppsV2TokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s AppsV2TokenResponseData) GoString() string {
	return s.String()
}

func (s *AppsV2TokenResponseData) SetExpiresIn(v int64) *AppsV2TokenResponseData {
	s.ExpiresIn = &v
	return s
}

func (s *AppsV2TokenResponseData) SetExpiresAt(v int64) *AppsV2TokenResponseData {
	s.ExpiresAt = &v
	return s
}

func (s *AppsV2TokenResponseData) SetAccessToken(v string) *AppsV2TokenResponseData {
	s.AccessToken = &v
	return s
}

type AriNotifyRequest struct {
	DateRange   *AriNotifyRequestDateRange `json:"date_range,omitempty" xml:"date_range,omitempty" require:"true"`
	NotifyScene []*int                     `json:"notify_scene,omitempty" xml:"notify_scene,omitempty" type:"Repeated"`
	Header      map[string]*string         `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string                    `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RatePlanIds []*string                  `json:"rate_plan_ids,omitempty" xml:"rate_plan_ids,omitempty" type:"Repeated"`
	AccountId   *string                    `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s AriNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyRequest) GoString() string {
	return s.String()
}

func (s *AriNotifyRequest) SetDateRange(v *AriNotifyRequestDateRange) *AriNotifyRequest {
	s.DateRange = v
	return s
}

func (s *AriNotifyRequest) SetNotifyScene(v []*int) *AriNotifyRequest {
	s.NotifyScene = v
	return s
}

func (s *AriNotifyRequest) SetHeader(v map[string]*string) *AriNotifyRequest {
	s.Header = v
	return s
}

func (s *AriNotifyRequest) SetAccessToken(v string) *AriNotifyRequest {
	s.AccessToken = &v
	return s
}

func (s *AriNotifyRequest) SetRatePlanIds(v []*string) *AriNotifyRequest {
	s.RatePlanIds = v
	return s
}

func (s *AriNotifyRequest) SetAccountId(v string) *AriNotifyRequest {
	s.AccountId = &v
	return s
}

type AriNotifyRequestDateRange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s AriNotifyRequestDateRange) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyRequestDateRange) GoString() string {
	return s.String()
}

func (s *AriNotifyRequestDateRange) SetEnd(v string) *AriNotifyRequestDateRange {
	s.End = &v
	return s
}

func (s *AriNotifyRequestDateRange) SetStart(v string) *AriNotifyRequestDateRange {
	s.Start = &v
	return s
}

type AriNotifyResponse struct {
	Data  *AriNotifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *AriNotifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AriNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyResponse) GoString() string {
	return s.String()
}

func (s *AriNotifyResponse) SetData(v *AriNotifyResponseData) *AriNotifyResponse {
	s.Data = v
	return s
}

func (s *AriNotifyResponse) SetExtra(v *AriNotifyResponseExtra) *AriNotifyResponse {
	s.Extra = v
	return s
}

type AriNotifyResponseData struct {
	GwDescription *string                                `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Code          *int32                                 `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Message       *string                                `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	SaveResult    []*AriNotifyResponseDataSaveResultItem `json:"save_result,omitempty" xml:"save_result,omitempty" type:"Repeated"`
	GwErrorCode   *int32                                 `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AriNotifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyResponseData) GoString() string {
	return s.String()
}

func (s *AriNotifyResponseData) SetGwDescription(v string) *AriNotifyResponseData {
	s.GwDescription = &v
	return s
}

func (s *AriNotifyResponseData) SetCode(v int32) *AriNotifyResponseData {
	s.Code = &v
	return s
}

func (s *AriNotifyResponseData) SetMessage(v string) *AriNotifyResponseData {
	s.Message = &v
	return s
}

func (s *AriNotifyResponseData) SetSaveResult(v []*AriNotifyResponseDataSaveResultItem) *AriNotifyResponseData {
	s.SaveResult = v
	return s
}

func (s *AriNotifyResponseData) SetGwErrorCode(v int32) *AriNotifyResponseData {
	s.GwErrorCode = &v
	return s
}

type AriNotifyResponseDataSaveResultItem struct {
	Message    *string `json:"message,omitempty" xml:"message,omitempty" require:"true"`
	Pass       *bool   `json:"pass,omitempty" xml:"pass,omitempty" require:"true"`
	RatePlanId *string `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty" require:"true"`
}

func (s AriNotifyResponseDataSaveResultItem) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyResponseDataSaveResultItem) GoString() string {
	return s.String()
}

func (s *AriNotifyResponseDataSaveResultItem) SetMessage(v string) *AriNotifyResponseDataSaveResultItem {
	s.Message = &v
	return s
}

func (s *AriNotifyResponseDataSaveResultItem) SetPass(v bool) *AriNotifyResponseDataSaveResultItem {
	s.Pass = &v
	return s
}

func (s *AriNotifyResponseDataSaveResultItem) SetRatePlanId(v string) *AriNotifyResponseDataSaveResultItem {
	s.RatePlanId = &v
	return s
}

type AriNotifyResponseExtra struct {
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
}

func (s AriNotifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AriNotifyResponseExtra) GoString() string {
	return s.String()
}

func (s *AriNotifyResponseExtra) SetNow(v int64) *AriNotifyResponseExtra {
	s.Now = &v
	return s
}

func (s *AriNotifyResponseExtra) SetSubDescription(v string) *AriNotifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AriNotifyResponseExtra) SetSubErrorCode(v int32) *AriNotifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AriNotifyResponseExtra) SetDescription(v string) *AriNotifyResponseExtra {
	s.Description = &v
	return s
}

func (s *AriNotifyResponseExtra) SetErrorCode(v int32) *AriNotifyResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AriNotifyResponseExtra) SetLogid(v string) *AriNotifyResponseExtra {
	s.Logid = &v
	return s
}

type AuditGetRequest struct {
	Cursor      *int64             `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Count       *int32             `json:"count,omitempty" xml:"count,omitempty" require:"true"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AuditGetRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditGetRequest) GoString() string {
	return s.String()
}

func (s *AuditGetRequest) SetCursor(v int64) *AuditGetRequest {
	s.Cursor = &v
	return s
}

func (s *AuditGetRequest) SetCount(v int32) *AuditGetRequest {
	s.Count = &v
	return s
}

func (s *AuditGetRequest) SetOpenId(v string) *AuditGetRequest {
	s.OpenId = &v
	return s
}

func (s *AuditGetRequest) SetHeader(v map[string]*string) *AuditGetRequest {
	s.Header = v
	return s
}

func (s *AuditGetRequest) SetAccessToken(v string) *AuditGetRequest {
	s.AccessToken = &v
	return s
}

type AuditGetResponse struct {
	Data      *AuditGetResponseData      `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ApplyList *AuditGetResponseApplyList `json:"apply_list,omitempty" xml:"apply_list,omitempty" require:"true"`
	Extra     *AuditGetResponseExtra     `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AuditGetResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditGetResponse) GoString() string {
	return s.String()
}

func (s *AuditGetResponse) SetData(v *AuditGetResponseData) *AuditGetResponse {
	s.Data = v
	return s
}

func (s *AuditGetResponse) SetApplyList(v *AuditGetResponseApplyList) *AuditGetResponse {
	s.ApplyList = v
	return s
}

func (s *AuditGetResponse) SetExtra(v *AuditGetResponseExtra) *AuditGetResponse {
	s.Extra = v
	return s
}

type AuditGetResponseApplyList struct {
	HasMore *bool                                `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
	List    []*AuditGetResponseApplyListListItem `json:"list,omitempty" xml:"list,omitempty" require:"true" type:"Repeated"`
	Cursor  *int64                               `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
}

func (s AuditGetResponseApplyList) String() string {
	return tea.Prettify(s)
}

func (s AuditGetResponseApplyList) GoString() string {
	return s.String()
}

func (s *AuditGetResponseApplyList) SetHasMore(v bool) *AuditGetResponseApplyList {
	s.HasMore = &v
	return s
}

func (s *AuditGetResponseApplyList) SetList(v []*AuditGetResponseApplyListListItem) *AuditGetResponseApplyList {
	s.List = v
	return s
}

func (s *AuditGetResponseApplyList) SetCursor(v int64) *AuditGetResponseApplyList {
	s.Cursor = &v
	return s
}

type AuditGetResponseApplyListListItem struct {
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty" require:"true"`
	ApplyId     *string `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	GroupId     *string `json:"group_id,omitempty" xml:"group_id,omitempty" require:"true"`
	ApplyStatus *int32  `json:"apply_status,omitempty" xml:"apply_status,omitempty" require:"true"`
	CreateTime  *int64  `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
}

func (s AuditGetResponseApplyListListItem) String() string {
	return tea.Prettify(s)
}

func (s AuditGetResponseApplyListListItem) GoString() string {
	return s.String()
}

func (s *AuditGetResponseApplyListListItem) SetUserId(v string) *AuditGetResponseApplyListListItem {
	s.UserId = &v
	return s
}

func (s *AuditGetResponseApplyListListItem) SetApplyId(v string) *AuditGetResponseApplyListListItem {
	s.ApplyId = &v
	return s
}

func (s *AuditGetResponseApplyListListItem) SetGroupId(v string) *AuditGetResponseApplyListListItem {
	s.GroupId = &v
	return s
}

func (s *AuditGetResponseApplyListListItem) SetApplyStatus(v int32) *AuditGetResponseApplyListListItem {
	s.ApplyStatus = &v
	return s
}

func (s *AuditGetResponseApplyListListItem) SetCreateTime(v int64) *AuditGetResponseApplyListListItem {
	s.CreateTime = &v
	return s
}

type AuditGetResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AuditGetResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuditGetResponseData) GoString() string {
	return s.String()
}

func (s *AuditGetResponseData) SetGwErrorCode(v int32) *AuditGetResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AuditGetResponseData) SetGwDescription(v string) *AuditGetResponseData {
	s.GwDescription = &v
	return s
}

type AuditGetResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AuditGetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AuditGetResponseExtra) GoString() string {
	return s.String()
}

func (s *AuditGetResponseExtra) SetDescription(v string) *AuditGetResponseExtra {
	s.Description = &v
	return s
}

func (s *AuditGetResponseExtra) SetSubErrorCode(v int32) *AuditGetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AuditGetResponseExtra) SetSubDescription(v string) *AuditGetResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AuditGetResponseExtra) SetLogid(v string) *AuditGetResponseExtra {
	s.Logid = &v
	return s
}

func (s *AuditGetResponseExtra) SetNow(v int64) *AuditGetResponseExtra {
	s.Now = &v
	return s
}

func (s *AuditGetResponseExtra) SetErrorCode(v int32) *AuditGetResponseExtra {
	s.ErrorCode = &v
	return s
}

type AuditNotifyRequest struct {
	AfterSaleId    *string                         `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty" require:"true"`
	IsApproved     *bool                           `json:"is_approved,omitempty" xml:"is_approved,omitempty" require:"true"`
	OutAfterSaleId *string                         `json:"out_after_sale_id,omitempty" xml:"out_after_sale_id,omitempty"`
	RejectReason   *AuditNotifyRequestRejectReason `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	Header         map[string]*string              `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken    *string                         `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AuditNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditNotifyRequest) GoString() string {
	return s.String()
}

func (s *AuditNotifyRequest) SetAfterSaleId(v string) *AuditNotifyRequest {
	s.AfterSaleId = &v
	return s
}

func (s *AuditNotifyRequest) SetIsApproved(v bool) *AuditNotifyRequest {
	s.IsApproved = &v
	return s
}

func (s *AuditNotifyRequest) SetOutAfterSaleId(v string) *AuditNotifyRequest {
	s.OutAfterSaleId = &v
	return s
}

func (s *AuditNotifyRequest) SetRejectReason(v *AuditNotifyRequestRejectReason) *AuditNotifyRequest {
	s.RejectReason = v
	return s
}

func (s *AuditNotifyRequest) SetHeader(v map[string]*string) *AuditNotifyRequest {
	s.Header = v
	return s
}

func (s *AuditNotifyRequest) SetAccessToken(v string) *AuditNotifyRequest {
	s.AccessToken = &v
	return s
}

type AuditNotifyRequestRejectReason struct {
	Desc       *string  `json:"desc,omitempty" xml:"desc,omitempty"`
	ReasonCode []*int64 `json:"reason_code,omitempty" xml:"reason_code,omitempty" require:"true" type:"Repeated"`
}

func (s AuditNotifyRequestRejectReason) String() string {
	return tea.Prettify(s)
}

func (s AuditNotifyRequestRejectReason) GoString() string {
	return s.String()
}

func (s *AuditNotifyRequestRejectReason) SetDesc(v string) *AuditNotifyRequestRejectReason {
	s.Desc = &v
	return s
}

func (s *AuditNotifyRequestRejectReason) SetReasonCode(v []*int64) *AuditNotifyRequestRejectReason {
	s.ReasonCode = v
	return s
}

type AuditNotifyResponse struct {
	Data  *AuditNotifyResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Extra *AuditNotifyResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
}

func (s AuditNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditNotifyResponse) GoString() string {
	return s.String()
}

func (s *AuditNotifyResponse) SetData(v *AuditNotifyResponseData) *AuditNotifyResponse {
	s.Data = v
	return s
}

func (s *AuditNotifyResponse) SetExtra(v *AuditNotifyResponseExtra) *AuditNotifyResponse {
	s.Extra = v
	return s
}

type AuditNotifyResponseData struct {
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s AuditNotifyResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuditNotifyResponseData) GoString() string {
	return s.String()
}

func (s *AuditNotifyResponseData) SetGwErrorCode(v int32) *AuditNotifyResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *AuditNotifyResponseData) SetGwDescription(v string) *AuditNotifyResponseData {
	s.GwDescription = &v
	return s
}

type AuditNotifyResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AuditNotifyResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AuditNotifyResponseExtra) GoString() string {
	return s.String()
}

func (s *AuditNotifyResponseExtra) SetLogid(v string) *AuditNotifyResponseExtra {
	s.Logid = &v
	return s
}

func (s *AuditNotifyResponseExtra) SetNow(v int64) *AuditNotifyResponseExtra {
	s.Now = &v
	return s
}

func (s *AuditNotifyResponseExtra) SetSubDescription(v string) *AuditNotifyResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *AuditNotifyResponseExtra) SetSubErrorCode(v int32) *AuditNotifyResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AuditNotifyResponseExtra) SetDescription(v string) *AuditNotifyResponseExtra {
	s.Description = &v
	return s
}

func (s *AuditNotifyResponseExtra) SetErrorCode(v int32) *AuditNotifyResponseExtra {
	s.ErrorCode = &v
	return s
}

type AuditSetRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Status      *int64             `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	ApplyId     *string            `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AuditSetRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditSetRequest) GoString() string {
	return s.String()
}

func (s *AuditSetRequest) SetOpenId(v string) *AuditSetRequest {
	s.OpenId = &v
	return s
}

func (s *AuditSetRequest) SetStatus(v int64) *AuditSetRequest {
	s.Status = &v
	return s
}

func (s *AuditSetRequest) SetApplyId(v string) *AuditSetRequest {
	s.ApplyId = &v
	return s
}

func (s *AuditSetRequest) SetHeader(v map[string]*string) *AuditSetRequest {
	s.Header = v
	return s
}

func (s *AuditSetRequest) SetAccessToken(v string) *AuditSetRequest {
	s.AccessToken = &v
	return s
}

type AuditSetResponse struct {
	Extra   *AuditSetResponseExtra `json:"extra,omitempty" xml:"extra,omitempty"`
	Data    *AuditSetResponseData  `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	Success *bool                  `json:"success,omitempty" xml:"success,omitempty" require:"true"`
}

func (s AuditSetResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditSetResponse) GoString() string {
	return s.String()
}

func (s *AuditSetResponse) SetExtra(v *AuditSetResponseExtra) *AuditSetResponse {
	s.Extra = v
	return s
}

func (s *AuditSetResponse) SetData(v *AuditSetResponseData) *AuditSetResponse {
	s.Data = v
	return s
}

func (s *AuditSetResponse) SetSuccess(v bool) *AuditSetResponse {
	s.Success = &v
	return s
}

type AuditSetResponseData struct {
	GwDescription *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	GwErrorCode   *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
}

func (s AuditSetResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuditSetResponseData) GoString() string {
	return s.String()
}

func (s *AuditSetResponseData) SetGwDescription(v string) *AuditSetResponseData {
	s.GwDescription = &v
	return s
}

func (s *AuditSetResponseData) SetGwErrorCode(v int32) *AuditSetResponseData {
	s.GwErrorCode = &v
	return s
}

type AuditSetResponseExtra struct {
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
}

func (s AuditSetResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s AuditSetResponseExtra) GoString() string {
	return s.String()
}

func (s *AuditSetResponseExtra) SetLogid(v string) *AuditSetResponseExtra {
	s.Logid = &v
	return s
}

func (s *AuditSetResponseExtra) SetNow(v int64) *AuditSetResponseExtra {
	s.Now = &v
	return s
}

func (s *AuditSetResponseExtra) SetErrorCode(v int32) *AuditSetResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *AuditSetResponseExtra) SetDescription(v string) *AuditSetResponseExtra {
	s.Description = &v
	return s
}

func (s *AuditSetResponseExtra) SetSubErrorCode(v int32) *AuditSetResponseExtra {
	s.SubErrorCode = &v
	return s
}

func (s *AuditSetResponseExtra) SetSubDescription(v string) *AuditSetResponseExtra {
	s.SubDescription = &v
	return s
}

type AuthGetRelatedIdRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AuthGetRelatedIdRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthGetRelatedIdRequest) GoString() string {
	return s.String()
}

func (s *AuthGetRelatedIdRequest) SetOpenId(v string) *AuthGetRelatedIdRequest {
	s.OpenId = &v
	return s
}

func (s *AuthGetRelatedIdRequest) SetHeader(v map[string]*string) *AuthGetRelatedIdRequest {
	s.Header = v
	return s
}

func (s *AuthGetRelatedIdRequest) SetAccessToken(v string) *AuthGetRelatedIdRequest {
	s.AccessToken = &v
	return s
}

type AuthGetRelatedIdResponse struct {
	ErrMsg *string                       `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                        `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                       `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AuthGetRelatedIdResponseData `json:"data,omitempty" xml:"data,omitempty"`
}

func (s AuthGetRelatedIdResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthGetRelatedIdResponse) GoString() string {
	return s.String()
}

func (s *AuthGetRelatedIdResponse) SetErrMsg(v string) *AuthGetRelatedIdResponse {
	s.ErrMsg = &v
	return s
}

func (s *AuthGetRelatedIdResponse) SetErrNo(v int32) *AuthGetRelatedIdResponse {
	s.ErrNo = &v
	return s
}

func (s *AuthGetRelatedIdResponse) SetLogId(v string) *AuthGetRelatedIdResponse {
	s.LogId = &v
	return s
}

func (s *AuthGetRelatedIdResponse) SetData(v *AuthGetRelatedIdResponseData) *AuthGetRelatedIdResponse {
	s.Data = v
	return s
}

type AuthGetRelatedIdResponseData struct {
	AlliedId *string `json:"allied_id,omitempty" xml:"allied_id,omitempty"`
}

func (s AuthGetRelatedIdResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuthGetRelatedIdResponseData) GoString() string {
	return s.String()
}

func (s *AuthGetRelatedIdResponseData) SetAlliedId(v string) *AuthGetRelatedIdResponseData {
	s.AlliedId = &v
	return s
}

type AuthorizeStatusRequest struct {
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	COpenId     *string            `json:"c_open_id,omitempty" xml:"c_open_id,omitempty" require:"true"`
	AppId       *string            `json:"app_id,omitempty" xml:"app_id,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
}

func (s AuthorizeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeStatusRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeStatusRequest) SetOpenId(v string) *AuthorizeStatusRequest {
	s.OpenId = &v
	return s
}

func (s *AuthorizeStatusRequest) SetCOpenId(v string) *AuthorizeStatusRequest {
	s.COpenId = &v
	return s
}

func (s *AuthorizeStatusRequest) SetAppId(v string) *AuthorizeStatusRequest {
	s.AppId = &v
	return s
}

func (s *AuthorizeStatusRequest) SetHeader(v map[string]*string) *AuthorizeStatusRequest {
	s.Header = v
	return s
}

func (s *AuthorizeStatusRequest) SetAccessToken(v string) *AuthorizeStatusRequest {
	s.AccessToken = &v
	return s
}

type AuthorizeStatusResponse struct {
	Data   *AuthorizeStatusResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                      `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                       `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
	LogId  *string                      `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
}

func (s AuthorizeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeStatusResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeStatusResponse) SetData(v *AuthorizeStatusResponseData) *AuthorizeStatusResponse {
	s.Data = v
	return s
}

func (s *AuthorizeStatusResponse) SetErrMsg(v string) *AuthorizeStatusResponse {
	s.ErrMsg = &v
	return s
}

func (s *AuthorizeStatusResponse) SetErrNo(v int32) *AuthorizeStatusResponse {
	s.ErrNo = &v
	return s
}

func (s *AuthorizeStatusResponse) SetLogId(v string) *AuthorizeStatusResponse {
	s.LogId = &v
	return s
}

type AuthorizeStatusResponseData struct {
	EnterFrom   *string                                 `json:"enter_from,omitempty" xml:"enter_from,omitempty"`
	Path        *string                                 `json:"path,omitempty" xml:"path,omitempty"`
	Query       *string                                 `json:"query,omitempty" xml:"query,omitempty"`
	Status      *int64                                  `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	Attribution *AuthorizeStatusResponseDataAttribution `json:"attribution,omitempty" xml:"attribution,omitempty"`
	DataImExtra *string                                 `json:"data_im_extra,omitempty" xml:"data_im_extra,omitempty"`
}

func (s AuthorizeStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeStatusResponseData) GoString() string {
	return s.String()
}

func (s *AuthorizeStatusResponseData) SetEnterFrom(v string) *AuthorizeStatusResponseData {
	s.EnterFrom = &v
	return s
}

func (s *AuthorizeStatusResponseData) SetPath(v string) *AuthorizeStatusResponseData {
	s.Path = &v
	return s
}

func (s *AuthorizeStatusResponseData) SetQuery(v string) *AuthorizeStatusResponseData {
	s.Query = &v
	return s
}

func (s *AuthorizeStatusResponseData) SetStatus(v int64) *AuthorizeStatusResponseData {
	s.Status = &v
	return s
}

func (s *AuthorizeStatusResponseData) SetAttribution(v *AuthorizeStatusResponseDataAttribution) *AuthorizeStatusResponseData {
	s.Attribution = v
	return s
}

func (s *AuthorizeStatusResponseData) SetDataImExtra(v string) *AuthorizeStatusResponseData {
	s.DataImExtra = &v
	return s
}

type AuthorizeStatusResponseDataAttribution struct {
	GroupId        *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	MessageId      *string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	RoomId         *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	AnchorId       *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	AuthorId       *string `json:"author_id,omitempty" xml:"author_id,omitempty"`
	ConversationId *string `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
	FromAttUid     *string `json:"from_att_uid,omitempty" xml:"from_att_uid,omitempty"`
}

func (s AuthorizeStatusResponseDataAttribution) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeStatusResponseDataAttribution) GoString() string {
	return s.String()
}

func (s *AuthorizeStatusResponseDataAttribution) SetGroupId(v string) *AuthorizeStatusResponseDataAttribution {
	s.GroupId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetMessageId(v string) *AuthorizeStatusResponseDataAttribution {
	s.MessageId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetRoomId(v string) *AuthorizeStatusResponseDataAttribution {
	s.RoomId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetAnchorId(v string) *AuthorizeStatusResponseDataAttribution {
	s.AnchorId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetAuthorId(v string) *AuthorizeStatusResponseDataAttribution {
	s.AuthorId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetConversationId(v string) *AuthorizeStatusResponseDataAttribution {
	s.ConversationId = &v
	return s
}

func (s *AuthorizeStatusResponseDataAttribution) SetFromAttUid(v string) *AuthorizeStatusResponseDataAttribution {
	s.FromAttUid = &v
	return s
}

type AuthorizeUserListRequest struct {
	PageSize    *int64             `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Cursor      *string            `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Limit       *int64             `json:"limit,omitempty" xml:"limit,omitempty"`
	PageNum     *int64             `json:"page_num,omitempty" xml:"page_num,omitempty"`
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	OpenId      *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
}

func (s AuthorizeUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeUserListRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeUserListRequest) SetPageSize(v int64) *AuthorizeUserListRequest {
	s.PageSize = &v
	return s
}

func (s *AuthorizeUserListRequest) SetCursor(v string) *AuthorizeUserListRequest {
	s.Cursor = &v
	return s
}

func (s *AuthorizeUserListRequest) SetLimit(v int64) *AuthorizeUserListRequest {
	s.Limit = &v
	return s
}

func (s *AuthorizeUserListRequest) SetPageNum(v int64) *AuthorizeUserListRequest {
	s.PageNum = &v
	return s
}

func (s *AuthorizeUserListRequest) SetHeader(v map[string]*string) *AuthorizeUserListRequest {
	s.Header = v
	return s
}

func (s *AuthorizeUserListRequest) SetAccessToken(v string) *AuthorizeUserListRequest {
	s.AccessToken = &v
	return s
}

func (s *AuthorizeUserListRequest) SetOpenId(v string) *AuthorizeUserListRequest {
	s.OpenId = &v
	return s
}

type AuthorizeUserListResponse struct {
	LogId  *string                        `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *AuthorizeUserListResponseData `json:"data,omitempty" xml:"data,omitempty"`
	ErrMsg *string                        `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	ErrNo  *int32                         `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s AuthorizeUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeUserListResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeUserListResponse) SetLogId(v string) *AuthorizeUserListResponse {
	s.LogId = &v
	return s
}

func (s *AuthorizeUserListResponse) SetData(v *AuthorizeUserListResponseData) *AuthorizeUserListResponse {
	s.Data = v
	return s
}

func (s *AuthorizeUserListResponse) SetErrMsg(v string) *AuthorizeUserListResponse {
	s.ErrMsg = &v
	return s
}

func (s *AuthorizeUserListResponse) SetErrNo(v int32) *AuthorizeUserListResponse {
	s.ErrNo = &v
	return s
}

type AuthorizeUserListResponseData struct {
	NextCursor   *string                                          `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	AuthUserList []*AuthorizeUserListResponseDataAuthUserListItem `json:"auth_user_list,omitempty" xml:"auth_user_list,omitempty" type:"Repeated"`
	HasMore      *bool                                            `json:"has_more,omitempty" xml:"has_more,omitempty"`
}

func (s AuthorizeUserListResponseData) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeUserListResponseData) GoString() string {
	return s.String()
}

func (s *AuthorizeUserListResponseData) SetNextCursor(v string) *AuthorizeUserListResponseData {
	s.NextCursor = &v
	return s
}

func (s *AuthorizeUserListResponseData) SetAuthUserList(v []*AuthorizeUserListResponseDataAuthUserListItem) *AuthorizeUserListResponseData {
	s.AuthUserList = v
	return s
}

func (s *AuthorizeUserListResponseData) SetHasMore(v bool) *AuthorizeUserListResponseData {
	s.HasMore = &v
	return s
}

type AuthorizeUserListResponseDataAuthUserListItem struct {
	EnterFrom           *string                                                   `json:"enter_from,omitempty" xml:"enter_from,omitempty" require:"true"`
	Path                *string                                                   `json:"path,omitempty" xml:"path,omitempty"`
	Query               *string                                                   `json:"query,omitempty" xml:"query,omitempty"`
	TargetOpenId        *string                                                   `json:"target_open_id,omitempty" xml:"target_open_id,omitempty"`
	Attribution         *AuthorizeUserListResponseDataAuthUserListItemAttribution `json:"attribution,omitempty" xml:"attribution,omitempty"`
	AuthUserSourceAppId *string                                                   `json:"auth_user_source_app_id,omitempty" xml:"auth_user_source_app_id,omitempty"`
	DataImExtra         *string                                                   `json:"data_im_extra,omitempty" xml:"data_im_extra,omitempty"`
}

func (s AuthorizeUserListResponseDataAuthUserListItem) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeUserListResponseDataAuthUserListItem) GoString() string {
	return s.String()
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetEnterFrom(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.EnterFrom = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetPath(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.Path = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetQuery(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.Query = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetTargetOpenId(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.TargetOpenId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetAttribution(v *AuthorizeUserListResponseDataAuthUserListItemAttribution) *AuthorizeUserListResponseDataAuthUserListItem {
	s.Attribution = v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetAuthUserSourceAppId(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.AuthUserSourceAppId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItem) SetDataImExtra(v string) *AuthorizeUserListResponseDataAuthUserListItem {
	s.DataImExtra = &v
	return s
}

type AuthorizeUserListResponseDataAuthUserListItemAttribution struct {
	AnchorId       *string `json:"anchor_id,omitempty" xml:"anchor_id,omitempty"`
	AuthorId       *string `json:"author_id,omitempty" xml:"author_id,omitempty"`
	ConversationId *string `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
	FromAttUid     *string `json:"from_att_uid,omitempty" xml:"from_att_uid,omitempty"`
	GroupId        *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	MessageId      *string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	RoomId         *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

func (s AuthorizeUserListResponseDataAuthUserListItemAttribution) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeUserListResponseDataAuthUserListItemAttribution) GoString() string {
	return s.String()
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetAnchorId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.AnchorId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetAuthorId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.AuthorId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetConversationId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.ConversationId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetFromAttUid(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.FromAttUid = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetGroupId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.GroupId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetMessageId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.MessageId = &v
	return s
}

func (s *AuthorizeUserListResponseDataAuthUserListItemAttribution) SetRoomId(v string) *AuthorizeUserListResponseDataAuthUserListItemAttribution {
	s.RoomId = &v
	return s
}

type BatchRollbackConsumeCouponRequest struct {
	OrderId              *string            `json:"order_id,omitempty" xml:"order_id,omitempty"`
	ConsumeOutNo         *string            `json:"consume_out_no,omitempty" xml:"consume_out_no,omitempty" require:"true"`
	OpenId               *string            `json:"open_id,omitempty" xml:"open_id,omitempty" require:"true"`
	AppId                *string            `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	RollbackConsumeOutNo *string            `json:"rollback_consume_out_no,omitempty" xml:"rollback_consume_out_no,omitempty" require:"true"`
	CouponIdList         []*string          `json:"coupon_id_list,omitempty" xml:"coupon_id_list,omitempty" require:"true" type:"Repeated"`
	AccessToken          *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RollbackConsumeTime  *int64             `json:"rollback_consume_time,omitempty" xml:"rollback_consume_time,omitempty" require:"true"`
	Header               map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s BatchRollbackConsumeCouponRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRollbackConsumeCouponRequest) GoString() string {
	return s.String()
}

func (s *BatchRollbackConsumeCouponRequest) SetOrderId(v string) *BatchRollbackConsumeCouponRequest {
	s.OrderId = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetConsumeOutNo(v string) *BatchRollbackConsumeCouponRequest {
	s.ConsumeOutNo = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetOpenId(v string) *BatchRollbackConsumeCouponRequest {
	s.OpenId = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetAppId(v string) *BatchRollbackConsumeCouponRequest {
	s.AppId = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetRollbackConsumeOutNo(v string) *BatchRollbackConsumeCouponRequest {
	s.RollbackConsumeOutNo = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetCouponIdList(v []*string) *BatchRollbackConsumeCouponRequest {
	s.CouponIdList = v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetAccessToken(v string) *BatchRollbackConsumeCouponRequest {
	s.AccessToken = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetRollbackConsumeTime(v int64) *BatchRollbackConsumeCouponRequest {
	s.RollbackConsumeTime = &v
	return s
}

func (s *BatchRollbackConsumeCouponRequest) SetHeader(v map[string]*string) *BatchRollbackConsumeCouponRequest {
	s.Header = v
	return s
}

type BatchRollbackConsumeCouponResponse struct {
	ErrMsg *string                                 `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	LogId  *string                                 `json:"log_id,omitempty" xml:"log_id,omitempty" require:"true"`
	Data   *BatchRollbackConsumeCouponResponseData `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	ErrNo  *int32                                  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s BatchRollbackConsumeCouponResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRollbackConsumeCouponResponse) GoString() string {
	return s.String()
}

func (s *BatchRollbackConsumeCouponResponse) SetErrMsg(v string) *BatchRollbackConsumeCouponResponse {
	s.ErrMsg = &v
	return s
}

func (s *BatchRollbackConsumeCouponResponse) SetLogId(v string) *BatchRollbackConsumeCouponResponse {
	s.LogId = &v
	return s
}

func (s *BatchRollbackConsumeCouponResponse) SetData(v *BatchRollbackConsumeCouponResponseData) *BatchRollbackConsumeCouponResponse {
	s.Data = v
	return s
}

func (s *BatchRollbackConsumeCouponResponse) SetErrNo(v int32) *BatchRollbackConsumeCouponResponse {
	s.ErrNo = &v
	return s
}

type BatchRollbackConsumeCouponResponseData struct {
	Results []*BatchRollbackConsumeCouponResponseDataResultsItem `json:"results,omitempty" xml:"results,omitempty" require:"true" type:"Repeated"`
}

func (s BatchRollbackConsumeCouponResponseData) String() string {
	return tea.Prettify(s)
}

func (s BatchRollbackConsumeCouponResponseData) GoString() string {
	return s.String()
}

func (s *BatchRollbackConsumeCouponResponseData) SetResults(v []*BatchRollbackConsumeCouponResponseDataResultsItem) *BatchRollbackConsumeCouponResponseData {
	s.Results = v
	return s
}

type BatchRollbackConsumeCouponResponseDataResultsItem struct {
	ErrMsg   *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
	CouponId *string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty" require:"true"`
	ErrNo    *int32  `json:"err_no,omitempty" xml:"err_no,omitempty" require:"true"`
}

func (s BatchRollbackConsumeCouponResponseDataResultsItem) String() string {
	return tea.Prettify(s)
}

func (s BatchRollbackConsumeCouponResponseDataResultsItem) GoString() string {
	return s.String()
}

func (s *BatchRollbackConsumeCouponResponseDataResultsItem) SetErrMsg(v string) *BatchRollbackConsumeCouponResponseDataResultsItem {
	s.ErrMsg = &v
	return s
}

func (s *BatchRollbackConsumeCouponResponseDataResultsItem) SetCouponId(v string) *BatchRollbackConsumeCouponResponseDataResultsItem {
	s.CouponId = &v
	return s
}

func (s *BatchRollbackConsumeCouponResponseDataResultsItem) SetErrNo(v int32) *BatchRollbackConsumeCouponResponseDataResultsItem {
	s.ErrNo = &v
	return s
}

type BillCateringQueryRequest struct {
	AccountId     *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	BillDate      *string            `json:"bill_date,omitempty" xml:"bill_date,omitempty" require:"true"`
	BizType       *int               `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	Cursor        *string            `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	WithdrawId    *string            `json:"withdraw_id,omitempty" xml:"withdraw_id,omitempty"`
	Header        map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken   *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	RootAccountId *string            `json:"root_account_id,omitempty" xml:"root_account_id,omitempty"`
	Size          *int64             `json:"size,omitempty" xml:"size,omitempty" require:"true"`
}

func (s BillCateringQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s BillCateringQueryRequest) GoString() string {
	return s.String()
}

func (s *BillCateringQueryRequest) SetAccountId(v string) *BillCateringQueryRequest {
	s.AccountId = &v
	return s
}

func (s *BillCateringQueryRequest) SetBillDate(v string) *BillCateringQueryRequest {
	s.BillDate = &v
	return s
}

func (s *BillCateringQueryRequest) SetBizType(v int) *BillCateringQueryRequest {
	s.BizType = &v
	return s
}

func (s *BillCateringQueryRequest) SetCursor(v string) *BillCateringQueryRequest {
	s.Cursor = &v
	return s
}

func (s *BillCateringQueryRequest) SetWithdrawId(v string) *BillCateringQueryRequest {
	s.WithdrawId = &v
	return s
}

func (s *BillCateringQueryRequest) SetHeader(v map[string]*string) *BillCateringQueryRequest {
	s.Header = v
	return s
}

func (s *BillCateringQueryRequest) SetAccessToken(v string) *BillCateringQueryRequest {
	s.AccessToken = &v
	return s
}

func (s *BillCateringQueryRequest) SetRootAccountId(v string) *BillCateringQueryRequest {
	s.RootAccountId = &v
	return s
}

func (s *BillCateringQueryRequest) SetSize(v int64) *BillCateringQueryRequest {
	s.Size = &v
	return s
}

type BillCateringQueryResponse struct {
	Extra *BillCateringQueryResponseExtra `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
	Data  *BillCateringQueryResponseData  `json:"data,omitempty" xml:"data,omitempty"`
}

func (s BillCateringQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s BillCateringQueryResponse) GoString() string {
	return s.String()
}

func (s *BillCateringQueryResponse) SetExtra(v *BillCateringQueryResponseExtra) *BillCateringQueryResponse {
	s.Extra = v
	return s
}

func (s *BillCateringQueryResponse) SetData(v *BillCateringQueryResponseData) *BillCateringQueryResponse {
	s.Data = v
	return s
}

type BillCateringQueryResponseData struct {
	LedgerRecords []*BillCateringQueryResponseDataLedgerRecordsItem `json:"ledger_records,omitempty" xml:"ledger_records,omitempty" require:"true" type:"Repeated"`
	GwErrorCode   *int32                                            `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	GwDescription *string                                           `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	Cursor        *string                                           `json:"cursor,omitempty" xml:"cursor,omitempty" require:"true"`
	HasMore       *bool                                             `json:"has_more,omitempty" xml:"has_more,omitempty" require:"true"`
}

func (s BillCateringQueryResponseData) String() string {
	return tea.Prettify(s)
}

func (s BillCateringQueryResponseData) GoString() string {
	return s.String()
}

func (s *BillCateringQueryResponseData) SetLedgerRecords(v []*BillCateringQueryResponseDataLedgerRecordsItem) *BillCateringQueryResponseData {
	s.LedgerRecords = v
	return s
}

func (s *BillCateringQueryResponseData) SetGwErrorCode(v int32) *BillCateringQueryResponseData {
	s.GwErrorCode = &v
	return s
}

func (s *BillCateringQueryResponseData) SetGwDescription(v string) *BillCateringQueryResponseData {
	s.GwDescription = &v
	return s
}

func (s *BillCateringQueryResponseData) SetCursor(v string) *BillCateringQueryResponseData {
	s.Cursor = &v
	return s
}

func (s *BillCateringQueryResponseData) SetHasMore(v bool) *BillCateringQueryResponseData {
	s.HasMore = &v
	return s
}

type BillCateringQueryResponseDataLedgerRecordsItem struct {
	VerifyId                *string           `json:"verify_id,omitempty" xml:"verify_id,omitempty"`
	PoiId                   *string           `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
	SettleTime              *int64            `json:"settle_time,omitempty" xml:"settle_time,omitempty"`
	SettleAccountPoiId      *string           `json:"settle_account_poi_id,omitempty" xml:"settle_account_poi_id,omitempty"`
	FundAmountType          *int64            `json:"fund_amount_type,omitempty" xml:"fund_amount_type,omitempty"`
	VerifyPoiAccountId      *string           `json:"verify_poi_account_id,omitempty" xml:"verify_poi_account_id,omitempty"`
	Amount                  map[string]*int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	OuterVerifyCouponNumber *string           `json:"outer_verify_coupon_number,omitempty" xml:"outer_verify_coupon_number,omitempty"`
	CertificateId           *string           `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	FlowId                  *string           `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	LedgerId                *string           `json:"ledger_id,omitempty" xml:"ledger_id,omitempty"`
	OuterVerifyOrderNumber  *string           `json:"outer_verify_order_number,omitempty" xml:"outer_verify_order_number,omitempty"`
	OrderExtId              *string           `json:"order_ext_id,omitempty" xml:"order_ext_id,omitempty"`
	SettleTypeDesc          *string           `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	ItemOrderId             *string           `json:"item_order_id,omitempty" xml:"item_order_id,omitempty"`
	VerifySerialNum         *int32            `json:"verify_serial_num,omitempty" xml:"verify_serial_num,omitempty"`
	GroupName               *string           `json:"group_name,omitempty" xml:"group_name,omitempty"`
	SkuId                   *string           `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	AccountId               *string           `json:"account_id,omitempty" xml:"account_id,omitempty"`
	ShopOrderId             *string           `json:"shop_order_id,omitempty" xml:"shop_order_id,omitempty"`
	SettleSubType           *int64            `json:"settle_sub_type,omitempty" xml:"settle_sub_type,omitempty"`
	SettleType              *int64            `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	BizType                 *int              `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	Code                    *string           `json:"code,omitempty" xml:"code,omitempty"`
	AssociatedShopOrderId   *string           `json:"associated_shop_order_id,omitempty" xml:"associated_shop_order_id,omitempty"`
	FundAmount              *int64            `json:"fund_amount,omitempty" xml:"fund_amount,omitempty"`
}

func (s BillCateringQueryResponseDataLedgerRecordsItem) String() string {
	return tea.Prettify(s)
}

func (s BillCateringQueryResponseDataLedgerRecordsItem) GoString() string {
	return s.String()
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetVerifyId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.VerifyId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetPoiId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.PoiId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSettleTime(v int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SettleTime = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSettleAccountPoiId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SettleAccountPoiId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetFundAmountType(v int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.FundAmountType = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetVerifyPoiAccountId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.VerifyPoiAccountId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetAmount(v map[string]*int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.Amount = v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetOuterVerifyCouponNumber(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.OuterVerifyCouponNumber = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetCertificateId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.CertificateId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetFlowId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.FlowId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetLedgerId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.LedgerId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetOuterVerifyOrderNumber(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.OuterVerifyOrderNumber = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetOrderExtId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.OrderExtId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSettleTypeDesc(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SettleTypeDesc = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetItemOrderId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.ItemOrderId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetVerifySerialNum(v int32) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.VerifySerialNum = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetGroupName(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.GroupName = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSkuId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SkuId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetAccountId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.AccountId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetShopOrderId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.ShopOrderId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSettleSubType(v int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SettleSubType = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetSettleType(v int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.SettleType = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetBizType(v int) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.BizType = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetCode(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.Code = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetAssociatedShopOrderId(v string) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.AssociatedShopOrderId = &v
	return s
}

func (s *BillCateringQueryResponseDataLedgerRecordsItem) SetFundAmount(v int64) *BillCateringQueryResponseDataLedgerRecordsItem {
	s.FundAmount = &v
	return s
}

type BillCateringQueryResponseExtra struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
	ErrorCode      *int32  `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	Logid          *string `json:"logid,omitempty" xml:"logid,omitempty" require:"true"`
	Now            *int64  `json:"now,omitempty" xml:"now,omitempty" require:"true"`
	SubDescription *string `json:"sub_description,omitempty" xml:"sub_description,omitempty" require:"true"`
	SubErrorCode   *int32  `json:"sub_error_code,omitempty" xml:"sub_error_code,omitempty" require:"true"`
}

func (s BillCateringQueryResponseExtra) String() string {
	return tea.Prettify(s)
}

func (s BillCateringQueryResponseExtra) GoString() string {
	return s.String()
}

func (s *BillCateringQueryResponseExtra) SetDescription(v string) *BillCateringQueryResponseExtra {
	s.Description = &v
	return s
}

func (s *BillCateringQueryResponseExtra) SetErrorCode(v int32) *BillCateringQueryResponseExtra {
	s.ErrorCode = &v
	return s
}

func (s *BillCateringQueryResponseExtra) SetLogid(v string) *BillCateringQueryResponseExtra {
	s.Logid = &v
	return s
}

func (s *BillCateringQueryResponseExtra) SetNow(v int64) *BillCateringQueryResponseExtra {
	s.Now = &v
	return s
}

func (s *BillCateringQueryResponseExtra) SetSubDescription(v string) *BillCateringQueryResponseExtra {
	s.SubDescription = &v
	return s
}

func (s *BillCateringQueryResponseExtra) SetSubErrorCode(v int32) *BillCateringQueryResponseExtra {
	s.SubErrorCode = &v
	return s
}

type BillQueryLegerUrlRequest struct {
	Header      map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
	AccessToken *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	LaunchDate  *string            `json:"launch_date,omitempty" xml:"launch_date,omitempty" require:"true"`
	AccountId   *string            `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
}

func (s BillQueryLegerUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s BillQueryLegerUrlRequest) GoString() string {
	return s.String()
}

func (s *BillQueryLegerUrlRequest) SetHeader(v map[string]*string) *BillQueryLegerUrlRequest {
	s.Header = v
	return s
}

func (s *BillQueryLegerUrlRequest) SetAccessToken(v string) *BillQueryLegerUrlRequest {
	s.AccessToken = &v
	return s
}

func (s *BillQueryLegerUrlRequest) SetLaunchDate(v string) *BillQueryLegerUrlRequest {
	s.LaunchDate = &v
	return s
}

func (s *BillQueryLegerUrlRequest) SetAccountId(v string) *BillQueryLegerUrlRequest {
	s.AccountId = &v
	return s
}
