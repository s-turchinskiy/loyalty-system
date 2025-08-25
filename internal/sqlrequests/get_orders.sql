SELECT
    orderstatuses.order_id AS number,
    statuses.status AS status,
    coalesce(balances.sum,0) AS accrual,
    orderstatuses.uploaded_at AS uploaded_at
    FROM loyaltysystem.ordersstatuses as orderstatuses
    inner join loyaltysystem.statuses as statuses ON
        orderstatuses.status_id = statuses.id
    left join loyaltysystem.balances as balances ON
        orderstatuses.order_id = balances.order_id
order by
    orderstatuses.uploaded_at desc
