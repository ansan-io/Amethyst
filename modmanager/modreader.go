package modmanager

import (
	"archive/zip"
	"fmt"
	"io"
)

type ModInfo struct {
	ModLoader     string `toml:"modLoader"`
	LoaderVersion string `toml:"loaderVersion"`
	License       string `toml:"license"`
	Mods          []struct {
		ModID       string `toml:"modId"`
		Version     string `toml:"version"`
		DisplayName string `toml:"displayName"`
		LogoFile    string `toml:"logoFile"`
		Credits     string `toml:"credits"`
		Authors     string `toml:"authors"`
		Description string `toml:"description"`
	} `toml:"mods"`
	Dependencies struct {
		mod []struct {
			ModID        string `toml:"modId"`
			Mandatory    bool   `toml:"mandatory"`
			VersionRange string `toml:"versionRange"`
			Ordering     string `toml:"ordering"`
			Side         string `toml:"side"`
		} `toml:"examplemod"`
	} `toml:"dependencies"`
}

func ReadModFile(target string) error {

  zf, err := zip.OpenReader(target)

  defer zf.Close();

  if err != nil {
    fmt.Println("Error opening zip file")
    return err
  }

  for _, file := range zf.File {
	  if file.Name == "META-INF/mods.toml" {
     	 content, err := file.Open();
      	defer content.Close();

      	if err != nil {
        	fmt.Println("Unable to read  content of file");
        	return err
      	}

     	 reader, err := io.ReadAll(content)
      	fmt.Println(string(reader))
	  	break
    }
  }
  return err
}

