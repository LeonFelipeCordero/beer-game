create
extension if not exists "pgcrypto";

create table board
(
    board_id    uuid      not null default gen_random_uuid(),
    name        text      not null,
    state       text      not null default 'CREATED',
    is_full     boolean   not null default false,
    is_finished boolean   not null default false,
    created_at  timestamp not null default current_timestamp,
    updated_at  timestamp not null default current_timestamp,
    constraint board_pk primary key (board_id)
);
create unique index board_name_ids on board (name);

create table player
(
    player_id    uuid      not null default gen_random_uuid(),
    name         text      not null,
    role         text      not null,
    stock        bigint    not null,
    backlog      bigint    not null,
    weekly_order bigint    not null,
    last_order   bigint    not null,
    cpu          boolean   not null,
    board_id     uuid      not null,
    created_at   timestamp not null default current_timestamp,
    updated_at   timestamp not null default current_timestamp,
    constraint player_pk primary key (player_id),
    constraint player_board_fk foreign key (board_id) references board (board_id)
);

create table orders
(
    order_id        uuid      not null default gen_random_uuid(),
    amount          bigint    not null,
    original_amount bigint    not null,
    type            text      not null,
    state           text      not null default 'PENDING',
    sender_id       uuid      not null,
    receiver_id     uuid,
    created_at      timestamp not null default current_timestamp,
    updated_at      timestamp not null default current_timestamp,
    constraint order_pk primary key (order_id),
    constraint order_player_sender_fk foreign key (sender_id) references player (player_id),
    constraint order_player_receiver_fk foreign key (receiver_id) references player (player_id)
);