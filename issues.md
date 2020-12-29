# MySQL常见问题
:Tech:Application:

#### MySQL不可以一条语句update select同一个表。（sqlalchemy可以。）

>The reason for this error is that MySQL doesn’t allow updates to a table when you are also using that same table in an inner select as your update criteria. The article goes on to provide a solution, which is to use a temporary table.

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

### partitioning(分表):

#### 主要形式

1. 水平分表(horizontal): 把不同主键的数据放到不同表中。
2. 垂直分表(vertical): 把表的不同列分为不同的表。

### sql_mode
Server SQL modes define what SQL syntax MySQL should support and what kind of data validation checks it should perform. 
#### tips 
1. 不同的mode下对sql语法的校验规则不同。如对autoincrement属性插入值是否、not null字段不能为空等的校验反馈不同。有的事warning、有的是直接Error。
2. 一般在测试环境下使用严格模式。
3. sql_mode属性可以动态修改，但可能会引起数据丢失或损坏（loss or corruption of data）。
4. 场景：`global`， `session`。
5. 查询sql_mode状态：
```mysql
// 查询
SELECT @@GLOBAL.sql_mode; 
SELECT @@SESSION.sql_mode;
SELECT @@sql_mode
```
#### strict_mode

sql_mod包含一个 `TRADITIONAL`, `STRICT_TRANS_TABLES`, `STRICT_ALL_TABLES`一种以上设置。


### like 语句相关索引。

1. 使用 `select X%`。
2. 使用*覆盖索引（Covering Index）*。

### clustered index(聚簇索引) & unclustered index（非聚簇索引/HeapTree）

> 1. It is generally faster to read from a clustered index if you want to get back all the columns. You do not have to go first to the index and then to the table.Writing to a table with a clustered index can be slower, if there is a need to rearrange the data.
> 2. Accessing a row through the clustered index is fast because the index search leads directly to the page with all the row data.
> 3. . All indexes other than the clustered index are known as [secondary indexes](https://dev.mysql.com/doc/refman/5.7/en/glossary.html#glos_secondary_index).

#### InnoDB
>
> 1. 默认主键为聚簇索引。
>
> 2. 如果没有主键，则选择首个**unique index**+**not null**作为聚簇索引。
>
> 3. 如果没有会创建隐藏clusetered index。
>
> 4. 如果使用非自增主键（如果身份证号或学号等），由于每次插入主键的值近似于随机，因此每次新纪录都要被插到现有索引页得中间某个位置：
>
>    <img src="../https://raw.githubusercontent.com/e1nfalda/IAaFaJdFLzSk/ignore/uPic/14.png" alt="img" style="zoom:50%;" />
>
>    此时MySQL不得不为了将新记录插到合适位置而移动数据，甚至目标页面可能已经被回写到磁盘上而从缓存中清掉，此时又要从磁盘上读回来，这增加了很多开销，同时频繁的移动、分页操作造成了大量的碎片，得到了不够紧凑的索引结构，后续不得不通过OPTIMIZE TABLE来重建表并优化填充页面。
>
> #### MyISAM
>
> 没有clustered index。

### 其他

- `B-tree`vs`binary-tree`: the *B* stands for: *Boeing*, *balanced*, *broad*, *bushy*, and *Bayer* have been suggested(wiki).

----

[^覆盖索引]: Covering Index， an index that contains all of, and possibly more, the columns you need for your query.

[^聚簇索引vs非聚簇索引]: https://web.archive.org/web/20200602095140/https://www.guru99.com/clustered-vs-non-clustered-index.html
[^ B树 MySQL InnoDB MyASIM]: https://web.archive.org/web/20200426152916/http://blog.codinglabs.org/articles/theory-of-mysql-index.html 

