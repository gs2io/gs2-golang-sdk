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

package inbox

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                *string              `json:"namespaceId"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	QueueNamespaceId           *string              `json:"queueNamespaceId"`
	KeyId                      *string              `json:"keyId"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	CreatedAt                  *int64               `json:"createdAt"`
	UpdatedAt                  *int64               `json:"updatedAt"`
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
		IsAutomaticDeletingEnabled: core.CastBool(data["isAutomaticDeletingEnabled"]),
		ReceiveMessageScript:       NewScriptSettingFromDict(core.CastMap(data["receiveMessageScript"])).Pointer(),
		ReadMessageScript:          NewScriptSettingFromDict(core.CastMap(data["readMessageScript"])).Pointer(),
		DeleteMessageScript:        NewScriptSettingFromDict(core.CastMap(data["deleteMessageScript"])).Pointer(),
		QueueNamespaceId:           core.CastString(data["queueNamespaceId"]),
		KeyId:                      core.CastString(data["keyId"]),
		ReceiveNotification:        NewNotificationSettingFromDict(core.CastMap(data["receiveNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
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
	var isAutomaticDeletingEnabled *bool
	if p.IsAutomaticDeletingEnabled != nil {
		isAutomaticDeletingEnabled = p.IsAutomaticDeletingEnabled
	}
	var receiveMessageScript map[string]interface{}
	if p.ReceiveMessageScript != nil {
		receiveMessageScript = p.ReceiveMessageScript.ToDict()
	}
	var readMessageScript map[string]interface{}
	if p.ReadMessageScript != nil {
		readMessageScript = p.ReadMessageScript.ToDict()
	}
	var deleteMessageScript map[string]interface{}
	if p.DeleteMessageScript != nil {
		deleteMessageScript = p.DeleteMessageScript.ToDict()
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var receiveNotification map[string]interface{}
	if p.ReceiveNotification != nil {
		receiveNotification = p.ReceiveNotification.ToDict()
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
	return map[string]interface{}{
		"namespaceId":                namespaceId,
		"name":                       name,
		"description":                description,
		"isAutomaticDeletingEnabled": isAutomaticDeletingEnabled,
		"receiveMessageScript":       receiveMessageScript,
		"readMessageScript":          readMessageScript,
		"deleteMessageScript":        deleteMessageScript,
		"queueNamespaceId":           queueNamespaceId,
		"keyId":                      keyId,
		"receiveNotification":        receiveNotification,
		"logSetting":                 logSetting,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
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

type Message struct {
	MessageId          *string         `json:"messageId"`
	Name               *string         `json:"name"`
	UserId             *string         `json:"userId"`
	Metadata           *string         `json:"metadata"`
	IsRead             *bool           `json:"isRead"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ReceivedAt         *int64          `json:"receivedAt"`
	ReadAt             *int64          `json:"readAt"`
	ExpiresAt          *int64          `json:"expiresAt"`
}

func NewMessageFromJson(data string) Message {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMessageFromDict(dict)
}

func NewMessageFromDict(data map[string]interface{}) Message {
	return Message{
		MessageId:          core.CastString(data["messageId"]),
		Name:               core.CastString(data["name"]),
		UserId:             core.CastString(data["userId"]),
		Metadata:           core.CastString(data["metadata"]),
		IsRead:             core.CastBool(data["isRead"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ReceivedAt:         core.CastInt64(data["receivedAt"]),
		ReadAt:             core.CastInt64(data["readAt"]),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
	}
}

func (p Message) ToDict() map[string]interface{} {

	var messageId *string
	if p.MessageId != nil {
		messageId = p.MessageId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var isRead *bool
	if p.IsRead != nil {
		isRead = p.IsRead
	}
	var readAcquireActions []interface{}
	if p.ReadAcquireActions != nil {
		readAcquireActions = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	var receivedAt *int64
	if p.ReceivedAt != nil {
		receivedAt = p.ReceivedAt
	}
	var readAt *int64
	if p.ReadAt != nil {
		readAt = p.ReadAt
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	return map[string]interface{}{
		"messageId":          messageId,
		"name":               name,
		"userId":             userId,
		"metadata":           metadata,
		"isRead":             isRead,
		"readAcquireActions": readAcquireActions,
		"receivedAt":         receivedAt,
		"readAt":             readAt,
		"expiresAt":          expiresAt,
	}
}

func (p Message) Pointer() *Message {
	return &p
}

func CastMessages(data []interface{}) []Message {
	v := make([]Message, 0)
	for _, d := range data {
		v = append(v, NewMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMessagesFromDict(data []Message) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentMessageMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentMessageMasterFromJson(data string) CurrentMessageMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentMessageMasterFromDict(dict)
}

func NewCurrentMessageMasterFromDict(data map[string]interface{}) CurrentMessageMaster {
	return CurrentMessageMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentMessageMaster) ToDict() map[string]interface{} {

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

func (p CurrentMessageMaster) Pointer() *CurrentMessageMaster {
	return &p
}

func CastCurrentMessageMasters(data []interface{}) []CurrentMessageMaster {
	v := make([]CurrentMessageMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentMessageMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentMessageMastersFromDict(data []CurrentMessageMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalMessageMaster struct {
	GlobalMessageId    *string         `json:"globalMessageId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	CreatedAt          *int64          `json:"createdAt"`
	ExpiresAt          *int64          `json:"expiresAt"`
}

func NewGlobalMessageMasterFromJson(data string) GlobalMessageMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGlobalMessageMasterFromDict(dict)
}

func NewGlobalMessageMasterFromDict(data map[string]interface{}) GlobalMessageMaster {
	return GlobalMessageMaster{
		GlobalMessageId:    core.CastString(data["globalMessageId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ExpiresTimeSpan:    NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
	}
}

func (p GlobalMessageMaster) ToDict() map[string]interface{} {

	var globalMessageId *string
	if p.GlobalMessageId != nil {
		globalMessageId = p.GlobalMessageId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var readAcquireActions []interface{}
	if p.ReadAcquireActions != nil {
		readAcquireActions = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	var expiresTimeSpan map[string]interface{}
	if p.ExpiresTimeSpan != nil {
		expiresTimeSpan = p.ExpiresTimeSpan.ToDict()
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	return map[string]interface{}{
		"globalMessageId":    globalMessageId,
		"name":               name,
		"metadata":           metadata,
		"readAcquireActions": readAcquireActions,
		"expiresTimeSpan":    expiresTimeSpan,
		"createdAt":          createdAt,
		"expiresAt":          expiresAt,
	}
}

func (p GlobalMessageMaster) Pointer() *GlobalMessageMaster {
	return &p
}

func CastGlobalMessageMasters(data []interface{}) []GlobalMessageMaster {
	v := make([]GlobalMessageMaster, 0)
	for _, d := range data {
		v = append(v, NewGlobalMessageMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalMessageMastersFromDict(data []GlobalMessageMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalMessage struct {
	GlobalMessageId    *string         `json:"globalMessageId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	ExpiresAt          *int64          `json:"expiresAt"`
}

func NewGlobalMessageFromJson(data string) GlobalMessage {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGlobalMessageFromDict(dict)
}

func NewGlobalMessageFromDict(data map[string]interface{}) GlobalMessage {
	return GlobalMessage{
		GlobalMessageId:    core.CastString(data["globalMessageId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		ReadAcquireActions: CastAcquireActions(core.CastArray(data["readAcquireActions"])),
		ExpiresTimeSpan:    NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer(),
		ExpiresAt:          core.CastInt64(data["expiresAt"]),
	}
}

func (p GlobalMessage) ToDict() map[string]interface{} {

	var globalMessageId *string
	if p.GlobalMessageId != nil {
		globalMessageId = p.GlobalMessageId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var readAcquireActions []interface{}
	if p.ReadAcquireActions != nil {
		readAcquireActions = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	var expiresTimeSpan map[string]interface{}
	if p.ExpiresTimeSpan != nil {
		expiresTimeSpan = p.ExpiresTimeSpan.ToDict()
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	return map[string]interface{}{
		"globalMessageId":    globalMessageId,
		"name":               name,
		"metadata":           metadata,
		"readAcquireActions": readAcquireActions,
		"expiresTimeSpan":    expiresTimeSpan,
		"expiresAt":          expiresAt,
	}
}

func (p GlobalMessage) Pointer() *GlobalMessage {
	return &p
}

func CastGlobalMessages(data []interface{}) []GlobalMessage {
	v := make([]GlobalMessage, 0)
	for _, d := range data {
		v = append(v, NewGlobalMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalMessagesFromDict(data []GlobalMessage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Received struct {
	ReceivedId                 *string  `json:"receivedId"`
	UserId                     *string  `json:"userId"`
	ReceivedGlobalMessageNames []string `json:"receivedGlobalMessageNames"`
	CreatedAt                  *int64   `json:"createdAt"`
	UpdatedAt                  *int64   `json:"updatedAt"`
}

func NewReceivedFromJson(data string) Received {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceivedFromDict(dict)
}

func NewReceivedFromDict(data map[string]interface{}) Received {
	return Received{
		ReceivedId:                 core.CastString(data["receivedId"]),
		UserId:                     core.CastString(data["userId"]),
		ReceivedGlobalMessageNames: core.CastStrings(core.CastArray(data["receivedGlobalMessageNames"])),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
	}
}

func (p Received) ToDict() map[string]interface{} {

	var receivedId *string
	if p.ReceivedId != nil {
		receivedId = p.ReceivedId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var receivedGlobalMessageNames []interface{}
	if p.ReceivedGlobalMessageNames != nil {
		receivedGlobalMessageNames = core.CastStringsFromDict(
			p.ReceivedGlobalMessageNames,
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
	return map[string]interface{}{
		"receivedId":                 receivedId,
		"userId":                     userId,
		"receivedGlobalMessageNames": receivedGlobalMessageNames,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
	}
}

func (p Received) Pointer() *Received {
	return &p
}

func CastReceiveds(data []interface{}) []Received {
	v := make([]Received, 0)
	for _, d := range data {
		v = append(v, NewReceivedFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceivedsFromDict(data []Received) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Config struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func NewConfigFromJson(data string) Config {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConfigFromDict(dict)
}

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Config) ToDict() map[string]interface{} {

	var key *string
	if p.Key != nil {
		key = p.Key
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"key":   key,
		"value": value,
	}
}

func (p Config) Pointer() *Config {
	return &p
}

func CastConfigs(data []interface{}) []Config {
	v := make([]Config, 0)
	for _, d := range data {
		v = append(v, NewConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConfigsFromDict(data []Config) []interface{} {
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

type TimeSpan struct {
	Days    *int32 `json:"days"`
	Hours   *int32 `json:"hours"`
	Minutes *int32 `json:"minutes"`
}

func NewTimeSpanFromJson(data string) TimeSpan {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTimeSpanFromDict(dict)
}

func NewTimeSpanFromDict(data map[string]interface{}) TimeSpan {
	return TimeSpan{
		Days:    core.CastInt32(data["days"]),
		Hours:   core.CastInt32(data["hours"]),
		Minutes: core.CastInt32(data["minutes"]),
	}
}

func (p TimeSpan) ToDict() map[string]interface{} {

	var days *int32
	if p.Days != nil {
		days = p.Days
	}
	var hours *int32
	if p.Hours != nil {
		hours = p.Hours
	}
	var minutes *int32
	if p.Minutes != nil {
		minutes = p.Minutes
	}
	return map[string]interface{}{
		"days":    days,
		"hours":   hours,
		"minutes": minutes,
	}
}

func (p TimeSpan) Pointer() *TimeSpan {
	return &p
}

func CastTimeSpans(data []interface{}) []TimeSpan {
	v := make([]TimeSpan, 0)
	for _, d := range data {
		v = append(v, NewTimeSpanFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeSpansFromDict(data []TimeSpan) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewAcquireActionFromJson(data string) AcquireAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionFromDict(dict)
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p AcquireAction) Pointer() *AcquireAction {
	return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
