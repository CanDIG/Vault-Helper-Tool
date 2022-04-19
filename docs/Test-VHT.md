# User Guide on How to Test the Vault Helper Tool

## How to run
Ensure that you have followed the commands in `install-docker.md` to initialize Candigv2 with docker, then run 
```
make compose
make init-authx
```
After this, look at the `keys.txt` file in the vault folder to find the root access token. This token will need to be updated in the `token.txt` file in the root directory. 
From the root of the `Candigv2` repository, run 
```
./path-to-cli {command} {optional-arguments}
```
## How to use Tool

- To call `write`, use:
```
./path-to-cli write {json file}
```
or after running the cli as 
```
write {json file}
```

- To call `read`, use:

```
./path-to-cli read {user's name}
```
or after running the cli as 
```
read {user's name}
```

- To call `delete`, use:
```
./path-to-cli delete {user's name}
```
or after running the cli as
```
delete {user's name}
```

- To call `list`, use:

```
./path-to-cli list
``` 
or after running the cli as 
```
list 
```

- To call `updateRole`, use:
```
$ ./path-to-cli updateRole {path-to-json-for-role} {role}
```
or after running the cli as 
```
updateRole {path-to-json-for-role} {role}
```

- To call `help`, use:

```
./path-to-cli -h
``` 
or after running the cli as 
```
./path-to-cli 
```

## Examples for Proper usage
- Write:
```
$ ./path-to-cli write Vault-Helper-Tool/example.json
Secret written successfully.
```
- Read:
```
$ ./path-to-cli read entity_1cd0efa6
Connecting to Vault using token in token.txt
{"dataset123":"4","dataset321":"4"}
```
- Delete:
```
$ ./path-to-cli delete entity_1cd0efa6
User deleted successfully.
```
- List: 
```
$ ./path-to-cli list
Connecting to Vault using token in token.txt
entity_1cd0efa6
{"dataset123":"4","dataset321":"4"}
-------------------------
entity_c65b1f1a
{"dataset123":"1","dataset321":"1"}
-------------------------
```
- updateRole
```
$ ./path-to-cli updateRole researcher Vault-Helper-Tool/role.json
Connecting to Vault using token in token.txt
Role updated successfully.
```
## How to Trigger Errors

### Incorrect number of arguments
```
$ ./path-to-cli write
Connecting to Vault using token in token.txt
2022/04/12 05:49:59 middleware errored: validation failed: file name not provided
```

```
$ ./path-to-cli read
Connecting to Vault using token in token.txt
2022/04/12 05:49:32 middleware errored: validation failed: no arguments provided, missing user's name
```

```
./path-to-cli delete
Connecting to Vault using token in token.txt
2022/04/12 05:50:35 middleware errored: validation failed: no arguments provided, missing user's name
```
```
./path-to-cli updateRole
Connecting to Vault using token in token.txt
2022/04/12 05:50:35 middleware errored: validation failed: no arguments provided, missing filename
```

### Wrong file name
```
$ ./path-to-cli write non-file.json
Connecting to Vault using token in token.txt
2022/04/12 05:53:07 middleware errored: handling failed: could not open file. open non-file.json: no such file or directory

```

### User/Role does not exist in vault

```
$ ./path-to-cli read non-user
Connecting to Vault using token in token.txt
2022/04/12 05:52:36 middleware errored: handling failed: non-user does not exist in Vault.
```

```
$ ./path-to-cli delete non-user
Connecting to Vault using token in token.txt
2022/04/14 10:43:15 middleware errored: handling failed: non-user does not exist in Vault.

```

```
$ ./path-to-cli ur Vault-Helper-Tool/non-role.json researcher
Connecting to Vault using token in token.txt
2022/04/19 08:20:39 middleware errored: handling failed: could not open file. open Vault-Helper-Tool/non-role.json: no such file or directory

```