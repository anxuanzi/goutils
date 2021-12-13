package ftacodec

import "testing"

func TestAes_Encrypt(t *testing.T) {
	res, err := NewAes("QZ8j7zVImYd6cxWs", "QZ8j7zVImYd6cxWs").Encrypt("anxuanzi")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(res)
}

func TestAes_Decrypt(t *testing.T) {
	res, err := NewAes("QZ8j7zVImYd6cxWs", "QZ8j7zVImYd6cxWs").Decrypt("sDl3HwiEGurUU3LhywonCQ==")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log(res)
}
