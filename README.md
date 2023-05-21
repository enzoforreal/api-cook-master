# api-cook-master en construction
Endpoint : /auth/login

Méthode : POST
Description : Permet à un utilisateur de s'authentifier.
Paramètres de requête :
email : Adresse e-mail de l'utilisateur.
password : Mot de passe de l'utilisateur.
Réponse :
Statut 200 OK : Succès de l'authentification.
Statut 401 Unauthorized : Échec de l'authentification.
Corps de la réponse : Utilisateur authentifié au format JSON.
Endpoint : /register

Méthode : POST
Description : Permet à un utilisateur de s'inscrire.
Paramètres de requête :
nom : Nom de l'utilisateur.
prenom : Prénom de l'utilisateur.
adresse : Adresse de l'utilisateur.
email : Adresse e-mail de l'utilisateur.
mot_de_passe : Mot de passe de l'utilisateur.
telephone : Numéro de téléphone de l'utilisateur.
Réponse :
Statut 201 Created : L'utilisateur a été créé avec succès.
Statut 400 Bad Request : Échec de la création de l'utilisateur en raison de données incorrectes ou manquantes.
Corps de la réponse : Utilisateur créé au format JSON.
Endpoint : /users/{id}

Méthode : GET
Description : Récupère les informations d'un utilisateur spécifié par son ID.
Paramètres de requête :
id : ID de l'utilisateur.
Réponse :
Statut 200 OK : Succès de la récupération des informations de l'utilisateur.
Statut 404 Not Found : L'utilisateur avec l'ID spécifié n'a pas été trouvé.
Corps de la réponse : Informations de l'utilisateur au format JSON.
Endpoint : /users/{id}

Méthode : PUT
Description : Met à jour les informations d'un utilisateur spécifié par son ID.
Paramètres de requête :
id : ID de l'utilisateur.
Corps de la requête : Nouvelles informations de l'utilisateur au format JSON.
Réponse :
Statut 200 OK : Succès de la mise à jour des informations de l'utilisateur.
Statut 404 Not Found : L'utilisateur avec l'ID spécifié n'a pas été trouvé.
Corps de la réponse : Utilisateur mis à jour au format JSON.
Endpoint : /users/{id}

Méthode : DELETE
Description : Supprime un utilisateur spécifié par son ID.
Paramètres de requête :
id : ID de l'utilisateur.
Réponse :
Statut 204 No Content : Succès de la suppression de l'utilisateur.
Statut 404 Not Found : L'utilisateur avec l'ID spécifié n'a pas été trouvé.
Ceci est une description générale de l'API Cookmaster en Go
