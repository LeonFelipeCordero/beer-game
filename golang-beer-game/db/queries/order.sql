-- name: SaveOrder :one 
insert into orders(amount, original_amount, type, sender_id, receiver_id)
values ($1, $1, $2, $3, $4)
returning *;

-- name: FindOrderById :one
select *
from orders
where order_id = $1;

-- name: FindOrderByBoardId :many
with boad_players as (select b.board_id,
                             player_id
                      from board b
                               left join player p on b.board_id = p.board_id
                      where b.board_id = $1)
select o.*
from orders o,
     boad_players bp
where o.sender_id = bp.player_id
   or o.receiver_id = bp.player_id
group by o.order_id;

-- name: FindOrderByPlayerId :many
select *
from orders
where sender_id = $1
   or receiver_id = $1;

-- name: MarkAsFilled :one
update orders
set state  = 'DELIVERED',
    amount = $2
where order_id = $1
returning *;

-- name: DeleteAllOrders :exec
delete
from orders;