// Code generated by yaml_to_go. DO NOT EDIT.
// source: g2_compressed.yaml

package spectest

type MsgHashCompressedTest struct {
	Input struct {
		Message string `json:"message"`
		Domain  string `json:"domain"`
	} `json:"input"`
	Output []string `json:"output"`
}
