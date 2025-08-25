INSERT INTO loyaltysystem.ordersStatuses (order_id, status_id)
	VALUES ($1, (select id from loyaltysystem.statuses where status = 'NEW'))
	ON CONFLICT (order_id) DO NOTHING