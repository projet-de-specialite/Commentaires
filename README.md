# API Commentaires

L’api commentaire permet de gérer des commentaires liés à un post, ainsi que les avis de ces commentaires (Pouce rouge/vert).

Elle est programmée en Goland. Elle nécessite l’environnement de celui-ci (https://go.dev/dl/).
Elle utilise aussi une base de données Postgres.

## Les Structures :
Commentaires :
- Id_Commentaire      INT
- Id_User             INT
- Texte               VARCHAR
- Date_Commentaire    DATE
- Id_Post             UUID

Les commentaires comprennent un identifiant, l’identifiant de l’utilisateur ayant posté le commentaire, le Texte du commentaire, la Date avec l’année, le mois, le jour, l’heure, les minutes et les secondes à laquelle le commentaire a été posté. Et pour finir l’identifiant du poste sur lequel le commentaire a été fait.


Avis : 
- Id_Commentaire  INT
- Id_User INT
- Valeur  BOOL

Les avis sont identifiés par l’identifiant du commentaire sur lequel l’avis est donné, ainsi que l’identifiant de l’utilisateur qui a donné son avis. Ils contiennent aussi une valeur booléenne qui correspond a pouce vert si la valeur est vraie ou pouce rouge si la valeur est fausse.

## Les Fonctionnalité : 

Nouveau Commentaire : **api/v1/comment/newComment**

La fonction « New Comment » est un POST qui permet de créer un commentaire sur un post. Elle récupère les identifiants du Post et de l’utilisateur ainsi que le Texte par un Json. Elle récupère ensuite la date et heure du système (Aspect a modifié pour mondialisation du programme) puis importe le nouveau commentaire dans la base de données.

Type : POST

Entrée : json contenant un Id_User, un Texte et un Id_Post 
{
    "Id_User": 2,
    "Texte": "Trop bien!",
    "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
}

Sortie : json contenant le mot "succes" qui contient un commentaire 
{
    "succes": {
        "Id_Commentaire": 0,
        "Id_User": 2,
        "Texte": "Trop bien!",
        "Date_Commentaire": "2023-05-18T00:04:28.987164+02:00",
        "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
    }
}
*******
Nouvelle Avis : **api/v1/comment/newAvis**

La fonction « New Avis » est un POST qui permet à un utilisateur de donner son avis sur un commentaire représenté par un pouce rouge pour un désaccord ou un pouce vert pour un accord. La valeur de cet avis est un booléen qui est vrai pour un pouce vert et faux pour un pouce rouge. La fonction récupère l’identifiant du commentaire et de l’utilisateur ainsi que la valeur par un Json afin de le stocker en base de données.

Type : POST

Entrée : json contenant un Id_Commentaire, un Id_User et une Valeur 
{
    "Id_Commentaire": 3,
    "Id_User": 1,
    "Valeur": true
}


Sortie : json contenant le mot "succes" qui contient un avis
{
    "succes": {
        "Id_Commentaire": 3,
        "Id_User": 2,
        "Valeur": true
    }
}
*******
Retrouver un commentaire : **api/v1/comment/findCommentById/:Id_Commentaire**

La fonction   « Find Comment By Id » est un GET qui permet de retrouver un commentaire par son identifiant donné dans l’URL. Elle renvoie un Json contenant le commentaire trouvé.

Type : GET

Entrée : "Id_Commentaire" en parametre

Sortie : json contenant un commentaire
{
    "Id_Commentaire": 3,
    "Id_User": 2,
    "Texte": "Trop bien!",
    "Date_Commentaire": "2023-05-14T21:55:47Z",
    "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
}
*******
Compter les Avis : **api/v1/comment/nombreAvis/:Id_Commentaire**

La fonction « Nombre d’Avis » est un GET qui permet de compter le nombre de pouce rouge et de pouce vert d’un commentaire donner dans l’URL. Elle renvoie ensuite un Json contenant deux valeurs, la première le nombre de pouce vert, et la deuxième le nombre de pouce rouge.

Type : GET

Entrée : "Id_Commentaire" en parametre

Sortie : json contenant deux valeur [ nombre pouce vert, nombre de pouce rouge ]
[
    0,
    0
]


*******
Afficher les commentaires : **api/v1/comment/printComment/:Id_Post**

La fonction « Print Comment » est un GET qui permet d’afficher tous les commentaires d’un post donner dans l’URL. Elle retourne un Json contentant un tableau de Commentaire.

Type : GET

Entrée : "Id_Post" en parametre

Sortie : json contenant plusieurs commentaires
[
    {
        "Id_Commentaire": 3,
        "Id_User": 2,
        "Texte": "Trop bien!",
        "Date_Commentaire": "2023-05-14T21:55:47Z",
        "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
    },
    {
        "Id_Commentaire": 4,
        "Id_User": 2,
        "Texte": "Trop bien!",
        "Date_Commentaire": "2023-05-18T00:01:27Z",
        "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
    },
    {
        "Id_Commentaire": 5,
        "Id_User": 2,
        "Texte": "Trop bien!",
        "Date_Commentaire": "2023-05-18T00:04:28Z",
        "Id_Post": "158c9998-18da-4bef-a8db-4891b1736574"
    }
]
*******
Supprimer un commentaire : **api/v1/comment/deleteComment/:Id_Commentaire**

La fonction « Delete Comment » est un DELETE qui permet de supprimer un commentaire donner dans l’URL. Elle supprime aussi tous les avis liés à ce commentaire.

Type : DELETE

Entrée : "Id_Commentaire" en parametre*

Sortie : json contenant le mot "succes" puis un string 
{
    "succes": "Row deleted"
}
*******
