// generate.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generate.go <layer> <name>")
		return
	}

	layer := os.Args[1]               // 如 "web"
	name := os.Args[2]                // 如 "channel"
	structName := strings.Title(name) // 将名称首字母大写

	// 定义模板路径和输出文件路径
	tmplPath := fmt.Sprintf(".templates/%s.go.tpl", layer)
	outputPath := fmt.Sprintf("internal/%s/%s.go", layer, name)

	// 检查模板文件是否存在
	if _, err := os.Stat(tmplPath); os.IsNotExist(err) {
		fmt.Printf("Template %s for layer '%s' does not exist.\n", tmplPath, layer)
		return
	}

	// 自定义模板函数
	funcMap := template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
	}

	// 解析模板文件并添加自定义函数
	tmpl, err := template.New(filepath.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	// 创建输出文件
	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// 执行模板并写入文件
	err = tmpl.Execute(f, map[string]string{
		"StructName": structName,
	})
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Printf("Generated %s\n", outputPath)
}
