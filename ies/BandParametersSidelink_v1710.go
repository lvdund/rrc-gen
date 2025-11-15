package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelink_v1710 struct {
	tx_IUC_Scheme1_Mode2Sidelink_r17 *BandParametersSidelink_v1710_tx_IUC_Scheme1_Mode2Sidelink_r17 `optional`
	tx_IUC_Scheme2_Mode2Sidelink_r17 *BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17 `optional`
}

func (ie *BandParametersSidelink_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil, ie.tx_IUC_Scheme2_Mode2Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if ie.tx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
		if err = ie.tx_IUC_Scheme2_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelink_v1710) Decode(r *uper.UperReader) error {
	var err error
	var tx_IUC_Scheme1_Mode2Sidelink_r17Present bool
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_IUC_Scheme2_Mode2Sidelink_r17Present bool
	if tx_IUC_Scheme2_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present {
		ie.tx_IUC_Scheme1_Mode2Sidelink_r17 = new(BandParametersSidelink_v1710_tx_IUC_Scheme1_Mode2Sidelink_r17)
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if tx_IUC_Scheme2_Mode2Sidelink_r17Present {
		ie.tx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandParametersSidelink_v1710_tx_IUC_Scheme2_Mode2Sidelink_r17)
		if err = ie.tx_IUC_Scheme2_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}
