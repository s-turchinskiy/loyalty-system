select
    coalesce(
            (select sum(sum) from loyaltysystem.balances where user_id = (select id from loyaltysystem.users where login = $1))
        , 0) as current,
    coalesce(
            -1 * (select sum(sum)
             from loyaltysystem.balances
             where operation_id = (select id from loyaltysystem.operations where operation = 'WITHDRAWAL')
             and user_id = (select id from loyaltysystem.users where login = $1)
            ), 0) as withdrawn