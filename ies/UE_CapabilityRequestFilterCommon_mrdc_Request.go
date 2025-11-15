package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon_mrdc_Request struct {
	omitEN_DC    *UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC    `optional`
	includeNR_DC *UE_CapabilityRequestFilterCommon_mrdc_Request_includeNR_DC `optional`
	includeNE_DC *UE_CapabilityRequestFilterCommon_mrdc_Request_includeNE_DC `optional`
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.omitEN_DC != nil, ie.includeNR_DC != nil, ie.includeNE_DC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.omitEN_DC != nil {
		if err = ie.omitEN_DC.Encode(w); err != nil {
			return utils.WrapError("Encode omitEN_DC", err)
		}
	}
	if ie.includeNR_DC != nil {
		if err = ie.includeNR_DC.Encode(w); err != nil {
			return utils.WrapError("Encode includeNR_DC", err)
		}
	}
	if ie.includeNE_DC != nil {
		if err = ie.includeNE_DC.Encode(w); err != nil {
			return utils.WrapError("Encode includeNE_DC", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request) Decode(r *uper.UperReader) error {
	var err error
	var omitEN_DCPresent bool
	if omitEN_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var includeNR_DCPresent bool
	if includeNR_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var includeNE_DCPresent bool
	if includeNE_DCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if omitEN_DCPresent {
		ie.omitEN_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC)
		if err = ie.omitEN_DC.Decode(r); err != nil {
			return utils.WrapError("Decode omitEN_DC", err)
		}
	}
	if includeNR_DCPresent {
		ie.includeNR_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_includeNR_DC)
		if err = ie.includeNR_DC.Decode(r); err != nil {
			return utils.WrapError("Decode includeNR_DC", err)
		}
	}
	if includeNE_DCPresent {
		ie.includeNE_DC = new(UE_CapabilityRequestFilterCommon_mrdc_Request_includeNE_DC)
		if err = ie.includeNE_DC.Decode(r); err != nil {
			return utils.WrapError("Decode includeNE_DC", err)
		}
	}
	return nil
}
