-- +goose Up
-- +goose StatementBegin
create table if not exists chats
(   
    id uuid PRIMARY KEY,
    user_id uuid not null,
    only_read boolean not null,
    time_to_life time 
);

create table if not exists messages
(   
    id uuid PRIMARY KEY,
    chat_id uuid,
    msg text,
    FOREIGN KEY (chat_id) REFERENCES chats(id)
);

-- +goose StatementEnd

