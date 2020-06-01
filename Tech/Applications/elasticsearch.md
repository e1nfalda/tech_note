| ES | Index（type） | document | field  | schema | DSL |
|:-:|---|:-:|---|---|---|
| RDBMS | database | row      | column | mapping | SQL |

### 反向索引

### beatmap

### Mapping: 

 process of defining how a document, and the fields it contains, are stored and indexed.

#### MetaFields

Meta-fields are used to customize how a document’s metadata associated is treated.

> 1. 针对整个document的设置。类似RDBMS的表属性。

| 分类                | 属性值                        |
| ------------------- | ----------------------------- |
| identity meta_field | _id, _index, _type            |
| document source     | _source, _size(document size) |
| indexing            | _field_names, _ignored        |
| routing             | _routing(to_shard)            |
| other               | _meta                         |

#### Mapping Parameter

> 1. 针对每个field设置的属性。

#### Fields or Property

`store`: 指定某field是否存储。如果fieldA指定，则可以检索,以列表形势返回。

> 很多场景`_source`可以替换store作用，但有些场景需要`_source:{enabled: false}`时，则需要明确哪些字段需要store。

`index`: 当设置*false*时即明确field不需要检索，如果查询语句该field，会报错。

`enabled`: 是否自动检索全部field。只检索top-level，不检索Object内部。

`dynamic`: 是否自动检索mapping新增field。true:检索；false: 忽略；strict: 报错。 



___

[^基本概念]: [10个基本概念](file:///Users/yao/Dropbox/文章/Tech/ES/10 Elasticsearch Concepts You Need to Learn _ Logz.io __ Reader View.pdf)
[^ metafield]: [meta fields 5大类常用属性](file:///Users/yao/Dropbox/文章/Tech/ES/meta_fields.pdf)

