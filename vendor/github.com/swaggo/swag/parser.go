package swag

import (
	"go/ast"
	goparser "go/parser"
	"go/token"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"regexp"
	"strconv"

	"github.com/go-openapi/jsonreference"
	"github.com/go-openapi/spec"
)

// Parser implements a parser for Go source files.
type Parser struct {
	// swagger represents the root document object for the API specification
	swagger *spec.Swagger

	//files is a map that stores map[real_go_file_path][astFile]
	files map[string]*ast.File

	// TypeDefinitions is a map that stores [package name][type name][*ast.TypeSpec]
	TypeDefinitions map[string]map[string]*ast.TypeSpec

	//registerTypes is a map that stores [refTypeName][*ast.TypeSpec]
	registerTypes map[string]*ast.TypeSpec
}

// New creates a new Parser with default properties.
func New() *Parser {
	parser := &Parser{
		swagger: &spec.Swagger{
			SwaggerProps: spec.SwaggerProps{
				Info: &spec.Info{
					InfoProps: spec.InfoProps{
						Contact: &spec.ContactInfo{},
						License: &spec.License{},
					},
				},
				Paths: &spec.Paths{
					Paths: make(map[string]spec.PathItem),
				},
				Definitions: make(map[string]spec.Schema),
			},
		},
		files:           make(map[string]*ast.File),
		TypeDefinitions: make(map[string]map[string]*ast.TypeSpec),
		registerTypes:   make(map[string]*ast.TypeSpec),
	}
	return parser
}

// ParseApi parses general api info for gived searchDir and mainApiFile
func (parser *Parser) ParseApi(searchDir string, mainApiFile string) {
	log.Println("Generate general API Info")
	parser.getAllGoFileInfo(searchDir)
	parser.ParseGeneralApiInfo(path.Join(searchDir, mainApiFile))

	for _, astFile := range parser.files {
		parser.ParseType(astFile)
	}

	for _, astFile := range parser.files {
		parser.ParseRouterApiInfo(astFile)
	}

	parser.ParseDefinitions()
}

// ParseGeneralApiInfo parses general api info for gived mainApiFile path
func (parser *Parser) ParseGeneralApiInfo(mainApiFile string) {
	fileSet := token.NewFileSet()
	fileTree, err := goparser.ParseFile(fileSet, mainApiFile, nil, goparser.ParseComments)

	if err != nil {
		log.Panicf("ParseGeneralApiInfo occur error:%+v", err)
	}

	parser.swagger.Swagger = "2.0"

	if fileTree.Comments != nil {
		for _, comment := range fileTree.Comments {
			for _, commentLine := range strings.Split(comment.Text(), "\n") {
				attribute := strings.ToLower(strings.Split(commentLine, " ")[0])
				switch attribute {
				case "@version":
					parser.swagger.Info.Version = strings.TrimSpace(commentLine[len(attribute):])
				case "@title":
					parser.swagger.Info.Title = strings.TrimSpace(commentLine[len(attribute):])
				case "@description":
					parser.swagger.Info.Description = strings.TrimSpace(commentLine[len(attribute):])
				case "@termsofservice":
					parser.swagger.Info.TermsOfService = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.name":
					parser.swagger.Info.Contact.Name = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.email":
					parser.swagger.Info.Contact.Email = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.url":
					parser.swagger.Info.Contact.URL = strings.TrimSpace(commentLine[len(attribute):])
				case "@license.name":
					parser.swagger.Info.License.Name = strings.TrimSpace(commentLine[len(attribute):])
				case "@license.url":
					parser.swagger.Info.License.URL = strings.TrimSpace(commentLine[len(attribute):])
				case "@host":
					parser.swagger.Host = strings.TrimSpace(commentLine[len(attribute):])
				case "@basepath":
					parser.swagger.BasePath = strings.TrimSpace(commentLine[len(attribute):])
				case "@schemes":
					parser.swagger.Schemes = GetSchemes(commentLine)
				}
			}
		}
	}
}

//自定义

// ParseApi parses general api info for gived searchDir and mainApiFile
func (parser *Parser) ParseApiCustom(searchDir string, mainApiFile string, host string) {
	log.Println("Generate general API Info")
	parser.getAllGoFileInfo(searchDir)
	parser.ParseGeneralApiInfoCustom(path.Join(searchDir, mainApiFile), host)

	for _, astFile := range parser.files {
		parser.ParseType(astFile)
	}

	for _, astFile := range parser.files {
		parser.ParseRouterApiInfo(astFile)
	}

	parser.ParseDefinitions()
}

// ParseGeneralApiInfo parses general api info for gived mainApiFile path
func (parser *Parser) ParseGeneralApiInfoCustom(mainApiFile string, host string) {
	fileSet := token.NewFileSet()
	fileTree, err := goparser.ParseFile(fileSet, mainApiFile, nil, goparser.ParseComments)

	if err != nil {
		log.Panicf("ParseGeneralApiInfo occur error:%+v", err)
	}

	parser.swagger.Swagger = "2.0"

	if fileTree.Comments != nil {
		for _, comment := range fileTree.Comments {
			for _, commentLine := range strings.Split(comment.Text(), "\n") {
				attribute := strings.ToLower(strings.Split(commentLine, " ")[0])
				switch attribute {
				case "@version":
					parser.swagger.Info.Version = strings.TrimSpace(commentLine[len(attribute):])
				case "@title":
					parser.swagger.Info.Title = strings.TrimSpace(commentLine[len(attribute):])
				case "@description":
					parser.swagger.Info.Description = strings.TrimSpace(commentLine[len(attribute):])
				case "@termsofservice":
					parser.swagger.Info.TermsOfService = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.name":
					parser.swagger.Info.Contact.Name = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.email":
					parser.swagger.Info.Contact.Email = strings.TrimSpace(commentLine[len(attribute):])
				case "@contact.url":
					parser.swagger.Info.Contact.URL = strings.TrimSpace(commentLine[len(attribute):])
				case "@license.name":
					parser.swagger.Info.License.Name = strings.TrimSpace(commentLine[len(attribute):])
				case "@license.url":
					parser.swagger.Info.License.URL = strings.TrimSpace(commentLine[len(attribute):])
				case "@basepath":
					parser.swagger.BasePath = strings.TrimSpace(commentLine[len(attribute):])
				case "@schemes":
					parser.swagger.Schemes = GetSchemes(commentLine)
				case "@host":
					if host == "" {
						parser.swagger.Host = strings.TrimSpace(commentLine[len(attribute):])
					} else {
						parser.swagger.Host = host
					}
				}
			}
		}
	}
}

// GetSchemes parses swagger schemes for gived commentLine
func GetSchemes(commentLine string) []string {
	attribute := strings.ToLower(strings.Split(commentLine, " ")[0])
	return strings.Split(strings.TrimSpace(commentLine[len(attribute):]), " ")
}

// ParseRouterApiInfo parses router api info for gived astFile
func (parser *Parser) ParseRouterApiInfo(astFile *ast.File) {
	for _, astDescription := range astFile.Decls {
		switch astDeclaration := astDescription.(type) {
		case *ast.FuncDecl:
			if astDeclaration.Doc != nil && astDeclaration.Doc.List != nil {
				operation := NewOperation() //for per 'function' comment, create a new 'Operation' object
				operation.parser = parser
				for _, comment := range astDeclaration.Doc.List {
					if err := operation.ParseComment(comment.Text); err != nil {
						log.Panicf("ParseComment panic:%+v", err)
					}
				}
				var pathItem spec.PathItem
				var ok bool

				if pathItem, ok = parser.swagger.Paths.Paths[operation.Path]; !ok {
					pathItem = spec.PathItem{}
				}
				switch strings.ToUpper(operation.HttpMethod) {
				case http.MethodGet:
					pathItem.Get = &operation.Operation
				case http.MethodPost:
					pathItem.Post = &operation.Operation
				case http.MethodDelete:
					pathItem.Delete = &operation.Operation
				case http.MethodPut:
					pathItem.Put = &operation.Operation
				case http.MethodPatch:
					pathItem.Patch = &operation.Operation
				case http.MethodHead:
					pathItem.Head = &operation.Operation
				case http.MethodOptions:
					pathItem.Options = &operation.Operation
				}

				parser.swagger.Paths.Paths[operation.Path] = pathItem
			}
		}
	}
}

// ParseType parses type info for gived astFile
func (parser *Parser) ParseType(astFile *ast.File) {
	if _, ok := parser.TypeDefinitions[astFile.Name.String()]; !ok {
		parser.TypeDefinitions[astFile.Name.String()] = make(map[string]*ast.TypeSpec)
	}

	for _, astDeclaration := range astFile.Decls {
		if generalDeclaration, ok := astDeclaration.(*ast.GenDecl); ok && generalDeclaration.Tok == token.TYPE {
			for _, astSpec := range generalDeclaration.Specs {
				if typeSpec, ok := astSpec.(*ast.TypeSpec); ok {
					parser.TypeDefinitions[astFile.Name.String()][typeSpec.Name.String()] = typeSpec
				}
			}
		}
	}
}

// ParseDefinitions parses Swagger Api definitions
func (parser *Parser) ParseDefinitions() {
	for refTypeName, typeSpec := range parser.registerTypes {
		ss := strings.Split(refTypeName, ".")
		pkgName := ss[0]
		parser.ParseDefinition(pkgName, typeSpec, typeSpec.Name.Name)
	}
}

// Build
func (parser *Parser) ParseDefinition(pkgName string, typeSpec *ast.TypeSpec, typeName string) {
	var refTypeName string
	if len(pkgName) > 0 {
		refTypeName = pkgName + "." + typeName
	} else {
		refTypeName = typeName
	}
	if _, already := parser.swagger.Definitions[refTypeName]; already {
		log.Println("Skipping '" + refTypeName + "', already present.")
		return
	}
	properties := make(map[string]spec.Schema)
	required := make([]string, 0)
	parser.parseTypeSpec(pkgName, typeSpec, properties, &required)

	for _, prop := range properties {
		// todo find the pkgName of the property type
		tname := prop.SchemaProps.Type[0]
		if _, ok := parser.TypeDefinitions[pkgName][tname]; ok {
			tspec := parser.TypeDefinitions[pkgName][tname]
			parser.ParseDefinition(pkgName, tspec, tname)
		}
	}

	log.Println("Generating " + refTypeName)
	parser.swagger.Definitions[refTypeName] = spec.Schema{
		SchemaProps: spec.SchemaProps{
			Type:       []string{"object"},
			Properties: properties,
			Required:   required,
		},
	}
}

func (parser *Parser) parseTypeSpec(pkgName string, typeSpec *ast.TypeSpec, properties map[string]spec.Schema, requiredStrs *[]string) {
	switch typeSpec.Type.(type) {
	case *ast.StructType:
		structDecl := typeSpec.Type.(*ast.StructType)
		fields := structDecl.Fields.List

		for _, field := range fields {
			if field.Names == nil { //anonymous field
				parser.parseAnonymousField(pkgName, field, properties)
			} else {
				name, schemaType, arrayType := parser.parseField(field)
				// if defined -- ref it
				if _, ok := parser.TypeDefinitions[pkgName][schemaType]; ok {
					properties[name] = spec.Schema{
						SchemaProps:
						spec.SchemaProps{Type: []string{schemaType},
							Ref: spec.Ref{
								Ref: jsonreference.MustCreateRef("#/definitions/" + pkgName + "." + schemaType),
							},
						},
					}
				} else if schemaType == "array" {
					// if defined -- ref it
					if _, ok := parser.TypeDefinitions[pkgName][arrayType]; ok {
						parser.ParseDefinition(pkgName, parser.TypeDefinitions[pkgName][arrayType], arrayType)
						properties[name] = spec.Schema{
							SchemaProps: spec.SchemaProps{
								Type:  []string{schemaType},
								Items: &spec.SchemaOrArray{Schema: &spec.Schema{SchemaProps: spec.SchemaProps{Ref: spec.Ref{Ref: jsonreference.MustCreateRef("#/definitions/" + pkgName + "." + arrayType)}}}},
							},
						}
					}
				} else {
					example, description, required := parseFieldTag(field.Tag.Value)
					properties[name] = spec.Schema{
						SchemaProps: spec.SchemaProps{
							Type:        []string{schemaType},
							Description: description,
						},
					}
					if example != "" {
						schema := properties[name]
						schema.SwaggerSchemaProps = spec.SwaggerSchemaProps{
							Example: example,
						}
					}
					if required && requiredStrs != nil {
						*requiredStrs = append(*requiredStrs, name)
					}
				}
			}
		}

	case *ast.ArrayType:
		log.Panic("ParseDefinitions not supported 'Array' yet.")
	case *ast.InterfaceType:
		log.Panic("ParseDefinitions not supported 'Interface' yet.")
	case *ast.MapType:
		log.Panic("ParseDefinitions not supported 'Map' yet.")
	}
}

func parseFieldTag(tag string) (example, description string, required bool) {
	re, err := regexp.Compile(`example:\"([\S]*)\"[\s]?`)
	if err == nil {
		exampleTag := re.FindStringSubmatch(tag)
		if len(exampleTag) > 0 {
			example = exampleTag[1]
		}
	}
	re, err = regexp.Compile(`description:\"([\S]*)\"[\s]?`)
	if err == nil {
		descriptionTag := re.FindStringSubmatch(tag)
		if len(descriptionTag) > 0 {
			description = descriptionTag[1]
		}
	}

	re, err = regexp.Compile(`required:\"([\S]*)\"[\s]?`)
	if err == nil {
		requiredTag := re.FindStringSubmatch(tag)
		if len(requiredTag) > 0 {
			required, err = strconv.ParseBool(requiredTag[1])
		}
	}
	return
}

func (parser *Parser) parseAnonymousField(pkgName string, field *ast.Field, properties map[string]spec.Schema) {
	if astTypeIdent, ok := field.Type.(*ast.Ident); ok {
		findPgkName := pkgName
		findBaseTypeName := astTypeIdent.Name
		ss := strings.Split(astTypeIdent.Name, ".")
		if len(ss) > 1 {
			findPgkName = ss[0]
			findBaseTypeName = ss[1]
		}

		baseTypeSpec := parser.TypeDefinitions[findPgkName][findBaseTypeName]
		parser.parseTypeSpec(findPgkName, baseTypeSpec, properties, nil)
	}
}

func (parser *Parser) parseField(field *ast.Field) (propName, schemaType string, arrayType string) {
	schType, arrType := getPropertyName(field)
	re, err := regexp.Compile(`json:\"([\\s\\S]+)\"`)
	if err != nil {
		log.Panic(err)
	}
	var fieldJsonName string
	jsonTag := re.FindStringSubmatch(field.Tag.Value)
	if len(jsonTag) > 0 {
		fieldJsonName = jsonTag[1]
	} else {
		proName := field.Names[0].Name
		fieldJsonName = strings.ToLower(proName[0:1]) + proName[1:]
	}
	CheckSchemaType(schType)
	return fieldJsonName, schType, arrType
}

// GetAllGoFileInfo gets all Go source files information for gived searchDir.
func (parser *Parser) getAllGoFileInfo(searchDir string) {
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		//exclude vendor folder
		if ext := filepath.Ext(path); ext == ".go" && !strings.Contains(string(os.PathSeparator)+path, string(os.PathSeparator)+"vendor"+string(os.PathSeparator)) {
			fset := token.NewFileSet() // positions are relative to fset
			astFile, err := goparser.ParseFile(fset, path, nil, goparser.ParseComments)

			if err != nil {
				log.Panicf("ParseFile panic:%+v", err)
			}

			parser.files[path] = astFile

		}
		return nil
	})
}

// GetSwagger returns *spec.Swagger which is the root document object for the API specification.
func (parser *Parser) GetSwagger() *spec.Swagger {
	return parser.swagger
}
