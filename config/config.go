package config

import (
	"os"
	"path/filepath"
)

// Executables used by the server.
type Executables struct {
	BIDS string `json:"bids"`
}

// Directories used by the server for temporary and long term storage.
type Directories struct {
	Temp   string `json:"temp"`
	Result string `json:"result"`
	Log    string `json:"log"`
}

// Denotations provide any freuquently used file names or other denotations
// e.g. validation result files, badge or result folder names.
type Denotations struct {
	LogFile              string `json:"logfile"`
	ResultsFolder        string `json:"resultsfolder"`
	ResultsFile          string `json:"resultsfile"`
	ResultsBadge         string `json:"resultsbadge"`
	ValidationConfigFile string `json:"valcfgfile"`
}

// Settings provide the default server settings.
// "Validators" currently only supports "BIDS".
type Settings struct {
	RootURL     string   `json:"rooturl"`
	Port        string   `json:"port"`
	LogSize     int      `json:"logsize"`
	GINUser     string   `json:"ginuser"`
	GINPassword string   `json:"ginpassword"`
	ClientID    string   `json:"clientid"`
	HookSecret  string   `json:"hooksecret"`
	CookieName  string   `json:"cookiename"`
	Validators  []string `json:"validators"`
}

// ServerCfg holds the config used to setup the gin validation server and
// the paths to all required executables, temporary and permanent folders.
type ServerCfg struct {
	Settings Settings    `json:"settings"`
	Exec     Executables `json:"executables"`
	Dir      Directories `json:"directories"`
	Label    Denotations `json:"denotations"`
}

var defaultCfg = ServerCfg{
	Settings{
		Port:        "3033",
		LogSize:     1048576,
		GINUser:     "ServiceWaiter",
		GINPassword: "",
		ClientID:    "gin-valid",
		HookSecret:  "",
		CookieName:  "gin-valid-session",
		// NOTE: NIX isn't actually supported yet, but having a second value helps with testing
		Validators: []string{"bids", "nix"},
	},
	Executables{
		BIDS: "bids-validator",
	},
	Directories{
		Temp:   filepath.Join(os.Getenv("GINVALIDHOME"), "tmp"),
		Log:    filepath.Join(os.Getenv("GINVALIDHOME"), "log"),
		Result: filepath.Join(os.Getenv("GINVALIDHOME"), "results"),
	},
	Denotations{
		LogFile:              "ginvalid.log",
		ResultsFolder:        "latest",
		ResultsFile:          "results.json",
		ResultsBadge:         "results.svg",
		ValidationConfigFile: "ginvalidation.yaml",
	},
}

// Read returns the default server configuration.
func Read() ServerCfg {
	return defaultCfg
}

// Set sets the server configuration.
func Set(cfg ServerCfg) {
	defaultCfg = cfg
}
