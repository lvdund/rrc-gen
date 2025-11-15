package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_CSI_Resource struct {
	uplinkBandwidthPartId BWP_Id           `madatory`
	pucch_Resource        PUCCH_ResourceId `madatory`
}

func (ie *PUCCH_CSI_Resource) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.uplinkBandwidthPartId.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkBandwidthPartId", err)
	}
	if err = ie.pucch_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_Resource", err)
	}
	return nil
}

func (ie *PUCCH_CSI_Resource) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.uplinkBandwidthPartId.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkBandwidthPartId", err)
	}
	if err = ie.pucch_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_Resource", err)
	}
	return nil
}
