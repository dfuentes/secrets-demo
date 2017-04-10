# secrets-demo

This is a small demo to show how we could leverage credstash/unicreds/stealth to handle secret data.

## Setting up the infra

I've included some terraform config that will set up storage for a dev stealth backend.  You can set it up using aws-vault like:

```bash
$ aws-vault exec dev -- make plan
$ aws-vault exec dev -- make apply
```

## Writing secrets

Once you have the infra up, you can write some secrets like so:

```bash
$ make stealth
$ aws-vault exec dev -- stealth write --environment=development --service=test-service --key=key-name --value=some-value
```

## Loading secrets

I've included a small go script that you can use to load environment variables from stealth.  You can build it with:

```bash
$ make build
```

This will produce a binary called `bootstrap`.  This binary outputs a bash script that exports your secrets into environment variables (prefixed with `_SECRET`).  An example of how you would use this is included in `demo.sh`.  You can run it with

```bash
$ make build
$ aws-vault exec dev -- ./demo.sh
```

Which should show your secret that you wrote in the `Writing secrets` section.  You could imagine doing something like that inside ECS tasks to load secrets for your app, leveraging per-task IAM roles to allow access to the necessary dynamo tables and KMS key.
