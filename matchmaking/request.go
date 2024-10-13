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

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeNamespacesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeNamespacesRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeNamespacesRequest{}
	} else {
		*p = DescribeNamespacesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeNamespacesRequestFromJson(data string) (DescribeNamespacesRequest, error) {
	req := DescribeNamespacesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeNamespacesRequest{}, err
	}
	return req, nil
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	SourceRequestId                               *string              `json:"sourceRequestId"`
	RequestId                                     *string              `json:"requestId"`
	ContextStack                                  *string              `json:"contextStack"`
	Name                                          *string              `json:"name"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	EnableDisconnectDetection                     *string              `json:"enableDisconnectDetection"`
	DisconnectDetectionTimeoutSeconds             *int32               `json:"disconnectDetectionTimeoutSeconds"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	EnableCollaborateSeasonRating                 *string              `json:"enableCollaborateSeasonRating"`
	CollaborateSeasonRatingNamespaceId            *string              `json:"collaborateSeasonRatingNamespaceId"`
	CollaborateSeasonRatingTtl                    *int32               `json:"collaborateSeasonRatingTtl"`
	ChangeRatingScript                            *ScriptSetting       `json:"changeRatingScript"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification                      *NotificationSetting `json:"changeRatingNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
}

func (p *CreateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateNamespaceRequest{}
	} else {
		*p = CreateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["enableRating"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableRating)
		}
		if v, ok := d["enableDisconnectDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableDisconnectDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableDisconnectDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableDisconnectDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableDisconnectDetection)
				}
			}
		}
		if v, ok := d["disconnectDetectionTimeoutSeconds"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisconnectDetectionTimeoutSeconds)
		}
		if v, ok := d["createGatheringTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerType)
				}
			}
		}
		if v, ok := d["createGatheringTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["createGatheringTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerScriptId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerType)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerScriptId)
				}
			}
		}
		if v, ok := d["enableCollaborateSeasonRating"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableCollaborateSeasonRating = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableCollaborateSeasonRating = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableCollaborateSeasonRating = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableCollaborateSeasonRating)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CollaborateSeasonRatingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingNamespaceId)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingTtl"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingTtl)
		}
		if v, ok := d["changeRatingScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingScript)
		}
		if v, ok := d["joinNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinNotification)
		}
		if v, ok := d["leaveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveNotification)
		}
		if v, ok := d["completeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteNotification)
		}
		if v, ok := d["changeRatingNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingNotification)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
	}
	return nil
}

func NewCreateNamespaceRequestFromJson(data string) (CreateNamespaceRequest, error) {
	req := CreateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateNamespaceRequest{}, err
	}
	return req, nil
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		EnableRating: func() *bool {
			v, ok := data["enableRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableRating"])
		}(),
		EnableDisconnectDetection: func() *string {
			v, ok := data["enableDisconnectDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableDisconnectDetection"])
		}(),
		DisconnectDetectionTimeoutSeconds: func() *int32 {
			v, ok := data["disconnectDetectionTimeoutSeconds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["disconnectDetectionTimeoutSeconds"])
		}(),
		CreateGatheringTriggerType: func() *string {
			v, ok := data["createGatheringTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerType"])
		}(),
		CreateGatheringTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["createGatheringTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerRealtimeNamespaceId"])
		}(),
		CreateGatheringTriggerScriptId: func() *string {
			v, ok := data["createGatheringTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerScriptId"])
		}(),
		CompleteMatchmakingTriggerType: func() *string {
			v, ok := data["completeMatchmakingTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerType"])
		}(),
		CompleteMatchmakingTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["completeMatchmakingTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"])
		}(),
		CompleteMatchmakingTriggerScriptId: func() *string {
			v, ok := data["completeMatchmakingTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerScriptId"])
		}(),
		EnableCollaborateSeasonRating: func() *string {
			v, ok := data["enableCollaborateSeasonRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableCollaborateSeasonRating"])
		}(),
		CollaborateSeasonRatingNamespaceId: func() *string {
			v, ok := data["collaborateSeasonRatingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["collaborateSeasonRatingNamespaceId"])
		}(),
		CollaborateSeasonRatingTtl: func() *int32 {
			v, ok := data["collaborateSeasonRatingTtl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["collaborateSeasonRatingTtl"])
		}(),
		ChangeRatingScript: func() *ScriptSetting {
			v, ok := data["changeRatingScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer()
		}(),
		JoinNotification: func() *NotificationSetting {
			v, ok := data["joinNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer()
		}(),
		LeaveNotification: func() *NotificationSetting {
			v, ok := data["leaveNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer()
		}(),
		CompleteNotification: func() *NotificationSetting {
			v, ok := data["completeNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer()
		}(),
		ChangeRatingNotification: func() *NotificationSetting {
			v, ok := data["changeRatingNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                                          p.Name,
		"description":                                   p.Description,
		"enableRating":                                  p.EnableRating,
		"enableDisconnectDetection":                     p.EnableDisconnectDetection,
		"disconnectDetectionTimeoutSeconds":             p.DisconnectDetectionTimeoutSeconds,
		"createGatheringTriggerType":                    p.CreateGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     p.CreateGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                p.CreateGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                p.CompleteMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            p.CompleteMatchmakingTriggerScriptId,
		"enableCollaborateSeasonRating":                 p.EnableCollaborateSeasonRating,
		"collaborateSeasonRatingNamespaceId":            p.CollaborateSeasonRatingNamespaceId,
		"collaborateSeasonRatingTtl":                    p.CollaborateSeasonRatingTtl,
		"changeRatingScript": func() map[string]interface{} {
			if p.ChangeRatingScript == nil {
				return nil
			}
			return p.ChangeRatingScript.ToDict()
		}(),
		"joinNotification": func() map[string]interface{} {
			if p.JoinNotification == nil {
				return nil
			}
			return p.JoinNotification.ToDict()
		}(),
		"leaveNotification": func() map[string]interface{} {
			if p.LeaveNotification == nil {
				return nil
			}
			return p.LeaveNotification.ToDict()
		}(),
		"completeNotification": func() map[string]interface{} {
			if p.CompleteNotification == nil {
				return nil
			}
			return p.CompleteNotification.ToDict()
		}(),
		"changeRatingNotification": func() map[string]interface{} {
			if p.ChangeRatingNotification == nil {
				return nil
			}
			return p.ChangeRatingNotification.ToDict()
		}(),
		"logSetting": func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceStatusRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetNamespaceStatusRequest{}
	} else {
		*p = GetNamespaceStatusRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceStatusRequestFromJson(data string) (GetNamespaceStatusRequest, error) {
	req := GetNamespaceStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceStatusRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetNamespaceRequest{}
	} else {
		*p = GetNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceRequestFromJson(data string) (GetNamespaceRequest, error) {
	req := GetNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	SourceRequestId                               *string              `json:"sourceRequestId"`
	RequestId                                     *string              `json:"requestId"`
	ContextStack                                  *string              `json:"contextStack"`
	NamespaceName                                 *string              `json:"namespaceName"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	EnableDisconnectDetection                     *string              `json:"enableDisconnectDetection"`
	DisconnectDetectionTimeoutSeconds             *int32               `json:"disconnectDetectionTimeoutSeconds"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	EnableCollaborateSeasonRating                 *string              `json:"enableCollaborateSeasonRating"`
	CollaborateSeasonRatingNamespaceId            *string              `json:"collaborateSeasonRatingNamespaceId"`
	CollaborateSeasonRatingTtl                    *int32               `json:"collaborateSeasonRatingTtl"`
	ChangeRatingScript                            *ScriptSetting       `json:"changeRatingScript"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification                      *NotificationSetting `json:"changeRatingNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
}

func (p *UpdateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateNamespaceRequest{}
	} else {
		*p = UpdateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["enableRating"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableRating)
		}
		if v, ok := d["enableDisconnectDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableDisconnectDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableDisconnectDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableDisconnectDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableDisconnectDetection)
				}
			}
		}
		if v, ok := d["disconnectDetectionTimeoutSeconds"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisconnectDetectionTimeoutSeconds)
		}
		if v, ok := d["createGatheringTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerType)
				}
			}
		}
		if v, ok := d["createGatheringTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["createGatheringTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerScriptId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerType)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerScriptId)
				}
			}
		}
		if v, ok := d["enableCollaborateSeasonRating"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableCollaborateSeasonRating = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableCollaborateSeasonRating = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableCollaborateSeasonRating = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableCollaborateSeasonRating)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CollaborateSeasonRatingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingNamespaceId)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingTtl"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingTtl)
		}
		if v, ok := d["changeRatingScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingScript)
		}
		if v, ok := d["joinNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinNotification)
		}
		if v, ok := d["leaveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveNotification)
		}
		if v, ok := d["completeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteNotification)
		}
		if v, ok := d["changeRatingNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingNotification)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
	}
	return nil
}

func NewUpdateNamespaceRequestFromJson(data string) (UpdateNamespaceRequest, error) {
	req := UpdateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateNamespaceRequest{}, err
	}
	return req, nil
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		EnableRating: func() *bool {
			v, ok := data["enableRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableRating"])
		}(),
		EnableDisconnectDetection: func() *string {
			v, ok := data["enableDisconnectDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableDisconnectDetection"])
		}(),
		DisconnectDetectionTimeoutSeconds: func() *int32 {
			v, ok := data["disconnectDetectionTimeoutSeconds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["disconnectDetectionTimeoutSeconds"])
		}(),
		CreateGatheringTriggerType: func() *string {
			v, ok := data["createGatheringTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerType"])
		}(),
		CreateGatheringTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["createGatheringTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerRealtimeNamespaceId"])
		}(),
		CreateGatheringTriggerScriptId: func() *string {
			v, ok := data["createGatheringTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerScriptId"])
		}(),
		CompleteMatchmakingTriggerType: func() *string {
			v, ok := data["completeMatchmakingTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerType"])
		}(),
		CompleteMatchmakingTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["completeMatchmakingTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"])
		}(),
		CompleteMatchmakingTriggerScriptId: func() *string {
			v, ok := data["completeMatchmakingTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerScriptId"])
		}(),
		EnableCollaborateSeasonRating: func() *string {
			v, ok := data["enableCollaborateSeasonRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableCollaborateSeasonRating"])
		}(),
		CollaborateSeasonRatingNamespaceId: func() *string {
			v, ok := data["collaborateSeasonRatingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["collaborateSeasonRatingNamespaceId"])
		}(),
		CollaborateSeasonRatingTtl: func() *int32 {
			v, ok := data["collaborateSeasonRatingTtl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["collaborateSeasonRatingTtl"])
		}(),
		ChangeRatingScript: func() *ScriptSetting {
			v, ok := data["changeRatingScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer()
		}(),
		JoinNotification: func() *NotificationSetting {
			v, ok := data["joinNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer()
		}(),
		LeaveNotification: func() *NotificationSetting {
			v, ok := data["leaveNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer()
		}(),
		CompleteNotification: func() *NotificationSetting {
			v, ok := data["completeNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer()
		}(),
		ChangeRatingNotification: func() *NotificationSetting {
			v, ok := data["changeRatingNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                                 p.NamespaceName,
		"description":                                   p.Description,
		"enableRating":                                  p.EnableRating,
		"enableDisconnectDetection":                     p.EnableDisconnectDetection,
		"disconnectDetectionTimeoutSeconds":             p.DisconnectDetectionTimeoutSeconds,
		"createGatheringTriggerType":                    p.CreateGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     p.CreateGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                p.CreateGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                p.CompleteMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            p.CompleteMatchmakingTriggerScriptId,
		"enableCollaborateSeasonRating":                 p.EnableCollaborateSeasonRating,
		"collaborateSeasonRatingNamespaceId":            p.CollaborateSeasonRatingNamespaceId,
		"collaborateSeasonRatingTtl":                    p.CollaborateSeasonRatingTtl,
		"changeRatingScript": func() map[string]interface{} {
			if p.ChangeRatingScript == nil {
				return nil
			}
			return p.ChangeRatingScript.ToDict()
		}(),
		"joinNotification": func() map[string]interface{} {
			if p.JoinNotification == nil {
				return nil
			}
			return p.JoinNotification.ToDict()
		}(),
		"leaveNotification": func() map[string]interface{} {
			if p.LeaveNotification == nil {
				return nil
			}
			return p.LeaveNotification.ToDict()
		}(),
		"completeNotification": func() map[string]interface{} {
			if p.CompleteNotification == nil {
				return nil
			}
			return p.CompleteNotification.ToDict()
		}(),
		"changeRatingNotification": func() map[string]interface{} {
			if p.ChangeRatingNotification == nil {
				return nil
			}
			return p.ChangeRatingNotification.ToDict()
		}(),
		"logSetting": func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DeleteNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteNamespaceRequest{}
	} else {
		*p = DeleteNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDeleteNamespaceRequestFromJson(data string) (DeleteNamespaceRequest, error) {
	req := DeleteNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteNamespaceRequest{}, err
	}
	return req, nil
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type DumpUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DumpUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DumpUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DumpUserDataByUserIdRequest{}
	} else {
		*p = DumpUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDumpUserDataByUserIdRequestFromJson(data string) (DumpUserDataByUserIdRequest, error) {
	req := DumpUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DumpUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckDumpUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckDumpUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CheckDumpUserDataByUserIdRequest{}
	} else {
		*p = CheckDumpUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) (CheckDumpUserDataByUserIdRequest, error) {
	req := CheckDumpUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckDumpUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CleanUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CleanUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CleanUserDataByUserIdRequest{}
	} else {
		*p = CleanUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCleanUserDataByUserIdRequestFromJson(data string) (CleanUserDataByUserIdRequest, error) {
	req := CleanUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CleanUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckCleanUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckCleanUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CheckCleanUserDataByUserIdRequest{}
	} else {
		*p = CheckCleanUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) (CheckCleanUserDataByUserIdRequest, error) {
	req := CheckCleanUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckCleanUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *PrepareImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrepareImportUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PrepareImportUserDataByUserIdRequest{}
	} else {
		*p = PrepareImportUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) (PrepareImportUserDataByUserIdRequest, error) {
	req := PrepareImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PrepareImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *ImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ImportUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ImportUserDataByUserIdRequest{}
	} else {
		*p = ImportUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UploadToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UploadToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UploadToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UploadToken)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewImportUserDataByUserIdRequestFromJson(data string) (ImportUserDataByUserIdRequest, error) {
	req := ImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckImportUserDataByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CheckImportUserDataByUserIdRequest{}
	} else {
		*p = CheckImportUserDataByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UploadToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UploadToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UploadToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UploadToken)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) (CheckImportUserDataByUserIdRequest, error) {
	req := CheckImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeGatheringsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeGatheringsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGatheringsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGatheringsRequest{}
	} else {
		*p = DescribeGatheringsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeGatheringsRequestFromJson(data string) (DescribeGatheringsRequest, error) {
	req := DescribeGatheringsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGatheringsRequest{}, err
	}
	return req, nil
}

func NewDescribeGatheringsRequestFromDict(data map[string]interface{}) DescribeGatheringsRequest {
	return DescribeGatheringsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeGatheringsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGatheringsRequest) Pointer() *DescribeGatheringsRequest {
	return &p
}

type CreateGatheringRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	AccessToken        *string          `json:"accessToken"`
	Player             *Player          `json:"player"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles    []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds       []*string        `json:"allowUserIds"`
	ExpiresAt          *int64           `json:"expiresAt"`
	ExpiresAtTimeSpan  *TimeSpan        `json:"expiresAtTimeSpan"`
}

func (p *CreateGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGatheringRequest{}
	} else {
		*p = CreateGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["player"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Player)
		}
		if v, ok := d["attributeRanges"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttributeRanges)
		}
		if v, ok := d["capacityOfRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CapacityOfRoles)
		}
		if v, ok := d["allowUserIds"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.AllowUserIds = l
			}
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["expiresAtTimeSpan"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAtTimeSpan)
		}
	}
	return nil
}

func NewCreateGatheringRequestFromJson(data string) (CreateGatheringRequest, error) {
	req := CreateGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGatheringRequest{}, err
	}
	return req, nil
}

func NewCreateGatheringRequestFromDict(data map[string]interface{}) CreateGatheringRequest {
	return CreateGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Player: func() *Player {
			v, ok := data["player"]
			if !ok || v == nil {
				return nil
			}
			return NewPlayerFromDict(core.CastMap(data["player"])).Pointer()
		}(),
		AttributeRanges: func() []AttributeRange {
			if data["attributeRanges"] == nil {
				return nil
			}
			return CastAttributeRanges(core.CastArray(data["attributeRanges"]))
		}(),
		CapacityOfRoles: func() []CapacityOfRole {
			if data["capacityOfRoles"] == nil {
				return nil
			}
			return CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"]))
		}(),
		AllowUserIds: func() []*string {
			v, ok := data["allowUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		ExpiresAtTimeSpan: func() *TimeSpan {
			v, ok := data["expiresAtTimeSpan"]
			if !ok || v == nil {
				return nil
			}
			return NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer()
		}(),
	}
}

func (p CreateGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"player": func() map[string]interface{} {
			if p.Player == nil {
				return nil
			}
			return p.Player.ToDict()
		}(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt": p.ExpiresAt,
		"expiresAtTimeSpan": func() map[string]interface{} {
			if p.ExpiresAtTimeSpan == nil {
				return nil
			}
			return p.ExpiresAtTimeSpan.ToDict()
		}(),
	}
}

func (p CreateGatheringRequest) Pointer() *CreateGatheringRequest {
	return &p
}

type CreateGatheringByUserIdRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	UserId             *string          `json:"userId"`
	Player             *Player          `json:"player"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles    []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds       []*string        `json:"allowUserIds"`
	ExpiresAt          *int64           `json:"expiresAt"`
	ExpiresAtTimeSpan  *TimeSpan        `json:"expiresAtTimeSpan"`
	TimeOffsetToken    *string          `json:"timeOffsetToken"`
}

func (p *CreateGatheringByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGatheringByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGatheringByUserIdRequest{}
	} else {
		*p = CreateGatheringByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["player"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Player)
		}
		if v, ok := d["attributeRanges"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttributeRanges)
		}
		if v, ok := d["capacityOfRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CapacityOfRoles)
		}
		if v, ok := d["allowUserIds"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.AllowUserIds = l
			}
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["expiresAtTimeSpan"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAtTimeSpan)
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCreateGatheringByUserIdRequestFromJson(data string) (CreateGatheringByUserIdRequest, error) {
	req := CreateGatheringByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGatheringByUserIdRequest{}, err
	}
	return req, nil
}

func NewCreateGatheringByUserIdRequestFromDict(data map[string]interface{}) CreateGatheringByUserIdRequest {
	return CreateGatheringByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Player: func() *Player {
			v, ok := data["player"]
			if !ok || v == nil {
				return nil
			}
			return NewPlayerFromDict(core.CastMap(data["player"])).Pointer()
		}(),
		AttributeRanges: func() []AttributeRange {
			if data["attributeRanges"] == nil {
				return nil
			}
			return CastAttributeRanges(core.CastArray(data["attributeRanges"]))
		}(),
		CapacityOfRoles: func() []CapacityOfRole {
			if data["capacityOfRoles"] == nil {
				return nil
			}
			return CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"]))
		}(),
		AllowUserIds: func() []*string {
			v, ok := data["allowUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		ExpiresAtTimeSpan: func() *TimeSpan {
			v, ok := data["expiresAtTimeSpan"]
			if !ok || v == nil {
				return nil
			}
			return NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer()
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CreateGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"player": func() map[string]interface{} {
			if p.Player == nil {
				return nil
			}
			return p.Player.ToDict()
		}(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt": p.ExpiresAt,
		"expiresAtTimeSpan": func() map[string]interface{} {
			if p.ExpiresAtTimeSpan == nil {
				return nil
			}
			return p.ExpiresAtTimeSpan.ToDict()
		}(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CreateGatheringByUserIdRequest) Pointer() *CreateGatheringByUserIdRequest {
	return &p
}

type UpdateGatheringRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	GatheringName      *string          `json:"gatheringName"`
	AccessToken        *string          `json:"accessToken"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
}

func (p *UpdateGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGatheringRequest{}
	} else {
		*p = UpdateGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["attributeRanges"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttributeRanges)
		}
	}
	return nil
}

func NewUpdateGatheringRequestFromJson(data string) (UpdateGatheringRequest, error) {
	req := UpdateGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGatheringRequest{}, err
	}
	return req, nil
}

func NewUpdateGatheringRequestFromDict(data map[string]interface{}) UpdateGatheringRequest {
	return UpdateGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		AttributeRanges: func() []AttributeRange {
			if data["attributeRanges"] == nil {
				return nil
			}
			return CastAttributeRanges(core.CastArray(data["attributeRanges"]))
		}(),
	}
}

func (p UpdateGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
	}
}

func (p UpdateGatheringRequest) Pointer() *UpdateGatheringRequest {
	return &p
}

type UpdateGatheringByUserIdRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	GatheringName      *string          `json:"gatheringName"`
	UserId             *string          `json:"userId"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	TimeOffsetToken    *string          `json:"timeOffsetToken"`
}

func (p *UpdateGatheringByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGatheringByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGatheringByUserIdRequest{}
	} else {
		*p = UpdateGatheringByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["attributeRanges"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttributeRanges)
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewUpdateGatheringByUserIdRequestFromJson(data string) (UpdateGatheringByUserIdRequest, error) {
	req := UpdateGatheringByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGatheringByUserIdRequest{}, err
	}
	return req, nil
}

func NewUpdateGatheringByUserIdRequestFromDict(data map[string]interface{}) UpdateGatheringByUserIdRequest {
	return UpdateGatheringByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		AttributeRanges: func() []AttributeRange {
			if data["attributeRanges"] == nil {
				return nil
			}
			return CastAttributeRanges(core.CastArray(data["attributeRanges"]))
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p UpdateGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"userId":        p.UserId,
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateGatheringByUserIdRequest) Pointer() *UpdateGatheringByUserIdRequest {
	return &p
}

type DoMatchmakingByPlayerRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	NamespaceName           *string `json:"namespaceName"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func (p *DoMatchmakingByPlayerRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DoMatchmakingByPlayerRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DoMatchmakingByPlayerRequest{}
	} else {
		*p = DoMatchmakingByPlayerRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["player"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Player)
		}
		if v, ok := d["matchmakingContextToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MatchmakingContextToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MatchmakingContextToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MatchmakingContextToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MatchmakingContextToken)
				}
			}
		}
	}
	return nil
}

func NewDoMatchmakingByPlayerRequestFromJson(data string) (DoMatchmakingByPlayerRequest, error) {
	req := DoMatchmakingByPlayerRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DoMatchmakingByPlayerRequest{}, err
	}
	return req, nil
}

func NewDoMatchmakingByPlayerRequestFromDict(data map[string]interface{}) DoMatchmakingByPlayerRequest {
	return DoMatchmakingByPlayerRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Player: func() *Player {
			v, ok := data["player"]
			if !ok || v == nil {
				return nil
			}
			return NewPlayerFromDict(core.CastMap(data["player"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
	}
}

func (p DoMatchmakingByPlayerRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"player": func() map[string]interface{} {
			if p.Player == nil {
				return nil
			}
			return p.Player.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingByPlayerRequest) Pointer() *DoMatchmakingByPlayerRequest {
	return &p
}

type DoMatchmakingRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	AccessToken             *string `json:"accessToken"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func (p *DoMatchmakingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DoMatchmakingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DoMatchmakingRequest{}
	} else {
		*p = DoMatchmakingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["player"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Player)
		}
		if v, ok := d["matchmakingContextToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MatchmakingContextToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MatchmakingContextToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MatchmakingContextToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MatchmakingContextToken)
				}
			}
		}
	}
	return nil
}

func NewDoMatchmakingRequestFromJson(data string) (DoMatchmakingRequest, error) {
	req := DoMatchmakingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DoMatchmakingRequest{}, err
	}
	return req, nil
}

func NewDoMatchmakingRequestFromDict(data map[string]interface{}) DoMatchmakingRequest {
	return DoMatchmakingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Player: func() *Player {
			v, ok := data["player"]
			if !ok || v == nil {
				return nil
			}
			return NewPlayerFromDict(core.CastMap(data["player"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
	}
}

func (p DoMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"player": func() map[string]interface{} {
			if p.Player == nil {
				return nil
			}
			return p.Player.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingRequest) Pointer() *DoMatchmakingRequest {
	return &p
}

type DoMatchmakingByUserIdRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	UserId                  *string `json:"userId"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
	TimeOffsetToken         *string `json:"timeOffsetToken"`
}

func (p *DoMatchmakingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DoMatchmakingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DoMatchmakingByUserIdRequest{}
	} else {
		*p = DoMatchmakingByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["player"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Player)
		}
		if v, ok := d["matchmakingContextToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MatchmakingContextToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MatchmakingContextToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MatchmakingContextToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MatchmakingContextToken)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDoMatchmakingByUserIdRequestFromJson(data string) (DoMatchmakingByUserIdRequest, error) {
	req := DoMatchmakingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DoMatchmakingByUserIdRequest{}, err
	}
	return req, nil
}

func NewDoMatchmakingByUserIdRequestFromDict(data map[string]interface{}) DoMatchmakingByUserIdRequest {
	return DoMatchmakingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Player: func() *Player {
			v, ok := data["player"]
			if !ok || v == nil {
				return nil
			}
			return NewPlayerFromDict(core.CastMap(data["player"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DoMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"player": func() map[string]interface{} {
			if p.Player == nil {
				return nil
			}
			return p.Player.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"timeOffsetToken":         p.TimeOffsetToken,
	}
}

func (p DoMatchmakingByUserIdRequest) Pointer() *DoMatchmakingByUserIdRequest {
	return &p
}

type PingRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	AccessToken        *string `json:"accessToken"`
}

func (p *PingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PingRequest{}
	} else {
		*p = PingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
	}
	return nil
}

func NewPingRequestFromJson(data string) (PingRequest, error) {
	req := PingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PingRequest{}, err
	}
	return req, nil
}

func NewPingRequestFromDict(data map[string]interface{}) PingRequest {
	return PingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p PingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
	}
}

func (p PingRequest) Pointer() *PingRequest {
	return &p
}

type PingByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *PingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PingByUserIdRequest{}
	} else {
		*p = PingByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewPingByUserIdRequestFromJson(data string) (PingByUserIdRequest, error) {
	req := PingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PingByUserIdRequest{}, err
	}
	return req, nil
}

func NewPingByUserIdRequestFromDict(data map[string]interface{}) PingByUserIdRequest {
	return PingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p PingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PingByUserIdRequest) Pointer() *PingByUserIdRequest {
	return &p
}

type GetGatheringRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GatheringName   *string `json:"gatheringName"`
}

func (p *GetGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGatheringRequest{}
	} else {
		*p = GetGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
	}
	return nil
}

func NewGetGatheringRequestFromJson(data string) (GetGatheringRequest, error) {
	req := GetGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGatheringRequest{}, err
	}
	return req, nil
}

func NewGetGatheringRequestFromDict(data map[string]interface{}) GetGatheringRequest {
	return GetGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
	}
}

func (p GetGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
	}
}

func (p GetGatheringRequest) Pointer() *GetGatheringRequest {
	return &p
}

type CancelMatchmakingRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	AccessToken        *string `json:"accessToken"`
}

func (p *CancelMatchmakingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CancelMatchmakingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CancelMatchmakingRequest{}
	} else {
		*p = CancelMatchmakingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
	}
	return nil
}

func NewCancelMatchmakingRequestFromJson(data string) (CancelMatchmakingRequest, error) {
	req := CancelMatchmakingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CancelMatchmakingRequest{}, err
	}
	return req, nil
}

func NewCancelMatchmakingRequestFromDict(data map[string]interface{}) CancelMatchmakingRequest {
	return CancelMatchmakingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p CancelMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
	}
}

func (p CancelMatchmakingRequest) Pointer() *CancelMatchmakingRequest {
	return &p
}

type CancelMatchmakingByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *CancelMatchmakingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CancelMatchmakingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CancelMatchmakingByUserIdRequest{}
	} else {
		*p = CancelMatchmakingByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCancelMatchmakingByUserIdRequestFromJson(data string) (CancelMatchmakingByUserIdRequest, error) {
	req := CancelMatchmakingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CancelMatchmakingByUserIdRequest{}, err
	}
	return req, nil
}

func NewCancelMatchmakingByUserIdRequestFromDict(data map[string]interface{}) CancelMatchmakingByUserIdRequest {
	return CancelMatchmakingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CancelMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CancelMatchmakingByUserIdRequest) Pointer() *CancelMatchmakingByUserIdRequest {
	return &p
}

type EarlyCompleteRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	AccessToken        *string `json:"accessToken"`
}

func (p *EarlyCompleteRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EarlyCompleteRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = EarlyCompleteRequest{}
	} else {
		*p = EarlyCompleteRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
	}
	return nil
}

func NewEarlyCompleteRequestFromJson(data string) (EarlyCompleteRequest, error) {
	req := EarlyCompleteRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return EarlyCompleteRequest{}, err
	}
	return req, nil
}

func NewEarlyCompleteRequestFromDict(data map[string]interface{}) EarlyCompleteRequest {
	return EarlyCompleteRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p EarlyCompleteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
	}
}

func (p EarlyCompleteRequest) Pointer() *EarlyCompleteRequest {
	return &p
}

type EarlyCompleteByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *EarlyCompleteByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EarlyCompleteByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = EarlyCompleteByUserIdRequest{}
	} else {
		*p = EarlyCompleteByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewEarlyCompleteByUserIdRequestFromJson(data string) (EarlyCompleteByUserIdRequest, error) {
	req := EarlyCompleteByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return EarlyCompleteByUserIdRequest{}, err
	}
	return req, nil
}

func NewEarlyCompleteByUserIdRequestFromDict(data map[string]interface{}) EarlyCompleteByUserIdRequest {
	return EarlyCompleteByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p EarlyCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p EarlyCompleteByUserIdRequest) Pointer() *EarlyCompleteByUserIdRequest {
	return &p
}

type DeleteGatheringRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GatheringName   *string `json:"gatheringName"`
}

func (p *DeleteGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGatheringRequest{}
	} else {
		*p = DeleteGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
	}
	return nil
}

func NewDeleteGatheringRequestFromJson(data string) (DeleteGatheringRequest, error) {
	req := DeleteGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGatheringRequest{}, err
	}
	return req, nil
}

func NewDeleteGatheringRequestFromDict(data map[string]interface{}) DeleteGatheringRequest {
	return DeleteGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
	}
}

func (p DeleteGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
	}
}

func (p DeleteGatheringRequest) Pointer() *DeleteGatheringRequest {
	return &p
}

type DescribeRatingModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeRatingModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRatingModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRatingModelMastersRequest{}
	} else {
		*p = DescribeRatingModelMastersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeRatingModelMastersRequestFromJson(data string) (DescribeRatingModelMastersRequest, error) {
	req := DescribeRatingModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRatingModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeRatingModelMastersRequestFromDict(data map[string]interface{}) DescribeRatingModelMastersRequest {
	return DescribeRatingModelMastersRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeRatingModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRatingModelMastersRequest) Pointer() *DescribeRatingModelMastersRequest {
	return &p
}

type CreateRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	InitialValue    *int32  `json:"initialValue"`
	Volatility      *int32  `json:"volatility"`
}

func (p *CreateRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateRatingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateRatingModelMasterRequest{}
	} else {
		*p = CreateRatingModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["initialValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialValue)
		}
		if v, ok := d["volatility"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Volatility)
		}
	}
	return nil
}

func NewCreateRatingModelMasterRequestFromJson(data string) (CreateRatingModelMasterRequest, error) {
	req := CreateRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateRatingModelMasterRequestFromDict(data map[string]interface{}) CreateRatingModelMasterRequest {
	return CreateRatingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		InitialValue: func() *int32 {
			v, ok := data["initialValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialValue"])
		}(),
		Volatility: func() *int32 {
			v, ok := data["volatility"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["volatility"])
		}(),
	}
}

func (p CreateRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"initialValue":  p.InitialValue,
		"volatility":    p.Volatility,
	}
}

func (p CreateRatingModelMasterRequest) Pointer() *CreateRatingModelMasterRequest {
	return &p
}

type GetRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func (p *GetRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRatingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRatingModelMasterRequest{}
	} else {
		*p = GetRatingModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
	}
	return nil
}

func NewGetRatingModelMasterRequestFromJson(data string) (GetRatingModelMasterRequest, error) {
	req := GetRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetRatingModelMasterRequestFromDict(data map[string]interface{}) GetRatingModelMasterRequest {
	return GetRatingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
	}
}

func (p GetRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingModelMasterRequest) Pointer() *GetRatingModelMasterRequest {
	return &p
}

type UpdateRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	InitialValue    *int32  `json:"initialValue"`
	Volatility      *int32  `json:"volatility"`
}

func (p *UpdateRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateRatingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateRatingModelMasterRequest{}
	} else {
		*p = UpdateRatingModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["initialValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialValue)
		}
		if v, ok := d["volatility"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Volatility)
		}
	}
	return nil
}

func NewUpdateRatingModelMasterRequestFromJson(data string) (UpdateRatingModelMasterRequest, error) {
	req := UpdateRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateRatingModelMasterRequest {
	return UpdateRatingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		InitialValue: func() *int32 {
			v, ok := data["initialValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialValue"])
		}(),
		Volatility: func() *int32 {
			v, ok := data["volatility"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["volatility"])
		}(),
	}
}

func (p UpdateRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"initialValue":  p.InitialValue,
		"volatility":    p.Volatility,
	}
}

func (p UpdateRatingModelMasterRequest) Pointer() *UpdateRatingModelMasterRequest {
	return &p
}

type DeleteRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func (p *DeleteRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRatingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRatingModelMasterRequest{}
	} else {
		*p = DeleteRatingModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
	}
	return nil
}

func NewDeleteRatingModelMasterRequestFromJson(data string) (DeleteRatingModelMasterRequest, error) {
	req := DeleteRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteRatingModelMasterRequestFromDict(data map[string]interface{}) DeleteRatingModelMasterRequest {
	return DeleteRatingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
	}
}

func (p DeleteRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p DeleteRatingModelMasterRequest) Pointer() *DeleteRatingModelMasterRequest {
	return &p
}

type DescribeRatingModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeRatingModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRatingModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRatingModelsRequest{}
	} else {
		*p = DescribeRatingModelsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeRatingModelsRequestFromJson(data string) (DescribeRatingModelsRequest, error) {
	req := DescribeRatingModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRatingModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeRatingModelsRequestFromDict(data map[string]interface{}) DescribeRatingModelsRequest {
	return DescribeRatingModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeRatingModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRatingModelsRequest) Pointer() *DescribeRatingModelsRequest {
	return &p
}

type GetRatingModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func (p *GetRatingModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRatingModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRatingModelRequest{}
	} else {
		*p = GetRatingModelRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
	}
	return nil
}

func NewGetRatingModelRequestFromJson(data string) (GetRatingModelRequest, error) {
	req := GetRatingModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRatingModelRequest{}, err
	}
	return req, nil
}

func NewGetRatingModelRequestFromDict(data map[string]interface{}) GetRatingModelRequest {
	return GetRatingModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
	}
}

func (p GetRatingModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingModelRequest) Pointer() *GetRatingModelRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *ExportMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExportMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ExportMasterRequest{}
	} else {
		*p = ExportMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewExportMasterRequestFromJson(data string) (ExportMasterRequest, error) {
	req := ExportMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ExportMasterRequest{}, err
	}
	return req, nil
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
	return ExportMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
	return &p
}

type GetCurrentModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentModelMasterRequest{}
	} else {
		*p = GetCurrentModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetCurrentModelMasterRequestFromJson(data string) (GetCurrentModelMasterRequest, error) {
	req := GetCurrentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentModelMasterRequestFromDict(data map[string]interface{}) GetCurrentModelMasterRequest {
	return GetCurrentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentModelMasterRequest) Pointer() *GetCurrentModelMasterRequest {
	return &p
}

type UpdateCurrentModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentModelMasterRequest{}
	} else {
		*p = UpdateCurrentModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Settings = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Settings = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Settings = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Settings)
				}
			}
		}
	}
	return nil
}

func NewUpdateCurrentModelMasterRequestFromJson(data string) (UpdateCurrentModelMasterRequest, error) {
	req := UpdateCurrentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentModelMasterRequest {
	return UpdateCurrentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
	}
}

func (p UpdateCurrentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentModelMasterRequest) Pointer() *UpdateCurrentModelMasterRequest {
	return &p
}

type UpdateCurrentModelMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentModelMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentModelMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentModelMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentModelMasterFromGitHubRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["checkoutSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CheckoutSetting)
		}
	}
	return nil
}

func NewUpdateCurrentModelMasterFromGitHubRequestFromJson(data string) (UpdateCurrentModelMasterFromGitHubRequest, error) {
	req := UpdateCurrentModelMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentModelMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentModelMasterFromGitHubRequest {
	return UpdateCurrentModelMasterFromGitHubRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		CheckoutSetting: func() *GitHubCheckoutSetting {
			v, ok := data["checkoutSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"checkoutSetting": func() map[string]interface{} {
			if p.CheckoutSetting == nil {
				return nil
			}
			return p.CheckoutSetting.ToDict()
		}(),
	}
}

func (p UpdateCurrentModelMasterFromGitHubRequest) Pointer() *UpdateCurrentModelMasterFromGitHubRequest {
	return &p
}

type DescribeSeasonModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeSeasonModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSeasonModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSeasonModelsRequest{}
	} else {
		*p = DescribeSeasonModelsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeSeasonModelsRequestFromJson(data string) (DescribeSeasonModelsRequest, error) {
	req := DescribeSeasonModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSeasonModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeSeasonModelsRequestFromDict(data map[string]interface{}) DescribeSeasonModelsRequest {
	return DescribeSeasonModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeSeasonModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeSeasonModelsRequest) Pointer() *DescribeSeasonModelsRequest {
	return &p
}

type GetSeasonModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func (p *GetSeasonModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSeasonModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSeasonModelRequest{}
	} else {
		*p = GetSeasonModelRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
	}
	return nil
}

func NewGetSeasonModelRequestFromJson(data string) (GetSeasonModelRequest, error) {
	req := GetSeasonModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSeasonModelRequest{}, err
	}
	return req, nil
}

func NewGetSeasonModelRequestFromDict(data map[string]interface{}) GetSeasonModelRequest {
	return GetSeasonModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
	}
}

func (p GetSeasonModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p GetSeasonModelRequest) Pointer() *GetSeasonModelRequest {
	return &p
}

type DescribeSeasonModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSeasonModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSeasonModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSeasonModelMastersRequest{}
	} else {
		*p = DescribeSeasonModelMastersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeSeasonModelMastersRequestFromJson(data string) (DescribeSeasonModelMastersRequest, error) {
	req := DescribeSeasonModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSeasonModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeSeasonModelMastersRequestFromDict(data map[string]interface{}) DescribeSeasonModelMastersRequest {
	return DescribeSeasonModelMastersRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeSeasonModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSeasonModelMastersRequest) Pointer() *DescribeSeasonModelMastersRequest {
	return &p
}

type CreateSeasonModelMasterRequest struct {
	SourceRequestId        *string `json:"sourceRequestId"`
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	NamespaceName          *string `json:"namespaceName"`
	Name                   *string `json:"name"`
	Description            *string `json:"description"`
	Metadata               *string `json:"metadata"`
	MaximumParticipants    *int32  `json:"maximumParticipants"`
	ExperienceModelId      *string `json:"experienceModelId"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func (p *CreateSeasonModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateSeasonModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateSeasonModelMasterRequest{}
	} else {
		*p = CreateSeasonModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["maximumParticipants"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParticipants)
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceModelId)
				}
			}
		}
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewCreateSeasonModelMasterRequestFromJson(data string) (CreateSeasonModelMasterRequest, error) {
	req := CreateSeasonModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateSeasonModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateSeasonModelMasterRequestFromDict(data map[string]interface{}) CreateSeasonModelMasterRequest {
	return CreateSeasonModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		MaximumParticipants: func() *int32 {
			v, ok := data["maximumParticipants"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParticipants"])
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		ChallengePeriodEventId: func() *string {
			v, ok := data["challengePeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["challengePeriodEventId"])
		}(),
	}
}

func (p CreateSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"name":                   p.Name,
		"description":            p.Description,
		"metadata":               p.Metadata,
		"maximumParticipants":    p.MaximumParticipants,
		"experienceModelId":      p.ExperienceModelId,
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p CreateSeasonModelMasterRequest) Pointer() *CreateSeasonModelMasterRequest {
	return &p
}

type GetSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func (p *GetSeasonModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSeasonModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSeasonModelMasterRequest{}
	} else {
		*p = GetSeasonModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
	}
	return nil
}

func NewGetSeasonModelMasterRequestFromJson(data string) (GetSeasonModelMasterRequest, error) {
	req := GetSeasonModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSeasonModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetSeasonModelMasterRequestFromDict(data map[string]interface{}) GetSeasonModelMasterRequest {
	return GetSeasonModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
	}
}

func (p GetSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p GetSeasonModelMasterRequest) Pointer() *GetSeasonModelMasterRequest {
	return &p
}

type UpdateSeasonModelMasterRequest struct {
	SourceRequestId        *string `json:"sourceRequestId"`
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	NamespaceName          *string `json:"namespaceName"`
	SeasonName             *string `json:"seasonName"`
	Description            *string `json:"description"`
	Metadata               *string `json:"metadata"`
	MaximumParticipants    *int32  `json:"maximumParticipants"`
	ExperienceModelId      *string `json:"experienceModelId"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func (p *UpdateSeasonModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateSeasonModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateSeasonModelMasterRequest{}
	} else {
		*p = UpdateSeasonModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["maximumParticipants"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParticipants)
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceModelId)
				}
			}
		}
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewUpdateSeasonModelMasterRequestFromJson(data string) (UpdateSeasonModelMasterRequest, error) {
	req := UpdateSeasonModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateSeasonModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateSeasonModelMasterRequestFromDict(data map[string]interface{}) UpdateSeasonModelMasterRequest {
	return UpdateSeasonModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		MaximumParticipants: func() *int32 {
			v, ok := data["maximumParticipants"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParticipants"])
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		ChallengePeriodEventId: func() *string {
			v, ok := data["challengePeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["challengePeriodEventId"])
		}(),
	}
}

func (p UpdateSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"seasonName":             p.SeasonName,
		"description":            p.Description,
		"metadata":               p.Metadata,
		"maximumParticipants":    p.MaximumParticipants,
		"experienceModelId":      p.ExperienceModelId,
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p UpdateSeasonModelMasterRequest) Pointer() *UpdateSeasonModelMasterRequest {
	return &p
}

type DeleteSeasonModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
}

func (p *DeleteSeasonModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSeasonModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSeasonModelMasterRequest{}
	} else {
		*p = DeleteSeasonModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSeasonModelMasterRequestFromJson(data string) (DeleteSeasonModelMasterRequest, error) {
	req := DeleteSeasonModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSeasonModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteSeasonModelMasterRequestFromDict(data map[string]interface{}) DeleteSeasonModelMasterRequest {
	return DeleteSeasonModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
	}
}

func (p DeleteSeasonModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
	}
}

func (p DeleteSeasonModelMasterRequest) Pointer() *DeleteSeasonModelMasterRequest {
	return &p
}

type DescribeSeasonGatheringsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
	Season          *int64  `json:"season"`
	Tier            *int64  `json:"tier"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSeasonGatheringsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSeasonGatheringsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSeasonGatheringsRequest{}
	} else {
		*p = DescribeSeasonGatheringsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeSeasonGatheringsRequestFromJson(data string) (DescribeSeasonGatheringsRequest, error) {
	req := DescribeSeasonGatheringsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSeasonGatheringsRequest{}, err
	}
	return req, nil
}

func NewDescribeSeasonGatheringsRequestFromDict(data map[string]interface{}) DescribeSeasonGatheringsRequest {
	return DescribeSeasonGatheringsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeSeasonGatheringsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
		"season":        p.Season,
		"tier":          p.Tier,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSeasonGatheringsRequest) Pointer() *DescribeSeasonGatheringsRequest {
	return &p
}

type DescribeMatchmakingSeasonGatheringsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	SeasonName      *string `json:"seasonName"`
	Season          *int64  `json:"season"`
	Tier            *int64  `json:"tier"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeMatchmakingSeasonGatheringsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMatchmakingSeasonGatheringsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMatchmakingSeasonGatheringsRequest{}
	} else {
		*p = DescribeMatchmakingSeasonGatheringsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeMatchmakingSeasonGatheringsRequestFromJson(data string) (DescribeMatchmakingSeasonGatheringsRequest, error) {
	req := DescribeMatchmakingSeasonGatheringsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMatchmakingSeasonGatheringsRequest{}, err
	}
	return req, nil
}

func NewDescribeMatchmakingSeasonGatheringsRequestFromDict(data map[string]interface{}) DescribeMatchmakingSeasonGatheringsRequest {
	return DescribeMatchmakingSeasonGatheringsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeMatchmakingSeasonGatheringsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"seasonName":    p.SeasonName,
		"season":        p.Season,
		"tier":          p.Tier,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMatchmakingSeasonGatheringsRequest) Pointer() *DescribeMatchmakingSeasonGatheringsRequest {
	return &p
}

type DoSeasonMatchmakingRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	SeasonName              *string `json:"seasonName"`
	AccessToken             *string `json:"accessToken"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func (p *DoSeasonMatchmakingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DoSeasonMatchmakingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DoSeasonMatchmakingRequest{}
	} else {
		*p = DoSeasonMatchmakingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["matchmakingContextToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MatchmakingContextToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MatchmakingContextToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MatchmakingContextToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MatchmakingContextToken)
				}
			}
		}
	}
	return nil
}

func NewDoSeasonMatchmakingRequestFromJson(data string) (DoSeasonMatchmakingRequest, error) {
	req := DoSeasonMatchmakingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DoSeasonMatchmakingRequest{}, err
	}
	return req, nil
}

func NewDoSeasonMatchmakingRequestFromDict(data map[string]interface{}) DoSeasonMatchmakingRequest {
	return DoSeasonMatchmakingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
	}
}

func (p DoSeasonMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"seasonName":              p.SeasonName,
		"accessToken":             p.AccessToken,
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoSeasonMatchmakingRequest) Pointer() *DoSeasonMatchmakingRequest {
	return &p
}

type DoSeasonMatchmakingByUserIdRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	SeasonName              *string `json:"seasonName"`
	UserId                  *string `json:"userId"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
	TimeOffsetToken         *string `json:"timeOffsetToken"`
}

func (p *DoSeasonMatchmakingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DoSeasonMatchmakingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DoSeasonMatchmakingByUserIdRequest{}
	} else {
		*p = DoSeasonMatchmakingByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["matchmakingContextToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MatchmakingContextToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MatchmakingContextToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MatchmakingContextToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MatchmakingContextToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MatchmakingContextToken)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDoSeasonMatchmakingByUserIdRequestFromJson(data string) (DoSeasonMatchmakingByUserIdRequest, error) {
	req := DoSeasonMatchmakingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DoSeasonMatchmakingByUserIdRequest{}, err
	}
	return req, nil
}

func NewDoSeasonMatchmakingByUserIdRequestFromDict(data map[string]interface{}) DoSeasonMatchmakingByUserIdRequest {
	return DoSeasonMatchmakingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DoSeasonMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"seasonName":              p.SeasonName,
		"userId":                  p.UserId,
		"matchmakingContextToken": p.MatchmakingContextToken,
		"timeOffsetToken":         p.TimeOffsetToken,
	}
}

func (p DoSeasonMatchmakingByUserIdRequest) Pointer() *DoSeasonMatchmakingByUserIdRequest {
	return &p
}

type GetSeasonGatheringRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	SeasonName          *string `json:"seasonName"`
	Season              *int64  `json:"season"`
	Tier                *int64  `json:"tier"`
	SeasonGatheringName *string `json:"seasonGatheringName"`
}

func (p *GetSeasonGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSeasonGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSeasonGatheringRequest{}
	} else {
		*p = GetSeasonGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["seasonGatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringName)
				}
			}
		}
	}
	return nil
}

func NewGetSeasonGatheringRequestFromJson(data string) (GetSeasonGatheringRequest, error) {
	req := GetSeasonGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSeasonGatheringRequest{}, err
	}
	return req, nil
}

func NewGetSeasonGatheringRequestFromDict(data map[string]interface{}) GetSeasonGatheringRequest {
	return GetSeasonGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		SeasonGatheringName: func() *string {
			v, ok := data["seasonGatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringName"])
		}(),
	}
}

func (p GetSeasonGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"seasonName":          p.SeasonName,
		"season":              p.Season,
		"tier":                p.Tier,
		"seasonGatheringName": p.SeasonGatheringName,
	}
}

func (p GetSeasonGatheringRequest) Pointer() *GetSeasonGatheringRequest {
	return &p
}

type VerifyIncludeParticipantRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	SeasonName          *string `json:"seasonName"`
	Season              *int64  `json:"season"`
	Tier                *int64  `json:"tier"`
	SeasonGatheringName *string `json:"seasonGatheringName"`
	AccessToken         *string `json:"accessToken"`
	VerifyType          *string `json:"verifyType"`
}

func (p *VerifyIncludeParticipantRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeParticipantRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeParticipantRequest{}
	} else {
		*p = VerifyIncludeParticipantRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["seasonGatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
	}
	return nil
}

func NewVerifyIncludeParticipantRequestFromJson(data string) (VerifyIncludeParticipantRequest, error) {
	req := VerifyIncludeParticipantRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeParticipantRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeParticipantRequestFromDict(data map[string]interface{}) VerifyIncludeParticipantRequest {
	return VerifyIncludeParticipantRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		SeasonGatheringName: func() *string {
			v, ok := data["seasonGatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
	}
}

func (p VerifyIncludeParticipantRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"seasonName":          p.SeasonName,
		"season":              p.Season,
		"tier":                p.Tier,
		"seasonGatheringName": p.SeasonGatheringName,
		"accessToken":         p.AccessToken,
		"verifyType":          p.VerifyType,
	}
}

func (p VerifyIncludeParticipantRequest) Pointer() *VerifyIncludeParticipantRequest {
	return &p
}

type VerifyIncludeParticipantByUserIdRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	DuplicationAvoider  *string `json:"duplicationAvoider"`
	NamespaceName       *string `json:"namespaceName"`
	SeasonName          *string `json:"seasonName"`
	Season              *int64  `json:"season"`
	Tier                *int64  `json:"tier"`
	SeasonGatheringName *string `json:"seasonGatheringName"`
	UserId              *string `json:"userId"`
	VerifyType          *string `json:"verifyType"`
	TimeOffsetToken     *string `json:"timeOffsetToken"`
}

func (p *VerifyIncludeParticipantByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeParticipantByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeParticipantByUserIdRequest{}
	} else {
		*p = VerifyIncludeParticipantByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["seasonGatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewVerifyIncludeParticipantByUserIdRequestFromJson(data string) (VerifyIncludeParticipantByUserIdRequest, error) {
	req := VerifyIncludeParticipantByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeParticipantByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeParticipantByUserIdRequestFromDict(data map[string]interface{}) VerifyIncludeParticipantByUserIdRequest {
	return VerifyIncludeParticipantByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		SeasonGatheringName: func() *string {
			v, ok := data["seasonGatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p VerifyIncludeParticipantByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"seasonName":          p.SeasonName,
		"season":              p.Season,
		"tier":                p.Tier,
		"seasonGatheringName": p.SeasonGatheringName,
		"userId":              p.UserId,
		"verifyType":          p.VerifyType,
		"timeOffsetToken":     p.TimeOffsetToken,
	}
}

func (p VerifyIncludeParticipantByUserIdRequest) Pointer() *VerifyIncludeParticipantByUserIdRequest {
	return &p
}

type DeleteSeasonGatheringRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	SeasonName          *string `json:"seasonName"`
	Season              *int64  `json:"season"`
	Tier                *int64  `json:"tier"`
	SeasonGatheringName *string `json:"seasonGatheringName"`
}

func (p *DeleteSeasonGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSeasonGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSeasonGatheringRequest{}
	} else {
		*p = DeleteSeasonGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["seasonGatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSeasonGatheringRequestFromJson(data string) (DeleteSeasonGatheringRequest, error) {
	req := DeleteSeasonGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSeasonGatheringRequest{}, err
	}
	return req, nil
}

func NewDeleteSeasonGatheringRequestFromDict(data map[string]interface{}) DeleteSeasonGatheringRequest {
	return DeleteSeasonGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		SeasonGatheringName: func() *string {
			v, ok := data["seasonGatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringName"])
		}(),
	}
}

func (p DeleteSeasonGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"seasonName":          p.SeasonName,
		"season":              p.Season,
		"tier":                p.Tier,
		"seasonGatheringName": p.SeasonGatheringName,
	}
}

func (p DeleteSeasonGatheringRequest) Pointer() *DeleteSeasonGatheringRequest {
	return &p
}

type VerifyIncludeParticipantByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyIncludeParticipantByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeParticipantByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeParticipantByStampTaskRequest{}
	} else {
		*p = VerifyIncludeParticipantByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyIncludeParticipantByStampTaskRequestFromJson(data string) (VerifyIncludeParticipantByStampTaskRequest, error) {
	req := VerifyIncludeParticipantByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeParticipantByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeParticipantByStampTaskRequestFromDict(data map[string]interface{}) VerifyIncludeParticipantByStampTaskRequest {
	return VerifyIncludeParticipantByStampTaskRequest{
		StampTask: func() *string {
			v, ok := data["stampTask"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampTask"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p VerifyIncludeParticipantByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyIncludeParticipantByStampTaskRequest) Pointer() *VerifyIncludeParticipantByStampTaskRequest {
	return &p
}

type DescribeJoinedSeasonGatheringsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	SeasonName      *string `json:"seasonName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeJoinedSeasonGatheringsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeJoinedSeasonGatheringsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeJoinedSeasonGatheringsRequest{}
	} else {
		*p = DescribeJoinedSeasonGatheringsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeJoinedSeasonGatheringsRequestFromJson(data string) (DescribeJoinedSeasonGatheringsRequest, error) {
	req := DescribeJoinedSeasonGatheringsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeJoinedSeasonGatheringsRequest{}, err
	}
	return req, nil
}

func NewDescribeJoinedSeasonGatheringsRequestFromDict(data map[string]interface{}) DescribeJoinedSeasonGatheringsRequest {
	return DescribeJoinedSeasonGatheringsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeJoinedSeasonGatheringsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"seasonName":    p.SeasonName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeJoinedSeasonGatheringsRequest) Pointer() *DescribeJoinedSeasonGatheringsRequest {
	return &p
}

type DescribeJoinedSeasonGatheringsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	SeasonName      *string `json:"seasonName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeJoinedSeasonGatheringsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeJoinedSeasonGatheringsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeJoinedSeasonGatheringsByUserIdRequest{}
	} else {
		*p = DescribeJoinedSeasonGatheringsByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDescribeJoinedSeasonGatheringsByUserIdRequestFromJson(data string) (DescribeJoinedSeasonGatheringsByUserIdRequest, error) {
	req := DescribeJoinedSeasonGatheringsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeJoinedSeasonGatheringsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeJoinedSeasonGatheringsByUserIdRequestFromDict(data map[string]interface{}) DescribeJoinedSeasonGatheringsByUserIdRequest {
	return DescribeJoinedSeasonGatheringsByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DescribeJoinedSeasonGatheringsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"seasonName":      p.SeasonName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeJoinedSeasonGatheringsByUserIdRequest) Pointer() *DescribeJoinedSeasonGatheringsByUserIdRequest {
	return &p
}

type GetJoinedSeasonGatheringRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	SeasonName      *string `json:"seasonName"`
	Season          *int64  `json:"season"`
}

func (p *GetJoinedSeasonGatheringRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetJoinedSeasonGatheringRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetJoinedSeasonGatheringRequest{}
	} else {
		*p = GetJoinedSeasonGatheringRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetJoinedSeasonGatheringRequestFromJson(data string) (GetJoinedSeasonGatheringRequest, error) {
	req := GetJoinedSeasonGatheringRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetJoinedSeasonGatheringRequest{}, err
	}
	return req, nil
}

func NewGetJoinedSeasonGatheringRequestFromDict(data map[string]interface{}) GetJoinedSeasonGatheringRequest {
	return GetJoinedSeasonGatheringRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetJoinedSeasonGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"seasonName":    p.SeasonName,
		"season":        p.Season,
	}
}

func (p GetJoinedSeasonGatheringRequest) Pointer() *GetJoinedSeasonGatheringRequest {
	return &p
}

type GetJoinedSeasonGatheringByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	SeasonName      *string `json:"seasonName"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetJoinedSeasonGatheringByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetJoinedSeasonGatheringByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetJoinedSeasonGatheringByUserIdRequest{}
	} else {
		*p = GetJoinedSeasonGatheringByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewGetJoinedSeasonGatheringByUserIdRequestFromJson(data string) (GetJoinedSeasonGatheringByUserIdRequest, error) {
	req := GetJoinedSeasonGatheringByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetJoinedSeasonGatheringByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetJoinedSeasonGatheringByUserIdRequestFromDict(data map[string]interface{}) GetJoinedSeasonGatheringByUserIdRequest {
	return GetJoinedSeasonGatheringByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p GetJoinedSeasonGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"seasonName":      p.SeasonName,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetJoinedSeasonGatheringByUserIdRequest) Pointer() *GetJoinedSeasonGatheringByUserIdRequest {
	return &p
}

type DescribeRatingsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeRatingsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRatingsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRatingsRequest{}
	} else {
		*p = DescribeRatingsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeRatingsRequestFromJson(data string) (DescribeRatingsRequest, error) {
	req := DescribeRatingsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRatingsRequest{}, err
	}
	return req, nil
}

func NewDescribeRatingsRequestFromDict(data map[string]interface{}) DescribeRatingsRequest {
	return DescribeRatingsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeRatingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRatingsRequest) Pointer() *DescribeRatingsRequest {
	return &p
}

type DescribeRatingsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeRatingsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRatingsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRatingsByUserIdRequest{}
	} else {
		*p = DescribeRatingsByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDescribeRatingsByUserIdRequestFromJson(data string) (DescribeRatingsByUserIdRequest, error) {
	req := DescribeRatingsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRatingsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeRatingsByUserIdRequestFromDict(data map[string]interface{}) DescribeRatingsByUserIdRequest {
	return DescribeRatingsByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DescribeRatingsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeRatingsByUserIdRequest) Pointer() *DescribeRatingsByUserIdRequest {
	return &p
}

type GetRatingRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	RatingName      *string `json:"ratingName"`
}

func (p *GetRatingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRatingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRatingRequest{}
	} else {
		*p = GetRatingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
	}
	return nil
}

func NewGetRatingRequestFromJson(data string) (GetRatingRequest, error) {
	req := GetRatingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRatingRequest{}, err
	}
	return req, nil
}

func NewGetRatingRequestFromDict(data map[string]interface{}) GetRatingRequest {
	return GetRatingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
	}
}

func (p GetRatingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingRequest) Pointer() *GetRatingRequest {
	return &p
}

type GetRatingByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RatingName      *string `json:"ratingName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetRatingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRatingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRatingByUserIdRequest{}
	} else {
		*p = GetRatingByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewGetRatingByUserIdRequestFromJson(data string) (GetRatingByUserIdRequest, error) {
	req := GetRatingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRatingByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetRatingByUserIdRequestFromDict(data map[string]interface{}) GetRatingByUserIdRequest {
	return GetRatingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p GetRatingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"ratingName":      p.RatingName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetRatingByUserIdRequest) Pointer() *GetRatingByUserIdRequest {
	return &p
}

type PutResultRequest struct {
	SourceRequestId *string      `json:"sourceRequestId"`
	RequestId       *string      `json:"requestId"`
	ContextStack    *string      `json:"contextStack"`
	NamespaceName   *string      `json:"namespaceName"`
	RatingName      *string      `json:"ratingName"`
	GameResults     []GameResult `json:"gameResults"`
}

func (p *PutResultRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutResultRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutResultRequest{}
	} else {
		*p = PutResultRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gameResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GameResults)
		}
	}
	return nil
}

func NewPutResultRequestFromJson(data string) (PutResultRequest, error) {
	req := PutResultRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutResultRequest{}, err
	}
	return req, nil
}

func NewPutResultRequestFromDict(data map[string]interface{}) PutResultRequest {
	return PutResultRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GameResults: func() []GameResult {
			if data["gameResults"] == nil {
				return nil
			}
			return CastGameResults(core.CastArray(data["gameResults"]))
		}(),
	}
}

func (p PutResultRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
	}
}

func (p PutResultRequest) Pointer() *PutResultRequest {
	return &p
}

type DeleteRatingRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	RatingName         *string `json:"ratingName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteRatingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRatingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRatingRequest{}
	} else {
		*p = DeleteRatingRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDeleteRatingRequestFromJson(data string) (DeleteRatingRequest, error) {
	req := DeleteRatingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRatingRequest{}, err
	}
	return req, nil
}

func NewDeleteRatingRequestFromDict(data map[string]interface{}) DeleteRatingRequest {
	return DeleteRatingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DeleteRatingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"ratingName":      p.RatingName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteRatingRequest) Pointer() *DeleteRatingRequest {
	return &p
}

type GetBallotRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
	AccessToken     *string `json:"accessToken"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
}

func (p *GetBallotRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBallotRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetBallotRequest{}
	} else {
		*p = GetBallotRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["numberOfPlayer"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NumberOfPlayer)
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewGetBallotRequestFromJson(data string) (GetBallotRequest, error) {
	req := GetBallotRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBallotRequest{}, err
	}
	return req, nil
}

func NewGetBallotRequestFromDict(data map[string]interface{}) GetBallotRequest {
	return GetBallotRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		NumberOfPlayer: func() *int32 {
			v, ok := data["numberOfPlayer"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["numberOfPlayer"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p GetBallotRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"ratingName":     p.RatingName,
		"gatheringName":  p.GatheringName,
		"accessToken":    p.AccessToken,
		"numberOfPlayer": p.NumberOfPlayer,
		"keyId":          p.KeyId,
	}
}

func (p GetBallotRequest) Pointer() *GetBallotRequest {
	return &p
}

type GetBallotByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
	UserId          *string `json:"userId"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetBallotByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBallotByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetBallotByUserIdRequest{}
	} else {
		*p = GetBallotByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["numberOfPlayer"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NumberOfPlayer)
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TimeOffsetToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TimeOffsetToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TimeOffsetToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TimeOffsetToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewGetBallotByUserIdRequestFromJson(data string) (GetBallotByUserIdRequest, error) {
	req := GetBallotByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBallotByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetBallotByUserIdRequestFromDict(data map[string]interface{}) GetBallotByUserIdRequest {
	return GetBallotByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		NumberOfPlayer: func() *int32 {
			v, ok := data["numberOfPlayer"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["numberOfPlayer"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p GetBallotByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"ratingName":      p.RatingName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"numberOfPlayer":  p.NumberOfPlayer,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetBallotByUserIdRequest) Pointer() *GetBallotByUserIdRequest {
	return &p
}

type VoteRequest struct {
	SourceRequestId *string      `json:"sourceRequestId"`
	RequestId       *string      `json:"requestId"`
	ContextStack    *string      `json:"contextStack"`
	NamespaceName   *string      `json:"namespaceName"`
	BallotBody      *string      `json:"ballotBody"`
	BallotSignature *string      `json:"ballotSignature"`
	GameResults     []GameResult `json:"gameResults"`
	KeyId           *string      `json:"keyId"`
}

func (p *VoteRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VoteRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VoteRequest{}
	} else {
		*p = VoteRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ballotBody"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BallotBody = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BallotBody = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BallotBody = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BallotBody = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BallotBody = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BallotBody)
				}
			}
		}
		if v, ok := d["ballotSignature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BallotSignature = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BallotSignature = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BallotSignature = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BallotSignature = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BallotSignature = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BallotSignature)
				}
			}
		}
		if v, ok := d["gameResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GameResults)
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVoteRequestFromJson(data string) (VoteRequest, error) {
	req := VoteRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VoteRequest{}, err
	}
	return req, nil
}

func NewVoteRequestFromDict(data map[string]interface{}) VoteRequest {
	return VoteRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		BallotBody: func() *string {
			v, ok := data["ballotBody"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ballotBody"])
		}(),
		BallotSignature: func() *string {
			v, ok := data["ballotSignature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ballotSignature"])
		}(),
		GameResults: func() []GameResult {
			if data["gameResults"] == nil {
				return nil
			}
			return CastGameResults(core.CastArray(data["gameResults"]))
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p VoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"ballotBody":      p.BallotBody,
		"ballotSignature": p.BallotSignature,
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteRequest) Pointer() *VoteRequest {
	return &p
}

type VoteMultipleRequest struct {
	SourceRequestId *string        `json:"sourceRequestId"`
	RequestId       *string        `json:"requestId"`
	ContextStack    *string        `json:"contextStack"`
	NamespaceName   *string        `json:"namespaceName"`
	SignedBallots   []SignedBallot `json:"signedBallots"`
	GameResults     []GameResult   `json:"gameResults"`
	KeyId           *string        `json:"keyId"`
}

func (p *VoteMultipleRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VoteMultipleRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VoteMultipleRequest{}
	} else {
		*p = VoteMultipleRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["signedBallots"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SignedBallots)
		}
		if v, ok := d["gameResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GameResults)
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVoteMultipleRequestFromJson(data string) (VoteMultipleRequest, error) {
	req := VoteMultipleRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VoteMultipleRequest{}, err
	}
	return req, nil
}

func NewVoteMultipleRequestFromDict(data map[string]interface{}) VoteMultipleRequest {
	return VoteMultipleRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		SignedBallots: func() []SignedBallot {
			if data["signedBallots"] == nil {
				return nil
			}
			return CastSignedBallots(core.CastArray(data["signedBallots"]))
		}(),
		GameResults: func() []GameResult {
			if data["gameResults"] == nil {
				return nil
			}
			return CastGameResults(core.CastArray(data["gameResults"]))
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p VoteMultipleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"signedBallots": CastSignedBallotsFromDict(
			p.SignedBallots,
		),
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteMultipleRequest) Pointer() *VoteMultipleRequest {
	return &p
}

type CommitVoteRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
}

func (p *CommitVoteRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CommitVoteRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CommitVoteRequest{}
	} else {
		*p = CommitVoteRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
	}
	return nil
}

func NewCommitVoteRequestFromJson(data string) (CommitVoteRequest, error) {
	req := CommitVoteRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CommitVoteRequest{}, err
	}
	return req, nil
}

func NewCommitVoteRequestFromDict(data map[string]interface{}) CommitVoteRequest {
	return CommitVoteRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
	}
}

func (p CommitVoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"gatheringName": p.GatheringName,
	}
}

func (p CommitVoteRequest) Pointer() *CommitVoteRequest {
	return &p
}
