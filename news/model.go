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
		Version: func() *string {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["version"])
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
	if p.Version != nil {
		m["version"] = p.Version
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
		ProgressId: func() *string {
			v, ok := data["progressId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["progressId"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		Generated: func() *int32 {
			v, ok := data["generated"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["generated"])
		}(),
		PatternCount: func() *int32 {
			v, ok := data["patternCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["patternCount"])
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

func (p Progress) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProgressId != nil {
		m["progressId"] = p.ProgressId
	}
	if p.UploadToken != nil {
		m["uploadToken"] = p.UploadToken
	}
	if p.Generated != nil {
		m["generated"] = p.Generated
	}
	if p.PatternCount != nil {
		m["patternCount"] = p.PatternCount
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
		OutputId: func() *string {
			v, ok := data["outputId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["outputId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Text: func() *string {
			v, ok := data["text"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["text"])
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

func (p Output) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.OutputId != nil {
		m["outputId"] = p.OutputId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Text != nil {
		m["text"] = p.Text
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		Contents: func() []Content {
			if data["contents"] == nil {
				return nil
			}
			return CastContents(core.CastArray(data["contents"]))
		}(),
		RemoveContents: func() []Content {
			if data["removeContents"] == nil {
				return nil
			}
			return CastContents(core.CastArray(data["removeContents"]))
		}(),
	}
}

func (p View) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Contents != nil {
		m["contents"] = CastContentsFromDict(
			p.Contents,
		)
	}
	if p.RemoveContents != nil {
		m["removeContents"] = CastContentsFromDict(
			p.RemoveContents,
		)
	}
	return m
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
		Section: func() *string {
			v, ok := data["section"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["section"])
		}(),
		Content: func() *string {
			v, ok := data["content"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["content"])
		}(),
		FrontMatter: func() *string {
			v, ok := data["frontMatter"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["frontMatter"])
		}(),
	}
}

func (p Content) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Section != nil {
		m["section"] = p.Section
	}
	if p.Content != nil {
		m["content"] = p.Content
	}
	if p.FrontMatter != nil {
		m["frontMatter"] = p.FrontMatter
	}
	return m
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
		Section: func() *string {
			v, ok := data["section"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["section"])
		}(),
		Content: func() *string {
			v, ok := data["content"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["content"])
		}(),
		Title: func() *string {
			v, ok := data["title"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["title"])
		}(),
		ScheduleEventId: func() *string {
			v, ok := data["scheduleEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleEventId"])
		}(),
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		FrontMatter: func() *string {
			v, ok := data["frontMatter"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["frontMatter"])
		}(),
	}
}

func (p News) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Section != nil {
		m["section"] = p.Section
	}
	if p.Content != nil {
		m["content"] = p.Content
	}
	if p.Title != nil {
		m["title"] = p.Title
	}
	if p.ScheduleEventId != nil {
		m["scheduleEventId"] = p.ScheduleEventId
	}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.FrontMatter != nil {
		m["frontMatter"] = p.FrontMatter
	}
	return m
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

func (p SetCookieRequestEntry) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
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
