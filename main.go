package main // OBLIGATOIRE dans tout programme en go.

import "fmt" // Importation de la bibliothèque fmt

func main() { // Fontion principale

	fmt.Println("Hello World !") // Affiche Hello World !

}



/* Définir une variable :

var [nom_de_la_variable] [type]

*/

var vie int

var vie int = 12

var vie, argent, puissance int

fmt.Println("vie : ", vie)
fmt.Println("argent : ", argent)
fmt.Println("puissance : ", puissance)

var vie, argent, puissance int = 10, 20, 30 // Surcharge des variables

// Il est possible de déclarer plusieurs variables avec des types différents

var vie int = 20
var nom string = "Juju"
var vitesse float32 = 5.4

// Il est également possible d'utiliser une autre syntaxe pour un affichage plus clair

func main() {
	var (
		vie	int = 20
		nom string = "Default"
		vitesse float32 = 5.4
	)
}

/* Typage dynamique

(Variable dynamique = pas besoin de spécifier le type lors de la déclaration de la variable, l'ordinateur s'en occupe)

Pour déclarer une variable dynamique on supprime le mot clé var et on le remplace par := */

flt := 15.5 // sera automatiquement de type float
in := 5 //  sera automatiquement de type int
st := "hello" //  sera automatiquement de type string
bol := true //  sera automatiquement de type boolean

fmt.Printf("Le type de la varialbe flt est %T\n", flt)
fmt.Printf("Le type de la varialbe in est %T\n", in)
fmt.Printf("Le type de la varialbe st est %T\n", st)
fmt.Printf("Le type de la varialbe bol est %T\n", bol)

/*
    %T : affiche le type d'une valeur
    %d : affiche un entier
    %s : affiche une chaîne de caractères
    %f : affiche un nombre décimal
    %b : affiche une représentation binaire
*/

// Les constantes : possède une valeur fixe qui ne peut être modifié.

const maConstante int = 50

fmt.Println("ma Constante : ", maConstante)

// Les calculs 

/*

Addition = +
Soustraction = -
Multiplication = *
Division = /
Modulo = %

*/

var a int = 4
var b int = 2

fmt.Println("a + b = ", a+b)
fmt.Println("a + b = ", a-b)
fmt.Println("a + b = ", a*b)
fmt.Println("a + b = ", a/b)
fmt.Println("a + b = ", a%b)

// Opérateurs d'assignation 

/*

+= 
-=
*=
/=
%=

*/