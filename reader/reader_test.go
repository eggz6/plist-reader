package reader

import "testing"

func TestReaderFromFile(t *testing.T) {
	path := "Info.plist"

	res, err := ReadFromFile(path)
	if err != nil {
		t.Fatalf("read from file faile. err=%v", err)
	}

	t.Logf("read from file success. res=%v", res)
}
