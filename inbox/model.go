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

package inbox

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                *string              `json:"namespaceId"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	IsAutomaticDeletingEnabled *bool                `json:"isAutomaticDeletingEnabled"`
	TransactionSetting         *TransactionSetting  `json:"transactionSetting"`
	ReceiveMessageScript       *ScriptSetting       `json:"receiveMessageScript"`
	ReadMessageScript          *ScriptSetting       `json:"readMessageScript"`
	DeleteMessageScript        *ScriptSetting       `json:"deleteMessageScript"`
	ReceiveNotification        *NotificationSetting `json:"receiveNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	CreatedAt                  *int64               `json:"createdAt"`
	UpdatedAt                  *int64               `json:"updatedAt"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId    *string `json:"keyId"`
	Revision *int64  `json:"revision"`
}

func (p *Namespace) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Namespace{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Namespace{}
	} else {
		*p = Namespace{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["isAutomaticDeletingEnabled"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IsAutomaticDeletingEnabled)
		}
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["receiveMessageScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveMessageScript)
		}
		if v, ok := d["readMessageScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReadMessageScript)
		}
		if v, ok := d["deleteMessageScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DeleteMessageScript)
		}
		if v, ok := d["receiveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveNotification)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["queueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QueueNamespaceId)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewNamespaceFromJson(data string) Namespace {
	req := Namespace{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		IsAutomaticDeletingEnabled: func() *bool {
			v, ok := data["isAutomaticDeletingEnabled"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isAutomaticDeletingEnabled"])
		}(),
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		ReceiveMessageScript: func() *ScriptSetting {
			v, ok := data["receiveMessageScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["receiveMessageScript"])).Pointer()
		}(),
		ReadMessageScript: func() *ScriptSetting {
			v, ok := data["readMessageScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["readMessageScript"])).Pointer()
		}(),
		DeleteMessageScript: func() *ScriptSetting {
			v, ok := data["deleteMessageScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["deleteMessageScript"])).Pointer()
		}(),
		ReceiveNotification: func() *NotificationSetting {
			v, ok := data["receiveNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["receiveNotification"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.IsAutomaticDeletingEnabled != nil {
		m["isAutomaticDeletingEnabled"] = p.IsAutomaticDeletingEnabled
	}
	if p.TransactionSetting != nil {
		m["transactionSetting"] = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	if p.ReceiveMessageScript != nil {
		m["receiveMessageScript"] = func() map[string]interface{} {
			if p.ReceiveMessageScript == nil {
				return nil
			}
			return p.ReceiveMessageScript.ToDict()
		}()
	}
	if p.ReadMessageScript != nil {
		m["readMessageScript"] = func() map[string]interface{} {
			if p.ReadMessageScript == nil {
				return nil
			}
			return p.ReadMessageScript.ToDict()
		}()
	}
	if p.DeleteMessageScript != nil {
		m["deleteMessageScript"] = func() map[string]interface{} {
			if p.DeleteMessageScript == nil {
				return nil
			}
			return p.DeleteMessageScript.ToDict()
		}()
	}
	if p.ReceiveNotification != nil {
		m["receiveNotification"] = func() map[string]interface{} {
			if p.ReceiveNotification == nil {
				return nil
			}
			return p.ReceiveNotification.ToDict()
		}()
	}
	if p.LogSetting != nil {
		m["logSetting"] = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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

type Message struct {
	MessageId          *string         `json:"messageId"`
	Name               *string         `json:"name"`
	UserId             *string         `json:"userId"`
	Metadata           *string         `json:"metadata"`
	IsRead             *bool           `json:"isRead"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ReceivedAt         *int64          `json:"receivedAt"`
	ReadAt             *int64          `json:"readAt"`
	ExpiresAt          *int64          `json:"expiresAt"`
	Revision           *int64          `json:"revision"`
}

func (p *Message) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Message{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Message{}
	} else {
		*p = Message{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["messageId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MessageId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MessageId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MessageId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MessageId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MessageId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MessageId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["isRead"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IsRead)
		}
		if v, ok := d["readAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReadAcquireActions)
		}
		if v, ok := d["receivedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceivedAt)
		}
		if v, ok := d["readAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReadAt)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewMessageFromJson(data string) Message {
	req := Message{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMessageFromDict(data map[string]interface{}) Message {
	return Message{
		MessageId: func() *string {
			v, ok := data["messageId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["messageId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		IsRead: func() *bool {
			v, ok := data["isRead"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isRead"])
		}(),
		ReadAcquireActions: func() []AcquireAction {
			if data["readAcquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["readAcquireActions"]))
		}(),
		ReceivedAt: func() *int64 {
			v, ok := data["receivedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["receivedAt"])
		}(),
		ReadAt: func() *int64 {
			v, ok := data["readAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["readAt"])
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Message) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.MessageId != nil {
		m["messageId"] = p.MessageId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.IsRead != nil {
		m["isRead"] = p.IsRead
	}
	if p.ReadAcquireActions != nil {
		m["readAcquireActions"] = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	if p.ReceivedAt != nil {
		m["receivedAt"] = p.ReceivedAt
	}
	if p.ReadAt != nil {
		m["readAt"] = p.ReadAt
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p Message) Pointer() *Message {
	return &p
}

func CastMessages(data []interface{}) []Message {
	v := make([]Message, 0)
	for _, d := range data {
		v = append(v, NewMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMessagesFromDict(data []Message) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentMessageMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentMessageMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentMessageMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentMessageMaster{}
	} else {
		*p = CurrentMessageMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
				}
			}
		}
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Settings = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Settings = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Settings = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Settings)
				}
			}
		}
	}
	return nil
}

func NewCurrentMessageMasterFromJson(data string) CurrentMessageMaster {
	req := CurrentMessageMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentMessageMasterFromDict(data map[string]interface{}) CurrentMessageMaster {
	return CurrentMessageMaster{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
	}
}

func (p CurrentMessageMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentMessageMaster) Pointer() *CurrentMessageMaster {
	return &p
}

func CastCurrentMessageMasters(data []interface{}) []CurrentMessageMaster {
	v := make([]CurrentMessageMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentMessageMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentMessageMastersFromDict(data []CurrentMessageMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalMessageMaster struct {
	GlobalMessageId    *string         `json:"globalMessageId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	// Deprecated: should not be used
	ExpiresAt                     *int64  `json:"expiresAt"`
	MessageReceptionPeriodEventId *string `json:"messageReceptionPeriodEventId"`
	CreatedAt                     *int64  `json:"createdAt"`
	Revision                      *int64  `json:"revision"`
}

func (p *GlobalMessageMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalMessageMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GlobalMessageMaster{}
	} else {
		*p = GlobalMessageMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalMessageId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalMessageId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalMessageId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalMessageId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalMessageId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalMessageId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalMessageId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["readAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReadAcquireActions)
		}
		if v, ok := d["expiresTimeSpan"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresTimeSpan)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["messageReceptionPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MessageReceptionPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MessageReceptionPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MessageReceptionPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MessageReceptionPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MessageReceptionPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MessageReceptionPeriodEventId)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewGlobalMessageMasterFromJson(data string) GlobalMessageMaster {
	req := GlobalMessageMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalMessageMasterFromDict(data map[string]interface{}) GlobalMessageMaster {
	return GlobalMessageMaster{
		GlobalMessageId: func() *string {
			v, ok := data["globalMessageId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalMessageId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		ReadAcquireActions: func() []AcquireAction {
			if data["readAcquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["readAcquireActions"]))
		}(),
		ExpiresTimeSpan: func() *TimeSpan {
			v, ok := data["expiresTimeSpan"]
			if !ok || v == nil {
				return nil
			}
			return NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer()
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		MessageReceptionPeriodEventId: func() *string {
			v, ok := data["messageReceptionPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["messageReceptionPeriodEventId"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p GlobalMessageMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalMessageId != nil {
		m["globalMessageId"] = p.GlobalMessageId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ReadAcquireActions != nil {
		m["readAcquireActions"] = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	if p.ExpiresTimeSpan != nil {
		m["expiresTimeSpan"] = func() map[string]interface{} {
			if p.ExpiresTimeSpan == nil {
				return nil
			}
			return p.ExpiresTimeSpan.ToDict()
		}()
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
	}
	if p.MessageReceptionPeriodEventId != nil {
		m["messageReceptionPeriodEventId"] = p.MessageReceptionPeriodEventId
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p GlobalMessageMaster) Pointer() *GlobalMessageMaster {
	return &p
}

func CastGlobalMessageMasters(data []interface{}) []GlobalMessageMaster {
	v := make([]GlobalMessageMaster, 0)
	for _, d := range data {
		v = append(v, NewGlobalMessageMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalMessageMastersFromDict(data []GlobalMessageMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalMessage struct {
	GlobalMessageId    *string         `json:"globalMessageId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ReadAcquireActions []AcquireAction `json:"readAcquireActions"`
	ExpiresTimeSpan    *TimeSpan       `json:"expiresTimeSpan"`
	// Deprecated: should not be used
	ExpiresAt                     *int64  `json:"expiresAt"`
	MessageReceptionPeriodEventId *string `json:"messageReceptionPeriodEventId"`
}

func (p *GlobalMessage) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalMessage{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GlobalMessage{}
	} else {
		*p = GlobalMessage{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalMessageId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalMessageId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalMessageId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalMessageId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalMessageId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalMessageId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalMessageId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["readAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReadAcquireActions)
		}
		if v, ok := d["expiresTimeSpan"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresTimeSpan)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["messageReceptionPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MessageReceptionPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MessageReceptionPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MessageReceptionPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MessageReceptionPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MessageReceptionPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MessageReceptionPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewGlobalMessageFromJson(data string) GlobalMessage {
	req := GlobalMessage{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalMessageFromDict(data map[string]interface{}) GlobalMessage {
	return GlobalMessage{
		GlobalMessageId: func() *string {
			v, ok := data["globalMessageId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalMessageId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		ReadAcquireActions: func() []AcquireAction {
			if data["readAcquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["readAcquireActions"]))
		}(),
		ExpiresTimeSpan: func() *TimeSpan {
			v, ok := data["expiresTimeSpan"]
			if !ok || v == nil {
				return nil
			}
			return NewTimeSpanFromDict(core.CastMap(data["expiresTimeSpan"])).Pointer()
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
		}(),
		MessageReceptionPeriodEventId: func() *string {
			v, ok := data["messageReceptionPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["messageReceptionPeriodEventId"])
		}(),
	}
}

func (p GlobalMessage) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalMessageId != nil {
		m["globalMessageId"] = p.GlobalMessageId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ReadAcquireActions != nil {
		m["readAcquireActions"] = CastAcquireActionsFromDict(
			p.ReadAcquireActions,
		)
	}
	if p.ExpiresTimeSpan != nil {
		m["expiresTimeSpan"] = func() map[string]interface{} {
			if p.ExpiresTimeSpan == nil {
				return nil
			}
			return p.ExpiresTimeSpan.ToDict()
		}()
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
	}
	if p.MessageReceptionPeriodEventId != nil {
		m["messageReceptionPeriodEventId"] = p.MessageReceptionPeriodEventId
	}
	return m
}

func (p GlobalMessage) Pointer() *GlobalMessage {
	return &p
}

func CastGlobalMessages(data []interface{}) []GlobalMessage {
	v := make([]GlobalMessage, 0)
	for _, d := range data {
		v = append(v, NewGlobalMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalMessagesFromDict(data []GlobalMessage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Received struct {
	ReceivedId                 *string   `json:"receivedId"`
	UserId                     *string   `json:"userId"`
	ReceivedGlobalMessageNames []*string `json:"receivedGlobalMessageNames"`
	CreatedAt                  *int64    `json:"createdAt"`
	UpdatedAt                  *int64    `json:"updatedAt"`
	Revision                   *int64    `json:"revision"`
}

func (p *Received) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Received{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Received{}
	} else {
		*p = Received{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["receivedId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReceivedId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReceivedId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReceivedId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReceivedId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReceivedId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReceivedId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["receivedGlobalMessageNames"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.ReceivedGlobalMessageNames = l
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewReceivedFromJson(data string) Received {
	req := Received{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewReceivedFromDict(data map[string]interface{}) Received {
	return Received{
		ReceivedId: func() *string {
			v, ok := data["receivedId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["receivedId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ReceivedGlobalMessageNames: func() []*string {
			v, ok := data["receivedGlobalMessageNames"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Received) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ReceivedId != nil {
		m["receivedId"] = p.ReceivedId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ReceivedGlobalMessageNames != nil {
		m["receivedGlobalMessageNames"] = core.CastStringsFromDict(
			p.ReceivedGlobalMessageNames,
		)
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p Received) Pointer() *Received {
	return &p
}

func CastReceiveds(data []interface{}) []Received {
	v := make([]Received, 0)
	for _, d := range data {
		v = append(v, NewReceivedFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceivedsFromDict(data []Received) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TimeSpan struct {
	Days    *int32 `json:"days"`
	Hours   *int32 `json:"hours"`
	Minutes *int32 `json:"minutes"`
}

func (p *TimeSpan) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TimeSpan{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TimeSpan{}
	} else {
		*p = TimeSpan{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["days"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Days)
		}
		if v, ok := d["hours"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Hours)
		}
		if v, ok := d["minutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Minutes)
		}
	}
	return nil
}

func NewTimeSpanFromJson(data string) TimeSpan {
	req := TimeSpan{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTimeSpanFromDict(data map[string]interface{}) TimeSpan {
	return TimeSpan{
		Days: func() *int32 {
			v, ok := data["days"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["days"])
		}(),
		Hours: func() *int32 {
			v, ok := data["hours"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["hours"])
		}(),
		Minutes: func() *int32 {
			v, ok := data["minutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["minutes"])
		}(),
	}
}

func (p TimeSpan) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Days != nil {
		m["days"] = p.Days
	}
	if p.Hours != nil {
		m["hours"] = p.Hours
	}
	if p.Minutes != nil {
		m["minutes"] = p.Minutes
	}
	return m
}

func (p TimeSpan) Pointer() *TimeSpan {
	return &p
}

func CastTimeSpans(data []interface{}) []TimeSpan {
	v := make([]TimeSpan, 0)
	for _, d := range data {
		v = append(v, NewTimeSpanFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeSpansFromDict(data []TimeSpan) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func (p *AcquireAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireAction{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcquireAction{}
	} else {
		*p = AcquireAction{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
		}
	}
	return nil
}

func NewAcquireActionFromJson(data string) AcquireAction {
	req := AcquireAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Request: func() *string {
			v, ok := data["request"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["request"])
		}(),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Request != nil {
		m["request"] = p.Request
	}
	return m
}

func (p AcquireAction) Pointer() *AcquireAction {
	return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Config struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *Config) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Config{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Config{}
	} else {
		*p = Config{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
	}
	return nil
}

func NewConfigFromJson(data string) Config {
	req := Config{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
	}
}

func (p Config) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p Config) Pointer() *Config {
	return &p
}

func CastConfigs(data []interface{}) []Config {
	v := make([]Config, 0)
	for _, d := range data {
		v = append(v, NewConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConfigsFromDict(data []Config) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VerifyActionResult struct {
	Action        *string `json:"action"`
	VerifyRequest *string `json:"verifyRequest"`
	StatusCode    *int32  `json:"statusCode"`
	VerifyResult  *string `json:"verifyResult"`
}

func (p *VerifyActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyActionResult{}
	} else {
		*p = VerifyActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["verifyRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["verifyResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyResult)
				}
			}
		}
	}
	return nil
}

func NewVerifyActionResultFromJson(data string) VerifyActionResult {
	req := VerifyActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVerifyActionResultFromDict(data map[string]interface{}) VerifyActionResult {
	return VerifyActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		VerifyRequest: func() *string {
			v, ok := data["verifyRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		VerifyResult: func() *string {
			v, ok := data["verifyResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyResult"])
		}(),
	}
}

func (p VerifyActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.VerifyRequest != nil {
		m["verifyRequest"] = p.VerifyRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.VerifyResult != nil {
		m["verifyResult"] = p.VerifyResult
	}
	return m
}

func (p VerifyActionResult) Pointer() *VerifyActionResult {
	return &p
}

func CastVerifyActionResults(data []interface{}) []VerifyActionResult {
	v := make([]VerifyActionResult, 0)
	for _, d := range data {
		v = append(v, NewVerifyActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVerifyActionResultsFromDict(data []VerifyActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeActionResult struct {
	Action         *string `json:"action"`
	ConsumeRequest *string `json:"consumeRequest"`
	StatusCode     *int32  `json:"statusCode"`
	ConsumeResult  *string `json:"consumeResult"`
}

func (p *ConsumeActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeActionResult{}
	} else {
		*p = ConsumeActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["consumeRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["consumeResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeResult)
				}
			}
		}
	}
	return nil
}

func NewConsumeActionResultFromJson(data string) ConsumeActionResult {
	req := ConsumeActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeActionResultFromDict(data map[string]interface{}) ConsumeActionResult {
	return ConsumeActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		ConsumeRequest: func() *string {
			v, ok := data["consumeRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		ConsumeResult: func() *string {
			v, ok := data["consumeResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeResult"])
		}(),
	}
}

func (p ConsumeActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.ConsumeRequest != nil {
		m["consumeRequest"] = p.ConsumeRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.ConsumeResult != nil {
		m["consumeResult"] = p.ConsumeResult
	}
	return m
}

func (p ConsumeActionResult) Pointer() *ConsumeActionResult {
	return &p
}

func CastConsumeActionResults(data []interface{}) []ConsumeActionResult {
	v := make([]ConsumeActionResult, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionResultsFromDict(data []ConsumeActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireActionResult struct {
	Action         *string `json:"action"`
	AcquireRequest *string `json:"acquireRequest"`
	StatusCode     *int32  `json:"statusCode"`
	AcquireResult  *string `json:"acquireResult"`
}

func (p *AcquireActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcquireActionResult{}
	} else {
		*p = AcquireActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["acquireRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["acquireResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireResult)
				}
			}
		}
	}
	return nil
}

func NewAcquireActionResultFromJson(data string) AcquireActionResult {
	req := AcquireActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireActionResultFromDict(data map[string]interface{}) AcquireActionResult {
	return AcquireActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		AcquireRequest: func() *string {
			v, ok := data["acquireRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		AcquireResult: func() *string {
			v, ok := data["acquireResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireResult"])
		}(),
	}
}

func (p AcquireActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.AcquireRequest != nil {
		m["acquireRequest"] = p.AcquireRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.AcquireResult != nil {
		m["acquireResult"] = p.AcquireResult
	}
	return m
}

func (p AcquireActionResult) Pointer() *AcquireActionResult {
	return &p
}

func CastAcquireActionResults(data []interface{}) []AcquireActionResult {
	v := make([]AcquireActionResult, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionResultsFromDict(data []AcquireActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionResult struct {
	TransactionId  *string               `json:"transactionId"`
	VerifyResults  []VerifyActionResult  `json:"verifyResults"`
	ConsumeResults []ConsumeActionResult `json:"consumeResults"`
	AcquireResults []AcquireActionResult `json:"acquireResults"`
	HasError       *bool                 `json:"hasError"`
}

func (p *TransactionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TransactionResult{}
	} else {
		*p = TransactionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["verifyResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyResults)
		}
		if v, ok := d["consumeResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeResults)
		}
		if v, ok := d["acquireResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireResults)
		}
		if v, ok := d["hasError"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.HasError)
		}
	}
	return nil
}

func NewTransactionResultFromJson(data string) TransactionResult {
	req := TransactionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionResultFromDict(data map[string]interface{}) TransactionResult {
	return TransactionResult{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		VerifyResults: func() []VerifyActionResult {
			if data["verifyResults"] == nil {
				return nil
			}
			return CastVerifyActionResults(core.CastArray(data["verifyResults"]))
		}(),
		ConsumeResults: func() []ConsumeActionResult {
			if data["consumeResults"] == nil {
				return nil
			}
			return CastConsumeActionResults(core.CastArray(data["consumeResults"]))
		}(),
		AcquireResults: func() []AcquireActionResult {
			if data["acquireResults"] == nil {
				return nil
			}
			return CastAcquireActionResults(core.CastArray(data["acquireResults"]))
		}(),
		HasError: func() *bool {
			v, ok := data["hasError"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["hasError"])
		}(),
	}
}

func (p TransactionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.VerifyResults != nil {
		m["verifyResults"] = CastVerifyActionResultsFromDict(
			p.VerifyResults,
		)
	}
	if p.ConsumeResults != nil {
		m["consumeResults"] = CastConsumeActionResultsFromDict(
			p.ConsumeResults,
		)
	}
	if p.AcquireResults != nil {
		m["acquireResults"] = CastAcquireActionResultsFromDict(
			p.AcquireResults,
		)
	}
	if p.HasError != nil {
		m["hasError"] = p.HasError
	}
	return m
}

func (p TransactionResult) Pointer() *TransactionResult {
	return &p
}

func CastTransactionResults(data []interface{}) []TransactionResult {
	v := make([]TransactionResult, 0)
	for _, d := range data {
		v = append(v, NewTransactionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionResultsFromDict(data []TransactionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScriptSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ScriptSetting{}
	} else {
		*p = ScriptSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["triggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerTargetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerTargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerTargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerTargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerTargetType)
				}
			}
		}
		if v, ok := d["doneTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerQueueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerQueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerQueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	req := ScriptSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId: func() *string {
			v, ok := data["triggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerScriptId"])
		}(),
		DoneTriggerTargetType: func() *string {
			v, ok := data["doneTriggerTargetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerTargetType"])
		}(),
		DoneTriggerScriptId: func() *string {
			v, ok := data["doneTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerScriptId"])
		}(),
		DoneTriggerQueueNamespaceId: func() *string {
			v, ok := data["doneTriggerQueueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerQueueNamespaceId"])
		}(),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		m["triggerScriptId"] = p.TriggerScriptId
	}
	if p.DoneTriggerTargetType != nil {
		m["doneTriggerTargetType"] = p.DoneTriggerTargetType
	}
	if p.DoneTriggerScriptId != nil {
		m["doneTriggerScriptId"] = p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		m["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
	}
	return m
}

func (p ScriptSetting) Pointer() *ScriptSetting {
	return &p
}

func CastScriptSettings(data []interface{}) []ScriptSetting {
	v := make([]ScriptSetting, 0)
	for _, d := range data {
		v = append(v, NewScriptSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScriptSettingsFromDict(data []ScriptSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NotificationSetting struct {
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func (p *NotificationSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NotificationSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = NotificationSetting{}
	} else {
		*p = NotificationSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["gatewayNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatewayNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatewayNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatewayNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatewayNamespaceId)
				}
			}
		}
		if v, ok := d["enableTransferMobileNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableTransferMobileNotification)
		}
		if v, ok := d["sound"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Sound = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Sound = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Sound = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Sound)
				}
			}
		}
	}
	return nil
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	req := NotificationSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId: func() *string {
			v, ok := data["gatewayNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatewayNamespaceId"])
		}(),
		EnableTransferMobileNotification: func() *bool {
			v, ok := data["enableTransferMobileNotification"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableTransferMobileNotification"])
		}(),
		Sound: func() *string {
			v, ok := data["sound"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sound"])
		}(),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GatewayNamespaceId != nil {
		m["gatewayNamespaceId"] = p.GatewayNamespaceId
	}
	if p.EnableTransferMobileNotification != nil {
		m["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
	}
	if p.Sound != nil {
		m["sound"] = p.Sound
	}
	return m
}

func (p NotificationSetting) Pointer() *NotificationSetting {
	return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func (p *GitHubCheckoutSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GitHubCheckoutSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GitHubCheckoutSetting{}
	} else {
		*p = GitHubCheckoutSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["apiKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApiKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApiKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApiKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApiKeyId)
				}
			}
		}
		if v, ok := d["repositoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepositoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepositoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepositoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepositoryName)
				}
			}
		}
		if v, ok := d["sourcePath"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourcePath = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourcePath = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourcePath = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourcePath)
				}
			}
		}
		if v, ok := d["referenceType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceType)
				}
			}
		}
		if v, ok := d["commitHash"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CommitHash = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CommitHash = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CommitHash = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CommitHash)
				}
			}
		}
		if v, ok := d["branchName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BranchName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BranchName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BranchName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BranchName)
				}
			}
		}
		if v, ok := d["tagName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TagName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TagName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TagName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TagName)
				}
			}
		}
	}
	return nil
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	req := GitHubCheckoutSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId: func() *string {
			v, ok := data["apiKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["apiKeyId"])
		}(),
		RepositoryName: func() *string {
			v, ok := data["repositoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repositoryName"])
		}(),
		SourcePath: func() *string {
			v, ok := data["sourcePath"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourcePath"])
		}(),
		ReferenceType: func() *string {
			v, ok := data["referenceType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["referenceType"])
		}(),
		CommitHash: func() *string {
			v, ok := data["commitHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["commitHash"])
		}(),
		BranchName: func() *string {
			v, ok := data["branchName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["branchName"])
		}(),
		TagName: func() *string {
			v, ok := data["tagName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["tagName"])
		}(),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ApiKeyId != nil {
		m["apiKeyId"] = p.ApiKeyId
	}
	if p.RepositoryName != nil {
		m["repositoryName"] = p.RepositoryName
	}
	if p.SourcePath != nil {
		m["sourcePath"] = p.SourcePath
	}
	if p.ReferenceType != nil {
		m["referenceType"] = p.ReferenceType
	}
	if p.CommitHash != nil {
		m["commitHash"] = p.CommitHash
	}
	if p.BranchName != nil {
		m["branchName"] = p.BranchName
	}
	if p.TagName != nil {
		m["tagName"] = p.TagName
	}
	return m
}

func (p GitHubCheckoutSetting) Pointer() *GitHubCheckoutSetting {
	return &p
}

func CastGitHubCheckoutSettings(data []interface{}) []GitHubCheckoutSetting {
	v := make([]GitHubCheckoutSetting, 0)
	for _, d := range data {
		v = append(v, NewGitHubCheckoutSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubCheckoutSettingsFromDict(data []GitHubCheckoutSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LogSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LogSetting{}
	} else {
		*p = LogSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["loggingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LoggingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LoggingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LoggingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LoggingNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewLogSettingFromJson(data string) LogSetting {
	req := LogSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: func() *string {
			v, ok := data["loggingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["loggingNamespaceId"])
		}(),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LoggingNamespaceId != nil {
		m["loggingNamespaceId"] = p.LoggingNamespaceId
	}
	return m
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

type TransactionSetting struct {
	EnableAutoRun             *bool   `json:"enableAutoRun"`
	EnableAtomicCommit        *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor *bool   `json:"transactionUseDistributor"`
	AcquireActionUseJobQueue  *bool   `json:"acquireActionUseJobQueue"`
	DistributorNamespaceId    *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func (p *TransactionSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TransactionSetting{}
	} else {
		*p = TransactionSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["enableAutoRun"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAutoRun)
		}
		if v, ok := d["enableAtomicCommit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAtomicCommit)
		}
		if v, ok := d["transactionUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionUseDistributor)
		}
		if v, ok := d["acquireActionUseJobQueue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionUseJobQueue)
		}
		if v, ok := d["distributorNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DistributorNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DistributorNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DistributorNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DistributorNamespaceId)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
		if v, ok := d["queueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	req := TransactionSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun: func() *bool {
			v, ok := data["enableAutoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAutoRun"])
		}(),
		EnableAtomicCommit: func() *bool {
			v, ok := data["enableAtomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAtomicCommit"])
		}(),
		TransactionUseDistributor: func() *bool {
			v, ok := data["transactionUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["transactionUseDistributor"])
		}(),
		AcquireActionUseJobQueue: func() *bool {
			v, ok := data["acquireActionUseJobQueue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["acquireActionUseJobQueue"])
		}(),
		DistributorNamespaceId: func() *string {
			v, ok := data["distributorNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["distributorNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EnableAutoRun != nil {
		m["enableAutoRun"] = p.EnableAutoRun
	}
	if p.EnableAtomicCommit != nil {
		m["enableAtomicCommit"] = p.EnableAtomicCommit
	}
	if p.TransactionUseDistributor != nil {
		m["transactionUseDistributor"] = p.TransactionUseDistributor
	}
	if p.AcquireActionUseJobQueue != nil {
		m["acquireActionUseJobQueue"] = p.AcquireActionUseJobQueue
	}
	if p.DistributorNamespaceId != nil {
		m["distributorNamespaceId"] = p.DistributorNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	return m
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
