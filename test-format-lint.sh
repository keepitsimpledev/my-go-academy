#!/bin/bash

# for local dev and QOL
# this mirrors and should be kept up-to-date with test-and-lint.yml
echo '    <build> -------------------------------------------------------------------'
go build -v ./...
echo '    </build> ------------------------------------------------------------------'
echo

echo '    <test> --------------------------------------------------------------------'
go test -race ./...
echo '    </test> -------------------------------------------------------------------'
echo

echo '    <fmt> ---------------------------------------------------------------------'
go fmt ./...
echo '    </fmt> --------------------------------------------------------------------'
echo

echo '    <vet> ---------------------------------------------------------------------'
go vet ./...
echo '    </vet> --------------------------------------------------------------------'
echo
    
echo '    <golangci-lint> -----------------------------------------------------------'
golangci-lint run -E asciicheck \
-E bidichk \
-E bodyclose \
-E decorder \
-E dogsled \
-E dupl \
-E dupword \
-E errname \
-E exhaustive \
-E exhaustruct \
-E gochecknoglobals \
-E gochecknoinits \
-E goconst \
-E godox \
-E gofmt \
-E gomnd \
-E gosec \
-E lll \
-E misspell \
-E nestif \
-E nilerr \
-E reassign \
-E revive \
-E stylecheck \
-E thelper \
-E unconvert \
-E unparam \
-E wastedassign \
-E whitespace \
-E wrapcheck \
-E wsl
echo '    </golangci-lint> ----------------------------------------------------------'
echo