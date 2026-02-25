package helpers

var SkillAlias = map[string]string{

	// ===== GO =====
	"golang":  "go",
	"go-lang": "go",
	"go lang": "go",

	// ===== JAVASCRIPT =====
	"js":         "javascript",
	"node":       "nodejs",
	"node.js":    "nodejs",
	"node-js":    "nodejs",
	"reactjs":    "react",
	"react.js":   "react",
	"angularjs":  "angular",
	"angular.js": "angular",
	"vuejs":      "vue",
	"vue.js":     "vue",
	"nextjs":     "next",
	"next.js":    "next",
	"nuxtjs":     "nuxt",
	"nuxt.js":    "nuxt",
	"expressjs":  "express",
	"express.js": "express",

	// ===== TYPESCRIPT =====
	"ts": "typescript",

	// ===== PYTHON =====
	"py": "python",

	// ===== JAVA =====
	"springboot":  "spring",
	"spring-boot": "spring",
	"java 8":      "java",
	"java 11":     "java",
	"java 17":     "java",

	// ===== .NET =====
	".net":      "dotnet",
	".net core": "dotnet",
	"dot net":   "dotnet",
	"c#":        "csharp",
	"c sharp":   "csharp",
	"asp.net":   "aspnet",

	// ===== C / C++ =====
	"c++":         "cpp",
	"c plus plus": "cpp",

	// ===== CLOUD =====
	"amazon web services": "aws",
	"ec2":                 "aws",
	"s3":                  "aws",
	"lambda":              "aws",
	"gcloud":              "gcp",
	"google cloud":        "gcp",
	"microsoft azure":     "azure",

	// ===== DEVOPS =====
	"k8":                     "kubernetes",
	"k8s":                    "kubernetes",
	"docker swarm":           "docker",
	"helm charts":            "helm",
	"ci/cd":                  "cicd",
	"ci-cd":                  "cicd",
	"continuous integration": "cicd",
	"continuous delivery":    "cicd",

	// ===== DATABASES =====
	"postgres":   "postgresql",
	"mongo":      "mongodb",
	"ms sql":     "mssql",
	"sql server": "mssql",
	"oracle db":  "oracle",

	// ===== DATA =====
	"pyspark":          "spark",
	"apache spark":     "spark",
	"hadoop ecosystem": "hadoop",

	// ===== MOBILE =====
	"react native": "reactnative",
	"flutter sdk":  "flutter",
	"android sdk":  "android",
	"ios sdk":      "ios",

	// ===== TOOLS =====
	"gitlab ci":           "gitlab",
	"github actions":      "github",
	"bitbucket pipelines": "bitbucket",

	// ===== AI / ML =====
	"machine learning":        "ml",
	"deep learning":           "dl",
	"artificial intelligence": "ai",
	"tensorflow 2":            "tensorflow",
	"pytorch lightning":       "pytorch",

	// ===== TESTING =====
	"unit testing":        "testing",
	"integration testing": "testing",
	"jestjs":              "jest",
	"mocha js":            "mocha",

	// ===== COMMON CLEANUPS =====
	"rest api":       "rest",
	"restful api":    "rest",
	"micro services": "microservices",

	// ===== WHITE COLLAR SKILLS =====
	// Management & Leadership
	"team management":   "management",
	"people management": "management",
	"team leadership":   "leadership",
	"project mgmt":      "projectmanagement",
	"project manager":   "projectmanagement",
	"pm":                "projectmanagement",
	"scrum master":      "scrummaster",
	"agile coach":       "agilecoach",
	"change management": "changemanagement",

	// Business & Strategy
	"business development":             "businessdevelopment",
	"biz dev":                          "businessdevelopment",
	"strategic planning":               "strategy",
	"business analysis":                "businessanalysis",
	"business analyst":                 "businessanalysis",
	"market research":                  "marketresearch",
	"competitive analysis":             "competitiveanalysis",
	"customer relationship management": "crm",
	"salesforce crm":                   "salesforce",
	"ms excel":                         "excel",
	"microsoft excel":                  "excel",
	"ms word":                          "word",
	"microsoft word":                   "word",
	"ms powerpoint":                    "powerpoint",
	"microsoft powerpoint":             "powerpoint",
	"power bi":                         "powerbi",
	"tableau software":                 "tableau",

	// Finance & Accounting
	"financial planning":   "finance",
	"financial analysis":   "finance",
	"bookkeeping":          "accounting",
	"accounts payable":     "accounting",
	"accounts receivable":  "accounting",
	"financial modeling":   "financialmodeling",
	"budgeting":            "budgeting",
	"forecasting":          "forecasting",
	"general ledger":       "accounting",
	"tally erp":            "tally",
	"gst filing":           "gst",
	"taxation":             "tax",
	"chartered accountant": "ca",

	// Sales & Marketing
	"digital marketing":          "digitalmarketing",
	"social media marketing":     "socialmarketing",
	"content marketing":          "contentmarketing",
	"email marketing":            "emailmarketing",
	"seo":                        "seo",
	"sem":                        "sem",
	"pay per click":              "ppc",
	"google ads":                 "googleads",
	"facebook ads":               "facebookads",
	"sales development":          "sales",
	"inside sales":               "sales",
	"outside sales":              "sales",
	"account management":         "accountmanagement",
	"customer success":           "customersuccess",
	"crm":                        "crm",
	"salesforce":                 "salesforce",
	"search engine optimization": "seo",
	"search engine marketing":    "sem",
	"google analytics":           "analytics",

	// HR & Administration
	"human resources":         "hr",
	"talent acquisition":      "recruiting",
	"talent management":       "hr",
	"employee relations":      "hr",
	"compensation":            "compensation",
	"benefits administration": "benefits",
	"performance management":  "hr",

	// Legal & Compliance
	"contract management":   "contracts",
	"legal research":        "legal",
	"compliance":            "compliance",
	"risk management":       "riskmanagement",
	"intellectual property": "ip",

	// Operations
	"supply chain":            "supplychain",
	"logistics":               "logistics",
	"procurement":             "procurement",
	"vendor management":       "vendormanagement",
	"quality assurance":       "qa",
	"process improvement":     "processimprovement",
	"lean six sigma":          "leansixsigma",
	"six sigma":               "sixsigma",
	"project management":      "pm",
	"program management":      "pm",
	"operations management":   "operations",
	"supply chain management": "supplychain",
	"logistics management":    "logistics",

	// Communication & Soft Skills
	"public speaking":        "publicspeaking",
	"presentation skills":    "presentations",
	"written communication":  "writing",
	"technical writing":      "technicalwriting",
	"copywriting":            "copywriting",
	"stakeholder management": "stakeholdermanagement",
	"client relations":       "clientrelations",

	// Healthcare
	"registered nurse":            "nurse",
	"staff nurse":                 "nurse",
	"medical doctor":              "doctor",
	"general physician":           "doctor",
	"certified nursing assistant": "cna",
	"pharmacy technician":         "pharmacist",
	"healthcare assistant":        "healthcare",

	// Education
	"primary teacher":     "teacher",
	"secondary teacher":   "teacher",
	"assistant professor": "professor",
	"lecturer":            "teacher",

	// ===== BLUE COLLAR SKILLS =====
	// Construction & Trades
	"construction":             "construction",
	"general contractor":       "construction",
	"carpentry":                "carpentry",
	"framing":                  "carpentry",
	"cabinet making":           "carpentry",
	"electrical work":          "electrical",
	"electrician":              "electrical",
	"wiring":                   "electrical",
	"plumbing":                 "plumbing",
	"pipe fitting":             "plumbing",
	"hvac":                     "hvac",
	"heating":                  "hvac",
	"ventilation":              "hvac",
	"air conditioning":         "hvac",
	"roofing":                  "roofing",
	"flooring":                 "flooring",
	"tiling":                   "tiling",
	"painting":                 "painting",
	"drywall":                  "drywall",
	"masonry":                  "masonry",
	"concrete work":            "concrete",
	"welding":                  "welding",
	"arc welding":              "welding",
	"mig welding":              "welding",
	"tig welding":              "welding",
	"electrician technician":   "electrician",
	"industrial electrician":   "electrician",
	"welder fabricator":        "welder",
	"mason worker":             "mason",
	"heavy equipment operator": "operator",
	"forklift operator":        "operator",
	"crane operator":           "operator",

	// Manufacturing & Production
	"manufacturing":        "manufacturing",
	"production":           "production",
	"assembly line":        "assembly",
	"quality control":      "qualitycontrol",
	"machine operation":    "machineoperation",
	"cnc operation":        "cnc",
	"cnc programming":      "cnc",
	"forklift operation":   "forklift",
	"warehouse operations": "warehouse",
	"inventory management": "inventory",
	"shipping receiving":   "shipping",
	"packaging":            "packaging",
	"machine operator":     "operator",
	"production operator":  "operator",
	"assembly line worker": "assembly",
	"quality inspector":    "qualitycontrol",

	// Automotive & Mechanical
	"automotive repair":   "automotive",
	"auto mechanic":       "automotive",
	"engine repair":       "automotive",
	"brake repair":        "automotive",
	"transmission repair": "automotive",
	"diesel mechanic":     "diesel",
	"heavy equipment":     "heavyequipment",
	"mechanical repair":   "mechanical",
	"automobile mechanic": "mechanic",
	"car mechanic":        "mechanic",

	// Food Service & Hospitality
	"food preparation":    "foodprep",
	"cooking":             "cooking",
	"kitchen management":  "kitchen",
	"food safety":         "foodsafety",
	"restaurant service":  "restaurant",
	"bartending":          "bartending",
	"customer service":    "customerservice",
	"hotel management":    "hospitality",
	"chef de partie":      "chef",
	"sous chef":           "chef",
	"kitchen staff":       "cook",
	"wait staff":          "waiter",
	"food service worker": "waiter",

	// Transportation & Logistics
	"truck driving":       "trucking",
	"commercial driving":  "commercialdriving",
	"cdl":                 "cdl",
	"delivery":            "delivery",
	"freight management":  "freight",
	"truck driver":        "driver",
	"delivery driver":     "driver",
	"commercial driver":   "driver",
	"warehouse worker":    "warehouse",
	"warehouse associate": "warehouse",

	// Retail
	"store manager":                   "retail",
	"retail sales associate":          "retail",
	"cashier operator":                "cashier",
	"customer service representative": "customerservice",
	"call center executive":           "callcenter",
	"bpo executive":                   "bpo",

	// Maintenance & Facilities
	"facility maintenance": "maintenance",
	"building maintenance": "maintenance",
	"grounds keeping":      "groundskeeping",
	"janitorial":           "janitorial",
	"custodial":            "custodial",

	// Safety & Security
	"occupational safety": "safety",
	"workplace safety":    "safety",
	"security":            "security",
	"loss prevention":     "lossprevention",
	"security officer":    "security",
	"security guard":      "security",

	// Cleaning
	"house keeping":       "housekeeping",
	"janitorial services": "janitor",
}
