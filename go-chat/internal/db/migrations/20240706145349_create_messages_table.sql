-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages (
  id uuid default gen_random_uuid(),
  sender_id VARCHAR(255) NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  message VARCHAR(255) NOT NULL,
  created_time timestamp default 'now()',
  updated_time  timestamp,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages
-- +goose StatementEnd
