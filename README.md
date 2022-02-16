# Alien attack

## Please look through the Makefile if you'd like to start quick

## In order to run unit tests type:

```
make test
```

## You can run this project in dev mode with:
```
make run-dev
``` 

## You can run this project in dev mode with flexibility:
```
go run -tags=starcraft main.go -w=<file> -a=<file>
``` 

## You can run this project in prod mode by building it first:
```
make build
make run-prod
``` 

## Areas of improvements:
-  Provide more tests for alien entity
-  Refactor cmd/root by extracting functions to be cleaner and more testable
-  Add destroy city functionality(may be cuncurent as well) (right now  only reacts on 10000 steps finished)
-  Could be better optimised, as there are some places where too much of the memory allocations happening,
perhaps using sync.pool would bring some performance boost  

## Problems during execution:
-  Too much time spent on providing world map file
-  Should not jump straight away for cuncurent way of doing things (start simple is better and then refactor on the way) 
-  Perhaps usinig graph was not the best data structure to choose here
