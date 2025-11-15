package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_FormatConfig struct {
	interslotFrequencyHopping *PUCCH_FormatConfig_interslotFrequencyHopping `optional`
	additionalDMRS            *PUCCH_FormatConfig_additionalDMRS            `optional`
	maxCodeRate               *PUCCH_MaxCodeRate                            `optional`
	nrofSlots                 *PUCCH_FormatConfig_nrofSlots                 `optional`
	pi2BPSK                   *PUCCH_FormatConfig_pi2BPSK                   `optional`
	simultaneousHARQ_ACK_CSI  *PUCCH_FormatConfig_simultaneousHARQ_ACK_CSI  `optional`
}

func (ie *PUCCH_FormatConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.interslotFrequencyHopping != nil, ie.additionalDMRS != nil, ie.maxCodeRate != nil, ie.nrofSlots != nil, ie.pi2BPSK != nil, ie.simultaneousHARQ_ACK_CSI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interslotFrequencyHopping != nil {
		if err = ie.interslotFrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode interslotFrequencyHopping", err)
		}
	}
	if ie.additionalDMRS != nil {
		if err = ie.additionalDMRS.Encode(w); err != nil {
			return utils.WrapError("Encode additionalDMRS", err)
		}
	}
	if ie.maxCodeRate != nil {
		if err = ie.maxCodeRate.Encode(w); err != nil {
			return utils.WrapError("Encode maxCodeRate", err)
		}
	}
	if ie.nrofSlots != nil {
		if err = ie.nrofSlots.Encode(w); err != nil {
			return utils.WrapError("Encode nrofSlots", err)
		}
	}
	if ie.pi2BPSK != nil {
		if err = ie.pi2BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode pi2BPSK", err)
		}
	}
	if ie.simultaneousHARQ_ACK_CSI != nil {
		if err = ie.simultaneousHARQ_ACK_CSI.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousHARQ_ACK_CSI", err)
		}
	}
	return nil
}

func (ie *PUCCH_FormatConfig) Decode(r *uper.UperReader) error {
	var err error
	var interslotFrequencyHoppingPresent bool
	if interslotFrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalDMRSPresent bool
	if additionalDMRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCodeRatePresent bool
	if maxCodeRatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofSlotsPresent bool
	if nrofSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pi2BPSKPresent bool
	if pi2BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousHARQ_ACK_CSIPresent bool
	if simultaneousHARQ_ACK_CSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if interslotFrequencyHoppingPresent {
		ie.interslotFrequencyHopping = new(PUCCH_FormatConfig_interslotFrequencyHopping)
		if err = ie.interslotFrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode interslotFrequencyHopping", err)
		}
	}
	if additionalDMRSPresent {
		ie.additionalDMRS = new(PUCCH_FormatConfig_additionalDMRS)
		if err = ie.additionalDMRS.Decode(r); err != nil {
			return utils.WrapError("Decode additionalDMRS", err)
		}
	}
	if maxCodeRatePresent {
		ie.maxCodeRate = new(PUCCH_MaxCodeRate)
		if err = ie.maxCodeRate.Decode(r); err != nil {
			return utils.WrapError("Decode maxCodeRate", err)
		}
	}
	if nrofSlotsPresent {
		ie.nrofSlots = new(PUCCH_FormatConfig_nrofSlots)
		if err = ie.nrofSlots.Decode(r); err != nil {
			return utils.WrapError("Decode nrofSlots", err)
		}
	}
	if pi2BPSKPresent {
		ie.pi2BPSK = new(PUCCH_FormatConfig_pi2BPSK)
		if err = ie.pi2BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode pi2BPSK", err)
		}
	}
	if simultaneousHARQ_ACK_CSIPresent {
		ie.simultaneousHARQ_ACK_CSI = new(PUCCH_FormatConfig_simultaneousHARQ_ACK_CSI)
		if err = ie.simultaneousHARQ_ACK_CSI.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousHARQ_ACK_CSI", err)
		}
	}
	return nil
}
