-- Active: 1678732327886@@127.0.0.1@5432@commentaires@public
CREATE TABLE commentaires(  
    Id_Commentaire int NOT NULL UNIQUE PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    Id_User VARCHAR(25),
    Texte VARCHAR(500),
    Date_Commentaire TIMESTAMP,
    Id_Post UUID
);

CREATE TABLE avis(
    Id_Commentaire int REFERENCES commentaires(id_commentaire),
    Id_User int,
    Valeur  BOOLEAN,
    PRIMARY KEY(Id_Commentaire, Id_User)
);
