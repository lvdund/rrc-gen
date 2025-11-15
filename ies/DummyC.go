package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyC struct {
	maxNumberTxPortsPerResource    DummyC_maxNumberTxPortsPerResource `madatory`
	maxNumberResources             int64                              `lb:1,ub:64,madatory`
	totalNumberTxPorts             int64                              `lb:2,ub:256,madatory`
	supportedCodebookMode          DummyC_supportedCodebookMode       `madatory`
	supportedNumberPanels          DummyC_supportedNumberPanels       `madatory`
	maxNumberCSI_RS_PerResourceSet int64                              `lb:1,ub:8,madatory`
}

func (ie *DummyC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberTxPortsPerResource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberTxPortsPerResource", err)
	}
	if err = w.WriteInteger(ie.maxNumberResources, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberResources", err)
	}
	if err = w.WriteInteger(ie.totalNumberTxPorts, &uper.Constraint{Lb: 2, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger totalNumberTxPorts", err)
	}
	if err = ie.supportedCodebookMode.Encode(w); err != nil {
		return utils.WrapError("Encode supportedCodebookMode", err)
	}
	if err = ie.supportedNumberPanels.Encode(w); err != nil {
		return utils.WrapError("Encode supportedNumberPanels", err)
	}
	if err = w.WriteInteger(ie.maxNumberCSI_RS_PerResourceSet, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *DummyC) Decode(r *uper.UperReader) error {
	var err error
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
	if err = ie.supportedCodebookMode.Decode(r); err != nil {
		return utils.WrapError("Decode supportedCodebookMode", err)
	}
	if err = ie.supportedNumberPanels.Decode(r); err != nil {
		return utils.WrapError("Decode supportedNumberPanels", err)
	}
	var tmp_int_maxNumberCSI_RS_PerResourceSet int64
	if tmp_int_maxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	ie.maxNumberCSI_RS_PerResourceSet = tmp_int_maxNumberCSI_RS_PerResourceSet
	return nil
}
