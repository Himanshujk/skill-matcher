package helpers

// Don't remove or alter these keys as they represent versions or certifications that include numbers.
var PreserveNumbers = map[string]bool{

	// =========================
	// WEB / FRONTEND STANDARDS
	// =========================
	"html5":  true,
	"css3":   true,
	"web2":   true,
	"web3":   true,
	"web2.0": true,
	"web3.0": true,
	"es6":    true,
	"es7":    true,
	"es8":    true,
	"es2015": true,
	"es2016": true,
	"es2017": true,
	"es2018": true,
	"es2019": true,
	"es2020": true,
	"es2021": true,
	"es2022": true,

	// =========================
	// AI / ML MODELS
	// =========================
	"gpt2":     true,
	"gpt3":     true,
	"gpt4":     true,
	"llama2":   true,
	"bert2":    true,
	"yolov3":   true,
	"yolov4":   true,
	"yolov5":   true,
	"resnet50": true,

	// =========================
	// ANGULAR (major architectural shift after v2)
	// =========================
	"angular2":  true,
	"angular4":  true,
	"angular5":  true,
	"angular6":  true,
	"angular7":  true,
	"angular8":  true,
	"angular9":  true,
	"angular10": true,
	"angular11": true,
	"angular12": true,
	"angular13": true,
	"angular14": true,
	"angular15": true,
	"angular16": true,

	// =========================
	// .NET MODERN VERSIONS
	// =========================
	".net5":   true,
	".net6":   true,
	".net7":   true,
	".net8":   true,
	"dotnet5": true,
	"dotnet6": true,
	"dotnet7": true,
	"dotnet8": true,

	// =========================
	// DATABASE MAJOR SHIFTS
	// =========================
	"mysql8":     true,
	"postgres14": true,
	"postgres15": true,
	"mongodb4":   true,
	"mongodb5":   true,

	// =========================
	// CLOUD CERTIFICATIONS
	// =========================
	"aws-saa-c02": true,
	"aws-saa-c03": true,
	"aws-dva-c01": true,
	"aws-mla-c01": true,
	"az-900":      true,
	"az-104":      true,
	"az-204":      true,
	"dp-900":      true,
	"dp-203":      true,
	"pl-300":      true,
	"pl-900":      true,
	"gcp-pca":     true,
	"gcp-ace":     true,

	// =========================
	// FINANCE / ACCOUNTING CERTIFICATIONS
	// =========================
	"ca-inter":   true,
	"ca-final":   true,
	"cfa-level1": true,
	"cfa-level2": true,
	"cfa-level3": true,
	"frm-level1": true,
	"frm-level2": true,

	// =========================
	// HEALTHCARE LICENSES
	// =========================
	"rn1":   true,
	"rn2":   true,
	"bls2":  true,
	"acls":  true,
	"pals":  true,
	"mbbs1": true,

	// =========================
	// TRADE / SAFETY CERTIFICATIONS
	// =========================
	"iosh":      true,
	"nebosh":    true,
	"osha10":    true,
	"osha30":    true,
	"iso9001":   true,
	"iso27001":  true,
	"iso14001":  true,
	"sixsigma":  true,
	"sixsigma6": true,
	"sixsigma3": true,

	// =========================
	// GOVERNMENT / EXAMS
	// =========================
	"ssc-cgl":    true,
	"ssc-chsl":   true,
	"ibps-po":    true,
	"ibps-clerk": true,
	"gate2023":   true,
	"gate2024":   true,
	"neet2023":   true,
	"neet2024":   true,

	// =========================
	// ERP SYSTEM VERSIONS (major)
	// =========================
	"sap-s4":     true,
	"sap-s4hana": true,

	// =========================
	// CAD / DESIGN
	// =========================
	"catia-v5": true,
	"nx12":     true,
	"nx10":     true,
}
