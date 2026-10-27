
## Usage

Start the mailpit container.

```sh
docker run --name mailpit --rm -p 8025:8025 -p 1025:1025 axllent/mailpit 
```

Add *-e MP_VERBOSE=1* for more verbose logging.

Build the mail client.

```sh
go build .
```

Run the client.

```
./mailclient
```