Quick Go program to calculate minor allele frequencies across large SNP datasets. In this case for the barley 50K SNP chip developed at the James Hutton Institute/International Barley Hub but may be useful for anyone with data in a similar format.

Data needs to be in markers x lines format. If your data is in the other orientation use something like datamash (https://www.gnu.org/software/datamash/) to transpose the file. 

The example data file sample_data.txt shows the format. Line names are not important for this as all its doing is counting the number of alleles for each marker then working out the allele with the lowest count and using this as the minor allele which it then displays as a count/percentage etc.

To run: go run main.go > [output.txt] after you have updated main.go with your input filename. I will change this to take command line input in the future but at the moment you are stuck with it like this :-)
