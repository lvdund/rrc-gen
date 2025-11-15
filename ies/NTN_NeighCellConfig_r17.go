package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NTN_NeighCellConfig_r17 struct {
	ntn_Config_r17  *NTN_Config_r17 `optional`
	carrierFreq_r17 *ARFCN_ValueNR  `optional`
	physCellId_r17  *PhysCellId     `optional`
}

func (ie *NTN_NeighCellConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ntn_Config_r17 != nil, ie.carrierFreq_r17 != nil, ie.physCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ntn_Config_r17 != nil {
		if err = ie.ntn_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ntn_Config_r17", err)
		}
	}
	if ie.carrierFreq_r17 != nil {
		if err = ie.carrierFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode carrierFreq_r17", err)
		}
	}
	if ie.physCellId_r17 != nil {
		if err = ie.physCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode physCellId_r17", err)
		}
	}
	return nil
}

func (ie *NTN_NeighCellConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var ntn_Config_r17Present bool
	if ntn_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var carrierFreq_r17Present bool
	if carrierFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var physCellId_r17Present bool
	if physCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ntn_Config_r17Present {
		ie.ntn_Config_r17 = new(NTN_Config_r17)
		if err = ie.ntn_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ntn_Config_r17", err)
		}
	}
	if carrierFreq_r17Present {
		ie.carrierFreq_r17 = new(ARFCN_ValueNR)
		if err = ie.carrierFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode carrierFreq_r17", err)
		}
	}
	if physCellId_r17Present {
		ie.physCellId_r17 = new(PhysCellId)
		if err = ie.physCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode physCellId_r17", err)
		}
	}
	return nil
}
