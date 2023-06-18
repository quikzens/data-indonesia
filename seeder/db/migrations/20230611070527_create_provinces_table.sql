-- migrate:up
CREATE TABLE IF NOT EXISTS provinces
(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE IF EXISTS provinces;
