// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configschema // import "github.com/open-telemetry/opentelemetry-collector-contrib/lib/configschema"

import (
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/structtag"
)

// cfgField holds attributes and subfields of a config struct.
type cfgField struct {
	Name      string      `yaml:",omitempty"`
	Type      string      `yaml:",omitempty"`
	Kind      string      `yaml:",omitempty"`
	Default   interface{} `yaml:",omitempty"`
	Doc       string      `yaml:",omitempty"`
	CfgFields []*cfgField `yaml:"fields,omitempty"`
}

// readFields accepts both a config struct's Value, as well as a commentReaderFunc,
// and returns a cfgField pointer for the top level struct as well as all of its
// recursive subfields.
func readFields(v reflect.Value, commentReader commentReaderFunc) (*cfgField, error) {
	cfgType := v.Type()
	field := &cfgField{
		Type: cfgType.String(),
	}
	err := refl(field, v, commentReader)
	return field, err
}

func refl(field *cfgField, v reflect.Value, commentReader commentReaderFunc) error {
	if v.Kind() == reflect.Ptr {
		err := refl(field, v.Elem(), commentReader)
		if err != nil {
			return err
		}
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	comments, err := commentReader(v)
	if err != nil {
		return err
	}

	// _struct comments are those that are on the struct type itself. Here we check
	// if field.Doc is empty, thus preventing a squashed type with struct comments
	// from overwriting the containing struct's comments.
	if sc, ok := comments["_struct"]; ok && field.Doc == "" {
		field.Doc = sc
	}

	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		if !structField.IsExported() {
			continue
		}
		tagName, options, err := mapstructure(structField.Tag)
		if err != nil {
			log.Printf("error parsing mapstructure tag for type: %s: %s: %v", field.Type, structField.Tag, err)
			// not fatal, can keep going
		}
		if tagName == "-" {
			continue
		}
		fv := v.Field(i)
		next := field
		if !containsSquash(options) {
			name := tagName
			if name == "" {
				name = strings.ToLower(structField.Name)
			}
			kindStr := fv.Kind().String()
			typeStr := fv.Type().String()
			if typeStr == kindStr {
				typeStr = "" // omit if redundant
			}
			next = &cfgField{
				Name: name,
				Type: typeStr,
				Kind: kindStr,
				Doc:  comments[structField.Name],
			}
			field.CfgFields = append(field.CfgFields, next)
		}
		err = handleKind(fv, next, commentReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleKind(v reflect.Value, f *cfgField, commentReader commentReaderFunc) (err error) {
	switch v.Kind() {
	case reflect.Struct:
		err = refl(f, v, commentReader)
	case reflect.Ptr:
		if v.IsNil() {
			err = refl(f, reflect.New(v.Type().Elem()), commentReader)
		} else {
			err = refl(f, v.Elem(), commentReader)
		}
	case reflect.Slice:
		e := v.Type().Elem()
		if e.Kind() == reflect.Struct {
			err = refl(f, reflect.New(e), commentReader)
		} else if e.Kind() == reflect.Ptr {
			err = refl(f, reflect.New(e.Elem()), commentReader)
		}
	case reflect.String:
		f.Default = v.String()
	case reflect.Bool:
		f.Default = v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() != 0 {
			if v.Type() == reflect.TypeOf(time.Duration(0)) {
				f.Default = time.Duration(v.Int()).String()
			} else {
				f.Default = v.Int()
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		f.Default = v.Uint()
	}
	return
}

func mapstructure(st reflect.StructTag) (string, []string, error) {
	tag := string(st)
	if tag == "" {
		return "", nil, nil
	}
	tags, err := structtag.Parse(tag)
	if err != nil {
		return "", nil, err
	}
	ms, err := tags.Get("mapstructure")
	if err != nil {
		return "", nil, err
	}
	return ms.Name, ms.Options, nil
}

func containsSquash(options []string) bool {
	for _, option := range options {
		if option == "squash" {
			return true
		}
	}
	return false
}
