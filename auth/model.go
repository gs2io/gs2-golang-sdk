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

package auth

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type AccessToken struct {
	Token                *string `json:"token"`
	UserId               *string `json:"userId"`
	FederationFromUserId *string `json:"federationFromUserId"`
	Expire               *int64  `json:"expire"`
	TimeOffset           *int32  `json:"timeOffset"`
}

func (p *AccessToken) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AccessToken{}
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
		*p = AccessToken{}
	} else {
		*p = AccessToken{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["token"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Token = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Token = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Token = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Token = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Token = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Token)
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
		if v, ok := d["federationFromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FederationFromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FederationFromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FederationFromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FederationFromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FederationFromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FederationFromUserId)
				}
			}
		}
		if v, ok := d["expire"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Expire)
		}
		if v, ok := d["timeOffset"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TimeOffset)
		}
	}
	return nil
}

func NewAccessTokenFromJson(data string) AccessToken {
	req := AccessToken{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAccessTokenFromDict(data map[string]interface{}) AccessToken {
	return AccessToken{
		Token:                core.CastString(data["token"]),
		UserId:               core.CastString(data["userId"]),
		FederationFromUserId: core.CastString(data["federationFromUserId"]),
		Expire:               core.CastInt64(data["expire"]),
		TimeOffset:           core.CastInt32(data["timeOffset"]),
	}
}

func (p AccessToken) ToDict() map[string]interface{} {

	var token *string
	if p.Token != nil {
		token = p.Token
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var federationFromUserId *string
	if p.FederationFromUserId != nil {
		federationFromUserId = p.FederationFromUserId
	}
	var expire *int64
	if p.Expire != nil {
		expire = p.Expire
	}
	var timeOffset *int32
	if p.TimeOffset != nil {
		timeOffset = p.TimeOffset
	}
	return map[string]interface{}{
		"token":                token,
		"userId":               userId,
		"federationFromUserId": federationFromUserId,
		"expire":               expire,
		"timeOffset":           timeOffset,
	}
}

func (p AccessToken) Pointer() *AccessToken {
	return &p
}

func CastAccessTokens(data []interface{}) []AccessToken {
	v := make([]AccessToken, 0)
	for _, d := range data {
		v = append(v, NewAccessTokenFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccessTokensFromDict(data []AccessToken) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
