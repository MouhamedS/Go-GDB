				package main

				import (
					  "fmt"
					  "github.com/cyrus-and/gdb"
					  "io"
					  "os"
					  //"path/filepath"
					"strconv"
				)
					var debug *gdb.Gdb
				func main() {


					  // start a new instance and pipe the target output to stdout
					  debug, _ = gdb.New(nil)
					  go io.Copy(os.Stdout, debug)

			

					  // load and run a program
					  debug.Send("file-exec-and-symbols", os.Args[1])
					fmt.Println(debug.Send("gdb-version"))

					 
					  var input string
					  
				 for input != "quit"   {
					  fmt.Println("Que voulez-vous faire ?")
					  fmt.Scanln(&input)
					
						switch input {
						
						//Break
						case "break": breake()
						

						//Break List						
						case "break-list" : breaklist()

						
						//Break delete						
						case "delete" : delete_break()

						//Run	
						case "run" : start()
					
						//Step 
						case "step" : step()

						//Reverse Stepping
						case "step-reverse" :step_reverse()
						
						// Continue
						case "continue" :  continuee()

						//Reverse continue 
						case "continue-reverse" : continue_reverse()


						//Print
						case "print" : print()

						//List variables locals
						case "list-variables" : list_variables ()

						
						//Backtrace

						case "backtrace" :fmt.Println(backtrace())

						//Watchpoints
						case "watch" : watch()

						//Where
						case "where" : where ()
						
						//quit
						case "quit":
						//stop
						case "stop" : stop()
						// Default Case					
						default: fmt.Println("Commandes non valides")  

					}
					  
				}	

					  
	}

	func start(){
					
				fmt.Println(debug.Send("exec-run"))
				debug.Send("interpreter-exec","console","record")	

			}
	func step(){
		
		 output, err := debug.Send("exec-step")
			if err != nil {
						fmt.Println(err)		
						}
					
							notif := output["class"]
							fmt.Println("Notification : ", notif) 
		
	}
	func breaklist(){

			output,err :=debug.Send("break-list")
						if err !=nil {
							fmt.Println(err)				
							}

						pay:=output["payload"]
						payAssert := pay.(map[string]interface {})
					
						breakpointTable := payAssert["BreakpointTable"]
						breakpointTableAssert := breakpointTable.(map[string]interface {})
					
						Array := breakpointTableAssert["body"]
						ArrayAssert := Array.([]interface{})
						nbreVar:=len(ArrayAssert)
					
						for i:=0; i<=nbreVar-1 ; i++{
							mapSepare := ArrayAssert[i]
							mapSepareAssert := mapSepare.(map[string]interface {})
							bkpt := mapSepareAssert["bkpt"]
							bkptAssert := bkpt.(map[string]interface {})
							number:=bkptAssert["number"]
							typeB :=bkptAssert["type"]
							enabled :=bkptAssert["enabled"]
							times :=bkptAssert["times"]
							disp :=bkptAssert["disp"]
							fun :=bkptAssert["func"]
							line :=bkptAssert["line"]
							fmt.Println("number:",number,"type:",typeB,"enabled:",enabled,"times:",times,"disp:",disp,"function:",fun,"line:",line)
						}


		}
	func delete_break(){
				
						var numero_break string		
				
						fmt.Println("Supprimer un breakpoint(n°) ou tous les breakpoints")
						fmt.Scanln(&numero_break )
						
						if numero_break != "" {
							debug.Send("break-delete",numero_break )
						} else {
							debug.Send("break-delete")
							}	

	}

	func step_reverse(){
		output,err := debug.Send("exec-step","--reverse")
								if err != nil {
									fmt.Println(err)		
									}
					
							notif := output["class"]
							fmt.Println("Notification : ",notif) 
		}

	func continuee(){
			output,err := debug.Send("exec-continue")
			if err != nil {
					fmt.Println(err)		
						}
					
				notif := output["class"]
				fmt.Println("Notification : ",notif) 
		}
	func continue_reverse(){
			output,err := debug.Send("exec-continue","--reverse")

								if err != nil {
								fmt.Println(err)		
									}
					
								notif := output["class"]
								fmt.Println("Notification : ",notif) 
		}
	func backtrace() string {
			
				
								output,_:=debug.Send("stack-list-frames")
								pay:=output["payload"]
								
								payAssert:=pay.(map[string]interface{})
				
								stack:=payAssert["stack"]
							
								stackAssert:=stack.([]interface{})
								nbreFct:=len(stackAssert)
			
								str :=""
								str_variables := ""

								for i:=0; i<=nbreFct-1 ; i++{
									stackSepare:=stackAssert[i]
									stackSepareAssert:=stackSepare.(map[string]interface{})
								
									frame:=stackSepareAssert["frame"]
									frameAssert:=frame.(map[string]interface{})
									
									index := strconv.Itoa(i)
									
									//list variables by frame 
									output_variables,_  := debug.Send("stack-list-variables","--thread","1","--frame",index,"--simple-values")
									map_variables := output_variables["payload"]
									m_variables := map_variables.(map[string]interface{})
									
									variables := m_variables["variables"]
									
									
									variables_slice := variables.([]interface{})
									str_variables = ""
								
								for j:=0; j<len(variables_slice); j++ {
										Separe_variables := variables_slice[j]
										Separe_variablesAssert :=Separe_variables.(map[string]interface{})
										
										name := Separe_variablesAssert["name"]
										str_name := name.(string)
										
										typee := Separe_variablesAssert["type"]
										str_type := typee.(string)
						
										value := Separe_variablesAssert["value"]
										str_value := value.(string)
										
										if arg, ok := Separe_variablesAssert["arg"] ;ok{
										str_arg := arg.(string)
											str_variables +="variable name : " + str_name + " type :" + str_type + " value :" + str_value +
																" arg : " + str_arg + "\n"
										}	else {
											str_variables +="variable name : " + str_name + " type :" + str_type +" value :" + str_value+ "\n"
											}
										
									}
								
									fun:=frameAssert["func"]
									line:=frameAssert["line"]
									level:=frameAssert["level"]
									
									function := fun.(string)
									ligne := line.(string)
									niveau := level.(string)
								
									str += "Frame " +index + "\n"+"level : " + niveau + " function : "+  function + "  line : " + ligne +"\n" + 														str_variables + "\n"
									
									
								}
			return str
		}

	func watch(){

				var input_watch string
				
				fmt.Println("Rentrez la variable suivi")
				fmt.Scanln(&input_watch)
				fmt.Println(debug.Send("break-watch", input_watch))
		}

	func where (){

			output,_ := debug.Send("stack-list-frames")
							pay:=output["payload"]

							payAssert:=pay.(map[string]interface{})
				
							stack:=payAssert["stack"]
							stackAssert:=stack.([]interface{})
							
							//Premier stack			
							stackSepare:=stackAssert[0]
							stackSepareAssert:=stackSepare.(map[string]interface{})
								
							frame:=stackSepareAssert["frame"]
							frameAssert:=frame.(map[string]interface{})
							fun:=frameAssert["func"]
							line:=frameAssert["line"]
							fmt.Println("function : ",fun ,"  line : ",line)
		}

	func breake(){

			var input_break string
			fmt.Println("rentrez votre breakpoint")
			fmt.Scanln(&input_break)
			debug.Send("break-insert", input_break)
		}

	func print(){

			var var_gdb string
			var var_cible string
		
			fmt.Println("Entrez le nom de la variable")
			fmt.Scanln(&var_gdb)
		
			fmt.Println("rentrez la variable cible")
			fmt.Scanln(&var_cible)
		
			debug.Send("var-create", var_gdb, "@", var_cible)
			output,err := debug.Send("var-evaluate-expression", var_gdb)	
				if err !=nil{
								fmt.Println(err)
						}

			fmt.Println(output["payload"])
		}
	func list_variables (){
			
			expr,err := debug.Send("stack-list-variables", "--all-values")	
									if err !=nil {
										fmt.Println(err)
										}

								variables := expr["payload"]
				
								variablesAssert := variables.(map[string]interface {})
								Array := variablesAssert["variables"]
								ArrayAssert := Array.([]interface{})
							
							nbreVar:=len(ArrayAssert)
							for i:=0; i<=nbreVar-1 ; i++{
							mapListe := ArrayAssert[i]
							mapListeAssert := mapListe.(map[string]interface {})
							name := mapListeAssert["name"]
							value := mapListeAssert["value"]
							arg := mapListeAssert["arg"]
							fmt.Println("name : ", name , "value : ",value , "arg : ",arg)
								
							}
				

		}

	func stop(){
		 debug.Interrupt()
		 debug.Exit()
}

