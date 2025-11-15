package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1550 struct {
	reducedCP_Latency    *UE_NR_Capability_v1550_reducedCP_Latency `optional`
	nonCriticalExtension *UE_NR_Capability_v1560                   `optional`
}

func (ie *UE_NR_Capability_v1550) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.reducedCP_Latency != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.reducedCP_Latency != nil {
		if err = ie.reducedCP_Latency.Encode(w); err != nil {
			return utils.WrapError("Encode reducedCP_Latency", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1550) Decode(r *uper.UperReader) error {
	var err error
	var reducedCP_LatencyPresent bool
	if reducedCP_LatencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if reducedCP_LatencyPresent {
		ie.reducedCP_Latency = new(UE_NR_Capability_v1550_reducedCP_Latency)
		if err = ie.reducedCP_Latency.Decode(r); err != nil {
			return utils.WrapError("Decode reducedCP_Latency", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1560)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
