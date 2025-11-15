package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig_transformPrecodingEnabled struct {
	nPUSCH_Identity                   *int64                                                            `lb:0,ub:1007,optional`
	sequenceGroupHopping              *DMRS_UplinkConfig_transformPrecodingEnabled_sequenceGroupHopping `optional`
	sequenceHopping                   *DMRS_UplinkConfig_transformPrecodingEnabled_sequenceHopping      `optional`
	dmrs_UplinkTransformPrecoding_r16 *DMRS_UplinkTransformPrecoding_r16                                `optional,setuprelease`
}

func (ie *DMRS_UplinkConfig_transformPrecodingEnabled) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nPUSCH_Identity != nil, ie.sequenceGroupHopping != nil, ie.sequenceHopping != nil, ie.dmrs_UplinkTransformPrecoding_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nPUSCH_Identity != nil {
		if err = w.WriteInteger(*ie.nPUSCH_Identity, &uper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
			return utils.WrapError("Encode nPUSCH_Identity", err)
		}
	}
	if ie.sequenceGroupHopping != nil {
		if err = ie.sequenceGroupHopping.Encode(w); err != nil {
			return utils.WrapError("Encode sequenceGroupHopping", err)
		}
	}
	if ie.sequenceHopping != nil {
		if err = ie.sequenceHopping.Encode(w); err != nil {
			return utils.WrapError("Encode sequenceHopping", err)
		}
	}
	if ie.dmrs_UplinkTransformPrecoding_r16 != nil {
		tmp_dmrs_UplinkTransformPrecoding_r16 := utils.SetupRelease[*DMRS_UplinkTransformPrecoding_r16]{
			Setup: ie.dmrs_UplinkTransformPrecoding_r16,
		}
		if err = tmp_dmrs_UplinkTransformPrecoding_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_UplinkTransformPrecoding_r16", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig_transformPrecodingEnabled) Decode(r *uper.UperReader) error {
	var err error
	var nPUSCH_IdentityPresent bool
	if nPUSCH_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sequenceGroupHoppingPresent bool
	if sequenceGroupHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sequenceHoppingPresent bool
	if sequenceHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_UplinkTransformPrecoding_r16Present bool
	if dmrs_UplinkTransformPrecoding_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nPUSCH_IdentityPresent {
		var tmp_int_nPUSCH_Identity int64
		if tmp_int_nPUSCH_Identity, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
			return utils.WrapError("Decode nPUSCH_Identity", err)
		}
		ie.nPUSCH_Identity = &tmp_int_nPUSCH_Identity
	}
	if sequenceGroupHoppingPresent {
		ie.sequenceGroupHopping = new(DMRS_UplinkConfig_transformPrecodingEnabled_sequenceGroupHopping)
		if err = ie.sequenceGroupHopping.Decode(r); err != nil {
			return utils.WrapError("Decode sequenceGroupHopping", err)
		}
	}
	if sequenceHoppingPresent {
		ie.sequenceHopping = new(DMRS_UplinkConfig_transformPrecodingEnabled_sequenceHopping)
		if err = ie.sequenceHopping.Decode(r); err != nil {
			return utils.WrapError("Decode sequenceHopping", err)
		}
	}
	if dmrs_UplinkTransformPrecoding_r16Present {
		tmp_dmrs_UplinkTransformPrecoding_r16 := utils.SetupRelease[*DMRS_UplinkTransformPrecoding_r16]{}
		if err = tmp_dmrs_UplinkTransformPrecoding_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_UplinkTransformPrecoding_r16", err)
		}
		ie.dmrs_UplinkTransformPrecoding_r16 = tmp_dmrs_UplinkTransformPrecoding_r16.Setup
	}
	return nil
}
