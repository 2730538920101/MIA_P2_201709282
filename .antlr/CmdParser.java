// Generated from /home/ec2-user/MIA_P2/Cmd.g4 by ANTLR 4.8

        import "../ast"
        import arrayList "github.com/colegno/arraylist"

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
		TOK_IGUAL=1, TOK_MKDISK=2, TOK_RMDISK=3, TOK_FDISK=4, TOK_MOUNT=5, TOK_MKFS=6, 
		TOK_LOGIN=7, TOK_LOGOUT=8, TOK_MKGRP=9, TOK_RMGRP=10, TOK_MKUSR=11, TOK_RMUSR=12, 
		TOK_MKFILE=13, TOK_MKDIR=14, TOK_EXEC=15, TOK_PAUSE=16, TOK_REP=17, TOK_PATH=18, 
		TOK_FIT=19, TOK_SIZE=20, TOK_UNIT=21, TOK_NAME=22, TOK_TYPE=23, TOK_ID=24, 
		TOK_USUARIO=25, TOK_PASSWORD=26, TOK_PWD=27, TOK_CONT=28, TOK_GRP=29, 
		TOK_RUTA=30, TOK_R=31, TOK_P=32, TOK_KB=33, TOK_MB=34, TOK_BYTES=35, TOK_PRIMARIA=36, 
		TOK_EXTENDIDA=37, TOK_LOGICA=38, TOK_FIRST=39, TOK_WORST=40, TOK_BEST=41, 
		TOK_FAST=42, TOK_FULL=43, TOK_CADENA=44, TOK_NUMERO=45, TOK_IDENTIFICADOR=46, 
		TOK_CAMINO=47, TOK_PALABRA=48, COMENTARIOS=49, WHITESPACE=50;
	public static final int
		RULE_start = 0, RULE_comandList = 1, RULE_comando = 2, RULE_param_list = 3, 
		RULE_param = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "comandList", "comando", "param_list", "param"
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
			null, "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", 
			"TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", 
			"TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", 
			"TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", 
			"TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", 
			"TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB", "TOK_BYTES", "TOK_PRIMARIA", 
			"TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST", "TOK_WORST", "TOK_BEST", 
			"TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", 
			"TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE"
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
		public ast.Ast ast;
		public ComandListContext comandList;
		public ComandListContext comandList() {
			return getRuleContext(ComandListContext.class,0);
		}
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
			((StartContext)_localctx).comandList = comandList(0);
			_localctx.ast = ast.NewAst(((StartContext)_localctx).comandList.lista)
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

	public static class ComandListContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ComandListContext AUXCLIST;
		public ComandoContext comando;
		public ComandoContext comando() {
			return getRuleContext(ComandoContext.class,0);
		}
		public ComandListContext comandList() {
			return getRuleContext(ComandListContext.class,0);
		}
		public ComandListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comandList; }
	}

	public final ComandListContext comandList() throws RecognitionException {
		return comandList(0);
	}

	private ComandListContext comandList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ComandListContext _localctx = new ComandListContext(_ctx, _parentState);
		ComandListContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_comandList, _p);
		 _localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(14);
			((ComandListContext)_localctx).comando = comando();

			                        _localctx.lista.Add(((ComandListContext)_localctx).comando.com)
			                   
			}
			_ctx.stop = _input.LT(-1);
			setState(23);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ComandListContext(_parentctx, _parentState);
					_localctx.AUXCLIST = _prevctx;
					_localctx.AUXCLIST = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_comandList);
					setState(17);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(18);
					((ComandListContext)_localctx).comando = comando();

					                                  ((ComandListContext)_localctx).AUXCLIST.lista.Add(((ComandListContext)_localctx).comando.com)
					                                  _localctx.lista = ((ComandListContext)_localctx).AUXCLIST.lista
					                             
					}
					} 
				}
				setState(25);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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

	public static class ComandoContext extends ParserRuleContext {
		public ast.Command com;
		public TerminalNode TOK_MKDISK() { return getToken(CmdParser.TOK_MKDISK, 0); }
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public TerminalNode TOK_RMDISK() { return getToken(CmdParser.TOK_RMDISK, 0); }
		public TerminalNode TOK_FDISK() { return getToken(CmdParser.TOK_FDISK, 0); }
		public TerminalNode TOK_MOUNT() { return getToken(CmdParser.TOK_MOUNT, 0); }
		public TerminalNode TOK_MKFS() { return getToken(CmdParser.TOK_MKFS, 0); }
		public TerminalNode TOK_LOGIN() { return getToken(CmdParser.TOK_LOGIN, 0); }
		public TerminalNode TOK_MKGRP() { return getToken(CmdParser.TOK_MKGRP, 0); }
		public TerminalNode TOK_RMGRP() { return getToken(CmdParser.TOK_RMGRP, 0); }
		public TerminalNode TOK_MKUSR() { return getToken(CmdParser.TOK_MKUSR, 0); }
		public TerminalNode TOK_RMUSR() { return getToken(CmdParser.TOK_RMUSR, 0); }
		public TerminalNode TOK_MKFILE() { return getToken(CmdParser.TOK_MKFILE, 0); }
		public TerminalNode TOK_MKDIR() { return getToken(CmdParser.TOK_MKDIR, 0); }
		public TerminalNode TOK_EXEC() { return getToken(CmdParser.TOK_EXEC, 0); }
		public TerminalNode TOK_REP() { return getToken(CmdParser.TOK_REP, 0); }
		public TerminalNode TOK_PAUSE() { return getToken(CmdParser.TOK_PAUSE, 0); }
		public TerminalNode TOK_LOGOUT() { return getToken(CmdParser.TOK_LOGOUT, 0); }
		public ComandoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comando; }
	}

	public final ComandoContext comando() throws RecognitionException {
		ComandoContext _localctx = new ComandoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_comando);
		try {
			setState(56);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TOK_MKDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				match(TOK_MKDISK);
				setState(27);
				param_list(0);
				}
				break;
			case TOK_RMDISK:
				enterOuterAlt(_localctx, 2);
				{
				setState(28);
				match(TOK_RMDISK);
				setState(29);
				param_list(0);
				}
				break;
			case TOK_FDISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(30);
				match(TOK_FDISK);
				setState(31);
				param_list(0);
				}
				break;
			case TOK_MOUNT:
				enterOuterAlt(_localctx, 4);
				{
				setState(32);
				match(TOK_MOUNT);
				setState(33);
				param_list(0);
				}
				break;
			case TOK_MKFS:
				enterOuterAlt(_localctx, 5);
				{
				setState(34);
				match(TOK_MKFS);
				setState(35);
				param_list(0);
				}
				break;
			case TOK_LOGIN:
				enterOuterAlt(_localctx, 6);
				{
				setState(36);
				match(TOK_LOGIN);
				setState(37);
				param_list(0);
				}
				break;
			case TOK_MKGRP:
				enterOuterAlt(_localctx, 7);
				{
				setState(38);
				match(TOK_MKGRP);
				setState(39);
				param_list(0);
				}
				break;
			case TOK_RMGRP:
				enterOuterAlt(_localctx, 8);
				{
				setState(40);
				match(TOK_RMGRP);
				setState(41);
				param_list(0);
				}
				break;
			case TOK_MKUSR:
				enterOuterAlt(_localctx, 9);
				{
				setState(42);
				match(TOK_MKUSR);
				setState(43);
				param_list(0);
				}
				break;
			case TOK_RMUSR:
				enterOuterAlt(_localctx, 10);
				{
				setState(44);
				match(TOK_RMUSR);
				setState(45);
				param_list(0);
				}
				break;
			case TOK_MKFILE:
				enterOuterAlt(_localctx, 11);
				{
				setState(46);
				match(TOK_MKFILE);
				setState(47);
				param_list(0);
				}
				break;
			case TOK_MKDIR:
				enterOuterAlt(_localctx, 12);
				{
				setState(48);
				match(TOK_MKDIR);
				setState(49);
				param_list(0);
				}
				break;
			case TOK_EXEC:
				enterOuterAlt(_localctx, 13);
				{
				setState(50);
				match(TOK_EXEC);
				setState(51);
				param_list(0);
				}
				break;
			case TOK_REP:
				enterOuterAlt(_localctx, 14);
				{
				setState(52);
				match(TOK_REP);
				setState(53);
				param_list(0);
				}
				break;
			case TOK_PAUSE:
				enterOuterAlt(_localctx, 15);
				{
				setState(54);
				match(TOK_PAUSE);
				}
				break;
			case TOK_LOGOUT:
				enterOuterAlt(_localctx, 16);
				{
				setState(55);
				match(TOK_LOGOUT);
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
		public *arrayList.List lista;
		public Param_listContext AUXPLIST;
		public ParamContext param;
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
		_localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(59);
			((Param_listContext)_localctx).param = param();

			                   _localctx.lista.Add(((Param_listContext)_localctx).param.par)
			                
			}
			_ctx.stop = _input.LT(-1);
			setState(68);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Param_listContext(_parentctx, _parentState);
					_localctx.AUXPLIST = _prevctx;
					_localctx.AUXPLIST = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_param_list);
					setState(62);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(63);
					((Param_listContext)_localctx).param = param();

					                             ((Param_listContext)_localctx).AUXPLIST.lista.Add(((Param_listContext)_localctx).param.par)
					                             _localctx.lista = ((Param_listContext)_localctx).AUXPLIST.lista
					                          
					}
					} 
				}
				setState(70);
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
		public ast.Parametro par;
		public TerminalNode TOK_SIZE() { return getToken(CmdParser.TOK_SIZE, 0); }
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_NUMERO() { return getToken(CmdParser.TOK_NUMERO, 0); }
		public TerminalNode TOK_PATH() { return getToken(CmdParser.TOK_PATH, 0); }
		public TerminalNode TOK_CAMINO() { return getToken(CmdParser.TOK_CAMINO, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public TerminalNode TOK_FIT() { return getToken(CmdParser.TOK_FIT, 0); }
		public TerminalNode TOK_FIRST() { return getToken(CmdParser.TOK_FIRST, 0); }
		public TerminalNode TOK_WORST() { return getToken(CmdParser.TOK_WORST, 0); }
		public TerminalNode TOK_BEST() { return getToken(CmdParser.TOK_BEST, 0); }
		public TerminalNode TOK_UNIT() { return getToken(CmdParser.TOK_UNIT, 0); }
		public TerminalNode TOK_BYTES() { return getToken(CmdParser.TOK_BYTES, 0); }
		public TerminalNode TOK_KB() { return getToken(CmdParser.TOK_KB, 0); }
		public TerminalNode TOK_MB() { return getToken(CmdParser.TOK_MB, 0); }
		public TerminalNode TOK_NAME() { return getToken(CmdParser.TOK_NAME, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public TerminalNode TOK_USUARIO() { return getToken(CmdParser.TOK_USUARIO, 0); }
		public TerminalNode TOK_GRP() { return getToken(CmdParser.TOK_GRP, 0); }
		public TerminalNode TOK_PASSWORD() { return getToken(CmdParser.TOK_PASSWORD, 0); }
		public TerminalNode TOK_PWD() { return getToken(CmdParser.TOK_PWD, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_PRIMARIA() { return getToken(CmdParser.TOK_PRIMARIA, 0); }
		public TerminalNode TOK_LOGICA() { return getToken(CmdParser.TOK_LOGICA, 0); }
		public TerminalNode TOK_EXTENDIDA() { return getToken(CmdParser.TOK_EXTENDIDA, 0); }
		public TerminalNode TOK_FAST() { return getToken(CmdParser.TOK_FAST, 0); }
		public TerminalNode TOK_FULL() { return getToken(CmdParser.TOK_FULL, 0); }
		public TerminalNode TOK_ID() { return getToken(CmdParser.TOK_ID, 0); }
		public TerminalNode TOK_IDENTIFICADOR() { return getToken(CmdParser.TOK_IDENTIFICADOR, 0); }
		public TerminalNode TOK_CONT() { return getToken(CmdParser.TOK_CONT, 0); }
		public TerminalNode TOK_RUTA() { return getToken(CmdParser.TOK_RUTA, 0); }
		public TerminalNode TOK_P() { return getToken(CmdParser.TOK_P, 0); }
		public TerminalNode TOK_R() { return getToken(CmdParser.TOK_R, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_param);
		try {
			setState(154);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(71);
				match(TOK_SIZE);
				setState(72);
				match(TOK_IGUAL);
				setState(73);
				match(TOK_NUMERO);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				match(TOK_PATH);
				setState(75);
				match(TOK_IGUAL);
				setState(76);
				match(TOK_CAMINO);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				match(TOK_PATH);
				setState(78);
				match(TOK_IGUAL);
				setState(79);
				match(TOK_CADENA);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(80);
				match(TOK_FIT);
				setState(81);
				match(TOK_IGUAL);
				setState(82);
				match(TOK_FIRST);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(83);
				match(TOK_FIT);
				setState(84);
				match(TOK_IGUAL);
				setState(85);
				match(TOK_WORST);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(86);
				match(TOK_FIT);
				setState(87);
				match(TOK_IGUAL);
				setState(88);
				match(TOK_BEST);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(89);
				match(TOK_UNIT);
				setState(90);
				match(TOK_IGUAL);
				setState(91);
				match(TOK_BYTES);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(92);
				match(TOK_UNIT);
				setState(93);
				match(TOK_IGUAL);
				setState(94);
				match(TOK_KB);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(95);
				match(TOK_UNIT);
				setState(96);
				match(TOK_IGUAL);
				setState(97);
				match(TOK_MB);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(98);
				match(TOK_NAME);
				setState(99);
				match(TOK_IGUAL);
				setState(100);
				match(TOK_PALABRA);
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(101);
				match(TOK_NAME);
				setState(102);
				match(TOK_IGUAL);
				setState(103);
				match(TOK_CADENA);
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(104);
				match(TOK_USUARIO);
				setState(105);
				match(TOK_IGUAL);
				setState(106);
				match(TOK_PALABRA);
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(107);
				match(TOK_USUARIO);
				setState(108);
				match(TOK_IGUAL);
				setState(109);
				match(TOK_CADENA);
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(110);
				match(TOK_GRP);
				setState(111);
				match(TOK_IGUAL);
				setState(112);
				match(TOK_PALABRA);
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(113);
				match(TOK_GRP);
				setState(114);
				match(TOK_IGUAL);
				setState(115);
				match(TOK_CADENA);
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(116);
				match(TOK_PASSWORD);
				setState(117);
				match(TOK_IGUAL);
				setState(118);
				match(TOK_PALABRA);
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(119);
				match(TOK_PWD);
				setState(120);
				match(TOK_IGUAL);
				setState(121);
				match(TOK_PALABRA);
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(122);
				match(TOK_TYPE);
				setState(123);
				match(TOK_IGUAL);
				setState(124);
				match(TOK_PRIMARIA);
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(125);
				match(TOK_TYPE);
				setState(126);
				match(TOK_IGUAL);
				setState(127);
				match(TOK_LOGICA);
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(128);
				match(TOK_TYPE);
				setState(129);
				match(TOK_IGUAL);
				setState(130);
				match(TOK_EXTENDIDA);
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(131);
				match(TOK_TYPE);
				setState(132);
				match(TOK_IGUAL);
				setState(133);
				match(TOK_FAST);
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(134);
				match(TOK_TYPE);
				setState(135);
				match(TOK_IGUAL);
				setState(136);
				match(TOK_FULL);
				}
				break;
			case 23:
				enterOuterAlt(_localctx, 23);
				{
				setState(137);
				match(TOK_ID);
				setState(138);
				match(TOK_IGUAL);
				setState(139);
				match(TOK_IDENTIFICADOR);
				}
				break;
			case 24:
				enterOuterAlt(_localctx, 24);
				{
				setState(140);
				match(TOK_CONT);
				setState(141);
				match(TOK_IGUAL);
				setState(142);
				match(TOK_CADENA);
				}
				break;
			case 25:
				enterOuterAlt(_localctx, 25);
				{
				setState(143);
				match(TOK_CONT);
				setState(144);
				match(TOK_IGUAL);
				setState(145);
				match(TOK_CAMINO);
				}
				break;
			case 26:
				enterOuterAlt(_localctx, 26);
				{
				setState(146);
				match(TOK_RUTA);
				setState(147);
				match(TOK_IGUAL);
				setState(148);
				match(TOK_CADENA);
				}
				break;
			case 27:
				enterOuterAlt(_localctx, 27);
				{
				setState(149);
				match(TOK_RUTA);
				setState(150);
				match(TOK_IGUAL);
				setState(151);
				match(TOK_CAMINO);
				}
				break;
			case 28:
				enterOuterAlt(_localctx, 28);
				{
				setState(152);
				match(TOK_P);
				}
				break;
			case 29:
				enterOuterAlt(_localctx, 29);
				{
				setState(153);
				match(TOK_R);
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
		case 1:
			return comandList_sempred((ComandListContext)_localctx, predIndex);
		case 3:
			return param_list_sempred((Param_listContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean comandList_sempred(ComandListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean param_list_sempred(Param_listContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\64\u009f\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\7\3\30\n\3\f\3\16\3\33\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\5\4;\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\7\5E\n\5\f\5\16"+
		"\5H\13\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u009d"+
		"\n\6\3\6\2\4\4\b\7\2\4\6\b\n\2\2\2\u00c6\2\f\3\2\2\2\4\17\3\2\2\2\6:\3"+
		"\2\2\2\b<\3\2\2\2\n\u009c\3\2\2\2\f\r\5\4\3\2\r\16\b\2\1\2\16\3\3\2\2"+
		"\2\17\20\b\3\1\2\20\21\5\6\4\2\21\22\b\3\1\2\22\31\3\2\2\2\23\24\f\4\2"+
		"\2\24\25\5\6\4\2\25\26\b\3\1\2\26\30\3\2\2\2\27\23\3\2\2\2\30\33\3\2\2"+
		"\2\31\27\3\2\2\2\31\32\3\2\2\2\32\5\3\2\2\2\33\31\3\2\2\2\34\35\7\4\2"+
		"\2\35;\5\b\5\2\36\37\7\5\2\2\37;\5\b\5\2 !\7\6\2\2!;\5\b\5\2\"#\7\7\2"+
		"\2#;\5\b\5\2$%\7\b\2\2%;\5\b\5\2&\'\7\t\2\2\';\5\b\5\2()\7\13\2\2);\5"+
		"\b\5\2*+\7\f\2\2+;\5\b\5\2,-\7\r\2\2-;\5\b\5\2./\7\16\2\2/;\5\b\5\2\60"+
		"\61\7\17\2\2\61;\5\b\5\2\62\63\7\20\2\2\63;\5\b\5\2\64\65\7\21\2\2\65"+
		";\5\b\5\2\66\67\7\23\2\2\67;\5\b\5\28;\7\22\2\29;\7\n\2\2:\34\3\2\2\2"+
		":\36\3\2\2\2: \3\2\2\2:\"\3\2\2\2:$\3\2\2\2:&\3\2\2\2:(\3\2\2\2:*\3\2"+
		"\2\2:,\3\2\2\2:.\3\2\2\2:\60\3\2\2\2:\62\3\2\2\2:\64\3\2\2\2:\66\3\2\2"+
		"\2:8\3\2\2\2:9\3\2\2\2;\7\3\2\2\2<=\b\5\1\2=>\5\n\6\2>?\b\5\1\2?F\3\2"+
		"\2\2@A\f\4\2\2AB\5\n\6\2BC\b\5\1\2CE\3\2\2\2D@\3\2\2\2EH\3\2\2\2FD\3\2"+
		"\2\2FG\3\2\2\2G\t\3\2\2\2HF\3\2\2\2IJ\7\26\2\2JK\7\3\2\2K\u009d\7/\2\2"+
		"LM\7\24\2\2MN\7\3\2\2N\u009d\7\61\2\2OP\7\24\2\2PQ\7\3\2\2Q\u009d\7.\2"+
		"\2RS\7\25\2\2ST\7\3\2\2T\u009d\7)\2\2UV\7\25\2\2VW\7\3\2\2W\u009d\7*\2"+
		"\2XY\7\25\2\2YZ\7\3\2\2Z\u009d\7+\2\2[\\\7\27\2\2\\]\7\3\2\2]\u009d\7"+
		"%\2\2^_\7\27\2\2_`\7\3\2\2`\u009d\7#\2\2ab\7\27\2\2bc\7\3\2\2c\u009d\7"+
		"$\2\2de\7\30\2\2ef\7\3\2\2f\u009d\7\62\2\2gh\7\30\2\2hi\7\3\2\2i\u009d"+
		"\7.\2\2jk\7\33\2\2kl\7\3\2\2l\u009d\7\62\2\2mn\7\33\2\2no\7\3\2\2o\u009d"+
		"\7.\2\2pq\7\37\2\2qr\7\3\2\2r\u009d\7\62\2\2st\7\37\2\2tu\7\3\2\2u\u009d"+
		"\7.\2\2vw\7\34\2\2wx\7\3\2\2x\u009d\7\62\2\2yz\7\35\2\2z{\7\3\2\2{\u009d"+
		"\7\62\2\2|}\7\31\2\2}~\7\3\2\2~\u009d\7&\2\2\177\u0080\7\31\2\2\u0080"+
		"\u0081\7\3\2\2\u0081\u009d\7(\2\2\u0082\u0083\7\31\2\2\u0083\u0084\7\3"+
		"\2\2\u0084\u009d\7\'\2\2\u0085\u0086\7\31\2\2\u0086\u0087\7\3\2\2\u0087"+
		"\u009d\7,\2\2\u0088\u0089\7\31\2\2\u0089\u008a\7\3\2\2\u008a\u009d\7-"+
		"\2\2\u008b\u008c\7\32\2\2\u008c\u008d\7\3\2\2\u008d\u009d\7\60\2\2\u008e"+
		"\u008f\7\36\2\2\u008f\u0090\7\3\2\2\u0090\u009d\7.\2\2\u0091\u0092\7\36"+
		"\2\2\u0092\u0093\7\3\2\2\u0093\u009d\7\61\2\2\u0094\u0095\7 \2\2\u0095"+
		"\u0096\7\3\2\2\u0096\u009d\7.\2\2\u0097\u0098\7 \2\2\u0098\u0099\7\3\2"+
		"\2\u0099\u009d\7\61\2\2\u009a\u009d\7\"\2\2\u009b\u009d\7!\2\2\u009cI"+
		"\3\2\2\2\u009cL\3\2\2\2\u009cO\3\2\2\2\u009cR\3\2\2\2\u009cU\3\2\2\2\u009c"+
		"X\3\2\2\2\u009c[\3\2\2\2\u009c^\3\2\2\2\u009ca\3\2\2\2\u009cd\3\2\2\2"+
		"\u009cg\3\2\2\2\u009cj\3\2\2\2\u009cm\3\2\2\2\u009cp\3\2\2\2\u009cs\3"+
		"\2\2\2\u009cv\3\2\2\2\u009cy\3\2\2\2\u009c|\3\2\2\2\u009c\177\3\2\2\2"+
		"\u009c\u0082\3\2\2\2\u009c\u0085\3\2\2\2\u009c\u0088\3\2\2\2\u009c\u008b"+
		"\3\2\2\2\u009c\u008e\3\2\2\2\u009c\u0091\3\2\2\2\u009c\u0094\3\2\2\2\u009c"+
		"\u0097\3\2\2\2\u009c\u009a\3\2\2\2\u009c\u009b\3\2\2\2\u009d\13\3\2\2"+
		"\2\6\31:F\u009c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}