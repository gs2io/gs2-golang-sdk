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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesRequestFromDict(dict)
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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusRequestFromDict(dict)
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

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceRequestFromDict(dict)
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
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	LogSetting         *LogSetting         `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"transactionSetting": p.TransactionSetting.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
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

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeScriptsRequestFromJson(data string) DescribeScriptsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScriptsRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Script          *string `json:"script"`
}

func NewCreateScriptRequestFromJson(data string) CreateScriptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateScriptRequestFromDict(dict)
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
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	Name            *string                `json:"name"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewCreateScriptFromGitHubRequestFromJson(data string) CreateScriptFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateScriptFromGitHubRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ScriptName      *string `json:"scriptName"`
}

func NewGetScriptRequestFromJson(data string) GetScriptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScriptRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ScriptName      *string `json:"scriptName"`
	Description     *string `json:"description"`
	Script          *string `json:"script"`
}

func NewUpdateScriptRequestFromJson(data string) UpdateScriptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateScriptRequestFromDict(dict)
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
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	ScriptName      *string                `json:"scriptName"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateScriptFromGitHubRequestFromJson(data string) UpdateScriptFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateScriptFromGitHubRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ScriptName      *string `json:"scriptName"`
}

func NewDeleteScriptRequestFromJson(data string) DeleteScriptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteScriptRequestFromDict(dict)
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
	SourceRequestId    *string       `json:"sourceRequestId"`
	RequestId          *string       `json:"requestId"`
	ContextStack       *string       `json:"contextStack"`
	DuplicationAvoider *string       `json:"duplicationAvoider"`
	ScriptId           *string       `json:"scriptId"`
	UserId             *string       `json:"userId"`
	Args               *string       `json:"args"`
	RandomStatus       *RandomStatus `json:"randomStatus"`
	TimeOffsetToken    *string       `json:"timeOffsetToken"`
}

func NewInvokeScriptRequestFromJson(data string) InvokeScriptRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInvokeScriptRequestFromDict(dict)
}

func NewInvokeScriptRequestFromDict(data map[string]interface{}) InvokeScriptRequest {
	return InvokeScriptRequest{
		ScriptId:        core.CastString(data["scriptId"]),
		UserId:          core.CastString(data["userId"]),
		Args:            core.CastString(data["args"]),
		RandomStatus:    NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer(),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p InvokeScriptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"scriptId":        p.ScriptId,
		"userId":          p.UserId,
		"args":            p.Args,
		"randomStatus":    p.RandomStatus.ToDict(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p InvokeScriptRequest) Pointer() *InvokeScriptRequest {
	return &p
}

type DebugInvokeRequest struct {
	SourceRequestId *string       `json:"sourceRequestId"`
	RequestId       *string       `json:"requestId"`
	ContextStack    *string       `json:"contextStack"`
	Script          *string       `json:"script"`
	Args            *string       `json:"args"`
	RandomStatus    *RandomStatus `json:"randomStatus"`
}

func NewDebugInvokeRequestFromJson(data string) DebugInvokeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDebugInvokeRequestFromDict(dict)
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
