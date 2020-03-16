package server

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	pb "github.com/phil0522/taskd/proto"
	"gopkg.in/ini.v1"
)

const (
	repositoryPath = ".taskd"
)

type snipperManger struct {
	globalShellSnippets []*pb.ShellSnippet
}

func (s *snipperManger) initialize() {
	userRoot := os.Getenv("HOME")
	dirPath := path.Join(userRoot, repositoryPath)
	log.Printf("scan dir %s", dirPath)

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("walk file error %s", err.Error())
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ini") {
			return nil
		}

		log.Printf("loading %s\n", path)
		cfg, err := ini.LoadSources(ini.LoadOptions{
			AllowPythonMultilineValues: true,
		}, path)

		if err != nil {
			log.Printf("failed to load %s, error: %s", path, err.Error())
		}

		basePath, err := filepath.Rel(dirPath, path)
		basePath = strings.TrimSuffix(basePath, ".ini")
		for _, section := range cfg.Sections() {
			snippet := &pb.ShellSnippet{}
			snippet.SnippetName = basePath + "/" + section.Name()
			log.Printf("name: %s", snippet.SnippetName)
			snippet.SnippetDescription = section.Key("desc").MustString("")
			snippet.SnippetCommand = section.Key("cmd").MustString("")

			if snippet.SnippetCommand == "" {
				continue
			}
			s.globalShellSnippets = append(s.globalShellSnippets, snippet)
		}
		return nil
	})
}