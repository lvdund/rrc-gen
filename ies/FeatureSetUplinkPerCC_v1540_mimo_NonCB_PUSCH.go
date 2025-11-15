package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH struct {
	maxNumberSRS_ResourcePerSet         int64 `lb:1,ub:4,madatory`
	maxNumberSimultaneousSRS_ResourceTx int64 `lb:1,ub:4,madatory`
}

func (ie *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberSRS_ResourcePerSet, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSRS_ResourcePerSet", err)
	}
	if err = w.WriteInteger(ie.maxNumberSimultaneousSRS_ResourceTx, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSimultaneousSRS_ResourceTx", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberSRS_ResourcePerSet int64
	if tmp_int_maxNumberSRS_ResourcePerSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSRS_ResourcePerSet", err)
	}
	ie.maxNumberSRS_ResourcePerSet = tmp_int_maxNumberSRS_ResourcePerSet
	var tmp_int_maxNumberSimultaneousSRS_ResourceTx int64
	if tmp_int_maxNumberSimultaneousSRS_ResourceTx, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSimultaneousSRS_ResourceTx", err)
	}
	ie.maxNumberSimultaneousSRS_ResourceTx = tmp_int_maxNumberSimultaneousSRS_ResourceTx
	return nil
}
