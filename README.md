# Social Network Project

A Facebook-like social network built with Next.js and Go.

## Features

- User Authentication (Registration and Login)
- User Profiles (Public and Private)
- Posts and Comments
- Follow/Unfollow System
- Groups and Events
- Real-time Chat
- Notifications

## Tech Stack

- Frontend: Next.js
- Backend: Go
- Database: SQLite
- Containerization: Docker

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/dgthegeek/social-network.git
   cd social-network
   ```

2. Build and run the Docker containers:
   ```
   docker-compose up --build
   ```

3. Access the application at `http://localhost:3000`

## Features in Detail

### User Authentication

- Registration with email, password, and profile information
- Login with email and password
- Session management using cookies

### User Profiles

- Public and private profile options
- Display user information, posts, followers, and following

### Posts and Comments

- Create posts with text, images, or GIFs
- Comment on posts
- Set post privacy (public, private, or custom)

### Follow System

- Send follow requests
- Accept or decline follow requests
- Automatic following for public profiles

### Groups

- Create and join groups
- Invite users to groups
- Request to join groups
- Create and respond to group events

### Chat

- Real-time private messaging using WebSockets
- Group chat rooms
- Emoji support

### Notifications

- Real-time notifications for various actions (follow requests, group invitations, etc.)

## Database Migrations

The project uses SQL migrations to manage the database schema. Migrations are located in `backend/pkg/db/migrations/sqlite/`.

# social-network
