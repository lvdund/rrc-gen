package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ResourceMapping struct {
	frequencyDomainAllocation    CSI_RS_ResourceMapping_frequencyDomainAllocation `lb:4,ub:4,madatory`
	nrofPorts                    CSI_RS_ResourceMapping_nrofPorts                 `madatory`
	firstOFDMSymbolInTimeDomain  int64                                            `lb:0,ub:13,madatory`
	firstOFDMSymbolInTimeDomain2 *int64                                           `lb:2,ub:12,optional`
	cdm_Type                     CSI_RS_ResourceMapping_cdm_Type                  `madatory`
	density                      CSI_RS_ResourceMapping_density                   `madatory`
	freqBand                     CSI_FrequencyOccupation                          `madatory`
}

func (ie *CSI_RS_ResourceMapping) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.firstOFDMSymbolInTimeDomain2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.frequencyDomainAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyDomainAllocation", err)
	}
	if err = ie.nrofPorts.Encode(w); err != nil {
		return utils.WrapError("Encode nrofPorts", err)
	}
	if err = w.WriteInteger(ie.firstOFDMSymbolInTimeDomain, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("WriteInteger firstOFDMSymbolInTimeDomain", err)
	}
	if ie.firstOFDMSymbolInTimeDomain2 != nil {
		if err = w.WriteInteger(*ie.firstOFDMSymbolInTimeDomain2, &uper.Constraint{Lb: 2, Ub: 12}, false); err != nil {
			return utils.WrapError("Encode firstOFDMSymbolInTimeDomain2", err)
		}
	}
	if err = ie.cdm_Type.Encode(w); err != nil {
		return utils.WrapError("Encode cdm_Type", err)
	}
	if err = ie.density.Encode(w); err != nil {
		return utils.WrapError("Encode density", err)
	}
	if err = ie.freqBand.Encode(w); err != nil {
		return utils.WrapError("Encode freqBand", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping) Decode(r *uper.UperReader) error {
	var err error
	var firstOFDMSymbolInTimeDomain2Present bool
	if firstOFDMSymbolInTimeDomain2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.frequencyDomainAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyDomainAllocation", err)
	}
	if err = ie.nrofPorts.Decode(r); err != nil {
		return utils.WrapError("Decode nrofPorts", err)
	}
	var tmp_int_firstOFDMSymbolInTimeDomain int64
	if tmp_int_firstOFDMSymbolInTimeDomain, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("ReadInteger firstOFDMSymbolInTimeDomain", err)
	}
	ie.firstOFDMSymbolInTimeDomain = tmp_int_firstOFDMSymbolInTimeDomain
	if firstOFDMSymbolInTimeDomain2Present {
		var tmp_int_firstOFDMSymbolInTimeDomain2 int64
		if tmp_int_firstOFDMSymbolInTimeDomain2, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode firstOFDMSymbolInTimeDomain2", err)
		}
		ie.firstOFDMSymbolInTimeDomain2 = &tmp_int_firstOFDMSymbolInTimeDomain2
	}
	if err = ie.cdm_Type.Decode(r); err != nil {
		return utils.WrapError("Decode cdm_Type", err)
	}
	if err = ie.density.Decode(r); err != nil {
		return utils.WrapError("Decode density", err)
	}
	if err = ie.freqBand.Decode(r); err != nil {
		return utils.WrapError("Decode freqBand", err)
	}
	return nil
}
