// Code generated by "enumer -linecomment -sql -json -text -yaml -gqlgen -type=ApplicationType,AccessTokenType"; DO NOT EDIT.

package op

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const _ApplicationTypeName = "webuser_agentnative"

var _ApplicationTypeIndex = [...]uint8{0, 3, 13, 19}

const _ApplicationTypeLowerName = "webuser_agentnative"

func (i ApplicationType) String() string {
	if i < 0 || i >= ApplicationType(len(_ApplicationTypeIndex)-1) {
		return fmt.Sprintf("ApplicationType(%d)", i)
	}
	return _ApplicationTypeName[_ApplicationTypeIndex[i]:_ApplicationTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ApplicationTypeNoOp() {
	var x [1]struct{}
	_ = x[ApplicationTypeWeb-(0)]
	_ = x[ApplicationTypeUserAgent-(1)]
	_ = x[ApplicationTypeNative-(2)]
}

var _ApplicationTypeValues = []ApplicationType{ApplicationTypeWeb, ApplicationTypeUserAgent, ApplicationTypeNative}

var _ApplicationTypeNameToValueMap = map[string]ApplicationType{
	_ApplicationTypeName[0:3]:        ApplicationTypeWeb,
	_ApplicationTypeLowerName[0:3]:   ApplicationTypeWeb,
	_ApplicationTypeName[3:13]:       ApplicationTypeUserAgent,
	_ApplicationTypeLowerName[3:13]:  ApplicationTypeUserAgent,
	_ApplicationTypeName[13:19]:      ApplicationTypeNative,
	_ApplicationTypeLowerName[13:19]: ApplicationTypeNative,
}

var _ApplicationTypeNames = []string{
	_ApplicationTypeName[0:3],
	_ApplicationTypeName[3:13],
	_ApplicationTypeName[13:19],
}

// ApplicationTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ApplicationTypeString(s string) (ApplicationType, error) {
	if val, ok := _ApplicationTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ApplicationTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ApplicationType values", s)
}

// ApplicationTypeValues returns all values of the enum
func ApplicationTypeValues() []ApplicationType {
	return _ApplicationTypeValues
}

// ApplicationTypeStrings returns a slice of all String values of the enum
func ApplicationTypeStrings() []string {
	strs := make([]string, len(_ApplicationTypeNames))
	copy(strs, _ApplicationTypeNames)
	return strs
}

// IsAApplicationType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ApplicationType) IsAApplicationType() bool {
	for _, v := range _ApplicationTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ApplicationType
func (i ApplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplicationType
func (i *ApplicationType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ApplicationType should be a string, got %s", data)
	}

	var err error
	*i, err = ApplicationTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for ApplicationType
func (i ApplicationType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ApplicationType
func (i *ApplicationType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ApplicationTypeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for ApplicationType
func (i ApplicationType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ApplicationType
func (i *ApplicationType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ApplicationTypeString(s)
	return err
}

func (i ApplicationType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *ApplicationType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ApplicationType: %[1]T(%[1]v)", value)
	}

	val, err := ApplicationTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for ApplicationType
func (i ApplicationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(i.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for ApplicationType
func (i *ApplicationType) UnmarshalGQL(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("ApplicationType should be a string, got %T", value)
	}

	var err error
	*i, err = ApplicationTypeString(str)
	return err
}

const _AccessTokenTypeName = "bearerJWT"

var _AccessTokenTypeIndex = [...]uint8{0, 6, 9}

const _AccessTokenTypeLowerName = "bearerjwt"

func (i AccessTokenType) String() string {
	if i < 0 || i >= AccessTokenType(len(_AccessTokenTypeIndex)-1) {
		return fmt.Sprintf("AccessTokenType(%d)", i)
	}
	return _AccessTokenTypeName[_AccessTokenTypeIndex[i]:_AccessTokenTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _AccessTokenTypeNoOp() {
	var x [1]struct{}
	_ = x[AccessTokenTypeBearer-(0)]
	_ = x[AccessTokenTypeJWT-(1)]
}

var _AccessTokenTypeValues = []AccessTokenType{AccessTokenTypeBearer, AccessTokenTypeJWT}

var _AccessTokenTypeNameToValueMap = map[string]AccessTokenType{
	_AccessTokenTypeName[0:6]:      AccessTokenTypeBearer,
	_AccessTokenTypeLowerName[0:6]: AccessTokenTypeBearer,
	_AccessTokenTypeName[6:9]:      AccessTokenTypeJWT,
	_AccessTokenTypeLowerName[6:9]: AccessTokenTypeJWT,
}

var _AccessTokenTypeNames = []string{
	_AccessTokenTypeName[0:6],
	_AccessTokenTypeName[6:9],
}

// AccessTokenTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccessTokenTypeString(s string) (AccessTokenType, error) {
	if val, ok := _AccessTokenTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _AccessTokenTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AccessTokenType values", s)
}

// AccessTokenTypeValues returns all values of the enum
func AccessTokenTypeValues() []AccessTokenType {
	return _AccessTokenTypeValues
}

// AccessTokenTypeStrings returns a slice of all String values of the enum
func AccessTokenTypeStrings() []string {
	strs := make([]string, len(_AccessTokenTypeNames))
	copy(strs, _AccessTokenTypeNames)
	return strs
}

// IsAAccessTokenType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AccessTokenType) IsAAccessTokenType() bool {
	for _, v := range _AccessTokenTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AccessTokenType
func (i AccessTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccessTokenType
func (i *AccessTokenType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccessTokenType should be a string, got %s", data)
	}

	var err error
	*i, err = AccessTokenTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for AccessTokenType
func (i AccessTokenType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for AccessTokenType
func (i *AccessTokenType) UnmarshalText(text []byte) error {
	var err error
	*i, err = AccessTokenTypeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for AccessTokenType
func (i AccessTokenType) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for AccessTokenType
func (i *AccessTokenType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = AccessTokenTypeString(s)
	return err
}

func (i AccessTokenType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *AccessTokenType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of AccessTokenType: %[1]T(%[1]v)", value)
	}

	val, err := AccessTokenTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface for AccessTokenType
func (i AccessTokenType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(i.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface for AccessTokenType
func (i *AccessTokenType) UnmarshalGQL(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("AccessTokenType should be a string, got %T", value)
	}

	var err error
	*i, err = AccessTokenTypeString(str)
	return err
}