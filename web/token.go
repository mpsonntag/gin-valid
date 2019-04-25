package web

import (
	"encoding/base32"
	"encoding/gob"
	"os"
	"path/filepath"

	gweb "github.com/G-Node/gin-cli/web"
	"github.com/G-Node/gin-valid/config"
	"github.com/G-Node/gin-valid/log"
)

// saveToken writes a token to disk using the username as filename.
// The location is defined by config.Dir.Tokens.
func saveToken(ut gweb.UserToken) error {
	cfg := config.Read()
	filename := filepath.Join(cfg.Dir.Tokens, ut.Username)
	tokenfile, err := os.Create(filename)
	defer tokenfile.Close()
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(tokenfile)
	return encoder.Encode(ut)
}

// getTokenByUsername reads a token from disk using the username as filename.
// The location is defined by config.Dir.Tokens.
func getTokenByUsername(username string) (gweb.UserToken, error) {
	cfg := config.Read()
	filename := filepath.Join(cfg.Dir.Tokens, username)
	return loadToken(filename)
}

// loadToken loads a token from the provided path
func loadToken(path string) (gweb.UserToken, error) {
	ut := gweb.UserToken{}
	tokenfile, err := os.Open(path)
	if err != nil {
		log.Write("[Error] Failed to load token from %s", path)
		return ut, err
	}
	defer tokenfile.Close()

	decoder := gob.NewDecoder(tokenfile)
	err = decoder.Decode(&ut)
	return ut, err
}

// linkToSession links a sessionID to a user's token.
func linkToSession(username string, sessionid string) error {
	cfg := config.Read()
	tokendir, _ := filepath.Abs(cfg.Dir.Tokens)
	utfile := filepath.Join(tokendir, username)
	sidfile := filepath.Join(tokendir, "by-sessionid", b32(sessionid))
	return os.Symlink(utfile, sidfile)
}

// getTokenBySession loads a user's access token using the session ID found in
// the user's cookie store.
func getTokenBySession(sessionid string) (gweb.UserToken, error) {
	cfg := config.Read()
	tokendir := cfg.Dir.Tokens
	filename := filepath.Join(tokendir, "by-sessionid", b32(sessionid))
	return loadToken(filename)
}

// linkToRepo links a repository name to a user's token.
// This token will be used for cloning a repository to run a validator when a
// web hook is triggered.
func linkToRepo(username string, repopath string) error {
	cfg := config.Read()
	tokendir := cfg.Dir.Tokens
	utfile := filepath.Join(tokendir, username)
	sidfile := filepath.Join(tokendir, "by-repo", b32(repopath))
	return os.Symlink(utfile, sidfile)
}

// getTokenByRepo loads a user's access token using a repository path.
func getTokenByRepo(repopath string) (gweb.UserToken, error) {
	cfg := config.Read()
	tokendir := cfg.Dir.Tokens
	filename := filepath.Join(tokendir, "by-repo", b32(repopath))
	return loadToken(filename)
}

// rmTokenRepoLink deletes a repository -> token link, removing our ability to
// clone the repository.
func rmTokenRepoLink(repopath string) error {
	cfg := config.Read()
	tokendir := cfg.Dir.Tokens
	filename := filepath.Join(tokendir, "by-repo", b32(repopath))
	return os.Remove(filename)
}

// b32 encodes a string to base 32. Use this to make strings such as IDs or
// repopaths filename friendly.
func b32(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}