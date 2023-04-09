package email

import (
	"sort"

	"golang.org/x/exp/slices"
)

func findDateEmail(output map[int]DateEmail, d string) (DateEmail, int) {
	for i, v := range output {
		if v.Date == d {
			return v, i
		}
	}
	return DateEmail{Date: d}, -1
}

func GroupByDate(emails []Email) map[int]DateEmail {
	output := make(map[int]DateEmail)

	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		// slim the date down so that we are specific on the day, rather than
		// day/time
		d := m.Received.Format("Jan 02 2006")
		v, i := findDateEmail(output, d)
		v.Emails = append(v.Emails, m)

		if i == -1 {
			output[len(output)] = v
		} else {
			output[i] = v
		}
	}

	return output
}

// get the top 5 contacts in an email array
func GetTopContacts(emails []Email) []Email {
	output := make(map[string][]Email)
	// take the keys and put them in a key/value [email.From]: [email..]
	for _, m := range emails {
		output[m.From] = append(output[m.From], m)
	}

	keys := make([]string, 0, len(output))
	for key := range output {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return len(output[keys[i]]) > len(output[keys[j]]) })

	sliced := keys[0:5]
	var newEmails []Email

	for _, v := range sliced {
		idx := slices.IndexFunc(emails, func(c Email) bool { return c.From == v })
		newEmails = append(newEmails, emails[idx])
	}

	return newEmails
}

// https://stackoverflow.com/a/56141678
func filesize(s string) int {
	l := len(s)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if s[l-1] == '=' {
			eq++
		}
		if s[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	//
	// eq == 0 :    bits-wasted = 0
	// eq == 1 :    bits-wasted = 2
	// eq == 2 :    bits-wasted = 4

	// each base64 character = 6 bits

	// so orig length ==  (l*6 - eq*2) / 8

	return (l*3 - eq) / 4

}
