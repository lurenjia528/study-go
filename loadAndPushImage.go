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
   path := getCurrentDir()
   i := readDir(path)
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
            pushShell := strings.Replace(split[k], "Loaded image:", "docker push", -1)
            fmt.Print("正在执行...")
            fmt.Println(pushShell)
            execCommand(pushShell)
         }
      } else {
         pushShell := strings.Replace(result, "Loaded image:", "docker push", -1)
         fmt.Print("正在执行...")
         fmt.Println(pushShell)
         execCommand(pushShell)
      }
   }
}

// 获取当前目录
func getCurrentDir() string {
   dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
   if err != nil {
      panic(err)
   }
   return dir
}

// 递归读取当前目录下所有tar文件
func readDir(path string) *list.List {
   fileList := list.New()
   infos, err := ioutil.ReadDir(path)
   if err != nil {
      panic(err)
   }
   for _, file := range infos {
      fileName := file.Name()
      if file.IsDir() {
         readDir(fileName)
      } else {
         if strings.HasSuffix(fileName, "tar") {
            fileList.PushFront(fileName)
         }
      }
   }
   return fileList
}

// 执行shell命令
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
