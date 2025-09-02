## Golang String Obfuscator

```
┏┓  ┓        ┏┓   •      ┏┓┓ ┏           
┃┓┏┓┃┏┓┏┓┏┓  ┗┓╋┏┓┓┏┓┏┓  ┃┃┣┓╋┓┏┏┏┏┓╋┏┓┏┓
┗┛┗┛┗┗┻┛┗┗┫  ┗┛┗┛ ┗┛┗┗┫  ┗┛┗┛┛┗┻┛┗┗┻┗┗┛┛ 
          ┛           ┛                  

Golang String Obfuscator                  
version: 1.0.0
made by love: RUSTY J0y
------------------------------------------------------------
```


I’ve been coding in C++ for years and recently switched to Go.

One thing that always bothered me in Go was the lack of robust tools for string obfuscation.

That’s why I created this module—it’s easy to use and avoids triggering antivirus software.

Below, I’ll show you how to use it

***
#### Step 1:
Download and build this repository.

```cmd
go build -o RustyJ0yObf.exe
```
***
#### Step 2:
Pass your desired text as an argument to the program.

The program will give you an output similar to this:

```cmd
C:\Users\RustyJ0y\Desktop> RustyJ0yObf.exe del /f /q %temp%\*.*

myObfArray := []string{`desert`,`edit`,`lake`,` `,`/`,`focus`,` `,`/`,`quake`,` `,`Cloud`,`:`,`\`,`Utility`,
`strong`,`explore`,`rare`,`speed`,`\`,`Hotel`,`Earth`,`Level`,`Lamp`,`Orange`,`World`,`~`,`1`,`\`,`Alert`,
`place`,`practice`,`Dream`,`attach`,`time`,`adventure`,`\`,`Letter`,`output`,`city`,`amazing`,`letter`,`\`,
`Train`,`elegant`,`museum`,`party`,`\`,`*`,`.`,`*`}
```

Copy the printed array and hardcode it into your program.

***
#### Step 3:
Now we need to retrieve the original string at runtime!

For this, use the following function in your code:


```go

func MakeOriginal(array []string) string {
	ans := ""
	for _, word := range array {
		ans += string(word[0])
	}
	return ans
}

```

***

Below is an example of using this function as well:

```go
package main

import (
	"fmt"
	"os/exec"
)

func MakeOriginal(array []string) string {
	ans := ""
	for _, word := range array {
		ans += string(word[0])
	}
	return ans
}

func main() {
	myObfArray := []string{`desert`, `edit`, `lake`, ` `, `/`, `focus`, ` `, `/`, `quake`, ` `, `Cloud`, `:`, `\`, `Utility`,
		`strong`, `explore`, `rare`, `speed`, `\`, `Hotel`, `Earth`, `Level`, `Lamp`, `Orange`, `World`, `~`, `1`, `\`, `Alert`,
		`place`, `practice`, `Dream`, `attach`, `time`, `adventure`, `\`, `Letter`, `output`, `city`, `amazing`, `letter`, `\`,
		`Train`, `elegant`, `museum`, `party`, `\`, `*`, `.`, `*`}

	myCommand := MakeOriginal(myObfArray)

	cmd := exec.Command("cmd", "/C", myCommand)
	err := cmd.Run()
	if err != nil {
		fmt.Println("[!] Command Execution Error:", err)
	}
	fmt.Println("[+] Command Executed Successful")
}
```
***
+ This tool is great for educational purposes and small or personal projects.
+ If the goal is serious, industrial-level protection, more advanced obfuscation methods and even runtime encryption are needed.
+ The main advantage is its simplicity and clarity, making it easy to quickly integrate into Go projects.

---

```
       ┓   ┓    ┓      
 ┏┳┓┏┓┏┫┏┓ ┣┓┓┏ ┃┏┓┓┏┏┓
 ┛┗┗┗┻┗┻┗  ┗┛┗┫ ┗┗┛┗┛┗ 
              ┛         
  ┳┓┳┳┏┓┏┳┓┓┏  ┏┳┏┓┓┏     
  ┣┫┃┃┗┓ ┃ ┗┫   ┃┃┫┗┫     
  ┛┗┗┛┗┛ ┻ ┗┛  ┗┛┗┛┗┛ 
```
