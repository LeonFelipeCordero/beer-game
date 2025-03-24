-- name: SavePlayer :one 
insert into player(name,
                   role,
                   stock,
                   backlog,
                   weekly_order,
                   last_order,
                   cpu,
                   board_id)
values ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8)
returning *;

-- name: FindPlayerById :one
select *
from player
where player_id = $1;

-- name: UpdatePlayerNumbers :exec
update player
set stock        = $2,
    backlog      = $3,
    weekly_order = $4,
    last_order   = $5
where player_id = $1;

-- name: FindPlayerByBoardId :many
select *
from player
where board_id = $1;

-- name: DeleteAllPlayers :exec
delete from player;

-- name: GetPlayerByRoleAndBoardId :one
select *
from player
where board_id = $1
and role = $2;