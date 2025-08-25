select
    (
        select
            sum(sum)
        from
            loyaltysystem.balances) as current