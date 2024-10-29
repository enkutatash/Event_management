CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(id),  
    event_id UUID REFERENCES events(id), 
    booking_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    number_of_tickets INT DEFAULT 1
);