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

import "github.com/gs2io/gs2-golang-sdk/core"

type DescribeStacksResult struct {
	Items         []Stack `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeStacksAsyncResult struct {
	result *DescribeStacksResult
	err    error
}

func NewDescribeStacksResultFromDict(data map[string]interface{}) DescribeStacksResult {
	return DescribeStacksResult{
		Items:         CastStacks(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStacksResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStacksFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStacksResult) Pointer() *DescribeStacksResult {
	return &p
}

type CreateStackResult struct {
	Item *Stack `json:"item"`
}

type CreateStackAsyncResult struct {
	result *CreateStackResult
	err    error
}

func NewCreateStackResultFromDict(data map[string]interface{}) CreateStackResult {
	return CreateStackResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateStackResult) Pointer() *CreateStackResult {
	return &p
}

type CreateStackFromGitHubResult struct {
	Item *Stack `json:"item"`
}

type CreateStackFromGitHubAsyncResult struct {
	result *CreateStackFromGitHubResult
	err    error
}

func NewCreateStackFromGitHubResultFromDict(data map[string]interface{}) CreateStackFromGitHubResult {
	return CreateStackFromGitHubResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateStackFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateStackFromGitHubResult) Pointer() *CreateStackFromGitHubResult {
	return &p
}

type ValidateResult struct {
}

type ValidateAsyncResult struct {
	result *ValidateResult
	err    error
}

func NewValidateResultFromDict(data map[string]interface{}) ValidateResult {
	return ValidateResult{}
}

func (p ValidateResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ValidateResult) Pointer() *ValidateResult {
	return &p
}

type GetStackStatusResult struct {
	Status *string `json:"status"`
}

type GetStackStatusAsyncResult struct {
	result *GetStackStatusResult
	err    error
}

func NewGetStackStatusResultFromDict(data map[string]interface{}) GetStackStatusResult {
	return GetStackStatusResult{
		Status: core.CastString(data["status"]),
	}
}

func (p GetStackStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
	}
}

func (p GetStackStatusResult) Pointer() *GetStackStatusResult {
	return &p
}

type GetStackResult struct {
	Item *Stack `json:"item"`
}

type GetStackAsyncResult struct {
	result *GetStackResult
	err    error
}

func NewGetStackResultFromDict(data map[string]interface{}) GetStackResult {
	return GetStackResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStackResult) Pointer() *GetStackResult {
	return &p
}

type UpdateStackResult struct {
	Item *Stack `json:"item"`
}

type UpdateStackAsyncResult struct {
	result *UpdateStackResult
	err    error
}

func NewUpdateStackResultFromDict(data map[string]interface{}) UpdateStackResult {
	return UpdateStackResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateStackResult) Pointer() *UpdateStackResult {
	return &p
}

type UpdateStackFromGitHubResult struct {
	Item *Stack `json:"item"`
}

type UpdateStackFromGitHubAsyncResult struct {
	result *UpdateStackFromGitHubResult
	err    error
}

func NewUpdateStackFromGitHubResultFromDict(data map[string]interface{}) UpdateStackFromGitHubResult {
	return UpdateStackFromGitHubResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateStackFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateStackFromGitHubResult) Pointer() *UpdateStackFromGitHubResult {
	return &p
}

type DeleteStackResult struct {
	Item *Stack `json:"item"`
}

type DeleteStackAsyncResult struct {
	result *DeleteStackResult
	err    error
}

func NewDeleteStackResultFromDict(data map[string]interface{}) DeleteStackResult {
	return DeleteStackResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStackResult) Pointer() *DeleteStackResult {
	return &p
}

type ForceDeleteStackResult struct {
	Item *Stack `json:"item"`
}

type ForceDeleteStackAsyncResult struct {
	result *ForceDeleteStackResult
	err    error
}

func NewForceDeleteStackResultFromDict(data map[string]interface{}) ForceDeleteStackResult {
	return ForceDeleteStackResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ForceDeleteStackResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ForceDeleteStackResult) Pointer() *ForceDeleteStackResult {
	return &p
}

type DeleteStackResourcesResult struct {
	Item *Stack `json:"item"`
}

type DeleteStackResourcesAsyncResult struct {
	result *DeleteStackResourcesResult
	err    error
}

func NewDeleteStackResourcesResultFromDict(data map[string]interface{}) DeleteStackResourcesResult {
	return DeleteStackResourcesResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStackResourcesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStackResourcesResult) Pointer() *DeleteStackResourcesResult {
	return &p
}

type DeleteStackEntityResult struct {
	Item *Stack `json:"item"`
}

type DeleteStackEntityAsyncResult struct {
	result *DeleteStackEntityResult
	err    error
}

func NewDeleteStackEntityResultFromDict(data map[string]interface{}) DeleteStackEntityResult {
	return DeleteStackEntityResult{
		Item: NewStackFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStackEntityResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStackEntityResult) Pointer() *DeleteStackEntityResult {
	return &p
}

type DescribeResourcesResult struct {
	Items         []Resource `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeResourcesAsyncResult struct {
	result *DescribeResourcesResult
	err    error
}

func NewDescribeResourcesResultFromDict(data map[string]interface{}) DescribeResourcesResult {
	return DescribeResourcesResult{
		Items:         CastResources(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeResourcesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastResourcesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeResourcesResult) Pointer() *DescribeResourcesResult {
	return &p
}

type GetResourceResult struct {
	Item *Resource `json:"item"`
}

type GetResourceAsyncResult struct {
	result *GetResourceResult
	err    error
}

func NewGetResourceResultFromDict(data map[string]interface{}) GetResourceResult {
	return GetResourceResult{
		Item: NewResourceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetResourceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetResourceResult) Pointer() *GetResourceResult {
	return &p
}

type DescribeEventsResult struct {
	Items         []Event `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

func NewDescribeEventsResultFromDict(data map[string]interface{}) DescribeEventsResult {
	return DescribeEventsResult{
		Items:         CastEvents(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEventsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEventsResult) Pointer() *DescribeEventsResult {
	return &p
}

type GetEventResult struct {
	Item *Event `json:"item"`
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

func NewGetEventResultFromDict(data map[string]interface{}) GetEventResult {
	return GetEventResult{
		Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEventResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetEventResult) Pointer() *GetEventResult {
	return &p
}

type DescribeOutputsResult struct {
	Items         []Output `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeOutputsAsyncResult struct {
	result *DescribeOutputsResult
	err    error
}

func NewDescribeOutputsResultFromDict(data map[string]interface{}) DescribeOutputsResult {
	return DescribeOutputsResult{
		Items:         CastOutputs(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeOutputsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastOutputsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeOutputsResult) Pointer() *DescribeOutputsResult {
	return &p
}

type GetOutputResult struct {
	Item *Output `json:"item"`
}

type GetOutputAsyncResult struct {
	result *GetOutputResult
	err    error
}

func NewGetOutputResultFromDict(data map[string]interface{}) GetOutputResult {
	return GetOutputResult{
		Item: NewOutputFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetOutputResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetOutputResult) Pointer() *GetOutputResult {
	return &p
}
