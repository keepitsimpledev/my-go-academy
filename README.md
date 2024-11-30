# my-go-academy
bjss.learnamp.com/en/learnlists/golang-academy

## environment
* ubuntu:
  * https://learn.microsoft.com/en-us/windows/wsl/install
  * https://apps.microsoft.com/detail/9PN20MSR04DW
* go: https://go.dev/doc/install
* golangci-lint: (https://golangci-lint.run/usage/install/, https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go)
```
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.56.2
$ export PATH=$PATH:/home/kenny/go-bin
$ source $HOME/.profile
```
* vscode: https://code.visualstudio.com/Download
  * vscode go extension: github.com/golang/vscode-go
* godoc (https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world):
  ```
  $ sudo apt install net-tools # for ifconfig to find IP address https://linuxiac.com/how-to-find-ip-address-in-linux/
  $ go install golang.org/x/tools/cmd/godoc@latest
  $ export PATH=$PATH:/home/kenny/go/bin
  $ source $HOME/.profile
  ```

reference script template (useful when installing):
```
$ export PATH=$PATH:<path to add>
$ source $HOME/.profile # so it's immedately referrable
```

