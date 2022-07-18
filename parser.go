package cleanupInput 


import "bufio"
import "strings"

var tmp string
func cleanupInput(ini_input string) string {
	scanner := bufio.NewScanner(strings.NewReader(ini_input))
	var ini_WithNoComments string
	for scanner.Scan() {
	
		tmp = scanner.Text()
		if scanner.Text() == "" {
			continue
		}
		tmp = strings.TrimPrefix(tmp," ")
		ini_WithNoComments += tmp + "\n"
		
	}
	ini_WithNoComments = strings.TrimSuffix(ini_WithNoComments, "\n")
		return ini_WithNoComments
}
