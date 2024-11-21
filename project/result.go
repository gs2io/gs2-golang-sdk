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
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateAccountResult) Pointer() *CreateAccountResult {
	return &p
}

type VerifyResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyResult) Pointer() *VerifyResult {
	return &p
}

type SignInResult struct {
	Item         *Account             `json:"item"`
	AccountToken *string              `json:"accountToken"`
	Metadata     *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p SignInResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"accountToken": p.AccountToken,
	}
}

func (p SignInResult) Pointer() *SignInResult {
	return &p
}

type ForgetResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
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
	return ForgetResult{}
}

func (p ForgetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ForgetResult) Pointer() *ForgetResult {
	return &p
}

type IssuePasswordResult struct {
	NewPassword *string              `json:"newPassword"`
	Metadata    *core.ResultMetadata `json:"metadata"`
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
		NewPassword: func() *string {
			v, ok := data["newPassword"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newPassword"])
		}(),
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
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateAccountResult) Pointer() *UpdateAccountResult {
	return &p
}

type EnableMfaResult struct {
	Item           *Account             `json:"item"`
	ChallengeToken *string              `json:"challengeToken"`
	Metadata       *core.ResultMetadata `json:"metadata"`
}

type EnableMfaAsyncResult struct {
	result *EnableMfaResult
	err    error
}

func NewEnableMfaResultFromJson(data string) EnableMfaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEnableMfaResultFromDict(dict)
}

func NewEnableMfaResultFromDict(data map[string]interface{}) EnableMfaResult {
	return EnableMfaResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ChallengeToken: func() *string {
			v, ok := data["challengeToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["challengeToken"])
		}(),
	}
}

func (p EnableMfaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"challengeToken": p.ChallengeToken,
	}
}

func (p EnableMfaResult) Pointer() *EnableMfaResult {
	return &p
}

type ChallengeMfaResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ChallengeMfaAsyncResult struct {
	result *ChallengeMfaResult
	err    error
}

func NewChallengeMfaResultFromJson(data string) ChallengeMfaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewChallengeMfaResultFromDict(dict)
}

func NewChallengeMfaResultFromDict(data map[string]interface{}) ChallengeMfaResult {
	return ChallengeMfaResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ChallengeMfaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ChallengeMfaResult) Pointer() *ChallengeMfaResult {
	return &p
}

type DisableMfaResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DisableMfaAsyncResult struct {
	result *DisableMfaResult
	err    error
}

func NewDisableMfaResultFromJson(data string) DisableMfaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDisableMfaResultFromDict(dict)
}

func NewDisableMfaResultFromDict(data map[string]interface{}) DisableMfaResult {
	return DisableMfaResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DisableMfaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DisableMfaResult) Pointer() *DisableMfaResult {
	return &p
}

type DeleteAccountResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
	return DeleteAccountResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteAccountResult) Pointer() *DeleteAccountResult {
	return &p
}

type DescribeProjectsResult struct {
	Items         []Project            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Project {
			if data["items"] == nil {
				return nil
			}
			return CastProjects(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateProjectResult) Pointer() *CreateProjectResult {
	return &p
}

type GetProjectResult struct {
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetProjectResult) Pointer() *GetProjectResult {
	return &p
}

type GetProjectTokenResult struct {
	Item         *Project             `json:"item"`
	OwnerId      *string              `json:"ownerId"`
	ProjectToken *string              `json:"projectToken"`
	Metadata     *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		OwnerId: func() *string {
			v, ok := data["ownerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ownerId"])
		}(),
		ProjectToken: func() *string {
			v, ok := data["projectToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectToken"])
		}(),
	}
}

func (p GetProjectTokenResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"ownerId":      p.OwnerId,
		"projectToken": p.ProjectToken,
	}
}

func (p GetProjectTokenResult) Pointer() *GetProjectTokenResult {
	return &p
}

type GetProjectTokenByIdentifierResult struct {
	Item         *Project             `json:"item"`
	OwnerId      *string              `json:"ownerId"`
	ProjectToken *string              `json:"projectToken"`
	Metadata     *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		OwnerId: func() *string {
			v, ok := data["ownerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ownerId"])
		}(),
		ProjectToken: func() *string {
			v, ok := data["projectToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectToken"])
		}(),
	}
}

func (p GetProjectTokenByIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"ownerId":      p.OwnerId,
		"projectToken": p.ProjectToken,
	}
}

func (p GetProjectTokenByIdentifierResult) Pointer() *GetProjectTokenByIdentifierResult {
	return &p
}

type UpdateProjectResult struct {
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateProjectResult) Pointer() *UpdateProjectResult {
	return &p
}

type ActivateRegionResult struct {
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ActivateRegionAsyncResult struct {
	result *ActivateRegionResult
	err    error
}

func NewActivateRegionResultFromJson(data string) ActivateRegionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewActivateRegionResultFromDict(dict)
}

func NewActivateRegionResultFromDict(data map[string]interface{}) ActivateRegionResult {
	return ActivateRegionResult{
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ActivateRegionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ActivateRegionResult) Pointer() *ActivateRegionResult {
	return &p
}

type WaitActivateRegionResult struct {
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WaitActivateRegionAsyncResult struct {
	result *WaitActivateRegionResult
	err    error
}

func NewWaitActivateRegionResultFromJson(data string) WaitActivateRegionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitActivateRegionResultFromDict(dict)
}

func NewWaitActivateRegionResultFromDict(data map[string]interface{}) WaitActivateRegionResult {
	return WaitActivateRegionResult{
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p WaitActivateRegionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p WaitActivateRegionResult) Pointer() *WaitActivateRegionResult {
	return &p
}

type DeleteProjectResult struct {
	Item     *Project             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Project {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteProjectResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteProjectResult) Pointer() *DeleteProjectResult {
	return &p
}

type DescribeBillingMethodsResult struct {
	Items         []BillingMethod      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []BillingMethod {
			if data["items"] == nil {
				return nil
			}
			return CastBillingMethods(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *BillingMethod       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *BillingMethod {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateBillingMethodResult) Pointer() *CreateBillingMethodResult {
	return &p
}

type GetBillingMethodResult struct {
	Item     *BillingMethod       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *BillingMethod {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetBillingMethodResult) Pointer() *GetBillingMethodResult {
	return &p
}

type UpdateBillingMethodResult struct {
	Item     *BillingMethod       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *BillingMethod {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateBillingMethodResult) Pointer() *UpdateBillingMethodResult {
	return &p
}

type DeleteBillingMethodResult struct {
	Item     *BillingMethod       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *BillingMethod {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBillingMethodFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteBillingMethodResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteBillingMethodResult) Pointer() *DeleteBillingMethodResult {
	return &p
}

type DescribeReceiptsResult struct {
	Items         []Receipt            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Receipt {
			if data["items"] == nil {
				return nil
			}
			return CastReceipts(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Items    []Billing            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Billing {
			if data["items"] == nil {
				return nil
			}
			return CastBillings(core.CastArray(data["items"]))
		}(),
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

type DescribeDumpProgressesResult struct {
	Items         []DumpProgress       `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeDumpProgressesAsyncResult struct {
	result *DescribeDumpProgressesResult
	err    error
}

func NewDescribeDumpProgressesResultFromJson(data string) DescribeDumpProgressesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDumpProgressesResultFromDict(dict)
}

func NewDescribeDumpProgressesResultFromDict(data map[string]interface{}) DescribeDumpProgressesResult {
	return DescribeDumpProgressesResult{
		Items: func() []DumpProgress {
			if data["items"] == nil {
				return nil
			}
			return CastDumpProgresses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeDumpProgressesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDumpProgressesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDumpProgressesResult) Pointer() *DescribeDumpProgressesResult {
	return &p
}

type GetDumpProgressResult struct {
	Item     *DumpProgress        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDumpProgressAsyncResult struct {
	result *GetDumpProgressResult
	err    error
}

func NewGetDumpProgressResultFromJson(data string) GetDumpProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDumpProgressResultFromDict(dict)
}

func NewGetDumpProgressResultFromDict(data map[string]interface{}) GetDumpProgressResult {
	return GetDumpProgressResult{
		Item: func() *DumpProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDumpProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetDumpProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetDumpProgressResult) Pointer() *GetDumpProgressResult {
	return &p
}

type WaitDumpUserDataResult struct {
	Item     *DumpProgress        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WaitDumpUserDataAsyncResult struct {
	result *WaitDumpUserDataResult
	err    error
}

func NewWaitDumpUserDataResultFromJson(data string) WaitDumpUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitDumpUserDataResultFromDict(dict)
}

func NewWaitDumpUserDataResultFromDict(data map[string]interface{}) WaitDumpUserDataResult {
	return WaitDumpUserDataResult{
		Item: func() *DumpProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDumpProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p WaitDumpUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p WaitDumpUserDataResult) Pointer() *WaitDumpUserDataResult {
	return &p
}

type ArchiveDumpUserDataResult struct {
	Item     *DumpProgress        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ArchiveDumpUserDataAsyncResult struct {
	result *ArchiveDumpUserDataResult
	err    error
}

func NewArchiveDumpUserDataResultFromJson(data string) ArchiveDumpUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewArchiveDumpUserDataResultFromDict(dict)
}

func NewArchiveDumpUserDataResultFromDict(data map[string]interface{}) ArchiveDumpUserDataResult {
	return ArchiveDumpUserDataResult{
		Item: func() *DumpProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDumpProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ArchiveDumpUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ArchiveDumpUserDataResult) Pointer() *ArchiveDumpUserDataResult {
	return &p
}

type DumpUserDataResult struct {
	Item     *DumpProgress        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DumpUserDataAsyncResult struct {
	result *DumpUserDataResult
	err    error
}

func NewDumpUserDataResultFromJson(data string) DumpUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataResultFromDict(dict)
}

func NewDumpUserDataResultFromDict(data map[string]interface{}) DumpUserDataResult {
	return DumpUserDataResult{
		Item: func() *DumpProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDumpProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DumpUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DumpUserDataResult) Pointer() *DumpUserDataResult {
	return &p
}

type GetDumpUserDataResult struct {
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDumpUserDataAsyncResult struct {
	result *GetDumpUserDataResult
	err    error
}

func NewGetDumpUserDataResultFromJson(data string) GetDumpUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDumpUserDataResultFromDict(dict)
}

func NewGetDumpUserDataResultFromDict(data map[string]interface{}) GetDumpUserDataResult {
	return GetDumpUserDataResult{
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
	}
}

func (p GetDumpUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p GetDumpUserDataResult) Pointer() *GetDumpUserDataResult {
	return &p
}

type DescribeCleanProgressesResult struct {
	Items         []CleanProgress      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeCleanProgressesAsyncResult struct {
	result *DescribeCleanProgressesResult
	err    error
}

func NewDescribeCleanProgressesResultFromJson(data string) DescribeCleanProgressesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCleanProgressesResultFromDict(dict)
}

func NewDescribeCleanProgressesResultFromDict(data map[string]interface{}) DescribeCleanProgressesResult {
	return DescribeCleanProgressesResult{
		Items: func() []CleanProgress {
			if data["items"] == nil {
				return nil
			}
			return CastCleanProgresses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeCleanProgressesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCleanProgressesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCleanProgressesResult) Pointer() *DescribeCleanProgressesResult {
	return &p
}

type GetCleanProgressResult struct {
	Item     *CleanProgress       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCleanProgressAsyncResult struct {
	result *GetCleanProgressResult
	err    error
}

func NewGetCleanProgressResultFromJson(data string) GetCleanProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCleanProgressResultFromDict(dict)
}

func NewGetCleanProgressResultFromDict(data map[string]interface{}) GetCleanProgressResult {
	return GetCleanProgressResult{
		Item: func() *CleanProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCleanProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCleanProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCleanProgressResult) Pointer() *GetCleanProgressResult {
	return &p
}

type WaitCleanUserDataResult struct {
	Item     *CleanProgress       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WaitCleanUserDataAsyncResult struct {
	result *WaitCleanUserDataResult
	err    error
}

func NewWaitCleanUserDataResultFromJson(data string) WaitCleanUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitCleanUserDataResultFromDict(dict)
}

func NewWaitCleanUserDataResultFromDict(data map[string]interface{}) WaitCleanUserDataResult {
	return WaitCleanUserDataResult{
		Item: func() *CleanProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCleanProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p WaitCleanUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p WaitCleanUserDataResult) Pointer() *WaitCleanUserDataResult {
	return &p
}

type CleanUserDataResult struct {
	Item     *CleanProgress       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CleanUserDataAsyncResult struct {
	result *CleanUserDataResult
	err    error
}

func NewCleanUserDataResultFromJson(data string) CleanUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataResultFromDict(dict)
}

func NewCleanUserDataResultFromDict(data map[string]interface{}) CleanUserDataResult {
	return CleanUserDataResult{
		Item: func() *CleanProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCleanProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CleanUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CleanUserDataResult) Pointer() *CleanUserDataResult {
	return &p
}

type DescribeImportProgressesResult struct {
	Items         []ImportProgress     `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeImportProgressesAsyncResult struct {
	result *DescribeImportProgressesResult
	err    error
}

func NewDescribeImportProgressesResultFromJson(data string) DescribeImportProgressesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeImportProgressesResultFromDict(dict)
}

func NewDescribeImportProgressesResultFromDict(data map[string]interface{}) DescribeImportProgressesResult {
	return DescribeImportProgressesResult{
		Items: func() []ImportProgress {
			if data["items"] == nil {
				return nil
			}
			return CastImportProgresses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeImportProgressesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastImportProgressesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeImportProgressesResult) Pointer() *DescribeImportProgressesResult {
	return &p
}

type GetImportProgressResult struct {
	Item     *ImportProgress      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetImportProgressAsyncResult struct {
	result *GetImportProgressResult
	err    error
}

func NewGetImportProgressResultFromJson(data string) GetImportProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetImportProgressResultFromDict(dict)
}

func NewGetImportProgressResultFromDict(data map[string]interface{}) GetImportProgressResult {
	return GetImportProgressResult{
		Item: func() *ImportProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewImportProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetImportProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetImportProgressResult) Pointer() *GetImportProgressResult {
	return &p
}

type WaitImportUserDataResult struct {
	Item     *ImportProgress      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WaitImportUserDataAsyncResult struct {
	result *WaitImportUserDataResult
	err    error
}

func NewWaitImportUserDataResultFromJson(data string) WaitImportUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitImportUserDataResultFromDict(dict)
}

func NewWaitImportUserDataResultFromDict(data map[string]interface{}) WaitImportUserDataResult {
	return WaitImportUserDataResult{
		Item: func() *ImportProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewImportProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p WaitImportUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p WaitImportUserDataResult) Pointer() *WaitImportUserDataResult {
	return &p
}

type PrepareImportUserDataResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PrepareImportUserDataAsyncResult struct {
	result *PrepareImportUserDataResult
	err    error
}

func NewPrepareImportUserDataResultFromJson(data string) PrepareImportUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataResultFromDict(dict)
}

func NewPrepareImportUserDataResultFromDict(data map[string]interface{}) PrepareImportUserDataResult {
	return PrepareImportUserDataResult{
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		UploadUrl: func() *string {
			v, ok := data["uploadUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadUrl"])
		}(),
	}
}

func (p PrepareImportUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
	}
}

func (p PrepareImportUserDataResult) Pointer() *PrepareImportUserDataResult {
	return &p
}

type ImportUserDataResult struct {
	Item     *ImportProgress      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ImportUserDataAsyncResult struct {
	result *ImportUserDataResult
	err    error
}

func NewImportUserDataResultFromJson(data string) ImportUserDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataResultFromDict(dict)
}

func NewImportUserDataResultFromDict(data map[string]interface{}) ImportUserDataResult {
	return ImportUserDataResult{
		Item: func() *ImportProgress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewImportProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ImportUserDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ImportUserDataResult) Pointer() *ImportUserDataResult {
	return &p
}

type DescribeImportErrorLogsResult struct {
	Items         []ImportErrorLog     `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeImportErrorLogsAsyncResult struct {
	result *DescribeImportErrorLogsResult
	err    error
}

func NewDescribeImportErrorLogsResultFromJson(data string) DescribeImportErrorLogsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeImportErrorLogsResultFromDict(dict)
}

func NewDescribeImportErrorLogsResultFromDict(data map[string]interface{}) DescribeImportErrorLogsResult {
	return DescribeImportErrorLogsResult{
		Items: func() []ImportErrorLog {
			if data["items"] == nil {
				return nil
			}
			return CastImportErrorLogs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeImportErrorLogsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastImportErrorLogsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeImportErrorLogsResult) Pointer() *DescribeImportErrorLogsResult {
	return &p
}

type GetImportErrorLogResult struct {
	Item     *ImportErrorLog      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetImportErrorLogAsyncResult struct {
	result *GetImportErrorLogResult
	err    error
}

func NewGetImportErrorLogResultFromJson(data string) GetImportErrorLogResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetImportErrorLogResultFromDict(dict)
}

func NewGetImportErrorLogResultFromDict(data map[string]interface{}) GetImportErrorLogResult {
	return GetImportErrorLogResult{
		Item: func() *ImportErrorLog {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewImportErrorLogFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetImportErrorLogResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetImportErrorLogResult) Pointer() *GetImportErrorLogResult {
	return &p
}
