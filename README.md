# Go Component SDK with Custom WIT

This started out as an example
["Go Component SDK with Custom WIT"](https://github.com/wasmCloud/go/tree/main/examples/component/invoke), <br/>
a WebAssembly component that demonstrates how to use the wasmCloud <br/> 
[Go Component SDK](https://github.com/wasmCloud/go/tree/main/component)
in conjunction with your own custom
[WIT interfaces](https://wasmcloud.com/docs/concepts/interfaces). 

You can find this application's custom interface in `wit/world.wit`:

```wit
interface invoker {
  call: func() -> string;
}
```

## Dependencies

Go 1.23+, tinygo, wasm-tools, wasmCloud Shell (`wash`)

## Run the example

```shell
git clone https://github.com/wasmCloud/go.git
cd examples/component/invoke
```

### Build the component

We will build and deploy this example manually, since we will be using the
[`wash call` subcommand](https://wasmcloud.com/docs/cli/wash#wash-call)
to interact with the application, requiring a stable identity to call.
(There is also the 
[`wash dev` subcommand](https://wasmcloud.com/docs/cli/wash#wash-dev).

Build the component:

```shell
wash build
```

### Start a wasmCloud environment

Start a local wasmCloud environment (using
`-d`/`--detached` to run in the background):

```shell
wash up -d
```

### Deploy the application

Deploy the component using the application manifest (`wadm.yaml`):

```shell
wash app deploy wadm.yaml
```

Rerun ths until the application has reached `Deployed` status:

```shell
wash app list
```

### Invoke the component

Once the application is deployed, you can call the component using
[`wash call`](https://wasmcloud.com/docs/cli/wash#wash-call),
which invokes a function on a component:

```shell
wash call invoke_example-invoker example:invoker/invoker.call
```
```text
Hello from the invoker!
```

### Clean up

You can delete an application from your wasmCloud environment
by referring either to its application name (`invoke-example`)
or the original application manifest:

```shell
wash app delete wadm.yaml
```

Stop your local wasmCloud environment:

```shell
wash down
```

### Bonus: Calling when running with `wash dev`

When running with `wash dev`, wasmCloud uses a generated
ID for the component. If you have `jq` installed,
you can run the following command to call the component:

```bash
wash call "$(wash get inventory -o json | jq -r '.inventories[0].components[0].id')" example:invoker/invoker.call
Hello from the invoker!
```

## 📖 Further reading

For more on custom interfaces, see the
[Interface Developer Guide](https://wasmcloud.com/docs/developer/interfaces/creating-an-interface)
in the wasmCloud documentation. 
