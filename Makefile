.PHONY: migrate-up migrate-down import-b2bclients import-products

migrate-up:
	docker run --network farmaze-backend_farmaze_backend_default \
		-e POSTGRES_USER=$$DB_USERNAME \
		-e POSTGRES_PASSWORD=$$DB_PASSWORD \
		-e POSTGRES_HOST=$$DB_HOST \
		-e POSTGRES_PORT=$$DB_PORT \
		-e POSTGRES_DB=$$DB_NAME \
		-v $(PWD)/db/migrations:/db/migrations/ \
		migrate/migrate:latest \
		-path=/db/migrations/ \
		-database=postgres://$$DB_USERNAME:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable up

migrate-down:
	docker run --network farmaze-backend_farmaze_backend_default \
		-e POSTGRES_USER=$$DB_USERNAME \
		-e POSTGRES_PASSWORD=$$DB_PASSWORD \
		-e POSTGRES_HOST=$$DB_HOST \
		-e POSTGRES_PORT=$$DB_PORT \
		-e POSTGRES_DB=$$DB_NAME \
		-v $(PWD)/db/migrations:/db/migrations/ \
		migrate/migrate:latest \
		-path=/db/migrations/ \
		-database=postgres://$$DB_USERNAME:$$DB_PASSWORD@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable down

import-b2bclients:
	docker exec -i farmaze-backend-db psql -U $(DB_USERNAME) -d $(DB_NAME) -c "\copy b2b_clients (id, company_name, contact_name, email, phone_number) from '/db/data/b2bclients.csv' with (format csv, header true);"

import-products:
	docker exec -i farmaze-backend-db psql -U $$DB_USERNAME -d $$DB_NAME -c "\copy products (id,name, price, description, available_quantity, category,unit) from '/db/data/products.csv' with (format csv, header true);"

cleanup-products:
	docker exec -i farmaze-backend-db psql -U $$DB_USERNAME -d $$DB_NAME -c "TRUNCATE products CASCADE;"
