package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os" // Import the os package
	"strings"

	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
)

type SecondApplicationData struct {
	Retailer_name                  string `json:"retailer_name"`
	New_consent                    string `json:"new_consent"`
	Existing_consent               string `json:"existing_consent"`
	Consent_reference_or_DPID      string `json:"consent_reference_or_DPID"`
	Spid                           string `json:"spid"`
	Company_applicant              string `json:"company_applicant"`
	Sole_trader_applicant          string `json:"sole_trader_applicant"`
	Applicant_name                 string `json:"applicant_name"`
	Registered_address_1           string `json:"registered_address_1"`
	Registered_address_2           string `json:"registered_address_2"`
	Registered_address_3           string `json:"registered_address_3"`
	Registered_address_4           string `json:"registered_address_4"`
	Registered_address_postcode    string `json:"registered_address_postcode"`
	Organisation_is_a_company      string `json:"organisation_is_a_company"`
	Organisation_is_not_a_company  string `json:"organisation_is_not_a_company"`
	Company_number                 string `json:"company_number"`
	Business                       string `json:"business"`
	Site_name                      string `json:"site_name"`
	Site_address_line_1            string `json:"site_address_line_1"`
	Site_address_line_2            string `json:"site_address_line_2"`
	Site_address_line_3            string `json:"site_address_line_3"`
	Site_address_line_4            string `json:"site_address_line_4"`
	Site_postcode                  string `json:"site_postcode"`
	Fp_role                        string `json:"fp_role"`
	Landline                       string `json:"landline"`
	Fp_mobile                      string `json:"fp_mobile"`
	Fp_email                       string `json:"fp_email"`
	Fp_name                        string `json:"fp_name"`
	Fp_contact_name                string `json:"fp_contact_name"`
	Opening_hours                  string `json:"opening_hours"`
	Same_contact_for_out_of_hours  string `json:"same_contact_for_out_of_hours"`
	Owner_is_applying              string `json:"owner_is_applying"`
	Landlord_name                  string `json:"landlord_name"`
	Landlord_address_line_1        string `json:"landlord_address_line_1"`
	Landlord_address_line_2        string `json:"landlord_address_line_2"`
	Landlord_address_line_3        string `json:"landlord_address_line_3"`
	Landlord_address_line_4        string `json:"landlord_address_line_4"`
	Landlord_postcode              string `json:"landlord_postcode"`
	Landlord_number                string `json:"landlord_number"`
	Production_1                   string `json:"production_1"`
	Production_information_1       string `json:"production_information_1"`
	Production_information_2       string `json:"production_information_2"`
	Trade_effluent_treatment       string `json:"trade_effluent_treatment"`
	M3_per_day                     string `json:"m3_per_day"`
	Litres_per_second              string `json:"litres_per_second"`
	Start_date                     string `json:"start_date"`
	Normal_hours                   string `json:"normal_hours"`
	Batch                          string `json:"batch"`
	Pumped                         string `json:"pumped"`
	Sample_location_part_1         string `json:"sample_location_part_1"`
	Sample_location_part_2         string `json:"sample_location_part_2"`
	Public_connection              string `json:"public_connection"`
	New_or_existing_connection     string `json:"new_or_existing_connection"`
	Foul_sewer                     string `json:"foul_sewer"`
	Contaminated_water             string `json:"contaminated_water"`
	Contaminated_water_volume      string `json:"contaminated_water_volume"`
	Ph_monitoring                  string `json:"ph_monitoring"`
	Meter_present                  string `json:"meter_present"`
	Meter_serial                   string `json:"meter_serial"`
	Average_water_use              string `json:"average_water_use"`
	Average_water_other_sources    string `json:"average_water_other_sources"`
	Average_days_per_week          string `json:"average_days_per_week"`
	Average_weeks_per_year         string `json:"average_weeks_per_year"`
	Closures                       string `json:"closures"`
	Reasons_why_water_not_returned string `json:"reasons_why_water_not_returned"`
	Hse_info                       string `json:"hse_info"`
	Env_permit                     string `json:"env_permit"`
	Env_permit_application_started string `json:"env_permit_application_started"`
	Opening_date                   string `json:"opening_date"`
	Job_role                       string `json:"job_role"`
	Company_name                   string `json:"company_name"`
	Phone_no                       string `json:"phone_no"`
	Email                          string `json:"email"`
	Signature                      string `json:"signature"`
	Coshh_mention                  string `json:"coshh_mention"`



}

func StringPtr(s string) *string {
	return &s
}

func CreateFields2(name string, value string) formfill.FDFField {
    if strings.Contains(name, "Yes") || strings.Contains(name, "No") || strings.Contains(name, "Tick") {
        // For checkboxes or similar fields where 'Yes', 'No', or 'Tick' are valid values
        fmt.Println("Creating checkbox/radio field ", name, " with value ", value)
        return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
    } else if strings.Contains(name, "Company number") {
        // For the 'company number' field, which is a text field
        fmt.Println("Creating text field ", name, " with value ", value)
        return formfill.FDFField{
            T: "3",
            Kids: []formfill.FDFField{
                {
                    T: name,
                    Values: formfill.Values{V: formfill.FDFText(value)}, // Use FDFText for text fields
                },
            },
        }
    } else {
        // For other general text fields
        fmt.Println("Creating text field ", name, " with value ", value)
        return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFText(value)}}
    }
}
func CreateFields(name string, value string) formfill.FDFField {
	if strings.Contains(name, "Companies' House Reg No"){
		// For the 'company number' field, which is a text field
		fmt.Println("Creating text field in main ", name, " with value ", value)
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFText(value)}}
	} else if strings.Contains(name, "Yes") {
		fmt.Println("Creating field ", name, " with value ", value)
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else if strings.Contains(name, "No") {
		fmt.Println("Creating field ", name, " with value ", value)
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else if strings.Contains(name, "Tick") {
		fmt.Println("Creating field ", name, " with value ", value)
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFName(value)}}
	} else {
		fmt.Println("Creating field ", name, " with value ", value)
		return formfill.FDFField{T: name, Values: formfill.Values{V: formfill.FDFText(value)}}
	}
}
func MapData2ToFields(appData SecondApplicationData) []formfill.FDFField {
	var formFields []formfill.FDFField
	if appData.Retailer_name != "" {
		formFields = append(formFields, CreateFields("1 Retailer name", appData.Retailer_name))
	}
	if appData.New_consent != "" {
		formFields = append(formFields, CreateFields("Tick1", appData.New_consent))
	}
	if appData.Existing_consent != "" {
		formFields = append(formFields, CreateFields("Tick4", appData.Existing_consent))
	}
	// page 5))
	// section 2.2 - spid/dpid ))
	if appData.Consent_reference_or_DPID != "" {
		formFields = append(formFields, CreateFields("2.2 Trade Effluent Consent Reference Number", appData.Consent_reference_or_DPID))
	}
	if appData.Spid != "" {
		formFields = append(formFields, CreateFields("2.2 SPID(s) relating to the premises", appData.Spid))
	}
	doc.Catalog.Flatten
	// section 3.1 - organisation))
	if appData.Company_applicant != "" {
		formFields = append(formFields, CreateFields("Tick7", appData.Company_applicant))
	}
	if appData.Sole_trader_applicant != "" {
		formFields = append(formFields, CreateFields("Tick9", appData.Sole_trader_applicant))
		formFields = append(formFields, CreateFields("No1", appData.Sole_trader_applicant))
	}
	if appData.Business != "" {
		formFields = append(formFields, CreateFields("2", appData.Applicant_name))
	}

	// need logic to determine is sole trader or company
	if appData.Organisation_is_a_company != "" {
		formFields = append(formFields, CreateFields("3.1 Full legal name of company or name of Partner 1", appData.Applicant_name))
		formFields = append(formFields, CreateFields("Yes1", appData.Organisation_is_a_company))
	}

	// page 6))
	// section 3.2 registered address))
	if appData.Registered_address_1 != "" {
		formFields = append(formFields, CreateFields("3.2 Address line 1", appData.Registered_address_1))
	}
	if appData.Registered_address_2 != "" {
		formFields = append(formFields, CreateFields("3.2 Address line 2", appData.Registered_address_2))
	}
	if appData.Registered_address_3 != "" {
		formFields = append(formFields, CreateFields("3.2 Address line 3", appData.Registered_address_3))
	}
	if appData.Registered_address_4 != "" {
		formFields = append(formFields, CreateFields("3.2 Address line 4", appData.Registered_address_4))
	}
	if appData.Registered_address_postcode != "" {
		formFields = append(formFields, CreateFields("3.2 Postcode", appData.Registered_address_postcode))
	}
	if appData.Company_number != "" {
		fmt.Println("Company number is ", appData.Company_number)
		formFields = append(formFields, CreateFields("3.2 Companies' House Reg No", appData.Company_number))
	}
	// section 3.3 Trade premises))
	if appData.Business != "" {
		formFields = append(formFields, CreateFields("3.3 Premises name 1", appData.Business))
	}

	//
	//should be Hand Car Wash))
	if appData.Site_name != "" {
		formFields = append(formFields, CreateFields("3.3 Premises name 2", appData.Site_name))
	}

	// should be site name))
	if appData.Site_address_line_1 != "" {
		formFields = append(formFields, CreateFields("3.3 Address line 1", appData.Site_address_line_1))
	}
	if appData.Site_address_line_2 != "" {
		formFields = append(formFields, CreateFields("3.3 Address line 2", appData.Site_address_line_2))
	}
	if appData.Site_address_line_3 != "" {
		formFields = append(formFields, CreateFields("3.3 Address line 3", appData.Site_address_line_3))
	}
	if appData.Site_address_line_4 != "" {
		formFields = append(formFields, CreateFields("3.3 Address line 4", appData.Site_address_line_4))
	}
	if appData.Site_postcode != "" {
		formFields = append(formFields, CreateFields("3.3 Postcode", appData.Site_postcode))
	}
	if appData.Fp_role != "" {
		formFields = append(formFields, CreateFields("3.3 Job title", appData.Fp_role))
	}
	// Franchisee / Site manager))
	if appData.Landline != "" {
		formFields = append(formFields, CreateFields("3.3 Landline", appData.Landline))
	}
	if appData.Fp_mobile != "" {
		formFields = append(formFields, CreateFields("3.3 Mobile", appData.Fp_mobile))
	}
	if appData.Fp_email != "" {
		formFields = append(formFields, CreateFields("3.3 Email address", appData.Fp_email))
	}
	if appData.Fp_name != "" {
		formFields = append(formFields, CreateFields("3.3 Name of contact", appData.Fp_name))
	}
	// page 7))
	// section 3.4 - hours and contact))
	if appData.Opening_hours != "" {
		formFields = append(formFields, CreateFields("3.4 Operational hours for the premises", appData.Opening_hours))
	}

	// Mon-Sat 9:00-17:00 Sun 10:00-16:00))
	if appData.Same_contact_for_out_of_hours != "" {
		formFields = append(formFields, CreateFields("Yes3", appData.Same_contact_for_out_of_hours))
	}
	if appData.Owner_is_applying != "" {
		formFields = append(formFields, CreateFields("No4", appData.Owner_is_applying))
	}

	// section 3.4 - owner))
	if appData.Landlord_name != "" {
		formFields = append(formFields, CreateFields("3.5 Name of owner of the premises", appData.Landlord_name))
	}
	if appData.Landlord_address_line_1 != "" {
		formFields = append(formFields, CreateFields("3.5 Address 1", appData.Landlord_address_line_1))
	}
	if appData.Landlord_address_line_2 != "" {
		formFields = append(formFields, CreateFields("3.5 Address 2", appData.Landlord_address_line_2))
	}
	if appData.Landlord_address_line_3 != "" {
		formFields = append(formFields, CreateFields("3.5 Address 3", appData.Landlord_address_line_3))
	}
	if appData.Landlord_address_line_4 != "" {
		formFields = append(formFields, CreateFields("3.5 Address 4", appData.Landlord_address_line_4))
	}
	if appData.Landlord_postcode != "" {
		formFields = append(formFields, CreateFields("3.5 Postcode", appData.Landlord_postcode))
	}
	if appData.Landlord_number != "" {
		formFields = append(formFields, CreateFields("3.5 Telephone number", appData.Landlord_number))
	}
	// page 8))
	// section 5.1 - Production))
	if appData.Production_1 != "" {
		formFields = append(formFields, CreateFields("5.1 Trade conducted at the premises 1", appData.Production_1))
	}
	// Hand car wash and valeting))
	if appData.Production_information_1 != "" {
		formFields = append(formFields, CreateFields("5.1 Further information 1", appData.Production_information_1))
	}
	// "Trade effluent will arrise from the washing / rinsing of vehicles on the wash pad "))
	if appData.Production_information_2 != "" {
		formFields = append(formFields, CreateFields("5.1 Further information 2", appData.Production_information_2))
	}

	// "in the wash bay area." ))
	// section 5.2 - Treatment))
	if appData.Trade_effluent_treatment != "" {
		formFields = append(formFields, CreateFields("5.2 chemical or biological treatment 1", appData.Trade_effluent_treatment))
	}
	// Silt trap and a two stage interceptor pit"))
	// page 10))
	// section 5.4 - Volume of trade effluent))
	if appData.M3_per_day != "" {
		formFields = append(formFields, CreateFields("5.4 m3 per 24 hours", appData.M3_per_day)) // 5m3))
	}
	if appData.Litres_per_second != "" {
		formFields = append(formFields, CreateFields("5.4 litres per second", appData.Litres_per_second)) // 0.5l/s))
	}
	if appData.Start_date != "" {
		formFields = append(formFields, CreateFields("5.4 Proposed starting date", appData.Start_date))
	}
	if appData.Normal_hours != "" {
		formFields = append(formFields, CreateFields("5.4 Periods of discharge in 24 hour period", appData.Normal_hours))
	}
	if appData.Batch != "" {
		formFields = append(formFields, CreateFields("Tick12", appData.Batch))
	}
	if appData.Pumped != "" {
		formFields = append(formFields, CreateFields("Tick14", appData.Pumped))
	}
	// page 11 ))
	// section 6 - trade effluent sampling and monitoring))
	if appData.Sample_location_part_1 != "" {

		formFields = append(formFields, CreateFields("6.1 Location 1", appData.Sample_location_part_1)) // "Sample can be obtained from the Tesco foul manhole that the car wash"
	}
	if appData.Sample_location_part_2 != "" {
		formFields = append(formFields, CreateFields("6.1 Location 2", appData.Sample_location_part_2)) // "discharges into."))
	}

	// section 6.3 Connection to public sewer location))
	if appData.Public_connection != "" {
		formFields = append(formFields, CreateFields("6.3 Location", appData.Public_connection)) // Connects to Tesco's sewer))
	}

	// section 6.4 Connection to public sewer location))
	if appData.New_or_existing_connection != "" {
		formFields = append(formFields, CreateFields("Tick16", appData.New_or_existing_connection))
	}

	if appData.Foul_sewer != "" {
		formFields = append(formFields, CreateFields("Tick17", appData.Foul_sewer))
	}
	// section 6.5 contaminated water:))
	if appData.Contaminated_water != "" {
		formFields = append(formFields, CreateFields("Yes43", appData.Contaminated_water))
	}

	if appData.Contaminated_water_volume != "" {
		formFields = append(formFields, CreateFields("6.5 Surface area draining through sample point", appData.Contaminated_water_volume)) // 54))
	}
	// section 6.6 PH monitoring))
	if appData.Ph_monitoring != "" {
		formFields = append(formFields, CreateFields("No44", appData.Ph_monitoring))
	}
	// page 12))
	// section 7.1 water meters))
	if appData.Meter_present != "" {
		formFields = append(formFields, CreateFields("No45", appData.Meter_present))
	}
	if appData.Meter_serial != "" {
		formFields = append(formFields, CreateFields("6.7 Meter Serial No 1", appData.Meter_serial))
	}
	// section 7.2 site consumption))
	if appData.Average_water_use != "" {
		formFields = append(formFields, CreateFields("7.2 Average water consumption", appData.Average_water_use))
	}
	if appData.Average_water_other_sources != "" {
		formFields = append(formFields, CreateFields("7.2 Av consumption from other sources", appData.Average_water_other_sources)) // 0))
	}
	if appData.Average_days_per_week != "" {
		formFields = append(formFields, CreateFields("7.2 Average number of days worked per week", appData.Average_days_per_week)) // 7))
	}
	if appData.Average_weeks_per_year != "" {
		formFields = append(formFields, CreateFields("7.2 Average number of weeks worked per year", appData.Average_weeks_per_year)) // 52))
	}
	if appData.Closures != "" {
		formFields = append(formFields, CreateFields("7.2 Approximate date(s) and total days/year", appData.Closures)) // Chrismas day, Easter Sunday, New Year's day))
	}
	if appData.Reasons_why_water_not_returned != "" {
		formFields = append(formFields, CreateFields("7.8 Reasons why water not returned to sewer", appData.Reasons_why_water_not_returned)) // N/A))
	}

	// page 13))
	// section 9.1 Health and safety info))
	if appData.Hse_info != "" {
		formFields = append(formFields, CreateFields("9.1 Enter details 1", appData.Hse_info))
	}
	// section 9.2 ))
	if appData.Env_permit != "" {
		formFields = append(formFields, CreateFields("No46", appData.Env_permit))
	}
	if appData.Env_permit_application_started != "" {
		formFields = append(formFields, CreateFields("Yes47", appData.Env_permit_application_started))
	}
	// page 14))
	// section 12 - Declaration))
	if appData.Opening_date != "" {
		formFields = append(formFields, CreateFields("12 Date", appData.Opening_date))
	}

	if appData.Fp_name != "" {
		formFields = append(formFields, CreateFields("12 FULL NAME", appData.Fp_name))
	}
	if appData.Job_role != "" {
		formFields = append(formFields, CreateFields("12 Role in company or job title", appData.Job_role))
	}

	if appData.Company_name != "" {
		formFields = append(formFields, CreateFields("12 Company name", appData.Company_name))
	}
	if appData.Fp_mobile != "" {
		formFields = append(formFields, CreateFields("12 Telephone number", appData.Fp_mobile))
	}
	if appData.Fp_email != "" {
		formFields = append(formFields, CreateFields("12 Email address", appData.Fp_email))
	}
	if appData.Signature != "" {
		formFields = append(formFields, CreateFields("12 Signature 1", appData.Signature))
	}
	// page 16 ))
	// more information))
	if appData.Coshh_mention != "" {
		formFields = append(formFields, CreateFields("15 Please use this page to add further information 1", appData.Coshh_mention))
	}
	return formFields
}
func main() {


	// get data from rust server as a test

	resp, err := http.Get("http://localhost:9999/sites")
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()




	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Raw JSON data:", string(body))

	//get the body into the Application data struct
	var response Response
	marshall_err := json.Unmarshal(body, &response)
	if marshall_err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
    fmt.Println(response)


	const path = "g02.pdf"
	doc, _, err := reader.ParsePDFFile(path, reader.Options{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1) // Exit the program if there is an error
	}
	//te_application_data := NewApplicationData()
	
	to_write := MapData2ToFields(response.Data[0])
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

type Response struct {
	Data []SecondApplicationData `json:"data"`
}
