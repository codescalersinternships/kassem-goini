package parser
import ("testing"
 		"reflect"
 		"errors"
		"os"
		)

var iniTemplate =`[DEFAULT]
ServerAliveInterval = 45
Compression = yes
CompressionLevel = 9
ForwardX11 = yes
[bitbucket.null]
User = hg
[topsecret.server.com]
Port = 50022
ForwardX11 = no`
var null string =""
var wanted_parsedMap= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.null": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}

func TestParse(t *testing.T) {
	
	t.Run("parse with clear input", func(t *testing.T) {
		got ,_ :=Parse(iniTemplate)
		want:= wanted_parsedMap
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
	})
	t.Run("parse with empty lines", func(t *testing.T) {
		got ,_ :=Parse("\n\n"+iniTemplate+"\n\n")
		want:= wanted_parsedMap
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
		
	})
	t.Run("parse with comments", func(t *testing.T) {
		got ,_ :=Parse(";comment1\n"+iniTemplate+"\n;comment2")
		want:= wanted_parsedMap
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
		
	})
	t.Run("parse with pre spaces and tabs", func(t *testing.T) {
		got ,_ :=Parse("  \t"+iniTemplate)
		want:= wanted_parsedMap
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
		
	})

}

func TestLoadFromString(t *testing.T) {
	t.Run("get from clear string", func(t *testing.T) {
		parser1:= Parser{}
		got  := parser1.LoadFromString(iniTemplate)
		
		if  got!=nil  {
			t.Errorf("expected no error but got '%s'",got)
		}	
	})

	t.Run("empty string", func(t *testing.T) {
		parser2:= Parser{}
		got  := parser2.LoadFromString(null)
		want := errors.New("invalid input: empty string")
		if  !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
}

func TestLoadFromFile(t *testing.T){
	t.Run("get from exists file", func(t *testing.T) {
		parser1:= Parser{}
		got  := parser1.LoadFromFile("/parse.ini")
		
		if  got!=nil  {
			t.Errorf("expected no error but got '%s'",got)
		}	
	})

	t.Run("not exists file", func(t *testing.T) {
		
		parser1:= Parser{}
		got  := parser1.LoadFromFile(null)

		want := errors.New("open {"+null+"}: no such file ")
		if  !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
}

func TestGetSectionNames(t *testing.T){
	parser1:= Parser{wanted_parsedMap}
	sections:= parser1.GetSectionNames()
	expected := []string{"DEFAULT","bitbucket.null","topsecret.server.com"}

	if  !contains(expected,sections)   {
		t.Errorf("expected '%s' but got '%s'", expected, sections)
	}	
}

func TestGetSections(t *testing.T) {
	parser1:= Parser{wanted_parsedMap}
	got:= parser1.GetSections()
	want:= wanted_parsedMap
	if  !reflect.DeepEqual(want, got )  {
		t.Errorf("expected '%s' but got '%s'", want, got)
	}	
}
func TestGet(t *testing.T){
	t.Run("get value of [DEFAULT],ServerAliveInterval",func(t *testing.T){
		parser1:= Parser{wanted_parsedMap}
		got := parser1.Get("DEFAULT","ServerAliveInterval")
		want:="45"
		if  got!=want {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
	t.Run("get value of [DEFAULT],null",func(t *testing.T){
		parser1:= Parser{wanted_parsedMap}
		got := parser1.Get("DEFAULT",null)
		want:=""
		if  got!=want {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
	t.Run("get value of null,null",func(t *testing.T){
		parser1:= Parser{wanted_parsedMap}
		got := parser1.Get(null,null)
		want:=""
		if  got!=want {
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
}
func TestSet(t *testing.T) {
	t.Run("value of [DEFAULT],ServerAliveInterval as {test}",func(t *testing.T){
		parser1:= Parser{wanted_parsedMap}
		got := parser1.Set("DEFAULT","ServerAliveInterval","{test}")
		want := map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"{test}","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.null": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
		if  !reflect.DeepEqual(want, got ){
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
	t.Run("value of [DEFAULT],ServerAliveInterval as null",func(t *testing.T){
		parser1:= Parser{wanted_parsedMap}
		got := parser1.Set("DEFAULT","ServerAliveInterval","")
		want := map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.null": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
		if  !reflect.DeepEqual(want, got ){
			t.Errorf("expected '%s' but got '%s'", want, got)
		}	
	})
}
func TestSaveToFile(t *testing.T){
	got:= SaveToFile("test.ini",iniTemplate)
	inputString, _ := os.ReadFile("test.ini")
	if string(inputString)!= iniTemplate || got != nil{
			t.Errorf("expected '%s' but got '%s', error: %v",iniTemplate,inputString,got)
	}
}


func contains(expected []string, out []string) bool {
	flag := false
	for _, expected_name := range expected {
		flag=false
		for _, out_name := range out {
			if expected_name == out_name {
				flag=true
				break
			}
		}
	}
	return flag
}