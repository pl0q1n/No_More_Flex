docker-compose up -d
docker cp scripts/init_db.sql nmf_postgres:init_db.sql
docker exec -u postgres nmf_postgres psql postgres postgres -f init_db.sql
