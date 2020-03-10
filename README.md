# `terraform-provider-harbor`

This is a [Terraform](https://terraform.io) provider for the [Harbor](https://goharbor.io) registry.

Right now it only implements the two things I needed: projects and robot accounts.

This is also one of my first Go project, feel free to tell me if I'm doing things wrong.

The API client is generated using [`go-swagger`](https://github.com/go-swagger/go-swagger).

Steps:
```bash
# clone Harbor
git clone git@github.com:goharbor/harbor.git
# checkout tag, currently generated from: remotes/origin/release-1.10.0
git checkout remotes/origin/release-1.10.0
# generate spec
swagger generate client -f ../harbor/api/harbor/swagger.yaml -t api
```

# License

MIT
