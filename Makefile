all:

api:
	cd api && make

web:
	cd web && npm run dev

.phony: all api web