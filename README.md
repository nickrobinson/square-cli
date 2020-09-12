# Square CLI 
[![Build Status](https://travis-ci.com/nickrobinson/square-cli.svg?token=mYDTz49qs6zeiYaoGsHS&branch=master)](https://travis-ci.com/nickrobinson/square-cli)

`square` makes all [Square Connect APIs](https://developer.squareup.com/explorer) available on the command line.

## Usage

- `square init` - Initialize Configuration
- `square get customers`
- `square get invoices -d location_id=L471AVFQJ8Z7J`
- `square delete customers/93Y9K6BQ8WRPV1C45GVMG9JG6M`
- `square put customers/93Y9K6BQ8WRPV1C45GVMG9JG6M -d '{"company_name": "Square"}'`
- `square post customers -d '{"given_name": "Jack", "family_name": "Dorsey"}'`
