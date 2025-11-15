package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasQuantityResultsEUTRA struct {
	rsrp *RSRP_RangeEUTRA `optional`
	rsrq *RSRQ_RangeEUTRA `optional`
	sinr *SINR_RangeEUTRA `optional`
}

func (ie *MeasQuantityResultsEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rsrp != nil, ie.rsrq != nil, ie.sinr != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rsrp != nil {
		if err = ie.rsrp.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp", err)
		}
	}
	if ie.rsrq != nil {
		if err = ie.rsrq.Encode(w); err != nil {
			return utils.WrapError("Encode rsrq", err)
		}
	}
	if ie.sinr != nil {
		if err = ie.sinr.Encode(w); err != nil {
			return utils.WrapError("Encode sinr", err)
		}
	}
	return nil
}

func (ie *MeasQuantityResultsEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var rsrpPresent bool
	if rsrpPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrqPresent bool
	if rsrqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sinrPresent bool
	if sinrPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if rsrpPresent {
		ie.rsrp = new(RSRP_RangeEUTRA)
		if err = ie.rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp", err)
		}
	}
	if rsrqPresent {
		ie.rsrq = new(RSRQ_RangeEUTRA)
		if err = ie.rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode rsrq", err)
		}
	}
	if sinrPresent {
		ie.sinr = new(SINR_RangeEUTRA)
		if err = ie.sinr.Decode(r); err != nil {
			return utils.WrapError("Decode sinr", err)
		}
	}
	return nil
}
