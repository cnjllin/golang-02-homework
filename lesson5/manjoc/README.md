![](https://i.imgur.com/JfxuYKM.png)

**上：ide， 下： github**  

<pre>
// 支持重定向
func redirectCommand(line string) {
	// fmt.Println("redirect command.")
	lineList := strings.Split(line, ">")
	cmdLine := lineList[0]
	outFileName := strings.Fields(lineList[1])[0]
	outPutFile, _ := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	args := strings.Fields(cmdLine)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = outPutFile
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
</pre>
