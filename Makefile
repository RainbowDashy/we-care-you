.PHONY: newdb

backend/data.db:
	sqlite3 backend/data.db < backend/migration.sql