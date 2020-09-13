# Square CLI 
[![Build Status](https://travis-ci.com/nickrobinson/square-cli.svg?token=mYDTz49qs6zeiYaoGsHS&branch=master)](https://travis-ci.com/nickrobinson/square-cli)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/nickrobinson/square-cli)

`square` makes all [Square Connect APIs](https://developer.squareup.com/explorer) available on the command line.

## Getting Started
Run the following command to initialize the Square CLI. If you haven't done so already, you will need to head over the the [Square Developer Dashboard](https://developer.squareup.com/apps/) to create a new App, and retrieve your Access Key.

`# square init`

## Usage

- `square init` - Initialize Configuration
- `square get customers`
- `square get invoices -d location_id=L471AVFQJ8Z7J`
- `square delete customers/93Y9K6BQ8WRPV1C45GVMG9JG6M`
- `square put customers/93Y9K6BQ8WRPV1C45GVMG9JG6M -d '{"company_name": "Square"}'`
- `square post customers -d '{"given_name": "Jack", "family_name": "Dorsey"}'`
