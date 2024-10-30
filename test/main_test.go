package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// create test values
var testMap = [][]string{
	{"BOOK", "A0", "1"},
	{"CANCEL", "A0", "1"},
	{"BOOK", "A0", "1"},
	{"BOOK", "A0", "1"},
	{"BOOK", "A2", "4"},
	{"BOOK", "A5", "1"},
	{"BOOK", "A6", "3"},
	{"BOOK", "A8", "1"},
	{"BOOK", "U1", "1"},
	{"CANCEL", "A5", "1"},
	{"BOOK", "B2", "0"},
	{"BOOK", "B7"},
	{"BOOK", "C0", "6"},
}

func TestMain(t *testing.T) {
	// test Main app with test values
	// clean up booking file created after testing
	for _, v := range testMap {
		if len(v) < 3 {
			v = append(v, "0")
		}
		cmd := exec.Command("go", "run", "../.", v[0], v[1], v[2])
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("Error: %s", err)
		}
		if !strings.Contains(string(out), "SUCCESS") && !strings.Contains(string(out), "FAIL") {
			fmt.Println(string(out))
			t.Fatalf("Test Failed: Parameters: %s %s %s \nExpected SUCCESS: Returned: %v", v[0], v[1], v[2], out)
		}
	}

	err := os.Remove("booking.csv")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
}

// func TestHelpers(t *testing.T) {
// 	fmt.Println("test canceling")
// }
