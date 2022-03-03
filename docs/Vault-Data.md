## Defining Format of the Data
From the [following document](https://github.com/CanDIG/CanDIGv2/blob/develop/docs/authz-permissions.md#draft-v003), this is the shape of the data that the vault helper tool will read:
```
{
  "aud": "cq_candig",
  "exp": 1603988812,
  "iat": 1603902412,
  "iss": "/v1/identity/oidc",
  
  "ga4gh_passport_v1": {
    "ga4gh_visa_v1": {
      "type": "ControlledAccessGrants",
      "value": {
        "dataset1234": {
          "access": 4
        }
      }
    }
  },

  "sub": "b6a4b63c...9a7a247db34f"
}
```
Addtionally, here are some of the paramters of an entity:
## Parameters of an Entity
<ul>
<li>name (string: entity- `UUID`) – Name of the entity.</li>

<li>id (string: `optional`) - ID of the entity. If set, updates the corresponding existing entity.</li>

<li>metadata (key-value-map: {}) – Metadata to be associated with the entity.</li>

<li>policies (list of strings: []) – Policies to be tied to the entity.</li>

<li>disabled (bool: `false`) – Whether the entity is disabled. Disabled entities' associated tokens cannot be used, but are not revoked.</li>
</ul>

[Taken from here](https://www.vaultproject.io/api-docs/secret/identity/entity).

Note that the token above has it's metadata as the dataset numebr and access. 
TO read the metadata in Vault, one could either parse this json or call `vault metadata identity/entity-alias` to ge the metadata. We will further exporte this with the HTTP API when implementing `read`. 