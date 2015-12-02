# Change Log

This project adheres to [Semantic Versioning](http://semver.org/).


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


[Unreleased]: https://github.com/christophgockel/goony/compare/1.0.0...HEAD
[1.0.0]: https://github.com/christophgockel/goony/compare/0.0.1...1.0.0

