package main

import (
	"bufio"
	"fmt"
	"github.com/kataras/iris"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CodeModel struct {
	Code string
	Desc string
}

var MdFile = map[string]string{"CodeMd":"code.md"}


var PathSeparator string

func main() {
	//检查目标目录是否存在
	if os.IsPathSeparator('/') {
		PathSeparator = "/"
	} else {
		PathSeparator = "\\"
	}


	app := iris.Default()
	tmpl := iris.HTML( "./views" , ".html")
	tmpl.Reload(true)
	app.RegisterView( tmpl )

	app.Handle("GET" , "/" , func( ctx iris.Context ){
		code := ctx.Params().Get("code")

		var  filePath string = GetCurrentDirectory() + PathSeparator + MdFile["CodeMd"]
		if isExists( filePath ) {
			fmt.Println( filePath + " exists")
		} else {
			ctx.HTML("not thing found")
			_  , err := os.Create( filePath)
			if err != nil {
				log.Fatal( err )
				panic( "init file error")
			}
		}
		mdFile , err := os.OpenFile( filePath , os.O_RDWR | os.O_CREATE , 0 )

		if err != nil {
			log.Fatalln( "read md file fail" , err )
		}
		defer mdFile.Close()

		//line := make([]byte , 4096 )
		lineRd := bufio.NewReader( mdFile)
		for {
			line , err := lineRd.ReadString( '\n' )
			line = strings.TrimSpace( line )
			if err != nil {
				if err == io.EOF {
					fmt.Println("file read ok")
					break
				} else {
					panic("file read error")
				}
				//close
			}
			if code != "" {
				//如果code不为空
			} else {

			}
			fmt.Println(line)
		}


		ctx.View("search.html")
	})

	app.Get( "/ping" , func( ctx iris.Context ){
		ctx.WriteString("pong")
	})

	app.Get("/hello" , func( ctx iris.Context) {
		ctx.JSON( iris.Map{"message":"Hello iris web frameword."})
	})

	app.Run( iris.Addr(":8080"))
	iris.WithoutServerError( iris.ErrServerClosed )

	//a := blackfriday.Run( input )
}

func createDir ( dir string ) (bool , error) {
	dirPath := filepath.Dir( dir )
	if !isExists( dirPath ) {
		err := os.MkdirAll( dirPath , os.ModePerm)
		if err != nil {
			return false , err
		}
	}
	return true , nil
}

/**
 * 文件是否存在
 */
func isExists(path string) bool {
	path = strings.Replace( path , "//" , "/" , 0 )
	fmt.Println( path )
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}


func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}