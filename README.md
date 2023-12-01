# Polarbear

Simple CLI offering HR and ECG streaming in real-time from Polar H10.

## Installation

`go install github.com/siiimooon/polarbear@latest`

## Usage
``` bash
$ polarbear scan [-d <duration time.Duration>]
device name:some-uuid-one       alice-polar-h10
device name:some-uuid-two       bob-polar-h10
device name:some-uuid-three     eve-polar-h10

$ polarbear ecg stream -u some-uuid-one
<number of directional changes within an R-R interval> - [array of each segment, representing directional changes]

$ polarbear hr -u some-uuid-one
Heart rate: ..., RR interval(s): [...]
```

## Disclaimer

**DISCLAIMER:** This code is not intended for production use. Please be aware that using this code in a production environment may result in unexpected outcomes, security vulnerabilities, or errors. Use at your own risk.

