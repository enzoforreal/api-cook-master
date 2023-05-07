# Utilise une image de Go spécifique comme image de base pour la compilation
FROM golang:1.17-alpine3.13 AS build

# Répertoire de travail dans le conteneur
WORKDIR /app

# Copie les fichiers du projet dans le conteneur
COPY . .

# Compile l'application
RUN cd src && go build -o /app/main

# Utilise une image Alpine Linux comme image de base pour l'exécution
FROM alpine:3.14

# Installe PostgreSQL
RUN apk update && \
    apk add postgresql postgresql-client && \
    mkdir /run/postgresql && \
    chown postgres /run/postgresql && \
    su -c "initdb -D /var/lib/postgresql/data" postgres && \
    echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf && \
    echo "listen_addresses='*'" >> /var/lib/postgresql/data/postgresql.conf

# Copie le script de migration dans le conteneur
COPY ./migrations/001_initial_database.sql /docker-entrypoint-initdb.d/

# Ajoute le fichier binaire de l'application dans le conteneur
COPY --from=build /app/main /usr/local/bin/api

# Copie le fichier .env dans le conteneur
COPY .env /

# Défini le répertoire de travail par défaut
WORKDIR /usr/local/bin

# Expose le port de l'API
EXPOSE 8000

# Défini les variables d'environnement pour l'API
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=cookmaster

# Lance l'application
CMD ["api"]
