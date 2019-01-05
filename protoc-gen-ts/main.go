package main

import (
	"github.com/gogo/protobuf/vanity/command"
	"github.com/golangper/protoc-gen-rorm/plugin"
)

func main() {
	req := command.Read()
	// files := req.GetProtoFile()
	// vanity.ForEachFile(files, vanity.TurnOffGogoImport)
	response := command.GeneratePlugin(req, &plugin.TsPlugin{}, ".pb.service.ts")
	for _, file := range response.GetFile() {

		file.Content = plugin.CleanImportsTs(file.Content)
	}
	command.Write(response)
}
