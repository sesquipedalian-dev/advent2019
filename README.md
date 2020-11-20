# advent2019
Advent of Code 2019 - only a year too late! https://adventofcode.com/2019

Go HTTP server with endpoints for each of the challenges in advent 2019! Available API:
* `GET /days` - list of days completed
* `GET /days/{day num}/challenges` - list of challenges that day
* `GET /days/{day num}/challenges/{day num}` - challenge answer for that day.  Any challenge input should be provided in the HTTP body as text/plain

e.g. `GET /days/1/challenges/1`

further edit the readme