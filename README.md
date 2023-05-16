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

**api/v1/comment/newComment**

La fonction « New Comment » est un POST qui permet de créer un commentaire sur un post. Elle récupère les identifiants du Post et de l’utilisateur ainsi que le Texte par un Json. Elle récupère ensuite la date et heure du système (Aspect a modifié pour mondialisation du programme) puis importe le nouveau commentaire dans la base de données.
