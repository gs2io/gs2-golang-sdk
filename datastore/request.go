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
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesRequestFromDict(dict)
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
    return DescribeNamespacesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
    return &p
}

type CreateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    LogSetting *LogSetting `json:"logSetting"`
    DoneUploadScript *ScriptSetting `json:"doneUploadScript"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        DoneUploadScript: NewScriptSettingFromDict(core.CastMap(data["doneUploadScript"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
        "doneUploadScript": p.DoneUploadScript.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusRequestFromDict(dict)
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
    return GetNamespaceStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
    return &p
}

type GetNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceRequestFromDict(dict)
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
    return GetNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
    return &p
}

type UpdateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    LogSetting *LogSetting `json:"logSetting"`
    DoneUploadScript *ScriptSetting `json:"doneUploadScript"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        DoneUploadScript: NewScriptSettingFromDict(core.CastMap(data["doneUploadScript"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "logSetting": p.LogSetting.ToDict(),
        "doneUploadScript": p.DoneUploadScript.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceRequestFromDict(dict)
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
    return DeleteNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
    return &p
}

type DescribeDataObjectsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Status *string `json:"status"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeDataObjectsRequestFromJson(data string) DescribeDataObjectsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDataObjectsRequestFromDict(dict)
}

func NewDescribeDataObjectsRequestFromDict(data map[string]interface{}) DescribeDataObjectsRequest {
    return DescribeDataObjectsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Status: core.CastString(data["status"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeDataObjectsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "status": p.Status,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeDataObjectsRequest) Pointer() *DescribeDataObjectsRequest {
    return &p
}

type DescribeDataObjectsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Status *string `json:"status"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeDataObjectsByUserIdRequestFromJson(data string) DescribeDataObjectsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDataObjectsByUserIdRequestFromDict(dict)
}

func NewDescribeDataObjectsByUserIdRequestFromDict(data map[string]interface{}) DescribeDataObjectsByUserIdRequest {
    return DescribeDataObjectsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Status: core.CastString(data["status"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeDataObjectsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "status": p.Status,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeDataObjectsByUserIdRequest) Pointer() *DescribeDataObjectsByUserIdRequest {
    return &p
}

type PrepareUploadRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Name *string `json:"name"`
    ContentType *string `json:"contentType"`
    Scope *string `json:"scope"`
    AllowUserIds []string `json:"allowUserIds"`
    UpdateIfExists *bool `json:"updateIfExists"`
}

func NewPrepareUploadRequestFromJson(data string) PrepareUploadRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareUploadRequestFromDict(dict)
}

func NewPrepareUploadRequestFromDict(data map[string]interface{}) PrepareUploadRequest {
    return PrepareUploadRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Name: core.CastString(data["name"]),
        ContentType: core.CastString(data["contentType"]),
        Scope: core.CastString(data["scope"]),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        UpdateIfExists: core.CastBool(data["updateIfExists"]),
    }
}

func (p PrepareUploadRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "name": p.Name,
        "contentType": p.ContentType,
        "scope": p.Scope,
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
        "updateIfExists": p.UpdateIfExists,
    }
}

func (p PrepareUploadRequest) Pointer() *PrepareUploadRequest {
    return &p
}

type PrepareUploadByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Name *string `json:"name"`
    ContentType *string `json:"contentType"`
    Scope *string `json:"scope"`
    AllowUserIds []string `json:"allowUserIds"`
    UpdateIfExists *bool `json:"updateIfExists"`
}

func NewPrepareUploadByUserIdRequestFromJson(data string) PrepareUploadByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareUploadByUserIdRequestFromDict(dict)
}

func NewPrepareUploadByUserIdRequestFromDict(data map[string]interface{}) PrepareUploadByUserIdRequest {
    return PrepareUploadByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Name: core.CastString(data["name"]),
        ContentType: core.CastString(data["contentType"]),
        Scope: core.CastString(data["scope"]),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        UpdateIfExists: core.CastBool(data["updateIfExists"]),
    }
}

func (p PrepareUploadByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "name": p.Name,
        "contentType": p.ContentType,
        "scope": p.Scope,
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
        "updateIfExists": p.UpdateIfExists,
    }
}

func (p PrepareUploadByUserIdRequest) Pointer() *PrepareUploadByUserIdRequest {
    return &p
}

type UpdateDataObjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    AccessToken *string `json:"accessToken"`
    Scope *string `json:"scope"`
    AllowUserIds []string `json:"allowUserIds"`
}

func NewUpdateDataObjectRequestFromJson(data string) UpdateDataObjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateDataObjectRequestFromDict(dict)
}

func NewUpdateDataObjectRequestFromDict(data map[string]interface{}) UpdateDataObjectRequest {
    return UpdateDataObjectRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Scope: core.CastString(data["scope"]),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
    }
}

func (p UpdateDataObjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "accessToken": p.AccessToken,
        "scope": p.Scope,
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
    }
}

func (p UpdateDataObjectRequest) Pointer() *UpdateDataObjectRequest {
    return &p
}

type UpdateDataObjectByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    UserId *string `json:"userId"`
    Scope *string `json:"scope"`
    AllowUserIds []string `json:"allowUserIds"`
}

func NewUpdateDataObjectByUserIdRequestFromJson(data string) UpdateDataObjectByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateDataObjectByUserIdRequestFromDict(dict)
}

func NewUpdateDataObjectByUserIdRequestFromDict(data map[string]interface{}) UpdateDataObjectByUserIdRequest {
    return UpdateDataObjectByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        UserId: core.CastString(data["userId"]),
        Scope: core.CastString(data["scope"]),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
    }
}

func (p UpdateDataObjectByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "userId": p.UserId,
        "scope": p.Scope,
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
    }
}

func (p UpdateDataObjectByUserIdRequest) Pointer() *UpdateDataObjectByUserIdRequest {
    return &p
}

type PrepareReUploadRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    AccessToken *string `json:"accessToken"`
    ContentType *string `json:"contentType"`
}

func NewPrepareReUploadRequestFromJson(data string) PrepareReUploadRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareReUploadRequestFromDict(dict)
}

func NewPrepareReUploadRequestFromDict(data map[string]interface{}) PrepareReUploadRequest {
    return PrepareReUploadRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        AccessToken: core.CastString(data["accessToken"]),
        ContentType: core.CastString(data["contentType"]),
    }
}

func (p PrepareReUploadRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "accessToken": p.AccessToken,
        "contentType": p.ContentType,
    }
}

func (p PrepareReUploadRequest) Pointer() *PrepareReUploadRequest {
    return &p
}

type PrepareReUploadByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    UserId *string `json:"userId"`
    ContentType *string `json:"contentType"`
}

func NewPrepareReUploadByUserIdRequestFromJson(data string) PrepareReUploadByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareReUploadByUserIdRequestFromDict(dict)
}

func NewPrepareReUploadByUserIdRequestFromDict(data map[string]interface{}) PrepareReUploadByUserIdRequest {
    return PrepareReUploadByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        UserId: core.CastString(data["userId"]),
        ContentType: core.CastString(data["contentType"]),
    }
}

func (p PrepareReUploadByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "userId": p.UserId,
        "contentType": p.ContentType,
    }
}

func (p PrepareReUploadByUserIdRequest) Pointer() *PrepareReUploadByUserIdRequest {
    return &p
}

type DoneUploadRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    AccessToken *string `json:"accessToken"`
}

func NewDoneUploadRequestFromJson(data string) DoneUploadRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDoneUploadRequestFromDict(dict)
}

func NewDoneUploadRequestFromDict(data map[string]interface{}) DoneUploadRequest {
    return DoneUploadRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p DoneUploadRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "accessToken": p.AccessToken,
    }
}

func (p DoneUploadRequest) Pointer() *DoneUploadRequest {
    return &p
}

type DoneUploadByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectName *string `json:"dataObjectName"`
    UserId *string `json:"userId"`
}

func NewDoneUploadByUserIdRequestFromJson(data string) DoneUploadByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDoneUploadByUserIdRequestFromDict(dict)
}

func NewDoneUploadByUserIdRequestFromDict(data map[string]interface{}) DoneUploadByUserIdRequest {
    return DoneUploadByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DoneUploadByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectName": p.DataObjectName,
        "userId": p.UserId,
    }
}

func (p DoneUploadByUserIdRequest) Pointer() *DoneUploadByUserIdRequest {
    return &p
}

type DeleteDataObjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectName *string `json:"dataObjectName"`
}

func NewDeleteDataObjectRequestFromJson(data string) DeleteDataObjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteDataObjectRequestFromDict(dict)
}

func NewDeleteDataObjectRequestFromDict(data map[string]interface{}) DeleteDataObjectRequest {
    return DeleteDataObjectRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
    }
}

func (p DeleteDataObjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectName": p.DataObjectName,
    }
}

func (p DeleteDataObjectRequest) Pointer() *DeleteDataObjectRequest {
    return &p
}

type DeleteDataObjectByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectName *string `json:"dataObjectName"`
}

func NewDeleteDataObjectByUserIdRequestFromJson(data string) DeleteDataObjectByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteDataObjectByUserIdRequestFromDict(dict)
}

func NewDeleteDataObjectByUserIdRequestFromDict(data map[string]interface{}) DeleteDataObjectByUserIdRequest {
    return DeleteDataObjectByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
    }
}

func (p DeleteDataObjectByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectName": p.DataObjectName,
    }
}

func (p DeleteDataObjectByUserIdRequest) Pointer() *DeleteDataObjectByUserIdRequest {
    return &p
}

type PrepareDownloadRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectId *string `json:"dataObjectId"`
}

func NewPrepareDownloadRequestFromJson(data string) PrepareDownloadRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadRequestFromDict(dict)
}

func NewPrepareDownloadRequestFromDict(data map[string]interface{}) PrepareDownloadRequest {
    return PrepareDownloadRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectId: core.CastString(data["dataObjectId"]),
    }
}

func (p PrepareDownloadRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectId": p.DataObjectId,
    }
}

func (p PrepareDownloadRequest) Pointer() *PrepareDownloadRequest {
    return &p
}

type PrepareDownloadByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectId *string `json:"dataObjectId"`
}

func NewPrepareDownloadByUserIdRequestFromJson(data string) PrepareDownloadByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadByUserIdRequestFromDict(dict)
}

func NewPrepareDownloadByUserIdRequestFromDict(data map[string]interface{}) PrepareDownloadByUserIdRequest {
    return PrepareDownloadByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectId: core.CastString(data["dataObjectId"]),
    }
}

func (p PrepareDownloadByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectId": p.DataObjectId,
    }
}

func (p PrepareDownloadByUserIdRequest) Pointer() *PrepareDownloadByUserIdRequest {
    return &p
}

type PrepareDownloadByGenerationRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectId *string `json:"dataObjectId"`
    Generation *string `json:"generation"`
}

func NewPrepareDownloadByGenerationRequestFromJson(data string) PrepareDownloadByGenerationRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadByGenerationRequestFromDict(dict)
}

func NewPrepareDownloadByGenerationRequestFromDict(data map[string]interface{}) PrepareDownloadByGenerationRequest {
    return PrepareDownloadByGenerationRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectId: core.CastString(data["dataObjectId"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p PrepareDownloadByGenerationRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectId": p.DataObjectId,
        "generation": p.Generation,
    }
}

func (p PrepareDownloadByGenerationRequest) Pointer() *PrepareDownloadByGenerationRequest {
    return &p
}

type PrepareDownloadByGenerationAndUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectId *string `json:"dataObjectId"`
    Generation *string `json:"generation"`
}

func NewPrepareDownloadByGenerationAndUserIdRequestFromJson(data string) PrepareDownloadByGenerationAndUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadByGenerationAndUserIdRequestFromDict(dict)
}

func NewPrepareDownloadByGenerationAndUserIdRequestFromDict(data map[string]interface{}) PrepareDownloadByGenerationAndUserIdRequest {
    return PrepareDownloadByGenerationAndUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectId: core.CastString(data["dataObjectId"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p PrepareDownloadByGenerationAndUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectId": p.DataObjectId,
        "generation": p.Generation,
    }
}

func (p PrepareDownloadByGenerationAndUserIdRequest) Pointer() *PrepareDownloadByGenerationAndUserIdRequest {
    return &p
}

type PrepareDownloadOwnDataRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectName *string `json:"dataObjectName"`
}

func NewPrepareDownloadOwnDataRequestFromJson(data string) PrepareDownloadOwnDataRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadOwnDataRequestFromDict(dict)
}

func NewPrepareDownloadOwnDataRequestFromDict(data map[string]interface{}) PrepareDownloadOwnDataRequest {
    return PrepareDownloadOwnDataRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
    }
}

func (p PrepareDownloadOwnDataRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectName": p.DataObjectName,
    }
}

func (p PrepareDownloadOwnDataRequest) Pointer() *PrepareDownloadOwnDataRequest {
    return &p
}

type PrepareDownloadByUserIdAndDataObjectNameRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectName *string `json:"dataObjectName"`
}

func NewPrepareDownloadByUserIdAndDataObjectNameRequestFromJson(data string) PrepareDownloadByUserIdAndDataObjectNameRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadByUserIdAndDataObjectNameRequestFromDict(dict)
}

func NewPrepareDownloadByUserIdAndDataObjectNameRequestFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameRequest {
    return PrepareDownloadByUserIdAndDataObjectNameRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
    }
}

func (p PrepareDownloadByUserIdAndDataObjectNameRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectName": p.DataObjectName,
    }
}

func (p PrepareDownloadByUserIdAndDataObjectNameRequest) Pointer() *PrepareDownloadByUserIdAndDataObjectNameRequest {
    return &p
}

type PrepareDownloadOwnDataByGenerationRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectName *string `json:"dataObjectName"`
    Generation *string `json:"generation"`
}

func NewPrepareDownloadOwnDataByGenerationRequestFromJson(data string) PrepareDownloadOwnDataByGenerationRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadOwnDataByGenerationRequestFromDict(dict)
}

func NewPrepareDownloadOwnDataByGenerationRequestFromDict(data map[string]interface{}) PrepareDownloadOwnDataByGenerationRequest {
    return PrepareDownloadOwnDataByGenerationRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p PrepareDownloadOwnDataByGenerationRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectName": p.DataObjectName,
        "generation": p.Generation,
    }
}

func (p PrepareDownloadOwnDataByGenerationRequest) Pointer() *PrepareDownloadOwnDataByGenerationRequest {
    return &p
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectName *string `json:"dataObjectName"`
    Generation *string `json:"generation"`
}

func NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationRequestFromJson(data string) PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationRequestFromDict(dict)
}

func NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationRequestFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest {
    return PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectName": p.DataObjectName,
        "generation": p.Generation,
    }
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest) Pointer() *PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest {
    return &p
}

type RestoreDataObjectRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    DataObjectId *string `json:"dataObjectId"`
}

func NewRestoreDataObjectRequestFromJson(data string) RestoreDataObjectRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRestoreDataObjectRequestFromDict(dict)
}

func NewRestoreDataObjectRequestFromDict(data map[string]interface{}) RestoreDataObjectRequest {
    return RestoreDataObjectRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DataObjectId: core.CastString(data["dataObjectId"]),
    }
}

func (p RestoreDataObjectRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "dataObjectId": p.DataObjectId,
    }
}

func (p RestoreDataObjectRequest) Pointer() *RestoreDataObjectRequest {
    return &p
}

type DescribeDataObjectHistoriesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectName *string `json:"dataObjectName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeDataObjectHistoriesRequestFromJson(data string) DescribeDataObjectHistoriesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDataObjectHistoriesRequestFromDict(dict)
}

func NewDescribeDataObjectHistoriesRequestFromDict(data map[string]interface{}) DescribeDataObjectHistoriesRequest {
    return DescribeDataObjectHistoriesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeDataObjectHistoriesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectName": p.DataObjectName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeDataObjectHistoriesRequest) Pointer() *DescribeDataObjectHistoriesRequest {
    return &p
}

type DescribeDataObjectHistoriesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectName *string `json:"dataObjectName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeDataObjectHistoriesByUserIdRequestFromJson(data string) DescribeDataObjectHistoriesByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDataObjectHistoriesByUserIdRequestFromDict(dict)
}

func NewDescribeDataObjectHistoriesByUserIdRequestFromDict(data map[string]interface{}) DescribeDataObjectHistoriesByUserIdRequest {
    return DescribeDataObjectHistoriesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeDataObjectHistoriesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectName": p.DataObjectName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeDataObjectHistoriesByUserIdRequest) Pointer() *DescribeDataObjectHistoriesByUserIdRequest {
    return &p
}

type GetDataObjectHistoryRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    DataObjectName *string `json:"dataObjectName"`
    Generation *string `json:"generation"`
}

func NewGetDataObjectHistoryRequestFromJson(data string) GetDataObjectHistoryRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDataObjectHistoryRequestFromDict(dict)
}

func NewGetDataObjectHistoryRequestFromDict(data map[string]interface{}) GetDataObjectHistoryRequest {
    return GetDataObjectHistoryRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p GetDataObjectHistoryRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "dataObjectName": p.DataObjectName,
        "generation": p.Generation,
    }
}

func (p GetDataObjectHistoryRequest) Pointer() *GetDataObjectHistoryRequest {
    return &p
}

type GetDataObjectHistoryByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    DataObjectName *string `json:"dataObjectName"`
    Generation *string `json:"generation"`
}

func NewGetDataObjectHistoryByUserIdRequestFromJson(data string) GetDataObjectHistoryByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDataObjectHistoryByUserIdRequestFromDict(dict)
}

func NewGetDataObjectHistoryByUserIdRequestFromDict(data map[string]interface{}) GetDataObjectHistoryByUserIdRequest {
    return GetDataObjectHistoryByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        Generation: core.CastString(data["generation"]),
    }
}

func (p GetDataObjectHistoryByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "dataObjectName": p.DataObjectName,
        "generation": p.Generation,
    }
}

func (p GetDataObjectHistoryByUserIdRequest) Pointer() *GetDataObjectHistoryByUserIdRequest {
    return &p
}