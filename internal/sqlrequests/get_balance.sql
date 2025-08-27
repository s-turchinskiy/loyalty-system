select
    (
        select
            sum(sum)
        from
            loyaltysystem.balances where user_id = (select id from loyaltysystem.users where login = $1)) as current