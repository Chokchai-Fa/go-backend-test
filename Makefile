server:
	go run cmd/main.go

migrateup:
	migrate -path db/migration -database "postgresql://root:Phukao98765@localhost:5432/svc_backend_team?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:Phukao98765@localhost:5432/svc_backend_team?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:Phukao98765@localhost:5432/svc_backend_team?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:Phukao98765@localhost:5432/svc_backend_team?sslmode=disable" -verbose down 1

.PHONY: server