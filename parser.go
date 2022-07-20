package parser 

// import "fmt"
import "bufio"
import "strings"


func cleanupInput(ini_input string) string {
	var tmp string
	scanner := bufio.NewScanner(strings.NewReader(ini_input))
	var ini_WithNoComments string
	for scanner.Scan() {
	
		tmp = scanner.Text()
		if scanner.Text() == "" {
			continue
		}
		if strings.HasPrefix(tmp, ";") {
			 	continue
			}
		tmp = strings.TrimLeft(tmp,"! ||!\t")
		ini_WithNoComments += tmp + "\n"
		
	}
	ini_WithNoComments = strings.TrimSuffix(ini_WithNoComments, "\n")
		return ini_WithNoComments
}

func loadString(inputFile string) []string{
	input_list := strings.Fields(inputFile)
	return input_list
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
