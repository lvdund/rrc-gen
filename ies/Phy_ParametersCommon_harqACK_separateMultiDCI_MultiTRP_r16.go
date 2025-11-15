package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16 struct {
	maxNumberLongPUCCHs_r16 *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16 `optional`
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberLongPUCCHs_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberLongPUCCHs_r16 != nil {
		if err = ie.maxNumberLongPUCCHs_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberLongPUCCHs_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberLongPUCCHs_r16Present bool
	if maxNumberLongPUCCHs_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberLongPUCCHs_r16Present {
		ie.maxNumberLongPUCCHs_r16 = new(Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16_maxNumberLongPUCCHs_r16)
		if err = ie.maxNumberLongPUCCHs_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberLongPUCCHs_r16", err)
		}
	}
	return nil
}
