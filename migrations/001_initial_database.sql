CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(50) NOT NULL,
    prenom VARCHAR(50) NOT NULL,
    adresse VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telephone VARCHAR(20) NOT NULL,
    mot_de_passe VARCHAR(100) NOT NULL,
    photo_de_profil VARCHAR(100),
    est_admin BOOLEAN NOT NULL CHECK (est_admin IN (true, false)),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE planification (
    id SERIAL PRIMARY KEY,
    id_evenement INT,
    date_debut TIMESTAMPTZ NOT NULL,
    date_fin DATE NOT NULL,
    lieu VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE abonnements (
    id SERIAL PRIMARY KEY,
    id_user INT,
    date_debut TIMESTAMPTZ NOT NULL,
    date_fin DATE NOT NULL,
    type_abonnement VARCHAR(20) NOT NULL,
    nb_utilisations INT,
    nb_reservations INT,
    statut_abonnement VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE fidelite (
    id SERIAL PRIMARY KEY,
    id_user INT,
    points_fidelite INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE ventes (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    id_produit INT NOT NULL,
    date_vente TIMESTAMPTZ NOT NULL,
    quantite INT DEFAULT 1,
    prix_total FLOAT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE devis (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    date_emission TIMESTAMPTZ NOT NULL,
    montant_total FLOAT NOT NULL,
    produits_services VARCHAR(255),
    statut_devis VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE factures (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    date_emission TIMESTAMPTZ NOT NULL,
    montant_total FLOAT NOT NULL,
    produits_services VARCHAR(20),
    statut_facture VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE evenements (
    id SERIAL PRIMARY KEY,
    date_evenement TIMESTAMPTZ NOT NULL,
    type_evenement VARCHAR(20) NOT NULL,
    salle_lieu VARCHAR(20) NOT NULL,
    intervenants VARCHAR(20),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE reservations (
    id SERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    id_evenement INT NOT NULL,
    id_cours INT,
    date_reservation TIMESTAMPTZ NOT NULL,
    nb_personnes INT NOT NULL,
    statut_reservation VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE cours (
    id SERIAL PRIMARY KEY,
    titre VARCHAR(30) NOT NULL,
    description VARCHAR(255),
    date_debut TIMESTAMP NOT NULL,
    date_fin TIMESTAMP NOT NULL,
    lien_cours VARCHAR(255),
    prix FLOAT NOT NULL,
    id_prestataire INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE catalogue (
    id SERIAL PRIMARY KEY,
    nom_service VARCHAR(30) NOT NULL,
    description VARCHAR(100),
    tarif FLOAT NOT NULL,
    id_prestataires INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE prestataires (
    id SERIAL PRIMARY KEY,
    nom_prestataire VARCHAR(30) NOT NULL,
    adresse VARCHAR(100),
    telephone VARCHAR(10),
    email VARCHAR(255) NOT NULL,
    description VARCHAR(100),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);


ALTER TABLE planification ADD FOREIGN KEY (id_evenement) REFERENCES evenements (id);

ALTER TABLE abonnements ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE fidelite ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE ventes ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE ventes ADD FOREIGN KEY (id_produit) REFERENCES catalogue (id);

ALTER TABLE devis ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE factures ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE reservations ADD FOREIGN KEY (id_user) REFERENCES users (id);

ALTER TABLE reservations ADD FOREIGN KEY (id_evenement) REFERENCES evenements (id);

ALTER TABLE reservations ADD FOREIGN KEY (id_cours) REFERENCES cours (id);

ALTER TABLE cours ADD FOREIGN KEY (id_prestataire) REFERENCES prestataires (id);

ALTER TABLE catalogue ADD FOREIGN KEY (id_prestataires) REFERENCES prestataires (id);
