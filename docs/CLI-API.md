### How to Update/`POST`

Note that the method to create or update an entity is `POST` at the path of `/identity/entity` in the identity secrets engine. Roughly speaking, the intention of the `POST` command in the helper tool is to update user permissions.
From the vault setup task, the first thing to do is to `vault write` the user permissions and then generate the JWT using the client-id token.
To update permissions (using CLI):
```
vault write identity/entity -<<EOF
{
    "name": "user",
    "metadata": {"dataset123": 4}
}
EOF
```
To generate a JWT (using HTTP API):
```
curl -H "X-Vault-Token: insert-client-id-token" http://docker.localhost:8200/v1/identity/oidc/token/insert-role
```
From this following [document](https://pkg.go.dev/github.com/hashicorp/vault/api#Logical.List), the first command corresponds to the following call in go:
```
_, err = client.Logical().Write("identity/entity", inputData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}
```
The paramters required are the name of the user (must be created prior) and the metadata (update to user permissions) for the first one and the role (must be created prior to this). 
The output should be of the form:
```
{
   "request_id":"ecc7015a-10de-79d1-6158-4c74d896c890",
   "lease_id":"",
   "renewable":false,
   "lease_duration":0,
   "data":{
      "client_id":"candig",
      "token":"eyJhbGciOiJSUzI1NiIsImtpZCI6IjY2NmRhNGVkLTllMWItODYyYy1hZWI2LTkxZTVlZjZmMTUxMiJ9.eyJhdWQiOiJjYW5kaWciLCJleHAiOjE2NDQ2MzQzNzAsImdhNGdoX3Bhc3Nwb3J0X3YxIjp7ImdhNGdoX3Zpc2FfdjEiOnsidHlwZSI6IkNvbnRyb2xsZWRBY2Nlc3NHcmFudHMiLCJ2YWx1ZSI6eyJkYXRhc2V0MTIzIjp7ImFjY2VzcyI6IjQifSwiZGF0YXNldDMyMSI6eyJhY2Nlc3MiOiIifX19fSwiaWF0IjoxNjQ0NTQ3OTcwLCJpc3MiOiJodHRwOi8vMC4wLjAuMDo4MjAwL3YxL2lkZW50aXR5L29pZGMiLCJuYW1lc3BhY2UiOiJyb290Iiwic3ViIjoiZjhjOTg2MTUtZjMwMi0wMjkyLTRhODEtMTBmODMzZmFjZTZkIn0.GAN0WNNCbLSBOmPfk7sKXhU1jm2MT5QSVa5oIMtxqNfnnzshbZcLoxl5o3vEfXu5uakKpDGE8e_CG3El1iwNAscMJ3sSJXAUdVpDN0kz2SkOPrLGR279MZRTtn6pVujArnscu-ult600b2SHm6O9ElPd8sP7pk_3wTTLMnxpTNESQgDxhpYOG18N1hgA7_ABNPqszcdcwRPr5woLcwI_TfGosqXkbN3RNyvhzcVndk3EOpsUuvnIGkyydCXHM13ICd2qUc-soDyQRyGieKSZ97nkbWxY5ZOibfIwROLtqo9JCbGIbI_-4pw8s0d3RTErjjOBPuEUw5gmt0sUASsOLQ",
      "ttl":86400
   },
   "wrap_info":null,
   "warnings":null,
   "auth":null
}
```
What we need is the `token` field that contains the new JWT.
Note that the input paramters might error if the name of the user is incorrect (not created) or not provided, if the metadata is not provided or if the role provided does not exist. 
It might be a little difficult to evaulate whether the user/role is valid or not prior to calling the API.
### How to Read/`GET` a Role

```
curl \
    --header "X-Vault-Token: $VAULT_TOKEN" \
    http://docker.localhost:8200/v1/identity/oidc/token/insert-role
```
The only paramter required here is the role. The output should be a json detailing the role (bound_audiences, user_claim, role_type, policies, ttl). If no role is provided or the role doesn't exist there should be an error. 
This would correspond roughly to call like so
```
inputData, err := client.Logical().Read("identity/entity")
    if err != nil {
        panic(err)
    }
    fmt.Println(data.inputData)
```
### How to `LIST` Tokens

Note that we can simply use:
```
vault auth list -format=json
```
to see the list of users like so:
```
{
  "jwt/": {
    "uuid": "e0eb095f-b9f9-e3ff-23d9-416c41b0b6b1",
    "type": "jwt",
    "description": "",
    "accessor": "auth_jwt_d99aba4e",
    "config": {
      "default_lease_ttl": 0,
      "max_lease_ttl": 0,
      "force_no_cache": false,
      "token_type": "default-service"
    },
    "options": null,
    "local": false,
    "seal_wrap": false,
    "external_entropy_access": false
  },
  "token/": {
    "uuid": "ac8efec2-9ae0-2675-4f69-9ddd5ee31253",
    "type": "token",
    "description": "token based credentials",
    ...snip...
    "seal_wrap": false,
    "external_entropy_access": false
  }
}
```
Vault also has an HTTP API that can call list. This will list all of the [roles](https://www.vaultproject.io/api/auth/jwt#groups_claim).
```
curl \
    --header "X-Vault-Token: insert-client-id-token" \
    --request LIST \
    http://docker.localhost:8200/v1/identity/oidc/token/insert-role
```
The corresponding call is:
```
List, err := client.Logical().List("identity/entity")
    if err != nil {
        panic(err)
    }
    fmt.Println(List)

```