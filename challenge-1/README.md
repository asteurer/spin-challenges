# Overview 
Welcome to the first challenge of Summer of Spin! In this challenge, you will demonstrate the following skills (see [writing applications](https://developer.fermyon.com/spin/v2/writing-apps) and [structuring applications](https://developer.fermyon.com/spin/v2/spin-application-structure)):
- HTTP calls
- Building Spin apps with pre-build components
- Reading from a file (optional)

# Building the application

### Requirements
- Most recent version of [Spin](https://developer.fermyon.com/spin/v2/install)
- See language-specific requirements
    - [Rust](https://developer.fermyon.com/spin/v2/rust-components)
    - [Golang](https://developer.fermyon.com/spin/v2/go-components)
    - [Python](https://developer.fermyon.com/spin/v2/python-components)
    - [JavaScript](https://developer.fermyon.com/spin/v2/javascript-components)
    - [Other Languages](https://developer.fermyon.com/spin/v2/other-languages)

### Instructions

You will be building a Spin app (using `encryption-module/main.wasm` as a component) that will do the following:
- Read `message.txt` and decrypt the string using the encryption module.
- Append your name to the end of the decrypted string.
- Re-encrypt the newly-created string using the encryption module.
- Return an HTTP response with the following values:
    - Response code: 200
    - Headers:
        - content-type: application/json
        - x-name: The name value you appended to the decrypted string
        - x-encryption-module-path: The HTTP trigger path you defined for the encryption module in the `spin.toml` file
    - Body: A utf-8-encoded JSON object with `"encryptedMessage"` as the key, and the re-encrypted string as the value.

# Testing the application

### Requirements
- Latest version of [Hurl](https://hurl.dev/) installed

### Instructions

- Run your Spin application (see [Running Spin Applications](https://developer.fermyon.com/spin/v2/running-apps)). 
- In your terminal, navigate to the directory containing the `test.hurl` file. 
- Run the command `hurl --test test.hurl`
- If your application fails the tests, try using the `--verbose` or `--very-verbose` flags in the `hurl` command to debug.

# Helpful hints

Some things to keep in mind:
- There is a bug in the TinyGo compiler that prevents Spin Golang applications from reading files. If you want to use Golang, you'll need to pass in the `message.txt` file string some other way (i.e. via the body of a curl request or hard-coding).
- When you create a Spin application with multiple components, the way they will interact with each other is through HTTP calls (i.e. localhost:3000/your-component-http-trigger-route). In order to do this, you must define an `allowed_outbound_host` in your `spin.toml` file as `http://localhost:3000`.
- If you get stuck, reach out via our [Discord channel](#TODO:_LINK_HERE), or see if [ChatGPT](https://chatgpt.com/) or [Google Gemini](https://gemini.google.com/) can help.