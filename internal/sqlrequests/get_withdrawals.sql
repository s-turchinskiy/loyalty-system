select
    w.order_id as order,
    -1 * w.sum as sum,
    w.uploaded_at AS processed_at
from
    loyaltysystem.balances as w
where
    w.operation_id = (select id from loyaltysystem.operations where operation = 'WITHDRAWAL')
    and w.user_id = (select id from loyaltysystem.users where login = $1)
order by
    w.uploaded_at desc
