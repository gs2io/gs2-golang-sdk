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

package identifier

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeUsersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateUserRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
}

type UpdateUserRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	Description *core.String	`json:"description"`
}

type GetUserRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type DeleteUserRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type DescribeSecurityPoliciesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type DescribeCommonSecurityPoliciesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	OwnerId *core.String	`json:"ownerId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	Policy *core.String	`json:"policy"`
}

type UpdateSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	SecurityPolicyName *core.String	`json:"securityPolicyName"`
	Description *core.String	`json:"description"`
	Policy *core.String	`json:"policy"`
}

type GetSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	SecurityPolicyName *core.String	`json:"securityPolicyName"`
}

type DeleteSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	SecurityPolicyName *core.String	`json:"securityPolicyName"`
}

type DescribeIdentifiersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateIdentifierRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type GetIdentifierRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	ClientId *core.String	`json:"clientId"`
}

type DeleteIdentifierRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	ClientId *core.String	`json:"clientId"`
}

type DescribePasswordsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreatePasswordRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	Password *core.String	`json:"password"`
}

type GetPasswordRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type DeletePasswordRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type GetHasSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
}

type AttachSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	SecurityPolicyId *core.String	`json:"securityPolicyId"`
}

type DetachSecurityPolicyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	SecurityPolicyId *core.String	`json:"securityPolicyId"`
}

type LoginRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	ClientId *core.String	`json:"clientId"`
	ClientSecret *core.String	`json:"clientSecret"`
}

type LoginByUserRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	UserName *core.String	`json:"userName"`
	Password *core.String	`json:"password"`
}
