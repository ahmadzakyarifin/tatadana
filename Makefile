.PHONY: up down build migrate migrate-create

# Menjalankan seluruh aplikasi (MySQL, Goose, Backend) di background
up:
	docker compose --env-file backend/.env up -d

# Mematikan seluruh aplikasi
down:
	docker compose --env-file backend/.env down

# Membangun ulang docker image backend (jika ada perubahan kode)
build:
	docker compose --env-file backend/.env build

# Menjalankan perintah Goose apapun (status, down, version, dll)
# Cara pakai: make migrate cmd=status
#             make migrate cmd=down
migrate:
	docker compose --env-file backend/.env run --rm migrate goose $(cmd)

# Perintah khusus untuk membuat file migrasi baru
# Cara pakai: make migrate-create name=add_users_table
migrate-create:
	docker compose --env-file backend/.env run --rm migrate goose create $(name) sql
