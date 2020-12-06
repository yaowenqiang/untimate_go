package example1

import (
    "net/http"
    "testing"
)
//go test -run Down -v

const succeed = "\u2713"
const failed = "\u2717"

func TestDownload(t *testing.T) {
    url := "https://www.baidu.com/"
    statusCode := 200
    t.Log("Given the need to test downloading content.")
    {
        testID := 0
        t.Logf("\tTest 0:\twhen checking %q for status code %d", url, statusCode)
        {
            resp, err := http.Get(url)
            if err != nil {
                t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
            }
            t.Logf("\t%s\tShould be able to make the Get call.", succeed)

            defer resp.Body.Close()

            if resp.StatusCode == statusCode {
                t.Logf("\t%s\tShould received a %d status code.", succeed, statusCode)
            } else {
                t.Errorf("\t%s\tTest %d Should received %d status code : %d", failed, testID, statusCode, resp.StatusCode)
            }
        }
    }
}
