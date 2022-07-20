package parser 

import "fmt"
import (
	"bufio"
 	"strings"
	"os"
	"io/ioutil"
	"errors"
	"path/filepath"
)
type Parser struct { 
	nested_map map[string]map[string]string
}

func Parse(ini_input string) (map[string]map[string]string,error)  {
	var tmp,title string
	scanner := bufio.NewScanner(strings.NewReader(ini_input))
	out_ini:= make(map[string]map[string]string)  
	for scanner.Scan() {
		////////////////////
		// clean up input //
		////////////////////
		if scanner.Text() == "" {
			continue
		}
		if strings.HasPrefix(scanner.Text() , ";") {
			 	continue
			}
		
		tmp = strings.TrimLeft(scanner.Text(),"! ||!\t")
		////////////////////
		// pars ini input //
		////////////////////
		if tmp[0] == '[' && tmp[len(tmp)-1] == ']'{
			title = tmp[1:len(tmp)-1]
			strings.TrimSpace(title)
			//create new map section with title
			out_ini[title]= make(map[string]string)
		} else {
			key_val:= strings.Split(tmp,"=")
			//to avoid index out of range
			if len(key_val)==2{
				out_ini[title][strings.TrimSpace(key_val[0])]= strings.TrimSpace(key_val[1])
			}
		}
		
	}
	return out_ini ,nil
}

func (parser *Parser) LoadFromString(inputString string) (err error) {
	if len(inputString)!=0 {
		parser.nested_map, err= Parse(inputString)

	}else{
		
		return errors.New("invalid input: empty string")
	}
	return nil
}

func (parser *Parser) LoadFromFile(filePath string) (err error) {
	abs_path,_ :=filepath.Abs(".")
	inputString, err := ioutil.ReadFile(abs_path+filePath) 
    if err != nil {
        return errors.New("open {"+filePath+"}: no such file ")
    }
	parser.nested_map, err= Parse(string(inputString))
	
	return nil 
}
func (parser *Parser)GetSectionNames() []string {
	res := []string{}
	for title, _ := range  parser.nested_map {
		res = append(res, title)
	}
	return res

}
func (parser *Parser)GetSections() map[string]map[string]string{
	return parser.nested_map
}

func (parser *Parser)Get(section_name string, key string) string {
		return parser.nested_map[section_name][key]
}

func (parser *Parser)Set(section_name string, key string, value string) map[string]map[string]string {
	parser.nested_map[section_name][key] = value
	return parser.nested_map
}

func (parser *Parser)ToString() string {

	ini_string:= ""
	for section, keyAndValue:= range parser.nested_map{
		fmt.Println(section)
		ini_string += "["+section+"]\n"
		for key, value := range keyAndValue{
			ini_string += key + " = "+value+"\n"
		}
	}
	return strings.TrimSuffix(ini_string, "\n")
}


func  SaveToFile(filePath string,ini_string string) (err error) {
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(ini_string)
	return err
}

