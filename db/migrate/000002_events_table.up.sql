CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    start_date DATE,
    end_date DATE,
    start_time TIME,
    end_time TIME,
    price DECIMAL(10, 2),
    quota INT,
    organizer VARCHAR(255)
);