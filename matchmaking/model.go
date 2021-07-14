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

package matchmaking

import "core"

type Namespace struct {
	NamespaceId                                   *string              `json:"namespaceId"`
	Name                                          *string              `json:"name"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
	CreatedAt                                     *int64               `json:"createdAt"`
	UpdatedAt                                     *int64               `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:                core.CastString(data["namespaceId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		EnableRating:               core.CastBool(data["enableRating"]),
		CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                                     core.CastInt64(data["createdAt"]),
		UpdatedAt:                                     core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":                p.NamespaceId,
		"name":                       p.Name,
		"description":                p.Description,
		"enableRating":               p.EnableRating,
		"createGatheringTriggerType": p.CreateGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     p.CreateGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                p.CreateGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                p.CompleteMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            p.CompleteMatchmakingTriggerScriptId,
		"joinNotification":                              p.JoinNotification.ToDict(),
		"leaveNotification":                             p.LeaveNotification.ToDict(),
		"completeNotification":                          p.CompleteNotification.ToDict(),
		"logSetting":                                    p.LogSetting.ToDict(),
		"createdAt":                                     p.CreatedAt,
		"updatedAt":                                     p.UpdatedAt,
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

type Gathering struct {
	GatheringId     *string          `json:"gatheringId"`
	Name            *string          `json:"name"`
	AttributeRanges []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds    []string         `json:"allowUserIds"`
	Metadata        *string          `json:"metadata"`
	ExpiresAt       *int64           `json:"expiresAt"`
	CreatedAt       *int64           `json:"createdAt"`
	UpdatedAt       *int64           `json:"updatedAt"`
}

func NewGatheringFromDict(data map[string]interface{}) Gathering {
	return Gathering{
		GatheringId:     core.CastString(data["gatheringId"]),
		Name:            core.CastString(data["name"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles: CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:    core.CastStrings(core.CastArray(data["allowUserIds"])),
		Metadata:        core.CastString(data["metadata"]),
		ExpiresAt:       core.CastInt64(data["expiresAt"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
	}
}

func (p Gathering) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"gatheringId": p.GatheringId,
		"name":        p.Name,
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"metadata":  p.Metadata,
		"expiresAt": p.ExpiresAt,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p Gathering) Pointer() *Gathering {
	return &p
}

func CastGatherings(data []interface{}) []Gathering {
	v := make([]Gathering, 0)
	for _, d := range data {
		v = append(v, NewGatheringFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGatheringsFromDict(data []Gathering) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RatingModelMaster struct {
	RatingModelId *string `json:"ratingModelId"`
	Name          *string `json:"name"`
	Metadata      *string `json:"metadata"`
	Description   *string `json:"description"`
	Volatility    *int32  `json:"volatility"`
	CreatedAt     *int64  `json:"createdAt"`
	UpdatedAt     *int64  `json:"updatedAt"`
}

func NewRatingModelMasterFromDict(data map[string]interface{}) RatingModelMaster {
	return RatingModelMaster{
		RatingModelId: core.CastString(data["ratingModelId"]),
		Name:          core.CastString(data["name"]),
		Metadata:      core.CastString(data["metadata"]),
		Description:   core.CastString(data["description"]),
		Volatility:    core.CastInt32(data["volatility"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
	}
}

func (p RatingModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"ratingModelId": p.RatingModelId,
		"name":          p.Name,
		"metadata":      p.Metadata,
		"description":   p.Description,
		"volatility":    p.Volatility,
		"createdAt":     p.CreatedAt,
		"updatedAt":     p.UpdatedAt,
	}
}

func (p RatingModelMaster) Pointer() *RatingModelMaster {
	return &p
}

func CastRatingModelMasters(data []interface{}) []RatingModelMaster {
	v := make([]RatingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRatingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelMastersFromDict(data []RatingModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RatingModel struct {
	RatingModelId *string `json:"ratingModelId"`
	Name          *string `json:"name"`
	Metadata      *string `json:"metadata"`
	Volatility    *int32  `json:"volatility"`
}

func NewRatingModelFromDict(data map[string]interface{}) RatingModel {
	return RatingModel{
		RatingModelId: core.CastString(data["ratingModelId"]),
		Name:          core.CastString(data["name"]),
		Metadata:      core.CastString(data["metadata"]),
		Volatility:    core.CastInt32(data["volatility"]),
	}
}

func (p RatingModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"ratingModelId": p.RatingModelId,
		"name":          p.Name,
		"metadata":      p.Metadata,
		"volatility":    p.Volatility,
	}
}

func (p RatingModel) Pointer() *RatingModel {
	return &p
}

func CastRatingModels(data []interface{}) []RatingModel {
	v := make([]RatingModel, 0)
	for _, d := range data {
		v = append(v, NewRatingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelsFromDict(data []RatingModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentRatingModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentRatingModelMasterFromDict(data map[string]interface{}) CurrentRatingModelMaster {
	return CurrentRatingModelMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRatingModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentRatingModelMaster) Pointer() *CurrentRatingModelMaster {
	return &p
}

func CastCurrentRatingModelMasters(data []interface{}) []CurrentRatingModelMaster {
	v := make([]CurrentRatingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentRatingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentRatingModelMastersFromDict(data []CurrentRatingModelMaster) []interface{} {
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

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"gatewayNamespaceId":               p.GatewayNamespaceId,
		"enableTransferMobileNotification": p.EnableTransferMobileNotification,
		"sound":                            p.Sound,
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

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
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
	return map[string]interface{}{
		"apiKeyId":       p.ApiKeyId,
		"repositoryName": p.RepositoryName,
		"sourcePath":     p.SourcePath,
		"referenceType":  p.ReferenceType,
		"commitHash":     p.CommitHash,
		"branchName":     p.BranchName,
		"tagName":        p.TagName,
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

type AttributeRange struct {
	Name *string `json:"name"`
	Min  *int32  `json:"min"`
	Max  *int32  `json:"max"`
}

func NewAttributeRangeFromDict(data map[string]interface{}) AttributeRange {
	return AttributeRange{
		Name: core.CastString(data["name"]),
		Min:  core.CastInt32(data["min"]),
		Max:  core.CastInt32(data["max"]),
	}
}

func (p AttributeRange) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name": p.Name,
		"min":  p.Min,
		"max":  p.Max,
	}
}

func (p AttributeRange) Pointer() *AttributeRange {
	return &p
}

func CastAttributeRanges(data []interface{}) []AttributeRange {
	v := make([]AttributeRange, 0)
	for _, d := range data {
		v = append(v, NewAttributeRangeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributeRangesFromDict(data []AttributeRange) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CapacityOfRole struct {
	RoleName     *string  `json:"roleName"`
	RoleAliases  []string `json:"roleAliases"`
	Capacity     *int32   `json:"capacity"`
	Participants []Player `json:"participants"`
}

func NewCapacityOfRoleFromDict(data map[string]interface{}) CapacityOfRole {
	return CapacityOfRole{
		RoleName:     core.CastString(data["roleName"]),
		RoleAliases:  core.CastStrings(core.CastArray(data["roleAliases"])),
		Capacity:     core.CastInt32(data["capacity"]),
		Participants: CastPlayers(core.CastArray(data["participants"])),
	}
}

func (p CapacityOfRole) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"roleName": p.RoleName,
		"roleAliases": core.CastStringsFromDict(
			p.RoleAliases,
		),
		"capacity": p.Capacity,
		"participants": CastPlayersFromDict(
			p.Participants,
		),
	}
}

func (p CapacityOfRole) Pointer() *CapacityOfRole {
	return &p
}

func CastCapacityOfRoles(data []interface{}) []CapacityOfRole {
	v := make([]CapacityOfRole, 0)
	for _, d := range data {
		v = append(v, NewCapacityOfRoleFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCapacityOfRolesFromDict(data []CapacityOfRole) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Attribute struct {
	Name  *string `json:"name"`
	Value *int32  `json:"value"`
}

func NewAttributeFromDict(data map[string]interface{}) Attribute {
	return Attribute{
		Name:  core.CastString(data["name"]),
		Value: core.CastInt32(data["value"]),
	}
}

func (p Attribute) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":  p.Name,
		"value": p.Value,
	}
}

func (p Attribute) Pointer() *Attribute {
	return &p
}

func CastAttributes(data []interface{}) []Attribute {
	v := make([]Attribute, 0)
	for _, d := range data {
		v = append(v, NewAttributeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributesFromDict(data []Attribute) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Player struct {
	UserId      *string     `json:"userId"`
	Attributes  []Attribute `json:"attributes"`
	RoleName    *string     `json:"roleName"`
	DenyUserIds []string    `json:"denyUserIds"`
}

func NewPlayerFromDict(data map[string]interface{}) Player {
	return Player{
		UserId:      core.CastString(data["userId"]),
		Attributes:  CastAttributes(core.CastArray(data["attributes"])),
		RoleName:    core.CastString(data["roleName"]),
		DenyUserIds: core.CastStrings(core.CastArray(data["denyUserIds"])),
	}
}

func (p Player) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
		"attributes": CastAttributesFromDict(
			p.Attributes,
		),
		"roleName": p.RoleName,
		"denyUserIds": core.CastStringsFromDict(
			p.DenyUserIds,
		),
	}
}

func (p Player) Pointer() *Player {
	return &p
}

func CastPlayers(data []interface{}) []Player {
	v := make([]Player, 0)
	for _, d := range data {
		v = append(v, NewPlayerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlayersFromDict(data []Player) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Rating struct {
	RatingId  *string  `json:"ratingId"`
	Name      *string  `json:"name"`
	UserId    *string  `json:"userId"`
	RateValue *float32 `json:"rateValue"`
	CreatedAt *int64   `json:"createdAt"`
	UpdatedAt *int64   `json:"updatedAt"`
}

func NewRatingFromDict(data map[string]interface{}) Rating {
	return Rating{
		RatingId:  core.CastString(data["ratingId"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		RateValue: core.CastFloat32(data["rateValue"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
	}
}

func (p Rating) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"ratingId":  p.RatingId,
		"name":      p.Name,
		"userId":    p.UserId,
		"rateValue": p.RateValue,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p Rating) Pointer() *Rating {
	return &p
}

func CastRatings(data []interface{}) []Rating {
	v := make([]Rating, 0)
	for _, d := range data {
		v = append(v, NewRatingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingsFromDict(data []Rating) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GameResult struct {
	Rank   *int32  `json:"rank"`
	UserId *string `json:"userId"`
}

func NewGameResultFromDict(data map[string]interface{}) GameResult {
	return GameResult{
		Rank:   core.CastInt32(data["rank"]),
		UserId: core.CastString(data["userId"]),
	}
}

func (p GameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rank":   p.Rank,
		"userId": p.UserId,
	}
}

func (p GameResult) Pointer() *GameResult {
	return &p
}

func CastGameResults(data []interface{}) []GameResult {
	v := make([]GameResult, 0)
	for _, d := range data {
		v = append(v, NewGameResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGameResultsFromDict(data []GameResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Ballot struct {
	UserId         *string `json:"userId"`
	RatingName     *string `json:"ratingName"`
	GatheringName  *string `json:"gatheringName"`
	NumberOfPlayer *int32  `json:"numberOfPlayer"`
}

func NewBallotFromDict(data map[string]interface{}) Ballot {
	return Ballot{
		UserId:         core.CastString(data["userId"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
	}
}

func (p Ballot) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":         p.UserId,
		"ratingName":     p.RatingName,
		"gatheringName":  p.GatheringName,
		"numberOfPlayer": p.NumberOfPlayer,
	}
}

func (p Ballot) Pointer() *Ballot {
	return &p
}

func CastBallots(data []interface{}) []Ballot {
	v := make([]Ballot, 0)
	for _, d := range data {
		v = append(v, NewBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBallotsFromDict(data []Ballot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SignedBallot struct {
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

func NewSignedBallotFromDict(data map[string]interface{}) SignedBallot {
	return SignedBallot{
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p SignedBallot) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p SignedBallot) Pointer() *SignedBallot {
	return &p
}

func CastSignedBallots(data []interface{}) []SignedBallot {
	v := make([]SignedBallot, 0)
	for _, d := range data {
		v = append(v, NewSignedBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignedBallotsFromDict(data []SignedBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type WrittenBallot struct {
	Ballot      *Ballot      `json:"ballot"`
	GameResults []GameResult `json:"gameResults"`
}

func NewWrittenBallotFromDict(data map[string]interface{}) WrittenBallot {
	return WrittenBallot{
		Ballot:      NewBallotFromDict(core.CastMap(data["ballot"])).Pointer(),
		GameResults: CastGameResults(core.CastArray(data["gameResults"])),
	}
}

func (p WrittenBallot) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"ballot": p.Ballot.ToDict(),
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
	}
}

func (p WrittenBallot) Pointer() *WrittenBallot {
	return &p
}

func CastWrittenBallots(data []interface{}) []WrittenBallot {
	v := make([]WrittenBallot, 0)
	for _, d := range data {
		v = append(v, NewWrittenBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWrittenBallotsFromDict(data []WrittenBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Vote struct {
	VoteId         *string         `json:"voteId"`
	RatingName     *string         `json:"ratingName"`
	GatheringName  *string         `json:"gatheringName"`
	WrittenBallots []WrittenBallot `json:"writtenBallots"`
	CreatedAt      *int64          `json:"createdAt"`
	UpdatedAt      *int64          `json:"updatedAt"`
}

func NewVoteFromDict(data map[string]interface{}) Vote {
	return Vote{
		VoteId:         core.CastString(data["voteId"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		WrittenBallots: CastWrittenBallots(core.CastArray(data["writtenBallots"])),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
	}
}

func (p Vote) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"voteId":        p.VoteId,
		"ratingName":    p.RatingName,
		"gatheringName": p.GatheringName,
		"writtenBallots": CastWrittenBallotsFromDict(
			p.WrittenBallots,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p Vote) Pointer() *Vote {
	return &p
}

func CastVotes(data []interface{}) []Vote {
	v := make([]Vote, 0)
	for _, d := range data {
		v = append(v, NewVoteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVotesFromDict(data []Vote) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
