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

type CreateAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	FullName        *string `json:"fullName"`
	CompanyName     *string `json:"companyName"`
	Password        *string `json:"password"`
	Lang            *string `json:"lang"`
}

func NewCreateAccountRequestFromJson(data string) CreateAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAccountRequestFromDict(dict)
}

func NewCreateAccountRequestFromDict(data map[string]interface{}) CreateAccountRequest {
	return CreateAccountRequest{
		Email:       core.CastString(data["email"]),
		FullName:    core.CastString(data["fullName"]),
		CompanyName: core.CastString(data["companyName"]),
		Password:    core.CastString(data["password"]),
		Lang:        core.CastString(data["lang"]),
	}
}

func (p CreateAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":       p.Email,
		"fullName":    p.FullName,
		"companyName": p.CompanyName,
		"password":    p.Password,
		"lang":        p.Lang,
	}
}

func (p CreateAccountRequest) Pointer() *CreateAccountRequest {
	return &p
}

type VerifyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	VerifyToken     *string `json:"verifyToken"`
}

func NewVerifyRequestFromJson(data string) VerifyRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRequestFromDict(dict)
}

func NewVerifyRequestFromDict(data map[string]interface{}) VerifyRequest {
	return VerifyRequest{
		VerifyToken: core.CastString(data["verifyToken"]),
	}
}

func (p VerifyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"verifyToken": p.VerifyToken,
	}
}

func (p VerifyRequest) Pointer() *VerifyRequest {
	return &p
}

type SignInRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	Password        *string `json:"password"`
}

func NewSignInRequestFromJson(data string) SignInRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSignInRequestFromDict(dict)
}

func NewSignInRequestFromDict(data map[string]interface{}) SignInRequest {
	return SignInRequest{
		Email:    core.CastString(data["email"]),
		Password: core.CastString(data["password"]),
	}
}

func (p SignInRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":    p.Email,
		"password": p.Password,
	}
}

func (p SignInRequest) Pointer() *SignInRequest {
	return &p
}

type ForgetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	Lang            *string `json:"lang"`
}

func NewForgetRequestFromJson(data string) ForgetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForgetRequestFromDict(dict)
}

func NewForgetRequestFromDict(data map[string]interface{}) ForgetRequest {
	return ForgetRequest{
		Email: core.CastString(data["email"]),
		Lang:  core.CastString(data["lang"]),
	}
}

func (p ForgetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email": p.Email,
		"lang":  p.Lang,
	}
}

func (p ForgetRequest) Pointer() *ForgetRequest {
	return &p
}

type IssuePasswordRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	IssuePasswordToken *string `json:"issuePasswordToken"`
}

func NewIssuePasswordRequestFromJson(data string) IssuePasswordRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssuePasswordRequestFromDict(dict)
}

func NewIssuePasswordRequestFromDict(data map[string]interface{}) IssuePasswordRequest {
	return IssuePasswordRequest{
		IssuePasswordToken: core.CastString(data["issuePasswordToken"]),
	}
}

func (p IssuePasswordRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"issuePasswordToken": p.IssuePasswordToken,
	}
}

func (p IssuePasswordRequest) Pointer() *IssuePasswordRequest {
	return &p
}

type UpdateAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	FullName        *string `json:"fullName"`
	CompanyName     *string `json:"companyName"`
	Password        *string `json:"password"`
	AccountToken    *string `json:"accountToken"`
}

func NewUpdateAccountRequestFromJson(data string) UpdateAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateAccountRequestFromDict(dict)
}

func NewUpdateAccountRequestFromDict(data map[string]interface{}) UpdateAccountRequest {
	return UpdateAccountRequest{
		Email:        core.CastString(data["email"]),
		FullName:     core.CastString(data["fullName"]),
		CompanyName:  core.CastString(data["companyName"]),
		Password:     core.CastString(data["password"]),
		AccountToken: core.CastString(data["accountToken"]),
	}
}

func (p UpdateAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":        p.Email,
		"fullName":     p.FullName,
		"companyName":  p.CompanyName,
		"password":     p.Password,
		"accountToken": p.AccountToken,
	}
}

func (p UpdateAccountRequest) Pointer() *UpdateAccountRequest {
	return &p
}

type DeleteAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
}

func NewDeleteAccountRequestFromJson(data string) DeleteAccountRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAccountRequestFromDict(dict)
}

func NewDeleteAccountRequestFromDict(data map[string]interface{}) DeleteAccountRequest {
	return DeleteAccountRequest{
		AccountToken: core.CastString(data["accountToken"]),
	}
}

func (p DeleteAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
	}
}

func (p DeleteAccountRequest) Pointer() *DeleteAccountRequest {
	return &p
}

type DescribeProjectsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeProjectsRequestFromJson(data string) DescribeProjectsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProjectsRequestFromDict(dict)
}

func NewDescribeProjectsRequestFromDict(data map[string]interface{}) DescribeProjectsRequest {
	return DescribeProjectsRequest{
		AccountToken: core.CastString(data["accountToken"]),
		PageToken:    core.CastString(data["pageToken"]),
		Limit:        core.CastInt32(data["limit"]),
	}
}

func (p DescribeProjectsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeProjectsRequest) Pointer() *DescribeProjectsRequest {
	return &p
}

type CreateProjectRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	AccountToken            *string `json:"accountToken"`
	Name                    *string `json:"name"`
	Description             *string `json:"description"`
	Plan                    *string `json:"plan"`
	Currency                *string `json:"currency"`
	ActivateRegionName      *string `json:"activateRegionName"`
	BillingMethodName       *string `json:"billingMethodName"`
	EnableEventBridge       *string `json:"enableEventBridge"`
	EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion    *string `json:"eventBridgeAwsRegion"`
}

func NewCreateProjectRequestFromJson(data string) CreateProjectRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProjectRequestFromDict(dict)
}

func NewCreateProjectRequestFromDict(data map[string]interface{}) CreateProjectRequest {
	return CreateProjectRequest{
		AccountToken:            core.CastString(data["accountToken"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Plan:                    core.CastString(data["plan"]),
		Currency:                core.CastString(data["currency"]),
		ActivateRegionName:      core.CastString(data["activateRegionName"]),
		BillingMethodName:       core.CastString(data["billingMethodName"]),
		EnableEventBridge:       core.CastString(data["enableEventBridge"]),
		EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
		EventBridgeAwsRegion:    core.CastString(data["eventBridgeAwsRegion"]),
	}
}

func (p CreateProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":            p.AccountToken,
		"name":                    p.Name,
		"description":             p.Description,
		"plan":                    p.Plan,
		"currency":                p.Currency,
		"activateRegionName":      p.ActivateRegionName,
		"billingMethodName":       p.BillingMethodName,
		"enableEventBridge":       p.EnableEventBridge,
		"eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
		"eventBridgeAwsRegion":    p.EventBridgeAwsRegion,
	}
}

func (p CreateProjectRequest) Pointer() *CreateProjectRequest {
	return &p
}

type GetProjectRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
}

func NewGetProjectRequestFromJson(data string) GetProjectRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectRequestFromDict(dict)
}

func NewGetProjectRequestFromDict(data map[string]interface{}) GetProjectRequest {
	return GetProjectRequest{
		AccountToken: core.CastString(data["accountToken"]),
		ProjectName:  core.CastString(data["projectName"]),
	}
}

func (p GetProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
	}
}

func (p GetProjectRequest) Pointer() *GetProjectRequest {
	return &p
}

type GetProjectTokenRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ProjectName     *string `json:"projectName"`
	AccountToken    *string `json:"accountToken"`
}

func NewGetProjectTokenRequestFromJson(data string) GetProjectTokenRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectTokenRequestFromDict(dict)
}

func NewGetProjectTokenRequestFromDict(data map[string]interface{}) GetProjectTokenRequest {
	return GetProjectTokenRequest{
		ProjectName:  core.CastString(data["projectName"]),
		AccountToken: core.CastString(data["accountToken"]),
	}
}

func (p GetProjectTokenRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectName":  p.ProjectName,
		"accountToken": p.AccountToken,
	}
}

func (p GetProjectTokenRequest) Pointer() *GetProjectTokenRequest {
	return &p
}

type GetProjectTokenByIdentifierRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountName     *string `json:"accountName"`
	ProjectName     *string `json:"projectName"`
	UserName        *string `json:"userName"`
	Password        *string `json:"password"`
}

func NewGetProjectTokenByIdentifierRequestFromJson(data string) GetProjectTokenByIdentifierRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProjectTokenByIdentifierRequestFromDict(dict)
}

func NewGetProjectTokenByIdentifierRequestFromDict(data map[string]interface{}) GetProjectTokenByIdentifierRequest {
	return GetProjectTokenByIdentifierRequest{
		AccountName: core.CastString(data["accountName"]),
		ProjectName: core.CastString(data["projectName"]),
		UserName:    core.CastString(data["userName"]),
		Password:    core.CastString(data["password"]),
	}
}

func (p GetProjectTokenByIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountName": p.AccountName,
		"projectName": p.ProjectName,
		"userName":    p.UserName,
		"password":    p.Password,
	}
}

func (p GetProjectTokenByIdentifierRequest) Pointer() *GetProjectTokenByIdentifierRequest {
	return &p
}

type UpdateProjectRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	AccountToken            *string `json:"accountToken"`
	ProjectName             *string `json:"projectName"`
	Description             *string `json:"description"`
	Plan                    *string `json:"plan"`
	BillingMethodName       *string `json:"billingMethodName"`
	EnableEventBridge       *string `json:"enableEventBridge"`
	EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion    *string `json:"eventBridgeAwsRegion"`
}

func NewUpdateProjectRequestFromJson(data string) UpdateProjectRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateProjectRequestFromDict(dict)
}

func NewUpdateProjectRequestFromDict(data map[string]interface{}) UpdateProjectRequest {
	return UpdateProjectRequest{
		AccountToken:            core.CastString(data["accountToken"]),
		ProjectName:             core.CastString(data["projectName"]),
		Description:             core.CastString(data["description"]),
		Plan:                    core.CastString(data["plan"]),
		BillingMethodName:       core.CastString(data["billingMethodName"]),
		EnableEventBridge:       core.CastString(data["enableEventBridge"]),
		EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
		EventBridgeAwsRegion:    core.CastString(data["eventBridgeAwsRegion"]),
	}
}

func (p UpdateProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":            p.AccountToken,
		"projectName":             p.ProjectName,
		"description":             p.Description,
		"plan":                    p.Plan,
		"billingMethodName":       p.BillingMethodName,
		"enableEventBridge":       p.EnableEventBridge,
		"eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
		"eventBridgeAwsRegion":    p.EventBridgeAwsRegion,
	}
}

func (p UpdateProjectRequest) Pointer() *UpdateProjectRequest {
	return &p
}

type ActivateRegionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	RegionName      *string `json:"regionName"`
}

func NewActivateRegionRequestFromJson(data string) ActivateRegionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewActivateRegionRequestFromDict(dict)
}

func NewActivateRegionRequestFromDict(data map[string]interface{}) ActivateRegionRequest {
	return ActivateRegionRequest{
		AccountToken: core.CastString(data["accountToken"]),
		ProjectName:  core.CastString(data["projectName"]),
		RegionName:   core.CastString(data["regionName"]),
	}
}

func (p ActivateRegionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
		"regionName":   p.RegionName,
	}
}

func (p ActivateRegionRequest) Pointer() *ActivateRegionRequest {
	return &p
}

type WaitActivateRegionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ProjectName     *string `json:"projectName"`
	RegionName      *string `json:"regionName"`
}

func NewWaitActivateRegionRequestFromJson(data string) WaitActivateRegionRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitActivateRegionRequestFromDict(dict)
}

func NewWaitActivateRegionRequestFromDict(data map[string]interface{}) WaitActivateRegionRequest {
	return WaitActivateRegionRequest{
		ProjectName: core.CastString(data["projectName"]),
		RegionName:  core.CastString(data["regionName"]),
	}
}

func (p WaitActivateRegionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectName": p.ProjectName,
		"regionName":  p.RegionName,
	}
}

func (p WaitActivateRegionRequest) Pointer() *WaitActivateRegionRequest {
	return &p
}

type DeleteProjectRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
}

func NewDeleteProjectRequestFromJson(data string) DeleteProjectRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProjectRequestFromDict(dict)
}

func NewDeleteProjectRequestFromDict(data map[string]interface{}) DeleteProjectRequest {
	return DeleteProjectRequest{
		AccountToken: core.CastString(data["accountToken"]),
		ProjectName:  core.CastString(data["projectName"]),
	}
}

func (p DeleteProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
	}
}

func (p DeleteProjectRequest) Pointer() *DeleteProjectRequest {
	return &p
}

type DescribeBillingMethodsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeBillingMethodsRequestFromJson(data string) DescribeBillingMethodsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBillingMethodsRequestFromDict(dict)
}

func NewDescribeBillingMethodsRequestFromDict(data map[string]interface{}) DescribeBillingMethodsRequest {
	return DescribeBillingMethodsRequest{
		AccountToken: core.CastString(data["accountToken"]),
		PageToken:    core.CastString(data["pageToken"]),
		Limit:        core.CastInt32(data["limit"]),
	}
}

func (p DescribeBillingMethodsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeBillingMethodsRequest) Pointer() *DescribeBillingMethodsRequest {
	return &p
}

type CreateBillingMethodRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	Description     *string `json:"description"`
	MethodType      *string `json:"methodType"`
	CardCustomerId  *string `json:"cardCustomerId"`
	PartnerId       *string `json:"partnerId"`
}

func NewCreateBillingMethodRequestFromJson(data string) CreateBillingMethodRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBillingMethodRequestFromDict(dict)
}

func NewCreateBillingMethodRequestFromDict(data map[string]interface{}) CreateBillingMethodRequest {
	return CreateBillingMethodRequest{
		AccountToken:   core.CastString(data["accountToken"]),
		Description:    core.CastString(data["description"]),
		MethodType:     core.CastString(data["methodType"]),
		CardCustomerId: core.CastString(data["cardCustomerId"]),
		PartnerId:      core.CastString(data["partnerId"]),
	}
}

func (p CreateBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":   p.AccountToken,
		"description":    p.Description,
		"methodType":     p.MethodType,
		"cardCustomerId": p.CardCustomerId,
		"partnerId":      p.PartnerId,
	}
}

func (p CreateBillingMethodRequest) Pointer() *CreateBillingMethodRequest {
	return &p
}

type GetBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
}

func NewGetBillingMethodRequestFromJson(data string) GetBillingMethodRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBillingMethodRequestFromDict(dict)
}

func NewGetBillingMethodRequestFromDict(data map[string]interface{}) GetBillingMethodRequest {
	return GetBillingMethodRequest{
		AccountToken:      core.CastString(data["accountToken"]),
		BillingMethodName: core.CastString(data["billingMethodName"]),
	}
}

func (p GetBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
	}
}

func (p GetBillingMethodRequest) Pointer() *GetBillingMethodRequest {
	return &p
}

type UpdateBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
	Description       *string `json:"description"`
}

func NewUpdateBillingMethodRequestFromJson(data string) UpdateBillingMethodRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBillingMethodRequestFromDict(dict)
}

func NewUpdateBillingMethodRequestFromDict(data map[string]interface{}) UpdateBillingMethodRequest {
	return UpdateBillingMethodRequest{
		AccountToken:      core.CastString(data["accountToken"]),
		BillingMethodName: core.CastString(data["billingMethodName"]),
		Description:       core.CastString(data["description"]),
	}
}

func (p UpdateBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
		"description":       p.Description,
	}
}

func (p UpdateBillingMethodRequest) Pointer() *UpdateBillingMethodRequest {
	return &p
}

type DeleteBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
}

func NewDeleteBillingMethodRequestFromJson(data string) DeleteBillingMethodRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBillingMethodRequestFromDict(dict)
}

func NewDeleteBillingMethodRequestFromDict(data map[string]interface{}) DeleteBillingMethodRequest {
	return DeleteBillingMethodRequest{
		AccountToken:      core.CastString(data["accountToken"]),
		BillingMethodName: core.CastString(data["billingMethodName"]),
	}
}

func (p DeleteBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
	}
}

func (p DeleteBillingMethodRequest) Pointer() *DeleteBillingMethodRequest {
	return &p
}

type DescribeReceiptsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeReceiptsRequestFromJson(data string) DescribeReceiptsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiptsRequestFromDict(dict)
}

func NewDescribeReceiptsRequestFromDict(data map[string]interface{}) DescribeReceiptsRequest {
	return DescribeReceiptsRequest{
		AccountToken: core.CastString(data["accountToken"]),
		PageToken:    core.CastString(data["pageToken"]),
		Limit:        core.CastInt32(data["limit"]),
	}
}

func (p DescribeReceiptsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeReceiptsRequest) Pointer() *DescribeReceiptsRequest {
	return &p
}

type DescribeBillingsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	Year            *int32  `json:"year"`
	Month           *int32  `json:"month"`
	Region          *string `json:"region"`
	Service         *string `json:"service"`
}

func NewDescribeBillingsRequestFromJson(data string) DescribeBillingsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBillingsRequestFromDict(dict)
}

func NewDescribeBillingsRequestFromDict(data map[string]interface{}) DescribeBillingsRequest {
	return DescribeBillingsRequest{
		AccountToken: core.CastString(data["accountToken"]),
		ProjectName:  core.CastString(data["projectName"]),
		Year:         core.CastInt32(data["year"]),
		Month:        core.CastInt32(data["month"]),
		Region:       core.CastString(data["region"]),
		Service:      core.CastString(data["service"]),
	}
}

func (p DescribeBillingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
		"year":         p.Year,
		"month":        p.Month,
		"region":       p.Region,
		"service":      p.Service,
	}
}

func (p DescribeBillingsRequest) Pointer() *DescribeBillingsRequest {
	return &p
}

type DescribeDumpProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeDumpProgressesRequestFromJson(data string) DescribeDumpProgressesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDumpProgressesRequestFromDict(dict)
}

func NewDescribeDumpProgressesRequestFromDict(data map[string]interface{}) DescribeDumpProgressesRequest {
	return DescribeDumpProgressesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeDumpProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeDumpProgressesRequest) Pointer() *DescribeDumpProgressesRequest {
	return &p
}

type GetDumpProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
}

func NewGetDumpProgressRequestFromJson(data string) GetDumpProgressRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDumpProgressRequestFromDict(dict)
}

func NewGetDumpProgressRequestFromDict(data map[string]interface{}) GetDumpProgressRequest {
	return GetDumpProgressRequest{
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p GetDumpProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetDumpProgressRequest) Pointer() *GetDumpProgressRequest {
	return &p
}

type WaitDumpUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewWaitDumpUserDataRequestFromJson(data string) WaitDumpUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitDumpUserDataRequestFromDict(dict)
}

func NewWaitDumpUserDataRequestFromDict(data map[string]interface{}) WaitDumpUserDataRequest {
	return WaitDumpUserDataRequest{
		TransactionId:    core.CastString(data["transactionId"]),
		UserId:           core.CastString(data["userId"]),
		MicroserviceName: core.CastString(data["microserviceName"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p WaitDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitDumpUserDataRequest) Pointer() *WaitDumpUserDataRequest {
	return &p
}

type ArchiveDumpUserDataRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
}

func NewArchiveDumpUserDataRequestFromJson(data string) ArchiveDumpUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewArchiveDumpUserDataRequestFromDict(dict)
}

func NewArchiveDumpUserDataRequestFromDict(data map[string]interface{}) ArchiveDumpUserDataRequest {
	return ArchiveDumpUserDataRequest{
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p ArchiveDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p ArchiveDumpUserDataRequest) Pointer() *ArchiveDumpUserDataRequest {
	return &p
}

type DumpUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDumpUserDataRequestFromJson(data string) DumpUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataRequestFromDict(dict)
}

func NewDumpUserDataRequestFromDict(data map[string]interface{}) DumpUserDataRequest {
	return DumpUserDataRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataRequest) Pointer() *DumpUserDataRequest {
	return &p
}

type GetDumpUserDataRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
}

func NewGetDumpUserDataRequestFromJson(data string) GetDumpUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDumpUserDataRequestFromDict(dict)
}

func NewGetDumpUserDataRequestFromDict(data map[string]interface{}) GetDumpUserDataRequest {
	return GetDumpUserDataRequest{
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p GetDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetDumpUserDataRequest) Pointer() *GetDumpUserDataRequest {
	return &p
}

type DescribeCleanProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCleanProgressesRequestFromJson(data string) DescribeCleanProgressesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCleanProgressesRequestFromDict(dict)
}

func NewDescribeCleanProgressesRequestFromDict(data map[string]interface{}) DescribeCleanProgressesRequest {
	return DescribeCleanProgressesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeCleanProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeCleanProgressesRequest) Pointer() *DescribeCleanProgressesRequest {
	return &p
}

type GetCleanProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
}

func NewGetCleanProgressRequestFromJson(data string) GetCleanProgressRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCleanProgressRequestFromDict(dict)
}

func NewGetCleanProgressRequestFromDict(data map[string]interface{}) GetCleanProgressRequest {
	return GetCleanProgressRequest{
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p GetCleanProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetCleanProgressRequest) Pointer() *GetCleanProgressRequest {
	return &p
}

type WaitCleanUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewWaitCleanUserDataRequestFromJson(data string) WaitCleanUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitCleanUserDataRequestFromDict(dict)
}

func NewWaitCleanUserDataRequestFromDict(data map[string]interface{}) WaitCleanUserDataRequest {
	return WaitCleanUserDataRequest{
		TransactionId:    core.CastString(data["transactionId"]),
		UserId:           core.CastString(data["userId"]),
		MicroserviceName: core.CastString(data["microserviceName"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p WaitCleanUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitCleanUserDataRequest) Pointer() *WaitCleanUserDataRequest {
	return &p
}

type CleanUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCleanUserDataRequestFromJson(data string) CleanUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataRequestFromDict(dict)
}

func NewCleanUserDataRequestFromDict(data map[string]interface{}) CleanUserDataRequest {
	return CleanUserDataRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataRequest) Pointer() *CleanUserDataRequest {
	return &p
}

type DescribeImportProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeImportProgressesRequestFromJson(data string) DescribeImportProgressesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeImportProgressesRequestFromDict(dict)
}

func NewDescribeImportProgressesRequestFromDict(data map[string]interface{}) DescribeImportProgressesRequest {
	return DescribeImportProgressesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeImportProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeImportProgressesRequest) Pointer() *DescribeImportProgressesRequest {
	return &p
}

type GetImportProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
}

func NewGetImportProgressRequestFromJson(data string) GetImportProgressRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetImportProgressRequestFromDict(dict)
}

func NewGetImportProgressRequestFromDict(data map[string]interface{}) GetImportProgressRequest {
	return GetImportProgressRequest{
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p GetImportProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetImportProgressRequest) Pointer() *GetImportProgressRequest {
	return &p
}

type WaitImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewWaitImportUserDataRequestFromJson(data string) WaitImportUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWaitImportUserDataRequestFromDict(dict)
}

func NewWaitImportUserDataRequestFromDict(data map[string]interface{}) WaitImportUserDataRequest {
	return WaitImportUserDataRequest{
		TransactionId:    core.CastString(data["transactionId"]),
		UserId:           core.CastString(data["userId"]),
		MicroserviceName: core.CastString(data["microserviceName"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p WaitImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitImportUserDataRequest) Pointer() *WaitImportUserDataRequest {
	return &p
}

type PrepareImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPrepareImportUserDataRequestFromJson(data string) PrepareImportUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataRequestFromDict(dict)
}

func NewPrepareImportUserDataRequestFromDict(data map[string]interface{}) PrepareImportUserDataRequest {
	return PrepareImportUserDataRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataRequest) Pointer() *PrepareImportUserDataRequest {
	return &p
}

type ImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewImportUserDataRequestFromJson(data string) ImportUserDataRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataRequestFromDict(dict)
}

func NewImportUserDataRequestFromDict(data map[string]interface{}) ImportUserDataRequest {
	return ImportUserDataRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataRequest) Pointer() *ImportUserDataRequest {
	return &p
}

type DescribeImportErrorLogsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeImportErrorLogsRequestFromJson(data string) DescribeImportErrorLogsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeImportErrorLogsRequestFromDict(dict)
}

func NewDescribeImportErrorLogsRequestFromDict(data map[string]interface{}) DescribeImportErrorLogsRequest {
	return DescribeImportErrorLogsRequest{
		TransactionId: core.CastString(data["transactionId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeImportErrorLogsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeImportErrorLogsRequest) Pointer() *DescribeImportErrorLogsRequest {
	return &p
}

type GetImportErrorLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	ErrorLogName    *string `json:"errorLogName"`
}

func NewGetImportErrorLogRequestFromJson(data string) GetImportErrorLogRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetImportErrorLogRequestFromDict(dict)
}

func NewGetImportErrorLogRequestFromDict(data map[string]interface{}) GetImportErrorLogRequest {
	return GetImportErrorLogRequest{
		TransactionId: core.CastString(data["transactionId"]),
		ErrorLogName:  core.CastString(data["errorLogName"]),
	}
}

func (p GetImportErrorLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
		"errorLogName":  p.ErrorLogName,
	}
}

func (p GetImportErrorLogRequest) Pointer() *GetImportErrorLogRequest {
	return &p
}
