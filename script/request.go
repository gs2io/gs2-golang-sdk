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

package script

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNamespacesRequestFromDict(dict2)
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
	RequestId    *string     `json:"requestId"`
	ContextStack *string     `json:"contextStack"`
	Name         *string     `json:"name"`
	Description  *string     `json:"description"`
	LogSetting   *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateNamespaceRequestFromDict(dict2)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceStatusRequestFromDict(dict2)
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
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceRequestFromDict(dict2)
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
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Description   *string     `json:"description"`
	LogSetting    *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateNamespaceRequestFromDict(dict2)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Description:   core.CastString(data["description"]),
		LogSetting:    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"logSetting":    p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteNamespaceRequestFromDict(dict2)
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

type DescribeScriptsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeScriptsRequestFromJson(data string) DescribeScriptsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeScriptsRequestFromDict(dict2)
}

func NewDescribeScriptsRequestFromDict(data map[string]interface{}) DescribeScriptsRequest {
	return DescribeScriptsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeScriptsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeScriptsRequest) Pointer() *DescribeScriptsRequest {
	return &p
}

type CreateScriptRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Script        *string `json:"script"`
}

func NewCreateScriptRequestFromJson(data string) CreateScriptRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateScriptRequestFromDict(dict2)
}

func NewCreateScriptRequestFromDict(data map[string]interface{}) CreateScriptRequest {
	return CreateScriptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Script:        core.CastString(data["script"]),
	}
}

func (p CreateScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"script":        p.Script,
	}
}

func (p CreateScriptRequest) Pointer() *CreateScriptRequest {
	return &p
}

type CreateScriptFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	Name            *string                `json:"name"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewCreateScriptFromGitHubRequestFromJson(data string) CreateScriptFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateScriptFromGitHubRequestFromDict(dict2)
}

func NewCreateScriptFromGitHubRequestFromDict(data map[string]interface{}) CreateScriptFromGitHubRequest {
	return CreateScriptFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Name:            core.CastString(data["name"]),
		Description:     core.CastString(data["description"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p CreateScriptFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"name":            p.Name,
		"description":     p.Description,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p CreateScriptFromGitHubRequest) Pointer() *CreateScriptFromGitHubRequest {
	return &p
}

type GetScriptRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ScriptName    *string `json:"scriptName"`
}

func NewGetScriptRequestFromJson(data string) GetScriptRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetScriptRequestFromDict(dict2)
}

func NewGetScriptRequestFromDict(data map[string]interface{}) GetScriptRequest {
	return GetScriptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ScriptName:    core.CastString(data["scriptName"]),
	}
}

func (p GetScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"scriptName":    p.ScriptName,
	}
}

func (p GetScriptRequest) Pointer() *GetScriptRequest {
	return &p
}

type UpdateScriptRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ScriptName    *string `json:"scriptName"`
	Description   *string `json:"description"`
	Script        *string `json:"script"`
}

func NewUpdateScriptRequestFromJson(data string) UpdateScriptRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateScriptRequestFromDict(dict2)
}

func NewUpdateScriptRequestFromDict(data map[string]interface{}) UpdateScriptRequest {
	return UpdateScriptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ScriptName:    core.CastString(data["scriptName"]),
		Description:   core.CastString(data["description"]),
		Script:        core.CastString(data["script"]),
	}
}

func (p UpdateScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"scriptName":    p.ScriptName,
		"description":   p.Description,
		"script":        p.Script,
	}
}

func (p UpdateScriptRequest) Pointer() *UpdateScriptRequest {
	return &p
}

type UpdateScriptFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	ScriptName      *string                `json:"scriptName"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateScriptFromGitHubRequestFromJson(data string) UpdateScriptFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateScriptFromGitHubRequestFromDict(dict2)
}

func NewUpdateScriptFromGitHubRequestFromDict(data map[string]interface{}) UpdateScriptFromGitHubRequest {
	return UpdateScriptFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		ScriptName:      core.CastString(data["scriptName"]),
		Description:     core.CastString(data["description"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateScriptFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"scriptName":      p.ScriptName,
		"description":     p.Description,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateScriptFromGitHubRequest) Pointer() *UpdateScriptFromGitHubRequest {
	return &p
}

type DeleteScriptRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ScriptName    *string `json:"scriptName"`
}

func NewDeleteScriptRequestFromJson(data string) DeleteScriptRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteScriptRequestFromDict(dict2)
}

func NewDeleteScriptRequestFromDict(data map[string]interface{}) DeleteScriptRequest {
	return DeleteScriptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ScriptName:    core.CastString(data["scriptName"]),
	}
}

func (p DeleteScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"scriptName":    p.ScriptName,
	}
}

func (p DeleteScriptRequest) Pointer() *DeleteScriptRequest {
	return &p
}

type InvokeScriptRequest struct {
	RequestId    *string       `json:"requestId"`
	ContextStack *string       `json:"contextStack"`
	ScriptId     *string       `json:"scriptId"`
	UserId       *string       `json:"userId"`
	Args         *string       `json:"args"`
	RandomStatus *RandomStatus `json:"randomStatus"`
}

func NewInvokeScriptRequestFromJson(data string) InvokeScriptRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewInvokeScriptRequestFromDict(dict2)
}

func NewInvokeScriptRequestFromDict(data map[string]interface{}) InvokeScriptRequest {
	return InvokeScriptRequest{
		ScriptId:     core.CastString(data["scriptId"]),
		UserId:       core.CastString(data["userId"]),
		Args:         core.CastString(data["args"]),
		RandomStatus: NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer(),
	}
}

func (p InvokeScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"scriptId":     p.ScriptId,
		"userId":       p.UserId,
		"args":         p.Args,
		"randomStatus": p.RandomStatus.ToDict(),
	}
}

func (p InvokeScriptRequest) Pointer() *InvokeScriptRequest {
	return &p
}

type DebugInvokeRequest struct {
	RequestId    *string       `json:"requestId"`
	ContextStack *string       `json:"contextStack"`
	Script       *string       `json:"script"`
	Args         *string       `json:"args"`
	RandomStatus *RandomStatus `json:"randomStatus"`
}

func NewDebugInvokeRequestFromJson(data string) DebugInvokeRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDebugInvokeRequestFromDict(dict2)
}

func NewDebugInvokeRequestFromDict(data map[string]interface{}) DebugInvokeRequest {
	return DebugInvokeRequest{
		Script:       core.CastString(data["script"]),
		Args:         core.CastString(data["args"]),
		RandomStatus: NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer(),
	}
}

func (p DebugInvokeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"script":       p.Script,
		"args":         p.Args,
		"randomStatus": p.RandomStatus.ToDict(),
	}
}

func (p DebugInvokeRequest) Pointer() *DebugInvokeRequest {
	return &p
}
