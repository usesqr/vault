# Welcome to the Sqr Vault Source Repository

Sqr Vault is an on-device, fully encrypted password manager designed to keep your sensitive logins secure.

Part of the Sqr series, Vault embodies our commitment to providing reliable, user-focused apps with an emphasis on privacy.

## Getting Started on Your Computer

### Prebuilt Binaries

Download prebuilt binaries from the "Releases" section of this repository.

### Installation Guide

- **Windows**: Download and run the `.exe` file.
- **macOS**: Download the `.app` file and move it to your Applications folder.
- **GNU/Linux**: Download the `.tar.xz` or `.tar.gz` file and extract it to your root directory.

### Building Your Own Binary

To build your own binary, follow the [Fyne setup guide](https://docs.fyne.io/started/) to prepare your development environment.

Clone this repository using the green "Code" button above or run:
`git clone https://github.com/usesqr/vault.git`

You can then:

1. Run Vault directly with `go run .`
2. Build a binary for your OS by following [these packaging steps](https://docs.fyne.io/started/packaging.html).

After building, refer to the "Installation Guide" above to install your binary.

## Features

- Vault features an intuitive, modern graphical interface.
- Your data remains on your device at all times.
- The password database is secured with AES-256 encryption.
- A securely generated salt is used with PBKDF2 to protect against rainbow-table attacks.

## Technical Details

- Developed in Go using the [Fyne](https://github.com/fyne-io/fyne) GUI library.
- Manages the SQLite database with [GORM](https://gorm.io).
- Encryption is handled by [gorm-sqlite-cipher](https://github.com/jackfr0st13/gorm-sqlite-cipher) using AES-256.
- The Go crypto library provides PBKDF2 for key derivation.

## License

All content in this repository is licensed under the MIT license unless otherwise noted. See the [LICENSE](LICENSE) file for more information.

Please note that third-party libraries included in the software may be licensed
under different terms.
