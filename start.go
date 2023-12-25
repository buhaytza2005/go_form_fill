package main

import (
	"fmt"
	"os" // Import the os package

	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
)

type ApplicationData struct {
	retailer_name *string
	new_consent *string
    existing_consent *string
    consent_reference_or_DPID *string
    spid *string
    company_applicant *string
    sole_trader_applicant *string
    applicant_name *string
    registered_address_1 *string
    registered_address_2 *string
    registered_address_3 *string
    registered_address_4 *string
    registered_address_postcode *string
    organisation_is_a_company *string
    organisation_is_not_a_company *string
    company_number *string
    business *string
    site_name *string
    site_address_line_1 *string
    site_address_line_2 *string
    site_address_line_3 *string
    site_address_line_4 *string
    site_postcode *string
    fp_role *string
    landline *string
    fp_mobile *string
    fp_email *string
    fp_name *string
    fp_contact_name *string
    opening_hours *string
    same_contact_for_out_of_hours *string
    owner_is_applying *string
    landlord_name *string
    landlord_address_line_1 *string
    landlord_address_line_2 *string
    landlord_address_line_3 *string
    landlord_address_line_4 *string
    landlord_postcode *string
    landlord_number *string
    production_1 *string
    production_information_1 *string
    production_information_2 *string
    trade_effluent_treatment *string
    m3_per_day *string
    litres_per_second *string
    start_date *string
    normal_hours *string
    batch *string
    pumped *string
    sample_location_part_1 *string
    sample_location_part_2 *string
    public_connection *string
    new_or_existing_connection *string
    foul_sewer *string
    contaminated_water *string
    contaminated_water_volume *string
    ph_monitoring *string
    meter_present *string
    meter_serial *string
    average_water_use *string
    average_water_other_sources *string
    average_days_per_week *string
    average_weeks_per_year *string
    closures *string
    reasons_why_water_not_returned *string
    hse_info *string
    env_permit *string
    env_permit_application_started *string
    opening_date *string
    job_role *string
    company_name *string
    phone_no *string
    email *string
    signature *string
    coshh_mention *string 
}


func NewApplicationData() *ApplicationData {
	return &ApplicationData{
		retailer_name: StringPtr("Water plus"),
		new_consent: nil,
		existing_consent: nil,
		spid: nil,
		applicant_name: nil,
		consent_reference_or_DPID: nil,
		company_applicant: nil,
		sole_trader_applicant: nil,
		registered_address_1: nil,
		registered_address_2: nil,
		registered_address_3: nil,
		registered_address_4: nil,
		registered_address_postcode: nil,
		organisation_is_a_company: nil,
		organisation_is_not_a_company: nil,
		company_number: nil,
		business: StringPtr("Hand car wash"),
		site_name: nil,
		site_address_line_1: nil,
		site_address_line_2: nil,
		site_address_line_3: nil,
		site_address_line_4: nil,
		site_postcode: nil,
		fp_role: StringPtr("Franchisee / Site manager"),
		landline: StringPtr("N/A"),
		fp_mobile: nil,
		fp_email: nil,
		fp_name: nil,
		fp_contact_name: nil,
		opening_hours: StringPtr("Mon-Sat 9:00-17:00 Sun 10:00-16:00"),
		same_contact_for_out_of_hours: StringPtr("Yes"),
		owner_is_applying: StringPtr("No"),
		landlord_name: nil,
		landlord_address_line_1: nil,
		landlord_address_line_2: nil,
		landlord_address_line_3: nil,
		landlord_address_line_4: nil,
		landlord_postcode: nil,
		landlord_number: nil,
		production_1: StringPtr("Hand car wash and valeting"),
		production_information_1: StringPtr("Trade effluent will arrise from the washing / rinsing of vehicles on the wash pad"),
		production_information_2: StringPtr("in the wash bay area."),
		trade_effluent_treatment: StringPtr("Silt trap and a two stage interceptor pit"),
		m3_per_day: StringPtr("5 m\u00b3"),
		litres_per_second: StringPtr("0.5l/s"),
		start_date: nil,
		normal_hours: StringPtr("09:00 - 17:00"),
		batch: StringPtr("Yes"),
		pumped: StringPtr("Yes"),
		sample_location_part_1: StringPtr("Sample can be obtained from the Store's foul manhole that the carwash"),
		sample_location_part_2: StringPtr(" discharges into."),
		public_connection: StringPtr("Connects to the Store's sewer"),
		new_or_existing_connection: StringPtr("existing"),
		foul_sewer: StringPtr("Yes"),
		contaminated_water: StringPtr("Yes"),
		contaminated_water_volume: StringPtr("54"),
		ph_monitoring: StringPtr("Yes"),
		meter_present: StringPtr("Yes"),
		meter_serial: nil,
		average_water_use: StringPtr("1"),
		average_water_other_sources: StringPtr("0"),
		average_days_per_week: StringPtr("7"),
		average_weeks_per_year: StringPtr("52"),
		closures: StringPtr("Chrismas Day, Easter Day, New Year's Day"),
		reasons_why_water_not_returned: StringPtr("N/A"),
		hse_info: StringPtr("nil"),
		env_permit: StringPtr("Yes"),
		env_permit_application_started: nil,
		opening_date: nil,
		job_role: nil,
		company_name: nil,
		phone_no: nil,
		email: nil,
		signature: nil,
		coshh_mention: StringPtr("Please see the attached COSHH sheets"),
	}
}

func StringPtr(s string) *string {
	return &s
}
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
	{T: "2 Trade Effluent Consent Reference Number", 
	Values: formfill.Values{
	    V: formfill.FDFText("test consent number"),
	    },
	},
	{T: "1 Retailer name", 
	Values: formfill.Values{
	    V: formfill.FDFText("test retailer name"),
	    },
	},
}

func MapDataToFromFieldsFDF(appData *ApplicationData) []formfill.FDFField {
	var formFields []formfill.FDFField
	formFields = append(formFields, formfill.FDFField{T: "1 Retailer name", Values: formfill.Values{V: formfill.FDFText(*appData.retailer_name)}})
	formFields = append(formFields, formfill.FDFField{T: "3 Job title", Values: formfill.Values{V: formfill.FDFText("TEsting carwash")}})
	formFields = append(formFields, formfill.FDFField{T: "3 Premises name 1", Values: formfill.Values{V: formfill.FDFText("TEsting carwash")}})
	formFields = append(formFields, formfill.FDFField{T: "2 Trade Effluent Consent Reference Number", Values: formfill.Values{V: formfill.FDFText("21312312312")}})
	formFields = append(formFields, formfill.FDFField{T: "2 SPID(s) relating to the premises", Values: formfill.Values{V: formfill.FDFText("21312312312")}})
	fmt.Println(formFields)
	return formFields
}
func main() {
	const path = "original_sample.pdf"
	//const path = "g02.pdf"
	doc, _, err := reader.ParsePDFFile(path, reader.Options{})
	fmt.Println(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}
	te_application_data := NewApplicationData()
	to_write := MapDataToFromFieldsFDF(te_application_data)
	fmt.Println(to_write)
	fmt.Println(data)
	err = formfill.FillForm(&doc, formfill.FDFDict{Fields: data}, true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	out, err := os.Create("original_sample_filled.pdf")
	//out, err := os.Create("g02filled.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	if err = doc.Write(out, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

}
