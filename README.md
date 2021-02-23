# td [![Go Reference](https://pkg.go.dev/badge/github.com/gotd/td.svg)](https://pkg.go.dev/github.com/gotd/td/telegram) [![codecov](https://codecov.io/gh/gotd/td/branch/main/graph/badge.svg?token=shR5PXG7Ds)](https://codecov.io/gh/gotd/td)

Telegram client and MTProto implementation in go.

* [Security policy](.github/SECURITY.md)
* [User support and dev chat](.github/SUPPORT.md)
* [Roadmap](ROADMAP.md)
* [Contributing](CONTRIBUTING.md)
* [Architecture](ARCHITECTURE.md)

Before using this library, read [How Not To Get Banned](https://github.com/gotd/td/blob/main/.github/SUPPORT.md#how-not-to-get-banned) guide.

## Status

Work in progress. Only go1.16 is supported and no backward compatibility is provided.

Goal of this project is to implement Telegram client while
providing building blocks for the other client or even server
implementation without performance bottlenecks.

This project is fully non-commercial and not affiliated with any commercial organization
(including Telegram LLC).

## Features

* Full MTProto 2.0 implementation in go
* Code for Telegram types generated by `./cmd/gotdgen` (based on [gotd/tl](https://github.com/gotd/tl) parser)
* Embedded [official documentation](https://core.telegram.org/schema)
* Pluggable session storage
* Automatic re-connects with keepalive
* Vendored Telegram public keys that are kept up-to-date
* End-to-end and unit testing, fuzzing
* No runtime reflection overhead
* Secure PRNG used for crypto
* 2FA support
* MTProxy support
* Connection pooling
* Helpers for uploads and downloads
* CDN support
* Automatic DC migration and redirect handling

## Example

You can see `cmd/gotdecho` for simple echo bot example or [gotd/bot](https://github.com/gotd/bot) that can
recover from restarts and fetch missed updates (and also is used as canary for stability testing).

### Auth

#### User

You can use `td/telegram/AuthFlow` to simplify user authentication flow.

```go
codePrompt := func(ctx context.Context) (string, error) {
    // NB: Use "golang.org/x/crypto/ssh/terminal" to prompt password.
    fmt.Print("Enter code: ")
    code, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(code), nil
}
// This will setup and perform authentication flow.
// If account does not require 2FA password, use telegram.CodeOnlyAuth
// instead of telegram.ConstantAuth.
if err := telegram.NewAuth(
    telegram.ConstantAuth(phone, password, telegram.CodeAuthenticatorFunc(codePrompt)),
    telegram.SendCodeOptions{},
).Run(ctx, client); err != nil {
    panic(err)
}
```
#### Bot

Use bot token from [@BotFather](https://telegram.me/BotFather).

```go
if err := client.AuthBot(ctx, "token:12345"); err != nil {
    panic(err)
}
```

### Calling MTProto directly

You can use generated `tg.Client` that allows calling any MTProto methods
directly.

```go
// Grab these from https://my.telegram.org/apps.
// Never share it or hardcode!
client := telegram.NewClient(appID, appHash, telegram.Options{})
client.Run(ctx, func(ctx context.Context) error) {
  // Grab token from @BotFather.
  if err := client.AuthBot(ctx, "token:12345"); err != nil {
    return err
  }
  state, err := tg.NewClient(client).UpdatesGetState(ctx)
  if err != nil {
    return err
  }
  // Got state: &{Pts:197 Qts:0 Date:1606855030 Seq:1 UnreadCount:106}
  // This will close client and cleanup resources.
  return nil
}
```

### Generated code

Code output of `gotdgen` contains references to TL types, examples, URL to
official documentation and [extracted](https://github.com/gotd/getdoc) comments from it.

For example, the [auth.Authorization](https://core.telegram.org/type/auth.Authorization) type in `tg/tl_auth_authorization_gen.go`:

```go
// AuthAuthorizationClass represents auth.Authorization generic type.
//
// See https://core.telegram.org/type/auth.Authorization for reference.
//
// Example:
//  g, err := DecodeAuthAuthorization(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *AuthAuthorization: // auth.authorization#cd050916
//  case *AuthAuthorizationSignUpRequired: // auth.authorizationSignUpRequired#44747e9a
//  default: panic(v)
//  }
type AuthAuthorizationClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthAuthorizationClass
}
```
Also, the corresponding [auth.signIn](https://core.telegram.org/method/auth.signIn) method:
```go
// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {}
```

The generated constructors contain detailed official documentation, including links:
```go
// FoldersDeleteFolderRequest represents TL type `folders.deleteFolder#1c295881`.
// Delete a peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/method/folders.deleteFolder for reference.
type FoldersDeleteFolderRequest struct {
    // Peer folder ID, for more info click here¹
    //
    // Links:
    //  1) https://core.telegram.org/api/folders#peer-folders
    FolderID int
}
```

## Contributions

Huge thanks to every contributor, dealing with a project of such scale is impossible alone.

Special thanks:
* [tdakkota](https://github.com/tdakkota)
    * Two-factor authentication (SRP)
    * Proxy support
    * Update dispatcher
    * Complete transport support (abridged, padded intermediate and full)
    * Telegram server for end-to-end testing
    * Multiple major refactorings, including critical cryptographical scope reduction
    * Code generation improvements (vector support, multiple modes for pretty-print)
    * And many other cool things and performance improvements
* [zweihander](https://github.com/zweihander)
    * Background pings
    * Links in generated documentation
    * Message acknowledgements
    * Retries
    * RPC Engine

## Reference

The MTProto protocol description is [hosted](https://core.telegram.org/mtproto#general-description) by Telegram.

Most important parts for client implementations:
* [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for client software developers

Current implementation [mostly conforms](https://github.com/gotd/td/issues/155) to security guidelines, but no
formal security audit were performed.

## Prior art

* [Lonami/grammers](https://github.com/Lonami/grammers) (Great Telegram client in Rust, many test vectors were used as reference)
* [sdidyk/mtproto](https://github.com/sdidyk/mtproto), [cjongseok/mtproto](https://github.com/cjongseok/mtproto), [xelaj/mtproto](https://github.com/xelaj/mtproto)  (MTProto 1.0 in go)

## License
MIT License

Created by Aleksandr (ernado) Razumov

2020
