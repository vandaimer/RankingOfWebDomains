package customerimporter

import "testing"

func TestSortByDomain(t *testing.T) {
        domains := map[string]int{
                "teamwork.com": 4,
                "bbb.com":      5,
                "aaa.com":      3,
        }

        expected := []Domain{
                {"aaa.com", 3},
                {"bbb.com", 3},
                {"teamwork.com", 3},
        }

        result := sortByDomain(domains)

        for index, item := range result {
                if domain := expected[index]; domain.Name != item.Name {
                        t.FailNow()
                }
        }

}

func TestReturnDomainFromEmail(t *testing.T) {
        domain := "domain.com"
        email := "my-email@" + domain

        expected := domain

        if result, _ := getDomainFromEmail(email); result != expected {
                t.FailNow()
        }
}

func TestReturnErrorWhenPassInvalidEmailToGetDomainFromEmail(t *testing.T) {
        email := "invalidemail"

        if _, err := getDomainFromEmail(email); err == nil {
                t.FailNow()
        }
}

func TestCountOccurrencesForEachDomain(t *testing.T) {
	domains := []string{
                "linkedin.com",
                "teamwork.com",
                "teamwork.com",
        }

        expected := map[string]int{
                "linkedin.com": 1,
                "teamwork.com": 2,
        }

        result := countOccurrencesForEachDomain(domains)

        for key, value := range result {
                if expected[key] != value {
                        t.FailNow()
                }
        }
}
