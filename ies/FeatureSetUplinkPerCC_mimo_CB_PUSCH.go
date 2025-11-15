package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_mimo_CB_PUSCH struct {
	maxNumberMIMO_LayersCB_PUSCH *MIMO_LayersUL `optional`
	maxNumberSRS_ResourcePerSet  int64          `lb:1,ub:2,madatory`
}

func (ie *FeatureSetUplinkPerCC_mimo_CB_PUSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberMIMO_LayersCB_PUSCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberMIMO_LayersCB_PUSCH != nil {
		if err = ie.maxNumberMIMO_LayersCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberMIMO_LayersCB_PUSCH", err)
		}
	}
	if err = w.WriteInteger(ie.maxNumberSRS_ResourcePerSet, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSRS_ResourcePerSet", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_mimo_CB_PUSCH) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberMIMO_LayersCB_PUSCHPresent bool
	if maxNumberMIMO_LayersCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberMIMO_LayersCB_PUSCHPresent {
		ie.maxNumberMIMO_LayersCB_PUSCH = new(MIMO_LayersUL)
		if err = ie.maxNumberMIMO_LayersCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberMIMO_LayersCB_PUSCH", err)
		}
	}
	var tmp_int_maxNumberSRS_ResourcePerSet int64
	if tmp_int_maxNumberSRS_ResourcePerSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSRS_ResourcePerSet", err)
	}
	ie.maxNumberSRS_ResourcePerSet = tmp_int_maxNumberSRS_ResourcePerSet
	return nil
}
