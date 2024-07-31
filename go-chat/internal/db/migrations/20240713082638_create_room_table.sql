-- +goose Up
-- +goose StatementBegin
CREATE TABLE rooms (
  id uuid default gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  created_time timestamp default current_timestamp,
  updated_time  timestamp default current_timestamp, 
  PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
