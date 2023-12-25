package main

import (
	"fmt"
	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
	"os" // Import the os package
)

var data = []formfill.FDFField{
	{T: "z1", Values: formfill.Values{V: formfill.FDFText("879-sde9-898")}},
	{T: "z2", Values: formfill.Values{V: formfill.FDFText("ACVE")}},
	{T: "z4", Values: formfill.Values{V: formfill.FDFText("La Maison du Rocher")}},
	{T: "z5", Values: formfill.Values{V: formfill.FDFText("26160")}},
	{T: "z5b", Values: formfill.Values{V: formfill.FDFText("CHAMALOC")}},
	{T: "z6", Values: formfill.Values{V: formfill.FDFText("Créer et gérer des séjours pour enfants, adolescents et adultes.")}},
	{T: "z7", Values: formfill.Values{V: formfill.FDFText("Faire connaître, à travers des animations adaptées à l’âge des participants, les valeurs chrétiennes.")}},
	{T: "z9", Values: formfill.Values{V: formfill.FDFName("Oui")}},
	{T: "d3", Values: formfill.Values{V: formfill.FDFText("1957")}},
	{T: "d3b", Values: formfill.Values{V: formfill.FDFText("1957")}},
	{T: "d1", Values: formfill.Values{V: formfill.FDFText("5")}},
	{T: "d1b", Values: formfill.Values{V: formfill.FDFText("29")}},
	{T: "d2", Values: formfill.Values{V: formfill.FDFText("1")}},
	{T: "d2b", Values: formfill.Values{V: formfill.FDFText("1")}},
	{T: "z29", Values: formfill.Values{V: formfill.FDFText("')='à=(kmlrk'")}},
	{T: "z30", Values: formfill.Values{V: formfill.FDFText("mldmskld8+-*")}},
	{T: "z31", Values: formfill.Values{V: formfill.FDFText("lmemzkd\ndlss\nzlkdsmkmdkmsdk")}},
	{T: "z32", Values: formfill.Values{V: formfill.FDFText("kdskdl")}},
	{T: "z33", Values: formfill.Values{V: formfill.FDFText("ùmdslsùmd")}},
	{T: "z34", Values: formfill.Values{V: formfill.FDFText("1457.46")}},
	{T: "z35", Values: formfill.Values{V: formfill.FDFText("mille quatre cent cinquante-sept euros et quarante-six centimes")}},
	{T: "z36", Values: formfill.Values{V: formfill.FDFText("25")}},
	{T: "z37", Values: formfill.Values{V: formfill.FDFText("11")}},
	{T: "z38", Values: formfill.Values{V: formfill.FDFText("2020")}},
	{T: "z50", Values: formfill.Values{V: formfill.FDFName("Oui")}},
	{T: "z39", Values: formfill.Values{V: formfill.FDFName("Oui")}},
	{T: "z46", Values: formfill.Values{V: formfill.FDFName("Oui")}},
	{T: "z44", Values: formfill.Values{V: formfill.FDFName("Oui")}},
	{T: "z52", Values: formfill.Values{V: formfill.FDFText("25")}},
	{T: "z53", Values: formfill.Values{V: formfill.FDFText("11")}},
	{T: "z54", Values: formfill.Values{V: formfill.FDFText("2020")}},
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
