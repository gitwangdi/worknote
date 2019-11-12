# Informer and factory

Informer pop the objects from the delta queue and deal with them.  
Factory is a set of informers.

## Informer

AddEventHandler -> Run

### AddEventHandler()

Add a set of funcs to the informer, including AddFunc, UpdateFunc and DeleteFunc.  
The specific function will be called when the event occurs.

### Run(channel)

Get the object from delta queue or wait.

## Factory

A set of Informers.

Contains two main function: Start and InformerFor.

### InformerFor

Add informer to factory and return the **informer**(after that, we can use AddEventHandler).

### Start

Run Run() for all the informers.

## Index

Index is an index to resource, default: \<namespace>/\<name>
