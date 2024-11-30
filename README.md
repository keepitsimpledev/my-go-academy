# my-go-academy
The Go Academy was a private pop-up academy.

It was an introduction to Go covering the programming language's syntax, concepts and everyday use.

The academy leveraged several resource types to support a learning journey and a list of assignments to work through. One of the assignments was to build a "To Do Application" with a REST API.

Resources:
* [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests)
  * this repository primarily consists of "Learn Go With Tests" exercises
* [Go by example](https://gobyexample.com/)
* [Learning Go, 2nd Edition](https://www.oreilly.com/library/view/learning-go-2nd/9781098139285/)

## environment (steps for Windows)
* wsl ([reference](https://learn.microsoft.com/en-us/windows/wsl/install-manual))
  * in "Windows Features", ensure "Virtual Machine Platform" and "Windows Subsystem for Linux" are enabled
  * in powershell, set version 2 as default with `$ wsl --set-default-version 2`
  * [update linux kernel](https://learn.microsoft.com/en-us/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package)
* ubuntu:
  * https://apps.microsoft.com/detail/9PN20MSR04DW
    * if installation hangs, [try sending SIGINT](https://github.com/microsoft/WSL/issues/6405)
  * check version with `$ lsb_release -a`
  * to re-install (and destroy/recreate environments):
    1. in Windows Powershell: `PS> wsl --unregister Ubuntu-22.04`
    1. in Add/Remove Programs, uninstall Ubuntu LTS
    1. Get Ubuntu LTS from https://apps.microsoft.com/detail/9PN20MSR04DW again
* git
  * don't forget to configure
    ```
    $ git config --global user.name "<yourname>"
    $ git config --global user.email <github email>
    ```
* go: https://go.dev/doc/install
  * check version with `$ go version`
  * `dlv` for debugging: `$ go install -v github.com/go-delve/delve/cmd/dlv@latest`
    * this seems to install to `~/go/bin`
* golangci-lint: (https://golangci-lint.run/welcome/install/, https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go)
  ```
  ## Install Steps
  # this installs in the relative directory `./bin/golangci-lint` - we'll install it in `~/go`:
  $ pushd ~/go
  $ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.56.2
  $ popd
  ```
  * run locally with default enabled libraries: `$ golangci-lint run`
  * run locally with multiple disabled libraries `$ golangci-lint run -E <library name> -E <library name> ...`
    * example: `$ golangci-lint run -E gofmt -E revive -E wsl`
* vscode: https://code.visualstudio.com/Download
  * be sure to enable `File > Auto Save`
  * wsl extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl
  * go extension: https://marketplace.visualstudio.com/items?itemName=golang.Go
    * I was prompted to installl `gopls, and so I did. For some reason, installing `gopls` with `sudo` placed it in the inaccessible `/root/go/bin/`, so i copied it somewhere accessible (and on the `$PATH`): `sudo cp /root/go/bin/gopls /usr/local/go/bin`
  * github actions (pipeline) extension: https://marketplace.visualstudio.com/items?itemName=GitHub.vscode-github-actions
    * for context-support when managing pipeline (`.yml`) files
* godoc (https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world):
  ```
  $ sudo apt install net-tools # for ifconfig to find IP address https://linuxiac.com/how-to-find-ip-address-in-linux/
  $ go install golang.org/x/tools/cmd/godoc@latest
  ```
  * run with `godoc -http=:6060`
* `CGO_ENABLED=1` and `gcc` to enable `go test -race`
  ```
  $ go env -w CGO_ENABLED=1
  $ sudo apt-get update
  $ sudo apt-get install gcc
  ```
* add necessary PATH exports to `.bashrc` so they don't have to be exported each time the terminal is restarted.
  * example snippet from `.bashrc` with some confirmation echoes:
  ```
  $ echo go academy startup routine
  $ echo executing:
  $ echo 'export PATH=$PATH:/usr/local/go/bin' # where go is
  $ echo 'export PATH=$PATH:/home/kenny/go/bin' # where golangci-lint and godoc are
  $ echo 'echo $PATH'
  $ echo 'which go'
  $ echo 'which godoc'
  $ echo 'which golangci-lint'
  $ echo
  $ echo output:
  $ export PATH=$PATH:/usr/local/go/bin # where go is
  $ export PATH=$PATH:/home/kenny/go/bin # where golangci-lint and godoc are
  $ echo $PATH
  $ echo
  $ echo $(which go)
  $ echo $(which godoc)
  $ echo $(which golangci-lint)
  ```

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
