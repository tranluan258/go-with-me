-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages ADD room_id uuid;
ALTER TABLE messages
ADD FOREIGN KEY (room_id) REFERENCES rooms(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
