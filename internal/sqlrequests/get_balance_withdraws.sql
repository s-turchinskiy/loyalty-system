select
    coalesce(
            (select sum(sum) from loyaltysystem.balances)
        , 0) as current,
    coalesce(
            (select sum(sum)
             from loyaltysystem.balances
             where operation_id = (select id from loyaltysystem.operations where operation = 'WITHDRAWAL')
            ), 0) as withdrawn