package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkTransformPrecoding_r16 struct {
	pi2BPSK_ScramblingID0 *int64 `lb:0,ub:65535,optional`
	pi2BPSK_ScramblingID1 *int64 `lb:0,ub:65535,optional`
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pi2BPSK_ScramblingID0 != nil, ie.pi2BPSK_ScramblingID1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pi2BPSK_ScramblingID0 != nil {
		if err = w.WriteInteger(*ie.pi2BPSK_ScramblingID0, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode pi2BPSK_ScramblingID0", err)
		}
	}
	if ie.pi2BPSK_ScramblingID1 != nil {
		if err = w.WriteInteger(*ie.pi2BPSK_ScramblingID1, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode pi2BPSK_ScramblingID1", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkTransformPrecoding_r16) Decode(r *uper.UperReader) error {
	var err error
	var pi2BPSK_ScramblingID0Present bool
	if pi2BPSK_ScramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pi2BPSK_ScramblingID1Present bool
	if pi2BPSK_ScramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pi2BPSK_ScramblingID0Present {
		var tmp_int_pi2BPSK_ScramblingID0 int64
		if tmp_int_pi2BPSK_ScramblingID0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode pi2BPSK_ScramblingID0", err)
		}
		ie.pi2BPSK_ScramblingID0 = &tmp_int_pi2BPSK_ScramblingID0
	}
	if pi2BPSK_ScramblingID1Present {
		var tmp_int_pi2BPSK_ScramblingID1 int64
		if tmp_int_pi2BPSK_ScramblingID1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode pi2BPSK_ScramblingID1", err)
		}
		ie.pi2BPSK_ScramblingID1 = &tmp_int_pi2BPSK_ScramblingID1
	}
	return nil
}
