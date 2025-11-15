package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MultiDCI_MultiTRP_r16 struct {
	maxNumberCORESET_r16              MultiDCI_MultiTRP_r16_maxNumberCORESET_r16              `madatory`
	maxNumberCORESETPerPoolIndex_r16  int64                                                   `lb:1,ub:3,madatory`
	maxNumberUnicastPDSCH_PerPool_r16 MultiDCI_MultiTRP_r16_maxNumberUnicastPDSCH_PerPool_r16 `madatory`
}

func (ie *MultiDCI_MultiTRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberCORESET_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCORESET_r16", err)
	}
	if err = w.WriteInteger(ie.maxNumberCORESETPerPoolIndex_r16, &uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberCORESETPerPoolIndex_r16", err)
	}
	if err = ie.maxNumberUnicastPDSCH_PerPool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}

func (ie *MultiDCI_MultiTRP_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberCORESET_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCORESET_r16", err)
	}
	var tmp_int_maxNumberCORESETPerPoolIndex_r16 int64
	if tmp_int_maxNumberCORESETPerPoolIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberCORESETPerPoolIndex_r16", err)
	}
	ie.maxNumberCORESETPerPoolIndex_r16 = tmp_int_maxNumberCORESETPerPoolIndex_r16
	if err = ie.maxNumberUnicastPDSCH_PerPool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberUnicastPDSCH_PerPool_r16", err)
	}
	return nil
}
