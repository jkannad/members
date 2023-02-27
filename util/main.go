package main

import (
	//"github.com/jkannad/spas/members/internal/service"
	"html/template"
	"log"
	"os"
	"strings"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type TemplateData struct {
	Data map[string]Person
}

func main() {
	/*err := service.LoadCountries()
	if err != nil {
		log.Println(err)
	}

	err = service.LoadStates()
	if err != nil {
		log.Println(err)
	} 

	err := service.LoadCities()
	if err != nil {
		log.Println(err)	
	} 

	err := service.LoadDialCodes()
	if err != nil {
		log.Println(err)
	} */

	//multipleDbInsert()
	//TestTemplate()
	var inter map[string]interface{}

	slice := []string{"text1", "text2", "text3"}
	inter["slice"] = slice

	//mp := map[string]string{"key1":"val1", "key2":"val2"}
	//inter["mp"] = mp

	sl := inter["slice"]
	for index, val := range sl {
		fmt.Printf("Index:%d Value:%s", index, val)
	}
}

func multipleDbInsert() {

	per1 := Person{
		Name:"Janarthanan",
		Age:45,
	}

	per2 := Person{
		Name:"Vaidegi",
		Age:44,
	}

	per3 := Person{
		Name:"Ashwin",
		Age:14,
	}

	persons := []Person{per1, per2, per3}


	query := "INSERT INTO product(product_name, product_price) VALUES "
    var inserts []string
    var params []interface{}
    for _, v := range persons {
        inserts = append(inserts, "(?, ?)")
        params = append(params, v.Name, v.Age)
    }
    queryVals := strings.Join(inserts, ",")
    query = query + queryVals
	fmt.Println(params)
    log.Println("query is", query)
}



func TestTemplate() {
	var td TemplateData
	data := make(map[string]Person)
	data["P1"] = Person{Name:"Jana", Age:45}
	data["P2"] = Person{Name:"Vaidegi", Age:44}
	data["P3"] = Person{Name:"Ashwin", Age:14}
	td.Data = data
	
	tmpl := template.New("test")

	//parse some content and generate a template
	tmpl, err := tmpl.Parse(`{{range $key, $value := .Data}}{{$key}}Name: {{$value.Name}} Age: {{$value.Age}}{{end}}`)
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	err1 := tmpl.Execute(os.Stdout,td)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)

	}
}

func TestTemplateIF() {
	var td TemplateData
	data := make(map[string]Person)
	data["P1"] = Person{Name:"Jana", Age:45}
	data["P2"] = Person{Name:"Vaidegi", Age:44}
	data["P3"] = Person{Name:"Ashwin", Age:14}
	td.Data = data
	
	tmpl := template.New("test")

	//parse some content and generate a template
	tmpl, err := tmpl.Parse(`{{range $key, $value := .Data}}{{$key}}Name: {{$value.Name}} Age: {{$value.Age}}{{end}}`)
	if err != nil {
		log.Fatal("Error Parsing template: ", err)
		return
	}
	err1 := tmpl.Execute(os.Stdout,td)
	if err1 != nil {
		log.Fatal("Error executing template: ", err1)

	}
}