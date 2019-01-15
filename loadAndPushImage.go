package main

import (
	"path/filepath"
	"os"
	"fmt"
	"io/ioutil"
	"container/list"
	"strings"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) == 5 {
		fmt.Println("缺少参数，请参考 --help")
		os.Exit(0)
	}

	if strings.Contains(os.Args[1], "--help") {
		fmt.Println("usage:./xxxx args1 args2 true \n " +
			"args1: \t required 仓库地址 \n " +
			"args2: \t required 项目地址 没有请填写nil \n " +
			"true | false: \t required  是否递归当前文件夹下的子文件夹 \n " +
			"push: \t required  是否push,需要写push,不需要写nopush \n " +
			"save: \t required  是否save,需要写save,不需要写nosave \n")
		os.Exit(0)
	}

	if os.Args[1] == "nil" && os.Args[2] == "nil" {
		fmt.Println("不符合官方推荐命名，请手动执行...")
		os.Exit(0)
	}
	path := getCurrentDir()
	fileList := list.New()
	i := readDir(path, fileList)
	for e := i.Front(); e != nil; e = e.Next() {
		tarName := e.Value.(string)
		loadShell := strings.Join([]string{"docker load -i ", tarName}, "")
		fmt.Print("正在执行...")
		fmt.Println(loadShell)
		result := execCommand(loadShell)
		count := strings.Count(result, "Loaded image:")
		if count > 1 {
			result = strings.TrimSpace(result)
			split := strings.Split(result, "\n")
			for k := range split {
				oldImageName := strings.Replace(split[k], "Loaded image:", "", -1)
				var newImageName string
				if strings.Contains(oldImageName, "/") {
					split := strings.Split(oldImageName, "/")
					imageName := split[len(split)-1]
					newImageName = getNewImageName(imageName)
				} else {
					newImageName = getNewImageName(oldImageName)
				}
				oldImageName = strings.TrimSpace(oldImageName)
				newImageName = strings.TrimSpace(newImageName)
				if k == 0 {
					if !(oldImageName == newImageName) {
						tagImage(oldImageName, newImageName)
						deleteImage(oldImageName,newImageName)
					}
					if os.Args[4] == "push" {
						pushImage(newImageName)
					}
					if os.Args[5] == "save" {
						saveImage(newImageName)
					}
				}
				if k > 0 {
					deleteImage(oldImageName,newImageName)
				}
			}
		} else {
			result = strings.TrimSpace(result)
			if strings.Contains(result, "/") {
				oldImageName := strings.Replace(result, "Loaded image:", "", -1)
				split := strings.Split(result, "/")
				imageName := split[len(split)-1]
				//newImageName := os.Args[1] + "/" + os.Args[2] + "/" + imageName
				newImageName := getNewImageName(imageName)
				oldImageName = strings.TrimSpace(oldImageName)
				newImageName = strings.TrimSpace(newImageName)
				if !(oldImageName == newImageName) {
					tagImage(oldImageName, newImageName)
					deleteImage(oldImageName,newImageName)
				}
				if os.Args[4] == "push" {
					pushImage(newImageName)
				}
				if os.Args[5] == "save" {
					saveImage(newImageName)
				}
			} else {
				oldImageName := strings.Replace(result, "Loaded image:", "", -1)
				//newImageName := os.Args[1] + "/" + os.Args[2] + "/" + oldImageName
				newImageName := getNewImageName(oldImageName)
				oldImageName = strings.TrimSpace(oldImageName)
				newImageName = strings.TrimSpace(newImageName)
				if !(oldImageName == newImageName) {
					tagImage(oldImageName, newImageName)
					deleteImage(oldImageName,newImageName)
				}
				if os.Args[4] == "push" {
					pushImage(newImageName)
				}
				if os.Args[5] == "save" {
					saveImage(newImageName)
				}
			}
		}
	}
}

func getNewImageName(imageName string) string {
	var newImageName string
	if os.Args[2] == "nil" {
		newImageName = os.Args[1] + "/" + imageName
	} else {
		newImageName = os.Args[1] + "/" + os.Args[2] + "/" + imageName
	}
	return strings.TrimSpace(newImageName)
}

func tagImage(oldImageName string, newImageName string) {
	tagShell := strings.Join([]string{"docker tag ", oldImageName, newImageName}, " ")
	fmt.Print("正在执行...")
	fmt.Println(tagShell)
	execCommand(tagShell)
}

func pushImage(newImageName string) {
	pushShell := strings.Join([]string{"docker push ", newImageName}, " ")
	fmt.Print("正在执行...")
	fmt.Println(pushShell)
	execCommand(pushShell)
}

func saveImage(newImageName string) {
	replace := strings.Replace(newImageName, "/", "_", -1)
	s := strings.Replace(replace, ":", "_", -1)
	tarName := s + ".tar"
	saveShell := strings.Join([]string{"docker save ", newImageName, "> ", tarName}, " ")
	fmt.Print("正在执行...")
	fmt.Println(saveShell)
	execCommand(saveShell)
}

func deleteImage(imageName,newImageName string) {
	if imageName == newImageName {
		return
	}
	rmShell := strings.Join([]string{"docker rmi ", imageName}, " ")
	fmt.Print("正在执行...")
	fmt.Println(rmShell)
	execCommand(rmShell)
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}

func readDir(path string, fileList *list.List) *list.List {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range infos {
		fileName := file.Name()
		if file.IsDir() {
			if os.Args[3] == "true" {
				readDir(path+"/"+fileName, fileList)
			}
		} else {
			if strings.HasSuffix(fileName, "tar") {
				fileList.PushFront(path + "/" + fileName)
			}
		}
	}
	return fileList
}

func execCommand(shell string) string {
	args := append([]string{"-c"}, shell)
	command := exec.Command(os.Getenv("SHELL"), args...)
	command.SysProcAttr = &syscall.SysProcAttr{}
	outPip, err := command.StdoutPipe()
	defer outPip.Close()
	if err != nil {
		panic(err)
	}
	err = command.Start()
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(outPip)
	if err != nil {
		panic(err)
	}
	return string(all)
}
