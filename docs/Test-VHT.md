# User Guide on How to Smoke-Test the Vault Helper Tool

## How to run
From the root of the `Candigv2` repository, run 
```
./Vault-Helper-Tool/cli/cli {command} {optional-arguments}
```
## How to use Tool

- To call `write`, use:
```
./Vault-Helper-Tool/cli/cli write {json file}
```
or after running the cli as 
```
write {json file}
```

- To call `read`, use:

```
./Vault-Helper-Tool/cli/cli read {user's name}
```
or after running the cli as 
```
read {user's name}
```

- To call `delete`, use:
```
./Vault-Helper-Tool/cli/cli delete {user's name}
```
or after running the cli as
```
delete {user's name}
```

- To call `list`, use:

```
./Vault-Helper-Tool/cli/cli list
``` 
or after running the cli as 
```
list 
```

- To call `help`, use:

```
./Vault-Helper-Tool/cli/cli -h
``` 
or after running the cli as 
```
./Vault-Helper-Tool/cli/cli 
```

## Examples for Proper usage
- Write:
```
$ ./Vault-Helper-Tool/cli/cli write ../example.json
Secret written successfully.
```
- Read:
```
$ ./Vault-Helper-Tool/cli/cli read entity_1cd0efa6
Connecting to Vault using token in token.txt
{"dataset123":"4","dataset321":"4"}
```
- Delete:
```
$ ./Vault-Helper-Tool/cli/cli delete entity_1cd0efa6
User deleted successfully.
```
- List: 
```
$ ./Vault-Helper-Tool/cli/cli list
Connecting to Vault using token in token.txt
entity_1cd0efa6
{"dataset123":"4","dataset321":"4"}
-------------------------
entity_c65b1f1a
{"dataset123":"1","dataset321":"1"}
-------------------------
```
## How to Trigger Errors

### Incorrect number of arguments
```
$ ./Vault-Helper-Tool/cli/cli write
Connecting to Vault using token in token.txt
2022/04/12 05:49:59 middleware errored: validation failed: file name not provided
```

```
$ ./Vault-Helper-Tool/cli/cli read
Connecting to Vault using token in token.txt
2022/04/12 05:49:32 middleware errored: validation failed: no arguments provided, missing user's name
```

```
./Vault-Helper-Tool/cli/cli delete
Connecting to Vault using token in token.txt
2022/04/12 05:50:35 middleware errored: validation failed: no arguments provided, missing user's name
```

### Wrong file name
```
$ ./Vault-Helper-Tool/cli/cli write non-file.json
Connecting to Vault using token in token.txt
2022/04/12 05:53:07 middleware errored: handling failed: could not open file. open non-file.json: no such file or directory

```

### User does not exist in vault

```
$ ./Vault-Helper-Tool/cli/cli read non-user
Connecting to Vault using token in token.txt
2022/04/12 05:52:36 middleware errored: handling failed: non-user does not exist in Vault.
```

```
$ ./Vault-Helper-Tool/cli/cli delete non-user
Connecting to Vault using token in token.txt
2022/04/14 10:43:15 middleware errored: handling failed: non-user does not exist in Vault.

```

