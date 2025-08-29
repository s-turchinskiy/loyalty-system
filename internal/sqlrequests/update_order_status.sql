UPDATE loyaltysystem.ordersStatuses
SET status_id =   (select id from loyaltysystem.statuses where status = $2)
WHERE order_id = $1;