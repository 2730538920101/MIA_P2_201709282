// Code generated from Cmd.g4 by ANTLR 4.10. DO NOT EDIT.

package parser // Cmd

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCmdListener is a complete listener for a parse tree produced by CmdParser.
type BaseCmdListener struct{}

var _ CmdListener = &BaseCmdListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCmdListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCmdListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCmdListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCmdListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCmdListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCmdListener) ExitStart(ctx *StartContext) {}

// EnterAddCommand is called when production AddCommand is entered.
func (s *BaseCmdListener) EnterAddCommand(ctx *AddCommandContext) {}

// ExitAddCommand is called when production AddCommand is exited.
func (s *BaseCmdListener) ExitAddCommand(ctx *AddCommandContext) {}

// EnterPause is called when production Pause is entered.
func (s *BaseCmdListener) EnterPause(ctx *PauseContext) {}

// ExitPause is called when production Pause is exited.
func (s *BaseCmdListener) ExitPause(ctx *PauseContext) {}

// EnterLogout is called when production Logout is entered.
func (s *BaseCmdListener) EnterLogout(ctx *LogoutContext) {}

// ExitLogout is called when production Logout is exited.
func (s *BaseCmdListener) ExitLogout(ctx *LogoutContext) {}

// EnterMkdisk is called when production Mkdisk is entered.
func (s *BaseCmdListener) EnterMkdisk(ctx *MkdiskContext) {}

// ExitMkdisk is called when production Mkdisk is exited.
func (s *BaseCmdListener) ExitMkdisk(ctx *MkdiskContext) {}

// EnterRmdisk is called when production Rmdisk is entered.
func (s *BaseCmdListener) EnterRmdisk(ctx *RmdiskContext) {}

// ExitRmdisk is called when production Rmdisk is exited.
func (s *BaseCmdListener) ExitRmdisk(ctx *RmdiskContext) {}

// EnterFdisk is called when production Fdisk is entered.
func (s *BaseCmdListener) EnterFdisk(ctx *FdiskContext) {}

// ExitFdisk is called when production Fdisk is exited.
func (s *BaseCmdListener) ExitFdisk(ctx *FdiskContext) {}

// EnterMount is called when production Mount is entered.
func (s *BaseCmdListener) EnterMount(ctx *MountContext) {}

// ExitMount is called when production Mount is exited.
func (s *BaseCmdListener) ExitMount(ctx *MountContext) {}

// EnterMkfs is called when production Mkfs is entered.
func (s *BaseCmdListener) EnterMkfs(ctx *MkfsContext) {}

// ExitMkfs is called when production Mkfs is exited.
func (s *BaseCmdListener) ExitMkfs(ctx *MkfsContext) {}

// EnterLogin is called when production Login is entered.
func (s *BaseCmdListener) EnterLogin(ctx *LoginContext) {}

// ExitLogin is called when production Login is exited.
func (s *BaseCmdListener) ExitLogin(ctx *LoginContext) {}

// EnterMkgrp is called when production Mkgrp is entered.
func (s *BaseCmdListener) EnterMkgrp(ctx *MkgrpContext) {}

// ExitMkgrp is called when production Mkgrp is exited.
func (s *BaseCmdListener) ExitMkgrp(ctx *MkgrpContext) {}

// EnterRmgrp is called when production Rmgrp is entered.
func (s *BaseCmdListener) EnterRmgrp(ctx *RmgrpContext) {}

// ExitRmgrp is called when production Rmgrp is exited.
func (s *BaseCmdListener) ExitRmgrp(ctx *RmgrpContext) {}

// EnterMkusr is called when production Mkusr is entered.
func (s *BaseCmdListener) EnterMkusr(ctx *MkusrContext) {}

// ExitMkusr is called when production Mkusr is exited.
func (s *BaseCmdListener) ExitMkusr(ctx *MkusrContext) {}

// EnterRmusr is called when production Rmusr is entered.
func (s *BaseCmdListener) EnterRmusr(ctx *RmusrContext) {}

// ExitRmusr is called when production Rmusr is exited.
func (s *BaseCmdListener) ExitRmusr(ctx *RmusrContext) {}

// EnterMkfile is called when production Mkfile is entered.
func (s *BaseCmdListener) EnterMkfile(ctx *MkfileContext) {}

// ExitMkfile is called when production Mkfile is exited.
func (s *BaseCmdListener) ExitMkfile(ctx *MkfileContext) {}

// EnterMkdir is called when production Mkdir is entered.
func (s *BaseCmdListener) EnterMkdir(ctx *MkdirContext) {}

// ExitMkdir is called when production Mkdir is exited.
func (s *BaseCmdListener) ExitMkdir(ctx *MkdirContext) {}

// EnterExec is called when production Exec is entered.
func (s *BaseCmdListener) EnterExec(ctx *ExecContext) {}

// ExitExec is called when production Exec is exited.
func (s *BaseCmdListener) ExitExec(ctx *ExecContext) {}

// EnterRep is called when production Rep is entered.
func (s *BaseCmdListener) EnterRep(ctx *RepContext) {}

// ExitRep is called when production Rep is exited.
func (s *BaseCmdListener) ExitRep(ctx *RepContext) {}

// EnterParam_list is called when production param_list is entered.
func (s *BaseCmdListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BaseCmdListener) ExitParam_list(ctx *Param_listContext) {}

// EnterSize is called when production Size is entered.
func (s *BaseCmdListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production Size is exited.
func (s *BaseCmdListener) ExitSize(ctx *SizeContext) {}

// EnterPath_R is called when production Path_R is entered.
func (s *BaseCmdListener) EnterPath_R(ctx *Path_RContext) {}

// ExitPath_R is called when production Path_R is exited.
func (s *BaseCmdListener) ExitPath_R(ctx *Path_RContext) {}

// EnterPath_S is called when production Path_S is entered.
func (s *BaseCmdListener) EnterPath_S(ctx *Path_SContext) {}

// ExitPath_S is called when production Path_S is exited.
func (s *BaseCmdListener) ExitPath_S(ctx *Path_SContext) {}

// EnterFitFF is called when production FitFF is entered.
func (s *BaseCmdListener) EnterFitFF(ctx *FitFFContext) {}

// ExitFitFF is called when production FitFF is exited.
func (s *BaseCmdListener) ExitFitFF(ctx *FitFFContext) {}

// EnterFitWF is called when production FitWF is entered.
func (s *BaseCmdListener) EnterFitWF(ctx *FitWFContext) {}

// ExitFitWF is called when production FitWF is exited.
func (s *BaseCmdListener) ExitFitWF(ctx *FitWFContext) {}

// EnterFitBF is called when production FitBF is entered.
func (s *BaseCmdListener) EnterFitBF(ctx *FitBFContext) {}

// ExitFitBF is called when production FitBF is exited.
func (s *BaseCmdListener) ExitFitBF(ctx *FitBFContext) {}

// EnterUnit_B is called when production Unit_B is entered.
func (s *BaseCmdListener) EnterUnit_B(ctx *Unit_BContext) {}

// ExitUnit_B is called when production Unit_B is exited.
func (s *BaseCmdListener) ExitUnit_B(ctx *Unit_BContext) {}

// EnterUnit_K is called when production Unit_K is entered.
func (s *BaseCmdListener) EnterUnit_K(ctx *Unit_KContext) {}

// ExitUnit_K is called when production Unit_K is exited.
func (s *BaseCmdListener) ExitUnit_K(ctx *Unit_KContext) {}

// EnterUnit_M is called when production Unit_M is entered.
func (s *BaseCmdListener) EnterUnit_M(ctx *Unit_MContext) {}

// ExitUnit_M is called when production Unit_M is exited.
func (s *BaseCmdListener) ExitUnit_M(ctx *Unit_MContext) {}

// EnterName_P is called when production Name_P is entered.
func (s *BaseCmdListener) EnterName_P(ctx *Name_PContext) {}

// ExitName_P is called when production Name_P is exited.
func (s *BaseCmdListener) ExitName_P(ctx *Name_PContext) {}

// EnterName_S is called when production Name_S is entered.
func (s *BaseCmdListener) EnterName_S(ctx *Name_SContext) {}

// ExitName_S is called when production Name_S is exited.
func (s *BaseCmdListener) ExitName_S(ctx *Name_SContext) {}

// EnterUsr_P is called when production Usr_P is entered.
func (s *BaseCmdListener) EnterUsr_P(ctx *Usr_PContext) {}

// ExitUsr_P is called when production Usr_P is exited.
func (s *BaseCmdListener) ExitUsr_P(ctx *Usr_PContext) {}

// EnterUsr_S is called when production Usr_S is entered.
func (s *BaseCmdListener) EnterUsr_S(ctx *Usr_SContext) {}

// ExitUsr_S is called when production Usr_S is exited.
func (s *BaseCmdListener) ExitUsr_S(ctx *Usr_SContext) {}

// EnterGrp_P is called when production Grp_P is entered.
func (s *BaseCmdListener) EnterGrp_P(ctx *Grp_PContext) {}

// ExitGrp_P is called when production Grp_P is exited.
func (s *BaseCmdListener) ExitGrp_P(ctx *Grp_PContext) {}

// EnterGrp_S is called when production Grp_S is entered.
func (s *BaseCmdListener) EnterGrp_S(ctx *Grp_SContext) {}

// ExitGrp_S is called when production Grp_S is exited.
func (s *BaseCmdListener) ExitGrp_S(ctx *Grp_SContext) {}

// EnterPass_P is called when production Pass_P is entered.
func (s *BaseCmdListener) EnterPass_P(ctx *Pass_PContext) {}

// ExitPass_P is called when production Pass_P is exited.
func (s *BaseCmdListener) ExitPass_P(ctx *Pass_PContext) {}

// EnterPwd_P is called when production Pwd_P is entered.
func (s *BaseCmdListener) EnterPwd_P(ctx *Pwd_PContext) {}

// ExitPwd_P is called when production Pwd_P is exited.
func (s *BaseCmdListener) ExitPwd_P(ctx *Pwd_PContext) {}

// EnterType_P is called when production Type_P is entered.
func (s *BaseCmdListener) EnterType_P(ctx *Type_PContext) {}

// ExitType_P is called when production Type_P is exited.
func (s *BaseCmdListener) ExitType_P(ctx *Type_PContext) {}

// EnterType_L is called when production Type_L is entered.
func (s *BaseCmdListener) EnterType_L(ctx *Type_LContext) {}

// ExitType_L is called when production Type_L is exited.
func (s *BaseCmdListener) ExitType_L(ctx *Type_LContext) {}

// EnterType_E is called when production Type_E is entered.
func (s *BaseCmdListener) EnterType_E(ctx *Type_EContext) {}

// ExitType_E is called when production Type_E is exited.
func (s *BaseCmdListener) ExitType_E(ctx *Type_EContext) {}

// EnterType_Fast is called when production Type_Fast is entered.
func (s *BaseCmdListener) EnterType_Fast(ctx *Type_FastContext) {}

// ExitType_Fast is called when production Type_Fast is exited.
func (s *BaseCmdListener) ExitType_Fast(ctx *Type_FastContext) {}

// EnterType_Full is called when production Type_Full is entered.
func (s *BaseCmdListener) EnterType_Full(ctx *Type_FullContext) {}

// ExitType_Full is called when production Type_Full is exited.
func (s *BaseCmdListener) ExitType_Full(ctx *Type_FullContext) {}

// EnterId is called when production Id is entered.
func (s *BaseCmdListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BaseCmdListener) ExitId(ctx *IdContext) {}

// EnterCont_S is called when production Cont_S is entered.
func (s *BaseCmdListener) EnterCont_S(ctx *Cont_SContext) {}

// ExitCont_S is called when production Cont_S is exited.
func (s *BaseCmdListener) ExitCont_S(ctx *Cont_SContext) {}

// EnterCont_R is called when production Cont_R is entered.
func (s *BaseCmdListener) EnterCont_R(ctx *Cont_RContext) {}

// ExitCont_R is called when production Cont_R is exited.
func (s *BaseCmdListener) ExitCont_R(ctx *Cont_RContext) {}

// EnterRuta_S is called when production Ruta_S is entered.
func (s *BaseCmdListener) EnterRuta_S(ctx *Ruta_SContext) {}

// ExitRuta_S is called when production Ruta_S is exited.
func (s *BaseCmdListener) ExitRuta_S(ctx *Ruta_SContext) {}

// EnterRuta_R is called when production Ruta_R is entered.
func (s *BaseCmdListener) EnterRuta_R(ctx *Ruta_RContext) {}

// ExitRuta_R is called when production Ruta_R is exited.
func (s *BaseCmdListener) ExitRuta_R(ctx *Ruta_RContext) {}

// EnterPp is called when production Pp is entered.
func (s *BaseCmdListener) EnterPp(ctx *PpContext) {}

// ExitPp is called when production Pp is exited.
func (s *BaseCmdListener) ExitPp(ctx *PpContext) {}

// EnterRr is called when production Rr is entered.
func (s *BaseCmdListener) EnterRr(ctx *RrContext) {}

// ExitRr is called when production Rr is exited.
func (s *BaseCmdListener) ExitRr(ctx *RrContext) {}
