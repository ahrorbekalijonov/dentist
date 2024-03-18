CREATE TABLE IF NOT EXISTS clients (
    id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    father_name VARCHAR(50),
    phone_number VARCHAR(20),
    address VARCHAR(255),
    birth_date VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS appointments (
    id UUID NOT NULL, 
    client_id UUID NOT NULL,
    date TIMESTAMP,
    diagnostics TEXT,
    treatment TEXT,
    amount INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);