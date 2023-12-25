package main

import (
	"fmt"
	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
	"os" // Import the os package
)

var data = []formfill.FDFField{
	{T: "z2", 
	Values: formfill.Values{
	    V: formfill.FDFText("testing_here"),
	},

    },
}

func main() {
	const path = "original_sample.pdf"
	doc, _, err := reader.ParsePDFFile(path, reader.Options{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	err = formfill.FillForm(&doc, formfill.FDFDict{Fields: data}, true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	out, err := os.Create("filled.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	if err = doc.Write(out, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

}
