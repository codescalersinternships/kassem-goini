package TestCleanUp 
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
}