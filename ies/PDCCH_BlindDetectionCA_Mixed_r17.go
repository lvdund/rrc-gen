package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionCA_Mixed_r17 struct {
	pdcch_BlindDetectionCA1_r17 *int64 `lb:1,ub:15,optional`
	pdcch_BlindDetectionCA2_r17 *int64 `lb:1,ub:15,optional`
}

func (ie *PDCCH_BlindDetectionCA_Mixed_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcch_BlindDetectionCA1_r17 != nil, ie.pdcch_BlindDetectionCA2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_BlindDetectionCA1_r17 != nil {
		if err = w.WriteInteger(*ie.pdcch_BlindDetectionCA1_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA1_r17", err)
		}
	}
	if ie.pdcch_BlindDetectionCA2_r17 != nil {
		if err = w.WriteInteger(*ie.pdcch_BlindDetectionCA2_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA2_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_BlindDetectionCA_Mixed_r17) Decode(r *uper.UperReader) error {
	var err error
	var pdcch_BlindDetectionCA1_r17Present bool
	if pdcch_BlindDetectionCA1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCA2_r17Present bool
	if pdcch_BlindDetectionCA2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_BlindDetectionCA1_r17Present {
		var tmp_int_pdcch_BlindDetectionCA1_r17 int64
		if tmp_int_pdcch_BlindDetectionCA1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA1_r17", err)
		}
		ie.pdcch_BlindDetectionCA1_r17 = &tmp_int_pdcch_BlindDetectionCA1_r17
	}
	if pdcch_BlindDetectionCA2_r17Present {
		var tmp_int_pdcch_BlindDetectionCA2_r17 int64
		if tmp_int_pdcch_BlindDetectionCA2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA2_r17", err)
		}
		ie.pdcch_BlindDetectionCA2_r17 = &tmp_int_pdcch_BlindDetectionCA2_r17
	}
	return nil
}
