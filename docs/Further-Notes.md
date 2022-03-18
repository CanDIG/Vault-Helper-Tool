## Further Notes

### Notes and Limitations of Tool
- Need root token to use for now (should theoretically be able to use other tokens as well)
- Cannot add q as command due to how urfave cli is structure, so added with input prompt
- `POST` overwrites metadata, since Vault's API command overwrites metadata
- Whenever adding new datasets, update the template in the [Setup Instructions](https://candig.atlassian.net/wiki/spaces/CA/pages/623116353/Authorisation+-+Vault+helper+tool) accordingly for generating the JWTJWT.

### Avenus for Improvement
- Some code reused between interactive mode and cli. Explore if ur fave cli has a similar interface that can be easily integrated
- Try this tool with non-root token (and correct permissions)