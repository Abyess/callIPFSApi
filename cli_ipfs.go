package main

import (
	"bytes"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"os"
)

var sh *shell.Shell

func main() {
	cli:=CLI{}
	hash := cli.Run()
	fmt.Print(hash)
}

//使用命令行分析
//1.所有的动作交给命令行
//2.上传文件
const Usage = `
	./ipfs upfile 	"xxxxx"		上传文件
	./ipfs downfile	"xxxxx"		下载文件
	./ipfs upstr 	"xxxxx"		上传字符串
	./ipfs downstr	"xxxxx"		下载字符串
`
type CLI struct {

}
//提供一个命令行解析，获取文件名的方法
func (cli *CLI)Run()string{
	cmds :=os.Args
	if len(cmds)<2{
		fmt.Printf(Usage)
		os.Exit(1)
	}
	switch cmds[1] {
	case "upfile":
		fmt.Printf("上传文件，文件名%s\n",cmds[2])
		data:=cmds[2]
		hash:=UploadFileIPFS(data)
		return hash
	case "downfile":
		fmt.Printf("下载文件，目标哈希%s\n",cmds[2])
		data:=cmds[2]
		_=CatFileIPFS(data)
		return ""
	case "upstr":
		fmt.Printf("上传字符，字符串%s\n",cmds[2])
		data:=cmds[2]
		hash:=UploadStrIPFS(data)
		return hash
	case "downstr":
		fmt.Printf("下载字符，目标哈希%s\n",cmds[2])
		data:=cmds[2]
		_=CatStrIPFS(data)
		return ""
	default:
		fmt.Printf("无效命令，请检查")
		fmt.Printf(Usage)
		return ""
	}
}

//数据上传到ipfs
func UploadStrIPFS(str string) string {
	sh = shell.NewShell("localhost:5001")

	hash, err := sh.Add(bytes.NewBufferString(str))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

//文件上传到ipfs
func UploadFileIPFS(str string) string {
	sh = shell.NewShell("localhost:5001")

	hash, err := sh.AddDir(str)
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	return hash
}

//从ipfs下载数据
func CatStrIPFS(hash string) string {
	sh = shell.NewShell("localhost:5001")

	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)

	return string(body)
}

//从ipfs下载数据
func CatFileIPFS(hash string) string {
	sh = shell.NewShell("localhost:5001")

	err := sh.Get(hash,"D:/BlockChain_Project/ipfs file")
	if err != nil {
		fmt.Println(err)
	}
	return ""
}
//func (s *Shell) Get(hash, outdir string) error