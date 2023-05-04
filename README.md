# api-cook-master
ApiCookMaster
ApiCookMaster est une API REST permettant de gérer une liste d'utilisateurs. Elle

Fonctionnalités
Affichage de la liste des utilisateurs
Affichage d'un utilisateur spécifique
Ajout d'un nouvel utilisateur
Modification d'un utilisateur existant
Suppression d'un utilisateur existant
Utilisation
Prérequis
Go version 1.16 ou supérieure
Installation
Clonez le dépôt GitHub :git clone https://github.com/enzoforreal/ApiCookMaster.git
Ouvrez un terminal dans le dossier du projet :cd ApiCookMaster
Exécutez la commande go mod downloadpour télécharger les dépendances du projet.
Démarrage
Exécutez la commande go run src/main.gopour démarrer le serveur.
accéder à l'API à l'adresse http://localhost:8000.
Points finaux
L'API contient les endpoints suivants :

GET /utilisateurs
Retourne la liste des utilisateurs.

OBTENIR /utilisateurs/{id}
Retourne les informations d'un utilisateur spécifique identifié par l'ID.

POST /utilisateurs
Ajout d'un nouvel utilisateur. Les données JSON doivent être envoyées dans le corps de la requête. Exemple :

json

Copier le code
{
"name": "John Doe"
}
PUT /users/{id}
Met à jour les informations d'un utilisateur identifié existant par l'ID. Les données JSON doivent être envoyées dans le corps de la requête. Exemple :

json

Copier le code
{
"name": "Jane Doe"
}
SUPPRIMER /users/{id}
Supprimer un utilisateur existant identifié par l'ID.