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
        return err
    }
	parser.nested_map, err= Parse(string(inputString))
	for k, v := range parser.nested_map {
		fmt.Println(k)
		for o,m := range v {
			fmt.Println(o,m)
		}
	}
	return nil 
}
func GetSectionNames(input_list []string) []string {
	res := []string{}
		for _,item := range input_list {
			if item[0] == '[' && item[len([]rune(item))-1] == ']' {
				res = append(res,item[1:len([]rune(item))-1])
			}
		}
		return res
}
func GetSections(input_list []string) map[string]map[string]string{
	//map of maps
	sections:= map[string]map[string]string{}
	for index,item := range input_list{
		if item[0] == '[' && item[len([]rune(item))-1] == ']'{
			section:= item[1:len([]rune(item))-1]
			sectionMap:=map[string]string{}
			//contains section's keys and values
			for j:= index+1; j<len(input_list); j++{
				if strings.HasPrefix(input_list[j],"[") && strings.HasSuffix(input_list[j],"]") {
					index = j - 1
					break
				}
				if input_list[j] == "="{
					sectionMap[input_list[j-1]]= input_list[j+1]
				}
			}
		sections[section]= sectionMap
	
		}	
	}
	return sections
}

func Get(sections map[string]map[string]string, section_name string, key string) string {
		value := sections[section_name][key]
		return value
}
func Set(sections map[string]map[string]string, section_name string, key string, value string) map[string]map[string]string {
	sections[section_name][key] = value
	return sections
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