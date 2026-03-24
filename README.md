# Events Management RESTful API

RESTful API untuk mengelola events/kegiatan dengan CRUD operations menggunakan Go + Gin framework, terintegrasi dengan Docker dan GitHub Actions.

[![CI - Unit Testing](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/ci.yml/badge.svg)](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/ci.yml)
[![CS - Security Scan](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/security.yml/badge.svg)](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/security.yml)

## 1. Deskripsi Project

**Events Management API** adalah RESTful API sederhana untuk mengelola data events/kegiatan. API ini mendukung operasi CRUD (Create, Read, Update, Delete) dengan response format JSON standar.

**Fitur:**
- GET /api/events - Mendapatkan semua events
- GET /api/events/:id - Mendapatkan event berdasarkan ID
- POST /api/events - Membuat event baru
- PUT /api/events/:id - Update event
- DELETE /api/events/:id - Hapus event
- GET /health - Health check endpoint

## 2. Dokumentasi API

### Base URL
```
Development: http://localhost:8080
Docker: http://localhost:8080
```

### Endpoints

| Method | Endpoint | Description | Status Code |
|--------|----------|-------------|-------------|
| GET | `/api/events` | Get all events | 200 OK |
| GET | `/api/events/:id` | Get event by ID | 200 OK / 404 Not Found |
| POST | `/api/events` | Create new event | 201 Created / 400 Bad Request |
| PUT | `/api/events/:id` | Update event | 200 OK / 404 Not Found |
| DELETE | `/api/events/:id` | Delete event | 200 OK / 404 Not Found |
| GET | `/health` | Health check | 200 OK |

### Request/Response Examples

#### GET /api/events

**Response (200 OK):**
```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "title": "Tech Conference 2026",
      "description": "Annual technology conference",
      "date": "2026-04-15",
      "location": "Jakarta Convention Center",
      "created_at": "2026-03-24T10:00:00Z"
    }
  ]
}
```

#### POST /api/events

**Request Body:**
```json
{
  "title": "Tech Conference 2026",
  "description": "Annual technology conference",
  "date": "2026-04-15",
  "location": "Jakarta Convention Center"
}
```

**Response (201 Created):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "Tech Conference 2026",
    "description": "Annual technology conference",
    "date": "2026-04-15",
    "location": "Jakarta Convention Center",
    "created_at": "2026-03-24T10:00:00Z"
  }
}
```

#### PUT /api/events/:id

**Request Body:**
```json
{
  "title": "Updated Event Title"
}
```

**Response (200 OK):**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "title": "Updated Event Title",
    "description": "Annual technology conference",
    "date": "2026-04-15",
    "location": "Jakarta Convention Center",
    "created_at": "2026-03-24T10:00:00Z"
  }
}
```

#### DELETE /api/events/:id

**Response (200 OK):**
```json
{
  "status": "success",
  "message": "Event deleted successfully"
}
```

#### GET /health

**Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2026-03-24T10:00:00Z"
}
```

#### Error Response (400 Bad Request):
```json
{
  "status": "error",
  "message": "Title is required"
}
```

#### Error Response (404 Not Found):
```json
{
  "status": "error",
  "message": "Event not found"
}
```

## 3. Panduan Instalasi (Docker)

### Prerequisites
- Docker Desktop installed
- Docker Compose installed

### Langkah-langkah

1. **Clone repository**
```bash
git clone <repository-url>
cd restful-api-ppl
```

2. **Build dan Run dengan Docker Compose**
```bash
docker-compose up --build
```

3. **Run di background**
```bash
docker-compose up -d
```

4. **Stop containers**
```bash
docker-compose down
```

5. **View logs**
```bash
docker-compose logs -f api
```

### Port Information

| Service | Host Port | Container Port | Description |
|---------|-----------|----------------|-------------|
| API | 8080 | 8080 | RESTful API endpoint |

Akses API di: `http://localhost:8080`

### Testing API

```bash
# Get all events
curl http://localhost:8080/api/events

# Create event
curl -X POST http://localhost:8080/api/events \
  -H "Content-Type: application/json" \
  -d '{"title":"Tech Conference","description":"Annual conference","date":"2026-04-15","location":"JCC"}'

# Get event by ID
curl http://localhost:8080/api/events/1

# Update event
curl -X PUT http://localhost:8080/api/events/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Event"}'

# Delete event
curl -X DELETE http://localhost:8080/api/events/1
```

## 4. Alur Kerja Git

### Branch Structure

```
main (production-ready)
  └── develop (integration branch)
        └── feat/* (feature branches)
```

| Branch | Description |
|--------|-------------|
| `main` | Production-ready code |
| `develop` | Integration branch untuk development |
| `feat/*` | Feature branches (ex: feat/add-events-crud) |

### Conventional Commits

Format: `<type>: <description>`

| Type | Description | Example |
|------|-------------|---------|
| `feat` | New feature | `feat: add events CRUD endpoints` |
| `fix` | Bug fix | `fix: handle invalid event ID` |
| `docs` | Documentation | `docs: update README with API docs` |
| `style` | Formatting | `style: format go code` |
| `refactor` | Code refactoring | `refactor: extract validation logic` |
| `test` | Adding tests | `test: add unit tests for handlers` |
| `chore` | Build process | `chore: add docker-compose config` |
| `ci` | CI configuration | `ci: add GitHub Actions workflow` |

### Contoh Commit

```bash
git commit -m "feat: add event model and struct"
git commit -m "feat: implement GET /api/events endpoint"
git commit -m "feat: implement POST /api/events endpoint"
git commit -m "test: add unit tests for event handlers"
git commit -m "chore: add Dockerfile configuration"
git commit -m "ci: add GitHub Actions CI workflow"
```

### Workflow

```bash
# Create feature branch dari develop
git checkout develop
git checkout -b feat/add-events-crud

# Make changes dan commit
git add .
git commit -m "feat: add event handlers"

# Push dan create Pull Request
git push -u origin feat/add-events-crud

# Setelah PR approved, merge ke develop
```

### Commit History (Bukti)

```
6c9cb1d fix: correct build command in CI workflow
1a7abb7 fix: update CI workflow and README badges
3e38841 docs: add API documentation
83c1317 ci: add GitHub Actions workflows
f5d7133 chore: add Docker configuration
9d0e49a test: add unit tests for events handlers
ccd229d feat: setup routes and main entry point
3ee7063 feat: implement events CRUD handlers
f56c0b5 feat: add event model and response struct
4be16b4 Initial commit
```

## 5. Status Automasi (GitHub Actions)

### Workflows

#### CI - Unit Testing (`.github/workflows/ci.yml`)

**Trigger:** Push/PR ke `main` dan `develop`

**Jobs:**
- Setup Go environment
- Cache Go modules
- Download dependencies
- Run unit tests dengan coverage
- Build application

**Status:** [![CI - Unit Testing](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/ci.yml/badge.svg)](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/ci.yml)

#### CS - Security Scan (`.github/workflows/security.yml`)

**Trigger:** Push/PR ke `main` dan `develop`

**Jobs:**
- Setup Go environment
- Install gosec
- Run security scan untuk vulnerability

**Status:** [![CS - Security Scan](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/security.yml/badge.svg)](https://github.com/Faruuuqqq/restful-api-ppl/actions/workflows/security.yml)

### Local Testing

```bash
# Run tests
go test -v ./...

# Build
go build -v ./...

# Run locally
go run ./src
```

## Tech Stack

- **Language:** Go 1.25
- **Framework:** Gin-Gonic
- **Testing:** testify
- **Security Scan:** gosec
- **Container:** Docker + Docker Compose
- **CI/CD:** GitHub Actions

## Project Structure

```
restful-api-ppl/
├── .github/
│   └── workflows/
│       ├── ci.yml
│       └── security.yml
├── src/
│   ├── main.go
│   ├── handlers/
│   │   └── events.go
│   ├── models/
│   │   └── events.go
│   └── routes/
│       └── routes.go
├── tests/
│   └── events_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## License

MIT License
