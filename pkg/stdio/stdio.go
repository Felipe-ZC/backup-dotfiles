package stdio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"example.com/dotfiles/pkg/domain"
)

type ioSvc struct {
	stdio domain.StdIOSvc
}

func NewIOSvc() domain.StdIOSvc {
	return ioSvc {
		//stdio: 	
	}
	//var stdio domain.StdIOSvc = ioSvc.stdio
	//return stdio 
}

func (stdio ioSvc) ReadCfgFile(filePath string) (domain.DotFileCfg, error) {
	var result domain.DotFileCfg
	var err error = readJsonFile(filePath, &result)
	return result, err
}

func (stdio ioSvc) ProcessCfgFile(cfgFile domain.DotFileCfg) error {
	return nil 
}

func readJsonFile(path string, dotCfg *domain.DotFileCfg) error {
	fmt.Println("reading...")
	// Open file for reading...
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	// Defer closing of file for later...
	defer jsonFile.Close()

	// Parse input file...
	byteValue, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &dotCfg)
	fmt.Println(dotCfg)
	
	return nil
}

func processConfigOpts(dotCfg domain.DotFileCfg) {
	fmt.Println(dotCfg)
    //for key, val := range dotCfg.dots {
        //log.Printf("Copying %s...", key)
        //switch val.(type) {
            //case string:
                //var destFileName string = filepath.Base(fmt.Sprint(val))
                //var dest string = fmt.Sprintf("%s/%s/%s", defaultDestDir, key, destFileName)
                //var err error = copyFileToDir(fmt.Sprint(val), dest)
                //if err != nil {
                    //log.Print(err)
                //}
            //case []interface{}:
                //log.Printf("%s is a []string", val)
        //}
    //}
}

func copyFileToDir(sourcePath string, destPath string) error {
    // Check if file exists...
    sourceFileStat, err := os.Stat(sourcePath)
    if err != nil {
        return fmt.Errorf("%s does not exist!", sourcePath)
    }
    // Check if file can be opened...
    if !sourceFileStat.Mode().IsRegular() {
        return fmt.Errorf("%s is not a regular file", sourcePath)
    }
    // Check if destination dir already exists...
    var destParentDir string = path.Dir(destPath)
    _, err = os.Stat(destParentDir)
    if err != nil {
        log.Printf("Creating %s...", destPath)
        err = os.MkdirAll(destParentDir, 0755)
        if err != nil {
            return fmt.Errorf("Could not create %s, error is: %s", destPath, err)
        }
    }
    // Read source file into a byte slice...
    source, err := ioutil.ReadFile(sourcePath)
    // Write file to destination...
    // 0644 is the octal representation of the file mode
    err = ioutil.WriteFile(destPath, source, 0644)
    if err != nil {
        return fmt.Errorf("Could not write %s to %s, error is: %s", sourcePath, destPath, err)
    }
    return nil
}
