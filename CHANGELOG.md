# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.5] - 2022-07-28
### Changed
- Add tracking fields https://github.com/AfterShip/aftership-sdk-go/pull/56

## [2.0.4] - 2022-07-20
### Added
- Add predict-batch endpoint and update tracking fields https://github.com/AfterShip/aftership-sdk-go/pull/42
### Changed
- Update tracking fields https://github.com/AfterShip/aftership-sdk-go/pull/41

## [2.0.3] - 2022-05-20
### Changed
- update tracking fields

## [2.0.2] - 2022-01-13
### Fixed
- Fixing https://github.com/AfterShip/aftership-sdk-go/issues/34

## [2.0.1] - 2022-05-06
### Changed
- Change shipment_weight type to float64

## [2.0.0] - 2022-01-13
### Breaking Changes
- Completely reorganized the SDK, see [Migrations](https://github.com/AfterShip/aftership-sdk-go#migrations)
- Removed `auto retry` feature, consumers need to retry the request by themselves.

Compatibility
- Go >= 1.13
### Added
- Support latest features in v4 API
- Use `go mod` to organize the Go modules
- Error handling

[2.0.5]: https://github.com/AfterShip/aftership-sdk-go/compare/2.0.4...2.0.5
[2.0.4]: https://github.com/AfterShip/aftership-sdk-go/compare/v2.0.3...2.0.4
[2.0.3]: https://github.com/AfterShip/aftership-sdk-go/compare/v2.0.2...v2.0.3
[2.0.2]: https://github.com/AfterShip/aftership-sdk-go/compare/v2.0.1...v2.0.2
[2.0.1]: https://github.com/AfterShip/aftership-sdk-go/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/AfterShip/aftership-sdk-go/releases/tag/v2.0.0