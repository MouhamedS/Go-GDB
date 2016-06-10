package main

import (
    "fmt"
    "gopkg.in/qml.v1"
    "os"
    "io/ioutil"    
    "path/filepath"
    "log"
)

func main() {

	


    if err := qml.Run(run); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {

	dir, err := filepath.Abs(filepath.Dir(os.Args[1]))
   	 if err != nil {
            log.Fatal(err)
    	}

	

    dat, err2 := ioutil.ReadFile(dir + "/" + os.Args[1])
    if err2 != nil {
            log.Fatal(err2)
    	}

    
   
    engine := qml.NewEngine()

    component, err := engine.LoadFile(dir + "/main.qml")
    if err != nil {
        return err
    }

    context := engine.Context()
    context.SetVar("fileOp",&File{Content : string(dat)})

    window := component.CreateWindow(nil)
	

    window.Show()
    window.Wait()

    return nil
}

type File struct {
    Name    string
    Content string
}
 /*
func (file *File) DebugRun() string {  
        GDBrun() 
}

*/


