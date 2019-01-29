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
	"flag"
)

var (
	reg     string
	pro     string
	r       bool
	push    bool
	save    bool
	h       bool
	v       bool
	VERSION string
	ARCH    string
)

func init() {
	flag.StringVar(&reg, "registry", "ygt", "镜像仓库名")
	flag.StringVar(&pro, "project", "", "镜像仓库项目名")
	flag.BoolVar(&r, "recursive", false, "是否递归子目录")
	flag.BoolVar(&push, "push", true, "是否推送到镜像仓库")
	flag.BoolVar(&save, "save", false, "是否保存修改后的镜像")
	flag.BoolVar(&h, "help", false, "帮助信息")
	flag.BoolVar(&v, "version", false, "版本信息")
	flag.Usage = usage
}

func usage() {
	fmt.Print(`遍历当前文件夹及其子文件夹,上传,推送,保存镜像
  --registry string
        required | 镜像仓库名 (default "kylincloud2.hub")
  --project string
        required | 镜像仓库项目名
  --recursive
        是否递归子目录
  --push
        是否推送到镜像仓库 (default true)
  --save
        是否保存修改后的镜像
  --version
        版本信息
  --help
        帮助信息
`)
}

func main() {
	flag.Parse()
	if h {
		usage()
		os.Exit(1)
	}

	if v {
		fmt.Printf("VERSION: %s\n", VERSION)
		fmt.Printf("ARCH: %s\n", ARCH)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}

	if reg == "" || pro == "" {
		fmt.Println("不符合官方推荐命名,请手动执行...")
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
						deleteImage(oldImageName, newImageName)
					}

					if push {
						pushImage(newImageName)
					}

					if save {
						saveImage(newImageName)
					}
				}
				if k > 0 {
					deleteImage(oldImageName, newImageName)
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
					deleteImage(oldImageName, newImageName)
				}

				if push {
					pushImage(newImageName)
				}

				if save {
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
					deleteImage(oldImageName, newImageName)
				}

				if push {
					pushImage(newImageName)
				}

				if save {
					saveImage(newImageName)
				}
			}
		}
	}
}

func getNewImageName(imageName string) string {
	var newImageName string
	if pro == "" {
		newImageName = reg + "/" + imageName
	} else {
		newImageName = reg + "/" + pro + "/" + imageName
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

func deleteImage(imageName, newImageName string) {
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
			if r {
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
