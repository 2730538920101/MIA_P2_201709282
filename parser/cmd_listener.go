// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CmdListener is a complete listener for a parse tree produced by CmdParser.
type CmdListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterAddCommand is called when entering the AddCommand production.
	EnterAddCommand(c *AddCommandContext)

	// EnterPause is called when entering the Pause production.
	EnterPause(c *PauseContext)

	// EnterLogout is called when entering the Logout production.
	EnterLogout(c *LogoutContext)

	// EnterMkdisk is called when entering the Mkdisk production.
	EnterMkdisk(c *MkdiskContext)

	// EnterRmdisk is called when entering the Rmdisk production.
	EnterRmdisk(c *RmdiskContext)

	// EnterFdisk is called when entering the Fdisk production.
	EnterFdisk(c *FdiskContext)

	// EnterMount is called when entering the Mount production.
	EnterMount(c *MountContext)

	// EnterMkfs is called when entering the Mkfs production.
	EnterMkfs(c *MkfsContext)

	// EnterLogin is called when entering the Login production.
	EnterLogin(c *LoginContext)

	// EnterMkgrp is called when entering the Mkgrp production.
	EnterMkgrp(c *MkgrpContext)

	// EnterRmgrp is called when entering the Rmgrp production.
	EnterRmgrp(c *RmgrpContext)

	// EnterMkusr is called when entering the Mkusr production.
	EnterMkusr(c *MkusrContext)

	// EnterRmusr is called when entering the Rmusr production.
	EnterRmusr(c *RmusrContext)

	// EnterMkfile is called when entering the Mkfile production.
	EnterMkfile(c *MkfileContext)

	// EnterMkdir is called when entering the Mkdir production.
	EnterMkdir(c *MkdirContext)

	// EnterExec is called when entering the Exec production.
	EnterExec(c *ExecContext)

	// EnterRep is called when entering the Rep production.
	EnterRep(c *RepContext)

	// EnterParam_list is called when entering the param_list production.
	EnterParam_list(c *Param_listContext)

	// EnterSize is called when entering the Size production.
	EnterSize(c *SizeContext)

	// EnterPath_R is called when entering the Path_R production.
	EnterPath_R(c *Path_RContext)

	// EnterPath_S is called when entering the Path_S production.
	EnterPath_S(c *Path_SContext)

	// EnterFitFF is called when entering the FitFF production.
	EnterFitFF(c *FitFFContext)

	// EnterFitWF is called when entering the FitWF production.
	EnterFitWF(c *FitWFContext)

	// EnterFitBF is called when entering the FitBF production.
	EnterFitBF(c *FitBFContext)

	// EnterUnit_B is called when entering the Unit_B production.
	EnterUnit_B(c *Unit_BContext)

	// EnterUnit_K is called when entering the Unit_K production.
	EnterUnit_K(c *Unit_KContext)

	// EnterUnit_M is called when entering the Unit_M production.
	EnterUnit_M(c *Unit_MContext)

	// EnterName_P is called when entering the Name_P production.
	EnterName_P(c *Name_PContext)

	// EnterName_S is called when entering the Name_S production.
	EnterName_S(c *Name_SContext)

	// EnterUsr_P is called when entering the Usr_P production.
	EnterUsr_P(c *Usr_PContext)

	// EnterUsr_S is called when entering the Usr_S production.
	EnterUsr_S(c *Usr_SContext)

	// EnterGrp_P is called when entering the Grp_P production.
	EnterGrp_P(c *Grp_PContext)

	// EnterGrp_S is called when entering the Grp_S production.
	EnterGrp_S(c *Grp_SContext)

	// EnterPass_P is called when entering the Pass_P production.
	EnterPass_P(c *Pass_PContext)

	// EnterPwd_P is called when entering the Pwd_P production.
	EnterPwd_P(c *Pwd_PContext)

	// EnterType_P is called when entering the Type_P production.
	EnterType_P(c *Type_PContext)

	// EnterType_L is called when entering the Type_L production.
	EnterType_L(c *Type_LContext)

	// EnterType_E is called when entering the Type_E production.
	EnterType_E(c *Type_EContext)

	// EnterType_Fast is called when entering the Type_Fast production.
	EnterType_Fast(c *Type_FastContext)

	// EnterType_Full is called when entering the Type_Full production.
	EnterType_Full(c *Type_FullContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// EnterCont_S is called when entering the Cont_S production.
	EnterCont_S(c *Cont_SContext)

	// EnterCont_R is called when entering the Cont_R production.
	EnterCont_R(c *Cont_RContext)

	// EnterRuta_S is called when entering the Ruta_S production.
	EnterRuta_S(c *Ruta_SContext)

	// EnterRuta_R is called when entering the Ruta_R production.
	EnterRuta_R(c *Ruta_RContext)

	// EnterPp is called when entering the Pp production.
	EnterPp(c *PpContext)

	// EnterRr is called when entering the Rr production.
	EnterRr(c *RrContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitAddCommand is called when exiting the AddCommand production.
	ExitAddCommand(c *AddCommandContext)

	// ExitPause is called when exiting the Pause production.
	ExitPause(c *PauseContext)

	// ExitLogout is called when exiting the Logout production.
	ExitLogout(c *LogoutContext)

	// ExitMkdisk is called when exiting the Mkdisk production.
	ExitMkdisk(c *MkdiskContext)

	// ExitRmdisk is called when exiting the Rmdisk production.
	ExitRmdisk(c *RmdiskContext)

	// ExitFdisk is called when exiting the Fdisk production.
	ExitFdisk(c *FdiskContext)

	// ExitMount is called when exiting the Mount production.
	ExitMount(c *MountContext)

	// ExitMkfs is called when exiting the Mkfs production.
	ExitMkfs(c *MkfsContext)

	// ExitLogin is called when exiting the Login production.
	ExitLogin(c *LoginContext)

	// ExitMkgrp is called when exiting the Mkgrp production.
	ExitMkgrp(c *MkgrpContext)

	// ExitRmgrp is called when exiting the Rmgrp production.
	ExitRmgrp(c *RmgrpContext)

	// ExitMkusr is called when exiting the Mkusr production.
	ExitMkusr(c *MkusrContext)

	// ExitRmusr is called when exiting the Rmusr production.
	ExitRmusr(c *RmusrContext)

	// ExitMkfile is called when exiting the Mkfile production.
	ExitMkfile(c *MkfileContext)

	// ExitMkdir is called when exiting the Mkdir production.
	ExitMkdir(c *MkdirContext)

	// ExitExec is called when exiting the Exec production.
	ExitExec(c *ExecContext)

	// ExitRep is called when exiting the Rep production.
	ExitRep(c *RepContext)

	// ExitParam_list is called when exiting the param_list production.
	ExitParam_list(c *Param_listContext)

	// ExitSize is called when exiting the Size production.
	ExitSize(c *SizeContext)

	// ExitPath_R is called when exiting the Path_R production.
	ExitPath_R(c *Path_RContext)

	// ExitPath_S is called when exiting the Path_S production.
	ExitPath_S(c *Path_SContext)

	// ExitFitFF is called when exiting the FitFF production.
	ExitFitFF(c *FitFFContext)

	// ExitFitWF is called when exiting the FitWF production.
	ExitFitWF(c *FitWFContext)

	// ExitFitBF is called when exiting the FitBF production.
	ExitFitBF(c *FitBFContext)

	// ExitUnit_B is called when exiting the Unit_B production.
	ExitUnit_B(c *Unit_BContext)

	// ExitUnit_K is called when exiting the Unit_K production.
	ExitUnit_K(c *Unit_KContext)

	// ExitUnit_M is called when exiting the Unit_M production.
	ExitUnit_M(c *Unit_MContext)

	// ExitName_P is called when exiting the Name_P production.
	ExitName_P(c *Name_PContext)

	// ExitName_S is called when exiting the Name_S production.
	ExitName_S(c *Name_SContext)

	// ExitUsr_P is called when exiting the Usr_P production.
	ExitUsr_P(c *Usr_PContext)

	// ExitUsr_S is called when exiting the Usr_S production.
	ExitUsr_S(c *Usr_SContext)

	// ExitGrp_P is called when exiting the Grp_P production.
	ExitGrp_P(c *Grp_PContext)

	// ExitGrp_S is called when exiting the Grp_S production.
	ExitGrp_S(c *Grp_SContext)

	// ExitPass_P is called when exiting the Pass_P production.
	ExitPass_P(c *Pass_PContext)

	// ExitPwd_P is called when exiting the Pwd_P production.
	ExitPwd_P(c *Pwd_PContext)

	// ExitType_P is called when exiting the Type_P production.
	ExitType_P(c *Type_PContext)

	// ExitType_L is called when exiting the Type_L production.
	ExitType_L(c *Type_LContext)

	// ExitType_E is called when exiting the Type_E production.
	ExitType_E(c *Type_EContext)

	// ExitType_Fast is called when exiting the Type_Fast production.
	ExitType_Fast(c *Type_FastContext)

	// ExitType_Full is called when exiting the Type_Full production.
	ExitType_Full(c *Type_FullContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)

	// ExitCont_S is called when exiting the Cont_S production.
	ExitCont_S(c *Cont_SContext)

	// ExitCont_R is called when exiting the Cont_R production.
	ExitCont_R(c *Cont_RContext)

	// ExitRuta_S is called when exiting the Ruta_S production.
	ExitRuta_S(c *Ruta_SContext)

	// ExitRuta_R is called when exiting the Ruta_R production.
	ExitRuta_R(c *Ruta_RContext)

	// ExitPp is called when exiting the Pp production.
	ExitPp(c *PpContext)

	// ExitRr is called when exiting the Rr production.
	ExitRr(c *RrContext)
}
