# Tetris Optimizer

<div align="center">

**Algorithme de résolution pour placer des pièces de tetris dans le plus petit cadre possible.**

</div>

## Aperçu

Tetris Optimizer est un outil de console de commande programmé en Go. L'objectif est de placer les tetrominos fournis de la façon la plus optimale possible.
C'est un projet d'études visant à étudier l'algorithmique pour apprendre à trouver le meilleur placement possible pour des pièces de Tetris dans un cadre carré de la plus petite taille possible.

## Fonctionnalités

- **Lecture des pièces** : Lit le fichier texte fourni au programme pour repérer d'éventuelles erreurs dans la construction des pièces. Si le fichier est valide, les pièces sont transmises à l'algorithme.
- **Tri des pièces** : Trie les pièces et analyse leur forme et leur position.
- **Placement des pièces** : Construit le plus petit carré possible avec les pièces transmises.

## Technologies utilisées

**Language:**

[![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)

## Utilisation

### Prérequis

Go doit être installé sur votre ordinateur.

- **Go**: Version 1.21 ou plus. Téléchargement : [golang.org](https://golang.org/dl/).

### Installation

1.  **Cloner le repo**

    ```bash
    git clone https://github.com/JadeJLC/Tetris-optimizer.git
    cd Tetris-optimizer
    ```

2.  ** Créer un fichier texte**
    Exemple de format attendu pour une pièce :

```
#...
#...
#...
#...
```

3.  **Lancer le fichier**
    Alternatively, to run without building an executable (for development):
    ```bash
    go run ./main . nomdufichier.txt
    ```

### **Sortie attendue**
Le format de sortie de fichier indique chaque pièce par une lettre, de cette façon :
```
ABB.
ABB.
A.C.
ACCC
```
Trois tétrominos (ligne verticale 1, carré B et "T" C) dans un carré de 4x4.

## **Erreurs**

Le programme renvoie une erreur si :

- Une des pièces est composée de plus ou moins que quatre carrés
- Le format de chaque pièce ne correspond pas au format attendu
- Le fichier n'est pas organisé correctement


## Structure du projet

```
project-root/
├── functions/      # Fonctions principale pour la manipulation des pièces
├── main/           # Package principal pour lancer l'application (main.go)
│   └── main.go     # Fichier de lancement de Tetris Optimize
├── go.mod          # Go module 
├── readme.md       # Ce fichier
```

## Development

### Tests

Le programme peut être testé en lançant le fichier `functions_test.go` :

```bash
go run ./functions/functions_test.go
```

## Autres informations

- Ce projet est un projet optionnel développé dans le cadre de ma formation à Zone01.
- Il a été en partie réalisé avec l'aide d'un autre apprenant versé en algorithmie.


<div align="center">

Par [JadeJLC](https://github.com/JadeJLC)

</div>
