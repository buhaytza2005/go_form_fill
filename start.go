package main

import (
	"fmt"
	"os" // Import the os package
	"strings"
	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
)

type ApplicationData struct {
	retailer_name                  *string
	new_consent                    *string
	existing_consent               *string
	consent_reference_or_DPID      *string
	spid                           *string
	company_applicant              *string
	sole_trader_applicant          *string
	applicant_name                 *string
	registered_address_1           *string
	registered_address_2           *string
	registered_address_3           *string
	registered_address_4           *string
	registered_address_postcode    *string
	organisation_is_a_company      *string
	organisation_is_not_a_company  *string
	company_number                 *string
	business                       *string
	site_name                      *string
	site_address_line_1            *string
	site_address_line_2            *string
	site_address_line_3            *string
	site_address_line_4            *string
	site_postcode                  *string
	fp_role                        *string
	landline                       *string
	fp_mobile                      *string
	fp_email                       *string
	fp_name                        *string
	fp_contact_name                *string
	opening_hours                  *string
	same_contact_for_out_of_hours  *string
	owner_is_applying              *string
	landlord_name                  *string
	landlord_address_line_1        *string
	landlord_address_line_2        *string
	landlord_address_line_3        *string
	landlord_address_line_4        *string
	landlord_postcode              *string
	landlord_number                *string
	production_1                   *string
	production_information_1       *string
	production_information_2       *string
	trade_effluent_treatment       *string
	m3_per_day                     *string
	litres_per_second              *string
	start_date                     *string
	normal_hours                   *string
	batch                          *string
	pumped                         *string
	sample_location_part_1         *string
	sample_location_part_2         *string
	public_connection              *string
	new_or_existing_connection     *string
	foul_sewer                     *string
	contaminated_water             *string
	contaminated_water_volume      *string
	ph_monitoring                  *string
	meter_present                  *string
	meter_serial                   *string
	average_water_use              *string
	average_water_other_sources    *string
	average_days_per_week          *string
	average_weeks_per_year         *string
	closures                       *string
	reasons_why_water_not_returned *string
	hse_info                       *string
	env_permit                     *string
	env_permit_application_started *string
	opening_date                   *string
	job_role                       *string
	company_name                   *string
	phone_no                       *string
	email                          *string
	signature                      *string
	coshh_mention                  *string
}

func NewApplicationData() *ApplicationData {
	return &ApplicationData{
		retailer_name:                  StringPtr("Water plus"),
		new_consent:                    nil,
		existing_consent:               nil,
		spid:                           nil,
		applicant_name:                 nil,
		consent_reference_or_DPID:      nil,
		company_applicant:              nil,
		sole_trader_applicant:          nil,
		registered_address_1:           nil,
		registered_address_2:           nil,
		registered_address_3:           nil,
		registered_address_4:           nil,
		registered_address_postcode:    nil,
		organisation_is_a_company:      nil,
		organisation_is_not_a_company:  nil,
		company_number:                 nil,
		business:                       StringPtr("Hand car wash"),
		site_name:                      nil,
		site_address_line_1:            nil,
		site_address_line_2:            nil,
		site_address_line_3:            nil,
		site_address_line_4:            nil,
		site_postcode:                  nil,
		fp_role:                        StringPtr("Franchisee / Site manager"),
		landline:                       StringPtr("N/A"),
		fp_mobile:                      nil,
		fp_email:                       nil,
		fp_name:                        nil,
		fp_contact_name:                nil,
		opening_hours:                  StringPtr("Mon-Sat 9:00-17:00 Sun 10:00-16:00"),
		same_contact_for_out_of_hours:  StringPtr("Yes"),
		owner_is_applying:              StringPtr("No"),
		landlord_name:                  nil,
		landlord_address_line_1:        nil,
		landlord_address_line_2:        nil,
		landlord_address_line_3:        nil,
		landlord_address_line_4:        nil,
		landlord_postcode:              nil,
		landlord_number:                nil,
		production_1:                   StringPtr("Hand car wash and valeting"),
		production_information_1:       StringPtr("Trade effluent will arrise from the washing / rinsing of vehicles on the wash pad"),
		production_information_2:       StringPtr("in the wash bay area."),
		trade_effluent_treatment:       StringPtr("Silt trap and a two stage interceptor pit"),
		m3_per_day:                     StringPtr("5 m\u00b3"),
		litres_per_second:              StringPtr("0.5l/s"),
		start_date:                     nil,
		normal_hours:                   StringPtr("09:00 - 17:00"),
		batch:                          StringPtr("Yes"),
		pumped:                         StringPtr("Yes"),
		sample_location_part_1:         StringPtr("Sample can be obtained from the Store's foul manhole that the carwash"),
		sample_location_part_2:         StringPtr(" discharges into."),
		public_connection:              StringPtr("Connects to the Store's sewer"),
		new_or_existing_connection:     StringPtr("existing"),
		foul_sewer:                     StringPtr("Yes"),
		contaminated_water:             StringPtr("Yes"),
		contaminated_water_volume:      StringPtr("54"),
		ph_monitoring:                  StringPtr("Yes"),
		meter_present:                  StringPtr("Yes"),
		meter_serial:                   nil,
		average_water_use:              StringPtr("1"),
		average_water_other_sources:    StringPtr("0"),
		average_days_per_week:          StringPtr("7"),
		average_weeks_per_year:         StringPtr("52"),
		closures:                       StringPtr("Chrismas Day, Easter Day, New Year's Day"),
		reasons_why_water_not_returned: StringPtr("N/A"),
		hse_info:                       StringPtr("nil"),
		env_permit:                     StringPtr("Yes"),
		env_permit_application_started: nil,
		opening_date:                   nil,
		job_role:                       nil,
		company_name:                   nil,
		phone_no:                       nil,
		email:                          nil,
		signature:                      nil,
		coshh_mention:                  StringPtr("Please see the attached COSHH sheets"),
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
}


func CreateFields(name string, value string) formfill.FDFField {
	if strings.Contains(name, "Yes") {
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else if strings.Contains(name, "No") {
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else if strings.Contains(name, "Tick") {
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else {
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFText(value)}}
	}
}

func MapDataToFromFieldsFDF(appData *ApplicationData) []formfill.FDFField {
	var formFields []formfill.FDFField
	if appData.retailer_name != nil {
		formFields = append(formFields, CreateFields("1 Retailer name", *appData.retailer_name))
	}
	if appData.new_consent != nil {
		formFields = append(formFields, CreateFields("Tick1", *appData.new_consent))
	}
	if appData.existing_consent != nil {
		formFields = append(formFields, CreateFields("Tick4", *appData.existing_consent))
	}
	// page 5))
	// section 2.2 - spid/dpid ))
	if appData.consent_reference_or_DPID != nil {
		formFields = append(formFields, CreateFields("2.2 Trade Effluent Consent Reference Number", *appData.consent_reference_or_DPID))
	}
	if appData.spid != nil {
		formFields = append(formFields, CreateFields("2.2 SPID(s) relating to the premises", *appData.spid))
	}
	// section 3.1 - organisation))
	if appData.company_applicant != nil {
		formFields = append(formFields, CreateFields("Tick7", *appData.company_applicant))
	}
	if appData.sole_trader_applicant != nil {
		formFields = append(formFields, CreateFields("Tick9", *appData.sole_trader_applicant))
	}
	if appData.applicant_name != nil {
		formFields = append(formFields, CreateFields("2", *appData.business))
	}
	if appData.organisation_is_a_company != nil {
		formFields = append(formFields, CreateFields("3.1 Full legal name of company or name of Partner 1", *appData.applicant_name))
	}
	if appData.organisation_is_not_a_company != nil {
		formFields = append(formFields, CreateFields("Yes1", *appData.organisation_is_a_company))
	}
	if appData.company_number != nil {
		formFields = append(formFields, CreateFields("No1", *appData.organisation_is_not_a_company))
	}

	// page 6))
	// section 3.2 registered address))
	if appData.registered_address_1 != nil {
		formFields = append(formFields, CreateFields("3.2 Address line 1", *appData.registered_address_1))
	}
	if appData.registered_address_2 != nil {
		formFields = append(formFields, CreateFields("3.2 Address line 2", *appData.registered_address_2))
	}
	if appData.registered_address_3 != nil {
		formFields = append(formFields, CreateFields("3.2 Address line 3", *appData.registered_address_3))
	}
	if appData.registered_address_4 != nil {
		formFields = append(formFields, CreateFields("3.2 Address line 4", *appData.registered_address_4))
	}
	if appData.registered_address_postcode != nil {
		formFields = append(formFields, CreateFields("3.2 Postcode", *appData.registered_address_postcode))
	}
	if appData.company_number != nil {
		formFields = append(formFields, CreateFields("3.2 Companies' House Reg No", *appData.company_number))
	}
	// section 3.3 Trade premises))
	if appData.business != nil {
		formFields = append(formFields, CreateFields("3.3 Premises name 1", *appData.business))
	}

	//
	//should be Hand Car Wash))
	if appData.site_name != nil {
		formFields = append(formFields, CreateFields("3.3 Premises name 2", *appData.site_name))
	}

	// should be site name))
	if appData.site_address_line_1 != nil {
		formFields = append(formFields, CreateFields("3.3 Address line 1", *appData.site_address_line_1))
	}
	if appData.site_address_line_2 != nil {
		formFields = append(formFields, CreateFields("3.3 Address line 2", *appData.site_address_line_2))
	}
	if appData.site_address_line_3 != nil {
		formFields = append(formFields, CreateFields("3.3 Address line 3", *appData.site_address_line_3))
	}
	if appData.site_address_line_4 != nil {
		formFields = append(formFields, CreateFields("3.3 Address line 4", *appData.site_address_line_4))
	}
	if appData.site_postcode != nil {
		formFields = append(formFields, CreateFields("3.3 Postcode", *appData.site_postcode))
	}
	if appData.fp_role != nil {
		formFields = append(formFields, CreateFields("3.3 Job title", *appData.fp_role))
	}
	// Franchisee / Site manager))
	if appData.landline != nil {
		formFields = append(formFields, CreateFields("3.3 Landline", *appData.landline))
	}
	if appData.fp_mobile != nil {
		formFields = append(formFields, CreateFields("3.3 Mobile", *appData.fp_mobile))
	}
	if appData.fp_email != nil {
		formFields = append(formFields, CreateFields("3.3 Email address", *appData.fp_email))
	}
	if appData.fp_name != nil {
		formFields = append(formFields, CreateFields("3.3 Name of contact", *appData.fp_contact_name))
	}
	// page 7))
	// section 3.4 - hours and contact))
	if appData.opening_hours != nil {
		formFields = append(formFields, CreateFields("3.4 Operational hours for the premises", *appData.opening_hours))
	}

	// Mon-Sat 9:00-17:00 Sun 10:00-16:00))
	if appData.same_contact_for_out_of_hours != nil {
		formFields = append(formFields, CreateFields("Yes3", *appData.same_contact_for_out_of_hours))
	}
	if appData.owner_is_applying != nil {
		formFields = append(formFields, CreateFields("No4", *appData.owner_is_applying))
	}

	// section 3.4 - owner))
	if appData.landlord_name != nil {
		formFields = append(formFields, CreateFields("3.5 Name of owner of the premises", *appData.landlord_name))
	}
	if appData.landlord_address_line_1 != nil {
		formFields = append(formFields, CreateFields("3.5 Address 1", *appData.landlord_address_line_1))
	}
	if appData.landlord_address_line_2 != nil {
		formFields = append(formFields, CreateFields("3.5 Address 2", *appData.landlord_address_line_2))
	}
	if appData.landlord_address_line_3 != nil {
		formFields = append(formFields, CreateFields("3.5 Address 3", *appData.landlord_address_line_3))
	}
	if appData.landlord_address_line_4 != nil {
		formFields = append(formFields, CreateFields("3.5 Address 4", *appData.landlord_address_line_4))
	}
	if appData.landlord_postcode != nil {
		formFields = append(formFields, CreateFields("3.5 Postcode", *appData.landlord_postcode))
	}
	if appData.landlord_number != nil {
		formFields = append(formFields, CreateFields("3.5 Telephone number", *appData.landlord_number))
	}
	// page 8))
	// section 5.1 - Production))
	if appData.production_1 != nil {
		formFields = append(formFields, CreateFields("5.1 Trade conducted at the premises 1", *appData.production_1))
	}
	// Hand car wash and valeting))
	if appData.production_information_1 != nil {
		formFields = append(formFields, CreateFields("5.1 Further information 1", *appData.production_information_1))
	}
	// "Trade effluent will arrise from the washing / rinsing of vehicles on the wash pad "))
	if appData.production_information_2 != nil {
		formFields = append(formFields, CreateFields("5.1 Further information 2", *appData.production_information_2))
	}

	// "in the wash bay area." ))
	// section 5.2 - Treatment))
	if appData.trade_effluent_treatment != nil {
		formFields = append(formFields, CreateFields("5.2 chemical or biological treatment 1", *appData.trade_effluent_treatment))
	}
	// Silt trap and a two stage interceptor pit"))
	// page 10))
	// section 5.4 - Volume of trade effluent))
	if appData.m3_per_day != nil {
		formFields = append(formFields, CreateFields("5.4 m3 per 24 hours", *appData.m3_per_day)) // 5m3))
	}
	if appData.litres_per_second != nil {
		formFields = append(formFields, CreateFields("5.4 litres per second", *appData.litres_per_second)) // 0.5l/s))
	}
	if appData.start_date != nil {
		formFields = append(formFields, CreateFields("5.4 Proposed starting date", *appData.start_date))
	}
	if appData.normal_hours != nil {
		formFields = append(formFields, CreateFields("5.4 Periods of discharge in 24 hour period", *appData.normal_hours))
	}
	if appData.batch != nil {
		formFields = append(formFields, CreateFields("Tick12", *appData.batch))
	}
	if appData.pumped != nil {
		formFields = append(formFields, CreateFields("Tick14", *appData.pumped))
	}
	// page 11 ))
	// section 6 - trade effluent sampling and monitoring))
	if appData.sample_location_part_1 != nil {

		formFields = append(formFields, CreateFields("6.1 Location 1", *appData.sample_location_part_1)) // "Sample can be obtained from the Tesco foul manhole that the car wash"
	}
	if appData.sample_location_part_2 != nil {
		formFields = append(formFields, CreateFields("6.1 Location 2", *appData.sample_location_part_2)) // "discharges into."))
	}

	// section 6.3 Connection to public sewer location))
	if appData.public_connection != nil {
		formFields = append(formFields, CreateFields("6.3 Location", *appData.public_connection)) // Connects to Tesco's sewer))
	}

	// section 6.4 Connection to public sewer location))
	if appData.new_or_existing_connection != nil {
		formFields = append(formFields, CreateFields("Tick16", *appData.new_or_existing_connection))
	}

	if appData.foul_sewer != nil {
		formFields = append(formFields, CreateFields("Tick17", *appData.foul_sewer))
	}
	// section 6.5 contaminated water:))
	if appData.contaminated_water != nil {
		formFields = append(formFields, CreateFields("Yes43", *appData.contaminated_water))
	}

	if appData.contaminated_water_volume != nil {
		formFields = append(formFields, CreateFields("6.5 Surface area draining through sample point", *appData.contaminated_water_volume)) // 54))
	}
	// section 6.6 PH monitoring))
	if appData.ph_monitoring != nil {
		formFields = append(formFields, CreateFields("No44", *appData.ph_monitoring))
	}
	// page 12))
	// section 7.1 water meters))
	if appData.meter_present != nil {
		formFields = append(formFields, CreateFields("No45", *appData.meter_present))
	}
	if appData.meter_serial != nil {
		formFields = append(formFields, CreateFields("6.7 Meter Serial No 1", *appData.meter_serial))
	}
	// section 7.2 site consumption))
	if appData.average_water_use != nil {
		formFields = append(formFields, CreateFields("7.2 Average water consumption", *appData.average_water_use))
	}
	if appData.average_water_other_sources != nil {
		formFields = append(formFields, CreateFields("7.2 Av consumption from other sources", *appData.average_water_other_sources)) // 0))
	}
	if appData.average_days_per_week != nil {
		formFields = append(formFields, CreateFields("7.2 Average number of days worked per week", *appData.average_days_per_week)) // 7))
	}
	if appData.average_weeks_per_year != nil {
		formFields = append(formFields, CreateFields("7.2 Average number of weeks worked per year", *appData.average_weeks_per_year)) // 52))
	}
	if appData.closures != nil {
		formFields = append(formFields, CreateFields("7.2 Approximate date(s) and total days/year", *appData.closures)) // Chrismas day, Easter Sunday, New Year's day))
	}
	if appData.reasons_why_water_not_returned != nil {
		formFields = append(formFields, CreateFields("7.8 Reasons why water not returned to sewer", *appData.reasons_why_water_not_returned)) // N/A))
	}

	// page 13))
	// section 9.1 Health and safety info))
	if appData.hse_info != nil {
		formFields = append(formFields, CreateFields("9.1 Enter details 1", *appData.hse_info))
	}
	// section 9.2 ))
	if appData.env_permit != nil {
		formFields = append(formFields, CreateFields("No46", *appData.env_permit))
	}
	if appData.env_permit_application_started != nil {
		formFields = append(formFields, CreateFields("Yes47", *appData.env_permit_application_started))
	}
	// page 14))
	// section 12 - Declaration))
	if appData.opening_date != nil {
		formFields = append(formFields, CreateFields("12 Date", *appData.opening_date))
	}

	if appData.fp_name != nil {
		formFields = append(formFields, CreateFields("12 FULL NAME", *appData.fp_name))
	}
	if appData.job_role != nil {
		formFields = append(formFields, CreateFields("12 Role in company or job title", *appData.job_role))
	}

	if appData.company_name != nil {
		formFields = append(formFields, CreateFields("12 Company name", *appData.company_name))
	}
	if appData.phone_no != nil {
		formFields = append(formFields, CreateFields("12 Telephone number", *appData.phone_no))
	}
	if appData.email != nil {
		formFields = append(formFields, CreateFields("12 Email address", *appData.email))
	}
	if appData.signature != nil {
		formFields = append(formFields, CreateFields("12 Signature 1", *appData.signature))
	}
	// page 16 ))
	// more information))
	if appData.coshh_mention != nil {
		formFields = append(formFields, CreateFields("15 Please use this page to add further information 1", *appData.coshh_mention))
	}
	return formFields
}
func main() {
	//const path = "original_sample.pdf"
	const path = "g02.pdf"
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
	err = formfill.FillForm(&doc, formfill.FDFDict{Fields: to_write}, true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	//out, err := os.Create("original_sample_filled.pdf")
	out, err := os.Create("g02filled.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

	if err = doc.Write(out, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}

}
