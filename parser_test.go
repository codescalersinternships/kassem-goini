package cleanupInput
import "testing"
const expectedOut =`[DEFAULT]
ServerAliveInterval = 45
Compression = yes
CompressionLevel = 9
ForwardX11 = yes
[bitbucket.org]
User = hg
[topsecret.server.com]
Port = 50022
ForwardX11 = no`
func TestCleanUp(t *testing.T){
	
	t.Run("impty lines",func(t *testing.T){ 
		var iniTemplate =`
[DEFAULT]
ServerAliveInterval = 45
Compression = yes
CompressionLevel = 9


ForwardX11 = yes
[bitbucket.org]

User = hg
[topsecret.server.com]
Port = 50022
ForwardX11 = no`
	noComments:= cleanupInput(iniTemplate)
	expected := expectedOut
	if noComments != expected {
		t.Errorf("expected '%s' but got '%s'", expected, noComments)
	}
})

	t.Run("with comments",func(t *testing.T){ 
	var iniTemplate =`
;sec1
[DEFAULT]
ServerAliveInterval = 45
Compression = yes
CompressionLevel = 9
;hello 
ForwardX11 = yes
[bitbucket.org]
;comments
User = hg
[topsecret.server.com]
Port = 50022
ForwardX11 = no`
	noComments:= cleanupInput(iniTemplate)
	expected := expectedOut
	if noComments != expected {
		t.Errorf("expected '%s' but got '%s'", expected, noComments)
	}
})
t.Run("with pre white spaces and tab",func(t *testing.T){ 
	var iniTemplate =`
;sec1
		[DEFAULT]
ServerAliveInterval = 45
  Compression = yes
	 CompressionLevel = 9
;hello 
ForwardX11 = yes
[bitbucket.org]
;comments
User = hg
[topsecret.server.com]
Port = 50022
ForwardX11 = no`
	noComments:= cleanupInput(iniTemplate)
	expected := expectedOut
	if noComments != expected {
		t.Errorf("expected '%s' but got '%s'", expected, noComments)
	}
})
		
}