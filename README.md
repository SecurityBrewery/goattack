# Mitre ATT&CK in Go

A Go module containing the Mitre ATT&CK framework tactics and techniques.

## Usage

```go
package main

import (
	"fmt"

	"github.com/SecurityBrewery/goattack"
)

func main() {
	if tactic, ok := goattack.Objects["TA0004"]; ok {
		fmt.Println(tactic.Name) // Privilege Escalation
		fmt.Println(tactic.URL)  // https://attack.mitre.org/tactics/TA0004
	}

	if technique, ok := goattack.Objects["T1548.002"]; ok {
		fmt.Println(technique.Name)     // Bypass User Account Control
		fmt.Println(technique.FullName) // Abuse Elevation Control Mechanism: Bypass User Account Control
		fmt.Println(technique.URL)      // https://attack.mitre.org/techniques/T1548/002
	}
}
```

## Generation

To generate the code copy the JSON files from
[github.com/mitre/cti/…/enterprise-attack.json](https://github.com/mitre/cti/blob/master/enterprise-attack/enterprise-attack.json)
to the `gen` directory and run `make generate`.
