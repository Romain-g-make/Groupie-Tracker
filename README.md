Overview

Groupie Tracker is a web application that displays artists, their information, concert dates, and locations, using data provided by the Groupie Trackers API. The application features a simple interface where the homepage displays all the artists, and a search bar allows you to filter results according to several criteria such as name, year, date, or location. Each artist has their own detailed page containing all their information and concerts, and navigation is facilitated by pagination and alphabetical sorting.

Installation

To install the application, simply have Go installed on your machine, clone the project, download the dependencies with the command `go mod download`, and then launch the application with `go run main.go`. Once this is done, simply open a browser and go to http://localhost:8080 to access the interface. 

Project Structure

The project is organized around a main.go file which serves as the entry point, and a Function_go/ folder which contains all the business logic such as search, sorting, pagination, API calls, and models.


Présentation

Groupie Tracker est une application web qui permet d’afficher des artistes, leurs informations, leurs dates de concerts ainsi que leurs localisations, en utilisant les données fournies par l’API Groupie Trackers. L’application propose une interface simple où la page d’accueil présente l’ensemble des artistes, et où une barre de recherche permet de filtrer les résultats selon plusieurs critères comme le nom, l’année, la date ou la localisation. Chaque artiste possède sa propre page détaillée regroupant toutes ses informations ainsi que ses concerts, et la navigation est facilitée par un système de pagination et un tri alphabétique.

Installation

Pour installer l’application, il suffit d’avoir Go installé sur la machine, de cloner le projet, de télécharger les dépendances avec la commande go mod download, puis de lancer l’application avec go run main.go. Une fois cela fait, il suffit d’ouvrir un navigateur et de se rendre à l’adresse http://localhost:8080 pour accéder à l’interface.

Structure du projet

Le projet est organisé autour d’un fichier main.go qui sert de point d’entrée, d’un dossier Fonction_go/ qui contient toute la logique métier comme la recherche, le tri, la pagination, les appels à l’API et les modèles
