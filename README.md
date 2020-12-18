这是一个通过命令行调用程序上传文件到IPFS的demo

运行目录中运行以下命令
	go run cli_ipfs.go
会得到
			./ipfs upfile   "xxxxx"         上传文件
			./ipfs downfile "xxxxx"         下载文件
			./ipfs upstr    "xxxxx"         上传字符串
			./ipfs downstr  "xxxxx"         下载字符串
	exit status 1
	
正确的使用方式是：
	go run cli_ipfs.go upfile "Filename"
	go run cli_ipfs.go	downfile Qm*************
  
  
