package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersXDD_Diff struct {
	dynamicSFI                      *Phy_ParametersXDD_Diff_dynamicSFI                      `optional`
	twoPUCCH_F0_2_ConsecSymbols     *Phy_ParametersXDD_Diff_twoPUCCH_F0_2_ConsecSymbols     `optional`
	twoDifferentTPC_Loop_PUSCH      *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH      `optional`
	twoDifferentTPC_Loop_PUCCH      *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUCCH      `optional`
	dl_SchedulingOffset_PDSCH_TypeA *Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeA `optional,ext-1`
	dl_SchedulingOffset_PDSCH_TypeB *Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeB `optional,ext-1`
	ul_SchedulingOffset             *Phy_ParametersXDD_Diff_ul_SchedulingOffset             `optional,ext-1`
}

func (ie *Phy_ParametersXDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dl_SchedulingOffset_PDSCH_TypeA != nil || ie.dl_SchedulingOffset_PDSCH_TypeB != nil || ie.ul_SchedulingOffset != nil
	preambleBits := []bool{hasExtensions, ie.dynamicSFI != nil, ie.twoPUCCH_F0_2_ConsecSymbols != nil, ie.twoDifferentTPC_Loop_PUSCH != nil, ie.twoDifferentTPC_Loop_PUCCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dynamicSFI != nil {
		if err = ie.dynamicSFI.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSFI", err)
		}
	}
	if ie.twoPUCCH_F0_2_ConsecSymbols != nil {
		if err = ie.twoPUCCH_F0_2_ConsecSymbols.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if ie.twoDifferentTPC_Loop_PUSCH != nil {
		if err = ie.twoDifferentTPC_Loop_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode twoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if ie.twoDifferentTPC_Loop_PUCCH != nil {
		if err = ie.twoDifferentTPC_Loop_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode twoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dl_SchedulingOffset_PDSCH_TypeA != nil || ie.dl_SchedulingOffset_PDSCH_TypeB != nil || ie.ul_SchedulingOffset != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dl_SchedulingOffset_PDSCH_TypeA != nil, ie.dl_SchedulingOffset_PDSCH_TypeB != nil, ie.ul_SchedulingOffset != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dl_SchedulingOffset_PDSCH_TypeA optional
			if ie.dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err = ie.dl_SchedulingOffset_PDSCH_TypeA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// encode dl_SchedulingOffset_PDSCH_TypeB optional
			if ie.dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err = ie.dl_SchedulingOffset_PDSCH_TypeB.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// encode ul_SchedulingOffset optional
			if ie.ul_SchedulingOffset != nil {
				if err = ie.ul_SchedulingOffset.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_SchedulingOffset", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *Phy_ParametersXDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSFIPresent bool
	if dynamicSFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_F0_2_ConsecSymbolsPresent bool
	if twoPUCCH_F0_2_ConsecSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoDifferentTPC_Loop_PUSCHPresent bool
	if twoDifferentTPC_Loop_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoDifferentTPC_Loop_PUCCHPresent bool
	if twoDifferentTPC_Loop_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dynamicSFIPresent {
		ie.dynamicSFI = new(Phy_ParametersXDD_Diff_dynamicSFI)
		if err = ie.dynamicSFI.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSFI", err)
		}
	}
	if twoPUCCH_F0_2_ConsecSymbolsPresent {
		ie.twoPUCCH_F0_2_ConsecSymbols = new(Phy_ParametersXDD_Diff_twoPUCCH_F0_2_ConsecSymbols)
		if err = ie.twoPUCCH_F0_2_ConsecSymbols.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if twoDifferentTPC_Loop_PUSCHPresent {
		ie.twoDifferentTPC_Loop_PUSCH = new(Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH)
		if err = ie.twoDifferentTPC_Loop_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode twoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if twoDifferentTPC_Loop_PUCCHPresent {
		ie.twoDifferentTPC_Loop_PUCCH = new(Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUCCH)
		if err = ie.twoDifferentTPC_Loop_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode twoDifferentTPC_Loop_PUCCH", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			dl_SchedulingOffset_PDSCH_TypeAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_SchedulingOffset_PDSCH_TypeBPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_SchedulingOffsetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dl_SchedulingOffset_PDSCH_TypeA optional
			if dl_SchedulingOffset_PDSCH_TypeAPresent {
				ie.dl_SchedulingOffset_PDSCH_TypeA = new(Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeA)
				if err = ie.dl_SchedulingOffset_PDSCH_TypeA.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// decode dl_SchedulingOffset_PDSCH_TypeB optional
			if dl_SchedulingOffset_PDSCH_TypeBPresent {
				ie.dl_SchedulingOffset_PDSCH_TypeB = new(Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeB)
				if err = ie.dl_SchedulingOffset_PDSCH_TypeB.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// decode ul_SchedulingOffset optional
			if ul_SchedulingOffsetPresent {
				ie.ul_SchedulingOffset = new(Phy_ParametersXDD_Diff_ul_SchedulingOffset)
				if err = ie.ul_SchedulingOffset.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_SchedulingOffset", err)
				}
			}
		}
	}
	return nil
}
