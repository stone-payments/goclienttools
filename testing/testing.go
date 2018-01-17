package basetesting

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/gomega"
)

//File load a file verifying possible errors
func File(path string) []byte {
	blob, err := ioutil.ReadFile(path)
	Expect(err).Should(BeNil())
	return blob
}

//LoadJSON loads a json verifying possible errors
func LoadJSON(path string, target interface{}) {
	blob := File(path)
	err := json.Unmarshal(blob, &target)
	Expect(err).Should(BeNil())
}
