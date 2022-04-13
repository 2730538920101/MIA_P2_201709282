// Generated from /home/ec2-user/MIA_P2/Cmd.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CmdParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TOK_MKDISK=1, TOK_RMDISK=2, TOK_FDISK=3, TOK_MOUNT=4, TOK_MKFS=5, TOK_LOGIN=6, 
		TOK_LOGOUT=7, TOK_MKGRP=8, TOK_RMGRP=9, TOK_MKUSR=10, TOK_RMUSR=11, TOK_MKFILE=12, 
		TOK_MKDIR=13, TOK_EXEC=14, TOK_PAUSE=15, TOK_REP=16, TOK_PATH=17, TOK_FIT=18, 
		TOK_SIZE=19, TOK_UNIT=20, TOK_NAME=21, TOK_TYPE=22, TOK_ID=23, TOK_USUARIO=24, 
		TOK_PASSWORD=25, TOK_PWD=26, TOK_CONT=27, TOK_GRP=28, TOK_RUTA=29, TOK_R=30, 
		TOK_P=31, TOK_FIRST=32, TOK_WORST=33, TOK_BEST=34, TOK_KB=35, TOK_MB=36, 
		TOK_BYTES=37, TOK_PRIMARIA=38, TOK_EXTENDIDA=39, TOK_LOGICA=40, TOK_FAST=41, 
		TOK_FULL=42, TOK_CADENA=43, TOK_NUMERO=44, TOK_IDENTIFICADOR=45, TOK_CAMINO=46, 
		TOK_PALABRA=47, TOK_IGUAL=48, COMENTARIOS=49, WHITESPACE=50;
	public static final int
		RULE_start = 0, RULE_comando = 1, RULE_comando_estado = 2, RULE_param_list = 3, 
		RULE_param = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "comando", "comando_estado", "param_list", "param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", 
			"TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", 
			"TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", "TOK_PATH", 
			"TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", "TOK_ID", 
			"TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", "TOK_RUTA", 
			"TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", "TOK_KB", "TOK_MB", 
			"TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FAST", 
			"TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", "TOK_CAMINO", 
			"TOK_PALABRA", "TOK_IGUAL", "COMENTARIOS", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Cmd.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CmdParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ComandoContext comando() {
			return getRuleContext(ComandoContext.class,0);
		}
		public TerminalNode EOF() { return getToken(CmdParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(10);
			comando();
			setState(11);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComandoContext extends ParserRuleContext {
		public ComandoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comando; }
	 
		public ComandoContext() { }
		public void copyFrom(ComandoContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PauseContext extends ComandoContext {
		public Token com;
		public TerminalNode TOK_PAUSE() { return getToken(CmdParser.TOK_PAUSE, 0); }
		public PauseContext(ComandoContext ctx) { copyFrom(ctx); }
	}
	public static class LogoutContext extends ComandoContext {
		public Token com;
		public TerminalNode TOK_LOGOUT() { return getToken(CmdParser.TOK_LOGOUT, 0); }
		public LogoutContext(ComandoContext ctx) { copyFrom(ctx); }
	}
	public static class AddCommandContext extends ComandoContext {
		public Comando_estadoContext comando_estado() {
			return getRuleContext(Comando_estadoContext.class,0);
		}
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public AddCommandContext(ComandoContext ctx) { copyFrom(ctx); }
	}

	public final ComandoContext comando() throws RecognitionException {
		ComandoContext _localctx = new ComandoContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_comando);
		try {
			setState(18);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TOK_MKDISK:
			case TOK_RMDISK:
			case TOK_FDISK:
			case TOK_MOUNT:
			case TOK_MKFS:
			case TOK_LOGIN:
			case TOK_MKGRP:
			case TOK_RMGRP:
			case TOK_MKUSR:
			case TOK_RMUSR:
			case TOK_MKFILE:
			case TOK_MKDIR:
			case TOK_EXEC:
			case TOK_REP:
				_localctx = new AddCommandContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(13);
				comando_estado();
				setState(14);
				param_list(0);
				}
				break;
			case TOK_PAUSE:
				_localctx = new PauseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(16);
				((PauseContext)_localctx).com = match(TOK_PAUSE);
				}
				break;
			case TOK_LOGOUT:
				_localctx = new LogoutContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(17);
				((LogoutContext)_localctx).com = match(TOK_LOGOUT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Comando_estadoContext extends ParserRuleContext {
		public Comando_estadoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comando_estado; }
	 
		public Comando_estadoContext() { }
		public void copyFrom(Comando_estadoContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class MkdirContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKDIR() { return getToken(CmdParser.TOK_MKDIR, 0); }
		public MkdirContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MkfileContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKFILE() { return getToken(CmdParser.TOK_MKFILE, 0); }
		public MkfileContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MkdiskContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKDISK() { return getToken(CmdParser.TOK_MKDISK, 0); }
		public MkdiskContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class RmdiskContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_RMDISK() { return getToken(CmdParser.TOK_RMDISK, 0); }
		public RmdiskContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MountContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MOUNT() { return getToken(CmdParser.TOK_MOUNT, 0); }
		public MountContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MkgrpContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKGRP() { return getToken(CmdParser.TOK_MKGRP, 0); }
		public MkgrpContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class FdiskContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_FDISK() { return getToken(CmdParser.TOK_FDISK, 0); }
		public FdiskContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class LoginContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_LOGIN() { return getToken(CmdParser.TOK_LOGIN, 0); }
		public LoginContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MkfsContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKFS() { return getToken(CmdParser.TOK_MKFS, 0); }
		public MkfsContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class RmusrContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_RMUSR() { return getToken(CmdParser.TOK_RMUSR, 0); }
		public RmusrContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class ExecContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_EXEC() { return getToken(CmdParser.TOK_EXEC, 0); }
		public ExecContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class MkusrContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_MKUSR() { return getToken(CmdParser.TOK_MKUSR, 0); }
		public MkusrContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class RmgrpContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_RMGRP() { return getToken(CmdParser.TOK_RMGRP, 0); }
		public RmgrpContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}
	public static class RepContext extends Comando_estadoContext {
		public Token com;
		public TerminalNode TOK_REP() { return getToken(CmdParser.TOK_REP, 0); }
		public RepContext(Comando_estadoContext ctx) { copyFrom(ctx); }
	}

	public final Comando_estadoContext comando_estado() throws RecognitionException {
		Comando_estadoContext _localctx = new Comando_estadoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_comando_estado);
		try {
			setState(34);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TOK_MKDISK:
				_localctx = new MkdiskContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(20);
				((MkdiskContext)_localctx).com = match(TOK_MKDISK);
				}
				break;
			case TOK_RMDISK:
				_localctx = new RmdiskContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(21);
				((RmdiskContext)_localctx).com = match(TOK_RMDISK);
				}
				break;
			case TOK_FDISK:
				_localctx = new FdiskContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(22);
				((FdiskContext)_localctx).com = match(TOK_FDISK);
				}
				break;
			case TOK_MOUNT:
				_localctx = new MountContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(23);
				((MountContext)_localctx).com = match(TOK_MOUNT);
				}
				break;
			case TOK_MKFS:
				_localctx = new MkfsContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(24);
				((MkfsContext)_localctx).com = match(TOK_MKFS);
				}
				break;
			case TOK_LOGIN:
				_localctx = new LoginContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(25);
				((LoginContext)_localctx).com = match(TOK_LOGIN);
				}
				break;
			case TOK_MKGRP:
				_localctx = new MkgrpContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(26);
				((MkgrpContext)_localctx).com = match(TOK_MKGRP);
				}
				break;
			case TOK_RMGRP:
				_localctx = new RmgrpContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(27);
				((RmgrpContext)_localctx).com = match(TOK_RMGRP);
				}
				break;
			case TOK_MKUSR:
				_localctx = new MkusrContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(28);
				((MkusrContext)_localctx).com = match(TOK_MKUSR);
				}
				break;
			case TOK_RMUSR:
				_localctx = new RmusrContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(29);
				((RmusrContext)_localctx).com = match(TOK_RMUSR);
				}
				break;
			case TOK_MKFILE:
				_localctx = new MkfileContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(30);
				((MkfileContext)_localctx).com = match(TOK_MKFILE);
				}
				break;
			case TOK_MKDIR:
				_localctx = new MkdirContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(31);
				((MkdirContext)_localctx).com = match(TOK_MKDIR);
				}
				break;
			case TOK_EXEC:
				_localctx = new ExecContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(32);
				((ExecContext)_localctx).com = match(TOK_EXEC);
				}
				break;
			case TOK_REP:
				_localctx = new RepContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(33);
				((RepContext)_localctx).com = match(TOK_REP);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Param_listContext extends ParserRuleContext {
		public ParamContext param() {
			return getRuleContext(ParamContext.class,0);
		}
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public Param_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param_list; }
	}

	public final Param_listContext param_list() throws RecognitionException {
		return param_list(0);
	}

	private Param_listContext param_list(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Param_listContext _localctx = new Param_listContext(_ctx, _parentState);
		Param_listContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_param_list, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(37);
			param();
			}
			_ctx.stop = _input.LT(-1);
			setState(43);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Param_listContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_param_list);
					setState(39);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(40);
					param();
					}
					} 
				}
				setState(45);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParamContext extends ParserRuleContext {
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	 
		public ParamContext() { }
		public void copyFrom(ParamContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class Cont_RContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_CONT() { return getToken(CmdParser.TOK_CONT, 0); }
		public TerminalNode TOK_CAMINO() { return getToken(CmdParser.TOK_CAMINO, 0); }
		public Cont_RContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Cont_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_CONT() { return getToken(CmdParser.TOK_CONT, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Cont_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class SizeContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_SIZE() { return getToken(CmdParser.TOK_SIZE, 0); }
		public TerminalNode TOK_NUMERO() { return getToken(CmdParser.TOK_NUMERO, 0); }
		public SizeContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Grp_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_GRP() { return getToken(CmdParser.TOK_GRP, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Grp_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Path_RContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_PATH() { return getToken(CmdParser.TOK_PATH, 0); }
		public TerminalNode TOK_CAMINO() { return getToken(CmdParser.TOK_CAMINO, 0); }
		public Path_RContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Grp_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_GRP() { return getToken(CmdParser.TOK_GRP, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public Grp_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Path_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_PATH() { return getToken(CmdParser.TOK_PATH, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Path_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class FitWFContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_FIT() { return getToken(CmdParser.TOK_FIT, 0); }
		public TerminalNode TOK_WORST() { return getToken(CmdParser.TOK_WORST, 0); }
		public FitWFContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Pass_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_PASSWORD() { return getToken(CmdParser.TOK_PASSWORD, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public Pass_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Type_LContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_LOGICA() { return getToken(CmdParser.TOK_LOGICA, 0); }
		public Type_LContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Type_FastContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_FAST() { return getToken(CmdParser.TOK_FAST, 0); }
		public Type_FastContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Type_EContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_EXTENDIDA() { return getToken(CmdParser.TOK_EXTENDIDA, 0); }
		public Type_EContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Name_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_NAME() { return getToken(CmdParser.TOK_NAME, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Name_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Name_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_NAME() { return getToken(CmdParser.TOK_NAME, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public Name_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Type_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_PRIMARIA() { return getToken(CmdParser.TOK_PRIMARIA, 0); }
		public Type_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class PpContext extends ParamContext {
		public Token par;
		public TerminalNode TOK_P() { return getToken(CmdParser.TOK_P, 0); }
		public PpContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class RrContext extends ParamContext {
		public Token par;
		public TerminalNode TOK_R() { return getToken(CmdParser.TOK_R, 0); }
		public RrContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Usr_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_USUARIO() { return getToken(CmdParser.TOK_USUARIO, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public Usr_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Type_FullContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_FULL() { return getToken(CmdParser.TOK_FULL, 0); }
		public Type_FullContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Usr_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_USUARIO() { return getToken(CmdParser.TOK_USUARIO, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Usr_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Unit_BContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_UNIT() { return getToken(CmdParser.TOK_UNIT, 0); }
		public TerminalNode TOK_BYTES() { return getToken(CmdParser.TOK_BYTES, 0); }
		public Unit_BContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Pwd_PContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_PWD() { return getToken(CmdParser.TOK_PWD, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public Pwd_PContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Ruta_SContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_RUTA() { return getToken(CmdParser.TOK_RUTA, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public Ruta_SContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Ruta_RContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_RUTA() { return getToken(CmdParser.TOK_RUTA, 0); }
		public TerminalNode TOK_CAMINO() { return getToken(CmdParser.TOK_CAMINO, 0); }
		public Ruta_RContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Unit_MContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_UNIT() { return getToken(CmdParser.TOK_UNIT, 0); }
		public TerminalNode TOK_MB() { return getToken(CmdParser.TOK_MB, 0); }
		public Unit_MContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class FitFFContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_FIT() { return getToken(CmdParser.TOK_FIT, 0); }
		public TerminalNode TOK_FIRST() { return getToken(CmdParser.TOK_FIRST, 0); }
		public FitFFContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class Unit_KContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_UNIT() { return getToken(CmdParser.TOK_UNIT, 0); }
		public TerminalNode TOK_KB() { return getToken(CmdParser.TOK_KB, 0); }
		public Unit_KContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class IdContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_ID() { return getToken(CmdParser.TOK_ID, 0); }
		public TerminalNode TOK_IDENTIFICADOR() { return getToken(CmdParser.TOK_IDENTIFICADOR, 0); }
		public IdContext(ParamContext ctx) { copyFrom(ctx); }
	}
	public static class FitBFContext extends ParamContext {
		public Token par;
		public Token res;
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_FIT() { return getToken(CmdParser.TOK_FIT, 0); }
		public TerminalNode TOK_BEST() { return getToken(CmdParser.TOK_BEST, 0); }
		public FitBFContext(ParamContext ctx) { copyFrom(ctx); }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_param);
		try {
			setState(129);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new SizeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				((SizeContext)_localctx).par = match(TOK_SIZE);
				setState(47);
				match(TOK_IGUAL);
				setState(48);
				((SizeContext)_localctx).res = match(TOK_NUMERO);
				}
				break;
			case 2:
				_localctx = new Path_RContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(49);
				((Path_RContext)_localctx).par = match(TOK_PATH);
				setState(50);
				match(TOK_IGUAL);
				setState(51);
				((Path_RContext)_localctx).res = match(TOK_CAMINO);
				}
				break;
			case 3:
				_localctx = new Path_SContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
				((Path_SContext)_localctx).par = match(TOK_PATH);
				setState(53);
				match(TOK_IGUAL);
				setState(54);
				((Path_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 4:
				_localctx = new FitFFContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(55);
				((FitFFContext)_localctx).par = match(TOK_FIT);
				setState(56);
				match(TOK_IGUAL);
				setState(57);
				((FitFFContext)_localctx).res = match(TOK_FIRST);
				}
				break;
			case 5:
				_localctx = new FitWFContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(58);
				((FitWFContext)_localctx).par = match(TOK_FIT);
				setState(59);
				match(TOK_IGUAL);
				setState(60);
				((FitWFContext)_localctx).res = match(TOK_WORST);
				}
				break;
			case 6:
				_localctx = new FitBFContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(61);
				((FitBFContext)_localctx).par = match(TOK_FIT);
				setState(62);
				match(TOK_IGUAL);
				setState(63);
				((FitBFContext)_localctx).res = match(TOK_BEST);
				}
				break;
			case 7:
				_localctx = new Unit_BContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(64);
				((Unit_BContext)_localctx).par = match(TOK_UNIT);
				setState(65);
				match(TOK_IGUAL);
				setState(66);
				((Unit_BContext)_localctx).res = match(TOK_BYTES);
				}
				break;
			case 8:
				_localctx = new Unit_KContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(67);
				((Unit_KContext)_localctx).par = match(TOK_UNIT);
				setState(68);
				match(TOK_IGUAL);
				setState(69);
				((Unit_KContext)_localctx).res = match(TOK_KB);
				}
				break;
			case 9:
				_localctx = new Unit_MContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(70);
				((Unit_MContext)_localctx).par = match(TOK_UNIT);
				setState(71);
				match(TOK_IGUAL);
				setState(72);
				((Unit_MContext)_localctx).res = match(TOK_MB);
				}
				break;
			case 10:
				_localctx = new Name_PContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(73);
				((Name_PContext)_localctx).par = match(TOK_NAME);
				setState(74);
				match(TOK_IGUAL);
				setState(75);
				((Name_PContext)_localctx).res = match(TOK_PALABRA);
				}
				break;
			case 11:
				_localctx = new Name_SContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(76);
				((Name_SContext)_localctx).par = match(TOK_NAME);
				setState(77);
				match(TOK_IGUAL);
				setState(78);
				((Name_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 12:
				_localctx = new Usr_PContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(79);
				((Usr_PContext)_localctx).par = match(TOK_USUARIO);
				setState(80);
				match(TOK_IGUAL);
				setState(81);
				((Usr_PContext)_localctx).res = match(TOK_PALABRA);
				}
				break;
			case 13:
				_localctx = new Usr_SContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(82);
				((Usr_SContext)_localctx).par = match(TOK_USUARIO);
				setState(83);
				match(TOK_IGUAL);
				setState(84);
				((Usr_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 14:
				_localctx = new Grp_PContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(85);
				((Grp_PContext)_localctx).par = match(TOK_GRP);
				setState(86);
				match(TOK_IGUAL);
				setState(87);
				((Grp_PContext)_localctx).res = match(TOK_PALABRA);
				}
				break;
			case 15:
				_localctx = new Grp_SContext(_localctx);
				enterOuterAlt(_localctx, 15);
				{
				setState(88);
				((Grp_SContext)_localctx).par = match(TOK_GRP);
				setState(89);
				match(TOK_IGUAL);
				setState(90);
				((Grp_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 16:
				_localctx = new Pass_PContext(_localctx);
				enterOuterAlt(_localctx, 16);
				{
				setState(91);
				((Pass_PContext)_localctx).par = match(TOK_PASSWORD);
				setState(92);
				match(TOK_IGUAL);
				setState(93);
				((Pass_PContext)_localctx).res = match(TOK_PALABRA);
				}
				break;
			case 17:
				_localctx = new Pwd_PContext(_localctx);
				enterOuterAlt(_localctx, 17);
				{
				setState(94);
				((Pwd_PContext)_localctx).par = match(TOK_PWD);
				setState(95);
				match(TOK_IGUAL);
				setState(96);
				((Pwd_PContext)_localctx).res = match(TOK_PALABRA);
				}
				break;
			case 18:
				_localctx = new Type_PContext(_localctx);
				enterOuterAlt(_localctx, 18);
				{
				setState(97);
				((Type_PContext)_localctx).par = match(TOK_TYPE);
				setState(98);
				match(TOK_IGUAL);
				setState(99);
				((Type_PContext)_localctx).res = match(TOK_PRIMARIA);
				}
				break;
			case 19:
				_localctx = new Type_LContext(_localctx);
				enterOuterAlt(_localctx, 19);
				{
				setState(100);
				((Type_LContext)_localctx).par = match(TOK_TYPE);
				setState(101);
				match(TOK_IGUAL);
				setState(102);
				((Type_LContext)_localctx).res = match(TOK_LOGICA);
				}
				break;
			case 20:
				_localctx = new Type_EContext(_localctx);
				enterOuterAlt(_localctx, 20);
				{
				setState(103);
				((Type_EContext)_localctx).par = match(TOK_TYPE);
				setState(104);
				match(TOK_IGUAL);
				setState(105);
				((Type_EContext)_localctx).res = match(TOK_EXTENDIDA);
				}
				break;
			case 21:
				_localctx = new Type_FastContext(_localctx);
				enterOuterAlt(_localctx, 21);
				{
				setState(106);
				((Type_FastContext)_localctx).par = match(TOK_TYPE);
				setState(107);
				match(TOK_IGUAL);
				setState(108);
				((Type_FastContext)_localctx).res = match(TOK_FAST);
				}
				break;
			case 22:
				_localctx = new Type_FullContext(_localctx);
				enterOuterAlt(_localctx, 22);
				{
				setState(109);
				((Type_FullContext)_localctx).par = match(TOK_TYPE);
				setState(110);
				match(TOK_IGUAL);
				setState(111);
				((Type_FullContext)_localctx).res = match(TOK_FULL);
				}
				break;
			case 23:
				_localctx = new IdContext(_localctx);
				enterOuterAlt(_localctx, 23);
				{
				setState(112);
				((IdContext)_localctx).par = match(TOK_ID);
				setState(113);
				match(TOK_IGUAL);
				setState(114);
				((IdContext)_localctx).res = match(TOK_IDENTIFICADOR);
				}
				break;
			case 24:
				_localctx = new Cont_SContext(_localctx);
				enterOuterAlt(_localctx, 24);
				{
				setState(115);
				((Cont_SContext)_localctx).par = match(TOK_CONT);
				setState(116);
				match(TOK_IGUAL);
				setState(117);
				((Cont_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 25:
				_localctx = new Cont_RContext(_localctx);
				enterOuterAlt(_localctx, 25);
				{
				setState(118);
				((Cont_RContext)_localctx).par = match(TOK_CONT);
				setState(119);
				match(TOK_IGUAL);
				setState(120);
				((Cont_RContext)_localctx).res = match(TOK_CAMINO);
				}
				break;
			case 26:
				_localctx = new Ruta_SContext(_localctx);
				enterOuterAlt(_localctx, 26);
				{
				setState(121);
				((Ruta_SContext)_localctx).par = match(TOK_RUTA);
				setState(122);
				match(TOK_IGUAL);
				setState(123);
				((Ruta_SContext)_localctx).res = match(TOK_CADENA);
				}
				break;
			case 27:
				_localctx = new Ruta_RContext(_localctx);
				enterOuterAlt(_localctx, 27);
				{
				setState(124);
				((Ruta_RContext)_localctx).par = match(TOK_RUTA);
				setState(125);
				match(TOK_IGUAL);
				setState(126);
				((Ruta_RContext)_localctx).res = match(TOK_CAMINO);
				}
				break;
			case 28:
				_localctx = new PpContext(_localctx);
				enterOuterAlt(_localctx, 28);
				{
				setState(127);
				((PpContext)_localctx).par = match(TOK_P);
				}
				break;
			case 29:
				_localctx = new RrContext(_localctx);
				enterOuterAlt(_localctx, 29);
				{
				setState(128);
				((RrContext)_localctx).par = match(TOK_R);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 3:
			return param_list_sempred((Param_listContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean param_list_sempred(Param_listContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\64\u0086\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\5\3\25"+
		"\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4%\n\4"+
		"\3\5\3\5\3\5\3\5\3\5\7\5,\n\5\f\5\16\5/\13\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u0084\n\6\3\6\2\3\b\7\2\4\6\b\n\2\2\2\u00ac"+
		"\2\f\3\2\2\2\4\24\3\2\2\2\6$\3\2\2\2\b&\3\2\2\2\n\u0083\3\2\2\2\f\r\5"+
		"\4\3\2\r\16\7\2\2\3\16\3\3\2\2\2\17\20\5\6\4\2\20\21\5\b\5\2\21\25\3\2"+
		"\2\2\22\25\7\21\2\2\23\25\7\t\2\2\24\17\3\2\2\2\24\22\3\2\2\2\24\23\3"+
		"\2\2\2\25\5\3\2\2\2\26%\7\3\2\2\27%\7\4\2\2\30%\7\5\2\2\31%\7\6\2\2\32"+
		"%\7\7\2\2\33%\7\b\2\2\34%\7\n\2\2\35%\7\13\2\2\36%\7\f\2\2\37%\7\r\2\2"+
		" %\7\16\2\2!%\7\17\2\2\"%\7\20\2\2#%\7\22\2\2$\26\3\2\2\2$\27\3\2\2\2"+
		"$\30\3\2\2\2$\31\3\2\2\2$\32\3\2\2\2$\33\3\2\2\2$\34\3\2\2\2$\35\3\2\2"+
		"\2$\36\3\2\2\2$\37\3\2\2\2$ \3\2\2\2$!\3\2\2\2$\"\3\2\2\2$#\3\2\2\2%\7"+
		"\3\2\2\2&\'\b\5\1\2\'(\5\n\6\2(-\3\2\2\2)*\f\4\2\2*,\5\n\6\2+)\3\2\2\2"+
		",/\3\2\2\2-+\3\2\2\2-.\3\2\2\2.\t\3\2\2\2/-\3\2\2\2\60\61\7\25\2\2\61"+
		"\62\7\62\2\2\62\u0084\7.\2\2\63\64\7\23\2\2\64\65\7\62\2\2\65\u0084\7"+
		"\60\2\2\66\67\7\23\2\2\678\7\62\2\28\u0084\7-\2\29:\7\24\2\2:;\7\62\2"+
		"\2;\u0084\7\"\2\2<=\7\24\2\2=>\7\62\2\2>\u0084\7#\2\2?@\7\24\2\2@A\7\62"+
		"\2\2A\u0084\7$\2\2BC\7\26\2\2CD\7\62\2\2D\u0084\7\'\2\2EF\7\26\2\2FG\7"+
		"\62\2\2G\u0084\7%\2\2HI\7\26\2\2IJ\7\62\2\2J\u0084\7&\2\2KL\7\27\2\2L"+
		"M\7\62\2\2M\u0084\7\61\2\2NO\7\27\2\2OP\7\62\2\2P\u0084\7-\2\2QR\7\32"+
		"\2\2RS\7\62\2\2S\u0084\7\61\2\2TU\7\32\2\2UV\7\62\2\2V\u0084\7-\2\2WX"+
		"\7\36\2\2XY\7\62\2\2Y\u0084\7\61\2\2Z[\7\36\2\2[\\\7\62\2\2\\\u0084\7"+
		"-\2\2]^\7\33\2\2^_\7\62\2\2_\u0084\7\61\2\2`a\7\34\2\2ab\7\62\2\2b\u0084"+
		"\7\61\2\2cd\7\30\2\2de\7\62\2\2e\u0084\7(\2\2fg\7\30\2\2gh\7\62\2\2h\u0084"+
		"\7*\2\2ij\7\30\2\2jk\7\62\2\2k\u0084\7)\2\2lm\7\30\2\2mn\7\62\2\2n\u0084"+
		"\7+\2\2op\7\30\2\2pq\7\62\2\2q\u0084\7,\2\2rs\7\31\2\2st\7\62\2\2t\u0084"+
		"\7/\2\2uv\7\35\2\2vw\7\62\2\2w\u0084\7-\2\2xy\7\35\2\2yz\7\62\2\2z\u0084"+
		"\7\60\2\2{|\7\37\2\2|}\7\62\2\2}\u0084\7-\2\2~\177\7\37\2\2\177\u0080"+
		"\7\62\2\2\u0080\u0084\7\60\2\2\u0081\u0084\7!\2\2\u0082\u0084\7 \2\2\u0083"+
		"\60\3\2\2\2\u0083\63\3\2\2\2\u0083\66\3\2\2\2\u00839\3\2\2\2\u0083<\3"+
		"\2\2\2\u0083?\3\2\2\2\u0083B\3\2\2\2\u0083E\3\2\2\2\u0083H\3\2\2\2\u0083"+
		"K\3\2\2\2\u0083N\3\2\2\2\u0083Q\3\2\2\2\u0083T\3\2\2\2\u0083W\3\2\2\2"+
		"\u0083Z\3\2\2\2\u0083]\3\2\2\2\u0083`\3\2\2\2\u0083c\3\2\2\2\u0083f\3"+
		"\2\2\2\u0083i\3\2\2\2\u0083l\3\2\2\2\u0083o\3\2\2\2\u0083r\3\2\2\2\u0083"+
		"u\3\2\2\2\u0083x\3\2\2\2\u0083{\3\2\2\2\u0083~\3\2\2\2\u0083\u0081\3\2"+
		"\2\2\u0083\u0082\3\2\2\2\u0084\13\3\2\2\2\6\24$-\u0083";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}