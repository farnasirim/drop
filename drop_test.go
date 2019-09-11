package drop

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtoGeneration(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "digest_go_mock")
	assert.Nil(t, err)
	defer func() {
		os.Remove(tmpDir)
	}()

	protoGenCommandTemplate := `protoc -I proto/ --go_out=plugins=grpc:%s proto/drop.proto`
	protoGenCommandBuilt := fmt.Sprintf(protoGenCommandTemplate, tmpDir)
	protoGenCommandArgs := strings.Split(protoGenCommandBuilt, " ")
	protoGenCmd := exec.Command(protoGenCommandArgs[0], protoGenCommandArgs[1:]...)

	err = protoGenCmd.Start()
	assert.Nil(t, err)

	err = protoGenCmd.Wait()
	assert.Nil(t, err)

	newContents, err := ioutil.ReadFile(path.Join(tmpDir, "drop.pb.go"))
	assert.Nil(t, err)

	origContents, err := ioutil.ReadFile("proto/drop.pb.go")
	assert.Nil(t, err)

	assert.Equal(t, origContents, newContents)
}
