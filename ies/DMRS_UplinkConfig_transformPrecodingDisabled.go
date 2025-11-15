package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig_transformPrecodingDisabled struct {
	scramblingID0   *int64                                                        `lb:0,ub:65535,optional`
	scramblingID1   *int64                                                        `lb:0,ub:65535,optional`
	dmrs_Uplink_r16 *DMRS_UplinkConfig_transformPrecodingDisabled_dmrs_Uplink_r16 `optional`
}

func (ie *DMRS_UplinkConfig_transformPrecodingDisabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scramblingID0 != nil, ie.scramblingID1 != nil, ie.dmrs_Uplink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scramblingID0 != nil {
		if err = w.WriteInteger(*ie.scramblingID0, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode scramblingID0", err)
		}
	}
	if ie.scramblingID1 != nil {
		if err = w.WriteInteger(*ie.scramblingID1, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode scramblingID1", err)
		}
	}
	if ie.dmrs_Uplink_r16 != nil {
		if err = ie.dmrs_Uplink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_Uplink_r16", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig_transformPrecodingDisabled) Decode(r *uper.UperReader) error {
	var err error
	var scramblingID0Present bool
	if scramblingID0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scramblingID1Present bool
	if scramblingID1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_Uplink_r16Present bool
	if dmrs_Uplink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scramblingID0Present {
		var tmp_int_scramblingID0 int64
		if tmp_int_scramblingID0, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode scramblingID0", err)
		}
		ie.scramblingID0 = &tmp_int_scramblingID0
	}
	if scramblingID1Present {
		var tmp_int_scramblingID1 int64
		if tmp_int_scramblingID1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode scramblingID1", err)
		}
		ie.scramblingID1 = &tmp_int_scramblingID1
	}
	if dmrs_Uplink_r16Present {
		ie.dmrs_Uplink_r16 = new(DMRS_UplinkConfig_transformPrecodingDisabled_dmrs_Uplink_r16)
		if err = ie.dmrs_Uplink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_Uplink_r16", err)
		}
	}
	return nil
}
