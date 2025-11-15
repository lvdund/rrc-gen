package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyD struct {
	maxNumberTxPortsPerResource    DummyD_maxNumberTxPortsPerResource `madatory`
	maxNumberResources             int64                              `lb:1,ub:64,madatory`
	totalNumberTxPorts             int64                              `lb:2,ub:256,madatory`
	parameterLx                    int64                              `lb:2,ub:4,madatory`
	amplitudeScalingType           DummyD_amplitudeScalingType        `madatory`
	amplitudeSubsetRestriction     *DummyD_amplitudeSubsetRestriction `optional`
	maxNumberCSI_RS_PerResourceSet int64                              `lb:1,ub:8,madatory`
}

func (ie *DummyD) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.amplitudeSubsetRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.maxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.maxNumberResources, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberResources", err)
	}
	if err = w.WriteInteger(ie.totalNumberTxPorts, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger totalNumberTxPorts", err)
	}
	if err = w.WriteInteger(ie.parameterLx, &uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger parameterLx", err)
	}
	if err = ie.amplitudeScalingType.Encode(w); err != nil {
		return utils.WrapError("Encode amplitudeScalingType", err)
	}
	if ie.amplitudeSubsetRestriction != nil {
		if err = ie.amplitudeSubsetRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode amplitudeSubsetRestriction", err)
		}
	}
	if err = w.WriteInteger(ie.maxNumberCSI_RS_PerResourceSet, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *DummyD) Decode(r *uper.UperReader) error {
	var err error
	var amplitudeSubsetRestrictionPresent bool
	if amplitudeSubsetRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.maxNumberTxPortsPerResource.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberTxPortsPerResource", err)
	}
	var tmp_int_maxNumberResources int64
	if tmp_int_maxNumberResources, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberResources", err)
	}
	ie.maxNumberResources = tmp_int_maxNumberResources
	var tmp_int_totalNumberTxPorts int64
	if tmp_int_totalNumberTxPorts, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger totalNumberTxPorts", err)
	}
	ie.totalNumberTxPorts = tmp_int_totalNumberTxPorts
	var tmp_int_parameterLx int64
	if tmp_int_parameterLx, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger parameterLx", err)
	}
	ie.parameterLx = tmp_int_parameterLx
	if err = ie.amplitudeScalingType.Decode(r); err != nil {
		return utils.WrapError("Decode amplitudeScalingType", err)
	}
	if amplitudeSubsetRestrictionPresent {
		ie.amplitudeSubsetRestriction = new(DummyD_amplitudeSubsetRestriction)
		if err = ie.amplitudeSubsetRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode amplitudeSubsetRestriction", err)
		}
	}
	var tmp_int_maxNumberCSI_RS_PerResourceSet int64
	if tmp_int_maxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	ie.maxNumberCSI_RS_PerResourceSet = tmp_int_maxNumberCSI_RS_PerResourceSet
	return nil
}
