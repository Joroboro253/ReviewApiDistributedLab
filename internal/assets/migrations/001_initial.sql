-- +migrate Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT
);

CREATE TABLE Reviews (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE rating (
    id SERIAL PRIMARY KEY,
    review_id INT NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    FOREIGN KEY (review_id) REFERENCES reviews(id)
);



INSERT INTO products (name, description) VALUES ('Тестовый продукт 1', 'Описание тестового продукта 1'), ('Тестовый продукт 2', 'Описание тестового продукта 2')

-- +migrate Down
