// Package initialism provides a list of alternative spellings, acronyms and initialisms for the named releasers.
//
// Alternative spellings are the same name but with different casing, spelling or punctuation (e.g. Coca-Cola, Coke).
//
// An acronym is an abbreviation formed from the initial letters of other words and pronounced as a word (e.g. NATO).
//
// An initialism is an abbreviation consisting of the initial letters pronounced invidivually (e.g. USA).
package initialism

import "strings"

// A Path is the partial URL path of the releaser.
type Path string

// List is a map of initialisms to releasers.
type List map[Path][]string

// path keys that are too long for the initialism map.
const (
	crue    = "cheat-requests-for-the-underground-elite"
	nappa   = "north-american-pirate_phreak-association"
	neua    = "national-elite-underground-alliance"
	nuaa    = "national-underground-application-alliance"
	tdttrsi = "the-dream-team*tristar-ampersand-red-sector-inc"
)

// all initialisms should be in their stylized form.
var initialisms = List{
	"2000ad":                                {"2KAD", "2000 AD"},
	"aces-of-ansi-art":                      {"AAA"},
	"acid-productions":                      {"ACiD", "ANSi Creators in Demand"},
	"advanced-art-of-cracking-group":        {"AAOCG"},
	"advanced-pirate-technology":            {"APT"},
	"affinity":                              {"AFT"},
	"air":                                   {"Team AiR", "Addiction In Releasing"},
	"alpha-flight":                          {"AFL"},
	"amnesia":                               {"AMN"},
	"anemia":                                {"ANM"},
	"arab-team-4-reverse-engineering":       {"AT4RE"},
	"arrogant-couriers-with-essays":         {"ACE"},
	"art-of-reverse-engineering":            {"AORE"},
	"artists-in-revolt":                     {"AiR"},
	"backlash":                              {"BLH"},
	"bad-ass-dudes":                         {"BAD"},
	"bentley-sidwell-productions":           {"BSP"},
	"blades-of-steel":                       {"BOS", "Blades"},
	"black-squadron":                        {"BS"},
	"blizzard":                              {"BLZ", "blizz"},
	"bitchin-ansi-design":                   {"BAD"},
	"boys-from-company-c":                   {"BCC"},
	"buds-biased-utils-report":              {"utils"},
	"canadian-pirates-inc":                  {"CPI"},
	"cd-images-for-the-elite":               {"CiFE"},
	"celerity-utilities-division":           {"CUD"},
	crue:                                    {"CRUE"},
	"chaos":                                 {"CHS"},
	"class":                                 {"CLS"},
	"classic":                               {"CLS"},
	"chemical-reaction":                     {"CRO"},
	"couriers-of-pirated-software":          {"COPS"},
	"courier-weektop-scorecard":             {"CWS"},
	"crackers-in-action":                    {"CIA"},
	"crack-in-morocco":                      {"CiM"},
	"creators-of-intense-art":               {"CIA"},
	"crude":                                 {"CRD"},
	"cybercrime-international-network":      {"CCi", "CyberCrime Inc."},
	"darksiders":                            {"DS"},
	"da-breaker-crew":                       {"DBC"},
	"damn-excellent-ansi-design":            {"DeAD"},
	"damn-excellent-ansi-designers":         {"DeAD"}, // Correct
	"delirium-tremens-group":                {"DTG"},
	"dead-on-arrival":                       {"DOA"},
	"dead-pirates-society":                  {"DPS"},
	"defacto":                               {"DF"},
	"defacto2":                              {"DF2"},
	"deviance":                              {"DEV", "DVN"},
	"devotion":                              {"DEV", "devot"},
	"digerati":                              {"DGT"},
	"digital-artists-of-the-rare-kind":      {"DARK"},
	"direct-from-stars":                     {"DFS"},
	"distinct":                              {"DTC", "DTN"},
	"divide-by-zero":                        {"DBZ"},
	"divine":                                {"DVN"},
	"drunken-rom-group":                     {"DRG", "Drunken"},
	"drink-or-die":                          {"DOD"},
	"dvt":                                   {"Devotion", "TeamDVT"},
	"dynasty":                               {"DYN"},
	"dynamix":                               {"DNX"},
	"dytec":                                 {"DYT", "DTC"},
	"dvniso":                                {"Deviance"},
	"eagle-soft-incorporated":               {"ESI"},
	"ebola-virus-crew":                      {"EVC"},
	"eclipse":                               {"ECL"},
	"eximius":                               {"XMS"},
	"empire-of-darkness":                    {"EOD"},
	"embrace":                               {"EMB"},
	"empire":                                {"EMP"},
	"energy":                                {"NRG"},
	"equinox":                               {"EQX"},
	"esp-pirates":                           {"ESP"},
	"fairlight":                             {"FLT"},
	"fairlight-dox":                         {"FDX", "FLTDOX"},
	"faith":                                 {"FTH"},
	"fantastic-4-cracking-group":            {"F4CG"},
	"fast-action-trading-elite":             {"fATE"},
	"fighting-for-fun":                      {"fff"},
	"fight-only-for-freedom":                {"FOFF"},
	"file-rappers":                          {"FR"},
	"flying-horse-cracking-force":           {"FHCF"},
	"future-crew":                           {"FC"},
	"future-brain-inc":                      {"FBi"},
	"fusion":                                {"FSN"},
	"genesis":                               {"GNS"},
	"graphic-revolution-in-progress":        {"GRiP"},
	"elite-couriers-group":                  {"ECG"},
	"epsilon":                               {"EPS"},
	"fyllecell":                             {"FLC"},
	"ghost-riders":                          {"GRS"},
	"graphics-rendered-in-magnificence":     {"GRiM"},
	"hard-to-beat-team":                     {"HTB"},
	"haze":                                  {"HZ"},
	"highroad":                              {"HR"},
	"hipe":                                  {"HPE"},
	"hoodlum":                               {"HLM"},
	"humble-dox":                            {"The Humble Guys DOX"},
	"hybrid":                                {"HYB"},
	"hype":                                  {"HYP"},
	"kalisto":                               {"KAL"},
	"kyrie-eleison":                         {"KE", "KEISO"},
	"illusion":                              {"iLL"},
	"independent":                           {"IND", "individual"},
	"independent-crackers-union":            {"ICU"},
	"insane-creators-enterprise":            {"iCE"},
	"international-network-of-crackers":     {"INC"},
	"international-cracking-crew":           {"iCC"},
	"inc-documentation-division":            {"IDD"},
	"inc-utility-division":                  {"IUD"},
	"influence":                             {"iNF"},
	"infinity":                              {"INF"},
	"jrp":                                   {"JRP"},
	"just-the-facts":                        {"JTF"},
	"laxity":                                {"LXT"},
	"legacy":                                {"LGC", "LGY"},
	"legend":                                {"LGD", "LEG"},
	"licensed-to-draw":                      {"LTD"},
	"light-speed-distributors":              {"LSD"},
	"lightforce":                            {"LFC", "LF"},
	"linezer0":                              {"Lz0", "Linezero"},
	"live-now-die-later":                    {"LnDL"},
	"lucid":                                 {"LCD"},
	"malicious-art-denomination":            {"MAD"},
	"malice":                                {"MAL"},
	"majic-12":                              {"M12"},
	"millennium":                            {"MnM"},
	"mutual-assured-destruction":            {"MAD"},
	"manifest":                              {"MFD", "Manifest Destiny"},
	"motiv8":                                {"M8"},
	"mirage":                                {"MIR"},
	"myth*deviance":                         {"MDVN"},
	neua:                                    {"NEUA", "North Eastern Underground Alliance"},
	nuaa:                                    {"NUAA"},
	"napalm":                                {"NPM"},
	"netrunners":                            {"NR"},
	"new-york-crackers":                     {"NYC"},
	"nexus":                                 {"NXS", "NX"},
	"nokturnal-trading-alliance":            {"NTA"},
	nappa:                                   {"NAPPA"},
	"nrp":                                   {"NRP"},
	"oddity":                                {"ODT"},
	"old-warez-inc":                         {"OWI"},
	"orion":                                 {"ORN"},
	"origin":                                {"OGN"},
	"originally-funny-guys":                 {"OFG"},
	"paradigm":                              {"PDM", "Zeus"},
	"paradox":                               {"PDX"},
	"pentagram":                             {"PTG"},
	"phoenix":                               {"PHX"},
	"phrozen-crew":                          {"PC"},
	"pirates-analyze-warez":                 {"PAW"},
	"pirates-gone-crazy":                    {"PGC"},
	"pirates-sick-of-initials":              {"PSi"},
	"pirates-with-attitudes":                {"PWA"},
	"prophecy":                              {"PCY"},
	"ptl-club":                              {"PTL"},
	"prestige":                              {"PSG", "PST"},
	"public-enemy":                          {"PE"},
	"public-enemy*red-sector-inc":           {"PE", "PE/RSI"},
	"razordox":                              {"RZR", "Razor"},
	"razor-1911":                            {"RZR", "Razor"},
	"razor-1911-cd-division":                {"RazorCD"},
	"razor-1911-demo":                       {"RZR", "Razor"},
	"reality-check-network":                 {"RCN"},
	"rebels":                                {"RBS"},
	"red-sector-inc":                        {"RSI"},
	"release-on-rampage":                    {"RoR"},
	"reloaded":                              {"RLD"},
	"reflux":                                {"RLX"},
	"relentless-pursuit-of-magnificence":    {"RPM"},
	"request-to-send":                       {"RTS"},
	"resistance-is-futile":                  {"RiF"},
	"resurrection":                          {"RSR", "RES"},
	"revenge-crew":                          {"REV"},
	"reverse-2-revolutionize":               {"R2R"},
	"reverse-engineers-dream":               {"RED"},
	"reverse-engineering-in-software":       {"REiS"},
	"reverse-engineering-passion-team":      {"REPT"},
	"rise-in-superior-couriering":           {"RiSC"},
	"seek-n-destroy":                        {"SND", "Seek and Destroy"},
	"skid-row":                              {"SR", "Skidrow"},
	"scoopex":                               {"SCX", "SPX"},
	"scandal":                               {"SCL"},
	"scienide":                              {"SCi"},
	"silicon-dream-artists":                 {"SDA"},
	"share-and-enjoy":                       {"SAE"},
	"shitonlygerman":                        {"SOG", "Scheisse Deutsch Only"},
	"skillion":                              {"SKN"},
	"sodom":                                 {"SDM"},
	"software-chronicles-digest":            {"SCD"},
	"software-pirates-inc":                  {"SPI"},
	"superior-art-creations":                {"SAC"},
	"surprise-productions":                  {"SP"},
	"team-technotrogens":                    {"TT3", "Team T3"},
	"the-crazed-asylum":                     {"TCA"},
	"the-console-division":                  {"TCD"},
	"the-council":                           {"CNC"},
	"the-dream-team":                        {"TDT"},
	"the-dream-team*skid-row":               {"TDT/SR"},
	tdttrsi:                                 {"TDT/TRSi"},
	"the-firm":                              {"FiRM", "FRM"},
	"the-force-team":                        {"TFT"},
	"the-game-review":                       {"TGR"},
	"the-game-scene-chart":                  {"TGSC"},
	"the-grand-council":                     {"TGC"},
	"the-humble-guys":                       {"THG", "Humble"},
	"the-lamerz-group":                      {"TLG"},
	"the-millennium-group":                  {"TMG"},
	"the-net-monkey-weekly-report":          {"TNMWR"},
	"the-nova-team":                         {"TNT"},
	"the-one-and-only":                      {"TOAO"},
	"the-outlaws":                           {"TOL", "OL"},
	"the-players-club":                      {"TPC"},
	"the-reversers-ultimate-epidemic":       {"tRUE"},
	"the-reviewers-guild":                   {"TRG"},
	"the-sabotage-rebellion-hackers":        {"TSRh"},
	"the-software-innovation-network":       {"SIN"},
	"the-sysops-association-network":        {"TSAN"},
	"the-unbiased-dox-report":               {"DR"},
	"the-underground-council":               {"UGC"},
	"the-untouchables":                      {"UNT"},
	"thg-fx":                                {"The Humble Guys FX"},
	"tristar":                               {"TRS"},
	"tristar-ampersand-red-sector-inc":      {"TRSi", "TRS", "Tristar"},
	"tyranny":                               {"TYR"},
	"ultra-tech":                            {"UT"},
	"under-seh-team":                        {"UST"},
	"union":                                 {"UNi"},
	"united-artist-association":             {"UAA"},
	"united-couriers":                       {"UC"},
	"united-cracking-force":                 {"UCF"},
	"united-group-international":            {"UGI"},
	"united-reverse-engineering-team":       {"URET"},
	"united-software-association*fairlight": {"USA/FLT"},
	"united-software-association":           {"USA"},
	"united-states-courier-report":          {"USCR"},
	"underpl":                               {"UPL"},
	"unleashed":                             {"UNL"},
	"untouchables":                          {"UNT"},
	"vengeance":                             {"VGN", "VEN"},
	"visions-of-reality":                    {"VOR"},
	"virility":                              {"VRL"},
	"wave":                                  {"The Wave", "CNC"},
	"wicked":                                {"WKD"},
	"x_force":                               {"XF"},
	"xtreeme":                               {"XT"},
	"zero-waiting-time":                     {"ZWT"},
	"zone":                                  {"z0ne"},
}

// Initialism returns the alternative spellings, acronyms and initialisms for the URL path.
// Or an empty slice if the URL path has no initialism.
//
// Example:
//
//	Initialism("the-firm") = []string{"FiRM, FRM"}
//	Initialism("defacto2") = []string{"DF2"}
func Initialism(path Path) []string {
	return initialisms[path]
}

// Initialisms returns the list of initialisms.
func Initialisms() List {
	return initialisms
}

// IsInitialism returns true if the URL path has an initialism.
//
// Example:
//
//	IsInitialism("the-firm") = true
//	IsInitialism("defacto2") = true
//	IsInitialism("some-random-bbs") = false
func IsInitialism(path Path) bool {
	_, ok := initialisms[path]
	return ok
}

// Join returns the alternative spellings, acronyms and initialisms for the
// URL path as a comma separated string.
// Or an empty string if the URL path has no initialism.
//
// Example:
//
//	Join("the-firm") = "FiRM, FRM"
//	Join("defacto2") = "DF2"
func Join(path Path) string {
	i := Initialism(path)
	if len(i) == 0 {
		return ""
	}
	return strings.Join(i, ", ")
}
