# Vault-Helper-Tool

This tool allows users to perform CRUD operations using the command-line in Linux. This allows for users to interact with Vault to add/remove/edit permissions for users that set the access level and user authorization. 

# Quick Start

This is a quick mock of how the Vault helper tool will work.
In order to get started, run `go build` in the cli folder to build the CLI.

## How To Use the CLI

Run the script `./cli` to set up a Vault dev server and run the code.

Note, there are 3 commands implemented:
<ul>
<li>`write`: Can use this command as 
```
./cli write {insert-token-here} {json file}
```
or after running the cli as 
```
write {insert-token-here} {json file}
```</li>
<li>`read`: Can use this command as 
```
./cli read {insert-token-here} {user's name}
``` or after running the cli as `read {insert-token-here} {user's name}`, the json file should be provided to it be changed to add more users.</li>
<li>`list`: Can use this command as 
```
./cli list {insert-token-here} {insert-token-here}
``` or after running the cli as 
```
list {insert-token-here}
```.</li>
<li>`help`: Can use this command as `./cli` or `./cli -h` or `./cli help`. This command will show information about the CLI.</li>
</ul>

There are two ways to access the CLI:
<ul>
<li>By running it in interactive mode (note the aliases can be used instead of the full command), for example: 
```
$ ./cli
# Enter command or enter q to quit: write {insert-token-here} example.json
# Enter command or enter q to quit: read {insert-token-here} user
# Enter command or enter q to quit: list {insert-token-here}
# Enter command or enter q to quit: w {insert-token-here} example.json
# Enter command or enter q to quit: r {insert-token-here} user
# Enter command or enter q to quit: l {insert-token-here}
# Enter command or enter q to quit: q
```
</li>
<li>By running it via the command line (note the aliases can be used instead of the full command), for example:
```
$ ./cli write {insert-token-here} example.json
$ ./cli read {insert-token-here} user
$ ./cli list {insert-token-here}
$ ./cli w {insert-token-here} example.json
$ ./cli r {insert-token-here} user
$ ./cli l {insert-token-here}
$ ./cli help
```
</li>
</ul>


## How to Generate a JWT After Writing

Follow the process outlined in the [Setup document](https://candig.atlassian.net/wiki/spaces/CA/pages/623116353/Authorisation+-+Vault+helper+tool) to initialize a user. Then, once you can sucessfully generate the sample JWT provided, complete the following steps to write to Vault and generate the JWT:

In the cli directory:
```
./cli write {insert-token-here} example.json
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