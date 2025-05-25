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

package freeze

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeStagesResult struct {
	Items    []Stage              `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeStagesAsyncResult struct {
	result *DescribeStagesResult
	err    error
}

func NewDescribeStagesResultFromJson(data string) DescribeStagesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStagesResultFromDict(dict)
}

func NewDescribeStagesResultFromDict(data map[string]interface{}) DescribeStagesResult {
	return DescribeStagesResult{
		Items: func() []Stage {
			if data["items"] == nil {
				return nil
			}
			return CastStages(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeStagesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStagesFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeStagesResult) Pointer() *DescribeStagesResult {
	return &p
}

type GetStageResult struct {
	Item     *Stage               `json:"item"`
	Source   []Microservice       `json:"source"`
	Current  []Microservice       `json:"current"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStageAsyncResult struct {
	result *GetStageResult
	err    error
}

func NewGetStageResultFromJson(data string) GetStageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStageResultFromDict(dict)
}

func NewGetStageResultFromDict(data map[string]interface{}) GetStageResult {
	return GetStageResult{
		Item: func() *Stage {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Source: func() []Microservice {
			if data["source"] == nil {
				return nil
			}
			return CastMicroservices(core.CastArray(data["source"]))
		}(),
		Current: func() []Microservice {
			if data["current"] == nil {
				return nil
			}
			return CastMicroservices(core.CastArray(data["current"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"source": CastMicroservicesFromDict(
			p.Source,
		),
		"current": CastMicroservicesFromDict(
			p.Current,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetStageResult) Pointer() *GetStageResult {
	return &p
}

type PromoteStageResult struct {
	Item     *Stage               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PromoteStageAsyncResult struct {
	result *PromoteStageResult
	err    error
}

func NewPromoteStageResultFromJson(data string) PromoteStageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPromoteStageResultFromDict(dict)
}

func NewPromoteStageResultFromDict(data map[string]interface{}) PromoteStageResult {
	return PromoteStageResult{
		Item: func() *Stage {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PromoteStageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PromoteStageResult) Pointer() *PromoteStageResult {
	return &p
}

type RollbackStageResult struct {
	Item     *Stage               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RollbackStageAsyncResult struct {
	result *RollbackStageResult
	err    error
}

func NewRollbackStageResultFromJson(data string) RollbackStageResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRollbackStageResultFromDict(dict)
}

func NewRollbackStageResultFromDict(data map[string]interface{}) RollbackStageResult {
	return RollbackStageResult{
		Item: func() *Stage {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStageFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RollbackStageResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RollbackStageResult) Pointer() *RollbackStageResult {
	return &p
}

type DescribeOutputsResult struct {
	Items         []Output             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeOutputsAsyncResult struct {
	result *DescribeOutputsResult
	err    error
}

func NewDescribeOutputsResultFromJson(data string) DescribeOutputsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeOutputsResultFromDict(dict)
}

func NewDescribeOutputsResultFromDict(data map[string]interface{}) DescribeOutputsResult {
	return DescribeOutputsResult{
		Items: func() []Output {
			if data["items"] == nil {
				return nil
			}
			return CastOutputs(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeOutputsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastOutputsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeOutputsResult) Pointer() *DescribeOutputsResult {
	return &p
}

type GetOutputResult struct {
	Item     *Output              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetOutputAsyncResult struct {
	result *GetOutputResult
	err    error
}

func NewGetOutputResultFromJson(data string) GetOutputResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetOutputResultFromDict(dict)
}

func NewGetOutputResultFromDict(data map[string]interface{}) GetOutputResult {
	return GetOutputResult{
		Item: func() *Output {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewOutputFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetOutputResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetOutputResult) Pointer() *GetOutputResult {
	return &p
}
