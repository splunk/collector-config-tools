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
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
)

type commentReaderFunc func(v reflect.Value) (map[string]string, error)

type commentReader struct {
	dr dirResolver
}

func (r commentReader) commentsForStruct(v reflect.Value) (map[string]string, error) {
	elem := v
	if v.Kind() == reflect.Ptr {
		elem = v.Elem()
	}
	packagePath, err := r.dr.typeToPackagePath(elem.Type())
	if err != nil {
		return nil, err
	}
	return searchDirsForComments(packagePath, elem.Type().String())
}

func searchDirsForComments(packageDir, typeName string) (map[string]string, error) {
	out := map[string]string{}
	err := filepath.WalkDir(packageDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			commentsForStructName(out, path, typeName)
		}
		return nil
	})
	return out, err
}

func commentsForStructName(comments map[string]string, dir, typeName string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	parts := strings.Split(typeName, ".")
	targetPkg := parts[0]
	targetType := parts[1]
	for pkgName, pkg := range pkgs {
		if pkgName != targetPkg {
			continue
		}
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				if gd, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range gd.Specs {
						if ts, ok := spec.(*ast.TypeSpec); ok {
							if ts.Name.Name == targetType {
								if structComments := gd.Doc.Text(); structComments != "" {
									comments["_struct"] = structComments
								}
								if st, ok := ts.Type.(*ast.StructType); ok {
									for _, field := range st.Fields.List {
										if name := fieldName(field); name != "" {
											comments[name] = field.Doc.Text()
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func fieldName(field *ast.Field) string {
	if field.Names != nil {
		return field.Names[0].Name
	} else if se, ok := field.Type.(*ast.SelectorExpr); ok {
		return se.Sel.Name
	} else if id, ok := field.Type.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}
