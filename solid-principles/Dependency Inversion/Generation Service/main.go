package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// -----------------------------
// Domain Model
// -----------------------------
type Report struct {
	Title   string
	Content string
}

// -----------------------------
// Abstraction (Formatter interface)
// -----------------------------
type ReportFormatter interface {
	Format(r Report) ([]byte, error)
}

// -----------------------------
// Low-Level Modules (Formatters)
// -----------------------------

// JSONFormatter produces JSON output.
type JSONFormatter struct{}

func (j JSONFormatter) Format(r Report) ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}

// CSVFormatter produces CSV output.
type CSVFormatter struct{}

func (c CSVFormatter) Format(r Report) ([]byte, error) {
	buf := &bytes.Buffer{}
	writer := csv.NewWriter(buf)

	// Header row
	if err := writer.Write([]string{"Title", "Content"}); err != nil {
		return nil, err
	}
	// Data row
	if err := writer.Write([]string{r.Title, r.Content}); err != nil {
		return nil, err
	}

	writer.Flush()
	return buf.Bytes(), writer.Error()
}

// XMLFormatter produces XML output.
type XMLFormatter struct{}

func (x XMLFormatter) Format(r Report) ([]byte, error) {
	// Wrap Report for XML marshalling
	type xmlReport struct {
		XMLName xml.Name `xml:"Report"`
		Title   string   `xml:"Title"`
		Content string   `xml:"Content"`
	}
	xr := xmlReport{Title: r.Title, Content: r.Content}
	return xml.MarshalIndent(xr, "", "  ")
}

// -----------------------------
// High-Level Module (Report Service)
// -----------------------------
type ReportService struct {
	formatter ReportFormatter
}

func NewReportService(f ReportFormatter) *ReportService {
	return &ReportService{formatter: f}
}

// Generate creates and outputs the formatted report.
func (s *ReportService) Generate(r Report) error {
	data, err := s.formatter.Format(r)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

// -----------------------------
// Main: Demonstration
// -----------------------------
func main() {
	report := Report{
		Title:   "Monthly Sales",
		Content: "Total Revenue: $42,000",
	}

	fmt.Println("----- JSON Report -----")
	jsonSvc := NewReportService(JSONFormatter{})
	if err := jsonSvc.Generate(report); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n----- CSV Report -----")
	csvSvc := NewReportService(CSVFormatter{})
	if err := csvSvc.Generate(report); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n----- XML Report -----")
	xmlSvc := NewReportService(XMLFormatter{})
	if err := xmlSvc.Generate(report); err != nil {
		fmt.Println("Error:", err)
	}

	// To add a new format (e.g., PDF), simply implement ReportFormatter—
	// no changes needed in ReportService or the main orchestration.
}
