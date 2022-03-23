package integration

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const GENE_INFO_FETCH_URL = "https://www.genenetwork.nl/api/v1/gene/"
const DNA_FETCH_BASE_URL = "https://api.genome.ucsc.edu/getData/sequence"

type GeneInfoResult struct {
	Status  int
	Message string
	Gene    struct {
		Chr         string
		Description string
		Start       int
		Stop        int
	}
}

type geneSequenceResult struct {
	Dna        string
	StatusCode int
	Error      string
}

func getGeneInfo(geneName string) (*GeneInfoResult, error) {
	resp, err := http.Get(GENE_INFO_FETCH_URL + geneName)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dat GeneInfoResult
	json.Unmarshal(body, &dat)
	if dat.Status != 0 {
		return nil, errors.New(dat.Message)
	}
	return &dat, err
}

func getDNA(chr string, start int, end int) (string, error) {
	url := fmt.Sprintf("%s?genome=hg38;chrom=chr%s;start=%d;end=%d", DNA_FETCH_BASE_URL, chr, start, end)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var dat geneSequenceResult
	json.Unmarshal(body, &dat)
	if dat.StatusCode != 0 {
		return "", errors.New(dat.Error)
	}
	return dat.Dna, nil
}

func GetDNAAndDescriptionForGene(geneName string) (string, string, error) {
	data, err := getGeneInfo(geneName)
	if err != nil {
		return "", "", err
	}
	dna, err := getDNA(data.Gene.Chr, data.Gene.Start, data.Gene.Stop)
	return dna, data.Gene.Description, err
}
