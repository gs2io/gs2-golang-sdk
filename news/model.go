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

package news

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string     `json:"namespaceId"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	Version     *string     `json:"version"`
	LogSetting  *LogSetting `json:"logSetting"`
	CreatedAt   *int64      `json:"createdAt"`
	UpdatedAt   *int64      `json:"updatedAt"`
	Revision    *int64      `json:"revision"`
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
		if v, ok := d["version"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Version = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Version = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Version = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Version = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Version = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Version)
				}
			}
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
		NamespaceId: core.CastString(data["namespaceId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Version:     core.CastString(data["version"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
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
	var version *string
	if p.Version != nil {
		version = p.Version
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"name":        name,
		"description": description,
		"version":     version,
		"logSetting":  logSetting,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
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

type Progress struct {
	ProgressId   *string `json:"progressId"`
	UploadToken  *string `json:"uploadToken"`
	Generated    *int32  `json:"generated"`
	PatternCount *int32  `json:"patternCount"`
	CreatedAt    *int64  `json:"createdAt"`
	UpdatedAt    *int64  `json:"updatedAt"`
	Revision     *int64  `json:"revision"`
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
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UploadToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UploadToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UploadToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UploadToken)
				}
			}
		}
		if v, ok := d["generated"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Generated)
		}
		if v, ok := d["patternCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PatternCount)
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
		ProgressId:   core.CastString(data["progressId"]),
		UploadToken:  core.CastString(data["uploadToken"]),
		Generated:    core.CastInt32(data["generated"]),
		PatternCount: core.CastInt32(data["patternCount"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		UpdatedAt:    core.CastInt64(data["updatedAt"]),
		Revision:     core.CastInt64(data["revision"]),
	}
}

func (p Progress) ToDict() map[string]interface{} {

	var progressId *string
	if p.ProgressId != nil {
		progressId = p.ProgressId
	}
	var uploadToken *string
	if p.UploadToken != nil {
		uploadToken = p.UploadToken
	}
	var generated *int32
	if p.Generated != nil {
		generated = p.Generated
	}
	var patternCount *int32
	if p.PatternCount != nil {
		patternCount = p.PatternCount
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
		"progressId":   progressId,
		"uploadToken":  uploadToken,
		"generated":    generated,
		"patternCount": patternCount,
		"createdAt":    createdAt,
		"updatedAt":    updatedAt,
		"revision":     revision,
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

type Output struct {
	OutputId  *string `json:"outputId"`
	Name      *string `json:"name"`
	Text      *string `json:"text"`
	CreatedAt *int64  `json:"createdAt"`
	Revision  *int64  `json:"revision"`
}

func (p *Output) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Output{}
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
		*p = Output{}
	} else {
		*p = Output{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["outputId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OutputId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OutputId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OutputId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OutputId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OutputId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OutputId)
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
		if v, ok := d["text"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Text = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Text = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Text = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Text = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Text = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Text)
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

func NewOutputFromJson(data string) Output {
	req := Output{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewOutputFromDict(data map[string]interface{}) Output {
	return Output{
		OutputId:  core.CastString(data["outputId"]),
		Name:      core.CastString(data["name"]),
		Text:      core.CastString(data["text"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Output) ToDict() map[string]interface{} {

	var outputId *string
	if p.OutputId != nil {
		outputId = p.OutputId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var text *string
	if p.Text != nil {
		text = p.Text
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"outputId":  outputId,
		"name":      name,
		"text":      text,
		"createdAt": createdAt,
		"revision":  revision,
	}
}

func (p Output) Pointer() *Output {
	return &p
}

func CastOutputs(data []interface{}) []Output {
	v := make([]Output, 0)
	for _, d := range data {
		v = append(v, NewOutputFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOutputsFromDict(data []Output) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type View struct {
	Contents       []Content `json:"contents"`
	RemoveContents []Content `json:"removeContents"`
}

func (p *View) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = View{}
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
		*p = View{}
	} else {
		*p = View{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["contents"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Contents)
		}
		if v, ok := d["removeContents"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RemoveContents)
		}
	}
	return nil
}

func NewViewFromJson(data string) View {
	req := View{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewViewFromDict(data map[string]interface{}) View {
	return View{
		Contents:       CastContents(core.CastArray(data["contents"])),
		RemoveContents: CastContents(core.CastArray(data["removeContents"])),
	}
}

func (p View) ToDict() map[string]interface{} {

	var contents []interface{}
	if p.Contents != nil {
		contents = CastContentsFromDict(
			p.Contents,
		)
	}
	var removeContents []interface{}
	if p.RemoveContents != nil {
		removeContents = CastContentsFromDict(
			p.RemoveContents,
		)
	}
	return map[string]interface{}{
		"contents":       contents,
		"removeContents": removeContents,
	}
}

func (p View) Pointer() *View {
	return &p
}

func CastViews(data []interface{}) []View {
	v := make([]View, 0)
	for _, d := range data {
		v = append(v, NewViewFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastViewsFromDict(data []View) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Content struct {
	Section     *string `json:"section"`
	Content     *string `json:"content"`
	FrontMatter *string `json:"frontMatter"`
}

func (p *Content) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Content{}
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
		*p = Content{}
	} else {
		*p = Content{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["section"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Section = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Section = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Section = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Section = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Section = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Section)
				}
			}
		}
		if v, ok := d["content"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Content = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Content = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Content = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Content = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Content = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Content)
				}
			}
		}
		if v, ok := d["frontMatter"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FrontMatter = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FrontMatter = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FrontMatter = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FrontMatter = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FrontMatter = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FrontMatter)
				}
			}
		}
	}
	return nil
}

func NewContentFromJson(data string) Content {
	req := Content{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewContentFromDict(data map[string]interface{}) Content {
	return Content{
		Section:     core.CastString(data["section"]),
		Content:     core.CastString(data["content"]),
		FrontMatter: core.CastString(data["frontMatter"]),
	}
}

func (p Content) ToDict() map[string]interface{} {

	var section *string
	if p.Section != nil {
		section = p.Section
	}
	var content *string
	if p.Content != nil {
		content = p.Content
	}
	var frontMatter *string
	if p.FrontMatter != nil {
		frontMatter = p.FrontMatter
	}
	return map[string]interface{}{
		"section":     section,
		"content":     content,
		"frontMatter": frontMatter,
	}
}

func (p Content) Pointer() *Content {
	return &p
}

func CastContents(data []interface{}) []Content {
	v := make([]Content, 0)
	for _, d := range data {
		v = append(v, NewContentFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastContentsFromDict(data []Content) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type News struct {
	Section         *string `json:"section"`
	Content         *string `json:"content"`
	Title           *string `json:"title"`
	ScheduleEventId *string `json:"scheduleEventId"`
	Timestamp       *int64  `json:"timestamp"`
	FrontMatter     *string `json:"frontMatter"`
}

func (p *News) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = News{}
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
		*p = News{}
	} else {
		*p = News{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["section"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Section = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Section = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Section = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Section = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Section = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Section)
				}
			}
		}
		if v, ok := d["content"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Content = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Content = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Content = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Content = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Content = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Content)
				}
			}
		}
		if v, ok := d["title"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Title = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Title = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Title = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Title = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Title = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Title)
				}
			}
		}
		if v, ok := d["scheduleEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleEventId)
				}
			}
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["frontMatter"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FrontMatter = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FrontMatter = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FrontMatter = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FrontMatter = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FrontMatter = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FrontMatter)
				}
			}
		}
	}
	return nil
}

func NewNewsFromJson(data string) News {
	req := News{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNewsFromDict(data map[string]interface{}) News {
	return News{
		Section:         core.CastString(data["section"]),
		Content:         core.CastString(data["content"]),
		Title:           core.CastString(data["title"]),
		ScheduleEventId: core.CastString(data["scheduleEventId"]),
		Timestamp:       core.CastInt64(data["timestamp"]),
		FrontMatter:     core.CastString(data["frontMatter"]),
	}
}

func (p News) ToDict() map[string]interface{} {

	var section *string
	if p.Section != nil {
		section = p.Section
	}
	var content *string
	if p.Content != nil {
		content = p.Content
	}
	var title *string
	if p.Title != nil {
		title = p.Title
	}
	var scheduleEventId *string
	if p.ScheduleEventId != nil {
		scheduleEventId = p.ScheduleEventId
	}
	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	var frontMatter *string
	if p.FrontMatter != nil {
		frontMatter = p.FrontMatter
	}
	return map[string]interface{}{
		"section":         section,
		"content":         content,
		"title":           title,
		"scheduleEventId": scheduleEventId,
		"timestamp":       timestamp,
		"frontMatter":     frontMatter,
	}
}

func (p News) Pointer() *News {
	return &p
}

func CastNewses(data []interface{}) []News {
	v := make([]News, 0)
	for _, d := range data {
		v = append(v, NewNewsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNewsesFromDict(data []News) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SetCookieRequestEntry struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *SetCookieRequestEntry) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetCookieRequestEntry{}
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
		*p = SetCookieRequestEntry{}
	} else {
		*p = SetCookieRequestEntry{}
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

func NewSetCookieRequestEntryFromJson(data string) SetCookieRequestEntry {
	req := SetCookieRequestEntry{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSetCookieRequestEntryFromDict(data map[string]interface{}) SetCookieRequestEntry {
	return SetCookieRequestEntry{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p SetCookieRequestEntry) ToDict() map[string]interface{} {

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

func (p SetCookieRequestEntry) Pointer() *SetCookieRequestEntry {
	return &p
}

func CastSetCookieRequestEntries(data []interface{}) []SetCookieRequestEntry {
	v := make([]SetCookieRequestEntry, 0)
	for _, d := range data {
		v = append(v, NewSetCookieRequestEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSetCookieRequestEntriesFromDict(data []SetCookieRequestEntry) []interface{} {
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
