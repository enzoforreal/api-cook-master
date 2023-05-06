# Utilisez une image de Go comme image de base pour la compilation
FROM golang:1.20-alpine3.14 AS build

# Répertoire de travail dans le conteneur
WORKDIR /app

# Copiez les fichiers du projet dans le conteneur
COPY . .

# Compilez l'application
RUN go build -o main

# Utilisez une image Alpine Linux comme image de base pour l'exécution
FROM alpine:3.14

# Ajoutez le fichier binaire de l'application dans le conteneur
COPY --from=build /app/main /usr/local/bin/api

# Installez PostgreSQL
RUN apk update && \
    apk add postgresql postgresql-client && \
    mkdir /run/postgresql && \
    chown postgres /run/postgresql && \
    su -c "initdb -D /var/lib/postgresql/data" postgres && \
    echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf && \
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

# Démarrez le service PostgreSQL et créez la base de données pour l'API
RUN su -c "pg_ctl start -D /var/lib/postgresql/data && \
           psql -U postgres -c 'CREATE DATABASE cookmaster;'" postgres

# Définissez le répertoire de travail par défaut
WORKDIR /usr/local/bin

# Exposez le port de l'API
EXPOSE 8080

# Lancez l'application
CMD ["api"]
