# Vault-Helper-Tool

This tool allows users to perform CRUD operations using the command-line in Linux. This allows for users to interact with Vault to add/remove/edit permissions for users that set the access level and user authorization. 

# Quick Start


This is a quick mock of how the Vault helper tool will work.
In order to get started, follow the steps outlined in the [Deployment document](https://learn.hashicorp.com/tutorials/vault/getting-started-deploy) provided by Vault. A root token should be provided to the user with the unseal key after running `vault operator unseal`. This token should be added to `secretFile.txt`. **NOTE: Make sure you do NOT have a trailing newline/space/words after the token**.
Then, run `go build` in the cli folder to build the CLI.

## How To Use the CLI

Run the script `./cli` to set up a Vault dev server and run the code.

Note, there are 3 commands implemented:

- `write`: Can use this command as 
```
./cli write {json file}
```
or after running the cli as 
```
write {json file}
```
- `read`: Can use this command as 
```
./cli read {user's name}
```
or after running the cli as 
```
read {user's name}
```
- `list`: Can use this command as 
```
./cli list
``` 
or after running the cli as 
```
list 
```
- `help`: Can use this command as `./cli` or `./cli -h` or `./cli help`. This command will show information about the CLI.

- `delete`: Can use the command as
```
./cli delete {user's name}
```
or after running the cli as
```
delete {user's name}
```

There are two ways to access the CLI:
- By running it in interactive mode (note the aliases can be used instead of the full command), for example: 
```
$ ./cli
# Enter command or enter q to quit: write example.json
# Enter command or enter q to quit: read user
# Enter command or enter q to quit: list
# Enter command or enter q to quit: delete user
# Enter command or enter q to quit: w example.json
# Enter command or enter q to quit: r user
# Enter command or enter q to quit: l
# Enter command or enter q to quit: d user
# Enter command or enter q to quit: q
```

- By running it via the command line (note the aliases can be used instead of the full command), for example:
```
$ ./cli write example.json
$ ./cli read user
$ ./cli list
$ ./cli read user
$ ./cli w example.json
$ ./cli r user
$ ./cli l
$ ./cli d user
$ ./cli help
```
## Verify that Data in Vault

Use the following vault commmand to list out the users in vault:
```
$ vault list identity/entity/name
```
And the following command to read all the information associated with a particual user:
```
$ vault read identity/entity/name/{user-name}
```

## How to Generate a JWT After Writing

Follow the process outlined in the [Setup document](https://candig.atlassian.net/wiki/spaces/CA/pages/623116353/Authorisation+-+Vault+helper+tool) to initialize a user (the root token should be added to `secretFile.txt`). Note, in order to run Vault in deployment, DO NOT use dev mode. Use the [Deployment document](https://learn.hashicorp.com/tutorials/vault/getting-started-deploy) provided by Vault instead. A root token should be provided to the user with the unseal key after running `vault operator unseal`. This token should be added to `secretFile.txt`. 
Following this proceess, once you can sucessfully generate the sample JWT provided, complete the following steps to write to Vault and generate the JWT:

In the cli directory:
```
./cli write example.json
```
Then, curl Keycloak:
```
curl -L -X POST 'http://docker.candig:8080/auth/realms/master/protocol/openid-connect/token' \
-H 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'client_id=candig' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_secret=iKq66fdBDSbNwiZZ7ntrdBxGbLJrWjwG' \
--data-urlencode 'scope=openid' \
--data-urlencode 'username=user' \
--data-urlencode 'password=user'
```
Then, keep the `id-token` generated, and curl Vault after modifying the `payload.json` file to be:
```
{
    "jwt": "insert-id-token-here",
    "role": "researcher"
}
```
After that, curl Vault:
```
curl --request POST --data @payload.json http://docker.candig:8200/v1/auth/jwt/login
```
Keep the `client-token` and modify the following command with it:
```
curl -H "X-Vault-Token: "insert-client-token-here" http://docker.localhost:8200/v1/identity/oidc/token/researcher
```
This should generate the JWT, and go to [JWT.io](https://jwt.io/) to verify if it is correct.

