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

package gateway

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	FirebaseSecret *string `json:"firebaseSecret"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
    return Namespace {
        NamespaceId: core.CastString(data["namespaceId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        FirebaseSecret: core.CastString(data["firebaseSecret"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Namespace) ToDict() map[string]interface{} {
    
    var namespaceId *string
    if p.NamespaceId != nil {
        namespaceId = p.NamespaceId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var firebaseSecret *string
    if p.FirebaseSecret != nil {
        firebaseSecret = p.FirebaseSecret
    }
    var logSetting map[string]interface{}
    if p.LogSetting != nil {
        logSetting = p.LogSetting.ToDict()
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "firebaseSecret": firebaseSecret,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Namespace) Pointer() *Namespace {
    return &p
}

func CastNamespaces(data []interface{}) []Namespace {
	v := make([]Namespace, 0)
	for _, d := range data {
		v = append(v, NewNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNamespacesFromDict(data []Namespace) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type WebSocketSession struct {
	WebSocketSessionId *string `json:"webSocketSessionId"`
	ConnectionId *string `json:"connectionId"`
	NamespaceName *string `json:"namespaceName"`
	UserId *string `json:"userId"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewWebSocketSessionFromJson(data string) WebSocketSession {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewWebSocketSessionFromDict(dict)
}

func NewWebSocketSessionFromDict(data map[string]interface{}) WebSocketSession {
    return WebSocketSession {
        WebSocketSessionId: core.CastString(data["webSocketSessionId"]),
        ConnectionId: core.CastString(data["connectionId"]),
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p WebSocketSession) ToDict() map[string]interface{} {
    
    var webSocketSessionId *string
    if p.WebSocketSessionId != nil {
        webSocketSessionId = p.WebSocketSessionId
    }
    var connectionId *string
    if p.ConnectionId != nil {
        connectionId = p.ConnectionId
    }
    var namespaceName *string
    if p.NamespaceName != nil {
        namespaceName = p.NamespaceName
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "webSocketSessionId": webSocketSessionId,
        "connectionId": connectionId,
        "namespaceName": namespaceName,
        "userId": userId,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p WebSocketSession) Pointer() *WebSocketSession {
    return &p
}

func CastWebSocketSessions(data []interface{}) []WebSocketSession {
	v := make([]WebSocketSession, 0)
	for _, d := range data {
		v = append(v, NewWebSocketSessionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWebSocketSessionsFromDict(data []WebSocketSession) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type FirebaseToken struct {
	FirebaseTokenId *string `json:"firebaseTokenId"`
	UserId *string `json:"userId"`
	Token *string `json:"token"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewFirebaseTokenFromJson(data string) FirebaseToken {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewFirebaseTokenFromDict(dict)
}

func NewFirebaseTokenFromDict(data map[string]interface{}) FirebaseToken {
    return FirebaseToken {
        FirebaseTokenId: core.CastString(data["firebaseTokenId"]),
        UserId: core.CastString(data["userId"]),
        Token: core.CastString(data["token"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p FirebaseToken) ToDict() map[string]interface{} {
    
    var firebaseTokenId *string
    if p.FirebaseTokenId != nil {
        firebaseTokenId = p.FirebaseTokenId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var token *string
    if p.Token != nil {
        token = p.Token
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "firebaseTokenId": firebaseTokenId,
        "userId": userId,
        "token": token,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p FirebaseToken) Pointer() *FirebaseToken {
    return &p
}

func CastFirebaseTokens(data []interface{}) []FirebaseToken {
	v := make([]FirebaseToken, 0)
	for _, d := range data {
		v = append(v, NewFirebaseTokenFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFirebaseTokensFromDict(data []FirebaseToken) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func NewLogSettingFromJson(data string) LogSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    
    var loggingNamespaceId *string
    if p.LoggingNamespaceId != nil {
        loggingNamespaceId = p.LoggingNamespaceId
    }
    return map[string]interface{} {
        "loggingNamespaceId": loggingNamespaceId,
    }
}

func (p LogSetting) Pointer() *LogSetting {
    return &p
}

func CastLogSettings(data []interface{}) []LogSetting {
	v := make([]LogSetting, 0)
	for _, d := range data {
		v = append(v, NewLogSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogSettingsFromDict(data []LogSetting) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}