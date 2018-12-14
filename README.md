# Query Manual

## Row
### Spec:
```
Row(<FIELD>=<VALUE>)
```
### Description:
`Row` returns all keys with value `<VALUE>` on field `<FIELD>` as a set.

### Examples:
Query all phone numbers that have accessed `jd.id`:
```
Row(domain="jd.id")
```

## Union
### Spec:
```
Union([ROW_CALL ...])
```
### Description:
`Union` performs a logical `OR` on the result sets of all `ROW_CALL` queries passed to it.

### Examples:
Query all phone numbers that have accessed `jd.id` or `alibaba.com`:
```
Union(Row(domain="jd.id"), Row(domain="alibaba.com"))
```

## Intersect
### Spec:
```
Intersect([ROW_CALL ...])
```
### Description:
`Intersect` performs a logical `AND` on the result sets of all `ROW_CALL` queries passed to it.

### Examples:
Query all phone numbers that have accessed `jd.id` and `alibaba.com`:
```
Intersect(Row(domain="jd.id"), Row(domain="alibaba.com"))
```

## Difference
### Spec:
```
Difference(<ROW_CALL>, [ROW_CALL ...])
```
### Description:
`Difference` performs a logical `complement` from the first set to the subsequent sets.

### Examples:
Query all phone numbers that have accessed `jd.id` but not `alibaba.com`:
```
Difference(Row(domain="jd.id"), Row(domain="alibaba.com"))
```