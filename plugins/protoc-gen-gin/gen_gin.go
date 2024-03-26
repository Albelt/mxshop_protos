package main

import (
	"bytes"
	_ "embed"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"html/template"
)

var importPaths = []protogen.GoImportPath{
	protogen.GoImportPath("github.com/gin-gonic/gin"),
	protogen.GoImportPath("net/http"),
}

//go:embed ginTemplate.gohtml
var templateString string //模板文件嵌入到变量

type GenGin struct{}

func (g *GenGin) GenerateFile(plugin *protogen.Plugin, file *protogen.File) error {
	if len(file.Services) == 0 {
		return nil
	}

	// 生成文件对象
	filename := file.GeneratedFilenamePrefix + "_gin.pb.go"
	f := plugin.NewGeneratedFile(filename, file.GoImportPath)

	// 开头说明
	f.P("// Code generated by protoc-gen-gin. DO NOT EDIT.")
	f.P()
	f.P("package ", file.GoPackageName)
	f.P()

	// 添加import
	for _, path := range importPaths {
		f.QualifiedGoIdent(path.Ident(""))
	}

	// 处理service(一个proto文件中最多一个service)
	text, err := g.GenerateService(file.Services[0])
	if err != nil {
		return err
	}
	f.P(text)

	return nil
}

func (g *GenGin) GenerateService(service *protogen.Service) (string, error) {
	if len(service.Methods) == 0 {
		return "", nil
	}

	srvDesc := serviceDesc{
		ServiceName: service.GoName,
	}

	// 从注解中获取http方法的信息
	for _, m := range service.Methods {
		rule, ok := proto.GetExtension(m.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)
		if rule != nil && ok {
			srvDesc.Methods = append(srvDesc.Methods, buildHttpRule(m, rule))
		}
	}

	// 使用模板渲染
	return g.RenderTmpl(&srvDesc)
}

func (g *GenGin) RenderTmpl(s *serviceDesc) (string, error) {
	tmpl, err := template.New("html").Parse(templateString)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, s); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func buildHttpRule(m *protogen.Method, rule *annotations.HttpRule) *methodDesc {
	var path, method string

	switch pattern := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		path = pattern.Get
		method = "GET"
	case *annotations.HttpRule_Post:
		path = pattern.Post
		method = "POST"
	case *annotations.HttpRule_Delete:
		path = pattern.Delete
		method = "DELETE"
	case *annotations.HttpRule_Put:
		path = pattern.Put
		method = "PUT"
	}

	return &methodDesc{
		Name:     m.GoName,
		Method:   method,
		Path:     path,
		Request:  m.Input.GoIdent.GoName,
		Response: m.Output.GoIdent.GoName,
	}
}

type serviceDesc struct {
	ServiceName string
	Methods     []*methodDesc
}

type methodDesc struct {
	Name     string //rpc方法名称
	Method   string //http方法
	Path     string //http路径
	Request  string //rpc请求名称
	Response string //rpc响应名称
}
