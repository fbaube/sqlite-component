## this was: Go Component SDK with Custom WIT
# it now is: Go SQLite component 

## (new text)

This started out as the example
["Go Component SDK with Custom WIT"](https://github.com/wasmCloud/go/tree/main/examples/component/invoke), <br/>
"a WebAssembly component that demonstrates how to use the wasmCloud <br/> 
[Go Component SDK](https://github.com/wasmCloud/go/tree/main/component)
in conjunction with your own custom
[WIT interfaces](https://wasmcloud.com/docs/concepts/interfaces)."

It has been heavily hacked to create multiple callable functions
and to <br/> stop naming every damned thing "invoke" or "invoker".

**But** there won't be any serious SQLite action in this repo until a prototype <br/>
of `wasi:filesystem` is made available by Bytecode Alliance / WasmCloud.

The license ia Apache 2.0, inherited from the WasmCloud code. 

## (old text)

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

We build and deploy this example manually, since we will be using the
[`wash call` <br/> subcommand](https://wasmcloud.com/docs/cli/wash#wash-call)
to interact with the application, requiring a stable identity to call. <br/>
(The alternative method is the 
[`wash dev` subcommand](https://wasmcloud.com/docs/cli/wash#wash-dev).)

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

Check this until the application has status `Deployed`:

```shell
wash app list
```

### Invoke the component

Once the application is deployed, you can call the component <br/>
using [`wash call`](https://wasmcloud.com/docs/cli/wash#wash-call),
which invokes a function on a component:

```shell
wash call invoke_example-invoker example:invoker/invoker.call
```
```text
Hello from the invoker!
```

### Clean up

You can delete an application from your wasmCloud environment by
referring either <br/> to its application name (`invoke-example`)
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
ID for the component. <br/> If you have `jq` installed,
run the following command to call the component:

```bash
wash call "$(wash get inventory -o json | jq -r \
'.inventories[0].components[0].id')" example:invoker/invoker.call
Hello from the invoker!
```

## Further reading

For more on custom interfaces, see the
[Interface Developer Guide](https://wasmcloud.com/docs/developer/interfaces/creating-an-interface)
in the wasmCloud documentation. 
