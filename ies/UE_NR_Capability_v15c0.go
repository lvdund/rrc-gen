package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v15c0 struct {
	nrdc_Parameters_v15c0     *NRDC_Parameters_v15c0                            `optional`
	partialFR2_FallbackRX_Req *UE_NR_Capability_v15c0_partialFR2_FallbackRX_Req `optional`
	nonCriticalExtension      *UE_NR_Capability_v15g0                           `optional`
}

func (ie *UE_NR_Capability_v15c0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrdc_Parameters_v15c0 != nil, ie.partialFR2_FallbackRX_Req != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrdc_Parameters_v15c0 != nil {
		if err = ie.nrdc_Parameters_v15c0.Encode(w); err != nil {
			return utils.WrapError("Encode nrdc_Parameters_v15c0", err)
		}
	}
	if ie.partialFR2_FallbackRX_Req != nil {
		if err = ie.partialFR2_FallbackRX_Req.Encode(w); err != nil {
			return utils.WrapError("Encode partialFR2_FallbackRX_Req", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v15c0) Decode(r *uper.UperReader) error {
	var err error
	var nrdc_Parameters_v15c0Present bool
	if nrdc_Parameters_v15c0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var partialFR2_FallbackRX_ReqPresent bool
	if partialFR2_FallbackRX_ReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrdc_Parameters_v15c0Present {
		ie.nrdc_Parameters_v15c0 = new(NRDC_Parameters_v15c0)
		if err = ie.nrdc_Parameters_v15c0.Decode(r); err != nil {
			return utils.WrapError("Decode nrdc_Parameters_v15c0", err)
		}
	}
	if partialFR2_FallbackRX_ReqPresent {
		ie.partialFR2_FallbackRX_Req = new(UE_NR_Capability_v15c0_partialFR2_FallbackRX_Req)
		if err = ie.partialFR2_FallbackRX_Req.Decode(r); err != nil {
			return utils.WrapError("Decode partialFR2_FallbackRX_Req", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v15g0)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
