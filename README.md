# DC20 Clerk

**DC20 Clerk** is a modular, full-stack web application designed to support gameplay for the tabletop RPG system **DC20**. It serves as both a learning platform for modern web development practices and a portfolio project demonstrating microservices architecture, CI/CD, and scalable design.

---

## Project Overview

DC20 Clerk is a suite of interactive utilities for players and Game Masters, including:

- A character sheet manager with stat calculations and version history
- A spell/ item/ monster library with support for custom creation
- A searchable rules reference
- A combat tracker with true random dice rolling
- A social system for managing encounters and shared sessions

While a typical VTT (Virtual Table Top) is a full replacement of the in-person tabletop experience, DC20 Clerk is intended more as a TTA (Table Top Assistant) for streamlining the in-person experience, or for providing a simpler way to play online without all the complexity of a VTT.

---

## Project Goals

- Learn and apply:
  - Microservices architecture with Go
  - Frontend development with React + Vite + TypeScript
  - Docker containerization
  - Authentication and persistent storage
  - CI/CD pipelines using GitHub Actions and GitOps
  - End-to-end testing with Playwright
  - Mobile-friendly design
- Build a polished, usable toolset for DC20 players
- Showcase advanced development practices in a public portfolio

---

## MVP Scope

The Minimum Viable Product includes:

- [ ] Homepage with navigation
- [ ] User authentication (login/register)
- [ ] Character Sheet builder:
  - [ ] Create and manage characters
  - [ ] Save/export/share character sheets
  - [ ] Form fillable sheet with no calculations
- [ ] Fully deployed and accessible
  - [ ] Frontend deployed to GitHub Pages
  - [ ] Backend deployed to Leapcell
  - [ ] Database and auth hosted by Supabase
  - [ ] Accessible and secure via custom domain

---

## Future Roadmap

- Character sheet
  - Fully calculated with all character options
  - Level up assistant
  - Inventory manager
  - Version history (rollback to previous levels)
- Combat/ Encounter assistant
  - Action point tracker
  - Integration for linking to characters and monsters
  - Fully rules-aware interface for performing actions
  - Damage estimation and calculation
  - Effects apply to character sheet automatically
  - True random dice rolls via random.org
- Spellbook
  - Quick reference to all spells
  - Spell modification interface
- Monster Library
  - Quick reference to all monsters with search and filters
  - Custom monster builder
- Armory
  - List of all weapons and items
  - Custom weapon and item builder
- Rulebook
  - Quick reference with search and filters
- Party system
  - Social features for adding friends and joining a party in a campaign
  - GM tools for scheduling and managing play sessions
  - Optional character sheet locking between sessions
  - Loot and xp interfaces
  - Item trading

---

## Tech Stack

The project is built using a modern full-stack architecture with a focus on test-driven development, containerization, and continuous deployment.

| Layer                | Technology                                     |
| -------------------- | ---------------------------------------------- |
| **Frontend**         | React, TypeScript, Vite                        |
| **Styling**          | TailwindCSS                                    |
| **Routing**          | React Router                                   |
| **State Mgmt**       | Redux                                          |
| **Testing**          | Vitest (unit), Playwright (E2E)                |
| **Backend**          | Go (Golang)                                    |
| **Architecture**     | Microservices + API Gateway                    |
| **Containerization** | Docker                                         |
| **CI/CD**            | GitHub Actions + GitOps                        |
| **Database + Auth**  | Supabase (PostgreSQL)                          |
| **Dice Rolls**       | random.org API                                 |
| **Deployment**       | GitHub Pages (frontend), Leapcell.io (backend) |
| **API Docs**         | OpenAPI spec + Swagger UI                      |

---

## Deployment Info

- **Frontend URL**: [dc20clerk.dronico.net](https://dc20clerk.dronico.net)
- **Branch Strategy**:
  - `main`: Protected release branch
  - `dev`: Default working branch
  - `feature/*`: Feature branches merged into `dev`
- **CI/CD**:
  - GitHub Actions deploy frontend on push to `main`
  - Backend services deployed to Leapcell.io via GitOps

---

© 2025 Luffy. All rights reserved.

This software is provided for personal use only. The code and implementation are original and owned by the author. However, the DC20 tabletop roleplaying system and its associated rules, mechanics, and intellectual property are owned by their respective creators.

This project is not affiliated with, endorsed by, or sponsored by the creators of DC20. All references to DC20 are made for educational and personal use purposes only.
