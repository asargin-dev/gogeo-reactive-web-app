package utils

import (
	"io/ioutil"
)

//Takes html dir name as parameter to redesign HTML
func HtmlDesigner (htmlPath string,htmlName string) {

	sourceHtml, err := ioutil.ReadFile(htmlPath)

	if err != nil {
		panic(err)
	}

	addTempKeys := string(sourceHtml)

	// Go template keywords to use in each other
	addTempKeys=`{{define "`+htmlName+`"}}`+ addTempKeys +`{{end}}`
	byteTypeOfSource:=[]byte(addTempKeys)

	//Redesign source code of HTML to create charts for selected parameters
	ioutil.WriteFile(htmlPath,byteTypeOfSource,0644)


}
