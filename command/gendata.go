package command

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "encoding/json"
    "github.com/codegangsta/cli"
)

func check_regexp(reg, str string) bool {
    return regexp.MustCompile(reg).Match([]byte(str))
}

func CmdGenData(c *cli.Context) {
	// Write your code here
    var fp *os.File
    var err error

    if c.NArg() < 1 {
        fp  = os.Stdin
    } else {
        fp, err = os.Open(c.Args()[0])
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    scanner := bufio.NewScanner(fp)
    m := make(map[string]string)
    mm := make(map[string]map[string]string)
    for scanner.Scan() {
        funcStr := ""
        var _ = funcStr

        check_regexp(`^func`, scanner.Text())
        // Instance method
        r := regexp.MustCompile(`^func \(.*\) (\w+)\(.*`)
        funcName := r.FindAllStringSubmatch(scanner.Text(), -1)
        if len(funcName) > 0 {
            if check_regexp(`^func`, scanner.Text()) == true {
                funcStr += scanner.Text() + "\n"
                for scanner.Scan() {
                    funcStr += scanner.Text() + "\n"
                    if check_regexp(`^}$`, scanner.Text()) == true {
                        m[funcName[0][1]] = funcStr
                        break
                    }
                }
            }
        }
        // normarl method
        r = regexp.MustCompile(`^func\s(\w+)\(.*`)
        funcName = r.FindAllStringSubmatch(scanner.Text(), -1)
        if len(funcName) > 0 {
            if check_regexp(`^func`, scanner.Text()) == true {
                funcStr += scanner.Text() + "\n"
                for scanner.Scan() {
                    funcStr += scanner.Text() + "\n"
                    if check_regexp(`^}$`, scanner.Text()) == true {
                        m[funcName[0][1]] = funcStr
                        break
                    }
                }
            }
        }
    }
    mm[c.Args()[0]] = m
    if err := scanner.Err(); err != nil {
        panic(err)
    }
    data, err := json.MarshalIndent(mm, "", "  ")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))

}
