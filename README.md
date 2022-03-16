# Vault-Helper-Tool

This tool allows users to perform CRUD operations using the command-line in Linux. This allows for users to interact with Vault to add/remove/edit permissions for users that set the access level and user authorization. 

# Quick Start

In order to get started, run `go build` to build the CLI.

## How To Use

Run the script `./cli` to set up a Vault dev server and run the code.

Note, there are 3 commands implemented:
<ul>
<li>`write`: Can use this command as `./cli write {json file}` or after running the cli as `write {json file}`</li>
<li>`read`: Can use this command as `./cli read {user's name}`(This will print out sample user's metadata) or after running the cli as `read {user's name}`, which will mimic Vault, in that the json file can be changed to add more users.</li>
<li>`list`: Can use this command as `./cli list` (This will print out sample user's metadata) or after running the cli as `list`, the latter will mimic Vault by printing out all the users added to the CLI.</li>
<li>`help`: Can use this command as `./cli` or `./cli -h`. This command will show information about the CLI.</li>
</ul>

## How to Test

You can change the `example.json` file to add/remove datasets and change the user's name. Then, simply run `./cli write example.json` to add the user to vault. You can also create other JSON files to add other users (make sure to keep the same structure in the example.json file).
