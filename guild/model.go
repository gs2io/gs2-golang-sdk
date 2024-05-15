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

package guild

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                *string              `json:"namespaceId"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
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
		JoinNotification:           NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:          NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		ChangeMemberNotification:   NewNotificationSettingFromDict(core.CastMap(data["changeMemberNotification"])).Pointer(),
		ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
		RemoveRequestNotification:  NewNotificationSettingFromDict(core.CastMap(data["removeRequestNotification"])).Pointer(),
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
	var joinNotification map[string]interface{}
	if p.JoinNotification != nil {
		joinNotification = p.JoinNotification.ToDict()
	}
	var leaveNotification map[string]interface{}
	if p.LeaveNotification != nil {
		leaveNotification = p.LeaveNotification.ToDict()
	}
	var changeMemberNotification map[string]interface{}
	if p.ChangeMemberNotification != nil {
		changeMemberNotification = p.ChangeMemberNotification.ToDict()
	}
	var receiveRequestNotification map[string]interface{}
	if p.ReceiveRequestNotification != nil {
		receiveRequestNotification = p.ReceiveRequestNotification.ToDict()
	}
	var removeRequestNotification map[string]interface{}
	if p.RemoveRequestNotification != nil {
		removeRequestNotification = p.RemoveRequestNotification.ToDict()
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
		"joinNotification":           joinNotification,
		"leaveNotification":          leaveNotification,
		"changeMemberNotification":   changeMemberNotification,
		"receiveRequestNotification": receiveRequestNotification,
		"removeRequestNotification":  removeRequestNotification,
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

type GuildModelMaster struct {
	GuildModelId              *string     `json:"guildModelId"`
	Name                      *string     `json:"name"`
	Description               *string     `json:"description"`
	Metadata                  *string     `json:"metadata"`
	DefaultMaximumMemberCount *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount        *int32      `json:"maximumMemberCount"`
	Roles                     []RoleModel `json:"roles"`
	GuildMasterRole           *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes     *int32      `json:"rejoinCoolTimeMinutes"`
	CreatedAt                 *int64      `json:"createdAt"`
	UpdatedAt                 *int64      `json:"updatedAt"`
	Revision                  *int64      `json:"revision"`
}

func NewGuildModelMasterFromJson(data string) GuildModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGuildModelMasterFromDict(dict)
}

func NewGuildModelMasterFromDict(data map[string]interface{}) GuildModelMaster {
	return GuildModelMaster{
		GuildModelId:              core.CastString(data["guildModelId"]),
		Name:                      core.CastString(data["name"]),
		Description:               core.CastString(data["description"]),
		Metadata:                  core.CastString(data["metadata"]),
		DefaultMaximumMemberCount: core.CastInt32(data["defaultMaximumMemberCount"]),
		MaximumMemberCount:        core.CastInt32(data["maximumMemberCount"]),
		Roles:                     CastRoleModels(core.CastArray(data["roles"])),
		GuildMasterRole:           core.CastString(data["guildMasterRole"]),
		GuildMemberDefaultRole:    core.CastString(data["guildMemberDefaultRole"]),
		RejoinCoolTimeMinutes:     core.CastInt32(data["rejoinCoolTimeMinutes"]),
		CreatedAt:                 core.CastInt64(data["createdAt"]),
		UpdatedAt:                 core.CastInt64(data["updatedAt"]),
		Revision:                  core.CastInt64(data["revision"]),
	}
}

func (p GuildModelMaster) ToDict() map[string]interface{} {

	var guildModelId *string
	if p.GuildModelId != nil {
		guildModelId = p.GuildModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var defaultMaximumMemberCount *int32
	if p.DefaultMaximumMemberCount != nil {
		defaultMaximumMemberCount = p.DefaultMaximumMemberCount
	}
	var maximumMemberCount *int32
	if p.MaximumMemberCount != nil {
		maximumMemberCount = p.MaximumMemberCount
	}
	var roles []interface{}
	if p.Roles != nil {
		roles = CastRoleModelsFromDict(
			p.Roles,
		)
	}
	var guildMasterRole *string
	if p.GuildMasterRole != nil {
		guildMasterRole = p.GuildMasterRole
	}
	var guildMemberDefaultRole *string
	if p.GuildMemberDefaultRole != nil {
		guildMemberDefaultRole = p.GuildMemberDefaultRole
	}
	var rejoinCoolTimeMinutes *int32
	if p.RejoinCoolTimeMinutes != nil {
		rejoinCoolTimeMinutes = p.RejoinCoolTimeMinutes
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
		"guildModelId":              guildModelId,
		"name":                      name,
		"description":               description,
		"metadata":                  metadata,
		"defaultMaximumMemberCount": defaultMaximumMemberCount,
		"maximumMemberCount":        maximumMemberCount,
		"roles":                     roles,
		"guildMasterRole":           guildMasterRole,
		"guildMemberDefaultRole":    guildMemberDefaultRole,
		"rejoinCoolTimeMinutes":     rejoinCoolTimeMinutes,
		"createdAt":                 createdAt,
		"updatedAt":                 updatedAt,
		"revision":                  revision,
	}
}

func (p GuildModelMaster) Pointer() *GuildModelMaster {
	return &p
}

func CastGuildModelMasters(data []interface{}) []GuildModelMaster {
	v := make([]GuildModelMaster, 0)
	for _, d := range data {
		v = append(v, NewGuildModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildModelMastersFromDict(data []GuildModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GuildModel struct {
	GuildModelId              *string     `json:"guildModelId"`
	Name                      *string     `json:"name"`
	Metadata                  *string     `json:"metadata"`
	DefaultMaximumMemberCount *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount        *int32      `json:"maximumMemberCount"`
	Roles                     []RoleModel `json:"roles"`
	GuildMasterRole           *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes     *int32      `json:"rejoinCoolTimeMinutes"`
}

func NewGuildModelFromJson(data string) GuildModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGuildModelFromDict(dict)
}

func NewGuildModelFromDict(data map[string]interface{}) GuildModel {
	return GuildModel{
		GuildModelId:              core.CastString(data["guildModelId"]),
		Name:                      core.CastString(data["name"]),
		Metadata:                  core.CastString(data["metadata"]),
		DefaultMaximumMemberCount: core.CastInt32(data["defaultMaximumMemberCount"]),
		MaximumMemberCount:        core.CastInt32(data["maximumMemberCount"]),
		Roles:                     CastRoleModels(core.CastArray(data["roles"])),
		GuildMasterRole:           core.CastString(data["guildMasterRole"]),
		GuildMemberDefaultRole:    core.CastString(data["guildMemberDefaultRole"]),
		RejoinCoolTimeMinutes:     core.CastInt32(data["rejoinCoolTimeMinutes"]),
	}
}

func (p GuildModel) ToDict() map[string]interface{} {

	var guildModelId *string
	if p.GuildModelId != nil {
		guildModelId = p.GuildModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var defaultMaximumMemberCount *int32
	if p.DefaultMaximumMemberCount != nil {
		defaultMaximumMemberCount = p.DefaultMaximumMemberCount
	}
	var maximumMemberCount *int32
	if p.MaximumMemberCount != nil {
		maximumMemberCount = p.MaximumMemberCount
	}
	var roles []interface{}
	if p.Roles != nil {
		roles = CastRoleModelsFromDict(
			p.Roles,
		)
	}
	var guildMasterRole *string
	if p.GuildMasterRole != nil {
		guildMasterRole = p.GuildMasterRole
	}
	var guildMemberDefaultRole *string
	if p.GuildMemberDefaultRole != nil {
		guildMemberDefaultRole = p.GuildMemberDefaultRole
	}
	var rejoinCoolTimeMinutes *int32
	if p.RejoinCoolTimeMinutes != nil {
		rejoinCoolTimeMinutes = p.RejoinCoolTimeMinutes
	}
	return map[string]interface{}{
		"guildModelId":              guildModelId,
		"name":                      name,
		"metadata":                  metadata,
		"defaultMaximumMemberCount": defaultMaximumMemberCount,
		"maximumMemberCount":        maximumMemberCount,
		"roles":                     roles,
		"guildMasterRole":           guildMasterRole,
		"guildMemberDefaultRole":    guildMemberDefaultRole,
		"rejoinCoolTimeMinutes":     rejoinCoolTimeMinutes,
	}
}

func (p GuildModel) Pointer() *GuildModel {
	return &p
}

func CastGuildModels(data []interface{}) []GuildModel {
	v := make([]GuildModel, 0)
	for _, d := range data {
		v = append(v, NewGuildModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildModelsFromDict(data []GuildModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Inbox struct {
	InboxId     *string   `json:"inboxId"`
	GuildName   *string   `json:"guildName"`
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
		GuildName:   core.CastString(data["guildName"]),
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
	var guildName *string
	if p.GuildName != nil {
		guildName = p.GuildName
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
		"guildName":   guildName,
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

type SendBox struct {
	SendBoxId        *string   `json:"sendBoxId"`
	UserId           *string   `json:"userId"`
	GuildModelName   *string   `json:"guildModelName"`
	TargetGuildNames []*string `json:"targetGuildNames"`
	CreatedAt        *int64    `json:"createdAt"`
	UpdatedAt        *int64    `json:"updatedAt"`
	Revision         *int64    `json:"revision"`
}

func NewSendBoxFromJson(data string) SendBox {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendBoxFromDict(dict)
}

func NewSendBoxFromDict(data map[string]interface{}) SendBox {
	return SendBox{
		SendBoxId:        core.CastString(data["sendBoxId"]),
		UserId:           core.CastString(data["userId"]),
		GuildModelName:   core.CastString(data["guildModelName"]),
		TargetGuildNames: core.CastStrings(core.CastArray(data["targetGuildNames"])),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
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
	var guildModelName *string
	if p.GuildModelName != nil {
		guildModelName = p.GuildModelName
	}
	var targetGuildNames []interface{}
	if p.TargetGuildNames != nil {
		targetGuildNames = core.CastStringsFromDict(
			p.TargetGuildNames,
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
		"sendBoxId":        sendBoxId,
		"userId":           userId,
		"guildModelName":   guildModelName,
		"targetGuildNames": targetGuildNames,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
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

type Guild struct {
	GuildId                   *string     `json:"guildId"`
	GuildModelName            *string     `json:"guildModelName"`
	Name                      *string     `json:"name"`
	DisplayName               *string     `json:"displayName"`
	Attribute1                *int32      `json:"attribute1"`
	Attribute2                *int32      `json:"attribute2"`
	Attribute3                *int32      `json:"attribute3"`
	Attribute4                *int32      `json:"attribute4"`
	Attribute5                *int32      `json:"attribute5"`
	JoinPolicy                *string     `json:"joinPolicy"`
	CustomRoles               []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	CurrentMaximumMemberCount *int32      `json:"currentMaximumMemberCount"`
	Members                   []Member    `json:"members"`
	CreatedAt                 *int64      `json:"createdAt"`
	UpdatedAt                 *int64      `json:"updatedAt"`
	Revision                  *int64      `json:"revision"`
}

func NewGuildFromJson(data string) Guild {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGuildFromDict(dict)
}

func NewGuildFromDict(data map[string]interface{}) Guild {
	return Guild{
		GuildId:                   core.CastString(data["guildId"]),
		GuildModelName:            core.CastString(data["guildModelName"]),
		Name:                      core.CastString(data["name"]),
		DisplayName:               core.CastString(data["displayName"]),
		Attribute1:                core.CastInt32(data["attribute1"]),
		Attribute2:                core.CastInt32(data["attribute2"]),
		Attribute3:                core.CastInt32(data["attribute3"]),
		Attribute4:                core.CastInt32(data["attribute4"]),
		Attribute5:                core.CastInt32(data["attribute5"]),
		JoinPolicy:                core.CastString(data["joinPolicy"]),
		CustomRoles:               CastRoleModels(core.CastArray(data["customRoles"])),
		GuildMemberDefaultRole:    core.CastString(data["guildMemberDefaultRole"]),
		CurrentMaximumMemberCount: core.CastInt32(data["currentMaximumMemberCount"]),
		Members:                   CastMembers(core.CastArray(data["members"])),
		CreatedAt:                 core.CastInt64(data["createdAt"]),
		UpdatedAt:                 core.CastInt64(data["updatedAt"]),
		Revision:                  core.CastInt64(data["revision"]),
	}
}

func (p Guild) ToDict() map[string]interface{} {

	var guildId *string
	if p.GuildId != nil {
		guildId = p.GuildId
	}
	var guildModelName *string
	if p.GuildModelName != nil {
		guildModelName = p.GuildModelName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var displayName *string
	if p.DisplayName != nil {
		displayName = p.DisplayName
	}
	var attribute1 *int32
	if p.Attribute1 != nil {
		attribute1 = p.Attribute1
	}
	var attribute2 *int32
	if p.Attribute2 != nil {
		attribute2 = p.Attribute2
	}
	var attribute3 *int32
	if p.Attribute3 != nil {
		attribute3 = p.Attribute3
	}
	var attribute4 *int32
	if p.Attribute4 != nil {
		attribute4 = p.Attribute4
	}
	var attribute5 *int32
	if p.Attribute5 != nil {
		attribute5 = p.Attribute5
	}
	var joinPolicy *string
	if p.JoinPolicy != nil {
		joinPolicy = p.JoinPolicy
	}
	var customRoles []interface{}
	if p.CustomRoles != nil {
		customRoles = CastRoleModelsFromDict(
			p.CustomRoles,
		)
	}
	var guildMemberDefaultRole *string
	if p.GuildMemberDefaultRole != nil {
		guildMemberDefaultRole = p.GuildMemberDefaultRole
	}
	var currentMaximumMemberCount *int32
	if p.CurrentMaximumMemberCount != nil {
		currentMaximumMemberCount = p.CurrentMaximumMemberCount
	}
	var members []interface{}
	if p.Members != nil {
		members = CastMembersFromDict(
			p.Members,
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
		"guildId":                   guildId,
		"guildModelName":            guildModelName,
		"name":                      name,
		"displayName":               displayName,
		"attribute1":                attribute1,
		"attribute2":                attribute2,
		"attribute3":                attribute3,
		"attribute4":                attribute4,
		"attribute5":                attribute5,
		"joinPolicy":                joinPolicy,
		"customRoles":               customRoles,
		"guildMemberDefaultRole":    guildMemberDefaultRole,
		"currentMaximumMemberCount": currentMaximumMemberCount,
		"members":                   members,
		"createdAt":                 createdAt,
		"updatedAt":                 updatedAt,
		"revision":                  revision,
	}
}

func (p Guild) Pointer() *Guild {
	return &p
}

func CastGuilds(data []interface{}) []Guild {
	v := make([]Guild, 0)
	for _, d := range data {
		v = append(v, NewGuildFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildsFromDict(data []Guild) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type JoinedGuild struct {
	JoinedGuildId  *string `json:"joinedGuildId"`
	GuildModelName *string `json:"guildModelName"`
	GuildName      *string `json:"guildName"`
	UserId         *string `json:"userId"`
	CreatedAt      *int64  `json:"createdAt"`
}

func NewJoinedGuildFromJson(data string) JoinedGuild {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewJoinedGuildFromDict(dict)
}

func NewJoinedGuildFromDict(data map[string]interface{}) JoinedGuild {
	return JoinedGuild{
		JoinedGuildId:  core.CastString(data["joinedGuildId"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		UserId:         core.CastString(data["userId"]),
		CreatedAt:      core.CastInt64(data["createdAt"]),
	}
}

func (p JoinedGuild) ToDict() map[string]interface{} {

	var joinedGuildId *string
	if p.JoinedGuildId != nil {
		joinedGuildId = p.JoinedGuildId
	}
	var guildModelName *string
	if p.GuildModelName != nil {
		guildModelName = p.GuildModelName
	}
	var guildName *string
	if p.GuildName != nil {
		guildName = p.GuildName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"joinedGuildId":  joinedGuildId,
		"guildModelName": guildModelName,
		"guildName":      guildName,
		"userId":         userId,
		"createdAt":      createdAt,
	}
}

func (p JoinedGuild) Pointer() *JoinedGuild {
	return &p
}

func CastJoinedGuilds(data []interface{}) []JoinedGuild {
	v := make([]JoinedGuild, 0)
	for _, d := range data {
		v = append(v, NewJoinedGuildFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJoinedGuildsFromDict(data []JoinedGuild) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentGuildMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentGuildMasterFromJson(data string) CurrentGuildMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentGuildMasterFromDict(dict)
}

func NewCurrentGuildMasterFromDict(data map[string]interface{}) CurrentGuildMaster {
	return CurrentGuildMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentGuildMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
}

func (p CurrentGuildMaster) Pointer() *CurrentGuildMaster {
	return &p
}

func CastCurrentGuildMasters(data []interface{}) []CurrentGuildMaster {
	v := make([]CurrentGuildMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentGuildMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentGuildMastersFromDict(data []CurrentGuildMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RoleModel struct {
	Name           *string `json:"name"`
	Metadata       *string `json:"metadata"`
	PolicyDocument *string `json:"policyDocument"`
}

func NewRoleModelFromJson(data string) RoleModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRoleModelFromDict(dict)
}

func NewRoleModelFromDict(data map[string]interface{}) RoleModel {
	return RoleModel{
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		PolicyDocument: core.CastString(data["policyDocument"]),
	}
}

func (p RoleModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var policyDocument *string
	if p.PolicyDocument != nil {
		policyDocument = p.PolicyDocument
	}
	return map[string]interface{}{
		"name":           name,
		"metadata":       metadata,
		"policyDocument": policyDocument,
	}
}

func (p RoleModel) Pointer() *RoleModel {
	return &p
}

func CastRoleModels(data []interface{}) []RoleModel {
	v := make([]RoleModel, 0)
	for _, d := range data {
		v = append(v, NewRoleModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRoleModelsFromDict(data []RoleModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Member struct {
	UserId   *string `json:"userId"`
	RoleName *string `json:"roleName"`
	JoinedAt *int64  `json:"joinedAt"`
}

func NewMemberFromJson(data string) Member {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMemberFromDict(dict)
}

func NewMemberFromDict(data map[string]interface{}) Member {
	return Member{
		UserId:   core.CastString(data["userId"]),
		RoleName: core.CastString(data["roleName"]),
		JoinedAt: core.CastInt64(data["joinedAt"]),
	}
}

func (p Member) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var roleName *string
	if p.RoleName != nil {
		roleName = p.RoleName
	}
	var joinedAt *int64
	if p.JoinedAt != nil {
		joinedAt = p.JoinedAt
	}
	return map[string]interface{}{
		"userId":   userId,
		"roleName": roleName,
		"joinedAt": joinedAt,
	}
}

func (p Member) Pointer() *Member {
	return &p
}

func CastMembers(data []interface{}) []Member {
	v := make([]Member, 0)
	for _, d := range data {
		v = append(v, NewMemberFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMembersFromDict(data []Member) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReceiveMemberRequest struct {
	UserId          *string `json:"userId"`
	TargetGuildName *string `json:"targetGuildName"`
}

func NewReceiveMemberRequestFromJson(data string) ReceiveMemberRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveMemberRequestFromDict(dict)
}

func NewReceiveMemberRequestFromDict(data map[string]interface{}) ReceiveMemberRequest {
	return ReceiveMemberRequest{
		UserId:          core.CastString(data["userId"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
	}
}

func (p ReceiveMemberRequest) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetGuildName *string
	if p.TargetGuildName != nil {
		targetGuildName = p.TargetGuildName
	}
	return map[string]interface{}{
		"userId":          userId,
		"targetGuildName": targetGuildName,
	}
}

func (p ReceiveMemberRequest) Pointer() *ReceiveMemberRequest {
	return &p
}

func CastReceiveMemberRequests(data []interface{}) []ReceiveMemberRequest {
	v := make([]ReceiveMemberRequest, 0)
	for _, d := range data {
		v = append(v, NewReceiveMemberRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiveMemberRequestsFromDict(data []ReceiveMemberRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendMemberRequest struct {
	UserId          *string `json:"userId"`
	TargetGuildName *string `json:"targetGuildName"`
}

func NewSendMemberRequestFromJson(data string) SendMemberRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendMemberRequestFromDict(dict)
}

func NewSendMemberRequestFromDict(data map[string]interface{}) SendMemberRequest {
	return SendMemberRequest{
		UserId:          core.CastString(data["userId"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
	}
}

func (p SendMemberRequest) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetGuildName *string
	if p.TargetGuildName != nil {
		targetGuildName = p.TargetGuildName
	}
	return map[string]interface{}{
		"userId":          userId,
		"targetGuildName": targetGuildName,
	}
}

func (p SendMemberRequest) Pointer() *SendMemberRequest {
	return &p
}

func CastSendMemberRequests(data []interface{}) []SendMemberRequest {
	v := make([]SendMemberRequest, 0)
	for _, d := range data {
		v = append(v, NewSendMemberRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendMemberRequestsFromDict(data []SendMemberRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTransactionSettingFromDict(dict)
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
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

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGitHubCheckoutSettingFromDict(dict)
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId:       core.CastString(data["apiKeyId"]),
		RepositoryName: core.CastString(data["repositoryName"]),
		SourcePath:     core.CastString(data["sourcePath"]),
		ReferenceType:  core.CastString(data["referenceType"]),
		CommitHash:     core.CastString(data["commitHash"]),
		BranchName:     core.CastString(data["branchName"]),
		TagName:        core.CastString(data["tagName"]),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
}

func (p GitHubCheckoutSetting) Pointer() *GitHubCheckoutSetting {
	return &p
}

func CastGitHubCheckoutSettings(data []interface{}) []GitHubCheckoutSetting {
	v := make([]GitHubCheckoutSetting, 0)
	for _, d := range data {
		v = append(v, NewGitHubCheckoutSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubCheckoutSettingsFromDict(data []GitHubCheckoutSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
