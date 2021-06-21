package reader

import (
	"io/ioutil"
	"os"

	"howett.net/plist"
)

type PList map[string]interface{}

func ReadFromFile(path string) (PList, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	res := make(PList)

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	_, err = plist.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
