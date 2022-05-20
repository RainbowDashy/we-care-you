data.db: migration.sql
	sqlite3 data.db < migration.sql

.PHONY: generate-test-data
generate-test-data: data.db testdata.sql
	sqlite3 data.db < testdata.sql
