# Description :

Interface GDB écrite en langage Go  en utilisant GDB/MI à travers l'interface  Go-GDB : https://github.com/cyrus-and/gdb et la partie graphique écrite en QML à travers le package Go-QML : https://github.com/go-qml/qml
				
## Pré-requis : 
	Installer Go
	Installer Qt
				


## Manuel Utilisateur :

•	Compiler votre programme à debogger en static avec l’option débogage (Exemple: gcc -g -static essai.c -o essai)

•	Il faut que le fichier source et son code compilé soient dans Projet-Go/Ressources , avec les code Go et QML dans Projet-Go
	
•	Le fichier debug_gdb.go dans Ressources correspond à la première version du code Go pour utiliser GDB seulement en mode console 
	
utilisable via la commande : go run debug_gdb.go  nom_fichier_compilé

•	Aller dans le dossier où est le code en Go et compiler le tout avec : go build

•	Lancer cette commande pour éviter une erreur survenue lors de la dernière version de Go avec Go-QML : export GODEBUG=cgocheck=0

•	Lancer le programme avec en indiquant le fichier source du programme à debugger en premier suivi de son code compilé : 
						
							Exemple :  ./Projet-Go essai.c essai
															
								        Projet-Go : programme compilé
														
									essai.c : code source à debugger
														
									essai : code compilé d’essai.c
 

•	Choisir un ou plusieurs breakpoint en cliquant sur des numéros de ligne 

•	Appuyer sur le bouton Play pour commencer le débogage (Run)

•	Appuyer sur le bouton Avance pour faire un débogage pas à pas (Step)

•	 Appuyer sur le bouton Retour pour revenir sur vos pas (Reverse-step)

•	Appuyer sur le bouton Avance rapide pour faire un saut (Continue)

•	Appuyer sur le bouton Retour rapide pour retourner du saut (Reverse-continue)

•	Les couches (frames) et variables par couches sont visibles dans rectangle de gauche






