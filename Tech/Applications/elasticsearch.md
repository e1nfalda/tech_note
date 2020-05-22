| ES | Index（type） | document | field  | schema | DSL |
|:-:|---|:-:|---|---|---|
| RDBMS | database | row      | column | mapping | SQL |

### 反向索引

### beatmap

### Mapping: 

 process of defining how a document, and the fields it contains, are stored and indexed.

#### MetaFields

Meta-fields are used to customize how a document’s metadata associated is treated.

[meta fields 5大类常用属性](file:///Users/yao/Dropbox/文章/Tech/ES/meta_fields.pdf)

#### Fields or Property





### `enabled` vs `index`

When settings `enabled` to false, you tell ES to completely ignore the parsing of the field, so it will neither be analyzed, nor indexed not stored (except in he `_source` field of course).

So, ES is not even aware that the field exists, and thus, it handles that case as if you were querying on any other non-existent field, basically as if the source didn't even contain the field. Result: ES doesn't return any document.

When setting `index` to false, ES is aware that the field exists (via the mapping), but it knows that it shouldn't be indexed. So when you query on it, ES tells you that you cannot do it since you've decided not to index that field. That's why ES throws an error since you're breaking the contract that you've declared in your mapping.



___

[^基本概念]: [10个基本概念](file:///Users/yao/Dropbox/文章/Tech/ES/10 Elasticsearch Concepts You Need to Learn _ Logz.io __ Reader View.pdf)

