docker compose up -d
docker exec -it mysql-container bash
mysql -u alphasoft -p dbmusic  < database.sql