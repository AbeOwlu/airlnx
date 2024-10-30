## Airline Seat Booking App
---
Application manages seat booking or cancelling on an airplane.

### Usage
Build application for specific architecture with;
```
export GOOS=<linux or window ...> # https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies
go build
```

A linux compatible, `airlnx` is provided in this repo. The application ca be exucuted passing required arguments to the executable, e.g;
```
./airlnx BOOK A1 1
```

If less arguments than required are passed, a helpful usage propmt is printed:
```
try: <filename> BOOK A0 1
```

### Tests and Enhancement
Program is tested using the `testMap` variable, covering basic expected usecases, and unexpected cases. Edge cases not exhaustively covered by test sample. Additionl test cases can be added to the testMap, and run with;
```
go test ./test -v
```
- Specific test cases for the handler functions can be introduced, for testing added features directly
- discovering and testing edge cases such as attempted (malicious or not) multi book or cancel - mutex on map and file store updates

---
- EnEnhancements can be added as features to allow cancelling consecutive seats that are booked for example.
- Seats that are booked together can also get a field indicating they are together. This consideration informed using ubiquitous csv file -instead of json for storing seat status. Additional fields can be added without complications.
- Program is unoptimized - profiling of performance and tuning would enhance program
- Using cobra package for interactive program execution, allowing mixed-case arguments for user ease-of-use, etc.