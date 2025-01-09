# Projet Horloge Temps Réel en Go

Ce projet est une application simple écrite en Go qui affiche en continu l'heure actuelle au format `HH:MM:SS` dans la console. L'heure est mise à jour chaque seconde.

## Fonctionnalités

- Affiche l'heure actuelle en temps réel.
- Met à jour l'affichage toutes les secondes.
- Code simple et facile à comprendre.

## Prérequis

- [Go](https://golang.org/) version 1.16 ou ultérieure installée sur votre système.

## Installation

1. Clonez ce dépôt sur votre machine locale

2. Assurez-vous que Go est correctement installé sur votre système en utilisant la commande suivante :

```bash
go version
```

3. Compilez et exécutez le programme en utilisant la commande suivante :

```bash
go build current_time.go
./current_time
```

## Utilisation

L'heure actuelle sera affichée en temps réel dans le terminal, mise à jour toutes les secondes.

Exemple de sortie :

```
Heure actuelle : 12:34:56
```

Pour arrêter le programme, utilisez `Ctrl+C` dans le terminal.

# Démarrer ce projet avec docker

```
docker build . -t current_time_go
docker run --rm current_time_go
```

Pour arrêter le programme, utilisez `Ctrl+C` dans le terminal.
