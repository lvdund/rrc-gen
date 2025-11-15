package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplinkPerCC_v1540 struct {
	mimo_NonCB_PUSCH *FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH `lb:1,ub:4,optional`
}

func (ie *FeatureSetUplinkPerCC_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mimo_NonCB_PUSCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mimo_NonCB_PUSCH != nil {
		if err = ie.mimo_NonCB_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode mimo_NonCB_PUSCH", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_v1540) Decode(r *uper.UperReader) error {
	var err error
	var mimo_NonCB_PUSCHPresent bool
	if mimo_NonCB_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mimo_NonCB_PUSCHPresent {
		ie.mimo_NonCB_PUSCH = new(FeatureSetUplinkPerCC_v1540_mimo_NonCB_PUSCH)
		if err = ie.mimo_NonCB_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode mimo_NonCB_PUSCH", err)
		}
	}
	return nil
}
