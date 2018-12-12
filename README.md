# XS

![Alt text](xs.png?raw=true "Clean Architecture")

[![Go Report Card](https://goreportcard.com/badge/github.com/zainul/xs)](https://goreportcard.com/report/github.com/zainul/xs)

XS is golang project starter kit try to follow guideline by Clean Architecture **Uncle Bob**

![Alt text](clean.jpeg?raw=true "Clean Architecture")

The constraints in the Clean Architecture are :
- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
    
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
    
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
    
- Independent of Database. You can swap out PostgreSQL or MySQL, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
    
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

## Project structure

### Configs
- Configs directory is for placed the config file

### Internal
- **Entity**
    - Encapsulate enterprise wide business rules. entity in Go is a set of data structures and functions.

- **Usecase**
    - Usecases is layer contains application specific business rules. It encapsulates and implements all of the use cases of the system 

- **Repository**
    - Repository is layer to communicate between Usecase to storage delivery mechanism

- **Delivery**
    - Delivery mechanism is package to delivery the output and receive the input from any model of delivery (http, grpc, queue database, UI)

- **Mock** 
    - Mock package is package for mocking the repository and usecase , it's for testing purpose
    
- **pkg**
    - Private library for your project

