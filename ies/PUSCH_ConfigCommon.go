package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_ConfigCommon struct {
	groupHoppingEnabledTransformPrecoding *PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding `optional`
	pusch_TimeDomainAllocationList        *PUSCH_TimeDomainResourceAllocationList                   `optional`
	msg3_DeltaPreamble                    *int64                                                    `lb:-1,ub:6,optional`
	p0_NominalWithGrant                   *int64                                                    `lb:-202,ub:24,optional`
}

func (ie *PUSCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.groupHoppingEnabledTransformPrecoding != nil, ie.pusch_TimeDomainAllocationList != nil, ie.msg3_DeltaPreamble != nil, ie.p0_NominalWithGrant != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.groupHoppingEnabledTransformPrecoding != nil {
		if err = ie.groupHoppingEnabledTransformPrecoding.Encode(w); err != nil {
			return utils.WrapError("Encode groupHoppingEnabledTransformPrecoding", err)
		}
	}
	if ie.pusch_TimeDomainAllocationList != nil {
		if err = ie.pusch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_TimeDomainAllocationList", err)
		}
	}
	if ie.msg3_DeltaPreamble != nil {
		if err = w.WriteInteger(*ie.msg3_DeltaPreamble, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode msg3_DeltaPreamble", err)
		}
	}
	if ie.p0_NominalWithGrant != nil {
		if err = w.WriteInteger(*ie.p0_NominalWithGrant, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode p0_NominalWithGrant", err)
		}
	}
	return nil
}

func (ie *PUSCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var groupHoppingEnabledTransformPrecodingPresent bool
	if groupHoppingEnabledTransformPrecodingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_TimeDomainAllocationListPresent bool
	if pusch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var msg3_DeltaPreamblePresent bool
	if msg3_DeltaPreamblePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_NominalWithGrantPresent bool
	if p0_NominalWithGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if groupHoppingEnabledTransformPrecodingPresent {
		ie.groupHoppingEnabledTransformPrecoding = new(PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding)
		if err = ie.groupHoppingEnabledTransformPrecoding.Decode(r); err != nil {
			return utils.WrapError("Decode groupHoppingEnabledTransformPrecoding", err)
		}
	}
	if pusch_TimeDomainAllocationListPresent {
		ie.pusch_TimeDomainAllocationList = new(PUSCH_TimeDomainResourceAllocationList)
		if err = ie.pusch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_TimeDomainAllocationList", err)
		}
	}
	if msg3_DeltaPreamblePresent {
		var tmp_int_msg3_DeltaPreamble int64
		if tmp_int_msg3_DeltaPreamble, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode msg3_DeltaPreamble", err)
		}
		ie.msg3_DeltaPreamble = &tmp_int_msg3_DeltaPreamble
	}
	if p0_NominalWithGrantPresent {
		var tmp_int_p0_NominalWithGrant int64
		if tmp_int_p0_NominalWithGrant, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode p0_NominalWithGrant", err)
		}
		ie.p0_NominalWithGrant = &tmp_int_p0_NominalWithGrant
	}
	return nil
}
