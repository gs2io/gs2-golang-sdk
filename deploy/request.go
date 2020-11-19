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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeStacksRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateStackRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	Name         *string            `json:"name"`
	Description  *string            `json:"description"`
	Template     *string            `json:"template"`
}

type CreateStackFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	Name            *string                `json:"name"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type ValidateRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	Template     *string            `json:"template"`
}

type GetStackStatusRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type GetStackRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type UpdateStackRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	Description  *string            `json:"description"`
	Template     *string            `json:"template"`
}

type UpdateStackFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	StackName       *string                `json:"stackName"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DeleteStackRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type ForceDeleteStackRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type DeleteStackResourcesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type DeleteStackEntityRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
}

type DescribeResourcesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type GetResourceRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	ResourceName *string            `json:"resourceName"`
}

type DescribeEventsRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type GetEventRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	EventName    *string            `json:"eventName"`
}

type DescribeOutputsRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type GetOutputRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	StackName    *string            `json:"stackName"`
	OutputName   *string            `json:"outputName"`
}
