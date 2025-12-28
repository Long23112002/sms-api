-- Tạo database httpsms_dedicated nếu chưa tồn tại
SELECT 'CREATE DATABASE httpsms_dedicated'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'httpsms_dedicated')\gexec

