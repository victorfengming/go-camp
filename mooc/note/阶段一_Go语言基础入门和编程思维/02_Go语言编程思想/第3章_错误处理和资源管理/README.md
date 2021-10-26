




# 3-1 defer调用.mp4

```go
func writeFile(filename string){
	file,err:=os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer:= bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer,string(i))
	}
}
```



等我代码写完了,我就忘记这个了

所以我就进就加一个这个



![1635231971395](README/1635231971395.png)

# 3-2 错误处理概念.mp4



# 3-3 服务器统一出错处理_浏览器需放大_.mp4



# 3-4 服务器统一出错处理.mp4

