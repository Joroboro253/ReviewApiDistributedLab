-- +migrate Up

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    rating INT NOT NULL,
    user_id INT NOT NULL,
    content TEXT NOT NULL CHECK (length(content) <= 255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_reviews_created_at ON reviews (created_at);

CREATE TABLE review_ratings (
    id SERIAL PRIMARY KEY,
    review_id INT NOT NULL,
    user_id INT NOT NULL,
    rating DECIMAL NOT NULL CHECK (rating >= 1.0 AND rating <= 5.0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE
);

CREATE INDEX idx_review_ratings_created_at ON review_ratings (created_at);

-- +migrate Down

DROP TABLE IF EXISTS review_ratings;

DROP TABLE IF EXISTS reviews;
