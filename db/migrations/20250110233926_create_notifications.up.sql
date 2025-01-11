CREATE TABLE notifications (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   user_id TEXT NOT NULL,
   message TEXT NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
