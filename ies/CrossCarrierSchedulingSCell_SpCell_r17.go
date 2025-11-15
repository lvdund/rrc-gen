package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingSCell_SpCell_r17 struct {
	supportedSCS_Combinations_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17 `lb:1,ub:496,optional`
	pdcch_MonitoringOccasion_r17  CrossCarrierSchedulingSCell_SpCell_r17_pdcch_MonitoringOccasion_r17   `madatory`
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedSCS_Combinations_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedSCS_Combinations_r17 != nil {
		if err = ie.supportedSCS_Combinations_r17.Encode(w); err != nil {
			return utils.WrapError("Encode supportedSCS_Combinations_r17", err)
		}
	}
	if err = ie.pdcch_MonitoringOccasion_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_MonitoringOccasion_r17", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17) Decode(r *uper.UperReader) error {
	var err error
	var supportedSCS_Combinations_r17Present bool
	if supportedSCS_Combinations_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedSCS_Combinations_r17Present {
		ie.supportedSCS_Combinations_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17)
		if err = ie.supportedSCS_Combinations_r17.Decode(r); err != nil {
			return utils.WrapError("Decode supportedSCS_Combinations_r17", err)
		}
	}
	if err = ie.pdcch_MonitoringOccasion_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_MonitoringOccasion_r17", err)
	}
	return nil
}
