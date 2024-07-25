-- +goose Up
-- +goose StatementBegin
create table chats
(
    id       serial primary key,
    name     text not null,
    users_ids int[]
);
create table messages
(
    id      serial,
    chat_id int references chats (id) on delete cascade,
    user_id int not null,
    text    text,
    created_at timestamp not null default now(),
    primary key (id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
DROP TABLE chats;
-- +goose StatementEnd
