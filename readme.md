# Ender

A Great Tool to Gather SQL ENDPOINTS Based on The Error Reported written in GO.

## Installation

Clone the repository, chmod u+x ender.go, run the script.

```bash
git clone https://github.com/DecodedTrojans/Ender.git && cd Ender
```

## Usage

```sh
$ go run ender.go target.com
```

## Output

```bash
[+]Checking the Internet Connection ...
Internet Check Done [✓]
[+]Checking the Domain.
Domain Check Done[✓]
[+]Gathering The End-Points for target.com
[+]EndPoints Gathered! Now Analyzing target.txt
[+]Searching The SQL parameters From The Input File target.txt.
[-]Found 10 SQL Errors in the target.com
```
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
