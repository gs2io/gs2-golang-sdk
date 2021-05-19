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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type CreateAccountRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Email *string	`json:"email"`
	FullName *string	`json:"fullName"`
	CompanyName *string	`json:"companyName"`
	Password *string	`json:"password"`
}

type VerifyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	VerifyToken *string	`json:"verifyToken"`
}

type SignInRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Email *string	`json:"email"`
	Password *string	`json:"password"`
}

type IssueAccountTokenRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountName *string	`json:"accountName"`
}

type ForgetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Email *string	`json:"email"`
}

type IssuePasswordRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	IssuePasswordToken *string	`json:"issuePasswordToken"`
}

type UpdateAccountRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Email *string	`json:"email"`
	FullName *string	`json:"fullName"`
	CompanyName *string	`json:"companyName"`
	Password *string	`json:"password"`
	AccountToken *string	`json:"accountToken"`
}

type DeleteAccountRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountName *string	`json:"accountName"`
}

type DescribeProjectsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateProjectRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Plan *string	`json:"plan"`
	BillingMethodName *string	`json:"billingMethodName"`
	EnableEventBridge *string	`json:"enableEventBridge"`
	EventBridgeAwsAccountId *string	`json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion *string	`json:"eventBridgeAwsRegion"`
}

type GetProjectRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	ProjectName *string	`json:"projectName"`
}

type GetProjectTokenRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	ProjectName *string	`json:"projectName"`
	AccountToken *string	`json:"accountToken"`
}

type GetProjectTokenByIdentifierRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountName *string	`json:"accountName"`
	ProjectName *string	`json:"projectName"`
	UserName *string	`json:"userName"`
	Password *string	`json:"password"`
}

type UpdateProjectRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	ProjectName *string	`json:"projectName"`
	Description *string	`json:"description"`
	Plan *string	`json:"plan"`
	BillingMethodName *string	`json:"billingMethodName"`
	EnableEventBridge *string	`json:"enableEventBridge"`
	EventBridgeAwsAccountId *string	`json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion *string	`json:"eventBridgeAwsRegion"`
}

type DeleteProjectRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	ProjectName *string	`json:"projectName"`
}

type DescribeBillingMethodsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateBillingMethodRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	Description *string	`json:"description"`
	MethodType *string	`json:"methodType"`
	CardCustomerId *string	`json:"cardCustomerId"`
	PartnerId *string	`json:"partnerId"`
}

type GetBillingMethodRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	BillingMethodName *string	`json:"billingMethodName"`
}

type UpdateBillingMethodRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	BillingMethodName *string	`json:"billingMethodName"`
	Description *string	`json:"description"`
}

type DeleteBillingMethodRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	BillingMethodName *string	`json:"billingMethodName"`
}

type DescribeReceiptsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type DescribeBillingsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	AccountToken *string	`json:"accountToken"`
	ProjectName *string	`json:"projectName"`
	Year *int32	`json:"year"`
	Month *int32	`json:"month"`
	Region *string	`json:"region"`
	Service *string	`json:"service"`
}
