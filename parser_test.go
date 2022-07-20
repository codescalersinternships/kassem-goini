package parser
import "testing"
import "reflect"
var iniTemplate ="\n;sec1\n[DEFAULT]\nServerAliveInterval = 45\n  \tCompression = yes\nCompressionLevel = 9\n;hello \nForwardX11 = yes\n[bitbucket.org]\n;comments\nUser = hg\n[topsecret.server.com]\nPort = 50022\nForwardX11 = no"
func TestCleanUp(t *testing.T){	

	cleaned_up:= cleanupInput(iniTemplate)
	expected := "[DEFAULT]\nServerAliveInterval = 45\nCompression = yes\nCompressionLevel = 9\nForwardX11 = yes\n[bitbucket.org]\nUser = hg\n[topsecret.server.com]\nPort = 50022\nForwardX11 = no"
	if cleaned_up != expected {
		t.Errorf("expected '%s' but got '%s'", expected, cleaned_up)
	}	
}

func TestGetSectionNames(t *testing.T){
	input:= cleanupInput(iniTemplate)
	sections:= GetSectionNames(loadString(input))
	expected := []string{"DEFAULT","bitbucket.org","topsecret.server.com"}
	if  !reflect.DeepEqual(sections, expected )  {
		t.Errorf("expected '%s' but got '%s'", expected, sections)
	}	
}

func TestGetSection(t *testing.T) {
	input:= cleanupInput(iniTemplate)
	got:=GetSections(loadString(input))
	want:= map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
	"ForwardX11" : "yes"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
	if  !reflect.DeepEqual(got, want)  {
		t.Errorf("expected '%s' but got '%s'", want,got)
	}	
}
func TestGet(t *testing.T) {
	input:= cleanupInput(iniTemplate)
	got:=Get(GetSections(loadString(input)),"DEFAULT","ServerAliveInterval")
	want :="45"
	if  !reflect.DeepEqual(got, want)  {
		t.Errorf("expected '%s' but got '%s'", want,got)
	}

}
func TestSet(t *testing.T) {
	input:= cleanupInput(iniTemplate)
	got:=Set(GetSections(loadString(input)),"DEFAULT","test","15")
	want :=map[string]map[string]string{"DEFAULT" : {"ServerAliveInterval":"45","Compression":"yes","CompressionLevel" : "9",
	"ForwardX11" : "yes","test":"15"},  "bitbucket.org": {"User" : "hg"}, "topsecret.server.com":{"Port":"50022","ForwardX11": "no"}}
	if  !reflect.DeepEqual(got, want)  {
		t.Errorf("expected '%s' but got '%s'", want,got)
	}

}