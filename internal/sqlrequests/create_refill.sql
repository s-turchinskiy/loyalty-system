INSERT INTO loyaltysystem.balances (order_id, sum, operation_id, user_id)
	VALUES ($1,
	        $2,
	        (select id from loyaltysystem.operations where operation = 'REFILL'),
           $3
           )