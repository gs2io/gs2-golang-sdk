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

package distributor

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    AssumeUserId *string `json:"assumeUserId"`
    LogSetting *LogSetting `json:"logSetting"`
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
        AssumeUserId: core.CastString(data["assumeUserId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "assumeUserId": p.AssumeUserId,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    AssumeUserId *string `json:"assumeUserId"`
    LogSetting *LogSetting `json:"logSetting"`
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
        AssumeUserId: core.CastString(data["assumeUserId"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "assumeUserId": p.AssumeUserId,
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
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

type DescribeDistributorModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeDistributorModelMastersRequestFromJson(data string) DescribeDistributorModelMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDistributorModelMastersRequestFromDict(dict)
}

func NewDescribeDistributorModelMastersRequestFromDict(data map[string]interface{}) DescribeDistributorModelMastersRequest {
    return DescribeDistributorModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeDistributorModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeDistributorModelMastersRequest) Pointer() *DescribeDistributorModelMastersRequest {
    return &p
}

type CreateDistributorModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    InboxNamespaceId *string `json:"inboxNamespaceId"`
    WhiteListTargetIds []string `json:"whiteListTargetIds"`
}

func NewCreateDistributorModelMasterRequestFromJson(data string) CreateDistributorModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateDistributorModelMasterRequestFromDict(dict)
}

func NewCreateDistributorModelMasterRequestFromDict(data map[string]interface{}) CreateDistributorModelMasterRequest {
    return CreateDistributorModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        InboxNamespaceId: core.CastString(data["inboxNamespaceId"]),
        WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
    }
}

func (p CreateDistributorModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "inboxNamespaceId": p.InboxNamespaceId,
        "whiteListTargetIds": core.CastStringsFromDict(
            p.WhiteListTargetIds,
        ),
    }
}

func (p CreateDistributorModelMasterRequest) Pointer() *CreateDistributorModelMasterRequest {
    return &p
}

type GetDistributorModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DistributorName *string `json:"distributorName"`
}

func NewGetDistributorModelMasterRequestFromJson(data string) GetDistributorModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDistributorModelMasterRequestFromDict(dict)
}

func NewGetDistributorModelMasterRequestFromDict(data map[string]interface{}) GetDistributorModelMasterRequest {
    return GetDistributorModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DistributorName: core.CastString(data["distributorName"]),
    }
}

func (p GetDistributorModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "distributorName": p.DistributorName,
    }
}

func (p GetDistributorModelMasterRequest) Pointer() *GetDistributorModelMasterRequest {
    return &p
}

type UpdateDistributorModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DistributorName *string `json:"distributorName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    InboxNamespaceId *string `json:"inboxNamespaceId"`
    WhiteListTargetIds []string `json:"whiteListTargetIds"`
}

func NewUpdateDistributorModelMasterRequestFromJson(data string) UpdateDistributorModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateDistributorModelMasterRequestFromDict(dict)
}

func NewUpdateDistributorModelMasterRequestFromDict(data map[string]interface{}) UpdateDistributorModelMasterRequest {
    return UpdateDistributorModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DistributorName: core.CastString(data["distributorName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        InboxNamespaceId: core.CastString(data["inboxNamespaceId"]),
        WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
    }
}

func (p UpdateDistributorModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "distributorName": p.DistributorName,
        "description": p.Description,
        "metadata": p.Metadata,
        "inboxNamespaceId": p.InboxNamespaceId,
        "whiteListTargetIds": core.CastStringsFromDict(
            p.WhiteListTargetIds,
        ),
    }
}

func (p UpdateDistributorModelMasterRequest) Pointer() *UpdateDistributorModelMasterRequest {
    return &p
}

type DeleteDistributorModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DistributorName *string `json:"distributorName"`
}

func NewDeleteDistributorModelMasterRequestFromJson(data string) DeleteDistributorModelMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteDistributorModelMasterRequestFromDict(dict)
}

func NewDeleteDistributorModelMasterRequestFromDict(data map[string]interface{}) DeleteDistributorModelMasterRequest {
    return DeleteDistributorModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DistributorName: core.CastString(data["distributorName"]),
    }
}

func (p DeleteDistributorModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "distributorName": p.DistributorName,
    }
}

func (p DeleteDistributorModelMasterRequest) Pointer() *DeleteDistributorModelMasterRequest {
    return &p
}

type DescribeDistributorModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeDistributorModelsRequestFromJson(data string) DescribeDistributorModelsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeDistributorModelsRequestFromDict(dict)
}

func NewDescribeDistributorModelsRequestFromDict(data map[string]interface{}) DescribeDistributorModelsRequest {
    return DescribeDistributorModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeDistributorModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeDistributorModelsRequest) Pointer() *DescribeDistributorModelsRequest {
    return &p
}

type GetDistributorModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DistributorName *string `json:"distributorName"`
}

func NewGetDistributorModelRequestFromJson(data string) GetDistributorModelRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetDistributorModelRequestFromDict(dict)
}

func NewGetDistributorModelRequestFromDict(data map[string]interface{}) GetDistributorModelRequest {
    return GetDistributorModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DistributorName: core.CastString(data["distributorName"]),
    }
}

func (p GetDistributorModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "distributorName": p.DistributorName,
    }
}

func (p GetDistributorModelRequest) Pointer() *GetDistributorModelRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExportMasterRequestFromDict(dict)
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
    return ExportMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
    return &p
}

type GetCurrentDistributorMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentDistributorMasterRequestFromJson(data string) GetCurrentDistributorMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentDistributorMasterRequestFromDict(dict)
}

func NewGetCurrentDistributorMasterRequestFromDict(data map[string]interface{}) GetCurrentDistributorMasterRequest {
    return GetCurrentDistributorMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentDistributorMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentDistributorMasterRequest) Pointer() *GetCurrentDistributorMasterRequest {
    return &p
}

type UpdateCurrentDistributorMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentDistributorMasterRequestFromJson(data string) UpdateCurrentDistributorMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentDistributorMasterRequestFromDict(dict)
}

func NewUpdateCurrentDistributorMasterRequestFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterRequest {
    return UpdateCurrentDistributorMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentDistributorMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentDistributorMasterRequest) Pointer() *UpdateCurrentDistributorMasterRequest {
    return &p
}

type UpdateCurrentDistributorMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentDistributorMasterFromGitHubRequestFromJson(data string) UpdateCurrentDistributorMasterFromGitHubRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentDistributorMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentDistributorMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterFromGitHubRequest {
    return UpdateCurrentDistributorMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentDistributorMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentDistributorMasterFromGitHubRequest) Pointer() *UpdateCurrentDistributorMasterFromGitHubRequest {
    return &p
}

type DistributeRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    DistributorName *string `json:"distributorName"`
    UserId *string `json:"userId"`
    DistributeResource *DistributeResource `json:"distributeResource"`
}

func NewDistributeRequestFromJson(data string) DistributeRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDistributeRequestFromDict(dict)
}

func NewDistributeRequestFromDict(data map[string]interface{}) DistributeRequest {
    return DistributeRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        DistributorName: core.CastString(data["distributorName"]),
        UserId: core.CastString(data["userId"]),
        DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
    }
}

func (p DistributeRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "distributorName": p.DistributorName,
        "userId": p.UserId,
        "distributeResource": p.DistributeResource.ToDict(),
    }
}

func (p DistributeRequest) Pointer() *DistributeRequest {
    return &p
}

type DistributeWithoutOverflowProcessRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    UserId *string `json:"userId"`
    DistributeResource *DistributeResource `json:"distributeResource"`
}

func NewDistributeWithoutOverflowProcessRequestFromJson(data string) DistributeWithoutOverflowProcessRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDistributeWithoutOverflowProcessRequestFromDict(dict)
}

func NewDistributeWithoutOverflowProcessRequestFromDict(data map[string]interface{}) DistributeWithoutOverflowProcessRequest {
    return DistributeWithoutOverflowProcessRequest {
        UserId: core.CastString(data["userId"]),
        DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
    }
}

func (p DistributeWithoutOverflowProcessRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "userId": p.UserId,
        "distributeResource": p.DistributeResource.ToDict(),
    }
}

func (p DistributeWithoutOverflowProcessRequest) Pointer() *DistributeWithoutOverflowProcessRequest {
    return &p
}

type RunStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewRunStampTaskRequestFromJson(data string) RunStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampTaskRequestFromDict(dict)
}

func NewRunStampTaskRequestFromDict(data map[string]interface{}) RunStampTaskRequest {
    return RunStampTaskRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p RunStampTaskRequest) Pointer() *RunStampTaskRequest {
    return &p
}

type RunStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewRunStampSheetRequestFromJson(data string) RunStampSheetRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampSheetRequestFromDict(dict)
}

func NewRunStampSheetRequestFromDict(data map[string]interface{}) RunStampSheetRequest {
    return RunStampSheetRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p RunStampSheetRequest) Pointer() *RunStampSheetRequest {
    return &p
}

type RunStampSheetExpressRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewRunStampSheetExpressRequestFromJson(data string) RunStampSheetExpressRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampSheetExpressRequestFromDict(dict)
}

func NewRunStampSheetExpressRequestFromDict(data map[string]interface{}) RunStampSheetExpressRequest {
    return RunStampSheetExpressRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampSheetExpressRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p RunStampSheetExpressRequest) Pointer() *RunStampSheetExpressRequest {
    return &p
}

type RunStampTaskWithoutNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    StampTask *string `json:"stampTask"`
    KeyId *string `json:"keyId"`
}

func NewRunStampTaskWithoutNamespaceRequestFromJson(data string) RunStampTaskWithoutNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampTaskWithoutNamespaceRequestFromDict(dict)
}

func NewRunStampTaskWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampTaskWithoutNamespaceRequest {
    return RunStampTaskWithoutNamespaceRequest {
        StampTask: core.CastString(data["stampTask"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampTaskWithoutNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampTask": p.StampTask,
        "keyId": p.KeyId,
    }
}

func (p RunStampTaskWithoutNamespaceRequest) Pointer() *RunStampTaskWithoutNamespaceRequest {
    return &p
}

type RunStampSheetWithoutNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewRunStampSheetWithoutNamespaceRequestFromJson(data string) RunStampSheetWithoutNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampSheetWithoutNamespaceRequestFromDict(dict)
}

func NewRunStampSheetWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampSheetWithoutNamespaceRequest {
    return RunStampSheetWithoutNamespaceRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampSheetWithoutNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p RunStampSheetWithoutNamespaceRequest) Pointer() *RunStampSheetWithoutNamespaceRequest {
    return &p
}

type RunStampSheetExpressWithoutNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewRunStampSheetExpressWithoutNamespaceRequestFromJson(data string) RunStampSheetExpressWithoutNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRunStampSheetExpressWithoutNamespaceRequestFromDict(dict)
}

func NewRunStampSheetExpressWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampSheetExpressWithoutNamespaceRequest {
    return RunStampSheetExpressWithoutNamespaceRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p RunStampSheetExpressWithoutNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p RunStampSheetExpressWithoutNamespaceRequest) Pointer() *RunStampSheetExpressWithoutNamespaceRequest {
    return &p
}