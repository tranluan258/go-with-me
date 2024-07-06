-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id uuid default gen_random_uuid(),
  username VARCHAR(30) unique NOT NULL,
  password VARCHAR(255),
  full_name VARCHAR(255) NOT NULL,
  avatar VARCHAR(255),
  created_time timestamp default 'now()',
  updated_time timestamp,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
