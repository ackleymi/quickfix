package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
	"os"
	"sort"
)

var (
	fieldMap     map[string]int
	fieldTypeMap map[string]*datadictionary.FieldType
	sortedTags   []string
	pkg          = "fix"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genEnums() {
	fileOut := "package enum\n"

	for _, fieldName := range sortedTags {
		fieldType, _ := fieldTypeMap[fieldName]
		if len(fieldType.Enums) == 0 {
			continue
		}

		sortedEnums := make([]string, len(fieldType.Enums))
		i := 0
		for enum, _ := range fieldType.Enums {
			sortedEnums[i] = enum
			i++
		}
		sort.Strings(sortedEnums)

		fileOut += fmt.Sprintf("//Enum values for %v\n", fieldName)
		fileOut += "const(\n"
		for _, enumVal := range sortedEnums {
			enum, _ := fieldType.Enums[enumVal]
			fileOut += fmt.Sprintf("%v_%v = \"%v\"\n", fieldName, enum.Description, enum.Value)
		}
		fileOut += ")\n"
	}

	gen.WriteFile("fix/enum/enums.go", fileOut)
}

func genFields() {
	fileOut := "package field\n"
	fileOut += "import(\n"
	fileOut += "\"github.com/quickfixgo/quickfix/fix\"\n"
	fileOut += "\"github.com/quickfixgo/quickfix/message\"\n"
	fileOut += "\"github.com/quickfixgo/quickfix/fix/tag\"\n"
	fileOut += ")\n"

	for _, tag := range sortedTags {
		field := fieldTypeMap[tag]

		baseType := ""
		goType := ""
		switch field.Type {
		case "STRING":
			baseType = "StringValue"
			goType = "string"
		case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
			baseType = "MultipleStringValue"
			goType = "string"
		case "MULTIPLECHARVALUE":
			baseType = "MultipleCharValue"
			goType = "string"
		case "CHAR":
			baseType = "CharValue"
			goType = "string"
		case "CURRENCY":
			baseType = "CurrencyValue"
			goType = "string"
		case "DATA":
			baseType = "DataValue"
			goType = "string"
		case "MONTHYEAR":
			baseType = "MonthYearValue"
			goType = "string"
		case "LOCALMKTDATE":
			baseType = "LocalMktDateValue"
			goType = "string"
		case "EXCHANGE":
			baseType = "ExchangeValue"
			goType = "string"
		case "LANGUAGE":
			baseType = "LanguageValue"
			goType = "string"
		case "XMLDATA":
			baseType = "XMLDataValue"
			goType = "string"
		case "COUNTRY":
			baseType = "CountryValue"
			goType = "string"
		case "UTCTIMEONLY":
			baseType = "UTCTimeOnlyValue"
		case "UTCDATEONLY":
			baseType = "UTCDateOnlyValue"
		case "TZTIMEONLY":
			baseType = "TZTimeOnlyValue"
		case "TZTIMESTAMP":
			baseType = "TZTimestampValue"
		case "BOOLEAN":
			baseType = "BooleanValue"
			goType = "bool"
		case "INT":
			baseType = "IntValue"
			goType = "int"
		case "LENGTH":
			baseType = "LengthValue"
			goType = "int"
		case "DAYOFMONTH":
			baseType = "DayOfMonthValue"
			goType = "int"
		case "NUMINGROUP":
			baseType = "NumInGroupValue"
			goType = "int"
		case "SEQNUM":
			baseType = "SeqNumValue"
			goType = "int"
		case "UTCTIMESTAMP":
			baseType = "UTCTimestampValue"
		case "FLOAT":
			baseType = "FloatValue"
			goType = "float64"
		case "QTY":
			baseType = "QtyValue"
			goType = "float64"
		case "AMT":
			baseType = "AmtValue"
			goType = "float64"
		case "PRICE":
			baseType = "PriceValue"
			goType = "float64"
		case "PRICEOFFSET":
			baseType = "PriceOffsetValue"
			goType = "float64"
		case "PERCENTAGE":
			baseType = "PercentageValue"
			goType = "float64"
		default:
			fmt.Printf("Unknown type '%v' for tag '%v'\n", field.Type, tag)
		}

		fileOut += fmt.Sprintf("//%v is a %v field\n", field.Name, field.Type)
		fileOut += fmt.Sprintf("type %v struct { message.%v }\n", field.Name, baseType)
		fileOut += fmt.Sprintf("//Tag returns tag.%v (%v)\n", field.Name, field.Tag)
		fileOut += fmt.Sprintf("func (f %v) Tag() fix.Tag {return tag.%v}\n", field.Name, field.Name)

		switch goType {
		case "string", "int", "float64", "bool":
			fileOut += fmt.Sprintf("//Build%v returns a new %v initialized with val\n", field.Name, field.Name)
			fileOut += fmt.Sprintf("func Build%v(val %v) %v {\n", field.Name, goType, field.Name)
			fileOut += fmt.Sprintf("var field %v\n", field.Name)
			fileOut += "field.Value = val\n"
			fileOut += "return field\n"
			fileOut += "}\n"
		}
	}

	gen.WriteFile("fix/field/fields.go", fileOut)
}

func genTags() {
	fileOut := "package tag\n"
	fileOut += "import(\"github.com/quickfixgo/quickfix/fix\")\n"

	fileOut += "const (\n"
	for _, tag := range sortedTags {
		fileOut += fmt.Sprintf("%v fix.Tag = %v\n", tag, fieldMap[tag])
	}
	fileOut += ")\n"

	gen.WriteFile("fix/tag/tag_numbers.go", fileOut)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	fieldMap = make(map[string]int)
	fieldTypeMap = make(map[string]*datadictionary.FieldType)

	for _, dataDict := range flag.Args() {
		spec, err := datadictionary.Parse(dataDict)

		if err != nil {
			panic(err)
		}

		for _, field := range spec.FieldTypeByTag {
			fieldMap[field.Name] = int(field.Tag)
			fieldTypeMap[field.Name] = field
		}
	}

	sortedTags = make([]string, len(fieldMap))
	i := 0
	for f := range fieldMap {
		sortedTags[i] = f
		i++
	}
	sort.Strings(sortedTags)

	genTags()
	genFields()
	genEnums()
}
