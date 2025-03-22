package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// SearchQuery struct - Immutable once created
type SearchQuery struct {
	keyword     string
	location    string
	startDate   string
	endDate     string
	fileType    string
	sortOrder   string
	maxResults  int
}

// SearchQueryBuilder - Builder Class
type SearchQueryBuilder struct {
	keyword     string
	location    string
	startDate   string
	endDate     string
	fileType    string
	sortOrder   string
	maxResults  int
}

// NewSearchQueryBuilder - Initializes a new SearchQueryBuilder
func NewSearchQueryBuilder(keyword string) *SearchQueryBuilder {
	return &SearchQueryBuilder{
		keyword: keyword,
	}
}

// WithLocation - Adds location filter
func (b *SearchQueryBuilder) WithLocation(location string) *SearchQueryBuilder {
	b.location = location
	return b
}

// WithDateRange - Adds start and end date filter
func (b *SearchQueryBuilder) WithDateRange(startDate, endDate string) *SearchQueryBuilder {
	b.startDate = startDate
	b.endDate = endDate
	return b
}

// WithFileType - Adds file type filter
func (b *SearchQueryBuilder) WithFileType(fileType string) *SearchQueryBuilder {
	b.fileType = fileType
	return b
}

// WithSortOrder - Adds sort order filter
func (b *SearchQueryBuilder) WithSortOrder(order string) *SearchQueryBuilder {
	b.sortOrder = order
	return b
}

// WithMaxResults - Limits the maximum number of results
func (b *SearchQueryBuilder) WithMaxResults(limit int) *SearchQueryBuilder {
	b.maxResults = limit
	return b
}

// Build - Constructs the immutable SearchQuery object
func (b *SearchQueryBuilder) Build() SearchQuery {
	return SearchQuery{
		keyword:    b.keyword,
		location:   b.location,
		startDate:  b.startDate,
		endDate:    b.endDate,
		fileType:   b.fileType,
		sortOrder:  b.sortOrder,
		maxResults: b.maxResults,
	}
}

// DisplayQuery - Displays the generated search query
func DisplayQuery(query SearchQuery) {
	var queryParams []string
	queryParams = append(queryParams, fmt.Sprintf("Keyword: %s", query.keyword))
	if query.location != "" {
		queryParams = append(queryParams, fmt.Sprintf("Location: %s", query.location))
	}
	if query.startDate != "" && query.endDate != "" {
		queryParams = append(queryParams, fmt.Sprintf("Date Range: %s to %s", query.startDate, query.endDate))
	}
	if query.fileType != "" {
		queryParams = append(queryParams, fmt.Sprintf("File Type: %s", query.fileType))
	}
	if query.sortOrder != "" {
		queryParams = append(queryParams, fmt.Sprintf("Sort Order: %s", query.sortOrder))
	}
	if query.maxResults > 0 {
		queryParams = append(queryParams, fmt.Sprintf("Max Results: %d", query.maxResults))
	}

	fmt.Println("🔎 Search Query:")
	fmt.Println(strings.Join(queryParams, " | "))

	// Simulated delay to demonstrate concurrent execution
	time.Sleep(2 * time.Second)
	fmt.Println("✅ Search results fetched successfully.")
}

// ConcurrentSearch - Executes searches concurrently using Goroutines
func ConcurrentSearch(wg *sync.WaitGroup, query SearchQuery) {
	defer wg.Done()
	DisplayQuery(query)
}

func main() {
	var wg sync.WaitGroup

	// Search 1: Complex Search with Multiple Filters
	query1 := NewSearchQueryBuilder("Artificial Intelligence").
		WithLocation("San Francisco").
		WithDateRange("2023-01-01", "2023-12-31").
		WithFileType("PDF").
		WithSortOrder("Ascending").
		WithMaxResults(50).
		Build()

	// Search 2: Simple Keyword Search with Limited Results
	query2 := NewSearchQueryBuilder("Machine Learning").
		WithMaxResults(20).
		Build()

	// Search 3: Additional Custom Search
	query3 := NewSearchQueryBuilder("Data Science").
		WithLocation("New York").
		WithFileType("DOCX").
		WithSortOrder("Descending").
		Build()

	// Run Concurrent Search Queries
	wg.Add(3)
	go ConcurrentSearch(&wg, query1)
	go ConcurrentSearch(&wg, query2)
	go ConcurrentSearch(&wg, query3)

	wg.Wait()

	fmt.Println("🚀 All search queries processed successfully.")
}
