/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package friend

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                *string              `json:"namespaceId"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	FollowScript               *ScriptSetting       `json:"followScript"`
	UnfollowScript             *ScriptSetting       `json:"unfollowScript"`
	SendRequestScript          *ScriptSetting       `json:"sendRequestScript"`
	CancelRequestScript        *ScriptSetting       `json:"cancelRequestScript"`
	AcceptRequestScript        *ScriptSetting       `json:"acceptRequestScript"`
	RejectRequestScript        *ScriptSetting       `json:"rejectRequestScript"`
	DeleteFriendScript         *ScriptSetting       `json:"deleteFriendScript"`
	UpdateProfileScript        *ScriptSetting       `json:"updateProfileScript"`
	FollowNotification         *NotificationSetting `json:"followNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	AcceptRequestNotification  *NotificationSetting `json:"acceptRequestNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	CreatedAt                  *int64               `json:"createdAt"`
	UpdatedAt                  *int64               `json:"updatedAt"`
	Revision                   *int64               `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:                core.CastString(data["namespaceId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		FollowScript:               NewScriptSettingFromDict(core.CastMap(data["followScript"])).Pointer(),
		UnfollowScript:             NewScriptSettingFromDict(core.CastMap(data["unfollowScript"])).Pointer(),
		SendRequestScript:          NewScriptSettingFromDict(core.CastMap(data["sendRequestScript"])).Pointer(),
		CancelRequestScript:        NewScriptSettingFromDict(core.CastMap(data["cancelRequestScript"])).Pointer(),
		AcceptRequestScript:        NewScriptSettingFromDict(core.CastMap(data["acceptRequestScript"])).Pointer(),
		RejectRequestScript:        NewScriptSettingFromDict(core.CastMap(data["rejectRequestScript"])).Pointer(),
		DeleteFriendScript:         NewScriptSettingFromDict(core.CastMap(data["deleteFriendScript"])).Pointer(),
		UpdateProfileScript:        NewScriptSettingFromDict(core.CastMap(data["updateProfileScript"])).Pointer(),
		FollowNotification:         NewNotificationSettingFromDict(core.CastMap(data["followNotification"])).Pointer(),
		ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
		AcceptRequestNotification:  NewNotificationSettingFromDict(core.CastMap(data["acceptRequestNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
		Revision:                   core.CastInt64(data["revision"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var followScript map[string]interface{}
	if p.FollowScript != nil {
		followScript = p.FollowScript.ToDict()
	}
	var unfollowScript map[string]interface{}
	if p.UnfollowScript != nil {
		unfollowScript = p.UnfollowScript.ToDict()
	}
	var sendRequestScript map[string]interface{}
	if p.SendRequestScript != nil {
		sendRequestScript = p.SendRequestScript.ToDict()
	}
	var cancelRequestScript map[string]interface{}
	if p.CancelRequestScript != nil {
		cancelRequestScript = p.CancelRequestScript.ToDict()
	}
	var acceptRequestScript map[string]interface{}
	if p.AcceptRequestScript != nil {
		acceptRequestScript = p.AcceptRequestScript.ToDict()
	}
	var rejectRequestScript map[string]interface{}
	if p.RejectRequestScript != nil {
		rejectRequestScript = p.RejectRequestScript.ToDict()
	}
	var deleteFriendScript map[string]interface{}
	if p.DeleteFriendScript != nil {
		deleteFriendScript = p.DeleteFriendScript.ToDict()
	}
	var updateProfileScript map[string]interface{}
	if p.UpdateProfileScript != nil {
		updateProfileScript = p.UpdateProfileScript.ToDict()
	}
	var followNotification map[string]interface{}
	if p.FollowNotification != nil {
		followNotification = p.FollowNotification.ToDict()
	}
	var receiveRequestNotification map[string]interface{}
	if p.ReceiveRequestNotification != nil {
		receiveRequestNotification = p.ReceiveRequestNotification.ToDict()
	}
	var acceptRequestNotification map[string]interface{}
	if p.AcceptRequestNotification != nil {
		acceptRequestNotification = p.AcceptRequestNotification.ToDict()
	}
	var logSetting map[string]interface{}
	if p.LogSetting != nil {
		logSetting = p.LogSetting.ToDict()
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":                namespaceId,
		"name":                       name,
		"description":                description,
		"followScript":               followScript,
		"unfollowScript":             unfollowScript,
		"sendRequestScript":          sendRequestScript,
		"cancelRequestScript":        cancelRequestScript,
		"acceptRequestScript":        acceptRequestScript,
		"rejectRequestScript":        rejectRequestScript,
		"deleteFriendScript":         deleteFriendScript,
		"updateProfileScript":        updateProfileScript,
		"followNotification":         followNotification,
		"receiveRequestNotification": receiveRequestNotification,
		"acceptRequestNotification":  acceptRequestNotification,
		"logSetting":                 logSetting,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
		"revision":                   revision,
	}
}

func (p Namespace) Pointer() *Namespace {
	return &p
}

func CastNamespaces(data []interface{}) []Namespace {
	v := make([]Namespace, 0)
	for _, d := range data {
		v = append(v, NewNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNamespacesFromDict(data []Namespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Profile struct {
	ProfileId       *string `json:"profileId"`
	UserId          *string `json:"userId"`
	PublicProfile   *string `json:"publicProfile"`
	FollowerProfile *string `json:"followerProfile"`
	FriendProfile   *string `json:"friendProfile"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
	Revision        *int64  `json:"revision"`
}

func NewProfileFromJson(data string) Profile {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewProfileFromDict(dict)
}

func NewProfileFromDict(data map[string]interface{}) Profile {
	return Profile{
		ProfileId:       core.CastString(data["profileId"]),
		UserId:          core.CastString(data["userId"]),
		PublicProfile:   core.CastString(data["publicProfile"]),
		FollowerProfile: core.CastString(data["followerProfile"]),
		FriendProfile:   core.CastString(data["friendProfile"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
		Revision:        core.CastInt64(data["revision"]),
	}
}

func (p Profile) ToDict() map[string]interface{} {

	var profileId *string
	if p.ProfileId != nil {
		profileId = p.ProfileId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var publicProfile *string
	if p.PublicProfile != nil {
		publicProfile = p.PublicProfile
	}
	var followerProfile *string
	if p.FollowerProfile != nil {
		followerProfile = p.FollowerProfile
	}
	var friendProfile *string
	if p.FriendProfile != nil {
		friendProfile = p.FriendProfile
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"profileId":       profileId,
		"userId":          userId,
		"publicProfile":   publicProfile,
		"followerProfile": followerProfile,
		"friendProfile":   friendProfile,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
		"revision":        revision,
	}
}

func (p Profile) Pointer() *Profile {
	return &p
}

func CastProfiles(data []interface{}) []Profile {
	v := make([]Profile, 0)
	for _, d := range data {
		v = append(v, NewProfileFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProfilesFromDict(data []Profile) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Follow struct {
	FollowId      *string   `json:"followId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func NewFollowFromJson(data string) Follow {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFollowFromDict(dict)
}

func NewFollowFromDict(data map[string]interface{}) Follow {
	return Follow{
		FollowId:      core.CastString(data["followId"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p Follow) ToDict() map[string]interface{} {

	var followId *string
	if p.FollowId != nil {
		followId = p.FollowId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserIds []interface{}
	if p.TargetUserIds != nil {
		targetUserIds = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"followId":      followId,
		"userId":        userId,
		"targetUserIds": targetUserIds,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p Follow) Pointer() *Follow {
	return &p
}

func CastFollows(data []interface{}) []Follow {
	v := make([]Follow, 0)
	for _, d := range data {
		v = append(v, NewFollowFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFollowsFromDict(data []Follow) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Friend struct {
	FriendId      *string   `json:"friendId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func NewFriendFromJson(data string) Friend {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendFromDict(dict)
}

func NewFriendFromDict(data map[string]interface{}) Friend {
	return Friend{
		FriendId:      core.CastString(data["friendId"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p Friend) ToDict() map[string]interface{} {

	var friendId *string
	if p.FriendId != nil {
		friendId = p.FriendId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserIds []interface{}
	if p.TargetUserIds != nil {
		targetUserIds = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"friendId":      friendId,
		"userId":        userId,
		"targetUserIds": targetUserIds,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p Friend) Pointer() *Friend {
	return &p
}

func CastFriends(data []interface{}) []Friend {
	v := make([]Friend, 0)
	for _, d := range data {
		v = append(v, NewFriendFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendsFromDict(data []Friend) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendBox struct {
	SendBoxId     *string   `json:"sendBoxId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func NewSendBoxFromJson(data string) SendBox {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendBoxFromDict(dict)
}

func NewSendBoxFromDict(data map[string]interface{}) SendBox {
	return SendBox{
		SendBoxId:     core.CastString(data["sendBoxId"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p SendBox) ToDict() map[string]interface{} {

	var sendBoxId *string
	if p.SendBoxId != nil {
		sendBoxId = p.SendBoxId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserIds []interface{}
	if p.TargetUserIds != nil {
		targetUserIds = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"sendBoxId":     sendBoxId,
		"userId":        userId,
		"targetUserIds": targetUserIds,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p SendBox) Pointer() *SendBox {
	return &p
}

func CastSendBoxes(data []interface{}) []SendBox {
	v := make([]SendBox, 0)
	for _, d := range data {
		v = append(v, NewSendBoxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendBoxesFromDict(data []SendBox) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Inbox struct {
	InboxId     *string   `json:"inboxId"`
	UserId      *string   `json:"userId"`
	FromUserIds []*string `json:"fromUserIds"`
	CreatedAt   *int64    `json:"createdAt"`
	UpdatedAt   *int64    `json:"updatedAt"`
	Revision    *int64    `json:"revision"`
}

func NewInboxFromJson(data string) Inbox {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInboxFromDict(dict)
}

func NewInboxFromDict(data map[string]interface{}) Inbox {
	return Inbox{
		InboxId:     core.CastString(data["inboxId"]),
		UserId:      core.CastString(data["userId"]),
		FromUserIds: core.CastStrings(core.CastArray(data["fromUserIds"])),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p Inbox) ToDict() map[string]interface{} {

	var inboxId *string
	if p.InboxId != nil {
		inboxId = p.InboxId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var fromUserIds []interface{}
	if p.FromUserIds != nil {
		fromUserIds = core.CastStringsFromDict(
			p.FromUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"inboxId":     inboxId,
		"userId":      userId,
		"fromUserIds": fromUserIds,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p Inbox) Pointer() *Inbox {
	return &p
}

func CastInboxes(data []interface{}) []Inbox {
	v := make([]Inbox, 0)
	for _, d := range data {
		v = append(v, NewInboxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxesFromDict(data []Inbox) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BlackList struct {
	BlackListId   *string   `json:"blackListId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func NewBlackListFromJson(data string) BlackList {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBlackListFromDict(dict)
}

func NewBlackListFromDict(data map[string]interface{}) BlackList {
	return BlackList{
		BlackListId:   core.CastString(data["blackListId"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p BlackList) ToDict() map[string]interface{} {

	var blackListId *string
	if p.BlackListId != nil {
		blackListId = p.BlackListId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserIds []interface{}
	if p.TargetUserIds != nil {
		targetUserIds = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"blackListId":   blackListId,
		"userId":        userId,
		"targetUserIds": targetUserIds,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p BlackList) Pointer() *BlackList {
	return &p
}

func CastBlackList(data []interface{}) []BlackList {
	v := make([]BlackList, 0)
	for _, d := range data {
		v = append(v, NewBlackListFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBlackListFromDict(data []BlackList) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FollowUser struct {
	UserId          *string `json:"userId"`
	PublicProfile   *string `json:"publicProfile"`
	FollowerProfile *string `json:"followerProfile"`
}

func NewFollowUserFromJson(data string) FollowUser {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFollowUserFromDict(dict)
}

func NewFollowUserFromDict(data map[string]interface{}) FollowUser {
	return FollowUser{
		UserId:          core.CastString(data["userId"]),
		PublicProfile:   core.CastString(data["publicProfile"]),
		FollowerProfile: core.CastString(data["followerProfile"]),
	}
}

func (p FollowUser) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var publicProfile *string
	if p.PublicProfile != nil {
		publicProfile = p.PublicProfile
	}
	var followerProfile *string
	if p.FollowerProfile != nil {
		followerProfile = p.FollowerProfile
	}
	return map[string]interface{}{
		"userId":          userId,
		"publicProfile":   publicProfile,
		"followerProfile": followerProfile,
	}
}

func (p FollowUser) Pointer() *FollowUser {
	return &p
}

func CastFollowUsers(data []interface{}) []FollowUser {
	v := make([]FollowUser, 0)
	for _, d := range data {
		v = append(v, NewFollowUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFollowUsersFromDict(data []FollowUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendUser struct {
	UserId        *string `json:"userId"`
	PublicProfile *string `json:"publicProfile"`
	FriendProfile *string `json:"friendProfile"`
}

func NewFriendUserFromJson(data string) FriendUser {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendUserFromDict(dict)
}

func NewFriendUserFromDict(data map[string]interface{}) FriendUser {
	return FriendUser{
		UserId:        core.CastString(data["userId"]),
		PublicProfile: core.CastString(data["publicProfile"]),
		FriendProfile: core.CastString(data["friendProfile"]),
	}
}

func (p FriendUser) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var publicProfile *string
	if p.PublicProfile != nil {
		publicProfile = p.PublicProfile
	}
	var friendProfile *string
	if p.FriendProfile != nil {
		friendProfile = p.FriendProfile
	}
	return map[string]interface{}{
		"userId":        userId,
		"publicProfile": publicProfile,
		"friendProfile": friendProfile,
	}
}

func (p FriendUser) Pointer() *FriendUser {
	return &p
}

func CastFriendUsers(data []interface{}) []FriendUser {
	v := make([]FriendUser, 0)
	for _, d := range data {
		v = append(v, NewFriendUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendUsersFromDict(data []FriendUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendRequest struct {
	UserId       *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func NewFriendRequestFromJson(data string) FriendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFriendRequestFromDict(dict)
}

func NewFriendRequestFromDict(data map[string]interface{}) FriendRequest {
	return FriendRequest{
		UserId:       core.CastString(data["userId"]),
		TargetUserId: core.CastString(data["targetUserId"]),
	}
}

func (p FriendRequest) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserId *string
	if p.TargetUserId != nil {
		targetUserId = p.TargetUserId
	}
	return map[string]interface{}{
		"userId":       userId,
		"targetUserId": targetUserId,
	}
}

func (p FriendRequest) Pointer() *FriendRequest {
	return &p
}

func CastFriendRequests(data []interface{}) []FriendRequest {
	v := make([]FriendRequest, 0)
	for _, d := range data {
		v = append(v, NewFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendRequestsFromDict(data []FriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendFriendRequest struct {
	UserId       *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func NewSendFriendRequestFromJson(data string) SendFriendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendFriendRequestFromDict(dict)
}

func NewSendFriendRequestFromDict(data map[string]interface{}) SendFriendRequest {
	return SendFriendRequest{
		UserId:       core.CastString(data["userId"]),
		TargetUserId: core.CastString(data["targetUserId"]),
	}
}

func (p SendFriendRequest) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserId *string
	if p.TargetUserId != nil {
		targetUserId = p.TargetUserId
	}
	return map[string]interface{}{
		"userId":       userId,
		"targetUserId": targetUserId,
	}
}

func (p SendFriendRequest) Pointer() *SendFriendRequest {
	return &p
}

func CastSendFriendRequests(data []interface{}) []SendFriendRequest {
	v := make([]SendFriendRequest, 0)
	for _, d := range data {
		v = append(v, NewSendFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendFriendRequestsFromDict(data []SendFriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReceiveFriendRequest struct {
	UserId       *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func NewReceiveFriendRequestFromJson(data string) ReceiveFriendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveFriendRequestFromDict(dict)
}

func NewReceiveFriendRequestFromDict(data map[string]interface{}) ReceiveFriendRequest {
	return ReceiveFriendRequest{
		UserId:       core.CastString(data["userId"]),
		TargetUserId: core.CastString(data["targetUserId"]),
	}
}

func (p ReceiveFriendRequest) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserId *string
	if p.TargetUserId != nil {
		targetUserId = p.TargetUserId
	}
	return map[string]interface{}{
		"userId":       userId,
		"targetUserId": targetUserId,
	}
}

func (p ReceiveFriendRequest) Pointer() *ReceiveFriendRequest {
	return &p
}

func CastReceiveFriendRequests(data []interface{}) []ReceiveFriendRequest {
	v := make([]ReceiveFriendRequest, 0)
	for _, d := range data {
		v = append(v, NewReceiveFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiveFriendRequestsFromDict(data []ReceiveFriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PublicProfile struct {
	UserId        *string `json:"userId"`
	PublicProfile *string `json:"publicProfile"`
}

func NewPublicProfileFromJson(data string) PublicProfile {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPublicProfileFromDict(dict)
}

func NewPublicProfileFromDict(data map[string]interface{}) PublicProfile {
	return PublicProfile{
		UserId:        core.CastString(data["userId"]),
		PublicProfile: core.CastString(data["publicProfile"]),
	}
}

func (p PublicProfile) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var publicProfile *string
	if p.PublicProfile != nil {
		publicProfile = p.PublicProfile
	}
	return map[string]interface{}{
		"userId":        userId,
		"publicProfile": publicProfile,
	}
}

func (p PublicProfile) Pointer() *PublicProfile {
	return &p
}

func CastPublicProfiles(data []interface{}) []PublicProfile {
	v := make([]PublicProfile, 0)
	for _, d := range data {
		v = append(v, NewPublicProfileFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPublicProfilesFromDict(data []PublicProfile) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewScriptSettingFromDict(dict)
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {

	var triggerScriptId *string
	if p.TriggerScriptId != nil {
		triggerScriptId = p.TriggerScriptId
	}
	var doneTriggerTargetType *string
	if p.DoneTriggerTargetType != nil {
		doneTriggerTargetType = p.DoneTriggerTargetType
	}
	var doneTriggerScriptId *string
	if p.DoneTriggerScriptId != nil {
		doneTriggerScriptId = p.DoneTriggerScriptId
	}
	var doneTriggerQueueNamespaceId *string
	if p.DoneTriggerQueueNamespaceId != nil {
		doneTriggerQueueNamespaceId = p.DoneTriggerQueueNamespaceId
	}
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
		"doneTriggerQueueNamespaceId": doneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Pointer() *ScriptSetting {
	return &p
}

func CastScriptSettings(data []interface{}) []ScriptSetting {
	v := make([]ScriptSetting, 0)
	for _, d := range data {
		v = append(v, NewScriptSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScriptSettingsFromDict(data []ScriptSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NotificationSetting struct {
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNotificationSettingFromDict(dict)
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {

	var gatewayNamespaceId *string
	if p.GatewayNamespaceId != nil {
		gatewayNamespaceId = p.GatewayNamespaceId
	}
	var enableTransferMobileNotification *bool
	if p.EnableTransferMobileNotification != nil {
		enableTransferMobileNotification = p.EnableTransferMobileNotification
	}
	var sound *string
	if p.Sound != nil {
		sound = p.Sound
	}
	return map[string]interface{}{
		"gatewayNamespaceId":               gatewayNamespaceId,
		"enableTransferMobileNotification": enableTransferMobileNotification,
		"sound":                            sound,
	}
}

func (p NotificationSetting) Pointer() *NotificationSetting {
	return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func NewLogSettingFromJson(data string) LogSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {

	var loggingNamespaceId *string
	if p.LoggingNamespaceId != nil {
		loggingNamespaceId = p.LoggingNamespaceId
	}
	return map[string]interface{}{
		"loggingNamespaceId": loggingNamespaceId,
	}
}

func (p LogSetting) Pointer() *LogSetting {
	return &p
}

func CastLogSettings(data []interface{}) []LogSetting {
	v := make([]LogSetting, 0)
	for _, d := range data {
		v = append(v, NewLogSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogSettingsFromDict(data []LogSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
