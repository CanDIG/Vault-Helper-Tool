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
From this following [document](https://pkg.go.dev/github.com/hashicorp/vault/api), the first command corresponds to the following call in go:
```
_, err = client.Logical().Write("identity/entity", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}
```
Note that secretData here is:
```
secretData := map[string]interface{}{
		"name": "user",
		"metadata": map[string]interface{}{
			"dataset123": 4,
		},
	}
```
The paramters required for the input are the name of the user (must be created prior) and the metadata (update to user permissions) for the first one and the role (must be created prior to this). 
After updating, we will have to go back to Vault to generate a JWT (using HTTP API):
```
curl -H "X-Vault-Token: insert-client-id-token" http://docker.localhost:8200/v1/identity/oidc/token/insert-role
```
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

The only paramter required here is the role. The output should be a json detailing the role (bound_audiences, user_claim, role_type, policies, ttl). If no role is provided or the role doesn't exist there should be an error. 
This would correspond roughly to call like so
```
secret, err := client.Logical().Read("identity/entity/name/user")
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}
```
### How to `LIST` Tokens

Note that we can simply use:
```
curl \
    --header "X-Vault-Token: ..." \
    --request LIST \
    http://127.0.0.1:8200/v1/identity/entity/name
```
to see the list of users like so:
```
{
  "data": {
    "keys": ["user"]
  }
}

```

The corresponding call using Vault's API is:
```
listSecret, err := client.Logical().List("identity/entity/name")
	if err != nil {
		log.Fatalf("Unable to list secret: %v", err)
	}
```