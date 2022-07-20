package parser
import "testing"
import "reflect"
import "errors"

var iniTemplate ="[DEFAULT]\nServerAliveInterval = 45\nCompression = yes\nCompressionLevel = 9\nForwardX11 = yes\n[bitbucket.org]\nUser = hg\n[topsecret.server.com]\nPort = 50022\nForwardX11 = no"
// func TestCleanUp(t *testing.T){	

// 	cleaned_up:= cleanupInput(iniTemplate)
// 	expected := "[DEFAULT]\nServerAliveInterval = 45\nCompression = yes\nCompressionLevel = 9\nForwardX11 = yes\n[bitbucket.org]\nUser = hg\n[topsecret.server.com]\nPort = 50022\nForwardX11 = no"
// 	if cleaned_up != expected {
// 		t.Errorf("expected '%s' but got '%s'", expected, cleaned_up)
// 	}	
// }


func TestParse(t *testing.T) {
	
	t.Run("parse with clear input", func(t *testing.T) {
		got ,_ :=Parse(iniTemplate)
		want:= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
	})
	t.Run("parse with empty lines", func(t *testing.T) {
		got ,_ :=Parse("\n\n"+iniTemplate+"\n\n")
		want:= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
		
	})
	t.Run("parse with comments", func(t *testing.T) {
		got ,_ :=Parse(";comment1\n"+iniTemplate+"\n;comment2")
		want:= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
		if  !reflect.DeepEqual(got, want)  {
			t.Errorf("expected '%s' but got '%s'", want,got)
		}	
		
	})
	t.Run("parse with pre spaces and tabs", func(t *testing.T) {
		got ,_ :=Parse("  \t"+iniTemplate)
		want:= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
		"ForwardX11" : "yes"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
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
		var org string =""
		parser2:= Parser{}
		got  := parser2.LoadFromString(org)
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

	// t.Run("not exists file", func(t *testing.T) {
	// 	var org string =""
	// 	parser1:= Parser{}
	// 	got  := parser1.LoadFromFile(org)

	// 	want := errors.New("open /parse.ini: no such file or directory")
	// 	if  !reflect.DeepEqual(got, want) {
	// 		t.Errorf("expected '%s' but got '%s'", want, got)
	// 	}	
	// })
}