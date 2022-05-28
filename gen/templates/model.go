//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	{{- $fmt := false }}{{ range .Attributes}}{{ if or (eq .Id true) (eq .Reference true) (eq .Type "List") (and (eq .Type "Bool") (eq .TypeYangBool "empty")) }}{{ $fmt = true }}{{ end}}{{ end}}
	{{- if $fmt }}
	"fmt"
	{{- end}}
	{{- $strconv := false }}{{ range .Attributes}}{{ if or (and (eq .Type "Int64") (ne .Id true) (ne .Reference true)) (eq .Type "List")}}{{ $strconv = true }}{{ end}}{{ end}}
	{{- if $strconv }}
	"strconv"
	{{- end}}
	{{- $reflect := false }}{{ range .Attributes}}{{ if eq .Type "List" }}{{ $reflect = true }}{{ end}}{{ end}}
	{{- if $reflect }}
	"reflect"
	{{- end}}

	"github.com/hashicorp/terraform-plugin-framework/types"
	{{- $sjson := false }}{{ range .Attributes}}{{ if and (ne .Id true) (ne .Reference true)}}{{ $sjson = true }}{{ end}}{{ end}}
	{{- if $sjson }}
	"github.com/tidwall/sjson"
	{{- end}}
	{{- $gjson := false }}{{ range .Attributes}}{{ if and (ne .Id true) (ne .Reference true) (ne .WriteOnly true)}}{{ $gjson = true }}{{ end}}{{ end}}
	{{- if $gjson }}
	"github.com/tidwall/gjson"
	{{- end}}
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}

{{- range .Attributes}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
}
{{- end}}
{{- end}}

func (data {{camelCase .Name}}) getPath() string {
{{- if hasId .Attributes}}
	return fmt.Sprintf("{{.Path}}"{{range .Attributes}}{{if or (eq .Id true) (eq .Reference true)}}, data.{{toGoName .TfName}}.Value{{end}}{{end}})
{{- else}}
	return "{{.Path}}"
{{- end}}
}

func (data {{camelCase .Name}}) toBody() string {
	body := "{}"
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .Id true) (ne .Type "List")}}
	if !data.{{toGoName .TfName}}.Null && !data.{{toGoName .TfName}}.Unknown {
		{{- if eq .Type "Int64"}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", strconv.FormatInt(data.{{toGoName .TfName}}.Value, 10))
		{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
		if data.{{toGoName .TfName}}.Value {
			body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", map[string]string{})
		}
		{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.Value)
		{{- else if eq .Type "String"}}
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", data.{{toGoName .TfName}}.Value)
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- range .Attributes}}
	{{- if eq .Type "List"}}
	{{- $list := toJsonPath .YangName .XPath }}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{toJsonPath .YangName .XPath}}", []interface{}{})
		for index, item := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			if !item.{{toGoName .TfName}}.Null && !item.{{toGoName .TfName}}.Unknown {
				{{- if eq .Type "Int64"}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", strconv.FormatInt(item.{{toGoName .TfName}}.Value, 10))
				{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
				if item.{{toGoName .TfName}}.Value {
					body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", map[string]string{})
				}
				{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", item.{{toGoName .TfName}}.Value)
				{{- else if eq .Type "String"}}
				body, _ = sjson.Set(body, "{{$list}}"+"."+strconv.Itoa(index)+"."+"{{toJsonPath .YangName .XPath}}", item.{{toGoName .TfName}}.Value)
				{{- end}}
			}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) updateFromBody(res []byte) {
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .Id true) (ne .WriteOnly true)}}
	{{- if eq .Type "Int64"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.Int()
	} else {
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else if eq .Type "Bool"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}}.Value = value.Bool()
		{{- else}}
		data.{{toGoName .TfName}}.Value = true
		{{- end}}
	} else {
		data.{{toGoName .TfName}}.Value = false
	}
	{{- else if eq .Type "String"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.String()
	} else {
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else if eq .Type "List"}}
	{{- $list := (toGoName .TfName)}}
	{{- $listPath := (toJsonPath .YangName .XPath)}}
	{{- $yangKey := ""}}
	for i := range data.{{$list}} {
		keys := [...]string{ {{range .Attributes}}{{if eq .Id true}}"{{.YangName}}", {{end}}{{end}} }
		keyValues := [...]string{ {{range .Attributes}}{{if eq .Id true}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.Value, 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.Value), {{else}}data.{{$list}}[i].{{toGoName .TfName}}.Value, {{end}}{{end}}{{end}} }

		var r gjson.Result
		gjson.GetBytes(res, "{{$listPath}}").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)

		{{- range .Attributes}}
		{{- if ne .WriteOnly true}}
		{{- if eq .Type "Int64"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.Int()
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- else if eq .Type "Bool"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() {
			{{- if eq .TypeYangBool "boolean"}}
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.Bool()
			{{- else}}
			data.{{$list}}[i].{{toGoName .TfName}}.Value = true
			{{- end}}
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = false
		}
		{{- else if eq .Type "String"}}
		if value := r.Get("{{toJsonPath .YangName .XPath}}"); value.Exists() {
			data.{{$list}}[i].{{toGoName .TfName}}.Value = value.String()
		} else {
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}

	{{- end}}
}

func (data *{{camelCase .Name}}) fromBody(res []byte) {
	{{- range .Attributes}}
	{{- if and (ne .Reference true) (ne .Id true) (ne .WriteOnly true)}}
	{{- if eq .Type "Int64"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.Int()
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "Bool"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		{{- if eq .TypeYangBool "boolean"}}
		data.{{toGoName .TfName}}.Value = value.Bool()
		{{- else}}
		data.{{toGoName .TfName}}.Value = true
		{{- end}}
		data.{{toGoName .TfName}}.Null = false
	} else {
		data.{{toGoName .TfName}}.Value = false
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "String"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}}.Value = value.String()
		data.{{toGoName .TfName}}.Null = false
	}
	{{- else if eq .Type "List"}}
	if value := gjson.GetBytes(res, "{{toJsonPath .YangName .XPath}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if ne .WriteOnly true}}
			if cValue := v.Get("{{toJsonPath .YangName .XPath}}"); cValue.Exists() {
				{{- if eq .Type "Int64"}}
				item.{{toGoName .TfName}}.Value = cValue.Int()
				{{- else if and (eq .Type "Bool") (eq .TypeYangBool "boolean")}}
				item.{{toGoName .TfName}}.Value = cValue.Bool()
				{{- else if and (eq .Type "Bool") (ne .TypeYangBool "boolean")}}
				item.{{toGoName .TfName}}.Value = true
				{{- else if eq .Type "String"}}
				item.{{toGoName .TfName}}.Value = cValue.String()
				{{- end}}
				item.{{toGoName .TfName}}.Null = false
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}


func (data *{{camelCase .Name}}) fromPlan(plan {{camelCase .Name}}) {
	data.Device = plan.Device
	{{- range .Attributes}}
	{{- if or (eq .Reference true) (eq .Id true) (eq .WriteOnly true)}}
	data.{{toGoName .TfName}}.Value = plan.{{toGoName .TfName}}.Value
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) setUnknownValues() {
	if data.Device.Unknown {
		data.Device.Unknown = false
		data.Device.Null = true
	}
	if data.Id.Unknown {
		data.Id.Unknown = false
		data.Id.Null = true
	}
	{{- range .Attributes}}
	{{- if ne .Type "List"}}
	if data.{{toGoName .TfName}}.Unknown {
		data.{{toGoName .TfName}}.Unknown = false
		data.{{toGoName .TfName}}.Null = true
	}
	{{- else}}
	{{- $list := (toGoName .TfName)}}
	for i := range data.{{$list}} {
		{{- range .Attributes}}
		if data.{{$list}}[i].{{toGoName .TfName}}.Unknown {
			data.{{$list}}[i].{{toGoName .TfName}}.Unknown = false
			data.{{$list}}[i].{{toGoName .TfName}}.Null = true
		}
		{{- end}}
	}
	{{- end}}
	{{- end}}
}

func (data *{{camelCase .Name}}) getDeletedListItems(state {{camelCase .Name}}) []string {
	deletedListItems := make([]string, 0)
	{{- range .Attributes}}
	{{- if eq .Type "List"}}
	{{- $goKey := ""}}
	{{- range .Attributes}}
	{{- if eq .Id true}}
	{{- $goKey = (toGoName .TfName)}}
	{{- end}}
	{{- end}}
	for i := range state.{{toGoName .TfName}} {
		{{- $list := (toGoName .TfName)}}
		keys := [...]string{ {{range .Attributes}}{{if eq .Id true}}"{{.YangName}}", {{end}}{{end}} }
		stateKeyValues := [...]string{ {{range .Attributes}}{{if eq .Id true}}{{if eq .Type "Int64"}}strconv.FormatInt(state.{{$list}}[i].{{toGoName .TfName}}.Value, 10), {{else if eq .Type "Bool"}}strconv.FormatBool(state.{{$list}}[i].{{toGoName .TfName}}.Value), {{else}}state.{{$list}}[i].{{toGoName .TfName}}.Value, {{end}}{{end}}{{end}} }
		
		emptyKeys := true
		{{- range .Attributes}}
		{{- if eq .Id true}}
		if !reflect.ValueOf(state.{{$list}}[i].{{toGoName .TfName}}.Value).IsZero() {
			emptyKeys = false
		}
		{{- end}}
		{{- end}}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.{{toGoName .TfName}} {
			found = true
			{{- range .Attributes}}
			{{- if eq .Id true}}
			if state.{{$list}}[i].{{toGoName .TfName}}.Value != data.{{$list}}[j].{{toGoName .TfName}}.Value {
				found = false
			} 
			{{- end}}
			{{- end}}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "["+keys[ki]+"="+stateKeyValues[ki]+"]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/{{.YangName}}%v", state.getPath(), keyString))
		}
	}
	{{- end}}
	{{- end}}
	return deletedListItems
}

func (data *{{camelCase .Name}}) getEmptyLeafsDelete() []string {
	emptyLeafsDelete := make([]string, 0)
	{{- range .Attributes}}
	{{- if and (eq .Type "Bool") (eq .TypeYangBool "empty")}}
	if !data.{{toGoName .TfName}}.Value {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/{{.YangName}}", data.getPath()))
	}
	{{- end}}
	{{- if eq .Type "List"}}
	{{- $hasEmpty := false}}
	{{ range .Attributes}}{{ if and (eq .Type "Bool") (eq .TypeYangBool "empty")}}{{ $hasEmpty = true}}{{ end}}{{- end}}
	{{- if $hasEmpty}}
	{{- $yangName := .YangName}}
	keys := [...]string{ {{range .Attributes}}{{if eq .Id true}}"{{.YangName}}", {{end}}{{end}} }
	for i := range data.{{toGoName .TfName}} {
		{{- $list := (toGoName .TfName)}}
		keyValues := [...]string{ {{range .Attributes}}{{if eq .Id true}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.Value, 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.Value), {{else}}data.{{$list}}[i].{{toGoName .TfName}}.Value, {{end}}{{end}}{{end}} }
		keyString := ""
		for ki := range keys {
			keyString += "["+keys[ki]+"="+keyValues[ki]+"]"
		}
		{{- range .Attributes}}
		{{- if and (eq .Type "Bool") (eq .TypeYangBool "empty")}}
		if !data.{{$list}}[i].{{toGoName .TfName}}.Value {
			emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/{{$yangName}}%v/{{.YangName}}", data.getPath(), keyString))
		}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return emptyLeafsDelete
}
