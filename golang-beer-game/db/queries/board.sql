-- name: SaveBoard :one 
insert into board(name)
values ($1)
returning *;

-- name: FindBoardById :one
select *
from board
where board_id = $1;

-- name: FindBoardByName :one
select *
from board
where name = $1;

-- name: FindBoardByPlayerId :one
select b.*
from board b
         left join player p on b.board_id = p.board_id
where p.player_id = $1;

-- name: GetRunningBoards :many
select *
from board
where state = 'RUNNING';

-- name: DeleteAllBoards :exec
delete
from board;

-- name: StartBoard :exec
update board
set state   = 'RUNNING',
    is_full = true
where board_id = $1;

-- name: GetAvailableRoles :many
select p.role
from board b
left join player p on b.board_id = p.board_id
where b.board_id = $1
group by p.player_id;
