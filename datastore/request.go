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

package datastore

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	Name             *string            `json:"name"`
	Description      *string            `json:"description"`
	LogSetting       *LogSetting        `json:"logSetting"`
	DoneUploadScript *ScriptSetting     `json:"doneUploadScript"`
}

type GetNamespaceStatusRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	Description      *string            `json:"description"`
	LogSetting       *LogSetting        `json:"logSetting"`
	DoneUploadScript *ScriptSetting     `json:"doneUploadScript"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeDataObjectsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Status                 *string            `json:"status"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeDataObjectsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Status                 *string            `json:"status"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareUploadRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Name                   *string            `json:"name"`
	ContentType            *string            `json:"contentType"`
	Scope                  *string            `json:"scope"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	UpdateIfExists         *bool              `json:"updateIfExists"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareUploadByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Name                   *string            `json:"name"`
	ContentType            *string            `json:"contentType"`
	Scope                  *string            `json:"scope"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	UpdateIfExists         *bool              `json:"updateIfExists"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UpdateDataObjectRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	Scope                  *string            `json:"scope"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UpdateDataObjectByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	UserId                 *string            `json:"userId"`
	Scope                  *string            `json:"scope"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareReUploadRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	ContentType            *string            `json:"contentType"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareReUploadByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	UserId                 *string            `json:"userId"`
	ContentType            *string            `json:"contentType"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DoneUploadRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DoneUploadByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteDataObjectRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteDataObjectByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectName         *string            `json:"dataObjectName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareDownloadRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectId           *string            `json:"dataObjectId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareDownloadByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectId           *string            `json:"dataObjectId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareDownloadByGenerationRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectId           *string            `json:"dataObjectId"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareDownloadByGenerationAndUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectId           *string            `json:"dataObjectId"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareDownloadOwnDataRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareDownloadByUserIdAndDataObjectNameRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectName         *string            `json:"dataObjectName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PrepareDownloadOwnDataByGenerationRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectName         *string            `json:"dataObjectName"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeDataObjectHistoriesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeDataObjectHistoriesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectName         *string            `json:"dataObjectName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetDataObjectHistoryRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	DataObjectName         *string            `json:"dataObjectName"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetDataObjectHistoryByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	DataObjectName         *string            `json:"dataObjectName"`
	Generation             *string            `json:"generation"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
