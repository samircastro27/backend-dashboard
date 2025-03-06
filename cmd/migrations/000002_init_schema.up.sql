-- Habilitar la extensión para UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Crear la tabla clients
CREATE TABLE IF NOT EXISTS clients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    login VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    company VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    progress INT NOT NULL,
    created_at DATE NOT NULL
);

-- Insertar datos iniciales (Seeder)
INSERT INTO clients (id, login, name, company, city, progress, created_at)
VALUES
    ('679d1303-9d92-4922-b78b-b874bd868ef1', 'scc.ing27', 'Samir Castro', 'Cuemby', 'Barranquilla', 70, '2023-03-03'),
    ('067e2106-e331-4043-b960-9a5aca93d9aa', 'mapa', 'Mariana Galindo', 'Big Bang', 'Soledad', 68, '2023-12-01'),
    ('90f17b01-495c-4d48-a9a9-a71a93052713', 'adbem', 'Kim Adbem', 'NN', 'New York', 38, '2023-05-04'),
    ('ee43a2d1-9865-43ab-b01c-a7a51f8dd25a', 'cr7', 'Lee Son', 'Seul LLC', 'Seul', 49, '2023-05-18');

