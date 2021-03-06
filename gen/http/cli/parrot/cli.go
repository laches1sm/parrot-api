// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// parrot HTTP client CLI support package
//
// Command:
// $ goa gen github.com/bdna/the_parrot_api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	parrotsvcc "github.com/bdna/the_parrot_api/gen/http/parrot/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `parrot add-parrot
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` parrot add-parrot --name "Dolor aut." --breed "Officiis aperiam ut voluptatem voluptate." --colour "Tempora dolores possimus."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		parrotFlags = flag.NewFlagSet("parrot", flag.ContinueOnError)

		parrotAddParrotFlags      = flag.NewFlagSet("add-parrot", flag.ExitOnError)
		parrotAddParrotNameFlag   = parrotAddParrotFlags.String("name", "REQUIRED", "Parrot name")
		parrotAddParrotBreedFlag  = parrotAddParrotFlags.String("breed", "REQUIRED", "Parrot breed")
		parrotAddParrotColourFlag = parrotAddParrotFlags.String("colour", "REQUIRED", "Parrot colour")
	)
	parrotFlags.Usage = parrotUsage
	parrotAddParrotFlags.Usage = parrotAddParrotUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "parrot":
			svcf = parrotFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "parrot":
			switch epn {
			case "add-parrot":
				epf = parrotAddParrotFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "parrot":
			c := parrotsvcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add-parrot":
				endpoint = c.AddParrot()
				data, err = parrotsvcc.BuildAddParrotPayload(*parrotAddParrotNameFlag, *parrotAddParrotBreedFlag, *parrotAddParrotColourFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// parrotUsage displays the usage of the parrot command and its subcommands.
func parrotUsage() {
	fmt.Fprintf(os.Stderr, `goa is trash
Usage:
    %s [globalflags] parrot COMMAND [flags]

COMMAND:
    add-parrot: AddParrot implements add-parrot.

Additional help:
    %s parrot COMMAND --help
`, os.Args[0], os.Args[0])
}
func parrotAddParrotUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] parrot add-parrot -name STRING -breed STRING -colour STRING

AddParrot implements add-parrot.
    -name STRING: Parrot name
    -breed STRING: Parrot breed
    -colour STRING: Parrot colour

Example:
    `+os.Args[0]+` parrot add-parrot --name "Dolor aut." --breed "Officiis aperiam ut voluptatem voluptate." --colour "Tempora dolores possimus."
`, os.Args[0])
}
