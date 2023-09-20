-- Criação da tabela de usuários
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL
);

-- Criação da tabela de transferências
CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    id_sender INT NOT NULL,
    id_receiver INT NOT NULL,
    value DECIMAL(10, 2) NOT NULL,
    date TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (id_sender) REFERENCES users(id),
    FOREIGN KEY (id_receiver) REFERENCES users(id)
);


-- Insert User 1
INSERT INTO "users" (name, balance)
VALUES ('User 1', 1000.00);

-- Insert User 2
INSERT INTO "users" (name, balance)
VALUES ('User 2', 750.50);

-- Insert User 3
INSERT INTO "users" (name, balance)
VALUES ('User 3', 500.00);

-- Insert User 4
INSERT INTO "users" (name, balance)
VALUES ('User 4', 300.50);
