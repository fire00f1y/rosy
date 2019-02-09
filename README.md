# Rosy

## Overview
Parse a file and look for xml content desired

## Usage
Run the binary with either the `-h` or `--help` flag to pring usage info
```bash
-c string
    Closing tag to match on. This only needs to be specified if it is not the companion tag to -t
-f string
    Which file to scan. Scans stdin if no file provided (enables piping to this utility)
-o string
    Output file. If left blank, will output to stdout
-t string
    Tag to look for. Will return whole file by default.
```

## Examples
### Reading file and outputing to stdout
```bash
$ rosy.exe -t "<xmlcontent>" -f sample_text.txt
<?xmlversion="1.0"encoding="utf-8"?><!--xslplanes.1.xsl.1AnXSLTstylesheetforxslplane.xmlusingchildtemplates--><xsl:stylesheetversion="1.0"xmlns:xsl="http://www.w3.org/1999/XSL/Transform"xmlns="http://www.w3.org/1999/xhtml"><!--Thetemplateforthewholedocument(theplaneelement)--><xsl:templatematch="plane"><html><head><title>Stylesheetforxslplane.xml</title></head><body><h2>AirplaneDescription</h2><!--Applythematchingtemplatestotheelementsinplane--><xsl:apply-templates/></body></html></xsl:template><!--Thetemplatestobeapplied(byapply-templates)totheelementsintheplaneelement--><xsl:templatematch="year"><spanstyle="font-style:italic;color:blue;">Year:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="make"><spanstyle="font-style:italic;color:blue;">Make:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="model"><spanstyle="font-style:italic;color:blue;">Model:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="color"><spanstyle="font-style:italic;color:blue;">Color:</span><xsl:value-ofselect="."/><br/></xsl:template></xsl:stylesheet></xmlcontent>from</xsl:stylesheet>
```

### Reading file and outputting to file
```bash
$ ls
go.mod  rosy.exe*  rosy.go  sample_text.txt
$ rosy.exe -t "<xmlcontent>" -f sample_text.txt  -o results.xml
$ ls
go.mod  results.xml  rosy.exe*  rosy.go  sample_text.txt
```

### Piping data into the utility into stdout
```bash
$ cat sample_text.txt | rosy.exe -t "<xmlcontent>"
<?xmlversion="1.0"encoding="utf-8"?><!--xslplanes.1.xsl.1AnXSLTstylesheetforxslplane.xmlusingchildtemplates--><xsl:stylesheetversion="1.0"xmlns:xsl="http://www.w3.org/1999/XSL/Transform"xmlns="http://www.w3.org/1999/xhtml"><!--Thetemplateforthewholedocument(theplaneelement)--><xsl:templatematch="plane"><html><head><title>Stylesheetforxslplane.xml</title></head><body><h2>AirplaneDescription</h2><!--Applythematchingtemplatestotheelementsinplane--><xsl:apply-templates/></body></html></xsl:template><!--Thetemplatestobeapplied(byapply-templates)totheelementsintheplaneelement--><xsl:templatematch="year"><spanstyle="font-style:italic;color:blue;">Year:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="make"><spanstyle="font-style:italic;color:blue;">Make:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="model"><spanstyle="font-style:italic;color:blue;">Model:</span><xsl:value-ofselect="."/><br/></xsl:template><xsl:templatematch="color"><spanstyle="font-style:italic;color:blue;">Color:</span><xsl:value-ofselect="."/><br/></xsl:template></xsl:stylesheet></xmlcontent>from</xsl:stylesheet>
```

### Piping data into the utility into new file
```bash
$ ls
go.mod  rosy.exe*  rosy.go  sample_text.txt
$ cat sample_text.txt | rosy.exe -t "<xmlcontent>" -o results.xml
$ ls
go.mod  results.xml  rosy.exe*  rosy.go  sample_text.txt
```