package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1530 struct {
	fdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530           `optional`
	tdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530           `optional`
	dummy                            *UE_NR_Capability_v1530_dummy                `optional`
	interRAT_Parameters              *InterRAT_Parameters                         `optional`
	inactiveState                    *UE_NR_Capability_v1530_inactiveState        `optional`
	delayBudgetReporting             *UE_NR_Capability_v1530_delayBudgetReporting `optional`
	nonCriticalExtension             *UE_NR_Capability_v1540                      `optional`
}

func (ie *UE_NR_Capability_v1530) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fdd_Add_UE_NR_Capabilities_v1530 != nil, ie.tdd_Add_UE_NR_Capabilities_v1530 != nil, ie.dummy != nil, ie.interRAT_Parameters != nil, ie.inactiveState != nil, ie.delayBudgetReporting != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fdd_Add_UE_NR_Capabilities_v1530 != nil {
		if err = ie.fdd_Add_UE_NR_Capabilities_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if ie.tdd_Add_UE_NR_Capabilities_v1530 != nil {
		if err = ie.tdd_Add_UE_NR_Capabilities_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.interRAT_Parameters != nil {
		if err = ie.interRAT_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode interRAT_Parameters", err)
		}
	}
	if ie.inactiveState != nil {
		if err = ie.inactiveState.Encode(w); err != nil {
			return utils.WrapError("Encode inactiveState", err)
		}
	}
	if ie.delayBudgetReporting != nil {
		if err = ie.delayBudgetReporting.Encode(w); err != nil {
			return utils.WrapError("Encode delayBudgetReporting", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1530) Decode(r *uper.UperReader) error {
	var err error
	var fdd_Add_UE_NR_Capabilities_v1530Present bool
	if fdd_Add_UE_NR_Capabilities_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_NR_Capabilities_v1530Present bool
	if tdd_Add_UE_NR_Capabilities_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var interRAT_ParametersPresent bool
	if interRAT_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var inactiveStatePresent bool
	if inactiveStatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var delayBudgetReportingPresent bool
	if delayBudgetReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if fdd_Add_UE_NR_Capabilities_v1530Present {
		ie.fdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
		if err = ie.fdd_Add_UE_NR_Capabilities_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if tdd_Add_UE_NR_Capabilities_v1530Present {
		ie.tdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
		if err = ie.tdd_Add_UE_NR_Capabilities_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if dummyPresent {
		ie.dummy = new(UE_NR_Capability_v1530_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if interRAT_ParametersPresent {
		ie.interRAT_Parameters = new(InterRAT_Parameters)
		if err = ie.interRAT_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode interRAT_Parameters", err)
		}
	}
	if inactiveStatePresent {
		ie.inactiveState = new(UE_NR_Capability_v1530_inactiveState)
		if err = ie.inactiveState.Decode(r); err != nil {
			return utils.WrapError("Decode inactiveState", err)
		}
	}
	if delayBudgetReportingPresent {
		ie.delayBudgetReporting = new(UE_NR_Capability_v1530_delayBudgetReporting)
		if err = ie.delayBudgetReporting.Decode(r); err != nil {
			return utils.WrapError("Decode delayBudgetReporting", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1540)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
