##### MySQL不可以一条语句update select同一个表,解决方案如下
> The reason for this error is that MySQL doesn’t allow updates to a table when you are also using that same table in an inner select as your update criteria. The article goes on to provide a solution, which is to use a temporary table.

Using this example, your update should be this:

```sql
update foo
set bar = bar - 1
where baz in
(
  select baz from
  (
    select baz
    from foo
    where fooID = '1'
  ) as arbitraryTableName
)
```

