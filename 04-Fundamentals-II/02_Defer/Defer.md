# Defer

The builtin keyword defer lets you do stuff that is executed when a function is returning.


```golang

func GetPerson(id string), Person, error {
	
	con := openDatabaseConnection()
	defer con.Close()
	
	return getPerson(con, id)
	
}

```

## Dealing with multiple resources

```golang

func GetPerson(id string), Person, error {
	
	con := openDatabaseConnection()
	defer con.Close()
	
	pst := createPreparedStatement(con, "Select * from Person where id = ?", id)
	defer pst.close
	
	return ...
	
}

```
