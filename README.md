# elrsbrute
bruteforce bind phrase from STDIN

go build elrsbrute.go

designed to work with pipes:
cat wordlist | ./elrsbrute -uid "79, 4, 253, 130, 33, 85"
./hashcat.bin -a3 "?l?l?l?l" -i --stdout | ./elrsbrute -uid "79, 4, 253, 130, 33, 85"
