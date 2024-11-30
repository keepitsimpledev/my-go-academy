# my-go-academy
bjss.learnamp.com/en/learnlists/golang-academy

## environment
* ubuntu:
  * https://learn.microsoft.com/en-us/windows/wsl/install
  * https://apps.microsoft.com/detail/9PN20MSR04DW
  * check version with `$ lsb_release -a`
* go: https://go.dev/doc/install
  * check version with `$ go version`
  * `dlv` for debugging: `$ go install -v github.com/go-delve/delve/cmd/dlv@latest`
* golangci-lint: (https://golangci-lint.run/usage/install/, https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go)
  * run locally with default enabled libraries: `$ golangci-lint run`
  * run locally with multiple disabled libraries `$ golangci-lint run -E <library name> -E <library name> ...`
    * example: `$ golangci-lint run -E gofmt -E revive -E wsl`
```
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.56.2
$ export PATH=$PATH:/home/kenny/go-bin
$ source $HOME/.profile
```
* vscode: https://code.visualstudio.com/Download
  * go extension: https://marketplace.visualstudio.com/items?itemName=golang.Go
  * github actions (pipeline) extension: https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-github-actions
* godoc (https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world):
  ```
  $ sudo apt install net-tools # for ifconfig to find IP address https://linuxiac.com/how-to-find-ip-address-in-linux/
  $ go install golang.org/x/tools/cmd/godoc@latest
  $ export PATH=$PATH:/home/kenny/go/bin
  $ source $HOME/.profile
  ```
  * run with `godoc -http=:6060`

### some useful things

* reference script template (useful when installing):
  ```
  $ export PATH=$PATH:<path to add>
  $ source $HOME/.profile # so it's immedately referrable
  ```
* `tree` for easy file/directory structure visualization
  ```
  $ sudo apt install tree # install
  $ tree --dirsfirst # use
  ```
* use `$ go fmt ./...` to format `go` files in project (doing so should appease the `gofmt` CI check)
