-- +goose Up
-- +goose StatementBegin
ALTER TABLE rooms ADD room_type VARCHAR(20);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
