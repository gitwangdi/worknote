## Tip  

Swapper(slice):func(i,j int) return a swapper function for the slice.

Copy(dst,src) copy array or slice.  

## Map

```go
//  ep.

iter := reflect.ValueOf(m).MapRange() // New a map item iterator. We should use Next() immediately. 
for iter.Next() { // Point to the next item.
    k := item.Key()
    v := item.Value()
    ...
}

```

## Type  

TypeOf(interface{}):type

### method

Mothed: Method(int) MethodByName(string) NumMethod()

Func: IsVariadic() In(int) Out(int) NumIn() NumOut()  

Struct: Field() FieldByIndex() FieldByName() FiedlByNameFunc(func(string) bool) 

Elem() type's element type.(Array/slice/map/chan/ptr)

...

### type constructor

* ArrayOf(int, type)  
* ChanOf(ChanDir, type)  
* FuncOf([]type, []type, bool)  
* MapOf(type, type)  
* PtrTo(type)  
* SliceOf(type)  
* StructOf([]StructField) 

## Value

operation for all kinds of interface.

## Usage

TypeOf()  
ValueOf()  
Value.Interface()  
Value = New(Type)  

