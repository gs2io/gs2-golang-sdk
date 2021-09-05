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
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Email *string `json:"email"`
    FullName *string `json:"fullName"`
    CompanyName *string `json:"companyName"`
    Password *string `json:"password"`
}

func NewCreateAccountRequestFromJson(data string) CreateAccountRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateAccountRequestFromDict(dict)
}

func NewCreateAccountRequestFromDict(data map[string]interface{}) CreateAccountRequest {
    return CreateAccountRequest {
        Email: core.CastString(data["email"]),
        FullName: core.CastString(data["fullName"]),
        CompanyName: core.CastString(data["companyName"]),
        Password: core.CastString(data["password"]),
    }
}

func (p CreateAccountRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "email": p.Email,
        "fullName": p.FullName,
        "companyName": p.CompanyName,
        "password": p.Password,
    }
}

func (p CreateAccountRequest) Pointer() *CreateAccountRequest {
    return &p
}

type VerifyRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    VerifyToken *string `json:"verifyToken"`
}

func NewVerifyRequestFromJson(data string) VerifyRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyRequestFromDict(dict)
}

func NewVerifyRequestFromDict(data map[string]interface{}) VerifyRequest {
    return VerifyRequest {
        VerifyToken: core.CastString(data["verifyToken"]),
    }
}

func (p VerifyRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "verifyToken": p.VerifyToken,
    }
}

func (p VerifyRequest) Pointer() *VerifyRequest {
    return &p
}

type SignInRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Email *string `json:"email"`
    Password *string `json:"password"`
}

func NewSignInRequestFromJson(data string) SignInRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSignInRequestFromDict(dict)
}

func NewSignInRequestFromDict(data map[string]interface{}) SignInRequest {
    return SignInRequest {
        Email: core.CastString(data["email"]),
        Password: core.CastString(data["password"]),
    }
}

func (p SignInRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "email": p.Email,
        "password": p.Password,
    }
}

func (p SignInRequest) Pointer() *SignInRequest {
    return &p
}

type IssueAccountTokenRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountName *string `json:"accountName"`
}

func NewIssueAccountTokenRequestFromJson(data string) IssueAccountTokenRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssueAccountTokenRequestFromDict(dict)
}

func NewIssueAccountTokenRequestFromDict(data map[string]interface{}) IssueAccountTokenRequest {
    return IssueAccountTokenRequest {
        AccountName: core.CastString(data["accountName"]),
    }
}

func (p IssueAccountTokenRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountName": p.AccountName,
    }
}

func (p IssueAccountTokenRequest) Pointer() *IssueAccountTokenRequest {
    return &p
}

type ForgetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Email *string `json:"email"`
}

func NewForgetRequestFromJson(data string) ForgetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewForgetRequestFromDict(dict)
}

func NewForgetRequestFromDict(data map[string]interface{}) ForgetRequest {
    return ForgetRequest {
        Email: core.CastString(data["email"]),
    }
}

func (p ForgetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "email": p.Email,
    }
}

func (p ForgetRequest) Pointer() *ForgetRequest {
    return &p
}

type IssuePasswordRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    IssuePasswordToken *string `json:"issuePasswordToken"`
}

func NewIssuePasswordRequestFromJson(data string) IssuePasswordRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssuePasswordRequestFromDict(dict)
}

func NewIssuePasswordRequestFromDict(data map[string]interface{}) IssuePasswordRequest {
    return IssuePasswordRequest {
        IssuePasswordToken: core.CastString(data["issuePasswordToken"]),
    }
}

func (p IssuePasswordRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "issuePasswordToken": p.IssuePasswordToken,
    }
}

func (p IssuePasswordRequest) Pointer() *IssuePasswordRequest {
    return &p
}

type UpdateAccountRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Email *string `json:"email"`
    FullName *string `json:"fullName"`
    CompanyName *string `json:"companyName"`
    Password *string `json:"password"`
    AccountToken *string `json:"accountToken"`
}

func NewUpdateAccountRequestFromJson(data string) UpdateAccountRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateAccountRequestFromDict(dict)
}

func NewUpdateAccountRequestFromDict(data map[string]interface{}) UpdateAccountRequest {
    return UpdateAccountRequest {
        Email: core.CastString(data["email"]),
        FullName: core.CastString(data["fullName"]),
        CompanyName: core.CastString(data["companyName"]),
        Password: core.CastString(data["password"]),
        AccountToken: core.CastString(data["accountToken"]),
    }
}

func (p UpdateAccountRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "email": p.Email,
        "fullName": p.FullName,
        "companyName": p.CompanyName,
        "password": p.Password,
        "accountToken": p.AccountToken,
    }
}

func (p UpdateAccountRequest) Pointer() *UpdateAccountRequest {
    return &p
}

type DeleteAccountRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
}

func NewDeleteAccountRequestFromJson(data string) DeleteAccountRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteAccountRequestFromDict(dict)
}

func NewDeleteAccountRequestFromDict(data map[string]interface{}) DeleteAccountRequest {
    return DeleteAccountRequest {
    }
}

func (p DeleteAccountRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteAccountRequest) Pointer() *DeleteAccountRequest {
    return &p
}

type DescribeProjectsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeProjectsRequestFromJson(data string) DescribeProjectsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeProjectsRequestFromDict(dict)
}

func NewDescribeProjectsRequestFromDict(data map[string]interface{}) DescribeProjectsRequest {
    return DescribeProjectsRequest {
        AccountToken: core.CastString(data["accountToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeProjectsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeProjectsRequest) Pointer() *DescribeProjectsRequest {
    return &p
}

type CreateProjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Plan *string `json:"plan"`
    BillingMethodName *string `json:"billingMethodName"`
    EnableEventBridge *string `json:"enableEventBridge"`
    EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
    EventBridgeAwsRegion *string `json:"eventBridgeAwsRegion"`
}

func NewCreateProjectRequestFromJson(data string) CreateProjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateProjectRequestFromDict(dict)
}

func NewCreateProjectRequestFromDict(data map[string]interface{}) CreateProjectRequest {
    return CreateProjectRequest {
        AccountToken: core.CastString(data["accountToken"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Plan: core.CastString(data["plan"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
        EnableEventBridge: core.CastString(data["enableEventBridge"]),
        EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
        EventBridgeAwsRegion: core.CastString(data["eventBridgeAwsRegion"]),
    }
}

func (p CreateProjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "name": p.Name,
        "description": p.Description,
        "plan": p.Plan,
        "billingMethodName": p.BillingMethodName,
        "enableEventBridge": p.EnableEventBridge,
        "eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
        "eventBridgeAwsRegion": p.EventBridgeAwsRegion,
    }
}

func (p CreateProjectRequest) Pointer() *CreateProjectRequest {
    return &p
}

type GetProjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    ProjectName *string `json:"projectName"`
}

func NewGetProjectRequestFromJson(data string) GetProjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetProjectRequestFromDict(dict)
}

func NewGetProjectRequestFromDict(data map[string]interface{}) GetProjectRequest {
    return GetProjectRequest {
        AccountToken: core.CastString(data["accountToken"]),
        ProjectName: core.CastString(data["projectName"]),
    }
}

func (p GetProjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "projectName": p.ProjectName,
    }
}

func (p GetProjectRequest) Pointer() *GetProjectRequest {
    return &p
}

type GetProjectTokenRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    ProjectName *string `json:"projectName"`
    AccountToken *string `json:"accountToken"`
}

func NewGetProjectTokenRequestFromJson(data string) GetProjectTokenRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetProjectTokenRequestFromDict(dict)
}

func NewGetProjectTokenRequestFromDict(data map[string]interface{}) GetProjectTokenRequest {
    return GetProjectTokenRequest {
        ProjectName: core.CastString(data["projectName"]),
        AccountToken: core.CastString(data["accountToken"]),
    }
}

func (p GetProjectTokenRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "projectName": p.ProjectName,
        "accountToken": p.AccountToken,
    }
}

func (p GetProjectTokenRequest) Pointer() *GetProjectTokenRequest {
    return &p
}

type GetProjectTokenByIdentifierRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountName *string `json:"accountName"`
    ProjectName *string `json:"projectName"`
    UserName *string `json:"userName"`
    Password *string `json:"password"`
}

func NewGetProjectTokenByIdentifierRequestFromJson(data string) GetProjectTokenByIdentifierRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetProjectTokenByIdentifierRequestFromDict(dict)
}

func NewGetProjectTokenByIdentifierRequestFromDict(data map[string]interface{}) GetProjectTokenByIdentifierRequest {
    return GetProjectTokenByIdentifierRequest {
        AccountName: core.CastString(data["accountName"]),
        ProjectName: core.CastString(data["projectName"]),
        UserName: core.CastString(data["userName"]),
        Password: core.CastString(data["password"]),
    }
}

func (p GetProjectTokenByIdentifierRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountName": p.AccountName,
        "projectName": p.ProjectName,
        "userName": p.UserName,
        "password": p.Password,
    }
}

func (p GetProjectTokenByIdentifierRequest) Pointer() *GetProjectTokenByIdentifierRequest {
    return &p
}

type UpdateProjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    ProjectName *string `json:"projectName"`
    Description *string `json:"description"`
    Plan *string `json:"plan"`
    BillingMethodName *string `json:"billingMethodName"`
    EnableEventBridge *string `json:"enableEventBridge"`
    EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
    EventBridgeAwsRegion *string `json:"eventBridgeAwsRegion"`
}

func NewUpdateProjectRequestFromJson(data string) UpdateProjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateProjectRequestFromDict(dict)
}

func NewUpdateProjectRequestFromDict(data map[string]interface{}) UpdateProjectRequest {
    return UpdateProjectRequest {
        AccountToken: core.CastString(data["accountToken"]),
        ProjectName: core.CastString(data["projectName"]),
        Description: core.CastString(data["description"]),
        Plan: core.CastString(data["plan"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
        EnableEventBridge: core.CastString(data["enableEventBridge"]),
        EventBridgeAwsAccountId: core.CastString(data["eventBridgeAwsAccountId"]),
        EventBridgeAwsRegion: core.CastString(data["eventBridgeAwsRegion"]),
    }
}

func (p UpdateProjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "projectName": p.ProjectName,
        "description": p.Description,
        "plan": p.Plan,
        "billingMethodName": p.BillingMethodName,
        "enableEventBridge": p.EnableEventBridge,
        "eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
        "eventBridgeAwsRegion": p.EventBridgeAwsRegion,
    }
}

func (p UpdateProjectRequest) Pointer() *UpdateProjectRequest {
    return &p
}

type DeleteProjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    ProjectName *string `json:"projectName"`
}

func NewDeleteProjectRequestFromJson(data string) DeleteProjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteProjectRequestFromDict(dict)
}

func NewDeleteProjectRequestFromDict(data map[string]interface{}) DeleteProjectRequest {
    return DeleteProjectRequest {
        AccountToken: core.CastString(data["accountToken"]),
        ProjectName: core.CastString(data["projectName"]),
    }
}

func (p DeleteProjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "projectName": p.ProjectName,
    }
}

func (p DeleteProjectRequest) Pointer() *DeleteProjectRequest {
    return &p
}

type DescribeBillingMethodsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeBillingMethodsRequestFromJson(data string) DescribeBillingMethodsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBillingMethodsRequestFromDict(dict)
}

func NewDescribeBillingMethodsRequestFromDict(data map[string]interface{}) DescribeBillingMethodsRequest {
    return DescribeBillingMethodsRequest {
        AccountToken: core.CastString(data["accountToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeBillingMethodsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeBillingMethodsRequest) Pointer() *DescribeBillingMethodsRequest {
    return &p
}

type CreateBillingMethodRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    Description *string `json:"description"`
    MethodType *string `json:"methodType"`
    CardCustomerId *string `json:"cardCustomerId"`
    PartnerId *string `json:"partnerId"`
}

func NewCreateBillingMethodRequestFromJson(data string) CreateBillingMethodRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateBillingMethodRequestFromDict(dict)
}

func NewCreateBillingMethodRequestFromDict(data map[string]interface{}) CreateBillingMethodRequest {
    return CreateBillingMethodRequest {
        AccountToken: core.CastString(data["accountToken"]),
        Description: core.CastString(data["description"]),
        MethodType: core.CastString(data["methodType"]),
        CardCustomerId: core.CastString(data["cardCustomerId"]),
        PartnerId: core.CastString(data["partnerId"]),
    }
}

func (p CreateBillingMethodRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "description": p.Description,
        "methodType": p.MethodType,
        "cardCustomerId": p.CardCustomerId,
        "partnerId": p.PartnerId,
    }
}

func (p CreateBillingMethodRequest) Pointer() *CreateBillingMethodRequest {
    return &p
}

type GetBillingMethodRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    BillingMethodName *string `json:"billingMethodName"`
}

func NewGetBillingMethodRequestFromJson(data string) GetBillingMethodRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBillingMethodRequestFromDict(dict)
}

func NewGetBillingMethodRequestFromDict(data map[string]interface{}) GetBillingMethodRequest {
    return GetBillingMethodRequest {
        AccountToken: core.CastString(data["accountToken"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
    }
}

func (p GetBillingMethodRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "billingMethodName": p.BillingMethodName,
    }
}

func (p GetBillingMethodRequest) Pointer() *GetBillingMethodRequest {
    return &p
}

type UpdateBillingMethodRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    BillingMethodName *string `json:"billingMethodName"`
    Description *string `json:"description"`
}

func NewUpdateBillingMethodRequestFromJson(data string) UpdateBillingMethodRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateBillingMethodRequestFromDict(dict)
}

func NewUpdateBillingMethodRequestFromDict(data map[string]interface{}) UpdateBillingMethodRequest {
    return UpdateBillingMethodRequest {
        AccountToken: core.CastString(data["accountToken"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
        Description: core.CastString(data["description"]),
    }
}

func (p UpdateBillingMethodRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "billingMethodName": p.BillingMethodName,
        "description": p.Description,
    }
}

func (p UpdateBillingMethodRequest) Pointer() *UpdateBillingMethodRequest {
    return &p
}

type DeleteBillingMethodRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    BillingMethodName *string `json:"billingMethodName"`
}

func NewDeleteBillingMethodRequestFromJson(data string) DeleteBillingMethodRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBillingMethodRequestFromDict(dict)
}

func NewDeleteBillingMethodRequestFromDict(data map[string]interface{}) DeleteBillingMethodRequest {
    return DeleteBillingMethodRequest {
        AccountToken: core.CastString(data["accountToken"]),
        BillingMethodName: core.CastString(data["billingMethodName"]),
    }
}

func (p DeleteBillingMethodRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "billingMethodName": p.BillingMethodName,
    }
}

func (p DeleteBillingMethodRequest) Pointer() *DeleteBillingMethodRequest {
    return &p
}

type DescribeReceiptsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiptsRequestFromJson(data string) DescribeReceiptsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReceiptsRequestFromDict(dict)
}

func NewDescribeReceiptsRequestFromDict(data map[string]interface{}) DescribeReceiptsRequest {
    return DescribeReceiptsRequest {
        AccountToken: core.CastString(data["accountToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiptsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiptsRequest) Pointer() *DescribeReceiptsRequest {
    return &p
}

type DescribeBillingsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    AccountToken *string `json:"accountToken"`
    ProjectName *string `json:"projectName"`
    Year *int32 `json:"year"`
    Month *int32 `json:"month"`
    Region *string `json:"region"`
    Service *string `json:"service"`
}

func NewDescribeBillingsRequestFromJson(data string) DescribeBillingsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBillingsRequestFromDict(dict)
}

func NewDescribeBillingsRequestFromDict(data map[string]interface{}) DescribeBillingsRequest {
    return DescribeBillingsRequest {
        AccountToken: core.CastString(data["accountToken"]),
        ProjectName: core.CastString(data["projectName"]),
        Year: core.CastInt32(data["year"]),
        Month: core.CastInt32(data["month"]),
        Region: core.CastString(data["region"]),
        Service: core.CastString(data["service"]),
    }
}

func (p DescribeBillingsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "accountToken": p.AccountToken,
        "projectName": p.ProjectName,
        "year": p.Year,
        "month": p.Month,
        "region": p.Region,
        "service": p.Service,
    }
}

func (p DescribeBillingsRequest) Pointer() *DescribeBillingsRequest {
    return &p
}