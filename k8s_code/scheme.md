# Scheme and schema

Schema is a package to deal with group, version and kind.

Scheme is a struct to store information of all the resources we may use.

## Scheme

When we write an operator, we fill the scheme with the object and default func we use. And then transport the scheme to the controller.

### unversionedTypes and unversionedKinds

map[type]gvk , map[kind]type

UnversionedType: the same kinds of the resource should have the same type(struct). (eg. deployment?)

### gvkToType and typeToGVK

### defaulterFuncs

map[type]func

Use AddTypeDefaultingFunc to add defaulterFuncs. In operator, use sheme.Default(object) to set some default field in object.

### fieldLabelConversionFuncs

map[gvk]func

Convert label in another version into this gvk.

AddFieldLabelConversionFunc() to set ...

ConvertFieldLabel() to use ...

### Converter

Convert one type to another type. (Notebook to unstructured object)
