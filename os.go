/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: SystemService
 Date: 2021/3/23 2:38 AM
*/
package system

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// 获取当前执行程序的绝对目录路径
func GetCurrentDir() string {
	currentPath := CurrentAndAbsPath()
	return filepath.Dir(currentPath)
}

// 当前执行程序的绝对文件路径
func GetCurrentFilename() string {
	return CurrentAndAbsPath()
}

// 当前执行程序的绝对路径
func CurrentAndAbsPath() string {
	current := SetCurrentPath()
	return GetAbsPath(current)
}

// 设置当前执行程序的相对路径
func SetCurrentPath() string {
	current := os.Args[0]
	path, err := exec.LookPath(current)
	if err != nil {
		log.Println("Set the current path error: ", err)
	}

	return path
}

// 获取当前执行程序的绝对路径
func GetAbsPath(current string) string {
	absPath, err := filepath.Abs(current)
	if err != nil {
		log.Println("Get the current absolute of path error: ", err)
	}

	return absPath
}

// 获取配置目录
func GetConfigDir() string {
	rootPath := GetRootPath()
	return filepath.Join(rootPath, "config", string(os.PathSeparator))
}

// 获取路径的上个目录
func GetLastPath(currentPath string) string {
	index := strings.LastIndex(currentPath, string(os.PathSeparator))
	return currentPath[:index]
}

// 获取项目根目录
func GetRootPath() string {
	dir := GetCurrentDir()
	rootPath := GetLastPath(dir)
	return filepath.Join(rootPath, string(os.PathSeparator))
}

// 环境变量解析: 根据环境变量的值替换字符串中的 ${var} or $var, 如果不存在任何环境变量, 则使用空字符串替换
func ParseEnvVar(varString string) string {
	return os.ExpandEnv(varString)
}

// Calls the OS default browser for open uri
func CallBrowser(uri string) error {
	browser := make(map[string]string)

	browser["windows"] = "start"
	browser["darwin"] = "open"
	browser["linux"] = "xdg-open"

	getOS := GetOS()
	run, ok := browser[getOS]
	if !ok {
		log.Println("Unknown the OS:", getOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

// Get OS types.
func GetOS() string {
	return runtime.GOOS
}

// Get absolute path by directory name.
func GetDirPath(dirname string) string {
	rootPath := GetRootPath()
	return filepath.Join(rootPath, dirname, string(os.PathSeparator))
}

// Get absolute path by directory name and file name.
func GetFilepath(dirname string, filename string) string {
	configDir := GetDirPath(dirname)
	return filepath.Join(configDir, filename)
}
