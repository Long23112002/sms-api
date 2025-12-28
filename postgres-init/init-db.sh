#!/bin/bash
set -e

# Đợi PostgreSQL sẵn sàng
until pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB"; do
  sleep 1
done

# Tạo database nếu chưa tồn tại
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    SELECT 'CREATE DATABASE httpsms_dedicated'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'httpsms_dedicated')\gexec
EOSQL

