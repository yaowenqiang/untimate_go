package example1

import (
    "net/http"
    "testing"
)
//go test -run Down -v

const succeed = "\u2713"
const failed = "\u2717"

func TestTableDownload(t *testing.T) {
    tests := []struct {
        url string
        statusCode int
    }{
        {"http://www.baidu.com", http.StatusOK},
        {"http://www.qq.com", http.StatusNotFound},
    }

    statusCode := 200
    t.Log("Given the need to test downloading content.")
    {
        for i, tt  := range tests {
            t.Logf("\tTest %d:\twhen checking %q for status code %d", i, tt.url, tt.statusCode)
            {
                resp, err := http.Get(tt.url)
                if err != nil {
                    t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
                }
                t.Logf("\t%s\tShould be able to make the Get call.", succeed)

                defer resp.Body.Close()

                if resp.StatusCode == tt.statusCode {
                    t.Logf("\t%s\tShould received a %d status code.", succeed, statusCode)
                } else {
                    t.Errorf("\t%s\tTest %d Should received %d status code : %d", failed, i, tt.statusCode, resp.StatusCode)
                }
            }

        }
    }
}
