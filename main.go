package main

import (
	"flag"
	"fmt"
	"example.com/dotfiles/pkg/backup"
	"example.com/dotfiles/pkg/stdio"
	//"example.com/dotfiles/pkg/domain"
)

const defaultDestDir string = "./test"

func main() {
    inputFilePath := parseFlags()
    if len(inputFilePath) == 0 {
        panic("No file path provided!")
    }
	ioSvc := stdio.NewIOSvc()
    backupSvc := backup.NewBackupSvc(ioSvc)
	configOpts, err := backupSvc.ReadCfgFile(inputFilePath)
	fmt.Println(configOpts, err)
}

//func processConfigOpts(configOpts map[string]interface{}) {
    //for key, val := range configOpts {
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
//}

//func copyFileToDir(sourcePath string, destPath string) error {
    //// Check if file exists...
    //sourceFileStat, err := os.Stat(sourcePath)
    //if err != nil {
        //return fmt.Errorf("%s does not exist!", sourcePath)
    //}
    //// Check if file can be opened...
    //if !sourceFileStat.Mode().IsRegular() {
        //return fmt.Errorf("%s is not a regular file", sourcePath)
    //}
    //// Check if destination dir already exists...
    //var destParentDir string = path.Dir(destPath)
    //_, err = os.Stat(destParentDir)
    //if err != nil {
        //log.Printf("Creating %s...", destPath)
        //err = os.MkdirAll(destParentDir, 0755)
        //if err != nil {
            //return fmt.Errorf("Could not create %s, error is: %s", destPath, err)
        //}
    //}
    //// Read source file into a byte slice...
    //source, err := ioutil.ReadFile(sourcePath)
    //// Write file to destination...
    //// 0644 is the octal representation of the file mode
    //err = ioutil.WriteFile(destPath, source, 0644)
    //if err != nil {
        //return fmt.Errorf("Could not write %s to %s, error is: %s", sourcePath, destPath, err)
    //}
    //return nil
//}

func parseFlags() string {
	var inputFilePath string
	flag.StringVar(&inputFilePath, "f", "", "Enter the full path to the config.json file")
	flag.Parse()
	return inputFilePath
}

//func readConfigFile(path string) map[string]interface{} {
	//// Open file for reading...
	//jsonFile, err := os.Open(path)
	//if err != nil {
		//log.Fatal(err)
	//}
	//// Defer closing of file for later...
	//defer jsonFile.Close()

	//// Parse input file...
	//byteValue, _ := ioutil.ReadAll(jsonFile)
	//var result map[string]interface{}
	//json.Unmarshal([]byte(byteValue), &result)

	//return result
//}
