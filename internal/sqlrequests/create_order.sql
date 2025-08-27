INSERT INTO loyaltysystem.ordersStatuses (order_id, status_id, user_id)
	VALUES ($1,
	        (select id from loyaltysystem.statuses where status = 'NEW'),
            (select id from loyaltysystem.users where login = $2))
	ON CONFLICT (order_id) DO NOTHING