package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SupportedCSI_RS_Resource struct {
	maxNumberTxPortsPerResource SupportedCSI_RS_Resource_maxNumberTxPortsPerResource `madatory`
	maxNumberResourcesPerBand   int64                                                `lb:1,ub:64,madatory`
	totalNumberTxPortsPerBand   int64                                                `lb:2,ub:256,madatory`
}

func (ie *SupportedCSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.maxNumberResourcesPerBand, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberResourcesPerBand", err)
	}
	if err = w.WriteInteger(ie.totalNumberTxPortsPerBand, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger totalNumberTxPortsPerBand", err)
	}
	return nil
}

func (ie *SupportedCSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberTxPortsPerResource.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberTxPortsPerResource", err)
	}
	var tmp_int_maxNumberResourcesPerBand int64
	if tmp_int_maxNumberResourcesPerBand, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberResourcesPerBand", err)
	}
	ie.maxNumberResourcesPerBand = tmp_int_maxNumberResourcesPerBand
	var tmp_int_totalNumberTxPortsPerBand int64
	if tmp_int_totalNumberTxPortsPerBand, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger totalNumberTxPortsPerBand", err)
	}
	ie.totalNumberTxPortsPerBand = tmp_int_totalNumberTxPortsPerBand
	return nil
}
