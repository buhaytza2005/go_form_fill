package main

import (
    "fmt"
    "os"

    "github.com/benoitkugler/pdf/reader"
)

func main() {
    const path = "g02.pdf"
    doc, _, err := reader.ParsePDFFile(path, reader.Options{})
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to parse PDF file: %s\n", err)
        os.Exit(1)
    }

    // Check if the PDF has an AcroForm
        fields := doc.Catalog.AcroForm.Flatten() // This will give you a flat map of fields
        for name, _ := range fields {
            fmt.Printf("Field name: %s\n", name)
            // You can also access field properties, such as field type (FT), value (V), etc.
            // Example: field.Field.FT, field.Field.V
        }
}

