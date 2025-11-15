package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands3
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands4
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands5
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands6
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands7
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands8
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands9
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands10
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands11
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands12
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands13
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands14
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands15
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands16
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands17
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands18
	CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands19_v1530
)

type CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand struct {
	Choice           uint64
	subbands3        uper.BitString `lb:3,ub:3,madatory`
	subbands4        uper.BitString `lb:4,ub:4,madatory`
	subbands5        uper.BitString `lb:5,ub:5,madatory`
	subbands6        uper.BitString `lb:6,ub:6,madatory`
	subbands7        uper.BitString `lb:7,ub:7,madatory`
	subbands8        uper.BitString `lb:8,ub:8,madatory`
	subbands9        uper.BitString `lb:9,ub:9,madatory`
	subbands10       uper.BitString `lb:10,ub:10,madatory`
	subbands11       uper.BitString `lb:11,ub:11,madatory`
	subbands12       uper.BitString `lb:12,ub:12,madatory`
	subbands13       uper.BitString `lb:13,ub:13,madatory`
	subbands14       uper.BitString `lb:14,ub:14,madatory`
	subbands15       uper.BitString `lb:15,ub:15,madatory`
	subbands16       uper.BitString `lb:16,ub:16,madatory`
	subbands17       uper.BitString `lb:17,ub:17,madatory`
	subbands18       uper.BitString `lb:18,ub:18,madatory`
	subbands19_v1530 uper.BitString `lb:19,ub:19,madatory`
}

func (ie *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands3:
		if err = w.WriteBitString(ie.subbands3.Bytes, uint(ie.subbands3.NumBits), &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode subbands3", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands4:
		if err = w.WriteBitString(ie.subbands4.Bytes, uint(ie.subbands4.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode subbands4", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands5:
		if err = w.WriteBitString(ie.subbands5.Bytes, uint(ie.subbands5.NumBits), &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode subbands5", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands6:
		if err = w.WriteBitString(ie.subbands6.Bytes, uint(ie.subbands6.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			err = utils.WrapError("Encode subbands6", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands7:
		if err = w.WriteBitString(ie.subbands7.Bytes, uint(ie.subbands7.NumBits), &uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode subbands7", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands8:
		if err = w.WriteBitString(ie.subbands8.Bytes, uint(ie.subbands8.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode subbands8", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands9:
		if err = w.WriteBitString(ie.subbands9.Bytes, uint(ie.subbands9.NumBits), &uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode subbands9", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands10:
		if err = w.WriteBitString(ie.subbands10.Bytes, uint(ie.subbands10.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode subbands10", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands11:
		if err = w.WriteBitString(ie.subbands11.Bytes, uint(ie.subbands11.NumBits), &uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
			err = utils.WrapError("Encode subbands11", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands12:
		if err = w.WriteBitString(ie.subbands12.Bytes, uint(ie.subbands12.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			err = utils.WrapError("Encode subbands12", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands13:
		if err = w.WriteBitString(ie.subbands13.Bytes, uint(ie.subbands13.NumBits), &uper.Constraint{Lb: 13, Ub: 13}, false); err != nil {
			err = utils.WrapError("Encode subbands13", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands14:
		if err = w.WriteBitString(ie.subbands14.Bytes, uint(ie.subbands14.NumBits), &uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			err = utils.WrapError("Encode subbands14", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands15:
		if err = w.WriteBitString(ie.subbands15.Bytes, uint(ie.subbands15.NumBits), &uper.Constraint{Lb: 15, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode subbands15", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands16:
		if err = w.WriteBitString(ie.subbands16.Bytes, uint(ie.subbands16.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode subbands16", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands17:
		if err = w.WriteBitString(ie.subbands17.Bytes, uint(ie.subbands17.NumBits), &uper.Constraint{Lb: 17, Ub: 17}, false); err != nil {
			err = utils.WrapError("Encode subbands17", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands18:
		if err = w.WriteBitString(ie.subbands18.Bytes, uint(ie.subbands18.NumBits), &uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
			err = utils.WrapError("Encode subbands18", err)
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands19_v1530:
		if err = w.WriteBitString(ie.subbands19_v1530.Bytes, uint(ie.subbands19_v1530.NumBits), &uper.Constraint{Lb: 19, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode subbands19_v1530", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands3:
		var tmp_bs_subbands3 []byte
		var n_subbands3 uint
		if tmp_bs_subbands3, n_subbands3, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode subbands3", err)
		}
		ie.subbands3 = uper.BitString{
			Bytes:   tmp_bs_subbands3,
			NumBits: uint64(n_subbands3),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands4:
		var tmp_bs_subbands4 []byte
		var n_subbands4 uint
		if tmp_bs_subbands4, n_subbands4, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode subbands4", err)
		}
		ie.subbands4 = uper.BitString{
			Bytes:   tmp_bs_subbands4,
			NumBits: uint64(n_subbands4),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands5:
		var tmp_bs_subbands5 []byte
		var n_subbands5 uint
		if tmp_bs_subbands5, n_subbands5, err = r.ReadBitString(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode subbands5", err)
		}
		ie.subbands5 = uper.BitString{
			Bytes:   tmp_bs_subbands5,
			NumBits: uint64(n_subbands5),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands6:
		var tmp_bs_subbands6 []byte
		var n_subbands6 uint
		if tmp_bs_subbands6, n_subbands6, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode subbands6", err)
		}
		ie.subbands6 = uper.BitString{
			Bytes:   tmp_bs_subbands6,
			NumBits: uint64(n_subbands6),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands7:
		var tmp_bs_subbands7 []byte
		var n_subbands7 uint
		if tmp_bs_subbands7, n_subbands7, err = r.ReadBitString(&uper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode subbands7", err)
		}
		ie.subbands7 = uper.BitString{
			Bytes:   tmp_bs_subbands7,
			NumBits: uint64(n_subbands7),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands8:
		var tmp_bs_subbands8 []byte
		var n_subbands8 uint
		if tmp_bs_subbands8, n_subbands8, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode subbands8", err)
		}
		ie.subbands8 = uper.BitString{
			Bytes:   tmp_bs_subbands8,
			NumBits: uint64(n_subbands8),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands9:
		var tmp_bs_subbands9 []byte
		var n_subbands9 uint
		if tmp_bs_subbands9, n_subbands9, err = r.ReadBitString(&uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode subbands9", err)
		}
		ie.subbands9 = uper.BitString{
			Bytes:   tmp_bs_subbands9,
			NumBits: uint64(n_subbands9),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands10:
		var tmp_bs_subbands10 []byte
		var n_subbands10 uint
		if tmp_bs_subbands10, n_subbands10, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode subbands10", err)
		}
		ie.subbands10 = uper.BitString{
			Bytes:   tmp_bs_subbands10,
			NumBits: uint64(n_subbands10),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands11:
		var tmp_bs_subbands11 []byte
		var n_subbands11 uint
		if tmp_bs_subbands11, n_subbands11, err = r.ReadBitString(&uper.Constraint{Lb: 11, Ub: 11}, false); err != nil {
			return utils.WrapError("Decode subbands11", err)
		}
		ie.subbands11 = uper.BitString{
			Bytes:   tmp_bs_subbands11,
			NumBits: uint64(n_subbands11),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands12:
		var tmp_bs_subbands12 []byte
		var n_subbands12 uint
		if tmp_bs_subbands12, n_subbands12, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
			return utils.WrapError("Decode subbands12", err)
		}
		ie.subbands12 = uper.BitString{
			Bytes:   tmp_bs_subbands12,
			NumBits: uint64(n_subbands12),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands13:
		var tmp_bs_subbands13 []byte
		var n_subbands13 uint
		if tmp_bs_subbands13, n_subbands13, err = r.ReadBitString(&uper.Constraint{Lb: 13, Ub: 13}, false); err != nil {
			return utils.WrapError("Decode subbands13", err)
		}
		ie.subbands13 = uper.BitString{
			Bytes:   tmp_bs_subbands13,
			NumBits: uint64(n_subbands13),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands14:
		var tmp_bs_subbands14 []byte
		var n_subbands14 uint
		if tmp_bs_subbands14, n_subbands14, err = r.ReadBitString(&uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode subbands14", err)
		}
		ie.subbands14 = uper.BitString{
			Bytes:   tmp_bs_subbands14,
			NumBits: uint64(n_subbands14),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands15:
		var tmp_bs_subbands15 []byte
		var n_subbands15 uint
		if tmp_bs_subbands15, n_subbands15, err = r.ReadBitString(&uper.Constraint{Lb: 15, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode subbands15", err)
		}
		ie.subbands15 = uper.BitString{
			Bytes:   tmp_bs_subbands15,
			NumBits: uint64(n_subbands15),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands16:
		var tmp_bs_subbands16 []byte
		var n_subbands16 uint
		if tmp_bs_subbands16, n_subbands16, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode subbands16", err)
		}
		ie.subbands16 = uper.BitString{
			Bytes:   tmp_bs_subbands16,
			NumBits: uint64(n_subbands16),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands17:
		var tmp_bs_subbands17 []byte
		var n_subbands17 uint
		if tmp_bs_subbands17, n_subbands17, err = r.ReadBitString(&uper.Constraint{Lb: 17, Ub: 17}, false); err != nil {
			return utils.WrapError("Decode subbands17", err)
		}
		ie.subbands17 = uper.BitString{
			Bytes:   tmp_bs_subbands17,
			NumBits: uint64(n_subbands17),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands18:
		var tmp_bs_subbands18 []byte
		var n_subbands18 uint
		if tmp_bs_subbands18, n_subbands18, err = r.ReadBitString(&uper.Constraint{Lb: 18, Ub: 18}, false); err != nil {
			return utils.WrapError("Decode subbands18", err)
		}
		ie.subbands18 = uper.BitString{
			Bytes:   tmp_bs_subbands18,
			NumBits: uint64(n_subbands18),
		}
	case CSI_ReportConfig_reportFreqConfiguration_csi_ReportingBand_Choice_subbands19_v1530:
		var tmp_bs_subbands19_v1530 []byte
		var n_subbands19_v1530 uint
		if tmp_bs_subbands19_v1530, n_subbands19_v1530, err = r.ReadBitString(&uper.Constraint{Lb: 19, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode subbands19_v1530", err)
		}
		ie.subbands19_v1530 = uper.BitString{
			Bytes:   tmp_bs_subbands19_v1530,
			NumBits: uint64(n_subbands19_v1530),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
