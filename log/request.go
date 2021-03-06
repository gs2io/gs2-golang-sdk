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

package log

import "core"

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
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
    Type *string `json:"type"`
    GcpCredentialJson *string `json:"gcpCredentialJson"`
    BigQueryDatasetName *string `json:"bigQueryDatasetName"`
    LogExpireDays *int32 `json:"logExpireDays"`
    AwsRegion *string `json:"awsRegion"`
    AwsAccessKeyId *string `json:"awsAccessKeyId"`
    AwsSecretAccessKey *string `json:"awsSecretAccessKey"`
    FirehoseStreamName *string `json:"firehoseStreamName"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Type: core.CastString(data["type"]),
        GcpCredentialJson: core.CastString(data["gcpCredentialJson"]),
        BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
        LogExpireDays: core.CastInt32(data["logExpireDays"]),
        AwsRegion: core.CastString(data["awsRegion"]),
        AwsAccessKeyId: core.CastString(data["awsAccessKeyId"]),
        AwsSecretAccessKey: core.CastString(data["awsSecretAccessKey"]),
        FirehoseStreamName: core.CastString(data["firehoseStreamName"]),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "type": p.Type,
        "gcpCredentialJson": p.GcpCredentialJson,
        "bigQueryDatasetName": p.BigQueryDatasetName,
        "logExpireDays": p.LogExpireDays,
        "awsRegion": p.AwsRegion,
        "awsAccessKeyId": p.AwsAccessKeyId,
        "awsSecretAccessKey": p.AwsSecretAccessKey,
        "firehoseStreamName": p.FirehoseStreamName,
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
    Type *string `json:"type"`
    GcpCredentialJson *string `json:"gcpCredentialJson"`
    BigQueryDatasetName *string `json:"bigQueryDatasetName"`
    LogExpireDays *int32 `json:"logExpireDays"`
    AwsRegion *string `json:"awsRegion"`
    AwsAccessKeyId *string `json:"awsAccessKeyId"`
    AwsSecretAccessKey *string `json:"awsSecretAccessKey"`
    FirehoseStreamName *string `json:"firehoseStreamName"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        Type: core.CastString(data["type"]),
        GcpCredentialJson: core.CastString(data["gcpCredentialJson"]),
        BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
        LogExpireDays: core.CastInt32(data["logExpireDays"]),
        AwsRegion: core.CastString(data["awsRegion"]),
        AwsAccessKeyId: core.CastString(data["awsAccessKeyId"]),
        AwsSecretAccessKey: core.CastString(data["awsSecretAccessKey"]),
        FirehoseStreamName: core.CastString(data["firehoseStreamName"]),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "type": p.Type,
        "gcpCredentialJson": p.GcpCredentialJson,
        "bigQueryDatasetName": p.BigQueryDatasetName,
        "logExpireDays": p.LogExpireDays,
        "awsRegion": p.AwsRegion,
        "awsAccessKeyId": p.AwsAccessKeyId,
        "awsSecretAccessKey": p.AwsSecretAccessKey,
        "firehoseStreamName": p.FirehoseStreamName,
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

type QueryAccessLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewQueryAccessLogRequestFromDict(data map[string]interface{}) QueryAccessLogRequest {
    return QueryAccessLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p QueryAccessLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p QueryAccessLogRequest) Pointer() *QueryAccessLogRequest {
    return &p
}

type CountAccessLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewCountAccessLogRequestFromDict(data map[string]interface{}) CountAccessLogRequest {
    return CountAccessLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p CountAccessLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p CountAccessLogRequest) Pointer() *CountAccessLogRequest {
    return &p
}

type QueryIssueStampSheetLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewQueryIssueStampSheetLogRequestFromDict(data map[string]interface{}) QueryIssueStampSheetLogRequest {
    return QueryIssueStampSheetLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p QueryIssueStampSheetLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p QueryIssueStampSheetLogRequest) Pointer() *QueryIssueStampSheetLogRequest {
    return &p
}

type CountIssueStampSheetLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewCountIssueStampSheetLogRequestFromDict(data map[string]interface{}) CountIssueStampSheetLogRequest {
    return CountIssueStampSheetLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p CountIssueStampSheetLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p CountIssueStampSheetLogRequest) Pointer() *CountIssueStampSheetLogRequest {
    return &p
}

type QueryExecuteStampSheetLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewQueryExecuteStampSheetLogRequestFromDict(data map[string]interface{}) QueryExecuteStampSheetLogRequest {
    return QueryExecuteStampSheetLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p QueryExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p QueryExecuteStampSheetLogRequest) Pointer() *QueryExecuteStampSheetLogRequest {
    return &p
}

type CountExecuteStampSheetLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewCountExecuteStampSheetLogRequestFromDict(data map[string]interface{}) CountExecuteStampSheetLogRequest {
    return CountExecuteStampSheetLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p CountExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p CountExecuteStampSheetLogRequest) Pointer() *CountExecuteStampSheetLogRequest {
    return &p
}

type QueryExecuteStampTaskLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewQueryExecuteStampTaskLogRequestFromDict(data map[string]interface{}) QueryExecuteStampTaskLogRequest {
    return QueryExecuteStampTaskLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p QueryExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p QueryExecuteStampTaskLogRequest) Pointer() *QueryExecuteStampTaskLogRequest {
    return &p
}

type CountExecuteStampTaskLogRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Service *string `json:"service"`
    Method *string `json:"method"`
    UserId *string `json:"userId"`
    Action *string `json:"action"`
    Begin *int64 `json:"begin"`
    End *int64 `json:"end"`
    LongTerm *bool `json:"longTerm"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewCountExecuteStampTaskLogRequestFromDict(data map[string]interface{}) CountExecuteStampTaskLogRequest {
    return CountExecuteStampTaskLogRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Begin: core.CastInt64(data["begin"]),
        End: core.CastInt64(data["end"]),
        LongTerm: core.CastBool(data["longTerm"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p CountExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "service": p.Service,
        "method": p.Method,
        "userId": p.UserId,
        "action": p.Action,
        "begin": p.Begin,
        "end": p.End,
        "longTerm": p.LongTerm,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p CountExecuteStampTaskLogRequest) Pointer() *CountExecuteStampTaskLogRequest {
    return &p
}