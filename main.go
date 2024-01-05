// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at

//   http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Parses a transposed file and calculates the minor allele frequency
// File should be in the format of markers x lines with each column
// being a line. The first column should be the marker name.
// The first line should be the header with the sample names.
// The output is a tab delimited file with the following columns:
// Marker, Minor Allele, MAF, MAF %, Total Allele Count, Missing, Counts
// The counts column is a tab delimited list of the allele counts for each
// allele.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// To run upate the file name in line 39 and run: go run main.go > output.txt

func main() {
	file, err := os.Open("sample_data.txt") // change this to your file name
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Marker\tMinor Allele\tMAF\tMAF %\tTotal Allele Count\tMissing\tCounts") // Header for the output

	const maxCapacity int = 2000000 // your required line length (approx 2MB)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	firstRow := true

	for scanner.Scan() { // This allows us to skip over the first row
		if firstRow {
			firstRow = false
			continue
		}
		minorAlleleFrequency(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func minorAlleleFrequency(calls string) {
	counts := make(map[string]int)
	total := 0
	missing := 0

	c := strings.Split(calls, "\t")

	for _, nucleotide := range c[1:] {
		if nucleotide == "H" || nucleotide == "-" { //account for the missing data
			missing++
		} else { // count the alleles
			counts[nucleotide]++
			total++
		}
	}

	minCount := total
	minorAllele := "MARKER FAILED" // This is the default if something goes wrong

	var stringBuild strings.Builder

	for i, count := range counts {
		if count < minCount {
			minCount = count
			minorAllele = i
		} // this gets the lowest count
		stringBuild.WriteString(i + ":" + fmt.Sprintf("%d", count) + "\t")
	}

	frequency := float64(minCount) / float64(total)
	frequency_percent := frequency * 100

	if frequency == 1 { // Then there is only one allele and we can't calculate MAF (its monomorphic)
		minorAllele = "MONOMORPHIC"
	}

	fmt.Printf("%s\t%s\t%.2f\t%.3f\t%d\t%d\t%s\n", c[0], minorAllele, frequency, frequency_percent, total, missing, &stringBuild)

}
