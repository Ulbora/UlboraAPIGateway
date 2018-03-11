package handlers

import (
	"net/http"
	"testing"
)

func TestGateway_doPostPutPatch(t *testing.T) {
	var pPost passParams
	rtn := doPostPutPatch(&pPost)
	if rtn.rtnCode != http.StatusOK {
		t.Fail()
	}
}
