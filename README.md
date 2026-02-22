# Word Stats

## Description

Word Stats est un projet en Go qui analyse une chaîne de texte pour compter :

- le nombre de lignes,
- le nombre de mots,
- le nombre de caractères (hors espaces).

Ce projet sert aussi à illustrer un workflow Git collaboratif : création de branches, fusions, résolution de conflits, et historique des commits.

---

## Fonctions principales

- `countLines(text string) int` : compte le nombre de lignes (`\n`) dans le texte.  
- `countWords(text string) int` : compte le nombre de mots séparés par des espaces.  
- `countChars(text string) int` : compte les caractères en ignorant les espaces, sauts de ligne et tabulations.  

---

## Structure du projet
word-stats/
│── main.go
│── main_test.go
│── README.md
│── history.txt


---

## Exécution

Pour lancer le programme :
```bash
go run main.go
```
Pour exécuter les tests unitaires avec couverture :
```bash
go test -cover
```

## Workflow Git

- **Initialisation** : `git init` et premier commit de la fonction `countLines`.
- **Branches** :
    - `count-words` : Développement de la fonction de comptage de mots.
    - `count-chars` : Développement de la fonction de comptage de caractères.
- **Fusion sans conflit** : Intégration de la branche `count-words` dans `main`.
- **Fusion avec conflit** :
    * Modifications concurrentes dans `main()`.
    * Tentative de fusion (`git merge count-chars`).
    * Résolution manuelle du conflit, test, puis commit final.
- **Historique** : Généré via `git log --oneline > history.txt`.

## Tests et couverture

- **Tests unitaires** : Chaque fonction possède son propre test dans le fichier `main_test.go`.
- **Taux de couverture** : L'objectif était d'avoir une couverture de code ≥ 80 %.
- **Test d'intégration** : La fonction `main()` est aussi testée pour garantir que l'exécution complète du programme fonctionne sans erreur.

## Objectifs atteints

- **Maîtrise des commandes Git** : Utilisation des commandes de base au quotidien.
- **Gestion des branches** : Création des fonctionnalités via des branches dédiées.
- **Résolution de conflits** : Capacité à corriger manuellement les conflits lors des fusions.
- **Développement en Go** : Écriture de fonctions simples, robustes et faciles à tester.
- **Qualité de code** : Analyse et optimisation de la couverture de tests.