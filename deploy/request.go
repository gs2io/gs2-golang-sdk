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

package deploy

import "core"

type DescribeStacksRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeStacksRequestFromDict(data map[string]interface{}) DescribeStacksRequest {
	return DescribeStacksRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeStacksRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeStacksRequest) Pointer() *DescribeStacksRequest {
	return &p
}

type CreateStackRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Template           *string `json:"template"`
}

func NewCreateStackRequestFromDict(data map[string]interface{}) CreateStackRequest {
	return CreateStackRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Template:    core.CastString(data["template"]),
	}
}

func (p CreateStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"template":    p.Template,
	}
}

func (p CreateStackRequest) Pointer() *CreateStackRequest {
	return &p
}

type CreateStackFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	Name               *string                `json:"name"`
	Description        *string                `json:"description"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewCreateStackFromGitHubRequestFromDict(data map[string]interface{}) CreateStackFromGitHubRequest {
	return CreateStackFromGitHubRequest{
		Name:            core.CastString(data["name"]),
		Description:     core.CastString(data["description"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p CreateStackFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":            p.Name,
		"description":     p.Description,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p CreateStackFromGitHubRequest) Pointer() *CreateStackFromGitHubRequest {
	return &p
}

type ValidateRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	Template           *string `json:"template"`
}

func NewValidateRequestFromDict(data map[string]interface{}) ValidateRequest {
	return ValidateRequest{
		Template: core.CastString(data["template"]),
	}
}

func (p ValidateRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"template": p.Template,
	}
}

func (p ValidateRequest) Pointer() *ValidateRequest {
	return &p
}

type GetStackStatusRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewGetStackStatusRequestFromDict(data map[string]interface{}) GetStackStatusRequest {
	return GetStackStatusRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p GetStackStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p GetStackStatusRequest) Pointer() *GetStackStatusRequest {
	return &p
}

type GetStackRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewGetStackRequestFromDict(data map[string]interface{}) GetStackRequest {
	return GetStackRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p GetStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p GetStackRequest) Pointer() *GetStackRequest {
	return &p
}

type UpdateStackRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	Description        *string `json:"description"`
	Template           *string `json:"template"`
}

func NewUpdateStackRequestFromDict(data map[string]interface{}) UpdateStackRequest {
	return UpdateStackRequest{
		StackName:   core.CastString(data["stackName"]),
		Description: core.CastString(data["description"]),
		Template:    core.CastString(data["template"]),
	}
}

func (p UpdateStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":   p.StackName,
		"description": p.Description,
		"template":    p.Template,
	}
}

func (p UpdateStackRequest) Pointer() *UpdateStackRequest {
	return &p
}

type UpdateStackFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	StackName          *string                `json:"stackName"`
	Description        *string                `json:"description"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateStackFromGitHubRequestFromDict(data map[string]interface{}) UpdateStackFromGitHubRequest {
	return UpdateStackFromGitHubRequest{
		StackName:       core.CastString(data["stackName"]),
		Description:     core.CastString(data["description"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateStackFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":       p.StackName,
		"description":     p.Description,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateStackFromGitHubRequest) Pointer() *UpdateStackFromGitHubRequest {
	return &p
}

type DeleteStackRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewDeleteStackRequestFromDict(data map[string]interface{}) DeleteStackRequest {
	return DeleteStackRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p DeleteStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackRequest) Pointer() *DeleteStackRequest {
	return &p
}

type ForceDeleteStackRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewForceDeleteStackRequestFromDict(data map[string]interface{}) ForceDeleteStackRequest {
	return ForceDeleteStackRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p ForceDeleteStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p ForceDeleteStackRequest) Pointer() *ForceDeleteStackRequest {
	return &p
}

type DeleteStackResourcesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewDeleteStackResourcesRequestFromDict(data map[string]interface{}) DeleteStackResourcesRequest {
	return DeleteStackResourcesRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p DeleteStackResourcesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackResourcesRequest) Pointer() *DeleteStackResourcesRequest {
	return &p
}

type DeleteStackEntityRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
}

func NewDeleteStackEntityRequestFromDict(data map[string]interface{}) DeleteStackEntityRequest {
	return DeleteStackEntityRequest{
		StackName: core.CastString(data["stackName"]),
	}
}

func (p DeleteStackEntityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackEntityRequest) Pointer() *DeleteStackEntityRequest {
	return &p
}

type DescribeResourcesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeResourcesRequestFromDict(data map[string]interface{}) DescribeResourcesRequest {
	return DescribeResourcesRequest{
		StackName: core.CastString(data["stackName"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeResourcesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeResourcesRequest) Pointer() *DescribeResourcesRequest {
	return &p
}

type GetResourceRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	ResourceName       *string `json:"resourceName"`
}

func NewGetResourceRequestFromDict(data map[string]interface{}) GetResourceRequest {
	return GetResourceRequest{
		StackName:    core.CastString(data["stackName"]),
		ResourceName: core.CastString(data["resourceName"]),
	}
}

func (p GetResourceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":    p.StackName,
		"resourceName": p.ResourceName,
	}
}

func (p GetResourceRequest) Pointer() *GetResourceRequest {
	return &p
}

type DescribeEventsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeEventsRequestFromDict(data map[string]interface{}) DescribeEventsRequest {
	return DescribeEventsRequest{
		StackName: core.CastString(data["stackName"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeEventsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeEventsRequest) Pointer() *DescribeEventsRequest {
	return &p
}

type GetEventRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	EventName          *string `json:"eventName"`
}

func NewGetEventRequestFromDict(data map[string]interface{}) GetEventRequest {
	return GetEventRequest{
		StackName: core.CastString(data["stackName"]),
		EventName: core.CastString(data["eventName"]),
	}
}

func (p GetEventRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"eventName": p.EventName,
	}
}

func (p GetEventRequest) Pointer() *GetEventRequest {
	return &p
}

type DescribeOutputsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeOutputsRequestFromDict(data map[string]interface{}) DescribeOutputsRequest {
	return DescribeOutputsRequest{
		StackName: core.CastString(data["stackName"]),
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeOutputsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeOutputsRequest) Pointer() *DescribeOutputsRequest {
	return &p
}

type GetOutputRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StackName          *string `json:"stackName"`
	OutputName         *string `json:"outputName"`
}

func NewGetOutputRequestFromDict(data map[string]interface{}) GetOutputRequest {
	return GetOutputRequest{
		StackName:  core.CastString(data["stackName"]),
		OutputName: core.CastString(data["outputName"]),
	}
}

func (p GetOutputRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":  p.StackName,
		"outputName": p.OutputName,
	}
}

func (p GetOutputRequest) Pointer() *GetOutputRequest {
	return &p
}
