# Change Log

This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

### Added
- Colored results can be enabled via a `--color` flag.
  - The request status and HTTP status code are colored depending on whether they were successful or not.


## [1.2.0] - 2016-01-22

### Added
- Content of route files can be repeated endlessly via the `--endless` flag.
- Added usage help for the executable via `--help`.

## [1.1.0] - 2015-12-18

### Added
- CSV output can be written to a file directly instead of stdout.

### Fixed
- Trying to read an unknown or inaccessible file yields an error message.
- Main executable returns proper error codes on error conditions (e.g. unknown file).


## [1.0.0] - 2015-11-25

### Changed
- Command line options are flags now instead of positional arguments.


## 0.0.1 - 2015-11-16

### Added
- Initial release.
- Reading a text file with routes.
- Every line of the file will be a request executed in a separate goroutine.
- Results are consolidated as CSV content printed to stdout.
- The target host and number of separate threads are configurable.


[Unreleased]: https://github.com/christophgockel/goony/compare/1.2.0...HEAD
[1.2.0]: https://github.com/christophgockel/goony/compare/1.1.0...1.2.0
[1.1.0]: https://github.com/christophgockel/goony/compare/1.0.0...1.1.0
[1.0.0]: https://github.com/christophgockel/goony/compare/0.0.1...1.0.0

