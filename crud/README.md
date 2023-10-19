# CRUD

Implements a simple REST API using `gin` to provide an example
of using `gorm` as an ORM.

This example, focuses on using an existing schema. gorm, like other ORMs such as
`sqlalchemy`, `hibernate`, etc, supports generating DB object from data models.  However,
in most cases, I like to perform DDL at the database. In other words, I like to 
model the database with native database tools, create the various DML and DDL in native
database files.  Versus generating a schema via the ORM.

This approach provides greater separation of roles.  Although in most cases,
I perform both the DBA and software development roles myself.

I used [gorm gen tool](https://gorm.io/gen/gen_tool.html) to reverse engineer the
schema to objects. 

# packages
* [gorm](https://gorm.io/docs/) - Go ORM 
* [gin](https://gin-gonic.com/docs/) - popular HTTP web framework
* [godotenv](https://github.com/joho/godotenv) - Nice Ruby dotenv project 
* [pq](https://github.com/lib/pq) - Postgres 

