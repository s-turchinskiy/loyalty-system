select
    u.login
from
    loyaltysystem.ordersstatuses o
        left join loyaltysystem.users u
                  on
                      o.user_id = u.id
where
    o.order_id = $1