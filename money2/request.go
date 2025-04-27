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

package money2

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
	DryRun       *bool   `json:"dryRun"`
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
	ContextStack                         *string              `json:"contextStack"`
	Name                                 *string              `json:"name"`
	CurrencyUsagePriority                *string              `json:"currencyUsagePriority"`
	Description                          *string              `json:"description"`
	SharedFreeCurrency                   *bool                `json:"sharedFreeCurrency"`
	PlatformSetting                      *PlatformSetting     `json:"platformSetting"`
	DepositBalanceScript                 *ScriptSetting       `json:"depositBalanceScript"`
	WithdrawBalanceScript                *ScriptSetting       `json:"withdrawBalanceScript"`
	SubscribeScript                      *string              `json:"subscribeScript"`
	RenewScript                          *string              `json:"renewScript"`
	UnsubscribeScript                    *string              `json:"unsubscribeScript"`
	TakeOverScript                       *ScriptSetting       `json:"takeOverScript"`
	ChangeSubscriptionStatusNotification *NotificationSetting `json:"changeSubscriptionStatusNotification"`
	LogSetting                           *LogSetting          `json:"logSetting"`
	DryRun                               *bool                `json:"dryRun"`
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
		if v, ok := d["currencyUsagePriority"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CurrencyUsagePriority = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CurrencyUsagePriority = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CurrencyUsagePriority = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CurrencyUsagePriority)
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
		if v, ok := d["sharedFreeCurrency"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SharedFreeCurrency)
		}
		if v, ok := d["platformSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PlatformSetting)
		}
		if v, ok := d["depositBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositBalanceScript)
		}
		if v, ok := d["withdrawBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawBalanceScript)
		}
		if v, ok := d["subscribeScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeScript)
				}
			}
		}
		if v, ok := d["renewScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RenewScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RenewScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RenewScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RenewScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RenewScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RenewScript)
				}
			}
		}
		if v, ok := d["unsubscribeScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UnsubscribeScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UnsubscribeScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UnsubscribeScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UnsubscribeScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UnsubscribeScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UnsubscribeScript)
				}
			}
		}
		if v, ok := d["takeOverScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TakeOverScript)
		}
		if v, ok := d["changeSubscriptionStatusNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeSubscriptionStatusNotification)
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
		CurrencyUsagePriority: func() *string {
			v, ok := data["currencyUsagePriority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currencyUsagePriority"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		SharedFreeCurrency: func() *bool {
			v, ok := data["sharedFreeCurrency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sharedFreeCurrency"])
		}(),
		PlatformSetting: func() *PlatformSetting {
			v, ok := data["platformSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformSettingFromDict(core.CastMap(data["platformSetting"])).Pointer()
		}(),
		DepositBalanceScript: func() *ScriptSetting {
			v, ok := data["depositBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["depositBalanceScript"])).Pointer()
		}(),
		WithdrawBalanceScript: func() *ScriptSetting {
			v, ok := data["withdrawBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["withdrawBalanceScript"])).Pointer()
		}(),
		SubscribeScript: func() *string {
			v, ok := data["subscribeScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeScript"])
		}(),
		RenewScript: func() *string {
			v, ok := data["renewScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["renewScript"])
		}(),
		UnsubscribeScript: func() *string {
			v, ok := data["unsubscribeScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["unsubscribeScript"])
		}(),
		TakeOverScript: func() *ScriptSetting {
			v, ok := data["takeOverScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["takeOverScript"])).Pointer()
		}(),
		ChangeSubscriptionStatusNotification: func() *NotificationSetting {
			v, ok := data["changeSubscriptionStatusNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeSubscriptionStatusNotification"])).Pointer()
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
		"name":                  p.Name,
		"currencyUsagePriority": p.CurrencyUsagePriority,
		"description":           p.Description,
		"sharedFreeCurrency":    p.SharedFreeCurrency,
		"platformSetting": func() map[string]interface{} {
			if p.PlatformSetting == nil {
				return nil
			}
			return p.PlatformSetting.ToDict()
		}(),
		"depositBalanceScript": func() map[string]interface{} {
			if p.DepositBalanceScript == nil {
				return nil
			}
			return p.DepositBalanceScript.ToDict()
		}(),
		"withdrawBalanceScript": func() map[string]interface{} {
			if p.WithdrawBalanceScript == nil {
				return nil
			}
			return p.WithdrawBalanceScript.ToDict()
		}(),
		"subscribeScript":   p.SubscribeScript,
		"renewScript":       p.RenewScript,
		"unsubscribeScript": p.UnsubscribeScript,
		"takeOverScript": func() map[string]interface{} {
			if p.TakeOverScript == nil {
				return nil
			}
			return p.TakeOverScript.ToDict()
		}(),
		"changeSubscriptionStatusNotification": func() map[string]interface{} {
			if p.ChangeSubscriptionStatusNotification == nil {
				return nil
			}
			return p.ChangeSubscriptionStatusNotification.ToDict()
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
	ContextStack                         *string              `json:"contextStack"`
	NamespaceName                        *string              `json:"namespaceName"`
	CurrencyUsagePriority                *string              `json:"currencyUsagePriority"`
	Description                          *string              `json:"description"`
	PlatformSetting                      *PlatformSetting     `json:"platformSetting"`
	DepositBalanceScript                 *ScriptSetting       `json:"depositBalanceScript"`
	WithdrawBalanceScript                *ScriptSetting       `json:"withdrawBalanceScript"`
	SubscribeScript                      *string              `json:"subscribeScript"`
	RenewScript                          *string              `json:"renewScript"`
	UnsubscribeScript                    *string              `json:"unsubscribeScript"`
	TakeOverScript                       *ScriptSetting       `json:"takeOverScript"`
	ChangeSubscriptionStatusNotification *NotificationSetting `json:"changeSubscriptionStatusNotification"`
	LogSetting                           *LogSetting          `json:"logSetting"`
	DryRun                               *bool                `json:"dryRun"`
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
		if v, ok := d["currencyUsagePriority"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CurrencyUsagePriority = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CurrencyUsagePriority = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CurrencyUsagePriority = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CurrencyUsagePriority = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CurrencyUsagePriority)
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
		if v, ok := d["platformSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PlatformSetting)
		}
		if v, ok := d["depositBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositBalanceScript)
		}
		if v, ok := d["withdrawBalanceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawBalanceScript)
		}
		if v, ok := d["subscribeScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeScript)
				}
			}
		}
		if v, ok := d["renewScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RenewScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RenewScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RenewScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RenewScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RenewScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RenewScript)
				}
			}
		}
		if v, ok := d["unsubscribeScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UnsubscribeScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UnsubscribeScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UnsubscribeScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UnsubscribeScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UnsubscribeScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UnsubscribeScript)
				}
			}
		}
		if v, ok := d["takeOverScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TakeOverScript)
		}
		if v, ok := d["changeSubscriptionStatusNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeSubscriptionStatusNotification)
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
		CurrencyUsagePriority: func() *string {
			v, ok := data["currencyUsagePriority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currencyUsagePriority"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		PlatformSetting: func() *PlatformSetting {
			v, ok := data["platformSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformSettingFromDict(core.CastMap(data["platformSetting"])).Pointer()
		}(),
		DepositBalanceScript: func() *ScriptSetting {
			v, ok := data["depositBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["depositBalanceScript"])).Pointer()
		}(),
		WithdrawBalanceScript: func() *ScriptSetting {
			v, ok := data["withdrawBalanceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["withdrawBalanceScript"])).Pointer()
		}(),
		SubscribeScript: func() *string {
			v, ok := data["subscribeScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeScript"])
		}(),
		RenewScript: func() *string {
			v, ok := data["renewScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["renewScript"])
		}(),
		UnsubscribeScript: func() *string {
			v, ok := data["unsubscribeScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["unsubscribeScript"])
		}(),
		TakeOverScript: func() *ScriptSetting {
			v, ok := data["takeOverScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["takeOverScript"])).Pointer()
		}(),
		ChangeSubscriptionStatusNotification: func() *NotificationSetting {
			v, ok := data["changeSubscriptionStatusNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeSubscriptionStatusNotification"])).Pointer()
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
		"namespaceName":         p.NamespaceName,
		"currencyUsagePriority": p.CurrencyUsagePriority,
		"description":           p.Description,
		"platformSetting": func() map[string]interface{} {
			if p.PlatformSetting == nil {
				return nil
			}
			return p.PlatformSetting.ToDict()
		}(),
		"depositBalanceScript": func() map[string]interface{} {
			if p.DepositBalanceScript == nil {
				return nil
			}
			return p.DepositBalanceScript.ToDict()
		}(),
		"withdrawBalanceScript": func() map[string]interface{} {
			if p.WithdrawBalanceScript == nil {
				return nil
			}
			return p.WithdrawBalanceScript.ToDict()
		}(),
		"subscribeScript":   p.SubscribeScript,
		"renewScript":       p.RenewScript,
		"unsubscribeScript": p.UnsubscribeScript,
		"takeOverScript": func() map[string]interface{} {
			if p.TakeOverScript == nil {
				return nil
			}
			return p.TakeOverScript.ToDict()
		}(),
		"changeSubscriptionStatusNotification": func() map[string]interface{} {
			if p.ChangeSubscriptionStatusNotification == nil {
				return nil
			}
			return p.ChangeSubscriptionStatusNotification.ToDict()
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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

type DescribeWalletsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeWalletsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeWalletsRequest{}
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
		*p = DescribeWalletsRequest{}
	} else {
		*p = DescribeWalletsRequest{}
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

func NewDescribeWalletsRequestFromJson(data string) (DescribeWalletsRequest, error) {
	req := DescribeWalletsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeWalletsRequest{}, err
	}
	return req, nil
}

func NewDescribeWalletsRequestFromDict(data map[string]interface{}) DescribeWalletsRequest {
	return DescribeWalletsRequest{
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

func (p DescribeWalletsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeWalletsRequest) Pointer() *DescribeWalletsRequest {
	return &p
}

type DescribeWalletsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeWalletsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeWalletsByUserIdRequest{}
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
		*p = DescribeWalletsByUserIdRequest{}
	} else {
		*p = DescribeWalletsByUserIdRequest{}
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

func NewDescribeWalletsByUserIdRequestFromJson(data string) (DescribeWalletsByUserIdRequest, error) {
	req := DescribeWalletsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeWalletsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeWalletsByUserIdRequestFromDict(data map[string]interface{}) DescribeWalletsByUserIdRequest {
	return DescribeWalletsByUserIdRequest{
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

func (p DescribeWalletsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeWalletsByUserIdRequest) Pointer() *DescribeWalletsByUserIdRequest {
	return &p
}

type GetWalletRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	Slot          *int32  `json:"slot"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetWalletRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetWalletRequest{}
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
		*p = GetWalletRequest{}
	} else {
		*p = GetWalletRequest{}
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
	}
	return nil
}

func NewGetWalletRequestFromJson(data string) (GetWalletRequest, error) {
	req := GetWalletRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetWalletRequest{}, err
	}
	return req, nil
}

func NewGetWalletRequestFromDict(data map[string]interface{}) GetWalletRequest {
	return GetWalletRequest{
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
	}
}

func (p GetWalletRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"slot":          p.Slot,
	}
}

func (p GetWalletRequest) Pointer() *GetWalletRequest {
	return &p
}

type GetWalletByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	Slot            *int32  `json:"slot"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetWalletByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetWalletByUserIdRequest{}
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
		*p = GetWalletByUserIdRequest{}
	} else {
		*p = GetWalletByUserIdRequest{}
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
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

func NewGetWalletByUserIdRequestFromJson(data string) (GetWalletByUserIdRequest, error) {
	req := GetWalletByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetWalletByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetWalletByUserIdRequestFromDict(data map[string]interface{}) GetWalletByUserIdRequest {
	return GetWalletByUserIdRequest{
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
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

func (p GetWalletByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetWalletByUserIdRequest) Pointer() *GetWalletByUserIdRequest {
	return &p
}

type DepositByUserIdRequest struct {
	ContextStack        *string              `json:"contextStack"`
	DuplicationAvoider  *string              `json:"duplicationAvoider"`
	NamespaceName       *string              `json:"namespaceName"`
	UserId              *string              `json:"userId"`
	Slot                *int32               `json:"slot"`
	DepositTransactions []DepositTransaction `json:"depositTransactions"`
	TimeOffsetToken     *string              `json:"timeOffsetToken"`
	DryRun              *bool                `json:"dryRun"`
}

func (p *DepositByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DepositByUserIdRequest{}
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
		*p = DepositByUserIdRequest{}
	} else {
		*p = DepositByUserIdRequest{}
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["depositTransactions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DepositTransactions)
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

func NewDepositByUserIdRequestFromJson(data string) (DepositByUserIdRequest, error) {
	req := DepositByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DepositByUserIdRequest{}, err
	}
	return req, nil
}

func NewDepositByUserIdRequestFromDict(data map[string]interface{}) DepositByUserIdRequest {
	return DepositByUserIdRequest{
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		DepositTransactions: func() []DepositTransaction {
			if data["depositTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["depositTransactions"]))
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

func (p DepositByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"slot":          p.Slot,
		"depositTransactions": CastDepositTransactionsFromDict(
			p.DepositTransactions,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DepositByUserIdRequest) Pointer() *DepositByUserIdRequest {
	return &p
}

type WithdrawRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Slot               *int32  `json:"slot"`
	WithdrawCount      *int32  `json:"withdrawCount"`
	PaidOnly           *bool   `json:"paidOnly"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *WithdrawRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawRequest{}
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
		*p = WithdrawRequest{}
	} else {
		*p = WithdrawRequest{}
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["withdrawCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawCount)
		}
		if v, ok := d["paidOnly"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PaidOnly)
		}
	}
	return nil
}

func NewWithdrawRequestFromJson(data string) (WithdrawRequest, error) {
	req := WithdrawRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WithdrawRequest{}, err
	}
	return req, nil
}

func NewWithdrawRequestFromDict(data map[string]interface{}) WithdrawRequest {
	return WithdrawRequest{
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		WithdrawCount: func() *int32 {
			v, ok := data["withdrawCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["withdrawCount"])
		}(),
		PaidOnly: func() *bool {
			v, ok := data["paidOnly"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["paidOnly"])
		}(),
	}
}

func (p WithdrawRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"slot":          p.Slot,
		"withdrawCount": p.WithdrawCount,
		"paidOnly":      p.PaidOnly,
	}
}

func (p WithdrawRequest) Pointer() *WithdrawRequest {
	return &p
}

type WithdrawByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Slot               *int32  `json:"slot"`
	WithdrawCount      *int32  `json:"withdrawCount"`
	PaidOnly           *bool   `json:"paidOnly"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *WithdrawByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawByUserIdRequest{}
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
		*p = WithdrawByUserIdRequest{}
	} else {
		*p = WithdrawByUserIdRequest{}
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
		if v, ok := d["slot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Slot)
		}
		if v, ok := d["withdrawCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WithdrawCount)
		}
		if v, ok := d["paidOnly"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PaidOnly)
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

func NewWithdrawByUserIdRequestFromJson(data string) (WithdrawByUserIdRequest, error) {
	req := WithdrawByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WithdrawByUserIdRequest{}, err
	}
	return req, nil
}

func NewWithdrawByUserIdRequestFromDict(data map[string]interface{}) WithdrawByUserIdRequest {
	return WithdrawByUserIdRequest{
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
		Slot: func() *int32 {
			v, ok := data["slot"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["slot"])
		}(),
		WithdrawCount: func() *int32 {
			v, ok := data["withdrawCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["withdrawCount"])
		}(),
		PaidOnly: func() *bool {
			v, ok := data["paidOnly"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["paidOnly"])
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

func (p WithdrawByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"slot":            p.Slot,
		"withdrawCount":   p.WithdrawCount,
		"paidOnly":        p.PaidOnly,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p WithdrawByUserIdRequest) Pointer() *WithdrawByUserIdRequest {
	return &p
}

type DepositByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *DepositByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DepositByStampSheetRequest{}
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
		*p = DepositByStampSheetRequest{}
	} else {
		*p = DepositByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
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

func NewDepositByStampSheetRequestFromJson(data string) (DepositByStampSheetRequest, error) {
	req := DepositByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DepositByStampSheetRequest{}, err
	}
	return req, nil
}

func NewDepositByStampSheetRequestFromDict(data map[string]interface{}) DepositByStampSheetRequest {
	return DepositByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
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

func (p DepositByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DepositByStampSheetRequest) Pointer() *DepositByStampSheetRequest {
	return &p
}

type WithdrawByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *WithdrawByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawByStampTaskRequest{}
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
		*p = WithdrawByStampTaskRequest{}
	} else {
		*p = WithdrawByStampTaskRequest{}
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

func NewWithdrawByStampTaskRequestFromJson(data string) (WithdrawByStampTaskRequest, error) {
	req := WithdrawByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WithdrawByStampTaskRequest{}, err
	}
	return req, nil
}

func NewWithdrawByStampTaskRequestFromDict(data map[string]interface{}) WithdrawByStampTaskRequest {
	return WithdrawByStampTaskRequest{
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

func (p WithdrawByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p WithdrawByStampTaskRequest) Pointer() *WithdrawByStampTaskRequest {
	return &p
}

type DescribeEventsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeEventsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeEventsByUserIdRequest{}
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
		*p = DescribeEventsByUserIdRequest{}
	} else {
		*p = DescribeEventsByUserIdRequest{}
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
		if v, ok := d["begin"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Begin)
		}
		if v, ok := d["end"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.End)
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

func NewDescribeEventsByUserIdRequestFromJson(data string) (DescribeEventsByUserIdRequest, error) {
	req := DescribeEventsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeEventsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeEventsByUserIdRequestFromDict(data map[string]interface{}) DescribeEventsByUserIdRequest {
	return DescribeEventsByUserIdRequest{
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
		Begin: func() *int64 {
			v, ok := data["begin"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["begin"])
		}(),
		End: func() *int64 {
			v, ok := data["end"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["end"])
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

func (p DescribeEventsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeEventsByUserIdRequest) Pointer() *DescribeEventsByUserIdRequest {
	return &p
}

type GetEventByTransactionIdRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	TransactionId *string `json:"transactionId"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetEventByTransactionIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetEventByTransactionIdRequest{}
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
		*p = GetEventByTransactionIdRequest{}
	} else {
		*p = GetEventByTransactionIdRequest{}
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
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetEventByTransactionIdRequestFromJson(data string) (GetEventByTransactionIdRequest, error) {
	req := GetEventByTransactionIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetEventByTransactionIdRequest{}, err
	}
	return req, nil
}

func NewGetEventByTransactionIdRequestFromDict(data map[string]interface{}) GetEventByTransactionIdRequest {
	return GetEventByTransactionIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetEventByTransactionIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"transactionId": p.TransactionId,
	}
}

func (p GetEventByTransactionIdRequest) Pointer() *GetEventByTransactionIdRequest {
	return &p
}

type VerifyReceiptRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	ContentName        *string  `json:"contentName"`
	Receipt            *Receipt `json:"receipt"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *VerifyReceiptRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReceiptRequest{}
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
		*p = VerifyReceiptRequest{}
	} else {
		*p = VerifyReceiptRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
		if v, ok := d["receipt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Receipt)
		}
	}
	return nil
}

func NewVerifyReceiptRequestFromJson(data string) (VerifyReceiptRequest, error) {
	req := VerifyReceiptRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReceiptRequest{}, err
	}
	return req, nil
}

func NewVerifyReceiptRequestFromDict(data map[string]interface{}) VerifyReceiptRequest {
	return VerifyReceiptRequest{
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
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		Receipt: func() *Receipt {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiptFromDict(core.CastMap(data["receipt"])).Pointer()
		}(),
	}
}

func (p VerifyReceiptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"contentName":   p.ContentName,
		"receipt": func() map[string]interface{} {
			if p.Receipt == nil {
				return nil
			}
			return p.Receipt.ToDict()
		}(),
	}
}

func (p VerifyReceiptRequest) Pointer() *VerifyReceiptRequest {
	return &p
}

type VerifyReceiptByUserIdRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	ContentName        *string  `json:"contentName"`
	Receipt            *Receipt `json:"receipt"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *VerifyReceiptByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReceiptByUserIdRequest{}
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
		*p = VerifyReceiptByUserIdRequest{}
	} else {
		*p = VerifyReceiptByUserIdRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
		if v, ok := d["receipt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Receipt)
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

func NewVerifyReceiptByUserIdRequestFromJson(data string) (VerifyReceiptByUserIdRequest, error) {
	req := VerifyReceiptByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReceiptByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyReceiptByUserIdRequestFromDict(data map[string]interface{}) VerifyReceiptByUserIdRequest {
	return VerifyReceiptByUserIdRequest{
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
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
		Receipt: func() *Receipt {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiptFromDict(core.CastMap(data["receipt"])).Pointer()
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

func (p VerifyReceiptByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"contentName":   p.ContentName,
		"receipt": func() map[string]interface{} {
			if p.Receipt == nil {
				return nil
			}
			return p.Receipt.ToDict()
		}(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p VerifyReceiptByUserIdRequest) Pointer() *VerifyReceiptByUserIdRequest {
	return &p
}

type VerifyReceiptByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyReceiptByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReceiptByStampTaskRequest{}
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
		*p = VerifyReceiptByStampTaskRequest{}
	} else {
		*p = VerifyReceiptByStampTaskRequest{}
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

func NewVerifyReceiptByStampTaskRequestFromJson(data string) (VerifyReceiptByStampTaskRequest, error) {
	req := VerifyReceiptByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReceiptByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyReceiptByStampTaskRequestFromDict(data map[string]interface{}) VerifyReceiptByStampTaskRequest {
	return VerifyReceiptByStampTaskRequest{
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

func (p VerifyReceiptByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyReceiptByStampTaskRequest) Pointer() *VerifyReceiptByStampTaskRequest {
	return &p
}

type DescribeSubscriptionStatusesRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscriptionStatusesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscriptionStatusesRequest{}
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
		*p = DescribeSubscriptionStatusesRequest{}
	} else {
		*p = DescribeSubscriptionStatusesRequest{}
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
	}
	return nil
}

func NewDescribeSubscriptionStatusesRequestFromJson(data string) (DescribeSubscriptionStatusesRequest, error) {
	req := DescribeSubscriptionStatusesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscriptionStatusesRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscriptionStatusesRequestFromDict(data map[string]interface{}) DescribeSubscriptionStatusesRequest {
	return DescribeSubscriptionStatusesRequest{
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
	}
}

func (p DescribeSubscriptionStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p DescribeSubscriptionStatusesRequest) Pointer() *DescribeSubscriptionStatusesRequest {
	return &p
}

type DescribeSubscriptionStatusesByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeSubscriptionStatusesByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscriptionStatusesByUserIdRequest{}
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
		*p = DescribeSubscriptionStatusesByUserIdRequest{}
	} else {
		*p = DescribeSubscriptionStatusesByUserIdRequest{}
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

func NewDescribeSubscriptionStatusesByUserIdRequestFromJson(data string) (DescribeSubscriptionStatusesByUserIdRequest, error) {
	req := DescribeSubscriptionStatusesByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscriptionStatusesByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscriptionStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeSubscriptionStatusesByUserIdRequest {
	return DescribeSubscriptionStatusesByUserIdRequest{
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
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DescribeSubscriptionStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSubscriptionStatusesByUserIdRequest) Pointer() *DescribeSubscriptionStatusesByUserIdRequest {
	return &p
}

type GetSubscriptionStatusRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscriptionStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscriptionStatusRequest{}
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
		*p = GetSubscriptionStatusRequest{}
	} else {
		*p = GetSubscriptionStatusRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewGetSubscriptionStatusRequestFromJson(data string) (GetSubscriptionStatusRequest, error) {
	req := GetSubscriptionStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscriptionStatusRequest{}, err
	}
	return req, nil
}

func NewGetSubscriptionStatusRequestFromDict(data map[string]interface{}) GetSubscriptionStatusRequest {
	return GetSubscriptionStatusRequest{
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
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p GetSubscriptionStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"contentName":   p.ContentName,
	}
}

func (p GetSubscriptionStatusRequest) Pointer() *GetSubscriptionStatusRequest {
	return &p
}

type GetSubscriptionStatusByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	ContentName     *string `json:"contentName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetSubscriptionStatusByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscriptionStatusByUserIdRequest{}
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
		*p = GetSubscriptionStatusByUserIdRequest{}
	} else {
		*p = GetSubscriptionStatusByUserIdRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
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

func NewGetSubscriptionStatusByUserIdRequestFromJson(data string) (GetSubscriptionStatusByUserIdRequest, error) {
	req := GetSubscriptionStatusByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscriptionStatusByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSubscriptionStatusByUserIdRequestFromDict(data map[string]interface{}) GetSubscriptionStatusByUserIdRequest {
	return GetSubscriptionStatusByUserIdRequest{
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
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
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

func (p GetSubscriptionStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"contentName":     p.ContentName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSubscriptionStatusByUserIdRequest) Pointer() *GetSubscriptionStatusByUserIdRequest {
	return &p
}

type AllocateSubscriptionStatusRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Receipt            *string `json:"receipt"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AllocateSubscriptionStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AllocateSubscriptionStatusRequest{}
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
		*p = AllocateSubscriptionStatusRequest{}
	} else {
		*p = AllocateSubscriptionStatusRequest{}
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
		if v, ok := d["receipt"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Receipt = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Receipt = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Receipt = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Receipt)
				}
			}
		}
	}
	return nil
}

func NewAllocateSubscriptionStatusRequestFromJson(data string) (AllocateSubscriptionStatusRequest, error) {
	req := AllocateSubscriptionStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AllocateSubscriptionStatusRequest{}, err
	}
	return req, nil
}

func NewAllocateSubscriptionStatusRequestFromDict(data map[string]interface{}) AllocateSubscriptionStatusRequest {
	return AllocateSubscriptionStatusRequest{
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
		Receipt: func() *string {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receipt"])
		}(),
	}
}

func (p AllocateSubscriptionStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"receipt":       p.Receipt,
	}
}

func (p AllocateSubscriptionStatusRequest) Pointer() *AllocateSubscriptionStatusRequest {
	return &p
}

type AllocateSubscriptionStatusByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Receipt            *string `json:"receipt"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AllocateSubscriptionStatusByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AllocateSubscriptionStatusByUserIdRequest{}
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
		*p = AllocateSubscriptionStatusByUserIdRequest{}
	} else {
		*p = AllocateSubscriptionStatusByUserIdRequest{}
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
		if v, ok := d["receipt"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Receipt = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Receipt = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Receipt = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Receipt)
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

func NewAllocateSubscriptionStatusByUserIdRequestFromJson(data string) (AllocateSubscriptionStatusByUserIdRequest, error) {
	req := AllocateSubscriptionStatusByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AllocateSubscriptionStatusByUserIdRequest{}, err
	}
	return req, nil
}

func NewAllocateSubscriptionStatusByUserIdRequestFromDict(data map[string]interface{}) AllocateSubscriptionStatusByUserIdRequest {
	return AllocateSubscriptionStatusByUserIdRequest{
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
		Receipt: func() *string {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receipt"])
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

func (p AllocateSubscriptionStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"receipt":         p.Receipt,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AllocateSubscriptionStatusByUserIdRequest) Pointer() *AllocateSubscriptionStatusByUserIdRequest {
	return &p
}

type TakeoverSubscriptionStatusRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	Receipt            *string `json:"receipt"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *TakeoverSubscriptionStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TakeoverSubscriptionStatusRequest{}
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
		*p = TakeoverSubscriptionStatusRequest{}
	} else {
		*p = TakeoverSubscriptionStatusRequest{}
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
		if v, ok := d["receipt"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Receipt = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Receipt = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Receipt = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Receipt)
				}
			}
		}
	}
	return nil
}

func NewTakeoverSubscriptionStatusRequestFromJson(data string) (TakeoverSubscriptionStatusRequest, error) {
	req := TakeoverSubscriptionStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return TakeoverSubscriptionStatusRequest{}, err
	}
	return req, nil
}

func NewTakeoverSubscriptionStatusRequestFromDict(data map[string]interface{}) TakeoverSubscriptionStatusRequest {
	return TakeoverSubscriptionStatusRequest{
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
		Receipt: func() *string {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receipt"])
		}(),
	}
}

func (p TakeoverSubscriptionStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"receipt":       p.Receipt,
	}
}

func (p TakeoverSubscriptionStatusRequest) Pointer() *TakeoverSubscriptionStatusRequest {
	return &p
}

type TakeoverSubscriptionStatusByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Receipt            *string `json:"receipt"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *TakeoverSubscriptionStatusByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TakeoverSubscriptionStatusByUserIdRequest{}
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
		*p = TakeoverSubscriptionStatusByUserIdRequest{}
	} else {
		*p = TakeoverSubscriptionStatusByUserIdRequest{}
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
		if v, ok := d["receipt"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Receipt = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Receipt = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Receipt = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Receipt = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Receipt)
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

func NewTakeoverSubscriptionStatusByUserIdRequestFromJson(data string) (TakeoverSubscriptionStatusByUserIdRequest, error) {
	req := TakeoverSubscriptionStatusByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return TakeoverSubscriptionStatusByUserIdRequest{}, err
	}
	return req, nil
}

func NewTakeoverSubscriptionStatusByUserIdRequestFromDict(data map[string]interface{}) TakeoverSubscriptionStatusByUserIdRequest {
	return TakeoverSubscriptionStatusByUserIdRequest{
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
		Receipt: func() *string {
			v, ok := data["receipt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receipt"])
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

func (p TakeoverSubscriptionStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"receipt":         p.Receipt,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p TakeoverSubscriptionStatusByUserIdRequest) Pointer() *TakeoverSubscriptionStatusByUserIdRequest {
	return &p
}

type DescribeRefundHistoriesByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeRefundHistoriesByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRefundHistoriesByUserIdRequest{}
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
		*p = DescribeRefundHistoriesByUserIdRequest{}
	} else {
		*p = DescribeRefundHistoriesByUserIdRequest{}
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

func NewDescribeRefundHistoriesByUserIdRequestFromJson(data string) (DescribeRefundHistoriesByUserIdRequest, error) {
	req := DescribeRefundHistoriesByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRefundHistoriesByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeRefundHistoriesByUserIdRequestFromDict(data map[string]interface{}) DescribeRefundHistoriesByUserIdRequest {
	return DescribeRefundHistoriesByUserIdRequest{
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

func (p DescribeRefundHistoriesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeRefundHistoriesByUserIdRequest) Pointer() *DescribeRefundHistoriesByUserIdRequest {
	return &p
}

type DescribeRefundHistoriesByDateRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Year          *int32  `json:"year"`
	Month         *int32  `json:"month"`
	Day           *int32  `json:"day"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeRefundHistoriesByDateRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRefundHistoriesByDateRequest{}
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
		*p = DescribeRefundHistoriesByDateRequest{}
	} else {
		*p = DescribeRefundHistoriesByDateRequest{}
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
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["day"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Day)
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

func NewDescribeRefundHistoriesByDateRequestFromJson(data string) (DescribeRefundHistoriesByDateRequest, error) {
	req := DescribeRefundHistoriesByDateRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRefundHistoriesByDateRequest{}, err
	}
	return req, nil
}

func NewDescribeRefundHistoriesByDateRequestFromDict(data map[string]interface{}) DescribeRefundHistoriesByDateRequest {
	return DescribeRefundHistoriesByDateRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
		}(),
		Day: func() *int32 {
			v, ok := data["day"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["day"])
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

func (p DescribeRefundHistoriesByDateRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"year":          p.Year,
		"month":         p.Month,
		"day":           p.Day,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRefundHistoriesByDateRequest) Pointer() *DescribeRefundHistoriesByDateRequest {
	return &p
}

type GetRefundHistoryRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	TransactionId *string `json:"transactionId"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetRefundHistoryRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRefundHistoryRequest{}
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
		*p = GetRefundHistoryRequest{}
	} else {
		*p = GetRefundHistoryRequest{}
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
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetRefundHistoryRequestFromJson(data string) (GetRefundHistoryRequest, error) {
	req := GetRefundHistoryRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRefundHistoryRequest{}, err
	}
	return req, nil
}

func NewGetRefundHistoryRequestFromDict(data map[string]interface{}) GetRefundHistoryRequest {
	return GetRefundHistoryRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetRefundHistoryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"transactionId": p.TransactionId,
	}
}

func (p GetRefundHistoryRequest) Pointer() *GetRefundHistoryRequest {
	return &p
}

type DescribeStoreContentModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeStoreContentModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStoreContentModelsRequest{}
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
		*p = DescribeStoreContentModelsRequest{}
	} else {
		*p = DescribeStoreContentModelsRequest{}
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

func NewDescribeStoreContentModelsRequestFromJson(data string) (DescribeStoreContentModelsRequest, error) {
	req := DescribeStoreContentModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStoreContentModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeStoreContentModelsRequestFromDict(data map[string]interface{}) DescribeStoreContentModelsRequest {
	return DescribeStoreContentModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeStoreContentModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStoreContentModelsRequest) Pointer() *DescribeStoreContentModelsRequest {
	return &p
}

type GetStoreContentModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetStoreContentModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStoreContentModelRequest{}
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
		*p = GetStoreContentModelRequest{}
	} else {
		*p = GetStoreContentModelRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewGetStoreContentModelRequestFromJson(data string) (GetStoreContentModelRequest, error) {
	req := GetStoreContentModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStoreContentModelRequest{}, err
	}
	return req, nil
}

func NewGetStoreContentModelRequestFromDict(data map[string]interface{}) GetStoreContentModelRequest {
	return GetStoreContentModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p GetStoreContentModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p GetStoreContentModelRequest) Pointer() *GetStoreContentModelRequest {
	return &p
}

type DescribeStoreContentModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeStoreContentModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStoreContentModelMastersRequest{}
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
		*p = DescribeStoreContentModelMastersRequest{}
	} else {
		*p = DescribeStoreContentModelMastersRequest{}
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

func NewDescribeStoreContentModelMastersRequestFromJson(data string) (DescribeStoreContentModelMastersRequest, error) {
	req := DescribeStoreContentModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStoreContentModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeStoreContentModelMastersRequestFromDict(data map[string]interface{}) DescribeStoreContentModelMastersRequest {
	return DescribeStoreContentModelMastersRequest{
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

func (p DescribeStoreContentModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStoreContentModelMastersRequest) Pointer() *DescribeStoreContentModelMastersRequest {
	return &p
}

type CreateStoreContentModelMasterRequest struct {
	ContextStack  *string               `json:"contextStack"`
	NamespaceName *string               `json:"namespaceName"`
	Name          *string               `json:"name"`
	Description   *string               `json:"description"`
	Metadata      *string               `json:"metadata"`
	AppleAppStore *AppleAppStoreContent `json:"appleAppStore"`
	GooglePlay    *GooglePlayContent    `json:"googlePlay"`
	DryRun        *bool                 `json:"dryRun"`
}

func (p *CreateStoreContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateStoreContentModelMasterRequest{}
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
		*p = CreateStoreContentModelMasterRequest{}
	} else {
		*p = CreateStoreContentModelMasterRequest{}
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
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
	}
	return nil
}

func NewCreateStoreContentModelMasterRequestFromJson(data string) (CreateStoreContentModelMasterRequest, error) {
	req := CreateStoreContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateStoreContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateStoreContentModelMasterRequestFromDict(data map[string]interface{}) CreateStoreContentModelMasterRequest {
	return CreateStoreContentModelMasterRequest{
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
		AppleAppStore: func() *AppleAppStoreContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlayContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p CreateStoreContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"appleAppStore": func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}(),
		"googlePlay": func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}(),
	}
}

func (p CreateStoreContentModelMasterRequest) Pointer() *CreateStoreContentModelMasterRequest {
	return &p
}

type GetStoreContentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetStoreContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStoreContentModelMasterRequest{}
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
		*p = GetStoreContentModelMasterRequest{}
	} else {
		*p = GetStoreContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewGetStoreContentModelMasterRequestFromJson(data string) (GetStoreContentModelMasterRequest, error) {
	req := GetStoreContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStoreContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetStoreContentModelMasterRequestFromDict(data map[string]interface{}) GetStoreContentModelMasterRequest {
	return GetStoreContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p GetStoreContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p GetStoreContentModelMasterRequest) Pointer() *GetStoreContentModelMasterRequest {
	return &p
}

type UpdateStoreContentModelMasterRequest struct {
	ContextStack  *string               `json:"contextStack"`
	NamespaceName *string               `json:"namespaceName"`
	ContentName   *string               `json:"contentName"`
	Description   *string               `json:"description"`
	Metadata      *string               `json:"metadata"`
	AppleAppStore *AppleAppStoreContent `json:"appleAppStore"`
	GooglePlay    *GooglePlayContent    `json:"googlePlay"`
	DryRun        *bool                 `json:"dryRun"`
}

func (p *UpdateStoreContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStoreContentModelMasterRequest{}
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
		*p = UpdateStoreContentModelMasterRequest{}
	} else {
		*p = UpdateStoreContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
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
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
	}
	return nil
}

func NewUpdateStoreContentModelMasterRequestFromJson(data string) (UpdateStoreContentModelMasterRequest, error) {
	req := UpdateStoreContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStoreContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateStoreContentModelMasterRequestFromDict(data map[string]interface{}) UpdateStoreContentModelMasterRequest {
	return UpdateStoreContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
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
		AppleAppStore: func() *AppleAppStoreContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlayContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlayContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p UpdateStoreContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"appleAppStore": func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}(),
		"googlePlay": func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}(),
	}
}

func (p UpdateStoreContentModelMasterRequest) Pointer() *UpdateStoreContentModelMasterRequest {
	return &p
}

type DeleteStoreContentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteStoreContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStoreContentModelMasterRequest{}
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
		*p = DeleteStoreContentModelMasterRequest{}
	} else {
		*p = DeleteStoreContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewDeleteStoreContentModelMasterRequestFromJson(data string) (DeleteStoreContentModelMasterRequest, error) {
	req := DeleteStoreContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStoreContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteStoreContentModelMasterRequestFromDict(data map[string]interface{}) DeleteStoreContentModelMasterRequest {
	return DeleteStoreContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p DeleteStoreContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p DeleteStoreContentModelMasterRequest) Pointer() *DeleteStoreContentModelMasterRequest {
	return &p
}

type DescribeStoreSubscriptionContentModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeStoreSubscriptionContentModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStoreSubscriptionContentModelsRequest{}
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
		*p = DescribeStoreSubscriptionContentModelsRequest{}
	} else {
		*p = DescribeStoreSubscriptionContentModelsRequest{}
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

func NewDescribeStoreSubscriptionContentModelsRequestFromJson(data string) (DescribeStoreSubscriptionContentModelsRequest, error) {
	req := DescribeStoreSubscriptionContentModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStoreSubscriptionContentModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeStoreSubscriptionContentModelsRequestFromDict(data map[string]interface{}) DescribeStoreSubscriptionContentModelsRequest {
	return DescribeStoreSubscriptionContentModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeStoreSubscriptionContentModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStoreSubscriptionContentModelsRequest) Pointer() *DescribeStoreSubscriptionContentModelsRequest {
	return &p
}

type GetStoreSubscriptionContentModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetStoreSubscriptionContentModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStoreSubscriptionContentModelRequest{}
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
		*p = GetStoreSubscriptionContentModelRequest{}
	} else {
		*p = GetStoreSubscriptionContentModelRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewGetStoreSubscriptionContentModelRequestFromJson(data string) (GetStoreSubscriptionContentModelRequest, error) {
	req := GetStoreSubscriptionContentModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStoreSubscriptionContentModelRequest{}, err
	}
	return req, nil
}

func NewGetStoreSubscriptionContentModelRequestFromDict(data map[string]interface{}) GetStoreSubscriptionContentModelRequest {
	return GetStoreSubscriptionContentModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p GetStoreSubscriptionContentModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p GetStoreSubscriptionContentModelRequest) Pointer() *GetStoreSubscriptionContentModelRequest {
	return &p
}

type DescribeStoreSubscriptionContentModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeStoreSubscriptionContentModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStoreSubscriptionContentModelMastersRequest{}
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
		*p = DescribeStoreSubscriptionContentModelMastersRequest{}
	} else {
		*p = DescribeStoreSubscriptionContentModelMastersRequest{}
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

func NewDescribeStoreSubscriptionContentModelMastersRequestFromJson(data string) (DescribeStoreSubscriptionContentModelMastersRequest, error) {
	req := DescribeStoreSubscriptionContentModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStoreSubscriptionContentModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeStoreSubscriptionContentModelMastersRequestFromDict(data map[string]interface{}) DescribeStoreSubscriptionContentModelMastersRequest {
	return DescribeStoreSubscriptionContentModelMastersRequest{
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

func (p DescribeStoreSubscriptionContentModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStoreSubscriptionContentModelMastersRequest) Pointer() *DescribeStoreSubscriptionContentModelMastersRequest {
	return &p
}

type CreateStoreSubscriptionContentModelMasterRequest struct {
	ContextStack        *string                           `json:"contextStack"`
	NamespaceName       *string                           `json:"namespaceName"`
	Name                *string                           `json:"name"`
	Description         *string                           `json:"description"`
	Metadata            *string                           `json:"metadata"`
	ScheduleNamespaceId *string                           `json:"scheduleNamespaceId"`
	TriggerName         *string                           `json:"triggerName"`
	TriggerExtendMode   *string                           `json:"triggerExtendMode"`
	RollupHour          *int32                            `json:"rollupHour"`
	ReallocateSpanDays  *int32                            `json:"reallocateSpanDays"`
	AppleAppStore       *AppleAppStoreSubscriptionContent `json:"appleAppStore"`
	GooglePlay          *GooglePlaySubscriptionContent    `json:"googlePlay"`
	DryRun              *bool                             `json:"dryRun"`
}

func (p *CreateStoreSubscriptionContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateStoreSubscriptionContentModelMasterRequest{}
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
		*p = CreateStoreSubscriptionContentModelMasterRequest{}
	} else {
		*p = CreateStoreSubscriptionContentModelMasterRequest{}
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
		if v, ok := d["scheduleNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleNamespaceId)
				}
			}
		}
		if v, ok := d["triggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerName)
				}
			}
		}
		if v, ok := d["triggerExtendMode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerExtendMode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerExtendMode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerExtendMode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerExtendMode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerExtendMode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerExtendMode)
				}
			}
		}
		if v, ok := d["rollupHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RollupHour)
		}
		if v, ok := d["reallocateSpanDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReallocateSpanDays)
		}
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
	}
	return nil
}

func NewCreateStoreSubscriptionContentModelMasterRequestFromJson(data string) (CreateStoreSubscriptionContentModelMasterRequest, error) {
	req := CreateStoreSubscriptionContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateStoreSubscriptionContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateStoreSubscriptionContentModelMasterRequestFromDict(data map[string]interface{}) CreateStoreSubscriptionContentModelMasterRequest {
	return CreateStoreSubscriptionContentModelMasterRequest{
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
		ScheduleNamespaceId: func() *string {
			v, ok := data["scheduleNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleNamespaceId"])
		}(),
		TriggerName: func() *string {
			v, ok := data["triggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerName"])
		}(),
		TriggerExtendMode: func() *string {
			v, ok := data["triggerExtendMode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerExtendMode"])
		}(),
		RollupHour: func() *int32 {
			v, ok := data["rollupHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rollupHour"])
		}(),
		ReallocateSpanDays: func() *int32 {
			v, ok := data["reallocateSpanDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["reallocateSpanDays"])
		}(),
		AppleAppStore: func() *AppleAppStoreSubscriptionContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreSubscriptionContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlaySubscriptionContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlaySubscriptionContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p CreateStoreSubscriptionContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"name":                p.Name,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"scheduleNamespaceId": p.ScheduleNamespaceId,
		"triggerName":         p.TriggerName,
		"triggerExtendMode":   p.TriggerExtendMode,
		"rollupHour":          p.RollupHour,
		"reallocateSpanDays":  p.ReallocateSpanDays,
		"appleAppStore": func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}(),
		"googlePlay": func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}(),
	}
}

func (p CreateStoreSubscriptionContentModelMasterRequest) Pointer() *CreateStoreSubscriptionContentModelMasterRequest {
	return &p
}

type GetStoreSubscriptionContentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetStoreSubscriptionContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStoreSubscriptionContentModelMasterRequest{}
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
		*p = GetStoreSubscriptionContentModelMasterRequest{}
	} else {
		*p = GetStoreSubscriptionContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewGetStoreSubscriptionContentModelMasterRequestFromJson(data string) (GetStoreSubscriptionContentModelMasterRequest, error) {
	req := GetStoreSubscriptionContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStoreSubscriptionContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetStoreSubscriptionContentModelMasterRequestFromDict(data map[string]interface{}) GetStoreSubscriptionContentModelMasterRequest {
	return GetStoreSubscriptionContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p GetStoreSubscriptionContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p GetStoreSubscriptionContentModelMasterRequest) Pointer() *GetStoreSubscriptionContentModelMasterRequest {
	return &p
}

type UpdateStoreSubscriptionContentModelMasterRequest struct {
	ContextStack        *string                           `json:"contextStack"`
	NamespaceName       *string                           `json:"namespaceName"`
	ContentName         *string                           `json:"contentName"`
	Description         *string                           `json:"description"`
	Metadata            *string                           `json:"metadata"`
	ScheduleNamespaceId *string                           `json:"scheduleNamespaceId"`
	TriggerName         *string                           `json:"triggerName"`
	TriggerExtendMode   *string                           `json:"triggerExtendMode"`
	RollupHour          *int32                            `json:"rollupHour"`
	ReallocateSpanDays  *int32                            `json:"reallocateSpanDays"`
	AppleAppStore       *AppleAppStoreSubscriptionContent `json:"appleAppStore"`
	GooglePlay          *GooglePlaySubscriptionContent    `json:"googlePlay"`
	DryRun              *bool                             `json:"dryRun"`
}

func (p *UpdateStoreSubscriptionContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStoreSubscriptionContentModelMasterRequest{}
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
		*p = UpdateStoreSubscriptionContentModelMasterRequest{}
	} else {
		*p = UpdateStoreSubscriptionContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
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
		if v, ok := d["scheduleNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleNamespaceId)
				}
			}
		}
		if v, ok := d["triggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerName)
				}
			}
		}
		if v, ok := d["triggerExtendMode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerExtendMode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerExtendMode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerExtendMode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerExtendMode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerExtendMode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerExtendMode)
				}
			}
		}
		if v, ok := d["rollupHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RollupHour)
		}
		if v, ok := d["reallocateSpanDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReallocateSpanDays)
		}
		if v, ok := d["appleAppStore"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppleAppStore)
		}
		if v, ok := d["googlePlay"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GooglePlay)
		}
	}
	return nil
}

func NewUpdateStoreSubscriptionContentModelMasterRequestFromJson(data string) (UpdateStoreSubscriptionContentModelMasterRequest, error) {
	req := UpdateStoreSubscriptionContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStoreSubscriptionContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateStoreSubscriptionContentModelMasterRequestFromDict(data map[string]interface{}) UpdateStoreSubscriptionContentModelMasterRequest {
	return UpdateStoreSubscriptionContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
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
		ScheduleNamespaceId: func() *string {
			v, ok := data["scheduleNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleNamespaceId"])
		}(),
		TriggerName: func() *string {
			v, ok := data["triggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerName"])
		}(),
		TriggerExtendMode: func() *string {
			v, ok := data["triggerExtendMode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerExtendMode"])
		}(),
		RollupHour: func() *int32 {
			v, ok := data["rollupHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rollupHour"])
		}(),
		ReallocateSpanDays: func() *int32 {
			v, ok := data["reallocateSpanDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["reallocateSpanDays"])
		}(),
		AppleAppStore: func() *AppleAppStoreSubscriptionContent {
			v, ok := data["appleAppStore"]
			if !ok || v == nil {
				return nil
			}
			return NewAppleAppStoreSubscriptionContentFromDict(core.CastMap(data["appleAppStore"])).Pointer()
		}(),
		GooglePlay: func() *GooglePlaySubscriptionContent {
			v, ok := data["googlePlay"]
			if !ok || v == nil {
				return nil
			}
			return NewGooglePlaySubscriptionContentFromDict(core.CastMap(data["googlePlay"])).Pointer()
		}(),
	}
}

func (p UpdateStoreSubscriptionContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"contentName":         p.ContentName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"scheduleNamespaceId": p.ScheduleNamespaceId,
		"triggerName":         p.TriggerName,
		"triggerExtendMode":   p.TriggerExtendMode,
		"rollupHour":          p.RollupHour,
		"reallocateSpanDays":  p.ReallocateSpanDays,
		"appleAppStore": func() map[string]interface{} {
			if p.AppleAppStore == nil {
				return nil
			}
			return p.AppleAppStore.ToDict()
		}(),
		"googlePlay": func() map[string]interface{} {
			if p.GooglePlay == nil {
				return nil
			}
			return p.GooglePlay.ToDict()
		}(),
	}
}

func (p UpdateStoreSubscriptionContentModelMasterRequest) Pointer() *UpdateStoreSubscriptionContentModelMasterRequest {
	return &p
}

type DeleteStoreSubscriptionContentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ContentName   *string `json:"contentName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteStoreSubscriptionContentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStoreSubscriptionContentModelMasterRequest{}
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
		*p = DeleteStoreSubscriptionContentModelMasterRequest{}
	} else {
		*p = DeleteStoreSubscriptionContentModelMasterRequest{}
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
		if v, ok := d["contentName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ContentName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ContentName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ContentName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ContentName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ContentName)
				}
			}
		}
	}
	return nil
}

func NewDeleteStoreSubscriptionContentModelMasterRequestFromJson(data string) (DeleteStoreSubscriptionContentModelMasterRequest, error) {
	req := DeleteStoreSubscriptionContentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStoreSubscriptionContentModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteStoreSubscriptionContentModelMasterRequestFromDict(data map[string]interface{}) DeleteStoreSubscriptionContentModelMasterRequest {
	return DeleteStoreSubscriptionContentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ContentName: func() *string {
			v, ok := data["contentName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["contentName"])
		}(),
	}
}

func (p DeleteStoreSubscriptionContentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"contentName":   p.ContentName,
	}
}

func (p DeleteStoreSubscriptionContentModelMasterRequest) Pointer() *DeleteStoreSubscriptionContentModelMasterRequest {
	return &p
}

type ExportMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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

type PreUpdateCurrentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *PreUpdateCurrentModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PreUpdateCurrentModelMasterRequest{}
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
		*p = PreUpdateCurrentModelMasterRequest{}
	} else {
		*p = PreUpdateCurrentModelMasterRequest{}
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

func NewPreUpdateCurrentModelMasterRequestFromJson(data string) (PreUpdateCurrentModelMasterRequest, error) {
	req := PreUpdateCurrentModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PreUpdateCurrentModelMasterRequest{}, err
	}
	return req, nil
}

func NewPreUpdateCurrentModelMasterRequestFromDict(data map[string]interface{}) PreUpdateCurrentModelMasterRequest {
	return PreUpdateCurrentModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p PreUpdateCurrentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p PreUpdateCurrentModelMasterRequest) Pointer() *PreUpdateCurrentModelMasterRequest {
	return &p
}

type UpdateCurrentModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Mode          *string `json:"mode"`
	Settings      *string `json:"settings"`
	UploadToken   *string `json:"uploadToken"`
	DryRun        *bool   `json:"dryRun"`
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
		if v, ok := d["mode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Mode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Mode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Mode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Mode)
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
		Mode: func() *string {
			v, ok := data["mode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["mode"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
	}
}

func (p UpdateCurrentModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"mode":          p.Mode,
		"settings":      p.Settings,
		"uploadToken":   p.UploadToken,
	}
}

func (p UpdateCurrentModelMasterRequest) Pointer() *UpdateCurrentModelMasterRequest {
	return &p
}

type UpdateCurrentModelMasterFromGitHubRequest struct {
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
	DryRun          *bool                  `json:"dryRun"`
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

type DescribeDailyTransactionHistoriesByCurrencyRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Currency      *string `json:"currency"`
	Year          *int32  `json:"year"`
	Month         *int32  `json:"month"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeDailyTransactionHistoriesByCurrencyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDailyTransactionHistoriesByCurrencyRequest{}
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
		*p = DescribeDailyTransactionHistoriesByCurrencyRequest{}
	} else {
		*p = DescribeDailyTransactionHistoriesByCurrencyRequest{}
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
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
				}
			}
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
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

func NewDescribeDailyTransactionHistoriesByCurrencyRequestFromJson(data string) (DescribeDailyTransactionHistoriesByCurrencyRequest, error) {
	req := DescribeDailyTransactionHistoriesByCurrencyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDailyTransactionHistoriesByCurrencyRequest{}, err
	}
	return req, nil
}

func NewDescribeDailyTransactionHistoriesByCurrencyRequestFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesByCurrencyRequest {
	return DescribeDailyTransactionHistoriesByCurrencyRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
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

func (p DescribeDailyTransactionHistoriesByCurrencyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"currency":      p.Currency,
		"year":          p.Year,
		"month":         p.Month,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeDailyTransactionHistoriesByCurrencyRequest) Pointer() *DescribeDailyTransactionHistoriesByCurrencyRequest {
	return &p
}

type DescribeDailyTransactionHistoriesRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Year          *int32  `json:"year"`
	Month         *int32  `json:"month"`
	Day           *int32  `json:"day"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeDailyTransactionHistoriesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDailyTransactionHistoriesRequest{}
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
		*p = DescribeDailyTransactionHistoriesRequest{}
	} else {
		*p = DescribeDailyTransactionHistoriesRequest{}
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
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["day"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Day)
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

func NewDescribeDailyTransactionHistoriesRequestFromJson(data string) (DescribeDailyTransactionHistoriesRequest, error) {
	req := DescribeDailyTransactionHistoriesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDailyTransactionHistoriesRequest{}, err
	}
	return req, nil
}

func NewDescribeDailyTransactionHistoriesRequestFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesRequest {
	return DescribeDailyTransactionHistoriesRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
		}(),
		Day: func() *int32 {
			v, ok := data["day"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["day"])
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

func (p DescribeDailyTransactionHistoriesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"year":          p.Year,
		"month":         p.Month,
		"day":           p.Day,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeDailyTransactionHistoriesRequest) Pointer() *DescribeDailyTransactionHistoriesRequest {
	return &p
}

type GetDailyTransactionHistoryRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Year          *int32  `json:"year"`
	Month         *int32  `json:"month"`
	Day           *int32  `json:"day"`
	Currency      *string `json:"currency"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetDailyTransactionHistoryRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDailyTransactionHistoryRequest{}
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
		*p = GetDailyTransactionHistoryRequest{}
	} else {
		*p = GetDailyTransactionHistoryRequest{}
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
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["day"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Day)
		}
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
				}
			}
		}
	}
	return nil
}

func NewGetDailyTransactionHistoryRequestFromJson(data string) (GetDailyTransactionHistoryRequest, error) {
	req := GetDailyTransactionHistoryRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDailyTransactionHistoryRequest{}, err
	}
	return req, nil
}

func NewGetDailyTransactionHistoryRequestFromDict(data map[string]interface{}) GetDailyTransactionHistoryRequest {
	return GetDailyTransactionHistoryRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
		}(),
		Day: func() *int32 {
			v, ok := data["day"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["day"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
	}
}

func (p GetDailyTransactionHistoryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"year":          p.Year,
		"month":         p.Month,
		"day":           p.Day,
		"currency":      p.Currency,
	}
}

func (p GetDailyTransactionHistoryRequest) Pointer() *GetDailyTransactionHistoryRequest {
	return &p
}

type DescribeUnusedBalancesRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeUnusedBalancesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeUnusedBalancesRequest{}
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
		*p = DescribeUnusedBalancesRequest{}
	} else {
		*p = DescribeUnusedBalancesRequest{}
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

func NewDescribeUnusedBalancesRequestFromJson(data string) (DescribeUnusedBalancesRequest, error) {
	req := DescribeUnusedBalancesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeUnusedBalancesRequest{}, err
	}
	return req, nil
}

func NewDescribeUnusedBalancesRequestFromDict(data map[string]interface{}) DescribeUnusedBalancesRequest {
	return DescribeUnusedBalancesRequest{
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

func (p DescribeUnusedBalancesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeUnusedBalancesRequest) Pointer() *DescribeUnusedBalancesRequest {
	return &p
}

type GetUnusedBalanceRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Currency      *string `json:"currency"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetUnusedBalanceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetUnusedBalanceRequest{}
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
		*p = GetUnusedBalanceRequest{}
	} else {
		*p = GetUnusedBalanceRequest{}
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
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
				}
			}
		}
	}
	return nil
}

func NewGetUnusedBalanceRequestFromJson(data string) (GetUnusedBalanceRequest, error) {
	req := GetUnusedBalanceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetUnusedBalanceRequest{}, err
	}
	return req, nil
}

func NewGetUnusedBalanceRequestFromDict(data map[string]interface{}) GetUnusedBalanceRequest {
	return GetUnusedBalanceRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
	}
}

func (p GetUnusedBalanceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"currency":      p.Currency,
	}
}

func (p GetUnusedBalanceRequest) Pointer() *GetUnusedBalanceRequest {
	return &p
}
