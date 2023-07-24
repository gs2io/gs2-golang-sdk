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

package stateMachine

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
    StartScript *ScriptSetting `json:"startScript"`
    PassScript *ScriptSetting `json:"passScript"`
    ErrorScript *ScriptSetting `json:"errorScript"`
    LowestStateMachineVersion *int64 `json:"lowestStateMachineVersion"`
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
        StartScript: NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
        PassScript: NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
        ErrorScript: NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
        LowestStateMachineVersion: core.CastInt64(data["lowestStateMachineVersion"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "startScript": p.StartScript.ToDict(),
        "passScript": p.PassScript.ToDict(),
        "errorScript": p.ErrorScript.ToDict(),
        "lowestStateMachineVersion": p.LowestStateMachineVersion,
        "logSetting": p.LogSetting.ToDict(),
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
    StartScript *ScriptSetting `json:"startScript"`
    PassScript *ScriptSetting `json:"passScript"`
    ErrorScript *ScriptSetting `json:"errorScript"`
    LowestStateMachineVersion *int64 `json:"lowestStateMachineVersion"`
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
        StartScript: NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
        PassScript: NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
        ErrorScript: NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
        LowestStateMachineVersion: core.CastInt64(data["lowestStateMachineVersion"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "startScript": p.StartScript.ToDict(),
        "passScript": p.PassScript.ToDict(),
        "errorScript": p.ErrorScript.ToDict(),
        "lowestStateMachineVersion": p.LowestStateMachineVersion,
        "logSetting": p.LogSetting.ToDict(),
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

type DescribeStateMachineMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeStateMachineMastersRequestFromJson(data string) DescribeStateMachineMastersRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeStateMachineMastersRequestFromDict(dict)
}

func NewDescribeStateMachineMastersRequestFromDict(data map[string]interface{}) DescribeStateMachineMastersRequest {
    return DescribeStateMachineMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeStateMachineMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeStateMachineMastersRequest) Pointer() *DescribeStateMachineMastersRequest {
    return &p
}

type UpdateStateMachineMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MainStateMachineName *string `json:"mainStateMachineName"`
    Payload *string `json:"payload"`
}

func NewUpdateStateMachineMasterRequestFromJson(data string) UpdateStateMachineMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateStateMachineMasterRequestFromDict(dict)
}

func NewUpdateStateMachineMasterRequestFromDict(data map[string]interface{}) UpdateStateMachineMasterRequest {
    return UpdateStateMachineMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MainStateMachineName: core.CastString(data["mainStateMachineName"]),
        Payload: core.CastString(data["payload"]),
    }
}

func (p UpdateStateMachineMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "mainStateMachineName": p.MainStateMachineName,
        "payload": p.Payload,
    }
}

func (p UpdateStateMachineMasterRequest) Pointer() *UpdateStateMachineMasterRequest {
    return &p
}

type GetStateMachineMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Version *int64 `json:"version"`
}

func NewGetStateMachineMasterRequestFromJson(data string) GetStateMachineMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStateMachineMasterRequestFromDict(dict)
}

func NewGetStateMachineMasterRequestFromDict(data map[string]interface{}) GetStateMachineMasterRequest {
    return GetStateMachineMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Version: core.CastInt64(data["version"]),
    }
}

func (p GetStateMachineMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "version": p.Version,
    }
}

func (p GetStateMachineMasterRequest) Pointer() *GetStateMachineMasterRequest {
    return &p
}

type DeleteStateMachineMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Version *int64 `json:"version"`
}

func NewDeleteStateMachineMasterRequestFromJson(data string) DeleteStateMachineMasterRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteStateMachineMasterRequestFromDict(dict)
}

func NewDeleteStateMachineMasterRequestFromDict(data map[string]interface{}) DeleteStateMachineMasterRequest {
    return DeleteStateMachineMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Version: core.CastInt64(data["version"]),
    }
}

func (p DeleteStateMachineMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "version": p.Version,
    }
}

func (p DeleteStateMachineMasterRequest) Pointer() *DeleteStateMachineMasterRequest {
    return &p
}

type DescribeStatusesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Status *string `json:"status"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeStatusesRequestFromJson(data string) DescribeStatusesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeStatusesRequestFromDict(dict)
}

func NewDescribeStatusesRequestFromDict(data map[string]interface{}) DescribeStatusesRequest {
    return DescribeStatusesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Status: core.CastString(data["status"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "status": p.Status,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
    return &p
}

type DescribeStatusesByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Status *string `json:"status"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeStatusesByUserIdRequestFromJson(data string) DescribeStatusesByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeStatusesByUserIdRequestFromDict(dict)
}

func NewDescribeStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeStatusesByUserIdRequest {
    return DescribeStatusesByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Status: core.CastString(data["status"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "status": p.Status,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeStatusesByUserIdRequest) Pointer() *DescribeStatusesByUserIdRequest {
    return &p
}

type GetStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    StatusName *string `json:"statusName"`
}

func NewGetStatusRequestFromJson(data string) GetStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStatusRequestFromDict(dict)
}

func NewGetStatusRequestFromDict(data map[string]interface{}) GetStatusRequest {
    return GetStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StatusName: core.CastString(data["statusName"]),
    }
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "statusName": p.StatusName,
    }
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
    return &p
}

type GetStatusByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    StatusName *string `json:"statusName"`
}

func NewGetStatusByUserIdRequestFromJson(data string) GetStatusByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetStatusByUserIdRequestFromDict(dict)
}

func NewGetStatusByUserIdRequestFromDict(data map[string]interface{}) GetStatusByUserIdRequest {
    return GetStatusByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        StatusName: core.CastString(data["statusName"]),
    }
}

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "statusName": p.StatusName,
    }
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
    return &p
}

type StartStateMachineByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Args *string `json:"args"`
    Ttl *int32 `json:"ttl"`
}

func NewStartStateMachineByUserIdRequestFromJson(data string) StartStateMachineByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStartStateMachineByUserIdRequestFromDict(dict)
}

func NewStartStateMachineByUserIdRequestFromDict(data map[string]interface{}) StartStateMachineByUserIdRequest {
    return StartStateMachineByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Args: core.CastString(data["args"]),
        Ttl: core.CastInt32(data["ttl"]),
    }
}

func (p StartStateMachineByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "args": p.Args,
        "ttl": p.Ttl,
    }
}

func (p StartStateMachineByUserIdRequest) Pointer() *StartStateMachineByUserIdRequest {
    return &p
}

type StartStateMachineByStampTaskRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewStartStateMachineByStampTaskRequestFromJson(data string) StartStateMachineByStampTaskRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStartStateMachineByStampTaskRequestFromDict(dict)
}

func NewStartStateMachineByStampTaskRequestFromDict(data map[string]interface{}) StartStateMachineByStampTaskRequest {
    return StartStateMachineByStampTaskRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p StartStateMachineByStampTaskRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p StartStateMachineByStampTaskRequest) Pointer() *StartStateMachineByStampTaskRequest {
    return &p
}

type EmitRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    StatusName *string `json:"statusName"`
    EventName *string `json:"eventName"`
    Args *string `json:"args"`
}

func NewEmitRequestFromJson(data string) EmitRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEmitRequestFromDict(dict)
}

func NewEmitRequestFromDict(data map[string]interface{}) EmitRequest {
    return EmitRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StatusName: core.CastString(data["statusName"]),
        EventName: core.CastString(data["eventName"]),
        Args: core.CastString(data["args"]),
    }
}

func (p EmitRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "statusName": p.StatusName,
        "eventName": p.EventName,
        "args": p.Args,
    }
}

func (p EmitRequest) Pointer() *EmitRequest {
    return &p
}

type EmitByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    StatusName *string `json:"statusName"`
    EventName *string `json:"eventName"`
    Args *string `json:"args"`
}

func NewEmitByUserIdRequestFromJson(data string) EmitByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEmitByUserIdRequestFromDict(dict)
}

func NewEmitByUserIdRequestFromDict(data map[string]interface{}) EmitByUserIdRequest {
    return EmitByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        StatusName: core.CastString(data["statusName"]),
        EventName: core.CastString(data["eventName"]),
        Args: core.CastString(data["args"]),
    }
}

func (p EmitByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "statusName": p.StatusName,
        "eventName": p.EventName,
        "args": p.Args,
    }
}

func (p EmitByUserIdRequest) Pointer() *EmitByUserIdRequest {
    return &p
}

type DeleteStatusByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    StatusName *string `json:"statusName"`
}

func NewDeleteStatusByUserIdRequestFromJson(data string) DeleteStatusByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteStatusByUserIdRequestFromDict(dict)
}

func NewDeleteStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteStatusByUserIdRequest {
    return DeleteStatusByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        StatusName: core.CastString(data["statusName"]),
    }
}

func (p DeleteStatusByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "statusName": p.StatusName,
    }
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
    return &p
}

type ExitStateMachineRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    StatusName *string `json:"statusName"`
}

func NewExitStateMachineRequestFromJson(data string) ExitStateMachineRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExitStateMachineRequestFromDict(dict)
}

func NewExitStateMachineRequestFromDict(data map[string]interface{}) ExitStateMachineRequest {
    return ExitStateMachineRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        StatusName: core.CastString(data["statusName"]),
    }
}

func (p ExitStateMachineRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "statusName": p.StatusName,
    }
}

func (p ExitStateMachineRequest) Pointer() *ExitStateMachineRequest {
    return &p
}

type ExitStateMachineByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    StatusName *string `json:"statusName"`
}

func NewExitStateMachineByUserIdRequestFromJson(data string) ExitStateMachineByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExitStateMachineByUserIdRequestFromDict(dict)
}

func NewExitStateMachineByUserIdRequestFromDict(data map[string]interface{}) ExitStateMachineByUserIdRequest {
    return ExitStateMachineByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        StatusName: core.CastString(data["statusName"]),
    }
}

func (p ExitStateMachineByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "statusName": p.StatusName,
    }
}

func (p ExitStateMachineByUserIdRequest) Pointer() *ExitStateMachineByUserIdRequest {
    return &p
}