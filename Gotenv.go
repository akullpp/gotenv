package gotenv

import (
	"bufio"
	"os"
	"strings"
)

type dotenv map[string]string

func read(f string, d dotenv) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		l := s.Text()
		if i := strings.Index(l, "="); i >= 0 {
			if k := strings.TrimSpace(l[:i]); len(k) > 0 {
				v := ""

				if len(l) > i {
					v = strings.TrimSpace(l[i+1:])
				}
				d[k] = v
			}
		}
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}

// Get returns the key/value map or an error
func Get() (map[string]string, error) {
	// TODO: Environment-based env files
	d := dotenv{}
	if err := read(".env", d); err != nil {
		return d, err
	}
	if err := read(".env.local", d); err != nil {
		return d, err
	}
	return d, nil
}
