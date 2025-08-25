select
    (
        select
            sum(sum)
        from
            loyaltysystem.balances) as current,
    (
        select
            sum(w.sum)
        from
            loyaltysystem.balances as w
        where
            w.operation_id = (
                select
                    id
                from
                    loyaltysystem.operations o
                where
                    o.operation = 'WITHDRAWAL')) as withdrawn