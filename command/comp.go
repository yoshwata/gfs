package command

import (
    "github.com/codegangsta/cli"
    "fmt"
    "github.com/juju/utils/set"
    "io/ioutil"
    "encoding/json"
)

func jaccard(a, b set.Strings) float64 {
    return float64(a.Intersection(b).Size()) / float64(a.Union(b).Size())
}

func ngrams(s string, n int) set.Strings {
    var result = set.NewStrings(s)
    for i := 0; i < len(s)-n+1; i++ {
        result.Add(s[i : i+n])
    }
    return result
}

func CmdComp(c *cli.Context) {
    // Write your code here
    raw1, err := ioutil.ReadFile(c.Args()[0])
    if err != nil {
        panic(err)
    }

    raw2, err := ioutil.ReadFile(c.Args()[1])
    if err != nil {
        panic(err)
    }

    m1 := make(map[string]map[string]string)
    m2 := make(map[string]map[string]string)

    json.Unmarshal(raw1, &m1)
    json.Unmarshal(raw2, &m2)
    if err != nil {
        panic(err)
    }

    for k1, v1 := range m1 {
        for kk1, vv1 := range v1 {
            for k2, v2 := range m2 {
                for kk2, vv2 := range v2 {
                        fmt.Printf("%s %s %s %s %v\n", k1, kk1, k2, kk2, jaccard(ngrams(vv1, 3), ngrams(vv2, 3)))
                }
            }
        }
    }

}
