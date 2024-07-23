package goattack_test

import (
	"fmt"

	"github.com/SecurityBrewery/goattack"
)

func Example() {
	if tactic, ok := goattack.Objects["TA0004"]; ok {
		fmt.Println(tactic.Name)
		fmt.Println(tactic.URL)
	}

	if technique, ok := goattack.Objects["T1548.002"]; ok {
		fmt.Println(technique.Name)
		fmt.Println(technique.FullName)
		fmt.Println(technique.URL)
	}

	// Output:
	// Privilege Escalation
	// https://attack.mitre.org/tactics/TA0004
	// Bypass User Account Control
	// Abuse Elevation Control Mechanism: Bypass User Account Control
	// https://attack.mitre.org/techniques/T1548/002
}
