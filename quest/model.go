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

package quest

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId         *string             `json:"namespaceId"`
	Name                *string             `json:"name"`
	Description         *string             `json:"description"`
	TransactionSetting  *TransactionSetting `json:"transactionSetting"`
	StartQuestScript    *ScriptSetting      `json:"startQuestScript"`
	CompleteQuestScript *ScriptSetting      `json:"completeQuestScript"`
	FailedQuestScript   *ScriptSetting      `json:"failedQuestScript"`
	LogSetting          *LogSetting         `json:"logSetting"`
	CreatedAt           *int64              `json:"createdAt"`
	UpdatedAt           *int64              `json:"updatedAt"`
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["startQuestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StartQuestScript)
		}
		if v, ok := d["completeQuestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteQuestScript)
		}
		if v, ok := d["failedQuestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FailedQuestScript)
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
		NamespaceId:         core.CastString(data["namespaceId"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		TransactionSetting:  NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		StartQuestScript:    NewScriptSettingFromDict(core.CastMap(data["startQuestScript"])).Pointer(),
		CompleteQuestScript: NewScriptSettingFromDict(core.CastMap(data["completeQuestScript"])).Pointer(),
		FailedQuestScript:   NewScriptSettingFromDict(core.CastMap(data["failedQuestScript"])).Pointer(),
		LogSetting:          NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:    core.CastString(data["queueNamespaceId"]),
		KeyId:               core.CastString(data["keyId"]),
		Revision:            core.CastInt64(data["revision"]),
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
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var startQuestScript map[string]interface{}
	if p.StartQuestScript != nil {
		startQuestScript = p.StartQuestScript.ToDict()
	}
	var completeQuestScript map[string]interface{}
	if p.CompleteQuestScript != nil {
		completeQuestScript = p.CompleteQuestScript.ToDict()
	}
	var failedQuestScript map[string]interface{}
	if p.FailedQuestScript != nil {
		failedQuestScript = p.FailedQuestScript.ToDict()
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
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":         namespaceId,
		"name":                name,
		"description":         description,
		"transactionSetting":  transactionSetting,
		"startQuestScript":    startQuestScript,
		"completeQuestScript": completeQuestScript,
		"failedQuestScript":   failedQuestScript,
		"logSetting":          logSetting,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"queueNamespaceId":    queueNamespaceId,
		"keyId":               keyId,
		"revision":            revision,
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

type QuestGroupModelMaster struct {
	QuestGroupModelId      *string `json:"questGroupModelId"`
	Name                   *string `json:"name"`
	Description            *string `json:"description"`
	Metadata               *string `json:"metadata"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
	Revision               *int64  `json:"revision"`
}

func (p *QuestGroupModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QuestGroupModelMaster{}
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
		*p = QuestGroupModelMaster{}
	} else {
		*p = QuestGroupModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["questGroupModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupModelId)
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
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
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

func NewQuestGroupModelMasterFromJson(data string) QuestGroupModelMaster {
	req := QuestGroupModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewQuestGroupModelMasterFromDict(data map[string]interface{}) QuestGroupModelMaster {
	return QuestGroupModelMaster{
		QuestGroupModelId:      core.CastString(data["questGroupModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p QuestGroupModelMaster) ToDict() map[string]interface{} {

	var questGroupModelId *string
	if p.QuestGroupModelId != nil {
		questGroupModelId = p.QuestGroupModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"questGroupModelId":      questGroupModelId,
		"name":                   name,
		"description":            description,
		"metadata":               metadata,
		"challengePeriodEventId": challengePeriodEventId,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
}

func (p QuestGroupModelMaster) Pointer() *QuestGroupModelMaster {
	return &p
}

func CastQuestGroupModelMasters(data []interface{}) []QuestGroupModelMaster {
	v := make([]QuestGroupModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelMastersFromDict(data []QuestGroupModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestModelMaster struct {
	QuestModelId                *string         `json:"questModelId"`
	QuestGroupName              *string         `json:"questGroupName"`
	Name                        *string         `json:"name"`
	Description                 *string         `json:"description"`
	Metadata                    *string         `json:"metadata"`
	Contents                    []Contents      `json:"contents"`
	ChallengePeriodEventId      *string         `json:"challengePeriodEventId"`
	FirstCompleteAcquireActions []AcquireAction `json:"firstCompleteAcquireActions"`
	ConsumeActions              []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions        []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames           []*string       `json:"premiseQuestNames"`
	CreatedAt                   *int64          `json:"createdAt"`
	UpdatedAt                   *int64          `json:"updatedAt"`
	Revision                    *int64          `json:"revision"`
}

func (p *QuestModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QuestModelMaster{}
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
		*p = QuestModelMaster{}
	} else {
		*p = QuestModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["questModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestModelId)
				}
			}
		}
		if v, ok := d["questGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupName)
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
		if v, ok := d["contents"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Contents)
		}
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
		if v, ok := d["firstCompleteAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FirstCompleteAcquireActions)
		}
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["failedAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FailedAcquireActions)
		}
		if v, ok := d["premiseQuestNames"]; ok && v != nil {
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
				p.PremiseQuestNames = l
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

func NewQuestModelMasterFromJson(data string) QuestModelMaster {
	req := QuestModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewQuestModelMasterFromDict(data map[string]interface{}) QuestModelMaster {
	return QuestModelMaster{
		QuestModelId:                core.CastString(data["questModelId"]),
		QuestGroupName:              core.CastString(data["questGroupName"]),
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		Metadata:                    core.CastString(data["metadata"]),
		Contents:                    CastContentses(core.CastArray(data["contents"])),
		ChallengePeriodEventId:      core.CastString(data["challengePeriodEventId"]),
		FirstCompleteAcquireActions: CastAcquireActions(core.CastArray(data["firstCompleteAcquireActions"])),
		ConsumeActions:              CastConsumeActions(core.CastArray(data["consumeActions"])),
		FailedAcquireActions:        CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
		PremiseQuestNames:           core.CastStrings(core.CastArray(data["premiseQuestNames"])),
		CreatedAt:                   core.CastInt64(data["createdAt"]),
		UpdatedAt:                   core.CastInt64(data["updatedAt"]),
		Revision:                    core.CastInt64(data["revision"]),
	}
}

func (p QuestModelMaster) ToDict() map[string]interface{} {

	var questModelId *string
	if p.QuestModelId != nil {
		questModelId = p.QuestModelId
	}
	var questGroupName *string
	if p.QuestGroupName != nil {
		questGroupName = p.QuestGroupName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var contents []interface{}
	if p.Contents != nil {
		contents = CastContentsesFromDict(
			p.Contents,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var firstCompleteAcquireActions []interface{}
	if p.FirstCompleteAcquireActions != nil {
		firstCompleteAcquireActions = CastAcquireActionsFromDict(
			p.FirstCompleteAcquireActions,
		)
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var failedAcquireActions []interface{}
	if p.FailedAcquireActions != nil {
		failedAcquireActions = CastAcquireActionsFromDict(
			p.FailedAcquireActions,
		)
	}
	var premiseQuestNames []interface{}
	if p.PremiseQuestNames != nil {
		premiseQuestNames = core.CastStringsFromDict(
			p.PremiseQuestNames,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"questModelId":                questModelId,
		"questGroupName":              questGroupName,
		"name":                        name,
		"description":                 description,
		"metadata":                    metadata,
		"contents":                    contents,
		"challengePeriodEventId":      challengePeriodEventId,
		"firstCompleteAcquireActions": firstCompleteAcquireActions,
		"consumeActions":              consumeActions,
		"failedAcquireActions":        failedAcquireActions,
		"premiseQuestNames":           premiseQuestNames,
		"createdAt":                   createdAt,
		"updatedAt":                   updatedAt,
		"revision":                    revision,
	}
}

func (p QuestModelMaster) Pointer() *QuestModelMaster {
	return &p
}

func CastQuestModelMasters(data []interface{}) []QuestModelMaster {
	v := make([]QuestModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelMastersFromDict(data []QuestModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentQuestMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentQuestMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentQuestMaster{}
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
		*p = CurrentQuestMaster{}
	} else {
		*p = CurrentQuestMaster{}
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

func NewCurrentQuestMasterFromJson(data string) CurrentQuestMaster {
	req := CurrentQuestMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentQuestMasterFromDict(data map[string]interface{}) CurrentQuestMaster {
	return CurrentQuestMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentQuestMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
}

func (p CurrentQuestMaster) Pointer() *CurrentQuestMaster {
	return &p
}

func CastCurrentQuestMasters(data []interface{}) []CurrentQuestMaster {
	v := make([]CurrentQuestMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentQuestMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentQuestMastersFromDict(data []CurrentQuestMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Contents struct {
	Metadata               *string         `json:"metadata"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	Weight                 *int32          `json:"weight"`
}

func (p *Contents) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Contents{}
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
		*p = Contents{}
	} else {
		*p = Contents{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["completeAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteAcquireActions)
		}
		if v, ok := d["weight"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Weight)
		}
	}
	return nil
}

func NewContentsFromJson(data string) Contents {
	req := Contents{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewContentsFromDict(data map[string]interface{}) Contents {
	return Contents{
		Metadata:               core.CastString(data["metadata"]),
		CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		Weight:                 core.CastInt32(data["weight"]),
	}
}

func (p Contents) ToDict() map[string]interface{} {

	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var completeAcquireActions []interface{}
	if p.CompleteAcquireActions != nil {
		completeAcquireActions = CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		)
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"metadata":               metadata,
		"completeAcquireActions": completeAcquireActions,
		"weight":                 weight,
	}
}

func (p Contents) Pointer() *Contents {
	return &p
}

func CastContentses(data []interface{}) []Contents {
	v := make([]Contents, 0)
	for _, d := range data {
		v = append(v, NewContentsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastContentsesFromDict(data []Contents) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Reward struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
	ItemId  *string `json:"itemId"`
	Value   *int32  `json:"value"`
}

func (p *Reward) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Reward{}
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
		*p = Reward{}
	} else {
		*p = Reward{}
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
		if v, ok := d["itemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemId)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewRewardFromJson(data string) Reward {
	req := Reward{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRewardFromDict(data map[string]interface{}) Reward {
	return Reward{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
		ItemId:  core.CastString(data["itemId"]),
		Value:   core.CastInt32(data["value"]),
	}
}

func (p Reward) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	var itemId *string
	if p.ItemId != nil {
		itemId = p.ItemId
	}
	var value *int32
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
		"itemId":  itemId,
		"value":   value,
	}
}

func (p Reward) Pointer() *Reward {
	return &p
}

func CastRewards(data []interface{}) []Reward {
	v := make([]Reward, 0)
	for _, d := range data {
		v = append(v, NewRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRewardsFromDict(data []Reward) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Progress struct {
	ProgressId    *string  `json:"progressId"`
	UserId        *string  `json:"userId"`
	TransactionId *string  `json:"transactionId"`
	QuestModelId  *string  `json:"questModelId"`
	RandomSeed    *int64   `json:"randomSeed"`
	Rewards       []Reward `json:"rewards"`
	FailedRewards []Reward `json:"failedRewards"`
	Metadata      *string  `json:"metadata"`
	CreatedAt     *int64   `json:"createdAt"`
	UpdatedAt     *int64   `json:"updatedAt"`
	Revision      *int64   `json:"revision"`
}

func (p *Progress) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Progress{}
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
		*p = Progress{}
	} else {
		*p = Progress{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["progressId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProgressId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProgressId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProgressId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProgressId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProgressId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProgressId)
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
		if v, ok := d["questModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestModelId)
				}
			}
		}
		if v, ok := d["randomSeed"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RandomSeed)
		}
		if v, ok := d["rewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rewards)
		}
		if v, ok := d["failedRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FailedRewards)
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

func NewProgressFromJson(data string) Progress {
	req := Progress{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewProgressFromDict(data map[string]interface{}) Progress {
	return Progress{
		ProgressId:    core.CastString(data["progressId"]),
		UserId:        core.CastString(data["userId"]),
		TransactionId: core.CastString(data["transactionId"]),
		QuestModelId:  core.CastString(data["questModelId"]),
		RandomSeed:    core.CastInt64(data["randomSeed"]),
		Rewards:       CastRewards(core.CastArray(data["rewards"])),
		FailedRewards: CastRewards(core.CastArray(data["failedRewards"])),
		Metadata:      core.CastString(data["metadata"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p Progress) ToDict() map[string]interface{} {

	var progressId *string
	if p.ProgressId != nil {
		progressId = p.ProgressId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var questModelId *string
	if p.QuestModelId != nil {
		questModelId = p.QuestModelId
	}
	var randomSeed *int64
	if p.RandomSeed != nil {
		randomSeed = p.RandomSeed
	}
	var rewards []interface{}
	if p.Rewards != nil {
		rewards = CastRewardsFromDict(
			p.Rewards,
		)
	}
	var failedRewards []interface{}
	if p.FailedRewards != nil {
		failedRewards = CastRewardsFromDict(
			p.FailedRewards,
		)
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"progressId":    progressId,
		"userId":        userId,
		"transactionId": transactionId,
		"questModelId":  questModelId,
		"randomSeed":    randomSeed,
		"rewards":       rewards,
		"failedRewards": failedRewards,
		"metadata":      metadata,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p Progress) Pointer() *Progress {
	return &p
}

func CastProgresses(data []interface{}) []Progress {
	v := make([]Progress, 0)
	for _, d := range data {
		v = append(v, NewProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProgressesFromDict(data []Progress) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CompletedQuestList struct {
	CompletedQuestListId *string   `json:"completedQuestListId"`
	UserId               *string   `json:"userId"`
	QuestGroupName       *string   `json:"questGroupName"`
	CompleteQuestNames   []*string `json:"completeQuestNames"`
	CreatedAt            *int64    `json:"createdAt"`
	UpdatedAt            *int64    `json:"updatedAt"`
	Revision             *int64    `json:"revision"`
}

func (p *CompletedQuestList) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CompletedQuestList{}
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
		*p = CompletedQuestList{}
	} else {
		*p = CompletedQuestList{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["completedQuestListId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompletedQuestListId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompletedQuestListId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompletedQuestListId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompletedQuestListId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompletedQuestListId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompletedQuestListId)
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
		if v, ok := d["questGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupName)
				}
			}
		}
		if v, ok := d["completeQuestNames"]; ok && v != nil {
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
				p.CompleteQuestNames = l
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

func NewCompletedQuestListFromJson(data string) CompletedQuestList {
	req := CompletedQuestList{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCompletedQuestListFromDict(data map[string]interface{}) CompletedQuestList {
	return CompletedQuestList{
		CompletedQuestListId: core.CastString(data["completedQuestListId"]),
		UserId:               core.CastString(data["userId"]),
		QuestGroupName:       core.CastString(data["questGroupName"]),
		CompleteQuestNames:   core.CastStrings(core.CastArray(data["completeQuestNames"])),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p CompletedQuestList) ToDict() map[string]interface{} {

	var completedQuestListId *string
	if p.CompletedQuestListId != nil {
		completedQuestListId = p.CompletedQuestListId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var questGroupName *string
	if p.QuestGroupName != nil {
		questGroupName = p.QuestGroupName
	}
	var completeQuestNames []interface{}
	if p.CompleteQuestNames != nil {
		completeQuestNames = core.CastStringsFromDict(
			p.CompleteQuestNames,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"completedQuestListId": completedQuestListId,
		"userId":               userId,
		"questGroupName":       questGroupName,
		"completeQuestNames":   completeQuestNames,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
}

func (p CompletedQuestList) Pointer() *CompletedQuestList {
	return &p
}

func CastCompletedQuestList(data []interface{}) []CompletedQuestList {
	v := make([]CompletedQuestList, 0)
	for _, d := range data {
		v = append(v, NewCompletedQuestListFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCompletedQuestListFromDict(data []CompletedQuestList) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestGroupModel struct {
	QuestGroupModelId      *string      `json:"questGroupModelId"`
	Name                   *string      `json:"name"`
	Metadata               *string      `json:"metadata"`
	Quests                 []QuestModel `json:"quests"`
	ChallengePeriodEventId *string      `json:"challengePeriodEventId"`
}

func (p *QuestGroupModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QuestGroupModel{}
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
		*p = QuestGroupModel{}
	} else {
		*p = QuestGroupModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["questGroupModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestGroupModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestGroupModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestGroupModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestGroupModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestGroupModelId)
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
		if v, ok := d["quests"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Quests)
		}
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewQuestGroupModelFromJson(data string) QuestGroupModel {
	req := QuestGroupModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewQuestGroupModelFromDict(data map[string]interface{}) QuestGroupModel {
	return QuestGroupModel{
		QuestGroupModelId:      core.CastString(data["questGroupModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Quests:                 CastQuestModels(core.CastArray(data["quests"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p QuestGroupModel) ToDict() map[string]interface{} {

	var questGroupModelId *string
	if p.QuestGroupModelId != nil {
		questGroupModelId = p.QuestGroupModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var quests []interface{}
	if p.Quests != nil {
		quests = CastQuestModelsFromDict(
			p.Quests,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	return map[string]interface{}{
		"questGroupModelId":      questGroupModelId,
		"name":                   name,
		"metadata":               metadata,
		"quests":                 quests,
		"challengePeriodEventId": challengePeriodEventId,
	}
}

func (p QuestGroupModel) Pointer() *QuestGroupModel {
	return &p
}

func CastQuestGroupModels(data []interface{}) []QuestGroupModel {
	v := make([]QuestGroupModel, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelsFromDict(data []QuestGroupModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestModel struct {
	QuestModelId                *string         `json:"questModelId"`
	Name                        *string         `json:"name"`
	Metadata                    *string         `json:"metadata"`
	Contents                    []Contents      `json:"contents"`
	ChallengePeriodEventId      *string         `json:"challengePeriodEventId"`
	FirstCompleteAcquireActions []AcquireAction `json:"firstCompleteAcquireActions"`
	ConsumeActions              []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions        []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames           []*string       `json:"premiseQuestNames"`
}

func (p *QuestModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QuestModel{}
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
		*p = QuestModel{}
	} else {
		*p = QuestModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["questModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QuestModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QuestModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QuestModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QuestModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QuestModelId)
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
		if v, ok := d["contents"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Contents)
		}
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
		if v, ok := d["firstCompleteAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FirstCompleteAcquireActions)
		}
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["failedAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FailedAcquireActions)
		}
		if v, ok := d["premiseQuestNames"]; ok && v != nil {
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
				p.PremiseQuestNames = l
			}
		}
	}
	return nil
}

func NewQuestModelFromJson(data string) QuestModel {
	req := QuestModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewQuestModelFromDict(data map[string]interface{}) QuestModel {
	return QuestModel{
		QuestModelId:                core.CastString(data["questModelId"]),
		Name:                        core.CastString(data["name"]),
		Metadata:                    core.CastString(data["metadata"]),
		Contents:                    CastContentses(core.CastArray(data["contents"])),
		ChallengePeriodEventId:      core.CastString(data["challengePeriodEventId"]),
		FirstCompleteAcquireActions: CastAcquireActions(core.CastArray(data["firstCompleteAcquireActions"])),
		ConsumeActions:              CastConsumeActions(core.CastArray(data["consumeActions"])),
		FailedAcquireActions:        CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
		PremiseQuestNames:           core.CastStrings(core.CastArray(data["premiseQuestNames"])),
	}
}

func (p QuestModel) ToDict() map[string]interface{} {

	var questModelId *string
	if p.QuestModelId != nil {
		questModelId = p.QuestModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var contents []interface{}
	if p.Contents != nil {
		contents = CastContentsesFromDict(
			p.Contents,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var firstCompleteAcquireActions []interface{}
	if p.FirstCompleteAcquireActions != nil {
		firstCompleteAcquireActions = CastAcquireActionsFromDict(
			p.FirstCompleteAcquireActions,
		)
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var failedAcquireActions []interface{}
	if p.FailedAcquireActions != nil {
		failedAcquireActions = CastAcquireActionsFromDict(
			p.FailedAcquireActions,
		)
	}
	var premiseQuestNames []interface{}
	if p.PremiseQuestNames != nil {
		premiseQuestNames = core.CastStringsFromDict(
			p.PremiseQuestNames,
		)
	}
	return map[string]interface{}{
		"questModelId":                questModelId,
		"name":                        name,
		"metadata":                    metadata,
		"contents":                    contents,
		"challengePeriodEventId":      challengePeriodEventId,
		"firstCompleteAcquireActions": firstCompleteAcquireActions,
		"consumeActions":              consumeActions,
		"failedAcquireActions":        failedAcquireActions,
		"premiseQuestNames":           premiseQuestNames,
	}
}

func (p QuestModel) Pointer() *QuestModel {
	return &p
}

func CastQuestModels(data []interface{}) []QuestModel {
	v := make([]QuestModel, 0)
	for _, d := range data {
		v = append(v, NewQuestModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelsFromDict(data []QuestModel) []interface{} {
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
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
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

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func (p *ConsumeAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeAction{}
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
		*p = ConsumeAction{}
	} else {
		*p = ConsumeAction{}
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

func NewConsumeActionFromJson(data string) ConsumeAction {
	req := ConsumeAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
	return ConsumeAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p ConsumeAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p ConsumeAction) Pointer() *ConsumeAction {
	return &p
}

func CastConsumeActions(data []interface{}) []ConsumeAction {
	v := make([]ConsumeAction, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionsFromDict(data []ConsumeAction) []interface{} {
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
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Config) ToDict() map[string]interface{} {

	var key *string
	if p.Key != nil {
		key = p.Key
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"key":   key,
		"value": value,
	}
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
		ApiKeyId:       core.CastString(data["apiKeyId"]),
		RepositoryName: core.CastString(data["repositoryName"]),
		SourcePath:     core.CastString(data["sourcePath"]),
		ReferenceType:  core.CastString(data["referenceType"]),
		CommitHash:     core.CastString(data["commitHash"]),
		BranchName:     core.CastString(data["branchName"]),
		TagName:        core.CastString(data["tagName"]),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
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
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {

	var triggerScriptId *string
	if p.TriggerScriptId != nil {
		triggerScriptId = p.TriggerScriptId
	}
	var doneTriggerTargetType *string
	if p.DoneTriggerTargetType != nil {
		doneTriggerTargetType = p.DoneTriggerTargetType
	}
	var doneTriggerScriptId *string
	if p.DoneTriggerScriptId != nil {
		doneTriggerScriptId = p.DoneTriggerScriptId
	}
	var doneTriggerQueueNamespaceId *string
	if p.DoneTriggerQueueNamespaceId != nil {
		doneTriggerQueueNamespaceId = p.DoneTriggerQueueNamespaceId
	}
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
		"doneTriggerQueueNamespaceId": doneTriggerQueueNamespaceId,
	}
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
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {

	var loggingNamespaceId *string
	if p.LoggingNamespaceId != nil {
		loggingNamespaceId = p.LoggingNamespaceId
	}
	return map[string]interface{}{
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

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
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
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
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
