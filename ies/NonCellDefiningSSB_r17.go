package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NonCellDefiningSSB_r17 struct {
	absoluteFrequencySSB_r17 ARFCN_ValueNR                               `madatory`
	ssb_Periodicity_r17      *NonCellDefiningSSB_r17_ssb_Periodicity_r17 `optional`
	ssb_TimeOffset_r17       *NonCellDefiningSSB_r17_ssb_TimeOffset_r17  `optional`
}

func (ie *NonCellDefiningSSB_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_Periodicity_r17 != nil, ie.ssb_TimeOffset_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.absoluteFrequencySSB_r17.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteFrequencySSB_r17", err)
	}
	if ie.ssb_Periodicity_r17 != nil {
		if err = ie.ssb_Periodicity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_Periodicity_r17", err)
		}
	}
	if ie.ssb_TimeOffset_r17 != nil {
		if err = ie.ssb_TimeOffset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_TimeOffset_r17", err)
		}
	}
	return nil
}

func (ie *NonCellDefiningSSB_r17) Decode(r *uper.UperReader) error {
	var err error
	var ssb_Periodicity_r17Present bool
	if ssb_Periodicity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_TimeOffset_r17Present bool
	if ssb_TimeOffset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.absoluteFrequencySSB_r17.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteFrequencySSB_r17", err)
	}
	if ssb_Periodicity_r17Present {
		ie.ssb_Periodicity_r17 = new(NonCellDefiningSSB_r17_ssb_Periodicity_r17)
		if err = ie.ssb_Periodicity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Periodicity_r17", err)
		}
	}
	if ssb_TimeOffset_r17Present {
		ie.ssb_TimeOffset_r17 = new(NonCellDefiningSSB_r17_ssb_TimeOffset_r17)
		if err = ie.ssb_TimeOffset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_TimeOffset_r17", err)
		}
	}
	return nil
}
