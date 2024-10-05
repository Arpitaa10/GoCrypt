# GoCrypt

GoCrypt is a powerful and user-friendly command-line tool built in Go that allows you to easily encrypt and decrypt files using a password. It utilizes the SHA-1 algorithm to ensure the security of your sensitive data, making it an ideal solution for your day-to-day encryption needs.

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/AmulyaParitosh/GoCrypt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
  - [Pre-built Binaries](#pre-built-binaries)
  - [Building from Source](#building-from-source)
- [Usage](#usage)
- [Examples](#examples)
- [Security Considerations](#security-considerations)
- [Contributing](#contributing)
- [License](#license)

## Features

- Simple and intuitive command-line interface
- File encryption using SHA-1 algorithm
- Password-based encryption and decryption
- Cross-platform compatibility (Windows and Ubuntu)
- Lightweight and fast execution

## Installation

### Pre-built Binaries

For quick installation, you can download the pre-built binaries for your operating system:

#### Ubuntu
```bash
wget https://amulyaparitosh.github.io/GoCrypt/build/ubuntu/gocrypt
chmod +x gocrypt
sudo mv gocrypt /usr/local/bin/
```

#### Windows
1. Download the executable from [this link](https://amulyaparitosh.github.io/GoCrypt/build/windows/gocrypt.exe).
2. Add the directory containing `gocrypt.exe` to your system's PATH environment variable.

### Building from Source

If you prefer to build the tool from source or want to contribute to its development, follow these steps:

1. Ensure you have Go installed on your system. If not, download and install it from [the official Go website](https://golang.org/dl/).

2. Clone the repository:
   ```bash
   git clone https://github.com/AmulyaParitosh/GoCrypt.git
   ```

3. Navigate to the project directory:
   ```bash
   cd GoCrypt
   ```

4. Build the project:
   ```bash
   go build -o gocrypt .
   ```

   This will create an executable named `gocrypt` in your current directory.

5. (Optional) Move the executable to a directory in your PATH for easy access:
   ```bash
   sudo mv gocrypt /usr/local/bin/
   ```

## Usage

GoCrypt is designed to be simple and straightforward to use. Here's the general syntax:

```
gocrypt <command> [arguments]
```

### Commands

- `encrypt`: Encrypts a file given a password
- `decrypt`: Tries to decrypt a file using a password
- `help`: Displays help text

### Basic Usage

To encrypt a file:
```bash
gocrypt encrypt /path/to/your/file
```

To decrypt a file:
```bash
gocrypt decrypt /path/to/your/encrypted/file
```

To display help:
```bash
gocrypt help
```

## Examples

1. Encrypting a file:
   ```bash
   gocrypt encrypt /home/user/documents/secret.txt
   ```
   You will be prompted to enter a password for encryption.

2. Decrypting a file:
   ```bash
   gocrypt decrypt /home/user/documents/secret.txt.encrypted
   ```
   You will be prompted to enter the password used for encryption.

3. Displaying help information:
   ```bash
   gocrypt help
   ```

## Security Considerations

- GoCrypt uses the SHA-1 algorithm for encryption. While SHA-1 is considered cryptographically broken for digital signatures, it can still provide a reasonable level of security for file encryption in many scenarios.
- Always use strong, unique passwords for each file you encrypt.
- Remember that the security of your encrypted files depends on the strength of your password and how well you protect it.
- For highly sensitive data, consider using more advanced encryption tools or consult with a security expert.

## Contributing

Contributions to GoCrypt are welcome! If you have suggestions for improvements or bug fixes, please feel free to:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
