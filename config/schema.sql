-- Active: 1678732327886@@127.0.0.1@5432@commentaires@public
CREATE TABLE table_name(  
    ID_Commentaires int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    Texte VARCHAR(500),
    Date_Commentaire DATE,
    Pouce_Rouge INT,
    Pouce_Vert INT,
    ID_Post UUID
);
