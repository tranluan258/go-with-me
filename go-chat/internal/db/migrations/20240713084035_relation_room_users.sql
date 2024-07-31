-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_room (
  user_id uuid,
  room_id uuid,
  created_time timestamp default current_timestamp,
  updated_time  timestamp default current_timestamp, 
  PRIMARY KEY(user_id,room_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (room_id) REFERENCES rooms(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
