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
    return map[string]interface{} {
        "namespaceId": p.NamespaceId,
        "name": p.Name,
        "description": p.Description,
        "firebaseSecret": p.FirebaseSecret,
        "logSetting": p.LogSetting.ToDict(),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
    return map[string]interface{} {
        "webSocketSessionId": p.WebSocketSessionId,
        "connectionId": p.ConnectionId,
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
    return map[string]interface{} {
        "firebaseTokenId": p.FirebaseTokenId,
        "userId": p.UserId,
        "token": p.Token,
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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
    return map[string]interface{} {
        "loggingNamespaceId": p.LoggingNamespaceId,
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