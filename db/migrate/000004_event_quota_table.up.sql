CREATE TABLE event_quotas (
    event_id UUID REFERENCES events(id) ON DELETE CASCADE, 
    remaining_quota INT NOT NULL DEFAULT 0  
);