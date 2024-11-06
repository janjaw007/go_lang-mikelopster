# Project Name

## Description

A brief description of what the project does.

## Prerequisites

- Docker installed
- PostgreSQL running in Docker container

## Setting Up the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/janjaw007/go_lang-mikelopster.git
   cd go_lang-mikelopster
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up
   ```

## Backup and Restore Scripts

### Backup Script

To back up the PostgreSQL database to a `backup.sql` file, run the following command:

### Restore Script

```bash
docker exec -t <container_name> pg_dumpall -c -U <POSTGRES_USER> > backup.sql

```

## old

```bash

docker exec -i <container_name> psql -U <POSTGRES_USER> -f backup.sql
```

## new

```bash
 docker exec -i postgres psql -U myuser -d mydatabase -f /backup.sql
```
