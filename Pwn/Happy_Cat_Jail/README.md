Flag:
```
^THJCC\{iT'[1-9][a-z][1-9][a-zA-Z][1-9]Sm[1-9]Y_gO[a-zA-Z]!!!!!L[a-z]AnG\}$
```

PoC:
```go
type catInterface struct {
	t uintptr
	v unsafe.Pointer
}

p := unsafe.Pointer(&target)

iface := (*catInterface)(p)

catStr := (*string)(iface.v)

fmt.Println(*catStr)
EOF
```
```go
fmt.Println(target.(*secret).flag)
_ = unsafe.Sizeof(0)
EOF
```
