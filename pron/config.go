package pron

import "os"

func parseConfig(location string) (tab []job, err error) {
	file, err := os.Open(location)
	defer file.Close()
	if err != nil {
		return []job{}, err
	}
	return nil, nil
}

func parseLine(s string) (j job, err error) {

}
