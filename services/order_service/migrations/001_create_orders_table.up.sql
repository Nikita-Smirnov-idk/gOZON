CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    amount INTEGER NOT NULL,
    description TEXT,
    status INTEGER NOT NULL DEFAULT 0
    created_at TIMESTAMPTZ DEFAULT NOW()
);