# AccessGuard

**Context-Aware Access Control & Secure File Management Platform**

AccessGuard est une application web de cybersécurité qui contrôle non seulement **qui** peut accéder à l'application, mais aussi **ce qu'il peut faire** et **dans quelles conditions**.

> 🚧 Projet en cours de développement (projet d'étude en cybersécurité).

## Fonctionnalités

- **Authentification contextuelle** : à la connexion, le backend évalue l'IP, le pays, l'appareil, l'heure et l'historique de connexion. Un moteur de règles décide : `ALLOW`, `BLOCK` ou `REQUIRE_MFA`.
- **Règles dynamiques** : configurables depuis le dashboard administrateur, jamais codées en dur.
- **RBAC** : rôles personnalisables, permissions attribuables dynamiquement, rôles assignables aux utilisateurs. Les permissions sont vérifiées par le backend Go à chaque opération sensible.
- **Gestion sécurisée des fichiers** : upload, liste, visualisation, téléchargement et suppression, avec propriété et isolation des fichiers imposées côté serveur.
- **Dashboard administrateur** : utilisateurs, rôles, permissions, règles contextuelles, audit logs, historique des connexions, événements de sécurité.
- **Audit logging** : journal en ajout seul des événements importants (connexions, MFA, fichiers, changements de rôles, permissions et règles).
- **HashiCorp Vault** : gestion des secrets (identifiants base de données, secrets applicatifs, configuration sensible).

## Principes de sécurité

Quatre couches distinctes :

| Couche | Question | Moment |
|---|---|---|
| Authentification | Qui est l'utilisateur ? | Login |
| Authentification contextuelle | Le contexte de connexion est-il acceptable ? | Login |
| Autorisation (RBAC) | Que peut faire cet utilisateur ? | Chaque requête protégée |
| Autorisation au niveau fichier | Peut-il accéder à ce fichier précis ? | Chaque opération sur fichier |

Le frontend n'est jamais considéré comme fiable : le backend est la seule source de vérité pour toutes les décisions d'autorisation.

## Stack technique

| Couche | Technologies |
|---|---|
| Frontend | React, TypeScript, Vite, CSS |
| Backend | Go, API REST, JWT |
| Base de données | MongoDB |
| Secrets | HashiCorp Vault |
| Infrastructure | Docker, Docker Compose, Nginx / Traefik |

## Architecture

```
Navigateur (React) → Reverse proxy → API Go → MongoDB
                                        ↓
                                  HashiCorp Vault
```

## Modèle de données

Collections MongoDB : `users`, `roles`, `permissions`, `files`, `rules`, `audit_logs`, `login_history`.

- Un utilisateur a un ou plusieurs rôles (`users.role_ids`).
- Un rôle contient plusieurs permissions (`roles.permissions`).