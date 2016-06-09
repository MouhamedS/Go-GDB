				package main

				import (
					  "fmt"
					  "github.com/cyrus-and/gdb"
					  "io"
					  "os"
					  //"path/filepath"
				)
					
				func main() {


					  // start a new instance and pipe the target output to stdout
					  gdb, _ := gdb.New(nil)
					  go io.Copy(os.Stdout, gdb	)

			

					  // load and run a program
					  gdb.Send("file-exec-and-symbols", os.Args[1])
					 
					  var input string
					  
				 for input != "quit"   {
					  fmt.Println("Que voulez-vous faire ?")
					  fmt.Scanln(&input)
					
						switch input {
						
						//Break
						case "break": breake(gdb)
						

						//Break List						
						case "break-list" : breaklist(gdb)

						
						//Break delete						
						case "delete" : 
{
							var numero_break string		
							fmt.Println("Supprimer un breakpoint(n°) ou tous les breakpoints")
							fmt.Scanln(&numero_break )
							if numero_break != "" {
							gdb.Send("break-delete", numero_break )
						} else {
							gdb.Send("break-delete")
							}			
						}
		
						//Run	
						case "run" : run(gdb)
					
						//Step 
						case "step" : step(gdb)

						//Reverse Stepping
						case "step-reverse" :step_reverse(gdb)
						
						// Continue
						case "continue" :  continuee(gdb)

						//Reverse continue 
						case "continue-reverse" : continue_reverse(gdb)


						//Print
						case "print" : print(gdb)

						//List variables locals
						case "list-variables" : list_variables (gdb)

						
						//Backtrace

						case "backtrace" :backtrace(gdb)

						//Watchpoints
						case "watch" : watch(gdb)

						//Where
						case "where" : where (gdb)
						
						//quit
						case "quit":
						// Default Case					
						default: fmt.Println("Commandes non valides")  

					}
					  
				}	

					  gdb.Exit()
	}

	func run (gdb *gdb.Gdb){
					
				fmt.Println(gdb.Send("exec-run"))
				gdb.Send("interpreter-exec","console","record")	

			}
	func step(gdb *gdb.Gdb){
		
		 output, err := gdb.Send("exec-step")
			if err != nil {
						fmt.Println(err)		
						}
					
							notif := output["class"]
							fmt.Println("Notification : ", notif) 
		
	}
	func breaklist(gdb *gdb.Gdb){

			output,err :=gdb.Send("break-list")
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
	/*func delete_break(gdb *gdb.Gdb){
				
						var numero_break string		
				
						fmt.Println("Supprimer un breakpoint(n°) ou tous les breakpoints")
						fmt.Scanln(&numero_break )
						
						if numero_break != "" {
							gdb.Send("break-delete","numero_break" )
						} else {
							gdb.Send("break-delete")
							}	

	}*/

	func step_reverse(gdb *gdb.Gdb){
		output,err := gdb.Send("exec-step","--reverse")
								if err != nil {
									fmt.Println(err)		
									}
					
							notif := output["class"]
							fmt.Println("Notification : ",notif) 
		}

	func continuee(gdb *gdb.Gdb){
			output,err := gdb.Send("exec-continue")
			if err != nil {
					fmt.Println(err)		
						}
					
				notif := output["class"]
				fmt.Println("Notification : ",notif) 
		}
	func continue_reverse(gdb *gdb.Gdb){
			output,err := gdb.Send("exec-continue","--reverse")

								if err != nil {
								fmt.Println(err)		
									}
					
								notif := output["class"]
								fmt.Println("Notification : ",notif) 
		}
	func backtrace(gdb *gdb.Gdb){
			
				
								output,_:=gdb.Send("stack-list-frames")
								pay:=output["payload"]

								payAssert:=pay.(map[string]interface{})
				
								stack:=payAssert["stack"]
							
								stackAssert:=stack.([]interface{})
								nbreFct:=len(stackAssert)
								for i:=0; i<=nbreFct-1 ; i++{
									stackSepare:=stackAssert[i]
									stackSepareAssert:=stackSepare.(map[string]interface{})
								
									frame:=stackSepareAssert["frame"]
									frameAssert:=frame.(map[string]interface{})
									fun:=frameAssert["func"]
									line:=frameAssert["line"]
									level:=frameAssert["level"]
									fmt.Println("level : ",level,"function : ",fun ,"  line : ",line)

								
								}
		}

	func watch(gdb *gdb.Gdb){

				var input_watch string
				
				fmt.Println("Rentrez la variable suivi")
				fmt.Scanln(&input_watch)
				fmt.Println(gdb.Send("break-watch", input_watch))
		}

	func where (gdb *gdb.Gdb){

			output,_ := gdb.Send("stack-list-frames")
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

	func breake(gdb *gdb.Gdb){

			var input_break string
			fmt.Println("rentrez votre breakpoint")
			fmt.Scanln(&input_break)
			gdb.Send("break-insert", input_break)
		}

	func print(gdb *gdb.Gdb){

			var var_gdb string
			var var_cible string
		
			fmt.Println("Entrez le nom de la variable")
			fmt.Scanln(&var_gdb)
		
			fmt.Println("rentrez la variable cible")
			fmt.Scanln(&var_cible)
		
			gdb.Send("var-create", var_gdb, "@", var_cible)
			output,err := gdb.Send("var-evaluate-expression", var_gdb)	
				if err !=nil{
								fmt.Println(err)
						}

			fmt.Println(output["payload"])
		}
	func list_variables (gdb *gdb.Gdb){
			
			expr,err := gdb.Send("stack-list-variables", "--all-values")	
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

