# Overview 
Welcome to the first challenge of Summer of Spin! In this challenge, you will demonstrate the following skills:
- HTTP calls
- Building Spin apps with pre-build components
- Reading from a file (optional)

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

Some things to keep in mind:
- There is a bug in the TinyGo compiler that prevents Spin Golang applications from reading files. If you wish to do this challenge in Go, you'll need to hard-code the values found in `message.txt`
- When you create a Spin application with multiple components, the way they will interact with each other is through HTTP calls (i.e. localhost:3000/your-component-http-trigger-route). In order to do this, you must define an `allowed_outbound_host` in your `spin.toml` file as `http://localhost:3000`.
- 
# 