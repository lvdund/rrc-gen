package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n8  uper.Enumerated = 0
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n16 uper.Enumerated = 1
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n32 uper.Enumerated = 2
	DummyG_maxNumberSSB_CSI_RS_ResourceOneTx_Enum_n64 uper.Enumerated = 3
)

type DummyG_maxNumberSSB_CSI_RS_ResourceOneTx struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *DummyG_maxNumberSSB_CSI_RS_ResourceOneTx) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode DummyG_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	return nil
}

func (ie *DummyG_maxNumberSSB_CSI_RS_ResourceOneTx) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode DummyG_maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
