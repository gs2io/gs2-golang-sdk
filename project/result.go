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

package project

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type CreateAccountResult struct {
	Item *Account `json:"item"`
}

type CreateAccountAsyncResult struct {
	result *CreateAccountResult
	err    error
}

func NewCreateAccountResultFromJson(data string) CreateAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAccountResultFromDict(dict)
}

func NewCreateAccountResultFromDict(data map[string]interface{}) CreateAccountResult {
	return CreateAccountResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateAccountResult) Pointer() *CreateAccountResult {
	return &p
}

type VerifyResult struct {
	Item *Account `json:"item"`
}

type VerifyAsyncResult struct {
	result *VerifyResult
	err    error
}

func NewVerifyResultFromJson(data string) VerifyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyResultFromDict(dict)
}

func NewVerifyResultFromDict(data map[string]interface{}) VerifyResult {
	return VerifyResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VerifyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p VerifyResult) Pointer() *VerifyResult {
	return &p
}

type SignInResult struct {
	Item         *Account `json:"item"`
	AccountToken *string  `json:"accountToken"`
}

type SignInAsyncResult struct {
	result *SignInResult
	err    error
}

func NewSignInResultFromJson(data string) SignInResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSignInResultFromDict(dict)
}

func NewSignInResultFromDict(data map[string]interface{}) SignInResult {
	return SignInResult{
		Item:         NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
		AccountToken: core.CastString(data["accountToken"]),
	}
}

func (p SignInResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"accountToken": p.AccountToken,
	}
}

func (p SignInResult) Pointer() *SignInResult {
	return &p
}

type IssueAccountTokenResult struct {
	AccountToken *string `json:"accountToken"`
}

type IssueAccountTokenAsyncResult struct {
	result *IssueAccountTokenResult
	err    error
}

func NewIssueAccountTokenResultFromJson(data string) IssueAccountTokenResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueAccountTokenResultFromDict(dict)
}

func NewIssueAccountTokenResultFromDict(data map[string]interface{}) IssueAccountTokenResult {
	return IssueAccountTokenResult{
		AccountToken: core.CastString(data["accountToken"]),
	}
}

func (p IssueAccountTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
	}
}

func (p IssueAccountTokenResult) Pointer() *IssueAccountTokenResult {
	return &p
}

type ForgetResult struct {
	IssuePasswordToken *string `json:"issuePasswordToken"`
}

type ForgetAsyncResult struct {
	result *ForgetResult
	err    error
}

func NewForgetResultFromJson(data string) ForgetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForgetResultFromDict(dict)
}

func NewForgetResultFromDict(data map[string]interface{}) ForgetResult {
	return ForgetResult{
		IssuePasswordToken: core.CastString(data["issuePasswordToken"]),
	}
}

func (p ForgetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"issuePasswordToken": p.IssuePasswordToken,
	}
}

func (p ForgetResult) Pointer() *ForgetResult {
	return &p
}

type IssuePasswordResult struct {
	NewPassword *string `json:"newPassword"`
}

type IssuePasswordAsyncResult struct {
	result *IssuePasswordResult
	err    error
}

func NewIssuePasswordResultFromJson(data string) IssuePasswordResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssuePasswordResultFromDict(dict)
}

func NewIssuePasswordResultFromDict(data map[string]interface{}) IssuePasswordResult {
	return IssuePasswordResult{
		NewPassword: core.CastString(data["newPassword"]),
	}
}

func (p IssuePasswordResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newPassword": p.NewPassword,
	}
}

func (p IssuePasswordResult) Pointer() *IssuePasswordResult {
	return &p
}

type UpdateAccountResult struct {
	Item *Account `json:"item"`
}

type UpdateAccountAsyncResult struct {
	result *UpdateAccountResult
	err    error
}

func NewUpdateAccountResultFromJson(data string) UpdateAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateAccountResultFromDict(dict)
}

func NewUpdateAccountResultFromDict(data map[string]interface{}) UpdateAccountResult {
	return UpdateAccountResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateAccountResult) Pointer() *UpdateAccountResult {
	return &p
}

type DeleteAccountResult struct {
}

type DeleteAccountAsyncResult struct {
	result *DeleteAccountResult
	err    error
}

func NewDeleteAccountResultFromJson(data string) DeleteAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAccountResultFromDict(dict)
}

func NewDeleteAccountResultFromDict(data map[string]interface{}) DeleteAccountResult {
	return DeleteAccountResult{}
}

func (p DeleteAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteAccountResult) Pointer() *DeleteAccountResult {
	return &p
}

type DescribeProjectsResult struct {
	Items         []Project `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeProjectsAsyncResult struct {
	result *DescribeProjectsResult
	err    error
}

func NewDescribeProjectsResultFromJson(data string) DescribeProjectsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProjectsResultFromDict(dict)
}

func NewDescribeProjectsResultFromDict(data map[string]interface{}) DescribeProjectsResult {
	return DescribeProjectsResult{
		Items:         CastProjects(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeProjectsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProjectsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeProjectsResult) Pointer() *DescribeProjectsResult {
	return &p
}

type CreateProjectResult struct {
	Item *Project `json:"item"`
}

type CreateProjectAsyncResult struct {
	result *CreateProjectResult
	err    error
}

func NewCreateProjectResultFromJson(data string) CreateProjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProjectResultFromDict(dict)
}

func NewCreateProjectResultFromDict(data map[string]interface{}) CreateProjectResult {
	return CreateProjectResult{
		Item: NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateProjectResult) Pointer() *CreateProjectResult {
	return &p
}

type GetProjectResult struct {
	Item *Project `json:"item"`
}

type GetProjectAsyncResult struct {
	result *GetProjectResult
	err    error
}

func NewGetProjectResultFromJson(data string) GetProjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectResultFromDict(dict)
}

func NewGetProjectResultFromDict(data map[string]interface{}) GetProjectResult {
	return GetProjectResult{
		Item: NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetProjectResult) Pointer() *GetProjectResult {
	return &p
}

type GetProjectTokenResult struct {
	Item         *Project `json:"item"`
	OwnerId      *string  `json:"ownerId"`
	ProjectToken *string  `json:"projectToken"`
}

type GetProjectTokenAsyncResult struct {
	result *GetProjectTokenResult
	err    error
}

func NewGetProjectTokenResultFromJson(data string) GetProjectTokenResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectTokenResultFromDict(dict)
}

func NewGetProjectTokenResultFromDict(data map[string]interface{}) GetProjectTokenResult {
	return GetProjectTokenResult{
		Item:         NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
		OwnerId:      core.CastString(data["ownerId"]),
		ProjectToken: core.CastString(data["projectToken"]),
	}
}

func (p GetProjectTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"ownerId":      p.OwnerId,
		"projectToken": p.ProjectToken,
	}
}

func (p GetProjectTokenResult) Pointer() *GetProjectTokenResult {
	return &p
}

type GetProjectTokenByIdentifierResult struct {
	Item         *Project `json:"item"`
	OwnerId      *string  `json:"ownerId"`
	ProjectToken *string  `json:"projectToken"`
}

type GetProjectTokenByIdentifierAsyncResult struct {
	result *GetProjectTokenByIdentifierResult
	err    error
}

func NewGetProjectTokenByIdentifierResultFromJson(data string) GetProjectTokenByIdentifierResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectTokenByIdentifierResultFromDict(dict)
}

func NewGetProjectTokenByIdentifierResultFromDict(data map[string]interface{}) GetProjectTokenByIdentifierResult {
	return GetProjectTokenByIdentifierResult{
		Item:         NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
		OwnerId:      core.CastString(data["ownerId"]),
		ProjectToken: core.CastString(data["projectToken"]),
	}
}

func (p GetProjectTokenByIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"ownerId":      p.OwnerId,
		"projectToken": p.ProjectToken,
	}
}

func (p GetProjectTokenByIdentifierResult) Pointer() *GetProjectTokenByIdentifierResult {
	return &p
}

type UpdateProjectResult struct {
	Item *Project `json:"item"`
}

type UpdateProjectAsyncResult struct {
	result *UpdateProjectResult
	err    error
}

func NewUpdateProjectResultFromJson(data string) UpdateProjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateProjectResultFromDict(dict)
}

func NewUpdateProjectResultFromDict(data map[string]interface{}) UpdateProjectResult {
	return UpdateProjectResult{
		Item: NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateProjectResult) Pointer() *UpdateProjectResult {
	return &p
}

type DeleteProjectResult struct {
	Item *Project `json:"item"`
}

type DeleteProjectAsyncResult struct {
	result *DeleteProjectResult
	err    error
}

func NewDeleteProjectResultFromJson(data string) DeleteProjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProjectResultFromDict(dict)
}

func NewDeleteProjectResultFromDict(data map[string]interface{}) DeleteProjectResult {
	return DeleteProjectResult{
		Item: NewProjectFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteProjectResult) Pointer() *DeleteProjectResult {
	return &p
}

type DescribeBillingMethodsResult struct {
	Items         []BillingMethod `json:"items"`
	NextPageToken *string         `json:"nextPageToken"`
}

type DescribeBillingMethodsAsyncResult struct {
	result *DescribeBillingMethodsResult
	err    error
}

func NewDescribeBillingMethodsResultFromJson(data string) DescribeBillingMethodsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBillingMethodsResultFromDict(dict)
}

func NewDescribeBillingMethodsResultFromDict(data map[string]interface{}) DescribeBillingMethodsResult {
	return DescribeBillingMethodsResult{
		Items:         CastBillingMethods(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeBillingMethodsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBillingMethodsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeBillingMethodsResult) Pointer() *DescribeBillingMethodsResult {
	return &p
}

type CreateBillingMethodResult struct {
	Item *BillingMethod `json:"item"`
}

type CreateBillingMethodAsyncResult struct {
	result *CreateBillingMethodResult
	err    error
}

func NewCreateBillingMethodResultFromJson(data string) CreateBillingMethodResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBillingMethodResultFromDict(dict)
}

func NewCreateBillingMethodResultFromDict(data map[string]interface{}) CreateBillingMethodResult {
	return CreateBillingMethodResult{
		Item: NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateBillingMethodResult) Pointer() *CreateBillingMethodResult {
	return &p
}

type GetBillingMethodResult struct {
	Item *BillingMethod `json:"item"`
}

type GetBillingMethodAsyncResult struct {
	result *GetBillingMethodResult
	err    error
}

func NewGetBillingMethodResultFromJson(data string) GetBillingMethodResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBillingMethodResultFromDict(dict)
}

func NewGetBillingMethodResultFromDict(data map[string]interface{}) GetBillingMethodResult {
	return GetBillingMethodResult{
		Item: NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetBillingMethodResult) Pointer() *GetBillingMethodResult {
	return &p
}

type UpdateBillingMethodResult struct {
	Item *BillingMethod `json:"item"`
}

type UpdateBillingMethodAsyncResult struct {
	result *UpdateBillingMethodResult
	err    error
}

func NewUpdateBillingMethodResultFromJson(data string) UpdateBillingMethodResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBillingMethodResultFromDict(dict)
}

func NewUpdateBillingMethodResultFromDict(data map[string]interface{}) UpdateBillingMethodResult {
	return UpdateBillingMethodResult{
		Item: NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateBillingMethodResult) Pointer() *UpdateBillingMethodResult {
	return &p
}

type DeleteBillingMethodResult struct {
	Item *BillingMethod `json:"item"`
}

type DeleteBillingMethodAsyncResult struct {
	result *DeleteBillingMethodResult
	err    error
}

func NewDeleteBillingMethodResultFromJson(data string) DeleteBillingMethodResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBillingMethodResultFromDict(dict)
}

func NewDeleteBillingMethodResultFromDict(data map[string]interface{}) DeleteBillingMethodResult {
	return DeleteBillingMethodResult{
		Item: NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteBillingMethodResult) Pointer() *DeleteBillingMethodResult {
	return &p
}

type DescribeReceiptsResult struct {
	Items         []Receipt `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeReceiptsAsyncResult struct {
	result *DescribeReceiptsResult
	err    error
}

func NewDescribeReceiptsResultFromJson(data string) DescribeReceiptsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiptsResultFromDict(dict)
}

func NewDescribeReceiptsResultFromDict(data map[string]interface{}) DescribeReceiptsResult {
	return DescribeReceiptsResult{
		Items:         CastReceipts(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeReceiptsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiptsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeReceiptsResult) Pointer() *DescribeReceiptsResult {
	return &p
}

type DescribeBillingsResult struct {
	Items []Billing `json:"items"`
}

type DescribeBillingsAsyncResult struct {
	result *DescribeBillingsResult
	err    error
}

func NewDescribeBillingsResultFromJson(data string) DescribeBillingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBillingsResultFromDict(dict)
}

func NewDescribeBillingsResultFromDict(data map[string]interface{}) DescribeBillingsResult {
	return DescribeBillingsResult{
		Items: CastBillings(core.CastArray(data["items"])),
	}
}

func (p DescribeBillingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBillingsFromDict(
			p.Items,
		),
	}
}

func (p DescribeBillingsResult) Pointer() *DescribeBillingsResult {
	return &p
}
