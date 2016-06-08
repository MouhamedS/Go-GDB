

package main

	

import (
    //"bufio"
    "fmt"
    //"io"
    "io/ioutil"
    "os"
	"path/filepath"
	"log"
	"gopkg.in/qml.v1"
)

	

func check(e error) {
    if e != nil {
        panic(e)
    }
}

	

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[1]))
   	 if err != nil {
            log.Fatal(err)
    	}

    dat, err := ioutil.ReadFile(dir  + "/" + os.Args[1])
    check(err)
    fmt.Print(string(dat))



    f, err := os.Open(dir  + "/" + os.Args[1])
    check(err)


/*
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))


	

    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))


	

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))


	

    _, err = f.Seek(0, 0)
    check(err)



    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

*/

	

    f.Close()

if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}	

}


func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}
	//context := engine.Context()
    
	


	window := controls.CreateWindow(nil)

	window.Show()
	window.Wait()
	return nil
}

type File struct {
	Name string
	Content string
}

func (file *File) Control(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[1]))
   	 if err != nil {
            log.Fatal(err)
    	}

    dat, err := ioutil.ReadFile(dir  + "/" + os.Args[1])
    check(err)
   // fmt.Print(string(dat))
	go func() {file.Content =string(dat)
	qml.Changed(file, &file.Content)	
	}()
}
