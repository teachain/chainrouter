package manager

import "testing"

func TestToPureName(t *testing.T){
  name:="payment.fisco.helloworld"
  t.Log(toPureName(name))
}

