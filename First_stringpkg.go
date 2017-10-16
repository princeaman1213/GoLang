package main
import (
	"fmt"
	"strings"
)
func main() {
	fmt.Println(
		strings.Contains("test","es"),"\n",
		strings.Count("test","e"),"\n",
		strings.HasPrefix("test","e"),"\n",  //true for t or te or tes or test
        strings.HasSuffix("test","st"),"\n",
		strings.Index("test","s"),"\n",
        strings.Join([]string{"a","b"}," and "),"\n",
		strings.Repeat("test",2),"\n",
		strings.Replace("ssss","s","r",1),"\n",
		strings.Split("doeton","d"),"\n",
		strings.ToLower("TEST"),"\n",
		strings.ToUpper("test"),

	)

}
