## Release History

## Change log

### 2.0.0-alpha (2020-04-30)
- New features
	- Support latest features in v4 API
	- Use `go mod` to organize the Go modules
	- Error handling
- Compatibility
	- Go >= 1.13
- Breaking changes
	- Completely reorganized the SDK, see [Migrations](https://github.com/AfterShip/aftership-sdk-go#migrations)
	- Removed `auto retry` feature, consumers need to retry the request by themselves.
