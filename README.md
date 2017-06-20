# jsoncheck
[![Build Status](https://travis-ci.org/vodkabears/jsoncheck.svg?branch=master)](https://travis-ci.org/vodkabears/jsoncheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodkabears/jsoncheck)](https://goreportcard.com/report/github.com/vodkabears/jsoncheck)

Tool for validating a batch of local or remote JSON documents by schemas.

## Usage

Download the latest release: https://github.com/vodkabears/jsoncheck/releases/latest

Run:

```bash
./jsoncheck [configuration file]
```

## Example of a configuration file

```json
[
	{
		"schema": "schemas/regions.json",
		"data": [
			"http://127.0.0.1:3000/api/regions",
			"http://127.0.0.1:3000/api/regions/?rgid=2"
		]
	},
	{
		"schema": "schemas/offer.json",
		"data": [
			"http://127.0.0.1:3000/api/offer/?id=1",
			"file:///home/me/offer.json"
		]
	}
]
```
