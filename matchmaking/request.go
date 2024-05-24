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
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
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
		Name:                              core.CastString(data["name"]),
		Description:                       core.CastString(data["description"]),
		EnableRating:                      core.CastBool(data["enableRating"]),
		EnableDisconnectDetection:         core.CastString(data["enableDisconnectDetection"]),
		DisconnectDetectionTimeoutSeconds: core.CastInt32(data["disconnectDetectionTimeoutSeconds"]),
		CreateGatheringTriggerType:        core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		EnableCollaborateSeasonRating:                 core.CastString(data["enableCollaborateSeasonRating"]),
		CollaborateSeasonRatingNamespaceId:            core.CastString(data["collaborateSeasonRatingNamespaceId"]),
		CollaborateSeasonRatingTtl:                    core.CastInt32(data["collaborateSeasonRatingTtl"]),
		ChangeRatingScript:                            NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		ChangeRatingNotification:                      NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
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
		"changeRatingScript":                            p.ChangeRatingScript.ToDict(),
		"joinNotification":                              p.JoinNotification.ToDict(),
		"leaveNotification":                             p.LeaveNotification.ToDict(),
		"completeNotification":                          p.CompleteNotification.ToDict(),
		"changeRatingNotification":                      p.ChangeRatingNotification.ToDict(),
		"logSetting":                                    p.LogSetting.ToDict(),
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
		NamespaceName: core.CastString(data["namespaceName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
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
		NamespaceName:                                 core.CastString(data["namespaceName"]),
		Description:                                   core.CastString(data["description"]),
		EnableRating:                                  core.CastBool(data["enableRating"]),
		EnableDisconnectDetection:                     core.CastString(data["enableDisconnectDetection"]),
		DisconnectDetectionTimeoutSeconds:             core.CastInt32(data["disconnectDetectionTimeoutSeconds"]),
		CreateGatheringTriggerType:                    core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		EnableCollaborateSeasonRating:                 core.CastString(data["enableCollaborateSeasonRating"]),
		CollaborateSeasonRatingNamespaceId:            core.CastString(data["collaborateSeasonRatingNamespaceId"]),
		CollaborateSeasonRatingTtl:                    core.CastInt32(data["collaborateSeasonRatingTtl"]),
		ChangeRatingScript:                            NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		ChangeRatingNotification:                      NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
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
		"changeRatingScript":                            p.ChangeRatingScript.ToDict(),
		"joinNotification":                              p.JoinNotification.ToDict(),
		"leaveNotification":                             p.LeaveNotification.ToDict(),
		"completeNotification":                          p.CompleteNotification.ToDict(),
		"changeRatingNotification":                      p.ChangeRatingNotification.ToDict(),
		"logSetting":                                    p.LogSetting.ToDict(),
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
		NamespaceName: core.CastString(data["namespaceName"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
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
		NamespaceName:     core.CastString(data["namespaceName"]),
		AccessToken:       core.CastString(data["accessToken"]),
		Player:            NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		AttributeRanges:   CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles:   CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:      core.CastStrings(core.CastArray(data["allowUserIds"])),
		ExpiresAt:         core.CastInt64(data["expiresAt"]),
		ExpiresAtTimeSpan: NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer(),
	}
}

func (p CreateGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"player":        p.Player.ToDict(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt":         p.ExpiresAt,
		"expiresAtTimeSpan": p.ExpiresAtTimeSpan.ToDict(),
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
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		Player:            NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		AttributeRanges:   CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles:   CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:      core.CastStrings(core.CastArray(data["allowUserIds"])),
		ExpiresAt:         core.CastInt64(data["expiresAt"]),
		ExpiresAtTimeSpan: NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer(),
		TimeOffsetToken:   core.CastString(data["timeOffsetToken"]),
	}
}

func (p CreateGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"player":        p.Player.ToDict(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt":         p.ExpiresAt,
		"expiresAtTimeSpan": p.ExpiresAtTimeSpan.ToDict(),
		"timeOffsetToken":   p.TimeOffsetToken,
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName:           core.CastString(data["namespaceName"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingByPlayerRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"player":                  p.Player.ToDict(),
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
		NamespaceName:           core.CastString(data["namespaceName"]),
		AccessToken:             core.CastString(data["accessToken"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"accessToken":             p.AccessToken,
		"player":                  p.Player.ToDict(),
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
		NamespaceName:           core.CastString(data["namespaceName"]),
		UserId:                  core.CastString(data["userId"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
		TimeOffsetToken:         core.CastString(data["timeOffsetToken"]),
	}
}

func (p DoMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"userId":                  p.UserId,
		"player":                  p.Player.ToDict(),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
		AccessToken:   core.CastString(data["accessToken"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
		AccessToken:   core.CastString(data["accessToken"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
		AccessToken:   core.CastString(data["accessToken"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
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

type GetCurrentRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentRatingModelMasterRequest{}
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
		*p = GetCurrentRatingModelMasterRequest{}
	} else {
		*p = GetCurrentRatingModelMasterRequest{}
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

func NewGetCurrentRatingModelMasterRequestFromJson(data string) (GetCurrentRatingModelMasterRequest, error) {
	req := GetCurrentRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) GetCurrentRatingModelMasterRequest {
	return GetCurrentRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentRatingModelMasterRequest) Pointer() *GetCurrentRatingModelMasterRequest {
	return &p
}

type UpdateCurrentRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentRatingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentRatingModelMasterRequest{}
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
		*p = UpdateCurrentRatingModelMasterRequest{}
	} else {
		*p = UpdateCurrentRatingModelMasterRequest{}
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

func NewUpdateCurrentRatingModelMasterRequestFromJson(data string) (UpdateCurrentRatingModelMasterRequest, error) {
	req := UpdateCurrentRatingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentRatingModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterRequest {
	return UpdateCurrentRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentRatingModelMasterRequest) Pointer() *UpdateCurrentRatingModelMasterRequest {
	return &p
}

type UpdateCurrentRatingModelMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentRatingModelMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentRatingModelMasterFromGitHubRequest{}
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
		*p = UpdateCurrentRatingModelMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentRatingModelMasterFromGitHubRequest{}
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

func NewUpdateCurrentRatingModelMasterFromGitHubRequestFromJson(data string) (UpdateCurrentRatingModelMasterFromGitHubRequest, error) {
	req := UpdateCurrentRatingModelMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentRatingModelMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentRatingModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterFromGitHubRequest {
	return UpdateCurrentRatingModelMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) Pointer() *UpdateCurrentRatingModelMasterFromGitHubRequest {
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
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		RatingName:    core.CastString(data["ratingName"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		RatingName:      core.CastString(data["ratingName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		GameResults:   CastGameResults(core.CastArray(data["gameResults"])),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		RatingName:      core.CastString(data["ratingName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName:  core.CastString(data["namespaceName"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
		KeyId:          core.CastString(data["keyId"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		RatingName:      core.CastString(data["ratingName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		NumberOfPlayer:  core.CastInt32(data["numberOfPlayer"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
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
		NamespaceName:   core.CastString(data["namespaceName"]),
		BallotBody:      core.CastString(data["ballotBody"]),
		BallotSignature: core.CastString(data["ballotSignature"]),
		GameResults:     CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:           core.CastString(data["keyId"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		SignedBallots: CastSignedBallots(core.CastArray(data["signedBallots"])),
		GameResults:   CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:         core.CastString(data["keyId"]),
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
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		GatheringName: core.CastString(data["gatheringName"]),
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
