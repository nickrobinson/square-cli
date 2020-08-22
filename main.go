package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego/swagger"
	"github.com/iancoleman/strcase"
)

func main() {
	jsonFile, err := os.Open("api/spec/api.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var swag swagger.Swagger
	json.Unmarshal(byteValue, &swag)

	for path, item := range swag.Paths {
		fmt.Println(path)
		if item.Get != nil {
			fmt.Println(strcase.ToKebab(item.Get.OperationID))
			if item.Get.Parameters != nil {
				for _, param := range item.Get.Parameters {
					fmt.Printf("\t%s (%s)\n", param.Name, param.Type)
				}
			}
		}

		if item.Post != nil {
			fmt.Println(strcase.ToKebab(item.Post.OperationID))
		}

		if item.Delete != nil {
			fmt.Println(strcase.ToKebab(item.Delete.OperationID))
		}
	}

}
