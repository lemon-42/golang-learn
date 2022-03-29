// Golang glossaire made by lemon_42 for Zone01

package main // Ou piscine dans notre cas

import "github.com/01-edu/z01" // Importation de la bibliothèque pour la piscine

func main() { // Fontion principale

	z01.PrintRune("Hello World !") // Affiche Hello World !

}



/* Définir une variable (syntaxe) :

var [nom_de_la_variable] [type]

*/

var nom_de_la_variable int 
var nom_de_la_variable int = 12 // Surcharge de la variable
var var1, var2, var3 = 1, 2, 3 // Surchargement des variables
ma_variable := 42 // Ecriture raccourcis pour tout types
float := 12.3
string := "Citron"

// Il est également possible d'utiliser une autre syntaxe pour un affichage plus clair

func main() {
	var (
		xp int = 0
		nom string = "Citron"
		vitesse float32 = 4.2
	)
}

/* Typage dynamique

(Variable dynamique = pas besoin de spécifier le type lors de la déclaration de la variable, l'ordinateur s'en occupe)

Pour déclarer une variable dynamique on supprime le mot clé var et on le remplace par := */

float := 15.5 // sera automatiquement de type float
in := 5 //  sera automatiquement de type int
str := "Citron" //  sera automatiquement de type string
bol := true, false //  sera automatiquement de type boolean

// Les constantes : possède une valeur fixe qui ne peut être modifié.

const maConstante int = 50

fmt.Println("ma Constante : ", maConstante)

// Les calculs 

/*

Addition = +
Soustraction = -
Multiplication = *
Division = /
Modulo = % (reste de la division, exemple le modulo de 10 par 3 est 1)

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

var a int = 4
var b int = 2

a += b
fmt.Println("a += b  = ", a)

a -= b
fmt.Println("a -= b  = ", a)

a *= b
fmt.Println("a *= b  = ", a)

a /= b
fmt.Println("a /= b  = ", a)

a %= 3
fmt.Println("a %= b  = ", a)

// Opérateurs d'incrémentation

++ // Incrémente d'une unité la variable
-- // Décrémente d'une unité la variable

// Opérateur relationel

== : Est strictement égal à
!= : Est différent de
> : Plus grand que
< : Plus petit que
>= : Plus grand ou égal
<= : Plus petit ou égal

// Opérateur logique

&& : Et
|| : Ou
! : Pas

// Boucle for : 

for i := 1; i < 3; i++ {
	z01.PrintRune("Citron")
}

// Pour i est égal à 1, tant que i est inférieur à 3, on incrémente i de 1

//Résultat :

Citron
Citron
Citron

// Déclinaison boucle for

for xp < 1000000 {
	z01.PrintRune("Ne lâche rien")
}


// Conditions If Else :

age := 20

if age < 18 {
	z01.PrintRune("Vous êtes mineur")
} else {
	z01.PrintRune("Vous êtes majeur")
}

// Résultat : 

Vous êtes majeur 

// Variante de If Else

age := 20

if age > 10 && age < 18 {
	z01.PrintRune("Tu es jeune")
} else if age > 18 && age < 30 {
	z01.PrintRune("Tu commence à te faire vieux")
}

// Résultat : 

Tu commence à te faire vieux

// Break && Continue 

/*
Break : Interompt la boucle
Continue : Reviens au début de la boucle
*/

// Fonctions 

// Délarer une fonction 

func maFonction (liste_de_paramêtre) type_de_retour {
	/* Votre code */
}

// Exemple : 

func Addition(a int, b int) int {
	
	resultat := a + b
	return resultat
}

// Fonction avec paramètre et sans valeurs de retour

func affichage(nom string, age int) {
    fmt.Println("Bonjour", nom, "vous avez", age, "ans")
}

func main() {
    affichage("Einstein", 76)
    affichage("Golang", 10)
}

// Fonction avec valeur de retour : 

func Addition(a int, b int) int {
	
	resultat := a + b
	return resultat
}

