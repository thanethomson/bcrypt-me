# bcrypt-me

A simple CLI for generating and validating
[bcrypt](https://en.wikipedia.org/wiki/Bcrypt) passwords/hashes.

## Requirements
Needs Go 1.12+

## Building and/or Installing
To build and/or install the application, clone the repo and run the following
from the root of the repo:

```bash
# Builds the executable to ./build/bcrypt-me
make

# Installs the executable into your $GOPATH
make install
```

## Usage
Running `bcrypt-me -h` will give you:

```
A simple CLI for generating and validating bcrypt hashes

Usage:
  bcrypt-me [command]

Available Commands:
  check       Validate the given password against the specified bcrypt hash
  gen         Generate a bcrypt hash of [password] with the given parameters
  help        Help about any command

Flags:
  -h, --help   help for bcrypt-me

Use "bcrypt-me [command] --help" for more information about a command.
```

## Usage Examples
A simple example on how to use the application:

```bash
# Generate bcrypt hash of "thisismypassword" with default cost (12)
bcrypt-me gen thisismypassword
$2a$12$K.OGx67Xys3uou.cPJg8qeFq47q2ytiUo8TQXuZVB4ol376fc54a.

# Then validate the password (note the escaping of the $ signs)
bcrypt-me check \$2a\$12\$K.OGx67Xys3uou.cPJg8qeFq47q2ytiUo8TQXuZVB4ol376fc54a. thisismypassword
Password matches!

# Give it an incorrect password
bcrypt-me check \$2a\$12\$K.OGx67Xys3uou.cPJg8qeFq47q2ytiUo8TQXuZVB4ol376fc54a. thisisNOTmypassword
crypto/bcrypt: hashedPassword is not the hash of the given password

# Generate bcrypt hash of "thisismypassword" with higher cost (16)
bcrypt-me gen thisismypassword -c 16
$2a$16$hd6AA3NEDtgIdWDfGIilWe8S2VtzWKUNZcQLZOibF0LlVCF.AZsle

# Validate the hash
bcrypt-me check \$2a\$16\$hd6AA3NEDtgIdWDfGIilWe8S2VtzWKUNZcQLZOibF0LlVCF.AZsle thisisNOTmypassword
crypto/bcrypt: hashedPassword is not the hash of the given password
```
