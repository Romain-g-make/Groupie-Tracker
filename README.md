# Presentation

Groupie Tracker is a web application that displays artists, their information, their concert dates, and their locations using data provided by the **Groupie Trackers API**.

The application offers a **simple** interface where the homepage displays all artists, and a search bar allows users to **filter results** based on several criteria such as name, year, date, or location. 
Each artist has their own detailed page gathering all their information as well as their concerts, and navigation is made easier with a **pagination system and alphabetical sorting**.

## Installation

To install the application, you simply need to have Go installed on your machine, clone the project, download the dependencies using the `go mod download` command, and then run the application with `go run main.go`. 
Once this is done, open a browser and go to http://localhost:8080 to access the interface.

## Project Structure

The project is organized around a **main.go** file, which launches the application (entry point).  
A **/Fonction_go/** directory contains all the business logic such as search, sorting, pagination functions, API calls, and models.  
A **/static/** directory stores the HTML and CSS files of the project.


# Présentation

Groupie Tracker est une application web qui permet d’afficher des artistes, leurs informations, leurs dates de concerts ainsi que leurs localisations, en utilisant les données fournies par **l’API Groupie Trackers**. 

L’application propose une interface **simple** où la page d’accueil présente l’ensemble des artistes, et où une barre de recherche permet de **filtrer les résultats** selon plusieurs critères comme le nom, l’année, la date ou la localisation. 
Chaque artiste possède sa propre page détaillée regroupant toutes ses informations ainsi que ses concerts, et la navigation est facilitée par un **système de pagination et un tri alphabétique**.

## Installation

Pour installer l’application, il suffit d’avoir Go installé sur la machine, de cloner le projet, de télécharger les dépendances avec la commande go mod download, puis de lancer l’application avec go run main.go. Une fois cela fait, il suffit d’ouvrir un navigateur et de se rendre à l’adresse http://localhost:8080 pour accéder à l’interface.

## Structure du projet

Le projet est organisé autour d’un fichier **main.go** qui permet de lancer le projet (point d'entrée). 
Un dossier **/Fonction_go/** qui contient toute la logique métier comme les fonctions de recherche, de tri, de pagination, les appels à l’API et les modèles.
Un dossier **/static/** qui stocke les fichiers HTML et CSS du projet.