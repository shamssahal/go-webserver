yourapp/
├── cmd/
│ └── api/
│ └── main.go
│
├── internal/
│ ├── http/
│ │ ├── handlers/
│ │ │ ├── users.go
│ │ │ └── health.go
│ │ │
│ │ ├── middleware/
│ │ │ ├── auth.go
│ │ │ ├── logging.go
│ │ │ └── recover.go
│ │ │
│ │ └── router.go
│ │
│ ├── core/
│ │ └── users/
│ │ ├── service.go
│ │ └── models.go
│ │
│ └── data/
│ ├── postgres.go # placeholder for real DB
│ ├── user_repo.go # in-memory repo or DB repo impl
│ └── migrations/
│ └── 0001_init.sql # placeholder migration file
│
├── pkg/
│ └── response/
│ └── response.go
│
├── go.mod
└── README.md
