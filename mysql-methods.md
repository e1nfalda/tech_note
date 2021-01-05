# MySql 常用方法
<!-- :Tech:Application: -->

- `group_concat(COLUMN)`搜索结果拼接：

    `select group_concat(id,"_",type_id) from table;`-> `"1_1,2_1,3_type,4,5,..."`

- `concat(row1, row2,str...)`: 字符串拼接

    `select concat(id,",", type_id)`-> `"1,type1"`; ` "2,type2"`...
