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

package key

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Description *core.String	`json:"description"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type DescribeKeysRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
}

type UpdateKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	KeyName *core.String	`json:"keyName"`
	Description *core.String	`json:"description"`
}

type GetKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	KeyName *core.String	`json:"keyName"`
}

type DeleteKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	KeyName *core.String	`json:"keyName"`
}

type EncryptRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	KeyName *core.String	`json:"keyName"`
	Data *core.String	`json:"data"`
}

type DecryptRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	KeyName *core.String	`json:"keyName"`
	Data *core.String	`json:"data"`
}

type DescribeGitHubApiKeysRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateGitHubApiKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	ApiKey *core.String	`json:"apiKey"`
	EncryptionKeyName *core.String	`json:"encryptionKeyName"`
}

type UpdateGitHubApiKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	ApiKeyName *core.String	`json:"apiKeyName"`
	Description *core.String	`json:"description"`
	ApiKey *core.String	`json:"apiKey"`
	EncryptionKeyName *core.String	`json:"encryptionKeyName"`
}

type GetGitHubApiKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	ApiKeyName *core.String	`json:"apiKeyName"`
}

type DeleteGitHubApiKeyRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	ApiKeyName *core.String	`json:"apiKeyName"`
}
