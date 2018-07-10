package config

// Executables used by the server.
type Executables struct {
	Gin  string
	BIDS string
}

// Directories used by the server for temporary and long term storage.
type Directories struct {
	Temp   string
	Result string
}

// Denotations provide any freuquently used file names or other denotations
// e.g. validation result files, badge or result folder names.
type Denotations struct {
	ResultsFolder string
	ResultsFile   string
	ResultsBadge  string
}

// ServerCfg holds the config used to setup the gin validation server and
// the paths to all required executables, temporary and permanent folders.
type ServerCfg struct {
	Exec  Executables
	Dir   Directories
	Label Denotations
}

var ginValidDefaultServer = ServerCfg{
	Executables{
		Gin:  "gin",
		BIDS: "/home/msonntag/node_modules/.bin/bids-validator",
	},
	Directories{
		Temp:   "/home/msonntag/Chaos/DL/val",
		Result: "/home/msonntag/Chaos/DL/valresults",
	},
	Denotations{
		ResultsFolder: "latest",
		ResultsFile:   "results.json",
		ResultsBadge:  "results.svg",
	},
}

// Read returns the default server configuration.
func Read() ServerCfg {
	return ginValidDefaultServer
}
