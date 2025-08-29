select
    o.order_id,
    s.status,
    o.user_id
from
    loyaltysystem.ordersstatuses as o
        inner join loyaltysystem.statuses as s
                   on
                       o.status_id = s.id
where
    o.status_id in (
        select
            id
        from
            loyaltysystem.statuses s
        where
            s.status in ('NEW', 'REGISTERED', 'PROCESSING'))
order by
    o.uploaded_at
