package sortimports
import (
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)
func SortInDir(path string) error{
	files, err :=  ioutil.ReadDir(path)
	if err!=nil{
		return err
	} else if len(files)==0{
		return errors.New("go file doesn't exist")
	}
	for _,v := range files {
		err := SortImports(path+"/"+v.Name())
		if err!=nil{
			return err
		}
	}
	return nil
}
func SortImports(path string) error{
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var start,end int
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			start=i+1
		} else if strings.Contains(line, ")"){
			end=i
			break
		}
	}
	sort.Strings(lines[start:end])
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
