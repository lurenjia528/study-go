package main

import (
	"regexp"
	"strings"
)

func main() {
	GetConfig()
}
func GetConfig() {
	imageString := "kylincloud2.hub/kube-system/virt-operator:v0.18.1"
	imageRegEx := regexp.MustCompile(`^(.*)/virt-operator(:.*)?$`)
	matches := imageRegEx.FindAllStringSubmatch(imageString, 1)
	//fmt.Println(matches)
	registry := matches[0][1]
	tag := strings.TrimPrefix(matches[0][2], ":")
	if tag == "" {
		tag = "latest"
	}
	println(registry)
}
