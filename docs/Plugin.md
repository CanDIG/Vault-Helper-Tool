## Developing Vault Helper Tool as a Plugin
--------------------------------------------
### Takeaways

The Vault plugin system solves the following use cases: allowing for customized authentication and customized secret engines. These do not have much overlap with our use case for the Vault Helper Tool. Thus there is no reason to implement the Vault Helper Tool as a plugin. Nonetheless, for future reference, some documentation about getting started with the Vault plugin system is provided below.

### Background on Vault Plugins
Vault communicates to plugins over a RPC interface, so it is possible to build and distribute a plugin for Vault without having to rebuild Vault itself.

Developing a plugin is simple. The only knowledge necessary to write a plugin is basic command-line skills and basic knowledge of the Go programming language.

The plugin implementation needs to satisfy the interface for the plugin type. (Definitions in the [docs for the backend running the plugin](https://www.vaultproject.io/docs/plugin)).

Plugin backends are the components in Vault that can be implemented separately from Vault's builtin backends. These backends can be either authentication or secrets engines.

The plus-side is that since it is a vault plug-in it will be easier for the user to navigate (As one san simply enable/disable this in Vault in order to keep everything within Vault). 

### How it Works

Each plugin acts as a server, and Vault makes API calls to that server. This is done over a mutually-authenticated TLS connection, all on the local machine. The process is as follows:

<ol>
<li>A Vault administrator registers a plugin the Vault's Plugin Catalog. The registration includes the path to the plugin binary on disk, the SHA256 checksum of the plugin binary, the name of the binary, and optionally the command to execute (defaults to the name of the plugin).</li>

<li>On execution, Vault verifies the plugin is registered (by name) in Vault's Plugin Catalog. Plugins must be registered in the catalog before use.</li>

<li>Vault ensures the checksum of the plugin on disk matches the registered checksum in the Plugin Catalog. This ensures the binary has not been tampered with or altered since installation.</li>

<li>Vault spawns the plugin, passing it a wrapped token containing TLS certificates and a private key. The wrapped token has exactly one use and a very small TTL.</li>

<li>The plugin unwraps the provided wrapped token by making an API call to Vault. The plugin extracts the unique TLS certificates and private key wrapped by the token. The plugin uses these TLS certificates and private key to start an RPC server encrypted with TLS.</li>

<li>Vault and the plugin communicate via RPC over TLS using mutual TLS.</li>
</ol>

### Is this Better than A CLI?

The major benefit of creating a vault plugin versus a CLI is that it will abstract away some of the set-up. Note, with `pathAuthLogin`, login can be implemented with the plugin, like so:

```
func (b *backend) pathAuthLogin(_ context.Context, req *logical.Request, d *framework.FieldData) (*logical.Response, error) {
  password := d.Get("password").(string)

  if subtle.ConstantTimeCompare([]byte(password), []byte("super-secret-password")) != 1 {
    return nil, logical.ErrPermissionDenied
  }

  ttl, _, err := b.SanitizeTTLStr("30s", "30s")
  if err != nil {
    return nil, err
  }

  // Compose the response
  return &logical.Response{
    Auth: &logical.Auth{
      InternalData: map[string]interface{}{
        "secret_value": "abcd1234",
      },
      Policies: []string{"my-policy", "other-policy"},
      Metadata: map[string]string{
        "fruit": "banana",
      },
      LeaseOptions: logical.LeaseOptions{
        TTL:       ttl,
        Renewable: true,
      },
    },
  }, nil
}
```
This code snippet abve implements the login by username and password instead of tokens, and returns the data, policies and time of expiration for the tokes.

Note that since the vault HTTP API is still used in the plug-in, the same calls can be implemented without issue. 
Much of the keycloak portion cannot be abstracted, but the generation of the JWT might be possible. Moreover, since the plugin can be enabled/disabled in Vault, and the user doesn't have to move between Vault and the CLI.

However, I'm not entirely sure how much of the [authorization procedure](https://candig.atlassian.net/wiki/spaces/CA/pages/623116353/WIP+Authorisation+-+Vault+helper+tool#Setup-Vault-for-the-task) can acutally be abstracted by implementing it as a Vault plugin as it is typically used to customize the authentication procedure and/or create a new secret engine.
### How to Get Started
This is some boiler-plate code that is necessary to configure a binary to be a plugin, manage the TLS handshake and serve the proper plugin APIs:
```
func main() {
  apiClientMeta := &pluginutil.APIClientMeta{}
  flags := apiClientMeta.FlagSet()
  flags.Parse(os.Args[1:])

  tlsConfig := apiClientMeta.GetTLSConfig()
  tlsProviderFunc := pluginutil.VaultPluginTLSProvider(tlsConfig)

  if err := plugin.Serve(&plugin.ServeOpts{
    BackendFactoryFunc: Factory,
    TLSProviderFunc:    tlsProviderFunc,
  }); err != nil {
    log.Fatal(err)
  }
}

func Factory(ctx context.Context, c *logical.BackendConfig) (logical.Backend, error) {
  // STUB
  return nil, nil
}
```
The main function above calls Vault's built-in plugin.Serve call, which builds all the required plugin APIs, TLS connections, and RPC server. Note that the factory function is a stub right now. The factory configures the plugin and is called by the backend. Vault backends are implemented as interfaces, so this makes it easier on the developer. 
```
type backend struct {
  *framework.Backend
}

func Backend(c *logical.BackendConfig) *backend {
  var b backend

  b.Backend = &framework.Backend{
    BackendType: logical.TypeCredential,
    AuthRenew:   b.pathAuthRenew,
    PathsSpecial: &logical.Paths{
      Unauthenticated: []string{"login"},
    },
    Paths: []*framework.Path{
      &framework.Path{
        Pattern: "login",
        Fields: map[string]*framework.FieldSchema{
          "password": &framework.FieldSchema{
            Type: framework.TypeString,
          },
        },
        Callbacks: map[logical.Operation]framework.OperationFunc{
          logical.UpdateOperation: b.pathAuthLogin,
        },
      },
    },
  }

  return &b
}
```

 
For futher information on how to implement this, see the following [document for an example](https://github.com/hashicorp/vault-auth-plugin-example).
Another example of how this is used is [Vault Auth Slack](https://github.com/sethvargo/vault-auth-slack).