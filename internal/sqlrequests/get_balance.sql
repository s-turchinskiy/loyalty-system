select
    coalesce(
            (
                select
                    sum(sum)
                from
                    loyaltysystem.balances where user_id = (select id from loyaltysystem.users where login = 'test'))
        , 0) as current