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

package dictionary

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	EntryScript *ScriptSetting	`json:"entryScript"`
	DuplicateEntryScript *ScriptSetting	`json:"duplicateEntryScript"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Description *string	`json:"description"`
	EntryScript *ScriptSetting	`json:"entryScript"`
	DuplicateEntryScript *ScriptSetting	`json:"duplicateEntryScript"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeEntryModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetEntryModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryName *string	`json:"entryName"`
}

type DescribeEntryModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateEntryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
}

type GetEntryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryName *string	`json:"entryName"`
}

type UpdateEntryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryName *string	`json:"entryName"`
	Description *string	`json:"description"`
	Metadata *string	`json:"metadata"`
}

type DeleteEntryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryName *string	`json:"entryName"`
}

type DescribeEntriesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeEntriesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddEntriesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	EntryModelNames []string	`json:"entryModelNames"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetEntryRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryModelName *string	`json:"entryModelName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetEntryByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	EntryModelName *string	`json:"entryModelName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetEntryWithSignatureRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	EntryModelName *string	`json:"entryModelName"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetEntryWithSignatureByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	EntryModelName *string	`json:"entryModelName"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ResetByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type AddEntriesByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *string	`json:"stampSheet"`
	KeyId *string	`json:"keyId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetCurrentEntryMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateCurrentEntryMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Settings *string	`json:"settings"`
}

type UpdateCurrentEntryMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}
