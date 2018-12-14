# Query Manual

## Row
### Spec:
```
Row(<FIELD>=<VALUE>)
```
### Description:
`Row` returns all keys with value `<VALUE>` on field `<FIELD>` as a set.

### Examples:
Query all users that have accessed `jd.com`:
```
Row(domain="jd.com")
```

## Union
### Spec:
```
Union([ROW_CALL ...])
```
### Description:
`Union` performs a logical `OR` on the result sets of all `ROW_CALL` queries passed to it.

### Examples:
Query all users that have accessed `jd.com` or `alibaba.com`:
```
Union(Row(domain="jd.com"), Row(domain="alibaba.com"))
```

## Intersect
### Spec:
```
Intersect([ROW_CALL ...])
```
### Description:
`Intersect` performs a logical `AND` on the result sets of all `ROW_CALL` queries passed to it.

### Examples:
Query all users that have accessed `jd.com` and `alibaba.com`:
```
Intersect(Row(domain="jd.com"), Row(domain="alibaba.com"))
```

## Difference
### Spec:
```
Difference(<ROW_CALL>, [ROW_CALL ...])
```
### Description:
`Difference` performs a logical `complement` from the first set to the subsequent sets.

### Examples:
Query all users that have accessed `jd.com` but not `alibaba.com`:
```
Difference(Row(domain="jd.com"), Row(domain="alibaba.com"))
```
