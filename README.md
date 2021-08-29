# Secret - MiniVault to store secret key-value pairs
> Note: For Learning Purpose.

> All credits to [Gophercises](https://gophercises.com/).

Secret is a package that can be used to store/retrieve secrets to/from a file. It encrypts the file contents using the key that you provide.

Secret uses AES for encryption and `crypto/cipher` package's cipher.Stream to create DecryptReader and EncryptWriter which can decrypt while reading from a file and encrypt while writing to a file respectively.

---

The repo also contains a command line tool (in `cmd/`) written using the `secret` package.
## Usage
```
secret [command]

Available Commands:
  get         Gets a secret from your secret storage
  set         Sets a secret in your secret storage.
  list        Lists all keys stored

Flags:
  -f, --file string   the path to file where secrets are stored
```
If `-f` is not provided, it stores secrets to `$HOME/.secrets` on Unix and `%USERPROFILE%\\.secrets` on Windows.

## Example:
```
$> secret set github_api_key 0imfnc8mVLWwsAawjYr4Rx-Af50DDqtlx
encoding key :
Value set successfully!

$> secret get github_api_key
encoding key :
0imfnc8mVLWwsAawjYr4Rx-Af50DDqtlx
```
## Installation / Tryout
```sh
git clone https://github.com/opxyc/secret && cd secret
go install cmd/cli.go
```
Now, ideally, it should generate a new executable `cli` in $GOPATH/bin. It can be renamed accordingly (Go does not currently provide an option to give a custom name for the binary generated during installation - as of 1.7).