package neo4jGenerator

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/ernesto-jimenez/gogen/cleanimports"
	"github.com/ernesto-jimenez/gogen/gogenutil"
	"github.com/ernesto-jimenez/gogen/importer"
	"github.com/ernesto-jimenez/gogen/imports"
	"github.com/gregpechiro/structFields"
)

// Neo4jGenerator will generate the UnmarshalMap function
type Neo4jGenerator struct {
	name       string
	targetName string
	namePkg    string
	pkg        *types.Package
	target     *types.Struct
}

// NewNeo4jGenerator initializes a Neo4jGenerator
func NewNeo4jGenerator(pkg, target string) (*Neo4jGenerator, error) {
	var err error
	if pkg == "" || pkg[0] == '.' {
		pkg, err = filepath.Abs(filepath.Clean(pkg))
		if err != nil {
			return nil, err
		}
		pkg = gogenutil.StripGopath(pkg)
	}
	p, err := importer.Default().Import(pkg)
	if err != nil {
		return nil, err
	}
	obj := p.Scope().Lookup(target)
	if obj == nil {
		return nil, fmt.Errorf("struct %s missing", target)
	}
	if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
		return nil, fmt.Errorf("%s should be an struct, was %s", target, obj.Type().Underlying())
	}
	return &Neo4jGenerator{
		targetName: target,
		pkg:        p,
		target:     obj.Type().Underlying().(*types.Struct),
	}, nil
}

func (g Neo4jGenerator) Fields() []structFields.Field {
	numFields := g.target.NumFields()

	fields := make([]structFields.Field, 0)
	for i := 0; i < numFields; i++ {
		f := structFields.Field{&g, g.target.Tag(i), g.target.Field(i)}

		if f.Field() != "" {
			fields = append(fields, f)
		}
	}
	return fields
}

func (g Neo4jGenerator) Indices() []structFields.Field {
	numFields := g.target.NumFields()

	indices := make([]structFields.Field, 0)
	for i := 0; i < numFields; i++ {
		tag := reflect.StructTag(g.target.Tag(i))
		n := tag.Get("neo4j")
		if n == "" {
			continue
		}
		if strings.Split(n, ",")[0] != "index" {
			continue
		}
		f := structFields.Field{&g, g.target.Tag(i), g.target.Field(i)}

		if f.Field() != "" {
			indices = append(indices, f)
		}
	}
	return indices
}

func (g Neo4jGenerator) Qf(pkg *types.Package) string {
	if g.pkg == pkg {
		return ""
	}
	return pkg.Name()
}

func (g Neo4jGenerator) Name() string {
	name := g.targetName
	return name
}

func (g Neo4jGenerator) Package() string {
	if g.namePkg != "" {
		return g.namePkg
	}
	return g.pkg.Name()
}

func (g *Neo4jGenerator) SetPackage(name string) {
	g.namePkg = name
}

func (g Neo4jGenerator) Imports() map[string]string {
	imports := imports.New(g.Package())
	fields := g.Fields()
	for i := 0; i < len(fields); i++ {
		m := fields[i]
		imports.AddImportsFrom(m.V.Type())
		imports.AddImportsFrom(m.UnderlyingType())
		if sub := m.UnderlyingTarget(); sub != nil {
			fields = append(fields, sub.Fields()...)
		}
	}
	return imports.Imports()
}

func (g Neo4jGenerator) FieldInputs() string {
	var q []string
	for _, field := range g.Fields() {
		q = append(q, fmt.Sprintf("%s:{%s%s}", field.Field(), toLowerFirst(g.Name()), field.Field()))
	}
	return strings.Join(q, ", ")
}

func (g Neo4jGenerator) Write(wr io.Writer) error {
	var buf bytes.Buffer
	if err := fnTmpl.Execute(&buf, g); err != nil {
		return err
	}
	return cleanimports.Clean(wr, buf.Bytes())
}

var (
	tempFuncs = template.FuncMap{
		"toLowerFirst": toLowerFirst,
		"neo4jName":    ToSnake,
	}

	fnTmpl = template.Must(template.New("func").Funcs(tempFuncs).Parse(`/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/neo4jGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package {{ .Package }}

import (
"fmt"
"log"
"io"
"strings"

"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
{{ range $path, $name := .Imports }}
{{ $name }} "{{ $path }}"{{ end }}
)
{{ $fields := .Fields }}{{ $strct := .Name }}{{ $fieldInputs := .FieldInputs }}
var No{{ $strct }}Found = fmt.Errorf("no {{ toLowerFirst $strct }} found")
var Multiple{{ $strct }}Found = fmt.Errorf("multiple {{ toLowerFirst $strct }} found")

func Add{{ $strct }}({{ toLowerFirst $strct }} {{ $strct }}) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE ({{ toLowerFirst $strct }}:{{ $strct }} { {{ $fieldInputs }} })", map[string]interface{}{
		{{ range $fields}}"{{ toLowerFirst $strct }}{{ .Field }}":{{ toLowerFirst $strct }}.{{ .Name }},
		{{ end }}
	})

	return err
}

func GetAll{{ $strct }}() ([]{{ $strct }}, error) {
	var {{ toLowerFirst $strct }}s []{{ $strct }}
	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}) RETURN {{ toLowerFirst $strct }}", nil)
	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return {{ toLowerFirst $strct }}s, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return {{ toLowerFirst $strct }}s, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		{{ toLowerFirst $strct }} := {{ $strct }}{}
		{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
			{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
		}
		{{ end }}
		{{ toLowerFirst $strct }}s = append({{ toLowerFirst $strct }}s, {{ toLowerFirst $strct }})
	}

	return {{ toLowerFirst $strct }}s, nil
}
{{ range $fields }}{{ if .UnderlyingIsBasic }}
{{ if .IsIndex }}func Index{{ $strct }}By{{ .Name }}() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :{{ $strct }}({{ .Field }})", nil)

	return err
}
{{ end }}
func Get{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) ({{ $strct }}, error) {
	{{ toLowerFirst $strct }} := {{ $strct }}{}

	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{ {{ .Field }}:{ {{ .Field }} } }) RETURN {{ toLowerFirst $strct }}", map[string]interface{}{
		"{{ .Field }}":{{ toLowerFirst .Name }},
	})
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return {{ toLowerFirst $strct }}, No{{ $strct }}Found
	}
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return {{ toLowerFirst $strct }}, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
		{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
	}
	{{ end }}

	return {{ toLowerFirst $strct }}, nil
}

func GetOnlyOne{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) ({{ $strct }}, error) {
	{{ toLowerFirst $strct }} := {{ $strct }}{}

	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{ {{ .Field }}:{ {{ .Field }} } }) RETURN {{ toLowerFirst $strct }}", map[string]interface{}{
		"{{ .Field }}":{{ toLowerFirst .Name }},
	})

	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return {{ toLowerFirst $strct }}, No{{ $strct }}Found
	}
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return {{ toLowerFirst $strct }}, Multiple{{ $strct }}Found
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return {{ toLowerFirst $strct }}, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
		{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
	}
	{{ end }}
	return {{ toLowerFirst $strct }}, nil
}

func GetAll{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) ([]{{ $strct }}, error) {
	var {{ toLowerFirst $strct }}s []{{ $strct }}
	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{ {{ .Field }}:{ {{ .Field }} } }) RETURN {{ toLowerFirst $strct }}", map[string]interface{}{
		"{{ .Field }}":{{ toLowerFirst .Name }},
	})

	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return {{ toLowerFirst $strct }}s, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return {{ toLowerFirst $strct }}s, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		{{ toLowerFirst $strct }} := {{ $strct }}{}
		{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
			{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
		}
		{{ end }}
		{{ toLowerFirst $strct }}s = append({{ toLowerFirst $strct }}s, {{ toLowerFirst $strct }})
	}

	return {{ toLowerFirst $strct }}s, nil
}

func UpdateAll{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}, {{ toLowerFirst $strct }} {{ $strct }}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{ {{ .Field }}:{ {{ .Field }} } }) SET {{ toLowerFirst $strct }} += { {{ $fieldInputs }} }", map[string]interface{}{
		"{{ .Field }}":{{ toLowerFirst .Name }},
		{{ range $fields}}"{{ toLowerFirst $strct }}{{ .Field }}":{{ toLowerFirst $strct }}.{{ .Name }},
		{{ end }}
	})
	return err
}

func DeleteAll{{ $strct }}By{{ .Name }}({{ toLowerFirst .Name }} {{ .Type }}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{ {{ .Field }}:{ {{ .Field }} } }) DETACH DELETE {{ toLowerFirst $strct }}", map[string]interface{}{
		"{{ .Field }}":{{ toLowerFirst .Name }},
	})
	return err
}
{{ end }}{{ end }}
func Get{{ $strct }}ByCustom(query map[string]interface{}) ({{ $strct }}, error) {
	{{ toLowerFirst $strct }} := {{ $strct }}{}

	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer conn.Close()

	queryStr := "MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN {{ toLowerFirst $strct }}"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return {{ toLowerFirst $strct }}, No{{ $strct }}Found
	}
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return {{ toLowerFirst $strct }}, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
		{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
	}
	{{ end }}

	return {{ toLowerFirst $strct }}, nil
}

func GetOnlyOne{{ $strct }}ByCustom(query map[string]interface{}) ({{ $strct }}, error) {
	{{ toLowerFirst $strct }} := {{ $strct }}{}

	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer conn.Close()

	queryStr := "MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN {{ toLowerFirst $strct }}"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return {{ toLowerFirst $strct }}, No{{ $strct }}Found
	}
	if err != nil {
		return {{ toLowerFirst $strct }}, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return {{ toLowerFirst $strct }}, Multiple{{ $strct }}Found
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return {{ toLowerFirst $strct }}, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
		{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
	}
	{{ end }}

	return {{ toLowerFirst $strct }}, nil
}

func GetAll{{ $strct }}ByCustom(query map[string]interface{}) ([]{{ $strct }}, error) {
	var {{ toLowerFirst $strct }}s []{{ $strct }}

	conn, err := driver.OpenPool()
	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer conn.Close()

	queryStr := "MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN {{ toLowerFirst $strct }}"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return {{ toLowerFirst $strct }}s, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return {{ toLowerFirst $strct }}s, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return {{ toLowerFirst $strct }}s, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		{{ toLowerFirst $strct }} := {{ $strct }}{}
		{{ range $fields }}if {{ .Field }}, ok := node.Properties["{{ .Field }}"].({{ .Type }}); ok {
			{{ toLowerFirst $strct }}.{{ .Name }} = {{ .Field }}
		}
		{{ end }}
		{{ toLowerFirst $strct }}s = append({{ toLowerFirst $strct }}s, {{ toLowerFirst $strct }})
	}

	return {{ toLowerFirst $strct }}s, nil
}

func UpdateAll{{ $strct }}ByCustom(params map[string]interface{}, {{ toLowerFirst $strct }} {{ $strct }}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET {{ toLowerFirst $strct }} += { {{ $fieldInputs }} }"

	{{ range $fields }}params["{{ toLowerFirst $strct }}{{ .Field }}"] = {{ toLowerFirst $strct }}.{{ .Name }}
	{{ end }}

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAll{{ $strct }}ByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH ({{ toLowerFirst $strct }}:{{ $strct }}{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE {{ toLowerFirst $strct }}"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
`))
)
