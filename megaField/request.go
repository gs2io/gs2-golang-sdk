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

package megaField

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeNamespacesRequestFromDict(dict2)
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	RequestId    *string     `json:"requestId"`
	ContextStack *string     `json:"contextStack"`
	Name         *string     `json:"name"`
	Description  *string     `json:"description"`
	LogSetting   *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateNamespaceRequestFromDict(dict2)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceStatusRequestFromDict(dict2)
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetNamespaceRequestFromDict(dict2)
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	RequestId     *string     `json:"requestId"`
	ContextStack  *string     `json:"contextStack"`
	NamespaceName *string     `json:"namespaceName"`
	Description   *string     `json:"description"`
	LogSetting    *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateNamespaceRequestFromDict(dict2)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Description:   core.CastString(data["description"]),
		LogSetting:    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"logSetting":    p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteNamespaceRequestFromDict(dict2)
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type DescribeAreaModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDescribeAreaModelsRequestFromJson(data string) DescribeAreaModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeAreaModelsRequestFromDict(dict2)
}

func NewDescribeAreaModelsRequestFromDict(data map[string]interface{}) DescribeAreaModelsRequest {
	return DescribeAreaModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeAreaModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeAreaModelsRequest) Pointer() *DescribeAreaModelsRequest {
	return &p
}

type GetAreaModelRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
}

func NewGetAreaModelRequestFromJson(data string) GetAreaModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetAreaModelRequestFromDict(dict2)
}

func NewGetAreaModelRequestFromDict(data map[string]interface{}) GetAreaModelRequest {
	return GetAreaModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p GetAreaModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p GetAreaModelRequest) Pointer() *GetAreaModelRequest {
	return &p
}

type DescribeAreaModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeAreaModelMastersRequestFromJson(data string) DescribeAreaModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeAreaModelMastersRequestFromDict(dict2)
}

func NewDescribeAreaModelMastersRequestFromDict(data map[string]interface{}) DescribeAreaModelMastersRequest {
	return DescribeAreaModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAreaModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAreaModelMastersRequest) Pointer() *DescribeAreaModelMastersRequest {
	return &p
}

type CreateAreaModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateAreaModelMasterRequestFromJson(data string) CreateAreaModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateAreaModelMasterRequestFromDict(dict2)
}

func NewCreateAreaModelMasterRequestFromDict(data map[string]interface{}) CreateAreaModelMasterRequest {
	return CreateAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateAreaModelMasterRequest) Pointer() *CreateAreaModelMasterRequest {
	return &p
}

type GetAreaModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
}

func NewGetAreaModelMasterRequestFromJson(data string) GetAreaModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetAreaModelMasterRequestFromDict(dict2)
}

func NewGetAreaModelMasterRequestFromDict(data map[string]interface{}) GetAreaModelMasterRequest {
	return GetAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p GetAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p GetAreaModelMasterRequest) Pointer() *GetAreaModelMasterRequest {
	return &p
}

type UpdateAreaModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewUpdateAreaModelMasterRequestFromJson(data string) UpdateAreaModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateAreaModelMasterRequestFromDict(dict2)
}

func NewUpdateAreaModelMasterRequestFromDict(data map[string]interface{}) UpdateAreaModelMasterRequest {
	return UpdateAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateAreaModelMasterRequest) Pointer() *UpdateAreaModelMasterRequest {
	return &p
}

type DeleteAreaModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
}

func NewDeleteAreaModelMasterRequestFromJson(data string) DeleteAreaModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteAreaModelMasterRequestFromDict(dict2)
}

func NewDeleteAreaModelMasterRequestFromDict(data map[string]interface{}) DeleteAreaModelMasterRequest {
	return DeleteAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p DeleteAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p DeleteAreaModelMasterRequest) Pointer() *DeleteAreaModelMasterRequest {
	return &p
}

type DescribeLayerModelsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
}

func NewDescribeLayerModelsRequestFromJson(data string) DescribeLayerModelsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeLayerModelsRequestFromDict(dict2)
}

func NewDescribeLayerModelsRequestFromDict(data map[string]interface{}) DescribeLayerModelsRequest {
	return DescribeLayerModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p DescribeLayerModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p DescribeLayerModelsRequest) Pointer() *DescribeLayerModelsRequest {
	return &p
}

type GetLayerModelRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AreaModelName  *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
}

func NewGetLayerModelRequestFromJson(data string) GetLayerModelRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetLayerModelRequestFromDict(dict2)
}

func NewGetLayerModelRequestFromDict(data map[string]interface{}) GetLayerModelRequest {
	return GetLayerModelRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p GetLayerModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p GetLayerModelRequest) Pointer() *GetLayerModelRequest {
	return &p
}

type DescribeLayerModelMastersRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeLayerModelMastersRequestFromJson(data string) DescribeLayerModelMastersRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDescribeLayerModelMastersRequestFromDict(dict2)
}

func NewDescribeLayerModelMastersRequestFromDict(data map[string]interface{}) DescribeLayerModelMastersRequest {
	return DescribeLayerModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeLayerModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeLayerModelMastersRequest) Pointer() *DescribeLayerModelMastersRequest {
	return &p
}

type CreateLayerModelMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AreaModelName *string `json:"areaModelName"`
	Name          *string `json:"name"`
	Description   *string `json:"description"`
	Metadata      *string `json:"metadata"`
}

func NewCreateLayerModelMasterRequestFromJson(data string) CreateLayerModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewCreateLayerModelMasterRequestFromDict(dict2)
}

func NewCreateLayerModelMasterRequestFromDict(data map[string]interface{}) CreateLayerModelMasterRequest {
	return CreateLayerModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateLayerModelMasterRequest) Pointer() *CreateLayerModelMasterRequest {
	return &p
}

type GetLayerModelMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AreaModelName  *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
}

func NewGetLayerModelMasterRequestFromJson(data string) GetLayerModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetLayerModelMasterRequestFromDict(dict2)
}

func NewGetLayerModelMasterRequestFromDict(data map[string]interface{}) GetLayerModelMasterRequest {
	return GetLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p GetLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p GetLayerModelMasterRequest) Pointer() *GetLayerModelMasterRequest {
	return &p
}

type UpdateLayerModelMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AreaModelName  *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
	Description    *string `json:"description"`
	Metadata       *string `json:"metadata"`
}

func NewUpdateLayerModelMasterRequestFromJson(data string) UpdateLayerModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateLayerModelMasterRequestFromDict(dict2)
}

func NewUpdateLayerModelMasterRequestFromDict(data map[string]interface{}) UpdateLayerModelMasterRequest {
	return UpdateLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
	}
}

func (p UpdateLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"description":    p.Description,
		"metadata":       p.Metadata,
	}
}

func (p UpdateLayerModelMasterRequest) Pointer() *UpdateLayerModelMasterRequest {
	return &p
}

type DeleteLayerModelMasterRequest struct {
	RequestId      *string `json:"requestId"`
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AreaModelName  *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
}

func NewDeleteLayerModelMasterRequestFromJson(data string) DeleteLayerModelMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewDeleteLayerModelMasterRequestFromDict(dict2)
}

func NewDeleteLayerModelMasterRequestFromDict(data map[string]interface{}) DeleteLayerModelMasterRequest {
	return DeleteLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p DeleteLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p DeleteLayerModelMasterRequest) Pointer() *DeleteLayerModelMasterRequest {
	return &p
}

type ExportMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewExportMasterRequestFromDict(dict2)
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
	return ExportMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
	return &p
}

type GetCurrentFieldMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentFieldMasterRequestFromJson(data string) GetCurrentFieldMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewGetCurrentFieldMasterRequestFromDict(dict2)
}

func NewGetCurrentFieldMasterRequestFromDict(data map[string]interface{}) GetCurrentFieldMasterRequest {
	return GetCurrentFieldMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentFieldMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentFieldMasterRequest) Pointer() *GetCurrentFieldMasterRequest {
	return &p
}

type UpdateCurrentFieldMasterRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
}

func NewUpdateCurrentFieldMasterRequestFromJson(data string) UpdateCurrentFieldMasterRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentFieldMasterRequestFromDict(dict2)
}

func NewUpdateCurrentFieldMasterRequestFromDict(data map[string]interface{}) UpdateCurrentFieldMasterRequest {
	return UpdateCurrentFieldMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentFieldMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentFieldMasterRequest) Pointer() *UpdateCurrentFieldMasterRequest {
	return &p
}

type UpdateCurrentFieldMasterFromGitHubRequest struct {
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentFieldMasterFromGitHubRequestFromJson(data string) UpdateCurrentFieldMasterFromGitHubRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewUpdateCurrentFieldMasterFromGitHubRequestFromDict(dict2)
}

func NewUpdateCurrentFieldMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentFieldMasterFromGitHubRequest {
	return UpdateCurrentFieldMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentFieldMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentFieldMasterFromGitHubRequest) Pointer() *UpdateCurrentFieldMasterFromGitHubRequest {
	return &p
}

type PutPositionRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Position           *Position `json:"position"`
	Vector             *Vector   `json:"vector"`
	R                  *float32  `json:"r"`
}

func NewPutPositionRequestFromJson(data string) PutPositionRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewPutPositionRequestFromDict(dict2)
}

func NewPutPositionRequestFromDict(data map[string]interface{}) PutPositionRequest {
	return PutPositionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Vector:         NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
	}
}

func (p PutPositionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"vector":         p.Vector.ToDict(),
		"r":              p.R,
	}
}

func (p PutPositionRequest) Pointer() *PutPositionRequest {
	return &p
}

type PutPositionByUserIdRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Position           *Position `json:"position"`
	Vector             *Vector   `json:"vector"`
	R                  *float32  `json:"r"`
}

func NewPutPositionByUserIdRequestFromJson(data string) PutPositionByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewPutPositionByUserIdRequestFromDict(dict2)
}

func NewPutPositionByUserIdRequestFromDict(data map[string]interface{}) PutPositionByUserIdRequest {
	return PutPositionByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Vector:         NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
	}
}

func (p PutPositionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"vector":         p.Vector.ToDict(),
		"r":              p.R,
	}
}

func (p PutPositionByUserIdRequest) Pointer() *PutPositionByUserIdRequest {
	return &p
}

type FetchPositionRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	UserIds            []*string `json:"userIds"`
}

func NewFetchPositionRequestFromJson(data string) FetchPositionRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewFetchPositionRequestFromDict(dict2)
}

func NewFetchPositionRequestFromDict(data map[string]interface{}) FetchPositionRequest {
	return FetchPositionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		UserIds:        core.CastStrings(core.CastArray(data["userIds"])),
	}
}

func (p FetchPositionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"userIds": core.CastStringsFromDict(
			p.UserIds,
		),
	}
}

func (p FetchPositionRequest) Pointer() *FetchPositionRequest {
	return &p
}

type FetchPositionFromSystemRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	UserIds            []*string `json:"userIds"`
}

func NewFetchPositionFromSystemRequestFromJson(data string) FetchPositionFromSystemRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewFetchPositionFromSystemRequestFromDict(dict2)
}

func NewFetchPositionFromSystemRequestFromDict(data map[string]interface{}) FetchPositionFromSystemRequest {
	return FetchPositionFromSystemRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		UserIds:        core.CastStrings(core.CastArray(data["userIds"])),
	}
}

func (p FetchPositionFromSystemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"userIds": core.CastStringsFromDict(
			p.UserIds,
		),
	}
}

func (p FetchPositionFromSystemRequest) Pointer() *FetchPositionFromSystemRequest {
	return &p
}

type NearUserIdsRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Point              *Position `json:"point"`
	R                  *float32  `json:"r"`
	Limit              *int32    `json:"limit"`
}

func NewNearUserIdsRequestFromJson(data string) NearUserIdsRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewNearUserIdsRequestFromDict(dict2)
}

func NewNearUserIdsRequestFromDict(data map[string]interface{}) NearUserIdsRequest {
	return NearUserIdsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Point:          NewPositionFromDict(core.CastMap(data["point"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p NearUserIdsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"point":          p.Point.ToDict(),
		"r":              p.R,
		"limit":          p.Limit,
	}
}

func (p NearUserIdsRequest) Pointer() *NearUserIdsRequest {
	return &p
}

type NearUserIdsFromSystemRequest struct {
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Point              *Position `json:"point"`
	R                  *float32  `json:"r"`
	Limit              *int32    `json:"limit"`
}

func NewNearUserIdsFromSystemRequestFromJson(data string) NearUserIdsFromSystemRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewNearUserIdsFromSystemRequestFromDict(dict2)
}

func NewNearUserIdsFromSystemRequestFromDict(data map[string]interface{}) NearUserIdsFromSystemRequest {
	return NearUserIdsFromSystemRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Point:          NewPositionFromDict(core.CastMap(data["point"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p NearUserIdsFromSystemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"point":          p.Point.ToDict(),
		"r":              p.R,
		"limit":          p.Limit,
	}
}

func (p NearUserIdsFromSystemRequest) Pointer() *NearUserIdsFromSystemRequest {
	return &p
}

type ActionRequest struct {
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	AccessToken        *string     `json:"accessToken"`
	AreaModelName      *string     `json:"areaModelName"`
	LayerModelName     *string     `json:"layerModelName"`
	Position           *MyPosition `json:"position"`
	Scopes             []Scope     `json:"scopes"`
}

func NewActionRequestFromJson(data string) ActionRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewActionRequestFromDict(dict2)
}

func NewActionRequestFromDict(data map[string]interface{}) ActionRequest {
	return ActionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewMyPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Scopes:         CastScopes(core.CastArray(data["scopes"])),
	}
}

func (p ActionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"scopes": CastScopesFromDict(
			p.Scopes,
		),
	}
}

func (p ActionRequest) Pointer() *ActionRequest {
	return &p
}

type ActionByUserIdRequest struct {
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	UserId             *string     `json:"userId"`
	AreaModelName      *string     `json:"areaModelName"`
	LayerModelName     *string     `json:"layerModelName"`
	Position           *MyPosition `json:"position"`
	Scopes             []Scope     `json:"scopes"`
}

func NewActionByUserIdRequestFromJson(data string) ActionByUserIdRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewActionByUserIdRequestFromDict(dict2)
}

func NewActionByUserIdRequestFromDict(data map[string]interface{}) ActionByUserIdRequest {
	return ActionByUserIdRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		UserId:         core.CastString(data["userId"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewMyPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Scopes:         CastScopes(core.CastArray(data["scopes"])),
	}
}

func (p ActionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"scopes": CastScopesFromDict(
			p.Scopes,
		),
	}
}

func (p ActionByUserIdRequest) Pointer() *ActionByUserIdRequest {
	return &p
}
