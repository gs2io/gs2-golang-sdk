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

package guard

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId    *string              `json:"namespaceId"`
	Name           *string              `json:"name"`
	Description    *string              `json:"description"`
	BlockingPolicy *BlockingPolicyModel `json:"blockingPolicy"`
	CreatedAt      *int64               `json:"createdAt"`
	UpdatedAt      *int64               `json:"updatedAt"`
	Revision       *int64               `json:"revision"`
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
		if v, ok := d["blockingPolicy"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BlockingPolicy)
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
		BlockingPolicy: func() *BlockingPolicyModel {
			v, ok := data["blockingPolicy"]
			if !ok || v == nil {
				return nil
			}
			return NewBlockingPolicyModelFromDict(core.CastMap(data["blockingPolicy"])).Pointer()
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
	if p.BlockingPolicy != nil {
		m["blockingPolicy"] = func() map[string]interface{} {
			if p.BlockingPolicy == nil {
				return nil
			}
			return p.BlockingPolicy.ToDict()
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

type BlockingPolicyModel struct {
	PassServices                 []*string `json:"passServices"`
	DefaultRestriction           *string   `json:"defaultRestriction"`
	LocationDetection            *string   `json:"locationDetection"`
	Locations                    []*string `json:"locations"`
	LocationRestriction          *string   `json:"locationRestriction"`
	AnonymousIpDetection         *string   `json:"anonymousIpDetection"`
	AnonymousIpRestriction       *string   `json:"anonymousIpRestriction"`
	HostingProviderIpDetection   *string   `json:"hostingProviderIpDetection"`
	HostingProviderIpRestriction *string   `json:"hostingProviderIpRestriction"`
	ReputationIpDetection        *string   `json:"reputationIpDetection"`
	ReputationIpRestriction      *string   `json:"reputationIpRestriction"`
	IpAddressesDetection         *string   `json:"ipAddressesDetection"`
	IpAddresses                  []*string `json:"ipAddresses"`
	IpAddressRestriction         *string   `json:"ipAddressRestriction"`
}

func (p *BlockingPolicyModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BlockingPolicyModel{}
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
		*p = BlockingPolicyModel{}
	} else {
		*p = BlockingPolicyModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["passServices"]; ok && v != nil {
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
				p.PassServices = l
			}
		}
		if v, ok := d["defaultRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DefaultRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DefaultRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DefaultRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DefaultRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DefaultRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DefaultRestriction)
				}
			}
		}
		if v, ok := d["locationDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LocationDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LocationDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LocationDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LocationDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LocationDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LocationDetection)
				}
			}
		}
		if v, ok := d["locations"]; ok && v != nil {
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
				p.Locations = l
			}
		}
		if v, ok := d["locationRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LocationRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LocationRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LocationRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LocationRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LocationRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LocationRestriction)
				}
			}
		}
		if v, ok := d["anonymousIpDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AnonymousIpDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AnonymousIpDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AnonymousIpDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AnonymousIpDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AnonymousIpDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AnonymousIpDetection)
				}
			}
		}
		if v, ok := d["anonymousIpRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AnonymousIpRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AnonymousIpRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AnonymousIpRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AnonymousIpRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AnonymousIpRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AnonymousIpRestriction)
				}
			}
		}
		if v, ok := d["hostingProviderIpDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.HostingProviderIpDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.HostingProviderIpDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.HostingProviderIpDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.HostingProviderIpDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.HostingProviderIpDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.HostingProviderIpDetection)
				}
			}
		}
		if v, ok := d["hostingProviderIpRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.HostingProviderIpRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.HostingProviderIpRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.HostingProviderIpRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.HostingProviderIpRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.HostingProviderIpRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.HostingProviderIpRestriction)
				}
			}
		}
		if v, ok := d["reputationIpDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReputationIpDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReputationIpDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReputationIpDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReputationIpDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReputationIpDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReputationIpDetection)
				}
			}
		}
		if v, ok := d["reputationIpRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReputationIpRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReputationIpRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReputationIpRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReputationIpRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReputationIpRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReputationIpRestriction)
				}
			}
		}
		if v, ok := d["ipAddressesDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.IpAddressesDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.IpAddressesDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.IpAddressesDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.IpAddressesDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.IpAddressesDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.IpAddressesDetection)
				}
			}
		}
		if v, ok := d["ipAddresses"]; ok && v != nil {
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
				p.IpAddresses = l
			}
		}
		if v, ok := d["ipAddressRestriction"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.IpAddressRestriction = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.IpAddressRestriction = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.IpAddressRestriction = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.IpAddressRestriction = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.IpAddressRestriction = &strValue
				default:
					_ = json.Unmarshal(*v, &p.IpAddressRestriction)
				}
			}
		}
	}
	return nil
}

func NewBlockingPolicyModelFromJson(data string) BlockingPolicyModel {
	req := BlockingPolicyModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBlockingPolicyModelFromDict(data map[string]interface{}) BlockingPolicyModel {
	return BlockingPolicyModel{
		PassServices: func() []*string {
			v, ok := data["passServices"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		DefaultRestriction: func() *string {
			v, ok := data["defaultRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["defaultRestriction"])
		}(),
		LocationDetection: func() *string {
			v, ok := data["locationDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["locationDetection"])
		}(),
		Locations: func() []*string {
			v, ok := data["locations"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		LocationRestriction: func() *string {
			v, ok := data["locationRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["locationRestriction"])
		}(),
		AnonymousIpDetection: func() *string {
			v, ok := data["anonymousIpDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["anonymousIpDetection"])
		}(),
		AnonymousIpRestriction: func() *string {
			v, ok := data["anonymousIpRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["anonymousIpRestriction"])
		}(),
		HostingProviderIpDetection: func() *string {
			v, ok := data["hostingProviderIpDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["hostingProviderIpDetection"])
		}(),
		HostingProviderIpRestriction: func() *string {
			v, ok := data["hostingProviderIpRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["hostingProviderIpRestriction"])
		}(),
		ReputationIpDetection: func() *string {
			v, ok := data["reputationIpDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["reputationIpDetection"])
		}(),
		ReputationIpRestriction: func() *string {
			v, ok := data["reputationIpRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["reputationIpRestriction"])
		}(),
		IpAddressesDetection: func() *string {
			v, ok := data["ipAddressesDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ipAddressesDetection"])
		}(),
		IpAddresses: func() []*string {
			v, ok := data["ipAddresses"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		IpAddressRestriction: func() *string {
			v, ok := data["ipAddressRestriction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ipAddressRestriction"])
		}(),
	}
}

func (p BlockingPolicyModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PassServices != nil {
		m["passServices"] = core.CastStringsFromDict(
			p.PassServices,
		)
	}
	if p.DefaultRestriction != nil {
		m["defaultRestriction"] = p.DefaultRestriction
	}
	if p.LocationDetection != nil {
		m["locationDetection"] = p.LocationDetection
	}
	if p.Locations != nil {
		m["locations"] = core.CastStringsFromDict(
			p.Locations,
		)
	}
	if p.LocationRestriction != nil {
		m["locationRestriction"] = p.LocationRestriction
	}
	if p.AnonymousIpDetection != nil {
		m["anonymousIpDetection"] = p.AnonymousIpDetection
	}
	if p.AnonymousIpRestriction != nil {
		m["anonymousIpRestriction"] = p.AnonymousIpRestriction
	}
	if p.HostingProviderIpDetection != nil {
		m["hostingProviderIpDetection"] = p.HostingProviderIpDetection
	}
	if p.HostingProviderIpRestriction != nil {
		m["hostingProviderIpRestriction"] = p.HostingProviderIpRestriction
	}
	if p.ReputationIpDetection != nil {
		m["reputationIpDetection"] = p.ReputationIpDetection
	}
	if p.ReputationIpRestriction != nil {
		m["reputationIpRestriction"] = p.ReputationIpRestriction
	}
	if p.IpAddressesDetection != nil {
		m["ipAddressesDetection"] = p.IpAddressesDetection
	}
	if p.IpAddresses != nil {
		m["ipAddresses"] = core.CastStringsFromDict(
			p.IpAddresses,
		)
	}
	if p.IpAddressRestriction != nil {
		m["ipAddressRestriction"] = p.IpAddressRestriction
	}
	return m
}

func (p BlockingPolicyModel) Pointer() *BlockingPolicyModel {
	return &p
}

func CastBlockingPolicyModels(data []interface{}) []BlockingPolicyModel {
	v := make([]BlockingPolicyModel, 0)
	for _, d := range data {
		v = append(v, NewBlockingPolicyModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBlockingPolicyModelsFromDict(data []BlockingPolicyModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
