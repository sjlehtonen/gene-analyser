package report

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/sjlehtonen/gene-analyser/server/dna"
	"github.com/sjlehtonen/gene-analyser/server/html"
)

const REPORTS_PATH = "/reports/"
const MIN_BASE_PAIRS = 200
const MAX_BASE_PAIRS = 3000
const BASE_PAIRS_INCREMENT = 100

type GeneReport struct {
	GeneName              string
	Html                  string
	CpGIslandsPerBasePair []dna.CPGIslandResult
}

func GenerateCpGIslandsTable(islandResults []dna.CPGIslandResult) string {
	sort.Slice(islandResults, func(i, j int) bool {
		return islandResults[i].Basepairs < islandResults[j].Basepairs
	})
	tableData := []html.TableData{}
	for _, v := range islandResults {
		tableData = append(tableData, html.TableData{ColumnName: fmt.Sprint(v.Basepairs), Data: []string{fmt.Sprint(v.Islands)}})
	}
	return html.Elements.Table(tableData)
}

func GenerateReportForGene(geneName string, dnaSequence string, description string) GeneReport {
	upperCaseDna := strings.ToUpper(dnaSequence)
	results := dna.CalculateCpGIslandsSimultaneously(upperCaseDna, MIN_BASE_PAIRS, MAX_BASE_PAIRS, BASE_PAIRS_INCREMENT)
	dnaHtml := html.Elements.P(upperCaseDna)
	cpgIslandsTable := GenerateCpGIslandsTable(results)
	tableStyleCss := "border:1px solid black;"
	tableContainerDivCss := "max-width: 60rem; overflow-wrap: anywhere;"
	htmlString := html.Elements.Html(html.Elements.Style([]string{"table", "th", "td"}, tableStyleCss), html.Elements.Style([]string{"div"}, tableContainerDivCss), html.Elements.P(html.Elements.B("Gene: "), geneName), html.Elements.P(html.Elements.B("Description: "), description), html.Elements.P(html.Elements.B("CpG Islands for selected basepair lengths: ")), cpgIslandsTable, html.Elements.P(html.Elements.B("DNA: "), html.Elements.Div(dnaHtml)))
	geneReport := GeneReport{GeneName: geneName, Html: htmlString, CpGIslandsPerBasePair: results}
	return geneReport
}

func GetReportURL(geneName string) string {
	return fmt.Sprintf("..%s%s.html", REPORTS_PATH, geneName)
}

const WRITE_FILE_PERMISSION = 0644

func (geneReport *GeneReport) WriteReport() {
	os.WriteFile(GetReportURL(geneReport.GeneName), []byte(geneReport.Html), WRITE_FILE_PERMISSION)
}
